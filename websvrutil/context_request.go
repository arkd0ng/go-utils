package websvrutil

import (
	"net/http"
	"strings"
)

// context_request.go provides HTTP request information retrieval methods for the Context type.
//
// This file contains methods for accessing various aspects of incoming HTTP requests:
//
// Request Information:
//   - Method(), Path(): Basic request metadata
//   - Query(), QueryDefault(): URL query parameter access
//
// Headers:
//   - Header(), SetHeader(), AddHeader(): Request/response header manipulation
//   - GetHeader(), GetHeaders(), HeaderExists(): Header access utilities
//   - ContentType(), UserAgent(), Referer(): Common header shortcuts
//
// Client IP Detection:
//   - ClientIP(): Intelligent client IP extraction considering proxies
//     Priority: X-Forwarded-For → X-Real-IP → RemoteAddr
//     Handles reverse proxies (nginx, HAProxy, CloudFlare)
//     Security: Uses only first IP from proxy chains to prevent spoofing
//
// HTTP Method Checks:
//   - IsGET(), IsPOST(), IsPUT(), IsPATCH(), IsDELETE(), IsHEAD(), IsOPTIONS()
//     Boolean methods for quick method type checking
//
// Request Type Detection:
//   - IsAjax(): Detects AJAX/XMLHttpRequest calls
//   - IsWebSocket(): Identifies WebSocket upgrade requests
//
// Content Negotiation:
//   - AcceptsJSON(), AcceptsHTML(), AcceptsXML()
//     Checks Accept header for response format selection
//
// All methods operate on the Context's underlying http.Request and http.ResponseWriter,
// providing a convenient and type-safe API for request handling.
//
// context_request.go는 Context 타입을 위한 HTTP 요청 정보 조회 메서드를 제공합니다.
//
// 이 파일은 들어오는 HTTP 요청의 다양한 측면에 접근하는 메서드를 포함합니다:
//
// 요청 정보:
//   - Method(), Path(): 기본 요청 메타데이터
//   - Query(), QueryDefault(): URL 쿼리 매개변수 접근
//
// 헤더:
//   - Header(), SetHeader(), AddHeader(): 요청/응답 헤더 조작
//   - GetHeader(), GetHeaders(), HeaderExists(): 헤더 접근 유틸리티
//   - ContentType(), UserAgent(), Referer(): 일반 헤더 단축키
//
// 클라이언트 IP 감지:
//   - ClientIP(): 프록시를 고려한 지능형 클라이언트 IP 추출
//     우선순위: X-Forwarded-For → X-Real-IP → RemoteAddr
//     리버스 프록시 처리 (nginx, HAProxy, CloudFlare)
//     보안: 스푸핑 방지를 위해 프록시 체인에서 첫 번째 IP만 사용
//
// HTTP 메서드 확인:
//   - IsGET(), IsPOST(), IsPUT(), IsPATCH(), IsDELETE(), IsHEAD(), IsOPTIONS()
//     빠른 메서드 타입 확인을 위한 부울 메서드
//
// 요청 타입 감지:
//   - IsAjax(): AJAX/XMLHttpRequest 호출 감지
//   - IsWebSocket(): WebSocket 업그레이드 요청 식별
//
// 콘텐츠 협상:
//   - AcceptsJSON(), AcceptsHTML(), AcceptsXML()
//     응답 형식 선택을 위한 Accept 헤더 확인
//
// 모든 메서드는 Context의 기본 http.Request 및 http.ResponseWriter에서 작동하며,
// 요청 처리를 위한 편리하고 타입 안전한 API를 제공합니다.

// ============================================================================
// Request Information
// 요청 정보
// ============================================================================

// Method returns the HTTP request method.
// Method는 HTTP 요청 메서드를 반환합니다.
func (c *Context) Method() string {
	return c.Request.Method
}

// Path returns the URL path.
// Path는 URL 경로를 반환합니다.
func (c *Context) Path() string {
	return c.Request.URL.Path
}

