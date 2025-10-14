package sliceutil

import "golang.org/x/exp/constraints"

// All checks if all elements in the slice satisfy the predicate.
// All은 슬라이스의 모든 요소가 조건을 만족하는지 확인합니다.
//
// Returns true if all elements satisfy the predicate, false otherwise.
// 모든 요소가 조건을 만족하면 true를, 그렇지 않으면 false를 반환합니다.
//
// An empty slice returns true (vacuous truth).
// 비어있는 슬라이스는 true를 반환합니다 (공허한 진리).
//
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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

	// Create a map for O(1) lookup / O(1) 조회를 위한 맵 생성
	sliceMap := make(map[T]bool, len(slice))
	for _, item := range slice {
		sliceMap[item] = true
	}

	// Check if all items are in the map / 모든 항목이 맵에 있는지 확인
	for _, item := range items {
		if !sliceMap[item] {
			return false
		}
	}

	return true
}
