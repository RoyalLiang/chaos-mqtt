package area

import "testing"

func TestPrintVehicles(t *testing.T) {
	vehicles := getVehicles()
	printVehicles(vehicles)
}
