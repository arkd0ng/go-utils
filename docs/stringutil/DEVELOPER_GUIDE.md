# Stringutil Package - Developer Guide
# Stringutil 패키지 - 개발자 가이드

**Version / 버전**: v1.5.018+
**Package / 패키지**: `github.com/arkd0ng/go-utils/stringutil`
**Design Philosophy / 설계 철학**: "20 lines → 1 line" (Extreme Simplicity / 극도의 간결함)
**Function Count / 함수 개수**: 53 functions across 9 categories / 9개 카테고리에 걸친 53개 함수

> **Note**: This guide was initially written for v1.5.x (37 functions). The package has been expanded to 53 functions with additional files:
> - `comparison.go` (NEW, 3 functions): EqualFold, HasPrefix, HasSuffix
> - `manipulation.go` (extended, 6 functions): Repeat, Substring, Left, Right, Insert, SwapCase
> - `case.go` (extended, 4 functions): ToTitle, Slugify, Quote, Unquote
> - `unicode.go` (NEW, 3 functions): RuneCount, Width, Normalize
>
> For complete API reference, see [stringutil/README.md](../../stringutil/README.md)
>
> **참고**: 이 가이드는 처음에 v1.5.x (37개 함수)용으로 작성되었습니다. 패키지는 추가 파일과 함께 53개 함수로 확장되었습니다.
> 전체 API 참조는 [stringutil/README.md](../../stringutil/README.md)를 참조하세요.

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#architecture-overview--아키텍처-개요)
2. [Package Structure / 패키지 구조](#package-structure--패키지-구조)
3. [Core Components / 핵심 컴포넌트](#core-components--핵심-컴포넌트)
4. [Internal Implementation / 내부 구현](#internal-implementation--내부-구현)
5. [Design Patterns / 디자인 패턴](#design-patterns--디자인-패턴)
6. [Adding New Features / 새 기능 추가](#adding-new-features--새-기능-추가)
7. [Testing Guide / 테스트 가이드](#testing-guide--테스트-가이드)
8. [Performance / 성능](#performance--성능)
9. [Contributing Guidelines / 기여 가이드라인](#contributing-guidelines--기여-가이드라인)
10. [Code Style / 코드 스타일](#code-style--코드-스타일)

---

## Architecture Overview / 아키텍처 개요

### Design Philosophy / 설계 철학

The stringutil package follows the principle of **"20 lines → 1 line"** - taking common string operations that typically require 10-20 lines of code and reducing them to a single function call.

stringutil 패키지는 **"20줄 → 1줄"** 원칙을 따릅니다 - 일반적으로 10-20줄의 코드가 필요한 일반적인 문자열 작업을 단일 함수 호출로 줄입니다.

**Key Principles / 주요 원칙**:

1. **Extreme Simplicity / 극도의 간결함**: Every function should be as simple as possible
2. **Unicode Safety / 유니코드 안전**: All functions use `[]rune` for proper Unicode handling
3. **Zero Dependencies / 제로 의존성**: Only standard library, no external dependencies
4. **Practical over Perfect / 완벽보다 실용성**: 99% coverage is better than 100% complexity
5. **Composability / 조합 가능성**: Functions can be combined for complex operations

1. **극도의 간결함**: 모든 함수는 최대한 간단해야 합니다
2. **유니코드 안전**: 모든 함수는 적절한 유니코드 처리를 위해 `[]rune`을 사용합니다
3. **제로 의존성**: 표준 라이브러리만, 외부 의존성 없음
4. **완벽보다 실용성**: 99% 커버리지가 100% 복잡성보다 낫습니다
5. **조합 가능성**: 복잡한 작업을 위해 함수를 결합할 수 있습니다

---

### High-Level Architecture / 상위 수준 아키텍처

```
stringutil/
├── stringutil.go         # Package documentation / 패키지 문서
├── case.go              # Case conversion functions / 케이스 변환 함수
├── manipulation.go      # String manipulation / 문자열 조작
├── validation.go        # Validation functions / 검증 함수
├── search.go            # Search & replace / 검색 및 치환
├── utils.go             # Utility functions / 유틸리티 함수
├── case_test.go         # Case conversion tests / 케이스 변환 테스트
├── manipulation_test.go # Manipulation tests / 조작 테스트
├── validation_test.go   # Validation tests / 검증 테스트
└── README.md            # API documentation / API 문서
```

**Architecture Diagram / 아키텍처 다이어그램**:

```
┌─────────────────────────────────────────────────────────────┐
│                   stringutil Package                         │
│                                                              │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │     Case     │  │ Manipulation │  │  Validation  │     │
│  │  Conversion  │  │              │  │              │     │
│  │  (5 funcs)   │  │  (9 funcs)   │  │  (8 funcs)   │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
│                                                              │
│  ┌──────────────┐  ┌──────────────┐                        │
│  │   Search &   │  │  Utilities   │                        │
│  │   Replace    │  │              │                        │
│  │  (6 funcs)   │  │  (9 funcs)   │                        │
│  └──────────────┘  └──────────────┘                        │
│                                                              │
│           All functions are Unicode-safe                    │
│           모든 함수는 유니코드 안전                          │
└─────────────────────────────────────────────────────────────┘
```

---

### Design Decisions / 설계 결정

#### 1. Unicode Safety / 유니코드 안전

**Decision / 결정**: Use `[]rune` instead of byte operations.

**Rationale / 근거**: Go strings are UTF-8 encoded byte slices. Using byte indices breaks multi-byte characters (emoji, CJK, etc.).

**Example / 예제**:
```go
// Bad: Byte-based (breaks Unicode) / 나쁨: 바이트 기반 (유니코드 깨짐)
func truncateBad(s string, length int) string {
    if len(s) <= length {
        return s
    }
    return s[:length] + "..."  // ❌ Breaks at byte boundary!
}

// Good: Rune-based (Unicode-safe) / 좋음: Rune 기반 (유니코드 안전)
func Truncate(s string, length int) string {
    runes := []rune(s)  // Convert to runes / rune으로 변환
    if len(runes) <= length {
        return s
    }
    return string(runes[:length]) + "..."  // ✅ Safe!
}
```

---

#### 2. Zero Dependencies / 제로 의존성

**Decision / 결정**: Only use Go standard library.

**Rationale / 근거**:
- Reduces dependency bloat / 의존성 비대화 감소
- Improves security (fewer attack vectors) / 보안 향상 (공격 벡터 감소)
- Simplifies maintenance / 유지보수 간소화
- Faster compilation / 더 빠른 컴파일

**Standard Library Imports / 표준 라이브러리 임포트**:
```go
import (
    "regexp"          // For validation patterns / 검증 패턴용
    "strings"         // String operations / 문자열 작업
    "unicode"         // Unicode checks / 유니코드 확인
)
```

---

#### 3. Practical Validation / 실용적 검증

**Decision / 결정**: Provide practical validation (99% coverage) instead of RFC-compliant validation (100% complexity).

**Rationale / 근거**:
- RFC 5322 email validation is extremely complex (6,535 lines of regex)
- 99% of real-world emails match simple patterns
- Users can use specialized libraries for strict validation

**Example / 예제**:
```go
// Practical email validation (99% coverage) / 실용적 이메일 검증 (99% 커버리지)
func IsEmail(s string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(s)
}
// Covers: user@example.com, user+tag@example.com, user.name@example.co.uk
// Doesn't cover: rare edge cases like quoted strings, comments, etc.
```

---

#### 4. Smart Case Conversion / 스마트 케이스 변환

**Decision / 결정**: Handle multiple input formats intelligently.

**Rationale / 근거**: Users input strings in various formats (PascalCase, camelCase, snake_case, kebab-case). The functions should handle all of them.

**Implementation / 구현**: Use `splitIntoWords()` helper that detects:
- Delimiters: `-`, `_`, ` ` (space)
- Case changes: `lowercase` → `Uppercase`
- Consecutive uppercase: `HTTP` → `H`, `T`, `T`, `P` (special handling)

---

## Package Structure / 패키지 구조

### File Organization / 파일 구성

Each file contains a specific category of functions:

각 파일에는 특정 카테고리의 함수가 포함됩니다:

#### `stringutil.go` (Package Documentation)

```go
// Package stringutil provides extreme simplicity string utility functions.
// Design Philosophy: "20 lines → 1 line"
//
// Categories:
// - Case Conversion: ToSnakeCase, ToCamelCase, ToKebabCase, ToPascalCase
// - String Manipulation: Truncate, Reverse, Capitalize, Clean
// - Validation: IsEmail, IsURL, IsAlphanumeric, IsNumeric
// - Search & Replace: ContainsAny, ContainsAll, ReplaceAll
// - Utilities: CountWords, PadLeft, Lines, Words
package stringutil
```

**Purpose / 목적**: Package-level documentation and imports.

---

#### `case.go` (Case Conversion - 163 lines)

**Functions / 함수**:
- `ToSnakeCase(s string) string`
- `ToCamelCase(s string) string`
- `ToKebabCase(s string) string`
- `ToPascalCase(s string) string`
- `ToScreamingSnakeCase(s string) string`

**Helper / 헬퍼**:
- `splitIntoWords(s string) []string` - Smart word splitting / 스마트 단어 분할

**Purpose / 목적**: Convert between naming conventions.

---

#### `manipulation.go` (String Manipulation - 139 lines)

**Functions / 함수**:
- `Truncate(s string, length int) string`
- `TruncateWithSuffix(s string, length int, suffix string) string`
- `Reverse(s string) string`
- `Capitalize(s string) string`
- `CapitalizeFirst(s string) string`
- `RemoveDuplicates(s string) string`
- `RemoveSpaces(s string) string`
- `RemoveSpecialChars(s string) string`
- `Clean(s string) string`

**Purpose / 목적**: Manipulate string content (truncate, reverse, clean, etc.).

---

#### `validation.go` (Validation - 170 lines)

**Functions / 함수**:
- `IsEmail(s string) bool`
- `IsURL(s string) bool`
- `IsAlphanumeric(s string) bool`
- `IsNumeric(s string) bool`
- `IsAlpha(s string) bool`
- `IsBlank(s string) bool`
- `IsLower(s string) bool`
- `IsUpper(s string) bool`

**Purpose / 목적**: Validate string format and content.

---

#### `search.go` (Search & Replace - 114 lines)

**Functions / 함수**:
- `ContainsAny(s string, substrs []string) bool`
- `ContainsAll(s string, substrs []string) bool`
- `StartsWithAny(s string, prefixes []string) bool`
- `EndsWithAny(s string, suffixes []string) bool`
- `ReplaceAll(s string, replacements map[string]string) string`
- `ReplaceIgnoreCase(s, old, new string) string`

**Purpose / 목적**: Search for patterns and replace text.

---

#### `utils.go` (Utilities - 128 lines)

**Functions / 함수**:
- `CountWords(s string) int`
- `CountOccurrences(s, substr string) int`
- `Join(strs []string, sep string) string`
- `Map(strs []string, fn func(string) string) []string`
- `Filter(strs []string, fn func(string) bool) []string`
- `PadLeft(s string, length int, pad string) string`
- `PadRight(s string, length int, pad string) string`
- `Lines(s string) []string`
- `Words(s string) []string`

**Purpose / 목적**: Utility helpers for common operations.

---

### Dependencies / 의존성

**Internal Dependencies / 내부 의존성**: None (files are independent)

**External Dependencies / 외부 의존성**: Zero (standard library only)

```go
// Standard library imports used / 사용된 표준 라이브러리 임포트
import (
    "regexp"      // For validation patterns / 검증 패턴용
    "strings"     // String operations / 문자열 작업
    "unicode"     // Unicode character checks / 유니코드 문자 확인
)
```

---

## Core Components / 핵심 컴포넌트

### 1. Smart Word Splitting / 스마트 단어 분할

**Location / 위치**: `case.go`

**Function / 함수**:
```go
// splitIntoWords splits a string into words intelligently.
// It handles delimiters (-, _, space) and case changes.
// splitIntoWords는 문자열을 지능적으로 단어로 분할합니다.
// 구분자 (-, _, 공백) 및 케이스 변경을 처리합니다.
func splitIntoWords(s string) []string
```

**Algorithm / 알고리즘**:

1. Convert string to runes / 문자열을 rune으로 변환
2. Iterate through runes / rune을 반복
3. Detect word boundaries / 단어 경계 감지:
   - Delimiter characters: `-`, `_`, ` `
   - Case change: `lowercase` → `Uppercase`
   - Consecutive uppercase: `HTTPServer` → `HTTP`, `Server`
4. Build word list / 단어 목록 생성
5. Filter empty words / 빈 단어 필터링

**Example Flow / 예제 흐름**:

```
Input: "UserProfileHTTPData"

Step 1: Convert to runes
['U', 's', 'e', 'r', 'P', 'r', 'o', 'f', 'i', 'l', 'e', 'H', 'T', 'T', 'P', 'D', 'a', 't', 'a']

Step 2: Detect boundaries
'U' → Start of word 1
's' → Continue
'e' → Continue
'r' → Continue
'P' → Case change! Start of word 2
...
'H' → Case change! Start of word 3
'T' → Uppercase continue
'T' → Uppercase continue
'P' → Uppercase continue
'D' → Case change after uppercase run! Split before 'D'
...

Result: ["User", "Profile", "HTTP", "Data"]
```

**Implementation Details / 구현 세부사항**:

```go
func splitIntoWords(s string) []string {
    if s == "" {
        return []string{}
    }

    runes := []rune(s)
    var words []string
    var currentWord []rune

    for i := 0; i < len(runes); i++ {
        r := runes[i]

        // Check if delimiter / 구분자 확인
        if r == '-' || r == '_' || r == ' ' {
            if len(currentWord) > 0 {
                words = append(words, string(currentWord))
                currentWord = []rune{}
            }
            continue
        }

        // Check case changes / 케이스 변경 확인
        if i > 0 {
            prevRune := runes[i-1]

            // Lowercase to uppercase transition / 소문자에서 대문자 전환
            if unicode.IsLower(prevRune) && unicode.IsUpper(r) {
                if len(currentWord) > 0 {
                    words = append(words, string(currentWord))
                    currentWord = []rune{}
                }
            }

            // Uppercase run followed by lowercase / 대문자 다음 소문자
            // Example: "HTTPServer" → "HTTP" + "Server"
            if i > 1 && unicode.IsUpper(prevRune) && unicode.IsUpper(runes[i-2]) && unicode.IsLower(r) {
                if len(currentWord) > 1 {
                    words = append(words, string(currentWord[:len(currentWord)-1]))
                    currentWord = []rune{prevRune}
                }
            }
        }

        currentWord = append(currentWord, r)
    }

    if len(currentWord) > 0 {
        words = append(words, string(currentWord))
    }

    return words
}
```

---

### 2. Unicode-Safe Truncation / 유니코드 안전 자르기

**Location / 위치**: `manipulation.go`

**Function / 함수**:
```go
func Truncate(s string, length int) string {
    return TruncateWithSuffix(s, length, "...")
}

func TruncateWithSuffix(s string, length int, suffix string) string
```

**Why Runes? / 왜 Rune인가?**

```go
// Problem with bytes / 바이트 문제
text := "안녕하세요"  // 5 characters, 15 bytes (UTF-8)
fmt.Println(len(text))        // 15 (bytes!)
fmt.Println(text[:3])         // "안" (WRONG! Shows garbage)

// Solution with runes / Rune 솔루션
runes := []rune(text)
fmt.Println(len(runes))       // 5 (characters!)
fmt.Println(string(runes[:3]))  // "안녕하" (CORRECT!)
```

**Implementation / 구현**:

```go
func TruncateWithSuffix(s string, length int, suffix string) string {
    // Convert to runes for Unicode safety / 유니코드 안전을 위해 rune으로 변환
    runes := []rune(s)

    // If string is shorter than limit, return as-is
    // 문자열이 제한보다 짧으면 그대로 반환
    if len(runes) <= length {
        return s
    }

    // Truncate and add suffix / 자르고 접미사 추가
    return string(runes[:length]) + suffix
}
```

**Test Cases / 테스트 케이스**:

```go
// ASCII
Truncate("Hello World", 8)  // "Hello Wo..."

// Korean
Truncate("안녕하세요", 3)    // "안녕하..."

// Emoji
Truncate("😀😁😂😃😄", 3)  // "😀😁😂..."

// Mixed
Truncate("Hello 世界", 8)   // "Hello 世界"
```

---

### 3. Practical Email Validation / 실용적 이메일 검증

**Location / 위치**: `validation.go`

**Function / 함수**:
```go
func IsEmail(s string) bool
```

**Pattern / 패턴**:
```regex
^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$
```

**Pattern Breakdown / 패턴 분석**:

```
^                        # Start of string / 문자열 시작
[a-zA-Z0-9._%+\-]+      # Local part: letters, digits, ._%+- / 로컬 부분
@                        # @ symbol / @ 기호
[a-zA-Z0-9.\-]+         # Domain: letters, digits, .- / 도메인
\.                       # Dot before TLD / TLD 전 점
[a-zA-Z]{2,}            # TLD: at least 2 letters / TLD: 최소 2글자
$                        # End of string / 문자열 끝
```

**Covers / 커버**:
- ✅ `user@example.com`
- ✅ `user.name@example.com`
- ✅ `user+tag@example.com`
- ✅ `user_123@example.co.uk`

**Doesn't Cover / 커버하지 않음**:
- ❌ Quoted strings: `"user name"@example.com`
- ❌ Comments: `user(comment)@example.com`
- ❌ IP addresses: `user@[192.168.1.1]`
- ❌ Unicode domains: `user@例え.jp`

**Why This Trade-off? / 왜 이런 절충안인가?**

- RFC 5322 regex: 6,535 characters, complex parsing
- Our regex: 56 characters, covers 99% of real emails
- Users can use specialized libraries for strict validation

---

### 4. Map and Filter (Functional Programming) / Map과 Filter (함수형 프로그래밍)

**Location / 위치**: `utils.go`

**Functions / 함수**:
```go
func Map(strs []string, fn func(string) string) []string
func Filter(strs []string, fn func(string) bool) []string
```

**Purpose / 목적**: Enable functional programming patterns for string transformations.

함수형 프로그래밍 패턴으로 문자열 변환을 가능하게 합니다.

**Map Implementation / Map 구현**:

```go
// Map applies a function to each string in a slice and returns a new slice.
// Map은 슬라이스의 각 문자열에 함수를 적용하고 새 슬라이스를 반환합니다.
func Map(strs []string, fn func(string) string) []string {
    // Pre-allocate result slice / 결과 슬라이스 사전 할당
    result := make([]string, len(strs))

    // Apply function to each element / 각 요소에 함수 적용
    for i, s := range strs {
        result[i] = fn(s)
    }

    return result
}
```

**Filter Implementation / Filter 구현**:

```go
// Filter returns a new slice containing only strings that match the predicate.
// Filter는 조건에 일치하는 문자열만 포함하는 새 슬라이스를 반환합니다.
func Filter(strs []string, fn func(string) bool) []string {
    // Don't pre-allocate (we don't know final size)
    // 사전 할당 안 함 (최종 크기를 모름)
    var result []string

    // Add matching elements / 일치하는 요소 추가
    for _, s := range strs {
        if fn(s) {
            result = append(result, s)
        }
    }

    return result
}
```

**Usage Examples / 사용 예제**:

```go
// Map: Transform all strings / Map: 모든 문자열 변환
words := []string{"hello", "world", "foo"}

// Built-in function / 내장 함수
upper := stringutil.Map(words, strings.ToUpper)
// Result: ["HELLO", "WORLD", "FOO"]

// Custom function / 사용자 정의 함수
prefixed := stringutil.Map(words, func(s string) string {
    return "prefix_" + s
})
// Result: ["prefix_hello", "prefix_world", "prefix_foo"]

// Filter: Select matching strings / Filter: 일치하는 문자열 선택
long := stringutil.Filter(words, func(s string) bool {
    return len(s) > 3
})
// Result: ["hello", "world"]
```

---

### 5. Multi-Pattern Replace / 다중 패턴 치환

**Location / 위치**: `search.go`

**Function / 함수**:
```go
func ReplaceAll(s string, replacements map[string]string) string
```

**Purpose / 목적**: Replace multiple patterns in a single pass.

단일 패스에서 여러 패턴을 치환합니다.

**Implementation / 구현**:

```go
// ReplaceAll replaces all occurrences of multiple patterns.
// ReplaceAll은 여러 패턴의 모든 발생을 치환합니다.
func ReplaceAll(s string, replacements map[string]string) string {
    // Iterate through replacement map / 치환 맵 반복
    for old, new := range replacements {
        s = strings.ReplaceAll(s, old, new)
    }
    return s
}
```

**Note on Map Iteration / 맵 반복 참고사항**:

Go maps have **undefined iteration order**. If replacements overlap, results may be unpredictable.

Go 맵은 **정의되지 않은 반복 순서**를 가집니다. 치환이 겹치면 결과가 예측 불가능할 수 있습니다.

**Example / 예제**:

```go
// Safe: No overlapping patterns / 안전: 겹치는 패턴 없음
replacements := map[string]string{
    "hello": "hi",
    "world": "universe",
    "foo":   "bar",
}
result := stringutil.ReplaceAll("hello world foo", replacements)
// Result: "hi universe bar"

// Unsafe: Overlapping patterns / 불안전: 겹치는 패턴
overlapping := map[string]string{
    "hello": "hi",
    "hi":    "hey",  // ⚠️ Overlap!
}
result := stringutil.ReplaceAll("hello", overlapping)
// Result: Undefined! Could be "hi" or "hey"
```

---

## Internal Implementation / 내부 구현

### Flow Diagram: ToSnakeCase / 흐름 다이어그램: ToSnakeCase

```
Input: "UserProfileData"
    ↓
splitIntoWords()
    ↓
["User", "Profile", "Data"]
    ↓
Convert each to lowercase
    ↓
["user", "profile", "data"]
    ↓
Join with "_"
    ↓
Output: "user_profile_data"
```

---

### Flow Diagram: Truncate / 흐름 다이어그램: Truncate

```
Input: "안녕하세요 반갑습니다", length=5
    ↓
Convert to []rune
    ↓
[안, 녕, 하, 세, 요,  , 반, 갑, 습, 니, 다]
    ↓
len(runes) > length? (11 > 5 = true)
    ↓
Take first 5 runes
    ↓
[안, 녕, 하, 세, 요]
    ↓
Convert to string and add "..."
    ↓
Output: "안녕하세요..."
```

---

### Flow Diagram: Map / 흐름 다이어그램: Map

```
Input: ["hello", "world"], fn=strings.ToUpper
    ↓
Create result slice (length=2)
    ↓
For i=0: Apply fn("hello") → "HELLO"
result[0] = "HELLO"
    ↓
For i=1: Apply fn("world") → "WORLD"
result[1] = "WORLD"
    ↓
Output: ["HELLO", "WORLD"]
```

---

### Flow Diagram: Filter / 흐름 다이어그램: Filter

```
Input: ["hello", "world", "a", "foo"], fn=len(s)>3
    ↓
Create empty result slice
    ↓
Check "hello": len>3? Yes → Add to result
result = ["hello"]
    ↓
Check "world": len>3? Yes → Add to result
result = ["hello", "world"]
    ↓
Check "a": len>3? No → Skip
    ↓
Check "foo": len>3? No → Skip
    ↓
Output: ["hello", "world"]
```

---

## Design Patterns / 디자인 패턴

### 1. Helper Function Pattern / 헬퍼 함수 패턴

**Pattern / 패턴**: Extract common logic into helper functions.

공통 로직을 헬퍼 함수로 추출합니다.

**Example / 예제**:

```go
// Public API / 공개 API
func ToSnakeCase(s string) string {
    words := splitIntoWords(s)  // Use helper / 헬퍼 사용
    for i, word := range words {
        words[i] = strings.ToLower(word)
    }
    return strings.Join(words, "_")
}

func ToCamelCase(s string) string {
    words := splitIntoWords(s)  // Reuse same helper / 동일한 헬퍼 재사용
    for i, word := range words {
        if i == 0 {
            words[i] = strings.ToLower(word)
        } else {
            words[i] = strings.Title(strings.ToLower(word))
        }
    }
    return strings.Join(words, "")
}

// Private helper (shared logic) / 비공개 헬퍼 (공유 로직)
func splitIntoWords(s string) []string {
    // Complex word splitting logic / 복잡한 단어 분할 로직
    // ...
}
```

**Benefits / 이점**:
- ✅ DRY (Don't Repeat Yourself)
- ✅ Single source of truth for word splitting
- ✅ Easier to test and maintain

---

### 2. Wrapper Pattern / 래퍼 패턴

**Pattern / 패턴**: Provide convenience wrappers with default parameters.

기본 매개변수가 있는 편의 래퍼를 제공합니다.

**Example / 예제**:

```go
// Core function with full control / 완전한 제어를 가진 핵심 함수
func TruncateWithSuffix(s string, length int, suffix string) string {
    runes := []rune(s)
    if len(runes) <= length {
        return s
    }
    return string(runes[:length]) + suffix
}

// Convenience wrapper with default suffix / 기본 접미사가 있는 편의 래퍼
func Truncate(s string, length int) string {
    return TruncateWithSuffix(s, length, "...")  // Default: "..."
}
```

**Benefits / 이점**:
- ✅ Simple API for common cases
- ✅ Full control when needed
- ✅ Backward compatibility

---

### 3. Higher-Order Function Pattern / 고차 함수 패턴

**Pattern / 패턴**: Functions that accept functions as parameters.

함수를 매개변수로 받는 함수입니다.

**Example / 예제**:

```go
// Higher-order function / 고차 함수
func Map(strs []string, fn func(string) string) []string {
    result := make([]string, len(strs))
    for i, s := range strs {
        result[i] = fn(s)  // Call provided function / 제공된 함수 호출
    }
    return result
}

// Usage with different functions / 다른 함수로 사용
upper := Map(words, strings.ToUpper)
lower := Map(words, strings.ToLower)
snake := Map(words, ToSnakeCase)
custom := Map(words, func(s string) string {
    return "prefix_" + s
})
```

**Benefits / 이점**:
- ✅ Highly flexible and reusable
- ✅ Functional programming style
- ✅ Composable operations

---

### 4. Predicate Pattern / 조건 패턴

**Pattern / 패턴**: Use boolean-returning functions for filtering.

필터링을 위해 부울 반환 함수를 사용합니다.

**Example / 예제**:

```go
// Filter with predicate / 조건으로 필터링
func Filter(strs []string, fn func(string) bool) []string {
    var result []string
    for _, s := range strs {
        if fn(s) {  // Call predicate / 조건 호출
            result = append(result, s)
        }
    }
    return result
}

// Predicate functions / 조건 함수
func isLong(s string) bool {
    return len(s) > 5
}

func hasDigit(s string) bool {
    for _, r := range s {
        if unicode.IsDigit(r) {
            return true
        }
    }
    return false
}

// Usage / 사용
long := Filter(words, isLong)
withDigits := Filter(words, hasDigit)
```

---

### 5. Builder Pattern (strings.Builder) / 빌더 패턴

**Pattern / 패턴**: Use `strings.Builder` for efficient string concatenation.

효율적인 문자열 연결을 위해 `strings.Builder`를 사용합니다.

**Example / 예제**:

```go
// Bad: String concatenation (inefficient) / 나쁨: 문자열 연결 (비효율적)
func joinBad(strs []string, sep string) string {
    result := ""
    for i, s := range strs {
        if i > 0 {
            result += sep  // ❌ Creates new string each time!
        }
        result += s
    }
    return result
}

// Good: Use strings.Builder / 좋음: strings.Builder 사용
func Join(strs []string, sep string) string {
    var builder strings.Builder
    for i, s := range strs {
        if i > 0 {
            builder.WriteString(sep)  // ✅ Efficient!
        }
        builder.WriteString(s)
    }
    return builder.String()
}
```

**Why Builder? / 왜 Builder인가?**

- Strings are immutable in Go / Go에서 문자열은 불변입니다
- Each `+=` creates a new string (O(n²) complexity)
- `strings.Builder` uses a growable buffer (O(n) complexity)

---

## Adding New Features / 새 기능 추가

### Step-by-Step Guide / 단계별 가이드

#### Step 1: Identify the Category / 카테고리 식별

Determine which category your function belongs to:

함수가 속한 카테고리를 결정하세요:

- **Case Conversion** (`case.go`): Naming convention changes
- **Manipulation** (`manipulation.go`): Content modification
- **Validation** (`validation.go`): Format/content checks
- **Search & Replace** (`search.go`): Pattern matching
- **Utilities** (`utils.go`): Helper functions

---

#### Step 2: Write the Function / 함수 작성

**Template / 템플릿**:

```go
// FunctionName does X and returns Y.
// FunctionName는 X를 수행하고 Y를 반환합니다.
//
// Example / 예제:
//  result := stringutil.FunctionName("input")
//  fmt.Println(result)  // output
//
// Unicode Support / 유니코드 지원: Yes/No
func FunctionName(param1 type1, param2 type2) returnType {
    // Implementation / 구현
    // ...
}
```

**Example: Adding TitleCase / 예제: TitleCase 추가**:

```go
// ToTitleCase converts a string to Title Case format (每个单词首字母大写).
// ToTitleCase는 문자열을 Title Case 형식으로 변환합니다 (각 단어 첫 글자 대문자).
//
// Example / 예제:
//  result := stringutil.ToTitleCase("hello world")
//  fmt.Println(result)  // Hello World
//
// Unicode Support / 유니코드 지원: Yes
func ToTitleCase(s string) string {
    words := splitIntoWords(s)
    for i, word := range words {
        runes := []rune(word)
        if len(runes) > 0 {
            runes[0] = unicode.ToUpper(runes[0])
            for j := 1; j < len(runes); j++ {
                runes[j] = unicode.ToLower(runes[j])
            }
            words[i] = string(runes)
        }
    }
    return strings.Join(words, " ")
}
```

---

#### Step 3: Add Tests / 테스트 추가

**Test Template / 테스트 템플릿**:

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    type
        expected type
    }{
        {"description1", input1, expected1},
        {"description2", input2, expected2},
        // ... more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FunctionName(tt.input)
            if result != tt.expected {
                t.Errorf("FunctionName(%v) = %v, want %v", tt.input, result, tt.expected)
            }
        })
    }
}
```

**Example Test / 예제 테스트**:

```go
func TestToTitleCase(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"lowercase", "hello world", "Hello World"},
        {"uppercase", "HELLO WORLD", "Hello World"},
        {"mixed", "hELLo WoRLd", "Hello World"},
        {"snake_case", "hello_world", "Hello World"},
        {"camelCase", "helloWorld", "Hello World"},
        {"empty", "", ""},
        {"unicode", "안녕 세계", "안녕 세계"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ToTitleCase(tt.input)
            if result != tt.expected {
                t.Errorf("ToTitleCase(%q) = %q, want %q", tt.input, result, tt.expected)
            }
        })
    }
}
```

---

#### Step 4: Update README / README 업데이트

Add the function to the appropriate category table in `stringutil/README.md`:

`stringutil/README.md`의 적절한 카테고리 테이블에 함수를 추가하세요:

```markdown
### Case Conversion / 케이스 변환

| Function / 함수 | Description / 설명 | Example / 예제 |
|-----------------|-------------------|----------------|
| ... | ... | ... |
| `ToTitleCase(s string) string` | Converts to Title Case / Title Case로 변환 | `ToTitleCase("hello world")` → `"Hello World"` |
```

---

#### Step 5: Update Documentation / 문서 업데이트

Add comprehensive documentation to `docs/stringutil/USER_MANUAL.md`:

`docs/stringutil/USER_MANUAL.md`에 포괄적인 문서를 추가하세요:

```markdown
#### ToTitleCase

Converts a string to Title Case format (first letter of each word capitalized).

문자열을 Title Case 형식으로 변환합니다 (각 단어의 첫 글자 대문자).

**Signature / 시그니처**:
func ToTitleCase(s string) string
```

**Example / 예제**:
```go
result := stringutil.ToTitleCase("hello world")
fmt.Println(result)  // Hello World
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예
```

---

#### Step 6: Run Tests / 테스트 실행

```bash
# Run all tests / 모든 테스트 실행
go test ./stringutil -v

# Run specific test / 특정 테스트 실행
go test ./stringutil -v -run TestToTitleCase

# Check coverage / 커버리지 확인
go test ./stringutil -cover
```

---

#### Step 7: Commit Changes / 변경사항 커밋

```bash
# Stage files / 파일 스테이징
git add stringutil/case.go stringutil/case_test.go stringutil/README.md docs/stringutil/USER_MANUAL.md

# Commit / 커밋
git commit -m "Feat: Add ToTitleCase function to stringutil

- Added ToTitleCase function for Title Case conversion
- Unicode-safe implementation using runes
- Added comprehensive tests with 7 test cases
- Updated README and USER_MANUAL documentation
"

# Push / 푸시
git push
```

---

## Testing Guide / 테스트 가이드

### Test Structure / 테스트 구조

All tests use **table-driven testing** for clarity and maintainability.

모든 테스트는 명확성과 유지보수성을 위해 **테이블 기반 테스트**를 사용합니다.

**Test File Naming / 테스트 파일 명명**:
- `case_test.go` - Case conversion tests
- `manipulation_test.go` - Manipulation tests
- `validation_test.go` - Validation tests
- `search_test.go` - Search & replace tests (future)
- `utils_test.go` - Utility tests (future)

---

### Running Tests / 테스트 실행

```bash
# Run all tests / 모든 테스트 실행
go test ./stringutil -v

# Run specific test / 특정 테스트 실행
go test ./stringutil -v -run TestToSnakeCase

# Run tests with coverage / 커버리지와 함께 테스트 실행
go test ./stringutil -cover

# Generate coverage report / 커버리지 리포트 생성
go test ./stringutil -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks / 벤치마크 실행
go test ./stringutil -bench=.

# Run specific benchmark / 특정 벤치마크 실행
go test ./stringutil -bench=BenchmarkToSnakeCase
```

---

### Writing Good Tests / 좋은 테스트 작성

#### 1. Test All Input Formats / 모든 입력 형식 테스트

```go
func TestToSnakeCase(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        // Different input formats / 다양한 입력 형식
        {"PascalCase", "pascal_case"},
        {"camelCase", "camel_case"},
        {"kebab-case", "kebab_case"},
        {"SCREAMING_SNAKE_CASE", "screaming_snake_case"},
        {"Mixed-Format_String", "mixed_format_string"},

        // Edge cases / 엣지 케이스
        {"", ""},
        {"A", "a"},
        {"ABC", "abc"},

        // Special cases / 특수 케이스
        {"HTTPServer", "http_server"},
        {"XMLParser", "xml_parser"},
    }
    // ...
}
```

---

#### 2. Test Unicode Support / 유니코드 지원 테스트

```go
func TestTruncateUnicode(t *testing.T) {
    tests := []struct {
        input    string
        length   int
        expected string
    }{
        // Korean / 한글
        {"안녕하세요", 3, "안녕하..."},

        // Japanese / 일본어
        {"こんにちは", 3, "こんに..."},

        // Chinese / 중국어
        {"你好世界", 2, "你好..."},

        // Emoji / 이모지
        {"😀😁😂😃", 2, "😀😁..."},

        // Mixed / 혼합
        {"Hello 世界", 8, "Hello 世界"},
    }
    // ...
}
```

---

#### 3. Test Edge Cases / 엣지 케이스 테스트

```go
func TestEdgeCases(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        // Empty / 빈 문자열
        {"empty string", "", ""},

        // Single character / 단일 문자
        {"single char", "a", "a"},

        // Whitespace / 공백
        {"only spaces", "   ", ""},
        {"tabs and newlines", "\t\n", ""},

        // Special characters / 특수 문자
        {"special chars", "!@#$%", "!@#$%"},

        // Very long string / 매우 긴 문자열
        {"long string", strings.Repeat("a", 10000), strings.Repeat("a", 10000)},
    }
    // ...
}
```

---

### Benchmark Tests / 벤치마크 테스트

**Template / 템플릿**:

```go
func BenchmarkFunctionName(b *testing.B) {
    input := "test input"

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        FunctionName(input)
    }
}
```

**Example / 예제**:

```go
func BenchmarkToSnakeCase(b *testing.B) {
    inputs := []string{
        "UserProfileData",
        "HTTPServer",
        "veryLongCamelCaseStringWithManyWords",
    }

    for _, input := range inputs {
        b.Run(input, func(b *testing.B) {
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                ToSnakeCase(input)
            }
        })
    }
}
```

---

## Performance / 성능

### Time Complexity / 시간 복잡도

| Function / 함수 | Time Complexity / 시간 복잡도 | Notes / 참고사항 |
|-----------------|-------------------------------|-----------------|
| ToSnakeCase | O(n) | n = string length |
| ToCamelCase | O(n) | n = string length |
| Truncate | O(n) | n = string length |
| Reverse | O(n) | n = string length |
| IsEmail | O(n) | Regex matching |
| ContainsAny | O(n*m) | n = string length, m = patterns |
| Map | O(n) | n = slice length |
| Filter | O(n) | n = slice length |

---

### Space Complexity / 공간 복잡도

| Function / 함수 | Space Complexity / 공간 복잡도 | Notes / 참고사항 |
|-----------------|-------------------------------|-----------------|
| ToSnakeCase | O(n) | New string created |
| Truncate | O(n) | Rune slice created |
| Map | O(n) | New slice created |
| Filter | O(n) | New slice (worst case) |

---

### Optimization Techniques / 최적화 기법

#### 1. Pre-allocate Slices / 슬라이스 사전 할당

```go
// Bad: Growing slice / 나쁨: 슬라이스 확장
func mapBad(strs []string, fn func(string) string) []string {
    var result []string  // ❌ No pre-allocation
    for _, s := range strs {
        result = append(result, fn(s))  // May reallocate multiple times
    }
    return result
}

