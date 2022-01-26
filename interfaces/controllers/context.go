package controllers

type Context interface {
	Param(string) string
	JSON(int, interface{}) error
}
