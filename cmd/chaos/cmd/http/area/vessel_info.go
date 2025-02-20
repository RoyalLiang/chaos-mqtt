package area

import (
	"encoding/json"
	"fms-awesome-tools/cmd/chaos/internal/http"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

var (
	keep       bool
	vid        string
	firstPrint = true
	tableRows  = 0
)

var GetVesselCmd = &cobra.Command{
	Use:   "vessels_status",
	Short: "获取所有船舶/指定船舶的CA状态及等待队列",
	Run: func(cmd *cobra.Command, args []string) {
		if keep {
			for {
				if vessels := getVessels(); vessels != nil {
					parseVesselInfo(vessels.Data.Values)
				}
				time.Sleep(10 * time.Second)
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
	// 计算每列的最大宽度
	colWidths := make([]int, 9)
	// 设置表头宽度作为初始值
	colWidths[0] = 8  // 船舶ID
	colWidths[1] = 12 // CA
	colWidths[2] = 2  // 容量
	colWidths[3] = 6  // 锁定状态
	colWidths[4] = 2  // 绑定车道
	colWidths[5] = 6  // QC状态
	colWidths[6] = 10 // QC队列
	colWidths[7] = 32 // 集卡队列
	colWidths[8] = 40 // 等待队列

	border := "="
	header := ""
	for i, width := range colWidths {
		border += strings.Repeat("=", width) + "="
		headerText := []string{"船舶ID", "CA", "容量", "锁定状态", "绑定车道", "QC状态", "QC队列", "集卡队列", "等待队列"}[i]
		header += fmt.Sprintf(" %-*s ", width-1, headerText) + "|"
	}

	h := strings.Split(header, "|")

	// 如果不是首次打印，移动光标到表格开始位置并清除表格区域
	if !firstPrint {
		// 移动光标到表格开始位置（上移tableRows行）
		fmt.Printf("\033[%dA", tableRows)
		// 清除从光标到屏幕底部的内容
		fmt.Print("\033[J")
	}

	// 打印表格
	fmt.Println(border + "==========================")
	fmt.Println(strings.Join(h[0:len(h)-1], "|"))
	fmt.Println(border + "==========================")

	// 打印数据行
	for _, ca := range cas {
		var bindLane int
		crane := getAssignedQCData(vessels, ca.Crane)

		if ca.FixedWorkLane != nil {
			bindLane = *ca.FixedWorkLane
		} else {
			bindLane = ca.BindLane
		}

		fmt.Printf("| %-*s | %-*s | %-*d | %-*s | %-*d | %-*s | %-*s | %-*s | %-*s |\n",
			colWidths[0]-1, ca.VesselId,
			colWidths[1]-1, ca.Name,
			colWidths[2]+1, ca.Capacity,
			colWidths[3]+2, getLockedStatus(ca.Locked),
			colWidths[4]+5, bindLane,
			colWidths[5], getLockedStatus(crane.Locked),
			colWidths[6], crane.VehicleID,
			colWidths[7]+2, strings.Join(ca.Vehicles, ","),
			colWidths[8], "")
	}
	fmt.Println(border + "==========================")

	// 更新表格行数和首次打印标志
	tableRows = len(cas) + 4 // 表头2行 + 数据行 + 底部边框1行
	firstPrint = false
}

func init() {
	GetVesselCmd.Flags().BoolVarP(&keep, "keepalive", "k", false, "是否保持刷新(refresh every 5s)")
	GetVesselCmd.Flags().StringVarP(&vid, "vessel-id", "v", "", "船舶ID")
}
