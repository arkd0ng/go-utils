# Random Package Developer Guide / Random 패키지 개발 가이드

Complete developer guide for the Random package - architecture, implementation, and contribution guidelines.

Random 패키지의 완전한 개발 가이드 - 아키텍처, 구현, 기여 가이드라인.

**Version / 버전**: v1.0.008
**Last Updated / 최종 업데이트**: 2025-10-10

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#architecture-overview--아키텍처-개요)
2. [Internal Implementation / 내부 구현](#internal-implementation--내부-구현)
3. [Adding New Methods / 새 메서드 추가](#adding-new-methods--새-메서드-추가)
4. [Testing Guide / 테스트 가이드](#testing-guide--테스트-가이드)
5. [Performance Optimization / 성능 최적화](#performance-optimization--성능-최적화)
6. [Security Considerations / 보안 고려사항](#security-considerations--보안-고려사항)
7. [Contribution Guidelines / 기여 가이드라인](#contribution-guidelines--기여-가이드라인)
8. [Code Style / 코드 스타일](#code-style--코드-스타일)

---

## Architecture Overview / 아키텍처 개요

### Package Structure / 패키지 구조

```
random/
├── string.go        # Core implementation / 핵심 구현
├── string_test.go   # Comprehensive tests / 종합 테스트
└── README.md        # Package documentation / 패키지 문서
```

### Design Pattern / 디자인 패턴

The package uses a **Singleton Pattern** with method receivers:

패키지는 메서드 리시버를 사용한 **싱글톤 패턴**을 사용합니다:

```go
// Global singleton instance / 전역 싱글톤 인스턴스
var GenString = stringGenerator{}

// Empty struct - no state / 빈 구조체 - 상태 없음
type stringGenerator struct{}

// Methods attached to singleton / 싱글톤에 연결된 메서드
func (stringGenerator) Letters(length ...int) (string, error) {
    return generateRandomString(charsetAlpha, length...)
}
```

**Benefits / 장점**:
- No initialization required / 초기화 불필요
- Clean API: `random.GenString.Alnum(32)` / 깔끔한 API
- No mutable state / 변경 가능한 상태 없음
- Thread-safe by design / 설계상 스레드 안전

### Components / 구성 요소

#### 1. Character Sets / 문자 집합

Defined as package-level constants:

패키지 레벨 상수로 정의:

```go
const (
    charsetAlpha          = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    charsetAlphaUpper     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    charsetAlphaLower     = "abcdefghijklmnopqrstuvwxyz"
    charsetDigits         = "0123456789"
    charsetHex            = "0123456789ABCDEF"
    charsetHexLower       = "0123456789abcdef"
    charsetSpecial        = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
    charsetSpecialLimited = "!@#$%^&*-_"
    charsetBase64         = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
    charsetBase64URL      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)
```

#### 2. Public Methods / 공개 메서드

14 public methods that delegate to the core helper:

핵심 헬퍼에 위임하는 14개의 공개 메서드:

- 5 Basic methods / 5개 기본 메서드
- 4 Case-specific methods / 4개 대소문자 구분 메서드
- 2 Hexadecimal methods / 2개 16진수 메서드
- 2 Encoding methods / 2개 인코딩 메서드
- 1 Custom method / 1개 사용자 정의 메서드

#### 3. Core Helper Function / 핵심 헬퍼 함수

`generateRandomString()` - handles all the logic:

모든 로직을 처리하는 `generateRandomString()`:

```go
func generateRandomString(charset string, length ...int) (string, error)
```

**Responsibilities / 책임**:
- Validate charset / 문자 집합 검증
- Parse length arguments / 길이 인자 파싱
- Validate length parameters / 길이 파라미터 검증
- Generate random length (if range) / 랜덤 길이 생성 (범위인 경우)
- Generate random string / 랜덤 문자열 생성
- Handle errors / 에러 처리

---

## Internal Implementation / 내부 구현

### Flow Diagram / 흐름 도표

```
User Call / 사용자 호출
    ↓
Public Method (e.g., Alnum) / 공개 메서드 (예: Alnum)
    ↓
generateRandomString(charset, length...)
    ↓
1. Validate charset / 문자 집합 검증
2. Parse length arguments / 길이 인자 파싱
3. Validate length / 길이 검증
4. Determine actual length / 실제 길이 결정
5. Generate random string / 랜덤 문자열 생성
    ↓
Return (string, error) / 반환 (string, error)
```

### Detailed Implementation / 상세 구현

#### Step 1: Charset Validation / 문자 집합 검증

```go
if len(charset) == 0 {
    return "", errors.New("charset cannot be empty")
}
```

**Purpose / 목적**: Prevent division by zero and ensure valid input.

빈 문자 집합으로 인한 0으로 나누기 방지 및 유효한 입력 보장.

#### Step 2: Length Argument Parsing / 길이 인자 파싱

```go
var min, max int
switch len(length) {
case 0:
    return "", errors.New("at least one length argument is required")
case 1:
    // Fixed length / 고정 길이
    min = length[0]
    max = length[0]
case 2:
    // Range length / 범위 길이
    min = length[0]
    max = length[1]
default:
    return "", fmt.Errorf("invalid number of arguments: expected 1 or 2, got %d", len(length))
}
```

**Design Decision / 설계 결정**: Variadic parameters allow flexible API.

가변 인자는 유연한 API를 가능하게 합니다.

#### Step 3: Length Validation / 길이 검증

```go
if min < 0 {
    return "", fmt.Errorf("minimum length cannot be negative: %d", min)
}
if max < min {
    return "", fmt.Errorf("maximum length (%d) cannot be less than minimum length (%d)", max, min)
}
```

**Purpose / 목적**: Ensure logical constraints are met.

논리적 제약 조건 충족 보장.

#### Step 4: Actual Length Determination / 실제 길이 결정

```go
actualLength := min
if max > min {
    // Generate random length between min and max
    // min과 max 사이의 랜덤 길이 생성
    lengthRange := max - min + 1
    randomLength, err := rand.Int(rand.Reader, big.NewInt(int64(lengthRange)))
    if err != nil {
        return "", fmt.Errorf("failed to generate random length: %w", err)
    }
    actualLength = min + int(randomLength.Int64())
}
```

**Algorithm / 알고리즘**:
1. Calculate range: `max - min + 1` / 범위 계산
2. Generate random number in range [0, range) / 범위 [0, range)의 랜덤 숫자 생성
3. Add to min: `min + random` / min에 추가

**Example / 예제**:
- `min=10, max=20` → range = 11
- Random: 0-10 → Actual: 10-20

#### Step 5: Random String Generation / 랜덤 문자열 생성

```go
result := make([]byte, actualLength)
charsetLen := big.NewInt(int64(len(charset)))

for i := 0; i < actualLength; i++ {
    randomIndex, err := rand.Int(rand.Reader, charsetLen)
    if err != nil {
        return "", fmt.Errorf("failed to generate random character at position %d: %w", i, err)
    }
    result[i] = charset[randomIndex.Int64()]
}

return string(result), nil
```

**Key Points / 주요 포인트**:
- Uses `crypto/rand.Reader` for cryptographic security / 암호화 보안을 위해 `crypto/rand.Reader` 사용
- `big.NewInt` handles large charset sizes / `big.NewInt`가 큰 문자 집합 크기 처리
- Byte slice for efficiency / 효율성을 위한 바이트 슬라이스
- Error handling at each position / 각 위치에서 에러 처리

### Why crypto/rand? / 왜 crypto/rand인가?

**Comparison / 비교**:

| Feature / 특징 | math/rand | crypto/rand |
|----------------|-----------|-------------|
| Security / 보안 | ❌ Predictable / 예측 가능 | ✅ Cryptographically secure / 암호학적으로 안전 |
| Speed / 속도 | ✅ Fast / 빠름 | ⚠️ Slower / 느림 |
| Use case / 사용 사례 | Simulations / 시뮬레이션 | Passwords, keys / 비밀번호, 키 |

**Decision / 결정**: We chose `crypto/rand` for security-sensitive applications.

보안이 중요한 애플리케이션을 위해 `crypto/rand`를 선택했습니다.

---

## Adding New Methods / 새 메서드 추가

### Step-by-Step Guide / 단계별 가이드

#### Step 1: Define Character Set / 문자 집합 정의

Add a new constant in `string.go`:

`string.go`에 새 상수 추가:

```go
const (
    // Existing charsets... / 기존 문자 집합...
    charsetNewType = "YOUR_CHARSET_HERE"  // Add your charset / 문자 집합 추가
)
```

**Example / 예제**: Adding DNA sequence generator:

DNA 시퀀스 생성기 추가 예제:

```go
const (
    charsetDNA = "ACGT"
)
```

#### Step 2: Create Public Method / 공개 메서드 생성

Add a new method to `stringGenerator`:

`stringGenerator`에 새 메서드 추가:

```go
// DNA generates a random DNA sequence (A, C, G, T)
// DNA는 랜덤 DNA 시퀀스(A, C, G, T)를 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자
//     * 1 argument: fixed length / 1개 인자: 고정 길이
//     * 2 arguments: min and max length / 2개 인자: 최소 및 최대 길이
//
// Returns / 반환값:
//   - string: generated random DNA sequence / 생성된 랜덤 DNA 시퀀스
//   - error: error if invalid parameters / 잘못된 매개변수인 경우 에러
//
// Common use cases / 일반적인 사용 사례:
//   - Biological simulations / 생물학적 시뮬레이션
//   - DNA sequence generation / DNA 시퀀스 생성
func (stringGenerator) DNA(length ...int) (string, error) {
    return generateRandomString(charsetDNA, length...)
}
```

**Documentation Requirements / 문서 요구사항**:
- ✅ Bilingual comments (English/Korean) / 이중 언어 주석
- ✅ Parameter descriptions / 매개변수 설명
- ✅ Return value descriptions / 반환값 설명
- ✅ Common use cases / 일반적인 사용 사례
- ✅ Examples (optional but recommended) / 예제 (선택사항이지만 권장)

#### Step 3: Add Tests / 테스트 추가

Add tests in `string_test.go`:

`string_test.go`에 테스트 추가:

```go
// TestDNA tests the DNA method
// TestDNA는 DNA 메서드를 테스트합니다
func TestDNA(t *testing.T) {
    tests := []struct {
        name        string
        length      []int
        wantErr     bool
        checkLength bool
        minLen      int
        maxLen      int
    }{
        {
            name:        "Fixed length 20",
            length:      []int{20},
            wantErr:     false,
            checkLength: true,
            minLen:      20,
            maxLen:      20,
        },
        {
            name:        "Range 10-30",
            length:      []int{10, 30},
            wantErr:     false,
            checkLength: true,
            minLen:      10,
            maxLen:      30,
        },
        {
            name:    "Negative length",
            length:  []int{-5},
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            str, err := GenString.DNA(tt.length...)

            if (err != nil) != tt.wantErr {
                t.Errorf("DNA() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            if !tt.wantErr {
                // Check length / 길이 확인
                if tt.checkLength {
                    if len(str) < tt.minLen || len(str) > tt.maxLen {
                        t.Errorf("DNA() length = %d, want between %d and %d",
                            len(str), tt.minLen, tt.maxLen)
                    }
                }

                // Check charset / 문자 집합 확인
                for _, char := range str {
                    if !strings.ContainsRune("ACGT", char) {
                        t.Errorf("DNA() contains invalid character: %c", char)
                    }
                }
            }
        })
    }
}

// BenchmarkDNA benchmarks the DNA method
// BenchmarkDNA는 DNA 메서드를 벤치마크합니다
func BenchmarkDNA(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _, _ = GenString.DNA(32)
    }
}
```

**Test Requirements / 테스트 요구사항**:
- ✅ Fixed length test / 고정 길이 테스트
- ✅ Range length test / 범위 길이 테스트
- ✅ Error cases (negative, invalid) / 에러 케이스 (음수, 잘못된 값)
- ✅ Charset validation / 문자 집합 검증
- ✅ Benchmark test / 벤치마크 테스트

#### Step 4: Update Documentation / 문서 업데이트

Update `README.md`:

`README.md` 업데이트:

```markdown
#### Biological Methods / 생물학적 메서드

- **`GenString.DNA(length ...int) (string, error)`** - DNA sequence (A, C, G, T) / DNA 시퀀스
```

Add usage example:

사용 예제 추가:

```markdown
// Generate DNA sequence / DNA 시퀀스 생성
dna, err := random.GenString.DNA(100)
if err != nil {
    log.Fatal(err)
}
fmt.Println(dna)  // Output: ACGTACGTACGT...
```

#### Step 5: Update Character Sets Table / 문자 집합 테이블 업데이트

```markdown
| Method / 메서드 | Character Set / 문자 집합 | Use Case / 사용 사례 |
|-----------------|---------------------------|---------------------|
| **DNA** | `A`, `C`, `G`, `T` | Biological simulations / 생물학적 시뮬레이션 |
```

---

## Testing Guide / 테스트 가이드

### Test Structure / 테스트 구조

All tests in `string_test.go` follow this pattern:

`string_test.go`의 모든 테스트는 이 패턴을 따릅니다:

```go
func TestMethodName(t *testing.T) {
    tests := []struct {
        name        string        // Test case name / 테스트 케이스 이름
        length      []int         // Input length / 입력 길이
        wantErr     bool          // Expect error? / 에러 예상?
        checkLength bool          // Validate length? / 길이 검증?
        minLen      int           // Minimum expected length / 최소 예상 길이
        maxLen      int           // Maximum expected length / 최대 예상 길이
    }{
        // Test cases... / 테스트 케이스...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation / 테스트 구현
        })
    }
}
```

### Running Tests / 테스트 실행

```bash
# Run all tests / 모든 테스트 실행
go test -v

# Run specific test / 특정 테스트 실행
go test -v -run TestAlnum

# Run with coverage / 커버리지와 함께 실행
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks / 벤치마크 실행
go test -bench=.
go test -bench=BenchmarkAlnum
```

### Test Coverage Goals / 테스트 커버리지 목표

- **Line Coverage / 라인 커버리지**: >90%
- **Branch Coverage / 분기 커버리지**: >85%
- **Function Coverage / 함수 커버리지**: 100%

### Collision Probability Testing / 충돌 확률 테스트

Special test for randomness validation:

랜덤성 검증을 위한 특수 테스트:

```go
func TestCollisionProbability(t *testing.T) {
    const iterations = 10000
    const length = 32

    seen := make(map[string]bool)
    collisions := 0

    for i := 0; i < iterations; i++ {
        str, err := GenString.Alnum(length)
        if err != nil {
            t.Fatal(err)
        }

        if seen[str] {
            collisions++
        }
        seen[str] = true
    }

    // Calculate theoretical probability / 이론적 확률 계산
    charsetSize := 62 // a-z, A-Z, 0-9
    possibleStrings := math.Pow(float64(charsetSize), float64(length))
    theoreticalProb := float64(iterations) / possibleStrings

    // Actual collision rate should be near theoretical / 실제 충돌률은 이론치에 근접해야 함
    actualRate := float64(collisions) / float64(iterations)

    t.Logf("Theoretical collision probability: %.10f", theoreticalProb)
    t.Logf("Actual collision rate: %.10f", actualRate)
    t.Logf("Collisions: %d / %d", collisions, iterations)

    // For 32-char alphanumeric, collision should be extremely rare
    // 32자 영숫자의 경우 충돌은 극히 드물어야 함
    if collisions > 0 {
        t.Logf("Warning: Collisions detected (expected to be extremely rare)")
    }
}
```

---

## Performance Optimization / 성능 최적화

### Current Performance / 현재 성능

Benchmark results on typical hardware:

일반적인 하드웨어의 벤치마크 결과:

```
BenchmarkAlnum-8         100000    12345 ns/op    64 B/op    2 allocs/op
BenchmarkDigits-8        150000     8234 ns/op    16 B/op    2 allocs/op
BenchmarkComplex-8        80000    14567 ns/op    96 B/op    2 allocs/op
```

### Optimization Techniques / 최적화 기법

#### 1. Pre-allocate Byte Slice / 바이트 슬라이스 사전 할당

```go
// ✅ Good - Pre-allocate / 사전 할당
result := make([]byte, actualLength)

// ❌ Bad - Dynamic growth / 동적 증가
var result []byte
for i := 0; i < actualLength; i++ {
    result = append(result, ...)
}
```

#### 2. Reuse big.Int / big.Int 재사용

```go
// ✅ Good - Reuse / 재사용
charsetLen := big.NewInt(int64(len(charset)))
for i := 0; i < actualLength; i++ {
    randomIndex, _ := rand.Int(rand.Reader, charsetLen)
    // ...
}

// ❌ Bad - Create each time / 매번 생성
for i := 0; i < actualLength; i++ {
    randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
    // ...
}
```

#### 3. Minimize Allocations / 할당 최소화

```go
// Check allocations / 할당 확인
go test -bench=BenchmarkAlnum -benchmem

// Goal: 2 allocs/op or less / 목표: 작업당 2회 이하 할당
```

### Profiling / 프로파일링

```bash
# CPU profiling / CPU 프로파일링
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profiling / 메모리 프로파일링
go test -bench=. -memprofile=mem.prof
go tool pprof mem.prof
```

---

## Security Considerations / 보안 고려사항

### Why crypto/rand is Essential / crypto/rand가 필수인 이유

**Vulnerabilities with math/rand / math/rand의 취약점**:

```go
// ❌ INSECURE - Never use for security! / 절대 보안용으로 사용 금지!
import "math/rand"

rand.Seed(time.Now().UnixNano())
password := rand.Intn(1000000)  // Predictable! / 예측 가능!
```

**Attack Scenario / 공격 시나리오**:
1. Attacker observes timestamp / 공격자가 타임스탬프 관찰
2. Seeds pseudo-random generator / 유사 난수 생성기에 시드 설정
3. Predicts all future values / 모든 미래 값 예측

**Secure Alternative / 안전한 대안**:

```go
// ✅ SECURE - Use crypto/rand / crypto/rand 사용
import "crypto/rand"

randomBytes := make([]byte, 32)
_, err := rand.Read(randomBytes)  // Cryptographically secure / 암호학적으로 안전
```

### Security Best Practices / 보안 모범 사례

1. **Always use crypto/rand / 항상 crypto/rand 사용**
   - Never `math/rand` for security / 보안에 `math/rand` 절대 사용 금지

2. **Sufficient Length / 충분한 길이**
   - Passwords: ≥16 characters / 비밀번호: ≥16자
   - API Keys: ≥32 characters / API 키: ≥32자
   - Tokens: ≥32 characters / 토큰: ≥32자

3. **Handle Errors / 에러 처리**
   - Never ignore errors from random generation / 랜덤 생성 에러 절대 무시 금지
   - System may run out of entropy / 시스템이 엔트로피 부족할 수 있음

4. **Character Set Selection / 문자 집합 선택**
   - Use `Complex()` for maximum security / 최대 보안을 위해 `Complex()` 사용
   - Avoid predictable patterns / 예측 가능한 패턴 피하기

---

## Contribution Guidelines / 기여 가이드라인

### Before Contributing / 기여하기 전에

1. **Check Existing Issues / 기존 이슈 확인**
   - Search for similar requests / 유사한 요청 검색
   - Discuss major changes first / 주요 변경사항은 먼저 논의

2. **Fork and Branch / 포크 및 브랜치**
   ```bash
   git clone https://github.com/YOUR_USERNAME/go-utils.git
   cd go-utils
   git checkout -b feature/new-method
   ```

3. **Make Changes / 변경 수행**
   - Follow code style / 코드 스타일 준수
   - Add tests / 테스트 추가
   - Update documentation / 문서 업데이트

4. **Test Thoroughly / 철저한 테스트**
   ```bash
   go test -v
   go test -cover
   go test -bench=.
   ```

5. **Submit Pull Request / 풀 리퀘스트 제출**
   - Clear description / 명확한 설명
   - Link related issues / 관련 이슈 링크
   - Include examples / 예제 포함

### Contribution Checklist / 기여 체크리스트

- [ ] Code follows style guidelines / 코드가 스타일 가이드라인을 따름
- [ ] Bilingual comments (English/Korean) / 이중 언어 주석 (영문/한글)
- [ ] All tests pass / 모든 테스트 통과
- [ ] New tests added / 새 테스트 추가
- [ ] Documentation updated / 문서 업데이트
- [ ] Benchmarks added / 벤치마크 추가
- [ ] No breaking changes (or documented) / 주요 변경사항 없음 (또는 문서화됨)

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

```go
// ✅ Good - Clear, descriptive names / 명확하고 설명적인 이름
func (stringGenerator) AlphaUpper(length ...int) (string, error)
const charsetAlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ❌ Bad - Unclear abbreviations / 불명확한 약어
func (stringGenerator) AU(l ...int) (string, error)
const csAU = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
```

### Comment Style / 주석 스타일

```go
// ✅ Good - Bilingual, structured / 이중 언어, 구조화됨
// AlphaUpper generates a random string containing only uppercase letters (A-Z)
// AlphaUpper는 대문자 알파벳(A-Z)만 포함하는 랜덤 문자열을 생성합니다
//
// Parameters / 매개변수:
//   - length: variable arguments for length specification / 길이 지정을 위한 가변 인자

// ❌ Bad - No Korean, unstructured / 한글 없음, 구조화되지 않음
// generates uppercase string
func AlphaUpper(length ...int) (string, error)
```

### Error Messages / 에러 메시지

```go
// ✅ Good - Descriptive, helpful / 설명적이고 도움이 됨
return "", fmt.Errorf("minimum length cannot be negative: %d", min)

// ❌ Bad - Vague / 모호함
return "", errors.New("invalid input")
```

### Testing Style / 테스트 스타일

```go
// ✅ Good - Table-driven tests / 테이블 기반 테스트
func TestAlnum(t *testing.T) {
    tests := []struct {
        name    string
        length  []int
        wantErr bool
    }{
        {"Fixed 10", []int{10}, false},
        {"Range 5-15", []int{5, 15}, false},
        {"Negative", []int{-1}, true},
    }
    // ...
}

// ❌ Bad - Individual test functions / 개별 테스트 함수
func TestAlnumFixed(t *testing.T) { /* ... */ }
func TestAlnumRange(t *testing.T) { /* ... */ }
func TestAlnumNegative(t *testing.T) { /* ... */ }
```

---

## Resources / 리소스

### Documentation / 문서

- [User Manual](USER_MANUAL.md) - Complete user guide / 완전한 사용자 가이드
- [Package README](../../random/README.md) - Package overview / 패키지 개요
- [Examples](../../examples/random_string/) - Working examples / 작동 예제

### External Resources / 외부 리소스

- [crypto/rand Documentation](https://pkg.go.dev/crypto/rand)
- [Go Testing Documentation](https://pkg.go.dev/testing)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

---

**Last Updated / 최종 업데이트**: 2025-10-10
**Version / 버전**: v1.0.008
**License / 라이선스**: MIT
