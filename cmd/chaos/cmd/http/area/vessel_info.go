package area

import (
	"context"
	"encoding/json"
	"fms-awesome-tools/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/configs"

	"github.com/redis/go-redis/v9"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	keep   bool
	vid    string
	t      = table.NewWriter()
	header = table.Row{
		"Vessel ID", "D", "Berth", "Ingress", "Egress", "QC", "WM", "Locked", "Occupy", "QC Queue", "CA", "Work lane", "CA Status",
		"CA Capacity", "Ca Queue", "TCA Info", "TCA Seq",
	}
)

const (
	moveCursor    = "\033[s" // 保存光标位置
	restoreCursor = "\033[u" // 恢复光标位置
	clearScreen   = "\033[J" // 清除从光标到屏幕底部的内容
)

var GetVesselCmd = &cobra.Command{
	Use:   "vessels",
	Short: "获取所有船舶/指定船舶的CA状态及等待队列",
	Run: func(cmd *cobra.Command, args []string) {
		t.AppendHeader(header)
		if keep {
			fmt.Print(moveCursor)
			printVesselsForever()
		} else {
			if vessels := getVessels(); vessels != nil {
				printVessels(vessels.Data.Values)
			}
		}
	},
}

type vesselManager struct {
	sync.Mutex
	vessels map[string]fms.VesselInfo
}

func (vm *vesselManager) addVessel(v fms.VesselInfo) {
	vm.Lock()
	defer vm.Unlock()
	vm.vessels[v.VesselInfo.VesselId] = v
}

func (vm *vesselManager) getVessels() []fms.VesselInfo {
	vessels := make(fms.VesselsInfo, 0)
	for _, vs := range vm.vessels {
		vessels = append(vessels, vs)
	}
	sort.Sort(vessels)
	return vessels
}

func printVesselsForever() {
	var (
		ctx       = context.Background()
		msgChan   = make(chan *redis.Message, 100)
		sleepTime = time.Second * 2
		exitChan  = make(chan os.Signal, 1)
		manager   = &vesselManager{
			vessels: make(map[string]fms.VesselInfo),
		}
	)

	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)
	sub, err := service.Subscribe(ctx, "vessel_status", msgChan)
	if err != nil {
		cobra.CheckErr(err)
	}

	defer sub.Close()

	ticker := time.NewTicker(sleepTime)
	defer ticker.Stop()
	for {
		select {
		case msg, ok := <-msgChan:
			if !ok {
				return
			}
			vesselInfo := &fms.VesselInfo{}
			if err = json.Unmarshal([]byte(msg.Payload), vesselInfo); err != nil {
				cobra.CheckErr(err)
			}
			manager.addVessel(*vesselInfo)
		case <-ticker.C:
			fmt.Print("\033[u\033[J")
			printVessels(manager.getVessels())
		case <-exitChan:
			fmt.Print("\033[0;0H")
			return
		}
	}
}

func getAssignedCraneCaData(crane string, cas []fms.VesselCAInfo) []fms.VesselCAInfo {
	res := make([]fms.VesselCAInfo, 0)
	for _, c := range cas {
		if strings.HasPrefix(c.Name, crane) {
			res = append(res, c)
		}
	}
	return res
}

func getAssignedQCTCAInfo(crane string, tcaMapping map[string]fms.TCAMapping) string {
	mapping, ok := tcaMapping[crane]
	if ok && mapping.Capacity > 0 {
		return fmt.Sprintf("%d/%s", mapping.Capacity, strings.Join(mapping.Vehicles, ","))
	}
	return "-/-"
}

func printVessels(vessels fms.VesselsInfo) {
	t.ResetRows()
	//rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	for _, vs := range vessels {
		for _, crane := range vs.Cranes {
			rows := make([]table.Row, 0)
			cas := getAssignedCraneCaData(crane.Name, vs.CAs)
			for _, ca := range cas {
				tcaInfo := getAssignedQCTCAInfo(ca.Crane, vs.TCAMapping)
				row := table.Row{
					ca.VesselId, vs.VesselInfo.Direction, vs.Wms(), vs.Ingress.WharfMarkEnd, vs.Egress.WharfMarkEnd, crane.Name, crane.WharfMark, getLockedStatus(crane.Locked),
					crane.VehicleID, strings.Join(vs.CAArrives, ","), ca.Name, ca.GetWorkLane(),
					getLockedStatus(ca.Locked), ca.Capacity, strings.Join(ca.Vehicles, ","), tcaInfo, vs.TCACallSeq,
				}
				rows = append(rows, row)
			}
			t.AppendRows(rows)
		}
	}

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true}, {Number: 2, AutoMerge: true}, {Number: 3, AutoMerge: true}, {Number: 4, AutoMerge: true},
		{Number: 5, AutoMerge: true}, {Number: 6, AutoMerge: true}, {Number: 7, AutoMerge: true}, {Number: 8, AutoMerge: true},
		{Number: 12, AutoMerge: true}, {Number: 17, AutoMerge: true},
	})

	fmt.Println(t.Render())
}

func getVessels() *fms.GetVesselsResponse {
	address := configs.Chaos.FMS.Area.Address
	url := address + fms.GetVesselsURL
	if vid != "" {
		url += "?vessel_id=" + vid
	}

	resp, err := fms.Get(url)
	if err != nil {
		logger.Warnf("获取船舶信息失败: %s", err.Error())
		return nil
	}

	vesselInfo := &fms.GetVesselsResponse{}
	if err := json.Unmarshal(resp, vesselInfo); err != nil {
		logger.Warnf("解析船舶信息失败: %s", err.Error())
		return nil
	}
	logger.Infof("船舶信息获取成功: %s", vesselInfo.String())
	return vesselInfo
}

func getLockedStatus(status int) string {
	if status == 1 {
		return "Locked"
	}
	return ""
}

func init() {
	GetVesselCmd.Flags().BoolVarP(&keep, "keepalive", "k", false, "自动刷新🔄️️(1/2s)")
	GetVesselCmd.Flags().StringVarP(&vid, "vessel-id", "i", "", "船舶ID🚢")
}
