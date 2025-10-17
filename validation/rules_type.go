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
// Requires exact boolean true value for strict validation.
//
// True는 값이 불리언 true인지 검증합니다.
// 엄격한 검증을 위해 정확한 불리언 true 값이 필요합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Requires boolean type
//     불리언 타입 필요
//   - Must be exactly true
//     정확히 true여야 함
//   - Fails on false or non-boolean
//     false 또는 불리언이 아니면 실패
//   - No truthy conversion
//     truthy 변환 없음
//
// Use Cases / 사용 사례:
//   - Checkbox acceptance / 체크박스 동의
//   - Terms and conditions / 약관 동의
//   - Feature flags / 기능 플래그
//   - Permission grants / 권한 부여
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Type assertion only
//     타입 단언만
//
// This validator checks if the value is exactly boolean true.
// Useful for checkbox validation, terms acceptance, etc.
// 이 검증기는 값이 정확히 불리언 true인지 확인합니다.
// 체크박스 검증, 약관 동의 등에 유용합니다.
//
// Example / 예시:
//
//	// Valid - true / 유효 - true
//	accepted := true
//	v := validation.New(accepted, "terms")
//	v.True()  // Passes
//
//	// Invalid - false / 무효 - false
//	v = validation.New(false, "terms")
//	v.True()  // Fails
//
//	// Invalid - not boolean / 무효 - 불리언 아님
//	v = validation.New(1, "value")
//	v.True()  // Fails (not bool type)
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
// Requires exact boolean false value for strict validation.
//
// False는 값이 불리언 false인지 검증합니다.
// 엄격한 검증을 위해 정확한 불리언 false 값이 필요합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Requires boolean type
//     불리언 타입 필요
//   - Must be exactly false
//     정확히 false여야 함
//   - Fails on true or non-boolean
//     true 또는 불리언이 아니면 실패
//   - No falsy conversion
//     falsy 변환 없음
//
// Use Cases / 사용 사례:
//   - Negative confirmation / 부정 확인
//   - Opt-out validation / 선택 해제 검증
//   - Disabled state check / 비활성화 상태 확인
//   - Default false requirement / 기본 false 요구사항
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Type assertion only
//     타입 단언만
//
// This validator checks if the value is exactly boolean false.
// Useful for negative confirmation validation.
// 이 검증기는 값이 정확히 불리언 false인지 확인합니다.
// 부정 확인 검증에 유용합니다.
//
// Example / 예시:
//
//	// Valid - false / 유효 - false
//	declined := false
//	v := validation.New(declined, "marketing")
//	v.False()  // Passes
//
//	// Invalid - true / 무효 - true
//	v = validation.New(true, "opt_out")
//	v.False()  // Fails
//
//	// Invalid - not boolean / 무효 - 불리언 아님
//	v = validation.New(0, "value")
//	v.False()  // Fails (not bool type)
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
// Uses reflection to check nil pointers, interfaces, slices, maps, channels, and functions.
//
// Nil은 값이 nil인지 검증합니다.
// 리플렉션을 사용하여 nil 포인터, 인터페이스, 슬라이스, 맵, 채널 및 함수를 확인합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts nil value
//     nil 값 허용
//   - Checks nil pointers via reflection
//     리플렉션을 통한 nil 포인터 확인
//   - Supports: pointer, interface, slice, map, channel, func
//     지원: 포인터, 인터페이스, 슬라이스, 맵, 채널, 함수
//   - Fails on non-nil values
//     nil이 아닌 값에 실패
//
// Use Cases / 사용 사례:
//   - Optional field validation / 선택적 필드 검증
//   - Null state verification / null 상태 확인
//   - Pointer absence check / 포인터 부재 확인
//   - Resource cleanup validation / 리소스 정리 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Reflection-based nil check
//     리플렉션 기반 nil 확인
//
// This validator checks if the value is nil (pointer, interface, slice, map, channel, func).
// Useful for validating optional fields that should be nil.
// 이 검증기는 값이 nil인지 확인합니다 (포인터, 인터페이스, 슬라이스, 맵, 채널, 함수).
// nil이어야 하는 선택적 필드를 검증하는 데 유용합니다.
//
// Example / 예시:
//
//	// Valid - nil pointer / 유효 - nil 포인터
//	var ptr *string
//	v := validation.New(ptr, "optional")
//	v.Nil()  // Passes
//
//	// Valid - nil slice / 유효 - nil 슬라이스
//	var slice []int
//	v = validation.New(slice, "data")
//	v.Nil()  // Passes
//
//	// Invalid - non-nil / 무효 - nil 아님
//	str := "value"
//	v = validation.New(&str, "ptr")
//	v.Nil()  // Fails
//
//	// Invalid - value type / 무효 - 값 타입
//	v = validation.New(42, "number")
//	v.Nil()  // Fails (not nillable type)
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
// Uses reflection to ensure pointers, interfaces, slices, maps, channels, and functions are not nil.
//
// NotNil은 값이 nil이 아닌지 검증합니다.
// 리플렉션을 사용하여 포인터, 인터페이스, 슬라이스, 맵, 채널 및 함수가 nil이 아님을 보장합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Rejects nil value
//     nil 값 거부
//   - Checks non-nil pointers via reflection
//     리플렉션을 통한 nil이 아닌 포인터 확인
//   - Supports: pointer, interface, slice, map, channel, func
//     지원: 포인터, 인터페이스, 슬라이스, 맵, 채널, 함수
//   - Fails on nil values
//     nil 값에 실패
//
// Use Cases / 사용 사례:
//   - Required pointer fields / 필수 포인터 필드
//   - Non-null validation / non-null 검증
//   - Reference existence check / 참조 존재 확인
//   - Dependency validation / 의존성 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Reflection-based nil check
//     리플렉션 기반 nil 확인
//
// This validator checks if the value is not nil (pointer, interface, slice, map, channel, func).
// Useful for validating required pointer fields.
// 이 검증기는 값이 nil이 아닌지 확인합니다 (포인터, 인터페이스, 슬라이스, 맵, 채널, 함수).
// 필수 포인터 필드를 검증하는 데 유용합니다.
//
// Example / 예시:
//
//	// Valid - non-nil pointer / 유효 - nil이 아닌 포인터
//	str := "value"
//	ptr := &str
//	v := validation.New(ptr, "required_ptr")
//	v.NotNil()  // Passes
//
//	// Valid - initialized slice / 유효 - 초기화된 슬라이스
//	slice := []int{1, 2, 3}
//	v = validation.New(slice, "data")
//	v.NotNil()  // Passes
//
//	// Invalid - nil pointer / 무효 - nil 포인터
//	var nilPtr *string
//	v = validation.New(nilPtr, "ptr")
//	v.NotNil()  // Fails
//
//	// Invalid - nil slice / 무효 - nil 슬라이스
//	var nilSlice []int
//	v = validation.New(nilSlice, "data")
//	v.NotNil()  // Fails
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
// Uses reflection to compare runtime type with expected type name.
//
// Type은 값이 지정된 타입인지 검증합니다.
// 리플렉션을 사용하여 런타임 타입과 예상 타입 이름을 비교합니다.
//
// Parameters / 매개변수:
//   - typeName: Expected type name (e.g., "string", "int", "slice")
//     예상 타입 이름 (예: "string", "int", "slice")
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflect.Kind() for type checking
//     타입 확인에 reflect.Kind() 사용
//   - Compares type name strings
//     타입 이름 문자열 비교
//   - Fails on nil value
//     nil 값에 실패
//   - Case-sensitive type matching
//     대소문자 구분 타입 매칭
//
// Use Cases / 사용 사례:
//   - Dynamic type validation / 동적 타입 검증
//   - Generic function validation / 제네릭 함수 검증
//   - Interface type checking / 인터페이스 타입 확인
//   - Runtime type enforcement / 런타임 타입 강제
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Reflection type comparison
//     리플렉션 타입 비교
//
// This validator checks if the value matches the specified type using reflection.
// Useful for dynamic type validation.
// 이 검증기는 리플렉션을 사용하여 값이 지정된 타입과 일치하는지 확인합니다.
// 동적 타입 검증에 유용합니다.
//
// Example / 예시:
//
//	// Valid string type / 유효한 문자열 타입
//	value := "hello"
//	v := validation.New(value, "text")
//	v.Type("string")  // Passes
//
//	// Valid int type / 유효한 정수 타입
//	v = validation.New(42, "number")
//	v.Type("int")  // Passes
//
//	// Invalid type / 무효한 타입
//	v = validation.New("123", "number")
//	v.Type("int")  // Fails (is string)
//
//	// Nil value / Nil 값
//	v = validation.New(nil, "value")
//	v.Type("string")  // Fails (is nil)
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
// Checks for zero values based on type: "", 0, false, nil, empty slices/maps.
//
// Empty는 값이 비어있는지 (제로 값) 검증합니다.
// 타입에 따른 제로 값 확인: "", 0, false, nil, 빈 슬라이스/맵.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - String: empty string ""
//     문자열: 빈 문자열 ""
//   - Numbers: 0
//     숫자: 0
//   - Boolean: false
//     불리언: false
//   - Slice/Map: nil or length 0
//     슬라이스/맵: nil 또는 길이 0
//   - Pointer: nil
//     포인터: nil
//
// Use Cases / 사용 사례:
//   - Optional field check / 선택적 필드 확인
//   - Default value validation / 기본값 검증
//   - Uninitialized state check / 초기화되지 않은 상태 확인
//   - Zero value detection / 제로 값 감지
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1) for most types
//     시간 복잡도: 대부분 타입에서 O(1)
//   - Reflection-based type switch
//     리플렉션 기반 타입 스위치
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
//	// Valid empty string / 유효한 빈 문자열
//	value := ""
//	v := validation.New(value, "optional")
//	v.Empty()  // Passes
//
//	// Valid zero number / 유효한 제로 숫자
//	v = validation.New(0, "count")
//	v.Empty()  // Passes
//
//	// Valid false boolean / 유효한 false 불리언
//	v = validation.New(false, "flag")
//	v.Empty()  // Passes
//
//	// Invalid - non-empty / 무효 - 비어있지 않음
//	v = validation.New("hello", "text")
//	v.Empty()  // Fails
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
// Ensures value is not the zero value for its type, works across all types.
//
// NotEmpty는 값이 비어있지 않은지 (제로 값이 아닌지) 검증합니다.
// 값이 해당 타입의 제로 값이 아님을 보장하며 모든 타입에서 작동합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - String: non-empty string
//     문자열: 비어있지 않은 문자열
//   - Numbers: non-zero
//     숫자: 0이 아님
//   - Boolean: true
//     불리언: true
//   - Slice/Map: non-nil and length > 0
//     슬라이스/맵: nil이 아니고 길이 > 0
//   - Pointer: non-nil
//     포인터: nil이 아님
//
// Use Cases / 사용 사례:
//   - Required field validation / 필수 필드 검증
//   - Non-zero value check / 0이 아닌 값 확인
//   - Initialized state verification / 초기화된 상태 확인
//   - Generic required validation / 제네릭 필수 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1) for most types
//     시간 복잡도: 대부분 타입에서 O(1)
//   - Reflection-based type switch
//     리플렉션 기반 타입 스위치
//
// This validator checks if the value is not the zero value for its type.
// Similar to Required() but works for any type.
// 이 검증기는 값이 해당 타입의 제로 값이 아닌지 확인합니다.
// Required()와 유사하지만 모든 타입에 작동합니다.
//
// Example / 예시:
//
//	// Valid non-empty string / 유효한 비어있지 않은 문자열
//	value := "hello"
//	v := validation.New(value, "required")
//	v.NotEmpty()  // Passes
//
//	// Valid non-zero number / 유효한 0이 아닌 숫자
//	v = validation.New(42, "count")
//	v.NotEmpty()  // Passes
//
//	// Valid true boolean / 유효한 true 불리언
//	v = validation.New(true, "flag")
//	v.NotEmpty()  // Passes
//
//	// Invalid - empty string / 무효 - 빈 문자열
//	v = validation.New("", "text")
//	v.NotEmpty()  // Fails
//
//	// Invalid - zero / 무효 - 제로
//	v = validation.New(0, "number")
//	v.NotEmpty()  // Fails
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
