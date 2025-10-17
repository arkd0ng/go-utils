package maputil

import (
	"fmt"
	"strings"
)

// transform.go provides map transformation operations for Go.
//
// This file implements functions that transform maps structurally or by applying
// functions to their keys and values. All operations follow functional programming
// principles and return new maps (immutable).
//
// Transformation Categories:
//
// Value Transformation:
//   - Map(m, fn): Transform values with key-value function
//     Time: O(n), Space: O(n)
//     Function receives both key and value
//     Returns new map with transformed values
//     Example: Map({"a": 1, "b": 2}, func(k, v) { return v*2 }) = {"a": 2, "b": 4}
//     Use cases: Data conversion, calculations, formatting
//
//   - MapValues(m, fn): Transform values with value-only function
//     Time: O(n), Space: O(n)
//     Function receives only value (simpler)
//     Useful when key is not needed
//     Example: MapValues({"a": 1, "b": 2}, func(v) { return v*2 }) = {"a": 2, "b": 4}
//     Use cases: Simple transformations, math operations, type conversions
//
// Key Transformation:
//   - MapKeys(m, fn): Transform keys using function
//     Time: O(n), Space: O(n)
//     Function receives key and value
//     Duplicate keys: Last value wins
//     Example: MapKeys({"a": 1, "b": 2}, func(k, v) { return strings.ToUpper(k) }) = {"A": 1, "B": 2}
//     Use cases: Case normalization, key formatting, namespace prefixing
//
// Entry Transformation:
//   - MapEntries(m, fn): Transform both keys and values
//     Time: O(n), Space: O(n)
//     Function receives key and value, returns new key and value
//     Allows type changes for both
//     Duplicate keys: Last value wins
//     Example: MapEntries({"a": 1}, func(k, v) { return v, k }) = {1: "a"} (swap)
//     Use cases: Complex transformations, type swapping, restructuring
//
// Structure Transformation:
//   - Invert(m): Swap keys and values
//     Time: O(n), Space: O(n)
//     Requires both K and V to be comparable
//     Duplicate values: Last key wins
//     Example: Invert({"a": 1, "b": 2}) = {1: "a", 2: "b"}
//     Use cases: Reverse lookups, ID to name mapping, index creation
//
//   - Flatten(m, delimiter): Nested map to flat map
//     Time: O(n*m) where m = average nested size, Space: O(n*m)
//     Joins nested keys with delimiter
//     Example: Flatten({"user": {"name": "Alice", "age": 30}}, ".") = {"user.name": "Alice", "user.age": 30}
//     Use cases: Configuration flattening, JSON to flat structure, path-based access
//
//   - Unflatten(m, delimiter): Flat map to nested map
//     Time: O(n*k) where k = average key depth, Space: O(n*k)
//     Splits keys by delimiter and builds nested structure
//     Example: Unflatten({"user.name": "Alice", "user.age": 30}, ".") = {"user": {"name": "Alice", "age": 30}}
//     Use cases: Config parsing, path-based data, hierarchical structures
//
// Partitioning:
//   - Chunk(m, size): Split into smaller maps
//     Time: O(n), Space: O(n)
//     Divides map into chunks of specified size
//     Last chunk may be smaller
//     Order not guaranteed (map iteration)
//     Example: Chunk({"a": 1, "b": 2, "c": 3}, 2) = [{"a": 1, "b": 2}, {"c": 3}]
//     Use cases: Batch processing, pagination, parallel processing
//
//   - Partition(m, fn): Split into two maps by predicate
//     Time: O(n), Space: O(n)
//     First map: entries satisfying predicate
//     Second map: entries not satisfying predicate
//     Example: Partition({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 1 })
//              = ({"b": 2, "c": 3}, {"a": 1})
//     Use cases: Filtering with remainder, validation, classification
//
// Value Filtering:
//   - Compact(m): Remove zero values
//     Time: O(n), Space: O(n)
//     Removes entries where value equals zero value
//     Example: Compact({"a": 1, "b": 0, "c": 3}) = {"a": 1, "c": 3}
//     Use cases: Cleanup, sparse data, removing defaults
//
// Design Principles:
//   - Immutability: All transformations return new maps
//   - Composability: Functions can be chained
//   - Type Safety: Leverage generics for compile-time safety
//   - Predictability: Consistent behavior with edge cases
//   - Flexibility: Multiple transformation options for different needs
//
// Comparison: Map vs MapValues:
//   - Map: Access to both key and value in function
//   - MapValues: Access to value only (simpler signature)
//   - Map: More flexible, can use key in transformation
//   - MapValues: Cleaner when key is not needed
//
// Comparison: MapKeys vs MapEntries:
//   - MapKeys: Transform keys only, values unchanged
//   - MapEntries: Transform both keys and values
//   - MapKeys: Simpler when only keys need change
//   - MapEntries: More flexible, allows type changes
//
// Comparison: Partition vs Filter:
//   - Partition: Returns both matching and non-matching
//   - Filter: Returns only matching entries
//   - Partition: No data loss, both groups preserved
//   - Filter: Discards non-matching entries
//
// Comparison: Flatten vs Manual:
//   - Flatten: One-line, consistent delimiter handling
//   - Manual: Requires nested loops, error-prone
//   - Flatten: Limited to two levels of nesting
//   - Manual: Can handle arbitrary depth
//
// Performance Characteristics:
//
// Time Complexity:
//   - Map/MapValues/MapKeys/MapEntries: O(n) - Single iteration
//   - Invert/Compact: O(n) - Single iteration
//   - Chunk/Partition: O(n) - Single iteration
//   - Flatten: O(n*m) - Nested iteration
//   - Unflatten: O(n*k) - String splitting and nesting
//
// Space Complexity:
//   - All operations: O(n) to O(n*m) - New maps allocated
//   - Chunk: O(n) + slice overhead
//   - Partition: 2*O(n) - Two maps
//   - Flatten/Unflatten: O(n*m) - Keys may be longer
//
// Memory Allocation:
//   - All create new maps (immutable operations)
//   - MapEntries may change type sizes
//   - Chunk creates slice of maps
//   - Unflatten creates nested map structures
//
// Common Usage Patterns:
//
//	// Chain transformations
//	m := map[string]int{"a": 1, "b": 0, "c": 3}
//	result := maputil.MapValues(
//	    maputil.Compact(m),
//	    func(v int) int { return v * 2 },
//	) // {"a": 2, "c": 6}
//
//	// Normalize keys
//	config := map[string]string{"Host": "localhost", "Port": "8080"}
//	normalized := maputil.MapKeys(config, func(k string, v string) string {
//	    return strings.ToLower(k)
//	}) // {"host": "localhost", "port": "8080"}
//
//	// Create reverse index
//	idToName := map[int]string{1: "Alice", 2: "Bob"}
//	nameToId := maputil.Invert(idToName) // {"Alice": 1, "Bob": 2}
//
//	// Flatten nested config
//	config := map[string]map[string]int{
//	    "database": {"timeout": 30, "pool": 10},
//	    "cache": {"ttl": 300},
//	}
//	flat := maputil.Flatten(config, ".")
//	// {"database.timeout": 30, "database.pool": 10, "cache.ttl": 300}
//
//	// Batch processing
//	data := map[string]int{...} // 1000 entries
//	batches := maputil.Chunk(data, 100) // 10 batches of 100
//	for _, batch := range batches {
//	    processBatch(batch)
//	}
//
//	// Separate valid and invalid
//	entries := map[string]int{"a": 1, "b": -1, "c": 5, "d": -2}
//	valid, invalid := maputil.Partition(entries, func(k string, v int) bool {
//	    return v > 0
//	})
//	// valid = {"a": 1, "c": 5}, invalid = {"b": -1, "d": -2}
//
// Type Transformation:
//
//	// Convert types
//	intMap := map[string]int{"a": 1, "b": 2}
//	stringMap := maputil.MapValues(intMap, func(v int) string {
//	    return fmt.Sprintf("%d", v)
//	}) // map[string]string{"a": "1", "b": "2"}
//
//	// Swap key-value types
//	original := map[string]int{"a": 1, "b": 2}
//	swapped := maputil.MapEntries(original, func(k string, v int) (int, string) {
//	    return v, k
//	}) // map[int]string{1: "a", 2: "b"}
//
// Edge Cases:
//   - Map on empty map: Returns empty map
//   - MapKeys with duplicate results: Last key-value wins
//   - MapEntries with duplicate keys: Last transformation wins
//   - Invert with duplicate values: Last key wins
//   - Flatten on empty nested: Returns empty map
//   - Unflatten with malformed keys: Creates partial structure
//   - Chunk with size <= 0: Returns nil
//   - Chunk with size > len(m): Returns single chunk
//   - Partition on empty: Returns two empty maps
//   - Compact with all zeros: Returns empty map
//
// Duplicate Key Handling:
//   - MapKeys: If fn produces same key for multiple entries, last value wins
//   - MapEntries: If fn produces same key for multiple entries, last wins
//   - Invert: If multiple keys have same value, last key wins
//   - Map iteration order is random, so "last" is non-deterministic
//   - For predictable results, ensure transformations produce unique keys
//
// Thread Safety:
//   - All operations create new maps (safe for concurrent reads of input)
//   - Not safe if input map is modified concurrently
//   - Output maps are independent (safe to use in different goroutines)
//   - Transformation functions must be goroutine-safe if accessing shared state
//
// Performance Tips:
//   - MapValues is slightly faster than Map when key not needed
//   - Preallocate result maps when size is known
//   - Chunk size affects memory locality and GC pressure
//   - Flatten/Unflatten have string concatenation overhead
//   - Consider caching Invert results for frequently accessed reverse maps
//   - Partition is more efficient than Filter + complementary Filter
//
// Nil Map Behavior:
//   - Map(nil, fn): Returns empty map
//   - MapValues(nil, fn): Returns empty map
//   - MapKeys(nil, fn): Returns empty map
//   - MapEntries(nil, fn): Returns empty map
//   - Invert(nil): Returns empty map
//   - Flatten(nil, delim): Returns empty map
//   - Unflatten(nil, delim): Returns empty map
//   - Chunk(nil, size): Returns nil or empty slice
//   - Partition(nil, fn): Returns two empty maps
//   - Compact(nil): Returns empty map
//
// transform.go는 Go를 위한 맵 변환 작업을 제공합니다.
//
// 이 파일은 맵을 구조적으로 또는 키와 값에 함수를 적용하여 변환하는
// 함수를 구현합니다. 모든 작업은 함수형 프로그래밍 원칙을 따르며
// 새 맵을 반환합니다 (불변).
//
// 변환 카테고리:
//
// 값 변환:
//   - Map(m, fn): 키-값 함수로 값 변환
//     시간: O(n), 공간: O(n)
//     함수가 키와 값 모두 수신
//     변환된 값이 있는 새 맵 반환
//     예: Map({"a": 1, "b": 2}, func(k, v) { return v*2 }) = {"a": 2, "b": 4}
//     사용 사례: 데이터 변환, 계산, 포맷팅
//
//   - MapValues(m, fn): 값만 사용하는 함수로 값 변환
//     시간: O(n), 공간: O(n)
//     함수가 값만 수신 (더 단순)
//     키가 필요 없을 때 유용
//     예: MapValues({"a": 1, "b": 2}, func(v) { return v*2 }) = {"a": 2, "b": 4}
//     사용 사례: 간단한 변환, 수학 연산, 타입 변환
//
// 키 변환:
//   - MapKeys(m, fn): 함수를 사용하여 키 변환
//     시간: O(n), 공간: O(n)
//     함수가 키와 값 수신
//     중복 키: 마지막 값이 우선
//     예: MapKeys({"a": 1, "b": 2}, func(k, v) { return strings.ToUpper(k) }) = {"A": 1, "B": 2}
//     사용 사례: 대소문자 정규화, 키 포맷팅, 네임스페이스 접두사
//
// 항목 변환:
//   - MapEntries(m, fn): 키와 값 모두 변환
//     시간: O(n), 공간: O(n)
//     함수가 키와 값 수신, 새 키와 값 반환
//     둘 다 타입 변경 가능
//     중복 키: 마지막 값이 우선
//     예: MapEntries({"a": 1}, func(k, v) { return v, k }) = {1: "a"} (교환)
//     사용 사례: 복잡한 변환, 타입 교환, 구조 재편성
//
// 구조 변환:
//   - Invert(m): 키와 값 교환
//     시간: O(n), 공간: O(n)
//     K와 V 모두 comparable 필요
//     중복 값: 마지막 키가 우선
//     예: Invert({"a": 1, "b": 2}) = {1: "a", 2: "b"}
//     사용 사례: 역 조회, ID에서 이름 매핑, 인덱스 생성
//
//   - Flatten(m, delimiter): 중첩 맵을 평면 맵으로
//     시간: O(n*m) (m = 평균 중첩 크기), 공간: O(n*m)
//     중첩 키를 구분자로 결합
//     예: Flatten({"user": {"name": "Alice", "age": 30}}, ".") = {"user.name": "Alice", "user.age": 30}
//     사용 사례: 설정 평면화, JSON을 평면 구조로, 경로 기반 접근
//
//   - Unflatten(m, delimiter): 평면 맵을 중첩 맵으로
//     시간: O(n*k) (k = 평균 키 깊이), 공간: O(n*k)
//     구분자로 키를 분할하고 중첩 구조 구축
//     예: Unflatten({"user.name": "Alice", "user.age": 30}, ".") = {"user": {"name": "Alice", "age": 30}}
//     사용 사례: 설정 파싱, 경로 기반 데이터, 계층 구조
//
// 분할:
//   - Chunk(m, size): 더 작은 맵으로 분할
//     시간: O(n), 공간: O(n)
//     지정된 크기의 청크로 맵 분할
//     마지막 청크는 더 작을 수 있음
//     순서 보장 없음 (맵 반복)
//     예: Chunk({"a": 1, "b": 2, "c": 3}, 2) = [{"a": 1, "b": 2}, {"c": 3}]
//     사용 사례: 배치 처리, 페이지네이션, 병렬 처리
//
//   - Partition(m, fn): 조건으로 두 맵으로 분할
//     시간: O(n), 공간: O(n)
//     첫 번째 맵: 조건 만족 항목
//     두 번째 맵: 조건 불만족 항목
//     예: Partition({"a": 1, "b": 2, "c": 3}, func(k, v) { return v > 1 })
//              = ({"b": 2, "c": 3}, {"a": 1})
//     사용 사례: 나머지가 있는 필터링, 검증, 분류
//
// 값 필터링:
//   - Compact(m): zero 값 제거
//     시간: O(n), 공간: O(n)
//     값이 zero 값과 같은 항목 제거
//     예: Compact({"a": 1, "b": 0, "c": 3}) = {"a": 1, "c": 3}
//     사용 사례: 정리, 희소 데이터, 기본값 제거
//
// 설계 원칙:
//   - 불변성: 모든 변환이 새 맵 반환
//   - 구성 가능성: 함수 체인 가능
//   - 타입 안전성: 컴파일 타임 안전성을 위해 제네릭 활용
//   - 예측 가능성: 엣지 케이스에서 일관된 동작
//   - 유연성: 다른 필요에 대한 여러 변환 옵션
//
// 비교: Map vs MapValues:
//   - Map: 함수에서 키와 값 모두 접근
//   - MapValues: 값만 접근 (더 단순한 시그니처)
//   - Map: 더 유연, 변환에 키 사용 가능
//   - MapValues: 키가 필요 없을 때 더 깔끔
//
// 비교: MapKeys vs MapEntries:
//   - MapKeys: 키만 변환, 값 변경 없음
//   - MapEntries: 키와 값 모두 변환
//   - MapKeys: 키만 변경 필요할 때 더 단순
//   - MapEntries: 더 유연, 타입 변경 가능
//
// 비교: Partition vs Filter:
//   - Partition: 일치 및 불일치 모두 반환
//   - Filter: 일치 항목만 반환
//   - Partition: 데이터 손실 없음, 두 그룹 모두 보존
//   - Filter: 불일치 항목 폐기
//
// 비교: Flatten vs 수동:
//   - Flatten: 한 줄, 일관된 구분자 처리
//   - 수동: 중첩 루프 필요, 오류 가능성
//   - Flatten: 두 단계 중첩으로 제한
//   - 수동: 임의 깊이 처리 가능
//
// 성능 특성:
//
// 시간 복잡도:
//   - Map/MapValues/MapKeys/MapEntries: O(n) - 단일 반복
//   - Invert/Compact: O(n) - 단일 반복
//   - Chunk/Partition: O(n) - 단일 반복
//   - Flatten: O(n*m) - 중첩 반복
//   - Unflatten: O(n*k) - 문자열 분할 및 중첩
//
// 공간 복잡도:
//   - 모든 작업: O(n)에서 O(n*m) - 새 맵 할당
//   - Chunk: O(n) + 슬라이스 오버헤드
//   - Partition: 2*O(n) - 두 맵
//   - Flatten/Unflatten: O(n*m) - 키가 더 길 수 있음
//
// 메모리 할당:
//   - 모두 새 맵 생성 (불변 작업)
//   - MapEntries는 타입 크기 변경 가능
//   - Chunk는 맵의 슬라이스 생성
//   - Unflatten은 중첩 맵 구조 생성
//
// 일반적인 사용 패턴:
//
//	// 변환 체인
//	m := map[string]int{"a": 1, "b": 0, "c": 3}
//	result := maputil.MapValues(
//	    maputil.Compact(m),
//	    func(v int) int { return v * 2 },
//	) // {"a": 2, "c": 6}
//
//	// 키 정규화
//	config := map[string]string{"Host": "localhost", "Port": "8080"}
//	normalized := maputil.MapKeys(config, func(k string, v string) string {
//	    return strings.ToLower(k)
//	}) // {"host": "localhost", "port": "8080"}
//
//	// 역 인덱스 생성
//	idToName := map[int]string{1: "Alice", 2: "Bob"}
//	nameToId := maputil.Invert(idToName) // {"Alice": 1, "Bob": 2}
//
//	// 중첩 설정 평면화
//	config := map[string]map[string]int{
//	    "database": {"timeout": 30, "pool": 10},
//	    "cache": {"ttl": 300},
//	}
//	flat := maputil.Flatten(config, ".")
//	// {"database.timeout": 30, "database.pool": 10, "cache.ttl": 300}
//
//	// 배치 처리
//	data := map[string]int{...} // 1000 항목
//	batches := maputil.Chunk(data, 100) // 100개의 배치 10개
//	for _, batch := range batches {
//	    processBatch(batch)
//	}
//
//	// 유효 및 무효 분리
//	entries := map[string]int{"a": 1, "b": -1, "c": 5, "d": -2}
//	valid, invalid := maputil.Partition(entries, func(k string, v int) bool {
//	    return v > 0
//	})
//	// valid = {"a": 1, "c": 5}, invalid = {"b": -1, "d": -2}
//
// 타입 변환:
//
//	// 타입 변환
//	intMap := map[string]int{"a": 1, "b": 2}
//	stringMap := maputil.MapValues(intMap, func(v int) string {
//	    return fmt.Sprintf("%d", v)
//	}) // map[string]string{"a": "1", "b": "2"}
//
//	// 키-값 타입 교환
//	original := map[string]int{"a": 1, "b": 2}
//	swapped := maputil.MapEntries(original, func(k string, v int) (int, string) {
//	    return v, k
//	}) // map[int]string{1: "a", 2: "b"}
//
// 엣지 케이스:
//   - 빈 맵에 Map: 빈 맵 반환
//   - 중복 결과로 MapKeys: 마지막 키-값이 우선
//   - 중복 키로 MapEntries: 마지막 변환이 우선
//   - 중복 값으로 Invert: 마지막 키가 우선
//   - 빈 중첩에 Flatten: 빈 맵 반환
//   - 잘못된 키로 Unflatten: 부분 구조 생성
//   - size <= 0으로 Chunk: nil 반환
//   - size > len(m)로 Chunk: 단일 청크 반환
//   - 빈 것에 Partition: 두 빈 맵 반환
//   - 모두 zero로 Compact: 빈 맵 반환
//
// 중복 키 처리:
//   - MapKeys: fn이 여러 항목에 대해 같은 키 생성 시 마지막 값이 우선
//   - MapEntries: fn이 여러 항목에 대해 같은 키 생성 시 마지막이 우선
//   - Invert: 여러 키가 같은 값을 가지면 마지막 키가 우선
//   - 맵 반복 순서는 무작위이므로 "마지막"은 비결정적
//   - 예측 가능한 결과를 위해 변환이 고유 키 생성하도록 보장
//
// 스레드 안전성:
//   - 모든 작업이 새 맵 생성 (입력의 동시 읽기 안전)
//   - 입력 맵이 동시에 수정되면 안전하지 않음
//   - 출력 맵은 독립적 (다른 고루틴에서 사용 안전)
//   - 변환 함수가 공유 상태 접근 시 고루틴 안전해야 함
//
// 성능 팁:
//   - 키가 필요 없을 때 Map보다 MapValues가 약간 빠름
//   - 크기를 알 때 결과 맵 사전 할당
//   - 청크 크기가 메모리 지역성 및 GC 압력에 영향
//   - Flatten/Unflatten은 문자열 연결 오버헤드 있음
//   - 자주 접근하는 역 맵은 Invert 결과 캐싱 고려
//   - Partition이 Filter + 보완 Filter보다 효율적
//
// Nil 맵 동작:
//   - Map(nil, fn): 빈 맵 반환
//   - MapValues(nil, fn): 빈 맵 반환
//   - MapKeys(nil, fn): 빈 맵 반환
//   - MapEntries(nil, fn): 빈 맵 반환
//   - Invert(nil): 빈 맵 반환
//   - Flatten(nil, delim): 빈 맵 반환
//   - Unflatten(nil, delim): 빈 맵 반환
//   - Chunk(nil, size): nil 또는 빈 슬라이스 반환
//   - Partition(nil, fn): 두 빈 맵 반환
//   - Compact(nil): 빈 맵 반환

