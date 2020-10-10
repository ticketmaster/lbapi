package handler

import (
	"github.com/ticketmaster/lbapi/common"
	"github.com/ticketmaster/lbapi/userenv"
)

// Create ...
type Create interface {
	Create([]byte, *userenv.User) (common.DbRecord, error)
}

// CreateBulk ...
type CreateBulk interface {
	CreateBulk([]byte, *userenv.User) (common.DbRecordCollection, error)
}

// ImportAll ...
type ImportAll interface {
	ImportAll(*userenv.User) (common.DbRecordCollection, error)
}

// Definiton ...
type Definiton interface {
}

// Delete ...
type Delete interface {
	Delete(string, *userenv.User) (common.DbRecord, error)
}

// Fetch ...
type Fetch interface {
	Fetch(map[string][]string, int, *userenv.User) (common.DbRecordCollection, error)
}

// FetchVirtualServices ...
type FetchVirtualServices interface {
	FetchVirtualServices(map[string][]string, int, *userenv.User) (common.VsDbRecordCollection, error)
}

// FetchByID ...
type FetchByID interface {
	FetchByID(string, *userenv.User) (common.DbRecordCollection, error)
}

// StageMigration ...
type StageMigration interface {
	StageMigration([]byte, string, *userenv.User) (common.DbRecord, error)
}

// FetchStaged ...
type FetchStaged interface {
	FetchStaged(string, *userenv.User) (common.DbRecord, error)
}

// Init ...
type Init interface {
	Init() (r interface{}, err error)
}

// routeHandler ...
type routeHandler interface {
	GetRoute() string
}

// Handler ...
type Handler struct {
	Definition interface{}
}

// Modify ...
type Modify interface {
	Modify([]byte, *userenv.User) (r common.DbRecord, err error)
}

// Backup ...
type Backup interface {
	Backup(*userenv.User) (err error)
}

// Migrate ...
type Migrate interface {
	Migrate(string, *userenv.User) (common.DbRecord, error)
}
