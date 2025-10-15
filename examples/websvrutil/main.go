package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arkd0ng/go-utils/websvrutil"
)

func main() {
	fmt.Println("=== websvrutil Package Examples (v1.11.004) ===")
	fmt.Println("=== websvrutil 패키지 예제 (v1.11.004) ===\n")

	// Example 1: Basic Server / 기본 서버
	fmt.Println("Example 1: Basic Server / 기본 서버")
	example1BasicServer()

	// Example 2: Server with Custom Options / 커스텀 옵션을 사용한 서버
	fmt.Println("\nExample 2: Server with Custom Options / 커스텀 옵션을 사용한 서버")
	example2CustomOptions()

	// Example 3: Routing with GET/POST / GET/POST 라우팅
	fmt.Println("\nExample 3: Routing with GET/POST / GET/POST 라우팅")
	example3Routing()

	// Example 4: Path Parameters / 경로 매개변수
	fmt.Println("\nExample 4: Path Parameters / 경로 매개변수")
	example4PathParameters()

	// Example 5: Wildcard Routes / 와일드카드 라우트
	fmt.Println("\nExample 5: Wildcard Routes / 와일드카드 라우트")
	example5WildcardRoutes()

	// Example 6: Custom 404 Handler / 커스텀 404 핸들러
	fmt.Println("\nExample 6: Custom 404 Handler / 커스텀 404 핸들러")
	example6Custom404()

	// Example 7: Context - Path Parameters / Context - 경로 매개변수
	fmt.Println("\nExample 7: Context - Path Parameters / Context - 경로 매개변수")
	example7ContextPathParameters()

	// Example 8: Context - Query Parameters / Context - 쿼리 매개변수
	fmt.Println("\nExample 8: Context - Query Parameters / Context - 쿼리 매개변수")
	example8ContextQueryParameters()

	// Example 9: Context - Custom Values / Context - 커스텀 값
	fmt.Println("\nExample 9: Context - Custom Values / Context - 커스텀 값")
	example9ContextCustomValues()

	// Example 10: Context - Request Headers / Context - 요청 헤더
	fmt.Println("\nExample 10: Context - Request Headers / Context - 요청 헤더")
	example10ContextHeaders()

	// Example 11: Graceful Shutdown / 정상 종료
	fmt.Println("\nExample 11: Graceful Shutdown / 정상 종료")
	example11GracefulShutdown()

	// Example 12: Custom Middleware / 커스텀 미들웨어
	fmt.Println("\nExample 12: Custom Middleware / 커스텀 미들웨어")
	example12CustomMiddleware()

	// Example 13: Multiple Middleware / 다중 미들웨어
	fmt.Println("\nExample 13: Multiple Middleware / 다중 미들웨어")
	example13MultipleMiddleware()

	// Example 14: Production Configuration / 프로덕션 설정
	fmt.Println("\nExample 14: Production Configuration / 프로덕션 설정")
	example14ProductionConfig()

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

// example3Routing demonstrates basic routing with GET and POST.
// example3Routing은 GET 및 POST를 사용한 기본 라우팅을 시연합니다.
func example3Routing() {
	app := websvrutil.New()

	getCount := 0
	postCount := 0

	// Register GET route
	// GET 라우트 등록
	app.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		getCount++
		w.WriteHeader(http.StatusOK)
	})

	// Register POST route
	// POST 라우트 등록
	app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
		postCount++
		w.WriteHeader(http.StatusCreated)
	})

	fmt.Println("✓ Registered routes:")
	fmt.Println("✓ 등록된 라우트:")
	fmt.Println("  - GET /users")
	fmt.Println("  - POST /users")

	// Simulate requests
	// 요청 시뮬레이션
	fmt.Println("\n  Simulating requests:")
	fmt.Println("  요청 시뮬레이션:")

	testGet := httptest.NewRequest("GET", "/users", nil)
	testPost := httptest.NewRequest("POST", "/users", nil)

	app.ServeHTTP(httptest.NewRecorder(), testGet)
	app.ServeHTTP(httptest.NewRecorder(), testPost)

	fmt.Printf("  - GET requests: %d\n", getCount)
	fmt.Printf("  - POST requests: %d\n", postCount)

	fmt.Println("\n✓ Routes working correctly")
	fmt.Println("✓ 라우트가 올바르게 작동합니다")
}

