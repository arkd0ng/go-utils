package validation

import (
	"fmt"
	"strings"
)

// ValidationError represents a single validation error with detailed context.
// It contains all information needed to understand and handle a validation failure.
//
// ValidationError는 상세한 컨텍스트를 포함한 단일 검증 에러를 나타냅니다.
// 검증 실패를 이해하고 처리하는 데 필요한 모든 정보를 포함합니다.
//
// Fields / 필드:
//   - Field: Name of the field that failed (e.g., "email", "age")
//     실패한 필드 이름 (예: "email", "age")
//   - Value: The actual value that failed validation (any type)
//     검증 실패한 실제 값 (모든 타입)
//   - Rule: Name of the failed validation rule (e.g., "required", "email")
//     실패한 검증 규칙 이름 (예: "required", "email")
//   - Message: Human-readable error message (bilingual supported)
//     사람이 읽을 수 있는 에러 메시지 (이중 언어 지원)
//
// Use Cases / 사용 사례:
//   - API error responses with detailed field errors
//     상세한 필드 에러가 있는 API 에러 응답
//   - Form validation feedback to users
//     사용자에게 폼 검증 피드백
//   - Logging validation failures with context
//     컨텍스트와 함께 검증 실패 로깅
//   - Conditional error handling based on field or rule
//     필드 또는 규칙 기반 조건부 에러 처리
//
// Thread Safety / 스레드 안전성:
//   - Immutable after creation
//     생성 후 불변
//   - Safe to read concurrently
//     동시 읽기 안전
//
// Example / 예제:
//   err := ValidationError{
//       Field:   "email",
//       Value:   "invalid-email",
//       Rule:    "email",
//       Message: "Invalid email format",
//   }
//   
//   // Access error details / 에러 상세 정보 액세스
//   fmt.Printf("Field: %s\n", err.Field)
//   fmt.Printf("Rule: %s\n", err.Rule)
//   fmt.Printf("Message: %s\n", err.Error())
type ValidationError struct {
	Field   string      // Field name / 필드 이름
	Value   interface{} // Invalid value / 유효하지 않은 값
	Rule    string      // Failed rule / 실패한 규칙
	Message string      // Error message / 에러 메시지
}

// Error returns the error message for this validation error.
// It implements the error interface for standard Go error handling.
//
// Error는 이 검증 에러에 대한 에러 메시지를 반환합니다.
// 표준 Go 에러 처리를 위해 error 인터페이스를 구현합니다.
//
// Returns / 반환값:
//   - string: Custom message if set, or default formatted message
//     설정된 경우 사용자 정의 메시지, 그렇지 않으면 기본 포맷된 메시지
//
// Behavior / 동작:
//   - Returns custom message if Message field is non-empty
//     Message 필드가 비어있지 않으면 사용자 정의 메시지 반환
//   - Returns default format: "{Rule} validation failed for field '{Field}'"
//     기본 형식 반환: "{규칙} validation failed for field '{필드}'"
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Read-only operation
//     스레드 안전: 읽기 전용 작업
//
// Example / 예제:
//   err := ValidationError{
//       Field:   "age",
//       Value:   -5,
//       Rule:    "min",
//       Message: "Age must be at least 0",
//   }
//   fmt.Println(err.Error())  // "Age must be at least 0"
//   
//   // Without custom message / 사용자 정의 메시지 없이
//   err2 := ValidationError{
//       Field: "email",
//       Rule:  "required",
//   }
//   fmt.Println(err2.Error())  // "required validation failed for field 'email'"
func (ve ValidationError) Error() string {
	if ve.Message != "" {
		return ve.Message
	}
	return fmt.Sprintf("%s validation failed for field '%s'", ve.Rule, ve.Field)
}

