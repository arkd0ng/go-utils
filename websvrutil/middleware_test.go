package websvrutil

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

// TestRecovery tests the Recovery middleware.
// TestRecovery는 Recovery 미들웨어를 테스트합니다.
func TestRecovery(t *testing.T) {
	// Capture log output
	// 로그 출력 캡처
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	middleware := Recovery()

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Status code = %d, want %d", rec.Code, http.StatusInternalServerError)
	}

	logOutput := buf.String()
	if !strings.Contains(logOutput, "PANIC") {
		t.Error("Log should contain PANIC message")
	}
	if !strings.Contains(logOutput, "test panic") {
		t.Error("Log should contain panic message")
	}
}

// TestRecoveryNoPanic tests Recovery middleware with no panic.
// TestRecoveryNoPanic은 패닉이 없는 Recovery 미들웨어를 테스트합니다.
func TestRecoveryNoPanic(t *testing.T) {
	middleware := Recovery()

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d, want %d", rec.Code, http.StatusOK)
	}

	if rec.Body.String() != "OK" {
		t.Errorf("Body = %s, want OK", rec.Body.String())
	}
}

// TestRecoveryWithConfig tests Recovery middleware with custom config.
// TestRecoveryWithConfig는 커스텀 설정으로 Recovery 미들웨어를 테스트합니다.
func TestRecoveryWithConfig(t *testing.T) {
	var capturedErr interface{}
	var capturedStack []byte

	config := RecoveryConfig{
		PrintStack: true,
		LogFunc: func(err interface{}, stack []byte) {
			capturedErr = err
			capturedStack = stack
		},
	}

	middleware := RecoveryWithConfig(config)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("custom panic")
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if capturedErr == nil {
		t.Error("LogFunc should have been called")
	}

	if capturedErr != "custom panic" {
		t.Errorf("Captured error = %v, want 'custom panic'", capturedErr)
	}

	if len(capturedStack) == 0 {
		t.Error("Stack trace should not be empty")
	}
}

// TestLogger tests the Logger middleware.
// TestLogger는 Logger 미들웨어를 테스트합니다.
func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	middleware := Logger()

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	logOutput := buf.String()
	if !strings.Contains(logOutput, "GET") {
		t.Error("Log should contain method")
	}
	if !strings.Contains(logOutput, "/test") {
		t.Error("Log should contain path")
	}
	if !strings.Contains(logOutput, "200") {
		t.Error("Log should contain status code")
	}
}

// TestLoggerWithConfig tests Logger middleware with custom config.
// TestLoggerWithConfig는 커스텀 설정으로 Logger 미들웨어를 테스트합니다.
func TestLoggerWithConfig(t *testing.T) {
	var loggedMethod, loggedPath string
	var loggedStatus int
	var loggedDuration time.Duration

	config := LoggerConfig{
		LogFunc: func(method, path string, status int, duration time.Duration) {
			loggedMethod = method
			loggedPath = path
			loggedStatus = status
			loggedDuration = duration
		},
	}

	middleware := LoggerWithConfig(config)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.WriteHeader(http.StatusCreated)
	}))

	req := httptest.NewRequest("POST", "/users", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if loggedMethod != "POST" {
		t.Errorf("Logged method = %s, want POST", loggedMethod)
	}

	if loggedPath != "/users" {
		t.Errorf("Logged path = %s, want /users", loggedPath)
	}

	if loggedStatus != http.StatusCreated {
		t.Errorf("Logged status = %d, want %d", loggedStatus, http.StatusCreated)
	}

	if loggedDuration < 10*time.Millisecond {
		t.Errorf("Logged duration = %v, should be >= 10ms", loggedDuration)
	}
}

// TestCORS tests the CORS middleware.
// TestCORS는 CORS 미들웨어를 테스트합니다.
func TestCORS(t *testing.T) {
	middleware := CORS()

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://example.com")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	origin := rec.Header().Get("Access-Control-Allow-Origin")
	if origin == "" {
		t.Error("Access-Control-Allow-Origin header should be set")
	}

	methods := rec.Header().Get("Access-Control-Allow-Methods")
	if methods == "" {
		t.Error("Access-Control-Allow-Methods header should be set")
	}

	headers := rec.Header().Get("Access-Control-Allow-Headers")
	if headers == "" {
		t.Error("Access-Control-Allow-Headers header should be set")
	}
}

