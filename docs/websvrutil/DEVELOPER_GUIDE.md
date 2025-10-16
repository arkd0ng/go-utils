# Websvrutil Package - Developer Guide / 개발자 가이드

**Package**: `github.com/arkd0ng/go-utils/websvrutil`
**Version**: v1.11.023
**Last Updated**: 2025-10-16

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#architecture-overview--아키텍처-개요)
2. [Package Structure / 패키지 구조](#package-structure--패키지-구조)
3. [Core Components / 핵심 컴포넌트](#core-components--핵심-컴포넌트)
4. [Internal Implementation / 내부 구현](#internal-implementation--내부-구현)
5. [Design Patterns / 디자인 패턴](#design-patterns--디자인-패턴)
6. [Adding New Features / 새 기능 추가](#adding-new-features--새-기능-추가)
7. [Testing Guide / 테스트 가이드](#testing-guide--테스트-가이드)
8. [Performance / 성능](#performance--성능)
9. [Contributing Guidelines / 기여 가이드라인](#contributing-guidelines--기여-가이드라인)
10. [Code Style / 코드 스타일](#code-style--코드-스타일)

---

## Architecture Overview / 아키텍처 개요

### Design Principles / 설계 원칙

The websvrutil package is built on several core design principles:

websvrutil 패키지는 여러 핵심 설계 원칙에 기반합니다:

1. **Developer Convenience First** / **개발자 편의 우선**
   - Intuitive API that reduces boilerplate code
   - 보일러플레이트 코드를 줄이는 직관적인 API
   - Method names that are self-documenting
   - 자체 문서화되는 메서드 이름

2. **Built on Standard Library** / **표준 라이브러리 기반**
   - Wraps `net/http` without hiding it
   - `net/http`를 숨기지 않고 래핑
   - Full compatibility with existing Go ecosystem
   - 기존 Go 생태계와 완전한 호환성

3. **Type Safety** / **타입 안전성**
   - Leverage Go's type system
   - Go의 타입 시스템 활용
   - Compile-time error detection
   - 컴파일 타임 에러 감지

4. **Performance** / **성능**
   - Zero-allocation context retrieval
   - 제로 할당 컨텍스트 검색
   - Efficient routing with trie-based router
   - 트라이 기반 라우터로 효율적인 라우팅
   - Minimal overhead over raw `net/http`
   - 원시 `net/http`에 비해 최소한의 오버헤드

5. **Thread Safety** / **스레드 안전성**
   - All components safe for concurrent use
   - 모든 컴포넌트가 동시 사용에 안전함
   - sync.RWMutex for shared state
   - 공유 상태를 위한 sync.RWMutex

### High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────┐
│                    HTTP Request                      │
└───────────────────┬─────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────────────────┐
│                  App (ServeHTTP)                     │
│  - Receives all HTTP requests                       │
│  - Routes to appropriate handler                    │
└───────────────────┬─────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────────────────┐
│               Middleware Chain                       │
│  - Global middleware (logging, recovery, etc.)      │
│  - Route-specific middleware (auth, etc.)           │
└───────────────────┬─────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────────────────┐
│                  Router (Match)                      │
│  - Trie-based path matching                         │
│  - HTTP method checking                             │
│  - Path parameter extraction                        │
└───────────────────┬─────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────────────────┐
│              Context (NewContext)                    │
│  - Created for each request                         │
│  - Stored in request context                        │
│  - Provides helper methods                          │
└───────────────────┬─────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────────────────┐
│                Handler Function                      │
│  - Business logic                                   │
│  - Uses Context helpers                             │
│  - Sends response                                   │
└─────────────────────────────────────────────────────┘
```

### Request Flow / 요청 흐름

```
1. HTTP Request arrives
   HTTP 요청 도착

2. App.ServeHTTP receives request
   App.ServeHTTP가 요청 수신

3. Context created and stored
   Context 생성 및 저장

4. Global middleware executes
   전역 미들웨어 실행

5. Router matches path and method
   Router가 경로 및 메서드 매칭

6. Path parameters extracted
   경로 매개변수 추출

7. Route-specific middleware executes
   라우트별 미들웨어 실행

8. Handler function executes
   핸들러 함수 실행

9. Response sent
   응답 전송

10. Cleanup (deferred functions)
    정리 (지연된 함수)
```

---

## Package Structure / 패키지 구조

### File Organization / 파일 조직

```
websvrutil/
├── websvrutil.go          # Package documentation and version / 패키지 문서 및 버전
├── app.go                 # App struct and lifecycle methods / App 구조체 및 생명주기 메서드
├── context.go             # Context struct and helper methods / Context 구조체 및 헬퍼 메서드
├── router.go              # Router struct and path matching / Router 구조체 및 경로 매칭
├── middleware.go          # Middleware chain management / 미들웨어 체인 관리
├── template.go            # Template rendering / 템플릿 렌더링
├── binding.go             # Request binding (JSON, XML) / 요청 바인딩 (JSON, XML)
├── session.go             # Session management / 세션 관리
├── static.go              # Static file serving / 정적 파일 서빙
├── storage.go             # Context storage (key-value) / 컨텍스트 저장소 (키-값)
├── shutdown.go            # Graceful shutdown / 우아한 종료
│
├── app_test.go            # App tests / App 테스트
├── context_test.go        # Context tests / Context 테스트
├── router_test.go         # Router tests / Router 테스트
├── middleware_test.go     # Middleware tests / 미들웨어 테스트
├── template_test.go       # Template tests / 템플릿 테스트
├── binding_test.go        # Binding tests / 바인딩 테스트
├── session_test.go        # Session tests / 세션 테스트
├── static_test.go         # Static file tests / 정적 파일 테스트
├── storage_test.go        # Storage tests / 저장소 테스트
├── shutdown_test.go       # Shutdown tests / 종료 테스트
├── cookie_test.go         # Cookie tests / 쿠키 테스트
├── method_test.go         # HTTP method tests / HTTP 메서드 테스트
└── error_test.go          # Error response tests / 에러 응답 테스트
```

### Responsibilities by File / 파일별 책임

| File / 파일 | Responsibility / 책임 |
|------------|---------------------|
| `app.go` | Application lifecycle (Run, Shutdown), HTTP server setup / 애플리케이션 생명주기 (Run, Shutdown), HTTP 서버 설정 |
| `context.go` | Request context helpers (Query, Param, JSON, etc.) / 요청 컨텍스트 헬퍼 (Query, Param, JSON 등) |
| `router.go` | Route registration and matching / 라우트 등록 및 매칭 |
| `middleware.go` | Middleware chain construction / 미들웨어 체인 구성 |
| `template.go` | HTML template loading and rendering / HTML 템플릿 로드 및 렌더링 |
| `binding.go` | Request body binding (JSON, XML) / 요청 본문 바인딩 (JSON, XML) |
| `session.go` | Cookie-based session management / 쿠키 기반 세션 관리 |
| `static.go` | Static file and directory serving / 정적 파일 및 디렉토리 서빙 |
| `storage.go` | Context key-value storage / 컨텍스트 키-값 저장소 |
| `shutdown.go` | Graceful shutdown with signal handling / 신호 처리를 포함한 우아한 종료 |

---

## Core Components / 핵심 컴포넌트

### 1. App / 앱

The `App` is the main application instance.

`App`은 메인 애플리케이션 인스턴스입니다.

**Structure:**

```go
type App struct {
    router           *Router
    middleware       []func(http.Handler) http.Handler
    server           *http.Server
    htmlTemplate     *template.Template
    funcMap          template.FuncMap
    mu               sync.RWMutex
    staticDirs       map[string]string // prefix -> directory
}
```

**Key Methods:**

- `New()` - Creates new App instance / 새 App 인스턴스 생성
- `Run(addr string)` - Starts HTTP server / HTTP 서버 시작
- `RunTLS(addr, certFile, keyFile string)` - Starts HTTPS server / HTTPS 서버 시작
- `RunWithGracefulShutdown(addr, timeout)` - Runs with signal handling / 신호 처리와 함께 실행
- `Shutdown(ctx)` - Gracefully shuts down server / 서버를 우아하게 종료
- `ServeHTTP(w, r)` - Implements http.Handler interface / http.Handler 인터페이스 구현

**Lifecycle:**

```go
app := New()              // 1. Create / 생성
app.Use(middleware)       // 2. Configure middleware / 미들웨어 설정
app.GET("/", handler)     // 3. Register routes / 라우트 등록
app.Run(":8080")          // 4. Start server / 서버 시작
// ... server running ... // ... 서버 실행 중 ...
app.Shutdown(ctx)         // 5. Graceful shutdown / 우아한 종료
```

### 2. Context / 컨텍스트

The `Context` provides request/response helpers.

`Context`는 요청/응답 헬퍼를 제공합니다.

**Structure:**

```go
type Context struct {
    Request        *http.Request
    ResponseWriter http.ResponseWriter
    Params         map[string]string
    statusCode     int
    storage        map[string]interface{}
    mu             sync.RWMutex
}
```

**Key Methods:**

**Request Methods / 요청 메서드:**
- `Param(key)` - Get path parameter / 경로 매개변수 가져오기
- `Query(key)` - Get query parameter / 쿼리 매개변수 가져오기
- `DefaultQuery(key, defaultValue)` - Get query with default / 기본값과 함께 쿼리 가져오기
- `BindJSON(v)` - Parse JSON body / JSON 본문 파싱
- `BindXML(v)` - Parse XML body / XML 본문 파싱

**Response Methods / 응답 메서드:**
- `String(code, format, ...values)` - Send string response / 문자열 응답 전송
- `JSON(code, obj)` - Send JSON response / JSON 응답 전송
- `XML(code, obj)` - Send XML response / XML 응답 전송
- `HTML(code, name, data)` - Render template / 템플릿 렌더링
- `Data(code, contentType, data)` - Send binary data / 바이너리 데이터 전송

**Header Methods / 헤더 메서드:**
- `GetHeader(key)` - Get request header / 요청 헤더 가져오기
- `SetHeader(key, value)` - Set response header / 응답 헤더 설정
- `SetHeaders(headers)` - Set multiple headers / 여러 헤더 설정

**Cookie Methods / 쿠키 메서드:**
- `Cookie(name)` - Get cookie / 쿠키 가져오기
- `CookieValue(name)` - Get cookie value (convenience) / 쿠키 값 가져오기 (편의)
- `SetCookie(cookie)` - Set cookie / 쿠키 설정
- `SetCookieAdvanced(opts)` - Set cookie with options / 옵션과 함께 쿠키 설정
- `DeleteCookie(name, path)` - Delete cookie / 쿠키 삭제

**Storage Methods / 저장소 메서드:**
- `Set(key, value)` - Store value in context / 컨텍스트에 값 저장
- `Get(key)` - Retrieve value / 값 검색
- `MustGet(key)` - Get or panic / 가져오거나 패닉
- `GetString(key)`, `GetInt(key)`, `GetBool(key)` - Type-safe getters / 타입 안전 getter

**Error Methods / 에러 메서드:**
- `ErrorJSON(code, message)` - Send error JSON / 에러 JSON 전송
- `SuccessJSON(code, message, data)` - Send success JSON / 성공 JSON 전송
- `AbortWithStatus(code)` - Abort with status code / 상태 코드로 중단
- `NotFound()`, `Unauthorized()`, etc. - HTTP error shortcuts / HTTP 에러 단축

**Context Creation / 컨텍스트 생성:**

```go
// Internal: Created by App for each request / 내부: App이 각 요청에 대해 생성
func NewContext(w http.ResponseWriter, r *http.Request) *Context

// Public: Retrieve context from request / 공개: 요청에서 컨텍스트 검색
func GetContext(r *http.Request) *Context
```

### 3. Router / 라우터

The `Router` handles path matching and route registration.

`Router`는 경로 매칭과 라우트 등록을 처리합니다.

**Structure:**

```go
type Router struct {
    routes map[string]*routeNode // method -> root node
    mu     sync.RWMutex
}

type routeNode struct {
    path       string
    isParam    bool
    paramName  string
    handler    http.Handler
    middleware []func(http.Handler) http.Handler
    children   map[string]*routeNode
}
```

**Path Matching Algorithm:**

The router uses a trie (prefix tree) data structure for efficient path matching.

라우터는 효율적인 경로 매칭을 위해 트라이(접두사 트리) 자료 구조를 사용합니다.

```
Example routes:
/users
/users/:id
/users/:id/posts
/posts

Trie structure:
      root
       |
     users
      / \
    :id posts
     |
   posts
```

**Matching Process:**

1. Split path into segments / 경로를 세그먼트로 분할
2. Traverse trie from root / 루트에서 트라이 순회
3. Match static segments first / 정적 세그먼트 먼저 매칭
4. Fall back to parameter segments / 매개변수 세그먼트로 대체
5. Extract parameter values / 매개변수 값 추출

**Time Complexity:**
- Registration: O(m) where m = number of path segments / 등록: O(m), m = 경로 세그먼트 수
- Matching: O(m) where m = number of path segments / 매칭: O(m), m = 경로 세그먼트 수

### 4. Middleware / 미들웨어

Middleware is a function that wraps an `http.Handler`.

미들웨어는 `http.Handler`를 래핑하는 함수입니다.

**Type:**

```go
type Middleware func(http.Handler) http.Handler
```

**Execution Order:**

```
Request → Middleware1 → Middleware2 → Middleware3 → Handler
          ↓                ↓               ↓
       (before)        (before)        (before)
          ↓                ↓               ↓
       Handler execution / 핸들러 실행
          ↓                ↓               ↓
       (after)          (after)         (after)
          ↓                ↓               ↓
Response ← Middleware1 ← Middleware2 ← Middleware3 ← Handler
```

**Example Implementation:**

```go
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Before handler / 핸들러 전
        start := time.Now()

        // Call next handler / 다음 핸들러 호출
        next.ServeHTTP(w, r)

        // After handler / 핸들러 후
        duration := time.Since(start)
        log.Printf("%s %s - %v", r.Method, r.URL.Path, duration)
    })
}
```

### 5. Session Management / 세션 관리

**Components:**

```go
type SessionStore struct {
    mu       sync.RWMutex
    sessions map[string]*Session
    options  SessionOptions
}

type Session struct {
    ID        string
    Data      map[string]interface{}
    CreatedAt time.Time
    ExpiresAt time.Time
    mu        sync.RWMutex
}

type SessionOptions struct {
    CookieName  string
    MaxAge      time.Duration
    Secure      bool
    HttpOnly    bool
    SameSite    http.SameSite
    CleanupTime time.Duration
    Path        string
    Domain      string
}
```

**Session ID Generation:**

Uses cryptographically secure random bytes:

암호학적으로 안전한 랜덤 바이트 사용:

```go
func (s *SessionStore) generateSessionID() string {
    b := make([]byte, 32) // 256 bits
    if _, err := rand.Read(b); err != nil {
        // Fallback to timestamp-based ID
        return base64.URLEncoding.EncodeToString([]byte(time.Now().String()))
    }
    return base64.URLEncoding.EncodeToString(b)
}
```

**Automatic Cleanup:**

Background goroutine removes expired sessions:

백그라운드 고루틴이 만료된 세션 제거:

```go
func (s *SessionStore) cleanupExpiredSessions() {
    ticker := time.NewTicker(s.options.CleanupTime)
    defer ticker.Stop()

    for range ticker.C {
        now := time.Now()
        s.mu.Lock()
        for id, session := range s.sessions {
            if session.ExpiresAt.Before(now) {
                delete(s.sessions, id)
            }
        }
        s.mu.Unlock()
    }
}
```

---

## Internal Implementation / 내부 구현

### Context Storage / 컨텍스트 저장

The Context is stored in the request's context.Context using a unique key:

Context는 고유 키를 사용하여 요청의 context.Context에 저장됩니다:

```go
type contextKey string

const contextKeyContext contextKey = "websvrutil_context"

// Store context in request / 요청에 컨텍스트 저장
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ctx := NewContext(w, r)
    r = r.WithContext(context.WithValue(r.Context(), contextKeyContext, ctx))

    // ... handle request ... / ... 요청 처리 ...
}

// Retrieve context from request / 요청에서 컨텍스트 검색
func GetContext(r *http.Request) *Context {
    ctx, ok := r.Context().Value(contextKeyContext).(*Context)
    if !ok {
        return NewContext(nil, r)
    }
    return ctx
}
```

**Benefits:**
- Zero allocation after initial creation / 초기 생성 후 제로 할당
- Type-safe retrieval / 타입 안전 검색
- Automatic cleanup / 자동 정리

### Router Implementation / 라우터 구현

**Registration:**

```go
func (r *Router) addRoute(method, path string, handler http.Handler, middleware ...func(http.Handler) http.Handler) {
    r.mu.Lock()
    defer r.mu.Unlock()

    // Get or create root node for method / 메서드에 대한 루트 노드 가져오기 또는 생성
    root, exists := r.routes[method]
    if !exists {
        root = &routeNode{path: "/", children: make(map[string]*routeNode)}
        r.routes[method] = root
    }

    // Split path into segments / 경로를 세그먼트로 분할
    segments := strings.Split(strings.Trim(path, "/"), "/")

    // Traverse/create trie nodes / 트라이 노드 순회/생성
    current := root
    for _, segment := range segments {
        if segment == "" {
            continue
        }

        // Check if parameter / 매개변수인지 확인
        isParam := strings.HasPrefix(segment, ":")
        key := segment
        if isParam {
            key = ":" // All params use same key / 모든 매개변수는 동일한 키 사용
        }

        // Get or create child node / 자식 노드 가져오기 또는 생성
        child, exists := current.children[key]
        if !exists {
            child = &routeNode{
                path:     segment,
                isParam:  isParam,
                children: make(map[string]*routeNode),
            }
            if isParam {
                child.paramName = strings.TrimPrefix(segment, ":")
            }
            current.children[key] = child
        }

        current = child
    }

    // Set handler and middleware at final node / 최종 노드에 핸들러 및 미들웨어 설정
    current.handler = handler
    current.middleware = middleware
}
```

**Matching:**

```go
func (r *Router) match(method, path string) (*routeNode, map[string]string) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    root, exists := r.routes[method]
    if !exists {
        return nil, nil
    }

    segments := strings.Split(strings.Trim(path, "/"), "/")
    params := make(map[string]string)

    current := root
    for _, segment := range segments {
        if segment == "" {
            continue
        }

        // Try static match first / 먼저 정적 매칭 시도
        if child, exists := current.children[segment]; exists {
            current = child
            continue
        }

        // Try parameter match / 매개변수 매칭 시도
        if child, exists := current.children[":"]; exists {
            params[child.paramName] = segment
            current = child
            continue
        }

        // No match / 매칭 없음
        return nil, nil
    }

    if current.handler == nil {
        return nil, nil
    }

    return current, params
}
```

### Middleware Chain / 미들웨어 체인

**Construction:**

```go
func buildMiddlewareChain(handler http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
    // Apply middleware in reverse order / 역순으로 미들웨어 적용
    for i := len(middleware) - 1; i >= 0; i-- {
        handler = middleware[i](handler)
    }
    return handler
}
```

**Example:**

```go
middleware := []func(http.Handler) http.Handler{
    LoggerMiddleware,
    RecoveryMiddleware,
    AuthMiddleware,
}

handler := buildMiddlewareChain(finalHandler, middleware...)

// Results in: / 결과:
// LoggerMiddleware(RecoveryMiddleware(AuthMiddleware(finalHandler)))
```

### Template Rendering / 템플릿 렌더링

**Template Loading:**

```go
func (a *App) LoadHTMLGlob(pattern string) {
    a.mu.Lock()
    defer a.mu.Unlock()

    tmpl := template.New("")

    // Apply custom functions / 커스텀 함수 적용
    if a.funcMap != nil {
        tmpl = tmpl.Funcs(a.funcMap)
    }

    // Parse templates / 템플릿 파싱
    var err error
    a.htmlTemplate, err = tmpl.ParseGlob(pattern)
    if err != nil {
        panic(err)
    }
}
```

**Template Execution:**

```go
func (c *Context) HTML(code int, name string, data interface{}) {
    c.SetHeader("Content-Type", "text/html; charset=utf-8")
    c.Status(code)

    // Get template / 템플릿 가져오기
    tmpl := getHTMLTemplate() // From App

    // Execute template / 템플릿 실행
    if err := tmpl.ExecuteTemplate(c.ResponseWriter, name, data); err != nil {
        http.Error(c.ResponseWriter, err.Error(), http.StatusInternalServerError)
    }
}
```

### Request Binding / 요청 바인딩

**JSON Binding:**

```go
func (c *Context) BindJSON(v interface{}) error {
    if c.Request.Body == nil {
        return errors.New("request body is empty")
    }

    // Decode JSON / JSON 디코드
    decoder := json.NewDecoder(c.Request.Body)
    if err := decoder.Decode(v); err != nil {
        return err
    }

    // Validate (if validator is available) / 검증 (검증자가 있는 경우)
    if err := validate(v); err != nil {
        return err
    }

    return nil
}
```

**Validation:**

Uses struct tags for validation:

검증을 위해 구조체 태그 사용:

```go
type User struct {
    Name  string `json:"name" binding:"required,min=3,max=50"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=130"`
}
```

---

## Design Patterns / 디자인 패턴

### 1. Context Pattern / 컨텍스트 패턴

**Problem:** Need to pass request-scoped data through handlers without function parameters.

**문제:** 함수 매개변수 없이 핸들러를 통해 요청 범위 데이터를 전달해야 함.

**Solution:** Store context in `context.Context` with unique key.

**해결:** 고유 키로 `context.Context`에 컨텍스트 저장.

```go
// Store / 저장
r = r.WithContext(context.WithValue(r.Context(), contextKeyContext, ctx))

// Retrieve / 검색
ctx := GetContext(r)
```

### 2. Middleware Chain Pattern / 미들웨어 체인 패턴

**Problem:** Need to wrap handlers with cross-cutting concerns (logging, auth, etc.).

**문제:** 횡단 관심사(로깅, 인증 등)로 핸들러를 래핑해야 함.

**Solution:** Function composition with `http.Handler` interface.

**해결:** `http.Handler` 인터페이스를 사용한 함수 합성.

```go
type Middleware func(http.Handler) http.Handler

func buildChain(handler http.Handler, mw ...Middleware) http.Handler {
    for i := len(mw) - 1; i >= 0; i-- {
        handler = mw[i](handler)
    }
    return handler
}
```

### 3. Builder Pattern / 빌더 패턴

**Problem:** Many optional configuration parameters.

**문제:** 많은 선택적 설정 매개변수.

**Solution:** Functional options pattern.

**해결:** 함수형 옵션 패턴.

```go
type SessionOptions struct {
    CookieName  string
    MaxAge      time.Duration
    Secure      bool
    // ... more options ... / ... 더 많은 옵션 ...
}

func DefaultSessionOptions() SessionOptions {
    return SessionOptions{
        CookieName: "sessionid",
        MaxAge:     24 * time.Hour,
        // ... defaults ... / ... 기본값 ...
    }
}

// Usage / 사용
store := NewSessionStore(SessionOptions{
    CookieName: "my_session",
    MaxAge:     12 * time.Hour,
})
```

### 4. Singleton Pattern / 싱글톤 패턴

**Problem:** Need single shared instance of session store or other resources.

**문제:** 세션 저장소 또는 기타 리소스의 단일 공유 인스턴스 필요.

**Solution:** Global variable with initialization function.

**해결:** 초기화 함수가 있는 전역 변수.

```go
var sessionStore *SessionStore

func init() {
    sessionStore = NewSessionStore(DefaultSessionOptions())
}
```

### 5. Template Method Pattern / 템플릿 메서드 패턴

**Problem:** Common request handling logic with customizable parts.

**문제:** 사용자 정의 가능한 부분이 있는 공통 요청 처리 로직.

**Solution:** Base handler with extension points.

**해결:** 확장 지점이 있는 기본 핸들러.

```go
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // 1. Create context / 컨텍스트 생성
    ctx := NewContext(w, r)
    r = r.WithContext(context.WithValue(r.Context(), contextKeyContext, ctx))

    // 2. Apply global middleware / 전역 미들웨어 적용
    handler := buildMiddlewareChain(a.router, a.middleware...)

    // 3. Execute handler / 핸들러 실행
    handler.ServeHTTP(w, r)

    // 4. Cleanup (deferred) / 정리 (지연)
}
```

### 6. Observer Pattern (Sessions) / 옵저버 패턴 (세션)

**Problem:** Need to clean up expired sessions automatically.

**문제:** 만료된 세션을 자동으로 정리해야 함.

**Solution:** Background goroutine with ticker.

**해결:** 티커가 있는 백그라운드 고루틴.

```go
func (s *SessionStore) cleanupExpiredSessions() {
    ticker := time.NewTicker(s.options.CleanupTime)
    defer ticker.Stop()

    for range ticker.C {
        // Clean up expired sessions / 만료된 세션 정리
        now := time.Now()
        s.mu.Lock()
        for id, session := range s.sessions {
            if session.ExpiresAt.Before(now) {
                delete(s.sessions, id)
            }
        }
        s.mu.Unlock()
    }
}
```

---

## Adding New Features / 새 기능 추가

### Adding a New Context Method / 새로운 Context 메서드 추가

**Example: Add `GetClientIP()` method**

**예제: `GetClientIP()` 메서드 추가**

1. **Add method to `context.go`:**

```go
// GetClientIP returns the client's IP address.
// GetClientIP는 클라이언트의 IP 주소를 반환합니다.
func (c *Context) GetClientIP() string {
    // Try X-Forwarded-For header / X-Forwarded-For 헤더 시도
    if ip := c.GetHeader("X-Forwarded-For"); ip != "" {
        // Get first IP in list / 목록에서 첫 번째 IP 가져오기
        if idx := strings.Index(ip, ","); idx != -1 {
            return strings.TrimSpace(ip[:idx])
        }
        return ip
    }

    // Try X-Real-IP header / X-Real-IP 헤더 시도
    if ip := c.GetHeader("X-Real-IP"); ip != "" {
        return ip
    }

    // Fall back to RemoteAddr / RemoteAddr로 대체
    ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
    return ip
}
```

2. **Add tests to `context_test.go`:**

```go
func TestGetClientIP(t *testing.T) {
    tests := []struct {
        name         string
        headers      map[string]string
        remoteAddr   string
        expectedIP   string
    }{
        {
            name: "X-Forwarded-For single IP",
            headers: map[string]string{
                "X-Forwarded-For": "203.0.113.1",
            },
            expectedIP: "203.0.113.1",
        },
        {
            name: "X-Forwarded-For multiple IPs",
            headers: map[string]string{
                "X-Forwarded-For": "203.0.113.1, 203.0.113.2",
            },
            expectedIP: "203.0.113.1",
        },
        {
            name: "X-Real-IP",
            headers: map[string]string{
                "X-Real-IP": "203.0.113.3",
            },
            expectedIP: "203.0.113.3",
        },
        {
            name:       "RemoteAddr",
            remoteAddr: "203.0.113.4:12345",
            expectedIP: "203.0.113.4",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodGet, "/", nil)
            if tt.remoteAddr != "" {
                req.RemoteAddr = tt.remoteAddr
            }
            for k, v := range tt.headers {
                req.Header.Set(k, v)
            }

            w := httptest.NewRecorder()
            ctx := NewContext(w, req)

            ip := ctx.GetClientIP()
            if ip != tt.expectedIP {
                t.Errorf("Expected %s, got %s", tt.expectedIP, ip)
            }
        })
    }
}
```

3. **Update README.md** with new method documentation.

4. **Update CHANGELOG.md** with the addition.

### Adding a New Middleware / 새로운 미들웨어 추가

**Example: Add Rate Limiting Middleware**

**예제: 속도 제한 미들웨어 추가**

1. **Create `ratelimit.go`:**

```go
package websvrutil

import (
    "net/http"
    "sync"
    "time"
)

// RateLimiter implements token bucket algorithm.
// RateLimiter는 토큰 버킷 알고리즘을 구현합니다.
type RateLimiter struct {
    tokens    int
    maxTokens int
    rate      time.Duration
    mu        sync.Mutex
    lastRefill time.Time
}

// NewRateLimiter creates a new rate limiter.
// NewRateLimiter는 새로운 속도 제한기를 생성합니다.
func NewRateLimiter(maxTokens int, rate time.Duration) *RateLimiter {
    return &RateLimiter{
        tokens:    maxTokens,
        maxTokens: maxTokens,
        rate:      rate,
        lastRefill: time.Now(),
    }
}

// Allow checks if request is allowed.
// Allow는 요청이 허용되는지 확인합니다.
func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    // Refill tokens / 토큰 리필
    now := time.Now()
    elapsed := now.Sub(rl.lastRefill)
    tokensToAdd := int(elapsed / rl.rate)

    if tokensToAdd > 0 {
        rl.tokens = min(rl.tokens+tokensToAdd, rl.maxTokens)
        rl.lastRefill = now
    }

    // Check if tokens available / 토큰 사용 가능 여부 확인
    if rl.tokens > 0 {
        rl.tokens--
        return true
    }

    return false
}

