package sliceutil

import "golang.org/x/exp/constraints"

// predicate.go provides predicate-based testing operations for slices.
//
// This file implements operations for checking conditions across entire slices,
// validating element properties, and testing slice characteristics:
//
// Quantifier Operations (Universal/Existential):
//
// All (Universal Quantifier - ∀):
//   - All(slice, predicate): Check if all elements satisfy condition
//     Time: O(n) worst case, O(1) best case, Space: O(1)
//     Short-circuits on first false (early termination)
//     Empty slice returns true (vacuous truth)
//     Logical equivalent: ∀x ∈ slice, P(x)
//     Example: All([2,4,6,8], isEven) = true
//     Use cases: Validation (all valid), constraint checking, invariant testing
//
// Any (Existential Quantifier - ∃):
//   - Any(slice, predicate): Check if at least one element satisfies condition
//     Time: O(n) worst case, O(1) best case, Space: O(1)
//     Short-circuits on first true (early termination)
//     Empty slice returns false
//     Logical equivalent: ∃x ∈ slice, P(x)
//     Example: Any([1,3,5,6], isEven) = true
//     Use cases: Existence checking, partial validation, feature detection
//
// None (Negated Existential - ¬∃):
//   - None(slice, predicate): Check if no elements satisfy condition
//     Time: O(n) worst case, O(1) best case, Space: O(1)
//     Short-circuits on first true (early termination)
//     Empty slice returns true
//     Logical equivalent: ¬∃x ∈ slice, P(x) OR ∀x ∈ slice, ¬P(x)
//     Complement of Any: None(s,p) = !Any(s,p)
//     Example: None([1,3,5,7], isEven) = true
//     Use cases: Constraint validation (none invalid), exclusion checking
//
// Logical Relationships:
//   - All(s, p) = !Any(s, !p)
//   - None(s, p) = !Any(s, p) = All(s, !p)
//   - Any(s, p) = !None(s, p) = !All(s, !p)
//
// Equality and Uniformity:
//
// AllEqual (Homogeneity Check):
//   - AllEqual(slice): Check if all elements are equal
//     Time: O(n) worst case, O(1) best case, Space: O(1)
//     Short-circuits on first difference
//     Empty slice returns true (vacuous truth)
//     Single element returns true
//     Requires comparable type
//     Example: AllEqual([5,5,5,5]) = true
//     Use cases: Constant array validation, uniform data checking
//
// Sortedness Verification:
//
// IsSortedBy (Order Validation):
//   - IsSortedBy(slice, keyFunc): Check if sorted by extracted key
//     Time: O(n), Space: O(1)
//     Ascending order only
//     Compares each element with previous
//     Empty/single element returns true
//     Key function must return Ordered type
//     Example: IsSortedBy(people, p → p.Age) checks age-sorted
//     Use cases: Verify sorted data, binary search precondition, order invariants
//
// Set Membership:
//
// ContainsAll (Subset Check):
//   - ContainsAll(slice, items...): Check if slice contains all specified items
//     Time: O(n+m) where n=slice length, m=items length, Space: O(n)
//     Builds map for O(1) lookups
//     Empty items returns true (vacuous truth)
//     All items must be found for true
//     Requires comparable type
//     Example: ContainsAll([1,2,3,4,5], 2, 4) = true
//     Use cases: Required elements checking, superset validation
//
// Performance Characteristics:
//
// Short-Circuiting (Early Termination):
//   - All: Returns false immediately when first non-matching element found
//   - Any: Returns true immediately when first matching element found
//   - None: Returns false immediately when first matching element found
//   - AllEqual: Returns false immediately when first difference found
//   - IsSortedBy: Returns false immediately when order violation found
//
// Best Case Performance:
//   - All/Any/None/AllEqual/IsSortedBy: O(1) with early termination
//   - ContainsAll: O(n+m) always (must build map first)
//
// Worst Case Performance:
//   - All: Must check all elements (all true)
//   - Any: Must check all elements (all false)
//   - None: Must check all elements (all false for predicate)
//   - AllEqual: Must check all elements (all equal)
//   - IsSortedBy: Must check all pairs (fully sorted)
//   - ContainsAll: O(n+m) always
//
// Memory Usage:
//   - All/Any/None/AllEqual/IsSortedBy: O(1) - no extra allocation
//   - ContainsAll: O(n) - builds hash map of slice elements
//
// Empty Slice Behavior (Vacuous Truth):
//   - All: true (all of nothing satisfy condition)
//   - Any: false (none of nothing satisfy condition)
//   - None: true (none of nothing satisfy condition)
//   - AllEqual: true (all of nothing are equal)
//   - IsSortedBy: true (empty is trivially sorted)
//   - ContainsAll: true if items is empty, false otherwise
//
// Design Principles:
//   - Short-circuiting: Optimize for early termination
//   - Consistent semantics: Follow mathematical logic conventions
//   - Vacuous truth: Empty slices handled per logical conventions
//   - Type safety: Generic constraints ensure correctness
//   - Predictable: No side effects, pure functions
//
// Common Usage Patterns:
//
//	// Validate all items
//	if sliceutil.All(orders, func(o Order) bool {
//	    return o.Status == "completed"
//	}) {
//	    fmt.Println("All orders completed")
//	}
//
//	// Check for any errors
//	if sliceutil.Any(results, func(r Result) bool {
//	    return r.Error != nil
//	}) {
//	    fmt.Println("Some results have errors")
//	}
//
//	// Ensure no invalid items
//	if sliceutil.None(items, func(i Item) bool {
//	    return !i.IsValid()
//	}) {
//	    fmt.Println("All items are valid")
//	}
//
//	// Check uniformity
//	if sliceutil.AllEqual(grades) {
//	    fmt.Println("All students have same grade")
//	}
//
//	// Verify sorted data for binary search
//	if sliceutil.IsSortedBy(users, func(u User) string {
//	    return u.Name
//	}) {
//	    // Safe to use binary search
//	    result := binarySearch(users, target)
//	}
//
//	// Check required elements
//	if sliceutil.ContainsAll(permissions, "read", "write") {
//	    fmt.Println("Has all required permissions")
//	}
//
// Combining Predicates:
//
//	// Combine with other operations
//	hasValidActive := sliceutil.Any(
//	    sliceutil.Filter(items, isActive),
//	    isValid,
//	)
//
//	// Logical combinations
//	allValidAndActive := sliceutil.All(items, func(i Item) bool {
//	    return i.IsValid() && i.IsActive()
//	})
//
// Comparison with Standard Library:
//   - More expressive than manual loops
//   - Short-circuiting built-in
//   - Type-safe with generics
//   - Functional composition style
//
// predicate.go는 슬라이스에 대한 조건자 기반 테스트 작업을 제공합니다.
//
// 이 파일은 전체 슬라이스에 걸쳐 조건을 확인하고, 요소 속성을 검증하며,
// 슬라이스 특성을 테스트하는 작업을 구현합니다:
//
// 한정자 작업 (전칭/존재):
//
// All (전칭 한정자 - ∀):
//   - All(slice, predicate): 모든 요소가 조건 만족하는지 확인
//     시간: O(n) 최악, O(1) 최선, 공간: O(1)
//     첫 번째 false에서 단락 (조기 종료)
//     빈 슬라이스는 true 반환 (공허한 진리)
//     논리적 동등: ∀x ∈ slice, P(x)
//     예: All([2,4,6,8], isEven) = true
//     사용 사례: 검증 (모두 유효), 제약 확인, 불변 테스트
//
// Any (존재 한정자 - ∃):
//   - Any(slice, predicate): 하나 이상의 요소가 조건 만족하는지 확인
//     시간: O(n) 최악, O(1) 최선, 공간: O(1)
//     첫 번째 true에서 단락 (조기 종료)
//     빈 슬라이스는 false 반환
//     논리적 동등: ∃x ∈ slice, P(x)
//     예: Any([1,3,5,6], isEven) = true
//     사용 사례: 존재 확인, 부분 검증, 기능 감지
//
// None (부정 존재 - ¬∃):
//   - None(slice, predicate): 어떤 요소도 조건 만족 안 하는지 확인
//     시간: O(n) 최악, O(1) 최선, 공간: O(1)
//     첫 번째 true에서 단락 (조기 종료)
//     빈 슬라이스는 true 반환
//     논리적 동등: ¬∃x ∈ slice, P(x) OR ∀x ∈ slice, ¬P(x)
//     Any의 보수: None(s,p) = !Any(s,p)
//     예: None([1,3,5,7], isEven) = true
//     사용 사례: 제약 검증 (무효 없음), 배제 확인
//
// 논리적 관계:
//   - All(s, p) = !Any(s, !p)
//   - None(s, p) = !Any(s, p) = All(s, !p)
//   - Any(s, p) = !None(s, p) = !All(s, !p)
//
// 동등성 및 균일성:
//
// AllEqual (동질성 확인):
//   - AllEqual(slice): 모든 요소가 같은지 확인
//     시간: O(n) 최악, O(1) 최선, 공간: O(1)
//     첫 번째 차이에서 단락
//     빈 슬라이스는 true 반환 (공허한 진리)
//     단일 요소는 true 반환
//     comparable 타입 필요
//     예: AllEqual([5,5,5,5]) = true
//     사용 사례: 상수 배열 검증, 균일 데이터 확인
//
// 정렬 검증:
//
// IsSortedBy (순서 검증):
//   - IsSortedBy(slice, keyFunc): 추출된 키로 정렬되었는지 확인
//     시간: O(n), 공간: O(1)
//     오름차순만
//     각 요소를 이전 요소와 비교
//     빈/단일 요소는 true 반환
//     키 함수는 Ordered 타입 반환 필요
//     예: IsSortedBy(people, p → p.Age) 나이 정렬 확인
//     사용 사례: 정렬 데이터 확인, 이진 검색 전제 조건, 순서 불변
//
// 집합 멤버십:
//
// ContainsAll (부분 집합 확인):
//   - ContainsAll(slice, items...): 슬라이스가 지정된 모든 항목 포함하는지 확인
//     시간: O(n+m) (n=슬라이스 길이, m=항목 길이), 공간: O(n)
//     O(1) 조회를 위한 맵 구축
//     빈 items는 true 반환 (공허한 진리)
//     모든 항목을 찾아야 true
//     comparable 타입 필요
//     예: ContainsAll([1,2,3,4,5], 2, 4) = true
//     사용 사례: 필수 요소 확인, 상위 집합 검증
//
// 성능 특성:
//
// 단락 평가 (조기 종료):
//   - All: 첫 번째 불일치 요소 발견 시 즉시 false 반환
//   - Any: 첫 번째 일치 요소 발견 시 즉시 true 반환
//   - None: 첫 번째 일치 요소 발견 시 즉시 false 반환
//   - AllEqual: 첫 번째 차이 발견 시 즉시 false 반환
//   - IsSortedBy: 순서 위반 발견 시 즉시 false 반환
//
// 최선 성능:
//   - All/Any/None/AllEqual/IsSortedBy: 조기 종료로 O(1)
//   - ContainsAll: 항상 O(n+m) (먼저 맵 구축 필요)
//
// 최악 성능:
//   - All: 모든 요소 확인 필요 (모두 true)
//   - Any: 모든 요소 확인 필요 (모두 false)
//   - None: 모든 요소 확인 필요 (조건자에 대해 모두 false)
//   - AllEqual: 모든 요소 확인 필요 (모두 같음)
//   - IsSortedBy: 모든 쌍 확인 필요 (완전 정렬)
//   - ContainsAll: 항상 O(n+m)
//
// 메모리 사용:
//   - All/Any/None/AllEqual/IsSortedBy: O(1) - 추가 할당 없음
//   - ContainsAll: O(n) - 슬라이스 요소의 해시 맵 구축
//
// 빈 슬라이스 동작 (공허한 진리):
//   - All: true (아무것도 없는 것의 모두가 조건 만족)
//   - Any: false (아무것도 없는 것 중 아무것도 조건 만족 안 함)
//   - None: true (아무것도 없는 것 중 아무것도 조건 만족 안 함)
//   - AllEqual: true (아무것도 없는 것의 모두가 같음)
//   - IsSortedBy: true (빈 것은 자명하게 정렬됨)
//   - ContainsAll: items가 비었으면 true, 아니면 false
//
// 설계 원칙:
//   - 단락 평가: 조기 종료 최적화
//   - 일관된 의미론: 수학적 논리 규칙 따름
//   - 공허한 진리: 논리 규칙에 따라 빈 슬라이스 처리
//   - 타입 안전성: 제네릭 제약 조건으로 정확성 보장
//   - 예측 가능: 부작용 없음, 순수 함수
//
// 일반적인 사용 패턴:
//
//	// 모든 항목 검증
//	if sliceutil.All(orders, func(o Order) bool {
//	    return o.Status == "completed"
//	}) {
//	    fmt.Println("모든 주문 완료")
//	}
//
//	// 에러 확인
//	if sliceutil.Any(results, func(r Result) bool {
//	    return r.Error != nil
//	}) {
//	    fmt.Println("일부 결과에 에러 있음")
//	}
//
//	// 무효 항목 없는지 확인
//	if sliceutil.None(items, func(i Item) bool {
//	    return !i.IsValid()
//	}) {
//	    fmt.Println("모든 항목 유효")
//	}
//
//	// 균일성 확인
//	if sliceutil.AllEqual(grades) {
//	    fmt.Println("모든 학생이 같은 학점")
//	}
//
//	// 이진 검색을 위한 정렬 데이터 확인
//	if sliceutil.IsSortedBy(users, func(u User) string {
//	    return u.Name
//	}) {
//	    // 이진 검색 사용 안전
//	    result := binarySearch(users, target)
//	}
//
//	// 필수 요소 확인
//	if sliceutil.ContainsAll(permissions, "read", "write") {
//	    fmt.Println("필요한 모든 권한 있음")
//	}
//
// 조건자 결합:
//
//	// 다른 작업과 결합
//	hasValidActive := sliceutil.Any(
//	    sliceutil.Filter(items, isActive),
//	    isValid,
//	)
//
//	// 논리적 조합
//	allValidAndActive := sliceutil.All(items, func(i Item) bool {
//	    return i.IsValid() && i.IsActive()
//	})
//
// 표준 라이브러리와 비교:
//   - 수동 루프보다 표현력 높음
//   - 단락 평가 내장
//   - 제네릭으로 타입 안전
//   - 함수형 구성 스타일

