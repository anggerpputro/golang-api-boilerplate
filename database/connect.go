package database

import (
	connection "github.com/cone-partij/golang-api-boilerplate/database/connections"
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
