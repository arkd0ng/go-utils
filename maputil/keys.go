package maputil

import "sort"

// keys.go provides key-focused map operations for Go.
//
// This file implements functions that extract, transform, or manipulate map keys.
// Operations include key extraction, sorting, renaming, searching, and transformation.
// All operations are immutable and return new data structures.
//
// Key Operations:
//
// Key Extraction:
//   - KeysSlice(m): Extract all keys as slice
//     Time: O(n), Space: O(n)
//     Alias for Keys from convert.go
//     Order is non-deterministic (map iteration)
//     Example: KeysSlice({"a": 1, "b": 2, "c": 3}) = ["a", "b", "c"] (order varies)
//     Use cases: Key enumeration, iteration, set operations
//
//   - KeysSorted(m): Extract keys in sorted order
//     Time: O(n log n), Space: O(n)
//     Requires Ordered constraint (comparable + <)
//     Deterministic output order
//     Example: KeysSorted({"c": 3, "a": 1, "b": 2}) = ["a", "b", "c"]
//     Use cases: Deterministic output, ordered processing, display
//
//   - KeysBy(m, fn): Extract keys matching predicate
//     Time: O(n), Space: O(k) where k = matching keys
//     Function receives key and value
//     Returns slice of matching keys
//     Example: KeysBy({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 1 }) = ["b", "c"]
//     Use cases: Filtered key extraction, conditional selection, search
//
// Key Searching:
//   - FindKey(m, fn): Find first key matching predicate
//     Time: O(n) worst case, early termination
//     Returns (key, true) if found, (zero, false) otherwise
//     Non-deterministic due to map iteration order
//     Example: FindKey({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 2 }) = ("c", true)
//     Use cases: Key lookup by condition, existence check, search
//
//   - FindKeys(m, fn): Find all keys matching predicate
//     Time: O(n), Space: O(k) where k = matches
//     Alias for KeysBy (semantic clarity)
//     Returns all matching keys
//     Example: FindKeys({"a": 1, "b": 2, "c": 3, "d": 4}, func(k, v) { return v > 2 }) = ["c", "d"]
//     Use cases: Multi-match search, filtering, batch operations
//
// Key Renaming:
//   - RenameKey(m, oldKey, newKey): Rename single key
//     Time: O(n), Space: O(n)
//     Returns new map with renamed key
//     If oldKey missing: returns clone
//     If newKey exists: overwrites
//     Example: RenameKey({"a": 1, "b": 2}, "b", "B") = {"a": 1, "B": 2}
//     Use cases: Key normalization, schema migration, field renaming
//
//   - RenameKeys(m, mapping): Rename multiple keys
//     Time: O(n), Space: O(n)
//     mapping is map[OldKey]NewKey
//     Keys not in mapping unchanged
//     Example: RenameKeys({"a": 1, "b": 2}, {"a": "A"}) = {"A": 1, "b": 2}
//     Use cases: Batch renaming, schema transformation, mapping application
//
//   - SwapKeys(m, key1, key2): Swap two key values
//     Time: O(n), Space: O(n)
//     Exchanges values at two keys
//     If either key missing: returns clone
//     Example: SwapKeys({"a": 1, "b": 2}, "a", "b") = {"a": 2, "b": 1}
//     Use cases: Value exchange, reordering, data manipulation
//
// Key Transformation:
//   - TransformKeys(m, fn): Transform all keys
//     Time: O(n), Space: O(n)
//     Function receives key, returns new key
//     Duplicate keys: last value wins
//     Example: TransformKeys({"a": 1, "b": 2}, func(k) { return strings.ToUpper(k) }) = {"A": 1, "B": 2}
//     Use cases: Case normalization, key formatting, bulk transformation
//
//   - PrefixKeys(m, prefix): Add prefix to all keys
//     Time: O(n), Space: O(n)
//     String-specific operation
//     Concatenates prefix to each key
//     Example: PrefixKeys({"a": 1, "b": 2}, "key_") = {"key_a": 1, "key_b": 2}
//     Use cases: Namespacing, key scoping, prefix addition
//
//   - SuffixKeys(m, suffix): Add suffix to all keys
//     Time: O(n), Space: O(n)
//     String-specific operation
//     Concatenates suffix to each key
//     Example: SuffixKeys({"a": 1, "b": 2}, "_key") = {"a_key": 1, "b_key": 2}
//     Use cases: Namespacing, key scoping, suffix addition
//
// Design Principles:
//   - Immutability: All operations return new maps/slices
//   - Predictability: Sorted operations provide deterministic order
//   - Flexibility: Multiple renaming and transformation options
//   - Safety: Non-existent key operations return clones safely
//   - Efficiency: O(n) operations where possible, O(n log n) for sorting
//
// Comparison: KeysSlice vs KeysSorted:
//   - KeysSlice: O(n), non-deterministic order
//   - KeysSorted: O(n log n), deterministic order
//   - KeysSlice: Faster, use when order doesn't matter
//   - KeysSorted: Slower, use when consistent order needed
//
// Comparison: KeysBy vs FindKeys:
//   - Functionally identical (FindKeys is alias)
//   - KeysBy: Emphasizes extraction
//   - FindKeys: Emphasizes search semantics
//   - Use based on semantic clarity in context
//
// Comparison: RenameKey vs TransformKeys:
//   - RenameKey: Single key, specific old→new
//   - TransformKeys: All keys, function-based
//   - RenameKey: Simpler for one key
//   - TransformKeys: Better for bulk transformations
//
// Comparison: PrefixKeys/SuffixKeys vs TransformKeys:
//   - Prefix/Suffix: Specialized for strings
//   - TransformKeys: General, any type
//   - Prefix/Suffix: More readable for string operations
//   - TransformKeys: More flexible, custom logic
//
// Performance Characteristics:
//
// Time Complexity:
//   - KeysSlice/KeysBy/FindKeys: O(n) - Single iteration
//   - KeysSorted: O(n log n) - Extraction + sorting
//   - FindKey: O(n) worst case, O(1) best with early termination
//   - RenameKey/RenameKeys/SwapKeys: O(n) - Full map copy
//   - TransformKeys/PrefixKeys/SuffixKeys: O(n) - Single iteration
//
// Space Complexity:
//   - All extraction operations: O(n) or O(k) where k = matches
//   - All map-returning operations: O(n) - New map
//   - KeysSorted: O(n) - Slice allocation + sorting
//
// Memory Allocation:
//   - Extraction: Allocates slice
//   - Renaming/Transformation: Allocates new map
//   - FindKey: O(1) - No allocation (returns existing key)
//   - KeysSorted: Allocates slice + sorting overhead
//
// Common Usage Patterns:
//
//	// Get sorted keys for deterministic iteration
//	m := map[string]int{"c": 3, "a": 1, "b": 2}
//	sortedKeys := maputil.KeysSorted(m)
//	for _, key := range sortedKeys {
//	    fmt.Printf("%s: %d\n", key, m[key])
//	}
//	// Output: a: 1, b: 2, c: 3 (always same order)
//
//	// Find keys with high values
//	scores := map[string]int{"Alice": 85, "Bob": 92, "Charlie": 78, "Diana": 95}
//	topScorers := maputil.KeysBy(scores, func(k string, v int) bool {
//	    return v >= 90
//	}) // ["Bob", "Diana"] (order may vary)
//
//	// Rename key for consistency
//	config := map[string]string{"hostname": "localhost", "Port": "8080"}
//	normalized := maputil.RenameKey(config, "hostname", "host")
//	// {"host": "localhost", "Port": "8080"}
//
//	// Batch rename for schema migration
//	oldSchema := map[string]interface{}{"user_name": "Alice", "user_age": 30}
//	mapping := map[string]string{"user_name": "name", "user_age": "age"}
//	newSchema := maputil.RenameKeys(oldSchema, mapping)
//	// {"name": "Alice", "age": 30}
//
//	// Normalize all keys to uppercase
//	data := map[string]int{"a": 1, "b": 2, "c": 3}
//	uppercased := maputil.TransformKeys(data, func(k string) string {
//	    return strings.ToUpper(k)
//	}) // {"A": 1, "B": 2, "C": 3}
//
//	// Add namespace prefix
//	settings := map[string]string{"timeout": "30", "retries": "3"}
//	prefixed := maputil.PrefixKeys(settings, "app.")
//	// {"app.timeout": "30", "app.retries": "3"}
//
// Advanced Key Patterns:
//
//	// Find first key with specific property
//	users := map[string]User{
//	    "u1": {Name: "Alice", Role: "admin"},
//	    "u2": {Name: "Bob", Role: "user"},
//	    "u3": {Name: "Charlie", Role: "admin"},
//	}
//	adminKey, found := maputil.FindKey(users, func(k string, u User) bool {
//	    return u.Role == "admin"
//	})
//	// adminKey = "u1" or "u3" (non-deterministic)
//
//	// Swap primary and secondary
//	endpoints := map[string]string{"primary": "server1", "secondary": "server2"}
//	failover := maputil.SwapKeys(endpoints, "primary", "secondary")
//	// {"primary": "server2", "secondary": "server1"}
//
//	// Chain transformations
//	m := map[string]int{"old_a": 1, "old_b": 2}
//	result := maputil.PrefixKeys(
//	    maputil.TransformKeys(m, func(k string) string {
//	        return strings.TrimPrefix(k, "old_")
//	    }),
//	    "new_",
//	)
//	// {"new_a": 1, "new_b": 2}
//
// Edge Cases:
//   - KeysSlice on empty map: Returns empty slice
//   - KeysSorted on empty map: Returns empty slice
//   - KeysBy with no matches: Returns empty slice
//   - FindKey with no match: Returns (zero, false)
//   - FindKeys with no matches: Returns empty slice
//   - RenameKey with missing oldKey: Returns clone
//   - RenameKey with existing newKey: Overwrites
//   - RenameKeys with empty mapping: Returns clone
//   - SwapKeys with missing key: Returns clone
//   - TransformKeys with duplicate results: Last value wins
//   - PrefixKeys with empty prefix: Returns clone
//   - SuffixKeys with empty suffix: Returns clone
//
// Duplicate Key Handling:
//   - TransformKeys: If fn produces same key for multiple inputs, last wins
//   - RenameKeys: Later mappings may overwrite earlier results
//   - Map iteration order is random, so "last" is non-deterministic
//   - Use carefully when transformations may produce duplicates
//
// Nil Map Behavior:
//   - KeysSlice(nil): Returns empty slice
//   - KeysSorted(nil): Returns empty slice
//   - KeysBy(nil, fn): Returns empty slice
//   - FindKey(nil, fn): Returns (zero, false)
//   - FindKeys(nil, fn): Returns empty slice
//   - RenameKey(nil, k1, k2): Returns nil
//   - RenameKeys(nil, m): Returns empty map
//   - SwapKeys(nil, k1, k2): Returns nil
//   - TransformKeys(nil, fn): Returns empty map
//   - PrefixKeys(nil, p): Returns empty map
//   - SuffixKeys(nil, s): Returns empty map
//
// Thread Safety:
//   - All operations are read-only on input map (safe for concurrent reads)
//   - Not safe if input map is modified concurrently
//   - Returned maps/slices are independent (safe for different goroutines)
//   - Transformation functions must be goroutine-safe
//
// Performance Tips:
//   - Use KeysSlice over KeysSorted when order doesn't matter
//   - FindKey is faster than FindKeys when only one result needed
//   - RenameKeys is more efficient than multiple RenameKey calls
//   - TransformKeys is more efficient than manual iteration
//   - For string keys, Prefix/Suffix are clearer than TransformKeys
//   - Cache sorted keys if used multiple times
//
// String Key Operations:
//   - PrefixKeys/SuffixKeys only work with string keys
//   - TransformKeys can work with any comparable type
//   - For case normalization, use TransformKeys with strings.ToUpper/ToLower
//   - Consider key collision when transforming (e.g., "A" and "a" → "a")
//
// keys.go는 Go를 위한 키 중심 맵 작업을 제공합니다.
//
// 이 파일은 맵 키를 추출, 변환 또는 조작하는 함수를 구현합니다.
// 작업에는 키 추출, 정렬, 이름 변경, 검색 및 변환이 포함됩니다.
// 모든 작업은 불변이며 새 데이터 구조를 반환합니다.
//
// 키 작업:
//
// 키 추출:
//   - KeysSlice(m): 모든 키를 슬라이스로 추출
//     시간: O(n), 공간: O(n)
//     convert.go의 Keys 별칭
//     순서는 비결정적 (맵 반복)
//     예: KeysSlice({"a": 1, "b": 2, "c": 3}) = ["a", "b", "c"] (순서 변함)
//     사용 사례: 키 열거, 반복, 집합 연산
//
//   - KeysSorted(m): 정렬된 순서로 키 추출
//     시간: O(n log n), 공간: O(n)
//     Ordered 제약 필요 (comparable + <)
//     결정적 출력 순서
//     예: KeysSorted({"c": 3, "a": 1, "b": 2}) = ["a", "b", "c"]
//     사용 사례: 결정적 출력, 순서 처리, 표시
//
//   - KeysBy(m, fn): 조건 일치 키 추출
//     시간: O(n), 공간: O(k) (k = 일치 키)
//     함수가 키와 값 수신
//     일치하는 키의 슬라이스 반환
//     예: KeysBy({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 1 }) = ["b", "c"]
//     사용 사례: 필터링된 키 추출, 조건부 선택, 검색
//
// 키 검색:
//   - FindKey(m, fn): 조건 일치 첫 번째 키 찾기
//     시간: 최악 O(n), 조기 종료
//     찾으면 (키, true), 없으면 (zero, false) 반환
//     맵 반복 순서로 인해 비결정적
//     예: FindKey({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 2 }) = ("c", true)
//     사용 사례: 조건별 키 조회, 존재 확인, 검색
//
//   - FindKeys(m, fn): 조건 일치 모든 키 찾기
//     시간: O(n), 공간: O(k) (k = 일치)
//     KeysBy 별칭 (의미적 명확성)
//     모든 일치 키 반환
//     예: FindKeys({"a": 1, "b": 2, "c": 3, "d": 4}, func(k, v) { return v > 2 }) = ["c", "d"]
//     사용 사례: 다중 일치 검색, 필터링, 배치 작업
//
// 키 이름 변경:
//   - RenameKey(m, oldKey, newKey): 단일 키 이름 변경
//     시간: O(n), 공간: O(n)
//     이름이 변경된 키가 있는 새 맵 반환
//     oldKey 없으면: 복제 반환
//     newKey 존재 시: 덮어쓰기
//     예: RenameKey({"a": 1, "b": 2}, "b", "B") = {"a": 1, "B": 2}
//     사용 사례: 키 정규화, 스키마 마이그레이션, 필드 이름 변경
//
//   - RenameKeys(m, mapping): 여러 키 이름 변경
//     시간: O(n), 공간: O(n)
//     mapping은 map[OldKey]NewKey
//     매핑에 없는 키는 변경 없음
//     예: RenameKeys({"a": 1, "b": 2}, {"a": "A"}) = {"A": 1, "b": 2}
//     사용 사례: 배치 이름 변경, 스키마 변환, 매핑 적용
//
//   - SwapKeys(m, key1, key2): 두 키 값 교환
//     시간: O(n), 공간: O(n)
//     두 키의 값 교환
//     어느 키든 없으면: 복제 반환
//     예: SwapKeys({"a": 1, "b": 2}, "a", "b") = {"a": 2, "b": 1}
//     사용 사례: 값 교환, 재정렬, 데이터 조작
//
// 키 변환:
//   - TransformKeys(m, fn): 모든 키 변환
//     시간: O(n), 공간: O(n)
//     함수가 키 수신, 새 키 반환
//     중복 키: 마지막 값 우선
//     예: TransformKeys({"a": 1, "b": 2}, func(k) { return strings.ToUpper(k) }) = {"A": 1, "B": 2}
//     사용 사례: 대소문자 정규화, 키 포맷팅, 대량 변환
//
//   - PrefixKeys(m, prefix): 모든 키에 접두사 추가
//     시간: O(n), 공간: O(n)
//     문자열 전용 작업
//     각 키에 접두사 연결
//     예: PrefixKeys({"a": 1, "b": 2}, "key_") = {"key_a": 1, "key_b": 2}
//     사용 사례: 네임스페이싱, 키 범위 지정, 접두사 추가
//
//   - SuffixKeys(m, suffix): 모든 키에 접미사 추가
//     시간: O(n), 공간: O(n)
//     문자열 전용 작업
//     각 키에 접미사 연결
//     예: SuffixKeys({"a": 1, "b": 2}, "_key") = {"a_key": 1, "b_key": 2}
//     사용 사례: 네임스페이싱, 키 범위 지정, 접미사 추가
//
// 엣지 케이스:
//   - 빈 맵에 KeysSlice: 빈 슬라이스 반환
//   - 빈 맵에 KeysSorted: 빈 슬라이스 반환
//   - 일치 없는 KeysBy: 빈 슬라이스 반환
//   - 일치 없는 FindKey: (zero, false) 반환
//   - 일치 없는 FindKeys: 빈 슬라이스 반환
//   - 없는 oldKey로 RenameKey: 복제 반환
//   - 존재하는 newKey로 RenameKey: 덮어쓰기
//   - 빈 매핑으로 RenameKeys: 복제 반환
//   - 없는 키로 SwapKeys: 복제 반환
//   - 중복 결과로 TransformKeys: 마지막 값 우선
//   - 빈 접두사로 PrefixKeys: 복제 반환
//   - 빈 접미사로 SuffixKeys: 복제 반환

