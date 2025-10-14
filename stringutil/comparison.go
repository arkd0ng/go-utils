package stringutil

import "strings"

// EqualFold compares two strings case-insensitively.
// EqualFold는 두 문자열을 대소문자 구분 없이 비교합니다.
//
// Returns true if strings are equal ignoring case.
// 대소문자를 무시하고 문자열이 같으면 true를 반환합니다.
//
// Example:
//
//	EqualFold("hello", "HELLO")  // true
//	EqualFold("GoLang", "golang") // true
//	EqualFold("hello", "world")  // false
func EqualFold(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

// HasPrefix checks if string starts with the given prefix.
// HasPrefix는 문자열이 주어진 접두사로 시작하는지 확인합니다.
//
// Example:
//
//	HasPrefix("hello world", "hello")  // true
//	HasPrefix("golang", "go")          // true
//	HasPrefix("hello", "world")        // false
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix checks if string ends with the given suffix.
// HasSuffix는 문자열이 주어진 접미사로 끝나는지 확인합니다.
//
// Example:
//
//	HasSuffix("hello world", "world")  // true
//	HasSuffix("golang", "lang")        // true
//	HasSuffix("hello", "world")        // false
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
