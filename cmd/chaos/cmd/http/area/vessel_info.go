package area

import (
	"encoding/json"
	"fms-awesome-tools/cmd/chaos/internal/http"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

var (
	keep bool
	vid  string
	t    = table.NewWriter()
)

const (
	moveCursor = "\033[H"
)

var GetVesselCmd = &cobra.Command{
	Use:   "vessels_status",
	Short: "获取所有船舶/指定船舶的CA状态及等待队列",
	Run: func(cmd *cobra.Command, args []string) {
		header := table.Row{"VesselID", "CA", "Capacity", "CA Status", "Working lane", "QC Status", "QC Queue", "CA Queue", "DWA Queue"}
		t.AppendHeader(header)

		if keep {
			for {
				if vessels := getVessels(); vessels != nil {
					parseVesselInfo(vessels.Data.Values)
				}
				time.Sleep(10 * time.Second)
				// 移动光标到起始位置
				fmt.Print(moveCursor)
			}
		} else {
			if vessels := getVessels(); vessels != nil {
				parseVesselInfo(vessels.Data.Values)
			}
		}
	},
}

func getVessels() *http.GetVesselsResponse {
	var address string

	if configs.Chaos.FMS == nil {
		address = ""
	} else {
		for _, service := range configs.Chaos.FMS.Services {
			if service.Name == "area" {
				address = service.Address
				break
			}
		}
	}

	if address == "" {
		address = "http://10.1.205.3:8888"
	}

	url := address + http.GetVesselsURL
	if vid != "" {
		url += "?vessel_id=" + vid
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取船舶信息失败: ", err.Error())
		return nil
	}

	vesselInfo := &http.GetVesselsResponse{}
	if err := json.Unmarshal(resp, vesselInfo); err != nil {
		fmt.Println("解析船舶信息失败: ", err.Error())
		return nil
	}
	return vesselInfo
}

func parseVesselInfo(vessels []http.VesselInfo) {
	cas := make([]http.VesselCAInfo, 0)

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

func getAssignedQCData(vessels []http.VesselInfo, craneNo string) http.VesselCraneInfo {
	for _, vs := range vessels {
		for _, c := range vs.Cranes {
			if c.Name == craneNo {
				return c
			}
		}
	}
	return http.VesselCraneInfo{}
}

func getLockedStatus(status int) string {
	if status == 1 {
		return "Locked"
	}
	return ""
}

func printResult(vessels []http.VesselInfo, cas []http.VesselCAInfo) {
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
			ca.VesselId, ca.Name, ca.Capacity, getLockedStatus(ca.Locked),
			bindLane, getLockedStatus(crane.Locked), crane.VehicleID, strings.Join(ca.Vehicles, ","), "",
		}

		t.AppendRow(row)
	}

	// 清除当前行到屏幕底部
	//fmt.Print("\033[J")
	fmt.Print(t.Render())
}

func init() {
	GetVesselCmd.Flags().BoolVarP(&keep, "keepalive", "k", false, "是否保持刷新(refresh every 5s)")
	GetVesselCmd.Flags().StringVarP(&vid, "vessel-id", "v", "", "船舶ID")
}
