package model

// SvcbindingsWrapper wraps the object and serves as default response
type SvcbindingsWrapper struct {
	Errorcode   int          `json:"errorcode,omitempty"`
	Message     string       `json:"message,omitempty"`
	Severity    string       `json:"severity,omitempty"`
	Svcbindings []Svcbinding `json:"svcbindings,omitempty"`
}

type Svcbinding struct {
	Servicename string `json:"servicename"`
	Ipaddress   string `json:"ipaddress"`
	Port        int    `json:"port"`
	Svrstate    string `json:"svrstate"`
	Vservername string `json:"vservername"`
}
