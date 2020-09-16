package model

type RecordCnameResponse struct {
	Result RecordCnameResult `json:"result,omitempty"`
}

type RecordCnameResponseWrapper struct {
	NextPageID string              `json:"next_page_id,omitempty"`
	Result     []RecordCnameResult `json:"result,omitempty"`
	Error      string              `json:"Error,omitempty"`
	Code       string              `json:"code,omitempty"`
	Text       string              `json:"text,omitempty"`
}

type RecordCnameResult struct {
	Ref string `json:"_ref"`
	// Canonical is the A record reference name.
	Canonical string `json:"canonical"`
	Name      string `json:"name"`
	View      string `json:"view"`
}

type RecordCnameCreateRequest struct {
	// Canonical is the A record reference name.
	Canonical string `json:"canonical"`
	Name      string `json:"name,omitempty"`
}

type RecordCnameUpdateRequest struct {
	Name string `json:"name,omitempty"`
}
