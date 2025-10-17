package httputil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client wraps http.Client with additional functionality.
// Client는 추가 기능을 가진 http.Client를 래핑합니다.
type Client struct {
	client *http.Client
	config *config
	// Optional cookie jar with persistence
	// 지속성을 가진 선택적 쿠키 저장소
	cookieJar *CookieJar
}

// NewClient creates a new HTTP client with the given options.
// NewClient는 주어진 옵션으로 새로운 HTTP 클라이언트를 생성합니다.
func NewClient(opts ...Option) *Client {
	cfg := defaultConfig()
	cfg.apply(opts)

	// Create HTTP client
	// HTTP 클라이언트 생성
	client := &http.Client{
		Timeout: cfg.timeout,
	}

	// Configure redirect policy
	// 리디렉션 정책 설정
	if !cfg.followRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	} else if cfg.maxRedirects > 0 {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			if len(via) >= cfg.maxRedirects {
				return fmt.Errorf("stopped after %d redirects", cfg.maxRedirects)
			}
			return nil
		}
	}

	// Configure TLS
	// TLS 설정
	if cfg.tlsConfig != nil {
		transport := &http.Transport{
			TLSClientConfig: cfg.tlsConfig,
		}
		client.Transport = transport
	}

	// Configure proxy
	// 프록시 설정
	if cfg.proxyURL != "" {
		if proxyURL, err := url.Parse(cfg.proxyURL); err == nil {
			transport := client.Transport
			if transport == nil {
				transport = http.DefaultTransport
			}
			if t, ok := transport.(*http.Transport); ok {
				t.Proxy = http.ProxyURL(proxyURL)
				client.Transport = t
			}
		}
	}

	// Configure cookie jar
	// 쿠키 저장소 설정
	if cfg.cookieJar != nil {
		client.Jar = cfg.cookieJar
	}

	// Initialize custom cookie jar
	// 사용자 정의 쿠키 저장소 초기화
	var customCookieJar *CookieJar
	if cfg.cookieJarPath != "" {
		// Create persistent cookie jar
		// 지속성 쿠키 저장소 생성
		jar, err := NewPersistentCookieJar(cfg.cookieJarPath)
		if err == nil {
			customCookieJar = jar
			client.Jar = jar.jar
		}
	} else if cfg.enableCookieJar {
		// Create in-memory cookie jar
		// 메모리 내 쿠키 저장소 생성
		jar, err := NewCookieJar()
		if err == nil {
			customCookieJar = jar
			client.Jar = jar.jar
		}
	}

	return &Client{
		client:    client,
		config:    cfg,
		cookieJar: customCookieJar,
	}
}

// Get performs a GET request and decodes the JSON response into result.
// Get은 GET 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) Get(path string, result interface{}, opts ...Option) error {
	return c.GetContext(context.Background(), path, result, opts...)
}

// GetContext performs a GET request with context and decodes the JSON response into result.
// GetContext는 context와 함께 GET 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) GetContext(ctx context.Context, path string, result interface{}, opts ...Option) error {
	return c.doRequest(ctx, http.MethodGet, path, nil, result, opts...)
}

// Post performs a POST request with body and decodes the JSON response into result.
// Post는 body와 함께 POST 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) Post(path string, body, result interface{}, opts ...Option) error {
	return c.PostContext(context.Background(), path, body, result, opts...)
}

// PostContext performs a POST request with context, body and decodes the JSON response into result.
// PostContext는 context, body와 함께 POST 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) PostContext(ctx context.Context, path string, body, result interface{}, opts ...Option) error {
	return c.doRequest(ctx, http.MethodPost, path, body, result, opts...)
}

// Put performs a PUT request with body and decodes the JSON response into result.
// Put은 body와 함께 PUT 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) Put(path string, body, result interface{}, opts ...Option) error {
	return c.PutContext(context.Background(), path, body, result, opts...)
}

// PutContext performs a PUT request with context, body and decodes the JSON response into result.
// PutContext는 context, body와 함께 PUT 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) PutContext(ctx context.Context, path string, body, result interface{}, opts ...Option) error {
	return c.doRequest(ctx, http.MethodPut, path, body, result, opts...)
}