// RateLimitMiddleware creates rate limiting middleware.
// RateLimitMiddleware는 속도 제한 미들웨어를 생성합니다.
func RateLimitMiddleware(limiter *RateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := GetContext(r)

            if !limiter.Allow() {
                ctx.ErrorJSON(429, "Rate limit exceeded")
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

2. **Add tests to `ratelimit_test.go`:**

```go
func TestRateLimiter(t *testing.T) {
    limiter := NewRateLimiter(2, 100*time.Millisecond)

    // First two requests should succeed / 처음 두 요청은 성공해야 함
    if !limiter.Allow() {
        t.Error("First request should be allowed")
    }
    if !limiter.Allow() {
        t.Error("Second request should be allowed")
    }

    // Third request should fail / 세 번째 요청은 실패해야 함
    if limiter.Allow() {
        t.Error("Third request should be denied")
    }

    // Wait for refill / 리필 대기
    time.Sleep(150 * time.Millisecond)

    // Should allow after refill / 리필 후 허용해야 함
    if !limiter.Allow() {
        t.Error("Request after refill should be allowed")
    }
}
```

3. **Usage example in documentation:**

```go
limiter := websvrutil.NewRateLimiter(10, time.Second) // 10 req/sec
app.Use(websvrutil.RateLimitMiddleware(limiter))
```

---

## Testing Guide / 테스트 가이드

### Test Structure / 테스트 구조

Each feature has its own test file:

각 기능은 자체 테스트 파일을 가집니다:

```
websvrutil/
├── context.go        → context_test.go
├── router.go         → router_test.go
├── middleware.go     → middleware_test.go
└── session.go        → session_test.go
```

### Running Tests / 테스트 실행

```bash
# Run all tests / 모든 테스트 실행
go test ./websvrutil -v

# Run specific test file / 특정 테스트 파일 실행
go test ./websvrutil -v -run TestContext

# Run with coverage / 커버리지와 함께 실행
go test ./websvrutil -cover

# Generate coverage report / 커버리지 보고서 생성
go test ./websvrutil -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Writing Tests / 테스트 작성

**Example test structure:**

```go
func TestFeatureName(t *testing.T) {
    // Setup / 설정
    app := New()
    req := httptest.NewRequest(http.MethodGet, "/test", nil)
    w := httptest.NewRecorder()

    // Execute / 실행
    app.ServeHTTP(w, req)

    // Assert / 단언
    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}
```

**Table-driven tests:**

```go
func TestQuery(t *testing.T) {
    tests := []struct {
        name     string
        url      string
        key      string
        expected string
    }{
        {"simple query", "/test?name=John", "name", "John"},
        {"missing query", "/test", "name", ""},
        {"empty value", "/test?name=", "name", ""},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodGet, tt.url, nil)
            w := httptest.NewRecorder()
            ctx := NewContext(w, req)

            result := ctx.Query(tt.key)
            if result != tt.expected {
                t.Errorf("Expected %q, got %q", tt.expected, result)
            }
        })
    }
}
```

### Benchmarks / 벤치마크

**Example benchmark:**

```go
func BenchmarkRouterMatch(b *testing.B) {
    router := NewRouter()
    router.addRoute("GET", "/users/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        router.match("GET", "/users/123")
    }
}
```

**Running benchmarks:**

```bash
# Run all benchmarks / 모든 벤치마크 실행
go test ./websvrutil -bench=.

