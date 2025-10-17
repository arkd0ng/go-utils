package websvrutil

import (
	"time"
)

// Options holds all configuration settings for an App instance.
// This structure uses the functional options pattern for flexible, extensible configuration
// without breaking API compatibility when adding new options.
//
// Options provides:
// - Server timeouts (read, write, idle)
// - HTTP header size limits
// - Template and static file serving configuration
// - Development features (auto-reload)
// - Built-in middleware controls (logging, recovery)
// - Request size limits (uploads, body)
//
// Usage:
//
//	Use With* functions to configure:
//	  app := websvrutil.New(
//	      websvrutil.WithReadTimeout(10*time.Second),
//	      websvrutil.WithTemplateDir("./views"),
//	      websvrutil.WithLogger(true),
//	  )
//
// Default Values:
//
//	See defaultOptions() for complete default configuration.
//	All options have sensible production-ready defaults.
//
// Options는 App 인스턴스에 대한 모든 구성 설정을 보유합니다.
// 이 구조는 새 옵션을 추가할 때 API 호환성을 깨지 않고
// 유연하고 확장 가능한 구성을 위해 함수형 옵션 패턴을 사용합니다.
//
// Options는 다음을 제공합니다:
// - 서버 타임아웃(읽기, 쓰기, 유휴)
// - HTTP 헤더 크기 제한
// - 템플릿 및 정적 파일 서빙 구성
// - 개발 기능(자동 리로드)
// - 내장 미들웨어 제어(로깅, 복구)
// - 요청 크기 제한(업로드, 본문)
type Options struct {
	// ReadTimeout is the maximum duration for reading the entire request, including the body.
	// ReadTimeout은 본문을 포함하여 전체 요청을 읽는 최대 기간입니다.
	ReadTimeout time.Duration

	// WriteTimeout is the maximum duration before timing out writes of the response.
	// WriteTimeout은 응답 쓰기 시간 초과 전 최대 기간입니다.
	WriteTimeout time.Duration

	// IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled.
	// IdleTimeout은 keep-alive가 활성화된 경우 다음 요청을 기다리는 최대 시간입니다.
	IdleTimeout time.Duration

	// MaxHeaderBytes controls the maximum number of bytes the server will read parsing the request header.
	// MaxHeaderBytes는 서버가 요청 헤더를 파싱할 때 읽을 최대 바이트 수를 제어합니다.
	MaxHeaderBytes int

	// TemplateDir is the directory where HTML templates are stored.
	// TemplateDir은 HTML 템플릿이 저장된 디렉토리입니다.
	TemplateDir string

	// StaticDir is the directory where static files are served from.
	// StaticDir은 정적 파일이 제공되는 디렉토리입니다.
	StaticDir string

	// StaticPrefix is the URL prefix for static files.
	// StaticPrefix는 정적 파일의 URL 접두사입니다.
	StaticPrefix string

	// EnableAutoReload enables automatic template reloading in development mode.
	// EnableAutoReload은 개발 모드에서 자동 템플릿 재로드를 활성화합니다.
	EnableAutoReload bool

	// EnableLogger enables built-in request logging middleware.
	// EnableLogger는 내장 요청 로깅 미들웨어를 활성화합니다.
	EnableLogger bool

	// EnableRecovery enables built-in panic recovery middleware.
	// EnableRecovery는 내장 패닉 복구 미들웨어를 활성화합니다.
	EnableRecovery bool

	// MaxUploadSize is the maximum allowed file upload size in bytes.
	// MaxUploadSize는 허용되는 최대 파일 업로드 크기(바이트)입니다.
	MaxUploadSize int64

	// MaxBodySize is the maximum allowed request body size in bytes (for JSON, etc).
	// MaxBodySize는 허용되는 최대 요청 본문 크기(바이트)입니다 (JSON 등).
	MaxBodySize int64
}

