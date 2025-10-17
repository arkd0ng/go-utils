# validation - Fluent Validation for Go / Go를 위한 Fluent 검증

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-blue)](https://go.dev/)
[![Coverage](https://img.shields.io/badge/Coverage-97.7%25-brightgreen)](https://github.com/arkd0ng/go-utils)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Version](https://img.shields.io/badge/Version-v1.13.030-blue)](https://github.com/arkd0ng/go-utils)

**104+ validators** to reduce 20-30 lines of validation code to just 1-2 lines with fluent API.

**104개 이상의 검증기**로 20-30줄의 검증 코드를 단 1-2줄로 줄입니다.

## Design Philosophy / 설계 철학

**"30 lines → 2 lines"** - Extreme Simplicity

- Fluent API with method chaining / 메서드 체이닝으로 Fluent API
- Type-safe with Go generics / Go 제네릭으로 타입 안전
- Bilingual error messages (EN/KR) / 양방향 에러 메시지 (영어/한글)
- Zero external dependencies / 외부 의존성 제로
- 92.5%+ test coverage / 92.5% 이상 테스트 커버리지

## Quick Start / 빠른 시작

```go
import "github.com/arkd0ng/go-utils/validation"

// Simple validation / 간단한 검증
v := validation.New("john@example.com", "email")
v.Required().Email()
if err := v.Validate(); err != nil {
    log.Fatal(err)
}

// Multiple validations with chaining / 체이닝으로 여러 검증
v := validation.New(25, "age")
v.Positive().Min(18).Max(120)
err := v.Validate()

// Multi-field validation / 다중 필드 검증
mv := validation.NewValidator()

mv.Field(user.Name, "name").Required().MinLength(2).MaxLength(50)
mv.Field(user.Email, "email").Required().Email()
mv.Field(user.Age, "age").Positive().Min(18).Max(120)

if err := mv.Validate(); err != nil {
    // Handle validation errors
}
```

## Categories / 카테고리

### 1. String Validators (20개)

| Validator | Description | 설명 |
|-----------|-------------|------|
| `Required()` | Not empty | 비어있지 않음 |
| `MinLength(n)` | Min string length | 최소 길이 |
| `MaxLength(n)` | Max string length | 최대 길이 |
| `Length(n)` | Exact length | 정확한 길이 |
| `Email()` | Valid email | 유효한 이메일 |
| `URL()` | Valid URL | 유효한 URL |
| `Alpha()` | Only letters | 문자만 |
| `Alphanumeric()` | Letters and numbers | 문자와 숫자 |
| `Numeric()` | Only numbers | 숫자만 |
| `StartsWith(s)` | Starts with string | 문자열로 시작 |
| `EndsWith(s)` | Ends with string | 문자열로 끝남 |
| `Contains(s)` | Contains string | 문자열 포함 |
| `Regex(pattern)` | Match regex | 정규식 매칭 |
| `UUID()` | Valid UUID | 유효한 UUID |
| `JSON()` | Valid JSON | 유효한 JSON |
| `Base64()` | Valid Base64 | 유효한 Base64 |
| `Lowercase()` | All lowercase | 모두 소문자 |
| `Uppercase()` | All uppercase | 모두 대문자 |
| `Phone()` | Valid phone | 유효한 전화번호 |
| `CreditCard()` | Valid credit card | 유효한 신용카드 |

### 2. Numeric Validators (10개)

| Validator | Description | 설명 |
|-----------|-------------|------|
| `Min(n)` | Minimum value | 최소값 |
| `Max(n)` | Maximum value | 최대값 |
| `Between(min, max)` | Value range | 값 범위 |
| `Positive()` | Positive number | 양수 |
| `Negative()` | Negative number | 음수 |
| `Zero()` | Zero value | 0 |
| `NonZero()` | Non-zero value | 0이 아님 |
| `Even()` | Even number | 짝수 |
| `Odd()` | Odd number | 홀수 |
| `MultipleOf(n)` | Multiple of n | n의 배수 |

### 3. Collection Validators (10개)

| Validator | Description | 설명 |
|-----------|-------------|------|
| `In(...values)` | Value in list | 목록에 존재 |
| `NotIn(...values)` | Value not in list | 목록에 없음 |
| `ArrayLength(n)` | Exact array length | 정확한 배열 길이 |
| `ArrayMinLength(n)` | Min array length | 최소 배열 길이 |
| `ArrayMaxLength(n)` | Max array length | 최대 배열 길이 |
| `ArrayNotEmpty()` | Array not empty | 배열 비어있지 않음 |
| `ArrayUnique()` | Unique elements | 고유한 요소 |
| `MapHasKey(key)` | Map has key | 맵에 키 존재 |
| `MapHasKeys(...keys)` | Map has all keys | 맵에 모든 키 존재 |
| `MapNotEmpty()` | Map not empty | 맵 비어있지 않음 |

### 4. Comparison Validators (10개)

| Validator | Description | 설명 |
|-----------|-------------|------|
| `Equals(value)` | Equal to value | 값과 동일 |
| `NotEquals(value)` | Not equal to value | 값과 다름 |
| `GreaterThan(n)` | Greater than | 보다 큼 |
| `GreaterThanOrEqual(n)` | Greater or equal | 크거나 같음 |
| `LessThan(n)` | Less than | 보다 작음 |
| `LessThanOrEqual(n)` | Less or equal | 작거나 같음 |
| `Before(time)` | Before time | 시간 이전 |
| `After(time)` | After time | 시간 이후 |
| `BeforeOrEqual(time)` | Before or equal time | 시간 이전이거나 같음 |
| `AfterOrEqual(time)` | After or equal time | 시간 이후이거나 같음 |

## Advanced Features / 고급 기능

### Stop on First Error / 첫 에러에서 중지

```go
v := validation.New("", "email")
v.StopOnError().Required().Email().MaxLength(100)
// Stops at Required() if empty
// 비어있으면 Required()에서 중지
```

### Custom Error Messages / 사용자 정의 에러 메시지

```go
v := validation.New(user.Age, "age")
v.Min(18).WithMessage("You must be at least 18 years old")
v.Max(120).WithMessage("Invalid age")
```

### Custom Validators / 사용자 정의 검증기

```go
v := validation.New(password, "password")
v.Custom(func(val interface{}) bool {
    s := val.(string)
    return strings.ContainsAny(s, "!@#$%")
}, "Password must contain special characters")
```

### Multi-Field Validation / 다중 필드 검증

```go
mv := validation.NewValidator()

mv.Field(user.Name, "name").
    Required().
    MinLength(2).
    MaxLength(50)

mv.Field(user.Email, "email").
    Required().
    Email().
    MaxLength(100)

mv.Field(user.Password, "password").
    Required().
    MinLength(8).
    Regex(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)`)

mv.Field(user.Age, "age").
    Positive().
    Between(18, 120)

if err := mv.Validate(); err != nil {
    errors := err.(validation.ValidationErrors)
    for _, e := range errors {
        fmt.Printf("%s: %s\n", e.Field, e.Message)
    }
}
```

## Real-World Examples / 실제 사용 예제

### User Registration / 사용자 등록

```go
type User struct {
    Username string
    Email    string
    Password string
    Age      int
    Country  string
}

func ValidateUser(user User) error {
    mv := validation.NewValidator()

    mv.Field(user.Username, "username").
        Required().
        MinLength(3).
        MaxLength(20).
        Alphanumeric()

    mv.Field(user.Email, "email").
        Required().
        Email().
        MaxLength(100)

    mv.Field(user.Password, "password").
        Required().
        MinLength(8).
        MaxLength(100).
        Regex(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]`)

    mv.Field(user.Age, "age").
        Positive().
        Between(13, 120)

    mv.Field(user.Country, "country").
        Required().
        In("US", "KR", "JP", "CN", "UK", "FR", "DE")

    return mv.Validate()
}
```

### API Request Validation / API 요청 검증

```go
type CreatePostRequest struct {
    Title    string
    Content  string
    Tags     []string
    Category string
}

func ValidateCreatePost(req CreatePostRequest) error {
    mv := validation.NewValidator()

    mv.Field(req.Title, "title").
        Required().
        MinLength(5).
        MaxLength(100)

    mv.Field(req.Content, "content").
        Required().
        MinLength(20).
        MaxLength(5000)

    mv.Field(req.Tags, "tags").
        ArrayNotEmpty().
        ArrayMinLength(1).
        ArrayMaxLength(5).
        ArrayUnique()

    mv.Field(req.Category, "category").
        Required().
        In("tech", "business", "lifestyle", "news")

    return mv.Validate()
}
```

### Configuration Validation / 설정 검증

```go
type Config struct {
    Port        int
    Host        string
    DatabaseURL string
    Timeout     int
    Features    map[string]bool
}

func ValidateConfig(cfg Config) error {
    mv := validation.NewValidator()

    mv.Field(cfg.Port, "port").
        Positive().
        Between(1, 65535)

    mv.Field(cfg.Host, "host").
        Required().
        URL()

    mv.Field(cfg.DatabaseURL, "database_url").
        Required().
        StartsWith("postgres://")

    mv.Field(cfg.Timeout, "timeout").
        Positive().
        Between(1, 3600)

    mv.Field(cfg.Features, "features").
        MapNotEmpty().
        MapHasKeys("auth", "logging")

    return mv.Validate()
}
```

## Error Handling / 에러 처리

```go
err := mv.Validate()
if err != nil {
    // Type assertion to ValidationErrors
    // ValidationErrors로 타입 단언
    validationErrs := err.(validation.ValidationErrors)

    // Get all errors / 모든 에러 가져오기
    for _, e := range validationErrs {
        fmt.Printf("Field: %s, Rule: %s, Message: %s\n",
            e.Field, e.Rule, e.Message)
    }

    // Check specific field / 특정 필드 확인
    if validationErrs.HasField("email") {
        fmt.Println("Email validation failed")
    }

    // Get errors for specific field / 특정 필드의 에러 가져오기
    emailErrs := validationErrs.GetField("email")
    for _, e := range emailErrs {
        fmt.Println(e.Message)
    }

    // Get first error / 첫 번째 에러 가져오기
    firstErr := validationErrs.First()

    // Get error count / 에러 개수 가져오기
    count := validationErrs.Count()

    // Convert to map / 맵으로 변환
    errMap := validationErrs.ToMap()
}
```

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/validation
```

## Requirements / 요구사항

- Go 1.18 or higher / Go 1.18 이상

## Performance / 성능

- Zero allocation for simple validations / 간단한 검증은 할당 없음
- Efficient regex caching / 효율적인 정규식 캐싱
- Minimal reflection usage / 최소한의 reflection 사용
- ~92.5% test coverage / ~92.5% 테스트 커버리지

## Documentation / 문서

- [User Manual](../docs/validation/USER_MANUAL.md) - Comprehensive guide / 포괄적인 가이드
- [Developer Guide](../docs/validation/DEVELOPER_GUIDE.md) - Architecture and internals / 아키텍처와 내부 구조
- [Examples](../examples/validation/main.go) - Executable examples / 실행 가능한 예제

## License / 라이선스

MIT License

## Contributing / 기여

Contributions are welcome! / 기여를 환영합니다!

## Version / 버전

Current version: **v1.13.013**

---

**Built with ❤️ for Go developers** / **Go 개발자를 위해 ❤️로 제작**
