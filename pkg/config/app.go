package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dbCon, err := gorm.Open("mysql", "root:@/simplerest?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db = dbCon
}

func GetDB() *gorm.DB {
	return db
}
