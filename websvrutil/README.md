# websvrutil - Web Server Utilities / 웹 서버 유틸리티

**Version / 버전**: v1.11.023
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

## Current Features (v1.11.023) / 현재 기능

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
- `RunWithGracefulShutdown(addr string, timeout time.Duration) error` - Run with signal handling / 시그널 처리와 함께 실행
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

**Context Storage / 컨텍스트 저장소** (Enhanced in v1.11.017+):
- `Set(key string, value interface{})` - Store value / 값 저장
- `Get(key string) (interface{}, bool)` - Retrieve value / 값 검색
- `MustGet(key string) interface{}` - Get or panic / 가져오거나 패닉
- `GetString(key string) string` - Get string value / 문자열 값 가져오기
- `GetInt(key string) int` - Get int value / int 값 가져오기
- `GetInt64(key string) int64` - Get int64 value / int64 값 가져오기
- `GetFloat64(key string) float64` - Get float64 value / float64 값 가져오기
- `GetBool(key string) bool` - Get bool value / bool 값 가져오기
- `GetStringSlice(key string) []string` - Get string slice / 문자열 슬라이스 가져오기
- `GetStringMap(key string) map[string]interface{}` - Get string map / 문자열 맵 가져오기
- `Exists(key string) bool` - Check if key exists / 키 존재 확인
- `Delete(key string)` - Remove value / 값 제거
- `Keys() []string` - Get all keys / 모든 키 가져오기

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

**Request Binding / 요청 바인딩** (v1.11.013+):
- `Bind(obj interface{}) error` - Auto bind based on Content-Type / Content-Type에 따라 자동 바인딩
- `BindJSON(obj interface{}) error` - Bind JSON request body / JSON 요청 본문 바인딩
- `BindForm(obj interface{}) error` - Bind form data / 폼 데이터 바인딩
- `BindQuery(obj interface{}) error` - Bind query parameters / 쿼리 매개변수 바인딩

**Cookie Helpers / 쿠키 헬퍼** (v1.11.014+, Enhanced in v1.11.020):
- `Cookie(name string) (*http.Cookie, error)` - Get cookie by name / 이름으로 쿠키 가져오기
- `CookieValue(name string) string` - Get cookie value (returns empty string if not found) / 쿠키 값 가져오기 (없으면 빈 문자열)
- `SetCookie(cookie *http.Cookie)` - Set response cookie / 응답 쿠키 설정
- `SetCookieAdvanced(opts CookieOptions)` - Set cookie with advanced options / 고급 옵션으로 쿠키 설정
- `DeleteCookie(name, path string)` - Delete cookie / 쿠키 삭제
- `GetCookie(name string) string` - Get cookie value / 쿠키 값 가져오기

**HTTP Method Helpers / HTTP 메서드 헬퍼** (v1.11.023+):
- `IsGET() bool` - Check if request method is GET / GET 요청인지 확인
- `IsPOST() bool` - Check if request method is POST / POST 요청인지 확인
- `IsPUT() bool` - Check if request method is PUT / PUT 요청인지 확인
- `IsPATCH() bool` - Check if request method is PATCH / PATCH 요청인지 확인
- `IsDELETE() bool` - Check if request method is DELETE / DELETE 요청인지 확인
- `IsHEAD() bool` - Check if request method is HEAD / HEAD 요청인지 확인
- `IsOPTIONS() bool` - Check if request method is OPTIONS / OPTIONS 요청인지 확인
- `IsAjax() bool` - Check if request is AJAX (XMLHttpRequest) / AJAX 요청인지 확인
- `IsWebSocket() bool` - Check if request is WebSocket upgrade / WebSocket 업그레이드 요청인지 확인
- `AcceptsJSON() bool` - Check if client accepts JSON responses / JSON 응답 수락 확인
- `AcceptsHTML() bool` - Check if client accepts HTML responses / HTML 응답 수락 확인
- `AcceptsXML() bool` - Check if client accepts XML responses / XML 응답 수락 확인

