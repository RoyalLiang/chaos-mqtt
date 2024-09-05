package topics

import (
	"fmt"
	"strings"

	"fms-awesome-tools/cmd/mqtt/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	success  int64
	trailers string
)

var LogonResponseCmd = &cobra.Command{
	Use:   "logon_response",
	Short: "发送 logon_response",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("logon_response", constants.LogonResponse, genLogonResponseParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func genLogonResponseParam() constants.LogonResponseParam {
	return constants.LogonResponseParam{
		VehicleID: constants.VehicleID,
		Success:   success,
		Trailers:  strings.Split(trailers, ","),
	}
}

func init() {
	LogonResponseCmd.Flags().Int64VarP(&success, "success", "s", 1, "登录成功; 1: 成功, 0: 失败")
	LogonResponseCmd.Flags().StringVarP(&trailers, "trailers", "t", "C53525", "拖车号; 多个拖车号用逗号隔开")
}
