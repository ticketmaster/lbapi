package dao

import (
	"database/sql"
)

// Client defines all database options.
type Client struct {
	Database string
	Host     string
	Password string
	Port     int
	SSLMode  string
	UserName string
}

// DAO defines a sql connection object.
type DAO struct {
	Db     *sql.DB
	Client Client
}
