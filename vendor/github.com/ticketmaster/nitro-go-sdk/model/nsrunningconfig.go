package model

// NsrunningconfigWrapper wraps the object and serves as default response.
type NsrunningconfigWrapper struct {
	Errorcode       int             `json:"errorcode"`
	Message         string          `json:"message"`
	Severity        string          `json:"severity"`
	Nsrunningconfig Nsrunningconfig `json:"Nsrunningconfig"`
}

// Nsrunningconfig describes the object.
// This is the entire configuration of the load balancer.
type Nsrunningconfig struct {
	Response string `json:"response"`
}
