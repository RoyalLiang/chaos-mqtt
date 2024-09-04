package topics

import "github.com/spf13/cobra"

var JobInstructionCmd = &cobra.Command{
	Use:   "job_instruction",
	Short: "发送job_instruction",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
