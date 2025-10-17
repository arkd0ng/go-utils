package websvrutil

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// csrf.go provides Cross-Site Request Forgery (CSRF) protection middleware.
//
// This file implements a comprehensive CSRF protection system that prevents
// unauthorized commands from being transmitted from a user that the web application trusts:
//
// Core Components:
//
// CSRFConfig:
//   - Configuration for CSRF protection behavior
//   - Token generation, validation, cookie settings
//   - Customizable token lookup (header/form/query)
//   - Cookie attributes (Secure, HttpOnly, SameSite)
//   - Error handling and validation skipping
//
// csrfTokenStore:
//   - Thread-safe in-memory token storage with expiration
//   - RWMutex-protected concurrent access
//   - Automatic cleanup of expired tokens (hourly)
//   - Global singleton instance (globalCSRFStore)
//
// Middleware Functions:
//   - CSRF(): Default CSRF protection with standard config
//   - CSRFWithConfig(config): Customizable CSRF protection
//   - GetCSRFToken(ctx): Retrieve token for use in templates
//
// Token Management:
//   - generateCSRFToken(length): Generate cryptographically secure random tokens
//     Uses crypto/rand for unpredictability
//     Default length: 32 bytes (256 bits)
//     Encoded as base64 for safe transmission
//   - getOrCreateCSRFToken(): Get existing or create new token
//     Stores token in cookie for client-side persistence
//     Adds token to server-side store with expiration
//   - validateCSRFToken(): Verify token matches expected value
//     Uses subtle.ConstantTimeCompare to prevent timing attacks
//
// Token Lookup Strategies:
//   - Header: "header:X-CSRF-Token" (default, recommended for AJAX)
//   - Form: "form:csrf_token" (for standard form submissions)
//   - Query: "query:csrf_token" (less common, useful for GET-based APIs)
//
// Protection Flow:
//   1. Middleware intercepts incoming request
//   2. If safe method (GET, HEAD, OPTIONS, TRACE), generate/refresh token
//   3. If unsafe method (POST, PUT, PATCH, DELETE):
//      - Extract token from configured location (header/form/query)
//      - Validate token matches server-side stored value
//      - Return 403 Forbidden if validation fails
//   4. Token stored in cookie and server-side store
//   5. Application retrieves token via GetCSRFToken() for rendering
//
// Cookie Configuration:
//   - CookieName: Cookie name (default: "_csrf")
//   - CookiePath: Cookie path scope (default: "/")
//   - CookieDomain: Cookie domain restriction
//   - CookieSecure: HTTPS-only transmission (recommended for production)
//   - CookieHTTPOnly: Prevent JavaScript access (default: true)
//   - CookieSameSite: SameSite attribute (default: SameSiteStrictMode)
//     Options: SameSiteStrictMode, SameSiteLaxMode, SameSiteNoneMode
//   - CookieMaxAge: Token lifetime in seconds (default: 86400 = 24 hours)
//
// Security Features:
//   - Cryptographically secure token generation (crypto/rand)
//   - Constant-time comparison (subtle.ConstantTimeCompare)
//     Prevents timing attacks that could leak token information
//   - Token expiration with automatic cleanup
//   - Per-request token validation
//   - Safe method exemption (GET/HEAD/OPTIONS/TRACE)
//   - Configurable validation skipping (Skipper function)
//
// Token Store:
//   - Thread-safe with RWMutex
//   - In-memory storage (tokens lost on server restart)
//   - Automatic hourly cleanup of expired tokens
//   - Token expiration tracked server-side
//
// Error Handling:
//   - Default: Returns 403 Forbidden with error message
//   - Custom: Provide ErrorHandler function in config
//   - Validation errors include detailed reason
//
// Example usage:
//
//	// Basic CSRF protection
//	app := New()
//	app.Use(CSRF())
//
//	// Custom CSRF configuration
//	app.Use(CSRFWithConfig(CSRFConfig{
//	    TokenLength:    32,
//	    TokenLookup:    "header:X-CSRF-Token",
//	    CookieName:     "_csrf",
//	    CookieSecure:   true,  // HTTPS only
//	    CookieSameSite: http.SameSiteStrictMode,
//	    CookieMaxAge:   86400, // 24 hours
//	    Skipper: func(r *http.Request) bool {
//	        // Skip CSRF for API endpoints with Bearer auth
//	        return strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ")
//	    },
//	}))
//
//	// Using token in templates
//	app.GET("/form", func(w http.ResponseWriter, r *http.Request) {
//	    ctx := GetContext(r)
//	    csrfToken := GetCSRFToken(ctx)
//	    // Render form with: <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
//	    ctx.Render(200, "form.html", map[string]string{"CSRFToken": csrfToken})
//	})
//
//	// Form submission protected by CSRF
//	app.POST("/submit", func(w http.ResponseWriter, r *http.Request) {
//	    // CSRF validation happens automatically
//	    // Handler only executes if validation passes
//	    // Process form submission...
//	})
//
// Performance:
//   - Token generation: ~10µs (crypto/rand)
//   - Validation: O(1) map lookup + constant-time comparison
//   - Memory: ~100 bytes per active token
//   - Cleanup: Runs hourly, removes expired tokens
//
// Limitations:
//   - In-memory storage (not suitable for multi-server deployments without sticky sessions)
//   - Tokens lost on server restart
//   - For distributed systems, consider external token store (Redis, database)
//
// Best Practices:
//   - Always use HTTPS in production (set CookieSecure: true)
//   - Use SameSiteStrictMode for maximum protection
//   - Set appropriate token expiration (balance security vs. user experience)
//   - Include CSRF token in all state-changing forms
//   - Use header-based tokens for AJAX requests
//
// csrf.go는 Cross-Site Request Forgery (CSRF) 보호 미들웨어를 제공합니다.
//
// 이 파일은 웹 애플리케이션이 신뢰하는 사용자로부터 전송되는
// 무단 명령을 방지하는 포괄적인 CSRF 보호 시스템을 구현합니다:
//
// 핵심 컴포넌트:
//
// CSRFConfig:
//   - CSRF 보호 동작 설정
//   - 토큰 생성, 검증, 쿠키 설정
//   - 커스터마이징 가능한 토큰 조회 (헤더/폼/쿼리)
//   - 쿠키 속성 (Secure, HttpOnly, SameSite)
//   - 에러 처리 및 검증 건너뛰기
//
// csrfTokenStore:
//   - 만료가 있는 스레드 안전 인메모리 토큰 저장소
//   - RWMutex로 보호된 동시 접근
//   - 만료된 토큰의 자동 정리 (시간당)
//   - 전역 싱글톤 인스턴스 (globalCSRFStore)
//
// 미들웨어 함수:
//   - CSRF(): 표준 설정으로 기본 CSRF 보호
//   - CSRFWithConfig(config): 커스터마이징 가능한 CSRF 보호
//   - GetCSRFToken(ctx): 템플릿에서 사용할 토큰 검색
//
// 토큰 관리:
//   - generateCSRFToken(length): 암호학적으로 안전한 랜덤 토큰 생성
//     예측 불가능성을 위해 crypto/rand 사용
//     기본 길이: 32바이트 (256비트)
//     안전한 전송을 위해 base64로 인코딩
//   - getOrCreateCSRFToken(): 기존 토큰 가져오기 또는 새 토큰 생성
//     클라이언트 측 지속성을 위해 쿠키에 토큰 저장
//     만료와 함께 서버 측 저장소에 토큰 추가
//   - validateCSRFToken(): 토큰이 예상 값과 일치하는지 확인
//     타이밍 공격 방지를 위해 subtle.ConstantTimeCompare 사용
//
// 토큰 조회 전략:
//   - 헤더: "header:X-CSRF-Token" (기본, AJAX에 권장)
//   - 폼: "form:csrf_token" (표준 폼 제출용)
//   - 쿼리: "query:csrf_token" (덜 일반적, GET 기반 API에 유용)
//
// 보호 흐름:
//   1. 미들웨어가 들어오는 요청 가로챔
//   2. 안전한 메서드 (GET, HEAD, OPTIONS, TRACE)인 경우, 토큰 생성/갱신
//   3. 안전하지 않은 메서드 (POST, PUT, PATCH, DELETE)인 경우:
//      - 설정된 위치 (헤더/폼/쿼리)에서 토큰 추출
//      - 토큰이 서버 측 저장 값과 일치하는지 검증
//      - 검증 실패 시 403 Forbidden 반환
//   4. 토큰을 쿠키 및 서버 측 저장소에 저장
//   5. 애플리케이션이 렌더링을 위해 GetCSRFToken()을 통해 토큰 검색
//
// 쿠키 설정:
//   - CookieName: 쿠키 이름 (기본: "_csrf")
//   - CookiePath: 쿠키 경로 범위 (기본: "/")
//   - CookieDomain: 쿠키 도메인 제한
//   - CookieSecure: HTTPS 전용 전송 (프로덕션에 권장)
//   - CookieHTTPOnly: JavaScript 접근 방지 (기본: true)
//   - CookieSameSite: SameSite 속성 (기본: SameSiteStrictMode)
//     옵션: SameSiteStrictMode, SameSiteLaxMode, SameSiteNoneMode
//   - CookieMaxAge: 토큰 수명 (초) (기본: 86400 = 24시간)
//
// 보안 기능:
//   - 암호학적으로 안전한 토큰 생성 (crypto/rand)
//   - 상수 시간 비교 (subtle.ConstantTimeCompare)
//     토큰 정보 유출 가능한 타이밍 공격 방지
//   - 자동 정리가 있는 토큰 만료
//   - 요청별 토큰 검증
//   - 안전한 메서드 면제 (GET/HEAD/OPTIONS/TRACE)
//   - 설정 가능한 검증 건너뛰기 (Skipper 함수)
//
// 토큰 저장소:
//   - RWMutex로 스레드 안전
//   - 인메모리 저장 (서버 재시작 시 토큰 손실)
//   - 만료된 토큰의 자동 시간당 정리
//   - 서버 측 토큰 만료 추적
//
// 에러 처리:
//   - 기본: 에러 메시지와 함께 403 Forbidden 반환
//   - 커스텀: 설정에서 ErrorHandler 함수 제공
//   - 검증 에러는 상세한 이유 포함
//
// 사용 예제:
//
//	// 기본 CSRF 보호
//	app := New()
//	app.Use(CSRF())
//
//	// 커스텀 CSRF 설정
//	app.Use(CSRFWithConfig(CSRFConfig{
//	    TokenLength:    32,
//	    TokenLookup:    "header:X-CSRF-Token",
//	    CookieName:     "_csrf",
//	    CookieSecure:   true,  // HTTPS 전용
//	    CookieSameSite: http.SameSiteStrictMode,
//	    CookieMaxAge:   86400, // 24시간
//	    Skipper: func(r *http.Request) bool {
//	        // Bearer 인증이 있는 API 엔드포인트는 CSRF 건너뜀
//	        return strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ")
//	    },
//	}))
//
//	// 템플릿에서 토큰 사용
//	app.GET("/form", func(w http.ResponseWriter, r *http.Request) {
//	    ctx := GetContext(r)
//	    csrfToken := GetCSRFToken(ctx)
//	    // 폼 렌더링: <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
//	    ctx.Render(200, "form.html", map[string]string{"CSRFToken": csrfToken})
//	})
//
//	// CSRF로 보호되는 폼 제출
//	app.POST("/submit", func(w http.ResponseWriter, r *http.Request) {
//	    // CSRF 검증이 자동으로 발생
//	    // 검증 통과 시에만 핸들러 실행
//	    // 폼 제출 처리...
//	})
//
// 성능:
//   - 토큰 생성: ~10µs (crypto/rand)
//   - 검증: O(1) 맵 조회 + 상수 시간 비교
//   - 메모리: 활성 토큰당 ~100바이트
//   - 정리: 시간당 실행, 만료된 토큰 제거
//
// 제한사항:
//   - 인메모리 저장 (스티키 세션 없이 다중 서버 배포에 적합하지 않음)
//   - 서버 재시작 시 토큰 손실
//   - 분산 시스템의 경우 외부 토큰 저장소 고려 (Redis, 데이터베이스)
//
// 모범 사례:
//   - 프로덕션에서 항상 HTTPS 사용 (CookieSecure: true 설정)
//   - 최대 보호를 위해 SameSiteStrictMode 사용
//   - 적절한 토큰 만료 설정 (보안과 사용자 경험 균형)
//   - 모든 상태 변경 폼에 CSRF 토큰 포함
//   - AJAX 요청에 헤더 기반 토큰 사용

