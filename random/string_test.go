package random

import (
	"strings"
	"testing"
	"unicode"
)

// TestAlpha tests the Alpha method
// TestAlpha는 Alpha 메서드를 테스트합니다
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
			result := GenString.Alpha(tt.min, tt.max)

			// Check length / 길이 확인
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("Alpha() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are alphabetic
			// 모든 문자가 알파벳인지 확인
			for _, char := range result {
				if !unicode.IsLetter(char) {
					t.Errorf("Alpha() contains non-alphabetic character: %c", char)
				}
			}
		})
	}
}

// TestAlphaNum tests the AlphaNum method
// TestAlphaNum은 AlphaNum 메서드를 테스트합니다
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
			result := GenString.AlphaNum(tt.min, tt.max)

			// Check length / 길이 확인
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("AlphaNum() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are alphanumeric
			// 모든 문자가 영숫자인지 확인
			for _, char := range result {
				if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
					t.Errorf("AlphaNum() contains invalid character: %c", char)
				}
			}
		})
	}
}

// TestAlphaNumSpecial tests the AlphaNumSpecial method
// TestAlphaNumSpecial은 AlphaNumSpecial 메서드를 테스트합니다
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
			result := GenString.AlphaNumSpecial(tt.min, tt.max)

			// Check length / 길이 확인
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("AlphaNumSpecial() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are from the expected charset
			// 모든 문자가 예상된 문자 집합에 포함되는지 확인
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
// TestAlphaNumSpecialLimited는 AlphaNumSpecialLimited 메서드를 테스트합니다
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
			result := GenString.AlphaNumSpecialLimited(tt.min, tt.max)

			// Check length / 길이 확인
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("AlphaNumSpecialLimited() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are from the expected charset
			// 모든 문자가 예상된 문자 집합에 포함되는지 확인
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
// TestCustom은 Custom 메서드를 테스트합니다
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
			result := GenString.Custom(tt.charset, tt.min, tt.max)

			// Check length / 길이 확인
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("Custom() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are from the custom charset
			// 모든 문자가 사용자 정의 문자 집합에 포함되는지 확인
			for _, char := range result {
				if !strings.ContainsRune(tt.charset, char) {
					t.Errorf("Custom() contains unexpected character: %c, charset: %s", char, tt.charset)
				}
			}
		})
	}
}

// TestEdgeCases tests edge cases
// TestEdgeCases는 엣지 케이스를 테스트합니다
func TestEdgeCases(t *testing.T) {
	t.Run("Min greater than max", func(t *testing.T) {
		result := GenString.Alpha(10, 5)
		if len(result) != 10 {
			t.Errorf("When min > max, length should be min (10), got %d", len(result))
		}
	})

	t.Run("Negative min", func(t *testing.T) {
		result := GenString.Alpha(-5, 10)
		if len(result) < 0 || len(result) > 10 {
			t.Errorf("With negative min, length should be between 0 and max")
		}
	})

	t.Run("Empty charset", func(t *testing.T) {
		result := GenString.Custom("", 5, 10)
		if result != "" {
			t.Errorf("Empty charset should return empty string, got %s", result)
		}
	})

	t.Run("Zero length", func(t *testing.T) {
		result := GenString.Alpha(0, 0)
		if len(result) != 0 {
			t.Errorf("Zero length should return empty string, got %s with length %d", result, len(result))
		}
	})
}

// TestRandomness tests that generated strings are actually random
// TestRandomness는 생성된 문자열이 실제로 랜덤인지 테스트합니다
func TestRandomness(t *testing.T) {
	// Generate multiple strings and check they're not all the same
	// 여러 문자열을 생성하고 모두 같지 않은지 확인
	results := make(map[string]bool)
	iterations := 100

	for i := 0; i < iterations; i++ {
		result := GenString.AlphaNum(10, 20)
		results[result] = true
	}

	// We expect at least some variation in 100 random strings
	// If we get less than 50 unique strings, something might be wrong
	// 100개의 랜덤 문자열에서 최소한의 변화를 기대합니다
	// 50개 미만의 고유 문자열이 나오면 문제가 있을 수 있습니다
	if len(results) < iterations/2 {
		t.Errorf("Randomness test failed: only %d unique strings out of %d iterations", len(results), iterations)
	}
}

// BenchmarkAlpha benchmarks the Alpha method
// BenchmarkAlpha는 Alpha 메서드의 성능을 벤치마크합니다
func BenchmarkAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenString.Alpha(10, 20)
	}
}

// BenchmarkAlphaNum benchmarks the AlphaNum method
// BenchmarkAlphaNum은 AlphaNum 메서드의 성능을 벤치마크합니다
func BenchmarkAlphaNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenString.AlphaNum(32, 128)
	}
}

// BenchmarkCustom benchmarks the Custom method
// BenchmarkCustom은 Custom 메서드의 성능을 벤치마크합니다
func BenchmarkCustom(b *testing.B) {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < b.N; i++ {
		GenString.Custom(charset, 16, 32)
	}
}
