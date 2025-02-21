package fms

import (
	"encoding/json"
	"time"
)

const prefix = "/fms/psa"

const (
	ManualModeURL     = ""
	SetBlockURL       = prefix + "/hatch_cover/op"
	GetVesselsURL     = prefix + "/vessels"
	GetAssignedVessel = prefix + "/vessel"
	GetVehiclesURL    = prefix + "/vehicles"
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

type VesselInfo struct {
	MaxWf      int               `json:"max_wf"`
	MinWf      int               `json:"min_wf"`
	VesselInfo VesselDetail      `json:"vessel_info"`
	CAs        []VesselCAInfo    `json:"cas"`
	WAArrives  []string          `json:"wa_arrives"`
	CAArrives  []string          `json:"ca_arrives"`
	Cranes     []VesselCraneInfo `json:"cranes"`
	Ingress    VesselGressInfo   `json:"ingress"`
	Egress     VesselGressInfo   `json:"egress"`
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
	Status string                 `json:"status"`
	Errno  interface{}            `json:"errno"`
	Msg    interface{}            `json:"msg"`
	Code   int                    `json:"code"`
	Data   []VehiclesResponseData `json:"data"`
}

type VehiclesResponseData struct {
	ID                 string                 `json:"id"`
	VesselID           string                 `json:"vessel_id"`
	CanGoCallIn        bool                   `json:"can_go_call_in"`
	Arrived            bool                   `json:"arrived"`
	Destination        VehicleDestination     `json:"destination"`
	CurrentDestination VehicleCurrDestination `json:"current_destination"`
	LastDestination    VehicleCurrDestination `json:"last_destination"`
}

type VehicleDestination struct {
	Name string `json:"name"`
	Lane int    `json:"lane"`
	Type string `json:"type"`
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
