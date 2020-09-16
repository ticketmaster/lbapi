package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetNsversion returns Nsversionfacts.
func (o Netscaler) GetNsversion() (r model.Nsversion, err error) {
	resp := model.NsversionWrapper{}
	err = o.Session.Get(BaseURI+"nsversion", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Nsversion
	return
}
