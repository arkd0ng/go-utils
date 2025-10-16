# errorutil Package Design Plan / errorutil 패키지 설계 계획서

## 1. Package Overview / 패키지 개요

### 1.1 Purpose / 목적

The `errorutil` package provides comprehensive error handling utilities for Go applications, offering enhanced error creation, wrapping, inspection, and formatting capabilities beyond the standard library.

`errorutil` 패키지는 Go 애플리케이션을 위한 포괄적인 에러 처리 유틸리티를 제공하며, 표준 라이브러리를 넘어서는 향상된 에러 생성, 래핑, 검사 및 포매팅 기능을 제공합니다.

### 1.2 Target Use Cases / 주요 사용 사례

- **Error Wrapping / 에러 래핑**: Add context to errors while preserving the original error chain / 원본 에러 체인을 유지하면서 컨텍스트 추가
- **Error Classification / 에러 분류**: Categorize errors by type (validation, network, database, etc.) / 타입별로 에러 분류 (검증, 네트워크, 데이터베이스 등)
- **Error Codes / 에러 코드**: Associate numeric or string codes with errors for API responses / API 응답을 위해 숫자 또는 문자열 코드를 에러와 연결
- **Stack Traces / 스택 트레이스**: Capture and display stack traces for debugging / 디버깅을 위한 스택 트레이스 캡처 및 표시
- **Error Inspection / 에러 검사**: Extract information from error chains / 에러 체인에서 정보 추출
- **Contextual Errors / 컨텍스트 에러**: Attach structured key-value data to errors / 구조화된 키-값 데이터를 에러에 첨부
- **Sentinel Errors / 센티널 에러**: Define and check for specific error conditions / 특정 에러 조건 정의 및 확인
- **Error Formatting / 에러 포매팅**: Customize error message presentation / 에러 메시지 표현 방식 커스터마이징

### 1.3 Design Principles / 설계 원칙

- **Standard Library Compatibility / 표준 라이브러리 호환성**: Works seamlessly with `errors` and `fmt` packages / `errors` 및 `fmt` 패키지와 원활하게 작동
- **Zero Dependencies / 제로 의존성**: No external dependencies beyond Go standard library / Go 표준 라이브러리 외 외부 의존성 없음
- **Performance / 성능**: Minimal overhead for error creation and wrapping / 에러 생성 및 래핑에 대한 최소 오버헤드
- **Type Safety / 타입 안정성**: Strongly typed error interfaces and type assertions / 강력한 타입의 에러 인터페이스 및 타입 단언
- **Immutability / 불변성**: Error values are immutable after creation / 에러 값은 생성 후 불변
- **Composability / 조합성**: Error utilities can be combined and layered / 에러 유틸리티는 조합 및 계층화 가능

---

## 2. Core Architecture / 핵심 아키텍처

### 2.1 Error Types Hierarchy / 에러 타입 계층 구조

```
Error Interface (built-in error) / 에러 인터페이스 (내장 error)
│
├── WrappedError (implements Unwrap) / 래핑된 에러 (Unwrap 구현)
│   └── Stores: message, cause error / 저장: 메시지, 원인 에러
│
├── CodedError (implements Code) / 코드 에러 (Code 구현)
│   └── Stores: message, code (string/int), cause error / 저장: 메시지, 코드, 원인 에러
│
├── StackError (implements StackTrace) / 스택 에러 (StackTrace 구현)
│   └── Stores: message, stack frames, cause error / 저장: 메시지, 스택 프레임, 원인 에러
│
├── ContextError (implements Context) / 컨텍스트 에러 (Context 구현)
│   └── Stores: message, key-value pairs, cause error / 저장: 메시지, 키-값 쌍, 원인 에러
│
└── CompositeError (multiple interfaces) / 복합 에러 (다중 인터페이스)
    └── Combines: Code + Stack + Context + Unwrap / 결합: 코드 + 스택 + 컨텍스트 + Unwrap
```

### 2.2 Key Interfaces / 주요 인터페이스

