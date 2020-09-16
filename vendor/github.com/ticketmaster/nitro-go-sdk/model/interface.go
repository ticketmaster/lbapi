package model

//	lacpactormode	<String>	Read-only	<br>* Active - The LA channel port of the NetScaler appliance generates LACPDU messages on a regular basis, regardless of any need expressed by its peer device to receive them.<br>* Passive - The LA channel port of the NetScaler appliance does not transmit LACPDU messages unless the peer device port is in the active mode. That is, the port does not speak unless spoken to.<br>* Disabled - Unbinds the interface from the LA channel. If this is the only interface in the LA channel, the LA channel is removed.<br>Possible values = DISABLED, ACTIVE, PASSIVE
//	flowctl	<String>	Read-write	802.3x flow control setting for the interface. The 802.3x specification does not define flow control for 10 Mbps and 100 Mbps speeds, but if a Gigabit Ethernet interface operates at those speeds, the flow control settings can be applied. The flow control setting that is finally applied to an interface depends on auto-negotiation. With the ON option, the peer negotiates the flow control, but the appliance then forces two-way flow control for the interface.<br>Default value: OFF<br>Possible values = OFF, RX, TX, RXTX, ON
//	trunkmode	<String>	Read-write	Accept and send 802.1q VLAN tagged packets, based on Allowed Vlan List of this interface.<br>Default value: OFF<br>Possible values = ON, OFF
//	autonegresult	<Double>	Read-only	Actual auto-negotiation setting for this interface.
//	actduplex	<String>	Read-only	Actual duplex setting for this interface.<br>Possible values = AUTO, HALF, FULL
//	actflowctl	<String>	Read-only	Actual flow control setting for this interface.<br>Possible values = OFF, RX, TX, RXTX, ON
//	actmedia	<String>	Read-only	Actual media setting for this interface.<br>Possible values = AUTO, UTP, FIBER
//	actspeed	<String>	Read-only	Actual speed setting for this interface.<br>Possible values = AUTO, 10, 100, 1000, 10000, 40000
//	actthroughput	<Double>	Read-only	Actual throughput for the interface.
//	tagall	<String>	Read-write	Add a four-byte 802.1q tag to every packet sent on this interface. The ON setting applies the tag for this interfaces native VLAN. OFF applies the tag for all VLANs other than the native VLAN.<br>Default value: OFF<br>Possible values = ON, OFF
//	ifalias	<String>	Read-write	Alias name for the interface. Used only to enhance readability. To perform any operations, you have to specify the interface ID.<br>Default value: " "<br>Maximum length = 31
//	autoneg	<String>	Read-write	Auto-negotiation state of the interface. With the ENABLED setting, the NetScaler appliance auto-negotiates the speed and duplex settings with the peer network device on the link. The NetScaler appliance auto-negotiates the settings of only those parameters (speed or duplex mode) for which the value is set as AUTO.<br>Default value: NSA_DVC_AUTONEG_ON<br>Possible values = DISABLED, ENABLED
//	lacpmode	<String>	Read-write	Bind the interface to a LA channel created by the Link Aggregation control protocol (LACP). <br>Available settings function as follows:<br>* Active - The LA channel port of the NetScaler appliance generates LACPDU messages on a regular basis, regardless of any need expressed by its peer device to receive them.<br>* Passive - The LA channel port of the NetScaler appliance does not transmit LACPDU messages unless the peer device port is in the active mode. That is, the port does not speak unless spoken to.<br>* Disabled - Unbinds the interface from the LA channel. If this is the only interface in the LA channel, the LA channel is removed.<br>Default value: DISABLED<br>Possible values = DISABLED, ACTIVE, PASSIVE
//	ifnum	<String[]>	Read-only	Contains the LA Master, if the interface is part of LA channel.
//	intfstate	<Double>	Read-only	Current state of the specified interface. The interface state set to UP only if the link state is UP and administrative state is ENABLED.
//	description	<String>	Read-only	Display the type of interface, the speeds at which this interface can operate, and, if applicable, the type of SFP,.
//	slaveduplex	<Double>	Read-only	Duplex of the member interfaces.
//	downtime	<Double>	Read-only	Duration for which the interface has been DOWN. This value is reset when the interface state changes to UP.(Example: 3 hours 1 minute 1 second).
//	uptime	<Double>	Read-only	Duration for which the interface has been UP (Example: 3 hours 1 minute 1 second). This value is reset when the interface state changes to DOWN..
//	taggedautolearn	<Double>	Read-only	Dynamic VLAN membership autolearning enabled or disabled on this interface.
//	speed	<String>	Read-write	Ethernet speed of the interface, in Mbps. <br>Notes:<br>* If you set the speed as AUTO, the NetScaler appliance attempts to auto-negotiate or auto-sense the link speed of the interface when it is UP. You must enable auto negotiation on the interface.<br>* If you set a speed other than AUTO, you must specify the same speed for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors. <br>Some interfaces do not support certain speeds. If you specify an unsupported speed, an error message appears.<br>Default value: AUTO<br>Possible values = AUTO, 10, 100, 1000, 10000, 40000
//	flags	<Double>	Read-only	Flags for this interface. Used for communicating the device states.
//	slaveflowctl	<Double>	Read-only	Flowcontrol of the member interfaces.
//	hangdetect	<Double>	Read-only	Hang detection enabled or disabled for this interface.
//	hangreset	<Double>	Read-only	Hang reset enabled or disabled for this interface.
//	bandwidthhigh	<Double>	Read-write	High threshold value for the bandwidth usage of the interface, in Mbps. The NetScaler appliance generates an SNMP trap message when the bandwidth usage of the interface is greater than or equal to the specified high threshold value.<br>Minimum value = 0<br>Maximum value = 160000
//	lacppartnerexpired	<String>	Read-only	If the LACPDUs are received for timeout period, the Receive Machine enters the Expired state and the timer is restarted with the timeout value of SHORT timeout.<br>Possible values = NS_EMPTY_STR, EXPIRED
//	lacppartnerdefaulted	<String>	Read-only	If the timer expires in the Expired state, the Receive Machine enters the Defaulted state.<br>Possible values = NS_EMPTY_STR, DEFAULTED
//	hamonitor	<String>	Read-write	In a High Availability (HA) configuration, monitor the interface for failure events. In an HA configuration, an interface that has HA MON enabled and is not bound to any Failover Interface Set (FIS), is a critical interface. Failure or disabling of any critical interface triggers HA failover.<br>Default value: ON<br>Possible values = ON, OFF
//	haheartbeat	<String>	Read-write	In a High Availability (HA) or Cluster configuration, configure the interface for sending heartbeats. In an HA or Cluster configuration, an interface that has HA Heartbeat disabled should not send the heartbeats.<br>Default value: ON<br>Possible values = OFF, ON
//	lacpkey	<Double>	Read-write	Integer identifying the LACP LA channel to which the interface is to be bound. <br>For an LA channel of the NetScaler appliance, this digit specifies the variable x of an LA channel in LA/x notation, where x can range from 1 to 8. For example, if you specify 3 as the LACP key for an LA channel, the interface is bound to the LA channel LA/3.<br>For an LA channel of a cluster configuration, this digit specifies the variable y of a cluster LA channel in CLA/(y-4) notation, where y can range from 5 to 8. For example, if you specify 6 as the LACP key for a cluster LA channel, the interface is bound to the cluster LA channel CLA/2.<br>Minimum value = 1<br>Maximum value = 8
//	mode	<String>	Read-only	Interface Link Aggregation mode (Auto/Manual) setting.<br>Possible values = MANUAL, AUTO
//	id	<String>	Read-write	Interface number, in C/U format, where C can take one of the following values:<br>* 0 - Indicates a management interface.<br>* 1 - Indicates a 1 Gbps port.<br>* 10 - Indicates a 10 Gbps port.<br>* LA - Indicates a link aggregation port.<br>* LO - Indicates a loop back port.<br>U is a unique integer for representing an interface in a particular port group.
//	taggedany	<Double>	Read-only	Interface setting to accept/drop all tagged packets.
//	intftype	<String>	Read-only	Interface Type, this field will have the interface type either it is virtual, physical or loopback. .<br>Possible values = BROADCOM 5700/5701, TIGON1/TIGON2, INTEL 82546, INTEL 8255X(PRO), Link Aggregate, Loopback, Intel 82541/47, Broadcom 5704, Chelsio 1G, Intel 8247X, Intel 82576 VF, Xen Virtual, Intel 10G, Intel 82599 VF, Hyper v, Cluster LAG, Intel 8247X SFP, XEN Interface, Chelsio 10G, KVM Virtio, VMXNET3
//	lacptimeout	<String>	Read-write	Interval at which the NetScaler appliance sends LACPDU messages to the peer device on the LA channel.<br>Available settings function as follows:<br>LONG - 30 seconds.<br>SHORT - 1 second.<br>Default value: NSA_LACP_TIMEOUT_LONG<br>Possible values = LONG, SHORT
//	lacpactortimeout	<String>	Read-only	Interval at which the NetScaler appliance sends LACPDU messages to the peer device on the LA channel.<br>Available settings function as follows:<br>LONG - 30 seconds.<br>SHORT - 1 second.<br>Possible values = LONG, SHORT
//	lacpactorportno	<Double>	Read-only	LACP Actor port number. LACP uses the port priority with the port number to form the port identifier.
//	lacpactorpriority	<Double>	Read-only	LACP Actor Priority. A LACP port priority is configured on each port using LACP. LACP uses the port priority with the port number to form the port identifier. The port priority determines which ports should be put in standby mode when there is a hardware limitation that prevents all compatible ports from aggregating.
//	lacppartnerkey	<Double>	Read-only	LACP Partner Key. The LACP key used by the partner port.
//	lacppartnerportno	<Double>	Read-only	LACP Partner Port number. LACP uses the port priority with the port number to form the port identifier.
//	lacppartnerpriority	<Double>	Read-only	LACP Partner Priority. A LACP port priority is configured on each port using LACP. LACP uses the port priority with the port number to form the port identifier. <br> The port priority determines which ports should be put in standby mode when there is a hardware limitation that prevents all compatible ports from aggregating.
//	lacppartnerstate	<String>	Read-only	LACP Partner State. Whether the port is in Active or Passive negotiating state.<br>Possible values = MANUAL, AUTO
//	lacppartnersystemmac	<String>	Read-only	LACP Partner System MAC.
//	lacppartnersystempriority	<Double>	Read-only	LACP Partner System Priority. The LACP partners system priority. The values for the priority range from 0 to 65535. The lower the value, the higher the system priority. The switch with the lower system priority value determines which links between LACP partner are active and which are in the standby for each LACP Channel.
//	lacpportmuxstate	<String>	Read-only	LACP Port MUX state. The state of the MUX control machine. The Mux Control Machine attaches the physical port to an aggregate port, using the Selection Logic to choose an appropriate port, and turns the distributor and collector for the physical port on or off as required by protocol information.<br>Possible values = DETACHED, WAITING, ATTACHED, COLLECTING, DISTRIBUTING
//	lacppriority	<Double>	Read-write	LACP port priority, expressed as an integer. The lower the number, the higher the priority. The NetScaler appliance limits the number of interfaces in an LA channel to sixteen.<br>Default value: 32768<br>Minimum value = 1<br>Maximum value = 65535
//	lacpportrxstat	<String>	Read-only	LACP Port RX state. The state of the Receive machine. The Receive Machine maintains partner information, recording protocol information from LACPDUs sent by remote partner(s). Received information is subject to a timeout, and if sufficient time elapses the receive machine will revert to using default partner information.<br>Possible values = INIT, PORT_DISABLED, LACP_DISABLED, EXPIRED, DEFAULTED, CURRENT
//	lacpportselectstate	<String>	Read-only	LACP Port SELECT state. The state of the SELECT state machine, It could be SELECTED or UNSELECTED.<br>Possible values = UNSELECTED, SELECTED, STANDBY
//	lldpmode	<String>	Read-write	Link Layer Discovery Protocol (LLDP) mode for an interface. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.<br>Possible values = NONE, TRANSMITTER, RECEIVER, TRANSCEIVER
//	linkredundancy	<String>	Read-write	Link Redundancy for Cluster LAG.<br>Default value: OFF<br>Possible values = ON, OFF
//	state	<String>	Read-only	Link state of the interface (UP/DOWN).<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
//	throughput	<Double>	Read-write	Low threshold value for the throughput of the interface, in Mbps. In an HA configuration, failover is triggered if the interface has HA MON enabled and the throughput is below the specified the threshold.<br>Minimum value = 0<br>Maximum value = 160000
//	lractiveintf	<Boolean>	Read-only	LR set member interface state(active/inactive).
//	lrsetpriority	<Double>	Read-write	LRSET port priority, expressed as an integer ranging from 1 to 1024. The highest priority is 1. The NetScaler limits the number of interfaces in an LRSET to 8. Within a LRSET the highest LR Priority Interface is considered as the first candidate for the Active interface, if the interface is UP.<br>Default value: 1024<br>Minimum value = 1<br>Maximum value = 1024
//	mac	<String>	Read-only	MAC address for this interface.
//	slavemedia	<Double>	Read-only	Media type of the member interfaces.
//	reqthroughput	<Double>	Read-only	Minimum required throughput for an interface. Failover is triggered if the operating throughput of a Link Aggregation (LA) channel for which HAMON is ON falls below this value. The possible values are:<br>1. 1000Mbps for 1G interfaces.<br>2. 10000Mbps for 10G interfaces.<br>3. 160000Mbps for Link Aggregation channels.<br>Minimum value = 0<br>Maximum value = 160000
//	actualmtu	<Double>	Read-only	MTU for this interface (the largest frame that can transit this interface).
//	devicename	<String>	Read-only	Name of the interface.
//	vlan	<Double>	Read-only	Native VLAN for this interface.
//	bandwidthnormal	<Double>	Read-write	Normal threshold value for the bandwidth usage of the interface, in Mbps. When the bandwidth usage of the interface becomes less than or equal to the specified normal threshold after exceeding the high threshold, the NetScaler appliance generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.<br>Minimum value = 0<br>Maximum value = 160000
//	rxbytes	<Double>	Read-only	Number of bytes received by an interface since the NetScaler appliance was started or the interface statistics were cleared.
//	txbytes	<Double>	Read-only	Number of bytes transmitted by an interface since the NetScaler appliance was started or the interface statistics were cleared.
//	indisc	<Double>	Read-only	Number of error-free inbound packets discarded by the specified interface because of a lack of resources (for example, insufficient receive buffers).
//	outdisc	<Double>	Read-only	Number of error-free outbound packets discarded by the specified interface because of a lack of resources. This statistic is not available on:<br> (1) 10G ports of NetScaler MPX 12500/12500/15500-10G platforms.<br> (2) 10G data ports on NetScaler MPX 17500/19500/21500 platforms.<br> .
//	rxerrors	<Double>	Read-only	Number of inbound packets dropped by the hardware on a specified interface since the NetScaler appliance was started or the interface statistics were cleared. Packets can be dropped because of CRC, length (undersize or oversize), or alignment errors.
//	rxdrops	<Double>	Read-only	Number of inbound packets dropped by the specified interface. Commonly dropped packets are multicast frames, spanning tree BPDUs, packets destined to a MAC not owned by the NetScaler appliance when L2 mode is disabled, or packets tagged for a VLAN that is not bound to the interface. In most healthy networks, this statistic increments at a steady rate regardless of traffic load. A sharp spike in dropped packets generally indicates an issue with connected L2 switches, such as a forwarding database overflow resulting in packets being broadcast on all ports.
//	bdgmacmoved	<Double>	Read-only	Number of MAC moves between ports. A high rate of MAC moves typically indicates a bridge loop between two interfaces.
//	txerrors	<Double>	Read-only	Number of outbound packets dropped by the hardware on a specified interface since the NetScaler appliance was started or the interface statistics were cleared. Packets can be dropped because of length (undersize or oversize) errors or a lack of resources. This statistic is available only for:<br> (1) Loop back interface (LO) of all platforms.<br> (2) All data ports on the NetScaler 12000 platform.<br> (3) Management ports on the Netscaler MPX 15000 and 17000 platforms.<br> .
//	txdrops	<Double>	Read-only	Number of packets dropped in transmission by the specified interface for one of the following reasons.<br> (1) VLAN mismatch.<br> (2) Oversized packets.<br> (3) Interface congestion.<br> (4) Loopback packets sent on non loop back interface.<br> .
//	rxpackets	<Double>	Read-only	Number of packets received by an interface since the NetScaler appliance was started or the interface statistics were cleared.
//	txpackets	<Double>	Read-only	Number of packets transmitted by an interface since the NetScaler appliance was started or the interface statistics were cleared.
//	fctls	<Double>	Read-only	Number of times flow control is performed on the specified interface because of received pause frames.
//	rxstalls	<Double>	Read-only	Number of times the interface stalled, when receiving packets, since the NetScaler appliance was started or the interface statistics were cleared. Receive (Rx) stalls are detected when the following conditions are met:<br> (1)The link is up for more than 10 minutes.<br> (2)Packets are transmitted, but no packets are received for 16 seconds.<br> .
//	txstalls	<Double>	Read-only	Number of times the interface stalled, when transmitting packets, since the NetScaler appliance was started or the interface statistics were cleared. Transmit (Tx) stalls are detected when a packet posted for transmission is not transmitted in 4 seconds.
//	hangs	<Double>	Read-only	Number of times the specified interface detected hangs in the transmit and receive paths since the NetScaler appliance was started or the interface statistics were cleared.
//	bdgmuted	<Double>	Read-only	Number of times the specified interface stopped transmitting and receiving packets because of MAC moves between ports.
//	stsstalls	<Double>	Read-only	Number of times the status updates for a specified interface were stalled since the NetScaler appliance was started or the interface statistics were cleared. A status stall is detected when the status of the interface is not updated by the NIC hardware within 0.8 seconds of the last update.
//	reqduplex	<String>	Read-only	Requested duplex setting for this interface.<br>Possible values = AUTO, HALF, FULL
//	reqflowcontrol	<String>	Read-only	Requested flow control setting for this interface.<br>Possible values = OFF, RX, TX, RXTX, ON
//	reqmedia	<String>	Read-only	Requested media setting for this interface.<br>Possible values = AUTO, UTP, FIBER
//	reqspeed	<String>	Read-only	Requested speed setting for this interface.<br>Possible values = AUTO, 10, 100, 1000, 10000, 40000
//	slavespeed	<Double>	Read-only	Speed of the member interfaces.
//	slavestate	<Double>	Read-only	State of the member interfaces.
//	lacppartneraggregation	<String>	Read-only	The Aggregation flag indicates that the participant will allow the link to be used as part of an aggregate. Otherwise the link is to be used as an individual link, i.e. not aggregated with any other.<br>Possible values = NS_EMPTY_STR, AGG
//	lacpactoraggregation	<String>	Read-only	The Aggregation flag indicates that the participant will allow the link to be used as part of an aggregate. Otherwise the link is to be used as an individual link, i.e. not aggregated with any other.<br>Possible values = NS_EMPTY_STR, AGG
//	backplane	<String>	Read-only	The cluster backplane status of the interface. If the status is enabled, the interface is part of the cluster backplane. By default, the backplane status is disabled.<br>Possible values = ENABLED, DISABLED
//	lacppartnercollecting	<String>	Read-only	The Collecting flag indicates that the participant.s collector, i.e. the reception component of the mux, is definitely on. If set the flag communicates collecting.<br>Possible values = NS_EMPTY_STR, COLLECTING
//	lacpactorcollecting	<String>	Read-only	The Collecting flag indicates that the participant.s collector, i.e. the reception component of the mux, is definitely on. If set the flag communicates collecting.<br>Possible values = NS_EMPTY_STR, COLLECTING
//	linkstate	<Double>	Read-only	The current state of the link associated with the interface. For logical interfaces (LA), the state of the link is dependent on the state of the slave interfaces. For the link to be UP at least one of the slave interfaces needs to be UP.
//	lacppartnerdistributing	<String>	Read-only	The Distributing flag indicates that the participant.s distributor is not definitely off. If reset the flag indicates not distributing.<br>Possible values = NS_EMPTY_STR, DISTRIBUTING
//	lacpactordistributing	<String>	Read-only	The Distributing flag indicates that the participant.s distributor is not definitely off. If reset the flag indicates not distributing.<br>Possible values = NS_EMPTY_STR, DISTRIBUTING
//	duplex	<String>	Read-write	The duplex mode for the interface. Notes:* If you set the duplex mode to AUTO, the NetScaler appliance attempts to auto-negotiate the duplex mode of the interface when it is UP. You must enable auto negotiation on the interface. If you set a duplex mode other than AUTO, you must specify the same duplex mode for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors.<br>Default value: AUTO<br>Possible values = AUTO, HALF, FULL
//	mtu	<Double>	Read-write	The maximum transmission unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes crc, that can be transmitted and received by this interface. Default value of MTU is 1500 on all the interface of Netscaler appliance any value configured more than 1500 on the interface will make the interface as jumbo enabled. In case of cluster backplane interface MTU value will be changed to 1514 by default, user has to change the backplane interface value to maximum mtu configured on any of the interface in cluster system plus 14 bytes more for backplane interface if Jumbo is enabled on any of the interface in a cluster system. Changing the backplane will bring back the MTU of backplane interface to default value of 1500. If a channel is configured as backplane then the same holds true for channel as well as member interfaces.<br>Default value: 1500<br>Minimum value = 1500<br>Maximum value = 9216
//	lacppartnerinsync	<String>	Read-only	The Synchronization flag indicates that the transmitting participant.s mux component is in sync with the system id and key information transmitted.<br>Possible values = NS_EMPTY_STR, SYNC
//	lacpactorinsync	<String>	Read-only	The Synchronization flag indicates that the transmitting participant.s mux component is in sync with the system id and key information transmitted.<br>Possible values = NS_EMPTY_STR, SYNC
//	lacppartnertimeout	<String>	Read-only	The timeout value for the information revieved in LACPDUs. It can have values as SHORT or LONG. The SHORT timeout is 3s and the LONG timeout is 90s.<br>Possible values = LONG, SHORT
//	trunk	<String>	Read-write	This argument is deprecated by tagall.<br>Default value: OFF<br>Possible values = ON, OFF
//	cleartime	<Double>	Read-only	Time since the interface stats are cleared last time.
//	lagtype	<String>	Read-write	Type of entity (NetScaler appliance or cluster configuration) for which to create the channel.<br>Default value: NODE<br>Possible values = NODE, CLUSTER
//	unit	<Double>	Read-only	Unit number for this interface, signifying the sequence number in which this interface is discovered on this Netscaler.
//	slavetime	<Double>	Read-only	UP time of the member interfaces.
//	vmac6	<String>	Read-only	Virtual MAC for IPv6 of this interface.
//	vmac	<String>	Read-only	Virtual MAC of this interface.
//	trunkallowedvlan	<String[]>	Read-write	VLAN ID or range of VLAN IDs will be allowed on this trunk interface. In the command line interface, separate the range with a hyphen. For example: 40-90.<br>Minimum length = 1<br>Maximum length = 4094
//	tagged	<Double>	Read-only	VLAN tags setting on this channel.

