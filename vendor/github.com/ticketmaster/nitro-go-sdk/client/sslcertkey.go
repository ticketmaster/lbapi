package client

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// AddSslcertkey creates a resource.
func (o Netscaler) AddSslcertkey(req model.SslcertkeyAdd) (r model.Sslcertkey, err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"sslcertkey", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetSslcertkey(req.Sslcertkey.Certkey)
}

// DeleteSslcertkey deletes a resource.
func (o Netscaler) DeleteSslcertkey(certkey string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"sslcertkey/"+certkey, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// LinkSslcertkey deletes a resource.
func (o Netscaler) LinkSslcertkey(req model.SslcertkeyLink) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"sslcertkey?action=link", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// UnlinkSslcertkey deletes a resource.
func (o Netscaler) UnlinkSslcertkey(req model.SslcertkeyUnlink) (err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"sslcertkey?action=unlink", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}

// ChangeSslcertkey deletes a resource.
func (o Netscaler) ChangeSslcertkey(req model.SslcertkeyChange) (r model.Sslcertkey, err error) {
	resp := model.Response{}
	err = o.Session.Post(BaseURI+"sslcertkey?action=change", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return o.GetSslcertkey(req.Sslcertkey.Certkey)
}

// GetSslcertkeys returns all vservers.
func (o Netscaler) GetSslcertkeys() (r []model.Sslcertkey, err error) {
	resp := model.SslcertkeyWrapper{}
	err = o.Session.Get(BaseURI+"sslcertkey", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Sslcertkey
	return
}

// GetSslcertkey returns a vserver.
func (o Netscaler) GetSslcertkey(certkey string) (r model.Sslcertkey, err error) {
	resp := []model.Sslcertkey{}
	resp, err = o.GetSslcertkeys()
	if err != nil {
		return
	}
	mResp := make(map[string]model.Sslcertkey)
	for _, v := range resp {
		mResp[v.Certkey] = v
	}
	if mResp[certkey].Certkey == "" {
		err = fmt.Errorf("%s does not exist", certkey)
		return
	}
	r = mResp[certkey]
	return
}

// UpdateSslcertkey updates the vserver to support ssl.
func (o Netscaler) UpdateSslcertkey(req model.SslcertkeyUpdate) (r model.Sslcertkey, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"sslcertkey", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetSslcertkey(req.Sslcertkey.Certkey)
}
