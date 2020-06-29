package models

// State holds infrmation of the stat
type State struct {
	ID       string   `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Counties []County `json:"counties,omitempty"`
}
