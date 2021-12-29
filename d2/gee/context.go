package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type H map[string]interface{}


// Context 上下文结构
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	Path   string
	Method string
	StatusCode int
}


func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}


// PostForm post
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}


// Query get请求参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}


// Status 响应状态
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}


// SetHeader 设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}


// string
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}


// JSON 返回 json
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}


// Data 返回数据流
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}


func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}


