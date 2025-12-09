package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	tools "fms-awesome-tools/utils"

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

func (wf *Workflow) callInRequest(vehicle *vehicleTask) {
	delay := tools.GetCustomSecond(1, 4)
	go func() {
		time.Sleep(time.Duration(delay) * time.Second)
		call := &messages.CallInRequest{
			APMID: vehicle.vehicleID,
			Data: messages.CallInRequestData{
				CallInMode: 0, Crane: vehicle.task.Data.NextLocation,
			},
		}
		wf.response("call_in_request", call.String())
	}()
}

func (wf *Workflow) mount(vehicle *vehicleTask) {
	mount := messages.MountInstruction{
		APMID: vehicle.vehicleID,
		Data:  messages.MountInstructionData{},
	}
	wf.response("mount_instruction", mount.String())
}

func (wf *Workflow) offload(vehicle *vehicleTask) {
	offload := messages.OffloadInstruction{
		APMID: vehicle.vehicleID,
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

func (wf *Workflow) sendJob(vehicle *vehicleTask) {
	var message string
	if vehicle.onlyStandby {
		message = messages.GenerateRouteRequestJob(vehicle.vehicleID, vehicle.destination, vehicle.lane, "S", "5", 1, 1, 40, 1)
	} else {
		message = messages.GenerateRouteRequestJob(vehicle.vehicleID, vehicle.destination, vehicle.lane, "S", "5", vehicle.activity, 1, 40, 1)
	}
	if err := PublishAssignedTopic("route_request_job_instruction", "", message); err != nil {
		fmt.Printf("[%s] [%s] 任务下发失败 ==> [%s]", time.Now().Local().String(), vehicle.vehicleID, err)
	} else {
		fmt.Printf("[%s] [%s] route_request_job_instruction ==> [%s]\n\n", time.Now().Local().String(), vehicle.vehicleID, message)
	}
}

func (wf *Workflow) routeJobResponseHandler(vehicle *vehicleTask, message []byte) {
	data := internal.ParseToRouteResponseJobInstruction(message)
	if data.Data.Success == 0 {
		fmt.Printf("\x1b[41m[%s]: 任务下发失败, 尝试重发， 失败原因: %s \x1b[0m", data.APMID, data.Data.RejectionCode)
		go wf.sendJob(vehicle)
	} else {
		job := &messages.JobInstruction{}
		if err := json.Unmarshal(message, job); err != nil {
			fmt.Println("job_instruction解析失败...")
			os.Exit(1)
		}
		job.Data.RouteMandate = "Y"
		wf.response("job_instruction", job.String())
		vehicle.task = &data
	}
}

func (wf *Workflow) readyForIngressToCallInHandler(vehicle *vehicleTask, message []byte) {
	ready := internal.ParseToReadyForIngressToCallIn(message)
	d := messages.IngressToCallIn{
		APMID: vehicle.vehicleID,
		Data: messages.IngressToCallInData{
			RouteDag: ready.Data.RouteDAG, RouteType: ready.Data.RouteType, LaneAvailability: ready.Data.LaneAvailability,
			DestinationName: ready.Data.DestinationName, DestinationLane: ready.Data.DestinationLane,
		},
	}
	wf.response("ingress_to_call_in", d.String())
}

func (wf *Workflow) readyForMoveToQCheHandler(vehicle *vehicleTask, message []byte) {
	data := internal.ParseToReadyForMoveToQC(message)
	move := &messages.MoveToQCRequest{
		APMID: vehicle.vehicleID,
		Data: messages.MOveToQCRequestData{
			RouteDag: data.Data.RouteDAG,
		},
	}
	wf.response("move_to_qc", move.String())
}

func (wf *Workflow) apmArrivalHandler(vehicle *vehicleTask, message []byte) {
	data := internal.ParseToApmArrivedRequest(message)
	if vehicle.task == nil {
		return
	}

	if (strings.HasSuffix(data.Data.Location, "Pre-Ingress") || strings.Contains(data.Data.Location, "Wait")) && wf.autoCallIn {
		wf.callInRequest(vehicle)
		return
	}

	switch vehicle.activity {
	case 1, 5:
		wf.sendNewTask(vehicle)
		return
	}

	go func() {
		if strings.TrimSpace(data.Data.Location) == strings.TrimSpace(vehicle.destination) {
			time.Sleep(time.Second * 150)
			switch vehicle.activity {
			case 2, 3, 4:
				wf.mount(vehicle)
			case 6, 7, 8:
				wf.offload(vehicle)
			default:
				return
			}

			time.Sleep(time.Second * 1)
			wf.sendNewTask(vehicle)
		}
	}()
}

func (wf *Workflow) pathUpdateRequestHandler(message []byte) {

}

func (wf *Workflow) dockPositionResponseHandler(message []byte) {

}

func (wf *Workflow) sendNewTask(vt *vehicleTask) {
	if wf.loop < 0 {
		fmt.Println(tools.CustomTitle(fmt.Sprintf("\n          [%s]: 当前流程已结束, 待执行下一个流程...          \n", vt.vehicleID)))
	} else if wf.loop == 0 || wf.loopCount > wf.loop {
		fmt.Println(tools.CustomTitle("\n          流程已全部执行结束...          \n"))
		os.Exit(1)
	} else {
		fmt.Println(tools.CustomTitle("\n          当前流程已结束, 开始执行下一个流程...          \n"))
		wf.loopCount++
	}

	if vt.onlyStandby {
		go func() {
			time.Sleep(time.Second)
			wf.sendJob(vt)
		}()
		return
	}

	if strings.Contains(vt.destination, "PQC") {
		vt.updateBlockTask()
	} else {
		vt.updateQCTask()
	}
	wf.sendJob(vt)
	//message := messages.GenerateRouteRequestJob(vt.vehicleID, vt.destination, vt.lane, "S", "5", vt.activity, 1, 40, 1)
	//if err := PublishAssignedTopic("route_request_job_instruction", "", message); err != nil {
	//	fmt.Printf("[%s] 任务下发失败: %s, 程序退出...", time.Now().Local().String(), err)
	//	os.Exit(1)
	//} else {
	//	fmt.Printf("[%s] send message to <%s>: %s\n", time.Now().Local().String(), "route_request_job_instruction", message)
	//}
}
