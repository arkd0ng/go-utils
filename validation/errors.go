package validation

import (
	"fmt"
	"strings"
)

// ValidationError represents a single validation error.
// ValidationError는 단일 검증 에러를 나타냅니다.
type ValidationError struct {
	Field   string      // Field name / 필드 이름
	Value   interface{} // Invalid value / 유효하지 않은 값
	Rule    string      // Failed rule / 실패한 규칙
	Message string      // Error message / 에러 메시지
}

// Error returns the error message.
// Error는 에러 메시지를 반환합니다.
func (ve ValidationError) Error() string {
	if ve.Message != "" {
		return ve.Message
	}
	return fmt.Sprintf("%s validation failed for field '%s'", ve.Rule, ve.Field)
}

// ValidationErrors is a collection of validation errors.
// ValidationErrors는 검증 에러의 모음입니다.
type ValidationErrors []ValidationError

// Error returns a formatted error message for all validation errors.
// Error는 모든 검증 에러에 대한 포맷된 에러 메시지를 반환합니다.
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
