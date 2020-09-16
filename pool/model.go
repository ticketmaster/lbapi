package pool

import (
	"github.com/ticketmaster/lbapi/certificate"
	"github.com/ticketmaster/lbapi/monitor"
	"github.com/ticketmaster/lbapi/persistence"
)

// Pool - implements pool package logic.
type Pool struct {
	// Avi - implements Avi logic.
	Avi *Avi
	// Netscaler - implements Netscaler logic.
	Netscaler *Netscaler
}

// RemovedArtifacts - objects removed during modify and marked for deletion
type RemovedArtifacts struct {
	// HealthMonitors - removed health monitors.
	HealthMonitors []monitor.Data
	// PersistenceProfiles - removed persistence profiles.
	PersistenceProfiles []persistence.Data
	// Certificates - removed certificates (Avi).
	Certificates []certificate.Data
	// Servers - removed backend servers (Netscaler).
	Servers []Server
}

// Data - resource configuration.
type Data struct {
	// Name - friendly name of the resource. Typically inherited from parent.
	Name string `json:"name"`
	// DefaultPort [optional] - this port will be mapped to all bindings.
	DefaultPort int `json:"default_port,omitempty"`
	// Enabled - controls the state of the pool.
	Enabled bool `json:"enabled"`
	// Certificate - certificate object. Only supported by Avi. This will allow
	// key authentication between the VIP and backend servers.
	Certificate certificate.Data `json:"certificate,omitempty"`
	// SSLEnabled - enables ssl between pool members and the vip.
	SSLEnabled bool `json:"ssl_enabled"`
	// GracefulDisableTimeout - only supported by Avi.
	GracefulDisableTimeout int `json:"graceful_disable_timeout,omitempty"`
	// Bindings [optional] - list of all backend servers.
	Bindings []MemberBinding `json:"bindings,omitempty"`
	// DisableDelay [optional] - time in seconds before pool is disabled.
	DisableDelay int `json:"disable_delay,omitempty"`
	// GracefulDisable - waits for all sessions to complete before
	// disabling the resource.
	GracefulDisable bool `json:"graceful_disable"`
	// HealthMonitors [optional] - list of all health monitors.
	HealthMonitors []monitor.Data `json:"health_monitors"`
	// IsNsrService - creates a service instead of a servicegroup.
	// Only supported by Netscaler.
	IsNsrService bool `json:"is_nsr_service"`
	// MaxClientConnections [optional] - maximum number of client connections.
	MaxClientConnections int `json:"max_client_connections,omitempty"`
	// Persistence [optional] - persistence object.
	Persistence persistence.Data `json:"persistence,omitempty"`
	// Priority - used only in pool group configurations. Only supported by Avi.
	// Equal priority results in Round Robin. Pools with higher priority take
	// take precedence, resulting in Active/Standby.
	Priority int `json:"priority,omitempty"`
	// Weight - used only in pool group configurations. Only supported by Avi.
	Weight int `json:"weight,omitempty"`
	// SourceStatus [system] - status of resource.
	SourceStatus string `json:"_status,omitempty"`
	// SourceUUID [system] - uuid of resource.
	SourceUUID string `json:"_uuid,omitempty"`
	// SourceLoadBalancingMethod [system] - lb method inherited from virtual service.
	SourceLoadBalancingMethod string `json:"_load_balancing_method,omitempty"`
	// SourceVrfRef [system] - vrf reference from Avi.
	SourceVrfRef string `json:"_vrf_ref,omitempty"`
	// SourceServiceType [system] - service type inherited from virtual service.
	SourceServiceType string `json:"_service_type,omitempty"`
}

// MemberBinding - backend server configuration.
type MemberBinding struct {
	// Port - service port on backend.
	Port int `json:"port"`
	// Server - server object.
	Server Server `json:"server"`
	// Enabled - controls whether or not the backend is enabled.
	Enabled bool `json:"enabled"`
	// DisableDelay [optional] - time in seconds before pool is disabled
	DisableDelay int `json:"disable_delay,omitempty"`
	// GracefulDisable - waits for all sessions to complete before
	// disabling the resource.
	GracefulDisable bool `json:"graceful_disable"`
}

// Collection - map of all pools.
type Collection struct {
	// System - map of all records indexed by name from source.
	System map[string]Data
	// Source - map of all related records indexed by uuid from source.
	Source map[string]Data
}

// MemberBindingsCollection - map of all bindings.
type MemberBindingsCollection struct {
	// Source - map of all related records indexed by uuid from source.
	Source map[string]map[string]MemberBinding
}

// Server - server configuration.
type Server struct {
	// IP - IPV4 of server.
	IP string `json:"ip,omitempty"`
	// SourceDNS [system] - reverse DNS result of IP.
	SourceDNS []string `json:"_dns,omitempty"`
	// SourceUUID [system] - uuid.
	SourceUUID string `json:"_uuid,omitempty"`
}
