package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetLbvserverBindings returns all records of the specific resource.
func (o Netscaler) GetLbvserverBindings() (r []model.LbvserverBinding, err error) {
	resp := model.LbvserverBindingWrapper{}
	err = o.Session.Get(BaseURI+"lbvserver_binding?bulkbindings=yes", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.LbvserverBinding
	return
}

// GetLbvserverBinding returns a single record matching the resource identfier.
func (o Netscaler) GetLbvserverBinding(name string) (r model.LbvserverBinding, err error) {
	resp := model.LbvserverBindingWrapper{}
	err = o.Session.Get(BaseURI+"lbvserver_binding/"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.LbvserverBinding[0]
	return
}
