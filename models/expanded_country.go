package models

// ExpandedCountry contains all the surrounding countries plus all the CBP ports
type ExpandedCountry struct {
	Countries *[]Country
	Ports     *[]PortCBP
}
