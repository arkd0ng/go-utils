// Package main demonstrates the httputil package capabilities
// main 패키지는 httputil 패키지의 기능을 보여줍니다
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/arkd0ng/go-utils/httputil"
)

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE DATA STRUCTURES / 예제 데이터 구조
// ═══════════════════════════════════════════════════════════════════════════

// User represents a user resource / User는 사용자 리소스를 나타냅니다
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Post represents a post resource / Post는 게시물 리소스를 나타냅니다
type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 1: BASIC HTTP METHODS (Phase 1) / 예제 1: 기본 HTTP 메서드 (1단계)
// ═══════════════════════════════════════════════════════════════════════════

func exampleBasicHTTPMethods() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 1: Basic HTTP Methods (Phase 1)")
	fmt.Println("예제 1: 기본 HTTP 메서드 (1단계)")
	fmt.Println(strings.Repeat("=", 80))

	// GET request / GET 요청
	fmt.Println("\n1.1 Simple GET Request / 간단한 GET 요청:")
	var users []User
	err := httputil.Get("https://jsonplaceholder.typicode.com/users", &users)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   → Fetched %d users / %d명의 사용자를 가져왔습니다\n", len(users))

	// GET with options / 옵션이 있는 GET
	fmt.Println("\n1.2 GET with Options / 옵션이 있는 GET:")
	var user User
	err = httputil.Get("https://jsonplaceholder.typicode.com/users/1", &user,
		httputil.WithTimeout(10*time.Second),
		httputil.WithHeader("X-Custom-Header", "example-value"))
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   → User: %s (%s)\n", user.Name, user.Email)

	// POST request / POST 요청
	fmt.Println("\n1.3 POST Request / POST 요청:")
	newPost := Post{
		UserID: 1,
		Title:  "Test Post from httputil",
		Body:   "This is a test post created using httputil package.",
	}
	var createdPost Post
	err = httputil.Post("https://jsonplaceholder.typicode.com/posts", newPost, &createdPost)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   → Created Post ID: %d\n", createdPost.ID)

	// PUT request / PUT 요청
	fmt.Println("\n1.4 PUT Request / PUT 요청:")
	updatedPost := Post{
		ID:     1,
		UserID: 1,
		Title:  "Updated Title",
		Body:   "Updated body content.",
	}
	var resultPost Post
	err = httputil.Put("https://jsonplaceholder.typicode.com/posts/1", updatedPost, &resultPost)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   → Updated Post: %s\n", resultPost.Title)

	// DELETE request / DELETE 요청
	fmt.Println("\n1.5 DELETE Request / DELETE 요청:")
	err = httputil.Delete("https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println("   → Post deleted successfully / 게시물이 성공적으로 삭제되었습니다")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 2: CLIENT WITH BASE URL (Phase 1) / 예제 2: Base URL을 가진 클라이언트 (1단계)
// ═══════════════════════════════════════════════════════════════════════════

func exampleClientWithBaseURL() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 2: Client with Base URL (Phase 1)")
	fmt.Println("예제 2: Base URL을 가진 클라이언트 (1단계)")
	fmt.Println(strings.Repeat("=", 80))

	// Create configured client / 설정된 클라이언트 생성
	client := httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithTimeout(30*time.Second),
		httputil.WithRetry(3),
	)

	// Multiple requests using the client / 클라이언트를 사용한 여러 요청
	fmt.Println("\n2.1 Fetching Multiple Resources / 여러 리소스 가져오기:")

	var users []User
	err := client.Get("/users?_limit=3", &users)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   → Fetched %d users / %d명의 사용자\n", len(users))

	var posts []Post
	err = client.Get("/posts?_limit=5", &posts)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   → Fetched %d posts / %d개의 게시물\n", len(posts))
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 3: CONTEXT AND TIMEOUT (Phase 1) / 예제 3: Context 및 타임아웃 (1단계)
// ═══════════════════════════════════════════════════════════════════════════

func exampleContextAndTimeout() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 3: Context and Timeout (Phase 1)")
	fmt.Println("예제 3: Context 및 타임아웃 (1단계)")
	fmt.Println(strings.Repeat("=", 80))

	// Create context with timeout / 타임아웃이 있는 context 생성
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("\n3.1 Request with Context / Context가 있는 요청:")
	var users []User
	err := httputil.GetContext(ctx, "https://jsonplaceholder.typicode.com/users", &users)
	if err != nil {
		if httputil.IsTimeoutError(err) {
			fmt.Println("   → Request timed out / 요청 타임아웃")
		} else {
			log.Printf("Error: %v", err)
		}
		return
	}
	fmt.Printf("   → Successfully fetched %d users / %d명의 사용자를 성공적으로 가져왔습니다\n", len(users))
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 4: ERROR HANDLING (Phase 1) / 예제 4: 에러 처리 (1단계)
// ═══════════════════════════════════════════════════════════════════════════

