# validation Package - Design Plan / 설계 계획

Design and architecture planning for the validation utility package.

검증 유틸리티 패키지의 설계 및 아키텍처 계획입니다.

**Version / 버전**: v1.13.001
**Created / 생성**: 2025-10-17
**Status / 상태**: Design Phase / 설계 단계

---

## 1. Package Overview / 패키지 개요

### Purpose / 목적

Provide extreme simplicity validation utilities that reduce 30-50 lines of validation code to just 2-3 lines with comprehensive validation rules, custom validators, and detailed error reporting.

30-50줄의 검증 코드를 단 2-3줄로 줄이는 극도로 간단한 검증 유틸리티를 제공하며, 포괄적인 검증 규칙, 사용자 정의 검증기, 상세한 에러 보고를 제공합니다.

### Key Features / 주요 기능

1. **Fluent API / Fluent API**
   - Chainable validation rules / 체이닝 가능한 검증 규칙
   - Readable and intuitive / 읽기 쉽고 직관적

2. **Built-in Validators / 내장 검증기**
   - 50+ pre-built validation rules / 50개 이상의 사전 구축된 검증 규칙
   - String, numeric, date, email, URL, etc. / 문자열, 숫자, 날짜, 이메일, URL 등

3. **Custom Validators / 사용자 정의 검증기**
   - Easy to create custom rules / 사용자 정의 규칙 쉽게 생성
   - Reusable validators / 재사용 가능한 검증기

4. **Struct Validation / 구조체 검증**
   - Tag-based validation / 태그 기반 검증
   - Nested struct support / 중첩 구조체 지원

5. **Detailed Error Messages / 상세한 에러 메시지**
   - Field-specific errors / 필드별 에러
   - Bilingual error messages / 이중 언어 에러 메시지
   - Custom error messages / 사용자 정의 에러 메시지

6. **Zero Dependencies / 제로 의존성**
   - Standard library only / 표준 라이브러리만 사용
   - No external packages / 외부 패키지 없음

### Target Users / 대상 사용자

- Go developers building web APIs / 웹 API를 구축하는 Go 개발자
- Backend services requiring input validation / 입력 검증이 필요한 백엔드 서비스
- Form validation in web applications / 웹 애플리케이션의 폼 검증
- Configuration validation / 설정 검증

---

## 2. Architecture / 아키텍처

### File Structure / 파일 구조

```
validation/
├── validator.go           # Core Validator type and fluent API
├── validator_test.go      # Tests for Validator
├── rules.go              # Built-in validation rules
├── rules_test.go         # Tests for rules
├── rules_string.go       # String-specific rules
├── rules_string_test.go  # Tests for string rules
├── rules_numeric.go      # Numeric rules
├── rules_numeric_test.go # Tests for numeric rules
├── rules_time.go         # Date/time rules
├── rules_time_test.go    # Tests for time rules
├── struct.go             # Struct validation with tags
├── struct_test.go        # Tests for struct validation
├── errors.go             # Error types and formatting
├── errors_test.go        # Tests for errors
├── types.go              # Type definitions
├── helpers.go            # Helper functions
├── options.go            # Configuration options
├── version.go            # Version constant
└── README.md             # Package documentation
```

### Core Components / 핵심 구성요소

#### 1. Validator Type / Validator 타입

```go
type Validator struct {
    value     interface{}
    fieldName string
    errors    []ValidationError
    stopOnError bool
}
```

**Methods / 메서드:**
- `New(value interface{}, fieldName string) *Validator`
- `Required() *Validator`
- `MinLength(n int) *Validator`
- `MaxLength(n int) *Validator`
- `Email() *Validator`
- `URL() *Validator`
- `Custom(fn func(interface{}) bool, message string) *Validator`
- `Validate() error`
- `GetErrors() []ValidationError`

#### 2. ValidationError Type / ValidationError 타입

```go
type ValidationError struct {
    Field   string
    Value   interface{}
    Rule    string
    Message string
}
```

#### 3. Rule Function / 규칙 함수

