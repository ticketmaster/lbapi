package client

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetLbvserverServiceBindings returns all records of the specific resource.
func (o Netscaler) GetLbvserverServiceBindings() (r []model.LbvserverServiceBinding, err error) {
	resp := model.LbvserverServiceBindingWrapper{}
	err = o.Session.Get(BaseURI+"lbvserver_service_binding?bulkbindings=yes", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.LbvserverServiceBinding
	return
}

// GetLbvserverServiceBinding returns a single record matching the resource identfier.
func (o Netscaler) GetLbvserverServiceBinding(name string, servicename string) (r model.LbvserverServiceBinding, err error) {
	resp := []model.LbvserverServiceBinding{}
	resp, err = o.GetLbvserverServiceBindings()
	if err != nil {
		return
	}
	mResp := make(map[string]model.LbvserverServiceBinding)
	for _, v := range resp {
		mResp[v.Name+"-"+v.Servicename] = v
	}
	if mResp[name+"-"+servicename].Servicename == "" {
		err = fmt.Errorf("resource does not exist %s - %s", name, servicename)
		return
	}
	r = mResp[name+"-"+servicename]
	return
}

// AddLbvserverServiceBinding creates a resource.
func (o Netscaler) AddLbvserverServiceBinding(req model.LbvserverServiceBindingAdd) (r model.LbvserverServiceBinding, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"lbvserver_service_binding", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetLbvserverServiceBinding(req.LbvserverServiceBinding.Name, req.LbvserverServiceBinding.Servicename)
}

// DeleteLbvserverServiceBinding deletes a resource.
func (o Netscaler) DeleteLbvserverServiceBinding(name string, servicename string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"lbvserver_service_binding/"+name+"?args=servicename:"+servicename, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableLbvserverServiceBinding disables a resource. This fun
func (o Netscaler) DisableLbvserverServiceBinding(req model.ServiceDisable) (err error) {
	return o.DisableService(req)
}

// EnableLbvserverServiceBinding enables a resource.
func (o Netscaler) EnableLbvserverServiceBinding(req model.ServiceEnable) (err error) {
	return o.EnableService(req)
}
