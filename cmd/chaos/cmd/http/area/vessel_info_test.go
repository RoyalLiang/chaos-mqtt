package area

import (
	"testing"
)

func TestPrintResult(te *testing.T) {
	vessels := getVessels()

	t.AppendHeader(header)
	if vessels == nil {
		return
	}
	if vessels.Data.Values == nil {
		return
	}
	printVessels(vessels.Data.Values)
}
