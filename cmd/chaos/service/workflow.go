package service

import (
	"encoding/json"
	"fms-awesome-tools/cmd/chaos/internal"
	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fms-awesome-tools/constants"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"sync"
)

var wf *workflow

type workflow struct {
	UUID   string
	client *MqttClient
	wg     sync.WaitGroup
}

func StartWorkflow() error {

	topics := map[string]byte{}
	for _, v := range constants.TopicFromFMS {
		topics[v] = 1
	}

	wf.client.SubscribeMultiple(topics, wf.messageHandler)
	return nil
}

func (w *workflow) messageHandler(client mqtt.Client, message mqtt.Message) {
	if message.Topic() != "heartbeat" {
		fmt.Printf("接收到 <%s> 数据 ==> %s\n", message.Topic(), string(message.Payload()))
	}
	switch message.Topic() {
	case "heartbeat":
		return
	case "power_on_request":
		fmt.Println("power_on_request")
	case "update_trailer":
		data := internal.ParseToLogonRequest(message.Payload())
		resp := internal.GenerateToLogonResponse(data)
		if err := w.client.Publish("logon_response", resp.String()); err != nil {
			fmt.Println("publish error, ", err)
		} else {
			fmt.Printf("发送到 <%s> 数据 ==> %s\n", message.Topic(), resp.String())
		}
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
		data := internal.ParseToRouteResponseJobInstruction(message.Payload())
		if len(data.Data.RouteDAG) != 0 {
			job := &messages.JobInstruction{}
			if err := json.Unmarshal(message.Payload(), job); err != nil {
				fmt.Println("job unmarshal error, ", err)
			}
			if err := w.client.Publish("job_instruction", job.String()); err != nil {
				fmt.Println("publish job error, ", err)
			}
		} else {
			fmt.Println("route response no route dag, ignore")
		}
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

func Close() {
	wf.client.Close()
}

func (w *workflow) connect() error {
	var err error
	w.client, err = NewMQTTClientWithConfig("workflow")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	wf = &workflow{
		UUID: uuid.NewString(),
		wg:   sync.WaitGroup{},
	}

	if err := wf.connect(); err != nil {
		panic(err)
	}
}
