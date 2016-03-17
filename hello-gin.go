package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/echo", func(c *gin.Context) {
		// Parse JSON body
		var json struct {
			Input string `json:"in" binding:"required"`
		}

		if c.Bind(&json) == nil {
			c.JSON(200, gin.H{"out": json.Input})
		}
	})

	r.Run(":8000")
}
