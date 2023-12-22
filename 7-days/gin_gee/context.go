package gin_gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 是一个方便的别名，用于构造 JSON 数据。
type H map[string]interface{}

// Context 是一个封装了 HTTP 请求和响应的结构体。
type Context struct {
	// 原始对象
	Writer http.ResponseWriter // HTTP 响应
	Req    *http.Request       // HTTP 请求

	// 请求信息
	Path   string // 请求路径
	Method string // 请求方法

	// 响应信息
	StatusCode int // HTTP 状态码
}

// NewContext 创建并返回一个新的 Context 实例。
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 从表单中获取指定键的值。
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 从 URL 查询参数中获取指定键的值。
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置响应的 HTTP 状态码。
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置 HTTP 响应头。
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 发送一个字符串类型的响应。
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 发送一个 JSON 格式的响应。
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 发送一个二进制数据响应。
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 发送一个 HTML 格式的响应。
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
