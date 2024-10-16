package messages

import (
	"encoding/json"
	"fms-awesome-tools/constants"
	tools "fms-awesome-tools/utils"
	"strconv"
	"strings"
)

type RouteRequestJobInstructionRequest struct {
	APMID string                                `json:"apm_id"`
	Data  RouteRequestJobInstructionRequestData `json:"data"`
}

type RouteRequestJobInstructionRequestData struct {
	APMID                   string     `json:"apm_id"`
	ID                      string     `json:"id"`
	Timestamp               int64      `json:"timestamp"`
	Message                 string     `json:"message"`
	MapVersion              string     `json:"map_version"`
	RouteDAG                []RouteDag `json:"route_dag"`
	RouteMandate            string     `json:"route_mandate"`
	OperationalGroup        string     `json:"operational_group"`
	LiftType                int64      `json:"lift_type"`
	APMDirection            string     `json:"apm_direction"`
	DualCycle               string     `json:"dual_cycle"`
	Activity                int64      `json:"activity"`
	NextLocation            string     `json:"next_location"`
	NextLocationLane        string     `json:"next_location_lane"`
	AssignedCntrSize        string     `json:"assigned_cntr_size"`
	AssignedCntrType        string     `json:"assigned_cntr_type"`
	NumMountedCntr          int64      `json:"num_mounted_cntr"`
	TargetDockPosition      string     `json:"target_dock_position"`
	OperationalTypes        []string   `json:"operational_types"`
	JobTypes                []string   `json:"job_types"`
	OperationalGroups       []string   `json:"operational_groups"`
	CntrNumbers             []string   `json:"cntr_numbers"`
	CntrSizes               []string   `json:"cntr_sizes"`
	CntrWeights             []string   `json:"cntr_weights"`
	CntrCategorys           []string   `json:"cntr_categorys"`
	CntrTypes               []string   `json:"cntr_types"`
	CntrStatus              []string   `json:"cntr_status"`
	DGGroups                []string   `json:"dg_groups"`
	IMOClass                []string   `json:"imo_class"`
	ReferTemperatures       []string   `json:"refer_temperatures"`
	CntrLocationsOnAPM      []int      `json:"cntr_locations_on_apm"`
	SourceLocations         []string   `json:"source_locations"`
	DestLocations           []string   `json:"dest_locations"`
	OffloadSequences        []string   `json:"offload_sequences"`
	Cones                   []string   `json:"cones"`
	OperationalQCSequences  []string   `json:"operational_qc_sequences"`
	OperationalJobSequences []string   `json:"operational_job_sequences"`
	TrailerPositions        []string   `json:"trailer_positions"`
	WeightClass             []string   `json:"weight_class"`
	UrGents                 []string   `json:"urgents"`
	DGs                     []string   `json:"dgs"`
	PlugRequireds           []string   `json:"plug_requireds"`
	MotorDirections         []string   `json:"motor_directions"`
}

func (r *RouteRequestJobInstructionRequest) String() string {
	v, _ := json.Marshal(r)
	return string(v)
}

type LogonResponseData struct {
	TrailerSeqNumbers     []int    `json:"trailer_seq_numbers"`
	TrailerUnladenWeights []int    `json:"trailer_unladen_weights"`
	Success               int64    `json:"success"`
	TrailerLengths        []int    `json:"trailer_lengths"`
	TrailerWidths         []int    `json:"trailer_widths"`
	TrailerHeights        []int    `json:"trailer_heights"`
	TrailerTypes          []string `json:"trailer_types"`
	TrailerPayloads       []int    `json:"trailer_payloads"`
	TrailerNumbers        []string `json:"trailer_numbers"`
	Message               string   `json:"message"`
	NumTrailers           int      `json:"num_trailers"`
	Timestamp             int64    `json:"timestamp"`
}

type LogonResponse struct {
	APMID string            `json:"apm_id"`
	Data  LogonResponseData `json:"data"`
}

func (lr LogonResponse) String() string {
	v, _ := json.Marshal(lr)
	return string(v)
}

type JobInstruction struct {
	APMID string             `json:"apm_id"`
	Data  JobInstructionData `json:"data"`
}

