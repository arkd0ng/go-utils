package stringutil

import (
	"strings"
)

// CountWords counts the number of words in a string (split by whitespace).
// CountWords는 문자열의 단어 수를 셉니다 (공백으로 분리).
//
// Example:
//
//	CountWords("hello world")  // 2
//	CountWords("  a  b  c  ")  // 3
func CountWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

// CountOccurrences counts the number of times a substring appears in a string.
// CountOccurrences는 부분 문자열이 문자열에 나타나는 횟수를 셉니다.
//
// Example:
//
//	CountOccurrences("hello hello", "hello")  // 2
//	CountOccurrences("abcabc", "abc")          // 2
func CountOccurrences(s, substr string) int {
	return strings.Count(s, substr)
}

// Join joins a slice of strings with a separator (wrapper for strings.Join).
// Join은 구분자로 문자열 슬라이스를 연결합니다 (strings.Join의 래퍼).
//
// Example:
//
//	Join([]string{"a", "b", "c"}, "-")  // "a-b-c"
//	Join([]string{"hello", "world"}, " ")  // "hello world"
func Join(strs []string, sep string) string {
	return strings.Join(strs, sep)
}

// Map applies a function to all strings in a slice.
// Map은 슬라이스의 모든 문자열에 함수를 적용합니다.
//
// Example:
//
//	Map([]string{"a", "b"}, strings.ToUpper)  // ["A", "B"]
//	Map([]string{"hello", "world"}, func(s string) string { return s + "!" })  // ["hello!", "world!"]
func Map(strs []string, fn func(string) string) []string {
	result := make([]string, len(strs))
	for i, s := range strs {
		result[i] = fn(s)
	}
	return result
}

// Filter filters strings by a predicate function.
// Filter는 조건 함수로 문자열을 필터링합니다.
//
// Example:
//
//	Filter([]string{"a", "ab", "abc"}, func(s string) bool { return len(s) > 2 })  // ["abc"]
//	Filter([]string{"hello", "world", "hi"}, func(s string) bool { return len(s) > 3 })  // ["hello", "world"]
func Filter(strs []string, fn func(string) bool) []string {
	result := make([]string, 0) // Initialize to empty slice, not nil / nil이 아닌 빈 슬라이스로 초기화
	for _, s := range strs {
		if fn(s) {
			result = append(result, s)
		}
	}
	return result
}

// PadLeft pads a string on the left side to reach the specified length.
// PadLeft는 지정된 길이에 도달하도록 문자열의 왼쪽에 패딩을 추가합니다.
//
// Example:
//
//	PadLeft("5", 3, "0")    // "005"
//	PadLeft("42", 5, "0")   // "00042"
func PadLeft(s string, length int, pad string) string {
	runes := []rune(s)
	if len(runes) >= length {
		return s
	}
	padCount := length - len(runes)
	padding := strings.Repeat(pad, padCount)
	return padding + s
}

// PadRight pads a string on the right side to reach the specified length.
// PadRight는 지정된 길이에 도달하도록 문자열의 오른쪽에 패딩을 추가합니다.
//
// Example:
//
//	PadRight("5", 3, "0")   // "500"
//	PadRight("42", 5, "0")  // "42000"
func PadRight(s string, length int, pad string) string {
	runes := []rune(s)
	if len(runes) >= length {
		return s
	}
	padCount := length - len(runes)
	padding := strings.Repeat(pad, padCount)
	return s + padding
}

// Lines splits a string by newlines.
// Lines는 줄바꿈으로 문자열을 분리합니다.
//
// Example:
//
//	Lines("line1\nline2\nline3")  // ["line1", "line2", "line3"]
//	Lines("a\nb\nc")               // ["a", "b", "c"]
func Lines(s string) []string {
	return strings.Split(s, "\n")
}

// Words splits a string by whitespace.
// Words는 공백으로 문자열을 분리합니다.
//
// Example:
//
//	Words("hello world foo")  // ["hello", "world", "foo"]
//	Words("  a  b  c  ")       // ["a", "b", "c"]
func Words(s string) []string {
	return strings.Fields(s)
}
