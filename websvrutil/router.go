package websvrutil

import (
	"context"
	"net/http"
	"strings"
	"sync"
)

// Router is the HTTP request router that matches incoming requests to registered route handlers.
// This is the core routing engine used by App to dispatch HTTP requests to appropriate handlers
// based on HTTP method and URL path patterns with support for parameters and wildcards.
//
// Router provides:
// - Method-based routing (GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD)
// - URL pattern matching with parameters (:id) and wildcards (*)
// - Efficient route lookup with O(n) complexity where n = number of routes for method
// - Thread-safe concurrent route registration and request handling
// - Custom 404 Not Found handler support
// - Zero external dependencies (uses only Go standard library)
//
// Routing Features:
// - Static routes: Exact path matching (e.g., "/users", "/api/v1/status")
// - Named parameters: Capture path segments (e.g., "/users/:id" → id=123)
// - Wildcard routes: Match remaining path (e.g., "/files/*" → */docs/file.pdf)
// - Case-sensitive matching: "/Users" and "/users" are different routes
// - Automatic parameter extraction and storage in request context
// - Support for all standard HTTP methods
//
// Thread-Safety:
// - Fully thread-safe for concurrent use
// - Read-write mutex protects route registration and lookup
// - Multiple requests can be handled concurrently (read locks)
// - Route registration blocks request handling briefly (write locks)
//
// Performance Characteristics:
// - Route lookup: O(n) where n = number of routes for HTTP method
// - Pattern parsing: O(m) where m = number of path segments (done once at registration)
// - Memory: Minimal overhead, stores only registered routes
// - No route compilation or preprocessing beyond initial pattern parsing
// - Suitable for applications with hundreds of routes
//
// Router는 등록된 라우트 핸들러에 들어오는 요청을 매칭하는 HTTP 요청 라우터입니다.
// 이것은 App이 HTTP 메서드와 URL 경로 패턴을 기반으로 HTTP 요청을
// 적절한 핸들러로 디스패치하는 데 사용하는 핵심 라우팅 엔진이며,
// 매개변수와 와일드카드를 지원합니다.
//
// Router는 다음을 제공합니다:
// - 메서드 기반 라우팅(GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD)
// - 매개변수(:id) 및 와일드카드(*)가 있는 URL 패턴 매칭
// - O(n) 복잡도의 효율적인 라우트 조회(n = 메서드의 라우트 수)
// - 스레드 안전 동시 라우트 등록 및 요청 처리
// - 커스텀 404 Not Found 핸들러 지원
// - 외부 종속성 없음(Go 표준 라이브러리만 사용)
//
// 라우팅 기능:
// - 정적 라우트: 정확한 경로 매칭(예: "/users", "/api/v1/status")
// - 명명된 매개변수: 경로 세그먼트 캡처(예: "/users/:id" → id=123)
// - 와일드카드 라우트: 나머지 경로 일치(예: "/files/*" → */docs/file.pdf)
// - 대소문자 구분 매칭: "/Users"와 "/users"는 다른 라우트
// - 자동 매개변수 추출 및 요청 컨텍스트에 저장
// - 모든 표준 HTTP 메서드 지원
//
// 스레드 안전성:
// - 동시 사용에 완전히 스레드 안전
// - 읽기-쓰기 뮤텍스가 라우트 등록 및 조회 보호
// - 여러 요청을 동시에 처리할 수 있음(읽기 잠금)
// - 라우트 등록은 요청 처리를 잠깐 차단(쓰기 잠금)
//
// 성능 특성:
// - 라우트 조회: O(n), n = HTTP 메서드의 라우트 수
// - 패턴 파싱: O(m), m = 경로 세그먼트 수(등록 시 한 번만 수행)
// - 메모리: 최소 오버헤드, 등록된 라우트만 저장
// - 초기 패턴 파싱 이외에는 라우트 컴파일이나 전처리 없음
// - 수백 개의 라우트를 가진 애플리케이션에 적합
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

