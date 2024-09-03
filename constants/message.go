package constants

type VehicleParam struct {
	ID                 string
	VehicleID          string
	Activity           int64
	NextLocation       string
	NextLocationLane   int64
	LiftType           int64
	TargetDockPosition int
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
            "next_location": "{{.NextLocation}}          ",
            "next_location_lane": "{{.NextLocationLane}}",
            "assigned_cntr_size": "40",
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
