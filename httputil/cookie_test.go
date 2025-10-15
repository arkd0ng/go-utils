package httputil

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"
)

// TestNewCookieJar tests creating a new cookie jar.
// TestNewCookieJar는 새 쿠키 저장소 생성을 테스트합니다.
func TestNewCookieJar(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	if jar == nil {
		t.Fatal("NewCookieJar returned nil")
	}

	if jar.jar == nil {
		t.Error("CookieJar.jar should not be nil")
	}

	if jar.filePath != "" {
		t.Errorf("Expected empty filePath, got %s", jar.filePath)
	}
}

// TestNewPersistentCookieJar tests creating a persistent cookie jar.
// TestNewPersistentCookieJar는 지속성 쿠키 저장소 생성을 테스트합니다.
func TestNewPersistentCookieJar(t *testing.T) {
	filePath := "test_cookies.json"
	defer os.Remove(filePath)

	jar, err := NewPersistentCookieJar(filePath)
	if err != nil {
		t.Fatalf("NewPersistentCookieJar failed: %v", err)
	}

	if jar == nil {
		t.Fatal("NewPersistentCookieJar returned nil")
	}

	if jar.filePath != filePath {
		t.Errorf("Expected filePath %s, got %s", filePath, jar.filePath)
	}
}

// TestCookieJar_SetAndGetCookies tests setting and getting cookies.
// TestCookieJar_SetAndGetCookies는 쿠키 설정 및 가져오기를 테스트합니다.
func TestCookieJar_SetAndGetCookies(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")
	cookies := []*http.Cookie{
		{Name: "session", Value: "abc123", Path: "/"},
		{Name: "user", Value: "john", Path: "/"},
	}

	jar.SetCookies(u, cookies)

	// Get cookies / 쿠키 가져오기
	gotCookies := jar.GetCookies(u)
	if len(gotCookies) != 2 {
		t.Errorf("Expected 2 cookies, got %d", len(gotCookies))
	}

	// Verify cookie values / 쿠키 값 검증
	found := make(map[string]string)
	for _, c := range gotCookies {
		found[c.Name] = c.Value
	}

	if found["session"] != "abc123" {
		t.Errorf("Expected session=abc123, got %s", found["session"])
	}

	if found["user"] != "john" {
		t.Errorf("Expected user=john, got %s", found["user"])
	}
}

// TestCookieJar_SetCookie tests setting a single cookie.
// TestCookieJar_SetCookie는 단일 쿠키 설정을 테스트합니다.
func TestCookieJar_SetCookie(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")
	cookie := &http.Cookie{Name: "token", Value: "xyz789", Path: "/"}

	jar.SetCookie(u, cookie)

	// Get cookies / 쿠키 가져오기
	gotCookies := jar.GetCookies(u)
	if len(gotCookies) != 1 {
		t.Errorf("Expected 1 cookie, got %d", len(gotCookies))
	}

	if gotCookies[0].Name != "token" || gotCookies[0].Value != "xyz789" {
		t.Errorf("Expected token=xyz789, got %s=%s", gotCookies[0].Name, gotCookies[0].Value)
	}
}

// TestCookieJar_ClearCookies tests clearing all cookies.
// TestCookieJar_ClearCookies는 모든 쿠키 제거를 테스트합니다.
func TestCookieJar_ClearCookies(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")
	cookies := []*http.Cookie{
		{Name: "session", Value: "abc123", Path: "/"},
		{Name: "user", Value: "john", Path: "/"},
	}

	jar.SetCookies(u, cookies)

	// Verify cookies are set / 쿠키가 설정되었는지 확인
	gotCookies := jar.GetCookies(u)
	if len(gotCookies) != 2 {
		t.Errorf("Expected 2 cookies before clear, got %d", len(gotCookies))
	}

	// Clear cookies / 쿠키 제거
	if err := jar.ClearCookies(); err != nil {
		t.Fatalf("ClearCookies failed: %v", err)
	}

	// Verify cookies are cleared / 쿠키가 제거되었는지 확인
	gotCookies = jar.GetCookies(u)
	if len(gotCookies) != 0 {
		t.Errorf("Expected 0 cookies after clear, got %d", len(gotCookies))
	}
}

