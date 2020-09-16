package model

// acctapplicationid	<Double[]>	Read-write	List of Acct-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message.<br>Minimum value = 0<br>Maximum value = 4294967295
// action	<String>	Read-write	Action to perform when the response to an inline monitor (a monitor of type HTTP-INLINE) indicates that the service is down. A service monitored by an inline monitor is considered DOWN if the response code is not one of the codes that have been specified for the Response Code parameter. <br>Available settings function as follows: <br>* NONE - Do not take any action. However, the show service command and the show lb monitor command indicate the total number of responses that were checked and the number of consecutive error responses received after the last successful probe.<br>* LOG - Log the event in NSLOG or SYSLOG. <br>* DOWN - Mark the service as being down, and then do not direct any traffic to the service until the configured down time has expired. Persistent connections to the service are terminated as soon as the service is marked as DOWN. Also, log the event in NSLOG or SYSLOG.<br>Default value: DOWN<br>Possible values = NONE, LOG, DOWN
// alertretries	<Integer>	Read-write	Number of consecutive probe failures after which the appliance generates an SNMP trap called monProbeFailed.<br>Minimum value = 0<br>Maximum value = 32
// application	<String>	Read-write	Name of the application used to determine the state of the service. Applicable to monitors of type CITRIX-XML-SERVICE.<br>Minimum length = 1
// attribute	<String>	Read-write	Attribute to evaluate when the LDAP server responds to the query. Success or failure of the monitoring probe depends on whether the attribute exists in the response. Optional.<br>Minimum length = 1
// authapplicationid	<Double[]>	Read-write	List of Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring CER message.<br>Minimum value = 0<br>Maximum value = 4294967295
// basedn	<String>	Read-write	The base distinguished name of the LDAP service, from where the LDAP server can begin the search for the attributes in the monitoring query. Required for LDAP service monitoring.<br>Minimum length = 1
// binddn	<String>	Read-write	The distinguished name with which an LDAP monitor can perform the Bind operation on the LDAP server. Optional. Applicable to LDAP monitors.<br>Minimum length = 1
// customheaders	<String>	Read-write	Custom header string to include in the monitoring probes.
// database	<String>	Read-write	Name of the database to connect to during authentication.<br>Minimum length = 1
// destip	<String>	Read-write	IP address of the service to which to send probes. If the parameter is set to 0, the IP address of the server to which the monitor is bound is considered the destination IP address.
// destport	<Integer>	Read-write	TCP or UDP port to which to send the probe. If the parameter is set to 0, the port number of the service to which the monitor is bound is considered the destination port. For a monitor of type USER, however, the destination port is the port number that is included in the HTTP request sent to the dispatcher. Does not apply to monitors of type PING.
// deviation	<Double>	Read-write	Time value added to the learned average response time in dynamic response time monitoring (DRTM). When a deviation is specified, the appliance learns the average response time of bound services and adds the deviation to the average. The final value is then continually adjusted to accommodate response time variations over time. Specified in milliseconds, seconds, or minutes.<br>Minimum value = 0<br>Maximum value = 20939
// dispatcherip	<String>	Read-write	IP address of the dispatcher to which to send the probe.
// dispatcherport	<Integer>	Read-write	Port number on which the dispatcher listens for the monitoring probe.
// domain	<String>	Read-write	Domain in which the XenDesktop Desktop Delivery Controller (DDC) servers or Web Interface servers are present. Required by CITRIX-XD-DDC and CITRIX-WI-EXTENDED monitors for logging on to the DDC servers and Web Interface servers, respectively.
// downtime	<Integer>	Read-write	Time duration for which to wait before probing a service that has been marked as DOWN. Expressed in milliseconds, seconds, or minutes.<br>Default value: 30<br>Minimum value = 1<br>Maximum value = 20939
// dup_state	<String>	Read-only	.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// dup_weight	<Double>	Read-only	.<br>Default value: 1<br>Minimum value = 1<br>Maximum value = 100
// dynamicinterval	<Integer>	Read-only	Interval between monitoring probes for DRTM enabled monitor , calculated dynamically based monitor response time.
// dynamicresponsetimeout	<Integer>	Read-only	Response timeout of the DRTM enabled monitor , calculated dynamically based on the history and current response time.
// evalrule	<String>	Read-write	Default syntax expression that evaluates the database servers response to a MYSQL-ECV or MSSQL-ECV monitoring query. Must produce a Boolean result. The result determines the state of the server. If the expression returns TRUE, the probe succeeds. <br>For example, if you want the appliance to evaluate the error message to determine the state of the server, use the rule MYSQL.RES.ROW(10) .TEXT_ELEM(2).EQ("MySQL").
// failureretries	<Integer>	Read-write	Number of retries that must fail, out of the number specified for the Retries parameter, for a service to be marked as DOWN. For example, if the Retries parameter is set to 10 and the Failure Retries parameter is set to 6, out of the ten probes sent, at least six probes must fail if the service is to be marked as DOWN. The default value of 0 means that all the retries must fail if the service is to be marked as DOWN.<br>Minimum value = 0<br>Maximum value = 32
// filename	<String>	Read-write	Name of a file on the FTP server. The appliance monitors the FTP service by periodically checking the existence of the file on the server. Applicable to FTP-EXTENDED monitors.<br>Minimum length = 1
// filter	<String>	Read-write	Filter criteria for the LDAP query. Optional.<br>Minimum length = 1
// firmwarerevision	<Double>	Read-write	Firmware-Revision value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
// group	<String>	Read-write	Name of a newsgroup available on the NNTP service that is to be monitored. The appliance periodically generates an NNTP query for the name of the newsgroup and evaluates the response. If the newsgroup is found on the server, the service is marked as UP. If the newsgroup does not exist or if the search fails, the service is marked as DOWN. Applicable to NNTP monitors.<br>Minimum length = 1
// hostipaddress	<String>	Read-write	Host-IP-Address value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. If Host-IP-Address is not specified, the appliance inserts the mapped IP (MIP) address or subnet IP (SNIP) address from which the CER request (the monitoring probe) is sent.<br>Minimum length = 1
// hostname	<String>	Read-write	Hostname in the FQDN format (Example: porche.cars.org). Applicable to STOREFRONT monitors.<br>Minimum length = 1
// httprequest	<String>	Read-write	HTTP request to send to the server (for example, "HEAD /file.html").
// inbandsecurityid	<String>	Read-write	Inband-Security-Id for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.<br>Possible values = NO_INBAND_SECURITY, TLS
// interval	<Integer>	Read-write	Time interval between two successive probes. Must be greater than the value of Response Time-out.<br>Default value: 5<br>Minimum value = 1<br>Maximum value = 20940
// ipaddress	<String[]>	Read-write	Set of IP addresses expected in the monitoring response from the DNS server, if the record type is A or AAAA. Applicable to DNS monitors.<br>Minimum length = 1
// iptunnel	<String>	Read-write	Send the monitoring probe to the service through an IP tunnel. A destination IP address must be specified.<br>Default value: NO<br>Possible values = YES, NO
// kcdaccount	<String>	Read-write	KCD Account used by MSSQL monitor.<br>Minimum length = 1<br>Maximum length = 32
// lasversion	<String>	Read-write	Version number of the Citrix Advanced Access Control Logon Agent. Required by the CITRIX-AAC-LAS monitor.
// logonpointname	<String>	Read-write	Name of the logon point that is configured for the Citrix Access Gateway Advanced Access Control software. Required if you want to monitor the associated login page or Logon Agent. Applicable to CITRIX-AAC-LAS and CITRIX-AAC-LOGINPAGE monitors.
// lrtm	<String>	Read-write	Calculate the least response times for bound services. If this parameter is not enabled, the appliance does not learn the response times of the bound services. Also used for LRTM load balancing.<br>Possible values = ENABLED, DISABLED
// lrtmconf	<Integer>	Read-only	State of LRTM configuration on the monitor.
// lrtmconfstr	<String>	Read-only	State of LRTM configuration on the monitor as STRING.<br>Possible values = ENABLED, DISABLED
// maxforwards	<Double>	Read-write	Maximum number of hops that the SIP request used for monitoring can traverse to reach the server. Applicable only to monitors of type SIP-UDP.<br>Default value: 1<br>Minimum value = 0<br>Maximum value = 255
// metric	<String>	Read-write	Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation.<br>Minimum length = 1<br>Maximum length = 37
// metrictable	<String>	Read-write	Metric table to which to bind metrics.<br>Minimum length = 1<br>Maximum length = 99
// metricthreshold	<Double>	Read-write	Threshold to be used for that metric.
// metricweight	<Double>	Read-write	The weight for the specified service metric with respect to others.<br>Minimum value = 1<br>Maximum value = 100
// monitorname	<String>	Read-write	Name for the monitor. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.<br><br>CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my monitor" or my monitor).<br>Minimum length = 1
// mssqlprotocolversion	<String>	Read-write	Version of MSSQL server that is to be monitored.<br>Default value: 70<br>Possible values = 70, 2000, 2000SP1, 2005, 2008, 2008R2, 2012, 2014
// multimetrictable	<String[]>	Read-only	Metric table to which to bind metrics, to be used only for output purposes.
// netprofile	<String>	Read-write	Name of the network profile.<br>Minimum length = 1<br>Maximum length = 127
// oraclesid	<String>	Read-write	Name of the service identifier that is used to connect to the Oracle database during authentication.<br>Minimum length = 1
// originhost	<String>	Read-write	Origin-Host value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.<br>Minimum length = 1
// originrealm	<String>	Read-write	Origin-Realm value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.<br>Minimum length = 1
// password	<String>	Read-write	Password that is required for logging on to the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC-ECV or CITRIX-XDM server. Used in conjunction with the user name specified for the User Name parameter.<br>Minimum length = 1
// productname	<String>	Read-write	Product-Name value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.<br>Minimum length = 1
// query	<String>	Read-write	Domain name to resolve as part of monitoring the DNS service (for example, example.com).
// querytype	<String>	Read-write	Type of DNS record for which to send monitoring queries. Set to Address for querying A records, AAAA for querying AAAA records, and Zone for querying the SOA record.<br>Possible values = Address, Zone, AAAA
// radaccountsession	<String>	Read-write	Account Session ID to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.<br>Minimum length = 1
// radaccounttype	<Double>	Read-write	Account Type to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.<br>Default value: 1<br>Minimum value = 0<br>Maximum value = 15
// radapn	<String>	Read-write	Called Station Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.<br>Minimum length = 1
// radframedip	<String>	Read-write	Source ip with which the packet will go out . Applicable to monitors of type RADIUS_ACCOUNTING.
// radkey	<String>	Read-write	Authentication key (shared secret text string) for RADIUS clients and servers to exchange. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.<br>Minimum length = 1
// radmsisdn	<String>	Read-write	Calling Stations Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.<br>Minimum length = 1
// radnasid	<String>	Read-write	NAS-Identifier to send in the Access-Request packet. Applicable to monitors of type RADIUS.<br>Minimum length = 1
// radnasip	<String>	Read-write	Network Access Server (NAS) IP address to use as the source IP address when monitoring a RADIUS server. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.
// recv	<String>	Read-write	String expected from the server for the service to be marked as UP. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.
// respcode	<String[]>	Read-write	Response codes for which to mark the service as UP. For any other response code, the action performed depends on the monitor type. HTTP monitors and RADIUS monitors mark the service as DOWN, while HTTP-INLINE monitors perform the action indicated by the Action parameter.
// resptimeout	<Integer>	Read-write	Amount of time for which the appliance must wait before it marks a probe as FAILED. Must be less than the value specified for the Interval parameter.<br><br>Note: For UDP-ECV monitors for which a receive string is not configured, response timeout does not apply. For UDP-ECV monitors with no receive string, probe failure is indicated by an ICMP port unreachable error received from the service.<br>Default value: 2<br>Minimum value = 1<br>Maximum value = 20939
// resptimeoutthresh	<Double>	Read-write	Response time threshold, specified as a percentage of the Response Time-out parameter. If the response to a monitor probe has not arrived when the threshold is reached, the appliance generates an SNMP trap called monRespTimeoutAboveThresh. After the response time returns to a value below the threshold, the appliance generates a monRespTimeoutBelowThresh SNMP trap. For the traps to be generated, the "MONITOR-RTO-THRESHOLD" alarm must also be enabled.<br>Minimum value = 0<br>Maximum value = 100
// retries	<Integer>	Read-write	Maximum number of probes to send to establish the state of a service for which a monitoring probe failed.<br>Default value: 3<br>Minimum value = 1<br>Maximum value = 127
// reverse	<String>	Read-write	Mark a service as DOWN, instead of UP, when probe criteria are satisfied, and as UP instead of DOWN when probe criteria are not satisfied.<br>Default value: NO<br>Possible values = YES, NO
// rtsprequest	<String>	Read-write	RTSP request to send to the server (for example, "OPTIONS *").
// scriptargs	<String>	Read-write	String of arguments for the script. The string is copied verbatim into the request.
// scriptname	<String>	Read-write	Path and name of the script to execute. The script must be available on the NetScaler appliance, in the /nsconfig/monitors/ directory.<br>Minimum length = 1
// secondarypassword	<String>	Read-write	Secondary password that users might have to provide to log on to the Access Gateway server. Applicable to CITRIX-AG monitors.
// secure	<String>	Read-write	Use a secure SSL connection when monitoring a service. Applicable only to TCP based monitors. The secure option cannot be used with a CITRIX-AG monitor, because a CITRIX-AG monitor uses a secure connection by default.<br>Default value: NO<br>Possible values = YES, NO
// send	<String>	Read-write	String to send to the service. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.
// servicegroupname	<String>	Read-write	The name of the service group to which the monitor is to be bound.<br>Minimum length = 1
// servicename	<String>	Read-write	The name of the service to which the monitor is bound.<br>Minimum length = 1
// sipmethod	<String>	Read-write	SIP method to use for the query. Applicable only to monitors of type SIP-UDP.<br>Possible values = OPTIONS, INVITE, REGISTER
// sipreguri	<String>	Read-write	SIP user to be registered. Applicable only if the monitor is of type SIP-UDP and the SIP Method parameter is set to REGISTER.<br>Minimum length = 1
// sipuri	<String>	Read-write	SIP URI string to send to the service (for example, sip:sip.test). Applicable only to monitors of type SIP-UDP.<br>Minimum length = 1
// sitepath	<String>	Read-write	URL of the logon page. For monitors of type CITRIX-WEB-INTERFACE, to monitor a dynamic page under the site path, terminate the site path with a slash (/). Applicable to CITRIX-WEB-INTERFACE, CITRIX-WI-EXTENDED and CITRIX-XDM monitors.<br>Minimum length = 1
// snmpcommunity	<String>	Read-write	Community name for SNMP monitors.<br>Minimum length = 1
// Snmpoid	<String>	Read-write	SNMP OID for SNMP monitors.<br>Minimum length = 1
// snmpthreshold	<String>	Read-write	Threshold for SNMP monitors.<br>Minimum length = 1
// snmpversion	<String>	Read-write	SNMP version to be used for SNMP monitors.<br>Possible values = V1, V2
// sqlquery	<String>	Read-write	SQL query for a MYSQL-ECV or MSSQL-ECV monitor. Sent to the database server after the server authenticates the connection.<br>Minimum length = 1
// sslprofile	<String>	Read-write	SSL Profile associated with the monitor.<br>Minimum length = 1<br>Maximum length = 127
// state	<String>	Read-write	State of the monitor. The DISABLED setting disables not only the monitor being configured, but all monitors of the same type, until the parameter is set to ENABLED. If the monitor is bound to a service, the state of the monitor is not taken into account when the state of the service is determined.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// storedb	<String>	Read-write	Store the database list populated with the responses to monitor probes. Used in database specific load balancing if MSSQL-ECV/MYSQL-ECV monitor is configured.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// storefrontacctservice	<String>	Read-write	Enable/Disable probing for Account Service. Applicable only to Store Front monitors. For multi-tenancy configuration users my skip account service.<br>Default value: YES<br>Possible values = YES, NO
// storefrontcheckbackendservices	<String>	Read-write	This option will enable monitoring of services running on storefront server. Storefront services are monitored by probing to a Windows service that runs on the Storefront server and exposes details of which storefront services are running.<br>Default value: NO<br>Possible values = YES, NO
// storename	<String>	Read-write	Store Name. For monitors of type STOREFRONT, STORENAME is an optional argument defining storefront service store name. Applicable to STOREFRONT monitors.<br>Minimum length = 1
// successretries	<Integer>	Read-write	Number of consecutive successful probes required to transition a services state from DOWN to UP.<br>Default value: 1<br>Minimum value = 1<br>Maximum value = 32
// supportedvendorids	<Double[]>	Read-write	List of Supported-Vendor-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum eight of these AVPs are supported in a monitoring message.<br>Minimum value = 1<br>Maximum value = 4294967295
// tos	<String>	Read-write	Probe the service by encoding the destination IP address in the IP TOS (6) bits.<br>Possible values = YES, NO
// tosid	<Double>	Read-write	The TOS ID of the specified destination IP. Applicable only when the TOS parameter is set.<br>Minimum value = 1<br>Maximum value = 63
// transparent	<String>	Read-write	The monitor is bound to a transparent device such as a firewall or router. The state of a transparent device depends on the responsiveness of the services behind it. If a transparent device is being monitored, a destination IP address must be specified. The probe is sent to the specified IP address by using the MAC address of the transparent device.<br>Default value: NO<br>Possible values = YES, NO
// trofscode	<Double>	Read-write	Code expected when the server is under maintenance.
// trofsstring	<String>	Read-write	String expected from the server for the service to be marked as trofs. Applicable to HTTP-ECV/TCP-ECV monitors.
// type	<String>	Read-write	Type of monitor that you want to create.<br>Possible values = PING, TCP, HTTP, TCP-ECV, HTTP-ECV, UDP-ECV, DNS, FTP, LDNS-PING, LDNS-TCP, LDNS-DNS, RADIUS, USER, HTTP-INLINE, SIP-UDP, SIP-TCP, LOAD, FTP-EXTENDED, SMTP, SNMP, NNTP, MYSQL, MYSQL-ECV, MSSQL-ECV, ORACLE-ECV, LDAP, POP3, CITRIX-XML-SERVICE, CITRIX-WEB-INTERFACE, DNS-TCP, RTSP, ARP, CITRIX-AG, CITRIX-AAC-LOGINPAGE, CITRIX-AAC-LAS, CITRIX-XD-DDC, ND6, CITRIX-WI-EXTENDED, DIAMETER, RADIUS_ACCOUNTING, STOREFRONT, APPC, SMPP, CITRIX-XNC-ECV, CITRIX-XDM, CITRIX-STA-SERVICE, CITRIX-STA-SERVICE-NHOP
// units1	<String>	Read-write	Unit of measurement for the Deviation parameter. Cannot be changed after the monitor is created.<br>Default value: SEC<br>Possible values = SEC, MSEC, MIN
// units2	<String>	Read-write	Unit of measurement for the Down Time parameter. Cannot be changed after the monitor is created.<br>Default value: SEC<br>Possible values = SEC, MSEC, MIN
// units3	<String>	Read-write	monitor interval units.<br>Default value: SEC<br>Possible values = SEC, MSEC, MIN
// units4	<String>	Read-write	monitor response timeout units.<br>Default value: SEC<br>Possible values = SEC, MSEC, MIN
// username	<String>	Read-write	User name with which to probe the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC or CITRIX-XDM server.<br>Minimum length = 1
// validatecred	<String>	Read-write	Validate the credentials of the Xen Desktop DDC server user. Applicable to monitors of type CITRIX-XD-DDC.<br>Default value: NO<br>Possible values = YES, NO
// vendorid	<Double>	Read-write	Vendor-Id value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
// vendorspecificacctapplicationids	<Double[]>	Read-write	List of Vendor-Specific-Acct-Application-Id attribute value pairs (AVPs) to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.<br>Minimum value = 0<br>Maximum value = 4294967295
// vendorspecificauthapplicationids	<Double[]>	Read-write	List of Vendor-Specific-Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.<br>Minimum value = 0<br>Maximum value = 4294967295
// vendorspecificvendorid	<Double>	Read-write	Vendor-Id to use in the Vendor-Specific-Application-Id grouped attribute-value pair (AVP) in the monitoring CER message. To specify Auth-Application-Id or Acct-Application-Id in Vendor-Specific-Application-Id, use vendorSpecificAuthApplicationIds or vendorSpecificAcctApplicationIds, respectively. Only one Vendor-Id is supported for all the Vendor-Specific-Application-Id AVPs in a CER monitoring message.<br>Minimum value = 1
// weight	<Double>	Read-only	.<br>Minimum value = 1<br>Maximum value = 100

