package api

import (
	"github.com/cone-partij/golang-api-boilerplate/core"
	"github.com/gin-gonic/gin"
)

type UserController core.Controller

func NewUserController() *UserController {
	return &UserController{}
}

func (c UserController) Index(gc *gin.Context) {
	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"message": "login",
	})
}
