package validation

import (
	"fmt"
)

// ============================================================================
// LOGICAL/CONDITIONAL VALIDATORS
// ============================================================================
//
// This file provides logical and conditional validation functions for:
// - OneOf (value must match one of the provided values)
// - NotOneOf (value must not match any of the provided values)
// - When (conditional validation based on predicate)
// - Unless (conditional validation unless predicate is true)
//
// 이 파일은 다음을 위한 논리 및 조건부 검증 함수를 제공합니다:
// - OneOf (값이 제공된 값 중 하나와 일치해야 함)
// - NotOneOf (값이 제공된 값 중 어느 것과도 일치하지 않아야 함)
// - When (조건부 검증 - 조건이 참일 때)
// - Unless (조건부 검증 - 조건이 거짓일 때)
//
// ============================================================================

// OneOf validates that the value matches one of the provided values.
// OneOf는 값이 제공된 값 중 하나와 일치하는지 검증합니다.
//
// This validator checks if the value equals any of the allowed values.
// Useful for validating enum-like fields or restricted choice fields.
// 이 검증기는 값이 허용된 값 중 하나와 같은지 확인합니다.
// 열거형 필드 또는 제한된 선택 필드를 검증하는 데 유용합니다.
//
// Example / 예시:
//
//	status := "active"
//	v := validation.New(status, "status")
//	v.OneOf("active", "inactive", "pending")
//
// Validation rules / 검증 규칙:
//   - Value must equal one of the provided values / 값이 제공된 값 중 하나와 같아야 함
//   - Comparison uses == operator / 비교는 == 연산자 사용
//   - At least one allowed value must be provided / 최소 하나의 허용 값이 제공되어야 함
func (v *Validator) OneOf(values ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if len(values) == 0 {
		v.addError("one_of", fmt.Sprintf("%s OneOf requires at least one value / %s OneOf는 최소 하나의 값이 필요합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if value matches any of the provided values
	for _, allowed := range values {
		if v.value == allowed {
			return v // Match found
		}
	}

	// No match found - create error message
	v.addError("one_of", fmt.Sprintf("%s must be one of the allowed values / %s은(는) 허용된 값 중 하나여야 합니다", v.fieldName, v.fieldName))
	return v
}

// NotOneOf validates that the value does not match any of the provided values.
// NotOneOf는 값이 제공된 값 중 어느 것과도 일치하지 않는지 검증합니다.
//
// This validator checks if the value is different from all forbidden values.
// Useful for blacklisting specific values or preventing reserved keywords.
// 이 검증기는 값이 모든 금지된 값과 다른지 확인합니다.
// 특정 값을 블랙리스트에 추가하거나 예약어를 방지하는 데 유용합니다.
//
// Example / 예시:
//
//	username := "admin"
//	v := validation.New(username, "username")
//	v.NotOneOf("admin", "root", "administrator")
//
// Validation rules / 검증 규칙:
//   - Value must not equal any of the provided values / 값이 제공된 값 중 어느 것과도 같지 않아야 함
//   - Comparison uses == operator / 비교는 == 연산자 사용
//   - At least one forbidden value must be provided / 최소 하나의 금지 값이 제공되어야 함
func (v *Validator) NotOneOf(values ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if len(values) == 0 {
		v.addError("not_one_of", fmt.Sprintf("%s NotOneOf requires at least one value / %s NotOneOf는 최소 하나의 값이 필요합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if value matches any of the forbidden values
	for _, forbidden := range values {
		if v.value == forbidden {
			v.addError("not_one_of", fmt.Sprintf("%s must not be one of the forbidden values / %s은(는) 금지된 값 중 하나가 아니어야 합니다", v.fieldName, v.fieldName))
			return v
		}
	}

	return v // Value is not in forbidden list
}

// When executes the validation function only if the predicate returns true.
// When은 조건이 참을 반환할 때만 검증 함수를 실행합니다.
//
// This validator allows conditional validation based on runtime conditions.
// The validation function is only executed if the predicate evaluates to true.
// 이 검증기는 런타임 조건에 따라 조건부 검증을 허용합니다.
// 검증 함수는 조건이 참으로 평가될 때만 실행됩니다.
//
// Example / 예시:
//
//	age := 25
//	isAdult := age >= 18
//	v := validation.New(age, "age")
//	v.When(isAdult, func(val *Validator) {
//	    val.Min(18)
//	})
//
// Validation rules / 검증 규칙:
//   - If predicate is true, validation function is executed / 조건이 참이면 검증 함수 실행
//   - If predicate is false, validation is skipped / 조건이 거짓이면 검증 건너뜀
//   - Useful for complex conditional validation logic / 복잡한 조건부 검증 로직에 유용
func (v *Validator) When(predicate bool, fn func(*Validator)) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if predicate {
		fn(v)
	}

	return v
}

// Unless executes the validation function only if the predicate returns false.
// Unless는 조건이 거짓을 반환할 때만 검증 함수를 실행합니다.
//
// This validator is the inverse of When - it executes validation when condition is false.
// Useful for "validate unless X" scenarios.
// 이 검증기는 When의 반대입니다 - 조건이 거짓일 때 검증을 실행합니다.
// "X가 아닌 경우 검증" 시나리오에 유용합니다.
//
// Example / 예시:
//
//	email := "user@example.com"
//	isGuest := false
//	v := validation.New(email, "email")
//	v.Unless(isGuest, func(val *Validator) {
//	    val.Required().Email()
//	})
//
// Validation rules / 검증 규칙:
//   - If predicate is false, validation function is executed / 조건이 거짓이면 검증 함수 실행
//   - If predicate is true, validation is skipped / 조건이 참이면 검증 건너뜀
//   - Inverse of When validator / When 검증기의 반대
func (v *Validator) Unless(predicate bool, fn func(*Validator)) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if !predicate {
		fn(v)
	}

	return v
}
