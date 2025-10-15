# httputil Package - Developer Guide
# httputil 패키지 - 개발자 가이드

**Version / 버전**: v1.10.002
**Last Updated / 최종 업데이트**: 2025-10-15

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#1-architecture-overview--아키텍처-개요)
2. [Package Structure / 패키지 구조](#2-package-structure--패키지-구조)
3. [Core Components / 핵심 컴포넌트](#3-core-components--핵심-컴포넌트)
4. [Internal Implementation / 내부 구현](#4-internal-implementation--내부-구현)
5. [Design Patterns / 디자인 패턴](#5-design-patterns--디자인-패턴)
6. [Adding New Features / 새 기능 추가](#6-adding-new-features--새-기능-추가)
7. [Testing Guide / 테스트 가이드](#7-testing-guide--테스트-가이드)
8. [Performance / 성능](#8-performance--성능)
9. [Contributing Guidelines / 기여 가이드라인](#9-contributing-guidelines--기여-가이드라인)
10. [Code Style / 코드 스타일](#10-code-style--코드-스타일)

---

## 1. Architecture Overview / 아키텍처 개요

### 1.1 Design Philosophy / 설계 철학

httputil은 "극도의 간결함 (Extreme Simplicity)" 원칙을 따릅니다:

**Core Principle / 핵심 원칙:**
- **30+ lines → 2-3 lines**: Reduce boilerplate code dramatically
- **30줄 이상 → 2-3줄**: 보일러플레이트 코드를 극적으로 감소
- **Zero Configuration**: Sensible defaults for 99% of use cases
- **제로 설정**: 99% 사용 사례에 대한 합리적인 기본값
- **Auto Everything**: Automatic JSON handling, retry, error wrapping
- **모든 것 자동화**: 자동 JSON 처리, 재시도, 에러 래핑

### 1.2 High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────────────────┐
│                         User Code                                │
│                         사용자 코드                               │
└─────────────────────────────────┬───────────────────────────────┘
                                  │
                  ┌───────────────┴───────────────┐
                  │                               │
         ┌────────▼────────┐           ┌────────▼────────┐
         │   Simple API    │           │  Client API     │
         │   간단한 API     │           │  클라이언트 API  │
         │  (simple.go)    │           │  (client.go)    │
         └────────┬────────┘           └────────┬────────┘
                  │                               │
                  └───────────────┬───────────────┘
                                  │
                     ┌────────────▼────────────┐
                     │   Options Pattern       │
                     │   옵션 패턴              │
                     │   (options.go)          │
                     └────────────┬────────────┘
                                  │
                     ┌────────────▼────────────┐
                     │   HTTP Client Core      │
                     │   HTTP 클라이언트 핵심   │
                     │   (client.go)           │
                     └────────────┬────────────┘
                                  │
         ┌────────────────────────┼────────────────────────┐
         │                        │                        │
    ┌────▼─────┐         ┌───────▼───────┐      ┌────────▼────────┐
    │  Retry   │         │  JSON Codec   │      │  Error Types    │
    │  Logic   │         │  JSON 코덱     │      │  에러 타입       │
    │  재시도  │         │               │      │  (errors.go)    │
    └──────────┘         └───────────────┘      └─────────────────┘
         │                        │                        │
         └────────────────────────┼────────────────────────┘
                                  │
                     ┌────────────▼────────────┐
                     │    Standard Library     │
                     │    표준 라이브러리       │
                     │  net/http, context      │
                     └─────────────────────────┘
```

### 1.3 Design Goals / 설계 목표

1. **Simplicity / 간결성**
   - Minimal API surface
   - Intuitive method names
   - 최소 API 표면
   - 직관적인 메서드 이름

2. **Reliability / 신뢰성**
   - Automatic retry with exponential backoff
   - Rich error types with context
   - 지수 백오프를 통한 자동 재시도
   - 컨텍스트가 있는 풍부한 에러 타입

3. **Performance / 성능**
   - Connection reuse
   - Minimal allocations
   - 연결 재사용
   - 최소 할당

4. **Safety / 안전성**
   - Type-safe operations
   - Context support for cancellation
   - 타입 안전 작업
   - 취소를 위한 Context 지원

5. **Zero Dependencies / 제로 의존성**
   - Standard library only
   - 표준 라이브러리만

---

## 2. Package Structure / 패키지 구조

### 2.1 File Organization / 파일 구성

```
httputil/
├── httputil.go       # Package doc, version management / 패키지 문서, 버전 관리
├── client.go         # Core HTTP client / 핵심 HTTP 클라이언트
├── simple.go         # Package-level convenience functions / 패키지 레벨 편의 함수
├── options.go        # Functional options pattern / 함수형 옵션 패턴
├── errors.go         # Error types / 에러 타입
├── response.go       # Response wrapper with helpers / 헬퍼가 있는 응답 래퍼
├── file.go           # File download/upload operations / 파일 다운로드/업로드 작업
├── url.go            # URL builder and utilities / URL 빌더 및 유틸리티
├── form.go           # Form builder and utilities / Form 빌더 및 유틸리티
├── httputil_test.go  # Tests / 테스트
└── README.md         # Package documentation / 패키지 문서
```

### 2.2 File Responsibilities / 파일별 책임

#### httputil.go (164 lines)
- Package-level documentation
- Version management via cfg/app.yaml
- Version constant export
- 패키지 레벨 문서
- cfg/app.yaml을 통한 버전 관리
- 버전 상수 내보내기

#### client.go (290 lines)
- `Client` struct definition
- Core HTTP methods (GET, POST, PUT, PATCH, DELETE)
- Context variants for all methods
- Retry logic with exponential backoff
- Request/response handling
- Client 구조체 정의
- 핵심 HTTP 메서드
- 모든 메서드의 Context 변형
- 지수 백오프를 통한 재시도 로직
- 요청/응답 처리

#### simple.go (142 lines)
- Package-level convenience functions
- Default client instance
- Wrappers around Client methods
- `SetDefaultClient` for global configuration
- 패키지 레벨 편의 함수
- 기본 클라이언트 인스턴스
- Client 메서드 래퍼
- 전역 설정을 위한 SetDefaultClient

#### options.go (236 lines)
- `Option` functional option type
- `config` struct with all configuration
- 12 built-in options
- Default configuration factory
- Option 함수형 옵션 타입
- 모든 설정이 있는 config 구조체
- 12개 내장 옵션
- 기본 설정 팩토리

#### errors.go (104 lines)
- `HTTPError` type (status code, body, URL)
- `RetryError` type (failed attempts)
- `TimeoutError` type (timeout detection)
- Helper functions (IsHTTPError, GetStatusCode, etc.)
- HTTPError 타입
- RetryError 타입
- TimeoutError 타입
- 헬퍼 함수

#### response.go (280 lines) **[Phase 2]**
- `Response` wrapper struct embedding *http.Response
- Body methods (Body(), String(), JSON())
- Status check methods (IsSuccess(), IsOK(), IsNotFound(), etc. - 15 methods)
- Header methods (Header(), Headers(), ContentType())
- Cached body for multiple reads
- Response 래퍼 구조체 (*http.Response 임베딩)
- 본문 메서드 (Body(), String(), JSON())
- 상태 확인 메서드 (IsSuccess(), IsOK(), IsNotFound() 등 - 15개 메서드)
- 헤더 메서드 (Header(), Headers(), ContentType())
- 여러 번 읽기를 위한 캐시된 본문

#### file.go (340 lines) **[Phase 3]**
- `ProgressFunc` callback type for progress tracking
- `progressReader` for progress callbacks
- File download methods (DownloadFile, DownloadFileContext, Download, DownloadContext)
- File upload methods (UploadFile, UploadFileContext, UploadFiles, UploadFilesContext)
- Multipart form data handling
- 진행 상황 추적을 위한 ProgressFunc 콜백 타입
- 진행 상황 콜백을 위한 progressReader
- 파일 다운로드 메서드
- 파일 업로드 메서드
- multipart form data 처리

#### url.go (180 lines) **[Phase 4]**
- `URLBuilder` fluent API struct
- URL building methods (Path(), Param(), ParamIf(), Build())
- URL utility functions (JoinURL, AddQueryParams, GetDomain, GetScheme, etc. - 6 functions)
- URLBuilder Fluent API 구조체
- URL 구축 메서드 (Path(), Param(), ParamIf(), Build())
- URL 유틸리티 함수 (JoinURL, AddQueryParams, GetDomain, GetScheme 등 - 6개 함수)

#### form.go (200 lines) **[Phase 4]**
- `FormBuilder` fluent API struct
- Form building methods (Set(), Add(), AddIf(), AddMultiple(), etc.)
- Form utility functions (ParseForm, EncodeForm)
- PostForm methods for Client and package-level
- FormBuilder Fluent API 구조체
- 폼 구축 메서드 (Set(), Add(), AddIf(), AddMultiple() 등)
- 폼 유틸리티 함수 (ParseForm, EncodeForm)
- Client 및 패키지 레벨용 PostForm 메서드

#### httputil_test.go (456 lines)
- Unit tests for all public API (Phase 1-4)
- Error type tests
- Configuration tests
- Response wrapper tests (17 sub-tests)
- URL Builder tests (3 sub-tests)
- Form Builder tests (5 sub-tests)
- URL utilities tests (6 sub-tests)
- Form utilities tests (2 sub-tests)
- Progress reader tests
- Total: 13 tests, 43+ sub-tests
- 모든 공개 API에 대한 단위 테스트 (1-4단계)
- 에러 타입 테스트
- 설정 테스트
- Response 래퍼 테스트 (17개 하위 테스트)
- URL Builder 테스트 (3개 하위 테스트)
- Form Builder 테스트 (5개 하위 테스트)
- URL 유틸리티 테스트 (6개 하위 테스트)
- Form 유틸리티 테스트 (2개 하위 테스트)
- Progress reader 테스트
- 총 13개 테스트, 43개 이상 하위 테스트

---

## 3. Core Components / 핵심 컴포넌트

### 3.1 Client Struct / Client 구조체

**Definition / 정의:**
```go
// Client wraps http.Client with additional functionality.
// Client는 추가 기능을 가진 http.Client를 래핑합니다.
type Client struct {
    client *http.Client  // Standard library HTTP client / 표준 라이브러리 HTTP 클라이언트
    config *config       // Configuration / 설정
}
```

**Responsibilities / 책임:**
- Wrap standard `http.Client`
- Store configuration
- Provide HTTP methods
- Handle retry logic
- 표준 http.Client 래핑
- 설정 저장
- HTTP 메서드 제공
- 재시도 로직 처리

**Thread Safety / 스레드 안전성:**
- `Client` instances are safe for concurrent use
- `http.Client` is thread-safe
- `config` is read-only after creation
- Client 인스턴스는 동시 사용이 안전함
- http.Client는 스레드 안전함
- config는 생성 후 읽기 전용

### 3.2 Config Struct / Config 구조체

**Definition / 정의:**
```go
type config struct {
    // Request configuration / 요청 설정
    headers     map[string]string
    queryParams map[string]string
    timeout     time.Duration
    userAgent   string

    // Authentication / 인증
    bearerToken   string
    basicAuthUser string
    basicAuthPass string

    // Retry configuration / 재시도 설정
    maxRetries int
    retryMin   time.Duration
    retryMax   time.Duration

    // Client configuration / 클라이언트 설정
    baseURL         string
    followRedirects bool
    maxRedirects    int
    tlsConfig       *tls.Config
    proxyURL        string
    cookieJar       http.CookieJar

    // Logging / 로깅
    logger Logger
}
```

**Default Values / 기본값:**
```go
func defaultConfig() *config {
    return &config{
        headers:         make(map[string]string),
        queryParams:     make(map[string]string),
        timeout:         30 * time.Second,     // 30 seconds / 30초
        userAgent:       "go-utils/httputil v" + Version,
        maxRetries:      3,                     // 3 attempts / 3번 시도
        retryMin:        100 * time.Millisecond, // Min backoff / 최소 백오프
        retryMax:        5 * time.Second,       // Max backoff / 최대 백오프
        followRedirects: true,                  // Follow redirects / 리디렉션 따르기
        maxRedirects:    10,                    // Max 10 redirects / 최대 10개 리디렉션
    }
}
```

### 3.3 Error Types / 에러 타입

#### HTTPError

**Purpose / 목적:** Represent HTTP errors with full context

**HTTP 에러를 전체 컨텍스트와 함께 나타냄**

```go
type HTTPError struct {
    StatusCode int    // e.g., 404, 500
    Status     string // e.g., "404 Not Found"
    Body       string // Response body for debugging / 디버깅용 응답 본문
    URL        string // Full request URL / 전체 요청 URL
    Method     string // e.g., "GET", "POST"
}
```

**Error Message Format / 에러 메시지 형식:**
```
HTTP 404 404 Not Found: Not Found (URL: GET https://api.example.com/users/999, Body: {"error":"user not found"})
```

#### RetryError

**Purpose / 목적:** Track failed retry attempts

**실패한 재시도 시도 추적**

```go
type RetryError struct {
    Attempts int   // Number of attempts / 시도 횟수
    LastErr  error // Last error / 마지막 에러
    URL      string
    Method   string
}
```

**Error Message Format / 에러 메시지 형식:**
```
request failed after 3 attempts (URL: POST https://api.example.com/data): network error
```

#### TimeoutError

**Purpose / 목적:** Indicate request timeout

**요청 타임아웃 표시**

```go
type TimeoutError struct {
    URL    string
    Method string
}
```

**Error Message Format / 에러 메시지 형식:**
```
request timeout (URL: GET https://api.example.com/slow)
```

### 3.4 Option Type / Option 타입

**Definition / 정의:**
```go
// Option is a functional option for configuring HTTP requests and clients.
// Option은 HTTP 요청 및 클라이언트를 설정하기 위한 함수형 옵션입니다.
type Option func(*config)
```

**Built-in Options (12) / 내장 옵션 (12개):**

1. **Request Configuration / 요청 설정:**
   - `WithTimeout(time.Duration)`
   - `WithHeaders(map[string]string)`
   - `WithHeader(key, value string)`
   - `WithQueryParams(map[string]string)`
   - `WithUserAgent(string)`

2. **Authentication / 인증:**
   - `WithBearerToken(string)`
   - `WithBasicAuth(user, pass string)`

3. **Retry Configuration / 재시도 설정:**
   - `WithRetry(maxRetries int)`
   - `WithRetryBackoff(min, max time.Duration)`

4. **Client Configuration / 클라이언트 설정:**
   - `WithBaseURL(string)`
   - `WithFollowRedirects(bool)`
   - `WithMaxRedirects(int)`

### 3.5 Response Wrapper / Response 래퍼 **[Phase 2]**

The `Response` type wraps `*http.Response` with additional helper methods:

Response 타입은 추가 헬퍼 메서드로 *http.Response를 래핑합니다:

```go
type Response struct {
    *http.Response
    body []byte  // Cached for multiple reads / 여러 번 읽기를 위해 캐시됨
}
```

**Key Methods / 주요 메서드:**

1. **Body Access / 본문 접근:**
   - `Body() []byte` - Get body as bytes
   - `String() string` - Get body as string
   - `JSON(result interface{}) error` - Decode JSON

2. **Status Checks / 상태 확인:**
   - `IsSuccess() bool` - 2xx status codes
   - `IsError() bool` - 4xx or 5xx
   - `IsClientError() bool` - 4xx
   - `IsServerError() bool` - 5xx
   - Plus 11 specific status checks (IsOK, IsNotFound, etc.)

3. **Header Access / 헤더 접근:**
   - `Header(key string) string` - Single header
   - `Headers() map[string]string` - All headers
   - `ContentType() string` - Content-Type header

**Design Benefits / 설계 이점:**
- Body caching allows multiple reads without re-fetching
- Intuitive status check methods reduce boilerplate
- Type-safe JSON decoding
- 본문 캐싱으로 다시 가져오지 않고 여러 번 읽기 가능
- 직관적인 상태 확인 메서드로 보일러플레이트 감소
- 타입 안전 JSON 디코딩

### 3.6 File Operations / 파일 작업 **[Phase 3]**

File download and upload with progress tracking:

진행 상황 추적과 함께 파일 다운로드 및 업로드:

```go
type ProgressFunc func(bytesRead, totalBytes int64)

type progressReader struct {
    reader   io.Reader
    progress ProgressFunc
    total    int64
    read     int64
}
```

**Download Methods / 다운로드 메서드:**
- `DownloadFile(url, filepath string, opts ...Option) error`
- `DownloadFileContext(ctx, url, filepath string, progress ProgressFunc, opts ...Option) error`
- `Download(url string, opts ...Option) ([]byte, error)`
- `DownloadContext(ctx, url string, opts ...Option) ([]byte, error)`

**Upload Methods / 업로드 메서드:**
- `UploadFile(url, fieldName, filepath string, result interface{}, opts ...Option) error`
- `UploadFileContext(ctx, url, fieldName, filepath string, result interface{}, opts ...Option) error`
- `UploadFiles(url string, files map[string]string, result interface{}, opts ...Option) error`
- `UploadFilesContext(ctx, url string, files map[string]string, result interface{}, opts ...Option) error`

**Progress Tracking / 진행 상황 추적:**
The `progressReader` wraps an `io.Reader` and calls the progress callback after each read, allowing real-time progress monitoring for large file operations.

progressReader는 io.Reader를 래핑하고 각 읽기 후 진행 상황 콜백을 호출하여 대용량 파일 작업에 대한 실시간 진행 상황 모니터링을 가능하게 합니다.

### 3.7 URL Builder / URL 빌더 **[Phase 4]**

Fluent API for building URLs with parameters:

매개변수와 함께 URL을 구축하기 위한 Fluent API:

```go
type URLBuilder struct {
    baseURL string
    path    []string
    params  url.Values
}
```

**Builder Methods / 빌더 메서드:**
- `NewURL(baseURL string) *URLBuilder` - Create builder
- `Path(segments ...string) *URLBuilder` - Add path segments
- `Param(key, value string) *URLBuilder` - Add single parameter
- `Params(map[string]string) *URLBuilder` - Add multiple parameters
- `ParamIf(condition bool, key, value string) *URLBuilder` - Conditional parameter
- `Build() string` - Build final URL

**Utility Functions / 유틸리티 함수:**
- `JoinURL(baseURL string, paths ...string) string`
- `AddQueryParams(urlStr string, params map[string]string) (string, error)`
- `GetDomain(urlStr string) (string, error)`
- `GetScheme(urlStr string) (string, error)`
- `IsAbsoluteURL(urlStr string) bool`
- `NormalizeURL(urlStr string) string`

**Design Pattern / 디자인 패턴:**
Method chaining (Fluent API) allows building complex URLs in a readable, declarative way.

메서드 체이닝 (Fluent API)은 읽기 쉽고 선언적인 방식으로 복잡한 URL을 구축할 수 있게 합니다.

### 3.8 Form Builder / Form 빌더 **[Phase 4]**

Fluent API for building form data:

폼 데이터를 구축하기 위한 Fluent API:

```go
type FormBuilder struct {
    values url.Values
}
```

**Builder Methods / 빌더 메서드:**
- `NewForm() *FormBuilder` - Create builder
- `Add(key, value string) *FormBuilder` - Add value (allows duplicates)
- `Set(key, value string) *FormBuilder` - Set value (replaces existing)
- `AddMultiple(key string, values ...string) *FormBuilder` - Add multiple values
- `AddIf(condition bool, key, value string) *FormBuilder` - Conditional add
- `Get(key string) string` - Get single value
- `GetAll(key string) []string` - Get all values
- `Has(key string) bool` - Check if key exists
- `Del(key string) *FormBuilder` - Delete key
- `Clone() *FormBuilder` - Clone form
- `Map() map[string]string` - Convert to map
- `Encode() string` - Encode as URL-encoded string

**Utility Functions / 유틸리티 함수:**
- `ParseForm(data string) (map[string]string, error)` - Parse URL-encoded string
- `EncodeForm(data map[string]string) string` - Encode map to URL-encoded string

**PostForm Methods / PostForm 메서드:**
- `PostForm(url string, data map[string]string, result interface{}, opts ...Option) error`
- `PostFormContext(ctx, url string, data map[string]string, result interface{}, opts ...Option) error`
- `Client.PostForm(path string, data map[string]string, result interface{}, opts ...Option) error`
- `Client.PostFormContext(ctx, path string, data map[string]string, result interface{}, opts ...Option) error`

---

### 3.9 Cookie Management / 쿠키 관리 **[Phase 5a]**

Thread-safe cookie jar with optional file persistence:

파일 지속성을 가진 스레드 안전 쿠키 저장소:

```go
// CookieJar manages HTTP cookies with optional persistence
// CookieJar는 선택적 지속성을 가진 HTTP 쿠키를 관리합니다
type CookieJar struct {
    jar      http.CookieJar    // Standard library cookie jar
    filePath string             // Optional file path for persistence
    mu       sync.RWMutex       // Thread-safe access
}
```

**Architecture Design / 아키텍처 설계:**

```
┌─────────────────────────────────────────────────────┐
│                   Client                            │
│  ┌───────────────────────────────────────────────┐  │
│  │           cookieJar: *CookieJar               │  │
│  └───────────────┬───────────────────────────────┘  │
│                  │                                   │
│                  │ delegates cookie operations       │
│                  │ 쿠키 작업 위임                     │
│                  ▼                                   │
│  ┌───────────────────────────────────────────────┐  │
│  │          CookieJar                            │  │
│  │  ┌─────────────────────────────────────────┐  │  │
│  │  │  jar: http.CookieJar (stdlib)           │  │  │
│  │  │  filePath: string (optional)            │  │  │
│  │  │  mu: sync.RWMutex (thread-safe)        │  │  │
│  │  └─────────────────────────────────────────┘  │  │
│  │                                               │  │
│  │  Operations / 작업:                           │  │
│  │  • SetCookies(u, cookies)                    │  │
│  │  • GetCookies(u) []*http.Cookie             │  │
│  │  • ClearCookies()                            │  │
│  │  • SaveCookies() → JSON file                │  │
│  │  • LoadCookies() ← JSON file                │  │
│  │  • HasCookie(u, name) bool                   │  │
│  │  • GetCookie(u, name) *http.Cookie          │  │
│  │  • RemoveCookie(u, name)                    │  │
│  │  • CountCookies(u) int                      │  │
│  └───────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────┘
```

**Key Components / 주요 컴포넌트:**

1. **CookieJar Struct / CookieJar 구조체:**
   - Wraps `http.CookieJar` from standard library
   - Optional file persistence with JSON serialization
   - Thread-safe with `sync.RWMutex`
   - 표준 라이브러리의 http.CookieJar 래핑
   - JSON 직렬화를 사용한 선택적 파일 지속성
   - sync.RWMutex로 스레드 안전

2. **Constructor Methods / 생성자 메서드:**
   ```go
   // Create in-memory cookie jar / 메모리 내 쿠키 저장소 생성
   func NewCookieJar() (*CookieJar, error)

   // Create persistent cookie jar / 지속성 쿠키 저장소 생성
   func NewPersistentCookieJar(filePath string) (*CookieJar, error)
   ```

3. **Client Integration / 클라이언트 통합:**
   - Client has optional `cookieJar *CookieJar` field
   - Initialized via options: `WithCookies()`, `WithPersistentCookies(path)`
   - Client delegates all cookie operations to CookieJar
   - Client는 선택적 cookieJar 필드 보유
   - 옵션을 통해 초기화: WithCookies(), WithPersistentCookies()
   - Client는 모든 쿠키 작업을 CookieJar에 위임

4. **Persistence Mechanism / 지속성 메커니즘:**
   ```go
   // Cookie entry for JSON serialization / JSON 직렬화를 위한 쿠키 엔트리
   type cookieEntry struct {
       URL     string
       Cookies []*http.Cookie
   }

   // Save: iterate jar → serialize to JSON → write file (0600)
   // Load: read file → deserialize JSON → populate jar
   // 저장: jar 순회 → JSON 직렬화 → 파일 작성 (0600)
   // 로드: 파일 읽기 → JSON 역직렬화 → jar 채우기
   ```

**Implementation Details / 구현 세부사항:**

1. **Thread Safety / 스레드 안전성:**
   ```go
   func (cj *CookieJar) SetCookie(u *url.URL, cookie *http.Cookie) {
       cj.mu.Lock()         // Write lock / 쓰기 잠금
       defer cj.mu.Unlock()
       // ... operation
   }

   func (cj *CookieJar) GetCookies(u *url.URL) []*http.Cookie {
       cj.mu.RLock()        // Read lock / 읽기 잠금
       defer cj.mu.RUnlock()
       // ... operation
   }
   ```

2. **Automatic Loading / 자동 로딩:**
   - When `NewPersistentCookieJar` is called, it automatically loads cookies if file exists
   - No manual `LoadCookies()` call needed on initialization
   - NewPersistentCookieJar 호출 시 파일이 존재하면 자동으로 쿠키 로드
   - 초기화 시 수동 LoadCookies() 호출 불필요

3. **File Permissions / 파일 권한:**
   - Cookie files are saved with `0600` permissions (owner read/write only)
   - Security: prevents other users from reading session cookies
   - 쿠키 파일은 0600 권한으로 저장 (소유자 읽기/쓰기만)
   - 보안: 다른 사용자가 세션 쿠키를 읽는 것 방지

4. **Public Suffix List / 공용 접미사 목록:**
   - Uses `golang.org/x/net/publicsuffix` for proper domain handling
   - Prevents cookies from being set on TLDs (e.g., `.com`, `.org`)
   - 적절한 도메인 처리를 위해 golang.org/x/net/publicsuffix 사용
   - TLD에 쿠키 설정 방지 (예: .com, .org)

**Client Methods / 클라이언트 메서드:**

All client methods check if `cookieJar` is nil and gracefully handle the case:

모든 클라이언트 메서드는 cookieJar가 nil인지 확인하고 우아하게 처리:

```go
func (c *Client) GetCookies(u *url.URL) []*http.Cookie {
    if c.cookieJar == nil {
        return nil  // Graceful degradation / 우아한 성능 저하
    }
    return c.cookieJar.GetCookies(u)
}

func (c *Client) SetCookie(u *url.URL, cookie *http.Cookie) {
    if c.cookieJar == nil {
        return  // No-op if cookie jar not enabled / 쿠키 저장소 미활성화 시 무작동
    }
    c.cookieJar.SetCookie(u, cookie)
}
```

**Configuration Options / 설정 옵션:**

```go
// Option 1: In-memory cookies (temporary) / 옵션 1: 메모리 내 쿠키 (임시)
func WithCookies() Option {
    return func(c *config) {
        c.enableCookieJar = true
    }
}

// Option 2: Persistent cookies (file storage) / 옵션 2: 지속성 쿠키 (파일 저장)
func WithPersistentCookies(filePath string) Option {
    return func(c *config) {
        c.cookieJarPath = filePath
    }
}

// Option 3: Custom cookie jar / 옵션 3: 사용자 정의 쿠키 저장소
func WithCookieJar(jar http.CookieJar) Option {
    return func(c *config) {
        c.cookieJar = jar
    }
}
```

**Usage Patterns / 사용 패턴:**

1. **Session-Based Authentication / 세션 기반 인증:**
   ```go
   client := httputil.NewClient(
       httputil.WithBaseURL("https://api.example.com"),
       httputil.WithPersistentCookies("session.json"),
   )

   // Login once / 한 번 로그인
   client.Post("/login", credentials, &result)
   client.SaveCookies()

   // Later: reuse session / 나중에: 세션 재사용
   client2 := httputil.NewClient(
       httputil.WithPersistentCookies("session.json"), // Auto-loads / 자동 로드
   )
   client2.Get("/profile", &profile) // Authenticated / 인증됨
   ```

2. **Multi-Step Workflow / 다단계 워크플로우:**
   ```go
   client := httputil.NewClient(httputil.WithCookies())

   // Step 1: Start workflow / 1단계: 워크플로우 시작
   client.Post("/workflow/start", nil, &step1)

   // Step 2: Server sets tracking cookie / 2단계: 서버가 추적 쿠키 설정
   client.Post("/workflow/step2", data, &step2)

   // Step 3: Cookie automatically sent / 3단계: 쿠키 자동 전송
   client.Post("/workflow/complete", final, &result)
   ```

**Performance / 성능:**

- **SetCookie**: ~680 ns/op, 128 B/op, 6 allocs/op
- **GetCookies**: ~600 ns/op, 424 B/op, 4 allocs/op
- **HasCookie**: ~600 ns/op, 424 B/op, 4 allocs/op

Benchmarks show minimal overhead for cookie operations.

벤치마크는 쿠키 작업에 대한 최소 오버헤드를 보여줍니다.

**Testing Strategy / 테스트 전략:**

23 test functions covering:
- Basic operations (NewCookieJar, SetCookies, GetCookies)
- Persistence (SaveAndLoadCookies)
- Client integration (CookieIntegration, PersistentCookies)
- Edge cases (ThreadSafety, ExpiredCookies, NoCookieJar)
- Utility methods (HasCookie, GetCookie, CountCookies, RemoveCookie)

23개 테스트 함수가 다음을 커버:
- 기본 작업
- 지속성
- 클라이언트 통합
- 엣지 케이스
- 유틸리티 메서드

---

## 4. Internal Implementation / 내부 구현

### 4.1 Request Flow / 요청 흐름

```
User calls Get/Post/etc
사용자가 Get/Post 등 호출
         │
         ▼
┌────────────────────┐
│  Simple API        │
│  (simple.go)       │
│  httputil.Get()    │
└────────┬───────────┘
         │
         ▼
┌────────────────────┐
│  Client Method     │
│  client.Get()      │
└────────┬───────────┘
         │
         ▼
┌────────────────────┐
│  doRequest()       │
│  Core logic        │
│  핵심 로직          │
└────────┬───────────┘
         │
         ├──────────────┬──────────────┬──────────────┐
         │              │              │              │
         ▼              ▼              ▼              ▼
┌─────────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────┐
│ Build URL   │  │ Marshal  │  │ Set      │  │ Retry Loop   │
│ URL 구축    │  │ JSON     │  │ Headers  │  │ 재시도 루프   │
└─────────────┘  └──────────┘  └──────────┘  └──────┬───────┘
                                                      │
                                          ┌───────────┴────────────┐
                                          │                        │
                                          ▼                        ▼
                                   ┌──────────┐            ┌─────────────┐
                                   │ Success  │            │ Error       │
                                   │ 성공     │            │ (Retry?)    │
                                   └────┬─────┘            │ 에러(재시도?) │
                                        │                  └─────────────┘
                                        ▼
                                 ┌──────────┐
                                 │ Unmarshal│
                                 │ JSON     │
                                 └────┬─────┘
                                      │
                                      ▼
                                 ┌──────────┐
                                 │ Return   │
                                 │ 반환     │
                                 └──────────┘
```

### 4.2 doRequest Method / doRequest 메서드

**Signature / 시그니처:**
```go
func (c *Client) doRequest(
    ctx context.Context,
    method string,
    path string,
    body interface{},
    result interface{},
    opts ...Option,
) error
```

**Steps / 단계:**

1. **Merge Configuration / 설정 병합**
   ```go
   cfg := *c.config
   cfg.apply(opts)
   ```

2. **Build Full URL / 전체 URL 구축**
   ```go
   fullURL := path
   if cfg.baseURL != "" && !isAbsoluteURL(path) {
       fullURL = cfg.baseURL + "/" + strings.TrimLeft(path, "/")
   }
   ```

3. **Add Query Parameters / 쿼리 매개변수 추가**
   ```go
   if len(cfg.queryParams) > 0 {
       u, _ := url.Parse(fullURL)
       q := u.Query()
       for k, v := range cfg.queryParams {
           q.Add(k, v)
       }
       u.RawQuery = q.Encode()
       fullURL = u.String()
   }
   ```

4. **Marshal Request Body / 요청 본문 마샬링**
   ```go
   var bodyReader io.Reader
   if body != nil {
       jsonData, err := json.Marshal(body)
       if err != nil {
           return fmt.Errorf("failed to marshal request body: %w", err)
       }
       bodyReader = bytes.NewReader(jsonData)
   }
   ```

5. **Retry Loop / 재시도 루프**
   ```go
   for attempt := 0; attempt <= cfg.maxRetries; attempt++ {
       // Create request / 요청 생성
       req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)

       // Set headers / 헤더 설정
       req.Header.Set("Content-Type", "application/json")
       req.Header.Set("User-Agent", cfg.userAgent)
       // ... more headers

       // Execute request / 요청 실행
       resp, err := c.client.Do(req)

       // Handle errors / 에러 처리
       // Retry on network errors and 5xx / 네트워크 에러 및 5xx 시 재시도

       // Calculate backoff / 백오프 계산
       backoff := calculateBackoff(attempt, cfg.retryMin, cfg.retryMax)
       time.Sleep(backoff)
   }
   ```

6. **Unmarshal Response / 응답 언마샬링**
   ```go
   if result != nil {
       if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
           return fmt.Errorf("failed to decode response: %w", err)
       }
   }
   ```

### 4.3 Retry Logic / 재시도 로직

**Retry Conditions / 재시도 조건:**
- Network errors (connection refused, timeout)
- 5xx server errors
- NOT 4xx client errors (user's fault)
- 네트워크 에러 (연결 거부, 타임아웃)
- 5xx 서버 에러
- 4xx 클라이언트 에러는 재시도 안 함 (사용자 오류)

**Backoff Calculation / 백오프 계산:**
```go
func calculateBackoff(attempt int, min, max time.Duration) time.Duration {
    // Exponential backoff / 지수 백오프
    backoff := min * time.Duration(math.Pow(2, float64(attempt)))
    if backoff > max {
        backoff = max
    }

    // Add jitter (±25%) to prevent thundering herd
    // 지터 추가 (±25%) - 썬더링 허드 방지
    jitter := time.Duration(rand.Int63n(int64(backoff / 4)))
    if rand.Intn(2) == 0 {
        backoff += jitter
    } else {
        backoff -= jitter
    }

    return backoff
}
```

**Example Backoff Sequence / 백오프 시퀀스 예제:**
- Attempt 0: 100ms ± 25ms = 75-125ms
- Attempt 1: 200ms ± 50ms = 150-250ms
- Attempt 2: 400ms ± 100ms = 300-500ms
- Attempt 3: 800ms ± 200ms = 600-1000ms
- Attempt 4: 1600ms ± 400ms = 1200-2000ms
- Max: 5000ms

### 4.4 Error Handling Flow / 에러 처리 흐름

```
Execute HTTP Request
HTTP 요청 실행
         │
         ▼
    ┌────────┐
    │Success?│
    │성공?   │
    └───┬────┘
        │
    ┌───┴────┐
    │        │
    Yes      No
    예       아니오
    │        │
    │        ▼
    │   ┌──────────────┐
    │   │ Error Type?  │
    │   │ 에러 타입?    │
    │   └───┬──────────┘
    │       │
    │   ┌───┴────────────────────────┬──────────────────┐
    │   │                            │                  │
    │   ▼                            ▼                  ▼
    │ ┌─────────────┐         ┌──────────┐      ┌────────────┐
    │ │Network Error│         │4xx Error │      │5xx Error   │
    │ │네트워크 에러 │         │          │      │            │
    │ └──────┬──────┘         └────┬─────┘      └─────┬──────┘
    │        │                     │                   │
    │        ▼                     ▼                   ▼
    │  ┌──────────┐          ┌──────────┐       ┌──────────┐
    │  │Retry?    │          │Return    │       │Retry?    │
    │  │재시도?   │          │HTTPError │       │재시도?   │
    │  └────┬─────┘          └──────────┘       └────┬─────┘
    │       │                                         │
    │   ┌───┴─────┐                              ┌───┴─────┐
    │   │         │                              │         │
    │  Yes        No                            Yes        No
    │  예         아니오                         예         아니오
    │   │         │                              │         │
    │   ▼         ▼                              │         ▼
    │ ┌────┐  ┌────────┐                        │    ┌────────┐
    │ │Wait│  │Return  │                        │    │Return  │
    │ │대기│  │Retry   │                        │    │HTTP    │
    │ │    │  │Error   │                        │    │Error   │
    │ └─┬──┘  └────────┘                        │    └────────┘
    │   │                                        │
    │   └────────────────────────────────────────┘
    │
    ▼
┌──────────┐
│Decode    │
│JSON      │
└────┬─────┘
     │
     ▼
┌──────────┐
│Return    │
│Success   │
└──────────┘
```

---

## 5. Design Patterns / 디자인 패턴

### 5.1 Functional Options Pattern / 함수형 옵션 패턴

**Purpose / 목적:** Flexible configuration without breaking API

**API를 깨뜨리지 않고 유연한 설정**

**Implementation / 구현:**
```go
// Option is a function that modifies config
// Option은 config를 수정하는 함수입니다
type Option func(*config)

// Example option / 옵션 예제
func WithTimeout(timeout time.Duration) Option {
    return func(c *config) {
        c.timeout = timeout
    }
}

// Usage / 사용법
client := httputil.NewClient(
    WithTimeout(30*time.Second),
    WithRetry(3),
    WithBearerToken("token"),
)
```

**Benefits / 이점:**
- **Backward Compatible**: Adding options doesn't break existing code
- **Self-Documenting**: Option names are clear
- **Composable**: Options can be combined easily
- **역방향 호환**: 옵션 추가가 기존 코드를 깨뜨리지 않음
- **자체 문서화**: 옵션 이름이 명확함
- **조합 가능**: 옵션을 쉽게 결합할 수 있음

### 5.2 Singleton Pattern / 싱글톤 패턴

**Purpose / 목적:** Provide default client for convenience

**편의를 위한 기본 클라이언트 제공**

**Implementation / 구현:**
```go
// Package-level default client / 패키지 레벨 기본 클라이언트
var defaultClient = NewClient()

// Package-level functions use defaultClient
// 패키지 레벨 함수는 defaultClient를 사용합니다
func Get(url string, result interface{}, opts ...Option) error {
    return defaultClient.Get(url, result, opts...)
}

// Allow users to replace default / 사용자가 기본값을 교체할 수 있도록 허용
func SetDefaultClient(client *Client) {
    defaultClient = client
}
```

**Benefits / 이점:**
- **Quick Start**: No client creation needed
- **Global Configuration**: Configure once, use everywhere
- **빠른 시작**: 클라이언트 생성 불필요
- **전역 설정**: 한 번 설정, 어디서나 사용

### 5.3 Builder Pattern (Implicit) / 빌더 패턴 (암시적)

**Purpose / 목적:** Construct complex Client instances

**복잡한 Client 인스턴스 구축**

**Implementation / 구현:**
```go
func NewClient(opts ...Option) *Client {
    // Start with defaults / 기본값으로 시작
    cfg := defaultConfig()

    // Apply options / 옵션 적용
    cfg.apply(opts)

    // Build http.Client / http.Client 구축
    client := &http.Client{
        Timeout: cfg.timeout,
    }

    // Configure redirect policy / 리디렉션 정책 설정
    if !cfg.followRedirects {
        client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        }
    }

    // ... more configuration

    return &Client{
        client: client,
        config: cfg,
    }
}
```

### 5.4 Facade Pattern / 파사드 패턴

**Purpose / 목적:** Hide complexity of standard library

**표준 라이브러리의 복잡성 숨기기**

**Before (Standard Library) / 이전 (표준 라이브러리):**
```go
// 30+ lines of code
client := &http.Client{Timeout: 30 * time.Second}
req, _ := http.NewRequest("GET", url, nil)
req.Header.Set("Authorization", "Bearer token")
resp, _ := client.Do(req)
defer resp.Body.Close()
// ... error handling, JSON decoding, etc.
```

**After (httputil) / 이후 (httputil):**
```go
// 2 lines
var result MyStruct
err := httputil.Get(url, &result, httputil.WithBearerToken("token"))
```

### 5.5 Decorator Pattern / 데코레이터 패턴

**Purpose / 목적:** Add retry and error handling transparently

**재시도 및 에러 처리를 투명하게 추가**

**Implementation / 구현:**
```go
func (c *Client) doRequest(...) error {
    // Wrap standard http.Client.Do with:
    // 표준 http.Client.Do를 다음으로 래핑:
    // - Retry logic / 재시도 로직
    // - Error enrichment / 에러 강화
    // - JSON encoding/decoding / JSON 인코딩/디코딩

    for attempt := 0; attempt <= maxRetries; attempt++ {
        resp, err := c.client.Do(req)
        // ... retry and error wrapping
    }
}
```

---

## 6. Adding New Features / 새 기능 추가

### 6.1 Adding a New Option / 새 옵션 추가

**Steps / 단계:**

1. **Add field to config struct / config 구조체에 필드 추가**
   ```go
   type config struct {
       // ... existing fields
       myNewOption string
   }
   ```

2. **Update defaultConfig / defaultConfig 업데이트**
   ```go
   func defaultConfig() *config {
       return &config{
           // ... existing defaults
           myNewOption: "default-value",
       }
   }
   ```

3. **Create option function / 옵션 함수 생성**
   ```go
   // WithMyNewOption sets my new option.
   // WithMyNewOption은 새 옵션을 설정합니다.
   func WithMyNewOption(value string) Option {
       return func(c *config) {
           c.myNewOption = value
       }
   }
   ```

4. **Use in doRequest / doRequest에서 사용**
   ```go
   func (c *Client) doRequest(...) error {
       cfg := *c.config
       cfg.apply(opts)

       // Use cfg.myNewOption / cfg.myNewOption 사용
       // ...
   }
   ```

5. **Add tests / 테스트 추가**
   ```go
   func TestMyNewOption(t *testing.T) {
       client := NewClient(WithMyNewOption("test-value"))
       if client.config.myNewOption != "test-value" {
           t.Errorf("Expected 'test-value', got '%s'", client.config.myNewOption)
       }
   }
   ```

6. **Update documentation / 문서 업데이트**
   - README.md
   - USER_MANUAL.md
   - DEVELOPER_GUIDE.md (this file)

### 6.2 Adding a New HTTP Method / 새 HTTP 메서드 추가

**Example: Adding HEAD method / 예제: HEAD 메서드 추가**

1. **Add to Client / Client에 추가**
   ```go
   // Head performs a HEAD request.
   // Head는 HEAD 요청을 수행합니다.
   func (c *Client) Head(path string, opts ...Option) (*http.Response, error) {
       return c.HeadContext(context.Background(), path, opts...)
   }

   // HeadContext performs a HEAD request with context.
   // HeadContext는 context와 함께 HEAD 요청을 수행합니다.
   func (c *Client) HeadContext(ctx context.Context, path string, opts ...Option) (*http.Response, error) {
       // HEAD requests don't have body or result
       // HEAD 요청은 본문이나 결과가 없습니다
       return c.doRequestRaw(ctx, http.MethodHead, path, nil, opts...)
   }
   ```

2. **Add to simple.go / simple.go에 추가**
   ```go
   // Head performs a HEAD request using the default client.
   // Head는 기본 클라이언트를 사용하여 HEAD 요청을 수행합니다.
   func Head(url string, opts ...Option) (*http.Response, error) {
       return defaultClient.Head(url, opts...)
   }

   // HeadContext performs a HEAD request with context using the default client.
   // HeadContext는 기본 클라이언트를 사용하여 context와 함께 HEAD 요청을 수행합니다.
   func HeadContext(ctx context.Context, url string, opts ...Option) (*http.Response, error) {
       return defaultClient.HeadContext(ctx, url, opts...)
   }
   ```

3. **Add tests / 테스트 추가**
   ```go
   func TestHead(t *testing.T) {
       server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           if r.Method != http.MethodHead {
               t.Errorf("Expected HEAD, got %s", r.Method)
           }
           w.WriteHeader(http.StatusOK)
       }))
       defer server.Close()

       resp, err := httputil.Head(server.URL)
       if err != nil {
           t.Fatal(err)
       }
       if resp.StatusCode != http.StatusOK {
           t.Errorf("Expected 200, got %d", resp.StatusCode)
       }
   }
   ```

### 6.3 Adding a New Error Type / 새 에러 타입 추가

**Example: Adding RateLimitError / 예제: RateLimitError 추가**

1. **Define error type / 에러 타입 정의**
   ```go
   // RateLimitError represents a rate limit error.
   // RateLimitError는 속도 제한 에러를 나타냅니다.
   type RateLimitError struct {
       URL          string
       Method       string
       RetryAfter   time.Duration // How long to wait / 대기 시간
       Limit        int           // Rate limit / 속도 제한
       Remaining    int           // Remaining requests / 남은 요청
   }

   func (e *RateLimitError) Error() string {
       return fmt.Sprintf("rate limit exceeded (URL: %s %s), retry after %v",
           e.Method, e.URL, e.RetryAfter)
   }
   ```

2. **Add helper functions / 헬퍼 함수 추가**
   ```go
   // IsRateLimitError checks if an error is a RateLimitError.
   // IsRateLimitError는 에러가 RateLimitError인지 확인합니다.
   func IsRateLimitError(err error) bool {
       _, ok := err.(*RateLimitError)
       return ok
   }
   ```

3. **Use in doRequest / doRequest에서 사용**
   ```go
   func (c *Client) doRequest(...) error {
       // ...
       if resp.StatusCode == 429 {
           retryAfter := parseRetryAfter(resp.Header.Get("Retry-After"))
           return &RateLimitError{
               URL:        fullURL,
               Method:     method,
               RetryAfter: retryAfter,
           }
       }
       // ...
   }
   ```

4. **Add tests and documentation / 테스트 및 문서 추가**

---

## 7. Testing Guide / 테스트 가이드

### 7.1 Test Structure / 테스트 구조

**Current Tests (httputil_test.go) / 현재 테스트:**
- `TestVersion`: Version loading
- `TestNewClient`: Client creation
- `TestNewClientWithOptions`: Client with options
- `TestHTTPError`: HTTPError functionality
- `TestRetryError`: RetryError functionality
- `TestTimeoutError`: TimeoutError functionality
- `TestDefaultConfig`: Default configuration values

### 7.2 Writing Unit Tests / 단위 테스트 작성

**Template / 템플릿:**
```go
func TestFeatureName(t *testing.T) {
    // Setup / 설정
    // ... create test data

    // Execute / 실행
    // ... call function

    // Assert / 검증
    // ... verify results
}
```

**Example / 예제:**
```go
func TestWithTimeout(t *testing.T) {
    timeout := 10 * time.Second
    client := NewClient(WithTimeout(timeout))

    if client.config.timeout != timeout {
        t.Errorf("Expected timeout %v, got %v",
            timeout, client.config.timeout)
    }
}
```

### 7.3 Integration Tests / 통합 테스트

**Using httptest / httptest 사용:**
```go
func TestGetRequest(t *testing.T) {
    // Create test server / 테스트 서버 생성
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verify request / 요청 검증
        if r.Method != "GET" {
            t.Errorf("Expected GET, got %s", r.Method)
        }

        // Send response / 응답 전송
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status": "ok",
        })
    }))
    defer server.Close()

    // Test httputil / httputil 테스트
    var result map[string]string
    err := httputil.Get(server.URL, &result)

    if err != nil {
        t.Fatal(err)
    }

    if result["status"] != "ok" {
        t.Errorf("Expected 'ok', got '%s'", result["status"])
    }
}
```

### 7.4 Testing Retry Logic / 재시도 로직 테스트

```go
func TestRetryOnNetworkError(t *testing.T) {
    attempts := 0

    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        attempts++
        if attempts < 3 {
            // Fail first 2 attempts / 처음 2번 실패
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        // Succeed on 3rd attempt / 3번째 시도에서 성공
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    }))
    defer server.Close()

    var result map[string]string
    err := httputil.Get(server.URL, &result,
        httputil.WithRetry(3),
        httputil.WithRetryBackoff(1*time.Millisecond, 10*time.Millisecond))

    if err != nil {
        t.Fatal(err)
    }

    if attempts != 3 {
        t.Errorf("Expected 3 attempts, got %d", attempts)
    }
}
```

### 7.5 Running Tests / 테스트 실행

```bash
# Run all tests / 모든 테스트 실행
go test ./httputil -v