```go
// Core error wrapping interface (compatible with Go 1.13+)
// 핵심 에러 래핑 인터페이스 (Go 1.13+ 호환)
type Unwrapper interface {
    Unwrap() error
}

// Error with associated code / 코드가 연결된 에러
type Coder interface {
    error
    Code() string
}

// Error with numeric code / 숫자 코드를 가진 에러
type NumericCoder interface {
    error
    Code() int
}

// Error with stack trace / 스택 트레이스를 가진 에러
type StackTracer interface {
    error
    StackTrace() []Frame
}

// Error with contextual data / 컨텍스트 데이터를 가진 에러
type Contexter interface {
    error
    Context() map[string]interface{}
}

// Stack frame information / 스택 프레임 정보
type Frame struct {
    File     string  // 파일 경로
    Line     int     // 라인 번호
    Function string  // 함수명
}
```

---

## 3. Feature Modules / 기능 모듈

### 3.1 Error Creation Module / 에러 생성 모듈

**Functions / 함수:**

- `New(message string) error`: Create basic error / 기본 에러 생성
- `Newf(format string, args ...interface{}) error`: Create formatted error / 포맷된 에러 생성
- `WithCode(message, code string) error`: Create error with string code / 문자열 코드를 가진 에러 생성
- `WithNumericCode(message string, code int) error`: Create error with numeric code / 숫자 코드를 가진 에러 생성
- `WithStack(message string) error`: Create error with stack trace / 스택 트레이스를 가진 에러 생성
- `WithContext(message string, ctx map[string]interface{}) error`: Create error with context / 컨텍스트를 가진 에러 생성

**Example / 예제:**

```go
// Create error with code / 코드를 가진 에러 생성
err := errorutil.WithCode("user not found", "USER_NOT_FOUND")

// Create error with stack trace / 스택 트레이스를 가진 에러 생성
err := errorutil.WithStack("database connection failed")
```

### 3.2 Error Wrapping Module / 에러 래핑 모듈

**Functions / 함수:**

- `Wrap(err error, message string) error`: Wrap error with message / 메시지로 에러 래핑
- `Wrapf(err error, format string, args ...interface{}) error`: Wrap with formatted message / 포맷된 메시지로 래핑
- `WrapWithCode(err error, message, code string) error`: Wrap with code / 코드와 함께 래핑
- `WrapWithStack(err error, message string) error`: Wrap with stack trace / 스택 트레이스와 함께 래핑
- `WrapWithContext(err error, message string, ctx map[string]interface{}) error`: Wrap with context / 컨텍스트와 함께 래핑

**Example / 예제:**

```go
// Wrap error with code / 코드와 함께 에러 래핑
if err := db.Query(); err != nil {
    return errorutil.WrapWithCode(err, "failed to fetch users", "DB_QUERY_ERROR")
}
```

### 3.3 Error Inspection Module / 에러 검사 모듈

**Functions / 함수:**

- `Unwrap(err error) error`: Unwrap one level (standard library) / 한 단계 언래핑 (표준 라이브러리)
- `UnwrapAll(err error) []error`: Get all errors in chain / 체인의 모든 에러 가져오기
- `Root(err error) error`: Get root cause error / 근본 원인 에러 가져오기
- `HasCode(err error, code string) bool`: Check if error chain has code / 에러 체인에 코드가 있는지 확인
- `GetCode(err error) (string, bool)`: Extract code from error chain / 에러 체인에서 코드 추출
- `GetNumericCode(err error) (int, bool)`: Extract numeric code / 숫자 코드 추출
- `GetStack(err error) ([]Frame, bool)`: Extract stack trace / 스택 트레이스 추출
- `GetContext(err error) (map[string]interface{}, bool)`: Extract context data / 컨텍스트 데이터 추출
- `Contains(err error, target error) bool`: Check if error chain contains target / 에러 체인에 대상이 포함되어 있는지 확인

**Example / 예제:**

```go
// Extract error code / 에러 코드 추출
if code, ok := errorutil.GetCode(err); ok {
    log.Printf("Error code: %s", code)
}

// Check for specific code / 특정 코드 확인
if errorutil.HasCode(err, "DB_CONNECTION_ERROR") {
    // Handle database connection errors / 데이터베이스 연결 에러 처리
}
```

### 3.4 Error Classification Module / 에러 분류 모듈

**Pre-defined Error Categories / 미리 정의된 에러 카테고리:**

- `ErrValidation`: Validation errors / 검증 에러
- `ErrNotFound`: Resource not found errors / 리소스를 찾을 수 없음 에러
- `ErrPermission`: Permission denied errors / 권한 거부 에러
- `ErrNetwork`: Network-related errors / 네트워크 관련 에러
- `ErrTimeout`: Timeout errors / 타임아웃 에러
- `ErrDatabase`: Database errors / 데이터베이스 에러
- `ErrInternal`: Internal server errors / 내부 서버 에러

