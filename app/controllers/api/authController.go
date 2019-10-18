package api

import (
	"github.com/cone-partij/golang-api-boilerplate/core"
	"github.com/gin-gonic/gin"
)

type AuthController core.Controller

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c AuthController) Login(gc *gin.Context) {
	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"message": "login",
	})
}

func (c AuthController) Register(gc *gin.Context) {
	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"message": "register",
	})
}

func (c AuthController) Logout(gc *gin.Context) {
	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"message": "logout",
	})
}

func (c AuthController) Me(gc *gin.Context) {
	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"message": "me",
	})
}

func (c AuthController) Refresh(gc *gin.Context) {
	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"message": "refresh",
	})
}
