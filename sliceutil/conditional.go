package sliceutil

// conditional.go provides conditional replacement and update operations for slices.
//
// This file implements operations that selectively modify elements based on
// conditions, creating new slices with targeted transformations:
//
// Conditional Replacement Operations:
//
// ReplaceIf (Predicate-Based Replacement):
//   - ReplaceIf(slice, predicate, newValue): Replace matching elements with fixed value
//     Time: O(n), Space: O(n)
//     Tests each element with predicate
//     Matching elements replaced with newValue
//     Non-matching elements copied unchanged
//     Immutable: Creates new slice
//     Example: ReplaceIf([1,2,3,4,5,6], isEven, 0) → [1,0,3,0,5,0]
//     Use cases: Sanitization, masking, normalization
//
// ReplaceAll (Value-Based Replacement):
//   - ReplaceAll(slice, oldValue, newValue): Replace all occurrences of specific value
//     Time: O(n), Space: O(n)
//     Tests each element for equality with oldValue
//     All matching values replaced with newValue
//     Simpler than ReplaceIf when replacing specific value
//     Requires comparable type
//     Immutable: Creates new slice
//     Example: ReplaceAll([1,2,3,2,4,2], 2, 99) → [1,99,3,99,4,99]
//     Use cases: Value substitution, error code replacement, placeholder replacement
//
// UpdateWhere (Functional Transformation):
//   - UpdateWhere(slice, predicate, updater): Apply function to matching elements
//     Time: O(n), Space: O(n)
//     Most flexible conditional operation
//     Matching elements transformed by updater function
//     Non-matching elements copied unchanged
//     Updater can access original element values
//     Works with any type (not limited to comparable)
//     Immutable: Creates new slice
//     Example: UpdateWhere([1,2,3,4,5], isEven, n → n*10) → [1,20,3,40,5]
//     Use cases: Conditional calculation, partial updates, state transitions
//
// Operation Comparison:
//
// ReplaceIf vs ReplaceAll:
//   - ReplaceIf: Uses predicate function (flexible condition)
//   - ReplaceAll: Uses value equality (simpler, comparable types only)
//   - ReplaceIf: Can match based on properties or calculations
//   - ReplaceAll: Direct value matching only
//   - Example: ReplaceIf can do "n > 10", ReplaceAll only "n == 10"
//
// ReplaceIf vs UpdateWhere:
//   - ReplaceIf: Replaces with fixed value (constant replacement)
//   - UpdateWhere: Transforms with function (dynamic transformation)
//   - ReplaceIf: Simpler when newValue is constant
//   - UpdateWhere: More flexible, can use original values
//   - Example: ReplaceIf(s, isEven, 0) vs UpdateWhere(s, isEven, n → n*2)
//
// When to Use Each:
//   - ReplaceAll: When replacing specific known value (e.g., replace null markers)
//   - ReplaceIf: When replacing based on condition with fixed value (e.g., cap outliers)
//   - UpdateWhere: When transformation depends on element value (e.g., percentage increase)
//
// Design Principles:
//   - Immutability: All operations return new slices
//   - Predictable: Always allocates full-size result slice
//   - Pure functions: No side effects on original slice
//   - Composable: Can chain with other sliceutil operations
//   - Type-safe: Generic implementations for any type
//
// Performance Characteristics:
//   - Time: All operations O(n) - must visit each element
//   - Space: All operations O(n) - must create full copy
//   - No short-circuiting: Must process entire slice
//   - Memory allocation: Single allocation of len(slice) capacity
//
// Memory Efficiency:
//   - Always allocates new slice of same size
//   - Cannot optimize away non-changed elements
//   - Pre-allocation eliminates append overhead
//   - Consider memory usage for large slices
//
// Immutability Pattern:
//   - Original slice never modified
//   - Safe for concurrent reads of original
//   - Result is independent copy
//   - Follows functional programming principles
//
// Common Usage Patterns:
//
//	// Sanitize sensitive data
//	sanitized := sliceutil.ReplaceIf(emails, func(e string) bool {
//	    return !isValidEmail(e)
//	}, "redacted@example.com")
//
//	// Remove placeholder values
//	cleaned := sliceutil.ReplaceAll(data, -1, 0)
//
//	// Apply discount to expensive items
//	discounted := sliceutil.UpdateWhere(prices, func(p float64) bool {
//	    return p > 100.0
//	}, func(p float64) float64 {
//	    return p * 0.9 // 10% discount
//	})
//
//	// Normalize outliers (cap at threshold)
//	normalized := sliceutil.ReplaceIf(measurements, func(m float64) bool {
//	    return m > threshold
//	}, threshold)
//
//	// Update user status
//	type User struct { Name string; Active bool }
//	activated := sliceutil.UpdateWhere(users,
//	    func(u User) bool { return !u.Active },
//	    func(u User) User {
//	        u.Active = true
//	        return u
//	    })
//
// Chaining Operations:
//
//	// Can compose with other operations
//	result := sliceutil.Filter(
//	    sliceutil.ReplaceIf(data, isInvalid, defaultValue),
//	    isNonZero,
//	)
//
//	// Apply multiple transformations
//	processed := sliceutil.UpdateWhere(
//	    sliceutil.ReplaceAll(raw, oldCode, newCode),
//	    needsUpdate,
//	    applyUpdate,
//	)
//
// Comparison with Standard Library:
//   - More declarative than for loops
//   - Clearer intent than manual copy + conditional logic
//   - Type-safe with generics
//   - Functional composition style
//
// Alternative Approaches:
//   - Map with ternary-like logic: Map(s, func(x) { if pred(x) return new else return x })
//   - Filter + Map combination (but changes length)
//   - Manual loop with copy (more verbose)
//
// Performance Considerations:
//   - Single pass through slice (efficient)
//   - Cannot avoid O(n) allocation for immutability
//   - For in-place modification, use manual loop
//   - Predicate/updater functions called once per element
//
// conditional.go는 슬라이스에 대한 조건부 교체 및 업데이트 작업을 제공합니다.
//
// 이 파일은 조건에 따라 선택적으로 요소를 수정하여, 목표한 변환을 통해
// 새 슬라이스를 생성하는 작업을 구현합니다:
//
// 조건부 교체 작업:
//
// ReplaceIf (조건자 기반 교체):
//   - ReplaceIf(slice, predicate, newValue): 일치하는 요소를 고정 값으로 교체
//     시간: O(n), 공간: O(n)
//     각 요소를 조건자로 테스트
//     일치하는 요소를 newValue로 교체
//     불일치 요소는 변경 없이 복사
//     불변: 새 슬라이스 생성
//     예: ReplaceIf([1,2,3,4,5,6], isEven, 0) → [1,0,3,0,5,0]
//     사용 사례: 정제, 마스킹, 정규화
//
// ReplaceAll (값 기반 교체):
//   - ReplaceAll(slice, oldValue, newValue): 특정 값의 모든 발생 교체
//     시간: O(n), 공간: O(n)
//     각 요소를 oldValue와 동등성 테스트
//     모든 일치 값을 newValue로 교체
//     특정 값 교체 시 ReplaceIf보다 단순
//     comparable 타입 필요
//     불변: 새 슬라이스 생성
//     예: ReplaceAll([1,2,3,2,4,2], 2, 99) → [1,99,3,99,4,99]
//     사용 사례: 값 대체, 오류 코드 교체, 플레이스홀더 교체
//
// UpdateWhere (함수형 변환):
//   - UpdateWhere(slice, predicate, updater): 일치하는 요소에 함수 적용
//     시간: O(n), 공간: O(n)
//     가장 유연한 조건부 작업
//     일치하는 요소를 updater 함수로 변환
//     불일치 요소는 변경 없이 복사
//     Updater가 원본 요소 값에 접근 가능
//     모든 타입 작동 (comparable 제한 없음)
//     불변: 새 슬라이스 생성
//     예: UpdateWhere([1,2,3,4,5], isEven, n → n*10) → [1,20,3,40,5]
//     사용 사례: 조건부 계산, 부분 업데이트, 상태 전환
//
// 작업 비교:
//
// ReplaceIf vs ReplaceAll:
//   - ReplaceIf: 조건자 함수 사용 (유연한 조건)
//   - ReplaceAll: 값 동등성 사용 (더 단순, comparable 타입만)
//   - ReplaceIf: 속성이나 계산 기반 일치 가능
//   - ReplaceAll: 직접 값 일치만
//   - 예: ReplaceIf는 "n > 10" 가능, ReplaceAll은 "n == 10"만
//
// ReplaceIf vs UpdateWhere:
//   - ReplaceIf: 고정 값으로 교체 (상수 교체)
//   - UpdateWhere: 함수로 변환 (동적 변환)
//   - ReplaceIf: newValue가 상수일 때 더 단순
//   - UpdateWhere: 더 유연, 원본 값 사용 가능
//   - 예: ReplaceIf(s, isEven, 0) vs UpdateWhere(s, isEven, n → n*2)
//
// 각각 사용 시기:
//   - ReplaceAll: 특정 알려진 값 교체 (예: null 마커 교체)
//   - ReplaceIf: 조건 기반 고정 값 교체 (예: 이상값 제한)
//   - UpdateWhere: 변환이 요소 값에 의존 (예: 백분율 증가)
//
// 설계 원칙:
//   - 불변성: 모든 작업은 새 슬라이스 반환
//   - 예측 가능: 항상 전체 크기 결과 슬라이스 할당
//   - 순수 함수: 원본 슬라이스에 부작용 없음
//   - 구성 가능: 다른 sliceutil 작업과 체인 가능
//   - 타입 안전: 모든 타입에 대한 제네릭 구현
//
// 성능 특성:
//   - 시간: 모든 작업 O(n) - 각 요소 방문 필요
//   - 공간: 모든 작업 O(n) - 전체 복사 생성 필요
//   - 단락 없음: 전체 슬라이스 처리 필요
//   - 메모리 할당: len(slice) 용량의 단일 할당
//
// 메모리 효율성:
//   - 항상 같은 크기의 새 슬라이스 할당
//   - 변경되지 않은 요소 최적화 불가
//   - 사전 할당으로 append 오버헤드 제거
//   - 큰 슬라이스의 메모리 사용량 고려
//
// 불변성 패턴:
//   - 원본 슬라이스 절대 수정 안 함
//   - 원본의 동시 읽기 안전
//   - 결과는 독립적인 복사본
//   - 함수형 프로그래밍 원칙 따름
//
// 일반적인 사용 패턴:
//
//	// 민감 데이터 정제
//	sanitized := sliceutil.ReplaceIf(emails, func(e string) bool {
//	    return !isValidEmail(e)
//	}, "redacted@example.com")
//
//	// 플레이스홀더 값 제거
//	cleaned := sliceutil.ReplaceAll(data, -1, 0)
//
//	// 비싼 항목에 할인 적용
//	discounted := sliceutil.UpdateWhere(prices, func(p float64) bool {
//	    return p > 100.0
//	}, func(p float64) float64 {
//	    return p * 0.9 // 10% 할인
//	})
//
//	// 이상값 정규화 (임계값에서 제한)
//	normalized := sliceutil.ReplaceIf(measurements, func(m float64) bool {
//	    return m > threshold
//	}, threshold)
//
//	// 사용자 상태 업데이트
//	type User struct { Name string; Active bool }
//	activated := sliceutil.UpdateWhere(users,
//	    func(u User) bool { return !u.Active },
//	    func(u User) User {
//	        u.Active = true
//	        return u
//	    })
//
// 작업 체인:
//
//	// 다른 작업과 구성 가능
//	result := sliceutil.Filter(
//	    sliceutil.ReplaceIf(data, isInvalid, defaultValue),
//	    isNonZero,
//	)
//
//	// 여러 변환 적용
//	processed := sliceutil.UpdateWhere(
//	    sliceutil.ReplaceAll(raw, oldCode, newCode),
//	    needsUpdate,
//	    applyUpdate,
//	)
//
// 표준 라이브러리와 비교:
//   - for 루프보다 선언적
//   - 수동 복사 + 조건 논리보다 의도 명확
//   - 제네릭으로 타입 안전
//   - 함수형 구성 스타일
//
// 대안 접근법:
//   - 삼항 유사 논리로 Map: Map(s, func(x) { if pred(x) return new else return x })
//   - Filter + Map 조합 (하지만 길이 변경)
//   - 복사를 포함한 수동 루프 (더 장황)
//
// 성능 고려사항:
//   - 슬라이스를 한 번만 통과 (효율적)
//   - 불변성을 위한 O(n) 할당 회피 불가
//   - 제자리 수정은 수동 루프 사용
//   - 조건자/업데이터 함수는 요소당 한 번 호출

