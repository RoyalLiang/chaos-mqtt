package requests

import (
	"testing"
)

func TestPrintResult(t *testing.T) {
	vessels := getVessels()

	if vessels == nil {
		return
	}
	if vessels.Data.Values == nil {
		return
	}
	parseVesselInfo(vessels.Data.Values)
}