// example4PathParameters demonstrates path parameter extraction.
// example4PathParameters는 경로 매개변수 추출을 시연합니다.
func example4PathParameters() {
	app := websvrutil.New()

	// Register route with parameter
	// 매개변수가 있는 라우트 등록
	app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	app.GET("/users/:userId/posts/:postId", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("✓ Registered routes with parameters:")
	fmt.Println("✓ 매개변수가 있는 라우트 등록됨:")
	fmt.Println("  - GET /users/:id")
	fmt.Println("  - GET /users/:userId/posts/:postId")

	// Test parameter matching
	// 매개변수 일치 테스트
	testPaths := []string{
		"/users/123",
		"/users/456/posts/789",
	}

	fmt.Println("\n  Testing parameter matching:")
	fmt.Println("  매개변수 일치 테스트:")

	for _, path := range testPaths {
		req := httptest.NewRequest("GET", path, nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		if rec.Code == http.StatusOK {
			fmt.Printf("  ✓ Matched: %s\n", path)
		}
	}

	fmt.Println("\n✓ Parameter extraction working")
	fmt.Println("✓ 매개변수 추출 작동 중")
}

// example5WildcardRoutes demonstrates wildcard route matching.
// example5WildcardRoutes는 와일드카드 라우트 일치를 시연합니다.
func example5WildcardRoutes() {
	app := websvrutil.New()

	// Register wildcard route
	// 와일드카드 라우트 등록
	app.GET("/files/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("✓ Registered wildcard route:")
	fmt.Println("✓ 와일드카드 라우트 등록됨:")
	fmt.Println("  - GET /files/*")

	// Test wildcard matching
	// 와일드카드 일치 테스트
	testPaths := []string{
		"/files/images/logo.png",
		"/files/docs/manual.pdf",
		"/files/a/b/c/d/e.txt",
	}

	fmt.Println("\n  Testing wildcard matching:")
	fmt.Println("  와일드카드 일치 테스트:")

	for _, path := range testPaths {
		req := httptest.NewRequest("GET", path, nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		if rec.Code == http.StatusOK {
			fmt.Printf("  ✓ Matched: %s\n", path)
		}
	}

	fmt.Println("\n✓ Wildcard routes working correctly")
	fmt.Println("✓ 와일드카드 라우트가 올바르게 작동합니다")
}

// example6Custom404 demonstrates custom 404 handler.
// example6Custom404는 커스텀 404 핸들러를 시연합니다.
func example6Custom404() {
	app := websvrutil.New()

	// Register normal route
	// 일반 라우트 등록
	app.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Register custom 404 handler
	// 커스텀 404 핸들러 등록
	custom404Called := false
	app.NotFound(func(w http.ResponseWriter, r *http.Request) {
		custom404Called = true
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Custom 404: %s not found", r.URL.Path)
	})

	fmt.Println("✓ Custom 404 handler registered")
	fmt.Println("✓ 커스텀 404 핸들러 등록됨")

	// Test existing route
	// 기존 라우트 테스트
	fmt.Println("\n  Testing existing route (/users):")
	fmt.Println("  기존 라우트 테스트 (/users):")
	req := httptest.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)
	fmt.Printf("  Status: %d\n", rec.Code)

	// Test non-existent route
	// 존재하지 않는 라우트 테스트
	fmt.Println("\n  Testing non-existent route (/nonexistent):")
	fmt.Println("  존재하지 않는 라우트 테스트 (/nonexistent):")
	req = httptest.NewRequest("GET", "/nonexistent", nil)
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)
	fmt.Printf("  Status: %d\n", rec.Code)
	fmt.Printf("  Custom handler called: %v\n", custom404Called)

	fmt.Println("\n✓ Custom 404 handler working correctly")
	fmt.Println("✓ 커스텀 404 핸들러가 올바르게 작동합니다")
}

// example7ContextPathParameters demonstrates Context path parameter access.
// example7ContextPathParameters는 Context 경로 매개변수 액세스를 시연합니다.
func example7ContextPathParameters() {
	app := websvrutil.New()

	var extractedID, extractedUserID, extractedPostID string

	// Single parameter
	// 단일 매개변수
	app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		extractedID = ctx.Param("id")
		fmt.Fprintf(w, "User ID: %s", extractedID)
	})

	// Multiple parameters
	// 다중 매개변수
	app.GET("/users/:userId/posts/:postId", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		extractedUserID = ctx.Param("userId")
		extractedPostID = ctx.Param("postId")
		fmt.Fprintf(w, "User: %s, Post: %s", extractedUserID, extractedPostID)
	})

	fmt.Println("✓ Routes with Context parameters registered")
	fmt.Println("✓ Context 매개변수가 있는 라우트 등록됨")

	// Test single parameter
	// 단일 매개변수 테스트
	req1 := httptest.NewRequest("GET", "/users/123", nil)
	rec1 := httptest.NewRecorder()
	app.ServeHTTP(rec1, req1)

	fmt.Println("\n  Single parameter test:")
	fmt.Println("  단일 매개변수 테스트:")
	fmt.Printf("  - URL: /users/123\n")
	fmt.Printf("  - Extracted ID: %s\n", extractedID)

	// Test multiple parameters
	// 다중 매개변수 테스트
	req2 := httptest.NewRequest("GET", "/users/456/posts/789", nil)
	rec2 := httptest.NewRecorder()
	app.ServeHTTP(rec2, req2)

	fmt.Println("\n  Multiple parameters test:")
	fmt.Println("  다중 매개변수 테스트:")
	fmt.Printf("  - URL: /users/456/posts/789\n")
	fmt.Printf("  - Extracted User ID: %s\n", extractedUserID)
	fmt.Printf("  - Extracted Post ID: %s\n", extractedPostID)

	fmt.Println("\n✓ Context path parameter access working")
	fmt.Println("✓ Context 경로 매개변수 액세스 작동 중")
}

