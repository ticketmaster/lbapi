package gin

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/ticketmaster/authentication/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// Authentication provides the gin handler function to authenticate requests
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.ToLower(c.Request.RequestURI) == "/login" {
			// Skip auth for login page
			c.Next()
			return
		}

		session := sessions.Default(c)
		var user *common.User
		var err error
		// Validate JWT stores in session
		jwtToken := session.Get("jwt")
		glog.Infoln("Cookie:", jwtToken)
		tokenString := ""
		if jwtToken != nil {
			tokenString = jwtToken.(string)
		}

		if currentOptions.EnableJwtAuthentication && len(tokenString) > 0 {
			user, err = validateJwt(c, tokenString)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"message": err.Error()})
				return
			}
		} else if auth := c.Request.Header.Get("Authorization"); auth != "" {
			authHeader := strings.Split(auth, " ")
			switch authHeader[0] {
			case "Bearer":
				if !currentOptions.EnableJwtAuthentication {
					err := errors.New("JWT provided, but is not enabled for authentication")
					c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"message": err.Error()})
					return
				}

				user, err = validateJwt(c, authHeader[1])
				if err != nil {
					c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"message": err.Error()})
					return
				}
			case "Basic":
				username, password, err := getCredentials(authHeader[1])
				if err != nil {
					c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"message": err.Error()})
					return
				}

				user, err = currentOptions.manager.ValidateCredentials(username, password)
				if err != nil {
					if strings.Index(strings.ToLower(err.Error()), "invalid credentials") == -1 {
						c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"message": err.Error()})
						return
					}

					unauthorized(c, currentOptions)
					return
				}
			default:
				err := errors.New("unrecognized authentication header")
				c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"message": err.Error()})
				return
			}
		} else {
			if currentOptions.manager.EnableAnonymousAccess {
				user, err = currentOptions.manager.CreateAnonymousUser()
				if err != nil {
					c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"message": err.Error()})
					return
				}
			} else {
				if currentOptions.manager.Authorization.Default != "allow" {
					unauthorized(c, currentOptions)
					return
				}
			}

		}

		if user != nil {
			authorized := currentOptions.manager.IsAuthorized(user, map[string]string{"action": "", "route": c.Request.RequestURI, "method": c.Request.Method})
			if !authorized {
				c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"message": "Not authorized"})
				return
			}
			setUserData(c, user)
		} else {
			err := errors.New("unrecognized user")
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"message": err.Error()})
			return
		}
		c.Next()
	}
}

func validateJwt(c *gin.Context, tokenString string) (*common.User, error) {
	if currentOptions.EnableJwtAuthentication && len(tokenString) > 0 {
		user, err := currentOptions.manager.CreateUserFromTokenString(tokenString)
		if err != nil {
			return nil, err
		}

		// Refresh token as needed
		newTokenString, err := currentOptions.manager.RefreshJwt(user)
		if err != nil {
			return nil, err
		}
		if tokenString != newTokenString {
			session := sessions.Default(c)
			session.Set("jwt", newTokenString)
		}

		return user, nil
	}

	return nil, errors.New("no token string provided")
}

func setUserData(c *gin.Context, user *common.User) {
	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Set("name", user.Name)
	session.Set("email", user.Email)
	session.Set("roles", base64.StdEncoding.EncodeToString([]byte(strings.Join(user.Roles, ","))))
}

func unauthorized(c *gin.Context, currentOptions *AuthenticationOptions) {
	if currentOptions.EnableBasicAuthentication {
		c.Writer.Header().Add("WWW-Authenticate", `Basic realm="gin"`)
	}

	if currentOptions.EnableJwtAuthentication {
		c.Writer.Header().Add("WWW-Authenticate", `Bearer`)
	}

	glog.Infof(c.Writer.Header().Get("WWW-Authenticate"))
	c.AbortWithStatusJSON(401, map[string]string{"message": "Not authorized"})
}

func getCredentials(data string) (username, password string, err error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", "", err
	}
	strData := strings.Split(string(decodedData), ":")
	username = strData[0]
	password = strData[1]
	return
}
