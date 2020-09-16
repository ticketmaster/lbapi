package common

import (
	"fmt"

	"github.com/ticketmaster/lbapi/backup"
	"github.com/ticketmaster/lbapi/userenv"
)

// Backup inserts object record.
func (o *Common) Backup(oUser *userenv.User) (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Print(fmt.Sprintf("backing up data: %s", o.Route))
	////////////////////////////////////////////////////////////////////////////
	m := make(map[string][]string)
	r, err := o.Fetch(m, 0, oUser)
	if err != nil {
		return
	}
	b := backup.New()
	err = b.Commit(r)
	if err != nil {
		return
	}
	return
}
