package config

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"rest-api-go-jwt/structs"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/go-api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		println(err)
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.User{})
	db.AutoMigrate(structs.Book{})
	return db
}

//tes ji
