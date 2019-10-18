package database

import (
	"fmt"

	"github.com/cone-partij/golang-api-boilerplate/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreConnection struct{}

func (c PostgreConnection) Connect() (db *gorm.DB, err error) {
	user := utils.Env("DB_USERNAME")
	password := utils.Env("DB_PASSWORD")
	address := utils.Env("DB_HOST")
	dbname := utils.Env("DB_DATABASE")

	raw := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", address, user, dbname, password)

	return gorm.Open("postgres", raw)
}
