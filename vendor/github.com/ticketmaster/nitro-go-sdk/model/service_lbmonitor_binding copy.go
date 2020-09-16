package model

// ServiceLbmonitorBindingWrapper ...
type ServiceLbmonitorBindingWrapper struct {
	Errorcode               int                       `json:"errorcode,omitempty"`
	Message                 string                    `json:"message,omitempty"`
	Severity                string                    `json:"severity,omitempty"`
	ServiceLbmonitorBinding []ServiceLbmonitorBinding `json:"service_lbmonitor_binding,omitempty"`
}

// ServiceLbmonitorBindingWrapperAdd ...
type ServiceLbmonitorBindingWrapperAdd struct {
	ServiceLbmonitorBinding ServiceLbmonitorBindingPayload `json:"service_lbmonitor_binding,omitempty"`
}

// ServiceLbmonitorBinding ...
type ServiceLbmonitorBinding struct {
	Weight                     string `json:"weight,omitempty"`
	Name                       string `json:"name,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	MonitorName                string `json:"monitor_name,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Monstatcode                int    `json:"monstatcode,omitempty"`
	DupWeight                  string `json:"dup_weight,omitempty"`
	Responsetime               string `json:"responsetime,omitempty"`
	Totalfailedprobes          string `json:"totalfailedprobes,omitempty"`
	Monstatparam2              int    `json:"monstatparam2,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Failedprobes               string `json:"failedprobes,omitempty"`
	Monstatparam3              int    `json:"monstatparam3,omitempty"`
	Totalprobes                string `json:"totalprobes,omitempty"`
	MonitorState               string `json:"monitor_state,omitempty"`
	DupState                   string `json:"dup_state,omitempty"`
	Monitortotalprobes         string `json:"monitortotalprobes,omitempty"`
	Monstatparam1              int    `json:"monstatparam1,omitempty"`
	Monitortotalfailedprobes   string `json:"monitortotalfailedprobes,omitempty"`
	Monitorcurrentfailedprobes string `json:"monitorcurrentfailedprobes,omitempty"`
}

// ServiceLbmonitorBindingPayload ...
type ServiceLbmonitorBindingPayload struct {
	Name        string `json:"name"`
	MonitorName string `json:"monitor_name"`
	Monstate    string `json:"monstate,omitempty"`
	Weight      string `json:"weight,omitempty"`
	Passive     bool   `json:"passive,omitempty"`
}

// ServiceLbmonitorBindingAdd ...
type ServiceLbmonitorBindingAdd struct {
	ServiceLbmonitorBinding ServiceLbmonitorBindingPayload `json:"service_lbmonitor_binding,omitempty"`
}
