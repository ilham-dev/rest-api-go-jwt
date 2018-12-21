package controller

import (
	"fmt"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"rest-api-go-jwt/structs"
	"net/http"
	"time"
)
var (
	mysupersecretpassword = "mysupersecretpasswords"
)
// to get one data with {id}
func (idb *InDB) GetUser(c *gin.Context) {
	var (
		user structs.User
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
func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		persons []structs.User
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"status": 200,
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data to the database
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		person structs.User
		result gin.H
	)

	user_name := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	person.Username = user_name
	person.Email = email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	person.Password = string(hashedPassword)
	idb.DB.Create(&person)
	result = gin.H{
		"status": 200,
		"result": person,
	}
	c.JSON(http.StatusOK, result)
}

// update data with {id} as query
func (idb *InDB) UpdateUser(c *gin.Context) {
	id := c.Query("id")
	username := c.PostForm("username")
	email := c.PostForm("email")
	var (
		person    structs.User
		newPerson structs.User
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	println(err)
	if err != nil {
		result = gin.H{
			"status": 404,
			"result": "data not found",
		}
	}
	newPerson.Username = username
	newPerson.Email = email
	err = idb.DB.Model(&person).Updates(newPerson).Error
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
func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		person structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"status": 404,
			"result": "data not found",
		}
	}
	println(id)
	err = idb.DB.Where("id = ?",id).Delete(&person).Error
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

func (idb *InDB) Login(c *gin.Context){
	var (
		user structs.User
		result gin.H
	)
	username := c.PostForm("username")
	password := c.PostForm("password")

	_ = idb.DB.Where("username = ?", username).First(&user).Error

	//deskripsi dan compare password
	var password_tes = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if password_tes == nil {
		// Login success
		// Create the token
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
		// Set some claims
		token.Claims = jwt_lib.MapClaims{
			"Id":  "Christopher",
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		}
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(mysupersecretpassword))
		if err != nil {
			result = gin.H{
				"status": 500,
				"result": "failed generate password",
			}
		}else{
			result = gin.H{
				"status": 200,
				"token": tokenString,
				"result": "succeess generate password",
			}
		}
	} else {
		//login failed
		result = gin.H{
			"status": 400,
			"result": "username or password wrong",
		}
	}
	c.JSON(http.StatusOK, result)
}

func Index(c *gin.Context)  {
	c.JSON(200, gin.H{
		"status": 200,
		"message": "ok welcome to golang",
	})
}