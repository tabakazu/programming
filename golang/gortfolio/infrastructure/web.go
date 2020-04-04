package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/gortfolio/interfaces/controller"
)

var Web *gin.Engine

func init() {
	router := gin.Default()
	usersController := controller.NewUsersController()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/users", func(c *gin.Context) { usersController.Create(c) })
		}
	}

	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	Web = router
}
