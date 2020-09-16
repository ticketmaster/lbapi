package model

// ServicegroupLbmonitorBindingWrapper ...
type ServicegroupLbmonitorBindingWrapper struct {
	Errorcode                    int                            `json:"errorcode,omitempty"`
	Message                      string                         `json:"message,omitempty"`
	Severity                     string                         `json:"severity,omitempty"`
	ServicegroupLbmonitorBinding []ServicegroupLbmonitorBinding `json:"servicegroup_lbmonitor_binding,omitempty"`
}

// ServicegroupLbmonitorBindingWrapperAdd ...
type ServicegroupLbmonitorBindingWrapperAdd struct {
	ServicegroupLbmonitorBinding ServicegroupLbmonitorBindingPayload `json:"servicegroup_lbmonitor_binding,omitempty"`
}

// ServicegroupLbmonitorBinding ...
type ServicegroupLbmonitorBinding struct {
	Weight                     string `json:"weight,omitempty"`
	Servicegroupname           string `json:"servicegroupname,omitempty"`
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

// ServicegroupLbmonitorBindingPayload ...
type ServicegroupLbmonitorBindingPayload struct {
	Servicegroupname string `json:"servicegroupname"`
	MonitorName      string `json:"monitor_name"`
	Monstate         string `json:"monstate,omitempty"`
	Weight           string `json:"weight,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
}

// ServicegroupLbmonitorBindingAdd ...
type ServicegroupLbmonitorBindingAdd struct {
	ServicegroupLbmonitorBinding ServicegroupLbmonitorBindingPayload `json:"servicegroup_lbmonitor_binding,omitempty"`
}
