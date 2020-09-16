package model

// Cpufrequncy. CPU Frequency.
// Encodedserialno. Encoded serial no.
// Host. host id.
// Hostid. host id.
// Hwdescription. Hardware and its ports detail.
// Manufactureday. Manufacturing day.
// Manufacturemonth. Manufacturing month.
// Manufactureyear. Manufacturing year.
// Serialno. Serial no.
// Sysid. System id.

// NshardwareWrapper wraps the object and serves as default response
type NshardwareWrapper struct {
	Errorcode  int        `json:"errorcode"`
	Message    string     `json:"message"`
	Severity   string     `json:"severity"`
	Nshardware Nshardware `json:"nshardware"`
}

// Nshardware describes the object.
type Nshardware struct {
	Hwdescription    string `json:"hwdescription"`
	Sysid            string `json:"sysid"`
	Manufactureday   int    `json:"manufactureday"`
	Manufacturemonth int    `json:"manufacturemonth"`
	Manufactureyear  int    `json:"manufactureyear"`
	Cpufrequncy      int    `json:"cpufrequncy"`
	Hostid           int    `json:"hostid"`
	Host             string `json:"host"`
	Serialno         string `json:"serialno"`
	Encodedserialno  string `json:"encodedserialno"`
}
