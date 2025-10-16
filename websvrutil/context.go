package websvrutil

import (
	"context"
	"sync"
	"net/http"
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
//
// Performance optimization / 성능 최적화:
//   - values map is lazily allocated (nil by default)
//   - values 맵은 지연 할당됩니다 (기본적으로 nil)
//   - Only created when first value is set via Set()
//   - Set()을 통해 첫 번째 값이 설정될 때만 생성됩니다
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:        r,
		ResponseWriter: w,
		params:         make(map[string]string),
		values:         nil, // Lazy allocation in Set()
	}
}

// ============================================================================
// URL Parameters / URL 매개변수
// ============================================================================

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

// ============================================================================
// Context Values / 컨텍스트 값
// ============================================================================

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

	// Lazy map initialization / 지연 맵 초기화
	// Only create the map when first value is set
	// 첫 번째 값이 설정될 때만 맵 생성
	if c.values == nil {
		c.values = make(map[string]interface{})
	}

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

// GetInt64 retrieves an int64 value from the context.
// GetInt64는 컨텍스트에서 int64 값을 검색합니다.
func (c *Context) GetInt64(key string) int64 {
	value, exists := c.Get(key)
	if !exists {
		return 0
	}
	i64, _ := value.(int64)
	return i64
}

// GetFloat64 retrieves a float64 value from the context.
// GetFloat64는 컨텍스트에서 float64 값을 검색합니다.
func (c *Context) GetFloat64(key string) float64 {
	value, exists := c.Get(key)
	if !exists {
		return 0.0
	}
	f64, _ := value.(float64)
	return f64
}

// GetStringSlice retrieves a []string value from the context.
// GetStringSlice는 컨텍스트에서 []string 값을 검색합니다.
func (c *Context) GetStringSlice(key string) []string {
	value, exists := c.Get(key)
	if !exists {
		return nil
	}
	slice, _ := value.([]string)
	return slice
}

// GetStringMap retrieves a map[string]interface{} value from the context.
// GetStringMap은 컨텍스트에서 map[string]interface{} 값을 검색합니다.
func (c *Context) GetStringMap(key string) map[string]interface{} {
	value, exists := c.Get(key)
	if !exists {
		return nil
	}
	m, _ := value.(map[string]interface{})
	return m
}

// Exists checks if a key exists in the context.
// Exists는 컨텍스트에 키가 존재하는지 확인합니다.
func (c *Context) Exists(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.values[key]
	return exists
}

// Delete removes a value from the context.
// Delete는 컨텍스트에서 값을 제거합니다.
func (c *Context) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.values, key)
}

// Keys returns all keys in the context.
// Keys는 컨텍스트의 모든 키를 반환합니다.
func (c *Context) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	keys := make([]string, 0, len(c.values))
	for key := range c.values {
		keys = append(keys, key)
	}
	return keys
}

// ============================================================================
// Go Context / Go 컨텍스트
// ============================================================================

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

// ============================================================================
// Context Retrieval / 컨텍스트 검색
// ============================================================================

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
