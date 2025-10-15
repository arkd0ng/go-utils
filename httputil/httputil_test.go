package httputil

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

// TestVersion tests that the version is properly loaded.
// TestVersion은 버전이 제대로 로드되는지 테스트합니다.
func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}

	if Version != "v1.10.002" {
		t.Errorf("Expected version 'v1.10.002', got '%s'", Version)
	}
}

// TestNewClient tests that a new client can be created.
// TestNewClient는 새 클라이언트를 생성할 수 있는지 테스트합니다.
func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("NewClient should not return nil")
	}

	if client.client == nil {
		t.Error("Client should have an http.Client")
	}

	if client.config == nil {
		t.Error("Client should have a config")
	}
}

// TestNewClientWithOptions tests creating a client with options.
// TestNewClientWithOptions는 옵션으로 클라이언트를 생성하는 것을 테스트합니다.
func TestNewClientWithOptions(t *testing.T) {
	client := NewClient(
		WithBaseURL("https://api.example.com"),
		WithBearerToken("test-token"),
		WithRetry(5),
	)

	if client == nil {
		t.Fatal("NewClient should not return nil")
	}

	if client.config.baseURL != "https://api.example.com" {
		t.Errorf("Expected baseURL 'https://api.example.com', got '%s'", client.config.baseURL)
	}

	if client.config.bearerToken != "test-token" {
		t.Errorf("Expected bearerToken 'test-token', got '%s'", client.config.bearerToken)
	}

	if client.config.maxRetries != 5 {
		t.Errorf("Expected maxRetries 5, got %d", client.config.maxRetries)
	}
}

// TestHTTPError tests HTTPError functionality.
// TestHTTPError는 HTTPError 기능을 테스트합니다.
func TestHTTPError(t *testing.T) {
	err := &HTTPError{
		StatusCode: 404,
		Status:     "404 Not Found",
		Body:       "Resource not found",
		URL:        "https://api.example.com/notfound",
		Method:     "GET",
	}

	expected := "HTTP 404 404 Not Found: Not Found (URL: GET https://api.example.com/notfound, Body: Resource not found)"
	if err.Error() != expected {
		t.Errorf("Expected error message:\n%s\nGot:\n%s", expected, err.Error())
	}

	if !IsHTTPError(err) {
		t.Error("IsHTTPError should return true for HTTPError")
	}

	if GetStatusCode(err) != 404 {
		t.Errorf("GetStatusCode should return 404, got %d", GetStatusCode(err))
	}
}

// TestRetryError tests RetryError functionality.
// TestRetryError는 RetryError 기능을 테스트합니다.
func TestRetryError(t *testing.T) {
	err := &RetryError{
		Attempts: 3,
		LastErr:  &HTTPError{StatusCode: 500},
		URL:      "https://api.example.com/retry",
		Method:   "POST",
	}

	if !IsRetryError(err) {
		t.Error("IsRetryError should return true for RetryError")
	}

	if err.Unwrap() == nil {
		t.Error("Unwrap should return the last error")
	}
}

