package validation

import "fmt"

// Min validates that the numeric value is greater than or equal to the minimum threshold.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// Min은 숫자 값이 최소 임계값 이상인지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Parameters / 매개변수:
//   - n: Minimum value (inclusive)
//     최소값 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val >= n
//     val >= n이면 허용
//   - Converts value to float64 for comparison
//     비교를 위해 값을 float64로 변환
//   - Inclusive comparison (equals is valid)
//     포함 비교 (같음은 유효)
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Age validation (minimum age)
//     나이 검증 (최소 나이)
//   - Price validation (minimum price)
//     가격 검증 (최소 가격)
//   - Quantity validation / 수량 검증
//   - Score/rating thresholds / 점수/평점 임계값
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
// Example / 예제:
//
//	v := validation.New(25, "age")
//	v.Min(18)  // Passes / 성공
//
//	v := validation.New(15, "age")
//	v.Min(18)  // Fails / 실패
//
//	v := validation.New(18, "age")
//	v.Min(18)  // Passes (inclusive) / 성공 (포함)
//
//	// Price validation / 가격 검증
//	v := validation.New(price, "price")
//	v.Positive().Min(0.01).Max(99999.99)
//
//	// Age requirement / 나이 요구사항
//	v := validation.New(age, "age")
//	v.Required().Min(18).Max(120)
func (v *Validator) Min(n float64) *Validator {
	return validateNumeric(v, "min", func(val float64) bool {
		return val >= n
	}, fmt.Sprintf("%s must be at least %v / %s은(는) 최소 %v 이상이어야 합니다", v.fieldName, n, v.fieldName, n))
}

// Max validates that the numeric value is less than or equal to the maximum threshold.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// Max는 숫자 값이 최대 임계값 이하인지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Parameters / 매개변수:
//   - n: Maximum value (inclusive)
//     최대값 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val <= n
//     val <= n이면 허용
//   - Converts value to float64 for comparison
//     비교를 위해 값을 float64로 변환
//   - Inclusive comparison (equals is valid)
//     포함 비교 (같음은 유효)
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Age validation (maximum age)
//     나이 검증 (최대 나이)
//   - Price caps / 가격 상한
//   - Quantity limits / 수량 제한
//   - Score/rating upper bounds / 점수/평점 상한
//   - Database column constraints / 데이터베이스 컬럼 제약
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
// Example / 예제:
//
//	v := validation.New(15, "age")
//	v.Max(18)  // Passes / 성공
//
//	v := validation.New(25, "age")
//	v.Max(18)  // Fails / 실패
//
//	v := validation.New(18, "age")
//	v.Max(18)  // Passes (inclusive) / 성공 (포함)
//
//	// Percentage validation / 백분율 검증
//	v := validation.New(percentage, "discount")
//	v.Min(0).Max(100)
//
//	// Rating validation / 평점 검증
//	v := validation.New(rating, "rating")
//	v.Positive().Min(1).Max(5)
func (v *Validator) Max(n float64) *Validator {
	return validateNumeric(v, "max", func(val float64) bool {
		return val <= n
	}, fmt.Sprintf("%s must be at most %v / %s은(는) 최대 %v 이하여야 합니다", v.fieldName, n, v.fieldName, n))
}

