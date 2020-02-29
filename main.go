package main

import (
	"net/http"
	"tinyhttp/tinyhttp"
)


func main()  {
	engine := tinyhttp.New()
	engine.Get("/", func(c *tinyhttp.Context) {
		c.HTML(http.StatusOK, "<h1>hello tinyhttp</h1>")
	})

	engine.Get("/hello", func(c *tinyhttp.Context) {
		c.String(http.StatusOK, "hello %s, path %s", c.Query("name"), c.Path)
	})

	engine.Get("/hello/:name", func(c *tinyhttp.Context) {
		c.String(http.StatusOK, "hello %s, path %s", c.Param("name"), c.Path)
	})

	engine.Get("/assets/*filepath", func(c *tinyhttp.Context) {
		c.String(http.StatusOK, "hello %s, path %s", c.Param("filepath"), c.Path)
	})

	engine.Post("/login", func(c *tinyhttp.Context) {
		c.Json(http.StatusOK, map[string]string{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	engine.Run(":9999")
}
