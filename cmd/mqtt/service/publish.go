package service

import (
	"fms-awesome-tools/constants"
	"fmt"
	"os"
	"text/template"
)

func getTemplateMessage(content string, param interface{}) string {
	t, _ := template.New("").Parse(content)
	_ = t.Execute(os.Stdout, &param)
	a := fmt.Sprintf("%v", t)
	fmt.Println("==========")
	fmt.Println(a)
	return a
}

func PublishAssignedTopic(topic, vehicleID string, activity int64, args ...string) error {

	fmt.Println("Publishing assigned topic:", topic)
	fmt.Println("Vehicle ID:", vehicleID)
	fmt.Println("Activity:", activity)
	return nil
}

func PublishRouteRequestJobInstruction(topic, vehicleID, destination string, lane, activity int64) error {
	param := constants.VehicleParam{
		ID:               "001",
		VehicleID:        vehicleID,
		Activity:         activity,
		NextLocationLane: lane,
		NextLocation:     destination,
	}
	message := getTemplateMessage(topic, param)
	fmt.Println("message: ", message)
	return MQTTClient.Publish(topic, message)
}

func PublishJobInstruction() error {
	return nil
}

func PublishDockPosition() error {
	return nil
}

func PublishCallInRequest() error {
	return nil
}

func PublishVesselBerth() error {
	return nil
}

func PublishMoveToQC() error {
	return nil
}

func PublishIngressToCallIn() error {
	return nil
}

func PublishCancelJob() error {
	return nil
}
