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

#### Basic Methods / 기본 메서드

- **`GenString.Letters(min, max int)`** - Alphabetic characters (a-z, A-Z) / 알파벳 (a-z, A-Z)
- **`GenString.Alnum(min, max int)`** - Alphanumeric (a-z, A-Z, 0-9) / 영숫자 (a-z, A-Z, 0-9)
- **`GenString.Digits(min, max int)`** - Numeric digits only (0-9) / 숫자만 (0-9)
- **`GenString.Complex(min, max int)`** - Alphanumeric + all special characters / 영숫자 + 모든 특수문자
- **`GenString.Standard(min, max int)`** - Alphanumeric + safe special characters (!@#$%^&*-_) / 영숫자 + 안전한 특수문자

#### Case-Specific Methods / 대소문자 구분 메서드

- **`GenString.AlphaUpper(min, max int)`** - Uppercase letters only (A-Z) / 대문자만 (A-Z)
- **`GenString.AlphaLower(min, max int)`** - Lowercase letters only (a-z) / 소문자만 (a-z)
- **`GenString.AlnumUpper(min, max int)`** - Uppercase + digits (A-Z, 0-9) / 대문자 + 숫자 (A-Z, 0-9)
- **`GenString.AlnumLower(min, max int)`** - Lowercase + digits (a-z, 0-9) / 소문자 + 숫자 (a-z, 0-9)

#### Hexadecimal Methods / 16진수 메서드

- **`GenString.Hex(min, max int)`** - Uppercase hexadecimal (0-9, A-F) / 대문자 16진수 (0-9, A-F)
- **`GenString.HexLower(min, max int)`** - Lowercase hexadecimal (0-9, a-f) / 소문자 16진수 (0-9, a-f)

#### Encoding Methods / 인코딩 메서드

- **`GenString.Base64(min, max int)`** - Base64 character set (A-Z, a-z, 0-9, +, /) / Base64 문자 집합
- **`GenString.Base64URL(min, max int)`** - URL-safe Base64 (A-Z, a-z, 0-9, -, _) / URL-safe Base64

#### Custom Method / 사용자 정의 메서드

- **`GenString.Custom(charset string, min, max int)`** - Custom character set / 사용자 정의 문자 집합

## Usage / 사용법

### Basic Examples / 기본 예제

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/random"
)

func main() {
    // Generate alphabetic string (32-128 characters)
    // 알파벳 문자열 생성 (32-128자)
    str1 := random.GenString.Letters(32, 128)
    fmt.Println(str1)

    // Generate alphanumeric string (32-128 characters)
    // 영숫자 문자열 생성 (32-128자)
    str2 := random.GenString.Alnum(32, 128)
    fmt.Println(str2)

    // Generate string with special characters (16-32 characters)
    // 특수 문자 포함 문자열 생성 (16-32자)
    str3 := random.GenString.Complex(16, 32)
    fmt.Println(str3)

    // Generate string with limited special characters (20-40 characters)
    // 제한된 특수 문자 포함 문자열 생성 (20-40자)
    str4 := random.GenString.Standard(20, 40)
    fmt.Println(str4)

    // Generate string with custom character set (10-20 characters)
    // 사용자 정의 문자 집합으로 문자열 생성 (10-20자)
    str5 := random.GenString.Custom("ABC123xyz", 10, 20)
    fmt.Println(str5)
}
```

### Fixed Length String / 고정 길이 문자열

To generate a string with a fixed length, set `min` and `max` to the same value:

고정 길이 문자열을 생성하려면 `min`과 `max`를 같은 값으로 설정하세요:

```go
// Generate exactly 32 characters / 정확히 32자 생성
password := random.GenString.Alnum(32, 32)
```

### Common Use Cases / 일반적인 사용 사례

```go
import "github.com/arkd0ng/go-utils/random"

// Generate a secure password / 안전한 비밀번호 생성
password := random.GenString.Complex(16, 24)

// Generate a random API key / 랜덤 API 키 생성
apiKey := random.GenString.Alnum(40, 40)

// Generate a random username / 랜덤 사용자명 생성
username := random.GenString.AlphaLower(8, 12)

// Generate a PIN code / PIN 코드 생성
pin := random.GenString.Digits(6, 6)

// Generate a hex color code / 16진수 색상 코드 생성
color := random.GenString.Hex(6, 6) // e.g., "A3F5C2"

// Generate a license key / 라이선스 키 생성
license := random.GenString.AlnumUpper(20, 20)

// Generate a URL-safe token / URL-safe 토큰 생성
token := random.GenString.Base64URL(32, 32)

// Generate a custom verification code / 사용자 정의 인증 코드 생성
code := random.GenString.Custom("0123456789", 6, 6)
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
