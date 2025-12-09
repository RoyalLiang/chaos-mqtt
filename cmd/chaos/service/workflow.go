package service

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	tools "fms-awesome-tools/utils"

	"fms-awesome-tools/cmd/chaos/internal/messages"

	"fms-awesome-tools/cmd/chaos/internal"
	"fms-awesome-tools/constants"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

var (
	Tasks     = []string{"IYS", "STANDBY"}
	Blocks    = []string{"TB", "TC", "TD", "TE", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TP", "TQ", "TR", "TU", "TS", "TT"}
	BlockNums = []string{"01", "02", "03", "04"}
	Slots     = []int64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}
	Lanes     = []string{"0", "11"}
	QCLanes   = []string{"2", "3", "5", "6"}
)

type vehicleTask struct {
	activity     int64
	vehicleID    string
	destination  string
	lane         string
	assignedQC   string
	assignedLane string
	taskType     string
	noStandby    bool
	onlyStandby  bool
	task         *messages.RouteResponseJobInstruction
}

type Workflow struct {
	client     *MqttClient
	wg         sync.WaitGroup
	UUID       string
	autoCallIn bool
	loop       int64
	loopCount  int64
	vehicles   map[string]*vehicleTask
}

func (vt *vehicleTask) updateDest(activity int64, dest string) {
	if strings.HasPrefix(dest, "PQC") {
		vt.destination = fmt.Sprintf("P,%s          ", dest)
		vt.taskType = "QC"
	} else if activity == 1 || activity == 5 {
		vt.taskType = "STANDBY"
	} else {
		vt.taskType = "IYS"
	}
}

func NewWorkflow(loopNum, activity, vehicleNum, startNumber int64, lane, vehicleID, dest, aQC, aLane string, autoCallIn, noStandby, onlyStandby bool) *Workflow {
	w := &Workflow{
		UUID:       uuid.NewString(),
		wg:         sync.WaitGroup{},
		autoCallIn: autoCallIn,
		loop:       loopNum,
		loopCount:  1,
		vehicles: func() map[string]*vehicleTask {
			vs := make(map[string]*vehicleTask)
			if vehicleID != "" && vehicleNum <= 0 {
				vehicleNum = 1
			}

			var vid string
			for i := 0; i < int(vehicleNum); i++ {
				if vehicleID != "" {
					vid = vehicleID
				} else {
					vid = fmt.Sprintf("APM%d", 9000+i+int(startNumber))
				}
				vs[vid] = &vehicleTask{
					activity:     activity,
					vehicleID:    vid,
					destination:  dest,
					lane:         lane,
					assignedQC:   aQC,
					assignedLane: aLane,
					noStandby:    noStandby,
					onlyStandby:  onlyStandby,
				}
				vs[vid].updateDest(activity, dest)
				fmt.Printf("%s, ", vid)
			}
			fmt.Println()
			return vs
		}(),
	}

	if err := w.connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return w
}

func sendLogon(vehicleID string) {
	message := messages.LogonResponse{
		APMID: vehicleID,
		Data: messages.LogonResponseData{
			Success: 1, TrailerSeqNumbers: []int{1}, TrailerLengths: []int{20}, TrailerUnladenWeights: []int{11},
			TrailerTypes: strings.Split("CST", ","), TrailerPayloads: []int{200}, TrailerWidths: make([]int, 0),
			TrailerHeights: make([]int, 0), TrailerNumbers: strings.Split("C53525", ","),
		},
	}.String()
	_ = PublishAssignedTopic("logon_response", "", message)
}

func choiceTaskType() string {
	index := rand.IntN(len(Tasks))
	return Tasks[index]
}

func choiceQCLane() string {
	index := rand.IntN(len(QCLanes))
	return QCLanes[index]
}

func choiceBlockLocation() string {
	bi := rand.IntN(len(Blocks))
	slots := BlockNums
	switch Blocks[bi] {
	case "TB", "TD", "TG", "TJ":
		slots = []string{"01", "02", "03", "04", "05", "06", "07"}
	case "TS", "TR", "TT":
		slots = []string{"01", "02", "03"}
	case "TN", "TP":
		slots = []string{"02", "03", "04"}
	}
	bni := rand.IntN(len(slots))
	si := rand.IntN(len(Slots))
	return fmt.Sprintf("Y,V,,%s%s,%02d,%02d,10,  ", Blocks[bi], slots[bni], Slots[si], Slots[si]+1)
}

func choiceBlockLane(dest string) string {
	if strings.Contains(dest, "TU01") || strings.Contains(dest, "TN01") || strings.Contains(dest, "TP01") {
		return "11"
	}

	li := rand.IntN(len(Lanes))
	return Lanes[li]
}

