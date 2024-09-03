package main

import (
	"os"
	"text/template"

	"fms-awesome-tools/constants"
)

func main() {
	t, _ := template.New("").Parse(constants.RouteRequestJobInstruction)
	t.
		_ = t.Execute(os.Stdout, constants.VehicleParam{
		ID: "010101", VehicleID: "AT001", NextLocation: "P,PQC9234", NextLocationLane: 2, LiftType: 1, TargetDockPosition: 5,
	})
	//fmt.Sprintf(constants.RouteRequestJobInstruction, "AT001")
}
