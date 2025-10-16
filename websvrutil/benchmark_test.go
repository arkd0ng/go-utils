package websvrutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// BenchmarkContextGetValue benchmarks Context.Get operation.
// BenchmarkContextGetValue는 Context.Get 작업을 벤치마크합니다.
func BenchmarkContextGetValue(b *testing.B) {
	ctx := NewContext(nil, nil)
	ctx.Set("key", "value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.Get("key")
	}
}

// BenchmarkContextSetValue benchmarks Context.Set operation.
// BenchmarkContextSetValue는 Context.Set 작업을 벤치마크합니다.
func BenchmarkContextSetValue(b *testing.B) {
	ctx := NewContext(nil, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.Set("key", "value")
	}
}

// BenchmarkJSONResponse benchmarks JSON response rendering.
// BenchmarkJSONResponse는 JSON 응답 렌더링을 벤치마크합니다.
func BenchmarkJSONResponse(b *testing.B) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	ctx := NewContext(rec, req)

	data := map[string]interface{}{
		"message": "success",
		"code":    200,
		"data":    []int{1, 2, 3, 4, 5},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.JSON(200, data)
	}
}

// BenchmarkRouting benchmarks simple routing.
// BenchmarkRouting은 간단한 라우팅을 벤치마크합니다.
func BenchmarkRouting(b *testing.B) {
	app := New(WithTemplateDir(""))
	app.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)
	}
}

// BenchmarkRoutingWithParams benchmarks routing with parameters.
// BenchmarkRoutingWithParams는 파라미터가 있는 라우팅을 벤치마크합니다.
func BenchmarkRoutingWithParams(b *testing.B) {
	app := New(WithTemplateDir(""))
	app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		_ = ctx.Param("id")
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/users/123", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)
	}
}

// BenchmarkMiddleware benchmarks middleware execution.
// BenchmarkMiddleware는 미들웨어 실행을 벤치마크합니다.
func BenchmarkMiddleware(b *testing.B) {
	app := New(WithTemplateDir(""))
	app.Use(Logger())
	app.Use(Recovery())

	app.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)
	}
}

// BenchmarkCSRFToken benchmarks CSRF token generation.
// BenchmarkCSRFToken은 CSRF 토큰 생성을 벤치마크합니다.
func BenchmarkCSRFToken(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generateCSRFToken(32)
	}
}

// BenchmarkValidator benchmarks validation.
// BenchmarkValidator는 검증을 벤치마크합니다.
func BenchmarkValidator(b *testing.B) {
	type User struct {
		Name  string `validate:"required,min=3,max=50"`
		Email string `validate:"required,email"`
		Age   int    `validate:"required,gte=18,lte=100"`
	}

	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   25,
	}

	validator := &DefaultValidator{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator.Validate(&user)
	}
}

// BenchmarkClientIP benchmarks ClientIP extraction.
// BenchmarkClientIP는 ClientIP 추출을 벤치마크합니다.
func BenchmarkClientIP(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "203.0.113.195, 70.41.3.18, 150.172.238.178")
	req.RemoteAddr = "192.168.1.1:12345"
	ctx := NewContext(nil, req)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.ClientIP()
	}
}

// BenchmarkParamExtraction benchmarks URL parameter extraction.
// BenchmarkParamExtraction은 URL 파라미터 추출을 벤치마크합니다.
func BenchmarkParamExtraction(b *testing.B) {
	ctx := NewContext(nil, nil)
	params := map[string]string{
		"id":     "123",
		"name":   "test",
		"action": "view",
	}
	ctx.setParams(params)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ctx.Param("id")
		_ = ctx.Param("name")
		_ = ctx.Param("action")
	}
}
