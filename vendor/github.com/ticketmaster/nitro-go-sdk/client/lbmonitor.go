package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetLbmonitors returns all records of the specific resource.
func (o Netscaler) GetLbmonitors() (r []model.Lbmonitor, err error) {
	resp := model.LbmonitorWrapper{}
	err = o.Session.Get(BaseURI+"lbmonitor", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Lbmonitor
	return
}

// GetLbmonitor returns a single record matching the resource identfier.
func (o Netscaler) GetLbmonitor(name string) (r model.Lbmonitor, err error) {
	resp := model.LbmonitorWrapper{}
	err = o.Session.Get(BaseURI+"lbmonitor/"+name, &resp)
	if err != nil || len(resp.Lbmonitor) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Lbmonitor[0]
	return
}

// AddLbmonitor creates a resource.
func (o Netscaler) AddLbmonitor(req model.LbmonitorAdd) (r model.Lbmonitor, err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"lbmonitor", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetLbmonitor(req.Lbmonitor.Monitorname)
}

// UpdateLbmonitor updates a resource.
func (o Netscaler) UpdateLbmonitor(req model.LbmonitorUpdate) (r model.Lbmonitor, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"lbmonitor", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetLbmonitor(req.Lbmonitor.Monitorname)
}

// DeleteLbmonitor deletes a resource.
func (o Netscaler) DeleteLbmonitor(name string, montype string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"lbmonitor/"+name+"?args=type:"+montype, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableLbmonitor disables a resource.
func (o Netscaler) DisableLbmonitor(req model.LbmonitorDisable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"lbmonitor?action=disable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// EnableLbmonitor enables a resource.
func (o Netscaler) EnableLbmonitor(req model.LbmonitorEnable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"lbmonitor?action=enable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}
