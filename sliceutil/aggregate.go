package sliceutil

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// aggregate.go provides aggregation and reduction operations for slices.
//
// This file implements operations that combine multiple elements into single
// values, group elements by criteria, or compute statistical measures:
//
// Reduction Operations:
//
// Reduce (Left-to-Right Accumulation):
//   - Reduce(slice, initial, fn): Fold left with accumulator
//     Time: O(n), Space: O(1)
//     Processes elements from left to right
//     Most common reduction pattern
//     Accumulator type can differ from element type
//     Example: Sum: Reduce([1,2,3], 0, (acc,n) → acc+n) = 6
//     Use cases: Summing, concatenation, building complex structures
//
// ReduceRight (Right-to-Left Accumulation):
//   - ReduceRight(slice, initial, fn): Fold right with accumulator
//     Time: O(n), Space: O(1)
//     Processes elements from right to left
//     Result differs from Reduce for non-commutative operations
//     Example: Concat: ReduceRight(["a","b","c"], "", concat) = "cba"
//     Use cases: Reverse operations, building nested structures
//
// Basic Aggregations:
//
// Sum (Numeric Total):
//   - Sum(slice): Add all numeric elements
//     Time: O(n), Space: O(1)
//     Works with Integer and Float constraints
//     Returns zero for empty slice
//     No overflow protection for integers
//     Example: Sum([1,2,3,4,5]) = 15
//     Use cases: Total calculation, numeric accumulation
//
// Min (Minimum Element):
//   - Min(slice): Find smallest element
//     Time: O(n), Space: O(1)
//     Works with Ordered constraint (comparable types)
//     Returns error for empty slice
//     Single pass through slice
//     Example: Min([3,1,4,1,5]) = 1, nil
//     Use cases: Finding minimum value, range calculations
//
// Max (Maximum Element):
//   - Max(slice): Find largest element
//     Time: O(n), Space: O(1)
//     Works with Ordered constraint
//     Returns error for empty slice
//     Single pass through slice
//     Example: Max([3,1,4,1,5]) = 5, nil
//     Use cases: Finding maximum value, range calculations
//
// Average (Mean Value):
//   - Average(slice): Compute arithmetic mean
//     Time: O(n), Space: O(1)
//     Works with Integer and Float constraints
//     Always returns float64
//     Returns 0 for empty slice (no error)
//     Uses Sum internally
//     Example: Average([1,2,3,4,5]) = 3.0
//     Use cases: Statistical analysis, data summarization
//
// Grouping Operations:
//
// GroupBy (Group by Key):
//   - GroupBy(slice, keyFunc): Create map of key → []elements
//     Time: O(n), Space: O(n)
//     Groups elements sharing same key
//     Key function must return comparable type
//     Preserves element order within groups
//     Dynamic allocation as groups are found
//     Example: GroupBy([1,2,3,4], n → n%2) = {0:[2,4], 1:[1,3]}
//     Use cases: Categorization, data organization, histogram building
//
// CountBy (Count by Key):
//   - CountBy(slice, keyFunc): Create map of key → count
//     Time: O(n), Space: O(k) where k = unique keys
//     Similar to GroupBy but stores counts instead of elements
//     More memory efficient than GroupBy when only counts needed
//     Example: CountBy(["a","b","a","c","a"], identity) = {"a":3, "b":1, "c":1}
//     Use cases: Frequency analysis, histogram, occurrence counting
//
// Partition (Binary Split):
//   - Partition(slice, predicate): Split into (matching, non-matching)
//     Time: O(n), Space: O(n)
//     Returns two slices based on predicate
//     First slice: elements where predicate = true
//     Second slice: elements where predicate = false
//     Both slices pre-allocated to input capacity
//     Preserves relative order in both slices
//     Example: Partition([1,2,3,4,5,6], isEven) = ([2,4,6], [1,3,5])
//     Use cases: Binary classification, validation/error separation
//
// Custom Comparisons:
//
// MinBy (Minimum by Custom Key):
//   - MinBy(slice, keyFunc): Find element with minimum key
//     Time: O(n), Space: O(1)
//     Returns entire element, not just key
//     Key function must return Ordered type
//     Returns error for empty slice
//     Single pass with comparison by extracted key
//     Example: MinBy(people, p → p.Age) = youngest person
//     Use cases: Finding object with minimum property, custom sorting criteria
//
// MaxBy (Maximum by Custom Key):
//   - MaxBy(slice, keyFunc): Find element with maximum key
//     Time: O(n), Space: O(1)
//     Returns entire element, not just key
//     Key function must return Ordered type
//     Returns error for empty slice
//     Single pass with comparison by extracted key
//     Example: MaxBy(people, p → p.Age) = oldest person
//     Use cases: Finding object with maximum property, custom sorting criteria
//
// Design Principles:
//   - Functional: Higher-order functions for flexibility
//   - Type-safe: Generic constraints ensure compile-time safety
//   - Efficient: Single-pass algorithms where possible
//   - Predictable: Consistent error handling (empty slice → error for Min/Max)
//   - Composable: Can combine with other sliceutil operations
//
// Error Handling:
//   - Min/Max/MinBy/MaxBy: Return error for empty slices
//   - Sum/Average: Return zero values for empty slices (no error)
//   - GroupBy/CountBy/Partition: Return empty results for empty slices
//   - Reduce/ReduceRight: Depend on initial value, work with empty slices
//
// Performance Characteristics:
//   - All operations: O(n) time complexity
//   - Reduce/Sum/Min/Max/Average: O(1) space (in-place aggregation)
//   - GroupBy/Partition: O(n) space (new collections)
//   - CountBy: O(k) space where k = unique keys
//   - MinBy/MaxBy: O(1) space (single element tracking)
//
// Memory Allocation:
//   - Reduction operations: Minimal allocation (just result)
//   - GroupBy: Allocates map + slices for each group
//   - CountBy: Allocates map only (lighter than GroupBy)
//   - Partition: Pre-allocates both result slices with input capacity
//
// Common Usage Patterns:
//
//	// Calculate total from orders
//	total := sliceutil.Reduce(orders, 0.0, func(sum float64, order Order) float64 {
//	    return sum + order.Amount
//	})
//
//	// Find min/max efficiently
//	minPrice, _ := sliceutil.Min(prices)
//	maxPrice, _ := sliceutil.Max(prices)
//
//	// Group users by role
//	byRole := sliceutil.GroupBy(users, func(u User) string {
//	    return u.Role
//	})
//
//	// Split valid/invalid items
//	valid, invalid := sliceutil.Partition(items, func(i Item) bool {
//	    return i.IsValid()
//	})
//
//	// Count occurrences
//	wordCounts := sliceutil.CountBy(words, func(w string) string {
//	    return strings.ToLower(w)
//	})
//
//	// Find best candidate
//	best, _ := sliceutil.MaxBy(candidates, func(c Candidate) float64 {
//	    return c.Score
//	})
//
// Comparison with Standard Library:
//   - More expressive than manual loops
//   - Type-safe with generics (vs reflection)
//   - Functional composition style
//   - Reusable patterns for common operations
//
// aggregate.go는 슬라이스에 대한 집계 및 축약 작업을 제공합니다.
//
// 이 파일은 여러 요소를 단일 값으로 결합하거나, 기준에 따라 요소를
// 그룹화하거나, 통계적 측정을 계산하는 작업을 구현합니다:
//
// 축약 작업:
//
// Reduce (왼쪽에서 오른쪽 누적):
//   - Reduce(slice, initial, fn): 누적자로 왼쪽 폴드
//     시간: O(n), 공간: O(1)
//     요소를 왼쪽에서 오른쪽으로 처리
//     가장 일반적인 축약 패턴
//     누적자 타입이 요소 타입과 다를 수 있음
//     예: Sum: Reduce([1,2,3], 0, (acc,n) → acc+n) = 6
//     사용 사례: 합계, 연결, 복잡한 구조 구축
//
// ReduceRight (오른쪽에서 왼쪽 누적):
//   - ReduceRight(slice, initial, fn): 누적자로 오른쪽 폴드
//     시간: O(n), 공간: O(1)
//     요소를 오른쪽에서 왼쪽으로 처리
//     비가환 작업에서 Reduce와 결과 다름
//     예: Concat: ReduceRight(["a","b","c"], "", concat) = "cba"
//     사용 사례: 역순 작업, 중첩 구조 구축
//
// 기본 집계:
//
// Sum (숫자 합계):
//   - Sum(slice): 모든 숫자 요소 더하기
//     시간: O(n), 공간: O(1)
//     Integer 및 Float 제약 조건으로 작동
//     빈 슬라이스에 대해 0 반환
//     정수에 대한 오버플로우 보호 없음
//     예: Sum([1,2,3,4,5]) = 15
//     사용 사례: 총계 계산, 숫자 누적
//
// Min (최소 요소):
//   - Min(slice): 가장 작은 요소 찾기
//     시간: O(n), 공간: O(1)
//     Ordered 제약 조건으로 작동 (비교 가능 타입)
//     빈 슬라이스에 대해 에러 반환
//     슬라이스를 한 번만 통과
//     예: Min([3,1,4,1,5]) = 1, nil
//     사용 사례: 최소값 찾기, 범위 계산
//
// Max (최대 요소):
//   - Max(slice): 가장 큰 요소 찾기
//     시간: O(n), 공간: O(1)
//     Ordered 제약 조건으로 작동
//     빈 슬라이스에 대해 에러 반환
//     슬라이스를 한 번만 통과
//     예: Max([3,1,4,1,5]) = 5, nil
//     사용 사례: 최대값 찾기, 범위 계산
//
// Average (평균값):
//   - Average(slice): 산술 평균 계산
//     시간: O(n), 공간: O(1)
//     Integer 및 Float 제약 조건으로 작동
//     항상 float64 반환
//     빈 슬라이스에 대해 0 반환 (에러 없음)
//     내부적으로 Sum 사용
//     예: Average([1,2,3,4,5]) = 3.0
//     사용 사례: 통계 분석, 데이터 요약
//
// 그룹화 작업:
//
// GroupBy (키로 그룹화):
//   - GroupBy(slice, keyFunc): 키 → []요소 맵 생성
//     시간: O(n), 공간: O(n)
//     같은 키를 공유하는 요소 그룹화
//     키 함수는 comparable 타입 반환 필요
//     그룹 내 요소 순서 유지
//     그룹 발견 시 동적 할당
//     예: GroupBy([1,2,3,4], n → n%2) = {0:[2,4], 1:[1,3]}
//     사용 사례: 분류, 데이터 정리, 히스토그램 구축
//
// CountBy (키로 개수 세기):
//   - CountBy(slice, keyFunc): 키 → 개수 맵 생성
//     시간: O(n), 공간: O(k) (k = 고유 키)
//     GroupBy와 유사하지만 요소 대신 개수 저장
//     개수만 필요할 때 GroupBy보다 메모리 효율적
//     예: CountBy(["a","b","a","c","a"], identity) = {"a":3, "b":1, "c":1}
//     사용 사례: 빈도 분석, 히스토그램, 발생 횟수 세기
//
// Partition (이진 분할):
//   - Partition(slice, predicate): (일치, 불일치)로 분할
//     시간: O(n), 공간: O(n)
//     조건자 기반으로 두 슬라이스 반환
//     첫 번째 슬라이스: predicate = true인 요소
//     두 번째 슬라이스: predicate = false인 요소
//     두 슬라이스 모두 입력 용량으로 사전 할당
//     두 슬라이스 모두 상대적 순서 유지
//     예: Partition([1,2,3,4,5,6], isEven) = ([2,4,6], [1,3,5])
//     사용 사례: 이진 분류, 검증/오류 분리
//
// 사용자 정의 비교:
//
// MinBy (사용자 정의 키로 최소):
//   - MinBy(slice, keyFunc): 최소 키를 가진 요소 찾기
//     시간: O(n), 공간: O(1)
//     키가 아닌 전체 요소 반환
//     키 함수는 Ordered 타입 반환 필요
//     빈 슬라이스에 대해 에러 반환
//     추출된 키로 비교하며 한 번만 통과
//     예: MinBy(people, p → p.Age) = 가장 젊은 사람
//     사용 사례: 최소 속성을 가진 객체 찾기, 사용자 정의 정렬 기준
//
// MaxBy (사용자 정의 키로 최대):
//   - MaxBy(slice, keyFunc): 최대 키를 가진 요소 찾기
//     시간: O(n), 공간: O(1)
//     키가 아닌 전체 요소 반환
//     키 함수는 Ordered 타입 반환 필요
//     빈 슬라이스에 대해 에러 반환
//     추출된 키로 비교하며 한 번만 통과
//     예: MaxBy(people, p → p.Age) = 가장 나이 많은 사람
//     사용 사례: 최대 속성을 가진 객체 찾기, 사용자 정의 정렬 기준
//
// 설계 원칙:
//   - 함수형: 유연성을 위한 고차 함수
//   - 타입 안전: 제네릭 제약 조건으로 컴파일 시간 안전성 보장
//   - 효율적: 가능한 곳에서 단일 패스 알고리즘
//   - 예측 가능: 일관된 에러 처리 (빈 슬라이스 → Min/Max 에러)
//   - 구성 가능: 다른 sliceutil 작업과 결합 가능
//
// 에러 처리:
//   - Min/Max/MinBy/MaxBy: 빈 슬라이스에 대해 에러 반환
//   - Sum/Average: 빈 슬라이스에 대해 0 값 반환 (에러 없음)
//   - GroupBy/CountBy/Partition: 빈 슬라이스에 대해 빈 결과 반환
//   - Reduce/ReduceRight: 초기값에 따라 다름, 빈 슬라이스로 작동
//
// 성능 특성:
//   - 모든 작업: O(n) 시간 복잡도
//   - Reduce/Sum/Min/Max/Average: O(1) 공간 (제자리 집계)
//   - GroupBy/Partition: O(n) 공간 (새 컬렉션)
//   - CountBy: O(k) 공간 (k = 고유 키)
//   - MinBy/MaxBy: O(1) 공간 (단일 요소 추적)
//
// 메모리 할당:
//   - 축약 작업: 최소 할당 (결과만)
//   - GroupBy: 각 그룹에 대해 맵 + 슬라이스 할당
//   - CountBy: 맵만 할당 (GroupBy보다 가벼움)
//   - Partition: 입력 용량으로 두 결과 슬라이스 사전 할당
//
// 일반적인 사용 패턴:
//
//	// 주문에서 총계 계산
//	total := sliceutil.Reduce(orders, 0.0, func(sum float64, order Order) float64 {
//	    return sum + order.Amount
//	})
//
//	// 최소/최대 효율적으로 찾기
//	minPrice, _ := sliceutil.Min(prices)
//	maxPrice, _ := sliceutil.Max(prices)
//
//	// 역할별 사용자 그룹화
//	byRole := sliceutil.GroupBy(users, func(u User) string {
//	    return u.Role
//	})
//
//	// 유효/무효 항목 분할
//	valid, invalid := sliceutil.Partition(items, func(i Item) bool {
//	    return i.IsValid()
//	})
//
//	// 발생 횟수 세기
//	wordCounts := sliceutil.CountBy(words, func(w string) string {
//	    return strings.ToLower(w)
//	})
//
//	// 최고 후보 찾기
//	best, _ := sliceutil.MaxBy(candidates, func(c Candidate) float64 {
//	    return c.Score
//	})
//
// 표준 라이브러리와 비교:
//   - 수동 루프보다 표현력 높음
//   - 제네릭으로 타입 안전 (vs 리플렉션)
//   - 함수형 구성 스타일
//   - 일반 작업에 대한 재사용 가능한 패턴

