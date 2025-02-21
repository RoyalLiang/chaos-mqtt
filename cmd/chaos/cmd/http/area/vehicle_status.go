package area

import (
	"encoding/json"
	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
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
		header := table.Row{"ID", "Vehicle ID", "Task Type", "Current Destination", "Current Arrived", "Destination", "Destination Lane", "Call Status"}
		vehicleTable.AppendHeader(header)

		if keep {
			fmt.Print(moveCursor)
			for {
				vehicles := getVehicles()
				fmt.Print(restoreCursor, clearScreen)
				printVehicles(vehicles)
				time.Sleep(5 * time.Second)
			}
		} else {
			if vessels := getVessels(); vessels != nil {
				parseVesselInfo(vessels.Data.Values)
			}
		}
	},
}

func getVehicles() []fms.VehiclesResponseData {
	var address string

	if configs.Chaos.FMS == nil {
		address = ""
	} else {
		for _, service := range configs.Chaos.FMS.Services {
			if service.Name == "area" {
				address = service.Address
				break
			}
		}
	}

	if address == "" {
		address = "http://10.1.205.3:8888"
	}

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
	if err := json.Unmarshal(resp, respData); err != nil {
		fmt.Println("解析集卡数据失败: ", err.Error())
		return nil
	}
	return respData.Data
}

func printVehicles(vehicles []fms.VehiclesResponseData) {
	vehicleTable.ResetRows()
	for index, vehicle := range vehicles {

		called := ""
		if vehicle.CanGoCallIn {
			called = "Called"
		}

		arrived := "On the way"
		if vehicle.Arrived {
			arrived = "Arrived"
		}

		row := table.Row{
			index + 1, vehicle.ID, vehicle.Destination.Type, vehicle.CurrentDestination.Name, arrived,
			vehicle.Destination.Name, vehicle.Destination.Lane, called,
		}
		vehicleTable.AppendRow(row)
	}
	fmt.Println(t.Render())
}

func init() {
	VehicleCmd.Flags().BoolVarP(&k, "keepalive", "k", false, "是否保持刷新(F/5s)")
	VehicleCmd.Flags().StringVarP(&vehicleID, "vehicle", "v", "", "集卡号")
}
