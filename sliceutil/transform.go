package sliceutil

// transform.go contains transformation operations for slices.
// transform.go는 슬라이스 변환 작업을 포함합니다.

// Map applies a function to each element and returns a new slice with the results.
// Map은 각 요소에 함수를 적용하고 결과로 새 슬라이스를 반환합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	doubled := sliceutil.Map(numbers, func(n int) int {
//	    return n * 2
//	}) // [2, 4, 6, 8, 10]
//
//	words := []string{"hello", "world"}
//	lengths := sliceutil.Map(words, func(s string) int {
//	    return len(s)
//	}) // [5, 5]
func Map[T any, R any](slice []T, fn func(T) R) []R {
	if slice == nil {
		return nil
	}
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter returns a new slice containing only elements that satisfy the predicate.
// Filter는 조건을 만족하는 요소만 포함하는 새 슬라이스를 반환합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens := sliceutil.Filter(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // [2, 4, 6]
//
//	words := []string{"apple", "a", "banana", "pear", "ab"}
//	long := sliceutil.Filter(words, func(s string) bool {
//	    return len(s) > 2
//	}) // ["apple", "banana", "pear"]
func Filter[T any](slice []T, predicate func(T) bool) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// FlatMap applies a function to each element and flattens the results into a single slice.
// FlatMap은 각 요소에 함수를 적용하고 결과를 하나의 슬라이스로 평탄화합니다.
//
// Example / 예제:
//
//	words := []string{"hello", "world"}
//	chars := sliceutil.FlatMap(words, func(s string) []rune {
//	    return []rune(s)
//	}) // ['h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd']
//
//	numbers := []int{1, 2, 3}
//	pairs := sliceutil.FlatMap(numbers, func(n int) []int {
//	    return []int{n, n * 2}
//	}) // [1, 2, 2, 4, 3, 6]
func FlatMap[T any, R any](slice []T, fn func(T) []R) []R {
	if slice == nil {
		return nil
	}
	result := make([]R, 0, len(slice))
	for _, v := range slice {
		result = append(result, fn(v)...)
	}
	return result
}

// Flatten flattens a slice of slices into a single slice.
// Flatten은 슬라이스의 슬라이스를 하나의 슬라이스로 평탄화합니다.
//
// Example / 예제:
//
//	nested := [][]int{{1, 2}, {3, 4}, {5}}
//	flat := sliceutil.Flatten(nested) // [1, 2, 3, 4, 5]
//
//	words := [][]string{{"hello", "world"}, {"foo", "bar"}}
//	flat := sliceutil.Flatten(words) // ["hello", "world", "foo", "bar"]
func Flatten[T any](slice [][]T) []T {
	if slice == nil {
		return nil
	}
	// Calculate total length for pre-allocation
	totalLen := 0
	for _, sub := range slice {
		totalLen += len(sub)
	}
	result := make([]T, 0, totalLen)
	for _, sub := range slice {
		result = append(result, sub...)
	}
	return result
}
