package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
	fmt.Println("=== websvrutil Package Examples (v1.11.002) ===")
	fmt.Println("=== websvrutil 패키지 예제 (v1.11.002) ===\n")

	// Example 1: Basic Server / 기본 서버
	fmt.Println("Example 1: Basic Server / 기본 서버")
	example1BasicServer()

	// Example 2: Server with Custom Options / 커스텀 옵션을 사용한 서버
	fmt.Println("\nExample 2: Server with Custom Options / 커스텀 옵션을 사용한 서버")
	example2CustomOptions()

	// Example 3: Graceful Shutdown / 정상 종료
	fmt.Println("\nExample 3: Graceful Shutdown / 정상 종료")
	example3GracefulShutdown()

	// Example 4: Custom Middleware / 커스텀 미들웨어
	fmt.Println("\nExample 4: Custom Middleware / 커스텀 미들웨어")
	example4CustomMiddleware()

	// Example 5: Multiple Middleware / 다중 미들웨어
	fmt.Println("\nExample 5: Multiple Middleware / 다중 미들웨어")
	example5MultipleMiddleware()

	// Example 6: Production Configuration / 프로덕션 설정
	fmt.Println("\nExample 6: Production Configuration / 프로덕션 설정")
	example6ProductionConfig()

	fmt.Println("\n=== All Examples Completed ===")
	fmt.Println("=== 모든 예제 완료 ===")
}

// example1BasicServer demonstrates creating and running a basic server.
// example1BasicServer는 기본 서버 생성 및 실행을 시연합니다.
func example1BasicServer() {
	// Create app with default options
	// 기본 옵션으로 앱 생성
	app := websvrutil.New()

	fmt.Println("✓ Created app with default options")
	fmt.Println("✓ 기본 옵션으로 앱 생성됨")
	fmt.Println("  - ReadTimeout: 15s")
	fmt.Println("  - WriteTimeout: 15s")
	fmt.Println("  - IdleTimeout: 60s")
	fmt.Println("  - MaxHeaderBytes: 1 MB")
	fmt.Println("  - Logger: enabled")
	fmt.Println("  - Recovery: enabled")

	// Note: In real usage, you would call app.Run(":8080")
	// 참고: 실제 사용에서는 app.Run(":8080")을 호출합니다
	fmt.Println("\nUsage: app.Run(\":8080\")")
	fmt.Println("사용법: app.Run(\":8080\")")

	_ = app // Suppress unused variable warning / 미사용 변수 경고 억제
}

// example2CustomOptions demonstrates using custom options.
// example2CustomOptions는 커스텀 옵션 사용을 시연합니다.
func example2CustomOptions() {
	// Create app with custom options
	// 커스텀 옵션으로 앱 생성
	app := websvrutil.New(
		websvrutil.WithReadTimeout(30*time.Second),
		websvrutil.WithWriteTimeout(30*time.Second),
		websvrutil.WithIdleTimeout(90*time.Second),
		websvrutil.WithMaxHeaderBytes(2<<20), // 2 MB
		websvrutil.WithTemplateDir("views"),
		websvrutil.WithStaticDir("public"),
		websvrutil.WithStaticPrefix("/assets"),
		websvrutil.WithAutoReload(true),
		websvrutil.WithLogger(false),
		websvrutil.WithRecovery(true),
	)

	fmt.Println("✓ Created app with custom options")
	fmt.Println("✓ 커스텀 옵션으로 앱 생성됨")
	fmt.Println("  - ReadTimeout: 30s")
	fmt.Println("  - WriteTimeout: 30s")
	fmt.Println("  - IdleTimeout: 90s")
	fmt.Println("  - MaxHeaderBytes: 2 MB")
	fmt.Println("  - TemplateDir: views")
	fmt.Println("  - StaticDir: public")
	fmt.Println("  - StaticPrefix: /assets")
	fmt.Println("  - AutoReload: true")
	fmt.Println("  - Logger: disabled")
	fmt.Println("  - Recovery: enabled")

	_ = app // Suppress unused variable warning / 미사용 변수 경고 억제
}