// Between validates that the numeric value is within a specified range (inclusive).
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// Between은 숫자 값이 지정된 범위 내에 있는지 검증합니다 (포함).
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
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
//   - Accepts value if min <= val <= max
//     min <= val <= max이면 허용
//   - Converts value to float64 for comparison
//     비교를 위해 값을 float64로 변환
//   - Inclusive on both ends (boundaries are valid)
//     양 끝 포함 (경계값은 유효)
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Age range validation / 나이 범위 검증
//   - Price range filtering / 가격 범위 필터링
//   - Percentage validation (0-100) / 백분율 검증 (0-100)
//   - Rating systems (1-5, 1-10) / 평점 시스템 (1-5, 1-10)
//   - Temperature ranges / 온도 범위
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Two numeric comparisons
//     두 개의 숫자 비교
//
// Example / 예제:
//
//	v := validation.New(25, "age")
//	v.Between(18, 65)  // Passes / 성공
//
//	v := validation.New(10, "age")
//	v.Between(18, 65)  // Fails (< min) / 실패 (최소값 미만)
//
//	v := validation.New(70, "age")
//	v.Between(18, 65)  // Fails (> max) / 실패 (최대값 초과)
//
//	v := validation.New(18, "age")
//	v.Between(18, 65)  // Passes (boundary) / 성공 (경계값)
//
//	// Percentage validation / 백분율 검증
//	v := validation.New(percentage, "discount")
//	v.Between(0, 100)
//
//	// Rating validation / 평점 검증
//	v := validation.New(rating, "rating")
//	v.Between(1, 5)
//
//	// Temperature range / 온도 범위
//	v := validation.New(temp, "temperature")
//	v.Between(-20, 50)
func (v *Validator) Between(min, max float64) *Validator {
	return validateNumeric(v, "between", func(val float64) bool {
		return val >= min && val <= max
	}, fmt.Sprintf("%s must be between %v and %v / %s은(는) %v와 %v 사이여야 합니다", v.fieldName, min, max, v.fieldName, min, max))
}

// Positive validates that the numeric value is strictly greater than zero.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// Positive는 숫자 값이 0보다 엄격하게 큰지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val > 0
//     val > 0이면 허용
//   - Zero is NOT valid (use NonZero for zero/positive)
//     0은 유효하지 않음 (0/양수는 NonZero 사용)
//   - Converts value to float64 for comparison
//     비교를 위해 값을 float64로 변환
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Price validation (must be > 0)
//     가격 검증 (0보다 커야 함)
//   - Quantity validation / 수량 검증
//   - Age validation / 나이 검증
//   - Positive measurements / 양수 측정값
//   - Count/frequency values / 횟수/빈도 값
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Single numeric comparison
//     단일 숫자 비교
//
// Example / 예제:
//
//	v := validation.New(10, "price")
//	v.Positive()  // Passes / 성공
//
//	v := validation.New(0, "price")
//	v.Positive()  // Fails (zero not positive) / 실패 (0은 양수 아님)
//
//	v := validation.New(-5, "price")
//	v.Positive()  // Fails / 실패
//
//	v := validation.New(0.01, "amount")
//	v.Positive()  // Passes / 성공
//
//	// Product price / 제품 가격
//	v := validation.New(price, "price")
//	v.Required().Positive().Max(999999.99)
//
//	// Quantity / 수량
//	v := validation.New(quantity, "quantity")
//	v.Positive().Min(1).Max(1000)
func (v *Validator) Positive() *Validator {
	return validateNumeric(v, "positive", func(val float64) bool {
		return val > 0
	}, fmt.Sprintf("%s must be positive / %s은(는) 양수여야 합니다", v.fieldName, v.fieldName))
}

// Negative validates that the numeric value is strictly less than zero.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// Negative는 숫자 값이 0보다 엄격하게 작은지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val < 0
//     val < 0이면 허용
//   - Zero is NOT valid (use NonZero for zero/negative)
//     0은 유효하지 않음 (0/음수는 NonZero 사용)
//   - Converts value to float64 for comparison
//     비교를 위해 값을 float64로 변환
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Debt/loss values / 부채/손실 값
//   - Temperature below zero / 영하 온도
//   - Negative offsets / 음수 오프셋
//   - Coordinate systems (negative axes)
//     좌표계 (음수 축)
//   - Financial losses / 재무 손실
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Single numeric comparison
//     단일 숫자 비교
//
// Example / 예제:
//
//	v := validation.New(-10, "loss")
//	v.Negative()  // Passes / 성공
//
//	v := validation.New(0, "loss")
//	v.Negative()  // Fails (zero not negative) / 실패 (0은 음수 아님)
//
//	v := validation.New(5, "loss")
//	v.Negative()  // Fails / 실패
//
//	v := validation.New(-0.01, "adjustment")
//	v.Negative()  // Passes / 성공
//
//	// Financial loss / 재무 손실
//	v := validation.New(loss, "loss")
//	v.Negative().Max(-0.01)
//
//	// Temperature below freezing / 빙점 이하 온도
//	v := validation.New(temp, "temperature")
//	v.Negative().Between(-50, -0.01)
func (v *Validator) Negative() *Validator {
	return validateNumeric(v, "negative", func(val float64) bool {
		return val < 0
	}, fmt.Sprintf("%s must be negative / %s은(는) 음수여야 합니다", v.fieldName, v.fieldName))
}

