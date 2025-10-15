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
		CookieName:  "sessionid",
		MaxAge:      24 * time.Hour,
		Secure:      false,
		HttpOnly:    true,
		SameSite:    http.SameSiteLaxMode,
		CleanupTime: 5 * time.Minute,
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

	// Start cleanup goroutine / 정리 고루틴 시작
	go store.cleanupExpiredSessions()

	return store
}

// Get retrieves or creates a session for the request.
// Get은 요청에 대한 세션을 검색하거나 생성합니다.
func (s *SessionStore) Get(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie(s.options.CookieName)
	if err == nil {
		// Try to get existing session / 기존 세션 가져오기 시도
		s.mu.RLock()
		session, exists := s.sessions[cookie.Value]
		s.mu.RUnlock()

		if exists && session.ExpiresAt.After(time.Now()) {
			return session, nil
		}
	}

	// Create new session / 새 세션 생성
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
	// Update expiration / 만료 시간 업데이트
	session.ExpiresAt = time.Now().Add(s.options.MaxAge)

	// Set cookie / 쿠키 설정
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

	// Remove session from store / 저장소에서 세션 제거
	s.mu.Lock()
	delete(s.sessions, cookie.Value)
	s.mu.Unlock()

	// Clear cookie / 쿠키 지우기
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
func (s *SessionStore) generateSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		// Fallback to timestamp-based ID / 타임스탬프 기반 ID로 대체
		return base64.URLEncoding.EncodeToString([]byte(time.Now().String()))
	}
	return base64.URLEncoding.EncodeToString(b)
}

// cleanupExpiredSessions periodically removes expired sessions.
// cleanupExpiredSessions는 주기적으로 만료된 세션을 제거합니다.
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
