package client

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetInterfaces returns all records of the specific resource.
func (o Netscaler) GetInterfaces() (r []model.Interface, err error) {
	resp := model.InterfaceWrapper{}
	err = o.Session.Get(BaseURI+"interface", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Interface
	return
}

// GetInterface returns a single record matching the resource identfier.
func (o Netscaler) GetInterface(name string) (r model.Interface, err error) {
	resp := []model.Interface{}
	resp, err = o.GetInterfaces()
	if err != nil {
		return
	}
	mResp := make(map[string]model.Interface)
	for _, v := range resp {
		mResp[v.Devicename] = v
	}
	if mResp[name].Devicename == "" {
		err = fmt.Errorf("%s does not exist", name)
		return
	}
	r = mResp[name]
	return
}
