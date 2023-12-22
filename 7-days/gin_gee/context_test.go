package gin_gee

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPostForm 测试 PostForm 方法。
func TestPostForm(t *testing.T) {
	form := new(bytes.Buffer)
	form.WriteString("key=value")
	r := httptest.NewRequest("POST", "/", form)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	c := NewContext(w, r)

	if c.PostForm("key") != "value" {
		t.Errorf("PostForm key should be 'value', got '%s'", c.PostForm("key"))
	}
}

// TestQuery 测试 Query 方法。
func TestQuery(t *testing.T) {
	r := httptest.NewRequest("GET", "/?key=value", nil)
	w := httptest.NewRecorder()
	c := NewContext(w, r)

	if c.Query("key") != "value" {
		t.Errorf("Query key should be 'value', got '%s'", c.Query("key"))
	}
}

// TestString 测试 String 方法。
func TestString(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	c := NewContext(w, r)

	c.String(http.StatusOK, "Hello %s", "world")

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "Hello world" {
		t.Errorf("Expected body to be 'Hello world', got '%s'", string(body))
	}
}

// TestJSON 测试 JSON 方法。
func TestJSON(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	c := NewContext(w, r)

	c.JSON(http.StatusOK, H{"hello": "world"})

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var data map[string]string
	json.Unmarshal(body, &data)

	if data["hello"] != "world" {
		t.Errorf("Expected JSON 'hello' key to be 'world', got '%s'", data["hello"])
	}
}

// TestData 测试 Data 方法。
func TestData(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	c := NewContext(w, r)

	c.Data(http.StatusOK, []byte("hello world"))

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "hello world" {
		t.Errorf("Expected body to be 'hello world', got '%s'", string(body))
	}
}

// TestHTML 测试 HTML 方法。
func TestHTML(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	c := NewContext(w, r)

	c.HTML(http.StatusOK, "<h1>Hello World</h1>")

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "<h1>Hello World</h1>" {
		t.Errorf("Expected HTML to be '<h1>Hello World</h1>', got '%s'", string(body))
	}
}
