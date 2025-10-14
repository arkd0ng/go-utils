package stringutil

import "strings"

// ContainsAny returns true if the string contains any of the substrings.
// ContainsAny는 문자열이 부분 문자열 중 하나라도 포함하면 true를 반환합니다.
//
// Example:
//
//	ContainsAny("hello world", []string{"foo", "world"})  // true
//	ContainsAny("hello world", []string{"foo", "bar"})    // false
func ContainsAny(s string, substrs []string) bool {
	for _, substr := range substrs {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// ContainsAll returns true if the string contains all of the substrings.
// ContainsAll은 문자열이 모든 부분 문자열을 포함하면 true를 반환합니다.
//
// Example:
//
//	ContainsAll("hello world", []string{"hello", "world"})  // true
//	ContainsAll("hello world", []string{"hello", "foo"})    // false
func ContainsAll(s string, substrs []string) bool {
	for _, substr := range substrs {
		if !strings.Contains(s, substr) {
			return false
		}
	}
	return true
}

// StartsWithAny returns true if the string starts with any of the prefixes.
// StartsWithAny는 문자열이 접두사 중 하나로 시작하면 true를 반환합니다.
//
// Example:
//
//	StartsWithAny("https://example.com", []string{"http://", "https://"})  // true
//	StartsWithAny("ftp://example.com", []string{"http://", "https://"})    // false
func StartsWithAny(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// EndsWithAny returns true if the string ends with any of the suffixes.
// EndsWithAny는 문자열이 접미사 중 하나로 끝나면 true를 반환합니다.
//
// Example:
//
//	EndsWithAny("file.txt", []string{".txt", ".md"})  // true
//	EndsWithAny("file.jpg", []string{".txt", ".md"})  // false
func EndsWithAny(s string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// ReplaceAll replaces multiple strings at once using a replacement map.
// ReplaceAll은 치환 맵을 사용하여 여러 문자열을 한 번에 치환합니다.
//
// Example:
//
//	ReplaceAll("a b c", map[string]string{"a": "x", "b": "y"})  // "x y c"
//	ReplaceAll("hello world", map[string]string{"hello": "hi", "world": "there"})  // "hi there"
func ReplaceAll(s string, replacements map[string]string) string {
	for old, new := range replacements {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}

// ReplaceIgnoreCase replaces a substring ignoring case.
// ReplaceIgnoreCase는 대소문자를 무시하고 부분 문자열을 치환합니다.
//
// Example:
//
//	ReplaceIgnoreCase("Hello World", "hello", "hi")  // "hi World"
//	ReplaceIgnoreCase("HELLO World", "hello", "hi")  // "hi World"
func ReplaceIgnoreCase(s, old, new string) string {
	lowerS := strings.ToLower(s)
	lowerOld := strings.ToLower(old)

	var result strings.Builder
	for len(lowerS) > 0 {
		index := strings.Index(lowerS, lowerOld)
		if index == -1 {
			result.WriteString(s)
			break
		}

		// Write everything before the match
		result.WriteString(s[:index])
		// Write the replacement
		result.WriteString(new)

		// Move past the match
		s = s[index+len(old):]
		lowerS = lowerS[index+len(lowerOld):]
	}

	return result.String()
}
