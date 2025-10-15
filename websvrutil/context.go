package websvrutil

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
)

// Context represents the context of the current HTTP request.
// Context는 현재 HTTP 요청의 컨텍스트를 나타냅니다.
type Context struct {
	// Request is the HTTP request
	// Request는 HTTP 요청입니다
	Request *http.Request

	// ResponseWriter is the HTTP response writer
	// ResponseWriter는 HTTP 응답 작성기입니다
	ResponseWriter http.ResponseWriter

	// params stores URL path parameters
	// params는 URL 경로 매개변수를 저장합니다
	params map[string]string

	// values stores custom context values
	// values는 커스텀 컨텍스트 값을 저장합니다
	values map[string]interface{}

	// app is a reference to the App instance
	// app는 App 인스턴스에 대한 참조입니다
	app *App

	// mu protects concurrent access to values
	// mu는 values에 대한 동시 액세스를 보호합니다
	mu sync.RWMutex
}

// contextKey is the type used for context keys.
// contextKey는 컨텍스트 키에 사용되는 타입입니다.
type contextKey string

const (
	// contextKeyParams is the key for storing route parameters
	// contextKeyParams는 라우트 매개변수를 저장하는 키입니다
	contextKeyParams contextKey = "params"
)

// NewContext creates a new Context instance.
// NewContext는 새 Context 인스턴스를 생성합니다.
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:        r,
		ResponseWriter: w,
		params:         make(map[string]string),
		values:         make(map[string]interface{}),
	}
}

// Param returns the value of the URL parameter with the given name.
// Param은 주어진 이름의 URL 매개변수 값을 반환합니다.
//
// Example / 예제:
//
//	// Route: /users/:id
//	// URL: /users/123
//	id := ctx.Param("id") // Returns "123" / "123" 반환
func (c *Context) Param(name string) string {
	return c.params[name]
}

// Params returns all URL parameters as a map.
// Params는 모든 URL 매개변수를 맵으로 반환합니다.
func (c *Context) Params() map[string]string {
	// Return a copy to prevent external modification
	// 외부 수정을 방지하기 위해 복사본 반환
	result := make(map[string]string, len(c.params))
	for k, v := range c.params {
		result[k] = v
	}
	return result
}

// setParams sets the URL parameters (internal use only).
// setParams는 URL 매개변수를 설정합니다 (내부 사용 전용).
func (c *Context) setParams(params map[string]string) {
	c.params = params
}

// Set stores a value in the context.
// Set은 컨텍스트에 값을 저장합니다.
//
// Example / 예제:
//
//	ctx.Set("user", user)
//	ctx.Set("requestID", "12345")
func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
}

// Get retrieves a value from the context.
// Get은 컨텍스트에서 값을 검색합니다.
//
// Example / 예제:
//
//	user, exists := ctx.Get("user")
//	if exists {
//	    // Use user
//	}
func (c *Context) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.values[key]
	return value, exists
}

// MustGet retrieves a value from the context and panics if it doesn't exist.
// MustGet은 컨텍스트에서 값을 검색하고 존재하지 않으면 패닉합니다.
//
// Example / 예제:
//
//	user := ctx.MustGet("user").(User)
func (c *Context) MustGet(key string) interface{} {
	value, exists := c.Get(key)
	if !exists {
		panic("key not found: " + key)
	}
	return value
}

// GetString retrieves a string value from the context.
// GetString은 컨텍스트에서 문자열 값을 검색합니다.
func (c *Context) GetString(key string) string {
	value, exists := c.Get(key)
	if !exists {
		return ""
	}
	str, _ := value.(string)
	return str
}

// GetInt retrieves an int value from the context.
// GetInt은 컨텍스트에서 int 값을 검색합니다.
func (c *Context) GetInt(key string) int {
	value, exists := c.Get(key)
	if !exists {
		return 0
	}
	i, _ := value.(int)
	return i
}

