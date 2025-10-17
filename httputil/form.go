package httputil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// PostForm performs a POST request with form data (application/x-www-form-urlencoded).
// PostForm은 폼 데이터와 함께 POST 요청을 수행합니다 (application/x-www-form-urlencoded).
func (c *Client) PostForm(path string, data map[string]string, result interface{}, opts ...Option) error {
	return c.PostFormContext(context.Background(), path, data, result, opts...)
}

// PostFormContext performs a POST request with form data and context.
// PostFormContext는 폼 데이터 및 context와 함께 POST 요청을 수행합니다.
func (c *Client) PostFormContext(ctx context.Context, path string, data map[string]string, result interface{}, opts ...Option) error {
	// Merge client config with request-specific options
	// 클라이언트 설정을 요청별 옵션과 병합
	cfg := *c.config
	cfg.apply(opts)

	// Build full URL
	// 전체 URL 구축
	fullURL := path
	if cfg.baseURL != "" && !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		fullURL = strings.TrimRight(cfg.baseURL, "/") + "/" + strings.TrimLeft(path, "/")
	}

	// Add query parameters
	// 쿼리 매개변수 추가
	if len(cfg.queryParams) > 0 {
		u, err := url.Parse(fullURL)
		if err != nil {
			return fmt.Errorf("invalid URL: %w", err)
		}
		q := u.Query()
		for k, v := range cfg.queryParams {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
		fullURL = u.String()
	}

	// Encode form data
	// 폼 데이터 인코딩
	formData := url.Values{}
	for k, v := range data {
		formData.Set(k, v)
	}
	body := strings.NewReader(formData.Encode())

	// Create request
	// 요청 생성
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	// 헤더 설정
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", cfg.userAgent)
	for k, v := range cfg.headers {
		req.Header.Set(k, v)
	}

	// Set authentication
	// 인증 설정
	if cfg.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.bearerToken)
	}
	if cfg.basicAuthUser != "" {
		req.SetBasicAuth(cfg.basicAuthUser, cfg.basicAuthPass)
	}

	// Execute request
	// 요청 실행
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to post form: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	// 상태 코드 확인
	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(bodyBytes),
			URL:        fullURL,
			Method:     http.MethodPost,
		}
	}

	// Decode response
	// 응답 디코딩
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// FormBuilder helps build form data.
// FormBuilder는 폼 데이터를 구축하는 데 도움을 줍니다.
type FormBuilder struct {
	values url.Values
}

// NewForm creates a new form builder.
// NewForm은 새 폼 빌더를 생성합니다.
func NewForm() *FormBuilder {
	return &FormBuilder{
		values: make(url.Values),
	}
}

// Add adds a field to the form.
// Add는 폼에 필드를 추가합니다.
func (f *FormBuilder) Add(key, value string) *FormBuilder {
	f.values.Add(key, value)
	return f
}

// Set sets a field in the form (replaces existing value).
// Set은 폼에 필드를 설정합니다 (기존 값 교체).
func (f *FormBuilder) Set(key, value string) *FormBuilder {
	f.values.Set(key, value)
	return f
}

// AddIf adds a field if the condition is true.
// AddIf는 조건이 참이면 필드를 추가합니다.
func (f *FormBuilder) AddIf(condition bool, key, value string) *FormBuilder {
	if condition {
		f.values.Add(key, value)
	}
	return f
}

// AddMultiple adds multiple values for the same key.
// AddMultiple은 동일한 키에 대해 여러 값을 추가합니다.
func (f *FormBuilder) AddMultiple(key string, values ...string) *FormBuilder {
	for _, v := range values {
		f.values.Add(key, v)
	}
	return f
}

// Delete removes a field from the form.
// Delete는 폼에서 필드를 제거합니다.
func (f *FormBuilder) Delete(key string) *FormBuilder {
	f.values.Del(key)
	return f
}

// Has checks if a field exists in the form.
// Has는 폼에 필드가 있는지 확인합니다.
func (f *FormBuilder) Has(key string) bool {
	return f.values.Has(key)
}

// Get returns the value of a field.
// Get은 필드의 값을 반환합니다.
func (f *FormBuilder) Get(key string) string {
	return f.values.Get(key)
}

// GetAll returns all values for a field.
// GetAll은 필드의 모든 값을 반환합니다.
func (f *FormBuilder) GetAll(key string) []string {
	return f.values[key]
}

// Values returns the underlying url.Values.
// Values는 기본 url.Values를 반환합니다.
func (f *FormBuilder) Values() url.Values {
	return f.values
}

// Map returns the form data as a map (first value only).
// Map은 폼 데이터를 맵으로 반환합니다 (첫 번째 값만).
func (f *FormBuilder) Map() map[string]string {
	result := make(map[string]string)
	for k, v := range f.values {
		if len(v) > 0 {
			result[k] = v[0]
		}
	}
	return result
}

// Encode returns the form data as URL-encoded string.
// Encode는 폼 데이터를 URL 인코딩된 문자열로 반환합니다.
func (f *FormBuilder) Encode() string {
	return f.values.Encode()
}

// String returns the form data as URL-encoded string (same as Encode).
// String은 폼 데이터를 URL 인코딩된 문자열로 반환합니다 (Encode와 동일).
func (f *FormBuilder) String() string {
	return f.Encode()
}

// Clone creates a copy of the form builder.
// Clone은 폼 빌더의 복사본을 생성합니다.
func (f *FormBuilder) Clone() *FormBuilder {
	newForm := NewForm()
	for k, values := range f.values {
		for _, v := range values {
			newForm.values.Add(k, v)
		}
	}
	return newForm
}

// Clear removes all fields from the form.
// Clear는 폼에서 모든 필드를 제거합니다.
func (f *FormBuilder) Clear() *FormBuilder {
	f.values = make(url.Values)
	return f
}

// ParseForm parses URL-encoded form data into a map.
// ParseForm은 URL 인코딩된 폼 데이터를 맵으로 파싱합니다.
func ParseForm(data string) (map[string]string, error) {
	values, err := url.ParseQuery(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse form data: %w", err)
	}

	result := make(map[string]string)
	for k, v := range values {
		if len(v) > 0 {
			result[k] = v[0]
		}
	}
	return result, nil
}

// EncodeForm encodes a map as URL-encoded form data.
// EncodeForm은 맵을 URL 인코딩된 폼 데이터로 인코딩합니다.
func EncodeForm(data map[string]string) string {
	values := make(url.Values)
	for k, v := range data {
		values.Set(k, v)
	}
	return values.Encode()
}
