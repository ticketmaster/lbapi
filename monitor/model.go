package monitor

// Monitor - implements monitor package logic.
type Monitor struct {
	// Avi - implements Avi logic.
	Avi *Avi
	// Netscaler - implements Netscaler logic.
	Netscaler *Netscaler
}

// Data - resource configuration.
type Data struct {
	// Name - friendly name of the resource. Typically inherited from parent.
	Name string `json:"name,omitempty"`
	// SendInterval - number of seconds between health checks.
	SendInterval int `json:"send_interval,omitempty"`
	// Description [optional] provides additional details for the resource.
	Description string `json:"description,omitempty"`
	// ReceiveTimeout - number of seconds before the health check times out.
	ReceiveTimeout int `json:"receive_timeout,omitempty"`
	// Type options include: http, http-ecv, https, https-ecv, tcp, tcp-ecv, ping, udp, and external.
	Type string `json:"type,omitempty"`
	// SuccessfulCount - number of times a positive response must be received to mark the resource healthy.
	SuccessfulCount int `json:"successful_count,omitempty"`
	// FailedCount - number of times the test must fail before the resource is marked down.
	FailedCount int `json:"failed_count,omitempty"`
	// Request [optional] - data the monitor will send.
	Request string `json:"request,omitempty"`
	// Response [optional] - expected response.
	Response string `json:"response,omitempty"`
	// MaintenanceResponse [optional] is an Avi only field. This will mark the resource as down due to maintenance if response data match.
	MaintenanceResponse string `json:"maintenance_response,omitempty"`
	// MaintenanceResponseCodes [optional] is an Avi only field. This will mark the resource as down due to maintenance if response codes match.
	MaintenanceResponseCodes []string `json:"maintenance_response_codes,omitempty"`
	// ResponseCodes [optional] - expected response codes.
	ResponseCodes []string `json:"response_codes,omitempty"`
	// MonitorPort [optional] defaults to same port of pool.
	MonitorPort int `json:"monitor_port,omitempty"`
	// SourceUUID [system] - record id of the resource.
	SourceUUID string `json:"_uuid,omitempty"`
}

// Collection - collection of all monitors.
type Collection struct {
	// Source - map of all related records indexed by uuid from source.
	Source map[string]Data
	// System - map of all related records indexed by name from source.
	System map[string]Data
}

// Monitors - collection of all supported monitors.
type Monitors struct {
	// System - map of all related records indexed by uuid from lbapi.
	System map[string]sourceData
	// Source - map of all related records indexed by uuid from source.
	Source map[string]sourceData
}

// sourceData - data derived from source.
type sourceData struct {
	// Ref - object reference.
	Ref string
	// Default - name of default resource.
	Default string
	// UUID - uuid of resource.
	UUID string
}
