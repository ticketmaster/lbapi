package sdkfork

import (
	"errors"
	"fmt"
	"time"

	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/nitro-go-sdk/client"
)

// Netscaler stores netscaler methods.
type Netscaler struct {
	Setting *config.Setting
	Client  *client.Netscaler
}

// NewNetscaler ...
func NewNetscaler() *Netscaler {
	return &Netscaler{
		Setting: config.Set(),
	}
}

// Connect Netscaler object.
func (o *Netscaler) Connect(address string) (r *client.Netscaler, err error) {
	if address == "" {
		err = errors.New("address empty")
		return
	}
	sessionChan := make(chan *client.Netscaler)
	go func(val string) {
		defer close(sessionChan)
		session, err := client.New(address, o.Setting.Nsr.User, o.Setting.Nsr.Password)
		if err != nil {
			sessionChan <- nil
			return
		}
		nsr := client.Netscaler{}
		nsr.Session = session
		sessionChan <- &nsr
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
