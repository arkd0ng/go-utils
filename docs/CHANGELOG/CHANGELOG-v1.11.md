## [v1.11.036] - 2025-10-16

### Documentation / 문서화
- **Example Code Writing Guide** / **예제 코드 작성 가이드** (`docs/EXAMPLE_CODE_GUIDE.md`)
  - Comprehensive guide for writing example code across all packages
  - 모든 패키지에 걸쳐 예제 코드 작성을 위한 종합 가이드

  **Guide Contents / 가이드 내용**:
  - Directory structure and file naming conventions
  - Logging standards with backup management (keep 5 most recent)
  - Code structure templates (main function, example functions)
  - Example categories: Basic, Combination, Real-World, Edge Case, Performance, Integration
  - Documentation standards: Bilingual comments, structured logging, Unicode symbols
  - Package-specific guidelines (websvrutil: 14 categories with detailed requirements)
  - Best practices: Progressive examples, self-documenting code, error demonstration, resource cleanup
  - Comprehensive checklist for new examples (30+ items)

  **Key Standards / 주요 표준**:
  - All logs must be saved to `logs/` directory / 모든 로그는 `logs/` 디렉토리에 저장
  - Log file format: `logs/{package}-example.log` (current) / 로그 파일 형식
  - Backup format: `logs/{package}-example-YYYYMMDD-HHMMSS.log` / 백업 형식
  - Keep only 5 most recent backups, auto-delete older files / 최근 5개만 유지, 자동 삭제
  - Extremely detailed logging (logs should be self-documenting) / 매우 상세한 로깅
  - Bilingual comments and logs (English/Korean) / 이중 언어 주석 및 로그
  - Unicode symbols for console output (✅ ❌ ⚠️ ℹ️ 🔍 📊 🔧 🚀 🏁 📝) / 콘솔 출력용 유니코드 기호

### Test Results / 테스트 결과
- ✅ All existing tests passing
- ✅ Documentation guide created successfully
- ✅ 모든 기존 테스트 통과
- ✅ 문서화 가이드 생성 완료

---

## [v1.11.035] - 2025-10-16

### Test Coverage Final Improvement / 테스트 커버리지 최종 개선
- **Coverage increase from 88.5% to 89.5%** / **커버리지 88.5%에서 89.5%로 증가**
  - Additional 1% coverage improvement (88.5% → 89.5%)
  - 추가 1% 커버리지 향상

- **New complete coverage test file** / **새로운 완전 커버리지 테스트 파일** (`coverage_complete_test.go`)
  - **CSRF Token Expiry**: Tests CSRF token cleanup through expiry mechanism
  - **Template File Detection**: Tests template file detection through LoadGlob (HTML, HTM, TMPL, TPL vs other files)
  - **Template Auto-Reload**: Tests template auto-reload functionality with file watching
  - **Graceful Shutdown Complete**: Tests all shutdown paths (shutdown before start, timeout, multiple calls)
  - **Validator isZero Complete**: Tests isZero with all types (int, int8-64, uint, uint8-64, float32-64, string, bool, slice, map, ptr)
  - **Validator String Formats**: Email, alpha, alphanum, numeric validation edge cases
  - **Context Render Error Paths**: Tests render with invalid template, missing template
  - **Middleware Edge Cases**: Logger with custom function, body limit with exact/over limit
  - **Bind Edge Cases**: Unsupported content type handling
  - **GetContext Edge Cases**: Nil context value, wrong type in context
  - **Session Edge Cases**: Destroy with invalid cookie, get and save multiple times
  - **CSRF Multiple Token Sources**: Query parameter token lookup
  - **Route Group With Middleware**: Nested route groups, empty prefix groups
  - CSRF 토큰 만료, 템플릿 파일 감지, 자동 리로드, 우아한 종료, 검증자, 렌더링, 미들웨어, 바인딩, 컨텍스트, 세션, 라우트 그룹 테스트

### Test Results / 테스트 결과
- ✅ All tests passing
- ✅ Coverage: 89.5% of statements (increased from 88.5%)
- ✅ Total cumulative improvement: 79.9% → 89.5% (9.6% increase across 3 test files)
- ✅ 총 누적 향상: 79.9% → 89.5% (3개 테스트 파일에 걸쳐 9.6% 증가)

---

## [v1.11.034] - 2025-10-16

### Test Coverage Further Improvement / 테스트 커버리지 추가 개선
- **Coverage increase from 84.5% to 88.5%** / **커버리지 84.5%에서 88.5%로 증가**
  - Additional 4% coverage improvement (84.5% → 88.5%)
  - 추가 4% 커버리지 향상

- **New additional test file** / **새로운 추가 테스트 파일** (`coverage_additional_test.go`)
  - **Template Engine Methods**: TemplateEngine, LoadTemplate, LoadTemplates, ReloadTemplates, AddTemplateFunc, AddTemplateFuncs
  - **Render Methods**: Render, RenderWithLayout with actual templates
  - **Validator Coverage Improvement**: validateMin/Max/Gt/Gte/Lt/Lte with all types (int, uint, float, string, slice, map, array)
  - **Validator Edge Cases**: isZero, validateEmail, validateAlpha/Alphanum edge cases, BindWithValidation error paths
  - **Middleware Coverage**: RecoveryWithConfig panic recovery and normal flow
  - **CSRF Internal Methods**: Cleanup goroutine, token validation with invalid formats
  - **Template Internal Methods**: isTemplateFile, addBuiltinFuncs, RenderWithLayout
  - **Graceful Shutdown**: Multiple shutdown calls, shutdown before start
  - 템플릿 엔진 메서드, 렌더 메서드, 검증자 커버리지 향상, 미들웨어, CSRF 내부 메서드, 우아한 종료 테스트

### Test Results / 테스트 결과
- ✅ All tests passing
- ✅ Coverage: 88.5% of statements (increased from 84.5%)
- ✅ Total improvement: 79.9% → 88.5% (8.6% increase)
- ✅ 전체 향상: 79.9% → 88.5% (8.6% 증가)

---

## [v1.11.033] - 2025-10-16

### Test Coverage Improvement / 테스트 커버리지 개선
- **Coverage increase from 79.9% to 84.5%** / **커버리지 79.9%에서 84.5%로 증가**
  - Added 169 new tests (259 → 428 tests total)
  - 169개의 새로운 테스트 추가 (총 259 → 428개 테스트)

- **New comprehensive test file** / **새로운 종합 테스트 파일** (`coverage_test.go`)
  - **HTTP Methods Coverage**: Tests for PUT, PATCH, DELETE, OPTIONS, HEAD methods
  - **Context Request Methods**: HeaderExists, ContentType, UserAgent, Referer, ClientIP, AddHeader, GetHeader, GetHeaders
  - **Context Cookie Methods**: GetCookie
  - **NotFound Handler**: Custom 404 handler testing
  - **Max Body Size**: Request body size limit enforcement
  - **Binding Edge Cases**: BindJSON, BindForm, Bind with various error conditions
  - **File Upload**: File and FileAttachment methods
  - **CSRF Edge Cases**: GetCSRFToken, query parameter lookup
  - **Validator Edge Cases**: All validation tags with different types (int, float, string, slice)
  - **Graceful Shutdown**: RunWithGracefulShutdown functionality
  - **Security Input Validation**: SQL injection, XSS, path traversal, large headers, null bytes
  - **Concurrency**: Thread-safe Context operations
  - **Error Paths**: Invalid JSON, missing files, cookies, form data
  - HTTP 메서드 커버리지, 컨텍스트 메서드, 쿠키, 파일 업로드, CSRF, 검증자, 보안, 동시성, 에러 경로 테스트

### Test Results / 테스트 결과
- ✅ All 428 tests passing
- ✅ Coverage: 84.5% of statements (increased from 79.9%)
- ✅ No regressions

---

## [v1.11.032] - 2025-10-16

### Code Organization / 코드 구성
- **File Split** / **파일 분할** (`context.go` → 5 files)
  - Split large context.go file (1,475 lines) into 5 logical files
  - `context.go` (302 lines) - Core struct, params/values management
  - `context_request.go` (454 lines) - Request info methods
  - `context_response.go` (296 lines) - Response rendering methods
  - `context_bind.go` (305 lines) - Data binding & file operations
  - `context_helpers.go` (215 lines) - Helper methods
  - 거대한 context.go 파일 (1,475줄)을 5개의 논리적 파일로 분할
  - 코드 가독성 및 유지보수성 향상

### Security / 보안
- **CSRF Protection Middleware** / **CSRF 보호 미들웨어** (`csrf.go`)
  - Cryptographically secure token generation with crypto/rand
  - Flexible token lookup (header, form, query)
  - Cookie-based token storage with customizable options
  - Request method filtering (safe methods: GET, HEAD, OPTIONS, TRACE)
  - Token validation with constant-time comparison (timing attack prevention)
  - Skipper function support for conditional CSRF validation
  - Automatic token cleanup with periodic garbage collection
  - 암호학적으로 안전한 토큰 생성 (crypto/rand)
  - 유연한 토큰 조회 (헤더, 폼, 쿼리)
  - 쿠키 기반 토큰 저장 (커스텀 옵션 지원)
  - 타이밍 공격 방지를 위한 상수 시간 비교
  - Tests: 5 comprehensive tests (`csrf_test.go`)

### Validation / 검증
- **Validation Tag Support** / **검증 태그 지원** (`validator.go`)
  - Built-in validator with 14 validation tags:
    - `required`, `email`, `min`, `max`, `len`
    - `eq`, `ne`, `gt`, `gte`, `lt`, `lte`
    - `oneof`, `alpha`, `alphanum`, `numeric`
  - Support for multiple tags per field
  - Type-safe validation for string, int, uint, float, slice, map, array
  - Detailed validation errors with field and tag information
  - Context method: `BindWithValidation(obj)` for auto bind + validate
  - 14개의 검증 태그 지원
  - 여러 태그 조합 가능
  - 타입 안전 검증 (문자열, 정수, 실수, 슬라이스, 맵, 배열)
  - Tests: 12 comprehensive tests (`validator_test.go`)

### Testing / 테스트
- **Integration Tests** / **통합 테스트** (`integration_test.go`)
  - Full app integration test with multiple middlewares
  - Route groups integration test
  - CSRF + validation integration test
  - 여러 미들웨어가 있는 전체 앱 통합 테스트
  - 라우트 그룹 통합 테스트
  - CSRF + 검증 통합 테스트
  - Tests: 3 integration tests

- **Benchmark Tests** / **벤치마크 테스트** (`benchmark_test.go`)
  - Context operations (Get/Set)
  - JSON response rendering
  - Routing (simple and with params)
  - Middleware execution
  - CSRF token generation
  - Validation
  - ClientIP extraction
  - Parameter extraction
  - Benchmarks: 10 performance benchmarks

### Test Results / 테스트 결과
- ✅ All 259 tests pass (increased from 237)
  - 219 unit tests
  - 18 example tests
  - 5 CSRF tests
  - 12 validator tests
  - 3 integration tests
  - 10 benchmark tests (separate)
- ✅ Zero breaking changes
- ✅ All functionality preserved

### Code Quality Improvements / 코드 품질 개선
- Better file organization with logical separation
- Enhanced security with CSRF protection
- Improved data validation capabilities
- More comprehensive test coverage
- Performance benchmarking added
- 논리적 분리를 통한 더 나은 파일 구성
- CSRF 보호로 향상된 보안
- 개선된 데이터 검증 기능
- 더 포괄적인 테스트 커버리지
- 성능 벤치마크 추가

---

## Summary of v1.11.x - Code Quality Improvements / v1.11.x 요약 - 코드 품질 개선

이번 버전에서는 websvrutil 패키지의 코드 품질, 보안, 유지보수성을 대폭 개선했습니다:

**High Priority Tasks Completed (v1.11.024-027) / 고우선순위 작업 완료:**
1. ✅ **Comprehensive Bilingual Comments** (v1.11.024)
   - Enhanced router.go, bind.go, session.go, context.go, middleware.go
   - Added algorithm descriptions, time complexity, security considerations
   - Improved developer experience with detailed internal documentation

2. ✅ **Code Refactoring** (v1.11.025)
   - Removed HTTP method registration duplication (~120 lines → ~50 lines, 58% reduction)
   - Improved maintainability with DRY principle

3. ✅ **Security Enhancement** (v1.11.026)
   - Added request body size limits (default: 10 MB)
   - DoS protection with io.LimitReader
   - Configurable via WithMaxBodySize() option

4. ✅ **Error Message Improvements** (v1.11.027)
   - Enhanced error messages with type information
   - Better debugging experience with descriptive errors

**Medium Priority Tasks Completed (v1.11.028) / 중우선순위 작업 완료:**
5. ✅ **Route Group Support** (v1.11.028)
   - Added Group functionality for organizing routes hierarchically
   - Support for nested groups with prefix concatenation
   - Group-specific middleware with inheritance
   - Method chaining for fluent API

**Test Results / 테스트 결과:**
- ✅ All 208 tests pass (2 skipped) - increased from 199 tests
- ✅ Test coverage: maintained
- ✅ No regressions

**Overall Quality Metrics / 전체 품질 지표:**
- Code Quality: 9/10 (improved from 8/10)
- Documentation: 9/10 (improved from 7/10)
- Test Coverage: 8/10 (maintained)
- Security: 9/10 (improved with body size limits)
- Feature Completeness: 9/10 (improved with route groups)

---

## [v1.11.029] - 2025-10-16

### Performance / 성능
- **Performance Optimizations** / **성능 최적화**
  1. **ClientIP() String Operations Optimization** / **ClientIP() 문자열 작업 최적화** (`context.go`)
     - Replaced manual loop with `strings.IndexByte()` for better performance
     - Added `strings.TrimSpace()` for X-Forwarded-For IP extraction
     - Applied optimization to both X-Forwarded-For and RemoteAddr parsing
     - Performance improvement: O(n) manual loop → O(n) optimized stdlib function

  2. **Context Values Map Lazy Allocation** / **컨텍스트 값 맵 지연 할당** (`context.go`)
     - Changed `NewContext()` to not allocate `values` map by default (nil)
     - Implemented lazy initialization in `Set()` method
     - Memory saving: One less map allocation per request when not used
     - Optimization benefit: Reduces memory allocations for requests that don't use context storage

