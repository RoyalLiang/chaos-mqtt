package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"fms-awesome-tools/cmd/chaos/internal"
	"fms-awesome-tools/cmd/chaos/internal/messages"
)

func (wf *Workflow) response(topic, message string) {
	if err := wf.client.Publish(topic, message); err != nil {
		fmt.Printf("%s ==> %s\n", topic, err)
	} else {
		fmt.Printf("%s ==> %s\n", topic, message)
	}
}

func (wf *Workflow) callInRequest() {
	call := &messages.CallInRequest{
		APMID: wf.task.APMID,
		Data: messages.CallInRequestData{
			CallInMode: 0, Crane: wf.task.Data.NextLocation,
		},
	}
	time.Sleep(5 * time.Second)
	wf.response("call_in_request", call.String())
}

func (wf *Workflow) mount() {
	mount := messages.MountInstruction{
		APMID: wf.task.APMID,
		Data:  messages.MountInstructionData{},
	}
	time.Sleep(2 * time.Second)
	wf.response("mount_instruction", mount.String())
}

func (wf *Workflow) offload() {
	offload := messages.OffloadInstruction{
		APMID: wf.task.APMID,
		Data: messages.OffloadInstructionData{
			CntrNumber: "FFFF0000000",
		},
	}
	time.Sleep(2 * time.Second)
	wf.response("offload_instruction", offload.String())
}

func (wf *Workflow) logonHandler(message []byte) {
	data := internal.ParseToLogonRequest(message)
	resp := messages.LogonResponse{
		APMID: data.APMID,
		Data: messages.LogonResponseData{
			Success: 1, TrailerSeqNumbers: data.Data.TrailerSeqNumbers, TrailerNumbers: data.Data.TrailerNumbers,
			TrailerHeights: make([]int, 0), TrailerLengths: make([]int, 0), TrailerPayloads: make([]int, 0),
			TrailerTypes: make([]string, 0), TrailerUnladenWeights: make([]int, 0), TrailerWidths: make([]int, 0),
			NumTrailers: data.Data.NumTrailers,
		},
	}
	wf.response("logon_response", resp.String())
}

func (wf *Workflow) routeJobResponseHandler(message []byte) {
	data := internal.ParseToRouteResponseJobInstruction(message)
	if len(data.Data.RouteDAG) != 0 {
		job := &messages.JobInstruction{}
		if err := json.Unmarshal(message, job); err != nil {
			fmt.Println("job_instruction == > route_response_job_instruction解析失败")
			return
		}
		wf.response("job_instruction", job.String())
		wf.task = &data
		return
	}
	fmt.Println("job_instruction == > route dag为空, 任务下发失败")
}

func (wf *Workflow) readyForIngressToCallInHandler(message []byte) {
	ready := internal.ParseToReadyForIngressToCallIn(message)
	d := messages.IngressToCallIn{
		APMID: wf.task.APMID,
		Data: messages.IngressToCallInData{
			RouteDag: ready.Data.RouteDAG, RouteType: ready.Data.RouteType, LaneAvailability: ready.Data.LaneAvailability,
			DestinationName: ready.Data.DestinationName, DestinationLane: ready.Data.DestinationLane,
		},
	}
	wf.response("ingress_to_call_in", d.String())
}

func (wf *Workflow) readyForMoveToQCheHandler(message []byte) {
	data := internal.ParseToReadyForMoveToQC(message)
	move := &messages.MoveToQCRequest{
		APMID: wf.task.APMID,
		Data: messages.MOveToQCRequestData{
			RouteDag: data.Data.RouteDAG,
		},
	}
	wf.response("move_to_qc", move.String())
}

func (wf *Workflow) apmArrivalHandler(message []byte) {
	data := internal.ParseToApmArrivedRequest(message)
	if wf.task == nil {
		return
	}

	if strings.HasSuffix(data.Data.Location, "Pre-Ingress") && wf.autoCallIn {
		wf.callInRequest()
		return
	}

	if strings.HasSuffix(wf.task.Data.NextLocation, data.Data.Location) {
		switch wf.task.Data.Activity {
		case 2:
			wf.mount()
		case 6:
			wf.offload()
		default:
			return
		}
	}
}

func (wf *Workflow) pathUpdateRequestHandler(message []byte) {

}

func (wf *Workflow) dockPositionResponseHandler(message []byte) {

}
