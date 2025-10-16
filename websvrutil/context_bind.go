package websvrutil

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// ============================================================================
// Data Binding / 데이터 바인딩
// ============================================================================

// BindJSON binds the request body as JSON to the provided struct.
// BindJSON은 요청 본문을 JSON으로 제공된 구조체에 바인딩합니다.
//
// Body size limit / 본문 크기 제한:
//   - Enforces maximum body size from App.options.MaxBodySize
//   - App.options.MaxBodySize에서 최대 본문 크기 강제 적용
//   - Default: 10 MB (configurable with WithMaxBodySize option)
//   - 기본값: 10 MB (WithMaxBodySize 옵션으로 설정 가능)
//   - Returns error if request body exceeds limit
//   - 요청 본문이 제한을 초과하면 에러 반환
//
// Security considerations / 보안 고려사항:
//   - Prevents denial-of-service attacks with large payloads
//   - 대용량 페이로드를 사용한 서비스 거부 공격 방지
//   - Uses io.LimitReader to enforce limit at read level
//   - io.LimitReader를 사용하여 읽기 수준에서 제한 강제 적용
//
// Example / 예제:
//
//	var user User
//	if err := ctx.BindJSON(&user); err != nil {
//	    return ctx.Error(400, "Invalid JSON")
//	}
//
// Custom limit / 커스텀 제한:
//
//	app := websvrutil.New(
//	    websvrutil.WithMaxBodySize(5 * 1024 * 1024), // 5 MB
//	)
func (c *Context) BindJSON(obj interface{}) error {
	if c.Request.Body == nil {
		return fmt.Errorf("request body is nil")
	}

	// Get max body size from app options, default to DefaultMaxBodySize
	// 앱 옵션에서 최대 본문 크기 가져오기, 기본값은 DefaultMaxBodySize
	maxBodySize := int64(DefaultMaxBodySize)
	if c.app != nil && c.app.options != nil && c.app.options.MaxBodySize > 0 {
		maxBodySize = c.app.options.MaxBodySize
	}

	// Use io.LimitReader to enforce body size limit
	// io.LimitReader를 사용하여 본문 크기 제한 강제 적용
	limitedReader := io.LimitReader(c.Request.Body, maxBodySize+1) // +1 to detect over-limit

	decoder := json.NewDecoder(limitedReader)
	if err := decoder.Decode(obj); err != nil {
		// Check if error is due to body size limit
		// 에러가 본문 크기 제한 때문인지 확인
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			// Try to read one more byte to confirm size limit exceeded
			// 크기 제한 초과 확인을 위해 한 바이트 더 읽기 시도
			var buf [1]byte
			if n, _ := limitedReader.Read(buf[:]); n > 0 {
				return fmt.Errorf("request body too large (exceeds %d bytes)", maxBodySize)
			}
		}
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	// Check if there's more data (body size exceeded)
	// 더 많은 데이터가 있는지 확인 (본문 크기 초과)
	var buf [1]byte
	if n, _ := limitedReader.Read(buf[:]); n > 0 {
		return fmt.Errorf("request body too large (exceeds %d bytes)", maxBodySize)
	}

	return nil
}

// BindForm binds the form data to the provided struct.
// BindForm은 폼 데이터를 제공된 구조체에 바인딩합니다.
//
// The struct should use `form` tags to specify form field names.
// 구조체는 `form` 태그를 사용하여 폼 필드 이름을 지정해야 합니다.
//
// Example / 예제:
//
//	type LoginForm struct {
//	    Username string `form:"username"`
//	    Password string `form:"password"`
//	}
//	var form LoginForm
//	if err := ctx.BindForm(&form); err != nil {
//	    return ctx.Error(400, "Invalid form data")
//	}
func (c *Context) BindForm(obj interface{}) error {
	if err := c.Request.ParseForm(); err != nil {
		return fmt.Errorf("failed to parse form: %w", err)
	}

	return bindFormData(obj, c.Request.Form)
}

// BindQuery binds the query parameters to the provided struct.
// BindQuery는 쿼리 매개변수를 제공된 구조체에 바인딩합니다.
//
// The struct should use `form` tags to specify query parameter names.
// 구조체는 `form` 태그를 사용하여 쿼리 매개변수 이름을 지정해야 합니다.
//
// Example / 예제:
//
//	type SearchQuery struct {
//	    Q    string `form:"q"`
//	    Page int    `form:"page"`
//	}
//	var query SearchQuery
//	if err := ctx.BindQuery(&query); err != nil {
//	    return ctx.Error(400, "Invalid query parameters")
//	}
func (c *Context) BindQuery(obj interface{}) error {
	return bindFormData(obj, c.Request.URL.Query())
}

