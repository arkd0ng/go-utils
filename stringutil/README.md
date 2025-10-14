# Stringutil Package / 문자열 유틸리티 패키지

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Extreme simplicity string utility functions for Go. Reduces 10-20 lines of repetitive code to a single function call.

Go를 위한 극도로 간단한 문자열 유틸리티 함수. 10-20줄의 반복 코드를 단일 함수 호출로 줄입니다.

**Design Philosophy / 설계 철학**: "20 lines → 1 line"

## Features / 기능

✨ **37 utility functions** across 5 categories / 5개 카테고리에 걸쳐 37개 유틸리티 함수

🚀 **Zero dependencies** - standard library only / 제로 의존성 - 표준 라이브러리만

🌍 **Unicode-safe** - works with 한글, 日本語, emoji 🎉 / 유니코드 안전

📝 **Bilingual docs** - English/Korean / 이중 언어 문서

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

// String manipulation / 문자열 조작
stringutil.Truncate("Hello World", 8)      // "Hello..."
stringutil.Clean("  hello   world  ")      // "hello world"

// Validation / 유효성 검사
stringutil.IsEmail("user@example.com")     // true
stringutil.IsURL("https://example.com")    // true
```

## API Categories / API 카테고리

### 1. Case Conversion / 케이스 변환 (5 functions)

| Function | Input | Output |
|----------|-------|--------|
| `ToSnakeCase` | `"UserProfileData"` | `"user_profile_data"` |
| `ToCamelCase` | `"user_profile_data"` | `"userProfileData"` |
| `ToKebabCase` | `"UserProfileData"` | `"user-profile-data"` |
| `ToPascalCase` | `"user_profile_data"` | `"UserProfileData"` |
| `ToScreamingSnakeCase` | `"userProfileData"` | `"USER_PROFILE_DATA"` |

### 2. String Manipulation / 문자열 조작 (9 functions)

- `Truncate(s string, length int) string` - Truncates with "..."
- `TruncateWithSuffix(s string, length int, suffix string) string` - Custom suffix
- `Reverse(s string) string` - Reverses string (Unicode-safe)
- `Capitalize(s string) string` - Capitalizes each word
- `CapitalizeFirst(s string) string` - Capitalizes first letter only
- `RemoveDuplicates(s string) string` - Removes duplicate characters
- `RemoveSpaces(s string) string` - Removes all whitespace
- `RemoveSpecialChars(s string) string` - Keeps only alphanumeric
- `Clean(s string) string` - Trims and deduplicates spaces

### 3. Validation / 유효성 검사 (8 functions)

- `IsEmail(s string) bool` - Email format (practical, not RFC 5322)
- `IsURL(s string) bool` - URL with http:// or https://
- `IsAlphanumeric(s string) bool` - Only a-z, A-Z, 0-9
- `IsNumeric(s string) bool` - Only 0-9
- `IsAlpha(s string) bool` - Only a-z, A-Z
- `IsBlank(s string) bool` - Empty or whitespace only
- `IsLower(s string) bool` - All lowercase letters
- `IsUpper(s string) bool` - All uppercase letters

### 4. Search & Replace / 검색 및 치환 (6 functions)

- `ContainsAny(s string, substrs []string) bool`
- `ContainsAll(s string, substrs []string) bool`
- `StartsWithAny(s string, prefixes []string) bool`
- `EndsWithAny(s string, suffixes []string) bool`
- `ReplaceAll(s string, replacements map[string]string) string`
- `ReplaceIgnoreCase(s, old, new string) string`

### 5. Utilities / 유틸리티 (9 functions)

- `CountWords(s string) int`
- `CountOccurrences(s, substr string) int`
- `Join(strs []string, sep string) string`
- `Map(strs []string, fn func(string) string) []string`
- `Filter(strs []string, fn func(string) bool) []string`
- `PadLeft(s string, length int, pad string) string`
- `PadRight(s string, length int, pad string) string`
- `Lines(s string) []string`
- `Words(s string) []string`

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
