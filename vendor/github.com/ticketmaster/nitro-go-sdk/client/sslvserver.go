package client

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetSslvservers returns all vservers.
func (o Netscaler) GetSslvservers() (r []model.Sslvserver, err error) {
	resp := model.SslvserverWrapper{}
	err = o.Session.Get(BaseURI+"sslvserver", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.SSLVServer
	return
}

// GetSslvserver returns a vserver.
func (o Netscaler) GetSslvserver(vservername string) (r model.Sslvserver, err error) {
	resp := []model.Sslvserver{}
	resp, err = o.GetSslvservers()
	if err != nil {
		return
	}
	mResp := make(map[string]model.Sslvserver)
	for _, v := range resp {
		mResp[v.Vservername] = v
	}
	if mResp[vservername].Vservername == "" {
		err = fmt.Errorf("%s does not exist", vservername)
		return
	}
	r = mResp[vservername]
	return
}

// UpdateSslvserver updates the vserver to support ssl.
func (o Netscaler) UpdateSslvserver(req model.SslvserverUpdate) (r model.Sslvserver, err error) {
	resp := model.Response{}
	err = o.Session.Put(BaseURI+"sslvserver", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetSslvserver(req.Sslvserver.Vservername)
}
