package database

import (
	"fmt"

	"github.com/cone-partij/golang-api-boilerplate/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlConnection struct{}

func (c MysqlConnection) Connect() (db *gorm.DB, err error) {
	user := utils.Env("DB_USERNAME")
	password := utils.Env("DB_PASSWORD")
	address := utils.Env("DB_HOST")
	dbname := utils.Env("DB_DATABASE")

	raw := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, address, dbname)

	return gorm.Open("mysql", raw)
}
