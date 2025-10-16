package websvrutil

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestCSRF tests basic CSRF middleware functionality.
// TestCSRF는 기본 CSRF 미들웨어 기능을 테스트합니다.
func TestCSRF(t *testing.T) {
	app := New(WithTemplateDir(""))
	app.Use(CSRF())

	// GET request should work without CSRF token
	// GET 요청은 CSRF 토큰 없이 작동해야 함
	app.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Extract CSRF token from cookie
	// 쿠키에서 CSRF 토큰 추출
	cookies := rec.Result().Cookies()
	var csrfToken string
	for _, cookie := range cookies {
		if cookie.Name == "_csrf" {
			csrfToken = cookie.Value
			break
		}
	}

	if csrfToken == "" {
		t.Fatal("Expected CSRF token in cookie")
	}

	// POST request without CSRF token should fail
	// CSRF 토큰 없는 POST 요청은 실패해야 함
	app.POST("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req = httptest.NewRequest("POST", "/test", nil)
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Errorf("Expected status 403 without CSRF token, got %d", rec.Code)
	}

	// POST request with valid CSRF token should work
	// 유효한 CSRF 토큰이 있는 POST 요청은 작동해야 함
	req = httptest.NewRequest("POST", "/test", nil)
	req.AddCookie(&http.Cookie{Name: "_csrf", Value: csrfToken})
	req.Header.Set("X-CSRF-Token", csrfToken)
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200 with valid CSRF token, got %d", rec.Code)
	}
}

// TestCSRFWithConfig tests CSRF middleware with custom configuration.
// TestCSRFWithConfig는 커스텀 설정으로 CSRF 미들웨어를 테스트합니다.
func TestCSRFWithConfig(t *testing.T) {
	config := CSRFConfig{
		TokenLength: 16,
		CookieName:  "custom_csrf",
		TokenLookup: "header:X-Custom-Token",
	}

	app := New(WithTemplateDir(""))
	app.Use(CSRFWithConfig(config))

	app.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	// Extract CSRF token from custom cookie
	// 커스텀 쿠키에서 CSRF 토큰 추출
	cookies := rec.Result().Cookies()
	var csrfToken string
	for _, cookie := range cookies {
		if cookie.Name == "custom_csrf" {
			csrfToken = cookie.Value
			break
		}
	}

	if csrfToken == "" {
		t.Fatal("Expected CSRF token in custom cookie")
	}

	// POST request with custom header should work
	// 커스텀 헤더가 있는 POST 요청은 작동해야 함
	app.POST("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req = httptest.NewRequest("POST", "/test", nil)
	req.AddCookie(&http.Cookie{Name: "custom_csrf", Value: csrfToken})
	req.Header.Set("X-Custom-Token", csrfToken)
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200 with custom CSRF token, got %d", rec.Code)
	}
}

// TestCSRFFormToken tests CSRF token in form data.
// TestCSRFFormToken은 폼 데이터의 CSRF 토큰을 테스트합니다.
func TestCSRFFormToken(t *testing.T) {
	config := CSRFConfig{
		TokenLookup: "form:csrf_token",
	}

	app := New(WithTemplateDir(""))
	app.Use(CSRFWithConfig(config))

	app.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	// Extract CSRF token
	// CSRF 토큰 추출
	cookies := rec.Result().Cookies()
	var csrfToken string
	for _, cookie := range cookies {
		if cookie.Name == "_csrf" {
			csrfToken = cookie.Value
			break
		}
	}

	// POST with form data
	// 폼 데이터가 있는 POST
	app.POST("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	formData := "csrf_token=" + csrfToken
	req = httptest.NewRequest("POST", "/test", strings.NewReader(formData))
	req.AddCookie(&http.Cookie{Name: "_csrf", Value: csrfToken})
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200 with form CSRF token, got %d", rec.Code)
	}
}

// TestCSRFSkipper tests CSRF middleware with Skipper function.
// TestCSRFSkipper는 Skipper 함수가 있는 CSRF 미들웨어를 테스트합니다.
func TestCSRFSkipper(t *testing.T) {
	config := CSRFConfig{
		Skipper: func(r *http.Request) bool {
			// Skip CSRF for /api/* routes
			// /api/* 라우트는 CSRF 건너뛰기
			return strings.HasPrefix(r.URL.Path, "/api/")
		},
	}

	app := New(WithTemplateDir(""))
	app.Use(CSRFWithConfig(config))

	app.POST("/api/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	app.POST("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// POST to /api/test should work without CSRF token (skipped)
	// /api/test로의 POST는 CSRF 토큰 없이 작동해야 함 (건너뛰기)
	req := httptest.NewRequest("POST", "/api/test", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200 for skipped route, got %d", rec.Code)
	}

	// POST to /test should fail without CSRF token
	// /test로의 POST는 CSRF 토큰 없이 실패해야 함
	req = httptest.NewRequest("POST", "/test", nil)
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Errorf("Expected status 403 for non-skipped route without CSRF token, got %d", rec.Code)
	}
}

// TestCSRFTokenGeneration tests CSRF token generation.
// TestCSRFTokenGeneration은 CSRF 토큰 생성을 테스트합니다.
func TestCSRFTokenGeneration(t *testing.T) {
	// Generate multiple tokens and ensure they're unique
	// 여러 토큰을 생성하고 고유한지 확인
	tokens := make(map[string]bool)
	for i := 0; i < 100; i++ {
		token, err := generateCSRFToken(32)
		if err != nil {
			t.Fatalf("Failed to generate CSRF token: %v", err)
		}

		if token == "" {
			t.Error("Generated empty CSRF token")
		}

		if tokens[token] {
			t.Error("Generated duplicate CSRF token")
		}

		tokens[token] = true
	}
}

// TestIsSafeMethod tests the isSafeMethod function.
// TestIsSafeMethod는 isSafeMethod 함수를 테스트합니다.
func TestIsSafeMethod(t *testing.T) {
	safeMethods := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodOptions,
		http.MethodTrace,
	}

	unsafeMethods := []string{
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	}

	for _, method := range safeMethods {
		if !isSafeMethod(method) {
			t.Errorf("Expected %s to be safe method", method)
		}
	}

	for _, method := range unsafeMethods {
		if isSafeMethod(method) {
			t.Errorf("Expected %s to be unsafe method", method)
		}
	}
}

// TestSplitTokenLookup tests the splitTokenLookup function.
// TestSplitTokenLookup는 splitTokenLookup 함수를 테스트합니다.
func TestSplitTokenLookup(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"header:X-CSRF-Token", []string{"header", "X-CSRF-Token"}},
		{"form:csrf_token", []string{"form", "csrf_token"}},
		{"query:token", []string{"query", "token"}},
		{"invalid", []string{"invalid"}},
	}

	for _, tt := range tests {
		result := splitTokenLookup(tt.input)
		if len(result) != len(tt.expected) {
			t.Errorf("Expected %d parts, got %d for input %s", len(tt.expected), len(result), tt.input)
			continue
		}

		for i := range result {
			if result[i] != tt.expected[i] {
				t.Errorf("Expected part %d to be %s, got %s for input %s", i, tt.expected[i], result[i], tt.input)
			}
		}
	}
}