**Functions / 함수:**

- `IsValidation(err error) bool`: Check if validation error / 검증 에러인지 확인
- `IsNotFound(err error) bool`: Check if not found error / 찾을 수 없음 에러인지 확인
- `IsPermission(err error) bool`: Check if permission error / 권한 에러인지 확인
- `IsNetwork(err error) bool`: Check if network error / 네트워크 에러인지 확인
- `IsTimeout(err error) bool`: Check if timeout error / 타임아웃 에러인지 확인
- `IsDatabase(err error) bool`: Check if database error / 데이터베이스 에러인지 확인
- `IsInternal(err error) bool`: Check if internal error / 내부 에러인지 확인

**Example / 예제:**

```go
// Check error type and handle accordingly / 에러 타입 확인 및 처리
if errorutil.IsNotFound(err) {
    return http.StatusNotFound, "Resource not found"
}
```

### 3.5 Error Formatting Module / 에러 포매팅 모듈

**Functions / 함수:**

- `Format(err error, verbose bool) string`: Format error with optional verbosity / 선택적 상세 정보로 에러 포맷
- `FormatWithStack(err error) string`: Format with stack trace / 스택 트레이스와 함께 포맷
- `FormatChain(err error) []string`: Format entire error chain / 전체 에러 체인 포맷
- `ToJSON(err error) string`: Convert error to JSON format / 에러를 JSON 형식으로 변환
- `ToMap(err error) map[string]interface{}`: Convert error to map / 에러를 맵으로 변환

**Example / 예제:**

```go
// Verbose output with stack trace / 스택 트레이스와 함께 상세 출력
fmt.Println(errorutil.Format(err, true))

// JSON for API responses / API 응답용 JSON
jsonErr := errorutil.ToJSON(err)
```

### 3.6 Error Assertion Module / 에러 단언 모듈

**Functions / 함수:**

- `As(err error, target interface{}) bool`: Type assertion for error chain / 에러 체인에 대한 타입 단언
- `Is(err, target error) bool`: Check if error matches target / 에러가 대상과 일치하는지 확인
- `Must(err error)`: Panic if error is not nil (for initialization) / 에러가 nil이 아니면 패닉 (초기화용)
- `MustReturn(val T, err error) T`: Return value or panic / 값 반환 또는 패닉
- `Assert(condition bool, message string) error`: Create error if condition is false / 조건이 거짓이면 에러 생성

**Example / 예제:**

```go
// Type assertion / 타입 단언
var codeErr *CodedError
if errorutil.As(err, &codeErr) {
    fmt.Printf("Error code: %s\n", codeErr.Code())
}

// Must pattern for initialization / 초기화를 위한 Must 패턴
config := errorutil.MustReturn(loadConfig())
```

---

## 4. Data Structures / 데이터 구조

### 4.1 Wrapped Error / 래핑된 에러

```go
type wrappedError struct {
    msg   string  // Error message / 에러 메시지
    cause error   // Underlying error / 기저 에러
}
```

### 4.2 Coded Error / 코드 에러

```go
type codedError struct {
    msg   string  // Error message / 에러 메시지
    code  string  // Error code (or int for numeric codes) / 에러 코드
    cause error   // Underlying error / 기저 에러
}
```

### 4.3 Stack Error / 스택 에러

```go
type stackError struct {
    msg    string   // Error message / 에러 메시지
    frames []Frame  // Stack frames / 스택 프레임
    cause  error    // Underlying error / 기저 에러
}

type Frame struct {
    File     string   // File path / 파일 경로
    Line     int      // Line number / 라인 번호
    Function string   // Function name / 함수명
    PC       uintptr  // Program counter / 프로그램 카운터
}
```

### 4.4 Context Error / 컨텍스트 에러

```go
type contextError struct {
    msg     string                 // Error message / 에러 메시지
    context map[string]interface{} // Contextual data / 컨텍스트 데이터
    cause   error                  // Underlying error / 기저 에러
}
```

### 4.5 Composite Error / 복합 에러

