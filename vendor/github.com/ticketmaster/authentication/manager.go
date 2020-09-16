package authentication

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/ticketmaster/authentication/authorization"
	"github.com/ticketmaster/authentication/client"
	"github.com/ticketmaster/authentication/common"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

// Manager holds references to the authentication components
type Manager struct {
	AuthenticationClients []client.Client
	Authorization         *authorization.Authorization
	PrivateKey            *rsa.PrivateKey
	PublicKey             *rsa.PublicKey
	JwtExpiration         time.Duration
	EnableAnonymousAccess bool
}

// NewManager instantiates a new Manager struct from configuration
func NewManager() (*Manager, error) {
	manager := &Manager{}

	var configs = viper.Get("authenticationClient").([]interface{})
	for idx, config := range configs {
		provider := config.(map[interface{}]interface{})["provider"].(string)
		glog.V(2).Infof("Detected configuration for authentication client type: %s at index %v", provider, idx)
		added := false
		for scName, sc := range client.SupportedClients {
			if provider == scName {
				client, err := sc(config.(map[interface{}]interface{}))
				if err != nil {
					glog.Errorf("error creating client at index %v: %v", idx, err)
					continue
				}
				manager.AuthenticationClients = append(manager.AuthenticationClients, client)
				added = true
				break
			}
		}

		if !added {
			glog.Warningf("could not find type for authentication client. Skipping client at index %v", idx)
		}
	}

	authorizationConfig := viper.Get("authorization")
	if authorizationConfig == nil {
		glog.Infof("no authorization section present in config")
	} else {
		authorization, err := authorization.NewAuthorization(authorizationConfig.(map[string]interface{}))
		if err != nil {
			glog.Error(err)
		} else {
			manager.Authorization = authorization
		}
	}

	priv := viper.GetString("privateKey")
	if len(priv) > 0 {
		privBytes, err := ioutil.ReadFile(priv)
		if err != nil {
			return nil, err
		}
		privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
		if err != nil {
			return nil, err
		}

		manager.PrivateKey = privKey
	}

	pub := viper.GetString("publicKey")
	if len(pub) > 0 {
		pubBytes, err := ioutil.ReadFile(pub)
		if err != nil {
			return nil, err
		}
		pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
		if err != nil {
			return nil, err
		}

		manager.PublicKey = pubKey
	}

	expiration, err := time.ParseDuration(viper.GetString("jwtExpiration"))
	if err != nil {
		return nil, err
	}
	manager.JwtExpiration = expiration
	manager.EnableAnonymousAccess = viper.GetBool("enableAnonymousAccess")

	return manager, nil
}

// CreateAnonymousUser creates a User struct for anonymous users
func (m Manager) CreateAnonymousUser() (*common.User, error) {
	if !m.EnableAnonymousAccess {
		return nil, errors.New("anonymous access it not permitted")
	}
	user := common.User{Origin: "Anonymous", Username: "Anonymous", Name: "Anonymous User", Email: ""}
	user.Roles = append(user.Roles, "Anonymous")
	return &user, nil
}

// ValidateCredentials takes a set of credentials and returns a User struct if the credentials are valid
func (m Manager) ValidateCredentials(username string, password string) (*common.User, error) {
	var errStrings []string
	var u *common.User
	var err error
	start := time.Now()
	if len(m.AuthenticationClients) == 0 {
		return nil, errors.New("no authentication providers enabled")
	}

	for _, c := range m.AuthenticationClients {
		u, err = c.ValidateCredentials(username, password)
		if err != nil {
			errStrings = append(errStrings, fmt.Sprintf("%v: %v", c.GetOrigin(), err))
		}
	}

	e2 := time.Since(start)
	glog.V(2).Infof("Validated credentials in %s", e2)

	if u == nil {
		return nil, fmt.Errorf(strings.Join(errStrings, "\n"))
	}

	return u, nil
}

// CreateUserFromToken convers a Jwt into a User struct
func (m Manager) CreateUserFromToken(token *jwt.Token) (*common.User, error) {
	return common.CreateUserFromToken(token)
}

// CreateUserFromTokenString parses a Jwt token string and returns a User struct
func (m Manager) CreateUserFromTokenString(tokenString string) (*common.User, error) {
	return common.CreateUserFromTokenString(tokenString, m.PublicKey)
}

// GetJwt gets a JWT for a given user
func (m Manager) GetJwt(u *common.User) (string, error) {
	return u.GetJwt(m.PrivateKey, m.JwtExpiration)
}

// RefreshJwt refreshes a JWT for a given user. If the expiration window is not yet available, the existing token is returned.
func (m Manager) RefreshJwt(u *common.User) (string, error) {
	return u.RefreshJwt(m.PrivateKey, m.JwtExpiration)
}

// IsAuthorized determines if a user is authorized for the specified action
func (m Manager) IsAuthorized(u *common.User, actions map[string]string) bool {

	return m.Authorization.IsAuthorized(u, actions)
}
