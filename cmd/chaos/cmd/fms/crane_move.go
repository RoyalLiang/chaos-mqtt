package area

import (
	"context"
	"encoding/json"
	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/configs"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	moveCrane    string
	moveDistance int64
	moveTime     int64
)

var CraneMoveCmd = &cobra.Command{
	Use:   "crane_move",
	Short: "模拟QC移动",
	Run: func(cmd *cobra.Command, args []string) {

	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if moveCrane == "" && moveDistance == 0 && moveTime == 0 {
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

	coordinate := pos
	ctx := context.Background()
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.Tick(time.Duration(moveTime) * time.Second):
			<-ctx.Done()
		case <-time.After(time.Second):
			sendRequest(coordinate)
			coordinate = calcCoordinate(coordinate)
		}
	}
}

func getOffset(offset float64) (float64, float64) {
	rad := 31.0 / 180.0 * math.Pi
	x := offset * math.Cos(rad)
	y := offset * math.Sin(rad)
	return x, y
}

func calcCoordinate(coordinate *fms.Coordinate) *fms.Coordinate {
	xoff, yoff := getOffset(12.05418)

	theta := -2.11 - math.Pi
	if moveDistance > 0 {
		theta = -2.11 + math.Pi
	}

	coordinate.X += float64(moveDistance) * math.Cos(theta)
	coordinate.Y += float64(moveDistance) * math.Sin(theta)
	fmt.Println("xy: ", coordinate)
	coordinate.X += xoff
	coordinate.Y -= yoff
	return coordinate
}

func sendRequest(coordinate *fms.Coordinate) {
	url := configs.Chaos.FMS.CraneManager.Address + fms.SetCraneLocationURL
	req := fms.SetCraneLocationReq{
		DeviceID: moveCrane, HOPos: 15108.6240234375, TRPos: 246.0, SPRLocked: true, SpreaderType: "40",
		TRRun: true, CraneReady: true, CurrentBayID: "MT9.250", CurrentLane: 3, GPSStatus: "normal",
		X: strconv.FormatFloat(coordinate.X, 'E', -1, 64), DisconnectCauseGPS: "1", Height: 12934,
		Y: strconv.FormatFloat(coordinate.Y, 'E', -1, 64), CMSStatus: "1", Size: "0", Loaction: "6",
		OpenClose: true, DisconnectCauseCMS: "2",
	}

	resp, err := fms.Post(url, []byte(req.String()))
	if err != nil {
		cobra.CheckErr(err)
	}
	fmt.Println(string(resp))
}

func init() {
	CraneMoveCmd.Flags().StringVarP(&moveCrane, "crane", "c", "", "移动的岸桥号")
	CraneMoveCmd.Flags().Int64VarP(&moveDistance, "distance", "d", 1, "岸桥每次移动的距离\n>0 wharf mark ⬆️\n<0 wharf mark ⬇️\n")
	CraneMoveCmd.Flags().Int64VarP(&moveTime, "time", "t", 0, "移动时间")
	CraneMoveCmd.MarkFlagsRequiredTogether("distance", "time")
}
