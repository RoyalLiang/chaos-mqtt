package service

import (
	"fmt"
	"os"
	"sync"
	"time"

	tools "fms-awesome-tools/utils"

	"fms-awesome-tools/cmd/chaos/internal/messages"

	"fms-awesome-tools/cmd/chaos/internal"
	"fms-awesome-tools/constants"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type Workflow struct {
	UUID   string
	client *MqttClient
	wg     sync.WaitGroup

	task        *messages.RouteResponseJobInstruction
	vehicleID   string
	destination string
	lane        string
	activity    int64
	taskType    string
	autoCallIn  bool
}

func NewWorkflow(activity int64, lane, vehicleID, dest string, autoCallIn bool) *Workflow {
	w := &Workflow{
		UUID:        uuid.NewString(),
		wg:          sync.WaitGroup{},
		autoCallIn:  autoCallIn,
		activity:    activity,
		lane:        lane,
		vehicleID:   vehicleID,
		destination: dest,
	}
	if err := w.connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return w
}

func (wf *Workflow) StartWorkflow() error {

	topics := map[string]byte{}
	for _, v := range constants.TopicFromFMS {
		topics[v] = 1
	}

	go func() {
		time.Sleep(time.Second * 3)
		message := messages.GenerateRouteRequestJob(wf.destination, wf.lane, "5", 1, 40, 1)
		if err := PublishAssignedTopic("route_request_job_instruction", "", message); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("route_request_job_instruction ==> ", message)
		}
	}()

	fmt.Println(tools.CustomTitle("\n          Chaos Workflow Start Listen...          \n"))
	wf.client.SubscribeMultiple(topics, wf.messageHandler)
	return nil
}

func (wf *Workflow) messageHandler(client mqtt.Client, message mqtt.Message) {
	//msg := internal.ParseToTask(message.Payload())
	//if wf.task != nil && msg.APMID != "" && wf.task.APMID != msg.APMID && message.Topic() != "heartbeat" && message.Topic() != "route_response_job_instruction" {
	//	fmt.Printf("当前workflow集卡: %s, 获取到不匹配的集卡号: %s,忽略\n", wf.task.APMID, msg.APMID)
	//	return
	//}

	if message.Topic() != "heartbeat" {
		fmt.Printf("%s <== %s\n", message.Topic(), string(message.Payload()))
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
