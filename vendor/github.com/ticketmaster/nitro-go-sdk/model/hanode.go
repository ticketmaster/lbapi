package model

// Completedfliptime. To inform user whether flip time is elapsed or not..
// Curflips. Keeps track of number of flips that have happened till now in current interval..
// Deadinterval. Number of seconds after which a peer node is marked DOWN if heartbeat messages are not received from the peer node.<br>Default value: 3<br>Minimum value = 3<br>Maximum value = 60.
// Disifaces. Disabled interfaces..
// Enaifaces. Enabled interfaces..
// Failsafe. Keep one node primary if both nodes fail the health check, so that a partially available node can back up data and handle traffic. This mode is set independently on each node.<br>Default value: OFF<br>Possible values = ON, OFF.
// Flags. The flags for this entry..
// Haheartbeatifaces. HAHEARTBEAT OFF interfaces..
// Hamonifaces. HAMON ON interfaces..
// Haprop. Automatically propagate all commands from the primary to the secondary node, except the following:<br>* All HA configuration related commands. For example, add ha node, set ha node, and bind ha node. <br>* All Interface related commands. For example, set interface and unset interface.<br>* All channels related commands. For example, add channel, set channel, and bind channel.<br>The propagated command is executed on the secondary node before it is executed on the primary. If command propagation fails, or if command execution fails on the secondary, the primary node executes the command and logs an error. Command propagation uses port 3010.<br>Note: After enabling propagation, run force synchronization on either node.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Hastatus. The HA status of the node. The HA status STAYSECONDARY is used to force the secondary device stay as secondary independent of the state of the Primary device. For example, in an existing HA setup, the Primary node has to be upgraded and this process would take few seconds. During the upgradation, it is possible that the Primary node may suffer from a downtime for a few seconds. However, the Secondary should not take over as the Primary node. Thus, the Secondary node should remain as Secondary even if there is a failure in the Primary node.<br> STAYPRIMARY configuration keeps the node in primary state in case if it is healthy, even if the peer node was the primary node initially. If the node with STAYPRIMARY setting (and no peer node) is added to a primary node (which has this node as the peer) then this node takes over as the new primary and the older node becomes secondary. ENABLED state means normal HA operation without any constraints/preferences. DISABLED state disables the normal HA operation of the node.<br>Possible values = ENABLED, STAYSECONDARY, DISABLED, STAYPRIMARY.
// Hasync. Automatically maintain synchronization by duplicating the configuration of the primary node on the secondary node. This setting is not propagated. Automatic synchronization requires that this setting be enabled (the default) on the current secondary node. Synchronization uses TCP port 3010.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED.
// Hellointerval. Interval, in milliseconds, between heartbeat messages sent to the peer node. The heartbeat messages are UDP packets sent to port 3003 of the peer node.<br>Default value: 200<br>Minimum value = 200<br>Maximum value = 1000.
// Id. Number that uniquely identifies the node. For self node, it will always be 0. Peer node values can range from 1-64.<br>Minimum value = 1<br>Maximum value = 64.
// Ifaces. Interfaces on which non-multicast is not seen..
// Inc. This option is required if the HA nodes reside on different networks. When this mode is enabled, the following independent network entities and configurations are neither propagated nor synced to the other node: MIPs, SNIPs, VLANs, routes (except LLB routes), route monitors, RNAT rules (except any RNAT rule with a VIP as the NAT IP), and dynamic routing configurations. They are maintained independently on each node.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED.
// Ipaddress. The NSIP or NSIP6 address of the node to be added for an HA configuration. This setting is neither propagated nor synchronized.<br>Minimum length = 1.
// Masterstatetime. Time elapsed in current master state..
// Maxflips. Max number of flips allowed before becoming sticky primary.<br>Default value: 0.
// Maxfliptime. Interval after which flipping of node states can again start.<br>Default value: 0.
// Description. Description.
// Name. Node Name.
// Netmask. The netmask..
// Pfifaces. Interfaces causing Partial Failure..
// Routemonitor. The IP address (IPv4 or IPv6)..
// Routemonitorstate. State for route monitor.<br>Possible values = UP, DOWN.
// Ssl2. SSL card status.<br>Possible values = DOWN, UP, NOT PRESENT, UNKNOWN.
// State. HA Master State.<br>Possible values = Primary, Secondary, UNKNOWN.
// Syncvlan. Vlan on which HA related communication is sent. This include sync, propagation , connection mirroring , LB persistency config sync, persistent session sync and session state sync. However HA heartbeats can go all interfaces.<br>Minimum value = 1<br>Maximum value = 4094.

// HanodeAdd defines add request.
type HanodeAdd struct {
	Hanode struct {
		ID        int    `json:"string"`
		Ipaddress string `json:"ipaddress"`
		Inc       string `json:"inc"`
	} `json:"hanode"`
}

// HanodeUpdate defines update request.
type HanodeUpdate struct {
	Hanode HanodeUpdateBody `json:"hanode"`
}

// HanodeUpdateBody body for updating object.
type HanodeUpdateBody struct {
	ID            int    `json:"string"`
	Hastatus      string `json:"hastatus"`
	Hasync        string `json:"hasync"`
	Haprop        string `json:"haprop"`
	Hellointerval int    `json:"hellointerval"`
	Deadinterval  int    `json:"deadinterval"`
	Failsafe      string `json:"failsafe"`
	Maxflips      string `json:"maxflips"`
	Maxfliptime   string `json:"maxfliptime"`
	Syncvlan      int    `json:"syncvlan"`
}

// HanodeWrapper wraps the object and serves as default response.
type HanodeWrapper struct {
	Errorcode int      `json:"errorcode"`
	Message   string   `json:"message"`
	Severity  string   `json:"severity"`
	Hanode    []Hanode `json:"hanode"`
}

// Hanode describes the object.
type Hanode struct {
	ID                int         `json:"string"`
	Name              string      `json:"name"`
	Ipaddress         string      `json:"ipaddress"`
	Flags             string      `json:"flags"`
	Hastatus          string      `json:"hastatus"`
	State             string      `json:"state"`
	Hasync            string      `json:"hasync"`
	Haprop            string      `json:"haprop"`
	Enaifaces         string      `json:"enaifaces"`
	Disifaces         string      `json:"disifaces"`
	Hamonifaces       string      `json:"hamonifaces"`
	Haheartbeatifaces string      `json:"haheartbeatifaces"`
	Pfifaces          string      `json:"pfifaces"`
	Ifaces            string      `json:"ifaces"`
	Netmask           string      `json:"netmask"`
	Inc               string      `json:"inc"`
	Ssl2              string      `json:"ssl2"`
	Hellointerval     int         `json:"hellointerval"`
	Deadinterval      int         `json:"deadinterval"`
	Masterstatetime   int         `json:"masterstatetime"`
	Failsafe          string      `json:"failsafe"`
	Routemonitor      string      `json:"routemonitor"`
	Maxflips          string      `json:"maxflips"`
	Maxfliptime       string      `json:"maxfliptime"`
	Curflips          string      `json:"curflips"`
	Completedfliptime string      `json:"completedfliptime"`
	Syncvlan          interface{} `json:"syncvlan"`
	Routemonitorstate string      `json:"routemonitorstate"`
}
