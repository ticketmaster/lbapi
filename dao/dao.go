package dao

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/ticketmaster/lbapi/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	GlobalDAO *DAO
	log       = logrus.New()
)

// New creates a new dao object.
func New() *DAO {
	config := config.GlobalConfig
	client := new(Client)
	client.Database = config.Database.Database
	client.Password = config.Database.Password
	client.UserName = config.Database.User
	client.Database = config.Database.Database
	client.Port = config.Database.Port
	client.Host = config.Database.Host
	client.SSLMode = config.Database.SSLMode
	return client.Connect()
}

func SetGlobal() {
	GlobalDAO = New()
}

// Connect establishes a database connection.
func (o *Client) Connect() *DAO {
	var err error
	var resp DAO
	if o.Host == "" || o.Port == 0 || o.UserName == "" || o.Password == "" || o.Database == "" || o.SSLMode == "" {
		err = errors.Errorf(
			"all fields must be set (%s)",
			spew.Sdump(o))
		log.Panic(err)
	}
	// The first argument corresponds to the driver name that the driver.
	// (in this case, `lib/pq`) used to register itself in `database/sql`
	// The next argument specifies the parameters to be used in the env.DBO.
	// Details about this string can be seen at https://godoc.org/githubcom/lib/pq.
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		o.UserName, o.Password, o.Database, o.Host, strconv.Itoa(o.Port), o.SSLMode))
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open env.DBO to postgre database (%s)",
			spew.Sdump(o))
		log.Panic(err)
	}
	resp.Db = db
	return &resp
}

// Close terminates a database connection.
func (o *DAO) Close() (err error) {
	if o.Db == nil {
		return
	}
	err = o.Db.Close()
	if err != nil {
		log.Panic(err)
	}
	return
}

// Test validates if connection is still open.
func (o *DAO) Test() (err error) {
	qry := "select version()"
	_, err = o.Db.Query(qry)
	if err != nil {
		log.Panic(err)
	}
	return
}

// PurgeData deletes all data contained within a database table.
func (o *DAO) PurgeData(tables []string) {
	t := strings.Join(tables, ",")
	stmt, err := o.Db.Prepare(`TRUNCATE ` + t + ` RESTART IDENTITY CASCADE;`)
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()
}
