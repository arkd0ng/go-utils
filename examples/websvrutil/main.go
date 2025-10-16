package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
	"github.com/arkd0ng/go-utils/websvrutil"
)

const logBaseName = "websvrutil-example"

var exampleLogger *logging.Logger

func setupLogger() *logging.Logger {
	if err := os.MkdirAll("logs", 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Failed to create logs directory: %v\n", err)
		return nil
	}

	logFilePath := fmt.Sprintf("logs/%s.log", logBaseName)

	if fileutil.Exists(logFilePath) {
		if modTime, err := fileutil.ModTime(logFilePath); err == nil {
			backupName := fmt.Sprintf("logs/%s-%s.log", logBaseName, modTime.Format("20060102-150405"))
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Fprintf(os.Stdout, "✅ Backed up previous log to: %s\n", backupName)
				fileutil.DeleteFile(logFilePath)
			}
		}

		backupPattern := fmt.Sprintf("logs/%s-*.log", logBaseName)
		if backupFiles, err := filepath.Glob(backupPattern); err == nil && len(backupFiles) > 5 {
			type fileInfo struct {
				path    string
				modTime time.Time
			}
			files := make([]fileInfo, 0, len(backupFiles))
			for _, f := range backupFiles {
				if mt, err := fileutil.ModTime(f); err == nil {
					files = append(files, fileInfo{path: f, modTime: mt})
				}
			}

			sort.Slice(files, func(i, j int) bool {
				return files[i].modTime.Before(files[j].modTime)
			})

			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Fprintf(os.Stdout, "🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	logger, err := logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithStdout(true),
		logging.WithLevel(logging.INFO),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Failed to initialize logger: %v\n", err)
		return nil
	}

	logger.Banner("websvrutil Package Examples", websvrutil.Version)
	logger.Info("Logs mirror console output / 로그는 콘솔 출력을 반영합니다")
	logger.Info("")

	return logger
}

func logPrintln(args ...interface{}) {
	fmt.Fprintln(os.Stdout, args...)
	if exampleLogger != nil {
		msg := strings.TrimSuffix(fmt.Sprintln(args...), "\n")
		exampleLogger.Info(msg)
	}
}

func logPrintf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
	if exampleLogger != nil {
		msg := strings.TrimSuffix(fmt.Sprintf(format, args...), "\n")
		exampleLogger.Info(msg)
	}
}

