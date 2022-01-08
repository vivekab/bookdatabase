package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PWD")
	d, err := gorm.Open("mysql", user+":"+pwd+"/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
