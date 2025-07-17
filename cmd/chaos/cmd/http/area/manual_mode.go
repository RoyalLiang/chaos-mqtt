package area

import (
	"fmt"
	"strconv"
	"strings"

	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/cmd/chaos/internal/fms/area"
	"fms-awesome-tools/configs"

	"github.com/spf13/cobra"
)

var (
	vesselID string
	ingress  int64
	egress   int64
	qcs      []string
	turns    []string
	eMapping []string
	reset    bool
)

var ManualModeCmd = &cobra.Command{
	Use:   "manual_mode",
	Short: "手动设置船舶相关配置",
	Run: func(cmd *cobra.Command, args []string) {
		if vesselID == "" && ingress < 0 && egress < 0 && len(qcs) == 0 {
			_ = cmd.Help()
			return
		}

		if vesselID == "" && (ingress >= 0 || egress >= 0 || len(qcs) > 0 || reset) {
			cobra.CheckErr("未指定船舶ID")
		}

		if reset {
			resetRequest()
		} else {
			manualRequest()
		}
	},
}

func resetRequest() {
	url := "/fms/psa/vessel/" + vesselID + "/reset"
	sendRequest(url, make([]byte, 0))
}

func split(mapping []string) map[string]string {
	result := make(map[string]string)
	for _, item := range mapping {
		parts := strings.Split(item, "=")
		if len(parts) != 2 {
			cobra.CheckErr("无效的输入格式: %s，应为 <lane>=<direction>")
		}
		result[parts[0]] = parts[1]
	}
	return result
}

func manualRequest() {
	qcLaneMap := make(map[string]int64)
	ingressTurns := make(map[string]string)
	egressTurns := make(map[string]string)
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

	if len(turns) > 0 {
		ingressTurns = split(turns)
	}

	if len(eMapping) > 0 {
		egressTurns = split(eMapping)
	}

	body := area.ManualModeRequest{
		Ingress:     ingress,
		Egress:      egress,
		QCLanes:     qcLaneMap,
		TurnMapping: ingressTurns,
		EMapping:    egressTurns,
		Mode:        1,
	}
	sendRequest(fmt.Sprintf(fms.ManualModeURL, vesselID), []byte(body.String()))
}

func sendRequest(url string, data []byte) {
	address := configs.Chaos.FMS.Area.Address
	url = address + url
	resp, err := fms.Post(url, data)
	if err != nil {
		cobra.CheckErr(err)
	} else {
		fmt.Println(string(resp))
	}
}

func init() {
	ManualModeCmd.Flags().BoolVar(&reset, "reset", false, "重置船舶模式🆑")
	ManualModeCmd.Flags().StringVarP(&vesselID, "vessel-id", "v", "", "船舶ID🚢")
	ManualModeCmd.Flags().Int64VarP(&ingress, "ingress", "i", 0, "指定的ingress wm🚩")
	ManualModeCmd.Flags().Int64VarP(&egress, "egress", "e", 0, "指定的egress wm🚩")
	ManualModeCmd.Flags().StringSliceVarP(&qcs, "qc-config", "c", []string{}, "批量设置数据，格式: QC1=2🌉")
	ManualModeCmd.Flags().StringSliceVarP(&turns, "turn-mapping", "t", []string{}, "设置ingress不同车道的转向方式\n格式: 2=left🌌\n可选方向: left, right, ''\n")
	ManualModeCmd.Flags().StringSliceVarP(&eMapping, "egress-mapping", "o", []string{}, "设置egress不同车道的转向方式\n格式: 2=left🌌\n可选方向: left, right, ''\n")
}
