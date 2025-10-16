# Websvrutil Package - User Manual / 사용자 매뉴얼

**Package**: `github.com/arkd0ng/go-utils/websvrutil`
**Version**: v1.11.023
**Last Updated**: 2025-10-16

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Core Concepts / 핵심 개념](#core-concepts--핵심-개념)
5. [Routing / 라우팅](#routing--라우팅)
6. [Request Handling / 요청 처리](#request-handling--요청-처리)
7. [Response Handling / 응답 처리](#response-handling--응답-처리)
8. [Middleware / 미들웨어](#middleware--미들웨어)
9. [Template Rendering / 템플릿 렌더링](#template-rendering--템플릿-렌더링)
10. [Session Management / 세션 관리](#session-management--세션-관리)
11. [Static Files / 정적 파일](#static-files--정적-파일)
12. [File Upload / 파일 업로드](#file-upload--파일-업로드)
13. [Error Handling / 에러 처리](#error-handling--에러-처리)
14. [Graceful Shutdown / 우아한 종료](#graceful-shutdown--우아한-종료)
15. [Best Practices / 모범 사례](#best-practices--모범-사례)
16. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)
17. [FAQ](#faq)

---

## Introduction / 소개

The `websvrutil` package provides a lightweight, developer-friendly HTTP framework for Go applications. It offers an intuitive API similar to popular web frameworks while maintaining the simplicity and performance of Go's standard library.

`websvrutil` 패키지는 Go 애플리케이션을 위한 가볍고 개발자 친화적인 HTTP 프레임워크를 제공합니다. 인기 있는 웹 프레임워크와 유사한 직관적인 API를 제공하면서도 Go 표준 라이브러리의 단순성과 성능을 유지합니다.

### Key Features / 주요 기능

- **Simple Routing**: HTTP method-based routing with path parameters
- **Request Binding**: Automatic JSON/XML parsing and validation
- **Template Rendering**: Built-in HTML template support
- **Middleware Support**: Chainable middleware system
- **Session Management**: Cookie-based sessions with automatic cleanup
- **Static File Serving**: Efficient static file and directory serving
- **File Upload**: Multipart form handling with file upload support
- **Error Handling**: Standardized error responses
- **Graceful Shutdown**: Signal-based graceful server shutdown

- **간단한 라우팅**: HTTP 메서드 기반 라우팅과 경로 매개변수
- **요청 바인딩**: 자동 JSON/XML 파싱 및 검증
- **템플릿 렌더링**: 내장 HTML 템플릿 지원
- **미들웨어 지원**: 체이닝 가능한 미들웨어 시스템
- **세션 관리**: 자동 정리 기능이 있는 쿠키 기반 세션
- **정적 파일 서빙**: 효율적인 정적 파일 및 디렉토리 서빙
- **파일 업로드**: 파일 업로드 지원이 있는 멀티파트 폼 처리
- **에러 처리**: 표준화된 에러 응답
- **우아한 종료**: 신호 기반 우아한 서버 종료

### Design Philosophy / 설계 철학

The websvrutil package follows the **"Developer Convenience First"** philosophy:

websvrutil 패키지는 **"개발자 편의 우선"** 철학을 따릅니다:

1. **Intuitive API**: Method names and signatures that are easy to remember
2. **Minimal Boilerplate**: Reduce repetitive code patterns
3. **Safe Defaults**: Sensible defaults that work for most use cases
4. **Type Safety**: Leverage Go's type system for compile-time safety
5. **Performance**: Built on top of Go's standard `net/http` for performance

1. **직관적인 API**: 기억하기 쉬운 메서드 이름과 시그니처
2. **최소한의 보일러플레이트**: 반복적인 코드 패턴 감소
3. **안전한 기본값**: 대부분의 사용 사례에 적합한 합리적인 기본값
4. **타입 안전성**: 컴파일 타임 안전성을 위한 Go 타입 시스템 활용
5. **성능**: 성능을 위해 Go의 표준 `net/http` 위에 구축

---

## Installation / 설치

### Prerequisites / 전제 조건

- Go 1.18 or higher (for generics support)
- Go 1.18 이상 (제네릭 지원)

### Installing the Package / 패키지 설치

```bash
go get github.com/arkd0ng/go-utils/websvrutil
```

### Importing / 임포트

```go
import "github.com/arkd0ng/go-utils/websvrutil"
```

---

## Quick Start / 빠른 시작

### Example 1: Hello World / 예제 1: Hello World

```go
package main

import (
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    // Create new app / 새 앱 생성
    app := websvrutil.New()

    // Define route / 라우트 정의
    app.GET("/", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        ctx.String(200, "Hello, World!")
    })

    // Start server / 서버 시작
    app.Run(":8080")
}
```

### Example 2: JSON API / 예제 2: JSON API

```go
package main

import (
    "github.com/arkd0ng/go-utils/websvrutil"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    app := websvrutil.New()

    // GET endpoint / GET 엔드포인트
    app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        id := ctx.Param("id")

        user := User{ID: 1, Name: "John Doe"}
        ctx.JSON(200, user)
    })

    // POST endpoint / POST 엔드포인트
    app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        var user User
        if err := ctx.BindJSON(&user); err != nil {
            ctx.ErrorJSON(400, "Invalid JSON")
            return
        }

        ctx.SuccessJSON(201, "User created", user)
    })

    app.Run(":8080")
}
```

### Example 3: Template Rendering / 예제 3: 템플릿 렌더링

```go
package main

import (
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    app := websvrutil.New()

    // Load templates / 템플릿 로드
    app.LoadHTMLGlob("templates/*")

    // Render template / 템플릿 렌더링
    app.GET("/", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        ctx.HTML(200, "index.html", map[string]interface{}{
            "Title": "Home Page",
            "Name":  "John",
        })
    })

    app.Run(":8080")
}
```

---

## Core Concepts / 핵심 개념

### App / 앱

The `App` is the main application instance that handles routing, middleware, and server lifecycle.

`App`은 라우팅, 미들웨어, 서버 생명주기를 처리하는 메인 애플리케이션 인스턴스입니다.

```go
// Create new app / 새 앱 생성
app := websvrutil.New()

// Configure app / 앱 설정
app.LoadHTMLGlob("templates/*")
app.Static("/static", "./public")

// Start server / 서버 시작
app.Run(":8080")
```

### Context / 컨텍스트

The `Context` provides access to the request and response, along with helper methods for common operations.

`Context`는 요청과 응답에 대한 액세스를 제공하며 일반적인 작업을 위한 헬퍼 메서드를 제공합니다.

```go
app.GET("/user/:id", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Get path parameter / 경로 매개변수 가져오기
    id := ctx.Param("id")

    // Get query parameter / 쿼리 매개변수 가져오기
    name := ctx.Query("name")

    // Send JSON response / JSON 응답 전송
    ctx.JSON(200, map[string]string{"id": id, "name": name})
})
```

### Router / 라우터

The `Router` manages route registration and matching.

`Router`는 라우트 등록 및 매칭을 관리합니다.

```go
// HTTP method-based routing / HTTP 메서드 기반 라우팅
app.GET("/users", listUsers)
app.POST("/users", createUser)
app.PUT("/users/:id", updateUser)
app.DELETE("/users/:id", deleteUser)

// Route groups / 라우트 그룹
api := app.Group("/api")
api.GET("/users", listUsers)
api.GET("/posts", listPosts)
```

---

## Routing / 라우팅

### Basic Routing / 기본 라우팅

```go
app := websvrutil.New()

// GET request / GET 요청
app.GET("/", homeHandler)

// POST request / POST 요청
app.POST("/submit", submitHandler)

// PUT request / PUT 요청
app.PUT("/update", updateHandler)

// DELETE request / DELETE 요청
app.DELETE("/delete", deleteHandler)

// PATCH request / PATCH 요청
app.PATCH("/patch", patchHandler)

// HEAD request / HEAD 요청
app.HEAD("/head", headHandler)

// OPTIONS request / OPTIONS 요청
app.OPTIONS("/options", optionsHandler)
```

### Path Parameters / 경로 매개변수

```go
// Single parameter / 단일 매개변수
app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    id := ctx.Param("id")
    ctx.String(200, "User ID: %s", id)
})

// Multiple parameters / 다중 매개변수
app.GET("/posts/:category/:id", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    category := ctx.Param("category")
    id := ctx.Param("id")
    ctx.String(200, "Category: %s, ID: %s", category, id)
})
```

### Query Parameters / 쿼리 매개변수

```go
// GET /search?q=golang&page=2
app.GET("/search", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Get single query parameter / 단일 쿼리 매개변수 가져오기
    query := ctx.Query("q")

    // Get with default value / 기본값과 함께 가져오기
    page := ctx.DefaultQuery("page", "1")

    ctx.String(200, "Query: %s, Page: %s", query, page)
})
```

### Route Groups / 라우트 그룹

```go
app := websvrutil.New()

// API v1 group / API v1 그룹
v1 := app.Group("/api/v1")
{
    v1.GET("/users", listUsersV1)
    v1.POST("/users", createUserV1)
    v1.GET("/users/:id", getUserV1)
}

// API v2 group / API v2 그룹
v2 := app.Group("/api/v2")
{
    v2.GET("/users", listUsersV2)
    v2.POST("/users", createUserV2)
    v2.GET("/users/:id", getUserV2)
}

// Admin group / 관리자 그룹
admin := app.Group("/admin")
{
    admin.GET("/dashboard", adminDashboard)
    admin.GET("/users", adminUsers)
}
```

---

## Request Handling / 요청 처리

### Reading Request Body / 요청 본문 읽기

#### JSON Binding / JSON 바인딩

```go
type User struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=130"`
}

app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    var user User
    if err := ctx.BindJSON(&user); err != nil {
        ctx.ErrorJSON(400, "Invalid JSON: " + err.Error())
        return
    }

    ctx.SuccessJSON(201, "User created", user)
})
```

#### XML Binding / XML 바인딩

```go
type Product struct {
    XMLName xml.Name `xml:"product"`
    Name    string   `xml:"name"`
    Price   float64  `xml:"price"`
}

app.POST("/products", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    var product Product
    if err := ctx.BindXML(&product); err != nil {
        ctx.ErrorJSON(400, "Invalid XML: " + err.Error())
        return
    }

    ctx.XML(201, product)
})
```

### Request Headers / 요청 헤더

```go
app.GET("/headers", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Get single header / 단일 헤더 가져오기
    userAgent := ctx.GetHeader("User-Agent")

    // Get all headers / 모든 헤더 가져오기
    headers := ctx.GetAllHeaders()

    // Check if header exists / 헤더 존재 확인
    hasAuth := ctx.HasHeader("Authorization")

    ctx.JSON(200, map[string]interface{}{
        "user_agent": userAgent,
        "headers":    headers,
        "has_auth":   hasAuth,
    })
})
```

### Request Method Helpers / 요청 메서드 헬퍼

```go
app.Any("/endpoint", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Check request method / 요청 메서드 확인
    if ctx.IsGET() {
        ctx.String(200, "GET request")
    } else if ctx.IsPOST() {
        ctx.String(200, "POST request")
    } else if ctx.IsPUT() {
        ctx.String(200, "PUT request")
    }

    // Check request type / 요청 타입 확인
    if ctx.IsAjax() {
        ctx.String(200, "AJAX request")
    }

    if ctx.IsWebSocket() {
        ctx.String(200, "WebSocket upgrade request")
    }
})
```

### Content Negotiation / 컨텐츠 협상

```go
app.GET("/data", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    data := map[string]interface{}{
        "message": "Hello",
        "value":   123,
    }

    // Check Accept header / Accept 헤더 확인
    if ctx.AcceptsJSON() {
        ctx.JSON(200, data)
    } else if ctx.AcceptsXML() {
        ctx.XML(200, data)
    } else if ctx.AcceptsHTML() {
        ctx.HTML(200, "data.html", data)
    } else {
        ctx.String(200, "Unsupported content type")
    }
})
```

---

## Response Handling / 응답 처리

### String Response / 문자열 응답

```go
app.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.String(200, "Hello, World!")
})

app.GET("/formatted", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    name := "John"
    ctx.String(200, "Hello, %s!", name)
})
```

### JSON Response / JSON 응답

```go
app.GET("/json", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Simple JSON / 간단한 JSON
    ctx.JSON(200, map[string]string{
        "message": "Success",
    })
})

app.GET("/user", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Struct to JSON / 구조체를 JSON으로
    user := User{ID: 1, Name: "John"}
    ctx.JSON(200, user)
})
```

### XML Response / XML 응답

```go
type Book struct {
    XMLName xml.Name `xml:"book"`
    Title   string   `xml:"title"`
    Author  string   `xml:"author"`
}