**Header Helpers / 헤더 헬퍼** (v1.11.014+):
- `AddHeader(key, value string)` - Add response header / 응답 헤더 추가
- `GetHeader(key string) string` - Get request header / 요청 헤더 가져오기
- `GetHeaders(key string) []string` - Get all header values / 모든 헤더 값 가져오기
- `HeaderExists(key string) bool` - Check if header exists / 헤더 존재 확인
- `ContentType() string` - Get Content-Type header / Content-Type 헤더 가져오기
- `UserAgent() string` - Get User-Agent header / User-Agent 헤더 가져오기
- `Referer() string` - Get Referer header / Referer 헤더 가져오기
- `ClientIP() string` - Get client IP address / 클라이언트 IP 주소 가져오기

### Error Response Helpers / 에러 응답 헬퍼

The Context provides convenient methods for sending error responses:
Context는 에러 응답 전송을 위한 편리한 메서드를 제공합니다:

```go
// Abort with status code / 상태 코드로 중단
ctx.AbortWithStatus(http.StatusNotFound)
ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid input"))
ctx.AbortWithJSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})

// Standardized JSON responses / 표준화된 JSON 응답
ctx.ErrorJSON(400, "Invalid request")
// {"error":"Invalid request","status":400,"success":false}

ctx.SuccessJSON(200, "Operation successful", map[string]interface{}{"id": 123})
// {"message":"Operation successful","data":{"id":123},"status":200,"success":true}

// Common HTTP error shortcuts / 일반적인 HTTP 에러 단축
ctx.NotFound()              // 404
ctx.Unauthorized()          // 401
ctx.Forbidden()             // 403
ctx.BadRequest()            // 400
ctx.InternalServerError()   // 500
```


**File Upload / 파일 업로드** (v1.11.015+):
- `FormFile(name string) (*multipart.FileHeader, error)` - Get uploaded file / 업로드된 파일 가져오기
- `MultipartForm() (*multipart.Form, error)` - Get parsed multipart form / 파싱된 multipart 폼 가져오기
- `SaveUploadedFile(file *multipart.FileHeader, dst string) error` - Save uploaded file / 업로드된 파일 저장

**Static File Serving / 정적 파일 서빙** (v1.11.016+):
- `File(filepath string) error` - Send file response to client / 클라이언트에게 파일 응답 전송
- `FileAttachment(filepath, filename string) error` - Send file as downloadable attachment / 다운로드 가능한 첨부 파일로 전송
- `Static(prefix, dir string) *App` (App method) - Serve static files from directory / 디렉토리에서 정적 파일 제공

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
| `WithMaxUploadSize(size int64)` | 32 MB | Maximum file upload size / 최대 파일 업로드 크기 |

### Session Management / 세션 관리

Cookie-based session management with in-memory storage.

쿠키 기반 세션 관리 및 인메모리 저장소.

**SessionStore Methods / SessionStore 메서드**:
- `NewSessionStore(opts SessionOptions) *SessionStore` - Create session store / 세션 저장소 생성
- `Get(r *http.Request) (*Session, error)` - Get or create session / 세션 가져오기 또는 생성
- `New() *Session` - Create new session / 새 세션 생성
- `Save(w http.ResponseWriter, session *Session)` - Save session and set cookie / 세션 저장 및 쿠키 설정
- `Destroy(w http.ResponseWriter, r *http.Request) error` - Destroy session / 세션 파괴
- `Count() int` - Get active session count / 활성 세션 수 가져오기

**Session Methods / Session 메서드**:
- `Set(key string, value interface{})` - Store value / 값 저장
- `Get(key string) (interface{}, bool)` - Retrieve value / 값 검색
- `GetString(key string) string` - Get string value / 문자열 값 가져오기
- `GetInt(key string) int` - Get int value / int 값 가져오기
- `GetBool(key string) bool` - Get bool value / bool 값 가져오기
- `Delete(key string)` - Remove value / 값 제거
- `Clear()` - Clear all values / 모든 값 지우기

**SessionOptions / 세션 옵션**:
- `CookieName` - Session cookie name (default: "sessionid") / 세션 쿠키 이름
- `MaxAge` - Session expiration (default: 24h) / 세션 만료 시간
- `Secure` - HTTPS only (default: false) / HTTPS만 사용
- `HttpOnly` - Prevent JavaScript access (default: true) / JavaScript 액세스 방지
- `SameSite` - SameSite attribute (default: Lax) / SameSite 속성
- `Path` - Cookie path (default: "/") / 쿠키 경로
- `Domain` - Cookie domain / 쿠키 도메인
- `CleanupTime` - Cleanup interval (default: 5m) / 정리 간격

