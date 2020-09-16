package model

// Cookieipport. Encryped Ip address and port of the service that is inserted into the set-cookie http header.
// Cookiename. Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.
// Curstate. Current LB vserver state.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
// Dynamicweight. Dynamic weight.
// Ipv46. IPv4 or IPv6 address to assign to the virtual server.
// Name. Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or my vserver)..<br>Minimum length = 1.
// Port. Port number for the virtual server.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Preferredlocation. Used for displaying the location of bound services.
// Servicegroupname. The service group name bound to the selected load balancing virtual server.
// Servicetype. Protocol used by the service (also called the service type).<br>Possible values = HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, DNS, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, RTSP, PUSH, SSL_PUSH, RADIUS, RDP, MYSQL, MSSQL, DIAMETER, SSL_DIAMETER, TFTP, ORACLE, SMPP, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP.
// Vserverid. Vserver Id.
// Weight. Weight to assign to the specified service.<br>Minimum value = 1<br>Maximum value = 100.

// LbvserverServicegroupmemberBindingWrapper wraps the object and serves as default response
type LbvserverServicegroupmemberBindingWrapper struct {
	Errorcode                          int                                  `json:"errorcode"`
	Message                            string                               `json:"message"`
	Severity                           string                               `json:"severity"`
	LbvserverServicegroupmemberBinding []LbvserverServicegroupmemberBinding `json:"lbvserver_servicegroup_binding"`
}

// LbvserverServicegroupmemberBinding describes the object.
type LbvserverServicegroupmemberBinding []struct {
	Servicegroupname  string `json:"servicegroupname"`
	Name              string `json:"name"`
	Weight            int    `json:"weight"`
	Cookieipport      string `json:"cookieipport"`
	Port              int    `json:"port"`
	Curstate          string `json:"curstate"`
	Preferredlocation string `json:"preferredlocation"`
	Vserverid         string `json:"vserverid"`
	Ipv46             string `json:"ipv46"`
	Cookiename        string `json:"cookiename"`
	Dynamicweight     int    `json:"dynamicweight"`
	Servicetype       string `json:"servicetype"`
}
