# errorutil Package / errorutil 패키지

Comprehensive error handling utilities for Go applications. / Go 애플리케이션을 위한 포괄적인 에러 처리 유틸리티입니다.

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](../LICENSE)
[![Coverage](https://img.shields.io/badge/Coverage-99.2%25-brightgreen.svg)](.)

## Table of Contents / 목차

- [Overview](#overview--개요)
- [Features](#features--주요-기능)
- [Installation](#installation--설치)
- [Quick Start](#quick-start--빠른-시작)
- [API Reference](#api-reference--api-참조)
  - [Error Creation](#error-creation--에러-생성)
  - [Error Wrapping](#error-wrapping--에러-래핑)
  - [Error Inspection](#error-inspection--에러-검사)
- [Examples](#examples--예제)
- [Best Practices](#best-practices--모범-사례)
- [Documentation](#documentation--문서)

## Overview / 개요

The `errorutil` package provides enhanced error handling capabilities beyond Go's standard library, while maintaining full compatibility with `errors` package and error wrapping introduced in Go 1.13+.

`errorutil` 패키지는 Go 표준 라이브러리를 넘어서는 향상된 에러 처리 기능을 제공하며, Go 1.13+에서 도입된 `errors` 패키지 및 에러 래핑과 완전히 호환됩니다.

### Why errorutil? / 왜 errorutil을 사용하나요?

- **Error Codes**: Attach string or numeric codes for easy error categorization / 쉬운 에러 분류를 위한 문자열 또는 숫자 코드 첨부
- **Error Chain Inspection**: Traverse error chains to find specific errors or codes / 특정 에러 또는 코드를 찾기 위한 에러 체인 탐색
- **Type Safety**: Full compatibility with `errors.Is` and `errors.As` / `errors.Is` 및 `errors.As`와 완전 호환
- **Zero Dependencies**: No external dependencies except Go standard library / Go 표준 라이브러리 외에 외부 의존성 없음
- **High Performance**: Minimal overhead with 99.2% test coverage / 99.2% 테스트 커버리지로 최소 오버헤드

## Features / 주요 기능

✅ **Error Creation** - Create errors with or without codes / 코드와 함께 또는 코드 없이 에러 생성
✅ **Error Wrapping** - Add context while preserving the error chain / 에러 체인을 보존하면서 컨텍스트 추가
✅ **String Codes** - Categorize errors with string codes (e.g., "ERR001", "VALIDATION_ERROR") / 문자열 코드로 에러 분류
✅ **Numeric Codes** - Use numeric codes for HTTP status codes (e.g., 404, 500) / HTTP 상태 코드용 숫자 코드 사용
✅ **Error Inspection** - Check codes and extract information from error chains / 에러 체인에서 코드 확인 및 정보 추출
✅ **Go 1.13+ Compatible** - Works seamlessly with `errors.Is`, `errors.As`, and `errors.Unwrap` / `errors.Is`, `errors.As`, `errors.Unwrap`과 원활하게 작동

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/errorutil
```

## Quick Start / 빠른 시작

### Basic Usage / 기본 사용법

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/errorutil"
)

func main() {
    // Create a simple error / 간단한 에러 생성
    err := errorutil.New("something went wrong")
    fmt.Println(err) // Output: something went wrong

    // Create error with code / 코드와 함께 에러 생성
    err = errorutil.WithCode("ERR001", "invalid input")
    fmt.Println(err) // Output: [ERR001] invalid input

    // Create error with numeric code (e.g., HTTP status) / 숫자 코드로 에러 생성 (예: HTTP 상태)
    err = errorutil.WithNumericCode(404, "user not found")
    fmt.Println(err) // Output: [404] user not found
}
```

### Error Wrapping / 에러 래핑

```go
// Wrap errors to add context / 컨텍스트를 추가하기 위해 에러 래핑
func fetchUser(id int) error {
    err := database.Query(id)
    if err != nil {
        // Wrap with additional context / 추가 컨텍스트와 함께 래핑
        return errorutil.Wrapf(err, "failed to fetch user %d", id)
    }
    return nil
}

// Wrap with error code / 에러 코드와 함께 래핑
func processPayment(amount float64) error {
    err := paymentGateway.Charge(amount)
    if err != nil {
        return errorutil.WrapWithCode(err, "PAYMENT_FAILED", "payment processing failed")
    }
    return nil
}
```

### Error Inspection / 에러 검사

```go
err := fetchUser(123)

// Check if error has a specific code / 특정 코드를 가진 에러인지 확인
if errorutil.HasCode(err, "ERR001") {
    fmt.Println("Validation error occurred")
}

// Extract error code / 에러 코드 추출
if code, ok := errorutil.GetCode(err); ok {
    fmt.Printf("Error code: %s\n", code)
}

// Check numeric code (e.g., HTTP status) / 숫자 코드 확인 (예: HTTP 상태)
if errorutil.HasNumericCode(err, 404) {
    fmt.Println("Not found error")
}
```

## API Reference / API 참조

### Error Creation / 에러 생성

#### New(message string) error

Creates a new error with the given message. / 주어진 메시지로 새로운 에러를 생성합니다.

```go
err := errorutil.New("something went wrong")
```

#### Newf(format string, args ...interface{}) error

Creates a new error with a formatted message. / 포맷된 메시지로 새로운 에러를 생성합니다.

```go
err := errorutil.Newf("failed to process user %d", 123)
```

#### WithCode(code, message string) error

Creates an error with a string error code. / 문자열 에러 코드와 함께 에러를 생성합니다.

```go
err := errorutil.WithCode("ERR001", "invalid input")
```

#### WithCodef(code, format string, args ...interface{}) error

Creates an error with a string code and formatted message. / 문자열 코드와 포맷된 메시지로 에러를 생성합니다.

```go
err := errorutil.WithCodef("ERR001", "invalid user: %d", 123)
```

#### WithNumericCode(code int, message string) error

Creates an error with a numeric error code. / 숫자 에러 코드와 함께 에러를 생성합니다.

```go
err := errorutil.WithNumericCode(404, "user not found")
```

#### WithNumericCodef(code int, format string, args ...interface{}) error

Creates an error with a numeric code and formatted message. / 숫자 코드와 포맷된 메시지로 에러를 생성합니다.

```go
err := errorutil.WithNumericCodef(500, "database error: %s", "connection timeout")
```

### Error Wrapping / 에러 래핑

#### Wrap(cause error, message string) error

Wraps an existing error with a new message. Returns `nil` if cause is `nil`. / 기존 에러를 새로운 메시지로 래핑합니다. cause가 `nil`이면 `nil`을 반환합니다.

```go
wrapped := errorutil.Wrap(err, "failed to connect to database")
```

#### Wrapf(cause error, format string, args ...interface{}) error

Wraps an error with a formatted message. / 포맷된 메시지로 에러를 래핑합니다.

```go
wrapped := errorutil.Wrapf(err, "failed to find user %d", 123)
```

#### WrapWithCode(cause error, code, message string) error

Wraps an error with a message and string code. / 메시지와 문자열 코드로 에러를 래핑합니다.

```go
wrapped := errorutil.WrapWithCode(err, "ERR001", "invalid input")
```

#### WrapWithCodef(cause error, code, format string, args ...interface{}) error

Wraps an error with a code and formatted message. / 코드와 포맷된 메시지로 에러를 래핑합니다.

```go
wrapped := errorutil.WrapWithCodef(err, "ERR404", "user %d not found", 123)
```

#### WrapWithNumericCode(cause error, code int, message string) error

Wraps an error with a message and numeric code. / 메시지와 숫자 코드로 에러를 래핑합니다.

```go
wrapped := errorutil.WrapWithNumericCode(err, 500, "internal server error")
```

#### WrapWithNumericCodef(cause error, code int, format string, args ...interface{}) error

Wraps an error with a numeric code and formatted message. / 숫자 코드와 포맷된 메시지로 에러를 래핑합니다.

```go
wrapped := errorutil.WrapWithNumericCodef(err, 408, "timeout after %d seconds", 30)
```

### Error Inspection / 에러 검사

#### HasCode(err error, code string) bool

Checks if an error or any error in its chain has the specified string code. / 에러 또는 에러 체인의 어떤 에러가 지정된 문자열 코드를 가지고 있는지 확인합니다.

```go
if errorutil.HasCode(err, "ERR001") {
    // Handle ERR001 / ERR001 처리
}
```

#### HasNumericCode(err error, code int) bool

Checks if an error or any error in its chain has the specified numeric code. / 에러 또는 에러 체인의 어떤 에러가 지정된 숫자 코드를 가지고 있는지 확인합니다.

```go
if errorutil.HasNumericCode(err, 404) {
    // Handle 404 Not Found / 404 Not Found 처리
}
```

#### GetCode(err error) (string, bool)

Extracts the string code from an error or any error in its chain. / 에러 또는 에러 체인에서 문자열 코드를 추출합니다.

```go
if code, ok := errorutil.GetCode(err); ok {
    fmt.Printf("Error code: %s\n", code)
}
```

#### GetNumericCode(err error) (int, bool)

Extracts the numeric code from an error or any error in its chain. / 에러 또는 에러 체인에서 숫자 코드를 추출합니다.

```go
if code, ok := errorutil.GetNumericCode(err); ok {
    fmt.Printf("HTTP status: %d\n", code)
}
```

## Examples / 예제

### HTTP API Error Handling / HTTP API 에러 처리

```go
func handleUser(w http.ResponseWriter, r *http.Request) {
    err := processUserRequest(r)
    if err != nil {
        // Extract HTTP status code from error / 에러에서 HTTP 상태 코드 추출
        status := 500 // default / 기본값
        if code, ok := errorutil.GetNumericCode(err); ok {
            status = code
        }

        http.Error(w, err.Error(), status)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func processUserRequest(r *http.Request) error {
    user, err := findUser(r.Context(), 123)
    if err != nil {
        // Wrap with HTTP 404 code / HTTP 404 코드와 함께 래핑
        return errorutil.WrapWithNumericCode(err, 404, "user not found")
    }

    err = validateUser(user)
    if err != nil {
        // Wrap with HTTP 400 code / HTTP 400 코드와 함께 래핑
        return errorutil.WrapWithNumericCode(err, 400, "invalid user data")
    }

    return nil
}
```

### Error Classification / 에러 분류

```go
const (
    ErrCodeValidation = "VALIDATION_ERROR"
    ErrCodeDatabase   = "DATABASE_ERROR"
    ErrCodeAuth       = "AUTH_ERROR"
)

func processData(data string) error {
    if data == "" {
        return errorutil.WithCode(ErrCodeValidation, "data cannot be empty")
    }

    err := saveToDatabase(data)
    if err != nil {
        return errorutil.WrapWithCode(err, ErrCodeDatabase, "failed to save data")
    }

    return nil
}

// Error handling / 에러 처리
err := processData("")
if errorutil.HasCode(err, ErrCodeValidation) {
    // Handle validation error / 검증 에러 처리
    log.Warn("Validation failed:", err)
    return
}
if errorutil.HasCode(err, ErrCodeDatabase) {
    // Handle database error / 데이터베이스 에러 처리
    log.Error("Database error:", err)
    return
}
```

### Deeply Nested Error Chains / 깊게 중첩된 에러 체인

```go
func layer1() error {
    err := layer2()
    if err != nil {
        return errorutil.Wrap(err, "layer1 failed")
    }
    return nil
}

func layer2() error {
    err := layer3()
    if err != nil {
        return errorutil.WrapWithCode(err, "ERR002", "layer2 failed")
    }
    return nil
}

func layer3() error {
    return errorutil.WithCode("ERR001", "something went wrong")
}

// Usage / 사용
err := layer1()
// Error chain: layer1 -> layer2 -> layer3
// errorutil can find ERR001 through the entire chain
// errorutil은 전체 체인을 통해 ERR001을 찾을 수 있습니다

if errorutil.HasCode(err, "ERR001") {
    fmt.Println("Found ERR001 in the chain!")
}

if code, ok := errorutil.GetCode(err); ok {
    fmt.Printf("First error code in chain: %s\n", code) // Output: ERR002
}
```

## Best Practices / 모범 사례

### 1. Use Codes for Categorization / 분류를 위해 코드 사용

Define error codes as constants for consistency: / 일관성을 위해 에러 코드를 상수로 정의:

```go
const (
    ErrCodeInvalidInput   = "INVALID_INPUT"
    ErrCodeNotFound       = "NOT_FOUND"
    ErrCodeUnauthorized   = "UNAUTHORIZED"
)

err := errorutil.WithCode(ErrCodeInvalidInput, "email is required")
```

### 2. Wrap Errors at Layer Boundaries / 계층 경계에서 에러 래핑

Add context when errors cross architectural boundaries: / 에러가 아키텍처 경계를 넘을 때 컨텍스트 추가:

```go
// Service layer / 서비스 계층
func (s *UserService) CreateUser(data UserData) error {
    err := s.repo.Save(data)
    if err != nil {
        return errorutil.Wrapf(err, "failed to create user %s", data.Email)
    }
    return nil
}

// HTTP handler layer / HTTP 핸들러 계층
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    err := h.service.CreateUser(userData)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}
```

### 3. Check Errors Before Wrapping / 래핑 전에 에러 확인

All wrapping functions return `nil` if the cause is `nil`: / 모든 래핑 함수는 cause가 `nil`이면 `nil`을 반환:

```go
err := maybeReturnsNil()
// Safe - returns nil if err is nil / 안전 - err이 nil이면 nil 반환
wrapped := errorutil.Wrap(err, "operation failed")
```

### 4. Use Numeric Codes for HTTP / HTTP에는 숫자 코드 사용

Use numeric codes for HTTP status codes: / HTTP 상태 코드에는 숫자 코드 사용:

```go
func fetchResource(id int) error {
    resource, err := db.FindByID(id)
    if err == sql.ErrNoRows {
        return errorutil.WithNumericCode(404, "resource not found")
    }
    if err != nil {
        return errorutil.WithNumericCode(500, "database error")
    }
    return nil
}
```

## Documentation / 문서

For more detailed documentation, see: / 더 자세한 문서는 다음을 참조하세요:

- [USER_MANUAL.md](../docs/errorutil/USER_MANUAL.md) - Comprehensive user guide / 종합 사용자 가이드
- [DEVELOPER_GUIDE.md](../docs/errorutil/DEVELOPER_GUIDE.md) - Architecture and internals / 아키텍처 및 내부 구조
- [Example Code](../examples/errorutil/) - Runnable examples / 실행 가능한 예제

## Version / 버전

Current version: **v1.12.008** / 현재 버전: **v1.12.008**

## License / 라이선스

MIT License - see [LICENSE](../LICENSE) for details / MIT 라이선스 - 자세한 내용은 [LICENSE](../LICENSE) 참조

---

**Made with ❤️ by the go-utils team** / **go-utils 팀이 ❤️로 만들었습니다**
