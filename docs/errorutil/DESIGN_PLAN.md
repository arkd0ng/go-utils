# errorutil Package Design Plan
# errorutil 패키지 설계 계획서

## 1. Package Overview | 패키지 개요

### 1.1 Purpose | 목적
The `errorutil` package provides comprehensive error handling utilities for Go applications, offering enhanced error creation, wrapping, inspection, and formatting capabilities beyond the standard library.

`errorutil` 패키지는 Go 애플리케이션을 위한 포괄적인 에러 처리 유틸리티를 제공하며, 표준 라이브러리를 넘어서는 향상된 에러 생성, 래핑, 검사 및 포매팅 기능을 제공합니다.

### 1.2 Target Use Cases | 주요 사용 사례
- **Error Wrapping | 에러 래핑**: Add context to errors while preserving the original error chain | 원본 에러 체인을 유지하면서 컨텍스트 추가
- **Error Classification | 에러 분류**: Categorize errors by type (validation, network, database, etc.) | 타입별로 에러 분류 (검증, 네트워크, 데이터베이스 등)
- **Error Codes | 에러 코드**: Associate numeric or string codes with errors for API responses | API 응답을 위해 숫자 또는 문자열 코드를 에러와 연결
- **Stack Traces | 스택 트레이스**: Capture and display stack traces for debugging | 디버깅을 위한 스택 트레이스 캡처 및 표시
- **Error Inspection | 에러 검사**: Extract information from error chains | 에러 체인에서 정보 추출
- **Contextual Errors | 컨텍스트 에러**: Attach structured key-value data to errors | 구조화된 키-값 데이터를 에러에 첨부
- **Sentinel Errors | 센티널 에러**: Define and check for specific error conditions | 특정 에러 조건 정의 및 확인
- **Error Formatting | 에러 포매팅**: Customize error message presentation | 에러 메시지 표현 방식 커스터마이징

### 1.3 Design Principles | 설계 원칙
- **Standard Library Compatibility | 표준 라이브러리 호환성**: Works seamlessly with `errors` and `fmt` packages | `errors` 및 `fmt` 패키지와 원활하게 작동
- **Zero Dependencies | 제로 의존성**: No external dependencies beyond Go standard library | Go 표준 라이브러리 외 외부 의존성 없음
- **Performance | 성능**: Minimal overhead for error creation and wrapping | 에러 생성 및 래핑에 대한 최소 오버헤드
- **Type Safety | 타입 안정성**: Strongly typed error interfaces and type assertions | 강력한 타입의 에러 인터페이스 및 타입 단언
- **Immutability | 불변성**: Error values are immutable after creation | 에러 값은 생성 후 불변
- **Composability | 조합성**: Error utilities can be combined and layered | 에러 유틸리티는 조합 및 계층화 가능

## 2. Core Architecture

### 2.1 Error Types Hierarchy

```
Error Interface (built-in error)
│
├── WrappedError (implements Unwrap)
│   └── Stores: message, cause error
│
├── CodedError (implements Code)
│   └── Stores: message, code (string/int), cause error
│
├── StackError (implements StackTrace)
│   └── Stores: message, stack frames, cause error
│
├── ContextError (implements Context)
│   └── Stores: message, key-value pairs, cause error
│
└── CompositeError (multiple interfaces)
    └── Combines: Code + Stack + Context + Unwrap
```

### 2.2 Key Interfaces

```go
// Core error wrapping interface (compatible with Go 1.13+)
type Unwrapper interface {
    Unwrap() error
}

// Error with associated code
type Coder interface {
    error
    Code() string
}

// Error with numeric code
type NumericCoder interface {
    error
    Code() int
}

// Error with stack trace
type StackTracer interface {
    error
    StackTrace() []Frame
}

// Error with contextual data
type Contexter interface {
    error
    Context() map[string]interface{}
}

// Stack frame information
type Frame struct {
    File     string
    Line     int
    Function string
}
```

## 3. Feature Modules

### 3.1 Error Creation Module

**Functions:**
- `New(message string) error`: Create basic error
- `Newf(format string, args ...interface{}) error`: Create formatted error
- `WithCode(message, code string) error`: Create error with string code
- `WithNumericCode(message string, code int) error`: Create error with numeric code
- `WithStack(message string) error`: Create error with stack trace
- `WithContext(message string, ctx map[string]interface{}) error`: Create error with context

**Example:**
```go
err := errorutil.WithCode("user not found", "USER_NOT_FOUND")
err := errorutil.WithStack("database connection failed")
```

### 3.2 Error Wrapping Module

