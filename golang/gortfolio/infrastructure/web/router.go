package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/gortfolio/infrastructure/datastore"
	"github.com/tabakazu/gortfolio/interfaces/controller"
	"github.com/tabakazu/gortfolio/usecase/interactor"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	usersRepository := datastore.NewUserRepository()
	userRegisterUsecase := interactor.NewUserRegisterInteractor(usersRepository)
	userLoginUsecase := interactor.NewUserLoginInteractor(usersRepository)
	usersController := controller.NewUsersController(userRegisterUsecase)
	userSessionsController := controller.NewUserSessionsController(userLoginUsecase)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/users", func(c *gin.Context) { usersController.Create(c) })
			v1.POST("/login", func(c *gin.Context) { userSessionsController.Create(c) })
		}
	}

	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	Router = router
}
