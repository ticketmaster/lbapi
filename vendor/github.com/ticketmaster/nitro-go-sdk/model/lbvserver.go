package model

// Activeservices. Total number of active services bound to the vserver.
// Appflowlog. Apply AppFlow logging to the virtual server.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Authentication. Enable or disable user authentication.<br>Default value: OFF<br>Possible values = ON, OFF.
// Authenticationhost. Fully qualified domain name (FQDN) of the authentication virtual server to which the user must be redirected for authentication. Make sure that the Authentication parameter is set to ENABLED.<br>Minimum length = 3<br>Maximum length = 252.
// Authn401. Enable or disable user authentication with HTTP 401 responses.<br>Default value: OFF<br>Possible values = ON, OFF.
// Authnprofile. Name of the authentication profile to be used when authentication is turned on.
// Authnvsname. Name of an authentication virtual server with which to authenticate users.<br>Minimum length = 1<br>Maximum length = 252.
// Backuplbmethod. Backup load balancing method. Becomes operational if the primary load balancing me<br>thod fails or cannot be used.<br> Valid only if the primary method is based on static proximity.<br>Default value: ROUNDROBIN<br>Possible values = ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME, SOURCEIPHASH, LEASTBANDWIDTH, LEASTPACKETS, CUSTOMLOAD.
// Backuppersistencetimeout. Time period for which backup persistence is in effect.<br>Default value: 2<br>Minimum value = 2<br>Maximum value = 1440.
// Backupvserver. Name of the backup virtual server to which to forward requests if the primary virtual server goes DOWN or reaches its spillover threshold.<br>Minimum length = 1.
// Backupvserverstatus. Staus of BackUp Vserver .
// Bindpoint. The bindpoint to which the policy is bound.<br>Possible values = REQUEST, RESPONSE.
// Bypassaaaa. If this option is enabled while resolving DNS64 query AAAA queries are not sent to back end dns server.<br>Default value: NO<br>Possible values = YES, NO.
// Cacheable. Route cacheable requests to a cache redirection virtual server. The load balancing virtual server can forward requests only to a transparent cache redirection virtual server that has an IP address and port combination of *:80, so such a cache redirection virtual server must be configured on the appliance.<br>Default value: NO<br>Possible values = YES, NO.
// Cachevserver. Cache virtual server.
// Clttimeout. Idle time, in seconds, after which a client connection is terminated.<br>Minimum value = 0<br>Maximum value = 31536000.
// Comment. Any comments that you might want to associate with the virtual server.
// Connfailover. Mode in which the connection failover feature must operate for the virtual server. After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary appliance. Clients remain connected to the same servers. Available settings function as follows:<br>* STATEFUL - The primary appliance shares state information with the secondary appliance, in real time, resulting in some runtime processing overhead. <br>* STATELESS - State information is not shared, and the new primary appliance tries to re-create the packet flow on the basis of the information contained in the packets it receives. <br>* DISABLED - Connection failover does not occur.<br>Default value: DISABLED<br>Possible values = DISABLED, STATEFUL, STATELESS.
// Consolidatedlconn. Use consolidated stats for LeastConnection.<br>Default value: GLOBAL<br>Possible values = GLOBAL, NO, YES.
// Consolidatedlconngbl. Fetches Global setting.<br>Possible values = YES, NO.
// Cookiedomain. Domain name to be used in the set cookie header in case of cookie persistence.
// Cookiename. Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.
// Curstate. Current LB vserver state.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
// Datalength. Length of the token to be extracted from the data segment of an incoming packet, for use in the token method of load balancing. The length of the token, specified in bytes, must not be greater than 24 KB. Applicable to virtual servers of type TCP.<br>Minimum value = 1<br>Maximum value = 100.
// Dataoffset. Offset to be considered when extracting a token from the TCP payload. Applicable to virtual servers, of type TCP, using the token method of load balancing. Must be within the first 24 KB of the TCP payload.<br>Minimum value = 0<br>Maximum value = 25400.
// Dbprofilename. Name of the DB profile whose settings are to be applied to the virtual server.<br>Minimum length = 1<br>Maximum length = 127.
// Dbslb. Enable database specific load balancing for MySQL and MSSQL service types.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Disableprimaryondown. If the primary virtual server goes down, do not allow it to return to primary status until manually enabled.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Dns64. This argument is for enabling/disabling the dns64 on lbvserver.<br>Possible values = ENABLED, DISABLED.
// Dnsprofilename. Name of the DNS profile to be associated with the VServer. DNS profile properties will be applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.<br>Minimum length = 1<br>Maximum length = 127.
// Dnsvservername. DNS vserver name.
// Domain. Domain.
// Downstateflush. Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Effectivestate. Effective state of the LB vserver , based on the state of backup vservers.<br>Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
// Gotopriorityexpression. Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
// Groupname. LB group to which the lb vserver is to be bound.
// Gt2Gb. Allow for greater than 2 GB transactions on this vserver.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Hashlength. Number of bytes to consider for the hash value used in the URLHASH and DOMAINHASH load balancing methods.<br>Default value: 80<br>Minimum value = 1<br>Maximum value = 4096.
// Health. Health of vserver based on percentage of weights of active svcs/all svcs. This does not consider administratively disabled svcs.
// Healththreshold. Threshold in percent of active services below which vserver state is made down. If this threshold is 0, vserver state will be up even if one bound service is up.<br>Default value: 0<br>Minimum value = 0<br>Maximum value = 100.
// Homepage. Home page.
// Httpprofilename. Name of the HTTP profile whose settings are to be applied to the virtual server.<br>Minimum length = 1<br>Maximum length = 127.
// Httpsredirecturl. URL to which to redirect traffic if the traffic is recieved from redirect port.
// Icmpvsrresponse. How the NetScaler appliance responds to ping requests received for an IP address that is common to one or more virtual servers. Available settings function as follows:<br>* If set to PASSIVE on all the virtual servers that share the IP address, the appliance always responds to the ping requests.<br>* If set to ACTIVE on all the virtual servers that share the IP address, the appliance responds to the ping requests if at least one of the virtual servers is UP. Otherwise, the appliance does not respond.<br>* If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance responds if at least one virtual server with the ACTIVE setting is UP. Otherwise, the appliance does not respond.<br>Note: This parameter is available at the virtual server level. A similar parameter, ICMP Response, is available at the IP address level, for IPv4 addresses of type VIP. To set that parameter, use the add ip command in the CLI or the Create IP dialog box in the GUI.<br>Default value: PASSIVE<br>Possible values = PASSIVE, ACTIVE.
// Insertvserveripport. Insert an HTTP header, whose value is the IP address and port number of the virtual server, before forwarding a request to the server. The format of the header is ;lt;vipHeader;gt;: ;lt;virtual server IP address;gt;_;lt;port number ;gt;, where vipHeader is the name that you specify for the header. If the virtual server has an IPv6 address, the address in the header is enclosed in brackets ([ and ]) to separate it from the port number. If you have mapped an IPv4 address to a virtual servers IPv6 address, the value of this parameter determines which IP address is inserted in the header, as follows:<br>* VIPADDR - Insert the IP address of the virtual server in the HTTP header regardless of whether the virtual server has an IPv4 address or an IPv6 address. A mapped IPv4 address, if configured, is ignored.<br>* V6TOV4MAPPING - Insert the IPv4 address that is mapped to the virtual servers IPv6 address. If a mapped IPv4 address is not configured, insert the IPv6 address.<br>* OFF - Disable header insertion.<br>Possible values = OFF, VIPADDR, V6TOV4MAPPING.
// Invoke. Invoke policies bound to a virtual server or policy label.
// Ipmapping. The permanent mapping for the V6 Address.
// Ipmask. IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.
// Ippattern. IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern. Mutually exclusive with the IP Address parameter. <br>For example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254. You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask).<br>If a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if virtual servers vs1 and vs2 have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.
// Ipv46. IPv4 or IPv6 address to assign to the virtual server.
// Isgslb. This field is set to true if it is a GSLBVserver.
// L2Conn. Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple (;lt;source IP;gt;:;lt;source port;gt;::;lt;destination IP;gt;:;lt;destination port;gt;) that is used to identify a connection. Allows multiple TCP and non-TCP connections with the same 4-tuple to co-exist on the NetScaler appliance.<br>Possible values = ON, OFF.
// Labelname. Name of the label invoked.
// Labeltype. The invocation type.<br>Possible values = reqvserver, resvserver, policylabel.
// Lbmethod. Load balancing method. The available settings function as follows:<br>* ROUNDROBIN - Distribute requests in rotation, regardless of the load. Weights can be assigned to services to enforce weighted round robin distribution.<br>* LEASTCONNECTION (default) - Select the service with the fewest connections. <br>* LEASTRESPONSETIME - Select the service with the lowest average response time. <br>* LEASTBANDWIDTH - Select the service currently handling the least traffic.<br>* LEASTPACKETS - Select the service currently serving the lowest number of packets per second.<br>* CUSTOMLOAD - Base service selection on the SNMP metrics obtained by custom load monitors.<br>* LRTM - Select the service with the lowest response time. Response times are learned through monitoring probes. This method also takes the number of active connections into account.<br>Also available are a number of hashing methods, in which the appliance extracts a predetermined portion of the request, creates a hash of the portion, and then checks whether any previous requests had the same hash value. If it finds a match, it forwards the request to the service that served those previous requests. Following are the hashing methods: <br>* URLHASH - Create a hash of the request URL (or part of the URL).<br>* DOMAINHASH - Create a hash of the domain name in the request (or part of the domain name). The domain name is taken from either the URL or the Host header. If the domain name appears in both locations, the URL is preferred. If the request does not contain a domain name, the load balancing method defaults to LEASTCONNECTION.<br>* DESTINATIONIPHASH - Create a hash of the destination IP address in the IP header. <br>* SOURCEIPHASH - Create a hash of the source IP address in the IP header. <br>* TOKEN - Extract a token from the request, create a hash of the token, and then select the service to which any previous requests with the same token hash value were sent. <br>* SRCIPDESTIPHASH - Create a hash of the string obtained by concatenating the source IP address and destination IP address in the IP header. <br>* SRCIPSRCPORTHASH - Create a hash of the source IP address and source port in the IP header. <br>* CALLIDHASH - Create a hash of the SIP Call-ID header.<br>* USER_TOKEN - Same as TOKEN LB method but token needs to be provided from an extension.<br>Default value: LEASTCONNECTION<br>Possible values = ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME, URLHASH, DOMAINHASH, DESTINATIONIPHASH, SOURCEIPHASH, SRCIPDESTIPHASH, LEASTBANDWIDTH, LEASTPACKETS, TOKEN, SRCIPSRCPORTHASH, LRTM, CALLIDHASH, CUSTOMLOAD, LEASTREQUEST, AUDITLOGHASH, STATICPROXIMITY, USER_TOKEN.
// Lbprofilename. Name of the LB profile which is associated to the vserver.
// Lbrrreason. Reason why a vserver is in RR. The following are the reasons:<br>1 - MEP is DOWN (GSLB)<br>2 - LB method has changed<br>3 - Bound services state changed to UP<br>4 - A new service is bound<br>5 - Startup RR factor has changed<br>6 - LB feature is enabled<br>7 - Load monitor is not active on a service<br>8 - Vserver is Enabled<br>9 - SSL feature is Enabled<br>10 - All bound services have reached threshold. Using effective state to load balance (GSLB)<br>11 - Primary state of bound services are not UP. Using effective state to load balance (GSLB)<br>12 - No LB decision can be made as all bound services have either reached threshold or are not UP (GSLB)<br>13 - All load monitors are active<br>.
// Listenpolicy. Default syntax expression identifying traffic accepted by the virtual server. Can be either an expression (for example, CLIENT.IP.DST.IN_SUBNET(192.0.2.0/24) or the name of a named expression. In the above example, the virtual server accepts all requests whose destination IP address is in the 192.0.2.0/24 subnet.<br>Default value: "NONE".
// Listenpriority. Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.<br>Default value: 101<br>Minimum value = 0<br>Maximum value = 101.
// M. Redirection mode for load balancing. Available settings function as follows:<br>* IP - Before forwarding a request to a server, change the destination IP address to the servers IP address. <br>* MAC - Before forwarding a request to a server, change the destination MAC address to the servers MAC address. The destination IP address is not changed. MAC-based redirection mode is used mostly in firewall load balancing deployments. <br>* IPTUNNEL - Perform IP-in-IP encapsulation for client IP packets. In the outer IP headers, set the destination IP address to the IP address of the server and the source IP address to the subnet IP (SNIP). The client IP packets are not modified. Applicable to both IPv4 and IPv6 packets. <br>* TOS - Encode the virtual servers TOS ID in the TOS field of the IP header. <br>You can use either the IPTUNNEL or the TOS option to implement Direct Server Return (DSR).<br>Default value: IP<br>Possible values = IP, MAC, IPTUNNEL, TOS.
// Macmoderetainvlan. This option is used to retain vlan information of incoming packet when macmode is enabled.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Map. Map.<br>Possible values = ON, OFF.
// Maxautoscalemembers. Maximum number of members expected to be present when vserver is used in Autoscale.<br>Default value: 0<br>Minimum value = 0<br>Maximum value = 5000.
// Minautoscalemembers. Minimum number of members expected to be present when vserver is used in Autoscale.<br>Default value: 0<br>Minimum value = 0<br>Maximum value = 5000.
// Mssqlserverversion. For a load balancing virtual server of type MSSQL, the Microsoft SQL Server version. Set this parameter if you expect some clients to run a version different from the version of the database. This setting provides compatibility between the client-side and server-side connections by ensuring that all communication conforms to the servers version.<br>Default value: 2008R2<br>Possible values = 70, 2000, 2000SP1, 2005, 2008, 2008R2, 2012, 2014.
// Mysqlcharacterset. Character set that the virtual server advertises to clients.
// Mysqlprotocolversion. MySQL protocol version that the virtual server advertises to clients.
// Mysqlservercapabilities. Server capabilities that the virtual server advertises to clients.
// Mysqlserverversion. MySQL server version string that the virtual server advertises to clients.<br>Minimum length = 1<br>Maximum length = 31.
// Name. Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.<br><br>CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or my vserver). .<br>Minimum length = 1.
// Netmask. IPv4 subnet mask to apply to the destination IP address or source IP address when the load balancing method is DESTINATIONIPHASH or SOURCEIPHASH.<br>Minimum length = 1.
// Netprofile. Name of the network profile to associate with the virtual server. If you set this parameter, the virtual server uses only the IP addresses in the network profile as source IP addresses when initiating connections with servers.<br>Minimum length = 1<br>Maximum length = 127.
// Newname. New name for the virtual server.<br>Minimum length = 1.
// Newservicerequest. Number of requests, or percentage of the load on existing services, by which to increase the load on a new service at each interval in slow-start mode. A non-zero value indicates that slow-start is applicable. A zero value indicates that the global RR startup parameter is applied. Changing the value to zero will cause services currently in slow start to take the full traffic as determined by the LB method. Subsequently, any new services added will use the global RR factor.<br>Default value: 0.
// Newservicerequestincrementinterval. Interval, in seconds, between successive increments in the load on a new service or a service whose state has just changed from DOWN to UP. A value of 0 (zero) specifies manual slow start.<br>Default value: 0<br>Minimum value = 0<br>Maximum value = 3600.
// Newservicerequestunit. Units in which to increment load at each interval in slow-start mode.<br>Default value: PER_SECOND<br>Possible values = PER_SECOND, PERCENT.
// Ngname. Nodegroup name to which this lbvsever belongs to.
// Oracleserverversion. Oracle server version.<br>Default value: 10G<br>Possible values = 10G, 11G.
// Persistavpno. Persist AVP number for Diameter Persistency. <br> In case this AVP is not defined in Base RFC 3588 and it is nested inside a Grouped AVP, <br> define a sequence of AVP numbers (max 3) in order of parent to child. So say persist AVP number X <br> is nested inside AVP Y which is nested in Z, then define the list as Z Y X.<br>Minimum value = 1.
// Persistencebackup. Backup persistence type for the virtual server. Becomes operational if the primary persistence mechanism fails.<br>Possible values = SOURCEIP, NONE.
// Persistencetype. Type of persistence for the virtual server. Available settings function as follows:<br>* SOURCEIP - Connections from the same client IP address belong to the same persistence session.<br>* COOKIEINSERT - Connections that have the same HTTP Cookie, inserted by a Set-Cookie directive from a server, belong to the same persistence session. <br>* SSLSESSION - Connections that have the same SSL Session ID belong to the same persistence session.<br>* CUSTOMSERVERID - Connections with the same server ID form part of the same session. For this persistence type, set the Server ID (CustomServerID) parameter for each service and configure the Rule parameter to identify the server ID in a request.<br>* RULE - All connections that match a user defined rule belong to the same persistence session. <br>* URLPASSIVE - Requests that have the same server ID in the URL query belong to the same persistence session. The server ID is the hexadecimal representation of the IP address and port of the service to which the request must be forwarded. This persistence type requires a rule to identify the server ID in the request. <br>* DESTIP - Connections to the same destination IP address belong to the same persistence session.<br>* SRCIPDESTIP - Connections that have the same source IP address and destination IP address belong to the same persistence session.<br>* CALLID - Connections that have the same CALL-ID SIP header belong to the same persistence session.<br>* RTSPSID - Connections that have the same RTSP Session ID belong to the same persistence session.<br>* FIXSESSION - Connections that have the same SenderCompID and TargetCompID values belong to the same persistence session.<br>* USERSESSION - Persistence session is created based on the persistence parameter value provided from an extension.<br>Possible values = SOURCEIP, COOKIEINSERT, SSLSESSION, RULE, URLPASSIVE, CUSTOMSERVERID, DESTIP, SRCIPDESTIP, CALLID, RTSPSID, DIAMETER, FIXSESSION, USERSESSION, NONE.
// Persistmask. Persistence mask for IP based persistence types, for IPv4 virtual servers.<br>Minimum length = 1.
// Policyname. Name of the policy bound to the LB vserver.
// Port. Port number for the virtual server.<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Pq. Use priority queuing on the virtual server. based persistence types, for IPv6 virtual servers.<br>Default value: OFF<br>Possible values = ON, OFF.
// Precedence. Precedence.<br>Possible values = RULE, URL.
// Processlocal. By turning on this option packets destined to a vserver in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Push. Process traffic with the push virtual server that is bound to this load balancing virtual server.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Pushlabel. Expression for extracting a label from the servers response. Can be either an expression or the name of a named expression.<br>Default value: "none".
// Pushmulticlients. Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates.<br>Default value: NO<br>Possible values = YES, NO.
// Pushvserver. Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the load balancing virtual server that you are configuring.<br>Minimum length = 1.
// Range. Number of IP addresses that the appliance must generate and assign to the virtual server. The virtual server then functions as a network virtual server, accepting traffic on any of the generated IP addresses. The IP addresses are generated automatically, as follows: <br>* For a range of n, the last octet of the address specified by the IP Address parameter increments n-1 times. <br>* If the last octet exceeds 255, it rolls over to 0 and the third octet increments by 1.<br>Note: The Range parameter assigns multiple IP addresses to one virtual server. To generate an array of virtual servers, each of which owns only one IP address, use brackets in the IP Address and Name parameters to specify the range. For example:<br>add lb vserver my_vserver[1-3] HTTP 192.0.2.[1-3] 80.<br>Default value: 1<br>Minimum value = 1<br>Maximum value = 254.
// Recursionavailable. When set to YES, this option causes the DNS replies from this vserver to have the RA bit turned on. Typically one would set this option to YES, when the vserver is load balancing a set of DNS servers thatsupport recursive queries.<br>Default value: NO<br>Possible values = YES, NO.
// Redirect. Cache redirect type.<br>Possible values = CACHE, POLICY, ORIGIN.
// Redirectfromport. Port number for the virtual server, from which we absorb the traffic for http redirect.<br>Minimum value = 1<br>Range 1 - 65535<br>* in CLI is represented as 65535 in NITRO API.
// Redirectportrewrite. Rewrite the port and change the protocol to ensure successful HTTP redirects from services.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Redirurl. URL to which to redirect traffic if the virtual server becomes unavailable. <br>WARNING! Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.<br>Minimum length = 1.
// Redirurlflags. The redirect URL to be unset.
// Resrule. Default syntax expression specifying which part of a servers response to use for creating rule based persistence sessions (persistence type RULE). Can be either an expression or the name of a named expression.<br>Example:<br>HTTP.RES.HEADER("setcookie").VALUE(0).TYPECAST_NVLIST_T(=,;).VALUE("server1").<br>Default value: "none".
// Retainconnectionsoncluster. This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled.<br>Default value: NO<br>Possible values = YES, NO.
// Rhistate. Route Health Injection (RHI) functionality of the NetSaler appliance for advertising the route of the VIP address associated with the virtual server. When Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:<br>* If you set RHI STATE to PASSIVE on all virtual servers, the NetScaler ADC always advertises the route for the VIP address.<br>* If you set RHI STATE to ACTIVE on all virtual servers, the NetScaler ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.<br>* If you set RHI STATE to ACTIVE on some and PASSIVE on others, the NetScaler ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.<br>Default value: PASSIVE<br>Possible values = PASSIVE, ACTIVE.
// Rtspnat. Use network address translation (NAT) for RTSP data connections.<br>Default value: OFF<br>Possible values = ON, OFF.
// Rule. Expression, or name of a named expression, against which traffic is evaluated. Written in the classic or default syntax.<br>Note:<br>Maximum length of a string literal in the expression is 255 characters. A longer string can be split into smaller strings of up to 255 characters each, and the smaller strings concatenated with the + operator. For example, you can create a 500-character string as follows: ";lt;string of 255 characters;gt;" + ";lt;string of 245 characters;gt;"<br>The following requirements apply only to the NetScaler CLI:<br>* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.<br>* If the expression itself includes double quotation marks, escape the quotations by using the \\ character. <br>* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.<br>Default value: "none".
// Ruletype. Rule type.
// Sc. Use SureConnect on the virtual server.<br>Default value: OFF<br>Possible values = ON, OFF.
// Servicename. Service to bind to the virtual server.<br>Minimum length = 1.
// Servicetype. Protocol used by the service (also called the service type).<br>Possible values = HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, DNS, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, RTSP, PUSH, SSL_PUSH, RADIUS, RDP, MYSQL, MSSQL, DIAMETER, SSL_DIAMETER, TFTP, ORACLE, SMPP, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP.
// Sessionless. Perform load balancing on a per-packet basis, without establishing sessions. Recommended for load balancing of intrusion detection system (IDS) servers and scenarios involving direct server return (DSR), where session information is unnecessary.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Skippersistency. This argument decides the behavior incase the service which is selected from an existing persistence session has reached threshold.<br>Default value: None<br>Possible values = Bypass, ReLb, None.
// Sobackupaction. Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists.<br>Possible values = DROP, ACCEPT, REDIRECT.
// Somethod. Type of threshold that, when exceeded, triggers spillover. Available settings function as follows:<br>* CONNECTION - Spillover occurs when the number of client connections exceeds the threshold.<br>* DYNAMICCONNECTION - Spillover occurs when the number of client connections at the virtual server exceeds the sum of the maximum client (Max Clients) settings for bound services. Do not specify a spillover threshold for this setting, because the threshold is implied by the Max Clients settings of bound services.<br>* BANDWIDTH - Spillover occurs when the bandwidth consumed by the virtual servers incoming and outgoing traffic exceeds the threshold. <br>* HEALTH - Spillover occurs when the percentage of weights of the services that are UP drops below the threshold. For example, if services svc1, svc2, and svc3 are bound to a virtual server, with weights 1, 2, and 3, and the spillover threshold is 50%, spillover occurs if svc1 and svc3 or svc2 and svc3 transition to DOWN. <br>* NONE - Spillover does not occur.<br>Possible values = CONNECTION, DYNAMICCONNECTION, BANDWIDTH, HEALTH, NONE.
// Sopersistence. If spillover occurs, maintain source IP address based persistence for both primary and backup virtual servers.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Sopersistencetimeout. Timeout for spillover persistence, in minutes.<br>Default value: 2<br>Minimum value = 2<br>Maximum value = 1440.
// Sothreshold. Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).<br>Minimum value = 1<br>Maximum value = 4294967287.
// State. State of the load balancing virtual server.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Statechangetimemsec. Time at which last state change happened. Milliseconds part.
// Statechangetimesec. Time when last state change happened. Seconds part.
// Statechangetimeseconds. Time when last state change happened. Seconds part.
// Status. Current status of the lb vserver. During the initial phase if the configured lb method is not round robin , the vserver will adopt round robin to distribute traffic for a predefined number of requests.
// Tcpprofilename. Name of the TCP profile whose settings are to be applied to the virtual server.<br>Minimum length = 1<br>Maximum length = 127.
// Td. Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.<br>Minimum value = 0<br>Maximum value = 4094.
// Thresholdvalue. Tells whether threshold exceeded for this service participating in CUSTOMLB.
// Tickssincelaststatechange. Time in 10 millisecond ticks since the last state change.
// Timeout. Time period for which a persistence session is in effect.<br>Default value: 2<br>Minimum value = 0<br>Maximum value = 1440.
// Tosid. TOS ID of the virtual server. Applicable only when the load balancing redirection mode is set to TOS.<br>Minimum value = 1<br>Maximum value = 63.
// Totalservices. Total number of services bound to the vserver.
// Trofspersistence. When value is ENABLED, Trofs persistence is honored. When value is DISABLED, Trofs persistence is not honored.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Type. Type of LB vserver.<br>Possible values = CONTENT, ADDRESS.
// V6Netmasklen. Number of bits to consider in an IPv6 destination or source IP address, for creating the hash that is required by the DESTINATIONIPHASH and SOURCEIPHASH load balancing methods.<br>Default value: 128<br>Minimum value = 1<br>Maximum value = 128.
// V6Persistmasklen. Persistence mask for IP based persistence types, for IPv6 virtual servers.<br>Default value: 128<br>Minimum value = 1<br>Maximum value = 128.
// Value. SSL status.<br>Possible values = Certkey not bound, SSL feature disabled.
// Version. Cookie version.
// Vipheader. Name for the inserted header. The default name is vip-header.<br>Minimum length = 1.
// Vsvrdynconnsothreshold. Spillover threshold for dynamic connection.
// Weight. Weight to assign to the specified service.<br>Minimum value = 1<br>Maximum value = 100.