func binarySearch(arr []string, target string) bool {
	sort.Strings(arr)
	index := sort.SearchStrings(arr, target)
	return index < len(arr) && arr[index] == target
}

func getAgainstActivity(activity int64) int64 {
	switch activity {
	case 2, 3, 4:
		return 6
	case 6, 7, 8:
		return 2
	default:
		return 1
	}
}

func (vt *vehicleTask) updateBlockTask() {
	vt.taskType = choiceTaskType()
	if vt.noStandby {
		vt.taskType = "IYS"
		vt.activity = getAgainstActivity(vt.activity)
		vt.destination = choiceBlockLocation()
		vt.lane = choiceBlockLane(vt.destination)
		return
	}

	if vt.taskType == "STANDBY" {
		vt.activity = 1
		vt.destination = ""
		vt.lane = ""
	}
}

func (vt *vehicleTask) updateQCTask() {
	if vt.assignedQC != "" {
		vt.destination = fmt.Sprintf("P,%s          ", vt.assignedQC)
	} else {
		vt.destination = "P,PQC924          "
	}
	vt.taskType = "QC"
	vt.activity = getAgainstActivity(vt.activity)
	if vt.assignedLane != "" {
		vt.lane = vt.assignedLane
	} else {
		vt.lane = choiceQCLane()
	}
}

func (wf *Workflow) StartWorkflow() error {
	topics := map[string]byte{}
	for _, v := range constants.TopicFromFMS {
		topics[v] = 1
	}

	go func() {
		time.Sleep(time.Second * 2)
		for _, vt := range wf.vehicles {
			sendLogon(vt.vehicleID)
			time.Sleep(time.Millisecond * 500)
			wf.sendJob(vt)
			time.Sleep(time.Second)
		}
	}()

	fmt.Println(tools.CustomTitle("\n          Chaos Workflow Start...          \n"))
	wf.client.SubscribeMultiple(topics, wf.messageHandler)
	return nil
}

func (wf *Workflow) messageHandler(client mqtt.Client, message mqtt.Message) {
	data := &Message{}
	err := json.Unmarshal(message.Payload(), data)
	_, ok := wf.vehicles[data.APMID]
	if err != nil || !ok {
		return
	}

	if message.Topic() != "heartbeat" {
		fmt.Printf("[%s] receive message from <%s>: %s\n\n", time.Now().Local().String(), message.Topic(), string(message.Payload()))
	}

	vehicle := wf.getVehicle(data.APMID)
	switch message.Topic() {
	case "heartbeat":
		return
	case "power_on_request":
		fmt.Println("power_on_request")
	case "update_trailer":
		wf.logonHandler(message.Payload())
	case "route_response":
		wf.routeHandler(message.Payload())
	//case "logoff_request":
	//	fmt.Println("logoff_request")
	//case "power_off_request":
	//	fmt.Println("power_off_request")
	case "switch_mode_response":
		_ = internal.ParseToSwitchModeResponse(message.Payload())
		//fmt.Println("switch_mode_response: ", resp.Data)
	//case "mode_change_update":
	//	fmt.Println("mode_change_update")
	case "route_response_job_instruction":
		wf.routeJobResponseHandler(vehicle, message.Payload())
	case "apm_arrived_request":
		wf.apmArrivalHandler(vehicle, message.Payload())
	//case "dock_position_response":
	//	fmt.Println("dock_position_response")
	//case "mount_instruction_response":
	//	fmt.Println("mount_instruction_response")
	//case "offload_instruction_response":
	//	fmt.Println("offload_instruction_response")
	//case "wharf_dock_position":
	//	fmt.Println("wharf_dock_position")
	//case "ingress_ready_response":
	//	fmt.Println("ingress_ready_response")
	case "ready_for_move_to_qc":
		wf.readyForMoveToQCheHandler(vehicle, message.Payload())
	case "ready_for_ingress_to_call_in":
		wf.readyForIngressToCallInHandler(vehicle, message.Payload())
		//default:
		//	fmt.Println("get topic: ", message.Topic(), ", but not implement, ignore")
	}

}

func (wf *Workflow) connect() error {
	var err error
	wf.client, err = NewMQTTClientWithConfig("workflow")
	if err != nil {
		return err
	}
	return nil
}

func (wf *Workflow) getVehicle(vehicleID string) *vehicleTask {
	var vehicle *vehicleTask
	for _, v := range wf.vehicles {
		if v.vehicleID == vehicleID {
			vehicle = v
		}
	}
	return vehicle
}

func (wf *Workflow) Close() {
	wf.client.Close()
}
