package service

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
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
	Blocks    = []string{"TB", "TC", "TD", "TE", "TF", "TG", "TH"}
	BlockNums = []string{"01", "02", "03", "04", "05", "06", "07"}
	Slots     = []int64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	Lanes     = []string{"0", "11"}
	QCLanes   = []string{"2", "3", "5", "6"}
)

type Workflow struct {
	client      *MqttClient
	wg          sync.WaitGroup
	task        *messages.RouteResponseJobInstruction
	UUID        string
	vehicleID   string
	destination string
	lane        string
	activity    int64
	taskType    string
	autoCallIn  bool
	loop        int64
	loopCount   int64
}

func NewWorkflow(loopNum, activity int64, lane, vehicleID, dest string, autoCallIn bool) *Workflow {
	w := &Workflow{
		UUID:        uuid.NewString(),
		wg:          sync.WaitGroup{},
		autoCallIn:  autoCallIn,
		activity:    activity,
		lane:        lane,
		vehicleID:   vehicleID,
		destination: dest,
		loop:        loopNum,
		loopCount:   1,
	}

	if strings.HasPrefix(w.destination, "PQC") {
		w.destination = fmt.Sprintf("P,%s          ", w.destination)
		w.taskType = "QC"
	} else if w.activity == 1 || w.activity == 5 {
		w.taskType = "STANDBY"
	} else {
		w.taskType = "IYS"
	}

	if err := w.connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return w
}

func sendLogon() {
	message := messages.LogonResponse{
		APMID: constants.VehicleID,
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
	bni := rand.IntN(len(BlockNums))
	si := rand.IntN(len(Slots))
	return fmt.Sprintf("Y,V,,%s%s,%02d,%02d,10,  ", Blocks[bi], BlockNums[bni], Slots[si], Slots[si]+1)
}

func choiceBlockLane() string {
	li := rand.IntN(len(Lanes))
	return Lanes[li]
}

func (wf *Workflow) updateTask() {
	if strings.Contains(wf.destination, "PQC") {
		wf.updateBlockTask()
	} else {
		wf.updateQCTask()
	}
}

func (wf *Workflow) updateBlockTask() {
	wf.taskType = choiceTaskType()
	if wf.taskType == "STANDBY" {
		wf.activity = 1
		wf.destination = ""
		wf.lane = ""
	} else if wf.taskType == "IYS" {
		wf.activity = 2
		wf.lane = choiceBlockLane()
		wf.destination = choiceBlockLocation()
	}
}

func (wf *Workflow) updateQCTask() {
	wf.destination = "P,PQC924          "
	wf.activity = 2
	wf.taskType = "QC"
	wf.lane = choiceQCLane()
}

func (wf *Workflow) StartWorkflow() error {
	topics := map[string]byte{}
	for _, v := range constants.TopicFromFMS {
		topics[v] = 1
	}

	sendLogon()
	go func() {
		time.Sleep(time.Second * 3)
		message := messages.GenerateRouteRequestJob(wf.destination, wf.lane, "S", "5", 1, 40, 1)
		if err := PublishAssignedTopic("route_request_job_instruction", "", message); err != nil {
			fmt.Printf("[%s] 任务下发失败: %s, 程序退出...", time.Now().Local().String(), err)
			os.Exit(1)
		} else {
			fmt.Printf("[%s] send message to <%s>: %s\n\n", time.Now().Local().String(), "route_request_job_instruction", message)
		}
	}()

	fmt.Println(tools.CustomTitle("\n          Chaos Workflow Start...          \n"))
	wf.client.SubscribeMultiple(topics, wf.messageHandler)
	return nil
}

func (wf *Workflow) messageHandler(client mqtt.Client, message mqtt.Message) {
	data := &Message{}
	err := json.Unmarshal(message.Payload(), data)
	if err != nil || data.APMID != wf.vehicleID {
		return
	}

	if message.Topic() != "heartbeat" {
		fmt.Printf("[%s] receive message from <%s>: %s\n\n", time.Now().Local().String(), message.Topic(), string(message.Payload()))
	}

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
		wf.routeJobResponseHandler(message.Payload())
	case "apm_arrived_request":
		wf.apmArrivalHandler(message.Payload())
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
		wf.readyForMoveToQCheHandler(message.Payload())
	case "ready_for_ingress_to_call_in":
		wf.readyForIngressToCallInHandler(message.Payload())
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

func (wf *Workflow) Close() {
	wf.client.Close()
}