// Route represents a registered HTTP route with pattern matching capabilities.
// Each Route instance stores the HTTP method, URL pattern, handler function,
// and pre-parsed pattern segments for efficient request matching.
//
// Route contains:
// - HTTP method specification (GET, POST, PUT, etc.)
// - URL pattern with support for literals, parameters, and wildcards
// - Handler function to execute when route matches
// - Pre-parsed segments for O(1) pattern compilation and O(n) matching
//
// Pattern Syntax:
// - Literal segments: "/users/profile" - Must match exactly
// - Parameter segments: "/users/:id" - Captures value into named parameter
// - Wildcard segments: "/files/*" - Matches all remaining path components
// - Mixed patterns: "/api/:version/users/:id" - Multiple parameters allowed
//
// Matching Behavior:
// - Case-sensitive path matching
// - Parameter values are extracted and stored in Context
// - Wildcard captures everything after its position as single parameter "*"
// - First matching route wins (routes checked in registration order)
// - Failed matches return nil params and false status
//
// Performance:
// - Pattern parsing done once at registration (O(m) where m = segments)
// - Match checking is O(n) where n = number of segments in pattern
// - No regular expressions used (simple string comparison)
// - Minimal memory overhead per route
//
// Thread-Safety:
// - Route instances are read-only after creation
// - Safe for concurrent request matching
// - Handler function should be thread-safe
//
// Route는 패턴 매칭 기능을 가진 등록된 HTTP 라우트를 나타냅니다.
// 각 Route 인스턴스는 HTTP 메서드, URL 패턴, 핸들러 함수,
// 그리고 효율적인 요청 매칭을 위해 사전 파싱된 패턴 세그먼트를 저장합니다.
//
// Route는 다음을 포함합니다:
// - HTTP 메서드 사양(GET, POST, PUT 등)
// - 리터럴, 매개변수, 와일드카드를 지원하는 URL 패턴
// - 라우트가 일치할 때 실행할 핸들러 함수
// - O(1) 패턴 컴파일 및 O(n) 매칭을 위한 사전 파싱된 세그먼트
//
// 패턴 구문:
// - 리터럴 세그먼트: "/users/profile" - 정확히 일치해야 함
// - 매개변수 세그먼트: "/users/:id" - 명명된 매개변수로 값 캡처
// - 와일드카드 세그먼트: "/files/*" - 나머지 모든 경로 구성 요소 일치
// - 혼합 패턴: "/api/:version/users/:id" - 여러 매개변수 허용
//
// 매칭 동작:
// - 대소문자 구분 경로 매칭
// - 매개변수 값은 추출되어 Context에 저장됨
// - 와일드카드는 그 위치 이후의 모든 것을 단일 매개변수 "*"로 캡처
// - 처음 일치하는 라우트가 승리(등록 순서대로 라우트 확인)
// - 실패한 매칭은 nil params 및 false 상태 반환
//
// 성능:
// - 등록 시 패턴 파싱이 한 번 수행됨(O(m), m = 세그먼트 수)
// - 매칭 확인은 O(n), n = 패턴의 세그먼트 수
// - 정규 표현식 미사용(단순 문자열 비교)
// - 라우트당 최소 메모리 오버헤드
//
// 스레드 안전성:
// - Route 인스턴스는 생성 후 읽기 전용
// - 동시 요청 매칭에 안전
// - 핸들러 함수는 스레드 안전해야 함
type Route struct {
	// Method is the HTTP method (GET, POST, etc.)
	// Method는 HTTP 메서드입니다 (GET, POST 등)
	Method string

	// Pattern is the URL pattern (e.g., "/users/:id")
	// Pattern은 URL 패턴입니다 (예: "/users/:id")
	Pattern string

	// Handler is the function to call when the route matches
	// Handler는 라우트가 일치할 때 호출할 함수입니다
	Handler http.HandlerFunc

	// segments stores the parsed pattern segments
	// segments는 파싱된 패턴 세그먼트를 저장합니다
	segments []segment
}

// segment represents a single parsed component of a URL pattern for efficient route matching.
// Each segment identifies whether it's a literal string, a named parameter, or a wildcard,
// enabling fast pattern matching without regular expressions.
//
// Segment Types:
// 1. Literal Segment (default):
//   - value: "users", "api", "v1"
//   - isParam: false, isWildcard: false
//   - Matches only exact string value
//   - Example: pattern "/users/list" → segments: [{value:"users"}, {value:"list"}]
//
// 2. Parameter Segment (starts with : in pattern):
//   - value: parameter name without colon (e.g., "id", "name")
//   - isParam: true, isWildcard: false
//   - Matches any single path segment
//   - Captured value stored in params map with parameter name as key
//   - Example: pattern "/users/:id" → segments: [{value:"users"}, {value:"id", isParam:true}]
//
// 3. Wildcard Segment (* in pattern):
//   - value: empty or unused
//   - isParam: false, isWildcard: true
//   - Matches all remaining path segments
//   - Captured value stored in params map with key "*"
//   - Must be last segment in pattern (anything after is ignored)
//   - Example: pattern "/files/*" → segments: [{value:"files"}, {isWildcard:true}]
//
// Parsing Rules:
// - Segments are created by parsePattern() function
// - Pattern is split by "/" separator
// - Empty segments (from consecutive slashes) are filtered out
// - Leading/trailing slashes are removed before splitting
// - Segment type determined by first character:
//   - "*" → wildcard segment
//   - ":" → parameter segment (colon stripped, name stored)
//   - other → literal segment
//
// Matching Behavior:
// - Literal: Compared with == operator (case-sensitive)
// - Parameter: Always matches, value extracted to params map
// - Wildcard: Always matches, captures remaining path as single value
//
// Memory Layout:
// - Small memory footprint (16-24 bytes per segment)
// - value: string (16 bytes on 64-bit)
// - isParam: bool (1 byte, padded)
// - isWildcard: bool (1 byte, padded)
//
// Performance:
// - Segment type checks are simple boolean comparisons (O(1))
// - String comparison for literals is O(k) where k = segment length
// - No allocations during matching (params map allocated once)
//
// Examples:
//
//	Pattern: "/api/v1/users/:id/posts/:postId"
//	Segments: [
//	  {value:"api"},
//	  {value:"v1"},
//	  {value:"users"},
//	  {value:"id", isParam:true},
//	  {value:"posts"},
//	  {value:"postId", isParam:true}
//	]
//
//	Pattern: "/static/*"
//	Segments: [
//	  {value:"static"},
//	  {isWildcard:true}
//	]
//
// segment는 효율적인 라우트 매칭을 위한 URL 패턴의 단일 파싱된 구성 요소를 나타냅니다.
// 각 세그먼트는 리터럴 문자열, 명명된 매개변수 또는 와일드카드인지를 식별하여
// 정규 표현식 없이 빠른 패턴 매칭을 가능하게 합니다.
//
// 세그먼트 유형:
// 1. 리터럴 세그먼트(기본값):
//   - value: "users", "api", "v1"
//   - isParam: false, isWildcard: false
//   - 정확한 문자열 값만 일치
//   - 예: 패턴 "/users/list" → 세그먼트: [{value:"users"}, {value:"list"}]
//
// 2. 매개변수 세그먼트(패턴에서 :로 시작):
//   - value: 콜론 없는 매개변수 이름(예: "id", "name")
//   - isParam: true, isWildcard: false
//   - 모든 단일 경로 세그먼트와 일치
//   - 캡처된 값은 매개변수 이름을 키로 params 맵에 저장됨
//   - 예: 패턴 "/users/:id" → 세그먼트: [{value:"users"}, {value:"id", isParam:true}]
//
// 3. 와일드카드 세그먼트(패턴의 *):
//   - value: 비어있거나 사용되지 않음
//   - isParam: false, isWildcard: true
//   - 나머지 모든 경로 세그먼트와 일치
//   - 캡처된 값은 키 "*"로 params 맵에 저장됨
//   - 패턴의 마지막 세그먼트여야 함(이후의 모든 것은 무시됨)
//   - 예: 패턴 "/files/*" → 세그먼트: [{value:"files"}, {isWildcard:true}]
//
// 파싱 규칙:
// - 세그먼트는 parsePattern() 함수에 의해 생성됨
// - 패턴은 "/" 구분자로 분할됨
// - 빈 세그먼트(연속된 슬래시에서)는 필터링됨
// - 분할 전 앞뒤 슬래시가 제거됨
// - 세그먼트 유형은 첫 번째 문자로 결정:
//   - "*" → 와일드카드 세그먼트
//   - ":" → 매개변수 세그먼트(콜론 제거, 이름 저장)
//   - 기타 → 리터럴 세그먼트
//
// 매칭 동작:
// - 리터럴: == 연산자로 비교(대소문자 구분)
// - 매개변수: 항상 일치, 값을 params 맵으로 추출
// - 와일드카드: 항상 일치, 나머지 경로를 단일 값으로 캡처
//
// 성능:
// - 세그먼트 유형 확인은 단순 불린 비교(O(1))
// - 리터럴의 문자열 비교는 O(k), k = 세그먼트 길이
// - 매칭 중 할당 없음(params 맵은 한 번 할당)
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

