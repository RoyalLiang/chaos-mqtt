package topics

import (
	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fmt"
	"strings"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	success  int64
	trailers string
	tp       string
)

var LogonResponseCmd = &cobra.Command{
	Use:   "logon_response",
	Short: "发送 logon_response",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("logon_response", "", genLogonResponse()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func genLogonResponse() string {
	return messages.LogonResponse{
		APMID: constants.VehicleID,
		Data: messages.LogonResponseData{
			Success: success, TrailerSeqNumbers: []int{1}, TrailerLengths: []int{20}, TrailerUnladenWeights: []int{11},
			TrailerTypes: strings.Split(tp, ","), TrailerPayloads: []int{200}, TrailerWidths: make([]int, 0),
			TrailerHeights: make([]int, 0), TrailerNumbers: strings.Split(trailers, ","),
		},
	}.String()
}

func init() {
	LogonResponseCmd.Flags().Int64VarP(&success, "success", "s", 1, "登录成功; 1: 成功, 0: 失败")
	LogonResponseCmd.Flags().StringVarP(&trailers, "trailers", "n", "C53525", "拖车号; 多个拖车号用逗号隔开")
	LogonResponseCmd.Flags().StringVarP(&tp, "trailer-types", "t", "CST", "拖车类型; 多个拖车号用逗号隔开")
}
