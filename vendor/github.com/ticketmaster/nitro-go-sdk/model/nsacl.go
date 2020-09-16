package model

// aclaction	<String>	Read-write	Action to perform on incoming IPv4 packets that match the extended ACL rule. Available settings function as follows: * ALLOW - The NetScaler appliance processes the packet. * BRIDGE - The NetScaler appliance bridges the packet to the destination without processing it. * DENY - The NetScaler appliance drops the packet. Possible values = BRIDGE, DENY, ALLOW
// aclassociate	<String[]>	Read-only	ACL linked. Possible values = NAT, FORWARDINGSESSION, NAT64, LSN
// aclname	<String>	Read-write	Aclname for the extended ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Minimum length = 1
// destip	<Boolean>	Read-write	IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
// destipop	<String>	Read-write	Either the equals (=) or does not equal (!=) logical operator. Possible values = =, !=, EQ, NEQ
// destipval	<String>	Read-write	IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
// destport	<Boolean>	Read-write	Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.  Note: The destination port can be specified only for TCP and UDP protocols.
// destportop	<String>	Read-write	Either the equals (=) or does not equal (!=) logical operator. Possible values = =, !=, EQ, NEQ
// destportval	<String>	Read-write	Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.  Note: The destination port can be specified only for TCP and UDP protocols. Maximum length = 65535
// established	<Boolean>	Read-write	Allow only incoming TCP packets that have the ACK or RST bit set, if the action set for the ACL rule is ALLOW and these packets match the other conditions in the ACL rule.
// hits	<Double>	Read-only	The hits of this ACL.
// icmpcode	<Double>	Read-write	Code of a particular ICMP message type to match against the ICMP code of an incoming ICMP packet. For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code.  If you set this parameter, you must set the ICMP Type parameter. Minimum value = 0 Maximum value = 65536
// icmptype	<Double>	Read-write	ICMP Message type to match against the message type of an incoming ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type.  Note: This parameter can be specified only for the ICMP protocol. Minimum value = 0 Maximum value = 65536
// Interface	<String>	Read-write	ID of an interface. The NetScaler appliance applies the ACL rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL rule to the incoming packets of all interfaces.
// kernelstate	<String>	Read-only	The commit status of the ACL. Default value: NOTAPPLIED Possible values = APPLIED, NOTAPPLIED, RE-APPLY, SFAPPLIED, SFNOTAPPLIED, SFAPPLIED61, SFNOTAPPLIED61
// logstate	<String>	Read-write	Enable or disable logging of events related to the extended ACL rule. The log messages are stored in the configured syslog or auditlog server. Default value: DISABLED Possible values = ENABLED, DISABLED
// newname	<String>	Read-write	New aclname for the extended ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Minimum length = 1
// priority	<Double>	Read-write	Priority for the extended ACL rule that determines the order in which it is evaluated relative to the other extended ACL rules. If you do not specify priorities while creating extended ACL rules, the ACL rules are evaluated in the order in which they are created. Minimum value = 1 Maximum value = 100000
// protocol	<String>	Read-write	Protocol to match against the protocol of an incoming IPv4 packet. Possible values = ICMP, IGMP, TCP, EGP, IGP, ARGUS, UDP, RDP, RSVP, EIGRP, L2TP, ISIS
// protocolnumber	<Double>	Read-write	Protocol to match against the protocol of an incoming IPv4 packet. Minimum value = 1 Maximum value = 255
// ratelimit	<Double>	Read-write	Maximum number of log messages to be generated per second. If you set this parameter, you must enable the Log State parameter. Default value: 100 Minimum value = 1 Maximum value = 10000
// srcip	<Boolean>	Read-write	IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
// srcipop	<String>	Read-write	Either the equals (=) or does not equal (!=) logical operator. Possible values = =, !=, EQ, NEQ
// srcipval	<String>	Read-write	IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example:10.102.29.30-10.102.29.189.
// srcmac	<String>	Read-write	MAC address to match against the source MAC address of an incoming IPv4 packet.
// srcmacmask	<String>	Read-write	Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111". . Default value: "000000000000"
// srcport	<Boolean>	Read-write	Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
// srcportop	<String>	Read-write	Either the equals (=) or does not equal (!=) logical operator. Possible values = =, !=, EQ, NEQ
// srcportval	<String>	Read-write	Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Maximum length = 65535
// state	<String>	Read-write	Enable or disable the extended ACL rule. After you apply the extended ACL rules, the NetScaler appliance compares incoming packets against the enabled extended ACL rules. Default value: ENABLED Possible values = ENABLED, DISABLED
// td	<Double>	Read-write	Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value = 0 Maximum value = 4094
// ttl	<Double>	Read-write	Number of seconds, in multiples of four, after which the extended ACL rule expires. If you do not want the extended ACL rule to expire, do not specify a TTL value. Minimum value = 1 Maximum value = 2147483647
// vlan	<Double>	Read-write	ID of the VLAN. The NetScaler appliance applies the ACL rule only to the incoming packets of the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL rule to the incoming packets on all VLANs. Minimum value = 1 Maximum value = 4094
// vxlan	<Double>	Read-write	ID of the VXLAN. The NetScaler appliance applies the ACL rule only to the incoming packets of the specified VXLAN. If you do not specify a VXLAN ID, the appliance applies the ACL rule to the incoming packets on all VXLANs. Minimum value = 1 Maximum value = 16777215

