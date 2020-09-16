package virtualserver

import (
	"github.com/ticketmaster/lbapi/certificate"
	"github.com/ticketmaster/lbapi/pool"
	"github.com/ticketmaster/lbapi/poolgroup"
)

// VirtualServer - implements virtualserver package logic.
type VirtualServer struct {
	// Avi - implements Avi logic.
	Avi *Avi
	// Netscaler - implements Netscaler logic.
	Netscaler *Netscaler
}

// Data - resource configuration.
type Data struct {
	// Name - friendly name of the resource. System will preprend prdXXXX to the
	// name if prd code isn't part of the name already.
	Name string `json:"name,omitempty"`
	// ServiceType - http, https, https-no-secure-cookies, http-multiplex-disabled,
	// ssl-l4-app, l4-app, l4-app-udp and ssl-bridge.
	ServiceType string `json:"service_type,omitempty"`
	// IP - ipv4 address.
	IP string `json:"ip,omitempty"`
	// Ports - port configuration.
	Ports []Port `json:"ports,omitempty"`
	// DNS - dns names associated with vip.
	DNS []string `json:"dns,omitempty"`
	// Enabled - enables the vip.
	Enabled bool `json:"enabled"`
	// Certificate - certificate configuration. Only applies to Avi.
	Certificates []certificate.Data `json:"certificates,omitempty"`
	// LoadBalancingMethro - roundrobin or leastconnection.
	LoadBalancingMethod string `json:"load_balancing_method,omitempty"`
	// Pools - pool configurations.
	Pools []pool.Data `json:"pools,omitempty"`
	// ProductCode - product code associated with record.
	ProductCode int `json:"product_code,omitempty"`
	// SourcePoolGroupUUID - uuid of pool group. Only supported by Avi.
	SourcePoolGroupUUID string `json:"_pool_group_uuid,omitempty"`
	// SourceStatus [system] - status of resource. Only applies to Netscaler.
	SourceStatus string `json:"_status,omitempty"`
	// SourceApplicationPolicy [system] - application policy uuid on Avi.
	SourceApplicationPolicy string `json:"_application_policy,omitempty"`
	// SourceNetworkSecurityPolicyRef [system] - network security policy uuid on Avi.
	SourceNetworkSecurityPolicyRef string `json:"_network_security_policy_ref,omitempty"`
	// SourceUUID [system] - uuid of resource on load balancer.
	SourceUUID string `json:"_uuid,omitempty"`
	// SourceVrfRef [system] - virtual routing context uuid on Avi.
	SourceVrfRef string `json:"_vrf_ref,omitempty"`
	// SourceNsrBackupVip [system] - name of backup on Netscaler.
	SourceNsrBackupVip string `json:"_nsr_backup_vip,omitempty"`
	// SourceLast30 [system] - last 30 day uptime for Netscaler.
	SourceLast30 int `json:"_last_30"`
}

// Port ...
type Port struct {
	// Port - service port.
	Port int `json:"port"`
	// L4Profile - tcp, udp, udp-fast-path-vdi, or udp-per-pkt. Note that udp is
	// only supported by ServiceType(s) with udp in their name.
	L4Profile string `json:"l4_profile"`
	// SSLEnabled - enables ssl on the port.
	SSLEnabled bool `json:"ssl_enabled"`
}

// Collection - map of all virtualservices.
type Collection struct {
	// System - map of all records indexed by name from source.
	System map[string]Data
	// Source - map of all related records indexed by uuid from source.
	Source map[string]Data
}

// RemovedArtifacts - objects removed during modify and marked for deletion
type RemovedArtifacts struct {
	// Certificates - removed certificates.
	Certificates []certificate.Data
	// Pools - removed pools.
	Pools []pool.Data
	// PoolGroups - removed pool groups.
	PoolGroups []poolgroup.Data
}

// PoolBindingsCollection - map of all bindings.
type PoolBindingsCollection struct {
	// Source - map of all related records indexed by uuid from source.
	Source map[string]map[string]pool.Data
}

// PromHealthStatus - Returned data from prometheus for health status.
type PromHealthStatus struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Cluster    string `json:"cluster"`
				Instance   string `json:"instance"`
				IP         string `json:"ip"`
				Job        string `json:"job"`
				Name       string `json:"name"`
				Pool       string `json:"pool"`
				TenantUUID string `json:"tenant_uuid"`
				Units      string `json:"units"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}
