package validation

import (
	"fmt"
	"time"
)

// Equals checks if the value equals the given value.
// Equals는 값이 주어진 값과 같은지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New("password123", "password")
//	v.Equals("password123")
func (v *Validator) Equals(value interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if v.value != value {
		v.addError("equals", fmt.Sprintf("%s must equal '%v' / %s은(는) '%v'와 같아야 합니다", v.fieldName, value, v.fieldName, value))
	}

	return v
}

// NotEquals checks if the value does not equal the given value.
// NotEquals는 값이 주어진 값과 다른지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New("newpassword", "password")
//	v.NotEquals("oldpassword")
func (v *Validator) NotEquals(value interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if v.value == value {
		v.addError("notequals", fmt.Sprintf("%s must not equal '%v' / %s은(는) '%v'와 달라야 합니다", v.fieldName, value, v.fieldName, value))
	}

	return v
}

// GreaterThan checks if the numeric value is greater than the given value.
// GreaterThan은 숫자 값이 주어진 값보다 큰지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New(10, "score")
//	v.GreaterThan(5)
func (v *Validator) GreaterThan(value float64) *Validator {
	return validateNumeric(v, "greaterthan", func(n float64) bool {
		return n > value
	}, fmt.Sprintf("%s must be greater than %v / %s은(는) %v보다 커야 합니다", v.fieldName, value, v.fieldName, value))
}

// GreaterThanOrEqual checks if the numeric value is greater than or equal to the given value.
// GreaterThanOrEqual은 숫자 값이 주어진 값보다 크거나 같은지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New(10, "score")
//	v.GreaterThanOrEqual(10)
func (v *Validator) GreaterThanOrEqual(value float64) *Validator {
	return validateNumeric(v, "greaterthanorequal", func(n float64) bool {
		return n >= value
	}, fmt.Sprintf("%s must be greater than or equal to %v / %s은(는) %v보다 크거나 같아야 합니다", v.fieldName, value, v.fieldName, value))
}

// LessThan checks if the numeric value is less than the given value.
// LessThan은 숫자 값이 주어진 값보다 작은지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New(5, "score")
//	v.LessThan(10)
func (v *Validator) LessThan(value float64) *Validator {
	return validateNumeric(v, "lessthan", func(n float64) bool {
		return n < value
	}, fmt.Sprintf("%s must be less than %v / %s은(는) %v보다 작아야 합니다", v.fieldName, value, v.fieldName, value))
}

// LessThanOrEqual checks if the numeric value is less than or equal to the given value.
// LessThanOrEqual은 숫자 값이 주어진 값보다 작거나 같은지 확인합니다.
//
// Example / 예제:
//
//	v := validation.New(10, "score")
//	v.LessThanOrEqual(10)
func (v *Validator) LessThanOrEqual(value float64) *Validator {
	return validateNumeric(v, "lessthanorequal", func(n float64) bool {
		return n <= value
	}, fmt.Sprintf("%s must be less than or equal to %v / %s은(는) %v보다 작거나 같아야 합니다", v.fieldName, value, v.fieldName, value))
}

// Before checks if the time value is before the given time.
// Before는 시간 값이 주어진 시간보다 이전인지 확인합니다.
//
// Example / 예제:
//
//	now := time.Now()
//	yesterday := now.Add(-24 * time.Hour)
//	v := validation.New(yesterday, "date")
//	v.Before(now)
func (v *Validator) Before(t time.Time) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	timeVal, ok := v.value.(time.Time)
	if !ok {
		v.addError("before", fmt.Sprintf("%s must be a time.Time value / %s은(는) time.Time 값이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if !timeVal.Before(t) {
		v.addError("before", fmt.Sprintf("%s must be before %s / %s은(는) %s 이전이어야 합니다", v.fieldName, t.Format(time.RFC3339), v.fieldName, t.Format(time.RFC3339)))
	}

	return v
}

// After checks if the time value is after the given time.
// After는 시간 값이 주어진 시간보다 이후인지 확인합니다.
//
// Example / 예제:
//
//	now := time.Now()
//	tomorrow := now.Add(24 * time.Hour)
//	v := validation.New(tomorrow, "date")
//	v.After(now)
func (v *Validator) After(t time.Time) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	timeVal, ok := v.value.(time.Time)
	if !ok {
		v.addError("after", fmt.Sprintf("%s must be a time.Time value / %s은(는) time.Time 값이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if !timeVal.After(t) {
		v.addError("after", fmt.Sprintf("%s must be after %s / %s은(는) %s 이후여야 합니다", v.fieldName, t.Format(time.RFC3339), v.fieldName, t.Format(time.RFC3339)))
	}

	return v
}

// BeforeOrEqual checks if the time value is before or equal to the given time.
// BeforeOrEqual은 시간 값이 주어진 시간보다 이전이거나 같은지 확인합니다.
//
// Example / 예제:
//
//	now := time.Now()
//	v := validation.New(now, "date")
//	v.BeforeOrEqual(now)
func (v *Validator) BeforeOrEqual(t time.Time) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	timeVal, ok := v.value.(time.Time)
	if !ok {
		v.addError("beforeorequal", fmt.Sprintf("%s must be a time.Time value / %s은(는) time.Time 값이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if timeVal.After(t) {
		v.addError("beforeorequal", fmt.Sprintf("%s must be before or equal to %s / %s은(는) %s 이전이거나 같아야 합니다", v.fieldName, t.Format(time.RFC3339), v.fieldName, t.Format(time.RFC3339)))
	}

	return v
}

// AfterOrEqual checks if the time value is after or equal to the given time.
// AfterOrEqual은 시간 값이 주어진 시간보다 이후이거나 같은지 확인합니다.
//
// Example / 예제:
//
//	now := time.Now()
//	v := validation.New(now, "date")
//	v.AfterOrEqual(now)
func (v *Validator) AfterOrEqual(t time.Time) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	timeVal, ok := v.value.(time.Time)
	if !ok {
		v.addError("afterorequal", fmt.Sprintf("%s must be a time.Time value / %s은(는) time.Time 값이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if timeVal.Before(t) {
		v.addError("afterorequal", fmt.Sprintf("%s must be after or equal to %s / %s은(는) %s 이후이거나 같아야 합니다", v.fieldName, t.Format(time.RFC3339), v.fieldName, t.Format(time.RFC3339)))
	}

	return v
}