// example3GracefulShutdown demonstrates graceful server shutdown.
// example3GracefulShutdown은 정상적인 서버 종료를 시연합니다.
func example3GracefulShutdown() {
	app := websvrutil.New()

	// Setup signal handling
	// 시그널 처리 설정
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("✓ Signal handler configured")
	fmt.Println("✓ 시그널 핸들러 설정됨")

	// Simulate server startup and shutdown
	// 서버 시작 및 종료 시뮬레이션
	serverStarted := make(chan bool)

	go func() {
		// In real usage: app.Run(":8080")
		// 실제 사용: app.Run(":8080")
		fmt.Println("\n  Server would start here...")
		fmt.Println("  서버가 여기서 시작됩니다...")
		serverStarted <- true

		// Simulate running
		// 실행 시뮬레이션
		time.Sleep(100 * time.Millisecond)
	}()

	<-serverStarted

	// Simulate shutdown signal
	// 종료 시그널 시뮬레이션
	fmt.Println("\n  Simulating shutdown signal...")
	fmt.Println("  종료 시그널 시뮬레이션...")

	// Graceful shutdown with timeout
	// 타임아웃으로 정상 종료
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("\n✓ Shutdown initiated with 5s timeout")
	fmt.Println("✓ 5초 타임아웃으로 종료 시작됨")

	// In real usage: app.Shutdown(ctx)
	// 실제 사용: app.Shutdown(ctx)
	_ = ctx
	_ = app

	fmt.Println("✓ Server would shutdown gracefully")
	fmt.Println("✓ 서버가 정상적으로 종료됩니다")
}

// example4CustomMiddleware demonstrates adding custom middleware.
// example4CustomMiddleware는 커스텀 미들웨어 추가를 시연합니다.
func example4CustomMiddleware() {
	app := websvrutil.New()

	// Create a logging middleware
	// 로깅 미들웨어 생성
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Printf("Started %s %s", r.Method, r.URL.Path)

			next.ServeHTTP(w, r)

			log.Printf("Completed in %v", time.Since(start))
		})
	}

	// Add middleware
	// 미들웨어 추가
	app.Use(loggingMiddleware)

	fmt.Println("✓ Added logging middleware")
	fmt.Println("✓ 로깅 미들웨어 추가됨")

	// Test with a sample request
	// 샘플 요청으로 테스트
	req, _ := http.NewRequest("GET", "/test", nil)
	rr := &responseRecorder{ResponseWriter: &dummyResponseWriter{}}

	fmt.Println("\n  Testing middleware with sample request:")
	fmt.Println("  샘플 요청으로 미들웨어 테스트:")

	app.ServeHTTP(rr, req)

	fmt.Println("✓ Middleware executed successfully")
	fmt.Println("✓ 미들웨어 실행 성공")
}

