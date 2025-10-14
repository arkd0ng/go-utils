package stringutil

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"user@example.com", true},
		{"user+tag@example.com", true},
		{"invalid.email", false},
		{"@example.com", false},
		{"user@", false},
	}

	for _, tt := range tests {
		result := IsEmail(tt.input)
		if result != tt.expected {
			t.Errorf("IsEmail(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"https://example.com", true},
		{"http://example.com/path", true},
		{"example.com", false},
		{"htp://invalid", false},
	}

	for _, tt := range tests {
		result := IsURL(tt.input)
		if result != tt.expected {
			t.Errorf("IsURL(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsAlphanumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc123", true},
		{"ABC", true},
		{"abc-123", false},
		{"abc 123", false},
	}

	for _, tt := range tests {
		result := IsAlphanumeric(tt.input)
		if result != tt.expected {
			t.Errorf("IsAlphanumeric(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