// TestCORSPreflight tests CORS preflight requests.
// TestCORSPreflight는 CORS 프리플라이트 요청을 테스트합니다.
func TestCORSPreflight(t *testing.T) {
	middleware := CORS()

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called for OPTIONS request")
	}))

	req := httptest.NewRequest("OPTIONS", "/test", nil)
	req.Header.Set("Origin", "https://example.com")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("Status code = %d, want %d", rec.Code, http.StatusNoContent)
	}

	origin := rec.Header().Get("Access-Control-Allow-Origin")
	if origin == "" {
		t.Error("Access-Control-Allow-Origin header should be set")
	}
}

// TestCORSWithConfig tests CORS middleware with custom config.
// TestCORSWithConfig는 커스텀 설정으로 CORS 미들웨어를 테스트합니다.
func TestCORSWithConfig(t *testing.T) {
	config := CORSConfig{
		AllowOrigins:     []string{"https://example.com", "https://api.example.com"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           3600 * time.Second,
	}

	middleware := CORSWithConfig(config)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://example.com")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	origin := rec.Header().Get("Access-Control-Allow-Origin")
	if origin != "https://example.com" {
		t.Errorf("Origin = %s, want https://example.com", origin)
	}

	credentials := rec.Header().Get("Access-Control-Allow-Credentials")
	if credentials != "true" {
		t.Errorf("Credentials = %s, want true", credentials)
	}

	maxAge := rec.Header().Get("Access-Control-Max-Age")
	if maxAge != "3600" {
		t.Errorf("MaxAge = %s, want 3600", maxAge)
	}
}

// TestCORSNotAllowedOrigin tests CORS with non-allowed origin.
// TestCORSNotAllowedOrigin은 허용되지 않은 오리진으로 CORS를 테스트합니다.
func TestCORSNotAllowedOrigin(t *testing.T) {
	config := CORSConfig{
		AllowOrigins: []string{"https://example.com"},
	}

	middleware := CORSWithConfig(config)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://evil.com")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	origin := rec.Header().Get("Access-Control-Allow-Origin")
	if origin == "https://evil.com" {
		t.Error("Origin should not be set for non-allowed origin")
	}
}

// TestResponseWriter tests responseWriter wrapper.
// TestResponseWriter는 responseWriter 래퍼를 테스트합니다.
func TestResponseWriter(t *testing.T) {
	rec := httptest.NewRecorder()
	rw := &responseWriter{ResponseWriter: rec, statusCode: http.StatusOK}

	rw.WriteHeader(http.StatusCreated)

	if rw.statusCode != http.StatusCreated {
		t.Errorf("Status code = %d, want %d", rw.statusCode, http.StatusCreated)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Underlying status code = %d, want %d", rec.Code, http.StatusCreated)
	}
}

// TestIsOriginAllowed tests the isOriginAllowed helper.
// TestIsOriginAllowed는 isOriginAllowed 헬퍼를 테스트합니다.
func TestIsOriginAllowed(t *testing.T) {
	tests := []struct {
		origin  string
		allowed []string
		want    bool
	}{
		{"https://example.com", []string{"https://example.com"}, true},
		{"https://example.com", []string{"*"}, true},
		{"https://example.com", []string{"https://other.com"}, false},
		{"https://example.com", []string{"https://example.com", "https://other.com"}, true},
	}

	for _, tt := range tests {
		got := isOriginAllowed(tt.origin, tt.allowed)
		if got != tt.want {
			t.Errorf("isOriginAllowed(%s, %v) = %v, want %v", tt.origin, tt.allowed, got, tt.want)
		}
	}
}

// TestJoinStrings tests the joinStrings helper.
// TestJoinStrings는 joinStrings 헬퍼를 테스트합니다.
func TestJoinStrings(t *testing.T) {
	tests := []struct {
		strs []string
		sep  string
		want string
	}{
		{[]string{"a", "b", "c"}, ", ", "a, b, c"},
		{[]string{"GET", "POST"}, ", ", "GET, POST"},
		{[]string{"one"}, ", ", "one"},
		{[]string{}, ", ", ""},
	}

	for _, tt := range tests {
		got := joinStrings(tt.strs, tt.sep)
		if got != tt.want {
			t.Errorf("joinStrings(%v, %s) = %s, want %s", tt.strs, tt.sep, got, tt.want)
		}
	}
}

// BenchmarkRecovery benchmarks the Recovery middleware.
// BenchmarkRecovery는 Recovery 미들웨어를 벤치마크합니다.
func BenchmarkRecovery(b *testing.B) {
	middleware := Recovery()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkLogger benchmarks the Logger middleware.
// BenchmarkLogger는 Logger 미들웨어를 벤치마크합니다.
func BenchmarkLogger(b *testing.B) {
	log.SetOutput(&bytes.Buffer{})
	defer log.SetOutput(nil)

	middleware := Logger()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkCORS benchmarks the CORS middleware.
// BenchmarkCORS는 CORS 미들웨어를 벤치마크합니다.
func BenchmarkCORS(b *testing.B) {
	middleware := CORS()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://example.com")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// TestRequestID tests the RequestID middleware.
// TestRequestID는 RequestID 미들웨어를 테스트합니다.
func TestRequestID(t *testing.T) {
	middleware := RequestID()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request ID in context
		// 컨텍스트에서 요청 ID 확인
		requestID := r.Context().Value("request_id")
		if requestID == nil {
			t.Error("Request ID not found in context")
		}
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	// Check response header
	// 응답 헤더 확인
	requestID := rec.Header().Get("X-Request-ID")
	if requestID == "" {
		t.Error("X-Request-ID header not set")
	}
	if len(requestID) != 32 { // 16 bytes hex = 32 characters
		t.Errorf("Expected request ID length 32, got %d", len(requestID))
	}
}

// TestRequestIDWithExistingID tests that existing request ID is preserved.
// TestRequestIDWithExistingID는 기존 요청 ID가 보존되는지 테스트합니다.
func TestRequestIDWithExistingID(t *testing.T) {
	middleware := RequestID()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("X-Request-ID", "existing-id-12345")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	// Check that existing ID is preserved
	// 기존 ID가 보존되는지 확인
	requestID := rec.Header().Get("X-Request-ID")
	if requestID != "existing-id-12345" {
		t.Errorf("Expected request ID 'existing-id-12345', got '%s'", requestID)
	}
}

// TestRequestIDWithConfig tests custom configuration.
// TestRequestIDWithConfig는 커스텀 설정을 테스트합니다.
func TestRequestIDWithConfig(t *testing.T) {
	called := false
	middleware := RequestIDWithConfig(RequestIDConfig{
		Header: "X-Custom-Request-ID",
		Generator: func() string {
			called = true
			return "custom-id-67890"
		},
	})

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if !called {
		t.Error("Custom generator not called")
	}

	requestID := rec.Header().Get("X-Custom-Request-ID")
	if requestID != "custom-id-67890" {
		t.Errorf("Expected request ID 'custom-id-67890', got '%s'", requestID)
	}
}

// TestTimeout tests the Timeout middleware.
// TestTimeout는 Timeout 미들웨어를 테스트합니다.
func TestTimeout(t *testing.T) {
	middleware := Timeout(100 * time.Millisecond)
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Fast response, should not timeout
		// 빠른 응답, 타임아웃 발생하지 않음
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}

// TestTimeoutWithConfig tests custom timeout configuration.
// TestTimeoutWithConfig는 커스텀 타임아웃 설정을 테스트합니다.
func TestTimeoutWithConfig(t *testing.T) {
	middleware := TimeoutWithConfig(TimeoutConfig{
		Timeout: 50 * time.Millisecond,
		Message: "Request timed out",
	})

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}

// TestBasicAuth tests the BasicAuth middleware with valid credentials.
// TestBasicAuth는 유효한 자격 증명으로 BasicAuth 미들웨어를 테스트합니다.
func TestBasicAuth(t *testing.T) {
	middleware := BasicAuth("admin", "password")
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check username in context
		// 컨텍스트에서 사용자 이름 확인
		username := r.Context().Value("auth_username")
		if username != "admin" {
			t.Errorf("Expected username 'admin', got '%v'", username)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authenticated"))
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.SetBasicAuth("admin", "password")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
	if rec.Body.String() != "Authenticated" {
		t.Errorf("Expected body 'Authenticated', got '%s'", rec.Body.String())
	}
}

// TestBasicAuthUnauthorized tests BasicAuth with invalid credentials.
// TestBasicAuthUnauthorized는 잘못된 자격 증명으로 BasicAuth를 테스트합니다.
func TestBasicAuthUnauthorized(t *testing.T) {
	middleware := BasicAuth("admin", "password")
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called for unauthorized request")
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.SetBasicAuth("admin", "wrongpassword")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec.Code)
	}

	wwwAuth := rec.Header().Get("WWW-Authenticate")
	if wwwAuth != `Basic realm="Restricted"` {
		t.Errorf("Expected WWW-Authenticate header, got '%s'", wwwAuth)
	}
}

// TestBasicAuthNoCredentials tests BasicAuth without credentials.
// TestBasicAuthNoCredentials는 자격 증명 없이 BasicAuth를 테스트합니다.
func TestBasicAuthNoCredentials(t *testing.T) {
	middleware := BasicAuth("admin", "password")
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called for request without credentials")
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec.Code)
	}
}

// TestBasicAuthWithConfig tests custom BasicAuth configuration.
// TestBasicAuthWithConfig는 커스텀 BasicAuth 설정을 테스트합니다.
func TestBasicAuthWithConfig(t *testing.T) {
	middleware := BasicAuthWithConfig(BasicAuthConfig{
		Validator: func(username, password string) bool {
			return username == "user1" && password == "secret"
		},
		Realm: "Admin Area",
	})

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Test valid credentials
	// 유효한 자격 증명 테스트
	req := httptest.NewRequest("GET", "/test", nil)
	req.SetBasicAuth("user1", "secret")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Test invalid credentials
	// 잘못된 자격 증명 테스트
	req2 := httptest.NewRequest("GET", "/test", nil)
	req2.SetBasicAuth("user1", "wrong")
	rec2 := httptest.NewRecorder()

	handler.ServeHTTP(rec2, req2)

	if rec2.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec2.Code)
	}

	wwwAuth := rec2.Header().Get("WWW-Authenticate")
	if wwwAuth != `Basic realm="Admin Area"` {
		t.Errorf("Expected custom realm in WWW-Authenticate, got '%s'", wwwAuth)
	}
}