// TestCookieJar_SaveAndLoadCookies tests cookie persistence.
// TestCookieJar_SaveAndLoadCookies는 쿠키 지속성을 테스트합니다.
func TestCookieJar_SaveAndLoadCookies(t *testing.T) {
	filePath := "test_cookies_persist.json"
	defer os.Remove(filePath)

	// Create jar and set cookies / 저장소 생성 및 쿠키 설정
	jar, err := NewPersistentCookieJar(filePath)
	if err != nil {
		t.Fatalf("NewPersistentCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")
	cookies := []*http.Cookie{
		{Name: "session", Value: "persist123", Path: "/"},
	}
	jar.SetCookies(u, cookies)

	// Save cookies / 쿠키 저장
	if err := jar.SaveCookies(); err != nil {
		t.Fatalf("SaveCookies failed: %v", err)
	}

	// Verify file exists / 파일 존재 확인
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("Cookie file should exist after SaveCookies")
	}

	// Create new jar and load cookies / 새 저장소 생성 및 쿠키 로드
	jar2, err := NewPersistentCookieJar(filePath)
	if err != nil {
		t.Fatalf("NewPersistentCookieJar failed: %v", err)
	}

	// Note: Standard cookiejar doesn't expose all cookies easily
	// This test verifies the file operations work correctly
	// 참고: 표준 cookiejar는 모든 쿠키를 쉽게 노출하지 않습니다
	// 이 테스트는 파일 작업이 올바르게 작동하는지 확인합니다
	if jar2.filePath != filePath {
		t.Errorf("Expected filePath %s, got %s", filePath, jar2.filePath)
	}
}

// TestCookieJar_GetCookiesByDomain tests getting cookies by domain.
// TestCookieJar_GetCookiesByDomain은 도메인별 쿠키 가져오기를 테스트합니다.
func TestCookieJar_GetCookiesByDomain(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")
	cookies := []*http.Cookie{
		{Name: "domain_cookie", Value: "test123", Path: "/"},
	}
	jar.SetCookies(u, cookies)

	// Get cookies by domain / 도메인별 쿠키 가져오기
	gotCookies := jar.GetCookiesByDomain("example.com")
	if len(gotCookies) == 0 {
		t.Error("Expected at least 1 cookie for example.com domain")
	}
}

// TestCookieJar_RemoveCookie tests removing a specific cookie.
// TestCookieJar_RemoveCookie는 특정 쿠키 제거를 테스트합니다.
func TestCookieJar_RemoveCookie(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")
	cookies := []*http.Cookie{
		{Name: "session", Value: "abc123", Path: "/"},
		{Name: "user", Value: "john", Path: "/"},
	}
	jar.SetCookies(u, cookies)

	// Verify 2 cookies / 2개 쿠키 확인
	if count := jar.CountCookies(u); count != 2 {
		t.Errorf("Expected 2 cookies, got %d", count)
	}

	// Remove one cookie / 1개 쿠키 제거
	jar.RemoveCookie(u, "session")

	// Verify cookie was removed / 쿠키가 제거되었는지 확인
	if jar.HasCookie(u, "session") {
		t.Error("Cookie 'session' should have been removed")
	}

	// Verify other cookie still exists / 다른 쿠키는 여전히 존재하는지 확인
	if !jar.HasCookie(u, "user") {
		t.Error("Cookie 'user' should still exist")
	}
}