// Map transforms map values using the provided function.
// Map은 제공된 함수를 사용하여 맵 값을 변환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.Map(m, func(k string, v int) string {
//	    return fmt.Sprintf("%s=%d", k, v)
//	}) // map[string]string{"a": "a=1", "b": "b=2", "c": "c=3"}
func Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = fn(k, v)
	}
	return result
}

// MapKeys transforms map keys using the provided function.
// MapKeys는 제공된 함수를 사용하여 맵 키를 변환합니다.
//
// If the function produces duplicate keys, the last value wins.
// 함수가 중복 키를 생성하면 마지막 값이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.MapKeys(m, func(k string, v int) string {
//	    return strings.ToUpper(k)
//	}) // map[string]int{"A": 1, "B": 2, "C": 3}
func MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V {
	result := make(map[R]V, len(m))
	for k, v := range m {
		newKey := fn(k, v)
		result[newKey] = v
	}
	return result
}

// MapValues transforms map values using the provided function (value-only version).
// MapValues는 제공된 함수를 사용하여 맵 값을 변환합니다 (값만 사용하는 버전).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.MapValues(m, func(v int) int {
//	    return v * 2
//	}) // map[string]int{"a": 2, "b": 4, "c": 6}
func MapValues[K comparable, V any, R any](m map[K]V, fn func(V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = fn(v)
	}
	return result
}

