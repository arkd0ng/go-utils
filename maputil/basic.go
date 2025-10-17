package maputil

// basic.go provides fundamental map operations for Go.
//
// This file implements core map manipulation functions that form the foundation
// of the maputil package. All operations are type-safe using Go generics and
// follow immutable design patterns (original maps remain unchanged).
//
// Core Operations:
//
// Key-Value Access:
//   - Get(m, key): Retrieve value with existence check
//     Time: O(1), Space: O(1)
//     Returns (value, true) if key exists, (zero, false) otherwise
//     Example: Get({"a": 1}, "a") = (1, true)
//     Use cases: Safe key access, conditional logic, default handling
//
//   - GetOr(m, key, default): Get value with fallback
//     Time: O(1), Space: O(1)
//     Returns value if key exists, default otherwise
//     Example: GetOr({"a": 1}, "b", 10) = 10
//     Use cases: Configuration reading, safe defaults, missing key handling
//
// Map Modification:
//   - Set(m, key, value): Add or update key-value pair
//     Time: O(n), Space: O(n)
//     Creates new map with added/updated pair
//     Immutable: Original map unchanged
//     Example: Set({"a": 1}, "b", 2) = {"a": 1, "b": 2}
//     Use cases: Immutable updates, functional pipelines, state management
//
//   - Delete(m, keys...): Remove multiple keys
//     Time: O(n + k) where k = len(keys), Space: O(n)
//     Creates new map without specified keys
//     Immutable: Original map unchanged
//     Uses set for O(1) key lookup
//     Example: Delete({"a": 1, "b": 2, "c": 3}, "b", "c") = {"a": 1}
//     Use cases: Filtering, cleanup, selective removal
//
// Existence Checks:
//   - Has(m, key): Check if key exists
//     Time: O(1), Space: O(1)
//     Boolean existence check
//     Example: Has({"a": 1}, "a") = true
//     Use cases: Validation, conditional logic, pre-checks
//
//   - IsEmpty(m): Check if map has no entries
//     Time: O(1), Space: O(1)
//     Returns true if len(m) == 0
//     Example: IsEmpty({}) = true
//     Use cases: Validation, early returns, error checks
//
//   - IsNotEmpty(m): Check if map has entries
//     Time: O(1), Space: O(1)
//     Returns true if len(m) > 0
//     Example: IsNotEmpty({"a": 1}) = true
//     Use cases: Guard clauses, validation, positive checks
//
// Map Properties:
//   - Len(m): Get number of entries
//     Time: O(1), Space: O(1)
//     Returns count of key-value pairs
//     Example: Len({"a": 1, "b": 2}) = 2
//     Use cases: Size validation, capacity planning, statistics
//
//   - Clear(m): Create empty map
//     Time: O(1), Space: O(1)
//     Returns new empty map of same type
//     Immutable: Original map unchanged
//     Example: Clear({"a": 1, "b": 2}) = {}
//     Use cases: Reset state, initialization, cleanup
//
// Copying & Comparison:
//   - Clone(m): Create shallow copy
//     Time: O(n), Space: O(n)
//     Duplicates map structure
//     Nil-safe: Returns nil for nil input
//     Shallow: Copies references, not deep values
//     Example: Clone({"a": 1, "b": 2}) = {"a": 1, "b": 2} (independent)
//     Use cases: Defensive copying, immutable patterns, safe mutation
//
//   - Equal(m1, m2): Compare two maps for equality
//     Time: O(n), Space: O(1)
//     Checks same keys and values
//     Requires V to be comparable
//     Example: Equal({"a": 1}, {"a": 1}) = true
//     Use cases: Testing, validation, change detection
//
// Design Principles:
//   - Immutability: All modification operations return new maps
//   - Type Safety: Leverages Go generics for compile-time safety
//   - Nil Safety: Handles nil maps gracefully (Clone preserves nil)
//   - Performance: O(1) operations for access, O(n) for modifications
//   - Consistency: Predictable behavior across all operations
//
// Comparison: Set vs Native Assignment:
//   - Set: Returns new map, immutable, functional style
//   - Native (m[k] = v): Mutates original, imperative style
//   - Set: Suitable for concurrent access patterns
//   - Native: More efficient when mutation is safe
//
// Comparison: Delete vs Native delete():
//   - Delete: Returns new map, immutable, multi-key support
//   - Native delete(): Mutates original, single key
//   - Delete: Suitable for functional pipelines
//   - Native: More efficient for single deletions
//
// Comparison: Clone vs Manual Copy:
//   - Clone: One-line, nil-safe, clear intent
//   - Manual: Requires loop, manual nil checks
//   - Clone: Consistent behavior across codebase
//   - Manual: More verbose, error-prone
//
// Performance Characteristics:
//
// Time Complexity:
//   - Get/GetOr: O(1) - Hash table lookup
//   - Has/IsEmpty/IsNotEmpty/Len: O(1) - Direct checks
//   - Clear: O(1) - Allocates new empty map
//   - Set: O(n) - Copies all existing entries + new one
//   - Delete: O(n + k) - Copies all entries except k keys
//   - Clone: O(n) - Copies all entries
//   - Equal: O(n) - Compares all entries
//
// Space Complexity:
//   - Get/GetOr/Has/IsEmpty/IsNotEmpty/Len: O(1) - No allocation
//   - Clear: O(1) - Empty map
//   - Set/Delete/Clone: O(n) - New map with n entries
//   - Equal: O(1) - No allocation (comparison only)
//
// Memory Allocation:
//   - Immutable operations (Set, Delete, Clone) allocate new maps
//   - Read operations (Get, Has, Len) allocate nothing
//   - For high-frequency updates, consider mutable operations
//   - Clone is shallow: nested structures share references
//
// Common Usage Patterns:
//
//	// Safe key access with default
//	config := map[string]string{"host": "localhost"}
//	port := maputil.GetOr(config, "port", "8080")
//
//	// Immutable updates in pipelines
//	m := map[string]int{"a": 1, "b": 2}
//	m = maputil.Set(m, "c", 3)
//	m = maputil.Set(m, "d", 4)
//	// Original {"a": 1, "b": 2} unchanged at each step
//
//	// Defensive copying
//	original := map[string]int{"x": 1, "y": 2}
//	working := maputil.Clone(original)
//	working["z"] = 3 // original unchanged
//
//	// Bulk key removal
//	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	filtered := maputil.Delete(data, "b", "d")
//	// Result: {"a": 1, "c": 3}
//
//	// Validation chains
//	if maputil.IsNotEmpty(cache) && maputil.Has(cache, key) {
//	    value := cache[key]
//	    // Process value
//	}
//
//	// Equality testing
//	expected := map[string]int{"a": 1, "b": 2}
//	actual := processData()
//	if !maputil.Equal(expected, actual) {
//	    t.Errorf("mismatch: got %v, want %v", actual, expected)
//	}
//
// Immutability Patterns:
//
//	// Pattern 1: Functional updates
//	state := map[string]int{"count": 0}
//	state = maputil.Set(state, "count", 1)
//	state = maputil.Set(state, "max", 10)
//
//	// Pattern 2: Conditional updates
//	var newState map[string]int
//	if shouldUpdate {
//	    newState = maputil.Set(state, key, value)
//	} else {
//	    newState = state
//	}
//
//	// Pattern 3: Pipelined transformations
//	result := maputil.Set(
//	    maputil.Delete(original, "temp"),
//	    "final", computeValue(),
//	)
//
// Edge Cases:
//   - Get on nil map: Returns (zero, false) - no panic
//   - Set on nil map: Creates new map with one entry
//   - Delete with no keys: Returns clone of original
//   - Delete with non-existent keys: Ignores missing keys
//   - Clone of nil: Returns nil (preserves nil semantics)
//   - Equal with different lengths: Returns false immediately (optimization)
//   - IsEmpty on nil: Returns true (len(nil) == 0)
//   - Has on nil map: Returns false - no panic
//
// Thread Safety:
//   - Read operations (Get, Has, Len, etc.): Safe if no concurrent writes
//   - Write operations: Create new maps, safe for different maps
//   - Same map: Not thread-safe without synchronization
//   - Clone: Safe way to share map across goroutines
//   - Immutable pattern: Naturally safer for concurrency
//
// Performance Tips:
//   - Use GetOr instead of Get + if-else for defaults
//   - Batch multiple Set operations if possible
//   - Consider native mutations for high-frequency updates
//   - Clone only when necessary (copy cost is O(n))
//   - Use Has before Get only when existence matters
//   - Delete multiple keys at once rather than repeatedly
//
// Nil Map Behavior:
//   - Get(nil, key): Returns (zero, false) - safe
//   - GetOr(nil, key, def): Returns def - safe
//   - Set(nil, k, v): Returns map[k]v - creates new map
//   - Delete(nil): Returns nil - preserves nil
//   - Has(nil, key): Returns false - safe
//   - IsEmpty(nil): Returns true - len(nil) == 0
//   - Len(nil): Returns 0 - len(nil) == 0
//   - Clone(nil): Returns nil - preserves nil semantics
//   - Equal(nil, nil): Returns true - both empty
//
// basic.go는 Go를 위한 기본 맵 작업을 제공합니다.
//
// 이 파일은 maputil 패키지의 기초를 이루는 핵심 맵 조작 함수를 구현합니다.
// 모든 작업은 Go 제네릭을 사용하여 타입 안전하며 불변 설계 패턴을 따릅니다
// (원본 맵은 변경되지 않음).
//
// 핵심 작업:
//
// 키-값 접근:
//   - Get(m, key): 존재 확인과 함께 값 가져오기
//     시간: O(1), 공간: O(1)
//     키 존재 시 (값, true), 없으면 (zero, false) 반환
//     예: Get({"a": 1}, "a") = (1, true)
//     사용 사례: 안전한 키 접근, 조건부 로직, 기본값 처리
//
//   - GetOr(m, key, default): 폴백이 있는 값 가져오기
//     시간: O(1), 공간: O(1)
//     키 존재 시 값, 없으면 기본값 반환
//     예: GetOr({"a": 1}, "b", 10) = 10
//     사용 사례: 설정 읽기, 안전한 기본값, 누락된 키 처리
//
// 맵 수정:
//   - Set(m, key, value): 키-값 쌍 추가 또는 업데이트
//     시간: O(n), 공간: O(n)
//     추가/업데이트된 쌍이 있는 새 맵 생성
//     불변: 원본 맵 변경 없음
//     예: Set({"a": 1}, "b", 2) = {"a": 1, "b": 2}
//     사용 사례: 불변 업데이트, 함수형 파이프라인, 상태 관리
//
//   - Delete(m, keys...): 여러 키 제거
//     시간: O(n + k) (k = len(keys)), 공간: O(n)
//     지정된 키 없이 새 맵 생성
//     불변: 원본 맵 변경 없음
//     O(1) 키 조회를 위해 집합 사용
//     예: Delete({"a": 1, "b": 2, "c": 3}, "b", "c") = {"a": 1}
//     사용 사례: 필터링, 정리, 선택적 제거
//
// 존재 확인:
//   - Has(m, key): 키 존재 확인
//     시간: O(1), 공간: O(1)
//     불리언 존재 확인
//     예: Has({"a": 1}, "a") = true
//     사용 사례: 검증, 조건부 로직, 사전 확인
//
//   - IsEmpty(m): 맵에 항목이 없는지 확인
//     시간: O(1), 공간: O(1)
//     len(m) == 0이면 true 반환
//     예: IsEmpty({}) = true
//     사용 사례: 검증, 조기 반환, 에러 확인
//
//   - IsNotEmpty(m): 맵에 항목이 있는지 확인
//     시간: O(1), 공간: O(1)
//     len(m) > 0이면 true 반환
//     예: IsNotEmpty({"a": 1}) = true
//     사용 사례: 가드 절, 검증, 긍정 확인
//
// 맵 속성:
//   - Len(m): 항목 수 가져오기
//     시간: O(1), 공간: O(1)
//     키-값 쌍 개수 반환
//     예: Len({"a": 1, "b": 2}) = 2
//     사용 사례: 크기 검증, 용량 계획, 통계
//
//   - Clear(m): 빈 맵 생성
//     시간: O(1), 공간: O(1)
//     같은 타입의 새 빈 맵 반환
//     불변: 원본 맵 변경 없음
//     예: Clear({"a": 1, "b": 2}) = {}
//     사용 사례: 상태 재설정, 초기화, 정리
//
// 복사 및 비교:
//   - Clone(m): 얕은 복사본 생성
//     시간: O(n), 공간: O(n)
//     맵 구조 복제
//     Nil 안전: nil 입력 시 nil 반환
//     얕은 복사: 참조 복사, 깊은 값 아님
//     예: Clone({"a": 1, "b": 2}) = {"a": 1, "b": 2} (독립적)
//     사용 사례: 방어적 복사, 불변 패턴, 안전한 변경
//
//   - Equal(m1, m2): 두 맵의 동등성 비교
//     시간: O(n), 공간: O(1)
//     같은 키와 값 확인
//     V가 comparable 필요
//     예: Equal({"a": 1}, {"a": 1}) = true
//     사용 사례: 테스트, 검증, 변경 감지
//
// 설계 원칙:
//   - 불변성: 모든 수정 작업이 새 맵 반환
//   - 타입 안전성: 컴파일 타임 안전성을 위해 Go 제네릭 활용
//   - Nil 안전성: nil 맵을 우아하게 처리 (Clone은 nil 보존)
//   - 성능: 접근은 O(1), 수정은 O(n)
//   - 일관성: 모든 작업에서 예측 가능한 동작
//
// 비교: Set vs 네이티브 할당:
//   - Set: 새 맵 반환, 불변, 함수형 스타일
//   - 네이티브 (m[k] = v): 원본 변경, 명령형 스타일
//   - Set: 동시 접근 패턴에 적합
//   - 네이티브: 변경이 안전할 때 더 효율적
//
// 비교: Delete vs 네이티브 delete():
//   - Delete: 새 맵 반환, 불변, 다중 키 지원
//   - 네이티브 delete(): 원본 변경, 단일 키
//   - Delete: 함수형 파이프라인에 적합
//   - 네이티브: 단일 삭제에 더 효율적
//
// 비교: Clone vs 수동 복사:
//   - Clone: 한 줄, nil 안전, 명확한 의도
//   - 수동: 루프 필요, 수동 nil 확인
//   - Clone: 코드베이스 전체에서 일관된 동작
//   - 수동: 더 장황하고 오류 가능성
//
// 성능 특성:
//
// 시간 복잡도:
//   - Get/GetOr: O(1) - 해시 테이블 조회
//   - Has/IsEmpty/IsNotEmpty/Len: O(1) - 직접 확인
//   - Clear: O(1) - 새 빈 맵 할당
//   - Set: O(n) - 모든 기존 항목 + 새 항목 복사
//   - Delete: O(n + k) - k개 키 제외 모든 항목 복사
//   - Clone: O(n) - 모든 항목 복사
//   - Equal: O(n) - 모든 항목 비교
//
// 공간 복잡도:
//   - Get/GetOr/Has/IsEmpty/IsNotEmpty/Len: O(1) - 할당 없음
//   - Clear: O(1) - 빈 맵
//   - Set/Delete/Clone: O(n) - n개 항목의 새 맵
//   - Equal: O(1) - 할당 없음 (비교만)
//
// 메모리 할당:
//   - 불변 작업 (Set, Delete, Clone)은 새 맵 할당
//   - 읽기 작업 (Get, Has, Len)은 할당 없음
//   - 고빈도 업데이트의 경우 가변 작업 고려
//   - Clone은 얕은 복사: 중첩 구조는 참조 공유
//
// 일반적인 사용 패턴:
//
//	// 기본값이 있는 안전한 키 접근
//	config := map[string]string{"host": "localhost"}
//	port := maputil.GetOr(config, "port", "8080")
//
//	// 파이프라인의 불변 업데이트
//	m := map[string]int{"a": 1, "b": 2}
//	m = maputil.Set(m, "c", 3)
//	m = maputil.Set(m, "d", 4)
//	// 각 단계에서 원본 {"a": 1, "b": 2} 변경 없음
//
//	// 방어적 복사
//	original := map[string]int{"x": 1, "y": 2}
//	working := maputil.Clone(original)
//	working["z"] = 3 // original 변경 없음
//
//	// 대량 키 제거
//	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	filtered := maputil.Delete(data, "b", "d")
//	// 결과: {"a": 1, "c": 3}
//
//	// 검증 체인
//	if maputil.IsNotEmpty(cache) && maputil.Has(cache, key) {
//	    value := cache[key]
//	    // 값 처리
//	}
//
//	// 동등성 테스트
//	expected := map[string]int{"a": 1, "b": 2}
//	actual := processData()
//	if !maputil.Equal(expected, actual) {
//	    t.Errorf("불일치: got %v, want %v", actual, expected)
//	}
//
// 불변성 패턴:
//
//	// 패턴 1: 함수형 업데이트
//	state := map[string]int{"count": 0}
//	state = maputil.Set(state, "count", 1)
//	state = maputil.Set(state, "max", 10)
//
//	// 패턴 2: 조건부 업데이트
//	var newState map[string]int
//	if shouldUpdate {
//	    newState = maputil.Set(state, key, value)
//	} else {
//	    newState = state
//	}
//
//	// 패턴 3: 파이프라인 변환
//	result := maputil.Set(
//	    maputil.Delete(original, "temp"),
//	    "final", computeValue(),
//	)
//
// 엣지 케이스:
//   - nil 맵에 Get: (zero, false) 반환 - 패닉 없음
//   - nil 맵에 Set: 하나의 항목이 있는 새 맵 생성
//   - 키 없이 Delete: 원본의 복제 반환
//   - 존재하지 않는 키로 Delete: 누락된 키 무시
//   - nil의 Clone: nil 반환 (nil 의미론 보존)
//   - 다른 길이로 Equal: 즉시 false 반환 (최적화)
//   - nil에 IsEmpty: true 반환 (len(nil) == 0)
//   - nil 맵에 Has: false 반환 - 패닉 없음
//
// 스레드 안전성:
//   - 읽기 작업 (Get, Has, Len 등): 동시 쓰기 없으면 안전
//   - 쓰기 작업: 새 맵 생성, 다른 맵에 안전
//   - 같은 맵: 동기화 없이는 스레드 안전하지 않음
//   - Clone: 고루틴 간 맵 공유의 안전한 방법
//   - 불변 패턴: 동시성에 자연스럽게 더 안전
//
// 성능 팁:
//   - 기본값을 위해 Get + if-else 대신 GetOr 사용
//   - 가능하면 여러 Set 작업 일괄 처리
//   - 고빈도 업데이트는 네이티브 변경 고려
//   - 필요할 때만 Clone (복사 비용은 O(n))
//   - 존재가 중요할 때만 Get 전에 Has 사용
//   - 반복적으로가 아닌 여러 키를 한 번에 Delete
//
// Nil 맵 동작:
//   - Get(nil, key): (zero, false) 반환 - 안전
//   - GetOr(nil, key, def): def 반환 - 안전
//   - Set(nil, k, v): map[k]v 반환 - 새 맵 생성
//   - Delete(nil): nil 반환 - nil 보존
//   - Has(nil, key): false 반환 - 안전
//   - IsEmpty(nil): true 반환 - len(nil) == 0
//   - Len(nil): 0 반환 - len(nil) == 0
//   - Clone(nil): nil 반환 - nil 의미론 보존
//   - Equal(nil, nil): true 반환 - 둘 다 비어 있음

