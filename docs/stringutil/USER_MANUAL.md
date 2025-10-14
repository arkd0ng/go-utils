# Stringutil Package - User Manual
# Stringutil 패키지 - 사용자 매뉴얼

**Version / 버전**: v1.5.018+
**Package / 패키지**: `github.com/arkd0ng/go-utils/stringutil`
**Design Philosophy / 설계 철학**: "20 lines → 1 line" (Extreme Simplicity / 극도의 간결함)
**Function Count / 함수 개수**: 53 functions across 9 categories / 9개 카테고리에 걸친 53개 함수

> **Note**: This manual was initially written for v1.5.x (37 functions). The package has been expanded to 53 functions with additional categories:
> - Comparison (3 functions): EqualFold, HasPrefix, HasSuffix
> - Extended Manipulation (8 functions): Repeat, Substring, Left, Right, Insert, SwapCase, ToTitle, Slugify
> - Quote/Unquote (2 functions): Quote, Unquote
> - Unicode Operations (3 functions): RuneCount, Width, Normalize
>
> For complete API reference, see [stringutil/README.md](../../stringutil/README.md)
>
> **참고**: 이 매뉴얼은 처음에 v1.5.x (37개 함수)용으로 작성되었습니다. 패키지는 추가 카테고리와 함께 53개 함수로 확장되었습니다.
> 전체 API 참조는 [stringutil/README.md](../../stringutil/README.md)를 참조하세요.

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Configuration Reference / 설정 참조](#configuration-reference--설정-참조)
   - [Case Conversion / 케이스 변환](#case-conversion--케이스-변환)
   - [String Manipulation / 문자열 조작](#string-manipulation--문자열-조작)
   - [Validation / 유효성 검사](#validation--유효성-검사)
   - [Search & Replace / 검색 및 치환](#search--replace--검색-및-치환)
   - [Utilities / 유틸리티](#utilities--유틸리티)
5. [Usage Patterns / 사용 패턴](#usage-patterns--사용-패턴)
6. [Common Use Cases / 일반적인 사용 사례](#common-use-cases--일반적인-사용-사례)
7. [Best Practices / 모범 사례](#best-practices--모범-사례)
8. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)
9. [FAQ](#faq)

---

## Introduction / 소개

### What is Stringutil? / Stringutil이란?

The `stringutil` package provides extreme simplicity string utility functions for Go developers. The design philosophy is "20 lines → 1 line", dramatically reducing boilerplate code for common string operations.

`stringutil` 패키지는 Go 개발자를 위한 극도로 간단한 문자열 유틸리티 함수를 제공합니다. 설계 철학은 "20줄 → 1줄"로, 일반적인 문자열 작업에 대한 보일러플레이트 코드를 극적으로 줄입니다.

### Key Features / 주요 기능

- **37 utility functions** across 5 categories / 5개 카테고리에 걸친 **37개 유틸리티 함수**
- **Unicode-safe** - All functions use `[]rune` instead of byte operations / **유니코드 안전** - 모든 함수는 바이트 작업 대신 `[]rune` 사용
- **Zero external dependencies** - Standard library only / **외부 의존성 제로** - 표준 라이브러리만 사용
- **Bilingual documentation** - English and Korean / **이중 언어 문서** - 영문 및 한글
- **Practical validation** - Email/URL validation that works for 99% of use cases / **실용적인 검증** - 99%의 사용 사례에서 작동하는 이메일/URL 검증
- **Smart case conversion** - Handles PascalCase, camelCase, snake_case, kebab-case, SCREAMING_SNAKE_CASE / **스마트 케이스 변환** - PascalCase, camelCase, snake_case, kebab-case, SCREAMING_SNAKE_CASE 처리

### Use Cases / 사용 사례

Common scenarios where stringutil excels:

stringutil이 뛰어난 일반적인 시나리오:

1. **API Development** - Convert between different naming conventions (JSON camelCase ↔ Database snake_case)
2. **Data Validation** - Quick email, URL, alphanumeric validation
3. **Text Processing** - Truncate, clean, and manipulate user input
4. **String Transformation** - Map/filter operations on string slices
5. **Search & Replace** - Multi-pattern search and replacement

1. **API 개발** - 다양한 명명 규칙 간 변환 (JSON camelCase ↔ 데이터베이스 snake_case)
2. **데이터 검증** - 빠른 이메일, URL, 영숫자 검증
3. **텍스트 처리** - 사용자 입력 자르기, 정리, 조작
4. **문자열 변환** - 문자열 슬라이스에 대한 Map/Filter 작업
5. **검색 및 치환** - 다중 패턴 검색 및 치환

### Package Statistics / 패키지 통계

- **Total Functions / 전체 함수**: 37
- **Code Lines / 코드 라인**: ~714 lines
- **Test Coverage / 테스트 커버리지**: 9 tests (more coming)
- **External Dependencies / 외부 의존성**: 0
- **Supported Go Versions / 지원 Go 버전**: 1.18+

---

## Installation / 설치

### Prerequisites / 전제 조건

- Go 1.18 or higher / Go 1.18 이상
- No external dependencies required / 외부 의존성 필요 없음

### Install Package / 패키지 설치

```bash
go get github.com/arkd0ng/go-utils/stringutil
```

### Import in Your Code / 코드에 임포트

```go
import "github.com/arkd0ng/go-utils/stringutil"
```

### Verify Installation / 설치 확인

Create a simple test file:

간단한 테스트 파일 생성:

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func main() {
    // Test case conversion / 케이스 변환 테스트
    result := stringutil.ToSnakeCase("UserProfileData")
    fmt.Println(result) // Output: user_profile_data
}
```

Run the test:

테스트 실행:

```bash
go run main.go
```

---

## Quick Start / 빠른 시작

### Example 1: Case Conversion / 케이스 변환

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func main() {
    input := "UserProfileData"

    // Convert to different cases / 다양한 케이스로 변환
    fmt.Println(stringutil.ToSnakeCase(input))           // user_profile_data
    fmt.Println(stringutil.ToCamelCase(input))           // userProfileData
    fmt.Println(stringutil.ToKebabCase(input))           // user-profile-data
    fmt.Println(stringutil.ToPascalCase(input))          // UserProfileData
    fmt.Println(stringutil.ToScreamingSnakeCase(input))  // USER_PROFILE_DATA
}
```

### Example 2: String Manipulation / 문자열 조작

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func main() {
    // Truncate long text / 긴 텍스트 자르기
    text := "Hello World! This is a long sentence."
    fmt.Println(stringutil.Truncate(text, 20))  // Hello World! This i...

    // Reverse string / 문자열 뒤집기
    fmt.Println(stringutil.Reverse("hello"))    // olleh

    // Capitalize first letter / 첫 글자 대문자화
    fmt.Println(stringutil.Capitalize("hello world"))  // Hello world

    // Clean whitespace / 공백 정리
    fmt.Println(stringutil.Clean("  hello   world  "))  // hello world
}
```

### Example 3: Validation / 유효성 검사

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func main() {
    // Email validation / 이메일 검증
    fmt.Println(stringutil.IsEmail("user@example.com"))  // true
    fmt.Println(stringutil.IsEmail("invalid.email"))     // false

    // URL validation / URL 검증
    fmt.Println(stringutil.IsURL("https://example.com"))  // true
    fmt.Println(stringutil.IsURL("example.com"))          // false

    // Alphanumeric check / 영숫자 확인
    fmt.Println(stringutil.IsAlphanumeric("abc123"))      // true
    fmt.Println(stringutil.IsAlphanumeric("abc-123"))     // false
}
```

### Example 4: Search & Replace / 검색 및 치환

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func main() {
    // Check if string contains any of the substrings
    // 문자열이 부분 문자열 중 하나를 포함하는지 확인
    text := "https://example.com"
    fmt.Println(stringutil.ContainsAny(text, []string{"http://", "https://"}))  // true

    // Replace multiple patterns / 여러 패턴 치환
    replacements := map[string]string{
        "hello": "hi",
        "world": "universe",
    }
    fmt.Println(stringutil.ReplaceAll("hello world", replacements))  // hi universe
}
```

### Example 5: Utilities / 유틸리티

```go
package main

import (
    "fmt"
    "strings"
    "github.com/arkd0ng/go-utils/stringutil"
)

func main() {
    // Count words / 단어 개수 세기
    fmt.Println(stringutil.CountWords("hello world foo"))  // 3

    // Pad strings / 문자열 패딩
    fmt.Println(stringutil.PadLeft("5", 3, "0"))   // 005
    fmt.Println(stringutil.PadRight("5", 3, "0"))  // 500

    // Map transformation / Map 변환
    words := []string{"hello", "world"}
    upper := stringutil.Map(words, strings.ToUpper)
    fmt.Println(upper)  // [HELLO WORLD]

    // Filter strings / 문자열 필터링
    filtered := stringutil.Filter(words, func(s string) bool {
        return len(s) > 4
    })
    fmt.Println(filtered)  // [hello world]
}
```

---

## Configuration Reference / 설정 참조

### Case Conversion / 케이스 변환

#### ToSnakeCase

Converts a string to snake_case format.

문자열을 snake_case 형식으로 변환합니다.

**Signature / 시그니처**:
```go
func ToSnakeCase(s string) string
```

**Supported Input Formats / 지원 입력 형식**:
- PascalCase: `UserProfileData` → `user_profile_data`
- camelCase: `userProfileData` → `user_profile_data`
- kebab-case: `user-profile-data` → `user_profile_data`
- SCREAMING_SNAKE_CASE: `USER_PROFILE_DATA` → `user_profile_data`
- Mixed: `HTTPServer` → `http_server`

**Example / 예제**:
```go
result := stringutil.ToSnakeCase("UserProfileData")
fmt.Println(result)  // user_profile_data
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### ToCamelCase

Converts a string to camelCase format (first letter lowercase).

문자열을 camelCase 형식으로 변환합니다 (첫 글자 소문자).

**Signature / 시그니처**:
```go
func ToCamelCase(s string) string
```

**Supported Input Formats / 지원 입력 형식**:
- snake_case: `user_profile_data` → `userProfileData`
- kebab-case: `user-profile-data` → `userProfileData`
- PascalCase: `UserProfileData` → `userProfileData`
- SCREAMING_SNAKE_CASE: `USER_PROFILE_DATA` → `userProfileData`

**Example / 예제**:
```go
result := stringutil.ToCamelCase("user_profile_data")
fmt.Println(result)  // userProfileData
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### ToKebabCase

Converts a string to kebab-case format.

문자열을 kebab-case 형식으로 변환합니다.

**Signature / 시그니처**:
```go
func ToKebabCase(s string) string
```

**Supported Input Formats / 지원 입력 형식**:
- PascalCase: `UserProfileData` → `user-profile-data`
- camelCase: `userProfileData` → `user-profile-data`
- snake_case: `user_profile_data` → `user-profile-data`
- SCREAMING_SNAKE_CASE: `USER_PROFILE_DATA` → `user-profile-data`

**Example / 예제**:
```go
result := stringutil.ToKebabCase("UserProfileData")
fmt.Println(result)  // user-profile-data
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### ToPascalCase

Converts a string to PascalCase format (first letter uppercase).

문자열을 PascalCase 형식으로 변환합니다 (첫 글자 대문자).

**Signature / 시그니처**:
```go
func ToPascalCase(s string) string
```

**Supported Input Formats / 지원 입력 형식**:
- snake_case: `user_profile_data` → `UserProfileData`
- kebab-case: `user-profile-data` → `UserProfileData`
- camelCase: `userProfileData` → `UserProfileData`
- SCREAMING_SNAKE_CASE: `USER_PROFILE_DATA` → `UserProfileData`

**Example / 예제**:
```go
result := stringutil.ToPascalCase("user_profile_data")
fmt.Println(result)  // UserProfileData
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### ToScreamingSnakeCase

Converts a string to SCREAMING_SNAKE_CASE format (all uppercase with underscores).

문자열을 SCREAMING_SNAKE_CASE 형식으로 변환합니다 (모두 대문자, 밑줄 포함).

**Signature / 시그니처**:
```go
func ToScreamingSnakeCase(s string) string
```

**Supported Input Formats / 지원 입력 형식**:
- PascalCase: `UserProfileData` → `USER_PROFILE_DATA`
- camelCase: `userProfileData` → `USER_PROFILE_DATA`
- snake_case: `user_profile_data` → `USER_PROFILE_DATA`
- kebab-case: `user-profile-data` → `USER_PROFILE_DATA`

**Example / 예제**:
```go
result := stringutil.ToScreamingSnakeCase("UserProfileData")
fmt.Println(result)  // USER_PROFILE_DATA
```

**Use Case / 사용 사례**: Environment variables, constants / 환경 변수, 상수

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

### String Manipulation / 문자열 조작

#### Truncate

Truncates a string to the specified length and appends "..." if truncated.

문자열을 지정된 길이로 자르고 잘린 경우 "..."를 추가합니다.

**Signature / 시그니처**:
```go
func Truncate(s string, length int) string
```

**Parameters / 매개변수**:
- `s string`: Input string / 입력 문자열
- `length int`: Maximum length (in runes, not bytes) / 최대 길이 (바이트가 아닌 rune)

**Example / 예제**:
```go
result := stringutil.Truncate("Hello World! This is a long sentence.", 20)
fmt.Println(result)  // Hello World! This i...

// Unicode support / 유니코드 지원
result2 := stringutil.Truncate("안녕하세요, 반갑습니다!", 5)
fmt.Println(result2)  // 안녕하세요...
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### TruncateWithSuffix

Truncates a string to the specified length and appends a custom suffix.

문자열을 지정된 길이로 자르고 사용자 정의 접미사를 추가합니다.

**Signature / 시그니처**:
```go
func TruncateWithSuffix(s string, length int, suffix string) string
```

**Parameters / 매개변수**:
- `s string`: Input string / 입력 문자열
- `length int`: Maximum length (in runes) / 최대 길이 (rune)
- `suffix string`: Custom suffix / 사용자 정의 접미사

**Example / 예제**:
```go
result := stringutil.TruncateWithSuffix("Hello World", 8, "…")
fmt.Println(result)  // Hello Wo…

result2 := stringutil.TruncateWithSuffix("Hello World", 8, " [more]")
fmt.Println(result2)  // Hello Wo [more]
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### Reverse

Reverses a string (Unicode-safe).

문자열을 뒤집습니다 (유니코드 안전).

**Signature / 시그니처**:
```go
func Reverse(s string) string
```

**Example / 예제**:
```go
result := stringutil.Reverse("hello")
fmt.Println(result)  // olleh

// Unicode support / 유니코드 지원
result2 := stringutil.Reverse("안녕하세요")
fmt.Println(result2)  // 요세하녕안
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### Capitalize

Capitalizes the first letter of each word.

각 단어의 첫 글자를 대문자로 만듭니다.

**Signature / 시그니처**:
```go
func Capitalize(s string) string
```

**Example / 예제**:
```go
result := stringutil.Capitalize("hello world")
fmt.Println(result)  // Hello World

result2 := stringutil.Capitalize("go programming language")
fmt.Println(result2)  // Go Programming Language
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### CapitalizeFirst

Capitalizes only the first letter of the string.

문자열의 첫 글자만 대문자로 만듭니다.

**Signature / 시그니처**:
```go
func CapitalizeFirst(s string) string
```

**Example / 예제**:
```go
result := stringutil.CapitalizeFirst("hello world")
fmt.Println(result)  // Hello world

result2 := stringutil.CapitalizeFirst("the quick brown fox")
fmt.Println(result2)  // The quick brown fox
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### RemoveDuplicates

Removes duplicate consecutive characters.

중복된 연속 문자를 제거합니다.

**Signature / 시그니처**:
```go
func RemoveDuplicates(s string) string
```

**Example / 예제**:
```go
result := stringutil.RemoveDuplicates("heeelllooo")
fmt.Println(result)  // helo

result2 := stringutil.RemoveDuplicates("aaabbbccc")
fmt.Println(result2)  // abc
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### RemoveSpaces

Removes all whitespace characters.

모든 공백 문자를 제거합니다.

**Signature / 시그니처**:
```go
func RemoveSpaces(s string) string
```

**Example / 예제**:
```go
result := stringutil.RemoveSpaces("h e l l o")
fmt.Println(result)  // hello

result2 := stringutil.RemoveSpaces("  hello   world  ")
fmt.Println(result2)  // helloworld
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### RemoveSpecialChars

Removes all non-alphanumeric characters (keeps letters, digits, and spaces).

모든 비영숫자 문자를 제거합니다 (문자, 숫자, 공백 유지).

**Signature / 시그니처**:
```go
func RemoveSpecialChars(s string) string
```

**Example / 예제**:
```go
result := stringutil.RemoveSpecialChars("hello@world!123#")
fmt.Println(result)  // helloworld123

result2 := stringutil.RemoveSpecialChars("user@example.com")
fmt.Println(result2)  // userexamplecom
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### Clean

Trims whitespace and removes duplicate spaces.

공백을 제거하고 중복 공백을 제거합니다.

**Signature / 시그니처**:
```go
func Clean(s string) string
```

**Example / 예제**:
```go
result := stringutil.Clean("  hello   world  ")
fmt.Println(result)  // hello world

result2 := stringutil.Clean("\t\nhello\t\nworld")
fmt.Println(result2)  // hello world
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

### Validation / 유효성 검사

#### IsEmail

Validates if a string is a valid email address (practical validation, not RFC 5322).

문자열이 유효한 이메일 주소인지 검증합니다 (실용적 검증, RFC 5322 아님).

**Signature / 시그니처**:
```go
func IsEmail(s string) bool
```

**Validation Pattern / 검증 패턴**:
```
[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}
```

**Example / 예제**:
```go
fmt.Println(stringutil.IsEmail("user@example.com"))      // true
fmt.Println(stringutil.IsEmail("user+tag@example.com"))  // true
fmt.Println(stringutil.IsEmail("invalid.email"))         // false
fmt.Println(stringutil.IsEmail("@example.com"))          // false
fmt.Println(stringutil.IsEmail("user@"))                 // false
```

**Coverage / 커버리지**: Works for 99% of email addresses / 99%의 이메일 주소에서 작동

---

#### IsURL

Validates if a string is a valid URL (starts with http:// or https://).

문자열이 유효한 URL인지 검증합니다 (http:// 또는 https://로 시작).

**Signature / 시그니처**:
```go
func IsURL(s string) bool
```

**Example / 예제**:
```go
fmt.Println(stringutil.IsURL("https://example.com"))      // true
fmt.Println(stringutil.IsURL("http://example.com/path"))  // true
fmt.Println(stringutil.IsURL("example.com"))              // false
fmt.Println(stringutil.IsURL("ftp://example.com"))        // false
```

**Note / 참고**: Simple prefix check, not full URL validation / 단순 접두사 확인, 전체 URL 검증 아님

---

#### IsAlphanumeric

Checks if a string contains only letters and digits.

문자열이 문자와 숫자만 포함하는지 확인합니다.

**Signature / 시그니처**:
```go
func IsAlphanumeric(s string) bool
```

**Example / 예제**:
```go
fmt.Println(stringutil.IsAlphanumeric("abc123"))    // true
fmt.Println(stringutil.IsAlphanumeric("ABC"))       // true
fmt.Println(stringutil.IsAlphanumeric("123"))       // true
fmt.Println(stringutil.IsAlphanumeric("abc-123"))   // false
fmt.Println(stringutil.IsAlphanumeric("abc 123"))   // false
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### IsNumeric

Checks if a string contains only digits.

문자열이 숫자만 포함하는지 확인합니다.

**Signature / 시그니처**:
```go
func IsNumeric(s string) bool
```

**Example / 예제**:
```go
fmt.Println(stringutil.IsNumeric("12345"))    // true
fmt.Println(stringutil.IsNumeric("123.45"))   // false
fmt.Println(stringutil.IsNumeric("123a"))     // false
fmt.Println(stringutil.IsNumeric(""))         // false
```

**Unicode Support / 유니코드 지원**: ✅ Yes (supports Unicode digits) / 예 (유니코드 숫자 지원)

---

#### IsAlpha

Checks if a string contains only letters.

문자열이 문자만 포함하는지 확인합니다.

**Signature / 시그니처**:
```go
func IsAlpha(s string) bool
```

**Example / 예제**:
```go
fmt.Println(stringutil.IsAlpha("hello"))    // true
fmt.Println(stringutil.IsAlpha("Hello"))    // true
fmt.Println(stringutil.IsAlpha("hello123")) // false
fmt.Println(stringutil.IsAlpha(""))         // false
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### IsBlank

Checks if a string is empty or contains only whitespace.

문자열이 비어 있거나 공백만 포함하는지 확인합니다.

**Signature / 시그니처**:
```go
func IsBlank(s string) bool
```

**Example / 예제**:
```go
fmt.Println(stringutil.IsBlank(""))          // true
fmt.Println(stringutil.IsBlank("   "))       // true
fmt.Println(stringutil.IsBlank("\t\n"))      // true
fmt.Println(stringutil.IsBlank("hello"))     // false
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### IsLower

Checks if a string contains only lowercase letters.

문자열이 소문자만 포함하는지 확인합니다.

**Signature / 시그니처**:
```go
func IsLower(s string) bool
```

**Example / 예제**:
```go
fmt.Println(stringutil.IsLower("hello"))    // true
fmt.Println(stringutil.IsLower("Hello"))    // false
fmt.Println(stringutil.IsLower("hello123")) // false
fmt.Println(stringutil.IsLower(""))         // false
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### IsUpper

Checks if a string contains only uppercase letters.

문자열이 대문자만 포함하는지 확인합니다.

**Signature / 시그니처**:
```go
func IsUpper(s string) bool
```

**Example / 예제**:
```go
fmt.Println(stringutil.IsUpper("HELLO"))    // true
fmt.Println(stringutil.IsUpper("Hello"))    // false
fmt.Println(stringutil.IsUpper("HELLO123")) // false
fmt.Println(stringutil.IsUpper(""))         // false
```

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

### Search & Replace / 검색 및 치환

#### ContainsAny

Checks if a string contains any of the specified substrings.

문자열이 지정된 부분 문자열 중 하나를 포함하는지 확인합니다.

**Signature / 시그니처**:
```go
func ContainsAny(s string, substrs []string) bool
```

**Example / 예제**:
```go
text := "https://example.com"
result := stringutil.ContainsAny(text, []string{"http://", "https://"})
fmt.Println(result)  // true

text2 := "hello world"
result2 := stringutil.ContainsAny(text2, []string{"foo", "bar"})
fmt.Println(result2)  // false
```

**Use Case / 사용 사례**: Protocol detection, keyword matching / 프로토콜 탐지, 키워드 매칭

---

#### ContainsAll

Checks if a string contains all of the specified substrings.

문자열이 지정된 모든 부분 문자열을 포함하는지 확인합니다.

**Signature / 시그니처**:
```go
func ContainsAll(s string, substrs []string) bool
```

**Example / 예제**:
```go
text := "hello world foo bar"
result := stringutil.ContainsAll(text, []string{"hello", "world"})
fmt.Println(result)  // true

result2 := stringutil.ContainsAll(text, []string{"hello", "baz"})
fmt.Println(result2)  // false
```

**Use Case / 사용 사례**: Multi-keyword validation / 다중 키워드 검증

---

#### StartsWithAny

Checks if a string starts with any of the specified prefixes.

문자열이 지정된 접두사 중 하나로 시작하는지 확인합니다.

**Signature / 시그니처**:
```go
func StartsWithAny(s string, prefixes []string) bool
```

**Example / 예제**:
```go
url := "https://example.com"
result := stringutil.StartsWithAny(url, []string{"http://", "https://"})
fmt.Println(result)  // true

text := "hello world"
result2 := stringutil.StartsWithAny(text, []string{"foo", "bar"})
fmt.Println(result2)  // false
```

**Use Case / 사용 사례**: Protocol validation, prefix matching / 프로토콜 검증, 접두사 매칭

---

#### EndsWithAny

Checks if a string ends with any of the specified suffixes.

문자열이 지정된 접미사 중 하나로 끝나는지 확인합니다.

**Signature / 시그니처**:
```go
func EndsWithAny(s string, suffixes []string) bool
```

**Example / 예제**:
```go
filename := "document.pdf"
result := stringutil.EndsWithAny(filename, []string{".pdf", ".doc", ".txt"})
fmt.Println(result)  // true

filename2 := "image.jpg"
result2 := stringutil.EndsWithAny(filename2, []string{".pdf", ".doc"})
fmt.Println(result2)  // false
```

**Use Case / 사용 사례**: File extension validation / 파일 확장자 검증

---

#### ReplaceAll

Replaces all occurrences of multiple patterns with their corresponding replacements.

여러 패턴의 모든 발생을 해당 치환으로 바꿉니다.

**Signature / 시그니처**:
```go
func ReplaceAll(s string, replacements map[string]string) string
```

**Example / 예제**:
```go
replacements := map[string]string{
    "hello": "hi",
    "world": "universe",
    "foo":   "bar",
}
result := stringutil.ReplaceAll("hello world foo", replacements)
fmt.Println(result)  // hi universe bar
```

**Use Case / 사용 사례**: Template replacement, text sanitization / 템플릿 치환, 텍스트 정제

---

#### ReplaceIgnoreCase

Replaces all occurrences of a substring ignoring case.

대소문자를 무시하고 부분 문자열의 모든 발생을 바꿉니다.

**Signature / 시그니처**:
```go
func ReplaceIgnoreCase(s, old, new string) string
```

**Example / 예제**:
```go
result := stringutil.ReplaceIgnoreCase("Hello WORLD hello", "hello", "hi")
fmt.Println(result)  // hi WORLD hi

result2 := stringutil.ReplaceIgnoreCase("ABC abc AbC", "abc", "xyz")
fmt.Println(result2)  // xyz xyz xyz
```

**Use Case / 사용 사례**: Case-insensitive text replacement / 대소문자 무시 텍스트 치환

---

### Utilities / 유틸리티

#### CountWords

Counts the number of words in a string.

문자열의 단어 수를 셉니다.

**Signature / 시그니처**:
```go
func CountWords(s string) int
```

**Example / 예제**:
```go
fmt.Println(stringutil.CountWords("hello world foo"))  // 3
fmt.Println(stringutil.CountWords("  hello   world  "))  // 2
fmt.Println(stringutil.CountWords(""))  // 0
```

**Note / 참고**: Uses `strings.Fields()` which splits on whitespace / 공백으로 분할하는 `strings.Fields()` 사용

---

#### CountOccurrences

Counts the number of occurrences of a substring.

부분 문자열의 발생 횟수를 셉니다.

**Signature / 시그니처**:
```go
func CountOccurrences(s, substr string) int
```

**Example / 예제**:
```go
fmt.Println(stringutil.CountOccurrences("hello world hello", "hello"))  // 2
fmt.Println(stringutil.CountOccurrences("aaabbbccc", "aa"))  // 1
fmt.Println(stringutil.CountOccurrences("hello world", "foo"))  // 0
```

**Use Case / 사용 사례**: Text analysis, pattern frequency / 텍스트 분석, 패턴 빈도

---

#### Join

Joins a slice of strings with a separator.

구분자로 문자열 슬라이스를 결합합니다.

**Signature / 시그니처**:
```go
func Join(strs []string, sep string) string
```

**Example / 예제**:
```go
words := []string{"hello", "world", "foo"}
result := stringutil.Join(words, ", ")
fmt.Println(result)  // hello, world, foo

result2 := stringutil.Join(words, " | ")
fmt.Println(result2)  // hello | world | foo
```

**Note / 참고**: Wrapper around `strings.Join()` / `strings.Join()` 래퍼

---

#### Map

Applies a function to each string in a slice.

슬라이스의 각 문자열에 함수를 적용합니다.

**Signature / 시그니처**:
```go
func Map(strs []string, fn func(string) string) []string
```

**Example / 예제**:
```go
words := []string{"hello", "world", "foo"}
upper := stringutil.Map(words, strings.ToUpper)
fmt.Println(upper)  // [HELLO WORLD FOO]

// Custom transformation / 사용자 정의 변환
prefix := stringutil.Map(words, func(s string) string {
    return "prefix_" + s
})
fmt.Println(prefix)  // [prefix_hello prefix_world prefix_foo]
```

**Use Case / 사용 사례**: Batch string transformation / 일괄 문자열 변환

---

#### Filter

Filters strings in a slice based on a predicate function.

조건 함수를 기반으로 슬라이스의 문자열을 필터링합니다.

**Signature / 시그니처**:
```go
func Filter(strs []string, fn func(string) bool) []string
```

**Example / 예제**:
```go
words := []string{"hello", "world", "foo", "a", "bar"}

// Filter by length / 길이로 필터링
long := stringutil.Filter(words, func(s string) bool {
    return len(s) > 3
})
fmt.Println(long)  // [hello world]

// Filter by prefix / 접두사로 필터링
withF := stringutil.Filter(words, func(s string) bool {
    return strings.HasPrefix(s, "f")
})
fmt.Println(withF)  // [foo]
```

**Use Case / 사용 사례**: Selective string processing / 선택적 문자열 처리

---

#### PadLeft

Pads a string on the left to reach the specified length.

지정된 길이에 도달하도록 문자열 왼쪽을 패딩합니다.

**Signature / 시그니처**:
```go
func PadLeft(s string, length int, pad string) string
```

**Example / 예제**:
```go
fmt.Println(stringutil.PadLeft("5", 3, "0"))    // 005
fmt.Println(stringutil.PadLeft("42", 5, "0"))   // 00042
fmt.Println(stringutil.PadLeft("hello", 10, " "))  // "     hello"
```

**Use Case / 사용 사례**: Number formatting, text alignment / 숫자 형식화, 텍스트 정렬

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### PadRight

Pads a string on the right to reach the specified length.

지정된 길이에 도달하도록 문자열 오른쪽을 패딩합니다.

**Signature / 시그니처**:
```go
func PadRight(s string, length int, pad string) string
```

**Example / 예제**:
```go
fmt.Println(stringutil.PadRight("5", 3, "0"))    // 500
fmt.Println(stringutil.PadRight("42", 5, "0"))   // 42000
fmt.Println(stringutil.PadRight("hello", 10, " "))  // "hello     "
```

**Use Case / 사용 사례**: Text formatting, table columns / 텍스트 형식화, 테이블 열

**Unicode Support / 유니코드 지원**: ✅ Yes / 예

---

#### Lines

Splits a string into lines.

문자열을 줄로 분할합니다.

**Signature / 시그니처**:
```go
func Lines(s string) []string
```

**Example / 예제**:
```go
text := "line1\nline2\nline3"
lines := stringutil.Lines(text)
fmt.Println(lines)  // [line1 line2 line3]

text2 := "hello\r\nworld\r\nfoo"
lines2 := stringutil.Lines(text2)
fmt.Println(lines2)  // [hello world foo]
```

**Note / 참고**: Handles both `\n` and `\r\n` / `\n`과 `\r\n` 모두 처리

---

#### Words

Splits a string into words.

문자열을 단어로 분할합니다.

**Signature / 시그니처**:
```go
func Words(s string) []string
```

**Example / 예제**:
```go
text := "hello world foo"
words := stringutil.Words(text)
fmt.Println(words)  // [hello world foo]

text2 := "  hello   world  "
words2 := stringutil.Words(text2)
fmt.Println(words2)  // [hello world]
```

**Note / 참고**: Uses `strings.Fields()` which splits on whitespace / 공백으로 분할하는 `strings.Fields()` 사용

---

## Usage Patterns / 사용 패턴

### Pattern 1: API Request/Response Transformation

Convert between JSON (camelCase) and database (snake_case) naming conventions.

JSON (camelCase)과 데이터베이스 (snake_case) 명명 규칙 간 변환.

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

type JSONRequest struct {
    UserName  string
    FirstName string
    LastName  string
    EmailAddress string
}

type DatabaseModel struct {
    user_name     string
    first_name    string
    last_name     string
    email_address string
}

func main() {
    // Convert JSON field names to database column names
    // JSON 필드 이름을 데이터베이스 컬럼 이름으로 변환
    fields := []string{"UserName", "FirstName", "LastName", "EmailAddress"}

    dbColumns := stringutil.Map(fields, stringutil.ToSnakeCase)
    fmt.Println(dbColumns)
    // [user_name first_name last_name email_address]

    // Convert database columns to camelCase for JSON response
    // 데이터베이스 컬럼을 JSON 응답용 camelCase로 변환
    jsonFields := stringutil.Map(dbColumns, stringutil.ToCamelCase)
    fmt.Println(jsonFields)
    // [userName firstName lastName emailAddress]
}
```

---

### Pattern 2: User Input Validation and Sanitization

Validate and clean user input before processing.

처리 전에 사용자 입력을 검증하고 정리합니다.

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func ValidateUserInput(email, website, username string) error {
    // Validate email / 이메일 검증
    if !stringutil.IsEmail(email) {
        return fmt.Errorf("invalid email: %s", email)
    }

    // Validate website / 웹사이트 검증
    if website != "" && !stringutil.IsURL(website) {
        return fmt.Errorf("invalid website: %s", website)
    }

    // Validate username (alphanumeric only) / 사용자 이름 검증 (영숫자만)
    cleaned := stringutil.RemoveSpaces(username)
    if !stringutil.IsAlphanumeric(cleaned) {
        return fmt.Errorf("username must be alphanumeric: %s", username)
    }

    return nil
}

func main() {
    // Valid input / 유효한 입력
    err := ValidateUserInput("user@example.com", "https://example.com", "john123")
    fmt.Println(err)  // nil

    // Invalid email / 잘못된 이메일
    err = ValidateUserInput("invalid.email", "https://example.com", "john123")
    fmt.Println(err)  // invalid email: invalid.email

    // Invalid username / 잘못된 사용자 이름
    err = ValidateUserInput("user@example.com", "https://example.com", "john@123")
    fmt.Println(err)  // username must be alphanumeric: john@123
}
```

---

### Pattern 3: Text Truncation for Display

Truncate long text for UI display with proper Unicode handling.

UI 표시를 위해 적절한 유니코드 처리로 긴 텍스트를 자릅니다.

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

type Article struct {
    Title       string
    Description string
    Content     string
}

func (a *Article) DisplaySummary(maxTitleLen, maxDescLen int) {
    // Truncate title / 제목 자르기
    title := stringutil.Truncate(a.Title, maxTitleLen)

    // Truncate description / 설명 자르기
    description := stringutil.Truncate(a.Description, maxDescLen)

    fmt.Printf("Title: %s\n", title)
    fmt.Printf("Description: %s\n", description)
}

func main() {
    article := Article{
        Title:       "Introduction to Go Programming Language and Its Ecosystem",
        Description: "This article covers the basics of Go programming, including syntax, concurrency, and best practices for building scalable applications.",
        Content:     "...",
    }

    // Display with truncation / 자르기로 표시
    article.DisplaySummary(40, 60)
    // Output:
    // Title: Introduction to Go Programming Langua...
    // Description: This article covers the basics of Go programming, incl...
}
```

---

### Pattern 4: Multi-Pattern Search and Replace

Replace multiple patterns in a single pass.

단일 패스에서 여러 패턴을 치환합니다.

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func SanitizeHTML(text string) string {
    // Replace HTML entities / HTML 엔터티 치환
    replacements := map[string]string{
        "&":  "&amp;",
        "<":  "&lt;",
        ">":  "&gt;",
        "\"": "&quot;",
        "'":  "&#39;",
    }

    return stringutil.ReplaceAll(text, replacements)
}

func main() {
    input := "<script>alert('XSS')</script>"
    sanitized := SanitizeHTML(input)
    fmt.Println(sanitized)
    // Output: &lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;
}
```

---

### Pattern 5: Batch String Transformation

Process multiple strings with the same transformation.

동일한 변환으로 여러 문자열을 처리합니다.

```go
package main

import (
    "fmt"
    "strings"
    "github.com/arkd0ng/go-utils/stringutil"
)

func main() {
    // Convert all to uppercase / 모두 대문자로 변환
    names := []string{"alice", "bob", "charlie"}
    upper := stringutil.Map(names, strings.ToUpper)
    fmt.Println(upper)  // [ALICE BOB CHARLIE]

    // Convert all to snake_case / 모두 snake_case로 변환
    camelNames := []string{"firstName", "lastName", "emailAddress"}
    snakeNames := stringutil.Map(camelNames, stringutil.ToSnakeCase)
    fmt.Println(snakeNames)  // [first_name last_name email_address]

    // Filter and transform / 필터링 및 변환
    words := []string{"hello", "world", "a", "foo", "bar"}
    long := stringutil.Filter(words, func(s string) bool {
        return len(s) > 3
    })
    longUpper := stringutil.Map(long, strings.ToUpper)
    fmt.Println(longUpper)  // [HELLO WORLD]
}
```

---

### Pattern 6: Environment Variable Conversion

Convert between different case formats for environment variables.

환경 변수를 위한 다양한 케이스 형식 간 변환.

```go
package main

import (
    "fmt"
    "os"
    "github.com/arkd0ng/go-utils/stringutil"
)

func main() {
    // Configuration keys in different formats
    // 다양한 형식의 설정 키
    configKeys := []string{
        "databaseUrl",
        "apiTimeout",
        "maxConnections",
    }

    // Convert to SCREAMING_SNAKE_CASE for environment variables
    // 환경 변수를 위해 SCREAMING_SNAKE_CASE로 변환
    envVars := stringutil.Map(configKeys, stringutil.ToScreamingSnakeCase)
    fmt.Println(envVars)
    // [DATABASE_URL API_TIMEOUT MAX_CONNECTIONS]

    // Set environment variables / 환경 변수 설정
    for _, envVar := range envVars {
        os.Setenv(envVar, "example_value")
    }
}
```

---

### Pattern 7: File Extension Validation

Validate file extensions using suffix matching.

접미사 매칭을 사용한 파일 확장자 검증.

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func ValidateImageFile(filename string) bool {
    allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
    return stringutil.EndsWithAny(filename, allowedExtensions)
}

func ValidateDocumentFile(filename string) bool {
    allowedExtensions := []string{".pdf", ".doc", ".docx", ".txt", ".rtf"}
    return stringutil.EndsWithAny(filename, allowedExtensions)
}

func main() {
    // Image validation / 이미지 검증
    fmt.Println(ValidateImageFile("photo.jpg"))    // true
    fmt.Println(ValidateImageFile("document.pdf"))  // false

    // Document validation / 문서 검증
    fmt.Println(ValidateDocumentFile("report.pdf"))  // true
    fmt.Println(ValidateDocumentFile("image.png"))   // false
}
```

---

### Pattern 8: Text Cleaning and Normalization

Clean and normalize user-generated text.

사용자 생성 텍스트를 정리하고 정규화합니다.

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func NormalizeComment(comment string) string {
    // Remove leading/trailing whitespace and duplicate spaces
    // 앞뒤 공백 및 중복 공백 제거
    cleaned := stringutil.Clean(comment)

    // Truncate if too long / 너무 길면 자르기
    if len([]rune(cleaned)) > 200 {
        cleaned = stringutil.Truncate(cleaned, 200)
    }

    // Capitalize first letter / 첫 글자 대문자화
    cleaned = stringutil.CapitalizeFirst(cleaned)

    return cleaned
}

func main() {
    comment := "  hello   world!   this is a   comment   "
    normalized := NormalizeComment(comment)
    fmt.Println(normalized)  // Hello world! this is a comment
}
```

---

### Pattern 9: Protocol Detection

Detect URL protocols using prefix matching.

접두사 매칭을 사용한 URL 프로토콜 탐지.

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func IsSecureProtocol(url string) bool {
    return stringutil.StartsWithAny(url, []string{"https://", "ftps://", "sftp://"})
}

func GetProtocol(url string) string {
    protocols := []string{"https://", "http://", "ftp://", "ws://", "wss://"}

    for _, protocol := range protocols {
        if stringutil.StartsWithAny(url, []string{protocol}) {
            return protocol
        }
    }

    return "unknown"
}

func main() {
    url1 := "https://example.com"
    fmt.Println(IsSecureProtocol(url1))  // true
    fmt.Println(GetProtocol(url1))       // https://

    url2 := "http://example.com"
    fmt.Println(IsSecureProtocol(url2))  // false
    fmt.Println(GetProtocol(url2))       // http://
}
```

---

### Pattern 10: Word Count and Text Analysis

Analyze text statistics.

텍스트 통계를 분석합니다.

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func AnalyzeText(text string) {
    // Count words / 단어 개수 세기
    wordCount := stringutil.CountWords(text)
    fmt.Printf("Word count: %d\n", wordCount)

    // Split into words / 단어로 분할
    words := stringutil.Words(text)
    fmt.Printf("Words: %v\n", words)

    // Count occurrences of specific word / 특정 단어의 발생 횟수 세기
    helloCount := stringutil.CountOccurrences(text, "hello")
    fmt.Printf("'hello' appears: %d times\n", helloCount)

    // Filter long words / 긴 단어 필터링
    longWords := stringutil.Filter(words, func(s string) bool {
        return len(s) > 5
    })
    fmt.Printf("Long words (>5 chars): %v\n", longWords)
}

func main() {
    text := "hello world hello universe wonderful programming"
    AnalyzeText(text)
    // Output:
    // Word count: 6
    // Words: [hello world hello universe wonderful programming]
    // 'hello' appears: 2 times
    // Long words (>5 chars): [universe wonderful programming]
}
```

---

## Common Use Cases / 일반적인 사용 사례

### Use Case 1: REST API Field Name Conversion

**Scenario / 시나리오**: Convert JSON field names (camelCase) to database column names (snake_case) and vice versa.

JSON 필드 이름 (camelCase)을 데이터베이스 컬럼 이름 (snake_case)으로 변환하고 그 반대로 변환합니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

// JSON request from client / 클라이언트로부터의 JSON 요청
type UserRequest struct {
    FirstName    string `json:"firstName"`
    LastName     string `json:"lastName"`
    EmailAddress string `json:"emailAddress"`
    PhoneNumber  string `json:"phoneNumber"`
}

// Database model / 데이터베이스 모델
type UserModel struct {
    FirstName    string `db:"first_name"`
    LastName     string `db:"last_name"`
    EmailAddress string `db:"email_address"`
    PhoneNumber  string `db:"phone_number"`
}

func ConvertJSONToDBColumns(jsonFields []string) []string {
    return stringutil.Map(jsonFields, stringutil.ToSnakeCase)
}

func ConvertDBToJSONFields(dbColumns []string) []string {
    return stringutil.Map(dbColumns, stringutil.ToCamelCase)
}

func main() {
    // JSON field names / JSON 필드 이름
    jsonFields := []string{"firstName", "lastName", "emailAddress", "phoneNumber"}

    // Convert to database columns / 데이터베이스 컬럼으로 변환
    dbColumns := ConvertJSONToDBColumns(jsonFields)
    fmt.Println("Database columns:", dbColumns)
    // Output: Database columns: [first_name last_name email_address phone_number]

    // Convert back to JSON fields / JSON 필드로 다시 변환
    jsonFieldsBack := ConvertDBToJSONFields(dbColumns)
    fmt.Println("JSON fields:", jsonFieldsBack)
    // Output: JSON fields: [firstName lastName emailAddress phoneNumber]

    // Build SQL query dynamically / SQL 쿼리 동적 생성
    columns := stringutil.Join(dbColumns, ", ")
    query := fmt.Sprintf("SELECT %s FROM users", columns)
    fmt.Println("SQL Query:", query)
    // Output: SQL Query: SELECT first_name, last_name, email_address, phone_number FROM users
}
```

---

### Use Case 2: Email Validation in Registration Form

**Scenario / 시나리오**: Validate user email addresses during registration.

등록 중에 사용자 이메일 주소를 검증합니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

type RegistrationForm struct {
    Email    string
    Username string
    Password string
}

func (f *RegistrationForm) Validate() []string {
    var errors []string

    // Validate email / 이메일 검증
    if stringutil.IsBlank(f.Email) {
        errors = append(errors, "Email is required")
    } else if !stringutil.IsEmail(f.Email) {
        errors = append(errors, "Invalid email format")
    }

    // Validate username / 사용자 이름 검증
    if stringutil.IsBlank(f.Username) {
        errors = append(errors, "Username is required")
    } else {
        cleaned := stringutil.RemoveSpaces(f.Username)
        if !stringutil.IsAlphanumeric(cleaned) {
            errors = append(errors, "Username must be alphanumeric")
        }
        if len([]rune(cleaned)) < 3 {
            errors = append(errors, "Username must be at least 3 characters")
        }
    }

    // Validate password / 비밀번호 검증
    if stringutil.IsBlank(f.Password) {
        errors = append(errors, "Password is required")
    } else if len([]rune(f.Password)) < 8 {
        errors = append(errors, "Password must be at least 8 characters")
    }

    return errors
}

func main() {
    // Valid form / 유효한 폼
    form1 := RegistrationForm{
        Email:    "user@example.com",
        Username: "john123",
        Password: "password123",
    }
    errors1 := form1.Validate()
    fmt.Println("Form 1 errors:", errors1)
    // Output: Form 1 errors: []

    // Invalid form / 잘못된 폼
    form2 := RegistrationForm{
        Email:    "invalid.email",
        Username: "jo",
        Password: "short",
    }
    errors2 := form2.Validate()
    fmt.Println("Form 2 errors:", errors2)
    // Output: Form 2 errors: [Invalid email format Username must be at least 3 characters Password must be at least 8 characters]
}
```

---

### Use Case 3: Displaying Long Article Titles

**Scenario / 시나리오**: Truncate long article titles for display in a list view with Unicode support.

유니코드 지원으로 목록 보기에 표시할 긴 기사 제목을 자릅니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

type Article struct {
    ID          int
    Title       string
    Description string
    Author      string
}

func (a *Article) DisplayListView(maxTitleLen, maxDescLen int) {
    // Truncate title / 제목 자르기
    title := stringutil.Truncate(a.Title, maxTitleLen)

    // Truncate description / 설명 자르기
    description := stringutil.Truncate(a.Description, maxDescLen)

    // Display / 표시
    fmt.Printf("[%d] %s\n", a.ID, title)
    fmt.Printf("    %s\n", description)
    fmt.Printf("    Author: %s\n", a.Author)
    fmt.Println()
}

func main() {
    articles := []Article{
        {
            ID:          1,
            Title:       "Introduction to Go Programming Language and Its Ecosystem for Modern Application Development",
            Description: "This comprehensive article covers the fundamentals of Go programming, including syntax, concurrency patterns, best practices, and real-world applications in building scalable microservices.",
            Author:      "John Doe",
        },
        {
            ID:          2,
            Title:       "안녕하세요! Go 언어로 시작하는 웹 개발 완벽 가이드입니다. 초보자를 위한 상세한 설명이 포함되어 있습니다.",
            Description: "이 기사는 Go 언어를 사용한 웹 개발의 모든 것을 다룹니다. 기본 문법부터 고급 패턴까지, 실무에서 바로 사용할 수 있는 예제와 함께 설명합니다.",
            Author:      "홍길동",
        },
    }

    // Display all articles / 모든 기사 표시
    for _, article := range articles {
        article.DisplayListView(50, 80)
    }

    // Output:
    // [1] Introduction to Go Programming Language and Its...
    //     This comprehensive article covers the fundamentals of Go programming, includin...
    //     Author: John Doe
    //
    // [2] 안녕하세요! Go 언어로 시작하는 웹 개발 완벽 가이드입니다. 초보자를 위한 상세한 ...
    //     이 기사는 Go 언어를 사용한 웹 개발의 모든 것을 다룹니다. 기본 문법부터 고급 패턴까지, 실무에서 바로 사용할 수 있는 예제와 함께 ...
    //     Author: 홍길동
}
```

---

### Use Case 4: Configuration File Template Replacement

**Scenario / 시나리오**: Replace placeholders in configuration files with actual values.

설정 파일의 플레이스홀더를 실제 값으로 치환합니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

const ConfigTemplate = `
# Application Configuration
# Application Name: {{APP_NAME}}
# Environment: {{ENV}}

database:
  host: {{DB_HOST}}
  port: {{DB_PORT}}
  user: {{DB_USER}}
  password: {{DB_PASSWORD}}
  database: {{DB_NAME}}

server:
  host: {{SERVER_HOST}}
  port: {{SERVER_PORT}}
`

func GenerateConfig(values map[string]string) string {
    config := ConfigTemplate

    // Add placeholder markers / 플레이스홀더 마커 추가
    replacements := make(map[string]string)
    for key, value := range values {
        placeholder := "{{" + key + "}}"
        replacements[placeholder] = value
    }

    return stringutil.ReplaceAll(config, replacements)
}

func main() {
    // Configuration values / 설정 값
    values := map[string]string{
        "APP_NAME":    "MyApp",
        "ENV":         "production",
        "DB_HOST":     "localhost",
        "DB_PORT":     "3306",
        "DB_USER":     "root",
        "DB_PASSWORD": "secret",
        "DB_NAME":     "myapp_db",
        "SERVER_HOST": "0.0.0.0",
        "SERVER_PORT": "8080",
    }

    // Generate configuration / 설정 생성
    config := GenerateConfig(values)
    fmt.Println(config)

    // Output:
    // # Application Configuration
    // # Application Name: MyApp
    // # Environment: production
    //
    // database:
    //   host: localhost
    //   port: 3306
    //   user: root
    //   password: secret
    //   database: myapp_db
    //
    // server:
    //   host: 0.0.0.0
    //   port: 8080
}
```

---

### Use Case 5: Sanitizing HTML User Input

**Scenario / 시나리오**: Sanitize user-generated content to prevent XSS attacks.

XSS 공격을 방지하기 위해 사용자 생성 콘텐츠를 정제합니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

func SanitizeHTML(text string) string {
    // Replace HTML special characters / HTML 특수 문자 치환
    replacements := map[string]string{
        "&":  "&amp;",
        "<":  "&lt;",
        ">":  "&gt;",
        "\"": "&quot;",
        "'":  "&#39;",
    }

    return stringutil.ReplaceAll(text, replacements)
}

func SanitizeUserComment(comment string) string {
    // Clean whitespace / 공백 정리
    cleaned := stringutil.Clean(comment)

    // Sanitize HTML / HTML 정제
    sanitized := SanitizeHTML(cleaned)

    // Truncate if too long / 너무 길면 자르기
    if len([]rune(sanitized)) > 500 {
        sanitized = stringutil.Truncate(sanitized, 500)
    }

    return sanitized
}

func main() {
    // Malicious input / 악의적인 입력
    maliciousInput := "<script>alert('XSS attack!')</script>"
    sanitized := SanitizeUserComment(maliciousInput)
    fmt.Println("Sanitized:", sanitized)
    // Output: Sanitized: &lt;script&gt;alert(&#39;XSS attack!&#39;)&lt;/script&gt;

    // Normal comment with extra whitespace / 추가 공백이 있는 일반 주석
    comment := "  This is a   great   article!   <b>Thanks</b>  "
    sanitized2 := SanitizeUserComment(comment)
    fmt.Println("Sanitized:", sanitized2)
    // Output: Sanitized: This is a great article! &lt;b&gt;Thanks&lt;/b&gt;
}
```

---

### Use Case 6: Generating URL Slugs

**Scenario / 시나리오**: Convert article titles to URL-friendly slugs.

기사 제목을 URL 친화적인 슬러그로 변환합니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "fmt"
    "strings"
    "github.com/arkd0ng/go-utils/stringutil"
)

func GenerateSlug(title string) string {
    // Convert to lowercase / 소문자로 변환
    slug := strings.ToLower(title)

    // Remove special characters / 특수 문자 제거
    slug = stringutil.RemoveSpecialChars(slug)

    // Clean whitespace / 공백 정리
    slug = stringutil.Clean(slug)

    // Replace spaces with hyphens / 공백을 하이픈으로 치환
    slug = strings.ReplaceAll(slug, " ", "-")

    // Remove duplicate hyphens / 중복 하이픈 제거
    for strings.Contains(slug, "--") {
        slug = strings.ReplaceAll(slug, "--", "-")
    }

    // Trim hyphens / 하이픈 제거
    slug = strings.Trim(slug, "-")

    return slug
}

func main() {
    titles := []string{
        "Introduction to Go Programming!",
        "How to Build REST APIs with Go?",
        "Top 10 Go Libraries for 2024",
        "안녕하세요! Go 언어 가이드",
    }

    for _, title := range titles {
        slug := GenerateSlug(title)
        fmt.Printf("Title: %s\n", title)
        fmt.Printf("Slug:  %s\n\n", slug)
    }

    // Output:
    // Title: Introduction to Go Programming!
    // Slug:  introduction-to-go-programming
    //
    // Title: How to Build REST APIs with Go?
    // Slug:  how-to-build-rest-apis-with-go
    //
    // Title: Top 10 Go Libraries for 2024
    // Slug:  top-10-go-libraries-for-2024
    //
    // Title: 안녕하세요! Go 언어 가이드
    // Slug:  go-
}
```

---

### Use Case 7: Processing CSV Column Names

**Scenario / 시나리오**: Convert CSV column names between different naming conventions.

다양한 명명 규칙 간에 CSV 컬럼 이름을 변환합니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/stringutil"
)

type CSVProcessor struct {
    Headers []string
}

func (p *CSVProcessor) ConvertHeadersToSnakeCase() []string {
    return stringutil.Map(p.Headers, stringutil.ToSnakeCase)
}

func (p *CSVProcessor) ConvertHeadersToCamelCase() []string {
    return stringutil.Map(p.Headers, stringutil.ToCamelCase)
}

func (p *CSVProcessor) ConvertHeadersToPascalCase() []string {
    return stringutil.Map(p.Headers, stringutil.ToPascalCase)
}

func main() {
    // CSV headers from Excel (spaces) / Excel에서 온 CSV 헤더 (공백)
    processor := CSVProcessor{
        Headers: []string{
            "First Name",
            "Last Name",
            "Email Address",
            "Phone Number",
            "Date Of Birth",
        },
    }

    fmt.Println("Original Headers:", processor.Headers)
    fmt.Println("snake_case:", processor.ConvertHeadersToSnakeCase())
    fmt.Println("camelCase:", processor.ConvertHeadersToCamelCase())
    fmt.Println("PascalCase:", processor.ConvertHeadersToPascalCase())

    // Output:
    // Original Headers: [First Name Last Name Email Address Phone Number Date Of Birth]
    // snake_case: [first_name last_name email_address phone_number date_of_birth]
    // camelCase: [firstName lastName emailAddress phoneNumber dateOfBirth]
    // PascalCase: [FirstName LastName EmailAddress PhoneNumber DateOfBirth]
}
```

---

### Use Case 8: Building Dynamic SQL Queries

**Scenario / 시나리오**: Build SQL queries dynamically with proper field name conversion.

적절한 필드 이름 변환으로 SQL 쿼리를 동적으로 생성합니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "fmt"
    "strings"
    "github.com/arkd0ng/go-utils/stringutil"
)

type QueryBuilder struct {
    Table   string
    Fields  []string
    Where   map[string]interface{}
}

func (qb *QueryBuilder) BuildSelectQuery() string {
    // Convert field names to snake_case / 필드 이름을 snake_case로 변환
    columns := stringutil.Map(qb.Fields, stringutil.ToSnakeCase)
    columnsStr := stringutil.Join(columns, ", ")

    // Build query / 쿼리 생성
    query := fmt.Sprintf("SELECT %s FROM %s", columnsStr, qb.Table)

    // Add WHERE clause / WHERE 절 추가
    if len(qb.Where) > 0 {
        var conditions []string
        for key, value := range qb.Where {
            column := stringutil.ToSnakeCase(key)
            condition := fmt.Sprintf("%s = '%v'", column, value)
            conditions = append(conditions, condition)
        }
        query += " WHERE " + strings.Join(conditions, " AND ")
    }

    return query
}

func main() {
    // Build SELECT query / SELECT 쿼리 생성
    qb := QueryBuilder{
        Table:  "users",
        Fields: []string{"UserId", "FirstName", "LastName", "EmailAddress"},
        Where: map[string]interface{}{
            "IsActive": true,
            "UserRole": "admin",
        },
    }

    query := qb.BuildSelectQuery()
    fmt.Println(query)
    // Output: SELECT user_id, first_name, last_name, email_address FROM users WHERE is_active = 'true' AND user_role = 'admin'
}
```

---

### Use Case 9: Validating File Uploads

**Scenario / 시나리오**: Validate uploaded file names and extensions.

업로드된 파일 이름 및 확장자를 검증합니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "fmt"
    "strings"
    "github.com/arkd0ng/go-utils/stringutil"
)

type FileUpload struct {
    Filename string
    Size     int64
}

func (f *FileUpload) Validate(allowedExtensions []string, maxSizeMB int64) []string {
    var errors []string

    // Validate filename / 파일 이름 검증
    if stringutil.IsBlank(f.Filename) {
        errors = append(errors, "Filename is required")
        return errors
    }

    // Validate extension / 확장자 검증
    if !stringutil.EndsWithAny(strings.ToLower(f.Filename), allowedExtensions) {
        errors = append(errors, fmt.Sprintf("File extension not allowed. Allowed: %v", allowedExtensions))
    }

    // Validate size / 크기 검증
    maxSizeBytes := maxSizeMB * 1024 * 1024
    if f.Size > maxSizeBytes {
        errors = append(errors, fmt.Sprintf("File size exceeds maximum %dMB", maxSizeMB))
    }

    // Validate filename characters / 파일 이름 문자 검증
    cleaned := stringutil.RemoveSpecialChars(f.Filename)
    if cleaned == "" {
        errors = append(errors, "Filename contains invalid characters")
    }

    return errors
}

func main() {
    // Valid image upload / 유효한 이미지 업로드
    upload1 := FileUpload{
        Filename: "photo.jpg",
        Size:     1024 * 1024 * 2, // 2MB
    }
    errors1 := upload1.Validate([]string{".jpg", ".jpeg", ".png", ".gif"}, 5)
    fmt.Println("Upload 1 errors:", errors1)
    // Output: Upload 1 errors: []

    // Invalid upload (wrong extension) / 잘못된 업로드 (잘못된 확장자)
    upload2 := FileUpload{
        Filename: "document.exe",
        Size:     1024 * 1024, // 1MB
    }
    errors2 := upload2.Validate([]string{".jpg", ".jpeg", ".png", ".gif"}, 5)
    fmt.Println("Upload 2 errors:", errors2)
    // Output: Upload 2 errors: [File extension not allowed. Allowed: [.jpg .jpeg .png .gif]]

    // Invalid upload (too large) / 잘못된 업로드 (너무 큼)
    upload3 := FileUpload{
        Filename: "photo.jpg",
        Size:     1024 * 1024 * 10, // 10MB
    }
    errors3 := upload3.Validate([]string{".jpg", ".jpeg", ".png", ".gif"}, 5)
    fmt.Println("Upload 3 errors:", errors3)
    // Output: Upload 3 errors: [File size exceeds maximum 5MB]
}
```

---

### Use Case 10: Generating Environment Variable Names

**Scenario / 시나리오**: Convert configuration keys to environment variable names.

설정 키를 환경 변수 이름으로 변환합니다.

**Complete Example / 완전한 예제**:

```go
package main

import (
    "fmt"
    "os"
    "github.com/arkd0ng/go-utils/stringutil"
)

type Config struct {
    DatabaseUrl    string
    ApiKey         string
    MaxConnections int
    CacheEnabled   bool
}

func LoadConfigFromEnv() *Config {
    config := &Config{}

    // Define config keys / 설정 키 정의
    keys := []string{
        "DatabaseUrl",
        "ApiKey",
        "MaxConnections",
        "CacheEnabled",
    }

    // Convert to environment variable names / 환경 변수 이름으로 변환
    envVars := stringutil.Map(keys, stringutil.ToScreamingSnakeCase)

    fmt.Println("Loading config from environment variables:")
    for i, key := range keys {
        envVar := envVars[i]
        value := os.Getenv(envVar)
        fmt.Printf("  %s = %s\n", envVar, value)
    }

    return config
}

func main() {
    // Set environment variables / 환경 변수 설정
    os.Setenv("DATABASE_URL", "mysql://localhost:3306/mydb")
    os.Setenv("API_KEY", "secret-key-123")
    os.Setenv("MAX_CONNECTIONS", "100")
    os.Setenv("CACHE_ENABLED", "true")

    // Load configuration / 설정 로드
    config := LoadConfigFromEnv()
    _ = config

    // Output:
    // Loading config from environment variables:
    //   DATABASE_URL = mysql://localhost:3306/mydb
    //   API_KEY = secret-key-123
    //   MAX_CONNECTIONS = 100
    //   CACHE_ENABLED = true
}
```

---

## Best Practices / 모범 사례

### 1. Always Use Unicode-Safe Functions

When working with multi-byte characters (emoji, CJK characters), always use the `stringutil` functions instead of byte-based operations.

다중 바이트 문자 (이모지, CJK 문자)를 사용할 때는 바이트 기반 작업 대신 항상 `stringutil` 함수를 사용하세요.

**Bad / 나쁨**:
```go
// This breaks with Unicode! / 유니코드에서 깨집니다!
func truncateBad(s string, length int) string {
    if len(s) <= length {
        return s
    }
    return s[:length] + "..."
}

text := "안녕하세요"
fmt.Println(truncateBad(text, 3))  // ��� (broken!)
```

**Good / 좋음**:
```go
text := "안녕하세요"
fmt.Println(stringutil.Truncate(text, 3))  // 안녕하... (correct!)
```

---

### 2. Validate User Input Early

Always validate user input as early as possible in your application flow.

항상 애플리케이션 흐름에서 가능한 한 빨리 사용자 입력을 검증하세요.

**Good / 좋음**:
```go
func HandleUserRegistration(email, username string) error {
    // Validate immediately / 즉시 검증
    if !stringutil.IsEmail(email) {
        return fmt.Errorf("invalid email")
    }

    if !stringutil.IsAlphanumeric(username) {
        return fmt.Errorf("username must be alphanumeric")
    }

    // Continue with business logic / 비즈니스 로직 계속
    // ...
    return nil
}
```

---

### 3. Use Map and Filter for Batch Operations

When processing multiple strings, use `Map` and `Filter` instead of loops for cleaner code.

여러 문자열을 처리할 때는 더 깨끗한 코드를 위해 루프 대신 `Map`과 `Filter`를 사용하세요.

**Bad / 나쁨**:
```go
var snakeNames []string
for _, name := range camelNames {
    snakeNames = append(snakeNames, stringutil.ToSnakeCase(name))
}
```

**Good / 좋음**:
```go
snakeNames := stringutil.Map(camelNames, stringutil.ToSnakeCase)
```

---

### 4. Clean User Input Before Processing

Always clean user input (trim, remove extra spaces) before validation or processing.

검증이나 처리 전에 항상 사용자 입력을 정리하세요 (공백 제거, 추가 공백 제거).

**Good / 좋음**:
```go
func ProcessComment(comment string) string {
    // Clean first / 먼저 정리
    cleaned := stringutil.Clean(comment)

    // Then capitalize / 그 다음 대문자화
    cleaned = stringutil.CapitalizeFirst(cleaned)

    // Then truncate if needed / 필요하면 자르기
    if len([]rune(cleaned)) > 200 {
        cleaned = stringutil.Truncate(cleaned, 200)
    }

    return cleaned
}
```

---

### 5. Use Consistent Case Conventions

Use consistent case conventions throughout your application:
- JSON fields: camelCase
- Database columns: snake_case
- Environment variables: SCREAMING_SNAKE_CASE
- Constants: SCREAMING_SNAKE_CASE
- URL slugs: kebab-case

애플리케이션 전체에서 일관된 케이스 규칙을 사용하세요:
- JSON 필드: camelCase
- 데이터베이스 컬럼: snake_case
- 환경 변수: SCREAMING_SNAKE_CASE
- 상수: SCREAMING_SNAKE_CASE
- URL 슬러그: kebab-case

**Good / 좋음**:
```go
// JSON / JSON
type APIRequest struct {
    FirstName string `json:"firstName"`  // camelCase
}

// Database / 데이터베이스
type DBModel struct {
    FirstName string `db:"first_name"`  // snake_case
}

// Environment / 환경 변수
const DATABASE_URL = os.Getenv("DATABASE_URL")  // SCREAMING_SNAKE_CASE
```

---

### 6. Prefer ContainsAny over Multiple Contains

When checking for multiple substrings, use `ContainsAny` instead of multiple `Contains` calls.

여러 부분 문자열을 확인할 때는 여러 `Contains` 호출 대신 `ContainsAny`를 사용하세요.

**Bad / 나쁨**:
```go
if strings.Contains(url, "http://") || strings.Contains(url, "https://") {
    // ...
}
```

**Good / 좋음**:
```go
if stringutil.ContainsAny(url, []string{"http://", "https://"}) {
    // ...
}
```

---

### 7. Use ReplaceAll for Multi-Pattern Replacement

When replacing multiple patterns, use `ReplaceAll` with a map instead of multiple `Replace` calls.

여러 패턴을 치환할 때는 여러 `Replace` 호출 대신 맵과 함께 `ReplaceAll`을 사용하세요.

**Bad / 나쁨**:
```go
s = strings.ReplaceAll(s, "&", "&amp;")
s = strings.ReplaceAll(s, "<", "&lt;")
s = strings.ReplaceAll(s, ">", "&gt;")
```

**Good / 좋음**:
```go
s = stringutil.ReplaceAll(s, map[string]string{
    "&": "&amp;",
    "<": "&lt;",
    ">": "&gt;",
})
```

---

### 8. Use TruncateWithSuffix for Custom Ellipsis

When you need a custom truncation suffix, use `TruncateWithSuffix` instead of `Truncate`.

사용자 정의 자르기 접미사가 필요할 때는 `Truncate` 대신 `TruncateWithSuffix`를 사용하세요.

**Good / 좋음**:
```go
// For Unicode ellipsis / 유니코드 줄임표를 위해
text := stringutil.TruncateWithSuffix(longText, 50, "…")

// For custom suffix / 사용자 정의 접미사를 위해
text := stringutil.TruncateWithSuffix(longText, 50, " [Read more]")
```

---

### 9. Combine Functions for Complex Operations

Combine multiple `stringutil` functions for complex string processing.

복잡한 문자열 처리를 위해 여러 `stringutil` 함수를 결합하세요.

**Good / 좋음**:
```go
func NormalizeUsername(username string) string {
    // Clean whitespace / 공백 정리
    cleaned := stringutil.Clean(username)

    // Remove special characters / 특수 문자 제거
    cleaned = stringutil.RemoveSpecialChars(cleaned)

    // Convert to lowercase / 소문자로 변환
    cleaned = strings.ToLower(cleaned)

    // Remove spaces / 공백 제거
    cleaned = stringutil.RemoveSpaces(cleaned)

    return cleaned
}
```

---

### 10. Use Filter with Named Functions

For complex filter logic, use named functions instead of anonymous functions for better readability.

복잡한 필터 로직의 경우 더 나은 가독성을 위해 익명 함수 대신 이름 있는 함수를 사용하세요.

**Good / 좋음**:
```go
func isLongWord(s string) bool {
    return len(s) > 5
}

func containsDigit(s string) bool {
    for _, r := range s {
        if unicode.IsDigit(r) {
            return true
        }
    }
    return false
}

// Use named functions / 이름 있는 함수 사용
longWords := stringutil.Filter(words, isLongWord)
wordsWithDigits := stringutil.Filter(words, containsDigit)
```

---

### 11. Handle Empty Strings Gracefully

Always check for empty strings before processing to avoid unexpected behavior.

예상치 못한 동작을 피하기 위해 처리 전에 항상 빈 문자열을 확인하세요.

**Good / 좋음**:
```go
func ProcessText(text string) string {
    if stringutil.IsBlank(text) {
        return ""  // or default value
    }

    // Continue processing / 처리 계속
    return stringutil.Clean(text)
}
```

---

### 12. Use Join Instead of Manual Concatenation

For joining multiple strings, use `Join` instead of manual concatenation.

여러 문자열을 결합할 때는 수동 연결 대신 `Join`을 사용하세요.

**Bad / 나쁨**:
```go
result := words[0] + ", " + words[1] + ", " + words[2]
```

**Good / 좋음**:
```go
result := stringutil.Join(words, ", ")
```

---

### 13. Validate Before Converting

When converting case formats, validate the input first to ensure it's in an expected format.

케이스 형식을 변환할 때는 먼저 입력이 예상 형식인지 검증하세요.

**Good / 좋음**:
```go
func ConvertFieldName(field string) (string, error) {
    // Validate first / 먼저 검증
    if stringutil.IsBlank(field) {
        return "", fmt.Errorf("field name cannot be blank")
    }

    // Then convert / 그 다음 변환
    return stringutil.ToSnakeCase(field), nil
}
```

---

### 14. Use IsBlank Instead of Checking Empty String

Use `IsBlank` to check for both empty strings and whitespace-only strings.

빈 문자열과 공백만 있는 문자열 모두를 확인하려면 `IsBlank`를 사용하세요.

**Bad / 나쁨**:
```go
if s == "" || strings.TrimSpace(s) == "" {
    // ...
}
```

**Good / 좋음**:
```go
if stringutil.IsBlank(s) {
    // ...
}
```

---

### 15. Test with Unicode Characters

Always test your string processing with Unicode characters to ensure proper handling.

적절한 처리를 보장하기 위해 항상 유니코드 문자로 문자열 처리를 테스트하세요.

**Good / 좋음**:
```go
func TestTruncateUnicode(t *testing.T) {
    tests := []struct {
        input    string
        length   int
        expected string
    }{
        {"안녕하세요", 3, "안녕하..."},
        {"Hello 世界", 8, "Hello 世界"},
        {"😀😁😂", 2, "😀😁..."},
    }
    // ...
}
```

---

## Troubleshooting / 문제 해결

### Problem 1: Truncate Breaks Unicode Characters

**Symptom / 증상**: When truncating strings with Unicode characters, output shows broken characters (�).

유니코드 문자가 있는 문자열을 자를 때 출력에 깨진 문자 (�)가 표시됩니다.

**Cause / 원인**: Using byte-based truncation instead of rune-based.

rune 기반 대신 바이트 기반 자르기를 사용합니다.

**Solution / 해결책**: Always use `stringutil.Truncate()` which handles runes correctly.

rune을 올바르게 처리하는 `stringutil.Truncate()`를 항상 사용하세요.

```go
// Bad / 나쁨
text := "안녕하세요"
fmt.Println(text[:3])  // ��� (broken!)

// Good / 좋음
fmt.Println(stringutil.Truncate(text, 3))  // 안녕하... (correct!)
```

---

### Problem 2: Case Conversion Not Working as Expected

**Symptom / 증상**: `ToSnakeCase` or `ToCamelCase` produces unexpected output.

`ToSnakeCase` 또는 `ToCamelCase`가 예상치 못한 출력을 생성합니다.

**Cause / 원인**: Input string has unexpected format or delimiters.

입력 문자열에 예상치 못한 형식이나 구분자가 있습니다.

**Solution / 해결책**: Clean the input string first before conversion.

변환 전에 먼저 입력 문자열을 정리하세요.

```go
// Input with extra spaces / 추가 공백이 있는 입력
input := "  User  Profile  Data  "

// Clean first / 먼저 정리
cleaned := stringutil.Clean(input)
result := stringutil.ToSnakeCase(cleaned)
fmt.Println(result)  // user_profile_data
```

---

### Problem 3: Email Validation Too Strict or Too Loose

**Symptom / 증상**: `IsEmail` rejects valid emails or accepts invalid ones.

`IsEmail`이 유효한 이메일을 거부하거나 잘못된 이메일을 허용합니다.

**Cause / 원인**: `IsEmail` uses practical validation (99% of cases), not RFC 5322 compliant.

`IsEmail`은 실용적 검증 (99%의 경우)을 사용하며 RFC 5322를 준수하지 않습니다.

**Solution / 해결책**: For strict validation, use a specialized email validation library. For most use cases, `IsEmail` is sufficient.

엄격한 검증을 위해서는 전문 이메일 검증 라이브러리를 사용하세요. 대부분의 사용 사례에서는 `IsEmail`이 충분합니다.

```go
// Practical validation (99% of cases) / 실용적 검증 (99%의 경우)
if stringutil.IsEmail(email) {
    // Process email / 이메일 처리
}

// For strict validation, use external library
// 엄격한 검증을 위해서는 외부 라이브러리 사용
// import "github.com/go-mail/mail"
```

---

### Problem 4: ReplaceAll Order Matters

**Symptom / 증상**: `ReplaceAll` produces unexpected results when replacements overlap.

치환이 겹칠 때 `ReplaceAll`이 예상치 못한 결과를 생성합니다.

**Cause / 원인**: Replacements are applied in map iteration order (undefined).

치환이 맵 반복 순서 (정의되지 않음)로 적용됩니다.

**Solution / 해결책**: Be careful with overlapping replacements. Consider applying them in a specific order.

겹치는 치환에 주의하세요. 특정 순서로 적용하는 것을 고려하세요.

```go
// Overlapping replacements / 겹치는 치환
replacements := map[string]string{
    "hello": "hi",
    "hi":    "hey",
}

// Result is undefined! / 결과가 정의되지 않습니다!
result := stringutil.ReplaceAll("hello", replacements)

// Solution: Apply in specific order / 해결책: 특정 순서로 적용
s := "hello"
s = strings.ReplaceAll(s, "hello", "hi")
s = strings.ReplaceAll(s, "hi", "hey")
fmt.Println(s)  // hey
```

---

### Problem 5: Map and Filter Modify Original Slice

**Symptom / 증상**: Modifying the result of `Map` or `Filter` seems to affect the original slice.

`Map` 또는 `Filter`의 결과를 수정하면 원본 슬라이스에 영향을 미치는 것처럼 보입니다.

**Cause / 원인**: Misunderstanding - `Map` and `Filter` always return new slices.

오해 - `Map`과 `Filter`는 항상 새 슬라이스를 반환합니다.

**Solution / 해결책**: `Map` and `Filter` always create new slices. The original slice is never modified.

`Map`과 `Filter`는 항상 새 슬라이스를 생성합니다. 원본 슬라이스는 절대 수정되지 않습니다.

```go
original := []string{"hello", "world"}
upper := stringutil.Map(original, strings.ToUpper)

fmt.Println(original)  // [hello world] (unchanged)
fmt.Println(upper)     // [HELLO WORLD]
```

---

### Problem 6: PadLeft/PadRight Not Padding Enough

**Symptom / 증상**: `PadLeft` or `PadRight` doesn't pad to the expected length.

`PadLeft` 또는 `PadRight`가 예상 길이로 패딩하지 않습니다.

**Cause / 원인**: Length parameter is in runes, not bytes. Multi-character pad strings count as multiple runes.

길이 매개변수가 바이트가 아닌 rune입니다. 다중 문자 패드 문자열은 여러 rune으로 계산됩니다.

**Solution / 해결책**: Use single-character pad strings for predictable results.

예측 가능한 결과를 위해 단일 문자 패드 문자열을 사용하세요.

```go
// Single character pad / 단일 문자 패드
fmt.Println(stringutil.PadLeft("5", 3, "0"))  // 005

// Multi-character pad (may not work as expected)
// 다중 문자 패드 (예상대로 작동하지 않을 수 있음)
fmt.Println(stringutil.PadLeft("5", 5, "00"))  // 000005 (10 characters!)
```

---

### Problem 7: IsURL Rejects Valid URLs

**Symptom / 증상**: `IsURL` rejects URLs that are valid.

`IsURL`이 유효한 URL을 거부합니다.

**Cause / 원인**: `IsURL` only checks for http:// and https:// prefixes. It's a simple validation.

`IsURL`은 http:// 및 https:// 접두사만 확인합니다. 간단한 검증입니다.

**Solution / 해결책**: For more comprehensive URL validation, use `net/url` package or a specialized library.

더 포괄적인 URL 검증을 위해서는 `net/url` 패키지 또는 전문 라이브러리를 사용하세요.

```go
// Simple validation / 간단한 검증
if stringutil.IsURL(url) {
    // Only checks http:// or https:// prefix
    // http:// 또는 https:// 접두사만 확인
}

// Comprehensive validation / 포괄적 검증
import "net/url"

func IsValidURL(s string) bool {
    _, err := url.Parse(s)
    return err == nil
}
```

---

### Problem 8: Filter Returns Empty Slice

**Symptom / 증상**: `Filter` always returns an empty slice.

`Filter`가 항상 빈 슬라이스를 반환합니다.

**Cause / 원인**: Filter function always returns false.

필터 함수가 항상 false를 반환합니다.

**Solution / 해결책**: Debug the filter function to ensure it returns true for expected cases.

필터 함수가 예상 경우에 true를 반환하는지 디버그하세요.

```go
words := []string{"hello", "world", "foo"}

// Bug: Always returns false / 버그: 항상 false 반환
filtered := stringutil.Filter(words, func(s string) bool {
    return false  // Wrong!
})
fmt.Println(filtered)  // []

// Fix: Return true for matching cases / 수정: 일치하는 경우에 true 반환
filtered = stringutil.Filter(words, func(s string) bool {
    return len(s) > 3
})
fmt.Println(filtered)  // [hello world]
```

---

### Problem 9: RemoveSpecialChars Removes Too Much

**Symptom / 증상**: `RemoveSpecialChars` removes characters you want to keep.

`RemoveSpecialChars`가 유지하려는 문자를 제거합니다.

**Cause / 원인**: `RemoveSpecialChars` only keeps letters, digits, and spaces. Everything else is removed.

`RemoveSpecialChars`는 문자, 숫자, 공백만 유지합니다. 그 외는 모두 제거됩니다.

**Solution / 해결책**: Use `ReplaceAll` with a custom replacement map for more control.

더 많은 제어를 위해 사용자 정의 치환 맵과 함께 `ReplaceAll`을 사용하세요.

```go
// RemoveSpecialChars removes everything except alphanumeric and spaces
// RemoveSpecialChars는 영숫자와 공백을 제외한 모든 것을 제거합니다
text := "hello@world.com"
cleaned := stringutil.RemoveSpecialChars(text)
fmt.Println(cleaned)  // helloworld com

// For more control, use ReplaceAll / 더 많은 제어를 위해 ReplaceAll 사용
text = "hello@world.com"
cleaned = stringutil.ReplaceAll(text, map[string]string{
    "@": " at ",
    ".": " dot ",
})
fmt.Println(cleaned)  // hello at world dot com
```

---

### Problem 10: ToSnakeCase Produces Unexpected Results

**Symptom / 증상**: `ToSnakeCase` produces unexpected underscores or missing words.

`ToSnakeCase`가 예상치 못한 밑줄이나 누락된 단어를 생성합니다.

**Cause / 원인**: Input has consecutive uppercase letters (e.g., HTTPServer).

입력에 연속된 대문자가 있습니다 (예: HTTPServer).

**Solution / 해결책**: This is expected behavior. HTTPServer → http_server. If you need different behavior, pre-process the input.

이것은 예상된 동작입니다. HTTPServer → http_server. 다른 동작이 필요하면 입력을 사전 처리하세요.

```go
// Expected behavior / 예상 동작
fmt.Println(stringutil.ToSnakeCase("HTTPServer"))  // http_server

// If you want "h_t_t_p_server", pre-process / "h_t_t_p_server"를 원하면 사전 처리
// (not recommended / 권장하지 않음)
```

---

## FAQ

### Q1: Does stringutil support Unicode?

**A**: Yes! All functions in the stringutil package are Unicode-safe. They use `[]rune` instead of byte operations to properly handle multi-byte characters like emoji, CJK characters, etc.

**답변**: 네! stringutil 패키지의 모든 함수는 유니코드 안전합니다. 이모지, CJK 문자 등의 다중 바이트 문자를 올바르게 처리하기 위해 바이트 작업 대신 `[]rune`을 사용합니다.

---

### Q2: What is the difference between IsEmail and RFC 5322 compliant validation?

**A**: `IsEmail` uses practical validation that covers 99% of real-world email addresses. It's much simpler and faster than RFC 5322 compliant validation. For most applications, `IsEmail` is sufficient. If you need strict RFC 5322 compliance, use a specialized library.

**답변**: `IsEmail`은 실제 이메일 주소의 99%를 커버하는 실용적 검증을 사용합니다. RFC 5322 준수 검증보다 훨씬 간단하고 빠릅니다. 대부분의 애플리케이션에서는 `IsEmail`이 충분합니다. 엄격한 RFC 5322 준수가 필요한 경우 전문 라이브러리를 사용하세요.

---

### Q3: Can I use stringutil in production?

**A**: Yes! The package is designed for production use with:
- Zero external dependencies
- Unicode-safe operations
- Well-tested functions
- Comprehensive documentation

**답변**: 네! 패키지는 다음과 같은 프로덕션 사용을 위해 설계되었습니다:
- 외부 의존성 제로
- 유니코드 안전 작업
- 잘 테스트된 함수
- 포괄적인 문서

---

### Q4: How do I convert between different case formats?

**A**: Use the case conversion functions:
- `ToSnakeCase`: snake_case
- `ToCamelCase`: camelCase
- `ToKebabCase`: kebab-case
- `ToPascalCase`: PascalCase
- `ToScreamingSnakeCase`: SCREAMING_SNAKE_CASE

All functions intelligently handle multiple input formats.

**답변**: 케이스 변환 함수를 사용하세요:
- `ToSnakeCase`: snake_case
- `ToCamelCase`: camelCase
- `ToKebabCase`: kebab-case
- `ToPascalCase`: PascalCase
- `ToScreamingSnakeCase`: SCREAMING_SNAKE_CASE

모든 함수는 여러 입력 형식을 지능적으로 처리합니다.

---

### Q5: What's the performance of stringutil functions?

**A**: All functions are optimized for performance:
- Case conversion: O(n) where n is string length
- Validation: O(n) for most functions
- Map/Filter: O(n) where n is slice length
- No unnecessary allocations

**답변**: 모든 함수는 성능을 위해 최적화되어 있습니다:
- 케이스 변환: O(n), n은 문자열 길이
- 검증: 대부분의 함수에서 O(n)
- Map/Filter: O(n), n은 슬라이스 길이
- 불필요한 할당 없음

---

### Q6: How do I handle errors?

**A**: Most stringutil functions do not return errors because they are designed to be simple and predictable. For validation functions, check the boolean return value:

```go
if !stringutil.IsEmail(email) {
    return fmt.Errorf("invalid email")
}
```

**답변**: 대부분의 stringutil 함수는 간단하고 예측 가능하도록 설계되어 에러를 반환하지 않습니다. 검증 함수의 경우 부울 반환 값을 확인하세요:

```go
if !stringutil.IsEmail(email) {
    return fmt.Errorf("잘못된 이메일")
}
```

---

### Q7: Can I use Map with custom functions?

**A**: Yes! `Map` accepts any function with signature `func(string) string`:

```go
// Built-in function / 내장 함수
upper := stringutil.Map(words, strings.ToUpper)

// Custom function / 사용자 정의 함수
prefixed := stringutil.Map(words, func(s string) string {
    return "prefix_" + s
})
```

**답변**: 네! `Map`은 시그니처 `func(string) string`을 가진 모든 함수를 허용합니다.

---

### Q8: How do I validate file extensions?

**A**: Use `EndsWithAny` for efficient multi-extension validation:

```go
func IsImageFile(filename string) bool {
    return stringutil.EndsWithAny(filename, []string{".jpg", ".jpeg", ".png", ".gif"})
}
```

**답변**: 효율적인 다중 확장자 검증을 위해 `EndsWithAny`를 사용하세요.

---

### Q9: What's the difference between Clean and Truncate?

**A**:
- `Clean`: Trims whitespace and removes duplicate spaces
- `Truncate`: Cuts string to specified length and adds "..."

They serve different purposes and can be used together:

```go
cleaned := stringutil.Clean(userInput)
truncated := stringutil.Truncate(cleaned, 100)
```

**답변**:
- `Clean`: 공백을 제거하고 중복 공백을 제거합니다
- `Truncate`: 문자열을 지정된 길이로 자르고 "..."를 추가합니다

서로 다른 목적을 가지며 함께 사용할 수 있습니다.

---

### Q10: Can I contribute to stringutil?

**A**: Yes! Contributions are welcome. Follow the project's contribution guidelines:
1. Fork the repository
2. Create a feature branch
3. Write tests for new functions
4. Ensure all tests pass
5. Submit a pull request

**답변**: 네! 기여를 환영합니다. 프로젝트의 기여 가이드라인을 따르세요:
1. 리포지토리 포크
2. 기능 브랜치 생성
3. 새 함수에 대한 테스트 작성
4. 모든 테스트 통과 확인
5. 풀 리퀘스트 제출

---

### Q11: Does stringutil work with empty strings?

**A**: Yes! All functions handle empty strings gracefully:

```go
fmt.Println(stringutil.ToSnakeCase(""))      // ""
fmt.Println(stringutil.Truncate("", 10))     // ""
fmt.Println(stringutil.IsEmail(""))          // false
fmt.Println(stringutil.IsBlank(""))          // true
```

**답변**: 네! 모든 함수는 빈 문자열을 우아하게 처리합니다.

---

### Q12: How do I build SQL queries with stringutil?

**A**: Use case conversion functions to convert field names:

```go
fields := []string{"UserId", "FirstName", "LastName"}
columns := stringutil.Map(fields, stringutil.ToSnakeCase)
query := fmt.Sprintf("SELECT %s FROM users", stringutil.Join(columns, ", "))
// SELECT user_id, first_name, last_name FROM users
```

**답변**: 케이스 변환 함수를 사용하여 필드 이름을 변환하세요.

---

### Q13: What's the recommended way to sanitize user input?

**A**: Combine multiple functions:

```go
func SanitizeInput(input string) string {
    // Clean whitespace / 공백 정리
    cleaned := stringutil.Clean(input)

    // Truncate if too long / 너무 길면 자르기
    if len([]rune(cleaned)) > 200 {
        cleaned = stringutil.Truncate(cleaned, 200)
    }

    // Sanitize HTML / HTML 정제
    cleaned = stringutil.ReplaceAll(cleaned, map[string]string{
        "<":  "&lt;",
        ">":  "&gt;",
        "&":  "&amp;",
    })

    return cleaned
}
```

**답변**: 여러 함수를 결합하세요.

---

### Q14: Can I use stringutil for URL generation?

**A**: Yes! Use case conversion and character removal:

```go
func GenerateSlug(title string) string {
    slug := strings.ToLower(title)
    slug = stringutil.RemoveSpecialChars(slug)
    slug = strings.ReplaceAll(slug, " ", "-")
    return slug
}
```

**답변**: 네! 케이스 변환 및 문자 제거를 사용하세요.

---

### Q15: Where can I find more examples?

**A**: Check the following resources:
- Package README: `stringutil/README.md`
- Examples directory: `examples/stringutil/`
- Test files: `stringutil/*_test.go`
- Developer Guide: `docs/stringutil/DEVELOPER_GUIDE.md`

**답변**: 다음 리소스를 확인하세요:
- 패키지 README: `stringutil/README.md`
- 예제 디렉토리: `examples/stringutil/`
- 테스트 파일: `stringutil/*_test.go`
- 개발자 가이드: `docs/stringutil/DEVELOPER_GUIDE.md`

---

## Appendix: Function Reference Table / 부록: 함수 참조 표

| Function / 함수 | Category / 카테고리 | Unicode Safe / 유니코드 안전 | Description / 설명 |
|-----------------|---------------------|----------------------------|-------------------|
| ToSnakeCase | Case Conversion | ✅ | Converts to snake_case |
| ToCamelCase | Case Conversion | ✅ | Converts to camelCase |
| ToKebabCase | Case Conversion | ✅ | Converts to kebab-case |
| ToPascalCase | Case Conversion | ✅ | Converts to PascalCase |
| ToScreamingSnakeCase | Case Conversion | ✅ | Converts to SCREAMING_SNAKE_CASE |
| Truncate | Manipulation | ✅ | Truncates with "..." |
| TruncateWithSuffix | Manipulation | ✅ | Truncates with custom suffix |
| Reverse | Manipulation | ✅ | Reverses string |
| Capitalize | Manipulation | ✅ | Capitalizes each word |
| CapitalizeFirst | Manipulation | ✅ | Capitalizes first letter only |
| RemoveDuplicates | Manipulation | ✅ | Removes duplicate consecutive chars |
| RemoveSpaces | Manipulation | ✅ | Removes all whitespace |
| RemoveSpecialChars | Manipulation | ✅ | Removes non-alphanumeric chars |
| Clean | Manipulation | ✅ | Trims and deduplicates spaces |
| IsEmail | Validation | ✅ | Validates email format |
| IsURL | Validation | ✅ | Checks for http(s):// prefix |
| IsAlphanumeric | Validation | ✅ | Checks letters and digits only |
| IsNumeric | Validation | ✅ | Checks digits only |
| IsAlpha | Validation | ✅ | Checks letters only |
| IsBlank | Validation | ✅ | Checks empty or whitespace |
| IsLower | Validation | ✅ | Checks lowercase letters only |
| IsUpper | Validation | ✅ | Checks uppercase letters only |
| ContainsAny | Search & Replace | ✅ | Checks for any substring |
| ContainsAll | Search & Replace | ✅ | Checks for all substrings |
| StartsWithAny | Search & Replace | ✅ | Checks for any prefix |
| EndsWithAny | Search & Replace | ✅ | Checks for any suffix |
| ReplaceAll | Search & Replace | ✅ | Replaces multiple patterns |
| ReplaceIgnoreCase | Search & Replace | ✅ | Case-insensitive replace |
| CountWords | Utilities | ✅ | Counts words |
| CountOccurrences | Utilities | ✅ | Counts substring occurrences |
| Join | Utilities | ✅ | Joins strings with separator |
| Map | Utilities | ✅ | Applies function to each string |
| Filter | Utilities | ✅ | Filters strings by predicate |
| PadLeft | Utilities | ✅ | Pads string on left |
| PadRight | Utilities | ✅ | Pads string on right |
| Lines | Utilities | ✅ | Splits into lines |
| Words | Utilities | ✅ | Splits into words |

---

**End of User Manual / 사용자 매뉴얼 끝**

For developer documentation, see [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md)

개발자 문서는 [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md)를 참조하세요.