// Zero validates that the numeric value is exactly equal to zero.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// Zero는 숫자 값이 정확히 0과 같은지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val == 0
//     val == 0이면 허용
//   - Exact equality comparison
//     정확한 동등 비교
//   - Converts value to float64 for comparison
//     비교를 위해 값을 float64로 변환
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Initial state validation / 초기 상태 검증
//   - Reset value verification / 초기화 값 확인
//   - Default value checking / 기본값 확인
//   - Zero balance verification / 잔액 0 확인
//   - Counter reset validation / 카운터 초기화 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Single numeric comparison
//     단일 숫자 비교
//
// Floating Point Considerations / 부동 소수점 고려사항:
//   - Direct equality check, no epsilon tolerance
//     직접 동등 검사, 엡실론 허용 없음
//   - May have precision issues with float operations
//     부동 소수점 연산에서 정밀도 문제 발생 가능
//   - For floating point, consider Between(-0.0001, 0.0001)
//     부동 소수점의 경우 Between(-0.0001, 0.0001) 고려
//
// Example / 예제:
//
//	v := validation.New(0, "counter")
//	v.Zero()  // Passes / 성공
//
//	v := validation.New(1, "counter")
//	v.Zero()  // Fails / 실패
//
//	v := validation.New(-1, "counter")
//	v.Zero()  // Fails / 실패
//
//	v := validation.New(0.0, "balance")
//	v.Zero()  // Passes / 성공
//
//	// Balance verification / 잔액 확인
//	v := validation.New(balance, "balance")
//	v.Zero() // Must be settled
//
//	// Counter reset / 카운터 초기화
//	v := validation.New(counter, "counter")
//	v.Zero() // Must be reset
func (v *Validator) Zero() *Validator {
	return validateNumeric(v, "zero", func(val float64) bool {
		return val == 0
	}, fmt.Sprintf("%s must be zero / %s은(는) 0이어야 합니다", v.fieldName, v.fieldName))
}

// NonZero validates that the numeric value is not equal to zero.
// Works with all numeric types: int, int8-64, uint, uint8-64, float32, float64.
//
// NonZero는 숫자 값이 0이 아닌지 검증합니다.
// 모든 숫자 타입에서 작동합니다: int, int8-64, uint, uint8-64, float32, float64.
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Accepts value if val != 0
//     val != 0이면 허용
//   - Accepts both positive and negative values
//     양수와 음수 모두 허용
//   - Converts value to float64 for comparison
//     비교를 위해 값을 float64로 변환
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Ensure value has been set / 값이 설정되었는지 확인
//   - Prevent division by zero / 0으로 나누기 방지
//   - Non-default value validation / 비기본값 검증
//   - Active state verification / 활성 상태 확인
//   - Required numeric value / 필수 숫자 값
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Single numeric comparison
//     단일 숫자 비교
//
// Example / 예제:
//
//	v := validation.New(10, "divisor")
//	v.NonZero()  // Passes / 성공
//
//	v := validation.New(-5, "divisor")
//	v.NonZero()  // Passes / 성공
//
//	v := validation.New(0, "divisor")
//	v.NonZero()  // Fails / 실패
//
//	v := validation.New(0.001, "factor")
//	v.NonZero()  // Passes / 성공
//
//	// Division safety / 나누기 안전성
//	v := validation.New(divisor, "divisor")
//	v.NonZero() // Prevent division by zero
//
//	// Quantity must be set / 수량 설정 필수
//	v := validation.New(quantity, "quantity")
//	v.NonZero().Positive()
func (v *Validator) NonZero() *Validator {
	return validateNumeric(v, "nonzero", func(val float64) bool {
		return val != 0
	}, fmt.Sprintf("%s must not be zero / %s은(는) 0이 아니어야 합니다", v.fieldName, v.fieldName))
}

