package gin_gee

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNew tests the creation of a new Engine
func TestNew(t *testing.T) {
	engine := New()
	if engine.router == nil {
		t.Fatal("Expected router to be initialized, got nil")
	}
}

// TestRouteRegistration tests the registration of GET and POST routes
func TestRouteRegistration(t *testing.T) {
	engine := New()
	engine.GET("/test", func(w http.ResponseWriter, r *http.Request) {})
	engine.POST("/test", func(w http.ResponseWriter, r *http.Request) {})

	if _, ok := engine.router["GET-/test"]; !ok {
		t.Fatal("Failed to register GET route")
	}

	if _, ok := engine.router["POST-/test"]; !ok {
		t.Fatal("Failed to register POST route")
	}
}

// testHandler is a helper function for handling test requests
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test passed"))
}

// TestServeHTTP tests the ServeHTTP method
func TestServeHTTP(t *testing.T) {
	engine := New()
	engine.GET("/test", testHandler)

	// Test route found
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	// Test route not found
	req = httptest.NewRequest("GET", "/notfound", nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	resp = w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected status 404, got %d", resp.StatusCode)
	}
}

// TestRun tests the Run method
func TestRun(t *testing.T) {
	engine := New()

	// Run in a goroutine to prevent blocking
	go func() {
		if err := engine.Run(":8080"); err != nil {
			t.Fatal("Failed to run server")
		}
	}()

	// Further test cases to ensure the server is running can be added here
}