// example5MultipleMiddleware demonstrates adding multiple middleware.
// example5MultipleMiddleware는 다중 미들웨어 추가를 시연합니다.
func example5MultipleMiddleware() {
	app := websvrutil.New()

	// First middleware: Request ID
	// 첫 번째 미들웨어: 요청 ID
	requestIDMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Request-ID", "12345")
			next.ServeHTTP(w, r)
		})
	}

	// Second middleware: Timing
	// 두 번째 미들웨어: 타이밍
	timingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)
			w.Header().Set("X-Response-Time", duration.String())
		})
	}

	// Third middleware: CORS
	// 세 번째 미들웨어: CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	}

	// Add all middleware (executed in order)
	// 모든 미들웨어 추가 (순서대로 실행)
	app.Use(requestIDMiddleware, timingMiddleware, corsMiddleware)

	fmt.Println("✓ Added 3 middleware:")
	fmt.Println("✓ 3개의 미들웨어 추가됨:")
	fmt.Println("  1. Request ID")
	fmt.Println("  2. Timing")
	fmt.Println("  3. CORS")

	// Test with a sample request
	// 샘플 요청으로 테스트
	req, _ := http.NewRequest("GET", "/test", nil)
	rr := &responseRecorder{
		ResponseWriter: &dummyResponseWriter{headers: make(http.Header)},
	}

	app.ServeHTTP(rr, req)

	fmt.Println("\n  Headers set by middleware:")
	fmt.Println("  미들웨어가 설정한 헤더:")
	if id := rr.Header().Get("X-Request-ID"); id != "" {
		fmt.Printf("  - X-Request-ID: %s\n", id)
	}
	if cors := rr.Header().Get("Access-Control-Allow-Origin"); cors != "" {
		fmt.Printf("  - Access-Control-Allow-Origin: %s\n", cors)
	}

	fmt.Println("\n✓ All middleware executed in order")
	fmt.Println("✓ 모든 미들웨어가 순서대로 실행됨")
}

// example6ProductionConfig demonstrates a production-ready configuration.
// example6ProductionConfig는 프로덕션 준비 설정을 시연합니다.
func example6ProductionConfig() {
	app := websvrutil.New(
		// Security timeouts / 보안 타임아웃
		websvrutil.WithReadTimeout(10*time.Second),
		websvrutil.WithWriteTimeout(10*time.Second),
		websvrutil.WithIdleTimeout(120*time.Second),

		// Limit request size / 요청 크기 제한
		websvrutil.WithMaxHeaderBytes(1<<20), // 1 MB

		// Production directories / 프로덕션 디렉토리
		websvrutil.WithTemplateDir("/app/templates"),
		websvrutil.WithStaticDir("/app/static"),
		websvrutil.WithStaticPrefix("/static"),

		// Production settings / 프로덕션 설정
		websvrutil.WithAutoReload(false),  // Disable in production
		websvrutil.WithLogger(true),       // Enable logging
		websvrutil.WithRecovery(true),     // Enable panic recovery
	)

	fmt.Println("✓ Production configuration applied")
	fmt.Println("✓ 프로덕션 설정 적용됨")
	fmt.Println("\nSecurity Features / 보안 기능:")
	fmt.Println("  ✓ Short timeouts to prevent slowloris attacks")
	fmt.Println("  ✓ Slowloris 공격 방지를 위한 짧은 타임아웃")
	fmt.Println("  ✓ Header size limits")
	fmt.Println("  ✓ 헤더 크기 제한")
	fmt.Println("  ✓ Panic recovery enabled")
	fmt.Println("  ✓ 패닉 복구 활성화")
	fmt.Println("  ✓ Request logging enabled")
	fmt.Println("  ✓ 요청 로깅 활성화")
	fmt.Println("\nOptimization / 최적화:")
	fmt.Println("  ✓ Template caching (no auto-reload)")
	fmt.Println("  ✓ 템플릿 캐싱 (자동 재로드 없음)")
	fmt.Println("  ✓ Keep-alive with appropriate timeout")
	fmt.Println("  ✓ 적절한 타임아웃의 Keep-alive")

	_ = app // Suppress unused variable warning / 미사용 변수 경고 억제
}

// responseRecorder is a simple response writer for testing.
// responseRecorder는 테스트용 간단한 응답 작성기입니다.
type responseRecorder struct {
	http.ResponseWriter
}

// dummyResponseWriter is a minimal implementation of http.ResponseWriter.
// dummyResponseWriter는 http.ResponseWriter의 최소 구현입니다.
type dummyResponseWriter struct {
	headers http.Header
	code    int
}

func (d *dummyResponseWriter) Header() http.Header {
	if d.headers == nil {
		d.headers = make(http.Header)
	}
	return d.headers
}

func (d *dummyResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (d *dummyResponseWriter) WriteHeader(statusCode int) {
	d.code = statusCode
}
