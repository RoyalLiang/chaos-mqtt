package area

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/cmd/chaos/internal/fms/area"
	"fms-awesome-tools/configs"

	"github.com/spf13/cobra"
)

var (
	hClear bool
	hAdd   bool
	hDel   bool
	hStart int64
	hEnd   int64
)

var HatchCoverCmd = &cobra.Command{
	Use:   "hatch_cover",
	Short: "锁定/解锁/清除 QC后大梁锁闭区",
	Run: func(cmd *cobra.Command, args []string) {
		if !hClear && !hAdd && !hDel {
			_ = cmd.Help()
			return
		}

		if hClear {
			sendData(fms.ClearBlockURL, make([]byte, 0))
		} else {
			if hStart <= 0 || hEnd <= 0 {
				cobra.CheckErr("invalid start or end wm")
			}
			operateHatchCover()
		}
	},
}

func operateHatchCover() {
	var op string
	if hAdd {
		op = "add"
	} else if hDel {
		op = "delete"
	}

	body := area.HatchCoverConfigRequest{
		Op:    op,
		Start: hStart,
		End:   hEnd,
	}.String()
	sendData(fms.OpBlockURL, []byte(body))
}

func sendData(url string, data []byte) {
	address := configs.Chaos.FMS.Area.Address
	resp, err := fms.Post(address+url, data)
	if err != nil {
		cobra.CheckErr(err)
	} else {
		fmt.Println(string(resp))
	}
}

func init() {
	HatchCoverCmd.Flags().BoolVar(&hClear, "clear", false, "清除所有的后大梁锁闭区🆑🚧")
	HatchCoverCmd.Flags().BoolVarP(&hAdd, "add", "a", false, "添加后大梁锁闭区➕🚧")
	HatchCoverCmd.Flags().BoolVarP(&hDel, "delete", "d", false, "删除后大梁锁闭区➖🚧")
	HatchCoverCmd.Flags().Int64VarP(&hStart, "start", "s", 0, "开始位置🔛")
	HatchCoverCmd.Flags().Int64VarP(&hEnd, "end", "e", 0, "结束位置🔚")
	HatchCoverCmd.MarkFlagsMutuallyExclusive("add", "delete", "clear")
	HatchCoverCmd.MarkFlagsRequiredTogether("start", "end")
	//HatchCoverCmd.MarkFlagsRequiredTogether("delete", "start", "end")

}