**Functions:**
- `Wrap(err error, message string) error`: Wrap error with message
- `Wrapf(err error, format string, args ...interface{}) error`: Wrap with formatted message
- `WrapWithCode(err error, message, code string) error`: Wrap with code
- `WrapWithStack(err error, message string) error`: Wrap with stack trace
- `WrapWithContext(err error, message string, ctx map[string]interface{}) error`: Wrap with context

**Example:**
```go
if err := db.Query(); err != nil {
    return errorutil.WrapWithCode(err, "failed to fetch users", "DB_QUERY_ERROR")
}
```

### 3.3 Error Inspection Module

**Functions:**
- `Unwrap(err error) error`: Unwrap one level (standard library)
- `UnwrapAll(err error) []error`: Get all errors in chain
- `Root(err error) error`: Get root cause error
- `HasCode(err error, code string) bool`: Check if error chain has code
- `GetCode(err error) (string, bool)`: Extract code from error chain
- `GetNumericCode(err error) (int, bool)`: Extract numeric code
- `GetStack(err error) ([]Frame, bool)`: Extract stack trace
- `GetContext(err error) (map[string]interface{}, bool)`: Extract context data
- `Contains(err error, target error) bool`: Check if error chain contains target

**Example:**
```go
if code, ok := errorutil.GetCode(err); ok {
    log.Printf("Error code: %s", code)
}
if errorutil.HasCode(err, "DB_CONNECTION_ERROR") {
    // Handle database connection errors
}
```

### 3.4 Error Classification Module

**Pre-defined Error Categories:**
- `ErrValidation`: Validation errors
- `ErrNotFound`: Resource not found errors
- `ErrPermission`: Permission denied errors
- `ErrNetwork`: Network-related errors
- `ErrTimeout`: Timeout errors
- `ErrDatabase`: Database errors
- `ErrInternal`: Internal server errors

**Functions:**
- `IsValidation(err error) bool`: Check if validation error
- `IsNotFound(err error) bool`: Check if not found error
- `IsPermission(err error) bool`: Check if permission error
- `IsNetwork(err error) bool`: Check if network error
- `IsTimeout(err error) bool`: Check if timeout error
- `IsDatabase(err error) bool`: Check if database error
- `IsInternal(err error) bool`: Check if internal error

**Example:**
```go
if errorutil.IsNotFound(err) {
    return http.StatusNotFound, "Resource not found"
}
```

### 3.5 Error Formatting Module

**Functions:**
- `Format(err error, verbose bool) string`: Format error with optional verbosity
- `FormatWithStack(err error) string`: Format with stack trace
- `FormatChain(err error) []string`: Format entire error chain
- `ToJSON(err error) string`: Convert error to JSON format
- `ToMap(err error) map[string]interface{}`: Convert error to map

**Example:**
```go
// Verbose output with stack trace
fmt.Println(errorutil.Format(err, true))

// JSON for API responses
jsonErr := errorutil.ToJSON(err)
```

### 3.6 Error Assertion Module

**Functions:**
- `As(err error, target interface{}) bool`: Type assertion for error chain
- `Is(err, target error) bool`: Check if error matches target
- `Must(err error)`: Panic if error is not nil (for initialization)
- `MustReturn(val T, err error) T`: Return value or panic
- `Assert(condition bool, message string) error`: Create error if condition is false

**Example:**
```go
var codeErr *CodedError
if errorutil.As(err, &codeErr) {
    fmt.Printf("Error code: %s\n", codeErr.Code())
}

config := errorutil.MustReturn(loadConfig())
```

## 4. Data Structures

### 4.1 Wrapped Error
```go
type wrappedError struct {
    msg   string
    cause error
}
```

### 4.2 Coded Error
```go
type codedError struct {
    msg   string
    code  string  // or int for numeric codes
    cause error
}
```

### 4.3 Stack Error
```go
type stackError struct {
    msg    string
    frames []Frame
    cause  error
}

type Frame struct {
    File     string
    Line     int
    Function string
    PC       uintptr
}
```

### 4.4 Context Error
```go
type contextError struct {
    msg     string
    context map[string]interface{}
    cause   error
}
```

### 4.5 Composite Error
```go
type compositeError struct {
    msg     string
    code    string
    frames  []Frame
    context map[string]interface{}
    cause   error
}
```

## 5. API Design

### 5.1 Constructor Functions

| Function | Return Type | Purpose |
|----------|-------------|---------|
| `New(msg string)` | `error` | Basic error creation |
| `Newf(format string, args ...interface{})` | `error` | Formatted error creation |
| `WithCode(msg, code string)` | `error` | Error with string code |
| `WithNumericCode(msg string, code int)` | `error` | Error with int code |
| `WithStack(msg string)` | `error` | Error with stack trace |
| `WithContext(msg string, ctx map[string]interface{})` | `error` | Error with context |
| `Wrap(err error, msg string)` | `error` | Wrap with message |
| `Wrapf(err error, format string, args ...interface{})` | `error` | Wrap with formatted message |

