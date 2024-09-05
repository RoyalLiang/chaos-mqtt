package service

import (
	"context"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
	"time"
)

var topics = []string{
	"power_on_request", "heartbeat", "update_trailer", "logoff_request", "power_off_request", "request_job",
	"cancel_job_response", "switch_mode_response", "mode_change_update", "block_response", "blocks_response",
	"parking_state", "maintenance_response", "parking_response", "resume_response", "stop_response",
	"route_response_job_instruction", "route_response", "job_instruction_response", "apm_move_request",
	"apm_arrived_request", "dock_position_response", "mount_instruction_response", "offload_instruction_response",
	"wharf_dock_position", "intermediate_location", "ingress_ready_response", "ingress_to_call_in_response",
	"move_to_qc_response", "armg_request_response", "vessel_berth_response", "vessel_unberth_response",
	"derived_vessel_configuration", "call_in_status_response", "hatch_cover_ops_response", "qc_position_info",
	"path_update_available", "path_update_response", "ready_for_ingress_to_call_in", "ready_for_move_to_qc",
	"ready_for_ingress_to_qc", "ingress_to_qc_response", "call_in_response", "coning_deconing_completion_response",
	"pm_activity_info_response", "pm_navigation_info_response", "manual_exception_handling_response",
	"apm_arrived_response", "apm_acceptance_update_response",
}

type Workflow struct {
	UUID     string
	Truck    string
	Activity int64
	ctx      context.Context
	wg       sync.WaitGroup
	exit     chan struct{}
}

func (w *Workflow) StartWorkflow() error {
	w.wg = sync.WaitGroup{}
	for _, v := range topics {
		w.wg.Add(1)
		go MQTTClient.Subscribe(&w.wg, v, byte(0), w.MessageHandler)
		w.wg.Wait()
	}
	time.Sleep(1 * time.Second)

	select {
	case <-time.After(time.Second * 30):
		fmt.Println("running...")
	case <-w.exit:
		return nil
	default:
		time.Sleep(time.Second)
	}
	return nil
}

func (w *Workflow) MessageHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Println(message.Topic())
	fmt.Println(string(message.Payload()))
	for {
		switch message.Topic() {
		case "power_on_request":
			fmt.Println("power_on_request")
		case "heartbeat":
			fmt.Println("heartbeat")
		case "update_trailer":
			fmt.Println("update_trailer")
		case "logoff_request":
			fmt.Println("logoff_request")
		case "power_off_request":
			fmt.Println("power_off_request")
		case "request_job":
			fmt.Println("request_job")
		case "cancel_job_response":
			fmt.Println("cancel_job_response")
		case "switch_mode_response":
			fmt.Println("switch_mode_response")
		case "mode_change_update":
			fmt.Println("mode_change_update")
		case "block_response":
			fmt.Println("block_response")
		case "parking_state":
			fmt.Println("parking_state")
		case "maintenance_response":
			fmt.Println("maintenance_response")
		case "route_response_job_instruction":
			fmt.Println("route_response_job_instruction")
		case "route_response":
			fmt.Println("route_response")
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
		case "intermediate_location":
			fmt.Println("intermediate_location")
		case "ingress_ready_response":
			fmt.Println("ingress_ready_response")
		case "ingress_to_call_in_response":
			fmt.Println("ingress_to_call_in_response")
		}
	}

}
