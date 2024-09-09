package internal

import (
	"encoding/json"
	"reflect"

	"fms-awesome-tools/cmd/chaos/internal/messages"
)

func parse(object interface{}, content []byte) interface{} {
	if err := json.Unmarshal(content, object); err != nil {
		panic(err)
	}
	return object
}

func ParseToTask(content []byte) messages.Task {
	obj := parse(&messages.Task{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.Task)
}

func ParseToRouteResponseJobInstruction(content []byte) messages.RouteResponseJobInstruction {
	obj := parse(&messages.RouteResponseJobInstruction{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.RouteResponseJobInstruction)
}

func ParseToSwitchModeResponse(content []byte) messages.SwitchModeResponse {
	obj := parse(&messages.SwitchModeResponse{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.SwitchModeResponse)
}

func ParseToReadyForIngressToCallIn(content []byte) messages.ReadyForIngressToCallIn {
	obj := parse(&messages.ReadyForIngressToCallIn{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.ReadyForIngressToCallIn)
}

func ParseToReadyForMoveToQC(content []byte) messages.ReadyForMoveToQC {
	obj := parse(&messages.ReadyForMoveToQC{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.ReadyForMoveToQC)
}

func ParseToLogonRequest(content []byte) messages.LogonRequest {
	obj := parse(&messages.LogonRequest{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.LogonRequest)
}

func ParseToApmArrivedRequest(content []byte) messages.APMArrivedRequest {
	obj := parse(&messages.APMArrivedRequest{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.APMArrivedRequest)
}

func ParseToDockPositionResponse(content []byte) messages.DockPositionResponse {
	obj := parse(&messages.DockPositionResponse{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.DockPositionResponse)
}

func ParseToPathUpdateAvailable(content []byte) messages.PathUpdateAvailable {
	obj := parse(&messages.PathUpdateAvailable{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.PathUpdateAvailable)
}

func ParseToMountInstructionResponse(content []byte) messages.MountInstructionResponse {
	obj := parse(&messages.MountInstructionResponse{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.MountInstructionResponse)
}

func ParseToOffloadInstructionResponse(content []byte) messages.OffloadInstructionResponse {
	obj := parse(&messages.OffloadInstructionResponse{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.OffloadInstructionResponse)
}

func ParseToWharfDockPositionRequest(content []byte) messages.WharfDockPositionRequest {
	obj := parse(&messages.WharfDockPositionRequest{}, content)
	return reflect.ValueOf(obj).Elem().Interface().(messages.WharfDockPositionRequest)
}

func GenerateToLogonResponse(data messages.LogonRequest) messages.LogonResponse {
	return messages.LogonResponse{APMID: data.APMID, Data: messages.LogonResponseData{
		Success: 1, TrailerSeqNumbers: data.Data.TrailerSeqNumbers, TrailerNumbers: data.Data.TrailerNumbers,
		TrailerHeights: make([]int, 0), TrailerLengths: make([]int, 0), TrailerPayloads: make([]int, 0),
		TrailerTypes: make([]string, 0), TrailerUnladenWeights: make([]int, 0), TrailerWidths: make([]int, 0),
		NumTrailers: data.Data.NumTrailers,
	}}
}
