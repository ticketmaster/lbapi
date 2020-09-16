package model

// Cookieipport. Encryped Ip address and port of the service that is inserted into the set-cookie http header.
// Curstate. Current LB vserver state.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
// Dynamicweight. Dynamic weight.
// Ipv46. IPv4 or IPv6 address to assign to the virtual server.
// Name. Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or my vserver). .<br>Minimum length = 1.
// Port. Port number for the virtual server.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Preferredlocation. Used for displaying the location of bound services.
// Servicegroupname. Name of the service group.<br>Minimum length = 1.
// Servicename. Service to bind to the virtual server.<br>Minimum length = 1.
// Servicetype. Protocol used by the service (also called the service type).<br>Possible values = HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, DNS, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, RTSP, PUSH, SSL_PUSH, RADIUS, RDP, MYSQL, MSSQL, DIAMETER, SSL_DIAMETER, TFTP, ORACLE, SMPP, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP.
// Vserverid. Vserver Id.
// Vsvrbindsvcip. used for showing the ip of bound entities.
// Vsvrbindsvcport. used for showing ports of bound entities.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Weight. Weight to assign to the specified service.<br>Minimum value = 1<br>Maximum value = 100.

// LbvserverServiceBindingAdd defines add request
type LbvserverServiceBindingAdd struct {
	LbvserverServiceBinding LbvserverServiceBindingAddBody `json:"lbvserver_service_binding"`
}

// LbvserverServiceBindingAddBody body for adding object.
type LbvserverServiceBindingAddBody struct {
	Name             string `json:"name"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename"`
	Weight           string `json:"weight,omitempty"`
}

// LbvserverServiceBindingWrapper wraps the object and serves as default response.
type LbvserverServiceBindingWrapper struct {
	Errorcode               int                       `json:"errorcode,omitempty"`
	Message                 string                    `json:"message,omitempty"`
	Severity                string                    `json:"severity,omitempty"`
	LbvserverServiceBinding []LbvserverServiceBinding `json:"lbvserver_service_binding,omitempty"`
}

// LbvserverServiceBinding describes the object.
type LbvserverServiceBinding struct {
	Cookieipport      string `json:"cookieipport"`
	Curstate          string `json:"curstate"`
	Dynamicweight     string `json:"dynamicweight"`
	Ipv46             string `json:"ipv46"`
	Name              string `json:"name"`
	Port              int    `json:"port"`
	Preferredlocation string `json:"preferredlocation"`
	Servicegroupname  string `json:"servicegroupname"`
	Servicename       string `json:"servicename"`
	Servicetype       string `json:"servicetype"`
	Vserverid         string `json:"vserverid"`
	Vsvrbindsvcip     string `json:"vsvrbindsvcip"`
	Vsvrbindsvcport   int    `json:"vsvrbindsvcport"`
	Weight            string `json:"weight"`
}
