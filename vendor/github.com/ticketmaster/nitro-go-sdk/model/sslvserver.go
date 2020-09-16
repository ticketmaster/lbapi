package model

// Cipherurl. The redirect URL to be used with the Cipher Redirect feature.
// Clientcert. Type of client authentication. If this parameter is set to MANDATORY, the appliance terminates the SSL handshake if the SSL client does not provide a valid certificate. With the OPTIONAL setting, the appliance requests a certificate from the SSL clients but proceeds with the SSL transaction even if the client presents an invalid certificate.<br>Caution: Define proper access control policies before changing this setting to Optional.<br>Possible values = Mandatory, Optional
// Dhfile. Name of and, optionally, path to the DH parameter file, in PEM format, to be installed. /nsconfig/ssl/ is the default path.<br>Minimum length = 1
// Dtlsprofilename. Name of the DTLS profile whose settings are to be applied to the virtual server.<br>Minimum length = 1<br>Maximum length = 127
// Hsts. State of TLSv1.0 protocol support for the SSL Virtual Server.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Includesubdomains. Set include sub domain value for STS header.<br>Default value: NO<br>Possible values = YES, NO
// Maxage. Set max-age value for STS header.<br>Default value: 0<br>Minimum value = 0<br>Maximum value = 4294967294
// Sslprofile. Name of the SSL profile that contains SSL settings for the virtual server.<br>Minimum length = 1<br>Maximum length = 127
// Sslv2url. URL of the page to which to redirect the client in case of a protocol version mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from.
// Ca. CA certificate.
// Cipherredirect. State of Cipher Redirect. If cipher redirect is enabled, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Cleartextport. Port on which clear-text data is sent by the appliance to the server. Do not specify this parameter for SSL offloading with end-to-end encryption.<br>Default value: 0<br>Minimum value = 0<br>Maximum value = 65534
// Clientauth. State of client authentication. If client authentication is enabled, the virtual server terminates the SSL handshake if the SSL client does not provide a valid certificate.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Dh. State of Diffie-Hellman (DH) key exchange.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Dhcount. Number of interactions, between the client and the NetScaler appliance, after which the DH private-public pair is regenerated. A value of zero (0) specifies infinite use (no refresh).<br>Minimum value = 0<br>Maximum value = 65534
// Dhkeyexpsizelimit. This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Dtlsflag. The flag is used to indicate whether DTLS is set or not.
// Ersa. State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// Ersacount. Refresh count for regeneration of the RSA public-key and private-key pair. Zero (0) specifies infinite usage (no refresh).<br>Minimum value = 0<br>Maximum value = 65534
// Nonfipsciphers. The state of usage of non FIPS approved ciphers.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Ocspstapling. State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values:<br>ENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake.<br>DISABLED: The appliance does not check the status of the server certificate. .<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Pushenctrigger. Trigger encryption on the basis of the PUSH flag value. Available settings function as follows:<br>* ALWAYS - Any PUSH packet triggers encryption.<br>* IGNORE - Ignore PUSH packet for triggering encryption.<br>* MERGE - For a consecutive sequence of PUSH packets, the last PUSH packet triggers encryption.<br>* TIMER - PUSH packet triggering encryption is delayed by the time defined in the set ssl parameter command or in the Change Advanced SSL Settings dialog box.<br>Possible values = Always, Merge, Ignore, Timer
// Redirectportrewrite. State of the port rewrite while performing HTTPS redirect. If this parameter is ENABLED and the URL from the server does not contain the standard port, the port is rewritten to the standard.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Sendclosenotify. Enable sending SSL Close-Notify at the end of a transaction.<br>Default value: YES<br>Possible values = YES, NO
// Service. Service.<br>Minimum value = 0<br>Maximum value = 65534
// Sessreuse. State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// Sesstimeout. Time, in seconds, for which to keep the session active. Any session resumption request received after the timeout period will require a fresh SSL handshake and establishment of a new SSL session.<br>Default value: 120<br>Minimum value = 0<br>Maximum value = 4294967294
// Skipcaname. The flag is used to indicate whether this particular CA certificates CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.
// Snicert. The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.
// Snienable. State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Ssl2. State of SSLv2 protocol support for the SSL Virtual Server.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Ssl3. State of SSLv3 protocol support for the SSL Virtual Server.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// Sslredirect. State of HTTPS redirects for the SSL virtual server. <br><br>For an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect.<br>If SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Sslv2Redirect. State of SSLv2 Redirect. If SSLv2 redirect is enabled, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a protocol version mismatch between the virtual server or service and the client.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// Strictsigdigestcheck. Parameter indicating to check whether peer entity certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Netscaler.<br>Default value: DISABLED<br>Possible values = ENABLED, DISABLED
// TLS1. State of TLSv1.0 protocol support for the SSL Virtual Server.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// TLS11. State of TLSv1.1 protocol support for the SSL Virtual Server. TLSv1.1 protocol is supported only on the MPX appliance. Support is not available on a FIPS appliance or on a NetScaler VPX virtual appliance. On an SDX appliance, TLSv1.1 protocol is supported only if an SSL chip is assigned to the instance.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// TLS12. State of TLSv1.2 protocol support for the SSL Virtual Server. TLSv1.2 protocol is supported only on the MPX appliance. Support is not available on a FIPS appliance or on a NetScaler VPX virtual appliance. On an SDX appliance, TLSv1.2 protocol is supported only if an SSL chip is assigned to the instance.<br>Default value: ENABLED<br>Possible values = ENABLED, DISABLED
// Vservername. Name of the SSL virtual server for which to set advanced configuration.<br>Minimum length = 1

