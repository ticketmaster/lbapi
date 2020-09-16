package infoblox

type gqlResponse struct {
	Data struct {
		EndpointIPAddresses struct {
			Edges []struct {
				Node struct {
					EndpointIPAddress string `json:"endpointIpAddress"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"endpointIpAddresses"`
	} `json:"data"`
}

type network struct {
	Cidr string
	Ref  string
}

type host struct {
	Host string `json:"host"`
}

type huntList struct {
	Cidr    string  `json:"cidr"`
	Entries []entry `json:"entries"`
}

type entry struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}
