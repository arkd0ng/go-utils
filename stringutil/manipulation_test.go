package stringutil

import "testing"

func TestTruncate(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{"Hello World", 8, "Hello Wo..."},
		{"Hello", 10, "Hello"},
		{"안녕하세요", 3, "안녕하..."},
	}

	for _, tt := range tests {
		result := Truncate(tt.input, tt.length)
		if result != tt.expected {
			t.Errorf("Truncate(%q, %d) = %q; want %q", tt.input, tt.length, result, tt.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"안녕", "녕안"},
		{"", ""},
	}

	for _, tt := range tests {
		result := Reverse(tt.input)
		if result != tt.expected {
			t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  hello   world  ", "hello world"},
		{"\t\nhello\t\nworld", "hello world"},
		{"", ""},
	}

	for _, tt := range tests {
		result := Clean(tt.input)
		if result != tt.expected {
			t.Errorf("Clean(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
