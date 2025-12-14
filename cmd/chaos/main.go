package main

import (
	"fms-awesome-tools/pkg/logger"
	"os"

	"fms-awesome-tools/cmd/chaos/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Errorf("主程序错误: %s", err.Error())
		os.Exit(1)
	}
}