**Features / 기능**:
- Automatic session expiration / 자동 세션 만료
- Background cleanup of expired sessions / 만료된 세션의 백그라운드 정리
- Cryptographically secure session IDs / 암호학적으로 안전한 세션 ID
- Thread-safe operations / 스레드 안전 작업
- Type-safe getters / 타입 안전 getter

### Middleware / 미들웨어

Built-in middleware for common use cases.

일반적인 사용 사례를 위한 내장 미들웨어.

**Recovery Middleware / 복구 미들웨어**:
- `Recovery()` - Panic recovery with logging / 로깅과 함께 패닉 복구
- `RecoveryWithConfig(config)` - Custom recovery configuration / 커스텀 복구 설정
- Automatically logs panics with stack traces / 스택 트레이스와 함께 패닉 자동 로깅
- Sends 500 Internal Server Error on panic / 패닉 시 500 Internal Server Error 전송

**Logger Middleware / 로거 미들웨어**:
- `Logger()` - HTTP request logging / HTTP 요청 로깅
- `LoggerWithConfig(config)` - Custom logger configuration / 커스텀 로거 설정
- Logs method, path, status code, duration / 메서드, 경로, 상태 코드, 소요 시간 로깅
- Customizable log function / 커스터마이즈 가능한 로그 함수

**CORS Middleware / CORS 미들웨어**:
- `CORS()` - Cross-Origin Resource Sharing / Cross-Origin Resource Sharing
- `CORSWithConfig(config)` - Custom CORS configuration / 커스텀 CORS 설정
- Configurable origins, methods, headers / 설정 가능한 오리진, 메서드, 헤더
- Automatic preflight request handling / 자동 프리플라이트 요청 처리
- Credentials and max-age support / 자격 증명 및 max-age 지원

**RequestID Middleware / 요청 ID 미들웨어**:
- `RequestID()` - Adds unique request ID to each request / 각 요청에 고유한 요청 ID 추가
- `RequestIDWithConfig(config)` - Custom RequestID configuration / 커스텀 요청 ID 설정
- Stores ID in context and response headers / 컨텍스트 및 응답 헤더에 ID 저장
- Customizable header name and ID generator / 커스터마이즈 가능한 헤더 이름 및 ID 생성기
- Preserves existing request ID if present / 기존 요청 ID가 있으면 보존

**Timeout Middleware / 타임아웃 미들웨어**:
- `Timeout(duration)` - Enforces request timeout / 요청 타임아웃 적용
- `TimeoutWithConfig(config)` - Custom timeout configuration / 커스텀 타임아웃 설정
- Sends 503 Service Unavailable on timeout / 타임아웃 시 503 Service Unavailable 전송
- Configurable timeout duration and message / 설정 가능한 타임아웃 기간 및 메시지
- Default timeout: 30 seconds / 기본 타임아웃: 30초

**BasicAuth Middleware / Basic 인증 미들웨어**:
- `BasicAuth(username, password)` - HTTP Basic Authentication / HTTP Basic Authentication
- `BasicAuthWithConfig(config)` - Custom BasicAuth configuration / 커스텀 Basic 인증 설정
- Constant-time password comparison / 상수 시간 비밀번호 비교
- Customizable realm and validator function / 커스터마이즈 가능한 영역 및 검증자 함수
- Stores username in context for later use / 나중에 사용하기 위해 컨텍스트에 사용자 이름 저장

**RateLimiter Middleware / Rate Limiter 미들웨어**:
- `RateLimiter(requests, window)` - IP-based rate limiting / IP 기반 rate limiting
- `RateLimiterWithConfig(config)` - Custom rate limiter configuration / 커스텀 rate limiter 설정
- Token bucket algorithm with sliding window / 슬라이딩 윈도우와 함께 토큰 버킷 알고리즘
- Returns 429 Too Many Requests when limit exceeded / 제한 초과 시 429 Too Many Requests 반환
- X-RateLimit headers (Limit, Remaining, Reset) / X-RateLimit 헤더 (제한, 남은 수, 리셋 시간)
- Customizable key function (IP, API key, etc.) / 커스터마이즈 가능한 키 함수

