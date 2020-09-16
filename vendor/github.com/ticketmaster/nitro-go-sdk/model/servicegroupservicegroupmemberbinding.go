package model

// Customserverid. The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.<br>Default value: "None".
// Delay. Time, in seconds, allocated for a shutdown of the services in the service group. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).
// Graceful. Wait for all existing connections to the service to terminate before shutting down the service.<br>Default value: NO<br>Possible values = YES, NO.
// Hashid. The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.<br>Minimum value = 1.
// Ip. IP Address.
// Port. Server port number.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Serverid. The identifier for the service. This is used when the persistency type is set to Custom Server ID.
// Servername. Name of the server to which to bind the service group.<br>Minimum length = 1.
// Servicegroupname. Name of the service group.<br>Minimum length = 1.
// State. Initial state of the service group.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Statechangetimesec. Time when last state change occurred. Seconds part.
// Svrstate. The state of the service.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
// Tickssincelaststatechange. Time in 10 millisecond ticks since the last state change.
// Weight. Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.<br>Minimum value = 1<br>Maximum value = 100.

// ServicegroupServicegroupmemberBindingAdd defines add request.
type ServicegroupServicegroupmemberBindingAdd struct {
	ServicegroupServicegroupmemberBinding ServicegroupServicegroupmemberBindingAddBody `json:"servicegroup_servicegroupmember_binding,omitempty"`
}

// ServicegroupServicegroupmemberBindingAddBody body for adding object.
type ServicegroupServicegroupmemberBindingAddBody struct {
	Servicegroupname string `json:"servicegroupname"`
	IP               string `json:"ip,omitempty"`
	Servername       string `json:"servername,omitempty"`
	Port             int    `json:"port"`
	Weight           string `json:"weight,omitempty"`
	Customserverid   string `json:"customserverid,omitempty"`
	Serverid         string `json:"serverid,omitempty"`
	State            string `json:"state"`
	Hashid           string `json:"hashid,omitempty"`
}

// ServicegroupServicegroupmemberBindingWrapper wraps the object and serves as default response.
type ServicegroupServicegroupmemberBindingWrapper struct {
	Errorcode                             int                                     `json:"errorcode,omitempty"`
	Message                               string                                  `json:"message,omitempty"`
	Severity                              string                                  `json:"severity,omitempty"`
	ServicegroupServicegroupmemberBinding []ServicegroupServicegroupmemberBinding `json:"servicegroup_servicegroupmember_binding,omitempty"`
}

// ServicegroupServicegroupmemberBinding describes the object.
type ServicegroupServicegroupmemberBinding struct {
	Servicegroupname          string `json:"servicegroupname"`
	IP                        string `json:"ip"`
	Port                      int    `json:"port"`
	State                     string `json:"state"`
	Hashid                    string `json:"hashid"`
	Serverid                  string `json:"serverid"`
	Servername                string `json:"servername"`
	Customserverid            string `json:"customserverid"`
	Weight                    string `json:"weight"`
	Delay                     int    `json:"delay"`
	Statechangetimesec        string `json:"statechangetimesec"`
	Svrstate                  string `json:"svrstate"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange"`
	Graceful                  string `json:"graceful"`
}
