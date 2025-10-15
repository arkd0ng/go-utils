package websvrutil

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
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
