# Stringutil Package - Design Plan / 설계 계획서
# stringutil 패키지 - 설계 계획서

**Version / 버전**: v1.5.x
**Author / 작성자**: arkd0ng
**Created / 작성일**: 2025-10-14
**Status / 상태**: Final Design - Extreme Simplicity / 최종 설계 - 극도의 간결함

---

## Table of Contents / 목차

1. [Why This Package Exists / 왜 이 패키지가 존재하는가](#why-this-package-exists--왜-이-패키지가-존재하는가)
2. [Design Philosophy / 설계 철학](#design-philosophy--설계-철학)
3. [What Users Get / 사용자가 얻는 것](#what-users-get--사용자가-얻는-것)
4. [API Design / API 설계](#api-design--api-설계)
5. [Implementation Architecture / 구현 아키텍처](#implementation-architecture--구현-아키텍처)
6. [File Structure / 파일 구조](#file-structure--파일-구조)
7. [Detailed Features / 상세 기능](#detailed-features--상세-기능)

---

## Why This Package Exists / 왜 이 패키지가 존재하는가

### The Problem / 문제점

Go의 표준 라이브러리(`strings` 패키지)를 사용할 때 개발자가 겪는 불편함:

Working with Go's standard library (`strings` package), developers face inconveniences:

1. **복잡한 케이스 변환 / Complex case conversions**:
   ```go
   // Snake case로 변환하려면?
   // Convert to snake_case?
   str := "UserProfileData"
   // 직접 구현해야 함... 정규식? 반복문?
   // Have to implement yourself... regex? loops?
   ```

2. **문자열 잘라내기의 번거로움 / Cumbersome string truncation**:
   ```go
   // 20자로 자르고 "..." 붙이기
   // Truncate to 20 chars and append "..."
   if len(str) > 20 {
       str = str[:20] + "..."
   }
   // 하지만 유니코드 문자는? rune 처리?
   // But what about Unicode? rune handling?
   ```

3. **반복적인 유효성 검사 / Repetitive validation**:
   ```go
   // 이메일 체크
   // Email validation
   if strings.Contains(email, "@") && strings.Contains(email, ".") {
       // 너무 단순... 실제로는 정규식 필요
       // Too simple... actually needs regex
   }

   // URL 체크
   // URL validation
   _, err := url.Parse(str)
   if err != nil || !strings.HasPrefix(str, "http") {
       // ...
   }
   ```

4. **kebab-case, camelCase 변환의 부재 / Missing kebab-case, camelCase conversions**:
   ```go
   // "user-profile-data" → "UserProfileData" 변환?
   // Convert "user-profile-data" → "UserProfileData"?
   // 직접 파싱해야 함
   // Must parse manually
   ```

5. **문자열 정리의 번거로움 / Tedious string cleanup**:
   ```go
   // 공백 제거, 특수문자 제거, 중복 공백 정리...
   // Remove spaces, special chars, clean duplicate spaces...
   str = strings.TrimSpace(str)
   str = strings.ReplaceAll(str, "  ", " ")
   // 여러 단계 필요
   // Multiple steps needed
   ```

### The Solution / 해결책

**이 패키지는 자주 쓰이지만 번거로운 문자열 작업을 한 줄로 해결합니다**:

**This package solves frequently-used but cumbersome string operations in one line**:

```go
import "github.com/arkd0ng/go-utils/stringutil"

// 1. 케이스 변환 - 한 줄로
// Case conversion - one line
stringutil.ToSnakeCase("UserProfileData")  // "user_profile_data"
stringutil.ToCamelCase("user-profile-data") // "userProfileData"
stringutil.ToKebabCase("UserProfileData")   // "user-profile-data"
stringutil.ToPascalCase("user_profile_data") // "UserProfileData"

// 2. 안전한 문자열 자르기 (유니코드 지원)
// Safe string truncation (Unicode support)
stringutil.Truncate("Hello World", 8)          // "Hello..."
stringutil.TruncateWithSuffix("안녕하세요", 3, "…") // "안녕하…"

// 3. 간단한 유효성 검사
// Simple validation
stringutil.IsEmail("user@example.com")    // true
stringutil.IsURL("https://example.com")   // true
stringutil.IsAlphanumeric("abc123")       // true
stringutil.IsNumeric("12345")             // true

// 4. 문자열 정리
// String cleanup
stringutil.Clean("  hello   world  ")     // "hello world"
stringutil.RemoveSpaces("a b c")          // "abc"
stringutil.RemoveSpecialChars("hello@#$") // "hello"

// 5. 유용한 헬퍼
// Useful helpers
stringutil.Reverse("hello")               // "olleh"
stringutil.Contains(str, []string{"foo", "bar"}) // true if any match
stringutil.Capitalize("hello world")      // "Hello World"
stringutil.CountWords("hello world")      // 2
```

### If It's Not This Simple, Don't Build It / 이 정도로 간단하지 않으면 만들지 마세요

**핵심 원칙 / Core Principle**: "20줄 → 1줄" / "20 lines → 1 line"

---

## Design Philosophy / 설계 철학

### 1. Extreme Simplicity / 극도의 간결함

Every function should reduce 10-20 lines of repetitive code to a single function call.

모든 함수는 10-20줄의 반복 코드를 한 번의 함수 호출로 줄여야 합니다.

**Bad / 나쁨**:
```go
// Too complex API
stringutil.Convert(str, stringutil.ConversionOptions{
    From: stringutil.CamelCase,
    To:   stringutil.SnakeCase,
})
```

**Good / 좋음**:
```go
// Simple, direct
stringutil.ToSnakeCase(str)
```

### 2. No Dependencies / 의존성 없음

- Uses only Go standard library
- 오직 Go 표준 라이브러리만 사용
- No external regex libraries, no third-party validation packages
- 외부 정규식 라이브러리 없음, 서드파티 검증 패키지 없음

### 3. Unicode Support / 유니코드 지원

- All functions work correctly with multi-byte characters (한글, 日本語, emoji 🎉)
- 모든 함수는 멀티바이트 문자와 올바르게 동작 (한글, 日本語, emoji 🎉)
- Use `rune` instead of `byte` for length calculations
- 길이 계산에 `byte` 대신 `rune` 사용

### 4. Practical Over Perfect / 완벽보다 실용성

- Email validation: good enough for 99% of cases (not RFC 5322 compliant)
- 이메일 검증: 99%의 경우에 충분함 (RFC 5322 완전 준수 아님)
- URL validation: checks common patterns
- URL 검증: 일반적인 패턴 체크
- Focus on developer productivity, not academic correctness
- 학술적 정확성보다 개발자 생산성에 집중

### 5. Chainable When Useful / 유용할 때 체이닝 가능

```go
// Multiple operations
result := stringutil.Clean(stringutil.ToSnakeCase(str))

// Or create a Builder pattern if chaining is common
builder := stringutil.New("  Hello World  ")
result := builder.Clean().ToSnakeCase().String()
```

---

## What Users Get / 사용자가 얻는 것

### Before / 이전 (표준 라이브러리 사용)

```go
import (
    "regexp"
    "strings"
    "unicode"
)

// Snake case 변환 (20+ lines)
func toSnakeCase(s string) string {
    var result []rune
    for i, r := range s {
        if unicode.IsUpper(r) {
            if i > 0 {
                result = append(result, '_')
            }
            result = append(result, unicode.ToLower(r))
        } else {
            result = append(result, r)
        }
    }
    return string(result)
}

// 이메일 검증 (정규식 작성)
func isEmail(s string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(s)
}

// 문자열 자르기 (유니코드 안전)
func truncate(s string, length int) string {
    runes := []rune(s)
    if len(runes) <= length {
        return s
    }
    return string(runes[:length]) + "..."
}

// 총 50+ lines for 3 functions
```

### After / 이후 (stringutil 패키지 사용)

```go
import "github.com/arkd0ng/go-utils/stringutil"

// 3 lines total
result1 := stringutil.ToSnakeCase("UserProfileData")
result2 := stringutil.IsEmail("user@example.com")
result3 := stringutil.Truncate("Hello World", 8)
```

**Code reduction / 코드 감소**: 50+ lines → 3 lines (94% reduction)

---

## API Design / API 설계

### Category 1: Case Conversion / 케이스 변환

```go
// snake_case 변환
func ToSnakeCase(s string) string

// camelCase 변환
func ToCamelCase(s string) string

// kebab-case 변환
func ToKebabCase(s string) string

// PascalCase 변환
func ToPascalCase(s string) string

// SCREAMING_SNAKE_CASE 변환
func ToScreamingSnakeCase(s string) string
```

**Examples / 예제**:
```go
ToSnakeCase("UserProfileData")     // "user_profile_data"
ToCamelCase("user-profile-data")   // "userProfileData"
ToKebabCase("UserProfileData")     // "user-profile-data"
ToPascalCase("user_profile_data")  // "UserProfileData"
ToScreamingSnakeCase("userName")   // "USER_NAME"
```

### Category 2: String Manipulation / 문자열 조작

```go
// 문자열 자르기 (유니코드 안전, "..." 추가)
func Truncate(s string, length int) string

// 사용자 정의 suffix로 자르기
func TruncateWithSuffix(s string, length int, suffix string) string

// 문자열 뒤집기
func Reverse(s string) string

// 첫 글자 대문자 (각 단어)
func Capitalize(s string) string

// 첫 글자만 대문자
func CapitalizeFirst(s string) string

// 반복 제거
func RemoveDuplicates(s string) string

// 공백 모두 제거
func RemoveSpaces(s string) string

// 특수문자 제거 (영숫자와 공백만 남김)
func RemoveSpecialChars(s string) string

// 공백 정리 (중복 공백 → 단일 공백, trim)
func Clean(s string) string
```

**Examples / 예제**:
```go
Truncate("Hello World", 8)              // "Hello..."
TruncateWithSuffix("안녕하세요", 3, "…")   // "안녕하…"
Reverse("hello")                        // "olleh"
Capitalize("hello world")               // "Hello World"
CapitalizeFirst("hello world")          // "Hello world"
RemoveDuplicates("hello")               // "helo"
RemoveSpaces("h e l l o")               // "hello"
RemoveSpecialChars("hello@#$123")       // "hello123"
Clean("  hello   world  ")              // "hello world"
```

### Category 3: Validation / 유효성 검사

```go
// 이메일 형식 검증 (실용적 수준)
func IsEmail(s string) bool

// URL 형식 검증
func IsURL(s string) bool

// 영숫자 검증 (a-z, A-Z, 0-9만)
func IsAlphanumeric(s string) bool

// 숫자 검증 (0-9만)
func IsNumeric(s string) bool

// 알파벳 검증 (a-z, A-Z만)
func IsAlpha(s string) bool

// 빈 문자열 또는 공백만 체크
func IsBlank(s string) bool

// 소문자만 체크
func IsLower(s string) bool

// 대문자만 체크
func IsUpper(s string) bool
```

**Examples / 예제**:
```go
IsEmail("user@example.com")      // true
IsEmail("invalid.email")         // false
IsURL("https://example.com")     // true
IsAlphanumeric("abc123")         // true
IsNumeric("12345")               // true
IsAlpha("abcABC")                // true
IsBlank("   ")                   // true
IsLower("hello")                 // true
IsUpper("HELLO")                 // true
```

### Category 4: Search & Replace / 검색 및 치환

```go
// 여러 문자열 중 하나라도 포함하는지 체크
func ContainsAny(s string, substrs []string) bool

// 모든 문자열이 포함되는지 체크
func ContainsAll(s string, substrs []string) bool

// 여러 문자열 중 하나로 시작하는지 체크
func StartsWithAny(s string, prefixes []string) bool

// 여러 문자열 중 하나로 끝나는지 체크
func EndsWithAny(s string, suffixes []string) bool

// 여러 문자열 일괄 치환
func ReplaceAll(s string, replacements map[string]string) string

// 대소문자 구분 없이 치환
func ReplaceIgnoreCase(s, old, new string) string
```

**Examples / 예제**:
```go
ContainsAny("hello world", []string{"foo", "world"})  // true
ContainsAll("hello world", []string{"hello", "world"}) // true
StartsWithAny("https://example.com", []string{"http://", "https://"}) // true
EndsWithAny("file.txt", []string{".txt", ".md"})      // true
ReplaceAll("a b c", map[string]string{"a": "x", "b": "y"}) // "x y c"
ReplaceIgnoreCase("Hello World", "hello", "hi")       // "hi World"
```

### Category 5: Utilities / 유틸리티

```go
// 단어 개수 세기
func CountWords(s string) int

// 특정 문자열 출현 횟수
func CountOccurrences(s, substr string) int

// 문자열 배열 연결 (구분자 사용)
func Join(strs []string, sep string) string

// 문자열 배열에 모두 적용
func Map(strs []string, fn func(string) string) []string

// 필터링
func Filter(strs []string, fn func(string) bool) []string

// 패딩
func PadLeft(s string, length int, pad string) string
func PadRight(s string, length int, pad string) string

// 줄바꿈으로 분리
func Lines(s string) []string

// 단어로 분리 (공백 기준)
func Words(s string) []string
```

**Examples / 예제**:
```go
CountWords("hello world")                    // 2
CountOccurrences("hello hello", "hello")     // 2
Join([]string{"a", "b", "c"}, "-")          // "a-b-c"
Map([]string{"a", "b"}, strings.ToUpper)    // ["A", "B"]
Filter([]string{"a", "ab", "abc"}, func(s string) bool {
    return len(s) > 2
})                                          // ["abc"]
PadLeft("5", 3, "0")                        // "005"
PadRight("5", 3, "0")                       // "500"
Lines("line1\nline2\nline3")                // ["line1", "line2", "line3"]
Words("hello world foo")                    // ["hello", "world", "foo"]
```

---

## Implementation Architecture / 구현 아키텍처

### Design Pattern: Pure Functions / 순수 함수

- All functions are stateless and side-effect free
- 모든 함수는 상태가 없고 부작용이 없음
- Easy to test, easy to reason about
- 테스트하기 쉽고 이해하기 쉬움

```go
// No struct, no methods - just pure functions
// 구조체 없음, 메서드 없음 - 순수 함수만

package stringutil

func ToSnakeCase(s string) string {
    // implementation
}

func IsEmail(s string) bool {
    // implementation
}
```

### Alternative: Builder Pattern (Optional) / 빌더 패턴 (선택사항)

체이닝이 유용한 경우를 위한 선택적 빌더:

Optional builder for chaining when useful:

```go
type Builder struct {
    value string
}

func New(s string) *Builder {
    return &Builder{value: s}
}

func (b *Builder) Clean() *Builder {
    b.value = Clean(b.value)
    return b
}

func (b *Builder) ToSnakeCase() *Builder {
    b.value = ToSnakeCase(b.value)
    return b
}

func (b *Builder) String() string {
    return b.value
}

// Usage / 사용법
result := stringutil.New("  UserProfileData  ").
    Clean().
    ToSnakeCase().
    String()  // "user_profile_data"
```

---

## File Structure / 파일 구조

```
stringutil/
├── stringutil.go           # Core functions / 핵심 함수
├── case.go                 # Case conversion functions / 케이스 변환
├── validation.go           # Validation functions / 검증 함수
├── manipulation.go         # String manipulation / 문자열 조작
├── search.go               # Search and replace / 검색 및 치환
├── utils.go                # Utility functions / 유틸리티 함수
├── builder.go              # Optional builder pattern / 선택적 빌더 패턴
├── stringutil_test.go      # Tests / 테스트
├── case_test.go
├── validation_test.go
├── manipulation_test.go
├── search_test.go
├── utils_test.go
├── builder_test.go
└── README.md               # Package documentation / 패키지 문서
```

---

## Detailed Features / 상세 기능

### Feature 1: Unicode-Safe Operations / 유니코드 안전 작업

**Challenge / 과제**: Go의 `len(string)`은 바이트 수를 반환하며, 멀티바이트 문자(한글, 이모지)에서 오작동

**Solution / 해결책**: 모든 길이 계산에 `[]rune` 사용

```go
// ❌ Wrong / 잘못됨
func Truncate(s string, length int) string {
    if len(s) <= length {  // 바이트 길이!
        return s
    }
    return s[:length] + "..."  // 멀티바이트 문자 깨짐!
}

// ✅ Correct / 올바름
func Truncate(s string, length int) string {
    runes := []rune(s)
    if len(runes) <= length {
        return s
    }
    return string(runes[:length]) + "..."
}
```

### Feature 2: Practical Email Validation / 실용적 이메일 검증

**Not RFC 5322 compliant / RFC 5322 완전 준수 아님**, but good enough for 99% of use cases:

```go
func IsEmail(s string) bool {
    // Simple regex: local@domain.tld
    // 간단한 정규식: local@domain.tld
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(s)
}
```

**Validates / 검증**:
- ✅ `user@example.com`
- ✅ `user.name@example.co.uk`
- ✅ `user+tag@example.com`
- ❌ `invalid`
- ❌ `@example.com`
- ❌ `user@`

### Feature 3: Smart Case Conversion / 스마트 케이스 변환

**Handles multiple input formats / 여러 입력 형식 처리**:

```go
// Input can be any case / 입력은 어떤 케이스든 가능
ToSnakeCase("UserProfileData")    // PascalCase input → "user_profile_data"
ToSnakeCase("userProfileData")    // camelCase input → "user_profile_data"
ToSnakeCase("user-profile-data")  // kebab-case input → "user_profile_data"
ToSnakeCase("USER_PROFILE_DATA")  // SCREAMING_SNAKE_CASE → "user_profile_data"

// Algorithm / 알고리즘:
// 1. Split by delimiters (-, _, space) / 구분자로 분리
// 2. Split by uppercase letters / 대문자로 분리
// 3. Join with target delimiter / 목표 구분자로 결합
```

### Feature 4: Clean Function / Clean 함수

**"Clean" = Trim + Deduplicate Spaces / "정리" = 공백 제거 + 중복 공백 제거**

```go
func Clean(s string) string {
    // 1. Trim leading/trailing spaces / 앞뒤 공백 제거
    s = strings.TrimSpace(s)

    // 2. Replace multiple spaces with single space / 중복 공백을 단일 공백으로
    re := regexp.MustCompile(`\s+`)
    s = re.ReplaceAllString(s, " ")

    return s
}

// Examples / 예제:
Clean("  hello   world  ")   // "hello world"
Clean("\t\nhello\t\nworld")  // "hello world"
```

---

## Success Criteria / 성공 기준

This package is successful if / 이 패키지가 성공한 것은:

1. ✅ **Developers save 10-20 lines of code per function / 개발자가 함수당 10-20줄 절약**
2. ✅ **Zero external dependencies / 외부 의존성 제로**
3. ✅ **100% test coverage / 100% 테스트 커버리지**
4. ✅ **Works correctly with Unicode / 유니코드에서 올바르게 동작**
5. ✅ **Simple, predictable API / 간단하고 예측 가능한 API**

---

## Non-Goals / 비목표

What this package **does NOT** aim to do:

이 패키지가 **목표로 하지 않는** 것:

1. ❌ **Perfect RFC compliance / 완벽한 RFC 준수**
   - Email validation is practical, not RFC 5322 perfect
   - 이메일 검증은 실용적이지만 RFC 5322 완벽하지 않음

2. ❌ **Advanced NLP / 고급 자연어 처리**
   - No stemming, lemmatization, or language detection
   - 어간 추출, 표제어 추출, 언어 감지 없음

3. ❌ **Localization / 현지화**
   - No locale-specific string operations
   - 로케일별 문자열 작업 없음

4. ❌ **Performance over readability / 가독성보다 성능**
   - Code is optimized for clarity, not nanosecond-level performance
   - 코드는 명확성을 위해 최적화되며 나노초 수준 성능이 아님

---

## Conclusion / 결론

**Design Goal / 설계 목표**: 자주 쓰이지만 번거로운 문자열 작업을 한 줄로 해결

**Key Principle / 핵심 원칙**: "If it's not dramatically simpler, don't build it"
"극적으로 간단하지 않으면 만들지 마세요"

This package will save developers countless hours of writing repetitive string manipulation code.

이 패키지는 개발자들이 반복적인 문자열 조작 코드를 작성하는 데 드는 수많은 시간을 절약할 것입니다.
