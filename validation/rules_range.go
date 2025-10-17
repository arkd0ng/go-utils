package validation

import (
	"fmt"
	"time"
)

// IntRange validates that an integer value is within the specified range (inclusive).
// Accepts all integer types (int, int8-64, uint, uint8-64) with automatic conversion.
//
// IntRange는 정수 값이 지정된 범위 내에 있는지 검증합니다 (포함).
// 자동 변환을 통해 모든 정수 타입(int, int8-64, uint, uint8-64)을 허용합니다.
//
// Parameters / 매개변수:
//   - min: Minimum value (inclusive)
//     최소값 (포함)
//   - max: Maximum value (inclusive)
//     최대값 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts all signed integer types (int, int8, int16, int32, int64)
//     모든 부호 있는 정수 타입 허용 (int, int8, int16, int32, int64)
//   - Accepts all unsigned integer types (uint, uint8, uint16, uint32, uint64)
//     모든 부호 없는 정수 타입 허용 (uint, uint8, uint16, uint32, uint64)
//   - Inclusive range: min <= value <= max
//     포함 범위: min <= 값 <= max
//   - Automatically converts to int64 for comparison
//     비교를 위해 int64로 자동 변환
//   - Fails if value is not an integer type
//     값이 정수 타입이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Age range validation / 나이 범위 검증
//   - Quantity limits / 수량 제한
//   - Score validation / 점수 검증
//   - Rating systems / 평점 시스템
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Type switch + comparison
//     타입 스위치 + 비교
//
// The range is inclusive: min <= value <= max
//
// Example / 예제:
//
//	// Age validation / 나이 검증
//	v := validation.New(25, "age")
//	v.IntRange(18, 65)  // Passes (working age)
//
//	// Score validation / 점수 검증
//	v = validation.New(85, "score")
//	v.IntRange(0, 100)  // Passes
//
//	// Boundaries included / 경계값 포함
//	v = validation.New(18, "age")
//	v.IntRange(18, 65)  // Passes (min included)
//
//	v = validation.New(65, "age")
//	v.IntRange(18, 65)  // Passes (max included)
//
//	// Out of range / 범위 밖
//	v = validation.New(150, "age")
//	v.IntRange(0, 120)  // Fails
//
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

// FloatRange validates that a floating-point value is within the specified range (inclusive).
// Accepts float32, float64, and all integer types with automatic conversion.
//
// FloatRange는 부동소수점 값이 지정된 범위 내에 있는지 검증합니다 (포함).
// float32, float64 및 모든 정수 타입을 자동 변환과 함께 허용합니다.
//
// Parameters / 매개변수:
//   - min: Minimum value (inclusive)
//     최소값 (포함)
//   - max: Maximum value (inclusive)
//     최대값 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts float32 and float64 types
//     float32 및 float64 타입 허용
//   - Accepts all integer types (converted to float64)
//     모든 정수 타입 허용 (float64로 변환)
//   - Inclusive range: min <= value <= max
//     포함 범위: min <= 값 <= max
//   - Automatically converts to float64 for comparison
//     비교를 위해 float64로 자동 변환
//   - Fails if value is not numeric
//     값이 숫자가 아니면 실패
//
// Use Cases / 사용 사례:
//   - Temperature range validation / 온도 범위 검증
//   - Percentage validation / 백분율 검증
//   - Price range validation / 가격 범위 검증
//   - Measurement validation / 측정값 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Type switch + comparison
//     타입 스위치 + 비교
//
// The range is inclusive: min <= value <= max
//
// Example / 예제:
//
//	// Temperature validation / 온도 검증
//	v := validation.New(98.6, "temperature")
//	v.FloatRange(95.0, 105.0)  // Passes (normal body temp)
//
//	// Percentage validation / 백분율 검증
//	v = validation.New(85.5, "grade")
//	v.FloatRange(0.0, 100.0)  // Passes
//
//	// Boundaries included / 경계값 포함
//	v = validation.New(0.0, "value")
//	v.FloatRange(0.0, 10.0)  // Passes (min included)
//
//	v = validation.New(10.0, "value")
//	v.FloatRange(0.0, 10.0)  // Passes (max included)
//
//	// Integer accepted / 정수 허용
//	v = validation.New(98, "temp")
//	v.FloatRange(95.0, 105.0)  // Passes (converted to float)
//
//	// Out of range / 범위 밖
//	v = validation.New(110.5, "temperature")
//	v.FloatRange(95.0, 105.0)  // Fails
//
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

