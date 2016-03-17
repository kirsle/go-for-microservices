package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	var DB map[string]string = map[string]string{}

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong\n")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":   "bar",
		"admin": "passwd",
	}))

	authorized.POST("/user", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON body.
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	r.Run(":8000")
}
