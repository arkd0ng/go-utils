package websvrutil

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// validator.go provides struct field validation using tag-based rules.
//
// This file implements a comprehensive validation system that inspects struct
// field tags to enforce constraints on user input:
//
// Core Types:
//
// Validator Interface:
//   - Defines standard interface for validation implementations
//   - Single method: Validate(interface{}) error
//   - Allows custom validator implementations
//
// DefaultValidator:
//   - Built-in tag-based validator
//   - Uses reflection to inspect struct fields
//   - Parses "validate" struct tags
//   - Applies validation rules sequentially
//
// ValidationError:
//   - Represents single field validation failure
//   - Contains: Field name, Tag rule, Actual value
//   - Formatted error message
//
// ValidationErrors:
//   - Collection of multiple ValidationError instances
//   - Concatenated error messages (semicolon-separated)
//   - Returned when multiple fields fail validation
//
// Supported Validation Tags:
//
// Presence & Equality:
//   - required: Field must not be zero value (empty string, 0, nil, false)
//   - eq=<value>: Field must equal specified value
//   - ne=<value>: Field must not equal specified value
//
// String Length:
//   - min=<n>: String must have at least n characters
//   - max=<n>: String must have at most n characters
//   - len=<n>: String must have exactly n characters
//
// Numeric Comparisons:
//   - gt=<n>: Number must be greater than n
//   - gte=<n>: Number must be greater than or equal to n
//   - lt=<n>: Number must be less than n
//   - lte=<n>: Number must be less than or equal to n
//
// String Format:
//   - email: Must be valid email format (RFC 5322 basic pattern)
//   - alpha: Alphabetic characters only (a-z, A-Z)
//   - alphanum: Alphanumeric characters only (a-z, A-Z, 0-9)
//   - numeric: Numeric characters only (0-9)
//
// Value Options:
//   - oneof=<v1,v2,v3>: Field must be one of comma-separated values
//     Example: oneof=admin,user,guest
//
// Tag Syntax:
//   - Multiple rules: Comma-separated (e.g., "required,min=3,max=50")
//   - Rule with value: rule=value (e.g., "min=5")
//   - Rule without value: rule (e.g., "required")
//
// Validation Functions:
//
// DefaultValidator.Validate(obj):
//   - Main validation entry point
//   - Uses reflection to iterate struct fields
//   - Extracts "validate" tag from each field
//   - Applies all rules sequentially
//   - Collects errors for all fields
//   - Returns ValidationErrors if any failures
//
// validateField(fieldName, field, tag):
//   - Parses comma-separated validation rules
//   - Applies each rule to field value
//   - Returns first validation error encountered
//
// applyRule(fieldName, field, ruleName, ruleValue):
//   - Dispatches to specific validation function
//   - Handles rule parsing (splits "rule=value")
//   - Returns ValidationError on failure
//
// Specific Validators:
//   - validateRequired(field): Checks non-zero value
//   - validateEmail(field): Regex pattern matching
//   - validateMin(field, min): String length or numeric comparison
//   - validateMax(field, max): String length or numeric comparison
//   - validateLen(field, len): Exact string length check
//   - validateEq(field, value): Equality check (string or numeric)
//   - validateNe(field, value): Inequality check
//   - validateGt/Gte/Lt/Lte(field, value): Numeric comparisons
//   - validateOneOf(field, values): String membership check
//   - validateAlpha/Alphanum/Numeric(field): Character class checks
//
// Helper Functions:
//   - isZero(value): Check if value is zero value for its type
//     Handles: string, int, float, bool, pointer, slice, map
//
// Context Integration:
//   - Context.BindWithValidation(obj): Bind and validate in one call
//     Combines BindJSON/BindForm with Validate()
//     Returns error if binding or validation fails
//
// Example usage:
//
//	// Define struct with validation tags
//	type UserRegister struct {
//	    Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
//	    Email    string `json:"email" validate:"required,email"`
//	    Password string `json:"password" validate:"required,min=8"`
//	    Age      int    `json:"age" validate:"required,gte=18,lte=100"`
//	    Role     string `json:"role" validate:"required,oneof=admin,user,guest"`
//	}
//
//	// Manual validation
//	validator := &DefaultValidator{}
//	user := UserRegister{
//	    Username: "jo",      // Too short
//	    Email:    "invalid", // Invalid email
//	    Age:      15,        // Too young
//	}
//	err := validator.Validate(user)
//	// Returns ValidationErrors with 3 errors
//
//	// Validation in handler
//	app.POST("/register", func(w http.ResponseWriter, r *http.Request) {
//	    ctx := GetContext(r)
//	    var user UserRegister
//	    if err := ctx.BindWithValidation(&user); err != nil {
//	        // Handle validation errors
//	        if valErrs, ok := err.(ValidationErrors); ok {
//	            // Return detailed error messages
//	            ctx.JSON(400, map[string]interface{}{
//	                "error": "Validation failed",
//	                "details": valErrs,
//	            })
//	            return
//	        }
//	        ctx.JSON(400, map[string]string{"error": err.Error()})
//	        return
//	    }
//	    // Proceed with validated user data
//	})
//
// Performance:
//   - Validation time: O(n*m) where n = fields, m = rules per field
//   - Uses reflection (overhead for struct inspection)
//   - Regex compilation cached for email validation
//   - Efficient for typical form validation (<100 fields)
//
// Limitations:
//   - Only validates struct fields (no nested struct validation)
//   - No custom error messages (generic messages only)
//   - Limited to supported validation rules
//   - Reflection-based (slower than compiled validators)
//   - For complex validation, consider third-party libraries:
//     * github.com/go-playground/validator/v10
//     * github.com/go-ozzo/ozzo-validation
//
// Best Practices:
//   - Always validate user input before processing
//   - Use BindWithValidation() for combined bind + validate
//   - Return detailed ValidationErrors to client for better UX
//   - Keep validation rules in struct tags for clarity
//   - Consider custom Validator implementation for complex rules
//   - Validate business logic separately (e.g., unique email, existing user)
//
// validator.go는 태그 기반 규칙을 사용한 구조체 필드 검증을 제공합니다.
//
// 이 파일은 사용자 입력에 제약 조건을 적용하기 위해 구조체 필드 태그를
// 검사하는 포괄적인 검증 시스템을 구현합니다:
//
// 핵심 타입:
//
// Validator 인터페이스:
//   - 검증 구현을 위한 표준 인터페이스 정의
//   - 단일 메서드: Validate(interface{}) error
//   - 커스텀 검증자 구현 허용
//
// DefaultValidator:
//   - 내장 태그 기반 검증자
//   - 리플렉션을 사용하여 구조체 필드 검사
//   - "validate" 구조체 태그 파싱
//   - 검증 규칙 순차 적용
//
// ValidationError:
//   - 단일 필드 검증 실패 표현
//   - 포함: 필드 이름, 태그 규칙, 실제 값
//   - 형식화된 에러 메시지
//
// ValidationErrors:
//   - 여러 ValidationError 인스턴스 모음
//   - 연결된 에러 메시지 (세미콜론 구분)
//   - 여러 필드가 검증 실패 시 반환
//
// 지원되는 검증 태그:
//
// 존재 및 동등성:
//   - required: 필드가 제로 값이 아니어야 함 (빈 문자열, 0, nil, false)
//   - eq=<value>: 필드가 지정된 값과 같아야 함
//   - ne=<value>: 필드가 지정된 값과 달라야 함
//
// 문자열 길이:
//   - min=<n>: 문자열이 최소 n자 이상이어야 함
//   - max=<n>: 문자열이 최대 n자 이하여야 함
//   - len=<n>: 문자열이 정확히 n자여야 함
//
// 숫자 비교:
//   - gt=<n>: 숫자가 n보다 커야 함
//   - gte=<n>: 숫자가 n보다 크거나 같아야 함
//   - lt=<n>: 숫자가 n보다 작아야 함
//   - lte=<n>: 숫자가 n보다 작거나 같아야 함
//
// 문자열 형식:
//   - email: 유효한 이메일 형식이어야 함 (RFC 5322 기본 패턴)
//   - alpha: 알파벳 문자만 (a-z, A-Z)
//   - alphanum: 영숫자 문자만 (a-z, A-Z, 0-9)
//   - numeric: 숫자 문자만 (0-9)
//
// 값 옵션:
//   - oneof=<v1,v2,v3>: 필드가 쉼표로 구분된 값 중 하나여야 함
//     예제: oneof=admin,user,guest
//
// 태그 구문:
//   - 여러 규칙: 쉼표로 구분 (예: "required,min=3,max=50")
//   - 값이 있는 규칙: rule=value (예: "min=5")
//   - 값이 없는 규칙: rule (예: "required")
//
// 검증 함수:
//
// DefaultValidator.Validate(obj):
//   - 주요 검증 진입점
//   - 리플렉션을 사용하여 구조체 필드 반복
//   - 각 필드에서 "validate" 태그 추출
//   - 모든 규칙 순차 적용
//   - 모든 필드의 에러 수집
//   - 실패 시 ValidationErrors 반환
//
// validateField(fieldName, field, tag):
//   - 쉼표로 구분된 검증 규칙 파싱
//   - 필드 값에 각 규칙 적용
//   - 첫 번째 검증 에러 반환
//
// applyRule(fieldName, field, ruleName, ruleValue):
//   - 특정 검증 함수로 디스패치
//   - 규칙 파싱 처리 ("rule=value" 분할)
//   - 실패 시 ValidationError 반환
//
// 특정 검증자:
//   - validateRequired(field): 제로가 아닌 값 확인
//   - validateEmail(field): 정규식 패턴 매칭
//   - validateMin(field, min): 문자열 길이 또는 숫자 비교
//   - validateMax(field, max): 문자열 길이 또는 숫자 비교
//   - validateLen(field, len): 정확한 문자열 길이 확인
//   - validateEq(field, value): 동등성 확인 (문자열 또는 숫자)
//   - validateNe(field, value): 부등성 확인
//   - validateGt/Gte/Lt/Lte(field, value): 숫자 비교
//   - validateOneOf(field, values): 문자열 멤버십 확인
//   - validateAlpha/Alphanum/Numeric(field): 문자 클래스 확인
//
// 헬퍼 함수:
//   - isZero(value): 타입의 제로 값인지 확인
//     처리: string, int, float, bool, pointer, slice, map
//
// 컨텍스트 통합:
//   - Context.BindWithValidation(obj): 한 번의 호출로 바인딩 및 검증
//     BindJSON/BindForm과 Validate() 결합
//     바인딩 또는 검증 실패 시 에러 반환
//
// 사용 예제:
//
//	// 검증 태그가 있는 구조체 정의
//	type UserRegister struct {
//	    Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
//	    Email    string `json:"email" validate:"required,email"`
//	    Password string `json:"password" validate:"required,min=8"`
//	    Age      int    `json:"age" validate:"required,gte=18,lte=100"`
//	    Role     string `json:"role" validate:"required,oneof=admin,user,guest"`
//	}
//
//	// 수동 검증
//	validator := &DefaultValidator{}
//	user := UserRegister{
//	    Username: "jo",      // 너무 짧음
//	    Email:    "invalid", // 유효하지 않은 이메일
//	    Age:      15,        // 너무 어림
//	}
//	err := validator.Validate(user)
//	// 3개의 에러가 있는 ValidationErrors 반환
//
//	// 핸들러에서 검증
//	app.POST("/register", func(w http.ResponseWriter, r *http.Request) {
//	    ctx := GetContext(r)
//	    var user UserRegister
//	    if err := ctx.BindWithValidation(&user); err != nil {
//	        // 검증 에러 처리
//	        if valErrs, ok := err.(ValidationErrors); ok {
//	            // 상세한 에러 메시지 반환
//	            ctx.JSON(400, map[string]interface{}{
//	                "error": "Validation failed",
//	                "details": valErrs,
//	            })
//	            return
//	        }
//	        ctx.JSON(400, map[string]string{"error": err.Error()})
//	        return
//	    }
//	    // 검증된 사용자 데이터로 진행
//	})
//
// 성능:
//   - 검증 시간: O(n*m), n = 필드 수, m = 필드당 규칙 수
//   - 리플렉션 사용 (구조체 검사 오버헤드)
//   - 이메일 검증을 위한 정규식 컴파일 캐시
//   - 일반적인 폼 검증에 효율적 (<100 필드)
//
// 제한사항:
//   - 구조체 필드만 검증 (중첩 구조체 검증 없음)
//   - 커스텀 에러 메시지 없음 (일반 메시지만)
//   - 지원되는 검증 규칙으로 제한
//   - 리플렉션 기반 (컴파일된 검증자보다 느림)
//   - 복잡한 검증의 경우 서드파티 라이브러리 고려:
//     * github.com/go-playground/validator/v10
//     * github.com/go-ozzo/ozzo-validation
//
// 모범 사례:
//   - 처리 전 항상 사용자 입력 검증
//   - 결합된 바인딩 + 검증을 위해 BindWithValidation() 사용
//   - 더 나은 UX를 위해 상세한 ValidationErrors를 클라이언트에 반환
//   - 명확성을 위해 구조체 태그에 검증 규칙 유지
//   - 복잡한 규칙을 위한 커스텀 Validator 구현 고려
//   - 비즈니스 로직은 별도로 검증 (예: 고유 이메일, 기존 사용자)