// InterfaceWrapper wraps the object and serves as default response.
type InterfaceWrapper struct {
	Errorcode int         `json:"errorcode,omitempty"`
	Message   string      `json:"message,omitempty"`
	Severity  string      `json:"severity,omitempty"`
	Interface []Interface `json:"interface,omitempty"`
}

// Interface describes the object.
type Interface struct {
	ID                        string   `json:"id"`
	Devicename                string   `json:"devicename"`
	Unit                      string   `json:"unit"`
	Description               string   `json:"description"`
	Flags                     string   `json:"flags"`
	Mtu                       string   `json:"mtu"`
	Actualmtu                 string   `json:"actualmtu"`
	Vlan                      string   `json:"vlan"`
	Mac                       string   `json:"mac"`
	Uptime                    string   `json:"uptime"`
	Downtime                  string   `json:"downtime"`
	Reqmedia                  string   `json:"reqmedia,omitempty"`
	Reqspeed                  string   `json:"reqspeed,omitempty"`
	Reqduplex                 string   `json:"reqduplex,omitempty"`
	Reqflowcontrol            string   `json:"reqflowcontrol,omitempty"`
	Hamonitor                 string   `json:"hamonitor"`
	Haheartbeat               string   `json:"haheartbeat"`
	State                     string   `json:"state"`
	Autoneg                   string   `json:"autoneg"`
	Autonegresult             string   `json:"autonegresult"`
	Tagged                    string   `json:"tagged"`
	Tagall                    string   `json:"tagall"`
	Trunkmode                 string   `json:"trunkmode"`
	Taggedany                 string   `json:"taggedany"`
	Taggedautolearn           string   `json:"taggedautolearn"`
	Hangdetect                string   `json:"hangdetect"`
	Hangreset                 string   `json:"hangreset"`
	Linkstate                 string   `json:"linkstate"`
	Intfstate                 string   `json:"intfstate"`
	Rxpackets                 string   `json:"rxpackets"`
	Rxbytes                   string   `json:"rxbytes"`
	Rxerrors                  string   `json:"rxerrors"`
	Rxdrops                   string   `json:"rxdrops"`
	Txpackets                 string   `json:"txpackets"`
	Txbytes                   string   `json:"txbytes"`
	Txerrors                  string   `json:"txerrors"`
	Txdrops                   string   `json:"txdrops"`
	Indisc                    string   `json:"indisc"`
	Outdisc                   string   `json:"outdisc"`
	Fctls                     string   `json:"fctls"`
	Hangs                     string   `json:"hangs"`
	Stsstalls                 string   `json:"stsstalls"`
	Txstalls                  string   `json:"txstalls"`
	Rxstalls                  string   `json:"rxstalls"`
	Bdgmacmoved               string   `json:"bdgmacmoved"`
	Bdgmuted                  string   `json:"bdgmuted"`
	Vmac                      string   `json:"vmac"`
	Vmac6                     string   `json:"vmac6"`
	Lacpmode                  string   `json:"lacpmode,omitempty"`
	Lacppriority              string   `json:"lacppriority,omitempty"`
	Lacptimeout               string   `json:"lacptimeout"`
	Lagtype                   string   `json:"lagtype"`
	Reqthroughput             string   `json:"reqthroughput"`
	Linkredundancy            string   `json:"linkredundancy"`
	Actthroughput             string   `json:"actthroughput"`
	Bandwidthhigh             string   `json:"bandwidthhigh"`
	Bandwidthnormal           string   `json:"bandwidthnormal"`
	Backplane                 string   `json:"backplane"`
	Ifnum                     []string `json:"ifnum"`
	Cleartime                 string   `json:"cleartime"`
	Slavestate                string   `json:"slavestate"`
	Slavemedia                string   `json:"slavemedia"`
	Slavespeed                string   `json:"slavespeed"`
	Slaveduplex               string   `json:"slaveduplex"`
	Slaveflowctl              string   `json:"slaveflowctl"`
	Slavetime                 string   `json:"slavetime"`
	Intftype                  string   `json:"intftype"`
	Lacpactormode             string   `json:"lacpactormode,omitempty"`
	Lacpactorpriority         string   `json:"lacpactorpriority"`
	Lacpactorportno           string   `json:"lacpactorportno"`
	Lacppartneraggregation    string   `json:"lacppartneraggregation"`
	Lacppartnerinsync         string   `json:"lacppartnerinsync"`
	Lacppartnercollecting     string   `json:"lacppartnercollecting"`
	Lacppartnerdistributing   string   `json:"lacppartnerdistributing"`
	Lacppartnerdefaulted      string   `json:"lacppartnerdefaulted"`
	Lacppartnerexpired        string   `json:"lacppartnerexpired"`
	Lacppartnerpriority       string   `json:"lacppartnerpriority"`
	Lacppartnersystemmac      string   `json:"lacppartnersystemmac"`
	Lacppartnersystempriority string   `json:"lacppartnersystempriority"`
	Lacppartnerportno         string   `json:"lacppartnerportno"`
	Lacppartnerkey            string   `json:"lacppartnerkey"`
	Lacpactoraggregation      string   `json:"lacpactoraggregation"`
	Lacpactorinsync           string   `json:"lacpactorinsync"`
	Lacpactorcollecting       string   `json:"lacpactorcollecting"`
	Lacpactordistributing     string   `json:"lacpactordistributing"`
	Lacpportmuxstate          string   `json:"lacpportmuxstate"`
	Lacpportrxstat            string   `json:"lacpportrxstat"`
	Lacpportselectstate       string   `json:"lacpportselectstate"`
	Lldpmode                  string   `json:"lldpmode"`
	Lrsetpriority             string   `json:"lrsetpriority"`
	Lractiveintf              bool     `json:"lractiveintf"`
	Actmedia                  string   `json:"actmedia,omitempty"`
	Actspeed                  string   `json:"actspeed,omitempty"`
	Actduplex                 string   `json:"actduplex,omitempty"`
	Actflowctl                string   `json:"actflowctl,omitempty"`
	Lacpkey                   int      `json:"lacpkey,omitempty"`
	Lacpactortimeout          string   `json:"lacpactortimeout,omitempty"`
	Lacppartnerstate          string   `json:"lacppartnerstate,omitempty"`
	Lacppartnertimeout        string   `json:"lacppartnertimeout,omitempty"`
	Mode                      string   `json:"mode,omitempty"`
}
