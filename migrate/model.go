package migrate

type Response struct {
	SourceID           string  `json:"source_id"`
	ProductCode        int     `json:"product_code"`
	Source             Wrapper `json:"source"`
	Target             Wrapper `json:"destination"`
	TargetLoadBalancer string  `json:"target_load_balancer"`
	SourceLoadBalancer string  `json:"source_load_balancer"`
	Platform           string  `json:"platform"`
	ReadinessChecks    struct {
		Ready            bool             `json:"ready"`
		IPStatus         IPStatus         `json:"ip_status"`
		DependencyStatus DependencyStatus `json:"dependency_status"`
		NetworkStatus    NetworkStatus    `json:"network_status"`
		Pools            []PoolStatus     `json:"pool_status"`
		LoadBalancer     bool             `json:"load_balancer"`
		Error            string           `json:"error,omitempty"`
	} `json:"readiness_checks"`
}

type DependencyStatus struct {
	Ready bool     `json:"ready"`
	Error string   `json:"error,omitempty"`
	IPs   []string `json:"ips,omitempty"`
}

type Wrapper struct {
	VirtualServer interface{} `json:"data"`
}

type HealthMonitorStatus struct {
	Type  string `json:"type"`
	Ready bool   `json:"ready"`
	Error string `json:"error,omitempty"`
}

type PoolStatus struct {
	Ready          bool                  `json:"ready"`
	HealthMonitors []HealthMonitorStatus `json:"health_monitors"`
	Servers        []ServerStatus        `json:"servers"`
	Persistence    bool                  `json:"persistence"`
	Error          string                `json:"error,omitempty"`
}

type ServerStatus struct {
	Ready bool   `json:"ready"`
	IP    string `json:"ip"`
	Error string `json:"error,omitempty"`
}

type IPStatus struct {
	Ready bool   `json:"ready"`
	IP    string `json:"ip"`
	Error string `json:"error,omitempty"`
}

type NetworkStatus struct {
	Ready       bool   `json:"ready"`
	Port        int    `json:"port"`
	ServiceType string `json:"service_type"`
	Error       string `json:"error,omitempty"`
}
