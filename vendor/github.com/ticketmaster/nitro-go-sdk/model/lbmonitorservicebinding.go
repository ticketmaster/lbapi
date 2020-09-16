package model

// dup_state	<String>	Read-write	State of the monitor. The state setting for a monitor of a given type affects all monitors of that type. For example, if an HTTP monitor is enabled, all HTTP monitors on the appliance are (or remain) enabled. If an HTTP monitor is disabled, all HTTP monitors on the appliance are disabled.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// dup_weight	<Double>	Read-write	Weight to assign to the binding between the monitor and service.<br>Default value: 1<br>Minimum value = 1<br>Maximum value = 100
// monitorname	<String>	Read-write	Name of the monitor.<br>Minimum length = 1
// Operations
// servicegroupname	<String>	Read-write	Name of the service group.<br>Minimum length = 1
// servicename	<String>	Read-write	Name of the service or service group.<br>Minimum length = 1
// state	<String>	Read-write	State of the monitor. The state setting for a monitor of a given type affects all monitors of that type. For example, if an HTTP monitor is enabled, all HTTP monitors on the appliance are (or remain) enabled. If an HTTP monitor is disabled, all HTTP monitors on the appliance are disabled.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// weight	<Double>	Read-write	Weight to assign to the binding between the monitor and service.<br>Minimum value = 1<br>Maximum value = 100

// LbmonitorServiceBindingWrapper wraps the object and serves as default response.
type LbmonitorServiceBindingWrapper struct {
	Errorcode               int                       `json:"errorcode,omitempty"`
	Message                 string                    `json:"message,omitempty"`
	Severity                string                    `json:"severity,omitempty"`
	LbmonitorServiceBinding []LbmonitorServiceBinding `json:"lbmonitor_service_binding,omitempty"`
}

// LbmonitorServiceBinding describes the object.
type LbmonitorServiceBinding struct {
	Servicename      string  `json:"servicename,omitempty"`
	Servicegroupname string  `json:"servicegroupname,omitempty"`
	Monitorname      string  `json:"monitorname,omitempty"`
	DupState         string  `json:"dup_state,omitempty"`
	DupWeight        float64 `json:"dup_weight,omitempty"`
	State            string  `json:"state,omitempty"`
	Weight           float64 `json:"weight,omitempty"`
}

// LbmonitorServiceBindingAdd describes the object.
type LbmonitorServiceBindingAdd struct {
	LbmonitorServiceBinding LbmonitorServiceBinding `json:"lbmonitor_service_binding,omitempty"`
}
