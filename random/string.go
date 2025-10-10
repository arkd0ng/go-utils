package random

import (
	"crypto/rand"
	"math/big"
)

// Character sets for random string generation
// 랜덤 문자열 생성을 위한 문자 집합
const (
	charsetAlpha          = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	charsetAlphaUpper     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetAlphaLower     = "abcdefghijklmnopqrstuvwxyz"
	charsetDigits         = "0123456789"
	charsetHex            = "0123456789ABCDEF"
	charsetHexLower       = "0123456789abcdef"
	charsetSpecial        = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	charsetSpecialLimited = "!@#$%^&*-_"
	charsetBase64         = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	charsetBase64URL      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

// stringGenerator provides methods for generating random strings
// stringGenerator는 랜덤 문자열 생성 메서드를 제공합니다
type stringGenerator struct{}

// GenString is a global instance for generating random strings
// GenString은 랜덤 문자열 생성을 위한 전역 인스턴스입니다
var GenString = stringGenerator{}

// Letters generates a random string containing only alphabetic characters (a-z, A-Z)
// Letters는 알파벳 문자(a-z, A-Z)만 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
func (stringGenerator) Letters(min, max int) string {
	return generateRandomString(charsetAlpha, min, max)
}

// Alnum generates a random string containing alphanumeric characters (a-z, A-Z, 0-9)
// Alnum은 영숫자 문자(a-z, A-Z, 0-9)를 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
func (stringGenerator) Alnum(min, max int) string {
	return generateRandomString(charsetAlpha+charsetDigits, min, max)
}

// Complex generates a random string containing alphanumeric and all special characters
// Complex는 영숫자와 모든 특수 문자를 포함하는 랜덤 문자열을 생성합니다
//
// Special characters include: !@#$%^&*()-_=+[]{}|;:,.<>?/
// 특수 문자 포함: !@#$%^&*()-_=+[]{}|;:,.<>?/
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
func (stringGenerator) Complex(min, max int) string {
	return generateRandomString(charsetAlpha+charsetDigits+charsetSpecial, min, max)
}

// Standard generates a random string with alphanumeric and safe special characters
// Standard는 영숫자와 안전한 특수 문자를 포함하는 랜덤 문자열을 생성합니다
//
// Special characters include only: !@#$%^&*-_
// 특수 문자는 다음만 포함: !@#$%^&*-_
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
func (stringGenerator) Standard(min, max int) string {
	return generateRandomString(charsetAlpha+charsetDigits+charsetSpecialLimited, min, max)
}

// Digits generates a random string containing only numeric digits (0-9)
// Digits는 숫자(0-9)만 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
//
// Common use cases / 일반적인 사용 사례:
//   - PIN codes / PIN 코드
//   - Verification codes / 인증 코드
//   - OTP (One-Time Password) / 일회용 비밀번호
func (stringGenerator) Digits(min, max int) string {
	return generateRandomString(charsetDigits, min, max)
}

// Hex generates a random hexadecimal string with uppercase letters (0-9, A-F)
// Hex는 대문자 16진수 문자열(0-9, A-F)을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
//
// Common use cases / 일반적인 사용 사례:
//   - Color codes / 색상 코드
//   - Hash values / 해시값
//   - Hexadecimal identifiers / 16진수 식별자
func (stringGenerator) Hex(min, max int) string {
	return generateRandomString(charsetHex, min, max)
}

// HexLower generates a random hexadecimal string with lowercase letters (0-9, a-f)
// HexLower는 소문자 16진수 문자열(0-9, a-f)을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
//
// Common use cases / 일반적인 사용 사례:
//   - UUID generation / UUID 생성
//   - Hash values / 해시값
//   - Lowercase hexadecimal identifiers / 소문자 16진수 식별자
func (stringGenerator) HexLower(min, max int) string {
	return generateRandomString(charsetHexLower, min, max)
}

// AlphaUpper generates a random string containing only uppercase letters (A-Z)
// AlphaUpper는 대문자 알파벳(A-Z)만 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
//
// Common use cases / 일반적인 사용 사례:
//   - Ticket codes / 티켓 코드
//   - Coupon codes / 쿠폰 코드
//   - Uppercase identifiers / 대문자 식별자
func (stringGenerator) AlphaUpper(min, max int) string {
	return generateRandomString(charsetAlphaUpper, min, max)
}

// AlphaLower generates a random string containing only lowercase letters (a-z)
// AlphaLower는 소문자 알파벳(a-z)만 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
//
// Common use cases / 일반적인 사용 사례:
//   - Usernames / 사용자명
//   - Subdomain names / 서브도메인 이름
//   - Lowercase identifiers / 소문자 식별자
func (stringGenerator) AlphaLower(min, max int) string {
	return generateRandomString(charsetAlphaLower, min, max)
}

// AlnumUpper generates a random string with uppercase letters and digits (A-Z, 0-9)
// AlnumUpper는 대문자 알파벳과 숫자(A-Z, 0-9)를 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
//
// Common use cases / 일반적인 사용 사례:
//   - License keys / 라이선스 키
//   - Product codes / 제품 코드
//   - Uppercase alphanumeric codes / 대문자 영숫자 코드
func (stringGenerator) AlnumUpper(min, max int) string {
	return generateRandomString(charsetAlphaUpper+charsetDigits, min, max)
}

// AlnumLower generates a random string with lowercase letters and digits (a-z, 0-9)
// AlnumLower는 소문자 알파벳과 숫자(a-z, 0-9)를 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
//
// Common use cases / 일반적인 사용 사례:
//   - Tokens / 토큰
//   - Identifiers / 식별자
//   - Lowercase alphanumeric codes / 소문자 영숫자 코드
func (stringGenerator) AlnumLower(min, max int) string {
	return generateRandomString(charsetAlphaLower+charsetDigits, min, max)
}

// Base64 generates a random string using Base64 character set (A-Z, a-z, 0-9, +, /)
// Base64는 Base64 문자 집합(A-Z, a-z, 0-9, +, /)을 사용하여 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
//
// Common use cases / 일반적인 사용 사례:
//   - Base64-like encoding / Base64 형태 인코딩
//   - Data representation / 데이터 표현
func (stringGenerator) Base64(min, max int) string {
	return generateRandomString(charsetBase64, min, max)
}

// Base64URL generates a random string using URL-safe Base64 character set (A-Z, a-z, 0-9, -, _)
// Base64URL은 URL-safe Base64 문자 집합(A-Z, a-z, 0-9, -, _)을 사용하여 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
//
// Common use cases / 일반적인 사용 사례:
//   - URL-safe tokens / URL-safe 토큰
//   - File names / 파일명
//   - URL parameters / URL 매개변수
func (stringGenerator) Base64URL(min, max int) string {
	return generateRandomString(charsetBase64URL, min, max)
}

// Custom generates a random string using a custom character set
// Custom은 사용자 정의 문자 집합을 사용하여 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - charset: custom set of characters to use for generation / 생성에 사용할 사용자 정의 문자 집합
//   - min: minimum length of the generated string / 생성될 문자열의 최소 길이
//   - max: maximum length of the generated string / 생성될 문자열의 최대 길이
//
// Returns / 반환값:
//   - A random string with length between min and max (inclusive) / min과 max 사이 길이의 랜덤 문자열 (포함)
func (stringGenerator) Custom(charset string, min, max int) string {
	return generateRandomString(charset, min, max)
}

// generateRandomString is a helper function that generates a random string
// from the given charset with a length between min and max
// generateRandomString은 주어진 문자 집합에서 min과 max 사이의 길이로
// 랜덤 문자열을 생성하는 헬퍼 함수입니다
func generateRandomString(charset string, min, max int) string {
	if min < 0 {
		min = 0
	}
	if max < min {
		max = min
	}
	if len(charset) == 0 {
		return ""
	}

	// Determine the actual length of the string to generate
	// 생성할 문자열의 실제 길이 결정
	length := min
	if max > min {
		// Generate random length between min and max
		// min과 max 사이의 랜덤 길이 생성
		lengthRange := max - min + 1
		randomLength, err := rand.Int(rand.Reader, big.NewInt(int64(lengthRange)))
		if err == nil {
			length = min + int(randomLength.Int64())
		}
	}

	// Generate the random string
	// 랜덤 문자열 생성
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			// Fallback to first character if random generation fails
			// 랜덤 생성 실패 시 첫 번째 문자로 대체
			result[i] = charset[0]
			continue
		}
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result)
}
