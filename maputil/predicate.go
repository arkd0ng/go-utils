package maputil

// predicate.go provides predicate-based map testing operations for Go.
//
// This file implements boolean functions that test map properties and contents
// using predicates or direct value checks. All operations are read-only and
// return boolean results.
//
// Predicate Testing Operations:
//
// Quantifier Predicates:
//   - Every(m, fn): All entries satisfy predicate (∀)
//     Time: O(n) worst case, early termination
//     Returns true if all entries match
//     Returns true for empty map (vacuous truth)
//     Short-circuits on first false
//     Example: Every({"a": 2, "b": 4, "c": 6}, func(k, v) { return v%2 == 0 }) = true
//     Use cases: Validation, consistency checks, all-or-nothing logic
//
//   - Some(m, fn): At least one entry satisfies predicate (∃)
//     Time: O(n) worst case, early termination
//     Returns true if any entry matches
//     Returns false for empty map
//     Short-circuits on first true
//     Example: Some({"a": 1, "b": 2, "c": 3}, func(k, v) { return v%2 == 0 }) = true
//     Use cases: Existence checks, any-match logic, search predicates
//
//   - None(m, fn): No entries satisfy predicate (¬∃)
//     Time: O(n) worst case, early termination
//     Returns true if no entries match
//     Returns true for empty map (vacuous truth)
//     Equivalent to !Some(m, fn)
//     Short-circuits on first true
//     Example: None({"a": 1, "b": 3, "c": 5}, func(k, v) { return v%2 == 0 }) = true
//     Use cases: Exclusion validation, negative checks, prohibition rules
//
// Existence Checks:
//   - HasKey(m, key): Check if key exists
//     Time: O(1), Space: O(1)
//     Alias for Has from basic.go
//     Hash table lookup
//     Example: HasKey({"a": 1, "b": 2}, "a") = true
//     Use cases: Key validation, membership testing, precondition checks
//
//   - HasValue(m, value): Check if value exists
//     Time: O(n), Space: O(1)
//     Linear search through values
//     Requires V to be comparable
//     Example: HasValue({"a": 1, "b": 2, "c": 3}, 2) = true
//     Use cases: Value validation, reverse lookup existence, data presence
//
//   - HasEntry(m, key, value): Check if key-value pair exists
//     Time: O(1), Space: O(1)
//     Checks both key existence and value match
//     More specific than HasKey or HasValue
//     Example: HasEntry({"a": 1, "b": 2}, "b", 2) = true
//     Use cases: Exact match validation, state verification, pair checking
//
// Set Relationships:
//   - IsSubset(subset, superset): Check subset relationship
//     Time: O(n) where n = len(subset), Space: O(1)
//     True if all subset entries exist in superset
//     Empty map is subset of any map
//     Proper subset: len(subset) < len(superset)
//     Example: IsSubset({"a": 1}, {"a": 1, "b": 2}) = true
//     Use cases: Set inclusion, permission checking, partial matching
//
//   - IsSuperset(superset, subset): Check superset relationship
//     Time: O(n) where n = len(subset), Space: O(1)
//     True if superset contains all subset entries
//     Alias: IsSuperset(a, b) = IsSubset(b, a)
//     Any map is superset of empty map
//     Example: IsSuperset({"a": 1, "b": 2}, {"a": 1}) = true
//     Use cases: Set containment, capability checking, coverage validation
//
// Design Principles:
//   - Short-circuiting: Every/Some/None terminate early when possible
//   - Vacuous truth: Every/None return true for empty maps
//   - Type safety: HasValue/HasEntry require comparable values
//   - Efficiency: O(1) operations for key checks, O(n) for value searches
//   - Consistency: Predictable behavior with empty maps
//
// Comparison: Every vs Some vs None:
//   - Every: ∀x P(x) - All must match
//   - Some: ∃x P(x) - At least one matches
//   - None: ¬∃x P(x) - None match
//   - Logical relationships:
//     * None(m, fn) = !Some(m, fn)
//     * Every(m, fn) = None(m, !fn)
//     * !Every(m, fn) = Some(m, !fn)
//
// Comparison: HasKey vs HasValue:
//   - HasKey: O(1) hash lookup
//   - HasValue: O(n) linear search
//   - HasKey: More efficient
//   - HasValue: More expensive but necessary for reverse lookups
//
// Comparison: HasValue vs Some:
//   - HasValue: Direct value equality check
//   - Some: Predicate-based, more flexible
//   - HasValue: Simpler for exact matches
//   - Some: Better for complex conditions
//
// Comparison: IsSubset vs Equal:
//   - IsSubset: One-way containment
//   - Equal: Bidirectional equality
//   - IsSubset(a, b) && IsSubset(b, a) = Equal(a, b)
//   - IsSubset allows superset to have extra entries
//
// Performance Characteristics:
//
// Time Complexity:
//   - Every/Some/None: O(n) worst case
//     * Best case: O(1) with early termination
//     * Average: O(n/2) with random distribution
//   - HasKey/HasEntry: O(1) - Hash table lookup
//   - HasValue: O(n) - Must check all values
//   - IsSubset/IsSuperset: O(n) where n = subset size
//
// Space Complexity:
//   - All operations: O(1) - No additional allocation
//   - Read-only operations, no new data structures
//
// Short-Circuit Optimization:
//   - Every: Returns false on first non-matching entry
//   - Some: Returns true on first matching entry
//   - None: Returns false on first matching entry
//   - Practical speedup: Often O(1) instead of O(n)
//
// Common Usage Patterns:
//
//	// Validate all entries meet criteria
//	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 28}
//	allAdults := maputil.Every(ages, func(k string, v int) bool {
//	    return v >= 18
//	}) // true
//
//	// Check if any entry meets criteria
//	scores := map[string]int{"Alice": 85, "Bob": 72, "Charlie": 65}
//	anyonePassedExcellent := maputil.Some(scores, func(k string, v int) bool {
//	    return v >= 90
//	}) // false
//
//	// Ensure no entries violate rule
//	prices := map[string]float64{"item1": 10.5, "item2": 25.0, "item3": 5.99}
//	noNegativePrices := maputil.None(prices, func(k string, v float64) bool {
//	    return v < 0
//	}) // true
//
//	// Check value existence
//	userRoles := map[string]string{"user1": "admin", "user2": "editor"}
//	hasAdmin := maputil.HasValue(userRoles, "admin") // true
//
//	// Verify exact key-value pair
//	config := map[string]string{"env": "production", "debug": "false"}
//	isProduction := maputil.HasEntry(config, "env", "production") // true
//
//	// Check subset relationship
//	required := map[string]string{"name": "Alice", "role": "admin"}
//	actual := map[string]string{"name": "Alice", "role": "admin", "email": "alice@example.com"}
//	hasRequired := maputil.IsSubset(required, actual) // true
//
// Validation Patterns:
//
//	// All fields present and valid
//	fields := map[string]string{"name": "Alice", "email": "alice@example.com"}
//	allValid := maputil.Every(fields, func(k string, v string) bool {
//	    return v != "" && len(v) > 0
//	})
//
//	// No forbidden values
//	data := map[string]int{"a": 1, "b": 2, "c": 3}
//	noZeros := maputil.None(data, func(k string, v int) bool {
//	    return v == 0
//	})
//
//	// At least one required field present
//	optionalFields := map[string]string{"phone": "", "address": "123 Main St"}
//	hasAnyContact := maputil.Some(optionalFields, func(k string, v string) bool {
//	    return v != ""
//	})
//
// Logical Combinations:
//
//	// Combine predicates
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//
//	// All positive AND all less than 10
//	valid := maputil.Every(m, func(k string, v int) bool {
//	    return v > 0 && v < 10
//	})
//
//	// Has any even OR any greater than 5
//	hasSpecial := maputil.Some(m, func(k string, v int) bool {
//	    return v%2 == 0 || v > 5
//	})
//
//	// No negatives AND no zeros
//	noInvalid := maputil.None(m, func(k string, v int) bool {
//	    return v <= 0
//	})
//
// Set Operations Examples:
//
//	// Permission checking
//	requiredPerms := map[string]bool{"read": true, "write": true}
//	userPerms := map[string]bool{"read": true, "write": true, "delete": true}
//	hasPermission := maputil.IsSubset(requiredPerms, userPerms) // true
//
//	// Configuration validation
//	defaults := map[string]string{"timeout": "30", "retries": "3"}
//	userConfig := map[string]string{"timeout": "60", "retries": "3", "debug": "true"}
//	hasAllDefaults := maputil.IsSubset(defaults, userConfig) // false (timeout value differs)
//
//	// Exact subset (all values match)
//	template := map[string]int{"x": 1, "y": 2}
//	extended := map[string]int{"x": 1, "y": 2, "z": 3}
//	isValidExtension := maputil.IsSubset(template, extended) // true
//
// Edge Cases:
//   - Every on empty map: Returns true (vacuous truth)
//   - Some on empty map: Returns false (no elements to match)
//   - None on empty map: Returns true (vacuous truth)
//   - HasKey on nil map: Returns false (safe)
//   - HasValue on nil map: Returns false (safe)
//   - HasEntry on nil map: Returns false (safe)
//   - IsSubset(empty, any): Returns true (empty is subset of all)
//   - IsSubset(any, empty): Returns false if any is non-empty
//   - IsSubset(m, m): Returns true (reflexive)
//   - IsSuperset(m, m): Returns true (reflexive)
//
// Nil Map Behavior:
//   - Every(nil, fn): Returns true (vacuous truth)
//   - Some(nil, fn): Returns false (no elements)
//   - None(nil, fn): Returns true (vacuous truth)
//   - HasKey(nil, key): Returns false
//   - HasValue(nil, value): Returns false
//   - HasEntry(nil, k, v): Returns false
//   - IsSubset(nil, m): Returns true (empty subset)
//   - IsSubset(m, nil): Returns false if m non-empty
//   - IsSuperset(nil, m): Returns false if m non-empty
//   - IsSuperset(m, nil): Returns true (empty subset)
//
// Thread Safety:
//   - All operations are read-only (safe for concurrent reads)
//   - Not safe if map is modified concurrently
//   - Predicate functions must be goroutine-safe if accessing shared state
//   - No internal state, purely functional
//
// Performance Tips:
//   - Use HasKey instead of HasValue when possible (O(1) vs O(n))
//   - Every/Some/None short-circuit - put likely failures first
//   - For large maps, use Some/None over counting and comparing
//   - HasEntry is faster than HasValue + separate key check
//   - IsSubset is O(n) of subset size, not superset
//   - Cache predicate results if called multiple times
//
// Predicate Best Practices:
//   - Keep predicates pure (no side effects)
//   - Predicates should be fast (called frequently)
//   - Use descriptive predicate names
//   - Consider short-circuit behavior when designing predicates
//   - Test edge cases (empty maps, nil values)
//
// predicate.go는 Go를 위한 조건 기반 맵 테스트 작업을 제공합니다.
//
// 이 파일은 조건 또는 직접 값 확인을 사용하여 맵 속성 및 내용을 테스트하는
// 불리언 함수를 구현합니다. 모든 작업은 읽기 전용이며 불리언 결과를 반환합니다.
//
// 조건 테스트 작업:
//
// 수량자 조건:
//   - Every(m, fn): 모든 항목이 조건 만족 (∀)
//     시간: 최악 O(n), 조기 종료
//     모든 항목 일치 시 true 반환
//     빈 맵에 true 반환 (공허한 진실)
//     첫 false에서 단락
//     예: Every({"a": 2, "b": 4, "c": 6}, func(k, v) { return v%2 == 0 }) = true
//     사용 사례: 검증, 일관성 확인, 전부 아니면 전무 로직
//
//   - Some(m, fn): 최소 하나의 항목이 조건 만족 (∃)
//     시간: 최악 O(n), 조기 종료
//     어떤 항목 일치 시 true 반환
//     빈 맵에 false 반환
//     첫 true에서 단락
//     예: Some({"a": 1, "b": 2, "c": 3}, func(k, v) { return v%2 == 0 }) = true
//     사용 사례: 존재 확인, 어느 일치 로직, 검색 조건
//
//   - None(m, fn): 어떤 항목도 조건 불만족 (¬∃)
//     시간: 최악 O(n), 조기 종료
//     어떤 항목도 일치 안 하면 true 반환
//     빈 맵에 true 반환 (공허한 진실)
//     !Some(m, fn)과 동등
//     첫 true에서 단락
//     예: None({"a": 1, "b": 3, "c": 5}, func(k, v) { return v%2 == 0 }) = true
//     사용 사례: 제외 검증, 부정 확인, 금지 규칙
//
// 존재 확인:
//   - HasKey(m, key): 키 존재 확인
//     시간: O(1), 공간: O(1)
//     basic.go의 Has 별칭
//     해시 테이블 조회
//     예: HasKey({"a": 1, "b": 2}, "a") = true
//     사용 사례: 키 검증, 멤버십 테스트, 사전 조건 확인
//
//   - HasValue(m, value): 값 존재 확인
//     시간: O(n), 공간: O(1)
//     값을 통한 선형 검색
//     V가 comparable 필요
//     예: HasValue({"a": 1, "b": 2, "c": 3}, 2) = true
//     사용 사례: 값 검증, 역 조회 존재, 데이터 존재
//
//   - HasEntry(m, key, value): 키-값 쌍 존재 확인
//     시간: O(1), 공간: O(1)
//     키 존재와 값 일치 모두 확인
//     HasKey 또는 HasValue보다 구체적
//     예: HasEntry({"a": 1, "b": 2}, "b", 2) = true
//     사용 사례: 정확한 일치 검증, 상태 확인, 쌍 확인
//
// 집합 관계:
//   - IsSubset(subset, superset): 부분집합 관계 확인
//     시간: O(n) (n = len(subset)), 공간: O(1)
//     모든 subset 항목이 superset에 존재하면 true
//     빈 맵은 모든 맵의 부분집합
//     진부분집합: len(subset) < len(superset)
//     예: IsSubset({"a": 1}, {"a": 1, "b": 2}) = true
//     사용 사례: 집합 포함, 권한 확인, 부분 일치
//
//   - IsSuperset(superset, subset): 상위집합 관계 확인
//     시간: O(n) (n = len(subset)), 공간: O(1)
//     superset이 모든 subset 항목 포함 시 true
//     별칭: IsSuperset(a, b) = IsSubset(b, a)
//     모든 맵은 빈 맵의 상위집합
//     예: IsSuperset({"a": 1, "b": 2}, {"a": 1}) = true
//     사용 사례: 집합 포함, 기능 확인, 커버리지 검증
//
// 설계 원칙:
//   - 단락 평가: Every/Some/None은 가능할 때 조기 종료
//   - 공허한 진실: Every/None은 빈 맵에 true 반환
//   - 타입 안전성: HasValue/HasEntry는 comparable 값 필요
//   - 효율성: 키 확인은 O(1), 값 검색은 O(n)
//   - 일관성: 빈 맵에서 예측 가능한 동작
//
// 비교: Every vs Some vs None:
//   - Every: ∀x P(x) - 모두 일치해야 함
//   - Some: ∃x P(x) - 최소 하나 일치
//   - None: ¬∃x P(x) - 아무것도 일치 안 함
//   - 논리적 관계:
//     * None(m, fn) = !Some(m, fn)
//     * Every(m, fn) = None(m, !fn)
//     * !Every(m, fn) = Some(m, !fn)
//
// 비교: HasKey vs HasValue:
//   - HasKey: O(1) 해시 조회
//   - HasValue: O(n) 선형 검색
//   - HasKey: 더 효율적
//   - HasValue: 더 비싸지만 역 조회에 필요
//
// 비교: HasValue vs Some:
//   - HasValue: 직접 값 동등성 확인
//   - Some: 조건 기반, 더 유연
//   - HasValue: 정확한 일치에 더 단순
//   - Some: 복잡한 조건에 더 좋음
//
// 비교: IsSubset vs Equal:
//   - IsSubset: 단방향 포함
//   - Equal: 양방향 동등
//   - IsSubset(a, b) && IsSubset(b, a) = Equal(a, b)
//   - IsSubset은 superset이 추가 항목 가질 수 있음
//
// 성능 특성:
//
// 시간 복잡도:
//   - Every/Some/None: 최악 O(n)
//     * 최선: 조기 종료로 O(1)
//     * 평균: 무작위 분포로 O(n/2)
//   - HasKey/HasEntry: O(1) - 해시 테이블 조회
//   - HasValue: O(n) - 모든 값 확인 필요
//   - IsSubset/IsSuperset: O(n) (n = subset 크기)
//
// 공간 복잡도:
//   - 모든 작업: O(1) - 추가 할당 없음
//   - 읽기 전용 작업, 새 데이터 구조 없음
//
// 단락 최적화:
//   - Every: 첫 불일치 항목에서 false 반환
//   - Some: 첫 일치 항목에서 true 반환
//   - None: 첫 일치 항목에서 false 반환
//   - 실용적 속도 향상: 종종 O(n) 대신 O(1)
//
// 엣지 케이스:
//   - 빈 맵에 Every: true 반환 (공허한 진실)
//   - 빈 맵에 Some: false 반환 (일치할 요소 없음)
//   - 빈 맵에 None: true 반환 (공허한 진실)
//   - nil 맵에 HasKey: false 반환 (안전)
//   - nil 맵에 HasValue: false 반환 (안전)
//   - nil 맵에 HasEntry: false 반환 (안전)
//   - IsSubset(empty, any): true 반환 (공집합은 모든 것의 부분집합)
//   - IsSubset(any, empty): any가 비어있지 않으면 false
//   - IsSubset(m, m): true 반환 (반사적)
//   - IsSuperset(m, m): true 반환 (반사적)

// Every checks whether all entries in the map satisfy the predicate.
// Every는 맵의 모든 항목이 조건을 만족하는지 확인합니다.
//
// Returns true for empty maps.
// 빈 맵의 경우 true를 반환합니다.
//
// Time complexity: O(n) worst case, early termination on first false
// 시간 복잡도: 최악의 경우 O(n), 첫 번째 false에서 조기 종료
//
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
//
//	superset := map[string]int{"a": 1, "b": 2, "c": 3}
//	subset := map[string]int{"a": 1, "b": 2}
//	result := maputil.IsSuperset(superset, subset) // true
func IsSuperset[K comparable, V comparable](superset, subset map[K]V) bool {
	return IsSubset(subset, superset)
}
