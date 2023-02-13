package orm

import (
	"api-golang/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func CreateConnection() {
	url := config.GetUrlDatabase()
	if connection, err := gorm.Open(mysql.Open(url), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		db = connection
	}
}

func CreateTables() {
	db.AutoMigrate(&User{})
}
