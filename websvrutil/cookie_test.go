package websvrutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestContextCookie tests Cookie method / Cookie 메서드 테스트
func TestContextCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "test_cookie",
		Value: "test_value",
	})

	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	// Get existing cookie / 기존 쿠키 가져오기
	cookie, err := ctx.Cookie("test_cookie")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cookie.Name != "test_cookie" {
		t.Errorf("Expected cookie name 'test_cookie', got '%s'", cookie.Name)
	}

	if cookie.Value != "test_value" {
		t.Errorf("Expected cookie value 'test_value', got '%s'", cookie.Value)
	}

	// Get non-existent cookie / 존재하지 않는 쿠키 가져오기
	_, err = ctx.Cookie("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent cookie")
	}
}

// TestContextCookieValue tests CookieValue method / CookieValue 메서드 테스트
func TestContextCookieValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "user",
		Value: "john",
	})

	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	// Get existing cookie value / 기존 쿠키 값 가져오기
	value := ctx.CookieValue("user")
	if value != "john" {
		t.Errorf("Expected 'john', got '%s'", value)
	}

	// Get non-existent cookie value / 존재하지 않는 쿠키 값 가져오기
	value = ctx.CookieValue("nonexistent")
	if value != "" {
		t.Errorf("Expected empty string, got '%s'", value)
	}
}

// TestContextSetCookieExisting tests existing SetCookie method / 기존 SetCookie 메서드 테스트
func TestContextSetCookieExisting(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: "abc123",
		Path:  "/",
	}
	ctx.SetCookie(cookie)

	// Check cookie is set / 쿠키 설정 확인
	cookies := w.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatal("Expected cookie to be set")
	}

	setCookie := cookies[0]
	if setCookie.Name != "session" {
		t.Errorf("Expected cookie name 'session', got '%s'", setCookie.Name)
	}

	if setCookie.Value != "abc123" {
		t.Errorf("Expected cookie value 'abc123', got '%s'", setCookie.Value)
	}

	if setCookie.Path != "/" {
		t.Errorf("Expected cookie path '/', got '%s'", setCookie.Path)
	}
}

// TestContextSetCookieAdvanced tests SetCookieAdvanced method / SetCookieAdvanced 메서드 테스트
func TestContextSetCookieAdvanced(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.SetCookieAdvanced(CookieOptions{
		Name:     "advanced",
		Value:    "value123",
		Path:     "/api",
		Domain:   "example.com",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	// Check cookie is set with advanced options / 고급 옵션으로 쿠키 설정 확인
	cookies := w.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatal("Expected cookie to be set")
	}

	cookie := cookies[0]
	if cookie.Name != "advanced" {
		t.Errorf("Expected cookie name 'advanced', got '%s'", cookie.Name)
	}

	if cookie.Value != "value123" {
		t.Errorf("Expected cookie value 'value123', got '%s'", cookie.Value)
	}

	if cookie.Path != "/api" {
		t.Errorf("Expected cookie path '/api', got '%s'", cookie.Path)
	}

	if cookie.Domain != "example.com" {
		t.Errorf("Expected cookie domain 'example.com', got '%s'", cookie.Domain)
	}

	if cookie.MaxAge != 3600 {
		t.Errorf("Expected cookie MaxAge 3600, got %d", cookie.MaxAge)
	}

	if !cookie.Secure {
		t.Error("Expected cookie to be secure")
	}

	if !cookie.HttpOnly {
		t.Error("Expected cookie to be HttpOnly")
	}

	if cookie.SameSite != http.SameSiteStrictMode {
		t.Errorf("Expected SameSite Strict, got %v", cookie.SameSite)
	}
}

// TestContextSetCookieAdvancedDefaultPath tests default path / 기본 경로 테스트
func TestContextSetCookieAdvancedDefaultPath(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	// Set cookie without path / 경로 없이 쿠키 설정
	ctx.SetCookieAdvanced(CookieOptions{
		Name:  "test",
		Value: "value",
		// Path is empty, should default to "/" / Path가 비어 있으면 기본값 "/"
	})

	cookies := w.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatal("Expected cookie to be set")
	}

	cookie := cookies[0]
	if cookie.Path != "/" {
		t.Errorf("Expected default path '/', got '%s'", cookie.Path)
	}
}

