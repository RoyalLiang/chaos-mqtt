package constants

type (
	TopicFromFMS string
)

var TopicFromAVCS = map[string]string{
	"LOGON_RESPONSE":                        "logon_response",
	"CANCEL_JOB":                            "cancel_job",
	"SWITCH_MODE":                           "switch_mode",
	"MAP_WGS84":                             "/map/wgs84/",
	"BLOCK":                                 "block",
	"MAP_BLOCK":                             "/map/blocks/",
	"PARKING_ACTION":                        "parking_action",
	"MAINTENANCE":                           "maintenance",
	"PARKING":                               "parking",
	"RESUME":                                "resume",
	"STOP":                                  "stop",
	"ROUTE_REQUEST_JOB_INSTRUCTION":         "route_request_job_instruction",
	"ROUTE_REQUEST":                         "route_request",
	"JOB_INSTRUCTION":                       "job_instruction",
	"DOCK_POSITION":                         "dock_position",
	"MOUNT_INSTRUCTION":                     "mount_instruction",
	"OFFLOAD_INSTRUCTION":                   "offload_instruction",
	"WHARF_DOCK_POSITION_RESPONSE":          "wharf_dock_position_response",
	"INGRESS_TO_CALL_IN":                    "ingress_to_call_in",
	"MOVE_TO_QC":                            "move_to_qc",
	"ARMG_REQUEST":                          "armg_request",
	"VESSEL_BERTH":                          "vessel_berth",
	"VESSEL_UNBERTH":                        "vessel_unberth",
	"DERIVED_VESSEL_CONFIGURATION_RESPONSE": "derived_vessel_configuration_response",
	"CALL_IN_STATUS":                        "call_in_status",
	"HATCH_COVER_OPS":                       "hatch_cover_ops",
	"QC_POSITION_INFO_RESPONSE":             "qc_position_info_response",
	"PATH_UPDATE_AVAILABLE_RESPONSE":        "path_update_available_response",
	"PATH_UPDATE_REQUEST":                   "path_update_request",
	"READY_FOR_INGRESS_TO_CALL_IN_RESPONSE": "ready_for_ingress_to_call_in_response",
	"READY_FOR_MOVE_TO_QC_RESPONSE":         "ready_for_move_to_qc_response",
	"READY_FOR_INGRESS_TO_QC_RESPONSE":      "ready_for_ingress_to_qc_response",
	"INGRESS_TO_QC":                         "ingress_to_qc",
	"APM_HEARTBEAT":                         "/apm/heartbeat",
	"CALL_IN_REQUEST":                       "call_in_request",
	"CONING_DECONING_COMPLETION":            "coning_deconing_completion",
	"PM_ACTIVITY_INFO":                      "pm_activity_info",
	"PM_NAVIGATION_INFO":                    "pm_navigation_info",
	"MANUAL_EXCEPTION_HANDLING":             "manual_exception_handling",
	"APM_ARRIVED_REQUEST":                   "/apm/apm_arrived_request",
	"APM_ACCEPTANCE_UPDATE":                 "apm_acceptance_update",
}