// Query returns the query string parameter with the given name.
// Query는 주어진 이름의 쿼리 문자열 매개변수를 반환합니다.
//
// Example
// 예제:
//
// // URL: /search?q=golang&page=2
// q := ctx.Query("q")       // Returns "golang"
// "golang" 반환
//
// page := ctx.Query("page") // Returns "2"
// "2" 반환
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// QueryDefault returns the query string parameter with the given name,
// or the default value if it doesn't exist.
// QueryDefault는 주어진 이름의 쿼리 문자열 매개변수를 반환하거나,
// 존재하지 않으면 기본값을 반환합니다.
func (c *Context) QueryDefault(key, defaultValue string) string {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// ============================================================================
// Headers
// 헤더
// ============================================================================

// Header returns the request header with the given key.
// Header는 주어진 키의 요청 헤더를 반환합니다.
func (c *Context) Header(key string) string {
	return c.Request.Header.Get(key)
}

// SetHeader sets a response header.
// SetHeader는 응답 헤더를 설정합니다.
func (c *Context) SetHeader(key, value string) {
	c.ResponseWriter.Header().Set(key, value)
}

// AddHeader adds a header value to the response.
// AddHeader는 응답에 헤더 값을 추가합니다.
//
// Unlike SetHeader, this appends the value if the header already exists.
// SetHeader와 달리 헤더가 이미 존재하는 경우 값을 추가합니다.
//
// Example
// 예제:
//
//	ctx.AddHeader("Set-Cookie", "cookie1=value1")
//	ctx.AddHeader("Set-Cookie", "cookie2=value2")
func (c *Context) AddHeader(key, value string) {
	c.ResponseWriter.Header().Add(key, value)
}

// GetHeader returns the request header with the given key.
// GetHeader는 주어진 키의 요청 헤더를 반환합니다.
//
// This is an alias for Header() for consistency.
// 일관성을 위한 Header()의 별칭입니다.
//
// Example
// 예제:
//
//	userAgent := ctx.GetHeader("User-Agent")
func (c *Context) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

// GetHeaders returns all values for the given header key.
// GetHeaders는 주어진 헤더 키의 모든 값을 반환합니다.
//
// Example
// 예제:
//
//	acceptEncodings := ctx.GetHeaders("Accept-Encoding")
func (c *Context) GetHeaders(key string) []string {
	return c.Request.Header.Values(key)
}

// HeaderExists checks if a request header exists.
// HeaderExists는 요청 헤더가 존재하는지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.HeaderExists("Authorization") {
//	    // Process authentication
//	}
func (c *Context) HeaderExists(key string) bool {
	_, exists := c.Request.Header[key]
	return exists
}

// ContentType returns the Content-Type header of the request.
// ContentType은 요청의 Content-Type 헤더를 반환합니다.
//
// Example
// 예제:
//
//	contentType := ctx.ContentType()
func (c *Context) ContentType() string {
	return c.Request.Header.Get("Content-Type")
}

// UserAgent returns the User-Agent header of the request.
// UserAgent는 요청의 User-Agent 헤더를 반환합니다.
//
// Example
// 예제:
//
//	userAgent := ctx.UserAgent()
func (c *Context) UserAgent() string {
	return c.Request.Header.Get("User-Agent")
}

// Referer returns the Referer header of the request.
// Referer는 요청의 Referer 헤더를 반환합니다.
//
// Example
// 예제:
//
//	referer := ctx.Referer()
func (c *Context) Referer() string {
	return c.Request.Header.Get("Referer")
}

// ============================================================================
// Client IP
// 클라이언트 IP
// ============================================================================

// ClientIP returns the client IP address.
// ClientIP는 클라이언트 IP 주소를 반환합니다.
//
// Priority order for IP detection
// IP 감지 우선순위:
//
//  1. X-Forwarded-For header (first IP in comma-separated list)
//
//  2. X-Real-IP header
//
//  3. RemoteAddr (direct connection)
//
//  1. X-Forwarded-For 헤더 (쉼표로 구분된 목록의 첫 번째 IP)
//
//  2. X-Real-IP 헤더
//
//  3. RemoteAddr (직접 연결)
//
// Header details
// 헤더 세부정보:
//
// X-Forwarded-For:
// - Standard header set by proxies (nginx, HAProxy, CloudFlare, etc.)
// - 프록시(nginx, HAProxy, CloudFlare 등)가 설정하는 표준 헤더
// - Format: "client, proxy1, proxy2" (comma-separated chain)
// - 형식: "클라이언트, 프록시1, 프록시2" (쉼표로 구분된 체인)
// - Returns ONLY the first IP (original client) for security
// - 보안을 위해 첫 번째 IP(원본 클라이언트)만 반환
// - Example: "203.0.113.195, 70.41.3.18" → returns "203.0.113.195"
// - 예제: "203.0.113.195, 70.41.3.18" → "203.0.113.195" 반환
//
// X-Real-IP:
// - Non-standard but widely used by nginx reverse proxies
// - 비표준이지만 nginx 리버스 프록시에서 널리 사용
// - Contains single IP address (no chain)
// - 단일 IP 주소 포함 (체인 없음)
// - More reliable than X-Forwarded-For when available
// - 사용 가능한 경우 X-Forwarded-For보다 신뢰성 높음
//
// RemoteAddr:
// - Direct TCP connection source address
// - 직접 TCP 연결 소스 주소
// - Format: "IP:Port" (e.g., "192.168.1.100:54321")
// - 형식: "IP:포트" (예: "192.168.1.100:54321")
// - Returns IP only (strips port number)
// - IP만 반환 (포트 번호 제거)
// - Most reliable when no proxies involved
// - 프록시가 없을 때 가장 신뢰할 수 있음
//
// Security considerations
// 보안 고려사항:
// - X-Forwarded-For can be spoofed by malicious clients
// - X-Forwarded-For는 악의적인 클라이언트가 위조할 수 있음
// - Only use first IP to prevent proxy chain manipulation
// - 프록시 체인 조작을 방지하기 위해 첫 번째 IP만 사용
// - For critical security decisions, validate IP against trusted proxy list
// - 중요한 보안 결정의 경우 신뢰할 수 있는 프록시 목록에 대해 IP 검증
// - Consider implementing IP whitelist/blacklist if needed
// - 필요한 경우 IP 화이트리스트/블랙리스트 구현 고려
//
// Performance
// 성능:
// - Time complexity: O(n) where n = length of X-Forwarded-For or RemoteAddr
// - 시간 복잡도: O(n), n = X-Forwarded-For 또는 RemoteAddr 길이
// - Optimized with byte-by-byte comparison (faster than strings.Split)
// - 바이트 단위 비교로 최적화 (strings.Split보다 빠름)
// - No memory allocation for string operations
// - 문자열 작업에 메모리 할당 없음
//
// Example scenarios
// 시나리오 예제:
//
// Direct connection (no proxy):
//
//	RemoteAddr: "203.0.113.195:54321"
//	Returns: "203.0.113.195"
//
// Behind nginx reverse proxy:
//
//	X-Real-IP: "203.0.113.195"
//	RemoteAddr: "127.0.0.1:8080" (nginx internal)
//	Returns: "203.0.113.195" (from X-Real-IP)
//
// Behind CloudFlare CDN:
//
//	X-Forwarded-For: "203.0.113.195, 104.16.133.229"
//	RemoteAddr: "104.16.133.229:443" (CloudFlare IP)
//	Returns: "203.0.113.195" (first IP from X-Forwarded-For)
//
// Example usage
// 사용 예제:
//
//	ip := ctx.ClientIP()
//	if ip == "127.0.0.1" {
//	    // Local request
//	} else {
//	    // Remote request, log for rate limiting
//	    rateLimiter.Check(ip)
//	}
func (c *Context) ClientIP() string {
	// Check X-Forwarded-For header
	// X-Forwarded-For 헤더 확인
	if xff := c.Request.Header.Get("X-Forwarded-For"); xff != "" {
		// Return the first IP in the list
		// 목록의 첫 번째 IP 반환
		// Use strings.IndexByte for better performance
		// 더 나은 성능을 위해 strings.IndexByte 사용
		if idx := strings.IndexByte(xff, ','); idx != -1 {
			return strings.TrimSpace(xff[:idx])
		}
		return xff
	}

	// Check X-Real-IP header
	// X-Real-IP 헤더 확인
	if xri := c.Request.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Fall back to RemoteAddr
	// RemoteAddr로 대체
	// Use strings.IndexByte for better performance
	// 더 나은 성능을 위해 strings.IndexByte 사용
	if idx := strings.IndexByte(c.Request.RemoteAddr, ':'); idx != -1 {
		return c.Request.RemoteAddr[:idx]
	}
	return c.Request.RemoteAddr
}

// ============================================================================
// HTTP Method Checks
// HTTP 메서드 확인
// ============================================================================

// IsGET checks if the request method is GET.
// IsGET는 요청 메서드가 GET인지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.IsGET() {
//	    // Handle GET request
//	}
func (c *Context) IsGET() bool {
	return c.Request.Method == http.MethodGet
}

// IsPOST checks if the request method is POST.
// IsPOST는 요청 메서드가 POST인지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.IsPOST() {
//	    // Handle POST request
//	}
func (c *Context) IsPOST() bool {
	return c.Request.Method == http.MethodPost
}

// IsPUT checks if the request method is PUT.
// IsPUT는 요청 메서드가 PUT인지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.IsPUT() {
//	    // Handle PUT request
//	}
func (c *Context) IsPUT() bool {
	return c.Request.Method == http.MethodPut
}

// IsPATCH checks if the request method is PATCH.
// IsPATCH는 요청 메서드가 PATCH인지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.IsPATCH() {
//	    // Handle PATCH request
//	}
func (c *Context) IsPATCH() bool {
	return c.Request.Method == http.MethodPatch
}

// IsDELETE checks if the request method is DELETE.
// IsDELETE는 요청 메서드가 DELETE인지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.IsDELETE() {
//	    // Handle DELETE request
//	}
func (c *Context) IsDELETE() bool {
	return c.Request.Method == http.MethodDelete
}

// IsHEAD checks if the request method is HEAD.
// IsHEAD는 요청 메서드가 HEAD인지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.IsHEAD() {
//	    // Handle HEAD request
//	}
func (c *Context) IsHEAD() bool {
	return c.Request.Method == http.MethodHead
}

// IsOPTIONS checks if the request method is OPTIONS.
// IsOPTIONS는 요청 메서드가 OPTIONS인지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.IsOPTIONS() {
//	    // Handle OPTIONS request
//	}
func (c *Context) IsOPTIONS() bool {
	return c.Request.Method == http.MethodOptions
}

// ============================================================================
// Request Type Checks
// 요청 타입 확인
// ============================================================================

// IsAjax checks if the request is an AJAX request (XMLHttpRequest).
// IsAjax는 요청이 AJAX 요청(XMLHttpRequest)인지 확인합니다.
//
// It checks for the X-Requested-With header set to "XMLHttpRequest".
// X-Requested-With 헤더가 "XMLHttpRequest"로 설정되었는지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.IsAjax() {
//	    // Handle AJAX request
//	}
func (c *Context) IsAjax() bool {
	return c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

// IsWebSocket checks if the request is a WebSocket upgrade request.
// IsWebSocket는 요청이 WebSocket 업그레이드 요청인지 확인합니다.
//
// Example
// 예제:
//
//	if ctx.IsWebSocket() {
//	    // Handle WebSocket upgrade
//	}
func (c *Context) IsWebSocket() bool {
	upgrade := c.Request.Header.Get("Upgrade")
	return upgrade == "websocket"
}

// ============================================================================
// Accept Type Checks
// Accept 타입 확인
// ============================================================================

// AcceptsJSON checks if the client accepts JSON responses.
// AcceptsJSON은 클라이언트가 JSON 응답을 수락하는지 확인합니다.
//
// It checks the Accept header for "application/json".
// Accept 헤더에서 "application/json"을 확인합니다.
//
// Example
// 예제:
//
//	if ctx.AcceptsJSON() {
//	    ctx.JSON(http.StatusOK, data)
//	}
func (c *Context) AcceptsJSON() bool {
	accept := c.Request.Header.Get("Accept")
	return accept == "*/*" ||
		accept == "application/json" ||
		c.containsContentType(accept, "application/json")
}

// AcceptsHTML checks if the client accepts HTML responses.
// AcceptsHTML은 클라이언트가 HTML 응답을 수락하는지 확인합니다.
//
// It checks the Accept header for "text/html".
// Accept 헤더에서 "text/html"을 확인합니다.
//
// Example
// 예제:
//
//	if ctx.AcceptsHTML() {
//	    ctx.HTML(http.StatusOK, "index", data)
//	}
func (c *Context) AcceptsHTML() bool {
	accept := c.Request.Header.Get("Accept")
	return accept == "*/*" ||
		accept == "text/html" ||
		c.containsContentType(accept, "text/html")
}

// AcceptsXML checks if the client accepts XML responses.
// AcceptsXML은 클라이언트가 XML 응답을 수락하는지 확인합니다.
//
// It checks the Accept header for "application/xml" or "text/xml".
// Accept 헤더에서 "application/xml" 또는 "text/xml"을 확인합니다.
//
// Example
// 예제:
//
//	if ctx.AcceptsXML() {
//	    ctx.XML(http.StatusOK, data)
//	}
func (c *Context) AcceptsXML() bool {
	accept := c.Request.Header.Get("Accept")
	return accept == "*/*" ||
		accept == "application/xml" ||
		accept == "text/xml" ||
		c.containsContentType(accept, "application/xml") ||
		c.containsContentType(accept, "text/xml")
}

// containsContentType checks if the accept header contains a specific content type.
// containsContentType는 accept 헤더에 특정 콘텐츠 타입이 포함되어 있는지 확인합니다.
func (c *Context) containsContentType(accept, contentType string) bool {
	// Simple substring check for content type
	// 콘텐츠 타입에 대한 간단한 부분 문자열 확인
	for i := 0; i < len(accept); i++ {
		if i+len(contentType) <= len(accept) && accept[i:i+len(contentType)] == contentType {
			return true
		}
	}
	return false
}