**Compression Middleware / 압축 미들웨어**:
- `Compression()` - Gzip compression for responses / 응답에 대한 Gzip 압축
- `CompressionWithConfig(config)` - Custom compression configuration / 커스텀 압축 설정
- Automatic compression when client supports gzip / 클라이언트가 gzip을 지원할 때 자동 압축
- Configurable compression level (1-9) / 설정 가능한 압축 레벨
- Default minimum size: 1KB / 기본 최소 크기: 1KB
- Sets Content-Encoding header / Content-Encoding 헤더 설정

**SecureHeaders Middleware / 보안 헤더 미들웨어**:
- `SecureHeaders()` - Adds security-related HTTP headers / 보안 관련 HTTP 헤더 추가
- `SecureHeadersWithConfig(config)` - Custom secure headers configuration / 커스텀 보안 헤더 설정
- X-Frame-Options (SAMEORIGIN) - Clickjacking protection / 클릭재킹 보호
- X-Content-Type-Options (nosniff) - MIME sniffing protection / MIME 스니핑 보호
- X-XSS-Protection (1; mode=block) - XSS protection / XSS 보호
- Strict-Transport-Security (HSTS) - HTTPS enforcement / HTTPS 강제
- Content-Security-Policy - CSP support / CSP 지원
- Referrer-Policy - Referrer control / 리퍼러 제어

**BodyLimit Middleware / 본문 제한 미들웨어**:
- `BodyLimit(maxBytes)` - Limits maximum request body size / 최대 요청 본문 크기 제한
- `BodyLimitWithConfig(config)` - Custom body limit configuration / 커스텀 본문 제한 설정
- Default limit: 10MB / 기본 제한: 10MB
- Returns error when body exceeds limit / 본문이 제한을 초과하면 에러 반환
- Prevents memory exhaustion attacks / 메모리 고갈 공격 방지

**Static Middleware / 정적 파일 미들웨어**:
- `Static(root)` - Serves static files from directory / 디렉토리에서 정적 파일 제공
- `StaticWithConfig(config)` - Custom static file configuration / 커스텀 정적 파일 설정
- Automatic index.html serving / 자동 index.html 제공
- Optional directory browsing / 선택적 디렉토리 탐색
- Falls through to next handler if file not found / 파일을 찾을 수 없으면 다음 핸들러로 전달

**Redirect Middleware / 리디렉션 미들웨어**:
- `Redirect(to)` - Redirects all requests to URL / 모든 요청을 URL로 리디렉션
- `RedirectWithConfig(config)` - Custom redirect configuration / 커스텀 리디렉션 설정
- `HTTPSRedirect()` - Redirects HTTP to HTTPS / HTTP를 HTTPS로 리디렉션
- `WWWRedirect(addWWW)` - Adds or removes www prefix / www 접두사 추가 또는 제거
- Default: 301 Moved Permanently / 기본값: 301 Moved Permanently
- Configurable status code / 설정 가능한 상태 코드

### Template System / 템플릿 시스템

Built-in template engine for HTML rendering with auto-discovery and custom functions.

HTML 렌더링을 위한 내장 템플릿 엔진 (자동 발견 및 커스텀 함수 지원).

**TemplateEngine / 템플릿 엔진**:
- Automatic template loading from directory / 디렉토리에서 자동 템플릿 로딩
- Support for nested directories / 중첩 디렉토리 지원
- **Layout system** / **레이아웃 시스템** (v1.11.011+)
- **26+ built-in template functions** / **26개 이상의 내장 템플릿 함수** (v1.11.011+)
- **Hot reload in development** / **개발 시 핫 리로드** (v1.11.012+)
- Custom template functions / 커스텀 템플릿 함수
- Custom delimiters / 커스텀 구분자
- Thread-safe template caching / 스레드 안전 템플릿 캐싱

**Template Methods / 템플릿 메서드**:
- `NewTemplateEngine(dir)` - Create new template engine / 새 템플릿 엔진 생성
- `Load(name)` - Load single template / 단일 템플릿 로드
- `LoadGlob(pattern)` - Load templates by pattern / 패턴으로 템플릿 로드
- `LoadAll()` - Load all templates recursively / 모든 템플릿 재귀적 로드
- `Render(w, name, data)` - Render template to writer / writer에 템플릿 렌더링
- `AddFunc(name, fn)` - Add custom template function / 커스텀 템플릿 함수 추가
- `AddFuncs(funcs)` - Add multiple functions / 여러 함수 추가
- `SetDelimiters(left, right)` - Set custom delimiters / 커스텀 구분자 설정
- `Has(name)` - Check if template exists / 템플릿 존재 확인
- `List()` - List all loaded templates / 모든 로드된 템플릿 목록
- `Clear()` - Remove all templates / 모든 템플릿 제거

