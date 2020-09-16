package model

type RangeResponse struct {
	Result RangeResult `json:"result,omitempty"`
}

type RangeResponseWrapper struct {
	NextPageID string        `json:"next_page_id,omitempty"`
	Result     []RangeResult `json:"result,omitempty"`
	Error      string        `json:"Error,omitempty"`
	Code       string        `json:"code,omitempty"`
	Text       string        `json:"text,omitempty"`
}

type RangeResult struct {
	Ref         string `json:"_ref"`
	Comment     string `json:"comment,omitempty"`
	EndAddr     string `json:"end_addr"`
	Network     string `json:"network"`
	NetworkView string `json:"network_view"`
	StartAddr   string `json:"start_addr"`
}