// Get retrieves the value for a given key and returns whether the key exists.
// Get은 주어진 키의 값을 가져오고 키의 존재 여부를 반환합니다.
//
// Returns the value and true if the key exists, or the zero value and false otherwise.
// 키가 존재하면 값과 true를 반환하고, 그렇지 않으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(1)
// 시간 복잡도: O(1)
//
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//
// result := maputil.Set(m, "c", 3) // map[string]int{"a": 1, "b": 2, "c": 3}
// Original map m is unchanged
// 원본 맵 m은 변경되지 않음
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
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//
// result := maputil.Delete(m, "b", "c") // map[string]int{"a": 1}
// Original map m is unchanged
// 원본 맵 m은 변경되지 않음
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//
// result := maputil.Clear(m) // map[string]int{}
// Original map m is unchanged
// 원본 맵 m은 변경되지 않음
func Clear[K comparable, V any](m map[K]V) map[K]V {
	return make(map[K]V)
}

// Clone creates a shallow copy of the map.
// Clone은 맵의 얕은 복사본을 생성합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	clone := maputil.Clone(m) // map[string]int{"a": 1, "b": 2}
//
// clone["c"] = 3
// Original map m is unchanged
// 원본 맵 m은 변경되지 않음
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
// Example
// 예제:
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
