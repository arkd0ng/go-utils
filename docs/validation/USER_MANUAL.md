# Validation Package - User Manual / Validation 패키지 - 사용자 매뉴얼

**Version / 버전**: v1.13.013
**Last Updated / 최종 업데이트**: 2025-10-17

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Core Concepts / 핵심 개념](#core-concepts--핵심-개념)
5. [String Validators / 문자열 검증기](#string-validators--문자열-검증기)
6. [Numeric Validators / 숫자 검증기](#numeric-validators--숫자-검증기)
7. [Collection Validators / 컬렉션 검증기](#collection-validators--컬렉션-검증기)
8. [Comparison Validators / 비교 검증기](#comparison-validators--비교-검증기)
9. [Advanced Features / 고급 기능](#advanced-features--고급-기능)
10. [Error Handling / 에러 처리](#error-handling--에러-처리)
11. [Real-World Examples / 실제 사용 예제](#real-world-examples--실제-사용-예제)
12. [Best Practices / 모범 사례](#best-practices--모범-사례)
13. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)

---

## Introduction / 소개

The `validation` package provides a **fluent, type-safe validation library** for Go 1.18+. It reduces 20-30 lines of validation boilerplate to just 1-2 lines using method chaining and provides bilingual error messages (English/Korean).

`validation` 패키지는 Go 1.18+ 환경을 위한 **플루언트하고 타입 안전한 검증 라이브러리**를 제공합니다. 메서드 체이닝을 사용하여 20-30줄의 검증 보일러플레이트를 단 1-2줄로 줄이며, 양방향 에러 메시지(영어/한글)를 제공합니다.

### Key Features / 주요 기능

- ✅ **50+ Built-in Validators** / **50개 이상의 내장 검증기**
- ✅ **Fluent API with Method Chaining** / **메서드 체이닝을 통한 플루언트 API**
- ✅ **Type-Safe with Go Generics** / **Go 제네릭을 활용한 타입 안전성**
- ✅ **Bilingual Error Messages (EN/KR)** / **양방향 에러 메시지 (영어/한글)**
- ✅ **Zero External Dependencies** / **외부 의존성 제로**
- ✅ **92.5%+ Test Coverage** / **92.5% 이상의 테스트 커버리지**
- ✅ **Multi-Field Validation** / **다중 필드 검증**
- ✅ **Custom Validators** / **사용자 정의 검증기**
- ✅ **Stop-on-First-Error Support** / **첫 에러에서 멈춤 지원**

---

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/validation
```

**Requirements / 요구사항**: Go 1.18 or higher / Go 1.18 이상

---

## Quick Start / 빠른 시작

### Single Field Validation / 단일 필드 검증

```go
package main

import (
    "fmt"
    "log"
    "github.com/arkd0ng/go-utils/validation"
)

func main() {
    // Simple string validation / 간단한 문자열 검증
    email := "john@example.com"
    v := validation.New(email, "email")
    v.Required().Email().MaxLength(100)

    if err := v.Validate(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Email is valid!")
}
```

### Multi-Field Validation / 다중 필드 검증

```go
type User struct {
    Name  string
    Email string
    Age   int
}

func ValidateUser(user User) error {
    mv := validation.NewValidator()

    mv.Field(user.Name, "name").Required().MinLength(2).MaxLength(50)
    mv.Field(user.Email, "email").Required().Email()
    mv.Field(user.Age, "age").Positive().Min(18).Max(120)

    return mv.Validate()
}
```

---

## Core Concepts / 핵심 개념

### 1. Validator

The `Validator` is the main validation object for a single field.

`Validator`는 단일 필드를 위한 주요 검증 객체입니다.

```go
// Create a validator / 검증기 생성
v := validation.New(value, "fieldName")

// Chain validation rules / 검증 규칙 체이닝
v.Required().MinLength(5).MaxLength(100)

// Execute validation / 검증 실행
err := v.Validate()
```

### 2. MultiValidator

The `MultiValidator` validates multiple fields at once.

`MultiValidator`는 여러 필드를 한 번에 검증합니다.

```go
mv := validation.NewValidator()

mv.Field(user.Name, "name").Required()
mv.Field(user.Email, "email").Email()

err := mv.Validate()
```

### 3. ValidationError

Error structure containing detailed validation failure information.

검증 실패에 대한 상세 정보를 담은 에러 구조체입니다.

```go
type ValidationError struct {
    Field   string      // Field name / 필드 이름
    Value   interface{} // Field value / 필드 값
    Rule    string      // Failed rule name / 실패한 규칙 이름
    Message string      // Error message / 에러 메시지
}
```

### 4. ValidationErrors

Collection of multiple validation errors.

여러 검증 에러의 모음입니다.

```go
type ValidationErrors []ValidationError
```

---

## String Validators / 문자열 검증기

### Basic Validators / 기본 검증기

#### `Required()`
Field must not be empty / 필드가 비어있지 않아야 합니다.

```go
v := validation.New("", "username")
v.Required()
// Error: username is required / username은(는) 필수입니다
```

#### `NotEmpty()`
String must not be empty / 문자열이 비어있지 않아야 합니다.

```go
v := validation.New("", "name")
v.NotEmpty()
// Error: name must not be empty / name은(는) 비어있지 않아야 합니다
```

### Length Validators / 길이 검증기

#### `MinLength(min int)`
String must have minimum length / 문자열이 최소 길이를 가져야 합니다.

```go
v := validation.New("ab", "username")
v.MinLength(3)
// Error: username must be at least 3 characters long
//        username은(는) 최소 3자 이상이어야 합니다
```

#### `MaxLength(max int)`
String must not exceed maximum length / 문자열이 최대 길이를 초과하지 않아야 합니다.

```go
v := validation.New("verylongusername", "username")
v.MaxLength(10)
// Error: username must be at most 10 characters long
//        username은(는) 최대 10자 이하여야 합니다
```

#### `Length(exact int)`
String must have exact length / 문자열이 정확한 길이를 가져야 합니다.

```go
v := validation.New("12345", "zipcode")
v.Length(5)
// Pass! / 통과!
```

### Format Validators / 포맷 검증기

#### `Email()`
Valid email address format / 유효한 이메일 주소 형식

```go
v := validation.New("invalid-email", "email")
v.Email()
// Error: email must be a valid email address
//        email은(는) 유효한 이메일 주소여야 합니다
```

#### `URL()`
Valid URL format / 유효한 URL 형식

```go
v := validation.New("https://example.com", "website")
v.URL()
// Pass! / 통과!
```

#### `UUID()`
Valid UUID format / 유효한 UUID 형식

```go
v := validation.New("550e8400-e29b-41d4-a716-446655440000", "id")
v.UUID()
// Pass! / 통과!
```

#### `JSON()`
Valid JSON format / 유효한 JSON 형식

```go
v := validation.New(`{"name":"John"}`, "data")
v.JSON()
// Pass! / 통과!
```

#### `Base64()`
Valid Base64 encoding / 유효한 Base64 인코딩

```go
v := validation.New("SGVsbG8gV29ybGQ=", "encoded")
v.Base64()
// Pass! / 통과!
```

### Character Type Validators / 문자 타입 검증기

#### `Alpha()`
Only alphabetic characters / 문자만 포함

```go
v := validation.New("abc123", "code")
v.Alpha()
// Error: code must contain only alphabetic characters
//        code은(는) 문자만 포함해야 합니다
```

#### `AlphaNumeric()`
Only alphanumeric characters / 문자와 숫자만 포함

```go
v := validation.New("user123", "username")
v.AlphaNumeric()
// Pass! / 통과!
```

#### `Numeric()`
Only numeric characters / 숫자만 포함

```go
v := validation.New("12345", "pin")
v.Numeric()
// Pass! / 통과!
```

### Case Validators / 대소문자 검증기

#### `Lowercase()`
All characters must be lowercase / 모든 문자가 소문자여야 합니다.

```go
v := validation.New("Hello", "code")
v.Lowercase()
// Error: code must be lowercase / code은(는) 소문자여야 합니다
```

#### `Uppercase()`
All characters must be uppercase / 모든 문자가 대문자여야 합니다.

```go
v := validation.New("HELLO", "code")
v.Uppercase()
// Pass! / 통과!
```

### Pattern Validators / 패턴 검증기

#### `StartsWith(prefix string)`
String must start with prefix / 문자열이 접두사로 시작해야 합니다.

```go
v := validation.New("user_john", "username")
v.StartsWith("user_")
// Pass! / 통과!
```

#### `EndsWith(suffix string)`
String must end with suffix / 문자열이 접미사로 끝나야 합니다.

```go
v := validation.New("document.pdf", "filename")
v.EndsWith(".pdf")
// Pass! / 통과!
```

#### `Contains(substring string)`
String must contain substring / 문자열이 부분 문자열을 포함해야 합니다.

```go
v := validation.New("hello world", "message")
v.Contains("world")
// Pass! / 통과!
```

#### `NotContains(substring string)`
String must not contain substring / 문자열이 부분 문자열을 포함하지 않아야 합니다.

```go
v := validation.New("clean text", "content")
v.NotContains("spam")
// Pass! / 통과!
```

#### `Matches(pattern string)` (also: `Regex`)
String must match regular expression / 문자열이 정규식과 일치해야 합니다.

```go
v := validation.New("abc123", "code")
v.Matches(`^[a-z]+\d+$`)
// Pass! / 통과!
```

---

## Numeric Validators / 숫자 검증기

### Range Validators / 범위 검증기

#### `Min(min float64)`
Number must be at least min / 숫자가 최소값 이상이어야 합니다.

```go
v := validation.New(5, "age")
v.Min(18)
// Error: age must be at least 18 / age은(는) 최소 18 이상이어야 합니다
```

#### `Max(max float64)`
Number must be at most max / 숫자가 최대값 이하여야 합니다.

```go
v := validation.New(150, "age")
v.Max(120)
// Error: age must be at most 120 / age은(는) 최대 120 이하여야 합니다
```

#### `Between(min, max float64)`
Number must be between min and max / 숫자가 최소값과 최대값 사이여야 합니다.

```go
v := validation.New(25, "age")
v.Between(18, 65)
// Pass! / 통과!
```

### Sign Validators / 부호 검증기

#### `Positive()`
Number must be positive (> 0) / 숫자가 양수여야 합니다 (> 0).

```go
v := validation.New(-5, "amount")
v.Positive()
// Error: amount must be positive / amount은(는) 양수여야 합니다
```

#### `Negative()`
Number must be negative (< 0) / 숫자가 음수여야 합니다 (< 0).

```go
v := validation.New(-10, "debt")
v.Negative()
// Pass! / 통과!
```

#### `NonNegative()`
Number must be non-negative (>= 0) / 숫자가 음수가 아니어야 합니다 (>= 0).

```go
v := validation.New(0, "count")
v.NonNegative()
// Pass! / 통과!
```

#### `NonPositive()`
Number must be non-positive (<= 0) / 숫자가 양수가 아니어야 합니다 (<= 0).

```go
v := validation.New(0, "temperature_change")
v.NonPositive()
// Pass! / 통과!
```

### Integer Validators / 정수 검증기

#### `Integer()`
Number must be an integer / 숫자가 정수여야 합니다.

```go
v := validation.New(42.5, "count")
v.Integer()
// Error: count must be an integer / count은(는) 정수여야 합니다
```

#### `Even()`
Number must be even / 숫자가 짝수여야 합니다.

```go
v := validation.New(4, "number")
v.Even()
// Pass! / 통과!
```

#### `Odd()`
Number must be odd / 숫자가 홀수여야 합니다.

```go
v := validation.New(5, "number")
v.Odd()
// Pass! / 통과!
```

---

## Collection Validators / 컬렉션 검증기

### Inclusion Validators / 포함 검증기

#### `In(values ...interface{})`
Value must be in the given list / 값이 주어진 목록에 있어야 합니다.

```go
v := validation.New("red", "color")
v.In("red", "green", "blue")
// Pass! / 통과!
```

#### `NotIn(values ...interface{})`
Value must not be in the given list / 값이 주어진 목록에 없어야 합니다.

```go
v := validation.New("yellow", "color")
v.NotIn("red", "green", "blue")
// Pass! / 통과!
```

### Array/Slice Validators / 배열/슬라이스 검증기

#### `ArrayLength(n int)`
Array/slice must have exact length / 배열/슬라이스가 정확한 길이를 가져야 합니다.

```go
v := validation.New([]string{"a", "b", "c"}, "tags")
v.ArrayLength(3)
// Pass! / 통과!
```

#### `ArrayMinLength(min int)`
Array/slice must have minimum length / 배열/슬라이스가 최소 길이를 가져야 합니다.

```go
v := validation.New([]int{1, 2}, "numbers")
v.ArrayMinLength(3)
// Error: numbers must have at least 3 elements
//        numbers은(는) 최소 3개의 요소를 가져야 합니다
```

#### `ArrayMaxLength(max int)`
Array/slice must not exceed maximum length / 배열/슬라이스가 최대 길이를 초과하지 않아야 합니다.

```go
v := validation.New([]string{"a", "b"}, "tags")
v.ArrayMaxLength(5)
// Pass! / 통과!
```

#### `ArrayNotEmpty()`
Array/slice must not be empty / 배열/슬라이스가 비어있지 않아야 합니다.

```go
v := validation.New([]int{}, "items")
v.ArrayNotEmpty()
// Error: items must not be empty / items은(는) 비어있지 않아야 합니다
```

#### `ArrayUnique()`
Array/slice must contain only unique elements / 배열/슬라이스가 고유한 요소만 포함해야 합니다.

```go
v := validation.New([]int{1, 2, 2, 3}, "numbers")
v.ArrayUnique()
// Error: numbers must contain only unique elements
//        numbers은(는) 고유한 요소만 포함해야 합니다
```

### Map Validators / 맵 검증기

#### `MapHasKey(key string)`
Map must have the specified key / 맵이 지정된 키를 가져야 합니다.

```go
m := map[string]int{"age": 25}
v := validation.New(m, "data")
v.MapHasKey("name")
// Error: data must have key 'name' / data은(는) 'name' 키를 가져야 합니다
```

#### `MapHasKeys(keys ...string)`
Map must have all specified keys / 맵이 모든 지정된 키를 가져야 합니다.

```go
m := map[string]int{"age": 25, "height": 170}
v := validation.New(m, "data")
v.MapHasKeys("age", "height", "weight")
// Error: data must have all keys [age, height, weight]
//        data은(는) 모든 키 [age, height, weight]를 가져야 합니다
```

#### `MapNotEmpty()`
Map must not be empty / 맵이 비어있지 않아야 합니다.

```go
v := validation.New(map[string]int{}, "config")
v.MapNotEmpty()
// Error: config must not be empty / config은(는) 비어있지 않아야 합니다
```

---

## Comparison Validators / 비교 검증기

### Value Comparison / 값 비교

#### `Equals(value interface{})`
Value must equal the given value / 값이 주어진 값과 같아야 합니다.

```go
v := validation.New("password123", "confirmation")
v.Equals("password123")
// Pass! / 통과!
```

#### `NotEquals(value interface{})`
Value must not equal the given value / 값이 주어진 값과 다르아야 합니다.

```go
v := validation.New("newpassword", "password")
v.NotEquals("oldpassword")
// Pass! / 통과!
```

### Numeric Comparison / 숫자 비교

#### `GreaterThan(n float64)`
Number must be greater than n / 숫자가 n보다 커야 합니다.

```go
v := validation.New(10, "score")
v.GreaterThan(5)
// Pass! / 통과!
```

#### `GreaterThanOrEqual(n float64)`
Number must be greater than or equal to n / 숫자가 n보다 크거나 같아야 합니다.

```go
v := validation.New(18, "age")
v.GreaterThanOrEqual(18)
// Pass! / 통과!
```

#### `LessThan(n float64)`
Number must be less than n / 숫자가 n보다 작아야 합니다.

```go
v := validation.New(5, "attempts")
v.LessThan(10)
// Pass! / 통과!
```

#### `LessThanOrEqual(n float64)`
Number must be less than or equal to n / 숫자가 n보다 작거나 같아야 합니다.

```go
v := validation.New(100, "percentage")
v.LessThanOrEqual(100)
// Pass! / 통과!
```

### Time Comparison / 시간 비교

#### `Before(t time.Time)`
Time must be before the given time / 시간이 주어진 시간 이전이어야 합니다.

```go
now := time.Now()
past := now.Add(-24 * time.Hour)

v := validation.New(past, "startDate")
v.Before(now)
// Pass! / 통과!
```

#### `After(t time.Time)`
Time must be after the given time / 시간이 주어진 시간 이후여야 합니다.

```go
now := time.Now()
future := now.Add(24 * time.Hour)

v := validation.New(future, "endDate")
v.After(now)
// Pass! / 통과!
```

#### `BeforeOrEqual(t time.Time)`
Time must be before or equal to the given time / 시간이 주어진 시간 이전이거나 같아야 합니다.

```go
v := validation.New(time.Now(), "deadline")
v.BeforeOrEqual(time.Now())
// Pass! / 통과!
```

#### `AfterOrEqual(t time.Time)`
Time must be after or equal to the given time / 시간이 주어진 시간 이후이거나 같아야 합니다.

```go
v := validation.New(time.Now(), "startDate")
v.AfterOrEqual(time.Now().Add(-1 * time.Hour))
// Pass! / 통과!
```

---

## Advanced Features / 고급 기능

### Stop on First Error / 첫 에러에서 중지

By default, validators collect all errors. Use `StopOnError()` to stop at the first failure.

기본적으로 검증기는 모든 에러를 수집합니다. 첫 실패에서 멈추려면 `StopOnError()`를 사용하세요.

```go
v := validation.New("", "email")
v.StopOnError().
    Required().        // Fails here, stops validation
    Email().           // Not executed
    MaxLength(100)     // Not executed

err := v.Validate()
// Only returns "email is required" error
// "email is required" 에러만 반환
```

### Custom Error Messages / 사용자 정의 에러 메시지

Override default error messages with `WithMessage()`.

`WithMessage()`로 기본 에러 메시지를 덮어쓸 수 있습니다.

```go
v := validation.New(user.Age, "age")
v.Min(18).WithMessage("You must be at least 18 years old to register")
v.Max(120).WithMessage("Please enter a valid age")

err := v.Validate()
```

### Custom Validators / 사용자 정의 검증기

Create custom validation logic with `Custom()`.

`Custom()`으로 사용자 정의 검증 로직을 만들 수 있습니다.

```go
v := validation.New(password, "password")

// Must contain special character / 특수 문자 포함 필수
v.Custom(func(val interface{}) bool {
    s := val.(string)
    return strings.ContainsAny(s, "!@#$%^&*()")
}, "Password must contain at least one special character")

// Must not contain username / 사용자명 포함 불가
v.Custom(func(val interface{}) bool {
    pwd := val.(string)
    return !strings.Contains(pwd, username)
}, "Password must not contain your username")
```

### Multi-Field Validation / 다중 필드 검증

Validate multiple fields together using `MultiValidator`.

`MultiValidator`로 여러 필드를 함께 검증할 수 있습니다.

```go
type UserRegistration struct {
    Username        string
    Email           string
    Password        string
    ConfirmPassword string
    Age             int
    Country         string
    Terms           bool
}

func ValidateRegistration(reg UserRegistration) error {
    mv := validation.NewValidator()

    // Username validation / 사용자명 검증
    mv.Field(reg.Username, "username").
        Required().
        MinLength(3).
        MaxLength(20).
        AlphaNumeric()

    // Email validation / 이메일 검증
    mv.Field(reg.Email, "email").
        Required().
        Email().
        MaxLength(100)

    // Password validation / 비밀번호 검증
    mv.Field(reg.Password, "password").
        Required().
        MinLength(8).
        MaxLength(100).
        Matches(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])`)

    // Password confirmation / 비밀번호 확인
    mv.Field(reg.ConfirmPassword, "confirm_password").
        Required().
        Equals(reg.Password).WithMessage("Passwords do not match")

    // Age validation / 나이 검증
    mv.Field(reg.Age, "age").
        Positive().
        Between(13, 120)

    // Country validation / 국가 검증
    mv.Field(reg.Country, "country").
        Required().
        In("US", "KR", "JP", "CN", "UK", "FR", "DE")

    // Terms acceptance / 약관 동의
    mv.Field(reg.Terms, "terms").
        Equals(true).WithMessage("You must accept the terms and conditions")

    return mv.Validate()
}
```

---

## Error Handling / 에러 처리

### Basic Error Handling / 기본 에러 처리

```go
err := mv.Validate()
if err != nil {
    // Type assertion / 타입 단언
    validationErrs := err.(validation.ValidationErrors)

    // Iterate through errors / 에러 순회
    for _, e := range validationErrs {
        fmt.Printf("Field: %s\n", e.Field)
        fmt.Printf("Value: %v\n", e.Value)
        fmt.Printf("Rule: %s\n", e.Rule)
        fmt.Printf("Message: %s\n\n", e.Message)
    }
}
```

### ValidationErrors Helper Methods / ValidationErrors 헬퍼 메서드

#### `HasField(fieldName string) bool`
Check if a specific field has errors / 특정 필드에 에러가 있는지 확인

```go
if validationErrs.HasField("email") {
    fmt.Println("Email validation failed")
}
```

#### `GetField(fieldName string) []ValidationError`
Get all errors for a specific field / 특정 필드의 모든 에러 가져오기

```go
emailErrors := validationErrs.GetField("email")
for _, e := range emailErrors {
    fmt.Println(e.Message)
}
```

#### `First() ValidationError`
Get the first error / 첫 번째 에러 가져오기

```go
firstError := validationErrs.First()
fmt.Println(firstError.Message)
```

#### `Count() int`
Get total number of errors / 총 에러 개수 가져오기

```go
count := validationErrs.Count()
fmt.Printf("Total errors: %d\n", count)
```

#### `ToMap() map[string][]string`
Convert errors to map format / 에러를 맵 형식으로 변환

```go
errMap := validationErrs.ToMap()
// {
//   "email": ["email must be a valid email address"],
//   "age": ["age must be at least 18"]
// }
```

### HTTP API Error Response / HTTP API 에러 응답

```go
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var req UserRegistration
    json.NewDecoder(r.Body).Decode(&req)

    if err := ValidateRegistration(req); err != nil {
        validationErrs := err.(validation.ValidationErrors)

        response := map[string]interface{}{
            "error": "Validation failed",
            "fields": validationErrs.ToMap(),
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    // Proceed with registration...
}
```

---

## Real-World Examples / 실제 사용 예제

### Example 1: User Profile Update / 사용자 프로필 업데이트

```go
type UserProfile struct {
    Bio         string
    Website     string
    Location    string
    Age         int
    Skills      []string
    SocialLinks map[string]string
}

func ValidateProfile(profile UserProfile) error {
    mv := validation.NewValidator()

    mv.Field(profile.Bio, "bio").
        MaxLength(500)

    mv.Field(profile.Website, "website").
        URL()

    mv.Field(profile.Location, "location").
        MaxLength(100)

    mv.Field(profile.Age, "age").
        Positive().
        Between(13, 120)

    mv.Field(profile.Skills, "skills").
        ArrayMaxLength(10).
        ArrayUnique()

    mv.Field(profile.SocialLinks, "social_links").
        MapHasKeys("twitter", "github")

    return mv.Validate()
}
```

### Example 2: E-commerce Order Validation / 전자상거래 주문 검증

```go
type Order struct {
    CustomerEmail string
    Items         []OrderItem
    ShippingAddr  Address
    PaymentMethod string
    CouponCode    string
    TotalAmount   float64
}

type OrderItem struct {
    ProductID string
    Quantity  int
}

type Address struct {
    Street     string
    City       string
    PostalCode string
    Country    string
}

func ValidateOrder(order Order) error {
    mv := validation.NewValidator()

    // Customer validation / 고객 검증
    mv.Field(order.CustomerEmail, "customer_email").
        Required().
        Email()

    // Items validation / 상품 검증
    mv.Field(order.Items, "items").
        ArrayNotEmpty().
        ArrayMinLength(1).
        ArrayMaxLength(50)

    // Address validation / 주소 검증
    mv.Field(order.ShippingAddr.Street, "shipping_street").
        Required().
        MinLength(5).
        MaxLength(200)

    mv.Field(order.ShippingAddr.City, "shipping_city").
        Required().
        MinLength(2).
        MaxLength(100)

    mv.Field(order.ShippingAddr.PostalCode, "shipping_postal_code").
        Required().
        Matches(`^\d{5}(-\d{4})?$`)

    mv.Field(order.ShippingAddr.Country, "shipping_country").
        Required().
        In("US", "KR", "JP", "CN", "UK", "FR", "DE")

    // Payment method / 결제 방법
    mv.Field(order.PaymentMethod, "payment_method").
        Required().
        In("credit_card", "paypal", "bank_transfer")

    // Total amount / 총 금액
    mv.Field(order.TotalAmount, "total_amount").
        Positive().
        Min(0.01)

    return mv.Validate()
}
```

### Example 3: Configuration File Validation / 설정 파일 검증

```go
type AppConfig struct {
    ServerPort      int
    ServerHost      string
    DatabaseURL     string
    RedisURL        string
    JWTSecret       string
    AllowedOrigins  []string
    RateLimitPerMin int
    Features        map[string]bool
    LogLevel        string
}

func ValidateConfig(cfg AppConfig) error {
    mv := validation.NewValidator()

    mv.Field(cfg.ServerPort, "server_port").
        Positive().
        Between(1, 65535)

    mv.Field(cfg.ServerHost, "server_host").
        Required().
        URL()

    mv.Field(cfg.DatabaseURL, "database_url").
        Required().
        StartsWith("postgres://")

    mv.Field(cfg.RedisURL, "redis_url").
        Required().
        StartsWith("redis://")

    mv.Field(cfg.JWTSecret, "jwt_secret").
        Required().
        MinLength(32).
        MaxLength(256)

    mv.Field(cfg.AllowedOrigins, "allowed_origins").
        ArrayNotEmpty().
        ArrayUnique()

    mv.Field(cfg.RateLimitPerMin, "rate_limit").
        Positive().
        Between(1, 10000)

    mv.Field(cfg.Features, "features").
        MapNotEmpty().
        MapHasKeys("auth", "logging", "metrics")

    mv.Field(cfg.LogLevel, "log_level").
        Required().
        In("debug", "info", "warn", "error")

    return mv.Validate()
}
```

---

## Best Practices / 모범 사례

### 1. Use Multi-Field Validation for Complex Objects / 복잡한 객체에 다중 필드 검증 사용

```go
// Good ✅
func ValidateUser(user User) error {
    mv := validation.NewValidator()
    mv.Field(user.Name, "name").Required()
    mv.Field(user.Email, "email").Email()
    return mv.Validate()
}

// Avoid ❌
func ValidateUser(user User) error {
    v1 := validation.New(user.Name, "name").Required()
    if err := v1.Validate(); err != nil {
        return err
    }
    v2 := validation.New(user.Email, "email").Email()
    return v2.Validate()
}
```

### 2. Use StopOnError for Performance / 성능을 위해 StopOnError 사용

```go
// If subsequent validations are expensive
// 후속 검증이 비용이 많이 드는 경우
v := validation.New(data, "data").
    StopOnError().
    Required().              // Quick check
    JSON().                  // Moderate check
    Custom(expensiveCheck)   // Expensive check (only if above pass)
```

### 3. Create Reusable Validation Functions / 재사용 가능한 검증 함수 생성

```go
// Reusable password validator / 재사용 가능한 비밀번호 검증기
func ValidatePassword(password string, fieldName string) *validation.Validator {
    v := validation.New(password, fieldName)
    return v.Required().
        MinLength(8).
        MaxLength(100).
        Matches(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])`)
}

// Usage / 사용
mv := validation.NewValidator()
ValidatePassword(user.Password, "password")
ValidatePassword(user.NewPassword, "new_password")
```

### 4. Use WithMessage for User-Friendly Errors / 사용자 친화적 에러를 위해 WithMessage 사용

```go
v := validation.New(age, "age")
v.Min(18).WithMessage("You must be at least 18 years old to register")
v.Max(120).WithMessage("Please enter a valid age (maximum 120)")
```

### 5. Validate at Service Boundary / 서비스 경계에서 검증

```go
// HTTP Handler / HTTP 핸들러
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    json.NewDecoder(r.Body).Decode(&req)

    // Validate at entry point / 진입점에서 검증
    if err := ValidateCreateUser(req); err != nil {
        handleValidationError(w, err)
        return
    }

    // Proceed with business logic / 비즈니스 로직 진행
    user, err := userService.Create(req)
    // ...
}
```

---

## Troubleshooting / 문제 해결

### Problem: Type Mismatch Errors / 문제: 타입 불일치 에러

```go
// Wrong ❌
v := validation.New(123, "code")
v.MinLength(5) // Error: code must be a string

// Correct ✅
v := validation.New("123", "code")
v.MinLength(5) // OK
```

### Problem: Custom Validator Not Working / 문제: 사용자 정의 검증기가 작동하지 않음

```go
// Wrong ❌ - Always returns true
v.Custom(func(val interface{}) bool {
    return true // Never fails
}, "Must contain special char")

// Correct ✅
v.Custom(func(val interface{}) bool {
    s := val.(string)
    return strings.ContainsAny(s, "!@#$%")
}, "Must contain special char")
```

### Problem: StopOnError Not Stopping / 문제: StopOnError가 멈추지 않음

```go
// Wrong ❌ - StopOnError must be first
v := validation.New("", "email")
v.Required().StopOnError().Email()

// Correct ✅
v := validation.New("", "email")
v.StopOnError().Required().Email()
```

### Problem: Validation Errors Not Appearing / 문제: 검증 에러가 나타나지 않음

```go
// Wrong ❌ - Forgot to call Validate()
v := validation.New("", "name")
v.Required()
// Missing: err := v.Validate()

// Correct ✅
v := validation.New("", "name")
v.Required()
err := v.Validate() // Must call Validate()
```

---

## Performance Tips / 성능 팁

1. **Use StopOnError for Sequential Validation** / **순차 검증에 StopOnError 사용**
   - Stops at first failure, avoiding unnecessary checks
   - 첫 실패에서 멈춰 불필요한 검사 회피

2. **Compile Regex Once** / **정규식 한 번만 컴파일**
   ```go
   // Regex is compiled internally and cached
   // 정규식은 내부적으로 컴파일되고 캐시됨
   v.Matches(`^[a-z]+$`)
   ```

3. **Avoid Reflection When Possible** / **가능하면 리플렉션 회피**
   - Use specific validators instead of generic Custom()
   - 제네릭 Custom() 대신 특정 검증기 사용

4. **Batch Field Validations** / **필드 검증 일괄 처리**
   - Use MultiValidator to collect all errors at once
   - MultiValidator를 사용해 모든 에러를 한 번에 수집

---

## Conclusion / 결론

The `validation` package provides a powerful, flexible, and type-safe way to validate data in Go applications. With 50+ built-in validators, fluent API, and bilingual error messages, it significantly reduces boilerplate code while improving code readability and maintainability.

`validation` 패키지는 Go 애플리케이션에서 데이터를 검증하는 강력하고 유연하며 타입 안전한 방법을 제공합니다. 50개 이상의 내장 검증기, 플루언트 API, 양방향 에러 메시지를 통해 보일러플레이트 코드를 크게 줄이고 코드 가독성과 유지보수성을 향상시킵니다.

For more information, see:
- [Package README](../../validation/README.md)
- [Developer Guide](DEVELOPER_GUIDE.md)
- [Executable Examples](../../examples/validation/main.go)

자세한 정보는 다음을 참조하세요:
- [패키지 README](../../validation/README.md)
- [개발자 가이드](DEVELOPER_GUIDE.md)
- [실행 가능한 예제](../../examples/validation/main.go)

---

**Last Updated / 최종 업데이트**: 2025-10-17
**Version / 버전**: v1.13.013
**License / 라이선스**: MIT
