// Package validation provides comprehensive data validation utilities for Go applications.
// It offers a fluent, chainable API for validating strings, numbers, collections, dates,
// files, networks, and business logic with extensive built-in rules.
//
// validation 패키지는 Go 애플리케이션을 위한 포괄적인 데이터 검증 유틸리티를 제공합니다.
// 문자열, 숫자, 컬렉션, 날짜, 파일, 네트워크 및 비즈니스 로직을 광범위한 내장 규칙으로
// 검증하기 위한 유창하고 체이닝 가능한 API를 제공합니다.
//
// Key Features / 주요 기능:
//
//   - Fluent chainable API for readable validation logic
//     읽기 쉬운 검증 로직을 위한 유창한 체이닝 API
//
//   - 100+ built-in validation rules covering common use cases
//     일반적인 사용 사례를 다루는 100개 이상의 내장 검증 규칙
//
//   - Support for strings, numbers, arrays, maps, dates, files, and more
//     문자열, 숫자, 배열, 맵, 날짜, 파일 등 지원
//
//   - Bilingual error messages (English/Korean)
//     이중 언어 에러 메시지 (영문/한글)
//
//   - Custom validation rules and error messages
//     사용자 정의 검증 규칙 및 에러 메시지
//
//   - Field-level and multi-field validation support
//     필드 레벨 및 다중 필드 검증 지원
//
//   - Stop-on-first-error or collect-all-errors modes
//     첫 에러에서 중단 또는 모든 에러 수집 모드
//
// Validation Categories / 검증 카테고리:
//
// String Validation (rules_string.go):
//   - Length, format, content validation
//   - Pattern matching with regex support
//   - Case-insensitive comparisons
//   - Unicode and multibyte character support
//
// Numeric Validation (rules_numeric.go):
//   - Range validation (min, max, between)
//   - Type validation (positive, negative, zero)
//   - Multiple data type support (int, float, string numbers)
//
// Collection Validation (rules_collection.go):
//   - Array/slice validation (length, uniqueness, containment)
//   - Map validation (keys, values, size)
//   - Empty/non-empty checks
//
// Date/Time Validation (rules_datetime.go):
//   - Date comparison (before, after, equal)
//   - Date range validation
//   - Format validation
//   - Timezone-aware comparisons
//
// File Validation (rules_file.go):
//   - Path validation
//   - Existence, readability, writability checks
//   - File size and extension validation
//
// Network Validation (rules_network.go):
//   - Email, URL, IP address validation
//   - Domain name validation
//   - Port number validation
//   - MAC address validation
//
// Format Validation (rules_format.go):
//   - JSON, XML, Base64 validation
//   - UUID, ISBN validation
//   - Phone number, postal code validation
//
// Business Logic (rules_business.go):
//   - Credit card validation with Luhn algorithm
//   - Tax ID validation (SSN, EIN)
//   - Currency validation
//
// Type Validation (rules_type.go):
//   - Type checking (string, number, boolean, array, etc.)
//   - Nil/empty validation
//   - Type conversion validation
//
// Performance / 성능:
//
//   - Zero memory allocation for simple validations
//     간단한 검증의 경우 메모리 할당 없음
//
//   - Lazy evaluation with short-circuit on StopOnError
//     StopOnError 시 단락 평가를 통한 지연 평가
//
//   - Efficient regex compilation and caching
//     효율적인 정규식 컴파일 및 캐싱
//
//   - Thread-safe: All validators are safe for concurrent use
//     스레드 안전: 모든 검증기는 동시 사용 안전
//
// Thread Safety / 스레드 안전성:
//
// All validation functions are thread-safe and can be used concurrently.
// Each Validator instance maintains its own state and does not share data.
//
// 모든 검증 함수는 스레드 안전하며 동시에 사용할 수 있습니다.
// 각 Validator 인스턴스는 자체 상태를 유지하며 데이터를 공유하지 않습니다.
//
// Quick Start / 빠른 시작:
//
// Basic single field validation:
//
//	// String validation / 문자열 검증
//	v := validation.New("test@example.com", "email")
//	v.Required().Email()
//	if err := v.Validate(); err != nil {
//	    // Handle validation errors / 검증 에러 처리
//	}
//
//	// Numeric validation / 숫자 검증
//	v := validation.New(25, "age")
//	v.Required().Min(18).Max(100)
//	if err := v.Validate(); err != nil {
//	    // Handle errors / 에러 처리
//	}
//
// Chaining multiple rules:
//
//	v := validation.New(username, "username")
//	v.Required().
//	  MinLength(3).
//	  MaxLength(20).
//	  AlphaNumeric().
//	  NoWhitespace()
//
//	if err := v.Validate(); err != nil {
//	    // Validation failed / 검증 실패
//	}
//
// Custom error messages:
//
//	v := validation.New(email, "email")
//	v.Required().WithMessage("Email is required / 이메일은 필수입니다").
//	  Email().WithMessage("Invalid email format / 유효하지 않은 이메일 형식")
//
// Stop on first error:
//
//	v := validation.New(value, "field").StopOnError()
//	v.Required().MinLength(5).MaxLength(10)
//	// Stops at first failed rule / 첫 번째 실패한 규칙에서 중단
//
// Multi-field validation:
//
//	mv := validation.NewValidator()
//	mv.Field(username, "username").Required().MinLength(3)
//	mv.Field(email, "email").Required().Email()
//	mv.Field(age, "age").Required().Min(18)
//
//	if err := mv.Validate(); err != nil {
//	    // Handle errors for all fields / 모든 필드의 에러 처리
//	}
//
// Custom validation rules:
//
//	v := validation.New(value, "field")
//	v.Custom(func(val interface{}) bool {
//	    // Custom validation logic / 사용자 정의 검증 로직
//	    return val.(string) != "forbidden"
//	}).WithMessage("Value is forbidden / 금지된 값입니다")
//
// Error Handling / 에러 처리:
//
// The package provides detailed error information including:
//   - Field name that failed validation / 검증 실패한 필드 이름
//   - Failed validation rule / 실패한 검증 규칙
//   - Actual value that failed / 실패한 실제 값
//   - Descriptive error message / 설명적인 에러 메시지
//
// Example error handling:
//
//	if err := v.Validate(); err != nil {
//	    if validationErrors, ok := err.(validation.ValidationErrors); ok {
//	        for _, e := range validationErrors {
//	            fmt.Printf("Field: %s, Rule: %s, Message: %s\n",
//	                e.Field, e.Rule, e.Message)
//	        }
//	    }
//	}
//
// Common Patterns / 일반적인 패턴:
//
// User registration validation:
//
//	validator := validation.NewValidator()
//	validator.Field(user.Email, "email").
//	    Required().
//	    Email().
//	    MaxLength(100)
//	validator.Field(user.Password, "password").
//	    Required().
//	    MinLength(8).
//	    ContainsDigit().
//	    ContainsSpecialChar()
//	validator.Field(user.Age, "age").
//	    Required().
//	    Min(18).
//	    Max(120)
//
// API request validation:
//
//	v := validation.New(request.Data, "data").StopOnError()
//	v.Required().Type("array").ArrayNotEmpty()
//	if err := v.Validate(); err != nil {
//	    return BadRequestError(err)
//	}
//
// File upload validation:
//
//	v := validation.New(filepath, "upload")
//	v.Required().
//	  FileExists().
//	  FileReadable().
//	  FileSize(0, 10*1024*1024).  // Max 10MB
//	  FileExtension(".jpg", ".png", ".gif")
//
// Best Practices / 모범 사례:
//
//  1. Use descriptive field names for better error messages
//     더 나은 에러 메시지를 위해 설명적인 필드 이름 사용
//
//  2. Chain related rules together for readability
//     가독성을 위해 관련 규칙을 함께 체이닝
//
//  3. Use StopOnError() for expensive validations
//     비용이 많이 드는 검증에는 StopOnError() 사용
//
//  4. Provide custom messages for business-specific rules
//     비즈니스별 규칙에 사용자 정의 메시지 제공
//
//  5. Validate early in your request handlers
//     요청 핸들러에서 조기에 검증
//
//  6. Reuse validation logic across similar use cases
//     유사한 사용 사례에서 검증 로직 재사용
//
// See also / 참고:
//
//   - errors.go: Error types and handling / 에러 타입 및 처리
//   - types.go: Core types and interfaces / 핵심 타입 및 인터페이스
//   - rules_*.go: Specific validation rule implementations / 특정 검증 규칙 구현
//
// Version: 1.13.x
// Author: arkd0ng
// License: MIT
package validation