// LbvserverAdd defines add request.
type LbvserverAdd struct {
	Lbvserver LbvserverAddBody `json:"lbvserver,omitempty"`
}

// LbvserverAddBody body fro adding object.
type LbvserverAddBody struct {
	Appflowlog                         string `json:"appflowlog,omitempty"`
	Authentication                     string `json:"authentication,omitempty"`
	Authenticationhost                 string `json:"authenticationhost,omitempty"`
	Authn401                           string `json:"authn401,omitempty"`
	Authnprofile                       string `json:"authnprofile,omitempty"`
	Authnvsname                        string `json:"authnvsname,omitempty"`
	Backuplbmethod                     string `json:"backuplbmethod,omitempty"`
	Backuppersistencetimeout           int    `json:"backuppersistencetimeout,omitempty"`
	Backupvserver                      string `json:"backupvserver,omitempty"`
	Bypassaaaa                         string `json:"bypassaaaa,omitempty"`
	Cacheable                          string `json:"cacheable,omitempty"`
	Clttimeout                         string `json:"clttimeout,omitempty"`
	Comment                            string `json:"comment,omitempty"`
	Connfailover                       string `json:"connfailover,omitempty"`
	Cookiename                         string `json:"cookiename,omitempty"`
	Datalength                         string `json:"datalength,omitempty"`
	Dataoffset                         string `json:"dataoffset,omitempty"`
	Dbprofilename                      string `json:"dbprofilename,omitempty"`
	Dbslb                              string `json:"dbslb,omitempty"`
	Disableprimaryondown               string `json:"disableprimaryondown,omitempty"`
	DNS64                              string `json:"dns64,omitempty"`
	Dnsprofilename                     string `json:"dnsprofilename,omitempty"`
	Downstateflush                     string `json:"downstateflush,omitempty"`
	Hashlength                         int    `json:"hashlength,omitempty"`
	Healththreshold                    string `json:"healththreshold,omitempty"`
	Httpprofilename                    string `json:"httpprofilename,omitempty"`
	Httpsredirecturl                   string `json:"httpsredirecturl,omitempty"`
	Icmpvsrresponse                    string `json:"icmpvsrresponse,omitempty"`
	Insertvserveripport                string `json:"insertvserveripport,omitempty"`
	Ipmask                             string `json:"ipmask,omitempty"`
	Ippattern                          string `json:"ippattern,omitempty"`
	Ipv46                              string `json:"ipv46"`
	L2Conn                             string `json:"l2conn,omitempty"`
	Lbmethod                           string `json:"lbmethod,omitempty"`
	Lbprofilename                      string `json:"lbprofilename,omitempty"`
	Listenpolicy                       string `json:"listenpolicy,omitempty"`
	Listenpriority                     int    `json:"listenpriority,omitempty"`
	M                                  string `json:"m,omitempty"`
	Macmoderetainvlan                  string `json:"macmoderetainvlan,omitempty"`
	Maxautoscalemembers                string `json:"maxautoscalemembers,omitempty"`
	Minautoscalemembers                string `json:"minautoscalemembers,omitempty"`
	Mssqlserverversion                 string `json:"mssqlserverversion,omitempty"`
	Mysqlcharacterset                  string `json:"mysqlcharacterset,omitempty"`
	Mysqlprotocolversion               string `json:"mysqlprotocolversion,omitempty"`
	Mysqlservercapabilities            string `json:"mysqlservercapabilities,omitempty"`
	Mysqlserverversion                 string `json:"mysqlserverversion,omitempty"`
	Name                               string `json:"name"`
	Netmask                            string `json:"netmask,omitempty"`
	Netprofile                         string `json:"netprofile,omitempty"`
	Newservicerequest                  string `json:"newservicerequest,omitempty"`
	Newservicerequestincrementinterval string `json:"newservicerequestincrementinterval,omitempty"`
	Newservicerequestunit              string `json:"newservicerequestunit,omitempty"`
	Oracleserverversion                string `json:"oracleserverversion,omitempty"`
	Persistavpno                       []int  `json:"persistavpno,omitempty"`
	Persistencebackup                  string `json:"persistencebackup,omitempty"`
	Persistencetype                    string `json:"persistencetype,omitempty"`
	Persistmask                        string `json:"persistmask,omitempty"`
	Port                               int    `json:"port,omitempty"`
	Pq                                 string `json:"pq,omitempty"`
	Processlocal                       string `json:"processlocal,omitempty"`
	Push                               string `json:"push,omitempty"`
	Pushlabel                          string `json:"pushlabel,omitempty"`
	Pushmulticlients                   string `json:"pushmulticlients,omitempty"`
	Pushvserver                        string `json:"pushvserver,omitempty"`
	Range                              string `json:"range,omitempty"`
	Recursionavailable                 string `json:"recursionavailable,omitempty"`
	Redirectfromport                   int    `json:"redirectfromport,omitempty"`
	Redirectportrewrite                string `json:"redirectportrewrite,omitempty"`
	Redirurl                           string `json:"redirurl,omitempty"`
	Resrule                            string `json:"resrule,omitempty"`
	Retainconnectionsoncluster         string `json:"retainconnectionsoncluster,omitempty"`
	Rhistate                           string `json:"rhistate,omitempty"`
	Rtspnat                            string `json:"rtspnat,omitempty"`
	Rule                               string `json:"rule,omitempty"`
	Sc                                 string `json:"sc,omitempty"`
	Servicetype                        string `json:"servicetype"`
	Sessionless                        string `json:"sessionless,omitempty"`
	Skippersistency                    string `json:"skippersistency,omitempty"`
	Sobackupaction                     string `json:"sobackupaction,omitempty"`
	Somethod                           string `json:"somethod,omitempty"`
	Sopersistence                      string `json:"sopersistence,omitempty"`
	Sopersistencetimeout               string `json:"sopersistencetimeout,omitempty"`
	Sothreshold                        string `json:"sothreshold,omitempty"`
	State                              string `json:"state,omitempty"`
	Tcpprofilename                     string `json:"tcpprofilename,omitempty"`
	Td                                 string `json:"td,omitempty"`
	Timeout                            int    `json:"timeout,omitempty"`
	Tosid                              int    `json:"tosid,omitempty"`
	Trofspersistence                   string `json:"trofspersistence,omitempty"`
	V6Netmasklen                       string `json:"v6netmasklen,omitempty"`
	V6Persistmasklen                   string `json:"v6persistmasklen,omitempty"`
	Vipheader                          string `json:"vipheader,omitempty"`
}

