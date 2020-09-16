package client

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// GetSslciphersuites returns all cipher suites.
func GetSslciphersuites(s *Session) (r []model.Sslciphersuite, err error) {
	resp := model.SslciphersuiteWrapper{}
	err = s.Get(BaseURI+"sslciphersuite", &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	r = resp.Sslciphersuite
	return
}

// GetSslciphersuite returns a specific ciphersuite.
func GetSslciphersuite(s *Session, ciphername string) (r model.Sslciphersuite, err error) {
	resp := model.SslciphersuiteWrapper{}
	err = s.Get(BaseURI+"sslciphersuite/"+ciphername, &resp)
	if err != nil {
		return
	}
	if resp.Errorcode != 0 {
		err = errors.New(resp.Message)
	}
	r = resp.Sslciphersuite[0]
	return
}