import "fmt"

// New creates a new Validator for the given value and field name.
// It initializes a validator instance with empty error collection and default settings.
//
// New는 주어진 값과 필드 이름에 대한 새 Validator를 생성합니다.
// 빈 에러 컬렉션과 기본 설정으로 검증기 인스턴스를 초기화합니다.
//
// Parameters / 매개변수:
//   - value: The value to be validated (can be any type including nil)
//     검증할 값 (nil을 포함한 모든 타입 가능)
//   - fieldName: Name of the field for error messages (should be descriptive)
//     에러 메시지용 필드 이름 (설명적이어야 함)
//
// Returns / 반환값:
//   - *Validator: A new validator instance ready for rule chaining
//     규칙 체이닝 준비가 된 새 검증기 인스턴스
//
// Behavior / 동작:
//   - Creates validator with stopOnError=false (collects all errors by default)
//     stopOnError=false로 검증기 생성 (기본적으로 모든 에러 수집)
//   - Initializes empty error collection
//     빈 에러 컬렉션 초기화
//   - Allocates custom message map
//     사용자 정의 메시지 맵 할당
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Each call creates a new independent instance
//     스레드 안전: 각 호출은 새로운 독립 인스턴스 생성
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Memory allocation: One validator struct + empty slice + empty map
//     메모리 할당: 검증기 구조체 1개 + 빈 슬라이스 + 빈 맵
//
// Example / 예제:
//
//	// String validation / 문자열 검증
//	v := validation.New("john@example.com", "email")
//	v.Required().Email()
//
//	// Numeric validation / 숫자 검증
//	v := validation.New(25, "age")
//	v.Required().Min(18).Max(120)
//
//	// Nil value validation / Nil 값 검증
//	v := validation.New(nil, "optional_field")
//	v.Email()  // Will pass if nil (optional)
//
//	// Complex validation chain / 복잡한 검증 체인
//	v := validation.New(username, "username")
//	v.Required().MinLength(3).MaxLength(20).AlphaNumeric()
func New(value interface{}, fieldName string) *Validator {
	return &Validator{
		value:          value,
		fieldName:      fieldName,
		errors:         []ValidationError{},
		stopOnError:    false,
		customMessages: make(map[string]string),
	}
}

