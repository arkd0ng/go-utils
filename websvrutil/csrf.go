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

	// CookieName is the name of the CSRF cookie (default: "_csrf") / CookieName은 CSRF 쿠키의 이름입니다 (기본값: "_csrf")
	CookieName string

	// CookiePath is the path of the CSRF cookie (default: "/") / CookiePath는 CSRF 쿠키의 경로입니다 (기본값: "/")
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

	// ContextKey is the key used to store the CSRF token in context (default: "csrf_token") / ContextKey는 컨텍스트에 CSRF 토큰을 저장하는 데 사용되는 키입니다 (기본값: "csrf_token")
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
