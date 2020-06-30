package models

// GeographicInfo stores geographic information about the ports such as the city, state and county of a
// a port in the opposite site of the entry
type GeographicInfo struct {
	Countries []Country           `json:"countries,omitempty"`
	States    map[string][]State  `json:"states,omitempty"`
	Counties  map[string][]County `json:"counties,omitempty"`
	Exchange  Exchange            `json:"exchange,omitempty"`
}