// ============================================================================
// CSRF Protection
// CSRF 보호
// ============================================================================

// CSRFConfig represents the configuration for CSRF protection.
// CSRFConfig는 CSRF 보호를 위한 설정을 나타냅니다.
type CSRFConfig struct {
	// TokenLength is the length of the CSRF token in bytes (default: 32)
	// TokenLength는 CSRF 토큰의 바이트 길이입니다 (기본값: 32)
	TokenLength int

	// TokenLookup defines where to find the CSRF token
	// TokenLookup은 CSRF 토큰을 찾을 위치를 정의합니다
	// Format: "<source>:<name>"
	// Possible values:
	// - "header:<name>" (e.g., "header:X-CSRF-Token")
	// - "form:<name>" (e.g., "form:csrf_token")
	// - "query:<name>" (e.g., "query:csrf_token")
	// Default: "header:X-CSRF-Token"
	TokenLookup string

	// CookieName is the name of the CSRF cookie (default: "_csrf")
	// CookieName은 CSRF 쿠키의 이름입니다 (기본값: "_csrf")
	CookieName string

	// CookiePath is the path of the CSRF cookie (default: "/")
	// CookiePath는 CSRF 쿠키의 경로입니다 (기본값: "/")
	CookiePath string

	// CookieDomain is the domain of the CSRF cookie
	// CookieDomain은 CSRF 쿠키의 도메인입니다
	CookieDomain string

	// CookieSecure indicates if the cookie should only be sent over HTTPS (default: false)
	// CookieSecure는 쿠키가 HTTPS를 통해서만 전송되어야 하는지를 나타냅니다 (기본값: false)
	CookieSecure bool

	// CookieHTTPOnly indicates if the cookie should be HTTP only (default: true)
	// CookieHTTPOnly는 쿠키가 HTTP 전용이어야 하는지를 나타냅니다 (기본값: true)
	CookieHTTPOnly bool

	// CookieSameSite defines the SameSite cookie attribute (default: SameSiteStrictMode)
	// CookieSameSite는 SameSite 쿠키 속성을 정의합니다 (기본값: SameSiteStrictMode)
	CookieSameSite http.SameSite

	// CookieMaxAge is the max age of the CSRF cookie in seconds (default: 86400 = 24 hours)
	// CookieMaxAge는 CSRF 쿠키의 최대 수명(초)입니다 (기본값: 86400 = 24시간)
	CookieMaxAge int

	// ContextKey is the key used to store the CSRF token in context (default: "csrf_token")
	// ContextKey는 컨텍스트에 CSRF 토큰을 저장하는 데 사용되는 키입니다 (기본값: "csrf_token")
	ContextKey string

	// ErrorHandler is called when CSRF validation fails
	// ErrorHandler는 CSRF 검증이 실패할 때 호출됩니다
	// If not set, a default error handler will be used
	// 설정되지 않으면 기본 에러 핸들러가 사용됩니다
	ErrorHandler func(http.ResponseWriter, *http.Request, error)

	// Skipper defines a function to skip CSRF validation
	// Skipper는 CSRF 검증을 건너뛸 함수를 정의합니다
	// Return true to skip validation for the given request
	// 주어진 요청에 대한 검증을 건너뛰려면 true 반환
	Skipper func(*http.Request) bool
}

