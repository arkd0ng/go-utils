# Random Package

A package for generating cryptographically secure random strings with various character sets.

랜덤 문자열을 생성하는 패키지로, 다양한 문자 집합으로 암호학적으로 안전한 랜덤 문자열을 생성합니다.

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/random
```

## Features / 주요 기능

Generate random strings with customizable length ranges and character sets.

사용자 정의 길이 범위와 문자 집합으로 랜덤 문자열을 생성합니다.

### Available Methods / 사용 가능한 메서드

All methods return `(string, error)` to handle potential errors during random string generation.

모든 메서드는 랜덤 문자열 생성 중 발생할 수 있는 에러를 처리하기 위해 `(string, error)`를 반환합니다.

#### Basic Methods / 기본 메서드

- **`GenString.Letters(length ...int) (string, error)`** - Alphabetic characters (a-z, A-Z) / 알파벳 (a-z, A-Z)
- **`GenString.Alnum(length ...int) (string, error)`** - Alphanumeric (a-z, A-Z, 0-9) / 영숫자 (a-z, A-Z, 0-9)
- **`GenString.Digits(length ...int) (string, error)`** - Numeric digits only (0-9) / 숫자만 (0-9)
- **`GenString.Complex(length ...int) (string, error)`** - Alphanumeric + all special characters / 영숫자 + 모든 특수문자
- **`GenString.Standard(length ...int) (string, error)`** - Alphanumeric + safe special characters (!@#$%^&*-_) / 영숫자 + 안전한 특수문자

#### Case-Specific Methods / 대소문자 구분 메서드

- **`GenString.AlphaUpper(length ...int) (string, error)`** - Uppercase letters only (A-Z) / 대문자만 (A-Z)
- **`GenString.AlphaLower(length ...int) (string, error)`** - Lowercase letters only (a-z) / 소문자만 (a-z)
- **`GenString.AlnumUpper(length ...int) (string, error)`** - Uppercase + digits (A-Z, 0-9) / 대문자 + 숫자 (A-Z, 0-9)
- **`GenString.AlnumLower(length ...int) (string, error)`** - Lowercase + digits (a-z, 0-9) / 소문자 + 숫자 (a-z, 0-9)

#### Hexadecimal Methods / 16진수 메서드

- **`GenString.Hex(length ...int) (string, error)`** - Uppercase hexadecimal (0-9, A-F) / 대문자 16진수 (0-9, A-F)
- **`GenString.HexLower(length ...int) (string, error)`** - Lowercase hexadecimal (0-9, a-f) / 소문자 16진수 (0-9, a-f)

#### Encoding Methods / 인코딩 메서드

- **`GenString.Base64(length ...int) (string, error)`** - Base64 character set (A-Z, a-z, 0-9, +, /) / Base64 문자 집합
- **`GenString.Base64URL(length ...int) (string, error)`** - URL-safe Base64 (A-Z, a-z, 0-9, -, _) / URL-safe Base64

#### Custom Method / 사용자 정의 메서드

- **`GenString.Custom(charset string, length ...int) (string, error)`** - Custom character set / 사용자 정의 문자 집합

## Usage / 사용법

### Basic Examples / 기본 예제

```go
package main

import (
    "fmt"
    "log"
    "github.com/arkd0ng/go-utils/random"
)

func main() {
    // Generate alphabetic string (32-128 characters)
    // 알파벳 문자열 생성 (32-128자)
    str1, err := random.GenString.Letters(32, 128)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(str1)

    // Generate alphanumeric string (32-128 characters)
    // 영숫자 문자열 생성 (32-128자)
    str2, err := random.GenString.Alnum(32, 128)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(str2)

    // Generate string with special characters (16-32 characters)
    // 특수 문자 포함 문자열 생성 (16-32자)
    str3, err := random.GenString.Complex(16, 32)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(str3)

    // Generate string with limited special characters (20-40 characters)
    // 제한된 특수 문자 포함 문자열 생성 (20-40자)
    str4, err := random.GenString.Standard(20, 40)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(str4)

    // Generate string with custom character set (10-20 characters)
    // 사용자 정의 문자 집합으로 문자열 생성 (10-20자)
    str5, err := random.GenString.Custom("ABC123xyz", 10, 20)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(str5)
}
```

### Fixed Length String / 고정 길이 문자열

To generate a string with a fixed length, provide a single argument:

고정 길이 문자열을 생성하려면 하나의 인자를 제공하세요:

```go
// Generate exactly 32 characters / 정확히 32자 생성
password, err := random.GenString.Alnum(32)
if err != nil {
    log.Fatal(err)
}

