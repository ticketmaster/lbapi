package client

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetHanodes returns all hanodes.
func (o Netscaler) GetHanodes() (r []model.Hanode, err error) {
	resp := model.HanodeWrapper{}
	err = o.Session.Get(BaseURI+"Hanode", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Hanode
	return
}

// GetHanode returns a specific hanode.
func (o Netscaler) GetHanode(ipaddress string) (r model.Hanode, err error) {
	resp := []model.Hanode{}
	resp, err = o.GetHanodes()
	if err != nil {
		return
	}
	mResp := make(map[string]model.Hanode)
	for _, v := range resp {
		mResp[v.Ipaddress] = v
	}
	if mResp[ipaddress].Ipaddress == "" {
		err = fmt.Errorf("%s does not exist", ipaddress)
		return
	}
	r = mResp[ipaddress]
	return
}
