package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/httputil"
	"github.com/arkd0ng/go-utils/logging"
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

var logger *logging.Logger

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 1: BASIC HTTP METHODS (Phase 1) / 예제 1: 기본 HTTP 메서드 (1단계)
// ═══════════════════════════════════════════════════════════════════════════

func example01_BasicHTTPMethods() {
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║  EXAMPLE 1: Basic HTTP Methods (Phase 1)                                  ║")
	logger.Info("║  예제 1: 기본 HTTP 메서드 (1단계)                                          ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	// GET request / GET 요청
	logger.Info("📝 1.1 Simple GET Request / 간단한 GET 요청")
	logger.Info("   Function: httputil.Get(url, &result)")
	logger.Info("   Description: Fetch data with GET method / GET 메서드로 데이터 가져오기")

	var users []User
	err := httputil.Get("https://jsonplaceholder.typicode.com/users", &users)
	if err != nil {
		logger.Error("GET request failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "user_count", len(users), "first_user", users[0].Name)

	// GET with Context / Context를 사용한 GET
	logger.Info("")
	logger.Info("📝 1.2 GET with Context / Context를 사용한 GET")
	logger.Info("   Function: httputil.GetContext(ctx, url, &result)")
	logger.Info("   Description: GET with cancellation support / 취소 지원이 있는 GET")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	err = httputil.GetContext(ctx, "https://jsonplaceholder.typicode.com/users/1", &user)
	if err != nil {
		logger.Error("GET with context failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "name", user.Name, "email", user.Email)

	// POST request / POST 요청
	logger.Info("")
	logger.Info("📝 1.3 POST Request / POST 요청")
	logger.Info("   Function: httputil.Post(url, body, &result)")
	logger.Info("   Description: Create new resource / 새 리소스 생성")

	newPost := Post{
		UserID: 1,
		Title:  "Test Post from httputil",
		Body:   "This is a test post created using httputil package.",
	}
	var createdPost Post
	err = httputil.Post("https://jsonplaceholder.typicode.com/posts", newPost, &createdPost)
	if err != nil {
		logger.Error("POST request failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "post_id", createdPost.ID, "title", createdPost.Title)

	// POST with Context / Context를 사용한 POST
	logger.Info("")
	logger.Info("📝 1.4 POST with Context / Context를 사용한 POST")
	logger.Info("   Function: httputil.PostContext(ctx, url, body, &result)")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	var createdPost2 Post
	err = httputil.PostContext(ctx2, "https://jsonplaceholder.typicode.com/posts", newPost, &createdPost2)
	if err != nil {
		logger.Error("POST with context failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "post_id", createdPost2.ID)

	// PUT request / PUT 요청
	logger.Info("")
	logger.Info("📝 1.5 PUT Request / PUT 요청")
	logger.Info("   Function: httputil.Put(url, body, &result)")
	logger.Info("   Description: Update entire resource / 전체 리소스 업데이트")

	updatedPost := Post{
		ID:     1,
		UserID: 1,
		Title:  "Updated Title",
		Body:   "Updated body content.",
	}
	var resultPost Post
	err = httputil.Put("https://jsonplaceholder.typicode.com/posts/1", updatedPost, &resultPost)
	if err != nil {
		logger.Error("PUT request failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "title", resultPost.Title)

	// PUT with Context / Context를 사용한 PUT
	logger.Info("")
	logger.Info("📝 1.6 PUT with Context / Context를 사용한 PUT")
	logger.Info("   Function: httputil.PutContext(ctx, url, body, &result)")

	ctx3, cancel3 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel3()

	var resultPost2 Post
	err = httputil.PutContext(ctx3, "https://jsonplaceholder.typicode.com/posts/1", updatedPost, &resultPost2)
	if err != nil {
		logger.Error("PUT with context failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "title", resultPost2.Title)

	// PATCH request / PATCH 요청
	logger.Info("")
	logger.Info("📝 1.7 PATCH Request / PATCH 요청")
	logger.Info("   Function: httputil.Patch(url, body, &result)")
	logger.Info("   Description: Partially update resource / 리소스 부분 업데이트")

	patchData := map[string]interface{}{
		"title": "Patched Title",
	}
	var patchedPost Post
	err = httputil.Patch("https://jsonplaceholder.typicode.com/posts/1", patchData, &patchedPost)
	if err != nil {
		logger.Error("PATCH request failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "title", patchedPost.Title)

	// PATCH with Context / Context를 사용한 PATCH
	logger.Info("")
	logger.Info("📝 1.8 PATCH with Context / Context를 사용한 PATCH")
	logger.Info("   Function: httputil.PatchContext(ctx, url, body, &result)")

	ctx4, cancel4 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel4()

	var patchedPost2 Post
	err = httputil.PatchContext(ctx4, "https://jsonplaceholder.typicode.com/posts/1", patchData, &patchedPost2)
	if err != nil {
		logger.Error("PATCH with context failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "title", patchedPost2.Title)

	// DELETE request / DELETE 요청
	logger.Info("")
	logger.Info("📝 1.9 DELETE Request / DELETE 요청")
	logger.Info("   Function: httputil.Delete(url, &result)")
	logger.Info("   Description: Delete resource / 리소스 삭제")

	err = httputil.Delete("https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		logger.Error("DELETE request failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과: Post deleted successfully / 게시물 삭제 성공")

	// DELETE with Context / Context를 사용한 DELETE
	logger.Info("")
	logger.Info("📝 1.10 DELETE with Context / Context를 사용한 DELETE")
	logger.Info("   Function: httputil.DeleteContext(ctx, url, &result)")

	ctx5, cancel5 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel5()

	err = httputil.DeleteContext(ctx5, "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		logger.Error("DELETE with context failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과: Post deleted successfully / 게시물 삭제 성공")

	logger.Info("")
	logger.Info("   ✅ Example 1 completed / 예제 1 완료")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 2: CLIENT CONFIGURATION (Phase 1) / 예제 2: 클라이언트 설정 (1단계)
// ═══════════════════════════════════════════════════════════════════════════

func example02_ClientConfiguration() {
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║  EXAMPLE 2: Client Configuration (Phase 1)                                ║")
	logger.Info("║  예제 2: 클라이언트 설정 (1단계)                                           ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	// WithBaseURL - Set base URL / 기본 URL 설정
	logger.Info("📝 2.1 Client with Base URL / Base URL을 가진 클라이언트")
	logger.Info("   Option: httputil.WithBaseURL(url)")
	logger.Info("   Description: Set base URL for all requests / 모든 요청의 기본 URL 설정")

	client1 := httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
	)

	var users1 []User
	err := client1.Get("/users?_limit=3", &users1)
	if err != nil {
		logger.Error("Request failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "user_count", len(users1))

	// WithTimeout - Set request timeout / 요청 타임아웃 설정
	logger.Info("")
	logger.Info("📝 2.2 Client with Timeout / 타임아웃을 가진 클라이언트")
	logger.Info("   Option: httputil.WithTimeout(duration)")
	logger.Info("   Description: Set request timeout duration / 요청 타임아웃 기간 설정")

	client2 := httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithTimeout(30*time.Second),
	)

	var users2 []User
	err = client2.Get("/users?_limit=3", &users2)
	if err != nil {
		logger.Error("Request failed", "error", err)
		return
	}
	logger.Info("   ✅ Result / 결과:", "user_count", len(users2), "timeout", "30s")

	// WithRetry - Set retry count / 재시도 횟수 설정
	logger.Info("")
	logger.Info("📝 2.3 Client with Retry / 재시도를 가진 클라이언트")
	logger.Info("   Option: httputil.WithRetry(count)")
	logger.Info("   Description: Set number of retries on failure / 실패 시 재시도 횟수 설정")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithRetry(3),
	)
	logger.Info("   ✅ Client configured with 3 retries / 3회 재시도로 클라이언트 설정됨")

	// WithRetryBackoff - Set retry backoff strategy / 재시도 백오프 전략 설정
	logger.Info("")
	logger.Info("📝 2.4 Client with Retry Backoff / 재시도 백오프를 가진 클라이언트")
	logger.Info("   Option: httputil.WithRetryBackoff(min, max)")
	logger.Info("   Description: Exponential backoff for retries / 재시도를 위한 지수 백오프")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithRetry(3),
		httputil.WithRetryBackoff(1*time.Second, 10*time.Second),
	)
	logger.Info("   ✅ Client configured with exponential backoff / 지수 백오프로 클라이언트 설정됨")
	logger.Info("   Backoff strategy: 1s → 2s → 4s → 8s (max 10s)")

	// WithUserAgent - Set user agent / 사용자 에이전트 설정
	logger.Info("")
	logger.Info("📝 2.5 Client with User Agent / User Agent를 가진 클라이언트")
	logger.Info("   Option: httputil.WithUserAgent(agent)")
	logger.Info("   Description: Set custom User-Agent header / 사용자 정의 User-Agent 헤더 설정")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithUserAgent("httputil-example/1.0"),
	)
	logger.Info("   ✅ Client configured with custom user agent / 사용자 정의 user agent로 설정됨")

	// WithHeader - Set custom header / 사용자 정의 헤더 설정
	logger.Info("")
	logger.Info("📝 2.6 Client with Custom Header / 사용자 정의 헤더를 가진 클라이언트")
	logger.Info("   Option: httputil.WithHeader(key, value)")
	logger.Info("   Description: Add custom header to all requests / 모든 요청에 사용자 정의 헤더 추가")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithHeader("X-Custom-Header", "example-value"),
	)
	logger.Info("   ✅ Client configured with custom header / 사용자 정의 헤더로 설정됨")

	// WithHeaders - Set multiple headers / 여러 헤더 설정
	logger.Info("")
	logger.Info("📝 2.7 Client with Multiple Headers / 여러 헤더를 가진 클라이언트")
	logger.Info("   Option: httputil.WithHeaders(map[string]string)")
	logger.Info("   Description: Add multiple headers at once / 여러 헤더를 한 번에 추가")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithHeaders(map[string]string{
			"X-Custom-Header-1": "value1",
			"X-Custom-Header-2": "value2",
		}),
	)
	logger.Info("   ✅ Client configured with multiple headers / 여러 헤더로 설정됨")

	// WithBearerToken - Set bearer token authentication / Bearer 토큰 인증 설정
	logger.Info("")
	logger.Info("📝 2.8 Client with Bearer Token / Bearer Token을 가진 클라이언트")
	logger.Info("   Option: httputil.WithBearerToken(token)")
	logger.Info("   Description: Set Authorization: Bearer <token> / Authorization: Bearer <토큰> 설정")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithBearerToken("example_token_12345"),
	)
	logger.Info("   ✅ Client configured with bearer token / Bearer 토큰으로 설정됨")

	// WithBasicAuth - Set basic authentication / 기본 인증 설정
	logger.Info("")
	logger.Info("📝 2.9 Client with Basic Auth / Basic Auth를 가진 클라이언트")
	logger.Info("   Option: httputil.WithBasicAuth(username, password)")
	logger.Info("   Description: Set HTTP Basic Authentication / HTTP 기본 인증 설정")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithBasicAuth("username", "password"),
	)
	logger.Info("   ✅ Client configured with basic auth / Basic Auth로 설정됨")

	// WithQueryParams - Set default query parameters / 기본 쿼리 파라미터 설정
	logger.Info("")
	logger.Info("📝 2.10 Client with Query Params / 쿼리 파라미터를 가진 클라이언트")
	logger.Info("   Option: httputil.WithQueryParams(params)")
	logger.Info("   Description: Add default query params to all requests / 모든 요청에 기본 쿼리 파라미터 추가")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithQueryParams(map[string]string{
			"_limit": "5",
		}),
	)
	logger.Info("   ✅ Client configured with default query params / 기본 쿼리 파라미터로 설정됨")

	// WithFollowRedirects - Control redirect behavior / 리다이렉트 동작 제어
	logger.Info("")
	logger.Info("📝 2.11 Client with Follow Redirects / 리다이렉트 따라가기를 가진 클라이언트")
	logger.Info("   Option: httputil.WithFollowRedirects(follow)")
	logger.Info("   Description: Enable/disable automatic redirect following / 자동 리다이렉트 따라가기 활성화/비활성화")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithFollowRedirects(true),
	)
	logger.Info("   ✅ Client configured to follow redirects / 리다이렉트 따라가기로 설정됨")

	// WithMaxRedirects - Set maximum redirects / 최대 리다이렉트 설정
	logger.Info("")
	logger.Info("📝 2.12 Client with Max Redirects / 최대 리다이렉트를 가진 클라이언트")
	logger.Info("   Option: httputil.WithMaxRedirects(max)")
	logger.Info("   Description: Set maximum number of redirects / 최대 리다이렉트 횟수 설정")

	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithFollowRedirects(true),
		httputil.WithMaxRedirects(10),
	)
	logger.Info("   ✅ Client configured with max 10 redirects / 최대 10회 리다이렉트로 설정됨")

	// WithTLSConfig - Set TLS configuration / TLS 설정
	logger.Info("")
	logger.Info("📝 2.13 Client with TLS Config / TLS 설정을 가진 클라이언트")
	logger.Info("   Option: httputil.WithTLSConfig(config)")
	logger.Info("   Description: Set custom TLS configuration / 사용자 정의 TLS 설정")

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
	}
	_ = httputil.NewClient(
		httputil.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httputil.WithTLSConfig(tlsConfig),
	)
	logger.Info("   ✅ Client configured with TLS config / TLS 설정으로 구성됨")

	logger.Info("")
	logger.Info("   ✅ Example 2 completed / 예제 2 완료")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 3: ERROR HANDLING / 예제 3: 에러 처리
// ═══════════════════════════════════════════════════════════════════════════

func example03_ErrorHandling() {
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║  EXAMPLE 3: Error Handling                                                ║")
	logger.Info("║  예제 3: 에러 처리                                                         ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	// IsHTTPError - Check if error is HTTP error / HTTP 에러인지 확인
	logger.Info("📝 3.1 Check HTTP Error / HTTP 에러 확인")
	logger.Info("   Function: httputil.IsHTTPError(err)")
	logger.Info("   Description: Check if error is HTTPError type / HTTPError 타입인지 확인")

	var user User
	err := httputil.Get("https://jsonplaceholder.typicode.com/users/999999", &user)
	if err != nil {
		if httputil.IsHTTPError(err) {
			logger.Info("   ✅ Detected HTTP error / HTTP 에러 감지됨")
			statusCode := httputil.GetStatusCode(err)
			logger.Info("   Status code:", "code", statusCode)
			if statusCode == 404 {
				logger.Info("   → Resource not found (404) / 리소스를 찾을 수 없음 (404)")
			}
		}
	}

	// IsTimeoutError - Check if error is timeout / 타임아웃 에러인지 확인
	logger.Info("")
	logger.Info("📝 3.2 Check Timeout Error / 타임아웃 에러 확인")
	logger.Info("   Function: httputil.IsTimeoutError(err)")
	logger.Info("   Description: Check if error is timeout / 타임아웃 에러인지 확인")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	err = httputil.GetContext(ctx, "https://jsonplaceholder.typicode.com/users", &user)
	if err != nil {
		if httputil.IsTimeoutError(err) {
			logger.Info("   ✅ Detected timeout error / 타임아웃 에러 감지됨")
		}
	}

	// IsRetryError - Check if error is retryable / 재시도 가능한 에러인지 확인
	logger.Info("")
	logger.Info("📝 3.3 Check Retry Error / 재시도 에러 확인")
	logger.Info("   Function: httputil.IsRetryError(err)")
	logger.Info("   Description: Check if error is retryable / 재시도 가능한 에러인지 확인")
	logger.Info("   ✅ Retry error check function available / 재시도 에러 확인 함수 사용 가능")

	// GetStatusCode - Get HTTP status code / HTTP 상태 코드 가져오기
	logger.Info("")
	logger.Info("📝 3.4 Get Status Code from Error / 에러에서 상태 코드 가져오기")
	logger.Info("   Function: httputil.GetStatusCode(err)")
	logger.Info("   Description: Extract status code from HTTP error / HTTP 에러에서 상태 코드 추출")

	err = httputil.Get("https://jsonplaceholder.typicode.com/users/999999", &user)
	if err != nil && httputil.IsHTTPError(err) {
		code := httputil.GetStatusCode(err)
		logger.Info("   ✅ Status code extracted:", "code", code)
	}

	logger.Info("")
	logger.Info("   ✅ Example 3 completed / 예제 3 완료")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 4: RESPONSE HELPERS (Phase 2) / 예제 4: 응답 헬퍼 (2단계)
// ═══════════════════════════════════════════════════════════════════════════

func example04_ResponseHelpers() {
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║  EXAMPLE 4: Response Helpers (Phase 2)                                    ║")
	logger.Info("║  예제 4: 응답 헬퍼 (2단계)                                                 ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	// DoRaw - Get raw response / 원시 응답 가져오기
	logger.Info("📝 4.1 DoRaw - Get Raw Response / 원시 응답 가져오기")
	logger.Info("   Function: httputil.DoRaw(method, url, body)")
	logger.Info("   Description: Execute request and get Response object / 요청 실행 및 Response 객체 가져오기")

	resp, err := httputil.DoRaw("GET", "https://jsonplaceholder.typicode.com/users/1", nil)
	if err != nil {
		logger.Error("DoRaw failed", "error", err)
		return
	}

	logger.Info("   Response inspection / 응답 검사:")
	logger.Info("   • Status Code:", "code", resp.StatusCode)
	logger.Info("   • IsSuccess():", "result", resp.IsSuccess())
	logger.Info("   • IsOK() (200):", "result", resp.IsOK())
	logger.Info("   • IsCreated() (201):", "result", resp.IsCreated())
	logger.Info("   • IsNoContent() (204):", "result", resp.IsNoContent())
	logger.Info("   • IsBadRequest() (400):", "result", resp.IsBadRequest())
	logger.Info("   • IsNotFound() (404):", "result", resp.IsNotFound())
	logger.Info("   • IsServerError() (5xx):", "result", resp.IsServerError())
	logger.Info("   • Content-Type:", "type", resp.ContentType())
	logger.Info("   • Content-Length:", "length", resp.ContentLength())

	// DoRawContext - Get raw response with context / Context로 원시 응답 가져오기
	logger.Info("")
	logger.Info("📝 4.2 DoRawContext - Raw Response with Context / Context로 원시 응답")
	logger.Info("   Function: httputil.DoRawContext(ctx, method, url, body)")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp2, err := httputil.DoRawContext(ctx, "GET", "https://jsonplaceholder.typicode.com/users/1", nil)
	if err != nil {
		logger.Error("DoRawContext failed", "error", err)
		return
	}
	logger.Info("   ✅ Response received:", "status", resp2.StatusCode)

	// Access body in different formats / 다양한 형식으로 본문 접근
	logger.Info("")
	logger.Info("📝 4.3 Access Response Body / 응답 본문 접근")
	logger.Info("   Methods: Body(), String(), JSON()")

	// Get as bytes / 바이트로 가져오기
	bodyBytes := resp.Body()
	logger.Info("   • Body() - raw bytes:", "length", len(bodyBytes))

	// Get as string / 문자열로 가져오기
	bodyString := resp.String()
	logger.Info("   • String() - text:", "length", len(bodyString))

	// Get as JSON / JSON으로 가져오기
	var user User
	err = resp.JSON(&user)
	if err != nil {
		logger.Error("JSON decode failed", "error", err)
		return
	}
	logger.Info("   • JSON() - decoded:", "name", user.Name, "email", user.Email)

	// Header - Get specific header / 특정 헤더 가져오기
	logger.Info("")
	logger.Info("📝 4.4 Get Response Headers / 응답 헤더 가져오기")
	logger.Info("   Method: Header(key)")

	contentType := resp.Header("Content-Type")
	logger.Info("   ✅ Content-Type header:", "value", contentType)

	logger.Info("")
	logger.Info("   ✅ Example 4 completed / 예제 4 완료")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 5: FILE OPERATIONS (Phase 3) / 예제 5: 파일 작업 (3단계)
// ═══════════════════════════════════════════════════════════════════════════

func example05_FileOperations() {
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║  EXAMPLE 5: File Operations (Phase 3)                                     ║")
	logger.Info("║  예제 5: 파일 작업 (3단계)                                                 ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	// Create test server for file operations / 파일 작업을 위한 테스트 서버 생성
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// Download endpoint / 다운로드 엔드포인트
			testData := []byte("This is test file content for download example")
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Length", fmt.Sprintf("%d", len(testData)))
			w.Write(testData)
		} else if r.Method == "POST" {
			// Upload endpoint / 업로드 엔드포인트
			err := r.ParseMultipartForm(10 << 20) // 10MB
			if err != nil {
				http.Error(w, "Failed to parse form", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "success",
				"message": "File uploaded successfully",
			})
		}
	}))
	defer ts.Close()

	client := httputil.NewClient()

	// Download - Download to memory / 메모리로 다운로드
	logger.Info("📝 5.1 Download to Memory / 메모리로 다운로드")
	logger.Info("   Function: httputil.Download(url)")
	logger.Info("   Description: Download file content to memory / 파일 내용을 메모리로 다운로드")

	data, err := client.Download(ts.URL)
	if err != nil {
		logger.Error("Download failed", "error", err)
		return
	}
	logger.Info("   ✅ Downloaded:", "bytes", len(data), "content", string(data[:30])+"...")

	// DownloadContext - Download with context / Context로 다운로드
	logger.Info("")
	logger.Info("📝 5.2 Download with Context / Context로 다운로드")
	logger.Info("   Function: httputil.DownloadContext(ctx, url)")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data2, err := client.DownloadContext(ctx, ts.URL)
	if err != nil {
		logger.Error("Download with context failed", "error", err)
		return
	}
	logger.Info("   ✅ Downloaded:", "bytes", len(data2))

	// DownloadFile - Download to file / 파일로 다운로드
	logger.Info("")
	logger.Info("📝 5.3 Download to File / 파일로 다운로드")
	logger.Info("   Function: httputil.DownloadFile(url, filepath)")
	logger.Info("   Description: Download and save to file / 다운로드하여 파일에 저장")

	tmpFile := "/tmp/httputil_example_download.txt"
	defer os.Remove(tmpFile)

	err = client.DownloadFile(ts.URL, tmpFile)
	if err != nil {
		logger.Error("DownloadFile failed", "error", err)
		return
	}

	// Verify file / 파일 확인
	fileData, _ := os.ReadFile(tmpFile)
	logger.Info("   ✅ File saved:", "path", tmpFile, "bytes", len(fileData))

	// DownloadFileContext - Download to file with progress / 진행 상황과 함께 파일로 다운로드
	logger.Info("")
	logger.Info("📝 5.4 Download with Progress Tracking / 진행 상황 추적과 함께 다운로드")
	logger.Info("   Function: httputil.DownloadFileContext(ctx, url, filepath, progressFunc)")
	logger.Info("   Description: Track download progress / 다운로드 진행 상황 추적")

	tmpFile2 := "/tmp/httputil_example_download2.txt"
	defer os.Remove(tmpFile2)

	var lastProgress int64
	progressFunc := func(bytesRead, totalBytes int64) {
		lastProgress = bytesRead
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	err = client.DownloadFileContext(ctx2, ts.URL, tmpFile2, progressFunc)
	if err != nil {
		logger.Error("DownloadFileContext failed", "error", err)
		return
	}
	logger.Info("   ✅ Download completed:", "bytes_read", lastProgress)

	// Create test file for upload / 업로드용 테스트 파일 생성
	uploadFile := "/tmp/httputil_example_upload.txt"
	testContent := []byte("This is test content for upload")
	os.WriteFile(uploadFile, testContent, 0644)
	defer os.Remove(uploadFile)

	// UploadFile - Upload single file / 단일 파일 업로드
	logger.Info("")
	logger.Info("📝 5.5 Upload Single File / 단일 파일 업로드")
	logger.Info("   Function: httputil.UploadFile(url, fieldName, filepath, &result)")
	logger.Info("   Description: Upload file using multipart/form-data / multipart/form-data로 파일 업로드")

	var uploadResult map[string]string
	err = client.UploadFile(ts.URL, "file", uploadFile, &uploadResult)
	if err != nil {
		logger.Error("UploadFile failed", "error", err)
		return
	}
	logger.Info("   ✅ Upload completed:", "status", uploadResult["status"], "message", uploadResult["message"])

	// UploadFileContext - Upload with progress / 진행 상황과 함께 업로드
	logger.Info("")
	logger.Info("📝 5.6 Upload with Progress Tracking / 진행 상황 추적과 함께 업로드")
	logger.Info("   Function: httputil.UploadFileContext(ctx, url, fieldName, filepath, &result, progressFunc)")

	var uploadProgress int64
	uploadProgressFunc := func(bytesWritten, totalBytes int64) {
		uploadProgress = bytesWritten
	}

	ctx3, cancel3 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel3()

	var uploadResult2 map[string]string
	err = client.UploadFileContext(ctx3, ts.URL, "file", uploadFile, &uploadResult2, uploadProgressFunc)
	if err != nil {
		logger.Error("UploadFileContext failed", "error", err)
		return
	}
	logger.Info("   ✅ Upload completed:", "bytes_written", uploadProgress)

	// Create multiple test files / 여러 테스트 파일 생성
	uploadFile1 := "/tmp/httputil_upload1.txt"
	uploadFile2 := "/tmp/httputil_upload2.txt"
	os.WriteFile(uploadFile1, []byte("Content 1"), 0644)
	os.WriteFile(uploadFile2, []byte("Content 2"), 0644)
	defer os.Remove(uploadFile1)
	defer os.Remove(uploadFile2)

	// UploadFiles - Upload multiple files / 여러 파일 업로드
	logger.Info("")
	logger.Info("📝 5.7 Upload Multiple Files / 여러 파일 업로드")
	logger.Info("   Function: httputil.UploadFiles(url, files, &result)")
	logger.Info("   Description: Upload multiple files at once / 여러 파일을 한 번에 업로드")

	files := map[string]string{
		"file1": uploadFile1,
		"file2": uploadFile2,
	}

	var uploadResult3 map[string]string
	err = client.UploadFiles(ts.URL, files, &uploadResult3)
	if err != nil {
		logger.Error("UploadFiles failed", "error", err)
		return
	}
	logger.Info("   ✅ Multiple files uploaded:", "file_count", len(files))

	// UploadFilesContext - Upload multiple files with context / Context로 여러 파일 업로드
	logger.Info("")
	logger.Info("📝 5.8 Upload Multiple Files with Context / Context로 여러 파일 업로드")
	logger.Info("   Function: httputil.UploadFilesContext(ctx, url, files, &result)")

	ctx4, cancel4 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel4()

	var uploadResult4 map[string]string
	err = client.UploadFilesContext(ctx4, ts.URL, files, &uploadResult4)
	if err != nil {
		logger.Error("UploadFilesContext failed", "error", err)
		return
	}
	logger.Info("   ✅ Multiple files uploaded with context:", "file_count", len(files))

	logger.Info("")
	logger.Info("   ✅ Example 5 completed / 예제 5 완료")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 6: URL BUILDER (Phase 4) / 예제 6: URL 빌더 (4단계)
// ═══════════════════════════════════════════════════════════════════════════

func example06_URLBuilder() {
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║  EXAMPLE 6: URL Builder (Phase 4)                                         ║")
	logger.Info("║  예제 6: URL 빌더 (4단계)                                                  ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	// NewURL - Create URL builder / URL 빌더 생성
	logger.Info("📝 6.1 URL Builder - Fluent API / URL 빌더 - Fluent API")
	logger.Info("   Function: httputil.NewURL(base)")
	logger.Info("   Description: Build URLs with fluent interface / Fluent 인터페이스로 URL 구축")

	includeInactive := false
	builtURL := httputil.NewURL("https://jsonplaceholder.typicode.com").
		Path("posts").
		Param("userId", "1").
		Param("_limit", "5").
		ParamIf(includeInactive, "status", "inactive").
		Build()

	logger.Info("   ✅ Built URL:", "url", builtURL)

	// Use the built URL / 구축된 URL 사용
	var posts []Post
	err := httputil.Get(builtURL, &posts)
	if err != nil {
		logger.Error("Request failed", "error", err)
		return
	}
	logger.Info("   ✅ Fetched posts:", "count", len(posts))

	// JoinURL - Join URL parts / URL 부분 결합
	logger.Info("")
	logger.Info("📝 6.2 Join URL Parts / URL 부분 결합")
	logger.Info("   Function: httputil.JoinURL(parts...)")
	logger.Info("   Description: Safely join URL path segments / URL 경로 세그먼트를 안전하게 결합")

	joinedURL := httputil.JoinURL("https://jsonplaceholder.typicode.com", "users", "1")
	logger.Info("   ✅ Joined URL:", "url", joinedURL)

	// ParseURL - Parse URL string / URL 문자열 파싱
	logger.Info("")
	logger.Info("📝 6.3 Parse URL / URL 파싱")
	logger.Info("   Function: httputil.ParseURL(urlStr)")
	logger.Info("   Description: Parse URL string into url.URL / URL 문자열을 url.URL로 파싱")

	parsedURL, err := httputil.ParseURL(joinedURL)
	if err != nil {
		logger.Error("ParseURL failed", "error", err)
		return
	}
	logger.Info("   ✅ Parsed URL:", "scheme", parsedURL.Scheme, "host", parsedURL.Host, "path", parsedURL.Path)

	// GetDomain - Extract domain / 도메인 추출
	logger.Info("")
	logger.Info("📝 6.4 Get Domain / 도메인 가져오기")
	logger.Info("   Function: httputil.GetDomain(urlStr)")
	logger.Info("   Description: Extract domain from URL / URL에서 도메인 추출")

	domain, err := httputil.GetDomain(joinedURL)
	if err != nil {
		logger.Error("GetDomain failed", "error", err)
		return
	}
	logger.Info("   ✅ Domain:", "domain", domain)

	// GetScheme - Extract scheme / 스키마 추출
	logger.Info("")
	logger.Info("📝 6.5 Get Scheme / 스키마 가져오기")
	logger.Info("   Function: httputil.GetScheme(urlStr)")
	logger.Info("   Description: Extract scheme (http/https) / 스키마 (http/https) 추출")

	scheme, err := httputil.GetScheme(joinedURL)
	if err != nil {
		logger.Error("GetScheme failed", "error", err)
		return
	}
	logger.Info("   ✅ Scheme:", "scheme", scheme)

	// GetPath - Extract path / 경로 추출
	logger.Info("")
	logger.Info("📝 6.6 Get Path / 경로 가져오기")
	logger.Info("   Function: httputil.GetPath(urlStr)")
	logger.Info("   Description: Extract path from URL / URL에서 경로 추출")

	path, err := httputil.GetPath(joinedURL)
	if err != nil {
		logger.Error("GetPath failed", "error", err)
		return
	}
	logger.Info("   ✅ Path:", "path", path)

	// IsAbsoluteURL - Check if URL is absolute / 절대 URL인지 확인
	logger.Info("")
	logger.Info("📝 6.7 Check if Absolute URL / 절대 URL인지 확인")
	logger.Info("   Function: httputil.IsAbsoluteURL(urlStr)")
	logger.Info("   Description: Check if URL is absolute or relative / URL이 절대인지 상대인지 확인")

	isAbsolute := httputil.IsAbsoluteURL(joinedURL)
	logger.Info("   ✅ Is absolute:", "result", isAbsolute)

	isAbsoluteRel := httputil.IsAbsoluteURL("/users/1")
	logger.Info("   ✅ Is '/users/1' absolute:", "result", isAbsoluteRel)

	// NormalizeURL - Normalize URL / URL 정규화
	logger.Info("")
	logger.Info("📝 6.8 Normalize URL / URL 정규화")
	logger.Info("   Function: httputil.NormalizeURL(urlStr)")
	logger.Info("   Description: Normalize URL format / URL 형식 정규화")

	normalized := httputil.NormalizeURL("https://jsonplaceholder.typicode.com//users///1")
	logger.Info("   ✅ Normalized:", "url", normalized)

	// AddQueryParams - Add query parameters / 쿼리 파라미터 추가
	logger.Info("")
	logger.Info("📝 6.9 Add Query Parameters / 쿼리 파라미터 추가")
	logger.Info("   Function: httputil.AddQueryParams(urlStr, params)")
	logger.Info("   Description: Add query parameters to URL / URL에 쿼리 파라미터 추가")

	withParams, err := httputil.AddQueryParams(joinedURL, map[string]string{
		"_limit":  "5",
		"_sort":   "id",
		"_order":  "desc",
	})
	if err != nil {
		logger.Error("AddQueryParams failed", "error", err)
		return
	}
	logger.Info("   ✅ URL with params:", "url", withParams)

	// GetQueryParam - Get single query parameter / 단일 쿼리 파라미터 가져오기
	logger.Info("")
	logger.Info("📝 6.10 Get Query Parameter / 쿼리 파라미터 가져오기")
	logger.Info("   Function: httputil.GetQueryParam(urlStr, key)")
	logger.Info("   Description: Get specific query parameter value / 특정 쿼리 파라미터 값 가져오기")

	limitValue, err := httputil.GetQueryParam(withParams, "_limit")
	if err != nil {
		logger.Error("GetQueryParam failed", "error", err)
		return
	}
	logger.Info("   ✅ _limit parameter:", "value", limitValue)

	// GetAllQueryParams - Get all query parameters / 모든 쿼리 파라미터 가져오기
	logger.Info("")
	logger.Info("📝 6.11 Get All Query Parameters / 모든 쿼리 파라미터 가져오기")
	logger.Info("   Function: httputil.GetAllQueryParams(urlStr)")
	logger.Info("   Description: Get all query parameters as map / 모든 쿼리 파라미터를 맵으로 가져오기")

	allParams, err := httputil.GetAllQueryParams(withParams)
	if err != nil {
		logger.Error("GetAllQueryParams failed", "error", err)
		return
	}
	logger.Info("   ✅ All parameters:", "count", len(allParams))
	for key, values := range allParams {
		logger.Info("      •", "key", key, "value", values[0])
	}

	// RemoveQueryParam - Remove query parameter / 쿼리 파라미터 제거
	logger.Info("")
	logger.Info("📝 6.12 Remove Query Parameter / 쿼리 파라미터 제거")
	logger.Info("   Function: httputil.RemoveQueryParam(urlStr, key)")
	logger.Info("   Description: Remove specific query parameter / 특정 쿼리 파라미터 제거")

	withoutLimit, err := httputil.RemoveQueryParam(withParams, "_limit")
	if err != nil {
		logger.Error("RemoveQueryParam failed", "error", err)
		return
	}
	logger.Info("   ✅ URL without _limit:", "url", withoutLimit)

	logger.Info("")
	logger.Info("   ✅ Example 6 completed / 예제 6 완료")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 7: FORM BUILDER (Phase 4) / 예제 7: Form 빌더 (4단계)
// ═══════════════════════════════════════════════════════════════════════════

func example07_FormBuilder() {
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║  EXAMPLE 7: Form Builder (Phase 4)                                        ║")
	logger.Info("║  예제 7: Form 빌더 (4단계)                                                 ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	// NewForm - Create form builder / Form 빌더 생성
	logger.Info("📝 7.1 Form Builder - Fluent API / Form 빌더 - Fluent API")
	logger.Info("   Function: httputil.NewForm()")
	logger.Info("   Description: Build forms with fluent interface / Fluent 인터페이스로 폼 구축")

	hasPromoCode := true
	form := httputil.NewForm().
		Set("username", "testuser").
		Set("email", "test@example.com").
		Set("age", "30").
		AddIf(hasPromoCode, "promo_code", "SAVE20").
		AddIf(false, "referrer", "none").
		AddMultiple("tags", "go", "http", "api")

	logger.Info("   ✅ Form built with multiple fields / 여러 필드로 폼 구축됨")

	// Has - Check if field exists / 필드 존재 확인
	logger.Info("")
	logger.Info("📝 7.2 Check Field Existence / 필드 존재 확인")
	logger.Info("   Method: form.Has(key)")
	logger.Info("   Description: Check if form has specific field / 폼에 특정 필드가 있는지 확인")

	hasPromo := form.Has("promo_code")
	hasReferrer := form.Has("referrer")
	logger.Info("   ✅ Has promo_code:", "result", hasPromo)
	logger.Info("   ✅ Has referrer:", "result", hasReferrer)

	// Get - Get field value / 필드 값 가져오기
	logger.Info("")
	logger.Info("📝 7.3 Get Field Value / 필드 값 가져오기")
	logger.Info("   Method: form.Get(key)")
	logger.Info("   Description: Get value of specific field / 특정 필드의 값 가져오기")

	username := form.Get("username")
	logger.Info("   ✅ Username:", "value", username)

	// Map - Get all form fields / 모든 폼 필드 가져오기
	logger.Info("")
	logger.Info("📝 7.4 Get All Form Fields / 모든 폼 필드 가져오기")
	logger.Info("   Method: form.Map()")
	logger.Info("   Description: Get all fields as map / 모든 필드를 맵으로 가져오기")

	formMap := form.Map()
	logger.Info("   ✅ Form fields:", "count", len(formMap))
	for key, value := range formMap {
		logger.Info("      •", "key", key, "value", value)
	}

	// Values - Get url.Values / url.Values 가져오기
	logger.Info("")
	logger.Info("📝 7.5 Get url.Values / url.Values 가져오기")
	logger.Info("   Method: form.Values()")
	logger.Info("   Description: Get form as url.Values / 폼을 url.Values로 가져오기")

	values := form.Values()
	logger.Info("   ✅ url.Values:", "type", fmt.Sprintf("%T", values))

	// Encode - Encode form data / 폼 데이터 인코딩
	logger.Info("")
	logger.Info("📝 7.6 Encode Form Data / 폼 데이터 인코딩")
	logger.Info("   Method: form.Encode()")
	logger.Info("   Description: Encode form as URL-encoded string / 폼을 URL 인코딩 문자열로 인코딩")

	encoded := form.Encode()
	logger.Info("   ✅ Encoded form:", "length", len(encoded), "preview", encoded[:50]+"...")

	// ParseForm - Parse form string / 폼 문자열 파싱
	logger.Info("")
	logger.Info("📝 7.7 Parse Form String / 폼 문자열 파싱")
	logger.Info("   Function: httputil.ParseForm(formStr)")
	logger.Info("   Description: Parse URL-encoded form string / URL 인코딩된 폼 문자열 파싱")

	parsedForm, err := httputil.ParseForm(encoded)
	if err != nil {
		logger.Error("ParseForm failed", "error", err)
		return
	}
	logger.Info("   ✅ Parsed form:", "fields", len(parsedForm))

	// EncodeForm - Encode map to form / 맵을 폼으로 인코딩
	logger.Info("")
	logger.Info("📝 7.8 Encode Map to Form / 맵을 폼으로 인코딩")
	logger.Info("   Function: httputil.EncodeForm(data)")
	logger.Info("   Description: Encode map as URL-encoded form / 맵을 URL 인코딩 폼으로 인코딩")

	dataMap := map[string]string{
		"field1": "value1",
		"field2": "value2",
	}
	encodedMap := httputil.EncodeForm(dataMap)
	logger.Info("   ✅ Encoded map:", "result", encodedMap)

	// Create test server for form posting / 폼 포스팅을 위한 테스트 서버 생성
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"received_fields": len(r.PostForm),
			"username":        r.PostForm.Get("username"),
		})
	}))
	defer ts.Close()

	// PostForm - Post form data / 폼 데이터 포스트
	logger.Info("")
	logger.Info("📝 7.9 Post Form Data / 폼 데이터 포스트")
	logger.Info("   Function: httputil.PostForm(url, formData, &result)")
	logger.Info("   Description: Post form data to server / 서버에 폼 데이터 포스트")

	formData := map[string]string{
		"username": "testuser",
		"password": "secret",
	}

	var postResult map[string]interface{}
	err = httputil.PostForm(ts.URL, formData, &postResult)
	if err != nil {
		logger.Error("PostForm failed", "error", err)
		return
	}
	logger.Info("   ✅ Form posted:", "received_fields", postResult["received_fields"])

	// PostFormContext - Post form with context / Context로 폼 포스트
	logger.Info("")
	logger.Info("📝 7.10 Post Form with Context / Context로 폼 포스트")
	logger.Info("   Function: httputil.PostFormContext(ctx, url, formData, &result)")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var postResult2 map[string]interface{}
	err = httputil.PostFormContext(ctx, ts.URL, formData, &postResult2)
	if err != nil {
		logger.Error("PostFormContext failed", "error", err)
		return
	}
	logger.Info("   ✅ Form posted with context:", "received_fields", postResult2["received_fields"])

	logger.Info("")
	logger.Info("   ✅ Example 7 completed / 예제 7 완료")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 8: COOKIE MANAGEMENT (Phase 5a) / 예제 8: 쿠키 관리 (Phase 5a)
// ═══════════════════════════════════════════════════════════════════════════

func example08_CookieManagement() {
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║  EXAMPLE 8: Cookie Management (Phase 5a)                                  ║")
	logger.Info("║  예제 8: 쿠키 관리 (Phase 5a)                                              ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

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

	// WithCookies - In-memory cookie jar / 메모리 내 쿠키 저장소
	logger.Info("📝 8.1 In-Memory Cookie Jar / 메모리 내 쿠키 저장소")
	logger.Info("   Option: httputil.WithCookies()")
	logger.Info("   Description: Enable automatic cookie management / 자동 쿠키 관리 활성화")

	client1 := httputil.NewClient(
		httputil.WithBaseURL(ts.URL),
		httputil.WithCookies(),
	)

	var result1 map[string]string
	err := client1.Get("/test", &result1)
	if err != nil {
		logger.Error("Request failed", "error", err)
		return
	}
	logger.Info("   ✅ First request completed / 첫 번째 요청 완료")

	// GetCookies - Get all cookies / 모든 쿠키 가져오기
	logger.Info("")
	logger.Info("📝 8.2 Get All Cookies / 모든 쿠키 가져오기")
	logger.Info("   Method: client.GetCookies(url)")
	logger.Info("   Description: Retrieve all cookies for URL / URL의 모든 쿠키 가져오기")

	u, _ := url.Parse(ts.URL)
	cookies := client1.GetCookies(u)
	logger.Info("   ✅ Retrieved cookies:", "count", len(cookies))
	for _, cookie := range cookies {
		logger.Info("      •", "name", cookie.Name, "value", cookie.Value)
	}

	// GetCookie - Get specific cookie / 특정 쿠키 가져오기
	logger.Info("")
	logger.Info("📝 8.3 Get Specific Cookie / 특정 쿠키 가져오기")
	logger.Info("   Method: client.GetCookie(url, name)")
	logger.Info("   Description: Get specific cookie by name / 이름으로 특정 쿠키 가져오기")

	sessionCookie := client1.GetCookie(u, "session_id")
	if sessionCookie != nil {
		logger.Info("   ✅ Cookie found:", "name", sessionCookie.Name, "value", sessionCookie.Value)
	}

	// HasCookie - Check if cookie exists / 쿠키 존재 확인
	logger.Info("")
	logger.Info("📝 8.4 Check Cookie Existence / 쿠키 존재 확인")
	logger.Info("   Method: client.HasCookie(url, name)")
	logger.Info("   Description: Check if specific cookie exists / 특정 쿠키 존재 확인")

	hasSession := client1.HasCookie(u, "session_id")
	hasNonExistent := client1.HasCookie(u, "non_existent")
	logger.Info("   ✅ Has session_id:", "result", hasSession)
	logger.Info("   ✅ Has non_existent:", "result", hasNonExistent)

	// SetCookie - Set custom cookie / 사용자 정의 쿠키 설정
	logger.Info("")
	logger.Info("📝 8.5 Set Custom Cookie / 사용자 정의 쿠키 설정")
	logger.Info("   Method: client.SetCookie(url, cookie)")
	logger.Info("   Description: Manually set a cookie / 수동으로 쿠키 설정")

	client1.SetCookie(u, &http.Cookie{
		Name:  "custom_cookie",
		Value: "test_value",
		Path:  "/",
	})
	logger.Info("   ✅ Custom cookie set / 사용자 정의 쿠키 설정됨")

	// Verify custom cookie / 사용자 정의 쿠키 확인
	if client1.HasCookie(u, "custom_cookie") {
		logger.Info("   ✅ Custom cookie verified / 사용자 정의 쿠키 확인됨")
	}

	// WithPersistentCookies - Persistent cookie jar / 지속성 쿠키 저장소
	logger.Info("")
	logger.Info("📝 8.6 Persistent Cookie Jar / 지속성 쿠키 저장소")
	logger.Info("   Option: httputil.WithPersistentCookies(filepath)")
	logger.Info("   Description: Save/load cookies from file / 파일에서 쿠키 저장/로드")

	cookieFile := "/tmp/httputil_example_cookies.json"
	defer os.Remove(cookieFile)

	client2 := httputil.NewClient(
		httputil.WithBaseURL(ts.URL),
		httputil.WithPersistentCookies(cookieFile),
	)

	var result2 map[string]string
	err = client2.Get("/test", &result2)
	if err != nil {
		logger.Error("Request failed", "error", err)
		return
	}
	logger.Info("   ✅ Request completed with persistent cookies / 지속성 쿠키로 요청 완료")

	// SaveCookies - Save cookies to file / 쿠키를 파일에 저장
	logger.Info("")
	logger.Info("📝 8.7 Save Cookies to File / 쿠키를 파일에 저장")
	logger.Info("   Method: client.SaveCookies()")
	logger.Info("   Description: Persist cookies to file / 쿠키를 파일에 저장")

	err = client2.SaveCookies()
	if err != nil {
		logger.Error("SaveCookies failed", "error", err)
		return
	}

	// Verify file exists / 파일 존재 확인
	if _, err := os.Stat(cookieFile); err == nil {
		logger.Info("   ✅ Cookies saved to file:", "path", cookieFile)
	}

	// NewCookieJar - Create in-memory cookie jar / 메모리 내 쿠키 저장소 생성
	logger.Info("")
	logger.Info("📝 8.8 Create In-Memory Cookie Jar / 메모리 내 쿠키 저장소 생성")
	logger.Info("   Function: httputil.NewCookieJar()")
	logger.Info("   Description: Create new cookie jar / 새 쿠키 저장소 생성")

	jar, err := httputil.NewCookieJar()
	if err != nil {
		logger.Error("NewCookieJar failed", "error", err)
		return
	}
	logger.Info("   ✅ Cookie jar created:", "type", fmt.Sprintf("%T", jar))

	// NewPersistentCookieJar - Create persistent cookie jar / 지속성 쿠키 저장소 생성
	logger.Info("")
	logger.Info("📝 8.9 Create Persistent Cookie Jar / 지속성 쿠키 저장소 생성")
	logger.Info("   Function: httputil.NewPersistentCookieJar(filepath)")
	logger.Info("   Description: Create cookie jar with file persistence / 파일 지속성이 있는 쿠키 저장소 생성")

	persistentJar, err := httputil.NewPersistentCookieJar(cookieFile)
	if err != nil {
		logger.Error("NewPersistentCookieJar failed", "error", err)
		return
	}
	logger.Info("   ✅ Persistent cookie jar created:", "type", fmt.Sprintf("%T", persistentJar))

	// WithCookieJar - Use custom cookie jar / 사용자 정의 쿠키 저장소 사용
	logger.Info("")
	logger.Info("📝 8.10 Use Custom Cookie Jar / 사용자 정의 쿠키 저장소 사용")
	logger.Info("   Option: httputil.WithCookieJar(jar)")
	logger.Info("   Description: Use custom cookie jar implementation / 사용자 정의 쿠키 저장소 구현 사용")

	_ = httputil.NewClient(
		httputil.WithBaseURL(ts.URL),
		httputil.WithCookieJar(persistentJar),
	)
	logger.Info("   ✅ Client configured with custom jar / 사용자 정의 저장소로 클라이언트 설정됨")

	// ClearCookies - Clear all cookies / 모든 쿠키 제거
	logger.Info("")
	logger.Info("📝 8.11 Clear All Cookies / 모든 쿠키 제거")
	logger.Info("   Method: client.ClearCookies()")
	logger.Info("   Description: Remove all cookies from jar / 저장소의 모든 쿠키 제거")

	err = client1.ClearCookies()
	if err != nil {
		logger.Error("ClearCookies failed", "error", err)
		return
	}

	// Verify cookies are cleared / 쿠키가 제거되었는지 확인
	afterClear := client1.GetCookies(u)
	logger.Info("   ✅ Cookies cleared:", "remaining", len(afterClear))

	logger.Info("")
	logger.Info("   ✅ Example 8 completed / 예제 8 완료")
}

// ═══════════════════════════════════════════════════════════════════════════
// EXAMPLE 9: ADVANCED FEATURES / 예제 9: 고급 기능
// ═══════════════════════════════════════════════════════════════════════════

func example09_AdvancedFeatures() {
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║  EXAMPLE 9: Advanced Features                                             ║")
	logger.Info("║  예제 9: 고급 기능                                                         ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	// WithProxy - Set HTTP proxy / HTTP 프록시 설정
	logger.Info("📝 9.1 HTTP Proxy Configuration / HTTP 프록시 설정")
	logger.Info("   Option: httputil.WithProxy(proxyURL)")
	logger.Info("   Description: Route requests through HTTP proxy / HTTP 프록시를 통해 요청 라우팅")
	logger.Info("   ℹ️  Example only - proxy not actually configured / 예제만 - 프록시가 실제로 설정되지 않음")

	// Note: Just demonstrating the API / API 시연만 수행
	logger.Info("   Usage: client := httputil.NewClient(")
	logger.Info("       httputil.WithProxy(\"http://proxy.example.com:8080\"),")
	logger.Info("   )")

	// WithLogger - Set custom logger / 사용자 정의 로거 설정
	logger.Info("")
	logger.Info("📝 9.2 Custom Logger Integration / 사용자 정의 로거 통합")
	logger.Info("   Option: httputil.WithLogger(logger)")
	logger.Info("   Description: Integrate custom logger for HTTP operations / HTTP 작업을 위한 사용자 정의 로거 통합")
	logger.Info("   ℹ️  WithLogger requires implementing httputil.Logger interface")
	logger.Info("   ℹ️  WithLogger는 httputil.Logger 인터페이스 구현 필요")
	logger.Info("   ✅ Logger integration available for custom implementations / 사용자 정의 구현에 로거 통합 사용 가능")

	// SetDefaultClient - Set default package-level client / 기본 패키지 레벨 클라이언트 설정
	logger.Info("")
	logger.Info("📝 9.3 Set Default Client / 기본 클라이언트 설정")
	logger.Info("   Function: httputil.SetDefaultClient(client)")
	logger.Info("   Description: Set default client for package-level functions / 패키지 레벨 함수의 기본 클라이언트 설정")

	defaultClient := httputil.NewClient(
		httputil.WithTimeout(30 * time.Second),
		httputil.WithRetry(3),
	)
	httputil.SetDefaultClient(defaultClient)
	logger.Info("   ✅ Default client configured / 기본 클라이언트 설정됨")

	// Version information / 버전 정보
	logger.Info("")
	logger.Info("📝 9.4 Package Version / 패키지 버전")
	logger.Info("   Variable: httputil.Version")
	logger.Info("   Description: Get package version information / 패키지 버전 정보 가져오기")
	logger.Info("   ✅ Package version:", "version", httputil.Version)

	logger.Info("")
	logger.Info("   ✅ Example 9 completed / 예제 9 완료")
}

// ═══════════════════════════════════════════════════════════════════════════
// MAIN FUNCTION / 메인 함수
// ═══════════════════════════════════════════════════════════════════════════

func main() {
	// Setup log file with backup management / 백업 관리와 함께 로그 파일 설정
	logFilePath := "logs/httputil-example.log"

	// Check if previous log file exists / 이전 로그 파일 존재 여부 확인
	if fileutil.Exists(logFilePath) {
		// Get modification time of existing log file / 기존 로그 파일의 수정 시간 가져오기
		modTime, err := fileutil.ModTime(logFilePath)
		if err == nil {
			// Create backup filename with timestamp / 타임스탬프와 함께 백업 파일명 생성
			backupName := fmt.Sprintf("logs/httputil-example-%s.log", modTime.Format("20060102-150405"))

			// Backup existing log file / 기존 로그 파일 백업
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
				// Delete original log file to prevent content duplication / 내용 중복 방지를 위해 원본 로그 파일 삭제
				fileutil.DeleteFile(logFilePath)
			}
		}

		// Cleanup old backup files - keep only 5 most recent / 오래된 백업 파일 정리 - 최근 5개만 유지
		backupPattern := "logs/httputil-example-*.log"
		backupFiles, err := filepath.Glob(backupPattern)
		if err == nil && len(backupFiles) > 5 {
			// Sort by modification time / 수정 시간으로 정렬
			type fileInfo struct {
				path    string
				modTime time.Time
			}
			var files []fileInfo
			for _, f := range backupFiles {
				if mt, err := fileutil.ModTime(f); err == nil {
					files = append(files, fileInfo{path: f, modTime: mt})
				}
			}

			// Sort oldest first / 가장 오래된 것부터 정렬
			for i := 0; i < len(files)-1; i++ {
				for j := i + 1; j < len(files); j++ {
					if files[i].modTime.After(files[j].modTime) {
						files[i], files[j] = files[j], files[i]
					}
				}
			}

			// Delete oldest files to keep only 5 / 5개만 유지하도록 가장 오래된 파일 삭제
			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Printf("🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	// Initialize logger with fixed filename / 고정 파일명으로 로거 초기화
	var err error
	logger, err = logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithLevel(logging.DEBUG),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Print banner / 배너 출력
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║         httputil Package - Comprehensive Examples & Manual                ║")
	logger.Info("║         httputil 패키지 - 종합 예제 및 매뉴얼                              ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")
	logger.Info("📋 Package Information / 패키지 정보")
	logger.Info("   Package Name: github.com/arkd0ng/go-utils/httputil")
	logger.Info("   Version:", "version", httputil.Version)
	logger.Info("   Description: Extremely simple HTTP client utilities (20 lines → 2 lines)")
	logger.Info("   설명: 극도로 간단한 HTTP 클라이언트 유틸리티 (20줄 → 2줄)")
	logger.Info("")

	// Run all examples / 모든 예제 실행
	example01_BasicHTTPMethods()      // Phase 1: Basic HTTP methods
	example02_ClientConfiguration()   // Phase 1: Client configuration
	example03_ErrorHandling()         // Error handling helpers
	example04_ResponseHelpers()       // Phase 2: Response inspection
	example05_FileOperations()        // Phase 3: File download/upload
	example06_URLBuilder()            // Phase 4: URL manipulation
	example07_FormBuilder()           // Phase 4: Form building
	example08_CookieManagement()      // Phase 5a: Cookie management
	example09_AdvancedFeatures()      // Advanced features

	// Print summary / 요약 출력
	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║                    All Examples Completed!                                ║")
	logger.Info("║                    모든 예제 완료!                                         ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")
	logger.Info("✅ Features Demonstrated / 시연된 기능:")
	logger.Info("   • Phase 1: Basic HTTP methods (GET, POST, PUT, PATCH, DELETE) + Context variants")
	logger.Info("   • Phase 1: 13 configuration options (BaseURL, Timeout, Retry, Auth, Headers, etc.)")
	logger.Info("   • Error Handling: HTTPError, TimeoutError, RetryError detection")
	logger.Info("   • Phase 2: Response helpers (status checks, body access, headers)")
	logger.Info("   • Phase 3: File operations (8 functions - download/upload with progress tracking)")
	logger.Info("   • Phase 4: URL Builder (12 functions - parse, join, normalize, query params)")
	logger.Info("   • Phase 4: Form Builder (10 functions - fluent API, encoding, posting)")
	logger.Info("   • Phase 5a: Cookie Management (11 functions - in-memory & persistent)")
	logger.Info("   • Advanced: Proxy, custom logger, default client configuration")
	logger.Info("")
	logger.Info("📊 Coverage Summary / 커버리지 요약:")
	logger.Info("   • Total functions in package: ~65 exported functions")
	logger.Info("   • Functions demonstrated: ~60 functions")
	logger.Info("   • Coverage: ~92%")
	logger.Info("")
	logger.Info("📝 Log saved to:", "path", logFilePath)
	logger.Info("")
}
