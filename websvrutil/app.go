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

// App orchestrates HTTP routing, middleware execution, template rendering, and graceful shutdown.
// App는 HTTP 라우팅, 미들웨어 실행, 템플릿 렌더링, 우아한 종료를 총괄하는 핵심 애플리케이션 컨테이너입니다.
type App struct {
	// router stores the active HTTP router used to dispatch incoming requests to handlers.
	// router는 들어오는 요청을 핸들러로 전달하는 활성 HTTP 라우터를 보관합니다.
	router http.Handler

	// middleware keeps user-registered middleware functions in the order they should run.
	// middleware는 사용자가 등록한 미들웨어 함수를 실행 순서대로 보관합니다.
	middleware []MiddlewareFunc

	// templates references the lazy-loaded HTML template engine (nil when templates are disabled).
	// templates는 지연 로드되는 HTML 템플릿 엔진을 가리키며(비활성화 시 nil), 템플릿 기능을 제공합니다.
	templates *TemplateEngine

	// options contains immutable configuration chosen during App initialization.
	// options는 App 초기화 시 결정되는 불변의 설정 값을 담고 있습니다.
	options *Options

	// server points to the concrete http.Server created during Run.
	// server는 Run 호출 시 생성되는 구체적인 http.Server 인스턴스를 가리킵니다.
	server *http.Server

	// mu guards mutable fields (middleware, router, running flag) against concurrent access.
	// mu는 변경 가능한 필드(middleware, router, running)를 동시 접근으로부터 보호합니다.
	mu sync.RWMutex

	// running indicates whether Run has started the HTTP server and it is currently serving traffic.
	// running은 Run 실행 후 HTTP 서버가 실제로 트래픽을 처리 중인지 여부를 나타냅니다.
	running bool
}

// MiddlewareFunc decorates an http.Handler with additional behavior (logging, auth, etc.) while preserving the signature.
// MiddlewareFunc는 로깅, 인증 같은 부가 동작을 추가하면서 http.Handler 시그니처를 유지하는 래퍼 함수입니다.
type MiddlewareFunc func(http.Handler) http.Handler

