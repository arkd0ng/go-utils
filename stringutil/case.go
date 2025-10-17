package stringutil

import (
	"strings"
	"unicode"
)

// ToSnakeCase converts a string to snake_case.
// ToSnakeCase는 문자열을 snake_case로 변환합니다.
//
// Handles multiple input formats:
// 여러 입력 형식 처리:
//   - PascalCase: "UserProfileData" → "user_profile_data"
//   - camelCase: "userProfileData" → "user_profile_data"
//   - kebab-case: "user-profile-data" → "user_profile_data"
//   - SCREAMING_SNAKE_CASE: "USER_PROFILE_DATA" → "user_profile_data"
//
// Example:
//
//	ToSnakeCase("UserProfileData")  // "user_profile_data"
//	ToSnakeCase("userProfileData")  // "user_profile_data"
//	ToSnakeCase("user-profile-data") // "user_profile_data"
func ToSnakeCase(s string) string {
	words := splitIntoWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "_")
}

// ToCamelCase converts a string to camelCase.
// ToCamelCase는 문자열을 camelCase로 변환합니다.
//
// Example:
//
//	ToCamelCase("user_profile_data")  // "userProfileData"
//	ToCamelCase("user-profile-data")  // "userProfileData"
//	ToCamelCase("UserProfileData")    // "userProfileData"
func ToCamelCase(s string) string {
	words := splitIntoWords(s)
	if len(words) == 0 {
		return ""
	}

	// First word lowercase, rest capitalized
	// 첫 단어는 소문자, 나머지는 대문자로 시작
	result := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			result += strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
		}
	}
	return result
}

// ToKebabCase converts a string to kebab-case.
// ToKebabCase는 문자열을 kebab-case로 변환합니다.
//
// Example:
//
//	ToKebabCase("UserProfileData")   // "user-profile-data"
//	ToKebabCase("user_profile_data") // "user-profile-data"
//	ToKebabCase("userProfileData")   // "user-profile-data"
func ToKebabCase(s string) string {
	words := splitIntoWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "-")
}