// Or use min and max with the same value / 또는 min과 max를 같은 값으로 사용
password2, err := random.GenString.Alnum(32, 32)
if err != nil {
    log.Fatal(err)
}
```

### Common Use Cases / 일반적인 사용 사례

```go
import (
    "log"
    "github.com/arkd0ng/go-utils/random"
)

// Generate a secure password / 안전한 비밀번호 생성
password, err := random.GenString.Complex(16, 24)
if err != nil {
    log.Fatal(err)
}

// Generate a random API key / 랜덤 API 키 생성
apiKey, err := random.GenString.Alnum(40)
if err != nil {
    log.Fatal(err)
}

// Generate a random username / 랜덤 사용자명 생성
username, err := random.GenString.AlphaLower(8, 12)
if err != nil {
    log.Fatal(err)
}

// Generate a PIN code / PIN 코드 생성
pin, err := random.GenString.Digits(6)
if err != nil {
    log.Fatal(err)
}

// Generate a hex color code / 16진수 색상 코드 생성
color, err := random.GenString.Hex(6) // e.g., "A3F5C2"
if err != nil {
    log.Fatal(err)
}

// Generate a license key / 라이선스 키 생성
license, err := random.GenString.AlnumUpper(20)
if err != nil {
    log.Fatal(err)
}

// Generate a URL-safe token / URL-safe 토큰 생성
token, err := random.GenString.Base64URL(32)
if err != nil {
    log.Fatal(err)
}

// Generate a custom verification code / 사용자 정의 인증 코드 생성
code, err := random.GenString.Custom("0123456789", 6)
if err != nil {
    log.Fatal(err)
}
```

## Character Sets / 문자 집합

| Method / 메서드 | Character Set / 문자 집합 | Use Case / 사용 사례 |
|-----------------|---------------------------|---------------------|
| **Letters** | `A-Z`, `a-z` | General text / 일반 텍스트 |
| **Alnum** | `A-Z`, `a-z`, `0-9` | General codes / 일반 코드 |
| **Digits** | `0-9` | PIN, OTP / PIN, OTP |
| **Complex** | `A-Z`, `a-z`, `0-9`, `!@#$%^&*()-_=+[]{}|;:,.<>?/` | Secure passwords / 강력한 비밀번호 |
| **Standard** | `A-Z`, `a-z`, `0-9`, `!@#$%^&*-_` | Safe passwords / 안전한 비밀번호 |
| **AlphaUpper** | `A-Z` | Ticket codes / 티켓 코드 |
| **AlphaLower** | `a-z` | Usernames / 사용자명 |
| **AlnumUpper** | `A-Z`, `0-9` | License keys / 라이선스 키 |
| **AlnumLower** | `a-z`, `0-9` | Tokens / 토큰 |
| **Hex** | `0-9`, `A-F` | Color codes / 색상 코드 |
| **HexLower** | `0-9`, `a-f` | UUID, hashes / UUID, 해시 |
| **Base64** | `A-Z`, `a-z`, `0-9`, `+`, `/` | Base64 encoding / Base64 인코딩 |
| **Base64URL** | `A-Z`, `a-z`, `0-9`, `-`, `_` | URL-safe tokens / URL-safe 토큰 |

## Security / 보안

The random string generator uses `crypto/rand` for cryptographically secure random generation, making it suitable for:

이 랜덤 문자열 생성기는 `crypto/rand`를 사용하여 암호학적으로 안전한 랜덤 생성을 제공하므로 다음 용도에 적합합니다:

- Password generation / 비밀번호 생성
- API key generation / API 키 생성
- Token generation / 토큰 생성
- Security-sensitive applications / 보안이 중요한 애플리케이션

## Testing / 테스트

Run the test suite:

테스트 실행:

```bash
go test -v
```

Run benchmarks:

벤치마크 실행:

```bash
go test -bench=.
```

## Examples / 예제

See the [examples directory](../examples/random_string/) for complete working examples.

완전한 실행 예제는 [examples 디렉토리](../examples/random_string/)를 참조하세요.

## License / 라이선스

MIT License - see the [LICENSE](../LICENSE) file for details.

MIT 라이선스 - 자세한 내용은 [LICENSE](../LICENSE) 파일을 참조하세요.
