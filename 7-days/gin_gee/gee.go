package gin_gee

import "net/http"

// HandlerFunc 定义路由处理函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现 ServeHTTP 方法的接口
type Engine struct {
	router map[string]HandlerFunc
}

// New 新建 Engine 实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 添加路由
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// GET 添加 GET 请求路由
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST 添加 POST 请求路由
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// Run 启动 HTTP 服务
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP 实现 http.Handler 接口的方法
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("404 NOT FOUND: " + req.URL.Path))
	}
}
