package service

import (
	"fmt"
	"os"
	"sync"

	"fms-awesome-tools/cmd/chaos/internal"
	"fms-awesome-tools/constants"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type Workflow struct {
	UUID   string
	client *MqttClient
	wg     sync.WaitGroup
}

func NewWorkflow() *Workflow {
	w := &Workflow{
		UUID: uuid.NewString(),
		wg:   sync.WaitGroup{},
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

	wf.client.SubscribeMultiple(topics, wf.messageHandler)
	return nil
}

func (wf *Workflow) messageHandler(client mqtt.Client, message mqtt.Message) {
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
	case "logoff_request":
		fmt.Println("logoff_request")
	case "power_off_request":
		fmt.Println("power_off_request")
	case "switch_mode_response":
		resp := internal.ParseToSwitchModeResponse(message.Payload())
		fmt.Println("switch_mode_response: ", resp.Data)
	case "mode_change_update":
		fmt.Println("mode_change_update")
	case "route_response_job_instruction":
		wf.routeJobResponseHandler(message.Payload())
	case "apm_arrived_request":
		fmt.Println("apm_arrived_request")
	case "dock_position_response":
		fmt.Println("dock_position_response")
	case "mount_instruction_response":
		fmt.Println("mount_instruction_response")
	case "offload_instruction_response":
		fmt.Println("offload_instruction_response")
	case "wharf_dock_position":
		fmt.Println("wharf_dock_position")
	case "ingress_ready_response":
		fmt.Println("ingress_ready_response")
	case "ready_for_move_to_qc":
		resp := internal.ParseToReadyForMoveToQC(message.Payload())
		fmt.Println("ready_for_move_to_qc: ", resp)
	case "ready_for_ingress_to_call_in":
		resp := internal.ParseToReadyForIngressToCallIn(message.Payload())
		fmt.Println("ready_for_ingress_to_call_in: ", resp)
	default:
		fmt.Println("get topic: ", message.Topic(), ", but not implement, ignore")
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