// Reduce applies a reducer function to accumulate a single value from the slice.
// Reduce는 슬라이스에서 단일 값을 누적하기 위해 reducer 함수를 적용합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	sum := sliceutil.Reduce(numbers, 0, func(acc, n int) int {
//	    return acc + n
//	}) // 15
//
//	words := []string{"hello", "world"}
//	combined := sliceutil.Reduce(words, "", func(acc, w string) string {
//	    return acc + w
//	}) // "helloworld"
func Reduce[T any, R any](slice []T, initial R, reducer func(R, T) R) R {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

// ReduceRight applies a reducer function from right to left to accumulate a single value.
// ReduceRight는 오른쪽에서 왼쪽으로 reducer 함수를 적용하여 단일 값을 누적합니다.
//
// Similar to Reduce, but processes elements from right to left.
// Reduce와 유사하지만 요소를 오른쪽에서 왼쪽으로 처리합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := sliceutil.ReduceRight(numbers, 0, func(acc, n int) int {
//	    return acc + n
//	}) // 15 (same as Reduce for commutative operations)
//
//	words := []string{"hello", "world"}
//	combined := sliceutil.ReduceRight(words, "", func(acc, w string) string {
//	    return acc + w
//	}) // "worldhello" (reversed compared to Reduce)
//
// // Useful for operations where order matters
// 순서가 중요한 작업에 유용
//	nested := [][]int{{1, 2}, {3, 4}, {5}}
//	flattened := sliceutil.ReduceRight(nested, []int{}, func(acc, slice []int) []int {
//	    return append(slice, acc...)
//	}) // [5, 3, 4, 1, 2]
func ReduceRight[T any, R any](slice []T, initial R, reducer func(R, T) R) R {
	result := initial
	for i := len(slice) - 1; i >= 0; i-- {
		result = reducer(result, slice[i])
	}
	return result
}

// Sum returns the sum of all elements in the slice.
// Sum은 슬라이스의 모든 요소의 합을 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	sum := sliceutil.Sum(numbers) // 15
//
//	floats := []float64{1.5, 2.5, 3.0}
//	sum := sliceutil.Sum(floats) // 7.0
func Sum[T constraints.Integer | constraints.Float](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Min returns the minimum element in the slice.
// Returns an error if the slice is empty.
// Min은 슬라이스의 최소 요소를 반환합니다.
// 슬라이스가 비어있으면 에러를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{3, 1, 4, 1, 5}
//	min, _ := sliceutil.Min(numbers) // 1
//
//	words := []string{"banana", "apple", "cherry"}
//	min, _ := sliceutil.Min(words) // "apple"
func Min[T constraints.Ordered](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("cannot find min of empty slice")
	}
	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

// Max returns the maximum element in the slice.
// Returns an error if the slice is empty.
// Max는 슬라이스의 최대 요소를 반환합니다.
// 슬라이스가 비어있으면 에러를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{3, 1, 4, 1, 5}
//	max, _ := sliceutil.Max(numbers) // 5
//
//	words := []string{"banana", "apple", "cherry"}
//	max, _ := sliceutil.Max(words) // "cherry"
func Max[T constraints.Ordered](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("cannot find max of empty slice")
	}
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

// Average returns the average of all elements in the slice.
// Returns 0 if the slice is empty.
// Average는 슬라이스의 모든 요소의 평균을 반환합니다.
// 슬라이스가 비어있으면 0을 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	avg := sliceutil.Average(numbers) // 3.0
//
//	floats := []float64{1.5, 2.5, 3.0}
//	avg := sliceutil.Average(floats) // 2.333...
func Average[T constraints.Integer | constraints.Float](slice []T) float64 {
	if len(slice) == 0 {
		return 0
	}
	sum := Sum(slice)
	return float64(sum) / float64(len(slice))
}

// GroupBy groups elements by a key function and returns a map.
// GroupBy는 키 함수로 요소를 그룹화하고 맵을 반환합니다.
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
//	    {"Charlie", 25},
//	}
//	grouped := sliceutil.GroupBy(people, func(p Person) int {
//	    return p.Age
//	}) // map[25:[{Alice 25} {Charlie 25}] 30:[{Bob 30}]]
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	grouped := sliceutil.GroupBy(numbers, func(n int) string {
//	    if n%2 == 0 {
//	        return "even"
//	    }
//	    return "odd"
//	}) // map["even":[2 4 6] "odd":[1 3 5]]
func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := keyFunc(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Partition splits a slice into two slices based on a predicate.
// The first slice contains elements that satisfy the predicate.
// The second slice contains elements that don't satisfy the predicate.
// Partition은 조건에 따라 슬라이스를 두 개의 슬라이스로 분할합니다.
// 첫 번째 슬라이스는 조건을 만족하는 요소를 포함합니다.
// 두 번째 슬라이스는 조건을 만족하지 않는 요소를 포함합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens, odds := sliceutil.Partition(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // evens: [2, 4, 6], odds: [1, 3, 5]
//
//	words := []string{"apple", "a", "banana", "pear", "ab"}
//	long, short := sliceutil.Partition(words, func(s string) bool {
//	    return len(s) > 2
//	}) // long: ["apple", "banana", "pear"], short: ["a", "ab"]
func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T) {
	trueSlice := make([]T, 0, len(slice))
	falseSlice := make([]T, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			trueSlice = append(trueSlice, v)
		} else {
			falseSlice = append(falseSlice, v)
		}
	}
	return trueSlice, falseSlice
}

// CountBy counts elements by a key function and returns a map with counts.
// CountBy는 키 함수로 요소를 세고 개수를 포함한 맵을 반환합니다.
//
// Similar to GroupBy, but returns counts instead of grouped elements.
// GroupBy와 유사하지만 그룹화된 요소 대신 개수를 반환합니다.
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
//	    {"Charlie", 25},
//	}
//	counts := sliceutil.CountBy(people, func(p Person) int {
//	    return p.Age
//	}) // map[25:2 30:1]
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	counts := sliceutil.CountBy(numbers, func(n int) string {
//	    if n%2 == 0 {
//	        return "even"
//	    }
//	    return "odd"
//	}) // map["even":3 "odd":3]
func CountBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K]int {
	result := make(map[K]int)
	for _, v := range slice {
		key := keyFunc(v)
		result[key]++
	}
	return result
}

// MinBy returns the element with the minimum key extracted by keyFunc.
// Returns an error if the slice is empty.
// MinBy는 keyFunc으로 추출한 키가 최소인 요소를 반환합니다.
// 슬라이스가 비어있으면 에러를 반환합니다.
//
// Example
// 예제:
//
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	people := []Person{
//	    {"Alice", 30},
//	    {"Bob", 25},
//	    {"Charlie", 35},
//	}
//	youngest, _ := sliceutil.MinBy(people, func(p Person) int {
//	    return p.Age
//	}) // {Bob 25}
func MinBy[T any, K constraints.Ordered](slice []T, keyFunc func(T) K) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("cannot find min of empty slice")
	}
	minItem := slice[0]
	minKey := keyFunc(slice[0])
	for i := 1; i < len(slice); i++ {
		key := keyFunc(slice[i])
		if key < minKey {
			minKey = key
			minItem = slice[i]
		}
	}
	return minItem, nil
}

// MaxBy returns the element with the maximum key extracted by keyFunc.
// Returns an error if the slice is empty.
// MaxBy는 keyFunc으로 추출한 키가 최대인 요소를 반환합니다.
// 슬라이스가 비어있으면 에러를 반환합니다.
//
// Example
// 예제:
//
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	people := []Person{
//	    {"Alice", 30},
//	    {"Bob", 25},
//	    {"Charlie", 35},
//	}
//	oldest, _ := sliceutil.MaxBy(people, func(p Person) int {
//	    return p.Age
//	}) // {Charlie 35}
func MaxBy[T any, K constraints.Ordered](slice []T, keyFunc func(T) K) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("cannot find max of empty slice")
	}
	maxItem := slice[0]
	maxKey := keyFunc(slice[0])
	for i := 1; i < len(slice); i++ {
		key := keyFunc(slice[i])
		if key > maxKey {
			maxKey = key
			maxItem = slice[i]
		}
	}
	return maxItem, nil
}