// newRouter creates and initializes a new Router instance for HTTP request routing.
// This is an internal constructor used by App to create its routing engine.
// Applications should not call this directly; use websvrutil.New() instead.
//
// Initialization:
// - Creates empty routes map (map[string][]*Route)
// - Sets default 404 Not Found handler (http.NotFound)
// - No routes registered initially
// - Ready for concurrent route registration via Handle() method
//
// Default 404 Handler:
// - Returns standard HTTP 404 Not Found response
// - Uses http.NotFound() from standard library
// - Sets "Content-Type: text/plain; charset=utf-8" header
// - Writes "404 page not found\n" message
// - Can be customized via Router.NotFound() method
//
// Routes Map Structure:
//
//	map[string][]*Route where:
//	- Key: HTTP method string (uppercase: "GET", "POST", etc.)
//	- Value: Slice of Route pointers for that method
//	- Routes within slice are checked in registration order
//	- First matching route wins
//
// Memory:
// - Minimal allocation (empty map + one closure for 404 handler)
// - Routes map grows as routes are registered
// - No pre-allocation of method keys
//
// Thread-Safety:
// - Returned Router is thread-safe
// - Internal mutex (sync.RWMutex) protects concurrent access
// - Safe to register routes and serve requests concurrently
//
// Usage:
//
//	Internal use only:
//	  router := newRouter()
//
//	Application use:
//	  app := websvrutil.New()  // Creates router internally
//	  app.GET("/path", handler)
//
// newRouter는 HTTP 요청 라우팅을 위한 새 Router 인스턴스를 생성하고 초기화합니다.
// 이것은 App이 라우팅 엔진을 생성하기 위해 사용하는 내부 생성자입니다.
// 애플리케이션은 이것을 직접 호출해서는 안 되며, 대신 websvrutil.New()를 사용하세요.
//
// 초기화:
// - 빈 라우트 맵 생성(map[string][]*Route)
// - 기본 404 Not Found 핸들러 설정(http.NotFound)
// - 초기에는 등록된 라우트 없음
// - Handle() 메서드를 통한 동시 라우트 등록 준비 완료
//
// 기본 404 핸들러:
// - 표준 HTTP 404 Not Found 응답 반환
// - 표준 라이브러리의 http.NotFound() 사용
// - "Content-Type: text/plain; charset=utf-8" 헤더 설정
// - "404 page not found\n" 메시지 작성
// - Router.NotFound() 메서드를 통해 커스터마이즈 가능
//
// 라우트 맵 구조:
//
//	map[string][]*Route, 여기서:
//	- 키: HTTP 메서드 문자열(대문자: "GET", "POST" 등)
//	- 값: 해당 메서드에 대한 Route 포인터 슬라이스
//	- 슬라이스 내의 라우트는 등록 순서대로 확인됨
//	- 처음 일치하는 라우트가 승리
//
// 메모리:
// - 최소 할당(빈 맵 + 404 핸들러를 위한 하나의 클로저)
// - 라우트가 등록됨에 따라 라우트 맵 증가
// - 메서드 키의 사전 할당 없음
//
// 스레드 안전성:
// - 반환된 Router는 스레드 안전
// - 내부 뮤텍스(sync.RWMutex)가 동시 액세스 보호
// - 라우트 등록과 요청 서빙을 동시에 안전하게 수행
func newRouter() *Router {
	return &Router{
		routes: make(map[string][]*Route),
		notFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		}),
	}
}