// All checks if all elements in the slice satisfy the predicate.
// All은 슬라이스의 모든 요소가 조건을 만족하는지 확인합니다.
//
// Returns true if all elements satisfy the predicate, false otherwise.
// 모든 요소가 조건을 만족하면 true를, 그렇지 않으면 false를 반환합니다.
//
// An empty slice returns true (vacuous truth).
// 비어있는 슬라이스는 true를 반환합니다 (공허한 진리).
//
// Example
// 예제:
//
//	numbers := []int{2, 4, 6, 8}
//	allEven := sliceutil.All(numbers, func(n int) bool { return n%2 == 0 })
//	// allEven: true
//
//	numbers2 := []int{2, 4, 5, 8}
//	allEven2 := sliceutil.All(numbers2, func(n int) bool { return n%2 == 0 })
//	// allEven2: false
func All[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Any checks if at least one element in the slice satisfies the predicate.
// Any는 슬라이스의 최소한 하나의 요소가 조건을 만족하는지 확인합니다.
//
// Returns true if at least one element satisfies the predicate, false otherwise.
// 최소한 하나의 요소가 조건을 만족하면 true를, 그렇지 않으면 false를 반환합니다.
//
// An empty slice returns false.
// 비어있는 슬라이스는 false를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 3, 5, 6}
//	hasEven := sliceutil.Any(numbers, func(n int) bool { return n%2 == 0 })
//	// hasEven: true
//
//	numbers2 := []int{1, 3, 5, 7}
//	hasEven2 := sliceutil.Any(numbers2, func(n int) bool { return n%2 == 0 })
//	// hasEven2: false
func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

// None checks if no elements in the slice satisfy the predicate.
// None은 슬라이스의 어떤 요소도 조건을 만족하지 않는지 확인합니다.
//
// Returns true if no elements satisfy the predicate, false otherwise.
// 어떤 요소도 조건을 만족하지 않으면 true를, 그렇지 않으면 false를 반환합니다.
//
// An empty slice returns true.
// 비어있는 슬라이스는 true를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 3, 5, 7}
//	noEven := sliceutil.None(numbers, func(n int) bool { return n%2 == 0 })
//	// noEven: true
//
//	numbers2 := []int{1, 3, 5, 6}
//	noEven2 := sliceutil.None(numbers2, func(n int) bool { return n%2 == 0 })
//	// noEven2: false
func None[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return false
		}
	}
	return true
}