# Run specific benchmark / 특정 벤치마크 실행
go test ./websvrutil -bench=BenchmarkRouterMatch

# With memory stats / 메모리 통계와 함께
go test ./websvrutil -bench=. -benchmem
```

### Test Coverage Goals / 테스트 커버리지 목표

- Overall coverage: > 75% / 전체 커버리지: > 75%
- Critical paths (routing, context): > 90% / 중요 경로 (라우팅, 컨텍스트): > 90%
- Error handling: 100% / 에러 처리: 100%

Current coverage: **79.4%** ✓

현재 커버리지: **79.4%** ✓

---

## Performance / 성능

### Benchmarks / 벤치마크

**Router matching:**

```
BenchmarkRouterMatch-8         5000000       250 ns/op       0 B/op       0 allocs/op
```

**Context creation:**

```
BenchmarkNewContext-8         10000000       150 ns/op      128 B/op       2 allocs/op
```

**JSON response:**

```
BenchmarkJSONResponse-8        1000000      1200 ns/op      512 B/op       5 allocs/op
```

### Optimization Tips / 최적화 팁

1. **Pre-compile regular expressions / 정규 표현식 사전 컴파일**
   ```go
   var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
   ```

2. **Reuse buffers / 버퍼 재사용**
   ```go
   var bufferPool = sync.Pool{
       New: func() interface{} {
           return new(bytes.Buffer)
       },
   }
   ```

3. **Minimize allocations / 할당 최소화**
   ```go
   // Bad: Creates new slice / 나쁨: 새 슬라이스 생성
   segments := strings.Split(path, "/")

   // Good: Preallocate / 좋음: 사전 할당
   segments := make([]string, 0, 10)
   ```

4. **Use sync.Pool for frequently created objects / 자주 생성되는 객체에 sync.Pool 사용**
   ```go
   var contextPool = sync.Pool{
       New: func() interface{} {
           return &Context{}
       },
   }
   ```

---

## Contributing Guidelines / 기여 가이드라인

### Before Contributing / 기여 전

1. Read this developer guide / 이 개발자 가이드 읽기
2. Check existing issues and PRs / 기존 이슈 및 PR 확인
3. Discuss major changes in an issue first / 주요 변경 사항은 먼저 이슈에서 논의

### Pull Request Process / 풀 리퀘스트 프로세스

1. **Fork and clone the repository / 저장소 포크 및 복제**
   ```bash
   git clone https://github.com/arkd0ng/go-utils.git
   cd go-utils/websvrutil
   ```

2. **Create a feature branch / 기능 브랜치 생성**
   ```bash
   git checkout -b feature/my-new-feature
   ```

3. **Make changes / 변경**
   - Write code / 코드 작성
   - Add tests / 테스트 추가
   - Update documentation / 문서 업데이트

4. **Run tests / 테스트 실행**
   ```bash
   go test ./websvrutil -v
   go test ./websvrutil -cover
   ```

5. **Commit changes / 변경 커밋**
   ```bash
   git add .
   git commit -m "Feat: Add new feature"
   ```

6. **Push and create PR / 푸시 및 PR 생성**
   ```bash
   git push origin feature/my-new-feature
   ```

### Commit Message Format / 커밋 메시지 형식

```
Type: Short description (max 72 chars)

