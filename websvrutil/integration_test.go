package websvrutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestIntegrationFullApp tests a complete app with multiple features.
// TestIntegrationFullApp은 여러 기능이 있는 완전한 앱을 테스트합니다.
func TestIntegrationFullApp(t *testing.T) {
	// Create app with middleware / 미들웨어가 있는 앱 생성
	app := New(WithTemplateDir(""))
	app.Use(Logger())
	app.Use(Recovery())
	app.Use(CORS())

	// Define routes / 라우트 정의
	app.GET("/health", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		ctx.JSON(200, map[string]string{"status": "ok"})
	})

	type User struct {
		Name  string `json:"name" validate:"required,min=3"`
		Email string `json:"email" validate:"required,email"`
	}

	app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		var user User
		if err := ctx.BindWithValidation(&user); err != nil {
			ctx.ErrorJSON(400, err.Error())
			return
		}
		ctx.SuccessJSON(201, "User created", user)
	})

	// Test GET /health / GET /health 테스트
	req := httptest.NewRequest("GET", "/health", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Test POST /users with valid data / 유효한 데이터로 POST /users 테스트
	userData := User{Name: "John Doe", Email: "john@example.com"}
	body, _ := json.Marshal(userData)
	req = httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != 201 {
		t.Errorf("Expected status 201, got %d", rec.Code)
	}

	// Test POST /users with invalid data / 무효한 데이터로 POST /users 테스트
	invalidUser := User{Name: "Jo", Email: "invalid"}
	body, _ = json.Marshal(invalidUser)
	req = httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != 400 {
		t.Errorf("Expected status 400 for invalid user, got %d", rec.Code)
	}
}

// TestIntegrationRouteGroups tests route groups integration.
// TestIntegrationRouteGroups는 라우트 그룹 통합을 테스트합니다.
func TestIntegrationRouteGroups(t *testing.T) {
	app := New(WithTemplateDir(""))

	// Create API v1 group / API v1 그룹 생성
	v1 := app.Group("/api/v1")
	v1.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		ctx.JSON(200, []string{"user1", "user2"})
	})

	// Create API v2 group / API v2 그룹 생성
	v2 := app.Group("/api/v2")
	v2.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		ctx.JSON(200, []string{"user1", "user2", "user3"})
	})

	// Test v1 / v1 테스트
	req := httptest.NewRequest("GET", "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Test v2 / v2 테스트
	req = httptest.NewRequest("GET", "/api/v2/users", nil)
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}

// TestIntegrationCSRFWithValidation tests CSRF and validation together.
// TestIntegrationCSRFWithValidation은 CSRF와 검증을 함께 테스트합니다.
func TestIntegrationCSRFWithValidation(t *testing.T) {
	app := New(WithTemplateDir(""))
	app.Use(CSRF())

	type FormData struct {
		Name string `json:"name" validate:"required,min=3"`
	}

	app.POST("/submit", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		var data FormData
		if err := ctx.BindWithValidation(&data); err != nil {
			ctx.ErrorJSON(400, err.Error())
			return
		}
		ctx.SuccessJSON(200, "Success", data)
	})

	// First GET to obtain CSRF token / 먼저 GET으로 CSRF 토큰 얻기
	app.GET("/form", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	req := httptest.NewRequest("GET", "/form", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	// Extract CSRF token / CSRF 토큰 추출
	cookies := rec.Result().Cookies()
	var csrfToken string
	for _, cookie := range cookies {
		if cookie.Name == "_csrf" {
			csrfToken = cookie.Value
			break
		}
	}

	if csrfToken == "" {
		t.Fatal("CSRF token not found")
	}

	// POST with CSRF token and valid data / CSRF 토큰과 유효한 데이터로 POST
	formData := FormData{Name: "John"}
	body, _ := json.Marshal(formData)
	req = httptest.NewRequest("POST", "/submit", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CSRF-Token", csrfToken)
	req.AddCookie(&http.Cookie{Name: "_csrf", Value: csrfToken})
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}
