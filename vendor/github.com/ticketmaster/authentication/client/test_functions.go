package client

import (
	"bytes"

	"github.com/spf13/viper"
)

func GetConfigElement(config []byte) map[interface{}]interface{} {
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))
	client := viper.Get("authenticationClient")
	return client.([]interface{})[0].(map[interface{}]interface{})
}
