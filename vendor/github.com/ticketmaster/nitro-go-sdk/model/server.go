package model

// Appflowlog. Enable logging of AppFlow information for the specified service group.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Autoscale. Auto scale option for a servicegroup.<br>Default value: DISABLED<br>Possible values = DISABLED, DNS, POLICY.
// Boundtd. Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.<br>Minimum value = 0<br>Maximum value = 4094.
// Cacheable. Use the transparent cache redirection virtual server to forward the request to the cache server.<br>Default value: NO<br>Possible values = YES, NO.
// Cip. Before forwarding a request to the service, insert an HTTP header with the clients IPv4 or IPv6 address as its value. Used if the server needs the clients IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option.<br>Possible values = ENABLED, DISABLED.
// Cipheader. Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as clients IP header name.<br>Minimum length = 1.
// Cka. Enable client keep-alive for the service group.<br>Possible values = YES, NO.
// Clttimeout. Time, in seconds, after which to terminate an idle client connection.<br>Minimum value = 0<br>Maximum value = 31536000.
// Cmp. Enable compression for the specified service.<br>Possible values = YES, NO.
// Comment. Any information about the server.
// Customserverid. A positive integer to identify the service. Used when the persistency type is set to Custom Server ID.<br>Default value: "None".
// Delay. Time, in seconds, after which all the services configured on the server are disabled.
// Domain. Domain name of the server. For a domain based configuration, you must create the server first.<br>Minimum length = 1.
// Domainresolvenow. Immediately send a DNS query to resolve the servers domain name.
// Domainresolveretry. Time, in seconds, for which the NetScaler appliance must wait, after DNS resolution fails, before sending the next DNS query to resolve the domain name.<br>Default value: 5<br>Minimum value = 5<br>Maximum value = 20939.
// Downstateflush. Perform delayed clean-up of connections to all services in the service group.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Graceful. Shut down gracefully, without accepting any new connections, and disabling each service when all of its connections are closed.<br>Default value: NO<br>Possible values = YES, NO.
// Internal. Display names of the servers that have been created for internal use.
// Ipaddress. IPv4 or IPv6 address of the server. If you create an IP address based server, you can specify the name of the server, instead of its IP address, when creating a service. Note: If you do not create a server entry, the server IP address that you enter when you create a service becomes the name of the server.
// Ipv6Address. Support IPv6 addressing mode. If you configure a server with the IPv6 addressing mode, you cannot use the server in the IPv4 addressing mode.<br>Default value: NO<br>Possible values = YES, NO.
// Maxbandwidth. Maximum bandwidth, in Kbps, allocated for all the services in the service group.<br>Minimum value = 0<br>Maximum value = 4294967287.
// Maxclient. Maximum number of simultaneous open connections for the service group.<br>Minimum value = 0<br>Maximum value = 4294967294.
// Maxreq. Maximum number of requests that can be sent on a persistent connection to the service group. <br> Note: Connection requests beyond this value are rejected.<br>Minimum value = 0<br>Maximum value = 65535.
// Monthreshold. Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.<br>Minimum value = 0<br>Maximum value = 65535.
// Name. Name for the server. <br>Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.<br>Can be changed after the name is created.<br>Minimum length = 1.
// Newname. New name for the server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.<br>Minimum length = 1.
// Sc. State of the SureConnect feature for the service group.<br>Default value: OFF<br>Possible values = ON, OFF.
// Sp. Enable surge protection for the service group.<br>Default value: OFF<br>Possible values = ON, OFF.
// State. Initial state of the server.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Statechangetimesec. Time when last state change happened. Seconds part.
// Svrtimeout. Time, in seconds, after which to terminate an idle server connection.<br>Minimum value = 0<br>Maximum value = 31536000.
// Tcpb. Enable TCP buffering for the service group.<br>Possible values = YES, NO.
// Td. Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.<br>Minimum value = 0<br>Maximum value = 4094.
// Tickssincelaststatechange. Time in 10 millisecond ticks since the last state change.
// Translationip. IP address used to transform the servers DNS-resolved IP address.
// Translationmask. The netmask of the translation ip.
// Usip. Use the clients IP address as the source IP address when initiating a connection to the server. When creating a service, if you do not set this parameter, the service inherits the global Use Source IP setting (available in the enable ns mode and disable ns mode CLI commands, or in the System ;gt; Settings ;gt; Configure modes ;gt; Configure Modes dialog box). However, you can override this setting after you create the service.<br>Possible values = YES, NO.

