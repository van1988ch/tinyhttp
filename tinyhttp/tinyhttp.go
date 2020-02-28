package tinyhttp

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *Router
}

func New() *Engine{
	return &Engine{router:newRouter()}
}

func (e *Engine)addRouter(method string, pattern string, handler HandlerFunc){
	e.router.AddRouter(method, pattern, handler)
}

func (e *Engine)Get(pattern string, handler HandlerFunc){
	e.addRouter("GET", pattern, handler)
}

func (e *Engine)Post(pattern string, handler HandlerFunc){
	e.addRouter("POST", pattern, handler)
}

func (e *Engine)Run(addr string)(err error){
	return http.ListenAndServe(addr, e)
}

func (e *Engine)ServeHTTP(w http.ResponseWriter, req *http.Request){
	c := newContext(w,req)
	e.router.Handle(c)
}