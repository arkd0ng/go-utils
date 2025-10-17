package validation

import (
	"fmt"
	"reflect"
)

// ============================================================================
// TYPE-SPECIFIC VALIDATORS
// ============================================================================
//
// This file provides type-specific validation functions for:
// - True/False (boolean validation)
// - Nil/NotNil (nil check)
// - Type (type validation)
// - Empty/NotEmpty (empty value check)
//
// 이 파일은 다음을 위한 타입별 검증 함수를 제공합니다:
// - True/False (불리언 검증)
// - Nil/NotNil (nil 검사)
// - Type (타입 검증)
// - Empty/NotEmpty (빈 값 검사)
//
// ============================================================================

// True validates that the value is boolean true.
// True는 값이 불리언 true인지 검증합니다.
//
// This validator checks if the value is exactly boolean true.
// Useful for checkbox validation, terms acceptance, etc.
// 이 검증기는 값이 정확히 불리언 true인지 확인합니다.
// 체크박스 검증, 약관 동의 등에 유용합니다.
//
// Example / 예시:
//
//	accepted := true
//	v := validation.New(accepted, "terms")
//	v.True()
//
// Validation rules / 검증 규칙:
//   - Value must be boolean type / 값이 불리언 타입이어야 함
//   - Value must be true / 값이 true여야 함
func (v *Validator) True() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	boolVal, ok := v.value.(bool)
	if !ok {
		v.addError("true", fmt.Sprintf("%s must be boolean / %s은(는) 불리언이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if !boolVal {
		v.addError("true", fmt.Sprintf("%s must be true / %s은(는) true여야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// False validates that the value is boolean false.
// False는 값이 불리언 false인지 검증합니다.
//
// This validator checks if the value is exactly boolean false.
// Useful for negative confirmation validation.
// 이 검증기는 값이 정확히 불리언 false인지 확인합니다.
// 부정 확인 검증에 유용합니다.
//
// Example / 예시:
//
//	declined := false
//	v := validation.New(declined, "marketing")
//	v.False()
//
// Validation rules / 검증 규칙:
//   - Value must be boolean type / 값이 불리언 타입이어야 함
//   - Value must be false / 값이 false여야 함
func (v *Validator) False() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	boolVal, ok := v.value.(bool)
	if !ok {
		v.addError("false", fmt.Sprintf("%s must be boolean / %s은(는) 불리언이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if boolVal {
		v.addError("false", fmt.Sprintf("%s must be false / %s은(는) false여야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// Nil validates that the value is nil.
// Nil은 값이 nil인지 검증합니다.
//
// This validator checks if the value is nil (pointer, interface, slice, map, channel, func).
// Useful for validating optional fields that should be nil.
// 이 검증기는 값이 nil인지 확인합니다 (포인터, 인터페이스, 슬라이스, 맵, 채널, 함수).
// nil이어야 하는 선택적 필드를 검증하는 데 유용합니다.
//
// Example / 예시:
//
//	var ptr *string
//	v := validation.New(ptr, "optional")
//	v.Nil()
//
// Validation rules / 검증 규칙:
//   - Value must be nil / 값이 nil이어야 함
//   - Works with pointers, interfaces, slices, maps, channels, functions
func (v *Validator) Nil() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if v.value != nil {
		// Check if it's a nil pointer/interface using reflection
		rv := reflect.ValueOf(v.value)
		if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface ||
			rv.Kind() == reflect.Slice || rv.Kind() == reflect.Map ||
			rv.Kind() == reflect.Chan || rv.Kind() == reflect.Func {
			if !rv.IsNil() {
				v.addError("nil", fmt.Sprintf("%s must be nil / %s은(는) nil이어야 합니다", v.fieldName, v.fieldName))
			}
		} else {
			v.addError("nil", fmt.Sprintf("%s must be nil / %s은(는) nil이어야 합니다", v.fieldName, v.fieldName))
		}
	}

	return v
}

// NotNil validates that the value is not nil.
// NotNil은 값이 nil이 아닌지 검증합니다.
//
// This validator checks if the value is not nil (pointer, interface, slice, map, channel, func).
// Useful for validating required pointer fields.
// 이 검증기는 값이 nil이 아닌지 확인합니다 (포인터, 인터페이스, 슬라이스, 맵, 채널, 함수).
// 필수 포인터 필드를 검증하는 데 유용합니다.
//
// Example / 예시:
//
//	str := "value"
//	ptr := &str
//	v := validation.New(ptr, "required_ptr")
//	v.NotNil()
//
// Validation rules / 검증 규칙:
//   - Value must not be nil / 값이 nil이 아니어야 함
//   - Works with pointers, interfaces, slices, maps, channels, functions
func (v *Validator) NotNil() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if v.value == nil {
		v.addError("not_nil", fmt.Sprintf("%s must not be nil / %s은(는) nil이 아니어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if it's a nil pointer/interface using reflection
	rv := reflect.ValueOf(v.value)
	if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface ||
		rv.Kind() == reflect.Slice || rv.Kind() == reflect.Map ||
		rv.Kind() == reflect.Chan || rv.Kind() == reflect.Func {
		if rv.IsNil() {
			v.addError("not_nil", fmt.Sprintf("%s must not be nil / %s은(는) nil이 아니어야 합니다", v.fieldName, v.fieldName))
		}
	}

	return v
}

// Type validates that the value is of the specified type.
// Type은 값이 지정된 타입인지 검증합니다.
//
// This validator checks if the value matches the specified type using reflection.
// Useful for dynamic type validation.
// 이 검증기는 리플렉션을 사용하여 값이 지정된 타입과 일치하는지 확인합니다.
// 동적 타입 검증에 유용합니다.
//
// Example / 예시:
//
//	value := "hello"
//	v := validation.New(value, "text")
//	v.Type("string")
//
// Validation rules / 검증 규칙:
//   - Value type must match specified type name / 값 타입이 지정된 타입 이름과 일치해야 함
//   - Type names: "string", "int", "float64", "bool", "slice", "map", "struct", etc.
func (v *Validator) Type(typeName string) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	actualType := reflect.TypeOf(v.value)
	if actualType == nil {
		v.addError("type", fmt.Sprintf("%s must be of type %s but is nil / %s은(는) %s 타입이어야 하지만 nil입니다", v.fieldName, typeName, v.fieldName, typeName))
		return v
	}

	actualTypeName := actualType.Kind().String()

	// Handle specific type names
	if typeName != actualTypeName {
		v.addError("type", fmt.Sprintf("%s must be of type %s but is %s / %s은(는) %s 타입이어야 하지만 %s입니다", v.fieldName, typeName, actualTypeName, v.fieldName, typeName, actualTypeName))
	}

	return v
}

// Empty validates that the value is empty (zero value).
// Empty는 값이 비어있는지 (제로 값) 검증합니다.
//
// This validator checks if the value is the zero value for its type:
// - string: ""
// - numbers: 0
// - bool: false
// - slice/map: nil or len == 0
// - pointer: nil
// 이 검증기는 값이 해당 타입의 제로 값인지 확인합니다:
// - 문자열: ""
// - 숫자: 0
// - 불리언: false
// - 슬라이스/맵: nil 또는 len == 0
// - 포인터: nil
//
// Example / 예시:
//
//	value := ""
//	v := validation.New(value, "optional")
//	v.Empty()
//
// Validation rules / 검증 규칙:
//   - Value must be zero value for its type / 값이 해당 타입의 제로 값이어야 함
func (v *Validator) Empty() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if !isEmptyValue(v.value) {
		v.addError("empty", fmt.Sprintf("%s must be empty / %s은(는) 비어있어야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// NotEmpty validates that the value is not empty (not zero value).
// NotEmpty는 값이 비어있지 않은지 (제로 값이 아닌지) 검증합니다.
//
// This validator checks if the value is not the zero value for its type.
// Similar to Required() but works for any type.
// 이 검증기는 값이 해당 타입의 제로 값이 아닌지 확인합니다.
// Required()와 유사하지만 모든 타입에 작동합니다.
//
// Example / 예시:
//
//	value := "hello"
//	v := validation.New(value, "required")
//	v.NotEmpty()
//
// Validation rules / 검증 규칙:
//   - Value must not be zero value for its type / 값이 해당 타입의 제로 값이 아니어야 함
func (v *Validator) NotEmpty() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if isEmptyValue(v.value) {
		v.addError("not_empty", fmt.Sprintf("%s must not be empty / %s은(는) 비어있지 않아야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// isEmptyValue checks if a value is empty/zero value
func isEmptyValue(val interface{}) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Chan:
		return v.IsNil() || v.Len() == 0
	case reflect.Array:
		return v.Len() == 0
	}
	return false
}
