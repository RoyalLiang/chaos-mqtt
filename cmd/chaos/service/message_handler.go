package service

import (
	"encoding/json"
	"fmt"

	"fms-awesome-tools/cmd/chaos/internal"
	"fms-awesome-tools/cmd/chaos/internal/messages"
)

func (wf *Workflow) response(topic, message string) {
	if err := wf.client.Publish("logon_response", message); err != nil {
		fmt.Printf("%s ==> %s\n", topic, err)
	} else {
		fmt.Printf("%s ==> %s\n", topic, message)
	}
}

func (wf *Workflow) logonHandler(message []byte) {
	data := internal.ParseToLogonRequest(message)
	resp := internal.GenerateToLogonResponse(data)
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
		return
	}
	fmt.Println("job_instruction == > route dag为空, 忽略处理")
}

func (wf *Workflow) readyForIngressToCallInHandler(message []byte) {

}

func (wf *Workflow) readyForMoveToQCheHandler(message []byte) {

}

func (wf *Workflow) apmArrivalHandler(message []byte) {

}

func (wf *Workflow) pathUpdateRequestHandler(message []byte) {

}

func (wf *Workflow) dockPositionResponseHandler(message []byte) {

}
