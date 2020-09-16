package client

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetAllSslcipherSslciphersuiteBindings returns bindings for all cipher groups.
func (o Netscaler) GetAllSslcipherSslciphersuiteBindings() (r []model.SslcipherSslciphersuiteBinding, err error) {
	resp := model.SslcipherSslciphersuiteBindingWrapper{}
	err = o.Session.Get(BaseURI+"sslcipher_sslciphersuite_binding?bulkbindings=yes", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.SslcipherSslciphersuiteBinding
	return
}

// GetSslcipherSslciphersuiteBindings returns bindings for a single cipher group.
func (o Netscaler) GetSslcipherSslciphersuiteBindings(ciphergroupname string) (r []model.SslcipherSslciphersuiteBinding, err error) {
	resp := model.SslcipherSslciphersuiteBindingWrapper{}
	err = o.Session.Get(BaseURI+"sslcipher_sslciphersuite_binding/"+ciphergroupname, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	r = resp.SslcipherSslciphersuiteBinding
	return
}

// GetSslcipherSslciphersuiteBinding returns a specific binding based on ciphername and ciphergroup.
func (o Netscaler) GetSslcipherSslciphersuiteBinding(ciphergroupname string, ciphername string) (r model.SslcipherSslciphersuiteBinding, err error) {
	resp := []model.SslcipherSslciphersuiteBinding{}
	resp, err = o.GetSslcipherSslciphersuiteBindings(ciphergroupname)
	if err != nil {
		return
	}
	mResp := make(map[string]model.SslcipherSslciphersuiteBinding)
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

// AddSslcipherSslciphersuiteBinding creates a cipher group binding.
func (o Netscaler) AddSslcipherSslciphersuiteBinding(req model.SslcipherSslciphersuiteBindingAdd) (r model.SslcipherSslciphersuiteBinding, err error) {
	resp := model.Response{}
	// Add cipher group binding.
	err = o.Session.Put(BaseURI+"sslcipher_sslciphersuite_binding", req, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
		return
	}
	return o.GetSslcipherSslciphersuiteBinding(req.SSLCipherSSLCipherSuiteBinding.Ciphergroupname, req.SSLCipherSSLCipherSuiteBinding.Ciphername)
}

// DeleteSslcipherSslciphersuiteBinding deletes a cipher group binding.
func (o Netscaler) DeleteSslcipherSslciphersuiteBinding(ciphergroupname string, ciphername string) (err error) {
	resp := model.Response{}
	// Add cipher group binding.
	err = o.Session.Delete(BaseURI+"sslcipher_sslciphersuite_binding?args=ciphergroupname:"+ciphergroupname+",ciphername:"+ciphername, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	return
}
