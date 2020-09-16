package client

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetServicegroupServicegroupmemberBindings returns all records of the specific resource.
func (o Netscaler) GetServicegroupServicegroupmemberBindings() (r []model.ServicegroupServicegroupmemberBinding, err error) {
	resp := model.ServicegroupServicegroupmemberBindingWrapper{}
	err = o.Session.Get(BaseURI+"servicegroup_servicegroupmember_binding?bulkbindings=yes", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.ServicegroupServicegroupmemberBinding
	return
}

// GetServicegroupServicegroupmembers all bindings associated with servicegroupname.
func (o Netscaler) GetServicegroupServicegroupmembers(servicegroupname string) (r []model.ServicegroupServicegroupmemberBinding, err error) {
	resp := model.ServicegroupServicegroupmemberBindingWrapper{}
	filter := fmt.Sprintf("servicegroup_servicegroupmember_binding/%s", servicegroupname)
	err = o.Session.Get(BaseURI+filter, &resp)
	if err != nil || len(resp.ServicegroupServicegroupmemberBinding) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.ServicegroupServicegroupmemberBinding
	return
}

// GetServicegroupServicegroupmemberBindingByIP returns a single record matching the resource identfier.
func (o Netscaler) GetServicegroupServicegroupmemberBindingByIP(servicegroupname string, ip string, port int) (r model.ServicegroupServicegroupmemberBinding, err error) {
	resp := model.ServicegroupServicegroupmemberBindingWrapper{}
	filter := fmt.Sprintf("servicegroup_servicegroupmember_binding/%s?filter=ip:%s,port:%v", servicegroupname, ip, port)
	err = o.Session.Get(BaseURI+filter, &resp)
	if err != nil || len(resp.ServicegroupServicegroupmemberBinding) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.ServicegroupServicegroupmemberBinding[0]
	return
}

// AddServicegroupServicegroupmemberBinding creates a resource.
func (o Netscaler) AddServicegroupServicegroupmemberBinding(req model.ServicegroupServicegroupmemberBindingAdd) (r model.ServicegroupServicegroupmemberBinding, err error) {
	resp := model.Response{}

	IPFilter := req.ServicegroupServicegroupmemberBinding.IP

	if req.ServicegroupServicegroupmemberBinding.Servername != "" {
		req.ServicegroupServicegroupmemberBinding.IP = ""
	}

	err = o.Session.Put(BaseURI+"servicegroup_servicegroupmember_binding", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetServicegroupServicegroupmemberBindingByIP(req.ServicegroupServicegroupmemberBinding.Servicegroupname, IPFilter, req.ServicegroupServicegroupmemberBinding.Port)
}

// DeleteServicegroupServicegroupmemberBindingByIP deletes a resource.
func (o Netscaler) DeleteServicegroupServicegroupmemberBindingByIP(servicegroupname string, ip string, port int) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"servicegroup_servicegroupmember_binding/"+servicegroupname+"?args=ip:"+ip+",port:"+strconv.Itoa(port), &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableServicegroupServicegroupmemberBinding disables a resource. This fun
func (o Netscaler) DisableServicegroupServicegroupmemberBinding(req model.ServicegroupDisable) (err error) {
	return o.DisableServicegroup(req)
}

// EnableServicegroupServicegroupmemberBinding enables a resource.
func (o Netscaler) EnableServicegroupServicegroupmemberBinding(req model.ServicegroupEnable) (err error) {
	return o.EnableServicegroup(req)
}
