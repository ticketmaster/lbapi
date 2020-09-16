package client

import "github.com/ticketmaster/authentication/common"

// Client represents a connection to an LDAP server
type Client interface {
	GetOrigin() string
	ValidateCredentials(username string, password string) (*common.User, error)
}

// ClientConstructor is a function definition for client constructors
type ClientConstructor func(map[interface{}]interface{}) (Client, error)

var SupportedClients map[string]ClientConstructor

func init() {
	SupportedClients = make(map[string]ClientConstructor)
	RegisterSupportedClient("ldap", NewLdapClient)
	RegisterSupportedClient("memory", NewMemoryClient)
}

// RegisterSupportedClient registers a client for use
func RegisterSupportedClient(providerName string, constructor ClientConstructor) {
	SupportedClients[providerName] = constructor
}
