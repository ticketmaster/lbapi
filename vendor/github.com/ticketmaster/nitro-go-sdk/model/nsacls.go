package model

// Nsacls defines apply request.
type Nsacls struct {
	Nsacl NsaclsBody `json:"nsacls,omitempty"`
}

// NsaclsBody empty struct used for applying Nsacl. Why Citrix...
type NsaclsBody struct {
}

// NsaclsWrapper wraps the object and serves as default response.
type NsaclsWrapper struct {
	Errorcode int    `json:"errorcode,omitempty"`
	Message   string `json:"message,omitempty"`
	Severity  string `json:"severity,omitempty"`
}
