package model

//	arp	<String>	Read-write	Respond to ARP requests for this IP address.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
//	arpresponse	<String>	Read-write	Respond to ARP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows:<br><br>* NONE - The NetScaler appliance responds to any ARP request for the VIP address, irrespective of the states of the virtual servers associated with the address.<br>* ONE VSERVER - The NetScaler appliance responds to any ARP request for the VIP address if at least one of the associated virtual servers is in UP state.<br>* ALL VSERVER - The NetScaler appliance responds to any ARP request for the VIP address if all of the associated virtual servers are in UP state.<br>Default value: 5<br>Possible values = NONE, ONE_VSERVER, ALL_VSERVERS
//	backplane	<String>	Read-only	The cluster backplane status of the interface. If the status is enabled, the interface is part of the cluster backplane. By default, the backplane status is disabled.<br>Possible values = ENABLED, DISABLED
//	bgp	<String>	Read-write	Use this option to enable or disable BGP on this IP address for the entity.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
//	cleartime	<Double>	Read-only	Time since the interface stats are cleared last time.
//	duplex	<String>	Read-write	The duplex mode for the interface. Notes:* If you set the duplex mode to AUTO, the NetScaler appliance attempts to auto-negotiate the duplex mode of the interface when it is UP. You must enable auto negotiation on the interface. If you set a duplex mode other than AUTO, you must specify the same duplex mode for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors.<br>Default value: AUTO<br>Possible values = AUTO, HALF, FULL
//	dynamicrouting	<String>	Read-write	Allow dynamic routing on this IP address. Specific to Subnet IP (SNIP) address.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
//	flags	<Double>	Read-only	The flags for this entry.
//	freeports	<Double>	Read-only	Number of free Ports available on this IP.
//	ftp	<String>	Read-write	Allow File Transfer Protocol (FTP) access to this IP address.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
//	gui	<String>	Read-write	Allow graphical user interface (GUI) access to this IP address.<br>Default value: ENABLED<br>Possible values = ENABLED, SECUREONLY, DISABLED
//	hostroute	<String>	Read-write	Option to push the VIP to ZebOS routing table for Kernel route redistribution through dynamic routing protocols.<br>Possible values = ENABLED, DISABLED
//	hostrtgw	<String>	Read-write	IP address of the gateway of the route for this VIP address.<br>Default value: -1
//	hostrtgwact	<String>	Read-only	Actual Gateway used for advertising host route.
//	icmp	<String>	Read-write	Respond to ICMP requests for this IP address.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
//	icmpresponse	<String>	Read-write	Respond to ICMP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows:<br>* NONE - The NetScaler appliance responds to any ICMP request for the VIP address, irrespective of the states of the virtual servers associated with the address.<br>* ONE VSERVER - The NetScaler appliance responds to any ICMP request for the VIP address if at least one of the associated virtual servers is in UP state.<br>* ALL VSERVER - The NetScaler appliance responds to any ICMP request for the VIP address if all of the associated virtual servers are in UP state.<br>* VSVR_CNTRLD - The behavior depends on the ICMP VSERVER RESPONSE setting on all the associated virtual servers.<br><br>The following settings can be made for the ICMP VSERVER RESPONSE parameter on a virtual server:<br>* If you set ICMP VSERVER RESPONSE to PASSIVE on all virtual servers, NetScaler always responds.<br>* If you set ICMP VSERVER RESPONSE to ACTIVE on all virtual servers, NetScaler responds if even one virtual server is UP.<br>* When you set ICMP VSERVER RESPONSE to ACTIVE on some and PASSIVE on others, NetScaler responds if even one virtual server set to ACTIVE is UP.<br>Default value: 5<br>Possible values = NONE, ONE_VSERVER, ALL_VSERVERS, VSVR_CNTRLD
//	ipaddress	<String>	Read-write	IPv4 address to create on the NetScaler appliance. Cannot be changed after the IP address is created.<br>Minimum length = 1
//	iptype	<String[]>	Read-only	.<br>Possible values = SNIP, VIP, NSIP, GSLBsiteIP, CLIP, LSN
//	lacpactorcollecting	<String>	Read-only	The Collecting flag indicates that the participant.s collector, i.e. the reception component of the mux, is definitely on. If set the flag communicates collecting.<br>Possible values = NS_EMPTY_STR, COLLECTING
//	lacpactordistributing	<String>	Read-only	The Distributing flag indicates that the participant.s distributor is not definitely off. If reset the flag indicates not distributing.<br>Possible values = NS_EMPTY_STR, DISTRIBUTING
//	lacpactorinsync	<String>	Read-only	The Synchronization flag indicates that the transmitting participant.s mux component is in sync with the system id and key information transmitted.<br>Possible values = NS_EMPTY_STR, SYNC
//	lacppartnercollecting	<String>	Read-only	The Collecting flag indicates that the participant.s collector, i.e. the reception component of the mux, is definitely on. If set the flag communicates collecting.<br>Possible values = NS_EMPTY_STR, COLLECTING
//	lacppartnerdistributing	<String>	Read-only	The Distributing flag indicates that the participant.s distributor is not definitely off. If reset the flag indicates not distributing.<br>Possible values = NS_EMPTY_STR, DISTRIBUTING
//	lacppartnerinsync	<String>	Read-only	The Synchronization flag indicates that the transmitting participant.s mux component is in sync with the system id and key information transmitted.<br>Possible values = NS_EMPTY_STR, SYNC
//	lacppartnertimeout	<String>	Read-only	The timeout value for the information revieved in LACPDUs. It can have values as SHORT or LONG. The SHORT timeout is 3s and the LONG timeout is 90s.<br>Possible values = LONG, SHORT
//	lagtype	<String>	Read-write	Type of entity (NetScaler appliance or cluster configuration) for which to create the channel.<br>Default value: NODE<br>Possible values = NODE, CLUSTER
//	linkstate	<Double>	Read-only	The current state of the link associated with the interface. For logical interfaces (LA), the state of the link is dependent on the state of the slave interfaces. For the link to be UP at least one of the slave interfaces needs to be UP.
//	metric	<Integer>	Read-write	Integer value to add to or subtract from the cost of the route advertised for the VIP address.<br>Minimum value = -16777215
//	mgmtaccess	<String>	Read-write	Allow access to management applications on this IP address.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
//	mtu	<Double>	Read-write	The maximum transmission unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes crc, that can be transmitted and received by this interface. Default value of MTU is 1500 on all the interface of Netscaler appliance any value configured more than 1500 on the interface will make the interface as jumbo enabled. In case of cluster backplane interface MTU value will be changed to 1514 by default, user has to change the backplane interface value to maximum mtu configured on any of the interface in cluster system plus 14 bytes more for backplane interface if Jumbo is enabled on any of the interface in a cluster system. Changing the backplane will bring back the MTU of backplane interface to default value of 1500. If a channel is configured as backplane then the same holds true for channel as well as member interfaces.<br>Default value: 1500<br>Minimum value = 1500<br>Maximum value = 9216
//	netmask	<String>	Read-write	Subnet mask associated with the IP address.
//	networkroute	<String>	Read-write	Option to push the SNIP subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.<br>Possible values = ENABLED, DISABLED
//	ospf	<String>	Read-write	Use this option to enable or disable OSPF on this IP address for the entity.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
//	ospfarea	<Double>	Read-write	ID of the area in which the type1 link-state advertisements (LSAs) are to be advertised for this virtual IP (VIP) address by the OSPF protocol running on the NetScaler appliance. When this parameter is not set, the VIP is advertised on all areas.<br>Default value: -1<br>Minimum value = 0<br>Maximum value = 4294967294LU
//	ospfareaval	<Double>	Read-only	The area ID of the area in which OSPF Type1 LSAs are advertised.
//	ospflsatype	<String>	Read-write	Type of LSAs to be used by the OSPF protocol, running on the NetScaler appliance, for advertising the route for this VIP address.<br>Default value: TYPE5<br>Possible values = TYPE1, TYPE5
//	ownerdownresponse	<String>	Read-write	in cluster system, if the owner node is down, whether should it respond to icmp/arp.<br>Default value: YES<br>Possible values = YES, NO
//	ownernode	<Double>	Read-write	The owner node in a Cluster for this IP address. Owner node can vary from 0 to 31. If ownernode is not specified then the IP is treated as Striped IP.<br>Default value: 255
//	restrictaccess	<String>	Read-write	Block access to nonmanagement applications on this IP. This option is applicable for MIPs, SNIPs, and NSIP, and is disabled by default. Nonmanagement applications can run on the underlying NetScaler Free BSD operating system.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
//	rip	<String>	Read-write	Use this option to enable or disable RIP on this IP address for the entity.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
//	riserhimsgcode	<Integer>	Read-only	The code indicating the rise rhi status.
//	slavetime	<Double>	Read-only	UP time of the member interfaces.
//	snmp	<String>	Read-write	Allow Simple Network Management Protocol (SNMP) access to this IP address.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
//	ssh	<String>	Read-write	Allow secure shell (SSH) access to this IP address.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
//	state	<String>	Read-write	Enable or disable the IP address.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
//	tag	<Double>	Read-write	Tag value for the network/host route associated with this IP.<br>Default value: 0
//	tagged	<Double>	Read-only	VLAN tags setting on this channel.
//	td	<Double>	Read-write	Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.<br>Minimum value = 0<br>Maximum value = 4094
//	telnet	<String>	Read-write	Allow Telnet access to this IP address.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
//	trunk	<String>	Read-write	This argument is deprecated by tagall.<br>Default value: OFF<br>Possible values = ON, OFF
//	trunkallowedvlan	<String[]>	Read-write	VLAN ID or range of VLAN IDs will be allowed on this trunk interface. In the command line interface, separate the range with a hyphen. For example: 40-90.<br>Minimum length = 1<br>Maximum length = 4094
//	type	<String>	Read-write	Type of the IP address to create on the NetScaler appliance. Cannot be changed after the IP address is created. The following are the different types of NetScaler owned IP addresses:<br>* A Subnet IP (SNIP) address is used by the NetScaler ADC to communicate with the servers. The NetScaler also uses the subnet IP address when generating its own packets, such as packets related to dynamic routing protocols, or to send monitor probes to check the health of the servers.<br>* A Virtual IP (VIP) address is the IP address associated with a virtual server. It is the IP address to which clients connect. An appliance managing a wide range of traffic may have many VIPs configured. Some of the attributes of the VIP address are customized to meet the requirements of the virtual server.<br>* A GSLB site IP (GSLBIP) address is associated with a GSLB site. It is not mandatory to specify a GSLBIP address when you initially configure the NetScaler appliance. A GSLBIP address is used only when you create a GSLB site.<br>* A Cluster IP (CLIP) address is the management address of the cluster. All cluster configurations must be performed by accessing the cluster through this IP address.<br>Default value: SNIP<br>Possible values = SNIP, VIP, NSIP, GSLBsiteIP, CLIP, LSN
//	unit	<Double>	Read-only	Unit number for this interface, signifying the sequence number in which this interface is discovered on this Netscaler.
//	viprtadv2bsd	<Boolean>	Read-only	Whether this route is advertised to FreeBSD.
//	vipvsercount	<Double>	Read-only	Number of vservers bound to this VIP.
//	vipvserdowncount	<Double>	Read-only	Number of vservers bound to this VIP, which are down.
//	vipvsrvrrhiactivecount	<Double>	Read-only	Number of vservers that have RHI state ACTIVE.
//	vipvsrvrrhiactiveupcount	<Double>	Read-only	Number of vservers that have RHI state ACTIVE, which are UP.
//	vmac	<String>	Read-only	Virtual MAC of this interface.
//	vmac6	<String>	Read-only	Virtual MAC for IPv6 of this interface.
//	vrid	<Double>	Read-write	A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up NetScaler appliances in an active-active configuration using VRRP.<br>Minimum value = 1<br>Maximum value = 255
//	vserver	<String>	Read-write	Use this option to set (enable or disable) the virtual server attribute for this IP address.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
//	vserverrhilevel	<String>	Read-write	Advertise the route for the Virtual IP (VIP) address on the basis of the state of the virtual servers associated with that VIP.<br>* NONE - Advertise the route for the VIP address, regardless of the state of the virtual servers associated with the address.<br>* ONE VSERVER - Advertise the route for the VIP address if at least one of the associated virtual servers is in UP state.<br>* ALL VSERVER - Advertise the route for the VIP address if all of the associated virtual servers are in UP state.<br>* VSVR_CNTRLD - Advertise the route for the VIP address according to the RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states.<br><br>When Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:<br> * If you set RHI STATE to PASSIVE on all virtual servers, the NetScaler ADC always advertises the route for the VIP address.<br> * If you set RHI STATE to ACTIVE on all virtual servers, the NetScaler ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.<br> *If you set RHI STATE to ACTIVE on some and PASSIVE on others, the NetScaler ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.<br><br>Default value: ONE_VSERVER<br>Possible values = ONE_VSERVER, ALL_VSERVERS, NONE, VSVR_CNTRLD
//	vserverrhimode	<String>	Read-write	Advertise the route for the Virtual IP (VIP) address using dynamic routing protocols or using RISE<br>* DYNMAIC_ROUTING - Advertise the route for the VIP address using dynamic routing protocols (default)<br>* RISE - Advertise the route for the VIP address using RISE.<br>Default value: DYNAMIC_ROUTING<br>Possible values = DYNAMIC_ROUTING, RISE
// InterfaceWrapper wraps the object and serves as default response.

