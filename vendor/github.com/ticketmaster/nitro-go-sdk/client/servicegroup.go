package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetServicegroups returns all records of the specific resource.
func (o Netscaler) GetServicegroups() (r []model.Servicegroup, err error) {
	resp := model.ServicegroupWrapper{}
	err = o.Session.Get(BaseURI+"servicegroup", &resp)
	if err != nil {
		return
	}

	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}

	r = resp.Servicegroup

	return
}

// GetServicegroup returns a single record matching the resource identfier.
func (o Netscaler) GetServicegroup(servicegroupname string) (r model.Servicegroup, err error) {
	resp := model.ServicegroupWrapper{}
	err = o.Session.Get(BaseURI+"servicegroup/"+servicegroupname, &resp)
	if err != nil || len(resp.Servicegroup) == 0 {
		return
	}

	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}

	r = resp.Servicegroup[0]

	return
}

// AddServicegroup creates a resource.
func (o Netscaler) AddServicegroup(req model.ServicegroupAdd) (r model.Servicegroup, err error) {
	resp := model.Response{}

	err = o.Session.Post(BaseURI+"servicegroup", req, &resp)

	if err != nil {
		return
	}

	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}

	return o.GetServicegroup(req.Servicegroup.Servicegroupname)
}

// UpdateServicegroup updates a resource.
func (o Netscaler) UpdateServicegroup(req model.ServicegroupUpdate) (r model.Servicegroup, err error) {
	resp := model.Response{}

	err = o.Session.Put(BaseURI+"servicegroup", req, &resp)

	if err != nil {
		return
	}

	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}

	return o.GetServicegroup(req.Servicegroup.Servicegroupname)
}

// DeleteServicegroup deletes a resource.
func (o Netscaler) DeleteServicegroup(servicegroupname string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"servicegroup/"+servicegroupname, &resp)
	if err != nil {
		return
	}

	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableServicegroup disables a resource.
func (o Netscaler) DisableServicegroup(req model.ServicegroupDisable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"servicegroup?action=disable", req, &resp)
	if err != nil {
		return
	}

	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// EnableServicegroup enables a resource.
func (o Netscaler) EnableServicegroup(req model.ServicegroupEnable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"servicegroup?action=enable", req, &resp)
	if err != nil {
		return
	}

	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// RenameServicegroup renames a resource.
func (o Netscaler) RenameServicegroup(req model.ServicegroupRename) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"servicegroup?action=rename", req, &resp)
	if err != nil {
		return
	}

	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}
