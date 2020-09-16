package poolgroup

import (
	"github.com/ticketmaster/lbapi/pool"
)

// PoolGroup - implements poolgroup package logic.
type PoolGroup struct {
	// Avi - implements Avi logic.
	Avi *Avi
}

// Data - resource configuration.
type Data struct {
	// SourceUUID [system] - uuid of resource.
	SourceUUID string
	// Members - uuid of member pools.
	Members []pool.Data
	// Name - name of pool group.
	Name string
}

// Collection - collection of all monitors.
type Collection struct {
	// Source - map of all related records indexed by uuid from source.
	Source map[string]Data
	// System - map of all related records indexed by name from source.
	System map[string]Data
	// Members - map of member uuids and their group map.
	Members map[string]string
}

// RemovedArtifacts - objects removed during modify and marked for deletion
type RemovedArtifacts struct {
	// Pools - removed pools.
	Pools []pool.Data
}
