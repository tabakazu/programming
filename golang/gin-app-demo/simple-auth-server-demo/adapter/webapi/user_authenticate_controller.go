package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/simple-auth-server-demo/infra/serializer"
	"github.com/tabakazu/simple-auth-server-demo/interactor"
	"github.com/tabakazu/simple-auth-server-demo/interactor/usecase"
)

type UserAuthenticateController struct {
	UserAuthenticate usecase.UserAuthenticateUseCase
}

func NewUserAuthenticateController() *UserAuthenticateController {
	return &UserAuthenticateController{
		UserAuthenticate: interactor.NewUserAuthenticateInteractor(),
	}
}

func (ctrl *UserAuthenticateController) Create(c *gin.Context) {
	var json struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.UserAuthenticate.Handle(json.Email, json.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	serializedUser := serializer.SerializeUser(user)
	c.JSON(http.StatusOK, serializedUser)
	return
}
