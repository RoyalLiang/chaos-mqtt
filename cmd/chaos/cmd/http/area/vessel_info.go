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
		printTable()

		if keep {
			for {
				if vessels := getVessels(); vessels != nil {
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

func getVessels() *http.GetVesselsResponse {
	var baseUrl string
	for _, service := range configs.Chaos.FMS.Services {
		if service.Name == "area" {
			baseUrl = service.BaseUrl
			break
		}
	}

	resp, err := http.Get(baseUrl + http.GetVesselsURL)
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

func printTable() {

}

func printResult(_ []http.VesselInfo, cas []http.VesselCAInfo) {
	// 计算每列的最大宽度
	colWidths := make([]int, 7)
	// 设置表头宽度作为初始值
	colWidths[0] = 8
	colWidths[1] = 12
	colWidths[2] = 6
	colWidths[3] = 6
	colWidths[4] = 6
	colWidths[5] = 32
	colWidths[6] = 40

	border := "="
	header := ""
	for i, width := range colWidths {
		border += strings.Repeat("=", width) + "="
		headerText := []string{"船舶ID", "CA", "容量", "锁定状态", "绑定车道", "集卡队列", "等待队列"}[i]
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
	fmt.Print(border + "=\n")
	fmt.Print(strings.Join(h[0:len(h)-1], "|") + "\n")
	fmt.Print(border + "=\n")

	// 打印数据行
	for _, ca := range cas {
		lockStatus := "未锁定"
		if ca.Locked == 1 {
			lockStatus = "已锁定"
		}

		// 使用ANSI颜色代码设置背景色
		fmt.Printf("\033[%dm| %-*s | %-*s | %-*d | %-*s | %-*d | %-*s | %-*s |\033[0m\n",
			0,
			colWidths[0]-1, ca.VesselId,
			colWidths[1]-1, ca.Name,
			colWidths[2]+1, ca.Capacity,
			colWidths[3]-1, lockStatus,
			colWidths[4]+2, ca.BindLane,
			colWidths[5]+1, strings.Join(ca.Vehicles, ","),
			colWidths[6], "")
	}
	fmt.Print(border + "=\n")

	// 更新表格行数和首次打印标志
	tableRows = len(cas) + 4 // 表头3行 + 数据行 + 底部边框1行
	firstPrint = false
}

func init() {
	GetVesselCmd.Flags().BoolVarP(&keep, "keepalive", "k", false, "是否保持刷新(refresh every 5s)")
	GetVesselCmd.Flags().StringVarP(&vid, "vessel-id", "v", "", "船舶ID")
}