// csrfTokenStore holds CSRF tokens with expiration.
// csrfTokenStore는 만료 시간과 함께 CSRF 토큰을 보유합니다.
type csrfTokenStore struct {
	mu     sync.RWMutex
	tokens map[string]time.Time
}

// global token store
// 전역 토큰 저장소
var globalCSRFStore = &csrfTokenStore{
	tokens: make(map[string]time.Time),
}

// add adds a token to the store with expiration.
// add는 만료 시간과 함께 토큰을 저장소에 추가합니다.
func (s *csrfTokenStore) add(token string, maxAge int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[token] = time.Now().Add(time.Duration(maxAge) * time.Second)
}

// validate validates a token and removes it if expired.
// validate는 토큰을 검증하고 만료된 경우 제거합니다.
func (s *csrfTokenStore) validate(token string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	expiration, exists := s.tokens[token]
	if !exists {
		return false
	}

	// Check if token is expired
	// 토큰이 만료되었는지 확인
	if time.Now().After(expiration) {
		delete(s.tokens, token)
		return false
	}

	return true
}

// cleanup removes expired tokens from the store.
// cleanup은 저장소에서 만료된 토큰을 제거합니다.
func (s *csrfTokenStore) cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for token, expiration := range s.tokens {
		if now.After(expiration) {
			delete(s.tokens, token)
		}
	}
}