// LbvserverUpdate defines update request.
type LbvserverUpdate struct {
	Lbvserver LbvserverUpdateBody `json:"lbvserver,omitempty"`
}

// LbvserverUpdateBody body for updating object.
type LbvserverUpdateBody struct {
	Appflowlog                         string `json:"appflowlog,omitempty"`
	Authentication                     string `json:"authentication,omitempty"`
	Authenticationhost                 string `json:"authenticationhost,omitempty"`
	Authn401                           string `json:"authn401,omitempty"`
	Authnprofile                       string `json:"authnprofile,omitempty"`
	Authnvsname                        string `json:"authnvsname,omitempty"`
	Backuplbmethod                     string `json:"backuplbmethod,omitempty"`
	Backuppersistencetimeout           int    `json:"backuppersistencetimeout,omitempty"`
	Backupvserver                      string `json:"backupvserver,omitempty"`
	Bypassaaaa                         string `json:"bypassaaaa,omitempty"`
	Cacheable                          string `json:"cacheable,omitempty"`
	Clttimeout                         string `json:"clttimeout,omitempty"`
	Comment                            string `json:"comment,omitempty"`
	Connfailover                       string `json:"connfailover,omitempty"`
	Cookiename                         string `json:"cookiename,omitempty"`
	Datalength                         string `json:"datalength,omitempty"`
	Dataoffset                         string `json:"dataoffset,omitempty"`
	Dbprofilename                      string `json:"dbprofilename,omitempty"`
	Dbslb                              string `json:"dbslb,omitempty"`
	Disableprimaryondown               string `json:"disableprimaryondown,omitempty"`
	DNS64                              string `json:"dns64,omitempty"`
	Dnsprofilename                     string `json:"dnsprofilename,omitempty"`
	Downstateflush                     string `json:"downstateflush,omitempty"`
	Hashlength                         int    `json:"hashlength,omitempty"`
	Healththreshold                    string `json:"healththreshold,omitempty"`
	Httpprofilename                    string `json:"httpprofilename,omitempty"`
	Httpsredirecturl                   string `json:"httpsredirecturl,omitempty"`
	Icmpvsrresponse                    string `json:"icmpvsrresponse,omitempty"`
	Insertvserveripport                string `json:"insertvserveripport,omitempty"`
	Ipmask                             string `json:"ipmask,omitempty"`
	Ippattern                          string `json:"ippattern,omitempty"`
	Ipv46                              string `json:"ipv46"`
	L2Conn                             string `json:"l2conn,omitempty"`
	Lbmethod                           string `json:"lbmethod,omitempty"`
	Lbprofilename                      string `json:"lbprofilename,omitempty"`
	Listenpolicy                       string `json:"listenpolicy,omitempty"`
	Listenpriority                     int    `json:"listenpriority,omitempty"`
	M                                  string `json:"m,omitempty"`
	Macmoderetainvlan                  string `json:"macmoderetainvlan,omitempty"`
	Maxautoscalemembers                string `json:"maxautoscalemembers,omitempty"`
	Minautoscalemembers                string `json:"minautoscalemembers,omitempty"`
	Mssqlserverversion                 string `json:"mssqlserverversion,omitempty"`
	Mysqlcharacterset                  string `json:"mysqlcharacterset,omitempty"`
	Mysqlprotocolversion               string `json:"mysqlprotocolversion,omitempty"`
	Mysqlservercapabilities            string `json:"mysqlservercapabilities,omitempty"`
	Mysqlserverversion                 string `json:"mysqlserverversion,omitempty"`
	Name                               string `json:"name"`
	Netmask                            string `json:"netmask,omitempty"`
	Netprofile                         string `json:"netprofile,omitempty"`
	Newservicerequest                  string `json:"newservicerequest,omitempty"`
	Newservicerequestincrementinterval string `json:"newservicerequestincrementinterval,omitempty"`
	Newservicerequestunit              string `json:"newservicerequestunit,omitempty"`
	Oracleserverversion                string `json:"oracleserverversion,omitempty"`
	Persistavpno                       []int  `json:"persistavpno,omitempty"`
	Persistencebackup                  string `json:"persistencebackup,omitempty"`
	Persistencetype                    string `json:"persistencetype,omitempty"`
	Persistmask                        string `json:"persistmask,omitempty"`
	Pq                                 string `json:"pq,omitempty"`
	Processlocal                       string `json:"processlocal,omitempty"`
	Push                               string `json:"push,omitempty"`
	Pushlabel                          string `json:"pushlabel,omitempty"`
	Pushmulticlients                   string `json:"pushmulticlients,omitempty"`
	Pushvserver                        string `json:"pushvserver,omitempty"`
	Recursionavailable                 string `json:"recursionavailable,omitempty"`
	Redirectfromport                   int    `json:"redirectfromport,omitempty"`
	Redirectportrewrite                string `json:"redirectportrewrite,omitempty"`
	Redirurl                           string `json:"redirurl,omitempty"`
	Resrule                            string `json:"resrule,omitempty"`
	Retainconnectionsoncluster         string `json:"retainconnectionsoncluster,omitempty"`
	Rhistate                           string `json:"rhistate,omitempty"`
	Rtspnat                            string `json:"rtspnat,omitempty"`
	Rule                               string `json:"rule,omitempty"`
	Sc                                 string `json:"sc,omitempty"`
	Servicename                        string `json:"servicename,omitempty"`
	Sessionless                        string `json:"sessionless,omitempty"`
	Skippersistency                    string `json:"skippersistency,omitempty"`
	Sobackupaction                     string `json:"sobackupaction,omitempty"`
	Somethod                           string `json:"somethod,omitempty"`
	Sopersistence                      string `json:"sopersistence,omitempty"`
	Sopersistencetimeout               string `json:"sopersistencetimeout,omitempty"`
	Sothreshold                        string `json:"sothreshold,omitempty"`
	Tcpprofilename                     string `json:"tcpprofilename,omitempty"`
	Timeout                            int    `json:"timeout,omitempty"`
	Tosid                              int    `json:"tosid,omitempty"`
	Trofspersistence                   string `json:"trofspersistence,omitempty"`
	V6Netmasklen                       string `json:"v6netmasklen,omitempty"`
	V6Persistmasklen                   string `json:"v6persistmasklen,omitempty"`
	Vipheader                          string `json:"vipheader,omitempty"`
	Weight                             int    `json:"weight,omitempty"`
}

