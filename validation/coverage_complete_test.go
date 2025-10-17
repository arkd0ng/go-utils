package validation

import (
	"testing"
	"time"
)

// Additional tests to achieve 100% coverage
// 100% 커버리지 달성을 위한 추가 테스트

// TestInWithNilValue tests In validator with nil value not in list
// TestInWithNilValue는 목록에 없는 nil 값으로 In 검증기를 테스트합니다
func TestInWithNilValue(t *testing.T) {
	v := New(nil, "field")
	v.In("a", "b", "c") // nil is not in this list
	err := v.Validate()
	if err == nil {
		t.Error("expected error for nil value not in list, got nil")
	}
}

// TestInWithValueFound tests In validator when value is found in list (success case)
// TestInWithValueFound는 값이 목록에 있을 때 In 검증기를 테스트합니다 (성공 케이스)
func TestInWithValueFound(t *testing.T) {
	v := New("b", "field")
	v.In("a", "b", "c") // "b" is in the list
	err := v.Validate()
	if err != nil {
		t.Errorf("expected no error when value is in list, got %v", err)
	}
}

// TestNotInWithValueNotFound tests NotIn validator when value is not in list (success case)
// TestNotInWithValueNotFound는 값이 목록에 없을 때 NotIn 검증기를 테스트합니다 (성공 케이스)
func TestNotInWithValueNotFound(t *testing.T) {
	v := New("d", "field")
	v.NotIn("a", "b", "c") // "d" is not in the forbidden list
	err := v.Validate()
	if err != nil {
		t.Errorf("expected no error when value is not in forbidden list, got %v", err)
	}
}

// TestArrayUniqueWithUniqueElements tests ArrayUnique with all unique elements (success case)
// TestArrayUniqueWithUniqueElements는 모든 요소가 고유할 때 ArrayUnique를 테스트합니다 (성공 케이스)
func TestArrayUniqueWithUniqueElements(t *testing.T) {
	v := New([]int{1, 2, 3, 4, 5}, "numbers")
	v.ArrayUnique() // All elements are unique
	err := v.Validate()
	if err != nil {
		t.Errorf("expected no error for unique array, got %v", err)
	}
}

// TestStopOnErrorWithBeforeValidator tests stop-on-error with Before validator
// TestStopOnErrorWithBeforeValidator는 Before 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithBeforeValidator(t *testing.T) {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	yesterday := now.Add(-24 * time.Hour)

	v := New(tomorrow, "date")
	v.StopOnError()
	v.Before(now)        // First error - tomorrow is not before now
	v.Before(yesterday)  // Should hit stopOnError return in Before()
	err := v.Validate()
	if err == nil {
		t.Error("expected error from Before validator")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestNotInWithNilValueInList tests NotIn validator with nil value in forbidden list
// TestNotInWithNilValueInList는 금지 목록에 있는 nil 값으로 NotIn 검증기를 테스트합니다
func TestNotInWithNilValueInList(t *testing.T) {
	v := New(nil, "field")
	v.NotIn(nil, "a", "b") // nil is in the forbidden list
	err := v.Validate()
	if err == nil {
		t.Error("expected error for nil value in forbidden list, got nil")
	}
}

// TestArrayMinLengthWithNonArray tests ArrayMinLength with non-array value
// TestArrayMinLengthWithNonArray는 배열이 아닌 값으로 ArrayMinLength를 테스트합니다
func TestArrayMinLengthWithNonArray(t *testing.T) {
	v := New("not an array", "field")
	v.ArrayMinLength(2)
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-array value, got nil")
	}
}

// TestArrayMaxLengthWithNonArray tests ArrayMaxLength with non-array value
// TestArrayMaxLengthWithNonArray는 배열이 아닌 값으로 ArrayMaxLength를 테스트합니다
func TestArrayMaxLengthWithNonArray(t *testing.T) {
	v := New("not an array", "field")
	v.ArrayMaxLength(5)
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-array value, got nil")
	}
}

// TestMapHasKeysWithNonMap tests MapHasKeys with non-map value
// TestMapHasKeysWithNonMap은 맵이 아닌 값으로 MapHasKeys를 테스트합니다
func TestMapHasKeysWithNonMap(t *testing.T) {
	v := New("not a map", "field")
	v.MapHasKeys("key1", "key2")
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-map value, got nil")
	}
}

