package constants

type JobParam struct {
	ID                 string
	VehicleID          string
	Activity           int64
	Destination        string
	Lane               int64
	LiftType           int64
	TargetDockPosition int64
	ContainerSize      string
}

const RouteRequestJobInstruction = `
{
        "apm_id": "{{.VehicleID}}",
        "data": {
            "apm_id": "{{.VehicleID}}",
            "id": "{{.ID}}",
            "timestamp": 1721374079082,
            "map_version": "PPT-456-20220620-20240112",
            "route_dag": [],
            "route_mandate": "N",
            "operational_group": "IYO50OO01",
            "lift_type": {{.LiftType}},
            "apm_direction": "S",
            "dual_cycle": "N",
            "activity": {{.Activity}},
            "next_location": "{{.Destination}}",
            "next_location_lane": "{{.Lane}}",
            "assigned_cntr_size": "{{.ContainerSize}}",
            "assigned_cntr_type": "GP",
            "num_mounted_cntr": 0,
            "target_dock_position": "{{.TargetDockPosition}}",
            "operational_types": [],
            "job_types": [],
            "operational_groups": [],
            "cntr_numbers": [],
            "cntr_sizes": [],
            "cntr_weights": [],
            "cntr_categorys": [],
            "cntr_types": [],
            "cntr_status": [],
            "dg_groups": [],
            "imo_class": [],
            "refer_temperatures": [],
            "cntr_locations_on_apm": [],
            "source_locations": [],
            "dest_locations": [],
            "offload_sequences": [],
            "cones": [],
            "operational_qc_sequences": [],
            "operational_job_sequences": [],
            "trailer_positions": [],
            "weight_class": [],
            "urgents": [],
            "dgs": [],
            "plug_requireds": [],
            "motor_directions": []
        }
    }
`

type JobInstructionParam struct {
}

const JobInstruction = `
{
        "apm_id": "{{.VehicleID}}",
        "data": {
            "apm_id": "{{.VehicleID}}",
            "id": "{{.ID}}",
            "timestamp": 1721353033201,
            "map_version": "PPT-456-20220620-20240112",
            "route_dag": [],
            "route_mandate": "N",
            "operational_group": "IYO50OO01",
            "lift_type": {{.LiftType}},
            "apm_direction": "B",
            "dual_cycle": "N",
            "activity": {{.Activity}},
            "next_location": "{{.Destination}}",
            "next_location_lane": "{{.Lane}}",
            "assigned_cntr_size": "{{.ContainerSize}}",
            "assigned_cntr_type": "GP",
            "num_mounted_cntr": 0,
            "target_dock_position": "{{.TargetDockPosition}}",
            "operational_types": [],
            "job_types": [],
            "operational_groups": [],
            "cntr_numbers": [],
            "cntr_sizes": [],
            "cntr_weights": [],
            "cntr_categorys": [],
            "cntr_types": [],
            "cntr_status": [],
            "dg_groups": [],
            "imo_class": [],
            "refer_temperatures": [],
            "cntr_locations_on_apm": [],
            "source_locations": [],
            "dest_locations": [],
            "offload_sequences": [],
            "cones": [],
            "operational_qc_sequences": [],
            "operational_job_sequences": [],
            "trailer_positions": [],
            "weight_class": [],
            "urgents": [],
            "dgs": [],
            "plug_requireds": [],
            "motor_directions": [],
            "route_type": "G"
        }
    }
`

type SwitchModeParam struct {
	VehicleID string
	Mode      string
}

const SwitchMode = `
{
        "apm_id": "{{.VehicleID}}",
        "data": {
            "id": "dab5cc52-a035-4f66-9879-bdebe26201a0",
            "set_mode": "{{.Mode}}",
            "timestamp": 1718848612000
        }
    }
`

type CallInRequestParam struct {
	VehicleID string
	Mode      int64
}

const CallInRequest = `
{
        "apm_id": "{{.VehicleID}}",
        "data": {
            "id": "6898bb69-4618-4855-b105-df38af45620f",
            "qc_number": "PQC921   ",
            "call_in_mode": {{.Mode}}
        }
    }
`

type DockPositionParam struct {
	VehicleID   string
	Activity    int64
	ConLocation int64
}

const DockPosition = `
{
        "apm_id": "{{.VehicleID}}",
        "data": {
            "activity": {{.Activity}},
            "cntr_location_on_apm": {{.ConLocation}},
            "chassis_lane": 11,
            "block": "TB03",
            "id": "MOAPM833020122023110519",
            "terminal": "V",
            "slot": 32,
            "operational_group": "IYO50OO01 ",
            "message": "YOS OYOS-YCPMJOB0-H 20230307123142945013ADDP APM8330 VTB03032IYO50OO01 1131 ",
            "timestamp": 1703042038879
        }
    }
`

type IngressToCallInParam struct {
	VehicleID string
}

const IngressToCallIn = `
{
        "apm_id": "{{.VehicleID}}",
        "data": {
            "destination_waypoint": "ca_wid_24_wm_1460_lane_3",
            "ingress_b_s": "",
            "ingress_s_b": "956",
            "pm_direction": "S",
            "route_type": "G",
            "route_dag": [],
            "destination_lane": "3",
            "lane_availability": [],
            "destination_name": "P,PQC913           _Call In Area (S-B) Lane 3",
            "id": "WFMOAPM860113032024111845",
            "wharf_stretch_id": "24",
            "message": "VOS OWTMS-TRAMNG-H 20240313112321850373APGQ PQC913 APM8601 0 ",
            "timestamp": 1710300202492
        }
    }
`

type MoveToQCParam struct {
	VehicleID string
}

const MoveToQC = `
{
        "apm_id": "{{.VehicleID}}",
        "data": {
            "id": "WFMOAPM900111072024101004",
            "timestamp": 1720664438565,
            "destination_name": "P,QC921",
            "target_dock_position": "",
            "container_size": "",
            "destination_lane": "3",
            "destination_waypoint": "ca_wid_24_wm_1429_lane_6",
            "route_dag": [],
            "route_type": "G"
        }
    }
`

type VesselBerthParam struct {
	VesselID  string
	Direction string
	StartPos  int64
	EndPos    int64
	Cranes    string
}

const VesselBerth = `
{
        "id": "{{.VesselID}}",
        "vessel_name": "ONE HARBOR",
        "wharf_side_indicator": "{{.Direction}}",
        "terminal": "P",
        "berth": "29",
        "position_from": {{.StartPos}},
        "position_to": {{.EndPos}},
        "assigned_qc": "{{.Cranes}}",
        "timestamp": 12345678901123
    }
`

type VesselUnberthParam struct {
	VesselID string
}

const VesselUnberth = `
{"id":"{{.VesselID}}","timestamp":12345678901123}
`
