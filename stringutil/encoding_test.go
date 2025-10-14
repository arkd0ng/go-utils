package stringutil

import "testing"

// Test Base64 encoding/decoding / Base64 인코딩/디코딩 테스트
func TestBase64Encode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "aGVsbG8="},
		{"hello world", "aGVsbG8gd29ybGQ="},
		{"", ""},
		{"안녕하세요", "7JWI64WV7ZWY7IS47JqU"},
	}

	for _, tt := range tests {
		result := Base64Encode(tt.input)
		if result != tt.expected {
			t.Errorf("Base64Encode(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestBase64Decode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"aGVsbG8=", "hello", false},
		{"aGVsbG8gd29ybGQ=", "hello world", false},
		{"", "", false},
		{"7JWI64WV7ZWY7IS47JqU", "안녕하세요", false},
		{"invalid!", "", true},
	}

	for _, tt := range tests {
		result, err := Base64Decode(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("Base64Decode(%q) should return error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("Base64Decode(%q) unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Base64Decode(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		}
	}
}

// Test Base64 round-trip / Base64 왕복 테스트
func TestBase64RoundTrip(t *testing.T) {
	inputs := []string{
		"hello",
		"hello world",
		"안녕하세요",
		"hello世界",
		"🎉🎊",
		"The quick brown fox jumps over the lazy dog",
	}

	for _, input := range inputs {
		encoded := Base64Encode(input)
		decoded, err := Base64Decode(encoded)
		if err != nil {
			t.Errorf("Base64 round-trip error for %q: %v", input, err)
		}
		if decoded != input {
			t.Errorf("Base64 round-trip failed: input=%q, decoded=%q", input, decoded)
		}
	}
}

// Test Base64URL encoding/decoding / Base64URL 인코딩/디코딩 테스트
func TestBase64URLEncode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello?world", "aGVsbG8_d29ybGQ="},
		{"hello/world", "aGVsbG8vd29ybGQ="},
		{"", ""},
	}

	for _, tt := range tests {
		result := Base64URLEncode(tt.input)
		if result != tt.expected {
			t.Errorf("Base64URLEncode(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestBase64URLDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"aGVsbG8_d29ybGQ=", "hello?world", false},
		{"aGVsbG8vd29ybGQ=", "hello/world", false},
		{"", "", false},
		{"invalid!", "", true},
	}

	for _, tt := range tests {
		result, err := Base64URLDecode(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("Base64URLDecode(%q) should return error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("Base64URLDecode(%q) unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Base64URLDecode(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		}
	}
}

// Test URL encoding/decoding / URL 인코딩/디코딩 테스트
func TestURLEncode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello+world"},
		{"hello/world", "hello%2Fworld"},
		{"hello?world", "hello%3Fworld"},
		{"hello&world", "hello%26world"},
		{"", ""},
		{"안녕하세요", "%EC%95%88%EB%85%95%ED%95%98%EC%84%B8%EC%9A%94"},
	}

	for _, tt := range tests {
		result := URLEncode(tt.input)
		if result != tt.expected {
			t.Errorf("URLEncode(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestURLDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"hello+world", "hello world", false},
		{"hello%2Fworld", "hello/world", false},
		{"hello%3Fworld", "hello?world", false},
		{"hello%26world", "hello&world", false},
		{"", "", false},
		{"%EC%95%88%EB%85%95%ED%95%98%EC%84%B8%EC%9A%94", "안녕하세요", false},
		{"%ZZ", "", true},
	}

	for _, tt := range tests {
		result, err := URLDecode(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("URLDecode(%q) should return error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("URLDecode(%q) unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("URLDecode(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		}
	}
}

// Test URL round-trip / URL 왕복 테스트
func TestURLRoundTrip(t *testing.T) {
	inputs := []string{
		"hello world",
		"hello/world",
		"hello?world",
		"hello&world",
		"안녕하세요",
	}

	for _, input := range inputs {
		encoded := URLEncode(input)
		decoded, err := URLDecode(encoded)
		if err != nil {
			t.Errorf("URL round-trip error for %q: %v", input, err)
		}
		if decoded != input {
			t.Errorf("URL round-trip failed: input=%q, decoded=%q", input, decoded)
		}
	}
}

// Test HTML escape/unescape / HTML 이스케이프/언이스케이프 테스트
func TestHTMLEscape(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"<div>hello</div>", "&lt;div&gt;hello&lt;/div&gt;"},
		{"'hello' & \"world\"", "&#39;hello&#39; &amp; &#34;world&#34;"},
		{"<script>alert('XSS')</script>", "&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;"},
		{"", ""},
		{"no special chars", "no special chars"},
	}

	for _, tt := range tests {
		result := HTMLEscape(tt.input)
		if result != tt.expected {
			t.Errorf("HTMLEscape(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestHTMLUnescape(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"&lt;div&gt;hello&lt;/div&gt;", "<div>hello</div>"},
		{"&#39;hello&#39; &amp; &#34;world&#34;", "'hello' & \"world\""},
		{"&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;", "<script>alert('XSS')</script>"},
		{"", ""},
		{"no entities", "no entities"},
	}

	for _, tt := range tests {
		result := HTMLUnescape(tt.input)
		if result != tt.expected {
			t.Errorf("HTMLUnescape(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

// Test HTML round-trip / HTML 왕복 테스트
func TestHTMLRoundTrip(t *testing.T) {
	inputs := []string{
		"<div>hello</div>",
		"'hello' & \"world\"",
		"<script>alert('XSS')</script>",
	}

	for _, input := range inputs {
		escaped := HTMLEscape(input)
		unescaped := HTMLUnescape(escaped)
		if unescaped != input {
			t.Errorf("HTML round-trip failed: input=%q, unescaped=%q", input, unescaped)
		}
	}
}

// Benchmark encoding functions / 인코딩 함수 벤치마크
func BenchmarkBase64Encode(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	for i := 0; i < b.N; i++ {
		_ = Base64Encode(s)
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	s := Base64Encode("The quick brown fox jumps over the lazy dog")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Base64Decode(s)
	}
}

func BenchmarkURLEncode(b *testing.B) {
	s := "hello world & hello/world"
	for i := 0; i < b.N; i++ {
		_ = URLEncode(s)
	}
}

func BenchmarkURLDecode(b *testing.B) {
	s := URLEncode("hello world & hello/world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = URLDecode(s)
	}
}

func BenchmarkHTMLEscape(b *testing.B) {
	s := "<div>hello & 'world'</div>"
	for i := 0; i < b.N; i++ {
		_ = HTMLEscape(s)
	}
}

func BenchmarkHTMLUnescape(b *testing.B) {
	s := HTMLEscape("<div>hello & 'world'</div>")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = HTMLUnescape(s)
	}
}
