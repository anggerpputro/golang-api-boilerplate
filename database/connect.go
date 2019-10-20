package database

import (
	"log"

	connection "github.com/cone-partij/golang-api-boilerplate/database/connections"
	ctlUtil "github.com/cone-partij/golang-api-boilerplate/utils"
	"github.com/jinzhu/gorm"
)

type DbConnection interface {
	Connect() (db *gorm.DB, err error)
}

func GetConnection(connectionName string) (db *gorm.DB, err error) {
	switch connectionName {
	case "mysql":
		conn := &connection.MysqlConnection{}
		return conn.Connect()
	case "postgre":
		conn := &connection.PostgreConnection{}
		return conn.Connect()
	default:
		return nil, nil
	}
}

func GetDefaultConnection() (db *gorm.DB) {
	db, err := GetConnection(ctlUtil.Env("DB_CONNECTION"))
	// defer db.Close()

	if err != nil {
		log.Panicln(err)
	}

	return db
}
