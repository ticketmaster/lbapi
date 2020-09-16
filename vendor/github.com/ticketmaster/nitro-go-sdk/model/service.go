package model

// Accessdown. Use Layer 2 mode to bridge the packets sent to this service if it is marked as DOWN. If the service is DOWN, and this parameter is disabled, the packets are dropped.<br>Default value: NO<br>Possible values = YES, NO.
// All. Display both user-configured and dynamically learned services.
// Appflowlog. Enable logging of AppFlow information.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Cacheable. Use the transparent cache redirection virtual server to forward requests to the cache server.<br>Note: Do not specify this parameter if you set the Cache Type parameter.<br>Default value: NO<br>Possible values = YES, NO.
// Cachetype. Cache type supported by the cache server.<br>Possible values = TRANSPARENT, REVERSE, FORWARD.
// Cip. Before forwarding a request to the service, insert an HTTP header with the clients IPv4 or IPv6 address as its value. Used if the server needs the clients IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option.<br>Possible values = ENABLED, DISABLED.
// Cipheader. Name for the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If you set the Client IP parameter, and you do not specify a name for the header, the appliance uses the header name specified for the global Client IP Header parameter (the cipHeader parameter in the set ns param CLI command or the Client IP Header parameter in the Configure HTTP Parameters dialog box at System ;gt; Settings ;gt; Change HTTP parameters). If the global Client IP Header parameter is not specified, the appliance inserts a header with the name "client-ip.".<br>Minimum length = 1.
// Cka. Enable client keep-alive for the service.<br>Possible values = YES, NO.
// Cleartextport. Port to which clear text data must be sent after the appliance decrypts incoming SSL traffic. Applicable to transparent SSL services.<br>Minimum value = 1.
// Clmonowner. Tells the mon owner of the service.<br>Minimum value = 0<br>Maximum value = 32.
// Clmonview. Tells the view id of the monitoring owner.
// Clttimeout. Time, in seconds, after which to terminate an idle client connection.<br>Minimum value = 0<br>Maximum value = 31536000.
// Cmp. Enable compression for the service.<br>Possible values = YES, NO.
// Comment. Any information about the service.
// Customserverid. Unique identifier for the service. Used when the persistency type for the virtual server is set to Custom Server ID.<br>Default value: "None".
// Delay. Time, in seconds, allocated to the NetScaler appliance for a graceful shutdown of the service. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).
// Dnsprofilename. Name of the DNS profile to be associated with the service. DNS profile properties will applied to the transactions processed by a service. This parameter is valid only for ADNS and ADNS-TCP services.<br>Minimum length = 1<br>Maximum length = 127.
// Downstateflush. Flush all active transactions associated with a service whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Dup_State. Added this field for getting state value from table.<br>Possible values = ENABLED, DISABLED.
// Graceful. Shut down gracefully, not accepting any new connections, and disabling the service when all of its connections are closed.<br>Default value: NO<br>Possible values = YES, NO.
// Gslb. The GSLB option for the corresponding virtual server.<br>Possible values = REMOTE, LOCAL.
// Hashid. A numerical identifier that can be used by hash based load balancing methods. Must be unique for each service.<br>Minimum value = 1.
// Healthmonitor. Monitor the health of this service. Available settings function as follows:<br>YES - Send probes to check the health of the service.<br>NO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times.<br>Default value: YES<br>Possible values = YES, NO.
// Httpprofilename. Name of the HTTP profile that contains HTTP configuration settings for the service.<br>Minimum length = 1<br>Maximum length = 127.
// Internal. Display only dynamically learned services.
// Ip. IP to assign to the service.<br>Minimum length = 1.
// Ipaddress. The new IP address of the service.
// Lastresponse. The string form of monstatcode.
// Maxbandwidth. Maximum bandwidth, in Kbps, allocated to the service.<br>Minimum value = 0<br>Maximum value = 4294967287.
// Maxclient. Maximum number of simultaneous open connections to the service.<br>Minimum value = 0<br>Maximum value = 4294967294.
// Maxreq. Maximum number of requests that can be sent on a persistent connection to the service. <br>Note: Connection requests beyond this value are rejected.<br>Minimum value = 0<br>Maximum value = 65535.
// Monconnectionclose. Close monitoring connections by sending the service a connection termination message with the specified bit set.<br>Default value: NONE<br>Possible values = RESET, FIN.
// Monitor_Name_Svc. Name of the monitor bound to the specified service.<br>Minimum length = 1.
// Monitor_State. The running state of the monitor on this service.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
// Monstatcode. The code indicating the monitor response.
// Monstatparam1. First parameter for use with message code.
// Monstatparam2. Second parameter for use with message code.
// Monstatparam3. Third parameter for use with message code.
// Monthreshold. Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.<br>Minimum value = 0<br>Maximum value = 65535.
// Name. Name for the service. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the service has been created.<br>Minimum length = 1.
// Netprofile. Network profile to use for the service.<br>Minimum length = 1<br>Maximum length = 127.
// Newname. New name for the service. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.<br>Minimum length = 1.
// Numofconnections. This will tell the number of client side connections are still open.
// Oracleserverversion. Oracle server version.<br>Default value: 10G<br>Possible values = 10G, 11G.
// Pathmonitor. Path monitoring for clustering.<br>Possible values = YES, NO.
// Pathmonitorindv. Individual Path monitoring decisions.<br>Possible values = YES, NO.
// Policyname. The name of the policyname for which this service is bound.
// Port. Port number of the service.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Processlocal. By turning on this option packets destined to a service in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Publicip. public ip.
// Publicport. public port.<br>Minimum value = 1<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Responsetime. Response time of this monitor.
// Riseapbrstatsmsgcode. The code indicating the rise apbr status.
// Riseapbrstatsmsgcode2. The code indicating other rise stats.
// Rtspsessionidremap. Enable RTSP session ID mapping for the service.<br>Default value: OFF<br>Possible values = ON, OFF.
// Sc. State of SureConnect for the service.<br>Default value: OFF<br>Possible values = ON, OFF.
// Serverid. The identifier for the service. This is used when the persistency type is set to Custom Server ID.
// Servername. Name of the server that hosts the service.<br>Minimum length = 1.
// Serviceconftype. The configuration type of the service.
// Serviceconftype2. The configuration type of the service (Internal/Dynamic/Configured).<br>Possible values = Configured, Dynamic, Internal.
// Serviceipstr. This field has been intorduced to show the dbs services ip.
// Servicetype. Protocol in which data is exchanged with the service.<br>Possible values = HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP.
// Sp. Enable surge protection for the service.<br>Possible values = ON, OFF.
// State. Initial state of the service.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Statechangetimemsec. Time at which last state change happened. Milliseconds part.
// Statechangetimesec. Time when last state change happened. Seconds part.
// Stateupdatereason. Checks state update reason on the secondary node.
// Svrstate. The state of the service.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
// Svrtimeout. Time, in seconds, after which to terminate an idle server connection.<br>Minimum value = 0<br>Maximum value = 31536000.
// Tcpb. Enable TCP buffering for the service.<br>Possible values = YES, NO.
// Tcpprofilename. Name of the TCP profile that contains TCP configuration settings for the service.<br>Minimum length = 1<br>Maximum length = 127.
// Td. Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.<br>Minimum value = 0<br>Maximum value = 4094.
// Tickssincelaststatechange. Time in 10 millisecond ticks since the last state change.
// Useproxyport. Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection. <br>Note: This parameter is available only when the Use Source IP (USIP) parameter is set to YES.<br>Possible values = YES, NO.
// Usip. Use the clients IP address as the source IP address when initiating a connection to the server. When creating a service, if you do not set this parameter, the service inherits the global Use Source IP setting (available in the enable ns mode and disable ns mode CLI commands, or in the System ;gt; Settings ;gt; Configure modes ;gt; Configure Modes dialog box). However, you can override this setting after you create the service.<br>Possible values = YES, NO.
// Value. SSL status.<br>Possible values = Certkey not bound, SSL feature disabled.
// Weight. Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.<br>Minimum value = 1<br>Maximum value = 100.

