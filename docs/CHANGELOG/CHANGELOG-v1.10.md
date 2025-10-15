# CHANGELOG - v1.10.x

This file contains detailed change logs for the v1.10.x releases of go-utils, focusing on the httputil package.

이 파일은 httputil 패키지에 중점을 둔 go-utils의 v1.10.x 릴리스에 대한 상세한 변경 로그를 포함합니다.

---

## [v1.10.003] - 2025-10-15

### Added / 추가됨

#### Documentation and Examples / 문서 및 예제

**Documentation Updates / 문서 업데이트**:
- Updated `httputil/README.md` with Phase 2-4 features and examples
- Updated `docs/httputil/USER_MANUAL.md` with comprehensive Phase 2-4 usage patterns (섹션 5.6-5.9, 6.6-6.7 추가)
- Updated `docs/httputil/DEVELOPER_GUIDE.md` with architecture details for Phase 2-4 components (섹션 3.5-3.8 추가)
- All documentation fully bilingual (English/Korean)
- httputil/README.md를 Phase 2-4 기능 및 예제로 업데이트
- 종합적인 Phase 2-4 사용 패턴으로 USER_MANUAL.md 업데이트
- Phase 2-4 컴포넌트를 위한 아키텍처 세부 정보로 DEVELOPER_GUIDE.md 업데이트
- 모든 문서 완전히 이중 언어 (영문/한글)

**Examples / 예제**:
- Created comprehensive example program (`examples/httputil/main.go`)
- 10 complete examples demonstrating all Phase 1-4 features:
  - Example 1: Basic HTTP Methods (Phase 1)
  - Example 2: Client with Base URL (Phase 1)
  - Example 3: Context and Timeout (Phase 1)
  - Example 4: Error Handling (Phase 1)
  - Example 5: Response Helpers (Phase 2)
  - Example 6: File Download (Phase 3)
  - Example 7: File Upload (Phase 3)
  - Example 8: URL Builder (Phase 4)
  - Example 9: Form Builder (Phase 4)
  - Example 10: Comprehensive Workflow (All Phases)
- Fully bilingual comments and explanations
- Example compiles and runs successfully
- 종합 예제 프로그램 생성
- 모든 Phase 1-4 기능을 보여주는 10개의 완전한 예제
- 완전히 이중 언어 주석 및 설명
- 예제 컴파일 및 실행 성공

**README Updates / README 업데이트**:
- Added 5 new Quick Start examples (4-8) for Phase 2-4
- Added API Reference sections for Response, URL Builder, Form Builder
- Added WithProgress option to configuration table
- Updated version history with v1.10.002 changelog
- Phase 2-4를 위한 5개의 새로운 빠른 시작 예제 추가
- Response, URL Builder, Form Builder를 위한 API 참조 섹션 추가
- 설정 테이블에 WithProgress 옵션 추가
- v1.10.002 변경 로그로 버전 히스토리 업데이트

**USER_MANUAL Updates / USER_MANUAL 업데이트**:
- Added subsections 5.6-5.9 in Usage Patterns section:
  - 5.6: Response Helpers
  - 5.7: File Download
  - 5.8: File Upload
  - 5.9: URL and Form Builders
- Added subsections 6.6-6.7 in Common Use Cases section:
  - 6.6: File Download Service with Progress (완전한 예제)
  - 6.7: File Upload Service (완전한 예제)
- 사용 패턴 섹션에 하위 섹션 5.6-5.9 추가
- 일반적인 사용 사례 섹션에 하위 섹션 6.6-6.7 추가

**DEVELOPER_GUIDE Updates / DEVELOPER_GUIDE 업데이트**:
- Updated file organization and responsibilities sections
- Added subsections 3.5-3.8 in Core Components section:
  - 3.5: Response Wrapper (Phase 2)
  - 3.6: File Operations (Phase 3)
  - 3.7: URL Builder (Phase 4)
  - 3.8: Form Builder (Phase 4)
- Complete architecture diagrams and design patterns
- 파일 구성 및 책임 섹션 업데이트
- 핵심 컴포넌트 섹션에 하위 섹션 3.5-3.8 추가
- 완전한 아키텍처 다이어그램 및 디자인 패턴

**Test Statistics / 테스트 통계**:
- Updated httputil_test.go: 456 lines
- Total: 13 tests, 43+ sub-tests
- All tests passing (0.643s execution time)
- httputil_test.go 업데이트: 456줄
- 총 13개 테스트, 43개 이상 하위 테스트
- 모든 테스트 통과 (0.643초 실행 시간)

### Changed / 변경됨
- None / 없음

### Fixed / 수정됨
- None / 없음

---

## [v1.10.002] - 2025-10-15

