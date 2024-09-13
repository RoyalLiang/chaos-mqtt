package topics

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"
)

var LogonRequestCmd = &cobra.Command{
	Use:   "logon_request",
	Short: "发送 logon_request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("update_trailer", "", genLogonRequestParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func genLogonRequestParam() string {
	ts := strings.Split(trailers, ",")
	return messages.LogonRequest{
		APMID: constants.VehicleID,
		Data: messages.LogonRequestData{
			TrailerNumbers: ts, TrailerSeqNumbers: make([]int, 0),
		},
	}.String()
}

func init() {
	LogonRequestCmd.Flags().Int64VarP(&success, "success", "s", 1, "登录成功; 1: 成功, 0: 失败")
	LogonRequestCmd.Flags().StringVarP(&trailers, "trailers", "t", "C53525", "拖车号; 多个拖车号用逗号隔开")
}
