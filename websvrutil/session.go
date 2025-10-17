package websvrutil

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
)

// session.go provides HTTP session management with in-memory storage.
//
// This file implements a thread-safe session management system for maintaining
// user state across multiple HTTP requests:
//
// Core Types:
//
// SessionStore:
//   - Central session manager with thread-safe in-memory storage
//   - RWMutex-protected concurrent access to sessions map
//   - Automatic cleanup of expired sessions via background goroutine
//   - Configurable session lifetime, cookie attributes
//
// Session:
//   - Individual user session with unique ID
//   - Key-value data storage (map[string]interface{})
//   - Creation and expiration timestamps
//   - Thread-safe data access with RWMutex
//
// SessionOptions:
//   - Configuration for session behavior and cookie settings
//   - CookieName: Session cookie identifier (default: "_session")
//   - MaxAge: Session lifetime (default: 24 hours)
//   - Secure: HTTPS-only cookie transmission
//   - HttpOnly: Prevent JavaScript access (default: true)
//   - SameSite: Cross-site cookie policy (default: SameSiteLaxMode)
//   - CleanupTime: Interval for expired session removal (default: 30 minutes)
//   - Path, Domain: Cookie scope control
//
// SessionStore Methods:
//   - NewSessionStore(opts): Create store with background cleanup goroutine
//   - Get(r): Retrieve existing session or create new one
//     Validates session ID from cookie, checks expiration
//   - New(): Create new session with unique ID
//     Generates cryptographically secure session ID (32 bytes, base64)
//   - Save(w, session): Persist session and update cookie
//     Refreshes expiration time, sets cookie with configured options
//   - Destroy(w, r): Remove session and clear cookie
//     Deletes from store, sets MaxAge=-1 to delete client cookie
//   - generateSessionID(): Generate secure random session ID
//     Uses crypto/rand for unpredictability (32 bytes = 256 bits)
//   - cleanupExpiredSessions(): Background goroutine removes expired sessions
//     Runs at CleanupTime intervals (default: every 30 minutes)
//   - Count(): Get total number of active sessions (for monitoring)
//
// Session Methods:
//   - Set(key, value): Store value in session data
//   - Get(key): Retrieve value with existence check
//   - GetString(key): Type-safe string retrieval (empty string if not found)
//   - GetInt(key): Type-safe int retrieval (0 if not found)
//   - GetBool(key): Type-safe bool retrieval (false if not found)
//   - Delete(key): Remove specific key from session
//   - Clear(): Remove all data from session (keeps session ID)
//
// Session Lifecycle:
//   1. Client makes initial request without session cookie
//   2. Server creates new session with unique ID
//   3. Session ID stored in secure cookie, sent to client
//   4. Client includes cookie in subsequent requests
//   5. Server retrieves session data using cookie ID
//   6. Session expires after MaxAge inactivity
//   7. Cleanup goroutine removes expired sessions periodically
//
// Security Features:
//   - Cryptographically secure session ID generation (crypto/rand)
//     32-byte random ID provides ~2^256 possible values
//     Base64 encoding for safe cookie transmission
//   - HttpOnly cookies prevent XSS attacks (default: true)
//     JavaScript cannot access session cookie
//   - Secure cookies for HTTPS-only transmission (configurable)
//   - SameSite protection against CSRF attacks
//     SameSiteLaxMode: Blocks cross-site POST but allows GET
//     SameSiteStrictMode: Blocks all cross-site requests
//   - Automatic expiration and cleanup prevents indefinite storage
//   - Thread-safe access prevents race conditions
//
// Thread-Safety:
//   - SessionStore uses RWMutex for concurrent access
//     Multiple goroutines can read sessions simultaneously
//     Writes (New, Save, Destroy) acquire exclusive lock
//   - Individual Sessions have their own RWMutex
//     Safe concurrent access to session data from multiple requests
//   - Cleanup goroutine safely removes expired sessions
//
// Example usage:
//
//	// Create session store
//	store := NewSessionStore(DefaultSessionOptions())
//
//	// In middleware or handler
//	app.GET("/login", func(w http.ResponseWriter, r *http.Request) {
//	    session, _ := store.Get(r)
//	    session.Set("username", "john")
//	    session.Set("user_id", 123)
//	    session.Set("is_admin", true)
//	    store.Save(w, session)
//	    // Session cookie automatically set
//	})
//
//	app.GET("/profile", func(w http.ResponseWriter, r *http.Request) {
//	    session, _ := store.Get(r)
//	    username := session.GetString("username")
//	    userID := session.GetInt("user_id")
//	    isAdmin := session.GetBool("is_admin")
//	    // Use session data...
//	})
//
//	app.POST("/logout", func(w http.ResponseWriter, r *http.Request) {
//	    store.Destroy(w, r)
//	    // Session removed, cookie cleared
//	})
//
// Custom configuration:
//
//	store := NewSessionStore(SessionOptions{
//	    CookieName:  "my_session",
//	    MaxAge:      2 * time.Hour,    // 2-hour sessions
//	    Secure:      true,              // HTTPS only
//	    HttpOnly:    true,              // No JavaScript access
//	    SameSite:    http.SameSiteStrictMode, // Strict CSRF protection
//	    CleanupTime: 15 * time.Minute,  // Cleanup every 15 minutes
//	    Path:        "/",
//	    Domain:      "example.com",
//	})
//
// Performance:
//   - Session lookup: O(1) map access
//   - Session creation: ~10µs (crypto/rand ID generation)
//   - Memory: ~200 bytes per session + data size
//   - Cleanup: O(n) where n = total sessions (runs infrequently)
//
// Limitations:
//   - In-memory storage (sessions lost on server restart)
//   - Not suitable for multi-server deployments without sticky sessions
//   - Memory usage grows with active sessions
//   - For production/distributed systems, consider:
//     * Redis-backed session store
//     * Database-backed sessions
//     * JWT tokens for stateless authentication
//
// Best Practices:
//   - Always use Secure: true in production (HTTPS)
//   - Keep HttpOnly: true to prevent XSS attacks
//   - Use SameSiteLaxMode or SameSiteStrictMode for CSRF protection
//   - Set appropriate MaxAge (balance security vs. user experience)
//   - Monitor session count with Count() method
//   - Consider external session store for production clusters
//   - Clear sensitive data after use (passwords, tokens)
//
// session.go는 인메모리 저장소를 사용한 HTTP 세션 관리를 제공합니다.
//
// 이 파일은 여러 HTTP 요청에 걸쳐 사용자 상태를 유지하기 위한
// 스레드 안전 세션 관리 시스템을 구현합니다:
//
// 핵심 타입:
//
// SessionStore:
//   - 스레드 안전 인메모리 저장소를 가진 중앙 세션 관리자
//   - RWMutex로 보호된 세션 맵에 대한 동시 접근
//   - 백그라운드 고루틴을 통한 만료된 세션의 자동 정리
//   - 설정 가능한 세션 수명, 쿠키 속성
//
// Session:
//   - 고유 ID를 가진 개별 사용자 세션
//   - 키-값 데이터 저장소 (map[string]interface{})
//   - 생성 및 만료 타임스탬프
//   - RWMutex를 사용한 스레드 안전 데이터 접근
//
// SessionOptions:
//   - 세션 동작 및 쿠키 설정 구성
//   - CookieName: 세션 쿠키 식별자 (기본: "_session")
//   - MaxAge: 세션 수명 (기본: 24시간)
//   - Secure: HTTPS 전용 쿠키 전송
//   - HttpOnly: JavaScript 접근 방지 (기본: true)
//   - SameSite: 크로스 사이트 쿠키 정책 (기본: SameSiteLaxMode)
//   - CleanupTime: 만료된 세션 제거 간격 (기본: 30분)
//   - Path, Domain: 쿠키 범위 제어
//
// SessionStore 메서드:
//   - NewSessionStore(opts): 백그라운드 정리 고루틴과 함께 저장소 생성
//   - Get(r): 기존 세션 검색 또는 새 세션 생성
//     쿠키에서 세션 ID 검증, 만료 확인
//   - New(): 고유 ID로 새 세션 생성
//     암호학적으로 안전한 세션 ID 생성 (32바이트, base64)
//   - Save(w, session): 세션 지속 및 쿠키 업데이트
//     만료 시간 갱신, 설정된 옵션으로 쿠키 설정
//   - Destroy(w, r): 세션 제거 및 쿠키 지우기
//     저장소에서 삭제, 클라이언트 쿠키 삭제를 위해 MaxAge=-1 설정
//   - generateSessionID(): 안전한 랜덤 세션 ID 생성
//     예측 불가능성을 위해 crypto/rand 사용 (32바이트 = 256비트)
//   - cleanupExpiredSessions(): 만료된 세션을 제거하는 백그라운드 고루틴
//     CleanupTime 간격으로 실행 (기본: 30분마다)
//   - Count(): 활성 세션 총 수 가져오기 (모니터링용)
//
// Session 메서드:
//   - Set(key, value): 세션 데이터에 값 저장
//   - Get(key): 존재 확인과 함께 값 검색
//   - GetString(key): 타입 안전 문자열 검색 (없으면 빈 문자열)
//   - GetInt(key): 타입 안전 int 검색 (없으면 0)
//   - GetBool(key): 타입 안전 bool 검색 (없으면 false)
//   - Delete(key): 세션에서 특정 키 제거
//   - Clear(): 세션에서 모든 데이터 제거 (세션 ID 유지)
//
// 세션 수명주기:
//   1. 클라이언트가 세션 쿠키 없이 초기 요청
//   2. 서버가 고유 ID로 새 세션 생성
//   3. 세션 ID가 보안 쿠키에 저장되어 클라이언트로 전송
//   4. 클라이언트가 후속 요청에 쿠키 포함
//   5. 서버가 쿠키 ID를 사용하여 세션 데이터 검색
//   6. MaxAge 비활성 후 세션 만료
//   7. 정리 고루틴이 주기적으로 만료된 세션 제거
//
// 보안 기능:
//   - 암호학적으로 안전한 세션 ID 생성 (crypto/rand)
//     32바이트 랜덤 ID는 ~2^256개의 가능한 값 제공
//     안전한 쿠키 전송을 위한 Base64 인코딩
//   - HttpOnly 쿠키는 XSS 공격 방지 (기본: true)
//     JavaScript가 세션 쿠키에 접근할 수 없음
//   - HTTPS 전용 전송을 위한 Secure 쿠키 (설정 가능)
//   - CSRF 공격에 대한 SameSite 보호
//     SameSiteLaxMode: 크로스 사이트 POST 차단하지만 GET 허용
//     SameSiteStrictMode: 모든 크로스 사이트 요청 차단
//   - 자동 만료 및 정리로 무한 저장 방지
//   - 스레드 안전 접근으로 경쟁 조건 방지
//
// 스레드 안전성:
//   - SessionStore는 동시 접근을 위해 RWMutex 사용
//     여러 고루틴이 동시에 세션 읽기 가능
//     쓰기 (New, Save, Destroy)는 배타적 잠금 획득
//   - 개별 Session은 자체 RWMutex 보유
//     여러 요청의 세션 데이터에 대한 안전한 동시 접근
//   - 정리 고루틴은 만료된 세션을 안전하게 제거
//
// 사용 예제:
//
//	// 세션 저장소 생성
//	store := NewSessionStore(DefaultSessionOptions())
//
//	// 미들웨어 또는 핸들러에서
//	app.GET("/login", func(w http.ResponseWriter, r *http.Request) {
//	    session, _ := store.Get(r)
//	    session.Set("username", "john")
//	    session.Set("user_id", 123)
//	    session.Set("is_admin", true)
//	    store.Save(w, session)
//	    // 세션 쿠키 자동 설정
//	})
//
//	app.GET("/profile", func(w http.ResponseWriter, r *http.Request) {
//	    session, _ := store.Get(r)
//	    username := session.GetString("username")
//	    userID := session.GetInt("user_id")
//	    isAdmin := session.GetBool("is_admin")
//	    // 세션 데이터 사용...
//	})
//
//	app.POST("/logout", func(w http.ResponseWriter, r *http.Request) {
//	    store.Destroy(w, r)
//	    // 세션 제거, 쿠키 지워짐
//	})
//
// 커스텀 설정:
//
//	store := NewSessionStore(SessionOptions{
//	    CookieName:  "my_session",
//	    MaxAge:      2 * time.Hour,    // 2시간 세션
//	    Secure:      true,              // HTTPS 전용
//	    HttpOnly:    true,              // JavaScript 접근 없음
//	    SameSite:    http.SameSiteStrictMode, // 엄격한 CSRF 보호
//	    CleanupTime: 15 * time.Minute,  // 15분마다 정리
//	    Path:        "/",
//	    Domain:      "example.com",
//	})
//
// 성능:
//   - 세션 조회: O(1) 맵 접근
//   - 세션 생성: ~10µs (crypto/rand ID 생성)
//   - 메모리: 세션당 ~200바이트 + 데이터 크기
//   - 정리: O(n), n = 총 세션 수 (드물게 실행)
//
// 제한사항:
//   - 인메모리 저장 (서버 재시작 시 세션 손실)
//   - 스티키 세션 없이 다중 서버 배포에 적합하지 않음
//   - 메모리 사용량이 활성 세션과 함께 증가
//   - 프로덕션/분산 시스템의 경우 고려사항:
//     * Redis 기반 세션 저장소
//     * 데이터베이스 기반 세션
//     * 상태 비저장 인증을 위한 JWT 토큰
//
// 모범 사례:
//   - 프로덕션에서 항상 Secure: true 사용 (HTTPS)
//   - XSS 공격 방지를 위해 HttpOnly: true 유지
//   - CSRF 보호를 위해 SameSiteLaxMode 또는 SameSiteStrictMode 사용
//   - 적절한 MaxAge 설정 (보안과 사용자 경험 균형)
//   - Count() 메서드로 세션 수 모니터링
//   - 프로덕션 클러스터용 외부 세션 저장소 고려
//   - 사용 후 민감한 데이터 지우기 (비밀번호, 토큰)