app.GET("/xml", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    book := Book{Title: "Go Programming", Author: "John Doe"}
    ctx.XML(200, book)
})
```

### HTML Response / HTML 응답

```go
app.GET("/page", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    ctx.HTML(200, "page.html", map[string]interface{}{
        "Title":   "My Page",
        "Message": "Welcome!",
    })
})
```

### File Response / 파일 응답

```go
// Send file inline / 파일 인라인 전송
app.GET("/view", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.File("./files/document.pdf")
})

// Send file as attachment / 파일을 첨부 파일로 전송
app.GET("/download", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.FileAttachment("./files/report.pdf", "monthly-report.pdf")
})
```

### Data Response / 데이터 응답

```go
app.GET("/binary", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f} // "Hello"
    ctx.Data(200, "application/octet-stream", data)
})
```

### Response Headers / 응답 헤더

```go
app.GET("/headers", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Set single header / 단일 헤더 설정
    ctx.SetHeader("X-Custom-Header", "value")

    // Set multiple headers / 다중 헤더 설정
    ctx.SetHeaders(map[string]string{
        "X-Request-ID": "12345",
        "X-Version":    "1.0",
    })

    ctx.String(200, "Headers set")
})
```

### Error Responses / 에러 응답

```go
app.GET("/error", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Standard error response / 표준 에러 응답
    ctx.ErrorJSON(400, "Invalid request")
    // {"error":"Invalid request","status":400,"success":false}
})