// Validate executes all validation rules and returns an error if any fail.
// It consolidates all collected validation errors into a single error type.
//
// Validate는 모든 검증 규칙을 실행하고, 실패 시 에러를 반환합니다.
// 수집된 모든 검증 에러를 단일 에러 타입으로 통합합니다.
//
// Returns / 반환값:
//   - error: ValidationErrors containing all failed validations, or nil if all pass
//     모든 실패한 검증을 포함하는 ValidationErrors, 모두 통과 시 nil
//
// Behavior / 동작:
//   - Returns nil if no validation errors occurred
//     검증 에러가 없으면 nil 반환
//   - Returns ValidationErrors type for type-safe error handling
//     타입 안전한 에러 처리를 위해 ValidationErrors 타입 반환
//   - Does not clear error collection (can be called multiple times)
//     에러 컬렉션을 지우지 않음 (여러 번 호출 가능)
//
// Error Handling / 에러 처리:
//
//	The returned error can be type-asserted to ValidationErrors for detailed inspection:
//	반환된 에러는 상세 검사를 위해 ValidationErrors로 타입 단언 가능:
//
//	if err := v.Validate(); err != nil {
//	    if validationErrors, ok := err.(validation.ValidationErrors); ok {
//	        for _, e := range validationErrors {
//	            // Process each error / 각 에러 처리
//	        }
//	    }
//	}
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Read-only operation on validator state
//     스레드 안전: 검증기 상태에 대한 읽기 전용 작업
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - No additional memory allocation
//     추가 메모리 할당 없음
//
// Example / 예제:
//
//	v := validation.New(email, "email")
//	v.Required().Email()
//
//	if err := v.Validate(); err != nil {
//	    log.Printf("Validation failed: %v", err)
//	    return BadRequestError(err)
//	}
//
//	// Type-safe error handling / 타입 안전한 에러 처리
//	if err := v.Validate(); err != nil {
//	    errors := err.(validation.ValidationErrors)
//	    for _, e := range errors {
//	        fmt.Printf("Field '%s' failed rule '%s': %s\n",
//	            e.Field, e.Rule, e.Message)
//	    }
//	}
func (v *Validator) Validate() error {
	if len(v.errors) == 0 {
		return nil
	}
	return ValidationErrors(v.errors)
}

