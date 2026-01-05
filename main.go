package main

import (
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	name  string
	price float64
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