// BenchmarkRequestID benchmarks the RequestID middleware.
// BenchmarkRequestID는 RequestID 미들웨어를 벤치마크합니다.
func BenchmarkRequestID(b *testing.B) {
	middleware := RequestID()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkTimeout benchmarks the Timeout middleware.
// BenchmarkTimeout는 Timeout 미들웨어를 벤치마크합니다.
func BenchmarkTimeout(b *testing.B) {
	middleware := Timeout(5 * time.Second)
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkBasicAuth benchmarks the BasicAuth middleware.
// BenchmarkBasicAuth는 BasicAuth 미들웨어를 벤치마크합니다.
func BenchmarkBasicAuth(b *testing.B) {
	middleware := BasicAuth("admin", "password")
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.SetBasicAuth("admin", "password")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// TestRateLimiter tests the RateLimiter middleware.
// TestRateLimiter는 RateLimiter 미들웨어를 테스트합니다.
func TestRateLimiter(t *testing.T) {
	middleware := RateLimiter(5, time.Minute)
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Make 5 requests (should all succeed)
	// 5개 요청 (모두 성공해야 함)
	for i := 0; i < 5; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "127.0.0.1:12345"
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Request %d: expected status 200, got %d", i+1, rec.Code)
		}

		limit := rec.Header().Get("X-RateLimit-Limit")
		if limit != "5" {
			t.Errorf("Expected X-RateLimit-Limit '5', got '%s'", limit)
		}
	}

	// 6th request should be rate limited
	// 6번째 요청은 rate limited 되어야 함
	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "127.0.0.1:12345"
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status 429, got %d", rec.Code)
	}

	remaining := rec.Header().Get("X-RateLimit-Remaining")
	if remaining != "0" {
		t.Errorf("Expected X-RateLimit-Remaining '0', got '%s'", remaining)
	}
}

