package area

import (
	"context"
	"encoding/json"
	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/configs"
	"fms-awesome-tools/constants"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"
)

const (
	moveCursor    = "\033[s" // 保存光标位置
	restoreCursor = "\033[u" // 恢复光标位置
	clearScreen   = "\033[J" // 清除从光标到屏幕底部的内容
)

var (
	vehicleID    string
	k            bool
	vehicleTable = table.NewWriter()
	redisClient  *redis.Client
)

var VehicleCmd = &cobra.Command{
	Use:   "vehicles",
	Short: "获取所有/指定集卡状态",
	Run: func(cmd *cobra.Command, args []string) {
		header := table.Row{"ID", "Vehicle ID", "Task Type", "Current Destination", "Destination Type", "Current Arrived", "Destination", "Destination Lane", "Call Status", "Mode", "Ready Status", "Manual Status", "SSA"}
		vehicleTable.AppendHeader(header)

		if !k && constants.VehicleID == "" {
			_ = cmd.Help()
			return
		}

		var err error
		redisClient, err = service.NewRedisClient()
		if err != nil {
			cobra.CheckErr(err)
		}

		if k {
			fmt.Print(moveCursor)
			subs()
		} else {
			vehicles := getVehicles()
			printVehicles(context.Background(), vehicles)
		}
	},
}

type VehicleManager struct {
	sync.Mutex
	vehicles map[string]*fms.VehiclesResponseData
}

func (vm *VehicleManager) Add(vehicle *fms.VehiclesResponseData) {
	vm.Lock()
	defer vm.Unlock()
	vm.vehicles[vehicle.ID] = vehicle
}

func (vm *VehicleManager) GetSorted() fms.Vehicles {
	vm.Lock()
	defer vm.Unlock()

	vehicles := make(fms.Vehicles, 0, len(vm.vehicles))
	for _, v := range vm.vehicles {
		vehicles = append(vehicles, *v)
	}
	sort.Sort(vehicles)
	return vehicles
}

func subs() {

	var (
		ctx          = context.Background()
		msgChan      = make(chan *redis.Message, 100)
		batchTimeout = time.Second
		manager      = &VehicleManager{
			vehicles: make(map[string]*fms.VehiclesResponseData),
		}
	)

	sub := redisClient.Subscribe(ctx, "vehicle_status")
	defer sub.Close()
	go func() {
		for {
			msg, err := sub.ReceiveMessage(ctx)
			if err != nil {
				fmt.Println("subs error:", err.Error())
				close(msgChan)
				return
			}
			msgChan <- msg
		}
	}()

	var messages []*redis.Message
	ticker := time.NewTicker(batchTimeout)
	defer ticker.Stop()

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case msg, ok := <-msgChan:
			if !ok {
				return
			}
			vehicle := &fms.VehiclesResponseData{}
			if err := json.Unmarshal([]byte(msg.Payload), vehicle); err == nil {
				manager.Add(vehicle)
			}
		case <-ticker.C:
			if len(messages) == 0 {
				continue
			}

			vehicles := manager.GetSorted()
			fmt.Print("\033[u\033[J")
			printVehicles(ctx, vehicles)
		case <-exitChan:
			fmt.Print("\033[0;0H") // 复位光标
			return
		}
	}
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

func printVehicles(ctx context.Context, vehicles fms.Vehicles) {
	vehicleTable.ResetRows()
	sort.Sort(vehicles)
	for index, vehicle := range vehicles {

		modeData, _ := redisClient.HGet(ctx, "psa_vehicle_status", vehicle.ID).Result()
		_ = json.Unmarshal([]byte(modeData), &vehicle)

		called := ""
		if vehicle.CanGoCallIn {
			called = "Called"
		}

		arrived := "On the way"
		if vehicle.Arrived {
			arrived = "Arrived"
		}

		ssa := ""
		if vehicle.SSA == 1 {
			ssa = "ON"
		}

		ready := ""
		if vehicle.ReadyStatus == 0 {
			ready = "OFF"
		}

		manual := ""
		if vehicle.ManualStatus == 1 {
			manual = "ON"
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
			index + 1, vehicle.ID, vehicle.Destination.Type, name, dtype, arrived, vehicle.Destination.Name,
			vehicle.Destination.Lane, called, vehicle.Mode, ready, manual, ssa,
		}
		vehicleTable.AppendRow(row)
	}
	fmt.Print(vehicleTable.Render())
}

func init() {
	VehicleCmd.Flags().BoolVarP(&k, "keepalive", "k", false, "是否保持刷新(F/5s)")
	VehicleCmd.Flags().StringVarP(&vehicleID, "vehicle", "v", "", "集卡号")
}