```go
type compositeError struct {
    msg     string                 // Error message / 에러 메시지
    code    string                 // Error code / 에러 코드
    frames  []Frame                // Stack frames / 스택 프레임
    context map[string]interface{} // Contextual data / 컨텍스트 데이터
    cause   error                  // Underlying error / 기저 에러
}
```

---

## 5. API Design / API 설계

### 5.1 Constructor Functions / 생성자 함수

| Function / 함수 | Return Type / 반환 타입 | Purpose / 목적 |
|----------|-------------|---------|
| `New(msg string)` | `error` | Basic error creation / 기본 에러 생성 |
| `Newf(format string, args ...interface{})` | `error` | Formatted error creation / 포맷된 에러 생성 |
| `WithCode(msg, code string)` | `error` | Error with string code / 문자열 코드를 가진 에러 |
| `WithNumericCode(msg string, code int)` | `error` | Error with int code / 정수 코드를 가진 에러 |
| `WithStack(msg string)` | `error` | Error with stack trace / 스택 트레이스를 가진 에러 |
| `WithContext(msg string, ctx map[string]interface{})` | `error` | Error with context / 컨텍스트를 가진 에러 |
| `Wrap(err error, msg string)` | `error` | Wrap with message / 메시지로 래핑 |
| `Wrapf(err error, format string, args ...interface{})` | `error` | Wrap with formatted message / 포맷된 메시지로 래핑 |

### 5.2 Inspection Functions / 검사 함수

| Function / 함수 | Return Type / 반환 타입 | Purpose / 목적 |
|----------|-------------|---------|
| `Unwrap(err error)` | `error` | Unwrap one level / 한 단계 언래핑 |
| `UnwrapAll(err error)` | `[]error` | Get all errors in chain / 체인의 모든 에러 |
| `Root(err error)` | `error` | Get root cause / 근본 원인 가져오기 |
| `HasCode(err error, code string)` | `bool` | Check for code / 코드 확인 |
| `GetCode(err error)` | `(string, bool)` | Extract code / 코드 추출 |
| `GetStack(err error)` | `([]Frame, bool)` | Extract stack / 스택 추출 |
| `GetContext(err error)` | `(map[string]interface{}, bool)` | Extract context / 컨텍스트 추출 |
| `Contains(err, target error)` | `bool` | Check containment / 포함 여부 확인 |

### 5.3 Classification Functions / 분류 함수

| Function / 함수 | Return Type / 반환 타입 | Purpose / 목적 |
|----------|-------------|---------|
| `IsValidation(err error)` | `bool` | Check validation error / 검증 에러 확인 |
| `IsNotFound(err error)` | `bool` | Check not found error / 찾을 수 없음 에러 확인 |
| `IsPermission(err error)` | `bool` | Check permission error / 권한 에러 확인 |
| `IsNetwork(err error)` | `bool` | Check network error / 네트워크 에러 확인 |
| `IsTimeout(err error)` | `bool` | Check timeout error / 타임아웃 에러 확인 |
| `IsDatabase(err error)` | `bool` | Check database error / 데이터베이스 에러 확인 |
| `IsInternal(err error)` | `bool` | Check internal error / 내부 에러 확인 |

### 5.4 Formatting Functions / 포매팅 함수

| Function / 함수 | Return Type / 반환 타입 | Purpose / 목적 |
|----------|-------------|---------|
| `Format(err error, verbose bool)` | `string` | Format error message / 에러 메시지 포맷 |
| `FormatWithStack(err error)` | `string` | Format with stack / 스택과 함께 포맷 |
| `FormatChain(err error)` | `[]string` | Format error chain / 에러 체인 포맷 |
| `ToJSON(err error)` | `string` | Convert to JSON / JSON으로 변환 |
| `ToMap(err error)` | `map[string]interface{}` | Convert to map / 맵으로 변환 |

---

## 6. Error Message Guidelines / 에러 메시지 가이드라인

### 6.1 Message Format / 메시지 형식

- **Clear and Concise / 명확하고 간결**: Describe what went wrong / 무엇이 잘못되었는지 설명
- **Contextual / 컨텍스트**: Include relevant information (file names, IDs, etc.) / 관련 정보 포함 (파일명, ID 등)
- **Actionable / 실행 가능**: Suggest what to do next when possible / 가능하면 다음에 할 일 제안
- **Lowercase Start / 소문자 시작**: Follow Go convention (lowercase first letter) / Go 규칙 준수 (첫 글자 소문자)
- **No Punctuation / 구두점 없음**: No trailing period / 마침표 없음

