package stringutil

import (
	"regexp"
	"strings"
	"unicode"
)

// Truncate truncates a string to the specified length and appends "...". / Truncate는 문자열을 지정된 길이로 자르고 "..."를 추가합니다.
//
// Unicode-safe: uses rune count, not byte count.
// 유니코드 안전: 바이트 수가 아닌 rune 수 사용.
//
// Example:
//
// Truncate("Hello World", 8)    // "Hello..." / Truncate("안녕하세요", 3)        // "안녕하..."
func Truncate(s string, length int) string {
	return TruncateWithSuffix(s, length, "...")
}

// TruncateWithSuffix truncates a string with a custom suffix.
// TruncateWithSuffix는 사용자 정의 suffix로 문자열을 자릅니다.
//
// Example:
//
// TruncateWithSuffix("Hello World", 8, "…")  // "Hello Wo…" / TruncateWithSuffix("안녕하세요", 3, "…")      // "안녕하…"
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
// Reverse("hello")  // "olleh" / Reverse("안녕")    // "녕안"
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
	// Trim leading/trailing spaces
	// 앞뒤 공백 제거
	s = strings.TrimSpace(s)

	// Replace multiple spaces with single space
	// 중복 공백을 단일 공백으로
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
// Repeat("hello", 3)  // "hellohellohello" / Repeat("안녕", 2)     // "안녕안녕"
//	Repeat("*", 5)      // "*****"
func Repeat(s string, count int) string {
	if count < 0 {
		return ""
	}
	return strings.Repeat(s, count)
}

// Substring extracts a substring from start to end index (Unicode-safe).
// Substring은 start부터 end 인덱스까지 부분 문자열을 추출합니다 (유니코드 안전).
//
// Parameters:
// - start: starting index (inclusive)
// - end: ending index (exclusive)
//
// If indices are out of bounds, they are adjusted to valid range.
// 인덱스가 범위를 벗어나면 유효한 범위로 조정됩니다.
//
// Example:
//
//	Substring("hello world", 0, 5)   // "hello"
// Substring("hello world", 6, 11)  // "world" / Substring("안녕하세요", 0, 2)       // "안녕"
//	Substring("hello", 0, 100)       // "hello" (auto-adjusted)
func Substring(s string, start, end int) string {
	runes := []rune(s)
	length := len(runes)

	// Adjust negative indices
	// 음수 인덱스 조정
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}

	// Adjust out-of-bounds indices
	// 범위 초과 인덱스 조정
	if start > length {
		start = length
	}
	if end > length {
		end = length
	}

	// Ensure start <= end
	// start <= end 보장
	if start > end {
		start, end = end, start
	}

	return string(runes[start:end])
}

// Left returns the leftmost n characters of a string (Unicode-safe).
// Left는 문자열의 가장 왼쪽 n개 문자를 반환합니다 (유니코드 안전).
//
// If n is greater than string length, returns the entire string.
// n이 문자열 길이보다 크면 전체 문자열을 반환합니다.
//
// Example:
//
// Left("hello world", 5)  // "hello" / Left("안녕하세요", 2)      // "안녕"
//	Left("hello", 10)       // "hello"
func Left(s string, n int) string {
	if n <= 0 {
		return ""
	}
	runes := []rune(s)
	if n >= len(runes) {
		return s
	}
	return string(runes[:n])
}

// Right returns the rightmost n characters of a string (Unicode-safe).
// Right는 문자열의 가장 오른쪽 n개 문자를 반환합니다 (유니코드 안전).
//
// If n is greater than string length, returns the entire string.
// n이 문자열 길이보다 크면 전체 문자열을 반환합니다.
//
// Example:
//
// Right("hello world", 5)  // "world" / Right("안녕하세요", 2)       // "세요"
//	Right("hello", 10)       // "hello"
func Right(s string, n int) string {
	if n <= 0 {
		return ""
	}
	runes := []rune(s)
	length := len(runes)
	if n >= length {
		return s
	}
	return string(runes[length-n:])
}

// Insert inserts a string at the specified index (Unicode-safe).
// Insert는 지정된 인덱스에 문자열을 삽입합니다 (유니코드 안전).
//
// If index is negative or greater than length, it's adjusted to valid range.
// 인덱스가 음수이거나 길이보다 크면 유효한 범위로 조정됩니다.
//
// Example:
//
//	Insert("hello world", 5, ",")    // "hello, world"
// Insert("hello", 0, "say ")       // "say hello" / Insert("안녕하세요", 2, " 반갑습니다 ")  // "안녕 반갑습니다 하세요"
func Insert(s string, index int, insert string) string {
	runes := []rune(s)
	length := len(runes)

	// Adjust negative index
	// 음수 인덱스 조정
	if index < 0 {
		index = 0
	}
	// Adjust out-of-bounds index
	// 범위 초과 인덱스 조정
	if index > length {
		index = length
	}

	// Build result
	// 결과 생성
	result := make([]rune, 0, length+len([]rune(insert)))
	result = append(result, runes[:index]...)
	result = append(result, []rune(insert)...)
	result = append(result, runes[index:]...)

	return string(result)
}

// SwapCase swaps the case of all letters in a string.
// SwapCase는 문자열의 모든 글자의 대소문자를 반전합니다.
//
// Uppercase becomes lowercase and vice versa.
// 대문자는 소문자로, 소문자는 대문자로 변환됩니다.
//
// Example:
//
//	SwapCase("Hello World")  // "hELLO wORLD"
//	SwapCase("GoLang")       // "gOlANG"
//	SwapCase("ABC123xyz")    // "abc123XYZ"
func SwapCase(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return string(runes)
}
