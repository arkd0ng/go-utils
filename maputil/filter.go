package maputil

// Filter returns a new map containing only the entries that satisfy the predicate.
// Filter는 조건을 만족하는 항목만 포함하는 새 맵을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	result := maputil.Filter(m, func(k string, v int) bool {
//	    return v > 2
//	}) // map[string]int{"c": 3, "d": 4}
func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if fn(k, v) {
			result[k] = v
		}
	}

	return result
}

// FilterKeys returns a new map containing only the entries with keys that satisfy the predicate.
// FilterKeys는 조건을 만족하는 키를 가진 항목만 포함하는 새 맵을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
//	result := maputil.FilterKeys(m, func(k string) bool {
//	    return len(k) > 5
//	}) // map[string]int{"banana": 2, "cherry": 3}
func FilterKeys[K comparable, V any](m map[K]V, fn func(K) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if fn(k) {
			result[k] = v
		}
	}

	return result
}

// FilterValues returns a new map containing only the entries with values that satisfy the predicate.
// FilterValues는 조건을 만족하는 값을 가진 항목만 포함하는 새 맵을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	result := maputil.FilterValues(m, func(v int) bool {
//	    return v%2 == 0
//	}) // map[string]int{"b": 2, "d": 4}
func FilterValues[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if fn(v) {
			result[k] = v
		}
	}

	return result
}

// Omit returns a new map excluding the specified keys.
// Omit은 지정된 키를 제외한 새 맵을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	result := maputil.Omit(m, "b", "d") // map[string]int{"a": 1, "c": 3}
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	if len(keys) == 0 {
		return Clone(m)
	}

	// Create a set of keys to omit for O(1) lookup
	// 생략할 키의 집합을 생성하여 O(1) 조회
	toOmit := make(map[K]struct{}, len(keys))
	for _, key := range keys {
		toOmit[key] = struct{}{}
	}

	result := make(map[K]V)
	for k, v := range m {
		if _, shouldOmit := toOmit[k]; !shouldOmit {
			result[k] = v
		}
	}

	return result
}

// Pick returns a new map containing only the specified keys.
// Pick은 지정된 키만 포함하는 새 맵을 반환합니다.
//
// Keys that don't exist in the original map are ignored.
// 원본 맵에 존재하지 않는 키는 무시됩니다.
//
// Time complexity: O(k) where k is the number of keys to pick
// 시간 복잡도: O(k) 여기서 k는 선택할 키의 개수
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	result := maputil.Pick(m, "a", "c", "e") // map[string]int{"a": 1, "c": 3}
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	result := make(map[K]V, len(keys))

	for _, key := range keys {
		if value, exists := m[key]; exists {
			result[key] = value
		}
	}

	return result
}

// OmitBy returns a new map excluding entries that satisfy the predicate.
// OmitBy는 조건을 만족하는 항목을 제외한 새 맵을 반환합니다.
//
// This is the inverse of Filter.
// 이것은 Filter의 반대입니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	result := maputil.OmitBy(m, func(k string, v int) bool {
//	    return v%2 == 0
//	}) // map[string]int{"a": 1, "c": 3}
func OmitBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if !fn(k, v) {
			result[k] = v
		}
	}

	return result
}

// PickBy returns a new map containing only entries that satisfy the predicate.
// PickBy는 조건을 만족하는 항목만 포함하는 새 맵을 반환합니다.
//
// This is an alias for Filter.
// 이것은 Filter의 별칭입니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	result := maputil.PickBy(m, func(k string, v int) bool {
//	    return v > 2
//	}) // map[string]int{"c": 3, "d": 4}
func PickBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	return Filter(m, fn)
}
