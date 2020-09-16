package common

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
)

// User represents a user who has been logged in
type User struct {
	Origin   string
	Username string
	Name     string
	Email    string
	Roles    []string
	Token    *jwt.Token
}

// CreateUserFromToken convers a Jwt into a User struct
func CreateUserFromToken(token *jwt.Token) (*User, error) {
	if !token.Valid {
		return nil, errors.New("token is not valid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token claims can not be parsed")
	}

	var user *User
	err := mapstructure.Decode(claims, &user)
	if err != nil {
		return nil, err
	}
	user.Token = token

	return user, nil
}

// CreateUserFromTokenString parses a Jwt token string and returns a User struct
func CreateUserFromTokenString(tokenString string, verifyKey *rsa.PublicKey) (*User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	return CreateUserFromToken(token)
}

// GetJwt creates a sign Jwt
func (u *User) GetJwt(signingKey *rsa.PrivateKey, expiration time.Duration) (string, error) {
	if u == nil {
		return "", errors.New("user reference is nil")
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(expiration).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = u.Username

	claims["origin"] = u.Origin
	claims["username"] = u.Username
	claims["name"] = u.Name
	claims["email"] = u.Email
	claims["roles"] = u.Roles

	token.Claims = claims
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	u.Token = token
	return tokenString, nil
}

// RefreshJwt refreshes the token if needed
func (u *User) RefreshJwt(signingKey *rsa.PrivateKey, expirationDuration time.Duration) (string, error) {
	if u == nil {
		return "", errors.New("user reference is nil")
	}
	if u.Token == nil {
		return u.GetJwt(signingKey, expirationDuration)
	}
	token := u.Token
	err := token.Claims.Valid()
	if err != nil {
		return "", err
	}

	// calculate refresh window
	c := token.Claims.(jwt.MapClaims)
	var iat int64
	tiat, ok := c["iat"].(float64)
	if ok {
		iat = int64(tiat)
	} else {
		iat = c["iat"].(int64)
	}
	issued := time.Unix(iat, 0)
	window := issued.Add(expirationDuration - (expirationDuration / 4))
	glog.V(5).Infof("Token issued at %v, refresh window begins at %v.", issued, window)
	if time.Now().After(window) {
		return u.GetJwt(signingKey, expirationDuration)
	}

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// HasRole returns true if the User has the specified role
func (u User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}

	return false
}
