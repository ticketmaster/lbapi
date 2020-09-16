package sdkfork

import (
	"errors"
	"fmt"
	"time"

	"github.com/ticketmaster/lbapi/config"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/session"
)

// Avi stores avi methods.
type Avi struct {
	Setting *config.Setting
	Client  *clients.AviClient
}

// NewAvi ...
func NewAvi() *Avi {
	return &Avi{
		Setting: config.Set(),
	}
}

// Connect Avi Object creates an avi client session.
func (o *Avi) Connect(address string) (r *clients.AviClient, err error) {
	if address == "" {
		err = errors.New("address empty")
		return
	}
	sessionChan := make(chan *clients.AviClient)
	go func(val string) {
		defer close(sessionChan)
		c, err := clients.NewAviClient(address, o.Setting.Avi.User,
			session.SetPassword(o.Setting.Avi.Password),
			session.SetTenant(o.Setting.Avi.Tenant),
			session.SetVersion(o.Setting.Avi.SDKVersion),
			session.SetInsecure)
		if err != nil {
			sessionChan <- nil
			return
		}
		sessionChan <- c
	}(address)
	select {
	case r = <-sessionChan:
	case <-time.After(30 * time.Second):
		err = fmt.Errorf("timout connecting to %s", address)
		return
	}
	if r == nil {
		err = fmt.Errorf("unable to connect to %s", address)
	}
	return
}