// Patch performs a PATCH request with body and decodes the JSON response into result.
// Patch는 body와 함께 PATCH 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) Patch(path string, body, result interface{}, opts ...Option) error {
	return c.PatchContext(context.Background(), path, body, result, opts...)
}

// PatchContext performs a PATCH request with context, body and decodes the JSON response into result.
// PatchContext는 context, body와 함께 PATCH 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) PatchContext(ctx context.Context, path string, body, result interface{}, opts ...Option) error {
	return c.doRequest(ctx, http.MethodPatch, path, body, result, opts...)
}

// Delete performs a DELETE request and decodes the JSON response into result.
// Delete는 DELETE 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) Delete(path string, result interface{}, opts ...Option) error {
	return c.DeleteContext(context.Background(), path, result, opts...)
}

// DeleteContext performs a DELETE request with context and decodes the JSON response into result.
// DeleteContext는 context와 함께 DELETE 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
func (c *Client) DeleteContext(ctx context.Context, path string, result interface{}, opts ...Option) error {
	return c.doRequest(ctx, http.MethodDelete, path, nil, result, opts...)
}

// doRequest performs an HTTP request with retry logic.
// doRequest는 재시도 로직을 사용하여 HTTP 요청을 수행합니다.
func (c *Client) doRequest(ctx context.Context, method, path string, body, result interface{}, opts ...Option) error {
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

	// Prepare request body
	// 요청 본문 준비
	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	// Retry logic
	// 재시도 로직
	var lastErr error
	for attempt := 0; attempt <= cfg.maxRetries; attempt++ {
		// Create request
		// 요청 생성
		req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		// Set headers
		// 헤더 설정
		req.Header.Set("Content-Type", "application/json")
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
			// Check if it's a timeout error
			// 타임아웃 에러인지 확인
			if ctx.Err() == context.DeadlineExceeded {
				return &TimeoutError{URL: fullURL, Method: method}
			}

			lastErr = err

			// Retry on network errors
			// 네트워크 에러 시 재시도
			if attempt < cfg.maxRetries {
				// Calculate backoff
				// 백오프 계산
				backoff := calculateBackoff(attempt, cfg.retryMin, cfg.retryMax)
				time.Sleep(backoff)
				continue
			}
			break
		}
		defer resp.Body.Close()

		// Check status code
		// 상태 코드 확인
		if resp.StatusCode >= 400 {
			// Read error body
			// 에러 본문 읽기
			bodyBytes, _ := io.ReadAll(resp.Body)
			httpErr := &HTTPError{
				StatusCode: resp.StatusCode,
				Status:     resp.Status,
				Body:       string(bodyBytes),
				URL:        fullURL,
				Method:     method,
			}

			// Retry on 5xx errors
			// 5xx 에러 시 재시도
			if resp.StatusCode >= 500 && attempt < cfg.maxRetries {
				lastErr = httpErr
				backoff := calculateBackoff(attempt, cfg.retryMin, cfg.retryMax)
				time.Sleep(backoff)
				continue
			}

			return httpErr
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

	// All retries failed
	// 모든 재시도 실패
	if lastErr != nil {
		return &RetryError{
			Attempts: cfg.maxRetries + 1,
			LastErr:  lastErr,
			URL:      fullURL,
			Method:   method,
		}
	}

	return fmt.Errorf("request failed with unknown error")
}

// calculateBackoff calculates the backoff duration for the given attempt.
// calculateBackoff는 주어진 시도에 대한 백오프 기간을 계산합니다.
func calculateBackoff(attempt int, min, max time.Duration) time.Duration {
	// Exponential backoff with jitter
	// 지터가 있는 지수 백오프
	backoff := min * time.Duration(math.Pow(2, float64(attempt)))
	if backoff > max {
		backoff = max
	}

	// Add jitter (±25%)
	// 지터 추가 (±25%)
	jitter := time.Duration(rand.Int63n(int64(backoff / 4)))
	if rand.Intn(2) == 0 {
		backoff += jitter
	} else {
		backoff -= jitter
	}

	return backoff
}
