package websvrutil

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewContext tests creating a new Context.
// TestNewContext는 새 Context 생성을 테스트합니다.
func TestNewContext(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)

	if ctx == nil {
		t.Fatal("NewContext() returned nil")
	}

	if ctx.Request != req {
		t.Error("Request not set correctly")
	}

	if ctx.ResponseWriter != rec {
		t.Error("ResponseWriter not set correctly")
	}

	if ctx.params == nil {
		t.Error("params map is nil")
	}

	if ctx.values == nil {
		t.Error("values map is nil")
	}
}

// TestContextParam tests parameter retrieval.
// TestContextParam은 매개변수 검색을 테스트합니다.
func TestContextParam(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/123", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)
	ctx.setParams(map[string]string{
		"id":   "123",
		"name": "john",
	})

	if ctx.Param("id") != "123" {
		t.Errorf("Param(\"id\") = %s, want \"123\"", ctx.Param("id"))
	}

	if ctx.Param("name") != "john" {
		t.Errorf("Param(\"name\") = %s, want \"john\"", ctx.Param("name"))
	}

	if ctx.Param("nonexistent") != "" {
		t.Error("Param(\"nonexistent\") should return empty string")
	}
}

// TestContextParams tests retrieving all parameters.
// TestContextParams는 모든 매개변수 검색을 테스트합니다.
func TestContextParams(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)
	ctx.setParams(map[string]string{
		"id":     "123",
		"postId": "456",
	})

	params := ctx.Params()

	if len(params) != 2 {
		t.Fatalf("Params() length = %d, want 2", len(params))
	}

	if params["id"] != "123" {
		t.Errorf("params[\"id\"] = %s, want \"123\"", params["id"])
	}

	if params["postId"] != "456" {
		t.Errorf("params[\"postId\"] = %s, want \"456\"", params["postId"])
	}

	// Test that modifying returned map doesn't affect original
	// 반환된 맵 수정이 원본에 영향을 주지 않는지 테스트
	params["id"] = "999"
	if ctx.Param("id") != "123" {
		t.Error("Params() should return a copy, not the original map")
	}
}

// TestContextSetGet tests setting and getting values.
// TestContextSetGet은 값 설정 및 가져오기를 테스트합니다.
func TestContextSetGet(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)

	// Test Set and Get
	// Set 및 Get 테스트
	ctx.Set("user", "john")
	ctx.Set("count", 42)
	ctx.Set("active", true)

	value, exists := ctx.Get("user")
	if !exists {
		t.Error("Get(\"user\") should exist")
	}
	if value != "john" {
		t.Errorf("Get(\"user\") = %v, want \"john\"", value)
	}

	value, exists = ctx.Get("count")
	if !exists {
		t.Error("Get(\"count\") should exist")
	}
	if value != 42 {
		t.Errorf("Get(\"count\") = %v, want 42", value)
	}

	// Test non-existent key
	// 존재하지 않는 키 테스트
	_, exists = ctx.Get("nonexistent")
	if exists {
		t.Error("Get(\"nonexistent\") should not exist")
	}
}

// TestContextMustGet tests MustGet method.
// TestContextMustGet은 MustGet 메서드를 테스트합니다.
func TestContextMustGet(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)
	ctx.Set("user", "john")

	value := ctx.MustGet("user")
	if value != "john" {
		t.Errorf("MustGet(\"user\") = %v, want \"john\"", value)
	}

	// Test panic on non-existent key
	// 존재하지 않는 키에 대한 패닉 테스트
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustGet should panic for non-existent key")
		}
	}()
	ctx.MustGet("nonexistent")
}

