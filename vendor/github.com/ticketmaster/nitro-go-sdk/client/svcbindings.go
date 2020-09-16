package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetSvcBindings returns all the lbvserver bindings associated with the service.
func (o Netscaler) GetSvcBindings(name string) (r []model.Svcbinding, err error) {
	resp := model.SvcbindingsWrapper{}
	err = o.Session.Get(BaseURI+"svcbindings/"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 || len(resp.Svcbindings) == 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Svcbindings
	return
}
