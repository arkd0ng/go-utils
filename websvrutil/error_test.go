package websvrutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAbortWithStatus tests AbortWithStatus method / AbortWithStatus 메서드 테스트
func TestAbortWithStatus(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.AbortWithStatus(http.StatusUnauthorized)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

// TestAbortWithError tests AbortWithError method / AbortWithError 메서드 테스트
func TestAbortWithError(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.AbortWithError(http.StatusBadRequest, "Invalid input")

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	body := w.Body.String()
	if body == "" {
		t.Error("Expected error message in body")
	}
}

// TestAbortWithJSON tests AbortWithJSON method / AbortWithJSON 메서드 테스트
func TestAbortWithJSON(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	errorData := map[string]string{"error": "Invalid input"}
	ctx.AbortWithJSON(http.StatusBadRequest, errorData)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if response["error"] != "Invalid input" {
		t.Errorf("Expected error 'Invalid input', got '%s'", response["error"])
	}
}

// TestErrorJSON tests ErrorJSON method / ErrorJSON 메서드 테스트
func TestErrorJSON(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.ErrorJSON(http.StatusNotFound, "User not found")

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if response["error"] != "User not found" {
		t.Errorf("Expected error 'User not found', got '%v'", response["error"])
	}

	if response["success"] != false {
		t.Error("Expected success to be false")
	}

	if int(response["status"].(float64)) != http.StatusNotFound {
		t.Errorf("Expected status %d, got %v", http.StatusNotFound, response["status"])
	}
}

// TestSuccessJSON tests SuccessJSON method / SuccessJSON 메서드 테스트
func TestSuccessJSON(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	data := map[string]string{"name": "John"}
	ctx.SuccessJSON(http.StatusOK, "Operation completed", data)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if response["message"] != "Operation completed" {
		t.Errorf("Expected message 'Operation completed', got '%v'", response["message"])
	}

	if response["success"] != true {
		t.Error("Expected success to be true")
	}

	if int(response["status"].(float64)) != http.StatusOK {
		t.Errorf("Expected status %d, got %v", http.StatusOK, response["status"])
	}

	responseData := response["data"].(map[string]interface{})
	if responseData["name"] != "John" {
		t.Errorf("Expected name 'John', got '%v'", responseData["name"])
	}
}

// TestNotFound tests NotFound method / NotFound 메서드 테스트
func TestNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.NotFound()

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

// TestUnauthorized tests Unauthorized method / Unauthorized 메서드 테스트
func TestUnauthorized(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.Unauthorized()

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

// TestForbidden tests Forbidden method / Forbidden 메서드 테스트
func TestForbidden(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.Forbidden()

	if w.Code != http.StatusForbidden {
		t.Errorf("Expected status %d, got %d", http.StatusForbidden, w.Code)
	}
}

// TestBadRequest tests BadRequest method / BadRequest 메서드 테스트
func TestBadRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.BadRequest()

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

// TestInternalServerError tests InternalServerError method / InternalServerError 메서드 테스트
func TestInternalServerError(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.InternalServerError()

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

// BenchmarkErrorJSON benchmarks ErrorJSON method / ErrorJSON 메서드 벤치마크
func BenchmarkErrorJSON(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w = httptest.NewRecorder()
		ctx.ResponseWriter = w
		ctx.ErrorJSON(http.StatusNotFound, "Not found")
	}
}

// BenchmarkSuccessJSON benchmarks SuccessJSON method / SuccessJSON 메서드 벤치마크
func BenchmarkSuccessJSON(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	data := map[string]string{"name": "John"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w = httptest.NewRecorder()
		ctx.ResponseWriter = w
		ctx.SuccessJSON(http.StatusOK, "Success", data)
	}
}