func exampleErrorHandling() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 4: Error Handling (Phase 1)")
	fmt.Println("예제 4: 에러 처리 (1단계)")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Println("\n4.1 Handling 404 Not Found / 404 Not Found 처리:")
	var user User
	err := httputil.Get("https://jsonplaceholder.typicode.com/users/999999", &user)
	if err != nil {
		if httputil.IsHTTPError(err) {
			statusCode := httputil.GetStatusCode(err)
			fmt.Printf("   → HTTP Error: Status Code %d\n", statusCode)
			if statusCode == 404 {
				fmt.Println("   → Resource not found / 리소스를 찾을 수 없습니다")
			}
		} else {
			log.Printf("Error: %v", err)
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 5: RESPONSE HELPERS (Phase 2) / 예제 5: 응답 헬퍼 (2단계)
// ═══════════════════════════════════════════════════════════════════════════

func exampleResponseHelpers() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 5: Response Helpers (Phase 2)")
	fmt.Println("예제 5: 응답 헬퍼 (2단계)")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Println("\n5.1 Using DoRaw for Response Inspection / 응답 검사를 위한 DoRaw 사용:")
	resp, err := httputil.DoRaw("GET", "https://jsonplaceholder.typicode.com/users/1", nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Check status with helper methods / 헬퍼 메서드로 상태 확인
	fmt.Printf("   → Status Code: %d\n", resp.StatusCode)
	fmt.Printf("   → Is Success: %v\n", resp.IsSuccess())
	fmt.Printf("   → Is OK (200): %v\n", resp.IsOK())

	// Access headers / 헤더 접근
	fmt.Printf("   → Content-Type: %s\n", resp.ContentType())

	// Access body in different formats / 다양한 형식으로 본문 접근
	bodyString := resp.String()
	fmt.Printf("   → Body length: %d bytes\n", len(bodyString))

	// Decode JSON / JSON 디코딩
	var user User
	err = resp.JSON(&user)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return
	}
	fmt.Printf("   → User: %s (%s)\n", user.Name, user.Email)
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 6: FILE DOWNLOAD (Phase 3) / 예제 6: 파일 다운로드 (3단계)
// ═══════════════════════════════════════════════════════════════════════════

func exampleFileDownload() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 6: File Download (Phase 3)")
	fmt.Println("예제 6: 파일 다운로드 (3단계)")
	fmt.Println(strings.Repeat("=", 80))

	// Note: This example demonstrates the API but uses a small file
	// 참고: 이 예제는 API를 보여주지만 작은 파일을 사용합니다
	fmt.Println("\n6.1 Download to Memory / 메모리로 다운로드:")
	data, err := httputil.Download("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   → Downloaded %d bytes / %d 바이트 다운로드됨\n", len(data))

	// Download with progress (commented out as it requires a writable file path)
	// 진행 상황과 함께 다운로드 (쓰기 가능한 파일 경로가 필요하므로 주석 처리됨)
	/*
		fmt.Println("\n6.2 Download with Progress / 진행 상황과 함께 다운로드:")
		ctx := context.Background()
		err = httputil.DownloadFileContext(ctx,
			"https://example.com/large-file.zip",
			"./downloads/file.zip",
			func(bytesRead, totalBytes int64) {
				if totalBytes > 0 {
					progress := float64(bytesRead) / float64(totalBytes) * 100
					fmt.Printf("\r   → Progress: %.2f%%", progress)
				}
			})
		if err != nil {
			log.Printf("Error: %v", err)
		}
		fmt.Println()
	*/
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 7: FILE UPLOAD (Phase 3) / 예제 7: 파일 업로드 (3단계)
// ═══════════════════════════════════════════════════════════════════════════

func exampleFileUpload() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 7: File Upload (Phase 3)")
	fmt.Println("예제 7: 파일 업로드 (3단계)")
	fmt.Println(strings.Repeat("=", 80))

	// Note: This example demonstrates the API usage
	// 참고: 이 예제는 API 사용법을 보여줍니다
	fmt.Println("\n7.1 File Upload API Usage / 파일 업로드 API 사용법:")
	fmt.Println(`
   // Upload single file / 단일 파일 업로드
   var result map[string]interface{}
   err := httputil.UploadFile(
       "https://api.example.com/upload",
       "document",
       "./files/report.pdf",
       &result)

   // Upload multiple files / 여러 파일 업로드
   err = httputil.UploadFiles(
       "https://api.example.com/upload-multiple",
       map[string]string{
           "file1": "./images/image1.jpg",
           "file2": "./images/image2.jpg",
       },
       &result)
   `)
	fmt.Println("   → File upload methods demonstrated / 파일 업로드 메서드 시연됨")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 8: URL BUILDER (Phase 4) / 예제 8: URL 빌더 (4단계)
// ═══════════════════════════════════════════════════════════════════════════

func exampleURLBuilder() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 8: URL Builder (Phase 4)")
	fmt.Println("예제 8: URL 빌더 (4단계)")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Println("\n8.1 Building URLs with Fluent API / Fluent API로 URL 구축:")

	// Build complex URL / 복잡한 URL 구축
	includeInactive := false
	url := httputil.NewURL("https://jsonplaceholder.typicode.com").
		Path("posts").
		Param("userId", "1").
		Param("_limit", "5").
		ParamIf(includeInactive, "status", "inactive").
		Build()

	fmt.Printf("   → Built URL: %s\n", url)

	// Use the built URL / 구축된 URL 사용
	var posts []Post
	err := httputil.Get(url, &posts)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   → Fetched %d posts / %d개의 게시물을 가져왔습니다\n", len(posts))

	fmt.Println("\n8.2 URL Utility Functions / URL 유틸리티 함수:")

	// Join URL parts / URL 부분 결합
	joinedURL := httputil.JoinURL("https://jsonplaceholder.typicode.com", "users", "1")
	fmt.Printf("   → Joined URL: %s\n", joinedURL)

	// Get domain / 도메인 가져오기
	domain, _ := httputil.GetDomain(joinedURL)
	fmt.Printf("   → Domain: %s\n", domain)

	// Check if absolute / 절대 URL인지 확인
	isAbsolute := httputil.IsAbsoluteURL(joinedURL)
	fmt.Printf("   → Is Absolute: %v\n", isAbsolute)
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 9: FORM BUILDER (Phase 4) / 예제 9: Form 빌더 (4단계)
// ═══════════════════════════════════════════════════════════════════════════

func exampleFormBuilder() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 9: Form Builder (Phase 4)")
	fmt.Println("예제 9: Form 빌더 (4단계)")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Println("\n9.1 Building Forms with Fluent API / Fluent API로 폼 구축:")

	// Build form with conditional fields / 조건부 필드가 있는 폼 구축
	hasPromoCode := true
	form := httputil.NewForm().
		Set("username", "testuser").
		Set("email", "test@example.com").
		Set("age", "30").
		AddIf(hasPromoCode, "promo_code", "SAVE20").
		AddIf(false, "referrer", "none").
		AddMultiple("tags", "go", "http", "api")

	// Check if field exists / 필드 존재 확인
	if form.Has("promo_code") {
		fmt.Println("   → Promo code applied / 프로모 코드 적용됨")
	}

	// Get form data / 폼 데이터 가져오기
	formMap := form.Map()
	fmt.Printf("   → Form fields: %d\n", len(formMap))
	for key, value := range formMap {
		fmt.Printf("      • %s: %s\n", key, value)
	}

	// Encode form / 폼 인코딩
	encoded := form.Encode()
	fmt.Printf("   → Encoded length: %d bytes\n", len(encoded))
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 10: COMPREHENSIVE WORKFLOW / 예제 10: 종합 워크플로우
// ═══════════════════════════════════════════════════════════════════════════

func exampleComprehensiveWorkflow() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("EXAMPLE 10: Comprehensive Workflow (All Phases)")
	fmt.Println("예제 10: 종합 워크플로우 (모든 단계)")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Println("\n10.1 Building a Complete API Client / 완전한 API 클라이언트 구축:")

	// Create configured client / 설정된 클라이언트 생성
	client := httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithTimeout(30*time.Second),
		httputil.WithRetry(3),
		httputil.WithUserAgent("httputil-example/1.0"),
	)

	// Use URL builder to create query / URL 빌더를 사용하여 쿼리 생성
	url := httputil.NewURL("").
		Path("posts").
		Param("userId", "1").
		Param("_limit", "3").
		Build()

	// Make request with raw response / 원시 응답으로 요청 실행
	resp, err := client.DoRaw("GET", url, nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Check response status / 응답 상태 확인
	if !resp.IsSuccess() {
		fmt.Println("   ✗ Request failed / 요청 실패")
		return
	}
	fmt.Println("   ✓ Request successful / 요청 성공")

	// Decode response / 응답 디코딩
	var posts []Post
	err = resp.JSON(&posts)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("   → Fetched %d posts:\n", len(posts))
	for i, post := range posts {
		fmt.Printf("      %d. %s (ID: %d)\n", i+1, post.Title, post.ID)
	}

	fmt.Println("\n   → All phases working together successfully!")
	fmt.Println("   → 모든 단계가 함께 성공적으로 작동합니다!")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 11: COOKIE MANAGEMENT (PHASE 5a)
// 예제 11: 쿠키 관리 (Phase 5a)
// ═══════════════════════════════════════════════════════════════════════════
func exampleCookieManagement() {
	fmt.Println("\n" + strings.Repeat("─", 80))
	fmt.Println("Example 11: Cookie Management (Phase 5a)")
	fmt.Println("예제 11: 쿠키 관리 (Phase 5a)")
	fmt.Println(strings.Repeat("─", 80))

	// Create test server that handles cookies / 쿠키를 처리하는 테스트 서버 생성
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set cookies in response / 응답에 쿠키 설정
		http.SetCookie(w, &http.Cookie{
			Name:  "session_id",
			Value: "abc123xyz",
			Path:  "/",
		})
		http.SetCookie(w, &http.Cookie{
			Name:  "user_pref",
			Value: "dark_mode",
			Path:  "/",
		})

		// Check for cookies in request / 요청의 쿠키 확인
		if cookie, err := r.Cookie("session_id"); err == nil {
			w.Header().Set("X-Session-Found", cookie.Value)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Cookie test endpoint",
		})
	}))
	defer ts.Close()

	fmt.Println("\n1. In-Memory Cookie Jar (temporary) / 메모리 내 쿠키 저장소 (임시)")
	fmt.Println("   ─────────────────────────────────────────────────────")

	// Create client with in-memory cookies / 메모리 내 쿠키를 사용하는 클라이언트 생성
	client1 := httputil.NewClient(
		httputil.WithBaseURL(ts.URL),
		httputil.WithCookies(), // Enable in-memory cookie jar / 메모리 내 쿠키 저장소 활성화
	)

	// Make first request / 첫 번째 요청
	var result1 map[string]string
	err := client1.Get("/test", &result1)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println("   → First request completed / 첫 번째 요청 완료")
	fmt.Printf("   → Response: %s\n", result1["message"])

	// Check cookies / 쿠키 확인
	u, _ := url.Parse(ts.URL)
	cookies := client1.GetCookies(u)
	fmt.Printf("   → Received %d cookies from server:\n", len(cookies))
	for _, cookie := range cookies {
		fmt.Printf("      • %s = %s\n", cookie.Name, cookie.Value)
	}

	// Second request will automatically send cookies / 두 번째 요청은 자동으로 쿠키 전송
	var result2 map[string]string
	err = client1.Get("/test", &result2)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println("   → Second request automatically sent cookies / 두 번째 요청이 자동으로 쿠키 전송")

	fmt.Println("\n2. Persistent Cookie Jar (saved to file) / 지속성 쿠키 저장소 (파일에 저장)")
	fmt.Println("   ─────────────────────────────────────────────────────")

	// Create temporary file for cookies / 쿠키를 위한 임시 파일 생성
	cookieFile := "example_cookies.json"
	defer os.Remove(cookieFile) // Clean up / 정리

	// Create client with persistent cookies / 지속성 쿠키를 사용하는 클라이언트 생성
	client2 := httputil.NewClient(
		httputil.WithBaseURL(ts.URL),
		httputil.WithPersistentCookies(cookieFile), // Enable persistent cookies / 지속성 쿠키 활성화
	)

	// Make request / 요청 수행
	var result3 map[string]string
	err = client2.Get("/test", &result3)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println("   → Request completed / 요청 완료")

	// Save cookies to file / 쿠키를 파일에 저장
	err = client2.SaveCookies()
	if err != nil {
		log.Printf("Error saving cookies: %v", err)
		return
	}

	fmt.Printf("   → Cookies saved to %s\n", cookieFile)

	// Verify file exists / 파일 존재 확인
	if _, err := os.Stat(cookieFile); err == nil {
		fmt.Println("   → Cookie file exists and can be loaded later / 쿠키 파일이 존재하며 나중에 로드 가능")
	}

	fmt.Println("\n3. Cookie Operations / 쿠키 작업")
	fmt.Println("   ─────────────────────────────────────────────────────")

	// Manual cookie operations / 수동 쿠키 작업
	testURL, _ := url.Parse(ts.URL)

	// Set a custom cookie / 사용자 정의 쿠키 설정
	client1.SetCookie(testURL, &http.Cookie{
		Name:  "custom_cookie",
		Value: "test_value",
		Path:  "/",
	})
	fmt.Println("   → Set custom cookie: custom_cookie=test_value")

	// Check if cookie exists / 쿠키 존재 확인
	if client1.HasCookie(testURL, "custom_cookie") {
		fmt.Println("   → Cookie 'custom_cookie' exists / 쿠키 'custom_cookie'가 존재합니다")
	}

	// Get specific cookie / 특정 쿠키 가져오기
	if cookie := client1.GetCookie(testURL, "session_id"); cookie != nil {
		fmt.Printf("   → Retrieved cookie: %s = %s\n", cookie.Name, cookie.Value)
	}

	// Count cookies / 쿠키 개수 세기
	allCookies := client1.GetCookies(testURL)
	fmt.Printf("   → Total cookies: %d\n", len(allCookies))

	// Clear all cookies / 모든 쿠키 제거
	err = client1.ClearCookies()
	if err != nil {
		log.Printf("Error clearing cookies: %v", err)
		return
	}
	fmt.Println("   → All cookies cleared / 모든 쿠키 제거됨")

	// Verify cookies are cleared / 쿠키가 제거되었는지 확인
	afterClear := client1.GetCookies(testURL)
	fmt.Printf("   → Cookies after clear: %d\n", len(afterClear))

	fmt.Println("\n   → Cookie management demonstration complete!")
	fmt.Println("   → 쿠키 관리 데모 완료!")
}

