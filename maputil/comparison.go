package maputil

// EqualMaps checks whether two maps are equal.
// EqualMaps는 두 맵이 동일한지 확인합니다.
//
// Two maps are equal if they have the same keys and values.
// 두 맵은 같은 키와 값을 가지면 동일합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
//	equal := maputil.EqualMaps(m1, m2) // true
func EqualMaps[K comparable, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		v2, exists := m2[k]
		if !exists || v1 != v2 {
			return false
		}
	}

	return true
}

// EqualFunc checks whether two maps are equal using a custom comparator function.
// EqualFunc은 사용자 정의 비교 함수를 사용하여 두 맵이 동일한지 확인합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m1 := map[string][]int{"a": {1, 2}, "b": {3, 4}}
//	m2 := map[string][]int{"a": {1, 2}, "b": {3, 4}}
//	equal := maputil.EqualFunc(m1, m2, func(v1, v2 []int) bool {
//	    if len(v1) != len(v2) {
//	        return false
//	    }
//	    for i := range v1 {
//	        if v1[i] != v2[i] {
//	            return false
//	        }
//	    }
//	    return true
//	}) // true
func EqualFunc[K comparable, V any](m1, m2 map[K]V, eq func(V, V) bool) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		v2, exists := m2[k]
		if !exists || !eq(v1, v2) {
			return false
		}
	}

	return true
}

// Diff returns a map containing entries that differ between two maps.
// Diff는 두 맵 간에 다른 항목을 포함하는 맵을 반환합니다.
//
// An entry differs if:
// - The key exists in both maps but values are different
// - The key exists in m1 but not in m2
// - The key exists in m2 but not in m1
//
// 다음과 같은 경우 항목이 다릅니다:
// - 키가 두 맵 모두에 존재하지만 값이 다름
// - 키가 m1에는 있지만 m2에는 없음
// - 키가 m2에는 있지만 m1에는 없음
//
// Time complexity: O(n+m)
// 시간 복잡도: O(n+m)
//
// Example
// 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"a": 1, "b": 20, "d": 4}
//	diff := maputil.Diff(m1, m2) // map[string]int{"b": 20, "c": 3, "d": 4}
func Diff[K comparable, V comparable](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V)

	// Check m1 against m2
	// m1과 m2 비교
	for k, v1 := range m1 {
		if v2, exists := m2[k]; !exists || v1 != v2 {
			if exists {
				// Use value from m2
				// m2의 값 사용
				result[k] = v2
			} else {
				// Use value from m1
				// m1의 값 사용
				result[k] = v1
			}
		}
	}

	// Check m2 for keys not in m1
	// m1에 없는 m2의 키 확인
	for k, v2 := range m2 {
		if _, exists := m1[k]; !exists {
			result[k] = v2
		}
	}

	return result
}

// DiffKeys returns a slice of keys that differ between two maps.
// DiffKeys는 두 맵 간에 다른 키를 슬라이스로 반환합니다.
//
// A key differs if:
// - It exists in one map but not the other
// - It exists in both maps but values are different
//
// 다음과 같은 경우 키가 다릅니다:
// - 한 맵에는 존재하지만 다른 맵에는 없음
// - 두 맵 모두에 존재하지만 값이 다름
//
// Time complexity: O(n+m)
// 시간 복잡도: O(n+m)
//
// Example
// 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"a": 1, "b": 20, "d": 4}
//	keys := maputil.DiffKeys(m1, m2) // []string{"b", "c", "d"} (order may vary)
func DiffKeys[K comparable, V comparable](m1, m2 map[K]V) []K {
	keys := make([]K, 0)
	seen := make(map[K]struct{})

	// Check m1 against m2
	// m1과 m2 비교
	for k, v1 := range m1 {
		if v2, exists := m2[k]; !exists || v1 != v2 {
			keys = append(keys, k)
			seen[k] = struct{}{}
		}
	}

	// Check m2 for keys not in m1
	// m1에 없는 m2의 키 확인
	for k := range m2 {
		if _, exists := m1[k]; !exists {
			if _, alreadyAdded := seen[k]; !alreadyAdded {
				keys = append(keys, k)
			}
		}
	}

	return keys
}

