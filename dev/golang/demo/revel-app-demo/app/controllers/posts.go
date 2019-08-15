package controllers

import (
	"github.com/revel/revel"
)

type Posts struct {
	App
}

func (c Posts) Index() revel.Result {
	return c.Render()
}
