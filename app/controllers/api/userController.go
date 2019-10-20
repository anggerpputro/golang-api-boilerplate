package api

import (
	"fmt"
	"log"

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

	var users []ctlModel.User
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

	var user ctlModel.User
	// Get first matched record
	db.First(&user, id)

	if user.ID == 0 {
		errMessage := fmt.Sprintf("User with ID: %v not found!", id)
		core.NewJsonResponse().ResponseNotFound(gc, gin.H{
			"error": map[string]string{
				"message": errMessage,
			},
		})
		log.Panicln(errMessage)
	}

	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"user": user,
	})
}

func (c UserController) Store(gc *gin.Context) {
	db := ctlDB.GetDefaultConnection()
	defer db.Close()

	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		core.NewJsonResponse().ResponseError(gc, gin.H{
			"error": err,
		})
		log.Panicln(err)
	}

	user := ctlModel.NewUser()
	user.Name = gc.PostForm("name")
	user.Username = gc.PostForm("username")
	user.Email = gc.PostForm("email")
	user.Password = gc.PostForm("password")

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		core.NewJsonResponse().ResponseError(gc, gin.H{
			"error": err,
		})
		log.Panicln(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		core.NewJsonResponse().ResponseError(gc, gin.H{
			"error": err,
		})
		log.Panicln(err)
	}

	core.NewJsonResponse().ResponseSuccess(gc, gin.H{
		"store": user,
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
