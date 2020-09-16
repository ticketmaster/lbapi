package model

// Appflowlog. Enable logging of AppFlow information for the specified service group.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Boundtd. Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.<br>Minimum value = 0<br>Maximum value = 4094.
// Cacheable. Use the transparent cache redirection virtual server to forward the request to the cache server.<br>Default value: NO<br>Possible values = YES, NO.
// Cip. Before forwarding a request to the service, insert an HTTP header with the clients IPv4 or IPv6 address as its value. Used if the server needs the clients IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option.<br>Possible values = ENABLED, DISABLED.
// Cipheader. Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as clients IP header name.<br>Minimum length = 1.
// Cka. Enable client keep-alive for the service group.<br>Possible values = YES, NO.
// Clttimeout. Time, in seconds, after which to terminate an idle client connection.<br>Minimum value = 0<br>Maximum value = 31536000.
// Cmp. Enable compression for the specified service.<br>Possible values = YES, NO.
// Customserverid. A positive integer to identify the service. Used when the persistency type is set to Custom Server ID.<br>Default value: "None".
// Downstateflush. Perform delayed clean-up of connections to all services in the service group.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Dup_Port. port of the service.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Dup_Svctype. service type of the service.<br>Possible values = HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP.
// Maxbandwidth. Maximum bandwidth, in Kbps, allocated for all the services in the service group.<br>Minimum value = 0<br>Maximum value = 4294967287.
// Maxclient. Maximum number of simultaneous open connections for the service group.<br>Minimum value = 0<br>Maximum value = 4294967294.
// Maxreq. Maximum number of requests that can be sent on a persistent connection to the service group. Note: Connection requests beyond this value are rejected.<br>Minimum value = 0<br>Maximum value = 65535.
// Monthreshold. Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.<br>Minimum value = 0<br>Maximum value = 65535.
// Name. Name of the server for which to display parameters.<br>Minimum length = 1.
// Port. The port number to be used for the bound service.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Sc. State of the SureConnect feature for the service group.<br>Default value: OFF<br>Possible values = ON, OFF.
// Servicegroupname. servicegroups bind to this server.
// Serviceipaddress. The IP address of the bound service.
// Serviceipstr. This field has been intorduced to show the dbs services ip.
// Sp. Enable surge protection for the service group.<br>Default value: OFF<br>Possible values = ON, OFF.
// Svctype. The type of bound service.<br>Possible values = HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP.
// Svrcfgflags. service flags to denote its a db enabled.
// Svrstate. The state of the bound service.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
// Svrtimeout. Time, in seconds, after which to terminate an idle server connection.<br>Minimum value = 0<br>Maximum value = 31536000.
// Tcpb. Enable TCP buffering for the service group.<br>Possible values = YES, NO.
// Usip. Use the clients IP address as the source IP address when initiating a connection to the server. When creating a service, if you do not set this parameter, the service inherits the global Use Source IP setting (available in the enable ns mode and disable ns mode CLI commands, or in the System ;gt; Settings ;gt; Configure modes ;gt; Configure Modes dialog box). However, you can override this setting after you create the service.<br>Possible values = YES, NO.

// ServerServicegroupBindingWrapper wraps the object and serves as default response.
type ServerServicegroupBindingWrapper struct {
	Errorcode                 int                         `json:"errorcode"`
	Message                   string                      `json:"message"`
	Severity                  string                      `json:"severity"`
	ServerServicegroupBinding []ServerServicegroupBinding `json:"server_servicegroup_binding"`
}

// ServerServicegroupBinding describes the resource.
type ServerServicegroupBinding struct {
	Name             string `json:"name"`
	Servicegroupname string `json:"servicegroupname"`
	Sp               string `json:"sp"`
	Sc               string `json:"sc"`
	Svrtimeout       int    `json:"svrtimeout"`
	Serviceipstr     string `json:"serviceipstr"`
	Tcpb             string `json:"tcpb"`
	Cip              string `json:"cip"`
	DupPort          int    `json:"dup_port"`
	Customserverid   string `json:"customserverid"`
	Boundtd          string `json:"boundtd"`
	Downstateflush   string `json:"downstateflush"`
	Svctype          string `json:"svctype"`
	Cka              string `json:"cka"`
	Svrstate         string `json:"svrstate"`
	Appflowlog       string `json:"appflowlog"`
	Cipheader        string `json:"cipheader"`
	Maxclient        string `json:"maxclient"`
	Clttimeout       int    `json:"clttimeout"`
	Port             int    `json:"port"`
	DupSvctype       string `json:"dup_svctype"`
	Cmp              string `json:"cmp"`
	Monthreshold     string `json:"monthreshold"`
	Maxreq           string `json:"maxreq"`
	Serviceipaddress string `json:"serviceipaddress"`
	Usip             string `json:"usip"`
	Cacheable        string `json:"cacheable"`
	Maxbandwidth     string `json:"maxbandwidth"`
	Svrcfgflags      string `json:"svrcfgflags"`
}
