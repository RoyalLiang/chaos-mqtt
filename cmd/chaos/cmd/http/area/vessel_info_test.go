package area

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
	printVessels(vessels.Data.Values)
}
