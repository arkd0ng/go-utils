package websvrutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewRouter tests creating a new router.
// TestNewRouter는 새 라우터 생성을 테스트합니다.
func TestNewRouter(t *testing.T) {
	router := newRouter()

	if router == nil {
		t.Fatal("newRouter() returned nil")
	}

	if router.routes == nil {
		t.Fatal("routes map is nil")
	}

	if router.notFoundHandler == nil {
		t.Fatal("notFoundHandler is nil")
	}
}

// TestRouterGET tests registering a GET route.
// TestRouterGET은 GET 라우트 등록을 테스트합니다.
func TestRouterGET(t *testing.T) {
	router := newRouter()

	called := false
	router.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if !called {
		t.Error("handler was not called")
	}

	if rec.Code != http.StatusOK {
		t.Errorf("status code = %d, want %d", rec.Code, http.StatusOK)
	}
}

// TestRouterPOST tests registering a POST route.
// TestRouterPOST는 POST 라우트 등록을 테스트합니다.
func TestRouterPOST(t *testing.T) {
	router := newRouter()

	called := false
	router.POST("/users", func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusCreated)
	})

	req := httptest.NewRequest("POST", "/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if !called {
		t.Error("handler was not called")
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("status code = %d, want %d", rec.Code, http.StatusCreated)
	}
}

// TestRouterAllMethods tests all HTTP methods.
// TestRouterAllMethods는 모든 HTTP 메서드를 테스트합니다.
func TestRouterAllMethods(t *testing.T) {
	router := newRouter()

	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"}

	for _, method := range methods {
		called := false
		handler := func(w http.ResponseWriter, r *http.Request) {
			called = true
			w.WriteHeader(http.StatusOK)
		}

		switch method {
		case "GET":
			router.GET("/test", handler)
		case "POST":
			router.POST("/test", handler)
		case "PUT":
			router.PUT("/test", handler)
		case "PATCH":
			router.PATCH("/test", handler)
		case "DELETE":
			router.DELETE("/test", handler)
		case "OPTIONS":
			router.OPTIONS("/test", handler)
		case "HEAD":
			router.HEAD("/test", handler)
		}

		req := httptest.NewRequest(method, "/test", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if !called {
			t.Errorf("%s handler was not called", method)
		}
	}
}

// TestRouterParameterExtraction tests parameter extraction from URLs.
// TestRouterParameterExtraction은 URL에서 매개변수 추출을 테스트합니다.
func TestRouterParameterExtraction(t *testing.T) {
	router := newRouter()

	router.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
		// Parameters will be accessible via Context in v1.11.004
		// 매개변수는 v1.11.004에서 Context를 통해 액세스 가능
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/users/123", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status code = %d, want %d", rec.Code, http.StatusOK)
	}
}