Longer description if needed. Wrap at 72 characters. Explain what
and why vs. how.

Fixes #123
```

**Types / 타입:**
- `Feat`: New feature / 새 기능
- `Fix`: Bug fix / 버그 수정
- `Docs`: Documentation / 문서
- `Refactor`: Code refactoring / 코드 리팩토링
- `Test`: Tests / 테스트
- `Chore`: Build, config, etc. / 빌드, 설정 등

### Code Review Checklist / 코드 리뷰 체크리스트

- [ ] Code follows style guide / 코드가 스타일 가이드를 따름
- [ ] Tests added and passing / 테스트 추가 및 통과
- [ ] Documentation updated / 문서 업데이트
- [ ] No breaking changes (or documented) / 호환성 파괴 변경 없음 (또는 문서화됨)
- [ ] Performance impact considered / 성능 영향 고려
- [ ] Error handling complete / 에러 처리 완료
- [ ] Thread safety verified / 스레드 안전성 검증

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

**Variables / 변수:**
- Use camelCase / camelCase 사용
- Descriptive names / 설명적인 이름
- Avoid single-letter names (except loop counters) / 단일 문자 이름 피하기 (루프 카운터 제외)

```go
// Good / 좋음
var userCount int
var sessionStore *SessionStore

// Bad / 나쁨
var uc int
var ss *SessionStore
```

**Functions / 함수:**
- Use PascalCase for exported / 내보내기에는 PascalCase 사용
- Use camelCase for unexported / 내보내지 않기에는 camelCase 사용
- Start with verb / 동사로 시작

```go
// Good / 좋음
func GetContext(r *http.Request) *Context
func buildMiddlewareChain(handler http.Handler) http.Handler

