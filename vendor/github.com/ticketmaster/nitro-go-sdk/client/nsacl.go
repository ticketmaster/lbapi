package client

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetNsacls returns all records of the specific resource.
func (o Netscaler) GetNsacls() (r []model.Nsacl, err error) {
	resp := model.NsaclWrapper{}
	err = o.Session.Get(BaseURI+"nsacl", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Nsacl
	return
}

// GetNsacl returns a single record matching the resource identfier.
func (o Netscaler) GetNsacl(aclname string) (r model.Nsacl, err error) {
	resp := []model.Nsacl{}
	resp, err = o.GetNsacls()
	if err != nil {
		return
	}
	mResp := make(map[string]model.Nsacl)
	for _, v := range resp {
		mResp[v.Aclname] = v
	}
	if mResp[aclname].Aclname == "" {
		err = fmt.Errorf("%s does not exist", aclname)
		return
	}
	r = mResp[aclname]
	return
}

// AddNsacl creates a resource.
func (o Netscaler) AddNsacl(req model.NsaclAdd) (r model.Nsacl, err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"nsacl", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetNsacl(req.Nsacl.Aclname)
}

// UpdateNsacl updates a resource.
func (o Netscaler) UpdateNsacl(req model.NsaclUpdate) (r model.Nsacl, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"nsacl", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetNsacl(req.Nsacl.Aclname)
}

// DeleteNsacl deletes a resource.
func (o Netscaler) DeleteNsacl(aclname string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"nsacl/"+aclname, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// DisableNsacl disables a resource.
func (o Netscaler) DisableNsacl(req model.NsaclDisable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"nsacl?action=disable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// EnableNsacl enables a resource.
func (o Netscaler) EnableNsacl(req model.NsaclEnable) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"nsacl?action=enable", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// RenameNsacl renames a resource.
func (o Netscaler) RenameNsacl(req model.NsaclRename) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"nsacl?action=rename", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}