// MapEntries transforms both keys and values using the provided function.
// MapEntries는 제공된 함수를 사용하여 키와 값을 모두 변환합니다.
//
// If the function produces duplicate keys, the last value wins.
// 함수가 중복 키를 생성하면 마지막 값이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//
// result := maputil.MapEntries(m, func(k string, v int) (int, string) {
// return v, k // Swap key and value types
// 키와 값 타입 교환
//
//	}) // map[int]string{1: "a", 2: "b"}
func MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, fn func(K1, V1) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(m))
	for k, v := range m {
		newKey, newValue := fn(k, v)
		result[newKey] = newValue
	}
	return result
}

// Invert swaps keys and values in the map.
// Invert는 맵의 키와 값을 교환합니다.
//
// If multiple keys have the same value, the last key-value pair wins.
// 여러 키가 같은 값을 가지면 마지막 키-값 쌍이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.Invert(m) // map[int]string{1: "a", 2: "b", 3: "c"}
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// Flatten converts a nested map to a flat map using a delimiter.
// Flatten은 구분자를 사용하여 중첩된 맵을 평면 맵으로 변환합니다.
//
// Time complexity: O(n*m) where n is the number of keys and m is the average nested map size
// 시간 복잡도: O(n*m) 여기서 n은 키 개수, m은 평균 중첩 맵 크기
//
// Example
// 예제:
//
//	m := map[string]map[string]int{
//	    "user1": {"age": 25, "score": 100},
//	    "user2": {"age": 30, "score": 95},
//	}
//	result := maputil.Flatten(m, ".") // map[string]int{"user1.age": 25, "user1.score": 100, ...}
func Flatten[K comparable, V any](m map[K]map[string]V, delimiter string) map[string]V {
	result := make(map[string]V)
	for k, nested := range m {
		prefix := fmt.Sprintf("%v", k)
		for nestedKey, nestedValue := range nested {
			flatKey := prefix + delimiter + nestedKey
			result[flatKey] = nestedValue
		}
	}
	return result
}