// Handle registers a new HTTP route with the specified method, pattern, and handler function.
// This is the core route registration method used by all HTTP method shortcuts (GET, POST, etc.).
// Routes are matched in registration order; first match wins.
//
// Handle is the foundation of the routing system, providing:
// - Method-specific route registration (GET, POST, PUT, etc.)
// - Pattern parsing with parameters and wildcards
// - Handler association for matched routes
// - Thread-safe route registration for concurrent use
//
// Purpose:
//   - Registers a route that maps HTTP method + URL pattern to handler
//   - Parses pattern into segments for efficient matching
//   - Stores route in method-specific list for fast lookup
//   - Enables RESTful API design with method-based routing
//
// Parameters:
//
//   - method: HTTP method string (case-insensitive, converted to uppercase).
//     Standard methods: "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"
//     Custom methods: Any string is accepted (e.g., "PROPFIND", "LOCK", "UNLOCK")
//     Empty string: Valid but not recommended (matches empty method requests)
//
//   - pattern: URL pattern string with optional parameters and wildcards.
//     Pattern Syntax:
//
//   - Literal paths: "/users", "/api/v1/status"
//
//   - Must match exactly (case-sensitive)
//
//   - Leading/trailing slashes are trimmed during parsing
//
//   - Named parameters: "/users/:id", "/posts/:postId/comments/:commentId"
//
//   - Starts with colon (:)
//
//   - Matches any single path segment
//
//   - Captured value accessible via Context.Param(name)
//
//   - Parameter name follows colon (alphanumeric + underscore recommended)
//
//   - Wildcards: "/files/*", "/static/*"
//
//   - Single asterisk (*)
//
//   - Matches all remaining path segments
//
//   - Must be last segment (anything after is ignored)
//
//   - Captured value accessible via Context.Param("*")
//
//   - Mixed patterns: "/api/:version/users/:id", "/docs/:lang/*"
//
//   - Combine literals, parameters, and wildcards
//
//   - Parameters can appear multiple times
//
//   - Wildcard (if present) must be last
//
//     Pattern Examples:
//
//   - "/users" → matches: "/users" only
//
//   - "/users/:id" → matches: "/users/123", "/users/abc"
//
//   - "/users/:id/posts" → matches: "/users/123/posts"
//
//   - "/files/*" → matches: "/files/a", "/files/a/b", "/files/a/b/c"
//
//   - "/:lang/docs/*" → matches: "/en/docs/guide", "/fr/docs/api/v1"
//
//   - handler: http.HandlerFunc to execute when route matches.
//     Handler receives http.ResponseWriter and *http.Request
//     Should write response status, headers, and body
//     Has access to route parameters via Context (extract from request.Context())
//
// Thread-Safety:
//   - Fully thread-safe with write lock (mu.Lock)
//   - Blocks concurrent Handle() calls during registration
//   - Blocks ServeHTTP() route lookup briefly during registration
//   - Safe to register routes while serving requests (though not recommended)
//
// Registration Behavior:
//   - Routes are appended to method-specific list
//   - Registration order determines matching priority
//   - No duplicate detection; same pattern can be registered multiple times
//   - Later registrations shadow earlier ones (first match wins)
//   - Pattern parsing happens once at registration (compiled for fast matching)
//
// Pattern Parsing:
//   - Pattern is parsed into segments immediately
//   - Parsing happens once (O(m) where m = number of segments)
//   - Segments stored in Route for O(n) matching during requests
//   - No regular expressions used (simple string operations)
//
// Performance:
//   - Registration: O(m) where m = pattern segment count
//   - Memory: One Route struct + segments slice per registration
//   - Lookup: O(n) where n = number of routes for method
//   - No route compilation or code generation
//
// Common Use Cases:
//   - RESTful API endpoints: Handle("GET", "/api/users", listUsers)
//   - Resource-specific actions: Handle("POST", "/users/:id/activate", activate)
//   - File serving: Handle("GET", "/static/*", serveFiles)
//   - Custom HTTP methods: Handle("PROPFIND", "/dav/*", webdavPropfind)
//   - Health checks: Handle("GET", "/health", healthCheck)
//
// Example - Basic Route Registration:
//
//	router := newRouter()
//	router.Handle("GET", "/users", func(w http.ResponseWriter, r *http.Request) {
//	    w.Write([]byte("User list"))
//	})
//
// Example - Route with Parameters:
//
//	router.Handle("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
//	    ctx := GetContext(r)
//	    id := ctx.Param("id")
//	    fmt.Fprintf(w, "User ID: %s", id)
//	})
//
// Example - Route with Wildcard:
//
//	router.Handle("GET", "/files/*", func(w http.ResponseWriter, r *http.Request) {
//	    ctx := GetContext(r)
//	    filepath := ctx.Param("*")
//	    fmt.Fprintf(w, "File path: %s", filepath)
//	})
//
// Example - Multiple Parameters:
//
//	router.Handle("GET", "/api/:version/users/:id", func(w http.ResponseWriter, r *http.Request) {
//	    ctx := GetContext(r)
//	    version := ctx.Param("version")
//	    id := ctx.Param("id")
//	    fmt.Fprintf(w, "API %s, User %s", version, id)
//	})
//
// Example - Custom HTTP Method:
//
//	router.Handle("PROPFIND", "/dav/*", func(w http.ResponseWriter, r *http.Request) {
//	    // WebDAV PROPFIND implementation
//	    w.WriteHeader(207) // Multi-Status
//	})
//
// Method Shortcuts:
//
//	Instead of calling Handle directly, use method shortcuts:
//	- router.GET(pattern, handler) → Handle("GET", pattern, handler)
//	- router.POST(pattern, handler) → Handle("POST", pattern, handler)
//	- router.PUT(pattern, handler) → Handle("PUT", pattern, handler)
//	- router.PATCH(pattern, handler) → Handle("PATCH", pattern, handler)
//	- router.DELETE(pattern, handler) → Handle("DELETE", pattern, handler)
//	- router.OPTIONS(pattern, handler) → Handle("OPTIONS", pattern, handler)
//	- router.HEAD(pattern, handler) → Handle("HEAD", pattern, handler)
//
// Best Practices:
//   - Register routes during application initialization, not per-request
//   - Use method shortcuts (GET, POST) for standard methods
//   - Keep patterns simple and RESTful
//   - Document route parameters in handler comments
//   - Validate parameter values in handlers
//   - Use consistent naming for parameters across routes
//   - Test route matching thoroughly
//
// Limitations:
//   - No route groups or nested routers
//   - No automatic parameter validation
//   - No route naming or URL generation
//   - No regex pattern support
//   - Wildcard must be last segment
//   - No per-route middleware (use App-level middleware)
//
// Handle은 지정된 메서드, 패턴 및 핸들러 함수로 새 HTTP 라우트를 등록합니다.
// 이것은 모든 HTTP 메서드 단축키(GET, POST 등)에서 사용하는 핵심 라우트 등록 메서드입니다.
// 라우트는 등록 순서대로 매칭되며, 첫 번째 일치가 승리합니다.
//
// Handle은 라우팅 시스템의 기반이며 다음을 제공합니다:
// - 메서드별 라우트 등록(GET, POST, PUT 등)
// - 매개변수 및 와일드카드가 있는 패턴 파싱
// - 매칭된 라우트에 대한 핸들러 연결
// - 동시 사용을 위한 스레드 안전 라우트 등록
//
// 목적:
//   - HTTP 메서드 + URL 패턴을 핸들러에 매핑하는 라우트 등록
//   - 효율적인 매칭을 위해 패턴을 세그먼트로 파싱
//   - 빠른 조회를 위해 메서드별 목록에 라우트 저장
//   - 메서드 기반 라우팅으로 RESTful API 디자인 활성화
//
// 매개변수:
//
//   - method: HTTP 메서드 문자열(대소문자 구분 없음, 대문자로 변환됨).
//     표준 메서드: "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"
//     커스텀 메서드: 모든 문자열 허용(예: "PROPFIND", "LOCK", "UNLOCK")
//     빈 문자열: 유효하지만 권장하지 않음(빈 메서드 요청과 일치)
//
//   - pattern: 선택적 매개변수 및 와일드카드가 있는 URL 패턴 문자열.
//     패턴 구문:
//
//   - 리터럴 경로: "/users", "/api/v1/status"
//
//   - 정확히 일치해야 함(대소문자 구분)
//
//   - 파싱 중 앞뒤 슬래시 제거
//
//   - 명명된 매개변수: "/users/:id", "/posts/:postId/comments/:commentId"
//
//   - 콜론(:)으로 시작
//
//   - 모든 단일 경로 세그먼트와 일치
//
//   - 캡처된 값은 Context.Param(name)으로 액세스 가능
//
//   - 매개변수 이름은 콜론 다음에 위치(영숫자 + 밑줄 권장)
//
//   - 와일드카드: "/files/*", "/static/*"
//
//   - 단일 별표(*)
//
//   - 나머지 모든 경로 세그먼트와 일치
//
//   - 마지막 세그먼트여야 함(이후의 모든 것은 무시됨)
//
//   - 캡처된 값은 Context.Param("*")로 액세스 가능
//
//   - 혼합 패턴: "/api/:version/users/:id", "/docs/:lang/*"
//
//   - 리터럴, 매개변수, 와일드카드 결합
//
//   - 매개변수는 여러 번 나타날 수 있음
//
//   - 와일드카드(있는 경우)는 마지막이어야 함
//
//   - handler: 라우트가 일치할 때 실행할 http.HandlerFunc.
//     핸들러는 http.ResponseWriter 및 *http.Request를 받음
//     응답 상태, 헤더 및 본문을 작성해야 함
//     Context를 통해 라우트 매개변수에 액세스(request.Context()에서 추출)
//
// 스레드 안전성:
//   - 쓰기 잠금(mu.Lock)으로 완전히 스레드 안전
//   - 등록 중 동시 Handle() 호출 차단
//   - 등록 중 ServeHTTP() 라우트 조회를 잠깐 차단
//   - 요청을 서빙하는 동안 라우트를 등록하는 것은 안전하지만 권장하지 않음
//
// 등록 동작:
//   - 라우트는 메서드별 목록에 추가됨
//   - 등록 순서가 매칭 우선순위 결정
//   - 중복 감지 없음; 동일한 패턴을 여러 번 등록할 수 있음
//   - 나중 등록이 이전 등록을 가림(첫 번째 일치가 승리)
//   - 패턴 파싱은 등록 시 한 번 발생(빠른 매칭을 위해 컴파일)
//
// 모범 사례:
//   - 요청당이 아닌 애플리케이션 초기화 중에 라우트 등록
//   - 표준 메서드에는 메서드 단축키(GET, POST) 사용
//   - 패턴을 간단하고 RESTful하게 유지
//   - 핸들러 주석에 라우트 매개변수 문서화
//   - 핸들러에서 매개변수 값 검증
//   - 라우트 전체에서 매개변수에 일관된 이름 사용
//   - 라우트 매칭을 철저히 테스트
//
// 제한사항:
//   - 라우트 그룹 또는 중첩 라우터 없음
//   - 자동 매개변수 검증 없음
//   - 라우트 이름 지정 또는 URL 생성 없음
//   - 정규식 패턴 지원 없음
//   - 와일드카드는 마지막 세그먼트여야 함
//   - 라우트별 미들웨어 없음(App 레벨 미들웨어 사용)
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