// Good: Pre-allocated slice / 좋음: 사전 할당된 슬라이스
func Map(strs []string, fn func(string) string) []string {
    result := make([]string, len(strs))  // ✅ Pre-allocate
    for i, s := range strs {
        result[i] = fn(s)  // No reallocation needed
    }
    return result
}
```

---

#### 2. Use strings.Builder / strings.Builder 사용

```go
// Bad: String concatenation / 나쁨: 문자열 연결
func joinBad(strs []string, sep string) string {
    result := ""
    for i, s := range strs {
        if i > 0 {
            result += sep  // ❌ Creates new string each time!
        }
        result += s
    }
    return result
}

// Good: strings.Builder / 좋음: strings.Builder
func Join(strs []string, sep string) string {
    var builder strings.Builder  // ✅ Efficient buffer
    for i, s := range strs {
        if i > 0 {
            builder.WriteString(sep)
        }
        builder.WriteString(s)
    }
    return builder.String()
}
```

---

#### 3. Compile Regex Once / Regex 한 번 컴파일

```go
// Bad: Compile every time / 나쁨: 매번 컴파일
func IsEmailBad(s string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(s)  // ❌ Compiles on every call!
}

// Good: Package-level regex / 좋음: 패키지 레벨 regex
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsEmail(s string) bool {
    return emailRegex.MatchString(s)  // ✅ Already compiled!
}
```

---

#### 4. Early Returns / 조기 반환

```go
// Good: Early return for empty string / 좋음: 빈 문자열에 대한 조기 반환
func Truncate(s string, length int) string {
    if s == "" {  // ✅ Early return
        return ""
    }

    runes := []rune(s)
    if len(runes) <= length {  // ✅ Early return
        return s
    }

    return string(runes[:length]) + "..."
}
```

---

## Contributing Guidelines / 기여 가이드라인

### How to Contribute / 기여 방법

1. **Fork the Repository / 리포지토리 포크**
   ```bash
   # Fork on GitHub, then clone / GitHub에서 포크 후 클론
   git clone https://github.com/your-username/go-utils.git
   cd go-utils
   ```

2. **Create a Feature Branch / 기능 브랜치 생성**
   ```bash
   git checkout -b feature/add-title-case
   ```

3. **Make Changes / 변경사항 작성**
   - Add function to appropriate file
   - Write comprehensive tests
   - Update documentation

4. **Run Tests / 테스트 실행**
   ```bash
   go test ./... -v
   go test ./... -cover
   ```

5. **Commit Changes / 변경사항 커밋**
   ```bash
   git add .
   git commit -m "Feat: Add ToTitleCase function"
   ```

6. **Push and Create PR / 푸시 및 PR 생성**
   ```bash
   git push origin feature/add-title-case
   # Create Pull Request on GitHub
   ```

---

### Code Review Checklist / 코드 리뷰 체크리스트

**Functionality / 기능성**:
- [ ] Function works as expected
- [ ] Handles edge cases (empty string, special characters, etc.)
- [ ] Unicode-safe (uses runes where appropriate)
- [ ] No panics or crashes

**Tests / 테스트**:
- [ ] Comprehensive test coverage
- [ ] Table-driven tests
- [ ] Unicode test cases included
- [ ] Edge cases tested
- [ ] All tests pass

**Documentation / 문서**:
- [ ] Function has doc comment
- [ ] Doc comment includes example
- [ ] README updated
- [ ] USER_MANUAL updated (if user-facing)
- [ ] DEVELOPER_GUIDE updated (if architecture change)

**Code Quality / 코드 품질**:
- [ ] Follows Go conventions
- [ ] Clear variable names
- [ ] No unnecessary complexity
- [ ] Efficient implementation
- [ ] No external dependencies

**Performance / 성능**:
- [ ] No unnecessary allocations
- [ ] Efficient algorithms
- [ ] Benchmark tests included (for complex functions)

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

**Functions / 함수**:
- Use **PascalCase** for exported functions: `ToSnakeCase`, `IsEmail`
- Use **camelCase** for private functions: `splitIntoWords`
- Use descriptive verbs: `Convert`, `Validate`, `Check`, `Remove`

**Variables / 변수**:
- Use **short names** for local variables: `s`, `r`, `i`
- Use **descriptive names** for important variables: `currentWord`, `emailRegex`
- Avoid abbreviations unless common: `fn` (function), `sep` (separator)

**Constants / 상수**:
- Use **PascalCase** for exported constants
- Use **camelCase** for private constants
- Group related constants together

---

### Comment Style / 주석 스타일

**Function Comments / 함수 주석**:

```go
// FunctionName does X and returns Y.
// FunctionName는 X를 수행하고 Y를 반환합니다.
//
// Parameters:
// - param1: Description / 설명
// - param2: Description / 설명
//
// Returns:
// - Description / 설명
//
// Example / 예제:
//  result := stringutil.FunctionName("input")
//  fmt.Println(result)  // output
//
// Unicode Support / 유니코드 지원: Yes/No
func FunctionName(param1 type1, param2 type2) returnType {
    // ...
}
```

**Inline Comments / 인라인 주석**:

```go
// Good: Explain why, not what / 좋음: 무엇이 아닌 왜를 설명
// Use strings.Builder for efficiency (O(n) vs O(n²))
// 효율성을 위해 strings.Builder 사용 (O(n) vs O(n²))
var builder strings.Builder

