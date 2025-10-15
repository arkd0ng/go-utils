package websvrutil

import (
	"compress/gzip"
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

// Recovery returns a middleware that recovers from panics.
// Recovery는 패닉에서 복구하는 미들웨어를 반환합니다.
//
// When a panic occurs, it logs the error and stack trace, then sends a 500 response.
// 패닉이 발생하면 에러와 스택 트레이스를 로깅하고 500 응답을 전송합니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.Recovery())
func Recovery() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// Log the panic with stack trace
					// 스택 트레이스와 함께 패닉 로깅
					log.Printf("PANIC: %v\n%s", err, debug.Stack())

					// Send 500 Internal Server Error
					// 500 Internal Server Error 전송
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

// RecoveryWithConfig returns a Recovery middleware with custom configuration.
// RecoveryWithConfig는 커스텀 설정으로 Recovery 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.RecoveryWithConfig(websvrutil.RecoveryConfig{
//	    PrintStack: true,
//	    LogFunc: func(err interface{}, stack []byte) {
//	        log.Printf("Panic: %v\n%s", err, stack)
//	    },
//	}))
func RecoveryWithConfig(config RecoveryConfig) MiddlewareFunc {
	// Set defaults if not provided
	// 제공되지 않은 경우 기본값 설정
	if config.LogFunc == nil {
		config.LogFunc = func(err interface{}, stack []byte) {
			if config.PrintStack {
				log.Printf("PANIC: %v\n%s", err, stack)
			} else {
				log.Printf("PANIC: %v", err)
			}
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					stack := debug.Stack()
					config.LogFunc(err, stack)

					// Send error response
					// 에러 응답 전송
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

// RecoveryConfig defines configuration for Recovery middleware.
// RecoveryConfig는 Recovery 미들웨어의 설정을 정의합니다.
type RecoveryConfig struct {
	// PrintStack determines whether to print the stack trace
	// PrintStack은 스택 트레이스를 출력할지 결정합니다
	PrintStack bool

	// LogFunc is called when a panic is recovered
	// LogFunc은 패닉이 복구될 때 호출됩니다
	LogFunc func(err interface{}, stack []byte)
}

// Logger returns a middleware that logs HTTP requests.
// Logger는 HTTP 요청을 로깅하는 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.Logger())
func Logger() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Create a response writer wrapper to capture status code
			// 상태 코드를 캡처하기 위한 응답 작성기 래퍼 생성
			wrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			// Process request
			// 요청 처리
			next.ServeHTTP(wrapper, r)

			// Log request
			// 요청 로깅
			duration := time.Since(start)
			log.Printf("%s %s %d %v", r.Method, r.URL.Path, wrapper.statusCode, duration)
		})
	}
}

// LoggerWithConfig returns a Logger middleware with custom configuration.
// LoggerWithConfig는 커스텀 설정으로 Logger 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.LoggerWithConfig(websvrutil.LoggerConfig{
//	    Format: "${method} ${path} ${status} ${duration}",
//	    LogFunc: func(format string, args ...interface{}) {
//	        log.Printf(format, args...)
//	    },
//	}))
func LoggerWithConfig(config LoggerConfig) MiddlewareFunc {
	// Set defaults if not provided
	// 제공되지 않은 경우 기본값 설정
	if config.LogFunc == nil {
		config.LogFunc = func(method, path string, status int, duration time.Duration) {
			log.Printf("%s %s %d %v", method, path, status, duration)
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(wrapper, r)

			duration := time.Since(start)
			config.LogFunc(r.Method, r.URL.Path, wrapper.statusCode, duration)
		})
	}
}

// LoggerConfig defines configuration for Logger middleware.
// LoggerConfig는 Logger 미들웨어의 설정을 정의합니다.
type LoggerConfig struct {
	// LogFunc is called to log each request
	// LogFunc은 각 요청을 로깅하기 위해 호출됩니다
	LogFunc func(method, path string, status int, duration time.Duration)
}

// CORS returns a middleware that handles Cross-Origin Resource Sharing.
// CORS는 Cross-Origin Resource Sharing을 처리하는 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.CORS())
func CORS() MiddlewareFunc {
	return CORSWithConfig(CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	})
}

