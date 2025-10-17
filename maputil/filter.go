package maputil

// filter.go provides map filtering operations for Go.
//
// This file implements functions for selectively including or excluding map entries
// based on predicates or explicit key lists. All operations are immutable and return
// new maps without modifying the original.
//
// Filtering Operations:
//
// Predicate-Based Filtering:
//   - Filter(m, fn): Include entries matching predicate
//     Time: O(n), Space: O(n)
//     Function receives key and value
//     Returns new map with matching entries
//     Example: Filter({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 1 }) = {"b": 2, "c": 3}
//     Use cases: Data filtering, validation, conditional selection
//
//   - FilterKeys(m, fn): Include entries with matching keys
//     Time: O(n), Space: O(n)
//     Function receives only key (simpler)
//     Values passed through unchanged
//     Example: FilterKeys({"apple": 1, "banana": 2}, func(k) { return len(k) > 5 }) = {"banana": 2}
//     Use cases: Key pattern matching, namespace filtering, prefix selection
//
//   - FilterValues(m, fn): Include entries with matching values
//     Time: O(n), Space: O(n)
//     Function receives only value (simpler)
//     Keys passed through unchanged
//     Example: FilterValues({"a": 1, "b": 2, "c": 3}, func(v) { return v%2 == 0 }) = {"b": 2}
//     Use cases: Value-based filtering, range selection, type filtering
//
// Explicit Key Selection:
//   - Pick(m, keys...): Include only specified keys
//     Time: O(k) where k = len(keys), Space: O(k)
//     Direct key selection (no predicate)
//     Non-existent keys ignored
//     Example: Pick({"a": 1, "b": 2, "c": 3}, "a", "c") = {"a": 1, "c": 3}
//     Use cases: API projections, field selection, whitelisting
//
//   - Omit(m, keys...): Exclude specified keys
//     Time: O(n), Space: O(n)
//     Direct key exclusion (no predicate)
//     Non-existent keys ignored
//     Example: Omit({"a": 1, "b": 2, "c": 3}, "b") = {"a": 1, "c": 3}
//     Use cases: Removing sensitive fields, blacklisting, cleanup
//
// Inverse Operations:
//   - OmitBy(m, fn): Exclude entries matching predicate
//     Time: O(n), Space: O(n)
//     Inverse of Filter
//     Function receives key and value
//     Returns entries NOT matching predicate
//     Example: OmitBy({"a": 1, "b": 2, "c": 3}, func(k, v) { return v < 2 }) = {"b": 2, "c": 3}
//     Use cases: Exclusion filtering, blacklisting by condition, negative selection
//
//   - PickBy(m, fn): Include entries matching predicate (alias)
//     Time: O(n), Space: O(n)
//     Alias for Filter (semantic clarity)
//     Functionally identical to Filter
//     Example: PickBy({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 1 }) = {"b": 2, "c": 3}
//     Use cases: Same as Filter, emphasizes "picking" semantics
//
// Design Principles:
//   - Immutability: All operations return new maps
//   - Flexibility: Multiple filtering strategies (predicate, keys, values)
//   - Consistency: Predictable behavior with edge cases
//   - Performance: Efficient algorithms for each use case
//   - Clarity: Semantic function names (Filter vs Omit, Pick vs PickBy)
//
// Comparison: Filter vs FilterKeys/FilterValues:
//   - Filter: Access to both key and value in predicate
//   - FilterKeys: Access to key only (simpler, clearer intent)
//   - FilterValues: Access to value only (simpler, clearer intent)
//   - Filter: More flexible but potentially more complex
//   - FilterKeys/FilterValues: More focused and easier to reason about
//
// Comparison: Pick vs Filter:
//   - Pick: Explicit key list, O(k) where k = len(keys)
//   - Filter: Predicate-based, O(n) where n = len(map)
//   - Pick: More efficient for small key sets
//   - Filter: More flexible for complex conditions
//
// Comparison: Omit vs OmitBy:
//   - Omit: Explicit key list, direct exclusion
//   - OmitBy: Predicate-based, conditional exclusion
//   - Omit: Simpler for known keys
//   - OmitBy: Flexible for dynamic conditions
//
// Comparison: Filter vs OmitBy:
//   - Filter: Include matching entries
//   - OmitBy: Exclude matching entries
//   - Filter(m, fn) ≈ OmitBy(m, !fn)
//   - Use Filter for positive selection (what to keep)
//   - Use OmitBy for negative selection (what to remove)
//
// Performance Characteristics:
//
// Time Complexity:
//   - Filter/FilterKeys/FilterValues: O(n) - Full iteration
//   - OmitBy/PickBy: O(n) - Full iteration
//   - Pick: O(k) - Only iterate requested keys
//   - Omit: O(n) - Must check all keys against omit set
//
// Space Complexity:
//   - All operations: O(m) where m = result size
//   - Filter/FilterKeys/FilterValues: Proportional to matches
//   - Pick: At most O(k) where k = requested keys
//   - Omit: O(n - k) where k = omitted keys
//   - OmitBy/PickBy: Proportional to non-matches/matches
//
// Memory Allocation:
//   - All create new maps (immutable)
//   - Result size depends on filter selectivity
//   - Pick/Omit create temporary key sets for O(1) lookup
//   - Empty results still allocate map header
//
// Common Usage Patterns:
//
//	// Filter by value range
//	scores := map[string]int{"Alice": 85, "Bob": 92, "Charlie": 78}
//	passing := maputil.FilterValues(scores, func(v int) bool {
//	    return v >= 80
//	}) // {"Alice": 85, "Bob": 92}
//
//	// Filter by key pattern
//	config := map[string]string{"db.host": "localhost", "db.port": "5432", "app.name": "myapp"}
//	dbConfig := maputil.FilterKeys(config, func(k string) bool {
//	    return strings.HasPrefix(k, "db.")
//	}) // {"db.host": "localhost", "db.port": "5432"}
//
//	// Pick specific fields (API projection)
//	user := map[string]interface{}{
//	    "id": 123, "name": "Alice", "email": "alice@example.com",
//	    "password": "secret", "internal_id": 456,
//	}
//	publicUser := maputil.Pick(user, "id", "name", "email")
//	// {"id": 123, "name": "Alice", "email": "alice@example.com"}
//
//	// Omit sensitive fields
//	data := map[string]string{
//	    "username": "alice", "password": "secret",
//	    "email": "alice@example.com", "token": "xyz123",
//	}
//	safeData := maputil.Omit(data, "password", "token")
//	// {"username": "alice", "email": "alice@example.com"}
//
//	// Remove invalid entries
//	entries := map[string]int{"a": 1, "b": -1, "c": 5, "d": -2}
//	valid := maputil.OmitBy(entries, func(k string, v int) bool {
//	    return v < 0
//	}) // {"a": 1, "c": 5}
//
// Chaining Filters:
//
//	// Multiple filter stages
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
//	result := maputil.FilterValues(
//	    maputil.FilterKeys(m, func(k string) bool {
//	        return k >= "b" && k <= "d"
//	    }),
//	    func(v int) bool {
//	        return v%2 == 0
//	    },
//	) // {"b": 2, "d": 4}
//
//	// Combined with transformations
//	data := map[string]int{"a": 1, "b": 0, "c": 3, "d": 0}
//	processed := maputil.MapValues(
//	    maputil.Compact(data), // Remove zeros
//	    func(v int) int { return v * 2 },
//	) // {"a": 2, "c": 6}
//
// Edge Cases:
//   - Filter on empty map: Returns empty map
//   - Filter with predicate that matches nothing: Returns empty map
//   - FilterKeys/FilterValues on empty: Returns empty map
//   - Pick with no keys: Returns empty map
//   - Pick with non-existent keys: Ignores missing keys
//   - Pick(m, "x") where "x" not in m: Returns empty map
//   - Omit with no keys: Returns clone of original
//   - Omit with all keys: Returns empty map
//   - OmitBy with predicate matching all: Returns empty map
//   - PickBy is exactly Filter (implementation alias)
//
// Nil Map Behavior:
//   - Filter(nil, fn): Returns empty map
//   - FilterKeys(nil, fn): Returns empty map
//   - FilterValues(nil, fn): Returns empty map
//   - Pick(nil, keys...): Returns empty map
//   - Omit(nil, keys...): Returns nil if no keys, else empty map
//   - OmitBy(nil, fn): Returns empty map
//   - PickBy(nil, fn): Returns empty map (calls Filter)
//
// Thread Safety:
//   - All operations create new maps (safe for concurrent reads of input)
//   - Not safe if input map is modified concurrently
//   - Output maps are independent (safe to use in different goroutines)
//   - Predicate functions must be goroutine-safe if accessing shared state
//
// Performance Tips:
//   - Use Pick instead of Filter when you know exact keys
//   - FilterKeys/FilterValues are clearer and slightly faster than Filter
//   - For small key sets, Pick is O(k) vs Filter's O(n)
//   - Omit builds temporary set for O(1) key lookup
//   - Chain filters carefully: each creates intermediate map
//   - Consider combining predicates in single Filter call
//   - Omit with many keys: consider FilterKeys with inverse predicate
//
// Predicate Best Practices:
//   - Keep predicates pure (no side effects)
//   - Predicates should be fast (called for every entry)
//   - Avoid expensive operations in predicates
//   - Consider caching results if predicate is expensive
//   - Use FilterKeys/FilterValues when only one component needed
//
// filter.go는 Go를 위한 맵 필터링 작업을 제공합니다.
//
// 이 파일은 조건 또는 명시적 키 목록을 기반으로 맵 항목을 선택적으로 포함하거나
// 제외하는 함수를 구현합니다. 모든 작업은 불변이며 원본을 수정하지 않고 새 맵을
// 반환합니다.
//
// 필터링 작업:
//
// 조건 기반 필터링:
//   - Filter(m, fn): 조건과 일치하는 항목 포함
//     시간: O(n), 공간: O(n)
//     함수가 키와 값 수신
//     일치하는 항목이 있는 새 맵 반환
//     예: Filter({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 1 }) = {"b": 2, "c": 3}
//     사용 사례: 데이터 필터링, 검증, 조건부 선택
//
//   - FilterKeys(m, fn): 일치하는 키가 있는 항목 포함
//     시간: O(n), 공간: O(n)
//     함수가 키만 수신 (더 단순)
//     값은 변경 없이 전달
//     예: FilterKeys({"apple": 1, "banana": 2}, func(k) { return len(k) > 5 }) = {"banana": 2}
//     사용 사례: 키 패턴 매칭, 네임스페이스 필터링, 접두사 선택
//
//   - FilterValues(m, fn): 일치하는 값이 있는 항목 포함
//     시간: O(n), 공간: O(n)
//     함수가 값만 수신 (더 단순)
//     키는 변경 없이 전달
//     예: FilterValues({"a": 1, "b": 2, "c": 3}, func(v) { return v%2 == 0 }) = {"b": 2}
//     사용 사례: 값 기반 필터링, 범위 선택, 타입 필터링
//
// 명시적 키 선택:
//   - Pick(m, keys...): 지정된 키만 포함
//     시간: O(k) (k = len(keys)), 공간: O(k)
//     직접 키 선택 (조건 없음)
//     존재하지 않는 키 무시
//     예: Pick({"a": 1, "b": 2, "c": 3}, "a", "c") = {"a": 1, "c": 3}
//     사용 사례: API 프로젝션, 필드 선택, 화이트리스트
//
//   - Omit(m, keys...): 지정된 키 제외
//     시간: O(n), 공간: O(n)
//     직접 키 제외 (조건 없음)
//     존재하지 않는 키 무시
//     예: Omit({"a": 1, "b": 2, "c": 3}, "b") = {"a": 1, "c": 3}
//     사용 사례: 민감한 필드 제거, 블랙리스트, 정리
//
// 역 작업:
//   - OmitBy(m, fn): 조건과 일치하는 항목 제외
//     시간: O(n), 공간: O(n)
//     Filter의 역
//     함수가 키와 값 수신
//     조건과 일치하지 않는 항목 반환
//     예: OmitBy({"a": 1, "b": 2, "c": 3}, func(k, v) { return v < 2 }) = {"b": 2, "c": 3}
//     사용 사례: 제외 필터링, 조건별 블랙리스트, 부정 선택
//
//   - PickBy(m, fn): 조건과 일치하는 항목 포함 (별칭)
//     시간: O(n), 공간: O(n)
//     Filter의 별칭 (의미적 명확성)
//     기능적으로 Filter와 동일
//     예: PickBy({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 1 }) = {"b": 2, "c": 3}
//     사용 사례: Filter와 동일, "선택" 의미 강조
//
// 설계 원칙:
//   - 불변성: 모든 작업이 새 맵 반환
//   - 유연성: 여러 필터링 전략 (조건, 키, 값)
//   - 일관성: 엣지 케이스에서 예측 가능한 동작
//   - 성능: 각 사용 사례에 효율적인 알고리즘
//   - 명확성: 의미적 함수 이름 (Filter vs Omit, Pick vs PickBy)
//
// 비교: Filter vs FilterKeys/FilterValues:
//   - Filter: 조건에서 키와 값 모두 접근
//   - FilterKeys: 키만 접근 (더 단순, 의도 명확)
//   - FilterValues: 값만 접근 (더 단순, 의도 명확)
//   - Filter: 더 유연하지만 잠재적으로 더 복잡
//   - FilterKeys/FilterValues: 더 집중적이고 추론하기 쉬움
//
// 비교: Pick vs Filter:
//   - Pick: 명시적 키 목록, O(k) (k = len(keys))
//   - Filter: 조건 기반, O(n) (n = len(map))
//   - Pick: 작은 키 집합에 더 효율적
//   - Filter: 복잡한 조건에 더 유연
//
// 비교: Omit vs OmitBy:
//   - Omit: 명시적 키 목록, 직접 제외
//   - OmitBy: 조건 기반, 조건부 제외
//   - Omit: 알려진 키에 더 단순
//   - OmitBy: 동적 조건에 유연
//
// 비교: Filter vs OmitBy:
//   - Filter: 일치하는 항목 포함
//   - OmitBy: 일치하는 항목 제외
//   - Filter(m, fn) ≈ OmitBy(m, !fn)
//   - 긍정 선택(유지할 것)에 Filter 사용
//   - 부정 선택(제거할 것)에 OmitBy 사용
//
// 성능 특성:
//
// 시간 복잡도:
//   - Filter/FilterKeys/FilterValues: O(n) - 전체 반복
//   - OmitBy/PickBy: O(n) - 전체 반복
//   - Pick: O(k) - 요청된 키만 반복
//   - Omit: O(n) - 생략 집합에 대해 모든 키 확인 필요
//
// 공간 복잡도:
//   - 모든 작업: O(m) (m = 결과 크기)
//   - Filter/FilterKeys/FilterValues: 일치 비율에 비례
//   - Pick: 최대 O(k) (k = 요청된 키)
//   - Omit: O(n - k) (k = 생략된 키)
//   - OmitBy/PickBy: 불일치/일치 비율에 비례
//
// 메모리 할당:
//   - 모두 새 맵 생성 (불변)
//   - 결과 크기는 필터 선택성에 따라 다름
//   - Pick/Omit은 O(1) 조회를 위해 임시 키 집합 생성
//   - 빈 결과도 맵 헤더 할당
//
// 일반적인 사용 패턴:
//
//	// 값 범위로 필터링
//	scores := map[string]int{"Alice": 85, "Bob": 92, "Charlie": 78}
//	passing := maputil.FilterValues(scores, func(v int) bool {
//	    return v >= 80
//	}) // {"Alice": 85, "Bob": 92}
//
//	// 키 패턴으로 필터링
//	config := map[string]string{"db.host": "localhost", "db.port": "5432", "app.name": "myapp"}
//	dbConfig := maputil.FilterKeys(config, func(k string) bool {
//	    return strings.HasPrefix(k, "db.")
//	}) // {"db.host": "localhost", "db.port": "5432"}
//
//	// 특정 필드 선택 (API 프로젝션)
//	user := map[string]interface{}{
//	    "id": 123, "name": "Alice", "email": "alice@example.com",
//	    "password": "secret", "internal_id": 456,
//	}
//	publicUser := maputil.Pick(user, "id", "name", "email")
//	// {"id": 123, "name": "Alice", "email": "alice@example.com"}
//
//	// 민감한 필드 생략
//	data := map[string]string{
//	    "username": "alice", "password": "secret",
//	    "email": "alice@example.com", "token": "xyz123",
//	}
//	safeData := maputil.Omit(data, "password", "token")
//	// {"username": "alice", "email": "alice@example.com"}
//
//	// 유효하지 않은 항목 제거
//	entries := map[string]int{"a": 1, "b": -1, "c": 5, "d": -2}
//	valid := maputil.OmitBy(entries, func(k string, v int) bool {
//	    return v < 0
//	}) // {"a": 1, "c": 5}
//
// 필터 체인:
//
//	// 여러 필터 단계
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
//	result := maputil.FilterValues(
//	    maputil.FilterKeys(m, func(k string) bool {
//	        return k >= "b" && k <= "d"
//	    }),
//	    func(v int) bool {
//	        return v%2 == 0
//	    },
//	) // {"b": 2, "d": 4}
//
//	// 변환과 결합
//	data := map[string]int{"a": 1, "b": 0, "c": 3, "d": 0}
//	processed := maputil.MapValues(
//	    maputil.Compact(data), // zero 제거
//	    func(v int) int { return v * 2 },
//	) // {"a": 2, "c": 6}
//
// 엣지 케이스:
//   - 빈 맵에 Filter: 빈 맵 반환
//   - 일치하는 것 없는 조건으로 Filter: 빈 맵 반환
//   - 빈 것에 FilterKeys/FilterValues: 빈 맵 반환
//   - 키 없이 Pick: 빈 맵 반환
//   - 존재하지 않는 키로 Pick: 누락된 키 무시
//   - Pick(m, "x") (m에 "x" 없음): 빈 맵 반환
//   - 키 없이 Omit: 원본의 복제 반환
//   - 모든 키로 Omit: 빈 맵 반환
//   - 모두 일치하는 조건으로 OmitBy: 빈 맵 반환
//   - PickBy는 정확히 Filter (구현 별칭)
//
// Nil 맵 동작:
//   - Filter(nil, fn): 빈 맵 반환
//   - FilterKeys(nil, fn): 빈 맵 반환
//   - FilterValues(nil, fn): 빈 맵 반환
//   - Pick(nil, keys...): 빈 맵 반환
//   - Omit(nil, keys...): 키 없으면 nil, 있으면 빈 맵 반환
//   - OmitBy(nil, fn): 빈 맵 반환
//   - PickBy(nil, fn): 빈 맵 반환 (Filter 호출)
//
// 스레드 안전성:
//   - 모든 작업이 새 맵 생성 (입력의 동시 읽기 안전)
//   - 입력 맵이 동시에 수정되면 안전하지 않음
//   - 출력 맵은 독립적 (다른 고루틴에서 사용 안전)
//   - 조건 함수가 공유 상태 접근 시 고루틴 안전해야 함
//
// 성능 팁:
//   - 정확한 키를 알 때 Filter 대신 Pick 사용
//   - FilterKeys/FilterValues가 Filter보다 명확하고 약간 빠름
//   - 작은 키 집합의 경우 Pick은 O(k) vs Filter의 O(n)
//   - Omit은 O(1) 키 조회를 위해 임시 집합 구축
//   - 필터 체인 신중히: 각각 중간 맵 생성
//   - 조건을 단일 Filter 호출로 결합 고려
//   - 많은 키로 Omit: 역 조건으로 FilterKeys 고려
//
// 조건 모범 사례:
//   - 조건을 순수하게 유지 (부수 효과 없음)
//   - 조건은 빨라야 함 (모든 항목마다 호출)
//   - 조건에서 비용이 큰 작업 피하기
//   - 조건이 비싸면 결과 캐싱 고려
//   - 한 구성 요소만 필요할 때 FilterKeys/FilterValues 사용

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