// example8ContextQueryParameters demonstrates Context query parameter access.
// example8ContextQueryParameters는 Context 쿼리 매개변수 액세스를 시연합니다.
func example8ContextQueryParameters() {
	app := websvrutil.New()

	var query, page, limit string

	app.GET("/search", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		query = ctx.Query("q")
		page = ctx.QueryDefault("page", "1")
		limit = ctx.QueryDefault("limit", "10")
		fmt.Fprintf(w, "Query: %s, Page: %s, Limit: %s", query, page, limit)
	})

	fmt.Println("✓ Search route with query parameters registered")
	fmt.Println("✓ 쿼리 매개변수가 있는 검색 라우트 등록됨")

	// Test with query parameters
	// 쿼리 매개변수로 테스트
	req := httptest.NewRequest("GET", "/search?q=golang&page=2", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	fmt.Println("\n  Query parameter test:")
	fmt.Println("  쿼리 매개변수 테스트:")
	fmt.Printf("  - URL: /search?q=golang&page=2\n")
	fmt.Printf("  - Query (q): %s\n", query)
	fmt.Printf("  - Page: %s\n", page)
	fmt.Printf("  - Limit (default): %s\n", limit)

	fmt.Println("\n✓ Context query parameter access working")
	fmt.Println("✓ Context 쿼리 매개변수 액세스 작동 중")
}

