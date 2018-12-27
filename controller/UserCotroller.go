package controller

import (
	"fmt"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest-api-go-jwt/structs"
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
		statuscode int
		messages string
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		statuscode = 404
		messages = "Not Found"
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		statuscode = 200
		messages = "Success Get Data User"
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
}

// get all data in person
func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		persons []structs.User
		result  gin.H
		statuscode int
		messages string
	)

	idb.DB.Find(&persons)
	statuscode = 200
	messages = "Success Get All Users"
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
}

// create new data to the database
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		person structs.User
		result gin.H
		statuscode int
		messages string
	)

	user_name := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	person.Username = user_name
	person.Email = email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	person.Password = string(hashedPassword)
	idb.DB.Create(&person)
	statuscode = 200
	messages = "Success Register"
	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
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
		statuscode int
		messages string
	)

	err := idb.DB.First(&person, id).Error
	println(err)
	if err != nil {
		statuscode = 404
		messages = "data not found"
		result = gin.H{
			"result": "data not found",
		}
	}
	newPerson.Username = username
	newPerson.Email = email
	err = idb.DB.Model(&person).Updates(newPerson).Error
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
			"status": 200,
			"result": err,
		}
	}
	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
}

// delete data with {id}
func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		person structs.User
		result gin.H
		statuscode int
		messages string
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		statuscode = 400
		messages = "data not found"
		result = gin.H{}
	}
	println(id)
	err = idb.DB.Where("id = ?",id).Delete(&person).Error
	fmt.Println(err)
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

func (idb *InDB) Login(c *gin.Context){
	var (
		user structs.User
		result gin.H
		statuscode int
		messages string
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
			statuscode = 500
			messages = "failed generate password"
			result = gin.H{
				"error" : err,
			}
		}else{
			statuscode = 200
			messages = "succeess generate password"
			result = gin.H{
				"token": tokenString,
			}
		}
	} else {
		//login failed
		statuscode = 400
		messages = "username or password wrong"
		result = gin.H{
			"result": "username or password wrong",
		}
	}
	c.JSON(http.StatusOK, structs.Json{statuscode,messages,result})
}

func Index(c *gin.Context)  {
	c.JSON(200, gin.H{
		"status": 200,
		"message": "ok welcome to golang",
	})
}