package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now(),
		})
	})

	r.Run()
}
