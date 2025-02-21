package area

import "encoding/json"

type ManualModeRequest struct {
	Ingress int64            `json:"ingress"`
	Egress  int64            `json:"egress"`
	QCLanes map[string]int64 `json:"qc_cfg"`
	Mode    int64            `json:"mode"`
}

func (req ManualModeRequest) String() string {
	v, _ := json.Marshal(req)
	return string(v)
}

type HatchCoverConfigRequest struct {
	Name  string `json:"name"`
	Op    string `json:"op"`
	Start int64  `json:"start"`
	End   int64  `json:"end"`
}

func (req HatchCoverConfigRequest) String() string {
	v, _ := json.Marshal(req)
	return string(v)
}
