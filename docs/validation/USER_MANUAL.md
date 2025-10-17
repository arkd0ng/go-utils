# Validation Package - User Manual / Validation 패키지 - 사용자 매뉴얼

**Version / 버전**: v1.13.024
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
9. [Network Validators / 네트워크 검증기](#network-validators--네트워크-검증기)
10. [DateTime Validators / 날짜/시간 검증기](#datetime-validators--날짜시간-검증기)
11. [Range Validators / 범위 검증기](#range-validators--범위-검증기)
12. [File Validators / 파일 검증기](#file-validators--파일-검증기)
13. [Credit Card Validators / 신용카드 검증기](#credit-card-validators--신용카드-검증기)
14. [Business/ID Validators / 비즈니스/ID 검증기](#businessid-validators--비즈니스id-검증기)
15. [Geographic Validators / 지리 좌표 검증기](#geographic-validators--지리-좌표-검증기)
16. [Security Validators / 보안 검증기](#security-validators--보안-검증기) 🆕
17. [Advanced Features / 고급 기능](#advanced-features--고급-기능)
18. [Error Handling / 에러 처리](#error-handling--에러-처리)
19. [Real-World Examples / 실제 사용 예제](#real-world-examples--실제-사용-예제)
20. [Best Practices / 모범 사례](#best-practices--모범-사례)
21. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)

---

## Introduction / 소개

The `validation` package provides a **fluent, type-safe validation library** for Go 1.18+. It reduces 20-30 lines of validation boilerplate to just 1-2 lines using method chaining and provides bilingual error messages (English/Korean).

`validation` 패키지는 Go 1.18+ 환경을 위한 **플루언트하고 타입 안전한 검증 라이브러리**를 제공합니다. 메서드 체이닝을 사용하여 20-30줄의 검증 보일러플레이트를 단 1-2줄로 줄이며, 양방향 에러 메시지(영어/한글)를 제공합니다.

### Key Features / 주요 기능

- ✅ **85+ Built-in Validators** / **85개 이상의 내장 검증기**
- ✅ **Fluent API with Method Chaining** / **메서드 체이닝을 통한 플루언트 API**
- ✅ **Type-Safe with Go Generics** / **Go 제네릭을 활용한 타입 안전성**
- ✅ **Bilingual Error Messages (EN/KR)** / **양방향 에러 메시지 (영어/한글)**
- ✅ **Zero External Dependencies** / **외부 의존성 제로**
- ✅ **100% Test Coverage** / **100% 테스트 커버리지**
- ✅ **Multi-Field Validation** / **다중 필드 검증**
- ✅ **Custom Validators** / **사용자 정의 검증기**
- ✅ **Stop-on-First-Error Support** / **첫 에러에서 멈춤 지원**
- ✅ **Network Validators (IPv4, IPv6, CIDR, MAC)** / **네트워크 검증기**
- ✅ **DateTime Validators (DateFormat, TimeFormat, DateBefore, DateAfter)** / **날짜/시간 검증기**
- ✅ **Range Validators (IntRange, FloatRange, DateRange)** / **범위 검증기**
- ✅ **Format Validators (UUIDv4, XML, Hex)** / **포맷 검증기**
- ✅ **File Validators (FilePath, FileExists, FileReadable, FileWritable, FileSize, FileExtension)** / **파일 검증기**
- ✅ **Credit Card Validators (CreditCard, CreditCardType, Luhn)** / **신용카드 검증기**
- ✅ **Business/ID Validators (ISBN, ISSN, EAN)** / **비즈니스/ID 검증기**
- ✅ **Geographic Validators (Latitude, Longitude, Coordinate)** / **지리 좌표 검증기**
- ✅ **Security Validators (JWT, BCrypt, MD5, SHA1, SHA256, SHA512)** 🆕 / **보안 검증기** 🆕

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

#### `UUIDv4()` 🆕
Valid UUID version 4 format / 유효한 UUID 버전 4 형식

```go
v := validation.New("550e8400-e29b-41d4-a716-446655440000", "request_id")
v.UUIDv4()
// Pass! / 통과!

v2 := validation.New("6ba7b810-9dad-11d1-80b4-00c04fd430c8", "id")
v2.UUIDv4()
// Error: id must be a valid UUID v4 (this is UUID v1)
//        id은(는) 유효한 UUID v4여야 합니다 (이것은 UUID v1입니다)
```

#### `XML()` 🆕
Valid XML format / 유효한 XML 형식

```go
xmlData := `<?xml version="1.0"?>
<person>
    <name>John Doe</name>
    <age>30</age>
</person>`

v := validation.New(xmlData, "user_data")
v.XML()
// Pass! / 통과!
```

#### `Hex()` 🆕
Valid hexadecimal format / 유효한 16진수 형식

```go
v := validation.New("0xdeadbeef", "color_code")
v.Hex()
// Pass! / 통과!

v2 := validation.New("ABCD1234", "hash")
v2.Hex()
// Pass! (0x prefix is optional / 0x 접두사는 선택사항)
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

## Network Validators / 네트워크 검증기

**New in v1.13.016** 🆕 / **v1.13.016의 새 기능** 🆕

Network validators validate IP addresses, CIDR notation, and MAC addresses using Go's standard `net` package.

네트워크 검증기는 Go의 표준 `net` 패키지를 사용하여 IP 주소, CIDR 표기법, MAC 주소를 검증합니다.

### IPv4()

Validates IPv4 addresses only. / IPv4 주소만 검증합니다.

**Validation Rules** / **검증 규칙**:
- Must be valid IPv4 format (xxx.xxx.xxx.xxx)
- Each octet must be 0-255
- No leading zeros (except 0 itself)

**Valid Examples** / **유효한 예시**:
```go
v := validation.New("192.168.1.1", "server_ip")
v.IPv4()
// Pass! / 통과!

v := validation.New("10.0.0.1", "gateway")
v.IPv4()
// Pass! / 통과!

v := validation.New("255.255.255.255", "broadcast")
v.IPv4()
// Pass! / 통과!
```

**Invalid Examples** / **잘못된 예시**:
```go
v := validation.New("256.1.1.1", "ip")
v.IPv4()
// Fail: octet > 255 / 실패: 옥텟이 255보다 큼

v := validation.New("192.168.1", "ip")
v.IPv4()
// Fail: incomplete / 실패: 불완전

v := validation.New("2001:db8::1", "ip")
v.IPv4()
// Fail: this is IPv6 / 실패: IPv6임
```

### IPv6()

Validates IPv6 addresses only. / IPv6 주소만 검증합니다.

**Validation Rules** / **검증 규칙**:
- Must be valid IPv6 format
- Supports compressed notation (::)
- Supports full and partial addresses

**Valid Examples** / **유효한 예시**:
```go
v := validation.New("2001:0db8:85a3:0000:0000:8a2e:0370:7334", "ipv6")
v.IPv6()
// Pass! Full format / 통과! 전체 형식

v := validation.New("2001:db8:85a3::8a2e:370:7334", "ipv6")
v.IPv6()
// Pass! Compressed format / 통과! 압축 형식

v := validation.New("::1", "loopback")
v.IPv6()
// Pass! IPv6 loopback / 통과! IPv6 루프백

v := validation.New("fe80::1", "link_local")
v.IPv6()
// Pass! Link-local address / 통과! 링크-로컬 주소
```

**Invalid Examples** / **잘못된 예시**:
```go
v := validation.New("192.168.1.1", "ip")
v.IPv6()
// Fail: this is IPv4 / 실패: IPv4임

v := validation.New("gggg::1", "ip")
v.IPv6()
// Fail: invalid hex / 실패: 잘못된 16진수

v := validation.New("2001:db8::1::2", "ip")
v.IPv6()
// Fail: double :: / 실패: :: 중복
```

### IP()

Validates both IPv4 and IPv6 addresses. / IPv4와 IPv6 주소 모두 검증합니다.

**Use this when** / **다음의 경우 사용**:
- You want to accept both IPv4 and IPv6 / IPv4와 IPv6를 모두 허용하려는 경우
- IP version doesn't matter / IP 버전이 중요하지 않은 경우

**Examples** / **예시**:
```go
v := validation.New("192.168.1.1", "ip")
v.IP()
// Pass! IPv4 accepted / 통과! IPv4 허용됨

v := validation.New("2001:db8::1", "ip")
v.IP()
// Pass! IPv6 accepted / 통과! IPv6 허용됨

v := validation.New("not-an-ip", "ip")
v.IP()
// Fail: invalid format / 실패: 잘못된 형식
```

### CIDR()

Validates CIDR notation (IP address with prefix length). / CIDR 표기법(접두사 길이가 있는 IP 주소)을 검증합니다.

**Validation Rules** / **검증 규칙**:
- Format: `<IP>/<prefix>`
- IP can be IPv4 or IPv6
- Prefix must be valid:
  - IPv4: 0-32
  - IPv6: 0-128

**Valid Examples** / **유효한 예시**:
```go
v := validation.New("192.168.1.0/24", "subnet")
v.CIDR()
// Pass! Common private network / 통과! 일반적인 사설 네트워크

v := validation.New("10.0.0.0/8", "network")
v.CIDR()
// Pass! Class A private network / 통과! 클래스 A 사설 네트워크

v := validation.New("192.168.1.1/32", "host")
v.CIDR()
// Pass! Single host / 통과! 단일 호스트

v := validation.New("2001:db8::/32", "ipv6_network")
v.CIDR()
// Pass! IPv6 network / 통과! IPv6 네트워크
```

**Invalid Examples** / **잘못된 예시**:
```go
v := validation.New("192.168.1.0", "network")
v.CIDR()
// Fail: missing prefix / 실패: 접두사 누락

v := validation.New("192.168.1.0/33", "network")
v.CIDR()
// Fail: prefix > 32 for IPv4 / 실패: IPv4의 경우 접두사가 32보다 큼

v := validation.New("invalid/24", "network")
v.CIDR()
// Fail: invalid IP / 실패: 잘못된 IP
```

### MAC()

Validates MAC (Media Access Control) addresses. / MAC(미디어 액세스 제어) 주소를 검증합니다.

**Supported Formats** / **지원되는 형식**:
- Colon-separated: `00:1A:2B:3C:4D:5E`
- Hyphen-separated: `00-1A-2B-3C-4D-5E`
- Dot-separated (Cisco): `001A.2B3C.4D5E`
- Case-insensitive / 대소문자 구분 안 함

**Valid Examples** / **유효한 예시**:
```go
v := validation.New("00:1A:2B:3C:4D:5E", "mac")
v.MAC()
// Pass! Colon-separated uppercase / 통과! 콜론 구분 대문자

v := validation.New("00-1a-2b-3c-4d-5e", "mac")
v.MAC()
// Pass! Hyphen-separated lowercase / 통과! 하이픈 구분 소문자

v := validation.New("001A.2B3C.4D5E", "mac")
v.MAC()
// Pass! Cisco dot format / 통과! Cisco 점 형식

v := validation.New("FF:FF:FF:FF:FF:FF", "broadcast_mac")
v.MAC()
// Pass! Broadcast MAC / 통과! 브로드캐스트 MAC
```

**Invalid Examples** / **잘못된 예시**:
```go
v := validation.New("00:1A:2B:3C:4D", "mac")
v.MAC()
// Fail: too short / 실패: 너무 짧음

v := validation.New("GG:1A:2B:3C:4D:5E", "mac")
v.MAC()
// Fail: invalid hex / 실패: 잘못된 16진수

v := validation.New("00:1A:2B:3C:4D:5E:6F", "mac")
v.MAC()
// Fail: too long / 실패: 너무 김
```

### Common Use Cases / 일반적인 사용 사례

#### API Endpoint IP Filtering / API 엔드포인트 IP 필터링

```go
type APIConfig struct {
    AllowedIPs []string
    Subnet     string
}

func ValidateAPIConfig(config APIConfig) error {
    mv := validation.NewValidator()

    // Validate subnet
    mv.Field(config.Subnet, "subnet").Required().CIDR()

    // Validate each allowed IP
    for i, ip := range config.AllowedIPs {
        fieldName := fmt.Sprintf("allowed_ips[%d]", i)
        mv.Field(ip, fieldName).Required().IP()
    }

    return mv.Validate()
}
```

#### Network Device Configuration / 네트워크 장치 구성

```go
type NetworkDevice struct {
    IPAddress  string
    Gateway    string
    Subnet     string
    MACAddress string
}

func ValidateNetworkDevice(device NetworkDevice) error {
    mv := validation.NewValidator()

    mv.Field(device.IPAddress, "ip_address").Required().IPv4()
    mv.Field(device.Gateway, "gateway").Required().IPv4()
    mv.Field(device.Subnet, "subnet").Required().CIDR()
    mv.Field(device.MACAddress, "mac_address").Required().MAC()

    return mv.Validate()
}
```

#### Firewall Rule Validation / 방화벽 규칙 검증

```go
type FirewallRule struct {
    SourceIP      string
    DestinationIP string
    Network       string
}

func ValidateFirewallRule(rule FirewallRule) error {
    mv := validation.NewValidator()

    // Source and destination can be any IP (v4 or v6)
    mv.Field(rule.SourceIP, "source_ip").Required().IP()
    mv.Field(rule.DestinationIP, "destination_ip").Required().IP()

    // Network must be CIDR notation
    mv.Field(rule.Network, "network").Required().CIDR()

    return mv.Validate()
}
```

### Performance Characteristics / 성능 특성

Network validators use Go's standard `net` package which is highly optimized:

네트워크 검증기는 고도로 최적화된 Go의 표준 `net` 패키지를 사용합니다:

| Validator | Avg Time | Description |
|-----------|----------|-------------|
| IPv4() | ~29 ns/op | Very fast, simple parsing / 매우 빠름, 단순 파싱 |
| IPv6() | ~92 ns/op | Fast, handles compression / 빠름, 압축 처리 |
| IP() | ~24 ns/op | Fastest, accepts both / 가장 빠름, 둘 다 허용 |
| CIDR() | ~145 ns/op | Slightly slower, parses prefix / 약간 느림, 접두사 파싱 |
| MAC() | ~64 ns/op | Fast, multiple format support / 빠름, 여러 형식 지원 |

### Tips and Best Practices / 팁 및 모범 사례

1. **Use Specific Validators When Possible** / **가능한 한 특정 검증기 사용**
   ```go
   // Good: Specific requirement
   v.IPv4()  // If you only accept IPv4

   // Less specific: May accept unwanted formats
   v.IP()    // Accepts both IPv4 and IPv6
   ```

2. **Validate CIDR for Network Configuration** / **네트워크 구성에 CIDR 검증**
   ```go
   // Always use CIDR for subnets and network ranges
   v.CIDR()  // Ensures proper network notation with prefix
   ```

3. **MAC Address Case Doesn't Matter** / **MAC 주소 대소문자는 중요하지 않음**
   ```go
   // All valid, case-insensitive
   "00:1A:2B:3C:4D:5E"  // Uppercase
   "00:1a:2b:3c:4d:5e"  // Lowercase
   "00:1a:2B:3C:4d:5E"  // Mixed
   ```

4. **Combine with Other Validators** / **다른 검증기와 결합**
   ```go
   v := validation.New(serverIP, "server_ip")
   v.Required().IPv4().
       Custom(func(val interface{}) bool {
           // Additional business logic
           ip := val.(string)
           return !strings.HasPrefix(ip, "127.")  // Reject localhost
       }, "Server IP cannot be localhost")
   ```

---

### DateTime Validators / 날짜/시간 검증기

DateTime validators validate date and time formats and ranges.

DateTime 검증기는 날짜 및 시간 형식과 범위를 검증합니다.

#### Available Validators / 사용 가능한 검증기

| Validator | Description | 설명 |
|-----------|-------------|------|
| `DateFormat(format)` | Validates date string format | 날짜 문자열 형식 검증 |
| `TimeFormat(format)` | Validates time string format | 시간 문자열 형식 검증 |
| `DateBefore(time)` | Validates date is before specified time | 지정된 시간 이전인지 검증 |
| `DateAfter(time)` | Validates date is after specified time | 지정된 시간 이후인지 검증 |

#### DateFormat(format) - Date Format Validation / 날짜 형식 검증

Validates that a string matches a specific date format using Go's time.Parse format.

Go의 time.Parse 형식을 사용하여 문자열이 특정 날짜 형식과 일치하는지 검증합니다.

**Validation Rules** / **검증 규칙**:
- Value must be a string / 값은 문자열이어야 함
- Must match the specified format exactly / 지정된 형식과 정확히 일치해야 함
- Date must be valid (e.g., no Feb 30) / 날짜가 유효해야 함 (예: 2월 30일 불가)

**Examples** / **예제**:

```go
// ISO 8601 format (YYYY-MM-DD)
v := validation.New("2025-10-17", "birth_date")
v.DateFormat("2006-01-02")
// Valid: "2025-10-17", "2025-01-01"
// Invalid: "10/17/2025", "2025-13-01", "not-a-date"

// US format (MM/DD/YYYY)
v := validation.New("10/17/2025", "event_date")
v.DateFormat("01/02/2006")
// Valid: "10/17/2025", "01/31/2025"
// Invalid: "2025-10-17", "13/01/2025"

// EU format (DD/MM/YYYY)
v := validation.New("17/10/2025", "meeting_date")
v.DateFormat("02/01/2006")
// Valid: "17/10/2025", "31/12/2025"
// Invalid: "10/17/2025", "32/01/2025"
```

#### TimeFormat(format) - Time Format Validation / 시간 형식 검증

Validates that a string matches a specific time format.

문자열이 특정 시간 형식과 일치하는지 검증합니다.

**Validation Rules** / **검증 규칙**:
- Value must be a string / 값은 문자열이어야 함
- Must match the specified format exactly / 지정된 형식과 정확히 일치해야 함
- Time components must be valid / 시간 구성요소가 유효해야 함

**Examples** / **예제**:

```go
// 24-hour format (HH:MM:SS)
v := validation.New("14:30:00", "meeting_time")
v.TimeFormat("15:04:05")
// Valid: "14:30:00", "00:00:00", "23:59:59"
// Invalid: "2:30 PM", "25:00:00", "14:60:00"

// 24-hour format without seconds (HH:MM)
v := validation.New("14:30", "start_time")
v.TimeFormat("15:04")
// Valid: "14:30", "00:00", "23:59"
// Invalid: "14:30:00", "2:30 PM"

// 12-hour format (hh:MM:SS AM/PM)
v := validation.New("02:30:00 PM", "appointment")
v.TimeFormat("03:04:05 PM")
// Valid: "02:30:00 PM", "11:59:59 AM"
// Invalid: "14:30:00", "13:00:00 PM"
```

#### DateBefore(time) - Date Before Validation / 날짜 이전 검증

Validates that a date is before the specified time.

날짜가 지정된 시간 이전인지 검증합니다.

**Supported Input Types** / **지원되는 입력 타입**:
- `time.Time` object / time.Time 객체
- RFC3339 string: `"2006-01-02T15:04:05Z07:00"`
- ISO 8601 string: `"2006-01-02"`

**Examples** / **예제**:

```go
// Using time.Time
maxDate := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)
testDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)
v := validation.New(testDate, "expiry_date")
v.DateBefore(maxDate)
// Valid: any date before 2025-12-31
// Invalid: 2025-12-31 or later

// Using RFC3339 string
v := validation.New("2025-10-17T12:00:00Z", "deadline")
v.DateBefore(time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC))

// Using ISO 8601 string
v := validation.New("2025-10-17", "event_date")
v.DateBefore(time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC))
```

#### DateAfter(time) - Date After Validation / 날짜 이후 검증

Validates that a date is after the specified time.

날짜가 지정된 시간 이후인지 검증합니다.

**Supported Input Types** / **지원되는 입력 타입**:
- `time.Time` object / time.Time 객체
- RFC3339 string: `"2006-01-02T15:04:05Z07:00"`
- ISO 8601 string: `"2006-01-02"`

**Examples** / **예제**:

```go
// Using time.Time
minDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
testDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)
v := validation.New(testDate, "start_date")
v.DateAfter(minDate)
// Valid: any date after 2025-01-01
// Invalid: 2025-01-01 or earlier

// Using RFC3339 string
v := validation.New("2025-10-17T12:00:00Z", "publish_date")
v.DateAfter(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))

// Using ISO 8601 string
v := validation.New("2025-10-17", "launch_date")
v.DateAfter(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
```

#### Common Use Cases / 일반적인 사용 사례

**1. Event Scheduling Validation / 이벤트 일정 검증**

```go
minDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
maxDate := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)

mv := validation.NewValidator()
mv.Field("2025-10-17", "event_date").Required().DateFormat("2006-01-02")
mv.Field("14:30:00", "event_time").Required().TimeFormat("15:04:05")
mv.Field(eventDateTime, "event_datetime").DateAfter(minDate).DateBefore(maxDate)
```

**2. User Registration (Birth Date) / 사용자 등록 (생년월일)**

```go
minAge := time.Now().AddDate(-120, 0, 0)  // Max 120 years old
maxAge := time.Now().AddDate(-18, 0, 0)   // Min 18 years old

mv := validation.NewValidator()
mv.Field("1990-05-15", "birth_date").
    Required().
    DateFormat("2006-01-02").
    DateAfter(minAge).
    DateBefore(maxAge)
```

**3. Booking System (Date Range) / 예약 시스템 (날짜 범위)**

```go
now := time.Now()
minBooking := now.AddDate(0, 0, 1)   // Tomorrow
maxBooking := now.AddDate(0, 6, 0)   // 6 months from now

mv := validation.NewValidator()
mv.Field(checkInDate, "check_in").
    Required().
    DateAfter(minBooking).
    DateBefore(maxBooking)
mv.Field(checkOutDate, "check_out").
    Required().
    DateAfter(checkInDate)  // Must be after check-in
```

**4. Document Expiry Validation / 문서 만료 검증**

```go
now := time.Now()

v := validation.New(expiryDate, "passport_expiry")
v.Required().DateAfter(now)  // Must not be expired
err := v.Validate()
```

#### Performance Characteristics / 성능 특성

| Validator | Time Complexity | Avg Time | Allocations |
|-----------|----------------|----------|-------------|
| DateFormat | O(n) | ~76 ns/op | 0 allocs |
| TimeFormat | O(n) | ~69 ns/op | 0 allocs |
| DateBefore | O(1) | ~32 ns/op | 1 alloc |
| DateAfter | O(1) | ~32 ns/op | 1 alloc |

**Notes** / **참고사항**:
- DateFormat and TimeFormat parse strings, so they're slightly slower / DateFormat과 TimeFormat은 문자열을 파싱하므로 약간 느립니다
- DateBefore and DateAfter are very fast for time.Time objects / DateBefore와 DateAfter는 time.Time 객체에 대해 매우 빠릅니다
- All validators have minimal memory allocations / 모든 검증기는 최소한의 메모리 할당을 합니다

#### Tips and Best Practices / 팁 및 모범 사례

1. **Use Standard Formats** / **표준 형식 사용**
   - Prefer ISO 8601 (`2006-01-02`) for portability
   - ISO 8601 형식은 이식성을 위해 선호됩니다

2. **Validate Format Before Range** / **범위 전에 형식 검증**
   ```go
   // Good: Format validation first
   v.DateFormat("2006-01-02").DateAfter(minDate).DateBefore(maxDate)
   ```

3. **Use UTC for Server-Side Validation** / **서버 측 검증에는 UTC 사용**
   ```go
   now := time.Now().UTC()
   v.DateAfter(now)
   ```

4. **Combine with Custom Validators** / **사용자 정의 검증기와 결합**
   ```go
   v := validation.New(date, "meeting_date")
   v.DateFormat("2006-01-02").
       Custom(func(val interface{}) bool {
           // Check if date is a weekday
           dateStr := val.(string)
           t, _ := time.Parse("2006-01-02", dateStr)
           return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday
       }, "Meeting date must be a weekday")
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

### Range Validators / 범위 검증기

Range validators validate that values are within specified inclusive ranges.

범위 검증기는 값이 지정된 포함 범위 내에 있는지 검증합니다.

#### Available Validators / 사용 가능한 검증기

| Validator | Description | 설명 |
|-----------|-------------|------|
| `IntRange(min, max)` | Validates integer is within range | 정수가 범위 내에 있는지 검증 |
| `FloatRange(min, max)` | Validates float is within range | 실수가 범위 내에 있는지 검증 |
| `DateRange(start, end)` | Validates date is within range | 날짜가 범위 내에 있는지 검증 |

#### IntRange(min, max) - Integer Range Validation / 정수 범위 검증

```go
v := validation.New(25, "age")
v.IntRange(18, 65)
// Valid: 18-65 (inclusive)
// Supports all int types (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64)
```

#### FloatRange(min, max) - Float Range Validation / 실수 범위 검증

```go
v := validation.New(98.6, "temperature")
v.FloatRange(95.0, 105.0)
// Valid: 95.0-105.0 (inclusive)
// Supports float32, float64, and all int types
```

#### DateRange(start, end) - Date Range Validation / 날짜 범위 검증

```go
start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
v := validation.New(eventDate, "event_date")
v.DateRange(start, end)
// Accepts time.Time, RFC3339, or ISO 8601 strings
```

#### Performance / 성능

| Validator | Avg Time | Allocations |
|-----------|----------|-------------|
| IntRange | ~7 ns/op | 0 allocs |
| FloatRange | ~7 ns/op | 0 allocs |
| DateRange | ~35 ns/op | 1 alloc |

---


### File Validators / 파일 검증기

File validators validate file paths, existence, permissions, sizes, and extensions. Perfect for file upload validation, configuration file validation, and file system operations.

파일 검증기는 파일 경로, 존재 여부, 권한, 크기 및 확장자를 검증합니다. 파일 업로드 검증, 구성 파일 검증 및 파일 시스템 작업에 완벽합니다.

#### Available Validators / 사용 가능한 검증기

| Validator | Description | 설명 |
|-----------|-------------|------|
| `FilePath()` | Validates file path format | 파일 경로 형식 검증 |
| `FileExists()` | Validates file/directory exists | 파일/디렉토리 존재 검증 |
| `FileReadable()` | Validates file is readable | 파일 읽기 가능 검증 |
| `FileWritable()` | Validates file is writable | 파일 쓰기 가능 검증 |
| `FileSize(min, max)` | Validates file size in bytes | 바이트 단위 파일 크기 검증 |
| `FileExtension(exts...)` | Validates file extension | 파일 확장자 검증 |

#### FilePath() - File Path Format Validation / 파일 경로 형식 검증

```go
v := validation.New("./config/app.json", "config_file")
v.FilePath()
// Valid: any valid path format (absolute or relative)
// 유효: 모든 유효한 경로 형식 (절대 또는 상대)
```

#### FileExists() - File Existence Validation / 파일 존재 검증

```go
v := validation.New("/etc/hosts", "hosts_file")
v.FileExists()
// Valid: file or directory must exist on filesystem
// 유효: 파일 또는 디렉토리가 파일 시스템에 존재해야 함
```

#### FileReadable() - File Readability Validation / 파일 읽기 가능 검증

```go
v := validation.New("/var/log/app.log", "log_file")
v.FileReadable()
// Valid: file must be readable (opens file to test)
// 유효: 파일이 읽기 가능해야 함 (파일을 열어 테스트)
```

#### FileWritable() - File Writability Validation / 파일 쓰기 가능 검증

```go
v := validation.New("/tmp/output.txt", "output_file")
v.FileWritable()
// Valid: existing file is writable or parent directory is writable for new files
// 유효: 기존 파일은 쓰기 가능하거나 새 파일의 경우 부모 디렉토리가 쓰기 가능
```

#### FileSize(min, max) - File Size Validation / 파일 크기 검증

```go
v := validation.New("/path/to/upload.jpg", "upload_file")
v.FileSize(1024, 10485760) // 1KB - 10MB
// Valid: file size must be between min and max bytes (inclusive)
// 유효: 파일 크기가 최소와 최대 바이트 사이여야 함 (포함)

// Common sizes / 일반적인 크기
// 1 KB = 1024 bytes
// 1 MB = 1048576 bytes (1024 * 1024)
// 10 MB = 10485760 bytes
```

#### FileExtension(extensions...) - File Extension Validation / 파일 확장자 검증

```go
v := validation.New("document.pdf", "file_name")
v.FileExtension(".pdf", ".doc", ".docx")
// Valid: file must have one of the allowed extensions
// 유효: 파일이 허용된 확장자 중 하나를 가져야 함

// Extensions can be specified with or without dot
// 확장자는 점 포함 또는 제외로 지정 가능
v.FileExtension("pdf", "doc", "docx") // Also valid / 또한 유효
```

#### Comprehensive Example / 종합 예제

```go
// File upload validation
mv := validation.NewValidator()
mv.Field(uploadPath, "upload_file").
	FileExists().
	FileReadable().
	FileSize(1024, 10485760).        // 1KB - 10MB
	FileExtension(".jpg", ".png", ".gif")

err := mv.Validate()
if err != nil {
	// Handle validation errors
	// 검증 에러 처리
	fmt.Println(err.Error())
}
```

#### Performance / 성능

| Validator | Avg Time | Allocations | Note |
|-----------|----------|-------------|------|
| FilePath | ~30 ns/op | 0 allocs | Path format check only / 경로 형식만 확인 |
| FileExists | ~1,879 ns/op | 3 allocs | OS stat call / OS stat 호출 |
| FileReadable | ~10,046 ns/op | 4 allocs | Opens file / 파일 열기 |
| FileSize | ~1,915 ns/op | 3 allocs | OS stat call / OS stat 호출 |
| FileExtension | ~10 ns/op | 0 allocs | String comparison / 문자열 비교 |

**Note**: File I/O operations are naturally slower than in-memory validations. FileReadable is the slowest because it actually opens the file to test read permissions.

**참고**: 파일 I/O 작업은 메모리 내 검증보다 자연스럽게 느립니다. FileReadable은 읽기 권한을 테스트하기 위해 실제로 파일을 열기 때문에 가장 느립니다.

#### Use Cases / 사용 사례

**File Upload Validation** / **파일 업로드 검증**
```go
mv.Field(uploadFile, "upload").
	FileSize(0, 5242880).            // Max 5MB
	FileExtension(".jpg", ".png")
```

**Configuration File Validation** / **구성 파일 검증**
```go
mv.Field(configPath, "config").
	FileExists().
	FileReadable().
	FileExtension(".json", ".yaml")
```

**Log File Validation** / **로그 파일 검증**
```go
mv.Field(logPath, "log_file").
	FileWritable()                   // Must be writable
```

---

## Credit Card Validators / 신용카드 검증기

Credit card validators provide validation for credit card numbers, specific card types, and Luhn algorithm checking. Perfect for payment processing, e-commerce platforms, and financial applications.

신용카드 검증기는 신용카드 번호, 특정 카드 타입 및 Luhn 알고리즘 확인을 위한 검증을 제공합니다. 결제 처리, 전자상거래 플랫폼 및 금융 애플리케이션에 완벽합니다.

### Available Validators / 사용 가능한 검증기

| Validator | Description | 설명 |
|-----------|-------------|------|
| `CreditCard()` | Validates credit card number using Luhn algorithm | Luhn 알고리즘을 사용한 신용카드 번호 검증 |
| `CreditCardType(cardType)` | Validates specific card type (Visa, Mastercard, etc.) | 특정 카드 타입 검증 (Visa, Mastercard 등) |
| `Luhn()` | Validates using Luhn algorithm (mod 10 checksum) | Luhn 알고리즘 검증 (mod 10 체크섬) |

### CreditCard() - Credit Card Number Validation / 신용카드 번호 검증

Validates a credit card number using the Luhn algorithm. Accepts numbers with spaces or hyphens, which are automatically removed. The card must be 13-19 digits long and pass the Luhn checksum.

Luhn 알고리즘을 사용하여 신용카드 번호를 검증합니다. 공백이나 하이픈이 있는 번호를 허용하며, 자동으로 제거됩니다. 카드는 13-19자리여야 하며 Luhn 체크섬을 통과해야 합니다.

```go
v := validation.New("4532015112830366", "card_number")
v.CreditCard()
// Valid: passes Luhn algorithm, 16 digits
// 유효: Luhn 알고리즘 통과, 16자리

// With spaces (automatically cleaned)
v := validation.New("4532 0151 1283 0366", "card_number")
v.CreditCard()
// Valid: spaces are removed before validation
// 유효: 검증 전 공백 제거됨

// With hyphens (automatically cleaned)
v := validation.New("4532-0151-1283-0366", "card_number")
v.CreditCard()
// Valid: hyphens are removed before validation
// 유효: 검증 전 하이픈 제거됨
```

**Validation Rules / 검증 규칙:**
- Must be a string / 문자열이어야 함
- After cleaning, must contain only digits / 정리 후 숫자만 포함해야 함
- Length must be 13-19 digits / 길이는 13-19자리여야 함
- Must pass Luhn algorithm check / Luhn 알고리즘 검사를 통과해야 함

### CreditCardType(cardType) - Card Type Validation / 카드 타입 검증

Validates a credit card number against a specific card type pattern. Supports major card networks worldwide.

특정 카드 타입 패턴에 대해 신용카드 번호를 검증합니다. 전 세계 주요 카드 네트워크를 지원합니다.

```go
// Visa validation
v := validation.New("4532015112830366", "card_number")
v.CreditCardType("visa")
// Valid: starts with 4, 13 or 16 digits, passes Luhn
// 유효: 4로 시작, 13 또는 16자리, Luhn 통과

// Mastercard validation
v := validation.New("5425233430109903", "card_number")
v.CreditCardType("mastercard")
// Valid: starts with 51-55, 16 digits, passes Luhn
// 유효: 51-55로 시작, 16자리, Luhn 통과

// American Express validation
v := validation.New("374245455400126", "card_number")
v.CreditCardType("amex")
// Valid: starts with 34 or 37, 15 digits, passes Luhn
// 유효: 34 또는 37로 시작, 15자리, Luhn 통과
```

**Supported Card Types / 지원되는 카드 타입:**

| Card Type | Pattern | Length | Example |
|-----------|---------|--------|---------|
| `visa` | Starts with 4 / 4로 시작 | 13 or 16 digits | 4532015112830366 |
| `mastercard` | Starts with 51-55 / 51-55로 시작 | 16 digits | 5425233430109903 |
| `amex` | Starts with 34 or 37 / 34 또는 37로 시작 | 15 digits | 374245455400126 |
| `discover` | Starts with 6011 or 65 / 6011 또는 65로 시작 | 16 digits | 6011111111111117 |
| `jcb` | Starts with 2131, 1800, or 35 / 2131, 1800, 또는 35로 시작 | 16 digits | 3530111333300000 |
| `dinersclub` | Starts with 300-305, 36, or 38 / 300-305, 36, 또는 38로 시작 | 14 digits | 30569309025904 |
| `unionpay` | Starts with 62 / 62로 시작 | 16-19 digits | 6200000000000005 |

**Note**: Card type names are case-insensitive. You can use "visa", "Visa", or "VISA".

**참고**: 카드 타입 이름은 대소문자를 구분하지 않습니다. "visa", "Visa", "VISA"를 사용할 수 있습니다.

### Luhn() - Luhn Algorithm Validation / Luhn 알고리즘 검증

Validates any number using the Luhn algorithm (mod 10 checksum). Useful for validating identification numbers, account numbers, or any number that uses Luhn validation.

Luhn 알고리즘(mod 10 체크섬)을 사용하여 숫자를 검증합니다. 식별 번호, 계좌 번호 또는 Luhn 검증을 사용하는 모든 번호를 검증하는 데 유용합니다.

```go
v := validation.New("79927398713", "identifier")
v.Luhn()
// Valid: passes Luhn algorithm
// 유효: Luhn 알고리즘 통과

// Credit card number
v := validation.New("4532015112830366", "number")
v.Luhn()
// Valid: any valid Luhn number
// 유효: 유효한 Luhn 번호
```

**How Luhn Algorithm Works / Luhn 알고리즘 작동 방식:**

1. Starting from the rightmost digit, double every second digit / 오른쪽 끝 자리부터 두 번째 자리마다 두 배로 만듦
2. If doubling results in a number > 9, subtract 9 / 두 배가 9보다 크면 9를 뺌
3. Sum all digits / 모든 자리를 더함
4. If sum % 10 == 0, the number is valid / 합계 % 10 == 0이면 번호가 유효함

**Example / 예시:**
```
Number: 79927398713
Step 1: 7 9 9 2 7 3 9 8 7 1 3
Step 2: 7 18 9 4 7 6 9 16 7 2 3  (double every 2nd from right)
Step 3: 7 9 9 4 7 6 9 7 7 2 3    (subtract 9 if > 9)
Step 4: 7+9+9+4+7+6+9+7+7+2+3 = 70
Step 5: 70 % 10 = 0 ✓ Valid!
```

### Comprehensive Example / 종합 예제

```go
// Payment validation with multiple checks
mv := validation.NewValidator()

// Validate credit card number
mv.Field(cardNumber, "card_number").
	Required().
	CreditCard().
	CreditCardType("visa")

// Validate CVV
mv.Field(cvv, "cvv").
	Required().
	Length(3, 4).
	Numeric()

// Validate expiration date
mv.Field(expiryDate, "expiry_date").
	Required().
	DateFormat("01/06").  // MM/YY format
	DateAfter(time.Now())

err := mv.Validate()
if err != nil {
	// Handle validation errors
	// 검증 에러 처리
	fmt.Println("Payment validation failed:", err)
	return
}

fmt.Println("Payment information validated successfully")
```

### Performance / 성능

| Validator | Avg Time | Allocations | Note |
|-----------|----------|-------------|------|
| CreditCard | ~550 ns/op | 2 allocs | Includes Luhn check / Luhn 체크 포함 |
| CreditCardType | ~950 ns/op | 2 allocs | Pattern matching + Luhn / 패턴 매칭 + Luhn |
| Luhn | ~450 ns/op | 2 allocs | Pure Luhn algorithm / 순수 Luhn 알고리즘 |

**Note**: Credit card validation is very fast (<1 microsecond) and suitable for real-time validation in payment forms.

**참고**: 신용카드 검증은 매우 빠르며(<1 마이크로초) 결제 양식의 실시간 검증에 적합합니다.

### Use Cases / 사용 사례

**E-commerce Payment Validation** / **전자상거래 결제 검증**
```go
mv.Field(cardNumber, "card_number").
	CreditCard().
	CreditCardType("visa")
```

**Multi-Card Type Support** / **다중 카드 타입 지원**
```go
// Accept Visa, Mastercard, or Amex
cardType := detectCardType(cardNumber)
mv.Field(cardNumber, "card_number").
	CreditCardType(cardType)
```

**Generic Luhn Validation** / **일반 Luhn 검증**
```go
// For any Luhn-validated number (IMEI, etc.)
mv.Field(imeiNumber, "imei").
	Luhn()
```

### Security Considerations / 보안 고려사항

**Important**: These validators only check the format and checksum of credit card numbers. They do NOT verify if the card is active, has sufficient balance, or belongs to a specific person.

**중요**: 이러한 검증기는 신용카드 번호의 형식과 체크섬만 확인합니다. 카드가 활성화되어 있는지, 잔액이 충분한지, 특정 사람에게 속하는지는 확인하지 않습니다.

**For production payment processing / 프로덕션 결제 처리의 경우:**
- Use a payment gateway like Stripe, PayPal, or Square / Stripe, PayPal, Square 같은 결제 게이트웨이 사용
- Never store full credit card numbers / 전체 신용카드 번호를 저장하지 말 것
- Use PCI DSS compliant storage if required / 필요한 경우 PCI DSS 준수 스토리지 사용
- Log only masked card numbers (e.g., "****1234") / 마스킹된 카드 번호만 로그에 기록 (예: "****1234")
- Transmit card data only over HTTPS / 카드 데이터는 HTTPS로만 전송

**Test Card Numbers / 테스트 카드 번호:**

The following are standard test card numbers that pass Luhn validation (use these for testing):

다음은 Luhn 검증을 통과하는 표준 테스트 카드 번호입니다(테스트에 사용):

- **Visa**: 4532015112830366, 4532015112830
- **Mastercard**: 5425233430109903, 5105105105105100
- **Amex**: 374245455400126, 340000000000009
- **Discover**: 6011111111111117, 6500000000000002
- **JCB**: 3530111333300000
- **Diners Club**: 30569309025904

---

## Business/ID Validators / 비즈니스/ID 검증기

Business/ID validators validate international standard identifiers used in commerce, publishing, and inventory systems. Perfect for e-commerce platforms, library systems, inventory management, and publishing applications.

비즈니스/ID 검증기는 상거래, 출판 및 재고 시스템에서 사용되는 국제 표준 식별자를 검증합니다. 전자상거래 플랫폼, 도서관 시스템, 재고 관리 및 출판 애플리케이션에 완벽합니다.

### Available Validators / 사용 가능한 검증기

| Validator | Description | 설명 |
|-----------|-------------|------|
| `ISBN()` | Validates International Standard Book Number (ISBN-10 or ISBN-13) | 국제 표준 도서 번호 검증 (ISBN-10 또는 ISBN-13) |
| `ISSN()` | Validates International Standard Serial Number (ISSN-8) | 국제 표준 연속 간행물 번호 검증 (ISSN-8) |
| `EAN()` | Validates European Article Number (EAN-8 or EAN-13) | 유럽 상품 코드 검증 (EAN-8 또는 EAN-13) |

### ISBN() - Book Number Validation / 도서 번호 검증

Validates International Standard Book Number with checksum algorithm. Supports both ISBN-10 and ISBN-13 formats with or without hyphens.

체크섬 알고리즘을 사용하여 국제 표준 도서 번호를 검증합니다. 하이픈 포함/미포함 ISBN-10 및 ISBN-13 형식을 모두 지원합니다.

```go
// ISBN-13 validation
v := validation.New("978-0-596-52068-7", "book_isbn")
v.ISBN()
// Valid: proper ISBN-13 format with correct checksum
// 유효: 올바른 체크섬이 있는 적절한 ISBN-13 형식

// ISBN-10 validation
v := validation.New("0-596-52068-9", "book_isbn")
v.ISBN()
// Valid: proper ISBN-10 format with correct checksum
// 유효: 올바른 체크섬이 있는 적절한 ISBN-10 형식

// Without hyphens
v := validation.New("9780596520687", "book_isbn")
v.ISBN()
// Valid: hyphens are optional
// 유효: 하이픈은 선택 사항
```

**ISBN-10 Format / ISBN-10 형식:**
- 10 characters: 9 digits + checksum (0-9 or X)
- 10자: 9자리 숫자 + 체크섬 (0-9 또는 X)
- Checksum algorithm: weighted sum mod 11
- 체크섬 알고리즘: 가중 합계 mod 11

**ISBN-13 Format / ISBN-13 형식:**
- 13 digits with alternating weights (1 and 3)
- 교대 가중치(1과 3)가 있는 13자리
- Checksum: (10 - (sum mod 10)) mod 10
- 체크섬: (10 - (합계 mod 10)) mod 10

### ISSN() - Serial Number Validation / 연속 간행물 번호 검증

Validates International Standard Serial Number for periodicals, journals, and magazines.

정기간행물, 저널 및 잡지에 대한 국제 표준 연속 간행물 번호를 검증합니다.

```go
v := validation.New("2049-3630", "journal_issn")
v.ISSN()
// Valid: proper ISSN format (XXXX-XXXX)
// 유효: 적절한 ISSN 형식 (XXXX-XXXX)

// Without hyphen
v := validation.New("20493630", "journal_issn")
v.ISSN()
// Valid: hyphen is optional
// 유효: 하이픈은 선택 사항

// ISSN ending with X (checksum digit)
v := validation.New("0317-847X", "journal_issn")
v.ISSN()
// Valid: X is valid checksum digit
// 유효: X는 유효한 체크섬 자리
```

**ISSN Format / ISSN 형식:**
- 8 characters: 7 digits + checksum (0-9 or X)
- 8자: 7자리 숫자 + 체크섬 (0-9 또는 X)
- Format: XXXX-XXXX (hyphen after 4th digit is optional)
- 형식: XXXX-XXXX (4번째 자리 뒤의 하이픈은 선택 사항)
- Checksum algorithm: weighted sum mod 11
- 체크섬 알고리즘: 가중 합계 mod 11

### EAN() - Product Barcode Validation / 제품 바코드 검증

Validates European Article Number used in retail product barcodes. Supports both EAN-8 and EAN-13 formats.

소매 제품 바코드에 사용되는 유럽 상품 코드를 검증합니다. EAN-8 및 EAN-13 형식을 모두 지원합니다.

```go
// EAN-13 (most common)
v := validation.New("4006381333931", "product_ean")
v.EAN()
// Valid: 13-digit product barcode
// 유효: 13자리 제품 바코드

// EAN-8 (compact format)
v := validation.New("96385074", "product_ean")
v.EAN()
// Valid: 8-digit compact barcode
// 유효: 8자리 컴팩트 바코드

// With spaces or hyphens (auto-cleaned)
v := validation.New("400-6381-333-931", "product_ean")
v.EAN()
// Valid: spaces and hyphens are removed
// 유효: 공백과 하이픈 제거됨
```

**EAN-8 Format / EAN-8 형식:**
- 8 digits with alternating weights (3 and 1)
- 교대 가중치(3과 1)가 있는 8자리
- Used for small products / 소형 제품에 사용

**EAN-13 Format / EAN-13 형식:**
- 13 digits with alternating weights (1 and 3)
- 교대 가중치(1과 3)가 있는 13자리
- Standard product barcode / 표준 제품 바코드
- Compatible with UPC / UPC와 호환

### Comprehensive Example / 종합 예제

```go
// E-commerce product validation
mv := validation.NewValidator()

// Validate book ISBN
mv.Field(bookISBN, "book_isbn").
	Required().
	ISBN()

// Validate magazine ISSN
mv.Field(magazineISSN, "magazine_issn").
	Required().
	ISSN()

// Validate product barcode
mv.Field(productEAN, "product_ean").
	Required().
	EAN()

err := mv.Validate()
if err != nil {
	// Handle validation errors
	// 검증 에러 처리
	fmt.Println("Invalid identifiers:", err)
	return
}

fmt.Println("All identifiers validated successfully")
```

### Performance / 성능

| Validator | Avg Time | Allocations | Note |
|-----------|----------|-------------|------|
| ISBN | ~650 ns/op | 2 allocs | Includes checksum validation / 체크섬 검증 포함 |
| ISSN | ~550 ns/op | 2 allocs | Mod 11 checksum / Mod 11 체크섬 |
| EAN | ~600 ns/op | 2 allocs | Alternating weight checksum / 교대 가중치 체크섬 |

**Note**: All validators are very fast (<1 microsecond) and suitable for real-time validation in e-commerce and inventory systems.

**참고**: 모든 검증기는 매우 빠르며(<1 마이크로초) 전자상거래 및 재고 시스템의 실시간 검증에 적합합니다.

### Use Cases / 사용 사례

**Online Bookstore** / **온라인 서점**
```go
mv.Field(bookISBN, "isbn").
	ISBN()
```

**Library Management System** / **도서관 관리 시스템**
```go
// Book
mv.Field(isbn, "book_identifier").ISBN()

// Journal/Magazine
mv.Field(issn, "journal_identifier").ISSN()
```

**E-commerce Product Catalog** / **전자상거래 제품 카탈로그**
```go
mv.Field(productBarcode, "barcode").
	EAN()
```

**Inventory Management** / **재고 관리**
```go
// Validate all product identifiers
products := []struct {
	ISBN string
	EAN  string
}{
	{"978-0-596-52068-7", "4006381333931"},
	// ... more products
}

for _, p := range products {
	mv.Field(p.ISBN, "isbn").ISBN()
	mv.Field(p.EAN, "ean").EAN()
}
```

### Validation Rules / 검증 규칙

**ISBN:**
- Must be 10 or 13 digits (after removing hyphens/spaces)
- 10 또는 13자리여야 함 (하이픈/공백 제거 후)
- ISBN-10: Last digit can be 0-9 or X
- ISBN-10: 마지막 자리는 0-9 또는 X 가능
- Must pass checksum validation
- 체크섬 검증을 통과해야 함

**ISSN:**
- Must be 8 characters (after removing hyphens/spaces)
- 8자여야 함 (하이픈/공백 제거 후)
- Last digit can be 0-9 or X
- 마지막 자리는 0-9 또는 X 가능
- Format: XXXX-XXXX (hyphen optional)
- 형식: XXXX-XXXX (하이픈 선택 사항)

**EAN:**
- Must be 8 or 13 digits (after removing hyphens/spaces)
- 8 또는 13자리여야 함 (하이픈/공백 제거 후)
- All digits only (no letters)
- 숫자만 가능 (문자 불가)
- Must pass checksum validation
- 체크섬 검증을 통과해야 함

### Common Validation Scenarios / 일반적인 검증 시나리오

**Book Publishing** / **도서 출판**
```go
// Validate both ISBN-10 and ISBN-13
mv.Field(isbn10, "isbn_10").ISBN()  // 0-596-52068-9
mv.Field(isbn13, "isbn_13").ISBN()  // 978-0-596-52068-7
```

**Magazine Subscription** / **잡지 구독**
```go
mv.Field(issn, "magazine_issn").
	Required().
	ISSN()
```

**Retail POS System** / **소매 POS 시스템**
```go
// Scan product barcode
mv.Field(scannedBarcode, "barcode").
	EAN()
```

**Import/Export** / **수입/수출**
```go
// Validate international product codes
mv.Field(ean13, "product_code").
	EAN()  // EAN-13 for international products
```

---

## Geographic Validators / 지리 좌표 검증기

Geographic validators ensure that location data (latitude, longitude, coordinates) is valid according to standard geographic coordinate systems. These validators are essential for mapping applications, location services, and geographic information systems (GIS).

지리 좌표 검증기는 위치 데이터(위도, 경도, 좌표)가 표준 지리 좌표 시스템에 따라 유효한지 확인합니다. 이러한 검증기는 지도 애플리케이션, 위치 서비스 및 지리 정보 시스템(GIS)에 필수적입니다.

### Available Validators / 사용 가능한 검증기

| Validator | Description (EN) | Description (KR) | Supported Types |
|-----------|------------------|------------------|-----------------|
| `Latitude()` | Validates latitude coordinates (-90 to 90 degrees) | 위도 좌표를 검증합니다 (-90 ~ 90도) | `float64`, `float32`, `int`, `int64`, `string` |
| `Longitude()` | Validates longitude coordinates (-180 to 180 degrees) | 경도 좌표를 검증합니다 (-180 ~ 180도) | `float64`, `float32`, `int`, `int64`, `string` |
| `Coordinate()` | Validates coordinate pairs in "lat,lon" format | "위도,경도" 형식의 좌표 쌍을 검증합니다 | `string` |

### 1. Latitude Validator / 위도 검증기

The `Latitude()` validator ensures that a value represents a valid latitude coordinate. Latitude values must be between -90° (South Pole) and +90° (North Pole).

`Latitude()` 검증기는 값이 유효한 위도 좌표를 나타내는지 확인합니다. 위도 값은 -90°(남극)와 +90°(북극) 사이여야 합니다.

**Validation Rules / 검증 규칙:**
- **Range**: -90.0 ≤ latitude ≤ 90.0 / **범위**: -90.0 ≤ 위도 ≤ 90.0
- **Supported Types**: `float64`, `float32`, `int`, `int64`, `string` / **지원 타입**: `float64`, `float32`, `int`, `int64`, `string`
- **String Format**: Must be a parseable number / **문자열 형식**: 파싱 가능한 숫자여야 함

**Examples / 예시:**

```go
// Basic latitude validation / 기본 위도 검증
latitude := 37.5665  // Seoul latitude
v := validation.New(latitude, "latitude")
v.Latitude()

if err := v.Validate(); err != nil {
    fmt.Println(err)  // No error - valid latitude
}

// Validate latitude from different types / 다양한 타입의 위도 검증
v1 := validation.New(37.5665, "lat").Latitude()          // float64
v2 := validation.New(float32(37.5), "lat").Latitude()    // float32
v3 := validation.New(45, "lat").Latitude()                // int
v4 := validation.New("37.5665", "lat").Latitude()        // string

// Invalid latitudes / 유효하지 않은 위도
v5 := validation.New(90.1, "lat").Latitude()             // Too high / 너무 높음
v6 := validation.New(-90.1, "lat").Latitude()            // Too low / 너무 낮음
v7 := validation.New("abc", "lat").Latitude()            // Non-numeric / 숫자가 아님
```

**Boundary Cases / 경계 케이스:**
```go
// Exactly at boundaries (valid) / 경계값 (유효)
v1 := validation.New(90.0, "lat").Latitude()    // North Pole / 북극 ✅
v2 := validation.New(-90.0, "lat").Latitude()   // South Pole / 남극 ✅

// Just outside boundaries (invalid) / 경계 밖 (유효하지 않음)
v3 := validation.New(90.0001, "lat").Latitude()  // ❌
v4 := validation.New(-90.0001, "lat").Latitude() // ❌
```

### 2. Longitude Validator / 경도 검증기

The `Longitude()` validator ensures that a value represents a valid longitude coordinate. Longitude values must be between -180° (International Date Line, west) and +180° (International Date Line, east).

`Longitude()` 검증기는 값이 유효한 경도 좌표를 나타내는지 확인합니다. 경도 값은 -180°(국제 날짜 변경선, 서쪽)와 +180°(국제 날짜 변경선, 동쪽) 사이여야 합니다.

**Validation Rules / 검증 규칙:**
- **Range**: -180.0 ≤ longitude ≤ 180.0 / **범위**: -180.0 ≤ 경도 ≤ 180.0
- **Supported Types**: `float64`, `float32`, `int`, `int64`, `string` / **지원 타입**: `float64`, `float32`, `int`, `int64`, `string`
- **String Format**: Must be a parseable number / **문자열 형식**: 파싱 가능한 숫자여야 함

**Examples / 예시:**

```go
// Basic longitude validation / 기본 경도 검증
longitude := 126.9780  // Seoul longitude
v := validation.New(longitude, "longitude")
v.Longitude()

if err := v.Validate(); err != nil {
    fmt.Println(err)  // No error - valid longitude
}

// Validate longitude from different types / 다양한 타입의 경도 검증
v1 := validation.New(126.9780, "lon").Longitude()        // float64
v2 := validation.New(float32(126.9), "lon").Longitude()  // float32
v3 := validation.New(90, "lon").Longitude()               // int
v4 := validation.New("126.9780", "lon").Longitude()      // string

// Invalid longitudes / 유효하지 않은 경도
v5 := validation.New(180.1, "lon").Longitude()           // Too high / 너무 높음
v6 := validation.New(-180.1, "lon").Longitude()          // Too low / 너무 낮음
v7 := validation.New("xyz", "lon").Longitude()           // Non-numeric / 숫자가 아님
```

**Boundary Cases / 경계 케이스:**
```go
// Exactly at boundaries (valid) / 경계값 (유효)
v1 := validation.New(180.0, "lon").Longitude()   // International Date Line / 국제 날짜 변경선 ✅
v2 := validation.New(-180.0, "lon").Longitude()  // International Date Line / 국제 날짜 변경선 ✅

// Just outside boundaries (invalid) / 경계 밖 (유효하지 않음)
v3 := validation.New(180.0001, "lon").Longitude()  // ❌
v4 := validation.New(-180.0001, "lon").Longitude() // ❌
```

### 3. Coordinate Validator / 좌표 검증기

The `Coordinate()` validator validates coordinate pairs in "latitude,longitude" format. It parses the string, validates both components, and ensures they are within valid ranges.

`Coordinate()` 검증기는 "위도,경도" 형식의 좌표 쌍을 검증합니다. 문자열을 파싱하여 두 구성 요소를 모두 검증하고 유효한 범위 내에 있는지 확인합니다.

**Validation Rules / 검증 규칙:**
- **Format**: "latitude,longitude" (comma-separated) / **형식**: "위도,경도" (쉼표로 구분)
- **Optional Spaces**: Spaces around comma are allowed / **선택적 공백**: 쉼표 주변 공백 허용
- **Latitude Range**: -90.0 ≤ latitude ≤ 90.0 / **위도 범위**: -90.0 ≤ 위도 ≤ 90.0
- **Longitude Range**: -180.0 ≤ longitude ≤ 180.0 / **경도 범위**: -180.0 ≤ 경도 ≤ 180.0
- **Type**: String only / **타입**: 문자열만

**Examples / 예시:**

```go
// Basic coordinate validation / 기본 좌표 검증
coordinate := "37.5665,126.9780"  // Seoul, South Korea
v := validation.New(coordinate, "location")
v.Coordinate()

if err := v.Validate(); err != nil {
    fmt.Println(err)  // No error - valid coordinate
}

// Various valid formats / 다양한 유효 형식
v1 := validation.New("37.5665,126.9780", "loc").Coordinate()   // No spaces
v2 := validation.New("37.5665, 126.9780", "loc").Coordinate()  // Space after comma
v3 := validation.New("  37.5665  ,  126.9780  ", "loc").Coordinate()  // Extra spaces
v4 := validation.New("0,0", "loc").Coordinate()                 // Null Island
v5 := validation.New("-90,-180", "loc").Coordinate()            // Min values
v6 := validation.New("90,180", "loc").Coordinate()              // Max values

// Famous locations / 유명한 위치
vSeoul := validation.New("37.5665,126.9780", "Seoul").Coordinate()
vNewYork := validation.New("40.7128,-74.0060", "New York").Coordinate()
vLondon := validation.New("51.5074,-0.1278", "London").Coordinate()
vTokyo := validation.New("35.6762,139.6503", "Tokyo").Coordinate()

// Invalid coordinates / 유효하지 않은 좌표
v7 := validation.New("91,0", "loc").Coordinate()              // Latitude out of range
v8 := validation.New("0,181", "loc").Coordinate()             // Longitude out of range
v9 := validation.New("37.5665", "loc").Coordinate()           // Missing longitude
v10 := validation.New("37.5665 126.9780", "loc").Coordinate() // No comma
v11 := validation.New("abc,xyz", "loc").Coordinate()          // Non-numeric
```

**Error Messages / 에러 메시지:**
```go
v := validation.New("91,0", "location")
v.Coordinate()
// Error: "location latitude must be between -90 and 90 / location 위도는 -90과 90 사이여야 합니다"

v2 := validation.New("0,181", "location")
v2.Coordinate()
// Error: "location longitude must be between -180 and 180 / location 경도는 -180과 180 사이여야 합니다"

v3 := validation.New("abc,xyz", "location")
v3.Coordinate()
// Error: "location latitude must be a valid number / location 위도는 유효한 숫자여야 합니다"
```

### Multi-Field Geographic Validation / 다중 필드 지리 좌표 검증

Validate multiple geographic fields together for location-based data:

위치 기반 데이터를 위한 여러 지리 필드를 함께 검증합니다:

```go
type Location struct {
    Latitude    float64
    Longitude   float64
    Coordinate  string
    Altitude    float64
}

func ValidateLocation(loc Location) error {
    mv := validation.NewValidator()

    // Validate separate latitude/longitude fields
    // 개별 위도/경도 필드 검증
    mv.Field(loc.Latitude, "latitude").
        Required().
        Latitude()

    mv.Field(loc.Longitude, "longitude").
        Required().
        Longitude()

    // Validate coordinate string
    // 좌표 문자열 검증
    mv.Field(loc.Coordinate, "coordinate").
        Required().
        Coordinate()

    // Validate altitude (optional)
    // 고도 검증 (선택적)
    if loc.Altitude != 0 {
        mv.Field(loc.Altitude, "altitude").
            FloatRange(-500.0, 9000.0)  // Sea level to Everest
    }

    return mv.Validate()
}
```

### Chaining with Other Validators / 다른 검증기와 체이닝

Combine geographic validators with other validation rules:

지리 좌표 검증기를 다른 검증 규칙과 결합합니다:

```go
// Validate required coordinate field
// 필수 좌표 필드 검증
v := validation.New(coordinate, "user_location")
v.Required().Coordinate()

// Validate optional latitude with custom error handling
// 사용자 정의 에러 처리로 선택적 위도 검증
v2 := validation.New(latitude, "optional_lat").StopOnError()
if latitude != 0 {  // Only validate if provided
    v2.Latitude()
}

// Multi-field validation with stop-on-error
// 첫 에러에서 멈춤과 함께 다중 필드 검증
mv := validation.NewValidator()
mv.Field(location.Lat, "latitude").StopOnError().Required().Latitude()
mv.Field(location.Lon, "longitude").StopOnError().Required().Longitude()
```

### Real-World Use Cases / 실제 사용 사례

**Location-Based Services** / **위치 기반 서비스**
```go
// Validate user's current location
mv.Field(userLat, "user_latitude").
    Required().
    Latitude()

mv.Field(userLon, "user_longitude").
    Required().
    Longitude()
```

**Mapping and Navigation** / **지도 및 내비게이션**
```go
// Validate destination coordinates from user input
mv.Field(destination, "destination").
    Required().
    Coordinate()

// Validate waypoint coordinates
for i, waypoint := range waypoints {
    mv.Field(waypoint, fmt.Sprintf("waypoint_%d", i)).
        Coordinate()
}
```

**GIS and Geospatial Applications** / **GIS 및 공간 정보 애플리케이션**
```go
// Validate boundary box for map query
mv.Field(minLat, "min_latitude").Required().Latitude()
mv.Field(maxLat, "max_latitude").Required().Latitude()
mv.Field(minLon, "min_longitude").Required().Longitude()
mv.Field(maxLon, "max_longitude").Required().Longitude()

// Also validate logical constraints
if minLat >= maxLat {
    return errors.New("min_latitude must be less than max_latitude")
}
if minLon >= maxLon {
    return errors.New("min_longitude must be less than max_longitude")
}
```

**Delivery and Logistics** / **배송 및 물류**
```go
// Validate pickup and delivery locations
mv.Field(pickupLocation, "pickup_location").
    Required().
    Coordinate()

mv.Field(deliveryLocation, "delivery_location").
    Required().
    Coordinate()
```

**IoT and Telemetry** / **IoT 및 원격 측정**
```go
// Validate GPS coordinates from IoT devices
mv.Field(deviceLat, "device_latitude").
    Latitude()

mv.Field(deviceLon, "device_longitude").
    Longitude()

// Coordinate validation from GPS string
mv.Field(gpsData, "gps_coordinates").
    Coordinate()
```

### Performance / 성능

Geographic validators are highly optimized for common use cases:

지리 좌표 검증기는 일반적인 사용 사례에 맞게 고도로 최적화되어 있습니다:

- **Latitude**: ~300-400 ns/op (sub-microsecond)
- **Longitude**: ~300-400 ns/op (sub-microsecond)
- **Coordinate**: ~600-800 ns/op (string parsing + dual validation)

**Benchmarks:**
```
BenchmarkLatitude-8    3000000    350 ns/op
BenchmarkLongitude-8   3000000    350 ns/op
BenchmarkCoordinate-8  2000000    750 ns/op
```

---

## Security Validators / 보안 검증기

Security validators ensure that cryptographic hashes, tokens, and security-related data formats are valid. These validators are essential for authentication systems, data integrity verification, and secure API communications.

보안 검증기는 암호화 해시, 토큰 및 보안 관련 데이터 형식이 유효한지 확인합니다. 이러한 검증기는 인증 시스템, 데이터 무결성 검증 및 안전한 API 통신에 필수적입니다.

### Available Validators / 사용 가능한 검증기

| Validator | Description (EN) | Description (KR) | Format |
|-----------|------------------|------------------|--------|
| `JWT()` | Validates JWT (JSON Web Token) format | JWT 형식을 검증합니다 | header.payload.signature |
| `BCrypt()` | Validates BCrypt password hash format | BCrypt 비밀번호 해시 형식을 검증합니다 | $2[abxy]$cost$hash |
| `MD5()` | Validates MD5 hash (32 hex characters) | MD5 해시를 검증합니다 (32자리 16진수) | 32 hex chars |
| `SHA1()` | Validates SHA1 hash (40 hex characters) | SHA1 해시를 검증합니다 (40자리 16진수) | 40 hex chars |
| `SHA256()` | Validates SHA256 hash (64 hex characters) | SHA256 해시를 검증합니다 (64자리 16진수) | 64 hex chars |
| `SHA512()` | Validates SHA512 hash (128 hex characters) | SHA512 해시를 검증합니다 (128자리 16진수) | 128 hex chars |

### 1. JWT Validator / JWT 검증기

The `JWT()` validator ensures that a value is a valid JSON Web Token format. It validates the three-part structure (header.payload.signature) and base64url encoding of each part.

`JWT()` 검증기는 값이 유효한 JSON Web Token 형식인지 확인합니다. 세 부분 구조(header.payload.signature)와 각 부분의 base64url 인코딩을 검증합니다.

**Validation Rules / 검증 규칙:**
- **Format**: `header.payload.signature` (three parts separated by dots)
- **Encoding**: Each part must be valid base64url
- **Non-empty**: Header and payload must not be empty

**Examples / 예시:**
```go
// Valid JWT token
token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U"
v := validation.New(token, "auth_token")
v.JWT()

// Authenticate API request
mv := validation.NewValidator()
mv.Field(authHeader, "authorization").Required().JWT()
```

**Use Cases:**
- API authentication token validation
- OAuth 2.0 / OpenID Connect token verification
- Microservice inter-service communication
- Mobile app authentication

### 2. BCrypt Validator / BCrypt 검증기

The `BCrypt()` validator validates BCrypt password hash format. BCrypt is a widely-used password hashing function with built-in salt.

`BCrypt()` 검증기는 BCrypt 비밀번호 해시 형식을 검증합니다. BCrypt는 내장 솔트가 있는 널리 사용되는 비밀번호 해싱 함수입니다.

**Validation Rules:**
- **Prefix**: Must start with `$2a$`, `$2b$`, `$2x$`, or `$2y$`
- **Length**: Exactly 60 characters
- **Format**: `$2[abxy]$[cost]$[salt][hash]`

**Examples:**
```go
// Validate password hash from database
hash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
v := validation.New(hash, "password_hash")
v.BCrypt()

// User registration validation
mv.Field(user.PasswordHash, "password").Required().BCrypt()
```

**Use Cases:**
- Password storage validation
- User authentication systems
- Secure credential verification
- Password migration validation

### 3. Hash Validators (MD5, SHA1, SHA256, SHA512) / 해시 검증기

Hash validators ensure cryptographic hash values are correctly formatted. These are commonly used for file integrity, data verification, and checksum validation.

해시 검증기는 암호화 해시 값이 올바르게 형식화되었는지 확인합니다. 파일 무결성, 데이터 검증 및 체크섬 검증에 일반적으로 사용됩니다.

**MD5 (32 hex characters):**
```go
hash := "5d41402abc4b2a76b9719d911017c592"
v := validation.New(hash, "file_md5")
v.MD5()
```

**SHA1 (40 hex characters):**
```go
hash := "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
v := validation.New(hash, "commit_hash")
v.SHA1()
```

**SHA256 (64 hex characters):**
```go
hash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
v := validation.New(hash, "file_hash")
v.SHA256()
```

**SHA512 (128 hex characters):**
```go
hash := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"
v := validation.New(hash, "secure_hash")
v.SHA512()
```

### Multi-Field Security Validation / 다중 필드 보안 검증

```go
type SecureRequest struct {
    Token          string
    PasswordHash   string
    FileChecksum   string
}

func ValidateSecureRequest(req SecureRequest) error {
    mv := validation.NewValidator()

    mv.Field(req.Token, "token").
        Required().
        JWT()

    mv.Field(req.PasswordHash, "password").
        Required().
        BCrypt()

    mv.Field(req.FileChecksum, "checksum").
        Required().
        SHA256()

    return mv.Validate()
}
```

### Real-World Use Cases / 실제 사용 사례

**API Authentication:**
```go
// Validate JWT bearer token
mv.Field(bearerToken, "authorization").
    Required().
    JWT()
```

**Password Management:**
```go
// Validate stored password hash
mv.Field(user.PasswordHash, "password").
    Required().
    BCrypt()
```

**File Integrity Verification:**
```go
// Validate file checksums
mv.Field(uploadedFileHash, "file_hash").
    Required().
    SHA256()

mv.Field(expectedHash, "expected_hash").
    Required().
    SHA256()
```

**Git Commit Validation:**
```go
// Validate commit hashes
mv.Field(commitSHA, "commit").
    Required().
    SHA1()
```

**Blockchain/Cryptocurrency:**
```go
// Validate transaction hashes
mv.Field(txHash, "transaction").
    Required().
    SHA256()
```

### Performance / 성능

Security validators are highly optimized with regex matching:

보안 검증기는 정규식 매칭으로 고도로 최적화되어 있습니다:

- **JWT**: ~800-1000 ns/op (base64 decoding + validation)
- **BCrypt**: ~200-300 ns/op (regex pattern matching)
- **MD5**: ~150-200 ns/op (32-char hex validation)
- **SHA1**: ~150-200 ns/op (40-char hex validation)
- **SHA256**: ~150-200 ns/op (64-char hex validation)
- **SHA512**: ~150-200 ns/op (128-char hex validation)

**Note**: Hash validators only validate format, not cryptographic correctness. For actual hash verification, use Go's `crypto` package.

**참고**: 해시 검증기는 형식만 검증하며 암호화 정확성은 검증하지 않습니다. 실제 해시 검증을 위해서는 Go의 `crypto` 패키지를 사용하세요.

---
