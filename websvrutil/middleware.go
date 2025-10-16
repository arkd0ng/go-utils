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

// Recovery returns a middleware that recovers from panics.
// Recovery는 패닉에서 복구하는 미들웨어를 반환합니다.
//
// Panic recovery mechanism / 패닉 복구 메커니즘:
//   - Wraps handler execution in defer/recover block
//   - defer/recover 블록으로 핸들러 실행 래핑
//   - Prevents server crashes from unhandled panics
//   - 처리되지 않은 패닉으로 인한 서버 충돌 방지
//   - Logs panic details with full stack trace
//   - 전체 스택 트레이스와 함께 패닉 세부정보 로깅
//   - Returns 500 Internal Server Error to client
//   - 클라이언트에 500 Internal Server Error 반환
//
// Why panic recovery is critical / 패닉 복구가 중요한 이유:
//   - Go's panic mechanism immediately terminates the goroutine
//   - Go의 패닉 메커니즘은 즉시 고루틴을 종료함
//   - Without recovery, a single panic in any handler kills the entire server
//   - 복구 없이는 모든 핸들러의 단일 패닉이 전체 서버를 중단시킴
//   - Recovery middleware ensures server stability and availability
//   - 복구 미들웨어는 서버 안정성과 가용성을 보장함
//   - Allows graceful degradation instead of catastrophic failure
//   - 재앙적 실패 대신 우아한 성능 저하 허용
//
// What gets logged / 로깅되는 내용:
//   - Panic value (error message, value, or custom type)
//   - 패닉 값 (에러 메시지, 값 또는 커스텀 타입)
//   - Full goroutine stack trace using runtime/debug.Stack()
//   - runtime/debug.Stack()을 사용한 전체 고루틴 스택 트레이스
//   - File paths, line numbers, and function names
//   - 파일 경로, 줄 번호 및 함수 이름
//
// Stack trace format example / 스택 트레이스 형식 예제:
//   PANIC: runtime error: index out of range [5] with length 3
//   goroutine 123 [running]:
//   runtime/debug.Stack()
//       /usr/local/go/src/runtime/debug/stack.go:24 +0x65
//   main.(*App).Recovery.func1.1()
//       /app/middleware.go:37 +0x89
//   ...
//
// Response to client / 클라이언트에 대한 응답:
//   - HTTP 500 Internal Server Error
//   - Generic "Internal Server Error" message (no stack trace exposed)
//   - 일반적인 "Internal Server Error" 메시지 (스택 트레이스 노출 없음)
//   - Prevents information disclosure to attackers
//   - 공격자에게 정보 공개 방지
//   - Detailed error only in server logs for debugging
//   - 디버깅을 위한 서버 로그에만 상세 에러 기록
//
// Thread safety / 스레드 안전성:
//   - Each request runs in separate goroutine with own defer/recover
//   - 각 요청은 자체 defer/recover를 가진 별도 고루틴에서 실행
//   - Panic in one request does NOT affect other concurrent requests
//   - 하나의 요청에서 패닉이 발생해도 다른 동시 요청에 영향 없음
//   - Server continues serving other requests normally
//   - 서버는 다른 요청을 정상적으로 계속 제공
//
// Performance / 성능:
//   - Negligible overhead (defer is cheap in Go)
//   - 무시할 만한 오버헤드 (Go에서 defer는 저렴함)
//   - Only activates on actual panic (zero cost in happy path)
//   - 실제 패닉 발생 시에만 활성화 (정상 경로에서 비용 없음)
//   - Stack trace generation only happens during panic
//   - 스택 트레이스 생성은 패닉 중에만 발생
//
// Common panic causes / 일반적인 패닉 원인:
//   - Nil pointer dereference: var p *User; p.Name
//   - Out of bounds access: arr[999] when len(arr) < 999
//   - Type assertion failure: val.(int) when val is not int
//   - Division by zero: x / 0 (for integers)
//   - Map concurrent read/write without mutex
//
// Best practices / 모범 사례:
//   - ALWAYS use Recovery() as the FIRST middleware
//   - 항상 Recovery()를 첫 번째 미들웨어로 사용
//   - Ensures all subsequent middleware/handlers are protected
//   - 모든 후속 미들웨어/핸들러가 보호되도록 보장
//   - Use RecoveryWithConfig() for custom logging behavior
//   - 커스텀 로깅 동작을 위해 RecoveryWithConfig() 사용
//   - Monitor panic frequency (frequent panics indicate bugs)
//   - 패닉 빈도 모니터링 (빈번한 패닉은 버그를 나타냄)
//
// Example usage / 사용 예제:
//
//	app := websvrutil.New()
//	app.Use(websvrutil.Recovery()) // MUST be first middleware
//	app.Use(websvrutil.Logger())
//	app.Use(websvrutil.CORS())
//
//	app.GET("/panic", func(w http.ResponseWriter, r *http.Request) {
//	    var user *User
//	    fmt.Fprintf(w, "Name: %s", user.Name) // PANIC: nil pointer dereference
//	    // Recovery catches this, logs stack trace, returns 500
//	})
//
// Advanced configuration / 고급 설정:
//
//	app.Use(websvrutil.RecoveryWithConfig(websvrutil.RecoveryConfig{
//	    PrintStack: true,
//	    LogFunc: func(err interface{}, stack []byte) {
//	        // Send to external logging service
//	        logger.Error("PANIC", "error", err, "stack", string(stack))
//	        // Send alert to monitoring system
//	        monitoring.SendAlert("Server panic", err)
//	    },
//	}))
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
// Example / 예제:
//
//	// Limit request body to 5MB
//	// 요청 본문을 5MB로 제한
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
		config.MaxBytes = 10 << 20 // 10MB default
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
// Example / 예제:
//
//	// Serve static files from "./public" directory
//	// "./public" 디렉토리에서 정적 파일 제공
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
// Example / 예제:
//
//	// Redirect all HTTP to HTTPS
//	// 모든 HTTP를 HTTPS로 리디렉션
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
// Example / 예제:
//
//	// Redirect all HTTP to HTTPS
//	// 모든 HTTP를 HTTPS로 리디렉션
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
// Example / 예제:
//
//	// Redirect to www version
//	// www 버전으로 리디렉션
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