// Even validates that the numeric value is an even number (divisible by 2).
// Converts to integer before checking, so decimal parts are truncated.
//
// Even은 숫자 값이 짝수(2로 나누어떨어짐)인지 검증합니다.
// 검사 전 정수로 변환하므로 소수 부분은 버려집니다.
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Converts value to int before checking
//     검사 전 int로 변환
//   - Truncates decimal part (2.9 becomes 2)
//     소수 부분 버림 (2.9는 2가 됨)
//   - Uses modulo operation: val % 2 == 0
//     나머지 연산 사용: val % 2 == 0
//   - Zero is considered even
//     0은 짝수로 간주
//   - Negative even numbers are valid (-2, -4, etc.)
//     음수 짝수도 유효 (-2, -4 등)
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Even number requirements / 짝수 요구사항
//   - Pair-based calculations / 쌍 기반 계산
//   - Alternating patterns / 교대 패턴
//   - Divisibility validation / 나누기 가능성 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Single modulo operation
//     단일 나머지 연산
//
// Important Note / 중요 참고:
//   - Float values are truncated, not rounded
//     부동 소수점 값은 반올림이 아닌 버림
//   - 2.9 becomes 2 (even), not 3 (odd)
//     2.9는 3(홀수)이 아닌 2(짝수)가 됨
//   - For precise integer checking, use integer types
//     정확한 정수 검사를 위해 정수 타입 사용
//
// Example / 예제:
//
//	v := validation.New(4, "value")
//	v.Even()  // Passes / 성공
//
//	v := validation.New(3, "value")
//	v.Even()  // Fails / 실패
//
//	v := validation.New(0, "value")
//	v.Even()  // Passes (zero is even) / 성공 (0은 짝수)
//
//	v := validation.New(-2, "value")
//	v.Even()  // Passes / 성공
//
//	v := validation.New(2.9, "value")
//	v.Even()  // Passes (truncated to 2) / 성공 (2로 버림)
//
//	// Pair validation / 쌍 검증
//	v := validation.New(count, "pair_count")
//	v.Positive().Even()
func (v *Validator) Even() *Validator {
	return validateNumeric(v, "even", func(val float64) bool {
		return int(val)%2 == 0
	}, fmt.Sprintf("%s must be an even number / %s은(는) 짝수여야 합니다", v.fieldName, v.fieldName))
}

// Odd validates that the numeric value is an odd number (not divisible by 2).
// Converts to integer before checking, so decimal parts are truncated.
//
// Odd는 숫자 값이 홀수(2로 나누어떨어지지 않음)인지 검증합니다.
// 검사 전 정수로 변환하므로 소수 부분은 버려집니다.
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Converts value to int before checking
//     검사 전 int로 변환
//   - Truncates decimal part (3.9 becomes 3)
//     소수 부분 버림 (3.9는 3이 됨)
//   - Uses modulo operation: val % 2 != 0
//     나머지 연산 사용: val % 2 != 0
//   - Zero is NOT odd (it's even)
//     0은 홀수가 아님 (짝수임)
//   - Negative odd numbers are valid (-1, -3, etc.)
//     음수 홀수도 유효 (-1, -3 등)
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Odd number requirements / 홀수 요구사항
//   - Alternating patterns / 교대 패턴
//   - Index validation (odd indices)
//     인덱스 검증 (홀수 인덱스)
//   - Mathematical constraints / 수학적 제약
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Single modulo operation
//     단일 나머지 연산
//
// Important Note / 중요 참고:
//   - Float values are truncated, not rounded
//     부동 소수점 값은 반올림이 아닌 버림
//   - 3.1 becomes 3 (odd), not 4 (even)
//     3.1은 4(짝수)가 아닌 3(홀수)가 됨
//   - For precise integer checking, use integer types
//     정확한 정수 검사를 위해 정수 타입 사용
//
// Example / 예제:
//
//	v := validation.New(3, "value")
//	v.Odd()  // Passes / 성공
//
//	v := validation.New(4, "value")
//	v.Odd()  // Fails / 실패
//
//	v := validation.New(0, "value")
//	v.Odd()  // Fails (zero is even) / 실패 (0은 짝수)
//
//	v := validation.New(-3, "value")
//	v.Odd()  // Passes / 성공
//
//	v := validation.New(3.9, "value")
//	v.Odd()  // Passes (truncated to 3) / 성공 (3으로 버림)
//
//	// Page number validation / 페이지 번호 검증
//	v := validation.New(page, "page")
//	v.Positive().Odd()
func (v *Validator) Odd() *Validator {
	return validateNumeric(v, "odd", func(val float64) bool {
		return int(val)%2 != 0
	}, fmt.Sprintf("%s must be an odd number / %s은(는) 홀수여야 합니다", v.fieldName, v.fieldName))
}