// TestTimeoutError tests TimeoutError functionality.
// TestTimeoutError는 TimeoutError 기능을 테스트합니다.
func TestTimeoutError(t *testing.T) {
	err := &TimeoutError{
		URL:    "https://api.example.com/slow",
		Method: "GET",
	}

	if !IsTimeoutError(err) {
		t.Error("IsTimeoutError should return true for TimeoutError")
	}

	if !err.Timeout() {
		t.Error("Timeout() should return true for TimeoutError")
	}

	expected := "request timeout (URL: GET https://api.example.com/slow)"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

// TestDefaultConfig tests that default configuration has sensible values.
// TestDefaultConfig는 기본 설정이 합리적인 값을 가지는지 테스트합니다.
func TestDefaultConfig(t *testing.T) {
	cfg := defaultConfig()

	if cfg.timeout == 0 {
		t.Error("Default timeout should not be 0")
	}

	if cfg.maxRetries < 0 {
		t.Error("Default maxRetries should not be negative")
	}

	if cfg.userAgent == "" {
		t.Error("Default userAgent should not be empty")
	}

	if !cfg.followRedirects {
		t.Error("Default should follow redirects")
	}
}

// TestResponse tests Response wrapper functionality.
// TestResponse는 Response 래퍼 기능을 테스트합니다.
func TestResponse(t *testing.T) {
	t.Run("Body methods", func(t *testing.T) {
		resp := &Response{
			body: []byte(`{"status":"ok"}`),
		}

		// Test Body()
		if string(resp.Body()) != `{"status":"ok"}` {
			t.Errorf("Body() returned unexpected value: %s", resp.Body())
		}

		// Test String()
		if resp.String() != `{"status":"ok"}` {
			t.Errorf("String() returned unexpected value: %s", resp.String())
		}

		// Test JSON()
		var result map[string]string
		if err := resp.JSON(&result); err != nil {
			t.Errorf("JSON() failed: %v", err)
		}
		if result["status"] != "ok" {
			t.Errorf("Expected status 'ok', got '%s'", result["status"])
		}
	})

	t.Run("Status checks", func(t *testing.T) {
		tests := []struct {
			name       string
			statusCode int
			checkFunc  func(*Response) bool
			expected   bool
		}{
			{"IsSuccess with 200", 200, (*Response).IsSuccess, true},
			{"IsSuccess with 404", 404, (*Response).IsSuccess, false},
			{"IsError with 404", 404, (*Response).IsError, true},
			{"IsError with 200", 200, (*Response).IsError, false},
			{"IsClientError with 404", 404, (*Response).IsClientError, true},
			{"IsClientError with 500", 500, (*Response).IsClientError, false},
			{"IsServerError with 500", 500, (*Response).IsServerError, true},
			{"IsServerError with 404", 404, (*Response).IsServerError, false},
			{"IsOK with 200", 200, (*Response).IsOK, true},
			{"IsOK with 201", 201, (*Response).IsOK, false},
			{"IsCreated with 201", 201, (*Response).IsCreated, true},
			{"IsNotFound with 404", 404, (*Response).IsNotFound, true},
			{"IsUnauthorized with 401", 401, (*Response).IsUnauthorized, true},
			{"IsForbidden with 403", 403, (*Response).IsForbidden, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				resp := &Response{
					Response: &http.Response{
						StatusCode: tt.statusCode,
					},
				}
				if got := tt.checkFunc(resp); got != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, got)
				}
			})
		}
	})
}

// TestURLBuilder tests URL building functionality.
// TestURLBuilder는 URL 구축 기능을 테스트합니다.
func TestURLBuilder(t *testing.T) {
	t.Run("Basic URL building", func(t *testing.T) {
		url := NewURL("https://api.example.com").
			Path("users", "123").
			Param("include", "posts").
			Build()

		expected := "https://api.example.com/users/123?include=posts"
		if url != expected {
			t.Errorf("Expected URL '%s', got '%s'", expected, url)
		}
	})

	t.Run("Multiple parameters", func(t *testing.T) {
		url := NewURL("https://api.example.com").
			Path("search").
			Params(map[string]string{
				"q":     "golang",
				"page":  "1",
				"limit": "20",
			}).
			Build()

		if !strings.Contains(url, "q=golang") {
			t.Error("URL should contain 'q=golang'")
		}
		if !strings.Contains(url, "page=1") {
			t.Error("URL should contain 'page=1'")
		}
	})

	t.Run("Conditional parameters", func(t *testing.T) {
		hasFilter := true
		url := NewURL("https://api.example.com").
			Path("users").
			ParamIf(hasFilter, "filter", "active").
			ParamIf(false, "deleted", "true").
			Build()

		if !strings.Contains(url, "filter=active") {
			t.Error("URL should contain 'filter=active'")
		}
		if strings.Contains(url, "deleted") {
			t.Error("URL should not contain 'deleted' parameter")
		}
	})
}

// TestFormBuilder tests form building functionality.
// TestFormBuilder는 폼 구축 기능을 테스트합니다.
func TestFormBuilder(t *testing.T) {
	t.Run("Basic form building", func(t *testing.T) {
		form := NewForm().
			Set("username", "john").
			Set("email", "john@example.com").
			Set("age", "30")

		if form.Get("username") != "john" {
			t.Errorf("Expected username 'john', got '%s'", form.Get("username"))
		}

		encoded := form.Encode()
		if !strings.Contains(encoded, "username=john") {
			t.Error("Encoded form should contain 'username=john'")
		}
	})

	t.Run("Multiple values", func(t *testing.T) {
		form := NewForm().
			AddMultiple("tags", "go", "http", "api")

		values := form.GetAll("tags")
		if len(values) != 3 {
			t.Errorf("Expected 3 tags, got %d", len(values))
		}
	})

	t.Run("Conditional fields", func(t *testing.T) {
		hasPromo := true
		form := NewForm().
			Set("username", "john").
			AddIf(hasPromo, "promo_code", "SAVE20").
			AddIf(false, "referrer", "none")

		if !form.Has("promo_code") {
			t.Error("Form should have 'promo_code'")
		}
		if form.Has("referrer") {
			t.Error("Form should not have 'referrer'")
		}
	})

	t.Run("Clone", func(t *testing.T) {
		original := NewForm().Set("key", "value")
		cloned := original.Clone().Set("key2", "value2")

		if cloned.Get("key") != "value" {
			t.Error("Cloned form should have original data")
		}
		if original.Has("key2") {
			t.Error("Original form should not have new key")
		}
	})

	t.Run("Map conversion", func(t *testing.T) {
		form := NewForm().
			Set("name", "John").
			Set("city", "Seoul")

		m := form.Map()
		if m["name"] != "John" || m["city"] != "Seoul" {
			t.Error("Map conversion failed")
		}
	})
}

