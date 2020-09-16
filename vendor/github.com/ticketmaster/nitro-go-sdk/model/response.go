package model

// Response is the native Netscaler response object.
type Response struct {
	Errorcode int    `json:"errorcode"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`
}
