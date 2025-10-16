package websvrutil

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// bindFormData binds URL values to a struct using reflection.
// bindFormData는 리플렉션을 사용하여 URL 값을 구조체에 바인딩합니다.
//
// Process / 프로세스:
//   1. Validate obj is a pointer to struct (returns error if not)
//   2. Iterate through struct fields
//   3. Extract "form" tag for field name mapping (uses field name if tag absent)
//   4. Retrieve corresponding value from URL values
//   5. Convert string value to appropriate type using strconv
//   6. Set field value using reflection
//
// Supported types / 지원되는 타입:
//   - string: Direct assignment
//   - int, int8, int16, int32, int64: ParseInt with base 10
//   - uint, uint8, uint16, uint32, uint64: ParseUint with base 10
//   - float32, float64: ParseFloat
//   - bool: ParseBool (accepts "1", "t", "T", "true", "TRUE", "True", "0", "f", "F", "false", "FALSE", "False")
//
// Struct tag format / 구조체 태그 형식:
//   type User struct {
//       Name  string `form:"name"`      // Maps to "name" form field
//       Email string `form:"email"`     // Maps to "email" form field
//       Age   int    `form:"age"`       // Maps to "age" form field
//       Admin bool   // Maps to "Admin" (no tag, uses field name)
//   }
//
// Error handling / 에러 처리:
//   - Returns error if obj is not a pointer to struct
//   - Returns error if type conversion fails (e.g., "abc" -> int)
//   - Skips fields that:
//     - Are unexported (cannot be set)
//     - Have no corresponding form value (empty string)
//
// Example usage / 사용 예제:
//   values := url.Values{
//       "name":  []string{"John"},
//       "age":   []string{"30"},
//       "admin": []string{"true"},
//   }
//   var user User
//   err := bindFormData(&user, values)
//   // user.Name = "John", user.Age = 30, user.Admin = true
//
// Performance / 성능:
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