// ServerAdd defines add request.
type ServerAdd struct {
	Server ServerAddBody `json:"server"`
}

// ServerAddBody body for adding object.
type ServerAddBody struct {
	Name               string `json:"name"`
	Ipaddress          string `json:"ipaddress"`
	Domain             string `json:"domain,omitempty"`
	Translationip      string `json:"translationip,omitempty"`
	Translationmask    string `json:"translationmask,omitempty"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	State              string `json:"state,omitempty"`
	Ipv6Address        string `json:"ipv6address,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Td                 int    `json:"td,omitempty"`
}

// ServerUpdateBody body for updating object.
type ServerUpdateBody struct {
	Name               string `json:"name"`
	Ipaddress          string `json:"ipaddress"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	Translationip      string `json:"translationip,omitempty"`
	Translationmask    string `json:"translationmask,omitempty"`
	Domainresolvenow   bool   `json:"domainresolvenow,omitempty"`
	Comment            string `json:"comment,omitempty"`
}

// ServerEnableBody body for enabling object.
type ServerEnableBody struct {
	Name string `json:"name"`
}

// ServerDisableBody body for disabling object.
type ServerDisableBody struct {
	Name     string `json:"name"`
	Delay    int    `json:"delay"`
	Graceful string `json:"graceful"`
}

// ServerRenameBody body for renaming object.
type ServerRenameBody struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

// ServerUpdate defines update request.
type ServerUpdate struct {
	Server ServerUpdateBody `json:"server"`
}

// ServerEnable defines enable request.
type ServerEnable struct {
	Server ServerEnableBody `json:"server"`
}

// ServerDisable defines disable request.
type ServerDisable struct {
	Server ServerDisableBody `json:"server"`
}

// ServerRename defines rename request.
type ServerRename struct {
	Server ServerRenameBody `json:"server"`
}

// ServerWrapper wraps the object and serves as default response.
type ServerWrapper struct {
	Errorcode int      `json:"errorcode"`
	Message   string   `json:"message"`
	Severity  string   `json:"severity"`
	Server    []Server `json:"server"`
}

// Server describes the resource.
type Server struct {
	Name                      string `json:"name"`
	Ipaddress                 string `json:"ipaddress"`
	State                     string `json:"state"`
	Translationip             string `json:"translationip"`
	Translationmask           string `json:"translationmask"`
	Svctype                   string `json:"svctype"`
	Port                      int    `json:"port"`
	Svrstate                  string `json:"svrstate"`
	Statechangetimesec        string `json:"statechangetimesec"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange"`
	Ipv6Address               string `json:"ipv6address"`
	DupPort                   int    `json:"dup_port"`
	DupSvctype                string `json:"dup_svctype"`
	Svrcfgflags               string `json:"svrcfgflags"`
	Td                        string `json:"td"`
	Maxreq                    string `json:"maxreq"`
	Maxbandwidth              string `json:"maxbandwidth"`
	Usip                      string `json:"usip"`
	Cka                       string `json:"cka"`
	Tcpb                      string `json:"tcpb"`
	Comment                   string `json:"comment"`
	Cmp                       string `json:"cmp"`
	Clttimeout                int    `json:"clttimeout"`
	Svrtimeout                int    `json:"svrtimeout"`
	Cip                       string `json:"cip"`
	Cacheable                 string `json:"cacheable"`
	Sp                        string `json:"sp"`
	Downstateflush            string `json:"downstateflush"`
	Appflowlog                string `json:"appflowlog"`
	Boundtd                   string `json:"boundtd"`
}
