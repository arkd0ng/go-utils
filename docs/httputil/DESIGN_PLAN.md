# httputil Package Design Plan
# httputil 패키지 설계 계획

## Overview / 개요

The `httputil` package provides extreme simplicity HTTP utilities for Go, reducing 30+ lines of repetitive HTTP code to just 2-3 lines.

`httputil` 패키지는 Go를 위한 극도로 간단한 HTTP 유틸리티를 제공하며, 30줄 이상의 반복적인 HTTP 코드를 단 2-3줄로 줄입니다.

## Design Philosophy / 설계 철학

**"30 lines → 2-3 lines"** - Extreme Simplicity for HTTP operations

Following the same principles as other packages:
- If not dramatically simpler, don't build it
- Auto everything: retries, timeouts, error handling
- Type-safe operations with generics
- Zero configuration needed

다른 패키지와 동일한 원칙 준수:
- 극적으로 간단하지 않으면 만들지 않음
- 모든 것 자동: 재시도, 타임아웃, 에러 처리
- 제네릭을 사용한 타입 안전 작업
- 설정 불필요

## Core Features / 핵심 기능

### 1. Simple HTTP Client / 간단한 HTTP 클라이언트
- **GET, POST, PUT, PATCH, DELETE** - RESTful methods with automatic JSON handling
- **Automatic retries** - Configurable retry logic with exponential backoff
- **Timeout management** - Smart timeout defaults with override options
- **Context support** - Cancellation and deadline propagation

### 2. Request Helpers / 요청 헬퍼
- **JSON encoding/decoding** - Automatic serialization and deserialization
- **Query parameters** - Easy query string building
- **Headers** - Fluent header management
- **Form data** - Multipart form and URL-encoded forms
- **File uploads** - Simple file upload handling

### 3. Response Helpers / 응답 헬퍼
- **Status code checking** - IsSuccess, IsRedirect, IsClientError, IsServerError
- **JSON parsing** - Type-safe JSON response parsing with generics
- **Error handling** - Smart error detection and wrapping
- **Body readers** - Convenient body reading utilities

### 4. Middleware Support / 미들웨어 지원
- **Request/Response logging** - Built-in logging middleware
- **Authentication** - Bearer token, Basic auth helpers
- **Custom headers** - Global header injection
- **Rate limiting** - Optional rate limiting support

## Package Structure / 패키지 구조

```
httputil/
├── httputil.go          # Package documentation and main types
├── client.go            # HTTP client wrapper with retry logic
├── request.go           # Request building helpers
├── response.go          # Response parsing helpers
├── middleware.go        # Middleware support
├── options.go           # Functional options pattern
├── errors.go            # Custom error types
├── helpers.go           # Utility functions
├── client_test.go       # Client tests
├── request_test.go      # Request tests
├── response_test.go     # Response tests
└── README.md            # Package documentation
```

## API Design / API 설계

### Simple API (most common use cases)

```go
// GET request / GET 요청
var result MyStruct
err := httputil.Get("https://api.example.com/data", &result)

// POST request / POST 요청
payload := MyPayload{Name: "test"}
var response MyResponse
err := httputil.Post("https://api.example.com/create", payload, &response)

// With options / 옵션 포함
err := httputil.Get("https://api.example.com/data", &result,
    httputil.WithTimeout(30*time.Second),
    httputil.WithHeaders(map[string]string{"Authorization": "Bearer token"}),
    httputil.WithRetry(3),
)
```

### Client API (for advanced scenarios)

```go
// Create client with defaults / 기본값으로 클라이언트 생성
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithTimeout(30*time.Second),
    httputil.WithRetry(3),
    httputil.WithBearerToken("your-token"),
)

// Make requests / 요청 실행
var result MyStruct
err := client.Get("/data", &result)
```

## Function Categories / 함수 카테고리

### Category 1: Simple HTTP Methods (10 functions)
- `Get(url string, result interface{}, opts ...Option) error`
- `GetContext(ctx context.Context, url string, result interface{}, opts ...Option) error`
- `Post(url string, body, result interface{}, opts ...Option) error`
- `PostContext(ctx context.Context, url string, body, result interface{}, opts ...Option) error`
- `Put(url string, body, result interface{}, opts ...Option) error`
- `PutContext(ctx context.Context, url string, body, result interface{}, opts ...Option) error`
- `Patch(url string, body, result interface{}, opts ...Option) error`
- `PatchContext(ctx context.Context, url string, body, result interface{}, opts ...Option) error`
- `Delete(url string, result interface{}, opts ...Option) error`
- `DeleteContext(ctx context.Context, url string, result interface{}, opts ...Option) error`

### Category 2: Request Builders (8 functions)
- `NewRequest(method, url string, body interface{}) (*Request, error)`
- `WithQueryParams(params map[string]string) Option`
- `WithHeaders(headers map[string]string) Option`
- `WithHeader(key, value string) Option`
- `WithBearerToken(token string) Option`
- `WithBasicAuth(username, password string) Option`
- `WithFormData(data map[string]string) Option`
- `WithMultipartForm(files map[string]string, fields map[string]string) Option`

### Category 3: Response Helpers (10 functions)
- `ParseJSON[T any](resp *http.Response) (T, error)` - Generic JSON parsing
- `IsSuccess(statusCode int) bool` - 2xx check
- `IsRedirect(statusCode int) bool` - 3xx check
- `IsClientError(statusCode int) bool` - 4xx check
- `IsServerError(statusCode int) bool` - 5xx check
- `ReadBody(resp *http.Response) ([]byte, error)`
- `ReadString(resp *http.Response) (string, error)`
- `CheckStatus(resp *http.Response) error` - Auto error on non-2xx
- `GetHeader(resp *http.Response, key string) string`
- `GetHeaders(resp *http.Response) map[string]string`

