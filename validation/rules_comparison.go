package validation

import (
	"fmt"
	"time"
)

// Equals validates that the value is exactly equal to the given value.
// Uses Go's == operator for comparison, supporting all comparable types.
//
// Equals는 값이 주어진 값과 정확히 같은지 검증합니다.
// 비교 가능한 모든 타입을 지원하는 Go의 == 연산자를 사용합니다.
//
// Parameters / 매개변수:
//   - value: The expected value to compare against
//     비교할 기대값
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses == operator for comparison
//     비교에 == 연산자 사용
//   - Type-sensitive comparison (5 != "5")
//     타입 구분 비교 (5 != "5")
//   - Case-sensitive for strings
//     문자열은 대소문자 구분
//   - Exact match required
//     정확한 일치 필요
//
// Use Cases / 사용 사례:
//   - Password confirmation / 비밀번호 확인
//   - Exact value validation / 정확한 값 검증
//   - Field matching validation / 필드 일치 검증
//   - Expected value checks / 예상 값 확인
//   - Agreement/consent validation / 동의/승인 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1) for primitive types
//     시간 복잡도: 기본 타입의 경우 O(1)
//   - Direct comparison, very fast
//     직접 비교, 매우 빠름
//
// Limitations / 제한사항:
//   - Not suitable for complex type comparison (use reflect.DeepEqual instead)
//     복잡한 타입 비교에는 부적합 (대신 reflect.DeepEqual 사용)
//   - Slice/map/struct comparison may not work as expected
//     슬라이스/맵/구조체 비교는 예상대로 작동하지 않을 수 있음
//
// Example / 예제:
//   // Password confirmation / 비밀번호 확인
//   v := validation.New("password123", "confirm_password")
//   v.Equals(originalPassword)  // Must match exactly
//
//   // Exact string match / 정확한 문자열 일치
//   v := validation.New("admin", "role")
//   v.Equals("admin")  // Passes / 성공
//
//   // Numeric equality / 숫자 동등성
//   v := validation.New(42, "answer")
//   v.Equals(42)  // Passes / 성공
//
//   // Case-sensitive / 대소문자 구분
//   v := validation.New("Admin", "role")
//   v.Equals("admin")  // Fails / 실패
//
//   // Type-sensitive / 타입 구분
//   v := validation.New(5, "number")
//   v.Equals("5")  // Fails (different types) / 실패 (다른 타입)
func (v *Validator) Equals(value interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if v.value != value {
		v.addError("equals", fmt.Sprintf("%s must equal '%v' / %s은(는) '%v'와 같아야 합니다", v.fieldName, value, v.fieldName, value))
	}

	return v
}

// NotEquals validates that the value is NOT equal to the given value.
// Uses Go's != operator for comparison, supporting all comparable types.
//
// NotEquals는 값이 주어진 값과 같지 않은지 검증합니다.
// 비교 가능한 모든 타입을 지원하는 Go의 != 연산자를 사용합니다.
//
// Parameters / 매개변수:
//   - value: The forbidden value to compare against
//     비교할 금지 값
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses != operator for comparison
//     비교에 != 연산자 사용
//   - Type-sensitive comparison (5 != "5")
//     타입 구분 비교 (5 != "5")
//   - Case-sensitive for strings
//     문자열은 대소문자 구분
//   - Fails only on exact match
//     정확히 일치할 때만 실패
//
// Use Cases / 사용 사례:
//   - Prevent reusing old password / 이전 비밀번호 재사용 방지
//   - Forbidden value validation / 금지된 값 검증
//   - Ensure value has changed / 값이 변경되었는지 확인
//   - Default value prevention / 기본값 방지
//   - Duplicate prevention / 중복 방지
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1) for primitive types
//     시간 복잡도: 기본 타입의 경우 O(1)
//   - Direct comparison, very fast
//     직접 비교, 매우 빠름
//
// Example / 예제:
//   // Prevent old password / 이전 비밀번호 방지
//   v := validation.New("newpassword123", "new_password")
//   v.NotEquals(oldPassword)  // Must be different
//
//   // Ensure value changed / 값 변경 확인
//   v := validation.New(newEmail, "email")
//   v.NotEquals(currentEmail)
//
//   // Forbidden value / 금지된 값
//   v := validation.New("guest", "username")
//   v.NotEquals("admin")  // Cannot be admin
//
//   // Default value prevention / 기본값 방지
//   v := validation.New(status, "status")
//   v.NotEquals("unknown")  // Must set a valid status
//
//   // Multiple forbidden values / 여러 금지 값
//   v := validation.New(role, "role")
//   v.NotEquals("root").NotEquals("system")
func (v *Validator) NotEquals(value interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if v.value == value {
		v.addError("notequals", fmt.Sprintf("%s must not equal '%v' / %s은(는) '%v'와 달라야 합니다", v.fieldName, value, v.fieldName, value))
	}

	return v
}