### Added / 추가됨

#### Phase 2-4 Features / Phase 2-4 기능

**Phase 2: Response Helpers (response.go)** / **Phase 2: 응답 헬퍼**:
- `Response` struct wrapping http.Response with additional methods
- `DoRaw/DoRawContext` methods returning raw response
- Response body methods: `Body()`, `String()`, `JSON()`
- Status check methods: `IsSuccess()`, `IsError()`, `IsClientError()`, `IsServerError()`
- Specific status checks: `IsOK()`, `IsCreated()`, `IsNotFound()`, `IsUnauthorized()` etc.
- Header helpers: `Header()`, `Headers()`, `ContentType()`, `ContentLength()`
- Response 구조체로 http.Response 래핑 및 추가 메서드
- 원시 응답 반환하는 DoRaw/DoRawContext 메서드
- 응답 본문 메서드
- 상태 확인 메서드
- 헤더 헬퍼

**Phase 3: File Download/Upload (file.go)** / **Phase 3: 파일 다운로드/업로드**:
- `DownloadFile/DownloadFileContext` - Download file with progress callback
- `Download/DownloadContext` - Download data as bytes
- `UploadFile/UploadFileContext` - Upload single file with multipart/form-data
- `UploadFiles/UploadFilesContext` - Upload multiple files
- `ProgressFunc` callback for tracking upload/download progress
- 진행 상황 콜백과 함께 파일 다운로드
- 바이트로 데이터 다운로드
- multipart/form-data로 단일 파일 업로드
- 여러 파일 업로드
- 업로드/다운로드 진행 상황 추적 콜백

**Phase 4: URL and Form Utilities (url.go, form.go)** / **Phase 4: URL 및 폼 유틸리티**:
- `URLBuilder` - Fluent API for building URLs
- URL utilities: `JoinURL()`, `AddQueryParams()`, `ParseURL()`, `GetQueryParam()`
- URL helpers: `GetDomain()`, `GetScheme()`, `GetPath()`, `IsAbsoluteURL()`, `NormalizeURL()`
- `FormBuilder` - Fluent API for building form data
- `PostForm/PostFormContext` - POST with application/x-www-form-urlencoded
- Form utilities: `ParseForm()`, `EncodeForm()`
- URLBuilder로 URL 구축을 위한 Fluent API
- URL 유틸리티 함수들
- FormBuilder로 폼 데이터 구축을 위한 Fluent API
- application/x-www-form-urlencoded로 POST
- 폼 유틸리티

**Package Updates / 패키지 업데이트**:
- Extended `simple.go` with package-level functions for all new features
- All new methods available on both Client and package level
- Maintained zero external dependencies principle
- All tests passing
- 모든 새 기능에 대한 패키지 레벨 함수로 simple.go 확장
- Client 및 패키지 레벨 모두에서 사용 가능한 모든 새 메서드
- 외부 의존성 제로 원칙 유지
- 모든 테스트 통과

**New Files / 새 파일**:
- `response.go` - 280 lines, response helper methods
- `file.go` - 340 lines, file upload/download with progress
- `url.go` - 180 lines, URL builder and utilities
- `form.go` - 200 lines, form builder and utilities

**Statistics / 통계**:
- Total new code: ~1,000 lines
- New methods: 50+ functions/methods
- Test coverage: All core functionality tested
- 총 새 코드: 약 1,000줄
- 새 메서드: 50개 이상 함수/메서드
- 테스트 커버리지: 모든 핵심 기능 테스트됨

### Changed / 변경됨
- None / 없음

### Fixed / 수정됨
- None / 없음

---

## [v1.10.001] - 2025-10-15

### Added / 추가됨

#### Initial httputil Package Implementation / httputil 패키지 초기 구현

**Core Features / 핵심 기능:**
- Extreme simplicity HTTP utilities reducing 30+ lines to 2-3 lines
- Complete RESTful HTTP methods (GET, POST, PUT, PATCH, DELETE)
- Automatic JSON encoding/decoding for requests and responses
- Smart retry logic with exponential backoff and jitter
- Context support for cancellation and timeouts
- Type-safe operations with rich error types
- Zero external dependencies (standard library only)
- 30줄 이상의 코드를 2-3줄로 줄이는 극도로 간단한 HTTP 유틸리티
- 완전한 RESTful HTTP 메서드 (GET, POST, PUT, PATCH, DELETE)
- 요청 및 응답에 대한 자동 JSON 인코딩/디코딩
- 지수 백오프 및 지터를 사용한 스마트 재시도 로직
- 취소 및 타임아웃을 위한 Context 지원
- 풍부한 에러 타입을 사용한 타입 안전 작업
- 외부 의존성 제로 (표준 라이브러리만)