// GET registers a route for HTTP GET requests.
// Convenience method that calls Handle("GET", pattern, handler).
//
// GET은 HTTP GET 요청에 대한 라우트를 등록합니다.
// Handle("GET", pattern, handler)를 호출하는 편의 메서드입니다.
func (ro *Router) GET(pattern string, handler http.HandlerFunc) {
	ro.Handle("GET", pattern, handler)
}

// POST registers a route for HTTP POST requests.
// Convenience method that calls Handle("POST", pattern, handler).
//
// POST는 HTTP POST 요청에 대한 라우트를 등록합니다.
// Handle("POST", pattern, handler)를 호출하는 편의 메서드입니다.
func (ro *Router) POST(pattern string, handler http.HandlerFunc) {
	ro.Handle("POST", pattern, handler)
}

// PUT registers a route for HTTP PUT requests.
// Convenience method that calls Handle("PUT", pattern, handler).
//
// PUT은 HTTP PUT 요청에 대한 라우트를 등록합니다.
// Handle("PUT", pattern, handler)를 호출하는 편의 메서드입니다.
func (ro *Router) PUT(pattern string, handler http.HandlerFunc) {
	ro.Handle("PUT", pattern, handler)
}

// PATCH registers a route for HTTP PATCH requests.
// Convenience method that calls Handle("PATCH", pattern, handler).
//
// PATCH는 HTTP PATCH 요청에 대한 라우트를 등록합니다.
// Handle("PATCH", pattern, handler)를 호출하는 편의 메서드입니다.
func (ro *Router) PATCH(pattern string, handler http.HandlerFunc) {
	ro.Handle("PATCH", pattern, handler)
}

