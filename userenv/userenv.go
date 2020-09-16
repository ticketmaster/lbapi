package userenv

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/ticketmaster/authentication/common"
	"github.com/ticketmaster/lbapi/shared"
)

// New - package constructor
func New(c *gin.Context) *User {
	return &User{
		Group:    fetchRoles(c),
		Username: fetchUser(c),
		Context:  c,
	}
}

// FetchRoles parses a session token and returns all roles.
// associated to a logged in user.
func fetchRoles(c *gin.Context) []string {
	////////////////////////////////////////////////////////////////////////////
	// Fetch authorization (Post-Token Auth).
	////////////////////////////////////////////////////////////////////////////
	authString := c.Request.Header.Get("Authorization")
	if authString != "" && strings.Contains(authString, "Bearer") {
		var user *common.User
		token := strings.Replace(authString, "Bearer ", "", 1)
		parsed, _ := jwt.Parse(token, nil)
		claims, _ := parsed.Claims.(jwt.MapClaims)
		mapstructure.Decode(claims, &user)
		return user.Roles
	}
	////////////////////////////////////////////////////////////////////////////
	// Fetch authorization from session.
	////////////////////////////////////////////////////////////////////////////
	s := sessions.Default(c)
	if s == nil {
		return nil
	}

	u := s.Get("roles")
	if u == nil {
		return nil
	}

	r, _ := b64.StdEncoding.DecodeString(u.(string))
	return strings.Split(string(r), ",")
}

// FetchUser parses a session token and returns the user object.
func fetchUser(c *gin.Context) (r string) {
	////////////////////////////////////////////////////////////////////////////
	// Fetch request body and restore it to buffer (Pre-Token Auth).
	////////////////////////////////////////////////////////////////////////////
	if c.Request.Body != nil && c.Request.URL.String() == "/login" {
		var login Login
		var body []byte
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Error(err)
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, &login)
		if err != nil {
			c.Error(err)
			return
		}
		return login.Username
	}
	////////////////////////////////////////////////////////////////////////////
	// Fetch authorization (Post-Token Auth).
	////////////////////////////////////////////////////////////////////////////
	authString := c.Request.Header.Get("Authorization")
	if authString != "" && strings.Contains(authString, "Bearer") {
		var user *common.User
		token := strings.Replace(authString, "Bearer ", "", 1)
		parsed, _ := jwt.Parse(token, nil)
		claims, _ := parsed.Claims.(jwt.MapClaims)
		mapstructure.Decode(claims, &user)
		return user.Username
	}
	////////////////////////////////////////////////////////////////////////////
	// Fetch authorization from session.
	////////////////////////////////////////////////////////////////////////////
	s := sessions.Default(c)
	if s == nil {
		return
	}
	u := s.Get("username")
	err := shared.MarshalInterface(u, &r)
	if err != nil {
		c.Error(err)
		return
	}
	return r
}

// HasAdminRight matches a product code to roles associated with a user.
func (o *User) HasAdminRight(code string) (err error) {
	if code == "" {
		return errors.New("you did not provide a product code for this asset")
	}
	////////////////////////////////////////////////////////////////////////////
	operator, err := o.hasRole("prd"+code+"-operator", o.Group)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	limited, err := o.hasRole("prd"+code+"-limitedaccess", o.Group)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if operator == false && limited == false {
		return errors.New("you are not authorized to make modifications to this asset. product code: " + code)
	}
	return nil
}

// hasRole matches a role (defined by regexp) to a security group.
// in LDAP.
func (o *User) hasRole(role string, roles []string) (matched bool, err error) {
	for _, r := range roles {
		matched, err = regexp.MatchString(role, strings.ToLower(r))
		if matched == true {
			return matched, err
		}
		adminRole := "prd1234-operator"
		matched, err = regexp.MatchString(adminRole, strings.ToLower(r))
		if matched == true {
			return matched, err
		}
	}
	return matched, err
}
