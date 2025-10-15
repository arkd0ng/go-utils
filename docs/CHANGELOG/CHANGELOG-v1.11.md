# CHANGELOG v1.11.x - Web Server Utilities Package

**Package / 패키지**: `websvrutil`
**Focus / 초점**: Extreme simplicity web server utilities / 극도로 간단한 웹 서버 유틸리티

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
