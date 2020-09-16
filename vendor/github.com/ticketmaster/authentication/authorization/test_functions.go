package authorization

import (
	"bytes"

	"github.com/spf13/viper"
)

func getConfigElement(config []byte) map[string]interface{} {
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))
	client := viper.Get("authorization")
	return client.(map[string]interface{})
}
