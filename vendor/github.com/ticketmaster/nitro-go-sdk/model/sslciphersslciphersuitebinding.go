package model

// Ciphergroupname. Name for the user-defined cipher group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the cipher group is created. The following requirement applies only to the NetScaler CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ciphergroup" or my ciphergroup).Minimum length = 1.
// Ciphergrpals. A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or user defined cipher-group name. Minimum length = 1
// Ciphername. Cipher name.
// Cipheroperation. The operation that is performed when adding the cipher-suite. Possible cipher operations are: ADD - Appends the given cipher-suite to the existing one configured for the virtual server. REM - Removes the given cipher-suite from the existing one configured for the virtual server. ORD - Overrides the current configured cipher-suite for the virtual server with the given cipher-suite. Default value: 0 Possible values = ADD, REM, ORD.
// Cipherpriority. This indicates priority assigned to the particular cipher. Minimum value = 1.
// Description. Cipher suite description.

// SslcipherSslciphersuiteBindingAdd defines add request.
type SslcipherSslciphersuiteBindingAdd struct {
	SSLCipherSSLCipherSuiteBinding SslcipherSslciphersuiteBindingAddBody `json:"sslcipher_sslciphersuite_binding"`
}

// SslcipherSslciphersuiteBindingAddBody body for adding object.
type SslcipherSslciphersuiteBindingAddBody struct {
	Ciphergroupname string `json:"ciphergroupname"`
	Ciphergrpals    string `json:"ciphgrpals,omitempty"`
	Ciphername      string `json:"ciphername"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Cipherpriority  string `json:"cipherpriority"`
}

// SslcipherSslciphersuiteBindingWrapper wraps the object and serves as default response.
type SslcipherSslciphersuiteBindingWrapper struct {
	Errorcode                      int                              `json:"errorcode"`
	Message                        string                           `json:"message"`
	Severity                       string                           `json:"severity"`
	SslcipherSslciphersuiteBinding []SslcipherSslciphersuiteBinding `json:"sslcipher_sslciphersuite_binding"`
}

// SslcipherSslciphersuiteBinding describes the object.
type SslcipherSslciphersuiteBinding struct {
	Ciphergroupname string `json:"ciphergroupname"`
	Ciphername      string `json:"ciphername"`
	Cipherpriority  string `json:"cipherpriority"`
	Description     string `json:"description"`
	Peflags         string `json:"peflags"`
	Stateflag       string `json:"stateflag"`
}
