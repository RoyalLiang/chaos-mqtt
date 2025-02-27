package area

import (
	"context"
	"encoding/json"
	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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

const (
	taskTypes     = "\nQC\nYARD (STANDBY)\nIYS\n"
	HatchCoverOps = "psa_hatch_cover_ops"
	GetTaskInfo   = "psa_task_info"
)

var (
	vehicleID     string
	k             bool
	vehicleReset  bool
	vehicleTable  = table.NewWriter()
	redisClient   *redis.Client
	vehicleFilter string
)

var VehicleCmd = &cobra.Command{
	Use:   "vehicles",
	Short: "获取所有/指定集卡状态",
	Run: func(cmd *cobra.Command, args []string) {
		header := table.Row{
			"ID", "Vehicle ID", "Task Type", "ISO", "Dock Position", "Start Time", "Destination", "Lane",
			"Curr Destination", "Curr Type", "Arrived", "Call Status", "Mode", "Ready Status", "Manual",
		}
		vehicleTable.AppendHeader(header)

		if vehicleReset {
			if vehicleID == "" {
				cobra.CheckErr("[集卡重置] 缺失集卡号...")
			}
			resetVehicle()
			return
		}

		if k {
			fmt.Print(moveCursor)
			subs()
		} else {
			vehicles := getVehicles()
			printVehicles(context.Background(), vehicles)
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !k && vehicleID == "" && !vehicleReset {
			_ = cmd.Help()
			return
		}

		var err error
		redisClient, err = service.NewRedisClient()
		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

type VehicleManager struct {
	sync.Mutex
	vehicles map[string]*fms.VehiclesResponseData
}

func (vm *VehicleManager) Add(vehicle *fms.VehiclesResponseData) {
	if vehicleID != "" && vehicle.ID != vehicleID {
		return
	}

	if (vehicle.ID == "AT001" || vehicle.ID == "AT002") && vehicleID == "" {
		return
	}

	if vehicleFilter != "" && vehicle.Destination.Type != vehicleFilter {
		if _, ok := vm.vehicles[vehicleID]; ok {
			delete(vm.vehicles, vehicleID)
		}
		return
	}

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

func resetVehicle() {
	address := configs.Chaos.FMS.Area.Address
	url := address + fmt.Sprintf("%s/%s/clear", fms.ResetVehicleURL, vehicleID)
	resp, err := fms.Post(url, make([]byte, 0))
	if err != nil {
		cobra.CheckErr(err)
	}
	fmt.Println(string(resp))
}

func subs() {

	var (
		ctx          = context.Background()
		msgChan      = make(chan *redis.Message, 100)
		batchTimeout = time.Second * 2
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
			vehicle := &fms.VehiclesResponseData{TaskInfo: &fms.VehicleTaskInfo{}}
			if err := json.Unmarshal([]byte(msg.Payload), vehicle); err == nil {
				manager.Add(vehicle)
			}
		case <-ticker.C:
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

		arrived := ""
		if vehicle.Arrived {
			arrived = "Arrived"
		} else if vehicle.Destination.Name != "" {
			arrived = "On the way"
		}

		//ssa := ""
		//if vehicle.SSA == 1 {
		//	ssa = "ON"
		//}

		ready := ""
		if vehicle.ReadyStatus == 0 {
			ready = "OFF"
		}

		manual := ""
		if vehicle.ManualStatus == 1 {
			manual = "ON"
		}

		lane := ""
		if vehicle.Destination.Lane >= 0 && vehicle.Destination.Name != "" {
			lane = fmt.Sprintf("%d", vehicle.Destination.Lane)
		}

		name := vehicle.CurrentDestination.Name
		if vehicle.CurrentDestination.Type == "Pre-Ingress" {
			name = vehicle.CurrentDestination.Type
		}
		if vehicle.Destination.Type == "YARD" {
			name = vehicle.Destination.Name
		}

		if vehicle.Destination.Type == "QC" {
			hatchData, _ := redisClient.HGet(ctx, HatchCoverOps, vehicle.Destination.Name).Result()
			if hatchData != "" {
				vehicle.HatchCover = "ON"
			}
		}

		task, _ := redisClient.HGet(ctx, GetTaskInfo, vehicle.ID).Result()
		_ = json.Unmarshal([]byte(task), &vehicle.TaskInfo)

		st := vehicle.Destination.CreateTime
		if len(st) >= 12 {
			st = st[11:]
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
			index + 1, vehicle.ID, vehicle.Destination.Type, vehicle.TaskInfo.ContainerSize,
			vehicle.TaskInfo.DestLocation, st, vehicle.Destination.Name,
			lane, name, dtype, arrived, called, vehicle.Mode, ready, manual,
		}
		vehicleTable.AppendRow(row)

		vehicleTable.SetRowPainter(func(row table.Row) text.Colors {
			if row[12].(string) == "MA" {
				return text.Colors{text.FgRed}
			} else if row[12].(string) == "TN" {
				return text.Colors{text.FgYellow}
			}
			return nil
		})

	}
	fmt.Print(vehicleTable.Render())
}

func init() {
	VehicleCmd.Flags().BoolVarP(&k, "keepalive", "k", false, "是否保持刷新(F/5s)")
	VehicleCmd.Flags().StringVarP(&vehicleID, "vehicle", "v", "", "集卡号")
	VehicleCmd.Flags().StringVarP(&vehicleFilter, "filter", "f", "", "指定的作业类型"+taskTypes)
	VehicleCmd.Flags().BoolVar(&vehicleReset, "reset", false, "重置集卡")
}
