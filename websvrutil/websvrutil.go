// Package websvrutil provides extreme simplicity web server utilities for building production-ready HTTP servers.
// It dramatically reduces the complexity of web server setup by providing a high-level, easy-to-use API
// that handles common server patterns including routing, middleware, context management, sessions,
// CSRF protection, template rendering, and graceful shutdown.
//
// The package is designed with the following principles:
//   - Simplicity: Reduce 50+ lines of boilerplate server setup code to just 5 lines
//   - Safety: Built-in CSRF protection, secure session management, and panic recovery
//   - Performance: Optimized context pooling, efficient middleware chain, and zero-allocation paths
//   - Flexibility: Extensible middleware system and customizable options
//   - Production-ready: Graceful shutdown, health checks, and comprehensive error handling
//
// Main Features:
//   - App Management: Create and manage HTTP server lifecycle with App struct
//   - Routing: Flexible route registration with support for path parameters and HTTP methods
//   - Context: Enhanced request/response handling through Context abstraction
//   - Middleware: Composable middleware chain for cross-cutting concerns
//   - Session: Secure session management with multiple storage backends
//   - CSRF: Built-in Cross-Site Request Forgery protection
//   - Template: Template engine integration with layout support
//   - Validation: Request data validation with detailed error messages
//   - Graceful Shutdown: Safe server shutdown with connection draining
//
// Basic Usage:
//   app := websvrutil.New()
//   app.GET("/", func(c *websvrutil.Context) error {
//       return c.String(200, "Hello, World!")
//   })
//   app.Start(":8080")
//
// Advanced Usage with Middleware:
//   app := websvrutil.New(
//       websvrutil.WithLogger(logger),
//       websvrutil.WithCSRF(),
//       websvrutil.WithSession(sessionStore),
//   )
//   
//   // Group routes with common middleware
//   api := app.Group("/api", authMiddleware)
//   api.GET("/users", listUsers)
//   api.POST("/users", createUser)
//   
//   // Start with graceful shutdown
//   app.StartWithGracefulShutdown(":8080", 30*time.Second)
//
// Performance Characteristics:
//   - Context pooling reduces GC pressure by reusing Context objects
//   - Middleware chain is pre-compiled for zero-overhead execution
//   - Route matching uses efficient trie-based algorithm
//   - Template caching minimizes repeated parsing
//
// Thread Safety:
//   - App struct is safe for concurrent use after initialization
//   - Context objects are NOT thread-safe and should not be shared between goroutines
//   - Session operations are thread-safe when using appropriate storage backend
//
// Version information is loaded dynamically from cfg/app.yaml.
//
// websvrutil 패키지는 프로덕션 수준의 HTTP 서버를 구축하기 위한 매우 간단한 웹 서버 유틸리티를 제공합니다.
// 라우팅, 미들웨어, 컨텍스트 관리, 세션, CSRF 보호, 템플릿 렌더링, 그레이스풀 셧다운을 포함한
// 일반적인 서버 패턴을 처리하는 높은 수준의 사용하기 쉬운 API를 제공하여
// 웹 서버 설정의 복잡성을 극적으로 줄입니다.
//
// 이 패키지는 다음 원칙으로 설계되었습니다:
//   - 단순성: 50줄 이상의 보일러플레이트 서버 설정 코드를 단 5줄로 축소
//   - 안전성: 내장 CSRF 보호, 안전한 세션 관리, 패닉 복구
//   - 성능: 최적화된 컨텍스트 풀링, 효율적인 미들웨어 체인, 제로 할당 경로
//   - 유연성: 확장 가능한 미들웨어 시스템과 커스터마이징 가능한 옵션
//   - 프로덕션 준비: 그레이스풀 셧다운, 헬스 체크, 포괄적인 에러 처리
//
// 주요 기능:
//   - 앱 관리: App 구조체로 HTTP 서버 생명주기 생성 및 관리
//   - 라우팅: 경로 매개변수와 HTTP 메서드를 지원하는 유연한 라우트 등록
//   - 컨텍스트: Context 추상화를 통한 향상된 요청/응답 처리
//   - 미들웨어: 횡단 관심사를 위한 조합 가능한 미들웨어 체인
//   - 세션: 여러 저장소 백엔드를 지원하는 안전한 세션 관리
//   - CSRF: 내장 크로스 사이트 요청 위조 보호
//   - 템플릿: 레이아웃 지원을 갖춘 템플릿 엔진 통합
//   - 검증: 상세한 에러 메시지를 제공하는 요청 데이터 검증
//   - 그레이스풀 셧다운: 연결 드레이닝을 통한 안전한 서버 종료
//
// 기본 사용법:
//   app := websvrutil.New()
//   app.GET("/", func(c *websvrutil.Context) error {
//       return c.String(200, "Hello, World!")
//   })
//   app.Start(":8080")
//
// 미들웨어를 사용한 고급 사용법:
//   app := websvrutil.New(
//       websvrutil.WithLogger(logger),
//       websvrutil.WithCSRF(),
//       websvrutil.WithSession(sessionStore),
//   )
//   
//   // 공통 미들웨어로 라우트 그룹화
//   api := app.Group("/api", authMiddleware)
//   api.GET("/users", listUsers)
//   api.POST("/users", createUser)
//   
//   // 그레이스풀 셧다운으로 시작
//   app.StartWithGracefulShutdown(":8080", 30*time.Second)
//
// 성능 특성:
//   - 컨텍스트 풀링은 Context 객체를 재사용하여 GC 압력을 줄임
//   - 미들웨어 체인은 오버헤드 없는 실행을 위해 사전 컴파일됨
//   - 라우트 매칭은 효율적인 트라이 기반 알고리즘 사용
//   - 템플릿 캐싱은 반복적인 파싱을 최소화
//
// 스레드 안전성:
//   - App 구조체는 초기화 후 동시 사용이 안전함
//   - Context 객체는 스레드 안전하지 않으며 고루틴 간 공유하면 안 됨
//   - 세션 작업은 적절한 저장소 백엔드를 사용할 때 스레드 안전함
//
// 버전 정보는 cfg/app.yaml에서 동적으로 로드됩니다.
package websvrutil

