package goutils

import (
	"strings"
	"testing"
	"unicode"
)

// TestAlpha tests the Alpha method
func TestAlpha(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"Fixed length", 10, 10},
		{"Variable length", 5, 15},
		{"Short length", 1, 5},
		{"Long length", 50, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenRandomString.Alpha(tt.min, tt.max)

			// Check length
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("Alpha() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are alphabetic
			for _, char := range result {
				if !unicode.IsLetter(char) {
					t.Errorf("Alpha() contains non-alphabetic character: %c", char)
				}
			}
		})
	}
}

// TestAlphaNum tests the AlphaNum method
func TestAlphaNum(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"Fixed length", 32, 32},
		{"Variable length", 32, 128},
		{"Short length", 8, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenRandomString.AlphaNum(tt.min, tt.max)

			// Check length
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("AlphaNum() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are alphanumeric
			for _, char := range result {
				if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
					t.Errorf("AlphaNum() contains invalid character: %c", char)
				}
			}
		})
	}
}

// TestAlphaNumSpecial tests the AlphaNumSpecial method
func TestAlphaNumSpecial(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"Fixed length", 20, 20},
		{"Variable length", 10, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenRandomString.AlphaNumSpecial(tt.min, tt.max)

			// Check length
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("AlphaNumSpecial() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are from the expected charset
			validChars := charsetAlpha + charsetDigits + charsetSpecial
			for _, char := range result {
				if !strings.ContainsRune(validChars, char) {
					t.Errorf("AlphaNumSpecial() contains unexpected character: %c", char)
				}
			}
		})
	}
}

// TestAlphaNumSpecialLimited tests the AlphaNumSpecialLimited method
func TestAlphaNumSpecialLimited(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"Fixed length", 15, 15},
		{"Variable length", 8, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenRandomString.AlphaNumSpecialLimited(tt.min, tt.max)

			// Check length
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("AlphaNumSpecialLimited() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are from the expected charset
			validChars := charsetAlpha + charsetDigits + charsetSpecialLimited
			for _, char := range result {
				if !strings.ContainsRune(validChars, char) {
					t.Errorf("AlphaNumSpecialLimited() contains unexpected character: %c", char)
				}
			}
		})
	}
}

// TestCustom tests the Custom method
func TestCustom(t *testing.T) {
	tests := []struct {
		name    string
		charset string
		min     int
		max     int
	}{
		{"Numbers only", "0123456789", 5, 10},
		{"Custom chars", "ABC123", 8, 12},
		{"Single char", "X", 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenRandomString.Custom(tt.charset, tt.min, tt.max)

			// Check length
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("Custom() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are from the custom charset
			for _, char := range result {
				if !strings.ContainsRune(tt.charset, char) {
					t.Errorf("Custom() contains unexpected character: %c, charset: %s", char, tt.charset)
				}
			}
		})
	}
}

// TestEdgeCases tests edge cases
func TestEdgeCases(t *testing.T) {
	t.Run("Min greater than max", func(t *testing.T) {
		result := GenRandomString.Alpha(10, 5)
		if len(result) != 10 {
			t.Errorf("When min > max, length should be min (10), got %d", len(result))
		}
	})

	t.Run("Negative min", func(t *testing.T) {
		result := GenRandomString.Alpha(-5, 10)
		if len(result) < 0 || len(result) > 10 {
			t.Errorf("With negative min, length should be between 0 and max")
		}
	})

	t.Run("Empty charset", func(t *testing.T) {
		result := GenRandomString.Custom("", 5, 10)
		if result != "" {
			t.Errorf("Empty charset should return empty string, got %s", result)
		}
	})

	t.Run("Zero length", func(t *testing.T) {
		result := GenRandomString.Alpha(0, 0)
		if len(result) != 0 {
			t.Errorf("Zero length should return empty string, got %s with length %d", result, len(result))
		}
	})
}

// TestRandomness tests that generated strings are actually random
func TestRandomness(t *testing.T) {
	// Generate multiple strings and check they're not all the same
	results := make(map[string]bool)
	iterations := 100

	for i := 0; i < iterations; i++ {
		result := GenRandomString.AlphaNum(10, 20)
		results[result] = true
	}

	// We expect at least some variation in 100 random strings
	// If we get less than 50 unique strings, something might be wrong
	if len(results) < iterations/2 {
		t.Errorf("Randomness test failed: only %d unique strings out of %d iterations", len(results), iterations)
	}
}

// BenchmarkAlpha benchmarks the Alpha method
func BenchmarkAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenRandomString.Alpha(10, 20)
	}
}

// BenchmarkAlphaNum benchmarks the AlphaNum method
func BenchmarkAlphaNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenRandomString.AlphaNum(32, 128)
	}
}

// BenchmarkCustom benchmarks the Custom method
func BenchmarkCustom(b *testing.B) {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < b.N; i++ {
		GenRandomString.Custom(charset, 16, 32)
	}
}
