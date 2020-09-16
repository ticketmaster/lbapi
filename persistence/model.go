package persistence

// Persistence - implements persistence package logic.
type Persistence struct {
	// Avi - implements Avi logic.
	Avi *Avi
	// Netscaler - implements Netscaler logic.
	Netscaler *Netscaler
}

// Data - resource configuration.
type Data struct {
	// Name - friendly name of the resource. Typically inherited from parent.
	Name string `json:"name,omitempty"`
	// Description [optional] provides additional details for the resource.
	Description string `json:"description,omitempty"`
	// ObjName [optional] - name of persistence setting resource. This is either the name
	// of a cookie, or header.
	ObjName string `json:"obj_name,omitempty"`
	// Type - types include client-ip, http-cookie, custom-http-header,
	// app-cookie and tls.
	Type string `json:"type,omitempty"`
	// Timeout [optional] - number of seconds before the session times out.
	// 0 means no timeout.
	Timeout int `json:"timeout,omitempty"`
	// SourceUUID [system] - record id of the resource.
	SourceUUID string `json:"_uuid,omitempty"`
}

// Collection - collection of source records with ETL applied.
type Collection struct {
	// System - map of all records indexed by name from source.
	System map[string]Data
	// Source - map of all records indexed by uuid from source.
	Source map[string]Data
}

// Types - collection of all persistence types.
type Types struct {
	// System - map of all related records indexed by uuid from lbapi.
	System map[string]sourceData
	// Source - map of all related records indexed by uuid from source.
	Source map[string]sourceData
}

// sourceData - data derived from source.
type sourceData struct {
	// Ref - object reference.
	Ref string
	// Default - name of default resource.
	Default string
	// UUID - uuid of resource.
	UUID string
}
