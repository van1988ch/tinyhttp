package tinyhttp

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix string
	middlewares []HandlerFunc
	parent *RouterGroup
	engine *Engine
}

type Engine struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup
}

func New() *Engine{
	 engine := &Engine{router:newRouter()}
	 engine.RouterGroup = &RouterGroup{engine:engine}
	 engine.groups = []*RouterGroup{engine.RouterGroup}
	 return engine
}

func (r *RouterGroup)Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		prefix:r.prefix + prefix,
		parent:r,
		engine:engine,
	}

	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (r *RouterGroup)addRouter(method string, comp string, handler HandlerFunc)  {
	pattern := r.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	r.engine.router.AddRouter(method, pattern, handler)
}

func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRouter("GET", pattern, handler)
}

func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRouter("POST", pattern, handler)
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