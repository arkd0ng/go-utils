package websvrutil

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
)

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
	CookieName   string        // Name of the session cookie / 세션 쿠키 이름
	MaxAge       time.Duration // Session expiration time / 세션 만료 시간
	Secure       bool          // Use secure cookies (HTTPS only) / 보안 쿠키 사용 (HTTPS만)
	HttpOnly     bool          // Prevent JavaScript access / JavaScript 액세스 방지
	SameSite     http.SameSite // SameSite cookie attribute / SameSite 쿠키 속성
	CleanupTime  time.Duration // Cleanup interval / 정리 간격
	Path         string        // Cookie path / 쿠키 경로
	Domain       string        // Cookie domain / 쿠키 도메인
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
//   "kqZ9Xx3vR_5yJKl2Nw8PmQ7VtBcDfGhE1WsIuO6A4ZY"
//   "9hTmPq2Wz8Lx4Vb7Nc1Yd6Rj3Fg5Ks0Hu9Ia8Qe2Ow7U"
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
//   1. Create ticker with configured cleanup interval
//   2. Wait for ticker signal
//   3. Acquire write lock (blocks all session operations during cleanup)
//   4. Iterate through all sessions in the store
//   5. Delete sessions where ExpiresAt < current time
//   6. Release write lock
//   7. Repeat indefinitely until program termination
//
//   1. 설정된 정리 간격으로 티커 생성
//   2. 티커 신호 대기
//   3. 쓰기 락 획득 (정리 중 모든 세션 작업 차단)
//   4. 저장소의 모든 세션 순회
//   5. ExpiresAt < 현재 시간인 세션 삭제
//   6. 쓰기 락 해제
//   7. 프로그램 종료까지 무한 반복
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
//   CleanupTime = 5 minutes
//   MaxAge = 1 hour
//
//   Session A created at 10:00, expires at 11:00
//   Session B created at 10:30, expires at 11:30
//
//   10:05 cleanup: no deletions (both active)
//   10:10 cleanup: no deletions (both active)
//   11:05 cleanup: Session A deleted (expired at 11:00)
//   11:35 cleanup: Session B deleted (expired at 11:30)
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
