package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200,"Hello World")
	})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"Who are you?")
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK,"Hello %s",name)
	})
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})
	// post
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password","000000")
		c.JSON(http.StatusOK,gin.H{
			"username": username,
			"password": password,
		})
	})

	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page","0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username","000000")

		c.JSON(http.StatusOK,gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})
	r.Run()
}