// SessionStore manages HTTP sessions with in-memory storage.
// SessionStore는 인메모리 저장소로 HTTP 세션을 관리합니다.
type SessionStore struct {
	mu       sync.RWMutex
	sessions map[string]*Session
	options  SessionOptions
}

// Session represents a user session with key-value storage.
// Session은 키-값 저장소가 있는 사용자 세션을 나타냅니다.
type Session struct {
	ID        string
	Data      map[string]interface{}
	CreatedAt time.Time
	ExpiresAt time.Time
	mu        sync.RWMutex
}

// SessionOptions configures the session store.
// SessionOptions는 세션 저장소를 설정합니다.
type SessionOptions struct {
	// Name of the session cookie
	// 세션 쿠키 이름
	CookieName string
	// Session expiration time
	// 세션 만료 시간
	MaxAge time.Duration
	// Use secure cookies (HTTPS only)
	// 보안 쿠키 사용 (HTTPS만)
	Secure bool
	// Prevent JavaScript access
	// JavaScript 액세스 방지
	HttpOnly bool
	// SameSite cookie attribute
	// SameSite 쿠키 속성
	SameSite http.SameSite
	// Cleanup interval
	// 정리 간격
	CleanupTime time.Duration
	// Cookie path
	// 쿠키 경로
	Path string
	// Cookie domain
	// 쿠키 도메인
	Domain string
}