# Run specific test / 특정 테스트 실행
go test ./httputil -v -run TestVersion

# Run with coverage / 커버리지와 함께 실행
go test ./httputil -cover

# Generate coverage report / 커버리지 리포트 생성
go test ./httputil -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks / 벤치마크 실행
go test ./httputil -bench=.
```

---

## 8. Performance / 성능

### 8.1 Connection Reuse / 연결 재사용

**Best Practice / 모범 사례:**
```go
// Good: Reuse client / 좋음: 클라이언트 재사용
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"))

for i := 0; i < 1000; i++ {
    client.Get(fmt.Sprintf("/users/%d", i), &user)
}

// Bad: Create new client each time / 나쁨: 매번 새 클라이언트 생성
for i := 0; i < 1000; i++ {
    httputil.Get(fmt.Sprintf("https://api.example.com/users/%d", i), &user)
}
```

**Why / 이유:**
- Standard `http.Client` maintains connection pool
- Reusing client reuses TCP connections
- Reduces latency and improves throughput
- 표준 http.Client는 연결 풀을 유지함
- 클라이언트 재사용은 TCP 연결을 재사용함
- 지연 시간 감소 및 처리량 향상

### 8.2 Memory Allocations / 메모리 할당

**Optimization Points / 최적화 지점:**

1. **Avoid Repeated JSON Marshal/Unmarshal / 반복적인 JSON Marshal/Unmarshal 방지**
   ```go
   // Bad: Marshal same payload multiple times / 나쁨: 동일한 페이로드를 여러 번 마샬링
   for i := 0; i < 100; i++ {
       payload := MyPayload{Data: "same"}
       httputil.Post(url, payload, &result)
   }

   // Good: Marshal once / 좋음: 한 번만 마샬링
   payload := MyPayload{Data: "same"}
   for i := 0; i < 100; i++ {
       httputil.Post(url, payload, &result)
   }
   ```

2. **Reuse Buffers / 버퍼 재사용**
   ```go
   // Internal: bytes.NewReader reuses buffer
   // 내부: bytes.NewReader가 버퍼를 재사용
   bodyReader = bytes.NewReader(jsonData)
   ```

### 8.3 Benchmarking / 벤치마킹

**Example Benchmark / 벤치마크 예제:**
```go
func BenchmarkGet(b *testing.B) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    }))
    defer server.Close()

    client := httputil.NewClient()
    var result map[string]string

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        client.Get(server.URL, &result)
    }
}
```

**Run Benchmark / 벤치마크 실행:**
```bash
go test ./httputil -bench=BenchmarkGet -benchmem
```

---

## 9. Contributing Guidelines / 기여 가이드라인

### 9.1 Development Setup / 개발 설정

**Prerequisites / 전제 조건:**
- Go 1.18 or later
- Git
- go 1.18 이상
- Git

**Clone Repository / 저장소 클론:**
```bash
git clone https://github.com/arkd0ng/go-utils.git
cd go-utils/httputil
```

**Install Dependencies / 의존성 설치:**
```bash
go mod download
```

### 9.2 Making Changes / 변경 사항 만들기

**Workflow / 워크플로우:**

1. **Create Feature Branch / 기능 브랜치 생성**
   ```bash
   git checkout -b feature/my-new-feature
   ```

2. **Make Changes / 변경 사항 작성**
   - Write code / 코드 작성
   - Add tests / 테스트 추가
   - Update documentation / 문서 업데이트

3. **Run Tests / 테스트 실행**
   ```bash
   go test ./httputil -v
   go test ./httputil -cover
   ```

4. **Format Code / 코드 포맷**
   ```bash
   go fmt ./httputil
   ```

5. **Commit Changes / 변경 사항 커밋**
   ```bash
   git add .
   git commit -m "Feat: Add new feature"
   ```

6. **Push and Create PR / 푸시 및 PR 생성**
   ```bash
   git push origin feature/my-new-feature
   ```

### 9.3 Contribution Checklist / 기여 체크리스트

**Before Submitting PR / PR 제출 전:**
- [ ] All tests pass / 모든 테스트 통과
- [ ] Code is formatted / 코드 포맷팅됨
- [ ] Documentation updated / 문서 업데이트됨
- [ ] Bilingual comments (EN/KR) / 이중 언어 주석 (영문/한글)
- [ ] Commit message follows format / 커밋 메시지가 형식을 따름
- [ ] No external dependencies added / 외부 의존성 추가 안 됨

**Commit Message Format / 커밋 메시지 형식:**
```
Type: Brief description