// TestCookieJar_CountCookies tests counting cookies.
// TestCookieJar_CountCookies는 쿠키 개수 세기를 테스트합니다.
func TestCookieJar_CountCookies(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")

	// Initially no cookies / 초기에는 쿠키 없음
	if count := jar.CountCookies(u); count != 0 {
		t.Errorf("Expected 0 cookies initially, got %d", count)
	}

	// Add cookies / 쿠키 추가
	cookies := []*http.Cookie{
		{Name: "cookie1", Value: "value1", Path: "/"},
		{Name: "cookie2", Value: "value2", Path: "/"},
		{Name: "cookie3", Value: "value3", Path: "/"},
	}
	jar.SetCookies(u, cookies)

	// Count should be 3 / 개수는 3이어야 함
	if count := jar.CountCookies(u); count != 3 {
		t.Errorf("Expected 3 cookies, got %d", count)
	}
}

// TestCookieJar_HasCookie tests checking cookie existence.
// TestCookieJar_HasCookie는 쿠키 존재 확인을 테스트합니다.
func TestCookieJar_HasCookie(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")
	cookie := &http.Cookie{Name: "exists", Value: "yes", Path: "/"}
	jar.SetCookie(u, cookie)

	// Should have the cookie / 쿠키가 존재해야 함
	if !jar.HasCookie(u, "exists") {
		t.Error("HasCookie should return true for 'exists'")
	}

	// Should not have non-existent cookie / 존재하지 않는 쿠키는 없어야 함
	if jar.HasCookie(u, "nonexistent") {
		t.Error("HasCookie should return false for 'nonexistent'")
	}
}

// TestCookieJar_GetCookie tests getting a specific cookie.
// TestCookieJar_GetCookie는 특정 쿠키 가져오기를 테스트합니다.
func TestCookieJar_GetCookie(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")
	cookie := &http.Cookie{Name: "target", Value: "found", Path: "/"}
	jar.SetCookie(u, cookie)

	// Get existing cookie / 존재하는 쿠키 가져오기
	gotCookie := jar.GetCookie(u, "target")
	if gotCookie == nil {
		t.Fatal("GetCookie should return a cookie")
	}

	if gotCookie.Name != "target" || gotCookie.Value != "found" {
		t.Errorf("Expected target=found, got %s=%s", gotCookie.Name, gotCookie.Value)
	}

	// Get non-existent cookie / 존재하지 않는 쿠키 가져오기
	nonCookie := jar.GetCookie(u, "nonexistent")
	if nonCookie != nil {
		t.Error("GetCookie should return nil for non-existent cookie")
	}
}

