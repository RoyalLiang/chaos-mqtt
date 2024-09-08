package messages

import "encoding/json"

type LogonResponseData struct {
	TrailerSeqNumbers     []int    `json:"trailer_seq_numbers"`
	TrailerUnladenWeights []int    `json:"trailer_unladen_weights"`
	Success               int      `json:"success"`
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
	LiftType                int        `json:"lift_type"`
	APMDirection            string     `json:"apm_direction"`
	DualCycle               string     `json:"dual_cycle"`
	Activity                int        `json:"activity"`
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