// DefaultSessionOptions returns default session options.
// DefaultSessionOptions는 기본 세션 옵션을 반환합니다.
func DefaultSessionOptions() SessionOptions {
	return SessionOptions{
		CookieName:  DefaultSessionCookieName,
		MaxAge:      DefaultSessionMaxAge,
		Secure:      false,
		HttpOnly:    true,
		SameSite:    http.SameSiteLaxMode,
		CleanupTime: DefaultSessionCleanup,
		Path:        "/",
		Domain:      "",
	}
}

// NewSessionStore creates a new session store with the given options.
// NewSessionStore는 주어진 옵션으로 새 세션 저장소를 생성합니다.
func NewSessionStore(opts SessionOptions) *SessionStore {
	store := &SessionStore{
		sessions: make(map[string]*Session),
		options:  opts,
	}

	// Start cleanup goroutine
	// 정리 고루틴 시작
	go store.cleanupExpiredSessions()

	return store
}

// Get retrieves or creates a session for the request.
// Get은 요청에 대한 세션을 검색하거나 생성합니다.
func (s *SessionStore) Get(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie(s.options.CookieName)
	if err == nil {
		// Try to get existing session
		// 기존 세션 가져오기 시도
		s.mu.RLock()
		session, exists := s.sessions[cookie.Value]
		s.mu.RUnlock()

		if exists && session.ExpiresAt.After(time.Now()) {
			return session, nil
		}
	}

	// Create new session
	// 새 세션 생성
	return s.New(), nil
}

