# httputil Package Work Plan
# httputil 패키지 작업 계획

## Project Overview / 프로젝트 개요

**Package**: `github.com/arkd0ng/go-utils/httputil`
**Version Range**: v1.10.x
**Goal**: Create extreme simplicity HTTP utilities reducing 30+ lines to 2-3 lines
**목표**: 30줄 이상의 코드를 2-3줄로 줄이는 극도로 간단한 HTTP 유틸리티 생성

## Development Phases / 개발 단계

### Phase 1: Core Foundation (v1.10.001-003) ✅ COMPLETED

**Status**: ✅ Completed in v1.10.001
**상태**: ✅ v1.10.001에서 완료

#### Tasks Completed / 완료된 작업:

1. **Package Structure Setup** ✅
   - Created httputil directory
   - Set up package documentation
   - Version management from cfg/app.yaml
   - 패키지 디렉토리 생성
   - 패키지 문서 설정
   - cfg/app.yaml에서 버전 관리

2. **Error Types** ✅ (errors.go)
   - HTTPError with status code and body
   - RetryError for retry failures
   - TimeoutError for timeouts
   - Helper functions: IsHTTPError, IsRetryError, IsTimeoutError, GetStatusCode
   - 상태 코드 및 본문을 포함한 HTTPError
   - 재시도 실패를 위한 RetryError
   - 타임아웃을 위한 TimeoutError
   - 헬퍼 함수들

3. **Options Pattern** ✅ (options.go)
   - Functional options for configuration
   - 12 configuration options implemented:
     - WithTimeout, WithHeaders, WithHeader
     - WithQueryParams, WithBearerToken, WithBasicAuth
     - WithRetry, WithRetryBackoff, WithUserAgent
     - WithBaseURL, WithFollowRedirects, WithMaxRedirects
   - Logger interface for optional logging
   - 설정을 위한 함수형 옵션
   - 12개 설정 옵션 구현
   - 선택적 로깅을 위한 Logger 인터페이스

4. **HTTP Client** ✅ (client.go)
   - Client struct wrapping http.Client
   - NewClient with configuration options
   - Smart retry logic with exponential backoff
   - Automatic JSON encoding/decoding
   - Context support for all methods
   - 10 HTTP methods (GET/POST/PUT/PATCH/DELETE + Context variants)
   - Client 구조체
   - 스마트 재시도 로직
   - 자동 JSON 인코딩/디코딩
   - 모든 메서드에 Context 지원

5. **Simple API** ✅ (simple.go)
   - Package-level convenience functions
   - Default client for quick usage
   - SetDefaultClient for global configuration
   - 패키지 레벨 편의 함수
   - 빠른 사용을 위한 기본 클라이언트

6. **Basic Tests** ✅ (httputil_test.go)
   - Version test
   - Client creation tests
   - Error type tests
   - Configuration tests
   - All tests passing
   - 모든 테스트 통과

**Files Created**: 6 files, ~890 lines of code
**생성된 파일**: 6개 파일, 약 890줄의 코드

---

### Phase 2: Response Helpers (v1.10.002-004) 📋 PLANNED

**Status**: 📋 Planned
**상태**: 📋 계획됨

#### Tasks / 작업:

1. **Response Utilities** (response.go)
   - `ParseJSON[T](resp) (T, error)` - Generic JSON parsing
   - `ReadBody(resp) ([]byte, error)` - Read response body
   - `ReadString(resp) (string, error)` - Read as string
   - `CheckStatus(resp) error` - Auto error on non-2xx
   - 제네릭 JSON 파싱
   - 응답 본문 읽기 유틸리티

2. **Status Code Helpers**
   - `IsSuccess(statusCode) bool` - Check 2xx
   - `IsRedirect(statusCode) bool` - Check 3xx
   - `IsClientError(statusCode) bool` - Check 4xx
   - `IsServerError(statusCode) bool` - Check 5xx
   - 상태 코드 확인 함수들

3. **Header Utilities**
   - `GetHeader(resp, key) string` - Get single header
   - `GetHeaders(resp) map[string]string` - Get all headers
   - `GetContentType(resp) string` - Get Content-Type
   - `GetContentLength(resp) int64` - Get Content-Length
   - 헤더 유틸리티 함수들

4. **Tests**
   - Response parsing tests
   - Status code check tests
   - Header utility tests
   - Mock HTTP server for testing
   - 응답 파싱 테스트
   - Mock HTTP 서버

**Estimated LOC**: ~300 lines
**예상 코드 줄 수**: 약 300줄

---

### Phase 3: Download/Upload (v1.10.005-007) 📋 PLANNED

**Status**: 📋 Planned
**상태**: 📋 계획됨

