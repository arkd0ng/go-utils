package websvrutil

import (
	"time"
)

// Options holds configuration options for the App.
// Options는 App의 설정 옵션을 보유합니다.
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
}

// Option is a functional option for configuring the App.
// Option은 App 설정을 위한 함수형 옵션입니다.
type Option func(*Options)

// defaultOptions returns the default options for the App.
// defaultOptions는 App의 기본 옵션을 반환합니다.
func defaultOptions() *Options {
	return &Options{
		ReadTimeout:      15 * time.Second,
		WriteTimeout:     15 * time.Second,
		IdleTimeout:      60 * time.Second,
		MaxHeaderBytes:   1 << 20, // 1 MB
		TemplateDir:      "templates",
		StaticDir:        "static",
		StaticPrefix:     "/static",
		EnableAutoReload: false,
		EnableLogger:     true,
		EnableRecovery:   true,
		MaxUploadSize:    32 << 20, // 32 MB default
	}
}

// WithReadTimeout sets the read timeout for the server.
// WithReadTimeout은 서버의 읽기 시간 초과를 설정합니다.
func WithReadTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = d
	}
}

// WithWriteTimeout sets the write timeout for the server.
// WithWriteTimeout은 서버의 쓰기 시간 초과를 설정합니다.
func WithWriteTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = d
	}
}

// WithIdleTimeout sets the idle timeout for the server.
// WithIdleTimeout은 서버의 유휴 시간 초과를 설정합니다.
func WithIdleTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.IdleTimeout = d
	}
}

// WithMaxHeaderBytes sets the maximum header bytes for the server.
// WithMaxHeaderBytes는 서버의 최대 헤더 바이트를 설정합니다.
func WithMaxHeaderBytes(n int) Option {
	return func(o *Options) {
		o.MaxHeaderBytes = n
	}
}

// WithTemplateDir sets the template directory.
// WithTemplateDir은 템플릿 디렉토리를 설정합니다.
func WithTemplateDir(dir string) Option {
	return func(o *Options) {
		o.TemplateDir = dir
	}
}

// WithStaticDir sets the static files directory.
// WithStaticDir은 정적 파일 디렉토리를 설정합니다.
func WithStaticDir(dir string) Option {
	return func(o *Options) {
		o.StaticDir = dir
	}
}

// WithStaticPrefix sets the URL prefix for static files.
// WithStaticPrefix는 정적 파일의 URL 접두사를 설정합니다.
func WithStaticPrefix(prefix string) Option {
	return func(o *Options) {
		o.StaticPrefix = prefix
	}
}

// WithAutoReload enables or disables automatic template reloading.
// WithAutoReload은 자동 템플릿 재로드를 활성화하거나 비활성화합니다.
func WithAutoReload(enable bool) Option {
	return func(o *Options) {
		o.EnableAutoReload = enable
	}
}

// WithLogger enables or disables the built-in logger middleware.
// WithLogger는 내장 로거 미들웨어를 활성화하거나 비활성화합니다.
func WithLogger(enable bool) Option {
	return func(o *Options) {
		o.EnableLogger = enable
	}
}

// WithRecovery enables or disables the built-in recovery middleware.
// WithRecovery는 내장 복구 미들웨어를 활성화하거나 비활성화합니다.
func WithRecovery(enable bool) Option {
	return func(o *Options) {
		o.EnableRecovery = enable
	}
}

// WithMaxUploadSize sets the maximum file upload size in bytes.
// WithMaxUploadSize는 최대 파일 업로드 크기(바이트)를 설정합니다.
func WithMaxUploadSize(size int64) Option {
	return func(o *Options) {
		o.MaxUploadSize = size
	}
}