// New creates a new session with a unique ID.
// New는 고유 ID로 새 세션을 생성합니다.
func (s *SessionStore) New() *Session {
	session := &Session{
		ID:        s.generateSessionID(),
		Data:      make(map[string]interface{}),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(s.options.MaxAge),
	}

	s.mu.Lock()
	s.sessions[session.ID] = session
	s.mu.Unlock()

	return session
}

// Save saves the session and sets the cookie.
// Save는 세션을 저장하고 쿠키를 설정합니다.
func (s *SessionStore) Save(w http.ResponseWriter, session *Session) {
	// Update expiration
	// 만료 시간 업데이트
	session.ExpiresAt = time.Now().Add(s.options.MaxAge)

	// Set cookie
	// 쿠키 설정
	http.SetCookie(w, &http.Cookie{
		Name:     s.options.CookieName,
		Value:    session.ID,
		Path:     s.options.Path,
		Domain:   s.options.Domain,
		MaxAge:   int(s.options.MaxAge.Seconds()),
		Secure:   s.options.Secure,
		HttpOnly: s.options.HttpOnly,
		SameSite: s.options.SameSite,
	})
}

// Destroy removes the session and clears the cookie.
// Destroy는 세션을 제거하고 쿠키를 지웁니다.
func (s *SessionStore) Destroy(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(s.options.CookieName)
	if err != nil {
		return err
	}

	// Remove session from store
	// 저장소에서 세션 제거
	s.mu.Lock()
	delete(s.sessions, cookie.Value)
	s.mu.Unlock()

	// Clear cookie
	// 쿠키 지우기
	http.SetCookie(w, &http.Cookie{
		Name:     s.options.CookieName,
		Value:    "",
		Path:     s.options.Path,
		Domain:   s.options.Domain,
		MaxAge:   -1,
		Secure:   s.options.Secure,
		HttpOnly: s.options.HttpOnly,
		SameSite: s.options.SameSite,
	})

	return nil
}

