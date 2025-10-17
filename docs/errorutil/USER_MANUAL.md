# errorutil User Manual / errorutil 사용자 매뉴얼

**Version**: v1.12.011
**Package**: github.com/arkd0ng/go-utils/errorutil
**Go Version**: 1.18+

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Core Concepts / 핵심 개념](#core-concepts--핵심-개념)
5. [Error Creation / 에러 생성](#error-creation--에러-생성)
6. [Error Wrapping / 에러 래핑](#error-wrapping--에러-래핑)
7. [Error Inspection / 에러 검사](#error-inspection--에러-검사)
8. [Advanced Usage / 고급 사용법](#advanced-usage--고급-사용법)
9. [Best Practices / 모범 사례](#best-practices--모범-사례)
10. [Common Patterns / 일반적인 패턴](#common-patterns--일반적인-패턴)
11. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)
12. [API Reference / API 참조](#api-reference--api-참조)

---

## Introduction / 소개

### What is errorutil? / errorutil이란?

The `errorutil` package provides comprehensive error handling utilities for Go applications. It extends the standard library's error handling capabilities while maintaining full compatibility with `errors.Is` and `errors.As`.

`errorutil` 패키지는 Go 애플리케이션을 위한 포괄적인 에러 처리 유틸리티를 제공합니다. 표준 라이브러리의 에러 처리 기능을 확장하면서 `errors.Is` 및 `errors.As`와 완전한 호환성을 유지합니다.

### Why use errorutil? / errorutil을 사용하는 이유는?

**Problems with Standard Library / 표준 라이브러리의 문제점:**

```go
// Standard library - verbose and limited
// 표준 라이브러리 - 장황하고 제한적
err := errors.New("validation failed")
// No error code / 에러 코드 없음
// No structured context / 구조화된 컨텍스트 없음
// Manual wrapping required / 수동 래핑 필요
```

**errorutil Solution / errorutil 해결책:**

```go
// errorutil - concise and feature-rich
// errorutil - 간결하고 기능이 풍부함
err := errorutil.WithCode("VALIDATION_ERROR", "validation failed")
// Has error code / 에러 코드 있음
// Easy to categorize / 분류하기 쉬움
// Simple wrapping / 간단한 래핑
```

### Key Features / 주요 기능

✅ **Error Codes** - String and numeric codes for categorization
✅ **에러 코드** - 분류를 위한 문자열 및 숫자 코드

✅ **Error Wrapping** - Add context while preserving error chains
✅ **에러 래핑** - 에러 체인을 보존하면서 컨텍스트 추가

✅ **Error Inspection** - Check codes and extract information
✅ **에러 검사** - 코드 확인 및 정보 추출

✅ **Go 1.13+ Compatible** - Works with `errors.Is`, `errors.As`
✅ **Go 1.13+ 호환** - `errors.Is`, `errors.As`와 함께 작동

✅ **Stack Traces** - Capture call stacks for debugging
✅ **스택 트레이스** - 디버깅을 위한 호출 스택 캡처

✅ **Contextual Data** - Attach structured data to errors
✅ **컨텍스트 데이터** - 에러에 구조화된 데이터 첨부

---

## Installation / 설치

### Requirements / 요구사항

- Go 1.18 or higher / Go 1.18 이상
- No external dependencies / 외부 의존성 없음

### Install Package / 패키지 설치

```bash
go get github.com/arkd0ng/go-utils/errorutil
```

### Import in Your Code / 코드에서 임포트

```go
import "github.com/arkd0ng/go-utils/errorutil"
```

### Verify Installation / 설치 확인

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/errorutil"
)

func main() {
    err := errorutil.New("test error")
    fmt.Println(err) // Output: test error
}
```

---

## Quick Start / 빠른 시작

### 1. Create a Simple Error / 간단한 에러 생성

```go
// Create error with message
// 메시지로 에러 생성
err := errorutil.New("something went wrong")

// Create formatted error
// 포맷된 에러 생성
err := errorutil.Newf("failed to process user %d", 123)
```

### 2. Create Error with Code / 코드와 함께 에러 생성

```go
// String code
// 문자열 코드
err := errorutil.WithCode("ERR001", "invalid input")

// Numeric code (HTTP status)
// 숫자 코드 (HTTP 상태)
err := errorutil.WithNumericCode(404, "user not found")
```

### 3. Wrap an Error / 에러 래핑

```go
// Original error
// 원본 에러
originalErr := errorutil.WithCode("DB_ERROR", "connection failed")

// Wrap with additional context
// 추가 컨텍스트와 함께 래핑
wrappedErr := errorutil.Wrap(originalErr, "failed to save user")

fmt.Println(wrappedErr)
// Output: failed to save user: [DB_ERROR] connection failed
```

### 4. Check Error Code / 에러 코드 확인

```go
err := errorutil.WithCode("VALIDATION_ERROR", "invalid email")

// Check if error has specific code
// 에러가 특정 코드를 가지는지 확인
if errorutil.HasCode(err, "VALIDATION_ERROR") {
    fmt.Println("Validation error occurred")
    // Handle validation error
    // 검증 에러 처리
}
```

### 5. Complete Example / 완전한 예제

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/errorutil"
)

func validateEmail(email string) error {
    if email == "" {
        return errorutil.WithCode("VALIDATION_ERROR", "email is required")
    }
    return nil
}

func saveUser(email string) error {
    if err := validateEmail(email); err != nil {
        return errorutil.Wrap(err, "failed to save user")
    }
    // Save user logic...
    // 사용자 저장 로직...
    return nil
}

func main() {
    err := saveUser("")
    if err != nil {
        // Check error code
        // 에러 코드 확인
        if errorutil.HasCode(err, "VALIDATION_ERROR") {
            fmt.Println("Validation failed")
            // Return 400 Bad Request
            // 400 Bad Request 반환
        }
        fmt.Println(err)
        // Output: failed to save user: [VALIDATION_ERROR] email is required
    }
}
```

---

## Core Concepts / 핵심 개념

### Error Chains / 에러 체인

Go 1.13+ supports error chains through the `Unwrap()` method. The errorutil package fully supports this concept.

Go 1.13+는 `Unwrap()` 메서드를 통해 에러 체인을 지원합니다. errorutil 패키지는 이 개념을 완전히 지원합니다.

```go
// Error chain example
// 에러 체인 예제
err1 := errorutil.WithCode("DB_ERROR", "connection timeout")
err2 := errorutil.Wrap(err1, "failed to fetch user")
err3 := errorutil.Wrap(err2, "API call failed")

// err3 wraps err2, which wraps err1
// err3은 err2를 래핑하고, err2는 err1을 래핑합니다
fmt.Println(err3)
// Output: API call failed: failed to fetch user: [DB_ERROR] connection timeout
```

### Error Codes / 에러 코드

Error codes help categorize and identify errors programmatically.

에러 코드는 프로그래밍 방식으로 에러를 분류하고 식별하는 데 도움이 됩니다.

**String Codes / 문자열 코드:**
- Use for application-specific errors / 애플리케이션 특정 에러에 사용
- Examples: `"VALIDATION_ERROR"`, `"AUTH_FAILED"`, `"DB_ERROR"`

**Numeric Codes / 숫자 코드:**
- Use for HTTP status codes / HTTP 상태 코드에 사용
- Examples: `404`, `500`, `401`

```go
// String code for business logic
// 비즈니스 로직을 위한 문자열 코드
err1 := errorutil.WithCode("PAYMENT_FAILED", "insufficient funds")

// Numeric code for HTTP responses
// HTTP 응답을 위한 숫자 코드
err2 := errorutil.WithNumericCode(402, "payment required")
```

### Error Interfaces / 에러 인터페이스

The package defines several interfaces for error capabilities:

패키지는 에러 기능을 위한 여러 인터페이스를 정의합니다:

```go
// Unwrapper - supports error chains
// Unwrapper - 에러 체인 지원
type Unwrapper interface {
    error
    Unwrap() error
}

// Coder - has string code
// Coder - 문자열 코드 있음
type Coder interface {
    error
    Code() string
}

// NumericCoder - has numeric code
// NumericCoder - 숫자 코드 있음
type NumericCoder interface {
    error
    Code() int
}

// StackTracer - captures stack trace
// StackTracer - 스택 트레이스 캡처
type StackTracer interface {
    error
    StackTrace() []Frame
}

// Contexter - carries contextual data
// Contexter - 컨텍스트 데이터 전달
type Contexter interface {
    error
    Context() map[string]interface{}
}
```

---

## Error Creation / 에러 생성

### Basic Error Creation / 기본 에러 생성

#### New - Create Simple Error / 간단한 에러 생성

```go
// Create error with message
// 메시지로 에러 생성
err := errorutil.New("file not found")
fmt.Println(err) // Output: file not found
```

#### Newf - Create Formatted Error / 포맷된 에러 생성

```go
// Create error with formatted message
// 포맷된 메시지로 에러 생성
userID := 123
err := errorutil.Newf("user %d not found", userID)
fmt.Println(err) // Output: user 123 not found
```

### Errors with String Codes / 문자열 코드가 있는 에러

#### WithCode - Create Error with Code / 코드와 함께 에러 생성

```go
// Create error with string code
// 문자열 코드와 함께 에러 생성
err := errorutil.WithCode("VALIDATION_ERROR", "invalid email format")
fmt.Println(err) // Output: [VALIDATION_ERROR] invalid email format

// Check the code
// 코드 확인
if errorutil.HasCode(err, "VALIDATION_ERROR") {
    fmt.Println("Handle validation error")
}
```

#### WithCodef - Create Formatted Error with Code / 코드와 포맷된 에러 생성

```go
// Create formatted error with code
// 코드와 포맷된 에러 생성
field := "email"
err := errorutil.WithCodef("VALIDATION_ERROR", "field %s is required", field)
fmt.Println(err) // Output: [VALIDATION_ERROR] field email is required
```

### Errors with Numeric Codes / 숫자 코드가 있는 에러

#### WithNumericCode - Create Error with Numeric Code / 숫자 코드와 에러 생성

```go
// Create error with HTTP status code
// HTTP 상태 코드와 함께 에러 생성
err := errorutil.WithNumericCode(404, "resource not found")
fmt.Println(err) // Output: [404] resource not found

// Use for HTTP responses
// HTTP 응답에 사용
if errorutil.HasNumericCode(err, 404) {
    // Return 404 Not Found
    // 404 Not Found 반환
}
```

#### WithNumericCodef - Create Formatted Error with Numeric Code / 숫자 코드와 포맷된 에러 생성

```go
// Create formatted error with numeric code
// 숫자 코드와 포맷된 에러 생성
resourceID := "user-123"
err := errorutil.WithNumericCodef(404, "resource %s not found", resourceID)
fmt.Println(err) // Output: [404] resource user-123 not found
```

### When to Use Each Type / 각 타입을 사용하는 경우

| Function | Use Case | Example |
|----------|----------|---------|
| `New` / `Newf` | Internal errors, no categorization needed<br/>내부 에러, 분류 불필요 | File I/O errors<br/>파일 I/O 에러 |
| `WithCode` / `WithCodef` | Business logic errors<br/>비즈니스 로직 에러 | Validation, authorization<br/>검증, 권한 부여 |
| `WithNumericCode` / `WithNumericCodef` | HTTP/API errors<br/>HTTP/API 에러 | REST API responses<br/>REST API 응답 |

---

## Error Wrapping / 에러 래핑

### Why Wrap Errors? / 에러를 래핑하는 이유는?

Error wrapping adds context to errors as they propagate up the call stack, helping you understand where and why errors occurred.

에러 래핑은 호출 스택을 따라 전파될 때 에러에 컨텍스트를 추가하여 에러가 어디서 왜 발생했는지 이해하는 데 도움이 됩니다.

```go
// Without wrapping - loses context
// 래핑 없이 - 컨텍스트 손실
func fetchUser(id int) error {
    return errors.New("connection timeout")
    // Where did this happen? What was the user ID?
    // 어디서 발생했나? 사용자 ID는 무엇인가?
}

// With wrapping - preserves context
// 래핑 사용 - 컨텍스트 보존
func fetchUser(id int) error {
    err := db.Connect()
    if err != nil {
        return errorutil.Wrapf(err, "failed to fetch user %d", id)
        // Clear context: what failed and why
        // 명확한 컨텍스트: 무엇이 왜 실패했는지
    }
    return nil
}
```

### Basic Wrapping / 기본 래핑

#### Wrap - Wrap Error with Message / 메시지로 에러 래핑

```go
// Original error
// 원본 에러
dbErr := errorutil.WithCode("DB_ERROR", "connection timeout")

// Wrap with context
// 컨텍스트와 함께 래핑
err := errorutil.Wrap(dbErr, "failed to save user")

fmt.Println(err)
// Output: failed to save user: [DB_ERROR] connection timeout

// Original error is preserved
// 원본 에러 보존됨
if errorutil.HasCode(err, "DB_ERROR") {
    fmt.Println("Database error detected")
}
```

#### Wrapf - Wrap Error with Formatted Message / 포맷된 메시지로 에러 래핑

```go
// Wrap with formatted context
// 포맷된 컨텍스트로 래핑
userID := 123
err := errorutil.Wrapf(dbErr, "failed to save user %d", userID)

fmt.Println(err)
// Output: failed to save user 123: [DB_ERROR] connection timeout
```

### Wrapping with Codes / 코드와 함께 래핑

#### WrapWithCode - Add Code While Wrapping / 래핑하면서 코드 추가

```go
// Original error (no code)
// 원본 에러 (코드 없음)
ioErr := errors.New("file not found")

// Wrap and add code
// 래핑하고 코드 추가
err := errorutil.WrapWithCode(ioErr, "FILE_ERROR", "failed to read config")

fmt.Println(err)
// Output: [FILE_ERROR] failed to read config: file not found

// Can now check by code
// 이제 코드로 확인 가능
if errorutil.HasCode(err, "FILE_ERROR") {
    fmt.Println("File error occurred")
}
```

#### WrapWithCodef - Formatted Wrap with Code / 코드와 포맷된 래핑

```go
// Wrap with code and formatted message
// 코드와 포맷된 메시지로 래핑
filename := "config.yaml"
err := errorutil.WrapWithCodef(ioErr, "FILE_ERROR",
    "failed to read %s", filename)

fmt.Println(err)
// Output: [FILE_ERROR] failed to read config.yaml: file not found
```

### Wrapping with Numeric Codes / 숫자 코드와 함께 래핑

#### WrapWithNumericCode - Add Numeric Code While Wrapping / 래핑하면서 숫자 코드 추가

```go
// Original error
// 원본 에러
dbErr := errors.New("user not found")

// Wrap with HTTP status code
// HTTP 상태 코드와 함께 래핑
err := errorutil.WrapWithNumericCode(dbErr, 404, "failed to fetch user")

fmt.Println(err)
// Output: [404] failed to fetch user: user not found

// Use in HTTP handler
// HTTP 핸들러에서 사용
if errorutil.HasNumericCode(err, 404) {
    http.Error(w, err.Error(), 404)
}
```

#### WrapWithNumericCodef - Formatted Wrap with Numeric Code / 숫자 코드와 포맷된 래핑

```go
// Wrap with numeric code and formatted message
// 숫자 코드와 포맷된 메시지로 래핑
userID := 123
err := errorutil.WrapWithNumericCodef(dbErr, 404,
    "failed to fetch user %d", userID)

fmt.Println(err)
// Output: [404] failed to fetch user 123: user not found
```

### Multiple Levels of Wrapping / 다중 레벨 래핑

You can wrap errors multiple times to build a detailed error chain:

에러를 여러 번 래핑하여 상세한 에러 체인을 만들 수 있습니다:

```go
// Layer 1: Database
// 레이어 1: 데이터베이스
func dbQuery() error {
    return errorutil.WithCode("DB_TIMEOUT", "connection timeout")
}

// Layer 2: Repository
// 레이어 2: 저장소
func getUser(id int) error {
    err := dbQuery()
    if err != nil {
        return errorutil.Wrapf(err, "failed to query user %d", id)
    }
    return nil
}

// Layer 3: Service
// 레이어 3: 서비스
func fetchUserProfile(id int) error {
    err := getUser(id)
    if err != nil {
        return errorutil.Wrap(err, "failed to fetch user profile")
    }
    return nil
}

// Layer 4: HTTP Handler
// 레이어 4: HTTP 핸들러
func handleGetUser(w http.ResponseWriter, r *http.Request) {
    err := fetchUserProfile(123)
    if err != nil {
        // Full error chain preserved
        // 전체 에러 체인 보존됨
        fmt.Println(err)
        // Output: failed to fetch user profile: failed to query user 123: [DB_TIMEOUT] connection timeout

        // Original code still accessible
        // 원본 코드 여전히 접근 가능
        if errorutil.HasCode(err, "DB_TIMEOUT") {
            http.Error(w, "Database timeout", 503)
        }
    }
}
```

---

## Error Inspection / 에러 검사

### Checking Error Codes / 에러 코드 확인

#### HasCode - Check String Code / 문자열 코드 확인

```go
err := errorutil.WithCode("AUTH_FAILED", "invalid credentials")

// Check if error has specific code
// 에러가 특정 코드를 가지는지 확인
if errorutil.HasCode(err, "AUTH_FAILED") {
    fmt.Println("Authentication failed")
    // Return 401 Unauthorized
    // 401 Unauthorized 반환
}

// Works through error chains
// 에러 체인을 통해 작동
wrappedErr := errorutil.Wrap(err, "failed to login")
if errorutil.HasCode(wrappedErr, "AUTH_FAILED") {
    fmt.Println("Still detects AUTH_FAILED through wrapping")
}
```

#### HasNumericCode - Check Numeric Code / 숫자 코드 확인

```go
err := errorutil.WithNumericCode(404, "user not found")

// Check if error has specific numeric code
// 에러가 특정 숫자 코드를 가지는지 확인
if errorutil.HasNumericCode(err, 404) {
    fmt.Println("Not found error")
    // Return 404 Not Found
    // 404 Not Found 반환
}

// Also works through error chains
// 에러 체인을 통해서도 작동
wrappedErr := errorutil.Wrap(err, "failed to fetch user")
if errorutil.HasNumericCode(wrappedErr, 404) {
    fmt.Println("404 still detected")
}
```

### Extracting Error Information / 에러 정보 추출

#### GetCode - Extract String Code / 문자열 코드 추출

```go
err := errorutil.WithCode("VALIDATION_ERROR", "invalid input")

// Get the error code
// 에러 코드 가져오기
code, ok := errorutil.GetCode(err)
if ok {
    fmt.Printf("Error code: %s\n", code)
    // Output: Error code: VALIDATION_ERROR
}

// Returns empty string and false if no code
// 코드가 없으면 빈 문자열과 false 반환
plainErr := errors.New("plain error")
code, ok = errorutil.GetCode(plainErr)
fmt.Printf("Code: %q, Found: %v\n", code, ok)
// Output: Code: "", Found: false
```

#### GetNumericCode - Extract Numeric Code / 숫자 코드 추출

```go
err := errorutil.WithNumericCode(500, "internal server error")

// Get the numeric code
// 숫자 코드 가져오기
code, ok := errorutil.GetNumericCode(err)
if ok {
    fmt.Printf("HTTP status: %d\n", code)
    // Output: HTTP status: 500

    // Use for HTTP response
    // HTTP 응답에 사용
    http.Error(w, err.Error(), code)
}

// Returns 0 and false if no code
// 코드가 없으면 0과 false 반환
plainErr := errors.New("plain error")
code, ok = errorutil.GetNumericCode(plainErr)
fmt.Printf("Code: %d, Found: %v\n", code, ok)
// Output: Code: 0, Found: false
```

### Error Code Decision Flow / 에러 코드 결정 흐름

```go
func handleError(err error) {
    // Try to get string code first
    // 먼저 문자열 코드 가져오기 시도
    if code, ok := errorutil.GetCode(err); ok {
        switch code {
        case "VALIDATION_ERROR":
            // Handle validation error
            // 검증 에러 처리
            logValidationError(err)
        case "AUTH_FAILED":
            // Handle auth error
            // 인증 에러 처리
            logAuthError(err)
        default:
            logGenericError(code, err)
        }
        return
    }

    // Try numeric code for HTTP errors
    // HTTP 에러를 위해 숫자 코드 시도
    if code, ok := errorutil.GetNumericCode(err); ok {
        if code >= 400 && code < 500 {
            // Client error
            // 클라이언트 에러
            handleClientError(code, err)
        } else if code >= 500 {
            // Server error
            // 서버 에러
            handleServerError(code, err)
        }
        return
    }

    // No code found, handle as generic error
    // 코드를 찾지 못함, 일반 에러로 처리
    handleGenericError(err)
}
```

---

## Advanced Usage / 고급 사용법

### Using with Standard Library / 표준 라이브러리와 함께 사용

errorutil errors are fully compatible with Go's standard `errors` package:

errorutil 에러는 Go의 표준 `errors` 패키지와 완전히 호환됩니다:

```go
import (
    "errors"
    "github.com/arkd0ng/go-utils/errorutil"
)

// errors.Is works
// errors.Is 작동
var ErrNotFound = errorutil.WithCode("NOT_FOUND", "resource not found")

err := errorutil.Wrap(ErrNotFound, "failed to fetch user")

if errors.Is(err, ErrNotFound) {
    fmt.Println("Not found error detected")
}

// errors.As works
// errors.As 작동
var coder errorutil.Coder
if errors.As(err, &coder) {
    fmt.Printf("Error code: %s\n", coder.Code())
    // Output: Error code: NOT_FOUND
}
```

### HTTP API Error Handling / HTTP API 에러 처리

Complete example of using errorutil in HTTP handlers:

HTTP 핸들러에서 errorutil을 사용하는 완전한 예제:

```go
package main

import (
    "encoding/json"
    "net/http"
    "github.com/arkd0ng/go-utils/errorutil"
)

// Error response structure
// 에러 응답 구조
type ErrorResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

// HTTP error handler
// HTTP 에러 핸들러
func handleAPIError(w http.ResponseWriter, err error) {
    var statusCode int
    var errorCode string

    // Get error code and status
    // 에러 코드와 상태 가져오기
    if code, ok := errorutil.GetNumericCode(err); ok {
        statusCode = code
    } else {
        statusCode = 500
    }

    if code, ok := errorutil.GetCode(err); ok {
        errorCode = code
    } else {
        errorCode = "INTERNAL_ERROR"
    }

    // Send JSON response
    // JSON 응답 전송
    w.Header().Set("Content-Type", "application/json")
    w.WriteStatus(statusCode)
    json.NewEncoder(w).Encode(ErrorResponse{
        Code:    errorCode,
        Message: err.Error(),
    })
}

// Example handler
// 예제 핸들러
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("id")

    if userID == "" {
        err := errorutil.WithNumericCode(400, "user ID required")
        handleAPIError(w, err)
        return
    }

    user, err := fetchUser(userID)
    if err != nil {
        // Error already has appropriate code
        // 에러에 이미 적절한 코드 있음
        handleAPIError(w, err)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func fetchUser(id string) (*User, error) {
    user, err := db.GetUser(id)
    if err != nil {
        if isNotFoundError(err) {
            return nil, errorutil.WrapWithNumericCode(err, 404,
                "user not found")
        }
        return nil, errorutil.WrapWithNumericCode(err, 500,
            "failed to fetch user")
    }
    return user, nil
}
```

### Custom Error Types / 커스텀 에러 타입

You can create your own error types that implement errorutil interfaces:

errorutil 인터페이스를 구현하는 자체 에러 타입을 만들 수 있습니다:

```go
// Custom error type with code
// 코드가 있는 커스텀 에러 타입
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error in field %s: %s", e.Field, e.Message)
}

// Implement Coder interface
// Coder 인터페이스 구현
func (e *ValidationError) Code() string {
    return "VALIDATION_ERROR"
}

// Now works with errorutil functions
// 이제 errorutil 함수와 함께 작동
func validateInput(email string) error {
    if email == "" {
        return &ValidationError{
            Field:   "email",
            Message: "email is required",
        }
    }
    return nil
}

func main() {
    err := validateInput("")

    // errorutil.HasCode works with custom type
    // errorutil.HasCode가 커스텀 타입과 작동
    if errorutil.HasCode(err, "VALIDATION_ERROR") {
        fmt.Println("Validation failed")
    }

    // Can still wrap
    // 여전히 래핑 가능
    wrapped := errorutil.Wrap(err, "failed to process input")
    if errorutil.HasCode(wrapped, "VALIDATION_ERROR") {
        fmt.Println("Still detected through wrapping")
    }
}
```

### Error Classification System / 에러 분류 시스템

Build a comprehensive error classification system:

포괄적인 에러 분류 시스템 구축:

```go
// Define error categories
// 에러 카테고리 정의
const (
    // Client errors (4xx)
    // 클라이언트 에러 (4xx)
    ErrCodeValidation   = "VALIDATION_ERROR"
    ErrCodeAuth         = "AUTH_ERROR"
    ErrCodeNotFound     = "NOT_FOUND"
    ErrCodeConflict     = "CONFLICT"

    // Server errors (5xx)
    // 서버 에러 (5xx)
    ErrCodeDatabase     = "DATABASE_ERROR"
    ErrCodeExternal     = "EXTERNAL_SERVICE_ERROR"
    ErrCodeInternal     = "INTERNAL_ERROR"
)

// Error factory functions
// 에러 팩토리 함수
func NewValidationError(message string) error {
    return errorutil.WithNumericCode(400, message)
}

func NewAuthError(message string) error {
    return errorutil.WithCode(ErrCodeAuth, message)
}

func NewNotFoundError(resource string) error {
    return errorutil.WithNumericCodef(404, "%s not found", resource)
}

func NewDatabaseError(err error) error {
    return errorutil.WrapWithCode(err, ErrCodeDatabase,
        "database operation failed")
}

// Centralized error handler
// 중앙 집중식 에러 핸들러
func classifyError(err error) (httpStatus int, logLevel string) {
    // Client errors
    // 클라이언트 에러
    if errorutil.HasCode(err, ErrCodeValidation) {
        return 400, "INFO"
    }
    if errorutil.HasCode(err, ErrCodeAuth) {
        return 401, "WARN"
    }
    if errorutil.HasCode(err, ErrCodeNotFound) {
        return 404, "INFO"
    }

    // Server errors
    // 서버 에러
    if errorutil.HasCode(err, ErrCodeDatabase) {
        return 503, "ERROR"
    }
    if errorutil.HasCode(err, ErrCodeExternal) {
        return 502, "ERROR"
    }

    // Default
    // 기본값
    return 500, "ERROR"
}
```

---

## Best Practices / 모범 사례

### 1. Use Codes for Categorization / 분류를 위해 코드 사용

**✅ Good / 좋음:**
```go
// Easy to categorize and handle
// 분류하고 처리하기 쉬움
err := errorutil.WithCode("VALIDATION_ERROR", "invalid email")

if errorutil.HasCode(err, "VALIDATION_ERROR") {
    // Handle validation errors
    // 검증 에러 처리
}
```

**❌ Bad / 나쁨:**
```go
// Hard to categorize
// 분류하기 어려움
err := errors.New("invalid email")

// Must parse error message (fragile)
// 에러 메시지를 파싱해야 함 (취약함)
if strings.Contains(err.Error(), "invalid") {
    // Fragile and error-prone
    // 취약하고 오류가 발생하기 쉬움
}
```

### 2. Wrap Errors at Layer Boundaries / 레이어 경계에서 에러 래핑

**✅ Good / 좋음:**
```go
// Each layer adds context
// 각 레이어가 컨텍스트 추가
func (s *UserService) GetUser(id int) (*User, error) {
    user, err := s.repo.FindByID(id)
    if err != nil {
        // Add service-level context
        // 서비스 레벨 컨텍스트 추가
        return nil, errorutil.Wrapf(err, "failed to get user %d", id)
    }
    return user, nil
}

func (r *UserRepository) FindByID(id int) (*User, error) {
    user, err := r.db.Query(id)
    if err != nil {
        // Add repository-level context
        // 저장소 레벨 컨텍스트 추가
        return nil, errorutil.Wrap(err, "failed to query database")
    }
    return user, nil
}
```

**❌ Bad / 나쁨:**
```go
// No context added
// 컨텍스트 추가 없음
func (s *UserService) GetUser(id int) (*User, error) {
    return s.repo.FindByID(id) // Just passes error through
}
```

### 3. Check Errors Before Wrapping / 래핑 전에 에러 확인

**✅ Good / 좋음:**
```go
user, err := fetchUser(123)
if err != nil {
    // Only wrap when error occurs
    // 에러가 발생했을 때만 래핑
    return errorutil.Wrap(err, "failed to process user")
}
```

**❌ Bad / 나쁨:**
```go
user, err := fetchUser(123)
// Wrapping nil error
// nil 에러 래핑
return user, errorutil.Wrap(err, "failed to process user")
```

### 4. Use Numeric Codes for HTTP / HTTP에는 숫자 코드 사용

**✅ Good / 좋음:**
```go
// Clear HTTP status mapping
// 명확한 HTTP 상태 매핑
err := errorutil.WithNumericCode(404, "user not found")

if code, ok := errorutil.GetNumericCode(err); ok {
    http.Error(w, err.Error(), code)
}
```

**❌ Bad / 나쁨:**
```go
// Have to map manually
// 수동으로 매핑해야 함
err := errorutil.WithCode("NOT_FOUND", "user not found")

// Manual mapping required
// 수동 매핑 필요
var statusCode int
if errorutil.HasCode(err, "NOT_FOUND") {
    statusCode = 404
} else {
    statusCode = 500
}
http.Error(w, err.Error(), statusCode)
```

### 5. Provide Meaningful Context / 의미 있는 컨텍스트 제공

**✅ Good / 좋음:**
```go
// Clear what failed and why
// 무엇이 왜 실패했는지 명확함
return errorutil.Wrapf(err, "failed to save user %d to database", userID)
```

**❌ Bad / 나쁨:**
```go
// Vague, no useful information
// 모호함, 유용한 정보 없음
return errorutil.Wrap(err, "error occurred")
```

### 6. Don't Overuse Wrapping / 래핑을 과도하게 사용하지 마세요

**✅ Good / 좋음:**
```go
// Wrap at meaningful boundaries
// 의미 있는 경계에서 래핑
func handler() error {
    err := service.DoWork()
    if err != nil {
        return errorutil.Wrap(err, "handler failed")
    }
    return nil
}

func service() error {
    err := repository.Save()
    if err != nil {
        return errorutil.Wrap(err, "service failed")
    }
    return nil
}
```

**❌ Bad / 나쁨:**
```go
// Too much wrapping, loses clarity
// 과도한 래핑, 명확성 상실
func helper1() error {
    err := helper2()
    return errorutil.Wrap(err, "helper1")
}

func helper2() error {
    err := helper3()
    return errorutil.Wrap(err, "helper2")
}

func helper3() error {
    err := doWork()
    return errorutil.Wrap(err, "helper3")
}
// Result: helper1: helper2: helper3: actual error
```

---

## Common Patterns / 일반적인 패턴

### Pattern 1: Validation Errors / 검증 에러

```go
// Validation function
// 검증 함수
func validateUser(user *User) error {
    if user.Email == "" {
        return errorutil.WithCode("VALIDATION_ERROR",
            "email is required")
    }

    if !isValidEmail(user.Email) {
        return errorutil.WithCodef("VALIDATION_ERROR",
            "invalid email format: %s", user.Email)
    }

    if user.Age < 18 {
        return errorutil.WithCode("VALIDATION_ERROR",
            "user must be 18 or older")
    }

    return nil
}

// Usage in handler
// 핸들러에서 사용
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    if err := validateUser(&user); err != nil {
        if errorutil.HasCode(err, "VALIDATION_ERROR") {
            http.Error(w, err.Error(), 400)
            return
        }
    }

    // Create user...
    // 사용자 생성...
}
```

### Pattern 2: Database Errors / 데이터베이스 에러

```go
// Repository layer
// 저장소 레이어
func (r *UserRepository) GetByID(id int) (*User, error) {
    var user User
    err := r.db.Get(&user, "SELECT * FROM users WHERE id = ?", id)

    if err == sql.ErrNoRows {
        // Not found - client error
        // 찾지 못함 - 클라이언트 에러
        return nil, errorutil.WithNumericCodef(404,
            "user %d not found", id)
    }

    if err != nil {
        // Database error - server error
        // 데이터베이스 에러 - 서버 에러
        return nil, errorutil.WrapWithCode(err, "DB_ERROR",
            "failed to query user")
    }

    return &user, nil
}

// Service layer wraps with more context
// 서비스 레이어가 더 많은 컨텍스트로 래핑
func (s *UserService) GetUser(id int) (*User, error) {
    user, err := s.repo.GetByID(id)
    if err != nil {
        return nil, errorutil.Wrapf(err,
            "failed to fetch user profile %d", id)
    }
    return user, nil
}

// HTTP handler checks error type
// HTTP 핸들러가 에러 타입 확인
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    id := getIDFromRequest(r)
    user, err := service.GetUser(id)

    if err != nil {
        if code, ok := errorutil.GetNumericCode(err); ok {
            http.Error(w, err.Error(), code)
            return
        }
        http.Error(w, "Internal server error", 500)
        return
    }

    json.NewEncoder(w).Encode(user)
}
```

### Pattern 3: External Service Errors / 외부 서비스 에러

```go
// External API client
// 외부 API 클라이언트
func (c *PaymentClient) Charge(amount int) error {
    resp, err := c.httpClient.Post("/charge", amount)

    if err != nil {
        // Network error
        // 네트워크 에러
        return errorutil.WrapWithNumericCode(err, 502,
            "payment service unavailable")
    }

    if resp.StatusCode == 402 {
        // Payment required
        // 결제 필요
        return errorutil.WithNumericCode(402,
            "insufficient funds")
    }

    if resp.StatusCode >= 500 {
        // External service error
        // 외부 서비스 에러
        return errorutil.WithNumericCodef(503,
            "payment service error: %d", resp.StatusCode)
    }

    return nil
}

// Service layer handles payment errors
// 서비스 레이어가 결제 에러 처리
func (s *OrderService) ProcessPayment(orderID int) error {
    err := s.paymentClient.Charge(calculateTotal(orderID))

    if err != nil {
        // Add order context
        // 주문 컨텍스트 추가
        return errorutil.Wrapf(err,
            "failed to process payment for order %d", orderID)
    }

    return nil
}

// Handler provides appropriate response
// 핸들러가 적절한 응답 제공
func processOrderHandler(w http.ResponseWriter, r *http.Request) {
    orderID := getOrderID(r)
    err := service.ProcessPayment(orderID)

    if err != nil {
        if errorutil.HasNumericCode(err, 402) {
            // Payment required - client's fault
            // 결제 필요 - 클라이언트 책임
            http.Error(w, "Insufficient funds", 402)
            return
        }

        if errorutil.HasNumericCode(err, 503) {
            // Service unavailable - retry later
            // 서비스 사용 불가 - 나중에 재시도
            http.Error(w, "Payment service temporarily unavailable", 503)
            return
        }

        http.Error(w, "Payment failed", 500)
        return
    }

    w.WriteHeader(200)
}
```

### Pattern 4: Retry with Error Codes / 에러 코드로 재시도

```go
// Retry function based on error code
// 에러 코드 기반 재시도 함수
func retryableOperation(fn func() error) error {
    maxRetries := 3
    retryDelay := time.Second

    for attempt := 1; attempt <= maxRetries; attempt++ {
        err := fn()

        if err == nil {
            return nil // Success
        }

        // Don't retry client errors
        // 클라이언트 에러는 재시도하지 않음
        if code, ok := errorutil.GetNumericCode(err); ok {
            if code >= 400 && code < 500 {
                return err // Client error, don't retry
            }
        }

        // Don't retry permanent errors
        // 영구 에러는 재시도하지 않음
        if errorutil.HasCode(err, "VALIDATION_ERROR") {
            return err
        }

        // Retry server errors
        // 서버 에러 재시도
        if attempt < maxRetries {
            time.Sleep(retryDelay)
            retryDelay *= 2 // Exponential backoff
            continue
        }

        // Max retries exceeded
        // 최대 재시도 횟수 초과
        return errorutil.Wrapf(err,
            "failed after %d attempts", maxRetries)
    }

    return nil
}

// Usage
// 사용법
err := retryableOperation(func() error {
    return callExternalAPI()
})
```

---

## Troubleshooting / 문제 해결

### Problem: Error Code Not Detected Through Chain / 문제: 체인을 통해 에러 코드가 감지되지 않음

**Symptom / 증상:**
```go
err := errorutil.WithCode("ERR001", "original error")
wrapped := fmt.Errorf("wrapped: %w", err)

// Returns false!
// false 반환!
errorutil.HasCode(wrapped, "ERR001")
```

**Solution / 해결책:**

Use errorutil wrapping functions, not `fmt.Errorf`:

`fmt.Errorf` 대신 errorutil 래핑 함수 사용:

```go
// ✅ Correct
// 올바름
wrapped := errorutil.Wrap(err, "wrapped")
errorutil.HasCode(wrapped, "ERR001") // true
```

**Why / 이유:**

`fmt.Errorf` creates a standard error that doesn't preserve errorutil's interfaces. Always use errorutil's wrapping functions.

`fmt.Errorf`는 errorutil의 인터페이스를 보존하지 않는 표준 에러를 생성합니다. 항상 errorutil의 래핑 함수를 사용하세요.

### Problem: GetCode Returns Empty String / 문제: GetCode가 빈 문자열 반환

**Symptom / 증상:**
```go
err := errorutil.WithNumericCode(404, "not found")
code, ok := errorutil.GetCode(err) // "", false
```

**Solution / 해결책:**

Use the correct getter for the code type:

코드 타입에 맞는 올바른 getter 사용:

```go
// ✅ Correct
// 올바름
code, ok := errorutil.GetNumericCode(err) // 404, true
```

**Why / 이유:**

`GetCode` is for string codes, `GetNumericCode` is for numeric codes. They are separate interfaces.

`GetCode`는 문자열 코드용이고, `GetNumericCode`는 숫자 코드용입니다. 별도의 인터페이스입니다.

### Problem: Nil Pointer When Wrapping / 문제: 래핑 시 Nil 포인터

**Symptom / 증상:**
```go
var err error // nil
wrapped := errorutil.Wrap(err, "wrapped")
// wrapped is nil, not an error with message "wrapped"
```

**Solution / 해결책:**

Always check if error is nil before wrapping:

래핑하기 전에 항상 에러가 nil인지 확인:

```go
// ✅ Correct
// 올바름
if err != nil {
    return errorutil.Wrap(err, "wrapped")
}
return nil
```

**Why / 이유:**

Wrapping a nil error returns nil. This is intentional to allow clean error handling flow.

nil 에러를 래핑하면 nil을 반환합니다. 이것은 깔끔한 에러 처리 흐름을 허용하기 위한 의도적인 동작입니다.

### Problem: Error Message Too Long / 문제: 에러 메시지가 너무 김

**Symptom / 증상:**
```go
// Deeply nested wrapping
// 깊게 중첩된 래핑
err.Error() // "layer5: layer4: layer3: layer2: layer1: original error"
```

**Solution / 해결책:**

Only wrap at meaningful boundaries:

의미 있는 경계에서만 래핑:

```go
// ✅ Better
// 더 나음
// Only wrap at architectural layer boundaries
// 아키텍처 레이어 경계에서만 래핑
// - HTTP Handler
// - Service Layer
// - Repository Layer
```

**Why / 이유:**

Too much wrapping adds noise. Wrap only where it adds meaningful context.

과도한 래핑은 잡음을 추가합니다. 의미 있는 컨텍스트를 추가하는 곳에서만 래핑하세요.

---

## API Reference / API 참조

### Error Creation Functions / 에러 생성 함수

#### New
```go
func New(message string) error
```
Creates a new error with the given message.
주어진 메시지로 새 에러를 생성합니다.

**Parameters / 매개변수:**
- `message`: Error message / 에러 메시지

**Returns / 반환:**
- `error`: New error / 새 에러

**Example / 예제:**
```go
err := errorutil.New("something went wrong")
```

#### Newf
```go
func Newf(format string, args ...interface{}) error
```
Creates a new error with a formatted message.
포맷된 메시지로 새 에러를 생성합니다.

**Parameters / 매개변수:**
- `format`: Format string / 포맷 문자열
- `args`: Format arguments / 포맷 인자

**Returns / 반환:**
- `error`: New formatted error / 새 포맷된 에러

**Example / 예제:**
```go
err := errorutil.Newf("user %d not found", 123)
```

#### WithCode
```go
func WithCode(code, message string) error
```
Creates an error with a string code.
문자열 코드와 함께 에러를 생성합니다.

**Parameters / 매개변수:**
- `code`: Error code / 에러 코드
- `message`: Error message / 에러 메시지

**Returns / 반환:**
- `error`: Error with code / 코드가 있는 에러

**Example / 예제:**
```go
err := errorutil.WithCode("ERR001", "validation failed")
```

#### WithCodef
```go
func WithCodef(code, format string, args ...interface{}) error
```
Creates an error with a string code and formatted message.
문자열 코드와 포맷된 메시지로 에러를 생성합니다.

**Parameters / 매개변수:**
- `code`: Error code / 에러 코드
- `format`: Format string / 포맷 문자열
- `args`: Format arguments / 포맷 인자

**Returns / 반환:**
- `error`: Error with code and formatted message / 코드와 포맷된 메시지가 있는 에러

**Example / 예제:**
```go
err := errorutil.WithCodef("ERR001", "field %s is invalid", "email")
```

#### WithNumericCode
```go
func WithNumericCode(code int, message string) error
```
Creates an error with a numeric code.
숫자 코드와 함께 에러를 생성합니다.

**Parameters / 매개변수:**
- `code`: Numeric error code (e.g., HTTP status) / 숫자 에러 코드 (예: HTTP 상태)
- `message`: Error message / 에러 메시지

**Returns / 반환:**
- `error`: Error with numeric code / 숫자 코드가 있는 에러

**Example / 예제:**
```go
err := errorutil.WithNumericCode(404, "user not found")
```

#### WithNumericCodef
```go
func WithNumericCodef(code int, format string, args ...interface{}) error
```
Creates an error with a numeric code and formatted message.
숫자 코드와 포맷된 메시지로 에러를 생성합니다.

**Parameters / 매개변수:**
- `code`: Numeric error code / 숫자 에러 코드
- `format`: Format string / 포맷 문자열
- `args`: Format arguments / 포맷 인자

**Returns / 반환:**
- `error`: Error with numeric code and formatted message / 숫자 코드와 포맷된 메시지가 있는 에러

**Example / 예제:**
```go
err := errorutil.WithNumericCodef(404, "user %d not found", 123)
```

### Error Wrapping Functions / 에러 래핑 함수

#### Wrap
```go
func Wrap(cause error, message string) error
```
Wraps an error with additional context.
추가 컨텍스트와 함께 에러를 래핑합니다.

**Parameters / 매개변수:**
- `cause`: Original error to wrap / 래핑할 원본 에러
- `message`: Additional context message / 추가 컨텍스트 메시지

**Returns / 반환:**
- `error`: Wrapped error, or nil if cause is nil / 래핑된 에러, cause가 nil이면 nil

**Example / 예제:**
```go
err := errorutil.Wrap(dbErr, "failed to save user")
```

#### Wrapf
```go
func Wrapf(cause error, format string, args ...interface{}) error
```
Wraps an error with a formatted message.
포맷된 메시지로 에러를 래핑합니다.

**Parameters / 매개변수:**
- `cause`: Original error to wrap / 래핑할 원본 에러
- `format`: Format string / 포맷 문자열
- `args`: Format arguments / 포맷 인자

**Returns / 반환:**
- `error`: Wrapped error with formatted message / 포맷된 메시지가 있는 래핑된 에러

**Example / 예제:**
```go
err := errorutil.Wrapf(dbErr, "failed to save user %d", userID)
```

#### WrapWithCode
```go
func WrapWithCode(cause error, code, message string) error
```
Wraps an error and adds a string code.
에러를 래핑하고 문자열 코드를 추가합니다.

**Parameters / 매개변수:**
- `cause`: Original error to wrap / 래핑할 원본 에러
- `code`: Error code to add / 추가할 에러 코드
- `message`: Wrapping message / 래핑 메시지

**Returns / 반환:**
- `error`: Wrapped error with code / 코드가 있는 래핑된 에러

**Example / 예제:**
```go
err := errorutil.WrapWithCode(ioErr, "FILE_ERROR", "failed to read config")
```

#### WrapWithCodef
```go
func WrapWithCodef(cause error, code, format string, args ...interface{}) error
```
Wraps an error with a code and formatted message.
코드와 포맷된 메시지로 에러를 래핑합니다.

**Parameters / 매개변수:**
- `cause`: Original error / 원본 에러
- `code`: Error code / 에러 코드
- `format`: Format string / 포맷 문자열
- `args`: Format arguments / 포맷 인자

**Returns / 반환:**
- `error`: Wrapped error with code and formatted message / 코드와 포맷된 메시지가 있는 래핑된 에러

**Example / 예제:**
```go
err := errorutil.WrapWithCodef(ioErr, "FILE_ERROR", "failed to read %s", filename)
```

#### WrapWithNumericCode
```go
func WrapWithNumericCode(cause error, code int, message string) error
```
Wraps an error and adds a numeric code.
에러를 래핑하고 숫자 코드를 추가합니다.

**Parameters / 매개변수:**
- `cause`: Original error / 원본 에러
- `code`: Numeric error code / 숫자 에러 코드
- `message`: Wrapping message / 래핑 메시지

**Returns / 반환:**
- `error`: Wrapped error with numeric code / 숫자 코드가 있는 래핑된 에러

**Example / 예제:**
```go
err := errorutil.WrapWithNumericCode(dbErr, 404, "user not found")
```

#### WrapWithNumericCodef
```go
func WrapWithNumericCodef(cause error, code int, format string, args ...interface{}) error
```
Wraps an error with a numeric code and formatted message.
숫자 코드와 포맷된 메시지로 에러를 래핑합니다.

**Parameters / 매개변수:**
- `cause`: Original error / 원본 에러
- `code`: Numeric error code / 숫자 에러 코드
- `format`: Format string / 포맷 문자열
- `args`: Format arguments / 포맷 인자

**Returns / 반환:**
- `error`: Wrapped error with numeric code and formatted message / 숫자 코드와 포맷된 메시지가 있는 래핑된 에러

**Example / 예제:**
```go
err := errorutil.WrapWithNumericCodef(dbErr, 404, "user %d not found", userID)
```

### Error Inspection Functions / 에러 검사 함수

#### HasCode
```go
func HasCode(err error, code string) bool
```
Checks if an error has a specific string code.
에러가 특정 문자열 코드를 가지는지 확인합니다.

**Parameters / 매개변수:**
- `err`: Error to check / 확인할 에러
- `code`: Code to look for / 찾을 코드

**Returns / 반환:**
- `bool`: true if code is found / 코드를 찾으면 true

**Example / 예제:**
```go
if errorutil.HasCode(err, "VALIDATION_ERROR") {
    // Handle validation error
}
```

#### HasNumericCode
```go
func HasNumericCode(err error, code int) bool
```
Checks if an error has a specific numeric code.
에러가 특정 숫자 코드를 가지는지 확인합니다.

**Parameters / 매개변수:**
- `err`: Error to check / 확인할 에러
- `code`: Numeric code to look for / 찾을 숫자 코드

**Returns / 반환:**
- `bool`: true if code is found / 코드를 찾으면 true

**Example / 예제:**
```go
if errorutil.HasNumericCode(err, 404) {
    // Handle 404 error
}
```

#### GetCode
```go
func GetCode(err error) (string, bool)
```
Extracts the string code from an error.
에러에서 문자열 코드를 추출합니다.

**Parameters / 매개변수:**
- `err`: Error to extract code from / 코드를 추출할 에러

**Returns / 반환:**
- `string`: Error code, or empty string if not found / 에러 코드, 찾지 못하면 빈 문자열
- `bool`: true if code was found / 코드를 찾으면 true

**Example / 예제:**
```go
if code, ok := errorutil.GetCode(err); ok {
    fmt.Printf("Error code: %s\n", code)
}
```

#### GetNumericCode
```go
func GetNumericCode(err error) (int, bool)
```
Extracts the numeric code from an error.
에러에서 숫자 코드를 추출합니다.

**Parameters / 매개변수:**
- `err`: Error to extract code from / 코드를 추출할 에러

**Returns / 반환:**
- `int`: Numeric code, or 0 if not found / 숫자 코드, 찾지 못하면 0
- `bool`: true if code was found / 코드를 찾으면 true

**Example / 예제:**
```go
if code, ok := errorutil.GetNumericCode(err); ok {
    http.Error(w, err.Error(), code)
}
```

#### GetStackTrace
```go
func GetStackTrace(err error) ([]Frame, bool)
```
Extracts the stack trace from an error or any error in its chain.
에러 또는 에러 체인에서 스택 트레이스를 추출합니다.

**Parameters / 매개변수:**
- `err`: Error to extract stack trace from / 스택 트레이스를 추출할 에러

**Returns / 반환:**
- `[]Frame`: Stack trace, or nil if not found / 스택 트레이스, 찾지 못하면 nil
- `bool`: true if stack trace was found / 스택 트레이스를 찾으면 true

**Example / 예제:**
```go
if stack, ok := errorutil.GetStackTrace(err); ok {
    for _, frame := range stack {
        fmt.Println(frame.String())
    }
}
```

#### GetContext
```go
func GetContext(err error) (map[string]interface{}, bool)
```
Extracts the context data from an error or any error in its chain.
에러 또는 에러 체인에서 컨텍스트 데이터를 추출합니다.

**Parameters / 매개변수:**
- `err`: Error to extract context from / 컨텍스트를 추출할 에러

**Returns / 반환:**
- `map[string]interface{}`: Context data, or nil if not found / 컨텍스트 데이터, 찾지 못하면 nil
- `bool`: true if context was found / 컨텍스트를 찾으면 true

**Example / 예제:**
```go
if ctx, ok := errorutil.GetContext(err); ok {
    fmt.Printf("User ID: %v\n", ctx["user_id"])
    fmt.Printf("Action: %v\n", ctx["action"])
}
```

#### Root
```go
func Root(err error) error
```
Returns the root (innermost) error in the error chain.
에러 체인의 루트(가장 안쪽) 에러를 반환합니다.

**Parameters / 매개변수:**
- `err`: Error to get root from / 루트를 가져올 에러

**Returns / 반환:**
- `error`: Root error, or nil if input is nil / 루트 에러, 입력이 nil이면 nil

**Example / 예제:**
```go
baseErr := errors.New("base error")
err1 := errorutil.Wrap(baseErr, "layer 1")
err2 := errorutil.Wrap(err1, "layer 2")
err3 := errorutil.Wrap(err2, "layer 3")

root := errorutil.Root(err3)
fmt.Println(root) // Output: base error
```

**Use Cases / 사용 사례:**
- Finding the original error in a long chain / 긴 체인에서 원본 에러 찾기
- Logging root causes / 근본 원인 로깅
- Error analysis / 에러 분석

#### UnwrapAll
```go
func UnwrapAll(err error) []error
```
Returns all errors in the error chain as a slice. The first element is the outermost error, and the last is the root error.
에러 체인의 모든 에러를 슬라이스로 반환합니다. 첫 번째 요소는 가장 바깥쪽 에러이고, 마지막은 루트 에러입니다.

**Parameters / 매개변수:**
- `err`: Error to unwrap / 언래핑할 에러

**Returns / 반환:**
- `[]error`: Slice of all errors in the chain, or nil if input is nil / 체인의 모든 에러 슬라이스, 입력이 nil이면 nil

**Example / 예제:**
```go
baseErr := errors.New("base error")
err1 := errorutil.Wrap(baseErr, "layer 1")
err2 := errorutil.Wrap(err1, "layer 2")

chain := errorutil.UnwrapAll(err2)
for i, e := range chain {
    fmt.Printf("Level %d: %v\n", i, e)
}
// Output:
// Level 0: layer 2: layer 1: base error
// Level 1: layer 1: base error
// Level 2: base error
```

**Use Cases / 사용 사례:**
- Analyzing error chains / 에러 체인 분석
- Detailed error logging / 상세 에러 로깅
- Debugging multi-layer errors / 다층 에러 디버깅

#### Contains
```go
func Contains(err error, target error) bool
```
Checks if the error chain contains a specific error using errors.Is().
errors.Is()를 사용하여 에러 체인이 특정 에러를 포함하는지 확인합니다.

**Parameters / 매개변수:**
- `err`: Error chain to check / 확인할 에러 체인
- `target`: Error to look for / 찾을 에러

**Returns / 반환:**
- `bool`: true if target is found in the chain, false otherwise / 체인에서 target을 찾으면 true, 아니면 false

**Example / 예제:**
```go
var ErrNotFound = errors.New("not found")
var ErrValidation = errors.New("validation error")

err := errorutil.Wrap(ErrNotFound, "failed to get user")

if errorutil.Contains(err, ErrNotFound) {
    fmt.Println("This is a not found error")
}

if errorutil.Contains(err, ErrValidation) {
    fmt.Println("This won't print")
}
```

**Use Cases / 사용 사례:**
- Checking for sentinel errors / 센티널 에러 확인
- Error type classification / 에러 타입 분류
- Conditional error handling / 조건부 에러 처리

---

## Additional Resources / 추가 자료

### Documentation / 문서

- [Package README](../../errorutil/README.md) - Quick start and overview
- [DEVELOPER_GUIDE.md](./DEVELOPER_GUIDE.md) - Internal architecture and design
- [Go Blog: Error Handling](https://go.dev/blog/go1.13-errors) - Official Go 1.13 error handling

### Related Packages / 관련 패키지

- Standard `errors` package - Go's built-in error handling
- Standard `fmt` package - For `fmt.Errorf` compatibility

### Source Code / 소스 코드

- [GitHub Repository](https://github.com/arkd0ng/go-utils)
- [errorutil Package](https://github.com/arkd0ng/go-utils/tree/main/errorutil)

---

**Last Updated / 최종 업데이트**: 2025-10-17
**Version / 버전**: v1.12.011
**Package / 패키지**: errorutil
