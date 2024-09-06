package service

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkflow(t *testing.T) {
	if err := StartWorkflow(); err != nil {
		fmt.Println("Failed to start workflow:", err)
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)
	Close()
}
