package websvrutil

import (
	"context"
	"net/http"
	"strings"
	"sync"
)

// Router is the HTTP request router.
// Router는 HTTP 요청 라우터입니다.
type Router struct {
	// routes stores all registered routes by HTTP method
	// routes는 HTTP 메서드별로 등록된 모든 라우트를 저장합니다
	routes map[string][]*Route

	// notFoundHandler is called when no route matches
	// notFoundHandler는 일치하는 라우트가 없을 때 호출됩니다
	notFoundHandler http.HandlerFunc

	// mu protects concurrent access to routes
	// mu는 라우트에 대한 동시 액세스를 보호합니다
	mu sync.RWMutex
}

// Route represents a registered route.
// Route는 등록된 라우트를 나타냅니다.
type Route struct {
	// Method is the HTTP method (GET, POST, etc.)
	// Method는 HTTP 메서드입니다 (GET, POST 등)
	Method string

	// Pattern is the URL pattern (e.g., "/users/:id") / Pattern은 URL 패턴입니다 (예: "/users/:id")
	Pattern string

	// Handler is the function to call when the route matches
	// Handler는 라우트가 일치할 때 호출할 함수입니다
	Handler http.HandlerFunc

	// segments stores the parsed pattern segments
	// segments는 파싱된 패턴 세그먼트를 저장합니다
	segments []segment
}

// segment represents a single part of a URL pattern.
// segment는 URL 패턴의 단일 부분을 나타냅니다.
type segment struct {
	// value is the literal value or parameter name
	// value는 리터럴 값 또는 매개변수 이름입니다
	value string

	// isParam indicates if this segment is a parameter (starts with :)
	// isParam은 이 세그먼트가 매개변수인지 나타냅니다 (:로 시작)
	isParam bool

	// isWildcard indicates if this segment is a wildcard (*)
	// isWildcard는 이 세그먼트가 와일드카드인지 나타냅니다 (*)
	isWildcard bool
}

// newRouter creates a new Router instance.
// newRouter는 새 Router 인스턴스를 생성합니다.
func newRouter() *Router {
	return &Router{
		routes: make(map[string][]*Route),
		notFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		}),
	}
}

// Handle registers a new route with the given method, pattern, and handler.
// Handle은 주어진 메서드, 패턴 및 핸들러로 새 라우트를 등록합니다.
//
// Pattern syntax
// 패턴 구문:
//   - "/users" - exact match / 정확한 일치
//   - "/users/:id" - parameter (accessible via Context.Param("id")) / 매개변수
//   - "/files/*" - wildcard (matches everything after /files/) / 와일드카드
//
// Example
// 예제:
//
//	router.Handle("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
//	    // Handler implementation
//	})
func (ro *Router) Handle(method, pattern string, handler http.HandlerFunc) {
	ro.mu.Lock()
	defer ro.mu.Unlock()

	// Parse the pattern into segments
	// 패턴을 세그먼트로 파싱
	segments := parsePattern(pattern)

	// Create the route
	// 라우트 생성
	route := &Route{
		Method:   strings.ToUpper(method),
		Pattern:  pattern,
		Handler:  handler,
		segments: segments,
	}

	// Add to routes map
	// 라우트 맵에 추가
	ro.routes[route.Method] = append(ro.routes[route.Method], route)
}

// GET registers a GET route.
// GET은 GET 라우트를 등록합니다.
func (ro *Router) GET(pattern string, handler http.HandlerFunc) {
	ro.Handle("GET", pattern, handler)
}

// POST registers a POST route.
// POST는 POST 라우트를 등록합니다.
func (ro *Router) POST(pattern string, handler http.HandlerFunc) {
	ro.Handle("POST", pattern, handler)
}

// PUT registers a PUT route.
// PUT은 PUT 라우트를 등록합니다.
func (ro *Router) PUT(pattern string, handler http.HandlerFunc) {
	ro.Handle("PUT", pattern, handler)
}

// PATCH registers a PATCH route.
// PATCH는 PATCH 라우트를 등록합니다.
func (ro *Router) PATCH(pattern string, handler http.HandlerFunc) {
	ro.Handle("PATCH", pattern, handler)
}

// DELETE registers a DELETE route.
// DELETE는 DELETE 라우트를 등록합니다.
func (ro *Router) DELETE(pattern string, handler http.HandlerFunc) {
	ro.Handle("DELETE", pattern, handler)
}

