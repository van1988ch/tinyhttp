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
	
	
	v1 := engine.Group("/v1")
	{
		v1.GET("/", func(c *tinyhttp.Context) {
			c.HTML(http.StatusOK, "<h1>Hello</h1>")
		})
		
		v1.GET("/hello", func(c *tinyhttp.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := engine.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *tinyhttp.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *tinyhttp.Context) {
			c.Json(http.StatusOK, map[string]string{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	engine.Run(":9999")
}