// TestContextGetTyped tests typed getter methods.
// TestContextGetTyped는 타입 지정 getter 메서드를 테스트합니다.
func TestContextGetTyped(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)
	ctx.Set("name", "john")
	ctx.Set("age", 25)
	ctx.Set("active", true)

	// Test GetString
	if ctx.GetString("name") != "john" {
		t.Errorf("GetString(\"name\") = %s, want \"john\"", ctx.GetString("name"))
	}

	if ctx.GetString("nonexistent") != "" {
		t.Error("GetString for non-existent key should return empty string")
	}

	// Test GetInt
	if ctx.GetInt("age") != 25 {
		t.Errorf("GetInt(\"age\") = %d, want 25", ctx.GetInt("age"))
	}

	if ctx.GetInt("nonexistent") != 0 {
		t.Error("GetInt for non-existent key should return 0")
	}

	// Test GetBool
	if ctx.GetBool("active") != true {
		t.Error("GetBool(\"active\") should be true")
	}

	if ctx.GetBool("nonexistent") != false {
		t.Error("GetBool for non-existent key should return false")
	}
}

// TestContextQuery tests query parameter retrieval.
// TestContextQuery는 쿼리 매개변수 검색을 테스트합니다.
func TestContextQuery(t *testing.T) {
	req := httptest.NewRequest("GET", "/search?q=golang&page=2&limit=10", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)

	if ctx.Query("q") != "golang" {
		t.Errorf("Query(\"q\") = %s, want \"golang\"", ctx.Query("q"))
	}

	if ctx.Query("page") != "2" {
		t.Errorf("Query(\"page\") = %s, want \"2\"", ctx.Query("page"))
	}

	if ctx.Query("limit") != "10" {
		t.Errorf("Query(\"limit\") = %s, want \"10\"", ctx.Query("limit"))
	}

	if ctx.Query("nonexistent") != "" {
		t.Error("Query for non-existent parameter should return empty string")
	}
}

// TestContextQueryDefault tests query parameter with default value.
// TestContextQueryDefault는 기본값이 있는 쿼리 매개변수를 테스트합니다.
func TestContextQueryDefault(t *testing.T) {
	req := httptest.NewRequest("GET", "/search?q=golang", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)

	if ctx.QueryDefault("q", "default") != "golang" {
		t.Error("QueryDefault should return actual value when it exists")
	}

	if ctx.QueryDefault("page", "1") != "1" {
		t.Error("QueryDefault should return default value when parameter doesn't exist")
	}
}

// TestContextHeader tests request header retrieval.
// TestContextHeader는 요청 헤더 검색을 테스트합니다.
func TestContextHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer token123")
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ctx := NewContext(rec, req)

	if ctx.Header("Authorization") != "Bearer token123" {
		t.Errorf("Header(\"Authorization\") = %s", ctx.Header("Authorization"))
	}

	if ctx.Header("Content-Type") != "application/json" {
		t.Errorf("Header(\"Content-Type\") = %s", ctx.Header("Content-Type"))
	}

	if ctx.Header("Nonexistent") != "" {
		t.Error("Header for non-existent header should return empty string")
	}
}

// TestContextSetHeader tests setting response headers.
// TestContextSetHeader는 응답 헤더 설정을 테스트합니다.
func TestContextSetHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)
	ctx.SetHeader("Content-Type", "application/json")
	ctx.SetHeader("X-Custom-Header", "custom-value")

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Error("SetHeader should set Content-Type header")
	}

	if rec.Header().Get("X-Custom-Header") != "custom-value" {
		t.Error("SetHeader should set custom header")
	}
}

// TestContextStatus tests setting response status.
// TestContextStatus는 응답 상태 설정을 테스트합니다.
func TestContextStatus(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)
	ctx.Status(http.StatusCreated)

	if rec.Code != http.StatusCreated {
		t.Errorf("Status code = %d, want %d", rec.Code, http.StatusCreated)
	}
}

// TestContextWrite tests writing data to response.
// TestContextWrite는 응답에 데이터 쓰기를 테스트합니다.
func TestContextWrite(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)

	data := []byte("Hello, World!")
	n, err := ctx.Write(data)

	if err != nil {
		t.Errorf("Write error: %v", err)
	}

	if n != len(data) {
		t.Errorf("Write returned %d bytes, want %d", n, len(data))
	}

	if rec.Body.String() != "Hello, World!" {
		t.Errorf("Response body = %s, want \"Hello, World!\"", rec.Body.String())
	}
}