### 5.2 Inspection Functions

| Function | Return Type | Purpose |
|----------|-------------|---------|
| `Unwrap(err error)` | `error` | Unwrap one level |
| `UnwrapAll(err error)` | `[]error` | Get all errors in chain |
| `Root(err error)` | `error` | Get root cause |
| `HasCode(err error, code string)` | `bool` | Check for code |
| `GetCode(err error)` | `(string, bool)` | Extract code |
| `GetStack(err error)` | `([]Frame, bool)` | Extract stack |
| `GetContext(err error)` | `(map[string]interface{}, bool)` | Extract context |
| `Contains(err, target error)` | `bool` | Check containment |

### 5.3 Classification Functions

| Function | Return Type | Purpose |
|----------|-------------|---------|
| `IsValidation(err error)` | `bool` | Check validation error |
| `IsNotFound(err error)` | `bool` | Check not found error |
| `IsPermission(err error)` | `bool` | Check permission error |
| `IsNetwork(err error)` | `bool` | Check network error |
| `IsTimeout(err error)` | `bool` | Check timeout error |
| `IsDatabase(err error)` | `bool` | Check database error |
| `IsInternal(err error)` | `bool` | Check internal error |

### 5.4 Formatting Functions

| Function | Return Type | Purpose |
|----------|-------------|---------|
| `Format(err error, verbose bool)` | `string` | Format error message |
| `FormatWithStack(err error)` | `string` | Format with stack |
| `FormatChain(err error)` | `[]string` | Format error chain |
| `ToJSON(err error)` | `string` | Convert to JSON |
| `ToMap(err error)` | `map[string]interface{}` | Convert to map |

## 6. Error Message Guidelines

### 6.1 Message Format
- **Clear and Concise**: Describe what went wrong
- **Contextual**: Include relevant information (file names, IDs, etc.)
- **Actionable**: Suggest what to do next when possible
- **Lowercase Start**: Follow Go convention (lowercase first letter)
- **No Punctuation**: No trailing period

**Examples:**
```go
// Good
errorutil.New("failed to open configuration file")
errorutil.Wrapf(err, "unable to connect to database %s", dbName)

// Bad
errorutil.New("Error occurred")  // Too vague
errorutil.New("Failed to open configuration file.")  // Capitalized, has period
```

### 6.2 Error Context Guidelines
- Add context at each layer where meaningful
- Don't repeat information from lower layers
- Include operation-specific details

**Example:**
```go
// Controller layer
if err := service.GetUser(id); err != nil {
    return errorutil.WrapWithCode(err, "failed to fetch user profile", "PROFILE_ERROR")
}

// Service layer  
if err := repo.FindByID(id); err != nil {
    return errorutil.Wrapf(err, "user lookup failed for id=%d", id)
}

// Repository layer
if err := db.Query(sql, id); err != nil {
    return errorutil.Wrap(err, "database query execution failed")
}
```

## 7. Performance Considerations

### 7.1 Stack Trace Capture
- **Lazy Capture**: Only capture stack when explicitly requested
- **Depth Limit**: Default to 32 frames, configurable
- **Skip Frames**: Skip errorutil internal frames

### 7.2 Error Chain Traversal
- **Early Exit**: Stop traversal when condition is met
- **Memoization**: Cache inspection results when appropriate
- **Nil Checks**: Fast path for nil errors

### 7.3 Memory Footprint
- **Shared Context**: Reuse context maps where possible
- **String Interning**: Consider interning common error codes
- **Frame Pooling**: Pool frame slices for stack traces

## 8. Testing Strategy

### 8.1 Unit Tests (60%+ coverage target: 80%+)
- Error creation and wrapping
- Error inspection and unwrapping
- Error classification
- Error formatting
- Type assertions (As, Is)
- Edge cases (nil errors, empty chains, etc.)

### 8.2 Benchmark Tests
- Error creation performance
- Wrapping overhead
- Chain traversal performance
- Stack trace capture cost
- Formatting performance

### 8.3 Example Tests
- Real-world error handling scenarios
- Integration with standard library
- Error chain construction
- API response error formatting

## 9. Documentation Requirements

### 9.1 Package Documentation (README.md)
- Quick start guide
- Feature overview
- Installation instructions
- Basic usage examples
- API reference summary
- Best practices
- Migration guide from standard errors

### 9.2 Function Documentation (Bilingual)
- English: Primary documentation
- Korean: Secondary documentation
- Examples for each major function
- Edge case behavior
- Performance notes

