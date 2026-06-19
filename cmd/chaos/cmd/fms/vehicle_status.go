package area

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/configs"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

const (
	moveCursor    = "\033[s" // 保存光标位置
	restoreCursor = "\033[u" // 恢复光标位置
	clearScreen   = "\033[J" // 清除从光标到屏幕底部的内容
)

const (
	taskTypes     = "\nQC\nYARD (STANDBY)\nIYS\n"
	vehicleModes  = "\nMA\nTN\nOP\n"
	HatchCoverOps = "psa_hatch_cover_ops"
	GetTaskInfo   = "psa_task_info"
)

var (
	vehicleID     string
	k             bool
	vehicleReset  bool
	vehicleTable  = table.NewWriter()
	redisClient   *redis.Client
	vehicleFilter []string
	exclude       string
	withQtruck    bool
)

var VehicleCmd = &cobra.Command{
	Use:   "vehicles",
	Short: "获取所有/指定集卡状态",
	Run: func(cmd *cobra.Command, args []string) {
		header := table.Row{
			"ID", "Vehicle ID", "Task Type", "Job Type", "Lift", "Cones", "ISO", "Start Time", "Destination",
			"Lane", "Curr Destination", "WTA/s", "QCHT/s", "Arrived", "Call Status", "Mode", "Ready", "Manual",
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

		ensureRedisClient()

		if exclude != "" && (exclude != "MA" && exclude != "TN" && exclude != "OP") {
			cobra.CheckErr("--exclude 选项错误")
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
	for _, vehicle := range vm.vehicles {
		if vehicleID != "" && vehicle.ID != vehicleID {
			continue
		}

		if !withQtruck && strings.HasPrefix(vehicle.ID, "AT") {
			continue
		}

		if exclude != "" && vehicle.Mode == exclude {
			continue
		}

		if len(vehicleFilter) > 0 {
			matched := false
			for _, filter := range vehicleFilter {
				if vehicle.Destination.Type == filter {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
		}

		vehicles = append(vehicles, *vehicle)
	}
	sort.Sort(vehicles)
	return vehicles
}

func ensureRedisClient() {
	if redisClient != nil {
		return
	}

	var err error
	redisClient, err = service.NewRedisClient()
	if err != nil {
		cobra.CheckErr(err)
	}
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
	ensureRedisClient()

	var (
		ctx          = context.Background()
		msgChan      = make(chan *redis.Message, 999)
		batchTimeout = time.Second * 3
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

		dname := vehicle.CurrentDestination.Name
		if vehicle.CurrentDestination.Type == "Pre-Ingress" {
			dname = vehicle.CurrentDestination.Type
		}
		if vehicle.Destination.Type == "YARD" {
			dname = vehicle.Destination.Name
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

		cons := ""
		if len(vehicle.TaskInfo.Containers) > 0 {
			cons = strconv.FormatInt(int64(len(vehicle.TaskInfo.Containers)), 10)
		}

		job := ""
		switch vehicle.TaskInfo.Activity {
		case 2, 3, 4:
			job = "MOUNT"
		case 6, 7, 8:
			job = "OFFLOAD"
		case 1, 5:
			job = "STANDBY"
		}

		switch vehicle.CurrentDestination.Type {
		case "CRANE_AREA":
			dname += " (QC)"
		case "CALLIN_AREA":
			dname += " (CA)"
		case "WAIT_AREA":
			dname += " (DWA)"
		}

		waitTime := ""
		qcTime := ""
		var waitColor text.Colors
		if vehicle.KaTime == 0 || vehicle.CallTime == 0 {
			waitTime = ""
		} else if vehicle.KaTime > 0 && vehicle.CallTime == 0 {
			dt := (time.Now().UnixMilli() - int64(vehicle.KaTime)) / 1000
			if 120 <= dt && dt < 300 {
				waitColor = text.Colors{text.BgGreen}
			} else if dt > 300 && dt < 600 {
				waitColor = text.Colors{text.FgYellow}
			} else {
				waitColor = text.Colors{text.FgRed}
			}
			waitTime = strconv.FormatInt(dt, 10)
		} else {
			waitTime = strconv.FormatInt(int64(vehicle.CallTime-vehicle.KaTime)/1000, 10)
			waitColor = text.Colors{}
		}

		if vehicle.Arrived && vehicle.CurrentDestination.Type == "CRANE_AREA" {
			dt := time.Now().UnixMilli()

			layout := "2006-01-02 15:04:05"
			loc, _ := time.LoadLocation("Asia/Shanghai")
			t, err := time.ParseInLocation(layout, vehicle.AT, loc)
			if err == nil {
				qcTime = strconv.FormatInt((dt-t.UnixMilli())/1000, 10)
			} else {
				fmt.Println("time err: ", err.Error())
			}
		}

		row := table.Row{
			index + 1, vehicle.ID, vehicle.Destination.Type, job, vehicle.TaskInfo.LiftType, cons, vehicle.TaskInfo.ContainerSize,
			st, vehicle.Destination.Name, lane, dname, waitTime, qcTime, arrived, called, vehicle.Mode, ready, manual,
		}

		vehicleTable.AppendRow(row)
		vehicleTable.SetRowPainter(func(row table.Row) text.Colors {
			if row[14].(string) == "MA" {
				return text.Colors{text.FgRed}
			} else if row[14].(string) == "TN" {
				return text.Colors{text.FgYellow}
			} else if waitColor != nil && waitTime != "" {
				return waitColor
			}

			return nil
		})

	}
	fmt.Print(vehicleTable.Render())
}

func init() {
	VehicleCmd.Flags().BoolVarP(&k, "keepalive", "k", false, "自动刷新🔄️(1/2s)")
	VehicleCmd.Flags().StringVarP(&vehicleID, "vehicle", "v", "", "集卡号🚗")
	VehicleCmd.Flags().StringSliceVarP(&vehicleFilter, "filter", "f", make([]string, 0), "指定的作业类型"+taskTypes)
	VehicleCmd.Flags().StringVarP(&exclude, "exclude", "e", "", "过滤指定模式的集卡"+vehicleModes)
	VehicleCmd.Flags().BoolVar(&vehicleReset, "reset", false, "重置集卡⭕")
	VehicleCmd.Flags().BoolVar(&withQtruck, "with-qtruck", false, "同时显示Qtruck🚗")
}
