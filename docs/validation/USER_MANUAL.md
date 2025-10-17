# Validation Package - User Manual / Validation 패키지 - 사용자 매뉴얼

**Version / 버전**: v1.13.020
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
12. [File Validators / 파일 검증기](#file-validators--파일-검증기) 🆕
13. [Advanced Features / 고급 기능](#advanced-features--고급-기능)
13. [Error Handling / 에러 처리](#error-handling--에러-처리)
14. [Real-World Examples / 실제 사용 예제](#real-world-examples--실제-사용-예제)
15. [Best Practices / 모범 사례](#best-practices--모범-사례)
16. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)

---

## Introduction / 소개

The `validation` package provides a **fluent, type-safe validation library** for Go 1.18+. It reduces 20-30 lines of validation boilerplate to just 1-2 lines using method chaining and provides bilingual error messages (English/Korean).

`validation` 패키지는 Go 1.18+ 환경을 위한 **플루언트하고 타입 안전한 검증 라이브러리**를 제공합니다. 메서드 체이닝을 사용하여 20-30줄의 검증 보일러플레이트를 단 1-2줄로 줄이며, 양방향 에러 메시지(영어/한글)를 제공합니다.

### Key Features / 주요 기능

- ✅ **70+ Built-in Validators** / **70개 이상의 내장 검증기**
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
- ✅ **File Validators (FilePath, FileExists, FileReadable, FileWritable, FileSize, FileExtension)** 🆕 / **파일 검증기** 🆕

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