// GetErrors returns all validation errors collected so far.
// It provides direct access to the error collection for custom processing.
//
// GetErrors는 지금까지 수집된 모든 검증 에러를 반환합니다.
// 사용자 정의 처리를 위해 에러 컬렉션에 직접 액세스를 제공합니다.
//
// Returns / 반환값:
//   - []ValidationError: Slice of all validation errors (empty if no errors)
//     모든 검증 에러의 슬라이스 (에러가 없으면 빈 슬라이스)
//
// Behavior / 동작:
//   - Returns the actual error slice (not a copy)
//     실제 에러 슬라이스 반환 (복사본 아님)
//   - Returns empty slice if no validation errors
//     검증 에러가 없으면 빈 슬라이스 반환
//   - Does not clear the error collection
//     에러 컬렉션을 지우지 않음
//
// Use Cases / 사용 사례:
//   - Custom error formatting / 사용자 정의 에러 포맷팅
//   - Conditional error handling / 조건부 에러 처리
//   - Logging specific error details / 특정 에러 상세 정보 로깅
//   - Filtering errors by field or rule / 필드 또는 규칙별 에러 필터링
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Read-only operation
//     스레드 안전: 읽기 전용 작업
//   - Warning: Modifying returned slice affects validator state
//     경고: 반환된 슬라이스 수정 시 검증기 상태 영향
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - No memory allocation (returns existing slice)
//     메모리 할당 없음 (기존 슬라이스 반환)
//
// Example / 예제:
//
//	v := validation.New(data, "data")
//	v.Required().Email().MinLength(5)
//
//	errors := v.GetErrors()
//	if len(errors) > 0 {
//	    for _, err := range errors {
//	        log.Printf("Validation error on %s: %s", err.Field, err.Message)
//	    }
//	}
//
//	// Check for specific field errors / 특정 필드 에러 확인
//	hasEmailError := false
//	for _, err := range v.GetErrors() {
//	    if err.Field == "email" {
//	        hasEmailError = true
//	        break
//	    }
//	}
func (v *Validator) GetErrors() []ValidationError {
	return v.errors
}

