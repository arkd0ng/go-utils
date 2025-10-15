package maputil

// Get retrieves the value for a given key and returns whether the key exists.
// Get은 주어진 키의 값을 가져오고 키의 존재 여부를 반환합니다.
//
// Returns the value and true if the key exists, or the zero value and false otherwise.
// 키가 존재하면 값과 true를 반환하고, 그렇지 않으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	value, ok := maputil.Get(m, "a") // value = 1, ok = true
//	value, ok = maputil.Get(m, "c")  // value = 0, ok = false
func Get[K comparable, V any](m map[K]V, key K) (V, bool) {
	value, ok := m[key]
	return value, ok
}

// GetOr retrieves the value for a given key, or returns a default value if the key doesn't exist.
// GetOr는 주어진 키의 값을 가져오거나, 키가 없으면 기본값을 반환합니다.
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	value := maputil.GetOr(m, "a", 10) // value = 1
//	value = maputil.GetOr(m, "c", 10)  // value = 10
func GetOr[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if value, ok := m[key]; ok {
		return value
	}
	return defaultValue
}

// Set creates a new map with the specified key-value pair added or updated.
// Set은 지정된 키-값 쌍이 추가되거나 업데이트된 새 맵을 생성합니다.
//
// The original map is not modified (immutable operation).
// 원본 맵은 수정되지 않습니다 (불변 작업).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	result := maputil.Set(m, "c", 3) // map[string]int{"a": 1, "b": 2, "c": 3}
//	// Original map m is unchanged / 원본 맵 m은 변경되지 않음
func Set[K comparable, V any](m map[K]V, key K, value V) map[K]V {
	result := make(map[K]V, len(m)+1)
	for k, v := range m {
		result[k] = v
	}
	result[key] = value
	return result
}

// Delete creates a new map with the specified keys removed.
// Delete는 지정된 키가 제거된 새 맵을 생성합니다.
//
// The original map is not modified (immutable operation).
// 원본 맵은 수정되지 않습니다 (불변 작업).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.Delete(m, "b", "c") // map[string]int{"a": 1}
//	// Original map m is unchanged / 원본 맵 m은 변경되지 않음
func Delete[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	if len(keys) == 0 {
		return Clone(m)
	}

	// Create a set of keys to delete for O(1) lookup
	// 삭제할 키의 집합을 생성하여 O(1) 조회
	toDelete := make(map[K]struct{}, len(keys))
	for _, key := range keys {
		toDelete[key] = struct{}{}
	}

	result := make(map[K]V, len(m))
	for k, v := range m {
		if _, shouldDelete := toDelete[k]; !shouldDelete {
			result[k] = v
		}
	}
	return result
}

// Has checks whether a key exists in the map.
// Has는 맵에 키가 존재하는지 확인합니다.
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	exists := maputil.Has(m, "a") // true
//	exists = maputil.Has(m, "c")  // false
func Has[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// IsEmpty checks whether the map is empty.
// IsEmpty는 맵이 비어 있는지 확인합니다.
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example / 예제:
//
//	m := map[string]int{}
//	empty := maputil.IsEmpty(m) // true
//	m["a"] = 1
//	empty = maputil.IsEmpty(m)  // false
func IsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) == 0
}

// IsNotEmpty checks whether the map is not empty.
// IsNotEmpty는 맵이 비어 있지 않은지 확인합니다.
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1}
//	notEmpty := maputil.IsNotEmpty(m) // true
//	m = map[string]int{}
//	notEmpty = maputil.IsNotEmpty(m)  // false
func IsNotEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) > 0
}

// Len returns the number of key-value pairs in the map.
// Len은 맵의 키-값 쌍의 개수를 반환합니다.
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	length := maputil.Len(m) // 3
func Len[K comparable, V any](m map[K]V) int {
	return len(m)
}

// Clear returns an empty map of the same type.
// Clear는 같은 타입의 빈 맵을 반환합니다.
//
// The original map is not modified (immutable operation).
// 원본 맵은 수정되지 않습니다 (불변 작업).
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.Clear(m) // map[string]int{}
//	// Original map m is unchanged / 원본 맵 m은 변경되지 않음
func Clear[K comparable, V any](m map[K]V) map[K]V {
	return make(map[K]V)
}

// Clone creates a shallow copy of the map.
// Clone은 맵의 얕은 복사본을 생성합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	clone := maputil.Clone(m) // map[string]int{"a": 1, "b": 2}
//	clone["c"] = 3
//	// Original map m is unchanged / 원본 맵 m은 변경되지 않음
func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// Equal checks whether two maps are equal (same keys and values).
// Equal은 두 맵이 동일한지 확인합니다 (같은 키와 값).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2}
//	m2 := map[string]int{"a": 1, "b": 2}
//	m3 := map[string]int{"a": 1, "b": 3}
//	equal := maputil.Equal(m1, m2) // true
//	equal = maputil.Equal(m1, m3)  // false
func Equal[K comparable, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}
