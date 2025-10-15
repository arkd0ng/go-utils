# websvrutil - Web Server Utilities / 웹 서버 유틸리티

**Version / 버전**: v1.11.005
**Package / 패키지**: `github.com/arkd0ng/go-utils/websvrutil`

## Overview / 개요

The `websvrutil` package provides extreme simplicity web server utilities for Golang. It reduces 50+ lines of typical web server setup code to just 5 lines, prioritizing developer convenience over raw performance.

`websvrutil` 패키지는 Golang을 위한 극도로 간단한 웹 서버 유틸리티를 제공합니다. 일반적인 웹 서버 설정 코드 50줄 이상을 단 5줄로 줄여주며, 순수 성능보다 개발자 편의성을 우선시합니다.

## Design Philosophy / 설계 철학

- **Developer Convenience First** / **개발자 편의성 우선**: 50+ lines → 5 lines
- **Smart Defaults** / **스마트 기본값**: Zero configuration for 99% of use cases / 99% 사용 사례에 대한 제로 설정
- **Standard Library Compatible** / **표준 라이브러리 호환**: Built on `net/http`, no magic / `net/http` 기반, 마법 없음
- **Easy Middleware Chaining** / **쉬운 미들웨어 체이닝**: Simple and intuitive / 간단하고 직관적
- **Auto Template Discovery** / **자동 템플릿 발견**: Smart template loading and hot reload / 스마트 템플릿 로딩 및 핫 리로드

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/websvrutil
```

## Current Features (v1.11.005) / 현재 기능

### App Struct / App 구조체

The main application instance that manages your web server.

웹 서버를 관리하는 주요 애플리케이션 인스턴스입니다.

**Methods / 메서드**:
- `New(opts ...Option) *App` - Create new app instance / 새 앱 인스턴스 생성
- `Use(middleware ...MiddlewareFunc) *App` - Add middleware / 미들웨어 추가
- `GET/POST/PUT/PATCH/DELETE/OPTIONS/HEAD(pattern, handler)` - Register routes / 라우트 등록
- `NotFound(handler)` - Custom 404 handler / 커스텀 404 핸들러
- `Run(addr string) error` - Start server / 서버 시작
- `Shutdown(ctx context.Context) error` - Graceful shutdown / 정상 종료
- `ServeHTTP(w http.ResponseWriter, r *http.Request)` - Implement http.Handler / http.Handler 구현

### Router / 라우터

Fast HTTP request router with parameter and wildcard support.

매개변수 및 와일드카드 지원을 갖춘 빠른 HTTP 요청 라우터.

**Features / 기능**:
- HTTP method routing (GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD) / HTTP 메서드 라우팅
- Path parameters (`:id`, `:name`) / 경로 매개변수
- Wildcard routes (`*`) / 와일드카드 라우트
- Custom 404 handler / 커스텀 404 핸들러
- Thread-safe / 스레드 안전

**Pattern Syntax / 패턴 구문**:
- `/users` - Exact match / 정확한 일치
- `/users/:id` - Parameter (e.g., `/users/123`) / 매개변수
- `/files/*` - Wildcard (matches everything) / 와일드카드 (모든 것과 일치)

### Context / 컨텍스트

Request context for accessing path parameters, query strings, headers, and storing custom values.

경로 매개변수, 쿼리 문자열, 헤더에 액세스하고 커스텀 값을 저장하기 위한 요청 컨텍스트.

**Parameter Access / 매개변수 액세스**:
- `Param(name string) string` - Get path parameter / 경로 매개변수 가져오기
- `Params() map[string]string` - Get all parameters / 모든 매개변수 가져오기

**Custom Values / 커스텀 값**:
- `Set(key string, value interface{})` - Store value / 값 저장
- `Get(key string) (interface{}, bool)` - Retrieve value / 값 검색
- `MustGet(key string) interface{}` - Get or panic / 가져오거나 패닉
- `GetString(key string) string` - Get string value / 문자열 값 가져오기
- `GetInt(key string) int` - Get int value / int 값 가져오기
- `GetBool(key string) bool` - Get bool value / bool 값 가져오기

**Request Helpers / 요청 헬퍼**:
- `Query(key string) string` - Get query parameter / 쿼리 매개변수 가져오기
- `QueryDefault(key, defaultValue string) string` - Get with default / 기본값으로 가져오기
- `Header(key string) string` - Get request header / 요청 헤더 가져오기
- `Method() string` - Get HTTP method / HTTP 메서드 가져오기
- `Path() string` - Get URL path / URL 경로 가져오기
- `Context() context.Context` - Get request context / 요청 컨텍스트 가져오기
- `WithContext(ctx context.Context) *Context` - Replace context / 컨텍스트 교체

**Basic Response / 기본 응답**:
- `SetHeader(key, value string)` - Set response header / 응답 헤더 설정
- `Status(code int)` - Set status code / 상태 코드 설정
- `Write(data []byte) (int, error)` - Write response / 응답 작성
- `WriteString(s string) (int, error)` - Write string / 문자열 작성

**JSON Response / JSON 응답**:
- `JSON(code int, data interface{}) error` - Send JSON response / JSON 응답 전송
- `JSONPretty(code int, data interface{}) error` - Send pretty JSON / 보기 좋은 JSON 전송
- `JSONIndent(code int, data, prefix, indent string) error` - Custom indent JSON / 커스텀 들여쓰기 JSON
- `Error(code int, message string) error` - Send JSON error / JSON 에러 전송

**HTML Response / HTML 응답**:
- `HTML(code int, html string) error` - Send HTML response / HTML 응답 전송
- `HTMLTemplate(code int, tmpl string, data interface{}) error` - Render template / 템플릿 렌더링

**Text Response / 텍스트 응답**:
- `Text(code int, text string) error` - Send plain text / 일반 텍스트 전송
- `Textf(code int, format string, args ...interface{}) error` - Send formatted text / 형식화된 텍스트 전송

**Other Responses / 기타 응답**:
- `XML(code int, xml string) error` - Send XML response / XML 응답 전송
- `Redirect(code int, url string)` - HTTP redirect / HTTP 리다이렉트
- `NoContent()` - Send 204 No Content / 204 No Content 전송

**Helper Function / 헬퍼 함수**:
- `GetContext(r *http.Request) *Context` - Get Context from request / 요청에서 Context 가져오기

**Thread-safe / 스레드 안전**: All Context operations are protected by sync.RWMutex / 모든 Context 작업은 sync.RWMutex로 보호됩니다

### Options Pattern / 옵션 패턴

Flexible configuration using functional options.

함수형 옵션을 사용한 유연한 설정.

**Available Options / 사용 가능한 옵션**:

| Option / 옵션 | Default / 기본값 | Description / 설명 |
|---------------|------------------|-------------------|
| `WithReadTimeout(d time.Duration)` | 15s | Server read timeout / 서버 읽기 시간 초과 |
| `WithWriteTimeout(d time.Duration)` | 15s | Server write timeout / 서버 쓰기 시간 초과 |
| `WithIdleTimeout(d time.Duration)` | 60s | Server idle timeout / 서버 유휴 시간 초과 |
| `WithMaxHeaderBytes(n int)` | 1 MB | Maximum header size / 최대 헤더 크기 |
| `WithTemplateDir(dir string)` | "templates" | Template directory / 템플릿 디렉토리 |
| `WithStaticDir(dir string)` | "static" | Static files directory / 정적 파일 디렉토리 |
| `WithStaticPrefix(prefix string)` | "/static" | Static files URL prefix / 정적 파일 URL 접두사 |
| `WithAutoReload(enable bool)` | false | Auto template reload / 자동 템플릿 재로드 |
| `WithLogger(enable bool)` | true | Enable logger middleware / 로거 미들웨어 활성화 |
| `WithRecovery(enable bool)` | true | Enable recovery middleware / 복구 미들웨어 활성화 |

## Quick Start / 빠른 시작

### Basic Server with Routes / 라우트가 있는 기본 서버

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    // Create app with defaults
    // 기본값으로 앱 생성
    app := websvrutil.New()

    // Register routes
    // 라우트 등록
    app.GET("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome!")
    })

    app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
        // Get Context to access path parameters
        // 경로 매개변수 액세스를 위한 Context 가져오기
        ctx := websvrutil.GetContext(r)
        id := ctx.Param("id")
        fmt.Fprintf(w, "User ID: %s", id)
    })

    app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusCreated)
        fmt.Fprintf(w, "User created")
    })

    // Start server
    // 서버 시작
    if err := app.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
```

### Server with Custom Options / 커스텀 옵션을 사용한 서버

```go
package main

import (
    "log"
    "time"
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    // Create app with custom options
    // 커스텀 옵션으로 앱 생성
    app := websvrutil.New(
        websvrutil.WithReadTimeout(30 * time.Second),
        websvrutil.WithWriteTimeout(30 * time.Second),
        websvrutil.WithTemplateDir("views"),
        websvrutil.WithStaticDir("public"),
        websvrutil.WithAutoReload(true), // Enable in development
    )

    // Start server
    // 서버 시작
    if err := app.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
```

### Graceful Shutdown / 정상 종료

```go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    app := websvrutil.New()

    // Setup signal handling
    // 시그널 처리 설정
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    // Start server in goroutine
    // 고루틴에서 서버 시작
    go func() {
        if err := app.Run(":8080"); err != nil {
            log.Printf("Server error: %v", err)
        }
    }()

    // Wait for interrupt signal
    // 인터럽트 시그널 대기
    <-quit
    log.Println("Shutting down server...")

    // Graceful shutdown with 5 second timeout
    // 5초 타임아웃으로 정상 종료
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := app.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    log.Println("Server exited")
}
```

### Custom Middleware / 커스텀 미들웨어

```go
package main

import (
    "log"
    "net/http"
    "time"
    "github.com/arkd0ng/go-utils/websvrutil"
)

// Logging middleware example
// 로깅 미들웨어 예제
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)

        next.ServeHTTP(w, r)

        log.Printf("Completed in %v", time.Since(start))
    })
}

// Authentication middleware example
// 인증 미들웨어 예제
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")

        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Validate token here
        // 여기서 토큰 검증

        next.ServeHTTP(w, r)
    })
}

func main() {
    app := websvrutil.New()

    // Add middleware (executed in order)
    // 미들웨어 추가 (순서대로 실행)
    app.Use(loggingMiddleware)
    app.Use(authMiddleware)

    if err := app.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
```

### Context Usage / Context 사용

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    app := websvrutil.New()

    // Path parameters
    // 경로 매개변수
    app.GET("/users/:id/posts/:postId", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        userID := ctx.Param("id")
        postID := ctx.Param("postId")
        fmt.Fprintf(w, "User: %s, Post: %s", userID, postID)
    })

    // Query parameters
    // 쿼리 매개변수
    app.GET("/search", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        q := ctx.Query("q")
        page := ctx.QueryDefault("page", "1")
        fmt.Fprintf(w, "Query: %s, Page: %s", q, page)
    })

    // Custom values storage
    // 커스텀 값 저장
    app.GET("/user/:id", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        // Store custom values
        // 커스텀 값 저장
        ctx.Set("userId", ctx.Param("id"))
        ctx.Set("authenticated", true)

        // Retrieve values
        // 값 검색
        if authenticated := ctx.GetBool("authenticated"); authenticated {
            userID := ctx.GetString("userId")
            fmt.Fprintf(w, "Authenticated user: %s", userID)
        }
    })

    // Request headers
    // 요청 헤더
    app.GET("/api/data", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        authToken := ctx.Header("Authorization")
        contentType := ctx.Header("Content-Type")

        // Set response headers
        // 응답 헤더 설정
        ctx.SetHeader("X-API-Version", "1.0")
        ctx.SetHeader("Content-Type", "application/json")

        fmt.Fprintf(w, "Auth: %s, Type: %s", authToken, contentType)
    })

    if err := app.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
```

### Wildcard and Custom 404 / 와일드카드 및 커스텀 404

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    app := websvrutil.New()

    // Exact match
    // 정확한 일치
    app.GET("/users", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Users list")
    })

    // Parameter match
    // 매개변수 일치
    app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        id := ctx.Param("id")
        fmt.Fprintf(w, "User ID: %s", id)
    })

    // Wildcard match (catches all paths starting with /files/)
    // 와일드카드 일치 (/files/로 시작하는 모든 경로 포착)
    app.GET("/files/*", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "File: %s", r.URL.Path)
    })

    // Custom 404 handler
    // 커스텀 404 핸들러
    app.NotFound(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Page not found: %s", r.URL.Path)
    })

    if err := app.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
```

## Upcoming Features / 예정된 기능

The following features are planned for future releases:

다음 기능이 향후 릴리스에 계획되어 있습니다:

- **Middleware System** (v1.11.006-010): Built-in middleware (recovery, logger, CORS, auth) / 내장 미들웨어
- **Template System** (v1.11.011-015): Auto-discovery, layouts, hot reload / 자동 발견, 레이아웃, 핫 리로드
- **Advanced Features** (v1.11.016-020): File upload, static serving, cookie helpers / 파일 업로드, 정적 제공

## Development Status / 개발 상태

**Current Phase / 현재 단계**: Phase 1 - Core Foundation (v1.11.001-005)

**Progress / 진행 상황**:
- ✅ v1.11.001: Project setup and planning / 프로젝트 설정 및 계획
- ✅ v1.11.002: App & Options / 앱 및 옵션
- ✅ v1.11.003: Router / 라우터
- ✅ v1.11.004: Context (Part 1) / 컨텍스트 (1부)
- ✅ v1.11.005: Response Helpers / 응답 헬퍼

## Documentation / 문서

- **Design Plan** / **설계 계획**: [docs/websvrutil/DESIGN_PLAN.md](../docs/websvrutil/DESIGN_PLAN.md)
- **Work Plan** / **작업 계획**: [docs/websvrutil/WORK_PLAN.md](../docs/websvrutil/WORK_PLAN.md)
- **Development Guide** / **개발 가이드**: [docs/websvrutil/PACKAGE_DEVELOPMENT_GUIDE.md](../docs/websvrutil/PACKAGE_DEVELOPMENT_GUIDE.md)
- **Changelog** / **변경 로그**: [docs/CHANGELOG/CHANGELOG-v1.11.md](../docs/CHANGELOG/CHANGELOG-v1.11.md)

## Testing / 테스트

```bash
# Run all tests
# 모든 테스트 실행
go test ./websvrutil -v

# Run with coverage
# 커버리지와 함께 실행
go test ./websvrutil -cover

# Run benchmarks
# 벤치마크 실행
go test ./websvrutil -bench=.
```

## License / 라이선스

MIT License - see [LICENSE](../LICENSE) for details.

MIT 라이선스 - 자세한 내용은 [LICENSE](../LICENSE)를 참조하세요.
