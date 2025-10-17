package maputil

import "sort"

// ValuesSlice returns all values from the map as a slice (alias for Values).
// ValuesSlice는 맵의 모든 값을 슬라이스로 반환합니다 (Values의 별칭).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
func ValuesSlice[K comparable, V any](m map[K]V) []V {
	return Values(m)
}

// ValuesSorted returns all values from the map as a sorted slice.
// ValuesSorted는 맵의 모든 값을 정렬된 슬라이스로 반환합니다.
//
// Time complexity: O(n log n)
// 시간 복잡도: O(n log n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	values := maputil.ValuesSorted(m) // []int{1, 2, 3}
func ValuesSorted[K comparable, V Ordered](m map[K]V) []V {
	values := Values(m)
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}

// ValuesBy returns values from the map that satisfy the predicate.
// ValuesBy는 조건을 만족하는 맵의 값을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	values := maputil.ValuesBy(m, func(k string, v int) bool {
//	    return v > 2
//	}) // []int{3, 4} (order may vary)
func ValuesBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []V {
	values := make([]V, 0)
	for k, v := range m {
		if fn(k, v) {
			values = append(values, v)
		}
	}
	return values
}

// UniqueValues returns a slice of unique values from the map.
// UniqueValues는 맵의 고유한 값들의 슬라이스를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3, "e": 2}
//	unique := maputil.UniqueValues(m) // []int{1, 2, 3} (order may vary)
func UniqueValues[K comparable, V comparable](m map[K]V) []V {
	seen := make(map[V]struct{})
	result := make([]V, 0)

	for _, v := range m {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// FindValue finds the first value that satisfies the predicate.
// FindValue는 조건을 만족하는 첫 번째 값을 찾습니다.
//
// Returns the value and true if found, or zero value and false otherwise.
// 찾으면 값과 true를 반환하고, 그렇지 않으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n) worst case, early termination when found
// 시간 복잡도: 최악의 경우 O(n), 찾으면 조기 종료
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	value, found := maputil.FindValue(m, func(k string, v int) bool {
//	    return v > 2
//	}) // value = 3, found = true (may vary due to map iteration order)
func FindValue[K comparable, V any](m map[K]V, fn func(K, V) bool) (V, bool) {
	for k, v := range m {
		if fn(k, v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

// ReplaceValue creates a new map with all occurrences of a value replaced.
// ReplaceValue는 특정 값의 모든 발생을 대체한 새 맵을 생성합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3}
//	result := maputil.ReplaceValue(m, 1, 10) // map[string]int{"a": 10, "b": 2, "c": 10, "d": 3}
func ReplaceValue[K comparable, V comparable](m map[K]V, oldValue, newValue V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		if v == oldValue {
			result[k] = newValue
		} else {
			result[k] = v
		}
	}
	return result
}

// UpdateValues creates a new map with all values transformed by the function.
// UpdateValues는 함수로 모든 값을 변환한 새 맵을 생성합니다.
//
// This is similar to MapValues but uses key-value pairs.
// 이것은 MapValues와 유사하지만 키-값 쌍을 사용합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.UpdateValues(m, func(k string, v int) int {
//	    return v * 10
//	}) // map[string]int{"a": 10, "b": 20, "c": 30}
func UpdateValues[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = fn(k, v)
	}
	return result
}

// MinValue returns the minimum value from a map.
// MinValue는 맵에서 최솟값을 반환합니다.
//
// Returns the minimum value and true if found, or zero value and false if map is empty.
// 찾은 경우 최솟값과 true를, 맵이 비어있으면 제로값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	min, found := maputil.MinValue(m) // min = 1, found = true
func MinValue[K comparable, V Ordered](m map[K]V) (V, bool) {
	if len(m) == 0 {
		var zero V
		return zero, false
	}

	var min V
	first := true
	for _, v := range m {
		if first || v < min {
			min = v
			first = false
		}
	}
	return min, true
}

// MaxValue returns the maximum value from a map.
// MaxValue는 맵에서 최댓값을 반환합니다.
//
// Returns the maximum value and true if found, or zero value and false if map is empty.
// 찾은 경우 최댓값과 true를, 맵이 비어있으면 제로값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	max, found := maputil.MaxValue(m) // max = 3, found = true
func MaxValue[K comparable, V Ordered](m map[K]V) (V, bool) {
	if len(m) == 0 {
		var zero V
		return zero, false
	}

	var max V
	first := true
	for _, v := range m {
		if first || v > max {
			max = v
			first = false
		}
	}
	return max, true
}

// SumValues returns the sum of all values in a map.
// SumValues는 맵의 모든 값의 합을 반환합니다.
//
// Returns zero if the map is empty.
// 맵이 비어있으면 0을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	sum := maputil.SumValues(m) // sum = 6
func SumValues[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
