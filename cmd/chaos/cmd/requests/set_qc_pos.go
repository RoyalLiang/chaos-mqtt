package requests

import (
	"github.com/spf13/cobra"
)

var QCPosCmd = &cobra.Command{
	Use:   "set_qc_position",
	Short: "set qc position manually",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
