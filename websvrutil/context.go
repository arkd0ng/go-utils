package websvrutil

import (
	"context"
	"net/http"
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