// TestMapHasKeysWithMissingKeys tests MapHasKeys with missing required keys
// TestMapHasKeysWithMissingKeys는 필수 키가 누락된 경우 MapHasKeys를 테스트합니다
func TestMapHasKeysWithMissingKeys(t *testing.T) {
	m := map[string]interface{}{
		"key1": "value1",
	}
	v := New(m, "field")
	v.MapHasKeys("key1", "key2", "key3")
	err := v.Validate()
	if err == nil {
		t.Error("expected error for missing keys, got nil")
	}
}

// TestEqualsWithDifferentTypes tests Equals with different types
// TestEqualsWithDifferentTypes는 다른 타입으로 Equals를 테스트합니다
func TestEqualsWithDifferentTypes(t *testing.T) {
	v := New("string", "field")
	v.Equals(123) // Different type
	err := v.Validate()
	if err == nil {
		t.Error("expected error for different types, got nil")
	}
}

// TestNotEqualsWithSameValue tests NotEquals with same value
// TestNotEqualsWithSameValue는 같은 값으로 NotEquals를 테스트합니다
func TestNotEqualsWithSameValue(t *testing.T) {
	v := New("same", "field")
	v.NotEquals("same")
	err := v.Validate()
	if err == nil {
		t.Error("expected error for equal values, got nil")
	}
}

// TestBeforeWithNonTimeValue tests Before with non-time value
// TestBeforeWithNonTimeValue는 시간이 아닌 값으로 Before를 테스트합니다
func TestBeforeWithNonTimeValue(t *testing.T) {
	v := New("not a time", "field")
	v.Before(time.Now())
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-time value, got nil")
	}
}

// TestAfterWithNonTimeValue tests After with non-time value
// TestAfterWithNonTimeValue는 시간이 아닌 값으로 After를 테스트합니다
func TestAfterWithNonTimeValue(t *testing.T) {
	v := New(123, "field")
	v.After(time.Now())
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-time value, got nil")
	}
}

// TestBeforeOrEqualWithNonTimeValue tests BeforeOrEqual with non-time value
// TestBeforeOrEqualWithNonTimeValue는 시간이 아닌 값으로 BeforeOrEqual을 테스트합니다
func TestBeforeOrEqualWithNonTimeValue(t *testing.T) {
	v := New([]int{1, 2, 3}, "field")
	v.BeforeOrEqual(time.Now())
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-time value, got nil")
	}
}

// TestAfterOrEqualWithNonTimeValue tests AfterOrEqual with non-time value
// TestAfterOrEqualWithNonTimeValue는 시간이 아닌 값으로 AfterOrEqual을 테스트합니다
func TestAfterOrEqualWithNonTimeValue(t *testing.T) {
	v := New(map[string]string{}, "field")
	v.AfterOrEqual(time.Now())
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-time value, got nil")
	}
}

// TestCustomValidatorWithError tests Custom validator returning false
// TestCustomValidatorWithError는 false를 반환하는 Custom 검증기를 테스트합니다
func TestCustomValidatorWithError(t *testing.T) {
	v := New("test", "field")
	v.Custom(func(val interface{}) bool {
		return false // Always fail
	}, "custom validation failed")

	err := v.Validate()
	if err == nil {
		t.Error("expected error from custom validator, got nil")
	}

	verrs, ok := err.(ValidationErrors)
	if !ok {
		t.Fatal("expected ValidationErrors type")
	}

	if len(verrs) != 1 {
		t.Errorf("expected 1 error, got %d", len(verrs))
	}

	if verrs[0].Message != "custom validation failed" {
		t.Errorf("got message %q, want %q", verrs[0].Message, "custom validation failed")
	}
}

// TestValidateStringWithNonStringValue tests validateString helper with non-string
// TestValidateStringWithNonStringValue는 문자열이 아닌 값으로 validateString 헬퍼를 테스트합니다
func TestValidateStringWithNonStringValue(t *testing.T) {
	v := New(123, "field")
	v.Required() // Calls validateString internally
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-string value in string validator, got nil")
	}
}

// TestValidateNumericWithUnsupportedType tests validateNumeric with unsupported type
// TestValidateNumericWithUnsupportedType은 지원하지 않는 타입으로 validateNumeric을 테스트합니다
func TestValidateNumericWithUnsupportedType(t *testing.T) {
	type customType struct{}
	v := New(customType{}, "field")
	v.Positive() // Calls validateNumeric internally
	err := v.Validate()
	if err == nil {
		t.Error("expected error for unsupported type in numeric validator, got nil")
	}
}

// TestArrayNotEmptyWithNonArray tests ArrayNotEmpty with non-array value
// TestArrayNotEmptyWithNonArray는 배열이 아닌 값으로 ArrayNotEmpty를 테스트합니다
func TestArrayNotEmptyWithNonArray(t *testing.T) {
	v := New(123, "field")
	v.ArrayNotEmpty()
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-array value, got nil")
	}
}

