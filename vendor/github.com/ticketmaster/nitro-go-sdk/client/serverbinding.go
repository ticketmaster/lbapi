package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetServerBindings returns all records of the specific resource.
func (o Netscaler) GetServerBindings() (r []model.ServerBinding, err error) {
	resp := model.ServerBindingWrapper{}
	err = o.Session.Get(BaseURI+"server_binding?bulkbindings=yes", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.ServerBinding
	return
}

// GetServerBinding returns a single record matching the resource identfier.
func (o Netscaler) GetServerBinding(name string) (r []model.ServerBinding, err error) {
	resp := model.ServerBindingWrapper{}
	err = o.Session.Get(BaseURI+"server_binding/"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 || len(resp.ServerBinding) == 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.ServerBinding
	return
}