// CORSWithConfig returns a CORS middleware with custom configuration.
// CORSWithConfig는 커스텀 설정으로 CORS 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.CORSWithConfig(websvrutil.CORSConfig{
//	    AllowOrigins: []string{"https://example.com"},
//	    AllowMethods: []string{"GET", "POST"},
//	    AllowHeaders: []string{"Content-Type"},
//	    AllowCredentials: true,
//	}))
func CORSWithConfig(config CORSConfig) MiddlewareFunc {
	// Set defaults if not provided
	// 제공되지 않은 경우 기본값 설정
	if len(config.AllowOrigins) == 0 {
		config.AllowOrigins = []string{"*"}
	}
	if len(config.AllowMethods) == 0 {
		config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	}
	if len(config.AllowHeaders) == 0 {
		config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			// CORS 헤더 설정
			origin := r.Header.Get("Origin")
			if origin != "" && isOriginAllowed(origin, config.AllowOrigins) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			} else if len(config.AllowOrigins) == 1 && config.AllowOrigins[0] == "*" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}

			if len(config.AllowMethods) > 0 {
				w.Header().Set("Access-Control-Allow-Methods", joinStrings(config.AllowMethods, ", "))
			}

			if len(config.AllowHeaders) > 0 {
				w.Header().Set("Access-Control-Allow-Headers", joinStrings(config.AllowHeaders, ", "))
			}

			if config.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if config.MaxAge > 0 {
				w.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", int(config.MaxAge.Seconds())))
			}

			// Handle preflight requests
			// 프리플라이트 요청 처리
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// CORSConfig defines configuration for CORS middleware.
// CORSConfig는 CORS 미들웨어의 설정을 정의합니다.
type CORSConfig struct {
	// AllowOrigins defines allowed origins
	// AllowOrigins는 허용된 오리진을 정의합니다
	AllowOrigins []string

	// AllowMethods defines allowed HTTP methods
	// AllowMethods는 허용된 HTTP 메서드를 정의합니다
	AllowMethods []string

	// AllowHeaders defines allowed headers
	// AllowHeaders는 허용된 헤더를 정의합니다
	AllowHeaders []string

	// AllowCredentials indicates whether credentials are allowed
	// AllowCredentials는 자격 증명 허용 여부를 나타냅니다
	AllowCredentials bool

	// MaxAge indicates how long preflight results can be cached
	// MaxAge는 프리플라이트 결과를 캐시할 수 있는 시간을 나타냅니다
	MaxAge time.Duration
}

// responseWriter wraps http.ResponseWriter to capture status code.
// responseWriter는 상태 코드를 캡처하기 위해 http.ResponseWriter를 래핑합니다.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and calls the underlying WriteHeader.
// WriteHeader는 상태 코드를 캡처하고 기본 WriteHeader를 호출합니다.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// isOriginAllowed checks if an origin is in the allowed list.
// isOriginAllowed는 오리진이 허용 목록에 있는지 확인합니다.
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowed := range allowedOrigins {
		if allowed == "*" || allowed == origin {
			return true
		}
	}
	return false
}

// joinStrings joins a slice of strings with a separator.
// joinStrings는 구분자로 문자열 슬라이스를 결합합니다.
func joinStrings(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}

// RequestID returns a middleware that adds a unique request ID to each request.
// RequestID는 각 요청에 고유한 요청 ID를 추가하는 미들웨어를 반환합니다.
//
// The request ID is added to the request context and response headers.
// 요청 ID는 요청 컨텍스트와 응답 헤더에 추가됩니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.RequestID())
//
//	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
//	    requestID := r.Context().Value("request_id").(string)
//	    fmt.Fprintf(w, "Request ID: %s", requestID)
//	})
func RequestID() MiddlewareFunc {
	return RequestIDWithConfig(RequestIDConfig{
		Header: "X-Request-ID",
	})
}

// RequestIDConfig defines the configuration for RequestID middleware.
// RequestIDConfig는 RequestID 미들웨어의 설정을 정의합니다.
type RequestIDConfig struct {
	// Header is the name of the header to store the request ID.
	// Header는 요청 ID를 저장할 헤더의 이름입니다.
	// Default: "X-Request-ID"
	Header string

	// Generator is the function to generate request IDs.
	// Generator는 요청 ID를 생성하는 함수입니다.
	// Default: generateRequestID (random 16-byte hex string)
	Generator func() string
}

