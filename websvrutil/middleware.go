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
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

// middleware.go provides HTTP middleware components for the websvrutil framework.
//
// This file implements a comprehensive collection of middleware functions that can be
// chained together to add cross-cutting concerns to HTTP request handling:
//
// Error Recovery:
//   - Recovery(): Catches panics, logs stack traces, sends 500 responses
//   - RecoveryWithConfig(): Customizable panic recovery with LogFunc option
//
// Request Logging:
//   - Logger(): Logs method, path, status, and latency for each request
//   - LoggerWithConfig(): Customizable logging with user-defined LogFunc
//
// Cross-Origin Resource Sharing (CORS):
//   - CORS(): Default CORS with Allow-Origin: *
//   - CORSWithConfig(): Fine-grained control over origins, methods, headers
//     Supports wildcard origins, preflight OPTIONS handling, credentials
//
// Request Identification:
//   - RequestID(): Generates unique request IDs (16-byte hex)
//   - RequestIDWithConfig(): Custom ID generation and header name
//     Useful for distributed tracing and log correlation
//
// Timeout Management:
//   - Timeout(): Enforces request timeout with context cancellation
//   - TimeoutWithConfig(): Custom timeout and error handler
//     Prevents long-running requests from exhausting resources
//
// Authentication:
//   - BasicAuth(): HTTP Basic Authentication with username/password
//   - BasicAuthWithConfig(): Customizable validator function and realm
//     Uses constant-time comparison to prevent timing attacks
//
// Rate Limiting:
//   - RateLimiter(): Token bucket rate limiting per client IP
//   - RateLimiterWithConfig(): Custom rate, window, and key function
//     Prevents abuse and DoS attacks with sliding window algorithm
//     Thread-safe with mutex-protected rate limit tracking
//
// Response Compression:
//   - Compression(): Gzip compression for responses
//   - CompressionWithConfig(): Custom compression level and min size
//     Checks Accept-Encoding header, reduces bandwidth usage
//
// Security Headers:
//   - SecureHeaders(): Adds common security headers (CSP, X-Frame-Options, etc.)
//   - SecureHeadersWithConfig(): Customizable security policy
//     Protects against XSS, clickjacking, MIME sniffing
//
// Body Size Limiting:
//   - BodyLimit(): Enforces maximum request body size
//   - BodyLimitWithConfig(): Custom size limit and error handler
//     Prevents memory exhaustion from large payloads
//
// Static File Serving:
//   - Static(): Serves files from filesystem directory
//   - StaticWithConfig(): Custom root, index files, browse option
//
// URL Redirection:
//   - Redirect(): Redirects all requests to specified URL
//   - RedirectWithConfig(): Custom status code (301, 302, 307, 308)
//   - HTTPSRedirect(): Redirects HTTP to HTTPS
//   - WWWRedirect(): Adds or removes www prefix
//
// All middleware follow the standard MiddlewareFunc signature:
//
//	type MiddlewareFunc func(http.Handler) http.Handler
//
// Middleware can be chained using App.Use():
//
//	app := New()
//	app.Use(Recovery())
//	app.Use(Logger())
//	app.Use(CORS())
//	app.Use(RateLimiter(100, time.Minute))
//
// Performance Considerations:
// - Middleware is executed in registration order for each request
// - Rate limiting uses efficient token bucket with cleanup goroutine
// - Compression checks content size before applying gzip
// - Request ID generation uses crypto/rand for uniqueness
//
// Security Features:
// - Panic recovery prevents information leakage
// - Basic auth uses constant-time comparison (subtle.ConstantTimeCompare)
// - Rate limiting prevents DoS attacks
// - Secure headers protect against common web vulnerabilities
// - Body limiting prevents memory exhaustion
//
// middleware.go는 websvrutil 프레임워크를 위한 HTTP 미들웨어 컴포넌트를 제공합니다.
//
// 이 파일은 HTTP 요청 처리에 횡단 관심사를 추가하기 위해 함께 연결될 수 있는
// 포괄적인 미들웨어 함수 모음을 구현합니다:
//
// 에러 복구:
//   - Recovery(): 패닉을 잡고, 스택 트레이스를 로깅하고, 500 응답을 전송
//   - RecoveryWithConfig(): LogFunc 옵션이 있는 사용자 정의 가능한 패닉 복구
//
// 요청 로깅:
//   - Logger(): 각 요청의 메서드, 경로, 상태, 지연시간 로깅
//   - LoggerWithConfig(): 사용자 정의 LogFunc이 있는 커스터마이징 가능한 로깅
//
// Cross-Origin Resource Sharing (CORS):
//   - CORS(): Allow-Origin: *가 있는 기본 CORS
//   - CORSWithConfig(): 출처, 메서드, 헤더에 대한 세밀한 제어
//     와일드카드 출처, preflight OPTIONS 처리, 자격 증명 지원
//
// 요청 식별:
//   - RequestID(): 고유 요청 ID 생성 (16바이트 hex)
//   - RequestIDWithConfig(): 사용자 정의 ID 생성 및 헤더 이름
//     분산 추적 및 로그 상관관계에 유용
//
// 타임아웃 관리:
//   - Timeout(): 컨텍스트 취소와 함께 요청 타임아웃 강제
//   - TimeoutWithConfig(): 사용자 정의 타임아웃 및 에러 핸들러
//     장시간 실행 요청이 리소스를 고갈시키는 것을 방지
//
// 인증:
//   - BasicAuth(): 사용자 이름/비밀번호로 HTTP 기본 인증
//   - BasicAuthWithConfig(): 사용자 정의 가능한 검증 함수 및 realm
//     타이밍 공격 방지를 위해 상수 시간 비교 사용
//
// 속도 제한:
//   - RateLimiter(): 클라이언트 IP당 토큰 버킷 속도 제한
//   - RateLimiterWithConfig(): 사용자 정의 속도, 윈도우, 키 함수
//     슬라이딩 윈도우 알고리즘으로 남용 및 DoS 공격 방지
//     뮤텍스로 보호된 속도 제한 추적으로 스레드 안전
//
// 응답 압축:
//   - Compression(): 응답에 대한 Gzip 압축
//   - CompressionWithConfig(): 사용자 정의 압축 수준 및 최소 크기
//     Accept-Encoding 헤더 확인, 대역폭 사용량 감소
//
// 보안 헤더:
//   - SecureHeaders(): 일반 보안 헤더 추가 (CSP, X-Frame-Options 등)
//   - SecureHeadersWithConfig(): 사용자 정의 가능한 보안 정책
//     XSS, 클릭재킹, MIME 스니핑으로부터 보호
//
// 본문 크기 제한:
//   - BodyLimit(): 최대 요청 본문 크기 강제
//   - BodyLimitWithConfig(): 사용자 정의 크기 제한 및 에러 핸들러
//     대용량 페이로드로 인한 메모리 고갈 방지
//
// 정적 파일 서빙:
//   - Static(): 파일시스템 디렉토리에서 파일 서빙
//   - StaticWithConfig(): 사용자 정의 루트, 인덱스 파일, 브라우즈 옵션
//
// URL 리다이렉션:
//   - Redirect(): 모든 요청을 지정된 URL로 리다이렉트
//   - RedirectWithConfig(): 사용자 정의 상태 코드 (301, 302, 307, 308)
//   - HTTPSRedirect(): HTTP를 HTTPS로 리다이렉트
//   - WWWRedirect(): www 접두사 추가 또는 제거
//
// 모든 미들웨어는 표준 MiddlewareFunc 시그니처를 따릅니다:
//
//	type MiddlewareFunc func(http.Handler) http.Handler
//
// 미들웨어는 App.Use()를 사용하여 연결할 수 있습니다:
//
//	app := New()
//	app.Use(Recovery())
//	app.Use(Logger())
//	app.Use(CORS())
//	app.Use(RateLimiter(100, time.Minute))
//
// 성능 고려사항:
// - 미들웨어는 각 요청에 대해 등록 순서대로 실행됨
// - 속도 제한은 정리 고루틴이 있는 효율적인 토큰 버킷 사용
// - 압축은 gzip 적용 전 콘텐츠 크기 확인
// - 요청 ID 생성은 고유성을 위해 crypto/rand 사용
//
// 보안 기능:
// - 패닉 복구는 정보 유출 방지
// - 기본 인증은 상수 시간 비교 사용 (subtle.ConstantTimeCompare)
// - 속도 제한은 DoS 공격 방지
// - 보안 헤더는 일반적인 웹 취약점으로부터 보호
// - 본문 제한은 메모리 고갈 방지

