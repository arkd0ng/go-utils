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

- **`GenString.Alpha(min, max int)`** - Alphabetic characters only (a-z, A-Z) / 알파벳만 (a-z, A-Z)
- **`GenString.AlphaNum(min, max int)`** - Alphanumeric characters (a-z, A-Z, 0-9) / 영숫자 (a-z, A-Z, 0-9)
- **`GenString.AlphaNumSpecial(min, max int)`** - Alphanumeric + all special characters / 영숫자 + 모든 특수문자
- **`GenString.AlphaNumSpecialLimited(min, max int)`** - Alphanumeric + limited special characters (!@#$%^&*-_) / 영숫자 + 제한된 특수문자
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
    str1 := random.GenString.Alpha(32, 128)
    fmt.Println(str1)

    // Generate alphanumeric string (32-128 characters)
    // 영숫자 문자열 생성 (32-128자)
    str2 := random.GenString.AlphaNum(32, 128)
    fmt.Println(str2)

    // Generate string with special characters (16-32 characters)
    // 특수 문자 포함 문자열 생성 (16-32자)
    str3 := random.GenString.AlphaNumSpecial(16, 32)
    fmt.Println(str3)

    // Generate string with limited special characters (20-40 characters)
    // 제한된 특수 문자 포함 문자열 생성 (20-40자)
    str4 := random.GenString.AlphaNumSpecialLimited(20, 40)
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
password := random.GenString.AlphaNum(32, 32)
```

### Common Use Cases / 일반적인 사용 사례

```go
import "github.com/arkd0ng/go-utils/random"

// Generate a secure password / 안전한 비밀번호 생성
password := random.GenString.AlphaNumSpecial(16, 24)

// Generate a random API key / 랜덤 API 키 생성
apiKey := random.GenString.AlphaNum(40, 40)

// Generate a random username / 랜덤 사용자명 생성
username := random.GenString.Alpha(8, 12)

// Generate a verification code with numbers only / 숫자로만 인증 코드 생성
code := random.GenString.Custom("0123456789", 6, 6)
```

## Character Sets / 문자 집합

- **Alpha**: `A-Z`, `a-z`
- **AlphaNum**: `A-Z`, `a-z`, `0-9`
- **AlphaNumSpecial**: `A-Z`, `a-z`, `0-9`, `!@#$%^&*()-_=+[]{}|;:,.<>?/`
- **AlphaNumSpecialLimited**: `A-Z`, `a-z`, `0-9`, `!@#$%^&*-_`

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