**Examples / 예제:**

```go
// Good / 좋은 예
errorutil.New("failed to open configuration file")
errorutil.Wrapf(err, "unable to connect to database %s", dbName)

// Bad / 나쁜 예
errorutil.New("Error occurred")  // Too vague / 너무 모호함
errorutil.New("Failed to open configuration file.")  // Capitalized, has period / 대문자 시작, 마침표 있음
```

### 6.2 Error Context Guidelines / 에러 컨텍스트 가이드라인

- Add context at each layer where meaningful / 의미 있는 각 계층에서 컨텍스트 추가
- Don't repeat information from lower layers / 하위 계층의 정보 반복하지 않기
- Include operation-specific details / 작업별 세부사항 포함

**Example / 예제:**

```go
// Controller layer / 컨트롤러 계층
if err := service.GetUser(id); err != nil {
    return errorutil.WrapWithCode(err, "failed to fetch user profile", "PROFILE_ERROR")
}

// Service layer / 서비스 계층
if err := repo.FindByID(id); err != nil {
    return errorutil.Wrapf(err, "user lookup failed for id=%d", id)
}

// Repository layer / 저장소 계층
if err := db.Query(sql, id); err != nil {
    return errorutil.Wrap(err, "database query execution failed")
}
```

---

## 7. Performance Considerations / 성능 고려사항

### 7.1 Stack Trace Capture / 스택 트레이스 캡처

- **Lazy Capture / 지연 캡처**: Only capture stack when explicitly requested / 명시적으로 요청할 때만 스택 캡처
- **Depth Limit / 깊이 제한**: Default to 32 frames, configurable / 기본 32 프레임, 설정 가능
- **Skip Frames / 프레임 건너뛰기**: Skip errorutil internal frames / errorutil 내부 프레임 건너뛰기

### 7.2 Error Chain Traversal / 에러 체인 순회

- **Early Exit / 조기 종료**: Stop traversal when condition is met / 조건이 충족되면 순회 중지
- **Memoization / 메모이제이션**: Cache inspection results when appropriate / 적절한 경우 검사 결과 캐시
- **Nil Checks / Nil 확인**: Fast path for nil errors / nil 에러에 대한 빠른 경로

### 7.3 Memory Footprint / 메모리 사용량

- **Shared Context / 공유 컨텍스트**: Reuse context maps where possible / 가능한 경우 컨텍스트 맵 재사용
- **String Interning / 문자열 인터닝**: Consider interning common error codes / 일반적인 에러 코드 인터닝 고려
- **Frame Pooling / 프레임 풀링**: Pool frame slices for stack traces / 스택 트레이스용 프레임 슬라이스 풀링

---

## 8. Testing Strategy / 테스트 전략

### 8.1 Unit Tests / 단위 테스트

**Coverage target: 80%+ / 커버리지 목표: 80% 이상**

- Error creation and wrapping / 에러 생성 및 래핑
- Error inspection and unwrapping / 에러 검사 및 언래핑
- Error classification / 에러 분류
- Error formatting / 에러 포매팅
- Type assertions (As, Is) / 타입 단언
- Edge cases (nil errors, empty chains, etc.) / 엣지 케이스 (nil 에러, 빈 체인 등)

### 8.2 Benchmark Tests / 벤치마크 테스트

- Error creation performance / 에러 생성 성능
- Wrapping overhead / 래핑 오버헤드
- Chain traversal performance / 체인 순회 성능
- Stack trace capture cost / 스택 트레이스 캡처 비용
- Formatting performance / 포매팅 성능

### 8.3 Example Tests / 예제 테스트

- Real-world error handling scenarios / 실제 에러 처리 시나리오
- Integration with standard library / 표준 라이브러리와의 통합
- Error chain construction / 에러 체인 구성
- API response error formatting / API 응답 에러 포매팅

---

## 9. Documentation Requirements / 문서 요구사항

### 9.1 Package Documentation / 패키지 문서

**README.md (bilingual / 이중 언어):**
- Quick start guide / 빠른 시작 가이드
- Feature overview / 기능 개요
- Installation instructions / 설치 지침
- Basic usage examples / 기본 사용 예제
- API reference summary / API 참조 요약
- Best practices / 모범 사례
- Migration guide from standard errors / 표준 에러에서의 마이그레이션 가이드

