package webapi

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine {
	r := gin.Default()

	userRegistration := NewUserRegistrationController()
	r.POST("/auth/registration", userRegistration.Create)

	userAuthenticate := NewUserAuthenticateController()
	r.POST("/auth", userAuthenticate.Create)

	user := NewUserController()
	r.GET("/user", user.Show)

	return r
}
