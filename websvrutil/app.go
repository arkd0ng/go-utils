package websvrutil

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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

	// templates is the template engine
	// templates는 템플릿 엔진입니다
	templates *TemplateEngine

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
// Example
// 예제:
//
//	app := websvrutil.New()
//	app := websvrutil.New(
//	    websvrutil.WithReadTimeout(30 * time.Second),
//	    websvrutil.WithLogger(true),
//	)
func New(opts ...Option) *App {
	// Apply default options
	// 기본 옵션 적용
	options := defaultOptions()

	// Apply user-provided options
	// 사용자 제공 옵션 적용
	for _, opt := range opts {
		opt(options)
	}

	// Create the router
	// 라우터 생성
	router := newRouter()

	// Create template engine if template directory is set
	// 템플릿 디렉토리가 설정된 경우 템플릿 엔진 생성
	var templateEngine *TemplateEngine
	if options.TemplateDir != "" {
		templateEngine = NewTemplateEngine(options.TemplateDir)

		// Auto-load templates if enabled
		// 활성화된 경우 템플릿 자동 로드
		if err := templateEngine.LoadAll(); err != nil {
			// Log error but don't fail - templates might be loaded later
			// 에러 로그하지만 실패하지 않음 - 템플릿은 나중에 로드될 수 있음
			fmt.Printf("Warning: failed to auto-load templates: %v\n", err)
		}

		// Auto-load layouts if layout directory exists
		// 레이아웃 디렉토리가 존재하면 자동 로드
		if err := templateEngine.LoadAllLayouts(); err != nil {
			// Log error but don't fail - layouts might be loaded later
			// 에러 로그하지만 실패하지 않음 - 레이아웃은 나중에 로드될 수 있음
			fmt.Printf("Warning: failed to auto-load layouts: %v\n", err)
		}

		// Enable auto-reload if configured
		// 설정된 경우 자동 재로드 활성화
		if options.EnableAutoReload {
			if err := templateEngine.EnableAutoReload(); err != nil {
				fmt.Printf("Warning: failed to enable auto-reload: %v\n", err)
			} else {
				fmt.Println("[Template Hot Reload] Auto-reload enabled for templates and layouts")
			}
		}
	}

	// Create the app instance
	// 앱 인스턴스 생성
	app := &App{
		router:     router,
		middleware: make([]MiddlewareFunc, 0),
		templates:  templateEngine,
		options:    options,
		// Will be created in Run()
		// Run()에서 생성 예정
		server:  nil,
		running: false,
	}

	return app
}

// Use adds middleware to the application's middleware chain.
// Use는 애플리케이션의 미들웨어 체인에 미들웨어를 추가합니다.
//
// Middleware functions are executed in the order they are added.
// 미들웨어 함수는 추가된 순서대로 실행됩니다.
//
// Example
// 예제:
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

// registerRoute is a helper method that registers a route with the given HTTP method.
// registerRoute는 주어진 HTTP 메서드로 라우트를 등록하는 헬퍼 메서드입니다.
//
// This helper reduces code duplication across GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD methods.
// 이 헬퍼는 GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD 메서드 전반에 걸친 코드 중복을 줄입니다.
//
// Parameters
// 매개변수:
//   - method: HTTP method name (e.g., "GET", "POST", "PUT")
//   - pattern: URL pattern with parameters (e.g., "/users/:id")
//   - handler: HTTP handler function
//
// Thread safety
// 스레드 안전성:
// - Acquires mutex lock to prevent concurrent route registration
// - Mutex 락을 획득하여 동시 라우트 등록 방지
// - Panics if routes are added while server is running
// - 서버 실행 중 라우트가 추가되면 패닉 발생
//
// Returns
// 반환:
// - *App for method chaining
// - 메서드 체이닝을 위한 *App
func (a *App) registerRoute(method, pattern string, handler http.HandlerFunc) *App {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.running {
		panic("cannot add routes while server is running")
	}

	if router, ok := a.router.(*Router); ok {
		// Call the appropriate router method based on HTTP method
		// HTTP 메서드에 따라 적절한 라우터 메서드 호출
		switch method {
		case "GET":
			router.GET(pattern, handler)
		case "POST":
			router.POST(pattern, handler)
		case "PUT":
			router.PUT(pattern, handler)
		case "PATCH":
			router.PATCH(pattern, handler)
		case "DELETE":
			router.DELETE(pattern, handler)
		case "OPTIONS":
			router.OPTIONS(pattern, handler)
		case "HEAD":
			router.HEAD(pattern, handler)
		}
	}
	return a
}