### 9.2 Function Documentation / 함수 문서

**All function comments must be bilingual / 모든 함수 주석은 이중 언어:**

```go
// New creates a new error with the given message.
// It returns a basic error that implements the error interface.
// New는 주어진 메시지로 새로운 에러를 생성합니다.
// error 인터페이스를 구현하는 기본 에러를 반환합니다.
func New(message string) error {
    // Implementation / 구현
}
```

### 9.3 Example Code / 예제 코드

**Example directories / 예제 디렉토리:**
- `/examples/errorutil/basic/` - Basic error creation and wrapping / 기본 에러 생성 및 래핑
- `/examples/errorutil/advanced/` - Advanced features (codes, stack, context) / 고급 기능
- `/examples/errorutil/http_handler/` - HTTP API error handling / HTTP API 에러 처리
- `/examples/errorutil/middleware/` - Error handling middleware / 에러 처리 미들웨어

---

## 10. Migration Path / 마이그레이션 경로

### 10.1 From Standard Library / 표준 라이브러리에서

```go
// Before (standard library) / 이전 (표준 라이브러리)
import "errors"
err := errors.New("something went wrong")
err2 := fmt.Errorf("failed: %w", err)

// After (errorutil) / 이후 (errorutil)
import "github.com/arkd0ng/go-utils/errorutil"
err := errorutil.New("something went wrong")
err2 := errorutil.Wrap(err, "failed")
```

### 10.2 From pkg/errors / pkg/errors에서

```go
// Before (pkg/errors) / 이전 (pkg/errors)
import "github.com/pkg/errors"
err := errors.Wrap(err, "failed")

// After (errorutil) / 이후 (errorutil)
import "github.com/arkd0ng/go-utils/errorutil"
err := errorutil.Wrap(err, "failed")  // Compatible API / 호환 가능한 API
```

---

## 11. Future Enhancements / 향후 개선사항

**(Post v1.12.x / v1.12.x 이후)**

### 11.1 Potential Features / 잠재적 기능

- **Error Aggregation / 에러 집계**: Combine multiple errors (like errors.Join in Go 1.20+) / 다중 에러 결합
- **Structured Logging Integration / 구조화된 로깅 통합**: Direct integration with logging package / logging 패키지와 직접 통합
- **Error Metrics / 에러 메트릭**: Built-in error rate tracking / 내장 에러 발생률 추적
- **Error Recovery / 에러 복구**: Helpers for panic recovery with error context / 에러 컨텍스트와 함께 패닉 복구 헬퍼
- **gRPC Status / gRPC 상태**: Conversion to/from gRPC status codes / gRPC 상태 코드 간 변환
- **HTTP Status / HTTP 상태**: Automatic HTTP status code mapping / 자동 HTTP 상태 코드 매핑

### 11.2 Advanced Error Types / 고급 에러 타입

- **Temporary Errors / 임시 에러**: Errors that may succeed on retry / 재시도 시 성공할 수 있는 에러
- **Retryable Errors / 재시도 가능 에러**: Errors with retry policies / 재시도 정책을 가진 에러
- **Multi-Errors / 다중 에러**: Aggregated errors from parallel operations / 병렬 작업의 집계된 에러

---

## 12. Dependencies / 의존성

### 12.1 Standard Library Only / 표준 라이브러리만

- `errors` - Standard error handling / 표준 에러 처리
- `fmt` - String formatting / 문자열 포매팅
- `runtime` - Stack trace capture / 스택 트레이스 캡처
- `encoding/json` - JSON formatting / JSON 포매팅
- `strings` - String manipulation / 문자열 조작

### 12.2 No External Dependencies / 외부 의존성 없음

Following go-utils package independence principle, errorutil will not depend on any external packages.

go-utils 패키지 독립성 원칙에 따라 errorutil은 외부 패키지에 의존하지 않습니다.

---

## 13. File Structure / 파일 구조

