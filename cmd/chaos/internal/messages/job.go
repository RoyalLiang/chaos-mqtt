package messages

type Task struct {
	APMID string      `json:"apm_id"`
	Data  interface{} `json:"data"`
}

type RouteResponseJobInstruction struct {
	APMID string    `json:"apm_id"`
	Data  RouteData `json:"data"`
}

type RouteData struct {
	Timestamp               int64      `json:"timestamp"`
	Success                 int64      `json:"success"`
	APMID                   string     `json:"apm_id"`
	APMDirection            string     `json:"apm_direction"`
	CntrSizes               []string   `json:"cntr_sizes"`
	IMOClass                []string   `json:"imo_class"`
	Activity                int        `json:"activity"`
	SourceLocations         []string   `json:"source_locations"`
	AssignedCntrType        string     `json:"assigned_cntr_type"`
	TargetDockPosition      string     `json:"target_dock_position"`
	OperationalGroups       []string   `json:"operational_groups"`
	CntrStatus              []string   `json:"cntr_status"`
	OperationalQCSequences  []string   `json:"operational_qc_sequences"`
	Urgents                 []string   `json:"urgents"`
	CntrLocationsOnAPM      []string   `json:"cntr_locations_on_apm"`
	DGs                     []string   `json:"dgs"`
	DualCycle               string     `json:"dual_cycle"`
	DGGroups                []string   `json:"dg_groups"`
	NextLocation            string     `json:"next_location"`
	RouteMandate            string     `json:"route_mandate"`
	RouteType               string     `json:"route_type"`
	RouteDAG                []RouteDag `json:"route_dag"`
	Cones                   []int64    `json:"cones"`
	NextLocationLane        string     `json:"next_location_lane"`
	ID                      string     `json:"id"`
	PlugRequireds           []string   `json:"plug_requireds"`
	MapVersion              string     `json:"map_version"`
	NumMountedCntr          int        `json:"num_mounted_cntr"`
	TrailerPositions        []string   `json:"trailer_positions"`
	CntrCategorys           []string   `json:"cntr_categorys"`
	OperationalJobSequences []string   `json:"operational_job_sequences"`
	OperationalTypes        []string   `json:"operational_types"`
	JobTypes                []string   `json:"job_types"`
	DestLocations           []string   `json:"dest_locations"`
	ReferTemperatures       []string   `json:"refer_temperatures"`
	Message                 string     `json:"message"`
	AssignedCntrSize        string     `json:"assigned_cntr_size"`
	WeightClass             []string   `json:"weight_class"`
	MotorDirections         []string   `json:"motor_directions"`
	LiftType                int        `json:"lift_type"`
	CntrTypes               []string   `json:"cntr_types"`
	CntrWeights             []string   `json:"cntr_weights"`
	OperationalGroup        string     `json:"operational_group"`
	CntrNumbers             []string   `json:"cntr_numbers"`
	OffloadSequences        []string   `json:"offload_sequences"`
	RejectionCode           string     `json:"rejection_code"`
}

type JobInstructionResponse struct {
	APMID string                     `json:"apm_id"`
	Data  JobInstructionResponseData `json:"data"`
}

type JobInstructionResponseData struct {
	Timestamp            int64      `json:"timestamp"`
	ID                   string     `json:"id"`
	Success              int        `json:"success"`
	RejectionCode        string     `json:"rejection_code"`
	CurrentRoute         []Position `json:"current_route"`
	AffectedByRoadBlocks int        `json:"affected_by_road_blocks"`
	DestLocation         string     `json:"dest_location"`
	DestinationName      string     `json:"destination_name"`
}

type Position struct {
	Convention int       `json:"convention"`
	Position   []float64 `json:"position"`
	Heading    float64   `json:"heading"`
}

type CallInResponse struct {
	APMID string             `json:"apm_id"`
	Data  CallInResponseData `json:"data"`
}

