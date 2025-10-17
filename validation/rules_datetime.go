package validation

import (
	"fmt"
	"time"
)

// DateFormat validates that the value is a valid date string in the specified format.
// Uses Go's time.Parse with the provided layout to verify date string validity.
//
// DateFormat은 값이 지정된 형식의 유효한 날짜 문자열인지 검증합니다.
// 제공된 레이아웃으로 Go의 time.Parse를 사용하여 날짜 문자열 유효성을 확인합니다.
//
// Parameters / 매개변수:
//   - format: Go time layout format string
//     Go 시간 레이아웃 형식 문자열
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses time.Parse for validation
//     time.Parse를 사용하여 검증
//   - Accepts any Go time layout format
//     모든 Go 시간 레이아웃 형식 허용
//   - Common formats:
//     일반 형식:
//   - "2006-01-02" - ISO 8601 date (YYYY-MM-DD)
//   - "02/01/2006" - DD/MM/YYYY
//   - "01/02/2006" - MM/DD/YYYY
//   - "2006/01/02" - YYYY/MM/DD
//   - "Jan 2, 2006" - Month Day, Year
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//   - Fails if parsing fails
//     파싱 실패 시 실패
//
// Use Cases / 사용 사례:
//   - Date input validation / 날짜 입력 검증
//   - Form date field validation / 양식 날짜 필드 검증
//   - Date string format enforcement / 날짜 문자열 형식 적용
//   - API date parameter validation / API 날짜 매개변수 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single time.Parse call
//     단일 time.Parse 호출
//
// Supported formats / 지원 형식:
//   - "2006-01-02" (ISO 8601 date)
//   - "02/01/2006" (DD/MM/YYYY)
//   - "01/02/2006" (MM/DD/YYYY)
//   - "2006/01/02" (YYYY/MM/DD)
//   - Any valid Go time.Parse format
//
// Example / 예제:
//
//	// ISO 8601 format / ISO 8601 형식
//	v := validation.New("2025-10-17", "birth_date")
//	v.DateFormat("2006-01-02")  // Passes
//
//	// US format / 미국 형식
//	v = validation.New("10/17/2025", "date")
//	v.DateFormat("01/02/2006")  // Passes
//
//	// European format / 유럽 형식
//	v = validation.New("17/10/2025", "date")
//	v.DateFormat("02/01/2006")  // Passes
//
//	// Invalid format / 무효한 형식
//	v = validation.New("2025-10-17", "date")
//	v.DateFormat("01/02/2006")  // Fails (format mismatch)
//
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

// TimeFormat validates that the value is a valid time string in the specified format.
// Uses Go's time.Parse with the provided layout to verify time string validity.
//
// TimeFormat은 값이 지정된 형식의 유효한 시간 문자열인지 검증합니다.
// 제공된 레이아웃으로 Go의 time.Parse를 사용하여 시간 문자열 유효성을 확인합니다.
//
// Parameters / 매개변수:
//   - format: Go time layout format string
//     Go 시간 레이아웃 형식 문자열
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses time.Parse for validation
//     time.Parse를 사용하여 검증
//   - Accepts any Go time layout format
//     모든 Go 시간 레이아웃 형식 허용
//   - Common formats:
//     일반 형식:
//   - "15:04:05" - 24-hour with seconds (HH:MM:SS)
//   - "15:04" - 24-hour without seconds (HH:MM)
//   - "03:04:05 PM" - 12-hour with seconds
//   - "03:04 PM" - 12-hour without seconds
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//   - Fails if parsing fails
//     파싱 실패 시 실패
//
// Use Cases / 사용 사례:
//   - Time input validation / 시간 입력 검증
//   - Form time field validation / 양식 시간 필드 검증
//   - Schedule time validation / 일정 시간 검증
//   - API time parameter validation / API 시간 매개변수 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single time.Parse call
//     단일 time.Parse 호출
//
// Supported formats / 지원 형식:
//   - "15:04:05" (HH:MM:SS)
//   - "15:04" (HH:MM)
//   - "03:04:05 PM" (12-hour with seconds)
//   - "03:04 PM" (12-hour)
//   - Any valid Go time.Parse format
//
// Example / 예제:
//
//	// 24-hour format / 24시간 형식
//	v := validation.New("14:30:00", "meeting_time")
//	v.TimeFormat("15:04:05")  // Passes
//
//	// 24-hour without seconds / 초 없는 24시간 형식
//	v = validation.New("14:30", "time")
//	v.TimeFormat("15:04")  // Passes
//
//	// 12-hour format / 12시간 형식
//	v = validation.New("02:30 PM", "appointment")
//	v.TimeFormat("03:04 PM")  // Passes
//
//	// Invalid format / 무효한 형식
//	v = validation.New("14:30:00", "time")
//	v.TimeFormat("03:04 PM")  // Fails (format mismatch)
//
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

