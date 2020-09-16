package model

// Name. Name of the server for which to display parameters.<br>Minimum length = 1.
// Server_Gslbservice_Binding. gslbservice that can be bound to server.
// Server_Service_Binding. service that can be bound to server.
// Server_Servicegroup_Binding. servicegroup that can be bound to server.

// ServerBindingWrapper wraps the object and serves as default response.
type ServerBindingWrapper struct {
	Errorcode     int             `json:"errorcode"`
	Message       string          `json:"message"`
	Severity      string          `json:"severity"`
	ServerBinding []ServerBinding `json:"server_binding"`
}

// ServerBinding describes the resource.
type ServerBinding struct {
	Name                      string                      `json:"name"`
	ServerServicegroupBinding []ServerServicegroupBinding `json:"server_servicegroup_binding"`
	ServerServiceBinding      []ServerServiceBinding      `json:"server_service_binding"`
	ServerGslbserviceBinding  []interface{}               `json:"server_gslbservice_binding"`
}