type JobInstructionData struct {
	APMID                   string     `json:"apm_id"`
	ID                      string     `json:"id"`
	Timestamp               int64      `json:"timestamp"`
	MapVersion              string     `json:"map_version"`
	RouteDAG                []RouteDag `json:"route_dag"`
	RouteMandate            string     `json:"route_mandate"`
	OperationalGroup        string     `json:"operational_group"`
	LiftType                int64      `json:"lift_type"`
	APMDirection            string     `json:"apm_direction"`
	DualCycle               string     `json:"dual_cycle"`
	Activity                int64      `json:"activity"`
	NextLocation            string     `json:"next_location"`
	NextLocationLane        string     `json:"next_location_lane"`
	AssignedCntrSize        string     `json:"assigned_cntr_size"`
	AssignedCntrType        string     `json:"assigned_cntr_type"`
	NumMountedCntr          int        `json:"num_mounted_cntr"`
	TargetDockPosition      string     `json:"target_dock_position"`
	OperationalTypes        []string   `json:"operational_types"`
	JobTypes                []string   `json:"job_types"`
	OperationalGroups       []string   `json:"operational_groups"`
	CNTRNumbers             []string   `json:"cntr_numbers"`
	CNTRSizes               []string   `json:"cntr_sizes"`
	CNTRWeights             []float64  `json:"cntr_weights"`
	CNTRCategorys           []string   `json:"cntr_categorys"`
	CNTRTypes               []string   `json:"cntr_types"`
	CNTRStatus              []string   `json:"cntr_status"`
	DGGroups                []string   `json:"dg_groups"`
	IMOClass                []string   `json:"imo_class"`
	ReferTemperatures       []float64  `json:"refer_temperatures"`
	CNTRLocationsOnAPM      []string   `json:"cntr_locations_on_apm"`
	SourceLocations         []string   `json:"source_locations"`
	DestLocations           []string   `json:"dest_locations"`
	OffloadSequences        []string   `json:"offload_sequences"`
	Cones                   []string   `json:"cones"`
	OperationalQCSequences  []string   `json:"operational_qc_sequences"`
	OperationalJobSequences []string   `json:"operational_job_sequences"`
	TrailerPositions        []string   `json:"trailer_positions"`
	WeightClass             []string   `json:"weight_class"`
	Urgents                 []string   `json:"urgents"`
	DGS                     []string   `json:"dgs"`
	PlugRequireds           []string   `json:"plug_requireds"`
	MotorDirections         []string   `json:"motor_directions"`
	RouteType               string     `json:"route_type"`
}

func (job JobInstruction) String() string {
	v, _ := json.Marshal(job)
	return string(v)
}

type MountInstruction struct {
	APMID string               `json:"apm_id"`
	Data  MountInstructionData `json:"data"`
}

type MountInstructionData struct {
	ID                     string `json:"id"`
	Timestamp              int64  `json:"timestamp"`
	OperationalType        string `json:"operational_type"`
	JobType                string `json:"job_type"`
	OperationalGroup       string `json:"operational_group"`
	CntrNumber             string `json:"cntr_number"`
	CntrSize               string `json:"cntr_size"`
	CntrWeight             int    `json:"cntr_weight"`
	CntrCategory           string `json:"cntr_category"`
	CntrType               string `json:"cntr_type"`
	CntrStatus             string `json:"cntr_status"`
	DGGroup                string `json:"dg_group"`
	IMOClass               string `json:"imo_class"`
	ReeferTemperature      string `json:"reefer_temperature"`
	CntrLocationOnAPM      int    `json:"cntr_location_on_apm"`
	SourceLocation         string `json:"source_location"`
	DestLocation           string `json:"dest_location"`
	OffloadSequence        string `json:"offload_sequence"`
	Cone                   int    `json:"cone"`
	OperationalQCSequence  int    `json:"operational_qc_sequence"`
	OperationalJobSequence int    `json:"operational_job_sequence"`
	TrailerPosition        string `json:"trailer_position"`
	WeightClass            string `json:"weight_class"`
	Urgent                 string `json:"urgent"`
	DG                     string `json:"dg"`
	PlugRequired           string `json:"plug_required"`
	LiftType               int    `json:"lift_type"`
	Message                string `json:"message"`
}

