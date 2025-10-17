package websvrutil

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// context_response.go provides HTTP response writing methods for the Context type.
//
// This file contains methods for generating various types of HTTP responses:
//
// Basic Response Writing:
//   - Status(): Set HTTP status code
//   - Write(), WriteString(): Low-level response body writing
//
// JSON Responses:
//   - JSON(): Standard JSON response with automatic marshaling
//   - JSONIndent(): JSON with custom indentation (debugging)
//   - JSONPretty(): JSON with 2-space indentation (convenience)
//
// HTML Responses:
//   - HTML(): Send raw HTML string
//   - HTMLTemplate(): Parse and execute inline template with data
//
// Text Responses:
//   - Text(): Plain text response
//   - Textf(): Formatted plain text (uses fmt.Sprintf)
//
// XML Responses:
//   - XML(): Send XML string with proper content-type
//
// Template Rendering:
//   - Render(): Render template file from template engine
//   - RenderWithLayout(): Render template with layout wrapper
//   (Requires App's template engine to be initialized)
//
// Redirects:
//   - Redirect(): HTTP redirect with status code (301, 302, 307, 308)
//   - NoContent(): Send 204 No Content (useful for DELETE)
//
// Error Responses:
//   - Error(): Convenience method for JSON error responses
//     Automatically includes error text, message, and status code
//
// All methods automatically set appropriate Content-Type headers and handle
// encoding/marshaling internally, providing a high-level API for response generation.
//
// context_response.go는 Context 타입을 위한 HTTP 응답 작성 메서드를 제공합니다.
//
// 이 파일은 다양한 유형의 HTTP 응답을 생성하는 메서드를 포함합니다:
//
// 기본 응답 작성:
//   - Status(): HTTP 상태 코드 설정
//   - Write(), WriteString(): 저수준 응답 본문 작성
//
// JSON 응답:
//   - JSON(): 자동 마샬링이 있는 표준 JSON 응답
//   - JSONIndent(): 사용자 정의 들여쓰기가 있는 JSON (디버깅)
//   - JSONPretty(): 2칸 들여쓰기가 있는 JSON (편의)
//
// HTML 응답:
//   - HTML(): 원시 HTML 문자열 전송
//   - HTMLTemplate(): 데이터로 인라인 템플릿 파싱 및 실행
//
// 텍스트 응답:
//   - Text(): 일반 텍스트 응답
//   - Textf(): 형식화된 일반 텍스트 (fmt.Sprintf 사용)
//
// XML 응답:
//   - XML(): 적절한 콘텐츠 타입과 함께 XML 문자열 전송
//
// 템플릿 렌더링:
//   - Render(): 템플릿 엔진에서 템플릿 파일 렌더링
//   - RenderWithLayout(): 레이아웃 래퍼와 함께 템플릿 렌더링
//   (App의 템플릿 엔진이 초기화되어야 함)
//
// 리다이렉트:
//   - Redirect(): 상태 코드가 있는 HTTP 리다이렉트 (301, 302, 307, 308)
//   - NoContent(): 204 No Content 전송 (DELETE에 유용)
//
// 에러 응답:
//   - Error(): JSON 에러 응답을 위한 편의 메서드
//     자동으로 에러 텍스트, 메시지, 상태 코드 포함
//
// 모든 메서드는 적절한 Content-Type 헤더를 자동으로 설정하고
// 인코딩/마샬링을 내부적으로 처리하여 응답 생성을 위한 고수준 API를 제공합니다.

// ============================================================================
// Response Writing
// 응답 작성
// ============================================================================

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

// ============================================================================
// JSON Responses
// JSON 응답
// ============================================================================

