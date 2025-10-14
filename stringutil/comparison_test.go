package stringutil

import "testing"

func TestEqualFold(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{"same case", "hello", "hello", true},
		{"different case", "hello", "HELLO", true},
		{"mixed case", "GoLang", "golang", true},
		{"not equal", "hello", "world", false},
		{"empty strings", "", "", true},
		{"one empty", "hello", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EqualFold(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("EqualFold(%q, %q) = %v, want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		prefix   string
		expected bool
	}{
		{"has prefix", "hello world", "hello", true},
		{"has prefix go", "golang", "go", true},
		{"no prefix", "hello", "world", false},
		{"empty prefix", "hello", "", true},
		{"empty string", "", "hello", false},
		{"exact match", "hello", "hello", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasPrefix(tt.s, tt.prefix)
			if result != tt.expected {
				t.Errorf("HasPrefix(%q, %q) = %v, want %v", tt.s, tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		suffix   string
		expected bool
	}{
		{"has suffix", "hello world", "world", true},
		{"has suffix lang", "golang", "lang", true},
		{"no suffix", "hello", "world", false},
		{"empty suffix", "hello", "", true},
		{"empty string", "", "hello", false},
		{"exact match", "hello", "hello", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasSuffix(tt.s, tt.suffix)
			if result != tt.expected {
				t.Errorf("HasSuffix(%q, %q) = %v, want %v", tt.s, tt.suffix, result, tt.expected)
			}
		})
	}
}
