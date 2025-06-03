package fms

import (
	"encoding/json"
	"fmt"
	"time"
)

const prefix = "/fms/psa"

const (
	ManualModeURL       = prefix + `/vessel/%s/manualModel`
	OpBlockURL          = prefix + "/hatch_cover/op"
	ClearBlockURL       = prefix + "/hatch_cover/clear"
	GetVesselsURL       = prefix + "/vessels"
	GetAssignedVessel   = prefix + "/vessel"
	GetVehiclesURL      = prefix + "/vehicles"
	ResetVehicleURL     = prefix + "/truck"
	GetCraneLocationURL = prefix + "/cranes"

	LockURL    = prefix + `/vessel/%s/lock`
	ReleaseURL = prefix + `/vessel/%s/unlock`
)

const (
	SetCraneLocationURL = "/api/only_simulation_test"
)

type GetVesselsResponse struct {
	Status string  `json:"status"`
	Errno  int64   `json:"errno"`
	Msg    string  `json:"msg"`
	Data   Vessels `json:"data"`
}

type Vessels struct {
	Names  []string     `json:"names"`
	Values []VesselInfo `json:"values"`
}

func (vessel Vessels) String() string {
	v, _ := json.Marshal(vessel)
	return string(v)
}

type Coordinate struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
	Theta float64 `json:"theta"`
}

func (c Coordinate) String() string {
	v, _ := json.Marshal(c)
	return string(v)
}

type VesselInfo struct {
	MaxWf      int                    `json:"max_wf"`
	MinWf      int                    `json:"min_wf"`
	VesselInfo VesselDetail           `json:"vessel_info"`
	CAs        []VesselCAInfo         `json:"cas"`
	WAArrives  []VesselWaArrivedTruck `json:"wa_arrives"`
	CAArrives  []string               `json:"ca_arrives"`
	Cranes     []VesselCraneInfo      `json:"cranes"`
	Ingress    VesselGressInfo        `json:"ingress"`
	Egress     VesselGressInfo        `json:"egress"`
}

type VesselsInfo []VesselInfo

func (v VesselsInfo) Len() int {
	return len(v)
}

func (v VesselsInfo) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v VesselsInfo) Less(i, j int) bool {
	return v[i].VesselInfo.StartPos < v[j].VesselInfo.StartPos
}

func (ve VesselInfo) Wms() string {
	return fmt.Sprintf("%d-%d", ve.VesselInfo.StartPos, ve.VesselInfo.EndPos)
}

func (ve VesselInfo) Gress() string {
	return fmt.Sprintf("%d-%d", ve.Ingress.WharfMarkStart, ve.Egress.WharfMarkEnd)
}

func (ve VesselInfo) String() string {
	v, _ := json.Marshal(ve)
	return string(v)
}

type VesselWaArrivedTruck struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Block string `json:"block"`
	Bay   string `json:"bay"`
	Lane  int    `json:"lane"`
}

type VesselDetail struct {
	VesselId   string   `json:"vessel_id"`
	Cranes     []string `json:"cranes"`
	BerthNo    string   `json:"berth_no"`
	StartPos   int      `json:"start_pos"`
	EndPos     int      `json:"end_pos"`
	Direction  string   `json:"direction"`
	VesselName string   `json:"vessel_name"`
	Length     int      `json:"length"`
}

type VesselCAInfo struct {
	Locked         int        `json:"locked"`
	VesselId       string     `json:"vessel_id"`
	Width          int        `json:"width"`
	Capacity       int        `json:"capacity"`
	FixedWorkLane  *int       `json:"fixed_work_lane"`
	Vehicles       []string   `json:"vehicles"`
	Name           string     `json:"name"`
	Pos            Coordinate `json:"pos"`
	Lane           int        `json:"lane"`
	BindLane       int        `json:"bind_lane"`
	Index          *int       `json:"index"`
	WharfMarkStart int        `json:"wharf_mark_start"`
	WharfMarkEnd   int        `json:"wharf_mark_end"`
	Crane          string     `json:"crane"`
}

func (ca *VesselCAInfo) GetWorkLane() int {
	workLane := ca.BindLane
	if ca.FixedWorkLane != nil {
		workLane = *ca.FixedWorkLane
	}
	return workLane
}

type VesselCraneInfo struct {
	Locked            int        `json:"locked"`
	Type              string     `json:"type"`
	VehicleID         string     `json:"vehicle_id"`
	Name              string     `json:"name"`
	Pos               Coordinate `json:"pos"`
	WharfMark         int        `json:"wharf_mark"`
	Status            int        `json:"status"`
	LastStatus        int        `json:"last_status"`
	LastPos           Coordinate `json:"last_pos"`
	Moving            bool       `json:"moving"`
	MovementThreshold float64    `json:"movement_threshold"`
	StopDuration      int        `json:"stop_duration"`
	LastMovementTime  time.Time  `json:"last_movement_time"`
}

