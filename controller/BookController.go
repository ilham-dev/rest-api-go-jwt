package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-go-jwt/structs"
)

func (idb *InDB) GetBook(c *gin.Context) {
	var (
		user structs.Book
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data in person
func (idb *InDB) GetBooks(c *gin.Context) {
	var (
		buku []structs.Book
		result  gin.H
	)

	idb.DB.Find(&buku)
	if len(buku) <= 0 {
		result = gin.H{
			"status": 200,
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": buku,
			"count":  len(buku),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data to the database
func (idb *InDB) CreatBook(c *gin.Context) {
	var (
		buku structs.Book
		result gin.H
	)

	author := c.PostForm("author")
	name := c.PostForm("name")
	status := c.PostForm("status")
	buku.Author = author
	buku.Name = name
	buku.Status = status
	idb.DB.Create(&buku)
	result = gin.H{
		"status": 200,
		"result": buku,
	}
	c.JSON(http.StatusOK, result)
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
	)

	err := idb.DB.First(&buku, id).Error
	println(err)
	if err != nil {
		result = gin.H{
			"status": 404,
			"result": "data not found",
		}
	}
	bukubaru.Author = author
	bukubaru.Name = name
	bukubaru.Status = status
	err = idb.DB.Model(&buku).Updates(bukubaru).Error
	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeleteBook(c *gin.Context) {
	var (
		buku structs.Book
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&buku, id).Error
	if err != nil {
		result = gin.H{
			"status": 404,
			"result": "data not found",
		}
	}
	println(id)
	err = idb.DB.Where("id = ?",id).Delete(&buku).Error
	fmt.Println(err)
	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