import "github.com/arkd0ng/go-utils/internal/version"

// Version is the current version of the websvrutil package.
// It is automatically loaded from the cfg/app.yaml configuration file at package initialization time.
// If the version cannot be loaded (e.g., file not found or invalid format), it defaults to "v0.0.0-dev".
//
// The version string typically follows semantic versioning (e.g., "1.0.0", "2.1.3").
// This version information can be used for:
//   - Logging and monitoring: Include version in application logs
//   - Health checks: Expose version in health check endpoints
//   - Debugging: Identify which version is running in production
//   - API versioning: Use as part of API version negotiation
//
// Example usage:
//   log.Printf("Starting websvrutil version: %s", websvrutil.Version)
//
//   // In a health check endpoint:
//   func healthCheck(c *websvrutil.Context) error {
//       return c.JSON(200, map[string]string{
//           "status": "ok",
//           "version": websvrutil.Version,
//       })
//   }
//
// Note: The version is loaded once at package initialization and remains constant
// throughout the application lifecycle. It cannot be modified at runtime.
//
// Version은 websvrutil 패키지의 현재 버전입니다.
// 패키지 초기화 시점에 cfg/app.yaml 설정 파일에서 자동으로 로드됩니다.
// 버전을 로드할 수 없는 경우(예: 파일을 찾을 수 없거나 형식이 잘못됨), 기본값 "v0.0.0-dev"로 설정됩니다.
//
// 버전 문자열은 일반적으로 시맨틱 버전 관리를 따릅니다(예: "1.0.0", "2.1.3").
// 이 버전 정보는 다음 용도로 사용할 수 있습니다:
//   - 로깅 및 모니터링: 애플리케이션 로그에 버전 포함
//   - 헬스 체크: 헬스 체크 엔드포인트에서 버전 노출
//   - 디버깅: 프로덕션에서 실행 중인 버전 식별
//   - API 버전 관리: API 버전 협상의 일부로 사용
//
// 사용 예제:
//   log.Printf("Starting websvrutil version: %s", websvrutil.Version)
//
//   // 헬스 체크 엔드포인트에서:
//   func healthCheck(c *websvrutil.Context) error {
//       return c.JSON(200, map[string]string{
//           "status": "ok",
//           "version": websvrutil.Version,
//       })
//   }
//
// 주의: 버전은 패키지 초기화 시 한 번 로드되며 애플리케이션 생명주기 동안 일정하게 유지됩니다.
// 런타임에 수정할 수 없습니다.
var Version = version.Get()
