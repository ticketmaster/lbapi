package client

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetLbvserverServicegroupBindings returns all records of the specific resource.
func (o Netscaler) GetLbvserverServicegroupBindings() (r []model.LbvserverServicegroupBinding, err error) {
	resp := model.LbvserverServicegroupBindingWrapper{}
	err = o.Session.Get(BaseURI+"lbvserver_servicegroup_binding?bulkbindings=yes", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.LbvserverServicegroupBinding
	return
}

// GetLbvserverServicegroupBinding returns a single record matching the resource identfier.
func (o Netscaler) GetLbvserverServicegroupBinding(name string, servicegroupname string) (r model.LbvserverServicegroupBinding, err error) {
	resp := []model.LbvserverServicegroupBinding{}
	resp, err = o.GetLbvserverServicegroupBindings()
	if err != nil {
		return
	}
	mResp := make(map[string]model.LbvserverServicegroupBinding)
	for _, v := range resp {
		mResp[v.Name+"-"+v.Servicegroupname] = v
	}
	if mResp[name+"-"+servicegroupname].Servicegroupname == "" {
		err = fmt.Errorf("resource does not exist %s - %s", name, servicegroupname)
		return
	}
	r = mResp[name+"-"+servicegroupname]
	return
}

// AddLbvserverServicegroupBinding creates a resource.
func (o Netscaler) AddLbvserverServicegroupBinding(req model.LbvserverServicegroupBindingAdd) (r model.LbvserverServicegroupBinding, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"lbvserver_servicegroup_binding", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetLbvserverServicegroupBinding(req.LbvserverServicegroupBinding.Name, req.LbvserverServicegroupBinding.Servicegroupname)
}

// DeleteLbvserverServicegroupBinding deletes a resource.
func (o Netscaler) DeleteLbvserverServicegroupBinding(name string, servicegroupname string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"lbvserver_servicegroup_binding/"+name+"?args=servicegroupname:"+servicegroupname, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableLbvserverServicegroupBinding disables a resource. This fun
func (o Netscaler) DisableLbvserverServicegroupBinding(req model.ServicegroupDisable) (err error) {
	return o.DisableServicegroup(req)
}

// EnableLbvserverServicegroupBinding enables a resource.
func (o Netscaler) EnableLbvserverServicegroupBinding(req model.ServicegroupEnable) (err error) {
	return o.EnableServicegroup(req)
}