### Testing / 테스트
- **Updated Context Tests for Lazy Allocation** / **지연 할당을 위한 컨텍스트 테스트 업데이트** (`context_test.go`)
  - Modified `TestNewContext` to verify values map is nil initially
  - Added lazy allocation verification in `TestContextSetGet`
  - Verifies values map is created after first `Set()` call
  - All 219 tests pass ✅ (increased from 208)

### Code Quality / 코드 품질
- **Import Statement Updated** / **Import 문 업데이트** (`context.go`)
  - Added `strings` package import for `IndexByte()` and `TrimSpace()` functions

### Technical Details / 기술 세부사항
- **Before (ClientIP)** / **이전 (ClientIP)**:
  ```go
  for idx := 0; idx < len(xff); idx++ {
      if xff[idx] == ',' {
          return xff[:idx]
      }
  }
  ```
- **After (ClientIP)** / **이후 (ClientIP)**:
  ```go
  if idx := strings.IndexByte(xff, ','); idx != -1 {
      return strings.TrimSpace(xff[:idx])
  }
  ```

- **Before (NewContext)** / **이전 (NewContext)**:
  ```go
  values: make(map[string]interface{}), // Always allocated
  ```
- **After (NewContext)** / **이후 (NewContext)**:
  ```go
  values: nil, // Lazy allocation in Set()
  ```

### Benefits / 이점
- Improved performance for ClientIP() with stdlib optimized functions
- Reduced memory allocations per request (one less map when not needed)
- Better IP extraction with whitespace trimming
- Maintains all existing functionality with zero breaking changes

---

## [v1.11.030] - 2025-10-16

### Code Organization / 코드 조직화
- **Constants File Added** / **상수 파일 추가** (`constants.go`)
  - Created centralized constants file for better code organization
  - Extracted all magic numbers and hardcoded strings into named constants
  - Improved code maintainability and readability

### Constants Added / 추가된 상수
1. **Default Timeout Configurations** / **기본 타임아웃 설정**:
   - `DefaultReadTimeout` = 15 seconds
   - `DefaultWriteTimeout` = 15 seconds
   - `DefaultIdleTimeout` = 60 seconds

2. **Default Size Limits** / **기본 크기 제한**:
   - `DefaultMaxHeaderBytes` = 1 MB
   - `DefaultMaxBodySize` = 10 MB (for JSON/form data)
   - `DefaultMaxUploadSize` = 32 MB (for file uploads)

3. **Default Session Configurations** / **기본 세션 설정**:
   - `DefaultSessionMaxAge` = 24 hours
   - `DefaultSessionCookieName` = "sessionid"
   - `DefaultSessionCleanup` = 5 minutes

4. **Content-Type Constants** / **Content-Type 상수**:
   - `ContentTypeJSON` = "application/json; charset=utf-8"
   - `ContentTypeHTML` = "text/html; charset=utf-8"
   - `ContentTypeXML` = "application/xml; charset=utf-8"
   - `ContentTypeText` = "text/plain; charset=utf-8"
   - `ContentTypeForm` = "application/x-www-form-urlencoded"
   - `ContentTypeMultipart` = "multipart/form-data"

5. **HTTP Header Constants** / **HTTP 헤더 상수**:
   - `HeaderContentType`, `HeaderAccept`, `HeaderAuthorization`
   - `HeaderUserAgent`, `HeaderXForwardedFor`, `HeaderXRealIP`

### Refactoring / 리팩토링
- **Updated All Files to Use Constants** / **모든 파일을 상수 사용으로 업데이트**:
  - `options.go`: defaultOptions() now uses named constants
  - `context.go`: All Content-Type and size limits use constants
  - `session.go`: DefaultSessionOptions() uses named constants
  - `middleware.go`: BodyLimitWithConfig() uses DefaultMaxBodySize

### Benefits / 이점
- Eliminated magic numbers throughout the codebase
- Single source of truth for configuration values
- Easier to maintain and update default values
- Better code documentation through named constants
- Improved code readability and maintainability
- Reduced risk of inconsistent default values

### Testing / 테스트
- All 219 tests pass ✅
- No breaking changes
- All existing functionality preserved

---

## [v1.11.031] - 2025-10-16

### Documentation / 문서화
- **Godoc Examples Added** / **Godoc 예제 추가** (`example_test.go`)
  - Added 18 comprehensive testable examples for godoc
  - All examples demonstrate real-world usage patterns
  - Examples cover key functionality across the package

### Examples Added / 추가된 예제
1. **Application Setup** / **애플리케이션 설정**:
   - `Example()`: Complete application with middleware and routes
   - `ExampleNew()`: Creating a new application
   - `ExampleNew_withOptions()`: Creating app with custom options

2. **Route Registration** / **라우트 등록**:
   - `ExampleApp_GET()`: Registering GET routes
   - `ExampleApp_POST()`: Registering POST routes
   - `ExampleApp_Group()`: Creating route groups

3. **Context Operations** / **컨텍스트 작업**:
   - `ExampleContext_Param()`: URL parameter retrieval
   - `ExampleContext_JSON()`: Sending JSON responses
   - `ExampleContext_BindJSON()`: Binding JSON request bodies
   - `ExampleContext_Query()`: Query parameter retrieval
   - `ExampleContext_SetCookie()`: Setting cookies

4. **Middleware** / **미들웨어**:
   - `ExampleLogger()`: Logger middleware usage
   - `ExampleRecovery()`: Recovery middleware usage
   - `ExampleCORS()`: CORS middleware with custom config

5. **Session Management** / **세션 관리**:
   - `ExampleNewSessionStore()`: Creating session store

### Benefits / 이점
- Improved godoc documentation with runnable examples
- Better developer experience with copy-paste ready code
- All examples are testable and verified to work
- Covers most common use cases

### Testing / 테스트
- All 219 unit tests pass ✅
- All 18 example tests pass ✅
- No breaking changes
- Total: 237 passing tests

### Code Review Progress / 코드 리뷰 진행사항
- ✅ Task 10 completed: Add Godoc examples
- ✅ Task 11 evaluated: context.go file split (deferred - too complex)
- Remaining: Low-priority tasks (CSRF, validation, integration tests)

---

## [v1.11.028] - 2025-10-16

### Features / 기능
- **Route Group Support Added** / **라우트 그룹 지원 추가** (`group.go`)
  - Implemented `Group` struct with prefix, middleware, and app reference
  - Added `App.Group(prefix)` method to create route groups
  - Added `Group.Group(prefix)` method for nested groups with prefix concatenation
  - Added `Group.Use(middleware...)` for group-specific middleware
  - Implemented all HTTP methods on Group: GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD
  - Method chaining support for fluent API
  - Middleware inheritance: nested groups inherit parent middleware
  - Middleware wrapping in reverse order for correct execution sequence

### Features / 기능
- **Route Organization** / **라우트 구성**
  - Hierarchical route structure with common prefixes
  - Example: `/api/v1/admin/users` from nested groups
  - Group-specific middleware applied to all routes in group
  - Cleaner, more maintainable route organization

### Testing / 테스트
- **Comprehensive Group Tests Added** / **종합 그룹 테스트 추가** (`group_test.go`)
  - 9 new test functions covering all Group functionality:
    - `TestGroup_BasicGroupCreation`: Basic group creation and route registration
    - `TestGroup_NestedGroups`: Nested group prefix concatenation
    - `TestGroup_GroupMiddleware`: Group-specific middleware application
    - `TestGroup_MiddlewareInheritance`: Middleware inheritance in nested groups
    - `TestGroup_AllHTTPMethods`: All 7 HTTP methods (GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD)
    - `TestGroup_MethodChaining`: Fluent API method chaining
    - `TestGroup_MultipleMiddleware`: Multiple middleware execution order
    - `TestGroup_EmptyPrefix`: Group with empty prefix
    - `TestGroup_DeepNesting`: Deeply nested groups (4 levels)
  - All 9 tests pass ✅
  - Total test count: 208 tests (increased from 199)

### Documentation / 문서화
- Added comprehensive bilingual documentation to `group.go`:
  - Group struct documentation with features and examples
  - App.Group() method documentation
  - Group.Group() method documentation for nested groups
  - Group.Use() method documentation for middleware
  - All HTTP method documentation (GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD)
  - registerRoute() internal method documentation with process flow

### Examples / 예제
```go
// Create API group with authentication middleware
// 인증 미들웨어가 있는 API 그룹 생성
api := app.Group("/api")
api.Use(AuthMiddleware())

// Create v1 API subgroup
// v1 API 하위 그룹 생성
v1 := api.Group("/v1")  // Prefix: /api/v1

// Create admin subgroup with additional middleware
// 추가 미들웨어가 있는 admin 하위 그룹 생성
admin := v1.Group("/admin")  // Prefix: /api/v1/admin
admin.Use(AdminMiddleware())

// Register routes - all inherit middleware
// 라우트 등록 - 모든 라우트가 미들웨어 상속
admin.GET("/users", listUsers)      // Route: /api/v1/admin/users
admin.POST("/users", createUser)    // Route: /api/v1/admin/users
admin.DELETE("/users/:id", deleteUser) // Route: /api/v1/admin/users/:id

// Method chaining support
// 메서드 체이닝 지원
v1.GET("/stats", getStats).
   POST("/reports", createReport).
   PUT("/config", updateConfig)
```

---

## [v1.11.027] - 2025-10-16

### Improvements / 개선
- **Error Messages Enhanced with Type Information** / **타입 정보를 포함한 에러 메시지 개선** (`bind.go`)
  - Improved `bindFormData()` error messages to include actual type received
    - "binding requires a pointer to a struct, got %s"
    - "got pointer to %s" for non-struct pointers
  - Enhanced `setFieldValue()` error messages with:
    - Actual value attempted to convert: `cannot convert value "abc" to int`
    - Target type information: includes full type (int, bool, float64, etc.)
    - Expected format hints for bool: "(expected true/false, 1/0, t/f)"
    - Supported types list: "unsupported field type %s (supported: string, int, uint, float, bool)"
  - Better debugging experience with descriptive error messages

### Examples / 예제
- Before: `failed to parse int`
- After: `cannot convert value "abc" to int: invalid syntax`
- Before: `unsupported field type: map`
- After: `unsupported field type map[string]string (supported: string, int, uint, float, bool)`

---

## [v1.11.026] - 2025-10-16

### Security / 보안
- **Request Body Size Limits Added** / **요청 본문 크기 제한 추가**
  - Added `MaxBodySize` option to `Options` struct (default: 10 MB)
  - Added `WithMaxBodySize(size int64)` configuration function
  - Enhanced `BindJSON()` to enforce body size limits using `io.LimitReader`
  - Prevents denial-of-service attacks with large JSON payloads
  - Returns descriptive error message when limit exceeded
  - Configurable per-app via options

### Documentation / 문서화
- Enhanced `BindJSON()` documentation with security considerations
- Added examples for custom body size limits

---

## [v1.11.025] - 2025-10-16

### Refactoring / 리팩토링
- **HTTP Method Registration Duplication Removed** / **HTTP 메서드 등록 중복 제거** (`app.go`)
  - Added `registerRoute(method, pattern, handler)` helper method
  - Refactored GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD methods to use helper
  - Reduced code from ~120 lines to ~50 lines (58% reduction)
  - Improved maintainability with DRY principle
  - All tests pass (199 tests)

---

## [v1.11.024] - 2025-10-16

### Code Quality / 코드 품질
- **Comprehensive Bilingual Comments Added** / **종합 이중 언어 주석 추가**
  - Enhanced `router.go` internal functions:
    - `parsePattern()` - Algorithm description, time complexity O(n), examples
    - `parsePath()` - Process documentation, cleaning logic
    - `match()` - Matching algorithm, parameter extraction details
  - Enhanced `bind.go` internal functions:
    - `bindFormData()` - Extensive reflection usage documentation, supported types, error handling
  - Enhanced `session.go` internal functions:
    - `generateSessionID()` - Security properties (256-bit entropy), collision probability (~1/10^77), fallback strategy
    - `cleanupExpiredSessions()` - Cleanup strategy, thread safety, performance considerations
  - Enhanced `context.go` internal functions:
    - `ClientIP()` - Priority order (X-Forwarded-For, X-Real-IP, RemoteAddr), security considerations, proxy scenarios
  - Enhanced `middleware.go` internal functions:
    - `Recovery()` - Panic recovery mechanism, thread safety, common panic causes, best practices

### Documentation / 문서화
- **CODE_REVIEW_REPORT.md** created with comprehensive analysis:
  - 10 categories of improvements identified
  - Overall assessment: Code quality 8/10, Documentation 7/10, Test coverage 8/10 (79.4%)
  - High Priority (1-4): Comments, refactoring, body limits, error messages
  - Medium Priority (5-7): Route groups, optimization, tests
  - Low Priority (8-10): File splitting, examples, security features

### Tests / 테스트
- All 199 tests pass
- Test coverage: 79.4%

---

## [v1.11.023] - 2025-10-16

### Added / 추가
- Package finalization with comprehensive documentation (USER_MANUAL: 1067 lines, DEVELOPER_GUIDE: 1084 lines) / 종합 문서를 포함한 패키지 완료
- Final README update with all feature sections / 모든 기능 섹션을 포함한 최종 README 업데이트


# CHANGELOG v1.11.x - Web Server Utilities Package

**Package / 패키지**: `websvrutil`
**Focus / 초점**: Extreme simplicity web server utilities / 극도로 간단한 웹 서버 유틸리티

---

## [v1.11.022] - 2025-10-16

### Added / 추가
- **Error Response Helpers** / **에러 응답 헬퍼** (`context.go`)
  - `AbortWithStatus(code int)` - Abort with status code
  - `AbortWithError(code int, message string)` - Abort with error message
  - `AbortWithJSON(code int, obj interface{})` - Abort with JSON
  - `ErrorJSON(code int, message string)` - Standardized JSON error
  - `SuccessJSON(code int, message string, data interface{})` - Standardized JSON success
  - `NotFound()` - Send 404 response
  - `Unauthorized()` - Send 401 response
  - `Forbidden()` - Send 403 response
  - `BadRequest()` - Send 400 response
  - `InternalServerError()` - Send 500 response

### Tests / 테스트
  - 11 test functions + 2 benchmarks (error_test.go)
  - 79.4% coverage

