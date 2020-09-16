package model

// NssavedconfigWrapper wraps the object and serves as default response.
type NssavedconfigWrapper struct {
	Errorcode     int           `json:"errorcode"`
	Message       string        `json:"message"`
	Severity      string        `json:"severity"`
	Nssavedconfig Nssavedconfig `json:"nssavedconfig"`
}

// Nssavedconfig describes the object.
// This is the entire configuration of the load balancer.
type Nssavedconfig struct {
	Textblob string `json:"textblob"`
}