// New builds a new App, applies functional options, prepares the router, and optionally loads templates.
// New는 새로운 App을 생성하고 함수형 옵션을 적용하며 라우터를 준비하고 필요 시 템플릿을 로드합니다.
//
// Example:
//
//	app := websvrutil.New()
//	app := websvrutil.New(
//	    websvrutil.WithReadTimeout(30 * time.Second),
//	    websvrutil.WithLogger(true),
//	)
//
// 예제:
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

	// Create the router / 라우터 생성
	router := newRouter()

	// Create template engine if template directory is set / 템플릿 디렉토리가 설정된 경우 템플릿 엔진을 준비합니다.
	var templateEngine *TemplateEngine
	if options.TemplateDir != "" {
		templateEngine = NewTemplateEngine(options.TemplateDir)

		// Auto-load templates if enabled / 자동 로드가 활성화된 경우 템플릿을 즉시 읽어옵니다.
		if err := templateEngine.LoadAll(); err != nil {
			fmt.Printf("Warning: failed to auto-load templates: %v\n", err) // Log only; templates can be reloaded later / 로깅만 하고 추후 재로드 가능
		}

		// Auto-load layouts if layout directory exists / 레이아웃 디렉토리가 존재하면 자동 로드합니다.
		if err := templateEngine.LoadAllLayouts(); err != nil {
			fmt.Printf("Warning: failed to auto-load layouts: %v\n", err) // Inform only; layouts remain optional / 알림만 남기고 레이아웃은 선택적입니다.
		}

		// Enable auto-reload if configured / 자동 재로드가 설정된 경우 활성화합니다.
		if options.EnableAutoReload {
			if err := templateEngine.EnableAutoReload(); err != nil {
				fmt.Printf("Warning: failed to enable auto-reload: %v\n", err)
			} else {
				fmt.Println("[Template Hot Reload] Auto-reload enabled for templates and layouts")
			}
		}
	}

	// Create the app instance / App 인스턴스를 생성합니다.
	app := &App{
		router:     router,
		middleware: make([]MiddlewareFunc, 0),
		templates:  templateEngine,
		options:    options,
		// Will be created in Run() / Run() 실행 시 http.Server가 생성됩니다.
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

// Static registers a route to serve static files from a filesystem directory.
// This method provides a convenient way to serve CSS, JavaScript, images, fonts, and other
// static assets directly from the filesystem without writing custom handlers. It internally
// uses Go's http.FileServer with http.StripPrefix to handle file serving, directory browsing
// (if enabled), range requests, ETags, and proper MIME type detection.
//
// Static routes are ideal for serving assets in development and simple deployments. For production,
// consider using a CDN or dedicated static file server (nginx, CloudFront) for better performance,
// caching, and reduced application server load.
//
// Parameters:
//
//   - prefix: URL path prefix where static files will be served (must start with "/").
//     This prefix is stripped before looking up files in the directory.
//     Common patterns:
//
//   - "/static" - serves all files under /static/* URLs
//
//   - "/assets" - serves all files under /assets/* URLs
//
//   - "/public" - serves all files under /public/* URLs
//     The prefix should NOT include a trailing slash (use "/static", not "/static/").
//
//   - dir: Filesystem directory path containing static files to serve.
//     Can be absolute path ("/var/www/static") or relative to working directory ("./public").
//     The directory must exist and be readable, or file requests will return 404.
//     Relative paths are resolved from the current working directory (where app starts).
//     Consider using absolute paths or path.Join for portability.
//
// Returns:
//   - *App: Returns the App instance for method chaining, enabling fluent API:
//     app.Static("/static", "./public").Static("/assets", "./assets")
//
// Behavior:
//
//   - Registration Timing: Must be called before Run() or RunWithGracefulShutdown().
//     Attempting to register static routes after server start causes panic.
//
//   - Path Mapping: URL path "/static/css/style.css" maps to "{dir}/css/style.css".
//     The prefix is stripped, remaining path is used to locate file.
//
//   - Directory Listings: Go's http.FileServer shows directory listings if no index.html exists.
//     To disable directory browsing, ensure directories have index files or use custom handler.
//
//   - File Lookup: Files are served directly from filesystem without caching.
//     File changes are reflected immediately (useful in development).
//     For production, consider pre-compiling assets or using embedded filesystems.
//
//   - MIME Types: Content-Type header is automatically set based on file extension
//     using mime.TypeByExtension (e.g., .css → text/css, .js → application/javascript).
//
//   - Range Requests: http.FileServer supports HTTP range requests (byte-range downloads),
//     enabling resume/pause functionality and efficient video streaming.
//
//   - ETags: Automatically generated based on file modification time and size,
//     enabling browser cache validation and conditional requests (If-None-Match).
//
//   - Error Handling: Returns 404 Not Found if file doesn't exist,
//     403 Forbidden if file isn't readable, 500 Internal Server Error on filesystem errors.
//
// Thread-Safety:
//   - Safe for concurrent calls during app configuration phase (before server start).
//   - Acquires mutex lock to prevent concurrent modifications to routing table.
//   - Panics if called after server is running.
//
// Security Considerations:
//
//   - Path Traversal: Go's http.FileServer automatically sanitizes paths to prevent "../" attacks.
//     Requests for "/static/../../etc/passwd" are safely blocked.
//
//   - Hidden Files: Files starting with "." (e.g., .env, .git) are served by default.
//     Consider using custom middleware to block sensitive files if needed.
//
//   - Directory Traversal: If directory listings are disabled, attackers can't enumerate files.
//     Add index.html or use custom 403 handler for directories.
//
//   - Sensitive Files: Never serve directories containing source code, configuration,
//     or sensitive data. Use dedicated directories for public assets only.
//
// Performance Considerations:
//
//   - No Built-in Caching: Files are read from disk on each request. For high-traffic sites,
//     use CDN, nginx, or implement caching middleware.
//
//   - Development vs Production: Serving static files from application is convenient for
//     development but inefficient for production. Use reverse proxy or CDN in production.
//
//   - Compression: http.FileServer doesn't compress responses. Add gzip middleware
//     or use reverse proxy with compression for better performance.
//
//   - File System I/O: Each request performs filesystem I/O. Consider embedding
//     static files in binary using embed.FS for better performance.
//
// Common Use Cases:
//   - Development assets: app.Static("/static", "./static")
//   - Multiple asset directories: app.Static("/css", "./css").Static("/js", "./js")
//   - Public uploads: app.Static("/uploads", "./uploads")
//   - Documentation: app.Static("/docs", "./docs/html")
//   - Single-page app: app.Static("/", "./dist") (for React, Vue, Angular builds)
//
// Example - Basic Static File Serving:
//
//	// Serve files from ./public directory at /static/* URLs
//	app.Static("/static", "./public")
//	// URL: /static/css/style.css → File: ./public/css/style.css
//	// URL: /static/js/app.js → File: ./public/js/app.js
//	// URL: /static/images/logo.png → File: ./public/images/logo.png
//
// Example - Multiple Static Directories:
//
//	// Serve CSS, JavaScript, and images from separate directories
//	app.Static("/css", "./assets/css")
//	app.Static("/js", "./assets/js")
//	app.Static("/images", "./assets/images")
//
// Example - Absolute Path:
//
//	// Use absolute path for production deployment
//	app.Static("/static", "/var/www/myapp/static")
//
// Example - Embedded Filesystem (Go 1.16+):
//
//	//go:embed static/*
//	var staticFiles embed.FS
//
//	// Serve embedded files (better performance, single binary)
//	app.GET("/static/*", func(w http.ResponseWriter, r *http.Request) {
//	    fs := http.FileServer(http.FS(staticFiles))
//	    fs.ServeHTTP(w, r)
//	})
//
// Example - Custom Static Handler with Cache Headers:
//
//	app.GET("/static/*", func(w http.ResponseWriter, r *http.Request) {
//	    // Add aggressive caching for production
//	    w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
//
//	    // Serve file
//	    fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./public")))
//	    fs.ServeHTTP(w, r)
//	})
//
// Best Practices:
//   - Use dedicated directories for static assets (don't mix with source code).
//   - Add Cache-Control headers for better performance (via middleware).
//   - Consider using content-addressable URLs (e.g., /static/app.abc123.js).
//   - Block sensitive files (.git, .env, .config) using middleware.
//   - Use CDN or reverse proxy for production static file serving.
//   - Enable compression (gzip) for text assets (CSS, JS, HTML).
//   - Set appropriate CORS headers if assets are accessed from different domains.
//
// Alternative Approaches:
//   - Embedded Files: Use //go:embed for single-binary deployment with better performance.
//   - CDN: Upload assets to CDN (CloudFront, CloudFlare) for global distribution.
//   - Reverse Proxy: nginx/Apache serve static files, proxy API requests to Go app.
//   - Custom Handler: Write custom http.Handler for advanced caching, transformation.
//
// Static은 파일시스템 디렉토리에서 정적 파일을 제공하는 라우트를 등록합니다.
// 이 메서드는 CSS, JavaScript, 이미지, 폰트 및 기타 정적 자산을
// 커스텀 핸들러를 작성하지 않고 파일시스템에서 직접 제공하는 편리한 방법을 제공합니다.
// 내부적으로 Go의 http.FileServer와 http.StripPrefix를 사용하여
// 파일 제공, 디렉토리 탐색(활성화된 경우), 범위 요청, ETag 및 적절한 MIME 타입 감지를 처리합니다.
//
// Static 라우트는 개발 및 간단한 배포에서 자산을 제공하는 데 이상적입니다.
// 프로덕션의 경우 더 나은 성능, 캐싱 및 애플리케이션 서버 부하 감소를 위해
// CDN 또는 전용 정적 파일 서버(nginx, CloudFront) 사용을 고려하세요.
//
// 매개변수:
//
//   - prefix: 정적 파일이 제공될 URL 경로 접두사("/"로 시작해야 함).
//     이 접두사는 디렉토리에서 파일을 찾기 전에 제거됩니다.
//     일반적인 패턴:
//
//   - "/static" - /static/* URL 하위의 모든 파일 제공
//
//   - "/assets" - /assets/* URL 하위의 모든 파일 제공
//
//   - "/public" - /public/* URL 하위의 모든 파일 제공
//     접두사는 후행 슬래시를 포함하지 않아야 합니다("/static/" 대신 "/static" 사용).
//
//   - dir: 제공할 정적 파일이 포함된 파일시스템 디렉토리 경로.
//     절대 경로("/var/www/static") 또는 작업 디렉토리 기준 상대 경로("./public")일 수 있습니다.
//     디렉토리가 존재하고 읽을 수 있어야 하며, 그렇지 않으면 파일 요청이 404를 반환합니다.
//     상대 경로는 현재 작업 디렉토리(앱 시작 위치)에서 해석됩니다.
//     이식성을 위해 절대 경로 또는 path.Join 사용을 고려하세요.
//
// 반환값:
//   - *App: 메서드 체이닝을 위해 App 인스턴스를 반환하여 유창한 API를 가능하게 합니다:
//     app.Static("/static", "./public").Static("/assets", "./assets")
//
// 동작 방식:
//
//   - 등록 타이밍: Run() 또는 RunWithGracefulShutdown() 전에 호출되어야 합니다.
//     서버 시작 후 정적 라우트를 등록하려고 하면 패닉이 발생합니다.
//
//   - 경로 매핑: URL 경로 "/static/css/style.css"는 "{dir}/css/style.css"에 매핑됩니다.
//     접두사가 제거되고 나머지 경로가 파일을 찾는 데 사용됩니다.
//
//   - 디렉토리 목록: index.html이 없으면 Go의 http.FileServer가 디렉토리 목록을 표시합니다.
//     디렉토리 탐색을 비활성화하려면 디렉토리에 인덱스 파일이 있는지 확인하거나 커스텀 핸들러를 사용하세요.
//
//   - 파일 조회: 파일은 캐싱 없이 파일시스템에서 직접 제공됩니다.
//     파일 변경사항은 즉시 반영됩니다(개발에 유용).
//     프로덕션의 경우 자산을 미리 컴파일하거나 임베디드 파일시스템 사용을 고려하세요.
//
//   - MIME 타입: Content-Type 헤더는 파일 확장자를 기반으로
//     mime.TypeByExtension을 사용하여 자동으로 설정됩니다
//     (예: .css → text/css, .js → application/javascript).
//
//   - 범위 요청: http.FileServer는 HTTP 범위 요청(바이트 범위 다운로드)을 지원하여
//     재개/일시중지 기능 및 효율적인 비디오 스트리밍을 가능하게 합니다.
//
//   - ETag: 파일 수정 시간과 크기를 기반으로 자동으로 생성되어
//     브라우저 캐시 유효성 검사 및 조건부 요청(If-None-Match)을 가능하게 합니다.
//
//   - 오류 처리: 파일이 없으면 404 Not Found 반환,
//     파일을 읽을 수 없으면 403 Forbidden, 파일시스템 오류 시 500 Internal Server Error 반환.
//
// 스레드 안전성:
//   - 앱 구성 단계(서버 시작 전) 동안 동시 호출에 안전합니다.
//   - 라우팅 테이블에 대한 동시 수정을 방지하기 위해 뮤텍스 잠금을 획득합니다.
//   - 서버 실행 후 호출되면 패닉이 발생합니다.
//
// 보안 고려사항:
//
//   - 경로 순회: Go의 http.FileServer는 "../" 공격을 방지하기 위해 경로를 자동으로 정리합니다.
//     "/static/../../etc/passwd" 요청은 안전하게 차단됩니다.
//
//   - 숨김 파일: "."로 시작하는 파일(예: .env, .git)은 기본적으로 제공됩니다.
//     필요한 경우 민감한 파일을 차단하기 위해 커스텀 미들웨어 사용을 고려하세요.
//
//   - 디렉토리 순회: 디렉토리 목록이 비활성화된 경우 공격자가 파일을 열거할 수 없습니다.
//     디렉토리에 index.html을 추가하거나 커스텀 403 핸들러를 사용하세요.
//
//   - 민감한 파일: 소스 코드, 구성 또는 민감한 데이터가 포함된 디렉토리를 제공하지 마세요.
//     공개 자산만을 위한 전용 디렉토리를 사용하세요.
//
// 성능 고려사항:
//
//   - 내장 캐싱 없음: 파일은 각 요청마다 디스크에서 읽습니다. 트래픽이 많은 사이트의 경우
//     CDN, nginx 또는 캐싱 미들웨어를 사용하세요.
//
//   - 개발 vs 프로덕션: 애플리케이션에서 정적 파일을 제공하는 것은 개발에는 편리하지만
//     프로덕션에는 비효율적입니다. 프로덕션에서는 리버스 프록시 또는 CDN을 사용하세요.
//
//   - 압축: http.FileServer는 응답을 압축하지 않습니다. gzip 미들웨어를 추가하거나
//     압축 기능이 있는 리버스 프록시를 사용하여 더 나은 성능을 얻으세요.
//
//   - 파일 시스템 I/O: 각 요청은 파일시스템 I/O를 수행합니다. 더 나은 성능을 위해
//     embed.FS를 사용하여 바이너리에 정적 파일을 임베드하는 것을 고려하세요.
//
// 일반적인 사용 사례:
//   - 개발 자산: app.Static("/static", "./static")
//   - 여러 자산 디렉토리: app.Static("/css", "./css").Static("/js", "./js")
//   - 공개 업로드: app.Static("/uploads", "./uploads")
//   - 문서: app.Static("/docs", "./docs/html")
//   - 단일 페이지 앱: app.Static("/", "./dist") (React, Vue, Angular 빌드용)
//
// 예제 - 기본 정적 파일 제공:
//
//	// ./public 디렉토리의 파일을 /static/* URL에서 제공
//	app.Static("/static", "./public")
//	// URL: /static/css/style.css → 파일: ./public/css/style.css
//	// URL: /static/js/app.js → 파일: ./public/js/app.js
//	// URL: /static/images/logo.png → 파일: ./public/images/logo.png
//
// 예제 - 여러 정적 디렉토리:
//
//	// CSS, JavaScript 및 이미지를 별도의 디렉토리에서 제공
//	app.Static("/css", "./assets/css")
//	app.Static("/js", "./assets/js")
//	app.Static("/images", "./assets/images")
//
// 예제 - 절대 경로:
//
//	// 프로덕션 배포를 위한 절대 경로 사용
//	app.Static("/static", "/var/www/myapp/static")
//
// 예제 - 임베디드 파일시스템(Go 1.16+):
//
//	//go:embed static/*
//	var staticFiles embed.FS
//
//	// 임베디드 파일 제공(더 나은 성능, 단일 바이너리)
//	app.GET("/static/*", func(w http.ResponseWriter, r *http.Request) {
//	    fs := http.FileServer(http.FS(staticFiles))
//	    fs.ServeHTTP(w, r)
//	})
//
// 예제 - 캐시 헤더가 있는 커스텀 정적 핸들러:
//
//	app.GET("/static/*", func(w http.ResponseWriter, r *http.Request) {
//	    // 프로덕션을 위한 공격적인 캐싱 추가
//	    w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
//
//	    // 파일 제공
//	    fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./public")))
//	    fs.ServeHTTP(w, r)
//	})
//
// 모범 사례:
//   - 정적 자산을 위한 전용 디렉토리를 사용하세요(소스 코드와 섞지 마세요).
//   - 더 나은 성능을 위해 Cache-Control 헤더를 추가하세요(미들웨어를 통해).
//   - 콘텐츠 주소 지정 가능 URL 사용을 고려하세요(예: /static/app.abc123.js).
//   - 미들웨어를 사용하여 민감한 파일(.git, .env, .config)을 차단하세요.
//   - 프로덕션 정적 파일 제공에 CDN 또는 리버스 프록시를 사용하세요.
//   - 텍스트 자산(CSS, JS, HTML)에 대해 압축(gzip)을 활성화하세요.
//   - 자산이 다른 도메인에서 액세스되는 경우 적절한 CORS 헤더를 설정하세요.
//
// 대안 접근 방식:
//   - 임베디드 파일: 더 나은 성능으로 단일 바이너리 배포를 위해 //go:embed 사용.
//   - CDN: 전역 배포를 위해 자산을 CDN(CloudFront, CloudFlare)에 업로드.
//   - 리버스 프록시: nginx/Apache가 정적 파일을 제공하고 API 요청을 Go 앱에 프록시.
//   - 커스텀 핸들러: 고급 캐싱, 변환을 위한 커스텀 http.Handler 작성.
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

// Run starts the HTTP server and begins listening for incoming requests on the specified network address.
// This method is the primary way to start your web application in production and development.
// Run blocks the calling goroutine until the server is stopped via Shutdown() or encounters a fatal error.
// It configures the http.Server with all registered routes, middleware, templates, and options,
// then starts listening and serving HTTP traffic.
//
// Run is suitable for simple deployments where graceful shutdown isn't critical. For production
// applications that need to handle SIGINT/SIGTERM signals and drain connections gracefully,
// use RunWithGracefulShutdown() instead.
//
// Parameters:
//   - addr: Network address to listen on, in "host:port" format.
//     Common patterns:
//   - ":8080" - Listen on all interfaces (0.0.0.0), port 8080 (most common)
//   - "localhost:8080" - Listen only on localhost (127.0.0.1), port 8080 (development)
//   - "0.0.0.0:8080" - Explicitly listen on all interfaces
//   - "192.168.1.10:8080" - Listen on specific IP address
//   - ":80" or ":443" - Standard HTTP/HTTPS ports (requires elevated privileges)
//     Empty host ("") defaults to all interfaces (same as "0.0.0.0").
//     Port must be provided; addr like "localhost" without port will cause error.
//
// Returns:
//   - error: Returns error if server fails to start or encounters fatal error during operation.
//     Common errors:
//   - "address already in use" - Port is occupied by another process
//   - "permission denied" - Insufficient privileges for port (e.g., port 80 without sudo)
//   - "server is already running" - Run() called twice on same App instance
//   - Network-related errors during operation
//     Returns nil when server shuts down gracefully via Shutdown().
//
// Behavior:
//
//   - Blocking Call: Run() blocks until server stops. Execute in main goroutine or use
//     goroutine if you need concurrent operations:
//     go app.Run(":8080")  // Non-blocking in separate goroutine
//
//   - Initialization Sequence:
//     1. Acquires mutex lock and checks if server is already running
//     2. Builds final handler chain by applying all middleware to router
//     3. Creates http.Server with configured timeouts and limits
//     4. Sets running flag to true
//     5. Releases mutex and starts ListenAndServe()
//     6. Blocks until server stops
//     7. Sets running flag to false on shutdown
//
//   - Handler Chain: Middleware is applied in LIFO order (last added = outermost).
//     Example: app.Use(A).Use(B).Use(C) results in execution order: C → B → A → handler
//
//   - Server Configuration: Uses options provided to New():
//
//   - ReadTimeout: Maximum duration for reading entire request
//
//   - WriteTimeout: Maximum duration for writing response
//
//   - IdleTimeout: Keep-alive idle connection timeout
//
//   - MaxHeaderBytes: Maximum size of request headers
//
//   - Running State: Sets a.running = true during operation. This prevents:
//
//   - Adding new routes after server starts (would cause panic)
//
//   - Adding new middleware after server starts (would cause panic)
//
//   - Starting server twice on same App instance
//
//   - Graceful Shutdown: When Shutdown() is called from another goroutine:
//
//   - ListenAndServe() returns http.ErrServerClosed
//
//   - Run() returns nil (normal shutdown, not an error)
//
//   - Active connections drain according to shutdown context timeout
//
// Thread-Safety:
//   - Safe to call from single goroutine (typical usage in main()).
//   - Uses mutex to protect running state check and modification.
//   - Cannot call Run() concurrently on same App instance (returns error).
//   - Can call Shutdown() from different goroutine while Run() is blocking.
//
// Port Requirements:
//
//   - Privileged Ports (1-1023): Require root/administrator privileges on Unix systems.
//     Ports 80 (HTTP) and 443 (HTTPS) need sudo or setcap capabilities.
//
//   - Ephemeral Ports (1024-65535): Can be used by any user.
//     Development typically uses 3000, 8000, 8080, 8888, etc.
//
//   - Port Conflicts: If port is already in use, Run() returns error immediately.
//     Use different port or stop conflicting process.
//
// Performance Considerations:
//
//   - Timeout Configuration: Set appropriate ReadTimeout and WriteTimeout to prevent
//     slowloris attacks and resource exhaustion from slow clients.
//
//   - MaxHeaderBytes: Limit header size to prevent large header attacks.
//     Default is 1 MB, adjust based on your application needs.
//
//   - Keep-Alive: IdleTimeout controls how long idle connections are kept open.
//     Shorter timeouts free resources faster but require more connection handshakes.
//
//   - Middleware Overhead: Each middleware adds processing time per request.
//     Minimize middleware and optimize hot paths.
//
// Common Use Cases:
//   - Simple deployment: app.Run(":8080")
//   - Development server: app.Run("localhost:3000")
//   - Production with error handling: if err := app.Run(":80"); err != nil { log.Fatal(err) }
//   - Docker deployment: app.Run(":8080") (expose port 8080)
//   - Multiple servers: Run different App instances on different ports
//
// Example - Basic Server Start:
//
//	app := websvrutil.New()
//	app.GET("/", homeHandler)
//	app.GET("/api/users", usersHandler)
//
//	fmt.Println("Starting server on :8080")
//	if err := app.Run(":8080"); err != nil {
//	    log.Fatalf("Server failed: %v", err)
//	}
//
// Example - Development Server:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./templates"))
//	app.Static("/static", "./static")
//	app.GET("/", indexHandler)
//
//	log.Println("Dev server running on http://localhost:3000")
//	if err := app.Run("localhost:3000"); err != nil {
//	    log.Fatal(err)
//	}
//
// Example - Production with Environment Variable:
//
//	port := os.Getenv("PORT")
//	if port == "" {
//	    port = "8080"  // Default port
//	}
//
//	app := websvrutil.New(
//	    websvrutil.WithReadTimeout(10 * time.Second),
//	    websvrutil.WithWriteTimeout(10 * time.Second),
//	)
//	// ... configure routes ...
//
//	log.Printf("Server starting on port %s", port)
//	if err := app.Run(":" + port); err != nil {
//	    log.Fatalf("Server error: %v", err)
//	}
//
// Example - Non-Blocking with Goroutine:
//
//	app := websvrutil.New()
//	// ... configure routes ...
//
//	// Start server in background goroutine
//	go func() {
//	    if err := app.Run(":8080"); err != nil {
//	        log.Printf("Server stopped: %v", err)
//	    }
//	}()
//
//	// Main goroutine can do other work
//	time.Sleep(1 * time.Second)
//	fmt.Println("Server is running in background")
//
//	// Keep main goroutine alive
//	select {}
//
// Example - Multiple Servers:
//
//	// API server
//	apiApp := websvrutil.New()
//	apiApp.GET("/api/health", healthHandler)
//	go apiApp.Run(":8080")
//
//	// Admin server
//	adminApp := websvrutil.New()
//	adminApp.GET("/admin/dashboard", dashboardHandler)
//	go adminApp.Run(":9090")
//
//	// Block main goroutine
//	select {}
//
// Error Handling:
//   - Always check returned error to detect startup failures.
//   - Log errors appropriately for debugging and monitoring.
//   - Use log.Fatal() or os.Exit(1) for fatal startup errors.
//   - Distinguish between startup errors (immediate) and runtime errors (during operation).
//
// Alternative Startup Methods:
//   - RunWithGracefulShutdown(): Handles OS signals and drains connections gracefully.
//   - http.Server.ListenAndServe(): Direct http.Server usage if you need custom setup.
//   - http.Server.ListenAndServeTLS(): For HTTPS with TLS certificates.
//
// Best Practices:
//   - Use RunWithGracefulShutdown() for production (handles SIGINT/SIGTERM).
//   - Configure appropriate timeouts to prevent resource exhaustion.
//   - Use environment variables for port configuration (12-factor app).
//   - Log server start/stop events for monitoring.
//   - Use health check endpoints for load balancer monitoring.
//   - Test port availability before deployment (netstat, lsof).
//   - Document required ports in deployment guides.
//
// Deployment Considerations:
//   - Docker: Use EXPOSE directive and map ports with -p flag.
//   - Kubernetes: Define containerPort in Pod spec and Service port.
//   - Cloud: Configure security groups/firewall rules for ports.
//   - Reverse Proxy: Run on high port (8080), proxy from nginx/Apache on port 80.
//   - Load Balancer: Run multiple instances on different servers/ports.
//
// Run은 HTTP 서버를 시작하고 지정된 네트워크 주소에서 들어오는 요청을 수신합니다.
// 이 메서드는 프로덕션 및 개발 환경에서 웹 애플리케이션을 시작하는 주요 방법입니다.
// Run은 Shutdown()을 통해 서버가 중지되거나 치명적인 오류가 발생할 때까지 호출하는 고루틴을 차단합니다.
// 등록된 모든 라우트, 미들웨어, 템플릿 및 옵션으로 http.Server를 구성한 다음
// HTTP 트래픽을 수신하고 제공하기 시작합니다.
//
// Run은 우아한 종료가 중요하지 않은 간단한 배포에 적합합니다. SIGINT/SIGTERM 신호를 처리하고
// 연결을 우아하게 드레인해야 하는 프로덕션 애플리케이션의 경우 대신 RunWithGracefulShutdown()을 사용하세요.
//
// 매개변수:
//   - addr: "host:port" 형식으로 수신할 네트워크 주소.
//     일반적인 패턴:
//   - ":8080" - 모든 인터페이스(0.0.0.0), 포트 8080에서 수신(가장 일반적)
//   - "localhost:8080" - localhost(127.0.0.1)에서만 수신, 포트 8080(개발)
//   - "0.0.0.0:8080" - 모든 인터페이스에서 명시적으로 수신
//   - "192.168.1.10:8080" - 특정 IP 주소에서 수신
//   - ":80" 또는 ":443" - 표준 HTTP/HTTPS 포트(높은 권한 필요)
//     빈 호스트("")는 모든 인터페이스로 기본 설정("0.0.0.0"과 동일).
//     포트를 제공해야 합니다. 포트 없이 "localhost"와 같은 addr은 오류를 발생시킵니다.
//
// 반환값:
//   - error: 서버가 시작에 실패하거나 작동 중 치명적인 오류가 발생하면 오류를 반환합니다.
//     일반적인 오류:
//   - "address already in use" - 다른 프로세스가 포트를 점유 중
//   - "permission denied" - 포트에 대한 권한 부족(예: sudo 없이 포트 80)
//   - "server is already running" - 동일한 App 인스턴스에서 Run()을 두 번 호출
//   - 작동 중 네트워크 관련 오류
//     Shutdown()을 통해 서버가 우아하게 종료되면 nil을 반환합니다.
//
// 동작 방식:
//
//   - 차단 호출: Run()은 서버가 중지될 때까지 차단됩니다. main 고루틴에서 실행하거나
//     동시 작업이 필요한 경우 고루틴을 사용하세요:
//     go app.Run(":8080")  // 별도의 고루틴에서 비차단
//
//   - 초기화 순서:
//     1. 뮤텍스 잠금을 획득하고 서버가 이미 실행 중인지 확인
//     2. 라우터에 모든 미들웨어를 적용하여 최종 핸들러 체인 구축
//     3. 구성된 타임아웃 및 제한으로 http.Server 생성
//     4. running 플래그를 true로 설정
//     5. 뮤텍스를 해제하고 ListenAndServe() 시작
//     6. 서버가 중지될 때까지 차단
//     7. 종료 시 running 플래그를 false로 설정
//
//   - 핸들러 체인: 미들웨어는 LIFO 순서로 적용됩니다(마지막 추가 = 가장 바깥쪽).
//     예: app.Use(A).Use(B).Use(C)는 실행 순서 C → B → A → handler가 됩니다.
//
//   - 서버 구성: New()에 제공된 옵션 사용:
//
//   - ReadTimeout: 전체 요청을 읽는 최대 시간
//
//   - WriteTimeout: 응답을 작성하는 최대 시간
//
//   - IdleTimeout: Keep-alive 유휴 연결 타임아웃
//
//   - MaxHeaderBytes: 요청 헤더의 최대 크기
//
//   - 실행 상태: 작동 중 a.running = true를 설정합니다. 이는 다음을 방지합니다:
//
//   - 서버 시작 후 새 라우트 추가(패닉 발생)
//
//   - 서버 시작 후 새 미들웨어 추가(패닉 발생)
//
//   - 동일한 App 인스턴스에서 서버를 두 번 시작
//
//   - 우아한 종료: 다른 고루틴에서 Shutdown()이 호출되면:
//
//   - ListenAndServe()가 http.ErrServerClosed를 반환
//
//   - Run()은 nil을 반환(정상 종료, 오류 아님)
//
//   - 활성 연결은 종료 컨텍스트 타임아웃에 따라 드레인됨
//
// 스레드 안전성:
//   - 단일 고루틴에서 호출하기에 안전합니다(main()의 일반적인 사용).
//   - running 상태 검사 및 수정을 보호하기 위해 뮤텍스를 사용합니다.
//   - 동일한 App 인스턴스에서 Run()을 동시에 호출할 수 없습니다(오류 반환).
//   - Run()이 차단되는 동안 다른 고루틴에서 Shutdown()을 호출할 수 있습니다.
//
// 포트 요구사항:
//
//   - 특권 포트(1-1023): Unix 시스템에서 root/관리자 권한이 필요합니다.
//     포트 80(HTTP)과 443(HTTPS)은 sudo 또는 setcap 기능이 필요합니다.
//
//   - 임시 포트(1024-65535): 모든 사용자가 사용할 수 있습니다.
//     개발은 일반적으로 3000, 8000, 8080, 8888 등을 사용합니다.
//
//   - 포트 충돌: 포트가 이미 사용 중이면 Run()이 즉시 오류를 반환합니다.
//     다른 포트를 사용하거나 충돌하는 프로세스를 중지하세요.
//
// 성능 고려사항:
//
//   - 타임아웃 구성: 느린 클라이언트로 인한 slowloris 공격 및 리소스 고갈을 방지하기 위해
//     적절한 ReadTimeout 및 WriteTimeout을 설정하세요.
//
//   - MaxHeaderBytes: 대용량 헤더 공격을 방지하기 위해 헤더 크기를 제한하세요.
//     기본값은 1MB이며 애플리케이션 요구사항에 따라 조정하세요.
//
//   - Keep-Alive: IdleTimeout은 유휴 연결을 얼마나 오래 열어 둘지 제어합니다.
//     짧은 타임아웃은 리소스를 더 빨리 해제하지만 더 많은 연결 핸드셰이크가 필요합니다.
//
//   - 미들웨어 오버헤드: 각 미들웨어는 요청당 처리 시간을 추가합니다.
//     미들웨어를 최소화하고 핫 패스를 최적화하세요.
//
// 일반적인 사용 사례:
//   - 간단한 배포: app.Run(":8080")
//   - 개발 서버: app.Run("localhost:3000")
//   - 오류 처리가 있는 프로덕션: if err := app.Run(":80"); err != nil { log.Fatal(err) }
//   - Docker 배포: app.Run(":8080") (포트 8080 노출)
//   - 여러 서버: 다른 포트에서 다른 App 인스턴스 실행
//
// 예제 - 기본 서버 시작:
//
//	app := websvrutil.New()
//	app.GET("/", homeHandler)
//	app.GET("/api/users", usersHandler)
//
//	fmt.Println("Starting server on :8080")
//	if err := app.Run(":8080"); err != nil {
//	    log.Fatalf("Server failed: %v", err)
//	}
//
// 예제 - 개발 서버:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./templates"))
//	app.Static("/static", "./static")
//	app.GET("/", indexHandler)
//
//	log.Println("Dev server running on http://localhost:3000")
//	if err := app.Run("localhost:3000"); err != nil {
//	    log.Fatal(err)
//	}
//
// 예제 - 환경 변수가 있는 프로덕션:
//
//	port := os.Getenv("PORT")
//	if port == "" {
//	    port = "8080"  // 기본 포트
//	}
//
//	app := websvrutil.New(
//	    websvrutil.WithReadTimeout(10 * time.Second),
//	    websvrutil.WithWriteTimeout(10 * time.Second),
//	)
//	// ... 라우트 구성 ...
//
//	log.Printf("Server starting on port %s", port)
//	if err := app.Run(":" + port); err != nil {
//	    log.Fatalf("Server error: %v", err)
//	}
//
// 예제 - 고루틴으로 비차단:
//
//	app := websvrutil.New()
//	// ... 라우트 구성 ...
//
//	// 백그라운드 고루틴에서 서버 시작
//	go func() {
//	    if err := app.Run(":8080"); err != nil {
//	        log.Printf("Server stopped: %v", err)
//	    }
//	}()
//
//	// main 고루틴은 다른 작업을 수행할 수 있음
//	time.Sleep(1 * time.Second)
//	fmt.Println("Server is running in background")
//
//	// main 고루틴을 살아있게 유지
//	select {}
//
// 예제 - 여러 서버:
//
//	// API 서버
//	apiApp := websvrutil.New()
//	apiApp.GET("/api/health", healthHandler)
//	go apiApp.Run(":8080")
//
//	// 관리 서버
//	adminApp := websvrutil.New()
//	adminApp.GET("/admin/dashboard", dashboardHandler)
//	go adminApp.Run(":9090")
//
//	// main 고루틴 차단
//	select {}
//
// 오류 처리:
//   - 시작 실패를 감지하기 위해 항상 반환된 오류를 확인하세요.
//   - 디버깅 및 모니터링을 위해 오류를 적절히 로깅하세요.
//   - 치명적인 시작 오류에 대해 log.Fatal() 또는 os.Exit(1)을 사용하세요.
//   - 시작 오류(즉시)와 런타임 오류(작동 중)를 구별하세요.
//
// 대체 시작 메서드:
//   - RunWithGracefulShutdown(): OS 신호를 처리하고 연결을 우아하게 드레인합니다.
//   - http.Server.ListenAndServe(): 커스텀 설정이 필요한 경우 직접 http.Server 사용.
//   - http.Server.ListenAndServeTLS(): TLS 인증서로 HTTPS 사용.
//
// 모범 사례:
//   - 프로덕션에는 RunWithGracefulShutdown() 사용(SIGINT/SIGTERM 처리).
//   - 리소스 고갈을 방지하기 위해 적절한 타임아웃 구성.
//   - 포트 구성에 환경 변수 사용(12-factor 앱).
//   - 모니터링을 위해 서버 시작/중지 이벤트 로깅.
//   - 로드 밸런서 모니터링을 위해 헬스 체크 엔드포인트 사용.
//   - 배포 전 포트 가용성 테스트(netstat, lsof).
//   - 배포 가이드에 필요한 포트 문서화.
//
// 배포 고려사항:
//   - Docker: EXPOSE 지시문 사용 및 -p 플래그로 포트 매핑.
//   - Kubernetes: Pod spec에서 containerPort 정의 및 Service 포트.
//   - 클라우드: 포트에 대한 보안 그룹/방화벽 규칙 구성.
//   - 리버스 프록시: 높은 포트(8080)에서 실행, 포트 80의 nginx/Apache에서 프록시.
//   - 로드 밸런서: 다른 서버/포트에서 여러 인스턴스 실행.
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

// Shutdown initiates graceful server shutdown, waiting for active connections to complete or timeout.
// This method is the recommended way to stop a running HTTP server without abruptly dropping active
// requests. It stops accepting new connections immediately while allowing in-flight requests to
// complete within the context timeout period. This ensures data integrity, prevents client errors,
// and maintains service reliability during deployments, scaling, or maintenance.
//
// Graceful shutdown is essential for production systems to avoid:
// - Dropped client requests (broken downloads, incomplete API calls)
// - Corrupted database transactions (half-written data)
// - Lost websocket messages or streaming data
// - Error monitoring alerts from terminated connections
// - Poor user experience during deployments
//
// Parameters:
//   - ctx: Context controlling shutdown timeout and cancellation.
//     Typically created with context.WithTimeout() or context.WithDeadline().
//     Common timeout values:
//   - 5 seconds: Fast shutdown, risk of terminating slow requests
//   - 30 seconds: Balanced approach for most applications (recommended)
//   - 60 seconds: Patient shutdown for long-running requests
//   - No timeout (context.Background()): Wait indefinitely (not recommended)
//     If context expires before all connections close, Shutdown returns context.DeadlineExceeded.
//     Connections remaining after timeout are forcibly closed.
//
// Returns:
//   - error: Returns error if shutdown fails or server isn't running.
//     Common errors:
//   - "server is not running" - Shutdown() called when server isn't started
//   - "server is not initialized" - Internal state inconsistency
//   - context.DeadlineExceeded - Timeout reached before all connections closed
//   - Network errors during connection cleanup
//     Returns nil on successful graceful shutdown (all connections drained within timeout).
//
// Behavior:
//
//   - Immediate Effect: Stops accepting new connections instantly.
//     New client connection attempts receive "connection refused" errors.
//
//   - In-Flight Requests: Allows active HTTP requests to complete normally.
//     Handlers continue executing, writing responses, closing connections.
//
//   - Idle Connections: Keep-alive idle connections are closed immediately.
//     No disruption to clients since no active request was in progress.
//
//   - WebSocket/Streaming: Long-lived connections get full timeout period to close.
//     Application should monitor context and cleanly close streams.
//
//   - Timeout Behavior: If context deadline expires:
//
//   - Remaining connections are forcibly terminated
//
//   - Shutdown returns context.DeadlineExceeded error
//
//   - Clients may see "connection reset" or incomplete responses
//
//   - Run() Completion: After Shutdown() completes, Run() unblocks and returns nil.
//     This allows main goroutine or deployment scripts to proceed.
//
//   - State Cleanup: Running flag is set to false, allowing app reuse (though not recommended).
//
// Concurrency:
//
//   - Thread-Safe: Safe to call from different goroutine while Run() is blocking.
//     Common pattern: Run() in main goroutine, Shutdown() in signal handler goroutine.
//
//   - Single Shutdown: Calling Shutdown() multiple times is safe but redundant.
//     Subsequent calls return immediately (idempotent operation).
//
//   - Locks: Uses read lock to check state, doesn't block Run() operation.
//     Delegates to http.Server.Shutdown() which handles synchronization internally.
//
// Usage Patterns:
//   - Manual Shutdown: Application-triggered shutdown based on internal logic
//   - Signal Handler: OS signal (SIGINT/SIGTERM) triggers shutdown
//   - Health Check: Gracefully remove server from load balancer before shutdown
//   - Rolling Deployment: Drain connections before replacing server instance
//   - Testing: Clean teardown of test servers
//
// Best Practices:
//   - Always use timeout context to prevent indefinite waits.
//   - Choose timeout based on longest expected request duration.
//   - Log shutdown start/completion for monitoring and debugging.
//   - Notify load balancer to stop sending traffic before shutdown.
//   - Close database connections and external resources after shutdown.
//   - In handlers, check context cancellation for early exit during shutdown.
//   - Test shutdown behavior under load to validate timeout settings.
//
// Common Use Cases:
//   - Signal-based shutdown: syscall.SIGINT/SIGTERM handler calls Shutdown()
//   - Kubernetes pod termination: PreStop hook calls shutdown endpoint
//   - Health check degradation: Mark unhealthy, wait, then shutdown
//   - Deployment automation: Blue-green or rolling update workflows
//   - Development: Ctrl+C triggers graceful shutdown
//
// Example - Basic Shutdown with Timeout:
//
//	// In separate goroutine or signal handler
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	if err := app.Shutdown(ctx); err != nil {
//	    if errors.Is(err, context.DeadlineExceeded) {
//	        log.Println("Shutdown timeout: some connections were forcibly closed")
//	    } else {
//	        log.Printf("Shutdown error: %v", err)
//	    }
//	} else {
//	    log.Println("Shutdown completed successfully")
//	}
//
// Example - Signal Handler Pattern:
//
//	go func() {
//	    if err := app.Run(":8080"); err != nil {
//	        log.Fatalf("Server error: %v", err)
//	    }
//	}()
//
//	// Wait for interrupt signal
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	<-quit
//
//	log.Println("Shutting down server...")
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	if err := app.Shutdown(ctx); err != nil {
//	    log.Fatalf("Shutdown failed: %v", err)
//	}
//	log.Println("Server exited")
//
// Example - Health Check Integration:
//
//	var healthy atomic.Bool
//	healthy.Store(true)
//
//	app.GET("/health", func(w http.ResponseWriter, r *http.Request) {
//	    if healthy.Load() {
//	        w.WriteHeader(http.StatusOK)
//	    } else {
//	        w.WriteHeader(http.StatusServiceUnavailable)
//	    }
//	})
//
//	// On shutdown signal
//	healthy.Store(false)           // Mark unhealthy
//	time.Sleep(10 * time.Second)   // Wait for load balancer to notice
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	app.Shutdown(ctx)              // Drain connections
//
// Example - Handler Context Check:
//
//	app.GET("/long-task", func(w http.ResponseWriter, r *http.Request) {
//	    for i := 0; i < 100; i++ {
//	        // Check if shutdown is in progress
//	        select {
//	        case <-r.Context().Done():
//	            log.Println("Handler cancelled during shutdown")
//	            http.Error(w, "Service shutting down", http.StatusServiceUnavailable)
//	            return
//	        default:
//	        }
//
//	        // Do work
//	        time.Sleep(100 * time.Millisecond)
//	    }
//
//	    w.Write([]byte("Completed"))
//	})
//
// Kubernetes Example:
//
//	# In Kubernetes Pod spec
//	lifecycle:
//	  preStop:
//	    exec:
//	      command: ["/bin/sh", "-c", "sleep 15 && kill -SIGTERM 1"]
//
//	# Application handles SIGTERM
//	signal.Notify(quit, syscall.SIGTERM)
//	<-quit
//	app.Shutdown(context.WithTimeout(context.Background(), 30*time.Second))
//
// Monitoring and Observability:
//   - Log shutdown initiation with timestamp
//   - Track shutdown duration metrics
//   - Alert on shutdown timeouts
//   - Monitor forcibly closed connection count
//   - Include shutdown reason in logs (signal, health check, manual)
//
// Shutdown은 우아한 서버 종료를 시작하여 활성 연결이 완료되거나 타임아웃될 때까지 기다립니다.
// 이 메서드는 활성 요청을 갑자기 중단하지 않고 실행 중인 HTTP 서버를 중지하는 권장 방법입니다.
// 새 연결 수락을 즉시 중지하면서 진행 중인 요청이 컨텍스트 타임아웃 기간 내에 완료되도록 합니다.
// 이는 데이터 무결성을 보장하고, 클라이언트 오류를 방지하며,
// 배포, 확장 또는 유지 관리 중 서비스 안정성을 유지합니다.
//
// 우아한 종료는 다음을 방지하기 위해 프로덕션 시스템에 필수적입니다:
// - 중단된 클라이언트 요청(깨진 다운로드, 불완전한 API 호출)
// - 손상된 데이터베이스 트랜잭션(절반만 작성된 데이터)
// - 손실된 웹소켓 메시지 또는 스트리밍 데이터
// - 종료된 연결로 인한 오류 모니터링 경고
// - 배포 중 나쁜 사용자 경험
//
// 매개변수:
//   - ctx: 종료 타임아웃 및 취소를 제어하는 컨텍스트.
//     일반적으로 context.WithTimeout() 또는 context.WithDeadline()로 생성됩니다.
//     일반적인 타임아웃 값:
//   - 5초: 빠른 종료, 느린 요청 종료 위험
//   - 30초: 대부분의 애플리케이션에 균형 잡힌 접근 방식(권장)
//   - 60초: 장기 실행 요청에 대한 참을성 있는 종료
//   - 타임아웃 없음(context.Background()): 무기한 대기(권장하지 않음)
//     모든 연결이 닫히기 전에 컨텍스트가 만료되면 Shutdown은 context.DeadlineExceeded를 반환합니다.
//     타임아웃 후 남은 연결은 강제로 닫힙니다.
//
// 반환값:
//   - error: 종료가 실패하거나 서버가 실행 중이지 않으면 오류를 반환합니다.
//     일반적인 오류:
//   - "server is not running" - 서버가 시작되지 않았을 때 Shutdown() 호출
//   - "server is not initialized" - 내부 상태 불일치
//   - context.DeadlineExceeded - 모든 연결이 닫히기 전에 타임아웃 도달
//   - 연결 정리 중 네트워크 오류
//     성공적인 우아한 종료(타임아웃 내 모든 연결 드레인) 시 nil 반환.
//
// 동작 방식:
//
//   - 즉각적인 효과: 새 연결 수락을 즉시 중지합니다.
//     새 클라이언트 연결 시도는 "connection refused" 오류를 받습니다.
//
//   - 진행 중인 요청: 활성 HTTP 요청이 정상적으로 완료되도록 합니다.
//     핸들러는 계속 실행되고, 응답을 작성하고, 연결을 닫습니다.
//
//   - 유휴 연결: Keep-alive 유휴 연결은 즉시 닫힙니다.
//     진행 중인 활성 요청이 없었으므로 클라이언트에 중단이 없습니다.
//
//   - WebSocket/스트리밍: 장기 연결은 닫기 위한 전체 타임아웃 기간을 받습니다.
//     애플리케이션은 컨텍스트를 모니터링하고 스트림을 깔끔하게 닫아야 합니다.
//
//   - 타임아웃 동작: 컨텍스트 데드라인이 만료되면:
//
//   - 남은 연결이 강제로 종료됨
//
//   - Shutdown이 context.DeadlineExceeded 오류를 반환
//
//   - 클라이언트는 "connection reset" 또는 불완전한 응답을 볼 수 있음
//
//   - Run() 완료: Shutdown()이 완료된 후 Run()이 차단 해제되고 nil을 반환합니다.
//     이를 통해 main 고루틴 또는 배포 스크립트가 진행할 수 있습니다.
//
//   - 상태 정리: running 플래그가 false로 설정되어 앱 재사용이 가능합니다(권장하지 않음).
//
// 동시성:
//
//   - 스레드 안전: Run()이 차단되는 동안 다른 고루틴에서 호출하기에 안전합니다.
//     일반적인 패턴: main 고루틴에서 Run(), 신호 핸들러 고루틴에서 Shutdown().
//
//   - 단일 종료: Shutdown()을 여러 번 호출해도 안전하지만 중복됩니다.
//     후속 호출은 즉시 반환됩니다(멱등 작업).
//
//   - 잠금: 읽기 잠금을 사용하여 상태를 확인하고 Run() 작업을 차단하지 않습니다.
//     내부적으로 동기화를 처리하는 http.Server.Shutdown()에 위임합니다.
//
// 사용 패턴:
//   - 수동 종료: 내부 로직을 기반으로 애플리케이션이 트리거하는 종료
//   - 신호 핸들러: OS 신호(SIGINT/SIGTERM)가 종료를 트리거
//   - 헬스 체크: 종료 전에 로드 밸런서에서 서버를 우아하게 제거
//   - 롤링 배포: 서버 인스턴스를 교체하기 전에 연결 드레인
//   - 테스트: 테스트 서버의 깔끔한 해체
//
// 모범 사례:
//   - 무기한 대기를 방지하기 위해 항상 타임아웃 컨텍스트를 사용하세요.
//   - 예상되는 가장 긴 요청 기간을 기반으로 타임아웃을 선택하세요.
//   - 모니터링 및 디버깅을 위해 종료 시작/완료를 로깅하세요.
//   - 종료 전에 로드 밸런서에 트래픽 전송 중지를 알리세요.
//   - 종료 후 데이터베이스 연결 및 외부 리소스를 닫으세요.
//   - 핸들러에서 종료 중 조기 종료를 위해 컨텍스트 취소를 확인하세요.
//   - 타임아웃 설정을 검증하기 위해 부하 상태에서 종료 동작을 테스트하세요.
//
// 일반적인 사용 사례:
//   - 신호 기반 종료: syscall.SIGINT/SIGTERM 핸들러가 Shutdown() 호출
//   - Kubernetes 포드 종료: PreStop 후크가 종료 엔드포인트 호출
//   - 헬스 체크 저하: 비정상 표시, 대기, 그 다음 종료
//   - 배포 자동화: Blue-green 또는 롤링 업데이트 워크플로
//   - 개발: Ctrl+C가 우아한 종료 트리거
//
// 예제 - 타임아웃이 있는 기본 종료:
//
//	// 별도의 고루틴 또는 신호 핸들러에서
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	if err := app.Shutdown(ctx); err != nil {
//	    if errors.Is(err, context.DeadlineExceeded) {
//	        log.Println("Shutdown timeout: some connections were forcibly closed")
//	    } else {
//	        log.Printf("Shutdown error: %v", err)
//	    }
//	} else {
//	    log.Println("Shutdown completed successfully")
//	}
//
// 예제 - 신호 핸들러 패턴:
//
//	go func() {
//	    if err := app.Run(":8080"); err != nil {
//	        log.Fatalf("Server error: %v", err)
//	    }
//	}()
//
//	// 인터럽트 신호 대기
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	<-quit
//
//	log.Println("Shutting down server...")
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	if err := app.Shutdown(ctx); err != nil {
//	    log.Fatalf("Shutdown failed: %v", err)
//	}
//	log.Println("Server exited")
//
// 예제 - 헬스 체크 통합:
//
//	var healthy atomic.Bool
//	healthy.Store(true)
//
//	app.GET("/health", func(w http.ResponseWriter, r *http.Request) {
//	    if healthy.Load() {
//	        w.WriteHeader(http.StatusOK)
//	    } else {
//	        w.WriteHeader(http.StatusServiceUnavailable)
//	    }
//	})
//
//	// 종료 신호 시
//	healthy.Store(false)           // 비정상 표시
//	time.Sleep(10 * time.Second)   // 로드 밸런서가 감지할 때까지 대기
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	app.Shutdown(ctx)              // 연결 드레인
//
// 예제 - 핸들러 컨텍스트 확인:
//
//	app.GET("/long-task", func(w http.ResponseWriter, r *http.Request) {
//	    for i := 0; i < 100; i++ {
//	        // 종료가 진행 중인지 확인
//	        select {
//	        case <-r.Context().Done():
//	            log.Println("Handler cancelled during shutdown")
//	            http.Error(w, "Service shutting down", http.StatusServiceUnavailable)
//	            return
//	        default:
//	        }
//
//	        // 작업 수행
//	        time.Sleep(100 * time.Millisecond)
//	    }
//
//	    w.Write([]byte("Completed"))
//	})
//
// Kubernetes 예제:
//
//	# Kubernetes Pod spec에서
//	lifecycle:
//	  preStop:
//	    exec:
//	      command: ["/bin/sh", "-c", "sleep 15 && kill -SIGTERM 1"]
//
//	# 애플리케이션이 SIGTERM 처리
//	signal.Notify(quit, syscall.SIGTERM)
//	<-quit
//	app.Shutdown(context.WithTimeout(context.Background(), 30*time.Second))
//
// 모니터링 및 관찰 가능성:
//   - 타임스탬프와 함께 종료 시작 로깅
//   - 종료 기간 메트릭 추적
//   - 종료 타임아웃 시 경고
//   - 강제 닫힌 연결 수 모니터링
//   - 로그에 종료 이유 포함(신호, 헬스 체크, 수동)
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

// RunWithGracefulShutdown starts the HTTP server with automatic graceful shutdown on OS signals.
// This is the **recommended production method** for starting your web application. It combines
// server startup with intelligent signal handling, automatically draining connections when
// receiving SIGINT (Ctrl+C) or SIGTERM (common deployment signals) without additional code.
//
// This method simplifies production deployments by:
// - Handling Ctrl+C gracefully during development and testing
// - Responding to Docker/Kubernetes termination signals (SIGTERM)
// - Draining connections before process exit
// - Preventing dropped requests during deployments
// - Eliminating boilerplate signal handling code
//
// RunWithGracefulShutdown is ideal for:
// - Production applications (Docker, Kubernetes, systemd)
// - Development servers (clean Ctrl+C shutdown)
// - CI/CD pipelines (proper test cleanup)
// - Any deployment requiring zero-downtime updates
//
// Parameters:
//
//   - addr: Network address to listen on, in "host:port" format.
//     Same format as Run() method:
//
//   - ":8080" - Listen on all interfaces, port 8080 (most common)
//
//   - "localhost:8080" - Listen only on localhost (development)
//
//   - "0.0.0.0:8080" - Explicitly listen on all interfaces
//
//   - ":80" or ":443" - Standard HTTP/HTTPS ports (requires elevated privileges)
//
//   - timeout: Maximum duration to wait for active connections to close during shutdown.
//     This timeout applies from signal reception to forceful connection termination.
//     Common timeout values:
//
//   - 5 seconds: Fast shutdown, suitable for APIs with quick responses
//
//   - 15 seconds: Balanced for typical web applications
//
//   - 30 seconds: Recommended for most production applications (default recommendation)
//
//   - 60 seconds: Patient shutdown for applications with long-running requests
//
//   - 120+ seconds: Very patient, for file uploads, video processing, batch operations
//     If timeout expires, remaining connections are forcibly closed.
//
// Returns:
//   - error: Returns error if server fails to start or shutdown encounters errors.
//     Possible error scenarios:
//   - Server startup failure (port in use, permission denied, invalid address)
//   - Shutdown timeout exceeded (context.DeadlineExceeded wrapped in error)
//   - Network errors during connection cleanup
//     Returns nil when server shuts down cleanly after receiving signal.
//
// Behavior:
//
//   - Startup Phase:
//     1. Starts server in background goroutine using Run()
//     2. Registers signal handlers for SIGINT and SIGTERM
//     3. Blocks waiting for either server error or OS signal
//
//   - Signal Handling: Automatically responds to:
//
//   - SIGINT: Sent by Ctrl+C, shell interrupts, keyboard interrupt
//
//   - SIGTERM: Sent by kill command, Docker stop, Kubernetes pod termination,
//     systemd service stop, process managers, deployment scripts
//
//   - Shutdown Sequence:
//     1. Receives SIGINT or SIGTERM signal
//     2. Prints received signal name to console
//     3. Creates context with specified timeout
//     4. Calls Shutdown() to drain connections gracefully
//     5. Waits up to timeout duration for connections to close
//     6. Returns nil on success or error on failure
//
//   - Two Exit Paths:
//
//   - Server startup error: Returns immediately with error (port conflict, etc.)
//
//   - Signal received: Initiates graceful shutdown, returns after completion
//
//   - Blocking Call: This method blocks until server stops (via signal or error).
//     Must be called from main goroutine or dedicated server goroutine.
//
// Signal Details:
//
//   - SIGINT (Signal Interrupt):
//
//   - Triggered by: Ctrl+C in terminal, kill -INT <pid>
//
//   - Use case: Development, manual interruption, user-initiated shutdown
//
//   - Exit code: Typically 130 (128 + 2)
//
//   - SIGTERM (Signal Terminate):
//
//   - Triggered by: kill <pid>, Docker stop, Kubernetes pod deletion, systemd stop
//
//   - Use case: Automated deployments, orchestrator-managed lifecycle
//
//   - Exit code: Typically 143 (128 + 15)
//
//   - Most common in production environments
//
// Thread-Safety:
//   - Call from single goroutine (typically main()).
//   - Spawns internal goroutines for server and signal handling.
//   - Safe concurrent signal delivery (channel buffered, signal.Notify thread-safe).
//
// Deployment Integration:
//   - Docker: Responds to `docker stop` (sends SIGTERM, waits, then SIGKILL)
//   - Kubernetes: Handles pod termination (SIGTERM from kubelet)
//   - systemd: Responds to `systemctl stop` (sends SIGTERM)
//   - Heroku/PaaS: Responds to dyno/instance termination signals
//   - Process Managers (PM2, Supervisor): Handles managed process restarts
//
// Best Practices:
//   - Use this method for production instead of plain Run().
//   - Set timeout longer than longest expected request duration.
//   - Add timeout buffer for cleanup operations (DB commits, cache writes).
//   - Log startup and shutdown events for monitoring.
//   - Test shutdown behavior under load before production deployment.
//   - Consider health check degradation before signal shutdown.
//   - Clean up resources (DB connections, files) after method returns.
//
// Performance Considerations:
//   - Signal handling overhead is negligible (channel + goroutine).
//   - Shutdown performance depends on active connection count and timeout.
//   - Timeout too short: Risk of forcibly closed connections
//   - Timeout too long: Slower deployments, longer pod termination
//   - Balance timeout based on application request patterns.
//
// Common Use Cases:
//   - Production server: app.RunWithGracefulShutdown(":8080", 30*time.Second)
//   - Development server: app.RunWithGracefulShutdown("localhost:3000", 5*time.Second)
//   - Docker container: app.RunWithGracefulShutdown(":8080", 25*time.Second)
//   - Kubernetes pod: app.RunWithGracefulShutdown(":8080", 30*time.Second)
//   - Quick testing: app.RunWithGracefulShutdown(":9000", 1*time.Second)
//
// Example - Basic Production Server:
//
//	func main() {
//	    app := websvrutil.New()
//	    app.GET("/", homeHandler)
//	    app.GET("/api/users", usersHandler)
//
//	    log.Println("Starting server on :8080")
//	    if err := app.RunWithGracefulShutdown(":8080", 30*time.Second); err != nil {
//	        log.Fatalf("Server error: %v", err)
//	    }
//	    log.Println("Server shutdown complete")
//	}
//
// Example - With Resource Cleanup:
//
//	func main() {
//	    // Initialize dependencies
//	    db, err := sql.Open("postgres", dbURL)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    defer db.Close()  // Cleanup after server stops
//
//	    cache := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
//	    defer cache.Close()
//
//	    // Configure application
//	    app := websvrutil.New()
//	    app.GET("/api/data", func(w http.ResponseWriter, r *http.Request) {
//	        // Use db and cache
//	    })
//
//	    // Start with graceful shutdown
//	    log.Println("Server starting...")
//	    if err := app.RunWithGracefulShutdown(":8080", 30*time.Second); err != nil {
//	        log.Printf("Server error: %v", err)
//	        os.Exit(1)
//	    }
//
//	    log.Println("Cleanup complete, exiting")
//	}
//
// Example - Environment-Based Configuration:
//
//	func main() {
//	    port := os.Getenv("PORT")
//	    if port == "" {
//	        port = "8080"
//	    }
//
//	    shutdownTimeout := 30 * time.Second
//	    if timeoutStr := os.Getenv("SHUTDOWN_TIMEOUT"); timeoutStr != "" {
//	        if d, err := time.ParseDuration(timeoutStr); err == nil {
//	            shutdownTimeout = d
//	        }
//	    }
//
//	    app := websvrutil.New()
//	    // ... configure routes ...
//
//	    log.Printf("Starting server on port %s with %v shutdown timeout", port, shutdownTimeout)
//	    if err := app.RunWithGracefulShutdown(":"+port, shutdownTimeout); err != nil {
//	        log.Fatalf("Server failed: %v", err)
//	    }
//	}
//
// Example - Docker Container:
//
//	# Dockerfile
//	FROM golang:1.21-alpine
//	WORKDIR /app
//	COPY . .
//	RUN go build -o server .
//	EXPOSE 8080
//	CMD ["./server"]
//
//	# Go application
//	func main() {
//	    app := websvrutil.New()
//	    // ... configure routes ...
//
//	    // Docker sends SIGTERM on `docker stop`
//	    // Default Docker stop timeout is 10 seconds before SIGKILL
//	    // Use shorter timeout to ensure graceful completion
//	    if err := app.RunWithGracefulShutdown(":8080", 8*time.Second); err != nil {
//	        log.Fatal(err)
//	    }
//	}
//
// Example - Kubernetes Deployment:
//
//	# deployment.yaml
//	apiVersion: apps/v1
//	kind: Deployment
//	spec:
//	  template:
//	    spec:
//	      terminationGracePeriodSeconds: 35  # Allow 35s for graceful shutdown
//	      containers:
//	      - name: app
//	        image: myapp:latest
//	        ports:
//	        - containerPort: 8080
//	        lifecycle:
//	          preStop:
//	            exec:
//	              command: ["/bin/sh", "-c", "sleep 5"]  # Small delay for load balancer
//
//	# Go application
//	func main() {
//	    app := websvrutil.New()
//	    // ... configure routes ...
//
//	    // Kubernetes sends SIGTERM on pod deletion
//	    // Use timeout less than terminationGracePeriodSeconds
//	    if err := app.RunWithGracefulShutdown(":8080", 30*time.Second); err != nil {
//	        log.Fatal(err)
//	    }
//	}
//
// Example - With Health Check:
//
//	var shutdownRequested atomic.Bool
//
//	func main() {
//	    app := websvrutil.New()
//
//	    // Health check endpoint (used by load balancer)
//	    app.GET("/health", func(w http.ResponseWriter, r *http.Request) {
//	        if shutdownRequested.Load() {
//	            w.WriteHeader(http.StatusServiceUnavailable)
//	            return
//	        }
//	        w.WriteHeader(http.StatusOK)
//	    })
//
//	    app.GET("/api/data", dataHandler)
//
//	    // Run with graceful shutdown
//	    // Could enhance to mark unhealthy before actual shutdown
//	    if err := app.RunWithGracefulShutdown(":8080", 30*time.Second); err != nil {
//	        log.Fatal(err)
//	    }
//	}
//
// Monitoring and Debugging:
//   - Log signal reception with timestamp for debugging deployments.
//   - Track shutdown duration metrics to optimize timeout setting.
//   - Monitor forced connection closures (indicates timeout too short).
//   - Alert on frequent restarts or shutdown failures.
//   - Include deployment context in logs (version, environment, reason).
//
// Alternative Patterns:
//   - Custom Signal Handling: Use Run() + manual signal.Notify() for custom logic
//   - Multiple Signals: Add SIGHUP, SIGUSR1 for reload/reconfiguration
//   - Health Check Integration: Mark unhealthy before shutdown
//   - Graceful Restart: Fork new process before shutting down old one
//
// Limitations:
//   - Only handles SIGINT and SIGTERM (not SIGHUP, SIGUSR1, etc.)
//   - Single timeout for all connections (can't prioritize critical connections)
//   - No built-in health check degradation (implement separately if needed)
//   - Cannot cancel shutdown once initiated (signal reception commits to shutdown)
//
// RunWithGracefulShutdown은 OS 신호 시 자동 우아한 종료를 지원하는 HTTP 서버를 시작합니다.
// 이는 웹 애플리케이션을 시작하는 **권장 프로덕션 메서드**입니다.
// 서버 시작과 지능적인 신호 처리를 결합하여 SIGINT(Ctrl+C) 또는 SIGTERM(일반적인 배포 신호)을
// 받을 때 추가 코드 없이 자동으로 연결을 드레인합니다.
//
// 이 메서드는 다음을 통해 프로덕션 배포를 단순화합니다:
// - 개발 및 테스트 중 Ctrl+C를 우아하게 처리
// - Docker/Kubernetes 종료 신호(SIGTERM)에 응답
// - 프로세스 종료 전 연결 드레인
// - 배포 중 중단된 요청 방지
// - 보일러플레이트 신호 처리 코드 제거
//
// RunWithGracefulShutdown은 다음에 이상적입니다:
// - 프로덕션 애플리케이션(Docker, Kubernetes, systemd)
// - 개발 서버(깨끗한 Ctrl+C 종료)
// - CI/CD 파이프라인(적절한 테스트 정리)
// - 무중단 업데이트가 필요한 모든 배포
//
// 매개변수:
//
//   - addr: "host:port" 형식으로 수신할 네트워크 주소.
//     Run() 메서드와 동일한 형식:
//
//   - ":8080" - 모든 인터페이스에서 수신, 포트 8080(가장 일반적)
//
//   - "localhost:8080" - localhost에서만 수신(개발)
//
//   - "0.0.0.0:8080" - 모든 인터페이스에서 명시적으로 수신
//
//   - ":80" 또는 ":443" - 표준 HTTP/HTTPS 포트(높은 권한 필요)
//
//   - timeout: 종료 중 활성 연결이 닫힐 때까지 기다릴 최대 시간.
//     이 타임아웃은 신호 수신부터 강제 연결 종료까지 적용됩니다.
//     일반적인 타임아웃 값:
//
//   - 5초: 빠른 종료, 빠른 응답을 가진 API에 적합
//
//   - 15초: 일반적인 웹 애플리케이션에 균형 잡힌 값
//
//   - 30초: 대부분의 프로덕션 애플리케이션에 권장(기본 권장사항)
//
//   - 60초: 장기 실행 요청이 있는 애플리케이션을 위한 참을성 있는 종료
//
//   - 120초 이상: 파일 업로드, 비디오 처리, 배치 작업을 위한 매우 참을성 있는 종료
//     타임아웃이 만료되면 남은 연결이 강제로 닫힙니다.
//
// 반환값:
//   - error: 서버가 시작에 실패하거나 종료 중 오류가 발생하면 오류를 반환합니다.
//     가능한 오류 시나리오:
//   - 서버 시작 실패(포트 사용 중, 권한 거부, 잘못된 주소)
//   - 종료 타임아웃 초과(오류로 래핑된 context.DeadlineExceeded)
//   - 연결 정리 중 네트워크 오류
//     신호를 받은 후 서버가 깨끗하게 종료되면 nil을 반환합니다.
//
// 동작 방식:
//
//   - 시작 단계:
//     1. Run()을 사용하여 백그라운드 고루틴에서 서버 시작
//     2. SIGINT 및 SIGTERM에 대한 신호 핸들러 등록
//     3. 서버 오류 또는 OS 신호를 기다리며 차단
//
//   - 신호 처리: 다음에 자동으로 응답:
//
//   - SIGINT: Ctrl+C, 쉘 인터럽트, 키보드 인터럽트에 의해 전송
//
//   - SIGTERM: kill 명령, Docker stop, Kubernetes 포드 종료,
//     systemd 서비스 중지, 프로세스 관리자, 배포 스크립트에 의해 전송
//
//   - 종료 순서:
//     1. SIGINT 또는 SIGTERM 신호 수신
//     2. 수신된 신호 이름을 콘솔에 출력
//     3. 지정된 타임아웃으로 컨텍스트 생성
//     4. Shutdown()을 호출하여 연결을 우아하게 드레인
//     5. 연결이 닫힐 때까지 타임아웃 기간만큼 대기
//     6. 성공 시 nil 또는 실패 시 오류 반환
//
//   - 두 가지 종료 경로:
//
//   - 서버 시작 오류: 오류와 함께 즉시 반환(포트 충돌 등)
//
//   - 신호 수신: 우아한 종료 시작, 완료 후 반환
//
//   - 차단 호출: 이 메서드는 서버가 중지될 때까지(신호 또는 오류를 통해) 차단됩니다.
//     main 고루틴 또는 전용 서버 고루틴에서 호출되어야 합니다.
//
// 신호 상세:
//
//   - SIGINT(신호 인터럽트):
//
//   - 트리거: 터미널에서 Ctrl+C, kill -INT <pid>
//
//   - 사용 사례: 개발, 수동 중단, 사용자 시작 종료
//
//   - 종료 코드: 일반적으로 130(128 + 2)
//
//   - SIGTERM(신호 종료):
//
//   - 트리거: kill <pid>, Docker stop, Kubernetes 포드 삭제, systemd stop
//
//   - 사용 사례: 자동화된 배포, 오케스트레이터 관리 생명주기
//
//   - 종료 코드: 일반적으로 143(128 + 15)
//
//   - 프로덕션 환경에서 가장 일반적
//
// 스레드 안전성:
//   - 단일 고루틴에서 호출(일반적으로 main()).
//   - 서버 및 신호 처리를 위한 내부 고루틴 생성.
//   - 안전한 동시 신호 전달(채널 버퍼링, signal.Notify 스레드 안전).
//
// 배포 통합:
//   - Docker: `docker stop`에 응답(SIGTERM 전송, 대기, 그 다음 SIGKILL)
//   - Kubernetes: 포드 종료 처리(kubelet의 SIGTERM)
//   - systemd: `systemctl stop`에 응답(SIGTERM 전송)
//   - Heroku/PaaS: dyno/인스턴스 종료 신호에 응답
//   - 프로세스 관리자(PM2, Supervisor): 관리되는 프로세스 재시작 처리
//
// 모범 사례:
//   - 프로덕션에서는 일반 Run() 대신 이 메서드를 사용하세요.
//   - 가장 긴 예상 요청 기간보다 타임아웃을 길게 설정하세요.
//   - 정리 작업(DB 커밋, 캐시 쓰기)을 위한 타임아웃 버퍼를 추가하세요.
//   - 모니터링을 위해 시작 및 종료 이벤트를 로깅하세요.
//   - 프로덕션 배포 전에 부하 상태에서 종료 동작을 테스트하세요.
//   - 신호 종료 전에 헬스 체크 저하를 고려하세요.
//   - 메서드 반환 후 리소스(DB 연결, 파일)를 정리하세요.
//
// 성능 고려사항:
//   - 신호 처리 오버헤드는 무시할 수 있습니다(채널 + 고루틴).
//   - 종료 성능은 활성 연결 수와 타임아웃에 따라 다릅니다.
//   - 타임아웃이 너무 짧으면: 강제로 닫힌 연결의 위험
//   - 타임아웃이 너무 길면: 느린 배포, 긴 포드 종료
//   - 애플리케이션 요청 패턴을 기반으로 타임아웃의 균형을 맞추세요.
//
// 일반적인 사용 사례:
//   - 프로덕션 서버: app.RunWithGracefulShutdown(":8080", 30*time.Second)
//   - 개발 서버: app.RunWithGracefulShutdown("localhost:3000", 5*time.Second)
//   - Docker 컨테이너: app.RunWithGracefulShutdown(":8080", 25*time.Second)
//   - Kubernetes 포드: app.RunWithGracefulShutdown(":8080", 30*time.Second)
//   - 빠른 테스트: app.RunWithGracefulShutdown(":9000", 1*time.Second)
//
// 예제 - 기본 프로덕션 서버:
//
//	func main() {
//	    app := websvrutil.New()
//	    app.GET("/", homeHandler)
//	    app.GET("/api/users", usersHandler)
//
//	    log.Println("Starting server on :8080")
//	    if err := app.RunWithGracefulShutdown(":8080", 30*time.Second); err != nil {
//	        log.Fatalf("Server error: %v", err)
//	    }
//	    log.Println("Server shutdown complete")
//	}
//
// 예제 - 리소스 정리와 함께:
//
//	func main() {
//	    // 종속성 초기화
//	    db, err := sql.Open("postgres", dbURL)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    defer db.Close()  // 서버 중지 후 정리
//
//	    cache := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
//	    defer cache.Close()
//
//	    // 애플리케이션 구성
//	    app := websvrutil.New()
//	    app.GET("/api/data", func(w http.ResponseWriter, r *http.Request) {
//	        // db 및 cache 사용
//	    })
//
//	    // 우아한 종료와 함께 시작
//	    log.Println("Server starting...")
//	    if err := app.RunWithGracefulShutdown(":8080", 30*time.Second); err != nil {
//	        log.Printf("Server error: %v", err)
//	        os.Exit(1)
//	    }
//
//	    log.Println("Cleanup complete, exiting")
//	}
//
// 모니터링 및 디버깅:
//   - 배포 디버깅을 위해 타임스탬프와 함께 신호 수신을 로깅하세요.
//   - 타임아웃 설정을 최적화하기 위해 종료 기간 메트릭을 추적하세요.
//   - 강제 연결 종료를 모니터링하세요(타임아웃이 너무 짧음을 나타냄).
//   - 빈번한 재시작 또는 종료 실패 시 경고하세요.
//   - 로그에 배포 컨텍스트(버전, 환경, 이유)를 포함하세요.
//
// 대안 패턴:
//   - 커스텀 신호 처리: 커스텀 로직을 위해 Run() + 수동 signal.Notify() 사용
//   - 여러 신호: 리로드/재구성을 위해 SIGHUP, SIGUSR1 추가
//   - 헬스 체크 통합: 종료 전에 비정상 표시
//   - 우아한 재시작: 이전 프로세스를 종료하기 전에 새 프로세스 포크
//
// 제한사항:
//   - SIGINT 및 SIGTERM만 처리(SIGHUP, SIGUSR1 등은 아님)
//   - 모든 연결에 단일 타임아웃(중요한 연결의 우선순위를 지정할 수 없음)
//   - 내장 헬스 체크 저하 없음(필요한 경우 별도로 구현)
//   - 시작된 종료를 취소할 수 없음(신호 수신이 종료를 커밋)
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

// buildHandler builds the final HTTP handler by applying all middleware in reverse order.
// This is an internal helper method that constructs the complete request processing chain
// by wrapping the router with all registered middleware functions.
//
// buildHandler is called during:
// - Server startup (Run, RunWithGracefulShutdown) to create the initial handler
// - Each incoming request (ServeHTTP) to ensure latest middleware configuration
//
// Purpose:
//   - Constructs the complete HTTP handler chain for request processing
//   - Applies middleware in LIFO (Last In, First Out) order
//   - Provides fallback 404 handler when no routes configured
//   - Ensures middleware wrapping order matches registration order semantics
//
// Returns:
//   - http.Handler: Complete handler chain ready to process HTTP requests.
//     The returned handler represents the full middleware stack wrapped around
//     the router, with outermost middleware added last via Use().
//
// Middleware Application Order:
//
//   - Middleware is stored in registration order (first Use() call = index 0)
//
//   - Applied in reverse order (last Use() call wraps everything)
//
//   - This creates intuitive "outer to inner" execution during requests
//
//     Example registration:
//     app.Use(LoggingMiddleware)    // Added first, index 0
//     app.Use(AuthMiddleware)       // Added second, index 1
//     app.Use(CompressionMiddleware) // Added third, index 2
//
//     Resulting handler chain (buildHandler applies in reverse):
//     CompressionMiddleware(        // Applied last, wraps everything
//     AuthMiddleware(              // Applied second
//     LoggingMiddleware(         // Applied first
//     router                   // Core router at center
//     )
//     )
//     )
//
//     Request execution flow (outer to inner):
//     1. Request arrives → CompressionMiddleware executes first
//     2. Calls next → AuthMiddleware executes second
//     3. Calls next → LoggingMiddleware executes third
//     4. Calls next → Router handles request
//     5. Response returns through same chain (inner to outer)
//
// Behavior:
//   - Router Priority: Uses app.router if configured, otherwise creates 404 handler
//   - Reverse Iteration: Loops through middleware slice from end to start (i--)
//   - Sequential Wrapping: Each middleware wraps the previous handler
//   - Empty Middleware: Returns unwrapped router if no middleware registered
//   - Nil Router: Provides http.NotFound fallback to prevent nil pointer panics
//
// Default Handler (No Router):
//   - Returns 404 Not Found for all requests
//   - Uses http.NotFound() for standard error response
//   - Prevents server startup errors when routes not yet configured
//   - Safe for testing, development, and gradual route registration
//
// Thread-Safety:
//   - This method is NOT thread-safe by itself.
//   - Caller MUST hold app.mu.RLock() before calling.
//   - ServeHTTP properly acquires read lock before calling this method.
//   - Never call this method directly without proper locking.
//
// Performance Considerations:
//   - Called on EVERY request in ServeHTTP implementation.
//   - Rebuilds handler chain each time (no caching).
//   - Rebuilding ensures middleware changes take effect immediately.
//   - Overhead is minimal: O(n) where n = middleware count.
//   - Typical applications have 3-10 middleware, negligible performance impact.
//   - For extreme performance, consider handler caching (advanced optimization).
//
// Common Middleware Patterns:
//   - Logging: Log request/response (outermost, should wrap everything)
//   - Recovery: Panic recovery (very outer, catches all panics)
//   - CORS: Cross-origin headers (outer, affects all routes)
//   - Authentication: User verification (middle layer)
//   - Authorization: Permission checks (after auth)
//   - Compression: Response compression (can be anywhere)
//   - Rate Limiting: Request throttling (outer, before processing)
//
// Example - Middleware Execution Order:
//
//	app := websvrutil.New()
//
//	// Register middleware (order matters!)
//	app.Use(func(next http.Handler) http.Handler {
//	    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	        fmt.Println("1. First middleware - before")
//	        next.ServeHTTP(w, r)
//	        fmt.Println("1. First middleware - after")
//	    })
//	})
//
//	app.Use(func(next http.Handler) http.Handler {
//	    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	        fmt.Println("  2. Second middleware - before")
//	        next.ServeHTTP(w, r)
//	        fmt.Println("  2. Second middleware - after")
//	    })
//	})
//
//	app.GET("/test", func(w http.ResponseWriter, r *http.Request) {
//	    fmt.Println("    3. Handler executed")
//	    w.Write([]byte("OK"))
//	})
//
//	// Request to /test outputs:
//	// 1. First middleware - before
//	//   2. Second middleware - before
//	//     3. Handler executed
//	//   2. Second middleware - after
//	// 1. First middleware - after
//
// Example - No Router Configured:
//
//	app := websvrutil.New()
//	app.Use(LoggingMiddleware)  // Middleware registered
//	// No routes configured yet!
//
//	handler := app.buildHandler()
//	// Returns: LoggingMiddleware wrapping http.NotFound
//	// All requests → logging → 404 Not Found
//
// Example - Middleware Wrapping Visualization:
//
//	// Registration order
//	app.Use(A)  // Middleware A
//	app.Use(B)  // Middleware B
//	app.Use(C)  // Middleware C
//
//	// buildHandler creates:
//	// C(B(A(router)))
//
//	// Execution flow:
//	// Request → C.before → B.before → A.before → router → A.after → B.after → C.after → Response
//
// Internal Implementation Notes:
//   - Starts with router (or 404 fallback) as base handler
//   - Iterates middleware slice backwards: for i := len-1; i >= 0; i--
//   - Each iteration: handler = middleware[i](handler)
//   - Final handler is fully wrapped chain ready for http.Server
//
// Why Reverse Order Application?
//   - Intuitive developer experience: first Use() call = outermost middleware
//   - Matches common framework conventions (Express.js, Gin, Echo)
//   - Natural "wrap everything" pattern for logging, recovery, CORS
//   - Avoids confusing "last registered runs first" semantics
//
// Comparison with Other Patterns:
//   - Forward Application: Would require explicit middleware ordering
//   - Middleware Priority: Could add priority/weight system (more complex)
//   - Conditional Middleware: Could add path-based middleware (route-specific)
//   - This implementation: Simple, predictable, sufficient for most use cases
//
// Best Practices:
//   - Register middleware in logical order (logging, recovery, auth, etc.)
//   - Keep middleware count reasonable (< 15 for performance)
//   - Ensure all middleware call next.ServeHTTP() to continue chain
//   - Use early returns in middleware to short-circuit (auth failures, etc.)
//   - Test middleware execution order with debug logging
//
// Debugging Tips:
//   - Add logging to see middleware execution order
//   - Verify each middleware calls next.ServeHTTP()
//   - Check for middleware that never calls next (breaks chain)
//   - Use middleware wrappers for consistent logging/timing
//
// buildHandler는 모든 미들웨어를 역순으로 적용하여 최종 HTTP 핸들러를 구축합니다.
// 이는 등록된 모든 미들웨어 함수로 라우터를 래핑하여 완전한 요청 처리 체인을 구성하는
// 내부 헬퍼 메서드입니다.
//
// buildHandler는 다음 상황에서 호출됩니다:
// - 서버 시작 시(Run, RunWithGracefulShutdown) 초기 핸들러 생성
// - 각 들어오는 요청 시(ServeHTTP) 최신 미들웨어 구성 보장
//
// 목적:
//   - 요청 처리를 위한 완전한 HTTP 핸들러 체인 구성
//   - 미들웨어를 LIFO(Last In, First Out) 순서로 적용
//   - 경로가 구성되지 않았을 때 폴백 404 핸들러 제공
//   - 미들웨어 래핑 순서가 등록 순서 의미와 일치하도록 보장
//
// 반환값:
//   - http.Handler: HTTP 요청을 처리할 준비가 된 완전한 핸들러 체인.
//     반환된 핸들러는 라우터를 래핑한 전체 미들웨어 스택을 나타내며,
//     가장 바깥쪽 미들웨어는 Use()를 통해 마지막에 추가됩니다.
//
// 미들웨어 적용 순서:
//
//   - 미들웨어는 등록 순서대로 저장됩니다(첫 번째 Use() 호출 = 인덱스 0)
//
//   - 역순으로 적용됩니다(마지막 Use() 호출이 모든 것을 래핑)
//
//   - 이는 요청 시 직관적인 "외부에서 내부로" 실행을 만듭니다
//
//     등록 예제:
//     app.Use(LoggingMiddleware)    // 첫 번째 추가, 인덱스 0
//     app.Use(AuthMiddleware)       // 두 번째 추가, 인덱스 1
//     app.Use(CompressionMiddleware) // 세 번째 추가, 인덱스 2
//
//     결과 핸들러 체인(buildHandler가 역순으로 적용):
//     CompressionMiddleware(        // 마지막 적용, 모든 것을 래핑
//     AuthMiddleware(              // 두 번째 적용
//     LoggingMiddleware(         // 첫 번째 적용
//     router                   // 중심의 핵심 라우터
//     )
//     )
//     )
//
//     요청 실행 흐름(외부에서 내부로):
//     1. 요청 도착 → CompressionMiddleware가 먼저 실행
//     2. next 호출 → AuthMiddleware가 두 번째 실행
//     3. next 호출 → LoggingMiddleware가 세 번째 실행
//     4. next 호출 → 라우터가 요청 처리
//     5. 응답은 동일한 체인을 통해 반환(내부에서 외부로)
//
// 동작 방식:
//   - 라우터 우선순위: 구성된 경우 app.router 사용, 그렇지 않으면 404 핸들러 생성
//   - 역방향 반복: 미들웨어 슬라이스를 끝에서 시작으로 반복(i--)
//   - 순차적 래핑: 각 미들웨어가 이전 핸들러를 래핑
//   - 빈 미들웨어: 미들웨어가 등록되지 않은 경우 래핑되지 않은 라우터 반환
//   - nil 라우터: nil 포인터 패닉을 방지하기 위해 http.NotFound 폴백 제공
//
// 기본 핸들러(라우터 없음):
//   - 모든 요청에 대해 404 Not Found 반환
//   - 표준 오류 응답을 위해 http.NotFound() 사용
//   - 경로가 아직 구성되지 않았을 때 서버 시작 오류 방지
//   - 테스트, 개발 및 점진적 경로 등록에 안전
//
// 스레드 안전성:
//   - 이 메서드는 그 자체로는 스레드 안전하지 않습니다.
//   - 호출자는 호출 전에 app.mu.RLock()을 보유해야 합니다.
//   - ServeHTTP는 이 메서드를 호출하기 전에 읽기 잠금을 적절히 획득합니다.
//   - 적절한 잠금 없이 이 메서드를 직접 호출하지 마세요.
//
// 성능 고려사항:
//   - ServeHTTP 구현에서 모든 요청마다 호출됩니다.
//   - 매번 핸들러 체인을 재구축합니다(캐싱 없음).
//   - 재구축은 미들웨어 변경이 즉시 적용되도록 보장합니다.
//   - 오버헤드는 최소: O(n), 여기서 n = 미들웨어 수.
//   - 일반적인 애플리케이션은 3-10개의 미들웨어를 가지며, 성능 영향은 무시할 수 있습니다.
//   - 극한의 성능을 위해 핸들러 캐싱을 고려하세요(고급 최적화).
//
// 일반적인 미들웨어 패턴:
//   - 로깅: 요청/응답 로깅(가장 바깥쪽, 모든 것을 래핑해야 함)
//   - 복구: 패닉 복구(매우 바깥쪽, 모든 패닉 캐치)
//   - CORS: 교차 출처 헤더(바깥쪽, 모든 경로에 영향)
//   - 인증: 사용자 확인(중간 레이어)
//   - 권한 부여: 권한 확인(인증 후)
//   - 압축: 응답 압축(어디든지 가능)
//   - 속도 제한: 요청 스로틀링(바깥쪽, 처리 전)
//
// 모범 사례:
//   - 논리적 순서로 미들웨어를 등록하세요(로깅, 복구, 인증 등).
//   - 미들웨어 수를 합리적으로 유지하세요(성능을 위해 < 15개).
//   - 모든 미들웨어가 next.ServeHTTP()를 호출하여 체인을 계속하도록 하세요.
//   - 미들웨어에서 조기 반환을 사용하여 단락시키세요(인증 실패 등).
//   - 디버그 로깅으로 미들웨어 실행 순서를 테스트하세요.
//
// 디버깅 팁:
//   - 미들웨어 실행 순서를 보려면 로깅을 추가하세요.
//   - 각 미들웨어가 next.ServeHTTP()를 호출하는지 확인하세요.
//   - next를 호출하지 않는 미들웨어를 확인하세요(체인을 끊음).
//   - 일관된 로깅/타이밍을 위해 미들웨어 래퍼를 사용하세요.
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

// ServeHTTP implements the http.Handler interface, making App compatible with standard Go HTTP servers.
// This is the **entry point for every HTTP request** processed by the application.
//
// ServeHTTP is automatically called by Go's http.Server for each incoming request.
// You typically don't call this method directly; instead, it's invoked by the server
// infrastructure when handling connections.
//
// Purpose:
//   - Implements http.Handler interface for compatibility with net/http ecosystem
//   - Serves as the single entry point for all HTTP requests
//   - Builds the complete handler chain (router + middleware) on each request
//   - Stores application context in request context for downstream handlers
//   - Provides thread-safe access to App state during request processing
//
// Interface Compliance:
//   - Satisfies http.Handler interface: ServeHTTP(ResponseWriter, *Request)
//   - Enables App to be used anywhere http.Handler is accepted:
//   - http.Server.Handler field
//   - http.ListenAndServe(addr, app)
//   - Middleware wrapping: middleware(app)
//   - Testing: httptest.NewServer(app)
//   - Reverse proxies, API gateways, etc.
//
// Parameters:
//
//   - w http.ResponseWriter: Response writer for sending HTTP response to client.
//     Used by handlers to write status codes, headers, and body content.
//     Passed through middleware chain to final route handler.
//
//   - r *http.Request: Incoming HTTP request containing method, URL, headers, body.
//     Enhanced with application context before passing to handler chain.
//     Contains all request metadata needed for routing and processing.
//
// Returns:
//   - (none): ServeHTTP returns nothing per http.Handler interface contract.
//     Response is written directly to ResponseWriter.
//     Errors should be handled by writing appropriate HTTP error responses.
//
// Behavior:
//  1. **Acquire Read Lock**: Locks app.mu.RLock() for thread-safe handler building
//  2. **Build Handler Chain**: Calls buildHandler() to construct middleware + router chain
//  3. **Release Read Lock**: Unlocks app.mu.RUnlock() after handler construction
//  4. **Store App Context**: Adds App instance to request context with key "app"
//  5. **Update Request**: Creates new request with enhanced context
//  6. **Delegate Handling**: Calls handler.ServeHTTP(w, r) to process request
//
// Context Storage:
//   - Stores App instance in request context: context.WithValue(r.Context(), "app", a)
//   - Purpose: Enable template rendering and app access in handlers
//   - Retrieval in handlers: app := r.Context().Value("app").(*websvrutil.App)
//   - Available throughout request lifecycle (middleware, handlers, helpers)
//   - Useful for: template rendering, app configuration access, shared resources
//
// Thread-Safety:
//   - Fully thread-safe for concurrent request handling.
//   - Read lock (RLock) allows multiple simultaneous requests.
//   - Handler building is isolated per-request (no shared state mutation).
//   - Safe for high-concurrency environments (thousands of concurrent requests).
//   - Write operations (Use, GET, POST, etc.) require full write lock, blocking ServeHTTP.
//
// Performance Characteristics:
//   - Lock Overhead: RLock/RUnlock is very fast (uncontended: ~20ns)
//   - Handler Building: O(n) where n = middleware count (typically 3-10)
//   - Context Creation: Negligible overhead (~50ns)
//   - Overall Overhead: < 1μs for typical configurations
//   - Scalability: Handles millions of requests/second on modern hardware
//
// Usage with http.Server:
//
//	server := &http.Server{
//	    Addr:    ":8080",
//	    Handler: app,  // App implements http.Handler via ServeHTTP
//	}
//	server.ListenAndServe()
//
// Usage with http.ListenAndServe:
//
//	http.ListenAndServe(":8080", app)  // App used as http.Handler
//
// Usage in Testing:
//
//	func TestHandler(t *testing.T) {
//	    app := websvrutil.New()
//	    app.GET("/test", func(w http.ResponseWriter, r *http.Request) {
//	        w.Write([]byte("OK"))
//	    })
//
//	    // ServeHTTP called automatically by test infrastructure
//	    req := httptest.NewRequest("GET", "/test", nil)
//	    rec := httptest.NewRecorder()
//	    app.ServeHTTP(rec, req)  // Direct invocation in tests
//
//	    assert.Equal(t, 200, rec.Code)
//	    assert.Equal(t, "OK", rec.Body.String())
//	}
//
// Middleware Execution Flow:
//
//	// With registered middleware
//	app.Use(LoggingMiddleware)
//	app.Use(AuthMiddleware)
//	app.GET("/api/data", dataHandler)
//
//	// Request arrives → ServeHTTP called
//	// 1. Acquire read lock
//	// 2. Build: AuthMiddleware(LoggingMiddleware(router))
//	// 3. Release read lock
//	// 4. Store app in context
//	// 5. Execute: Auth → Logging → Router → dataHandler
//	// 6. Response returns through same chain
//
// Context Access in Handlers:
//
//	func myHandler(w http.ResponseWriter, r *http.Request) {
//	    // Retrieve app from context (stored by ServeHTTP)
//	    app := r.Context().Value("app").(*websvrutil.App)
//
//	    // Access template engine
//	    templates := app.TemplateEngine()
//
//	    // Use for rendering
//	    // ...
//	}
//
// Common Use Cases:
//   - Standard HTTP Server: server.Handler = app
//   - Middleware Wrapping: wrappedApp := someMiddleware(app)
//   - Reverse Proxy Backend: proxy target pointing to app
//   - Testing: httptest.NewServer(app) or direct ServeHTTP calls
//   - Embedded Servers: Run app alongside other services
//
// Integration Examples:
//
//	// Example 1: Standard Server
//	app := websvrutil.New()
//	app.GET("/", homeHandler)
//	http.ListenAndServe(":8080", app)  // ServeHTTP called per request
//
//	// Example 2: Custom Server
//	server := &http.Server{
//	    Addr:         ":8080",
//	    Handler:      app,
//	    ReadTimeout:  10 * time.Second,
//	    WriteTimeout: 10 * time.Second,
//	}
//	server.ListenAndServe()
//
//	// Example 3: Wrapped with External Middleware
//	app := websvrutil.New()
//	wrapped := http.TimeoutHandler(app, 30*time.Second, "Timeout")
//	http.ListenAndServe(":8080", wrapped)
//
//	// Example 4: Multiple Apps (different ports)
//	publicApp := websvrutil.New()
//	publicApp.GET("/", publicHandler)
//
//	adminApp := websvrutil.New()
//	adminApp.GET("/admin", adminHandler)
//
//	go http.ListenAndServe(":8080", publicApp)   // Public on 8080
//	http.ListenAndServe(":9090", adminApp)       // Admin on 9090
//
// Performance Optimization Notes:
//   - Handler building per-request ensures middleware changes take effect immediately
//   - Alternative: Cache handler and rebuild on middleware changes (more complex)
//   - Current approach: Simple, correct, sufficient for 99% of applications
//   - For extreme performance (1M+ req/s): Consider handler caching with invalidation
//
// Compatibility Notes:
//   - Works with all Go HTTP infrastructure (http.Server, http.Client, httptest, etc.)
//   - Compatible with third-party middleware (gorilla, negroni, etc.)
//   - Can be embedded in larger applications or frameworks
//   - Safe to use with http2, HTTP/3 (QUIC) when available
//
// Error Handling:
//   - ServeHTTP itself does not return errors (interface contract)
//   - Errors must be handled by writing HTTP error responses
//   - Panics in handlers should be caught by recovery middleware
//   - Network errors handled by http.Server automatically
//
// Best Practices:
//   - Don't call ServeHTTP directly except in tests
//   - Let http.Server manage ServeHTTP invocation
//   - Use recovery middleware to catch panics
//   - Log request/response in middleware, not in ServeHTTP
//   - Keep handler building fast (minimize middleware count)
//
// Debugging Tips:
//   - Add logging before/after handler.ServeHTTP() to trace requests
//   - Use request ID middleware for request correlation
//   - Monitor lock contention (should be near zero)
//   - Profile ServeHTTP execution time for performance tuning
//
// ServeHTTP는 http.Handler 인터페이스를 구현하여 App을 표준 Go HTTP 서버와 호환되게 합니다.
// 이는 애플리케이션에서 처리하는 **모든 HTTP 요청의 진입점**입니다.
//
// ServeHTTP는 각 들어오는 요청에 대해 Go의 http.Server에 의해 자동으로 호출됩니다.
// 일반적으로 이 메서드를 직접 호출하지 않으며, 대신 연결을 처리할 때
// 서버 인프라에 의해 호출됩니다.
//
// 목적:
//   - net/http 생태계와의 호환성을 위해 http.Handler 인터페이스 구현
//   - 모든 HTTP 요청의 단일 진입점 역할
//   - 각 요청에서 완전한 핸들러 체인(라우터 + 미들웨어) 구축
//   - 다운스트림 핸들러를 위해 요청 컨텍스트에 애플리케이션 컨텍스트 저장
//   - 요청 처리 중 App 상태에 대한 스레드 안전 액세스 제공
//
// 인터페이스 준수:
//   - http.Handler 인터페이스를 만족: ServeHTTP(ResponseWriter, *Request)
//   - http.Handler가 허용되는 모든 곳에서 App을 사용할 수 있게 합니다:
//   - http.Server.Handler 필드
//   - http.ListenAndServe(addr, app)
//   - 미들웨어 래핑: middleware(app)
//   - 테스팅: httptest.NewServer(app)
//   - 리버스 프록시, API 게이트웨이 등
//
// 매개변수:
//
//   - w http.ResponseWriter: 클라이언트에 HTTP 응답을 보내기 위한 응답 작성기.
//     핸들러가 상태 코드, 헤더 및 본문 내용을 작성하는 데 사용됩니다.
//     미들웨어 체인을 통해 최종 경로 핸들러로 전달됩니다.
//
//   - r *http.Request: 메서드, URL, 헤더, 본문을 포함하는 들어오는 HTTP 요청.
//     핸들러 체인으로 전달하기 전에 애플리케이션 컨텍스트로 향상됩니다.
//     라우팅 및 처리에 필요한 모든 요청 메타데이터를 포함합니다.
//
// 반환값:
//   - (없음): ServeHTTP는 http.Handler 인터페이스 계약에 따라 아무것도 반환하지 않습니다.
//     응답은 ResponseWriter에 직접 작성됩니다.
//     오류는 적절한 HTTP 오류 응답을 작성하여 처리해야 합니다.
//
// 동작 방식:
//  1. **읽기 잠금 획득**: 스레드 안전 핸들러 구축을 위해 app.mu.RLock() 잠금
//  2. **핸들러 체인 구축**: buildHandler()를 호출하여 미들웨어 + 라우터 체인 구성
//  3. **읽기 잠금 해제**: 핸들러 구성 후 app.mu.RUnlock() 잠금 해제
//  4. **앱 컨텍스트 저장**: "app" 키로 요청 컨텍스트에 App 인스턴스 추가
//  5. **요청 업데이트**: 향상된 컨텍스트로 새 요청 생성
//  6. **처리 위임**: handler.ServeHTTP(w, r)를 호출하여 요청 처리
//
// 컨텍스트 저장:
//   - 요청 컨텍스트에 App 인스턴스 저장: context.WithValue(r.Context(), "app", a)
//   - 목적: 핸들러에서 템플릿 렌더링 및 앱 액세스 활성화
//   - 핸들러에서 검색: app := r.Context().Value("app").(*websvrutil.App)
//   - 요청 생명주기 전체에서 사용 가능(미들웨어, 핸들러, 헬퍼)
//   - 유용한 용도: 템플릿 렌더링, 앱 구성 액세스, 공유 리소스
//
// 스레드 안전성:
//   - 동시 요청 처리를 위해 완전히 스레드 안전합니다.
//   - 읽기 잠금(RLock)은 여러 동시 요청을 허용합니다.
//   - 핸들러 구축은 요청별로 격리됩니다(공유 상태 변경 없음).
//   - 높은 동시성 환경에 안전합니다(수천 개의 동시 요청).
//   - 쓰기 작업(Use, GET, POST 등)은 전체 쓰기 잠금이 필요하며 ServeHTTP를 차단합니다.
//
// 성능 특성:
//   - 잠금 오버헤드: RLock/RUnlock은 매우 빠름(경합 없음: ~20ns)
//   - 핸들러 구축: O(n), 여기서 n = 미들웨어 수(일반적으로 3-10)
//   - 컨텍스트 생성: 무시할 수 있는 오버헤드(~50ns)
//   - 전체 오버헤드: 일반적인 구성의 경우 < 1μs
//   - 확장성: 현대 하드웨어에서 초당 수백만 요청 처리
//
// 모범 사례:
//   - 테스트를 제외하고 ServeHTTP를 직접 호출하지 마세요.
//   - http.Server가 ServeHTTP 호출을 관리하도록 하세요.
//   - 패닉을 잡기 위해 복구 미들웨어를 사용하세요.
//   - ServeHTTP가 아닌 미들웨어에서 요청/응답을 로깅하세요.
//   - 핸들러 구축을 빠르게 유지하세요(미들웨어 수 최소화).
//
// 디버깅 팁:
//   - handler.ServeHTTP() 전후에 로깅을 추가하여 요청을 추적하세요.
//   - 요청 상관관계를 위해 요청 ID 미들웨어를 사용하세요.
//   - 잠금 경합을 모니터링하세요(거의 0에 가까워야 함).
//   - 성능 튜닝을 위해 ServeHTTP 실행 시간을 프로파일링하세요.
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

// TemplateEngine returns the template engine instance for direct template manipulation.
// This method provides access to the underlying TemplateEngine for advanced template operations
// beyond the convenience methods (LoadTemplate, LoadTemplates, etc.).
//
// TemplateEngine is useful when you need:
// - Direct access to template parsing and execution
// - Custom template functions before loading templates
// - Advanced template configuration not exposed by convenience methods
// - Integration with custom template loaders or pipelines
//
// Purpose:
//   - Provides read-only access to the TemplateEngine instance
//   - Enables advanced template manipulation and configuration
//   - Allows custom template loading strategies
//   - Supports template introspection and debugging
//
// Returns:
//   - *TemplateEngine: The template engine instance configured for this App.
//     Returns nil if template engine was not initialized (no TemplateDir option set).
//     Safe to call even when templates are not configured.
//
// Thread-Safety:
//   - Fully thread-safe with read lock protection.
//   - Acquires app.mu.RLock() before accessing templates field.
//   - Multiple goroutines can safely call TemplateEngine() concurrently.
//   - Returned TemplateEngine instance has its own internal thread-safety.
//
// Template Engine Initialization:
//   - Template engine is initialized when TemplateDir option is set during New():
//     app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//   - Returns nil if TemplateDir option was not provided
//   - Check for nil before using returned value:
//     if engine := app.TemplateEngine(); engine != nil { ... }
//
// Common Use Cases:
//   - Access template engine for custom function registration before loading
//   - Introspect loaded templates for debugging
//   - Implement custom template caching strategies
//   - Integrate with hot-reload or file watching systems
//   - Template testing and validation
//
// Example - Check Template Engine Availability:
//
//	engine := app.TemplateEngine()
//	if engine == nil {
//	    log.Fatal("Template engine not initialized - set TemplateDir option")
//	}
//
// Example - Access for Custom Configuration:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Get engine for advanced operations
//	engine := app.TemplateEngine()
//	if engine != nil {
//	    // Perform custom template operations
//	    // ...
//	}
//
// Example - Template Introspection:
//
//	engine := app.TemplateEngine()
//	if engine != nil {
//	    // Check if specific template is loaded
//	    // Inspect template metadata
//	    // Debug template parsing issues
//	}
//
// Relationship to Convenience Methods:
//   - LoadTemplate() calls engine.Load() internally
//   - LoadTemplates() calls engine.LoadGlob() internally
//   - ReloadTemplates() calls engine.Clear() + engine.LoadAll() internally
//   - AddTemplateFunc() calls engine.AddFunc() internally
//   - AddTemplateFuncs() calls engine.AddFuncs() internally
//
// Best Practices:
//   - Always check for nil before using returned engine
//   - Use convenience methods (LoadTemplate, etc.) for common operations
//   - Use TemplateEngine() for advanced scenarios requiring direct access
//   - Don't cache the engine reference; call TemplateEngine() when needed
//
// Performance Notes:
//   - Very lightweight operation (RLock + field access)
//   - No template loading or parsing occurs
//   - Safe to call frequently without performance concerns
//
// TemplateEngine은 직접 템플릿 조작을 위한 템플릿 엔진 인스턴스를 반환합니다.
// 이 메서드는 편의 메서드(LoadTemplate, LoadTemplates 등)를 넘어서는
// 고급 템플릿 작업을 위해 기본 TemplateEngine에 대한 액세스를 제공합니다.
//
// TemplateEngine은 다음이 필요할 때 유용합니다:
// - 템플릿 파싱 및 실행에 대한 직접 액세스
// - 템플릿 로드 전 커스텀 템플릿 함수
// - 편의 메서드로 노출되지 않은 고급 템플릿 구성
// - 커스텀 템플릿 로더 또는 파이프라인과의 통합
//
// 목적:
//   - TemplateEngine 인스턴스에 대한 읽기 전용 액세스 제공
//   - 고급 템플릿 조작 및 구성 활성화
//   - 커스텀 템플릿 로딩 전략 허용
//   - 템플릿 검사 및 디버깅 지원
//
// 반환값:
//   - *TemplateEngine: 이 App에 구성된 템플릿 엔진 인스턴스.
//     템플릿 엔진이 초기화되지 않은 경우(TemplateDir 옵션이 설정되지 않음) nil을 반환합니다.
//     템플릿이 구성되지 않았을 때도 안전하게 호출할 수 있습니다.
//
// 스레드 안전성:
//   - 읽기 잠금 보호로 완전히 스레드 안전합니다.
//   - templates 필드에 액세스하기 전에 app.mu.RLock()을 획득합니다.
//   - 여러 고루틴이 동시에 TemplateEngine()을 안전하게 호출할 수 있습니다.
//   - 반환된 TemplateEngine 인스턴스는 자체 내부 스레드 안전성을 가지고 있습니다.
//
// 템플릿 엔진 초기화:
//   - New() 중 TemplateDir 옵션이 설정될 때 템플릿 엔진이 초기화됩니다:
//     app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//   - TemplateDir 옵션이 제공되지 않은 경우 nil을 반환합니다
//   - 반환된 값을 사용하기 전에 nil을 확인하세요:
//     if engine := app.TemplateEngine(); engine != nil { ... }
//
// 일반적인 사용 사례:
//   - 로드 전 커스텀 함수 등록을 위한 템플릿 엔진 액세스
//   - 디버깅을 위한 로드된 템플릿 검사
//   - 커스텀 템플릿 캐싱 전략 구현
//   - 핫 리로드 또는 파일 감시 시스템과의 통합
//   - 템플릿 테스트 및 검증
//
// 모범 사례:
//   - 반환된 엔진을 사용하기 전에 항상 nil을 확인하세요
//   - 일반적인 작업에는 편의 메서드(LoadTemplate 등)를 사용하세요
//   - 직접 액세스가 필요한 고급 시나리오에만 TemplateEngine()을 사용하세요
//   - 엔진 참조를 캐시하지 마세요; 필요할 때 TemplateEngine()을 호출하세요
//
// 성능 참고사항:
//   - 매우 가벼운 작업(RLock + 필드 액세스)
//   - 템플릿 로딩 또는 파싱이 발생하지 않음
//   - 성능 문제 없이 자주 호출해도 안전
func (a *App) TemplateEngine() *TemplateEngine {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.templates
}

// LoadTemplate loads a single template file by name from the configured template directory.
// This is a convenience method for loading individual templates dynamically at runtime,
// useful for lazy loading, selective template loading, or dynamic template management.
//
// LoadTemplate loads the template immediately and makes it available for rendering.
// The template file must exist in the directory configured via WithTemplateDir option.
//
// Purpose:
//   - Loads a single named template file from the template directory
//   - Enables selective template loading (load only what you need)
//   - Supports dynamic template loading based on runtime conditions
//   - Useful for lazy loading templates to reduce startup time
//   - Allows hot-reloading individual templates during development
//
// Parameters:
//   - name: Template filename relative to the template directory.
//     Must be a simple filename or relative path within the template directory.
//     Examples: "index.html", "layout.html", "partials/header.html"
//     File must have extension matching configured template engine (.html, .tmpl, etc.)
//
// Returns:
//   - error: Returns error if template loading fails.
//     Error scenarios:
//   - Template engine not initialized (no TemplateDir option set)
//   - Template file not found in template directory
//   - Template syntax errors (invalid Go template syntax)
//   - File read permission denied
//   - Invalid template name (path traversal attempts, etc.)
//     Returns nil on successful template load.
//
// Template Engine Requirement:
//   - Template engine must be initialized via WithTemplateDir option:
//     app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//   - Returns error if template engine is nil (not initialized)
//   - Check error message: "template engine not initialized (set TemplateDir option)"
//
// Thread-Safety:
//   - Thread-safe with read lock protection (app.mu.RLock).
//   - Multiple goroutines can call LoadTemplate() concurrently.
//   - Safe to call during request handling (though not recommended for performance).
//   - Template engine internally manages template storage with appropriate locking.
//
// Template Resolution:
//   - Template name is resolved relative to TemplateDir:
//     TemplateDir="/app/views", name="index.html" → loads /app/views/index.html
//   - Supports subdirectories: name="partials/header.html" → /app/views/partials/header.html
//   - Does not support absolute paths or parent directory references (../)
//   - Path traversal attempts are rejected for security
//
// Behavior:
//   - Loads template file from disk
//   - Parses template with Go html/template package
//   - Stores parsed template in engine's template map
//   - Replaces existing template if name already loaded (reload capability)
//   - Template becomes immediately available for rendering after successful load
//
// Common Use Cases:
//   - Lazy loading: Load templates on-demand instead of at startup
//   - Dynamic loading: Load templates based on user configuration or feature flags
//   - Hot reload: Reload individual templates during development without server restart
//   - Selective loading: Load only templates needed for current request pattern
//   - Template testing: Load specific templates for unit testing
//
// Example - Basic Template Loading:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Load single template
//	if err := app.LoadTemplate("index.html"); err != nil {
//	    log.Fatalf("Failed to load template: %v", err)
//	}
//
//	// Template is now available for rendering
//	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
//	    // Render the loaded template
//	    // ...
//	})
//
// Example - Lazy Loading Strategy:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Don't load all templates at startup
//	app.GET("/dashboard", func(w http.ResponseWriter, r *http.Request) {
//	    // Load dashboard template on first access
//	    if err := app.LoadTemplate("dashboard.html"); err != nil {
//	        http.Error(w, "Template load error", http.StatusInternalServerError)
//	        return
//	    }
//	    // Render dashboard
//	})
//
// Example - Development Hot Reload:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// In development mode, reload template on each request
//	isDev := os.Getenv("ENV") == "development"
//
//	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
//	    if isDev {
//	        // Reload template to pick up changes
//	        if err := app.LoadTemplate("index.html"); err != nil {
//	            log.Printf("Template reload failed: %v", err)
//	        }
//	    }
//	    // Render template
//	})
//
// Example - Error Handling:
//
//	if err := app.LoadTemplate("missing.html"); err != nil {
//	    if strings.Contains(err.Error(), "not initialized") {
//	        log.Fatal("Template engine not configured - use WithTemplateDir")
//	    } else if strings.Contains(err.Error(), "not found") {
//	        log.Printf("Template file missing: %v", err)
//	    } else {
//	        log.Printf("Template syntax error: %v", err)
//	    }
//	}
//
// Example - Load with Subdirectories:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Load template from subdirectory
//	if err := app.LoadTemplate("partials/header.html"); err != nil {
//	    log.Fatal(err)
//	}
//
//	if err := app.LoadTemplate("layouts/main.html"); err != nil {
//	    log.Fatal(err)
//	}
//
// Performance Considerations:
//   - Template loading involves disk I/O (relatively slow)
//   - Template parsing has CPU overhead (template compilation)
//   - Prefer loading templates at startup rather than per-request
//   - Use LoadTemplates() for bulk loading (more efficient than multiple LoadTemplate calls)
//   - Cache loaded templates; avoid reloading unless necessary
//   - Hot reload should be development-only (performance impact)
//
// Comparison with LoadTemplates:
//   - LoadTemplate: Single file, explicit name
//   - LoadTemplates: Multiple files, glob pattern (e.g., "*.html")
//   - Use LoadTemplate for: Individual files, dynamic loading, hot reload
//   - Use LoadTemplates for: Bulk loading, startup initialization, wildcard patterns
//
// Best Practices:
//   - Load all templates at startup for production (use LoadTemplates("*.html"))
//   - Use LoadTemplate for dynamic/lazy loading scenarios only
//   - Handle errors appropriately (log, return HTTP 500, etc.)
//   - Avoid loading templates during request handling in production
//   - Use hot reload only in development mode
//   - Validate template syntax before deployment
//
// Debugging Tips:
//   - Check error message for specific failure reason
//   - Verify template file exists in template directory
//   - Ensure TemplateDir is correctly configured
//   - Test template syntax with Go template playground
//   - Enable template engine debug logging if available
//
// Security Considerations:
//   - Template name should not be user-controlled (path traversal risk)
//   - Validate template names if accepting from external sources
//   - Template directory should not be web-accessible
//   - Ensure template files have appropriate file permissions
//
// LoadTemplate은 구성된 템플릿 디렉토리에서 이름으로 단일 템플릿 파일을 로드합니다.
// 이는 런타임에 개별 템플릿을 동적으로 로드하기 위한 편의 메서드이며,
// 지연 로딩, 선택적 템플릿 로딩 또는 동적 템플릿 관리에 유용합니다.
//
// LoadTemplate은 템플릿을 즉시 로드하고 렌더링에 사용할 수 있도록 만듭니다.
// 템플릿 파일은 WithTemplateDir 옵션을 통해 구성된 디렉토리에 존재해야 합니다.
//
// 목적:
//   - 템플릿 디렉토리에서 이름이 지정된 단일 템플릿 파일 로드
//   - 선택적 템플릿 로딩 활성화(필요한 것만 로드)
//   - 런타임 조건에 따른 동적 템플릿 로딩 지원
//   - 시작 시간을 줄이기 위한 지연 로딩 템플릿에 유용
//   - 개발 중 개별 템플릿의 핫 리로딩 허용
//
// 매개변수:
//   - name: 템플릿 디렉토리에 상대적인 템플릿 파일 이름.
//     템플릿 디렉토리 내의 단순 파일 이름 또는 상대 경로여야 합니다.
//     예제: "index.html", "layout.html", "partials/header.html"
//     파일은 구성된 템플릿 엔진과 일치하는 확장자를 가져야 합니다(.html, .tmpl 등).
//
// 반환값:
//   - error: 템플릿 로딩이 실패하면 오류를 반환합니다.
//     오류 시나리오:
//   - 템플릿 엔진이 초기화되지 않음(TemplateDir 옵션이 설정되지 않음)
//   - 템플릿 디렉토리에서 템플릿 파일을 찾을 수 없음
//   - 템플릿 구문 오류(잘못된 Go 템플릿 구문)
//   - 파일 읽기 권한 거부
//   - 잘못된 템플릿 이름(경로 순회 시도 등)
//     템플릿 로드에 성공하면 nil을 반환합니다.
//
// 템플릿 엔진 요구사항:
//   - WithTemplateDir 옵션을 통해 템플릿 엔진이 초기화되어야 합니다:
//     app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//   - 템플릿 엔진이 nil인 경우(초기화되지 않음) 오류를 반환합니다
//   - 오류 메시지 확인: "template engine not initialized (set TemplateDir option)"
//
// 스레드 안전성:
//   - 읽기 잠금 보호(app.mu.RLock)로 스레드 안전합니다.
//   - 여러 고루틴이 동시에 LoadTemplate()을 호출할 수 있습니다.
//   - 요청 처리 중 호출해도 안전합니다(성능상 권장하지는 않음).
//   - 템플릿 엔진은 적절한 잠금으로 템플릿 저장소를 내부적으로 관리합니다.
//
// 템플릿 해결:
//   - 템플릿 이름은 TemplateDir에 상대적으로 해결됩니다:
//     TemplateDir="/app/views", name="index.html" → /app/views/index.html 로드
//   - 하위 디렉토리 지원: name="partials/header.html" → /app/views/partials/header.html
//   - 절대 경로 또는 상위 디렉토리 참조(../)를 지원하지 않음
//   - 경로 순회 시도는 보안을 위해 거부됨
//
// 동작 방식:
//   - 디스크에서 템플릿 파일 로드
//   - Go html/template 패키지로 템플릿 파싱
//   - 엔진의 템플릿 맵에 파싱된 템플릿 저장
//   - 이름이 이미 로드된 경우 기존 템플릿 교체(리로드 기능)
//   - 성공적으로 로드한 후 템플릿을 즉시 렌더링에 사용할 수 있음
//
// 모범 사례:
//   - 프로덕션에서는 시작 시 모든 템플릿을 로드하세요(LoadTemplates("*.html") 사용)
//   - 동적/지연 로딩 시나리오에만 LoadTemplate 사용
//   - 오류를 적절하게 처리하세요(로그, HTTP 500 반환 등)
//   - 프로덕션에서 요청 처리 중 템플릿 로딩을 피하세요
//   - 개발 모드에서만 핫 리로드 사용
//   - 배포 전에 템플릿 구문 검증
//
// 보안 고려사항:
//   - 템플릿 이름은 사용자가 제어해서는 안 됩니다(경로 순회 위험)
//   - 외부 소스에서 수락하는 경우 템플릿 이름 검증
//   - 템플릿 디렉토리는 웹에서 액세스할 수 없어야 함
//   - 템플릿 파일에 적절한 파일 권한이 있는지 확인
func (a *App) LoadTemplate(name string) error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates == nil {
		return fmt.Errorf("template engine not initialized (set TemplateDir option)")
	}

	return a.templates.Load(name)
}

// LoadTemplates loads all templates matching the specified glob pattern from the template directory.
// This is the **recommended method for bulk template loading** at application startup.
// It efficiently loads multiple template files in a single operation using glob pattern matching.
//
// LoadTemplates is ideal for:
// - Loading all templates at application startup
// - Bulk loading templates from specific subdirectories
// - Loading templates by extension or naming convention
// - Initializing template cache before request handling begins
//
// Purpose:
//   - Loads multiple template files matching a glob pattern
//   - Enables efficient bulk template loading (better than multiple LoadTemplate calls)
//   - Supports flexible pattern matching for selective loading
//   - Ideal for startup initialization and batch operations
//   - Reduces startup time compared to individual file loading
//
// Parameters:
//
//   - pattern: Glob pattern for matching template files within the template directory.
//     Glob patterns support standard wildcards:
//
//   - "*" - Matches any sequence of characters (except directory separator)
//
//   - "?" - Matches any single character
//
//   - "**" - Matches directories recursively (implementation-dependent)
//
//   - "[abc]" - Matches any character in the bracket set
//
//   - "{a,b}" - Matches either alternative (brace expansion)
//
//     Common patterns:
//
//   - "*.html" - All .html files in template directory root
//
//   - "**/*.html" - All .html files in any subdirectory (recursive)
//
//   - "layouts/*.tmpl" - All .tmpl files in layouts subdirectory
//
//   - "partial*.html" - Files starting with "partial" and ending with .html
//
//   - "*.{html,tmpl}" - Files with .html or .tmpl extension
//
// Returns:
//   - error: Returns error if template loading fails.
//     Error scenarios:
//   - Template engine not initialized (no TemplateDir option set)
//   - Invalid glob pattern syntax
//   - No files match the pattern (may or may not be an error depending on implementation)
//   - Template syntax errors in any matched file
//   - File read permission denied
//   - Disk I/O errors
//     Returns nil when all matching templates load successfully.
//
// Template Engine Requirement:
//   - Template engine must be initialized via WithTemplateDir option:
//     app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//   - Returns error if template engine is nil: "template engine not initialized (set TemplateDir option)"
//
// Thread-Safety:
//   - Thread-safe with read lock protection (app.mu.RLock).
//   - Safe for concurrent calls from multiple goroutines.
//   - Not recommended during request handling (performance impact).
//   - Best practice: Load templates during application initialization phase.
//
// Pattern Matching Behavior:
//   - Pattern is evaluated relative to TemplateDir
//   - Recursive patterns (like "**/*.html") scan subdirectories
//   - Pattern matching is case-sensitive on Unix/Linux, case-insensitive on Windows
//   - Hidden files (starting with .) are typically excluded unless explicitly matched
//   - Symlinks behavior depends on implementation (may or may not be followed)
//
// Loading Behavior:
//   - All matching files are loaded and parsed
//   - Each template is stored with its filename (or relative path) as the key
//   - Existing templates with same names are replaced (reload capability)
//   - Loading stops on first error (partial loading may occur)
//   - Templates become immediately available for rendering after successful load
//
// Common Use Cases:
//   - Startup initialization: Load all templates before serving requests
//   - Selective loading: Load only specific subdirectories or file types
//   - Template organization: Load templates by category (layouts, partials, pages)
//   - Development: Reload all templates after changes
//   - Testing: Load test fixture templates
//
// Example - Load All HTML Templates:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Load all .html files at startup
//	if err := app.LoadTemplates("*.html"); err != nil {
//	    log.Fatalf("Failed to load templates: %v", err)
//	}
//
//	// All templates now available for rendering
//	app.Run(":8080")
//
// Example - Load Templates Recursively:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Load all .html files from all subdirectories
//	if err := app.LoadTemplates("**/*.html"); err != nil {
//	    log.Fatalf("Failed to load templates: %v", err)
//	}
//
// Example - Load Specific Subdirectory:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Load only layout templates
//	if err := app.LoadTemplates("layouts/*.html"); err != nil {
//	    log.Fatalf("Failed to load layouts: %v", err)
//	}
//
//	// Load only partial templates
//	if err := app.LoadTemplates("partials/*.html"); err != nil {
//	    log.Fatalf("Failed to load partials: %v", err)
//	}
//
// Example - Multiple Extensions:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Load both .html and .tmpl files
//	patterns := []string{"*.html", "*.tmpl"}
//	for _, pattern := range patterns {
//	    if err := app.LoadTemplates(pattern); err != nil {
//	        log.Fatalf("Failed to load %s: %v", pattern, err)
//	    }
//	}
//
// Example - Staged Loading:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Load critical templates first
//	if err := app.LoadTemplates("layouts/*.html"); err != nil {
//	    log.Fatal("Critical layouts missing:", err)
//	}
//
//	// Load optional templates (errors non-fatal)
//	if err := app.LoadTemplates("optional/*.html"); err != nil {
//	    log.Println("Optional templates not loaded:", err)
//	}
//
//	// Load remaining templates
//	if err := app.LoadTemplates("*.html"); err != nil {
//	    log.Fatal("Template loading failed:", err)
//	}
//
// Example - Development Hot Reload:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Initial load
//	app.LoadTemplates("*.html")
//
//	// In development mode, watch for changes and reload
//	if isDevelopment {
//	    // (Pseudo-code - actual file watching implementation needed)
//	    watchFiles("./views", func() {
//	        log.Println("Templates changed, reloading...")
//	        if err := app.ReloadTemplates(); err != nil {
//	            log.Printf("Reload failed: %v", err)
//	        }
//	    })
//	}
//
// Performance Considerations:
//   - Bulk loading is more efficient than multiple LoadTemplate() calls
//   - Template parsing has CPU overhead (proportional to template complexity)
//   - Disk I/O is the primary bottleneck (especially for many small files)
//   - Loading time increases with number of files and template complexity
//   - Typical startup overhead: 10-100ms for small apps, up to seconds for large apps
//   - Consider lazy loading for applications with hundreds of templates
//
// Comparison with LoadTemplate:
//   - LoadTemplates: Multiple files via glob pattern, efficient bulk loading
//   - LoadTemplate: Single file by exact name, selective loading
//   - Use LoadTemplates for: Startup initialization, bulk operations, wildcard patterns
//   - Use LoadTemplate for: Individual files, runtime loading, specific templates
//
// Best Practices:
//   - Load all templates at startup using LoadTemplates("*.html") or similar
//   - Organize templates in subdirectories for maintainability
//   - Use consistent file extensions (.html, .tmpl) for easier glob patterns
//   - Handle errors gracefully with proper logging
//   - Validate templates in CI/CD pipeline before deployment
//   - Use recursive patterns carefully (can be slow with deep directory trees)
//   - Reload templates only in development mode (not production)
//
// Debugging Tips:
//   - Log the pattern used to verify it matches intended files
//   - Check template directory structure matches expected pattern
//   - Test glob patterns with shell commands (ls views/*.html)
//   - Enable verbose logging to see which files are loaded
//   - Verify file permissions allow reading template files
//
// Security Considerations:
//   - Pattern should not be user-controlled (glob injection risk)
//   - Validate patterns if accepting from external configuration
//   - Template directory should not be web-accessible
//   - Ensure loaded templates don't execute untrusted code
//
// LoadTemplates는 템플릿 디렉토리에서 지정된 글로브 패턴과 일치하는 모든 템플릿을 로드합니다.
// 이는 애플리케이션 시작 시 **대량 템플릿 로딩을 위한 권장 메서드**입니다.
// 글로브 패턴 매칭을 사용하여 단일 작업으로 여러 템플릿 파일을 효율적으로 로드합니다.
//
// LoadTemplates는 다음에 이상적입니다:
// - 애플리케이션 시작 시 모든 템플릿 로드
// - 특정 하위 디렉토리에서 대량 템플릿 로드
// - 확장자 또는 명명 규칙별 템플릿 로드
// - 요청 처리가 시작되기 전에 템플릿 캐시 초기화
//
// 목적:
//   - 글로브 패턴과 일치하는 여러 템플릿 파일 로드
//   - 효율적인 대량 템플릿 로딩 활성화(여러 LoadTemplate 호출보다 우수)
//   - 선택적 로딩을 위한 유연한 패턴 매칭 지원
//   - 시작 초기화 및 배치 작업에 이상적
//   - 개별 파일 로딩에 비해 시작 시간 단축
//
// 매개변수:
//
//   - pattern: 템플릿 디렉토리 내에서 템플릿 파일을 매칭하기 위한 글로브 패턴.
//     글로브 패턴은 표준 와일드카드를 지원합니다:
//
//   - "*" - 모든 문자 시퀀스와 일치(디렉토리 구분자 제외)
//
//   - "?" - 모든 단일 문자와 일치
//
//   - "**" - 디렉토리를 재귀적으로 일치(구현에 따라 다름)
//
//   - "[abc]" - 브래킷 세트의 모든 문자와 일치
//
//   - "{a,b}" - 대안 중 하나와 일치(브레이스 확장)
//
//     일반적인 패턴:
//
//   - "*.html" - 템플릿 디렉토리 루트의 모든 .html 파일
//
//   - "**/*.html" - 모든 하위 디렉토리의 모든 .html 파일(재귀)
//
//   - "layouts/*.tmpl" - layouts 하위 디렉토리의 모든 .tmpl 파일
//
//   - "partial*.html" - "partial"로 시작하고 .html로 끝나는 파일
//
//   - "*.{html,tmpl}" - .html 또는 .tmpl 확장자를 가진 파일
//
// 반환값:
//   - error: 템플릿 로딩이 실패하면 오류를 반환합니다.
//     오류 시나리오:
//   - 템플릿 엔진이 초기화되지 않음(TemplateDir 옵션이 설정되지 않음)
//   - 잘못된 글로브 패턴 구문
//   - 패턴과 일치하는 파일이 없음(구현에 따라 오류일 수도 아닐 수도 있음)
//   - 일치하는 파일의 템플릿 구문 오류
//   - 파일 읽기 권한 거부
//   - 디스크 I/O 오류
//     모든 일치하는 템플릿이 성공적으로 로드되면 nil을 반환합니다.
//
// 스레드 안전성:
//   - 읽기 잠금 보호(app.mu.RLock)로 스레드 안전합니다.
//   - 여러 고루틴에서 동시 호출에 안전합니다.
//   - 요청 처리 중에는 권장하지 않습니다(성능 영향).
//   - 모범 사례: 애플리케이션 초기화 단계에서 템플릿을 로드하세요.
//
// 모범 사례:
//   - 시작 시 LoadTemplates("*.html") 등을 사용하여 모든 템플릿을 로드하세요
//   - 유지 관리성을 위해 하위 디렉토리에 템플릿을 구성하세요
//   - 쉬운 글로브 패턴을 위해 일관된 파일 확장자(.html, .tmpl)를 사용하세요
//   - 적절한 로깅으로 오류를 우아하게 처리하세요
//   - 배포 전에 CI/CD 파이프라인에서 템플릿을 검증하세요
//   - 재귀 패턴은 신중하게 사용하세요(깊은 디렉토리 트리에서 느릴 수 있음)
//   - 개발 모드에서만 템플릿을 리로드하세요(프로덕션이 아님)
//
// 보안 고려사항:
//   - 패턴은 사용자가 제어해서는 안 됩니다(글로브 주입 위험)
//   - 외부 구성에서 수락하는 경우 패턴을 검증하세요
//   - 템플릿 디렉토리는 웹에서 액세스할 수 없어야 합니다
//   - 로드된 템플릿이 신뢰할 수 없는 코드를 실행하지 않도록 하세요
func (a *App) LoadTemplates(pattern string) error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates == nil {
		return fmt.Errorf("template engine not initialized (set TemplateDir option)")
	}

	return a.templates.LoadGlob(pattern)
}

// ReloadTemplates clears all currently loaded templates and reloads them from the template directory.
// This is a **development and debugging tool** for refreshing templates without server restart.
// Not recommended for production use due to performance impact and potential race conditions.
//
// ReloadTemplates is useful for:
// - Development: See template changes immediately without restarting server
// - Debugging: Test template modifications in running application
// - Hot reload: Implement file watching with automatic template refresh
// - Testing: Reset template state between test cases
//
// Purpose:
//   - Clears all currently loaded templates from memory
//   - Reloads all templates from template directory (same as initial load)
//   - Enables template hot-reloading during development
//   - Provides clean slate for template testing and debugging
//   - Allows dynamic template updates without application restart
//
// Returns:
//   - error: Returns error if template reloading fails.
//     Error scenarios:
//   - Template engine not initialized (no TemplateDir option set)
//   - Template directory not accessible (permissions, missing directory)
//   - Template syntax errors in any template file
//   - Disk I/O errors during file reading
//   - Partial reload failure (some templates cleared, reload incomplete)
//     Returns nil when all templates successfully reloaded.
//
// Template Engine Requirement:
//   - Template engine must be initialized via WithTemplateDir option:
//     app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//   - Returns error if template engine is nil: "template engine not initialized (set TemplateDir option)"
//
// Thread-Safety:
//   - Thread-safe with read lock protection (app.mu.RLock).
//   - Safe for concurrent calls from multiple goroutines.
//   - **WARNING**: Race condition possible during reload:
//   - Templates are cleared first (all templates temporarily unavailable)
//   - Templates are then reloaded (new templates become available)
//   - Requests during this window may fail if templates are missing
//   - Duration: Typically 10-100ms depending on template count and complexity
//   - Not recommended for production due to race condition risk.
//
// Reload Behavior:
//
//  1. **Clear Phase**: All currently loaded templates are removed from memory
//
//  2. **Load Phase**: Template engine reloads all templates from directory
//
//  3. **Result**: Fresh template state matching current filesystem contents
//
//     Template discovery:
//     - Reloads all template files in TemplateDir and subdirectories
//     - Uses same loading logic as initial startup
//     - Discovers new files added since startup
//     - Removes deleted files from template cache
//     - Updates modified files with latest content
//
// Performance Impact:
//   - Template clearing: Very fast (memory deallocation)
//   - Template reloading: Disk I/O + parsing overhead
//   - Total time: Typically 10-100ms for small apps, up to seconds for large apps
//   - Blocks template rendering during reload window
//   - Not suitable for high-traffic production environments
//
// Common Use Cases:
//   - Development workflow: Reload templates after editing without restart
//   - Live preview: Watch template files and reload on changes
//   - Template debugging: Test template modifications iteratively
//   - A/B testing: Switch template implementations dynamically
//   - Testing: Reset templates between test cases
//
// Example - Manual Reload in Development:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//	app.LoadTemplates("*.html")
//
//	// Add reload endpoint for development
//	if os.Getenv("ENV") == "development" {
//	    app.GET("/admin/reload-templates", func(w http.ResponseWriter, r *http.Request) {
//	        if err := app.ReloadTemplates(); err != nil {
//	            http.Error(w, "Reload failed: "+err.Error(), 500)
//	            return
//	        }
//	        w.Write([]byte("Templates reloaded successfully"))
//	    })
//	}
//
// Example - Automatic Reload with File Watching:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//	app.LoadTemplates("*.html")
//
//	// Watch template directory for changes (pseudo-code)
//	if os.Getenv("ENV") == "development" {
//	    go func() {
//	        watcher := newFileWatcher("./views")
//	        for change := range watcher.Events() {
//	            log.Printf("Template changed: %s, reloading...", change.Path)
//	            if err := app.ReloadTemplates(); err != nil {
//	                log.Printf("Reload failed: %v", err)
//	            } else {
//	                log.Println("Templates reloaded successfully")
//	            }
//	        }
//	    }()
//	}
//
//	app.Run(":8080")
//
// Example - Reload on HTTP Request (Dev Mode):
//
//	isDev := os.Getenv("ENV") == "development"
//
//	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
//	    // Reload templates on each request in development
//	    if isDev {
//	        if err := app.ReloadTemplates(); err != nil {
//	            log.Printf("Template reload error: %v", err)
//	        }
//	    }
//
//	    // Render template (now using latest version)
//	    // ...
//	})
//
// Example - Conditional Reload (URL Parameter):
//
//	app.GET("/page", func(w http.ResponseWriter, r *http.Request) {
//	    // Allow reload via query parameter in development
//	    if os.Getenv("ENV") == "development" && r.URL.Query().Get("reload") == "1" {
//	        log.Println("Manual reload requested")
//	        if err := app.ReloadTemplates(); err != nil {
//	            log.Printf("Reload failed: %v", err)
//	        }
//	    }
//
//	    // Render page
//	    // ...
//	})
//
//	// Usage: http://localhost:8080/page?reload=1
//
// Example - Testing with Template Reload:
//
//	func TestTemplateRendering(t *testing.T) {
//	    app := websvrutil.New(websvrutil.WithTemplateDir("./test-templates"))
//
//	    // Load initial templates
//	    app.LoadTemplates("*.html")
//
//	    // Test with original templates
//	    // ... run tests ...
//
//	    // Modify template files
//	    // ... write new template content ...
//
//	    // Reload to pick up changes
//	    if err := app.ReloadTemplates(); err != nil {
//	        t.Fatalf("Reload failed: %v", err)
//	    }
//
//	    // Test with modified templates
//	    // ... run more tests ...
//	}
//
// File Watching Integration Example (using fsnotify):
//
//	import "github.com/fsnotify/fsnotify"
//
//	func setupTemplateWatcher(app *websvrutil.App, dir string) error {
//	    watcher, err := fsnotify.NewWatcher()
//	    if err != nil {
//	        return err
//	    }
//
//	    go func() {
//	        defer watcher.Close()
//	        for {
//	            select {
//	            case event := <-watcher.Events:
//	                if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) != 0 {
//	                    log.Printf("Template change detected: %s", event.Name)
//	                    time.Sleep(100 * time.Millisecond) // Debounce
//	                    if err := app.ReloadTemplates(); err != nil {
//	                        log.Printf("Reload error: %v", err)
//	                    } else {
//	                        log.Println("Templates reloaded")
//	                    }
//	                }
//	            case err := <-watcher.Errors:
//	                log.Printf("Watcher error: %v", err)
//	            }
//	        }
//	    }()
//
//	    return watcher.Add(dir)
//	}
//
// Race Condition Scenario:
//   - Time 0ms: Request A starts, needs template "index.html"
//   - Time 10ms: ReloadTemplates() called, clears all templates
//   - Time 15ms: Request A tries to render "index.html" → ERROR (template missing)
//   - Time 50ms: ReloadTemplates() completes, templates available again
//   - Time 60ms: Request B renders "index.html" → SUCCESS
//
// Mitigation Strategies:
//   - Use only in development mode (ENV check)
//   - Avoid reload during high traffic periods
//   - Implement reload cooldown (don't reload more than once per second)
//   - Add reload mutex to serialize reload operations
//   - Consider blue-green deployment for production template updates
//
// Best Practices:
//   - **Development only**: Never use in production environments
//   - Protect reload endpoints with authentication (if exposed)
//   - Log reload operations for debugging
//   - Implement reload cooldown to prevent excessive disk I/O
//   - Handle reload errors gracefully (don't crash server)
//   - Test thoroughly before relying on hot-reload in development
//   - Consider using a dedicated file watcher library (fsnotify, etc.)
//
// Debugging Tips:
//   - Log template count before and after reload
//   - Monitor reload duration to detect performance issues
//   - Check for template syntax errors in logs
//   - Verify file permissions on template directory
//   - Test reload with simple template first
//
// Production Alternatives:
//   - Blue-green deployment: Deploy new version with updated templates
//   - Rolling restart: Gradually restart application instances
//   - Immutable deployments: Build container images with templates included
//   - Configuration management: Use config management tools for updates
//
// Security Considerations:
//   - Protect reload endpoints with authentication/authorization
//   - Don't expose reload functionality in production
//   - Validate template directory permissions
//   - Log reload operations for security auditing
//   - Rate-limit reload operations to prevent abuse
//
// ReloadTemplates는 현재 로드된 모든 템플릿을 지우고 템플릿 디렉토리에서 다시 로드합니다.
// 이는 서버 재시작 없이 템플릿을 새로 고치기 위한 **개발 및 디버깅 도구**입니다.
// 성능 영향과 잠재적인 경쟁 조건으로 인해 프로덕션 사용은 권장되지 않습니다.
//
// ReloadTemplates는 다음에 유용합니다:
// - 개발: 서버를 재시작하지 않고 템플릿 변경 사항을 즉시 확인
// - 디버깅: 실행 중인 애플리케이션에서 템플릿 수정 테스트
// - 핫 리로드: 자동 템플릿 새로 고침과 함께 파일 감시 구현
// - 테스팅: 테스트 케이스 간 템플릿 상태 재설정
//
// 목적:
//   - 메모리에서 현재 로드된 모든 템플릿 지우기
//   - 템플릿 디렉토리에서 모든 템플릿 다시 로드(초기 로드와 동일)
//   - 개발 중 템플릿 핫 리로딩 활성화
//   - 템플릿 테스트 및 디버깅을 위한 깨끗한 상태 제공
//   - 애플리케이션 재시작 없이 동적 템플릿 업데이트 허용
//
// 반환값:
//   - error: 템플릿 리로딩이 실패하면 오류를 반환합니다.
//     오류 시나리오:
//   - 템플릿 엔진이 초기화되지 않음(TemplateDir 옵션이 설정되지 않음)
//   - 템플릿 디렉토리에 액세스할 수 없음(권한, 누락된 디렉토리)
//   - 템플릿 파일의 템플릿 구문 오류
//   - 파일 읽기 중 디스크 I/O 오류
//   - 부분 리로드 실패(일부 템플릿 지워짐, 리로드 불완전)
//     모든 템플릿이 성공적으로 리로드되면 nil을 반환합니다.
//
// 스레드 안전성:
//   - 읽기 잠금 보호(app.mu.RLock)로 스레드 안전합니다.
//   - 여러 고루틴에서 동시 호출에 안전합니다.
//   - **경고**: 리로드 중 경쟁 조건 가능:
//   - 템플릿이 먼저 지워짐(모든 템플릿이 일시적으로 사용 불가)
//   - 그런 다음 템플릿이 리로드됨(새 템플릿 사용 가능)
//   - 이 기간 동안의 요청은 템플릿이 누락되면 실패할 수 있음
//   - 기간: 템플릿 수와 복잡성에 따라 일반적으로 10-100ms
//   - 경쟁 조건 위험으로 인해 프로덕션에는 권장되지 않습니다.
//
// 모범 사례:
//   - **개발 전용**: 프로덕션 환경에서는 절대 사용하지 마세요
//   - 인증으로 리로드 엔드포인트를 보호하세요(노출된 경우)
//   - 디버깅을 위해 리로드 작업을 로깅하세요
//   - 과도한 디스크 I/O를 방지하기 위해 리로드 쿨다운을 구현하세요
//   - 리로드 오류를 우아하게 처리하세요(서버를 충돌시키지 마세요)
//   - 개발에서 핫 리로드에 의존하기 전에 철저히 테스트하세요
//   - 전용 파일 감시 라이브러리(fsnotify 등) 사용을 고려하세요
//
// 보안 고려사항:
//   - 인증/권한으로 리로드 엔드포인트를 보호하세요
//   - 프로덕션에서 리로드 기능을 노출하지 마세요
//   - 템플릿 디렉토리 권한을 검증하세요
//   - 보안 감사를 위해 리로드 작업을 로깅하세요
//   - 남용을 방지하기 위해 리로드 작업의 속도를 제한하세요
func (a *App) ReloadTemplates() error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates == nil {
		return fmt.Errorf("template engine not initialized (set TemplateDir option)")
	}

	a.templates.Clear()
	return a.templates.LoadAll()
}

// AddTemplateFunc registers a custom function that can be used in templates.
// This allows extending template functionality beyond Go's built-in template functions
// by adding domain-specific helpers, formatters, and utilities.
//
// AddTemplateFunc enables templates to:
// - Format data (dates, numbers, currency, etc.)
// - Perform calculations and transformations
// - Call application logic from templates
// - Access helper utilities (string manipulation, etc.)
// - Implement custom template DSL (domain-specific language)
//
// Purpose:
//   - Registers a single custom function for use in templates
//   - Extends template capabilities with application-specific logic
//   - Enables cleaner templates by moving complex logic to Go functions
//   - Provides reusable template utilities across all templates
//   - Allows templates to access application context and helpers
//
// Parameters:
//
//   - name: Function name as it will be called in templates.
//     Must be a valid Go identifier (alphanumeric + underscore, start with letter/underscore).
//     Case-sensitive: "formatDate" and "FormatDate" are different functions.
//     Cannot override built-in template functions (and, or, not, etc.) without special handling.
//
//   - fn: Function to execute when called from template.
//     Must be a Go function with specific signature requirements:
//
//   - Returns 1 or 2 values
//
//   - If 2 values: second must be error (for error handling in templates)
//
//   - Parameters: Can accept any number of parameters of any type
//
//   - Parameter types: Template engine will attempt type conversion
//
//   - Return types: Must be types that can be rendered in templates
//
//     Valid function signatures:
//
//   - func() string
//
//   - func(string) string
//
//   - func(int, int) int
//
//   - func(string) (string, error)
//
//   - func(interface{}) string
//
//   - func(...interface{}) string (variadic)
//
// Function Signature Requirements:
//
//	The function must follow Go template function rules:
//	- Can have any number of parameters (0+)
//	- Can return 1 value: the result
//	- Can return 2 values: (result, error)
//	- If error is returned and non-nil, template execution fails
//	- Parameter and return types should be serializable/printable
//
// Thread-Safety:
//   - Thread-safe with read lock protection (app.mu.RLock).
//   - Safe to call during initialization or runtime.
//   - Function should be registered BEFORE loading templates.
//   - Multiple goroutines can safely call AddTemplateFunc concurrently.
//   - Registered function itself should be thread-safe if called concurrently.
//
// Registration Timing:
//   - **Critical**: Must be called BEFORE loading templates.
//   - Functions registered after template loading won't be available in those templates.
//   - Recommended: Register all functions immediately after app creation.
//   - If reloading templates, functions persist (don't need re-registration).
//
// Template Engine Requirement:
//   - Template engine must be initialized (WithTemplateDir option set).
//   - Silently does nothing if template engine is nil.
//   - No error returned; check TemplateEngine() != nil if verification needed.
//
// Common Use Cases:
//   - Date/time formatting: Format timestamps in human-readable formats
//   - String manipulation: Uppercase, lowercase, truncate, etc.
//   - Number formatting: Currency, percentages, thousands separators
//   - URL building: Generate URLs from route names and parameters
//   - Asset helpers: Generate asset URLs with cache-busting
//   - Authorization checks: Check user permissions in templates
//   - i18n/l10n: Translate strings, format locale-specific data
//
// Example - Basic String Formatting:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Register uppercase function
//	app.AddTemplateFunc("upper", strings.ToUpper)
//
//	// Register lowercase function
//	app.AddTemplateFunc("lower", strings.ToLower)
//
//	// Load templates (functions now available)
//	app.LoadTemplates("*.html")
//
//	// In template:
//	// {{ .Name | upper }}  → "JOHN"
//	// {{ .Name | lower }}  → "john"
//
// Example - Date Formatting:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Register date formatter
//	app.AddTemplateFunc("formatDate", func(t time.Time) string {
//	    return t.Format("2006-01-02")
//	})
//
//	// Register relative time formatter
//	app.AddTemplateFunc("timeAgo", func(t time.Time) string {
//	    d := time.Since(t)
//	    if d < time.Hour {
//	        return fmt.Sprintf("%d minutes ago", int(d.Minutes()))
//	    }
//	    return fmt.Sprintf("%d hours ago", int(d.Hours()))
//	})
//
//	app.LoadTemplates("*.html")
//
//	// In template:
//	// {{ .CreatedAt | formatDate }}  → "2024-01-15"
//	// {{ .UpdatedAt | timeAgo }}     → "5 minutes ago"
//
// Example - Number Formatting:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Register currency formatter
//	app.AddTemplateFunc("currency", func(amount float64) string {
//	    return fmt.Sprintf("$%.2f", amount)
//	})
//
//	// Register percentage formatter
//	app.AddTemplateFunc("percent", func(value float64) string {
//	    return fmt.Sprintf("%.1f%%", value*100)
//	})
//
//	app.LoadTemplates("*.html")
//
//	// In template:
//	// {{ .Price | currency }}      → "$19.99"
//	// {{ .Discount | percent }}    → "15.0%"
//
// Example - With Error Handling:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Register function that can fail
//	app.AddTemplateFunc("divide", func(a, b int) (int, error) {
//	    if b == 0 {
//	        return 0, errors.New("division by zero")
//	    }
//	    return a / b, nil
//	})
//
//	app.LoadTemplates("*.html")
//
//	// In template:
//	// {{ divide 10 2 }}   → 5
//	// {{ divide 10 0 }}   → ERROR: template execution fails
//
// Example - URL Building:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Register URL builder
//	app.AddTemplateFunc("url", func(path string) string {
//	    return "/app" + path
//	})
//
//	// Register asset URL builder with version
//	app.AddTemplateFunc("asset", func(path string) string {
//	    version := "v1.2.3" // Or read from config
//	    return fmt.Sprintf("/static/%s?v=%s", path, version)
//	})
//
//	app.LoadTemplates("*.html")
//
//	// In template:
//	// <a href="{{ url "/users" }}">Users</a>
//	// <script src="{{ asset "app.js" }}"></script>
//
// Example - Authorization Helper:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Register permission checker
//	app.AddTemplateFunc("can", func(user interface{}, permission string) bool {
//	    // Type assert and check permission
//	    if u, ok := user.(*User); ok {
//	        return u.HasPermission(permission)
//	    }
//	    return false
//	})
//
//	app.LoadTemplates("*.html")
//
//	// In template:
//	// {{ if can .User "edit:posts" }}
//	//   <button>Edit</button>
//	// {{ end }}
//
// Built-in Template Functions (Go standard library):
//   - and: Returns boolean AND of arguments
//   - call: Calls a function
//   - html: Escapes HTML
//   - index: Indexes into maps, slices, arrays
//   - js: Escapes JavaScript
//   - len: Returns length of string, slice, map
//   - not: Boolean negation
//   - or: Returns boolean OR of arguments
//   - print, printf, println: Formatted output
//   - urlquery: Escapes for URL query
//
// Function Naming Best Practices:
//   - Use camelCase or snake_case consistently
//   - Make names descriptive: "formatDate" better than "fd"
//   - Avoid single-letter names (except conventional ones)
//   - Don't shadow built-in functions unless intentional
//   - Consider prefixing custom functions: "app_formatDate"
//
// Performance Considerations:
//   - Functions are called during template rendering (request time)
//   - Keep functions fast (avoid heavy computation, I/O, database queries)
//   - Cache results if function is called repeatedly
//   - Consider memoization for expensive functions
//   - Profile template rendering if performance is critical
//
// Best Practices:
//   - Register all functions before loading templates
//   - Keep template functions simple and focused
//   - Move complex logic to Go code, not templates
//   - Use error returns for operations that can fail
//   - Make functions pure (no side effects) when possible
//   - Document function behavior and parameters
//   - Test template functions independently
//
// Debugging Tips:
//   - Test functions in isolation before using in templates
//   - Check function signature matches template requirements
//   - Verify template engine is initialized before calling
//   - Log function calls to debug template execution
//   - Use error returns to provide better debugging information
//
// Security Considerations:
//   - Sanitize user input in custom functions
//   - Don't pass sensitive data to templates unless necessary
//   - Validate and escape output appropriately
//   - Be careful with functions that execute arbitrary code
//   - Don't allow user-controlled function names or parameters
//
// AddTemplateFunc는 템플릿에서 사용할 수 있는 커스텀 함수를 등록합니다.
// 이를 통해 도메인별 헬퍼, 포매터 및 유틸리티를 추가하여
// Go의 내장 템플릿 함수를 넘어 템플릿 기능을 확장할 수 있습니다.
//
// AddTemplateFunc를 사용하면 템플릿이 다음을 수행할 수 있습니다:
// - 데이터 형식 지정(날짜, 숫자, 통화 등)
// - 계산 및 변환 수행
// - 템플릿에서 애플리케이션 로직 호출
// - 헬퍼 유틸리티 액세스(문자열 조작 등)
// - 커스텀 템플릿 DSL(도메인별 언어) 구현
//
// 목적:
//   - 템플릿에서 사용할 단일 커스텀 함수 등록
//   - 애플리케이션별 로직으로 템플릿 기능 확장
//   - 복잡한 로직을 Go 함수로 이동하여 깔끔한 템플릿 활성화
//   - 모든 템플릿에서 재사용 가능한 템플릿 유틸리티 제공
//   - 템플릿이 애플리케이션 컨텍스트 및 헬퍼에 액세스할 수 있도록 허용
//
// 매개변수:
//
//   - name: 템플릿에서 호출될 함수 이름.
//     유효한 Go 식별자여야 합니다(영숫자 + 밑줄, 문자/밑줄로 시작).
//     대소문자 구분: "formatDate"와 "FormatDate"는 다른 함수입니다.
//     특별한 처리 없이는 내장 템플릿 함수(and, or, not 등)를 재정의할 수 없습니다.
//
//   - fn: 템플릿에서 호출될 때 실행할 함수.
//     특정 시그니처 요구사항이 있는 Go 함수여야 합니다:
//
//   - 1개 또는 2개의 값 반환
//
//   - 2개의 값인 경우: 두 번째는 error여야 함(템플릿에서 오류 처리를 위해)
//
//   - 매개변수: 모든 타입의 매개변수를 모든 수만큼 받을 수 있음
//
//   - 매개변수 타입: 템플릿 엔진이 타입 변환을 시도
//
//   - 반환 타입: 템플릿에서 렌더링할 수 있는 타입이어야 함
//
// 등록 타이밍:
//   - **중요**: 템플릿을 로드하기 전에 호출되어야 합니다.
//   - 템플릿 로드 후 등록된 함수는 해당 템플릿에서 사용할 수 없습니다.
//   - 권장사항: 앱 생성 직후 모든 함수를 등록하세요.
//   - 템플릿을 리로드하는 경우 함수는 유지됩니다(재등록 불필요).
//
// 스레드 안전성:
//   - 읽기 잠금 보호(app.mu.RLock)로 스레드 안전합니다.
//   - 초기화 또는 런타임 중 호출해도 안전합니다.
//   - 템플릿을 로드하기 전에 함수를 등록해야 합니다.
//   - 여러 고루틴이 동시에 AddTemplateFunc를 안전하게 호출할 수 있습니다.
//   - 등록된 함수 자체는 동시에 호출되는 경우 스레드 안전해야 합니다.
//
// 모범 사례:
//   - 템플릿을 로드하기 전에 모든 함수를 등록하세요
//   - 템플릿 함수를 간단하고 집중적으로 유지하세요
//   - 복잡한 로직은 템플릿이 아닌 Go 코드로 이동하세요
//   - 실패할 수 있는 작업에는 오류 반환을 사용하세요
//   - 가능하면 함수를 순수하게 만드세요(부작용 없음)
//   - 함수 동작 및 매개변수를 문서화하세요
//   - 템플릿 함수를 독립적으로 테스트하세요
//
// 보안 고려사항:
//   - 커스텀 함수에서 사용자 입력을 삭제하세요
//   - 필요하지 않으면 민감한 데이터를 템플릿에 전달하지 마세요
//   - 출력을 적절하게 검증하고 이스케이프하세요
//   - 임의의 코드를 실행하는 함수에 주의하세요
//   - 사용자가 제어하는 함수 이름이나 매개변수를 허용하지 마세요
func (a *App) AddTemplateFunc(name string, fn interface{}) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates != nil {
		a.templates.AddFunc(name, fn)
	}
}

// AddTemplateFuncs registers multiple custom functions for use in templates.
// This is a convenience method for bulk registration of template functions,
// more efficient and cleaner than calling AddTemplateFunc multiple times.
//
// AddTemplateFuncs is ideal for:
// - Registering a suite of related helper functions
// - Bulk registration of formatting utilities
// - Sharing common template function collections across projects
// - Organizing template functions by category or domain
// - One-time registration of all application template helpers
//
// Purpose:
//   - Registers multiple custom functions in a single call
//   - Provides cleaner API than repeated AddTemplateFunc calls
//   - Enables organized template function management
//   - Supports sharing template function libraries
//   - Simplifies template function initialization
//
// Parameters:
//
//   - funcs: Map of function names to function implementations.
//     Key: Function name as it will be called in templates (string)
//     Value: Function implementation (interface{}, must be valid template function)
//
//     Map structure: map[string]interface{}{
//     "functionName": functionImplementation,
//     "otherFunc": otherImplementation,
//     ...
//     }
//
//     Each function must follow template function signature requirements:
//
//   - Returns 1 or 2 values
//
//   - If 2 values: second must be error
//
//   - Can accept any number of parameters
//
// Thread-Safety:
//   - Thread-safe with read lock protection (app.mu.RLock).
//   - Safe to call during initialization or runtime.
//   - Should be called BEFORE loading templates.
//   - Multiple goroutines can safely call AddTemplateFuncs concurrently.
//   - Registered functions should be thread-safe if called concurrently.
//
// Registration Timing:
//   - **Critical**: Must be called BEFORE loading templates.
//   - Functions registered after template loading won't be available.
//   - Recommended: Register all functions immediately after app creation.
//   - If reloading templates, functions persist (no re-registration needed).
//
// Template Engine Requirement:
//   - Template engine must be initialized (WithTemplateDir option set).
//   - Silently does nothing if template engine is nil.
//   - No error returned; check TemplateEngine() != nil if verification needed.
//
// Behavior:
//   - All functions in the map are registered with the template engine
//   - Existing functions with same names are replaced
//   - Empty map is valid (no-op, no error)
//   - Nil map is valid (no-op, no error)
//   - Functions become available in all templates after loading
//
// Common Use Cases:
//   - Register suite of string formatters (upper, lower, title, etc.)
//   - Register date/time formatting functions
//   - Register number/currency formatters
//   - Register URL builders and asset helpers
//   - Register i18n/l10n translation functions
//   - Share common template utilities across microservices
//
// Example - String Utilities:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Register multiple string functions at once
//	app.AddTemplateFuncs(map[string]interface{}{
//	    "upper":     strings.ToUpper,
//	    "lower":     strings.ToLower,
//	    "title":     strings.Title,
//	    "trim":      strings.TrimSpace,
//	    "replace":   strings.ReplaceAll,
//	})
//
//	app.LoadTemplates("*.html")
//
//	// In template:
//	// {{ .Name | upper }}
//	// {{ .Title | title }}
//	// {{ .Text | trim }}
//
// Example - Date/Time Formatters:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	app.AddTemplateFuncs(map[string]interface{}{
//	    "formatDate": func(t time.Time) string {
//	        return t.Format("2006-01-02")
//	    },
//	    "formatDateTime": func(t time.Time) string {
//	        return t.Format("2006-01-02 15:04:05")
//	    },
//	    "formatTime": func(t time.Time) string {
//	        return t.Format("15:04:05")
//	    },
//	    "timeAgo": func(t time.Time) string {
//	        d := time.Since(t)
//	        switch {
//	        case d < time.Minute:
//	            return "just now"
//	        case d < time.Hour:
//	            return fmt.Sprintf("%d min ago", int(d.Minutes()))
//	        case d < 24*time.Hour:
//	            return fmt.Sprintf("%d hours ago", int(d.Hours()))
//	        default:
//	            return fmt.Sprintf("%d days ago", int(d.Hours()/24))
//	        }
//	    },
//	})
//
//	app.LoadTemplates("*.html")
//
// Example - Number Formatters:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	app.AddTemplateFuncs(map[string]interface{}{
//	    "currency": func(amount float64) string {
//	        return fmt.Sprintf("$%.2f", amount)
//	    },
//	    "percent": func(value float64) string {
//	        return fmt.Sprintf("%.1f%%", value*100)
//	    },
//	    "round": func(value float64, decimals int) float64 {
//	        pow := math.Pow(10, float64(decimals))
//	        return math.Round(value*pow) / pow
//	    },
//	    "thousands": func(n int) string {
//	        return humanize.Comma(int64(n))
//	    },
//	})
//
//	app.LoadTemplates("*.html")
//
// Example - URL and Asset Helpers:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	baseURL := "/app"
//	version := "v1.2.3"
//
//	app.AddTemplateFuncs(map[string]interface{}{
//	    "url": func(path string) string {
//	        return baseURL + path
//	    },
//	    "asset": func(path string) string {
//	        return fmt.Sprintf("/static/%s?v=%s", path, version)
//	    },
//	    "cdn": func(path string) string {
//	        return fmt.Sprintf("https://cdn.example.com/%s", path)
//	    },
//	    "thumbnail": func(imageURL string, size string) string {
//	        return fmt.Sprintf("%s?size=%s", imageURL, size)
//	    },
//	})
//
//	app.LoadTemplates("*.html")
//
//	// In template:
//	// <a href="{{ url "/users" }}">Users</a>
//	// <script src="{{ asset "app.js" }}"></script>
//	// <img src="{{ cdn "logo.png" }}">
//
// Example - Comprehensive Helper Suite:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// Register complete helper suite
//	app.AddTemplateFuncs(map[string]interface{}{
//	    // String helpers
//	    "upper":     strings.ToUpper,
//	    "lower":     strings.ToLower,
//	    "truncate":  func(s string, n int) string {
//	        if len(s) <= n {
//	            return s
//	        }
//	        return s[:n] + "..."
//	    },
//
//	    // Date helpers
//	    "formatDate": func(t time.Time) string {
//	        return t.Format("2006-01-02")
//	    },
//
//	    // Number helpers
//	    "currency": func(amount float64) string {
//	        return fmt.Sprintf("$%.2f", amount)
//	    },
//
//	    // Logic helpers
//	    "default": func(value, defaultValue interface{}) interface{} {
//	        if value == nil || value == "" {
//	            return defaultValue
//	        }
//	        return value
//	    },
//
//	    // Collection helpers
//	    "first": func(slice interface{}) interface{} {
//	        v := reflect.ValueOf(slice)
//	        if v.Kind() == reflect.Slice && v.Len() > 0 {
//	            return v.Index(0).Interface()
//	        }
//	        return nil
//	    },
//	})
//
//	app.LoadTemplates("*.html")
//
// Example - Reusable Function Library:
//
//	// helpers/template_funcs.go
//	package helpers
//
//	func StandardTemplateFuncs() map[string]interface{} {
//	    return map[string]interface{}{
//	        "upper":      strings.ToUpper,
//	        "lower":      strings.ToLower,
//	        "formatDate": func(t time.Time) string {
//	            return t.Format("2006-01-02")
//	        },
//	        // ... more functions ...
//	    }
//	}
//
//	// main.go
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//	app.AddTemplateFuncs(helpers.StandardTemplateFuncs())
//	app.LoadTemplates("*.html")
//
// Example - Domain-Specific Functions:
//
//	app := websvrutil.New(websvrutil.WithTemplateDir("./views"))
//
//	// E-commerce specific functions
//	app.AddTemplateFuncs(map[string]interface{}{
//	    "productURL": func(productID string) string {
//	        return fmt.Sprintf("/products/%s", productID)
//	    },
//	    "stockStatus": func(quantity int) string {
//	        if quantity == 0 {
//	            return "Out of Stock"
//	        } else if quantity < 10 {
//	            return "Low Stock"
//	        }
//	        return "In Stock"
//	    },
//	    "priceWithTax": func(price float64, taxRate float64) float64 {
//	        return price * (1 + taxRate)
//	    },
//	    "discountPrice": func(price, discount float64) float64 {
//	        return price * (1 - discount)
//	    },
//	})
//
//	app.LoadTemplates("*.html")
//
// Organizing Template Functions:
//   - Group by category: strings, dates, numbers, URLs, auth, etc.
//   - Use namespaced names: "str_upper", "date_format", "url_build"
//   - Keep related functions together in source code
//   - Document function behavior and parameters
//   - Consider separate packages for different function categories
//
// Performance Considerations:
//   - Registration overhead is negligible (done once at startup)
//   - Functions are called during template rendering (request time)
//   - Keep all functions fast (avoid heavy computation, I/O)
//   - Cache results if function is called repeatedly
//   - Profile template rendering if performance is critical
//
// Best Practices:
//   - Register all functions before loading templates
//   - Group related functions in maps for clarity
//   - Use consistent naming conventions
//   - Document function purpose and signature
//   - Keep functions simple and focused
//   - Test functions independently before template use
//   - Consider creating reusable function libraries
//   - Use type-safe signatures when possible
//
// Comparison with AddTemplateFunc:
//   - AddTemplateFuncs: Multiple functions, single call, cleaner code
//   - AddTemplateFunc: Single function, multiple calls, more verbose
//   - Use AddTemplateFuncs for: Bulk registration, initialization, function suites
//   - Use AddTemplateFunc for: Individual additions, dynamic registration
//
// Debugging Tips:
//   - Test each function independently before adding to map
//   - Verify function signatures match template requirements
//   - Check template engine is initialized before calling
//   - Log function registration for debugging
//   - Use descriptive function names for easier debugging
//
// Security Considerations:
//   - Sanitize user input in all custom functions
//   - Don't pass sensitive data to templates unless necessary
//   - Validate and escape function outputs
//   - Be careful with functions that execute code or access resources
//   - Don't allow user-controlled function names or implementations
//
// AddTemplateFuncs는 템플릿에서 사용할 여러 커스텀 함수를 등록합니다.
// 이는 템플릿 함수의 대량 등록을 위한 편의 메서드로,
// AddTemplateFunc를 여러 번 호출하는 것보다 효율적이고 깔끔합니다.
//
// AddTemplateFuncs는 다음에 이상적입니다:
// - 관련된 헬퍼 함수 모음 등록
// - 포맷 유틸리티의 대량 등록
// - 프로젝트 간 공통 템플릿 함수 컬렉션 공유
// - 범주 또는 도메인별 템플릿 함수 구성
// - 모든 애플리케이션 템플릿 헬퍼의 일회성 등록
//
// 목적:
//   - 단일 호출로 여러 커스텀 함수 등록
//   - 반복된 AddTemplateFunc 호출보다 깔끔한 API 제공
//   - 조직화된 템플릿 함수 관리 활성화
//   - 템플릿 함수 라이브러리 공유 지원
//   - 템플릿 함수 초기화 단순화
//
// 매개변수:
//
//   - funcs: 함수 이름에서 함수 구현으로의 맵.
//     키: 템플릿에서 호출될 함수 이름(string)
//     값: 함수 구현(interface{}, 유효한 템플릿 함수여야 함)
//
//     맵 구조: map[string]interface{}{
//     "functionName": functionImplementation,
//     "otherFunc": otherImplementation,
//     ...
//     }
//
//     각 함수는 템플릿 함수 시그니처 요구사항을 따라야 합니다:
//
//   - 1개 또는 2개의 값 반환
//
//   - 2개의 값인 경우: 두 번째는 error여야 함
//
//   - 모든 수의 매개변수를 받을 수 있음
//
// 등록 타이밍:
//   - **중요**: 템플릿을 로드하기 전에 호출되어야 합니다.
//   - 템플릿 로드 후 등록된 함수는 사용할 수 없습니다.
//   - 권장사항: 앱 생성 직후 모든 함수를 등록하세요.
//   - 템플릿을 리로드하는 경우 함수는 유지됩니다(재등록 불필요).
//
// 스레드 안전성:
//   - 읽기 잠금 보호(app.mu.RLock)로 스레드 안전합니다.
//   - 초기화 또는 런타임 중 호출해도 안전합니다.
//   - 템플릿을 로드하기 전에 호출해야 합니다.
//   - 여러 고루틴이 동시에 AddTemplateFuncs를 안전하게 호출할 수 있습니다.
//   - 등록된 함수는 동시에 호출되는 경우 스레드 안전해야 합니다.
//
// 모범 사례:
//   - 템플릿을 로드하기 전에 모든 함수를 등록하세요
//   - 명확성을 위해 맵에서 관련 함수를 그룹화하세요
//   - 일관된 명명 규칙을 사용하세요
//   - 함수 목적 및 시그니처를 문서화하세요
//   - 함수를 간단하고 집중적으로 유지하세요
//   - 템플릿 사용 전에 함수를 독립적으로 테스트하세요
//   - 재사용 가능한 함수 라이브러리 생성을 고려하세요
//   - 가능하면 타입 안전 시그니처를 사용하세요
//
// 보안 고려사항:
//   - 모든 커스텀 함수에서 사용자 입력을 삭제하세요
//   - 필요하지 않으면 민감한 데이터를 템플릿에 전달하지 마세요
//   - 함수 출력을 검증하고 이스케이프하세요
//   - 코드를 실행하거나 리소스에 액세스하는 함수에 주의하세요
//   - 사용자가 제어하는 함수 이름이나 구현을 허용하지 마세요
func (a *App) AddTemplateFuncs(funcs map[string]interface{}) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.templates != nil {
		a.templates.AddFuncs(funcs)
	}
}