// JSON sends a JSON response with the given status code and data.
// JSON은 주어진 상태 코드와 데이터로 JSON 응답을 전송합니다.
//
// The data will be marshaled to JSON and sent with Content-Type: application/json.
// 데이터는 JSON으로 마샬링되어 Content-Type: application/json으로 전송됩니다.
//
// Example
// 예제:
//
//	ctx.JSON(200, map[string]string{"message": "success"})
func (c *Context) JSON(code int, data interface{}) error {
	c.SetHeader("Content-Type", ContentTypeJSON)
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
// Example
// 예제:
//
//	ctx.JSONIndent(200, data, "", "  ")
func (c *Context) JSONIndent(code int, data interface{}, prefix, indent string) error {
	c.SetHeader("Content-Type", ContentTypeJSON)
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
// Example
// 예제:
//
//	ctx.JSONPretty(200, data)
func (c *Context) JSONPretty(code int, data interface{}) error {
	return c.JSONIndent(code, data, "", "  ")
}

// ============================================================================
// HTML Responses
// HTML 응답
// ============================================================================

// HTML sends an HTML response with the given status code and HTML content.
// HTML은 주어진 상태 코드와 HTML 콘텐츠로 HTML 응답을 전송합니다.
//
// Example
// 예제:
//
//	ctx.HTML(200, "<h1>Hello World</h1>")
func (c *Context) HTML(code int, html string) error {
	c.SetHeader("Content-Type", ContentTypeHTML)
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
// Example
// 예제:
//
//	tmpl := "<h1>Hello {{.Name}}</h1>"
//	ctx.HTMLTemplate(200, tmpl, map[string]string{"Name": "World"})
func (c *Context) HTMLTemplate(code int, tmpl string, data interface{}) error {
	c.SetHeader("Content-Type", ContentTypeHTML)
	c.Status(code)

	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(c.ResponseWriter, data)
}

// ============================================================================
// Text Responses
// 텍스트 응답
// ============================================================================

// Text sends a plain text response.
// Text는 일반 텍스트 응답을 전송합니다.
//
// Example
// 예제:
//
//	ctx.Text(200, "Hello World")
func (c *Context) Text(code int, text string) error {
	c.SetHeader("Content-Type", ContentTypeText)
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
// Example
// 예제:
//
//	ctx.Textf(200, "Hello %s", "World")
func (c *Context) Textf(code int, format string, args ...interface{}) error {
	text := fmt.Sprintf(format, args...)
	return c.Text(code, text)
}

// ============================================================================
// XML Responses
// XML 응답
// ============================================================================

// XML sends an XML response.
// XML은 XML 응답을 전송합니다.
//
// Example
// 예제:
//
//	ctx.XML(200, "<root><message>success</message></root>")
func (c *Context) XML(code int, xml string) error {
	c.SetHeader("Content-Type", ContentTypeXML)
	c.Status(code)
	_, err := c.WriteString(xml)
	return err
}

// ============================================================================
// Template Rendering
// 템플릿 렌더링
// ============================================================================

// Render renders a template file with the given data.
// Render는 주어진 데이터로 템플릿 파일을 렌더링합니다.
//
// The template file is loaded from the template engine.
// 템플릿 파일은 템플릿 엔진에서 로드됩니다.
//
// Example
// 예제:
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
	c.SetHeader("Content-Type", ContentTypeHTML)
	c.Status(code)

	// Render template
	// 템플릿 렌더링
	return engine.Render(c.ResponseWriter, name, data)
}

// RenderWithLayout renders a template with a layout.
// RenderWithLayout는 레이아웃과 함께 템플릿을 렌더링합니다.
//
// Example
// 예제:
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
	c.SetHeader("Content-Type", ContentTypeHTML)
	c.Status(code)

	// Render template with layout
	// 레이아웃과 함께 템플릿 렌더링
	return engine.RenderWithLayout(c.ResponseWriter, layoutName, templateName, data)
}

// ============================================================================
// Redirects
// 리다이렉트
// ============================================================================

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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
//
//	ctx.Error(400, "Invalid input")
func (c *Context) Error(code int, message string) error {
	return c.JSON(code, map[string]interface{}{
		"error":   http.StatusText(code),
		"message": message,
		"status":  code,
	})
}