// Unflatten converts a flat map to a nested map using a delimiter.
// Unflatten은 구분자를 사용하여 평면 맵을 중첩된 맵으로 변환합니다.
//
// Time complexity: O(n*k) where n is the number of keys and k is the average key depth
// 시간 복잡도: O(n*k) 여기서 n은 키 개수, k는 평균 키 깊이
//
// Example
// 예제:
//
//	m := map[string]int{"user.name": 1, "user.age": 25, "admin.name": 2}
//	result := maputil.Unflatten(m, ".") // map[string]interface{}{"user": {...}, "admin": {...}}
func Unflatten[V any](m map[string]V, delimiter string) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range m {
		parts := strings.Split(key, delimiter)
		current := result

		for i := 0; i < len(parts)-1; i++ {
			part := parts[i]
			if _, exists := current[part]; !exists {
				current[part] = make(map[string]interface{})
			}
			current = current[part].(map[string]interface{})
		}

		current[parts[len(parts)-1]] = value
	}

	return result
}

// Chunk splits a map into multiple smaller maps of the specified size.
// Chunk는 맵을 지정된 크기의 여러 작은 맵으로 분할합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
//	result := maputil.Chunk(m, 2) // []map[string]int{{"a": 1, "b": 2}, {"c": 3, "d": 4}, {"e": 5}}
func Chunk[K comparable, V any](m map[K]V, size int) []map[K]V {
	if size <= 0 {
		return nil
	}

	numChunks := (len(m) + size - 1) / size
	result := make([]map[K]V, 0, numChunks)
	current := make(map[K]V, size)

	i := 0
	for k, v := range m {
		current[k] = v
		i++

		if i%size == 0 {
			result = append(result, current)
			current = make(map[K]V, size)
		}
	}

	if len(current) > 0 {
		result = append(result, current)
	}

	return result
}

