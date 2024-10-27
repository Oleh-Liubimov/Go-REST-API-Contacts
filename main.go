package main

import (
	"go-rest-api/db/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.InitDB()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
