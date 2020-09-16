package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetServices returns all records of the specific resource.
func (o Netscaler) GetServices() (r []model.Service, err error) {
	resp := model.ServiceWrapper{}
	err = o.Session.Get(BaseURI+"service", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Service
	return
}

// GetService returns a single record matching the resource identfier.
func (o Netscaler) GetService(name string) (r model.Service, err error) {
	resp := model.ServiceWrapper{}
	err = o.Session.Get(BaseURI+"service/"+name, &resp)
	if err != nil || len(resp.Service) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Service[0]
	return
}

// AddService creates a resource.
func (o Netscaler) AddService(req model.ServiceAdd) (r model.Service, err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"service", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetService(req.Service.Name)
}

// UpdateService updates a resource.
func (o Netscaler) UpdateService(req model.ServiceUpdate) (r model.Service, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"service", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetService(req.Service.Name)
}

// DeleteService deletes a resource.
func (o Netscaler) DeleteService(name string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"service/"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableService disables a resource.
func (o Netscaler) DisableService(req model.ServiceDisable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"service?action=disable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// EnableService enables a resource.
func (o Netscaler) EnableService(req model.ServiceEnable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"service?action=enable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// RenameService renames a resource.
func (o Netscaler) RenameService(req model.ServiceRename) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"service?action=rename", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}