// GreaterThan validates that the numeric value is strictly greater than the given value.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// GreaterThan은 숫자 값이 주어진 값보다 엄격하게 큰지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Parameters / 매개변수:
//   - value: The threshold value (exclusive)
//     임계값 (미포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val > threshold
//     val > 임계값이면 허용
//   - Exclusive comparison (equals fails)
//     미포함 비교 (같으면 실패)
//   - Converts to float64 for comparison
//     비교를 위해 float64로 변환
//   - Fails if value is not numeric
//     값이 숫자가 아니면 실패
//
// Use Cases / 사용 사례:
//   - Minimum threshold validation / 최소 임계값 검증
//   - Score requirements / 점수 요구사항
//   - Range validation (with LessThan)
//     범위 검증 (LessThan과 함께)
//   - Positive number validation / 양수 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Simple numeric comparison
//     단순 숫자 비교
//
// Comparison with Min() / Min()과의 비교:
//   - GreaterThan(5): value must be > 5 (6, 7, 8...)
//   - Min(5): value must be >= 5 (5, 6, 7...)
//   - Use GreaterThan for exclusive bounds
//     미포함 경계에는 GreaterThan 사용
//
// Example / 예제:
//   // Exclusive threshold / 미포함 임계값
//   v := validation.New(10, "score")
//   v.GreaterThan(5)  // Passes (10 > 5) / 성공
//
//   v := validation.New(5, "score")
//   v.GreaterThan(5)  // Fails (5 = 5, not >) / 실패
//
//   // Positive validation / 양수 검증
//   v := validation.New(amount, "amount")
//   v.GreaterThan(0)  // Must be positive
//
//   // Range validation / 범위 검증
//   v := validation.New(value, "value")
//   v.GreaterThan(0).LessThan(100)  // 0 < value < 100
func (v *Validator) GreaterThan(value float64) *Validator {
	return validateNumeric(v, "greaterthan", func(n float64) bool {
		return n > value
	}, fmt.Sprintf("%s must be greater than %v / %s은(는) %v보다 커야 합니다", v.fieldName, value, v.fieldName, value))
}