// TestURLUtilities tests URL utility functions.
// TestURLUtilities는 URL 유틸리티 함수를 테스트합니다.
func TestURLUtilities(t *testing.T) {
	t.Run("JoinURL", func(t *testing.T) {
		url := JoinURL("https://api.example.com", "v1", "users", "123")
		expected := "https://api.example.com/v1/users/123"
		if url != expected {
			t.Errorf("Expected '%s', got '%s'", expected, url)
		}
	})

	t.Run("AddQueryParams", func(t *testing.T) {
		url, err := AddQueryParams("https://api.example.com/users", map[string]string{
			"page":  "1",
			"limit": "20",
		})
		if err != nil {
			t.Fatalf("AddQueryParams failed: %v", err)
		}
		if !strings.Contains(url, "page=1") {
			t.Error("URL should contain 'page=1'")
		}
	})

	t.Run("GetDomain", func(t *testing.T) {
		domain, err := GetDomain("https://api.example.com:8080/path")
		if err != nil {
			t.Fatalf("GetDomain failed: %v", err)
		}
		if domain != "api.example.com:8080" {
			t.Errorf("Expected 'api.example.com:8080', got '%s'", domain)
		}
	})

	t.Run("GetScheme", func(t *testing.T) {
		scheme, err := GetScheme("https://api.example.com")
		if err != nil {
			t.Fatalf("GetScheme failed: %v", err)
		}
		if scheme != "https" {
			t.Errorf("Expected 'https', got '%s'", scheme)
		}
	})

	t.Run("IsAbsoluteURL", func(t *testing.T) {
		if !IsAbsoluteURL("https://example.com") {
			t.Error("https://example.com should be absolute")
		}
		if IsAbsoluteURL("/relative/path") {
			t.Error("/relative/path should not be absolute")
		}
	})

	t.Run("NormalizeURL", func(t *testing.T) {
		url := NormalizeURL("  https://example.com/  ")
		if url != "https://example.com" {
			t.Errorf("Expected 'https://example.com', got '%s'", url)
		}
	})
}

// TestFormUtilities tests form utility functions.
// TestFormUtilities는 폼 유틸리티 함수를 테스트합니다.
func TestFormUtilities(t *testing.T) {
	t.Run("ParseForm", func(t *testing.T) {
		data, err := ParseForm("name=John&city=Seoul&age=30")
		if err != nil {
			t.Fatalf("ParseForm failed: %v", err)
		}
		if data["name"] != "John" {
			t.Errorf("Expected name 'John', got '%s'", data["name"])
		}
		if data["city"] != "Seoul" {
			t.Errorf("Expected city 'Seoul', got '%s'", data["city"])
		}
	})

	t.Run("EncodeForm", func(t *testing.T) {
		encoded := EncodeForm(map[string]string{
			"username": "john",
			"password": "secret",
		})
		if !strings.Contains(encoded, "username=john") {
			t.Error("Encoded form should contain 'username=john'")
		}
	})
}

// TestProgressReader tests progress callback functionality.
// TestProgressReader는 진행 상황 콜백 기능을 테스트합니다.
func TestProgressReader(t *testing.T) {
	data := []byte("Hello, World!")
	reader := bytes.NewReader(data)

	var totalRead int64
	progressFunc := func(read, total int64) {
		totalRead = read
	}

	pr := &progressReader{
		reader:   reader,
		progress: progressFunc,
		total:    int64(len(data)),
	}

	buf := make([]byte, len(data))
	n, err := pr.Read(buf)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}

	if n != len(data) {
		t.Errorf("Expected to read %d bytes, got %d", len(data), n)
	}

	if totalRead != int64(len(data)) {
		t.Errorf("Progress callback should report %d bytes, got %d", len(data), totalRead)
	}
}