app.GET("/success", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Success response / 성공 응답
    ctx.SuccessJSON(200, "Operation completed", map[string]int{"count": 5})
    // {"message":"Operation completed","data":{"count":5},"status":200,"success":true}
})

// HTTP error shortcuts / HTTP 에러 단축
app.GET("/not-found", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.NotFound() // 404
})

app.GET("/unauthorized", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.Unauthorized() // 401
})

app.GET("/forbidden", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.Forbidden() // 403
})

app.GET("/bad-request", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.BadRequest() // 400
})

app.GET("/server-error", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.InternalServerError() // 500
})
```

### Redirect / 리다이렉트

```go
app.GET("/redirect", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.Redirect(302, "/new-location")
})
```

---

## Middleware / 미들웨어

### Using Middleware / 미들웨어 사용

```go
// Global middleware / 전역 미들웨어
app.Use(LoggerMiddleware)
app.Use(RecoveryMiddleware)

// Route-specific middleware / 라우트별 미들웨어
app.GET("/admin", AuthMiddleware, adminHandler)

// Group middleware / 그룹 미들웨어
admin := app.Group("/admin")
admin.Use(AuthMiddleware)
admin.GET("/dashboard", dashboardHandler)
```

### Logger Middleware / 로거 미들웨어

```go
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Call next handler / 다음 핸들러 호출
        next.ServeHTTP(w, r)

        // Log request / 요청 로그
        duration := time.Since(start)
        log.Printf("%s %s - %v", r.Method, r.URL.Path, duration)
    })
}

app.Use(LoggerMiddleware)
```

### Recovery Middleware / 복구 미들웨어

```go
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic: %v", err)
                http.Error(w, "Internal Server Error", 500)
            }
        }()

        next.ServeHTTP(w, r)
    })
}

app.Use(RecoveryMiddleware)
```

### Authentication Middleware / 인증 미들웨어

```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        // Check authorization header / 인증 헤더 확인
        token := ctx.GetHeader("Authorization")
        if token == "" {
            ctx.Unauthorized()
            return
        }

        // Validate token / 토큰 검증
        if !isValidToken(token) {
            ctx.Forbidden()
            return
        }

        next.ServeHTTP(w, r)
    })
}

app.GET("/protected", AuthMiddleware, protectedHandler)
```

### CORS Middleware / CORS 미들웨어

```go
func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        // Set CORS headers / CORS 헤더 설정
        ctx.SetHeader("Access-Control-Allow-Origin", "*")
        ctx.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        ctx.SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight / 프리플라이트 처리
        if r.Method == "OPTIONS" {
            ctx.Status(204)
            return
        }

        next.ServeHTTP(w, r)
    })
}

app.Use(CORSMiddleware)
```

---

## Template Rendering / 템플릿 렌더링

### Loading Templates / 템플릿 로드

```go
app := websvrutil.New()

// Load all templates from directory / 디렉토리에서 모든 템플릿 로드
app.LoadHTMLGlob("templates/*")

// Load specific template files / 특정 템플릿 파일 로드
app.LoadHTMLFiles("templates/index.html", "templates/about.html")
```

### Rendering Templates / 템플릿 렌더링

**Template file (templates/index.html):**

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>Welcome, {{.Name}}!</h1>
    <p>{{.Message}}</p>
</body>
</html>
```

**Go code:**

```go
app.GET("/", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    ctx.HTML(200, "index.html", map[string]interface{}{
        "Title":   "Home Page",
        "Name":    "John",
        "Message": "Welcome to our website!",
    })
})
```

### Custom Template Functions / 커스텀 템플릿 함수

```go
// Add custom template functions / 커스텀 템플릿 함수 추가
app.SetFuncMap(template.FuncMap{
    "upper": strings.ToUpper,
    "lower": strings.ToLower,
    "formatDate": func(t time.Time) string {
        return t.Format("2006-01-02")
    },
})

app.LoadHTMLGlob("templates/*")
```

**Template usage:**

```html
<h1>{{upper .Title}}</h1>
<p>Date: {{formatDate .CreatedAt}}</p>
```

---

## Session Management / 세션 관리

### Creating Session Store / 세션 저장소 생성

