package service

import (
	"encoding/json"
	tools "fms-awesome-tools/utils"
	"fmt"
	"os"
	"strings"
	"time"

	"fms-awesome-tools/cmd/chaos/internal"
	"fms-awesome-tools/cmd/chaos/internal/messages"
)

func (wf *Workflow) response(topic, message string) {
	if err := wf.client.Publish(topic, message); err != nil {
		fmt.Printf("[%s] send message to <%s> failed: %s\n", time.Now().Local().String(), topic, err.Error())
	} else {
		fmt.Printf("[%s] send message to <%s>: %s\n\n", time.Now().Local().String(), topic, message)
	}
}

func (wf *Workflow) callInRequest() {
	call := &messages.CallInRequest{
		APMID: wf.task.APMID,
		Data: messages.CallInRequestData{
			CallInMode: 0, Crane: wf.task.Data.NextLocation,
		},
	}
	wf.response("call_in_request", call.String())
}

func (wf *Workflow) mount() {
	mount := messages.MountInstruction{
		APMID: wf.task.APMID,
		Data:  messages.MountInstructionData{},
	}
	wf.response("mount_instruction", mount.String())
}

func (wf *Workflow) offload() {
	offload := messages.OffloadInstruction{
		APMID: wf.task.APMID,
		Data: messages.OffloadInstructionData{
			CntrNumber: "FFFF 0000000",
		},
	}
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

func (wf *Workflow) routeHandler(message []byte) {

}

func (wf *Workflow) routeJobResponseHandler(message []byte) {
	data := internal.ParseToRouteResponseJobInstruction(message)
	if len(data.Data.RouteDAG) != 0 {
		job := &messages.JobInstruction{}
		if err := json.Unmarshal(message, job); err != nil {
			fmt.Println("job_instruction解析失败...")
			os.Exit(1)
		}
		job.Data.RouteMandate = "Y"
		wf.response("job_instruction", job.String())
		wf.task = &data
		return
	}
	fmt.Println("job_instruction == > route dag为空, 任务下发失败, 流程结束")
	os.Exit(1)
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

	if strings.TrimSpace(data.Data.Location) == strings.TrimSpace(wf.destination) {
		switch wf.task.Data.Activity {
		case 2, 3, 4:
			wf.mount()
		case 6, 7, 8:
			wf.offload()
		case 1, 5:
			time.Sleep(time.Second)
		default:
			return
		}

		if wf.loop < 0 {
			fmt.Println(tools.CustomTitle("\n          当前流程已结束, 待执行下一个流程...          \n"))
		} else if wf.loop == 0 {
			fmt.Println(tools.CustomTitle("\n          当前流程已结束...          \n"))
			os.Exit(1)
		} else if wf.loopCount > wf.loop {
			fmt.Println(tools.CustomTitle("\n          流程已全部执行结束...          \n"))
			os.Exit(1)
		} else {
			fmt.Println(tools.CustomTitle("\n          当前流程已结束, 待执行下一个流程...          \n"))
			wf.loopCount++
		}
		wf.sendNewTask()
	}
}

func (wf *Workflow) pathUpdateRequestHandler(message []byte) {

}

func (wf *Workflow) dockPositionResponseHandler(message []byte) {

}

func (wf *Workflow) sendNewTask() {
	if strings.Contains(wf.destination, "PQC") {
		wf.updateBlockTask()
	} else {
		wf.updateQCTask()
	}
	time.Sleep(time.Second * 3)
	message := messages.GenerateRouteRequestJob(wf.destination, wf.lane, "S", "5", wf.activity, 1, 40, 1)
	if err := PublishAssignedTopic("route_request_job_instruction", "", message); err != nil {
		fmt.Printf("[%s] 任务下发失败: %s, 程序退出...", time.Now().Local().String(), err)
		os.Exit(1)
	} else {
		fmt.Printf("[%s] send message to <%s>: %s\n", time.Now().Local().String(), "route_request_job_instruction", message)
	}
}
