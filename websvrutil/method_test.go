package websvrutil

import (
	"net/http/httptest"
	"testing"
)

// TestContextIsGET tests IsGET method / IsGET 메서드 테스트
func TestContextIsGET(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	ctx := &Context{Request: req}

	if !ctx.IsGET() {
		t.Error("Expected IsGET() to return true for GET request")
	}

	req = httptest.NewRequest("POST", "/", nil)
	ctx = &Context{Request: req}

	if ctx.IsGET() {
		t.Error("Expected IsGET() to return false for POST request")
	}
}

// TestContextIsPOST tests IsPOST method / IsPOST 메서드 테스트
func TestContextIsPOST(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	ctx := &Context{Request: req}

	if !ctx.IsPOST() {
		t.Error("Expected IsPOST() to return true for POST request")
	}

	req = httptest.NewRequest("GET", "/", nil)
	ctx = &Context{Request: req}

	if ctx.IsPOST() {
		t.Error("Expected IsPOST() to return false for GET request")
	}
}

// TestContextIsPUT tests IsPUT method / IsPUT 메서드 테스트
func TestContextIsPUT(t *testing.T) {
	req := httptest.NewRequest("PUT", "/", nil)
	ctx := &Context{Request: req}

	if !ctx.IsPUT() {
		t.Error("Expected IsPUT() to return true for PUT request")
	}

	req = httptest.NewRequest("POST", "/", nil)
	ctx = &Context{Request: req}

	if ctx.IsPUT() {
		t.Error("Expected IsPUT() to return false for POST request")
	}
}

// TestContextIsPATCH tests IsPATCH method / IsPATCH 메서드 테스트
func TestContextIsPATCH(t *testing.T) {
	req := httptest.NewRequest("PATCH", "/", nil)
	ctx := &Context{Request: req}

	if !ctx.IsPATCH() {
		t.Error("Expected IsPATCH() to return true for PATCH request")
	}

	req = httptest.NewRequest("PUT", "/", nil)
	ctx = &Context{Request: req}

	if ctx.IsPATCH() {
		t.Error("Expected IsPATCH() to return false for PUT request")
	}
}

// TestContextIsDELETE tests IsDELETE method / IsDELETE 메서드 테스트
func TestContextIsDELETE(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/", nil)
	ctx := &Context{Request: req}

	if !ctx.IsDELETE() {
		t.Error("Expected IsDELETE() to return true for DELETE request")
	}

	req = httptest.NewRequest("POST", "/", nil)
	ctx = &Context{Request: req}

	if ctx.IsDELETE() {
		t.Error("Expected IsDELETE() to return false for POST request")
	}
}

// TestContextIsHEAD tests IsHEAD method / IsHEAD 메서드 테스트
func TestContextIsHEAD(t *testing.T) {
	req := httptest.NewRequest("HEAD", "/", nil)
	ctx := &Context{Request: req}

	if !ctx.IsHEAD() {
		t.Error("Expected IsHEAD() to return true for HEAD request")
	}

	req = httptest.NewRequest("GET", "/", nil)
	ctx = &Context{Request: req}

	if ctx.IsHEAD() {
		t.Error("Expected IsHEAD() to return false for GET request")
	}
}

// TestContextIsOPTIONS tests IsOPTIONS method / IsOPTIONS 메서드 테스트
func TestContextIsOPTIONS(t *testing.T) {
	req := httptest.NewRequest("OPTIONS", "/", nil)
	ctx := &Context{Request: req}

	if !ctx.IsOPTIONS() {
		t.Error("Expected IsOPTIONS() to return true for OPTIONS request")
	}

	req = httptest.NewRequest("GET", "/", nil)
	ctx = &Context{Request: req}

	if ctx.IsOPTIONS() {
		t.Error("Expected IsOPTIONS() to return false for GET request")
	}
}

// TestContextIsAjax tests IsAjax method / IsAjax 메서드 테스트
func TestContextIsAjax(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	ctx := &Context{Request: req}

	if !ctx.IsAjax() {
		t.Error("Expected IsAjax() to return true when X-Requested-With is XMLHttpRequest")
	}

	req = httptest.NewRequest("GET", "/", nil)
	ctx = &Context{Request: req}

	if ctx.IsAjax() {
		t.Error("Expected IsAjax() to return false without X-Requested-With header")
	}
}

