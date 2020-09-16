package model

// boundservicegroupsvrstate	<String>	Read-only	The state of the servicegroup.<br>Possible values = ENABLED, DISABLED
// monitorname	<String>	Read-write	The name of the monitor.<br>Minimum length = 1
// monstate	<String>	Read-only	The configured state (enable/disable) of Monitor on this service.<br>Possible values = ENABLED, DISABLED
// servicegroupname	<String>	Read-write	The name of the service group.
// servicetype	<String>	Read-only	The type of service.<br>Possible values = HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP

// LbmonbindingsServicegroupBindingWrapper wraps the object and serves as default response.
type LbmonbindingsServicegroupBindingWrapper struct {
	Errorcode                        int                                `json:"errorcode,omitempty"`
	Message                          string                             `json:"message,omitempty"`
	Severity                         string                             `json:"severity,omitempty"`
	LbmonbindingsServicegroupBinding []LbmonbindingsServicegroupBinding `json:"lbmonbindings_servicegroup_binding,omitempty"`
}

// LbmonbindingsServicegroupBinding describes the object.
type LbmonbindingsServicegroupBinding struct {
	Monitorname               string `json:"monitorname,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	Boundservicegroupsvrstate string `json:"boundservicegroupsvrstate,omitempty"`
	Monstate                  string `json:"monstate,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
}