// Bad: State the obvious / 나쁨: 명백한 것을 진술
// Create a builder / 빌더 생성
var builder strings.Builder
```

---

### Error Handling / 에러 처리

Most stringutil functions do not return errors because they are designed to be simple and predictable. However, when adding new functions that may fail, follow these guidelines:

대부분의 stringutil 함수는 간단하고 예측 가능하도록 설계되어 에러를 반환하지 않습니다. 그러나 실패할 수 있는 새 함수를 추가할 때는 다음 가이드라인을 따르세요:

```go
// Good: Return zero value on invalid input / 좋음: 잘못된 입력에 제로 값 반환
func CountWords(s string) int {
    if s == "" {
        return 0  // Zero value, not error
    }
    words := strings.Fields(s)
    return len(words)
}

// Good: Validation function returns bool / 좋음: 검증 함수는 bool 반환
func IsEmail(s string) bool {
    if s == "" {
        return false  // Invalid, not error
    }
    return emailRegex.MatchString(s)
}
```

---

### Import Organization / 임포트 구성

```go
// Standard library imports only / 표준 라이브러리 임포트만
import (
    "regexp"
    "strings"
    "unicode"
)
```

---

### Testing Best Practices / 테스팅 모범 사례

**Test Names / 테스트 이름**:
- Use `TestFunctionName` format
- Use descriptive subtest names

```go
func TestToSnakeCase(t *testing.T) {
    tests := []struct {
        name     string  // Descriptive name / 설명적 이름
        input    string
        expected string
    }{
        {"PascalCase input", "UserProfileData", "user_profile_data"},
        {"camelCase input", "userProfileData", "user_profile_data"},
        {"empty string", "", ""},
    }
    // ...
}
```

---

## Appendix: Complete Function Reference / 부록: 완전한 함수 참조

### Case Conversion (case.go)

```go
func ToSnakeCase(s string) string
func ToCamelCase(s string) string
func ToKebabCase(s string) string
func ToPascalCase(s string) string
func ToScreamingSnakeCase(s string) string

