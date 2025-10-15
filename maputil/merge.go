package maputil

// Merge merges multiple maps into a single map.
// Merge는 여러 맵을 단일 맵으로 병합합니다.
//
// If duplicate keys exist, the value from the last map wins.
// 중복 키가 있으면 마지막 맵의 값이 우선합니다.
//
// Time complexity: O(n*m) where n is the number of maps and m is the average map size
// 시간 복잡도: O(n*m) 여기서 n은 맵 개수, m은 평균 맵 크기
//
// Example / 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2}
//	m2 := map[string]int{"b": 3, "c": 4}
//	m3 := map[string]int{"c": 5, "d": 6}
//	result := maputil.Merge(m1, m2, m3) // map[string]int{"a": 1, "b": 3, "c": 5, "d": 6}
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	totalSize := 0
	for _, m := range maps {
		totalSize += len(m)
	}

	result := make(map[K]V, totalSize)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// MergeWith merges multiple maps using a custom resolver function for duplicate keys.
// MergeWith는 중복 키에 대해 사용자 정의 해결 함수를 사용하여 여러 맵을 병합합니다.
//
// The resolver function receives the old and new values and returns the value to use.
// 해결 함수는 이전 값과 새 값을 받아 사용할 값을 반환합니다.
//
// Time complexity: O(n*m) where n is the number of maps and m is the average map size
// 시간 복잡도: O(n*m) 여기서 n은 맵 개수, m은 평균 맵 크기
//
// Example / 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2}
//	m2 := map[string]int{"b": 3, "c": 4}
//	result := maputil.MergeWith(
//	    func(old, new int) int { return old + new }, // Sum on conflict / 충돌 시 합산
//	    m1, m2,
//	) // map[string]int{"a": 1, "b": 5, "c": 4}
func MergeWith[K comparable, V any](fn func(V, V) V, maps ...map[K]V) map[K]V {
	totalSize := 0
	for _, m := range maps {
		totalSize += len(m)
	}

	result := make(map[K]V, totalSize)
	for _, m := range maps {
		for k, v := range m {
			if existing, exists := result[k]; exists {
				result[k] = fn(existing, v)
			} else {
				result[k] = v
			}
		}
	}

	return result
}

// DeepMerge performs a deep merge of nested string maps.
// DeepMerge는 중첩된 문자열 맵을 깊이 병합합니다.
//
// Nested maps are recursively merged. For non-map values, the last value wins.
// 중첩된 맵은 재귀적으로 병합됩니다. 맵이 아닌 값의 경우 마지막 값이 우선합니다.
//
// Time complexity: O(n*m*d) where d is the average nesting depth
// 시간 복잡도: O(n*m*d) 여기서 d는 평균 중첩 깊이
//
// Example / 예제:
//
//	m1 := map[string]interface{}{
//	    "user": map[string]interface{}{"name": "Alice", "age": 25},
//	}
//	m2 := map[string]interface{}{
//	    "user": map[string]interface{}{"age": 26, "city": "Seoul"},
//	}
//	result := maputil.DeepMerge(m1, m2)
//	// result = {"user": {"name": "Alice", "age": 26, "city": "Seoul"}}
func DeepMerge(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, m := range maps {
		for k, v := range m {
			if existing, exists := result[k]; exists {
				// If both values are maps, merge them recursively
				// 두 값이 모두 맵이면 재귀적으로 병합
				if existingMap, ok := existing.(map[string]interface{}); ok {
					if vMap, ok := v.(map[string]interface{}); ok {
						result[k] = DeepMerge(existingMap, vMap)
						continue
					}
				}
			}
			result[k] = v
		}
	}

	return result
}

// Union is an alias for Merge.
// Union은 Merge의 별칭입니다.
//
// Time complexity: O(n*m)
// 시간 복잡도: O(n*m)
func Union[K comparable, V any](maps ...map[K]V) map[K]V {
	return Merge(maps...)
}

// Intersection returns a map containing only keys that exist in all input maps.
// Intersection은 모든 입력 맵에 존재하는 키만 포함하는 맵을 반환합니다.
//
// The values are taken from the first map.
// 값은 첫 번째 맵에서 가져옵니다.
//
// Time complexity: O(n*m) where n is the number of maps and m is the average map size
// 시간 복잡도: O(n*m) 여기서 n은 맵 개수, m은 평균 맵 크기
//
// Example / 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"b": 20, "c": 30, "d": 40}
//	m3 := map[string]int{"c": 100, "d": 200}
//	result := maputil.Intersection(m1, m2, m3) // map[string]int{"c": 3}
func Intersection[K comparable, V any](maps ...map[K]V) map[K]V {
	if len(maps) == 0 {
		return make(map[K]V)
	}
	if len(maps) == 1 {
		return Clone(maps[0])
	}

	result := make(map[K]V)

	// Iterate over the first map
	// 첫 번째 맵 반복
	for k, v := range maps[0] {
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
			result[k] = v
		}
	}

	return result
}

// Difference returns a map containing keys from the first map that are not in the second map.
// Difference는 첫 번째 맵에는 있지만 두 번째 맵에는 없는 키를 포함하는 맵을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"b": 20, "d": 40}
//	result := maputil.Difference(m1, m2) // map[string]int{"a": 1, "c": 3}
func Difference[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V)

	for k, v := range m1 {
		if _, exists := m2[k]; !exists {
			result[k] = v
		}
	}

	return result
}

// SymmetricDifference returns a map containing keys that exist in either map but not in both.
// SymmetricDifference는 어느 한 맵에는 있지만 두 맵 모두에는 없는 키를 포함하는 맵을 반환합니다.
//
// Time complexity: O(n+m)
// 시간 복잡도: O(n+m)
//
// Example / 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	m2 := map[string]int{"b": 20, "c": 30, "d": 40}
//	result := maputil.SymmetricDifference(m1, m2) // map[string]int{"a": 1, "d": 40}
func SymmetricDifference[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V)

	// Add keys from m1 that are not in m2
	// m1에는 있지만 m2에는 없는 키 추가
	for k, v := range m1 {
		if _, exists := m2[k]; !exists {
			result[k] = v
		}
	}

	// Add keys from m2 that are not in m1
	// m2에는 있지만 m1에는 없는 키 추가
	for k, v := range m2 {
		if _, exists := m1[k]; !exists {
			result[k] = v
		}
	}

	return result
}

// Assign merges source maps into the target map (mutating operation).
// Assign은 소스 맵을 대상 맵에 병합합니다 (변경 작업).
//
// WARNING: This function modifies the target map in place.
// 경고: 이 함수는 대상 맵을 직접 수정합니다.
//
// Time complexity: O(n*m)
// 시간 복잡도: O(n*m)
//
// Example / 예제:
//
//	target := map[string]int{"a": 1, "b": 2}
//	source := map[string]int{"b": 3, "c": 4}
//	result := maputil.Assign(target, source) // target is modified / target이 수정됨
//	// result = map[string]int{"a": 1, "b": 3, "c": 4}
func Assign[K comparable, V any](target map[K]V, sources ...map[K]V) map[K]V {
	for _, source := range sources {
		for k, v := range source {
			target[k] = v
		}
	}
	return target
}