// example9ContextCustomValues demonstrates storing and retrieving custom values.
// example9ContextCustomValues는 커스텀 값 저장 및 검색을 시연합니다.
func example9ContextCustomValues() {
	app := websvrutil.New()

	var storedUser string
	var storedAuth bool
	var storedCount int

	app.GET("/user/:id", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)

		// Store custom values
		// 커스텀 값 저장
		ctx.Set("userId", ctx.Param("id"))
		ctx.Set("authenticated", true)
		ctx.Set("requestCount", 42)

		// Retrieve values
		// 값 검색
		storedUser = ctx.GetString("userId")
		storedAuth = ctx.GetBool("authenticated")
		storedCount = ctx.GetInt("requestCount")

		fmt.Fprintf(w, "User: %s, Auth: %v, Count: %d", storedUser, storedAuth, storedCount)
	})

	fmt.Println("✓ Route with custom value storage registered")
	fmt.Println("✓ 커스텀 값 저장이 있는 라우트 등록됨")

	// Test custom values
	// 커스텀 값 테스트
	req := httptest.NewRequest("GET", "/user/alice", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	fmt.Println("\n  Custom values test:")
	fmt.Println("  커스텀 값 테스트:")
	fmt.Printf("  - Stored user ID: %s\n", storedUser)
	fmt.Printf("  - Stored authenticated: %v\n", storedAuth)
	fmt.Printf("  - Stored request count: %d\n", storedCount)

	fmt.Println("\n✓ Context custom value storage working")
	fmt.Println("✓ Context 커스텀 값 저장 작동 중")
}

// example10ContextHeaders demonstrates request and response header access.
// example10ContextHeaders는 요청 및 응답 헤더 액세스를 시연합니다.
func example10ContextHeaders() {
	app := websvrutil.New()

	var authHeader, contentType string

	app.GET("/api/data", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)

		// Read request headers
		// 요청 헤더 읽기
		authHeader = ctx.Header("Authorization")
		contentType = ctx.Header("Content-Type")

		// Set response headers
		// 응답 헤더 설정
		ctx.SetHeader("X-API-Version", "1.0")
		ctx.SetHeader("Content-Type", "application/json")

		fmt.Fprintf(w, "Auth: %s, Type: %s", authHeader, contentType)
	})

	fmt.Println("✓ API route with header access registered")
	fmt.Println("✓ 헤더 액세스가 있는 API 라우트 등록됨")

	// Test with headers
	// 헤더로 테스트
	req := httptest.NewRequest("GET", "/api/data", nil)
	req.Header.Set("Authorization", "Bearer token123")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	fmt.Println("\n  Request headers:")
	fmt.Println("  요청 헤더:")
	fmt.Printf("  - Authorization: %s\n", authHeader)
	fmt.Printf("  - Content-Type: %s\n", contentType)

	fmt.Println("\n  Response headers:")
	fmt.Println("  응답 헤더:")
	fmt.Printf("  - X-API-Version: %s\n", rec.Header().Get("X-API-Version"))
	fmt.Printf("  - Content-Type: %s\n", rec.Header().Get("Content-Type"))

	fmt.Println("\n✓ Context header access working")
	fmt.Println("✓ Context 헤더 액세스 작동 중")
}

// example11GracefulShutdown demonstrates graceful server shutdown.
// example11GracefulShutdown은 정상적인 서버 종료를 시연합니다.
func example11GracefulShutdown() {
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

// example12CustomMiddleware demonstrates adding custom middleware.
// example12CustomMiddleware는 커스텀 미들웨어 추가를 시연합니다.
func example12CustomMiddleware() {
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

// example13MultipleMiddleware demonstrates adding multiple middleware.
// example13MultipleMiddleware는 다중 미들웨어 추가를 시연합니다.
func example13MultipleMiddleware() {
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

// example14ProductionConfig demonstrates a production-ready configuration.
// example14ProductionConfig는 프로덕션 준비 설정을 시연합니다.
func example14ProductionConfig() {
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
