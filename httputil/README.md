# httputil Package

Extremely simple HTTP client utilities that reduce 30+ lines to 2-3 lines.

극도로 간단한 HTTP 클라이언트 유틸리티로 30줄 이상의 코드를 2-3줄로 줄입니다.

**Version / 버전**: v1.10.004
**Package Path / 패키지 경로**: `github.com/arkd0ng/go-utils/httputil`

---

## Table of Contents / 목차

- [Features / 주요 기능](#features--주요-기능)
- [Installation / 설치](#installation--설치)
- [Quick Start / 빠른 시작](#quick-start--빠른-시작)
- [API Reference / API 참조](#api-reference--api-참조)
  - [Simple API / 간단한 API](#simple-api--간단한-api)
  - [Client API](#client-api)
  - [Configuration Options / 설정 옵션](#configuration-options--설정-옵션)
  - [Error Types / 에러 타입](#error-types--에러-타입)
- [Usage Examples / 사용 예제](#usage-examples--사용-예제)
- [Best Practices / 모범 사례](#best-practices--모범-사례)
- [Documentation / 문서](#documentation--문서)

---

## Features / 주요 기능

### Extreme Simplicity / 극도의 간결함
- **30+ lines → 2-3 lines** of code reduction
- **30줄 이상 → 2-3줄**로 코드 감소
- Zero boilerplate for common HTTP operations
- 일반적인 HTTP 작업에 보일러플레이트 코드 불필요

### Core Features / 핵심 기능

✅ **RESTful HTTP Methods** / **RESTful HTTP 메서드**
- GET, POST, PUT, PATCH, DELETE
- Context variants for all methods (cancellation & timeout support)
- 모든 메서드의 Context 버전 (취소 및 타임아웃 지원)

✅ **Automatic JSON Handling** / **자동 JSON 처리**
- Automatic request body encoding
- Automatic response body decoding
- 요청 본문 자동 인코딩
- 응답 본문 자동 디코딩

✅ **Smart Retry Logic** / **스마트 재시도 로직**
- Configurable retry attempts (default: 3)
- Exponential backoff with jitter
- Automatic retry on network errors and 5xx responses
- 설정 가능한 재시도 횟수 (기본값: 3)
- 지터가 있는 지수 백오프
- 네트워크 오류 및 5xx 응답 시 자동 재시도

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
- URL utilities (JoinURL, AddQueryParams, GetDomain, etc.)
- URL 및 폼 구축을 위한 Fluent API
- 조건부 매개변수 (ParamIf, AddIf)
- URL 유틸리티 (JoinURL, AddQueryParams, GetDomain 등)

✅ **Rich Configuration** / **풍부한 설정**
- Functional options pattern for flexible configuration
- 12 built-in options (timeout, headers, auth, retry, etc.)
- 함수형 옵션 패턴으로 유연한 설정
- 12개 내장 옵션 (타임아웃, 헤더, 인증, 재시도 등)

✅ **Rich Error Types** / **풍부한 에러 타입**
- HTTPError (status code, body, URL)
- RetryError (failed attempts tracking)
- TimeoutError (timeout detection)
- HTTPError (상태 코드, 본문, URL)
- RetryError (실패한 시도 추적)
- TimeoutError (타임아웃 감지)

✅ **Zero External Dependencies** / **제로 외부 의존성**
- Standard library only (net/http, encoding/json, mime/multipart)
- 표준 라이브러리만 사용 (net/http, encoding/json, mime/multipart)

---

## Installation / 설치

### Prerequisites / 전제 조건

- Go 1.18 or later (for generics support)
- Go 1.18 이상 (제네릭 지원)

### Install Package / 패키지 설치

```bash
go get github.com/arkd0ng/go-utils/httputil
```

### Import / 임포트

```go
import "github.com/arkd0ng/go-utils/httputil"
```

---

## Quick Start / 빠른 시작

### 1. Simple GET Request / 간단한 GET 요청

**Before (30+ lines):**
```go
// Traditional approach / 전통적인 방식
client := &http.Client{Timeout: 30 * time.Second}
req, err := http.NewRequest("GET", "https://api.example.com/users", nil)
if err != nil {
    log.Fatal(err)
}
req.Header.Set("Authorization", "Bearer token123")
req.Header.Set("Content-Type", "application/json")

resp, err := client.Do(req)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

if resp.StatusCode != 200 {
    body, _ := io.ReadAll(resp.Body)
    log.Fatalf("HTTP %d: %s", resp.StatusCode, body)
}

var users []User
if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
    log.Fatal(err)
}
// 20+ more lines for retry logic, error handling...
```

**After (2 lines):**
```go
// httputil approach / httputil 방식
var users []User
err := httputil.Get("https://api.example.com/users", &users,
    httputil.WithBearerToken("token123"))
```

### 2. POST Request with Authentication / 인증과 함께 POST 요청

```go
// Create payload / 페이로드 생성
payload := User{
    Name:  "John Doe",
    Email: "john@example.com",
}

// Send POST request / POST 요청 전송
var result Response
err := httputil.Post("https://api.example.com/users", payload, &result,
    httputil.WithBearerToken("your-token"),
    httputil.WithTimeout(30*time.Second))
if err != nil {
    log.Fatal(err)
}
```

### 3. Client with Base URL / Base URL을 가진 클라이언트

```go
// Create configured client / 설정된 클라이언트 생성
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithBearerToken("your-token"),
    httputil.WithRetry(5),
    httputil.WithTimeout(60*time.Second),
)

// Use client for multiple requests / 여러 요청에 클라이언트 사용
var users []User
err := client.Get("/users", &users)

var user User
err = client.Get("/users/123", &user)

payload := User{Name: "Jane"}
var created User
err = client.Post("/users", payload, &created)
```

### 4. Response Helpers / 응답 헬퍼

```go
// Get raw response with helper methods / 헬퍼 메서드와 함께 원시 응답 가져오기
resp, err := httputil.DoRaw("GET", "https://api.example.com/users", nil,
    httputil.WithBearerToken("token123"))
if err != nil {
    log.Fatal(err)
}

// Check status / 상태 확인
if resp.IsSuccess() {
    log.Println("Request successful")
}

// Access body / 본문 접근
bodyString := resp.String()
bodyBytes := resp.Body()

// Decode JSON / JSON 디코딩
var users []User
err = resp.JSON(&users)
```

### 5. File Download with Progress / 진행 상황과 함께 파일 다운로드

```go
// Download file with progress tracking / 진행 상황 추적과 함께 파일 다운로드
err := httputil.DownloadFile(
    "https://example.com/large-file.zip",
    "./downloads/file.zip",
    httputil.WithProgress(func(bytesRead, totalBytes int64) {
        progress := float64(bytesRead) / float64(totalBytes) * 100
        fmt.Printf("\rDownloading: %.2f%%", progress)
    }),
)
if err != nil {
    log.Fatal(err)
}
```

### 6. File Upload / 파일 업로드

```go
// Upload single file / 단일 파일 업로드
var result map[string]interface{}
err := httputil.UploadFile(
    "https://api.example.com/upload",
    "file",
    "./document.pdf",
    &result,
    httputil.WithBearerToken("token123"),
)
if err != nil {
    log.Fatal(err)
}

// Upload multiple files / 여러 파일 업로드
err = httputil.UploadFiles(
    "https://api.example.com/upload-multiple",
    map[string]string{
        "file1": "./image1.jpg",
        "file2": "./image2.jpg",
    },
    &result,
    httputil.WithBearerToken("token123"),
)
```

### 7. URL Builder / URL 빌더

```go
// Build URL with fluent API / Fluent API로 URL 구축
url := httputil.NewURL("https://api.example.com").
    Path("users", "search").
    Param("q", "golang").
    Param("page", "1").
    ParamIf(includeInactive, "status", "inactive").
    Build()
// Result: https://api.example.com/users/search?q=golang&page=1

// Join URL paths / URL 경로 결합
url = httputil.JoinURL("https://api.example.com", "v1", "users", "123")
// Result: https://api.example.com/v1/users/123
```

### 8. Form Builder / Form 빌더

```go
// Build form data with fluent API / Fluent API로 폼 데이터 구축
form := httputil.NewForm().
    Set("username", "john").
    Set("email", "john@example.com").
    AddIf(hasPromoCode, "promo_code", "SAVE20").
    AddMultiple("tags", "go", "http", "api")

// Post form data / 폼 데이터 전송
var result map[string]interface{}
err := httputil.PostForm(
    "https://api.example.com/submit",
    form.Map(),
    &result,
)
```

---

## API Reference / API 참조

### Simple API / 간단한 API

Package-level convenience functions using a default client.

기본 클라이언트를 사용하는 패키지 레벨 편의 함수들.

#### GET Request / GET 요청

```go
func Get(url string, result interface{}, opts ...Option) error
func GetContext(ctx context.Context, url string, result interface{}, opts ...Option) error
```

**Example / 예제:**
```go
var data MyStruct
err := httputil.Get("https://api.example.com/data", &data)
```

#### POST Request / POST 요청

```go
func Post(url string, body, result interface{}, opts ...Option) error
func PostContext(ctx context.Context, url string, body, result interface{}, opts ...Option) error
```

**Example / 예제:**
```go
payload := MyPayload{Name: "test"}
var response MyResponse
err := httputil.Post("https://api.example.com/create", payload, &response)
```

#### PUT Request / PUT 요청

```go
func Put(url string, body, result interface{}, opts ...Option) error
func PutContext(ctx context.Context, url string, body, result interface{}, opts ...Option) error
```

**Example / 예제:**
```go
update := MyUpdate{ID: 1, Name: "updated"}
var response MyResponse
err := httputil.Put("https://api.example.com/update/1", update, &response)
```

#### PATCH Request / PATCH 요청

```go
func Patch(url string, body, result interface{}, opts ...Option) error
func PatchContext(ctx context.Context, url string, body, result interface{}, opts ...Option) error
```

**Example / 예제:**
```go
patch := map[string]interface{}{"name": "patched"}
var response MyResponse
err := httputil.Patch("https://api.example.com/update/1", patch, &response)
```

#### DELETE Request / DELETE 요청

```go
func Delete(url string, result interface{}, opts ...Option) error
func DeleteContext(ctx context.Context, url string, result interface{}, opts ...Option) error
```

**Example / 예제:**
```go
var response MyResponse
err := httputil.Delete("https://api.example.com/delete/1", &response)
```

#### Set Default Client / 기본 클라이언트 설정

```go
func SetDefaultClient(client *Client)
```

Configure the default client used by package-level functions.

패키지 레벨 함수에서 사용되는 기본 클라이언트를 설정합니다.

**Example / 예제:**
```go
httputil.SetDefaultClient(httputil.NewClient(
    httputil.WithTimeout(60*time.Second),
    httputil.WithRetry(5),
))
```

#### Raw Response / 원시 응답

```go
func DoRaw(method, url string, body interface{}, opts ...Option) (*Response, error)
func DoRawContext(ctx context.Context, method, url string, body interface{}, opts ...Option) (*Response, error)
```

Get raw Response with helper methods instead of auto-decoding JSON.

JSON 자동 디코딩 대신 헬퍼 메서드가 있는 원시 Response를 가져옵니다.

**Example / 예제:**
```go
resp, err := httputil.DoRaw("GET", "https://api.example.com/data", nil)
if resp.IsSuccess() {
    bodyString := resp.String()
}
```

#### File Operations / 파일 작업

```go
func DownloadFile(url, filepath string, opts ...Option) error
func DownloadFileContext(ctx context.Context, url, filepath string, progress ProgressFunc, opts ...Option) error
func Download(url string, opts ...Option) ([]byte, error)
func DownloadContext(ctx context.Context, url string, opts ...Option) ([]byte, error)
```

Download files from URL with optional progress tracking.

선택적 진행 상황 추적과 함께 URL에서 파일을 다운로드합니다.

**Example / 예제:**
```go
// Download to file / 파일로 다운로드
err := httputil.DownloadFile("https://example.com/file.zip", "./file.zip")

// Download to memory / 메모리로 다운로드
data, err := httputil.Download("https://example.com/data.json")

// With progress callback / 진행 상황 콜백과 함께
ctx := context.Background()
err = httputil.DownloadFileContext(ctx, url, filepath,
    func(read, total int64) {
        fmt.Printf("Progress: %.2f%%\n", float64(read)/float64(total)*100)
    })
```

```go
func UploadFile(url, fieldName, filepath string, result interface{}, opts ...Option) error
func UploadFileContext(ctx context.Context, url, fieldName, filepath string, result interface{}, opts ...Option) error
func UploadFiles(url string, files map[string]string, result interface{}, opts ...Option) error
func UploadFilesContext(ctx context.Context, url string, files map[string]string, result interface{}, opts ...Option) error
```

Upload files using multipart form data.

multipart form data를 사용하여 파일을 업로드합니다.

**Example / 예제:**
```go
// Upload single file / 단일 파일 업로드
var result map[string]interface{}
err := httputil.UploadFile(url, "document", "./file.pdf", &result)

// Upload multiple files / 여러 파일 업로드
err = httputil.UploadFiles(url, map[string]string{
    "file1": "./image1.jpg",
    "file2": "./image2.jpg",
}, &result)
```

#### Form Operations / Form 작업

```go
func PostForm(url string, data map[string]string, result interface{}, opts ...Option) error
func PostFormContext(ctx context.Context, url string, data map[string]string, result interface{}, opts ...Option) error
```

Post form data with `application/x-www-form-urlencoded` encoding.

`application/x-www-form-urlencoded` 인코딩으로 폼 데이터를 전송합니다.

**Example / 예제:**
```go
var result map[string]interface{}
err := httputil.PostForm(url, map[string]string{
    "username": "john",
    "email": "john@example.com",
}, &result)
```

---

### Client API

Create and configure custom HTTP clients.

사용자 정의 HTTP 클라이언트 생성 및 설정.

#### Create Client / 클라이언트 생성

```go
func NewClient(opts ...Option) *Client
```

**Example / 예제:**
```go
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithBearerToken("token123"),
    httputil.WithRetry(3),
)
```

#### Client Methods / 클라이언트 메서드

All methods available on Client instances:

Client 인스턴스에서 사용 가능한 모든 메서드:

```go
func (c *Client) Get(path string, result interface{}, opts ...Option) error
func (c *Client) GetContext(ctx context.Context, path string, result interface{}, opts ...Option) error

func (c *Client) Post(path string, body, result interface{}, opts ...Option) error
func (c *Client) PostContext(ctx context.Context, path string, body, result interface{}, opts ...Option) error

func (c *Client) Put(path string, body, result interface{}, opts ...Option) error
func (c *Client) PutContext(ctx context.Context, path string, body, result interface{}, opts ...Option) error

func (c *Client) Patch(path string, body, result interface{}, opts ...Option) error
func (c *Client) PatchContext(ctx context.Context, path string, body, result interface{}, opts ...Option) error

func (c *Client) Delete(path string, result interface{}, opts ...Option) error
func (c *Client) DeleteContext(ctx context.Context, path string, result interface{}, opts ...Option) error
```

#### Client File Operations / 클라이언트 파일 작업

```go
func (c *Client) DownloadFile(url, filepath string, opts ...Option) error
func (c *Client) DownloadFileContext(ctx context.Context, url, filepath string, progress ProgressFunc, opts ...Option) error
func (c *Client) Download(url string, opts ...Option) ([]byte, error)
func (c *Client) DownloadContext(ctx context.Context, url string, opts ...Option) ([]byte, error)

func (c *Client) UploadFile(url, fieldName, filepath string, result interface{}, opts ...Option) error
func (c *Client) UploadFileContext(ctx context.Context, url, fieldName, filepath string, result interface{}, opts ...Option) error
func (c *Client) UploadFiles(url string, files map[string]string, result interface{}, opts ...Option) error
func (c *Client) UploadFilesContext(ctx context.Context, url string, files map[string]string, result interface{}, opts ...Option) error
```

#### Client Form Operations / 클라이언트 Form 작업

```go
func (c *Client) PostForm(path string, data map[string]string, result interface{}, opts ...Option) error
func (c *Client) PostFormContext(ctx context.Context, path string, data map[string]string, result interface{}, opts ...Option) error
```

---

### Response API / Response API

Rich response wrapper with helper methods for easier response handling.

더 쉬운 응답 처리를 위한 헬퍼 메서드가 있는 풍부한 응답 래퍼.

```go
type Response struct {
    *http.Response
    // ... cached body
}
```

#### Body Methods / 본문 메서드

```go
func (r *Response) Body() []byte                          // Get response body as bytes
func (r *Response) String() string                        // Get response body as string
func (r *Response) JSON(result interface{}) error         // Decode JSON into result
```

**Example / 예제:**
```go
resp, _ := httputil.DoRaw("GET", url, nil)
bodyBytes := resp.Body()
bodyString := resp.String()

var data MyStruct
resp.JSON(&data)
```

#### Status Check Methods / 상태 확인 메서드

```go
func (r *Response) IsSuccess() bool        // 2xx status codes
func (r *Response) IsError() bool          // 4xx or 5xx status codes
func (r *Response) IsClientError() bool    // 4xx status codes
func (r *Response) IsServerError() bool    // 5xx status codes

// Specific status codes / 특정 상태 코드
func (r *Response) IsOK() bool             // 200 OK
func (r *Response) IsCreated() bool        // 201 Created
func (r *Response) IsAccepted() bool       // 202 Accepted
func (r *Response) IsNoContent() bool      // 204 No Content
func (r *Response) IsMovedPermanently() bool    // 301 Moved Permanently
func (r *Response) IsFound() bool          // 302 Found
func (r *Response) IsBadRequest() bool     // 400 Bad Request
func (r *Response) IsUnauthorized() bool   // 401 Unauthorized
func (r *Response) IsForbidden() bool      // 403 Forbidden
func (r *Response) IsNotFound() bool       // 404 Not Found
func (r *Response) IsInternalServerError() bool  // 500 Internal Server Error
```

**Example / 예제:**
```go
resp, _ := httputil.DoRaw("GET", url, nil)

if resp.IsSuccess() {
    log.Println("Request successful")
} else if resp.IsNotFound() {
    log.Println("Resource not found")
} else if resp.IsServerError() {
    log.Println("Server error")
}
```

#### Header Methods / 헤더 메서드

```go
func (r *Response) Header(key string) string           // Get single header value
func (r *Response) Headers() map[string]string        // Get all headers as map
func (r *Response) ContentType() string               // Get Content-Type header
```

**Example / 예제:**
```go
resp, _ := httputil.DoRaw("GET", url, nil)
contentType := resp.ContentType()
allHeaders := resp.Headers()
customHeader := resp.Header("X-Custom-Header")
```

---

### URL Builder API / URL Builder API

Fluent API for building URLs with parameters.

매개변수와 함께 URL을 구축하기 위한 Fluent API.

```go
type URLBuilder struct { /* ... */ }
```

#### Methods / 메서드

```go
func NewURL(baseURL string) *URLBuilder
func (u *URLBuilder) Path(segments ...string) *URLBuilder
func (u *URLBuilder) Param(key, value string) *URLBuilder
func (u *URLBuilder) Params(params map[string]string) *URLBuilder
func (u *URLBuilder) ParamIf(condition bool, key, value string) *URLBuilder
func (u *URLBuilder) Build() string
```

**Example / 예제:**
```go
// Build complex URL / 복잡한 URL 구축
url := httputil.NewURL("https://api.example.com").
    Path("v1", "users", "search").
    Param("q", "golang").
    Param("page", "1").
    Param("limit", "20").
    ParamIf(includeInactive, "status", "inactive").
    Build()

// Result: https://api.example.com/v1/users/search?q=golang&page=1&limit=20
```

#### URL Utility Functions / URL 유틸리티 함수

```go
func JoinURL(baseURL string, paths ...string) string
func AddQueryParams(urlStr string, params map[string]string) (string, error)
func GetDomain(urlStr string) (string, error)
func GetScheme(urlStr string) (string, error)
func IsAbsoluteURL(urlStr string) bool
func NormalizeURL(urlStr string) string
```

**Example / 예제:**
```go
// Join URL parts / URL 부분 결합
url := httputil.JoinURL("https://api.example.com", "v1", "users", "123")
// Result: https://api.example.com/v1/users/123

// Add query parameters / 쿼리 매개변수 추가
url, _ = httputil.AddQueryParams(url, map[string]string{
    "fields": "name,email",
})

// Get domain / 도메인 가져오기
domain, _ := httputil.GetDomain("https://api.example.com:8080/path")
// Result: api.example.com:8080

// Check if absolute / 절대 URL인지 확인
isAbsolute := httputil.IsAbsoluteURL("https://example.com")  // true
isAbsolute = httputil.IsAbsoluteURL("/relative/path")        // false
```

---

### Form Builder API / Form Builder API

Fluent API for building form data.

폼 데이터를 구축하기 위한 Fluent API.

```go
type FormBuilder struct { /* ... */ }
```

#### Methods / 메서드

```go
func NewForm() *FormBuilder
func (f *FormBuilder) Add(key, value string) *FormBuilder
func (f *FormBuilder) Set(key, value string) *FormBuilder
func (f *FormBuilder) AddMultiple(key string, values ...string) *FormBuilder
func (f *FormBuilder) AddIf(condition bool, key, value string) *FormBuilder
func (f *FormBuilder) Get(key string) string
func (f *FormBuilder) GetAll(key string) []string
func (f *FormBuilder) Has(key string) bool
func (f *FormBuilder) Del(key string) *FormBuilder
func (f *FormBuilder) Clone() *FormBuilder
func (f *FormBuilder) Map() map[string]string
func (f *FormBuilder) Encode() string
```

**Example / 예제:**
```go
// Build form with conditional fields / 조건부 필드가 있는 폼 구축
hasPromo := true
form := httputil.NewForm().
    Set("username", "john").
    Set("email", "john@example.com").
    Set("age", "30").
    AddIf(hasPromo, "promo_code", "SAVE20").
    AddIf(false, "referrer", "none").
    AddMultiple("tags", "go", "http", "api")

// Check if field exists / 필드 존재 확인
if form.Has("promo_code") {
    log.Println("Promo code applied")
}

// Get form data / 폼 데이터 가져오기
formMap := form.Map()
encoded := form.Encode()

// Clone form / 폼 복제
form2 := form.Clone().Set("email", "jane@example.com")
```

#### Form Utility Functions / Form 유틸리티 함수

```go
func ParseForm(data string) (map[string]string, error)
func EncodeForm(data map[string]string) string
```

**Example / 예제:**
```go
// Parse form data / 폼 데이터 파싱
formData, _ := httputil.ParseForm("name=John&city=Seoul&age=30")
// Result: map[string]string{"name": "John", "city": "Seoul", "age": "30"}

// Encode form data / 폼 데이터 인코딩
encoded := httputil.EncodeForm(map[string]string{
    "username": "john",
    "password": "secret",
})
// Result: "password=secret&username=john"
```

---

### Configuration Options / 설정 옵션

#### Request Configuration / 요청 설정

| Option / 옵션 | Type / 타입 | Default / 기본값 | Description / 설명 |
|---------------|-------------|------------------|-------------------|
| `WithTimeout(duration)` | `time.Duration` | 30 seconds / 30초 | Request timeout / 요청 타임아웃 |
| `WithHeaders(map)` | `map[string]string` | Empty / 빈 맵 | Custom headers / 사용자 정의 헤더 |
| `WithHeader(key, value)` | `string, string` | - | Single header / 단일 헤더 |
| `WithQueryParams(map)` | `map[string]string` | Empty / 빈 맵 | Query parameters / 쿼리 매개변수 |
| `WithUserAgent(agent)` | `string` | "go-utils/httputil v{version}" | User-Agent header |
| `WithProgress(callback)` | `ProgressFunc` | `nil` | Progress callback for file operations / 파일 작업을 위한 진행 상황 콜백 |

**Example / 예제:**
```go
err := httputil.Get(url, &result,
    httputil.WithTimeout(10*time.Second),
    httputil.WithHeader("X-Custom-Header", "value"),
    httputil.WithQueryParams(map[string]string{
        "page": "1",
        "limit": "100",
    }),
)
```

#### Authentication / 인증

| Option / 옵션 | Type / 타입 | Description / 설명 |
|---------------|-------------|-------------------|
| `WithBearerToken(token)` | `string` | Bearer token authentication / Bearer 토큰 인증 |
| `WithBasicAuth(user, pass)` | `string, string` | Basic authentication / 기본 인증 |

**Example / 예제:**
```go
// Bearer token / Bearer 토큰
err := httputil.Get(url, &result,
    httputil.WithBearerToken("your-token-here"))

// Basic auth / 기본 인증
err := httputil.Get(url, &result,
    httputil.WithBasicAuth("username", "password"))
```

#### Retry Configuration / 재시도 설정

| Option / 옵션 | Type / 타입 | Default / 기본값 | Description / 설명 |
|---------------|-------------|------------------|-------------------|
| `WithRetry(maxRetries)` | `int` | 3 | Maximum retry attempts / 최대 재시도 횟수 |
| `WithRetryBackoff(min, max)` | `time.Duration, time.Duration` | 100ms, 5s | Backoff time range / 백오프 시간 범위 |

**Example / 예제:**
```go
err := httputil.Get(url, &result,
    httputil.WithRetry(5),
    httputil.WithRetryBackoff(200*time.Millisecond, 10*time.Second))
```

#### Client Configuration / 클라이언트 설정

| Option / 옵션 | Type / 타입 | Default / 기본값 | Description / 설명 |
|---------------|-------------|------------------|-------------------|
| `WithBaseURL(baseURL)` | `string` | Empty / 빈 문자열 | Base URL for all requests / 모든 요청의 기본 URL |
| `WithFollowRedirects(follow)` | `bool` | `true` | Follow HTTP redirects / HTTP 리디렉션 따르기 |
| `WithMaxRedirects(max)` | `int` | 10 | Maximum redirects / 최대 리디렉션 수 |

**Example / 예제:**
```go
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com/v1"),
    httputil.WithFollowRedirects(false),
)
```

#### Cookie Management / 쿠키 관리

| Option / 옵션 | Type / 타입 | Description / 설명 |
|---------------|-------------|-------------------|
| `WithCookies()` | - | Enable in-memory cookie jar / 메모리 내 쿠키 저장소 활성화 |
| `WithPersistentCookies(filePath)` | `string` | Enable persistent cookie jar with file storage / 파일 저장소를 사용한 지속성 쿠키 저장소 활성화 |
| `WithCookieJar(jar)` | `http.CookieJar` | Use custom cookie jar / 사용자 정의 쿠키 저장소 사용 |

**Example / 예제:**
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

// Cookie operations / 쿠키 작업
u, _ := url.Parse("https://api.example.com")
client.SetCookie(u, &http.Cookie{Name: "session", Value: "abc123"})
cookies := client.GetCookies(u)
client.SaveCookies() // Save to file (persistent jar only) / 파일에 저장 (지속성 저장소만)
```

---

### Error Types / 에러 타입

#### HTTPError

HTTP error with status code and body.

상태 코드와 본문을 포함한 HTTP 에러.

```go
type HTTPError struct {
    StatusCode int
    Status     string
    Body       string
    URL        string
    Method     string
}
```

**Check and Handle / 확인 및 처리:**
```go
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsHTTPError(err) {
        statusCode := httputil.GetStatusCode(err)
        if statusCode == 404 {
            // Handle not found / 찾을 수 없음 처리
        }
    }
}
```

#### RetryError

Error after all retry attempts failed.

모든 재시도 시도 실패 후 에러.

```go
type RetryError struct {
    Attempts int
    LastErr  error
    URL      string
    Method   string
}
```

**Check and Handle / 확인 및 처리:**
```go
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsRetryError(err) {
        // All retries failed / 모든 재시도 실패
        retryErr := err.(*httputil.RetryError)
        log.Printf("Failed after %d attempts", retryErr.Attempts)
    }
}
```

#### TimeoutError

Request timeout error.

요청 타임아웃 에러.

```go
type TimeoutError struct {
    URL    string
    Method string
}
```

**Check and Handle / 확인 및 처리:**
```go
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsTimeoutError(err) {
        // Handle timeout / 타임아웃 처리
        log.Println("Request timed out")
    }
}
```

---

## Usage Examples / 사용 예제

### Example 1: REST API Client / REST API 클라이언트

```go
package main

import (
    "log"
    "time"

    "github.com/arkd0ng/go-utils/httputil"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    // Create API client / API 클라이언트 생성
    client := httputil.NewClient(
        httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
        httputil.WithTimeout(30*time.Second),
        httputil.WithRetry(3),
    )

    // GET: List users / GET: 사용자 목록
    var users []User
    if err := client.Get("/users", &users); err != nil {
        log.Fatal(err)
    }
    log.Printf("Found %d users", len(users))

    // GET: Single user / GET: 단일 사용자
    var user User
    if err := client.Get("/users/1", &user); err != nil {
        log.Fatal(err)
    }
    log.Printf("User: %+v", user)

    // POST: Create user / POST: 사용자 생성
    newUser := User{Name: "John Doe", Email: "john@example.com"}
    var created User
    if err := client.Post("/users", newUser, &created); err != nil {
        log.Fatal(err)
    }
    log.Printf("Created: %+v", created)

    // PUT: Update user / PUT: 사용자 업데이트
    updated := User{ID: 1, Name: "Jane Doe", Email: "jane@example.com"}
    var result User
    if err := client.Put("/users/1", updated, &result); err != nil {
        log.Fatal(err)
    }
    log.Printf("Updated: %+v", result)

    // DELETE: Delete user / DELETE: 사용자 삭제
    if err := client.Delete("/users/1", nil); err != nil {
        log.Fatal(err)
    }
    log.Println("Deleted successfully")
}
```

### Example 2: Context and Timeout / Context 및 타임아웃

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/arkd0ng/go-utils/httputil"
)

func main() {
    // Create context with timeout / 타임아웃이 있는 context 생성
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var result map[string]interface{}
    err := httputil.GetContext(ctx, "https://api.example.com/slow", &result)
    if err != nil {
        if httputil.IsTimeoutError(err) {
            log.Println("Request timed out")
        } else {
            log.Fatal(err)
        }
    }
}
```

### Example 3: Error Handling / 에러 처리

```go
package main

import (
    "log"

    "github.com/arkd0ng/go-utils/httputil"
)

func main() {
    var result map[string]interface{}
    err := httputil.Get("https://api.example.com/endpoint", &result,
        httputil.WithRetry(3))

    if err != nil {
        // Check error type / 에러 타입 확인
        switch {
        case httputil.IsHTTPError(err):
            statusCode := httputil.GetStatusCode(err)
            log.Printf("HTTP error: %d", statusCode)

            // Handle specific status codes / 특정 상태 코드 처리
            if statusCode == 404 {
                log.Println("Resource not found")
            } else if statusCode >= 500 {
                log.Println("Server error")
            }

        case httputil.IsRetryError(err):
            log.Println("All retry attempts failed")

        case httputil.IsTimeoutError(err):
            log.Println("Request timed out")

        default:
            log.Printf("Unknown error: %v", err)
        }
        return
    }

    log.Printf("Success: %+v", result)
}
```

### Example 4: Custom Headers and Query Parameters / 사용자 정의 헤더 및 쿼리 매개변수

```go
package main

import (
    "log"

    "github.com/arkd0ng/go-utils/httputil"
)

func main() {
    var result map[string]interface{}
    err := httputil.Get("https://api.example.com/search", &result,
        httputil.WithQueryParams(map[string]string{
            "q":     "golang",
            "page":  "1",
            "limit": "20",
        }),
        httputil.WithHeaders(map[string]string{
            "X-API-Version": "v1",
            "X-Request-ID":  "12345",
        }),
        httputil.WithBearerToken("your-token"),
    )

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Results: %+v", result)
}
```

---

## Best Practices / 모범 사례

### 1. Use Client for Multiple Requests / 여러 요청에 클라이언트 사용

Create a client instance when making multiple requests to the same API.

동일한 API에 여러 요청을 할 때는 클라이언트 인스턴스를 생성하세요.

```go
// Good: Reuse client / 좋음: 클라이언트 재사용
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithBearerToken("token"),
)
client.Get("/users", &users)
client.Get("/posts", &posts)

// Bad: Create new client each time / 나쁨: 매번 새 클라이언트 생성
httputil.Get("https://api.example.com/users", &users, httputil.WithBearerToken("token"))
httputil.Get("https://api.example.com/posts", &posts, httputil.WithBearerToken("token"))
```

### 2. Use Context for Cancellation / 취소를 위해 Context 사용

Always use Context variants for long-running requests.

장시간 실행되는 요청에는 항상 Context 변형을 사용하세요.

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

err := httputil.GetContext(ctx, url, &result)
```

### 3. Handle Errors Appropriately / 에러를 적절히 처리

Check error types and handle them specifically.

에러 타입을 확인하고 구체적으로 처리하세요.

```go
err := httputil.Get(url, &result)
if err != nil {
    if httputil.IsHTTPError(err) {
        // Handle HTTP errors / HTTP 에러 처리
    } else if httputil.IsTimeoutError(err) {
        // Handle timeouts / 타임아웃 처리
    }
}
```

### 4. Configure Retry Logic / 재시도 로직 설정

Adjust retry settings based on your needs.

필요에 따라 재시도 설정을 조정하세요.

```go
// For critical operations / 중요한 작업의 경우
client := httputil.NewClient(
    httputil.WithRetry(5),
    httputil.WithRetryBackoff(500*time.Millisecond, 30*time.Second),
)

// For non-critical operations / 중요하지 않은 작업의 경우
client := httputil.NewClient(
    httputil.WithRetry(1),
)
```

### 5. Set Appropriate Timeouts / 적절한 타임아웃 설정

Always set timeouts to prevent hanging requests.

항상 타임아웃을 설정하여 요청이 멈추는 것을 방지하세요.

```go
// Short timeout for health checks / 헬스 체크용 짧은 타임아웃
httputil.Get(url, &result, httputil.WithTimeout(5*time.Second))

// Longer timeout for complex operations / 복잡한 작업용 긴 타임아웃
httputil.Post(url, data, &result, httputil.WithTimeout(2*time.Minute))
```

---

## Documentation / 문서

### Additional Resources / 추가 자료

- **User Manual / 사용자 매뉴얼**: [docs/httputil/USER_MANUAL.md](../docs/httputil/USER_MANUAL.md)
- **Developer Guide / 개발자 가이드**: [docs/httputil/DEVELOPER_GUIDE.md](../docs/httputil/DEVELOPER_GUIDE.md)
- **Design Plan / 설계 계획**: [docs/httputil/DESIGN_PLAN.md](../docs/httputil/DESIGN_PLAN.md)
- **Work Plan / 작업 계획**: [docs/httputil/WORK_PLAN.md](../docs/httputil/WORK_PLAN.md)
- **Changelog / 변경 기록**: [docs/CHANGELOG/CHANGELOG-v1.10.md](../docs/CHANGELOG/CHANGELOG-v1.10.md)

### Package Documentation / 패키지 문서

View full package documentation:

전체 패키지 문서 보기:

```bash
go doc github.com/arkd0ng/go-utils/httputil
```

### Examples / 예제

See the examples directory for complete working examples:

완전한 작동 예제는 examples 디렉토리를 참조하세요:

```bash
# Run httputil example / httputil 예제 실행
go run examples/httputil/main.go
```

---

## Version History / 버전 히스토리

### v1.10.002 - 2025-10-15

**Phase 2-4 Features Added / 2-4단계 기능 추가**

- Response helpers (20+ methods) / 응답 헬퍼 (20개 이상 메서드)
- File operations (download/upload with progress) / 파일 작업 (진행 상황과 함께 다운로드/업로드)
- URL Builder (fluent API) / URL 빌더 (Fluent API)
- Form Builder (fluent API) / Form 빌더 (Fluent API)
- URL utilities (6 functions) / URL 유틸리티 (6개 함수)
- Form utilities (2 functions) / Form 유틸리티 (2개 함수)
- Extended Simple API (26+ functions) / 확장된 간단한 API (26개 이상 함수)
- Comprehensive tests (13 tests, 43+ sub-tests) / 종합 테스트 (13개 테스트, 43개 이상 하위 테스트)

### v1.10.001 - 2025-10-15

**Initial Release / 초기 릴리스**

- Core HTTP client with retry logic / 재시도 로직을 가진 핵심 HTTP 클라이언트
- Simple API (10 functions) / 간단한 API (10개 함수)
- Options pattern (12 options) / 옵션 패턴 (12개 옵션)
- Error types (3 types) / 에러 타입 (3개 타입)
- Comprehensive tests / 종합 테스트

---

## License / 라이선스

MIT License - See [LICENSE](../LICENSE) file for details.

MIT 라이선스 - 자세한 내용은 [LICENSE](../LICENSE) 파일을 참조하세요.

---

## Contributing / 기여

Contributions are welcome! Please see [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.

기여를 환영합니다! 가이드라인은 [CONTRIBUTING.md](../CONTRIBUTING.md)를 참조하세요.

---

**Package Author / 패키지 작성자**: arkd0ng
**Repository / 저장소**: https://github.com/arkd0ng/go-utils
**Go Version / Go 버전**: 1.18+