```go
// Create session store with default options / 기본 옵션으로 세션 저장소 생성
store := websvrutil.NewSessionStore(websvrutil.DefaultSessionOptions())

// Create with custom options / 커스텀 옵션으로 생성
store := websvrutil.NewSessionStore(websvrutil.SessionOptions{
    CookieName:  "my_session",
    MaxAge:      12 * time.Hour,
    Secure:      true,
    HttpOnly:    true,
    SameSite:    http.SameSiteStrictMode,
    CleanupTime: 10 * time.Minute,
    Path:        "/",
    Domain:      "example.com",
})
```

### Using Sessions / 세션 사용

```go
var sessionStore *websvrutil.SessionStore

func init() {
    sessionStore = websvrutil.NewSessionStore(websvrutil.DefaultSessionOptions())
}

// Set session data / 세션 데이터 설정
app.POST("/login", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Get or create session / 세션 가져오기 또는 생성
    session, _ := sessionStore.Get(r)

    // Store user data / 사용자 데이터 저장
    session.Set("user_id", 123)
    session.Set("username", "john")
    session.Set("is_admin", true)

    // Save session / 세션 저장
    sessionStore.Save(w, session)

    ctx.SuccessJSON(200, "Logged in", nil)
})

// Get session data / 세션 데이터 가져오기
app.GET("/profile", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Get session / 세션 가져오기
    session, err := sessionStore.Get(r)
    if err != nil {
        ctx.Unauthorized()
        return
    }

    // Retrieve values / 값 가져오기
    userID := session.GetInt("user_id")
    username := session.GetString("username")
    isAdmin := session.GetBool("is_admin")

    ctx.JSON(200, map[string]interface{}{
        "user_id":  userID,
        "username": username,
        "is_admin": isAdmin,
    })
})

// Destroy session / 세션 삭제
app.POST("/logout", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    sessionStore.Destroy(w, r)

    ctx.SuccessJSON(200, "Logged out", nil)
})
```

### Session Data Methods / 세션 데이터 메서드

```go
session, _ := sessionStore.Get(r)

// Set values / 값 설정
session.Set("key", "value")

// Get values / 값 가져오기
value, exists := session.Get("key")

// Type-safe getters / 타입 안전 getter
str := session.GetString("name")
num := session.GetInt("count")
flag := session.GetBool("enabled")

// Delete value / 값 삭제
session.Delete("key")

// Clear all values / 모든 값 지우기
session.Clear()
```

---

## Static Files / 정적 파일

### Serving Static Files / 정적 파일 서빙

```go
app := websvrutil.New()

// Serve static files from directory / 디렉토리에서 정적 파일 서빙
app.Static("/static", "./public")
// Access: /static/css/style.css → ./public/css/style.css

// Multiple static directories / 여러 정적 디렉토리
app.Static("/css", "./assets/css")
app.Static("/js", "./assets/js")
app.Static("/images", "./assets/images")
```

### Directory Structure Example / 디렉토리 구조 예제

```
project/
├── main.go
├── public/
│   ├── css/
│   │   └── style.css
│   ├── js/
│   │   └── app.js
│   └── images/
│       └── logo.png
└── templates/
    └── index.html
```

**Go code:**

```go
app.Static("/static", "./public")

// Access files:
// http://localhost:8080/static/css/style.css
// http://localhost:8080/static/js/app.js
// http://localhost:8080/static/images/logo.png
```

---

## File Upload / 파일 업로드

### Single File Upload / 단일 파일 업로드

```go
app.POST("/upload", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Get uploaded file / 업로드된 파일 가져오기
    file, header, err := ctx.FormFile("file")
    if err != nil {
        ctx.ErrorJSON(400, "No file uploaded")
        return
    }
    defer file.Close()

    // Save file / 파일 저장
    dst, err := os.Create("./uploads/" + header.Filename)
    if err != nil {
        ctx.InternalServerError()
        return
    }
    defer dst.Close()

    if _, err := io.Copy(dst, file); err != nil {
        ctx.InternalServerError()
        return
    }

    ctx.SuccessJSON(200, "File uploaded", map[string]string{
        "filename": header.Filename,
        "size":     fmt.Sprintf("%d", header.Size),
    })
})
```

### Multiple File Upload / 다중 파일 업로드

```go
app.POST("/upload-multiple", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Parse multipart form / 멀티파트 폼 파싱
    if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
        ctx.ErrorJSON(400, "Failed to parse form")
        return
    }

    // Get all files / 모든 파일 가져오기
    files := r.MultipartForm.File["files"]

    var uploadedFiles []string
    for _, fileHeader := range files {
        file, err := fileHeader.Open()
        if err != nil {
            continue
        }
        defer file.Close()

        // Save each file / 각 파일 저장
        dst, err := os.Create("./uploads/" + fileHeader.Filename)
        if err != nil {
            continue
        }
        defer dst.Close()

        io.Copy(dst, file)
        uploadedFiles = append(uploadedFiles, fileHeader.Filename)
    }

    ctx.SuccessJSON(200, "Files uploaded", map[string]interface{}{
        "count": len(uploadedFiles),
        "files": uploadedFiles,
    })
})
```

### HTML Form Example / HTML 폼 예제

```html
<!-- Single file upload / 단일 파일 업로드 -->
<form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="file">
    <button type="submit">Upload</button>
</form>

<!-- Multiple file upload / 다중 파일 업로드 -->
<form action="/upload-multiple" method="post" enctype="multipart/form-data">
    <input type="file" name="files" multiple>
    <button type="submit">Upload Files</button>
</form>
```

---

## Error Handling / 에러 처리

### Standard Error Responses / 표준 에러 응답

```go
app.GET("/error-demo", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Custom error JSON / 커스텀 에러 JSON
    ctx.ErrorJSON(400, "Invalid request parameters")
    // Response: {"error":"Invalid request parameters","status":400,"success":false}
})
```

### HTTP Error Shortcuts / HTTP 에러 단축

```go
app.GET("/not-found", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.NotFound() // 404
})

app.GET("/unauthorized", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.Unauthorized() // 401
})

app.GET("/forbidden", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.Forbidden() // 403
})

app.GET("/bad-request", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.BadRequest() // 400
})

app.GET("/server-error", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.InternalServerError() // 500
})
```

### Abort Methods / 중단 메서드

