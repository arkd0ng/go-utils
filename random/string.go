package random

import (
	"crypto/rand"
	"errors"
	"fmt"
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
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//       Example: Letters(10) generates a 10-character string / Letters(10)은 10자 문자열 생성
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//       Example: Letters(8, 12) generates 8-12 character string / Letters(8, 12)는 8-12자 문자열 생성
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
func (stringGenerator) Letters(length ...int) (string, error) {
	return generateRandomString(charsetAlpha, length...)
}

// Alnum generates a random string containing alphanumeric characters (a-z, A-Z, 0-9)
// Alnum은 영숫자 문자(a-z, A-Z, 0-9)를 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
func (stringGenerator) Alnum(length ...int) (string, error) {
	return generateRandomString(charsetAlpha+charsetDigits, length...)
}

// Complex generates a random string containing alphanumeric and all special characters
// Complex는 영숫자와 모든 특수 문자를 포함하는 랜덤 문자열을 생성합니다
//
// Special characters include: !@#$%^&*()-_=+[]{}|;:,.<>?/
// 특수 문자 포함: !@#$%^&*()-_=+[]{}|;:,.<>?/
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
func (stringGenerator) Complex(length ...int) (string, error) {
	return generateRandomString(charsetAlpha+charsetDigits+charsetSpecial, length...)
}

// Standard generates a random string with alphanumeric and safe special characters
// Standard는 영숫자와 안전한 특수 문자를 포함하는 랜덤 문자열을 생성합니다
//
// Special characters include only: !@#$%^&*-_
// 특수 문자는 다음만 포함: !@#$%^&*-_
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
func (stringGenerator) Standard(length ...int) (string, error) {
	return generateRandomString(charsetAlpha+charsetDigits+charsetSpecialLimited, length...)
}

// Digits generates a random string containing only numeric digits (0-9)
// Digits는 숫자(0-9)만 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - PIN codes / PIN 코드
//   - Verification codes / 인증 코드
//   - OTP (One-Time Password) / 일회용 비밀번호
func (stringGenerator) Digits(length ...int) (string, error) {
	return generateRandomString(charsetDigits, length...)
}

// Hex generates a random hexadecimal string with uppercase letters (0-9, A-F)
// Hex는 대문자 16진수 문자열(0-9, A-F)을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - Color codes / 색상 코드
//   - Hash values / 해시값
//   - Hexadecimal identifiers / 16진수 식별자
func (stringGenerator) Hex(length ...int) (string, error) {
	return generateRandomString(charsetHex, length...)
}

// HexLower generates a random hexadecimal string with lowercase letters (0-9, a-f)
// HexLower는 소문자 16진수 문자열(0-9, a-f)을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - UUID generation / UUID 생성
//   - Hash values / 해시값
//   - Lowercase hexadecimal identifiers / 소문자 16진수 식별자
func (stringGenerator) HexLower(length ...int) (string, error) {
	return generateRandomString(charsetHexLower, length...)
}

// AlphaUpper generates a random string containing only uppercase letters (A-Z)
// AlphaUpper는 대문자 알파벳(A-Z)만 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - Ticket codes / 티켓 코드
//   - Coupon codes / 쿠폰 코드
//   - Uppercase identifiers / 대문자 식별자
func (stringGenerator) AlphaUpper(length ...int) (string, error) {
	return generateRandomString(charsetAlphaUpper, length...)
}

// AlphaLower generates a random string containing only lowercase letters (a-z)
// AlphaLower는 소문자 알파벳(a-z)만 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - Usernames / 사용자명
//   - Subdomain names / 서브도메인 이름
//   - Lowercase identifiers / 소문자 식별자
func (stringGenerator) AlphaLower(length ...int) (string, error) {
	return generateRandomString(charsetAlphaLower, length...)
}

// AlnumUpper generates a random string with uppercase letters and digits (A-Z, 0-9)
// AlnumUpper는 대문자 알파벳과 숫자(A-Z, 0-9)를 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - License keys / 라이선스 키
//   - Product codes / 제품 코드
//   - Uppercase alphanumeric codes / 대문자 영숫자 코드
func (stringGenerator) AlnumUpper(length ...int) (string, error) {
	return generateRandomString(charsetAlphaUpper+charsetDigits, length...)
}

// AlnumLower generates a random string with lowercase letters and digits (a-z, 0-9)
// AlnumLower는 소문자 알파벳과 숫자(a-z, 0-9)를 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - Tokens / 토큰
//   - Identifiers / 식별자
//   - Lowercase alphanumeric codes / 소문자 영숫자 코드
func (stringGenerator) AlnumLower(length ...int) (string, error) {
	return generateRandomString(charsetAlphaLower+charsetDigits, length...)
}

// Base64 generates a random string using Base64 character set (A-Z, a-z, 0-9, +, /)
// Base64는 Base64 문자 집합(A-Z, a-z, 0-9, +, /)을 사용하여 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - Base64-like encoding / Base64 형태 인코딩
//   - Data representation / 데이터 표현
func (stringGenerator) Base64(length ...int) (string, error) {
	return generateRandomString(charsetBase64, length...)
}

// Base64URL generates a random string using URL-safe Base64 character set (A-Z, a-z, 0-9, -, _)
// Base64URL은 URL-safe Base64 문자 집합(A-Z, a-z, 0-9, -, _)을 사용하여 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//       Example: Base64URL(64) generates a 64-character string / Base64URL(64)은 64자 문자열 생성
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//       Example: Base64URL(64, 128) generates 64-128 character string / Base64URL(64, 128)은 64-128자 문자열 생성
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - URL-safe tokens / URL-safe 토큰
//   - File names / 파일명
//   - URL parameters / URL 매개변수
func (stringGenerator) Base64URL(length ...int) (string, error) {
	return generateRandomString(charsetBase64URL, length...)
}

// Custom generates a random string using a custom character set
// Custom은 사용자 정의 문자 집합을 사용하여 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - charset: custom set of characters to use for generation / 생성에 사용할 사용자 정의 문자 집합
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
func (stringGenerator) Custom(charset string, length ...int) (string, error) {
	return generateRandomString(charset, length...)
}

// generateRandomString is a helper function that generates a random string
// from the given charset with a specified length
// generateRandomString은 주어진 문자 집합에서 지정된 길이로
// 랜덤 문자열을 생성하는 헬퍼 함수입니다
//
// Parameters / 매개변수:
//   - charset: character set to use / 사용할 문자 집합
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random string / 생성된 랜덤 문자열
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
func generateRandomString(charset string, length ...int) (string, error) {
	// Validate charset / 문자 집합 검증
	if len(charset) == 0 {
		return "", errors.New("charset cannot be empty")
	}

	// Parse length arguments / 길이 인자 파싱
	var min, max int
	switch len(length) {
	case 0:
		return "", errors.New("at least one length argument is required")
	case 1:
		// Fixed length / 고정 길이
		min = length[0]
		max = length[0]
	case 2:
		// Range length / 범위 길이
		min = length[0]
		max = length[1]
	default:
		return "", fmt.Errorf("invalid number of arguments: expected 1 or 2, got %d", len(length))
	}

	// Validate length parameters / 길이 매개변수 검증
	if min < 0 {
		return "", fmt.Errorf("minimum length cannot be negative: %d", min)
	}
	if max < min {
		return "", fmt.Errorf("maximum length (%d) cannot be less than minimum length (%d)", max, min)
	}

	// Determine the actual length of the string to generate
	// 생성할 문자열의 실제 길이 결정
	actualLength := min
	if max > min {
		// Generate random length between min and max
		// min과 max 사이의 랜덤 길이 생성
		lengthRange := max - min + 1
		randomLength, err := rand.Int(rand.Reader, big.NewInt(int64(lengthRange)))
		if err != nil {
			return "", fmt.Errorf("failed to generate random length: %w", err)
		}
		actualLength = min + int(randomLength.Int64())
	}

	// Generate the random string
	// 랜덤 문자열 생성
	result := make([]byte, actualLength)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < actualLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", fmt.Errorf("failed to generate random character at position %d: %w", i, err)
		}
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result), nil
}
