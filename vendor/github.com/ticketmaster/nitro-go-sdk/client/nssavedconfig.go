package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetNssavedconfig returns Nssavedconfig.
func (o Netscaler) GetNssavedconfig() (r model.Nssavedconfig, err error) {
	resp := model.NssavedconfigWrapper{}
	err = o.Session.Get(BaseURI+"nssavedconfig", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Nssavedconfig
	return
}
