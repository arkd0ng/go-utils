# httputil Package - User Manual
# httputil 패키지 - 사용자 매뉴얼

**Version / 버전**: v1.10.002
**Last Updated / 최종 업데이트**: 2025-10-15

---

## Table of Contents / 목차

1. [Introduction / 소개](#1-introduction--소개)
2. [Installation / 설치](#2-installation--설치)
3. [Quick Start / 빠른 시작](#3-quick-start--빠른-시작)
4. [Configuration Reference / 설정 참조](#4-configuration-reference--설정-참조)
5. [Usage Patterns / 사용 패턴](#5-usage-patterns--사용-패턴)
6. [Common Use Cases / 일반적인 사용 사례](#6-common-use-cases--일반적인-사용-사례)
7. [Error Handling / 에러 처리](#7-error-handling--에러-처리)
8. [Best Practices / 모범 사례](#8-best-practices--모범-사례)
9. [Troubleshooting / 문제 해결](#9-troubleshooting--문제-해결)
10. [FAQ](#10-faq)

---

## 1. Introduction / 소개

### 1.1 What is httputil? / httputil이란?

`httputil` is an extremely simple HTTP client library that reduces 30+ lines of boilerplate code to just 2-3 lines. It provides a clean, intuitive API for making HTTP requests with automatic retry logic, JSON encoding/decoding, and comprehensive error handling.

`httputil`은 30줄 이상의 보일러플레이트 코드를 단 2-3줄로 줄이는 극도로 간단한 HTTP 클라이언트 라이브러리입니다. 자동 재시도 로직, JSON 인코딩/디코딩, 종합적인 에러 처리를 갖춘 깔끔하고 직관적인 API를 제공합니다.

### 1.2 Key Features / 주요 기능

✅ **Extreme Simplicity** / **극도의 간결함**
- Reduces 30+ lines to 2-3 lines
- Zero boilerplate for common operations
- 30줄 이상을 2-3줄로 감소
- 일반적인 작업에 보일러플레이트 불필요

✅ **RESTful HTTP Methods** / **RESTful HTTP 메서드**
- GET, POST, PUT, PATCH, DELETE
- Context variants for all methods
- 모든 메서드의 Context 버전

✅ **Automatic JSON Handling** / **자동 JSON 처리**
- Request body encoding
- Response body decoding
- 요청 본문 인코딩
- 응답 본문 디코딩

✅ **Smart Retry Logic** / **스마트 재시도 로직**
- Configurable retry attempts
- Exponential backoff with jitter
- Automatic retry on failures
- 설정 가능한 재시도 횟수
- 지터가 있는 지수 백오프
- 실패 시 자동 재시도

✅ **Response Helpers** / **응답 헬퍼**
- Rich Response wrapper with 20+ helper methods
- Status code checks (IsSuccess, IsOK, IsNotFound, etc.)
- Body access (Body(), String(), JSON())
- 20개 이상의 헬퍼 메서드를 가진 풍부한 Response 래퍼
- 상태 코드 확인 (IsSuccess, IsOK, IsNotFound 등)
- 본문 접근 (Body(), String(), JSON())

✅ **File Operations** / **파일 작업**
- File download with progress tracking
- File upload with multipart form data
- Progress callbacks for large files
- 진행 상황 추적과 함께 파일 다운로드
- multipart form data를 사용한 파일 업로드
- 대용량 파일을 위한 진행 상황 콜백

✅ **URL & Form Builders** / **URL 및 Form 빌더**
- Fluent API for building URLs and forms
- Conditional parameters (ParamIf, AddIf)
- URL utilities (JoinURL, AddQueryParams, etc.)
- URL 및 폼 구축을 위한 Fluent API
- 조건부 매개변수 (ParamIf, AddIf)
- URL 유틸리티 (JoinURL, AddQueryParams 등)

✅ **Rich Error Types** / **풍부한 에러 타입**
- HTTPError (status code, body, URL)
- RetryError (failed attempts)
- TimeoutError (timeout detection)
- HTTPError (상태 코드, 본문, URL)
- RetryError (실패한 시도)
- TimeoutError (타임아웃 감지)

✅ **Zero External Dependencies** / **제로 외부 의존성**
- Standard library only
- 표준 라이브러리만 사용

### 1.3 Use Cases / 사용 사례

Perfect for:
- REST API clients
- Microservices communication
- External API integration
- Webhook handling
- Health check endpoints
- Data fetching services

다음과 같은 경우에 완벽:
- REST API 클라이언트
- 마이크로서비스 통신
- 외부 API 통합
- 웹훅 처리
- 헬스 체크 엔드포인트
- 데이터 가져오기 서비스

---

## 2. Installation / 설치

### 2.1 Prerequisites / 전제 조건

**Go Version / Go 버전:**
- Go 1.18 or later (for generics support)
- Go 1.18 이상 (제네릭 지원)

**Operating System / 운영 체제:**
- Linux, macOS, Windows
- Any platform supported by Go
- Go가 지원하는 모든 플랫폼

### 2.2 Install Package / 패키지 설치

```bash
# Install httputil package / httputil 패키지 설치
go get github.com/arkd0ng/go-utils/httputil
```

### 2.3 Import in Your Code / 코드에서 임포트

```go
import "github.com/arkd0ng/go-utils/httputil"
```

### 2.4 Verify Installation / 설치 확인

Create a simple test file:

간단한 테스트 파일 생성:

```go
package main

import (
    "fmt"
    "log"

    "github.com/arkd0ng/go-utils/httputil"
)

func main() {
    fmt.Println("httputil version:", httputil.Version)

    // Simple GET request / 간단한 GET 요청
    var result map[string]interface{}
    err := httputil.Get("https://jsonplaceholder.typicode.com/users/1", &result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Success: %+v\n", result)
}
```

Run the test:

테스트 실행:

```bash
go run main.go
```

---

## 3. Quick Start / 빠른 시작

### 3.1 Simple GET Request / 간단한 GET 요청

**Traditional Approach (30+ lines):**

전통적인 방식 (30줄 이상):

```go
// Traditional way / 전통적인 방식
client := &http.Client{
    Timeout: 30 * time.Second,
}

req, err := http.NewRequest("GET", "https://api.example.com/users", nil)
if err != nil {
    return err
}

req.Header.Set("Authorization", "Bearer token123")
req.Header.Set("Content-Type", "application/json")

resp, err := client.Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()

if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
}

var users []User
if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
    return err
}

// Plus retry logic, error handling... 20+ more lines
// 재시도 로직, 에러 처리 등... 20줄 이상 추가
```

**httputil Approach (2 lines):**

httputil 방식 (2줄):

```go
var users []User
err := httputil.Get("https://api.example.com/users", &users,
    httputil.WithBearerToken("token123"))
```

### 3.2 Simple POST Request / 간단한 POST 요청

```go
// Define payload / 페이로드 정의
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CreateUserResponse struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    CreatedAt string `json:"created_at"`
}

// Make POST request / POST 요청 실행
payload := CreateUserRequest{
    Name:  "John Doe",
    Email: "john@example.com",
}

var response CreateUserResponse
err := httputil.Post("https://api.example.com/users", payload, &response,
    httputil.WithBearerToken("your-token"))

if err != nil {
    log.Fatal(err)
}

fmt.Printf("Created user: %+v\n", response)
```

### 3.3 Client with Base URL / Base URL을 가진 클라이언트

When making multiple requests to the same API, create a client with a base URL:

동일한 API에 여러 요청을 할 때는 base URL을 가진 클라이언트를 생성하세요:

```go
// Create configured client / 설정된 클라이언트 생성
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com/v1"),
    httputil.WithBearerToken("your-token"),
    httputil.WithTimeout(60*time.Second),
    httputil.WithRetry(3),
)

// Use relative paths / 상대 경로 사용
var users []User
err := client.Get("/users", &users)

var user User
err = client.Get("/users/123", &user)

newUser := User{Name: "Jane"}
var created User
err = client.Post("/users", newUser, &created)
```

---

## 4. Configuration Reference / 설정 참조

### 4.1 Request Configuration / 요청 설정

#### WithTimeout

Set request timeout.

요청 타임아웃 설정.

```go
func WithTimeout(timeout time.Duration) Option
```

**Default / 기본값**: 30 seconds / 30초

**Example / 예제:**
```go
// Short timeout for health checks / 헬스 체크용 짧은 타임아웃
err := httputil.Get(url, &result,
    httputil.WithTimeout(5*time.Second))

// Long timeout for large uploads / 대용량 업로드용 긴 타임아웃
err := httputil.Post(url, data, &result,
    httputil.WithTimeout(5*time.Minute))
```

#### WithHeaders

Set multiple custom headers.

여러 사용자 정의 헤더 설정.

```go
func WithHeaders(headers map[string]string) Option
```

**Example / 예제:**
```go
err := httputil.Get(url, &result,
    httputil.WithHeaders(map[string]string{
        "X-API-Version": "v1",
        "X-Request-ID":  "abc123",
        "X-Client-Type": "mobile",
    }))
```

#### WithHeader

Set a single custom header.

단일 사용자 정의 헤더 설정.

```go
func WithHeader(key, value string) Option
```

**Example / 예제:**
```go
err := httputil.Get(url, &result,
    httputil.WithHeader("X-Custom-Header", "custom-value"))
```

#### WithQueryParams

Set query parameters.

쿼리 매개변수 설정.

```go
func WithQueryParams(params map[string]string) Option
```

**Example / 예제:**
```go
err := httputil.Get("https://api.example.com/search", &result,
    httputil.WithQueryParams(map[string]string{
        "q":     "golang",
        "page":  "1",
        "limit": "20",
        "sort":  "desc",
    }))

// Results in: https://api.example.com/search?q=golang&page=1&limit=20&sort=desc
```

#### WithUserAgent

Set custom User-Agent header.

사용자 정의 User-Agent 헤더 설정.

```go
func WithUserAgent(userAgent string) Option
```

**Default / 기본값**: "go-utils/httputil v{version}"

**Example / 예제:**
```go
err := httputil.Get(url, &result,
    httputil.WithUserAgent("MyApp/1.0"))
```

### 4.2 Authentication / 인증

#### WithBearerToken

Set Bearer token for authentication.

인증을 위한 Bearer 토큰 설정.

```go
func WithBearerToken(token string) Option
```

**Example / 예제:**
```go
err := httputil.Get(url, &result,
    httputil.WithBearerToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."))
```

**Sets Header / 헤더 설정:**
```
Authorization: Bearer {token}
```

#### WithBasicAuth

Set Basic Authentication credentials.

기본 인증 자격 증명 설정.

```go
func WithBasicAuth(username, password string) Option
```

**Example / 예제:**
```go
err := httputil.Get(url, &result,
    httputil.WithBasicAuth("admin", "secretpassword"))
```

### 4.3 Retry Configuration / 재시도 설정

#### WithRetry

Set maximum number of retry attempts.

최대 재시도 횟수 설정.

```go
func WithRetry(maxRetries int) Option
```

**Default / 기본값**: 3

**Example / 예제:**
```go
// No retries / 재시도 없음
err := httputil.Get(url, &result,
    httputil.WithRetry(0))

// Aggressive retries for critical operations / 중요한 작업용 적극적 재시도
err := httputil.Get(url, &result,
    httputil.WithRetry(10))
```

**Retry Behavior / 재시도 동작:**
- Retries on network errors / 네트워크 에러 시 재시도
- Retries on 5xx server errors / 5xx 서버 에러 시 재시도
- NO retry on 4xx client errors / 4xx 클라이언트 에러는 재시도 안 함
- Exponential backoff between retries / 재시도 간 지수 백오프

#### WithRetryBackoff

Set minimum and maximum backoff time for retries.

재시도의 최소 및 최대 백오프 시간 설정.

```go
func WithRetryBackoff(min, max time.Duration) Option
```

**Default / 기본값**: min=100ms, max=5s

**Example / 예제:**
```go
// Faster retries / 빠른 재시도
err := httputil.Get(url, &result,
    httputil.WithRetry(5),
    httputil.WithRetryBackoff(50*time.Millisecond, 1*time.Second))

// Slower retries / 느린 재시도
err := httputil.Get(url, &result,
    httputil.WithRetry(5),
    httputil.WithRetryBackoff(1*time.Second, 30*time.Second))
```

**Backoff Formula / 백오프 공식:**
```
backoff = min * 2^attempt (capped at max)
backoff = backoff ± (backoff * 25%)  // jitter
```

### 4.4 Client Configuration / 클라이언트 설정

#### WithBaseURL

Set base URL for all requests.

모든 요청의 기본 URL 설정.

```go
func WithBaseURL(baseURL string) Option
```

**Example / 예제:**
```go
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com/v1"))

// Now use relative paths / 이제 상대 경로 사용
client.Get("/users", &users)       // -> https://api.example.com/v1/users
client.Get("/users/123", &user)    // -> https://api.example.com/v1/users/123
```

#### WithFollowRedirects

Enable or disable following HTTP redirects.

HTTP 리디렉션 따르기 활성화 또는 비활성화.

```go
func WithFollowRedirects(follow bool) Option
```

**Default / 기본값**: `true`

**Example / 예제:**
```go
// Disable redirects / 리디렉션 비활성화
client := httputil.NewClient(
    httputil.WithFollowRedirects(false))
```

#### WithMaxRedirects

Set maximum number of redirects to follow.

따를 최대 리디렉션 수 설정.

```go
func WithMaxRedirects(max int) Option
```

**Default / 기본값**: 10

**Example / 예제:**
```go
client := httputil.NewClient(
    httputil.WithMaxRedirects(5))
```

---

## 5. Usage Patterns / 사용 패턴

### 5.1 Package-Level Functions / 패키지 레벨 함수

Use package-level functions for one-off requests.

일회성 요청에는 패키지 레벨 함수를 사용하세요.

```go
// Simple GET / 간단한 GET
var result MyStruct
err := httputil.Get("https://api.example.com/data", &result)

// GET with options / 옵션이 있는 GET
err := httputil.Get(url, &result,
    httputil.WithTimeout(10*time.Second),
    httputil.WithBearerToken("token"))

// POST with body / 본문이 있는 POST
payload := MyPayload{Name: "test"}
var response MyResponse
err := httputil.Post(url, payload, &response)
```

### 5.2 Client Instance / 클라이언트 인스턴스

Create a client for multiple requests to the same API.

동일한 API에 여러 요청을 위한 클라이언트 생성.

```go
// Create configured client / 설정된 클라이언트 생성
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithBearerToken("your-token"),
    httputil.WithTimeout(30*time.Second),
    httputil.WithRetry(3),
)

// Make multiple requests / 여러 요청 실행
var users []User
client.Get("/users", &users)

var posts []Post
client.Get("/posts", &posts)

newComment := Comment{Text: "Hello"}
var created Comment
client.Post("/comments", newComment, &created)
```

### 5.3 Context-Based Requests / Context 기반 요청

Use Context variants for cancellation and timeout control.

취소 및 타임아웃 제어를 위해 Context 변형을 사용하세요.

```go
// Create context with timeout / 타임아웃이 있는 context 생성
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Use context / context 사용
var result MyStruct
err := httputil.GetContext(ctx, url, &result)

// Or with client / 또는 클라이언트와 함께
err = client.GetContext(ctx, "/data", &result)
```

### 5.4 Per-Request Options / 요청별 옵션

Override client configuration for specific requests.

특정 요청에 대해 클라이언트 설정 재정의.

```go
// Client with default timeout / 기본 타임아웃이 있는 클라이언트
client := httputil.NewClient(
    httputil.WithTimeout(30*time.Second))

// Override timeout for specific request / 특정 요청의 타임아웃 재정의
var result MyStruct
err := client.Get("/slow-endpoint", &result,
    httputil.WithTimeout(2*time.Minute))  // Override
```

### 5.5 Combining Options / 옵션 결합

Combine multiple options for complex configurations.

복잡한 설정을 위해 여러 옵션 결합.

```go
err := httputil.Get(url, &result,
    httputil.WithTimeout(30*time.Second),
    httputil.WithBearerToken("token"),
    httputil.WithRetry(5),
    httputil.WithRetryBackoff(200*time.Millisecond, 10*time.Second),
    httputil.WithHeaders(map[string]string{
        "X-API-Version": "v2",
        "X-Request-ID":  uuid.New().String(),
    }),
    httputil.WithQueryParams(map[string]string{
        "filter": "active",
        "sort":   "desc",
    }),
)
```

### 5.6 Response Helpers / 응답 헬퍼

Use DoRaw to get Response with helper methods instead of auto-decoding JSON.

JSON 자동 디코딩 대신 헬퍼 메서드가 있는 Response를 가져오려면 DoRaw를 사용하세요.

```go
// Get raw response / 원시 응답 가져오기
resp, err := httputil.DoRaw("GET", "https://api.example.com/users", nil,
    httputil.WithBearerToken("token123"))
if err != nil {
    log.Fatal(err)
}

// Check status with helper methods / 헬퍼 메서드로 상태 확인
if resp.IsSuccess() {
    log.Println("Request successful")
}

if resp.IsNotFound() {
    log.Println("Resource not found")
}

// Access body in different formats / 다양한 형식으로 본문 접근
bodyBytes := resp.Body()        // []byte
bodyString := resp.String()     // string

// Decode JSON when ready / 준비되었을 때 JSON 디코딩
var users []User
err = resp.JSON(&users)

// Get headers / 헤더 가져오기
contentType := resp.ContentType()
customHeader := resp.Header("X-Custom-Header")
allHeaders := resp.Headers()
```

### 5.7 File Download / 파일 다운로드

Download files with optional progress tracking.

선택적 진행 상황 추적과 함께 파일을 다운로드합니다.

```go
// Simple file download / 간단한 파일 다운로드
err := httputil.DownloadFile(
    "https://example.com/file.zip",
    "./downloads/file.zip")
if err != nil {
    log.Fatal(err)
}

// Download with progress callback / 진행 상황 콜백과 함께 다운로드
ctx := context.Background()
err = httputil.DownloadFileContext(ctx,
    "https://example.com/large-file.zip",
    "./downloads/large-file.zip",
    func(bytesRead, totalBytes int64) {
        progress := float64(bytesRead) / float64(totalBytes) * 100
        fmt.Printf("\rDownloading: %.2f%%", progress)
    })

// Download to memory / 메모리로 다운로드
data, err := httputil.Download("https://example.com/data.json")
if err != nil {
    log.Fatal(err)
}
```

### 5.8 File Upload / 파일 업로드

Upload files using multipart form data.

multipart form data를 사용하여 파일을 업로드합니다.

```go
// Upload single file / 단일 파일 업로드
var result map[string]interface{}
err := httputil.UploadFile(
    "https://api.example.com/upload",
    "document",                    // Field name / 필드 이름
    "./files/report.pdf",          // File path / 파일 경로
    &result,
    httputil.WithBearerToken("token123"))
if err != nil {
    log.Fatal(err)
}
log.Printf("Upload result: %+v", result)

// Upload multiple files / 여러 파일 업로드
err = httputil.UploadFiles(
    "https://api.example.com/upload-multiple",
    map[string]string{
        "file1": "./images/image1.jpg",
        "file2": "./images/image2.jpg",
        "file3": "./documents/doc.pdf",
    },
    &result,
    httputil.WithBearerToken("token123"))
```

### 5.9 URL and Form Builders / URL 및 Form 빌더

Build URLs and forms using fluent API.

Fluent API를 사용하여 URL과 폼을 구축합니다.

```go
// URL Builder / URL 빌더
includeInactive := false
url := httputil.NewURL("https://api.example.com").
    Path("v1", "users", "search").
    Param("q", "golang").
    Param("page", "1").
    Param("limit", "20").
    ParamIf(includeInactive, "status", "inactive").
    Build()
// Result: https://api.example.com/v1/users/search?q=golang&page=1&limit=20

// URL utilities / URL 유틸리티
baseURL := "https://api.example.com"
fullURL := httputil.JoinURL(baseURL, "v1", "users", "123")
// Result: https://api.example.com/v1/users/123

domain, _ := httputil.GetDomain(fullURL)
// Result: api.example.com

// Form Builder / Form 빌더
hasPromoCode := true
form := httputil.NewForm().
    Set("username", "john").
    Set("email", "john@example.com").
    Set("age", "30").
    AddIf(hasPromoCode, "promo_code", "SAVE20").
    AddMultiple("tags", "go", "http", "api")

// Check if field exists / 필드 존재 확인
if form.Has("promo_code") {
    log.Println("Promo code applied")
}

// Post form data / 폼 데이터 전송
var result map[string]interface{}
err := httputil.PostForm(
    "https://api.example.com/submit",
    form.Map(),
    &result)
```

---

### 5.10 Cookie Management / 쿠키 관리

Automatic cookie persistence and session management with file storage support.

파일 저장소를 지원하는 자동 쿠키 지속성 및 세션 관리.

#### 5.10.1 In-Memory Cookies (Temporary) / 메모리 내 쿠키 (임시)

Cookies are stored in memory and discarded when the client is closed.

쿠키는 메모리에 저장되며 클라이언트가 닫힐 때 삭제됩니다.

```go
// Enable in-memory cookie jar / 메모리 내 쿠키 저장소 활성화
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithCookies(), // Enable cookie management / 쿠키 관리 활성화
)

// Make request - cookies are automatically handled
// 요청 수행 - 쿠키가 자동으로 처리됨
var result Response
err := client.Get("/login", &result)

// Server sets cookies in response (e.g., session_id)
// 서버가 응답에 쿠키 설정 (예: session_id)

// Subsequent requests automatically send cookies
// 후속 요청은 자동으로 쿠키 전송
err = client.Get("/profile", &result) // Sends session_id cookie / session_id 쿠키 전송
```

#### 5.10.2 Persistent Cookies (File Storage) / 지속성 쿠키 (파일 저장)

Cookies are saved to a file and automatically loaded on initialization.

쿠키가 파일에 저장되고 초기화 시 자동으로 로드됩니다.

```go
// Enable persistent cookie jar / 지속성 쿠키 저장소 활성화
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithPersistentCookies("cookies.json"), // Save to file / 파일에 저장
)

// Login and get session cookie / 로그인 및 세션 쿠키 획득
var loginResult LoginResponse
err := client.Post("/login", LoginRequest{
    Username: "user@example.com",
    Password: "password",
}, &loginResult)

// Save cookies to file / 쿠키를 파일에 저장
err = client.SaveCookies()
if err != nil {
    log.Printf("Failed to save cookies: %v", err)
}

// Next time, cookies are automatically loaded / 다음에는 쿠키가 자동으로 로드됨
// No need to login again / 다시 로그인할 필요 없음
client2 := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithPersistentCookies("cookies.json"), // Auto-loads cookies / 쿠키 자동 로드
)

// Authenticated request without login / 로그인 없이 인증된 요청
var profile UserProfile
err = client2.Get("/profile", &profile) // Uses saved session cookie / 저장된 세션 쿠키 사용
```

#### 5.10.3 Manual Cookie Operations / 수동 쿠키 작업

Direct cookie manipulation for advanced use cases.

고급 사용 사례를 위한 직접 쿠키 조작.

```go
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithCookies(),
)

u, _ := url.Parse("https://api.example.com")

// Set a custom cookie / 사용자 정의 쿠키 설정
client.SetCookie(u, &http.Cookie{
    Name:   "preference",
    Value:  "dark_mode",
    Path:   "/",
    MaxAge: 86400, // 24 hours / 24시간
})

// Check if cookie exists / 쿠키 존재 확인
if client.HasCookie(u, "session_id") {
    log.Println("User is logged in")
}

// Get specific cookie / 특정 쿠키 가져오기
cookie := client.GetCookie(u, "session_id")
if cookie != nil {
    log.Printf("Session ID: %s", cookie.Value)
}

// Get all cookies for domain / 도메인의 모든 쿠키 가져오기
cookies := client.GetCookies(u)
log.Printf("Total cookies: %d", len(cookies))

// Clear all cookies / 모든 쿠키 제거
err := client.ClearCookies()
if err != nil {
    log.Printf("Failed to clear cookies: %v", err)
}
```

#### 5.10.4 Cookie Persistence Workflow / 쿠키 지속성 워크플로우

Complete workflow for session persistence across application restarts.

애플리케이션 재시작 간 세션 지속성을 위한 완전한 워크플로우.

```go
func loginAndSaveSession(email, password, cookieFile string) error {
    // Create client with persistent cookies / 지속성 쿠키를 사용하는 클라이언트 생성
    client := httputil.NewClient(
        httputil.WithBaseURL("https://api.example.com"),
        httputil.WithPersistentCookies(cookieFile),
    )

    // Login / 로그인
    var result LoginResponse
    err := client.Post("/auth/login", LoginRequest{
        Email:    email,
        Password: password,
    }, &result)
    if err != nil {
        return fmt.Errorf("login failed: %w", err)
    }

    // Save session cookies to file / 세션 쿠키를 파일에 저장
    err = client.SaveCookies()
    if err != nil {
        return fmt.Errorf("failed to save cookies: %w", err)
    }

    log.Println("✓ Login successful, session saved")
    return nil
}

func makeAuthenticatedRequest(cookieFile string) error {
    // Create client and auto-load saved cookies / 클라이언트 생성 및 저장된 쿠키 자동 로드
    client := httputil.NewClient(
        httputil.WithBaseURL("https://api.example.com"),
        httputil.WithPersistentCookies(cookieFile), // Auto-loads / 자동 로드
    )

    // Make authenticated request / 인증된 요청 수행
    var profile UserProfile
    err := client.Get("/api/profile", &profile)
    if err != nil {
        return fmt.Errorf("request failed: %w", err)
    }

    log.Printf("✓ Profile loaded: %s", profile.Name)
    return nil
}

// Usage / 사용
func main() {
    cookieFile := "session.json"

    // First time: login and save / 첫 번째: 로그인 및 저장
    err := loginAndSaveSession("user@example.com", "password", cookieFile)
    if err != nil {
        log.Fatal(err)
    }

    // Later: use saved session / 나중에: 저장된 세션 사용
    err = makeAuthenticatedRequest(cookieFile)
    if err != nil {
        log.Fatal(err)
    }
}
```

#### 5.10.5 Shopping Cart Example / 쇼핑 카트 예제

Using cookies for shopping cart persistence.

쇼핑 카트 지속성을 위한 쿠키 사용.

```go
func addToCartWithCookies() {
    client := httputil.NewClient(
        httputil.WithBaseURL("https://shop.example.com"),
        httputil.WithPersistentCookies("cart_cookies.json"),
    )

    // Add items to cart / 카트에 항목 추가
    items := []CartItem{
        {ProductID: "prod-123", Quantity: 2},
        {ProductID: "prod-456", Quantity: 1},
    }

    for _, item := range items {
        var result AddToCartResponse
        err := client.Post("/api/cart/add", item, &result)
        if err != nil {
            log.Printf("Failed to add item: %v", err)
            continue
        }
        log.Printf("✓ Added: %s (qty: %d)", item.ProductID, item.Quantity)
    }

    // Save cart session / 카트 세션 저장
    client.SaveCookies()
    log.Println("✓ Cart session saved")

    // Later: view cart (auto-loads session) / 나중에: 카트 보기 (세션 자동 로드)
    client2 := httputil.NewClient(
        httputil.WithBaseURL("https://shop.example.com"),
        httputil.WithPersistentCookies("cart_cookies.json"),
    )

    var cart Cart
    err := client2.Get("/api/cart", &cart)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("✓ Cart loaded: %d items, total: $%.2f",
        len(cart.Items), cart.Total)
}
```

---

## 6. Common Use Cases / 일반적인 사용 사례

### 6.1 REST API Client / REST API 클라이언트

Complete example of a REST API client:

REST API 클라이언트의 완전한 예제:

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/arkd0ng/go-utils/httputil"
)

// User represents a user resource / User는 사용자 리소스를 나타냅니다
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

// APIClient wraps httputil client / APIClient는 httputil 클라이언트를 래핑합니다
type APIClient struct {
    client *httputil.Client
}

// NewAPIClient creates a new API client / NewAPIClient는 새 API 클라이언트를 생성합니다
func NewAPIClient(baseURL, token string) *APIClient {
    return &APIClient{
        client: httputil.NewClient(
            httputil.WithBaseURL(baseURL),
            httputil.WithBearerToken(token),
            httputil.WithTimeout(30*time.Second),
            httputil.WithRetry(3),
        ),
    }
}

// ListUsers gets all users / ListUsers는 모든 사용자를 가져옵니다
func (c *APIClient) ListUsers() ([]User, error) {
    var users []User
    err := c.client.Get("/users", &users)
    return users, err
}

// GetUser gets a single user / GetUser는 단일 사용자를 가져옵니다
func (c *APIClient) GetUser(id int) (*User, error) {
    var user User
    err := c.client.Get(fmt.Sprintf("/users/%d", id), &user)
    return &user, err
}

// CreateUser creates a new user / CreateUser는 새 사용자를 생성합니다
func (c *APIClient) CreateUser(name, email string) (*User, error) {
    payload := map[string]string{
        "name":  name,
        "email": email,
    }

    var user User
    err := c.client.Post("/users", payload, &user)
    return &user, err
}

// UpdateUser updates an existing user / UpdateUser는 기존 사용자를 업데이트합니다
func (c *APIClient) UpdateUser(id int, name, email string) (*User, error) {
    payload := map[string]string{
        "name":  name,
        "email": email,
    }

    var user User
    err := c.client.Put(fmt.Sprintf("/users/%d", id), payload, &user)
    return &user, err
}

// DeleteUser deletes a user / DeleteUser는 사용자를 삭제합니다
func (c *APIClient) DeleteUser(id int) error {
    return c.client.Delete(fmt.Sprintf("/users/%d", id), nil)
}

func main() {
    // Create API client / API 클라이언트 생성
    api := NewAPIClient("https://api.example.com/v1", "your-token")

    // List users / 사용자 목록
    users, err := api.ListUsers()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d users\n", len(users))

    // Get specific user / 특정 사용자 가져오기
    user, err := api.GetUser(1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User: %+v\n", user)

    // Create user / 사용자 생성
    newUser, err := api.CreateUser("John Doe", "john@example.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created: %+v\n", newUser)

    // Update user / 사용자 업데이트
    updated, err := api.UpdateUser(newUser.ID, "Jane Doe", "jane@example.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Updated: %+v\n", updated)

    // Delete user / 사용자 삭제
    if err := api.DeleteUser(newUser.ID); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Deleted successfully")
}
```

### 6.2 Webhook Handler / 웹훅 핸들러

Handle incoming webhooks and make outbound requests:

수신 웹훅을 처리하고 아웃바운드 요청 실행:

```go
package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/arkd0ng/go-utils/httputil"
)

// WebhookPayload represents incoming webhook data
// WebhookPayload는 수신 웹훅 데이터를 나타냅니다
type WebhookPayload struct {
    Event string                 `json:"event"`
    Data  map[string]interface{} `json:"data"`
}

// NotificationService sends notifications
// NotificationService는 알림을 전송합니다
type NotificationService struct {
    client *httputil.Client
}

func NewNotificationService(apiURL, apiKey string) *NotificationService {
    return &NotificationService{
        client: httputil.NewClient(
            httputil.WithBaseURL(apiURL),
            httputil.WithBearerToken(apiKey),
            httputil.WithTimeout(10*time.Second),
            httputil.WithRetry(3),
        ),
    }
}

// SendNotification sends a notification / SendNotification은 알림을 전송합니다
func (s *NotificationService) SendNotification(ctx context.Context, message string) error {
    payload := map[string]string{
        "message": message,
        "channel": "alerts",
    }

    var result map[string]interface{}
    return s.client.PostContext(ctx, "/notifications", payload, &result)
}

// HandleWebhook processes incoming webhooks / HandleWebhook은 수신 웹훅을 처리합니다
func HandleWebhook(notifier *NotificationService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Parse webhook payload / 웹훅 페이로드 파싱
        var payload WebhookPayload
        if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
            http.Error(w, "Invalid payload", http.StatusBadRequest)
            return
        }

        // Create context with timeout / 타임아웃이 있는 context 생성
        ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
        defer cancel()

        // Send notification based on event / 이벤트에 따라 알림 전송
        message := "Event received: " + payload.Event
        if err := notifier.SendNotification(ctx, message); err != nil {
            log.Printf("Failed to send notification: %v", err)
            http.Error(w, "Internal error", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{
            "status": "success",
        })
    }
}

func main() {
    // Create notification service / 알림 서비스 생성
    notifier := NewNotificationService(
        "https://notifications.example.com",
        "your-api-key",
    )

    // Setup webhook handler / 웹훅 핸들러 설정
    http.HandleFunc("/webhook", HandleWebhook(notifier))

    // Start server / 서버 시작
    log.Println("Webhook server listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 6.3 Health Check Service / 헬스 체크 서비스

Monitor multiple services with health checks:

헬스 체크로 여러 서비스 모니터링:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "sync"
    "time"

    "github.com/arkd0ng/go-utils/httputil"
)

// ServiceStatus represents service health status
// ServiceStatus는 서비스 헬스 상태를 나타냅니다
type ServiceStatus struct {
    Name      string
    URL       string
    Healthy   bool
    Latency   time.Duration
    Error     error
    CheckedAt time.Time
}

// HealthChecker monitors service health / HealthChecker는 서비스 헬스를 모니터링합니다
type HealthChecker struct {
    services []string
    client   *httputil.Client
}

func NewHealthChecker(services []string) *HealthChecker {
    return &HealthChecker{
        services: services,
        client: httputil.NewClient(
            httputil.WithTimeout(5*time.Second),
            httputil.WithRetry(1),
        ),
    }
}

// CheckService checks a single service / CheckService는 단일 서비스를 확인합니다
func (h *HealthChecker) CheckService(ctx context.Context, url string) ServiceStatus {
    start := time.Now()
    status := ServiceStatus{
        Name:      url,
        URL:       url,
        CheckedAt: time.Now(),
    }

    var result map[string]interface{}
    err := h.client.GetContext(ctx, url, &result)

    status.Latency = time.Since(start)
    status.Healthy = err == nil
    status.Error = err

    return status
}

// CheckAll checks all services concurrently / CheckAll은 모든 서비스를 동시에 확인합니다
func (h *HealthChecker) CheckAll(ctx context.Context) []ServiceStatus {
    var wg sync.WaitGroup
    results := make([]ServiceStatus, len(h.services))

    for i, service := range h.services {
        wg.Add(1)
        go func(index int, url string) {
            defer wg.Done()
            results[index] = h.CheckService(ctx, url)
        }(i, service)
    }

    wg.Wait()
    return results
}

// PrintStatus prints health check results / PrintStatus는 헬스 체크 결과를 출력합니다
func PrintStatus(statuses []ServiceStatus) {
    fmt.Println("\n=== Service Health Check ===")
    for _, status := range statuses {
        healthStatus := "✅ HEALTHY"
        if !status.Healthy {
            healthStatus = "❌ UNHEALTHY"
        }

        fmt.Printf("%s %s (latency: %v)\n",
            healthStatus, status.Name, status.Latency)

        if status.Error != nil {
            fmt.Printf("  Error: %v\n", status.Error)
        }
    }
}

func main() {
    // Define services to monitor / 모니터링할 서비스 정의
    services := []string{
        "https://api.service1.com/health",
        "https://api.service2.com/health",
        "https://api.service3.com/health",
    }

    checker := NewHealthChecker(services)

    // Check services every 30 seconds / 30초마다 서비스 확인
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()

    for {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        statuses := checker.CheckAll(ctx)
        cancel()

        PrintStatus(statuses)

        <-ticker.C
    }
}
```

### 6.4 Data Fetching Service / 데이터 가져오기 서비스

Fetch and aggregate data from multiple sources:

여러 소스에서 데이터 가져오기 및 집계:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "sync"
    "time"

    "github.com/arkd0ng/go-utils/httputil"
)

// User data from API / API의 사용자 데이터
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Post data from API / API의 게시물 데이터
type Post struct {
    ID     int    `json:"id"`
    UserID int    `json:"userId"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

// UserProfile combines user and posts / UserProfile은 사용자와 게시물을 결합합니다
type UserProfile struct {
    User  User
    Posts []Post
}

// DataFetcher fetches data from multiple sources
// DataFetcher는 여러 소스에서 데이터를 가져옵니다
type DataFetcher struct {
    client *httputil.Client
}

func NewDataFetcher(baseURL string) *DataFetcher {
    return &DataFetcher{
        client: httputil.NewClient(
            httputil.WithBaseURL(baseURL),
            httputil.WithTimeout(30*time.Second),
            httputil.WithRetry(3),
        ),
    }
}

// FetchUser fetches user data / FetchUser는 사용자 데이터를 가져옵니다
func (f *DataFetcher) FetchUser(ctx context.Context, userID int) (*User, error) {
    var user User
    err := f.client.GetContext(ctx, fmt.Sprintf("/users/%d", userID), &user)
    return &user, err
}

// FetchUserPosts fetches user's posts / FetchUserPosts는 사용자의 게시물을 가져옵니다
func (f *DataFetcher) FetchUserPosts(ctx context.Context, userID int) ([]Post, error) {
    var posts []Post
    err := f.client.GetContext(ctx, "/posts", &posts,
        httputil.WithQueryParams(map[string]string{
            "userId": fmt.Sprintf("%d", userID),
        }))
    return posts, err
}

// FetchUserProfile fetches complete user profile concurrently
// FetchUserProfile은 전체 사용자 프로필을 동시에 가져옵니다
func (f *DataFetcher) FetchUserProfile(ctx context.Context, userID int) (*UserProfile, error) {
    var wg sync.WaitGroup
    profile := &UserProfile{}

    var userErr, postsErr error

    // Fetch user and posts concurrently / 사용자와 게시물을 동시에 가져오기
    wg.Add(2)

    go func() {
        defer wg.Done()
        user, err := f.FetchUser(ctx, userID)
        if err != nil {
            userErr = err
            return
        }
        profile.User = *user
    }()

    go func() {
        defer wg.Done()
        posts, err := f.FetchUserPosts(ctx, userID)
        if err != nil {
            postsErr = err
            return
        }
        profile.Posts = posts
    }()

    wg.Wait()

    // Check for errors / 에러 확인
    if userErr != nil {
        return nil, fmt.Errorf("failed to fetch user: %w", userErr)
    }
    if postsErr != nil {
        return nil, fmt.Errorf("failed to fetch posts: %w", postsErr)
    }

    return profile, nil
}

func main() {
    // Create data fetcher / 데이터 가져오기 생성
    fetcher := NewDataFetcher("https://jsonplaceholder.typicode.com")

    // Fetch user profile / 사용자 프로필 가져오기
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    profile, err := fetcher.FetchUserProfile(ctx, 1)
    if err != nil {
        log.Fatal(err)
    }

    // Print results / 결과 출력
    fmt.Printf("User: %s (%s)\n", profile.User.Name, profile.User.Email)
    fmt.Printf("Posts: %d\n", len(profile.Posts))
    for _, post := range profile.Posts {
        fmt.Printf("  - %s\n", post.Title)
    }
}
```

### 6.5 Rate-Limited API Client / 속도 제한 API 클라이언트

Client that respects API rate limits:

API 속도 제한을 준수하는 클라이언트:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/arkd0ng/go-utils/httputil"
    "golang.org/x/time/rate"
)

// RateLimitedClient wraps httputil with rate limiting
// RateLimitedClient는 속도 제한이 있는 httputil을 래핑합니다
type RateLimitedClient struct {
    client  *httputil.Client
    limiter *rate.Limiter
}

// NewRateLimitedClient creates a rate-limited client
// NewRateLimitedClient는 속도 제한 클라이언트를 생성합니다
func NewRateLimitedClient(baseURL, token string, requestsPerSecond int) *RateLimitedClient {
    return &RateLimitedClient{
        client: httputil.NewClient(
            httputil.WithBaseURL(baseURL),
            httputil.WithBearerToken(token),
            httputil.WithTimeout(30*time.Second),
            httputil.WithRetry(3),
        ),
        limiter: rate.NewLimiter(rate.Limit(requestsPerSecond), 1),
    }
}

// Get makes a rate-limited GET request / Get은 속도 제한 GET 요청을 실행합니다
func (c *RateLimitedClient) Get(ctx context.Context, path string, result interface{}, opts ...httputil.Option) error {
    // Wait for rate limiter / 속도 제한 대기
    if err := c.limiter.Wait(ctx); err != nil {
        return fmt.Errorf("rate limit wait failed: %w", err)
    }

    return c.client.GetContext(ctx, path, result, opts...)
}

// Post makes a rate-limited POST request / Post는 속도 제한 POST 요청을 실행합니다
func (c *RateLimitedClient) Post(ctx context.Context, path string, body, result interface{}, opts ...httputil.Option) error {
    // Wait for rate limiter / 속도 제한 대기
    if err := c.limiter.Wait(ctx); err != nil {
        return fmt.Errorf("rate limit wait failed: %w", err)
    }

    return c.client.PostContext(ctx, path, body, result, opts...)
}

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    // Create rate-limited client (10 requests per second)
    // 속도 제한 클라이언트 생성 (초당 10개 요청)
    client := NewRateLimitedClient(
        "https://api.example.com",
        "your-token",
        10, // 10 requests per second
    )

    ctx := context.Background()

    // Make multiple requests (automatically rate-limited)
    // 여러 요청 실행 (자동으로 속도 제한됨)
    for i := 1; i <= 50; i++ {
        var user User
        if err := client.Get(ctx, fmt.Sprintf("/users/%d", i), &user); err != nil {
            log.Printf("Failed to fetch user %d: %v", i, err)
            continue
        }
        fmt.Printf("Fetched user %d: %s\n", i, user.Name)
    }
}
```

### 6.6 File Download Service with Progress / 진행 상황과 함께 파일 다운로드 서비스

Complete example of a file download service with progress tracking:

진행 상황 추적과 함께 파일 다운로드 서비스의 완전한 예제:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/arkd0ng/go-utils/httputil"
)

// FileDownloader handles file downloads with progress tracking
// FileDownloader는 진행 상황 추적과 함께 파일 다운로드를 처리합니다
type FileDownloader struct {
    client *httputil.Client
}

// NewFileDownloader creates a new file downloader
// NewFileDownloader는 새 파일 다운로더를 생성합니다
func NewFileDownloader(timeout time.Duration) *FileDownloader {
    return &FileDownloader{
        client: httputil.NewClient(
            httputil.WithTimeout(timeout),
            httputil.WithRetry(3),
        ),
    }
}

// DownloadWithProgress downloads a file with progress tracking
// DownloadWithProgress는 진행 상황 추적과 함께 파일을 다운로드합니다
func (fd *FileDownloader) DownloadWithProgress(ctx context.Context, url, filePath string) error {
    startTime := time.Now()

    err := fd.client.DownloadFileContext(ctx, url, filePath,
        func(bytesRead, totalBytes int64) {
            if totalBytes > 0 {
                progress := float64(bytesRead) / float64(totalBytes) * 100
                speed := float64(bytesRead) / time.Since(startTime).Seconds() / 1024 / 1024 // MB/s
                fmt.Printf("\rProgress: %.2f%% (%.2f MB/s)", progress, speed)
            }
        })

    if err != nil {
        return fmt.Errorf("download failed: %w", err)
    }

    fmt.Printf("\nDownload completed in %v\n", time.Since(startTime))
    return nil
}

// DownloadMultiple downloads multiple files concurrently
// DownloadMultiple은 여러 파일을 동시에 다운로드합니다
func (fd *FileDownloader) DownloadMultiple(ctx context.Context, files map[string]string) error {
    errCh := make(chan error, len(files))

    for url, filePath := range files {
        go func(u, fp string) {
            fmt.Printf("Starting download: %s\n", fp)
            errCh <- fd.DownloadWithProgress(ctx, u, fp)
        }(url, filePath)
    }

    // Wait for all downloads / 모든 다운로드 대기
    for i := 0; i < len(files); i++ {
        if err := <-errCh; err != nil {
            return err
        }
    }

    fmt.Println("All downloads completed")
    return nil
}

func main() {
    downloader := NewFileDownloader(5 * time.Minute)
    ctx := context.Background()

    // Single file download / 단일 파일 다운로드
    err := downloader.DownloadWithProgress(ctx,
        "https://example.com/large-file.zip",
        "./downloads/file.zip")
    if err != nil {
        log.Fatal(err)
    }

    // Multiple files download / 여러 파일 다운로드
    files := map[string]string{
        "https://example.com/file1.zip": "./downloads/file1.zip",
        "https://example.com/file2.zip": "./downloads/file2.zip",
        "https://example.com/file3.zip": "./downloads/file3.zip",
    }

    err = downloader.DownloadMultiple(ctx, files)
    if err != nil {
        log.Fatal(err)
    }
}
```

### 6.7 File Upload Service / 파일 업로드 서비스

Complete example of a file upload service:

파일 업로드 서비스의 완전한 예제:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "path/filepath"
    "time"

    "github.com/arkd0ng/go-utils/httputil"
)

// UploadResult represents the result of a file upload
// UploadResult는 파일 업로드의 결과를 나타냅니다
type UploadResult struct {
    FileID   string `json:"file_id"`
    FileName string `json:"file_name"`
    FileSize int64  `json:"file_size"`
    URL      string `json:"url"`
}

// FileUploader handles file uploads
// FileUploader는 파일 업로드를 처리합니다
type FileUploader struct {
    client *httputil.Client
}

// NewFileUploader creates a new file uploader
// NewFileUploader는 새 파일 업로더를 생성합니다
func NewFileUploader(apiURL, token string) *FileUploader {
    return &FileUploader{
        client: httputil.NewClient(
            httputil.WithBaseURL(apiURL),
            httputil.WithBearerToken(token),
            httputil.WithTimeout(10*time.Minute),
            httputil.WithRetry(3),
        ),
    }
}

// UploadSingle uploads a single file
// UploadSingle은 단일 파일을 업로드합니다
func (fu *FileUploader) UploadSingle(ctx context.Context, filePath string) (*UploadResult, error) {
    var result UploadResult

    err := fu.client.UploadFileContext(ctx,
        "/upload",
        "file",
        filePath,
        &result)

    if err != nil {
        return nil, fmt.Errorf("upload failed: %w", err)
    }

    return &result, nil
}

// UploadBatch uploads multiple files in a single request
// UploadBatch는 단일 요청으로 여러 파일을 업로드합니다
func (fu *FileUploader) UploadBatch(ctx context.Context, filePaths []string) ([]UploadResult, error) {
    // Build file map / 파일 맵 구축
    files := make(map[string]string)
    for i, fp := range filePaths {
        fieldName := fmt.Sprintf("file%d", i+1)
        files[fieldName] = fp
    }

    var results []UploadResult
    err := fu.client.UploadFilesContext(ctx, "/upload-batch", files, &results)
    if err != nil {
        return nil, fmt.Errorf("batch upload failed: %w", err)
    }

    return results, nil
}

// UploadDirectory uploads all files in a directory
// UploadDirectory는 디렉토리의 모든 파일을 업로드합니다
func (fu *FileUploader) UploadDirectory(ctx context.Context, dirPath string, pattern string) error {
    matches, err := filepath.Glob(filepath.Join(dirPath, pattern))
    if err != nil {
        return fmt.Errorf("failed to list files: %w", err)
    }

    fmt.Printf("Found %d files to upload\n", len(matches))

    for i, filePath := range matches {
        fmt.Printf("Uploading %d/%d: %s\n", i+1, len(matches), filePath)

        result, err := fu.UploadSingle(ctx, filePath)
        if err != nil {
            return fmt.Errorf("failed to upload %s: %w", filePath, err)
        }

        fmt.Printf("  Uploaded: %s (ID: %s)\n", result.FileName, result.FileID)
    }

    fmt.Println("All files uploaded successfully")
    return nil
}

func main() {
    uploader := NewFileUploader("https://api.example.com", "your-token")
    ctx := context.Background()

    // Upload single file / 단일 파일 업로드
    result, err := uploader.UploadSingle(ctx, "./documents/report.pdf")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Uploaded: %s (URL: %s)\n", result.FileName, result.URL)

    // Upload batch / 배치 업로드
    results, err := uploader.UploadBatch(ctx, []string{
        "./images/image1.jpg",
        "./images/image2.jpg",
        "./images/image3.jpg",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Uploaded %d files\n", len(results))

    // Upload directory / 디렉토리 업로드
    err = uploader.UploadDirectory(ctx, "./documents", "*.pdf")
    if err != nil {
        log.Fatal(err)
    }
}
```

---

## 7. Error Handling / 에러 처리

### 7.1 Error Types / 에러 타입

httputil provides three rich error types:

httputil은 세 가지 풍부한 에러 타입을 제공합니다:

#### HTTPError

HTTP error with status code and response body.

상태 코드와 응답 본문을 포함한 HTTP 에러.

```go
type HTTPError struct {
    StatusCode int    // HTTP status code / HTTP 상태 코드
    Status     string // Status text / 상태 텍스트
    Body       string // Response body / 응답 본문
    URL        string // Request URL / 요청 URL
    Method     string // HTTP method / HTTP 메서드
}
```

**Check and Handle / 확인 및 처리:**
```go
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsHTTPError(err) {
        statusCode := httputil.GetStatusCode(err)

        switch {
        case statusCode == 404:
            // Not found / 찾을 수 없음
            fmt.Println("Resource not found")
        case statusCode >= 400 && statusCode < 500:
            // Client error / 클라이언트 에러
            fmt.Println("Bad request")
        case statusCode >= 500:
            // Server error / 서버 에러
            fmt.Println("Server error")
        }
    }
}
```

#### RetryError

Error after all retry attempts failed.

모든 재시도 시도 실패 후 에러.

```go
type RetryError struct {
    Attempts int    // Number of attempts / 시도 횟수
    LastErr  error  // Last error / 마지막 에러
    URL      string // Request URL / 요청 URL
    Method   string // HTTP method / HTTP 메서드
}
```

**Check and Handle / 확인 및 처리:**
```go
err := httputil.Get(url, &result, httputil.WithRetry(5))
if err != nil {
    if httputil.IsRetryError(err) {
        retryErr := err.(*httputil.RetryError)
        fmt.Printf("Failed after %d attempts: %v\n",
            retryErr.Attempts, retryErr.LastErr)
    }
}
```

#### TimeoutError

Request timeout error.

요청 타임아웃 에러.

```go
type TimeoutError struct {
    URL    string // Request URL / 요청 URL
    Method string // HTTP method / HTTP 메서드
}
```

**Check and Handle / 확인 및 처리:**
```go
err := httputil.Get(url, &result, httputil.WithTimeout(5*time.Second))
if err != nil {
    if httputil.IsTimeoutError(err) {
        fmt.Println("Request timed out")
        // Retry with longer timeout / 더 긴 타임아웃으로 재시도
        err = httputil.Get(url, &result, httputil.WithTimeout(30*time.Second))
    }
}
```

### 7.2 Error Handling Patterns / 에러 처리 패턴

#### Pattern 1: Type-Based Handling / 타입 기반 처리

```go
err := httputil.Get(url, &result)
if err != nil {
    switch {
    case httputil.IsHTTPError(err):
        handleHTTPError(err)
    case httputil.IsRetryError(err):
        handleRetryError(err)
    case httputil.IsTimeoutError(err):
        handleTimeoutError(err)
    default:
        handleUnknownError(err)
    }
}
```

#### Pattern 2: Status Code Handling / 상태 코드 처리

```go
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsHTTPError(err) {
        statusCode := httputil.GetStatusCode(err)

        if statusCode == 401 {
            // Refresh token and retry / 토큰 새로고침 및 재시도
            newToken := refreshToken()
            err = httputil.Get(url, &result,
                httputil.WithBearerToken(newToken))
        } else if statusCode == 429 {
            // Rate limited, wait and retry / 속도 제한, 대기 후 재시도
            time.Sleep(time.Minute)
            err = httputil.Get(url, &result)
        }
    }
}
```

#### Pattern 3: Wrapped Error Handling / 래핑된 에러 처리

```go
func fetchData(url string) (*Data, error) {
    var data Data
    err := httputil.Get(url, &data)
    if err != nil {
        if httputil.IsHTTPError(err) {
            return nil, fmt.Errorf("HTTP error fetching data: %w", err)
        }
        return nil, fmt.Errorf("failed to fetch data: %w", err)
    }
    return &data, nil
}
```

#### Pattern 4: Retry with Exponential Backoff / 지수 백오프로 재시도

```go
func fetchWithRetry(url string, maxAttempts int) (*Data, error) {
    var data Data

    for attempt := 0; attempt < maxAttempts; attempt++ {
        err := httputil.Get(url, &data,
            httputil.WithRetry(0), // Disable internal retry
            httputil.WithTimeout(time.Duration(10*(attempt+1))*time.Second))

        if err == nil {
            return &data, nil
        }

        // Don't retry on 4xx errors / 4xx 에러는 재시도 안 함
        if httputil.IsHTTPError(err) {
            statusCode := httputil.GetStatusCode(err)
            if statusCode >= 400 && statusCode < 500 {
                return nil, err
            }
        }

        // Wait before retry / 재시도 전 대기
        if attempt < maxAttempts-1 {
            backoff := time.Duration(math.Pow(2, float64(attempt))) * time.Second
            time.Sleep(backoff)
        }
    }

    return nil, fmt.Errorf("failed after %d attempts", maxAttempts)
}
```

---

## 8. Best Practices / 모범 사례

### 8.1 Client Reuse / 클라이언트 재사용

✅ **DO: Reuse client for multiple requests**

✅ **좋음: 여러 요청에 클라이언트 재사용**

```go
// Good: Create once, use many times / 좋음: 한 번 생성, 여러 번 사용
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithBearerToken("token"),
)

for i := 0; i < 100; i++ {
    client.Get(fmt.Sprintf("/users/%d", i), &user)
}
```

❌ **DON'T: Create new client for each request**

❌ **나쁨: 각 요청마다 새 클라이언트 생성**

```go
// Bad: Creates new client each time / 나쁨: 매번 새 클라이언트 생성
for i := 0; i < 100; i++ {
    httputil.Get(fmt.Sprintf("https://api.example.com/users/%d", i), &user,
        httputil.WithBearerToken("token"))
}
```

### 8.2 Context Usage / Context 사용

✅ **DO: Always use Context for long-running requests**

✅ **좋음: 장시간 실행 요청에 항상 Context 사용**

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

err := httputil.GetContext(ctx, url, &result)
```

✅ **DO: Propagate context from upstream**

✅ **좋음: 상위에서 context 전파**

```go
func HandleRequest(ctx context.Context) error {
    // Use context from request / 요청의 context 사용
    var result Data
    return httputil.GetContext(ctx, url, &result)
}
```

### 8.3 Timeout Configuration / 타임아웃 설정

✅ **DO: Set appropriate timeouts**

✅ **좋음: 적절한 타임아웃 설정**

```go
// Short timeout for health checks / 헬스 체크용 짧은 타임아웃
httputil.Get(healthURL, &status,
    httputil.WithTimeout(5*time.Second))

// Long timeout for uploads / 업로드용 긴 타임아웃
httputil.Post(uploadURL, data, &result,
    httputil.WithTimeout(5*time.Minute))
```

❌ **DON'T: Use same timeout for all operations**

❌ **나쁨: 모든 작업에 동일한 타임아웃 사용**

```go
// Bad: Same timeout for everything / 나쁨: 모든 것에 동일한 타임아웃
client := httputil.NewClient(
    httputil.WithTimeout(30*time.Second))
```

### 8.4 Error Handling / 에러 처리

✅ **DO: Handle specific error types**

✅ **좋음: 특정 에러 타입 처리**

```go
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsHTTPError(err) {
        // Handle HTTP error / HTTP 에러 처리
    } else if httputil.IsTimeoutError(err) {
        // Handle timeout / 타임아웃 처리
    }
}
```

❌ **DON'T: Ignore error details**

❌ **나쁨: 에러 세부 정보 무시**

```go
// Bad: Generic error handling / 나쁨: 일반적인 에러 처리
err := httputil.Get(url, &result)
if err != nil {
    log.Fatal("Request failed")
}
```

### 8.5 Retry Configuration / 재시도 설정

✅ **DO: Configure retry based on operation criticality**

✅ **좋음: 작업 중요도에 따라 재시도 설정**

```go
// Critical operation: aggressive retries / 중요한 작업: 적극적 재시도
err := httputil.Post(url, data, &result,
    httputil.WithRetry(10),
    httputil.WithRetryBackoff(1*time.Second, 30*time.Second))

// Non-critical: minimal retries / 중요하지 않음: 최소 재시도
err := httputil.Get(url, &result,
    httputil.WithRetry(1))
```

### 8.6 Request Options / 요청 옵션

✅ **DO: Combine options logically**

✅ **좋음: 옵션을 논리적으로 결합**

```go
commonOpts := []httputil.Option{
    httputil.WithTimeout(30*time.Second),
    httputil.WithBearerToken(token),
    httputil.WithRetry(3),
}

// Apply common options / 공통 옵션 적용
err := httputil.Get(url, &result, commonOpts...)
```

### 8.7 Base URL Usage / Base URL 사용

✅ **DO: Use base URL for API clients**

✅ **좋음: API 클라이언트에 base URL 사용**

```go
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com/v1"))

// Clean relative paths / 깔끔한 상대 경로
client.Get("/users", &users)
client.Get("/posts", &posts)
```

❌ **DON'T: Repeat full URLs**

❌ **나쁨: 전체 URL 반복**

```go
// Bad: Repetitive full URLs / 나쁨: 반복적인 전체 URL
httputil.Get("https://api.example.com/v1/users", &users)
httputil.Get("https://api.example.com/v1/posts", &posts)
```

### 8.8 Authentication / 인증

✅ **DO: Handle token refresh**

✅ **좋음: 토큰 새로고침 처리**

```go
func requestWithTokenRefresh(client *httputil.Client, url string, result interface{}) error {
    err := client.Get(url, result)
    if err != nil && httputil.IsHTTPError(err) {
        if httputil.GetStatusCode(err) == 401 {
            // Refresh token and retry / 토큰 새로고침 및 재시도
            newToken := refreshAuthToken()
            return client.Get(url, result,
                httputil.WithBearerToken(newToken))
        }
    }
    return err
}
```

### 8.9 Logging / 로깅

✅ **DO: Log important operations**

✅ **좋음: 중요한 작업 로깅**

```go
log.Printf("Fetching users from %s", url)
err := httputil.Get(url, &users)
if err != nil {
    log.Printf("Failed to fetch users: %v", err)
    return err
}
log.Printf("Successfully fetched %d users", len(users))
```

### 8.10 Testing / 테스트

✅ **DO: Use test servers**

✅ **좋음: 테스트 서버 사용**

```go
func TestAPIClient(t *testing.T) {
    // Create test server / 테스트 서버 생성
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{
            "status": "ok",
        })
    }))
    defer server.Close()

    // Test with test server / 테스트 서버로 테스트
    var result map[string]string
    err := httputil.Get(server.URL, &result)
    if err != nil {
        t.Fatal(err)
    }
}
```

---

## 9. Troubleshooting / 문제 해결

### 9.1 Common Issues / 일반적인 문제

#### Issue 1: Request Timeout / 요청 타임아웃

**Symptom / 증상:**
```
request timeout (URL: GET https://api.example.com/slow)
```

**Solution / 해결책:**
```go
// Increase timeout / 타임아웃 증가
err := httputil.Get(url, &result,
    httputil.WithTimeout(2*time.Minute))
```

#### Issue 2: All Retries Failed / 모든 재시도 실패

**Symptom / 증상:**
```
failed after 3 attempts (URL: POST https://api.example.com/data): network error
```

**Solution / 해결책:**
```go
// Increase retry attempts and backoff / 재시도 횟수 및 백오프 증가
err := httputil.Post(url, data, &result,
    httputil.WithRetry(10),
    httputil.WithRetryBackoff(1*time.Second, 30*time.Second))
```

#### Issue 3: 401 Unauthorized / 401 인증되지 않음

**Symptom / 증상:**
```
HTTP 401 Unauthorized: Unauthorized
```

**Solution / 해결책:**
```go
// Check token / 토큰 확인
err := httputil.Get(url, &result,
    httputil.WithBearerToken("your-valid-token"))

// Or refresh token / 또는 토큰 새로고침
if httputil.GetStatusCode(err) == 401 {
    newToken := refreshToken()
    err = httputil.Get(url, &result,
        httputil.WithBearerToken(newToken))
}
```

#### Issue 4: JSON Decode Error / JSON 디코딩 에러

**Symptom / 증상:**
```
failed to decode response: invalid character '<' looking for beginning of value
```

**Cause / 원인:** Server returned HTML instead of JSON

**서버가 JSON 대신 HTML을 반환함**

**Solution / 해결책:**
```go
// Check response status first / 먼저 응답 상태 확인
var result MyStruct
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsHTTPError(err) {
        // Server returned error HTML / 서버가 에러 HTML 반환
        httpErr := err.(*httputil.HTTPError)
        log.Printf("Response body: %s", httpErr.Body)
    }
}
```

#### Issue 5: Connection Refused / 연결 거부됨

**Symptom / 증상:**
```
failed after 3 attempts: dial tcp: connection refused
```

**Causes / 원인:**
- Server is down / 서버가 다운됨
- Wrong URL / 잘못된 URL
- Network issues / 네트워크 문제

**Solution / 해결책:**
```go
// 1. Verify URL / URL 확인
log.Printf("Connecting to: %s", url)

// 2. Check server status / 서버 상태 확인
// 3. Check network connectivity / 네트워크 연결 확인

// 4. Add more aggressive retries / 더 적극적인 재시도 추가
err := httputil.Get(url, &result,
    httputil.WithRetry(10),
    httputil.WithRetryBackoff(2*time.Second, 1*time.Minute))
```

### 9.2 Debugging Tips / 디버깅 팁

#### Enable Request Logging / 요청 로깅 활성화

```go
func logRequest(method, url string) {
    log.Printf("[HTTP] %s %s", method, url)
}

func logResponse(statusCode int, latency time.Duration) {
    log.Printf("[HTTP] Status %d (latency: %v)", statusCode, latency)
}

// Use in your code / 코드에서 사용
start := time.Now()
logRequest("GET", url)

var result MyStruct
err := httputil.Get(url, &result)

if err != nil {
    log.Printf("[HTTP] Error: %v", err)
} else {
    logResponse(200, time.Since(start))
}
```

#### Check Error Details / 에러 세부 정보 확인

```go
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsHTTPError(err) {
        httpErr := err.(*httputil.HTTPError)
        log.Printf("Status: %d", httpErr.StatusCode)
        log.Printf("URL: %s", httpErr.URL)
        log.Printf("Method: %s", httpErr.Method)
        log.Printf("Body: %s", httpErr.Body)
    }
}
```

#### Test with curl / curl로 테스트

```bash
# Test the same request with curl / curl로 동일한 요청 테스트
curl -v \
  -H "Authorization: Bearer token" \
  -H "Content-Type: application/json" \
  https://api.example.com/endpoint
```

---

## 10. FAQ

### Q1: How do I make a request without decoding JSON? / JSON 디코딩 없이 요청하려면?

**A:** Pass `nil` as the result parameter:

**답:** 결과 매개변수로 `nil`을 전달하세요:

```go
// Just check if request succeeds / 요청 성공 확인만
err := httputil.Get(url, nil)
if err != nil {
    log.Fatal("Request failed")
}
```

### Q2: How do I send raw JSON string? / 원시 JSON 문자열을 전송하려면?

**A:** Use a map or struct that represents the JSON:

**답:** JSON을 나타내는 맵 또는 구조체 사용:

```go
// Send as map / 맵으로 전송
payload := map[string]interface{}{
    "name": "John",
    "age":  30,
}

var result MyResponse
err := httputil.Post(url, payload, &result)
```

### Q3: How do I handle cookies? / 쿠키를 처리하려면?

**A:** Cookie management is fully supported with automatic persistence (v1.10.004+):

**답:** 쿠키 관리는 자동 지속성과 함께 완전히 지원됩니다 (v1.10.004+):

```go
// In-memory cookies (temporary) / 메모리 내 쿠키 (임시)
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithCookies(),
)

// Persistent cookies (saved to file) / 지속성 쿠키 (파일에 저장)
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithPersistentCookies("cookies.json"),
)

// Manual cookie operations / 수동 쿠키 작업
u, _ := url.Parse("https://api.example.com")
client.SetCookie(u, &http.Cookie{Name: "session", Value: "abc123"})
cookies := client.GetCookies(u)
client.SaveCookies() // Save to file / 파일에 저장
```

See section 5.10 for detailed usage patterns.

자세한 사용 패턴은 섹션 5.10을 참조하세요.

### Q4: Can I upload files? / 파일을 업로드할 수 있나요?

**A:** File upload support is planned for Phase 3 (v1.10.005-007).

**답:** 파일 업로드 지원은 Phase 3 (v1.10.005-007)에 계획되어 있습니다.

### Q5: How do I set custom TLS config? / 사용자 정의 TLS 설정을 하려면?

**A:** Currently not directly supported. This is on the roadmap.

**답:** 현재 직접 지원되지 않음. 로드맵에 있습니다.

### Q6: Can I use this with gRPC? / gRPC와 함께 사용할 수 있나요?

**A:** No, httputil is designed for REST/HTTP APIs only.

**답:** 아니요, httputil은 REST/HTTP API 전용으로 설계되었습니다.

### Q7: How do I debug request/response? / 요청/응답을 디버깅하려면?

**A:** Check the error details:

**답:** 에러 세부 정보 확인:

```go
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsHTTPError(err) {
        httpErr := err.(*httputil.HTTPError)
        fmt.Printf("Request: %s %s\n", httpErr.Method, httpErr.URL)
        fmt.Printf("Status: %d\n", httpErr.StatusCode)
        fmt.Printf("Body: %s\n", httpErr.Body)
    }
}
```

### Q8: Is it thread-safe? / 스레드 안전한가요?

**A:** Yes, Client instances are safe for concurrent use.

**답:** 예, Client 인스턴스는 동시 사용이 안전합니다.

```go
client := httputil.NewClient()

