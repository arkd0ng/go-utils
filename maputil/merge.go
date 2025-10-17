package maputil

// Package maputil/merge.go provides comprehensive map merging and set operations.
// This file contains functions for combining, intersecting, differencing, and
// deep-merging maps with various conflict resolution strategies.
//
// maputil/merge.go 패키지는 포괄적인 맵 병합 및 집합 연산을 제공합니다.
// 이 파일은 다양한 충돌 해결 전략을 통해 맵을 결합, 교차, 차집합, 깊은 병합하는 함수들을 포함합니다.
//
// # Overview | 개요
//
// The merge.go file provides 8 map combination operations organized into 3 categories:
//
// merge.go 파일은 3개 카테고리로 구성된 8개 맵 결합 작업을 제공합니다:
//
// 1. MERGING OPERATIONS | 병합 작업
//   - Merge: Simple merge with last-wins conflict resolution (O(n*m))
//   - MergeWith: Custom conflict resolution function (O(n*m))
//   - DeepMerge: Recursive merge for nested maps (O(n*m*d))
//   - Union: Alias for Merge (O(n*m))
//   - Assign: In-place mutation merge (O(n*m))
//
// 2. SET OPERATIONS | 집합 연산
//   - Intersection: Keys present in all maps (O(n*m))
//   - Difference: Keys in first map but not second (O(n))
//   - SymmetricDifference: Keys in either map but not both (O(n+m))
//
// 3. CONFLICT RESOLUTION | 충돌 해결
//   - Last-wins strategy (Merge, Union)
//   - Custom resolver (MergeWith)
//   - Recursive merge (DeepMerge)
//
// # Design Principles | 설계 원칙
//
// 1. IMMUTABILITY BY DEFAULT | 기본 불변성
//   - Most functions return new maps without modifying inputs
//   - Assign is the only mutating operation (clearly documented)
//   - Safe for concurrent read access to original maps
//
//   대부분의 함수는 입력을 수정하지 않고 새 맵을 반환합니다
//   Assign만 변경 작업입니다 (명확히 문서화됨)
//   원본 맵에 대한 동시 읽기 접근에 안전합니다
//
// 2. CONFLICT RESOLUTION STRATEGIES | 충돌 해결 전략
//   - Last-wins: Simplest, most performant (Merge, Union)
//   - Custom resolver: Maximum flexibility (MergeWith)
//   - Recursive: For nested structures (DeepMerge)
//   - Explicit control: Choose strategy based on use case
//
//   마지막 우선: 가장 단순하고 성능이 좋음 (Merge, Union)
//   커스텀 해결: 최대 유연성 (MergeWith)
//   재귀적: 중첩 구조용 (DeepMerge)
//   명시적 제어: 사용 사례에 따라 전략 선택
//
// 3. SET THEORY SEMANTICS | 집합 이론 의미론
//   - Intersection: A ∩ B ∩ C (all sets)
//   - Difference: A - B (relative complement)
//   - SymmetricDifference: A △ B (exclusive or)
//   - Union: A ∪ B ∪ C (alias for Merge)
//
//   교집합: A ∩ B ∩ C (모든 집합)
//   차집합: A - B (상대 여집합)
//   대칭차집합: A △ B (배타적 OR)
//   합집합: A ∪ B ∪ C (Merge의 별칭)
//
// 4. VARIADIC OPERATIONS | 가변 인자 연산
//   - Merge, MergeWith, Union, Intersection: Accept multiple maps
//   - Difference, SymmetricDifference: Binary operations (two maps)
//   - Sequential processing for multiple inputs
//
//   Merge, MergeWith, Union, Intersection: 여러 맵 허용
//   Difference, SymmetricDifference: 이진 연산 (두 맵)
//   여러 입력에 대한 순차 처리
//
// # Function Categories | 함수 카테고리
//
// BASIC MERGING OPERATIONS | 기본 병합 작업
//
// Merge(maps...) combines multiple maps into a single map using last-wins strategy.
// When duplicate keys exist across maps, the value from the rightmost (last) map
// takes precedence.
//
// Merge(maps...)는 마지막 우선 전략으로 여러 맵을 단일 맵으로 결합합니다.
// 맵 간 중복 키가 존재하면, 가장 오른쪽(마지막) 맵의 값이 우선합니다.
//
// Time Complexity: O(n*m) where n = number of maps, m = average map size
// Space Complexity: O(total keys) - allocates result map
// Conflict Resolution: Last-wins (rightmost value)
// Order: Processes maps left-to-right
// Nil Handling: Skips nil maps gracefully
// Immutability: Returns new map, originals unchanged
//
// 시간 복잡도: O(n*m) 여기서 n = 맵 개수, m = 평균 맵 크기
// 공간 복잡도: O(총 키 개수) - 결과 맵 할당
// 충돌 해결: 마지막 우선 (가장 오른쪽 값)
// 순서: 맵을 왼쪽에서 오른쪽으로 처리
// Nil 처리: nil 맵을 우아하게 건너뜀
// 불변성: 새 맵 반환, 원본 변경 없음
//
// Use Case: Simple configuration layering, defaults + overrides
// 사용 사례: 단순 구성 계층화, 기본값 + 재정의
//
// Example:
//   defaults := map[string]int{"timeout": 30, "retries": 3}
//   userConfig := map[string]int{"timeout": 60}
//   result := Merge(defaults, userConfig)
//   // result = {"timeout": 60, "retries": 3} (userConfig overrides)
//
// MergeWith(fn, maps...) combines maps using a custom conflict resolution function.
// The resolver receives (oldValue, newValue) and returns the value to use when
// keys conflict. This provides maximum flexibility for complex merging logic.
//
// MergeWith(fn, maps...)는 커스텀 충돌 해결 함수로 맵을 결합합니다.
// 해결자는 (기존값, 새값)을 받고 키가 충돌할 때 사용할 값을 반환합니다.
// 복잡한 병합 로직에 최대 유연성을 제공합니다.
//
// Time Complexity: O(n*m) + O(k*r) where k = conflicts, r = resolver cost
// Space Complexity: O(total keys)
// Conflict Resolution: Custom function (oldVal, newVal) -> resolvedVal
// Order: Processes maps left-to-right, calling resolver on conflicts
// Resolver Signature: func(V, V) V
// Immutability: Returns new map, originals unchanged
// Nil Handling: Skips nil maps
//
// 시간 복잡도: O(n*m) + O(k*r) 여기서 k = 충돌 개수, r = 해결자 비용
// 공간 복잡도: O(총 키 개수)
// 충돌 해결: 커스텀 함수 (기존값, 새값) -> 해결값
// 순서: 맵을 왼쪽에서 오른쪽으로 처리, 충돌 시 해결자 호출
// 해결자 시그니처: func(V, V) V
// 불변성: 새 맵 반환, 원본 변경 없음
// Nil 처리: nil 맵 건너뜀
//
// Use Case: Summing values, custom priority logic, conditional merging
// 사용 사례: 값 합산, 커스텀 우선순위 로직, 조건부 병합
//
// Example:
//   sales1 := map[string]int{"alice": 100, "bob": 150}
//   sales2 := map[string]int{"bob": 50, "charlie": 200}
//   total := MergeWith(func(old, new int) int { return old + new }, sales1, sales2)
//   // total = {"alice": 100, "bob": 200, "charlie": 200} (summed)
//
// DeepMerge(maps...) performs recursive merging for nested map structures.
// When both values at a key are maps, they are recursively merged. Otherwise,
// last-wins strategy applies. Only works with map[string]interface{} type.
//
// DeepMerge(maps...)는 중첩된 맵 구조에 대해 재귀적 병합을 수행합니다.
// 키의 두 값이 모두 맵이면, 재귀적으로 병합됩니다. 그렇지 않으면 마지막 우선 전략이 적용됩니다.
// map[string]interface{} 타입에서만 작동합니다.
//
// Time Complexity: O(n*m*d) where d = average nesting depth
// Space Complexity: O(total keys * depth)
// Conflict Resolution: Recursive for nested maps, last-wins for values
// Type Constraint: Only map[string]interface{} (no generics)
// Recursion: Handles arbitrary nesting depth
// Immutability: Returns new nested structure
// Nil Handling: Skips nil maps at all levels
//
// 시간 복잡도: O(n*m*d) 여기서 d = 평균 중첩 깊이
// 공간 복잡도: O(총 키 개수 * 깊이)
// 충돌 해결: 중첩 맵은 재귀, 값은 마지막 우선
// 타입 제약: map[string]interface{}만 (제네릭 없음)
// 재귀: 임의의 중첩 깊이 처리
// 불변성: 새 중첩 구조 반환
// Nil 처리: 모든 레벨에서 nil 맵 건너뜀
//
// Use Case: Configuration files, JSON merging, nested settings
// 사용 사례: 구성 파일, JSON 병합, 중첩 설정
//
// Example:
//   config1 := map[string]interface{}{
//       "database": map[string]interface{}{"host": "localhost", "port": 3306},
//   }
//   config2 := map[string]interface{}{
//       "database": map[string]interface{}{"port": 5432, "user": "admin"},
//   }
//   merged := DeepMerge(config1, config2)
//   // merged = {"database": {"host": "localhost", "port": 5432, "user": "admin"}}
//
// Union(maps...) is a semantic alias for Merge, emphasizing set union interpretation.
// Functionally identical to Merge, but name conveys set theory semantics (A ∪ B).
//
// Union(maps...)는 집합 합집합 해석을 강조하는 Merge의 의미론적 별칭입니다.
// 기능적으로 Merge와 동일하지만, 이름이 집합 이론 의미론(A ∪ B)을 전달합니다.
//
// Time Complexity: O(n*m) (same as Merge)
// Space Complexity: O(total keys)
// Conflict Resolution: Last-wins
// Semantics: Set union (∪) operation
// Immutability: Returns new map
//
// 시간 복잡도: O(n*m) (Merge와 동일)
// 공간 복잡도: O(총 키 개수)
// 충돌 해결: 마지막 우선
// 의미론: 집합 합집합 (∪) 연산
// 불변성: 새 맵 반환
//
// Use Case: When expressing set union semantics
// 사용 사례: 집합 합집합 의미론을 표현할 때
//
// Assign(target, sources...) is the ONLY mutating merge operation.
// It modifies the target map in-place by copying all entries from source maps.
// WARNING: This breaks immutability and can cause concurrent access issues.
//
// Assign(target, sources...)는 유일한 변경 병합 작업입니다.
// 소스 맵의 모든 항목을 복사하여 대상 맵을 제자리에서 수정합니다.
// 경고: 불변성을 깨트리고 동시 접근 문제를 일으킬 수 있습니다.
//
// Time Complexity: O(n*m)
// Space Complexity: O(1) (modifies in-place, no new map)
// MUTATING OPERATION: Directly modifies target map
// Return Value: Returns target map (for chaining)
// Conflict Resolution: Last-wins
// Concurrent Safety: NOT SAFE - requires external synchronization
// Nil Handling: Panics if target is nil (cannot mutate nil)
//
// 시간 복잡도: O(n*m)
// 공간 복잡도: O(1) (제자리 수정, 새 맵 없음)
// 변경 작업: 대상 맵을 직접 수정
// 반환값: 대상 맵 반환 (체이닝용)
// 충돌 해결: 마지막 우선
// 동시성 안전성: 안전하지 않음 - 외부 동기화 필요
// Nil 처리: target이 nil이면 패닉 (nil 변경 불가)
//
// Use Case: Performance-critical in-place updates (use with caution)
// 사용 사례: 성능 중요한 제자리 업데이트 (주의해서 사용)
//
// Example:
//   config := map[string]int{"a": 1}
//   overrides := map[string]int{"b": 2}
//   Assign(config, overrides) // config is now {"a": 1, "b": 2}
//
// SET OPERATIONS | 집합 연산
//
// Intersection(maps...) returns keys present in ALL input maps (∩ operation).
// The result contains only keys that exist in every single input map. Values
// are taken from the first map.
//
// Intersection(maps...)는 모든 입력 맵에 존재하는 키를 반환합니다 (∩ 연산).
// 결과는 모든 입력 맵에 존재하는 키만 포함합니다. 값은 첫 번째 맵에서 가져옵니다.
//
// Time Complexity: O(n*m) where n = number of maps, m = average map size
// Space Complexity: O(k) where k = number of common keys
// Set Operation: A ∩ B ∩ C (intersection of all sets)
// Value Source: First map (maps[0])
// Empty Input: Returns empty map
// Single Map: Returns clone of that map
// Immutability: Returns new map
// Nil Handling: Treats nil maps as empty
//
// 시간 복잡도: O(n*m) 여기서 n = 맵 개수, m = 평균 맵 크기
// 공간 복잡도: O(k) 여기서 k = 공통 키 개수
// 집합 연산: A ∩ B ∩ C (모든 집합의 교집합)
// 값 소스: 첫 번째 맵 (maps[0])
// 빈 입력: 빈 맵 반환
// 단일 맵: 해당 맵의 복제본 반환
// 불변성: 새 맵 반환
// Nil 처리: nil 맵을 빈 것으로 처리
//
// Use Case: Find common configuration keys, shared permissions
// 사용 사례: 공통 구성 키 찾기, 공유 권한
//
// Example:
//   users1 := map[string]bool{"alice": true, "bob": true, "charlie": true}
//   users2 := map[string]bool{"bob": true, "charlie": true, "dave": true}
//   users3 := map[string]bool{"charlie": true, "dave": true}
//   common := Intersection(users1, users2, users3)
//   // common = {"charlie": true} (only charlie in all three)
//
// Difference(m1, m2) returns keys in m1 that are NOT in m2 (A - B operation).
// This is the relative complement: elements in first set but not in second.
//
// Difference(m1, m2)는 m1에는 있지만 m2에는 없는 키를 반환합니다 (A - B 연산).
// 상대 여집합입니다: 첫 번째 집합에 있지만 두 번째에는 없는 원소.
//
// Time Complexity: O(n) where n = size of m1
// Space Complexity: O(k) where k = keys only in m1
// Set Operation: A - B (relative complement)
// Binary Operation: Takes exactly two maps
// Asymmetric: Difference(A, B) ≠ Difference(B, A)
// Value Source: First map (m1)
// Immutability: Returns new map
// Nil Handling: Treats nil as empty
//
// 시간 복잡도: O(n) 여기서 n = m1 크기
// 공간 복잡도: O(k) 여기서 k = m1에만 있는 키
// 집합 연산: A - B (상대 여집합)
// 이진 연산: 정확히 두 맵 사용
// 비대칭: Difference(A, B) ≠ Difference(B, A)
// 값 소스: 첫 번째 맵 (m1)
// 불변성: 새 맵 반환
// Nil 처리: nil을 빈 것으로 처리
//
// Use Case: Find removed items, detect deletions, filter exclusions
// 사용 사례: 제거된 항목 찾기, 삭제 감지, 필터 제외
//
// Example:
//   oldConfig := map[string]int{"timeout": 30, "retries": 3, "cache": 100}
//   newConfig := map[string]int{"timeout": 60, "retries": 5}
//   removed := Difference(oldConfig, newConfig)
//   // removed = {"cache": 100} (was in old but not new)
//
// SymmetricDifference(m1, m2) returns keys in EITHER map but NOT BOTH (A △ B).
// This is the exclusive or (XOR) operation in set theory: (A - B) ∪ (B - A).
//
// SymmetricDifference(m1, m2)는 어느 한 맵에는 있지만 둘 다에는 없는 키를 반환합니다 (A △ B).
// 집합 이론의 배타적 OR (XOR) 연산입니다: (A - B) ∪ (B - A).
//
// Time Complexity: O(n+m) where n = size of m1, m = size of m2
// Space Complexity: O(k) where k = non-overlapping keys
// Set Operation: A △ B (symmetric difference / XOR)
// Binary Operation: Takes exactly two maps
// Symmetric: SymmetricDifference(A, B) = SymmetricDifference(B, A)
// Value Source: From respective maps (m1 or m2)
// Immutability: Returns new map
// Nil Handling: Treats nil as empty
//
// 시간 복잡도: O(n+m) 여기서 n = m1 크기, m = m2 크기
// 공간 복잡도: O(k) 여기서 k = 겹치지 않는 키
// 집합 연산: A △ B (대칭차집합 / XOR)
// 이진 연산: 정확히 두 맵 사용
// 대칭: SymmetricDifference(A, B) = SymmetricDifference(B, A)
// 값 소스: 각 맵에서 (m1 또는 m2)
// 불변성: 새 맵 반환
// Nil 처리: nil을 빈 것으로 처리
//
// Use Case: Find differences between two states, change detection
// 사용 사례: 두 상태 간 차이 찾기, 변경 감지
//
// Example:
//   state1 := map[string]int{"a": 1, "b": 2, "c": 3}
//   state2 := map[string]int{"b": 20, "c": 30, "d": 4}
//   changes := SymmetricDifference(state1, state2)
//   // changes = {"a": 1, "d": 4} (unique to each side)
//
// # Comparisons with Related Functions | 관련 함수와 비교
//
// Merge vs. MergeWith:
//   - Merge: Simple last-wins, faster, no custom logic
//   - MergeWith: Custom resolver, more flexible, slightly slower
//   - Use Merge for simple overrides
//   - Use MergeWith for summing, priority rules, conditional merging
//
// Merge 대 MergeWith:
//   - Merge: 단순 마지막 우선, 더 빠름, 커스텀 로직 없음
//   - MergeWith: 커스텀 해결자, 더 유연, 약간 느림
//   - 단순 재정의에 Merge 사용
//   - 합산, 우선순위 규칙, 조건부 병합에 MergeWith 사용
//
// Merge vs. DeepMerge:
//   - Merge: Flat maps, O(n*m), all value types
//   - DeepMerge: Nested maps, O(n*m*d), only map[string]interface{}
//   - Use Merge for simple flat structures
//   - Use DeepMerge for configuration files, JSON merging
//
// Merge 대 DeepMerge:
//   - Merge: 평면 맵, O(n*m), 모든 값 타입
//   - DeepMerge: 중첩 맵, O(n*m*d), map[string]interface{}만
//   - 단순 평면 구조에 Merge 사용
//   - 구성 파일, JSON 병합에 DeepMerge 사용
//
// Merge vs. Assign:
//   - Merge: Immutable, returns new map, safe for concurrency
//   - Assign: Mutable, modifies target, NOT safe for concurrency
//   - Use Merge by default (immutability is safer)
//   - Use Assign only when performance critical and safe
//
// Merge 대 Assign:
//   - Merge: 불변, 새 맵 반환, 동시성에 안전
//   - Assign: 가변, 대상 수정, 동시성에 안전하지 않음
//   - 기본적으로 Merge 사용 (불변성이 더 안전)
//   - 성능 중요하고 안전한 경우에만 Assign 사용
//
// Union vs. Merge:
//   - Functionally identical (Union is alias)
//   - Union emphasizes set theory semantics
//   - Merge emphasizes data combination semantics
//   - Choose based on semantic clarity for your use case
//
// Union 대 Merge:
//   - 기능적으로 동일 (Union은 별칭)
//   - Union은 집합 이론 의미론 강조
//   - Merge는 데이터 결합 의미론 강조
//   - 사용 사례의 의미적 명확성에 따라 선택
//
// Intersection vs. Filter:
//   - Intersection: Multiple maps, finds common keys
//   - Filter: Single map, predicate-based filtering
//   - Intersection for multi-map common elements
//   - Filter for single-map conditional selection
//
// Intersection 대 Filter:
//   - Intersection: 여러 맵, 공통 키 찾기
//   - Filter: 단일 맵, 조건 기반 필터링
//   - 다중 맵 공통 요소에 Intersection
//   - 단일 맵 조건 선택에 Filter
//
// Difference vs. Omit:
//   - Difference: Two maps, A - B operation
//   - Omit: Single map + key list, removes specific keys
//   - Difference for map-based exclusion
//   - Omit for explicit key list exclusion
//
// Difference 대 Omit:
//   - Difference: 두 맵, A - B 연산
//   - Omit: 단일 맵 + 키 목록, 특정 키 제거
//   - 맵 기반 제외에 Difference
//   - 명시적 키 목록 제외에 Omit
//
// SymmetricDifference vs. Union + Intersection:
//   - SymmetricDifference: Direct XOR, single operation
//   - Union - Intersection: Two operations, same result
//   - SymmetricDifference more efficient
//   - Formula: SymDiff(A,B) = Union(A,B) - Intersection(A,B)
//
// SymmetricDifference 대 Union + Intersection:
//   - SymmetricDifference: 직접 XOR, 단일 연산
//   - Union - Intersection: 두 연산, 동일 결과
//   - SymmetricDifference가 더 효율적
//   - 공식: SymDiff(A,B) = Union(A,B) - Intersection(A,B)
//
// # Performance Characteristics | 성능 특성
//
// Time Complexities:
//   - O(n*m): Merge, MergeWith, Union, Intersection, Assign
//   - O(n): Difference
//   - O(n+m): SymmetricDifference
//   - O(n*m*d): DeepMerge (d = nesting depth)
//
// 시간 복잡도:
//   - O(n*m): Merge, MergeWith, Union, Intersection, Assign
//   - O(n): Difference
//   - O(n+m): SymmetricDifference
//   - O(n*m*d): DeepMerge (d = 중첩 깊이)
//
// Space Complexities:
//   - O(1): Assign (in-place mutation)
//   - O(k): Intersection, Difference, SymmetricDifference (k = result keys)
//   - O(total keys): Merge, MergeWith, Union
//   - O(total keys * depth): DeepMerge
//
// 공간 복잡도:
//   - O(1): Assign (제자리 변경)
//   - O(k): Intersection, Difference, SymmetricDifference (k = 결과 키)
//   - O(총 키 개수): Merge, MergeWith, Union
//   - O(총 키 개수 * 깊이): DeepMerge
//
// Optimization Tips:
//   - Pre-allocate result maps when total size known
//   - Use Difference instead of Filter when excluding another map's keys
//   - Use Assign only when mutation is safe and performance critical
//   - For multiple Difference calls, consider caching Intersection results
//   - DeepMerge is expensive; avoid on deep structures if possible
//
// 최적화 팁:
//   - 총 크기가 알려진 경우 결과 맵을 미리 할당
//   - 다른 맵의 키를 제외할 때 Filter 대신 Difference 사용
//   - 변경이 안전하고 성능이 중요할 때만 Assign 사용
//   - 여러 Difference 호출에 대해 Intersection 결과 캐싱 고려
//   - DeepMerge는 비용이 높음; 가능하면 깊은 구조 피하기
//
// # Common Usage Patterns | 일반적인 사용 패턴
//
// 1. Configuration Layering | 구성 계층화
//
//	defaults := map[string]int{"timeout": 30, "retries": 3, "cache": 100}
//	envConfig := map[string]int{"timeout": 60}
//	userConfig := map[string]int{"cache": 200}
//	final := maputil.Merge(defaults, envConfig, userConfig)
//	// final = {"timeout": 60, "retries": 3, "cache": 200}
//	// Priority: defaults < env < user
//
// 2. Summing Sales Data | 판매 데이터 합산
//
//	sales1 := map[string]int{"alice": 1000, "bob": 1500}
//	sales2 := map[string]int{"bob": 500, "charlie": 2000}
//	sales3 := map[string]int{"alice": 500, "charlie": 300}
//	total := maputil.MergeWith(
//	    func(old, new int) int { return old + new },
//	    sales1, sales2, sales3,
//	)
//	// total = {"alice": 1500, "bob": 2000, "charlie": 2300}
//
// 3. Deep Configuration Merge | 깊은 구성 병합
//
//	base := map[string]interface{}{
//	    "database": map[string]interface{}{
//	        "host": "localhost",
//	        "port": 3306,
//	        "timeout": 30,
//	    },
//	}
//	override := map[string]interface{}{
//	    "database": map[string]interface{}{
//	        "port": 5432,
//	        "user": "admin",
//	    },
//	}
//	config := maputil.DeepMerge(base, override)
//	// config = {"database": {"host": "localhost", "port": 5432, "timeout": 30, "user": "admin"}}
//
// 4. Find Common Permissions | 공통 권한 찾기
//
//	adminPerms := map[string]bool{"read": true, "write": true, "delete": true, "admin": true}
//	moderatorPerms := map[string]bool{"read": true, "write": true, "delete": true}
//	userPerms := map[string]bool{"read": true, "write": true}
//	commonToAll := maputil.Intersection(adminPerms, moderatorPerms, userPerms)
//	// commonToAll = {"read": true, "write": true}
//
// 5. Detect Removed Configuration Keys | 제거된 구성 키 감지
//
//	oldConfig := map[string]string{
//	    "api_key": "secret", "timeout": "30", "deprecated_option": "value",
//	}
//	newConfig := map[string]string{
//	    "api_key": "secret", "timeout": "60",
//	}
//	removed := maputil.Difference(oldConfig, newConfig)
//	// removed = {"deprecated_option": "value"}
//	fmt.Printf("Removed keys: %v\n", maputil.KeysSlice(removed))
//
// 6. Find Changed Keys | 변경된 키 찾기
//
//	state1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	state2 := map[string]int{"a": 1, "b": 20, "c": 3, "e": 5}
//	changed := maputil.SymmetricDifference(state1, state2)
//	// changed = {"d": 4, "e": 5} (keys unique to each)
//	// Note: This doesn't detect value changes for common keys
//
// 7. In-Place Update (Performance Critical) | 제자리 업데이트 (성능 중요)
//
//	cache := make(map[string]string, 1000)
//	// Populate cache...
//	updates := map[string]string{"key1": "newval1", "key2": "newval2"}
//	maputil.Assign(cache, updates) // Mutates cache directly
//	// Use with caution: not safe for concurrent access
//
// 8. Multi-Layer Merge with Priority | 우선순위가 있는 다중 계층 병합
//
//	defaultSettings := map[string]int{"volume": 50, "brightness": 70, "contrast": 50}
//	systemSettings := map[string]int{"brightness": 80}
//	userSettings := map[string]int{"volume": 80, "brightness": 90}
//	finalSettings := maputil.Merge(defaultSettings, systemSettings, userSettings)
//	// finalSettings = {"volume": 80, "brightness": 90, "contrast": 50}
//	// Order matters: later maps override earlier ones
//
// 9. Custom Conflict Resolution | 커스텀 충돌 해결
//
//	inventory1 := map[string]int{"apples": 10, "bananas": 5}
//	inventory2 := map[string]int{"bananas": 3, "oranges": 7}
//	combined := maputil.MergeWith(
//	    func(old, new int) int {
//	        if old > new {
//	            return old // Keep higher inventory
//	        }
//	        return new
//	    },
//	    inventory1, inventory2,
//	)
//	// combined = {"apples": 10, "bananas": 5, "oranges": 7}
//
// 10. Set Operations Chain | 집합 연산 체인
//
//	setA := map[string]bool{"a": true, "b": true, "c": true}
//	setB := map[string]bool{"b": true, "c": true, "d": true}
//	setC := map[string]bool{"c": true, "d": true, "e": true}
//
//	// Union of A, B, C
//	union := maputil.Union(setA, setB, setC)
//	// union = {"a": true, "b": true, "c": true, "d": true, "e": true}
//
//	// Intersection of A, B, C
//	intersection := maputil.Intersection(setA, setB, setC)
//	// intersection = {"c": true}
//
//	// A - B (elements in A but not B)
//	diff := maputil.Difference(setA, setB)
//	// diff = {"a": true}
//
//	// A △ B (symmetric difference)
//	symDiff := maputil.SymmetricDifference(setA, setB)
//	// symDiff = {"a": true, "d": true}
//
// # Edge Cases and Nil Handling | 엣지 케이스와 Nil 처리
//
// Empty Maps:
//   - Merge, MergeWith, Union: Return empty map if all inputs empty
//   - Intersection: Returns empty map (no common keys)
//   - Difference: Returns empty map if first map empty
//   - SymmetricDifference: Returns union if no overlap
//
// 빈 맵:
//   - Merge, MergeWith, Union: 모든 입력이 비었으면 빈 맵 반환
//   - Intersection: 빈 맵 반환 (공통 키 없음)
//   - Difference: 첫 맵이 비었으면 빈 맵 반환
//   - SymmetricDifference: 겹침이 없으면 합집합 반환
//
// Nil Maps:
//   - All functions treat nil as empty maps
//   - Exception: Assign panics if target is nil (cannot mutate nil)
//   - No nil pointer dereferences in read operations
//
// Nil 맵:
//   - 모든 함수가 nil을 빈 맵으로 처리
//   - 예외: Assign은 target이 nil이면 패닉 (nil 변경 불가)
//   - 읽기 작업에서 nil 포인터 역참조 없음
//
// Single Map:
//   - Merge, Union: Returns clone of the single map
//   - Intersection: Returns clone of the single map
//   - MergeWith: Resolver not called (no conflicts)
//
// 단일 맵:
//   - Merge, Union: 단일 맵의 복제본 반환
//   - Intersection: 단일 맵의 복제본 반환
//   - MergeWith: 해결자가 호출되지 않음 (충돌 없음)
//
// No Conflicts:
//   - MergeWith: Resolver not called if keys don't overlap
//   - Merge, MergeWith, Union: Equivalent when keys distinct
//
// 충돌 없음:
//   - MergeWith: 키가 겹치지 않으면 해결자 호출 안 됨
//   - Merge, MergeWith, Union: 키가 구별되면 동등
//
// Complete Overlap:
//   - Intersection: Returns first map's values for all keys
//   - Difference: Returns empty map if all keys in second map
//
// 완전 겹침:
//   - Intersection: 모든 키에 대해 첫 맵의 값 반환
//   - Difference: 모든 키가 두 번째 맵에 있으면 빈 맵 반환
//
// # Thread Safety | 스레드 안전성
//
// Immutable Operations (Safe for Concurrent Reads):
//   - Merge, MergeWith, DeepMerge, Union
//   - Intersection, Difference, SymmetricDifference
//   - All create new maps, don't modify originals
//   - Safe when input maps have concurrent readers
//
// 불변 작업 (동시 읽기 안전):
//   - Merge, MergeWith, DeepMerge, Union
//   - Intersection, Difference, SymmetricDifference
//   - 모두 새 맵 생성, 원본 수정 없음
//   - 입력 맵에 동시 읽기가 있을 때 안전
//
// Mutable Operation (NOT Safe):
//   - Assign: Modifies target map in-place
//   - Requires external synchronization (mutex) for concurrent access
//   - Not safe when other goroutines read or write target
//
// 가변 작업 (안전하지 않음):
//   - Assign: 대상 맵을 제자리에서 수정
//   - 동시 접근에 외부 동기화 (뮤텍스) 필요
//   - 다른 고루틴이 대상을 읽거나 쓸 때 안전하지 않음
//
// Concurrent Modification Warning:
//   - Do not modify input maps while merge operations read them
//   - Do not modify target map while Assign is running
//   - Use sync.RWMutex for concurrent map access
//
// 동시 수정 경고:
//   - 병합 작업이 읽는 동안 입력 맵 수정 금지
//   - Assign이 실행되는 동안 대상 맵 수정 금지
//   - 동시 맵 접근에 sync.RWMutex 사용
//
// # See Also | 참고
//
// Related files in maputil package:
//   - basic.go: Fundamental operations (Clone, Equal used internally)
//   - filter.go: Filtering operations (Pick, Omit for key selection)
//   - comparison.go: Equality and difference detection
//   - nested.go: Nested map navigation (complements DeepMerge)
//   - transform.go: Map transformation (MapKeys, MapValues for post-merge)
//
// maputil 패키지의 관련 파일:
//   - basic.go: 기본 작업 (내부적으로 사용되는 Clone, Equal)
//   - filter.go: 필터링 작업 (키 선택용 Pick, Omit)
//   - comparison.go: 동등성 및 차이 감지
//   - nested.go: 중첩 맵 탐색 (DeepMerge 보완)
//   - transform.go: 맵 변환 (병합 후 MapKeys, MapValues)

// Merge merges multiple maps into a single map.
// Merge는 여러 맵을 단일 맵으로 병합합니다.
//
// If duplicate keys exist, the value from the last map wins.
// 중복 키가 있으면 마지막 맵의 값이 우선합니다.
//
// Time complexity: O(n*m) where n is the number of maps and m is the average map size
// 시간 복잡도: O(n*m) 여기서 n은 맵 개수, m은 평균 맵 크기
//
// Example
// 예제:
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
// Example
// 예제:
//
//	m1 := map[string]int{"a": 1, "b": 2}
//	m2 := map[string]int{"b": 3, "c": 4}
//
// result := maputil.MergeWith(
// func(old, new int) int { return old + new }, // Sum on conflict
// 충돌 시 합산
//
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
//
//	target := map[string]int{"a": 1, "b": 2}
//
// source := map[string]int{"b": 3, "c": 4}
// result := maputil.Assign(target, source) // target is modified
// target이 수정됨
//
//	// result = map[string]int{"a": 1, "b": 3, "c": 4}
func Assign[K comparable, V any](target map[K]V, sources ...map[K]V) map[K]V {
	for _, source := range sources {
		for k, v := range source {
			target[k] = v
		}
	}
	return target
}
