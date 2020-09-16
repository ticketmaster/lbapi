package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetNsips returns all records of the specific resource.
func (o Netscaler) GetNsips() (r []model.Nsip, err error) {
	resp := model.NsipWrapper{}
	err = o.Session.Get(BaseURI+"nsip", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Nsip
	return
}

// GetNsip returns a single record matching the resource identfier.
func (o Netscaler) GetNsip(ipaddress string) (r model.Nsip, err error) {
	resp := model.NsipWrapper{}
	err = o.Session.Get(BaseURI+"nsip/"+ipaddress, &resp)
	if err != nil || len(resp.Nsip) == 0 {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Nsip[0]
	return
}

// UpdateNsip updates a resource.
func (o Netscaler) UpdateNsip(req model.NsipUpdate) (r model.Nsip, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"nsip", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetNsip(req.Nsip.Ipaddress)
}

// DeleteNsip deletes a resource.
func (o Netscaler) DeleteNsip(name string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"nsip/"+name, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableNsip disables a resource.
func (o Netscaler) DisableNsip(req model.NsipDisable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"nsip?action=disable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// EnableNsip enables a resource.
func (o Netscaler) EnableNsip(req model.NsipEnable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"nsip?action=enable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}
