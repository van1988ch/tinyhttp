package tinyhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	Params map[string]string
	StatusCode int

	handlers []HandlerFunc
	index int
}

func (c *Context)Param(key string)string {
	value, _ := c.Params[key]
	return value
}

//中间件可等待用户自己定义的 Handler处理结束后，再做一些额外的操作
func (c *Context)Next()  {
	c.index++
	s := len(c.handlers)
	for ; c.index < s ; c.index++  {
		c.handlers[c.index](c)
	}
}

func newContext(w http.ResponseWriter, req *http.Request)*Context  {
	return &Context{
		Writer:w,
		Req:req,
		Path:req.URL.Path,
		Method:req.Method,
		index:-1,
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.Json(code, map[string]interface{}{"message": err})
}

func (c *Context)PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context)Query(key string) string  {
	return c.Req.URL.Query().Get(key)
}

func (c *Context)Status(code int)  {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context)SetHeader(key string, value string)  {
	c.Writer.Header().Set(key, value)
}

func (c *Context)String(code int, format string, values ...interface{})  {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context)Json(code int, obj interface{})  {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context)Data(code int, data []byte)  {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context)HTML(code int, html string)  {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}