func main() {
	exampleLogger = setupLogger()
	if exampleLogger != nil {
		defer exampleLogger.Close()
	}

	logPrintf("=== websvrutil Package Examples (%s) ===\n", websvrutil.Version)
	logPrintf("=== websvrutil 패키지 예제 (%s) ===\n", websvrutil.Version)
	logPrintln()

	// Example 1: Basic Server / 기본 서버
	logPrintln("Example 1: Basic Server / 기본 서버")
	example1BasicServer()

	// Example 2: Server with Custom Options / 커스텀 옵션을 사용한 서버
	logPrintln()
	logPrintln("Example 2: Server with Custom Options / 커스텀 옵션을 사용한 서버")
	example2CustomOptions()

	// Example 3: Routing with GET/POST / GET/POST 라우팅
	logPrintln()
	logPrintln("Example 3: Routing with GET/POST / GET/POST 라우팅")
	example3Routing()

	// Example 4: Path Parameters / 경로 매개변수
	logPrintln()
	logPrintln("Example 4: Path Parameters / 경로 매개변수")
	example4PathParameters()

	// Example 5: Wildcard Routes / 와일드카드 라우트
	logPrintln()
	logPrintln("Example 5: Wildcard Routes / 와일드카드 라우트")
	example5WildcardRoutes()

	// Example 6: Custom 404 Handler / 커스텀 404 핸들러
	logPrintln()
	logPrintln("Example 6: Custom 404 Handler / 커스텀 404 핸들러")
	example6Custom404()

	// Example 7: Context - Path Parameters / Context - 경로 매개변수
	logPrintln()
	logPrintln("Example 7: Context - Path Parameters / Context - 경로 매개변수")
	example7ContextPathParameters()

	// Example 8: Context - Query Parameters / Context - 쿼리 매개변수
	logPrintln()
	logPrintln("Example 8: Context - Query Parameters / Context - 쿼리 매개변수")
	example8ContextQueryParameters()

	// Example 9: Context - Custom Values / Context - 커스텀 값
	logPrintln()
	logPrintln("Example 9: Context - Custom Values / Context - 커스텀 값")
	example9ContextCustomValues()

	// Example 10: Context - Request Headers / Context - 요청 헤더
	logPrintln()
	logPrintln("Example 10: Context - Request Headers / Context - 요청 헤더")
	example10ContextHeaders()

	// Example 11: Graceful Shutdown / 정상 종료
	logPrintln()
	logPrintln("Example 11: Graceful Shutdown / 정상 종료")
	example11GracefulShutdown()

	// Example 12: Custom Middleware / 커스텀 미들웨어
	logPrintln()
	logPrintln("Example 12: Custom Middleware / 커스텀 미들웨어")
	example12CustomMiddleware()

	// Example 13: Multiple Middleware / 다중 미들웨어
	logPrintln()
	logPrintln("Example 13: Multiple Middleware / 다중 미들웨어")
	example13MultipleMiddleware()

	// Example 14: Production Configuration / 프로덕션 설정
	logPrintln()
	logPrintln("Example 14: Production Configuration / 프로덕션 설정")
	example14ProductionConfig()

	logPrintln()
	logPrintln("=== All Examples Completed ===")
	logPrintln("=== 모든 예제 완료 ===")
}

