package certificate

// Certificate - implements certificate package logic.
type Certificate struct {
	// Avi - implements Avi logic.
	Avi *Avi
	// Netscaler - implements Netscaler logic.
	// Netscaler *Netscaler
}

// Data - resource configuration.
type Data struct {
	// Name - friendly name of the resource. Typically inherited from parent.
	Name string `json:"name,omitempty"`
	// PublicKey - PEM formated certificate. Replace line breaks with "\n".
	Certificate string `json:"certificate,omitempty"`
	// Key - resource configuration.
	Key Key `json:"key,omitempty"`
	// SourceCommonName [system] - common name of certificate.
	SourceCommonName string `json:"_common_name,omitempty"`
	// SourceSerialNumber [system] - serial number of the certificate.
	SourceSerialNumber string `json:"_serial_number,omitempty"`
	// SourceDistinguishedName [system] - distinguished name of certificate.
	SourceDistinguishedName string `json:"_distinguished_name,omitempty"`
	// SourceExpiry [system] - date the cert expires.
	SourceExpiry string `json:"_expiry,omitempty"`
	// SourceSignatureAlgorithm [system] - signature algorithm used to sign certificate.
	SourceSignatureAlgorithm string `json:"_signature_algorithm,omitempty"`
	// SourceSelfSigned [system] - true if certificate is self-signed.
	SourceSelfSigned bool `json:"_self_signed,omitempty"`
	// SourceUUID [system] - record id of the resource.
	SourceUUID string `json:"_uuid,omitempty"`
}

// Key - resource configuration.
type Key struct {
	//PrivateKey [required on creation] - PEM formatted private key. Replace line breaks with "\n".
	PrivateKey string `json:"private_key,omitempty"`
	// PassPhrase [optional] - secret for decrypting key.
	PassPhrase string `json:"passphrase,omitempty"`
	// SourceRSASize [system] - size of RSA key. Applicable to RSA certs.
	SourceRSASize string `json:"_rsa_size,omitempty"`
	// SourceECCurve [system] - eccurve used to encrypt certificate. Applicable to ECC certs.
	SourceECCurve string `json:"_ec_curve,omitempty"`
	// SourceAlgorithm [system] - algorithm used to encrypt certificate.
	SourceAlgorithm string `json:"_algorithm,omitempty"`
}

// AviCertificate - Struct for manipulating Avi Certificates.
type AviCertificate struct {
	// Certificate - resource configuration".
	Certificate PublicKey `json:"certificate"`
	// TenantRef - Avi URL for tenant.
	TenantRef string `json:"tenant_ref,omitempty"`
	// Name - friendly name of the resource. Typically inherited from parent.
	Name string `json:"name"`
	//Key [required on creation] - PEM formatted private key.
	Key string `json:"key"`
	// KeyPassphrase [optional] - secret for decrypting key.
	KeyPassphrase string `json:"key_passphrase,omitempty"`
}

// PublicKey - resource configuration.
type PublicKey struct {
	// Certificate - PEM formated certificate.
	Certificate string `json:"certificate"`
}

// CertificateCollection ...
type Collection struct {
	// Source - map of all related records indexed by uuid from source.
	Source map[string]Data
	// System - map of all related records indexed by name from source.
	System map[string]Data
}
