package websvrutil

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestNew tests the New function with default options.
// TestNew은 기본 옵션으로 New 함수를 테스트합니다.
func TestNew(t *testing.T) {
	app := New()

	if app == nil {
		t.Fatal("New() returned nil")
	}

	if app.options == nil {
		t.Fatal("options is nil")
	}

	if app.middleware == nil {
		t.Fatal("middleware is nil")
	}

	if app.running {
		t.Error("app should not be running initially")
	}
}

// TestNewWithOptions tests the New function with custom options.
// TestNewWithOptions는 커스텀 옵션으로 New 함수를 테스트합니다.
func TestNewWithOptions(t *testing.T) {
	customReadTimeout := 30 * time.Second
	customWriteTimeout := 45 * time.Second
	customTemplateDir := "custom/templates"

	app := New(
		WithReadTimeout(customReadTimeout),
		WithWriteTimeout(customWriteTimeout),
		WithTemplateDir(customTemplateDir),
		WithLogger(false),
	)

	if app.options.ReadTimeout != customReadTimeout {
		t.Errorf("ReadTimeout = %v, want %v", app.options.ReadTimeout, customReadTimeout)
	}

	if app.options.WriteTimeout != customWriteTimeout {
		t.Errorf("WriteTimeout = %v, want %v", app.options.WriteTimeout, customWriteTimeout)
	}

	if app.options.TemplateDir != customTemplateDir {
		t.Errorf("TemplateDir = %v, want %v", app.options.TemplateDir, customTemplateDir)
	}

	if app.options.EnableLogger {
		t.Error("EnableLogger should be false")
	}

	// Check that other options have default values
	// 다른 옵션이 기본값을 가지고 있는지 확인
	if app.options.IdleTimeout != 60*time.Second {
		t.Errorf("IdleTimeout = %v, want %v", app.options.IdleTimeout, 60*time.Second)
	}
}

// TestUse tests adding middleware to the app.
// TestUse는 앱에 미들웨어를 추가하는 것을 테스트합니다.
func TestUse(t *testing.T) {
	app := New()

	// Create test middleware
	// 테스트 미들웨어 생성
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Middleware-1", "true")
			next.ServeHTTP(w, r)
		})
	}

	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Middleware-2", "true")
			next.ServeHTTP(w, r)
		})
	}

	// Add middleware
	// 미들웨어 추가
	result := app.Use(middleware1)

	// Check that Use returns the app for chaining
	// Use가 체이닝을 위해 앱을 반환하는지 확인
	if result != app {
		t.Error("Use() should return the app for chaining")
	}

	if len(app.middleware) != 1 {
		t.Errorf("middleware count = %d, want 1", len(app.middleware))
	}

	// Add another middleware
	// 다른 미들웨어 추가
	app.Use(middleware2)

	if len(app.middleware) != 2 {
		t.Errorf("middleware count = %d, want 2", len(app.middleware))
	}
}

// TestUseMultiple tests adding multiple middleware at once.
// TestUseMultiple은 여러 미들웨어를 한 번에 추가하는 것을 테스트합니다.
func TestUseMultiple(t *testing.T) {
	app := New()

	middleware1 := func(next http.Handler) http.Handler {
		return next
	}

	middleware2 := func(next http.Handler) http.Handler {
		return next
	}

	middleware3 := func(next http.Handler) http.Handler {
		return next
	}

	app.Use(middleware1, middleware2, middleware3)

	if len(app.middleware) != 3 {
		t.Errorf("middleware count = %d, want 3", len(app.middleware))
	}
}

// TestServeHTTP tests the ServeHTTP method.
// TestServeHTTP는 ServeHTTP 메서드를 테스트합니다.
func TestServeHTTP(t *testing.T) {
	app := New()

	// Add test middleware
	// 테스트 미들웨어 추가
	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Test-Middleware", "executed")
			next.ServeHTTP(w, r)
		})
	})

	// Create test request
	// 테스트 요청 생성
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	// Call ServeHTTP
	// ServeHTTP 호출
	app.ServeHTTP(rec, req)

	// Check middleware was executed
	// 미들웨어가 실행되었는지 확인
	if rec.Header().Get("X-Test-Middleware") != "executed" {
		t.Error("middleware was not executed")
	}

	// Default handler should return 404
	// 기본 핸들러는 404를 반환해야 함
	if rec.Code != http.StatusNotFound {
		t.Errorf("status code = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

// TestMiddlewareOrder tests that middleware is executed in the correct order.
// TestMiddlewareOrder는 미들웨어가 올바른 순서로 실행되는지 테스트합니다.
func TestMiddlewareOrder(t *testing.T) {
	app := New()

	var executionOrder []string

	// Add middleware in order
	// 순서대로 미들웨어 추가
	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executionOrder = append(executionOrder, "first")
			next.ServeHTTP(w, r)
		})
	})

	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executionOrder = append(executionOrder, "second")
			next.ServeHTTP(w, r)
		})
	})

	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executionOrder = append(executionOrder, "third")
			next.ServeHTTP(w, r)
		})
	})

	// Create test request
	// 테스트 요청 생성
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	// Execute
	// 실행
	app.ServeHTTP(rec, req)

	// Check execution order
	// 실행 순서 확인
	expectedOrder := []string{"first", "second", "third"}
	if len(executionOrder) != len(expectedOrder) {
		t.Fatalf("execution order length = %d, want %d", len(executionOrder), len(expectedOrder))
	}

	for i, expected := range expectedOrder {
		if executionOrder[i] != expected {
			t.Errorf("execution order[%d] = %s, want %s", i, executionOrder[i], expected)
		}
	}
}

