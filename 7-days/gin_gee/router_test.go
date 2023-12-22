package gin_gee

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试新建路由器
func TestNewRouter(t *testing.T) {
	r := newRouter()
	if r == nil {
		t.Fatal("Expected new router to be non-nil")
	}
}

// 测试添加路由
func TestAddRoute(t *testing.T) {
	engine := New()
	engine.GET("/test", func(c *Context) {
		c.String(http.StatusOK, "test")
	})
	engine.POST("/test", func(c *Context) {
		c.String(http.StatusOK, "test")
	})
	key := "GET-" + "/test"
	if _, exists := engine.router.handlers[key]; !exists {
		t.Fatalf("Route %s is not added", key)
	}
}

// 测试处理存在的路由
func TestRouterHandleExistingRoute(t *testing.T) {
	r := newRouter()
	pattern := "/test"
	r.addRoute("GET", pattern, func(c *Context) {
		c.String(http.StatusOK, "test")
	})

	req := httptest.NewRequest("GET", pattern, nil)
	w := httptest.NewRecorder()
	c := NewContext(w, req)

	r.handle(c)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
}

// 测试处理不存在的路由
func TestRouterHandleNonExistingRoute(t *testing.T) {
	r := newRouter()
	req := httptest.NewRequest("GET", "/notexist", nil)
	w := httptest.NewRecorder()
	c := NewContext(w, req)

	r.handle(c)

	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected status code 404, got %d", resp.StatusCode)
	}
}
