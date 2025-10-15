package maputil

// Every checks whether all entries in the map satisfy the predicate.
// Every는 맵의 모든 항목이 조건을 만족하는지 확인합니다.
//
// Returns true for empty maps.
// 빈 맵의 경우 true를 반환합니다.
//
// Time complexity: O(n) worst case, early termination on first false
// 시간 복잡도: 최악의 경우 O(n), 첫 번째 false에서 조기 종료
//
// Example / 예제:
//
//	m := map[string]int{"a": 2, "b": 4, "c": 6}
//	allEven := maputil.Every(m, func(k string, v int) bool {
//	    return v%2 == 0
//	}) // true
func Every[K comparable, V any](m map[K]V, fn func(K, V) bool) bool {
	for k, v := range m {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Some checks whether at least one entry in the map satisfies the predicate.
// Some은 맵의 최소 하나의 항목이 조건을 만족하는지 확인합니다.
//
// Returns false for empty maps.
// 빈 맵의 경우 false를 반환합니다.
//
// Time complexity: O(n) worst case, early termination on first true
// 시간 복잡도: 최악의 경우 O(n), 첫 번째 true에서 조기 종료
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	hasEven := maputil.Some(m, func(k string, v int) bool {
//	    return v%2 == 0
//	}) // true (because of "b": 2)
func Some[K comparable, V any](m map[K]V, fn func(K, V) bool) bool {
	for k, v := range m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// None checks whether no entries in the map satisfy the predicate.
// None은 맵의 어떤 항목도 조건을 만족하지 않는지 확인합니다.
//
// Returns true for empty maps.
// 빈 맵의 경우 true를 반환합니다.
//
// Time complexity: O(n) worst case, early termination on first true
// 시간 복잡도: 최악의 경우 O(n), 첫 번째 true에서 조기 종료
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 3, "c": 5}
//	noEven := maputil.None(m, func(k string, v int) bool {
//	    return v%2 == 0
//	}) // true
func None[K comparable, V any](m map[K]V, fn func(K, V) bool) bool {
	for k, v := range m {
		if fn(k, v) {
			return false
		}
	}
	return true
}

// HasKey checks whether a key exists in the map (alias for Has).
// HasKey는 맵에 키가 존재하는지 확인합니다 (Has의 별칭).
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	exists := maputil.HasKey(m, "a") // true
func HasKey[K comparable, V any](m map[K]V, key K) bool {
	return Has(m, key)
}

// HasValue checks whether a value exists in the map.
// HasValue는 맵에 값이 존재하는지 확인합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	exists := maputil.HasValue(m, 2) // true
//	exists = maputil.HasValue(m, 5)  // false
func HasValue[K comparable, V comparable](m map[K]V, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// HasEntry checks whether a specific key-value pair exists in the map.
// HasEntry는 특정 키-값 쌍이 맵에 존재하는지 확인합니다.
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	exists := maputil.HasEntry(m, "b", 2) // true
//	exists = maputil.HasEntry(m, "b", 3)  // false
func HasEntry[K comparable, V comparable](m map[K]V, key K, value V) bool {
	if v, exists := m[key]; exists {
		return v == value
	}
	return false
}

// IsSubset checks whether the first map is a subset of the second map.
// IsSubset은 첫 번째 맵이 두 번째 맵의 부분집합인지 확인합니다.
//
// A map is a subset if all its key-value pairs exist in the superset.
// 모든 키-값 쌍이 상위집합에 존재하면 부분집합입니다.
//
// Time complexity: O(n) where n is the size of the subset
// 시간 복잡도: O(n) 여기서 n은 부분집합의 크기
//
// Example / 예제:
//
//	subset := map[string]int{"a": 1, "b": 2}
//	superset := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.IsSubset(subset, superset) // true
func IsSubset[K comparable, V comparable](subset, superset map[K]V) bool {
	if len(subset) > len(superset) {
		return false
	}

	for k, v := range subset {
		if superV, exists := superset[k]; !exists || v != superV {
			return false
		}
	}

	return true
}

// IsSuperset checks if a map is a superset of another map.
// IsSuperset는 맵이 다른 맵의 상위집합인지 확인합니다.
//
// A map is a superset if it contains all key-value pairs from the subset.
// 맵이 subset의 모든 키-값 쌍을 포함하면 상위집합입니다.
//
// Time complexity: O(n) where n is the size of subset
// 시간 복잡도: O(n) 여기서 n은 subset의 크기
//
// Example / 예제:
//
//	superset := map[string]int{"a": 1, "b": 2, "c": 3}
//	subset := map[string]int{"a": 1, "b": 2}
//	result := maputil.IsSuperset(superset, subset) // true
func IsSuperset[K comparable, V comparable](superset, subset map[K]V) bool {
	return IsSubset(subset, superset)
}
