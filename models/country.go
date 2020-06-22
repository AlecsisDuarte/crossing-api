package models

// Country information
type Country struct {
	ID       string   `json:"id"`
	Name     string   `json:"name,omitempty"`
	Currency string   `json:"currency,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Exchange float64  `json:"exchange,omitempty"`
}