```go
type RuleFunc func(interface{}) bool
```

#### 4. Struct Validation / 구조체 검증

```go
type StructValidator struct {
    errors []ValidationError
}

func ValidateStruct(s interface{}) error
```

**Tag Format / 태그 형식:**
```go
type User struct {
    Name  string `validate:"required,minlen=3,maxlen=50"`
    Email string `validate:"required,email"`
    Age   int    `validate:"required,min=18,max=120"`
}
```

---

## 3. API Design / API 설계

### Fluent API Example / Fluent API 예제

```go
// Single field validation / 단일 필드 검증
err := validation.New(email, "email").
    Required().
    Email().
    MaxLength(100).
    Validate()

// Multiple fields / 여러 필드
v := validation.NewValidator()
v.Field(name, "name").Required().MinLength(3).MaxLength(50)
v.Field(email, "email").Required().Email()
v.Field(age, "age").Required().Min(18).Max(120)

if err := v.Validate(); err != nil {
    // Handle errors / 에러 처리
    for _, e := range v.GetErrors() {
        fmt.Printf("%s: %s\n", e.Field, e.Message)
    }
}
```

### Struct Validation Example / 구조체 검증 예제

```go
type User struct {
    Name     string `validate:"required,minlen=3,maxlen=50"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"required,min=18,max=120"`
    Password string `validate:"required,minlen=8,contains=upper,contains=lower,contains=digit"`
    Website  string `validate:"url,optional"`
}

user := User{
    Name:     "John",
    Email:    "john@example.com",
    Age:      25,
    Password: "Secret123",
    Website:  "https://example.com",
}

if err := validation.ValidateStruct(user); err != nil {
    // Handle validation errors / 검증 에러 처리
}
```

### Custom Validator Example / 사용자 정의 검증기 예제

```go
// Custom validation function / 사용자 정의 검증 함수
isPalindrome := func(v interface{}) bool {
    s := v.(string)
    reversed := ""
    for i := len(s) - 1; i >= 0; i-- {
        reversed += string(s[i])
    }
    return s == reversed
}

err := validation.New("racecar", "text").
    Custom(isPalindrome, "must be a palindrome").
    Validate()
```

---

## 4. Built-in Validation Rules / 내장 검증 규칙

### String Validators (20 rules) / 문자열 검증기 (20개 규칙)

1. **Required** - Not empty / 비어있지 않음
2. **MinLength(n)** - Minimum length / 최소 길이
3. **MaxLength(n)** - Maximum length / 최대 길이
4. **Length(n)** - Exact length / 정확한 길이
5. **Email** - Valid email format / 유효한 이메일 형식
6. **URL** - Valid URL format / 유효한 URL 형식
7. **Alpha** - Only letters / 문자만
8. **Alphanumeric** - Letters and numbers / 문자와 숫자
9. **Numeric** - Only numbers / 숫자만
10. **StartsWith(prefix)** - Starts with string / 문자열로 시작
11. **EndsWith(suffix)** - Ends with string / 문자열로 끝남
12. **Contains(substring)** - Contains substring / 부분 문자열 포함
13. **Regex(pattern)** - Matches regex / 정규식 일치
14. **UUID** - Valid UUID / 유효한 UUID
15. **JSON** - Valid JSON / 유효한 JSON
16. **Base64** - Valid Base64 / 유효한 Base64
17. **Lowercase** - All lowercase / 모두 소문자
18. **Uppercase** - All uppercase / 모두 대문자
19. **Phone** - Valid phone number / 유효한 전화번호
20. **CreditCard** - Valid credit card / 유효한 신용카드

### Numeric Validators (10 rules) / 숫자 검증기 (10개 규칙)

1. **Min(n)** - Minimum value / 최소값
2. **Max(n)** - Maximum value / 최대값
3. **Between(min, max)** - In range / 범위 내
4. **Positive** - Greater than 0 / 0보다 큼
5. **Negative** - Less than 0 / 0보다 작음
6. **Zero** - Equal to 0 / 0과 같음
7. **NonZero** - Not equal to 0 / 0이 아님
8. **Even** - Even number / 짝수
9. **Odd** - Odd number / 홀수
10. **MultipleOf(n)** - Multiple of n / n의 배수

