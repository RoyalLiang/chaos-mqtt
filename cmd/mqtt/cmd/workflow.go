package cmd

//var (
//	truck string
//	flow  string
//	//activity int64
//)
//
//var workflowCmd = &cobra.Command{
//	Use:   "workflow",
//	Short: "Start workflow testing",
//	Run: func(cmd *cobra.Command, args []string) {
//		if truck == "" && flow == "" {
//			cmd.Help()
//			return
//		}
//		fmt.Println("truck:", truck)
//		fmt.Println("flow:", flow)
//		fmt.Println("activity:", activity)
//
//		startWorkflow()
//	},
//}
//
//func startWorkflow() {
//	workflow := service.Workflow{
//		UUID:     uuid.NewString(),
//		Truck:    truck,
//		Flow:     flow,
//		Activity: activity,
//	}
//	fmt.Println("workflow object:", workflow)
//	if err := workflow.StartWorkflow(); err != nil {
//		fmt.Println("Failed to start workflow:", err)
//		os.Exit(1)
//	}
//}
//
//func init() {
//	workflowCmd.PersistentFlags().StringVarP(&truck, "truck", "t", "", "which truck used to testing")
//	workflowCmd.PersistentFlags().StringVarP(&flow, "flow", "f", "", "QC\nYARD\nSTANDBY\n")
//	workflowCmd.PersistentFlags().Int64VarP(&activity, "activity", "a", 0, activities)
//	workflowCmd.MarkFlagsRequiredTogether("truck", "flow", "activity")
//	rootCmd.AddCommand(workflowCmd)
//}
