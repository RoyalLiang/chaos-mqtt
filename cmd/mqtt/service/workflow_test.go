package service

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestWorkflow(t *testing.T) {
	workflow := Workflow{
		UUID:     uuid.NewString(),
		Truck:    "AT001",
		Flow:     "QC",
		Activity: 1,
	}

	if err := workflow.StartWorkflow(); err != nil {
		fmt.Println("Failed to start workflow:", err)
		t.Fatal(err)
	}
}
