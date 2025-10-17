# errorutil Package - Developer Guide / 개발자 가이드

**Version / 버전**: v1.12.012
**Package / 패키지**: `github.com/arkd0ng/go-utils/errorutil`
**Go Version / Go 버전**: 1.18+

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#architecture-overview--아키텍처-개요)
2. [Package Structure / 패키지 구조](#package-structure--패키지-구조)
3. [Core Components / 핵심 컴포넌트](#core-components--핵심-컴포넌트)
4. [Design Patterns / 디자인 패턴](#design-patterns--디자인-패턴)
5. [Internal Implementation / 내부 구현](#internal-implementation--내부-구현)
6. [Adding New Features / 새 기능 추가](#adding-new-features--새-기능-추가)
7. [Testing Guide / 테스트 가이드](#testing-guide--테스트-가이드)
8. [Performance / 성능](#performance--성능)
9. [Contributing Guidelines / 기여 가이드라인](#contributing-guidelines--기여-가이드라인)
10. [Code Style / 코드 스타일](#code-style--코드-스타일)

---

## Architecture Overview / 아키텍처 개요

### Design Principles / 설계 원칙

The errorutil package follows these core design principles:

errorutil 패키지는 다음과 같은 핵심 설계 원칙을 따릅니다:

1. **Simplicity / 간결성**: Clean and intuitive API for error handling / 깔끔하고 직관적인 에러 처리 API
2. **Compatibility / 호환성**: Full compatibility with Go 1.13+ error handling (`errors.Is`, `errors.As`) / Go 1.13+ 에러 처리와 완전 호환
3. **Extensibility / 확장성**: Interface-based design for custom error types / 커스텀 에러 타입을 위한 인터페이스 기반 설계
4. **Immutability / 불변성**: Errors are immutable; wrapping creates new errors / 에러는 불변; 래핑은 새 에러 생성
5. **Zero Dependencies / 제로 의존성**: Only uses Go standard library / Go 표준 라이브러리만 사용
6. **Type Safety / 타입 안전성**: Strong typing for error codes and interfaces / 에러 코드와 인터페이스를 위한 강력한 타이핑
7. **Performance / 성능**: Minimal allocations and efficient error chain walking / 최소 할당 및 효율적인 에러 체인 탐색

### High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────────────────┐
│                        errorutil Package                        │
│                  github.com/arkd0ng/go-utils/errorutil          │
└─────────────────────────────────────────────────────────────────┘
                                  │
                ┌─────────────────┴─────────────────┐
                │                                   │
        ┌───────▼────────┐                 ┌───────▼────────┐
        │   Interfaces   │                 │   Functions    │
        │                │                 │   (18 total)   │
        │ - Unwrapper    │                 │                │
        │ - Coder        │                 │ 3 Categories:  │
        │ - NumericCoder │                 │                │
        │ - StackTracer  │                 │ 1. Creation    │
        │ - Contexter    │                 │    (6 funcs)   │
        │                │                 │                │
        └────────────────┘                 │ 2. Wrapping    │
                │                          │    (6 funcs)   │
                │                          │                │
        ┌───────▼────────┐                 │ 3. Inspection  │
        │  Error Types   │                 │    (6 funcs)   │
        │                │                 │                │
        │ - wrappedError │                 └────────────────┘
        │ - codedError   │
        │ - numericCoded │
        │   Error        │
        │ - stackError   │
        │ - contextError │
        │ - composite    │
        │   Error        │
        └────────────────┘
```

### Component Interaction / 컴포넌트 상호작용

```
User Code / 사용자 코드
    ↓
┌───────────────────────────┐
│   Public API Functions    │  ← Error creation & wrapping
│   (18 functions)          │    에러 생성 및 래핑
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Error Type Selection    │  ← Choose appropriate error type
│   (6 internal types)      │    적절한 에러 타입 선택
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Interface Checks        │  ← Coder, NumericCoder, etc.
│   (Type assertions)       │    타입 단언
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Error Chain Walking     │  ← errors.As, errors.Is
│   (Unwrap recursion)      │    Unwrap 재귀
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Go Standard Library     │  ← errors, fmt
│   (errors package)        │    errors 패키지
└───────────────────────────┘
```

---

## Package Structure / 패키지 구조

### File Organization / 파일 구성

```
errorutil/
├── types.go               # Package documentation, types, interfaces
│                          # 패키지 문서, 타입, 인터페이스
├── error.go               # Error creation and wrapping functions
│                          # 에러 생성 및 래핑 함수
├── inspect.go             # Error inspection functions
│                          # 에러 검사 함수
├── stack.go               # Stack trace capture and formatting
│                          # 스택 트레이스 캡처 및 포매팅
├── format.go              # Error formatting utilities
│                          # 에러 포매팅 유틸리티
├── error_test.go          # Tests for error creation/wrapping
│                          # 에러 생성/래핑 테스트
├── inspect_test.go        # Tests for inspection functions
│                          # 검사 함수 테스트
├── stack_test.go          # Tests for stack traces
│                          # 스택 트레이스 테스트
├── format_test.go         # Tests for formatting
│                          # 포매팅 테스트
└── README.md              # Package README
                           # 패키지 README
```

### File Responsibilities / 파일별 책임

| File / 파일 | Purpose / 목적 | Key Components / 주요 컴포넌트 | Lines / 줄 수 |
|-------------|---------------|----------------------------|--------------|
| `types.go` | Package documentation, type definitions / 패키지 문서, 타입 정의 | Interfaces: Unwrapper, Coder, NumericCoder, StackTracer, Contexter<br/>Types: Frame, wrappedError, codedError, numericCodedError, stackError, contextError, compositeError | ~383 |
| `error.go` | Error creation and wrapping functions / 에러 생성 및 래핑 함수 | Functions: New, Newf, WithCode, WithCodef, WithNumericCode, WithNumericCodef, Wrap, Wrapf, WrapWithCode, WrapWithCodef, WrapWithNumericCode, WrapWithNumericCodef | ~380 |
| `inspect.go` | Error inspection functions / 에러 검사 함수 | Functions: HasCode, HasNumericCode, GetCode, GetNumericCode, GetStackTrace, GetContext | ~311 |
| `stack.go` | Stack trace capture and formatting / 스택 트레이스 캡처 및 포매팅 | Functions: captureStack, NewWithStack, WrapWithStack | ~150 |
| `format.go` | Error formatting utilities / 에러 포매팅 유틸리티 | Functions: FormatError, FormatWithStack, FormatSimple | ~120 |
| `*_test.go` | Comprehensive tests / 종합 테스트 | Test functions, benchmarks | ~1,200 total |

**Total Package Size / 전체 패키지 크기**: ~2,544 lines (implementation + tests) / ~2,544줄 (구현 + 테스트)

**Test Coverage / 테스트 커버리지**: 99.2%

---

## Core Components / 핵심 컴포넌트

### 1. Interfaces / 인터페이스

**Location / 위치**: `errorutil/types.go`

The package defines five key interfaces for error capabilities:

패키지는 에러 기능을 위한 5개의 주요 인터페이스를 정의합니다:

#### Unwrapper Interface

```go
// Unwrapper is the interface for errors that wrap other errors.
// This is compatible with the standard library errors.Unwrap function.
// Unwrapper는 다른 에러를 래핑하는 에러를 위한 인터페이스입니다.
// 표준 라이브러리 errors.Unwrap 함수와 호환됩니다.
type Unwrapper interface {
    error
    Unwrap() error
}
```

**Purpose / 목적**: Enables error chain traversal for `errors.Is` and `errors.As` / `errors.Is` 및 `errors.As`를 위한 에러 체인 탐색 가능

**Implementers / 구현자**: All error types (wrappedError, codedError, numericCodedError, stackError, contextError, compositeError) / 모든 에러 타입

#### Coder Interface

```go
// Coder is the interface for errors that have an associated string code.
// Error codes are useful for API responses and error categorization.
// Coder는 연결된 문자열 코드를 가진 에러를 위한 인터페이스입니다.
// 에러 코드는 API 응답 및 에러 분류에 유용합니다.
type Coder interface {
    error
    Code() string
}
```

**Purpose / 목적**: Allows errors to carry string codes for categorization / 분류를 위해 에러가 문자열 코드를 전달하도록 허용

**Implementers / 구현자**: codedError, compositeError / codedError, compositeError

#### NumericCoder Interface

```go
// NumericCoder is the interface for errors that have an associated numeric code.
// Numeric codes are useful for HTTP status codes and error numbers.
// NumericCoder는 연결된 숫자 코드를 가진 에러를 위한 인터페이스입니다.
// 숫자 코드는 HTTP 상태 코드 및 에러 번호에 유용합니다.
type NumericCoder interface {
    error
    Code() int
}
```

**Purpose / 목적**: Allows errors to carry numeric codes (e.g., HTTP status codes) / 에러가 숫자 코드(예: HTTP 상태 코드)를 전달하도록 허용

**Implementers / 구현자**: numericCodedError, compositeError (when numCode != 0) / numericCodedError, compositeError (numCode != 0일 때)

#### StackTracer Interface

```go
// StackTracer is the interface for errors that capture stack traces.
// Stack traces help with debugging by showing where errors originated.
// StackTracer는 스택 트레이스를 캡처하는 에러를 위한 인터페이스입니다.
// 스택 트레이스는 에러가 어디서 발생했는지 보여줌으로써 디버깅에 도움을 줍니다.
type StackTracer interface {
    error
    StackTrace() []Frame
}
```

**Purpose / 목적**: Provides stack traces for debugging / 디버깅을 위한 스택 트레이스 제공

**Implementers / 구현자**: stackError, compositeError (when stack is set) / stackError, compositeError (stack이 설정된 경우)

#### Contexter Interface

```go
// Contexter is the interface for errors that carry structured contextual data.
// Context data provides additional information about the error condition.
// Contexter는 구조화된 컨텍스트 데이터를 전달하는 에러를 위한 인터페이스입니다.
// 컨텍스트 데이터는 에러 조건에 대한 추가 정보를 제공합니다.
type Contexter interface {
    error
    Context() map[string]interface{}
}
```

**Purpose / 목적**: Allows errors to carry structured contextual data / 에러가 구조화된 컨텍스트 데이터를 전달하도록 허용

**Implementers / 구현자**: contextError, compositeError (when ctx is set) / contextError, compositeError (ctx가 설정된 경우)

### 2. Error Types / 에러 타입

**Location / 위치**: `errorutil/types.go`

The package implements six internal error types:

패키지는 6개의 내부 에러 타입을 구현합니다:

#### wrappedError

```go
type wrappedError struct {
    msg   string
    cause error
}
```

**Purpose / 목적**: Basic error wrapping without additional features / 추가 기능 없는 기본 에러 래핑

**Used by / 사용처**: `New()`, `Newf()`, `Wrap()`, `Wrapf()`

**Interfaces Implemented / 구현된 인터페이스**: `error`, `Unwrapper`

#### codedError

```go
type codedError struct {
    msg   string
    code  string
    cause error
}
```

**Purpose / 목적**: Errors with string codes / 문자열 코드가 있는 에러

**Used by / 사용처**: `WithCode()`, `WithCodef()`, `WrapWithCode()`, `WrapWithCodef()`

**Interfaces Implemented / 구현된 인터페이스**: `error`, `Coder`, `Unwrapper`

**Error Format / 에러 형식**: `[CODE] message: cause`

#### numericCodedError

```go
type numericCodedError struct {
    msg   string
    code  int
    cause error
}
```

**Purpose / 목적**: Errors with numeric codes (HTTP status, etc.) / 숫자 코드가 있는 에러 (HTTP 상태 등)

**Used by / 사용처**: `WithNumericCode()`, `WithNumericCodef()`, `WrapWithNumericCode()`, `WrapWithNumericCodef()`

**Interfaces Implemented / 구현된 인터페이스**: `error`, `NumericCoder`, `Unwrapper`

**Error Format / 에러 형식**: `[404] message: cause`

#### stackError

```go
type stackError struct {
    msg   string
    stack []Frame
    cause error
}
```

**Purpose / 목적**: Errors with captured stack traces / 스택 트레이스가 캡처된 에러

**Used by / 사용처**: `NewWithStack()`, `WrapWithStack()` (defined in stack.go)

**Interfaces Implemented / 구현된 인터페이스**: `error`, `StackTracer`, `Unwrapper`

#### contextError

```go
type contextError struct {
    msg   string
    ctx   map[string]interface{}
    cause error
}
```

**Purpose / 목적**: Errors with structured context data / 구조화된 컨텍스트 데이터가 있는 에러

**Used by / 사용처**: `WithContext()`, `WrapWithContext()` (future extension)

**Interfaces Implemented / 구현된 인터페이스**: `error`, `Contexter`, `Unwrapper`

**Note / 참고**: Context is returned as a defensive copy to prevent mutation / 컨텍스트는 변경을 방지하기 위해 방어적 복사본으로 반환됨

#### compositeError

```go
type compositeError struct {
    msg     string
    code    string
    numCode int
    stack   []Frame
    ctx     map[string]interface{}
    cause   error
}
```

**Purpose / 목적**: Combines multiple error features (code + stack + context) / 여러 에러 기능 결합 (코드 + 스택 + 컨텍스트)

**Used by / 사용처**: Complex error scenarios requiring multiple capabilities / 여러 기능이 필요한 복잡한 에러 시나리오

**Interfaces Implemented / 구현된 인터페이스**: All interfaces depending on which fields are set / 설정된 필드에 따라 모든 인터페이스

**Note / 참고**: Most feature-rich error type, but not currently exposed in public API / 가장 기능이 풍부한 에러 타입이지만 현재 공개 API에 노출되지 않음

### 3. Frame Type / Frame 타입

```go
type Frame struct {
    File     string
    Line     int
    Function string
}
```

**Purpose / 목적**: Represents a single stack frame / 단일 스택 프레임 나타냄

**String Format / 문자열 형식**: `file:line function`

**Example / 예제**: `main.go:42 main.processUser`

---

## Design Patterns / 디자인 패턴

### 1. Factory Pattern / 팩토리 패턴

Error creation functions act as factories that instantiate appropriate error types:

에러 생성 함수는 적절한 에러 타입을 인스턴스화하는 팩토리 역할을 합니다:

```go
// Factory function for coded errors
// 코드가 있는 에러를 위한 팩토리 함수
func WithCode(code, message string) error {
    return &codedError{
        msg:  message,
        code: code,
    }
}

// Factory function for numeric coded errors
// 숫자 코드가 있는 에러를 위한 팩토리 함수
func WithNumericCode(code int, message string) error {
    return &numericCodedError{
        msg:  message,
        code: code,
    }
}
```

**Benefits / 이점**:
- Encapsulates error type selection / 에러 타입 선택 캡슐화
- User doesn't need to know internal types / 사용자가 내부 타입을 알 필요 없음
- Easy to extend with new error types / 새 에러 타입으로 확장하기 쉬움

### 2. Decorator Pattern / 데코레이터 패턴

Error wrapping adds additional context to existing errors without modifying them:

에러 래핑은 기존 에러를 수정하지 않고 추가 컨텍스트를 추가합니다:

```go
// Original error
// 원본 에러
err := WithCode("DB_ERROR", "connection timeout")

// Decorated with additional context
// 추가 컨텍스트로 장식됨
wrapped := Wrap(err, "failed to save user")

// Original error is preserved
// 원본 에러 보존됨
// wrapped.Unwrap() returns original err
```

**Benefits / 이점**:
- Preserves error chain / 에러 체인 보존
- Immutable errors / 불변 에러
- Composable error handling / 조합 가능한 에러 처리

### 3. Chain of Responsibility / 책임 연쇄 패턴

Error inspection walks the error chain to find capabilities:

에러 검사는 기능을 찾기 위해 에러 체인을 탐색합니다:

```go
func HasCode(err error, code string) bool {
    if err == nil {
        return false
    }

    // Check current error first
    // 현재 에러를 먼저 확인
    if coder, ok := err.(Coder); ok {
        if coder.Code() == code {
            return true
        }
    }

    // Walk the chain
    // 체인 탐색
    var coder Coder
    if errors.As(err, &coder) {
        return coder.Code() == code
    }

    return false
}
```

**Benefits / 이점**:
- Works through wrapped errors / 래핑된 에러를 통해 작동
- Compatible with standard library / 표준 라이브러리와 호환
- Flexible error handling / 유연한 에러 처리

### 4. Template Method Pattern / 템플릿 메서드 패턴

Error formatting follows a template:

에러 포매팅은 템플릿을 따릅니다:

```go
// Template for coded errors
// 코드가 있는 에러를 위한 템플릿
func (e *codedError) Error() string {
    if e.cause != nil {
        return fmt.Sprintf("[%s] %s: %v", e.code, e.msg, e.cause)
    }
    return fmt.Sprintf("[%s] %s", e.code, e.msg)
}
```

**Format Template / 형식 템플릿**:
- With code and cause: `[CODE] message: cause`
- With code only: `[CODE] message`
- Without code: `message: cause`

---

## Internal Implementation / 내부 구현

### Error Creation Flow / 에러 생성 흐름

```
User calls WithCode("ERR001", "error")
사용자가 WithCode("ERR001", "error") 호출
    ↓
┌────────────────────────────┐
│ WithCode function          │
│ WithCode 함수              │
└────────────────────────────┘
    ↓
┌────────────────────────────┐
│ Create codedError struct   │
│ codedError 구조체 생성     │
│ {                          │
│   msg: "error",            │
│   code: "ERR001",          │
│   cause: nil               │
│ }                          │
└────────────────────────────┘
    ↓
┌────────────────────────────┐
│ Return as error interface  │
│ error 인터페이스로 반환    │
└────────────────────────────┘
    ↓
User receives error
사용자가 error 수신
```

### Error Wrapping Flow / 에러 래핑 흐름

```
User calls Wrap(originalErr, "context")
사용자가 Wrap(originalErr, "context") 호출
    ↓
┌────────────────────────────┐
│ Wrap function              │
│ Wrap 함수                  │
└────────────────────────────┘
    ↓
Check if originalErr is nil
originalErr이 nil인지 확인
    ↓
    ├─ YES: return nil
    │   YES: nil 반환
    │
    └─ NO: continue
        NO: 계속
        ↓
┌────────────────────────────┐
│ Create wrappedError        │
│ wrappedError 생성          │
│ {                          │
│   msg: "context",          │
│   cause: originalErr       │
│ }                          │
└────────────────────────────┘
    ↓
┌────────────────────────────┐
│ Return as error interface  │
│ error 인터페이스로 반환    │
│ (implements Unwrap())      │
│ (Unwrap() 구현)            │
└────────────────────────────┘
```

### Error Inspection Flow / 에러 검사 흐름

```
User calls HasCode(err, "ERR001")
사용자가 HasCode(err, "ERR001") 호출
    ↓
┌────────────────────────────┐
│ HasCode function           │
│ HasCode 함수               │
└────────────────────────────┘
    ↓
Check if err is nil
err이 nil인지 확인
    ↓
    ├─ YES: return false
    │   YES: false 반환
    │
    └─ NO: continue
        NO: 계속
        ↓
┌────────────────────────────┐
│ Type assert to Coder       │
│ Coder로 타입 단언          │
│ if coder, ok := err.(Coder)│
└────────────────────────────┘
    ↓
    ├─ Success: Check code
    │   성공: 코드 확인
    │   ↓
    │   return coder.Code() == code
    │
    └─ Fail: Walk error chain
        실패: 에러 체인 탐색
        ↓
┌────────────────────────────┐
│ errors.As to find Coder    │
│ Coder를 찾기 위해 errors.As│
│ var coder Coder            │
│ if errors.As(err, &coder)  │
└────────────────────────────┘
    ↓
    ├─ Found: return coder.Code() == code
    │   찾음: coder.Code() == code 반환
    │
    └─ Not found: return false
        찾지 못함: false 반환
```

### Error Chain Walking / 에러 체인 탐색

The package relies on Go's `errors.As` for chain walking:

패키지는 체인 탐색을 위해 Go의 `errors.As`를 사용합니다:

```
Error Chain Example / 에러 체인 예제:

err3 := Wrap(err2, "layer 3")
err2 := Wrap(err1, "layer 2")
err1 := WithCode("ERR001", "original")

Chain structure / 체인 구조:
┌──────────────────┐
│ wrappedError     │ ← err3 (outermost)
│ msg: "layer 3"   │    err3 (최외곽)
│ cause: ───────┐  │
└───────────────┼──┘
                ↓
┌──────────────────┐
│ wrappedError     │ ← err2 (middle)
│ msg: "layer 2"   │    err2 (중간)
│ cause: ───────┐  │
└───────────────┼──┘
                ↓
┌──────────────────┐
│ codedError       │ ← err1 (innermost)
│ msg: "original"  │    err1 (최내곽)
│ code: "ERR001"   │
│ cause: nil       │
└──────────────────┘

HasCode(err3, "ERR001") walks:
HasCode(err3, "ERR001") 탐색:
1. Check err3 → not a Coder
   err3 확인 → Coder 아님
2. Unwrap to err2 → not a Coder
   err2로 Unwrap → Coder 아님
3. Unwrap to err1 → is a Coder!
   err1로 Unwrap → Coder임!
4. Check err1.Code() == "ERR001" → true
   err1.Code() == "ERR001" 확인 → true
```

---

## Adding New Features / 새 기능 추가

### Adding a New Error Type / 새 에러 타입 추가

**Step 1**: Define the error type in `types.go`

**1단계**: `types.go`에 에러 타입 정의

```go
// customError is an error with custom capability
// customError는 커스텀 기능이 있는 에러입니다
type customError struct {
    msg    string
    custom CustomData
    cause  error
}

func (e *customError) Error() string {
    if e.cause != nil {
        return fmt.Sprintf("%s (custom: %v): %v", e.msg, e.custom, e.cause)
    }
    return fmt.Sprintf("%s (custom: %v)", e.msg, e.custom)
}

func (e *customError) Unwrap() error {
    return e.cause
}

func (e *customError) CustomData() CustomData {
    return e.custom
}
```

**Step 2**: Define the interface in `types.go`

**2단계**: `types.go`에 인터페이스 정의

```go
// Customizer is the interface for errors with custom data
// Customizer는 커스텀 데이터가 있는 에러를 위한 인터페이스입니다
type Customizer interface {
    error
    CustomData() CustomData
}
```

**Step 3**: Add creation function in `error.go`

**3단계**: `error.go`에 생성 함수 추가

```go
// WithCustom creates an error with custom data
// WithCustom은 커스텀 데이터로 에러를 생성합니다
func WithCustom(message string, data CustomData) error {
    return &customError{
        msg:    message,
        custom: data,
        cause:  nil,
    }
}
```

**Step 4**: Add inspection function in `inspect.go`

**4단계**: `inspect.go`에 검사 함수 추가

```go
// GetCustomData extracts custom data from an error
// GetCustomData는 에러에서 커스텀 데이터를 추출합니다
func GetCustomData(err error) (CustomData, bool) {
    if err == nil {
        return nil, false
    }

    if customizer, ok := err.(Customizer); ok {
        return customizer.CustomData(), true
    }

    var customizer Customizer
    if errors.As(err, &customizer) {
        return customizer.CustomData(), true
    }

    return nil, false
}
```

**Step 5**: Add comprehensive tests

**5단계**: 종합 테스트 추가

```go
// TestWithCustom tests the WithCustom function
// TestWithCustom은 WithCustom 함수를 테스트합니다
func TestWithCustom(t *testing.T) {
    data := CustomData{...}
    err := WithCustom("test error", data)

    // Test error message
    // 에러 메시지 테스트
    if err.Error() != "test error (custom: ...)" {
        t.Errorf("unexpected error message")
    }

    // Test custom data extraction
    // 커스텀 데이터 추출 테스트
    gotData, ok := GetCustomData(err)
    if !ok {
        t.Fatal("GetCustomData failed")
    }
    // Verify gotData...
}
```

### Adding a New Inspection Function / 새 검사 함수 추가

Follow the pattern used by existing inspection functions:

기존 검사 함수가 사용하는 패턴을 따르세요:

```go
// Get{Feature} extracts {feature} from an error or any error in its chain.
// Get{Feature}는 에러 또는 에러 체인의 어떤 에러에서 {feature}를 추출합니다.
func Get{Feature}(err error) ({Type}, bool) {
    if err == nil {
        return {zero value}, false
    }

    // Check the current error first
    // 현재 에러를 먼저 확인
    if {interfacer}, ok := err.({Interface}); ok {
        return {interfacer}.{Method}(), true
    }

    // Walk the error chain
    // 에러 체인을 탐색
    var {interfacer} {Interface}
    if errors.As(err, &{interfacer}) {
        return {interfacer}.{Method}(), true
    }

    return {zero value}, false
}
```

---

## Testing Guide / 테스트 가이드

### Test Structure / 테스트 구조

Each test file follows this structure:

각 테스트 파일은 이 구조를 따릅니다:

```go
package errorutil

import (
    "errors"
    "testing"
)

// Test{FunctionName} tests the {FunctionName} function
// Test{FunctionName}은 {FunctionName} 함수를 테스트합니다
func Test{FunctionName}(t *testing.T) {
    tests := []struct {
        name     string
        // input fields
        // want fields
    }{
        {
            name: "description",
            // test case
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test logic
        })
    }
}
```

### Test Categories / 테스트 카테고리

1. **Creation Tests / 생성 테스트**:
   - Test all error creation functions / 모든 에러 생성 함수 테스트
   - Verify error messages / 에러 메시지 확인
   - Check interface implementation / 인터페이스 구현 확인

2. **Wrapping Tests / 래핑 테스트**:
   - Test wrapping with/without cause / cause가 있거나 없는 래핑 테스트
   - Verify nil handling / nil 처리 확인
   - Check error chain preservation / 에러 체인 보존 확인

3. **Inspection Tests / 검사 테스트**:
   - Test code detection / 코드 감지 테스트
   - Verify chain walking / 체인 탐색 확인
   - Check nil safety / nil 안전성 확인

4. **Integration Tests / 통합 테스트**:
   - Test complex error chains / 복잡한 에러 체인 테스트
   - Verify standard library compatibility / 표준 라이브러리 호환성 확인
   - Check real-world scenarios / 실제 시나리오 확인

### Running Tests / 테스트 실행

```bash
# Run all tests
# 모든 테스트 실행
go test ./errorutil -v

# Run with coverage
# 커버리지와 함께 실행
go test ./errorutil -cover

# Run specific test
# 특정 테스트 실행
go test ./errorutil -run TestHasCode -v

# Generate coverage report
# 커버리지 리포트 생성
go test ./errorutil -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Coverage Requirements / 테스트 커버리지 요구사항

- **Minimum / 최소**: 80%
- **Target / 목표**: 95%+
- **Current / 현재**: 99.2%

All exported functions must have tests covering:

모든 내보낸 함수는 다음을 커버하는 테스트가 있어야 합니다:

- Normal cases / 정상 케이스
- Edge cases (nil, empty, zero) / 엣지 케이스 (nil, 빈, 0)
- Error cases / 에러 케이스
- Integration with other functions / 다른 함수와의 통합

---

## Performance / 성능

### Allocation Benchmarks / 할당 벤치마크

```bash
# Run benchmarks
# 벤치마크 실행
go test ./errorutil -bench=. -benchmem

# Example results
# 예제 결과
BenchmarkNew-8                   5000000    250 ns/op    64 B/op    2 allocs/op
BenchmarkWithCode-8              3000000    380 ns/op    96 B/op    2 allocs/op
BenchmarkWrap-8                  4000000    290 ns/op    80 B/op    2 allocs/op
BenchmarkHasCode-8              20000000     55 ns/op     0 B/op    0 allocs/op
```

### Performance Considerations / 성능 고려사항

1. **Minimal Allocations / 최소 할당**:
   - Error creation: 1-2 allocations / 에러 생성: 1-2 할당
   - Wrapping: 1 allocation / 래핑: 1 할당
   - Inspection: 0 allocations (read-only) / 검사: 0 할당 (읽기 전용)

2. **Efficient Chain Walking / 효율적인 체인 탐색**:
   - Uses `errors.As` for optimized traversal / 최적화된 탐색을 위해 `errors.As` 사용
   - Short-circuits on first match / 첫 번째 일치 시 단락
   - No unnecessary allocations / 불필요한 할당 없음

3. **String Formatting / 문자열 포매팅**:
   - Lazy evaluation (only on Error() call) / 지연 평가 (Error() 호출 시에만)
   - Efficient fmt.Sprintf usage / 효율적인 fmt.Sprintf 사용

---

## Contributing Guidelines / 기여 가이드라인

### Development Workflow / 개발 워크플로우

1. **Before Starting / 시작 전**:
   - Read [DEVELOPMENT_WORKFLOW_GUIDE.md](../DEVELOPMENT_WORKFLOW_GUIDE.md)
   - Read [PACKAGE_DEVELOPMENT_GUIDE.md](../PACKAGE_DEVELOPMENT_GUIDE.md)
   - Bump version in `cfg/app.yaml`

2. **During Development / 개발 중**:
   - Write code following code style guidelines / 코드 스타일 가이드라인 따르기
   - Add bilingual comments (English/Korean) / 이중 언어 주석 추가 (영문/한글)
   - Write comprehensive tests / 종합 테스트 작성
   - Maintain >80% coverage / >80% 커버리지 유지

3. **Before Committing / 커밋 전**:
   - Run all tests: `go test ./errorutil -v` / 모든 테스트 실행
   - Check coverage: `go test ./errorutil -cover` / 커버리지 확인
   - Update CHANGELOG.md / CHANGELOG.md 업데이트
   - Verify bilingual documentation / 이중 언어 문서 확인

4. **Commit / 커밋**:
   - Follow commit message format / 커밋 메시지 형식 따르기
   - Include version in message / 메시지에 버전 포함

### Pull Request Checklist / Pull Request 체크리스트

- [ ] Version bumped in `cfg/app.yaml` / `cfg/app.yaml`에서 버전 증가
- [ ] All tests passing / 모든 테스트 통과
- [ ] Coverage >80% / 커버리지 >80%
- [ ] Documentation updated (bilingual) / 문서 업데이트 (이중 언어)
- [ ] CHANGELOG updated / CHANGELOG 업데이트
- [ ] Examples updated if applicable / 예제 업데이트 (해당하는 경우)
- [ ] Code follows style guidelines / 코드가 스타일 가이드라인 준수

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

1. **Functions / 함수**:
   - Creation: `New`, `With{Feature}` (e.g., `WithCode`)
   - Wrapping: `Wrap`, `WrapWith{Feature}`
   - Inspection: `Has{Feature}`, `Get{Feature}`
   - Use PascalCase for exported / 내보낸 것에 PascalCase 사용

2. **Types / 타입**:
   - Internal error types: `{feature}Error` (e.g., `codedError`)
   - Interfaces: `{Feature}er` (e.g., `Coder`, `Unwrapper`)
   - Use camelCase for unexported / 내보내지 않은 것에 camelCase 사용

3. **Variables / 변수**:
   - Use descriptive names / 설명적인 이름 사용
   - Avoid single letters except in loops / 루프를 제외하고 단일 문자 피하기
   - Error variables: `err`, `cause`, `wrapped`

### Comment Style / 주석 스타일

**All comments must be bilingual (English/Korean):**

**모든 주석은 이중 언어(영문/한글)여야 합니다:**

```go
// FunctionName does something useful.
// Additional details about the function.
//
// FunctionName은 유용한 작업을 수행합니다.
// 함수에 대한 추가 세부정보.
//
// Parameters
// 매개변수:
// - param1: Description
// 설명
//
// Returns
// 반환:
// - type: Description
// 설명
//
// Example
// 예제:
//
//  err := FunctionName(param1)
//  if err != nil {
//      // handle error
//      // 에러 처리
//  }
func FunctionName(param1 Type) ReturnType {
    // Implementation
    // 구현
}
```

### Error Message Format / 에러 메시지 형식

1. **Simple errors / 간단한 에러**: `"message"`
2. **Coded errors / 코드가 있는 에러**: `"[CODE] message"`
3. **Numeric coded / 숫자 코드**: `"[404] message"`
4. **With cause / cause가 있는 경우**: `"message: cause"`
5. **Full format / 전체 형식**: `"[CODE] message: cause"`

### Function Organization / 함수 구성

Within each file, organize functions by:

각 파일 내에서 함수를 다음으로 구성:

1. Type definitions / 타입 정의
2. Creation functions / 생성 함수
3. Wrapping functions / 래핑 함수
4. Inspection functions / 검사 함수
5. Helper functions / 헬퍼 함수

---

## Appendix / 부록

### Interface Compatibility Matrix / 인터페이스 호환성 매트릭스

| Error Type | Unwrapper | Coder | NumericCoder | StackTracer | Contexter |
|------------|-----------|-------|--------------|-------------|-----------|
| wrappedError | ✅ | ❌ | ❌ | ❌ | ❌ |
| codedError | ✅ | ✅ | ❌ | ❌ | ❌ |
| numericCodedError | ✅ | ❌ | ✅ | ❌ | ❌ |
| stackError | ✅ | ❌ | ❌ | ✅ | ❌ |
| contextError | ✅ | ❌ | ❌ | ❌ | ✅ |
| compositeError | ✅ | ✅* | ✅* | ✅* | ✅* |

\* compositeError implements interfaces conditionally based on which fields are set
\* compositeError는 설정된 필드에 따라 조건부로 인터페이스를 구현합니다

### Function Reference Table / 함수 참조 테이블

| Function | Returns | Implements | File |
|----------|---------|------------|------|
| New | wrappedError | Unwrapper | error.go |
| Newf | wrappedError | Unwrapper | error.go |
| WithCode | codedError | Coder, Unwrapper | error.go |
| WithCodef | codedError | Coder, Unwrapper | error.go |
| WithNumericCode | numericCodedError | NumericCoder, Unwrapper | error.go |
| WithNumericCodef | numericCodedError | NumericCoder, Unwrapper | error.go |
| Wrap | wrappedError | Unwrapper | error.go |
| Wrapf | wrappedError | Unwrapper | error.go |
| WrapWithCode | codedError | Coder, Unwrapper | error.go |
| WrapWithCodef | codedError | Coder, Unwrapper | error.go |
| WrapWithNumericCode | numericCodedError | NumericCoder, Unwrapper | error.go |
| WrapWithNumericCodef | numericCodedError | NumericCoder, Unwrapper | error.go |
| HasCode | bool | - | inspect.go |
| HasNumericCode | bool | - | inspect.go |
| GetCode | string, bool | - | inspect.go |
| GetNumericCode | int, bool | - | inspect.go |
| GetStackTrace | []Frame, bool | - | inspect.go |
| GetContext | map[string]interface{}, bool | - | inspect.go |
| Root | error | - | inspect.go |
| UnwrapAll | []error | - | inspect.go |
| Contains | bool | - | inspect.go |

---

**Last Updated / 최종 업데이트**: 2025-10-17
**Version / 버전**: v1.12.012
**Package / 패키지**: errorutil
