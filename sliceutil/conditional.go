package sliceutil

// conditional.go contains conditional slice operations.
// conditional.go는 조건부 슬라이스 작업을 포함합니다.

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
