package stringutil

import "testing"

func TestRepeat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		count    int
		expected string
	}{
		{"basic repeat", "abc", 3, "abcabcabc"},
		{"single repeat", "hello", 1, "hello"},
		{"zero repeat", "test", 0, ""},
		{"negative count", "test", -1, ""},
		{"unicode repeat", "안녕", 2, "안녕안녕"},
		{"emoji repeat", "🎉", 3, "🎉🎉🎉"},
		{"empty string", "", 5, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Repeat(tt.input, tt.count)
			if result != tt.expected {
				t.Errorf("Repeat(%q, %d) = %q, want %q", tt.input, tt.count, result, tt.expected)
			}
		})
	}
}

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
