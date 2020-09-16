package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetServers returns all records of the specific resource.
func (o Netscaler) GetServers() (r []model.Server, err error) {
	resp := model.ServerWrapper{}
	err = o.Session.Get(BaseURI+"server", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Server
	return
}

// GetServer returns a single record matching the resource identfier.
func (o Netscaler) GetServer(name string) (r model.Server, err error) {
	resp := model.ServerWrapper{}
	err = o.Session.Get(BaseURI+"server/"+name, &resp)
	if err != nil || len(resp.Server) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Server[0]
	return
}

// GetServerByIP returns a single record matching the resource identfier.
func (o Netscaler) GetServerByIP(ip string) (r model.Server, err error) {
	resp := model.ServerWrapper{}
	err = o.Session.Get(BaseURI+"server?filter=ipaddress:"+ip, &resp)
	if err != nil || len(resp.Server) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Server[0]
	return
}

// AddServer creates a resource.
func (o Netscaler) AddServer(req model.ServerAdd) (r model.Server, err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"server", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetServer(req.Server.Name)
}

// UpdateServer updates a resource.
func (o Netscaler) UpdateServer(req model.ServerUpdate) (r model.Server, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"server", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetServer(req.Server.Name)
}

// DeleteServer deletes a resource.
func (o Netscaler) DeleteServer(name string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"server/"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableServer disables a resource.
func (o Netscaler) DisableServer(req model.ServerDisable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"server?action=disable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// EnableServer enables a resource.
func (o Netscaler) EnableServer(req model.ServerEnable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"server?action=enable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// RenameServer renames a resource.
func (o Netscaler) RenameServer(req model.ServerRename) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"server?action=rename", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}
