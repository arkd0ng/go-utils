# websvrutil - Web Server Utilities / 웹 서버 유틸리티

**Version / 버전**: v1.11.002
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

## Current Features (v1.11.002) / 현재 기능

### App Struct / App 구조체

The main application instance that manages your web server.

웹 서버를 관리하는 주요 애플리케이션 인스턴스입니다.

**Methods / 메서드**:
- `New(opts ...Option) *App` - Create new app instance / 새 앱 인스턴스 생성
- `Use(middleware ...MiddlewareFunc) *App` - Add middleware / 미들웨어 추가
- `Run(addr string) error` - Start server / 서버 시작
- `Shutdown(ctx context.Context) error` - Graceful shutdown / 정상 종료
- `ServeHTTP(w http.ResponseWriter, r *http.Request)` - Implement http.Handler / http.Handler 구현

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

### Basic Server / 기본 서버

```go
package main

import (
    "log"
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    // Create app with defaults
    // 기본값으로 앱 생성
    app := websvrutil.New()

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

## Upcoming Features / 예정된 기능

The following features are planned for future releases:

다음 기능이 향후 릴리스에 계획되어 있습니다:

- **Router** (v1.11.003): HTTP routing with path parameters / 경로 매개변수가 있는 HTTP 라우팅
- **Context** (v1.11.004-005): Request context with parameter binding / 매개변수 바인딩이 있는 요청 컨텍스트
- **Middleware System** (v1.11.006-010): Built-in middleware (recovery, logger, CORS, auth) / 내장 미들웨어
- **Template System** (v1.11.011-015): Auto-discovery, layouts, hot reload / 자동 발견, 레이아웃, 핫 리로드
- **Advanced Features** (v1.11.016-020): File upload, static serving, cookie helpers / 파일 업로드, 정적 제공

## Development Status / 개발 상태

**Current Phase / 현재 단계**: Phase 1 - Core Foundation (v1.11.001-005)

**Progress / 진행 상황**:
- ✅ v1.11.001: Project setup and planning / 프로젝트 설정 및 계획
- ✅ v1.11.002: App & Options / 앱 및 옵션
- 📝 v1.11.003: Router / 라우터
- 📝 v1.11.004: Context (Part 1) / 컨텍스트 (1부)
- 📝 v1.11.005: Response Helpers / 응답 헬퍼

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