// Start periodic cleanup
// 주기적 정리 시작
func init() {
	// Run cleanup every hour
	// 1시간마다 정리 실행
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()
		for range ticker.C {
			globalCSRFStore.cleanup()
		}
	}()
}

// DefaultCSRFConfig returns default CSRF configuration.
// DefaultCSRFConfig는 기본 CSRF 설정을 반환합니다.
func DefaultCSRFConfig() CSRFConfig {
	return CSRFConfig{
		TokenLength:    32,
		TokenLookup:    "header:X-CSRF-Token",
		CookieName:     "_csrf",
		CookiePath:     "/",
		CookieSecure:   false,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		CookieMaxAge:   86400, // 24 hours
		ContextKey:     "csrf_token",
	}
}

// CSRF returns a CSRF middleware with default configuration.
// CSRF는 기본 설정으로 CSRF 미들웨어를 반환합니다.
//
// Example
// 예제:
//
//	app.Use(websvrutil.CSRF())
func CSRF() MiddlewareFunc {
	return CSRFWithConfig(DefaultCSRFConfig())
}

// CSRFWithConfig returns a CSRF middleware with custom configuration.
// CSRFWithConfig는 커스텀 설정으로 CSRF 미들웨어를 반환합니다.
//
// Example
// 예제:
//
//	app.Use(websvrutil.CSRFWithConfig(websvrutil.CSRFConfig{
//	    TokenLength:  32,
//	    CookieName:   "_csrf",
//	    CookieSecure: true,
//	}))
func CSRFWithConfig(config CSRFConfig) MiddlewareFunc {
	// Set defaults if not provided
	// 제공되지 않은 경우 기본값 설정
	if config.TokenLength == 0 {
		config.TokenLength = DefaultCSRFConfig().TokenLength
	}
	if config.TokenLookup == "" {
		config.TokenLookup = DefaultCSRFConfig().TokenLookup
	}
	if config.CookieName == "" {
		config.CookieName = DefaultCSRFConfig().CookieName
	}
	if config.CookiePath == "" {
		config.CookiePath = DefaultCSRFConfig().CookiePath
	}
	if config.CookieMaxAge == 0 {
		config.CookieMaxAge = DefaultCSRFConfig().CookieMaxAge
	}
	if config.ContextKey == "" {
		config.ContextKey = DefaultCSRFConfig().ContextKey
	}

	// Default error handler
	// 기본 에러 핸들러
	if config.ErrorHandler == nil {
		config.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, "CSRF token validation failed: "+err.Error(), http.StatusForbidden)
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip CSRF validation if Skipper returns true
			// Skipper가 true를 반환하면 CSRF 검증 건너뛰기
			if config.Skipper != nil && config.Skipper(r) {
				next.ServeHTTP(w, r)
				return
			}

			// Get or create CSRF token
			// CSRF 토큰 가져오기 또는 생성
			token, err := getOrCreateCSRFToken(w, r, &config)
			if err != nil {
				config.ErrorHandler(w, r, err)
				return
			}

			// Store token in context
			// 컨텍스트에 토큰 저장
			ctx := GetContext(r)
			if ctx != nil {
				ctx.Set(config.ContextKey, token)
			}

			// Validate CSRF token for unsafe methods
			// 안전하지 않은 메서드에 대한 CSRF 토큰 검증
			if !isSafeMethod(r.Method) {
				if err := validateCSRFToken(r, token, &config); err != nil {
					config.ErrorHandler(w, r, err)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

// getOrCreateCSRFToken retrieves or creates a CSRF token.
// getOrCreateCSRFToken은 CSRF 토큰을 검색하거나 생성합니다.
func getOrCreateCSRFToken(w http.ResponseWriter, r *http.Request, config *CSRFConfig) (string, error) {
	// Try to get token from cookie
	// 쿠키에서 토큰 가져오기 시도
	cookie, err := r.Cookie(config.CookieName)
	if err == nil && cookie.Value != "" {
		// Validate that token exists in store
		// 저장소에 토큰이 존재하는지 검증
		if globalCSRFStore.validate(cookie.Value) {
			return cookie.Value, nil
		}
	}

	// Generate new token
	// 새 토큰 생성
	token, err := generateCSRFToken(config.TokenLength)
	if err != nil {
		return "", fmt.Errorf("failed to generate CSRF token: %w", err)
	}

	// Store token
	// 토큰 저장
	globalCSRFStore.add(token, config.CookieMaxAge)

	// Set cookie
	// 쿠키 설정
	http.SetCookie(w, &http.Cookie{
		Name:     config.CookieName,
		Value:    token,
		Path:     config.CookiePath,
		Domain:   config.CookieDomain,
		MaxAge:   config.CookieMaxAge,
		Secure:   config.CookieSecure,
		HttpOnly: config.CookieHTTPOnly,
		SameSite: config.CookieSameSite,
	})

	return token, nil
}

// validateCSRFToken validates the CSRF token from the request.
// validateCSRFToken은 요청에서 CSRF 토큰을 검증합니다.
func validateCSRFToken(r *http.Request, expectedToken string, config *CSRFConfig) error {
	// Extract token from request
	// 요청에서 토큰 추출
	token := extractCSRFToken(r, config)
	if token == "" {
		return fmt.Errorf("CSRF token not found")
	}

	// Constant-time comparison to prevent timing attacks
	// 타이밍 공격을 방지하기 위한 상수 시간 비교
	if subtle.ConstantTimeCompare([]byte(token), []byte(expectedToken)) != 1 {
		return fmt.Errorf("CSRF token mismatch")
	}

	return nil
}

// extractCSRFToken extracts the CSRF token from the request.
// extractCSRFToken은 요청에서 CSRF 토큰을 추출합니다.
func extractCSRFToken(r *http.Request, config *CSRFConfig) string {
	// Parse TokenLookup
	// TokenLookup 파싱
	// Format: "<source>:<name>"
	source := "header"
	name := "X-CSRF-Token"

	if config.TokenLookup != "" {
		parts := splitTokenLookup(config.TokenLookup)
		if len(parts) == 2 {
			source = parts[0]
			name = parts[1]
		}
	}

	// Extract token based on source
	// 소스에 따라 토큰 추출
	switch source {
	case "header":
		return r.Header.Get(name)
	case "form":
		return r.FormValue(name)
	case "query":
		return r.URL.Query().Get(name)
	default:
		return r.Header.Get(name)
	}
}

// splitTokenLookup splits the TokenLookup string into source and name.
// splitTokenLookup은 TokenLookup 문자열을 소스와 이름으로 분할합니다.
func splitTokenLookup(lookup string) []string {
	for i := 0; i < len(lookup); i++ {
		if lookup[i] == ':' {
			return []string{lookup[:i], lookup[i+1:]}
		}
	}
	return []string{lookup}
}

// generateCSRFToken generates a cryptographically secure random token.
// generateCSRFToken은 암호학적으로 안전한 랜덤 토큰을 생성합니다.
func generateCSRFToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// isSafeMethod checks if the HTTP method is considered safe (doesn't modify data).
// isSafeMethod는 HTTP 메서드가 안전한 것으로 간주되는지 확인합니다 (데이터를 수정하지 않음).
func isSafeMethod(method string) bool {
	return method == http.MethodGet ||
		method == http.MethodHead ||
		method == http.MethodOptions ||
		method == http.MethodTrace
}

// GetCSRFToken retrieves the CSRF token from the context.
// GetCSRFToken은 컨텍스트에서 CSRF 토큰을 검색합니다.
//
// This is useful for rendering the token in HTML forms or JavaScript.
// HTML 폼이나 JavaScript에서 토큰을 렌더링하는 데 유용합니다.
//
// Example
// 예제:
//
//	token := websvrutil.GetCSRFToken(ctx)
//	// Use token in HTML form:
//	// <input type="hidden" name="csrf_token" value="{{.Token}}">
func GetCSRFToken(ctx *Context) string {
	return ctx.GetString("csrf_token")
}
