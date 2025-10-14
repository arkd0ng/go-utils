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

// splitIntoWords splits a string into words based on delimiters and case changes.
// splitIntoWords는 구분자와 케이스 변경을 기반으로 문자열을 단어로 분리합니다.
//
// Handles:
// 처리:
//   - Delimiters: -, _, space / 구분자: -, _, 공백
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

		// Skip delimiters / 구분자 건너뛰기
		if r == '-' || r == '_' || r == ' ' {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []rune{}
			}
			continue
		}

		// Handle case changes / 케이스 변경 처리
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