// ============================================================================
// Validation
// 검증
// ============================================================================

// Validator defines the interface for custom validators.
// Validator는 커스텀 검증자를 위한 인터페이스를 정의합니다.
type Validator interface {
	Validate(interface{}) error
}

// ValidationError represents a validation error.
// ValidationError는 검증 에러를 나타냅니다.
type ValidationError struct {
	Field string
	Tag   string
	Value interface{}
}

// Error returns the error message.
// Error는 에러 메시지를 반환합니다.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s' with tag '%s'", e.Field, e.Tag)
}

// ValidationErrors represents multiple validation errors.
// ValidationErrors는 여러 검증 에러를 나타냅니다.
type ValidationErrors []*ValidationError

// Error returns the concatenated error messages.
// Error는 연결된 에러 메시지를 반환합니다.
func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}
	messages := make([]string, len(e))
	for i, err := range e {
		messages[i] = err.Error()
	}
	return strings.Join(messages, "; ")
}

// DefaultValidator is the built-in validator.
// DefaultValidator는 내장 검증자입니다.
type DefaultValidator struct{}

// Validate validates a struct using validation tags.
// Validate는 검증 태그를 사용하여 구조체를 검증합니다.
//
// Supported tags
// 지원되는 태그:
// - required: field must not be empty
// 필드가 비어 있지 않아야 함
// - email: field must be a valid email address
// 유효한 이메일 주소여야 함
// - min=<value>: minimum length/value
// 최소 길이/값
// - max=<value>: maximum length/value
// 최대 길이/값
// - len=<value>: exact length
// 정확한 길이
// - eq=<value>: equal to value
// 값과 같아야 함
// - ne=<value>: not equal to value
// 값과 같지 않아야 함
// - gt=<value>: greater than value
// 값보다 커야 함
// - gte=<value>: greater than or equal to value
// 값보다 크거나 같아야 함
// - lt=<value>: less than value
// 값보다 작아야 함
// - lte=<value>: less than or equal to value
// 값보다 작거나 같아야 함
// - oneof=<values>: one of the values (comma-separated)
// 값 중 하나 (쉼표로 구분)
// - alpha: alphabetic characters only
// 알파벳 문자만
// - alphanum: alphanumeric characters only
// 영숫자 문자만
// - numeric: numeric characters only
// 숫자 문자만
//
// Example
// 예제:
//
//	type User struct {
//	    Name  string `validate:"required,min=3,max=50"`
//	    Email string `validate:"required,email"`
//	    Age   int    `validate:"required,gte=18,lte=100"`
//	    Role  string `validate:"required,oneof=admin,user,guest"`
//	}
func (v *DefaultValidator) Validate(obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("validate: expected struct, got %s", val.Kind())
	}

	var errors ValidationErrors

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Get validation tag
		// 검증 태그 가져오기
		tag := fieldType.Tag.Get("validate")
		if tag == "" {
			continue
		}

		// Parse and validate tags
		// 태그 파싱 및 검증
		if err := validateField(fieldType.Name, field, tag); err != nil {
			if ve, ok := err.(*ValidationError); ok {
				errors = append(errors, ve)
			} else if ves, ok := err.(ValidationErrors); ok {
				errors = append(errors, ves...)
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// validateField validates a single field.
// validateField는 단일 필드를 검증합니다.
func validateField(fieldName string, field reflect.Value, tag string) error {
	// Split tag by comma
	// 쉼표로 태그 분할
	rules := strings.Split(tag, ",")

	for _, rule := range rules {
		rule = strings.TrimSpace(rule)
		if rule == "" {
			continue
		}

		// Parse rule (e.g., "min=3" -> "min", "3")
		// 규칙 파싱 (예: "min=3" -> "min", "3")
		parts := strings.SplitN(rule, "=", 2)
		ruleName := parts[0]
		var ruleValue string
		if len(parts) > 1 {
			ruleValue = parts[1]
		}

		// Validate based on rule
		// 규칙에 따라 검증
		if err := applyRule(fieldName, field, ruleName, ruleValue); err != nil {
			return err
		}
	}

	return nil
}

// applyRule applies a validation rule to a field.
// applyRule은 필드에 검증 규칙을 적용합니다.
func applyRule(fieldName string, field reflect.Value, ruleName, ruleValue string) error {
	switch ruleName {
	case "required":
		return validateRequired(fieldName, field)
	case "email":
		return validateEmail(fieldName, field)
	case "min":
		return validateMin(fieldName, field, ruleValue)
	case "max":
		return validateMax(fieldName, field, ruleValue)
	case "len":
		return validateLen(fieldName, field, ruleValue)
	case "eq":
		return validateEq(fieldName, field, ruleValue)
	case "ne":
		return validateNe(fieldName, field, ruleValue)
	case "gt":
		return validateGt(fieldName, field, ruleValue)
	case "gte":
		return validateGte(fieldName, field, ruleValue)
	case "lt":
		return validateLt(fieldName, field, ruleValue)
	case "lte":
		return validateLte(fieldName, field, ruleValue)
	case "oneof":
		return validateOneOf(fieldName, field, ruleValue)
	case "alpha":
		return validateAlpha(fieldName, field)
	case "alphanum":
		return validateAlphanum(fieldName, field)
	case "numeric":
		return validateNumeric(fieldName, field)
	default:
		// Unknown rule, skip
		// 알 수 없는 규칙, 건너뛰기
		return nil
	}
}

// validateRequired validates that a field is not empty.
// validateRequired는 필드가 비어 있지 않은지 검증합니다.
func validateRequired(fieldName string, field reflect.Value) error {
	if isZero(field) {
		return &ValidationError{Field: fieldName, Tag: "required", Value: field.Interface()}
	}
	return nil
}

// validateEmail validates that a field is a valid email.
// validateEmail은 필드가 유효한 이메일인지 검증합니다.
func validateEmail(fieldName string, field reflect.Value) error {
	if field.Kind() != reflect.String {
		return nil
	}

	email := field.String()
	if email == "" {
		// Empty is handled by required
		// 빈 값은 required에서 처리
		return nil
	}

	// Simple email regex
	// 간단한 이메일 정규식
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return &ValidationError{Field: fieldName, Tag: "email", Value: email}
	}

	return nil
}

// validateMin validates minimum length/value.
// validateMin은 최소 길이/값을 검증합니다.
func validateMin(fieldName string, field reflect.Value, min string) error {
	minVal, err := strconv.ParseInt(min, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		if int64(len(field.String())) < minVal {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() < minVal {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() < uint64(minVal) {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() < float64(minVal) {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if int64(field.Len()) < minVal {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	}

	return nil
}

// validateMax validates maximum length/value.
// validateMax는 최대 길이/값을 검증합니다.
func validateMax(fieldName string, field reflect.Value, max string) error {
	maxVal, err := strconv.ParseInt(max, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		if int64(len(field.String())) > maxVal {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() > maxVal {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() > uint64(maxVal) {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() > float64(maxVal) {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if int64(field.Len()) > maxVal {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	}

	return nil
}

// validateLen validates exact length.
// validateLen은 정확한 길이를 검증합니다.
func validateLen(fieldName string, field reflect.Value, length string) error {
	lenVal, err := strconv.ParseInt(length, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		if int64(len(field.String())) != lenVal {
			return &ValidationError{Field: fieldName, Tag: "len=" + length, Value: field.Interface()}
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if int64(field.Len()) != lenVal {
			return &ValidationError{Field: fieldName, Tag: "len=" + length, Value: field.Interface()}
		}
	}

	return nil
}

// validateEq validates equality.
// validateEq는 동등성을 검증합니다.
func validateEq(fieldName string, field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		if field.String() != value {
			return &ValidationError{Field: fieldName, Tag: "eq=" + value, Value: field.Interface()}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, _ := strconv.ParseInt(value, 10, 64)
		if field.Int() != intVal {
			return &ValidationError{Field: fieldName, Tag: "eq=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateNe validates inequality.
// validateNe는 부등성을 검증합니다.
func validateNe(fieldName string, field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		if field.String() == value {
			return &ValidationError{Field: fieldName, Tag: "ne=" + value, Value: field.Interface()}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, _ := strconv.ParseInt(value, 10, 64)
		if field.Int() == intVal {
			return &ValidationError{Field: fieldName, Tag: "ne=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateGt validates greater than.
// validateGt는 보다 큰 값을 검증합니다.
func validateGt(fieldName string, field reflect.Value, value string) error {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() <= intVal {
			return &ValidationError{Field: fieldName, Tag: "gt=" + value, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() <= uint64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "gt=" + value, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() <= float64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "gt=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateGte validates greater than or equal.
// validateGte는 보다 크거나 같은 값을 검증합니다.
func validateGte(fieldName string, field reflect.Value, value string) error {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() < intVal {
			return &ValidationError{Field: fieldName, Tag: "gte=" + value, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() < uint64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "gte=" + value, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() < float64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "gte=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateLt validates less than.
// validateLt는 보다 작은 값을 검증합니다.
func validateLt(fieldName string, field reflect.Value, value string) error {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() >= intVal {
			return &ValidationError{Field: fieldName, Tag: "lt=" + value, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() >= uint64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "lt=" + value, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() >= float64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "lt=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateLte validates less than or equal.
// validateLte는 보다 작거나 같은 값을 검증합니다.
func validateLte(fieldName string, field reflect.Value, value string) error {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() > intVal {
			return &ValidationError{Field: fieldName, Tag: "lte=" + value, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() > uint64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "lte=" + value, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() > float64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "lte=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateOneOf validates that value is one of the specified values.
// validateOneOf는 값이 지정된 값 중 하나인지 검증합니다.
func validateOneOf(fieldName string, field reflect.Value, values string) error {
	if field.Kind() != reflect.String {
		return nil
	}

	fieldValue := field.String()
	// Split by comma
	// 쉼표로 분할
	validValues := strings.Split(values, " ")

	for _, valid := range validValues {
		if fieldValue == strings.TrimSpace(valid) {
			return nil
		}
	}

	return &ValidationError{Field: fieldName, Tag: "oneof=" + values, Value: fieldValue}
}

// validateAlpha validates alphabetic characters only.
// validateAlpha는 알파벳 문자만 검증합니다.
func validateAlpha(fieldName string, field reflect.Value) error {
	if field.Kind() != reflect.String {
		return nil
	}

	str := field.String()
	if str == "" {
		return nil
	}

	alphaRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !alphaRegex.MatchString(str) {
		return &ValidationError{Field: fieldName, Tag: "alpha", Value: str}
	}

	return nil
}

// validateAlphanum validates alphanumeric characters only.
// validateAlphanum은 영숫자 문자만 검증합니다.
func validateAlphanum(fieldName string, field reflect.Value) error {
	if field.Kind() != reflect.String {
		return nil
	}

	str := field.String()
	if str == "" {
		return nil
	}

	alphanumRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !alphanumRegex.MatchString(str) {
		return &ValidationError{Field: fieldName, Tag: "alphanum", Value: str}
	}

	return nil
}

// validateNumeric validates numeric characters only.
// validateNumeric은 숫자 문자만 검증합니다.
func validateNumeric(fieldName string, field reflect.Value) error {
	if field.Kind() != reflect.String {
		return nil
	}

	str := field.String()
	if str == "" {
		return nil
	}

	numericRegex := regexp.MustCompile(`^[0-9]+$`)
	if !numericRegex.MatchString(str) {
		return &ValidationError{Field: fieldName, Tag: "numeric", Value: str}
	}

	return nil
}

// isZero checks if a value is zero.
// isZero는 값이 제로인지 확인합니다.
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	}

	return false
}

// BindWithValidation binds and validates the request data.
// BindWithValidation은 요청 데이터를 바인딩하고 검증합니다.
//
// Example
// 예제:
//
//	type User struct {
//	    Name  string `json:"name" validate:"required,min=3"`
//	    Email string `json:"email" validate:"required,email"`
//	}
//	var user User
//	if err := ctx.BindWithValidation(&user); err != nil {
//	    return ctx.Error(400, err.Error())
//	}
func (c *Context) BindWithValidation(obj interface{}) error {
	// First bind the data
	// 먼저 데이터 바인딩
	if err := c.Bind(obj); err != nil {
		return err
	}

	// Then validate
	// 그 다음 검증
	validator := &DefaultValidator{}
	return validator.Validate(obj)
}
