package service

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkflow(t *testing.T) {
	w := NewWorkflow()
	if err := w.StartWorkflow(); err != nil {
		fmt.Println("Failed to start workflow:", err)
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)
	w.Close()
}