// Recovery returns middleware that wraps handlers with panic recovery and sends a safe HTTP 500 response when a panic occurs.
// Recovery는 핸들러를 패닉 복구 로직으로 감싸 패닉이 발생하면 안전한 HTTP 500 응답을 보냅니다.
//
// Key behaviors:
// - Wraps handler execution with defer/recover to prevent server crashes.
// - Logs the panic value together with the captured stack trace for diagnostics.
// - Sends a generic 500 response so internal details never leak to clients.
//
// 주요 동작:
// - defer/recover로 핸들러 실행을 감싸 서버 전체가 중단되는 것을 방지합니다.
// - 패닉 값과 캡처된 스택 트레이스를 함께 기록해 진단에 활용합니다.
// - 내부 정보가 노출되지 않도록 클라이언트에는 일반적인 500 응답만 전달합니다.
//
// Example:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.Recovery())
//	app.GET("/panic", func(w http.ResponseWriter, r *http.Request) {
//	    var user *User
//	    fmt.Fprintf(w, "Name: %s", user.Name) // triggers panic for demonstration
//	})
//
// 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.Recovery())
//	app.GET("/panic", func(w http.ResponseWriter, r *http.Request) {
//	    var user *User
//	    fmt.Fprintf(w, "Name: %s", user.Name) // 데모를 위해 패닉을 유도합니다
//	})
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
// Example
// 예제:
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

