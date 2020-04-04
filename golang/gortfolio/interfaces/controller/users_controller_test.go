package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tabakazu/gortfolio/interfaces/controller"
)

func TestUsersController_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	usersController := controller.NewUsersController()
	r.POST("/users", func(c *gin.Context) { usersController.Create(c) })

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
