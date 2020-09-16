package loadbalancer

// LoadBalancer - implements loadbalancer package logic.
type LoadBalancer struct {
	// Avi - implements Avi logic.
	Avi *Avi
	// Netscaler - implements Netscaler logic.
	Netscaler *Netscaler
}

// Data - resource configuration.
type Data struct {
	// ClusterUUID - unique id of cluster. Avi only.
	ClusterUUID string `json:"cluster_uuid,omitempty"`
	// ClusterIP - ip used to manage cluster wide resources. For Netscaler, this
	// would be the first MIP or SNIP configured for management. For Avi, this
	// is the control cluster ip.
	ClusterIP string `json:"cluster_ip,omitempty"`
	// ClusterDNS - reverse DNS lookup of cluster ip.
	ClusterDNS []string `json:"cluster_dns,omitempty"`
	// DNS - reverse DNS lookup of server ip.
	DNS []string `json:"dns,omitempty"`
	// Firmware - firmware running on load balancer.
	// Firmware - firmware running on load balancer.
	Firmware string `json:"firmware,omitempty"`
	// HAMembers - all members of the HA cluster. For Avi, this only relates to
	// the control plane.
	HAMembers []HAMember `json:"ha_partner,omitempty"`
	// Mfr - manufacturer. Either avi or netscaler.
	Mfr string `json:"mfr,omitempty"`
	// Model - model number of appliance. Netscaler only.
	Model string `json:"model,omitempty"`
	// Serial - serial number of appliance. Netscaler only.
	Serial string `json:"serial,omitempty"`
	// DeviceID - unique identifier for licensing. Netscaler only.
	ProductCode int    `json:"product_code,omitempty"`
	DeviceID    string `json:"device_id,omitempty"`
	// Status - system status.
	Status string `json:"status,omitempty"`
	// IPAddresses - list of all the ip addresses mapped to the unit.
	// Netscaler only.
	IPAddresses []IPAddress `json:"ipaddresses,omitempty"`
	// Interfaces - list of all the interfaces configured on the unit.
	// Netscaler only.
	Interfaces []Interface `json:"interfaces,omitempty"`
	// VRFContexts - list of all the vrfcontexts configured on the unit.
	// Avi only.
	VRFContexts []VrfContext `json:"vrf_contexts,omitempty"`
	// Routes - list of all the network routes in CIDR format.
	Routes map[string]string `json:"routes,omitempty"`
}

// HAMember - ha member configuration.
type HAMember struct {
	// IP - IPV4 of ha member.
	IP string `json:"ip,omitempty"`
	// Role - either active or standby.
	Role string `json:"role,omitempty"`
	// Status - status as it pertains to the cluster.
	Status string `json:"status,omitempty"`
}

// Interface - interface configuration.
type Interface struct {
	// ID - id of interface.
	ID string `json:"id,omitempty"`
	// MAC - mac address of interface.
	MAC string `json:"mac,omitempty"`
	// Enabled - state of interface.
	Enabled bool `json:"enabled"`
	// Lacpmode - lacp mode of interface.
	Lacpmode string `json:"lacpmode,omitempty"`
}

// IPAddress - ip address configuration.
type IPAddress struct {
	// IP - IPV4/IPV6 of object.
	IP string `json:"ip,omitempty"`
	// DNS -reverse dns lookup of ip.
	DNS []string `json:"dns,omitempty"`
	// Type ipv4 or ipv6.
	Type string `json:"type,omitempty"`
	// Enabled - state of interface.
	Enabled bool `json:"enabled"`
	// Netmask - netmask for ip.
	Netmask string `json:"netmask,omitempty"`
	// CIDR - network cidr.
	CIDR string `json:"cidr,omitempty"`
}

// Route - route configuration.
type Route struct {
	Network string `json:"network,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Mask    int    `json:"mask,omitempty"`
}

// Runtime - runtime configuration.
type Runtime struct {
	NodeInfo struct {
		UUID        string `json:"uuid"`
		Version     string `json:"version"`
		MgmtIP      string `json:"mgmt_ip"`
		ClusterUUID string `json:"cluster_uuid"`
	} `json:"node_info"`
	NodeStates []struct {
		MgmtIP string `json:"mgmt_ip"`
		Role   string `json:"role"`
	} `json:"node_states"`
}

// SSLProfileCollection - collection of ssl profiles.
type SSLProfileCollection struct {
	// System - map of objects keyed by friendly name.
	System map[string]SSLProfile
	// Source - map of objects keyd by uuid.
	Source map[string]SSLProfile
}

// ServiceTypeCollection - collection of service types.
type ServiceTypeCollection struct {
	// System - map of objects keyed by friendly name.
	System map[string]Service
	// Source - map of objects keyd by uuid.
	Source map[string]Service
}

// RouteCollection - collection of routes.
type RouteCollection struct {
	// System - map of objects keyed by network.
	System map[string]string
}

// NetworkProfileCollection - collection of network profiles.
type NetworkProfileCollection struct {
	// System - map of objects keyed by friendly name.
	System map[string]NetworkProfile
	// Source - map of objects keyd by uuid.
	Source map[string]NetworkProfile
}

// VrfContextCollection - collection of vrfcontexts.
type VrfContextCollection struct {
	// System - map of objects keyed by friendly name.
	System map[string]VrfContext
	// Source - map of objects keyd by uuid.
	Source map[string]VrfContext
}

// VsVipCollection - collection of vsvips.
type VsVipCollection struct {
	// System - map of objects keyed by friendly name.
	System map[string]VsVip
	// Source - map of objects keyd by uuid.
	Source map[string]VsVip
}

// NetworkProfile - part of collection.
type NetworkProfile struct {
	// Name - name of resource on the load balancer.
	Name string `json:"name"`
	// UUID -uuid of resource on load balancer.
	UUID string `json:"uuid,omitempty"`
	// APIName -name of resource in API.
	APIName string `json:"api_name,omitempty"`
}

// VsVip - part of collection.
type VsVip struct {
	// IP - IP address.
	IP string `json:"ip"`
	// UUID -uuid of resource on load balancer.
	UUID string `json:"uuid,omitempty"`
}

// VrfContext  - part of collection.
type VrfContext struct {
	// Name - name of resource on the load balancer.
	Name string `json:"name"`
	// UUID -UUID of object.
	UUID string `json:"uuid,omitempty"`
	// CloudRef -cloud ref of object.
	CloudRef string `json:"cloud_ref,omitempty"`
	// TenantRef - UUID of tenant.
	TenantRef string `json:"tenant_ref,omitempty"`
	// Routes - list of all routes associated to VrfContext.
	Routes []Route `json:"routes,omitempty"`
}

// SSLProfile  - part of collection.
type SSLProfile struct {
	// Name - name of resource on the load balancer.
	Name string `json:"name"`
	// UUID - uuid of resource on the load balancer.
	UUID string `json:"uuid,omitempty"`
}

// Service - part of collection.
type Service struct {
	Name                     string                    `json:"name"`
	UUID                     string                    `json:"uuid,omitempty"`
	APIName                  string                    `json:"api_name,omitempty"`
	SupportedNetworkProfiles map[string]NetworkProfile `json:"supported_network_profiles,omitempty"`
}
