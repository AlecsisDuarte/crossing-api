package models

// County information
type County struct {
	Name  string   `json:"name,omitempty"`
	Ports []string `json:"ports,omitempty"`
}
