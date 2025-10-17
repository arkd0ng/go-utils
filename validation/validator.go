package validation

import "fmt"

// New creates a new Validator for the given value and field name.
// New는 주어진 값과 필드 이름에 대한 새 Validator를 생성합니다.
func New(value interface{}, fieldName string) *Validator {
	return &Validator{
		value:       value,
		fieldName:   fieldName,
		errors:      []ValidationError{},
		stopOnError: false,
	}
}

// Validate executes all validation rules and returns an error if any fail.
// Validate는 모든 검증 규칙을 실행하고, 실패 시 에러를 반환합니다.
func (v *Validator) Validate() error {
	if len(v.errors) == 0 {
		return nil
	}
	return ValidationErrors(v.errors)
}

// GetErrors returns all validation errors.
// GetErrors는 모든 검증 에러를 반환합니다.
func (v *Validator) GetErrors() []ValidationError {
	return v.errors
}

// StopOnError sets the validator to stop on the first error.
// StopOnError는 첫 번째 에러에서 멈추도록 설정합니다.
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

	v.errors = append(v.errors, ValidationError{
		Field:   v.fieldName,
		Value:   v.value,
		Rule:    rule,
		Message: message,
	})
	v.lastRule = rule
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
