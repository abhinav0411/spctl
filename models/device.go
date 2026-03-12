package models

type PlayerDevice struct {
	ID         string `json:"id"`
	Active     bool   `json:"is_active"`
	Restricted bool   `json:"is_restricted"`
	Name       string `json:"name"`
	Type       string `json:"type"`
}