// Bad / 나쁨
func context(r *http.Request) *Context
func chain(handler http.Handler) http.Handler
```

**Types / 타입:**
- Use PascalCase / PascalCase 사용
- Avoid stuttering / 말더듬 피하기

```go
// Good / 좋음
type Router struct { ... }
type SessionStore struct { ... }

// Bad / 나쁨
type WebRouter struct { ... }  // "websvrutil.WebRouter" stutters
type SessionStoreStore struct { ... }
```

### Comments / 주석

**Package comments / 패키지 주석:**
```go
// Package websvrutil provides a lightweight HTTP framework.
// Package websvrutil는 가벼운 HTTP 프레임워크를 제공합니다.
package websvrutil
```

**Function comments / 함수 주석:**
```go
// GetContext retrieves the Context from the request.
// GetContext는 요청에서 Context를 검색합니다.
func GetContext(r *http.Request) *Context
```

**Inline comments / 인라인 주석:**
```go
// Check if path parameter / 경로 매개변수인지 확인
isParam := strings.HasPrefix(segment, ":")
```

### Error Handling / 에러 처리

**Always check errors / 항상 에러 확인:**
```go
// Good / 좋음
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()

// Bad / 나쁨
file, _ := os.Open("file.txt")
defer file.Close()
```

**Wrap errors with context / 컨텍스트와 함께 에러 래핑:**
```go
if err := validate(user); err != nil {
    return fmt.Errorf("failed to validate user: %w", err)
}
```

### Formatting / 포맷팅

**Use gofmt / gofmt 사용:**
```bash
gofmt -w .
```

**Use goimports / goimports 사용:**
```bash
goimports -w .
```

### Best Practices / 모범 사례

1. **Keep functions small / 함수를 작게 유지**
   - Max 50 lines / 최대 50줄
   - Single responsibility / 단일 책임

2. **Avoid global state / 전역 상태 피하기**
   - Use dependency injection / 의존성 주입 사용
   - Pass dependencies explicitly / 명시적으로 의존성 전달

3. **Use interfaces / 인터페이스 사용**
   - Accept interfaces, return structs / 인터페이스 받기, 구조체 반환
   - Keep interfaces small / 인터페이스를 작게 유지

4. **Document exported items / 내보낸 항목 문서화**
   - All exported types, functions, constants / 모든 내보낸 타입, 함수, 상수
   - Bilingual comments (English / Korean) / 이중 언어 주석 (영문 / 한글)

5. **Test coverage / 테스트 커버리지**
   - > 75% overall / 전체 > 75%
   - All error paths / 모든 에러 경로
   - Edge cases / 엣지 케이스

---

## Conclusion / 결론

This developer guide provides a comprehensive overview of the websvrutil package's internal architecture, design patterns, and best practices. By following these guidelines, contributors can maintain consistency and quality across the codebase.

이 개발자 가이드는 websvrutil 패키지의 내부 아키텍처, 디자인 패턴 및 모범 사례에 대한 포괄적인 개요를 제공합니다. 이러한 가이드라인을 따름으로써 기여자는 코드베이스 전체에서 일관성과 품질을 유지할 수 있습니다.

For questions or discussions, please open an issue on GitHub.

질문이나 논의가 있으면 GitHub에서 이슈를 여세요.

---

**Version**: v1.11.023
**Last Updated**: 2025-10-16
**License**: MIT