// GET registers a GET route.
// GET은 GET 라우트를 등록합니다.
//
// Example
// 예제:
//
//	app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
//	    // Handler implementation
//	})
func (a *App) GET(pattern string, handler http.HandlerFunc) *App {
	return a.registerRoute("GET", pattern, handler)
}

// POST registers a POST route.
// POST는 POST 라우트를 등록합니다.
func (a *App) POST(pattern string, handler http.HandlerFunc) *App {
	return a.registerRoute("POST", pattern, handler)
}

// PUT registers a PUT route.
// PUT은 PUT 라우트를 등록합니다.
func (a *App) PUT(pattern string, handler http.HandlerFunc) *App {
	return a.registerRoute("PUT", pattern, handler)
}

// PATCH registers a PATCH route.
// PATCH는 PATCH 라우트를 등록합니다.
func (a *App) PATCH(pattern string, handler http.HandlerFunc) *App {
	return a.registerRoute("PATCH", pattern, handler)
}

// DELETE registers a DELETE route.
// DELETE는 DELETE 라우트를 등록합니다.
func (a *App) DELETE(pattern string, handler http.HandlerFunc) *App {
	return a.registerRoute("DELETE", pattern, handler)
}

// OPTIONS registers an OPTIONS route.
// OPTIONS는 OPTIONS 라우트를 등록합니다.
func (a *App) OPTIONS(pattern string, handler http.HandlerFunc) *App {
	return a.registerRoute("OPTIONS", pattern, handler)
}

// HEAD registers a HEAD route.
// HEAD는 HEAD 라우트를 등록합니다.
func (a *App) HEAD(pattern string, handler http.HandlerFunc) *App {
	return a.registerRoute("HEAD", pattern, handler)
}

// NotFound sets the handler for 404 Not Found responses.
// NotFound는 404 Not Found 응답에 대한 핸들러를 설정합니다.
//
// Example
// 예제:
//
//	app.NotFound(func(w http.ResponseWriter, r *http.Request) {
//	    w.WriteHeader(http.StatusNotFound)
//	    fmt.Fprintf(w, "Custom 404 page")
//	})
func (a *App) NotFound(handler http.HandlerFunc) *App {
	a.mu.Lock()
	defer a.mu.Unlock()

	if router, ok := a.router.(*Router); ok {
		router.NotFound(handler)
	}
	return a
}

// Static registers a route to serve static files from a directory.
// Static은 디렉토리에서 정적 파일을 제공하는 라우트를 등록합니다.
//
// The prefix is the URL path prefix (e.g., "/static"), and dir is the directory path.
// prefix는 URL 경로 접두사(예: "/static")이고, dir은 디렉토리 경로입니다.
//
// Example
// 예제:
//
//	app.Static("/static", "./public")
//
// // Serves files from ./public directory at /static/* URLs
// ./public 디렉토리의 파일을 /static/* URL에서 제공
func (a *App) Static(prefix, dir string) *App {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.running {
		panic("cannot add routes while server is running")
	}

	// Create file server handler
	// 파일 서버 핸들러 생성
	fileServer := http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))

	// Register wildcard route for all files under prefix
	// prefix 하위의 모든 파일에 대한 와일드카드 라우트 등록
	pattern := prefix + "/*"

	if router, ok := a.router.(*Router); ok {
		router.GET(pattern, func(w http.ResponseWriter, r *http.Request) {
			fileServer.ServeHTTP(w, r)
		})
	}

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
// Example
// 예제:
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
// Example
// 예제:
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