// RequestIDWithConfig returns a RequestID middleware with custom configuration.
// RequestIDWithConfig는 커스텀 설정으로 RequestID 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.RequestIDWithConfig(websvrutil.RequestIDConfig{
//	    Header: "X-Custom-Request-ID",
//	    Generator: func() string {
//	        return uuid.New().String()
//	    },
//	}))
func RequestIDWithConfig(config RequestIDConfig) MiddlewareFunc {
	// Set defaults
	// 기본값 설정
	if config.Header == "" {
		config.Header = "X-Request-ID"
	}
	if config.Generator == nil {
		config.Generator = generateRequestID
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if request ID already exists in header
			// 헤더에 요청 ID가 이미 존재하는지 확인
			requestID := r.Header.Get(config.Header)
			if requestID == "" {
				// Generate new request ID
				// 새 요청 ID 생성
				requestID = config.Generator()
			}

			// Add request ID to response header
			// 응답 헤더에 요청 ID 추가
			w.Header().Set(config.Header, requestID)

			// Add request ID to request context
			// 요청 컨텍스트에 요청 ID 추가
			ctx := context.WithValue(r.Context(), "request_id", requestID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// generateRequestID generates a random 16-byte hex string.
// generateRequestID는 무작위 16바이트 16진수 문자열을 생성합니다.
func generateRequestID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Timeout returns a middleware that enforces a request timeout.
// Timeout은 요청 타임아웃을 적용하는 미들웨어를 반환합니다.
//
// If the handler does not complete within the timeout, a 503 Service Unavailable response is sent.
// 핸들러가 타임아웃 내에 완료되지 않으면 503 Service Unavailable 응답이 전송됩니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.Timeout(5 * time.Second))
func Timeout(timeout time.Duration) MiddlewareFunc {
	return TimeoutWithConfig(TimeoutConfig{
		Timeout: timeout,
	})
}

// TimeoutConfig defines the configuration for Timeout middleware.
// TimeoutConfig는 Timeout 미들웨어의 설정을 정의합니다.
type TimeoutConfig struct {
	// Timeout is the maximum duration for the request.
	// Timeout은 요청의 최대 지속 시간입니다.
	// Default: 30 seconds
	Timeout time.Duration

	// Message is the error message sent on timeout.
	// Message는 타임아웃 시 전송되는 에러 메시지입니다.
	// Default: "Service Unavailable"
	Message string
}

// TimeoutWithConfig returns a Timeout middleware with custom configuration.
// TimeoutWithConfig는 커스텀 설정으로 Timeout 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.TimeoutWithConfig(websvrutil.TimeoutConfig{
//	    Timeout: 10 * time.Second,
//	    Message: "Request timed out",
//	}))
func TimeoutWithConfig(config TimeoutConfig) MiddlewareFunc {
	// Set defaults
	// 기본값 설정
	if config.Timeout <= 0 {
		config.Timeout = 30 * time.Second
	}
	if config.Message == "" {
		config.Message = "Service Unavailable"
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a context with timeout
			// 타임아웃이 있는 컨텍스트 생성
			ctx, cancel := context.WithTimeout(r.Context(), config.Timeout)
			defer cancel()

			// Replace request context
			// 요청 컨텍스트 교체
			r = r.WithContext(ctx)

			// Use http.TimeoutHandler
			// http.TimeoutHandler 사용
			http.TimeoutHandler(next, config.Timeout, config.Message).ServeHTTP(w, r)
		})
	}
}

// BasicAuth returns a middleware that enforces HTTP Basic Authentication.
// BasicAuth는 HTTP Basic Authentication을 적용하는 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.BasicAuth("admin", "password"))
func BasicAuth(username, password string) MiddlewareFunc {
	return BasicAuthWithConfig(BasicAuthConfig{
		Username: username,
		Password: password,
	})
}

// BasicAuthConfig defines the configuration for BasicAuth middleware.
// BasicAuthConfig는 BasicAuth 미들웨어의 설정을 정의합니다.
type BasicAuthConfig struct {
	// Username is the required username.
	// Username은 필수 사용자 이름입니다.
	Username string

	// Password is the required password.
	// Password는 필수 비밀번호입니다.
	Password string

	// Realm is the authentication realm.
	// Realm은 인증 영역입니다.
	// Default: "Restricted"
	Realm string

	// Validator is a custom validation function.
	// Validator는 커스텀 검증 함수입니다.
	// If provided, Username and Password are ignored.
	// 제공되면 Username과 Password는 무시됩니다.
	Validator func(username, password string) bool
}

// BasicAuthWithConfig returns a BasicAuth middleware with custom configuration.
// BasicAuthWithConfig는 커스텀 설정으로 BasicAuth 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.BasicAuthWithConfig(websvrutil.BasicAuthConfig{
//	    Validator: func(username, password string) bool {
//	        return username == "admin" && password == "secret"
//	    },
//	    Realm: "Admin Area",
//	}))
func BasicAuthWithConfig(config BasicAuthConfig) MiddlewareFunc {
	// Set defaults
	// 기본값 설정
	if config.Realm == "" {
		config.Realm = "Restricted"
	}

	// Create validator function
	// 검증자 함수 생성
	var validator func(string, string) bool
	if config.Validator != nil {
		validator = config.Validator
	} else {
		// Use default validator (constant-time comparison)
		// 기본 검증자 사용 (상수 시간 비교)
		validator = func(username, password string) bool {
			usernameMatch := subtle.ConstantTimeCompare([]byte(username), []byte(config.Username)) == 1
			passwordMatch := subtle.ConstantTimeCompare([]byte(password), []byte(config.Password)) == 1
			return usernameMatch && passwordMatch
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get credentials from Authorization header
			// Authorization 헤더에서 자격 증명 가져오기
			username, password, ok := r.BasicAuth()
			if !ok || !validator(username, password) {
				// Send 401 Unauthorized with WWW-Authenticate header
				// WWW-Authenticate 헤더와 함께 401 Unauthorized 전송
				w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, config.Realm))
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Store username in context for later use
			// 나중에 사용하기 위해 컨텍스트에 사용자 이름 저장
			ctx := context.WithValue(r.Context(), "auth_username", username)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// RateLimiter returns a middleware that limits the number of requests per client.
// RateLimiter는 클라이언트당 요청 수를 제한하는 미들웨어를 반환합니다.
//
// Uses a simple token bucket algorithm with IP-based rate limiting.
// IP 기반 rate limiting과 함께 간단한 토큰 버킷 알고리즘을 사용합니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.RateLimiter(100, time.Minute)) // 100 requests per minute
func RateLimiter(requests int, window time.Duration) MiddlewareFunc {
	return RateLimiterWithConfig(RateLimiterConfig{
		Requests: requests,
		Window:   window,
	})
}

// RateLimiterConfig defines the configuration for RateLimiter middleware.
// RateLimiterConfig는 RateLimiter 미들웨어의 설정을 정의합니다.
type RateLimiterConfig struct {
	// Requests is the maximum number of requests allowed per window.
	// Requests는 윈도우당 허용되는 최대 요청 수입니다.
	// Default: 100
	Requests int

	// Window is the time window for rate limiting.
	// Window는 rate limiting을 위한 시간 윈도우입니다.
	// Default: 1 minute
	Window time.Duration

	// KeyFunc is the function to extract the rate limit key from the request.
	// KeyFunc는 요청에서 rate limit 키를 추출하는 함수입니다.
	// Default: uses client IP address
	KeyFunc func(r *http.Request) string
}

// rateLimitEntry stores rate limit information for a client.
// rateLimitEntry는 클라이언트의 rate limit 정보를 저장합니다.
type rateLimitEntry struct {
	count      int
	resetTime  time.Time
	mu         sync.Mutex
}

// RateLimiterWithConfig returns a RateLimiter middleware with custom configuration.
// RateLimiterWithConfig는 커스텀 설정으로 RateLimiter 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.RateLimiterWithConfig(websvrutil.RateLimiterConfig{
//	    Requests: 50,
//	    Window: 30 * time.Second,
//	    KeyFunc: func(r *http.Request) string {
//	        return r.Header.Get("X-API-Key")
//	    },
//	}))
func RateLimiterWithConfig(config RateLimiterConfig) MiddlewareFunc {
	// Set defaults
	// 기본값 설정
	if config.Requests <= 0 {
		config.Requests = 100
	}
	if config.Window <= 0 {
		config.Window = time.Minute
	}
	if config.KeyFunc == nil {
		config.KeyFunc = func(r *http.Request) string {
			// Extract IP from RemoteAddr
			// RemoteAddr에서 IP 추출
			ip := r.RemoteAddr
			if idx := strings.LastIndex(ip, ":"); idx != -1 {
				ip = ip[:idx]
			}
			return ip
		}
	}

	// Create rate limit store
	// rate limit 저장소 생성
	store := make(map[string]*rateLimitEntry)
	var storeMu sync.RWMutex

	// Cleanup goroutine to remove expired entries
	// 만료된 항목을 제거하는 정리 고루틴
	go func() {
		ticker := time.NewTicker(config.Window)
		defer ticker.Stop()
		for range ticker.C {
			storeMu.Lock()
			now := time.Now()
			for key, entry := range store {
				entry.mu.Lock()
				if now.After(entry.resetTime) {
					delete(store, key)
				}
				entry.mu.Unlock()
			}
			storeMu.Unlock()
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := config.KeyFunc(r)
			now := time.Now()

			// Get or create entry
			// 항목 가져오기 또는 생성
			storeMu.Lock()
			entry, exists := store[key]
			if !exists {
				entry = &rateLimitEntry{
					count:     0,
					resetTime: now.Add(config.Window),
				}
				store[key] = entry
			}
			storeMu.Unlock()

			// Check rate limit
			// rate limit 확인
			entry.mu.Lock()
			if now.After(entry.resetTime) {
				// Reset window
				// 윈도우 리셋
				entry.count = 0
				entry.resetTime = now.Add(config.Window)
			}

			if entry.count >= config.Requests {
				entry.mu.Unlock()
				// Rate limit exceeded
				// Rate limit 초과
				w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", config.Requests))
				w.Header().Set("X-RateLimit-Remaining", "0")
				w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", entry.resetTime.Unix()))
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			entry.count++
			remaining := config.Requests - entry.count
			entry.mu.Unlock()

			// Set rate limit headers
			// rate limit 헤더 설정
			w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", config.Requests))
			w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
			w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", entry.resetTime.Unix()))

			next.ServeHTTP(w, r)
		})
	}
}

// Compression returns a middleware that compresses HTTP responses using gzip.
// Compression은 gzip을 사용하여 HTTP 응답을 압축하는 미들웨어를 반환합니다.
//
// Automatically compresses responses when client supports gzip (Accept-Encoding: gzip).
// 클라이언트가 gzip을 지원할 때 자동으로 응답을 압축합니다 (Accept-Encoding: gzip).
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.Compression())
func Compression() MiddlewareFunc {
	return CompressionWithConfig(CompressionConfig{
		Level: gzip.DefaultCompression,
	})
}

// CompressionConfig defines the configuration for Compression middleware.
// CompressionConfig는 Compression 미들웨어의 설정을 정의합니다.
type CompressionConfig struct {
	// Level is the gzip compression level.
	// Level은 gzip 압축 레벨입니다.
	// Valid values: -1 (default), 0 (no compression), 1 (best speed) to 9 (best compression)
	// 유효한 값: -1 (기본), 0 (압축 없음), 1 (최고 속도) ~ 9 (최고 압축)
	Level int

	// MinLength is the minimum response size to compress (in bytes).
	// MinLength는 압축할 최소 응답 크기입니다 (바이트).
	// Default: 1024 (1KB)
	MinLength int
}

// gzipResponseWriter wraps http.ResponseWriter to support gzip compression.
// gzipResponseWriter는 gzip 압축을 지원하기 위해 http.ResponseWriter를 래핑합니다.
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
	statusCode int
}

func (w *gzipResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// CompressionWithConfig returns a Compression middleware with custom configuration.
// CompressionWithConfig는 커스텀 설정으로 Compression 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.CompressionWithConfig(websvrutil.CompressionConfig{
//	    Level: gzip.BestCompression,
//	    MinLength: 2048, // 2KB
//	}))
func CompressionWithConfig(config CompressionConfig) MiddlewareFunc {
	// Set defaults
	// 기본값 설정
	if config.Level == 0 {
		config.Level = gzip.DefaultCompression
	}
	if config.MinLength <= 0 {
		config.MinLength = 1024 // 1KB
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if client supports gzip
			// 클라이언트가 gzip을 지원하는지 확인
			if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				next.ServeHTTP(w, r)
				return
			}

			// Create gzip writer
			// gzip writer 생성
			gz, err := gzip.NewWriterLevel(w, config.Level)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			defer gz.Close()

			// Set Content-Encoding header
			// Content-Encoding 헤더 설정
			w.Header().Set("Content-Encoding", "gzip")
			w.Header().Del("Content-Length") // Content-Length is not accurate after compression

			// Wrap response writer
			// response writer 래핑
			gzw := &gzipResponseWriter{
				Writer:         gz,
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			next.ServeHTTP(gzw, r)
		})
	}
}

