package websvrutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestNewSessionStore tests creating a new session store / 새 세션 저장소 생성 테스트
func TestNewSessionStore(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)

	if store == nil {
		t.Fatal("Expected non-nil session store")
	}

	if store.options.CookieName != "sessionid" {
		t.Errorf("Expected cookie name 'sessionid', got '%s'", store.options.CookieName)
	}

	if store.options.MaxAge != 24*time.Hour {
		t.Errorf("Expected MaxAge 24 hours, got %v", store.options.MaxAge)
	}

	if store.Count() != 0 {
		t.Errorf("Expected 0 sessions, got %d", store.Count())
	}
}

// TestSessionStoreNew tests creating a new session / 새 세션 생성 테스트
func TestSessionStoreNew(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)

	session := store.New()

	if session == nil {
		t.Fatal("Expected non-nil session")
	}

	if session.ID == "" {
		t.Error("Expected non-empty session ID")
	}

	if session.Data == nil {
		t.Error("Expected non-nil session data map")
	}

	if session.CreatedAt.IsZero() {
		t.Error("Expected non-zero CreatedAt")
	}

	if session.ExpiresAt.IsZero() {
		t.Error("Expected non-zero ExpiresAt")
	}

	if store.Count() != 1 {
		t.Errorf("Expected 1 session, got %d", store.Count())
	}
}

// TestSessionSetGet tests session set and get operations / 세션 설정 및 가져오기 작업 테스트
func TestSessionSetGet(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)
	session := store.New()

	// Test string value / 문자열 값 테스트
	session.Set("username", "john")
	value, exists := session.Get("username")
	if !exists {
		t.Error("Expected username to exist")
	}
	if value != "john" {
		t.Errorf("Expected 'john', got '%v'", value)
	}

	// Test int value / int 값 테스트
	session.Set("count", 42)
	value, exists = session.Get("count")
	if !exists {
		t.Error("Expected count to exist")
	}
	if value != 42 {
		t.Errorf("Expected 42, got %v", value)
	}

	// Test bool value / bool 값 테스트
	session.Set("active", true)
	value, exists = session.Get("active")
	if !exists {
		t.Error("Expected active to exist")
	}
	if value != true {
		t.Errorf("Expected true, got %v", value)
	}

	// Test non-existent key / 존재하지 않는 키 테스트
	_, exists = session.Get("nonexistent")
	if exists {
		t.Error("Expected nonexistent key to not exist")
	}
}

// TestSessionGetTyped tests typed getter methods / 타입별 getter 메서드 테스트
func TestSessionGetTyped(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)
	session := store.New()

	session.Set("username", "alice")
	session.Set("age", 25)
	session.Set("active", true)

	// Test GetString / GetString 테스트
	username := session.GetString("username")
	if username != "alice" {
		t.Errorf("Expected 'alice', got '%s'", username)
	}

	// Test GetString with non-existent key / 존재하지 않는 키로 GetString 테스트
	missing := session.GetString("missing")
	if missing != "" {
		t.Errorf("Expected empty string, got '%s'", missing)
	}

	// Test GetInt / GetInt 테스트
	age := session.GetInt("age")
	if age != 25 {
		t.Errorf("Expected 25, got %d", age)
	}

	// Test GetInt with non-existent key / 존재하지 않는 키로 GetInt 테스트
	missingInt := session.GetInt("missing")
	if missingInt != 0 {
		t.Errorf("Expected 0, got %d", missingInt)
	}

	// Test GetBool / GetBool 테스트
	active := session.GetBool("active")
	if active != true {
		t.Errorf("Expected true, got %v", active)
	}

	// Test GetBool with non-existent key / 존재하지 않는 키로 GetBool 테스트
	missingBool := session.GetBool("missing")
	if missingBool != false {
		t.Errorf("Expected false, got %v", missingBool)
	}
}