func (m MountInstruction) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

type PathUpdateRequest struct {
	APMID string                `json:"apm_id"`
	Data  PathUpdateRequestData `json:"data"`
}

type PathUpdateRequestData struct {
	ID                  string     `json:"id"`
	DestinationName     string     `json:"destination_name"`
	TargetDockPosition  string     `json:"target_dock_position"`
	DestinationLane     string     `json:"destination_lane"`
	DestinationWaypoint string     `json:"destination_waypoint"`
	RouteDag            []RouteDag `json:"route_dag"`
	RouteType           string     `json:"route_type"`
	Timestamp           int64      `json:"timestamp"`
}

func (p PathUpdateRequest) String() string {
	v, _ := json.Marshal(p)
	return string(v)
}

type OffloadInstruction struct {
	APMID string                 `json:"apm_id"`
	Data  OffloadInstructionData `json:"data"`
}

type OffloadInstructionData struct {
	ID         string `json:"id"`
	Timestamp  int64  `json:"timestamp"`
	CntrNumber string `json:"cntr_number"`
	Message    string `json:"message"`
}

func (o OffloadInstruction) String() string {
	v, _ := json.Marshal(o)
	return string(v)
}

type WharfDockPositionResponse struct {
	APMID string                        `json:"apm_id"`
	Data  WharfDockPositionResponseData `json:"data"`
}

type WharfDockPositionResponseData struct {
	APMID         string `json:"apm_id"`
	Success       int    `json:"success"`
	RejectionCode string `json:"rejection_code"`
	ID            string `json:"id"`
	Timestamp     int64  `json:"timestamp"`
}

type CallInRequest struct {
	APMID string            `json:"apm_id"`
	Data  CallInRequestData `json:"data"`
}

type CallInRequestData struct {
	ID         string `json:"id"`
	Crane      string `json:"qc_number"`
	CallInMode int64  `json:"call_in_mode"`
}

func (c CallInRequest) String() string {
	v, _ := json.Marshal(c)
	return string(v)
}

type MoveToQCRequest struct {
	APMID string              `json:"apm_id"`
	Data  MOveToQCRequestData `json:"data"`
}

type MOveToQCRequestData struct {
	ID                  string     `json:"id"`
	Timestamp           int64      `json:"timestamp"`
	DestinationName     string     `json:"destination_name"`
	TargetDockPosition  string     `json:"target_dock_position"`
	ContainerSize       string     `json:"container_size"`
	DestinationLane     string     `json:"destination_lane"`
	DestinationWaypoint string     `json:"destination_waypoint"`
	RouteDag            []RouteDag `json:"route_dag"`
	RouteType           string     `json:"route_type"`
}

func (m MoveToQCRequest) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

type IngressToCallIn struct {
	APMID string              `json:"apm_id"`
	Data  IngressToCallInData `json:"data"`
}

type IngressToCallInData struct {
	ID                  string     `json:"id"`
	Timestamp           int64      `json:"timestamp"`
	DestinationName     string     `json:"destination_name"`
	DestinationWaypoint string     `json:"destination_waypoint"`
	RouteDag            []RouteDag `json:"route_dag"`
	DestinationLane     string     `json:"destination_lane"`
	LaneAvailability    []string   `json:"lane_availability"`
	IngressBS           string     `json:"ingress_b_s"`
	IngressSB           string     `json:"ingress_s_b"`
	PmDirection         string     `json:"pm_direction"`
	WharfStretchId      string     `json:"wharf_stretch_id"`
	RouteType           string     `json:"route_type"`
	Message             string     `json:"message"`
}

func (m IngressToCallIn) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

type FunctionalJobData struct {
	Id           string     `json:"id"`
	Timestamp    int64      `json:"timestamp"`
	MapVersion   string     `json:"map_version"`
	RouteDag     []RouteDag `json:"route_dag"`
	RouteType    string     `json:"route_type"`
	DestLocation string     `json:"dest_location"`
}