const (
	POWER_ON_REQUEST                    TopicFromFMS = "power_on_request"
	HEARTBEAT                           TopicFromFMS = "heartbeat"
	UPDATE_TRAILER                      TopicFromFMS = "update_trailer"
	LOGOFF_REQUEST                      TopicFromFMS = "logoff_request"
	POWER_OFF_REQUEST                   TopicFromFMS = "power_off_request"
	REQUEST_JOB                         TopicFromFMS = "request_job"
	CANCEL_JOB_RESPONSE                 TopicFromFMS = "cancel_job_response"
	SWITCH_MODE_RESPONSE                TopicFromFMS = "switch_mode_response"
	MODE_CHANGE_UPDATE                  TopicFromFMS = "mode_change_update"
	BLOCK_RESPONSE                      TopicFromFMS = "block_response"
	BLOCKS_RESPONSE                     TopicFromFMS = "blocks_response"
	PARKING_STATE                       TopicFromFMS = "parking_state"
	MAINTENANCE_RESPONSE                TopicFromFMS = "maintenance_response"
	PARKING_RESPONSE                    TopicFromFMS = "parking_response"
	RESUME_RESPONSE                     TopicFromFMS = "resume_response"
	STOP_RESPONSE                       TopicFromFMS = "stop_response"
	ROUTE_RESPONSE_JOB_INSTRUCTION      TopicFromFMS = "route_response_job_instruction"
	ROUTE_RESPONSE                      TopicFromFMS = "route_response"
	JOB_INSTRUCTION_RESPONSE            TopicFromFMS = "job_instruction_response"
	APM_MOVE_REQUEST                    TopicFromFMS = "apm_move_request"
	QTRUCK_ARRIVED_REQUEST              TopicFromFMS = "apm_arrived_request"
	DOCK_POSITION_RESPONSE              TopicFromFMS = "dock_position_response"
	MOUNT_INSTRUCTION_RESPONSE          TopicFromFMS = "mount_instruction_response"
	OFFLOAD_INSTRUCTION_RESPONSE        TopicFromFMS = "offload_instruction_response"
	WHARF_DOCK_POSITION                 TopicFromFMS = "wharf_dock_position"
	INTERMEDIATE_LOCATION               TopicFromFMS = "intermediate_location"
	INGRESS_READY_RESPONSE              TopicFromFMS = "ingress_ready_response"
	INGRESS_TO_CALL_IN_RESPONSE         TopicFromFMS = "ingress_to_call_in_response"
	MOVE_TO_QC_RESPONSE                 TopicFromFMS = "move_to_qc_response"
	ARMG_REQUEST_RESPONSE               TopicFromFMS = "armg_request_response"
	VESSEL_BERTH_RESPONSE               TopicFromFMS = "vessel_berth_response"
	VESSEL_UNBERTH_RESPONSE             TopicFromFMS = "vessel_unberth_response"
	DERIVED_VESSEL_CONFIGURATION        TopicFromFMS = "derived_vessel_configuration"
	CALL_IN_STATUS_RESPONSE             TopicFromFMS = "call_in_status_response"
	HATCH_COVER_OPS_RESPONSE            TopicFromFMS = "hatch_cover_ops_response"
	QC_POSITION_INFO                    TopicFromFMS = "qc_position_info"
	PATH_UPDATE_AVAILABLE               TopicFromFMS = "path_update_available"
	PATH_UPDATE_RESPONSE                TopicFromFMS = "path_update_response"
	READY_FOR_INGRESS_TO_CALL_IN        TopicFromFMS = "ready_for_ingress_to_call_in"
	READY_FOR_MOVE_TO_QC                TopicFromFMS = "ready_for_move_to_qc"
	READY_FOR_INGRESS_TO_QC             TopicFromFMS = "ready_for_ingress_to_qc"
	INGRESS_TO_QC_RESPONSE              TopicFromFMS = "ingress_to_qc_response"
	CALL_IN_RESPONSE                    TopicFromFMS = "call_in_response"
	CONING_DECONING_COMPLETION_RESPONSE TopicFromFMS = "coning_deconing_completion_response"
	PM_ACTIVITY_INFO_RESPONSE           TopicFromFMS = "pm_activity_info_response"
	PM_NAVIGATION_INFO_RESPONSE         TopicFromFMS = "pm_navigation_info_response"
	MANUAL_EXCEPTION_HANDLING_RESPONSE  TopicFromFMS = "manual_exception_handling_response"
	APM_ARRIVED_RESPONSE                TopicFromFMS = "apm_arrived_response"
	APM_ACCEPTANCE_UPDATE_RESPONSE      TopicFromFMS = "apm_acceptance_update_response"
)
