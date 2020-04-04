package controller

import "net/http"

type UsersController struct{}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (ctrl *UsersController) Create(c Context) {
	c.JSON(http.StatusCreated, nil)
}