// DateRange validates that a date value is within the specified date range (inclusive).
// Accepts time.Time objects or date strings in RFC3339/ISO 8601 formats.
//
// DateRange는 날짜 값이 지정된 날짜 범위 내에 있는지 검증합니다 (포함).
// time.Time 객체 또는 RFC3339/ISO 8601 형식의 날짜 문자열을 허용합니다.
//
// Parameters / 매개변수:
//   - start: Start date (inclusive)
//     시작 날짜 (포함)
//   - end: End date (inclusive)
//     종료 날짜 (포함)
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
//   - Inclusive range: start <= value <= end
//     포함 범위: 시작 <= 값 <= 종료
//   - Fails if value is not supported type
//     값이 지원되지 않는 타입이면 실패
//   - Fails if string parsing fails
//     문자열 파싱 실패 시 실패
//
// Use Cases / 사용 사례:
//   - Event date validation / 이벤트 날짜 검증
//   - Booking period validation / 예약 기간 검증
//   - Valid date window / 유효 날짜 창
//   - Historical period validation / 역사적 기간 검증
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
// The range is inclusive: start <= value <= end
//
// The value can be / 값은 다음이 될 수 있음:
//   - time.Time object
//   - string in RFC3339 format ("2006-01-02T15:04:05Z07:00")
//   - string in ISO 8601 format ("2006-01-02")
//
// Example / 예제:
//
//	// Event date range / 이벤트 날짜 범위
//	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
//	eventDate := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)
//	v := validation.New(eventDate, "event_date")
//	v.DateRange(start, end)  // Passes (within 2025)
//
//	// ISO 8601 string / ISO 8601 문자열
//	v = validation.New("2025-06-15", "date")
//	v.DateRange(start, end)  // Passes
//
//	// RFC3339 string / RFC3339 문자열
//	v = validation.New("2025-06-15T10:30:00Z", "datetime")
//	v.DateRange(start, end)  // Passes
//
//	// Boundaries included / 경계값 포함
//	v = validation.New(start, "date")
//	v.DateRange(start, end)  // Passes (start included)
//
//	v = validation.New(end, "date")
//	v.DateRange(start, end)  // Passes (end included)
//
//	// Out of range / 범위 밖
//	v = validation.New("2024-12-31", "date")
//	v.DateRange(start, end)  // Fails (before start)
//
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

// toInt64 converts various integer types to int64 for unified comparison.
// Internal helper function supporting all signed and unsigned integer types.
//
// toInt64는 통합 비교를 위해 다양한 정수 타입을 int64로 변환합니다.
// 모든 부호 있는/없는 정수 타입을 지원하는 내부 헬퍼 함수입니다.
//
// Parameters / 매개변수:
//   - val: Value to convert (any integer type)
//     변환할 값 (모든 정수 타입)
//
// Returns / 반환:
//   - int64: Converted integer value
//     int64: 변환된 정수 값
//   - error: Error if value is not an integer type
//     error: 값이 정수 타입이 아닌 경우 오류
//
// Supported Types / 지원 타입:
//   - Signed: int, int8, int16, int32, int64
//     부호 있는: int, int8, int16, int32, int64
//   - Unsigned: uint, uint8, uint16, uint32, uint64
//     부호 없는: uint, uint8, uint16, uint32, uint64
//
// Example / 예제:
//
//	val, err := v.toInt64(42)       // int -> int64
//	val, err := v.toInt64(int8(42))  // int8 -> int64
//	val, err := v.toInt64(uint(42))  // uint -> int64
//	val, err := v.toInt64(3.14)      // error (not integer)
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

// toFloat64 converts various numeric types to float64 for unified comparison.
// Internal helper function supporting float and integer types.
//
// toFloat64는 통합 비교를 위해 다양한 숫자 타입을 float64로 변환합니다.
// 부동소수점 및 정수 타입을 지원하는 내부 헬퍼 함수입니다.
//
// Parameters / 매개변수:
//   - val: Value to convert (float or integer type)
//     변환할 값 (부동소수점 또는 정수 타입)
//
// Returns / 반환:
//   - float64: Converted floating-point value
//     float64: 변환된 부동소수점 값
//   - error: Error if value is not a numeric type
//     error: 값이 숫자 타입이 아닌 경우 오류
//
// Supported Types / 지원 타입:
//   - Float: float32, float64
//     부동소수점: float32, float64
//   - Signed integers: int, int8, int16, int32, int64
//     부호 있는 정수: int, int8, int16, int32, int64
//   - Unsigned integers: uint, uint8, uint16, uint32, uint64
//     부호 없는 정수: uint, uint8, uint16, uint32, uint64
//
// Example / 예제:
//
//	val, err := v.toFloat64(3.14)        // float64 -> float64
//	val, err := v.toFloat64(float32(3.14)) // float32 -> float64
//	val, err := v.toFloat64(42)          // int -> float64
//	val, err := v.toFloat64("3.14")      // error (not numeric)
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