// TestRateLimiterWithConfig tests custom RateLimiter configuration.
// TestRateLimiterWithConfig는 커스텀 RateLimiter 설정을 테스트합니다.
func TestRateLimiterWithConfig(t *testing.T) {
	middleware := RateLimiterWithConfig(RateLimiterConfig{
		Requests: 2,
		Window:   time.Second,
		KeyFunc: func(r *http.Request) string {
			return r.Header.Get("X-API-Key")
		},
	})

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Test with API key
	// API 키로 테스트
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("X-API-Key", "test-key")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Request %d: expected status 200, got %d", i+1, rec.Code)
		}
	}

	// 3rd request should be rate limited
	// 3번째 요청은 rate limited 되어야 함
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("X-API-Key", "test-key")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status 429, got %d", rec.Code)
	}
}

// TestCompression tests the Compression middleware.
// TestCompression는 Compression 미들웨어를 테스트합니다.
func TestCompression(t *testing.T) {
	middleware := Compression()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World! This is a test response that should be compressed."))
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	contentEncoding := rec.Header().Get("Content-Encoding")
	if contentEncoding != "gzip" {
		t.Errorf("Expected Content-Encoding 'gzip', got '%s'", contentEncoding)
	}

	// Response should be gzip compressed
	// 응답이 gzip으로 압축되어야 함
	if len(rec.Body.Bytes()) == 0 {
		t.Error("Expected compressed response, got empty body")
	}
}