// TestRouterWildcard tests wildcard route matching.
// TestRouterWildcard는 와일드카드 라우트 일치를 테스트합니다.
func TestRouterWildcard(t *testing.T) {
	router := newRouter()

	called := false
	router.GET("/files/*", func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	tests := []string{
		"/files/images/logo.png",
		"/files/docs/manual.pdf",
		"/files/a/b/c/d/e.txt",
	}

	for _, path := range tests {
		req := httptest.NewRequest("GET", path, nil)
		rec := httptest.NewRecorder()

		called = false
		router.ServeHTTP(rec, req)

		if !called {
			t.Errorf("wildcard handler not called for path: %s", path)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("status code = %d for path %s", rec.Code, path)
		}
	}
}

// TestRouterNotFound tests 404 Not Found response.
// TestRouterNotFound는 404 Not Found 응답을 테스트합니다.
func TestRouterNotFound(t *testing.T) {
	router := newRouter()

	router.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/nonexistent", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("status code = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

// TestRouterCustomNotFound tests custom 404 handler.
// TestRouterCustomNotFound는 커스텀 404 핸들러를 테스트합니다.
func TestRouterCustomNotFound(t *testing.T) {
	router := newRouter()

	customCalled := false
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		customCalled = true
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Custom 404"))
	})

	req := httptest.NewRequest("GET", "/nonexistent", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if !customCalled {
		t.Error("custom not found handler was not called")
	}

	if rec.Body.String() != "Custom 404" {
		t.Errorf("response body = %s, want 'Custom 404'", rec.Body.String())
	}
}

// TestRouterMethodNotAllowed tests that different methods on same path don't interfere.
// TestRouterMethodNotAllowed는 동일 경로의 다른 메서드가 간섭하지 않는지 테스트합니다.
func TestRouterMethodNotAllowed(t *testing.T) {
	router := newRouter()

	router.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// POST to a GET-only route should return 404
	// GET 전용 라우트에 POST하면 404를 반환해야 함
	req := httptest.NewRequest("POST", "/test", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("status code = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

// TestRouteMatching tests the route matching logic.
// TestRouteMatching은 라우트 일치 로직을 테스트합니다.
func TestRouteMatching(t *testing.T) {
	tests := []struct {
		name      string
		pattern   string
		path      string
		shouldMatch bool
	}{
		{"exact match", "/users", "/users", true},
		{"exact mismatch", "/users", "/posts", false},
		{"parameter match", "/users/:id", "/users/123", true},
		{"parameter mismatch length", "/users/:id", "/users/123/posts", false},
		{"multiple parameters", "/users/:id/posts/:postId", "/users/123/posts/456", true},
		{"wildcard match", "/files/*", "/files/a/b/c", true},
		{"wildcard root", "/files/*", "/files/", true},
		{"trailing slash", "/users/", "/users", true},
		{"root path", "/", "/", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			route := &Route{
				Pattern:  tt.pattern,
				segments: parsePattern(tt.pattern),
			}

			_, matched := route.match(tt.path)

			if matched != tt.shouldMatch {
				t.Errorf("pattern %q with path %q: matched = %v, want %v",
					tt.pattern, tt.path, matched, tt.shouldMatch)
			}
		})
	}
}

// TestParsePattern tests URL pattern parsing.
// TestParsePattern은 URL 패턴 파싱을 테스트합니다.
func TestParsePattern(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		expected []segment
	}{
		{
			name:    "simple path",
			pattern: "/users",
			expected: []segment{
				{value: "users", isParam: false, isWildcard: false},
			},
		},
		{
			name:    "path with parameter",
			pattern: "/users/:id",
			expected: []segment{
				{value: "users", isParam: false, isWildcard: false},
				{value: "id", isParam: true, isWildcard: false},
			},
		},
		{
			name:    "path with wildcard",
			pattern: "/files/*",
			expected: []segment{
				{value: "files", isParam: false, isWildcard: false},
				{value: "", isParam: false, isWildcard: true},
			},
		},
		{
			name:     "root path",
			pattern:  "/",
			expected: []segment{},
		},
		{
			name:    "multiple parameters",
			pattern: "/users/:userId/posts/:postId",
			expected: []segment{
				{value: "users", isParam: false, isWildcard: false},
				{value: "userId", isParam: true, isWildcard: false},
				{value: "posts", isParam: false, isWildcard: false},
				{value: "postId", isParam: true, isWildcard: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parsePattern(tt.pattern)

			if len(result) != len(tt.expected) {
				t.Fatalf("segment count = %d, want %d", len(result), len(tt.expected))
			}

			for i, seg := range result {
				exp := tt.expected[i]
				if seg.value != exp.value || seg.isParam != exp.isParam || seg.isWildcard != exp.isWildcard {
					t.Errorf("segment[%d] = %+v, want %+v", i, seg, exp)
				}
			}
		})
	}
}

// TestParsePath tests URL path parsing.
// TestParsePath는 URL 경로 파싱을 테스트합니다.
func TestParsePath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected []string
	}{
		{"simple path", "/users", []string{"users"}},
		{"multiple segments", "/users/123/posts", []string{"users", "123", "posts"}},
		{"root path", "/", []string{}},
		{"trailing slash", "/users/", []string{"users"}},
		{"no leading slash", "users/123", []string{"users", "123"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parsePath(tt.path)

			if len(result) != len(tt.expected) {
				t.Fatalf("segment count = %d, want %d", len(result), len(tt.expected))
			}

			for i, seg := range result {
				if seg != tt.expected[i] {
					t.Errorf("segment[%d] = %s, want %s", i, seg, tt.expected[i])
				}
			}
		})
	}
}

// TestAppRouterIntegration tests router integration with App.
// TestAppRouterIntegration은 App과 라우터 통합을 테스트합니다.
func TestAppRouterIntegration(t *testing.T) {
	app := New()

	getCalled := false
	postCalled := false

	app.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		getCalled = true
		w.WriteHeader(http.StatusOK)
	})

	app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
		postCalled = true
		w.WriteHeader(http.StatusCreated)
	})

	// Test GET
	req := httptest.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if !getCalled {
		t.Error("GET handler was not called")
	}

	if rec.Code != http.StatusOK {
		t.Errorf("GET status = %d, want %d", rec.Code, http.StatusOK)
	}

	// Test POST
	req = httptest.NewRequest("POST", "/users", nil)
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if !postCalled {
		t.Error("POST handler was not called")
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("POST status = %d, want %d", rec.Code, http.StatusCreated)
	}
}

// BenchmarkRouterSimpleRoute benchmarks simple route matching.
// BenchmarkRouterSimpleRoute는 간단한 라우트 일치를 벤치마크합니다.
func BenchmarkRouterSimpleRoute(b *testing.B) {
	router := newRouter()
	router.GET("/users", func(w http.ResponseWriter, r *http.Request) {})

	req := httptest.NewRequest("GET", "/users", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
	}
}

// BenchmarkRouterParameterRoute benchmarks parameter route matching.
// BenchmarkRouterParameterRoute는 매개변수 라우트 일치를 벤치마크합니다.
func BenchmarkRouterParameterRoute(b *testing.B) {
	router := newRouter()
	router.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {})

	req := httptest.NewRequest("GET", "/users/123", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
	}
}

// BenchmarkRouterWildcardRoute benchmarks wildcard route matching.
// BenchmarkRouterWildcardRoute는 와일드카드 라우트 일치를 벤치마크합니다.
func BenchmarkRouterWildcardRoute(b *testing.B) {
	router := newRouter()
	router.GET("/files/*", func(w http.ResponseWriter, r *http.Request) {})

	req := httptest.NewRequest("GET", "/files/a/b/c/d/e.txt", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
	}
}

// BenchmarkParsePattern benchmarks pattern parsing.
// BenchmarkParsePattern은 패턴 파싱을 벤치마크합니다.
func BenchmarkParsePattern(b *testing.B) {
	pattern := "/users/:userId/posts/:postId/comments/:commentId"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parsePattern(pattern)
	}
}

// BenchmarkParsePath benchmarks path parsing.
// BenchmarkParsePath는 경로 파싱을 벤치마크합니다.
func BenchmarkParsePath(b *testing.B) {
	path := "/users/123/posts/456/comments/789"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parsePath(path)
	}
}