// TestArrayUniqueWithDifferentTypes tests ArrayUnique with various types
// TestArrayUniqueWithDifferentTypes는 다양한 타입으로 ArrayUnique를 테스트합니다
func TestArrayUniqueWithDifferentTypes(t *testing.T) {
	// Test with float slice
	floatSlice := []float64{1.1, 2.2, 1.1}
	v1 := New(floatSlice, "field")
	v1.ArrayUnique()
	err1 := v1.Validate()
	if err1 == nil {
		t.Error("expected error for duplicate floats, got nil")
	}

	// Test with string slice with duplicates
	stringSlice := []string{"a", "b", "a"}
	v2 := New(stringSlice, "field")
	v2.ArrayUnique()
	err2 := v2.Validate()
	if err2 == nil {
		t.Error("expected error for duplicate strings, got nil")
	}
}

// TestMapHasKeyWithNonMap tests MapHasKey with non-map value
// TestMapHasKeyWithNonMap은 맵이 아닌 값으로 MapHasKey를 테스트합니다
func TestMapHasKeyWithNonMap(t *testing.T) {
	v := New([]int{1, 2, 3}, "field")
	v.MapHasKey("key")
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-map value, got nil")
	}
}

// TestMapNotEmptyWithNonMap tests MapNotEmpty with non-map value
// TestMapNotEmptyWithNonMap은 맵이 아닌 값으로 MapNotEmpty를 테스트합니다
func TestMapNotEmptyWithNonMap(t *testing.T) {
	v := New("not a map", "field")
	v.MapNotEmpty()
	err := v.Validate()
	if err == nil {
		t.Error("expected error for non-map value, got nil")
	}
}

// TestEqualsWithNilValues tests Equals with nil values
// TestEqualsWithNilValues는 nil 값으로 Equals를 테스트합니다
func TestEqualsWithNilValues(t *testing.T) {
	// Both nil
	v1 := New(nil, "field")
	v1.Equals(nil)
	err1 := v1.Validate()
	if err1 != nil {
		t.Error("expected no error when both values are nil")
	}

	// One nil
	v2 := New(nil, "field")
	v2.Equals("not nil")
	err2 := v2.Validate()
	if err2 == nil {
		t.Error("expected error when comparing nil with non-nil, got nil")
	}
}

// TestNotEqualsWithNilValues tests NotEquals with nil values
// TestNotEqualsWithNilValues는 nil 값으로 NotEquals를 테스트합니다
func TestNotEqualsWithNilValues(t *testing.T) {
	// Both nil - should fail
	v1 := New(nil, "field")
	v1.NotEquals(nil)
	err1 := v1.Validate()
	if err1 == nil {
		t.Error("expected error when both values are nil, got nil")
	}

	// One nil - should pass
	v2 := New(nil, "field")
	v2.NotEquals("not nil")
	err2 := v2.Validate()
	if err2 != nil {
		t.Errorf("expected no error when comparing nil with non-nil, got %v", err2)
	}
}

// TestCustomValidatorReturnsTrueEdgeCase tests Custom validator with true return
// TestCustomValidatorReturnsTrueEdgeCase는 true를 반환하는 Custom 검증기를 테스트합니다
func TestCustomValidatorReturnsTrueEdgeCase(t *testing.T) {
	v := New("test", "field")
	v.Custom(func(val interface{}) bool {
		return true // Always pass
	}, "should not appear")

	err := v.Validate()
	if err != nil {
		t.Errorf("expected no error from passing custom validator, got %v", err)
	}
}

// TestValidateStringWithEmptyString tests validateString with empty string
// TestValidateStringWithEmptyString은 빈 문자열로 validateString을 테스트합니다
func TestValidateStringWithEmptyString(t *testing.T) {
	v := New("", "field")
	v.MinLength(1) // Should fail for empty string
	err := v.Validate()
	if err == nil {
		t.Error("expected error for empty string with MinLength(1), got nil")
	}
}

// TestValidateNumericWithAllNumericTypes tests validateNumeric with all supported types
// TestValidateNumericWithAllNumericTypes는 모든 지원되는 타입으로 validateNumeric을 테스트합니다
func TestValidateNumericWithAllNumericTypes(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"int8", int8(10)},
		{"int16", int16(10)},
		{"int32", int32(10)},
		{"uint", uint(10)},
		{"uint8", uint8(10)},
		{"uint16", uint16(10)},
		{"uint32", uint32(10)},
		{"uint64", uint64(10)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Positive()
			err := v.Validate()
			if err != nil {
				t.Errorf("expected no error for positive %s, got %v", tt.name, err)
			}
		})
	}
}

