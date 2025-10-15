package httputil

import (
	"fmt"
	"net/url"
	"strings"
)

// URLBuilder helps build URLs with path and query parameters.
// URLBuilder는 경로 및 쿼리 매개변수로 URL을 구축하는 데 도움을 줍니다.
type URLBuilder struct {
	baseURL string
	path    []string
	params  url.Values
}

// NewURL creates a new URL builder with the given base URL.
// NewURL은 주어진 기본 URL로 새 URL 빌더를 생성합니다.
func NewURL(baseURL string) *URLBuilder {
	return &URLBuilder{
		baseURL: strings.TrimRight(baseURL, "/"),
		path:    make([]string, 0),
		params:  make(url.Values),
	}
}

// Path adds path segments to the URL.
// Path는 URL에 경로 세그먼트를 추가합니다.
func (u *URLBuilder) Path(segments ...string) *URLBuilder {
	u.path = append(u.path, segments...)
	return u
}

// Param adds a query parameter to the URL.
// Param은 URL에 쿼리 매개변수를 추가합니다.
func (u *URLBuilder) Param(key, value string) *URLBuilder {
	u.params.Add(key, value)
	return u
}

// Params adds multiple query parameters to the URL.
// Params는 URL에 여러 쿼리 매개변수를 추가합니다.
func (u *URLBuilder) Params(params map[string]string) *URLBuilder {
	for k, v := range params {
		u.params.Add(k, v)
	}
	return u
}

// ParamIf adds a query parameter if the condition is true.
// ParamIf는 조건이 참이면 쿼리 매개변수를 추가합니다.
func (u *URLBuilder) ParamIf(condition bool, key, value string) *URLBuilder {
	if condition {
		u.params.Add(key, value)
	}
	return u
}

// Build returns the final URL string.
// Build는 최종 URL 문자열을 반환합니다.
func (u *URLBuilder) Build() string {
	// Build path / 경로 구축
	fullPath := u.baseURL
	if len(u.path) > 0 {
		for _, segment := range u.path {
			fullPath += "/" + strings.TrimLeft(segment, "/")
		}
	}

	// Add query parameters / 쿼리 매개변수 추가
	if len(u.params) > 0 {
		fullPath += "?" + u.params.Encode()
	}

	return fullPath
}

// String returns the final URL string (same as Build).
// String은 최종 URL 문자열을 반환합니다 (Build와 동일).
func (u *URLBuilder) String() string {
	return u.Build()
}

// URL utility functions / URL 유틸리티 함수들

// JoinURL joins base URL with path segments.
// JoinURL은 기본 URL을 경로 세그먼트와 결합합니다.
func JoinURL(baseURL string, paths ...string) string {
	result := strings.TrimRight(baseURL, "/")
	for _, p := range paths {
		result += "/" + strings.TrimLeft(p, "/")
	}
	return result
}

// AddQueryParams adds query parameters to a URL.
// AddQueryParams는 URL에 쿼리 매개변수를 추가합니다.
func AddQueryParams(urlStr string, params map[string]string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// ParseURL parses a URL string and returns *url.URL.
// ParseURL은 URL 문자열을 파싱하고 *url.URL을 반환합니다.
func ParseURL(urlStr string) (*url.URL, error) {
	return url.Parse(urlStr)
}

// GetQueryParam returns a query parameter value from URL.
// GetQueryParam은 URL에서 쿼리 매개변수 값을 반환합니다.
func GetQueryParam(urlStr, key string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}
	return u.Query().Get(key), nil
}

// GetAllQueryParams returns all query parameters from URL.
// GetAllQueryParams는 URL에서 모든 쿼리 매개변수를 반환합니다.
func GetAllQueryParams(urlStr string) (map[string]string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	params := make(map[string]string)
	for k, v := range u.Query() {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}
	return params, nil
}

// RemoveQueryParam removes a query parameter from URL.
// RemoveQueryParam은 URL에서 쿼리 매개변수를 제거합니다.
func RemoveQueryParam(urlStr, key string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	q := u.Query()
	q.Del(key)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// IsAbsoluteURL checks if a URL is absolute (has scheme).
// IsAbsoluteURL은 URL이 절대 경로인지 확인합니다 (스키마가 있음).
func IsAbsoluteURL(urlStr string) bool {
	return strings.HasPrefix(urlStr, "http://") || strings.HasPrefix(urlStr, "https://")
}

// NormalizeURL normalizes a URL (removes trailing slash, etc.).
// NormalizeURL은 URL을 정규화합니다 (후행 슬래시 제거 등).
func NormalizeURL(urlStr string) string {
	urlStr = strings.TrimSpace(urlStr)
	urlStr = strings.TrimRight(urlStr, "/")
	return urlStr
}

// GetDomain returns the domain from a URL.
// GetDomain은 URL에서 도메인을 반환합니다.
func GetDomain(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}
	return u.Host, nil
}

// GetScheme returns the scheme from a URL (http, https, etc.).
// GetScheme은 URL에서 스키마를 반환합니다 (http, https 등).
func GetScheme(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}
	return u.Scheme, nil
}

// GetPath returns the path from a URL.
// GetPath는 URL에서 경로를 반환합니다.
func GetPath(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}
	return u.Path, nil
}
