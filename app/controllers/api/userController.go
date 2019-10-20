package api

import (
	ctlModel "github.com/cone-partij/golang-api-boilerplate/app/models"
	"github.com/cone-partij/golang-api-boilerplate/core"
	ctlDB "github.com/cone-partij/golang-api-boilerplate/database"
	"github.com/gin-gonic/gin"
)

type UserController core.Controller

func NewUserController() *UserController {
	return &UserController{}
}

func (c UserController) Index(gc *gin.Context) {
	db := ctlDB.GetDefaultConnection()
	defer db.Close()

	users := ctlModel.NewUser()

	// Get all records
	db.Find(&users)

	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"users": users,
	})
}

func (c UserController) Show(gc *gin.Context) {
	id := gc.Param("id")

	db := ctlDB.GetDefaultConnection()
	defer db.Close()

	user := ctlModel.NewUser()

	// Get first matched record
	db.Where("id = ?", id).First(&user)

	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"user": user,
	})
}

func (c UserController) Store(gc *gin.Context) {
	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"store": true,
	})
}

func (c UserController) Update(gc *gin.Context) {
	id := gc.Param("id")

	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"update": id,
	})
}

func (c UserController) Destroy(gc *gin.Context) {
	id := gc.Param("id")

	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"destroy": id,
	})
}
