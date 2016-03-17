package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
)

type User struct {
	// gorm.Model
	ID        int64         `gorm:"AUTO_INCREMENT" json:"id"`
	Username  string        `json:"username"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Address   Address       `json:"address"`
	AddressID sql.NullInt64 `json:"addressId"`
}

type Address struct {
	// gorm.Model
	ID      int64  `gorm:"AUTO_INCREMENT" json:"id"`
	UserID  int64  `gorm:"index" json:"userId"`
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `gorm:"size:2" json:"state"`
	Zipcode string `json:"zip"`
}

var USER_NOT_FOUND gin.H = gin.H{
	"status": "error",
	"error":  "User not found.",
}

func main() {
	// Connect to the DB.
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}

	// Make tables.
	if db.HasTable(&User{}) != true {
		db.CreateTable(&User{})
	}
	if db.HasTable(&Address{}) != true {
		db.CreateTable(&Address{})
	}

	// Get the router from Gin.
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {

		var users []User
		db.Preload("Address").Find(&users)

		results := gin.H{
			"status": "ok",
			"users":  users,
		}
		c.JSON(200, results)
	})

	r.POST("/users", func(c *gin.Context) {
		json := User{}
		if c.Bind(&json) == nil {
			// Make sure the username isn't already taken!
			username := json.Username
			exists := []User{}
			db.Where("username = ?", username).Find(&exists)
			if len(exists) != 0 {
				c.JSON(409, gin.H{
					"status": "error",
					"error":  "Username already exists.",
				})
				return
			}

			// Create the user.
			db.Create(&User{
				Username:  json.Username,
				FirstName: json.FirstName,
				LastName:  json.LastName,
			})
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	r.GET("/user/:id", func(c *gin.Context) {
		userId := c.Params.ByName("id")

		user := User{}
		db.Where("id = ?", userId).Preload("Address").First(&user)

		// Not valid user?
		if user.ID == 0 {
			c.JSON(404, USER_NOT_FOUND)
			return
		}

		c.JSON(200, gin.H{
			"status": "ok",
			"user":   user,
		})
	})

	r.GET("/user/:id/address", func(c *gin.Context) {
		userId := c.Params.ByName("id")

		user := User{}
		db.Where("id = ?", userId).Preload("Address").First(&user)
		if user.ID == 0 {
			c.JSON(404, USER_NOT_FOUND)
			return
		}

		// User has no address?
		if user.Address.ID == 0 {
			c.JSON(404, gin.H{
				"status": "error",
				"error":  "User has no address on file.",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  "ok",
			"address": user.Address,
		})
	})

	r.POST("/user/:id/address", func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"status": "error",
				"error":  "User ID must be a number.",
			})
			return
		}

		user := User{}
		db.Where("id = ?", userId).First(&user)
		if user.ID == 0 {
			c.JSON(404, USER_NOT_FOUND)
			return
		}

		// Parse the address JSON.
		json := Address{}
		if c.Bind(&json) == nil {
			db.Create(&Address{
				UserID:  userId,
				Street1: json.Street1,
				Street2: json.Street2,
				City:    json.City,
				State:   json.State,
				Zipcode: json.Zipcode,
			})
			c.JSON(201, gin.H{"status": "ok"})
		}
	})

	r.Run(":8000")
}