type CallInResponseData struct {
	Timestamp     int64  `json:"timestamp"`
	ID            string `json:"id"`
	Success       int    `json:"success"`
	RejectionCode string `json:"rejection_code"`
}

type ReadyForMoveToQC struct {
	APMID string               `json:"apm_id"`
	Data  ReadyForMoveToQCData `json:"data"`
}

type ReadyForMoveToQCData struct {
	Timestamp           int64      `json:"timestamp"`
	ContainerSize       string     `json:"container_size"`
	DestinationWaypoint string     `json:"destination_waypoint"`
	RouteDAG            []RouteDag `json:"route_dag"`
	RouteType           string     `json:"route_type"`
	DestinationLane     string     `json:"destination_lane"`
	DestinationName     string     `json:"destination_name"`
	ID                  string     `json:"id"`
	TargetDockPosition  string     `json:"target_dock_position"`
}

type ReadyForIngressToCallIn struct {
	APMID string                      `json:"apm_id"`
	Data  ReadyForIngressToCallInData `json:"data"`
}

type ReadyForIngressToCallInData struct {
	Timestamp           int64      `json:"timestamp"`
	ID                  string     `json:"id"`
	DestinationWaypoint string     `json:"destination_waypoint"`
	IngressBS           string     `json:"ingress_b_s"`
	IngressSB           string     `json:"ingress_s_b"`
	PMDirection         string     `json:"pm_direction"`
	RouteType           string     `json:"route_type"`
	RouteDAG            []RouteDag `json:"route_dag"`
	DestinationLane     string     `json:"destination_lane"`
	WharfStretchID      string     `json:"wharf_stretch_id"`
	Message             string     `json:"message"`
	LaneAvailability    []string   `json:"lane_availability"`
	DestinationName     string     `json:"destination_name"`
}

type DockPositionResponse struct {
	APMID string                   `json:"apm_id"`
	Data  DockPositionResponseData `json:"data"`
}

type DockPositionResponseData struct {
	Timestamp     int64  `json:"timestamp"`
	ID            string `json:"id"`
	Success       int    `json:"success"`
	RejectionCode string `json:"rejection_code"`
}

type PathUpdateAvailable struct {
	APMID string                  `json:"apm_id"`
	Data  PathUpdateAvailableData `json:"data"`
}

type PathUpdateAvailableData struct {
	Timestamp           int64      `json:"timestamp"`
	ID                  string     `json:"id"`
	DestinationWaypoint string     `json:"destination_waypoint"`
	TargetDockPosition  string     `json:"target_dock_position"`
	RouteType           string     `json:"route_type"`
	DestinationLane     string     `json:"destination_lane"`
	DestinationName     string     `json:"destination_name"`
	RouteDag            []RouteDag `json:"route_dag"`
}

type MountInstructionResponse struct {
	APMID string                       `json:"apm_id"`
	Data  MountInstructionResponseData `json:"data"`
}

type MountInstructionResponseData struct {
	Timestamp     int64  `json:"timestamp"`
	ID            string `json:"id"`
	Success       int    `json:"success"`
	RejectionCode string `json:"rejection_code"`
}

type OffloadInstructionResponse struct {
	APMID string                         `json:"apm_id"`
	Data  OffloadInstructionResponseData `json:"data"`
}

type OffloadInstructionResponseData struct {
	Timestamp     int64  `json:"timestamp"`
	ID            string `json:"id"`
	Success       int    `json:"success"`
	RejectionCode string `json:"rejection_code"`
}

type WharfDockPositionRequest struct {
	APMID string                       `json:"apm_id"`
	Data  WharfDockPositionRequestData `json:"data"`
}

type WharfDockPositionRequestData struct {
	Timestamp          int64  `json:"timestamp"`
	ID                 string `json:"id"`
	TargetDockPosition string `json:"target_dock_position"`
	DestinationLane    string `json:"destination_lane"`
	DestinationName    string `json:"destination_name"`
}