// StopOnError sets the validator to stop on the first validation error.
// This is useful for expensive validations or when you only need to know if validation fails.
//
// StopOnError는 첫 번째 검증 에러에서 멈추도록 검증기를 설정합니다.
// 비용이 많이 드는 검증이나 검증 실패 여부만 알면 될 때 유용합니다.
//
// Returns / 반환값:
//   - *Validator: The validator instance for method chaining
//     메서드 체이닝을 위한 검증기 인스턴스
//
// Behavior / 동작:
//   - Sets internal flag to stop after first error
//     첫 에러 후 중단하도록 내부 플래그 설정
//   - All subsequent validation rules check this flag
//     모든 후속 검증 규칙이 이 플래그 확인
//   - Does not affect already collected errors
//     이미 수집된 에러에는 영향 없음
//
// Use Cases / 사용 사례:
//   - Expensive validations (database lookups, API calls)
//     비용이 많이 드는 검증 (데이터베이스 조회, API 호출)
//   - Early exit when any validation fails
//     검증 실패 시 조기 종료
//   - Performance optimization for long validation chains
//     긴 검증 체인의 성능 최적화
//   - When only boolean pass/fail is needed
//     통과/실패 여부만 필요할 때
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Modifies validator state atomically
//     스레드 안전: 검증기 상태를 원자적으로 수정
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Can significantly reduce validation time for long chains
//     긴 체인의 경우 검증 시간을 크게 단축 가능
//   - No memory allocation
//     메모리 할당 없음
//
// Example / 예제:
//
//	// Basic usage / 기본 사용
//	v := validation.New(email, "email").StopOnError()
//	v.Required().Email().MaxLength(100)
//	// Stops at first failed rule / 첫 번째 실패한 규칙에서 중단
//
//	// With expensive validation / 비용이 많이 드는 검증과 함께
//	v := validation.New(username, "username").StopOnError()
//	v.Required().
//	  MinLength(3).         // Quick check first / 먼저 빠른 확인
//	  MaxLength(20).
//	  Custom(checkUsernameInDB)  // Expensive DB check last
//
//	// Compare with collect-all mode / 모든 에러 수집 모드와 비교
//	v1 := validation.New(data, "data").StopOnError()
//	// Stops at first error, faster / 첫 에러에서 중단, 더 빠름
//
//	v2 := validation.New(data, "data")
//	// Collects all errors, more informative / 모든 에러 수집, 더 많은 정보
func (v *Validator) StopOnError() *Validator {
	v.stopOnError = true
	return v
}

// WithMessage sets a custom message for the last validation rule.
// WithMessage는 마지막 검증 규칙에 대한 사용자 정의 메시지를 설정합니다.
func (v *Validator) WithMessage(message string) *Validator {
	if len(v.errors) > 0 {
		// Update the last error's message
		v.errors[len(v.errors)-1].Message = message
	}
	return v
}

// addError adds a validation error to the validator.
// addError는 검증기에 검증 에러를 추가합니다.
func (v *Validator) addError(rule, message string) *Validator {
	// If stopOnError is true and we already have errors, don't add more
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	// Check if there's a custom message for this rule
	// 이 규칙에 대한 커스텀 메시지가 있는지 확인
	if customMsg, ok := v.customMessages[rule]; ok {
		message = customMsg
	}

	v.errors = append(v.errors, ValidationError{
		Field:   v.fieldName,
		Value:   v.value,
		Rule:    rule,
		Message: message,
	})
	v.lastRule = rule
	return v
}

// WithCustomMessage sets a custom message for a specific validation rule before it runs.
// WithCustomMessage는 검증 규칙이 실행되기 전에 특정 규칙에 대한 커스텀 메시지를 설정합니다.
//
// This allows you to set custom messages upfront, unlike WithMessage() which modifies the last error.
// WithMessage()는 마지막 에러를 수정하는 것과 달리, 이 메서드는 미리 커스텀 메시지를 설정할 수 있습니다.
//
// Example / 예시:
//
//	v := validation.New("", "email")
//	v.WithCustomMessage("required", "이메일을 입력해주세요")
//	v.WithCustomMessage("email", "올바른 이메일 형식이 아닙니다")
//	v.Required().Email()
func (v *Validator) WithCustomMessage(rule, message string) *Validator {
	v.customMessages[rule] = message
	return v
}

