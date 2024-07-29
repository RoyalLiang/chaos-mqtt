package main

import (
	"os"

	"fms-awesome-tools/cmd/mqtt/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
