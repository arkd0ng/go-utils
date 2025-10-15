package httputil

import (
	"crypto/tls"
	"net/http"
	"time"
)

// Option is a functional option for configuring HTTP requests and clients.
// Option은 HTTP 요청 및 클라이언트를 설정하기 위한 함수형 옵션입니다.
type Option func(*config)

// config holds the configuration for HTTP requests.
// config는 HTTP 요청의 설정을 보유합니다.
type config struct {
	// Request configuration / 요청 설정
	headers     map[string]string
	queryParams map[string]string
	timeout     time.Duration
	userAgent   string

	// Authentication / 인증
	bearerToken string
	basicAuthUser string
	basicAuthPass string

	// Retry configuration / 재시도 설정
	maxRetries int
	retryMin   time.Duration
	retryMax   time.Duration

	// Client configuration / 클라이언트 설정
	baseURL         string
	followRedirects bool
	maxRedirects    int
	tlsConfig       *tls.Config
	proxyURL        string
	cookieJar       http.CookieJar

	// Cookie jar configuration / 쿠키 저장소 설정
	enableCookieJar bool
	cookieJarPath   string

	// Logging / 로깅
	logger Logger
}

// defaultConfig returns a configuration with sensible defaults.
// defaultConfig는 합리적인 기본값을 가진 설정을 반환합니다.
func defaultConfig() *config {
	return &config{
		headers:         make(map[string]string),
		queryParams:     make(map[string]string),
		timeout:         30 * time.Second,
		userAgent:       "go-utils/httputil v" + Version,
		maxRetries:      3,
		retryMin:        100 * time.Millisecond,
		retryMax:        5 * time.Second,
		followRedirects: true,
		maxRedirects:    10,
	}
}

// apply applies all options to the configuration.
// apply는 모든 옵션을 설정에 적용합니다.
func (c *config) apply(opts []Option) {
	for _, opt := range opts {
		opt(c)
	}
}

// WithTimeout sets the request timeout.
// WithTimeout은 요청 타임아웃을 설정합니다.
//
// Default: 30 seconds
// 기본값: 30초
func WithTimeout(timeout time.Duration) Option {
	return func(c *config) {
		c.timeout = timeout
	}
}

// WithHeaders sets custom headers for the request.
// WithHeaders는 요청에 대한 사용자 정의 헤더를 설정합니다.
func WithHeaders(headers map[string]string) Option {
	return func(c *config) {
		for k, v := range headers {
			c.headers[k] = v
		}
	}
}

// WithHeader sets a single custom header for the request.
// WithHeader는 요청에 대한 단일 사용자 정의 헤더를 설정합니다.
func WithHeader(key, value string) Option {
	return func(c *config) {
		c.headers[key] = value
	}
}

// WithQueryParams sets query parameters for the request.
// WithQueryParams는 요청에 대한 쿼리 매개변수를 설정합니다.
func WithQueryParams(params map[string]string) Option {
	return func(c *config) {
		for k, v := range params {
			c.queryParams[k] = v
		}
	}
}

// WithBearerToken sets the Bearer token for authentication.
// WithBearerToken은 인증을 위한 Bearer 토큰을 설정합니다.
func WithBearerToken(token string) Option {
	return func(c *config) {
		c.bearerToken = token
	}
}

// WithBasicAuth sets Basic Authentication credentials.
// WithBasicAuth는 기본 인증 자격 증명을 설정합니다.
func WithBasicAuth(username, password string) Option {
	return func(c *config) {
		c.basicAuthUser = username
		c.basicAuthPass = password
	}
}

// WithRetry sets the maximum number of retry attempts.
// WithRetry는 최대 재시도 시도 횟수를 설정합니다.
//
// Default: 3
// 기본값: 3
func WithRetry(maxRetries int) Option {
	return func(c *config) {
		c.maxRetries = maxRetries
	}
}

// WithRetryBackoff sets the minimum and maximum backoff time for retries.
// WithRetryBackoff는 재시도에 대한 최소 및 최대 백오프 시간을 설정합니다.
//
// Default: min=100ms, max=5s
// 기본값: min=100ms, max=5s
func WithRetryBackoff(min, max time.Duration) Option {
	return func(c *config) {
		c.retryMin = min
		c.retryMax = max
	}
}