// ServiceAdd defines add request
type ServiceAdd struct {
	Service ServiceAddBody `json:"service,omitempty"`
}

// ServiceAddBody body for adding object.
type ServiceAddBody struct {
	Appflowlog         string `json:"appflowlog,omitempty"`
	Cacheable          string `json:"cacheable,omitempty"`
	Cachetype          string `json:"cachetype,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Cipheader          string `json:"cipheader,omitempty"`
	Cka                string `json:"cka,omitempty"`
	Cleartextport      int    `json:"cleartextport,omitempty"`
	Clttimeout         int    `json:"clttimeout,omitempty"`
	Cmp                string `json:"cmp,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Customserverid     string `json:"customserverid,omitempty"`
	Dnsprofilename     string `json:"dnsprofilename,omitempty"`
	Downstateflush     string `json:"downstateflush,omitempty"`
	Hashid             string `json:"hashid,omitempty"`
	Healthmonitor      string `json:"healthmonitor,omitempty"`
	Httpprofilename    string `json:"httpprofilename,omitempty"`
	IP                 string `json:"ip,omitempty"`
	Ipaddress          string `json:"ipaddress"`
	Maxbandwidth       string `json:"maxbandwidth,omitempty"`
	Maxclient          string `json:"maxclient,omitempty"`
	Maxreq             string `json:"maxreq,omitempty"`
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
	MonitorNameSvc     string `json:"monitor_name_svc,omitempty"`
	Monthreshold       string `json:"monthreshold,omitempty"`
	Name               string `json:"name"`
	Netprofile         string `json:"netprofile,omitempty"`
	Pathmonitor        string `json:"pathmonitor,omitempty"`
	Pathmonitorindv    string `json:"pathmonitorindv,omitempty"`
	Port               int    `json:"port"`
	Processlocal       string `json:"processlocal,omitempty"`
	Rtspsessionidremap string `json:"rtspsessionidremap,omitempty"`
	Sc                 string `json:"sc,omitempty"`
	Serverid           int    `json:"serverid,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Servicetype        string `json:"servicetype"`
	Sp                 string `json:"sp,omitempty"`
	State              string `json:"state,omitempty"`
	Svrtimeout         int    `json:"svrtimeout,omitempty"`
	Tcpb               string `json:"tcpb,omitempty"`
	Tcpprofilename     string `json:"tcpprofilename,omitempty"`
	Td                 string `json:"td,omitempty"`
	Useproxyport       string `json:"useproxyport,omitempty"`
	Usip               string `json:"usip,omitempty"`
	Weight             string `json:"weight,omitempty"`
}

// ServiceUpdateBody body for updating object.
type ServiceUpdateBody struct {
	Accessdown         string `json:"accessdown,omitempty"`
	Appflowlog         string `json:"appflowlog,omitempty"`
	Cacheable          string `json:"cacheable,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Cipheader          string `json:"cipheader,omitempty"`
	Cka                string `json:"cka,omitempty"`
	Clttimeout         int    `json:"clttimeout,omitempty"`
	Cmp                string `json:"cmp,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Customserverid     string `json:"customserverid,omitempty"`
	Dnsprofilename     string `json:"dnsprofilename,omitempty"`
	Downstateflush     string `json:"downstateflush,omitempty"`
	Hashid             string `json:"hashid,omitempty"`
	Healthmonitor      string `json:"healthmonitor,omitempty"`
	Httpprofilename    string `json:"httpprofilename,omitempty"`
	Ipaddress          string `json:"ipaddress"`
	Maxbandwidth       string `json:"maxbandwidth,omitempty"`
	Maxclient          string `json:"maxclient,omitempty"`
	Maxreq             string `json:"maxreq,omitempty"`
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
	MonitorNameSvc     string `json:"monitor_name_svc,omitempty"`
	Monthreshold       string `json:"monthreshold,omitempty"`
	Name               string `json:"name"`
	Netprofile         string `json:"netprofile,omitempty"`
	Pathmonitor        string `json:"pathmonitor,omitempty"`
	Pathmonitorindv    string `json:"pathmonitorindv,omitempty"`
	Processlocal       string `json:"processlocal,omitempty"`
	Rtspsessionidremap string `json:"rtspsessionidremap,omitempty"`
	Sc                 string `json:"sc,omitempty"`
	Serverid           string `json:"serverid,omitempty"`
	Sp                 string `json:"sp,omitempty"`
	Svrtimeout         string `json:"svrtimeout,omitempty"`
	Tcpb               string `json:"tcpb,omitempty"`
	Tcpprofilename     string `json:"tcpprofilename,omitempty"`
	Useproxyport       string `json:"useproxyport,omitempty"`
	Usip               string `json:"usip,omitempty"`
	Weight             string `json:"weight,omitempty"`
}

// ServiceEnableBody body for enabling object.
type ServiceEnableBody struct {
	Name string `json:"name,omitempty"`
}

// ServiceDisableBody body for disabling object.
type ServiceDisableBody struct {
	Name     string `json:"name,omitempty"`
	Delay    int    `json:"delay,omitempty"`
	Graceful string `json:"graceful,omitempty"`
}

//ServiceRenameBody body for renaming object.
type ServiceRenameBody struct {
	Name    string `json:"name,omitempty"`
	Newname string `json:"newname,omitempty"`
}

// ServiceUpdate defines update request.
type ServiceUpdate struct {
	Service ServiceUpdateBody `json:"service,omitempty"`
}

// ServiceEnable defines enable request.
type ServiceEnable struct {
	Service ServiceEnableBody `json:"service,omitempty"`
}

// ServiceDisable defines disable request.
type ServiceDisable struct {
	Service ServiceDisableBody `json:"service,omitempty"`
}

// ServiceRename defines rename request.
type ServiceRename struct {
	Service ServiceRenameBody `json:"service,omitempty"`
}

// ServiceWrapper wraps the object and serves as default response
type ServiceWrapper struct {
	Errorcode int       `json:"errorcode,omitempty"`
	Message   string    `json:"message,omitempty"`
	Severity  string    `json:"severity,omitempty"`
	Service   []Service `json:"service,omitempty"`
}

// Service describes the object.
type Service struct {
	Accessdown                string `json:"accessdown,omitempty"`
	All                       bool   `json:"all,omitempty"`
	Appflowlog                string `json:"appflowlog,omitempty"`
	Cacheable                 string `json:"cacheable,omitempty"`
	Cachetype                 string `json:"cachetype,omitempty"`
	Cip                       string `json:"cip,omitempty"`
	Cipheader                 string `json:"cipheader,omitempty"`
	Cka                       string `json:"cka,omitempty"`
	Cleartextport             int    `json:"cleartextport,omitempty"`
	Clmonowner                int    `json:"clmonowner,omitempty"`
	Clmonview                 string `json:"clmonview,omitempty"`
	Clttimeout                int    `json:"clttimeout,omitempty"`
	Cmp                       string `json:"cmp,omitempty"`
	Comment                   string `json:"comment,omitempty"`
	Customserverid            string `json:"customserverid,omitempty"`
	Delay                     int    `json:"delay,omitempty"`
	Dnsprofilename            string `json:"dnsprofilename,omitempty"`
	Downstateflush            string `json:"downstateflush,omitempty"`
	DupState                  string `json:"dup_state,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Gslb                      string `json:"gslb,omitempty"`
	Hashid                    string `json:"hashid,omitempty"`
	Healthmonitor             string `json:"healthmonitor,omitempty"`
	Httpprofilename           string `json:"httpprofilename,omitempty"`
	Internal                  bool   `json:"Internal,omitempty"`
	Ipaddress                 string `json:"ipaddress"`
	Lastresponse              string `json:"lastresponse,omitempty"`
	Maxbandwidth              string `json:"maxbandwidth,omitempty"`
	Maxclient                 string `json:"maxclient,omitempty"`
	Maxreq                    string `json:"maxreq,omitempty"`
	Monconnectionclose        string `json:"monconnectionclose,omitempty"`
	MonitorState              string `json:"monitor_state,omitempty"`
	Monstatcode               int    `json:"monstatcode,omitempty"`
	Monstatparam1             int    `json:"monstatparam1,omitempty"`
	Monstatparam2             int    `json:"monstatparam2,omitempty"`
	Monstatparam3             int    `json:"monstatparam3,omitempty"`
	Monthreshold              string `json:"monthreshold,omitempty"`
	Name                      string `json:"name"`
	Netprofile                string `json:"netprofile,omitempty"`
	Numofconnections          int    `json:"numofconnections,omitempty"`
	Oracleserverversion       string `json:"oracleserverversion,omitempty"`
	Pathmonitor               string `json:"pathmonitor,omitempty"`
	Pathmonitorindv           string `json:"pathmonitorindv,omitempty"`
	Policyname                string `json:"policyname,omitempty"`
	Port                      int    `json:"port"`
	Processlocal              string `json:"processlocal,omitempty"`
	Publicip                  string `json:"publicip,omitempty"`
	Publicport                int    `json:"publicport,omitempty"`
	Responsetime              string `json:"responsetime,omitempty"`
	Riseapbrstatsmsgcode      int    `json:"riseapbrstatsmsgcode,omitempty"`
	Riseapbrstatsmsgcode2     int    `json:"riseapbrstatsmsgcode2,omitempty"`
	Rtspsessionidremap        string `json:"rtspsessionidremap,omitempty"`
	Sc                        string `json:"sc,omitempty"`
	Serverid                  int    `json:"serverid,omitempty"`
	Servername                string `json:"servername,omitempty"`
	Serviceconftpye           bool   `json:"serviceconftpye,omitempty"`
	Serviceconftype           bool   `json:"serviceconftype,omitempty"`
	Serviceconftype2          string `json:"serviceconftype2,omitempty"`
	Serviceipstr              string `json:"serviceipstr,omitempty"`
	Servicetype               string `json:"servicetype"`
	Sp                        string `json:"sp,omitempty"`
	Statechangetimemsec       string `json:"statechangetimemsec,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Stateupdatereason         string `json:"stateupdatereason,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Svrtimeout                int    `json:"svrtimeout,omitempty"`
	Tcpb                      string `json:"tcpb,omitempty"`
	Tcpprofilename            string `json:"tcpprofilename,omitempty"`
	Td                        string `json:"td,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Timesincelaststatechange  string `json:"timesincelaststatechange,omitempty"`
	Useproxyport              string `json:"useproxyport,omitempty"`
	Usip                      string `json:"usip,omitempty"`
	Value                     string `json:"value,omitempty"`
	Weight                    string `json:"weight,omitempty"`
}