**App Template Methods / 앱 템플릿 메서드**:
- `app.TemplateEngine()` - Get template engine instance / 템플릿 엔진 인스턴스 가져오기
- `app.LoadTemplate(name)` - Load single template / 단일 템플릿 로드
- `app.LoadTemplates(pattern)` - Load templates by pattern / 패턴으로 템플릿 로드
- `app.ReloadTemplates()` - Reload all templates / 모든 템플릿 다시 로드
- `app.AddTemplateFunc(name, fn)` - Add custom function / 커스텀 함수 추가
- `app.AddTemplateFuncs(funcs)` - Add multiple functions / 여러 함수 추가

**Context Template Methods / Context 템플릿 메서드**:
- `ctx.Render(code, name, data)` - Render template file / 템플릿 파일 렌더링
- `ctx.RenderWithLayout(code, layout, template, data)` - Render with layout / 레이아웃과 함께 렌더링 (v1.11.011+)
- `ctx.HTMLTemplate(code, tmpl, data)` - Render inline template / 인라인 템플릿 렌더링

**Layout Methods / 레이아웃 메서드** (v1.11.011+):
- `SetLayoutDir(dir)` - Set layout directory / 레이아웃 디렉토리 설정
- `LoadLayout(name)` - Load single layout / 단일 레이아웃 로드
- `LoadAllLayouts()` - Load all layouts / 모든 레이아웃 로드
- `RenderWithLayout(w, layout, template, data)` - Render with layout / 레이아웃과 함께 렌더링
- `HasLayout(name)` - Check if layout exists / 레이아웃 존재 확인
- `ListLayouts()` - List all loaded layouts / 모든 로드된 레이아웃 목록

**Built-in Template Functions** (v1.11.011+):
- **String functions**: `upper`, `lower`, `title`, `trim`, `trimPrefix`, `trimSuffix`, `replace`, `contains`, `hasPrefix`, `hasSuffix`, `split`, `join`, `repeat`
- **Date/Time functions**: `now`, `formatDate`, `formatDateSimple`, `formatDateTime`, `formatTime`
- **URL functions**: `urlEncode`, `urlDecode`
- **Safe HTML**: `safeHTML`, `safeURL`, `safeJS`
- **Utility functions**: `default`, `len`

**Hot Reload Methods** (v1.11.012+):
- `EnableAutoReload()` - Enable automatic template reloading / 자동 템플릿 재로드 활성화
- `DisableAutoReload()` - Disable automatic template reloading / 자동 템플릿 재로드 비활성화
- `IsAutoReloadEnabled()` - Check if auto-reload is enabled / 자동 재로드 활성화 확인

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

### Graceful Shutdown (Simple) / 정상 종료 (간단)