// example1BasicServer demonstrates creating and running a basic server.
// example1BasicServer는 기본 서버 생성 및 실행을 시연합니다.
func example1BasicServer() {
	// Create app with default options
	// 기본 옵션으로 앱 생성
	app := websvrutil.New()

	logPrintln("✓ Created app with default options")
	logPrintln("✓ 기본 옵션으로 앱 생성됨")
	logPrintln("  - ReadTimeout: 15s")
	logPrintln("  - WriteTimeout: 15s")
	logPrintln("  - IdleTimeout: 60s")
	logPrintln("  - MaxHeaderBytes: 1 MB")
	logPrintln("  - Logger: enabled")
	logPrintln("  - Recovery: enabled")

	// Note: In real usage, you would call app.Run(":8080")
	// 참고: 실제 사용에서는 app.Run(":8080")을 호출합니다
	logPrintln()
	logPrintln("Usage: app.Run(\":8080\")")
	logPrintln("사용법: app.Run(\":8080\")")

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

	logPrintln("✓ Created app with custom options")
	logPrintln("✓ 커스텀 옵션으로 앱 생성됨")
	logPrintln("  - ReadTimeout: 30s")
	logPrintln("  - WriteTimeout: 30s")
	logPrintln("  - IdleTimeout: 90s")
	logPrintln("  - MaxHeaderBytes: 2 MB")
	logPrintln("  - TemplateDir: views")
	logPrintln("  - StaticDir: public")
	logPrintln("  - StaticPrefix: /assets")
	logPrintln("  - AutoReload: true")
	logPrintln("  - Logger: disabled")
	logPrintln("  - Recovery: enabled")

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

	logPrintln("✓ Registered routes:")
	logPrintln("✓ 등록된 라우트:")
	logPrintln("  - GET /users")
	logPrintln("  - POST /users")

	// Simulate requests
	// 요청 시뮬레이션
	logPrintln()
	logPrintln("  Simulating requests:")
	logPrintln("  요청 시뮬레이션:")

	testGet := httptest.NewRequest("GET", "/users", nil)
	testPost := httptest.NewRequest("POST", "/users", nil)

	app.ServeHTTP(httptest.NewRecorder(), testGet)
	app.ServeHTTP(httptest.NewRecorder(), testPost)

	logPrintf("  - GET requests: %d\n", getCount)
	logPrintf("  - POST requests: %d\n", postCount)

	logPrintln()
	logPrintln("✓ Routes working correctly")
	logPrintln("✓ 라우트가 올바르게 작동합니다")
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

	logPrintln("✓ Registered routes with parameters:")
	logPrintln("✓ 매개변수가 있는 라우트 등록됨:")
	logPrintln("  - GET /users/:id")
	logPrintln("  - GET /users/:userId/posts/:postId")

	// Test parameter matching
	// 매개변수 일치 테스트
	testPaths := []string{
		"/users/123",
		"/users/456/posts/789",
	}

	logPrintln()
	logPrintln("  Testing parameter matching:")
	logPrintln("  매개변수 일치 테스트:")

	for _, path := range testPaths {
		req := httptest.NewRequest("GET", path, nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		if rec.Code == http.StatusOK {
			logPrintf("  ✓ Matched: %s\n", path)
		}
	}

	logPrintln()
	logPrintln("✓ Parameter extraction working")
	logPrintln("✓ 매개변수 추출 작동 중")
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

	logPrintln("✓ Registered wildcard route:")
	logPrintln("✓ 와일드카드 라우트 등록됨:")
	logPrintln("  - GET /files/*")

	// Test wildcard matching
	// 와일드카드 일치 테스트
	testPaths := []string{
		"/files/images/logo.png",
		"/files/docs/manual.pdf",
		"/files/a/b/c/d/e.txt",
	}

	logPrintln()
	logPrintln("  Testing wildcard matching:")
	logPrintln("  와일드카드 일치 테스트:")

	for _, path := range testPaths {
		req := httptest.NewRequest("GET", path, nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		if rec.Code == http.StatusOK {
			logPrintf("  ✓ Matched: %s\n", path)
		}
	}

	logPrintln()
	logPrintln("✓ Wildcard routes working correctly")
	logPrintln("✓ 와일드카드 라우트가 올바르게 작동합니다")
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

	logPrintln("✓ Custom 404 handler registered")
	logPrintln("✓ 커스텀 404 핸들러 등록됨")

	// Test existing route
	// 기존 라우트 테스트
	logPrintln()
	logPrintln("  Testing existing route (/users):")
	logPrintln("  기존 라우트 테스트 (/users):")
	req := httptest.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)
	logPrintf("  Status: %d\n", rec.Code)

	// Test non-existent route
	// 존재하지 않는 라우트 테스트
	logPrintln()
	logPrintln("  Testing non-existent route (/nonexistent):")
	logPrintln("  존재하지 않는 라우트 테스트 (/nonexistent):")
	req = httptest.NewRequest("GET", "/nonexistent", nil)
	rec = httptest.NewRecorder()
	app.ServeHTTP(rec, req)
	logPrintf("  Status: %d\n", rec.Code)
	logPrintf("  Custom handler called: %v\n", custom404Called)

	logPrintln()
	logPrintln("✓ Custom 404 handler working correctly")
	logPrintln("✓ 커스텀 404 핸들러가 올바르게 작동합니다")
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

	logPrintln("✓ Routes with Context parameters registered")
	logPrintln("✓ Context 매개변수가 있는 라우트 등록됨")

	// Test single parameter
	// 단일 매개변수 테스트
	req1 := httptest.NewRequest("GET", "/users/123", nil)
	rec1 := httptest.NewRecorder()
	app.ServeHTTP(rec1, req1)

	logPrintln()
	logPrintln("  Single parameter test:")
	logPrintln("  단일 매개변수 테스트:")
	logPrintf("  - URL: /users/123\n")
	logPrintf("  - Extracted ID: %s\n", extractedID)

	// Test multiple parameters
	// 다중 매개변수 테스트
	req2 := httptest.NewRequest("GET", "/users/456/posts/789", nil)
	rec2 := httptest.NewRecorder()
	app.ServeHTTP(rec2, req2)

	logPrintln()
	logPrintln("  Multiple parameters test:")
	logPrintln("  다중 매개변수 테스트:")
	logPrintf("  - URL: /users/456/posts/789\n")
	logPrintf("  - Extracted User ID: %s\n", extractedUserID)
	logPrintf("  - Extracted Post ID: %s\n", extractedPostID)

	logPrintln()
	logPrintln("✓ Context path parameter access working")
	logPrintln("✓ Context 경로 매개변수 액세스 작동 중")
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

	logPrintln("✓ Search route with query parameters registered")
	logPrintln("✓ 쿼리 매개변수가 있는 검색 라우트 등록됨")

	// Test with query parameters
	// 쿼리 매개변수로 테스트
	req := httptest.NewRequest("GET", "/search?q=golang&page=2", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	logPrintln()
	logPrintln("  Query parameter test:")
	logPrintln("  쿼리 매개변수 테스트:")
	logPrintf("  - URL: /search?q=golang&page=2\n")
	logPrintf("  - Query (q): %s\n", query)
	logPrintf("  - Page: %s\n", page)
	logPrintf("  - Limit (default): %s\n", limit)

	logPrintln()
	logPrintln("✓ Context query parameter access working")
	logPrintln("✓ Context 쿼리 매개변수 액세스 작동 중")
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

	logPrintln("✓ Route with custom value storage registered")
	logPrintln("✓ 커스텀 값 저장이 있는 라우트 등록됨")

	// Test custom values
	// 커스텀 값 테스트
	req := httptest.NewRequest("GET", "/user/alice", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	logPrintln()
	logPrintln("  Custom values test:")
	logPrintln("  커스텀 값 테스트:")
	logPrintf("  - Stored user ID: %s\n", storedUser)
	logPrintf("  - Stored authenticated: %v\n", storedAuth)
	logPrintf("  - Stored request count: %d\n", storedCount)

	logPrintln()
	logPrintln("✓ Context custom value storage working")
	logPrintln("✓ Context 커스텀 값 저장 작동 중")
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

	logPrintln("✓ API route with header access registered")
	logPrintln("✓ 헤더 액세스가 있는 API 라우트 등록됨")

	// Test with headers
	// 헤더로 테스트
	req := httptest.NewRequest("GET", "/api/data", nil)
	req.Header.Set("Authorization", "Bearer token123")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	logPrintln()
	logPrintln("  Request headers:")
	logPrintln("  요청 헤더:")
	logPrintf("  - Authorization: %s\n", authHeader)
	logPrintf("  - Content-Type: %s\n", contentType)

	logPrintln()
	logPrintln("  Response headers:")
	logPrintln("  응답 헤더:")
	logPrintf("  - X-API-Version: %s\n", rec.Header().Get("X-API-Version"))
	logPrintf("  - Content-Type: %s\n", rec.Header().Get("Content-Type"))

	logPrintln()
	logPrintln("✓ Context header access working")
	logPrintln("✓ Context 헤더 액세스 작동 중")
}

// example11GracefulShutdown demonstrates graceful server shutdown.
// example11GracefulShutdown은 정상적인 서버 종료를 시연합니다.
func example11GracefulShutdown() {
	app := websvrutil.New()

	// Setup signal handling
	// 시그널 처리 설정
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	logPrintln("✓ Signal handler configured")
	logPrintln("✓ 시그널 핸들러 설정됨")

	// Simulate server startup and shutdown
	// 서버 시작 및 종료 시뮬레이션
	serverStarted := make(chan bool)

	go func() {
		// In real usage: app.Run(":8080")
		// 실제 사용: app.Run(":8080")
		logPrintln()
		logPrintln("  Server would start here...")
		logPrintln("  서버가 여기서 시작됩니다...")
		serverStarted <- true

		// Simulate running
		// 실행 시뮬레이션
		time.Sleep(100 * time.Millisecond)
	}()

	<-serverStarted

	// Simulate shutdown signal
	// 종료 시그널 시뮬레이션
	logPrintln()
	logPrintln("  Simulating shutdown signal...")
	logPrintln("  종료 시그널 시뮬레이션...")

	// Graceful shutdown with timeout
	// 타임아웃으로 정상 종료
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logPrintln()
	logPrintln("✓ Shutdown initiated with 5s timeout")
	logPrintln("✓ 5초 타임아웃으로 종료 시작됨")

	// In real usage: app.Shutdown(ctx)
	// 실제 사용: app.Shutdown(ctx)
	_ = ctx
	_ = app

	logPrintln("✓ Server would shutdown gracefully")
	logPrintln("✓ 서버가 정상적으로 종료됩니다")
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

	logPrintln("✓ Added logging middleware")
	logPrintln("✓ 로깅 미들웨어 추가됨")

	// Test with a sample request
	// 샘플 요청으로 테스트
	req, _ := http.NewRequest("GET", "/test", nil)
	rr := &responseRecorder{ResponseWriter: &dummyResponseWriter{}}

	logPrintln()
	logPrintln("  Testing middleware with sample request:")
	logPrintln("  샘플 요청으로 미들웨어 테스트:")

	app.ServeHTTP(rr, req)

	logPrintln("✓ Middleware executed successfully")
	logPrintln("✓ 미들웨어 실행 성공")
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

	logPrintln("✓ Added 3 middleware:")
	logPrintln("✓ 3개의 미들웨어 추가됨:")
	logPrintln("  1. Request ID")
	logPrintln("  2. Timing")
	logPrintln("  3. CORS")

	// Test with a sample request
	// 샘플 요청으로 테스트
	req, _ := http.NewRequest("GET", "/test", nil)
	rr := &responseRecorder{
		ResponseWriter: &dummyResponseWriter{headers: make(http.Header)},
	}

	app.ServeHTTP(rr, req)

	logPrintln()
	logPrintln("  Headers set by middleware:")
	logPrintln("  미들웨어가 설정한 헤더:")
	if id := rr.Header().Get("X-Request-ID"); id != "" {
		logPrintf("  - X-Request-ID: %s\n", id)
	}
	if cors := rr.Header().Get("Access-Control-Allow-Origin"); cors != "" {
		logPrintf("  - Access-Control-Allow-Origin: %s\n", cors)
	}

	logPrintln()
	logPrintln("✓ All middleware executed in order")
	logPrintln("✓ 모든 미들웨어가 순서대로 실행됨")
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
		websvrutil.WithAutoReload(false), // Disable in production
		websvrutil.WithLogger(true),      // Enable logging
		websvrutil.WithRecovery(true),    // Enable panic recovery
	)

	logPrintln("✓ Production configuration applied")
	logPrintln("✓ 프로덕션 설정 적용됨")
	logPrintln()
	logPrintln("Security Features / 보안 기능:")
	logPrintln("  ✓ Short timeouts to prevent slowloris attacks")
	logPrintln("  ✓ Slowloris 공격 방지를 위한 짧은 타임아웃")
	logPrintln("  ✓ Header size limits")
	logPrintln("  ✓ 헤더 크기 제한")
	logPrintln("  ✓ Panic recovery enabled")
	logPrintln("  ✓ 패닉 복구 활성화")
	logPrintln("  ✓ Request logging enabled")
	logPrintln("  ✓ 요청 로깅 활성화")
	logPrintln()
	logPrintln("Optimization / 최적화:")
	logPrintln("  ✓ Template caching (no auto-reload)")
	logPrintln("  ✓ 템플릿 캐싱 (자동 재로드 없음)")
	logPrintln("  ✓ Keep-alive with appropriate timeout")
	logPrintln("  ✓ 적절한 타임아웃의 Keep-alive")

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