```
errorutil/
├── errorutil.go           # Package documentation and main exports / 패키지 문서 및 주요 내보내기
├── create.go              # Error creation functions / 에러 생성 함수
├── wrap.go                # Error wrapping functions / 에러 래핑 함수
├── inspect.go             # Error inspection functions / 에러 검사 함수
├── classify.go            # Error classification / 에러 분류
├── format.go              # Error formatting functions / 에러 포매팅 함수
├── assert.go              # Error assertion utilities / 에러 단언 유틸리티
├── types.go               # Error type definitions / 에러 타입 정의
├── stack.go               # Stack trace utilities / 스택 트레이스 유틸리티
├── options.go             # Option pattern / 옵션 패턴
├── errors.go              # Standard error definitions / 표준 에러 정의
├── create_test.go         # Tests for create.go / create.go 테스트
├── wrap_test.go           # Tests for wrap.go / wrap.go 테스트
├── inspect_test.go        # Tests for inspect.go / inspect.go 테스트
├── classify_test.go       # Tests for classify.go / classify.go 테스트
├── format_test.go         # Tests for format.go / format.go 테스트
├── assert_test.go         # Tests for assert.go / assert.go 테스트
├── stack_test.go          # Tests for stack.go / stack.go 테스트
├── errorutil_test.go      # Integration tests and examples / 통합 테스트 및 예제
└── README.md              # Package documentation (bilingual) / 패키지 문서 (이중 언어)

docs/errorutil/
├── DESIGN_PLAN.md         # This file / 이 파일
└── WORK_PLAN.md           # Development work plan / 개발 작업 계획 (to be created / 생성 예정)

examples/errorutil/
├── basic/                 # Basic usage examples / 기본 사용 예제
├── advanced/              # Advanced features examples / 고급 기능 예제
├── http_handler/          # HTTP API error handling example / HTTP API 에러 처리 예제
└── middleware/            # Error handling middleware example / 에러 처리 미들웨어 예제
```

---

## 14. Version Plan / 버전 계획

- **v1.12.001-010**: Core error creation and wrapping / 핵심 에러 생성 및 래핑
- **v1.12.011-020**: Error inspection and classification / 에러 검사 및 분류
- **v1.12.021-030**: Error formatting and assertions / 에러 포매팅 및 단언
- **v1.12.031-040**: Stack traces and context errors / 스택 트레이스 및 컨텍스트 에러
- **v1.12.041-050**: Documentation and examples / 문서 및 예제
- **v1.12.051-060**: Testing and benchmarks / 테스트 및 벤치마크
- **v1.12.061-070**: Final review and polish / 최종 검토 및 다듬기

---

## Summary / 요약

The `errorutil` package will provide a comprehensive, production-ready error handling solution for Go applications. It maintains compatibility with the standard library while offering powerful features like error codes, stack traces, contextual data, and flexible formatting. The package follows go-utils design principles: zero dependencies, high performance, and excellent documentation.

`errorutil` 패키지는 Go 애플리케이션을 위한 포괄적이고 프로덕션 준비가 완료된 에러 처리 솔루션을 제공합니다. 에러 코드, 스택 트레이스, 컨텍스트 데이터 및 유연한 포매팅과 같은 강력한 기능을 제공하면서 표준 라이브러리와의 호환성을 유지합니다. 이 패키지는 go-utils 설계 원칙을 따릅니다: 제로 의존성, 높은 성능, 뛰어난 문서화.

**Key Differentiators / 주요 차별화 요소:**
- ✅ Zero external dependencies / 외부 의존성 없음
- ✅ Standard library compatible / 표준 라이브러리 호환
- ✅ Rich error inspection API / 풍부한 에러 검사 API
- ✅ Flexible error classification / 유연한 에러 분류
- ✅ Built-in stack traces / 내장 스택 트레이스
- ✅ Structured error context / 구조화된 에러 컨텍스트
- ✅ JSON-ready for APIs / API용 JSON 준비
- ✅ Comprehensive testing / 포괄적인 테스트
- ✅ Bilingual documentation / 이중 언어 문서

**Next Steps / 다음 단계:**
1. Create WORK_PLAN.md with detailed task breakdown / 상세한 작업 분류가 포함된 WORK_PLAN.md 생성
2. Implement core error types (v1.12.001-005) / 핵심 에러 타입 구현
3. Implement creation functions (v1.12.006-010) / 생성 함수 구현
4. Continue with inspection, classification, and formatting features / 검사, 분류 및 포매팅 기능 계속 진행
5. Create comprehensive examples and documentation / 포괄적인 예제 및 문서 생성
6. Achieve 80%+ test coverage / 80% 이상 테스트 커버리지 달성
7. Review and publish v1.12.x / v1.12.x 검토 및 게시
