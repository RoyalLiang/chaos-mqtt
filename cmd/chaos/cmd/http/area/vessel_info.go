package area

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/redis/go-redis/v9"

	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/configs"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	keep   bool
	vid    string
	t      = table.NewWriter()
	header = table.Row{
		"Vessel ID", "Ingress WM", "Egress WM", "QC", "CA", "Working lane", "CA Status", "CA Capacity", "Ca Queues",
		"QC Status", "QC Assigned", "QC Queues", "DWA Queues",
	}
)

const (
	moveCursor    = "\033[s" // 保存光标位置
	restoreCursor = "\033[u" // 恢复光标位置
	clearScreen   = "\033[J" // 清除从光标到屏幕底部的内容
)

var GetVesselCmd = &cobra.Command{
	Use:   "vessels_status",
	Short: "获取所有船舶/指定船舶的CA状态及等待队列",
	Run: func(cmd *cobra.Command, args []string) {
		if !keep && vesselID == "" {
			_ = cmd.Help()
			return
		}

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
	resp := make([]fms.VesselInfo, 0)
	for _, vs := range vm.vessels {
		resp = append(resp, vs)
	}
	return resp
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
			if err = json.Unmarshal([]byte(msg.Payload), vesselInfo); err == nil {
				manager.addVessel(*vesselInfo)
			}
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

func printVessels(vessels []fms.VesselInfo) {
	t.ResetRows()
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	for _, vs := range vessels {
		for _, crane := range vs.Cranes {
			rows := make([]table.Row, 0)
			cas := getAssignedCraneCaData(crane.Name, vs.CAs)
			for _, ca := range cas {
				row := table.Row{
					ca.VesselId, vs.Ingress.WharfMarkStart, vs.Egress.WharfMarkEnd, crane.Name, ca.Name, ca.GetWorkLane(),
					getLockedStatus(ca.Locked), ca.Capacity, strings.Join(ca.Vehicles, ","),
					getLockedStatus(crane.Locked), crane.VehicleID, "", "",
				}
				rows = append(rows, row)
			}
			t.AppendRows(rows, rowConfigAutoMerge)
		}
	}
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true, VAlign: text.VAlignMiddle, Align: text.AlignCenter},
		{Number: 2, AutoMerge: true, VAlign: text.VAlignMiddle, Align: text.AlignCenter},
		{Number: 3, AutoMerge: true, VAlign: text.VAlignMiddle, Align: text.AlignCenter},
		{Number: 4, AutoMerge: true, VAlign: text.VAlignMiddle, Align: text.AlignCenter},
	})

	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	fmt.Print(t.Render())
}

func getVessels() *fms.GetVesselsResponse {
	address := configs.Chaos.FMS.Area.Address
	url := address + fms.GetVesselsURL
	if vid != "" {
		url += "?vessel_id=" + vid
	}

	resp, err := fms.Get(url)
	if err != nil {
		fmt.Println("获取船舶信息失败: ", err.Error())
		return nil
	}

	vesselInfo := &fms.GetVesselsResponse{}
	if err := json.Unmarshal(resp, vesselInfo); err != nil {
		fmt.Println("解析船舶信息失败: ", err.Error())
		return nil
	}
	return vesselInfo
}

func getAssignedQCData(vessels []fms.VesselInfo, craneNo string) fms.VesselCraneInfo {
	for _, vs := range vessels {
		for _, c := range vs.Cranes {
			if c.Name == craneNo {
				return c
			}
		}
	}
	return fms.VesselCraneInfo{}
}

func getLockedStatus(status int) string {
	if status == 1 {
		return "Locked"
	}
	return ""
}

func printResult(vessels []fms.VesselInfo, cas []fms.VesselCAInfo) {
	t.ResetRows()
	for _, ca := range cas {
		var bindLane int
		crane := getAssignedQCData(vessels, ca.Crane)

		if ca.FixedWorkLane != nil {
			bindLane = *ca.FixedWorkLane
		} else {
			bindLane = ca.BindLane
		}

		row := table.Row{
			ca.VesselId, ca.Name, bindLane, ca.Capacity, getLockedStatus(ca.Locked), strings.Join(ca.Vehicles, ","),
			getLockedStatus(crane.Locked), crane.VehicleID, "", "",
		}
		t.AppendRow(row)
	}

	fmt.Print(t.Render())
}

func init() {
	GetVesselCmd.Flags().BoolVarP(&keep, "keepalive", "k", false, "自动刷新🔄️️(1/5s)")
	GetVesselCmd.Flags().StringVarP(&vid, "vessel-id", "i", "", "船舶ID🚢")
}
