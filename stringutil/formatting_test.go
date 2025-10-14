package stringutil

import "testing"

// Test FormatNumber / FormatNumber 테스트
func TestFormatNumber(t *testing.T) {
	tests := []struct {
		n         int
		separator string
		expected  string
	}{
		{1000000, ",", "1,000,000"},
		{1234567, ".", "1.234.567"},
		{1234567, " ", "1 234 567"},
		{123, ",", "123"},
		{-1000000, ",", "-1,000,000"},
		{0, ",", "0"},
	}

	for _, tt := range tests {
		result := FormatNumber(tt.n, tt.separator)
		if result != tt.expected {
			t.Errorf("FormatNumber(%d, %q) = %q, want %q", tt.n, tt.separator, result, tt.expected)
		}
	}
}

// Test FormatBytes / FormatBytes 테스트
func TestFormatBytes(t *testing.T) {
	tests := []struct {
		bytes    int64
		expected string
	}{
		{0, "0 B"},
		{1023, "1023 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{1048576, "1.0 MB"},
		{1073741824, "1.0 GB"},
		{1099511627776, "1.0 TB"},
	}

	for _, tt := range tests {
		result := FormatBytes(tt.bytes)
		if result != tt.expected {
			t.Errorf("FormatBytes(%d) = %q, want %q", tt.bytes, result, tt.expected)
		}
	}
}

// Test Pluralize / Pluralize 테스트
func TestPluralize(t *testing.T) {
	tests := []struct {
		count    int
		singular string
		plural   string
		expected string
	}{
		{0, "item", "items", "items"},
		{1, "item", "items", "item"},
		{2, "item", "items", "items"},
		{5, "person", "people", "people"},
		{1, "person", "people", "person"},
	}

	for _, tt := range tests {
		result := Pluralize(tt.count, tt.singular, tt.plural)
		if result != tt.expected {
			t.Errorf("Pluralize(%d, %q, %q) = %q, want %q",
				tt.count, tt.singular, tt.plural, result, tt.expected)
		}
	}
}

// Test FormatWithCount / FormatWithCount 테스트
func TestFormatWithCount(t *testing.T) {
	tests := []struct {
		count    int
		singular string
		plural   string
		expected string
	}{
		{0, "item", "items", "0 items"},
		{1, "item", "items", "1 item"},
		{5, "item", "items", "5 items"},
	}

	for _, tt := range tests {
		result := FormatWithCount(tt.count, tt.singular, tt.plural)
		if result != tt.expected {
			t.Errorf("FormatWithCount(%d, %q, %q) = %q, want %q",
				tt.count, tt.singular, tt.plural, result, tt.expected)
		}
	}
}

// Test Ellipsis / Ellipsis 테스트
func TestEllipsis(t *testing.T) {
	tests := []struct {
		s        string
		maxLen   int
		expected string
	}{
		{"verylongfilename.txt", 15, "verylo...me.txt"},
		{"short.txt", 20, "short.txt"},
		{"abcdefgh", 5, "a...h"},
		{"abc", 5, "abc"},
		{"abcdefgh", 3, "abc"},
	}

	for _, tt := range tests {
		result := Ellipsis(tt.s, tt.maxLen)
		if result != tt.expected {
			t.Errorf("Ellipsis(%q, %d) = %q, want %q", tt.s, tt.maxLen, result, tt.expected)
		}
	}
}

// Test Mask / Mask 테스트
func TestMask(t *testing.T) {
	tests := []struct {
		s        string
		first    int
		last     int
		maskChar string
		expected string
	}{
		{"1234567890", 2, 2, "*", "12******90"},
		{"secret", 1, 1, "#", "s####t"},
		{"short", 2, 2, "*", "sh*rt"},  // 5 chars, first 2 + last 2 + 1 masked
		{"hi", 1, 1, "*", "hi"},        // too short to mask
	}

	for _, tt := range tests {
		result := Mask(tt.s, tt.first, tt.last, tt.maskChar)
		if result != tt.expected {
			t.Errorf("Mask(%q, %d, %d, %q) = %q, want %q",
				tt.s, tt.first, tt.last, tt.maskChar, result, tt.expected)
		}
	}
}

// Test MaskEmail / MaskEmail 테스트
func TestMaskEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected string
	}{
		{"john.doe@example.com", "j******e@example.com"},
		{"a@example.com", "a@example.com"},
		{"ab@example.com", "ab@example.com"},
		{"invalid", "invalid"},
	}

	for _, tt := range tests {
		result := MaskEmail(tt.email)
		if result != tt.expected {
			t.Errorf("MaskEmail(%q) = %q, want %q", tt.email, result, tt.expected)
		}
	}
}

// Test MaskCreditCard / MaskCreditCard 테스트
func TestMaskCreditCard(t *testing.T) {
	tests := []struct {
		card     string
		expected string
	}{
		{"1234567890123456", "************3456"},
		{"1234-5678-9012-3456", "****-****-****-3456"},
		{"1234", "1234"},
	}

	for _, tt := range tests {
		result := MaskCreditCard(tt.card)
		if result != tt.expected {
			t.Errorf("MaskCreditCard(%q) = %q, want %q", tt.card, result, tt.expected)
		}
	}
}

// Test AddLineNumbers / AddLineNumbers 테스트
func TestAddLineNumbers(t *testing.T) {
	input := "line1\nline2\nline3"
	expected := "1: line1\n2: line2\n3: line3"
	result := AddLineNumbers(input)
	if result != expected {
		t.Errorf("AddLineNumbers(%q) = %q, want %q", input, result, expected)
	}
}

// Test Indent / Indent 테스트
func TestIndent(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected string
	}{
		{"line1\nline2", "  ", "  line1\n  line2"},
		{"line1\nline2", "\t", "\tline1\n\tline2"},
		{"single", ">>", ">>single"},
	}

	for _, tt := range tests {
		result := Indent(tt.s, tt.prefix)
		if result != tt.expected {
			t.Errorf("Indent(%q, %q) = %q, want %q", tt.s, tt.prefix, result, tt.expected)
		}
	}
}

// Test Dedent / Dedent 테스트
func TestDedent(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"  line1\n  line2", "line1\nline2"},
		{"    line1\n  line2", "  line1\nline2"},
		{"line1\nline2", "line1\nline2"},
	}

	for _, tt := range tests {
		result := Dedent(tt.s)
		if result != tt.expected {
			t.Errorf("Dedent(%q) = %q, want %q", tt.s, result, tt.expected)
		}
	}
}

// Test WrapText / WrapText 테스트
func TestWrapText(t *testing.T) {
	input := "The quick brown fox jumps"
	expected := "The quick\nbrown fox\njumps"
	result := WrapText(input, 10)
	if result != expected {
		t.Errorf("WrapText(%q, 10) = %q, want %q", input, result, expected)
	}

	// Test with width 0
	result = WrapText(input, 0)
	if result != input {
		t.Errorf("WrapText(%q, 0) should return original string", input)
	}
}

// Benchmark formatting functions / 포맷팅 함수 벤치마크
func BenchmarkFormatNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FormatNumber(1234567890, ",")
	}
}

func BenchmarkFormatBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FormatBytes(1073741824)
	}
}

func BenchmarkMaskEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MaskEmail("john.doe@example.com")
	}
}

func BenchmarkWrapText(b *testing.B) {
	text := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WrapText(text, 20)
	}
}