// SslvserverUpdate defines update request.
type SslvserverUpdate struct {
	Sslvserver SslvserverUpdateBody `json:"sslvserver,omitempty"`
}

// SslvserverUpdateBody body for updating object.
type SslvserverUpdateBody struct {
	Cipherredirect       string `json:"cipherredirect,omitempty"`
	Cipherurl            string `json:"cipherurl,omitempty"`
	Cleartextport        int    `json:"cleartextport,omitempty"`
	Clientauth           string `json:"clientauth,omitempty"`
	Clientcert           string `json:"clientcert,omitempty"`
	Dh                   string `json:"dh,omitempty"`
	Dhcount              string `json:"dhcount,omitempty"`
	Dhfile               string `json:"dhfile,omitempty"`
	Dhkeyexpsizelimit    string `json:"dhkeyexpsizelimit,omitempty"`
	Dtlsprofilename      string `json:"dtlsprofilename,omitempty"`
	Ersa                 string `json:"ersa,omitempty"`
	Ersacount            string `json:"ersacount,omitempty"`
	Hsts                 string `json:"hsts,omitempty"`
	Includesubdomains    string `json:"includesubdomains,omitempty"`
	Maxage               int    `json:"maxage,omitempty"`
	Ocspstapling         string `json:"ocspstapling,omitempty"`
	Pushenctrigger       string `json:"pushenctrigger,omitempty"`
	Redirectportrewrite  string `json:"redirectportrewrite,omitempty"`
	Sendclosenotify      string `json:"sendclosenotify,omitempty"`
	Sessreuse            string `json:"sessreuse,omitempty"`
	Sesstimeout          string `json:"sesstimeout,omitempty"`
	Snienable            string `json:"snienable,omitempty"`
	Ssl2                 string `json:"ssl2,omitempty"`
	Ssl3                 string `json:"ssl3,omitempty"`
	Sslprofile           string `json:"sslprofile,omitempty"`
	Sslredirect          string `json:"sslredirect,omitempty"`
	Sslv2Redirect        string `json:"sslv2redirect,omitempty"`
	Sslv2url             string `json:"sslv2url,omitempty"`
	Strictsigdigestcheck string `json:"strictsigdigestcheck,omitempty"`
	TLS1                 string `json:"tls1,omitempty"`
	TLS11                string `json:"tls11,omitempty"`
	TLS12                string `json:"tls12,omitempty"`
	Vservername          string `json:"vservername"`
}

// SslvserverWrapper wraps the object and serves as default response.
type SslvserverWrapper struct {
	Errorcode  int          `json:"errorcode,omitempty"`
	Message    string       `json:"message,omitempty"`
	Severity   string       `json:"severity,omitempty"`
	SSLVServer []Sslvserver `json:"sslvserver,omitempty"`
}

// Sslvserver describes the object.
type Sslvserver struct {
	Ca                   bool   `json:"ca,omitempty"`
	Cipherredirect       string `json:"cipherredirect,omitempty"`
	Cleartextport        int    `json:"cleartextport,omitempty"`
	Clientauth           string `json:"clientauth,omitempty"`
	Dh                   string `json:"dh,omitempty"`
	Dhcount              string `json:"dhcount,omitempty"`
	Dhkeyexpsizelimit    string `json:"dhkeyexpsizelimit,omitempty"`
	Dtlsflag             bool   `json:"dtlsflag,omitempty"`
	Ersa                 string `json:"ersa,omitempty"`
	Ersacount            string `json:"ersacount,omitempty"`
	Invoke               bool   `json:"invoke,omitempty"`
	Nonfipsciphers       string `json:"nonfipsciphers,omitempty"`
	Ocspstapling         string `json:"ocspstapling,omitempty"`
	Polinherit           string `json:"polinherit,omitempty"`
	Priority             string `json:"priority,omitempty"`
	Pushenctrigger       string `json:"pushenctrigger,omitempty"`
	Redirectportrewrite  string `json:"redirectportrewrite,omitempty"`
	Sendclosenotify      string `json:"sendclosenotify,omitempty"`
	Service              string `json:"service,omitempty"`
	Sessreuse            string `json:"sessreuse,omitempty"`
	Sesstimeout          string `json:"sesstimeout,omitempty"`
	Skipcaname           bool   `json:"skipcaname,omitempty"`
	Snicert              bool   `json:"snicert,omitempty"`
	Snienable            string `json:"snienable,omitempty"`
	Ssl2                 string `json:"ssl2,omitempty"`
	Ssl3                 string `json:"ssl3,omitempty"`
	Sslredirect          string `json:"sslredirect,omitempty"`
	Sslv2Redirect        string `json:"sslv2redirect,omitempty"`
	Strictsigdigestcheck string `json:"strictsigdigestcheck,omitempty"`
	TLS1                 string `json:"tls1,omitempty"`
	TLS11                string `json:"tls11,omitempty"`
	TLS12                string `json:"tls12,omitempty"`
	Vservername          string `json:"vservername"`
}