// OPTIONS registers an OPTIONS route.
// OPTIONS는 OPTIONS 라우트를 등록합니다.
func (ro *Router) OPTIONS(pattern string, handler http.HandlerFunc) {
	ro.Handle("OPTIONS", pattern, handler)
}

// HEAD registers a HEAD route.
// HEAD는 HEAD 라우트를 등록합니다.
func (ro *Router) HEAD(pattern string, handler http.HandlerFunc) {
	ro.Handle("HEAD", pattern, handler)
}

// NotFound sets the handler for 404 Not Found responses.
// NotFound는 404 Not Found 응답에 대한 핸들러를 설정합니다.
func (ro *Router) NotFound(handler http.HandlerFunc) {
	ro.mu.Lock()
	defer ro.mu.Unlock()
	ro.notFoundHandler = handler
}

// ServeHTTP implements the http.Handler interface.
// ServeHTTP는 http.Handler 인터페이스를 구현합니다.
func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ro.mu.RLock()
	defer ro.mu.RUnlock()

	// Get routes for this method
	// 이 메서드의 라우트 가져오기
	routes := ro.routes[r.Method]

	// Try to match a route
	// 라우트 일치 시도
	for _, route := range routes {
		if params, ok := route.match(r.URL.Path); ok {
			// Create context with parameters
			// 매개변수와 함께 컨텍스트 생성
			ctx := NewContext(w, r)
			if len(params) > 0 {
				ctx.setParams(params)
			}

			// Store context in request context
			// 요청 컨텍스트에 컨텍스트 저장
			r = r.WithContext(contextWithValue(r.Context(), ctx))

			route.Handler(w, r)
			return
		}
	}

	// No route matched, call not found handler
	// 일치하는 라우트 없음, not found 핸들러 호출
	ro.notFoundHandler(w, r)
}

// contextWithValue stores the Context in the request's context.Context.
// contextWithValue는 요청의 context.Context에 Context를 저장합니다.
func contextWithValue(ctx context.Context, c *Context) context.Context {
	return context.WithValue(ctx, contextKeyParams, c)
}

// match checks if the route matches the given path and extracts parameters.
// match는 라우트가 주어진 경로와 일치하는지 확인하고 매개변수를 추출합니다.
//
// Matching algorithm
// 매칭 알고리즘:
//   1. Parse incoming path into segments
//   2. Check segment count compatibility:
//      - Without wildcard: must match exactly
//      - With wildcard: request must have at least (pattern segments - 1)
//   3. Iterate through pattern segments:
//      - Wildcard (*): captures all remaining path segments
//      - Parameter (:name): captures current segment into parameter map
//      - Literal: must match exactly (case-sensitive)
//   4. Return parameters map and match status
//
// Return values
// 반환 값:
//   - params: Map of parameter names to values (e.g., {"id": "123"})
//   - matched: true if route matches, false otherwise
//
// Examples
// 예제:
//   Pattern: "/users/:id", Path: "/users/123"
//   -> params: {"id": "123"}, matched: true
//
//   Pattern: "/files/*", Path: "/files/docs/report.pdf"
//   -> params: {"*": "docs/report.pdf"}, matched: true
//
//   Pattern: "/users/:id", Path: "/posts/123"
//   -> params: nil, matched: false
//
// Time complexity: O(n) where n = number of segments
// 시간 복잡도: O(n), n = 세그먼트 수
func (route *Route) match(path string) (map[string]string, bool) {
	// Parse the request path into segments
	// 요청 경로를 세그먼트로 파싱
	pathSegments := parsePath(path)

	// Check if segment counts match (unless there's a wildcard)
	// 세그먼트 수 일치 확인 (와일드카드가 없는 경우)
	hasWildcard := false
	for _, seg := range route.segments {
		if seg.isWildcard {
			hasWildcard = true
			break
		}
	}

	if !hasWildcard && len(pathSegments) != len(route.segments) {
		return nil, false
	}

	if hasWildcard && len(pathSegments) < len(route.segments)-1 {
		return nil, false
	}

	// Match each segment and extract parameters
	// 각 세그먼트를 일치시키고 매개변수 추출
	params := make(map[string]string)

	for i, segment := range route.segments {
		if segment.isWildcard {
			// Wildcard matches everything remaining
			// 와일드카드는 나머지 모든 것과 일치
			if i < len(pathSegments) {
				params["*"] = strings.Join(pathSegments[i:], "/")
			} else {
				params["*"] = ""
			}
			return params, true
		}

		if i >= len(pathSegments) {
			return nil, false
		}

		if segment.isParam {
			// Parameter segment, store the value
			// 매개변수 세그먼트, 값 저장
			params[segment.value] = pathSegments[i]
		} else {
			// Literal segment, must match exactly
			// 리터럴 세그먼트, 정확히 일치해야 함
			if segment.value != pathSegments[i] {
				return nil, false
			}
		}
	}

	return params, true
}

