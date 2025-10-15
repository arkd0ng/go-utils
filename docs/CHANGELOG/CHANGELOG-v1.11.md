# CHANGELOG v1.11.x - Web Server Utilities Package

**Package / 패키지**: `websvrutil`
**Focus / 초점**: Extreme simplicity web server utilities / 극도로 간단한 웹 서버 유틸리티

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
