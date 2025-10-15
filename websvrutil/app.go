package websvrutil

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

// App is the main application instance for the web server.
// App는 웹 서버의 주요 애플리케이션 인스턴스입니다.
type App struct {
	// router is the HTTP request router (placeholder for now, will be implemented in v1.11.003)
	// router는 HTTP 요청 라우터입니다 (현재는 임시, v1.11.003에서 구현 예정)
	router http.Handler

	// middleware stores the middleware chain
	// middleware는 미들웨어 체인을 저장합니다
	middleware []MiddlewareFunc

	// templates is the template engine (placeholder for now, will be implemented in Phase 3)
	// templates는 템플릿 엔진입니다 (현재는 임시, Phase 3에서 구현 예정)
	templates interface{}

	// options holds the configuration options
	// options는 설정 옵션을 보유합니다
	options *Options

	// server is the underlying HTTP server
	// server는 기본 HTTP 서버입니다
	server *http.Server

	// mu protects concurrent access to the App
	// mu는 App에 대한 동시 액세스를 보호합니다
	mu sync.RWMutex

	// running indicates whether the server is currently running
	// running은 서버가 현재 실행 중인지 나타냅니다
	running bool
}

// MiddlewareFunc is a function that wraps an http.Handler.
// MiddlewareFunc는 http.Handler를 래핑하는 함수입니다.
type MiddlewareFunc func(http.Handler) http.Handler

// New creates a new App instance with the given options.
// New는 주어진 옵션으로 새 App 인스턴스를 생성합니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app := websvrutil.New(
//	    websvrutil.WithReadTimeout(30 * time.Second),
//	    websvrutil.WithLogger(true),
//	)
func New(opts ...Option) *App {
	// Apply default options / 기본 옵션 적용
	options := defaultOptions()

	// Apply user-provided options / 사용자 제공 옵션 적용
	for _, opt := range opts {
		opt(options)
	}

	// Create the app instance / 앱 인스턴스 생성
	app := &App{
		router:     nil, // Will be set in v1.11.003 / v1.11.003에서 설정 예정
		middleware: make([]MiddlewareFunc, 0),
		templates:  nil, // Will be set in Phase 3 / Phase 3에서 설정 예정
		options:    options,
		server:     nil, // Will be created in Run() / Run()에서 생성 예정
		running:    false,
	}

	return app
}

// Use adds middleware to the application's middleware chain.
// Use는 애플리케이션의 미들웨어 체인에 미들웨어를 추가합니다.
//
// Middleware functions are executed in the order they are added.
// 미들웨어 함수는 추가된 순서대로 실행됩니다.
//
// Example / 예제:
//
//	app.Use(loggingMiddleware)
//	app.Use(authMiddleware)
func (a *App) Use(middleware ...MiddlewareFunc) *App {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.running {
		// Cannot add middleware while server is running
		// 서버 실행 중에는 미들웨어를 추가할 수 없습니다
		panic("cannot add middleware while server is running")
	}

	a.middleware = append(a.middleware, middleware...)
	return a
}

// Run starts the HTTP server on the specified address.
// Run은 지정된 주소에서 HTTP 서버를 시작합니다.
//
// The address should be in the format "host:port" (e.g., "localhost:8080" or ":8080").
// 주소는 "host:port" 형식이어야 합니다 (예: "localhost:8080" 또는 ":8080").
//
// This method blocks until the server is shut down.
// 이 메서드는 서버가 종료될 때까지 차단됩니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	if err := app.Run(":8080"); err != nil {
//	    log.Fatal(err)
//	}
func (a *App) Run(addr string) error {
	a.mu.Lock()

	if a.running {
		a.mu.Unlock()
		return errors.New("server is already running")
	}

	// Build the handler chain with middleware
	// 미들웨어와 함께 핸들러 체인 구축
	handler := a.buildHandler()

	// Create the HTTP server
	// HTTP 서버 생성
	a.server = &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    a.options.ReadTimeout,
		WriteTimeout:   a.options.WriteTimeout,
		IdleTimeout:    a.options.IdleTimeout,
		MaxHeaderBytes: a.options.MaxHeaderBytes,
	}

	a.running = true
	a.mu.Unlock()

	// Start the server
	// 서버 시작
	fmt.Printf("Server starting on %s\n", addr)
	err := a.server.ListenAndServe()

	// Server stopped
	// 서버 중지됨
	a.mu.Lock()
	a.running = false
	a.mu.Unlock()

	// http.ErrServerClosed is expected when Shutdown is called
	// Shutdown이 호출되면 http.ErrServerClosed가 예상됩니다
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

// Shutdown gracefully shuts down the server without interrupting active connections.
// Shutdown은 활성 연결을 중단하지 않고 서버를 정상적으로 종료합니다.
//
// It waits for active connections to close until the context is cancelled.
// 컨텍스트가 취소될 때까지 활성 연결이 닫힐 때까지 기다립니다.
//
// Example / 예제:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	if err := app.Shutdown(ctx); err != nil {
//	    log.Printf("Server shutdown error: %v", err)
//	}
func (a *App) Shutdown(ctx context.Context) error {
	a.mu.RLock()
	server := a.server
	running := a.running
	a.mu.RUnlock()

	if !running {
		return errors.New("server is not running")
	}

	if server == nil {
		return errors.New("server is not initialized")
	}

	fmt.Println("Server shutting down gracefully...")
	return server.Shutdown(ctx)
}

// buildHandler builds the final HTTP handler by applying all middleware.
// buildHandler는 모든 미들웨어를 적용하여 최종 HTTP 핸들러를 구축합니다.
func (a *App) buildHandler() http.Handler {
	// Start with the router (placeholder for now)
	// 라우터로 시작 (현재는 임시)
	var handler http.Handler
	if a.router != nil {
		handler = a.router
	} else {
		// Default handler that returns 404 for all requests
		// 모든 요청에 대해 404를 반환하는 기본 핸들러
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		})
	}

	// Apply middleware in reverse order (last added = outermost)
	// 미들웨어를 역순으로 적용 (마지막 추가 = 가장 바깥쪽)
	for i := len(a.middleware) - 1; i >= 0; i-- {
		handler = a.middleware[i](handler)
	}

	return handler
}

// ServeHTTP implements the http.Handler interface.
// ServeHTTP는 http.Handler 인터페이스를 구현합니다.
//
// This allows the App to be used as an http.Handler.
// 이를 통해 App을 http.Handler로 사용할 수 있습니다.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mu.RLock()
	handler := a.buildHandler()
	a.mu.RUnlock()

	handler.ServeHTTP(w, r)
}
