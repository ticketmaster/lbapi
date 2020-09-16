package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetServiceLbmonitorBinding ...
func (o Netscaler) GetServiceLbmonitorBinding(name string) (r model.ServiceLbmonitorBinding, err error) {
	resp := model.ServiceLbmonitorBindingWrapper{}
	err = o.Session.Get(BaseURI+"service_lbmonitor_binding/"+name, &resp)
	if err != nil || len(resp.ServiceLbmonitorBinding) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.ServiceLbmonitorBinding[0]
	return
}

// GetServiceLbmonitorBindings ...
func (o Netscaler) GetServiceLbmonitorBindings() (r []model.ServiceLbmonitorBinding, err error) {
	resp := model.ServiceLbmonitorBindingWrapper{}
	err = o.Session.Get(BaseURI+"service_lbmonitor_binding?bulkbindings=yes", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.ServiceLbmonitorBinding
	return
}

// DeleteServiceLbmonitorBinding ...
func (o Netscaler) DeleteServiceLbmonitorBinding(name string, serviceName string) (err error) {
	resp := model.ServiceLbmonitorBindingWrapper{}
	err = o.Session.Delete(BaseURI+"service_lbmonitor_binding/"+serviceName+"?args=monitor_name:"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return
}

// AddServiceLbmonitorBinding ...
func (o Netscaler) AddServiceLbmonitorBinding(payload model.ServiceLbmonitorBindingPayload) (r model.ServiceLbmonitorBinding, err error) {
	resp := model.ServiceLbmonitorBindingWrapper{}
	req := model.ServiceLbmonitorBindingWrapperAdd{ServiceLbmonitorBinding: payload}
	err = o.Session.Put(BaseURI+"service_lbmonitor_binding", &req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetServiceLbmonitorBinding(req.ServiceLbmonitorBinding.Name)
}