// DELETE registers a route for HTTP DELETE requests.
// Convenience method that calls Handle("DELETE", pattern, handler).
//
// DELETE는 HTTP DELETE 요청에 대한 라우트를 등록합니다.
// Handle("DELETE", pattern, handler)를 호출하는 편의 메서드입니다.
func (ro *Router) DELETE(pattern string, handler http.HandlerFunc) {
	ro.Handle("DELETE", pattern, handler)
}

// OPTIONS registers a route for HTTP OPTIONS requests.
// Convenience method that calls Handle("OPTIONS", pattern, handler).
//
// OPTIONS는 HTTP OPTIONS 요청에 대한 라우트를 등록합니다.
// Handle("OPTIONS", pattern, handler)를 호출하는 편의 메서드입니다.
func (ro *Router) OPTIONS(pattern string, handler http.HandlerFunc) {
	ro.Handle("OPTIONS", pattern, handler)
}

// HEAD registers a route for HTTP HEAD requests.
// Convenience method that calls Handle("HEAD", pattern, handler).
//
// HEAD는 HTTP HEAD 요청에 대한 라우트를 등록합니다.
// Handle("HEAD", pattern, handler)를 호출하는 편의 메서드입니다.
func (ro *Router) HEAD(pattern string, handler http.HandlerFunc) {
	ro.Handle("HEAD", pattern, handler)
}

// NotFound sets a custom handler for 404 Not Found responses.
// This handler is called when no registered route matches the request.
//
// Purpose:
//   - Customize 404 error response (custom HTML, JSON, logging, etc.)
//   - Implement branded error pages
//   - Add telemetry/metrics for 404s
//   - Provide helpful error messages or suggestions
//
// Default Behavior:
//   - If NotFound is not called, uses http.NotFound() from standard library
//   - Default writes "404 page not found\n" with text/plain content type
//
// Thread-Safety:
//   - Thread-safe with write lock (mu.Lock)
//   - Safe to call while serving requests
//
// Example - Custom 404 Handler:
//
//	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
//	    w.WriteHeader(http.StatusNotFound)
//	    w.Write([]byte("Custom 404 - Page not found"))
//	})
//
// Example - JSON 404 Response:
//
//	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
//	    w.Header().Set("Content-Type", "application/json")
//	    w.WriteHeader(http.StatusNotFound)
//	    json.NewEncoder(w).Encode(map[string]string{
//	        "error": "not_found",
//	        "message": "The requested resource was not found",
//	    })
//	})
//
// NotFound는 404 Not Found 응답에 대한 커스텀 핸들러를 설정합니다.
// 이 핸들러는 등록된 라우트가 요청과 일치하지 않을 때 호출됩니다.
//
// 목적:
//   - 404 오류 응답 커스터마이즈(커스텀 HTML, JSON, 로깅 등)
//   - 브랜드화된 오류 페이지 구현
//   - 404에 대한 원격 측정/메트릭 추가
//   - 유용한 오류 메시지 또는 제안 제공
//
// 기본 동작:
//   - NotFound가 호출되지 않으면 표준 라이브러리의 http.NotFound() 사용
//   - 기본적으로 text/plain 콘텐츠 타입으로 "404 page not found\n" 작성
//
// 스레드 안전성:
//   - 쓰기 잠금(mu.Lock)으로 스레드 안전
//   - 요청을 서빙하는 동안 호출해도 안전
func (ro *Router) NotFound(handler http.HandlerFunc) {
	ro.mu.Lock()
	defer ro.mu.Unlock()
	ro.notFoundHandler = handler
}

