package area

import (
	"encoding/json"
	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"sort"
	"time"
)

var (
	vehicleID    string
	k            bool
	vehicleTable = table.NewWriter()
)

var VehicleCmd = &cobra.Command{
	Use:   "vehicles",
	Short: "获取所有/指定集卡状态",
	Run: func(cmd *cobra.Command, args []string) {
		header := table.Row{"ID", "Vehicle ID", "Task Type", "Current Destination", "Destination Type", "Current Arrived", "Destination", "Destination Lane", "Call Status"}
		vehicleTable.AppendHeader(header)

		if k {
			fmt.Print(moveCursor)
			for {
				vehicles := getVehicles()
				fmt.Print(restoreCursor, clearScreen)
				printVehicles(vehicles)
				time.Sleep(5 * time.Second)
			}
		} else {
			vehicles := getVehicles()
			printVehicles(vehicles)
		}
	},
}

func getVehicles() fms.Vehicles {
	address := configs.Chaos.FMS.Area.Address
	url := address + fms.GetVehiclesURL
	if vehicleID != "" {
		url += "?vehicle_id=" + vehicleID
	}

	resp, err := fms.Get(url)
	if err != nil {
		fmt.Println("获取集卡数据失败: ", err.Error())
		return nil
	}

	respData := &fms.VehiclesResponse{}
	if err = json.Unmarshal(resp, respData); err != nil {
		fmt.Println("解析集卡数据失败: ", err.Error())
		return nil
	}
	return respData.Data
}

func printVehicles(vehicles fms.Vehicles) {
	vehicleTable.ResetRows()
	sort.Sort(vehicles)
	for index, vehicle := range vehicles {

		called := ""
		if vehicle.CanGoCallIn {
			called = "Called"
		}

		arrived := "On the way"
		if vehicle.Arrived {
			arrived = "Arrived"
		}

		name := vehicle.CurrentDestination.Name
		if vehicle.CurrentDestination.Type == "Pre-Ingress" {
			name = vehicle.CurrentDestination.Type
		}
		if vehicle.Destination.Type == "YARD" {
			name = vehicle.Destination.Name
		}

		dtype := vehicle.CurrentDestination.Type
		switch vehicle.CurrentDestination.Type {
		case "CRANE_AREA":
			dtype = "QC"
		case "CALLIN_AREA":
			dtype = "CA"
		case "WAIT_AREA":
			dtype = "DWA"
		case "Pre-Ingress":
			dtype = "Pre-Ingress"
		default:
			dtype = vehicle.Destination.Type
		}
		row := table.Row{
			index + 1, vehicle.ID, vehicle.Destination.Type, name, dtype,
			arrived, vehicle.Destination.Name, vehicle.Destination.Lane, called,
		}
		vehicleTable.AppendRow(row)
	}
	fmt.Println(vehicleTable.Render())
}

func init() {
	VehicleCmd.Flags().BoolVarP(&k, "keepalive", "k", false, "是否保持刷新(F/5s)")
	VehicleCmd.Flags().StringVarP(&vehicleID, "vehicle", "v", "", "集卡号")
}
