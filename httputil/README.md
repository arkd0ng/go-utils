# httputil Package

Extremely simple HTTP client utilities that reduce 30+ lines to 2-3 lines.

극도로 간단한 HTTP 클라이언트 유틸리티로 30줄 이상의 코드를 2-3줄로 줄입니다.

**Version / 버전**: v1.10.001
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
- Standard library only (net/http, encoding/json)
- 표준 라이브러리만 사용 (net/http, encoding/json)

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