// ServeHTTP implements the http.Handler interface for Router.
// This is the request processing entry point that matches requests to routes and executes handlers.
//
// ServeHTTP is automatically called by Go's http.Server for each incoming request.
// It performs route matching, parameter extraction, context creation, and handler execution.
//
// Purpose:
//   - Implements http.Handler interface (Router can be used as http.Handler)
//   - Matches incoming requests to registered routes
//   - Extracts URL parameters from matched routes
//   - Creates and populates Context with request data and parameters
//   - Invokes matched route handler or 404 handler
//
// Process Flow:
//  1. Acquire read lock (ro.mu.RLock) for thread-safe route lookup
//  2. Get list of routes registered for request's HTTP method
//  3. Iterate through routes in registration order
//  4. For each route, attempt to match request path:
//     - If match succeeds: Extract parameters, create Context, call handler, return
//     - If match fails: Continue to next route
//  5. If no routes match: Call notFoundHandler (404)
//  6. Release read lock (ro.mu.RUnlock)
//
// Route Matching:
//   - Routes are checked in registration order (first match wins)
//   - Matching is case-sensitive
//   - Parameters and wildcards are extracted into params map
//   - Match function returns (params, true) on success or (nil, false) on failure
//
// Context Creation:
//   - NewContext(w, r) creates Context wrapper around request/response
//   - Parameters from route matching are stored via ctx.setParams(params)
//   - Context is stored in request's context.Context for handler access
//   - Handlers retrieve Context via GetContext(r) or similar function
//
// Thread-Safety:
//   - Uses read lock for concurrent request handling
//   - Multiple requests can be processed simultaneously
//   - Route lookup is thread-safe with RWMutex
//   - Route registration (Handle) briefly blocks ServeHTTP with write lock
//
// Performance:
//   - Route lookup: O(n) where n = number of routes for HTTP method
//   - Path matching: O(m) where m = number of segments in pattern
//   - Typical latency: < 10μs for small route sets
//   - Scales well up to hundreds of routes per method
//
// Parameters:
//   - w: http.ResponseWriter for writing response
//   - r: *http.Request containing request details
//
// Example - Router as http.Handler:
//
//	router := newRouter()
//	router.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
//	    w.Write([]byte("Hello, World!"))
//	})
//
//	http.ListenAndServe(":8080", router)  // Router implements http.Handler
//
// ServeHTTP는 Router에 대한 http.Handler 인터페이스를 구현합니다.
// 이것은 요청을 라우트와 매칭하고 핸들러를 실행하는 요청 처리 진입점입니다.
//
// ServeHTTP는 각 들어오는 요청에 대해 Go의 http.Server에 의해 자동으로 호출됩니다.
// 라우트 매칭, 매개변수 추출, 컨텍스트 생성 및 핸들러 실행을 수행합니다.
//
// 목적:
//   - http.Handler 인터페이스 구현(Router를 http.Handler로 사용 가능)
//   - 들어오는 요청을 등록된 라우트와 매칭
//   - 매칭된 라우트에서 URL 매개변수 추출
//   - 요청 데이터 및 매개변수로 Context 생성 및 채우기
//   - 매칭된 라우트 핸들러 또는 404 핸들러 호출
//
// 프로세스 흐름:
//  1. 스레드 안전 라우트 조회를 위해 읽기 잠금(ro.mu.RLock) 획득
//  2. 요청의 HTTP 메서드에 등록된 라우트 목록 가져오기
//  3. 등록 순서대로 라우트 반복
//  4. 각 라우트에 대해 요청 경로와 매칭 시도:
//     - 매칭 성공: 매개변수 추출, Context 생성, 핸들러 호출, 반환
//     - 매칭 실패: 다음 라우트로 계속
//  5. 라우트가 일치하지 않으면: notFoundHandler 호출(404)
//  6. 읽기 잠금(ro.mu.RUnlock) 해제
//
// 라우트 매칭:
//   - 라우트는 등록 순서대로 확인됨(첫 번째 일치가 승리)
//   - 매칭은 대소문자 구분
//   - 매개변수 및 와일드카드는 params 맵으로 추출됨
//   - 매칭 함수는 성공 시 (params, true) 또는 실패 시 (nil, false) 반환
//
// 스레드 안전성:
//   - 동시 요청 처리를 위해 읽기 잠금 사용
//   - 여러 요청을 동시에 처리할 수 있음
//   - RWMutex로 라우트 조회가 스레드 안전
//   - 라우트 등록(Handle)은 쓰기 잠금으로 ServeHTTP를 잠깐 차단
//
// 성능:
//   - 라우트 조회: O(n), n = HTTP 메서드의 라우트 수
//   - 경로 매칭: O(m), m = 패턴의 세그먼트 수
//   - 일반적인 지연 시간: 작은 라우트 세트의 경우 < 10μs
//   - 메서드당 수백 개의 라우트까지 잘 확장됨
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

