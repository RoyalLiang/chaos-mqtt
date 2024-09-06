package internal

import (
	"encoding/json"
	"fms-awesome-tools/cmd/chaos/internal/messages"
	"reflect"
)

func parse(object interface{}, content []byte) interface{} {
	if err := json.Unmarshal(content, object); err != nil {
		panic(err)
	}
	return object
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

func GenerateToLogonResponse(data messages.LogonRequest) messages.LogonResponse {
	return messages.LogonResponse{APMID: data.APMID, Data: messages.LogonResponseData{
		Success: 1, TrailerSeqNumbers: data.Data.TrailerSeqNumbers, TrailerNumbers: data.Data.TrailerNumbers,
		TrailerHeights: make([]int, 0), TrailerLengths: make([]int, 0), TrailerPayloads: make([]int, 0),
		TrailerTypes: make([]string, 0), TrailerUnladenWeights: make([]int, 0), TrailerWidths: make([]int, 0),
		NumTrailers: data.Data.NumTrailers,
	}}
}