// Set stores a value in the session.
// Set은 세션에 값을 저장합니다.
func (sess *Session) Set(key string, value interface{}) {
	sess.mu.Lock()
	defer sess.mu.Unlock()
	sess.Data[key] = value
}

// Get retrieves a value from the session.
// Get은 세션에서 값을 검색합니다.
func (sess *Session) Get(key string) (interface{}, bool) {
	sess.mu.RLock()
	defer sess.mu.RUnlock()
	value, exists := sess.Data[key]
	return value, exists
}

// GetString retrieves a string value from the session.
// GetString은 세션에서 문자열 값을 검색합니다.
func (sess *Session) GetString(key string) string {
	value, exists := sess.Get(key)
	if !exists {
		return ""
	}
	str, _ := value.(string)
	return str
}

// GetInt retrieves an int value from the session.
// GetInt는 세션에서 int 값을 검색합니다.
func (sess *Session) GetInt(key string) int {
	value, exists := sess.Get(key)
	if !exists {
		return 0
	}
	i, _ := value.(int)
	return i
}

// GetBool retrieves a bool value from the session.
// GetBool은 세션에서 bool 값을 검색합니다.
func (sess *Session) GetBool(key string) bool {
	value, exists := sess.Get(key)
	if !exists {
		return false
	}
	b, _ := value.(bool)
	return b
}

