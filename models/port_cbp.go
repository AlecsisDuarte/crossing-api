package models

// Structured used by CBP to enlist their ports
type PortCBP struct {
	PortNumber             string `json:"port_number,omitempty"`
	Border                 string `json:"border,omitempty"`
	PortName               string `json:"port_name,omitempty"`
	CrossingName           string `json:"crossing_name,omitempty"`
	Hours                  string `json:"hours,omitempty"`
	Date                   string `json:"date,omitempty"`
	Time                   string `json:"time,omitempty"`
	PortStatus             string `json:"port_status,omitempty"`
	CommercialVehicleLanes struct {
		MaximumLanes        string `json:"maximum_lanes,omitempty"`
		CvAutomationType    string `json:"cv_automation_type,omitempty"`
		CvSegmentFrom       string `json:"cv_segment_from,omitempty"`
		CvSegmentTo         string `json:"cv_segment_to,omitempty"`
		CvStandardTolerance string `json:"cv_standard_tolerance,omitempty"`
		CvFastTolerance     string `json:"cv_fast_tolerance,omitempty"`
		StandardLanes       struct {
			UpdateTime        string `json:"update_time,omitempty"`
			OperationalStatus string `json:"operational_status,omitempty"`
			DelayMinutes      string `json:"delay_minutes,omitempty"`
			LanesOpen         string `json:"lanes_open,omitempty"`
		} `json:"standard_lanes,omitempty"`
		FASTLanes struct {
			UpdateTime        string `json:"update_time,omitempty"`
			OperationalStatus string `json:"operational_status,omitempty"`
			DelayMinutes      string `json:"delay_minutes,omitempty"`
			LanesOpen         string `json:"lanes_open,omitempty"`
		} `json:"FAST_lanes,omitempty"`
	} `json:"commercial_vehicle_lanes,omitempty"`
	PassengerVehicleLanes struct {
		MaximumLanes           string `json:"maximum_lanes,omitempty"`
		PvAutomationType       string `json:"pv_automation_type,omitempty"`
		PvSegmentFrom          string `json:"pv_segment_from,omitempty"`
		PvSegmentTo            string `json:"pv_segment_to,omitempty"`
		PvStandardTolerance    string `json:"pv_standard_tolerance,omitempty"`
		PvNexusSentriTolerance string `json:"pv_nexus_sentri_tolerance,omitempty"`
		PvReadyTolerance       string `json:"pv_ready_tolerance,omitempty"`
		StandardLanes          struct {
			UpdateTime        string `json:"update_time,omitempty"`
			OperationalStatus string `json:"operational_status,omitempty"`
			DelayMinutes      string `json:"delay_minutes,omitempty"`
			LanesOpen         string `json:"lanes_open,omitempty"`
		} `json:"standard_lanes,omitempty"`
		NEXUSSENTRILanes struct {
			UpdateTime        string `json:"update_time,omitempty"`
			OperationalStatus string `json:"operational_status,omitempty"`
			DelayMinutes      string `json:"delay_minutes,omitempty"`
			LanesOpen         string `json:"lanes_open,omitempty"`
		} `json:"NEXUS_SENTRI_lanes,omitempty"`
		ReadyLanes struct {
			UpdateTime        string `json:"update_time,omitempty"`
			OperationalStatus string `json:"operational_status,omitempty"`
			DelayMinutes      string `json:"delay_minutes,omitempty"`
			LanesOpen         string `json:"lanes_open,omitempty"`
		} `json:"ready_lanes,omitempty"`
	} `json:"passenger_vehicle_lanes,omitempty"`
	PedestrianLanes struct {
		MaximumLanes         string `json:"maximum_lanes,omitempty"`
		PedAutomationType    string `json:"ped_automation_type,omitempty"`
		PedSegmentFrom       string `json:"ped_segment_from,omitempty"`
		PedSegmentTo         string `json:"ped_segment_to,omitempty"`
		PedStandardTolerance string `json:"ped_standard_tolerance,omitempty"`
		PedReadyTolerance    string `json:"ped_ready_tolerance,omitempty"`
		StandardLanes        struct {
			UpdateTime        string `json:"update_time,omitempty"`
			OperationalStatus string `json:"operational_status,omitempty"`
			DelayMinutes      string `json:"delay_minutes,omitempty"`
			LanesOpen         string `json:"lanes_open,omitempty"`
		} `json:"standard_lanes,omitempty"`
		ReadyLanes struct {
			UpdateTime        string `json:"update_time,omitempty"`
			OperationalStatus string `json:"operational_status,omitempty"`
			DelayMinutes      string `json:"delay_minutes,omitempty"`
			LanesOpen         string `json:"lanes_open,omitempty"`
		} `json:"ready_lanes,omitempty"`
	} `json:"pedestrian_lanes,omitempty"`
	ConstructionNotice string `json:"construction_notice,omitempty"`
	Automation         string `json:"automation,omitempty"`
	AutomationEnabled  string `json:"automation_enabled,omitempty"`
}