// TestContextWriteString tests writing string to response.
// TestContextWriteString은 응답에 문자열 쓰기를 테스트합니다.
func TestContextWriteString(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)

	str := "Hello, World!"
	n, err := ctx.WriteString(str)

	if err != nil {
		t.Errorf("WriteString error: %v", err)
	}

	if n != len(str) {
		t.Errorf("WriteString returned %d bytes, want %d", n, len(str))
	}

	if rec.Body.String() != str {
		t.Errorf("Response body = %s, want %s", rec.Body.String(), str)
	}
}

// TestContextMethodPath tests Method and Path methods.
// TestContextMethodPath는 Method 및 Path 메서드를 테스트합니다.
func TestContextMethodPath(t *testing.T) {
	req := httptest.NewRequest("POST", "/users/123", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)

	if ctx.Method() != "POST" {
		t.Errorf("Method() = %s, want \"POST\"", ctx.Method())
	}

	if ctx.Path() != "/users/123" {
		t.Errorf("Path() = %s, want \"/users/123\"", ctx.Path())
	}
}

// TestContextContext tests Context() method.
// TestContextContext는 Context() 메서드를 테스트합니다.
func TestContextContext(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)
	reqCtx := ctx.Context()

	if reqCtx != req.Context() {
		t.Error("Context() should return request's context")
	}
}

// TestContextWithContext tests WithContext method.
// TestContextWithContext는 WithContext 메서드를 테스트합니다.
func TestContextWithContext(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	ctx := NewContext(rec, req)

	newCtx := context.WithValue(context.Background(), "key", "value")
	ctx2 := ctx.WithContext(newCtx)

	// Should be a new context
	// 새 컨텍스트여야 함
	if ctx2 == ctx {
		t.Error("WithContext should return a new Context")
	}

	// Should have the new context.Context
	// 새 context.Context를 가져야 함
	if ctx2.Request.Context() != newCtx {
		t.Error("WithContext should update request's context")
	}
}

// TestGetContext tests retrieving Context from request.
// TestGetContext는 요청에서 Context 검색을 테스트합니다.
func TestGetContext(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/123", nil)
	rec := httptest.NewRecorder()

	// Create context and store in request
	// 컨텍스트 생성 및 요청에 저장
	ctx := NewContext(rec, req)
	ctx.setParams(map[string]string{"id": "123"})

	reqWithCtx := req.WithContext(contextWithValue(req.Context(), ctx))

	// Retrieve context
	// 컨텍스트 검색
	retrieved := GetContext(reqWithCtx)

	if retrieved == nil {
		t.Fatal("GetContext returned nil")
	}

	if retrieved.Param("id") != "123" {
		t.Errorf("Retrieved context has wrong parameter value")
	}
}

// TestGetContextNil tests GetContext with no stored context.
// TestGetContextNil은 저장된 컨텍스트가 없는 GetContext를 테스트합니다.
func TestGetContextNil(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	ctx := GetContext(req)

	if ctx == nil {
		t.Fatal("GetContext should return empty context, not nil")
	}

	if ctx.Request != req {
		t.Error("Empty context should have correct request")
	}
}

// BenchmarkNewContext benchmarks creating a new Context.
// BenchmarkNewContext는 새 Context 생성을 벤치마크합니다.
func BenchmarkNewContext(b *testing.B) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewContext(rec, req)
	}
}

// BenchmarkContextSetGet benchmarks Set and Get operations.
// BenchmarkContextSetGet은 Set 및 Get 작업을 벤치마크합니다.
func BenchmarkContextSetGet(b *testing.B) {
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	ctx := NewContext(rec, req)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.Set("key", "value")
		_, _ = ctx.Get("key")
	}
}

// BenchmarkContextParam benchmarks parameter retrieval.
// BenchmarkContextParam은 매개변수 검색을 벤치마크합니다.
func BenchmarkContextParam(b *testing.B) {
	req := httptest.NewRequest("GET", "/users/123", nil)
	rec := httptest.NewRecorder()
	ctx := NewContext(rec, req)
	ctx.setParams(map[string]string{"id": "123"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ctx.Param("id")
	}
}