#### Tasks / 작업:

1. **Download Functions** (download.go)
   - `DownloadFile(url, filepath, opts...) error` - Download to file
   - `DownloadFileContext(ctx, url, filepath, opts...) error` - With context
   - Progress callback support via options
   - Resume support for interrupted downloads
   - 파일 다운로드 함수
   - 진행 상황 콜백 지원

2. **Upload Functions** (upload.go)
   - `UploadFile(url, fieldName, filepath, opts...) error` - Single file upload
   - `UploadFileContext(ctx, url, fieldName, filepath, opts...) error` - With context
   - `UploadFiles(url, files map[string]string, opts...) error` - Multiple files
   - `UploadFilesContext(ctx, url, files, opts...) error` - With context
   - Multipart form data support
   - Progress callback support
   - 파일 업로드 함수
   - 멀티파트 폼 데이터 지원

3. **Progress Callback**
   - `WithProgress(callback func(written, total int64)) Option`
   - Real-time progress tracking
   - 실시간 진행 상황 추적

4. **Tests**
   - Download tests with mock server
   - Upload tests with multipart forms
   - Progress callback tests
   - 다운로드/업로드 테스트

**Estimated LOC**: ~400 lines
**예상 코드 줄 수**: 약 400줄

---

### Phase 4: Utility Functions (v1.10.008-010) 📋 PLANNED

**Status**: 📋 Planned
**상태**: 📋 계획됨

#### Tasks / 작업:

1. **URL Utilities** (utils.go)
   - `BuildURL(baseURL, path, params) string` - Build URL with query params
   - `EncodeQueryParams(params) string` - Encode query parameters
   - `ParseQueryParams(rawQuery) map[string]string` - Parse query string
   - `IsURL(str) bool` - Validate URL
   - `IsHTTPS(url) bool` - Check if HTTPS
   - URL 유틸리티 함수들

2. **Form Data Helpers**
   - `WithFormData(data map[string]string) Option` - URL-encoded form
   - `WithMultipartForm(files, fields) Option` - Multipart form
   - Form data encoding
   - 폼 데이터 인코딩

3. **Request Building**
   - `NewRequest(method, url, body) (*Request, error)` - Build raw request
   - Request wrapper for manual control
   - 수동 제어를 위한 요청 래퍼

4. **Tests**
   - URL building tests
   - Query parameter tests
   - Form data tests
   - URL 빌딩 테스트

**Estimated LOC**: ~200 lines
**예상 코드 줄 수**: 약 200줄

---

### Phase 5: Examples & Documentation (v1.10.011-015) 📋 PLANNED

**Status**: 📋 Planned
**상태**: 📋 계획됨

#### Tasks / 작업:

1. **Example Program** (examples/httputil/main.go)
   - GET request example
   - POST request with authentication
   - File download example
   - File upload example
   - Client with base URL
   - Error handling examples
   - 예제 프로그램 작성

2. **Package README** (httputil/README.md)
   - Quick start guide
   - API reference
   - Usage examples
   - Configuration options
   - 빠른 시작 가이드
   - API 참조

3. **User Manual** (docs/httputil/USER_MANUAL.md)
   - Comprehensive usage guide
   - All functions documented
   - Real-world examples
   - Best practices
   - Troubleshooting
   - FAQ
   - 종합 사용 가이드
   - 모든 함수 문서화

4. **Developer Guide** (docs/httputil/DEVELOPER_GUIDE.md)
   - Architecture overview
   - Internal implementation
   - Design patterns used
   - Testing guide
   - Contributing guide
   - 아키텍처 개요
   - 내부 구현

5. **Root Documentation Updates**
   - Update root README.md with httputil section
   - Update root CHANGELOG.md
   - Update CLAUDE.md
   - 루트 문서 업데이트

**Estimated LOC**: ~2000+ lines of documentation
**예상 문서 줄 수**: 약 2000줄 이상

---

## Implementation Guidelines / 구현 가이드라인

### Code Quality Standards / 코드 품질 기준

1. **Testing**
   - 100% test coverage goal
   - Unit tests for all functions
   - Integration tests with mock HTTP server
   - Benchmark tests for performance
   - 100% 테스트 커버리지 목표
   - 모든 함수에 대한 단위 테스트

2. **Documentation**
   - Bilingual comments (English/Korean)
   - GoDoc-style documentation
   - Usage examples in comments
   - 이중 언어 주석
   - 사용 예제 포함

3. **Error Handling**
   - Rich error types with context
   - Wrapped errors for debugging
   - Clear error messages
   - 컨텍스트를 포함한 풍부한 에러 타입

