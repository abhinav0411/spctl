package models

type Song struct {
	ContextURI string `json:"context_uri"`
	OffSet     struct {
		Position   int `json:"position"`
		PositionMs int `json:"position_ms"`
	} `json:"offset"`
}
