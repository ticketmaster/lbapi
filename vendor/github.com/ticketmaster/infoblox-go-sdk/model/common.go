package model

type Ipv4Addr struct {
	Ref              string `json:"_ref,omitempty"`
	ConfigureForDhcp bool   `json:"configure_for_dhcp,omitempty"`
	Host             string `json:"host,omitempty"`
	Ipv4Addr         string `json:"ipv4addr,omitempty"`
	Mac              string `json:"mac,omitempty"`
}
type ResponseError struct {
	Error string `json:"Error,omitempty"`
	Code  string `json:"code,omitempty"`
	Text  string `json:"text,omitempty"`
}
