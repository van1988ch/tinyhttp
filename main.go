package main

import (
	"log"
	"net/http"
	"time"
	"tinyhttp/tinyhttp"
)


func onlyForV2() tinyhttp.HandlerFunc {
	return func(c *tinyhttp.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

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
	v2.Use(onlyForV2())
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