// contextWithValue stores the websvrutil Context in the request's context.Context.
// This is an internal helper function used by ServeHTTP to make Context accessible to handlers.
//
// Purpose:
//   - Embeds websvrutil.Context into Go's standard context.Context
//   - Enables handlers to retrieve Context via request context
//   - Preserves existing context chain (doesn't replace, wraps)
//
// Key:
//   - Uses contextKeyParams constant as storage key
//   - Key is package-private to prevent external access/conflicts
//
// Thread-Safety:
//   - context.WithValue creates new immutable context
//   - Safe for concurrent use
//   - Original context remains unchanged
//
// Usage:
//
//	Internal only - called by ServeHTTP:
//	  ctx := NewContext(w, r)
//	  ctx.setParams(params)
//	  r = r.WithContext(contextWithValue(r.Context(), ctx))
//
// contextWithValue는 요청의 context.Context에 websvrutil Context를 저장합니다.
// 이것은 ServeHTTP가 Context를 핸들러에서 액세스 가능하게 만드는 데 사용하는 내부 헬퍼 함수입니다.
//
// 목적:
//   - Go의 표준 context.Context에 websvrutil.Context 포함
//   - 핸들러가 요청 컨텍스트를 통해 Context를 검색할 수 있게 함
//   - 기존 컨텍스트 체인 보존(교체하지 않고 래핑)
func contextWithValue(ctx context.Context, c *Context) context.Context {
	return context.WithValue(ctx, contextKeyParams, c)
}

// match checks if the route matches the given path and extracts parameters.
// match는 라우트가 주어진 경로와 일치하는지 확인하고 매개변수를 추출합니다.
//
// Matching algorithm
// 매칭 알고리즘:
//  1. Parse incoming path into segments
//  2. Check segment count compatibility:
//     - Without wildcard: must match exactly
//     - With wildcard: request must have at least (pattern segments - 1)
//  3. Iterate through pattern segments:
//     - Wildcard (*): captures all remaining path segments
//     - Parameter (:name): captures current segment into parameter map
//     - Literal: must match exactly (case-sensitive)
//  4. Return parameters map and match status
//
// Return values
// 반환 값:
//   - params: Map of parameter names to values (e.g., {"id": "123"})
//   - matched: true if route matches, false otherwise
//
// Examples
// 예제:
//
//	Pattern: "/users/:id", Path: "/users/123"
//	-> params: {"id": "123"}, matched: true
//
//	Pattern: "/files/*", Path: "/files/docs/report.pdf"
//	-> params: {"*": "docs/report.pdf"}, matched: true
//
//	Pattern: "/users/:id", Path: "/posts/123"
//	-> params: nil, matched: false
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
// - Static segments: "/users/profile" - Exact match required
// - 정적 세그먼트: "/users/profile" - 정확한 일치 필요
// - Parameters: "/users/:id" - Captures value into "id" parameter
// - 매개변수: "/users/:id" - "id" 매개변수로 값 캡처
// - Wildcards: "/files/*" - Matches all remaining path segments
// - 와일드카드: "/files/*" - 나머지 모든 경로 세그먼트 일치
//
// Implementation details
// 구현 세부사항:
//  1. Trim leading/trailing slashes from pattern
//  2. Split pattern by "/" separator
//  3. Identify segment types:
//     - "*" -> wildcard segment (isWildcard=true)
//     - ":name" -> parameter segment (isParam=true, value="name")
//     - "literal" -> static segment (value="literal")
//  4. Pre-allocate segments slice with capacity for performance
//
// Time complexity: O(n) where n = number of path segments
// 시간 복잡도: O(n), n = 경로 세그먼트 수
//
// Examples
// 예제:
//
//	parsePattern("/users") -> [segment{value: "users"}]
//	parsePattern("/users/:id") -> [segment{value: "users"}, segment{value: "id", isParam: true}]
//	parsePattern("/files/*") -> [segment{value: "files"}, segment{isWildcard: true}]
func parsePattern(pattern string) []segment {
	// Remove leading and trailing slashes
	// 앞뒤 슬래시 제거
	pattern = strings.Trim(pattern, "/")

	if pattern == "" {
		return []segment{}
	}

	// Split by
	//
	// 로 분할
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
			// Remove the : prefix
			// : 접두사 제거
			seg.value = part[1:]
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
//  1. Trim leading and trailing slashes
//  2. Handle empty path (returns empty slice)
//  3. Split by "/" separator
//  4. Filter out empty segments (from consecutive slashes)
//  5. Return cleaned segment list
//
// Examples
// 예제:
//
//	parsePath("/users/123") -> ["users", "123"]
//	parsePath("/users/123/") -> ["users", "123"] (trailing slash removed)
//	parsePath("/api//v1/users") -> ["api", "v1", "users"] (double slash cleaned)
//	parsePath("/") -> [] (root path)
//	parsePath("") -> [] (empty path)
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
	//
	// 로 분할
	parts := strings.Split(path, "/")
	segments := make([]string, 0, len(parts))

	for _, part := range parts {
		if part != "" {
			segments = append(segments, part)
		}
	}

	return segments
}
