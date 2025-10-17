package websvrutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGroup_BasicGroupCreation tests basic group creation and route registration.
// TestGroup_BasicGroupCreation은 기본 그룹 생성 및 라우트 등록을 테스트합니다.
func TestGroup_BasicGroupCreation(t *testing.T) {
	app := New()

	// Create a group
	// 그룹 생성
	api := app.Group("/api")

	// Register routes in the group
	// 그룹에 라우트 등록
	api.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET users"))
	})

	api.POST("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST users"))
	})

	// Test GET /api/users
	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	if w.Body.String() != "GET users" {
		t.Errorf("Expected 'GET users', got %s", w.Body.String())
	}

	// Test POST /api/users
	req = httptest.NewRequest("POST", "/api/users", nil)
	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	if w.Body.String() != "POST users" {
		t.Errorf("Expected 'POST users', got %s", w.Body.String())
	}
}

// TestGroup_NestedGroups tests nested group creation with prefix concatenation.
// TestGroup_NestedGroups는 접두사 연결을 사용한 중첩 그룹 생성을 테스트합니다.
func TestGroup_NestedGroups(t *testing.T) {
	app := New()

	// Create nested groups
	// 중첩 그룹 생성
	api := app.Group("/api")
	v1 := api.Group("/v1")
	admin := v1.Group("/admin")

	// Register route in nested group
	// 중첩 그룹에 라우트 등록
	admin.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin users"))
	})

	// Test GET /api/v1/admin/users
	req := httptest.NewRequest("GET", "/api/v1/admin/users", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	if w.Body.String() != "admin users" {
		t.Errorf("Expected 'admin users', got %s", w.Body.String())
	}
}

// TestGroup_GroupMiddleware tests group-specific middleware application.
// TestGroup_GroupMiddleware는 그룹별 미들웨어 적용을 테스트합니다.
func TestGroup_GroupMiddleware(t *testing.T) {
	app := New()

	// Middleware that adds header
	// 헤더를 추가하는 미들웨어
	addHeader := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Test", "middleware")
			next.ServeHTTP(w, r)
		})
	}

	// Create group with middleware
	// 미들웨어가 있는 그룹 생성
	api := app.Group("/api")
	api.Use(addHeader)

	api.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("users"))
	})

	// Route outside group (no middleware)
	// 그룹 외부 라우트 (미들웨어 없음)
	app.GET("/public", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("public"))
	})

	// Test group route (should have middleware)
	// 그룹 라우트 테스트 (미들웨어 있어야 함)
	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Header().Get("X-Test") != "middleware" {
		t.Errorf("Expected X-Test header, got none")
	}

	// Test non-group route (should not have middleware)
	// 비그룹 라우트 테스트 (미들웨어 없어야 함)
	req = httptest.NewRequest("GET", "/public", nil)
	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Header().Get("X-Test") != "" {
		t.Errorf("Expected no X-Test header, got %s", w.Header().Get("X-Test"))
	}
}

// TestGroup_MiddlewareInheritance tests middleware inheritance in nested groups.
// TestGroup_MiddlewareInheritance는 중첩 그룹에서 미들웨어 상속을 테스트합니다.
func TestGroup_MiddlewareInheritance(t *testing.T) {
	app := New()

	executionOrder := []string{}

	// Parent middleware
	// 부모 미들웨어
	parentMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executionOrder = append(executionOrder, "parent")
			next.ServeHTTP(w, r)
		})
	}

	// Child middleware
	// 자식 미들웨어
	childMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executionOrder = append(executionOrder, "child")
			next.ServeHTTP(w, r)
		})
	}

	// Create parent group with middleware
	// 미들웨어가 있는 부모 그룹 생성
	api := app.Group("/api")
	api.Use(parentMiddleware)

	// Create child group with additional middleware
	// 추가 미들웨어가 있는 자식 그룹 생성
	v1 := api.Group("/v1")
	v1.Use(childMiddleware)

	v1.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		executionOrder = append(executionOrder, "handler")
		w.Write([]byte("test"))
	})

	// Execute request
	// 요청 실행
	req := httptest.NewRequest("GET", "/api/v1/test", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	// Verify execution order: parent -> child -> handler
	// 실행 순서 확인: parent -> child -> handler
	if len(executionOrder) != 3 {
		t.Fatalf("Expected 3 executions, got %d", len(executionOrder))
	}
	if executionOrder[0] != "parent" {
		t.Errorf("Expected first execution to be 'parent', got %s", executionOrder[0])
	}
	if executionOrder[1] != "child" {
		t.Errorf("Expected second execution to be 'child', got %s", executionOrder[1])
	}
	if executionOrder[2] != "handler" {
		t.Errorf("Expected third execution to be 'handler', got %s", executionOrder[2])
	}
}