// parsePattern parses a URL pattern string into segments for efficient matching.
// parsePattern은 효율적인 매칭을 위해 URL 패턴 문자열을 세그먼트로 파싱합니다.
//
// Pattern syntax
// 패턴 구문:
// - Static segments: "/users/profile" - Exact match required / - 정적 세그먼트: "/users/profile" - 정확한 일치 필요
// - Parameters: "/users/:id" - Captures value into "id" parameter / - 매개변수: "/users/:id" - "id" 매개변수로 값 캡처
// - Wildcards: "/files/*" - Matches all remaining path segments / - 와일드카드: "/files/*" - 나머지 모든 경로 세그먼트 일치
//
// Implementation details
// 구현 세부사항:
//   1. Trim leading/trailing slashes from pattern
//   2. Split pattern by "/" separator
//   3. Identify segment types:
//      - "*" -> wildcard segment (isWildcard=true)
//      - ":name" -> parameter segment (isParam=true, value="name")
//      - "literal" -> static segment (value="literal")
//   4. Pre-allocate segments slice with capacity for performance
//
// Time complexity: O(n) where n = number of path segments
// 시간 복잡도: O(n), n = 경로 세그먼트 수
//
// Examples
// 예제:
//   parsePattern("/users") -> [segment{value: "users"}]
//   parsePattern("/users/:id") -> [segment{value: "users"}, segment{value: "id", isParam: true}]
//   parsePattern("/files/*") -> [segment{value: "files"}, segment{isWildcard: true}]
func parsePattern(pattern string) []segment {
	// Remove leading and trailing slashes
	// 앞뒤 슬래시 제거
	pattern = strings.Trim(pattern, "/")

	if pattern == "" {
		return []segment{}
	}

	// Split by
	// / 로 분할
	parts := strings.Split(pattern, "/")
	segments := make([]segment, 0, len(parts))

	for _, part := range parts {
		if part == "" {
			continue
		}

		seg := segment{}

		if part == "*" {
			// Wildcard segment
			// 와일드카드 세그먼트
			seg.isWildcard = true
		} else if strings.HasPrefix(part, ":") {
			// Parameter segment
			// 매개변수 세그먼트
			seg.isParam = true
			seg.value = part[1:] // Remove the : prefix / : 접두사 제거
		} else {
			// Literal segment
			// 리터럴 세그먼트
			seg.value = part
		}

		segments = append(segments, seg)
	}

	return segments
}

// parsePath parses a URL path into segments for route matching.
// parsePath는 라우트 매칭을 위해 URL 경로를 세그먼트로 파싱합니다.
//
// Process
// 프로세스:
//   1. Trim leading and trailing slashes
//   2. Handle empty path (returns empty slice)
//   3. Split by "/" separator
//   4. Filter out empty segments (from consecutive slashes)
//   5. Return cleaned segment list
//
// Examples
// 예제:
//   parsePath("/users/123") -> ["users", "123"]
//   parsePath("/users/123/") -> ["users", "123"] (trailing slash removed)
//   parsePath("/api//v1/users") -> ["api", "v1", "users"] (double slash cleaned)
//   parsePath("/") -> [] (root path)
//   parsePath("") -> [] (empty path)
//
// Time complexity: O(n) where n = path length
// 시간 복잡도: O(n), n = 경로 길이
func parsePath(path string) []string {
	// Remove leading and trailing slashes
	// 앞뒤 슬래시 제거
	path = strings.Trim(path, "/")

	if path == "" {
		return []string{}
	}

	// Split by
	// / 로 분할
	parts := strings.Split(path, "/")
	segments := make([]string, 0, len(parts))

	for _, part := range parts {
		if part != "" {
			segments = append(segments, part)
		}
	}

	return segments
}