```go
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    app := websvrutil.New()

    app.GET("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    // RunWithGracefulShutdown automatically handles SIGINT and SIGTERM signals
    // RunWithGracefulShutdown은 SIGINT 및 SIGTERM 시그널을 자동으로 처리합니다
    if err := app.RunWithGracefulShutdown(":8080", 5*time.Second); err != nil {
        log.Fatal(err)
    }

    log.Println("Server stopped gracefully")
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

### File Upload / 파일 업로드

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "path/filepath"
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    // Create app with custom max upload size (10 MB)
    // 커스텀 최대 업로드 크기(10 MB)로 앱 생성
    app := websvrutil.New(
        websvrutil.WithMaxUploadSize(10 << 20), // 10 MB
    )

    // Single file upload / 단일 파일 업로드
    app.POST("/upload", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        // Get uploaded file / 업로드된 파일 가져오기
        file, err := ctx.FormFile("file")
        if err != nil {
            http.Error(w, "Failed to get file: "+err.Error(), http.StatusBadRequest)
            return
        }

        // Save file to disk / 파일을 디스크에 저장
        dstPath := filepath.Join("uploads", file.Filename)
        if err := ctx.SaveUploadedFile(file, dstPath); err != nil {
            http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "File uploaded successfully: %s (%d bytes)",
            file.Filename, file.Size)
    })

    // Multiple file upload / 여러 파일 업로드
    app.POST("/upload-multiple", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        // Get multipart form / multipart 폼 가져오기
        form, err := ctx.MultipartForm()
        if err != nil {
            http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
            return
        }

        // Get all uploaded files / 모든 업로드된 파일 가져오기
        files := form.File["files"]

        for _, file := range files {
            dstPath := filepath.Join("uploads", file.Filename)

            // Save each file / 각 파일 저장
            if err := ctx.SaveUploadedFile(file, dstPath); err != nil {
                fmt.Fprintf(w, "Failed to save %s: %v\n", file.Filename, err)
                continue
            }

            fmt.Fprintf(w, "Saved: %s (%d bytes)\n", file.Filename, file.Size)
        }
    })

    // File upload with form fields / 폼 필드와 함께 파일 업로드
    app.POST("/upload-with-data", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)

        // Get multipart form / multipart 폼 가져오기
        form, err := ctx.MultipartForm()
        if err != nil {
            http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
            return
        }

        // Get form fields / 폼 필드 가져오기
        title := form.Value["title"][0]
        description := form.Value["description"][0]

        // Get uploaded file / 업로드된 파일 가져오기
        file := form.File["file"][0]
        dstPath := filepath.Join("uploads", file.Filename)

        if err := ctx.SaveUploadedFile(file, dstPath); err != nil {
            http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "Title: %s\nDescription: %s\nFile: %s (%d bytes)",
            title, description, file.Filename, file.Size)
    })

    log.Println("Server starting on :8080")
    if err := app.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
```

### Static File Serving / 정적 파일 서빙

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

    // Serve static files from ./public directory at /static/* URLs
    // ./public 디렉토리의 파일을 /static/* URL에서 제공
    app.Static("/static", "./public")

    // Serve static files from multiple directories
    // 여러 디렉토리에서 정적 파일 제공
    app.Static("/css", "./assets/css")
    app.Static("/js", "./assets/js")
    app.Static("/images", "./assets/images")

    // Serve a specific file / 특정 파일 제공
    app.GET("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        ctx.File("./public/favicon.ico")
    })

    // Serve file as download / 파일을 다운로드로 제공
    app.GET("/download/report", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        ctx.FileAttachment("./reports/monthly.pdf", "monthly-report.pdf")
    })

    // Dynamic file serving based on parameter / 매개변수 기반 동적 파일 제공
    app.GET("/files/:filename", func(w http.ResponseWriter, r *http.Request) {
        ctx := websvrutil.GetContext(r)
        filename := ctx.Param("filename")
        filepath := fmt.Sprintf("./uploads/%s", filename)
        ctx.File(filepath)
    })

    log.Println("Server starting on :8080")
    if err := app.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
