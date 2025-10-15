package httputil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Response wraps http.Response with additional functionality.
// Response는 추가 기능을 가진 http.Response를 래핑합니다.
type Response struct {
	*http.Response
	body []byte // Cached body for multiple reads / 여러 번 읽기 위한 캐시된 본문
}

// DoRaw performs an HTTP request and returns the raw response.
// DoRaw는 HTTP 요청을 수행하고 원시 응답을 반환합니다.
func (c *Client) DoRaw(method, path string, body interface{}, opts ...Option) (*Response, error) {
	return c.DoRawContext(context.Background(), method, path, body, opts...)
}

// DoRawContext performs an HTTP request with context and returns the raw response.
// DoRawContext는 context와 함께 HTTP 요청을 수행하고 원시 응답을 반환합니다.
func (c *Client) DoRawContext(ctx context.Context, method, path string, body interface{}, opts ...Option) (*Response, error) {
	// Merge client config with request-specific options
	// 클라이언트 설정을 요청별 옵션과 병합
	cfg := *c.config
	cfg.apply(opts)

	// Build full URL / 전체 URL 구축
	fullURL := path
	if cfg.baseURL != "" && !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		fullURL = strings.TrimRight(cfg.baseURL, "/") + "/" + strings.TrimLeft(path, "/")
	}

	// Add query parameters / 쿼리 매개변수 추가
	if len(cfg.queryParams) > 0 {
		u, err := url.Parse(fullURL)
		if err != nil {
			return nil, fmt.Errorf("invalid URL: %w", err)
		}
		q := u.Query()
		for k, v := range cfg.queryParams {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
		fullURL = u.String()
	}

	// Prepare request body / 요청 본문 준비
	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	// Create request / 요청 생성
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers / 헤더 설정
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", cfg.userAgent)
	for k, v := range cfg.headers {
		req.Header.Set(k, v)
	}

	// Set authentication / 인증 설정
	if cfg.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.bearerToken)
	}
	if cfg.basicAuthUser != "" {
		req.SetBasicAuth(cfg.basicAuthUser, cfg.basicAuthPass)
	}

	// Execute request / 요청 실행
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read and cache body / 본문 읽기 및 캐시
	bodyBytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return &Response{
		Response: resp,
		body:     bodyBytes,
	}, nil
}

// Body returns the response body as bytes.
// Body는 응답 본문을 바이트로 반환합니다.
func (r *Response) Body() []byte {
	return r.body
}

// String returns the response body as string.
// String은 응답 본문을 문자열로 반환합니다.
func (r *Response) String() string {
	return string(r.body)
}

// JSON decodes the response body as JSON into result.
// JSON은 응답 본문을 JSON으로 result에 디코딩합니다.
func (r *Response) JSON(result interface{}) error {
	if err := json.Unmarshal(r.body, result); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}
	return nil
}

// IsSuccess returns true if status code is 2xx.
// IsSuccess는 상태 코드가 2xx이면 true를 반환합니다.
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// IsRedirect returns true if status code is 3xx.
// IsRedirect는 상태 코드가 3xx이면 true를 반환합니다.
func (r *Response) IsRedirect() bool {
	return r.StatusCode >= 300 && r.StatusCode < 400
}

// IsClientError returns true if status code is 4xx.
// IsClientError는 상태 코드가 4xx이면 true를 반환합니다.
func (r *Response) IsClientError() bool {
	return r.StatusCode >= 400 && r.StatusCode < 500
}

// IsServerError returns true if status code is 5xx.
// IsServerError는 상태 코드가 5xx이면 true를 반환합니다.
func (r *Response) IsServerError() bool {
	return r.StatusCode >= 500 && r.StatusCode < 600
}

// IsError returns true if status code is 4xx or 5xx.
// IsError는 상태 코드가 4xx 또는 5xx이면 true를 반환합니다.
func (r *Response) IsError() bool {
	return r.StatusCode >= 400
}

// Header returns the value of the header with the given key.
// Header는 주어진 키의 헤더 값을 반환합니다.
func (r *Response) Header(key string) string {
	return r.Response.Header.Get(key)
}

// Headers returns all headers as a map.
// Headers는 모든 헤더를 맵으로 반환합니다.
func (r *Response) Headers() map[string]string {
	headers := make(map[string]string)
	for k, v := range r.Response.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}
	return headers
}

// ContentType returns the Content-Type header value.
// ContentType은 Content-Type 헤더 값을 반환합니다.
func (r *Response) ContentType() string {
	return r.Response.Header.Get("Content-Type")
}

// ContentLength returns the Content-Length header value.
// ContentLength는 Content-Length 헤더 값을 반환합니다.
func (r *Response) ContentLength() int64 {
	return r.Response.ContentLength
}

// Status checks / 상태 확인 함수들

// IsOK returns true if status code is 200.
// IsOK는 상태 코드가 200이면 true를 반환합니다.
func (r *Response) IsOK() bool {
	return r.StatusCode == http.StatusOK
}

// IsCreated returns true if status code is 201.
// IsCreated는 상태 코드가 201이면 true를 반환합니다.
func (r *Response) IsCreated() bool {
	return r.StatusCode == http.StatusCreated
}

// IsNoContent returns true if status code is 204.
// IsNoContent는 상태 코드가 204이면 true를 반환합니다.
func (r *Response) IsNoContent() bool {
	return r.StatusCode == http.StatusNoContent
}

// IsBadRequest returns true if status code is 400.
// IsBadRequest는 상태 코드가 400이면 true를 반환합니다.
func (r *Response) IsBadRequest() bool {
	return r.StatusCode == http.StatusBadRequest
}

// IsUnauthorized returns true if status code is 401.
// IsUnauthorized는 상태 코드가 401이면 true를 반환합니다.
func (r *Response) IsUnauthorized() bool {
	return r.StatusCode == http.StatusUnauthorized
}

// IsForbidden returns true if status code is 403.
// IsForbidden는 상태 코드가 403이면 true를 반환합니다.
func (r *Response) IsForbidden() bool {
	return r.StatusCode == http.StatusForbidden
}

// IsNotFound returns true if status code is 404.
// IsNotFound는 상태 코드가 404이면 true를 반환합니다.
func (r *Response) IsNotFound() bool {
	return r.StatusCode == http.StatusNotFound
}

// IsTooManyRequests returns true if status code is 429.
// IsTooManyRequests는 상태 코드가 429이면 true를 반환합니다.
func (r *Response) IsTooManyRequests() bool {
	return r.StatusCode == http.StatusTooManyRequests
}

// IsInternalServerError returns true if status code is 500.
// IsInternalServerError는 상태 코드가 500이면 true를 반환합니다.
func (r *Response) IsInternalServerError() bool {
	return r.StatusCode == http.StatusInternalServerError
}

// IsBadGateway returns true if status code is 502.
// IsBadGateway는 상태 코드가 502이면 true를 반환합니다.
func (r *Response) IsBadGateway() bool {
	return r.StatusCode == http.StatusBadGateway
}

// IsServiceUnavailable returns true if status code is 503.
// IsServiceUnavailable는 상태 코드가 503이면 true를 반환합니다.
func (r *Response) IsServiceUnavailable() bool {
	return r.StatusCode == http.StatusServiceUnavailable
}

// IsGatewayTimeout returns true if status code is 504.
// IsGatewayTimeout는 상태 코드가 504이면 true를 반환합니다.
func (r *Response) IsGatewayTimeout() bool {
	return r.StatusCode == http.StatusGatewayTimeout
}
