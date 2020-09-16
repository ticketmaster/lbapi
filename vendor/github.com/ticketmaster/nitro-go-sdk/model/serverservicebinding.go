package model

// Name. Name of the server for which to display parameters.<br>Minimum length = 1.
// Port. The port number to be used for the bound service.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Serviceipaddress. The IP address of the bound service.
// Serviceipstr. This field has been intorduced to show the dbs services ip.
// Servicename. The services attatched to the server.
// Svctype. The type of bound service.<br>Possible values = HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP.
// Svrstate. The state of the bound service.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.

// ServerServiceBindingWrapper wraps the object and serves as default response.
type ServerServiceBindingWrapper struct {
	Errorcode            int                    `json:"errorcode"`
	Message              string                 `json:"message"`
	Severity             string                 `json:"severity"`
	ServerServiceBinding []ServerServiceBinding `json:"server_service_binding"`
}

// ServerServiceBinding describes the resource.
type ServerServiceBinding struct {
	Name             string `json:"name"`
	Servicename      string `json:"servicename"`
	Serviceipstr     string `json:"serviceipstr"`
	Svctype          string `json:"svctype"`
	Svrstate         string `json:"svrstate"`
	Port             int    `json:"port"`
	Serviceipaddress string `json:"serviceipaddress"`
}