// TestCompressionWithoutGzip tests that compression is skipped when client doesn't support gzip.
// TestCompressionWithoutGzip는 클라이언트가 gzip을 지원하지 않을 때 압축을 건너뛰는지 테스트합니다.
func TestCompressionWithoutGzip(t *testing.T) {
	middleware := Compression()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	// No Accept-Encoding header
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	contentEncoding := rec.Header().Get("Content-Encoding")
	if contentEncoding != "" {
		t.Errorf("Expected no Content-Encoding, got '%s'", contentEncoding)
	}

	body := rec.Body.String()
	if body != "Hello, World!" {
		t.Errorf("Expected uncompressed body, got '%s'", body)
	}
}

// TestCompressionWithConfig tests custom Compression configuration.
// TestCompressionWithConfig는 커스텀 Compression 설정을 테스트합니다.
func TestCompressionWithConfig(t *testing.T) {
	middleware := CompressionWithConfig(CompressionConfig{
		Level:     gzip.BestCompression,
		MinLength: 1024,
	})

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test"))
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}

// TestSecureHeaders tests the SecureHeaders middleware.
// TestSecureHeaders는 SecureHeaders 미들웨어를 테스트합니다.
func TestSecureHeaders(t *testing.T) {
	middleware := SecureHeaders()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	// Check security headers
	// 보안 헤더 확인
	headers := []struct {
		name     string
		expected string
	}{
		{"X-Frame-Options", "SAMEORIGIN"},
		{"X-Content-Type-Options", "nosniff"},
		{"X-XSS-Protection", "1; mode=block"},
		{"Referrer-Policy", "strict-origin-when-cross-origin"},
	}

	for _, h := range headers {
		value := rec.Header().Get(h.name)
		if value != h.expected {
			t.Errorf("Expected %s '%s', got '%s'", h.name, h.expected, value)
		}
	}
}

// TestSecureHeadersWithConfig tests custom SecureHeaders configuration.
// TestSecureHeadersWithConfig는 커스텀 SecureHeaders 설정을 테스트합니다.
func TestSecureHeadersWithConfig(t *testing.T) {
	middleware := SecureHeadersWithConfig(SecureHeadersConfig{
		XFrameOptions:         "DENY",
		ContentSecurityPolicy: "default-src 'self'",
	})

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	xFrameOptions := rec.Header().Get("X-Frame-Options")
	if xFrameOptions != "DENY" {
		t.Errorf("Expected X-Frame-Options 'DENY', got '%s'", xFrameOptions)
	}

	csp := rec.Header().Get("Content-Security-Policy")
	if csp != "default-src 'self'" {
		t.Errorf("Expected CSP 'default-src 'self'', got '%s'", csp)
	}
}

