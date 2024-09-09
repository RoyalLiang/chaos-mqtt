package messages

import "encoding/json"

type SwitchModeResponse struct {
	APMID string                 `json:"apm_id"`
	Data  SwitchModeResponseData `json:"data"`
}

type SwitchModeResponseData struct {
	Timestamp     int64  `json:"timestamp"`
	ID            string `json:"id"`
	Success       int    `json:"success"`
	RejectionCode string `json:"rejection_code"`
	SetMode       string `json:"set_mode"`
}

type LogonRequest struct {
	APMID string           `json:"apm_id"`
	Data  LogonRequestData `json:"data"`
}

type LogonRequestData struct {
	Timestamp         int64    `json:"timestamp"`
	NumTrailers       int      `json:"num_trailers"`
	TrailerNumbers    []string `json:"trailer_numbers"`
	TrailerSeqNumbers []int    `json:"trailer_seq_numbers"`
}

func (lr LogonRequest) String() string {
	v, _ := json.Marshal(lr)
	return string(v)
}

type APMArrivedRequest struct {
	APMID string                `json:"apm_id"`
	Data  APMArrivedRequestData `json:"data"`
}

type APMArrivedRequestData struct {
	Timestamp          int64  `json:"timestamp"`
	AlternateLaneDock  string `json:"alternate_lane_dock"`
	Location           string `json:"location"`
	ID                 string `json:"id"`
	TargetDockPosition string `json:"target_dock_position"`
}