// Logger returns middleware that records method, path, status code, and latency for each HTTP request.
// Logger는 각 HTTP 요청의 메서드, 경로, 상태 코드, 지연 시간을 기록하는 미들웨어를 제공합니다.
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
			duration := time.Since(start) // 요청 처리에 걸린 시간
			log.Printf("%s %s %d %v", r.Method, r.URL.Path, wrapper.statusCode, duration)
		})
	}
}

// LoggerWithConfig returns middleware that applies a user-defined logging function per request.
// LoggerWithConfig는 요청마다 사용자 정의 로깅 함수를 실행하는 미들웨어를 제공합니다.
//
// Example / 예제:
//
//	app.Use(websvrutil.LoggerWithConfig(websvrutil.LoggerConfig{
//	    LogFunc: func(method, path string, status int, duration time.Duration) {
//	        log.Printf("[ACCESS] %s %s %d %v", method, path, status, duration)
//	    },
//	}))
func LoggerWithConfig(config LoggerConfig) MiddlewareFunc {
	// Set defaults if not provided / 제공되지 않은 경우 기본값을 설정합니다.
	if config.LogFunc == nil {
		config.LogFunc = func(method, path string, status int, duration time.Duration) {
			log.Printf("%s %s %d %v", method, path, status, duration)
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK} // 기본 상태 코드를 OK로 초기화

			next.ServeHTTP(wrapper, r) // 실제 핸들러 실행

			duration := time.Since(start)                                      // 요청 처리에 걸린 시간
			config.LogFunc(r.Method, r.URL.Path, wrapper.statusCode, duration) // 사용자 정의 로거 호출
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
	count     int
	resetTime time.Time
	mu        sync.Mutex
}

// RateLimiterWithConfig returns a RateLimiter middleware with custom configuration.
// RateLimiterWithConfig는 커스텀 설정으로 RateLimiter 미들웨어를 반환합니다.
//
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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

// BodyLimitConfig defines the configuration for BodyLimit middleware.
// BodyLimitConfig는 BodyLimit 미들웨어의 설정을 정의합니다.
type BodyLimitConfig struct {
	// MaxBytes is the maximum allowed request body size in bytes
	// MaxBytes는 허용되는 최대 요청 본문 크기(바이트)입니다
	MaxBytes int64
}

// BodyLimit returns a middleware that limits the maximum request body size.
// BodyLimit는 최대 요청 본문 크기를 제한하는 미들웨어를 반환합니다.
//
// Default limit is 10MB.
// 기본 제한은 10MB입니다.
//
// Example
// 예제:
//
// // Limit request body to 5MB
// 요청 본문을 5MB로 제한
//
//	server.Use(BodyLimit(5 * 1024 * 1024))
func BodyLimit(maxBytes int64) MiddlewareFunc {
	return BodyLimitWithConfig(BodyLimitConfig{
		MaxBytes: maxBytes,
	})
}

// BodyLimitWithConfig returns a middleware with custom configuration.
// BodyLimitWithConfig는 사용자 정의 설정으로 미들웨어를 반환합니다.
func BodyLimitWithConfig(config BodyLimitConfig) MiddlewareFunc {
	if config.MaxBytes <= 0 {
		config.MaxBytes = DefaultMaxBodySize
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Limit the request body size
			// 요청 본문 크기 제한
			r.Body = http.MaxBytesReader(w, r.Body, config.MaxBytes)

			next.ServeHTTP(w, r)
		})
	}
}

// StaticConfig defines the configuration for Static middleware.
// StaticConfig는 Static 미들웨어의 설정을 정의합니다.
type StaticConfig struct {
	// Root is the root directory to serve files from
	// Root는 파일을 제공할 루트 디렉토리입니다
	Root string

	// Index is the index file to serve (default: "index.html")
	// Index는 제공할 인덱스 파일입니다 (기본값: "index.html")
	Index string

	// Browse enables directory browsing
	// Browse는 디렉토리 탐색을 활성화합니다
	Browse bool
}

// Static returns a middleware that serves static files from the specified directory.
// Static는 지정된 디렉토리에서 정적 파일을 제공하는 미들웨어를 반환합니다.
//
// Example
// 예제:
//
// // Serve static files from "./public" directory
// "./public" 디렉토리에서 정적 파일 제공
//
//	server.Use(Static("./public"))
func Static(root string) MiddlewareFunc {
	return StaticWithConfig(StaticConfig{
		Root: root,
	})
}