// TestSessionDelete tests deleting session values / 세션 값 삭제 테스트
func TestSessionDelete(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)
	session := store.New()

	session.Set("key1", "value1")
	session.Set("key2", "value2")

	// Verify key1 exists / key1 존재 확인
	_, exists := session.Get("key1")
	if !exists {
		t.Error("Expected key1 to exist")
	}

	// Delete key1 / key1 삭제
	session.Delete("key1")

	// Verify key1 is deleted / key1 삭제 확인
	_, exists = session.Get("key1")
	if exists {
		t.Error("Expected key1 to be deleted")
	}

	// Verify key2 still exists / key2 여전히 존재 확인
	_, exists = session.Get("key2")
	if !exists {
		t.Error("Expected key2 to still exist")
	}
}

// TestSessionClear tests clearing all session values / 모든 세션 값 지우기 테스트
func TestSessionClear(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)
	session := store.New()

	session.Set("key1", "value1")
	session.Set("key2", "value2")
	session.Set("key3", "value3")

	// Clear all values / 모든 값 지우기
	session.Clear()

	// Verify all keys are cleared / 모든 키 지워짐 확인
	_, exists1 := session.Get("key1")
	_, exists2 := session.Get("key2")
	_, exists3 := session.Get("key3")

	if exists1 || exists2 || exists3 {
		t.Error("Expected all keys to be cleared")
	}

	if len(session.Data) != 0 {
		t.Errorf("Expected empty data map, got %d items", len(session.Data))
	}
}

// TestSessionStoreGetExisting tests getting an existing session from cookie / 쿠키에서 기존 세션 가져오기 테스트
func TestSessionStoreGetExisting(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)

	// Create a session / 세션 생성
	originalSession := store.New()
	originalSession.Set("username", "bob")

	// Create request with session cookie / 세션 쿠키가 있는 요청 생성
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  opts.CookieName,
		Value: originalSession.ID,
	})

	// Get session from request / 요청에서 세션 가져오기
	session, err := store.Get(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if session.ID != originalSession.ID {
		t.Errorf("Expected session ID '%s', got '%s'", originalSession.ID, session.ID)
	}

	// Verify data is preserved / 데이터 보존 확인
	username := session.GetString("username")
	if username != "bob" {
		t.Errorf("Expected username 'bob', got '%s'", username)
	}
}

// TestSessionStoreGetNew tests creating a new session when none exists / 세션이 없을 때 새 세션 생성 테스트
func TestSessionStoreGetNew(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)

	// Create request without session cookie / 세션 쿠키가 없는 요청 생성
	req := httptest.NewRequest("GET", "/", nil)

	// Get session from request / 요청에서 세션 가져오기
	session, err := store.Get(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if session == nil {
		t.Fatal("Expected non-nil session")
	}

	if session.ID == "" {
		t.Error("Expected non-empty session ID")
	}
}

// TestSessionStoreSave tests saving a session and setting cookie / 세션 저장 및 쿠키 설정 테스트
func TestSessionStoreSave(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)
	session := store.New()

	w := httptest.NewRecorder()
	store.Save(w, session)

	// Check cookie is set / 쿠키 설정 확인
	cookies := w.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatal("Expected cookie to be set")
	}

	cookie := cookies[0]
	if cookie.Name != opts.CookieName {
		t.Errorf("Expected cookie name '%s', got '%s'", opts.CookieName, cookie.Name)
	}

	if cookie.Value != session.ID {
		t.Errorf("Expected cookie value '%s', got '%s'", session.ID, cookie.Value)
	}

	if cookie.Path != opts.Path {
		t.Errorf("Expected cookie path '%s', got '%s'", opts.Path, cookie.Path)
	}

	if cookie.HttpOnly != opts.HttpOnly {
		t.Errorf("Expected HttpOnly %v, got %v", opts.HttpOnly, cookie.HttpOnly)
	}
}

