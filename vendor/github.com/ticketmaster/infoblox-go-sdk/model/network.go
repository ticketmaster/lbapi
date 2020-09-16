package model

type NextAvailable struct {
	Ips []string `json:"ips"`
}

type NextAvailableCount struct {
	Num int `json:"num"`
}

type NetworkResponse struct {
	Result NetworkResult `json:"result,omitempty"`
}

type NetworkResponseWrapper struct {
	NextPageID string          `json:"next_page_id,omitempty"`
	Result     []NetworkResult `json:"result,omitempty"`
	Error      string          `json:"Error,omitempty"`
	Code       string          `json:"code,omitempty"`
	Text       string          `json:"text,omitempty"`
}

type NetworkResult struct {
	Ref         string `json:"_ref"`
	Comment     string `json:"comment"`
	Network     string `json:"network"`
	NetworkView string `json:"network_view"`
}
