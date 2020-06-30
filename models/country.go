package models

// Country information
type Country struct {
	ID       string   `json:"id"`
	Name     string   `json:"name,omitempty"`
	Currency string   `json:"currency,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Exchange Exchange `json:"exchange,omitempty"`
	States   []State  `json:"states,omitempty"`
}
