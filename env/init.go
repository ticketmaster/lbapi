package env

import (
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/dao"
)

// Set prepares all environmental objects.
func Set() {
	config.SetGlobal()
	dao.SetGlobal()
}