4. **Performance**
   - Efficient retry logic
   - Connection pooling
   - Minimal allocations
   - Benchmark tests
   - 효율적인 재시도 로직
   - 연결 풀링

### Dependencies / 의존성

**Allowed Dependencies / 허용된 의존성:**
- Standard library: net/http, encoding/json, context, time, etc.
- gopkg.in/yaml.v3 (for version loading only)

**No External HTTP Libraries / 외부 HTTP 라이브러리 금지:**
- No third-party HTTP clients
- Pure standard library implementation
- 표준 라이브러리만 사용

### Design Principles / 설계 원칙

1. **Extreme Simplicity**
   - 30+ lines → 2-3 lines
   - Zero configuration for basic usage
   - Intuitive API naming
   - 극도의 간결함

2. **Type Safety**
   - Generic functions where appropriate
   - Compile-time type checking
   - No interface{} casting by users
   - 타입 안전성

3. **Auto Everything**
   - Automatic JSON handling
   - Automatic retries
   - Automatic error wrapping
   - 모든 것 자동화

4. **Context Support**
   - Context variants for all methods
   - Timeout and cancellation
   - Deadline propagation
   - Context 지원

---

## Progress Tracking / 진행 상황 추적

### Completed / 완료됨 ✅

- [x] Phase 1: Core Foundation (v1.10.001)
  - [x] Package structure
  - [x] Error types
  - [x] Options pattern
  - [x] HTTP client
  - [x] Simple API
  - [x] Basic tests
  - [x] Design plan
  - [x] Initial CHANGELOG

### In Progress / 진행 중 🚧

- [ ] Phase 1 Documentation
  - [ ] Package README
  - [ ] User Manual
  - [ ] Developer Guide
  - [ ] Root documentation updates

### Planned / 계획됨 📋

- [ ] Phase 2: Response Helpers (v1.10.002-004)
- [ ] Phase 3: Download/Upload (v1.10.005-007)
- [ ] Phase 4: Utility Functions (v1.10.008-010)
- [ ] Phase 5: Examples & Polish (v1.10.011-015)

---

## Version History / 버전 히스토리

### v1.10.001 - 2025-10-15 ✅
- Initial httputil package implementation
- Core HTTP client with retry logic
- Simple API (10 functions)
- Options pattern (12 options)
- Error types (3 types)
- Basic tests (7 tests passing)
- Design plan document

### v1.10.002 - Planned 📋
- Response helper functions
- Status code checks
- Header utilities
- Generic JSON parsing

### v1.10.003 - Planned 📋
- Response helper tests
- Mock HTTP server setup
- Integration tests

### v1.10.004 - Planned 📋
- Documentation improvements
- More usage examples
- Performance tests

### v1.10.005-007 - Planned 📋
- File download/upload
- Progress callbacks
- Multipart forms

### v1.10.008-010 - Planned 📋
- URL utilities
- Form data helpers
- Request building

### v1.10.011-015 - Planned 📋
- Example program
- Package README
- User Manual
- Developer Guide
- Root documentation updates

---

## Success Criteria / 성공 기준

### Functionality / 기능성
- ✅ Reduces 30+ lines to 2-3 lines
- ✅ RESTful HTTP methods (GET, POST, PUT, PATCH, DELETE)
- ✅ Automatic JSON handling
- ✅ Smart retry logic
- ✅ Context support
- ✅ Rich error types
- 📋 Response helpers
- 📋 File download/upload
- 📋 URL utilities

### Quality / 품질
- ✅ Zero external dependencies (except gopkg.in/yaml.v3)
- ✅ Type-safe operations
- ✅ All tests passing
- 📋 100% test coverage
- 📋 Comprehensive documentation
- 📋 Real-world examples

### Documentation / 문서화
- ✅ Package-level documentation
- ✅ Function-level documentation (bilingual)
- ✅ Design plan
- ✅ CHANGELOG entry
- 📋 Package README
- 📋 User Manual
- 📋 Developer Guide
- 📋 Root documentation updates

---

## Timeline / 타임라인

- **v1.10.001**: 2025-10-15 ✅ Completed
- **v1.10.002-004**: TBD (Response helpers)
- **v1.10.005-007**: TBD (Download/Upload)
- **v1.10.008-010**: TBD (Utilities)
- **v1.10.011-015**: TBD (Documentation & Examples)

---

**Status**: Phase 1 Complete, Documentation in Progress
**상태**: Phase 1 완료, 문서화 진행 중

**Next Immediate Task**: Complete documentation (README, USER_MANUAL, DEVELOPER_GUIDE)
**다음 즉시 작업**: 문서화 완료 (README, USER_MANUAL, DEVELOPER_GUIDE)
