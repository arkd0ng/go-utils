package validation

import (
	"fmt"
	"time"
)

// DateFormat checks if the value is a valid date string in the specified format.
// DateFormat은 값이 지정된 형식의 유효한 날짜 문자열인지 확인합니다.
//
// Supported formats:
//   - "2006-01-02" (ISO 8601 date)
//   - "02/01/2006" (DD/MM/YYYY)
//   - "01/02/2006" (MM/DD/YYYY)
//   - "2006/01/02" (YYYY/MM/DD)
//   - Any valid Go time.Parse format
//
// Example:
//
//	v := validation.New("2025-10-17", "birth_date")
//	v.DateFormat("2006-01-02")
//	err := v.Validate()
func (v *Validator) DateFormat(format string) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("date_format", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	_, err := time.Parse(format, str)
	if err != nil {
		v.addError("date_format", fmt.Sprintf("%s must be a valid date in format %s / %s은(는) %s 형식의 유효한 날짜여야 합니다", v.fieldName, format, v.fieldName, format))
		return v
	}

	return v
}

// TimeFormat checks if the value is a valid time string in the specified format.
// TimeFormat은 값이 지정된 형식의 유효한 시간 문자열인지 확인합니다.
//
// Supported formats:
//   - "15:04:05" (HH:MM:SS)
//   - "15:04" (HH:MM)
//   - "03:04:05 PM" (12-hour with seconds)
//   - "03:04 PM" (12-hour)
//   - Any valid Go time.Parse format
//
// Example:
//
//	v := validation.New("14:30:00", "meeting_time")
//	v.TimeFormat("15:04:05")
//	err := v.Validate()
func (v *Validator) TimeFormat(format string) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("time_format", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	_, err := time.Parse(format, str)
	if err != nil {
		v.addError("time_format", fmt.Sprintf("%s must be a valid time in format %s / %s은(는) %s 형식의 유효한 시간이어야 합니다", v.fieldName, format, v.fieldName, format))
		return v
	}

	return v
}

// DateBefore checks if the date value is before the specified date.
// DateBefore는 날짜 값이 지정된 날짜 이전인지 확인합니다.
//
// The value can be:
//   - time.Time object
//   - string in RFC3339 format ("2006-01-02T15:04:05Z07:00")
//   - string in ISO 8601 format ("2006-01-02")
//
// Example:
//
//	maxDate := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)
//	v := validation.New(userDate, "expiry_date")
//	v.DateBefore(maxDate)
//	err := v.Validate()
func (v *Validator) DateBefore(before time.Time) *Validator {
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
				v.addError("date_before", fmt.Sprintf("%s must be a valid date (time.Time or RFC3339/ISO 8601 string) / %s은(는) 유효한 날짜여야 합니다 (time.Time 또는 RFC3339/ISO 8601 문자열)", v.fieldName, v.fieldName))
				return v
			}
		}
	default:
		v.addError("date_before", fmt.Sprintf("%s must be a time.Time or date string / %s은(는) time.Time 또는 날짜 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if !targetTime.Before(before) {
		v.addError("date_before", fmt.Sprintf("%s must be before %s / %s은(는) %s 이전이어야 합니다", v.fieldName, before.Format("2006-01-02"), v.fieldName, before.Format("2006-01-02")))
		return v
	}

	return v
}

// DateAfter checks if the date value is after the specified date.
// DateAfter는 날짜 값이 지정된 날짜 이후인지 확인합니다.
//
// The value can be:
//   - time.Time object
//   - string in RFC3339 format ("2006-01-02T15:04:05Z07:00")
//   - string in ISO 8601 format ("2006-01-02")
//
// Example:
//
//	minDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	v := validation.New(userDate, "start_date")
//	v.DateAfter(minDate)
//	err := v.Validate()
func (v *Validator) DateAfter(after time.Time) *Validator {
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
				v.addError("date_after", fmt.Sprintf("%s must be a valid date (time.Time or RFC3339/ISO 8601 string) / %s은(는) 유효한 날짜여야 합니다 (time.Time 또는 RFC3339/ISO 8601 문자열)", v.fieldName, v.fieldName))
				return v
			}
		}
	default:
		v.addError("date_after", fmt.Sprintf("%s must be a time.Time or date string / %s은(는) time.Time 또는 날짜 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if !targetTime.After(after) {
		v.addError("date_after", fmt.Sprintf("%s must be after %s / %s은(는) %s 이후여야 합니다", v.fieldName, after.Format("2006-01-02"), v.fieldName, after.Format("2006-01-02")))
		return v
	}

	return v
}