// GreaterThanOrEqual validates that the numeric value is greater than or equal to the given value.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// GreaterThanOrEqual은 숫자 값이 주어진 값보다 크거나 같은지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Parameters / 매개변수:
//   - value: The minimum threshold value (inclusive)
//     최소 임계값 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val >= threshold
//     val >= 임계값이면 허용
//   - Inclusive comparison (equals passes)
//     포함 비교 (같으면 통과)
//   - Converts to float64 for comparison
//     비교를 위해 float64로 변환
//   - Fails if value is not numeric
//     값이 숫자가 아니면 실패
//
// Use Cases / 사용 사례:
//   - Minimum value validation / 최소값 검증
//   - Non-negative validation / 음수가 아닌 값 검증
//   - Threshold requirements / 임계값 요구사항
//   - Quota validation / 할당량 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Simple numeric comparison
//     단순 숫자 비교
//
// Comparison with Min() / Min()과의 비교:
//   - GreaterThanOrEqual(5): Same as Min(5)
//   - Both accept value >= 5
//   - Min() is more idiomatic for minimum bounds
//     최소 경계에는 Min()이 더 관용적
//
// Example / 예제:
//   // Inclusive minimum / 포함 최소값
//   v := validation.New(10, "score")
//   v.GreaterThanOrEqual(10)  // Passes (10 >= 10) / 성공
//
//   v := validation.New(9, "score")
//   v.GreaterThanOrEqual(10)  // Fails (9 < 10) / 실패
//
//   // Non-negative validation / 음수가 아닌 값 검증
//   v := validation.New(count, "count")
//   v.GreaterThanOrEqual(0)  // Can be 0 or positive
//
//   // Minimum age / 최소 나이
//   v := validation.New(age, "age")
//   v.GreaterThanOrEqual(18)  // 18 or older
func (v *Validator) GreaterThanOrEqual(value float64) *Validator {
	return validateNumeric(v, "greaterthanorequal", func(n float64) bool {
		return n >= value
	}, fmt.Sprintf("%s must be greater than or equal to %v / %s은(는) %v보다 크거나 같아야 합니다", v.fieldName, value, v.fieldName, value))
}

// LessThan validates that the numeric value is strictly less than the given value.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// LessThan은 숫자 값이 주어진 값보다 엄격하게 작은지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Parameters / 매개변수:
//   - value: The threshold value (exclusive)
//     임계값 (미포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val < threshold
//     val < 임계값이면 허용
//   - Exclusive comparison (equals fails)
//     미포함 비교 (같으면 실패)
//   - Converts to float64 for comparison
//     비교를 위해 float64로 변환
//   - Fails if value is not numeric
//     값이 숫자가 아니면 실패
//
// Use Cases / 사용 사례:
//   - Maximum threshold validation / 최대 임계값 검증
//   - Upper bound validation / 상한 검증
//   - Range validation (with GreaterThan)
//     범위 검증 (GreaterThan과 함께)
//   - Negative number validation / 음수 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Simple numeric comparison
//     단순 숫자 비교
//
// Comparison with Max() / Max()와의 비교:
//   - LessThan(10): value must be < 10 (9, 8, 7...)
//   - Max(10): value must be <= 10 (10, 9, 8...)
//   - Use LessThan for exclusive upper bounds
//     미포함 상한에는 LessThan 사용
//
// Example / 예제:
//   // Exclusive upper bound / 미포함 상한
//   v := validation.New(5, "score")
//   v.LessThan(10)  // Passes (5 < 10) / 성공
//
//   v := validation.New(10, "score")
//   v.LessThan(10)  // Fails (10 = 10, not <) / 실패
//
//   // Negative validation / 음수 검증
//   v := validation.New(value, "value")
//   v.LessThan(0)  // Must be negative
//
//   // Open range validation / 열린 범위 검증
//   v := validation.New(value, "value")
//   v.GreaterThan(0).LessThan(100)  // 0 < value < 100
func (v *Validator) LessThan(value float64) *Validator {
	return validateNumeric(v, "lessthan", func(n float64) bool {
		return n < value
	}, fmt.Sprintf("%s must be less than %v / %s은(는) %v보다 작아야 합니다", v.fieldName, value, v.fieldName, value))
}

