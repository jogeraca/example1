package main

import (
	"./models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ORM orm.Ormer

func init() {
	models.ConnectToDb()
	ORM = models.GetOrmObject()
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.POST("/createUser", createUser)
	router.GET("/readUsers", readUsers)
	router.PUT("/updateUser", updateUser)
	router.DELETE("/deleteUser", deleteUser)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8080")
	// router.Run(":3000") for a hard coded port
}

func createUser(c *gin.Context) {
	var newUser models.Users
	c.BindJSON(&newUser)
	_, err := ORM.Insert(&newUser)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"email":     newUser.Email,
			"user_name": newUser.UserName,
			"user_id":   newUser.UserId})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
	}
}
func readUsers(c *gin.Context) {
	var user []models.Users
	fmt.Println(ORM)

	_, err := ORM.QueryTable("users").All(&user)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": &user})
		//c.JSON(http.StatusInternalServerError,
		//gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the users"})
	} else {
		//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": &user})
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the users"})
	}
}

func updateUser(c *gin.Context) {
	var updateUser models.Users
	c.BindJSON(&updateUser)
	_, err := ORM.QueryTable("users").Filter("email", updateUser.Email).Update(
		orm.Params{"user_name": updateUser.UserName,
			"password": updateUser.Password})
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to update the users"})
	}
}

func deleteUser(c *gin.Context) {
	var delUser models.Users
	c.BindJSON(&delUser)
	_, err := ORM.QueryTable("users").Filter("email", delUser.Email).Delete()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to delete the users"})
	}
}