// NsipWrapper wraps the object and serves as default response.
type NsipWrapper struct {
	Errorcode int    `json:"errorcode,omitempty"`
	Message   string `json:"message,omitempty"`
	Severity  string `json:"severity,omitempty"`
	Nsip      []Nsip `json:"nsip,omitempty"`
}

// Nsip describes the object.
type Nsip struct {
	Ipaddress                string   `json:"ipaddress"`
	Td                       string   `json:"td"`
	Type                     string   `json:"type"`
	Netmask                  string   `json:"netmask"`
	Flags                    string   `json:"flags"`
	Arp                      string   `json:"arp"`
	Icmp                     string   `json:"icmp"`
	Vserver                  string   `json:"vserver"`
	Telnet                   string   `json:"telnet"`
	SSH                      string   `json:"ssh"`
	Gui                      string   `json:"gui"`
	Snmp                     string   `json:"snmp"`
	Ftp                      string   `json:"ftp"`
	Mgmtaccess               string   `json:"mgmtaccess"`
	Restrictaccess           string   `json:"restrictaccess"`
	Decrementttl             string   `json:"decrementttl"`
	Dynamicrouting           string   `json:"dynamicrouting"`
	Hostroute                string   `json:"hostroute"`
	Networkroute             string   `json:"networkroute"`
	Tag                      string   `json:"tag"`
	Hostrtgwact              string   `json:"hostrtgwact"`
	Metric                   int      `json:"metric"`
	Ospfareaval              string   `json:"ospfareaval"`
	Vserverrhilevel          string   `json:"vserverrhilevel"`
	Vserverrhimode           string   `json:"vserverrhimode"`
	Viprtadv2Bsd             bool     `json:"viprtadv2bsd"`
	Vipvsercount             string   `json:"vipvsercount"`
	Vipvserdowncount         string   `json:"vipvserdowncount"`
	Vipvsrvrrhiactivecount   string   `json:"vipvsrvrrhiactivecount"`
	Vipvsrvrrhiactiveupcount string   `json:"vipvsrvrrhiactiveupcount"`
	Ospflsatype              string   `json:"ospflsatype"`
	State                    string   `json:"state"`
	Freeports                string   `json:"freeports"`
	Riserhimsgcode           int      `json:"riserhimsgcode"`
	Iptype                   []string `json:"iptype"`
	Icmpresponse             string   `json:"icmpresponse"`
	Ownernode                string   `json:"ownernode"`
	Arpresponse              string   `json:"arpresponse"`
	Ownerdownresponse        string   `json:"ownerdownresponse"`
}

