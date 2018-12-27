package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-go-jwt/structs"
)

func (idb *InDB) GetBook(c *gin.Context) {
	var (
		user structs.Book
		result gin.H
		statuscode int
		messages string
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		statuscode = 404
		messages = "Book Not Found"
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		statuscode = 200
		messages = "Success Get Book"
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
}

// get all data in person
func (idb *InDB) GetBooks(c *gin.Context) {
	var (
		buku []structs.Book
		result  gin.H
		statuscode int
		messages string
	)

	idb.DB.Find(&buku)
	if len(buku) <= 0 {
		statuscode = 404
		messages = "Book Not Found"
		result = gin.H{
			"result": len(buku),
			"count":  len(buku),
		}
	} else {
		statuscode = 200
		messages = "Success Get All Book"
		result = gin.H{
			"result": buku,
			"count":  len(buku),
		}
	}

	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
}

// create new data to the database
func (idb *InDB) CreatBook(c *gin.Context) {
	var (
		buku structs.Book
		result gin.H
		statuscode int
		messages string
	)

	author := c.PostForm("author")
	name := c.PostForm("name")
	status := c.PostForm("status")
	buku.Author = author
	buku.Name = name
	buku.Status = status
	idb.DB.Create(&buku)
	statuscode = 200
	messages = "Success Create Book"
	result = gin.H{
		"result": buku,
	}
	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
}

// update data with {id} as query
func (idb *InDB) UpdateBook(c *gin.Context) {
	id := c.Query("id")
	author := c.PostForm("author")
	name := c.PostForm("name")
	status := c.PostForm("status")
	var (
		buku    structs.Book
		bukubaru structs.Book
		result    gin.H
		statuscode int
		messages string
	)

	err := idb.DB.First(&buku, id).Error
	println(err)
	if err != nil {
		statuscode = 404
		messages = "data not found"
		result = gin.H{
			"result": err,
		}
	}
	bukubaru.Author = author
	bukubaru.Name = name
	bukubaru.Status = status
	err = idb.DB.Model(&buku).Updates(bukubaru).Error
	if err != nil {
		statuscode = 400
		messages = "update failed"
		result = gin.H{
			"result": err,
		}
	} else {
		statuscode = 200
		messages = "successfully updated data"
		result = gin.H{
			"result": err,
		}
	}
	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
}

// delete data with {id}
func (idb *InDB) DeleteBook(c *gin.Context) {
	var (
		buku structs.Book
		result gin.H
		statuscode int
		messages string
	)
	id := c.Param("id")
	err := idb.DB.First(&buku, id).Error
	if err != nil {
		statuscode = 404
		messages = "data not found"
		result = gin.H{
			"status": 404,
			"result": "data not found",
		}
	}
	println(id)
	err = idb.DB.Where("id = ?",id).Delete(&buku).Error
	if err != nil {
		statuscode = 400
		messages = "delete failed"
		result = gin.H{
			"result": err,
		}
	} else {
		statuscode = 200
		messages = "Data deleted successfully"
		result = gin.H{
			"result": err,
		}
	}

	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
}
