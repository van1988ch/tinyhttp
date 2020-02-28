package tinyhttp

import (
	"log"
	"net/http"
)

type Router struct {
	router map[string]HandlerFunc
}

func newRouter()*Router  {
	return &Router{make(map[string]HandlerFunc)}
}

func (r *Router)AddRouter(method string, pattern string, handler HandlerFunc)  {
	log.Printf("Router %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.router[key] = handler
}

func (r *Router)Handle(c *Context)  {
	key := c.Method + "-" + c.Path
	if handler,ok := r.router[key];ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND:%s\n", c.Req.URL)
	}
}