---

## [v1.11.021] - 2025-10-16

### Added / 추가
- **HTTP Method Helpers** / **HTTP 메서드 헬퍼** (`context.go`)
  - `IsGET() bool` - Check if request method is GET
  - `IsPOST() bool` - Check if request method is POST
  - `IsPUT() bool` - Check if request method is PUT
  - `IsPATCH() bool` - Check if request method is PATCH
  - `IsDELETE() bool` - Check if request method is DELETE
  - `IsHEAD() bool` - Check if request method is HEAD
  - `IsOPTIONS() bool` - Check if request method is OPTIONS
  - Simple boolean check methods for common HTTP methods
  - No error handling needed, just true/false return

- **Request Type Helpers** / **요청 타입 헬퍼** (`context.go`)
  - `IsAjax() bool` - Check if request is AJAX (XMLHttpRequest)
  - Checks X-Requested-With header for "XMLHttpRequest"
  - `IsWebSocket() bool` - Check if request is WebSocket upgrade
  - Checks Upgrade header for "websocket"

- **Content Negotiation Helpers** / **콘텐츠 협상 헬퍼** (`context.go`)
  - `AcceptsJSON() bool` - Check if client accepts JSON responses
  - `AcceptsHTML() bool` - Check if client accepts HTML responses
  - `AcceptsXML() bool` - Check if client accepts XML responses
  - Checks Accept header for content type preferences
  - Supports wildcard (*/*) and content type lists
  - `containsContentType(accept, contentType string) bool` - Helper method

- **Comprehensive Tests** / **종합 테스트** (`method_test.go`)
  - `TestContextIsGET` - Test GET method check
  - `TestContextIsPOST` - Test POST method check
  - `TestContextIsPUT` - Test PUT method check
  - `TestContextIsPATCH` - Test PATCH method check
  - `TestContextIsDELETE` - Test DELETE method check
  - `TestContextIsHEAD` - Test HEAD method check
  - `TestContextIsOPTIONS` - Test OPTIONS method check
  - `TestContextIsAjax` - Test AJAX detection
  - `TestContextIsWebSocket` - Test WebSocket detection
  - `TestContextAcceptsJSON` - Test JSON acceptance (6 subtests)
  - `TestContextAcceptsHTML` - Test HTML acceptance (6 subtests)
  - `TestContextAcceptsXML` - Test XML acceptance (7 subtests)
  - `TestContextContainsContentType` - Test content type detection (5 subtests)
  - `BenchmarkIsGET` - Benchmark GET check
  - `BenchmarkIsAjax` - Benchmark AJAX check
  - `BenchmarkAcceptsJSON` - Benchmark JSON acceptance
  - **Total: 13 test functions (25 subtests) + 3 benchmarks** for HTTP method helpers

### Documentation / 문서
- **README.md**
  - Updated version to v1.11.021
  - Added "HTTP Method Helpers" section
  - Listed all 12 new helper methods
  - Added before "Header Helpers" section

### Performance / 성능
- **Test Coverage**: 79.2% of statements / 구문의 79.2%
- **Method Checks**: Direct string comparison, no overhead
- **Content Type Detection**: Simple substring search algorithm

---

## [v1.11.020] - 2025-10-16

### Added / 추가
- **Cookie Helpers Enhancement** / **쿠키 헬퍼 향상** (`context.go`)
  - `CookieValue(name string) string` - Get cookie value as string
  - Returns empty string if cookie not found (no error handling needed)
  - Convenience method for simple cookie value retrieval
  - `SetCookieAdvanced(opts CookieOptions)` - Set cookie with advanced options
  - Full control over cookie attributes (Path, Domain, MaxAge, Secure, HttpOnly, SameSite)
  - Default path "/" if not provided
  - `CookieOptions` struct with 8 configurable fields

- **Cookie Configuration** / **쿠키 설정**
  - `CookieOptions` type for advanced cookie configuration
  - Fields: Name, Value, Path, Domain, MaxAge, Secure, HttpOnly, SameSite
  - MaxAge: 0 for session cookies, -1 to delete
  - Compatible with existing Cookie, SetCookie, DeleteCookie methods

- **Comprehensive Tests** / **종합 테스트** (`cookie_test.go`)
  - `TestContextCookie` - Test Cookie method
  - `TestContextCookieValue` - Test CookieValue convenience method
  - `TestContextSetCookieExisting` - Test existing SetCookie method
  - `TestContextSetCookieAdvanced` - Test SetCookieAdvanced with all options
  - `TestContextSetCookieAdvancedDefaultPath` - Test default path behavior
  - `TestContextDeleteCookieExisting` - Test existing DeleteCookie method
  - `TestContextMultipleCookies` - Test setting multiple cookies
  - `TestCookieOptions` - Test CookieOptions struct
  - `BenchmarkContextSetCookieExisting` - Benchmark SetCookie
  - `BenchmarkContextCookieValue` - Benchmark CookieValue
  - `BenchmarkContextSetCookieAdvanced` - Benchmark SetCookieAdvanced
  - **Total: 8 test functions + 3 benchmarks** for cookie helpers

### Changed / 변경
- **Enhanced Cookie Helpers** / **향상된 쿠키 헬퍼**
  - Added convenience methods while maintaining backward compatibility
  - Existing methods (Cookie, SetCookie, DeleteCookie) unchanged
  - New methods work alongside existing methods

### Documentation / 문서
- **README.md**
  - Updated version to v1.11.020
  - Updated "Cookie Helpers" section with new methods
  - Added note "Enhanced in v1.11.020"
  - Listed CookieValue and SetCookieAdvanced methods

### Performance / 성능
- **Test Coverage**: 78.8% of statements / 구문의 78.8%
- **CookieValue**: No error handling overhead, direct string return
- **SetCookieAdvanced**: Efficient single-call configuration

---

## [v1.11.019] - 2025-10-16

### Added / 추가
- **Session Management System** / **세션 관리 시스템** (`session.go`)
  - `NewSessionStore(opts SessionOptions) *SessionStore` - Create session store with custom options
  - `SessionStore.Get(r *http.Request) (*Session, error)` - Get or create session from request
  - `SessionStore.New() *Session` - Create new session with unique ID
  - `SessionStore.Save(w http.ResponseWriter, session *Session)` - Save session and set cookie
  - `SessionStore.Destroy(w http.ResponseWriter, r *http.Request) error` - Destroy session and clear cookie
  - `SessionStore.Count() int` - Get active session count

- **Session Data Storage** / **세션 데이터 저장소**
  - `Session.Set(key string, value interface{})` - Store value in session
  - `Session.Get(key string) (interface{}, bool)` - Retrieve value from session
  - `Session.GetString(key string) string` - Get string value (type-safe)
  - `Session.GetInt(key string) int` - Get int value (type-safe)
  - `Session.GetBool(key string) bool` - Get bool value (type-safe)
  - `Session.Delete(key string)` - Remove value from session
  - `Session.Clear()` - Clear all session values
  - Thread-safe with sync.RWMutex protection

- **Session Configuration** / **세션 설정**
  - `SessionOptions` struct with 8 configurable fields
  - `DefaultSessionOptions()` function for quick setup
  - CookieName, MaxAge, Secure, HttpOnly, SameSite, Path, Domain, CleanupTime
  - Default: 24h expiration, HttpOnly=true, SameSite=Lax

- **Advanced Features** / **고급 기능**
  - Automatic session expiration with configurable MaxAge
  - Background cleanup goroutine for expired sessions
  - Cryptographically secure session IDs using crypto/rand
  - Base64 URL-safe encoding for session IDs
  - Fallback to timestamp-based IDs if crypto/rand fails

- **Comprehensive Tests** / **종합 테스트** (`session_test.go`)
  - `TestNewSessionStore` - Test session store creation
  - `TestSessionStoreNew` - Test new session creation
  - `TestSessionSetGet` - Test session set and get operations
  - `TestSessionGetTyped` - Test type-safe getter methods
  - `TestSessionDelete` - Test value deletion
  - `TestSessionClear` - Test clearing all values
  - `TestSessionStoreGetExisting` - Test getting existing session from cookie
  - `TestSessionStoreGetNew` - Test creating new session when none exists
  - `TestSessionStoreSave` - Test saving session and setting cookie
  - `TestSessionStoreDestroy` - Test destroying session
  - `TestSessionExpiration` - Test session expiration behavior
  - `TestSessionCleanup` - Test automatic cleanup of expired sessions
  - `TestSessionConcurrency` - Test concurrent access to session
  - `BenchmarkSessionSet` - Benchmark Set operation
  - `BenchmarkSessionGet` - Benchmark Get operation
  - `BenchmarkSessionStoreNew` - Benchmark session creation
  - **Total: 12 test functions + 3 benchmarks** for session management

### Documentation / 문서
- **README.md**
  - Updated version to v1.11.019
  - Added "Session Management" section with full API documentation
  - SessionStore methods, Session methods, SessionOptions reference
  - Added comprehensive Session Management example in Quick Start
  - Example shows login, protected routes, logout, and session info

### Performance / 성능
- **Test Coverage**: 80.6% of statements / 구문의 80.6%
- **Session Operations**: Thread-safe with minimal locking overhead
- **Cleanup**: Background goroutine with configurable interval (default: 5m)

---

## [v1.11.018] - 2025-10-16

### Added / 추가
- **Graceful Shutdown with Signal Handling** / **시그널 처리와 함께 정상 종료** (`app.go`)
  - `RunWithGracefulShutdown(addr string, timeout time.Duration) error` - Run server with automatic signal handling
  - Automatically listens for SIGINT and SIGTERM signals
  - Gracefully shuts down server when signal received
  - Configurable shutdown timeout
  - Simplified alternative to manual signal handling

- **Signal Handling** / **시그널 처리**
  - Automatic handling of SIGINT (Ctrl+C) and SIGTERM signals
  - Goroutine-based server startup for non-blocking signal handling
  - Select-based multiplexing for server errors and shutdown signals
  - Context-based timeout for graceful shutdown

- **Comprehensive Tests** / **종합 테스트** (`shutdown_test.go`)
  - `TestShutdownWithoutRun` - Test shutdown without running server
  - `TestShutdown` - Test basic shutdown functionality
  - `TestShutdownNotRunning` - Test error when server not running
  - `TestShutdownWithTimeout` - Test shutdown timeout behavior
  - `TestShutdownIdempotent` - Test multiple shutdown calls
  - `TestRunWithGracefulShutdown` - Test signal handling (manual test)
  - `TestGracefulShutdownWithActiveConnections` - Test connection draining (manual test)
  - `BenchmarkShutdown` - Benchmark shutdown operation
  - **Total: 7 test functions + 1 benchmark** for graceful shutdown

### Changed / 변경
- **Added Imports to app.go** / **app.go에 임포트 추가**
  - Added `os`, `os/signal`, `syscall`, `time` imports for signal handling

### Documentation / 문서
- **README.md**
  - Updated version to v1.11.018
  - Added `RunWithGracefulShutdown` method to App Struct section
  - Added "Graceful Shutdown (Simple)" example showing the new method
  - Example demonstrates automatic signal handling with 5-second timeout

### Performance / 성능
- **Test Coverage**: 79.5% of statements / 구문의 79.5%
- **RunWithGracefulShutdown**: Minimal overhead, uses goroutine and channel-based signaling

---

## [v1.11.017] - 2025-10-16

### Added / 추가
- **Context Storage Enhancement** / **컨텍스트 저장소 향상** (`context.go`)
  - `GetInt64(key string) int64` - Get int64 value from context
  - `GetFloat64(key string) float64` - Get float64 value from context
  - `GetStringSlice(key string) []string` - Get string slice from context
  - `GetStringMap(key string) map[string]interface{}` - Get string map from context
  - `Exists(key string) bool` - Check if key exists in context
  - `Delete(key string)` - Remove value from context
  - `Keys() []string` - Get all keys from context

- **Enhanced Type Safety** / **향상된 타입 안전성**
  - Type-safe getter methods for common data types
  - Zero-value returns for non-existent or wrong-type keys
  - Thread-safe operations with sync.RWMutex protection

- **Key Management** / **키 관리**
  - Exists() method for key existence check
  - Delete() method for value removal
  - Keys() method to list all stored keys
  - Support for nil values in storage

- **Comprehensive Tests** / **종합 테스트** (`storage_test.go`)
  - `TestContextStorageSetGet` - Test Set and Get operations
  - `TestContextStorageMustGet` - Test MustGet method
  - `TestContextStorageMustGetPanic` - Test MustGet panic on missing key
  - `TestContextGetString` - Test string type getter
  - `TestContextGetInt` - Test int type getter
  - `TestContextGetBool` - Test bool type getter
  - `TestContextGetInt64` - Test int64 type getter
  - `TestContextGetFloat64` - Test float64 type getter
  - `TestContextGetStringSlice` - Test string slice getter
  - `TestContextGetStringMap` - Test string map getter
  - `TestContextExists` - Test key existence check
  - `TestContextDelete` - Test value deletion
  - `TestContextKeys` - Test getting all keys
  - `TestContextStorageConcurrency` - Test thread-safety
  - `TestContextStorageTypes` - Test various data types
  - `BenchmarkContextSet` - Benchmark Set operation
  - `BenchmarkContextGet` - Benchmark Get operation
  - `BenchmarkContextGetString` - Benchmark GetString operation
  - **Total: 15 test functions + 3 benchmarks** for context storage

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.017
- Bumped version to v1.11.017 in `cfg/app.yaml`
- Updated `README.md` with enhanced Context Storage section
- Added new file: `storage_test.go`

### Testing Coverage / 테스트 커버리지
- **15 new test functions + 3 benchmarks** for context storage
- **Total: 205+ test functions**, **Total: 46 benchmark functions**
- **79.2% test coverage** - All tests passing ✅

### Notes / 참고사항
- All new methods are thread-safe with RWMutex protection
- Type-safe getters return zero values for missing/wrong-type keys
- Exists() checks for key presence (even if value is nil)
- Delete() safely removes keys without panic
- Keys() returns sorted slice for predictable iteration
- Enhanced from basic Set/Get/MustGet to full-featured storage system
- Supports all common Go types: string, int, int64, float64, bool, slices, maps
- Concurrency tests verify thread-safety under high load
- Next: v1.11.018+ will continue Phase 5 (Server Management) development

---

## [v1.11.016] - 2025-10-16

### Added / 추가
- **Static File Serving** / **정적 파일 서빙** (`app.go`, `context.go`)
  - `Static(prefix, dir string) *App` (App method) - Serve static files from directory
  - `File(filepath string) error` (Context method) - Send file response to client
  - `FileAttachment(filepath, filename string) error` (Context method) - Send file as downloadable attachment
  - Automatic Content-Type detection based on file extension
  - Support for multiple static directories with different prefixes
  - Built on standard library http.FileServer and http.ServeFile

- **File Serving Features** / **파일 서빙 기능**
  - Static directory serving with URL prefix stripping
  - Direct file serving with automatic MIME type detection
  - Download attachment with custom filename
  - Wildcard route registration for static files
  - Subdirectory support in static file serving

- **Comprehensive Tests** / **종합 테스트** (`static_test.go`)
  - `TestStaticFileServing` - Test serving files from directory
  - `TestStaticNotFound` - Test 404 for non-existent files
  - `TestStaticMultiplePrefixes` - Test multiple static directories
  - `TestFile` - Test File method with automatic Content-Type
  - `TestFileNotFound` - Test 404 for non-existent file
  - `TestFileAttachment` - Test file download with Content-Disposition
  - `TestFileAttachmentDifferentTypes` - Test various file types (PDF, ZIP, TXT, JPG)
  - `BenchmarkStaticFileServing` - Benchmark static file serving
  - `BenchmarkFile` - Benchmark File method
  - **Total: 7 test functions + 2 benchmarks** for static file serving

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.016
- Bumped version to v1.11.016 in `cfg/app.yaml`
- Updated `README.md` with Static File Serving section and examples
- Added new file: `static_test.go`

### Testing Coverage / 테스트 커버리지
- **7 new test functions + 2 benchmarks** for static file serving
- **Total: 190+ test functions**, **Total: 43 benchmark functions**
- **78.6% test coverage** - All tests passing ✅

### Notes / 참고사항
- Static() method uses http.FileServer for efficient file serving
- File() and FileAttachment() use http.ServeFile with automatic MIME detection
- Multiple static directories can be served with different URL prefixes
- FileAttachment() sets Content-Disposition header for download
- All file operations support subdirectories
- Examples include multiple static directories, specific files, and dynamic file serving
- Next: v1.11.017+ will add context storage and middleware enhancements

---

## [v1.11.015] - 2025-10-16

### Added / 추가
- **File Upload Support** / **파일 업로드 지원** (`context.go`)
  - `FormFile(name string) (*multipart.FileHeader, error)` - Get uploaded file by form field name
  - `MultipartForm() (*multipart.Form, error)` - Get parsed multipart form (files + fields)
  - `SaveUploadedFile(file *multipart.FileHeader, dst string) error` - Save uploaded file to destination
  - Automatic file size limit enforcement using `MaxUploadSize` option
  - Default 32 MB upload limit (configurable)

- **MaxUploadSize Option** / **MaxUploadSize 옵션** (`options.go`)
  - `WithMaxUploadSize(size int64)` - Set maximum file upload size in bytes
  - Default: 32 MB (32 << 20 bytes)
  - Configurable per App instance
  - Enforced in MultipartForm() method

- **Context App Reference** / **Context 앱 참조** (`context.go`)
  - Added `app *App` field to Context struct
  - Enables Context to access App options (e.g., MaxUploadSize)
  - Thread-safe access to app configuration

- **Comprehensive Tests** / **종합 테스트** (`upload_test.go`)
  - `TestFormFile` - Test getting uploaded file
  - `TestFormFileNotFound` - Test non-existent file field
  - `TestMultipartForm` - Test getting multipart form with files and fields
  - `TestMultipartFormMultipleFiles` - Test multiple file uploads
  - `TestSaveUploadedFile` - Test saving uploaded file to disk
  - `TestSaveUploadedFileLargeFile` - Test saving 1MB file
  - `TestSaveUploadedFileError` - Test error handling (invalid path)
  - `TestMultipartFormWithMaxUploadSize` - Test custom max upload size
  - `BenchmarkFormFile` - Benchmark file retrieval
  - `BenchmarkSaveUploadedFile` - Benchmark file saving (1KB file)
  - **Total: 8 test functions + 2 benchmarks** for file upload

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.015
- Bumped version to v1.11.015 in `cfg/app.yaml`
- Updated `README.md` with file upload section and examples
- Added new file: `upload_test.go`
- Added imports to `context.go`: `io`, `mime/multipart`, `os`
- Updated Options table in README with MaxUploadSize option

### Testing Coverage / 테스트 커버리지
- **8 new test functions + 2 benchmarks** for file upload
- **Total: 183+ test functions**, **Total: 41 benchmark functions**
- **78.4% test coverage** - All tests passing ✅

### Notes / 참고사항
- File upload supports single and multiple files
- MultipartForm() returns both files and form fields
- SaveUploadedFile() automatically creates parent directories if needed
- MaxUploadSize is configurable per App instance (default 32 MB)
- All file operations are thread-safe
- Examples include single file, multiple files, and files with form data
- Next: v1.11.016+ will add static file serving

---

## [v1.11.014] - 2025-10-16

### Added / 추가
- **Cookie Helpers** / **쿠키 헬퍼** (`context.go`)
  - `Cookie(name string) (*http.Cookie, error)` - Get cookie by name
  - `SetCookie(cookie *http.Cookie)` - Set response cookie with full options
  - `DeleteCookie(name, path string)` - Delete cookie by setting MaxAge to -1
  - `GetCookie(name string) string` - Convenience method to get cookie value

- **Header Helpers** / **헤더 헬퍼** (`context.go`)
  - `AddHeader(key, value string)` - Add header value (appends if exists)
  - `GetHeader(key string) string` - Get request header (alias for Header())
  - `GetHeaders(key string) []string` - Get all values for a header key
  - `HeaderExists(key string) bool` - Check if request header exists
  - `ContentType() string` - Get Content-Type header
  - `UserAgent() string` - Get User-Agent header
  - `Referer() string` - Get Referer header
  - `ClientIP() string` - Get client IP address with X-Forwarded-For, X-Real-IP support

- **Comprehensive Tests** / **종합 테스트** (`cookie_test.go`)
  - `TestCookie` - Test getting a cookie
  - `TestCookieNotFound` - Test non-existent cookie
  - `TestSetCookie` - Test setting a cookie
  - `TestDeleteCookie` - Test deleting a cookie
  - `TestGetCookie` - Test convenience method
  - `TestAddHeader` - Test adding multiple header values
  - `TestGetHeader` - Test getting request header
  - `TestGetHeaders` - Test getting multiple header values
  - `TestHeaderExists` - Test header existence check
  - `TestContentType` - Test Content-Type helper
  - `TestUserAgent` - Test User-Agent helper
  - `TestReferer` - Test Referer helper
  - `TestClientIP` - Test client IP extraction (4 sub-tests)
  - **Total: 13 test functions** for cookie and header helpers

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.014
- Bumped version to v1.11.014 in `cfg/app.yaml`
- Updated `README.md` with cookie and header helpers documentation
- Added new file: `cookie_test.go`

### Testing Coverage / 테스트 커버리지
- **13 new test functions** for cookie and header helpers
- **Total: 175+ test functions**, **Total: 39 benchmark functions**
- **78.0% test coverage** - All tests passing ✅

### Notes / 참고사항
- Cookie helpers provide easy cookie management with full HTTP cookie options
- Header helpers include convenience methods for common headers (Content-Type, User-Agent, Referer)
- ClientIP() intelligently extracts IP from X-Forwarded-For, X-Real-IP, or RemoteAddr
- All methods are thread-safe and integrated into Context
- Next: v1.11.015+ will add file upload support

---

## [v1.11.013] - 2025-10-16

### Added / 추가
- **Request Binding System** / **요청 바인딩 시스템** (`context.go`, `bind.go`)
  - `Bind(obj interface{}) error` - Auto bind request data based on Content-Type
  - `BindJSON(obj interface{}) error` - Bind JSON request body to struct
  - `BindForm(obj interface{}) error` - Bind form data to struct with `form` tags
  - `BindQuery(obj interface{}) error` - Bind query parameters to struct with `form` tags
  - Supports automatic Content-Type detection (application/json, application/x-www-form-urlencoded, multipart/form-data)
  - Reflection-based struct field mapping using `form` tags
  - Type conversion support: string, int, uint, float, bool

- **Form Data Binding Helper** / **폼 데이터 바인딩 헬퍼** (`bind.go`)
  - `bindFormData(obj interface{}, values url.Values) error` - Bind URL values to struct
  - `setFieldValue(field reflect.Value, value string) error` - Set struct field value by type
  - Supports int (int, int8, int16, int32, int64)
  - Supports uint (uint, uint8, uint16, uint32, uint64)
  - Supports float (float32, float64)
  - Supports bool
  - Supports string
  - Error handling for invalid type conversions

- **Comprehensive Tests** / **종합 테스트** (`bind_test.go`)
  - `TestBindJSON` - Test JSON binding with valid and invalid JSON
  - `TestBindForm` - Test form data binding
  - `TestBindQuery` - Test query parameter binding
  - `TestBind` - Test automatic binding based on Content-Type
  - `TestBindFormData` - Test helper function with all data types
  - `TestBindFormDataError` - Test error cases (not pointer, not struct)
  - `BenchmarkBindJSON` - Benchmark JSON binding performance
  - `BenchmarkBindForm` - Benchmark form binding performance
  - **Total: 6 test functions + 2 benchmarks** for request binding

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.013
- Bumped version to v1.11.013 in `cfg/app.yaml`
- Updated `README.md` with request binding documentation
- Added new files: `bind.go`, `bind_test.go`

### Testing Coverage / 테스트 커버리지
- **6 new test functions + 2 benchmarks** for request binding
- **Total: 162+ test functions**, **Total: 39 benchmark functions**
- **77.5% test coverage** - All tests passing ✅

### Notes / 참고사항
- Request binding uses reflection for automatic struct mapping
- Struct fields should use `form` tags to specify form/query field names
- JSON binding uses standard `json` tags
- Supports nested structs for complex data structures
- Type conversion errors are properly handled and reported
- Next: v1.11.014+ will add cookie & header helpers or file upload

---

## [v1.11.012] - 2025-10-16

### Added / 추가
- **Hot Reload System** / **핫 리로드 시스템** (`template.go`)
  - `EnableAutoReload()` - Enable automatic template reloading when files change
  - `DisableAutoReload()` - Disable automatic template reloading
  - `IsAutoReloadEnabled()` - Check if auto-reload is currently enabled
  - Polling-based file watcher (checks every 1 second)
  - Watches both template directory and layouts directory
  - Automatically reloads templates and layouts when files are modified
  - Console logging for reload events: "[Template Hot Reload] Detected changes, reloading templates..."
  - Thread-safe auto-reload management with stopChan
  - New fields in TemplateEngine: `autoReload bool`, `stopChan chan struct{}`

- **App Integration** / **앱 통합** (`app.go`)
  - Auto-enable hot reload when `WithAutoReload(true)` option is set
  - Logs: "[Template Hot Reload] Auto-reload enabled for templates and layouts"
  - References correct `options.EnableAutoReload` field

- **Helper Function** / **헬퍼 함수** (`template.go`)
  - `isTemplateFile(path string) bool` - Check if file is a template based on extension
  - Supports .html, .htm, .tmpl extensions

- **Comprehensive Tests** / **종합 테스트** (`template_test.go`)
  - `TestEnableAutoReload` - Test enabling auto-reload
  - `TestDisableAutoReload` - Test disabling auto-reload
  - `TestIsAutoReloadEnabled` - Test checking auto-reload status
  - **Total: 24 test functions** for complete template system

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.012
- Bumped version to v1.11.012 in `cfg/app.yaml`
- Updated `README.md` with hot reload documentation
- Modified TemplateEngine struct to include auto-reload fields

### Testing Coverage / 테스트 커버리지
- **24 test functions + 4 benchmarks** for complete template system with hot reload
- **Total: 156+ test functions**, **Total: 37 benchmark functions**
- **77.2% test coverage** - All tests passing ✅

### Notes / 참고사항
- Hot reload uses polling (1 second interval) instead of filesystem events (no external dependencies)
- Useful during development - automatically reloads templates when modified
- Enable with `WithAutoReload(true)` option or call `EnableAutoReload()` manually
- Disable in production for better performance
- Next: v1.11.013+ will add additional template features or move to Phase 4 (Advanced Features)

---

## [v1.11.011] - 2025-10-16

### Added / 추가
- **Layout System** / **레이아웃 시스템** (`template.go`)
  - `SetLayoutDir(dir)` - Set layout directory (default: "views/layouts")
  - `LoadLayout(name)` - Load single layout file
  - `LoadAllLayouts()` - Load all layouts from layout directory recursively
  - `RenderWithLayout(w, layoutName, templateName, data)` - Render template with layout
  - `HasLayout(name)` - Check if layout exists
  - `ListLayouts()` - List all loaded layouts
  - Layout templates use `{{template "content" .}}` to embed content
  - Separate storage for layouts (`layouts map[string]*template.Template`)
  - Auto-load layouts on app creation when TemplateDir is set

- **Built-in Template Functions** / **내장 템플릿 함수** (`template.go`)
  - **String functions** (13개):
    - `upper`, `lower`, `title` - Case conversion
    - `trim`, `trimPrefix`, `trimSuffix` - Whitespace/prefix/suffix removal
    - `replace` - String replacement
    - `contains`, `hasPrefix`, `hasSuffix` - String checking
    - `split`, `join` - String splitting/joining
    - `repeat` - String repetition
  - **Date/Time functions** (5개):
    - `now` - Current time
    - `formatDate` - Format time with layout
    - `formatDateSimple` - Simple date format (2006-01-02)
    - `formatDateTime` - DateTime format (2006-01-02 15:04:05)
    - `formatTime` - Time format (15:04:05)
  - **URL functions** (2개):
    - `urlEncode` - URL encode string
    - `urlDecode` - URL decode string
  - **Safe HTML functions** (3개):
    - `safeHTML` - Mark HTML as safe (template.HTML)
    - `safeURL` - Mark URL as safe (template.URL)
    - `safeJS` - Mark JavaScript as safe (template.JS)
  - **Utility functions** (2개):
    - `default` - Return default value if empty
    - `len` - Return length of string/slice/map
  - **Total: 26+ built-in functions**

- **Context Layout Rendering** / **Context 레이아웃 렌더링** (`context.go`)
  - `ctx.RenderWithLayout(code, layoutName, templateName, data)` - Render template with layout
  - Automatic content-type and status code setting
  - Access to template engine through request context

- **App Integration** / **앱 통합** (`app.go`)
  - Auto-load layouts on app creation (after loading templates)
  - Logs warning if layout loading fails but doesn't stop app

- **Comprehensive Tests** / **종합 테스트** (`template_test.go`)
  - 7 new test functions + 2 benchmarks for layouts and built-in functions
  - Tests for layout loading and rendering
  - Tests for built-in functions (Upper, Lower, SafeHTML)
  - Tests for SetLayoutDir
  - Tests for ListLayouts and HasLayout
  - Benchmark for built-in functions
  - Benchmark for RenderWithLayout
  - **Total: 21 test functions + 4 benchmarks** for template system

### Fixed / 수정
- Fixed `RenderWithLayout` to use `ExecuteTemplate()` instead of `Execute()`
  - Was: `layoutClone.Execute(w, data)` (executed content template)
  - Now: `layoutClone.ExecuteTemplate(w, layoutName, data)` (executes layout template)
  - This ensures the layout template is properly executed with the content template embedded

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.011
- Bumped version to v1.11.011 in `cfg/app.yaml`
- Updated `README.md` with layout system and built-in functions documentation
- Modified `NewTemplateEngine()` to call `addBuiltinFuncs()` on creation
- Template engine now has both `templates` and `layouts` maps
- Template engine now has `layoutDir` field (default: "views/layouts")

### Testing Coverage / 테스트 커버리지
- **21 test functions + 4 benchmarks** for complete template system
- **Total: 153+ test functions**, **Total: 37 benchmark functions**
- **80.4% test coverage** - All tests passing ✅

### Notes / 참고사항
- Built-in functions are automatically added to all templates and layouts
- Layout system uses Go's `template.Clone()` and `AddParseTree()` for composition
- Layout directory defaults to "views/layouts" but can be customized
- Both templates and layouts support nested directories
- Next: v1.11.012+ will add hot reload and additional template features

---

## [v1.11.010] - 2025-10-16

### Added / 추가
- **Template Engine Core** / **템플릿 엔진 핵심** (`template.go`)
  - `TemplateEngine` struct with thread-safe template management
  - Automatic template loading from directory (`LoadAll()`)
  - Single template loading (`Load()`) and glob pattern loading (`LoadGlob()`)
  - Template rendering to io.Writer (`Render()`)
  - Custom template functions support (`AddFunc()`, `AddFuncs()`)
  - Custom delimiter support (`SetDelimiters()`)
  - Template existence check (`Has()`)
  - List all loaded templates (`List()`)
  - Clear all templates (`Clear()`)
  - Support for .html, .htm, .tmpl files
  - Nested directory support for recursive template loading

- **App Template Integration** / **앱 템플릿 통합** (`app.go`)
  - Changed `templates` field from `interface{}` to `*TemplateEngine`
  - Auto-initialize template engine when `TemplateDir` option is set
  - Auto-load all templates on app creation
  - `TemplateEngine()` - Get template engine instance
  - `LoadTemplate(name)` - Load single template
  - `LoadTemplates(pattern)` - Load templates by pattern
  - `ReloadTemplates()` - Reload all templates
  - `AddTemplateFunc(name, fn)` - Add custom template function
  - `AddTemplateFuncs(funcs)` - Add multiple custom functions
  - Store app in request context for template access

- **Context Template Rendering** / **Context 템플릿 렌더링** (`context.go`)
  - `ctx.Render(code, name, data)` - Render template file with data
  - Automatic content-type and status code setting
  - Access to template engine through request context

- **Comprehensive Tests** / **종합 테스트** (`template_test.go`)
  - 14 test functions + 2 benchmarks
  - Tests for all template engine methods
  - Tests for custom functions and delimiters
  - Tests for glob loading and recursive loading
  - Tests for rendering with data
  - Benchmark for Load and Render operations

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.010
- Bumped version to v1.11.010 in `cfg/app.yaml`
- Updated `README.md` with comprehensive template system documentation
- Updated `app.go` ServeHTTP to store app in request context

### Testing Coverage / 테스트 커버리지
- **14 new test functions + 2 benchmarks** for template engine
- **Total: 146+ test functions**, **Total: 33 benchmark functions**
- **81.6% test coverage** - All tests passing ✅

### Notes / 참고사항
- **Phase 3 (Template System) started!** / Phase 3 (템플릿 시스템) 시작!
- Template engine is automatically initialized when `TemplateDir` option is set
- Templates are auto-loaded on app creation for convenience
- Thread-safe template caching with sync.RWMutex
- Support for custom template functions and delimiters
- Next: v1.11.011+ will add layout system, hot reload, and built-in template functions

---

## [v1.11.009] - 2025-10-16

### Added / 추가
- Added 5 new final middleware to `middleware.go` / middleware.go에 5개의 새로운 최종 미들웨어 추가
  - **BodyLimit Middleware / 본문 제한 미들웨어**:
    - `BodyLimit(maxBytes)` - Limits maximum request body size / 최대 요청 본문 크기 제한
    - `BodyLimitWithConfig(config)` - Custom body limit configuration / 커스텀 본문 제한 설정
    - Default limit: 10MB / 기본 제한: 10MB
    - Uses http.MaxBytesReader for efficient limiting / http.MaxBytesReader를 사용한 효율적인 제한
    - Prevents memory exhaustion attacks / 메모리 고갈 공격 방지
  - **Static Middleware / 정적 파일 미들웨어**:
    - `Static(root)` - Serves static files from directory / 디렉토리에서 정적 파일 제공
    - `StaticWithConfig(config)` - Custom static file configuration / 커스텀 정적 파일 설정
    - Automatic index.html serving / 자동 index.html 제공
    - Optional directory browsing / 선택적 디렉토리 탐색
    - Falls through to next handler if file not found / 파일을 찾을 수 없으면 다음 핸들러로 전달
  - **Redirect Middleware / 리디렉션 미들웨어**:
    - `Redirect(to)` - Redirects all requests to URL / 모든 요청을 URL로 리디렉션
    - `RedirectWithConfig(config)` - Custom redirect configuration / 커스텀 리디렉션 설정
    - Default: 301 Moved Permanently / 기본값: 301 Moved Permanently
    - Configurable status code / 설정 가능한 상태 코드
  - **HTTPSRedirect Middleware / HTTPS 리디렉션 미들웨어**:
    - `HTTPSRedirect()` - Redirects HTTP to HTTPS / HTTP를 HTTPS로 리디렉션
    - Detects protocol from TLS and X-Forwarded-Proto header / TLS 및 X-Forwarded-Proto 헤더에서 프로토콜 감지
    - 301 Permanent redirect / 301 영구 리디렉션
  - **WWWRedirect Middleware / WWW 리디렉션 미들웨어**:
    - `WWWRedirect(addWWW)` - Adds or removes www prefix / www 접두사 추가 또는 제거
    - Supports both HTTP and HTTPS / HTTP 및 HTTPS 모두 지원
    - Protocol-aware redirection / 프로토콜 인식 리디렉션
- Created comprehensive tests in `middleware_test.go`: 11 tests + 5 benchmarks
- Updated `README.md` with new middleware documentation
- **Total 14 middleware now available** / 총 14개 미들웨어 사용 가능

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.009
- Bumped version to v1.11.009 in `cfg/app.yaml`
- Fixed middleware imports: Added `os` to imports
- Fixed `middleware_test.go` imports: Added `io` and `os`
- Fixed WWWRedirect and HTTPSRedirect to use `r.URL.Path` instead of `r.RequestURI` for proper URL construction

### Testing Coverage / 테스트 커버리지
- **11 new tests + 5 benchmarks** for new middleware
- **Total: 132+ test functions**, **Total: 31 benchmark functions**
- **85.3% test coverage** - All tests passing ✅

### Notes / 참고사항
- **Phase 2 (Middleware System) complete!** / Phase 2 (미들웨어 시스템) 완료!
- Total 14 middleware available: Recovery, Logger, CORS, RequestID, Timeout, BasicAuth, RateLimiter, Compression, SecureHeaders, BodyLimit, Static, Redirect, HTTPSRedirect, WWWRedirect
- Comprehensive middleware suite for production-ready web servers / 프로덕션 준비 웹 서버를 위한 포괄적인 미들웨어 제품군
- Next: v1.11.010+ will focus on Phase 3 (Template System) / 다음: v1.11.010+는 Phase 3 (템플릿 시스템)에 집중

---

## [v1.11.008] - 2025-10-16

### Added / 추가
- Added 3 new advanced middleware to `middleware.go` / middleware.go에 3개의 새로운 고급 미들웨어 추가
  - **RateLimiter Middleware**: IP-based rate limiting with token bucket algorithm
  - **Compression Middleware**: Gzip compression for HTTP responses
  - **SecureHeaders Middleware**: Security-related HTTP headers (X-Frame-Options, CSP, HSTS, etc.)
- Created comprehensive tests in `middleware_test.go`: 7 tests + 3 benchmarks
- Updated `README.md` with new middleware documentation
- Total 9 middleware now available

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.008
- Bumped version to v1.11.008 in `cfg/app.yaml`

### Testing Coverage / 테스트 커버리지
- **7 new tests + 3 benchmarks** for new middleware
- **Total: 121+ test functions**, **Total: 26 benchmark functions**
- **85.0% test coverage** - All tests passing ✅

### Notes / 참고사항
- Phase 2 (Middleware System) nearly complete!
- Total 9 middleware available: Recovery, Logger, CORS, RequestID, Timeout, BasicAuth, RateLimiter, Compression, SecureHeaders
- Next: v1.11.009-010 may add final middleware features or move to Phase 3 (Template System)

---

## [v1.11.007] - 2025-10-16

### Added / 추가
- Added 3 new middleware to `middleware.go` / middleware.go에 3개의 새로운 미들웨어 추가
  - **RequestID Middleware / 요청 ID 미들웨어**:
    - `RequestID()` - Default RequestID middleware / 기본 요청 ID 미들웨어
    - `RequestIDWithConfig(config)` - Custom RequestID configuration / 커스텀 요청 ID 설정
    - Generates unique 16-byte hex request IDs / 고유한 16바이트 16진수 요청 ID 생성
    - Stores ID in context with key "request_id" / "request_id" 키로 컨텍스트에 ID 저장
    - Adds ID to response header (default: X-Request-ID) / 응답 헤더에 ID 추가 (기본: X-Request-ID)
    - Preserves existing request ID if present / 기존 요청 ID가 있으면 보존
    - Customizable header name and ID generator / 커스터마이즈 가능한 헤더 이름 및 ID 생성기
  - **Timeout Middleware / 타임아웃 미들웨어**:
    - `Timeout(duration)` - Default timeout middleware / 기본 타임아웃 미들웨어
    - `TimeoutWithConfig(config)` - Custom timeout configuration / 커스텀 타임아웃 설정
    - Enforces request timeout (default: 30 seconds) / 요청 타임아웃 적용 (기본: 30초)
    - Sends 503 Service Unavailable on timeout / 타임아웃 시 503 Service Unavailable 전송
    - Uses http.TimeoutHandler for implementation / 구현을 위해 http.TimeoutHandler 사용
    - Configurable timeout duration and error message / 설정 가능한 타임아웃 기간 및 에러 메시지
  - **BasicAuth Middleware / Basic 인증 미들웨어**:
    - `BasicAuth(username, password)` - Default BasicAuth middleware / 기본 Basic 인증 미들웨어
    - `BasicAuthWithConfig(config)` - Custom BasicAuth configuration / 커스텀 Basic 인증 설정
    - HTTP Basic Authentication enforcement / HTTP Basic Authentication 적용
    - Constant-time password comparison (secure) / 상수 시간 비밀번호 비교 (보안)
    - Sends 401 Unauthorized with WWW-Authenticate header / WWW-Authenticate 헤더와 함께 401 Unauthorized 전송
    - Customizable realm and validator function / 커스터마이즈 가능한 영역 및 검증자 함수
    - Stores username in context with key "auth_username" / "auth_username" 키로 컨텍스트에 사용자 이름 저장
- Added configuration structs / 설정 구조체 추가
  - `RequestIDConfig` - RequestID middleware configuration / 요청 ID 미들웨어 설정
  - `TimeoutConfig` - Timeout middleware configuration / 타임아웃 미들웨어 설정
  - `BasicAuthConfig` - BasicAuth middleware configuration / Basic 인증 미들웨어 설정
- Added helper function / 헬퍼 함수 추가
  - `generateRequestID()` - Generates random 16-byte hex string / 무작위 16바이트 16진수 문자열 생성
- Updated imports in `middleware.go` / middleware.go의 imports 업데이트
  - Added `context` for context operations / 컨텍스트 작업을 위한 context 추가
  - Added `crypto/rand` for secure random generation / 안전한 무작위 생성을 위한 crypto/rand 추가
  - Added `crypto/subtle` for constant-time comparison / 상수 시간 비교를 위한 crypto/subtle 추가
  - Added `encoding/hex` for hex encoding / 16진수 인코딩을 위한 encoding/hex 추가
- Created comprehensive tests in `middleware_test.go` / middleware_test.go에 포괄적인 테스트 생성
  - 9 new test functions for new middleware / 새 미들웨어를 위한 9개의 새로운 테스트 함수
  - RequestID tests: TestRequestID, TestRequestIDWithExistingID, TestRequestIDWithConfig / 요청 ID 테스트
  - Timeout tests: TestTimeout, TestTimeoutWithConfig / 타임아웃 테스트
  - BasicAuth tests: TestBasicAuth, TestBasicAuthUnauthorized, TestBasicAuthNoCredentials, TestBasicAuthWithConfig / Basic 인증 테스트
  - 3 new benchmark functions / 3개의 새로운 벤치마크 함수
- Updated `README.md` with new middleware documentation / 새 미들웨어 문서로 README.md 업데이트
  - Added RequestID, Timeout, BasicAuth middleware sections / 요청 ID, 타임아웃, Basic 인증 미들웨어 섹션 추가
  - Updated version to v1.11.007 / 버전을 v1.11.007로 업데이트
  - Updated progress status / 진행 상태 업데이트

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.007 / websvrutil.go 버전 상수를 v1.11.007로 업데이트
- Bumped version to v1.11.007 in `cfg/app.yaml` / cfg/app.yaml의 버전을 v1.11.007로 상향

### Technical Details / 기술 세부사항
- **RequestID Middleware Architecture / 요청 ID 미들웨어 아키텍처**:
  - Uses crypto/rand for cryptographically secure random IDs / 암호학적으로 안전한 무작위 ID를 위해 crypto/rand 사용
  - 16-byte random = 32-character hex string / 16바이트 무작위 = 32자 16진수 문자열
  - Checks for existing ID in request header / 요청 헤더에서 기존 ID 확인
  - Stores ID in both context and response header / 컨텍스트와 응답 헤더 모두에 ID 저장
  - Context key: "request_id" (string) / 컨텍스트 키: "request_id" (문자열)
- **Timeout Middleware Architecture / 타임아웃 미들웨어 아키텍처**:
  - Uses context.WithTimeout for timeout enforcement / 타임아웃 적용을 위해 context.WithTimeout 사용
  - Wraps handler with http.TimeoutHandler / http.TimeoutHandler로 핸들러 래핑
  - Default timeout: 30 seconds / 기본 타임아웃: 30초
  - Default message: "Service Unavailable" / 기본 메시지: "Service Unavailable"
  - Timeout is enforced by http.TimeoutHandler / 타임아웃은 http.TimeoutHandler에 의해 적용됨
- **BasicAuth Middleware Architecture / Basic 인증 미들웨어 아키텍처**:
  - Uses r.BasicAuth() to extract credentials / r.BasicAuth()를 사용하여 자격 증명 추출
  - Uses subtle.ConstantTimeCompare for secure password comparison / 안전한 비밀번호 비교를 위해 subtle.ConstantTimeCompare 사용
  - Prevents timing attacks / 타이밍 공격 방지
  - Returns 401 with WWW-Authenticate header on failure / 실패 시 WWW-Authenticate 헤더와 함께 401 반환
  - Stores username in context for later use / 나중에 사용하기 위해 컨텍스트에 사용자 이름 저장
  - Context key: "auth_username" (string) / 컨텍스트 키: "auth_username" (문자열)
- **Configuration Pattern / 설정 패턴**:
  - Default functions: RequestID(), Timeout(), BasicAuth() / 기본 함수
  - Config functions: RequestIDWithConfig(), TimeoutWithConfig(), BasicAuthWithConfig() / 설정 함수
  - Smart defaults for quick start / 빠른 시작을 위한 스마트 기본값
  - Custom validators and generators supported / 커스텀 검증자 및 생성기 지원

### Testing Coverage / 테스트 커버리지
- **9 new middleware test functions** / **9개의 새로운 미들웨어 테스트 함수**
- **3 new benchmark functions** (RequestID, Timeout, BasicAuth) / **3개의 새로운 벤치마크 함수**
- **Total: 114+ test functions** (105 from v1.11.006 + 9 new) / **총 114개 이상의 테스트 함수**
- **Total: 23 benchmark functions** (20 from v1.11.006 + 3 new) / **총 23개의 벤치마크 함수**
- **85.4% test coverage** - All tests passing ✅ / **85.4% 테스트 커버리지** - 모든 테스트 통과 ✅
- Tests cover: request ID generation/preservation, timeout enforcement, basic auth validation, custom configs / 테스트 범위: 요청 ID 생성/보존, 타임아웃 적용, basic 인증 검증, 커스텀 설정

### Performance / 성능
- Middleware benchmarks (sample results) / 미들웨어 벤치마크 (샘플 결과):
  - RequestID: ~300-400 ns/op (includes crypto/rand) / 요청 ID: ~300-400 ns/op (crypto/rand 포함)
  - Timeout: ~400-500 ns/op (includes context creation) / 타임아웃: ~400-500 ns/op (컨텍스트 생성 포함)
  - BasicAuth: ~500-600 ns/op (includes constant-time comparison) / Basic 인증: ~500-600 ns/op (상수 시간 비교 포함)
  - Still minimal overhead for production use / 여전히 프로덕션 사용을 위한 최소 오버헤드

### Notes / 참고사항
- Phase 2 (Middleware System) continued! / Phase 2 (미들웨어 시스템) 계속!
- Total 6 middleware now available (Recovery, Logger, CORS, RequestID, Timeout, BasicAuth) / 총 6개의 미들웨어 사용 가능
- RequestID is essential for request tracing and debugging / 요청 ID는 요청 추적 및 디버깅에 필수적
- Timeout prevents slow clients from blocking resources / 타임아웃은 느린 클라이언트가 리소스를 차단하는 것을 방지
- BasicAuth provides simple authentication for APIs / Basic 인증은 API를 위한 간단한 인증 제공
- All middleware follow consistent naming and config patterns / 모든 미들웨어는 일관된 명명 및 설정 패턴 따름
- Next: v1.11.008 may add more middleware (Rate Limiting, Compression, etc.) / 다음: v1.11.008은 더 많은 미들웨어 추가 예정 (Rate Limiting, Compression 등)

---

## [v1.11.006] - 2025-10-16

### Added / 추가
- Created `middleware.go` with built-in middleware implementations / 내장 미들웨어 구현이 있는 middleware.go 생성
  - **Recovery Middleware / 복구 미들웨어**:
    - `Recovery()` - Default recovery middleware with panic logging / 패닉 로깅이 있는 기본 복구 미들웨어
    - `RecoveryWithConfig(config)` - Custom recovery configuration / 커스텀 복구 설정
    - Captures panics and logs with stack traces / 패닉을 캡처하고 스택 트레이스와 함께 로깅
    - Returns 500 Internal Server Error on panic / 패닉 시 500 Internal Server Error 반환
    - Configurable logging and stack printing / 설정 가능한 로깅 및 스택 출력
  - **Logger Middleware / 로거 미들웨어**:
    - `Logger()` - Default logger middleware / 기본 로거 미들웨어
    - `LoggerWithConfig(config)` - Custom logger configuration / 커스텀 로거 설정
    - Logs method, path, status code, duration / 메서드, 경로, 상태 코드, 소요 시간 로깅
    - Custom log function support / 커스텀 로그 함수 지원
  - **CORS Middleware / CORS 미들웨어**:
    - `CORS()` - Default CORS with wildcard origin / 와일드카드 오리진이 있는 기본 CORS
    - `CORSWithConfig(config)` - Custom CORS configuration / 커스텀 CORS 설정
    - Configurable origins, methods, headers / 설정 가능한 오리진, 메서드, 헤더
    - Automatic preflight (OPTIONS) request handling / 자동 프리플라이트 요청 처리
    - Credentials and max-age support / 자격 증명 및 max-age 지원
- Added configuration structs / 설정 구조체 추가
  - `RecoveryConfig` - Recovery middleware configuration / 복구 미들웨어 설정
  - `LoggerConfig` - Logger middleware configuration / 로거 미들웨어 설정
  - `CORSConfig` - CORS middleware configuration / CORS 미들웨어 설정
- Added helper types / 헬퍼 타입 추가
  - `responseWriter` - Status code tracking wrapper / 상태 코드 추적 래퍼
  - Helper functions: `isOriginAllowed`, `joinStrings` / 헬퍼 함수
- Created comprehensive `middleware_test.go` with 15 test functions / 15개의 테스트 함수가 있는 포괄적인 middleware_test.go 생성
  - Recovery tests: TestRecovery, TestRecoveryNoPanic, TestRecoveryWithConfig / 복구 테스트
  - Logger tests: TestLogger, TestLoggerWithConfig / 로거 테스트
  - CORS tests: TestCORS, TestCORSPreflight, TestCORSWithConfig, TestCORSNotAllowedOrigin / CORS 테스트
  - Helper tests: TestResponseWriter, TestIsOriginAllowed, TestJoinStrings / 헬퍼 테스트
  - 3 benchmark functions / 3개의 벤치마크 함수
- Updated `README.md` with Middleware documentation / 미들웨어 문서로 README.md 업데이트
  - Added comprehensive Middleware section / 포괄적인 미들웨어 섹션 추가
  - Recovery, Logger, CORS features documentation / 복구, 로거, CORS 기능 문서
  - Updated version to v1.11.006 / 버전을 v1.11.006로 업데이트
  - Updated development status progress / 개발 상태 진행 상황 업데이트
  - Updated current phase to Phase 2 / 현재 단계를 Phase 2로 업데이트

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.006 / websvrutil.go 버전 상수를 v1.11.006로 업데이트
- Bumped version to v1.11.006 in `cfg/app.yaml` / cfg/app.yaml의 버전을 v1.11.006로 상향

### Technical Details / 기술 세부사항
- **Recovery Middleware Architecture / 복구 미들웨어 아키텍처**:
  - Uses defer/recover pattern to catch panics / defer/recover 패턴을 사용하여 패닉 캡처
  - Logs panic value and stack trace / 패닉 값 및 스택 트레이스 로깅
  - Returns 500 status code / 500 상태 코드 반환
  - Configurable: PrintStack, LogFunc / 설정 가능
- **Logger Middleware Architecture / 로거 미들웨어 아키텍처**:
  - Uses responseWriter wrapper to capture status code / responseWriter 래퍼를 사용하여 상태 코드 캡처
  - Measures request duration with time.Now() / time.Now()로 요청 소요 시간 측정
  - Logs after handler completes / 핸들러 완료 후 로깅
  - Custom log function support / 커스텀 로그 함수 지원
- **CORS Middleware Architecture / CORS 미들웨어 아키텍처**:
  - Sets Access-Control-* headers based on configuration / 설정에 따라 Access-Control-* 헤더 설정
  - Handles preflight OPTIONS requests / 프리플라이트 OPTIONS 요청 처리
  - Origin validation with wildcard support / 와일드카드 지원이 있는 오리진 검증
  - Supports credentials, max-age, exposed headers / 자격 증명, max-age, 노출 헤더 지원
- **Configuration Pattern / 설정 패턴**:
  - Default functions: Recovery(), Logger(), CORS() / 기본 함수
  - Config functions: RecoveryWithConfig(), LoggerWithConfig(), CORSWithConfig() / 설정 함수
  - Smart defaults for 99% use cases / 99% 사용 사례를 위한 스마트 기본값
- **responseWriter Helper / responseWriter 헬퍼**:
  - Wraps http.ResponseWriter / http.ResponseWriter 래핑
  - Tracks status code for logging / 로깅을 위한 상태 코드 추적
  - Defaults to 200 OK if not explicitly set / 명시적으로 설정하지 않으면 200 OK로 기본 설정

### Testing Coverage / 테스트 커버리지
- **15 new middleware test functions** / **15개의 새로운 미들웨어 테스트 함수**
- **3 new benchmark functions** (Recovery, Logger, CORS) / **3개의 새로운 벤치마크 함수**
- **Total: 105+ test functions** (90 from v1.11.005 + 15 new) / **총 105개 이상의 테스트 함수**
- **Total: 20 benchmark functions** (17 from v1.11.005 + 3 new) / **총 20개의 벤치마크 함수**
- **84.0% test coverage** - All tests passing ✅ / **84.0% 테스트 커버리지** - 모든 테스트 통과 ✅
- Tests cover: panic recovery, normal flow, custom configs, logging, CORS headers, preflight, origin validation / 테스트 범위: 패닉 복구, 정상 흐름, 커스텀 설정, 로깅, CORS 헤더, 프리플라이트, 오리진 검증

### Performance / 성능
- Middleware benchmarks (sample results) / 미들웨어 벤치마크 (샘플 결과):
  - Recovery: ~200-300 ns/op (no panic) / 복구: ~200-300 ns/op (패닉 없음)
  - Logger: ~300-400 ns/op / 로거: ~300-400 ns/op
  - CORS: ~200-300 ns/op / CORS: ~200-300 ns/op
  - Minimal overhead for production use / 프로덕션 사용을 위한 최소 오버헤드

### Notes / 참고사항
- Phase 2 (Middleware System) started! / Phase 2 (미들웨어 시스템) 시작!
- All three core middleware (Recovery, Logger, CORS) implemented in single version / 세 가지 핵심 미들웨어를 단일 버전에서 구현
- Smart defaults make middleware usage extremely simple / 스마트 기본값으로 미들웨어 사용이 극도로 간단함
- Custom configuration available for advanced use cases / 고급 사용 사례를 위한 커스텀 설정 제공
- responseWriter wrapper enables status code tracking / responseWriter 래퍼가 상태 코드 추적 가능
- Next: v1.11.007 will add more middleware features (Rate Limiting, Request ID, etc.) / 다음: v1.11.007은 더 많은 미들웨어 기능 추가 예정

---

## [v1.11.005] - 2025-10-16

### Added / 추가
- Added 11 response helper methods to Context / Context에 11개의 응답 헬퍼 메서드 추가
  - **JSON Response / JSON 응답**:
    - `JSON(code, data)` - Send JSON response / JSON 응답 전송
    - `JSONPretty(code, data)` - Send pretty JSON (2-space indent) / 보기 좋은 JSON 전송
    - `JSONIndent(code, data, prefix, indent)` - Custom indentation / 커스텀 들여쓰기
    - `Error(code, message)` - Send JSON error response / JSON 에러 응답 전송
  - **HTML Response / HTML 응답**:
    - `HTML(code, html)` - Send HTML response / HTML 응답 전송
    - `HTMLTemplate(code, tmpl, data)` - Render HTML template / HTML 템플릿 렌더링
  - **Text Response / 텍스트 응답**:
    - `Text(code, text)` - Send plain text / 일반 텍스트 전송
    - `Textf(code, format, args...)` - Send formatted text / 형식화된 텍스트 전송
  - **Other Responses / 기타 응답**:
    - `XML(code, xml)` - Send XML response / XML 응답 전송
    - `Redirect(code, url)` - HTTP redirect / HTTP 리다이렉트
    - `NoContent()` - Send 204 No Content / 204 No Content 전송
- Updated `context.go` imports / context.go imports 업데이트
  - Added `encoding/json` for JSON marshaling / JSON 마샬링을 위한 encoding/json 추가
  - Added `fmt` for string formatting / 문자열 형식화를 위한 fmt 추가
  - Added `html/template` for template rendering / 템플릿 렌더링을 위한 html/template 추가
- Created comprehensive tests for response helpers / 응답 헬퍼를 위한 포괄적인 테스트 생성
  - 14 new test functions covering all response methods / 모든 응답 메서드를 다루는 14개의 새로운 테스트 함수
  - Tests for JSON, JSONPretty, JSONIndent, HTML, HTMLTemplate, Text, Textf, XML, Redirect, NoContent, Error / JSON, JSONPretty, JSONIndent, HTML, HTMLTemplate, Text, Textf, XML, Redirect, NoContent, Error 테스트
  - Error handling tests (HTMLTemplate parsing error) / 에러 처리 테스트
  - 3 new benchmark functions / 3개의 새로운 벤치마크 함수
- Updated `README.md` with Response Helpers documentation / Response Helpers 문서로 README.md 업데이트
  - Added comprehensive response methods documentation / 포괄적인 응답 메서드 문서 추가
  - Organized by category (JSON, HTML, Text, Other) / 카테고리별 구성
  - Updated version to v1.11.005 / 버전을 v1.11.005로 업데이트
  - Updated progress status / 진행 상태 업데이트

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.005 / websvrutil.go 버전 상수를 v1.11.005로 업데이트
- Bumped version to v1.11.005 in `cfg/app.yaml` / cfg/app.yaml의 버전을 v1.11.005로 상향

### Technical Details / 기술 세부사항
- **Response Helper Categories / 응답 헬퍼 카테고리**:
  - JSON: Full JSON support with pretty-printing and custom indentation / 보기 좋은 출력 및 커스텀 들여쓰기를 포함한 완전한 JSON 지원
  - HTML: Direct HTML and template rendering / 직접 HTML 및 템플릿 렌더링
  - Text: Plain text and formatted text (Printf-style) / 일반 텍스트 및 형식화된 텍스트
  - Other: XML, redirects, and no-content responses / XML, 리다이렉트 및 콘텐츠 없음 응답
- **Content-Type Headers / Content-Type 헤더**:
  - JSON: `application/json; charset=utf-8`
  - HTML: `text/html; charset=utf-8`
  - Text: `text/plain; charset=utf-8`
  - XML: `application/xml; charset=utf-8`
- **Error Response Format / 에러 응답 형식**:
  - JSON object with `error`, `message`, and `status` fields / error, message, status 필드가 있는 JSON 객체
  - Automatically includes HTTP status text / HTTP 상태 텍스트 자동 포함
- **Template Rendering / 템플릿 렌더링**:
  - Uses Go's `html/template` package / Go의 html/template 패키지 사용
  - Inline template parsing / 인라인 템플릿 파싱
  - Error handling for invalid templates / 잘못된 템플릿에 대한 에러 처리

### Testing Coverage / 테스트 커버리지
- **14 new response helper test functions** / **14개의 새로운 응답 헬퍼 테스트 함수**
- **3 new benchmark functions** (JSON, HTML, Text) / **3개의 새로운 벤치마크 함수**
- **Total: 90+ test functions** (76 from v1.11.004 + 14 new) / **총 90개 이상의 테스트 함수**
- **Total: 17 benchmark functions** (14 from v1.11.004 + 3 new) / **총 17개의 벤치마크 함수**
- **82.7% test coverage** - All tests passing ✅ / **82.7% 테스트 커버리지** - 모든 테스트 통과 ✅
- Tests cover: JSON (standard/pretty/indent), HTML (direct/template), Text (plain/formatted), XML, Redirect, NoContent, Error / 테스트 범위: JSON (표준/보기좋은/들여쓰기), HTML (직접/템플릿), Text (일반/형식화), XML, Redirect, NoContent, Error

### Performance / 성능
- Response helper benchmarks (sample results) / 응답 헬퍼 벤치마크 (샘플 결과):
  - JSON encoding: ~1-2 μs/op (depends on data size) / JSON 인코딩: 데이터 크기에 따라 다름
  - HTML response: ~100-200 ns/op
  - Text response: ~100-200 ns/op
  - Template rendering: ~5-10 μs/op (simple templates) / 템플릿 렌더링: 간단한 템플릿 기준

### Notes / 참고사항
- Response helpers provide convenient methods for common response types / 응답 헬퍼는 일반적인 응답 타입을 위한 편리한 메서드 제공
- All methods automatically set appropriate Content-Type headers / 모든 메서드가 자동으로 적절한 Content-Type 헤더 설정
- JSON encoding uses streaming encoder for efficiency / JSON 인코딩은 효율성을 위해 스트리밍 인코더 사용
- Template rendering supports inline templates (file templates in future versions) / 템플릿 렌더링은 인라인 템플릿 지원 (파일 템플릿은 향후 버전에서)
- Phase 1 (Core Foundation) complete! / Phase 1 (핵심 기반) 완료!
- Next: Phase 2 - Middleware System (v1.11.006-010) / 다음: Phase 2 - 미들웨어 시스템

---

## [v1.11.004] - 2025-10-16

### Added / 추가
- Created `context.go` with Context implementation / Context 구현이 있는 context.go 생성
  - `Context` struct for request context management / 요청 컨텍스트 관리를 위한 Context 구조체
  - Thread-safe with sync.RWMutex / sync.RWMutex로 스레드 안전
  - Parameter access: `Param(name)`, `Params()` / 매개변수 액세스
  - Custom value storage: `Set(key, value)`, `Get(key)`, `MustGet(key)` / 커스텀 값 저장
  - Typed getters: `GetString()`, `GetInt()`, `GetBool()` / 타입 지정 getter
  - Query parameters: `Query(key)`, `QueryDefault(key, default)` / 쿼리 매개변수
  - Header access: `Header(key)`, `SetHeader(key, value)` / 헤더 액세스
  - Request helpers: `Method()`, `Path()`, `Context()`, `WithContext()` / 요청 헬퍼
  - Response helpers: `Status(code)`, `Write(data)`, `WriteString(s)` / 응답 헬퍼
  - Helper function: `GetContext(r *http.Request)` / 헬퍼 함수
- Modified `router.go` to create Context and store parameters / Context를 생성하고 매개변수를 저장하도록 router.go 수정
  - Router now creates Context for each request / Router가 각 요청에 대해 Context 생성
  - Parameters extracted from path are stored in Context / 경로에서 추출된 매개변수가 Context에 저장
  - Context stored in request's context.Context / Context가 요청의 context.Context에 저장
  - Added `contextWithValue()` helper function / contextWithValue() 헬퍼 함수 추가
- Created comprehensive `context_test.go` with 24+ test functions / 24개 이상의 테스트 함수가 있는 포괄적인 context_test.go 생성
  - Context creation tests / Context 생성 테스트
  - Parameter access tests (Param, Params) / 매개변수 액세스 테스트
  - Custom value storage tests (Set, Get, MustGet) / 커스텀 값 저장 테스트
  - Typed getter tests (GetString, GetInt, GetBool) / 타입 지정 getter 테스트
  - Query parameter tests (Query, QueryDefault) / 쿼리 매개변수 테스트
  - Header tests (Header, SetHeader) / 헤더 테스트
  - Response tests (Status, Write, WriteString) / 응답 테스트
  - Request helper tests (Method, Path, Context, WithContext) / 요청 헬퍼 테스트
  - GetContext helper tests / GetContext 헬퍼 테스트
  - 3 benchmark functions (NewContext, SetGet, Param) / 3개의 벤치마크 함수
- Updated `README.md` with Context documentation / Context 문서로 README.md 업데이트
  - Added comprehensive Context features section / 포괄적인 Context 기능 섹션 추가
  - Updated quick start with Context examples / Context 예제로 빠른 시작 업데이트
  - Added Context usage example with 4 scenarios / 4가지 시나리오가 있는 Context 사용 예제 추가
  - Updated version to v1.11.004 / 버전을 v1.11.004로 업데이트
  - Updated development status progress / 개발 상태 진행 상황 업데이트
- Updated `examples/websvrutil/main.go` with Context examples / Context 예제로 examples/websvrutil/main.go 업데이트
  - 14 total examples (added 4 new Context examples) / 총 14개 예제 (4개의 새로운 Context 예제 추가)
  - Example 7: Context - Path parameters / Context - 경로 매개변수
  - Example 8: Context - Query parameters / Context - 쿼리 매개변수
  - Example 9: Context - Custom values / Context - 커스텀 값
  - Example 10: Context - Request headers / Context - 요청 헤더
  - Renamed examples 7-10 to 11-14 / 예제 7-10을 11-14로 이름 변경

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.004 / websvrutil.go 버전 상수를 v1.11.004로 업데이트
- Bumped version to v1.11.004 in `cfg/app.yaml` / cfg/app.yaml의 버전을 v1.11.004로 상향
- Modified Router ServeHTTP to create and inject Context / Router ServeHTTP를 Context를 생성하고 주입하도록 수정

### Technical Details / 기술 세부사항
- **Context Architecture** / **Context 아키텍처**:
  - Request-scoped context for parameter and value storage / 매개변수 및 값 저장을 위한 요청 범위 컨텍스트
  - Thread-safe with sync.RWMutex (concurrent read, exclusive write) / sync.RWMutex로 스레드 안전 (동시 읽기, 배타적 쓰기)
  - Stored in request's context.Context for retrieval / 검색을 위해 요청의 context.Context에 저장
  - Provides convenient access to common request data / 일반적인 요청 데이터에 대한 편리한 액세스 제공
- **Context Features** / **Context 기능**:
  - Parameter access: Path parameters from route patterns / 매개변수 액세스: 라우트 패턴의 경로 매개변수
  - Custom values: Store/retrieve arbitrary values / 커스텀 값: 임의의 값 저장/검색
  - Query helpers: Easy query parameter access / 쿼리 헬퍼: 쉬운 쿼리 매개변수 액세스
  - Header helpers: Read request/write response headers / 헤더 헬퍼: 요청 헤더 읽기/응답 헤더 쓰기
  - Response helpers: Write status and body / 응답 헬퍼: 상태 및 본문 작성
- **Integration with Router** / **Router와의 통합**:
  - Router creates Context for each request / Router가 각 요청에 대해 Context 생성
  - Parameters from route matching stored in Context / 라우트 일치에서 나온 매개변수가 Context에 저장
  - Context accessible via `GetContext(r *http.Request)` / GetContext(r *http.Request)를 통해 Context 액세스 가능
  - Context stored using internal context key / 내부 컨텍스트 키를 사용하여 Context 저장

### Testing Coverage / 테스트 커버리지
- **24+ new context test functions** / **24개 이상의 새로운 context 테스트 함수**
- **3 context benchmark functions** / **3개의 context 벤치마크 함수**
- **Total: 76+ test functions** (52 from v1.11.003 + 24 new) / **총 76개 이상의 테스트 함수**
- **Total: 14 benchmark functions** (11 from v1.11.003 + 3 new) / **총 14개의 벤치마크 함수**
- Tests cover: Context creation, parameter access, custom values, query/headers, response helpers / 테스트 범위: Context 생성, 매개변수 액세스, 커스텀 값, 쿼리/헤더, 응답 헬퍼

### Performance / 성능
- Context benchmarks (sample results) / Context 벤치마크 (샘플 결과):
  - NewContext: ~100-150 ns/op
  - Set/Get operations: ~50-100 ns/op
  - Param access: ~10-20 ns/op
  - Thread-safe operations with minimal overhead / 최소 오버헤드로 스레드 안전 작업

### Notes / 참고사항
- Path parameters are now fully accessible via Context / 경로 매개변수는 이제 Context를 통해 완전히 액세스 가능
- Context provides convenient helpers for common request/response operations / Context는 일반적인 요청/응답 작업을 위한 편리한 헬퍼 제공
- Thread-safe for concurrent access (multiple goroutines can read simultaneously) / 동시 액세스에 안전 (여러 고루틴이 동시에 읽을 수 있음)
- Next: v1.11.005 will add JSON/HTML/Text response helpers / 다음: v1.11.005는 JSON/HTML/Text 응답 헬퍼 추가 예정

---

## [v1.11.003] - 2025-10-16

### Added / 추가
- Created `router.go` with Router implementation / Router 구현이 있는 router.go 생성
  - HTTP method routing (GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD) / HTTP 메서드 라우팅
  - Path parameter extraction (`:id`, `:name`) / 경로 매개변수 추출
  - Wildcard route matching (`*`) / 와일드카드 라우트 일치
  - Custom 404 handler support / 커스텀 404 핸들러 지원
  - Thread-safe route registration / 스레드 안전 라우트 등록
- Added routing methods to App struct / App 구조체에 라우팅 메서드 추가
  - `GET(pattern, handler)` - Register GET route / GET 라우트 등록
  - `POST(pattern, handler)` - Register POST route / POST 라우트 등록
  - `PUT(pattern, handler)` - Register PUT route / PUT 라우트 등록
  - `PATCH(pattern, handler)` - Register PATCH route / PATCH 라우트 등록
  - `DELETE(pattern, handler)` - Register DELETE route / DELETE 라우트 등록
  - `OPTIONS(pattern, handler)` - Register OPTIONS route / OPTIONS 라우트 등록
  - `HEAD(pattern, handler)` - Register HEAD route / HEAD 라우트 등록
  - `NotFound(handler)` - Set custom 404 handler / 커스텀 404 핸들러 설정
- Created comprehensive `router_test.go` with 24 test functions / 24개의 테스트 함수가 있는 포괄적인 router_test.go 생성
  - Route registration tests (GET, POST, all methods) / 라우트 등록 테스트
  - Parameter extraction tests / 매개변수 추출 테스트
  - Wildcard route tests / 와일드카드 라우트 테스트
  - Custom 404 handler tests / 커스텀 404 핸들러 테스트
  - Pattern parsing tests / 패턴 파싱 테스트
  - Path parsing tests / 경로 파싱 테스트
  - App router integration tests / 앱 라우터 통합 테스트
  - 5 benchmark functions (router performance) / 5개의 벤치마크 함수
- Updated `README.md` with Router documentation / Router 문서로 README.md 업데이트
  - Router features and pattern syntax / Router 기능 및 패턴 구문
  - Updated quick start with routing examples / 라우팅 예제로 빠른 시작 업데이트
  - Added wildcard and custom 404 example / 와일드카드 및 커스텀 404 예제 추가
- Updated `examples/websvrutil/main.go` with Router examples / Router 예제로 examples/websvrutil/main.go 업데이트
  - 10 total examples (added 4 new routing examples) / 총 10개 예제 (4개의 새로운 라우팅 예제 추가)
  - Example 3: Routing with GET/POST / GET/POST 라우팅
  - Example 4: Path parameters / 경로 매개변수
  - Example 5: Wildcard routes / 와일드카드 라우트
  - Example 6: Custom 404 handler / 커스텀 404 핸들러

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.003 / websvrutil.go 버전 상수를 v1.11.003으로 업데이트
- Bumped version to v1.11.003 in `cfg/app.yaml` / cfg/app.yaml의 버전을 v1.11.003으로 상향
- Modified App struct to use Router instead of placeholder / App 구조체를 임시 대신 Router를 사용하도록 수정
- Updated `New()` to automatically create router instance / `New()`가 라우터 인스턴스를 자동으로 생성하도록 업데이트

### Technical Details / 기술 세부사항
- **Router Architecture** / **라우터 아키텍처**:
  - Segment-based pattern matching for performance / 성능을 위한 세그먼트 기반 패턴 일치
  - Pattern parsing on registration (once) / 등록 시 패턴 파싱 (1회)
  - Path parsing on each request (fast) / 각 요청 시 경로 파싱 (빠름)
  - Thread-safe with sync.RWMutex / sync.RWMutex로 스레드 안전
- **Pattern Types** / **패턴 타입**:
  - Literal segments: `/users`, `/posts` / 리터럴 세그먼트
  - Parameter segments: `:id`, `:userId` / 매개변수 세그먼트
  - Wildcard segment: `*` (matches all remaining) / 와일드카드 세그먼트 (나머지 모두 일치)
- **Route Matching** / **라우트 일치**:
  - Exact match for literals / 리터럴 정확한 일치
  - Parameter extraction for `:name` segments / `:name` 세그먼트 매개변수 추출
  - Greedy match for wildcard `*` / 와일드카드 `*` 욕심 일치
  - Method-specific routing (GET /users != POST /users) / 메서드별 라우팅

### Testing Coverage / 테스트 커버리지
- **24 new router test functions** / **24개의 새로운 라우터 테스트 함수**
- **5 router benchmark functions** / **5개의 라우터 벤치마크 함수**
- **Total: 52 test functions** (28 from v1.11.002 + 24 new) / **총 52개의 테스트 함수**
- **Total: 11 benchmark functions** (6 from v1.11.002 + 5 new) / **총 11개의 벤치마크 함수**
- Tests cover: route registration, matching, parameters, wildcards, 404, integration / 테스트 범위: 라우트 등록, 일치, 매개변수, 와일드카드, 404, 통합

### Performance / 성능
- Router benchmarks (sample results) / 라우터 벤치마크 (샘플 결과):
  - Simple route: ~700 ns/op
  - Parameter route: ~700 ns/op
  - Wildcard route: ~700 ns/op
  - Pattern parsing: ~80 ns/op
  - Path parsing: ~50 ns/op

### Notes / 참고사항
- Path parameters are extracted but not yet accessible (coming in v1.11.004) / 경로 매개변수는 추출되지만 아직 액세스 불가 (v1.11.004에서 예정)
- Context API will provide parameter access in v1.11.004 / Context API는 v1.11.004에서 매개변수 액세스 제공
- Router is fully functional for route matching and method routing / Router는 라우트 일치 및 메서드 라우팅에 완전히 작동

---

## [v1.11.002] - 2025-10-16

### Added / 추가
- Created `app.go` with App struct and core methods / App 구조체 및 핵심 메서드가 있는 app.go 생성
  - `New(opts ...Option) *App` - Create new app instance / 새 앱 인스턴스 생성
  - `Use(middleware ...MiddlewareFunc) *App` - Add middleware / 미들웨어 추가
  - `Run(addr string) error` - Start HTTP server / HTTP 서버 시작
  - `Shutdown(ctx context.Context) error` - Graceful shutdown / 정상 종료
  - `ServeHTTP(w http.ResponseWriter, r *http.Request)` - Implement http.Handler / http.Handler 구현
  - `buildHandler()` - Build middleware chain / 미들웨어 체인 구축
- Created `options.go` with Options pattern / Options 패턴이 있는 options.go 생성
  - `Options` struct with 10 configuration fields / 10개의 설정 필드가 있는 Options 구조체
  - `defaultOptions()` - Smart default values / 스마트 기본값
  - 10 functional option functions (WithReadTimeout, WithWriteTimeout, etc.) / 10개의 함수형 옵션 함수
- Created comprehensive `app_test.go` with 15 test functions / 15개의 테스트 함수가 있는 포괄적인 app_test.go 생성
  - `TestNew`, `TestNewWithOptions` - App creation tests / 앱 생성 테스트
  - `TestUse`, `TestUseMultiple` - Middleware tests / 미들웨어 테스트
  - `TestServeHTTP`, `TestMiddlewareOrder` - HTTP handler tests / HTTP 핸들러 테스트
  - `TestShutdownWithoutRun`, `TestRunInvalidAddress` - Error handling tests / 에러 처리 테스트
  - `TestConcurrentUse` - Concurrency safety test / 동시성 안전성 테스트
  - 4 benchmark functions (New, NewWithOptions, Use, ServeHTTP) / 4개의 벤치마크 함수
- Created comprehensive `options_test.go` with 13 test functions / 13개의 테스트 함수가 있는 포괄적인 options_test.go 생성
  - Individual option tests for all 10 options / 10개 옵션 모두에 대한 개별 옵션 테스트
  - `TestMultipleOptions` - Combined options test / 결합된 옵션 테스트
  - `TestOptionsOverride` - Options priority test / 옵션 우선순위 테스트
  - `TestOptionsImmutability` - Instance isolation test / 인스턴스 격리 테스트
  - 2 benchmark functions (WithReadTimeout, MultipleOptions) / 2개의 벤치마크 함수
- Created comprehensive `README.md` for the package / 패키지를 위한 포괄적인 README.md 생성
  - Package overview and design philosophy / 패키지 개요 및 설계 철학
  - Installation instructions / 설치 지침
  - Current features documentation (App, Options) / 현재 기능 문서 (App, Options)
  - Configuration reference table / 설정 참조 테이블
  - 4 quick start examples with bilingual code / 4개의 빠른 시작 예제 (이중 언어 코드)
  - Upcoming features roadmap / 예정된 기능 로드맵
- Created comprehensive `examples/websvrutil/main.go` / 포괄적인 examples/websvrutil/main.go 생성
  - 6 complete examples demonstrating all features / 모든 기능을 시연하는 6개의 완전한 예제
  - Example 1: Basic server / 기본 서버
  - Example 2: Custom options / 커스텀 옵션
  - Example 3: Graceful shutdown / 정상 종료
  - Example 4: Custom middleware / 커스텀 미들웨어
  - Example 5: Multiple middleware / 다중 미들웨어
  - Example 6: Production configuration / 프로덕션 설정

### Changed / 변경
- Updated `websvrutil.go` version constant to v1.11.002 / websvrutil.go 버전 상수를 v1.11.002로 업데이트
- Bumped version to v1.11.002 in `cfg/app.yaml` / cfg/app.yaml의 버전을 v1.11.002로 상향

### Technical Details / 기술 세부사항
- **App struct**: Main application instance managing web server / 웹 서버를 관리하는 주요 애플리케이션 인스턴스
  - Manages middleware chain, router, templates, HTTP server / 미들웨어 체인, 라우터, 템플릿, HTTP 서버 관리
  - Thread-safe with sync.RWMutex / sync.RWMutex로 스레드 안전
  - Supports middleware chaining / 미들웨어 체이닝 지원
  - Implements http.Handler interface / http.Handler 인터페이스 구현
- **Options Pattern**: Functional options for flexible configuration / 유연한 설정을 위한 함수형 옵션
  - 10 configuration options with smart defaults / 스마트 기본값이 있는 10개의 설정 옵션
  - ReadTimeout, WriteTimeout, IdleTimeout (server timeouts) / 서버 타임아웃
  - MaxHeaderBytes (security limit) / 보안 제한
  - TemplateDir, StaticDir, StaticPrefix (directories) / 디렉토리
  - EnableAutoReload, EnableLogger, EnableRecovery (features) / 기능
- **Middleware System**: Standard http.Handler wrapping pattern / 표준 http.Handler 래핑 패턴
  - MiddlewareFunc type: `func(http.Handler) http.Handler`
  - Executed in order of addition (first added = outermost) / 추가 순서대로 실행
  - Cannot add middleware while server is running (panic) / 서버 실행 중 미들웨어 추가 불가

### Testing Coverage / 테스트 커버리지
- **28 test functions** total (15 app + 13 options) / 총 28개의 테스트 함수
- **6 benchmark functions** (4 app + 2 options) / 6개의 벤치마크 함수
- Tests cover: creation, configuration, middleware, HTTP handling, concurrency, error cases / 테스트 범위: 생성, 설정, 미들웨어, HTTP 처리, 동시성, 에러 케이스

### Notes / 참고사항
- Router, Context, and Template features are placeholders (coming in v1.11.003+) / Router, Context, Template 기능은 임시 (v1.11.003+ 예정)
- Default handler returns 404 for all requests until router is implemented / 라우터 구현 전까지 기본 핸들러는 모든 요청에 404 반환
- Graceful shutdown with context timeout support / 컨텍스트 타임아웃을 지원하는 정상 종료

---

## [v1.11.001] - 2025-10-16

### Added / 추가
- Created `websvrutil` package structure / websvrutil 패키지 구조 생성
- Created `websvrutil.go` with package information and version / 패키지 정보 및 버전이 포함된 websvrutil.go 생성
- Created comprehensive `DESIGN_PLAN.md` (60+ pages) / 포괄적인 DESIGN_PLAN.md 생성 (60페이지 이상)
  - Package overview and design philosophy / 패키지 개요 및 설계 철학
  - Architecture and core components / 아키텍처 및 핵심 컴포넌트
  - API design examples / API 설계 예제
  - Template system design / 템플릿 시스템 설계
  - Error handling and performance considerations / 에러 처리 및 성능 고려사항
- Created detailed `WORK_PLAN.md` with 6 phases / 6개 단계로 구성된 상세 WORK_PLAN.md 생성
  - Phase 1: Core Foundation (v1.11.001-005) / 핵심 기반
  - Phase 2: Middleware System (v1.11.006-010) / 미들웨어 시스템
  - Phase 3: Template System (v1.11.011-015) / 템플릿 시스템
  - Phase 4: Advanced Features (v1.11.016-020) / 고급 기능
  - Phase 5: Server Management (v1.11.021-025) / 서버 관리
  - Phase 6: Documentation & Polish (v1.11.026-030) / 문서화 및 마무리
- Created feature branch `feature/v1.11.x-websvrutil` / 기능 브랜치 생성
- Updated root `README.md` with websvrutil "In Development" status / 루트 README.md에 websvrutil "개발 중" 상태 업데이트
- Updated `CHANGELOG.md` with v1.11.x section / CHANGELOG.md에 v1.11.x 섹션 업데이트
- Created directory structure: `websvrutil/`, `docs/websvrutil/`, `examples/websvrutil/` / 디렉토리 구조 생성

### Changed / 변경
- Bumped version to v1.11.001 in `cfg/app.yaml` / cfg/app.yaml의 버전을 v1.11.001로 상향

### Notes / 참고사항
- **Design Philosophy / 설계 철학**: Developer convenience first (50+ lines → 5 lines) / 개발자 편의성 우선 (50줄 이상 → 5줄)
- **Key Principles / 주요 원칙**:
  - Extreme simplicity over performance / 성능보다 극도의 간결함
  - Smart defaults for 99% use cases / 99% 사용 사례를 위한 스마트 기본값
  - Auto template discovery and hot reload / 자동 템플릿 발견 및 핫 리로드
  - Easy middleware chaining / 쉬운 미들웨어 체이닝
  - Standard library compatible / 표준 라이브러리 호환

---

## Roadmap / 로드맵

### Phase 1: Core Foundation (v1.11.001-005)
- ✅ v1.11.001: Project setup and planning
- ✅ v1.11.002: App & Options
- ✅ v1.11.003: Router
- 📝 v1.11.004: Context (Part 1)
- 📝 v1.11.005: Response Helpers

### Phase 2: Middleware System (v1.11.006-010)
- 📝 v1.11.006: Middleware Chain
- 📝 v1.11.007: Recovery Middleware
- 📝 v1.11.008: Logger Middleware
- 📝 v1.11.009: CORS Middleware
- 📝 v1.11.010: Auth Middleware

### Phase 3: Template System (v1.11.011-015)
- 📝 v1.11.011: Template Engine Core
- 📝 v1.11.012: Auto Template Discovery
- 📝 v1.11.013: Layout System
- 📝 v1.11.014: Custom Template Functions
- 📝 v1.11.015: Hot Reload

### Phase 4: Advanced Features (v1.11.016-020)
- 📝 v1.11.016: Request Binding
- 📝 v1.11.017: Cookie & Header Helpers
- 📝 v1.11.018: File Upload
- 📝 v1.11.019: Static File Serving
- 📝 v1.11.020: Context Storage

### Phase 5: Server Management (v1.11.021-025)
- 📝 v1.11.021: Graceful Shutdown
- 📝 v1.11.022: Health Check
- 📝 v1.11.023: Route Groups
- 📝 v1.11.024: Error Handling
- 📝 v1.11.025: Server Utilities

### Phase 6: Documentation & Polish (v1.11.026-030)
- 📝 v1.11.026: USER_MANUAL.md
- 📝 v1.11.027: DEVELOPER_GUIDE.md
- 📝 v1.11.028: Comprehensive Examples
- 📝 v1.11.029: Testing & Benchmarks
- 📝 v1.11.030: Final Polish

---

**Legend / 범례**:
- ✅ Completed / 완료
- 🔄 In Progress / 진행 중
- 📝 Planned / 계획됨
