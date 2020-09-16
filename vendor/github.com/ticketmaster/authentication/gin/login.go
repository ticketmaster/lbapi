package gin

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type loginCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login provides an endpoint to log in and receive a JWT token
func Login(c *gin.Context) {
	request := struct {
		Username string
		Password string
	}{"", ""}
	err := c.BindJSON(&request)
	if err != nil {
		glog.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	if !currentOptions.EnableJwtAuthentication {
		err := errors.New("JWT authentication is disabled")
		glog.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	user, err := currentOptions.manager.ValidateCredentials(request.Username, request.Password)
	if err != nil {
		glog.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	token, err := currentOptions.manager.GetJwt(user)
	if err != nil {
		glog.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("jwt", token)
	session.Save()
	c.JSON(http.StatusOK, struct{ Token string }{token})
	return

}