// LessThanOrEqual validates that the numeric value is less than or equal to the given value.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// LessThanOrEqual은 숫자 값이 주어진 값보다 작거나 같은지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Parameters / 매개변수:
//   - value: The maximum threshold value (inclusive)
//     최대 임계값 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val <= threshold
//     val <= 임계값이면 허용
//   - Inclusive comparison (equals passes)
//     포함 비교 (같으면 통과)
//   - Converts to float64 for comparison
//     비교를 위해 float64로 변환
//   - Fails if value is not numeric
//     값이 숫자가 아니면 실패
//
// Use Cases / 사용 사례:
//   - Maximum value validation / 최대값 검증
//   - Upper limit validation / 상한 검증
//   - Capacity constraints / 용량 제약
//   - Budget limits / 예산 제한
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Simple numeric comparison
//     단순 숫자 비교
//
// Comparison with Max() / Max()와의 비교:
//   - LessThanOrEqual(10): Same as Max(10)
//   - Both accept value <= 10
//   - Max() is more idiomatic for maximum bounds
//     최대 경계에는 Max()가 더 관용적
//
// Example / 예제:
//   // Inclusive maximum / 포함 최대값
//   v := validation.New(10, "score")
//   v.LessThanOrEqual(10)  // Passes (10 <= 10) / 성공
//
//   v := validation.New(11, "score")
//   v.LessThanOrEqual(10)  // Fails (11 > 10) / 실패
//
//   // Maximum age / 최대 나이
//   v := validation.New(age, "age")
//   v.LessThanOrEqual(120)  // 120 or younger
//
//   // Closed range / 닫힌 범위
//   v := validation.New(value, "value")
//   v.GreaterThanOrEqual(0).LessThanOrEqual(100)  // 0 <= value <= 100
func (v *Validator) LessThanOrEqual(value float64) *Validator {
	return validateNumeric(v, "lessthanorequal", func(n float64) bool {
		return n <= value
	}, fmt.Sprintf("%s must be less than or equal to %v / %s은(는) %v보다 작거나 같아야 합니다", v.fieldName, value, v.fieldName, value))
}

// Before validates that the time value is strictly before the given time.
// Only works with time.Time values. Uses time.Before() for comparison.
//
// Before는 시간 값이 주어진 시간보다 엄격하게 이전인지 검증합니다.
// time.Time 값에서만 작동합니다. 비교에 time.Before()를 사용합니다.
//
// Parameters / 매개변수:
//   - t: The threshold time (exclusive)
//     임계 시간 (미포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses time.Before() for comparison
//     비교에 time.Before() 사용
//   - Exclusive comparison (exact time fails)
//     미포함 비교 (정확히 같으면 실패)
//   - Timezone-aware comparison
//     시간대 인식 비교
//   - Fails if value is not time.Time
//     값이 time.Time이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Event deadline validation / 이벤트 마감 검증
//   - Historical date validation / 과거 날짜 검증
//   - Date range validation / 날짜 범위 검증
//   - Expiration checks / 만료 확인
//   - Start/end date validation / 시작/종료 날짜 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Direct time comparison
//     직접 시간 비교
//
// Timezone Considerations / 시간대 고려사항:
//   - Compares absolute time, not wall clock time
//     벽시계 시간이 아닌 절대 시간 비교
//   - Timezone differences handled correctly
//     시간대 차이 올바르게 처리
//   - Use UTC for consistent comparisons
//     일관된 비교를 위해 UTC 사용
//
// Example / 예제:
//   // Deadline validation / 마감 검증
//   deadline := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
//   v := validation.New(submittedTime, "submitted_at")
//   v.Before(deadline)  // Must be submitted before deadline
//
//   // Past date validation / 과거 날짜 검증
//   now := time.Now()
//   v := validation.New(birthDate, "birth_date")
//   v.Before(now)  // Birth date must be in the past
//
//   // Event scheduling / 이벤트 스케줄링
//   v := validation.New(startTime, "start_time")
//   v.Before(endTime)  // Start must be before end
//
//   // Exact time fails / 정확한 시간 실패
//   now := time.Now()
//   v := validation.New(now, "time")
//   v.Before(now)  // Fails (not strictly before)
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