### Date/Time Validators (8 rules) / 날짜/시간 검증기 (8개 규칙)

1. **After(date)** - After date / 날짜 이후
2. **Before(date)** - Before date / 날짜 이전
3. **Between(start, end)** - In date range / 날짜 범위 내
4. **Today** - Is today / 오늘임
5. **Past** - In the past / 과거임
6. **Future** - In the future / 미래임
7. **Weekday** - Is weekday / 평일임
8. **Weekend** - Is weekend / 주말임

### Collection Validators (7 rules) / 컬렉션 검증기 (7개 규칙)

1. **MinItems(n)** - Minimum items / 최소 항목 수
2. **MaxItems(n)** - Maximum items / 최대 항목 수
3. **Unique** - All unique / 모두 고유함
4. **In(values)** - In whitelist / 화이트리스트 내
5. **NotIn(values)** - Not in blacklist / 블랙리스트 외
6. **Each(validator)** - Validate each item / 각 항목 검증
7. **Empty** - Is empty / 비어있음

### Comparison Validators (5 rules) / 비교 검증기 (5개 규칙)

1. **Equal(value)** - Equal to value / 값과 같음
2. **NotEqual(value)** - Not equal to value / 값과 다름
3. **GreaterThan(value)** - Greater than / 보다 큼
4. **LessThan(value)** - Less than / 보다 작음
5. **OneOf(values)** - One of values / 값 중 하나

---

## 5. Error Handling / 에러 처리

### Error Types / 에러 타입

```go
type ValidationError struct {
    Field   string        // Field name / 필드 이름
    Value   interface{}   // Invalid value / 유효하지 않은 값
    Rule    string        // Failed rule / 실패한 규칙
    Message string        // Error message / 에러 메시지
}

type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string
func (ve ValidationErrors) HasField(field string) bool
func (ve ValidationErrors) GetField(field string) []ValidationError
func (ve ValidationErrors) ToMap() map[string][]string
```

### Error Message Format / 에러 메시지 형식

**English / 영문:**
- `"email is required"`
- `"name must be at least 3 characters"`
- `"age must be between 18 and 120"`

**Korean / 한글:**
- `"email은(는) 필수입니다"`
- `"name은(는) 최소 3자 이상이어야 합니다"`
- `"age은(는) 18에서 120 사이여야 합니다"`

### Custom Error Messages / 사용자 정의 에러 메시지

```go
validation.New(email, "email").
    Required().WithMessage("Email is required / 이메일은 필수입니다").
    Email().WithMessage("Invalid email format / 유효하지 않은 이메일 형식").
    Validate()
```

---

## 6. Performance Considerations / 성능 고려사항

### Optimization Strategies / 최적화 전략

1. **Stop on First Error / 첫 에러에서 중지**
   ```go
   validation.New(value, "field").
       StopOnError().
       Required().
       Email().
       Validate()
   ```

2. **Lazy Validation / 지연 검증**
   - Only validate when `Validate()` is called / `Validate()` 호출 시에만 검증
   - Avoid unnecessary allocations / 불필요한 할당 방지

3. **Regex Compilation / 정규식 컴파일**
   - Pre-compile common regex patterns / 일반적인 정규식 패턴 사전 컴파일
   - Cache compiled patterns / 컴파일된 패턴 캐시

4. **Memory Efficiency / 메모리 효율성**
   - Reuse error slices where possible / 가능한 경우 에러 슬라이스 재사용
   - Minimize allocations / 할당 최소화

---

## 7. Testing Strategy / 테스트 전략

### Coverage Goal / 커버리지 목표

- **Target / 목표**: 100% test coverage / 100% 테스트 커버리지
- **Critical / 중요**: All validation rules must be tested / 모든 검증 규칙이 테스트되어야 함