// MultipleOf validates that the numeric value is a multiple of the specified number.
// Converts both values to integers before checking divisibility.
//
// MultipleOf는 숫자 값이 지정된 숫자의 배수인지 검증합니다.
// 나누기 가능성 검사 전 두 값을 정수로 변환합니다.
//
// Parameters / 매개변수:
//   - n: The divisor (must not be zero)
//     제수 (0이면 안 됨)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Converts both value and n to int
//     값과 n을 모두 int로 변환
//   - Fails if n is zero (division by zero)
//     n이 0이면 실패 (0으로 나누기)
//   - Uses modulo operation: val % n == 0
//     나머지 연산 사용: val % n == 0
//   - Truncates decimal parts
//     소수 부분 버림
//   - Skips validation if value is not numeric
//     값이 숫자가 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Batch size validation (multiples of 10, 100)
//     배치 크기 검증 (10, 100의 배수)
//   - Increment validation / 증분 검증
//   - Package quantity (multiples of 6, 12)
//     패키지 수량 (6, 12의 배수)
//   - Time intervals (multiples of 5, 15, 30)
//     시간 간격 (5, 15, 30의 배수)
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Single modulo operation
//     단일 나머지 연산
//
// Important Note / 중요 참고:
//   - Both value and n are converted to int
//     값과 n 모두 int로 변환
//   - Float precision is lost (10.5, 2.5 -> 10, 2)
//     부동 소수점 정밀도 손실 (10.5, 2.5 -> 10, 2)
//   - n = 0 always returns false (no division by zero)
//     n = 0은 항상 false 반환 (0으로 나누기 없음)
//
// Example / 예제:
//
//	v := validation.New(10, "quantity")
//	v.MultipleOf(5)  // Passes (10 = 5 * 2) / 성공
//
//	v := validation.New(12, "quantity")
//	v.MultipleOf(5)  // Fails / 실패
//
//	v := validation.New(0, "value")
//	v.MultipleOf(5)  // Passes (0 = 5 * 0) / 성공
//
//	v := validation.New(15, "minutes")
//	v.MultipleOf(15)  // Passes / 성공
//
//	v := validation.New(10, "value")
//	v.MultipleOf(0)  // Fails (division by zero) / 실패 (0으로 나누기)
//
//	// Batch size / 배치 크기
//	v := validation.New(quantity, "batch_size")
//	v.Positive().MultipleOf(100)
//
//	// Time interval in minutes / 분 단위 시간 간격
//	v := validation.New(minutes, "interval")
//	v.Between(0, 60).MultipleOf(5)
func (v *Validator) MultipleOf(n float64) *Validator {
	return validateNumeric(v, "multipleof", func(val float64) bool {
		if n == 0 {
			return false
		}
		return int(val)%int(n) == 0
	}, fmt.Sprintf("%s must be a multiple of %v / %s은(는) %v의 배수여야 합니다", v.fieldName, n, v.fieldName, n))
}