// NsipUpdate defines update request.
type NsipUpdate struct {
	Nsip NsipUpdateBody `json:"nsip,omitempty"`
}

// NsipUpdateBody body for updating object.
type NsipUpdateBody struct {
	Ipaddress      string `json:"ipaddress"`
	Td             string `json:"td"`
	Netmask        string `json:"netmask"`
	Arp            string `json:"arp"`
	Icmp           string `json:"icmp"`
	Vserver        string `json:"vserver"`
	Telnet         string `json:"telnet"`
	Ftp            string `json:"ftp"`
	Gui            string `json:"gui"`
	SSH            string `json:"ssh"`
	Snmp           string `json:"snmp"`
	Mgmtaccess     string `json:"mgmtaccess"`
	Restrictaccess string `json:"restrictaccess"`
	Dynamicrouting string `json:"dynamicrouting"`
	Hostroute      string `json:"hostroute"`
	Networkroute   string `json:"networkroute"`
	Metric         int    `json:"metric"`
	Arpresponse    string `json:"arpresponse"`
	Tag            string `json:"tag"`
	Icmpresponse   string `json:"icmpresponse"`
}

// NsipEnable defines enable request.
type NsipEnable struct {
	Nsip NsipEnableDisableBody `json:"nsip,omitempty"`
}

// NsipDisable defines disable request.
type NsipDisable struct {
	Nsip NsipEnableDisableBody `json:"nsip,omitempty"`
}

// NsipEnableDisableBody body for enabling/disable resource.
type NsipEnableDisableBody struct {
	Ipaddress string `json:"ipaddress,omitempty"`
}