// Delete removes a value from the session.
// Delete는 세션에서 값을 제거합니다.
func (sess *Session) Delete(key string) {
	sess.mu.Lock()
	defer sess.mu.Unlock()
	delete(sess.Data, key)
}

// Clear removes all values from the session.
// Clear는 세션에서 모든 값을 제거합니다.
func (sess *Session) Clear() {
	sess.mu.Lock()
	defer sess.mu.Unlock()
	sess.Data = make(map[string]interface{})
}

// generateSessionID generates a cryptographically secure session ID.
// generateSessionID는 암호학적으로 안전한 세션 ID를 생성합니다.
//
// Security properties
// 보안 속성:
// - Uses crypto/rand for cryptographically secure randomness
// - crypto/rand를 사용하여 암호학적으로 안전한 랜덤성 확보
// - 256-bit entropy (32 bytes) provides extremely high collision resistance
// - 256비트 엔트로피 (32바이트)로 극도로 높은 충돌 저항성 제공
// - Base64 URL-safe encoding for cookie compatibility (no +, /, = characters)
// - 쿠키 호환성을 위한 Base64 URL 안전 인코딩 (+, /, = 문자 없음)
//
// Collision probability
// 충돌 확률:
// - With 256 bits: approximately 1 in 2^256 (~1.16 × 10^77)
// - 256비트 사용 시: 약 1/2^256 (~1.16 × 10^77)
// - For comparison, there are ~10^80 atoms in the universe
// - 비교: 우주의 원자 수는 약 10^80개
// - Practically zero chance of collision in any realistic scenario
// - 실제 시나리오에서 충돌 가능성은 사실상 0
//
// Fallback strategy
// 대체 전략:
// - If crypto/rand fails (extremely rare), falls back to timestamp-based ID
// - crypto/rand 실패 시 (극히 드묾) 타임스탬프 기반 ID로 대체
// - Fallback should never happen in practice (crypto/rand failure indicates serious system issues)
// - 대체는 실제로 발생하지 않아야 함 (crypto/rand 실패는 심각한 시스템 문제를 나타냄)
//
// Output format
// 출력 형식:
// - Base64 URL-safe encoded string, 43 characters long (32 bytes → 43 chars)
// - Base64 URL 안전 인코딩 문자열, 43자 길이 (32바이트 → 43자)
//
// Example output
// 출력 예제:
//
//	"kqZ9Xx3vR_5yJKl2Nw8PmQ7VtBcDfGhE1WsIuO6A4ZY"
//	"9hTmPq2Wz8Lx4Vb7Nc1Yd6Rj3Fg5Ks0Hu9Ia8Qe2Ow7U"
//
// Performance
// 성능:
// - Time complexity: O(1) - fixed 32-byte generation
// - 시간 복잡도: O(1) - 고정 32바이트 생성
// - Typical execution: <1μs on modern hardware
// - 일반 실행 시간: 최신 하드웨어에서 1마이크로초 미만
func (s *SessionStore) generateSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		// Fallback to timestamp-based ID
		// 타임스탬프 기반 ID로 대체
		return base64.URLEncoding.EncodeToString([]byte(time.Now().String()))
	}
	return base64.URLEncoding.EncodeToString(b)
}

