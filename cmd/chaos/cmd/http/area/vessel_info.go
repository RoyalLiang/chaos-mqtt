package area

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/configs"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	keep bool
	vid  string
	t    = table.NewWriter()
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
		header := table.Row{"VesselID", "CA", "Working lane", "Capacity", "CA Status", "Ca Queues", "QC Status", "QC Assigned", "QC Queues", "DWA Queues"}
		t.AppendHeader(header)

		if keep {
			// 保存初始光标位置
			fmt.Print(moveCursor)
			for {
				if vessels := getVessels(); vessels != nil {
					// 恢复到保存的位置并清屏
					fmt.Print(restoreCursor, clearScreen)
					parseVesselInfo(vessels.Data.Values)
				}
				time.Sleep(5 * time.Second)
			}
		} else {
			if vessels := getVessels(); vessels != nil {
				parseVesselInfo(vessels.Data.Values)
			}
		}
	},
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

func parseVesselInfo(vessels []fms.VesselInfo) {
	cas := make([]fms.VesselCAInfo, 0)

	for _, vessel := range vessels {
		if vessel.CAs == nil {
			continue
		}

		for _, ca := range vessel.CAs {
			cas = append(cas, ca)
		}

	}
	printResult(vessels, cas)
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
	GetVesselCmd.Flags().StringVarP(&vid, "vessel-id", "v", "", "船舶ID🚢")
}
