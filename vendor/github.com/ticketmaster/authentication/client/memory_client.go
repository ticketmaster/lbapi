package client

import (
	"errors"

	"github.com/ticketmaster/authentication/common"
	"github.com/mitchellh/mapstructure"
)

// MemoryClient represents an in-memory user authentication store. This is highly insecure (plaintext passwords) and should be used for TESTING ONLY.
type MemoryClient struct {
	Origin string
	Users  []*configUser
}

type configUser struct {
	Name     string
	Username string
	Password string
	Email    string
	Roles    []string
}

func init() {

}

// NewMemoryClient creates a new client from the specified configuration
func NewMemoryClient(config map[interface{}]interface{}) (Client, error) {
	origin, ok := config["origin"].(string)
	if !ok {
		origin = "memory"
	}

	var users []*configUser
	u, ok := config["users"]
	if ok {
		err := mapstructure.Decode(u, &users)
		if err != nil {
			return nil, err
		}
	}

	return &MemoryClient{origin, users}, nil
}

// ValidateCredentials takes a set of credentials and returns a User struct if the credentials are valid
func (c MemoryClient) ValidateCredentials(username string, password string) (*common.User, error) {
	user := c.getUser(username)
	if user == nil || user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	newUser := common.User{Origin: c.GetOrigin(), Username: username, Name: user.Name, Email: user.Email, Roles: user.Roles}

	return &newUser, nil
}

func (c MemoryClient) getUser(username string) *configUser {
	for _, user := range c.Users {
		if user.Username == username {
			return user
		}
	}

	return nil
}

// GetOrigin returns the origin for this client
func (c MemoryClient) GetOrigin() string {
	return c.Origin
}