// Helper / 헬퍼
func splitIntoWords(s string) []string
```

### String Manipulation (manipulation.go)

```go
func Truncate(s string, length int) string
func TruncateWithSuffix(s string, length int, suffix string) string
func Reverse(s string) string
func Capitalize(s string) string
func CapitalizeFirst(s string) string
func RemoveDuplicates(s string) string
func RemoveSpaces(s string) string
func RemoveSpecialChars(s string) string
func Clean(s string) string
```

### Validation (validation.go)

```go
func IsEmail(s string) bool
func IsURL(s string) bool
func IsAlphanumeric(s string) bool
func IsNumeric(s string) bool
func IsAlpha(s string) bool
func IsBlank(s string) bool
func IsLower(s string) bool
func IsUpper(s string) bool
```

### Search & Replace (search.go)

```go
func ContainsAny(s string, substrs []string) bool
func ContainsAll(s string, substrs []string) bool
func StartsWithAny(s string, prefixes []string) bool
func EndsWithAny(s string, suffixes []string) bool
func ReplaceAll(s string, replacements map[string]string) string
func ReplaceIgnoreCase(s, old, new string) string
```

### Utilities (utils.go)

```go
func CountWords(s string) int
func CountOccurrences(s, substr string) int
func Join(strs []string, sep string) string
func Map(strs []string, fn func(string) string) []string
func Filter(strs []string, fn func(string) bool) []string
func PadLeft(s string, length int, pad string) string
func PadRight(s string, length int, pad string) string
func Lines(s string) []string
func Words(s string) []string
```

---

**End of Developer Guide / 개발자 가이드 끝**

For user documentation, see [USER_MANUAL.md](USER_MANUAL.md)

사용자 문서는 [USER_MANUAL.md](USER_MANUAL.md)를 참조하세요.