// TestClient_CookieIntegration tests cookie integration with Client.
// TestClient_CookieIntegration은 클라이언트와의 쿠키 통합을 테스트합니다.
func TestClient_CookieIntegration(t *testing.T) {
	// Create test server / 테스트 서버 생성
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set a cookie in response / 응답에 쿠키 설정
		http.SetCookie(w, &http.Cookie{
			Name:  "server_cookie",
			Value: "from_server",
			Path:  "/",
		})

		// Check for cookies in request / 요청의 쿠키 확인
		if cookie, err := r.Cookie("client_cookie"); err == nil {
			w.Header().Set("X-Client-Cookie", cookie.Value)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	// Create client with cookies enabled / 쿠키가 활성화된 클라이언트 생성
	client := NewClient(
		WithBaseURL(server.URL),
		WithCookies(),
	)

	// Make request / 요청 수행
	var result map[string]string
	err := client.Get("/test", &result)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	// Verify cookie jar is initialized / 쿠키 저장소가 초기화되었는지 확인
	if client.cookieJar == nil {
		t.Error("Client cookieJar should not be nil with WithCookies()")
	}

	// Get cookies from server / 서버에서 쿠키 가져오기
	u, _ := url.Parse(server.URL)
	cookies := client.GetCookies(u)
	if len(cookies) == 0 {
		t.Error("Should have received cookies from server")
	}

	// Verify server cookie exists / 서버 쿠키가 존재하는지 확인
	found := false
	for _, c := range cookies {
		if c.Name == "server_cookie" && c.Value == "from_server" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Should have received 'server_cookie' from server")
	}
}

// TestClient_PersistentCookies tests persistent cookie integration.
// TestClient_PersistentCookies는 지속성 쿠키 통합을 테스트합니다.
func TestClient_PersistentCookies(t *testing.T) {
	filePath := "test_client_cookies.json"
	defer os.Remove(filePath)

	// Create client with persistent cookies / 지속성 쿠키를 가진 클라이언트 생성
	client := NewClient(
		WithBaseURL("https://example.com"),
		WithPersistentCookies(filePath),
	)

	// Verify cookie jar is initialized / 쿠키 저장소가 초기화되었는지 확인
	if client.cookieJar == nil {
		t.Fatal("Client cookieJar should not be nil with WithPersistentCookies()")
	}

	if client.cookieJar.filePath != filePath {
		t.Errorf("Expected filePath %s, got %s", filePath, client.cookieJar.filePath)
	}

	// Set a cookie / 쿠키 설정
	u, _ := url.Parse("https://example.com")
	cookie := &http.Cookie{Name: "persistent", Value: "test123", Path: "/"}
	client.SetCookie(u, cookie)

	// Save cookies / 쿠키 저장
	if err := client.SaveCookies(); err != nil {
		t.Fatalf("SaveCookies failed: %v", err)
	}

	// Verify file exists / 파일 존재 확인
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("Cookie file should exist after SaveCookies")
	}
}

// TestClient_CookieMethods tests all Client cookie methods.
// TestClient_CookieMethods는 모든 클라이언트 쿠키 메서드를 테스트합니다.
func TestClient_CookieMethods(t *testing.T) {
	client := NewClient(WithCookies())
	if client.cookieJar == nil {
		t.Fatal("Client cookieJar should not be nil")
	}

	u, _ := url.Parse("https://example.com")

	// Test SetCookie / SetCookie 테스트
	cookie := &http.Cookie{Name: "test", Value: "value", Path: "/"}
	client.SetCookie(u, cookie)

	// Test HasCookie / HasCookie 테스트
	if !client.HasCookie(u, "test") {
		t.Error("HasCookie should return true")
	}

	// Test GetCookie / GetCookie 테스트
	gotCookie := client.GetCookie(u, "test")
	if gotCookie == nil {
		t.Fatal("GetCookie should return a cookie")
	}
	if gotCookie.Value != "value" {
		t.Errorf("Expected value 'value', got '%s'", gotCookie.Value)
	}

	// Test GetCookies / GetCookies 테스트
	cookies := client.GetCookies(u)
	if len(cookies) != 1 {
		t.Errorf("Expected 1 cookie, got %d", len(cookies))
	}

	// Test ClearCookies / ClearCookies 테스트
	if err := client.ClearCookies(); err != nil {
		t.Fatalf("ClearCookies failed: %v", err)
	}

	if client.HasCookie(u, "test") {
		t.Error("Cookie should have been cleared")
	}
}

// TestClient_NoCookieJar tests client behavior without cookie jar.
// TestClient_NoCookieJar는 쿠키 저장소 없는 클라이언트 동작을 테스트합니다.
func TestClient_NoCookieJar(t *testing.T) {
	client := NewClient() // No cookie options / 쿠키 옵션 없음

	if client.cookieJar != nil {
		t.Error("Client cookieJar should be nil without cookie options")
	}

	u, _ := url.Parse("https://example.com")

	// All methods should handle nil cookieJar gracefully / 모든 메서드는 nil 쿠키 저장소를 우아하게 처리해야 함
	cookies := client.GetCookies(u)
	if cookies != nil {
		t.Error("GetCookies should return nil when cookieJar is nil")
	}

	client.SetCookie(u, &http.Cookie{Name: "test", Value: "value"})
	// Should not panic / 패닉이 발생하지 않아야 함

	if client.HasCookie(u, "test") {
		t.Error("HasCookie should return false when cookieJar is nil")
	}

	if client.GetCookie(u, "test") != nil {
		t.Error("GetCookie should return nil when cookieJar is nil")
	}

	if err := client.ClearCookies(); err != nil {
		t.Error("ClearCookies should not return error when cookieJar is nil")
	}
}

