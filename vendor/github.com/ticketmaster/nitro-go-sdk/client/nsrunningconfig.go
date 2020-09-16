package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetNsrunningconfig returns Nsrunningconfig.
func (o Netscaler) GetNsrunningconfig() (r model.Nsrunningconfig, err error) {
	resp := model.NsrunningconfigWrapper{}
	err = o.Session.Get(BaseURI+"nsrunningconfig", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Nsrunningconfig
	return
}