// LbmonitorAdd defines add request
type LbmonitorAdd struct {
	Lbmonitor LbmonitorAddBody `json:"lbmonitor,omitempty"`
}

// LbmonitorAddBody body for adding object.
type LbmonitorAddBody struct {
	Monitorname                      string   `json:"monitorname,omitempty"`
	Type                             string   `json:"type,omitempty"`
	Action                           string   `json:"action,omitempty"`
	Respcode                         []string `json:"respcode,omitempty"`
	Httprequest                      string   `json:"httprequest,omitempty"`
	Rtsprequest                      string   `json:"rtsprequest,omitempty"`
	Customheaders                    string   `json:"customheaders,omitempty"`
	Maxforwards                      string   `json:"maxforwards,omitempty"`
	Sipmethod                        string   `json:"sipmethod,omitempty"`
	Sipuri                           string   `json:"sipuri,omitempty"`
	Sipreguri                        string   `json:"sipreguri,omitempty"`
	Send                             string   `json:"send,omitempty"`
	Recv                             string   `json:"recv,omitempty"`
	Query                            string   `json:"query,omitempty"`
	Querytype                        string   `json:"querytype,omitempty"`
	Scriptname                       string   `json:"scriptname,omitempty"`
	Scriptargs                       string   `json:"scriptargs,omitempty"`
	Dispatcherip                     string   `json:"dispatcherip,omitempty"`
	Dispatcherport                   int      `json:"dispatcherport,omitempty"`
	Username                         string   `json:"username,omitempty"`
	Password                         string   `json:"password,omitempty"`
	Secondarypassword                string   `json:"secondarypassword,omitempty"`
	Logonpointname                   string   `json:"logonpointname,omitempty"`
	Lasversion                       string   `json:"lasversion,omitempty"`
	Radkey                           string   `json:"radkey,omitempty"`
	Radnasid                         string   `json:"radnasid,omitempty"`
	Radnasip                         string   `json:"radnasip,omitempty"`
	Radaccounttype                   string   `json:"radaccounttype,omitempty"`
	Radframedip                      string   `json:"radframedip,omitempty"`
	Radapn                           string   `json:"radapn,omitempty"`
	Radmsisdn                        string   `json:"radmsisdn,omitempty"`
	Radaccountsession                string   `json:"radaccountsession,omitempty"`
	Lrtm                             string   `json:"lrtm,omitempty"`
	Deviation                        string   `json:"deviation,omitempty"`
	Units1                           string   `json:"units1,omitempty"`
	Interval                         int      `json:"interval,omitempty"`
	Units3                           string   `json:"units3,omitempty"`
	Resptimeout                      int      `json:"resptimeout,omitempty"`
	Units4                           string   `json:"units4,omitempty"`
	Resptimeoutthresh                string   `json:"resptimeoutthresh,omitempty"`
	Retries                          int      `json:"retries,omitempty"`
	Failureretries                   int      `json:"failureretries,omitempty"`
	Alertretries                     int      `json:"alertretries,omitempty"`
	Successretries                   int      `json:"successretries,omitempty"`
	Downtime                         int      `json:"downtime,omitempty"`
	Units2                           string   `json:"units2,omitempty"`
	Destip                           string   `json:"destip,omitempty"`
	Destport                         int      `json:"destport,omitempty"`
	State                            string   `json:"state,omitempty"`
	Reverse                          string   `json:"reverse,omitempty"`
	Transparent                      string   `json:"transparent,omitempty"`
	Iptunnel                         string   `json:"iptunnel,omitempty"`
	Tos                              string   `json:"tos,omitempty"`
	Tosid                            string   `json:"tosid,omitempty"`
	Secure                           string   `json:"secure,omitempty"`
	Validatecred                     string   `json:"validatecred,omitempty"`
	Domain                           string   `json:"domain,omitempty"`
	Ipaddress                        []string `json:"ipaddress,omitempty"`
	Group                            string   `json:"group,omitempty"`
	Filename                         string   `json:"filename,omitempty"`
	Basedn                           string   `json:"basedn,omitempty"`
	Binddn                           string   `json:"binddn,omitempty"`
	Filter                           string   `json:"filter,omitempty"`
	Attribute                        string   `json:"attribute,omitempty"`
	Database                         string   `json:"database,omitempty"`
	Oraclesid                        string   `json:"oraclesid,omitempty"`
	Sqlquery                         string   `json:"sqlquery,omitempty"`
	Evalrule                         string   `json:"evalrule,omitempty"`
	Mssqlprotocolversion             string   `json:"mssqlprotocolversion,omitempty"`
	Snmpoid                          string   `json:"Snmpoid,omitempty"`
	Snmpcommunity                    string   `json:"snmpcommunity,omitempty"`
	Snmpthreshold                    string   `json:"snmpthreshold,omitempty"`
	Snmpversion                      string   `json:"snmpversion,omitempty"`
	Metrictable                      string   `json:"metrictable,omitempty"`
	Application                      string   `json:"application,omitempty"`
	Sitepath                         string   `json:"sitepath,omitempty"`
	Storename                        string   `json:"storename,omitempty"`
	Storefrontacctservice            string   `json:"storefrontacctservice,omitempty"`
	Hostname                         string   `json:"hostname,omitempty"`
	Netprofile                       string   `json:"netprofile,omitempty"`
	Originhost                       string   `json:"originhost,omitempty"`
	Originrealm                      string   `json:"originrealm,omitempty"`
	Hostipaddress                    string   `json:"hostipaddress,omitempty"`
	Vendorid                         string   `json:"vendorid,omitempty"`
	Productname                      string   `json:"productname,omitempty"`
	Firmwarerevision                 string   `json:"firmwarerevision,omitempty"`
	Authapplicationid                []string `json:"authapplicationid,omitempty"`
	Acctapplicationid                []string `json:"acctapplicationid,omitempty"`
	Inbandsecurityid                 string   `json:"inbandsecurityid,omitempty"`
	Supportedvendorids               []string `json:"supportedvendorids,omitempty"`
	Vendorspecificvendorid           string   `json:"vendorspecificvendorid,omitempty"`
	Vendorspecificauthapplicationids []string `json:"vendorspecificauthapplicationids,omitempty"`
	Vendorspecificacctapplicationids []string `json:"vendorspecificacctapplicationids,omitempty"`
	Kcdaccount                       string   `json:"kcdaccount,omitempty"`
	Storedb                          string   `json:"storedb,omitempty"`
	Storefrontcheckbackendservices   string   `json:"storefrontcheckbackendservices,omitempty"`
	Trofscode                        string   `json:"trofscode,omitempty"`
	Trofsstring                      string   `json:"trofsstring,omitempty"`
	Sslprofile                       string   `json:"sslprofile,omitempty"`
}