// TestContextIsWebSocket tests IsWebSocket method / IsWebSocket 메서드 테스트
func TestContextIsWebSocket(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Upgrade", "websocket")
	ctx := &Context{Request: req}

	if !ctx.IsWebSocket() {
		t.Error("Expected IsWebSocket() to return true when Upgrade is websocket")
	}

	req = httptest.NewRequest("GET", "/", nil)
	ctx = &Context{Request: req}

	if ctx.IsWebSocket() {
		t.Error("Expected IsWebSocket() to return false without Upgrade header")
	}
}

// TestContextAcceptsJSON tests AcceptsJSON method / AcceptsJSON 메서드 테스트
func TestContextAcceptsJSON(t *testing.T) {
	tests := []struct {
		name   string
		accept string
		want   bool
	}{
		{"exact match", "application/json", true},
		{"wildcard", "*/*", true},
		{"with charset", "application/json; charset=utf-8", true},
		{"in list", "text/html, application/json, text/plain", true},
		{"no match", "text/html", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.accept != "" {
				req.Header.Set("Accept", tt.accept)
			}
			ctx := &Context{Request: req}

			if got := ctx.AcceptsJSON(); got != tt.want {
				t.Errorf("AcceptsJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestContextAcceptsHTML tests AcceptsHTML method / AcceptsHTML 메서드 테스트
func TestContextAcceptsHTML(t *testing.T) {
	tests := []struct {
		name   string
		accept string
		want   bool
	}{
		{"exact match", "text/html", true},
		{"wildcard", "*/*", true},
		{"with charset", "text/html; charset=utf-8", true},
		{"in list", "text/html, application/json", true},
		{"no match", "application/json", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.accept != "" {
				req.Header.Set("Accept", tt.accept)
			}
			ctx := &Context{Request: req}

			if got := ctx.AcceptsHTML(); got != tt.want {
				t.Errorf("AcceptsHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestContextAcceptsXML tests AcceptsXML method / AcceptsXML 메서드 테스트
func TestContextAcceptsXML(t *testing.T) {
	tests := []struct {
		name   string
		accept string
		want   bool
	}{
		{"application/xml", "application/xml", true},
		{"text/xml", "text/xml", true},
		{"wildcard", "*/*", true},
		{"with charset", "application/xml; charset=utf-8", true},
		{"in list", "text/html, application/xml", true},
		{"no match", "application/json", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.accept != "" {
				req.Header.Set("Accept", tt.accept)
			}
			ctx := &Context{Request: req}

			if got := ctx.AcceptsXML(); got != tt.want {
				t.Errorf("AcceptsXML() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestContextContainsContentType tests containsContentType method / containsContentType 메서드 테스트
func TestContextContainsContentType(t *testing.T) {
	ctx := &Context{}

	tests := []struct {
		accept      string
		contentType string
		want        bool
	}{
		{"application/json", "application/json", true},
		{"text/html, application/json", "application/json", true},
		{"application/json; charset=utf-8", "application/json", true},
		{"text/html", "application/json", false},
		{"", "application/json", false},
	}

	for _, tt := range tests {
		t.Run(tt.accept, func(t *testing.T) {
			if got := ctx.containsContentType(tt.accept, tt.contentType); got != tt.want {
				t.Errorf("containsContentType(%q, %q) = %v, want %v", tt.accept, tt.contentType, got, tt.want)
			}
		})
	}
}

// BenchmarkIsGET benchmarks IsGET method / IsGET 메서드 벤치마크
func BenchmarkIsGET(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	ctx := &Context{Request: req}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.IsGET()
	}
}

// BenchmarkIsAjax benchmarks IsAjax method / IsAjax 메서드 벤치마크
func BenchmarkIsAjax(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	ctx := &Context{Request: req}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.IsAjax()
	}
}

// BenchmarkAcceptsJSON benchmarks AcceptsJSON method / AcceptsJSON 메서드 벤치마크
func BenchmarkAcceptsJSON(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")
	ctx := &Context{Request: req}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.AcceptsJSON()
	}
}
