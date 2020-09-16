package model

type RecordAResponse struct {
	Result RecordAResult `json:"result,omitempty"`
}
type RecordAResponseWrapper struct {
	NextPageID string          `json:"next_page_id,omitempty"`
	Result     []RecordAResult `json:"result,omitempty"`
	Error      string          `json:"Error,omitempty"`
	Code       string          `json:"code,omitempty"`
	Text       string          `json:"text,omitempty"`
}
type RecordAResult struct {
	Ref      string `json:"_ref"`
	Ipv4Addr string `json:"ipv4addr"`
	Name     string `json:"name"`
	View     string `json:"view"`
}
type RecordACreateRequest struct {
	Ipv4Addr string `json:"ipv4addr"`
	Name     string `json:"name,omitempty"`
}
type RecordAUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Ipv4Addr string `json:"ipv4addr"`
}