```

### Session Management / 세션 관리

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
    app := websvrutil.New()

    // Create session store with custom options / 커스텀 옵션으로 세션 저장소 생성
    sessionStore := websvrutil.NewSessionStore(websvrutil.SessionOptions{
        CookieName:  "mysession",
        MaxAge:      24 * time.Hour,
        HttpOnly:    true,
        Secure:      false, // Set to true in production with HTTPS / HTTPS를 사용하는 프로덕션에서는 true로 설정
        SameSite:    http.SameSiteLaxMode,
        Path:        "/",
        CleanupTime: 5 * time.Minute,
    })

    // Login handler - create session / 로그인 핸들러 - 세션 생성
    app.POST("/login", func(w http.ResponseWriter, r *http.Request) {
        // Get or create session / 세션 가져오기 또는 생성
        session, _ := sessionStore.Get(r)

        // Store user data in session / 세션에 사용자 데이터 저장
        session.Set("user_id", 123)
        session.Set("username", "john")
        session.Set("authenticated", true)

        // Save session and set cookie / 세션 저장 및 쿠키 설정
        sessionStore.Save(w, session)

        fmt.Fprintf(w, "Login successful! Session ID: %s", session.ID)
    })

    // Protected route - check session / 보호된 라우트 - 세션 확인
    app.GET("/dashboard", func(w http.ResponseWriter, r *http.Request) {
        session, err := sessionStore.Get(r)
        if err != nil {
            http.Error(w, "Session error", http.StatusInternalServerError)
            return
        }

        // Check if user is authenticated / 사용자 인증 확인
        authenticated := session.GetBool("authenticated")
        if !authenticated {
            http.Error(w, "Not authenticated", http.StatusUnauthorized)
            return
        }

        // Get user data from session / 세션에서 사용자 데이터 가져오기
        username := session.GetString("username")
        userID := session.GetInt("user_id")

        fmt.Fprintf(w, "Welcome %s (ID: %d)!", username, userID)
    })

    // Logout handler - destroy session / 로그아웃 핸들러 - 세션 파괴
    app.POST("/logout", func(w http.ResponseWriter, r *http.Request) {
        if err := sessionStore.Destroy(w, r); err != nil {
            http.Error(w, "Logout failed", http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Logged out successfully")
    })

    // Session info endpoint / 세션 정보 엔드포인트
    app.GET("/session/info", func(w http.ResponseWriter, r *http.Request) {
        session, _ := sessionStore.Get(r)

        fmt.Fprintf(w, "Session ID: %s\n", session.ID)
        fmt.Fprintf(w, "Created: %s\n", session.CreatedAt)
        fmt.Fprintf(w, "Expires: %s\n", session.ExpiresAt)
        fmt.Fprintf(w, "Active Sessions: %d\n", sessionStore.Count())
    })

    log.Println("Server starting on :8080")
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

**Current Phase / 현재 단계**: Phase 2 - Middleware System (v1.11.006-010)

**Progress / 진행 상황**:
- ✅ v1.11.001: Project setup and planning / 프로젝트 설정 및 계획
- ✅ v1.11.002: App & Options / 앱 및 옵션
- ✅ v1.11.003: Router / 라우터
- ✅ v1.11.004: Context (Part 1) / 컨텍스트 (1부)
- ✅ v1.11.005: Response Helpers / 응답 헬퍼
- ✅ v1.11.006: Middleware (Recovery, Logger, CORS) / 미들웨어 (복구, 로거, CORS)
- ✅ v1.11.007: Additional Middleware (RequestID, Timeout, BasicAuth) / 추가 미들웨어 (요청 ID, 타임아웃, Basic 인증)
- ✅ v1.11.008: Advanced Middleware (RateLimiter, Compression, SecureHeaders) / 고급 미들웨어 (Rate Limiter, 압축, 보안 헤더)
- ✅ v1.11.009: Final Middleware (BodyLimit, Static, Redirect, HTTPSRedirect, WWWRedirect) / 최종 미들웨어 (본문 제한, 정적 파일, 리디렉션)
- ✅ v1.11.010: Template Engine Core (TemplateEngine, auto-loading, custom functions) / 템플릿 엔진 핵심 (TemplateEngine, 자동 로딩, 커스텀 함수)
- ✅ v1.11.011: Layout System and Built-in Functions (layouts, 26+ built-in functions) / 레이아웃 시스템 및 내장 함수 (레이아웃, 26개 이상의 내장 함수)
- ✅ v1.11.012: Hot Reload (automatic template reloading, file watching) / 핫 리로드 (자동 템플릿 재로드, 파일 감시)
- ✅ v1.11.013: Request Binding (Bind, BindJSON, BindForm, BindQuery) / 요청 바인딩 (Bind, BindJSON, BindForm, BindQuery)
- ✅ v1.11.014: Cookie & Header Helpers (Cookie, SetCookie, DeleteCookie, GetCookie, AddHeader, GetHeaders, ClientIP) / 쿠키 및 헤더 헬퍼
- ✅ v1.11.015: File Upload (FormFile, MultipartForm, SaveUploadedFile, MaxUploadSize) / 파일 업로드 (FormFile, MultipartForm, SaveUploadedFile, MaxUploadSize)
- ✅ v1.11.016: Static File Serving (Static, File, FileAttachment) / 정적 파일 서빙 (Static, File, FileAttachment)
- ✅ v1.11.017: Context Storage Enhancement (GetInt64, GetFloat64, GetStringSlice, GetStringMap, Exists, Delete, Keys) / 컨텍스트 저장소 향상

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
