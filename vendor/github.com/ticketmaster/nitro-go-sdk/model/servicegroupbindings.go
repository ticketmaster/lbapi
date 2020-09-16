package model

// ServicegroupbindingsWrapper wraps the object and serves as default response
type ServicegroupbindingsWrapper struct {
	Errorcode            int                   `json:"errorcode,omitempty"`
	Message              string                `json:"message,omitempty"`
	Severity             string                `json:"severity,omitempty"`
	Servicegroupbindings []Servicegroupbinding `json:"servicegroupbindings,omitempty"`
}

type Servicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname"`
	Ipaddress        string `json:"ipaddress"`
	Port             int    `json:"port"`
	State            string `json:"state"`
	Svrstate         string `json:"svrstate"`
	Vservername      string `json:"vservername"`
}
