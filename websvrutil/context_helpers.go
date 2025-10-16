package websvrutil

import (
	"net/http"
)

// ============================================================================
// Cookie Helpers / 쿠키 헬퍼
// ============================================================================

// CookieValue retrieves a cookie value by name.
// CookieValue는 이름으로 쿠키 값을 검색합니다.
//
// It returns the cookie value as a string. If the cookie is not found,
// it returns an empty string.
//
// 쿠키 값을 문자열로 반환합니다. 쿠키를 찾을 수 없으면 빈 문자열을 반환합니다.
//
// Example / 예제:
//
//	sessionID := ctx.CookieValue("session_id")
func (c *Context) CookieValue(name string) string {
	cookie, err := c.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

// SetCookieAdvanced sets a cookie with advanced options.
// SetCookieAdvanced는 고급 옵션으로 쿠키를 설정합니다.
//
// This method accepts a CookieOptions struct for full control over cookie attributes.
// 쿠키 속성을 완전히 제어하기 위해 CookieOptions 구조체를 받습니다.
//
// Example / 예제:
//
//	ctx.SetCookieAdvanced(CookieOptions{
//	    Name:     "session",
//	    Value:    "abc123",
//	    MaxAge:   3600,
//	    Path:     "/",
//	    HttpOnly: true,
//	    Secure:   true,
//	})
func (c *Context) SetCookieAdvanced(opts CookieOptions) {
	cookie := &http.Cookie{
		Name:     opts.Name,
		Value:    opts.Value,
		Path:     opts.Path,
		Domain:   opts.Domain,
		MaxAge:   opts.MaxAge,
		Secure:   opts.Secure,
		HttpOnly: opts.HttpOnly,
		SameSite: opts.SameSite,
	}

	// Set default path if not provided / 제공되지 않으면 기본 경로 설정
	if cookie.Path == "" {
		cookie.Path = "/"
	}

	http.SetCookie(c.ResponseWriter, cookie)
}

// CookieOptions represents options for setting a cookie.
// CookieOptions는 쿠키 설정을 위한 옵션을 나타냅니다.
type CookieOptions struct {
	// Name is the cookie name / Name은 쿠키 이름입니다
	Name string

	// Value is the cookie value / Value는 쿠키 값입니다
	Value string

	// Path is the cookie path (default: "/") / Path는 쿠키 경로입니다 (기본값: "/")
	Path string

	// Domain is the cookie domain / Domain은 쿠키 도메인입니다
	Domain string

	// MaxAge is the cookie max age in seconds / MaxAge는 쿠키 최대 수명(초)입니다
	// Use 0 for session cookies / 세션 쿠키의 경우 0 사용
	// Use -1 to delete the cookie / 쿠키를 삭제하려면 -1 사용
	MaxAge int

	// Secure indicates if cookie should only be sent over HTTPS / Secure는 쿠키가 HTTPS를 통해서만 전송되어야 하는지를 나타냅니다
	Secure bool

	// HttpOnly prevents JavaScript access to the cookie / HttpOnly는 JavaScript의 쿠키 액세스를 방지합니다
	HttpOnly bool

	// SameSite controls cross-site cookie behavior / SameSite는 크로스 사이트 쿠키 동작을 제어합니다
	SameSite http.SameSite
}

// ============================================================================
// Error Response Helpers / 에러 응답 헬퍼
// ============================================================================

// AbortWithStatus aborts the request with the specified status code.
// AbortWithStatus는 지정된 상태 코드로 요청을 중단합니다.
//
// Example / 예제:
//
//	ctx.AbortWithStatus(http.StatusUnauthorized)
func (c *Context) AbortWithStatus(code int) {
	c.Status(code)
	c.ResponseWriter.WriteHeader(code)
}

// AbortWithError aborts with status code and error message.
// AbortWithError는 상태 코드와 에러 메시지로 중단합니다.
//
// Example / 예제:
//
//	ctx.AbortWithError(http.StatusBadRequest, "Invalid input")
func (c *Context) AbortWithError(code int, message string) {
	c.Status(code)
	http.Error(c.ResponseWriter, message, code)
}

// AbortWithJSON aborts with status code and JSON error response.
// AbortWithJSON은 상태 코드와 JSON 에러 응답으로 중단합니다.
//
// Example / 예제:
//
//	ctx.AbortWithJSON(http.StatusBadRequest, map[string]string{
//	    "error": "Invalid input",
//	})
func (c *Context) AbortWithJSON(code int, obj interface{}) {
	c.JSON(code, obj)
}

// ErrorJSON sends a standardized JSON error response.
// ErrorJSON은 표준화된 JSON 에러 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.ErrorJSON(http.StatusNotFound, "User not found")
func (c *Context) ErrorJSON(code int, message string) {
	c.JSON(code, map[string]interface{}{
		"error":   message,
		"status":  code,
		"success": false,
	})
}

// SuccessJSON sends a standardized JSON success response.
// SuccessJSON은 표준화된 JSON 성공 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.SuccessJSON(http.StatusOK, "Operation completed", data)
func (c *Context) SuccessJSON(code int, message string, data interface{}) {
	c.JSON(code, map[string]interface{}{
		"message": message,
		"data":    data,
		"status":  code,
		"success": true,
	})
}

// ============================================================================
// HTTP Status Shortcuts / HTTP 상태 단축키
// ============================================================================

// NotFound sends a 404 Not Found response.
// NotFound는 404 Not Found 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.NotFound()
func (c *Context) NotFound() {
	c.AbortWithStatus(http.StatusNotFound)
}

// Unauthorized sends a 401 Unauthorized response.
// Unauthorized는 401 Unauthorized 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.Unauthorized()
func (c *Context) Unauthorized() {
	c.AbortWithStatus(http.StatusUnauthorized)
}

// Forbidden sends a 403 Forbidden response.
// Forbidden은 403 Forbidden 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.Forbidden()
func (c *Context) Forbidden() {
	c.AbortWithStatus(http.StatusForbidden)
}

// BadRequest sends a 400 Bad Request response.
// BadRequest는 400 Bad Request 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.BadRequest()
func (c *Context) BadRequest() {
	c.AbortWithStatus(http.StatusBadRequest)
}

// InternalServerError sends a 500 Internal Server Error response.
// InternalServerError는 500 Internal Server Error 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.InternalServerError()
func (c *Context) InternalServerError() {
	c.AbortWithStatus(http.StatusInternalServerError)
}
