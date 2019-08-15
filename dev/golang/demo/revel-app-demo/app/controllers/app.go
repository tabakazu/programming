package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Ping() revel.Result {
	type response struct {
		Message string `json:"message"`
	}
	result := response{Message: "pong"}
	return c.RenderJSON(result)
}
