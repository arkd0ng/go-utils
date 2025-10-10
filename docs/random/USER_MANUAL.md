# Random Package User Manual / Random 패키지 사용자 매뉴얼

Complete user manual for the Random package - cryptographically secure random string generation.

Random 패키지의 완전한 사용자 매뉴얼 - 암호학적으로 안전한 랜덤 문자열 생성.

**Version / 버전**: v1.0.008
**Last Updated / 최종 업데이트**: 2025-10-10

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Method Reference / 메서드 참조](#method-reference--메서드-참조)
5. [Usage Patterns / 사용 패턴](#usage-patterns--사용-패턴)
6. [Common Use Cases / 일반적인 사용 사례](#common-use-cases--일반적인-사용-사례)
7. [Error Handling / 에러 처리](#error-handling--에러-처리)
8. [Best Practices / 모범 사례](#best-practices--모범-사례)
9. [FAQ / 자주 묻는 질문](#faq--자주-묻는-질문)
10. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)

---

## Introduction / 소개

The Random package provides cryptographically secure random string generation with 14 different methods, flexible length parameters, and comprehensive error handling.

Random 패키지는 14가지 다양한 메서드, 유연한 길이 파라미터, 포괄적인 에러 처리를 갖춘 암호학적으로 안전한 랜덤 문자열 생성을 제공합니다.

### Key Features / 주요 기능

- **Cryptographically Secure / 암호학적으로 안전**: Uses `crypto/rand` for secure random generation / `crypto/rand` 사용
- **14 Different Methods / 14가지 메서드**: From basic to specialized character sets / 기본부터 특수화된 문자 집합까지
- **Flexible Length / 유연한 길이**: Fixed length or range (min-max) / 고정 길이 또는 범위 (최소-최대)
- **Error Handling / 에러 처리**: All methods return `(string, error)` / 모든 메서드가 `(string, error)` 반환
- **Thread-Safe / 스레드 안전**: Safe for concurrent use / 동시 사용 안전

### Security / 보안

This package is suitable for security-sensitive applications:

이 패키지는 보안이 중요한 애플리케이션에 적합합니다:

- ✅ Password generation / 비밀번호 생성
- ✅ API key generation / API 키 생성
- ✅ Token generation / 토큰 생성
- ✅ Session ID generation / 세션 ID 생성
- ✅ Cryptographic salt generation / 암호화 salt 생성

---

## Installation / 설치

### Prerequisites / 사전 요구사항

- Go 1.16 or higher / Go 1.16 이상
- No external dependencies / 외부 의존성 없음

### Install Command / 설치 명령

```bash
go get github.com/arkd0ng/go-utils/random
```

### Import / Import

```go
import "github.com/arkd0ng/go-utils/random"
```

---

## Quick Start / 빠른 시작

### Simple Example / 간단한 예제

```go
package main

import (
    "fmt"
    "log"
    "github.com/arkd0ng/go-utils/random"
)

func main() {
    // Generate a 32-character alphanumeric string
    // 32자 영숫자 문자열 생성
    str, err := random.GenString.Alnum(32)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Generated: %s\n", str)
    // Output example: Generated: a7B3xK9mP2qR5vN8zL1wD4cF6hJ0tY
}
```

### Range Length Example / 범위 길이 예제

```go
// Generate a string between 16 and 32 characters
// 16-32자 사이의 문자열 생성
str, err := random.GenString.Alnum(16, 32)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Length: %d, String: %s\n", len(str), str)
// Output example: Length: 24, String: x8K2mP9qR4vN7zL3wD5cF
```

---

## Method Reference / 메서드 참조

### Basic Methods / 기본 메서드

#### 1. Letters

Generates alphabetic characters only (a-z, A-Z).

알파벳 문자만 생성 (a-z, A-Z).

**Signature / 시그니처**:
```go
func (stringGenerator) Letters(length ...int) (string, error)
```

**Character Set / 문자 집합**: `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`

**Examples / 예제**:
```go
// Fixed length / 고정 길이
str1, err := random.GenString.Letters(10)
// Output: "AbCdEfGhIj"

// Range length / 범위 길이
str2, err := random.GenString.Letters(8, 16)
// Output: "KlMnOpQrStUv" (length between 8-16)
```

**Use Cases / 사용 사례**:
- Name generation / 이름 생성
- Text-only identifiers / 텍스트 전용 식별자
- Human-readable codes / 사람이 읽을 수 있는 코드

---

#### 2. Alnum

Generates alphanumeric characters (a-z, A-Z, 0-9).

영숫자 문자 생성 (a-z, A-Z, 0-9).

**Signature / 시그니처**:
```go
func (stringGenerator) Alnum(length ...int) (string, error)
```

**Character Set / 문자 집합**: `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`

**Examples / 예제**:
```go
// API Key generation / API 키 생성
apiKey, err := random.GenString.Alnum(32)
// Output: "a7B3xK9mP2qR5vN8zL1wD4cF6hJ0tY"

// Session ID / 세션 ID
sessionID, err := random.GenString.Alnum(64)
```

**Use Cases / 사용 사례**:
- API keys / API 키
- Session IDs / 세션 ID
- General-purpose tokens / 범용 토큰
- Database identifiers / 데이터베이스 식별자

---

#### 3. Digits

Generates numeric digits only (0-9).

숫자만 생성 (0-9).

**Signature / 시그니처**:
```go
func (stringGenerator) Digits(length ...int) (string, error)
```

**Character Set / 문자 집합**: `0123456789`

**Examples / 예제**:
```go
// 6-digit PIN / 6자리 PIN
pin, err := random.GenString.Digits(6)
// Output: "847293"

// OTP code / OTP 코드
otp, err := random.GenString.Digits(4, 8)
// Output: "52847" (length between 4-8)
```

**Use Cases / 사용 사례**:
- PIN codes / PIN 코드
- OTP (One-Time Password) / 일회용 비밀번호
- Verification codes / 인증 코드
- Phone number generation / 전화번호 생성

---

#### 4. Complex

Generates alphanumeric + all special characters.

영숫자 + 모든 특수 문자 생성.

**Signature / 시그니처**:
```go
func (stringGenerator) Complex(length ...int) (string, error)
```

**Character Set / 문자 집합**: `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/`

**Examples / 예제**:
```go
// Strong password / 강력한 비밀번호
password, err := random.GenString.Complex(16, 24)
// Output: "aB3!xK#9m@P2*qR&5v"

// Secure token / 보안 토큰
token, err := random.GenString.Complex(32)
```

**Use Cases / 사용 사례**:
- Strong passwords / 강력한 비밀번호
- High-security tokens / 고보안 토큰
- Encryption keys / 암호화 키

**Warning / 주의**: Contains special characters that may need escaping in URLs or shell commands.

URL이나 쉘 명령에서 이스케이프가 필요한 특수 문자를 포함합니다.

---

#### 5. Standard

Generates alphanumeric + safe special characters.

영숫자 + 안전한 특수 문자 생성.

**Signature / 시그니처**:
```go
func (stringGenerator) Standard(length ...int) (string, error)
```

**Character Set / 문자 집합**: `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*-_`

**Examples / 예제**:
```go
// User password / 사용자 비밀번호
password, err := random.GenString.Standard(12, 20)
// Output: "aB3@xK#9m-P2*qR"

// Safe token / 안전한 토큰
token, err := random.GenString.Standard(24)
```

**Use Cases / 사용 사례**:
- User passwords / 사용자 비밀번호
- Safe tokens / 안전한 토큰
- General secure strings / 일반 보안 문자열

---

### Case-Specific Methods / 대소문자 구분 메서드

#### 6. AlphaUpper

Uppercase letters only (A-Z).

대문자만 (A-Z).

**Signature / 시그니처**:
```go
func (stringGenerator) AlphaUpper(length ...int) (string, error)
```

**Examples / 예제**:
```go
// Ticket code / 티켓 코드
ticket, err := random.GenString.AlphaUpper(8)
// Output: "ABCDEFGH"

// Coupon code / 쿠폰 코드
coupon, err := random.GenString.AlphaUpper(10, 15)
```

**Use Cases / 사용 사례**:
- Ticket codes / 티켓 코드
- Coupon codes / 쿠폰 코드
- Uppercase-only identifiers / 대문자 전용 식별자

---

#### 7. AlphaLower

Lowercase letters only (a-z).

소문자만 (a-z).

**Signature / 시그니처**:
```go
func (stringGenerator) AlphaLower(length ...int) (string, error)
```

**Examples / 예제**:
```go
// Username / 사용자명
username, err := random.GenString.AlphaLower(8, 12)
// Output: "abcdefghij"

// Subdomain / 서브도메인
subdomain, err := random.GenString.AlphaLower(6)
```

**Use Cases / 사용 사례**:
- Usernames / 사용자명
- Subdomain names / 서브도메인 이름
- Lowercase-only identifiers / 소문자 전용 식별자

---

#### 8. AlnumUpper

Uppercase letters + digits (A-Z, 0-9).

대문자 + 숫자 (A-Z, 0-9).

**Signature / 시그니처**:
```go
func (stringGenerator) AlnumUpper(length ...int) (string, error)
```

**Examples / 예제**:
```go
// License key / 라이선스 키
license, err := random.GenString.AlnumUpper(20)
// Output: "ABC123DEF456GHI789JK"

// Product code / 제품 코드
product, err := random.GenString.AlnumUpper(16)
```

**Use Cases / 사용 사례**:
- License keys / 라이선스 키
- Product codes / 제품 코드
- Serial numbers / 시리얼 번호

---

#### 9. AlnumLower

Lowercase letters + digits (a-z, 0-9).

소문자 + 숫자 (a-z, 0-9).

**Signature / 시그니처**:
```go
func (stringGenerator) AlnumLower(length ...int) (string, error)
```

**Examples / 예제**:
```go
// Token / 토큰
token, err := random.GenString.AlnumLower(32)
// Output: "abc123def456ghi789jk012lmn345op"

// Identifier / 식별자
id, err := random.GenString.AlnumLower(16, 24)
```

**Use Cases / 사용 사례**:
- Tokens / 토큰
- Identifiers / 식별자
- Database keys / 데이터베이스 키

---

### Hexadecimal Methods / 16진수 메서드

#### 10. Hex

Uppercase hexadecimal (0-9, A-F).

대문자 16진수 (0-9, A-F).

**Signature / 시그니처**:
```go
func (stringGenerator) Hex(length ...int) (string, error)
```

**Examples / 예제**:
```go
// Color code / 색상 코드
color, err := random.GenString.Hex(6)
// Output: "A3F5C2"

// Hash value / 해시값
hash, err := random.GenString.Hex(32)
```

**Use Cases / 사용 사례**:
- Color codes / 색상 코드
- Hash values / 해시값
- Hexadecimal identifiers / 16진수 식별자

---

#### 11. HexLower

Lowercase hexadecimal (0-9, a-f).

소문자 16진수 (0-9, a-f).

**Signature / 시그니처**:
```go
func (stringGenerator) HexLower(length ...int) (string, error)
```

**Examples / 예제**:
```go
// UUID component / UUID 구성요소
uuid, err := random.GenString.HexLower(32)
// Output: "a3f5c2d1e4b7f8a9c0d3e6f9b2c5d8e1"

// Hash / 해시
hash, err := random.GenString.HexLower(40)
```

**Use Cases / 사용 사례**:
- UUID generation / UUID 생성
- Hash values / 해시값
- Lowercase hex identifiers / 소문자 16진수 식별자

---

### Encoding Methods / 인코딩 메서드

#### 12. Base64

Base64 character set (A-Z, a-z, 0-9, +, /).

Base64 문자 집합 (A-Z, a-z, 0-9, +, /).

**Signature / 시그니처**:
```go
func (stringGenerator) Base64(length ...int) (string, error)
```

**Examples / 예제**:
```go
// Base64-like string / Base64 형태 문자열
str, err := random.GenString.Base64(32)
// Output: "aB3+xK/9m=P2qR5vN8zL1wD4cF6hJ0"
```

**Use Cases / 사용 사례**:
- Base64-like encoding / Base64 형태 인코딩
- Data representation / 데이터 표현

---

#### 13. Base64URL

URL-safe Base64 (A-Z, a-z, 0-9, -, _).

URL-safe Base64 (A-Z, a-z, 0-9, -, _).

**Signature / 시그니처**:
```go
func (stringGenerator) Base64URL(length ...int) (string, error)
```

**Examples / 예제**:
```go
// URL-safe token / URL-safe 토큰
token, err := random.GenString.Base64URL(64)
// Output: "aB3-xK_9m-P2qR5vN8zL1wD4cF6hJ0tYuM3nP7qS9rT1wV5xZ8a"

// File name / 파일명
filename, err := random.GenString.Base64URL(16)
```

**Use Cases / 사용 사례**:
- URL-safe tokens / URL-safe 토큰
- File names / 파일명
- URL parameters / URL 매개변수

---

### Custom Method / 사용자 정의 메서드

#### 14. Custom

Custom character set.

사용자 정의 문자 집합.

**Signature / 시그니처**:
```go
func (stringGenerator) Custom(charset string, length ...int) (string, error)
```

**Examples / 예제**:
```go
// Custom charset / 사용자 정의 문자 집합
code, err := random.GenString.Custom("ACGT", 20)
// Output: "ACGTACGTACGTACGTACGT" (DNA sequence)

// Limited numbers / 제한된 숫자
lucky, err := random.GenString.Custom("13579", 6)
// Output: "137915"

// Custom alphabet / 사용자 정의 알파벳
vowels, err := random.GenString.Custom("AEIOU", 10, 15)
```

**Use Cases / 사용 사례**:
- Domain-specific codes / 도메인별 코드
- Custom alphabets / 사용자 정의 알파벳
- Specialized character sets / 특수화된 문자 집합

---

## Usage Patterns / 사용 패턴

### Pattern 1: Fixed Length / 고정 길이

```go
// Generate exactly 32 characters
// 정확히 32자 생성
str, err := random.GenString.Alnum(32)
```

### Pattern 2: Range Length / 범위 길이

```go
// Generate 16-32 characters
// 16-32자 생성
str, err := random.GenString.Alnum(16, 32)
```

### Pattern 3: Error Handling / 에러 처리

```go
// Proper error handling / 적절한 에러 처리
str, err := random.GenString.Alnum(32)
if err != nil {
    log.Printf("Failed to generate random string: %v", err)
    return err
}
```

### Pattern 4: Multiple Generations / 다중 생성

```go
// Generate multiple strings / 여러 문자열 생성
passwords := make([]string, 10)
for i := 0; i < 10; i++ {
    pwd, err := random.GenString.Complex(16, 24)
    if err != nil {
        return err
    }
    passwords[i] = pwd
}
```

---

## Common Use Cases / 일반적인 사용 사례

### Use Case 1: Password Generation / 비밀번호 생성

```go
// Strong password with special characters
// 특수 문자를 포함한 강력한 비밀번호
password, err := random.GenString.Complex(16, 24)
if err != nil {
    return err
}

// Medium-strength password
// 중간 강도 비밀번호
password2, err := random.GenString.Standard(12, 16)
```

### Use Case 2: API Key Generation / API 키 생성

```go
// 64-character API key
// 64자 API 키
apiKey, err := random.GenString.Alnum(64)
if err != nil {
    return err
}

fmt.Printf("API-KEY: %s\n", apiKey)
```

### Use Case 3: Session ID / 세션 ID

```go
// Session ID with timestamp
// 타임스탬프를 포함한 세션 ID
sessionID, err := random.GenString.Base64URL(32)
if err != nil {
    return err
}

fullSessionID := fmt.Sprintf("%d-%s", time.Now().Unix(), sessionID)
```

### Use Case 4: OTP Generation / OTP 생성

```go
// 6-digit OTP
// 6자리 OTP
otp, err := random.GenString.Digits(6)
if err != nil {
    return err
}

// Send OTP to user
sendOTP(userEmail, otp)
```

### Use Case 5: License Key / 라이선스 키

```go
// Format: XXXX-XXXX-XXXX-XXXX
// 형식: XXXX-XXXX-XXXX-XXXX
parts := make([]string, 4)
for i := 0; i < 4; i++ {
    part, err := random.GenString.AlnumUpper(4)
    if err != nil {
        return err
    }
    parts[i] = part
}

license := strings.Join(parts, "-")
// Output: "A3F5-C2D1-E4B7-F8A9"
```

---

## Error Handling / 에러 처리

### Common Errors / 일반적인 에러

#### Error 1: Invalid Length / 잘못된 길이

```go
// ❌ Negative length
str, err := random.GenString.Alnum(-5)
// Error: "minimum length cannot be negative: -5"

// ❌ Max < Min
str, err := random.GenString.Alnum(20, 10)
// Error: "maximum length (10) cannot be less than minimum length (20)"
```

#### Error 2: Empty Charset / 빈 문자 집합

```go
// ❌ Empty charset
str, err := random.GenString.Custom("", 10)
// Error: "charset cannot be empty"
```

#### Error 3: Missing Length / 길이 누락

```go
// ❌ No length argument
str, err := random.GenString.Alnum()
// Error: "at least one length argument is required"
```

### Error Handling Best Practices / 에러 처리 모범 사례

```go
func generateSecurePassword() (string, error) {
    password, err := random.GenString.Complex(16, 24)
    if err != nil {
        // Log the error / 에러 로깅
        log.Printf("Password generation failed: %v", err)

        // Return descriptive error / 설명적인 에러 반환
        return "", fmt.Errorf("failed to generate password: %w", err)
    }

    return password, nil
}
```

---

## Best Practices / 모범 사례

### 1. Choose Appropriate Length / 적절한 길이 선택

```go
// ✅ Good - Strong password
password, _ := random.GenString.Complex(16, 24)

// ❌ Bad - Too short for security
weak, _ := random.GenString.Complex(6)
```

### 2. Use Appropriate Method / 적절한 메서드 사용

```go
// ✅ Good - Use Digits for PIN
pin, _ := random.GenString.Digits(6)

// ❌ Bad - Using Alnum for PIN (contains letters)
pin, _ := random.GenString.Alnum(6)  // Could be "a3F5c2"
```

### 3. Handle Errors / 에러 처리

```go
// ✅ Good - Check and handle errors
str, err := random.GenString.Alnum(32)
if err != nil {
    log.Fatal(err)
}

// ❌ Bad - Ignore errors
str, _ := random.GenString.Alnum(32)  // Don't do this!
```

### 4. Use Range for Variability / 가변성을 위해 범위 사용

```go
// ✅ Good - Variable length adds unpredictability
token, _ := random.GenString.Base64URL(32, 64)

// ❌ Less secure - Always same length
token, _ := random.GenString.Base64URL(32)
```

---

## FAQ / 자주 묻는 질문

### Q1: Is this package cryptographically secure? / 이 패키지는 암호학적으로 안전한가요?

**A**: Yes. The package uses `crypto/rand` which provides cryptographically secure random generation suitable for security-sensitive applications.

**A**: 네. 이 패키지는 `crypto/rand`를 사용하여 보안이 중요한 애플리케이션에 적합한 암호학적으로 안전한 랜덤 생성을 제공합니다.

### Q2: Can I use this for passwords? / 비밀번호에 사용할 수 있나요?

**A**: Absolutely. Use `Complex()` or `Standard()` for password generation. We recommend 16-24 characters for strong passwords.

**A**: 물론입니다. 비밀번호 생성에는 `Complex()` 또는 `Standard()`를 사용하세요. 강력한 비밀번호를 위해 16-24자를 권장합니다.

### Q3: What's the difference between Complex and Standard? / Complex와 Standard의 차이는 무엇인가요?

**A**: `Complex()` includes all special characters (`!@#$%^&*()-_=+[]{}|;:,.<>?/`), while `Standard()` includes only safe characters (`!@#$%^&*-_`). Use `Complex()` for maximum security, `Standard()` for compatibility.

**A**: `Complex()`는 모든 특수 문자를 포함하고, `Standard()`는 안전한 문자만 포함합니다. 최대 보안을 위해서는 `Complex()`, 호환성을 위해서는 `Standard()`를 사용하세요.

### Q4: How do I generate a UUID? / UUID를 어떻게 생성하나요?

**A**: Use `HexLower(32)` and format it:

**A**: `HexLower(32)`를 사용하고 형식화하세요:

```go
hex, _ := random.GenString.HexLower(32)
uuid := fmt.Sprintf("%s-%s-%s-%s-%s",
    hex[0:8], hex[8:12], hex[12:16], hex[16:20], hex[20:32])
```

### Q5: Is it thread-safe? / 스레드 안전한가요?

**A**: Yes. All methods are safe for concurrent use as `crypto/rand.Reader` is thread-safe.

**A**: 네. `crypto/rand.Reader`가 스레드 안전하므로 모든 메서드는 동시 사용이 안전합니다.

### Q6: What's the maximum length I can generate? / 생성 가능한 최대 길이는 얼마인가요?

**A**: There's no hard limit, but very large lengths (>10MB) may cause memory issues. For typical use cases, lengths up to 1MB are safe.

**A**: 엄격한 제한은 없지만, 매우 큰 길이(>10MB)는 메모리 문제를 일으킬 수 있습니다. 일반적인 사용 사례의 경우 1MB까지는 안전합니다.

---

## Troubleshooting / 문제 해결

### Problem 1: "minimum length cannot be negative" / "최소 길이는 음수일 수 없습니다"

**Solution / 해결책**: Ensure length parameters are positive integers.

```go
// ❌ Wrong
str, err := random.GenString.Alnum(-5)

// ✅ Correct
str, err := random.GenString.Alnum(5)
```

### Problem 2: "maximum length cannot be less than minimum length" / "최대 길이는 최소 길이보다 작을 수 없습니다"

**Solution / 해결책**: Ensure max >= min.

```go
// ❌ Wrong
str, err := random.GenString.Alnum(20, 10)

// ✅ Correct
str, err := random.GenString.Alnum(10, 20)
```

### Problem 3: "charset cannot be empty" / "문자 집합은 비어있을 수 없습니다"

**Solution / 해결책**: Provide a non-empty charset for `Custom()`.

```go
// ❌ Wrong
str, err := random.GenString.Custom("", 10)

// ✅ Correct
str, err := random.GenString.Custom("ABC123", 10)
```

### Problem 4: Random strings are not random enough / 랜덤 문자열이 충분히 랜덤하지 않습니다

**Solution / 해결책**: This package uses `crypto/rand` which is cryptographically secure. If strings appear predictable, check that you're not using a fixed seed or deterministic logic elsewhere in your application.

**Solution / 해결책**: 이 패키지는 암호학적으로 안전한 `crypto/rand`를 사용합니다. 문자열이 예측 가능해 보이면 애플리케이션의 다른 곳에서 고정 시드나 결정론적 로직을 사용하고 있지 않은지 확인하세요.

---

## Support / 지원

For issues, questions, or contributions:

문제, 질문 또는 기여를 위해:

- **GitHub Issues**: [github.com/arkd0ng/go-utils/issues](https://github.com/arkd0ng/go-utils/issues)
- **Documentation**: [Main README](../../random/README.md)
- **Examples**: [examples/random_string/](../../examples/random_string/)

---

**Last Updated / 최종 업데이트**: 2025-10-10
**Version / 버전**: v1.0.008
**License / 라이선스**: MIT
