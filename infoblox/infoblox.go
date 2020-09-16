package infoblox

import (
	"github.com/ticketmaster/infoblox-go-sdk"
	ib "github.com/ticketmaster/infoblox-go-sdk"
	"github.com/ticketmaster/infoblox-go-sdk/client"
	"github.com/ticketmaster/lbapi/config"
	"github.com/sirupsen/logrus"
)

// Infoblox describes the infoblox object.
type Infoblox struct {
	Client *ib.Infoblox
	Log    *logrus.Entry
}

// NewInfoblox creates a new infoblox object.
func NewInfoblox() *Infoblox {
	ibo := new(Infoblox)
	config := config.Set()
	con := new(client.Host)
	con.UserName = config.Infoblox.User
	con.Password = config.Infoblox.Password
	con.Name = config.Infoblox.Host
	c := infoblox.Set(con)
	ibo.Client = &c
	ibo.Log = logrus.NewEntry(logrus.New())
	return ibo
}