// SecureHeaders returns a middleware that adds security-related HTTP headers.
// SecureHeaders는 보안 관련 HTTP 헤더를 추가하는 미들웨어를 반환합니다.
//
// Adds headers like X-Frame-Options, X-Content-Type-Options, X-XSS-Protection, etc.
// X-Frame-Options, X-Content-Type-Options, X-XSS-Protection 등의 헤더를 추가합니다.
//
// Example / 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.SecureHeaders())
func SecureHeaders() MiddlewareFunc {
	return SecureHeadersWithConfig(SecureHeadersConfig{})
}

// SecureHeadersConfig defines the configuration for SecureHeaders middleware.
// SecureHeadersConfig는 SecureHeaders 미들웨어의 설정을 정의합니다.
type SecureHeadersConfig struct {
	// XFrameOptions provides clickjacking protection.
	// XFrameOptions는 클릭재킹 보호를 제공합니다.
	// Default: "SAMEORIGIN"
	XFrameOptions string

	// XContentTypeOptions prevents MIME type sniffing.
	// XContentTypeOptions는 MIME 타입 스니핑을 방지합니다.
	// Default: "nosniff"
	XContentTypeOptions string

	// XXSSProtection enables XSS filter.
	// XXSSProtection은 XSS 필터를 활성화합니다.
	// Default: "1; mode=block"
	XXSSProtection string

	// ContentSecurityPolicy defines CSP policy.
	// ContentSecurityPolicy는 CSP 정책을 정의합니다.
	// Default: "" (not set)
	ContentSecurityPolicy string

	// StrictTransportSecurity enforces HTTPS.
	// StrictTransportSecurity는 HTTPS를 강제합니다.
	// Default: "max-age=31536000; includeSubDomains"
	StrictTransportSecurity string

	// ReferrerPolicy controls referrer information.
	// ReferrerPolicy는 리퍼러 정보를 제어합니다.
	// Default: "strict-origin-when-cross-origin"
	ReferrerPolicy string
}

