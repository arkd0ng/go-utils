# Stringutil Package / 문자열 유틸리티 패키지

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Extreme simplicity string utility functions for Go. Reduces 10-20 lines of repetitive code to a single function call.

Go를 위한 극도로 간단한 문자열 유틸리티 함수. 10-20줄의 반복 코드를 단일 함수 호출로 줄입니다.

**Design Philosophy / 설계 철학**: "20 lines → 1 line"

## Features / 기능

✨ **108+ utility functions** across 13 categories / 13개 카테고리에 걸쳐 108개 이상 유틸리티 함수

🚀 **Minimal dependencies** - standard library + golang.org/x/text / 최소 의존성

🌍 **Unicode-safe** - works with 한글, 日本語, emoji 🎉 / 유니코드 안전

📝 **Bilingual docs** - English/Korean / 이중 언어 문서

🎯 **Comprehensive** - string manipulation, validation, case conversion, Unicode operations, Builder pattern, encoding, distance algorithms, formatting / 포괄적

🔗 **Fluent API** - Builder pattern with 30+ chainable methods / 30개 이상의 체이닝 가능한 메서드를 가진 Builder 패턴

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/stringutil
```

## Quick Start / 빠른 시작

```go
import "github.com/arkd0ng/go-utils/stringutil"

// Builder pattern - fluent API / Builder 패턴 - 유창한 API
result := stringutil.NewBuilder().
    Append("  user profile data  ").
    Clean().
    ToSnakeCase().
    Truncate(15).
    Build()  // "user_profile_da..."

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

// Encoding/Decoding / 인코딩/디코딩
stringutil.Base64Encode("hello")           // "aGVsbG8="
stringutil.URLEncode("hello world")        // "hello+world"
stringutil.HTMLEscape("<script>")          // "&lt;script&gt;"

// Distance & Similarity / 거리 및 유사도
stringutil.LevenshteinDistance("kitten", "sitting")  // 3
stringutil.Similarity("hello", "hallo")              // 0.8

// Formatting / 포맷팅
stringutil.FormatNumber(1000000, ",")      // "1,000,000"
stringutil.FormatBytes(1536)               // "1.5 KB"
stringutil.MaskEmail("john.doe@example.com") // "j******e@example.com"

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

### 10. Builder Pattern / Builder 패턴 (30+ methods)

Fluent API for chaining string operations / 문자열 작업을 체이닝하기 위한 유창한 API

```go
// Basic usage / 기본 사용법
builder := stringutil.NewBuilder()
result := builder.Append("hello").ToUpper().Build()  // "HELLO"

// Complex chaining / 복잡한 체이닝
result := stringutil.NewBuilder().
    Append("  user profile data  ").
    Clean().
    ToSnakeCase().
    Truncate(15).
    Build()  // "user_profile_da..."
```

**Available methods / 사용 가능한 메서드**:
- **Construction**: `NewBuilder()`, `NewBuilderWithString(s)`
- **Case conversion**: `ToSnakeCase()`, `ToCamelCase()`, `ToKebabCase()`, `ToPascalCase()`, `ToTitle()`, `ToUpper()`, `ToLower()`
- **Manipulation**: `Append(s)`, `AppendLine(s)`, `Capitalize()`, `Reverse()`, `Trim()`, `Clean()`, `RemoveSpaces()`, `RemoveSpecialChars()`
- **Truncation**: `Truncate(length)`, `TruncateWithSuffix(length, suffix)`
- **Formatting**: `Slugify()`, `Quote()`, `Unquote()`, `PadLeft(length, pad)`, `PadRight(length, pad)`
- **Transformation**: `Replace(old, new)`, `Repeat(count)`
- **Utility**: `Build()`, `String()`, `Len()`, `Reset()`

### 11. Encoding & Decoding / 인코딩 및 디코딩 (8 functions)

- `Base64Encode(s string) string` - Standard Base64 encoding
- `Base64Decode(s string) (string, error)` - Standard Base64 decoding
- `Base64URLEncode(s string) string` - URL-safe Base64 encoding
- `Base64URLDecode(s string) (string, error)` - URL-safe Base64 decoding
- `URLEncode(s string) string` - URL query string encoding
- `URLDecode(s string) (string, error)` - URL query string decoding
- `HTMLEscape(s string) string` - HTML entity escaping
- `HTMLUnescape(s string) string` - HTML entity unescaping

### 12. Distance & Similarity / 거리 및 유사도 (4 functions)

- `LevenshteinDistance(a, b string) int` - Edit distance (insertions, deletions, substitutions)
- `Similarity(a, b string) float64` - Similarity score (0.0-1.0) based on Levenshtein
- `HammingDistance(a, b string) int` - Count of differing positions (equal-length strings)
- `JaroWinklerSimilarity(a, b string) float64` - Jaro-Winkler similarity (0.0-1.0)

**Use cases / 사용 사례**:
- Fuzzy search / 퍼지 검색
- Typo correction / 오타 수정
- Duplicate detection / 중복 감지
- String matching / 문자열 매칭

### 13. Formatting / 포맷팅 (13 functions)

- `FormatNumber(n int, separator string) string` - Format numbers with thousand separators
- `FormatBytes(bytes int64) string` - Human-readable byte sizes (KB, MB, GB, etc.)
- `Pluralize(count int, singular, plural string) string` - Singular/plural based on count
- `FormatWithCount(count int, singular, plural string) string` - Count with pluralized noun
- `Ellipsis(s string, maxLen int) string` - Truncate with ellipsis in middle
- `Mask(s string, first, last int, maskChar string) string` - Mask middle characters
- `MaskEmail(email string) string` - Mask email (show first/last + domain)
- `MaskCreditCard(card string) string` - Mask credit card (show last 4 digits)
- `AddLineNumbers(s string) string` - Add line numbers to text
- `Indent(s string, prefix string) string` - Indent each line with prefix
- `Dedent(s string) string` - Remove common leading whitespace
- `WrapText(s string, width int) string` - Wrap text to specified width

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
