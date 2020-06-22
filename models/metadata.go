package models

// Metadata stores extra information to enhance information suplied by CBP
type Metadata struct {
	GeographicInfo GeographicInfo `json:"geographic_info,omitempty"`
}
