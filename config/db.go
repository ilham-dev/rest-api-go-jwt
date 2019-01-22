package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
//tes ilham
