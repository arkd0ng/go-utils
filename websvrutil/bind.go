package websvrutil

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// bind.go provides reflection-based data binding utilities for converting
// URL form values into Go structs.
//
// This file implements the core binding logic used by Context.BindForm()
// and Context.BindQuery() methods to automatically populate struct fields
// from HTTP form data:
//
// Core Functions:
//   - bindFormData(): Main binding function that maps url.Values to struct
//     Uses reflection to iterate fields and set values
//     Respects "form" struct tags for field name mapping
//   - setFieldValue(): Type-specific value conversion and assignment
//     Converts string form values to appropriate Go types
//
// Supported Types:
//   - string: Direct string assignment
//   - Integers: int, int8, int16, int32, int64
//     Conversion: strconv.ParseInt(value, 10, 64)
//   - Unsigned integers: uint, uint8, uint16, uint32, uint64
//     Conversion: strconv.ParseUint(value, 10, 64)
//   - Floating point: float32, float64
//     Conversion: strconv.ParseFloat(value, 64)
//   - Boolean: bool
//     Conversion: strconv.ParseBool(value)
//     Accepts: "1", "t", "T", "true", "TRUE", "True", "0", "f", "F", "false", "FALSE", "False"
//
// Struct Tag Format:
//   - "form" tag specifies form field name for mapping
//   - If tag is absent, uses struct field name directly
//   - Example:
//     type User struct {
//         Username string `form:"username"`  // Maps to "username" form field
//         Email    string `form:"email"`     // Maps to "email" form field
//         Age      int    `form:"age"`       // Maps to "age" form field
//         IsAdmin  bool   // Maps to "IsAdmin" (no tag, uses field name)
//     }
//
// Binding Process:
//   1. Validate that obj is a pointer to struct (returns error if not)
//   2. Iterate through all struct fields using reflection
//   3. For each field:
//      - Extract "form" tag (or use field name if no tag)
//      - Retrieve corresponding value from url.Values
//      - Convert string value to field's type using strconv
//      - Set field value using reflection
//   4. Skip fields that:
//      - Are unexported (cannot be set via reflection)
//      - Have no corresponding form value (empty string)
//
// Error Handling:
//   - Returns error if obj is not a pointer to struct
//   - Returns error if type conversion fails (e.g., "abc" → int)
//   - Wraps conversion errors with field name and type information
//
// Example Usage:
//   // HTTP request: POST /signup?username=john&email=john@example.com&age=30
//   type SignupForm struct {
//       Username string `form:"username"`
//       Email    string `form:"email"`
//       Age      int    `form:"age"`
//   }
//   var form SignupForm
//   err := ctx.BindQuery(&form)
//   // form.Username = "john", form.Email = "john@example.com", form.Age = 30
//
// Performance:
//   - Time complexity: O(n) where n = number of struct fields
//   - Uses reflection, so has overhead compared to manual parsing
//   - Efficient for forms with many fields (eliminates repetitive parsing code)
//   - Memory: Minimal allocations (reuses existing struct instance)
//
// Security Considerations:
//   - Always validate bound data before using (binding != validation)
//   - Use validator package for comprehensive validation rules
//   - Be cautious with unexported fields (they are skipped during binding)
//
// Limitations:
//   - Does not support nested structs or complex types
//   - Does not support slice/array fields
//   - Requires exact type match (no automatic type coercion beyond strconv)
//
// bind.go는 URL 폼 값을 Go 구조체로 변환하기 위한 리플렉션 기반
// 데이터 바인딩 유틸리티를 제공합니다.
//
// 이 파일은 Context.BindForm() 및 Context.BindQuery() 메서드가
// HTTP 폼 데이터에서 구조체 필드를 자동으로 채우는 데 사용하는
// 핵심 바인딩 로직을 구현합니다:
//
// 핵심 함수:
//   - bindFormData(): url.Values를 구조체에 매핑하는 주요 바인딩 함수
//     리플렉션을 사용하여 필드를 반복하고 값 설정
//     필드 이름 매핑을 위해 "form" 구조체 태그 존중
//   - setFieldValue(): 타입별 값 변환 및 할당
//     문자열 폼 값을 적절한 Go 타입으로 변환
//
// 지원되는 타입:
//   - string: 직접 문자열 할당
//   - 정수: int, int8, int16, int32, int64
//     변환: strconv.ParseInt(value, 10, 64)
//   - 부호 없는 정수: uint, uint8, uint16, uint32, uint64
//     변환: strconv.ParseUint(value, 10, 64)
//   - 부동 소수점: float32, float64
//     변환: strconv.ParseFloat(value, 64)
//   - 부울: bool
//     변환: strconv.ParseBool(value)
//     허용: "1", "t", "T", "true", "TRUE", "True", "0", "f", "F", "false", "FALSE", "False"
//
// 구조체 태그 형식:
//   - "form" 태그는 매핑을 위한 폼 필드 이름 지정
//   - 태그가 없으면 구조체 필드 이름을 직접 사용
//   - 예제:
//     type User struct {
//         Username string `form:"username"`  // "username" 폼 필드에 매핑
//         Email    string `form:"email"`     // "email" 폼 필드에 매핑
//         Age      int    `form:"age"`       // "age" 폼 필드에 매핑
//         IsAdmin  bool   // "IsAdmin"에 매핑 (태그 없음, 필드 이름 사용)
//     }
//
// 바인딩 프로세스:
//   1. obj가 구조체 포인터인지 검증 (아니면 에러 반환)
//   2. 리플렉션을 사용하여 모든 구조체 필드 반복
//   3. 각 필드에 대해:
//      - "form" 태그 추출 (태그 없으면 필드 이름 사용)
//      - url.Values에서 해당 값 검색
//      - strconv를 사용하여 문자열 값을 필드 타입으로 변환
//      - 리플렉션을 사용하여 필드 값 설정
//   4. 다음 필드는 건너뜀:
//      - 내보내지지 않은 필드 (리플렉션으로 설정 불가)
//      - 해당 폼 값이 없는 필드 (빈 문자열)
//
// 에러 처리:
//   - obj가 구조체 포인터가 아니면 에러 반환
//   - 타입 변환 실패 시 에러 반환 (예: "abc" → int)
//   - 필드 이름 및 타입 정보와 함께 변환 에러 래핑
//
// 사용 예제:
//   // HTTP 요청: POST /signup?username=john&email=john@example.com&age=30
//   type SignupForm struct {
//       Username string `form:"username"`
//       Email    string `form:"email"`
//       Age      int    `form:"age"`
//   }
//   var form SignupForm
//   err := ctx.BindQuery(&form)
//   // form.Username = "john", form.Email = "john@example.com", form.Age = 30
//
// 성능:
//   - 시간 복잡도: O(n), n = 구조체 필드 수
//   - 리플렉션 사용으로 수동 파싱 대비 오버헤드 존재
//   - 필드가 많은 폼에 효율적 (반복적인 파싱 코드 제거)
//   - 메모리: 최소 할당 (기존 구조체 인스턴스 재사용)
//
// 보안 고려사항:
//   - 사용 전 항상 바인딩된 데이터 검증 (바인딩 ≠ 검증)
//   - 포괄적인 검증 규칙을 위해 validator 패키지 사용
//   - 내보내지지 않은 필드 주의 (바인딩 중 건너뜀)
//
// 제한사항:
//   - 중첩 구조체나 복잡한 타입 미지원
//   - 슬라이스/배열 필드 미지원
//   - 정확한 타입 일치 필요 (strconv 이상의 자동 타입 강제 변환 없음)