**Package Structure / 패키지 구조:**
- `httputil.go` - Package documentation and version management
- `client.go` - HTTP client wrapper with retry logic
- `simple.go` - Package-level convenience functions
- `options.go` - Functional options pattern for configuration
- `errors.go` - Custom error types (HTTPError, RetryError, TimeoutError)
- `httputil_test.go` - Comprehensive tests
- `httputil.go` - 패키지 문서 및 버전 관리
- `client.go` - 재시도 로직을 포함한 HTTP 클라이언트 래퍼
- `simple.go` - 패키지 레벨 편의 함수
- `options.go` - 설정을 위한 함수형 옵션 패턴
- `errors.go` - 사용자 정의 에러 타입
- `httputil_test.go` - 포괄적인 테스트

**API Design / API 설계:**

1. **Simple API (10 functions)** - Most common use cases / 가장 일반적인 사용 사례:
   - `Get(url, result, opts...)` - GET request with auto JSON decoding
   - `GetContext(ctx, url, result, opts...)` - GET with context
   - `Post(url, body, result, opts...)` - POST request with body
   - `PostContext(ctx, url, body, result, opts...)` - POST with context
   - `Put(url, body, result, opts...)` - PUT request with body
   - `PutContext(ctx, url, body, result, opts...)` - PUT with context
   - `Patch(url, body, result, opts...)` - PATCH request with body
   - `PatchContext(ctx, url, body, result, opts...)` - PATCH with context
   - `Delete(url, result, opts...)` - DELETE request
   - `DeleteContext(ctx, url, result, opts...)` - DELETE with context

2. **Client API** - For advanced scenarios / 고급 시나리오용:
   - `NewClient(opts...)` - Create configured client
   - `Client.Get/Post/Put/Patch/Delete` - Instance methods
   - `SetDefaultClient(client)` - Configure default client

3. **Configuration Options (12 options)** - Flexible configuration / 유연한 설정:
   - `WithTimeout(duration)` - Set request timeout
   - `WithHeaders(headers)` - Set custom headers
   - `WithHeader(key, value)` - Set single header
   - `WithQueryParams(params)` - Set query parameters
   - `WithBearerToken(token)` - Set Bearer authentication
   - `WithBasicAuth(user, pass)` - Set Basic authentication
   - `WithRetry(max)` - Set max retry attempts
   - `WithRetryBackoff(min, max)` - Configure retry backoff
   - `WithUserAgent(ua)` - Set custom User-Agent
   - `WithBaseURL(url)` - Set base URL for client
   - `WithFollowRedirects(follow)` - Enable/disable redirects
   - `WithMaxRedirects(max)` - Set max redirects

4. **Error Types** - Rich error information / 풍부한 에러 정보:
   - `HTTPError` - HTTP errors with status code and response body
   - `RetryError` - Errors after all retry attempts
   - `TimeoutError` - Request timeout errors
   - Helper functions: `IsHTTPError`, `IsRetryError`, `IsTimeoutError`, `GetStatusCode`

**Example Usage / 사용 예제:**

Simple GET request:
```go
var result MyStruct
err := httputil.Get("https://api.example.com/data", &result)
```

POST request with options:
```go
payload := MyPayload{Name: "test"}
var response MyResponse
err := httputil.Post("https://api.example.com/create", payload, &response,
    httputil.WithBearerToken("token"),
    httputil.WithTimeout(30*time.Second),
    httputil.WithRetry(3),
)
```

Using a client:
```go
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com"),
    httputil.WithBearerToken("token"),
)
var result MyStruct
err := client.Get("/data", &result)
```

**Design Philosophy / 설계 철학:**
- "30 lines → 2-3 lines" - Extreme Simplicity
- Auto everything: JSON handling, retries, error wrapping
- Type-safe with generics where possible
- Zero configuration needed for basic usage
- Context support everywhere

**Testing / 테스트:**
- ✅ Unit tests for all core functionality
- ✅ Error type tests
- ✅ Configuration tests
- ✅ Version loading tests
- ✅ All tests passing

**Documentation / 문서:**
- ✅ Package-level documentation with examples
- ✅ Function-level documentation (bilingual English/Korean)
- ✅ Design plan document (DESIGN_PLAN.md)
- ✅ Clear API with intuitive naming

**Dependencies / 의존성:**
- Standard library only: net/http, encoding/json, context, time
- gopkg.in/yaml.v3 (for version loading from cfg/app.yaml)
- Zero external HTTP dependencies

---

**Status / 상태**: Initial Release / 초기 릴리스
**Next Steps / 다음 단계**: Add response helpers, download/upload utilities, comprehensive examples
