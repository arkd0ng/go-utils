package validation

import (
	"fmt"
	"reflect"
)

// In checks if the value exists in the given list.
// In은 값이 주어진 목록에 존재하는지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New("red", "color")
//	v.In("red", "green", "blue")
func (v *Validator) In(values ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	found := false
	for _, val := range values {
		if reflect.DeepEqual(v.value, val) {
			found = true
			break
		}
	}

	if !found {
		v.addError("in", fmt.Sprintf("%s must be one of the allowed values / %s은(는) 허용된 값 중 하나여야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// NotIn checks if the value does not exist in the given list.
// NotIn은 값이 주어진 목록에 존재하지 않는지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New("admin", "role")
//	v.NotIn("guest", "anonymous")
func (v *Validator) NotIn(values ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	found := false
	for _, val := range values {
		if reflect.DeepEqual(v.value, val) {
			found = true
			break
		}
	}

	if found {
		v.addError("notin", fmt.Sprintf("%s must not be one of the forbidden values / %s은(는) 금지된 값이 아니어야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// ArrayLength checks if the array/slice has the exact length.
// ArrayLength는 배열/슬라이스가 정확한 길이를 가지는지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New([]int{1, 2, 3}, "numbers")
//	v.ArrayLength(3)
func (v *Validator) ArrayLength(length int) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arraylength", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() != length {
		v.addError("arraylength", fmt.Sprintf("%s length must be exactly %d / %s 길이는 정확히 %d여야 합니다", v.fieldName, length, v.fieldName, length))
	}

	return v
}

// ArrayMinLength checks if the array/slice has at least the minimum length.
// ArrayMinLength는 배열/슬라이스가 최소 길이를 가지는지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New([]string{"a", "b", "c"}, "items")
//	v.ArrayMinLength(2)
func (v *Validator) ArrayMinLength(min int) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arrayminlength", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() < min {
		v.addError("arrayminlength", fmt.Sprintf("%s length must be at least %d / %s 길이는 최소 %d여야 합니다", v.fieldName, min, v.fieldName, min))
	}

	return v
}

// ArrayMaxLength checks if the array/slice has at most the maximum length.
// ArrayMaxLength는 배열/슬라이스가 최대 길이를 넘지 않는지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New([]int{1, 2, 3}, "numbers")
//	v.ArrayMaxLength(5)
func (v *Validator) ArrayMaxLength(max int) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arraymaxlength", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() > max {
		v.addError("arraymaxlength", fmt.Sprintf("%s length must be at most %d / %s 길이는 최대 %d여야 합니다", v.fieldName, max, v.fieldName, max))
	}

	return v
}

// ArrayNotEmpty checks if the array/slice is not empty.
// ArrayNotEmpty는 배열/슬라이스가 비어있지 않은지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New([]string{"item"}, "items")
//	v.ArrayNotEmpty()
func (v *Validator) ArrayNotEmpty() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arraynotempty", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() == 0 {
		v.addError("arraynotempty", fmt.Sprintf("%s must not be empty / %s은(는) 비어있지 않아야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// ArrayUnique checks if all elements in the array/slice are unique.
// ArrayUnique는 배열/슬라이스의 모든 요소가 고유한지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New([]int{1, 2, 3}, "numbers")
//	v.ArrayUnique()
func (v *Validator) ArrayUnique() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arrayunique", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	seen := make(map[interface{}]bool)
	for i := 0; i < val.Len(); i++ {
		item := val.Index(i).Interface()
		if seen[item] {
			v.addError("arrayunique", fmt.Sprintf("%s must contain only unique elements / %s은(는) 고유한 요소만 포함해야 합니다", v.fieldName, v.fieldName))
			return v
		}
		seen[item] = true
	}

	return v
}

// MapHasKey checks if the map contains the specified key.
// MapHasKey는 맵이 지정된 키를 포함하는지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New(map[string]int{"age": 25}, "data")
//	v.MapHasKey("age")
func (v *Validator) MapHasKey(key interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Map {
		v.addError("maphaskey", fmt.Sprintf("%s must be a map / %s은(는) 맵이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	keyVal := reflect.ValueOf(key)
	if !val.MapIndex(keyVal).IsValid() {
		v.addError("maphaskey", fmt.Sprintf("%s must contain key '%v' / %s은(는) 키 '%v'를 포함해야 합니다", v.fieldName, key, v.fieldName, key))
	}

	return v
}

// MapHasKeys checks if the map contains all specified keys.
// MapHasKeys는 맵이 지정된 모든 키를 포함하는지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New(map[string]string{"name": "John", "email": "j@e.com"}, "user")
//	v.MapHasKeys("name", "email")
func (v *Validator) MapHasKeys(keys ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Map {
		v.addError("maphaskeys", fmt.Sprintf("%s must be a map / %s은(는) 맵이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	var missingKeys []interface{}
	for _, key := range keys {
		keyVal := reflect.ValueOf(key)
		if !val.MapIndex(keyVal).IsValid() {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) > 0 {
		v.addError("maphaskeys", fmt.Sprintf("%s is missing required keys: %v / %s에 필수 키가 없습니다: %v", v.fieldName, missingKeys, v.fieldName, missingKeys))
	}

	return v
}

// MapNotEmpty checks if the map is not empty.
// MapNotEmpty는 맵이 비어있지 않은지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New(map[string]int{"count": 1}, "data")
//	v.MapNotEmpty()
func (v *Validator) MapNotEmpty() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Map {
		v.addError("mapnotempty", fmt.Sprintf("%s must be a map / %s은(는) 맵이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() == 0 {
		v.addError("mapnotempty", fmt.Sprintf("%s must not be empty / %s은(는) 비어있지 않아야 합니다", v.fieldName, v.fieldName))
	}

	return v
}