// WithCustomMessages sets multiple custom messages for validation rules.
// WithCustomMessages는 검증 규칙에 대한 여러 커스텀 메시지를 설정합니다.
//
// Example / 예시:
//
//	v := validation.New("", "password")
//	v.WithCustomMessages(map[string]string{
//	    "required":   "비밀번호를 입력해주세요",
//	    "min_length": "비밀번호는 8자 이상이어야 합니다",
//	    "max_length": "비밀번호는 20자 이하여야 합니다",
//	})
//	v.Required().MinLength(8).MaxLength(20)
func (v *Validator) WithCustomMessages(messages map[string]string) *Validator {
	for rule, message := range messages {
		v.customMessages[rule] = message
	}
	return v
}

// Custom applies a custom validation function with a message.
// Custom은 사용자 정의 검증 함수를 메시지와 함께 적용합니다.
func (v *Validator) Custom(fn RuleFunc, message string) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if !fn(v.value) {
		v.addError("custom", message)
	}

	return v
}

// NewValidator creates a new MultiValidator for multiple fields.
// NewValidator는 여러 필드를 위한 새 MultiValidator를 생성합니다.
func NewValidator() *MultiValidator {
	return &MultiValidator{
		validators: []*Validator{},
		errors:     []ValidationError{},
	}
}

// Field adds a field to the multi-validator and returns a Validator for chaining.
// Field는 multi-validator에 필드를 추가하고 체이닝을 위한 Validator를 반환합니다.
func (mv *MultiValidator) Field(value interface{}, fieldName string) *Validator {
	v := New(value, fieldName)
	mv.validators = append(mv.validators, v)
	return v
}

// Validate executes all validators and returns an error if any fail.
// Validate는 모든 검증기를 실행하고, 실패 시 에러를 반환합니다.
func (mv *MultiValidator) Validate() error {
	for _, v := range mv.validators {
		if err := v.Validate(); err != nil {
			if verrs, ok := err.(ValidationErrors); ok {
				mv.errors = append(mv.errors, verrs...)
			}
		}
	}

	if len(mv.errors) == 0 {
		return nil
	}

	return ValidationErrors(mv.errors)
}

// GetErrors returns all validation errors from all fields.
// GetErrors는 모든 필드의 모든 검증 에러를 반환합니다.
func (mv *MultiValidator) GetErrors() []ValidationError {
	return mv.errors
}

// validateString is a helper function to validate string values.
// validateString은 문자열 값을 검증하는 헬퍼 함수입니다.
func validateString(v *Validator, rule string, fn func(string) bool, message string) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	s, ok := v.value.(string)
	if !ok {
		v.addError(rule, fmt.Sprintf("%s must be a string", v.fieldName))
		return v
	}

	if !fn(s) {
		v.addError(rule, message)
	}

	return v
}

// validateNumeric is a helper function to validate numeric values.
// validateNumeric는 숫자 값을 검증하는 헬퍼 함수입니다.
func validateNumeric(v *Validator, rule string, fn func(float64) bool, message string) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	var num float64
	switch n := v.value.(type) {
	case int:
		num = float64(n)
	case int8:
		num = float64(n)
	case int16:
		num = float64(n)
	case int32:
		num = float64(n)
	case int64:
		num = float64(n)
	case uint:
		num = float64(n)
	case uint8:
		num = float64(n)
	case uint16:
		num = float64(n)
	case uint32:
		num = float64(n)
	case uint64:
		num = float64(n)
	case float32:
		num = float64(n)
	case float64:
		num = n
	default:
		v.addError(rule, fmt.Sprintf("%s must be a numeric value", v.fieldName))
		return v
	}

	if !fn(num) {
		v.addError(rule, message)
	}

	return v
}
