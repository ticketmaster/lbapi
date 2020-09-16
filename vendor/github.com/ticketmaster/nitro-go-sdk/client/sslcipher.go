package client

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetSslciphergroups returns all ssl cipher groups.
func (o Netscaler) GetSslciphergroups() (r []model.Sslcipher, err error) {
	resp := model.SslcipherWrapper{}
	err = o.Session.Get(BaseURI+"sslcipher", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Sslcipher
	return
}

// GetSslciphergroup returns a cipher group.
func (o Netscaler) GetSslciphergroup(ciphergroupname string) (r model.Sslcipher, err error) {
	resp := []model.Sslcipher{}
	resp, err = o.GetSslciphergroups()
	if err != nil {
		return
	}
	mResp := make(map[string]model.Sslcipher)
	for _, v := range resp {
		mResp[v.Ciphergroupname] = v
	}
	if mResp[ciphergroupname].Ciphergroupname == "" {
		err = fmt.Errorf("%s does not exist", ciphergroupname)
		return
	}
	r = mResp[ciphergroupname]
	return
}

// GetSslciphergroupMember returns a cipher group member.
func (o Netscaler) GetSslciphergroupMember(ciphergroupname string, ciphername string) (r model.Sslcipher, err error) {
	resp := []model.Sslcipher{}
	resp, err = o.GetSslciphergroupMembers(ciphergroupname)
	if err != nil {
		return
	}
	mResp := make(map[string]model.Sslcipher)
	for _, v := range resp {
		mResp[v.Ciphername] = v
	}
	if mResp[ciphername].Ciphername == "" {
		err = fmt.Errorf("%s does not exist", ciphername)
		return
	}
	r = mResp[ciphername]
	return
}

// GetSslciphergroupMembers returns all the cipher suites assigned to a specific cipher group.
// IMHO why Citrix allows this is beyond me because the same data is returned by
// sslcipher_sslciphersuite_binding.
func (o Netscaler) GetSslciphergroupMembers(ciphergroupname string) (r []model.Sslcipher, err error) {
	resp := model.SslcipherWrapper{}
	err = o.Session.Get(BaseURI+"sslcipher/"+ciphergroupname, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.Sslcipher
	return
}

// AddSslciphergroup creates a cipher group object.
func (o Netscaler) AddSslciphergroup(req model.SslcipherAdd) (r model.Sslcipher, err error) {
	resp := model.Response{}
	// Add cipher group.
	err = o.Session.Post(BaseURI+"sslcipher", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetSslciphergroup(req.Sslcipher.Ciphergroupname)
}

// DeleteSslciphergroup deletes a cipher group.
func (o Netscaler) DeleteSslciphergroup(ciphergroupname string) (err error) {
	resp := model.Response{}
	err = o.Session.Delete(BaseURI+"sslcipher/"+ciphergroupname, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}