// StaticWithConfig returns a middleware with custom configuration.
// StaticWithConfig는 사용자 정의 설정으로 미들웨어를 반환합니다.
func StaticWithConfig(config StaticConfig) MiddlewareFunc {
	if config.Root == "" {
		config.Root = "."
	}
	if config.Index == "" {
		config.Index = "index.html"
	}

	fileServer := http.FileServer(http.Dir(config.Root))

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if file exists
			// 파일이 존재하는지 확인
			filePath := config.Root + r.URL.Path
			info, err := os.Stat(filePath)

			if err == nil {
				// If it's a directory, serve index file or browse
				// 디렉토리인 경우 인덱스 파일 제공 또는 탐색
				if info.IsDir() {
					indexPath := filePath + "/" + config.Index
					if _, err := os.Stat(indexPath); err == nil {
						fileServer.ServeHTTP(w, r)
						return
					}
					if !config.Browse {
						next.ServeHTTP(w, r)
						return
					}
				}
				// Serve the file
				// 파일 제공
				fileServer.ServeHTTP(w, r)
				return
			}

			// File not found, pass to next handler
			// 파일을 찾을 수 없으면 다음 핸들러로 전달
			next.ServeHTTP(w, r)
		})
	}
}

// RedirectConfig defines the configuration for Redirect middleware.
// RedirectConfig는 Redirect 미들웨어의 설정을 정의합니다.
type RedirectConfig struct {
	// Code is the HTTP status code for redirect (default: 301 Moved Permanently)
	// Code는 리디렉션을 위한 HTTP 상태 코드입니다 (기본값: 301 Moved Permanently)
	Code int

	// To is the destination URL
	// To는 대상 URL입니다
	To string
}

// Redirect returns a middleware that redirects all requests to the specified URL.
// Redirect는 모든 요청을 지정된 URL로 리디렉션하는 미들웨어를 반환합니다.
//
// Uses 301 Moved Permanently by default.
// 기본적으로 301 Moved Permanently를 사용합니다.
//
// Example
// 예제:
//
// // Redirect all HTTP to HTTPS
// 모든 HTTP를 HTTPS로 리디렉션
//
//	httpServer.Use(Redirect("https://example.com"))
func Redirect(to string) MiddlewareFunc {
	return RedirectWithConfig(RedirectConfig{
		To: to,
	})
}

// RedirectWithConfig returns a middleware with custom configuration.
// RedirectWithConfig는 사용자 정의 설정으로 미들웨어를 반환합니다.
func RedirectWithConfig(config RedirectConfig) MiddlewareFunc {
	if config.Code == 0 {
		config.Code = http.StatusMovedPermanently // 301
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Redirect to destination
			// 대상으로 리디렉션
			http.Redirect(w, r, config.To, config.Code)
		})
	}
}

// HTTPSRedirect returns a middleware that redirects HTTP requests to HTTPS.
// HTTPSRedirect는 HTTP 요청을 HTTPS로 리디렉션하는 미들웨어를 반환합니다.
//
// Example
// 예제:
//
// // Redirect all HTTP to HTTPS
// 모든 HTTP를 HTTPS로 리디렉션
//
//	httpServer.Use(HTTPSRedirect())
func HTTPSRedirect() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// If not HTTPS, redirect to HTTPS
			// HTTPS가 아니면 HTTPS로 리디렉션
			if r.TLS == nil && r.Header.Get("X-Forwarded-Proto") != "https" {
				host := r.Host
				if host == "" {
					host = "localhost"
				}
				url := "https://" + host + r.URL.Path
				http.Redirect(w, r, url, http.StatusMovedPermanently)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// WWWRedirect returns a middleware that redirects non-www to www or vice versa.
// WWWRedirect는 non-www를 www로 또는 그 반대로 리디렉션하는 미들웨어를 반환합니다.
//
// Example
// 예제:
//
// // Redirect to www version
// www 버전으로 리디렉션
//
//	server.Use(WWWRedirect(true))
func WWWRedirect(addWWW bool) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host := r.Host
			hasWWW := len(host) > 4 && host[:4] == "www."

			if addWWW && !hasWWW {
				// Add www prefix
				// www 접두사 추가
				scheme := "http"
				if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
					scheme = "https"
				}
				url := scheme + "://www." + host + r.URL.Path
				http.Redirect(w, r, url, http.StatusMovedPermanently)
				return
			} else if !addWWW && hasWWW {
				// Remove www prefix
				// www 접두사 제거
				scheme := "http"
				if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
					scheme = "https"
				}
				url := scheme + "://" + host[4:] + r.URL.Path
				http.Redirect(w, r, url, http.StatusMovedPermanently)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
