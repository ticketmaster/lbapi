package common

import (
	"crypto/rsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

func getKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privBytes, err := ioutil.ReadFile("../test-certificates/jwt.rsa")
	if err != nil {
		return nil, nil, err
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return nil, nil, err
	}

	pubBytes, err := ioutil.ReadFile("../test-certificates/jwt.rsa.pub")
	if err != nil {
		return nil, nil, err
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return nil, nil, err
	}

	return privKey, pubKey, nil
}

// CopyMap extends common
func CopyMap(source map[string]interface{}) map[string]interface{} {
	dest := make(map[string]interface{})
	for k, v := range source {
		dest[k] = v
	}

	return dest
}