// TestGroup_AllHTTPMethods tests all HTTP methods on groups.
// TestGroup_AllHTTPMethods는 그룹의 모든 HTTP 메서드를 테스트합니다.
func TestGroup_AllHTTPMethods(t *testing.T) {
	app := New()
	api := app.Group("/api")

	// Register all HTTP methods
	// 모든 HTTP 메서드 등록
	api.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET"))
	})
	api.POST("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST"))
	})
	api.PUT("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PUT"))
	})
	api.PATCH("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PATCH"))
	})
	api.DELETE("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DELETE"))
	})
	api.OPTIONS("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OPTIONS"))
	})
	api.HEAD("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Test each method
	// 각 메서드 테스트
	methods := []struct {
		method   string
		expected string
	}{
		{"GET", "GET"},
		{"POST", "POST"},
		{"PUT", "PUT"},
		{"PATCH", "PATCH"},
		{"DELETE", "DELETE"},
		{"OPTIONS", "OPTIONS"},
		{"HEAD", ""},
	}

	for _, m := range methods {
		req := httptest.NewRequest(m.method, "/api/test", nil)
		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("%s: Expected status 200, got %d", m.method, w.Code)
		}
		if m.method != "HEAD" && w.Body.String() != m.expected {
			t.Errorf("%s: Expected '%s', got '%s'", m.method, m.expected, w.Body.String())
		}
	}
}

// TestGroup_MethodChaining tests method chaining for fluent API.
// TestGroup_MethodChaining은 Fluent API를 위한 메서드 체이닝을 테스트합니다.
func TestGroup_MethodChaining(t *testing.T) {
	app := New()

	// Test method chaining
	// 메서드 체이닝 테스트
	api := app.Group("/api").
		GET("/users", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("users"))
		}).
		POST("/users", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("create user"))
		})

	// Verify both routes work
	// 두 라우트 모두 작동하는지 확인
	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Body.String() != "users" {
		t.Errorf("Expected 'users', got %s", w.Body.String())
	}

	req = httptest.NewRequest("POST", "/api/users", nil)
	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Body.String() != "create user" {
		t.Errorf("Expected 'create user', got %s", w.Body.String())
	}

	// Verify we can continue using the group
	// 그룹을 계속 사용할 수 있는지 확인
	api.PUT("/users/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("update user"))
	})

	req = httptest.NewRequest("PUT", "/api/users/123", nil)
	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Body.String() != "update user" {
		t.Errorf("Expected 'update user', got %s", w.Body.String())
	}
}

// TestGroup_MultipleMiddleware tests multiple middleware on a group.
// TestGroup_MultipleMiddleware는 그룹에 여러 미들웨어를 테스트합니다.
func TestGroup_MultipleMiddleware(t *testing.T) {
	app := New()

	executionOrder := []string{}

	// First middleware
	// 첫 번째 미들웨어
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executionOrder = append(executionOrder, "middleware1")
			next.ServeHTTP(w, r)
		})
	}

	// Second middleware
	// 두 번째 미들웨어
	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executionOrder = append(executionOrder, "middleware2")
			next.ServeHTTP(w, r)
		})
	}

	// Create group with multiple middleware
	// 여러 미들웨어가 있는 그룹 생성
	api := app.Group("/api")
	api.Use(middleware1, middleware2)

	api.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		executionOrder = append(executionOrder, "handler")
		w.Write([]byte("test"))
	})

	// Execute request
	// 요청 실행
	req := httptest.NewRequest("GET", "/api/test", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	// Verify execution order: middleware1 -> middleware2 -> handler
	// 실행 순서 확인: middleware1 -> middleware2 -> handler
	if len(executionOrder) != 3 {
		t.Fatalf("Expected 3 executions, got %d", len(executionOrder))
	}
	if executionOrder[0] != "middleware1" {
		t.Errorf("Expected first execution to be 'middleware1', got %s", executionOrder[0])
	}
	if executionOrder[1] != "middleware2" {
		t.Errorf("Expected second execution to be 'middleware2', got %s", executionOrder[1])
	}
	if executionOrder[2] != "handler" {
		t.Errorf("Expected third execution to be 'handler', got %s", executionOrder[2])
	}
}

// TestGroup_EmptyPrefix tests group with empty prefix.
// TestGroup_EmptyPrefix는 빈 접두사를 가진 그룹을 테스트합니다.
func TestGroup_EmptyPrefix(t *testing.T) {
	app := New()

	// Create group with empty prefix
	// 빈 접두사로 그룹 생성
	group := app.Group("")

	group.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})

	// Test route
	// 라우트 테스트
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	if w.Body.String() != "test" {
		t.Errorf("Expected 'test', got %s", w.Body.String())
	}
}

// TestGroup_DeepNesting tests deeply nested groups.
// TestGroup_DeepNesting은 깊게 중첩된 그룹을 테스트합니다.
func TestGroup_DeepNesting(t *testing.T) {
	app := New()

	// Create deeply nested groups
	// 깊게 중첩된 그룹 생성
	level1 := app.Group("/level1")
	level2 := level1.Group("/level2")
	level3 := level2.Group("/level3")
	level4 := level3.Group("/level4")

	level4.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("deep"))
	})

	// Test route
	// 라우트 테스트
	req := httptest.NewRequest("GET", "/level1/level2/level3/level4/test", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	if w.Body.String() != "deep" {
		t.Errorf("Expected 'deep', got %s", w.Body.String())
	}
}
