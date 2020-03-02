package main

import (
	"fmt"
	"html/template"
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

type student struct {
	Name string
	Age  int8
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {

	engine := tinyhttp.New()
	engine.Use(tinyhttp.Logger())

	engine.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	engine.LoadHTMLGlob("templates/*")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	engine.Static("/assets", "./static")

	engine.GET("/", func(c *tinyhttp.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	engine.GET("/students", func(c *tinyhttp.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", map[string]interface{}{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	engine.Get("/hello", func(c *tinyhttp.Context) {
		c.String(http.StatusOK, "hello %s, path %s", c.Query("name"), c.Path)
	})

	engine.Get("/hello/:name", func(c *tinyhttp.Context) {
		c.String(http.StatusOK, "hello %s, path %s", c.Param("name"), c.Path)
	})

	engine.Post("/login", func(c *tinyhttp.Context) {
		c.Json(http.StatusOK, map[string]string{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	v1 := engine.Group("/v1")
	{

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
