package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetNshardware returns nshardware facts.
func (o Netscaler) GetNshardware() (r model.Nshardware, err error) {
	resp := model.NshardwareWrapper{}
	err = o.Session.Get(BaseURI+"nshardware", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Nshardware
	return
}