Types:
- Feat: New feature
- Fix: Bug fix
- Docs: Documentation changes
- Refactor: Code refactoring
- Test: Test additions/changes
- Chore: Build, config changes
```

---

## 10. Code Style / 코드 스타일

### 10.1 Naming Conventions / 명명 규칙

**Exported Functions / 내보낸 함수:**
- Use PascalCase: `Get`, `Post`, `NewClient`
- Clear, descriptive names
- PascalCase 사용
- 명확하고 설명적인 이름

**Unexported Functions / 내보내지 않은 함수:**
- Use camelCase: `doRequest`, `calculateBackoff`
- camelCase 사용

**Variables / 변수:**
- Short names for local scope: `cfg`, `req`, `resp`, `err`
- Descriptive names for wider scope: `defaultClient`, `maxRetries`
- 로컬 스코프에는 짧은 이름
- 넓은 스코프에는 설명적인 이름

### 10.2 Comment Style / 주석 스타일

**Bilingual Comments / 이중 언어 주석:**
```go
// Function comment in English.
// 함수 주석을 한글로.
func MyFunction() {
    // Implementation comment in English
    // 구현 주석을 한글로
}
```

**GoDoc Format / GoDoc 형식:**
```go
// NewClient creates a new HTTP client with the given options.
// NewClient는 주어진 옵션으로 새로운 HTTP 클라이언트를 생성합니다.
//
// Example:
//
//	client := httputil.NewClient(
//	    httputil.WithTimeout(30*time.Second),
//	    httputil.WithRetry(3),
//	)
func NewClient(opts ...Option) *Client {
    // ...
}
```

### 10.3 Error Handling / 에러 처리

**Pattern / 패턴:**
```go
// Check error immediately / 즉시 에러 확인
result, err := someFunction()
if err != nil {
    return fmt.Errorf("context: %w", err)
}

