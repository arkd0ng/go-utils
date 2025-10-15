# websvrutil Design Plan / 설계 계획

**Version / 버전**: v1.11.001
**Created / 생성일**: 2025-10-16
**Package / 패키지**: `github.com/arkd0ng/go-utils/websvrutil`

---

## Table of Contents / 목차

- [Package Overview / 패키지 개요](#package-overview--패키지-개요)
- [Design Philosophy / 설계 철학](#design-philosophy--설계-철학)
- [Architecture / 아키텍처](#architecture--아키텍처)
- [Core Components / 핵심 컴포넌트](#core-components--핵심-컴포넌트)
- [API Design / API 설계](#api-design--api-설계)
- [Template System / 템플릿 시스템](#template-system--템플릿-시스템)
- [Error Handling / 에러 처리](#error-handling--에러-처리)
- [Performance Considerations / 성능 고려사항](#performance-considerations--성능-고려사항)

---

## Package Overview / 패키지 개요

### Purpose / 목적

Provide extreme simplicity for HTTP server development with focus on **developer convenience over raw performance**.

HTTP 서버 개발을 위한 극도의 간결함을 제공하며, **순수 성능보다 개발자 편의성**에 초점을 맞춥니다.

### Problem Statement / 문제 정의

**Before / 이전**:
- Setting up HTTP servers requires 50+ lines of boilerplate code / HTTP 서버 설정에 50줄 이상의 보일러플레이트 코드 필요
- Complex handler registration / 복잡한 핸들러 등록
- Tedious template management / 지루한 템플릿 관리
- Repetitive middleware setup / 반복적인 미들웨어 설정
- Manual error handling / 수동 에러 처리
- Verbose JSON/response writing / 장황한 JSON/응답 작성

**After / 이후**:
- 5-10 lines of code for complete server setup / 완전한 서버 설정을 5-10줄 코드로
- Intuitive handler registration / 직관적인 핸들러 등록
- Auto template discovery and hot reload / 자동 템플릿 발견 및 핫 리로드
- One-line middleware chaining / 한 줄 미들웨어 체이닝
- Automatic error handling / 자동 에러 처리
- Helper methods for JSON/HTML responses / JSON/HTML 응답을 위한 헬퍼 메서드

### Target Use Cases / 목표 사용 사례

1. **REST API Servers** / REST API 서버
   - Quick JSON API development / 빠른 JSON API 개발
   - CRUD operations / CRUD 작업

2. **Web Applications** / 웹 애플리케이션
   - Template-based rendering / 템플릿 기반 렌더링
   - Form handling / 폼 처리

3. **Microservices** / 마이크로서비스
   - Quick service endpoints / 빠른 서비스 엔드포인트
   - Health checks / 헬스 체크

4. **Admin Dashboards** / 관리자 대시보드
   - Simple CRUD interfaces / 간단한 CRUD 인터페이스
   - Data visualization / 데이터 시각화

---

## Design Philosophy / 설계 철학

### 1. Developer Convenience First / 개발자 편의성 우선

**Priority Order / 우선순위**:
1. **Ease of Use** / 사용 용이성 (최우선)
2. **Code Readability** / 코드 가독성
3. **Developer Productivity** / 개발자 생산성
4. **Performance** / 성능 (중요하지만 4순위)

### 2. Extreme Simplicity / 극도의 간결함

**Before vs After Example / 이전 vs 이후 예제**:

```go
// ❌ Before: 50+ lines with standard net/http
func main() {
    mux := http.NewServeMux()

    // Register routes / 라우트 등록
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"message": "Hello"})
    })

    mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        // Parse query parameters / 쿼리 파라미터 파싱
        // Fetch data / 데이터 가져오기
        // Handle errors / 에러 처리
        // Encode JSON / JSON 인코딩
        // ... 40+ more lines
    })

    // Add CORS / CORS 추가
    corsHandler := func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }
            h.ServeHTTP(w, r)
        })
    }

    // Start server / 서버 시작
    log.Fatal(http.ListenAndServe(":8080", corsHandler(mux)))
}

// ✅ After: 5-10 lines with websvrutil
func main() {
    app := websvrutil.New()

    app.GET("/", func(c *websvrutil.Context) {
        c.JSON(200, map[string]string{"message": "Hello"})
    })

    app.GET("/users", listUsers)

    app.Run(":8080")
}
```

### 3. Smart Defaults / 스마트 기본값

- Auto CORS handling / 자동 CORS 처리
- Auto recovery from panics / 패닉 자동 복구
- Auto request logging / 자동 요청 로깅
- Auto template discovery / 자동 템플릿 발견
- Auto graceful shutdown / 자동 우아한 종료

### 4. Standard Library Compatible / 표준 라이브러리 호환

- Built on top of `net/http` / `net/http` 위에 구축
- Can use standard `http.Handler` / 표준 `http.Handler` 사용 가능
- No magic, just convenience / 마법이 아닌, 편리함

---

## Architecture / 아키텍처

### High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────────┐
│                     websvrutil.App                       │
│  (Main server instance / 메인 서버 인스턴스)              │
└─────────────────────────────────────────────────────────┘
                        │
        ┌───────────────┼───────────────┐
        │               │               │
        ▼               ▼               ▼
┌──────────────┐ ┌──────────────┐ ┌──────────────┐
│   Router     │ │  Middleware  │ │   Context    │
│   (Routes)   │ │   (Chain)    │ │  (Request)   │
└──────────────┘ └──────────────┘ └──────────────┘
        │               │               │
        ▼               ▼               ▼
┌──────────────┐ ┌──────────────┐ ┌──────────────┐
│   Handlers   │ │   Recovery   │ │   Response   │
│              │ │     CORS     │ │    Helpers   │
│              │ │    Logging   │ │              │
└──────────────┘ └──────────────┘ └──────────────┘
                        │
                        ▼
                ┌──────────────┐
                │   Template   │
                │    System    │
                └──────────────┘
```

### File Structure / 파일 구조

```
websvrutil/
├── websvrutil.go       # Package info, version / 패키지 정보, 버전
├── app.go              # App (server) struct and constructor / App 구조체 및 생성자
├── router.go           # Routing logic (RESTful routes) / 라우팅 로직
├── context.go          # Request context wrapper / 요청 컨텍스트 래퍼
├── response.go         # Response helpers (JSON, HTML, etc.) / 응답 헬퍼
├── middleware.go       # Built-in middleware / 내장 미들웨어
├── template.go         # Template system / 템플릿 시스템
├── handler.go          # Handler utilities / 핸들러 유틸리티
├── server.go           # Server management (start, stop, graceful) / 서버 관리
├── errors.go           # Error types / 에러 타입
└── options.go          # Configuration options / 설정 옵션
```

---

## Core Components / 핵심 컴포넌트

### 1. App (Server Instance) / 앱 (서버 인스턴스)

**Purpose / 목적**: Main server instance that manages routes, middleware, and server lifecycle.

라우트, 미들웨어, 서버 생명주기를 관리하는 메인 서버 인스턴스.

**Type Definition / 타입 정의**:
```go
type App struct {
    router      *Router
    middleware  []MiddlewareFunc
    templates   *TemplateEngine
    options     *Options
    server      *http.Server
}

// Constructor / 생성자
func New(opts ...Option) *App

// Methods / 메서드
func (a *App) Use(middleware ...MiddlewareFunc) *App
func (a *App) GET(path string, handler HandlerFunc) *App
func (a *App) POST(path string, handler HandlerFunc) *App
func (a *App) PUT(path string, handler HandlerFunc) *App
func (a *App) PATCH(path string, handler HandlerFunc) *App
func (a *App) DELETE(path string, handler HandlerFunc) *App
func (a *App) Static(prefix, root string) *App
func (a *App) Run(addr string) error
func (a *App) Shutdown(ctx context.Context) error
```

### 2. Router / 라우터

**Purpose / 목적**: Handle route registration and matching with path parameters.

경로 매개변수를 사용한 라우트 등록 및 매칭 처리.

**Type Definition / 타입 정의**:
```go
type Router struct {
    routes map[string]map[string]HandlerFunc  // method -> path -> handler
    prefix string
}

// Methods / 메서드
func (r *Router) Add(method, path string, handler HandlerFunc)
func (r *Router) Match(method, path string) (HandlerFunc, map[string]string, bool)
func (r *Router) Group(prefix string) *Router
```

**Route Patterns / 라우트 패턴**:
- Static: `/users`, `/api/v1/products`
- Parameters: `/users/:id`, `/posts/:postId/comments/:commentId`
- Wildcard: `/static/*filepath`

### 3. Context / 컨텍스트

**Purpose / 목적**: Request context that wraps http.Request and provides convenience methods.

http.Request를 래핑하고 편의 메서드를 제공하는 요청 컨텍스트.

**Type Definition / 타입 정의**:
```go
type Context struct {
    Request  *http.Request
    Response http.ResponseWriter
    Params   map[string]string  // Path parameters / 경로 매개변수
    Query    url.Values         // Query parameters / 쿼리 매개변수

    // Private fields / 비공개 필드
    app      *App
    index    int  // Middleware index / 미들웨어 인덱스
}

// Request Methods / 요청 메서드
func (c *Context) Param(key string) string
func (c *Context) QueryParam(key string) string
func (c *Context) QueryParamDefault(key, defaultValue string) string
func (c *Context) Bind(obj interface{}) error
func (c *Context) BindJSON(obj interface{}) error
func (c *Context) BindForm(obj interface{}) error
func (c *Context) Cookie(name string) (*http.Cookie, error)
func (c *Context) SetCookie(cookie *http.Cookie)
func (c *Context) Header(key string) string
func (c *Context) SetHeader(key, value string)

// Response Methods / 응답 메서드
func (c *Context) JSON(code int, obj interface{})
func (c *Context) HTML(code int, name string, data interface{})
func (c *Context) String(code int, format string, values ...interface{})
func (c *Context) Data(code int, contentType string, data []byte)
func (c *Context) File(filepath string)
func (c *Context) Redirect(code int, location string)
func (c *Context) Error(code int, message string)

// Middleware Methods / 미들웨어 메서드
func (c *Context) Next()
func (c *Context) Abort()
func (c *Context) AbortWithStatus(code int)
func (c *Context) AbortWithError(code int, err error)

// Context Storage / 컨텍스트 저장소
func (c *Context) Set(key string, value interface{})
func (c *Context) Get(key string) (interface{}, bool)
func (c *Context) MustGet(key string) interface{}
```

### 4. Middleware / 미들웨어

**Purpose / 목적**: Interceptors for request/response processing.

요청/응답 처리를 위한 인터셉터.

**Type Definition / 타입 정의**:
```go
type MiddlewareFunc func(HandlerFunc) HandlerFunc
type HandlerFunc func(*Context)
```

**Built-in Middleware / 내장 미들웨어**:
```go
// Recovery - Panic 복구
func Recovery() MiddlewareFunc

// Logger - 요청 로깅
func Logger() MiddlewareFunc

// CORS - Cross-Origin Resource Sharing
func CORS(opts ...CORSOption) MiddlewareFunc

// Auth - Authentication
func Auth(validator func(token string) bool) MiddlewareFunc

// RateLimit - Rate limiting
func RateLimit(requests int, duration time.Duration) MiddlewareFunc

// Compress - Response compression
func Compress() MiddlewareFunc

// Timeout - Request timeout
func Timeout(duration time.Duration) MiddlewareFunc
```

### 5. Template System / 템플릿 시스템

**Purpose / 목적**: Easy template management with auto-discovery and hot reload.

자동 발견 및 핫 리로드를 제공하는 쉬운 템플릿 관리.

**Type Definition / 타입 정의**:
```go
type TemplateEngine struct {
    templates map[string]*template.Template
    funcs     template.FuncMap
    root      string
    ext       string
    hotReload bool
}

// Methods / 메서드
func (t *TemplateEngine) LoadTemplates(root, ext string) error
func (t *TemplateEngine) Render(w io.Writer, name string, data interface{}) error
func (t *TemplateEngine) AddFunc(name string, fn interface{})
func (t *TemplateEngine) EnableHotReload()
```

**Template Features / 템플릿 기능**:
- Auto-discovery of `.html` files / `.html` 파일 자동 발견
- Nested template support / 중첩 템플릿 지원
- Layout system / 레이아웃 시스템
- Custom template functions / 사용자 정의 템플릿 함수
- Hot reload in development / 개발 시 핫 리로드

---

## API Design / API 설계

### Simple REST API Example / 간단한 REST API 예제

```go
package main

import (
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    app := websvrutil.New()

    // Global middleware / 전역 미들웨어
    app.Use(websvrutil.Logger())
    app.Use(websvrutil.Recovery())
    app.Use(websvrutil.CORS())

    // Routes / 라우트
    app.GET("/", homeHandler)
    app.GET("/users", listUsers)
    app.GET("/users/:id", getUser)
    app.POST("/users", createUser)
    app.PUT("/users/:id", updateUser)
    app.DELETE("/users/:id", deleteUser)

    // Start server / 서버 시작
    app.Run(":8080")
}

func homeHandler(c *websvrutil.Context) {
    c.JSON(200, map[string]string{
        "message": "Welcome to the API",
        "version": "v1.0.0",
    })
}

func listUsers(c *websvrutil.Context) {
    users := []User{
        {ID: 1, Name: "Alice"},
        {ID: 2, Name: "Bob"},
    }
    c.JSON(200, users)
}

func getUser(c *websvrutil.Context) {
    id := c.Param("id")
    user := findUser(id)
    if user == nil {
        c.Error(404, "User not found")
        return
    }
    c.JSON(200, user)
}

func createUser(c *websvrutil.Context) {
    var user User
    if err := c.BindJSON(&user); err != nil {
        c.Error(400, "Invalid request")
        return
    }
    // Save user / 사용자 저장
    c.JSON(201, user)
}
```

### Web Application with Templates / 템플릿을 사용한 웹 애플리케이션

```go
package main

import (
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    app := websvrutil.New(
        websvrutil.WithTemplates("./templates", ".html"),
        websvrutil.WithHotReload(true),
    )

    app.Use(websvrutil.Logger())
    app.Use(websvrutil.Recovery())

    // Static files / 정적 파일
    app.Static("/static", "./static")

    // Routes / 라우트
    app.GET("/", indexPage)
    app.GET("/about", aboutPage)
    app.GET("/contact", contactPage)
    app.POST("/contact", submitContact)

    app.Run(":8080")
}

func indexPage(c *websvrutil.Context) {
    c.HTML(200, "index.html", map[string]interface{}{
        "Title": "Home Page",
        "User":  getCurrentUser(c),
    })
}

func aboutPage(c *websvrutil.Context) {
    c.HTML(200, "about.html", map[string]interface{}{
        "Title": "About Us",
    })
}

func contactPage(c *websvrutil.Context) {
    c.HTML(200, "contact.html", map[string]interface{}{
        "Title": "Contact",
    })
}

func submitContact(c *websvrutil.Context) {
    var form ContactForm
    if err := c.BindForm(&form); err != nil {
        c.Error(400, "Invalid form data")
        return
    }
    // Process form / 폼 처리
    c.Redirect(303, "/thank-you")
}
```

### Route Groups / 라우트 그룹

```go
func main() {
    app := websvrutil.New()

    // API v1 / API v1
    v1 := app.Group("/api/v1")
    v1.Use(websvrutil.Auth(validateToken))
    {
        v1.GET("/users", listUsers)
        v1.POST("/users", createUser)

        // Admin routes / 관리자 라우트
        admin := v1.Group("/admin")
        admin.Use(requireAdmin)
        {
            admin.GET("/stats", getStats)
            admin.DELETE("/users/:id", deleteUser)
        }
    }

    app.Run(":8080")
}
```

---

## Template System / 템플릿 시스템

### Directory Structure / 디렉토리 구조

```
templates/
├── layouts/
│   ├── base.html       # Base layout / 기본 레이아웃
│   └── admin.html      # Admin layout / 관리자 레이아웃
├── partials/
│   ├── header.html     # Header partial / 헤더 파셜
│   ├── footer.html     # Footer partial / 푸터 파셜
│   └── nav.html        # Navigation partial / 네비게이션 파셜
└── pages/
    ├── index.html      # Home page / 홈 페이지
    ├── about.html      # About page / 소개 페이지
    └── contact.html    # Contact page / 연락처 페이지
```

### Layout System / 레이아웃 시스템

**Base Layout (layouts/base.html)**:
```html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    {{template "header" .}}

    <main>
        {{template "content" .}}
    </main>

    {{template "footer" .}}

    <script src="/static/js/app.js"></script>
</body>
</html>
```

**Page Template (pages/index.html)**:
```html
{{define "content"}}
<div class="container">
    <h1>Welcome, {{.User.Name}}</h1>
    <p>{{.Message}}</p>
</div>
{{end}}
```

### Custom Template Functions / 사용자 정의 템플릿 함수

```go
app := websvrutil.New()

// Add custom functions / 사용자 정의 함수 추가
app.AddTemplateFunc("formatDate", func(t time.Time) string {
    return t.Format("2006-01-02")
})

app.AddTemplateFunc("upper", strings.ToUpper)

app.AddTemplateFunc("safeHTML", func(s string) template.HTML {
    return template.HTML(s)
})
```

**Usage in template / 템플릿에서 사용**:
```html
<p>Posted on: {{formatDate .PostDate}}</p>
<h1>{{upper .Title}}</h1>
<div>{{safeHTML .Content}}</div>
```

---

## Error Handling / 에러 처리

### Error Types / 에러 타입

```go
type HTTPError struct {
    Code    int
    Message string
    Internal error  // Original error / 원본 에러
}

func (e *HTTPError) Error() string {
    return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}
```

### Error Handler / 에러 핸들러

```go
// Default error handler / 기본 에러 핸들러
func defaultErrorHandler(c *Context, err error) {
    if httpErr, ok := err.(*HTTPError); ok {
        c.JSON(httpErr.Code, map[string]interface{}{
            "error": httpErr.Message,
        })
    } else {
        c.JSON(500, map[string]interface{}{
            "error": "Internal Server Error",
        })
    }
}

// Custom error handler / 사용자 정의 에러 핸들러
app := websvrutil.New(
    websvrutil.WithErrorHandler(func(c *Context, err error) {
        // Custom error handling / 사용자 정의 에러 처리
        log.Printf("Error: %v", err)
        c.HTML(500, "error.html", map[string]interface{}{
            "Error": err.Error(),
        })
    }),
)
```

---

## Performance Considerations / 성능 고려사항

### Optimization Strategies / 최적화 전략

1. **Router Performance / 라우터 성능**
   - Use map for exact route matching / 정확한 라우트 매칭을 위해 맵 사용
   - Trie for pattern matching (if needed) / 패턴 매칭을 위한 Trie (필요시)
   - Cache compiled route patterns / 컴파일된 라우트 패턴 캐시

2. **Template Performance / 템플릿 성능**
   - Pre-compile all templates on startup / 시작 시 모든 템플릿 미리 컴파일
   - Cache parsed templates / 파싱된 템플릿 캐시
   - Optional hot reload only in development / 개발 시에만 선택적 핫 리로드

3. **Middleware Performance / 미들웨어 성능**
   - Minimize middleware chain / 미들웨어 체인 최소화
   - Use sync.Pool for Context reuse / Context 재사용을 위해 sync.Pool 사용

4. **Response Writing / 응답 작성**
   - Buffer response writing / 응답 작성 버퍼링
   - Use json.Encoder for streaming / 스트리밍을 위해 json.Encoder 사용

### Benchmarking / 벤치마킹

Target performance (NOT a priority, but good to have):

목표 성능 (우선순위는 아니지만, 있으면 좋음):

- Simple route: < 10μs per request / 간단한 라우트: 요청당 10μs 미만
- JSON response: < 20μs per request / JSON 응답: 요청당 20μs 미만
- Template render: < 100μs per request / 템플릿 렌더링: 요청당 100μs 미만

**Note / 참고**: Performance is important but secondary to developer experience.

성능은 중요하지만 개발자 경험에 부차적입니다.

---

## Configuration Options / 설정 옵션

### Options Pattern / 옵션 패턴

```go
type Options struct {
    // Server / 서버
    ReadTimeout     time.Duration
    WriteTimeout    time.Duration
    MaxHeaderBytes  int

    // Templates / 템플릿
    TemplateRoot    string
    TemplateExt     string
    HotReload       bool

    // Middleware / 미들웨어
    EnableLogger    bool
    EnableRecovery  bool
    EnableCORS      bool

    // Error Handling / 에러 처리
    ErrorHandler    func(*Context, error)

    // Static Files / 정적 파일
    StaticRoot      string
    StaticPrefix    string
}

// Option functions / 옵션 함수
func WithReadTimeout(d time.Duration) Option
func WithWriteTimeout(d time.Duration) Option
func WithTemplates(root, ext string) Option
func WithHotReload(enable bool) Option
func WithLogger(enable bool) Option
func WithRecovery(enable bool) Option
func WithCORS(enable bool) Option
func WithErrorHandler(handler func(*Context, error)) Option
func WithStatic(prefix, root string) Option
```

**Usage / 사용법**:
```go
app := websvrutil.New(
    websvrutil.WithReadTimeout(30*time.Second),
    websvrutil.WithTemplates("./views", ".html"),
    websvrutil.WithHotReload(true),
    websvrutil.WithLogger(true),
)
```

---

## Comparison with Other Frameworks / 다른 프레임워크와 비교

### vs Gin

| Feature / 기능 | Gin | websvrutil |
|---------------|-----|------------|
| **Speed / 속도** | ⭐⭐⭐⭐⭐ (매우 빠름) | ⭐⭐⭐ (충분히 빠름) |
| **Ease of Use / 사용 용이성** | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ (최우선) |
| **Template System / 템플릿 시스템** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ (Auto-discovery) |
| **Learning Curve / 학습 곡선** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ (매우 쉬움) |
| **Boilerplate / 보일러플레이트** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ (최소화) |

### vs Echo

| Feature / 기능 | Echo | websvrutil |
|---------------|------|------------|
| **Speed / 속도** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ |
| **Ease of Use / 사용 용이성** | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **Middleware / 미들웨어** | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ (더 간단) |
| **Standard Compat / 표준 호환** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |

### vs net/http (Standard Library)

| Feature / 기능 | net/http | websvrutil |
|---------------|----------|------------|
| **Simplicity / 간결함** | ⭐⭐ (장황함) | ⭐⭐⭐⭐⭐ |
| **Features / 기능** | ⭐⭐ (기본만) | ⭐⭐⭐⭐⭐ |
| **Boilerplate / 보일러플레이트** | ⭐ (매우 많음) | ⭐⭐⭐⭐⭐ (최소) |

---

## Design Decisions / 설계 결정사항

### 1. Why Not Use Existing Frameworks? / 왜 기존 프레임워크를 사용하지 않나?

**Reason / 이유**: Focus on developer convenience, not performance.

개발자 편의성에 집중, 성능이 아님.

- Gin/Echo: Great performance but more verbose / 뛰어난 성능이지만 더 장황함
- This package: Extreme simplicity for 99% use cases / 99% 사용 사례를 위한 극도의 간결함

### 2. Why Built-in Template System? / 왜 내장 템플릿 시스템?

**Reason / 이유**: Auto-discovery and hot reload out of the box.

즉시 사용 가능한 자동 발견 및 핫 리로드.

- No manual template registration / 수동 템플릿 등록 불필요
- Hot reload in development / 개발 시 핫 리로드
- Layout system / 레이아웃 시스템

### 3. Why Context Wrapper? / 왜 Context 래퍼?

**Reason / 이유**: Convenience methods for common operations.

일반적인 작업을 위한 편의 메서드.

- `c.JSON()` instead of manual json.Encode / 수동 json.Encode 대신
- `c.Param()` instead of manual URL parsing / 수동 URL 파싱 대신
- `c.BindJSON()` for automatic body parsing / 자동 본문 파싱

---

## Future Enhancements / 향후 개선사항

### Phase 1 (v1.11.x)
- ✅ Core router
- ✅ Basic handlers
- ✅ Middleware
- ✅ Template system

### Phase 2 (v1.12.x)
- WebSocket support / WebSocket 지원
- Session management / 세션 관리
- File upload helpers / 파일 업로드 헬퍼
- Form validation / 폼 검증

### Phase 3 (v1.13.x)
- Database helpers / 데이터베이스 헬퍼
- ORM integration / ORM 통합
- Cache helpers / 캐시 헬퍼

---

## Conclusion / 결론

The `websvrutil` package prioritizes **developer convenience** over raw performance, making it ideal for:

`websvrutil` 패키지는 순수 성능보다 **개발자 편의성**을 우선시하여 다음에 이상적입니다:

- Rapid prototyping / 빠른 프로토타이핑
- Small to medium web applications / 중소형 웹 애플리케이션
- REST APIs / REST API
- Admin dashboards / 관리자 대시보드
- Microservices / 마이크로서비스

**Core Philosophy / 핵심 철학**: "Make the developer happy, performance will follow."

"개발자를 행복하게 하면, 성능은 따라온다."