// TestContextDeleteCookieExisting tests existing DeleteCookie method / 기존 DeleteCookie 메서드 테스트
func TestContextDeleteCookieExisting(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.DeleteCookie("session", "/")

	// Check cookie is deleted (MaxAge = -1) / 쿠키 삭제 확인 (MaxAge = -1)
	cookies := w.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatal("Expected cookie to be set for deletion")
	}

	cookie := cookies[0]
	if cookie.Name != "session" {
		t.Errorf("Expected cookie name 'session', got '%s'", cookie.Name)
	}

	if cookie.MaxAge != -1 {
		t.Errorf("Expected MaxAge -1, got %d", cookie.MaxAge)
	}

	if cookie.Value != "" {
		t.Errorf("Expected empty value, got '%s'", cookie.Value)
	}
}

// TestContextMultipleCookies tests setting multiple cookies / 여러 쿠키 설정 테스트
func TestContextMultipleCookies(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	ctx.SetCookie(&http.Cookie{Name: "cookie1", Value: "value1", Path: "/"})
	ctx.SetCookie(&http.Cookie{Name: "cookie2", Value: "value2", Path: "/"})
	ctx.SetCookie(&http.Cookie{Name: "cookie3", Value: "value3", Path: "/"})

	cookies := w.Result().Cookies()
	if len(cookies) != 3 {
		t.Fatalf("Expected 3 cookies, got %d", len(cookies))
	}

	// Check all cookies are set / 모든 쿠키 설정 확인
	cookieMap := make(map[string]string)
	for _, cookie := range cookies {
		cookieMap[cookie.Name] = cookie.Value
	}

	if cookieMap["cookie1"] != "value1" {
		t.Errorf("Expected cookie1='value1', got '%s'", cookieMap["cookie1"])
	}

	if cookieMap["cookie2"] != "value2" {
		t.Errorf("Expected cookie2='value2', got '%s'", cookieMap["cookie2"])
	}

	if cookieMap["cookie3"] != "value3" {
		t.Errorf("Expected cookie3='value3', got '%s'", cookieMap["cookie3"])
	}
}

// TestCookieOptions tests CookieOptions struct / CookieOptions 구조체 테스트
func TestCookieOptions(t *testing.T) {
	opts := CookieOptions{
		Name:     "test",
		Value:    "value",
		Path:     "/api",
		Domain:   "example.com",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	if opts.Name != "test" {
		t.Errorf("Expected Name 'test', got '%s'", opts.Name)
	}

	if opts.Value != "value" {
		t.Errorf("Expected Value 'value', got '%s'", opts.Value)
	}

	if opts.Path != "/api" {
		t.Errorf("Expected Path '/api', got '%s'", opts.Path)
	}

	if opts.Domain != "example.com" {
		t.Errorf("Expected Domain 'example.com', got '%s'", opts.Domain)
	}

	if opts.MaxAge != 3600 {
		t.Errorf("Expected MaxAge 3600, got %d", opts.MaxAge)
	}

	if !opts.Secure {
		t.Error("Expected Secure to be true")
	}

	if !opts.HttpOnly {
		t.Error("Expected HttpOnly to be true")
	}

	if opts.SameSite != http.SameSiteLaxMode {
		t.Errorf("Expected SameSite Lax, got %v", opts.SameSite)
	}
}

// BenchmarkContextSetCookieExisting benchmarks existing SetCookie method / 기존 SetCookie 메서드 벤치마크
func BenchmarkContextSetCookieExisting(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	cookie := &http.Cookie{
		Name:  "bench",
		Value: "value",
		Path:  "/",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.SetCookie(cookie)
	}
}

// BenchmarkContextCookieValue benchmarks CookieValue method / CookieValue 메서드 벤치마크
func BenchmarkContextCookieValue(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "bench",
		Value: "value",
	})

	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.CookieValue("bench")
	}
}

// BenchmarkContextSetCookieAdvanced benchmarks SetCookieAdvanced method / SetCookieAdvanced 메서드 벤치마크
func BenchmarkContextSetCookieAdvanced(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx := &Context{
		Request:        req,
		ResponseWriter: w,
	}

	opts := CookieOptions{
		Name:     "bench",
		Value:    "value",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.SetCookieAdvanced(opts)
	}
}