```go
// Abort with status code / 상태 코드로 중단
app.GET("/abort-status", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.AbortWithStatus(http.StatusTeapot)
})

// Abort with error / 에러로 중단
app.GET("/abort-error", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    err := errors.New("something went wrong")
    ctx.AbortWithError(500, err)
})

// Abort with JSON / JSON으로 중단
app.GET("/abort-json", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.AbortWithJSON(403, map[string]string{
        "error": "Access denied",
    })
})
```

### Success Response / 성공 응답

```go
app.POST("/create", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Create resource / 리소스 생성
    resource := map[string]interface{}{
        "id":   123,
        "name": "New Resource",
    }

    ctx.SuccessJSON(201, "Resource created successfully", resource)
    // Response: {"message":"Resource created successfully","data":{"id":123,"name":"New Resource"},"status":201,"success":true}
})
```

---

## Graceful Shutdown / 우아한 종료

### Basic Graceful Shutdown / 기본 우아한 종료

```go
app := websvrutil.New()

app.GET("/", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.String(200, "Hello, World!")
})

// Run with graceful shutdown / 우아한 종료와 함께 실행
// Listens for SIGINT and SIGTERM / SIGINT와 SIGTERM 수신
// 5 second timeout for graceful shutdown / 우아한 종료를 위한 5초 타임아웃
if err := app.RunWithGracefulShutdown(":8080", 5*time.Second); err != nil {
    log.Fatal(err)
}
```

### Manual Shutdown Control / 수동 종료 제어

```go
app := websvrutil.New()

// Start server in goroutine / 고루틴에서 서버 시작
go func() {
    if err := app.Run(":8080"); err != nil && err != http.ErrServerClosed {
        log.Fatal(err)
    }
}()

// Wait for interrupt signal / 인터럽트 신호 대기
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit

log.Println("Shutting down server...")

// Shutdown with timeout / 타임아웃과 함께 종료
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := app.Shutdown(ctx); err != nil {
    log.Fatal("Server forced to shutdown:", err)
}

log.Println("Server exited")
```

### Cleanup on Shutdown / 종료 시 정리

```go
func main() {
    app := websvrutil.New()

    // Resources to cleanup / 정리할 리소스
    db := connectDatabase()
    defer db.Close()

    cache := connectCache()
    defer cache.Close()

    // Setup routes / 라우트 설정
    app.GET("/", homeHandler)

    // Graceful shutdown handles cleanup automatically
    // 우아한 종료가 자동으로 정리 처리
    if err := app.RunWithGracefulShutdown(":8080", 5*time.Second); err != nil {
        log.Fatal(err)
    }

    // Deferred cleanup will run after shutdown
    // 지연된 정리가 종료 후 실행됨
}
```

---

## Best Practices / 모범 사례

### 1. Use Route Groups for Organization / 조직화를 위한 라우트 그룹 사용

```go
// Good: Organized by feature / 좋음: 기능별로 조직화
api := app.Group("/api")
{
    users := api.Group("/users")
    users.GET("", listUsers)
    users.POST("", createUser)
    users.GET("/:id", getUser)

    posts := api.Group("/posts")
    posts.GET("", listPosts)
    posts.POST("", createPost)
}

// Bad: Flat structure / 나쁨: 평면 구조
app.GET("/api/users", listUsers)
app.POST("/api/users", createUser)
app.GET("/api/posts", listPosts)
```

### 2. Validate Input Early / 입력을 조기에 검증

```go
// Good: Validate early / 좋음: 조기 검증
app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    var user User
    if err := ctx.BindJSON(&user); err != nil {
        ctx.ErrorJSON(400, "Invalid JSON")
        return
    }

    if user.Email == "" {
        ctx.ErrorJSON(400, "Email is required")
        return
    }

    // Proceed with business logic / 비즈니스 로직 진행
    createUser(user)
    ctx.SuccessJSON(201, "User created", user)
})
```

### 3. Use Middleware for Cross-Cutting Concerns / 횡단 관심사에 미들웨어 사용

```go
// Good: Centralized logging, auth, etc. / 좋음: 중앙화된 로깅, 인증 등
app.Use(LoggerMiddleware)
app.Use(RecoveryMiddleware)
app.Use(CORSMiddleware)

admin := app.Group("/admin")
admin.Use(AuthMiddleware)

// Bad: Repeated logic in handlers / 나쁨: 핸들러에서 로직 반복
app.GET("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {
    // Check auth... / 인증 확인...
    // Log request... / 요청 로그...
    // Handle CORS... / CORS 처리...
    // Actual handler logic / 실제 핸들러 로직
})
```

### 4. Return Consistent Error Formats / 일관된 에러 형식 반환

```go
// Good: Consistent error format / 좋음: 일관된 에러 형식
ctx.ErrorJSON(400, "Invalid email format")
ctx.ErrorJSON(401, "Authentication required")
ctx.ErrorJSON(500, "Database connection failed")

// All return: {"error":"...","status":XXX,"success":false}

// Bad: Inconsistent formats / 나쁨: 비일관적 형식
ctx.JSON(400, "Invalid email")
ctx.JSON(401, map[string]string{"message": "Not authenticated"})
ctx.String(500, "Error occurred")
```

### 5. Use Type-Safe Request Binding / 타입 안전 요청 바인딩 사용

```go
// Good: Type-safe binding / 좋음: 타입 안전 바인딩
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=130"`
}

app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    var req CreateUserRequest
    if err := ctx.BindJSON(&req); err != nil {
        ctx.ErrorJSON(400, err.Error())
        return
    }

    // req is validated and type-safe / req는 검증되고 타입 안전함
})

// Bad: Manual parsing / 나쁨: 수동 파싱
app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
    var data map[string]interface{}
    json.NewDecoder(r.Body).Decode(&data)
    name := data["name"].(string) // Type assertion risks / 타입 단언 위험
})
```

### 6. Handle Panics with Recovery Middleware / 복구 미들웨어로 패닉 처리

```go
// Good: Global recovery / 좋음: 전역 복구
app.Use(RecoveryMiddleware)

app.GET("/panic", func(w http.ResponseWriter, r *http.Request) {
    panic("something went wrong") // Caught by middleware / 미들웨어가 잡음
})

