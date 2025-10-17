package sliceutil

// basic.go contains basic slice operations for searching and checking.
// basic.go는 검색 및 확인을 위한 기본 슬라이스 작업을 포함합니다.

// Contains checks if slice contains the specified item.
// Contains는 슬라이스에 지정된 항목이 있는지 확인합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasThree := sliceutil.Contains(numbers, 3) // true
//	hasTen := sliceutil.Contains(numbers, 10)  // false
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// ContainsFunc checks if slice contains an item that satisfies the predicate.
// ContainsFunc는 조건을 만족하는 항목이 슬라이스에 있는지 확인합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasEven := sliceutil.ContainsFunc(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // true (has 2, 4)
func ContainsFunc[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of item in slice.
// Returns -1 if item is not found.
//
// IndexOf는 슬라이스에서 항목의 첫 번째 발생 인덱스를 반환합니다.
// 항목을 찾을 수 없으면 -1을 반환합니다.
//
// Example
// 예제:
//
//	fruits := []string{"apple", "banana", "cherry", "banana"}
//	index := sliceutil.IndexOf(fruits, "banana") // 1
//	index2 := sliceutil.IndexOf(fruits, "grape") // -1
func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of item in slice.
// Returns -1 if item is not found.
//
// LastIndexOf는 슬라이스에서 항목의 마지막 발생 인덱스를 반환합니다.
// 항목을 찾을 수 없으면 -1을 반환합니다.
//
// Example
// 예제:
//
//	fruits := []string{"apple", "banana", "cherry", "banana"}
//	index := sliceutil.LastIndexOf(fruits, "banana") // 3
//	index2 := sliceutil.LastIndexOf(fruits, "grape") // -1
func LastIndexOf[T comparable](slice []T, item T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// Find returns the first item in slice that satisfies the predicate.
// Returns the found item and true if found, zero value and false otherwise.
//
// Find는 조건을 만족하는 슬라이스의 첫 번째 항목을 반환합니다.
// 찾은 경우 항목과 true를 반환하고, 그렇지 않으면 제로 값과 false를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	even, found := sliceutil.Find(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // even = 2, found = true
//
//	negative, found := sliceutil.Find(numbers, func(n int) bool {
//	    return n < 0
//	}) // negative = 0, found = false
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, v := range slice {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindLast returns the last item in slice that satisfies the predicate.
// Returns the found item and true if found, zero value and false otherwise.
//
// FindLast는 조건을 만족하는 슬라이스의 마지막 항목을 반환합니다.
// 찾은 경우 항목과 true를 반환하고, 그렇지 않으면 제로 값과 false를 반환합니다.
//
// Similar to Find, but searches from right to left.
// Find와 유사하지만 오른쪽에서 왼쪽으로 검색합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	even, found := sliceutil.FindLast(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // even = 6, found = true (last even number)
//
//	words := []string{"apple", "banana", "apricot", "cherry"}
//	startsWithA, found := sliceutil.FindLast(words, func(s string) bool {
//	    return len(s) > 0 && s[0] == 'a'
//	}) // startsWithA = "apricot", found = true
func FindLast[T any](slice []T, predicate func(T) bool) (T, bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return slice[i], true
		}
	}
	var zero T
	return zero, false
}

// FindIndex returns the index of the first item that satisfies the predicate.
// Returns -1 if no item is found.
//
// FindIndex는 조건을 만족하는 첫 번째 항목의 인덱스를 반환합니다.
// 항목을 찾을 수 없으면 -1을 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	index := sliceutil.FindIndex(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // 1 (index of 2)
func FindIndex[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// Count returns the number of items that satisfy the predicate.
// Count는 조건을 만족하는 항목의 개수를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evenCount := sliceutil.Count(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // 3 (2, 4, 6)
func Count[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
		}
	}
	return count
}

// IsEmpty checks if the slice is empty or nil.
// IsEmpty는 슬라이스가 비어있거나 nil인지 확인합니다.
//
// Example
// 예제:
//
//	empty := []int{}
//	sliceutil.IsEmpty(empty) // true
//	sliceutil.IsEmpty(nil)   // true
//
//	nonEmpty := []int{1, 2, 3}
//	sliceutil.IsEmpty(nonEmpty) // false
func IsEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// IsNotEmpty checks if the slice has at least one item.
// IsNotEmpty는 슬라이스에 최소한 하나의 항목이 있는지 확인합니다.
//
// Example
// 예제:
//
//	empty := []int{}
//	sliceutil.IsNotEmpty(empty) // false
//
//	nonEmpty := []int{1, 2, 3}
//	sliceutil.IsNotEmpty(nonEmpty) // true
func IsNotEmpty[T any](slice []T) bool {
	return len(slice) > 0
}

// Equal checks if two slices are equal (same length and same elements in same order).
// Equal은 두 슬라이스가 같은지 확인합니다 (같은 길이와 같은 순서의 같은 요소).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3}
//	b := []int{1, 2, 3}
//	c := []int{1, 2, 4}
//
//	sliceutil.Equal(a, b) // true
//	sliceutil.Equal(a, c) // false
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
