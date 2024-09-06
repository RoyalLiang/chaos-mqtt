package messages

type RouteDag struct {
	Name      string  `json:"name"`
	Direction string  `json:"direction"`
	Type      string  `json:"type"`
	LaneNum   string  `json:"lane_num"`
	BlockNum  string  `json:"block_num"`
	SlotNum   string  `json:"slot_num"`
	Dev       float64 `json:"dev"`
	Pose      Pose    `json:"pose"`
}

type Pose struct {
	Convention int       `json:"convention"`
	Position   []float64 `json:"position"`
	Heading    float64   `json:"heading"`
}
