package model

// monitorname	<String>	Read-write	The name of the monitor.<br>Minimum length = 1
// lbmonbindings_servicegroup_binding	<lbmonbindings_servicegroup_binding[]>	Read-only	servicegroup that can be bound to lbmonbindings.
// lbmonbindings_service_binding	<lbmonbindings_service_binding[]>	Read-only	service that can be bound to lbmonbindings.

// LbmonbindingsBindingWrapper wraps the object and serves as default response.
type LbmonbindingsBindingWrapper struct {
	Errorcode            int                    `json:"errorcode,omitempty"`
	Message              string                 `json:"message,omitempty"`
	Severity             string                 `json:"severity,omitempty"`
	LbmonbindingsBinding []LbmonbindingsBinding `json:"lbmonbindings_binding,omitempty"`
}

// LbmonbindingsBinding describes the object.
type LbmonbindingsBinding struct {
	Monitorname                      string                             `json:"monitorname,omitempty"`
	LbmonbindingsServicegroupBinding []LbmonbindingsServicegroupBinding `json:"lbmonbindings_servicegroup_binding,omitempty"`
	LbmonbindingsServiceBinding      []LbmonbindingsServiceBinding      `json:"lbmonbindings_service_binding,omitempty"`
}
