package model

// ipaddress	<String>	Read-only	The IPAddress of the service.
// monitorname	<String>	Read-write	The name of the monitor.<br>Minimum length = 1
// monsvcstate	<String>	Read-only	The configured state (enable/disable) of Monitor on this service.<br>Possible values = ENABLED, DISABLED
// port	<Integer>	Read-only	The port of the service.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API
// servicename	<String>	Read-write	The name of the service.
// servicetype	<String>	Read-only	The type of service.<br>Possible values = HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP
// svrstate	<String>	Read-only	The state of the service.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED

// LbmonbindingsServiceBindingWrapper wraps the object and serves as default response.
type LbmonbindingsServiceBindingWrapper struct {
	Errorcode                   int                           `json:"errorcode,omitempty"`
	Message                     string                        `json:"message,omitempty"`
	Severity                    string                        `json:"severity,omitempty"`
	LbmonbindingsServiceBinding []LbmonbindingsServiceBinding `json:"lbmonbindings_service_binding,omitempty"`
}

// LbmonbindingsServiceBinding describes the object.
type LbmonbindingsServiceBinding struct {
	Servicename string `json:"servicename"`
	Monitorname string `json:"monitorname"`
	Servicetype string `json:"servicetype"`
	Svrstate    string `json:"svrstate"`
	Monsvcstate string `json:"monsvcstate"`
	Port        int    `json:"port"`
	Ipaddress   string `json:"ipaddress"`
}
