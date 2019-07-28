package controllers

import (
	"../../core"
	"github.com/gin-gonic/gin"
)

type HelloController core.Controller

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c HelloController) Ping(gc *gin.Context) {
	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"message": "pong",
	})
}

func (c HelloController) Hello(gc *gin.Context) {
	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"message": "Hello World!",
	})
}
