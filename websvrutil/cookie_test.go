package websvrutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCookie tests getting a cookie
// TestCookie는 쿠키 가져오기를 테스트합니다
func TestCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: "abc123"})
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	cookie, err := ctx.Cookie("session_id")
	if err != nil {
		t.Fatalf("Cookie() error = %v", err)
	}

	if cookie.Name != "session_id" {
		t.Errorf("Cookie name = %s, want session_id", cookie.Name)
	}

	if cookie.Value != "abc123" {
		t.Errorf("Cookie value = %s, want abc123", cookie.Value)
	}
}

// TestCookieNotFound tests getting a non-existent cookie
// TestCookieNotFound는 존재하지 않는 쿠키 가져오기를 테스트합니다
func TestCookieNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	_, err := ctx.Cookie("nonexistent")
	if err == nil {
		t.Error("Expected error for non-existent cookie")
	}
}

// TestSetCookie tests setting a cookie
// TestSetCookie는 쿠키 설정을 테스트합니다
func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "xyz789",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
	}

	ctx.SetCookie(cookie)

	// Check Set-Cookie header
	// Set-Cookie 헤더 확인
	setCookie := w.Header().Get("Set-Cookie")
	if setCookie == "" {
		t.Fatal("Set-Cookie header not set")
	}

	if !contains(setCookie, "session_id=xyz789") {
		t.Errorf("Set-Cookie header does not contain session_id=xyz789")
	}
}

// TestDeleteCookie tests deleting a cookie
// TestDeleteCookie는 쿠키 삭제를 테스트합니다
func TestDeleteCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	ctx.DeleteCookie("session_id", "/")

	// Check Set-Cookie header
	// Set-Cookie 헤더 확인
	setCookie := w.Header().Get("Set-Cookie")
	if setCookie == "" {
		t.Fatal("Set-Cookie header not set")
	}

	if !contains(setCookie, "Max-Age=0") && !contains(setCookie, "max-age=0") {
		t.Errorf("Set-Cookie header should have Max-Age=0 for deletion")
	}
}

// TestGetCookie tests the convenience method for getting cookie value
// TestGetCookie는 쿠키 값을 가져오는 편의 메서드를 테스트합니다
func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "user", Value: "john"})
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	value := ctx.GetCookie("user")
	if value != "john" {
		t.Errorf("GetCookie() = %s, want john", value)
	}

	// Test non-existent cookie
	// 존재하지 않는 쿠키 테스트
	value = ctx.GetCookie("nonexistent")
	if value != "" {
		t.Errorf("GetCookie() for non-existent cookie = %s, want empty string", value)
	}
}

// TestAddHeader tests adding a header
// TestAddHeader는 헤더 추가를 테스트합니다
func TestAddHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	ctx.AddHeader("X-Custom", "value1")
	ctx.AddHeader("X-Custom", "value2")

	values := w.Header().Values("X-Custom")
	if len(values) != 2 {
		t.Fatalf("Expected 2 values for X-Custom header, got %d", len(values))
	}

	if values[0] != "value1" || values[1] != "value2" {
		t.Errorf("Header values = %v, want [value1, value2]", values)
	}
}

// TestGetHeader tests getting a request header
// TestGetHeader는 요청 헤더 가져오기를 테스트합니다
func TestGetHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("User-Agent", "Test/1.0")
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	value := ctx.GetHeader("User-Agent")
	if value != "Test/1.0" {
		t.Errorf("GetHeader() = %s, want Test/1.0", value)
	}
}

// TestGetHeaders tests getting multiple header values
// TestGetHeaders는 여러 헤더 값 가져오기를 테스트합니다
func TestGetHeaders(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "text/html")
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	values := ctx.GetHeaders("Accept")
	if len(values) != 2 {
		t.Fatalf("Expected 2 values for Accept header, got %d", len(values))
	}

	if values[0] != "text/html" || values[1] != "application/json" {
		t.Errorf("Header values = %v, want [text/html, application/json]", values)
	}
}

// TestHeaderExists tests checking if a header exists
// TestHeaderExists는 헤더 존재 확인을 테스트합니다
func TestHeaderExists(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer token123")
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	if !ctx.HeaderExists("Authorization") {
		t.Error("HeaderExists() = false, want true")
	}

	if ctx.HeaderExists("X-Nonexistent") {
		t.Error("HeaderExists() = true, want false for non-existent header")
	}
}

// TestContentType tests getting Content-Type header
// TestContentType는 Content-Type 헤더 가져오기를 테스트합니다
func TestContentType(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	contentType := ctx.ContentType()
	if contentType != "application/json" {
		t.Errorf("ContentType() = %s, want application/json", contentType)
	}
}

// TestUserAgent tests getting User-Agent header
// TestUserAgent는 User-Agent 헤더 가져오기를 테스트합니다
func TestUserAgent(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	userAgent := ctx.UserAgent()
	if userAgent != "Mozilla/5.0" {
		t.Errorf("UserAgent() = %s, want Mozilla/5.0", userAgent)
	}
}

// TestReferer tests getting Referer header
// TestReferer는 Referer 헤더 가져오기를 테스트합니다
func TestReferer(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Referer", "https://example.com")
	w := httptest.NewRecorder()

	ctx := NewContext(w, req)

	referer := ctx.Referer()
	if referer != "https://example.com" {
		t.Errorf("Referer() = %s, want https://example.com", referer)
	}
}

// TestClientIP tests getting client IP address
// TestClientIP는 클라이언트 IP 주소 가져오기를 테스트합니다
func TestClientIP(t *testing.T) {
	tests := []struct {
		name           string
		remoteAddr     string
		xForwardedFor  string
		xRealIP        string
		expectedIP     string
	}{
		{
			name:       "from RemoteAddr",
			remoteAddr: "192.168.1.1:12345",
			expectedIP: "192.168.1.1",
		},
		{
			name:          "from X-Real-IP",
			remoteAddr:    "192.168.1.1:12345",
			xRealIP:       "10.0.0.1",
			expectedIP:    "10.0.0.1",
		},
		{
			name:          "from X-Forwarded-For single",
			remoteAddr:    "192.168.1.1:12345",
			xForwardedFor: "10.0.0.1",
			expectedIP:    "10.0.0.1",
		},
		{
			name:          "from X-Forwarded-For multiple",
			remoteAddr:    "192.168.1.1:12345",
			xForwardedFor: "10.0.0.1, 10.0.0.2, 10.0.0.3",
			expectedIP:    "10.0.0.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			req.RemoteAddr = tt.remoteAddr
			if tt.xForwardedFor != "" {
				req.Header.Set("X-Forwarded-For", tt.xForwardedFor)
			}
			if tt.xRealIP != "" {
				req.Header.Set("X-Real-IP", tt.xRealIP)
			}
			w := httptest.NewRecorder()

			ctx := NewContext(w, req)

			ip := ctx.ClientIP()
			if ip != tt.expectedIP {
				t.Errorf("ClientIP() = %s, want %s", ip, tt.expectedIP)
			}
		})
	}
}

// Helper function to check if a string contains a substring
// 문자열이 하위 문자열을 포함하는지 확인하는 헬퍼 함수
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