// BenchmarkRateLimiter benchmarks the RateLimiter middleware.
// BenchmarkRateLimiter는 RateLimiter 미들웨어를 벤치마크합니다.
func BenchmarkRateLimiter(b *testing.B) {
	middleware := RateLimiter(1000, time.Minute)
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "127.0.0.1:12345"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkCompression benchmarks the Compression middleware.
// BenchmarkCompression는 Compression 미들웨어를 벤치마크합니다.
func BenchmarkCompression(b *testing.B) {
	middleware := Compression()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World! This is a test response."))
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Accept-Encoding", "gzip")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkSecureHeaders benchmarks the SecureHeaders middleware.
// BenchmarkSecureHeaders는 SecureHeaders 미들웨어를 벤치마크합니다.
func BenchmarkSecureHeaders(b *testing.B) {
	middleware := SecureHeaders()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// TestBodyLimit tests the BodyLimit middleware
// TestBodyLimit는 BodyLimit 미들웨어를 테스트합니다
func TestBodyLimit(t *testing.T) {
	middleware := BodyLimit(100) // 100 bytes limit
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}))

	// Test with small body (should succeed)
	t.Run("SmallBody", func(t *testing.T) {
		smallBody := strings.Repeat("a", 50)
		req := httptest.NewRequest("POST", "/test", strings.NewReader(smallBody))
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rec.Code)
		}
	})

	// Test with large body (should fail)
	t.Run("LargeBody", func(t *testing.T) {
		largeBody := strings.Repeat("a", 200)
		req := httptest.NewRequest("POST", "/test", strings.NewReader(largeBody))
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		// The body should be truncated or return error
		// 본문이 잘리거나 에러가 반환되어야 함
		if rec.Code == http.StatusOK && len(rec.Body.String()) > 100 {
			t.Errorf("Body should be limited to 100 bytes")
		}
	})
}

// TestBodyLimitWithConfig tests the BodyLimit middleware with custom config
// TestBodyLimitWithConfig는 사용자 정의 설정으로 BodyLimit 미들웨어를 테스트합니다
func TestBodyLimitWithConfig(t *testing.T) {
	middleware := BodyLimitWithConfig(BodyLimitConfig{
		MaxBytes: 1024,
	})
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("POST", "/test", strings.NewReader("test"))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}

// TestStatic tests the Static middleware
// TestStatic는 Static 미들웨어를 테스트합니다
func TestStatic(t *testing.T) {
	// Create temporary directory and file for testing
	// 테스트를 위한 임시 디렉토리와 파일 생성
	tmpDir := t.TempDir()
	testFile := tmpDir + "/test.txt"
	os.WriteFile(testFile, []byte("Hello, Static!"), 0644)

	middleware := Static(tmpDir)
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}))

	// Test existing file
	t.Run("ExistingFile", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test.txt", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "Hello, Static!") {
			t.Errorf("Expected 'Hello, Static!' in response, got %s", rec.Body.String())
		}
	})

	// Test non-existing file (should fall through to next handler)
	t.Run("NonExistingFile", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/nonexistent.txt", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", rec.Code)
		}
	})
}

// TestStaticWithConfig tests the Static middleware with custom config
// TestStaticWithConfig는 사용자 정의 설정으로 Static 미들웨어를 테스트합니다
func TestStaticWithConfig(t *testing.T) {
	tmpDir := t.TempDir()
	indexFile := tmpDir + "/index.html"
	os.WriteFile(indexFile, []byte("<h1>Index</h1>"), 0644)

	middleware := StaticWithConfig(StaticConfig{
		Root:  tmpDir,
		Index: "index.html",
	})
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}

// TestRedirect tests the Redirect middleware
// TestRedirect는 Redirect 미들웨어를 테스트합니다
func TestRedirect(t *testing.T) {
	middleware := Redirect("https://example.com")
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status 301, got %d", rec.Code)
	}
	location := rec.Header().Get("Location")
	if location != "https://example.com" {
		t.Errorf("Expected Location 'https://example.com', got '%s'", location)
	}
}

