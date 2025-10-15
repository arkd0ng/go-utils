package httputil

import (
	"testing"
)

// TestVersion tests that the version is properly loaded.
// TestVersion은 버전이 제대로 로드되는지 테스트합니다.
func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}

	if Version != "v1.10.001" {
		t.Errorf("Expected version 'v1.10.001', got '%s'", Version)
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