// After validates that the time value is strictly after the given time.
// Only works with time.Time values. Uses time.After() for comparison.
//
// After는 시간 값이 주어진 시간보다 엄격하게 이후인지 검증합니다.
// time.Time 값에서만 작동합니다. 비교에 time.After()를 사용합니다.
//
// Parameters / 매개변수:
//   - t: The threshold time (exclusive)
//     임계 시간 (미포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses time.After() for comparison
//     비교에 time.After() 사용
//   - Exclusive comparison (exact time fails)
//     미포함 비교 (정확히 같으면 실패)
//   - Timezone-aware comparison
//     시간대 인식 비교
//   - Fails if value is not time.Time
//     값이 time.Time이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Future date validation / 미래 날짜 검증
//   - Event start time validation / 이벤트 시작 시간 검증
//   - Minimum date requirements / 최소 날짜 요구사항
//   - After-hours validation / 업무 시간 이후 검증
//   - Chronological order validation / 시간순 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Direct time comparison
//     직접 시간 비교
//
// Timezone Considerations / 시간대 고려사항:
//   - Compares absolute time, not wall clock time
//     벽시계 시간이 아닌 절대 시간 비교
//   - Timezone differences handled correctly
//     시간대 차이 올바르게 처리
//   - Use UTC for consistent comparisons
//     일관된 비교를 위해 UTC 사용
//
// Example / 예제:
//   // Future date validation / 미래 날짜 검증
//   now := time.Now()
//   v := validation.New(eventDate, "event_date")
//   v.After(now)  // Event must be in the future
//
//   // Start after reference / 참조 이후 시작
//   referenceTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//   v := validation.New(startTime, "start_time")
//   v.After(referenceTime)
//
//   // Event sequence / 이벤트 순서
//   v := validation.New(endTime, "end_time")
//   v.After(startTime)  // End must be after start
//
//   // Exact time fails / 정확한 시간 실패
//   now := time.Now()
//   v := validation.New(now, "time")
//   v.After(now)  // Fails (not strictly after)
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

// BeforeOrEqual validates that the time value is before or equal to the given time.
// Only works with time.Time values. Inclusive comparison allowing exact time match.
//
// BeforeOrEqual은 시간 값이 주어진 시간보다 이전이거나 같은지 검증합니다.
// time.Time 값에서만 작동합니다. 정확한 시간 일치를 허용하는 포함 비교입니다.
//
// Parameters / 매개변수:
//   - t: The threshold time (inclusive)
//     임계 시간 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts if time is before OR equal to threshold
//     시간이 임계값 이전이거나 같으면 허용
//   - Inclusive comparison (exact time passes)
//     포함 비교 (정확히 같으면 통과)
//   - Uses !timeVal.After(t) for implementation
//     구현에 !timeVal.After(t) 사용
//   - Timezone-aware comparison
//     시간대 인식 비교
//   - Fails if value is not time.Time
//     값이 time.Time이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Deadline validation (inclusive) / 마감 검증 (포함)
//   - Maximum date validation / 최대 날짜 검증
//   - End date validation / 종료 날짜 검증
//   - Historical or current date validation / 과거 또는 현재 날짜 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Direct time comparison
//     직접 시간 비교
//
// Example / 예제:
//   // Deadline validation (inclusive) / 마감 검증 (포함)
//   deadline := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
//   v := validation.New(submittedTime, "submitted_at")
//   v.BeforeOrEqual(deadline)  // Can submit up to and including deadline
//
//   // Current or past date / 현재 또는 과거 날짜
//   now := time.Now()
//   v := validation.New(date, "date")
//   v.BeforeOrEqual(now)  // Can be today or earlier
//
//   // Exact time passes / 정확한 시간 통과
//   deadline := time.Now()
//   v := validation.New(deadline, "time")
//   v.BeforeOrEqual(deadline)  // Passes (equal allowed)
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