// DateBefore validates that the date value is before the specified date.
// Accepts time.Time objects or date strings in RFC3339/ISO 8601 formats.
//
// DateBefore는 날짜 값이 지정된 날짜 이전인지 검증합니다.
// time.Time 객체 또는 RFC3339/ISO 8601 형식의 날짜 문자열을 허용합니다.
//
// Parameters / 매개변수:
//   - before: Maximum date (exclusive)
//     최대 날짜 (제외)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts time.Time objects directly
//     time.Time 객체 직접 허용
//   - Accepts string in RFC3339 format ("2006-01-02T15:04:05Z07:00")
//     RFC3339 형식 문자열 허용 ("2006-01-02T15:04:05Z07:00")
//   - Accepts string in ISO 8601 format ("2006-01-02")
//     ISO 8601 형식 문자열 허용 ("2006-01-02")
//   - Tries RFC3339 first, then ISO 8601
//     RFC3339 먼저 시도, 그 다음 ISO 8601
//   - Exclusive comparison (equal dates fail)
//     배타적 비교 (같은 날짜는 실패)
//   - Fails if value is not supported type
//     값이 지원되지 않는 타입이면 실패
//   - Fails if string parsing fails
//     문자열 파싱 실패 시 실패
//
// Use Cases / 사용 사례:
//   - Expiry date validation / 만료 날짜 검증
//   - Maximum date constraint / 최대 날짜 제약
//   - Past date validation / 과거 날짜 검증
//   - Date range upper bound / 날짜 범위 상한
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) for string parsing, O(1) for time.Time
//     시간 복잡도: 문자열 파싱 O(n), time.Time O(1)
//   - Up to 2 parse attempts for strings
//     문자열에 대해 최대 2번의 파싱 시도
//
// The value can be / 값은 다음이 될 수 있음:
//   - time.Time object
//   - string in RFC3339 format ("2006-01-02T15:04:05Z07:00")
//   - string in ISO 8601 format ("2006-01-02")
//
// Example / 예제:
//
//	// time.Time validation / time.Time 검증
//	maxDate := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)
//	userDate := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)
//	v := validation.New(userDate, "expiry_date")
//	v.DateBefore(maxDate)  // Passes
//
//	// ISO 8601 string / ISO 8601 문자열
//	v = validation.New("2025-06-15", "date")
//	v.DateBefore(maxDate)  // Passes
//
//	// RFC3339 string / RFC3339 문자열
//	v = validation.New("2025-06-15T10:30:00Z", "datetime")
//	v.DateBefore(maxDate)  // Passes
//
//	// Equal date fails / 같은 날짜는 실패
//	v = validation.New(maxDate, "date")
//	v.DateBefore(maxDate)  // Fails (not before)
//
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

// DateAfter validates that the date value is after the specified date.
// Accepts time.Time objects or date strings in RFC3339/ISO 8601 formats.
//
// DateAfter는 날짜 값이 지정된 날짜 이후인지 검증합니다.
// time.Time 객체 또는 RFC3339/ISO 8601 형식의 날짜 문자열을 허용합니다.
//
// Parameters / 매개변수:
//   - after: Minimum date (exclusive)
//     최소 날짜 (제외)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts time.Time objects directly
//     time.Time 객체 직접 허용
//   - Accepts string in RFC3339 format ("2006-01-02T15:04:05Z07:00")
//     RFC3339 형식 문자열 허용 ("2006-01-02T15:04:05Z07:00")
//   - Accepts string in ISO 8601 format ("2006-01-02")
//     ISO 8601 형식 문자열 허용 ("2006-01-02")
//   - Tries RFC3339 first, then ISO 8601
//     RFC3339 먼저 시도, 그 다음 ISO 8601
//   - Exclusive comparison (equal dates fail)
//     배타적 비교 (같은 날짜는 실패)
//   - Fails if value is not supported type
//     값이 지원되지 않는 타입이면 실패
//   - Fails if string parsing fails
//     문자열 파싱 실패 시 실패
//
// Use Cases / 사용 사례:
//   - Start date validation / 시작 날짜 검증
//   - Minimum date constraint / 최소 날짜 제약
//   - Future date validation / 미래 날짜 검증
//   - Date range lower bound / 날짜 범위 하한
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) for string parsing, O(1) for time.Time
//     시간 복잡도: 문자열 파싱 O(n), time.Time O(1)
//   - Up to 2 parse attempts for strings
//     문자열에 대해 최대 2번의 파싱 시도
//
// The value can be / 값은 다음이 될 수 있음:
//   - time.Time object
//   - string in RFC3339 format ("2006-01-02T15:04:05Z07:00")
//   - string in ISO 8601 format ("2006-01-02")
//
// Example / 예제:
//
//	// time.Time validation / time.Time 검증
//	minDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	userDate := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)
//	v := validation.New(userDate, "start_date")
//	v.DateAfter(minDate)  // Passes
//
//	// ISO 8601 string / ISO 8601 문자열
//	v = validation.New("2025-06-15", "date")
//	v.DateAfter(minDate)  // Passes
//
//	// RFC3339 string / RFC3339 문자열
//	v = validation.New("2025-06-15T10:30:00Z", "datetime")
//	v.DateAfter(minDate)  // Passes
//
//	// Equal date fails / 같은 날짜는 실패
//	v = validation.New(minDate, "date")
//	v.DateAfter(minDate)  // Fails (not after)
//
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
