package tinyhttp

import (
	"log"
	"net/http"
	"strings"
)

type Router struct {
	router map[string]HandlerFunc
	roots map[string]*node
}

func newRouter()*Router  {
	return &Router{
		router:make(map[string]HandlerFunc),
		roots:make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _,item := range vs {
		if item != ""{
			parts = append(parts, item)
			if item[0] == '*'{
				break;
			}
		}
	}
	return parts
}

func (r *Router)AddRouter(method string, pattern string, handler HandlerFunc)  {
	log.Printf("Router %4s - %s", method, pattern)

	parts := parsePattern(pattern)
	key := method + "-" + pattern

	_,ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.router[key] = handler
}

func (r *Router)GetRouter(method string, path string) (*node, map[string]string){
	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok{
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1{
				params[part[1:]] = strings.Join(searchParts[index:],"/")
				break;
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *Router)Handle(c *Context)  {
	n,params := r.GetRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.router[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}

	c.Next()
}
