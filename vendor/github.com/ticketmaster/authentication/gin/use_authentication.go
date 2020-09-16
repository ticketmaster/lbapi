package gin

import (
	"github.com/ticketmaster/authentication"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var currentOptions *AuthenticationOptions

// AuthenticationOptions is a struct that holds options relating to the authentication package
type AuthenticationOptions struct {
	EnableJwtAuthentication   bool
	EnableBasicAuthentication bool
	ConfigName                string
	ConfigPath                string
	EnvironmentVarPrefix      string
	manager                   *authentication.Manager
}

// NewAuthenticationOptions creates a set of default authentication options
func NewAuthenticationOptions() *AuthenticationOptions {
	return &AuthenticationOptions{EnableBasicAuthentication: true, EnableJwtAuthentication: true, ConfigName: "authentication", ConfigPath: "./", EnvironmentVarPrefix: "AUTH"}
}

// UseAuthentication is a helper function to set up authentication handlers and login API endpoints
func UseAuthentication(r *gin.Engine, options *AuthenticationOptions) error {
	viper.SetConfigName(options.ConfigName)
	viper.AddConfigPath(options.ConfigPath)
	viper.SetEnvPrefix("AUTH")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	manager, err := authentication.NewManager()
	if err != nil {
		return err
	}

	options.manager = manager
	currentOptions = options

	store := cookie.NewStore([]byte("secretkey"))
	r.Use(sessions.Sessions("auth-session", store))
	r.Use(Authentication())

	r.POST("/login", Login)
	return nil
}