// TestRedirectWithConfig tests the Redirect middleware with custom config
// TestRedirectWithConfig는 사용자 정의 설정으로 Redirect 미들웨어를 테스트합니다
func TestRedirectWithConfig(t *testing.T) {
	middleware := RedirectWithConfig(RedirectConfig{
		Code: http.StatusFound,
		To:   "https://new-site.com",
	})
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusFound {
		t.Errorf("Expected status 302, got %d", rec.Code)
	}
}

// TestHTTPSRedirect tests the HTTPSRedirect middleware
// TestHTTPSRedirect는 HTTPSRedirect 미들웨어를 테스트합니다
func TestHTTPSRedirect(t *testing.T) {
	middleware := HTTPSRedirect()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Test HTTP request (should redirect)
	t.Run("HTTPRequest", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://example.com/test", nil)
		req.Host = "example.com"
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusMovedPermanently {
			t.Errorf("Expected status 301, got %d", rec.Code)
		}
		location := rec.Header().Get("Location")
		if !strings.HasPrefix(location, "https://") {
			t.Errorf("Expected HTTPS redirect, got %s", location)
		}
	})

	// Test HTTPS request (should not redirect)
	t.Run("HTTPSRequest", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example.com/test", nil)
		req.Header.Set("X-Forwarded-Proto", "https")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rec.Code)
		}
	})
}

// TestWWWRedirect tests the WWWRedirect middleware
// TestWWWRedirect는 WWWRedirect 미들웨어를 테스트합니다
func TestWWWRedirect(t *testing.T) {
	// Test adding www prefix
	t.Run("AddWWW", func(t *testing.T) {
		middleware := WWWRedirect(true)
		handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "http://example.com/test", nil)
		req.Host = "example.com"
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusMovedPermanently {
			t.Errorf("Expected status 301, got %d", rec.Code)
		}
		location := rec.Header().Get("Location")
		if !strings.Contains(location, "www.example.com") {
			t.Errorf("Expected www.example.com in location, got %s", location)
		}
	})

	// Test removing www prefix
	t.Run("RemoveWWW", func(t *testing.T) {
		middleware := WWWRedirect(false)
		handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "http://www.example.com/test", nil)
		req.Host = "www.example.com"
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusMovedPermanently {
			t.Errorf("Expected status 301, got %d", rec.Code)
		}
		location := rec.Header().Get("Location")
		if strings.Contains(location, "www.") {
			t.Errorf("Expected no www in location, got %s", location)
		}
	})

	// Test no redirect when already correct
	t.Run("NoRedirect", func(t *testing.T) {
		middleware := WWWRedirect(true)
		handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "http://www.example.com/test", nil)
		req.Host = "www.example.com"
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rec.Code)
		}
	})
}

// BenchmarkBodyLimit benchmarks the BodyLimit middleware
// BenchmarkBodyLimit는 BodyLimit 미들웨어를 벤치마크합니다
func BenchmarkBodyLimit(b *testing.B) {
	middleware := BodyLimit(1024 * 1024) // 1MB
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/test", strings.NewReader("test data"))
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkStatic benchmarks the Static middleware
// BenchmarkStatic는 Static 미들웨어를 벤치마크합니다
func BenchmarkStatic(b *testing.B) {
	tmpDir := b.TempDir()
	middleware := Static(tmpDir)
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/nonexistent.txt", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkRedirect benchmarks the Redirect middleware
// BenchmarkRedirect는 Redirect 미들웨어를 벤치마크합니다
func BenchmarkRedirect(b *testing.B) {
	middleware := Redirect("https://example.com")
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkHTTPSRedirect benchmarks the HTTPSRedirect middleware
// BenchmarkHTTPSRedirect는 HTTPSRedirect 미들웨어를 벤치마크합니다
func BenchmarkHTTPSRedirect(b *testing.B) {
	middleware := HTTPSRedirect()
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "http://example.com/test", nil)
		req.Host = "example.com"
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}

// BenchmarkWWWRedirect benchmarks the WWWRedirect middleware
// BenchmarkWWWRedirect는 WWWRedirect 미들웨어를 벤치마크합니다
func BenchmarkWWWRedirect(b *testing.B) {
	middleware := WWWRedirect(true)
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "http://example.com/test", nil)
		req.Host = "example.com"
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}
