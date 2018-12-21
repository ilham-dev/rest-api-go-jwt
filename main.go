package main

import (
	"github.com/gin-gonic/gin"
	"rest-api-go-jwt/config"
	"rest-api-go-jwt/controller"
)

var (
	mysupersecretpassword = "mysupersecretpasswords"
)

func main() {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}
	route := gin.Default();
	/**
		Create Group Route
	 */
	private := route.Group("/api")
	private.Use(config.Auth(mysupersecretpassword))
	/*
		Global Route
	 */
	route.GET("/", controller.Index)
	route.POST("/login", inDB.Login)
	route.POST("/register", inDB.CreateUser)
	/*
		Private Route User
	 */
	user := private.Group("/user")

	user.GET("/user/:id", inDB.GetUser)
	user.GET("/users", inDB.GetUsers)
	user.PUT("/user", inDB.UpdateUser)
	user.DELETE("/user/:id", inDB.DeleteUser)

	/*
		Private Route Book
	 */
	book := private.Group("/book")
	book.GET("/book/:id", inDB.GetBook)
	book.POST("/book", inDB.CreatBook)
	book.GET("/books", inDB.GetBooks)
	book.PUT("/book", inDB.UpdateBook)
	book.DELETE("/book/:id", inDB.DeleteBook)

	_ = route.Run();
}