// ReplaceIf returns a new slice where elements matching the predicate are replaced with newValue.
// The original slice is not modified.
//
// ReplaceIf는 조건을 만족하는 요소를 newValue로 교체한 새 슬라이스를 반환합니다.
// 원본 슬라이스는 수정되지 않습니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	result := sliceutil.ReplaceIf(numbers, func(n int) bool {
//	    return n%2 == 0
//	}, 0)
//	// [1, 0, 3, 0, 5, 0] (even numbers replaced with 0)
func ReplaceIf[T any](slice []T, predicate func(T) bool, newValue T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		if predicate(v) {
			result[i] = newValue
		} else {
			result[i] = v
		}
	}
	return result
}

// ReplaceAll returns a new slice where all occurrences of oldValue are replaced with newValue.
// The original slice is not modified.
//
// ReplaceAll은 oldValue의 모든 발생을 newValue로 교체한 새 슬라이스를 반환합니다.
// 원본 슬라이스는 수정되지 않습니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 2, 4, 2}
//	result := sliceutil.ReplaceAll(numbers, 2, 99)
//	// [1, 99, 3, 99, 4, 99]
func ReplaceAll[T comparable](slice []T, oldValue, newValue T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		if v == oldValue {
			result[i] = newValue
		} else {
			result[i] = v
		}
	}
	return result
}

// UpdateWhere returns a new slice where elements matching the predicate are updated using the updater function.
// The original slice is not modified.
//
// UpdateWhere는 조건을 만족하는 요소를 updater 함수로 업데이트한 새 슬라이스를 반환합니다.
// 원본 슬라이스는 수정되지 않습니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := sliceutil.UpdateWhere(numbers,
//	    func(n int) bool { return n%2 == 0 },
//	    func(n int) int { return n * 10 })
//	// [1, 20, 3, 40, 5] (even numbers multiplied by 10)
//
//	// Example with structs:
//	type User struct { ID int; Active bool }
//	users := []User{{1, false}, {2, true}, {3, false}}
//	result := sliceutil.UpdateWhere(users,
//	    func(u User) bool { return !u.Active },
//	    func(u User) User { u.Active = true; return u })
//	// All inactive users are now active
func UpdateWhere[T any](slice []T, predicate func(T) bool, updater func(T) T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		if predicate(v) {
			result[i] = updater(v)
		} else {
			result[i] = v
		}
	}
	return result
}
