package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetLbmonbindingsBindings returns all records of the specific resource.
// This function is broken on the netscalers...so you cannot use ?bulkbindings=yes
func (o Netscaler) GetLbmonbindingsBindings() (r []model.LbmonbindingsBinding, err error) {
	// Get all monitors
	resp := model.LbmonitorWrapper{}
	err = o.Session.Get(BaseURI+"lbmonitor", &resp)
	if err != nil {
		return
	}
	// Get all bindings and add them to map
	for _, val := range resp.Lbmonitor {
		respBinding := model.LbmonbindingsBindingWrapper{}
		err = o.Session.Get(BaseURI+"lbmonbindings_binding/"+val.Monitorname, &respBinding)
		if err != nil {
			return
		}
		if len(respBinding.LbmonbindingsBinding) > 0 {
			r = append(r, respBinding.LbmonbindingsBinding[0])
		}
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return
}

// GetLbmonbindingsBinding returns a single record matching the resource identfier.
func (o Netscaler) GetLbmonbindingsBinding(name string) (r model.LbmonbindingsBinding, err error) {
	resp := model.LbmonbindingsBindingWrapper{}
	err = o.Session.Get(BaseURI+"lbmonbindings_binding/"+name, &resp)
	if err != nil || len(resp.LbmonbindingsBinding) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.LbmonbindingsBinding[0]
	return
}