// SecureHeadersWithConfig returns a SecureHeaders middleware with custom configuration.
// SecureHeadersWithConfig는 커스텀 설정으로 SecureHeaders 미들웨어를 반환합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.SecureHeadersWithConfig(websvrutil.SecureHeadersConfig{
//	    XFrameOptions: "DENY",
//	    ContentSecurityPolicy: "default-src 'self'",
//	}))
func SecureHeadersWithConfig(config SecureHeadersConfig) MiddlewareFunc {
	// Set defaults
	// 기본값 설정
	if config.XFrameOptions == "" {
		config.XFrameOptions = "SAMEORIGIN"
	}
	if config.XContentTypeOptions == "" {
		config.XContentTypeOptions = "nosniff"
	}
	if config.XXSSProtection == "" {
		config.XXSSProtection = "1; mode=block"
	}
	if config.StrictTransportSecurity == "" {
		config.StrictTransportSecurity = "max-age=31536000; includeSubDomains"
	}
	if config.ReferrerPolicy == "" {
		config.ReferrerPolicy = "strict-origin-when-cross-origin"
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set security headers
			// 보안 헤더 설정
			w.Header().Set("X-Frame-Options", config.XFrameOptions)
			w.Header().Set("X-Content-Type-Options", config.XContentTypeOptions)
			w.Header().Set("X-XSS-Protection", config.XXSSProtection)
			w.Header().Set("Referrer-Policy", config.ReferrerPolicy)

			// Set HSTS only for HTTPS
			// HTTPS에만 HSTS 설정
			if r.TLS != nil {
				w.Header().Set("Strict-Transport-Security", config.StrictTransportSecurity)
			}

			// Set CSP if configured
			// 설정된 경우 CSP 설정
			if config.ContentSecurityPolicy != "" {
				w.Header().Set("Content-Security-Policy", config.ContentSecurityPolicy)
			}

			next.ServeHTTP(w, r)
		})
	}
}
