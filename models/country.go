package models

// Country information
type Country struct {
	ID           string   `json:"id"`
	Name         string   `json:"name,omitempty"`
	Currency     string   `json:"currency,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	ExchangeRate float64  `json:"exchange_rate,omitempty"`
	States       []State  `json:"states,omitempty"`
}
