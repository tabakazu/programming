package controller

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
	Param(string) string
	ShouldBindJSON(interface{}) error
	Status(int)
}