// TestStopOnErrorWithArrayValidators tests stop-on-error behavior with array validators
// TestStopOnErrorWithArrayValidators는 배열 검증기에서 stop-on-error 동작을 테스트합니다
func TestStopOnErrorWithArrayValidators(t *testing.T) {
	// Test with ArrayMinLength - should stop after first error
	v1 := New("not an array", "field")
	v1.StopOnError()
	v1.ArrayMinLength(2)          // First error - not an array
	v1.ArrayMaxLength(5)          // Should not execute due to StopOnError
	err1 := v1.Validate()
	if err1 == nil {
		t.Error("expected error from ArrayMinLength")
	}
	verrs1 := err1.(ValidationErrors)
	if len(verrs1) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs1))
	}

	// Test with ArrayMaxLength
	v2 := New(123, "field")
	v2.StopOnError()
	v2.ArrayMaxLength(5)          // First error - not an array
	v2.ArrayNotEmpty()            // Should not execute
	err2 := v2.Validate()
	if err2 == nil {
		t.Error("expected error from ArrayMaxLength")
	}
	verrs2 := err2.(ValidationErrors)
	if len(verrs2) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs2))
	}
}

// TestStopOnErrorWithMapValidators tests stop-on-error behavior with map validators
// TestStopOnErrorWithMapValidators는 맵 검증기에서 stop-on-error 동작을 테스트합니다
func TestStopOnErrorWithMapValidators(t *testing.T) {
	m := map[string]interface{}{
		"key1": "value1",
	}
	v := New(m, "field")
	v.StopOnError()
	v.MapHasKeys("key1", "key2", "key3") // First error - missing keys
	v.MapNotEmpty()                       // Should not execute
	err := v.Validate()
	if err == nil {
		t.Error("expected error from MapHasKeys")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithComparisonValidators tests stop-on-error with comparison validators
// TestStopOnErrorWithComparisonValidators는 비교 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithComparisonValidators(t *testing.T) {
	v := New("test", "field")
	v.StopOnError()
	v.Equals("different")  // First error
	v.NotEquals("another") // Should not execute
	err := v.Validate()
	if err == nil {
		t.Error("expected error from Equals")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithTimeComparisons tests stop-on-error with time comparisons
// TestStopOnErrorWithTimeComparisons는 시간 비교에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithTimeComparisons(t *testing.T) {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)

	v := New(tomorrow, "field")
	v.StopOnError()
	v.Before(now)        // First error - tomorrow is not before now
	v.After(now.Add(-48 * time.Hour)) // Should not execute
	err := v.Validate()
	if err == nil {
		t.Error("expected error from Before")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithCustomValidator tests stop-on-error with custom validator
// TestStopOnErrorWithCustomValidator는 사용자 정의 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithCustomValidator(t *testing.T) {
	v := New("test", "field")
	v.StopOnError()
	v.Custom(func(val interface{}) bool {
		return false // Fail
	}, "custom error")
	v.Custom(func(val interface{}) bool {
		return false // Should hit stopOnError return in Custom()
	}, "second custom error")
	err := v.Validate()
	if err == nil {
		t.Error("expected error from Custom validator")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithStringValidators tests stop-on-error with string validators
// TestStopOnErrorWithStringValidators는 문자열 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithStringValidators(t *testing.T) {
	v := New(123, "field") // Not a string
	v.StopOnError()
	v.MinLength(5)  // First error - not a string
	v.MaxLength(10) // Should not execute
	err := v.Validate()
	if err == nil {
		t.Error("expected error from MinLength")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithNumericValidators tests stop-on-error with numeric validators
// TestStopOnErrorWithNumericValidators는 숫자 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithNumericValidators(t *testing.T) {
	type customType struct{}
	v := New(customType{}, "field") // Not a number
	v.StopOnError()
	v.Positive() // First error - not a number
	v.Min(0)     // Should not execute
	err := v.Validate()
	if err == nil {
		t.Error("expected error from Positive")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithInValidator tests stop-on-error with In validator
// TestStopOnErrorWithInValidator는 In 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithInValidator(t *testing.T) {
	v := New("yellow", "color")
	v.StopOnError()
	v.In("red", "green", "blue")    // First error - not in list
	v.In("purple", "orange", "pink") // Should hit stopOnError return in In()
	err := v.Validate()
	if err == nil {
		t.Error("expected error from In validator")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithNotInValidator tests stop-on-error with NotIn validator
// TestStopOnErrorWithNotInValidator는 NotIn 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithNotInValidator(t *testing.T) {
	v := New("red", "color")
	v.StopOnError()
	v.NotIn("red", "green")  // First error - value is in forbidden list
	v.NotIn("blue", "yellow") // Should hit stopOnError return in NotIn()
	err := v.Validate()
	if err == nil {
		t.Error("expected error from NotIn validator")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithArrayLength tests stop-on-error with ArrayLength validator
// TestStopOnErrorWithArrayLength는 ArrayLength 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithArrayLength(t *testing.T) {
	v := New([]int{1, 2}, "numbers")
	v.StopOnError()
	v.ArrayLength(3) // First error - wrong length
	v.ArrayLength(5) // Should hit stopOnError return in ArrayLength()
	err := v.Validate()
	if err == nil {
		t.Error("expected error from ArrayLength")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithArrayNotEmpty tests stop-on-error with ArrayNotEmpty validator
// TestStopOnErrorWithArrayNotEmpty는 ArrayNotEmpty 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithArrayNotEmpty(t *testing.T) {
	v := New([]int{}, "numbers")
	v.StopOnError()
	v.ArrayNotEmpty() // First error - empty array
	v.ArrayMinLength(1) // Should not execute
	err := v.Validate()
	if err == nil {
		t.Error("expected error from ArrayNotEmpty")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithArrayUnique tests stop-on-error with ArrayUnique validator
// TestStopOnErrorWithArrayUnique는 ArrayUnique 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithArrayUnique(t *testing.T) {
	v := New([]int{1, 2, 1}, "numbers")
	v.StopOnError()
	v.ArrayUnique()     // First error - duplicate values
	v.ArrayUnique()     // Should hit stopOnError return in ArrayUnique()
	err := v.Validate()
	if err == nil {
		t.Error("expected error from ArrayUnique")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithMapHasKey tests stop-on-error with MapHasKey validator
// TestStopOnErrorWithMapHasKey는 MapHasKey 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithMapHasKey(t *testing.T) {
	m := map[string]string{"key1": "value1"}
	v := New(m, "data")
	v.StopOnError()
	v.MapHasKey("key2") // First error - missing key
	v.MapHasKeys("key3", "key4") // Should hit stopOnError return in MapHasKeys()
	err := v.Validate()
	if err == nil {
		t.Error("expected error from MapHasKey")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithBeforeOrEqual tests stop-on-error with BeforeOrEqual validator
// TestStopOnErrorWithBeforeOrEqual는 BeforeOrEqual 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithBeforeOrEqual(t *testing.T) {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	yesterday := now.Add(-24 * time.Hour)

	v := New(tomorrow, "date")
	v.StopOnError()
	v.BeforeOrEqual(now) // First error - tomorrow is after now
	v.BeforeOrEqual(yesterday) // Should hit stopOnError return in BeforeOrEqual()
	err := v.Validate()
	if err == nil {
		t.Error("expected error from BeforeOrEqual")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithAfterOrEqual tests stop-on-error with AfterOrEqual validator
// TestStopOnErrorWithAfterOrEqual는 AfterOrEqual 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithAfterOrEqual(t *testing.T) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)

	v := New(yesterday, "date")
	v.StopOnError()
	v.AfterOrEqual(now) // First error - yesterday is before now
	v.AfterOrEqual(tomorrow) // Should hit stopOnError return in AfterOrEqual()
	err := v.Validate()
	if err == nil {
		t.Error("expected error from AfterOrEqual")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithEquals tests stop-on-error with Equals validator
// TestStopOnErrorWithEquals는 Equals 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithEquals(t *testing.T) {
	v := New(10, "value")
	v.StopOnError()
	v.Equals(20)    // First error - not equal
	v.Equals(30) // Should hit stopOnError return in Equals()
	err := v.Validate()
	if err == nil {
		t.Error("expected error from Equals")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestStopOnErrorWithMapNotEmpty tests stop-on-error with MapNotEmpty validator
// TestStopOnErrorWithMapNotEmpty는 MapNotEmpty 검증기에서 stop-on-error를 테스트합니다
func TestStopOnErrorWithMapNotEmpty(t *testing.T) {
	m := map[string]string{}
	v := New(m, "data")
	v.StopOnError()
	v.MapNotEmpty()     // First error - empty map
	v.MapHasKey("key1") // Should not execute
	err := v.Validate()
	if err == nil {
		t.Error("expected error from MapNotEmpty")
	}
	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}