// NsaclAdd defines add request
type NsaclAdd struct {
	Nsacl NsaclAddBody `json:"nsacl,omitempty"`
}

// NsaclAddBody body for adding object.
type NsaclAddBody struct {
	Aclname        string `json:"aclname,omitempty"`
	Aclaction      string `json:"aclaction,omitempty"`
	Td             string `json:"td,omitempty"`
	Srcip          bool   `json:"srcip,omitempty"`
	Srcipop        string `json:"srcipop,omitempty"`
	Srcipval       string `json:"srcipval,omitempty"`
	Srcport        bool   `json:"srcport,omitempty"`
	Srcportop      string `json:"srcportop,omitempty"`
	Srcportval     string `json:"srcportval,omitempty"`
	Destip         bool   `json:"destip,omitempty"`
	Destipop       string `json:"destipop,omitempty"`
	Destipval      string `json:"destipval,omitempty"`
	Destport       bool   `json:"destport,omitempty"`
	Destportop     string `json:"destportop,omitempty"`
	Destportval    string `json:"destportval,omitempty"`
	TTL            string `json:"ttl,omitempty"`
	Srcmac         string `json:"srcmac,omitempty"`
	Srcmacmask     string `json:"srcmacmask,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	Protocolnumber string `json:"protocolnumber,omitempty"`
	Vlan           string `json:"vlan,omitempty"`
	Vxlan          string `json:"vxlan,omitempty"`
	Interface      string `json:"Interface,omitempty"`
	Established    bool   `json:"established,omitempty"`
	Icmptype       string `json:"icmptype,omitempty"`
	Icmpcode       string `json:"icmpcode,omitempty"`
	Priority       string `json:"priority,omitempty"`
	State          string `json:"state,omitempty"`
	Logstate       string `json:"logstate,omitempty"`
	Ratelimit      string `json:"ratelimit,omitempty"`
}

// NsaclUpdate defines update request.
type NsaclUpdate struct {
	Nsacl NsaclUpdateBody `json:"nsacl,omitempty"`
}

// NsaclUpdateBody body for updating object.
type NsaclUpdateBody struct {
	Aclname        string `json:"aclname,omitempty"`
	Aclaction      string `json:"aclaction,omitempty"`
	Srcip          bool   `json:"srcip,omitempty"`
	Srcipop        string `json:"srcipop,omitempty"`
	Srcipval       string `json:"srcipval,omitempty"`
	Srcport        bool   `json:"srcport,omitempty"`
	Srcportop      string `json:"srcportop,omitempty"`
	Srcportval     string `json:"srcportval,omitempty"`
	Destip         bool   `json:"destip,omitempty"`
	Destipop       string `json:"destipop,omitempty"`
	Destipval      string `json:"destipval,omitempty"`
	Destport       bool   `json:"destport,omitempty"`
	Destportop     string `json:"destportop,omitempty"`
	Destportval    string `json:"destportval,omitempty"`
	Srcmac         string `json:"srcmac,omitempty"`
	Srcmacmask     string `json:"srcmacmask,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	Protocolnumber string `json:"protocolnumber,omitempty"`
	Icmptype       string `json:"icmptype,omitempty"`
	Icmpcode       string `json:"icmpcode,omitempty"`
	Vlan           string `json:"vlan,omitempty"`
	Vxlan          string `json:"vxlan,omitempty"`
	Interface      string `json:"Interface,omitempty"`
	Priority       string `json:"priority,omitempty"`
	Logstate       string `json:"logstate,omitempty"`
	Ratelimit      string `json:"ratelimit,omitempty"`
	Established    bool   `json:"established,omitempty"`
}