// LbvserverEnable defines enable request.
type LbvserverEnable struct {
	Lbvserver LbvserverEnableDisableBody `json:"lbvserver,omitempty"`
}

// LbvserverDisable defines disable request.
type LbvserverDisable struct {
	Lbvserver LbvserverEnableDisableBody `json:"lbvserver,omitempty"`
}

// LbvserverEnableDisableBody body for enabling/disable resource.
type LbvserverEnableDisableBody struct {
	Name string `json:"name,omitempty"`
}

// LbvserverRename defines rename request.
type LbvserverRename struct {
	Lbvserver LbvserverRenameBody `json:"lbvserver,omitempty"`
}

// LbvserverRenameBody body for renaming object.
type LbvserverRenameBody struct {
	Name    string `json:"name,omitempty"`
	Newname string `json:"newname,omitempty"`
}

// LbvserverWrapper wraps the object and serves as default response.
type LbvserverWrapper struct {
	Errorcode int         `json:"errorcode,omitempty"`
	Message   string      `json:"message,omitempty"`
	Severity  string      `json:"severity,omitempty"`
	Lbvserver []Lbvserver `json:"lbvserver,omitempty"`
}

// Lbvserver describes the object.
type Lbvserver struct {
	Activeservices                     string `json:"activeservices,omitempty"`
	Appflowlog                         string `json:"appflowlog,omitempty"`
	Authentication                     string `json:"authentication,omitempty"`
	Authenticationhost                 string `json:"authenticationhost,omitempty"`
	Authn401                           string `json:"authn401,omitempty"`
	Authnprofile                       string `json:"authnprofile,omitempty"`
	Authnvsname                        string `json:"authnvsname,omitempty"`
	Backuplbmethod                     string `json:"backuplbmethod,omitempty"`
	Backuppersistencetimeout           int    `json:"backuppersistencetimeout,omitempty"`
	Backupvserver                      string `json:"backupvserver,omitempty"`
	Backupvserverstatus                string `json:"backupvserverstatus,omitempty"`
	Bindpoint                          string `json:"bindpoint,omitempty"`
	Bypassaaaa                         string `json:"bypassaaaa,omitempty"`
	Cacheable                          string `json:"cacheable,omitempty"`
	Cachevserver                       string `json:"cachevserver,omitempty"`
	Clttimeout                         string `json:"clttimeout,omitempty"`
	Comment                            string `json:"comment,omitempty"`
	Connfailover                       string `json:"connfailover,omitempty"`
	Consolidatedlconn                  string `json:"consolidatedlconn,omitempty"`
	Consolidatedlconngbl               string `json:"consolidatedlconngbl,omitempty"`
	Cookiedomain                       string `json:"cookiedomain,omitempty"`
	Cookiename                         string `json:"cookiename,omitempty"`
	Curstate                           string `json:"curstate,omitempty"`
	Datalength                         string `json:"datalength,omitempty"`
	Dataoffset                         string `json:"dataoffset,omitempty"`
	Dbprofilename                      string `json:"dbprofilename,omitempty"`
	Dbslb                              string `json:"dbslb,omitempty"`
	Disableprimaryondown               string `json:"disableprimaryondown,omitempty"`
	DNS64                              string `json:"dns64,omitempty"`
	Dnsprofilename                     string `json:"dnsprofilename,omitempty"`
	Dnsvservername                     string `json:"dnsvservername,omitempty"`
	Domain                             string `json:"domain,omitempty"`
	Downstateflush                     string `json:"downstateflush,omitempty"`
	Effectivestate                     string `json:"effectivestate,omitempty"`
	Gotopriorityexpression             string `json:"gotopriorityexpression,omitempty"`
	Groupname                          string `json:"groupname,omitempty"`
	Gt2Gb                              string `json:"gt2gb,omitempty"`
	Hashlength                         int    `json:"hashlength,omitempty"`
	Health                             string `json:"health,omitempty"`
	Healththreshold                    string `json:"healththreshold,omitempty"`
	Homepage                           string `json:"homepage,omitempty"`
	Httpprofilename                    string `json:"httpprofilename,omitempty"`
	Httpsredirecturl                   string `json:"httpsredirecturl,omitempty"`
	Icmpvsrresponse                    string `json:"icmpvsrresponse,omitempty"`
	Insertvserveripport                string `json:"insertvserveripport,omitempty"`
	Invoke                             bool   `json:"invoke,omitempty"`
	IP                                 string `json:"ip,omitempty"`
	Ipmapping                          string `json:"ipmapping,omitempty"`
	Ipmask                             string `json:"ipmask,omitempty"`
	Ippattern                          string `json:"ippattern,omitempty"`
	Ipv46                              string `json:"ipv46"`
	Isgslb                             bool   `json:"isgslb,omitempty"`
	L2Conn                             string `json:"l2conn,omitempty"`
	Labelname                          string `json:"labelname,omitempty"`
	Labeltype                          string `json:"labeltype,omitempty"`
	Lbmethod                           string `json:"lbmethod,omitempty"`
	Lbprofilename                      string `json:"lbprofilename,omitempty"`
	Lbrrreason                         int    `json:"lbrrreason,omitempty"`
	Listenpolicy                       string `json:"listenpolicy,omitempty"`
	Listenpriority                     int    `json:"listenpriority,omitempty"`
	M                                  string `json:"m,omitempty"`
	Macmoderetainvlan                  string `json:"macmoderetainvlan,omitempty"`
	Map                                string `json:"map,omitempty"`
	Maxautoscalemembers                string `json:"maxautoscalemembers,omitempty"`
	Minautoscalemembers                string `json:"minautoscalemembers,omitempty"`
	Mssqlserverversion                 string `json:"mssqlserverversion,omitempty"`
	Mysqlcharacterset                  string `json:"mysqlcharacterset,omitempty"`
	Mysqlprotocolversion               string `json:"mysqlprotocolversion,omitempty"`
	Mysqlservercapabilities            string `json:"mysqlservercapabilities,omitempty"`
	Mysqlserverversion                 string `json:"mysqlserverversion,omitempty"`
	Name                               string `json:"name"`
	Netmask                            string `json:"netmask,omitempty"`
	Netprofile                         string `json:"netprofile,omitempty"`
	Newservicerequest                  string `json:"newservicerequest,omitempty"`
	Newservicerequestincrementinterval string `json:"newservicerequestincrementinterval,omitempty"`
	Newservicerequestunit              string `json:"newservicerequestunit,omitempty"`
	Ngname                             string `json:"ngname,omitempty"`
	Oracleserverversion                string `json:"oracleserverversion,omitempty"`
	Persistavpno                       []int  `json:"persistavpno,omitempty"`
	Persistencebackup                  string `json:"persistencebackup,omitempty"`
	Persistencetype                    string `json:"persistencetype,omitempty"`
	Persistmask                        string `json:"persistmask,omitempty"`
	Policyname                         string `json:"policyname,omitempty"`
	Port                               int    `json:"port,omitempty"`
	Pq                                 string `json:"pq,omitempty"`
	Precedence                         string `json:"precedence,omitempty"`
	Processlocal                       string `json:"processlocal,omitempty"`
	Push                               string `json:"push,omitempty"`
	Pushlabel                          string `json:"pushlabel,omitempty"`
	Pushmulticlients                   string `json:"pushmulticlients,omitempty"`
	Pushvserver                        string `json:"pushvserver,omitempty"`
	Range                              string `json:"range,omitempty"`
	Recursionavailable                 string `json:"recursionavailable,omitempty"`
	Redirect                           string `json:"redirect,omitempty"`
	Redirectfromport                   int    `json:"redirectfromport,omitempty"`
	Redirectportrewrite                string `json:"redirectportrewrite,omitempty"`
	Redirurl                           string `json:"redirurl,omitempty"`
	Resrule                            string `json:"resrule,omitempty"`
	Retainconnectionsoncluster         string `json:"retainconnectionsoncluster,omitempty"`
	Rhistate                           string `json:"rhistate,omitempty"`
	Rtspnat                            string `json:"rtspnat,omitempty"`
	Rule                               string `json:"rule,omitempty"`
	Ruletype                           string `json:"ruletype,omitempty"`
	Sc                                 string `json:"sc,omitempty"`
	Servicename                        string `json:"servicename,omitempty"`
	Servicetype                        string `json:"servicetype"`
	Sessionless                        string `json:"sessionless,omitempty"`
	Skippersistency                    string `json:"skippersistency,omitempty"`
	Sobackupaction                     string `json:"sobackupaction,omitempty"`
	Somethod                           string `json:"somethod,omitempty"`
	Sopersistence                      string `json:"sopersistence,omitempty"`
	Sopersistencetimeout               string `json:"sopersistencetimeout,omitempty"`
	Sothreshold                        string `json:"sothreshold,omitempty"`
	Statechangetimemsec                string `json:"statechangetimemsec,omitempty"`
	Statechangetimesec                 string `json:"statechangetimesec,omitempty"`
	Statechangetimeseconds             string `json:"statechangetimeseconds,omitempty"`
	Status                             int    `json:"status,omitempty"`
	Tcpprofilename                     string `json:"tcpprofilename,omitempty"`
	Td                                 string `json:"td,omitempty"`
	Thresholdvalue                     int    `json:"thresholdvalue,omitempty"`
	Tickssincelaststatechange          string `json:"tickssincelaststatechange,omitempty"`
	Timeout                            int    `json:"timeout,omitempty"`
	Tosid                              int    `json:"tosid,omitempty"`
	Totalservices                      string `json:"totalservices,omitempty"`
	Trofspersistence                   string `json:"trofspersistence,omitempty"`
	Type                               string `json:"type,omitempty"`
	V6Netmasklen                       string `json:"v6netmasklen,omitempty"`
	V6Persistmasklen                   string `json:"v6persistmasklen,omitempty"`
	Value                              string `json:"value,omitempty"`
	Version                            int    `json:"version,omitempty"`
	Vipheader                          string `json:"vipheader,omitempty"`
	Vsvrdynconnsothreshold             string `json:"vsvrdynconnsothreshold,omitempty"`
	Weight                             int    `json:"weight,omitempty"`
}
