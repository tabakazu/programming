package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/simple-auth-server-demo/infra/serializer"
	"github.com/tabakazu/simple-auth-server-demo/interactor"
	"github.com/tabakazu/simple-auth-server-demo/interactor/usecase"
)

type UserController struct {
	GetByToken usecase.UserGetByTokenUseCase
}

func NewUserController() *UserController {
	return &UserController{
		GetByToken: interactor.NewUserGetByTokenInteractor(),
	}
}

func (ctrl *UserController) Show(c *gin.Context) {
	var headers struct {
		APIToken string `header:"HTTP_AUTHORIZATION" binding:"required"`
	}

	if err := c.ShouldBindHeader(&headers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.GetByToken.Handle(headers.APIToken)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	serializedUser := serializer.SerializeUser(user)
	c.JSON(http.StatusOK, serializedUser)
	return
}
