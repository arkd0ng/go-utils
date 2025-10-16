# Websvrutil Package - User Manual / 사용자 매뉴얼

**Package**: `github.com/arkd0ng/go-utils/websvrutil`  
**Version**: v1.11.028  
**Last Updated**: 2025-10-16

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른-시작](#quick-start--빠른-시작)
4. [Application Setup / 애플리케이션-설정](#application-setup--애플리케이션-설정)
5. [Routing / 라우팅](#routing--라우팅)
6. [Context Helpers / 컨텍스트-헬퍼](#context-helpers--컨텍스트-헬퍼)
7. [Responses / 응답](#responses--응답)
8. [Middleware / 미들웨어](#middleware--미들웨어)
9. [Template Rendering / 템플릿-렌더링](#template-rendering--템플릿-렌더링)
10. [Session Management / 세션-관리](#session-management--세션-관리)
11. [File Handling & Static Assets / 파일-처리-및-정적-자산](#file-handling--static-assets--파일-처리-및-정적-자산)
12. [Graceful Shutdown / 우아한-종료](#graceful-shutdown--우아한-종료)
13. [FAQ](#faq)

---

## Introduction / 소개

`websvrutil` provides a lightweight HTTP toolkit that layers developer-friendly helpers on top of Go's `net/http`.  
`websvrutil`은 Go의 `net/http` 위에 개발자 친화적인 헬퍼를 더한 경량 HTTP 툴킷을 제공합니다.

It focuses on practical productivity: concise routing, convenient request/response helpers, and smart defaults.  
이 패키지는 실용적인 생산성에 초점을 맞추어 간결한 라우팅, 편리한 요청/응답 헬퍼, 스마트한 기본값을 제공합니다.

### Key Capabilities / 주요 기능

- **App lifecycle** helpers (`New`, `Run`, `RunWithGracefulShutdown`) with option-based configuration / 옵션 기반 설정이 가능한 **애플리케이션 생명주기** 헬퍼 (`New`, `Run`, `RunWithGracefulShutdown`)
- **Router & Groups** supporting GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD, and nested prefixes / GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD 및 중첩 접두사를 지원하는 **라우터와 그룹**
- **Context helpers** for params, queries, headers, cookies, and request-scoped storage / 파라미터, 쿼리, 헤더, 쿠키, 요청 범위 저장소를 위한 **컨텍스트 헬퍼**
- **Binding utilities** (`Bind`, `BindJSON`, `BindForm`, `BindQuery`) honoring size limits / 크기 제한을 준수하는 **데이터 바인딩 유틸리티** (`Bind`, `BindJSON`, `BindForm`, `BindQuery`)
- **Response helpers** for text, JSON (pretty/indent), HTML string, XML string, files, and standardized errors / 텍스트, JSON(예쁘게/들여쓰기), HTML 문자열, XML 문자열, 파일, 표준화된 에러 응답을 위한 **응답 헬퍼**
- **Template engine** with layouts, custom funcs, and optional auto-reload / 레이아웃, 사용자 정의 함수, 선택적 자동 재로드가 가능한 **템플릿 엔진**
- **Session store** using secure cookie IDs with configurable policies / 구성 가능한 정책을 가진 보안 쿠키 ID 기반 **세션 저장소**
- **Built-in middleware** including recovery, logging, CORS, request ID, timeout, auth, rate limiting, compression, security headers, redirects, and CSRF / Recovery, Logging, CORS, Request ID, Timeout, 인증, 레이트 리밋, 압축, 보안 헤더, 리다이렉트, CSRF 등을 포함한 **내장 미들웨어**
- **File helpers** for multipart uploads and static directory serving / 멀티파트 업로드와 정적 디렉터리 제공을 위한 **파일 헬퍼**

---

## Installation / 설치

### Prerequisites / 전제 조건

- Go 1.18 or newer (module-aware builds) / Go 1.18 이상 (모듈 기반 빌드)

### Install / 설치 방법

```bash
go get github.com/arkd0ng/go-utils/websvrutil
```

### Import / 임포트

```go
import "github.com/arkd0ng/go-utils/websvrutil"
```

---

## Quick Start / 빠른 시작

The snippet below spins up a basic server with logging and recovery enabled.  
아래 예제는 로깅과 복구 미들웨어를 활성화한 기본 서버를 실행합니다.

```go
package main

import (
    "net/http"

    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    app := websvrutil.New(
        websvrutil.WithLogger(true),  // Enable logging middleware / 로깅 미들웨어 활성화
        websvrutil.WithRecovery(true),// Enable recovery middleware / 복구 미들웨어 활성화
    )

    app.GET("/", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        ctx.Text(200, "Hello, websvrutil!") // Plain text response / 일반 텍스트 응답
    })

    if err := app.Run(":8080"); err != nil {
        panic(err)
    }
}
```

Create the app with `New`, register routes, and call `Run` to start listening.  
`New`로 앱을 만들고 라우트를 등록한 뒤 `Run`을 호출하여 서버를 실행합니다.

---

## Application Setup / 애플리케이션 설정

`New` accepts functional options so you can adjust server behavior without boilerplate.  
`New` 함수는 함수형 옵션을 받아 보일러플레이트 없이 서버 동작을 조정할 수 있습니다.

```go
app := websvrutil.New(
    websvrutil.WithReadTimeout(30*time.Second),   // Read timeout / 읽기 타임아웃
    websvrutil.WithWriteTimeout(30*time.Second),  // Write timeout / 쓰기 타임아웃
    websvrutil.WithIdleTimeout(2*time.Minute),    // Idle timeout / 유휴 타임아웃
    websvrutil.WithTemplateDir("templates"),     // Template directory / 템플릿 디렉터리
    websvrutil.WithStaticDir("./public"),        // Static file directory / 정적 파일 디렉터리
    websvrutil.WithStaticPrefix("/assets"),      // Static URL prefix / 정적 URL 접두사
    websvrutil.WithAutoReload(true),              // Template auto-reload / 템플릿 자동 재로드
    websvrutil.WithMaxBodySize(5<<20),            // Max body size / 최대 본문 크기 (5MB)
    websvrutil.WithMaxUploadSize(32<<20),         // Max upload size / 최대 업로드 크기 (32MB)
)
```

- `WithLogger` and `WithRecovery` toggle bundled middleware on or off.  
- `WithTemplateDir` creates a `TemplateEngine` so templates can load on demand.  
- `WithMaxBodySize` and `WithMaxUploadSize` enforce limits during binding and uploads.  

- `WithLogger`와 `WithRecovery`는 내장 미들웨어를 켜거나 끕니다.  
- `WithTemplateDir`는 템플릿을 필요할 때 로드할 수 있도록 `TemplateEngine`을 생성합니다.  
- `WithMaxBodySize`와 `WithMaxUploadSize`는 바인딩 및 업로드 시 크기 제한을 적용합니다.  

Add middleware with `Use` in registration order (first added = innermost).  
`Use`로 미들웨어를 추가하면 등록 순서대로 적용되며, 먼저 추가한 미들웨어가 안쪽에서 실행됩니다.

```go
app.Use(
    websvrutil.RequestID(),              // Request ID injection / 요청 ID 삽입
    websvrutil.Logger(),                 // Access logging / 접근 로깅
    websvrutil.RecoveryWithConfig(websvrutil.RecoveryConfig{LogStack: true}),
    websvrutil.CORSWithConfig(websvrutil.CORSConfig{AllowOrigins: []string{"https://example.com"}}),
)
```

---

## Routing / 라우팅

All HTTP verbs are exposed as fluent helpers on `App`.  
모든 HTTP 메서드는 `App`의 플루언트 헬퍼로 제공됩니다.

```go
app.GET("/status", statusHandler)
app.POST("/users", createUser)
app.PUT("/users/:id", updateUser)
app.PATCH("/users/:id", patchUser)
app.DELETE("/users/:id", deleteUser)
app.OPTIONS("/info", optionsHandler)
app.HEAD("/health", headHandler)
```

Route groups help share prefixes and middleware.  
라우트 그룹을 사용하면 접두사와 미들웨어를 쉽게 공유할 수 있습니다.

```go
api := app.Group("/api")
api.GET("/ping", pingHandler)

v1 := api.Group("/v1")
v1.Use(authMiddleware)
v1.GET("/profile", profileHandler)
```

Customize the not-found handler or expose static directories directly.  
사용자 정의 404 핸들러를 등록하거나 정적 디렉터리를 직접 노출할 수 있습니다.

```go
app.NotFound(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "resource not found", http.StatusNotFound)
})

app.Static("/static", "./public")
```

---

## Context Helpers / 컨텍스트 헬퍼

Retrieve the request-scoped `Context` inside handlers for convenience.  
핸들러 내부에서 요청 범위의 `Context`를 가져와 편리하게 사용할 수 있습니다.

```go
func profileHandler(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    id := ctx.Param("id")                          // Path parameter / 경로 매개변수
    verbose := ctx.Query("verbose")                // Query parameter / 쿼리 매개변수
    lang := ctx.QueryDefault("lang", "en")        // Default query / 기본 쿼리값

    ctx.Set("userID", id)                          // Store arbitrary data / 임의 데이터 저장
    storedID, _ := ctx.Get("userID")               // Retrieve stored data / 저장된 데이터 검색

    ctx.JSON(200, map[string]interface{}{
        "id":       id,
        "verbose": verbose,
        "lang":    lang,
        "agent":   ctx.UserAgent(),               // User-Agent header / 사용자 에이전트 헤더
        "ip":      ctx.ClientIP(),                // Client IP detection / 클라이언트 IP 확인
        "stored":  storedID,
    })
}
```

Popular helpers include `Param`, `Params`, `Method`, `Path`, `Query`, `QueryDefault`, `Header`, `SetHeader`, `AddHeader`, `HeaderExists`, `Cookie`, `CookieValue`, `GetCookie`, `SetCookie`, `SetCookieAdvanced`, `DeleteCookie`, `Set`, `Get`, `MustGet`, and typed getters (`GetString`, `GetInt`, `GetBool`, `GetInt64`, `GetFloat64`, `GetStringSlice`, `GetStringMap`).  
자주 사용하는 헬퍼로는 `Param`, `Params`, `Method`, `Path`, `Query`, `QueryDefault`, `Header`, `SetHeader`, `AddHeader`, `HeaderExists`, `Cookie`, `CookieValue`, `GetCookie`, `SetCookie`, `SetCookieAdvanced`, `DeleteCookie`, `Set`, `Get`, `MustGet`, 그리고 `GetString`, `GetInt`, `GetBool`, `GetInt64`, `GetFloat64`, `GetStringSlice`, `GetStringMap`와 같은 타입별 getter가 있습니다.

### Binding Data / 데이터 바인딩

`Bind` chooses an appropriate strategy based on `Content-Type`, while explicit helpers let you force a format.  
`Bind`는 `Content-Type`에 따라 적절한 전략을 선택하며, 명시적 헬퍼를 사용하면 원하는 포맷을 강제할 수 있습니다.

```go
type User struct {
    Name  string `json:"name" form:"name"`
    Email string `json:"email" form:"email"`
}

app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    var payload User
    if err := ctx.Bind(&payload); err != nil {          // Auto-detect binding / 자동 바인딩 선택
        ctx.ErrorJSON(http.StatusBadRequest, err.Error())
        return
    }

    ctx.SuccessJSON(http.StatusCreated, "created", payload) // Standard success payload / 표준 성공 응답
})
```

`BindJSON`, `BindForm`, and `BindQuery` are available when you need explicit control.  
명시적 제어가 필요하면 `BindJSON`, `BindForm`, `BindQuery`를 사용할 수 있습니다.

---

## Responses / 응답

Use the context to manage status codes and write response bodies.  
컨텍스트를 사용해 상태 코드를 설정하고 응답 본문을 작성합니다.

```go
ctx.Status(201)
ctx.Text(201, "created")
ctx.Textf(200, "Hello %s", name)
ctx.JSON(200, payload)
ctx.JSONIndent(200, payload, "", "  ")
ctx.JSONPretty(200, payload)
ctx.HTML(200, "<h1>Hi</h1>")
ctx.XML(200, "<note>OK</note>")
```

`Write`, `WriteString`, `File`, and `FileAttachment` cover byte streaming and downloads.  
`Write`, `WriteString`, `File`, `FileAttachment`을 사용하면 바이트 스트리밍과 다운로드를 처리할 수 있습니다.

Standardized helpers produce consistent error payloads.  
표준화된 헬퍼는 일관된 에러 페이로드를 생성합니다.

```go
ctx.AbortWithStatus(http.StatusUnauthorized)
ctx.AbortWithError(http.StatusBadRequest, "invalid request")
ctx.AbortWithJSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
ctx.ErrorJSON(http.StatusNotFound, "missing resource")
ctx.SuccessJSON(http.StatusOK, "ok", data)
ctx.NotFound()
ctx.Unauthorized()
ctx.Forbidden()
ctx.BadRequest()
ctx.InternalServerError()
```

---

## Middleware / 미들웨어

Register middleware with `Use`; the last added middleware executes first.  
`Use`로 미들웨어를 등록하면 마지막에 추가한 미들웨어가 가장 먼저 실행됩니다.

```go
app.Use(
    websvrutil.RequestIDWithConfig(websvrutil.RequestIDConfig{HeaderName: "X-Trace-ID"}),
    websvrutil.Timeout(5*time.Second),
    websvrutil.LoggerWithConfig(websvrutil.LoggerConfig{IncludeHeaders: true}),
)
```

Built-in middleware overview / 내장 미들웨어 개요:

- `Recovery`, `RecoveryWithConfig` – panic safety and custom logging / 패닉 안전성과 커스텀 로깅 지원
- `Logger`, `LoggerWithConfig` – request/response logging / 요청/응답 로깅
- `CORS`, `CORSWithConfig` – cross-origin access control / 교차 출처 접근 제어
- `RequestID`, `RequestIDWithConfig` – request ID propagation / 요청 ID 전파
- `Timeout`, `TimeoutWithConfig` – per-request deadlines / 요청별 데드라인 설정
- `BasicAuth`, `BasicAuthWithConfig` – simple HTTP authentication / 간단한 HTTP 인증
- `RateLimiter`, `RateLimiterWithConfig` – token bucket rate limiting / 토큰 버킷 레이트 리밋
- `Compression`, `CompressionWithConfig` – gzip compression / gzip 압축
- `SecureHeaders`, `SecureHeadersWithConfig` – security header defaults / 보안 헤더 기본값
- `BodyLimit`, `BodyLimitWithConfig` – middleware-level body cap / 미들웨어 수준 본문 제한
- `Static`, `StaticWithConfig` – static directory serving / 정적 디렉터리 제공
- `Redirect`, `RedirectWithConfig`, `HTTPSRedirect`, `WWWRedirect` – redirect helpers / 리다이렉트 헬퍼
- `CSRF`, `CSRFWithConfig` – CSRF token issuance & validation / CSRF 토큰 발급 및 검증

Custom middleware should match the same `func(http.Handler) http.Handler` signature.  
사용자 정의 미들웨어도 `func(http.Handler) http.Handler` 시그니처를 따라야 합니다.

---

## Template Rendering / 템플릿 렌더링

Initialize the template engine by specifying `WithTemplateDir` (and optionally enabling auto-reload).  
`WithTemplateDir`를 지정하고 필요하다면 자동 재로드를 활성화하여 템플릿 엔진을 초기화합니다.

```go
app := websvrutil.New(
    websvrutil.WithTemplateDir("web/templates"),
    websvrutil.WithAutoReload(true),
)

if err := app.LoadTemplates("*.html"); err != nil {
    log.Fatal(err)
}
```

Render templates inside handlers using `Render` or `RenderWithLayout`.  
핸들러 내부에서는 `Render` 또는 `RenderWithLayout`을 사용해 템플릿을 렌더링합니다.

```go
app.GET("/", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    if err := ctx.Render(http.StatusOK, "home.html", map[string]any{
        "Title":   "Home",
        "Message": "Welcome",
    }); err != nil {
        ctx.InternalServerError()
    }
})
```

Template engine features include layout directories, `AddTemplateFunc`, `AddTemplateFuncs`, `LoadTemplate`, `LoadTemplates`, `LoadAll`, `LoadLayout`, `LoadAllLayouts`, `ReloadTemplates`, `EnableAutoReload`, `DisableAutoReload`, and `IsAutoReloadEnabled`.  
템플릿 엔진 기능에는 레이아웃 디렉터리, `AddTemplateFunc`, `AddTemplateFuncs`, `LoadTemplate`, `LoadTemplates`, `LoadAll`, `LoadLayout`, `LoadAllLayouts`, `ReloadTemplates`, `EnableAutoReload`, `DisableAutoReload`, `IsAutoReloadEnabled`가 포함됩니다.

---

## Session Management / 세션 관리

`SessionStore` keeps sessions in memory and issues secure cookie IDs.  
`SessionStore`는 메모리에 세션을 저장하고 안전한 쿠키 ID를 발급합니다.

```go
store := websvrutil.NewSessionStore(websvrutil.DefaultSessionOptions())

app.GET("/login", func(w http.ResponseWriter, r *http.Request) {
    sess, _ := store.Get(r)      // Fetch or create session / 세션 조회 또는 생성
    sess.Set("user", "alice")   // Store data in session / 세션에 데이터 저장
    store.Save(w, sess)          // Persist cookie / 쿠키 저장

    ctx := websvrutil.GetContext(r)
    ctx.Text(http.StatusOK, "logged in")
})

app.GET("/me", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    sess, _ := store.Get(r)

    user := sess.GetString("user")
    if user == "" {
        ctx.Unauthorized()
        return
    }

    ctx.JSON(http.StatusOK, map[string]string{"user": user})
})
```

Adjust behavior with `SessionOptions` (cookie name, expiration, SameSite, cleanup interval, path, domain, secure flag, HTTP-only flag).  
`SessionOptions`(쿠키 이름, 만료, SameSite, 정리 주기, Path, Domain, Secure, HttpOnly)으로 동작을 원하는 대로 조정할 수 있습니다.

---

## File Handling & Static Assets / 파일 처리 및 정적 자산

Multipart helpers simplify uploads and reuse option-based size limits.  
멀티파트 헬퍼는 업로드를 단순화하고 옵션 기반 크기 제한을 재사용합니다.

```go
app.POST("/upload", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    header, err := ctx.FormFile("file")           // First file / 첫 번째 파일
    if err != nil {
        ctx.ErrorJSON(http.StatusBadRequest, err.Error())
        return
    }

    if err := ctx.SaveUploadedFile(header, "./uploads/"+header.Filename); err != nil {
        ctx.InternalServerError()
        return
    }

    ctx.SuccessJSON(http.StatusOK, "uploaded", header.Filename)
})
```

- `MultipartForm` parses and caches multipart payloads while respecting `WithMaxUploadSize`.  
- `File` streams files inline; `FileAttachment` forces download with a custom filename.  

- `MultipartForm`은 `WithMaxUploadSize`를 준수하며 멀티파트 페이로드를 파싱하고 캐시합니다.  
- `File`은 파일을 인라인으로 전송하고 `FileAttachment`는 사용자 정의 파일명으로 다운로드를 강제합니다.  

`app.Static(prefix, dir)` registers middleware-friendly static serving for directories.  
`app.Static(prefix, dir)`로 디렉터리에 대한 미들웨어 기반 정적 서빙을 구성할 수 있습니다.

---

## Graceful Shutdown / 우아한 종료

`RunWithGracefulShutdown` starts the server and listens for `SIGINT`/`SIGTERM`, applying a timeout before calling `Shutdown`.  
`RunWithGracefulShutdown`은 서버를 시작하고 `SIGINT`/`SIGTERM`을 감지하여 타임아웃 후 `Shutdown`을 호출합니다.

```go
if err := app.RunWithGracefulShutdown(":8080", 30*time.Second); err != nil {
    log.Fatal(err)
}
```

Use `Shutdown` manually when integrating with custom signal handling or orchestrators.  
사용자 정의 신호 처리나 오케스트레이터에 통합할 때는 `Shutdown`을 직접 호출하면 됩니다.

---

## FAQ

**Q1. How do I access the raw `http.ResponseWriter` or `*http.Request`? / Q1. 원래의 `http.ResponseWriter`나 `*http.Request`는 어떻게 사용하나요?**  
Use the handler parameters directly or reference `ctx.ResponseWriter` and `ctx.Request`.  
핸들러 매개변수를 직접 사용하거나 `ctx.ResponseWriter`, `ctx.Request`를 참조하면 됩니다.

**Q2. How can I reload templates automatically? / Q2. 템플릿을 자동으로 다시 로드하려면 어떻게 하나요?**  
Enable `WithAutoReload(true)` when creating the app, or call `TemplateEngine().EnableAutoReload()` for manual control.  
앱 생성 시 `WithAutoReload(true)`를 사용하거나, 필요할 때 `TemplateEngine().EnableAutoReload()`를 호출하세요.

**Q3. Can I mix standard handlers with websvrutil middleware? / Q3. 표준 핸들러와 websvrutil 미들웨어를 함께 사용할 수 있나요?**  
Yes. Register any `http.Handler` or `http.HandlerFunc`; inside the handler call `websvrutil.GetContext(r)` to access helpers.  
가능합니다. 모든 `http.Handler` 또는 `http.HandlerFunc`를 등록하고, 핸들러 내부에서 `websvrutil.GetContext(r)`를 호출해 헬퍼를 사용하세요.

**Q4. How do I test handlers? / Q4. 핸들러는 어떻게 테스트하나요?**  
Use `httptest.NewRequest` and `httptest.NewRecorder`, then call `app.ServeHTTP`.  
`httptest.NewRequest`와 `httptest.NewRecorder`를 사용한 뒤 `app.ServeHTTP`를 호출하세요.

---

Enjoy building with `websvrutil`!  
`websvrutil`과 함께 즐거운 개발 되세요!
