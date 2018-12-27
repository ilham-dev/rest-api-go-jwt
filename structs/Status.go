package structs

import "github.com/gin-gonic/gin"

type Json struct {
	Status int
	Message string
	Data gin.H
}