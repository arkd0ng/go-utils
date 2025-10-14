package stringutil

import (
	"regexp"
	"strings"
	"unicode"
)

// Truncate truncates a string to the specified length and appends "...".
// Truncate는 문자열을 지정된 길이로 자르고 "..."를 추가합니다.
//
// Unicode-safe: uses rune count, not byte count.
// 유니코드 안전: 바이트 수가 아닌 rune 수 사용.
//
// Example:
//
//	Truncate("Hello World", 8)    // "Hello..."
//	Truncate("안녕하세요", 3)        // "안녕하..."
func Truncate(s string, length int) string {
	return TruncateWithSuffix(s, length, "...")
}

// TruncateWithSuffix truncates a string with a custom suffix.
// TruncateWithSuffix는 사용자 정의 suffix로 문자열을 자릅니다.
//
// Example:
//
//	TruncateWithSuffix("Hello World", 8, "…")  // "Hello Wo…"
//	TruncateWithSuffix("안녕하세요", 3, "…")      // "안녕하…"
func TruncateWithSuffix(s string, length int, suffix string) string {
	runes := []rune(s)
	if len(runes) <= length {
		return s
	}
	return string(runes[:length]) + suffix
}

// Reverse reverses a string (Unicode-safe).
// Reverse는 문자열을 뒤집습니다 (유니코드 안전).
//
// Example:
//
//	Reverse("hello")  // "olleh"
//	Reverse("안녕")    // "녕안"
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Capitalize capitalizes the first letter of each word.
// Capitalize는 각 단어의 첫 글자를 대문자로 만듭니다.
//
// Example:
//
//	Capitalize("hello world")  // "Hello World"
//	Capitalize("hello-world")  // "Hello-World"
func Capitalize(s string) string {
	return strings.Title(s)
}

// CapitalizeFirst capitalizes only the first letter of the string.
// CapitalizeFirst는 문자열의 첫 글자만 대문자로 만듭니다.
//
// Example:
//
//	CapitalizeFirst("hello world")  // "Hello world"
//	CapitalizeFirst("HELLO WORLD")  // "HELLO WORLD"
func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// RemoveDuplicates removes duplicate characters from a string.
// RemoveDuplicates는 문자열에서 중복 문자를 제거합니다.
//
// Example:
//
//	RemoveDuplicates("hello")  // "helo"
//	RemoveDuplicates("aabbcc")  // "abc"
func RemoveDuplicates(s string) string {
	seen := make(map[rune]bool)
	var result []rune
	for _, r := range s {
		if !seen[r] {
			seen[r] = true
			result = append(result, r)
		}
	}
	return string(result)
}

// RemoveSpaces removes all whitespace from a string.
// RemoveSpaces는 문자열에서 모든 공백을 제거합니다.
//
// Example:
//
//	RemoveSpaces("h e l l o")  // "hello"
//	RemoveSpaces("  hello world  ")  // "helloworld"
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\t", ""), "\n", "")
}

// RemoveSpecialChars removes special characters, keeping only alphanumeric and spaces.
// RemoveSpecialChars는 특수 문자를 제거하고 영숫자와 공백만 유지합니다.
//
// Example:
//
//	RemoveSpecialChars("hello@#$123")  // "hello123"
//	RemoveSpecialChars("a!b@c#123")    // "abc123"
func RemoveSpecialChars(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	return re.ReplaceAllString(s, "")
}

// Clean trims whitespace and deduplicates spaces.
// Clean은 공백을 제거하고 중복 공백을 정리합니다.
//
// Example:
//
//	Clean("  hello   world  ")  // "hello world"
//	Clean("\t\nhello\t\nworld")  // "hello world"
func Clean(s string) string {
	// Trim leading/trailing spaces / 앞뒤 공백 제거
	s = strings.TrimSpace(s)

	// Replace multiple spaces with single space / 중복 공백을 단일 공백으로
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")

	return s
}

// Repeat repeats a string n times.
// Repeat는 문자열을 n번 반복합니다.
//
// Unicode-safe: works correctly with all Unicode characters.
// 유니코드 안전: 모든 유니코드 문자와 정상 작동.
//
// Example:
//
//	Repeat("hello", 3)  // "hellohellohello"
//	Repeat("안녕", 2)     // "안녕안녕"
//	Repeat("*", 5)      // "*****"
func Repeat(s string, count int) string {
	if count < 0 {
		return ""
	}
	return strings.Repeat(s, count)
}