// ToPascalCase converts a string to PascalCase.
// ToPascalCase는 문자열을 PascalCase로 변환합니다.
//
// Example:
//
//	ToPascalCase("user_profile_data") // "UserProfileData"
//	ToPascalCase("user-profile-data") // "UserProfileData"
//	ToPascalCase("userProfileData")   // "UserProfileData"
func ToPascalCase(s string) string {
	words := splitIntoWords(s)
	result := ""
	for _, word := range words {
		if len(word) > 0 {
			result += strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return result
}

// ToScreamingSnakeCase converts a string to SCREAMING_SNAKE_CASE.
// ToScreamingSnakeCase는 문자열을 SCREAMING_SNAKE_CASE로 변환합니다.
//
// Example:
//
//	ToScreamingSnakeCase("UserProfileData") // "USER_PROFILE_DATA"
//	ToScreamingSnakeCase("userProfileData") // "USER_PROFILE_DATA"
func ToScreamingSnakeCase(s string) string {
	words := splitIntoWords(s)
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}
	return strings.Join(words, "_")
}

// ToTitle converts a string to Title Case (each word capitalized).
// ToTitle은 문자열을 Title Case로 변환합니다 (각 단어의 첫 글자를 대문자로).
//
// Example:
//
//	ToTitle("hello world")       // "Hello World"
//	ToTitle("user_profile_data") // "User Profile Data"
//	ToTitle("hello-world")       // "Hello World"
func ToTitle(s string) string {
	words := splitIntoWords(s)
	result := make([]string, len(words))
	for i, word := range words {
		if len(word) > 0 {
			result[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(result, " ")
}

// Slugify converts a string to a URL-friendly slug.
// Slugify는 문자열을 URL 친화적인 슬러그로 변환합니다.
//
// Converts to lowercase, replaces spaces and special characters with hyphens,
// and removes consecutive hyphens.
// 소문자로 변환하고, 공백과 특수 문자를 하이픈으로 대체하며,
// 연속된 하이픈을 제거합니다.
//
// Example:
//
//	Slugify("Hello World!")           // "hello-world"
//	Slugify("User Profile Data")      // "user-profile-data"
//	Slugify("Go Utils -- Package")    // "go-utils-package"
func Slugify(s string) string {
	// Convert to lowercase
	// 소문자로 변환
	s = strings.ToLower(s)

	// Replace spaces and special characters with hyphens
	// 공백과 특수 문자를 하이픈으로 대체
	var result []rune
	lastWasHyphen := false

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
			lastWasHyphen = false
		} else if !lastWasHyphen {
			result = append(result, '-')
			lastWasHyphen = true
		}
	}

	// Trim leading and trailing hyphens
	// 앞뒤 하이픈 제거
	slug := strings.Trim(string(result), "-")
	return slug
}

// Quote wraps a string in double quotes and escapes internal quotes.
// Quote는 문자열을 큰따옴표로 감싸고 내부 따옴표를 이스케이프합니다.
//
// Example:
//
//	Quote("hello")       // "\"hello\""
//	Quote("say \"hi\"")  // "\"say \\\"hi\\\"\""
func Quote(s string) string {
	// Use strconv.Quote for proper escaping
	// 적절한 이스케이프를 위해 strconv.Quote 사용
	var result strings.Builder
	result.WriteRune('"')
	for _, r := range s {
		if r == '"' || r == '\\' {
			result.WriteRune('\\')
		}
		result.WriteRune(r)
	}
	result.WriteRune('"')
	return result.String()
}

// Unquote removes surrounding quotes from a string and unescapes internal quotes.
// Unquote는 문자열에서 주변 따옴표를 제거하고 내부 따옴표의 이스케이프를 해제합니다.
//
// Supports both double quotes (") and single quotes (').
// 큰따옴표(")와 작은따옴표(') 모두 지원합니다.
//
// Example:
//
//	Unquote("\"hello\"")       // "hello"
//	Unquote("'world'")         // "world"
//	Unquote("\"say \\\"hi\\\"\"") // "say \"hi\""
func Unquote(s string) string {
	if len(s) < 2 {
		return s
	}

	// Check if string is quoted
	// 문자열이 따옴표로 감싸져 있는지 확인
	if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
		s = s[1 : len(s)-1]
	}

	// Unescape internal quotes
	// 내부 따옴표 이스케이프 해제
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "\\'", "'")
	s = strings.ReplaceAll(s, "\\\\", "\\")

	return s
}

// splitIntoWords splits a string into words based on delimiters and case changes.
// splitIntoWords는 구분자와 케이스 변경을 기반으로 문자열을 단어로 분리합니다.
//
// Handles:
// 처리:
// - Delimiters: -, _, space
// 구분자: -, _, 공백
//   - Case changes: "UserProfile" → ["User", "Profile"]
//   - Consecutive uppercase: "HTTPServer" → ["HTTP", "Server"]
func splitIntoWords(s string) []string {
	if s == "" {
		return []string{}
	}

	var words []string
	var currentWord []rune

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// Skip delimiters
		// 구분자 건너뛰기
		if r == '-' || r == '_' || r == ' ' {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []rune{}
			}
			continue
		}

		// Handle case changes
		// 케이스 변경 처리
		if unicode.IsUpper(r) && len(currentWord) > 0 {
			// Check if previous character was lowercase
			// 이전 문자가 소문자였는지 확인
			if unicode.IsLower(currentWord[len(currentWord)-1]) {
				words = append(words, string(currentWord))
				currentWord = []rune{r}
				continue
			}

			// Check if next character is lowercase (e.g., "HTTPServer" → "HTTP" "Server")
			// 다음 문자가 소문자인지 확인 (예: "HTTPServer" → "HTTP" "Server")
			if i+1 < len(runes) && unicode.IsLower(runes[i+1]) && len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []rune{r}
				continue
			}
		}

		currentWord = append(currentWord, r)
	}

	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	return words
}
