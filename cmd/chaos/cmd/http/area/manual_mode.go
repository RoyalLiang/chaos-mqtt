package area

import (
	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/cmd/chaos/internal/fms/area"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var (
	vesselID string
	ingress  int64
	egress   int64
	qcs      []string
)

var ManualModeCmd = &cobra.Command{
	Use:   "manual_mode",
	Short: "手动设置船舶相关配置",
	Run: func(cmd *cobra.Command, args []string) {
		if ingress < 0 && egress < 0 && len(qcs) == 0 {
			_ = cmd.Help()
			return
		}

		qcLaneMap := make(map[string]int64)
		if len(qcs) > 0 {
			for _, item := range qcs {
				parts := strings.Split(item, "=")
				if len(parts) != 2 {
					fmt.Printf("无效的输入格式: %s，应为 QC=lane\n", item)
					return
				}
				lane, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					fmt.Printf("lane解析失败: %v\n", err)
					return
				}
				qcLaneMap[parts[0]] = lane
			}
		}

		sendRequest(qcLaneMap)
	},
}

func sendRequest(qcLaneMap map[string]int64) {
	var address string
	for _, service := range configs.Chaos.FMS.Services {
		if service.Name == "area" {
			address = service.Address
			break
		}
	}

	url := address + "/fms/psa/vessel/" + vesselID + "/manualModel"
	body := area.ManualModeRequest{
		Ingress: ingress,
		Egress:  egress,
		QCLanes: qcLaneMap,
		Mode:    1,
	}
	resp, err := fms.Post(url, []byte(body.String()))
	if err != nil {
		fmt.Println("模式设置失败: ", err.Error())
	} else {
		fmt.Println(string(resp))
	}

}

func init() {
	ManualModeCmd.Flags().StringVarP(&vesselID, "vessel-id", "v", "", "船舶ID")
	ManualModeCmd.Flags().Int64VarP(&ingress, "ingress", "i", 0, "指定的ingress wharf mark")
	ManualModeCmd.Flags().Int64VarP(&egress, "egress", "e", 0, "指定的egress wharf mark")
	ManualModeCmd.Flags().StringSliceVarP(&qcs, "qc-config", "c", []string{}, "批量设置数据，格式: QC1=2")
	_ = ManualModeCmd.MarkFlagRequired("vessel-id")
}
