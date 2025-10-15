package websvrutil

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
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
