package config

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result gin.H;
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			result = gin.H{
				"status": 401,
				"result": "Token Not Valid",
			}
			c.JSON(http.StatusOK, result)
			c.AbortWithError(401, err)
		}
	}
}
