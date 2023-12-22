package gin_gee

import (
	"log"
	"net/http"
)

//type HandlerFunc func(*Context)

// router 是一个用于 HTTP 请求路由的结构体。
type router struct {
	handlers map[string]HandlerFunc // 存储所有路由处理函数的映射
}

// newRouter 创建并返回一个新的 router 实例。
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// addRoute 为指定的 HTTP 方法和路径添加一个路由处理函数。
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern) // 记录路由信息
	key := method + "-" + pattern                 // 创建路由键
	r.handlers[key] = handler                     // 将处理函数存储到映射中
}

// handle 处理一个 HTTP 请求。
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path // 从 Context 中获取请求方法和路径
	if handler, ok := r.handlers[key]; ok {
		handler(c) // 如果找到路由处理函数，执行它
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path) // 否则返回 404 状态
	}
}
