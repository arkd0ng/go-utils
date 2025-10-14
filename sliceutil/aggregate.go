package sliceutil

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// aggregate.go contains aggregation operations for slices.
// aggregate.go는 슬라이스 집계 작업을 포함합니다.

// Reduce applies a reducer function to accumulate a single value from the slice.
// Reduce는 슬라이스에서 단일 값을 누적하기 위해 reducer 함수를 적용합니다.
//
// Example / 예제:
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

// Sum returns the sum of all elements in the slice.
// Sum은 슬라이스의 모든 요소의 합을 반환합니다.
//
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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