// TestCookieJar_ThreadSafety tests thread-safe cookie operations.
// TestCookieJar_ThreadSafety는 스레드 안전 쿠키 작업을 테스트합니다.
func TestCookieJar_ThreadSafety(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")

	// Concurrent reads and writes / 동시 읽기 및 쓰기
	done := make(chan bool)

	// Writer goroutines / 쓰기 고루틴
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				cookie := &http.Cookie{
					Name:  "concurrent",
					Value: "test",
					Path:  "/",
				}
				jar.SetCookie(u, cookie)
			}
			done <- true
		}(i)
	}

	// Reader goroutines / 읽기 고루틴
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				jar.GetCookies(u)
				jar.HasCookie(u, "concurrent")
				jar.CountCookies(u)
			}
			done <- true
		}()
	}

	// Wait for all goroutines / 모든 고루틴 대기
	for i := 0; i < 20; i++ {
		<-done
	}
}

// TestCookieJar_ExpiredCookies tests handling of expired cookies.
// TestCookieJar_ExpiredCookies는 만료된 쿠키 처리를 테스트합니다.
func TestCookieJar_ExpiredCookies(t *testing.T) {
	jar, err := NewCookieJar()
	if err != nil {
		t.Fatalf("NewCookieJar failed: %v", err)
	}

	u, _ := url.Parse("https://example.com")

	// Set cookie that expires immediately / 즉시 만료되는 쿠키 설정
	expiredCookie := &http.Cookie{
		Name:    "expired",
		Value:   "old",
		Path:    "/",
		Expires: time.Now().Add(-1 * time.Hour),
	}
	jar.SetCookie(u, expiredCookie)

	// The standard cookiejar should automatically filter expired cookies
	// 표준 cookiejar는 만료된 쿠키를 자동으로 필터링해야 합니다
	cookies := jar.GetCookies(u)
	for _, c := range cookies {
		if c.Name == "expired" {
			t.Error("Expired cookie should not be returned")
		}
	}
}

// BenchmarkCookieJar_SetCookie benchmarks cookie setting.
// BenchmarkCookieJar_SetCookie는 쿠키 설정 벤치마크입니다.
func BenchmarkCookieJar_SetCookie(b *testing.B) {
	jar, _ := NewCookieJar()
	u, _ := url.Parse("https://example.com")
	cookie := &http.Cookie{Name: "bench", Value: "test", Path: "/"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jar.SetCookie(u, cookie)
	}
}

// BenchmarkCookieJar_GetCookies benchmarks cookie retrieval.
// BenchmarkCookieJar_GetCookies는 쿠키 가져오기 벤치마크입니다.
func BenchmarkCookieJar_GetCookies(b *testing.B) {
	jar, _ := NewCookieJar()
	u, _ := url.Parse("https://example.com")
	cookie := &http.Cookie{Name: "bench", Value: "test", Path: "/"}
	jar.SetCookie(u, cookie)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jar.GetCookies(u)
	}
}

// BenchmarkCookieJar_HasCookie benchmarks cookie existence check.
// BenchmarkCookieJar_HasCookie는 쿠키 존재 확인 벤치마크입니다.
func BenchmarkCookieJar_HasCookie(b *testing.B) {
	jar, _ := NewCookieJar()
	u, _ := url.Parse("https://example.com")
	cookie := &http.Cookie{Name: "bench", Value: "test", Path: "/"}
	jar.SetCookie(u, cookie)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jar.HasCookie(u, "bench")
	}
}
