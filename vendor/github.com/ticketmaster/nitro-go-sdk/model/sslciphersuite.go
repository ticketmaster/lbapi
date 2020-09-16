package model

// Ciphername. Name of the cipher suite for which to show detailed information.
// Description. Cipher suite description.

// SslciphersuiteWrapper wraps the object and serves as default response
type SslciphersuiteWrapper struct {
	Errorcode      int              `json:"errorcode"`
	Message        string           `json:"message"`
	Severity       string           `json:"severity"`
	Sslciphersuite []Sslciphersuite `json:"sslciphersuite"`
}

// Sslciphersuite  describes the object.
type Sslciphersuite struct {
	Ciphername  string `json:"ciphername"`
	Description string `json:"description"`
}