// Option is a functional option type for configuring App instances.
// Follows the functional options pattern for clean, extensible API design.
//
// Pattern Benefits:
// - Backward compatible (new options don't break existing code)
// - Self-documenting (With* function names describe purpose)
// - Optional parameters (use only what you need)
// - Chainable configuration
//
// Usage:
//
//	app := websvrutil.New(
//	    WithReadTimeout(5*time.Second),
//	    WithLogger(false),
//	)
//
// Option은 App 인스턴스 구성을 위한 함수형 옵션 타입입니다.
// 깔끔하고 확장 가능한 API 디자인을 위한 함수형 옵션 패턴을 따릅니다.
type Option func(*Options)

// defaultOptions returns default configuration values for a new App.
// These defaults are production-ready and balanced for typical web applications.
//
// Default Values:
// - ReadTimeout: 15s (DefaultReadTimeout)
// - WriteTimeout: 15s (DefaultWriteTimeout)
// - IdleTimeout: 60s (DefaultIdleTimeout)
// - MaxHeaderBytes: 1MB (DefaultMaxHeaderBytes)
// - TemplateDir: "templates"
// - StaticDir: "static"
// - StaticPrefix: "/static"
// - EnableAutoReload: false (development only)
// - EnableLogger: true
// - EnableRecovery: true
// - MaxUploadSize: 32MB (DefaultMaxUploadSize)
// - MaxBodySize: 10MB (DefaultMaxBodySize)
//
// defaultOptions는 새 App에 대한 기본 구성 값을 반환합니다.
// 이러한 기본값은 프로덕션 준비가 되어 있으며 일반적인 웹 애플리케이션에 균형 잡혀 있습니다.
func defaultOptions() *Options {
	return &Options{
		ReadTimeout:      DefaultReadTimeout,
		WriteTimeout:     DefaultWriteTimeout,
		IdleTimeout:      DefaultIdleTimeout,
		MaxHeaderBytes:   DefaultMaxHeaderBytes,
		TemplateDir:      "templates",
		StaticDir:        "static",
		StaticPrefix:     "/static",
		EnableAutoReload: false,
		EnableLogger:     true,
		EnableRecovery:   true,
		MaxUploadSize:    DefaultMaxUploadSize,
		MaxBodySize:      DefaultMaxBodySize,
	}
}

// WithReadTimeout sets the maximum duration for reading entire request including body.
// Default: 15 seconds. Prevents slow clients from holding connections.
//
// WithReadTimeout은 본문을 포함한 전체 요청 읽기의 최대 기간을 설정합니다.
// 기본값: 15초. 느린 클라이언트가 연결을 보유하는 것을 방지합니다.
func WithReadTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = d
	}
}

// WithWriteTimeout sets the maximum duration for writing response before timeout.
// Default: 15 seconds. Prevents slow clients from holding connections.
//
// WithWriteTimeout은 시간 초과 전 응답 쓰기의 최대 기간을 설정합니다.
// 기본값: 15초. 느린 클라이언트가 연결을 보유하는 것을 방지합니다.
func WithWriteTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = d
	}
}

// WithIdleTimeout sets the maximum time to wait for next request when keep-alive is enabled.
// Default: 60 seconds. Balances connection reuse and resource consumption.
//
// WithIdleTimeout은 keep-alive가 활성화된 경우 다음 요청을 기다리는 최대 시간을 설정합니다.
// 기본값: 60초. 연결 재사용과 리소스 소비의 균형을 맞춥니다.
func WithIdleTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.IdleTimeout = d
	}
}

// WithMaxHeaderBytes sets the maximum bytes server will read parsing request headers.
// Default: 1MB. Prevents malicious large headers from consuming memory.
//
// WithMaxHeaderBytes는 서버가 요청 헤더를 파싱할 때 읽을 최대 바이트를 설정합니다.
// 기본값: 1MB. 악의적인 대형 헤더가 메모리를 소비하는 것을 방지합니다.
func WithMaxHeaderBytes(n int) Option {
	return func(o *Options) {
		o.MaxHeaderBytes = n
	}
}