// GetBool retrieves a bool value from the context.
// GetBool은 컨텍스트에서 bool 값을 검색합니다.
func (c *Context) GetBool(key string) bool {
	value, exists := c.Get(key)
	if !exists {
		return false
	}
	b, _ := value.(bool)
	return b
}

// Context returns the request's context.Context.
// Context는 요청의 context.Context를 반환합니다.
func (c *Context) Context() context.Context {
	return c.Request.Context()
}

// WithContext returns a shallow copy of Context with a new context.Context.
// WithContext는 새 context.Context를 가진 Context의 얕은 복사본을 반환합니다.
func (c *Context) WithContext(ctx context.Context) *Context {
	c2 := *c
	c2.Request = c.Request.WithContext(ctx)
	return &c2
}

// Method returns the HTTP request method.
// Method는 HTTP 요청 메서드를 반환합니다.
func (c *Context) Method() string {
	return c.Request.Method
}

// Path returns the URL path.
// Path는 URL 경로를 반환합니다.
func (c *Context) Path() string {
	return c.Request.URL.Path
}

// Query returns the query string parameter with the given name.
// Query는 주어진 이름의 쿼리 문자열 매개변수를 반환합니다.
//
// Example / 예제:
//
//	// URL: /search?q=golang&page=2
//	q := ctx.Query("q")       // Returns "golang" / "golang" 반환
//	page := ctx.Query("page") // Returns "2" / "2" 반환
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// QueryDefault returns the query string parameter with the given name,
// or the default value if it doesn't exist.
// QueryDefault는 주어진 이름의 쿼리 문자열 매개변수를 반환하거나,
// 존재하지 않으면 기본값을 반환합니다.
func (c *Context) QueryDefault(key, defaultValue string) string {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Header returns the request header with the given key.
// Header는 주어진 키의 요청 헤더를 반환합니다.
func (c *Context) Header(key string) string {
	return c.Request.Header.Get(key)
}

// SetHeader sets a response header.
// SetHeader는 응답 헤더를 설정합니다.
func (c *Context) SetHeader(key, value string) {
	c.ResponseWriter.Header().Set(key, value)
}

// Status sets the HTTP response status code.
// Status는 HTTP 응답 상태 코드를 설정합니다.
func (c *Context) Status(code int) {
	c.ResponseWriter.WriteHeader(code)
}

// Write writes data to the response body.
// Write는 응답 본문에 데이터를 씁니다.
func (c *Context) Write(data []byte) (int, error) {
	return c.ResponseWriter.Write(data)
}

// WriteString writes a string to the response body.
// WriteString은 응답 본문에 문자열을 씁니다.
func (c *Context) WriteString(s string) (int, error) {
	return c.ResponseWriter.Write([]byte(s))
}

// GetContext retrieves the Context from the request's context.Context.
// GetContext는 요청의 context.Context에서 Context를 검색합니다.
//
// Example / 예제:
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//	    ctx := websvrutil.GetContext(r)
//	    id := ctx.Param("id")
//	}
func GetContext(r *http.Request) *Context {
	value := r.Context().Value(contextKeyParams)
	if value == nil {
		// Return empty context if not found
		// 찾을 수 없으면 빈 컨텍스트 반환
		return NewContext(nil, r)
	}
	ctx, ok := value.(*Context)
	if !ok {
		return NewContext(nil, r)
	}
	return ctx
}

// JSON sends a JSON response with the given status code and data.
// JSON은 주어진 상태 코드와 데이터로 JSON 응답을 전송합니다.
//
// The data will be marshaled to JSON and sent with Content-Type: application/json.
// 데이터는 JSON으로 마샬링되어 Content-Type: application/json으로 전송됩니다.
//
// Example / 예제:
//
//	ctx.JSON(200, map[string]string{"message": "success"})
func (c *Context) JSON(code int, data interface{}) error {
	c.SetHeader("Content-Type", "application/json; charset=utf-8")
	c.Status(code)

	encoder := json.NewEncoder(c.ResponseWriter)
	return encoder.Encode(data)
}

// JSONIndent sends a JSON response with indentation for readability.
// JSONIndent는 가독성을 위해 들여쓰기가 있는 JSON 응답을 전송합니다.
//
// This is useful for debugging or development. For production, use JSON() instead.
// 디버깅이나 개발에 유용합니다. 프로덕션에서는 JSON()을 사용하세요.
//
// Example / 예제:
//
//	ctx.JSONIndent(200, data, "", "  ")
func (c *Context) JSONIndent(code int, data interface{}, prefix, indent string) error {
	c.SetHeader("Content-Type", "application/json; charset=utf-8")
	c.Status(code)

	encoder := json.NewEncoder(c.ResponseWriter)
	encoder.SetIndent(prefix, indent)
	return encoder.Encode(data)
}

// JSONPretty sends a JSON response with pretty-printing (2-space indentation).
// JSONPretty는 보기 좋게 출력된 JSON 응답을 전송합니다 (2칸 들여쓰기).
//
// This is a convenience wrapper around JSONIndent with default indentation.
// 기본 들여쓰기가 있는 JSONIndent의 편의 래퍼입니다.
//
// Example / 예제:
//
//	ctx.JSONPretty(200, data)
func (c *Context) JSONPretty(code int, data interface{}) error {
	return c.JSONIndent(code, data, "", "  ")
}

// HTML sends an HTML response with the given status code and HTML content.
// HTML은 주어진 상태 코드와 HTML 콘텐츠로 HTML 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.HTML(200, "<h1>Hello World</h1>")
func (c *Context) HTML(code int, html string) error {
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.Status(code)
	_, err := c.WriteString(html)
	return err
}

// HTMLTemplate renders an HTML template with the given data.
// HTMLTemplate은 주어진 데이터로 HTML 템플릿을 렌더링합니다.
//
// The template is parsed and executed with the provided data.
// 템플릿은 제공된 데이터로 파싱되고 실행됩니다.
//
// Example / 예제:
//
//	tmpl := "<h1>Hello {{.Name}}</h1>"
//	ctx.HTMLTemplate(200, tmpl, map[string]string{"Name": "World"})
func (c *Context) HTMLTemplate(code int, tmpl string, data interface{}) error {
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.Status(code)

	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(c.ResponseWriter, data)
}

// Text sends a plain text response.
// Text는 일반 텍스트 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.Text(200, "Hello World")
func (c *Context) Text(code int, text string) error {
	c.SetHeader("Content-Type", "text/plain; charset=utf-8")
	c.Status(code)
	_, err := c.WriteString(text)
	return err
}

// Textf sends a formatted plain text response.
// Textf는 형식화된 일반 텍스트 응답을 전송합니다.
//
// This uses fmt.Sprintf for formatting.
// fmt.Sprintf를 사용하여 형식화합니다.
//
// Example / 예제:
//
//	ctx.Textf(200, "Hello %s", "World")
func (c *Context) Textf(code int, format string, args ...interface{}) error {
	text := fmt.Sprintf(format, args...)
	return c.Text(code, text)
}

// XML sends an XML response.
// XML은 XML 응답을 전송합니다.
//
// Example / 예제:
//
//	ctx.XML(200, "<root><message>success</message></root>")
func (c *Context) XML(code int, xml string) error {
	c.SetHeader("Content-Type", "application/xml; charset=utf-8")
	c.Status(code)
	_, err := c.WriteString(xml)
	return err
}

// Redirect sends an HTTP redirect response.
// Redirect는 HTTP 리다이렉트 응답을 전송합니다.
//
// Common status codes:
// - 301: Moved Permanently
// - 302: Found (temporary redirect)
// - 303: See Other
// - 307: Temporary Redirect
// - 308: Permanent Redirect
//
// 일반적인 상태 코드:
// - 301: 영구 이동
// - 302: 발견 (임시 리다이렉트)
// - 303: 다른 것 보기
// - 307: 임시 리다이렉트
// - 308: 영구 리다이렉트
//
// Example / 예제:
//
//	ctx.Redirect(302, "/new-url")
func (c *Context) Redirect(code int, url string) {
	http.Redirect(c.ResponseWriter, c.Request, url, code)
}

// NoContent sends a 204 No Content response.
// NoContent는 204 No Content 응답을 전송합니다.
//
// This is commonly used for successful DELETE requests or when no response body is needed.
// DELETE 요청이 성공했거나 응답 본문이 필요없을 때 일반적으로 사용됩니다.
//
// Example / 예제:
//
//	ctx.NoContent()
func (c *Context) NoContent() {
	c.Status(http.StatusNoContent)
}

// Error sends an error response with the given status code and message.
// Error는 주어진 상태 코드와 메시지로 에러 응답을 전송합니다.
//
// This is a convenience method for sending JSON error responses.
// JSON 에러 응답 전송을 위한 편의 메서드입니다.
//
// Example / 예제:
//
//	ctx.Error(400, "Invalid input")
func (c *Context) Error(code int, message string) error {
	return c.JSON(code, map[string]interface{}{
		"error":   http.StatusText(code),
		"message": message,
		"status":  code,
	})
}

// Render renders a template file with the given data.
// Render는 주어진 데이터로 템플릿 파일을 렌더링합니다.
//
// The template file is loaded from the template engine.
// 템플릿 파일은 템플릿 엔진에서 로드됩니다.
//
// Example / 예제:
//
//	ctx.Render(200, "index.html", map[string]string{"Title": "Home"})
func (c *Context) Render(code int, name string, data interface{}) error {
	// Get app from request context
	// 요청 컨텍스트에서 앱 가져오기
	app, ok := c.Request.Context().Value("app").(*App)
	if !ok || app == nil {
		return fmt.Errorf("app not found in context")
	}

	// Get template engine
	// 템플릿 엔진 가져오기
	engine := app.TemplateEngine()
	if engine == nil {
		return fmt.Errorf("template engine not initialized")
	}

	// Set content type and status
	// Content-Type 및 상태 설정
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.Status(code)

	// Render template
	// 템플릿 렌더링
	return engine.Render(c.ResponseWriter, name, data)
}

// RenderWithLayout renders a template with a layout.
// RenderWithLayout는 레이아웃과 함께 템플릿을 렌더링합니다.
//
// Example / 예제:
//
//	ctx.RenderWithLayout(200, "base.html", "index.html", map[string]string{"Title": "Home"})
func (c *Context) RenderWithLayout(code int, layoutName, templateName string, data interface{}) error {
	// Get app from request context
	// 요청 컨텍스트에서 앱 가져오기
	app, ok := c.Request.Context().Value("app").(*App)
	if !ok || app == nil {
		return fmt.Errorf("app not found in context")
	}

	// Get template engine
	// 템플릿 엔진 가져오기
	engine := app.TemplateEngine()
	if engine == nil {
		return fmt.Errorf("template engine not initialized")
	}

	// Set content type and status
	// Content-Type 및 상태 설정
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.Status(code)

	// Render template with layout
	// 레이아웃과 함께 템플릿 렌더링
	return engine.RenderWithLayout(c.ResponseWriter, layoutName, templateName, data)
}

// BindJSON binds the request body as JSON to the provided struct.
// BindJSON은 요청 본문을 JSON으로 제공된 구조체에 바인딩합니다.
//
// Example / 예제:
//
//	var user User
//	if err := ctx.BindJSON(&user); err != nil {
//	    return ctx.Error(400, "Invalid JSON")
//	}
func (c *Context) BindJSON(obj interface{}) error {
	if c.Request.Body == nil {
		return fmt.Errorf("request body is nil")
	}

	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(obj); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
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
	if contentType == "application/json" || contentType == "application/json; charset=utf-8" {
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

// AddHeader adds a header value to the response.
// AddHeader는 응답에 헤더 값을 추가합니다.
//
// Unlike SetHeader, this appends the value if the header already exists.
// SetHeader와 달리 헤더가 이미 존재하는 경우 값을 추가합니다.
//
// Example / 예제:
//
//	ctx.AddHeader("Set-Cookie", "cookie1=value1")
//	ctx.AddHeader("Set-Cookie", "cookie2=value2")
func (c *Context) AddHeader(key, value string) {
	c.ResponseWriter.Header().Add(key, value)
}

// GetHeader returns the request header with the given key.
// GetHeader는 주어진 키의 요청 헤더를 반환합니다.
//
// This is an alias for Header() for consistency.
// 일관성을 위한 Header()의 별칭입니다.
//
// Example / 예제:
//
//	userAgent := ctx.GetHeader("User-Agent")
func (c *Context) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

// GetHeaders returns all values for the given header key.
// GetHeaders는 주어진 헤더 키의 모든 값을 반환합니다.
//
// Example / 예제:
//
//	acceptEncodings := ctx.GetHeaders("Accept-Encoding")
func (c *Context) GetHeaders(key string) []string {
	return c.Request.Header.Values(key)
}

// HeaderExists checks if a request header exists.
// HeaderExists는 요청 헤더가 존재하는지 확인합니다.
//
// Example / 예제:
//
//	if ctx.HeaderExists("Authorization") {
//	    // Process authentication
//	}
func (c *Context) HeaderExists(key string) bool {
	_, exists := c.Request.Header[key]
	return exists
}

// ContentType returns the Content-Type header of the request.
// ContentType은 요청의 Content-Type 헤더를 반환합니다.
//
// Example / 예제:
//
//	contentType := ctx.ContentType()
func (c *Context) ContentType() string {
	return c.Request.Header.Get("Content-Type")
}

// UserAgent returns the User-Agent header of the request.
// UserAgent는 요청의 User-Agent 헤더를 반환합니다.
//
// Example / 예제:
//
//	userAgent := ctx.UserAgent()
func (c *Context) UserAgent() string {
	return c.Request.Header.Get("User-Agent")
}

// Referer returns the Referer header of the request.
// Referer는 요청의 Referer 헤더를 반환합니다.
//
// Example / 예제:
//
//	referer := ctx.Referer()
func (c *Context) Referer() string {
	return c.Request.Header.Get("Referer")
}

// ClientIP returns the client IP address.
// ClientIP는 클라이언트 IP 주소를 반환합니다.
//
// It checks X-Forwarded-For, X-Real-IP headers first, then falls back to RemoteAddr.
// X-Forwarded-For, X-Real-IP 헤더를 먼저 확인한 후 RemoteAddr로 대체합니다.
//
// Example / 예제:
//
//	ip := ctx.ClientIP()
func (c *Context) ClientIP() string {
	// Check X-Forwarded-For header
	// X-Forwarded-For 헤더 확인
	if xff := c.Request.Header.Get("X-Forwarded-For"); xff != "" {
		// Return the first IP in the list
		// 목록의 첫 번째 IP 반환
		for idx := 0; idx < len(xff); idx++ {
			if xff[idx] == ',' {
				return xff[:idx]
			}
		}
		return xff
	}

	// Check X-Real-IP header
	// X-Real-IP 헤더 확인
	if xri := c.Request.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Fall back to RemoteAddr
	// RemoteAddr로 대체
	for idx := 0; idx < len(c.Request.RemoteAddr); idx++ {
		if c.Request.RemoteAddr[idx] == ':' {
			return c.Request.RemoteAddr[:idx]
		}
	}
	return c.Request.RemoteAddr
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
	maxSize := int64(32 << 20) // 32 MB default
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