// Partition splits a map into two maps based on a predicate function.
// Partition은 조건 함수를 기반으로 맵을 두 개의 맵으로 분할합니다.
//
// Returns two maps: the first contains entries that satisfy the predicate,
// the second contains entries that don't.
//
// 두 개의 맵을 반환합니다: 첫 번째는 조건을 만족하는 항목,
// 두 번째는 그렇지 않은 항목을 포함합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	even, odd := maputil.Partition(m, func(k string, v int) bool {
//	    return v%2 == 0
//	}) // even = {"b": 2, "d": 4}, odd = {"a": 1, "c": 3}
func Partition[K comparable, V any](m map[K]V, fn func(K, V) bool) (map[K]V, map[K]V) {
	trueMap := make(map[K]V)
	falseMap := make(map[K]V)

	for k, v := range m {
		if fn(k, v) {
			trueMap[k] = v
		} else {
			falseMap[k] = v
		}
	}

	return trueMap, falseMap
}

// Compact removes entries with zero values from the map.
// Compact는 맵에서 zero 값을 가진 항목을 제거합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 0, "c": 3, "d": 0}
//	result := maputil.Compact(m) // map[string]int{"a": 1, "c": 3}
func Compact[K comparable, V comparable](m map[K]V) map[K]V {
	var zero V
	result := make(map[K]V)

	for k, v := range m {
		if v != zero {
			result[k] = v
		}
	}

	return result
}
