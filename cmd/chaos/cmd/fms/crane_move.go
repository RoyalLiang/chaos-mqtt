package area

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/configs"

	"github.com/spf13/cobra"
)

var (
	moveCrane    string
	moveDistance int64
	moveTime     int64
	coordinate   *c
)

var CraneMoveCmd = &cobra.Command{
	Use:   "crane_move",
	Short: "模拟QC移动",
	Run: func(cmd *cobra.Command, args []string) {

	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if moveCrane == "" && moveTime == 0 {
			_ = cmd.Help()
			os.Exit(1)
		}

		if moveTime < 0 {
			cobra.CheckErr("移动时间必须大于0")
		}

		if moveDistance == 0 {
			cobra.CheckErr("移动距离不可为0")
		}

		craneMOve()
	},
}

type c struct {
	offsetCoordinate *fms.Coordinate
	centerCoordinate *fms.Coordinate
}

func getCraneData() (*fms.Coordinate, error) {
	url := configs.Chaos.FMS.Area.Address + fms.GetCraneLocationURL + "?crane_no=" + moveCrane
	resp, err := fms.Get(url)
	if err != nil {
		fmt.Println("", err.Error())
		return nil, fmt.Errorf("%s: %s", "获取岸桥信息失败", err.Error())
	}

	craneInfo := &fms.GetCranesResponse{}
	if err = json.Unmarshal(resp, craneInfo); err != nil {
		return nil, fmt.Errorf("%s: %s", "解析岸桥信息失败", err.Error())
	}
	return &craneInfo.Data.LatestPos, nil
}

func craneMOve() {
	pos, err := getCraneData()
	if err != nil {
		cobra.CheckErr(err)
	}
	if pos == nil {
		cobra.CheckErr("无法获取岸桥坐标...")
	}

	coordinate.centerCoordinate = pos
	ticker := time.NewTicker(time.Duration(moveTime) * time.Second)
	defer ticker.Stop()

Loop:
	for {
		select {
		case <-ticker.C:
			break Loop
		case <-time.After(time.Second):
			sendRequest(coordinate)
			coordinate = calcCoordinate(coordinate)
		}
	}
	fmt.Println("运行结束...")
}

func getOffset(offset float64) (float64, float64) {
	rad := 31.0 / 180.0 * math.Pi
	x := offset * math.Cos(rad)
	y := offset * math.Sin(rad)
	return x, y
}

func calcCoordinate(coordinate *c) *c {
	xoff, yoff := getOffset(12.05418)

	theta := -2.11 - math.Pi
	if moveDistance > 0 {
		theta = -2.11 + math.Pi
	}

	coordinate.centerCoordinate.X += float64(moveDistance) * math.Cos(theta)
	coordinate.centerCoordinate.Y += float64(moveDistance) * math.Sin(theta)
	coordinate.offsetCoordinate.X = coordinate.centerCoordinate.X + xoff
	coordinate.offsetCoordinate.Y = coordinate.centerCoordinate.Y - yoff
	return coordinate
}

func sendRequest(coordinate *c) {
	url := configs.Chaos.FMS.CraneManager.Address + fms.SetCraneLocationURL
	req := fms.SetCraneLocationReq{
		DeviceID: moveCrane, HOPos: 15108.6240234375, TRPos: 246.0, SPRLocked: true, SpreaderType: "40",
		TRRun: true, CraneReady: true, CurrentBayID: "MT9.250", CurrentLane: 3, GPSStatus: "normal",
		X: strconv.FormatFloat(coordinate.offsetCoordinate.X, 'E', -1, 64), DisconnectCauseGPS: "1", Height: 12934,
		Y: strconv.FormatFloat(coordinate.offsetCoordinate.Y, 'E', -1, 64), CMSStatus: "1", Size: "0", Loaction: "6",
		OpenClose: true, DisconnectCauseCMS: "2",
	}

	_, err := fms.Post(url, []byte(req.String()))
	if err != nil {
		cobra.CheckErr(err)
	}
	fmt.Printf("岸桥当前位置: x: %.5f, y: %.5f\n", coordinate.centerCoordinate.X, coordinate.centerCoordinate.Y)
}

func init() {
	coordinate = &c{
		offsetCoordinate: &fms.Coordinate{},
		centerCoordinate: &fms.Coordinate{},
	}

	CraneMoveCmd.Flags().StringVarP(&moveCrane, "crane", "c", "", "岸桥号🌉")
	CraneMoveCmd.Flags().Int64VarP(&moveDistance, "distance", "d", 1, "单次移动距离\n>0 wm⬆️\n<0 wm⬇️\n")
	CraneMoveCmd.Flags().Int64VarP(&moveTime, "time", "t", 0, "移动时间⏰")
	CraneMoveCmd.MarkFlagsRequiredTogether("distance", "time")
}
