package maputil

import "sort"

// KeysSlice returns all keys from the map as a slice (alias for Keys).
// KeysSlice는 맵의 모든 키를 슬라이스로 반환합니다 (Keys의 별칭).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
func KeysSlice[K comparable, V any](m map[K]V) []K {
	return Keys(m)
}

// KeysSorted returns all keys from the map as a sorted slice.
// KeysSorted는 맵의 모든 키를 정렬된 슬라이스로 반환합니다.
//
// Time complexity: O(n log n)
// 시간 복잡도: O(n log n)
//
// Example / 예제:
//
//	m := map[string]int{"c": 3, "a": 1, "b": 2}
//	keys := maputil.KeysSorted(m) // []string{"a", "b", "c"}
func KeysSorted[K Ordered, V any](m map[K]V) []K {
	keys := Keys(m)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

// KeysBy returns keys from the map that satisfy the predicate.
// KeysBy는 조건을 만족하는 맵의 키를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	keys := maputil.KeysBy(m, func(k string, v int) bool {
//	    return v > 2
//	}) // []string{"c", "d"} (order may vary)
func KeysBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []K {
	keys := make([]K, 0)
	for k, v := range m {
		if fn(k, v) {
			keys = append(keys, k)
		}
	}
	return keys
}

// RenameKey creates a new map with a key renamed.
// RenameKey는 키의 이름이 변경된 새 맵을 생성합니다.
//
// If the old key doesn't exist, returns a clone of the original map.
// If the new key already exists, it will be overwritten.
//
// 이전 키가 존재하지 않으면 원본 맵의 복사본을 반환합니다.
// 새 키가 이미 존재하면 덮어씁니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.RenameKey(m, "b", "B") // map[string]int{"a": 1, "B": 2, "c": 3}
func RenameKey[K comparable, V any](m map[K]V, oldKey, newKey K) map[K]V {
	result := Clone(m)

	if value, exists := result[oldKey]; exists {
		delete(result, oldKey)
		result[newKey] = value
	}

	return result
}

// RenameKeys creates a new map with multiple keys renamed according to the mapping.
// RenameKeys는 매핑에 따라 여러 키의 이름이 변경된 새 맵을 생성합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	mapping := map[string]string{"a": "A", "b": "B"}
//	result := maputil.RenameKeys(m, mapping) // map[string]int{"A": 1, "B": 2, "c": 3}
func RenameKeys[K comparable, V any](m map[K]V, mapping map[K]K) map[K]V {
	result := make(map[K]V, len(m))

	for k, v := range m {
		if newKey, shouldRename := mapping[k]; shouldRename {
			result[newKey] = v
		} else {
			result[k] = v
		}
	}

	return result
}

// SwapKeys creates a new map with two keys swapped.
// SwapKeys는 두 키가 교환된 새 맵을 생성합니다.
//
// If either key doesn't exist, returns a clone of the original map.
// 어느 한 키라도 존재하지 않으면 원본 맵의 복사본을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.SwapKeys(m, "a", "b") // map[string]int{"a": 2, "b": 1, "c": 3}
func SwapKeys[K comparable, V any](m map[K]V, key1, key2 K) map[K]V {
	result := Clone(m)

	val1, exists1 := result[key1]
	val2, exists2 := result[key2]

	if exists1 && exists2 {
		result[key1] = val2
		result[key2] = val1
	}

	return result
}

// FindKey finds the first key that satisfies the predicate.
// FindKey는 조건을 만족하는 첫 번째 키를 찾습니다.
//
// Returns the key and true if found, or zero value and false otherwise.
// 찾으면 키와 true를 반환하고, 그렇지 않으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n) worst case, early termination when found
// 시간 복잡도: 최악의 경우 O(n), 찾으면 조기 종료
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	key, found := maputil.FindKey(m, func(k string, v int) bool {
//	    return v > 2
//	}) // key = "c", found = true (may vary due to map iteration order)
func FindKey[K comparable, V any](m map[K]V, fn func(K, V) bool) (K, bool) {
	for k, v := range m {
		if fn(k, v) {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// FindKeys finds all keys that satisfy the predicate.
// FindKeys는 조건을 만족하는 모든 키를 찾습니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	keys := maputil.FindKeys(m, func(k string, v int) bool {
//	    return v > 2
//	}) // []string{"c", "d"} (order may vary)
func FindKeys[K comparable, V any](m map[K]V, fn func(K, V) bool) []K {
	return KeysBy(m, fn)
}
