package certificate

import (
	"fmt"

	"github.com/ticketmaster/lbapi/shared"
	"github.com/avinetworks/sdk/go/models"
)

func (o *Avi) etlCreate(data *Data) AviCertificate {
	return AviCertificate{
		Certificate: PublicKey{
			Certificate: data.Certificate,
		},
		Key:           data.Key.PrivateKey,
		KeyPassphrase: data.Key.PassPhrase,
		Name:          data.Name,
	}
}
func (o *Avi) etlModify(data *Data) (r *models.SSLKeyAndCertificate, err error) {
	////////////////////////////////////////////////////////////////////////////
	r, err = o.Client.SSLKeyAndCertificate.GetByName(data.Name)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if *r.Certificate.SelfSigned {
		err = fmt.Errorf("self-signed certificates cannot be modified")
		return
	}
	////////////////////////////////////////////////////////////////////////////
	r.Certificate.Certificate = &data.Certificate
	r.Key = &data.Key.PrivateKey
	r.KeyPassphrase = &data.Key.PassPhrase
	return
}
func (o *Avi) etlFetch(in *models.SSLKeyAndCertificate) (data *Data, err error) {
	var rsasize string
	var eccurve string
	var keyalgorithm string
	if in.KeyParams != nil {
		if in.KeyParams.RsaParams != nil {
			rsasize = *in.KeyParams.RsaParams.KeySize
		}
		////////////////////////////////////////////////////////////////////////////
		if in.KeyParams.EcParams != nil {
			eccurve = *in.KeyParams.EcParams.Curve
		}
		if in.KeyParams.Algorithm != nil {
			keyalgorithm = *in.KeyParams.Algorithm
		}
	}
	////////////////////////////////////////////////////////////////////////////
	data = &Data{
		SourceUUID:               shared.FormatAviRef(*in.UUID),
		SourceSelfSigned:         *in.Certificate.SelfSigned,
		SourceCommonName:         *in.Certificate.Subject.CommonName,
		Name:                     *in.Name,
		SourceExpiry:             *in.Certificate.NotAfter,
		Certificate:              *in.Certificate.Certificate,
		SourceDistinguishedName:  *in.Certificate.Subject.DistinguishedName,
		SourceSerialNumber:       *in.Certificate.SerialNumber,
		SourceSignatureAlgorithm: *in.Certificate.SignatureAlgorithm,
		Key: Key{
			SourceAlgorithm: keyalgorithm,
			SourceRSASize:   rsasize,
			SourceECCurve:   eccurve,
		},
	}
	return
}
