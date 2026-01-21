package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Name":    "Alwi",
			"Bio":     "Day 1 learning Golang with Gin",
			"message": "Hello World!",
		})
	})

	router.Run()
}
