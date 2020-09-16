package model

// Bundle. Parse the certificate chain as a single file after linking the server certificate to its issuers certificate within the file.<br>Default value: NO<br>Possible values = YES, NO.
// Cert. Name of and, optionally, path to the X509 certificate file that is used to form the certificate-key pair. The certificate file should be present on the appliances hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.<br>Minimum length = 1.
// Certificatetype. Specifies whether the certificate is of type root-CA, intermediate-CA, server, client, or client and server.<br>Possible values = ROOT_CERT, INTM_CERT, CLNT_CERT, SRVR_CERT, UNKNOWN, CLIENTANDSERVER_CERT.
// Certkey. Name for the certificate and private-key pair. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the certificate-key pair is created.<br><br>The following requirement applies only to the NetScaler CLI:<br>If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cert" or my cert).<br>Minimum length = 1.
// Clientcertnotafter. Not-After date.
// Clientcertnotbefore. Not-Before date.
// Data. Vserver Id.
// Daystoexpiration. Days remaining for the certificate to expire.
// Expirymonitor. Issue an alert when the certificate is about to expire.<br>Possible values = ENABLED, DISABLED.
// Fipskey. Name of the FIPS key that was created inside the Hardware Security Module (HSM) of a FIPS appliance, or a key that was imported into the HSM.<br>Minimum length = 1.
// Hsmkey. Name of the HSM key that was created in the External Hardware Security Module (HSM) of a FIPS appliance.<br>Minimum length = 1.
// Inform. Input format of the certificate and the private-key files. The three formats supported by the appliance are:<br>PEM - Privacy Enhanced Mail<br>DER - Distinguished Encoding Rule<br>PFX - Personal Information Exchange.<br>Default value: PEM<br>Possible values = DER, PEM, PFX.
// Issuer. Issuer name.
// Key. Name of and, optionally, path to the private-key file that is used to form the certificate-key pair. The certificate file should be present on the appliances hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.<br>Minimum length = 1.
// Linkcertkeyname. Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.<br>Minimum length = 1.
// Nodomaincheck. Override the check for matching domain names during a certificate update operation.
// Notificationperiod. Time, in number of days, before certificate expiration, at which to generate an alert that the certificate is about to expire.<br>Minimum value = 10<br>Maximum value = 100.
// Ocspresponsestatus. Ocsp response status of the certificate.<br>Possible values = NONE, EXPIRED, VALID.
// Passcrypt. Passcrypt.<br>Minimum length = 1.
// Passplain. Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.<br>Minimum length = 1.
// Password. Passphrase that was used to encrypt the private-key. Use this option to load encrypted private-keys in PEM format.
// Priority. ocsp priority.
// Publickey. Public key algorithm.
// Publickeysize. Size of the public key.
// Serial. Serial number.
// Servicename. Service name to which the certificate key pair is bound.
// Signaturealg. Signature algorithm.
// Status. Status of the certificate.<br>Possible values = Valid, Not yet valid, Expired.
// Subject. Subject name.

// SslcertkeyAdd defines add request
type SslcertkeyAdd struct {
	Sslcertkey SslcertkeyAddBody `json:"sslvserver,omitempty"`
}

// SslcertkeyAddBody body for adding object.
type SslcertkeyAddBody struct {
	Certkey            string `json:"certkey"`
	Cert               string `json:"cert"`
	Key                string `json:"key"`
	Password           bool   `json:"password"`
	Fipskey            string `json:"fipskey"`
	Hsmkey             string `json:"hsmkey"`
	Inform             string `json:"inform"`
	Passplain          string `json:"passplain"`
	Expirymonitor      string `json:"expirymonitor"`
	Notificationperiod string `json:"notificationperiod"`
	Bundle             string `json:"bundle"`
}

// SslcertkeyUpdateBody body for updating object.
type SslcertkeyUpdateBody struct {
	Certkey            string `json:"certkey"`
	Expirymonitor      string `json:"expirymonitor"`
	Notificationperiod string `json:"notificationperiod"`
}

// SslcertkeyUpdate defines update request.
type SslcertkeyUpdate struct {
	Sslcertkey SslcertkeyUpdateBody `json:"sslvserver,omitempty"`
}

// SslcertkeyLink defines link request.
type SslcertkeyLink struct {
	Sslcertkey SslcertkeyLinkBody `json:"sslvserver,omitempty"`
}

// SslcertkeyLinkBody body to link.
type SslcertkeyLinkBody struct {
	Certkey         string `json:"certkey"`
	Linkcertkeyname string `json:"linkcertkeyname"`
}

// SslcertkeyUnlinkBody body to unlink.
type SslcertkeyUnlinkBody struct {
	Certkey       string `json:"certkey"`
	Cert          string `json:"cert"`
	Key           string `json:"key"`
	Password      bool   `json:"password"`
	Fipskey       string `json:"fipskey"`
	Inform        string `json:"inform"`
	Passplain     string `json:"passplain"`
	Nodomaincheck bool   `json:"nodomaincheck"`
}

// SslcertkeyUnlink defines unlink request.
type SslcertkeyUnlink struct {
	Sslcertkey SslcertkeyUnlinkBody `json:"sslvserver,omitempty"`
}

// SslcertkeyChangeBody body to update.
type SslcertkeyChangeBody struct {
	Certkey       string `json:"certkey"`
	Cert          string `json:"cert"`
	Key           string `json:"key"`
	Password      bool   `json:"password"`
	Fipskey       string `json:"fipskey"`
	Inform        string `json:"inform"`
	Passplain     string `json:"passplain"`
	Nodomaincheck bool   `json:"nodomaincheck"`
}

// SslcertkeyChange defines change request.
type SslcertkeyChange struct {
	Sslcertkey SslcertkeyChangeBody `json:"sslvserver,omitempty"`
}

// SslcertkeyWrapper wraps the object and serves as default response.
type SslcertkeyWrapper struct {
	Errorcode  int          `json:"errorcode,omitempty"`
	Message    string       `json:"message,omitempty"`
	Severity   string       `json:"severity,omitempty"`
	Sslcertkey []Sslcertkey `json:"sslvserver,omitempty"`
}

// Sslcertkey describes the object.
type Sslcertkey struct {
	Certkey             string `json:"certkey"`
	Cert                string `json:"cert"`
	Key                 string `json:"key"`
	Inform              string `json:"inform"`
	Signaturealg        string `json:"signaturealg"`
	Certificatetype     string `json:"certificatetype"`
	Serial              string `json:"serial"`
	Issuer              string `json:"issuer"`
	Clientcertnotbefore string `json:"clientcertnotbefore"`
	Clientcertnotafter  string `json:"clientcertnotafter"`
	Daystoexpiration    int    `json:"daystoexpiration"`
	Subject             string `json:"subject"`
	Publickey           string `json:"publickey"`
	Publickeysize       int    `json:"publickeysize"`
	Version             int    `json:"version"`
	Priority            string `json:"priority"`
	Status              string `json:"status"`
	Fipskey             string `json:"fipskey"`
	Hsmkey              string `json:"hsmkey"`
	Passcrypt           string `json:"passcrypt"`
	Data                string `json:"data"`
	Servicename         string `json:"servicename"`
	Expirymonitor       string `json:"expirymonitor"`
	Notificationperiod  string `json:"notificationperiod"`
	Linkcertkeyname     string `json:"linkcertkeyname"`
	Ocspresponsestatus  string `json:"ocspresponsestatus"`
}
