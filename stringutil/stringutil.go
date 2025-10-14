// Package stringutil provides extreme simplicity string utility functions.
// stringutil 패키지는 극도로 간단한 문자열 유틸리티 함수를 제공합니다.
//
// This package reduces 10-20 lines of repetitive string manipulation code
// to a single function call. All functions are Unicode-safe and have no
// external dependencies (standard library only).
//
// 이 패키지는 10-20줄의 반복적인 문자열 조작 코드를 단일 함수 호출로 줄입니다.
// 모든 함수는 유니코드 안전하며 외부 의존성이 없습니다 (표준 라이브러리만).
//
// Design Philosophy: "20 lines → 1 line"
// 설계 철학: "20줄 → 1줄"
//
// Categories:
// - Case Conversion: ToSnakeCase, ToCamelCase, ToKebabCase, ToPascalCase
// - String Manipulation: Truncate, Reverse, Capitalize, Clean
// - Validation: IsEmail, IsURL, IsAlphanumeric, IsNumeric
// - Search & Replace: ContainsAny, ContainsAll, ReplaceAll
// - Utilities: CountWords, PadLeft, Lines, Words
//
// Example:
//
//	import "github.com/arkd0ng/go-utils/stringutil"
//
//	// Case conversion / 케이스 변환
//	stringutil.ToSnakeCase("UserProfileData")  // "user_profile_data"
//	stringutil.ToCamelCase("user-profile-data") // "userProfileData"
//
//	// String manipulation / 문자열 조작
//	stringutil.Truncate("Hello World", 8)      // "Hello..."
//	stringutil.Clean("  hello   world  ")      // "hello world"
//
//	// Validation / 유효성 검사
//	stringutil.IsEmail("user@example.com")     // true
//	stringutil.IsURL("https://example.com")    // true
package stringutil