// Bad: Unhandled panics / 나쁨: 처리되지 않은 패닉
app.GET("/panic", func(w http.ResponseWriter, r *http.Request) {
    panic("something went wrong") // Crashes server / 서버 충돌
})
```

### 7. Use Sessions Securely / 세션을 안전하게 사용

```go
// Good: Secure session configuration / 좋음: 안전한 세션 설정
store := websvrutil.NewSessionStore(websvrutil.SessionOptions{
    CookieName:  "session_id",
    MaxAge:      24 * time.Hour,
    Secure:      true,  // HTTPS only / HTTPS만
    HttpOnly:    true,  // No JavaScript access / JavaScript 접근 없음
    SameSite:    http.SameSiteStrictMode,
    CleanupTime: 5 * time.Minute,
})

// Bad: Insecure configuration / 나쁨: 안전하지 않은 설정
store := websvrutil.NewSessionStore(websvrutil.SessionOptions{
    Secure:   false, // Vulnerable to interception / 가로채기에 취약
    HttpOnly: false, // Vulnerable to XSS / XSS에 취약
})
```

### 8. Use Context Storage for Request-Scoped Data / 요청 범위 데이터에 컨텍스트 저장소 사용

```go
// Good: Use context storage / 좋음: 컨텍스트 저장소 사용
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        // Authenticate and store user / 인증 및 사용자 저장
        user := authenticateUser(r)
        ctx.Set("user", user)

        next.ServeHTTP(w, r)
    })
}

app.GET("/profile", AuthMiddleware, func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    user := ctx.MustGet("user").(User)
    ctx.JSON(200, user)
})
```

### 9. Use Graceful Shutdown in Production / 프로덕션에서 우아한 종료 사용

```go
// Good: Graceful shutdown / 좋음: 우아한 종료
app.RunWithGracefulShutdown(":8080", 10*time.Second)

// Bad: Abrupt shutdown / 나쁨: 갑작스러운 종료
app.Run(":8080")
```

### 10. Organize Handlers by Domain / 도메인별로 핸들러 조직화

```go
// Good: Organized by domain / 좋음: 도메인별로 조직화
// handlers/users.go
package handlers

func ListUsers(w http.ResponseWriter, r *http.Request) { ... }
func CreateUser(w http.ResponseWriter, r *http.Request) { ... }

// handlers/posts.go
package handlers

func ListPosts(w http.ResponseWriter, r *http.Request) { ... }
func CreatePost(w http.ResponseWriter, r *http.Request) { ... }

// main.go
app.GET("/users", handlers.ListUsers)
app.GET("/posts", handlers.ListPosts)