// WithUserAgent sets a custom User-Agent header.
// WithUserAgent는 사용자 정의 User-Agent 헤더를 설정합니다.
//
// Default: "go-utils/httputil v{version}"
// 기본값: "go-utils/httputil v{version}"
func WithUserAgent(userAgent string) Option {
	return func(c *config) {
		c.userAgent = userAgent
	}
}

// WithBaseURL sets the base URL for the client.
// WithBaseURL은 클라이언트의 기본 URL을 설정합니다.
//
// This is useful when making multiple requests to the same API.
// 동일한 API에 여러 요청을 할 때 유용합니다.
func WithBaseURL(baseURL string) Option {
	return func(c *config) {
		c.baseURL = baseURL
	}
}

// WithFollowRedirects enables or disables following HTTP redirects.
// WithFollowRedirects는 HTTP 리디렉션 따르기를 활성화하거나 비활성화합니다.
//
// Default: true
// 기본값: true
func WithFollowRedirects(follow bool) Option {
	return func(c *config) {
		c.followRedirects = follow
	}
}

// WithMaxRedirects sets the maximum number of redirects to follow.
// WithMaxRedirects는 따를 최대 리디렉션 수를 설정합니다.
//
// Default: 10
// 기본값: 10
func WithMaxRedirects(max int) Option {
	return func(c *config) {
		c.maxRedirects = max
	}
}

// WithTLSConfig sets a custom TLS configuration.
// WithTLSConfig는 사용자 정의 TLS 설정을 지정합니다.
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(c *config) {
		c.tlsConfig = tlsConfig
	}
}

// WithProxy sets the proxy URL.
// WithProxy는 프록시 URL을 설정합니다.
func WithProxy(proxyURL string) Option {
	return func(c *config) {
		c.proxyURL = proxyURL
	}
}

// WithCookieJar sets a custom cookie jar.
// WithCookieJar는 사용자 정의 쿠키 저장소를 설정합니다.
func WithCookieJar(jar http.CookieJar) Option {
	return func(c *config) {
		c.cookieJar = jar
	}
}

// WithCookies enables cookie management with an in-memory cookie jar.
// WithCookies는 메모리 내 쿠키 저장소를 사용한 쿠키 관리를 활성화합니다.
//
// This creates a temporary cookie jar that will be discarded when the client is closed.
// 이는 클라이언트가 닫힐 때 삭제되는 임시 쿠키 저장소를 생성합니다.
//
// For persistent cookies, use WithPersistentCookies instead.
// 지속적인 쿠키를 위해서는 WithPersistentCookies를 사용하세요.
func WithCookies() Option {
	return func(c *config) {
		c.enableCookieJar = true
	}
}

// WithPersistentCookies enables cookie management with file persistence.
// WithPersistentCookies는 파일 지속성을 가진 쿠키 관리를 활성화합니다.
//
// Cookies will be automatically saved to and loaded from the specified file path.
// 쿠키는 지정된 파일 경로에 자동으로 저장되고 로드됩니다.
//
// Example:
//   client := httputil.NewClient(
//       httputil.WithPersistentCookies("cookies.json"),
//   )
func WithPersistentCookies(filePath string) Option {
	return func(c *config) {
		c.cookieJarPath = filePath
	}
}

// WithLogger sets a custom logger.
// WithLogger는 사용자 정의 로거를 설정합니다.
func WithLogger(logger Logger) Option {
	return func(c *config) {
		c.logger = logger
	}
}

// Logger is an interface for logging HTTP requests and responses.
// Logger는 HTTP 요청 및 응답을 로깅하기 위한 인터페이스입니다.
type Logger interface {
	// Log logs a message with key-value pairs
	// Log는 키-값 쌍과 함께 메시지를 로깅합니다
	Log(msg string, keysAndValues ...interface{})
}

// noopLogger is a no-op logger that does nothing.
// noopLogger는 아무것도 하지 않는 no-op 로거입니다.
type noopLogger struct{}

func (noopLogger) Log(msg string, keysAndValues ...interface{}) {}