// AfterOrEqual validates that the time value is after or equal to the given time.
// Only works with time.Time values. Inclusive comparison allowing exact time match.
//
// AfterOrEqual은 시간 값이 주어진 시간보다 이후이거나 같은지 검증합니다.
// time.Time 값에서만 작동합니다. 정확한 시간 일치를 허용하는 포함 비교입니다.
//
// Parameters / 매개변수:
//   - t: The threshold time (inclusive)
//     임계 시간 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts if time is after OR equal to threshold
//     시간이 임계값 이후이거나 같으면 허용
//   - Inclusive comparison (exact time passes)
//     포함 비교 (정확히 같으면 통과)
//   - Uses !timeVal.Before(t) for implementation
//     구현에 !timeVal.Before(t) 사용
//   - Timezone-aware comparison
//     시간대 인식 비교
//   - Fails if value is not time.Time
//     값이 time.Time이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Minimum date validation / 최소 날짜 검증
//   - Start date validation / 시작 날짜 검증
//   - Future or current date validation / 미래 또는 현재 날짜 검증
//   - Availability start time / 가용성 시작 시간
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Direct time comparison
//     직접 시간 비교
//
// Example / 예제:
//   // Minimum start date / 최소 시작 날짜
//   minDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//   v := validation.New(startDate, "start_date")
//   v.AfterOrEqual(minDate)  // Can start on or after 2024-01-01
//
//   // Current or future date / 현재 또는 미래 날짜
//   now := time.Now()
//   v := validation.New(eventDate, "event_date")
//   v.AfterOrEqual(now)  // Can be today or later
//
//   // Exact time passes / 정확한 시간 통과
//   startTime := time.Now()
//   v := validation.New(startTime, "time")
//   v.AfterOrEqual(startTime)  // Passes (equal allowed)
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

// BetweenTime validates that the time value is within the given time range (inclusive).
// Only works with time.Time values. Both start and end boundaries are included.
//
// BetweenTime은 시간 값이 주어진 시간 범위 내에 있는지 검증합니다 (포함).
// time.Time 값에서만 작동합니다. 시작 및 종료 경계값 모두 포함됩니다.
//
// Parameters / 매개변수:
//   - start: The start time of the range (inclusive)
//     범위의 시작 시간 (포함)
//   - end: The end time of the range (inclusive)
//     범위의 종료 시간 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts if start <= time <= end
//     시작 <= 시간 <= 종료이면 허용
//   - Inclusive on both boundaries
//     양쪽 경계 모두 포함
//   - Timezone-aware comparison
//     시간대 인식 비교
//   - Fails if value is not time.Time
//     값이 time.Time이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Date range validation / 날짜 범위 검증
//   - Event time slot validation / 이벤트 시간대 검증
//   - Business hours validation / 업무 시간 검증
//   - Seasonal validation / 계절 검증
//   - Booking period validation / 예약 기간 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Two time comparisons
//     두 번의 시간 비교
//
// Example / 예제:
//   // Year 2024 validation / 2024년 검증
//   start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//   end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
//   v := validation.New(date, "date")
//   v.BetweenTime(start, end)  // Must be in 2024
//
//   // Business hours (9 AM - 5 PM) / 업무 시간 (오전 9시 - 오후 5시)
//   start := time.Date(2024, 1, 1, 9, 0, 0, 0, time.Local)
//   end := time.Date(2024, 1, 1, 17, 0, 0, 0, time.Local)
//   v := validation.New(appointmentTime, "appointment")
//   v.BetweenTime(start, end)
//
//   // Boundaries included / 경계값 포함
//   start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//   end := time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)
//   v := validation.New(start, "date")
//   v.BetweenTime(start, end)  // Passes (start included)
//
//   v := validation.New(end, "date")
//   v.BetweenTime(start, end)  // Passes (end included)
func (v *Validator) BetweenTime(start, end time.Time) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	timeVal, ok := v.value.(time.Time)
	if !ok {
		v.addError("betweentime", fmt.Sprintf("%s must be a time.Time value / %s은(는) time.Time 값이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if timeVal.Before(start) || timeVal.After(end) {
		v.addError("betweentime", fmt.Sprintf("%s must be between %s and %s / %s은(는) %s와(과) %s 사이여야 합니다",
			v.fieldName, start.Format(time.RFC3339), end.Format(time.RFC3339), v.fieldName, start.Format(time.RFC3339), end.Format(time.RFC3339)))
	}

	return v
}
