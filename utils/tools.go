package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Masterminds/semver"

	"go.uber.org/zap"
)

const (
	timer int = 5
)

// ParseConfig parse assigned config file
//
//	cfgPath	string - config file path
//	cfgObj	configs.PonyConfig - config object
//func ParseConfig(cfgPath string, cfgObj configs.PonyConfig) (*configs.PonyConfig, error) {
//
//	data, err := os.ReadFile(cfgPath)
//	err = yaml.Unmarshal(data, &cfgObj)
//	if err != nil {
//		fmt.Println("parse error:", err)
//		return &cfgObj, err
//	}
//	return &cfgObj, nil
//}

// Execute run assigned command and return output([]byte) & error
//
//	w	string - command name
//	args	...string - command args
func Execute(w string, args ...string) ([]byte, error) {
	cmd := exec.Command(w, args...)
	zap.S().Warnf("[ Execute Command ] %v", cmd.String())
	out, err := cmd.CombinedOutput()
	return out, err
}

// WriteFile write data to file, if not exist, create. exist, overwrite
//
//	path	string - file path
//	content	string - file content
func WriteFile(path, content string) {

	var file *os.File
	var err error
	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(path)
		}
	} else {
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	}

	if err != nil {
		fmt.Sprintf("fail to create %s, error: %s\n", path, err)
		os.Exit(1)
	}

	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("[ File Write Error ] fail to write ", path, ": ", content)
		os.Exit(1)
	}
	fmt.Sprintf("[ File Write Success ] write to %v success\n", path)
}

// ReadFile read assigned file content
//
//	filePath	string - file path
func ReadFile(filePath string) ([]byte, error) {

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		zap.S().Debug("File does not exist:", filePath)
		return nil, err
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		zap.S().Debug("Error reading file:", err)
		return nil, err
	}
	return content, nil
}

// CheckConfig rewrite config file if missing some important key-value
//
//	dir		string - program execute root dir
//	ponyCfg	*configs.PonyConfig - config object, will rewrite
//func CheckConfig(dir string, ponyCfg *configs.PonyConfig) {
//
//	var flag = false
//	version := filepath.Base(dir)
//
//	if version != "" && ponyCfg.Core.Version != version {
//		ponyCfg.Core.Version = version
//		flag = true
//	}
//
//	if ponyCfg.Core.Dir == "" {
//		ponyCfg.Core.Dir = dir
//		flag = true
//	}
//
//	if ponyCfg.Core.UUID == "" {
//		uid, _ := uuid.NewUUID()
//		ponyCfg.Core.UUID = uid.String()
//		flag = true
//	}
//
//	if flag {
//		configs.WriteConfig(viper.GetViper(), "core", ponyCfg.Core)
//	}
//}

// GetRootDir get root dir of current executable
func GetRootDir() string {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	return dir
}

// VersionCompare compare version, return true if nv > cv
//
//	cv	string - current version
//	nv	string - new version
func VersionCompare(cv, nv string) bool {
	cu, _ := semver.NewVersion(cv)
	ne, _ := semver.NewVersion(nv)

	if cu.Compare(ne) == -1 {
		return true
	}
	return false
}

func GetVehicleTaskID(vehicleID string) string {
	now := time.Now()

	formattedTime := now.Format("200601021504051")
	return "FAKE" + vehicleID + formattedTime
}
