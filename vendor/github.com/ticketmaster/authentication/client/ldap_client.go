package client

import (
	"crypto/tls"
	"errors"
	"fmt"
	"regexp"

	"github.com/ticketmaster/authentication/common"
	ldap "gopkg.in/ldap.v3"
)

// LdapClient represents a connection to an LDAP server
type LdapClient struct {
	Endpoint           string
	Port               int
	UseTLS             bool
	BaseDN             string
	ShortDomain        string
	InsecureSkipVerify bool
	TLSServerName      string
}

// NewLdapClient creates a new client from the specified configuration map
func NewLdapClient(config map[interface{}]interface{}) (Client, error) {
	endpoint, ok := config["endpoint"].(string)
	if !ok {
		return nil, errors.New("endpoint must be specified in configuration")
	}

	port, ok := config["port"].(int)
	if !ok {
		port = 389
	}

	useTLS, ok := config["useTLS"].(bool)
	if !ok {
		useTLS = true
	}

	baseDN, ok := config["baseDN"].(string)
	if !ok {
		return nil, errors.New("baseDN must be specified in configuration")
	}

	shortDomain, ok := config["shortDomain"].(string)
	if !ok {
		return nil, errors.New("shortDomain must be specified in configuration")
	}

	insecureSkipVerify, ok := config["insecureSkipVerify"].(bool)
	if !ok {
		insecureSkipVerify = false
	}

	tlsServerName, ok := config["tlsServerName"].(string)
	if !ok {
		tlsServerName = endpoint
	}

	return &LdapClient{
		endpoint,
		port,
		useTLS,
		baseDN,
		shortDomain,
		insecureSkipVerify,
		tlsServerName}, nil
}

// ValidateCredentials takes a set of credentials and returns a User struct if the credentials are valid
func (c LdapClient) ValidateCredentials(username string, password string) (*common.User, error) {
	var l *ldap.Conn
	var err error
	if c.UseTLS {
		l, err = ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", c.Endpoint, c.Port), &tls.Config{InsecureSkipVerify: c.InsecureSkipVerify, ServerName: c.TLSServerName})
		if err != nil {
			return nil, err
		}
	} else {
		l, err = ldap.Dial("tcp", fmt.Sprintf("%s:%d", c.Endpoint, c.Port))
		if err != nil {
			return nil, err
		}
	}

	defer l.Close()

	samAccountName, domain, err := c.parseUsername(username)
	if err != nil {
		return nil, err
	}

	err = l.Bind(fmt.Sprintf("%s\\%s", domain, samAccountName), password)
	if err != nil {
		return nil, err
	}

	request := ldap.NewSearchRequest(c.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 2, 0, false, fmt.Sprintf("(&(objectClass=user)(samAccountName=%s))", samAccountName), []string{"dn", "cn", "displayName", "mail"}, nil)
	sr, err := l.Search(request)
	if err != nil {
		return nil, err
	}

	if len(sr.Entries) == 0 {
		return nil, errors.New("could not find the logged in user")
	}

	if len(sr.Entries) > 1 {
		return nil, errors.New("found multiple users when searching for logged in user")
	}

	dn := sr.Entries[0].DN
	groupRequest := ldap.NewSearchRequest(c.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, fmt.Sprintf("(&(objectclass=group)(member:1.2.840.113556.1.4.1941:=%s))", dn), []string{"dn", "cn", "name"}, nil)
	gr, err := l.Search(groupRequest)
	if err != nil {
		return nil, err
	}

	user := common.User{Origin: c.GetOrigin(), Username: samAccountName, Name: sr.Entries[0].GetAttributeValue("displayName"), Email: sr.Entries[0].GetAttributeValue("mail")}
	for _, entry := range gr.Entries {
		user.Roles = append(user.Roles, entry.GetAttributeValue("name"))
	}

	return &user, nil
}

func (c LdapClient) parseUsername(username string) (string, string, error) {
	re, err := regexp.Compile(`^((.*?)\\)?(.*?)(@(.*?))?$`)
	if err != nil {
		return "", "", err
	}
	result := re.FindStringSubmatch(username)
	if result == nil {
		return "", "", errors.New("could not parse username")
	}
	samAccountName := result[3]
	domain := c.ShortDomain
	if len(result[2]) > 0 {
		domain = result[2]
	}

	return samAccountName, domain, nil
}

// GetOrigin returns the origin for this client
func (c LdapClient) GetOrigin() string {
	return c.ShortDomain
}
