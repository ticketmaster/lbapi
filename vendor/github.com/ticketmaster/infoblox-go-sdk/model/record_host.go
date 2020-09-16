package model

type RecordHostResponse struct {
	Result RecordHostResult `json:"result,omitempty"`
}

type RecordHostResponseWrapper struct {
	NextPageID string             `json:"next_page_id,omitempty"`
	Result     []RecordHostResult `json:"result,omitempty"`
	Error      string             `json:"Error,omitempty"`
	Code       string             `json:"code,omitempty"`
	Text       string             `json:"text,omitempty"`
}

type RecordHostResult struct {
	Ref       string     `json:"_ref,omitempty"`
	Name      string     `json:"name,omitempty"`
	Aliases   []string   `json:"aliases,omitempty"`
	Ipv4Addrs []Ipv4Addr `json:"ipv4addrs,omitempty"`
	View      string     `json:"view,omitempty"`
}

type RecordHostCreateRequest struct {
	Name      string     `json:"name,omitempty"`
	Aliases   []string   `json:"aliases,omitempty"`
	Ipv4Addrs []Ipv4Addr `json:"ipv4addrs,omitempty"`
	View      string     `json:"view,omitempty"`
}

type RecordHostUpdateRequest struct {
	Name      string     `json:"name,omitempty"`
	Aliases   []string   `json:"aliases,omitempty"`
	Ipv4Addrs []Ipv4Addr `json:"ipv4addrs,omitempty"`
	View      string     `json:"view,omitempty"`
}
