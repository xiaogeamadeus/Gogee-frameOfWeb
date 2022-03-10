package main

import (
	"gogee"
	"net/http"
)

func main() {
	r := gogee.New()
	r.GET("/", func(c *gogee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gogee</h1>")
	})

	r.GET("/hello", func(c *gogee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gogee.Context) {
		c.JSON(http.StatusOK, gogee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
