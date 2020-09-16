package model

// Name. Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or my vserver). .<br>Minimum length = 1.
// Servicegroupname. The service group name bound to the selected load balancing virtual server.
// Servicename. Service to bind to the virtual server.<br>Minimum length = 1.
// Weight. Integer specifying the weight of the service. A larger number specifies a greater weight. Defines the capacity of the service relative to the other services in the load balancing configuration. Determines the priority given to the service in load balancing decisions.<br>Default value: 1<br>Minimum value = 1<br>Maximum value = 100.

// LbvserverServicegroupBindingAdd defines add request
type LbvserverServicegroupBindingAdd struct {
	LbvserverServicegroupBinding LbvserverServicegroupBinding `json:"lbvserver_servicegroup_binding"`
}

// LbvserverServicegroupBindingWrapper wraps the object and serves as default response.
type LbvserverServicegroupBindingWrapper struct {
	Errorcode                    int                            `json:"errorcode,omitempty"`
	Message                      string                         `json:"message,omitempty"`
	Severity                     string                         `json:"severity,omitempty"`
	LbvserverServicegroupBinding []LbvserverServicegroupBinding `json:"lbvserver_servicegroup_binding"`
}

// LbvserverServicegroupBinding describes the object.
type LbvserverServicegroupBinding struct {
	Name             string `json:"name"`
	Servicegroupname string `json:"servicegroupname"`
	Servicename      string `json:"servicename,omitempty"`
	Weight           string `json:"weight,omitempty"`
}