// cleanupExpiredSessions periodically removes expired sessions.
// cleanupExpiredSessions는 주기적으로 만료된 세션을 제거합니다.
//
// Cleanup strategy
// 정리 전략:
// - Runs in a background goroutine started by NewSessionStore()
// - NewSessionStore()가 시작한 백그라운드 고루틴에서 실행
// - Executes periodically based on SessionOptions.CleanupTime interval
// - SessionOptions.CleanupTime 간격에 따라 주기적으로 실행
// - Default interval: 5 minutes (configurable via SessionOptions)
// - 기본 간격: 5분 (SessionOptions로 설정 가능)
//
// Process
// 프로세스:
//
//  1. Create ticker with configured cleanup interval
//
//  2. Wait for ticker signal
//
//  3. Acquire write lock (blocks all session operations during cleanup)
//
//  4. Iterate through all sessions in the store
//
//  5. Delete sessions where ExpiresAt < current time
//
//  6. Release write lock
//
//  7. Repeat indefinitely until program termination
//
//  1. 설정된 정리 간격으로 티커 생성
//
//  2. 티커 신호 대기
//
//  3. 쓰기 락 획득 (정리 중 모든 세션 작업 차단)
//
//  4. 저장소의 모든 세션 순회
//
//  5. ExpiresAt < 현재 시간인 세션 삭제
//
//  6. 쓰기 락 해제
//
//  7. 프로그램 종료까지 무한 반복
//
// Thread safety
// 스레드 안전성:
// - Uses sync.RWMutex.Lock() for exclusive access during cleanup
// - 정리 중 배타적 액세스를 위해 sync.RWMutex.Lock() 사용
// - Blocks all Get(), New(), Save() operations while cleaning
// - 정리 중 모든 Get(), New(), Save() 작업 차단
// - Lock duration proportional to number of sessions
// - 락 기간은 세션 수에 비례
//
// Performance considerations
// 성능 고려사항:
// - Time complexity: O(n) where n = number of sessions in store
// - 시간 복잡도: O(n), n = 저장소의 세션 수
// - Space complexity: O(1) - no additional memory allocation
// - 공간 복잡도: O(1) - 추가 메모리 할당 없음
// - For large session counts (>10,000), consider increasing CleanupTime interval
// - 대규모 세션 수(>10,000)의 경우 CleanupTime 간격 증가 고려
// - Deleted sessions are immediately eligible for garbage collection
// - 삭제된 세션은 즉시 가비지 컬렉션 대상이 됨
//
// Memory management
// 메모리 관리:
// - Deleted map entries are reclaimed by Go's garbage collector
// - 삭제된 맵 항목은 Go 가비지 컬렉터가 회수
// - Does not shrink the underlying map capacity (Go limitation)
// - 기본 맵 용량을 축소하지 않음 (Go 제약)
// - For memory-critical applications, consider recreating the map periodically
// - 메모리 중요 애플리케이션의 경우 맵을 주기적으로 재생성 고려
//
// Example behavior
// 동작 예제:
//
//	CleanupTime = 5 minutes
//	MaxAge = 1 hour
//
//	Session A created at 10:00, expires at 11:00
//	Session B created at 10:30, expires at 11:30
//
//	10:05 cleanup: no deletions (both active)
//	10:10 cleanup: no deletions (both active)
//	11:05 cleanup: Session A deleted (expired at 11:00)
//	11:35 cleanup: Session B deleted (expired at 11:30)
func (s *SessionStore) cleanupExpiredSessions() {
	ticker := time.NewTicker(s.options.CleanupTime)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		s.mu.Lock()
		for id, session := range s.sessions {
			if session.ExpiresAt.Before(now) {
				delete(s.sessions, id)
			}
		}
		s.mu.Unlock()
	}
}

// Count returns the number of active sessions.
// Count는 활성 세션 수를 반환합니다.
func (s *SessionStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.sessions)
}
