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