### Test Categories / 테스트 카테고리

1. **Unit Tests / 단위 테스트**
   - Each validation rule independently / 각 검증 규칙 독립적으로
   - Edge cases and boundary conditions / 엣지 케이스와 경계 조건

2. **Integration Tests / 통합 테스트**
   - Multiple rules chained together / 여러 규칙을 함께 체이닝
   - Struct validation scenarios / 구조체 검증 시나리오

3. **Benchmark Tests / 벤치마크 테스트**
   - Performance of individual rules / 개별 규칙의 성능
   - Complex validation scenarios / 복잡한 검증 시나리오

4. **Example Tests / 예제 테스트**
   - Documentation examples / 문서 예제
   - Real-world use cases / 실제 사용 사례

---

## 8. Documentation Requirements / 문서 요구사항

### Required Documentation / 필수 문서

1. **README.md**
   - Quick start guide / 빠른 시작 가이드
   - All validation rules table / 모든 검증 규칙 표
   - Basic examples / 기본 예제

2. **USER_MANUAL.md**
   - Comprehensive guide (2000+ lines) / 포괄적인 가이드 (2000줄 이상)
   - All validators with examples / 모든 검증기와 예제
   - Advanced patterns / 고급 패턴
   - Error handling / 에러 처리

3. **DEVELOPER_GUIDE.md**
   - Architecture details / 아키텍처 세부사항
   - How to add custom validators / 사용자 정의 검증기 추가 방법
   - Contributing guidelines / 기여 가이드라인
   - Testing guide / 테스트 가이드

4. **examples/validation/main.go**
   - Executable examples / 실행 가능한 예제
   - All major features demonstrated / 모든 주요 기능 시연
   - Detailed logging / 상세한 로깅

---

## 9. Before vs After / 이전 vs 이후

### ❌ Before: Standard Go Validation (50+ lines)

```go
type User struct {
    Name  string
    Email string
    Age   int
}

func validateUser(u User) []string {
    var errors []string

    // Name validation
    if u.Name == "" {
        errors = append(errors, "name is required")
    }
    if len(u.Name) < 3 {
        errors = append(errors, "name must be at least 3 characters")
    }
    if len(u.Name) > 50 {
        errors = append(errors, "name must be at most 50 characters")
    }

    // Email validation
    if u.Email == "" {
        errors = append(errors, "email is required")
    }
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(u.Email) {
        errors = append(errors, "email is invalid")
    }
    if len(u.Email) > 100 {
        errors = append(errors, "email must be at most 100 characters")
    }

    // Age validation
    if u.Age == 0 {
        errors = append(errors, "age is required")
    }
    if u.Age < 18 {
        errors = append(errors, "age must be at least 18")
    }
    if u.Age > 120 {
        errors = append(errors, "age must be at most 120")
    }

    return errors
}

// Usage
user := User{Name: "John", Email: "john@example.com", Age: 25}
if errors := validateUser(user); len(errors) > 0 {
    for _, err := range errors {
        fmt.Println(err)
    }
}
```

### ✅ After: validation Package (3 lines)

```go
type User struct {
    Name  string `validate:"required,minlen=3,maxlen=50"`
    Email string `validate:"required,email,maxlen=100"`
    Age   int    `validate:"required,min=18,max=120"`
}

// Usage
user := User{Name: "John", Email: "john@example.com", Age: 25}
if err := validation.ValidateStruct(user); err != nil {
    fmt.Println(err)
}
```

**Result / 결과**: **50+ lines → 3 lines** (94% reduction / 94% 감소)

---

## 10. Dependencies / 의존성

### External Dependencies / 외부 의존성

**NONE / 없음** - Standard library only / 표준 라이브러리만

### Standard Library Usage / 표준 라이브러리 사용

- `regexp` - Regular expression validation / 정규식 검증
- `strings` - String manipulation / 문자열 조작
- `strconv` - String conversion / 문자열 변환
- `time` - Date/time validation / 날짜/시간 검증
- `reflect` - Struct tag parsing / 구조체 태그 파싱
- `fmt` - Error formatting / 에러 포맷팅

