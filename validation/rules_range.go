package validation

import (
	"fmt"
	"time"
)

// IntRange checks if the integer value is within the specified range (inclusive).
// IntRange는 정수 값이 지정된 범위 내에 있는지 확인합니다 (포함).
//
// The range is inclusive: min <= value <= max
//
// Example:
//
//	v := validation.New(25, "age")
//	v.IntRange(18, 65)
//	err := v.Validate()
func (v *Validator) IntRange(min, max int) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	// Convert value to int64 using helper function
	intVal, err := v.toInt64(v.value)
	if err != nil {
		v.addError("int_range", fmt.Sprintf("%s must be an integer / %s은(는) 정수여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if intVal < int64(min) || intVal > int64(max) {
		v.addError("int_range", fmt.Sprintf("%s must be between %d and %d / %s은(는) %d와 %d 사이여야 합니다", v.fieldName, min, max, v.fieldName, min, max))
		return v
	}

	return v
}

// FloatRange checks if the float value is within the specified range (inclusive).
// FloatRange는 실수 값이 지정된 범위 내에 있는지 확인합니다 (포함).
//
// The range is inclusive: min <= value <= max
//
// Example:
//
//	v := validation.New(98.6, "temperature")
//	v.FloatRange(95.0, 105.0)
//	err := v.Validate()
func (v *Validator) FloatRange(min, max float64) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	// Convert value to float64 using helper function
	floatVal, err := v.toFloat64(v.value)
	if err != nil {
		v.addError("float_range", fmt.Sprintf("%s must be a number / %s은(는) 숫자여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if floatVal < min || floatVal > max {
		v.addError("float_range", fmt.Sprintf("%s must be between %.2f and %.2f / %s은(는) %.2f와 %.2f 사이여야 합니다", v.fieldName, min, max, v.fieldName, min, max))
		return v
	}

	return v
}

// DateRange checks if the date value is within the specified range (inclusive).
// DateRange는 날짜 값이 지정된 범위 내에 있는지 확인합니다 (포함).
//
// The range is inclusive: start <= value <= end
//
// The value can be:
//   - time.Time object
//   - string in RFC3339 format ("2006-01-02T15:04:05Z07:00")
//   - string in ISO 8601 format ("2006-01-02")
//
// Example:
//
//	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
//	v := validation.New(eventDate, "event_date")
//	v.DateRange(start, end)
//	err := v.Validate()
func (v *Validator) DateRange(start, end time.Time) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	var targetTime time.Time
	var parseErr error

	switch val := v.value.(type) {
	case time.Time:
		targetTime = val
	case string:
		// Try RFC3339 first
		targetTime, parseErr = time.Parse(time.RFC3339, val)
		if parseErr != nil {
			// Try ISO 8601 date format
			targetTime, parseErr = time.Parse("2006-01-02", val)
			if parseErr != nil {
				v.addError("date_range", fmt.Sprintf("%s must be a valid date (time.Time or RFC3339/ISO 8601 string) / %s은(는) 유효한 날짜여야 합니다 (time.Time 또는 RFC3339/ISO 8601 문자열)", v.fieldName, v.fieldName))
				return v
			}
		}
	default:
		v.addError("date_range", fmt.Sprintf("%s must be a time.Time or date string / %s은(는) time.Time 또는 날짜 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if targetTime.Before(start) || targetTime.After(end) {
		v.addError("date_range", fmt.Sprintf("%s must be between %s and %s / %s은(는) %s와 %s 사이여야 합니다", v.fieldName, start.Format("2006-01-02"), end.Format("2006-01-02"), v.fieldName, start.Format("2006-01-02"), end.Format("2006-01-02")))
		return v
	}

	return v
}

// Helper function to convert value to int64
func (v *Validator) toInt64(val interface{}) (int64, error) {
	switch value := val.(type) {
	case int:
		return int64(value), nil
	case int8:
		return int64(value), nil
	case int16:
		return int64(value), nil
	case int32:
		return int64(value), nil
	case int64:
		return value, nil
	case uint:
		return int64(value), nil
	case uint8:
		return int64(value), nil
	case uint16:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case uint64:
		return int64(value), nil
	default:
		return 0, fmt.Errorf("not an integer")
	}
}

// Helper function to convert value to float64
func (v *Validator) toFloat64(val interface{}) (float64, error) {
	switch value := val.(type) {
	case float32:
		return float64(value), nil
	case float64:
		return value, nil
	case int:
		return float64(value), nil
	case int8:
		return float64(value), nil
	case int16:
		return float64(value), nil
	case int32:
		return float64(value), nil
	case int64:
		return float64(value), nil
	case uint:
		return float64(value), nil
	case uint8:
		return float64(value), nil
	case uint16:
		return float64(value), nil
	case uint32:
		return float64(value), nil
	case uint64:
		return float64(value), nil
	default:
		return 0, fmt.Errorf("not a number")
	}
}
