package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// ClearNsacls clears all rules.
func (o Netscaler) ClearNsacls() (err error) {
	p := model.Nsacls{}
	resp := model.NsaclsWrapper{}
	err = o.Session.Post(BaseURI+"nsacls?action=clear", &p, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return
}

// ApplyNsacls applies all rules.
func (o Netscaler) ApplyNsacls() (err error) {
	p := model.Nsacls{}
	resp := model.NsaclsWrapper{}
	err = o.Session.Post(BaseURI+"nsacls?action=apply", &p, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return
}