// Bind automatically binds the request data based on Content-Type.
// Bind는 Content-Type에 따라 요청 데이터를 자동으로 바인딩합니다.
//
// It supports JSON (application/json) and form data (application/x-www-form-urlencoded, multipart/form-data).
// JSON (application/json) 및 폼 데이터 (application/x-www-form-urlencoded, multipart/form-data)를 지원합니다.
//
// Example / 예제:
//
//	var data RequestData
//	if err := ctx.Bind(&data); err != nil {
//	    return ctx.Error(400, "Invalid request data")
//	}
func (c *Context) Bind(obj interface{}) error {
	contentType := c.Request.Header.Get("Content-Type")

	// Check for JSON content type
	// JSON Content-Type 확인
	if contentType == "application/json" || contentType == ContentTypeJSON {
		return c.BindJSON(obj)
	}

	// Check for form content types
	// 폼 Content-Type 확인
	if contentType == "application/x-www-form-urlencoded" || contentType == "multipart/form-data" {
		return c.BindForm(obj)
	}

	// Default to form binding if no content type specified
	// Content-Type이 지정되지 않은 경우 폼 바인딩을 기본값으로 사용
	return c.BindForm(obj)
}

// ============================================================================
// Cookie Operations / 쿠키 작업
// ============================================================================

// Cookie returns the named cookie provided in the request.
// Cookie는 요청에서 제공된 이름이 지정된 쿠키를 반환합니다.
//
// Example / 예제:
//
//	cookie, err := ctx.Cookie("session_id")
//	if err != nil {
//	    // Cookie not found
//	}
func (c *Context) Cookie(name string) (*http.Cookie, error) {
	return c.Request.Cookie(name)
}

// SetCookie adds a Set-Cookie header to the response.
// SetCookie는 응답에 Set-Cookie 헤더를 추가합니다.
//
// Example / 예제:
//
//	cookie := &http.Cookie{
//	    Name:     "session_id",
//	    Value:    "abc123",
//	    Path:     "/",
//	    MaxAge:   3600,
//	    HttpOnly: true,
//	    Secure:   true,
//	}
//	ctx.SetCookie(cookie)
func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.ResponseWriter, cookie)
}

// DeleteCookie deletes a cookie by setting its MaxAge to -1.
// DeleteCookie는 MaxAge를 -1로 설정하여 쿠키를 삭제합니다.
//
// Example / 예제:
//
//	ctx.DeleteCookie("session_id", "/")
func (c *Context) DeleteCookie(name, path string) {
	cookie := &http.Cookie{
		Name:   name,
		Value:  "",
		Path:   path,
		MaxAge: -1,
	}
	http.SetCookie(c.ResponseWriter, cookie)
}

// GetCookie is a convenience method to get a cookie value.
// GetCookie는 쿠키 값을 가져오는 편의 메서드입니다.
//
// Example / 예제:
//
//	value := ctx.GetCookie("session_id")
func (c *Context) GetCookie(name string) string {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

// ============================================================================
// File Upload / 파일 업로드
// ============================================================================

// FormFile retrieves the first file for the provided form key.
// FormFile은 제공된 폼 키에 대한 첫 번째 파일을 가져옵니다.
func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	file, header, err := c.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return header, nil
}

// MultipartForm returns the parsed multipart form, including file uploads.
// MultipartForm은 파일 업로드를 포함한 파싱된 multipart 폼을 반환합니다.
func (c *Context) MultipartForm() (*multipart.Form, error) {
	maxSize := int64(DefaultMaxUploadSize)
	if c.app != nil && c.app.options != nil {
		maxSize = c.app.options.MaxUploadSize
	}
	if err := c.Request.ParseMultipartForm(maxSize); err != nil {
		return nil, err
	}
	return c.Request.MultipartForm, nil
}

// SaveUploadedFile saves the uploaded file to the destination path.
// SaveUploadedFile은 업로드된 파일을 대상 경로에 저장합니다.
func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// ============================================================================
// Static File Serving / 정적 파일 서빙
// ============================================================================

// File sends a file response to the client.
// File은 클라이언트에게 파일 응답을 전송합니다.
//
// The filepath should be the absolute or relative path to the file.
// filepath는 파일의 절대 경로 또는 상대 경로여야 합니다.
//
// Example / 예제:
//
//	ctx.File("./public/index.html")
func (c *Context) File(filepath string) error {
	http.ServeFile(c.ResponseWriter, c.Request, filepath)
	return nil
}

// FileAttachment sends a file as a downloadable attachment.
// FileAttachment는 파일을 다운로드 가능한 첨부 파일로 전송합니다.
//
// The filename parameter sets the name shown in the download dialog.
// filename 매개변수는 다운로드 대화상자에 표시되는 이름을 설정합니다.
//
// Example / 예제:
//
//	ctx.FileAttachment("./reports/report.pdf", "monthly-report.pdf")
func (c *Context) FileAttachment(filepath, filename string) error {
	c.SetHeader("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))
	http.ServeFile(c.ResponseWriter, c.Request, filepath)
	return nil
}
