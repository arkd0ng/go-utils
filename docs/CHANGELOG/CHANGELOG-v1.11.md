# CHANGELOG v1.11.x - Web Server Utilities Package

**Package / 패키지**: `websvrutil`
**Focus / 초점**: Extreme simplicity web server utilities / 극도로 간단한 웹 서버 유틸리티

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
- 📝 v1.11.003: Router
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