// RunWithGracefulShutdown starts the server and handles graceful shutdown on OS signals.
// RunWithGracefulShutdown은 서버를 시작하고 OS 신호에 대한 우아한 종료를 처리합니다.
//
// It listens for SIGINT (Ctrl+C) and SIGTERM signals and shuts down gracefully.
// SIGINT (Ctrl+C) 및 SIGTERM 신호를 수신하고 우아하게 종료합니다.
//
// The timeout parameter specifies how long to wait for active connections to close.
// timeout 매개변수는 활성 연결이 닫힐 때까지 기다릴 시간을 지정합니다.
//
// Example
// 예제:
//
//	app := websvrutil.New()
//	// ... configure routes ...
//	if err := app.RunWithGracefulShutdown(":8080", 30*time.Second); err != nil {
//	    log.Fatal(err)
//	}
func (a *App) RunWithGracefulShutdown(addr string, timeout time.Duration) error {
	// Start server in a goroutine
	// 고루틴에서 서버 시작
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- a.Run(addr)
	}()

	// Wait for interrupt signal
	// 인터럽트 신호 대기
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		// Server failed to start
		// 서버 시작 실패
		return err
	case sig := <-quit:
		// Received shutdown signal
		// 종료 신호 수신
		fmt.Printf("\nReceived signal: %v\n", sig)

		// Create shutdown context with timeout
		// 타임아웃이 있는 종료 컨텍스트 생성
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Attempt graceful shutdown
		// 우아한 종료 시도
		if err := a.Shutdown(ctx); err != nil {
			return fmt.Errorf("server shutdown failed: %w", err)
		}

		fmt.Println("Server stopped")
		return nil
	}
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

	// Store app in request context for template rendering
	// 템플릿 렌더링을 위해 요청 컨텍스트에 app 저장
	ctx := context.WithValue(r.Context(), "app", a)
	r = r.WithContext(ctx)

	handler.ServeHTTP(w, r)
}

// TemplateEngine returns the template engine instance.
// TemplateEngine은 템플릿 엔진 인스턴스를 반환합니다.
func (a *App) TemplateEngine() *TemplateEngine {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.templates
}

// LoadTemplate loads a single template file.
// LoadTemplate은 단일 템플릿 파일을 로드합니다.
//
// Example
// 예제:
//
//	err := app.LoadTemplate("index.html")
func (a *App) LoadTemplate(name string) error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates == nil {
		return fmt.Errorf("template engine not initialized (set TemplateDir option)")
	}

	return a.templates.Load(name)
}

// LoadTemplates loads all templates matching the pattern.
// LoadTemplates는 패턴과 일치하는 모든 템플릿을 로드합니다.
//
// Example
// 예제:
//
//	err := app.LoadTemplates("*.html")
func (a *App) LoadTemplates(pattern string) error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates == nil {
		return fmt.Errorf("template engine not initialized (set TemplateDir option)")
	}

	return a.templates.LoadGlob(pattern)
}

// ReloadTemplates reloads all templates from the template directory.
// ReloadTemplates는 템플릿 디렉토리에서 모든 템플릿을 다시 로드합니다.
//
// Example
// 예제:
//
//	err := app.ReloadTemplates()
func (a *App) ReloadTemplates() error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates == nil {
		return fmt.Errorf("template engine not initialized (set TemplateDir option)")
	}

	a.templates.Clear()
	return a.templates.LoadAll()
}

// AddTemplateFunc adds a custom template function.
// AddTemplateFunc는 커스텀 템플릿 함수를 추가합니다.
//
// Example
// 예제:
//
//	app.AddTemplateFunc("upper", strings.ToUpper)
func (a *App) AddTemplateFunc(name string, fn interface{}) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates != nil {
		a.templates.AddFunc(name, fn)
	}
}

// AddTemplateFuncs adds multiple custom template functions.
// AddTemplateFuncs는 여러 커스텀 템플릿 함수를 추가합니다.
//
// Example
// 예제:
//
//	app.AddTemplateFuncs(template.FuncMap{
//	    "upper": strings.ToUpper,
//	    "lower": strings.ToLower,
//	})
func (a *App) AddTemplateFuncs(funcs map[string]interface{}) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates != nil {
		a.templates.AddFuncs(funcs)
	}
}