// AllEqual checks if all elements in the slice are equal.
// AllEqual은 슬라이스의 모든 요소가 같은지 확인합니다.
//
// Returns true if all elements are equal, false otherwise.
// 모든 요소가 같으면 true를, 그렇지 않으면 false를 반환합니다.
//
// An empty slice returns true.
// 비어있는 슬라이스는 true를 반환합니다.
//
// A slice with one element returns true.
// 요소가 하나인 슬라이스는 true를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{5, 5, 5, 5}
//	allSame := sliceutil.AllEqual(numbers)
//	// allSame: true
//
//	numbers2 := []int{5, 5, 6, 5}
//	allSame2 := sliceutil.AllEqual(numbers2)
//	// allSame2: false
func AllEqual[T comparable](slice []T) bool {
	if len(slice) <= 1 {
		return true
	}

	first := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] != first {
			return false
		}
	}

	return true
}

// IsSortedBy checks if the slice is sorted in ascending order by the key extracted by keyFunc.
// IsSortedBy는 keyFunc으로 추출한 키를 기준으로 슬라이스가 오름차순으로 정렬되어 있는지 확인합니다.
//
// Returns true if the slice is sorted in ascending order by the extracted key, false otherwise.
// 추출한 키를 기준으로 슬라이스가 오름차순으로 정렬되어 있으면 true를, 그렇지 않으면 false를 반환합니다.
//
// An empty slice or a slice with one element is considered sorted.
// 비어있는 슬라이스나 요소가 하나인 슬라이스는 정렬된 것으로 간주됩니다.
//
// Example
// 예제:
//
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	people := []Person{
//	    {"Alice", 25},
//	    {"Bob", 30},
//	    {"Charlie", 35},
//	}
//	isSortedByAge := sliceutil.IsSortedBy(people, func(p Person) int { return p.Age })
//	// isSortedByAge: true
func IsSortedBy[T any, K constraints.Ordered](slice []T, keyFunc func(T) K) bool {
	if len(slice) <= 1 {
		return true
	}

	for i := 1; i < len(slice); i++ {
		if keyFunc(slice[i]) < keyFunc(slice[i-1]) {
			return false
		}
	}

	return true
}

// ContainsAll checks if the slice contains all of the specified items.
// ContainsAll은 슬라이스가 지정된 모든 항목을 포함하는지 확인합니다.
//
// Returns true if the slice contains all specified items, false otherwise.
// 슬라이스가 지정된 모든 항목을 포함하면 true를, 그렇지 않으면 false를 반환합니다.
//
// If no items are specified, returns true.
// 항목이 지정되지 않은 경우 true를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasAll := sliceutil.ContainsAll(numbers, 2, 4)
//	// hasAll: true
//
//	hasAll2 := sliceutil.ContainsAll(numbers, 2, 6)
//	// hasAll2: false
func ContainsAll[T comparable](slice []T, items ...T) bool {
	if len(items) == 0 {
		return true
	}

	// Create a map for O(1) lookup
	// O(1) 조회를 위한 맵 생성
	sliceMap := make(map[T]bool, len(slice))
	for _, item := range slice {
		sliceMap[item] = true
	}

	// Check if all items are in the map
	// 모든 항목이 맵에 있는지 확인
	for _, item := range items {
		if !sliceMap[item] {
			return false
		}
	}

	return true
}
