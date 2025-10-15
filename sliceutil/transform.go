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

// Unique returns a new slice with duplicate elements removed.
// Unique는 중복 요소가 제거된 새 슬라이스를 반환합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4}
//	unique := sliceutil.Unique(numbers) // [1, 2, 3, 4]
//
//	words := []string{"apple", "banana", "apple", "cherry"}
//	unique := sliceutil.Unique(words) // ["apple", "banana", "cherry"]
func Unique[T comparable](slice []T) []T {
	if slice == nil {
		return nil
	}
	seen := make(map[T]bool, len(slice))
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// UniqueBy returns a new slice with duplicate elements removed based on a key function.
// UniqueBy는 키 함수를 기반으로 중복 요소가 제거된 새 슬라이스를 반환합니다.
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
//	    {"Alice", 28},
//	}
//	unique := sliceutil.UniqueBy(people, func(p Person) string {
//	    return p.Name
//	}) // [{"Alice", 25}, {"Bob", 30}]
func UniqueBy[T any, K comparable](slice []T, keyFunc func(T) K) []T {
	if slice == nil {
		return nil
	}
	seen := make(map[K]bool, len(slice))
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		key := keyFunc(v)
		if !seen[key] {
			seen[key] = true
			result = append(result, v)
		}
	}
	return result
}

// Compact removes consecutive duplicate elements from the slice.
// Compact는 슬라이스에서 연속된 중복 요소를 제거합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
//	compact := sliceutil.Compact(numbers) // [1, 2, 3, 4, 5]
//
//	words := []string{"a", "a", "b", "b", "c"}
//	compact := sliceutil.Compact(words) // ["a", "b", "c"]
func Compact[T comparable](slice []T) []T {
	if slice == nil {
		return nil
	}
	if len(slice) == 0 {
		return []T{}
	}
	result := make([]T, 0, len(slice))
	result = append(result, slice[0])
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			result = append(result, slice[i])
		}
	}
	return result
}

// Reverse returns a new slice with elements in reverse order.
// Reverse는 요소가 역순으로 된 새 슬라이스를 반환합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	reversed := sliceutil.Reverse(numbers) // [5, 4, 3, 2, 1]
//
//	words := []string{"hello", "world"}
//	reversed := sliceutil.Reverse(words) // ["world", "hello"]
func Reverse[T any](slice []T) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}
