package tinyhttp

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine{
	return &Engine{router:make(map[string]HandlerFunc)}
}

func (e *Engine)AddRouter(method string, pattern string, handler HandlerFunc){
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine)Get(pattern string, handler HandlerFunc){
	e.AddRouter("GET", pattern, handler)
}

func (e *Engine)Post(pattern string, handler HandlerFunc){
	e.AddRouter("POST", pattern, handler)
}

func (e *Engine)Run(addr string)(err error){
	return http.ListenAndServe(addr, e)
}

func (e *Engine)ServeHTTP(w http.ResponseWriter, req *http.Request){
	key := req.Method + "-" + req.URL.Path
	if handler,ok := e.router[key];ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w,"404 NOT FOUND:%s\n", req.URL)
	}
}