// bindFormData binds URL values to a struct using reflection.
// bindFormData는 리플렉션을 사용하여 URL 값을 구조체에 바인딩합니다.
//
// Process
// 프로세스:
//  1. Validate obj is a pointer to struct (returns error if not)
//  2. Iterate through struct fields
//  3. Extract "form" tag for field name mapping (uses field name if tag absent)
//  4. Retrieve corresponding value from URL values
//  5. Convert string value to appropriate type using strconv
//  6. Set field value using reflection
//
// Supported types
// 지원되는 타입:
//   - string: Direct assignment
//   - int, int8, int16, int32, int64: ParseInt with base 10
//   - uint, uint8, uint16, uint32, uint64: ParseUint with base 10
//   - float32, float64: ParseFloat
//   - bool: ParseBool (accepts "1", "t", "T", "true", "TRUE", "True", "0", "f", "F", "false", "FALSE", "False")
//
// Struct tag format
// 구조체 태그 형식:
//
//	type User struct {
//	    Name  string `form:"name"`      // Maps to "name" form field
//	    Email string `form:"email"`     // Maps to "email" form field
//	    Age   int    `form:"age"`       // Maps to "age" form field
//	    Admin bool   // Maps to "Admin" (no tag, uses field name)
//	}
//
// Error handling
// 에러 처리:
//   - Returns error if obj is not a pointer to struct
//   - Returns error if type conversion fails (e.g., "abc" -> int)
//   - Skips fields that:
//   - Are unexported (cannot be set)
//   - Have no corresponding form value (empty string)
//
// Example usage
// 사용 예제:
//
//	values := url.Values{
//	    "name":  []string{"John"},
//	    "age":   []string{"30"},
//	    "admin": []string{"true"},
//	}
//	var user User
//	err := bindFormData(&user, values)
//	// user.Name = "John", user.Age = 30, user.Admin = true
//
// Performance
// 성능:
//   - Time complexity: O(n) where n = number of struct fields
//   - Uses reflection, so has overhead compared to manual parsing
//   - Efficient for forms with many fields (reduces boilerplate code)
func bindFormData(obj interface{}, values url.Values) error {
	// Get the reflect value and type of the object
	// 객체의 리플렉트 값과 타입 가져오기
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("binding requires a pointer to a struct, got %s", val.Kind())
	}

	// Dereference the pointer
	// 포인터 역참조
	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("binding requires a pointer to a struct, got pointer to %s", val.Kind())
	}

	typ := val.Type()

	// Iterate over struct fields
	// 구조체 필드 반복
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := typ.Field(i)

		// Skip unexported fields
		// 내보내지 않은 필드 건너뛰기
		if !field.CanSet() {
			continue
		}

		// Get the form tag
		// form 태그 가져오기
		tag := typeField.Tag.Get("form")
		if tag == "" {
			// Use field name if no tag specified
			// 태그가 지정되지 않은 경우 필드 이름 사용
			tag = typeField.Name
		}

		// Get the value from the form data
		// 폼 데이터에서 값 가져오기
		formValue := values.Get(tag)
		if formValue == "" {
			// Skip if no value provided
			// 값이 제공되지 않은 경우 건너뛰기
			continue
		}

		// Set the field value based on its type
		// 타입에 따라 필드 값 설정
		if err := setFieldValue(field, formValue); err != nil {
			return fmt.Errorf("failed to set field %s: %w", typeField.Name, err)
		}
	}

	return nil
}

// setFieldValue sets the value of a reflect.Value based on its kind.
// setFieldValue는 reflect.Value의 종류에 따라 값을 설정합니다.
func setFieldValue(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot convert value %q to %s: %w", value, field.Type(), err)
		}
		field.SetInt(intValue)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot convert value %q to %s: %w", value, field.Type(), err)
		}
		field.SetUint(uintValue)

	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("cannot convert value %q to %s: %w", value, field.Type(), err)
		}
		field.SetFloat(floatValue)

	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("cannot convert value %q to %s (expected true/false, 1/0, t/f): %w", value, field.Type(), err)
		}
		field.SetBool(boolValue)

	default:
		return fmt.Errorf("unsupported field type %s (supported: string, int, uint, float, bool)", field.Type())
	}

	return nil
}
