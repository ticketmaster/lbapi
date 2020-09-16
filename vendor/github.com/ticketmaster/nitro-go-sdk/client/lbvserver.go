package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetLbvservers returns all records of the specific resource.
func (o Netscaler) GetLbvservers() (r []model.Lbvserver, err error) {
	resp := model.LbvserverWrapper{}
	err = o.Session.Get(BaseURI+"lbvserver", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Lbvserver
	return
}

// GetLbvserver returns a single record matching the resource identfier.
func (o Netscaler) GetLbvserver(name string) (r model.Lbvserver, err error) {
	resp := model.LbvserverWrapper{}
	err = o.Session.Get(BaseURI+"lbvserver/"+name, &resp)
	if err != nil || len(resp.Lbvserver) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Lbvserver[0]
	return
}

// GetLbvserverByIP returns a single record matching the resource identfier.
func (o Netscaler) GetLbvserverByIP(ip string) (r []model.Lbvserver, err error) {
	resp := model.LbvserverWrapper{}
	err = o.Session.Get(BaseURI+"lbvserver?filter=ipv46:"+ip, &resp)
	if err != nil || len(resp.Lbvserver) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Lbvserver
	return
}

// AddLbvserver creates a resource.
func (o Netscaler) AddLbvserver(req model.LbvserverAdd) (r model.Lbvserver, err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"lbvserver", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetLbvserver(req.Lbvserver.Name)
}

// UpdateLbvserver updates a resource.
func (o Netscaler) UpdateLbvserver(req model.LbvserverUpdate) (r model.Lbvserver, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"lbvserver", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetLbvserver(req.Lbvserver.Name)
}

// DeleteLbvserver deletes a resource.
func (o Netscaler) DeleteLbvserver(name string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"lbvserver/"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableLbvserver disables a resource.
func (o Netscaler) DisableLbvserver(req model.LbvserverDisable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"lbvserver?action=disable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// EnableLbvserver enables a resource.
func (o Netscaler) EnableLbvserver(req model.LbvserverEnable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"lbvserver?action=enable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// RenameLbvserver renames a resource.
func (o Netscaler) RenameLbvserver(req model.LbvserverRename) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"lbvserver?action=rename", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}
