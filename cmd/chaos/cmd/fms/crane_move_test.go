package area

import (
	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fmt"
	"testing"
)

func TestCalcCranePos(t *testing.T) {
	moveDistance = 3

	c := &fms.Coordinate{
		X: 763.4479,
		Y: 490.4556,
	}

	calcCoordinate(c)
	fmt.Println(c)
}