// CommonKeys returns a slice of keys that exist in all input maps.
// CommonKeys는 모든 입력 맵에 존재하는 키를 슬라이스로 반환합니다.
//
// Time complexity: O(n*m) where n is the number of maps and m is the average map size
// 시간 복잡도: O(n*m) 여기서 n은 맵 개수, m은 평균 맵 크기
//
// Example
// 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"b": 20, "c": 30, "d": 40}
//	m3 := map[string]int{"c": 100, "d": 200}
//	common := maputil.CommonKeys(m1, m2, m3) // []string{"c"}
func CommonKeys[K comparable, V any](maps ...map[K]V) []K {
	if len(maps) == 0 {
		return []K{}
	}
	if len(maps) == 1 {
		return Keys(maps[0])
	}

	common := make([]K, 0)

	// Iterate over the first map
	// 첫 번째 맵 반복
	for k := range maps[0] {
		existsInAll := true

		// Check if the key exists in all other maps
		// 키가 다른 모든 맵에 존재하는지 확인
		for i := 1; i < len(maps); i++ {
			if _, exists := maps[i][k]; !exists {
				existsInAll = false
				break
			}
		}

		if existsInAll {
			common = append(common, k)
		}
	}

	return common
}

// AllKeys returns a slice of all unique keys from all input maps.
// AllKeys는 모든 입력 맵의 고유한 키를 슬라이스로 반환합니다.
//
// Time complexity: O(n*m)
// 시간 복잡도: O(n*m)
//
// Example
// 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2}
//	m2 := map[string]int{"b": 20, "c": 30}
//	m3 := map[string]int{"c": 100, "d": 200}
//	all := maputil.AllKeys(m1, m2, m3) // []string{"a", "b", "c", "d"} (order may vary)
func AllKeys[K comparable, V any](maps ...map[K]V) []K {
	seen := make(map[K]struct{})
	keys := make([]K, 0)

	for _, m := range maps {
		for k := range m {
			if _, exists := seen[k]; !exists {
				seen[k] = struct{}{}
				keys = append(keys, k)
			}
		}
	}

	return keys
}

// Compare performs a detailed comparison of two maps.
// Compare는 두 맵을 상세하게 비교합니다.
//
// Returns three maps:
// - added: entries that exist in m2 but not in m1
// - removed: entries that exist in m1 but not in m2
// - modified: entries that exist in both maps but with different values (values from m2)
//
// 세 개의 맵을 반환합니다:
// - added: m2에는 있지만 m1에는 없는 항목
// - removed: m1에는 있지만 m2에는 없는 항목
// - modified: 두 맵 모두에 존재하지만 값이 다른 항목 (m2의 값)
//
// Time complexity: O(n+m)
// 시간 복잡도: O(n+m)
//
// Example
// 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"a": 1, "b": 20, "d": 4}
//	added, removed, modified := maputil.Compare(m1, m2)
//	// added = {"d": 4}
//	// removed = {"c": 3}
//	// modified = {"b": 20}
func Compare[K comparable, V comparable](m1, m2 map[K]V) (added, removed, modified map[K]V) {
	added = make(map[K]V)
	removed = make(map[K]V)
	modified = make(map[K]V)

	// Check m1 against m2
	// m1과 m2 비교
	for k, v1 := range m1 {
		if v2, exists := m2[k]; exists {
			if v1 != v2 {
				modified[k] = v2
			}
		} else {
			removed[k] = v1
		}
	}

	// Check m2 for keys not in m1
	// m1에 없는 m2의 키 확인
	for k, v2 := range m2 {
		if _, exists := m1[k]; !exists {
			added[k] = v2
		}
	}

	return added, removed, modified
}