### Category 4: Client Configuration (12 functions)
- `NewClient(opts ...Option) *Client`
- `WithBaseURL(baseURL string) Option`
- `WithTimeout(timeout time.Duration) Option`
- `WithRetry(maxRetries int) Option`
- `WithRetryBackoff(min, max time.Duration) Option`
- `WithLogger(logger Logger) Option`
- `WithUserAgent(userAgent string) Option`
- `WithFollowRedirects(follow bool) Option`
- `WithMaxRedirects(max int) Option`
- `WithTLSConfig(config *tls.Config) Option`
- `WithProxy(proxyURL string) Option`
- `WithCookieJar(jar http.CookieJar) Option`

### Category 5: Download/Upload (6 functions)
- `DownloadFile(url, filepath string, opts ...Option) error`
- `DownloadFileContext(ctx context.Context, url, filepath string, opts ...Option) error`
- `UploadFile(url, fieldName, filepath string, opts ...Option) error`
- `UploadFileContext(ctx context.Context, url, fieldName, filepath string, opts ...Option) error`
- `UploadFiles(url string, files map[string]string, opts ...Option) error`
- `UploadFilesContext(ctx context.Context, url string, files map[string]string, opts ...Option) error`

### Category 6: Utilities (8 functions)
- `BuildURL(baseURL string, path string, params map[string]string) string`
- `EncodeQueryParams(params map[string]string) string`
- `ParseQueryParams(rawQuery string) map[string]string`
- `IsURL(str string) bool`
- `IsHTTPS(url string) bool`
- `GetContentType(resp *http.Response) string`
- `GetContentLength(resp *http.Response) int64`
- `SetDefaultClient(client *http.Client)`

**Total: ~54 functions across 6 categories**

## Key Design Decisions / 주요 설계 결정

### 1. Automatic JSON Handling
- Automatically marshal request body to JSON
- Automatically unmarshal response body from JSON
- Support for custom content types via options

### 2. Smart Retry Logic
- Automatic retry on network errors and 5xx responses
- Exponential backoff with jitter
- Configurable max retries (default: 3)
- Idempotent methods only (GET, PUT, DELETE) by default

### 3. Context Support
- All methods have Context variants
- Timeout and cancellation support
- Deadline propagation

### 4. Type Safety
- Generic response parsing: `ParseJSON[T]()`
- Compile-time type checking
- No interface{} casting needed

### 5. Error Handling
- Rich error types with status codes
- Wraps standard errors
- Includes request/response details for debugging

## Error Types / 에러 타입

```go
type HTTPError struct {
    StatusCode int
    Status     string
    Body       string
    URL        string
    Method     string
}

type RetryError struct {
    Attempts int
    LastErr  error
}

type TimeoutError struct {
    Timeout time.Duration
    URL     string
}
```

## Example Usage / 사용 예제

### Before (Standard Go) / 이전 (표준 Go)
```go
// 30+ lines
client := &http.Client{Timeout: 30 * time.Second}
reqBody, _ := json.Marshal(payload)
req, err := http.NewRequest("POST", "https://api.example.com/data", bytes.NewBuffer(reqBody))
if err != nil {
    return err
}
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer token")

resp, err := client.Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
    body, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
}

var result MyResponse
if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
    return err
}
// ... more error handling
```

### After (httputil) / 이후 (httputil)
```go
// 2-3 lines
var result MyResponse
err := httputil.Post("https://api.example.com/data", payload, &result,
    httputil.WithBearerToken("token"))
```

## Dependencies / 의존성

- **Standard library only**: net/http, encoding/json, context, time
- **Zero external dependencies**: Following the project's zero-dependency principle
- **Optional**: Could integrate with existing logging package

## Performance Considerations / 성능 고려사항

- Connection pooling via http.Client
- Configurable timeouts and retries
- Efficient JSON encoding/decoding
- Minimal allocations
- Benchmark against standard library

## Testing Strategy / 테스트 전략

1. **Unit tests**: All functions with mock HTTP server
2. **Integration tests**: Real HTTP requests to test endpoints
3. **Error handling tests**: Network errors, timeouts, retries
4. **Concurrent tests**: Thread-safety validation
5. **Benchmark tests**: Performance comparison

## Documentation Plan / 문서화 계획

1. **README.md**: Quick start, examples, API reference
2. **USER_MANUAL.md**: Comprehensive usage guide
3. **DEVELOPER_GUIDE.md**: Architecture, extending the package
4. **Examples**: Real-world usage scenarios

## Implementation Phases / 구현 단계

### Phase 1: Core Client (v1.10.001-003)
- Basic HTTP methods (GET, POST, PUT, DELETE)
- Automatic JSON handling
- Options pattern setup

### Phase 2: Advanced Features (v1.10.004-006)
- Retry logic with exponential backoff
- Context support
- Error types

### Phase 3: Request/Response Helpers (v1.10.007-009)
- Query parameter builders
- Header management
- Response parsing utilities

### Phase 4: File Operations (v1.10.010-012)
- File download
- File upload (single and multiple)
- Progress callbacks

### Phase 5: Documentation & Polish (v1.10.013-015)
- Comprehensive tests
- README and examples
- Performance optimization

## Success Criteria / 성공 기준

- ✅ Reduces 30+ lines to 2-3 lines for common HTTP operations
- ✅ 100% test coverage
- ✅ Zero external dependencies
- ✅ Type-safe with generics
- ✅ Automatic retry and error handling
- ✅ Context support for all operations
- ✅ Comprehensive documentation
- ✅ Real-world examples

---

**Version**: v1.10.x
**Status**: Design Phase
**Date**: 2025-10-15