// Safe to use from multiple goroutines / 여러 고루틴에서 안전하게 사용
go client.Get("/users", &users1)
go client.Get("/posts", &posts1)
```

### Q9: How do I cancel a request? / 요청을 취소하려면?

**A:** Use Context with cancel:

**답:** 취소가 있는 Context 사용:

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    time.Sleep(1 * time.Second)
    cancel() // Cancel after 1 second / 1초 후 취소
}()

err := httputil.GetContext(ctx, url, &result)
// Returns error if cancelled / 취소되면 에러 반환
```

### Q10: Can I customize retry logic? / 재시도 로직을 사용자 정의할 수 있나요?

**A:** Yes, use `WithRetry` and `WithRetryBackoff`:

**답:** 예, `WithRetry`와 `WithRetryBackoff` 사용:

```go
err := httputil.Get(url, &result,
    httputil.WithRetry(10),                                      // 10 attempts / 10번 시도
    httputil.WithRetryBackoff(500*time.Millisecond, 30*time.Second)) // Custom backoff / 사용자 정의 백오프
```

### Q11: What's the default User-Agent? / 기본 User-Agent는?

**A:** "go-utils/httputil v{version}"

**답:** "go-utils/httputil v{버전}"

```go
// Check version / 버전 확인
fmt.Println(httputil.Version) // v1.10.001
```

### Q12: How do I contribute? / 어떻게 기여하나요?

**A:** See [CONTRIBUTING.md](../../CONTRIBUTING.md)

**답:** [CONTRIBUTING.md](../../CONTRIBUTING.md) 참조

---

## Additional Resources / 추가 자료

- **Package README**: [httputil/README.md](../../httputil/README.md)
- **Developer Guide**: [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md)
- **Design Plan**: [DESIGN_PLAN.md](DESIGN_PLAN.md)
- **Work Plan**: [WORK_PLAN.md](WORK_PLAN.md)
- **Changelog**: [CHANGELOG-v1.10.md](../CHANGELOG/CHANGELOG-v1.10.md)
- **GitHub Repository**: https://github.com/arkd0ng/go-utils

---

**Last Updated / 최종 업데이트**: 2025-10-15
**Version / 버전**: v1.10.001
**Author / 작성자**: arkd0ng