type VesselGressInfo struct {
	Name           string     `json:"name"`
	Width          int        `json:"width"`
	Start          Coordinate `json:"start"`
	End            Coordinate `json:"end"`
	WharfMarkStart int        `json:"wharf_mark_start"`
	WharfMarkEnd   int        `json:"wharf_mark_end"`
}

type VehiclesResponse struct {
	Status string      `json:"status"`
	Errno  interface{} `json:"errno"`
	Msg    interface{} `json:"msg"`
	Code   int         `json:"code"`
	Data   Vehicles    `json:"data"`
}

type VehiclesResponseData struct {
	ID                 string                 `json:"id"`
	VesselID           string                 `json:"vessel_id"`
	CanGoCallIn        bool                   `json:"can_go_call_in"`
	Arrived            bool                   `json:"arrived"`
	Destination        VehicleDestination     `json:"destination"`
	CurrentDestination VehicleCurrDestination `json:"current_destination"`
	LastDestination    VehicleCurrDestination `json:"last_destination"`
	Mode               string                 `json:"mode"`
	SSA                int                    `json:"ssa"`
	ReadyStatus        int                    `json:"ready_status"`
	ManualStatus       int                    `json:"manual_status"`
	HatchCover         string                 `json:"hatch_cover"`
	TaskInfo           *VehicleTaskInfo       `json:"task_info"`
}

type VehicleTaskInfo struct {
	Activity            int      `json:"activity"`
	TosRawID            int      `json:"tos_raw_id"`
	DestLocation        string   `json:"dest_location"`
	ContainerSize       string   `json:"container_size"`
	DockPosition        string   `json:"dock_position"`
	OnlyOne             bool     `json:"only_one"`
	ArrivedLocation     string   `json:"arrived_location"`
	TargetDockPositions []string `json:"target_dock_positions"`
	Containers          []string `json:"containers"`
	TargetLane          string   `json:"target_lane"`
	LiftType            int      `json:"lift_type"`
	Crane               string   `json:"crane"`
}

type Vehicles []VehiclesResponseData

func (v Vehicles) Len() int {
	return len(v)
}

func (v Vehicles) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Vehicles) Less(i, j int) bool {
	return v[i].ID < v[j].ID
}

type VehicleDestination struct {
	Name       string `json:"name"`
	Lane       int    `json:"lane"`
	Type       string `json:"type"`
	CreateTime string `json:"create_time"`
}

type VehicleCurrDestination struct {
	Type      string  `json:"type"`
	VehicleID string  `json:"vehicle_id"`
	Name      string  `json:"name"`
	Block     *string `json:"block"`
	Bay       *string `json:"bay"`
	Lane      int     `json:"lane"`
	CraneNo   string  `json:"crane_no"`
	StayThere bool    `json:"stay_there"`
	Weight    int     `json:"weight"`
}

type GetCranesResponse struct {
	Status string                `json:"status"`
	Errno  *int64                `json:"errno"`
	Msg    *string               `json:"msg"`
	Code   int                   `json:"code"`
	Data   GetCranesResponseData `json:"data"`
}

type GetCranesResponseData struct {
	Locked            int        `json:"locked"`
	Type              string     `json:"type"`
	Name              string     `json:"name"`
	WharfMark         int        `json:"wharf_mark"`
	Status            int        `json:"status"`
	LastStatus        int        `json:"last_status"`
	Moving            bool       `json:"moving"`
	MovementThreshold float64    `json:"movement_threshold"`
	StopDuration      int        `json:"stop_duration"`
	LastMovementTime  time.Time  `json:"last_movement_time"`
	Pos               Coordinate `json:"pos"`
	LastPos           Coordinate `json:"last_pos"`
	LatestPos         Coordinate `json:"latest_pos"`
}

type SetCraneLocationReq struct {
	DeviceID           string  `json:"deviceID"`
	HOPos              float64 `json:"HO_Pos"`
	TRPos              float64 `json:"TR_Pos"`
	SPRLocked          bool    `json:"SPR_Locked"`
	SpreaderType       string  `json:"Spreader_Type"`
	TRRun              bool    `json:"TR_Run"`
	CraneReady         bool    `json:"crane_ready"`
	CurrentLane        int     `json:"current_lane"`
	CurrentBayID       string  `json:"current_bay_id"`
	X                  string  `json:"x"`
	Y                  string  `json:"y"`
	GPSStatus          string  `json:"gps_status"`
	DisconnectCauseGPS string  `json:"disconnect_cause_gps"`
	CMSStatus          string  `json:"cms_status"`
	Size               string  `json:"size"`
	Loaction           string  `json:"location"`
	Height             int     `json:"height"`
	OpenClose          bool    `json:"open_close"`
	DisconnectCauseCMS string  `json:"disconnect_cause_cms"`
}

func (slr SetCraneLocationReq) String() string {
	v, _ := json.Marshal(slr)
	return string(v)
}

type OperateReq struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func (o OperateReq) String() string {
	v, _ := json.Marshal(o)
	return string(v)
}