// KeysSlice returns all keys from the map as a slice (alias for Keys).
// KeysSlice는 맵의 모든 키를 슬라이스로 반환합니다 (Keys의 별칭).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
func KeysSlice[K comparable, V any](m map[K]V) []K {
	return Keys(m)
}

// KeysSorted returns all keys from the map as a sorted slice.
// KeysSorted는 맵의 모든 키를 정렬된 슬라이스로 반환합니다.
//
// Time complexity: O(n log n)
// 시간 복잡도: O(n log n)
//
// Example
// 예제:
//
//	m := map[string]int{"c": 3, "a": 1, "b": 2}
//	keys := maputil.KeysSorted(m) // []string{"a", "b", "c"}
func KeysSorted[K Ordered, V any](m map[K]V) []K {
	keys := Keys(m)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

// KeysBy returns keys from the map that satisfy the predicate.
// KeysBy는 조건을 만족하는 맵의 키를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	keys := maputil.KeysBy(m, func(k string, v int) bool {
//	    return v > 2
//	}) // []string{"c", "d"} (order may vary)
func KeysBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []K {
	keys := make([]K, 0)
	for k, v := range m {
		if fn(k, v) {
			keys = append(keys, k)
		}
	}
	return keys
}

// RenameKey creates a new map with a key renamed.
// RenameKey는 키의 이름이 변경된 새 맵을 생성합니다.
//
// If the old key doesn't exist, returns a clone of the original map.
// If the new key already exists, it will be overwritten.
//
// 이전 키가 존재하지 않으면 원본 맵의 복사본을 반환합니다.
// 새 키가 이미 존재하면 덮어씁니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.RenameKey(m, "b", "B") // map[string]int{"a": 1, "B": 2, "c": 3}
func RenameKey[K comparable, V any](m map[K]V, oldKey, newKey K) map[K]V {
	result := Clone(m)

	if value, exists := result[oldKey]; exists {
		delete(result, oldKey)
		result[newKey] = value
	}

	return result
}

// RenameKeys creates a new map with multiple keys renamed according to the mapping.
// RenameKeys는 매핑에 따라 여러 키의 이름이 변경된 새 맵을 생성합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	mapping := map[string]string{"a": "A", "b": "B"}
//	result := maputil.RenameKeys(m, mapping) // map[string]int{"A": 1, "B": 2, "c": 3}
func RenameKeys[K comparable, V any](m map[K]V, mapping map[K]K) map[K]V {
	result := make(map[K]V, len(m))

	for k, v := range m {
		if newKey, shouldRename := mapping[k]; shouldRename {
			result[newKey] = v
		} else {
			result[k] = v
		}
	}

	return result
}

// SwapKeys creates a new map with two keys swapped.
// SwapKeys는 두 키가 교환된 새 맵을 생성합니다.
//
// If either key doesn't exist, returns a clone of the original map.
// 어느 한 키라도 존재하지 않으면 원본 맵의 복사본을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.SwapKeys(m, "a", "b") // map[string]int{"a": 2, "b": 1, "c": 3}
func SwapKeys[K comparable, V any](m map[K]V, key1, key2 K) map[K]V {
	result := Clone(m)

	val1, exists1 := result[key1]
	val2, exists2 := result[key2]

	if exists1 && exists2 {
		result[key1] = val2
		result[key2] = val1
	}

	return result
}

// FindKey finds the first key that satisfies the predicate.
// FindKey는 조건을 만족하는 첫 번째 키를 찾습니다.
//
// Returns the key and true if found, or zero value and false otherwise.
// 찾으면 키와 true를 반환하고, 그렇지 않으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n) worst case, early termination when found
// 시간 복잡도: 최악의 경우 O(n), 찾으면 조기 종료
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	key, found := maputil.FindKey(m, func(k string, v int) bool {
//	    return v > 2
//	}) // key = "c", found = true (may vary due to map iteration order)
func FindKey[K comparable, V any](m map[K]V, fn func(K, V) bool) (K, bool) {
	for k, v := range m {
		if fn(k, v) {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// FindKeys finds all keys that satisfy the predicate.
// FindKeys는 조건을 만족하는 모든 키를 찾습니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	keys := maputil.FindKeys(m, func(k string, v int) bool {
//	    return v > 2
//	}) // []string{"c", "d"} (order may vary)
func FindKeys[K comparable, V any](m map[K]V, fn func(K, V) bool) []K {
	return KeysBy(m, fn)
}

// PrefixKeys adds a prefix to all keys in a map.
// PrefixKeys는 맵의 모든 키에 접두사를 추가합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	result := maputil.PrefixKeys(m, "key_")
//	// map[string]int{"key_a": 1, "key_b": 2}
func PrefixKeys[V any](m map[string]V, prefix string) map[string]V {
	result := make(map[string]V, len(m))
	for k, v := range m {
		result[prefix+k] = v
	}
	return result
}

// SuffixKeys adds a suffix to all keys in a map.
// SuffixKeys는 맵의 모든 키에 접미사를 추가합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	result := maputil.SuffixKeys(m, "_key")
//	// map[string]int{"a_key": 1, "b_key": 2}
func SuffixKeys[V any](m map[string]V, suffix string) map[string]V {
	result := make(map[string]V, len(m))
	for k, v := range m {
		result[k+suffix] = v
	}
	return result
}

// TransformKeys transforms all keys in a map using a function.
// TransformKeys는 함수를 사용하여 맵의 모든 키를 변환합니다.
//
// If multiple keys map to the same transformed key, the last value wins.
// 여러 키가 같은 변환된 키로 매핑되면 마지막 값이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.TransformKeys(m, func(k string) string {
//	    return strings.ToUpper(k)
//	}) // map[string]int{"A": 1, "B": 2, "C": 3}
func TransformKeys[K comparable, V any](m map[K]V, fn func(K) K) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		newKey := fn(k)
		result[newKey] = v
	}
	return result
}
