package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetServicegroupLbmonitorBinding ...
func (o Netscaler) GetServicegroupLbmonitorBinding(name string) (r model.ServicegroupLbmonitorBinding, err error) {
	resp := model.ServicegroupLbmonitorBindingWrapper{}
	err = o.Session.Get(BaseURI+"servicegroup_lbmonitor_binding/"+name, &resp)
	if err != nil || len(resp.ServicegroupLbmonitorBinding) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.ServicegroupLbmonitorBinding[0]
	return
}

// GetServicegroupLbmonitorBindings ...
func (o Netscaler) GetServicegroupLbmonitorBindings() (r []model.ServicegroupLbmonitorBinding, err error) {
	resp := model.ServicegroupLbmonitorBindingWrapper{}
	err = o.Session.Get(BaseURI+"servicegroup_lbmonitor_binding?bulkbindings=yes", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.ServicegroupLbmonitorBinding
	return
}

// DeleteServicegroupLbmonitorBinding ...
func (o Netscaler) DeleteServicegroupLbmonitorBinding(name string, servicegroupName string) (err error) {
	resp := model.ServicegroupLbmonitorBindingWrapper{}
	err = o.Session.Delete(BaseURI+"servicegroup_lbmonitor_binding/"+servicegroupName+"?args=monitor_name:"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return
}

// AddServicegroupLbmonitorBinding ...
func (o Netscaler) AddServicegroupLbmonitorBinding(payload model.ServicegroupLbmonitorBindingPayload) (r model.ServicegroupLbmonitorBinding, err error) {
	resp := model.ServicegroupLbmonitorBindingWrapper{}
	req := model.ServicegroupLbmonitorBindingWrapperAdd{ServicegroupLbmonitorBinding: payload}
	err = o.Session.Put(BaseURI+"servicegroup_lbmonitor_binding", &req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetServicegroupLbmonitorBinding(req.ServicegroupLbmonitorBinding.Servicegroupname)
}