// Bad: Everything in main.go / 나쁨: 모든 것이 main.go에
func main() {
    app.GET("/users", func(...) { /* hundreds of lines */ })
    app.GET("/posts", func(...) { /* hundreds of lines */ })
}
```

---

## Troubleshooting / 문제 해결

### Issue: Routes Not Matching / 문제: 라우트가 매칭되지 않음

**Symptom**: 404 Not Found for valid routes / 증상: 유효한 라우트에 대해 404 Not Found

**Causes and Solutions**:

1. **Trailing slash mismatch** / **후행 슬래시 불일치**
   ```go
   // Route defined as / 라우트 정의
   app.GET("/users", handler)

   // Request to / 요청
   GET /users/  // 404! Missing route for /users/

   // Solution: Define both / 해결: 둘 다 정의
   app.GET("/users", handler)
   app.GET("/users/", handler)
   ```

2. **HTTP method mismatch** / **HTTP 메서드 불일치**
   ```go
   // Route defined as / 라우트 정의
   app.GET("/users", handler)

   // Request / 요청
   POST /users  // 404! No POST route defined

   // Solution: Check method / 해결: 메서드 확인
   app.POST("/users", handler)
   ```

3. **Parameter pattern issues** / **매개변수 패턴 문제**
   ```go
   // Route defined as / 라우트 정의
   app.GET("/users/:id", handler)

   // These work / 동작함
   GET /users/123
   GET /users/abc

   // This doesn't work / 동작하지 않음
   GET /users/  // Missing parameter
   ```

### Issue: Template Not Found / 문제: 템플릿을 찾을 수 없음

**Symptom**: "template not found" error / 증상: "template not found" 에러

**Causes and Solutions**:

1. **Templates not loaded** / **템플릿이 로드되지 않음**
   ```go
   // Bad: Forgot to load templates / 나쁨: 템플릿 로드를 잊음
   app.GET("/", func(w http.ResponseWriter, r *http.Request) {
       ctx.HTML(200, "index.html", nil) // Error!
   })

   // Good: Load templates first / 좋음: 먼저 템플릿 로드
   app.LoadHTMLGlob("templates/*")
   app.GET("/", func(w http.ResponseWriter, r *http.Request) {
       ctx.HTML(200, "index.html", nil)
   })
   ```

2. **Wrong file path** / **잘못된 파일 경로**
   ```go
   // Check your directory structure / 디렉토리 구조 확인
   project/
   ├── main.go
   └── templates/
       └── index.html

   // Use relative path from main.go / main.go에서 상대 경로 사용
   app.LoadHTMLGlob("templates/*")  // Correct / 올바름
   app.LoadHTMLGlob("./templates/*") // Also works / 역시 동작함
   ```

3. **File extension mismatch** / **파일 확장자 불일치**
   ```go
   // Template file: index.html
   ctx.HTML(200, "index.html", nil)  // Correct / 올바름
   ctx.HTML(200, "index", nil)       // Error! / 에러!
   ```

### Issue: Session Not Persisting / 문제: 세션이 유지되지 않음

**Symptom**: Session data lost between requests / 증상: 요청 간 세션 데이터 손실

**Causes and Solutions**:

1. **Forgot to save session** / **세션 저장을 잊음**
   ```go
   // Bad: Set but didn't save / 나쁨: 설정했지만 저장하지 않음
   session, _ := store.Get(r)
   session.Set("user_id", 123)
   // Session not saved! / 세션이 저장되지 않음!

   // Good: Save session / 좋음: 세션 저장
   session, _ := store.Get(r)
   session.Set("user_id", 123)
   store.Save(w, session)  // Must save! / 저장해야 함!
   ```

2. **Cookie settings issue** / **쿠키 설정 문제**
   ```go
   // Bad: Secure=true on HTTP / 나쁨: HTTP에서 Secure=true
   store := websvrutil.NewSessionStore(websvrutil.SessionOptions{
       Secure: true,  // Only works on HTTPS! / HTTPS에서만 동작!
   })

   // Good: Match your protocol / 좋음: 프로토콜에 맞춤
   store := websvrutil.NewSessionStore(websvrutil.SessionOptions{
       Secure: false,  // For HTTP development / HTTP 개발용
   })
   ```

3. **Session expired** / **세션 만료**
   ```go
   // Check MaxAge setting / MaxAge 설정 확인
   store := websvrutil.NewSessionStore(websvrutil.SessionOptions{
       MaxAge: 1 * time.Hour,  // Sessions expire after 1 hour / 1시간 후 세션 만료
   })
   ```

### Issue: JSON Binding Fails / 문제: JSON 바인딩 실패

**Symptom**: "invalid JSON" or binding errors / 증상: "invalid JSON" 또는 바인딩 에러

**Causes and Solutions**:

1. **Struct tag mismatch** / **구조체 태그 불일치**
   ```go
   type User struct {
       Name string `json:"name"`  // JSON field: "name"
   }

   // Request body / 요청 본문
   {"Name": "John"}  // Wrong! Should be lowercase "name" / 틀림! 소문자 "name"이어야 함
   {"name": "John"}  // Correct! / 올바름!
   ```

2. **Content-Type header missing** / **Content-Type 헤더 누락**
   ```go
   // Client must send / 클라이언트가 보내야 함
   Content-Type: application/json

   // Otherwise binding fails / 그렇지 않으면 바인딩 실패
   ```

3. **Validation failures** / **검증 실패**
   ```go
   type User struct {
       Email string `json:"email" binding:"required,email"`
       Age   int    `json:"age" binding:"gte=0,lte=130"`
   }

   // This fails validation / 검증 실패
   {"email": "invalid", "age": 200}

   // This passes / 통과
   {"email": "john@example.com", "age": 25}
   ```

### Issue: File Upload Fails / 문제: 파일 업로드 실패

**Symptom**: "no file uploaded" or "file too large" / 증상: "no file uploaded" 또는 "file too large"

**Causes and Solutions**:

1. **Form encoding mismatch** / **폼 인코딩 불일치**
   ```html
   <!-- Bad: Wrong enctype / 나쁨: 잘못된 enctype -->
   <form action="/upload" method="post">

   <!-- Good: Correct enctype / 좋음: 올바른 enctype -->
   <form action="/upload" method="post" enctype="multipart/form-data">
   ```

2. **Input name mismatch** / **입력 이름 불일치**
   ```go
   // Handler expects / 핸들러 예상
   file, _, err := ctx.FormFile("document")

   // HTML must match / HTML이 일치해야 함
   <input type="file" name="document">  <!-- Correct / 올바름 -->
   <input type="file" name="file">      <!-- Wrong! / 틀림! -->
   ```

3. **File size limit** / **파일 크기 제한**
   ```go
   // Increase size limit / 크기 제한 증가
   r.ParseMultipartForm(10 << 20)  // 10 MB
   r.ParseMultipartForm(50 << 20)  // 50 MB
   ```

### Issue: Middleware Not Running / 문제: 미들웨어가 실행되지 않음

**Symptom**: Middleware logic not executing / 증상: 미들웨어 로직이 실행되지 않음

**Causes and Solutions**:

1. **Order matters** / **순서가 중요함**
   ```go
   // Bad: Middleware added after routes / 나쁨: 라우트 후 미들웨어 추가
   app.GET("/users", handler)
   app.Use(LoggerMiddleware)  // Too late! / 너무 늦음!

   // Good: Middleware before routes / 좋음: 라우트 전 미들웨어
   app.Use(LoggerMiddleware)
   app.GET("/users", handler)
   ```

2. **Forgot to call next** / **next 호출을 잊음**
   ```go
   // Bad: Doesn't call next / 나쁨: next를 호출하지 않음
   func MyMiddleware(next http.Handler) http.Handler {
       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           // Do something / 무언가 수행
           // Forgot: next.ServeHTTP(w, r) / 잊음: next.ServeHTTP(w, r)
       })
   }

   // Good: Calls next / 좋음: next를 호출
   func MyMiddleware(next http.Handler) http.Handler {
       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           // Do something / 무언가 수행
           next.ServeHTTP(w, r)  // Continue chain / 체인 계속
       })
   }
   ```

---

## FAQ

### 1. How do I serve both API and web pages? / API와 웹 페이지를 둘 다 서빙하려면?

```go
app := websvrutil.New()

// Load templates for web pages / 웹 페이지용 템플릿 로드
app.LoadHTMLGlob("templates/*")

// Web pages / 웹 페이지
app.GET("/", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.HTML(200, "index.html", nil)
})

// API endpoints / API 엔드포인트
api := app.Group("/api")
api.GET("/users", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)
    ctx.JSON(200, users)
})
```

### 2. How do I handle CORS? / CORS를 어떻게 처리하나요?

```go
func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        ctx.SetHeaders(map[string]string{
            "Access-Control-Allow-Origin":  "*",
            "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
            "Access-Control-Allow-Headers": "Content-Type, Authorization",
        })

        if r.Method == "OPTIONS" {
            ctx.Status(204)
            return
        }

        next.ServeHTTP(w, r)
    })
}

app.Use(CORSMiddleware)
```

### 3. How do I implement authentication? / 인증을 어떻게 구현하나요?

```go
// Use sessions / 세션 사용
var sessionStore = websvrutil.NewSessionStore(websvrutil.DefaultSessionOptions())

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        session, err := sessionStore.Get(r)
        if err != nil || session.GetInt("user_id") == 0 {
            ctx.Unauthorized()
            return
        }

        next.ServeHTTP(w, r)
    })
}

// Protected routes / 보호된 라우트
app.GET("/profile", AuthMiddleware, profileHandler)
```

### 4. How do I validate request data? / 요청 데이터를 어떻게 검증하나요?

```go
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required,min=3,max=50"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"required,gte=0,lte=130"`
}

app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    var req CreateUserRequest
    if err := ctx.BindJSON(&req); err != nil {
        ctx.ErrorJSON(400, "Validation failed: " + err.Error())
        return
    }

    // Data is validated / 데이터가 검증됨
    createUser(req)
})
```

### 5. How do I serve a Single Page Application (SPA)? / SPA를 어떻게 서빙하나요?

```go
app := websvrutil.New()

