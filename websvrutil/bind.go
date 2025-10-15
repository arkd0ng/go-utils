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
// The struct fields should use `form` tags to specify the form field names.
// 구조체 필드는 `form` 태그를 사용하여 폼 필드 이름을 지정해야 합니다.
func bindFormData(obj interface{}, values url.Values) error {
	// Get the reflect value and type of the object
	// 객체의 리플렉트 값과 타입 가져오기
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("binding requires a pointer to a struct")
	}

	// Dereference the pointer
	// 포인터 역참조
	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("binding requires a pointer to a struct")
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
			return fmt.Errorf("failed to parse int: %w", err)
		}
		field.SetInt(intValue)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse uint: %w", err)
		}
		field.SetUint(uintValue)

	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("failed to parse float: %w", err)
		}
		field.SetFloat(floatValue)

	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("failed to parse bool: %w", err)
		}
		field.SetBool(boolValue)

	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}

	return nil
}
