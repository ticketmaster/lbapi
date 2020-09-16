package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetServicegroupBindings returns a single record matching the resource identfier.
func (o Netscaler) GetServicegroupBindings(name string) (r []model.Servicegroupbinding, err error) {
	resp := model.ServicegroupbindingsWrapper{}
	err = o.Session.Get(BaseURI+"servicegroupbindings/"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 || len(resp.Servicegroupbindings) == 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Servicegroupbindings
	return
}