// Serve static files / 정적 파일 서빙
app.Static("/static", "./dist/static")

// Catch-all route for SPA / SPA를 위한 모든 라우트 캐치
app.GET("/*", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./dist/index.html")
})
```

### 6. How do I limit request body size? / 요청 본문 크기를 어떻게 제한하나요?

```go
func LimitMiddleware(maxSize int64) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            r.Body = http.MaxBytesReader(w, r.Body, maxSize)
            next.ServeHTTP(w, r)
        })
    }
}

// Limit to 10MB / 10MB로 제한
app.Use(LimitMiddleware(10 << 20))
```

### 7. How do I implement rate limiting? / 속도 제한을 어떻게 구현하나요?

```go
import "golang.org/x/time/rate"

var limiter = rate.NewLimiter(10, 20) // 10 req/sec, burst 20

func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        if !limiter.Allow() {
            ctx.ErrorJSON(429, "Rate limit exceeded")
            return
        }

        next.ServeHTTP(w, r)
    })
}

app.Use(RateLimitMiddleware)
```

### 8. How do I handle file downloads? / 파일 다운로드를 어떻게 처리하나요?

```go
app.GET("/download", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Force download with custom filename / 커스텀 파일명으로 다운로드 강제
    ctx.FileAttachment("./files/report.pdf", "monthly-report-2025.pdf")
})
```

### 9. How do I log requests? / 요청을 어떻게 로깅하나요?

```go
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Create response recorder to capture status / 상태를 캡처하기 위한 응답 레코더 생성
        recorder := &ResponseRecorder{ResponseWriter: w, StatusCode: 200}

        next.ServeHTTP(recorder, r)

        duration := time.Since(start)
        log.Printf(
            "%s %s %d %v",
            r.Method,
            r.URL.Path,
            recorder.StatusCode,
            duration,
        )
    })
}

type ResponseRecorder struct {
    http.ResponseWriter
    StatusCode int
}

func (r *ResponseRecorder) WriteHeader(code int) {
    r.StatusCode = code
    r.ResponseWriter.WriteHeader(code)
}
```

### 10. How do I implement HTTPS? / HTTPS를 어떻게 구현하나요?

```go
// Generate self-signed certificate for development / 개발용 자체 서명 인증서 생성
// openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes

app := websvrutil.New()

// Setup routes / 라우트 설정
app.GET("/", homeHandler)

// Run with TLS / TLS로 실행
if err := app.RunTLS(":443", "cert.pem", "key.pem"); err != nil {
    log.Fatal(err)
}
```

### 11. How do I use custom templates? / 커스텀 템플릿을 어떻게 사용하나요?

```go
// Create custom template with functions / 함수가 있는 커스텀 템플릿 생성
tmpl := template.New("").Funcs(template.FuncMap{
    "upper": strings.ToUpper,
    "add": func(a, b int) int {
        return a + b
    },
})

// Parse templates / 템플릿 파싱
tmpl, err := tmpl.ParseGlob("templates/*")
if err != nil {
    log.Fatal(err)
}

// Set custom template / 커스텀 템플릿 설정
app.SetHTMLTemplate(tmpl)
```

### 12. How do I implement WebSocket? / WebSocket을 어떻게 구현하나요?

```go
import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all origins / 모든 출처 허용
    },
}

app.GET("/ws", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    if !ctx.IsWebSocket() {
        ctx.BadRequest()
        return
    }

    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        ctx.InternalServerError()
        return
    }
    defer conn.Close()

    // Handle WebSocket messages / WebSocket 메시지 처리
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            break
        }

        conn.WriteMessage(messageType, message)
    }
})
```

### 13. How do I test handlers? / 핸들러를 어떻게 테스트하나요?

```go
import (
    "net/http/httptest"
    "testing"
)

func TestHandler(t *testing.T) {
    app := websvrutil.New()

    app.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        ctx.String(200, "Hello, World!")
    })

    req := httptest.NewRequest("GET", "/hello", nil)
    w := httptest.NewRecorder()

    app.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf("Expected status 200, got %d", w.Code)
    }

    if w.Body.String() != "Hello, World!" {
        t.Errorf("Expected 'Hello, World!', got %s", w.Body.String())
    }
}
```

### 14. How do I use environment variables? / 환경 변수를 어떻게 사용하나요?

```go
import "os"

func main() {
    app := websvrutil.New()

    // Get port from environment / 환경에서 포트 가져오기
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    app.Run(":" + port)
}
```

### 15. How do I implement pagination? / 페이지네이션을 어떻게 구현하나요?

```go
app.GET("/users", func(w http.ResponseWriter, r *http.Request) {
    ctx := websvrutil.GetContext(r)

    // Get pagination parameters / 페이지네이션 매개변수 가져오기
    page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

    // Calculate offset / 오프셋 계산
    offset := (page - 1) * limit

    // Get paginated data / 페이지네이션된 데이터 가져오기
    users := getUsersPaginated(offset, limit)
    total := getTotalUsers()

    ctx.JSON(200, map[string]interface{}{
        "data":       users,
        "page":       page,
        "limit":      limit,
        "total":      total,
        "total_pages": (total + limit - 1) / limit,
    })
})
```

---

## Conclusion / 결론

The websvrutil package provides a comprehensive, developer-friendly toolkit for building web applications in Go. By following the patterns and best practices outlined in this manual, you can create robust, maintainable, and performant web services.

websvrutil 패키지는 Go에서 웹 애플리케이션을 구축하기 위한 포괄적이고 개발자 친화적인 툴킷을 제공합니다. 이 매뉴얼에 설명된 패턴과 모범 사례를 따르면 강력하고 유지 관리 가능하며 성능이 우수한 웹 서비스를 만들 수 있습니다.

For more information, examples, and updates, please visit:
더 많은 정보, 예제 및 업데이트는 다음을 방문하세요:

- GitHub Repository: https://github.com/arkd0ng/go-utils
- Package Documentation: https://pkg.go.dev/github.com/arkd0ng/go-utils/websvrutil

---

**Version**: v1.11.023
**Last Updated**: 2025-10-16
**License**: MIT
