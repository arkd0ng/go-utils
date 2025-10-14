package stringutil

import (
	"regexp"
	"strings"
	"unicode"
)

// IsEmail validates if a string is an email address (practical validation).
// IsEmail은 문자열이 이메일 주소인지 검증합니다 (실용적 검증).
//
// Not RFC 5322 compliant, but good enough for 99% of cases.
// RFC 5322 완전 준수 아니지만 99%의 경우에 충분함.
//
// Example:
//
//	IsEmail("user@example.com")      // true
//	IsEmail("user+tag@example.com")  // true
//	IsEmail("invalid.email")         // false
//	IsEmail("@example.com")          // false
func IsEmail(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(s)
}

// IsURL validates if a string is a URL.
// IsURL은 문자열이 URL인지 검증합니다.
//
// Checks for http:// or https:// scheme.
// http:// 또는 https:// 스킴 확인.
//
// Example:
//
//	IsURL("https://example.com")       // true
//	IsURL("http://example.com/path")   // true
//	IsURL("example.com")               // false (no scheme)
//	IsURL("htp://invalid")             // false
func IsURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

// IsAlphanumeric checks if string contains only alphanumeric characters (a-z, A-Z, 0-9).
// IsAlphanumeric은 문자열이 영숫자만 포함하는지 확인합니다 (a-z, A-Z, 0-9).
//
// Example:
//
//	IsAlphanumeric("abc123")   // true
//	IsAlphanumeric("ABC")      // true
//	IsAlphanumeric("abc-123")  // false (has dash)
//	IsAlphanumeric("abc 123")  // false (has space)
func IsAlphanumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsNumeric checks if string contains only digits (0-9).
// IsNumeric은 문자열이 숫자만 포함하는지 확인합니다 (0-9).
//
// Example:
//
//	IsNumeric("12345")   // true
//	IsNumeric("0")       // true
//	IsNumeric("123.45")  // false (has dot)
//	IsNumeric("-123")    // false (has minus)
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsAlpha checks if string contains only letters (a-z, A-Z).
// IsAlpha는 문자열이 알파벳만 포함하는지 확인합니다 (a-z, A-Z).
//
// Example:
//
//	IsAlpha("abcABC")  // true
//	IsAlpha("hello")   // true
//	IsAlpha("abc123")  // false (has digits)
//	IsAlpha("abc-")    // false (has dash)
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsBlank checks if string is empty or contains only whitespace.
// IsBlank는 문자열이 비어있거나 공백만 포함하는지 확인합니다.
//
// Example:
//
//	IsBlank("")       // true
//	IsBlank("   ")    // true
//	IsBlank("\t\n")   // true
//	IsBlank("hello")  // false
//	IsBlank(" a ")    // false
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsLower checks if all letters in string are lowercase.
// IsLower는 문자열의 모든 글자가 소문자인지 확인합니다.
//
// Example:
//
//	IsLower("hello")   // true
//	IsLower("abc")     // true
//	IsLower("Hello")   // false
//	IsLower("ABC")     // false
//	IsLower("abc123")  // true (digits don't affect)
func IsLower(s string) bool {
	if s == "" {
		return false
	}
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsLower(r) {
				return false
			}
		}
	}
	return hasLetter
}

// IsUpper checks if all letters in string are uppercase.
// IsUpper는 문자열의 모든 글자가 대문자인지 확인합니다.
//
// Example:
//
//	IsUpper("HELLO")   // true
//	IsUpper("ABC")     // true
//	IsUpper("Hello")   // false
//	IsUpper("abc")     // false
//	IsUpper("ABC123")  // true (digits don't affect)
func IsUpper(s string) bool {
	if s == "" {
		return false
	}
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetter
}
