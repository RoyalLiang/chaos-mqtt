package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"fms-awesome-tools/constants"
)

func main() {
	t, _ := template.New("").Parse(constants.RouteRequestJobInstruction)
	_ = t.Execute(os.Stdout, constants.VehicleParam{
		ID: "010101", VehicleID: "AT001", NextLocation: "P,PQC9234", NextLocationLane: 2, LiftType: 1, TargetDockPosition: 5,
	})
	a := fmt.Sprintf("%v", t)
	fmt.Println(json.Marshal(a))
}