type FunctionalJob struct {
	APMID string            `json:"apm_id"`
	Data  FunctionalJobData `json:"data"`
}

func (f FunctionalJob) String() string {
	v, _ := json.Marshal(f)
	return string(v)
}

type SwitchModeRequestData struct {
	ID        string `json:"id"`
	SetMode   string `json:"set_mode"`
	Timestamp int64  `json:"timestamp"`
}

type SwitchModeRequest struct {
	ApmId string                `json:"apm_id"`
	Data  SwitchModeRequestData `json:"data"`
}

func (s SwitchModeRequest) String() string {
	v, _ := json.Marshal(s)
	return string(v)
}

type RouteRequestData struct {
	Timestamp  int64  `json:"timestamp"`
	Id         string `json:"id"`
	MapVersion string `json:"map_version"`
	Type       string `json:"type"`
	Data       string `json:"data"`
}

type RouteRequest struct {
	ApmId string           `json:"apm_id"`
	Data  RouteRequestData `json:"data"`
}

func (r RouteRequest) String() string {
	v, _ := json.Marshal(r)
	return string(v)
}

func GenerateRouteRequestJob(destination, lane, targetDockPos string, liftSize, container, quantity int64) string {
	var dest = ""
	if constants.Activity == 2 || constants.Activity == 6 {
		if strings.HasPrefix(destination, "PQC") {
			dest = "P," + destination + "          "
		} else {
			dest = destination
		}
	}

	routeJob := RouteRequestJobInstructionRequest{
		APMID: constants.VehicleID,
		Data: RouteRequestJobInstructionRequestData{
			APMID: constants.VehicleID, RouteDAG: make([]RouteDag, 0), Activity: constants.Activity,
			ID: tools.GetVehicleTaskID(constants.VehicleID, dest, constants.Activity), NextLocation: dest, NextLocationLane: lane,
			TargetDockPosition: targetDockPos, LiftType: liftSize, OperationalTypes: make([]string, 0),
			CntrCategorys: make([]string, 0), CntrStatus: make([]string, 0), CntrWeights: make([]string, 0),
			CntrNumbers: make([]string, 0), CntrSizes: make([]string, 0), CntrTypes: make([]string, 0),
			Cones: make([]string, 0), CntrLocationsOnAPM: make([]int, 0), OperationalJobSequences: make([]string, 0),
			OperationalGroups: make([]string, 0), OperationalQCSequences: make([]string, 0), JobTypes: make([]string, 0),
			UrGents: make([]string, 0), DestLocations: make([]string, 0), DGGroups: make([]string, 0),
			DGs: make([]string, 0), ReferTemperatures: make([]string, 0), IMOClass: make([]string, 0),
			OffloadSequences: make([]string, 0), TrailerPositions: make([]string, 0), WeightClass: make([]string, 0),
			PlugRequireds: make([]string, 0), SourceLocations: make([]string, 0), MotorDirections: make([]string, 0),
			AssignedCntrType: "GP", NumMountedCntr: 0, DualCycle: "N",
		},
	}

	routeJob.Data.CntrSizes = append(routeJob.Data.CntrSizes, strconv.FormatInt(container, 10))
	switch constants.Activity {
	case 2:
		routeJob.Data.AssignedCntrSize = strconv.FormatInt(container, 10)
	case 6:
		if container >= 40 {
			routeJob.Data.OffloadSequences = append(routeJob.Data.OffloadSequences, "FFFF0000000")
			routeJob.Data.CntrLocationsOnAPM = append(routeJob.Data.CntrLocationsOnAPM, 5)
			routeJob.Data.OffloadSequences = append(routeJob.Data.OffloadSequences, "0"+strconv.FormatInt(5, 64))
		}
		for c := range quantity {
			routeJob.Data.OffloadSequences = append(routeJob.Data.OffloadSequences, "FFFF0000000")
			routeJob.Data.CntrLocationsOnAPM = append(routeJob.Data.CntrLocationsOnAPM, int(c))
			routeJob.Data.OffloadSequences = append(routeJob.Data.OffloadSequences, "0"+strconv.FormatInt(c, 64))
		}
	}
	return routeJob.String()
}