// NsaclEnable defines enable request.
type NsaclEnable struct {
	Nsacl NsaclEnableDisable `json:"nsacl,omitempty"`
}

// NsaclDisable defines disable request.
type NsaclDisable struct {
	Nsacl NsaclEnableDisable `json:"nsacl,omitempty"`
}

// NsaclEnableDisable body for enabling/disabling object.
type NsaclEnableDisable struct {
	Aclname string `json:"aclname,omitempty"`
}

// NsaclRename defines rename request.
type NsaclRename struct {
	Nsacl NsaclRenameBody `json:"nsacl,omitempty"`
}

// NsaclRenameBody body for renaming object.
type NsaclRenameBody struct {
	Aclname string `json:"aclname,omitempty"`
	Newname string `json:"newname,omitempty"`
}

// NsaclWrapper wraps the object and serves as default response.
type NsaclWrapper struct {
	Errorcode int     `json:"errorcode,omitempty"`
	Message   string  `json:"message,omitempty"`
	Severity  string  `json:"severity,omitempty"`
	Nsacl     []Nsacl `json:"nsacl,omitempty"`
}

// Nsacl describes the object.
type Nsacl struct {
	Aclname        string   `json:"aclname,omitempty"`
	Td             string   `json:"td,omitempty"`
	Aclaction      string   `json:"aclaction,omitempty"`
	Srcmac         string   `json:"srcmac,omitempty"`
	Srcmacmask     string   `json:"srcmacmask,omitempty"`
	Protocol       string   `json:"protocol,omitempty"`
	Protocolnumber string   `json:"protocolnumber,omitempty"`
	Srcportval     string   `json:"srcportval,omitempty"`
	Srcportop      string   `json:"srcportop,omitempty"`
	Destportval    string   `json:"destportval,omitempty"`
	Destportop     string   `json:"destportop,omitempty"`
	Srcipval       string   `json:"srcipval,omitempty"`
	Srcipop        string   `json:"srcipop,omitempty"`
	Destipval      string   `json:"destipval,omitempty"`
	Destipop       string   `json:"destipop,omitempty"`
	Vlan           string   `json:"vlan,omitempty"`
	Vxlan          string   `json:"vxlan,omitempty"`
	State          string   `json:"state,omitempty"`
	TTL            string   `json:"ttl,omitempty"`
	Icmptype       string   `json:"icmptype,omitempty"`
	Icmpcode       string   `json:"icmpcode,omitempty"`
	Interface      string   `json:"Interface,omitempty"`
	Hits           string   `json:"hits,omitempty"`
	Established    bool     `json:"established,omitempty"`
	Priority       string   `json:"priority,omitempty"`
	Kernelstate    string   `json:"kernelstate,omitempty"`
	Logstate       string   `json:"logstate,omitempty"`
	Ratelimit      string   `json:"ratelimit,omitempty"`
	Aclassociate   []string `json:"aclassociate,omitempty"`
}
