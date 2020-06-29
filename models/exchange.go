package models

// Exchange is a wrapper structure for data fetched from https://api.exchangeratesapi.io
type Exchange struct {
	Rates       map[string]float64 `json:"rates"`
	Base        string             `json:"base,omitempty"`
	Date        string             `json:"date,omitempty"`
	Description string             `json:"description,omitempty"`
	Source      string             `json:"Source,omitempty"`
}