// TestSessionStoreDestroy tests destroying a session / 세션 파괴 테스트
func TestSessionStoreDestroy(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)
	session := store.New()

	// Create request with session cookie / 세션 쿠키가 있는 요청 생성
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  opts.CookieName,
		Value: session.ID,
	})

	w := httptest.NewRecorder()

	// Destroy session / 세션 파괴
	err := store.Destroy(w, req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify session is removed from store / 저장소에서 세션 제거 확인
	if store.Count() != 0 {
		t.Errorf("Expected 0 sessions, got %d", store.Count())
	}

	// Verify cookie is cleared / 쿠키 지워짐 확인
	cookies := w.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatal("Expected cookie to be set for clearing")
	}

	cookie := cookies[0]
	if cookie.MaxAge != -1 {
		t.Errorf("Expected MaxAge -1, got %d", cookie.MaxAge)
	}
}

// TestSessionExpiration tests session expiration / 세션 만료 테스트
func TestSessionExpiration(t *testing.T) {
	opts := DefaultSessionOptions()
	opts.MaxAge = 100 * time.Millisecond // Short expiration for testing / 테스트용 짧은 만료 시간
	store := NewSessionStore(opts)

	session := store.New()

	// Wait for session to expire / 세션 만료 대기
	time.Sleep(150 * time.Millisecond)

	// Try to get expired session / 만료된 세션 가져오기 시도
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  opts.CookieName,
		Value: session.ID,
	})

	newSession, err := store.Get(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should get a new session (different ID) / 새 세션을 얻어야 함 (다른 ID)
	if newSession.ID == session.ID {
		t.Error("Expected new session ID for expired session")
	}
}

// TestSessionCleanup tests automatic cleanup of expired sessions / 만료된 세션 자동 정리 테스트
func TestSessionCleanup(t *testing.T) {
	opts := DefaultSessionOptions()
	opts.MaxAge = 100 * time.Millisecond    // Short expiration / 짧은 만료 시간
	opts.CleanupTime = 150 * time.Millisecond // Short cleanup interval / 짧은 정리 간격
	store := NewSessionStore(opts)

	// Create multiple sessions / 여러 세션 생성
	session1 := store.New()
	session2 := store.New()
	session3 := store.New()

	if store.Count() != 3 {
		t.Fatalf("Expected 3 sessions, got %d", store.Count())
	}

	// Wait for sessions to expire and cleanup to run / 세션 만료 및 정리 실행 대기
	time.Sleep(300 * time.Millisecond)

	// All sessions should be cleaned up / 모든 세션이 정리되어야 함
	if store.Count() != 0 {
		t.Errorf("Expected 0 sessions after cleanup, got %d", store.Count())
	}

	_ = session1
	_ = session2
	_ = session3
}

// TestSessionConcurrency tests concurrent access to session / 세션에 대한 동시 액세스 테스트
func TestSessionConcurrency(t *testing.T) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)
	session := store.New()

	done := make(chan bool)

	// Concurrent writes / 동시 쓰기
	for i := 0; i < 100; i++ {
		go func(n int) {
			session.Set("key", n)
			done <- true
		}(i)
	}

	// Wait for all goroutines / 모든 고루틴 대기
	for i := 0; i < 100; i++ {
		<-done
	}

	// Verify session is still valid / 세션 여전히 유효 확인
	_, exists := session.Get("key")
	if !exists {
		t.Error("Expected key to exist after concurrent writes")
	}
}

// BenchmarkSessionSet benchmarks session Set operation / 세션 Set 작업 벤치마크
func BenchmarkSessionSet(b *testing.B) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)
	session := store.New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		session.Set("key", "value")
	}
}

// BenchmarkSessionGet benchmarks session Get operation / 세션 Get 작업 벤치마크
func BenchmarkSessionGet(b *testing.B) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)
	session := store.New()
	session.Set("key", "value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		session.Get("key")
	}
}

// BenchmarkSessionStoreNew benchmarks creating new sessions / 새 세션 생성 벤치마크
func BenchmarkSessionStoreNew(b *testing.B) {
	opts := DefaultSessionOptions()
	store := NewSessionStore(opts)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.New()
	}
}
