package model

type RecordPTRResponse struct {
	Result RecordPTRResult `json:"result,omitempty"`
}

type RecordPTRResponseWrapper struct {
	NextPageID string            `json:"next_page_id,omitempty"`
	Result     []RecordPTRResult `json:"result,omitempty"`
	Error      string            `json:"Error,omitempty"`
	Code       string            `json:"code,omitempty"`
	Text       string            `json:"text,omitempty"`
}

type RecordPTRResult struct {
	Ref      string `json:"_ref"`
	PTRDName string `json:"ptrdname"`
	Name     string `json:"name,omitempty"`
	Ipv4Addr string `json:"ipv4addr,omitempty"`
	View     string `json:"view"`
}

type RecordPTRCreateRequest struct {
	Ipv4Addr string `json:"ipv4addr"`
	PTRDName string `json:"ptrdname"`
	Name     string `json:"name,omitempty"`
}

type RecordPTRUpdateRequest struct {
	PTRDName string `json:"ptrdname"`
}