// Use result / 결과 사용
doSomething(result)
```

**Wrap Errors / 에러 래핑:**
```go
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}
```

### 10.4 Code Organization / 코드 구성

**File Structure / 파일 구조:**
```go
package httputil

// 1. Imports / 임포트
import (
    "context"
    "net/http"
)

// 2. Constants / 상수
const (
    DefaultTimeout = 30 * time.Second
)

// 3. Types / 타입
type Client struct {
    client *http.Client
    config *config
}

// 4. Constructors / 생성자
func NewClient(opts ...Option) *Client {
    // ...
}

// 5. Methods / 메서드
func (c *Client) Get(...) error {
    // ...
}

// 6. Helper Functions / 헬퍼 함수
func calculateBackoff(...) time.Duration {
    // ...
}
```

### 10.5 Best Practices / 모범 사례

1. **Keep Functions Small / 함수를 작게 유지**
   - One function, one responsibility
   - Max 50-100 lines per function
   - 하나의 함수, 하나의 책임
   - 함수당 최대 50-100줄

2. **Use Early Returns / 조기 반환 사용**
   ```go
   // Good / 좋음
   if err != nil {
       return err
   }
   // continue with success case

   // Bad / 나쁨
   if err == nil {
       // long success code
   } else {
       return err
   }
   ```

3. **Avoid Else After Return / 반환 후 else 방지**
   ```go
   // Good / 좋음
   if condition {
       return value1
   }
   return value2

   // Bad / 나쁨
   if condition {
       return value1
   } else {
       return value2
   }
   ```

4. **Use Table-Driven Tests / 테이블 기반 테스트 사용**
   ```go
   func TestSomething(t *testing.T) {
       tests := []struct {
           name     string
           input    string
           expected string
       }{
           {"case1", "input1", "expected1"},
           {"case2", "input2", "expected2"},
       }

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               result := doSomething(tt.input)
               if result != tt.expected {
                   t.Errorf("Expected %s, got %s", tt.expected, result)
               }
           })
       }
   }
   ```

---

## Additional Resources / 추가 자료

- **User Manual**: [USER_MANUAL.md](USER_MANUAL.md)
- **Design Plan**: [DESIGN_PLAN.md](DESIGN_PLAN.md)
- **Work Plan**: [WORK_PLAN.md](WORK_PLAN.md)
- **Package README**: [../../httputil/README.md](../../httputil/README.md)
- **Changelog**: [../CHANGELOG/CHANGELOG-v1.10.md](../CHANGELOG/CHANGELOG-v1.10.md)

---

**Last Updated / 최종 업데이트**: 2025-10-15
**Version / 버전**: v1.10.001
**Author / 작성자**: arkd0ng