### 9.3 Example Code
- `/examples/errorutil/basic/` - Basic error creation and wrapping
- `/examples/errorutil/advanced/` - Advanced features (codes, stack, context)
- `/examples/errorutil/http_handler/` - HTTP API error handling
- `/examples/errorutil/middleware/` - Error handling middleware

## 10. Migration Path

### 10.1 From Standard Library
```go
// Before (standard library)
import "errors"
err := errors.New("something went wrong")
err2 := fmt.Errorf("failed: %w", err)

// After (errorutil)
import "github.com/arkd0ng/go-utils/errorutil"
err := errorutil.New("something went wrong")
err2 := errorutil.Wrap(err, "failed")
```

### 10.2 From pkg/errors (if migrating)
```go
// Before (pkg/errors)
import "github.com/pkg/errors"
err := errors.Wrap(err, "failed")

// After (errorutil)
import "github.com/arkd0ng/go-utils/errorutil"
err := errorutil.Wrap(err, "failed")  // Compatible API
```

## 11. Future Enhancements (Post v1.12.x)

### 11.1 Potential Features
- **Error Aggregation**: Combine multiple errors (like errors.Join in Go 1.20+)
- **Structured Logging Integration**: Direct integration with logging package
- **Error Metrics**: Built-in error rate tracking
- **Error Recovery**: Helpers for panic recovery with error context
- **gRPC Status**: Conversion to/from gRPC status codes
- **HTTP Status**: Automatic HTTP status code mapping

### 11.2 Advanced Error Types
- **Temporary Errors**: Errors that may succeed on retry
- **Retryable Errors**: Errors with retry policies
- **Multi-Errors**: Aggregated errors from parallel operations

## 12. Dependencies

### 12.1 Standard Library Only
- `errors` - Standard error handling
- `fmt` - String formatting
- `runtime` - Stack trace capture
- `encoding/json` - JSON formatting
- `strings` - String manipulation

### 12.2 No External Dependencies
Following go-utils package independence principle, errorutil will not depend on any external packages.

## 13. File Structure

```
errorutil/
├── errorutil.go           # Package documentation and main exports
├── create.go              # Error creation functions
├── wrap.go                # Error wrapping functions
├── inspect.go             # Error inspection functions
├── classify.go            # Error classification (sentinel errors)
├── format.go              # Error formatting functions
├── assert.go              # Error assertion utilities
├── types.go               # Error type definitions
├── stack.go               # Stack trace utilities
├── options.go             # Option pattern for error creation
├── errors.go              # Standard error definitions
├── create_test.go         # Tests for create.go
├── wrap_test.go           # Tests for wrap.go
├── inspect_test.go        # Tests for inspect.go
├── classify_test.go       # Tests for classify.go
├── format_test.go         # Tests for format.go
├── assert_test.go         # Tests for assert.go
├── stack_test.go          # Tests for stack.go
├── errorutil_test.go      # Integration tests and examples
└── README.md              # Package documentation (bilingual)

docs/errorutil/
├── DESIGN_PLAN.md         # This file
└── WORK_PLAN.md           # Development work plan (to be created)

examples/errorutil/
├── basic/                 # Basic usage examples
├── advanced/              # Advanced features examples
├── http_handler/          # HTTP API error handling example
└── middleware/            # Error handling middleware example
```

## 14. Version Plan

- **v1.12.001-010**: Core error creation and wrapping
- **v1.12.011-020**: Error inspection and classification
- **v1.12.021-030**: Error formatting and assertions
- **v1.12.031-040**: Stack traces and context errors
- **v1.12.041-050**: Documentation and examples
- **v1.12.051-060**: Testing and benchmarks
- **v1.12.061-070**: Final review and polish

---

## Summary

The `errorutil` package will provide a comprehensive, production-ready error handling solution for Go applications. It maintains compatibility with the standard library while offering powerful features like error codes, stack traces, contextual data, and flexible formatting. The package follows go-utils design principles: zero dependencies, high performance, and excellent documentation.

**Key Differentiators:**
- ✅ Zero external dependencies
- ✅ Standard library compatible
- ✅ Rich error inspection API
- ✅ Flexible error classification
- ✅ Built-in stack traces
- ✅ Structured error context
- ✅ JSON-ready for APIs
- ✅ Comprehensive testing
- ✅ Bilingual documentation

**Next Steps:**
1. Create WORK_PLAN.md with detailed task breakdown
2. Implement core error types (v1.12.001-005)
3. Implement creation functions (v1.12.006-010)
4. Continue with inspection, classification, and formatting features
5. Create comprehensive examples and documentation
6. Achieve 80%+ test coverage
7. Review and publish v1.12.x