// TestShutdownWithoutRun tests shutting down a server that hasn't been started.
// TestShutdownWithoutRun은 시작되지 않은 서버를 종료하는 것을 테스트합니다.
func TestShutdownWithoutRun(t *testing.T) {
	app := New()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := app.Shutdown(ctx)
	if err == nil {
		t.Error("Shutdown() should return error when server is not running")
	}
}

// TestRunInvalidAddress tests running the server with an invalid address.
// TestRunInvalidAddress는 잘못된 주소로 서버를 실행하는 것을 테스트합니다.
func TestRunInvalidAddress(t *testing.T) {
	app := New()

	// Try to run on an invalid address
	// 잘못된 주소로 실행 시도
	go func() {
		err := app.Run("invalid:address:format")
		if err == nil {
			t.Error("Run() should return error for invalid address")
		}
	}()

	// Give it a moment to start
	// 시작할 시간 부여
	time.Sleep(100 * time.Millisecond)
}

// TestConcurrentUse tests that Use panics when called while server is running.
// TestConcurrentUse는 서버 실행 중 Use 호출 시 패닉이 발생하는지 테스트합니다.
func TestConcurrentUse(t *testing.T) {
	app := New()

	// Start the server in a goroutine
	// 고루틴에서 서버 시작
	go func() {
		_ = app.Run(":0") // Use :0 to let OS pick a free port
	}()

	// Wait for server to start
	// 서버 시작 대기
	time.Sleep(100 * time.Millisecond)

	// Try to add middleware while running
	// 실행 중 미들웨어 추가 시도
	defer func() {
		if r := recover(); r == nil {
			t.Error("Use() should panic when called while server is running")
		}

		// Shutdown the server
		// 서버 종료
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		_ = app.Shutdown(ctx)
	}()

	middleware := func(next http.Handler) http.Handler {
		return next
	}
	app.Use(middleware)
}

// TestDefaultOptions tests the default options values.
// TestDefaultOptions는 기본 옵션 값을 테스트합니다.
func TestDefaultOptions(t *testing.T) {
	opts := defaultOptions()

	tests := []struct {
		name     string
		got      interface{}
		want     interface{}
		testFunc func() bool
	}{
		{"ReadTimeout", opts.ReadTimeout, 15 * time.Second, func() bool { return opts.ReadTimeout == 15*time.Second }},
		{"WriteTimeout", opts.WriteTimeout, 15 * time.Second, func() bool { return opts.WriteTimeout == 15*time.Second }},
		{"IdleTimeout", opts.IdleTimeout, 60 * time.Second, func() bool { return opts.IdleTimeout == 60*time.Second }},
		{"MaxHeaderBytes", opts.MaxHeaderBytes, 1 << 20, func() bool { return opts.MaxHeaderBytes == 1<<20 }},
		{"TemplateDir", opts.TemplateDir, "templates", func() bool { return opts.TemplateDir == "templates" }},
		{"StaticDir", opts.StaticDir, "static", func() bool { return opts.StaticDir == "static" }},
		{"StaticPrefix", opts.StaticPrefix, "/static", func() bool { return opts.StaticPrefix == "/static" }},
		{"EnableAutoReload", opts.EnableAutoReload, false, func() bool { return opts.EnableAutoReload == false }},
		{"EnableLogger", opts.EnableLogger, true, func() bool { return opts.EnableLogger == true }},
		{"EnableRecovery", opts.EnableRecovery, true, func() bool { return opts.EnableRecovery == true }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.testFunc() {
				t.Errorf("%s = %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

// BenchmarkNew benchmarks the New function.
// BenchmarkNew은 New 함수를 벤치마크합니다.
func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New()
	}
}

// BenchmarkNewWithOptions benchmarks the New function with options.
// BenchmarkNewWithOptions는 옵션이 있는 New 함수를 벤치마크합니다.
func BenchmarkNewWithOptions(b *testing.B) {
	opts := []Option{
		WithReadTimeout(30 * time.Second),
		WithWriteTimeout(30 * time.Second),
		WithTemplateDir("custom"),
		WithLogger(false),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = New(opts...)
	}
}

// BenchmarkUse benchmarks adding middleware.
// BenchmarkUse는 미들웨어 추가를 벤치마크합니다.
func BenchmarkUse(b *testing.B) {
	middleware := func(next http.Handler) http.Handler {
		return next
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app := New()
		app.Use(middleware)
	}
}

// BenchmarkServeHTTP benchmarks the ServeHTTP method.
// BenchmarkServeHTTP는 ServeHTTP 메서드를 벤치마크합니다.
func BenchmarkServeHTTP(b *testing.B) {
	app := New()
	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	})

	req := httptest.NewRequest("GET", "/", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)
	}
}
