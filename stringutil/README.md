# Stringutil Package / 문자열 유틸리티 패키지

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Extreme simplicity string utility functions for Go. Reduces 10-20 lines of repetitive code to a single function call.

Go를 위한 극도로 간단한 문자열 유틸리티 함수. 10-20줄의 반복 코드를 단일 함수 호출로 줄입니다.

**Design Philosophy / 설계 철학**: "20 lines → 1 line"

## Features / 기능

✨ **53 utility functions** across 9 categories / 9개 카테고리에 걸쳐 53개 유틸리티 함수

🚀 **Minimal dependencies** - standard library + golang.org/x/text / 최소 의존성

🌍 **Unicode-safe** - works with 한글, 日本語, emoji 🎉 / 유니코드 안전

📝 **Bilingual docs** - English/Korean / 이중 언어 문서

🎯 **Comprehensive** - string manipulation, validation, case conversion, Unicode operations / 포괄적

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/stringutil
```

## Quick Start / 빠른 시작

```go
import "github.com/arkd0ng/go-utils/stringutil"

// Case conversion / 케이스 변환
stringutil.ToSnakeCase("UserProfileData")  // "user_profile_data"
stringutil.ToCamelCase("user-profile-data") // "userProfileData"
stringutil.ToTitle("hello world")           // "Hello World"
stringutil.Slugify("Hello World!")          // "hello-world"

// String manipulation / 문자열 조작
stringutil.Truncate("Hello World", 8)      // "Hello..."
stringutil.Clean("  hello   world  ")      // "hello world"
stringutil.Substring("hello world", 0, 5)  // "hello"
stringutil.Insert("hello world", 5, ",")   // "hello, world"
stringutil.SwapCase("Hello World")         // "hELLO wORLD"

// Validation / 유효성 검사
stringutil.IsEmail("user@example.com")     // true
stringutil.IsURL("https://example.com")    // true

// Comparison / 비교
stringutil.EqualFold("hello", "HELLO")     // true
stringutil.HasPrefix("hello world", "hello") // true

// Unicode operations / 유니코드 작업
stringutil.RuneCount("안녕하세요")          // 5 (not 15 bytes)
stringutil.Width("hello世界")               // 9 (5 + 4)
stringutil.Normalize("café", "NFC")        // "café" (composed)
```

## API Categories / API 카테고리

### 1. Case Conversion / 케이스 변환 (9 functions)

| Function | Input | Output |
|----------|-------|--------|
| `ToSnakeCase` | `"UserProfileData"` | `"user_profile_data"` |
| `ToCamelCase` | `"user_profile_data"` | `"userProfileData"` |
| `ToKebabCase` | `"UserProfileData"` | `"user-profile-data"` |
| `ToPascalCase` | `"user_profile_data"` | `"UserProfileData"` |
| `ToScreamingSnakeCase` | `"userProfileData"` | `"USER_PROFILE_DATA"` |
| `ToTitle` | `"hello world"` | `"Hello World"` |
| `Slugify` | `"Hello World!"` | `"hello-world"` |
| `Quote` | `"hello"` | `"\"hello\""` |
| `Unquote` | `"\"hello\""` | `"hello"` |

### 2. String Manipulation / 문자열 조작 (17 functions)

- `Truncate(s string, length int) string` - Truncates with "..."
- `TruncateWithSuffix(s string, length int, suffix string) string` - Custom suffix
- `Reverse(s string) string` - Reverses string (Unicode-safe)
- `Capitalize(s string) string` - Capitalizes each word
- `CapitalizeFirst(s string) string` - Capitalizes first letter only
- `RemoveDuplicates(s string) string` - Removes duplicate characters
- `RemoveSpaces(s string) string` - Removes all whitespace
- `RemoveSpecialChars(s string) string` - Keeps only alphanumeric
- `Clean(s string) string` - Trims and deduplicates spaces
- `Repeat(s string, count int) string` - Repeats string count times
- `Substring(s string, start, end int) string` - Extracts substring (Unicode-safe)
- `Left(s string, n int) string` - Gets leftmost n characters
- `Right(s string, n int) string` - Gets rightmost n characters
- `Insert(s string, index int, insert string) string` - Inserts at index
- `SwapCase(s string) string` - Swaps upper/lowercase

### 3. Validation / 유효성 검사 (8 functions)

- `IsEmail(s string) bool` - Email format (practical, not RFC 5322)
- `IsURL(s string) bool` - URL with http:// or https://
- `IsAlphanumeric(s string) bool` - Only a-z, A-Z, 0-9
- `IsNumeric(s string) bool` - Only 0-9
- `IsAlpha(s string) bool` - Only a-z, A-Z
- `IsBlank(s string) bool` - Empty or whitespace only
- `IsLower(s string) bool` - All lowercase letters
- `IsUpper(s string) bool` - All uppercase letters

### 4. Comparison / 비교 (3 functions)

- `EqualFold(s1, s2 string) bool` - Case-insensitive comparison
- `HasPrefix(s, prefix string) bool` - Check string prefix
- `HasSuffix(s, suffix string) bool` - Check string suffix

### 5. Search & Replace / 검색 및 치환 (6 functions)

- `ContainsAny(s string, substrs []string) bool`
- `ContainsAll(s string, substrs []string) bool`
- `StartsWithAny(s string, prefixes []string) bool`
- `EndsWithAny(s string, suffixes []string) bool`
- `ReplaceAll(s string, replacements map[string]string) string`
- `ReplaceIgnoreCase(s, old, new string) string`

### 6. Unicode Operations / 유니코드 작업 (3 functions)

- `RuneCount(s string) int` - Count Unicode characters (not bytes)
- `Width(s string) int` - Calculate display width (CJK double-width support)
- `Normalize(s string, form string) string` - Unicode normalization (NFC/NFD/NFKC/NFKD)

### 7. String Generation / 문자열 생성 (2 functions)

- `PadLeft(s string, length int, pad string) string`
- `PadRight(s string, length int, pad string) string`

### 8. String Parsing / 문자열 파싱 (2 functions)

- `Lines(s string) []string`
- `Words(s string) []string`

### 9. Collection Utilities / 컬렉션 유틸리티 (5 functions)

- `CountWords(s string) int`
- `CountOccurrences(s, substr string) int`
- `Join(strs []string, sep string) string`
- `Map(strs []string, fn func(string) string) []string`
- `Filter(strs []string, fn func(string) bool) []string`

## Examples / 예제

See [examples/stringutil](../examples/stringutil/) for complete examples.

완전한 예제는 [examples/stringutil](../examples/stringutil/)을 참조하세요.

## Testing / 테스팅

```bash
go test ./stringutil -v
```

## License / 라이선스

MIT License - see [LICENSE](../LICENSE) file

## Contributing / 기여

Contributions are welcome! Please see [CONTRIBUTING.md](../CONTRIBUTING.md)

기여를 환영합니다! [CONTRIBUTING.md](../CONTRIBUTING.md)를 참조하세요.
