package model

// Version. Version.
// Mode. Kernel mode (KMPE/VMPE).

// NsversionWrapper wraps the object and serves as default response
type NsversionWrapper struct {
	Errorcode int       `json:"errorcode"`
	Message   string    `json:"message"`
	Severity  string    `json:"severity"`
	Nsversion Nsversion `json:"nsversion"`
}

// Nsversion describes the object.
type Nsversion struct {
	Version string `json:"version"`
	Mode    string `json:"mode"`
}