---

## 11. Milestones / 마일스톤

### Phase 1: Core Implementation (v1.13.001-020)

- [ ] Package structure and types (v1.13.001-002)
- [ ] Validator type and fluent API (v1.13.003-005)
- [ ] String validators (v1.13.006-010)
- [ ] Numeric validators (v1.13.011-013)
- [ ] Basic tests (v1.13.014-017)
- [ ] README.md (v1.13.018-020)

### Phase 2: Advanced Features (v1.13.021-040)

- [ ] Date/time validators (v1.13.021-023)
- [ ] Collection validators (v1.13.024-026)
- [ ] Comparison validators (v1.13.027-028)
- [ ] Struct validation (v1.13.029-032)
- [ ] Custom validators (v1.13.033-035)
- [ ] Error handling (v1.13.036-038)
- [ ] Examples (v1.13.039-040)

### Phase 3: Documentation & Finalization (v1.13.041-060)

- [ ] USER_MANUAL.md (v1.13.041-048)
- [ ] DEVELOPER_GUIDE.md (v1.13.049-054)
- [ ] Performance benchmarks (v1.13.055-057)
- [ ] Final testing and coverage (v1.13.058-059)
- [ ] Merge to main (v1.13.060)

---

## 12. Success Criteria / 성공 기준

### Must Have / 필수 사항

- ✅ 100% test coverage / 100% 테스트 커버리지
- ✅ 50+ built-in validation rules / 50개 이상의 내장 검증 규칙
- ✅ Fluent API for chaining / 체이닝을 위한 Fluent API
- ✅ Struct validation with tags / 태그를 사용한 구조체 검증
- ✅ Custom validators support / 사용자 정의 검증기 지원
- ✅ Detailed error messages / 상세한 에러 메시지
- ✅ Bilingual documentation / 이중 언어 문서화
- ✅ Zero external dependencies / 외부 의존성 없음

### Quality Metrics / 품질 지표

- **Code Reduction / 코드 감소**: 50+ lines → 2-3 lines (95%+ reduction)
- **Performance / 성능**: < 100ns per simple validation
- **Memory / 메모리**: Minimal allocations
- **Usability / 사용성**: Intuitive and easy to learn

---

## 13. Future Enhancements / 향후 개선사항

### Potential Features / 잠재적 기능

1. **Conditional Validation / 조건부 검증**
   ```go
   validate:"required_if=Country,US"
   ```

2. **Cross-Field Validation / 필드 간 검증**
   ```go
   validate:"eqfield=ConfirmPassword"
   ```

3. **Localization / 현지화**
   - Support for multiple languages / 여러 언어 지원
   - Custom message templates / 사용자 정의 메시지 템플릿

4. **Async Validation / 비동기 검증**
   - For database lookups / 데이터베이스 조회용
   - API calls / API 호출

5. **Validation Groups / 검증 그룹**
   - Different rules for different scenarios / 시나리오별 다른 규칙

---

## Conclusion / 결론

This design plan provides a comprehensive roadmap for implementing the validation package with extreme simplicity, high quality, and zero dependencies. The package will follow all go-utils standards and conventions.

이 설계 계획은 극도로 간단하고, 높은 품질과 외부 의존성이 없는 validation 패키지를 구현하기 위한 포괄적인 로드맵을 제공합니다. 패키지는 모든 go-utils 표준과 규약을 따를 것입니다.

**Next Steps / 다음 단계:**
1. Create WORK_PLAN.md with detailed tasks / 상세 작업이 포함된 WORK_PLAN.md 생성
2. Begin Phase 1 implementation / Phase 1 구현 시작
3. Continuous testing and documentation / 지속적인 테스트와 문서화

---

**Document Version / 문서 버전**: v1.13.001
**Author / 작성자**: go-utils team
**Status / 상태**: Approved for Implementation / 구현 승인됨