// LbmonitorUpdate defines update request.
type LbmonitorUpdate struct {
	Lbmonitor LbmonitorUpdateBody `json:"lbmonitor,omitempty"`
}

// LbmonitorUpdateBody body for updating object.
type LbmonitorUpdateBody struct {
	Monitorname                      string   `json:"monitorname,omitempty"`
	Type                             string   `json:"type,omitempty"`
	Action                           string   `json:"action,omitempty"`
	Respcode                         []string `json:"respcode,omitempty"`
	Httprequest                      string   `json:"httprequest,omitempty"`
	Rtsprequest                      string   `json:"rtsprequest,omitempty"`
	Customheaders                    string   `json:"customheaders,omitempty"`
	Maxforwards                      string   `json:"maxforwards,omitempty"`
	Sipmethod                        string   `json:"sipmethod,omitempty"`
	Sipreguri                        string   `json:"sipreguri,omitempty"`
	Sipuri                           string   `json:"sipuri,omitempty"`
	Send                             string   `json:"send,omitempty"`
	Recv                             string   `json:"recv,omitempty"`
	Query                            string   `json:"query,omitempty"`
	Querytype                        string   `json:"querytype,omitempty"`
	Username                         string   `json:"username,omitempty"`
	Password                         string   `json:"password,omitempty"`
	Secondarypassword                string   `json:"secondarypassword,omitempty"`
	Logonpointname                   string   `json:"logonpointname,omitempty"`
	Lasversion                       string   `json:"lasversion,omitempty"`
	Radkey                           string   `json:"radkey,omitempty"`
	Radnasid                         string   `json:"radnasid,omitempty"`
	Radnasip                         string   `json:"radnasip,omitempty"`
	Radaccounttype                   string   `json:"radaccounttype,omitempty"`
	Radframedip                      string   `json:"radframedip,omitempty"`
	Radapn                           string   `json:"radapn,omitempty"`
	Radmsisdn                        string   `json:"radmsisdn,omitempty"`
	Radaccountsession                string   `json:"radaccountsession,omitempty"`
	Lrtm                             string   `json:"lrtm,omitempty"`
	Deviation                        string   `json:"deviation,omitempty"`
	Units1                           string   `json:"units1,omitempty"`
	Scriptname                       string   `json:"scriptname,omitempty"`
	Scriptargs                       string   `json:"scriptargs,omitempty"`
	Validatecred                     string   `json:"validatecred,omitempty"`
	Domain                           string   `json:"domain,omitempty"`
	Dispatcherip                     string   `json:"dispatcherip,omitempty"`
	Dispatcherport                   int      `json:"dispatcherport,omitempty"`
	Interval                         int      `json:"interval,omitempty"`
	Units3                           string   `json:"units3,omitempty"`
	Resptimeout                      int      `json:"resptimeout,omitempty"`
	Units4                           string   `json:"units4,omitempty"`
	Resptimeoutthresh                string   `json:"resptimeoutthresh,omitempty"`
	Retries                          int      `json:"retries,omitempty"`
	Failureretries                   int      `json:"failureretries,omitempty"`
	Alertretries                     int      `json:"alertretries,omitempty"`
	Successretries                   int      `json:"successretries,omitempty"`
	Downtime                         int      `json:"downtime,omitempty"`
	Units2                           string   `json:"units2,omitempty"`
	Destip                           string   `json:"destip,omitempty"`
	Destport                         int      `json:"destport,omitempty"`
	State                            string   `json:"state,omitempty"`
	Reverse                          string   `json:"reverse,omitempty"`
	Transparent                      string   `json:"transparent,omitempty"`
	Iptunnel                         string   `json:"iptunnel,omitempty"`
	Tos                              string   `json:"tos,omitempty"`
	Tosid                            string   `json:"tosid,omitempty"`
	Secure                           string   `json:"secure,omitempty"`
	Ipaddress                        []string `json:"ipaddress,omitempty"`
	Group                            string   `json:"group,omitempty"`
	Filename                         string   `json:"filename,omitempty"`
	Basedn                           string   `json:"basedn,omitempty"`
	Binddn                           string   `json:"binddn,omitempty"`
	Filter                           string   `json:"filter,omitempty"`
	Attribute                        string   `json:"attribute,omitempty"`
	Database                         string   `json:"database,omitempty"`
	Oraclesid                        string   `json:"oraclesid,omitempty"`
	Sqlquery                         string   `json:"sqlquery,omitempty"`
	Evalrule                         string   `json:"evalrule,omitempty"`
	Snmpoid                          string   `json:"Snmpoid,omitempty"`
	Snmpcommunity                    string   `json:"snmpcommunity,omitempty"`
	Snmpthreshold                    string   `json:"snmpthreshold,omitempty"`
	Snmpversion                      string   `json:"snmpversion,omitempty"`
	Metrictable                      string   `json:"metrictable,omitempty"`
	Metric                           string   `json:"metric,omitempty"`
	Metricthreshold                  string   `json:"metricthreshold,omitempty"`
	Metricweight                     string   `json:"metricweight,omitempty"`
	Application                      string   `json:"application,omitempty"`
	Sitepath                         string   `json:"sitepath,omitempty"`
	Storename                        string   `json:"storename,omitempty"`
	Storefrontacctservice            string   `json:"storefrontacctservice,omitempty"`
	Storefrontcheckbackendservices   string   `json:"storefrontcheckbackendservices,omitempty"`
	Hostname                         string   `json:"hostname,omitempty"`
	Netprofile                       string   `json:"netprofile,omitempty"`
	Mssqlprotocolversion             string   `json:"mssqlprotocolversion,omitempty"`
	Originhost                       string   `json:"originhost,omitempty"`
	Originrealm                      string   `json:"originrealm,omitempty"`
	Hostipaddress                    string   `json:"hostipaddress,omitempty"`
	Vendorid                         string   `json:"vendorid,omitempty"`
	Productname                      string   `json:"productname,omitempty"`
	Firmwarerevision                 string   `json:"firmwarerevision,omitempty"`
	Authapplicationid                []string `json:"authapplicationid,omitempty"`
	Acctapplicationid                []string `json:"acctapplicationid,omitempty"`
	Inbandsecurityid                 string   `json:"inbandsecurityid,omitempty"`
	Supportedvendorids               []string `json:"supportedvendorids,omitempty"`
	Vendorspecificvendorid           string   `json:"vendorspecificvendorid,omitempty"`
	Vendorspecificauthapplicationids []string `json:"vendorspecificauthapplicationids,omitempty"`
	Vendorspecificacctapplicationids []string `json:"vendorspecificacctapplicationids,omitempty"`
	Kcdaccount                       string   `json:"kcdaccount,omitempty"`
	Storedb                          string   `json:"storedb,omitempty"`
	Trofscode                        string   `json:"trofscode,omitempty"`
	Trofsstring                      string   `json:"trofsstring,omitempty"`
	Sslprofile                       string   `json:"sslprofile,omitempty"`
}

