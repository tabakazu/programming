package main

import (
	"fmt"
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
	c.Bind(&json)

	userUsecase := usecase.NewUserUsecase()
	user := userUsecase.Register(json.Email, json.Password)
	fmt.Println(user)
}
