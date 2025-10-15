package maputil

// ForEach executes a function for each key-value pair in the map.
// ForEach는 맵의 각 키-값 쌍에 대해 함수를 실행합니다.
//
// This is a side-effect operation that does not return a value.
// It is useful for logging, debugging, or performing operations
// that don't require creating a new map.
//
// 이것은 값을 반환하지 않는 부수 효과 작업입니다.
// 로깅, 디버깅 또는 새 맵을 생성할 필요가 없는 작업에 유용합니다.
//
// Time Complexity / 시간 복잡도: O(n)
// Space Complexity / 공간 복잡도: O(1)
//
// Parameters / 매개변수:
//   - m: The input map / 입력 맵
//   - fn: Function to execute for each entry / 각 항목에 대해 실행할 함수
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	maputil.ForEach(m, func(k string, v int) {
//	    fmt.Printf("%s=%d\n", k, v)
//	})
//	// Output (order may vary):
//	// a=1
//	// b=2
//	// c=3
//
// Use Case / 사용 사례:
//   - Logging all map entries / 모든 맵 항목 로깅
//   - Debugging map contents / 맵 내용 디버깅
//   - Performing side effects (e.g., sending notifications) / 부수 효과 수행 (예: 알림 전송)
//   - Collecting statistics without modifying the map / 맵을 수정하지 않고 통계 수집
func ForEach[K comparable, V any](m map[K]V, fn func(K, V)) {
	for k, v := range m {
		fn(k, v)
	}
}

// GetMany retrieves multiple values from the map by their keys.
// GetMany는 키로 맵에서 여러 값을 검색합니다.
//
// This function returns a slice of values corresponding to the provided keys.
// If a key does not exist in the map, the zero value for V is returned at that position.
//
// 이 함수는 제공된 키에 해당하는 값의 슬라이스를 반환합니다.
// 키가 맵에 존재하지 않으면 해당 위치에 V의 제로 값이 반환됩니다.
//
// Time Complexity / 시간 복잡도: O(k) where k is the number of keys / k는 키의 개수
// Space Complexity / 공간 복잡도: O(k)
//
// Parameters / 매개변수:
//   - m: The input map / 입력 맵
//   - keys: Variable number of keys to retrieve / 검색할 키의 가변 개수
//
// Returns / 반환값:
//   - []V: Slice of values corresponding to the keys / 키에 해당하는 값의 슬라이스
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	values := maputil.GetMany(m, "a", "c", "d")
//	// values: [1, 3, 0] (d doesn't exist, returns zero value)
//
// Use Case / 사용 사례:
//   - Batch retrieval of multiple values / 여러 값의 일괄 검색
//   - Configuration lookups / 설정 조회
//   - Data extraction for processing / 처리를 위한 데이터 추출
func GetMany[K comparable, V any](m map[K]V, keys ...K) []V {
	result := make([]V, len(keys))
	for i, key := range keys {
		result[i] = m[key]
	}
	return result
}

// SetMany sets multiple key-value pairs in the map at once.
// SetMany는 맵에 여러 키-값 쌍을 한 번에 설정합니다.
//
// This function creates a new map with the original entries plus the new entries.
// If a key already exists, its value is updated.
//
// 이 함수는 원본 항목과 새 항목을 포함하는 새 맵을 생성합니다.
// 키가 이미 존재하면 값이 업데이트됩니다.
//
// Time Complexity / 시간 복잡도: O(n + e) where n is map size, e is entries / n은 맵 크기, e는 항목 개수
// Space Complexity / 공간 복잡도: O(n + e)
//
// Parameters / 매개변수:
//   - m: The input map / 입력 맵
//   - entries: Variable number of Entry structs to set / 설정할 Entry 구조체의 가변 개수
//
// Returns / 반환값:
//   - map[K]V: New map with updated entries / 업데이트된 항목이 있는 새 맵
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	result := maputil.SetMany(m,
//	    maputil.Entry[string, int]{Key: "c", Value: 3},
//	    maputil.Entry[string, int]{Key: "d", Value: 4},
//	)
//	// result: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//
// Use Case / 사용 사례:
//   - Batch updates to configuration / 설정에 대한 일괄 업데이트
//   - Initializing map with multiple values / 여러 값으로 맵 초기화
//   - Merging multiple entries at once / 여러 항목을 한 번에 병합
func SetMany[K comparable, V any](m map[K]V, entries ...Entry[K, V]) map[K]V {
	result := make(map[K]V, len(m)+len(entries))
	for k, v := range m {
		result[k] = v
	}
	for _, entry := range entries {
		result[entry.Key] = entry.Value
	}
	return result
}

// Tap executes a side-effect function on the map and returns the original map.
// Tap은 맵에 대해 부수 효과 함수를 실행하고 원본 맵을 반환합니다.
//
// This function is useful for method chaining where you want to perform a side effect
// (like logging or debugging) without breaking the chain. The function receives the
// entire map and can inspect or operate on it, but the original map is returned unchanged.
//
// 이 함수는 체인을 끊지 않고 부수 효과(로깅 또는 디버깅 등)를 수행하려는
// 메서드 체이닝에 유용합니다. 함수는 전체 맵을 받아 검사하거나 작업할 수 있지만,
// 원본 맵은 변경되지 않고 반환됩니다.
//
// Time Complexity / 시간 복잡도: O(n) - depends on fn / fn에 따라 다름
// Space Complexity / 공간 복잡도: O(1)
//
// Parameters / 매개변수:
//   - m: The input map / 입력 맵
//   - fn: Side-effect function to execute / 실행할 부수 효과 함수
//
// Returns / 반환값:
//   - map[K]V: The original map (unchanged) / 원본 맵 (변경되지 않음)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.Filter(m, func(k string, v int) bool {
//	    return v > 1
//	}).
//	Tap(func(m map[string]int) {
//	    fmt.Printf("Filtered map: %v\n", m)
//	}).
//	Map(func(k string, v int) int {
//	    return v * 2
//	})
//	// Logs: Filtered map: map[b:2 c:3]
//	// result: map[string]int{"b": 4, "c": 6}
//
// Use Case / 사용 사례:
//   - Debugging in method chains / 메서드 체인에서 디버깅
//   - Logging intermediate results / 중간 결과 로깅
//   - Collecting statistics without breaking chain / 체인을 끊지 않고 통계 수집
//   - Performing validation or assertions / 유효성 검사 또는 단언 수행
func Tap[K comparable, V any](m map[K]V, fn func(map[K]V)) map[K]V {
	fn(m)
	return m
}