// ValidationErrors is a collection of validation errors.
// It implements the error interface and provides utility methods for error inspection.
//
// ValidationErrors는 검증 에러의 모음입니다.
// error 인터페이스를 구현하고 에러 검사를 위한 유틸리티 메서드를 제공합니다.
//
// Behavior / 동작:
//   - Acts as a slice of ValidationError
//     ValidationError의 슬라이스로 동작
//   - Can be empty (no errors)
//     비어있을 수 있음 (에러 없음)
//   - Provides formatted output of all errors
//     모든 에러의 포맷된 출력 제공
//
// Use Cases / 사용 사례:
//   - Collecting multiple field validation errors
//     여러 필드 검증 에러 수집
//   - Returning comprehensive validation feedback
//     포괄적인 검증 피드백 반환
//   - Filtering errors by field or rule
//     필드 또는 규칙별 에러 필터링
//   - API error responses with multiple violations
//     여러 위반 사항이 있는 API 에러 응답
//
// Thread Safety / 스레드 안전성:
//   - Safe to read concurrently
//     동시 읽기 안전
//   - Not safe to modify concurrently
//     동시 수정 안전하지 않음
//
// Example / 예제:
//   errors := validation.ValidationErrors{
//       {Field: "email", Rule: "required", Message: "Email is required"},
//       {Field: "age", Rule: "min", Message: "Age must be at least 18"},
//   }
//   
//   // Use as error / 에러로 사용
//   if len(errors) > 0 {
//       return errors  // Implements error interface
//   }
//   
//   // Iterate through errors / 에러 순회
//   for _, err := range errors {
//       log.Printf("Field %s failed: %s", err.Field, err.Message)
//   }
type ValidationErrors []ValidationError

// Error returns a formatted error message for all validation errors.
// Multiple errors are separated by semicolons for easy parsing.
//
// Error는 모든 검증 에러에 대한 포맷된 에러 메시지를 반환합니다.
// 여러 에러는 쉬운 파싱을 위해 세미콜론으로 구분됩니다.
//
// Returns / 반환값:
//   - string: Semicolon-separated error messages, or empty string if no errors
//     세미콜론으로 구분된 에러 메시지, 에러가 없으면 빈 문자열
//
// Format / 형식:
//   - Single error: "error message"
//     단일 에러: "에러 메시지"
//   - Multiple errors: "error1; error2; error3"
//     여러 에러: "에러1; 에러2; 에러3"
//   - No errors: ""
//     에러 없음: ""
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Read-only operation
//     스레드 안전: 읽기 전용 작업
//
// Performance / 성능:
//   - Time complexity: O(n) where n is number of errors
//     시간 복잡도: O(n) (n은 에러 개수)
//   - Memory allocation: One string allocation for join
//     메모리 할당: join을 위한 문자열 할당 1회
//
// Example / 예제:
//   errors := validation.ValidationErrors{
//       {Field: "email", Rule: "required", Message: "Email is required"},
//       {Field: "age", Rule: "min", Message: "Age too low"},
//   }
//   
//   msg := errors.Error()
//   // "Email is required; Age too low"
//   
//   // Empty errors / 빈 에러
//   empty := validation.ValidationErrors{}
//   fmt.Println(empty.Error())  // ""
func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return ""
	}

	var messages []string
	for _, err := range ve {
		messages = append(messages, err.Error())
	}

	return strings.Join(messages, "; ")
}

// HasField checks if there are errors for a specific field.
// HasField는 특정 필드에 대한 에러가 있는지 확인합니다.
func (ve ValidationErrors) HasField(field string) bool {
	for _, err := range ve {
		if err.Field == field {
			return true
		}
	}
	return false
}

// GetField returns all errors for a specific field.
// GetField는 특정 필드에 대한 모든 에러를 반환합니다.
func (ve ValidationErrors) GetField(field string) []ValidationError {
	var errors []ValidationError
	for _, err := range ve {
		if err.Field == field {
			errors = append(errors, err)
		}
	}
	return errors
}

// ToMap converts validation errors to a map of field names to error messages.
// ToMap은 검증 에러를 필드 이름과 에러 메시지의 맵으로 변환합니다.
func (ve ValidationErrors) ToMap() map[string][]string {
	result := make(map[string][]string)
	for _, err := range ve {
		result[err.Field] = append(result[err.Field], err.Error())
	}
	return result
}

// First returns the first validation error, or nil if there are no errors.
// First는 첫 번째 검증 에러를 반환하거나, 에러가 없으면 nil을 반환합니다.
func (ve ValidationErrors) First() *ValidationError {
	if len(ve) == 0 {
		return nil
	}
	return &ve[0]
}

// Count returns the number of validation errors.
// Count는 검증 에러의 개수를 반환합니다.
func (ve ValidationErrors) Count() int {
	return len(ve)
}