// WithTemplateDir sets the directory where HTML templates are stored.
// Default: "templates". Directory is relative to application working directory.
//
// WithTemplateDir은 HTML 템플릿이 저장되는 디렉토리를 설정합니다.
// 기본값: "templates". 디렉토리는 애플리케이션 작업 디렉토리에 상대적입니다.
func WithTemplateDir(dir string) Option {
	return func(o *Options) {
		o.TemplateDir = dir
	}
}

// WithStaticDir sets the directory where static files are served from.
// Default: "static". Used with Static() method for file serving.
//
// WithStaticDir은 정적 파일이 제공되는 디렉토리를 설정합니다.
// 기본값: "static". 파일 서빙을 위해 Static() 메서드와 함께 사용됩니다.
func WithStaticDir(dir string) Option {
	return func(o *Options) {
		o.StaticDir = dir
	}
}

// WithStaticPrefix sets the URL prefix for static file routes.
// Default: "/static". Example: "/static" maps to StaticDir on filesystem.
//
// WithStaticPrefix는 정적 파일 라우트의 URL 접두사를 설정합니다.
// 기본값: "/static". 예: "/static"은 파일 시스템의 StaticDir에 매핑됩니다.
func WithStaticPrefix(prefix string) Option {
	return func(o *Options) {
		o.StaticPrefix = prefix
	}
}

// WithAutoReload enables or disables automatic template reloading.
// Default: false. Enable only in development mode for hot reload.
//
// WithAutoReload은 자동 템플릿 재로드를 활성화하거나 비활성화합니다.
// 기본값: false. 핫 리로드를 위해 개발 모드에서만 활성화하세요.
func WithAutoReload(enable bool) Option {
	return func(o *Options) {
		o.EnableAutoReload = enable
	}
}

// WithLogger enables or disables built-in request logging middleware.
// Default: true. Logs HTTP method, path, status code, and duration.
//
// WithLogger는 내장 요청 로깅 미들웨어를 활성화하거나 비활성화합니다.
// 기본값: true. HTTP 메서드, 경로, 상태 코드 및 기간을 로깅합니다.
func WithLogger(enable bool) Option {
	return func(o *Options) {
		o.EnableLogger = enable
	}
}

// WithRecovery enables or disables built-in panic recovery middleware.
// Default: true. Catches panics and returns 500 Internal Server Error.
//
// WithRecovery는 내장 패닉 복구 미들웨어를 활성화하거나 비활성화합니다.
// 기본값: true. 패닉을 캐치하고 500 Internal Server Error를 반환합니다.
func WithRecovery(enable bool) Option {
	return func(o *Options) {
		o.EnableRecovery = enable
	}
}

// WithMaxUploadSize sets the maximum allowed file upload size in bytes.
// Default: 32MB. Prevents excessive memory usage from large file uploads.
//
// WithMaxUploadSize는 허용되는 최대 파일 업로드 크기(바이트)를 설정합니다.
// 기본값: 32MB. 대용량 파일 업로드로 인한 과도한 메모리 사용을 방지합니다.
func WithMaxUploadSize(size int64) Option {
	return func(o *Options) {
		o.MaxUploadSize = size
	}
}

// WithMaxBodySize sets the maximum allowed request body size in bytes.
// This limit applies to JSON, form data, and other request bodies (not file uploads).
// Default: 10MB. Prevents memory exhaustion from large payloads.
//
// Example:
//
//	app := websvrutil.New(
//	    websvrutil.WithMaxBodySize(5 * 1024 * 1024), // 5 MB
//	)
//
// WithMaxBodySize는 허용되는 최대 요청 본문 크기(바이트)를 설정합니다.
// 이 제한은 JSON, 폼 데이터 및 기타 요청 본문에 적용됩니다(파일 업로드 제외).
// 기본값: 10MB. 대용량 페이로드로 인한 메모리 고갈을 방지합니다.
func WithMaxBodySize(size int64) Option {
	return func(o *Options) {
		o.MaxBodySize = size
	}
}