// ═══════════════════════════════════════════════════════════════════════════
// MAIN FUNCTION / 메인 함수
// ═══════════════════════════════════════════════════════════════════════════

func main() {
	// Print banner / 배너 출력
	fmt.Println(strings.Repeat("═", 80))
	fmt.Println("                    httputil Package Examples")
	fmt.Println("                    httputil 패키지 예제")
	fmt.Println(strings.Repeat("═", 80))
	fmt.Printf("Version: %s\n", httputil.Version)
	fmt.Println(strings.Repeat("═", 80))

	// Run all examples / 모든 예제 실행
	exampleBasicHTTPMethods()         // Phase 1
	exampleClientWithBaseURL()         // Phase 1
	exampleContextAndTimeout()         // Phase 1
	exampleErrorHandling()             // Phase 1
	exampleResponseHelpers()           // Phase 2
	exampleFileDownload()              // Phase 3
	exampleFileUpload()                // Phase 3
	exampleURLBuilder()                // Phase 4
	exampleFormBuilder()               // Phase 4
	exampleCookieManagement()          // Phase 5a
	exampleComprehensiveWorkflow()     // All Phases

	// Print summary / 요약 출력
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("                    All Examples Completed!")
	fmt.Println("                    모든 예제 완료!")
	fmt.Println(strings.Repeat("═", 80))
	fmt.Println("\nFeatures demonstrated / 시연된 기능:")
	fmt.Println("  ✓ Phase 1: Basic HTTP methods (GET, POST, PUT, DELETE)")
	fmt.Println("  ✓ Phase 1: Client with base URL and configuration")
	fmt.Println("  ✓ Phase 1: Context and timeout handling")
	fmt.Println("  ✓ Phase 1: Error handling (HTTPError, TimeoutError)")
	fmt.Println("  ✓ Phase 2: Response helpers (status checks, body access)")
	fmt.Println("  ✓ Phase 3: File operations (download/upload with progress)")
	fmt.Println("  ✓ Phase 4: URL Builder (fluent API)")
	fmt.Println("  ✓ Phase 4: Form Builder (fluent API)")
	fmt.Println("  ✓ Phase 5a: Cookie Management (in-memory and persistent)")
	fmt.Println(strings.Repeat("═", 80))
}
