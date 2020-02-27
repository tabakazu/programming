package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/simple-auth-server-demo/usecase"
)

func main() {
	r := gin.Default()
	r.POST("/register", RegisterHandler)
	r.Run(":8080")
}

func RegisterHandler(c *gin.Context) {
	var json struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userUsecase := usecase.NewUserUsecase()
	user, err := userUsecase.Register(json.Email, json.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}
