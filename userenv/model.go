package userenv

import "github.com/gin-gonic/gin"

// User Object.
type User struct {
	Group    []string
	Username string
	Error    error
	Context  *gin.Context
}

type Login struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}