// LbmonitorEnable defines enable request.
type LbmonitorEnable struct {
	Lbmonitor LbmonitorEnableDisableBody `json:"lbmonitor,omitempty"`
}

// LbmonitorDisable defines disable request.
type LbmonitorDisable struct {
	Lbmonitor LbmonitorEnableDisableBody `json:"lbmonitor,omitempty"`
}

// LbmonitorEnableDisableBody body for enabling/disabling object.
type LbmonitorEnableDisableBody struct {
	Name             string `json:"name,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
}

// LbmonitorRename defines rename request.
type LbmonitorRename struct {
	Lbmonitor LbmonitorRenameBody `json:"lbmonitor,omitempty"`
}

// LbmonitorRenameBody body for renaming object.
type LbmonitorRenameBody struct {
	Name    string `json:"name,omitempty"`
	Newname string `json:"newname,omitempty"`
}

// LbmonitorWrapper wraps the object and serves as default response.
type LbmonitorWrapper struct {
	Errorcode int         `json:"errorcode,omitempty"`
	Message   string      `json:"message,omitempty"`
	Severity  string      `json:"severity,omitempty"`
	Lbmonitor []Lbmonitor `json:"lbmonitor,omitempty"`
}

// Lbmonitor describes the object.
type Lbmonitor struct {
	Monitorname                      string   `json:"monitorname,omitempty"`
	Type                             string   `json:"type,omitempty"`
	Interval                         int      `json:"interval,omitempty"`
	Units3                           string   `json:"units3,omitempty"`
	Resptimeout                      int      `json:"resptimeout,omitempty"`
	Resptimeoutthresh                string   `json:"resptimeoutthresh,omitempty"`
	Units4                           string   `json:"units4,omitempty"`
	Retries                          int      `json:"retries,omitempty"`
	Failureretries                   int      `json:"failureretries,omitempty"`
	Alertretries                     int      `json:"alertretries,omitempty"`
	Successretries                   int      `json:"successretries,omitempty"`
	Downtime                         int      `json:"downtime,omitempty"`
	Units2                           string   `json:"units2,omitempty"`
	Destip                           string   `json:"destip,omitempty"`
	Destport                         int      `json:"destport,omitempty"`
	State                            string   `json:"state,omitempty"`
	Reverse                          string   `json:"reverse,omitempty"`
	Transparent                      string   `json:"transparent,omitempty"`
	Iptunnel                         string   `json:"iptunnel,omitempty"`
	Tos                              string   `json:"tos,omitempty"`
	Tosid                            string   `json:"tosid,omitempty"`
	Secure                           string   `json:"secure,omitempty"`
	Action                           string   `json:"action,omitempty"`
	Respcode                         []string `json:"respcode,omitempty"`
	Httprequest                      string   `json:"httprequest,omitempty"`
	Rtsprequest                      string   `json:"rtsprequest,omitempty"`
	Send                             string   `json:"send,omitempty"`
	Recv                             string   `json:"recv,omitempty"`
	Query                            string   `json:"query,omitempty"`
	Querytype                        string   `json:"querytype,omitempty"`
	Username                         string   `json:"username,omitempty"`
	Password                         string   `json:"password,omitempty"`
	Secondarypassword                string   `json:"secondarypassword,omitempty"`
	Logonpointname                   string   `json:"logonpointname,omitempty"`
	Lasversion                       string   `json:"lasversion,omitempty"`
	Validatecred                     string   `json:"validatecred,omitempty"`
	Domain                           string   `json:"domain,omitempty"`
	Radkey                           string   `json:"radkey,omitempty"`
	Radnasid                         string   `json:"radnasid,omitempty"`
	Radnasip                         string   `json:"radnasip,omitempty"`
	Radaccounttype                   string   `json:"radaccounttype,omitempty"`
	Radframedip                      string   `json:"radframedip,omitempty"`
	Radapn                           string   `json:"radapn,omitempty"`
	Radmsisdn                        string   `json:"radmsisdn,omitempty"`
	Radaccountsession                string   `json:"radaccountsession,omitempty"`
	Lrtm                             string   `json:"lrtm,omitempty"`
	Lrtmconf                         int      `json:"lrtmconf,omitempty"`
	Lrtmconfstr                      string   `json:"lrtmconfstr,omitempty"`
	Deviation                        string   `json:"deviation,omitempty"`
	Units1                           string   `json:"units1,omitempty"`
	Dynamicresponsetimeout           int      `json:"dynamicresponsetimeout,omitempty"`
	Dynamicinterval                  int      `json:"dynamicinterval,omitempty"`
	Scriptname                       string   `json:"scriptname,omitempty"`
	Scriptargs                       string   `json:"scriptargs,omitempty"`
	Dispatcherip                     string   `json:"dispatcherip,omitempty"`
	Dispatcherport                   int      `json:"dispatcherport,omitempty"`
	Sipuri                           string   `json:"sipuri,omitempty"`
	Sipmethod                        string   `json:"sipmethod,omitempty"`
	Maxforwards                      string   `json:"maxforwards,omitempty"`
	Sipreguri                        string   `json:"sipreguri,omitempty"`
	Customheaders                    string   `json:"customheaders,omitempty"`
	Ipaddress                        []string `json:"ipaddress,omitempty"`
	Group                            string   `json:"group,omitempty"`
	Filename                         string   `json:"filename,omitempty"`
	Basedn                           string   `json:"basedn,omitempty"`
	Binddn                           string   `json:"binddn,omitempty"`
	Filter                           string   `json:"filter,omitempty"`
	Attribute                        string   `json:"attribute,omitempty"`
	Database                         string   `json:"database,omitempty"`
	Oraclesid                        string   `json:"oraclesid,omitempty"`
	Sqlquery                         string   `json:"sqlquery,omitempty"`
	Evalrule                         string   `json:"evalrule,omitempty"`
	Snmpoid                          string   `json:"Snmpoid,omitempty"`
	Snmpcommunity                    string   `json:"snmpcommunity,omitempty"`
	Snmpthreshold                    string   `json:"snmpthreshold,omitempty"`
	Snmpversion                      string   `json:"snmpversion,omitempty"`
	Metric                           string   `json:"metric,omitempty"`
	Metrictable                      string   `json:"metrictable,omitempty"`
	Multimetrictable                 []string `json:"multimetrictable,omitempty"`
	Metricthreshold                  string   `json:"metricthreshold,omitempty"`
	Metricweight                     string   `json:"metricweight,omitempty"`
	Application                      string   `json:"application,omitempty"`
	Sitepath                         string   `json:"sitepath,omitempty"`
	Storename                        string   `json:"storename,omitempty"`
	Storefrontacctservice            string   `json:"storefrontacctservice,omitempty"`
	Storefrontcheckbackendservices   string   `json:"storefrontcheckbackendservices,omitempty"`
	Hostname                         string   `json:"hostname,omitempty"`
	Netprofile                       string   `json:"netprofile,omitempty"`
	Mssqlprotocolversion             string   `json:"mssqlprotocolversion,omitempty"`
	Originhost                       string   `json:"originhost,omitempty"`
	Originrealm                      string   `json:"originrealm,omitempty"`
	Hostipaddress                    string   `json:"hostipaddress,omitempty"`
	Vendorid                         string   `json:"vendorid,omitempty"`
	Productname                      string   `json:"productname,omitempty"`
	Firmwarerevision                 string   `json:"firmwarerevision,omitempty"`
	Authapplicationid                []string `json:"authapplicationid,omitempty"`
	Acctapplicationid                []string `json:"acctapplicationid,omitempty"`
	Inbandsecurityid                 string   `json:"inbandsecurityid,omitempty"`
	Supportedvendorids               []string `json:"supportedvendorids,omitempty"`
	Vendorspecificvendorid           string   `json:"vendorspecificvendorid,omitempty"`
	Vendorspecificauthapplicationids []string `json:"vendorspecificauthapplicationids,omitempty"`
	Vendorspecificacctapplicationids []string `json:"vendorspecificacctapplicationids,omitempty"`
	Servicename                      string   `json:"servicename,omitempty"`
	DupState                         string   `json:"dup_state,omitempty"`
	DupWeight                        string   `json:"dup_weight,omitempty"`
	Servicegroupname                 string   `json:"servicegroupname,omitempty"`
	Weight                           string   `json:"weight,omitempty"`
	Kcdaccount                       string   `json:"kcdaccount,omitempty"`
	Storedb                          string   `json:"storedb,omitempty"`
	Trofscode                        string   `json:"trofscode,omitempty"`
	Trofsstring                      string   `json:"trofsstring,omitempty"`
	Sslprofile                       string   `json:"sslprofile,omitempty"`
}
