package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// PingAccessAgent ping access agent
// swagger:model PingAccessAgent
type PingAccessAgent struct {

	// UNIX time since epoch in microseconds. Units(MICROSECONDS).
	// Read Only: true
	LastModified *string `json:"_last_modified,omitempty"`

	//  Field introduced in 18.2.3.
	Description *string `json:"description,omitempty"`

	// Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.3.
	Labels []*KeyValue `json:"labels,omitempty"`

	// Name of the PingAccess Agent. Field introduced in 18.2.3.
	// Required: true
	Name *string `json:"name"`

	// Pool containing a primary PingAccess Server, as well as any failover servers included in the agent.properties file. It is a reference to an object of type Pool. Field introduced in 18.2.3.
	// Required: true
	PingaccessPoolRef *string `json:"pingaccess_pool_ref"`

	// The ip and port of the primary PingAccess Server. Field introduced in 18.2.3.
	// Required: true
	PrimaryServer *PoolServer `json:"primary_server"`

	// PingAccessAgent's agent.properties file generated by PingAccess server. Field introduced in 18.2.3.
	// Required: true
	PropertiesFileData *string `json:"properties_file_data"`

	//  It is a reference to an object of type Tenant. Field introduced in 18.2.3.
	TenantRef *string `json:"tenant_ref,omitempty"`

	// url
	// Read Only: true
	URL *string `json:"url,omitempty"`

	// UUID of the PingAccess Agent. Field introduced in 18.2.3.
	UUID *string `json:"uuid,omitempty"`
}