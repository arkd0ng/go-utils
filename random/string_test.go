package random

import (
	"math"
	"strings"
	"testing"
	"unicode"
)

// TestLetters tests the Letters method
// TestLetters는 Letters 메서드를 테스트합니다
func TestLetters(t *testing.T) {
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
			result := GenString.Letters(tt.min, tt.max)

			// Check length / 길이 확인
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("Letters() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are alphabetic
			// 모든 문자가 알파벳인지 확인
			for _, char := range result {
				if !unicode.IsLetter(char) {
					t.Errorf("Letters() contains non-alphabetic character: %c", char)
				}
			}
		})
	}
}

// TestAlnum tests the Alnum method
// TestAlnum은 Alnum 메서드를 테스트합니다
func TestAlnum(t *testing.T) {
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
			result := GenString.Alnum(tt.min, tt.max)

			// Check length / 길이 확인
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("Alnum() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are alphanumeric
			// 모든 문자가 영숫자인지 확인
			for _, char := range result {
				if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
					t.Errorf("Alnum() contains invalid character: %c", char)
				}
			}
		})
	}
}

// TestComplex tests the Complex method
// TestComplex는 Complex 메서드를 테스트합니다
func TestComplex(t *testing.T) {
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
			result := GenString.Complex(tt.min, tt.max)

			// Check length / 길이 확인
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("Complex() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are from the expected charset
			// 모든 문자가 예상된 문자 집합에 포함되는지 확인
			validChars := charsetAlpha + charsetDigits + charsetSpecial
			for _, char := range result {
				if !strings.ContainsRune(validChars, char) {
					t.Errorf("Complex() contains unexpected character: %c", char)
				}
			}
		})
	}
}

// TestStandard tests the Standard method
// TestStandard는 Standard 메서드를 테스트합니다
func TestStandard(t *testing.T) {
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
			result := GenString.Standard(tt.min, tt.max)

			// Check length / 길이 확인
			if len(result) < tt.min || len(result) > tt.max {
				t.Errorf("Standard() length = %d, want between %d and %d", len(result), tt.min, tt.max)
			}

			// Check that all characters are from the expected charset
			// 모든 문자가 예상된 문자 집합에 포함되는지 확인
			validChars := charsetAlpha + charsetDigits + charsetSpecialLimited
			for _, char := range result {
				if !strings.ContainsRune(validChars, char) {
					t.Errorf("Standard() contains unexpected character: %c", char)
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
		result := GenString.Letters(10, 5)
		if len(result) != 10 {
			t.Errorf("When min > max, length should be min (10), got %d", len(result))
		}
	})

	t.Run("Negative min", func(t *testing.T) {
		result := GenString.Letters(-5, 10)
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
		result := GenString.Letters(0, 0)
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
		result := GenString.Alnum(10, 20)
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

// TestCollisionProbability tests collision probability with large sample sizes
// TestCollisionProbability는 대량 샘플에서 충돌 확률을 테스트합니다
func TestCollisionProbability(t *testing.T) {
	tests := []struct {
		name       string
		length     int
		iterations int
		charset    string
		method     func(int, int) string
	}{
		{
			name:       "10-char Alnum (10,000 iterations)",
			length:     10,
			iterations: 10000,
			charset:    charsetAlpha + charsetDigits,
			method:     GenString.Alnum,
		},
		{
			name:       "12-char Alnum (50,000 iterations)",
			length:     12,
			iterations: 50000,
			charset:    charsetAlpha + charsetDigits,
			method:     GenString.Alnum,
		},
		{
			name:       "8-char Letters (10,000 iterations)",
			length:     8,
			iterations: 10000,
			charset:    charsetAlpha,
			method:     GenString.Letters,
		},
		{
			name:       "16-char Complex (10,000 iterations)",
			length:     16,
			iterations: 10000,
			charset:    charsetAlpha + charsetDigits + charsetSpecial,
			method:     GenString.Complex,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := make(map[string]bool)
			collisions := 0

			// Generate strings and track collisions
			// 문자열을 생성하고 충돌을 추적합니다
			for i := 0; i < tt.iterations; i++ {
				str := tt.method(tt.length, tt.length)

				if results[str] {
					collisions++
				}
				results[str] = true
			}

			// Calculate theoretical collision probability using Birthday Paradox
			// Birthday Paradox를 사용하여 이론적 충돌 확률을 계산합니다
			charsetSize := float64(len(tt.charset))
			possibleStrings := math.Pow(charsetSize, float64(tt.length))

			// Approximate collision probability: 1 - exp(-n^2 / 2N)
			// where n = iterations, N = possible strings
			// 근사 충돌 확률: 1 - exp(-n^2 / 2N)
			// n = 반복 횟수, N = 가능한 문자열 수
			n := float64(tt.iterations)
			theoreticalProb := 1 - math.Exp(-(n*n)/(2*possibleStrings))
			expectedCollisions := theoreticalProb * n

			t.Logf("Charset size / 문자 집합 크기: %d", len(tt.charset))
			t.Logf("String length / 문자열 길이: %d", tt.length)
			t.Logf("Possible unique strings / 가능한 고유 문자열 수: %.2e", possibleStrings)
			t.Logf("Iterations / 반복 횟수: %d", tt.iterations)
			t.Logf("Unique strings generated / 생성된 고유 문자열: %d", len(results))
			t.Logf("Collisions found / 발견된 충돌: %d", collisions)
			t.Logf("Theoretical collision probability / 이론적 충돌 확률: %.6f%%", theoreticalProb*100)
			t.Logf("Expected collisions / 예상 충돌 횟수: %.2f", expectedCollisions)
			t.Logf("Actual collision rate / 실제 충돌률: %.6f%%", float64(collisions)/float64(tt.iterations)*100)

			// For cryptographically secure random generation with reasonable parameters,
			// we expect very few or no collisions
			// 암호학적으로 안전한 랜덤 생성과 적절한 매개변수로,
			// 매우 적거나 충돌이 없을 것으로 예상합니다

			// Allow up to 2x the theoretical expected collisions as threshold
			// 이론적 예상 충돌의 2배까지 허용
			maxAllowedCollisions := int(expectedCollisions*2) + 1

			if collisions > maxAllowedCollisions {
				t.Errorf("Too many collisions: got %d, expected around %.2f (max allowed %d)",
					collisions, expectedCollisions, maxAllowedCollisions)
			}

			// Verify uniqueness percentage is high
			// 고유성 비율이 높은지 확인
			uniquenessRate := float64(len(results)) / float64(tt.iterations) * 100
			t.Logf("Uniqueness rate / 고유성 비율: %.2f%%", uniquenessRate)

			if uniquenessRate < 99.0 {
				t.Errorf("Uniqueness rate too low: %.2f%%, expected > 99%%", uniquenessRate)
			}
		})
	}
}

// BenchmarkLetters benchmarks the Letters method
// BenchmarkLetters는 Letters 메서드의 성능을 벤치마크합니다
func BenchmarkLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenString.Letters(10, 20)
	}
}

// BenchmarkAlnum benchmarks the Alnum method
// BenchmarkAlnum은 Alnum 메서드의 성능을 벤치마크합니다
func BenchmarkAlnum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenString.Alnum(32, 128)
	}
}

// BenchmarkComplex benchmarks the Complex method
// BenchmarkComplex는 Complex 메서드의 성능을 벤치마크합니다
func BenchmarkComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenString.Complex(16, 24)
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
