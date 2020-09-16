package model

// Ciphergroupname. Name for the user-defined cipher group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the cipher group is created.<br><br>The following requirement applies only to the NetScaler CLI:<br>If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ciphergroup" or my ciphergroup).<br>Minimum length = 1.
// Ciphername. Cipher name.
// Cipherpriority. This indicates priority assigned to the particular cipher.<br>Minimum value = 1.
// Ciphgrpalias. The individual cipher name(s), a user-defined cipher group, or a system predefined cipher alias that will be added to the predefined cipher alias that will be added to the group cipherGroupName.<br>If a cipher alias or a cipher group is specified, all the individual ciphers in the cipher alias or group will be added to the user-defined cipher group.<br>Minimum length = 1.
// Sslprofile. Name of the profile to which cipher is attached.

// SslcipherAdd defines add request.
type SslcipherAdd struct {
	Sslcipher SslcipherAddBody `json:"sslcipher"`
}

// SslcipherAddBody body for adding object.
type SslcipherAddBody struct {
	Ciphergroupname string `json:"ciphergroupname"`
	Ciphgrpalias    string `json:"ciphgrpalias"`
}

// SslcipherUpdateBody body for updating object.
type SslcipherUpdateBody struct {
	Ciphergroupname string `json:"ciphergroupname"`
	Ciphername      string `json:"ciphername"`
	Cipherpriority  string `json:"cipherpriority"`
}

// SslcipherUpdate defines update request.
type SslcipherUpdate struct {
	Sslcipher SslcipherUpdateBody `json:"sslcipher"`
}

// SslcipherWrapper wraps the object and serves as default response.
type SslcipherWrapper struct {
	Errorcode int         `json:"errorcode"`
	Message   string      `json:"message"`
	Severity  string      `json:"severity"`
	Sslcipher []Sslcipher `json:"sslcipher"`
}

// Sslcipher describes the object.
type Sslcipher struct {
	Ciphergroupname string `json:"ciphergroupname"`
	Ciphername      string `json:"ciphername"`
	Cipherpriority  string `json:"cipherpriority"`
	Description     string `json:"description"`
	Sslprofile      string `json:"sslprofile"`
}
