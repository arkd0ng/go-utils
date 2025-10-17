package maputil

// Package maputil/comparison.go provides comprehensive map comparison operations.
// This file contains functions for comparing maps, detecting differences,
// finding common keys, and performing detailed change analysis.
//
// maputil/comparison.go 패키지는 포괄적인 맵 비교 작업을 제공합니다.
// 이 파일은 맵 비교, 차이 감지, 공통 키 찾기, 상세한 변경 분석을 수행하는 함수들을 포함합니다.
//
// # Overview | 개요
//
// The comparison.go file provides 7 map comparison operations organized into 3 categories:
//
// comparison.go 파일은 3개 카테고리로 구성된 7개 맵 비교 작업을 제공합니다:
//
// 1. EQUALITY TESTING | 동등성 테스트
//   - EqualMaps: Check map equality with comparable values (O(n))
//   - EqualFunc: Check map equality with custom comparator (O(n))
//
// 2. DIFFERENCE DETECTION | 차이 감지
//   - Diff: Get map of differing entries (O(n+m))
//   - DiffKeys: Get slice of differing keys (O(n+m))
//   - Compare: Detailed diff with added/removed/modified (O(n+m))
//
// 3. KEY ANALYSIS | 키 분석
//   - CommonKeys: Keys present in all maps (O(n*m))
//   - AllKeys: All unique keys from all maps (O(n*m))
//
// # Design Principles | 설계 원칙
//
// 1. COMPREHENSIVE COMPARISON | 포괄적 비교
//   - Size comparison first (quick inequality check)
//   - Key existence and value equality checking
//   - Support for both comparable and non-comparable values
//   - Custom comparators for complex types
//
//   크기 먼저 비교 (빠른 불일치 검사)
//   키 존재와 값 동등성 확인
//   비교 가능 및 비교 불가능 값 모두 지원
//   복잡한 타입을 위한 커스텀 비교자
//
// 2. DETAILED CHANGE TRACKING | 상세한 변경 추적
//   - Compare function returns added, removed, modified separately
//   - Diff functions provide different views of changes
//   - DiffKeys for key-only change detection
//   - Support version control-like diff semantics
//
//   Compare 함수는 추가, 제거, 수정을 별도로 반환
//   Diff 함수는 변경사항의 다른 뷰 제공
//   키만의 변경 감지를 위한 DiffKeys
//   버전 관리 같은 diff 의미론 지원
//
// 3. FLEXIBLE VALUE COMPARISON | 유연한 값 비교
//   - EqualMaps: Built-in == operator for comparable types
//   - EqualFunc: Custom comparison logic for any types
//   - Diff/DiffKeys: Value-aware difference detection
//   - Support for slices, structs, complex types via EqualFunc
//
//   EqualMaps: 비교 가능 타입을 위한 내장 == 연산자
//   EqualFunc: 모든 타입을 위한 커스텀 비교 로직
//   Diff/DiffKeys: 값 인식 차이 감지
//   EqualFunc를 통한 슬라이스, 구조체, 복잡한 타입 지원
//
// 4. KEY AGGREGATION | 키 집계
//   - CommonKeys: Intersection-like operation for keys
//   - AllKeys: Union-like operation for keys
//   - Deduplication in AllKeys ensures unique results
//   - Both support variadic input for multiple maps
//
//   CommonKeys: 키를 위한 교집합 같은 연산
//   AllKeys: 키를 위한 합집합 같은 연산
//   AllKeys의 중복 제거로 고유 결과 보장
//   둘 다 여러 맵을 위한 가변 인자 지원
//
// # Function Categories | 함수 카테고리
//
// EQUALITY TESTING OPERATIONS | 동등성 테스트 작업
//
// EqualMaps(m1, m2) checks whether two maps are equal by comparing both keys and values.
// Returns true if maps have identical keys with identical values, false otherwise.
// Requires values to be comparable (support == operator).
//
// EqualMaps(m1, m2)는 키와 값을 모두 비교하여 두 맵이 동일한지 확인합니다.
// 맵이 동일한 값을 가진 동일한 키를 가지면 true, 그렇지 않으면 false를 반환합니다.
// 값이 비교 가능해야 합니다 (== 연산자 지원).
//
// Time Complexity: O(n) where n = size of smaller map
// Space Complexity: O(1) - no allocations
// Comparison: Uses built-in == operator for values
// Early Termination: Returns false on first mismatch or size difference
// Constraint: Requires comparable constraint for values
// Nil Handling: nil maps treated as empty, nil == empty returns true
//
// 시간 복잡도: O(n) 여기서 n = 더 작은 맵 크기
// 공간 복잡도: O(1) - 할당 없음
// 비교: 값에 대해 내장 == 연산자 사용
// 조기 종료: 첫 불일치나 크기 차이에서 false 반환
// 제약: 값에 comparable 제약 필요
// Nil 처리: nil 맵을 빈 것으로 처리, nil == empty는 true 반환
//
// Use Case: Quick equality check for simple types
// 사용 사례: 단순 타입의 빠른 동등성 확인
//
// Example:
//   config1 := map[string]int{"timeout": 30, "retries": 3}
//   config2 := map[string]int{"timeout": 30, "retries": 3}
//   config3 := map[string]int{"timeout": 60, "retries": 3}
//   EqualMaps(config1, config2) // true
//   EqualMaps(config1, config3) // false
//
// EqualFunc(m1, m2, eq) checks map equality using a custom comparator function.
// This allows comparison of non-comparable types (slices, structs, functions, etc.)
// by providing a custom equality function. The comparator receives two values and
// returns true if they should be considered equal.
//
// EqualFunc(m1, m2, eq)는 커스텀 비교자 함수를 사용하여 맵 동등성을 확인합니다.
// 커스텀 동등성 함수를 제공하여 비교 불가능 타입(슬라이스, 구조체, 함수 등)의
// 비교를 허용합니다. 비교자는 두 값을 받고 동일하게 간주되어야 하면 true를 반환합니다.
//
// Time Complexity: O(n) + O(n*c) where c = cost of comparator function
// Space Complexity: O(1) for comparison, depends on comparator
// Comparator Signature: func(V, V) bool
// Flexibility: Works with any value type
// Early Termination: Returns false on first inequality
// Custom Logic: Supports deep equality, fuzzy matching, tolerance-based comparison
//
// 시간 복잡도: O(n) + O(n*c) 여기서 c = 비교자 함수 비용
// 공간 복잡도: 비교는 O(1), 비교자에 따라 다름
// 비교자 시그니처: func(V, V) bool
// 유연성: 모든 값 타입과 함께 작동
// 조기 종료: 첫 불일치에서 false 반환
// 커스텀 로직: 깊은 동등성, 퍼지 매칭, 허용 오차 기반 비교 지원
//
// Use Case: Compare maps with complex value types (slices, nested structs)
// 사용 사례: 복잡한 값 타입을 가진 맵 비교 (슬라이스, 중첩 구조체)
//
// Example:
//   m1 := map[string][]int{"a": {1, 2, 3}, "b": {4, 5}}
//   m2 := map[string][]int{"a": {1, 2, 3}, "b": {4, 5}}
//   equal := EqualFunc(m1, m2, func(v1, v2 []int) bool {
//       if len(v1) != len(v2) {
//           return false
//       }
//       for i := range v1 {
//           if v1[i] != v2[i] {
//               return false
//           }
//       }
//       return true
//   })
//   // equal = true
//
// DIFFERENCE DETECTION OPERATIONS | 차이 감지 작업
//
// Diff(m1, m2) returns a map containing all entries that differ between two maps.
// An entry is included in the result if:
//   - Key exists in both maps but values differ (uses m2's value)
//   - Key exists only in m1 (uses m1's value)
//   - Key exists only in m2 (uses m2's value)
//
// Diff(m1, m2)는 두 맵 간에 다른 모든 항목을 포함하는 맵을 반환합니다.
// 다음 경우 항목이 결과에 포함됩니다:
//   - 키가 두 맵 모두에 존재하지만 값이 다름 (m2의 값 사용)
//   - 키가 m1에만 존재 (m1의 값 사용)
//   - 키가 m2에만 존재 (m2의 값 사용)
//
// Time Complexity: O(n+m) where n = size of m1, m = size of m2
// Space Complexity: O(d) where d = number of differences
// Value Selection: Modified entries use m2's value, unique entries use their own
// Symmetry: Not symmetric - Diff(A, B) ≠ Diff(B, A) for modified entries
// Empty Result: Returns empty map if maps are equal
// Nil Handling: Treats nil as empty map
//
// 시간 복잡도: O(n+m) 여기서 n = m1 크기, m = m2 크기
// 공간 복잡도: O(d) 여기서 d = 차이 개수
// 값 선택: 수정된 항목은 m2의 값, 고유 항목은 자신의 값 사용
// 대칭성: 비대칭 - 수정된 항목에 대해 Diff(A, B) ≠ Diff(B, A)
// 빈 결과: 맵이 동일하면 빈 맵 반환
// Nil 처리: nil을 빈 맵으로 처리
//
// Use Case: Find all changes between map versions
// 사용 사례: 맵 버전 간 모든 변경사항 찾기
//
// Example:
//   old := map[string]int{"a": 1, "b": 2, "c": 3}
//   new := map[string]int{"a": 1, "b": 20, "d": 4}
//   diff := Diff(old, new)
//   // diff = {"b": 20, "c": 3, "d": 4}
//   // b: modified (new value), c: removed (old value), d: added (new value)
//
// DiffKeys(m1, m2) returns a slice of keys that differ between two maps.
// Similar to Diff but returns only keys, not values. More efficient when you only
// need to know which keys changed, not what the values are.
//
// DiffKeys(m1, m2)는 두 맵 간에 다른 키를 슬라이스로 반환합니다.
// Diff와 유사하지만 값이 아닌 키만 반환합니다. 어떤 키가 변경되었는지만 알면 되고
// 값이 무엇인지는 필요하지 않을 때 더 효율적입니다.
//
// Time Complexity: O(n+m)
// Space Complexity: O(d) where d = number of differing keys
// Deduplication: Uses 'seen' map to avoid duplicate keys
// Order: Non-deterministic (map iteration order)
// Empty Result: Returns empty slice if maps are equal
// Nil Handling: Treats nil as empty map
//
// 시간 복잡도: O(n+m)
// 공간 복잡도: O(d) 여기서 d = 다른 키 개수
// 중복 제거: 중복 키를 피하기 위해 'seen' 맵 사용
// 순서: 비결정적 (맵 순회 순서)
// 빈 결과: 맵이 동일하면 빈 슬라이스 반환
// Nil 처리: nil을 빈 맵으로 처리
//
// Use Case: Detect which configuration keys changed
// 사용 사례: 어떤 구성 키가 변경되었는지 감지
//
// Example:
//   state1 := map[string]int{"a": 1, "b": 2, "c": 3}
//   state2 := map[string]int{"a": 1, "b": 20, "d": 4}
//   changed := DiffKeys(state1, state2)
//   // changed = []string{"b", "c", "d"} (order may vary)
//
// Compare(m1, m2) performs a detailed three-way comparison of two maps.
// Returns three separate maps:
//   - added: entries in m2 but not in m1 (new keys)
//   - removed: entries in m1 but not in m2 (deleted keys)
//   - modified: entries in both maps with different values (changed keys, m2's values)
//
// Compare(m1, m2)는 두 맵의 상세한 3방향 비교를 수행합니다.
// 세 개의 별도 맵을 반환합니다:
//   - added: m2에는 있지만 m1에는 없는 항목 (새 키)
//   - removed: m1에는 있지만 m2에는 없는 항목 (삭제된 키)
//   - modified: 두 맵 모두에 다른 값을 가진 항목 (변경된 키, m2의 값)
//
// Time Complexity: O(n+m)
// Space Complexity: O(a+r+mo) where a=added, r=removed, mo=modified
// Return Values: Three maps (added, removed, modified)
// Semantic Clarity: Clear categorization of changes
// Git-like: Similar to git diff (additions, deletions, modifications)
// Empty Maps: Returns three empty maps if inputs are equal
//
// 시간 복잡도: O(n+m)
// 공간 복잡도: O(a+r+mo) 여기서 a=추가, r=제거, mo=수정
// 반환값: 세 맵 (added, removed, modified)
// 의미적 명확성: 명확한 변경 분류
// Git 같은: git diff와 유사 (추가, 삭제, 수정)
// 빈 맵: 입력이 동일하면 세 개의 빈 맵 반환
//
// Use Case: Version control, change tracking, audit logging
// 사용 사례: 버전 관리, 변경 추적, 감사 로깅
//
// Example:
//   oldConfig := map[string]int{"timeout": 30, "retries": 3, "cache": 100}
//   newConfig := map[string]int{"timeout": 60, "retries": 3, "maxConn": 50}
//   added, removed, modified := Compare(oldConfig, newConfig)
//   // added = {"maxConn": 50}
//   // removed = {"cache": 100}
//   // modified = {"timeout": 60}
//
// KEY ANALYSIS OPERATIONS | 키 분석 작업
//
// CommonKeys(maps...) returns keys that exist in ALL input maps.
// This is similar to Intersection but returns only keys, not key-value pairs.
// Works with any number of maps (variadic).
//
// CommonKeys(maps...)는 모든 입력 맵에 존재하는 키를 반환합니다.
// Intersection과 유사하지만 키-값 쌍이 아닌 키만 반환합니다.
// 임의 개수의 맵과 함께 작동합니다 (가변 인자).
//
// Time Complexity: O(n*m) where n = number of maps, m = average map size
// Space Complexity: O(k) where k = number of common keys
// Variadic: Accepts any number of maps
// Empty Input: Returns empty slice
// Single Map: Returns all keys from that map
// Set Semantics: Intersection operation (∩) for keys
//
// 시간 복잡도: O(n*m) 여기서 n = 맵 개수, m = 평균 맵 크기
// 공간 복잡도: O(k) 여기서 k = 공통 키 개수
// 가변 인자: 임의 개수의 맵 허용
// 빈 입력: 빈 슬라이스 반환
// 단일 맵: 해당 맵의 모든 키 반환
// 집합 의미론: 키를 위한 교집합 연산 (∩)
//
// Use Case: Find shared configuration keys, common fields
// 사용 사례: 공유 구성 키 찾기, 공통 필드
//
// Example:
//   schema1 := map[string]string{"id": "int", "name": "string", "age": "int"}
//   schema2 := map[string]string{"name": "string", "age": "int", "email": "string"}
//   schema3 := map[string]string{"age": "int", "email": "string", "phone": "string"}
//   common := CommonKeys(schema1, schema2, schema3)
//   // common = []string{"age"} (only age in all three)
//
// AllKeys(maps...) returns all unique keys from all input maps (union of keys).
// Deduplicates keys that appear in multiple maps. Works with any number of maps.
//
// AllKeys(maps...)는 모든 입력 맵의 모든 고유 키를 반환합니다 (키의 합집합).
// 여러 맵에 나타나는 키를 중복 제거합니다. 임의 개수의 맵과 함께 작동합니다.
//
// Time Complexity: O(n*m) where n = number of maps, m = average map size
// Space Complexity: O(u) where u = number of unique keys
// Variadic: Accepts any number of maps
// Deduplication: Uses map for O(1) duplicate detection
// Order: Non-deterministic (depends on map iteration)
// Set Semantics: Union operation (∪) for keys
//
// 시간 복잡도: O(n*m) 여기서 n = 맵 개수, m = 평균 맵 크기
// 공간 복잡도: O(u) 여기서 u = 고유 키 개수
// 가변 인자: 임의 개수의 맵 허용
// 중복 제거: O(1) 중복 감지를 위해 맵 사용
// 순서: 비결정적 (맵 순회에 의존)
// 집합 의미론: 키를 위한 합집합 연산 (∪)
//
// Use Case: Collect all possible configuration keys, schema merging
// 사용 사례: 모든 가능한 구성 키 수집, 스키마 병합
//
// Example:
//   config1 := map[string]int{"timeout": 30, "retries": 3}
//   config2 := map[string]int{"retries": 5, "cache": 100}
//   config3 := map[string]int{"cache": 200, "maxConn": 50}
//   allKeys := AllKeys(config1, config2, config3)
//   // allKeys = []string{"timeout", "retries", "cache", "maxConn"} (order may vary)
//
// # Comparisons with Related Functions | 관련 함수와 비교
//
// EqualMaps vs. EqualFunc:
//   - EqualMaps: Built-in ==, faster, only comparable types
//   - EqualFunc: Custom comparator, slower, any types
//   - Use EqualMaps for simple types (int, string, etc.)
//   - Use EqualFunc for complex types (slices, structs, etc.)
//
// EqualMaps 대 EqualFunc:
//   - EqualMaps: 내장 ==, 더 빠름, 비교 가능 타입만
//   - EqualFunc: 커스텀 비교자, 더 느림, 모든 타입
//   - 단순 타입에 EqualMaps 사용 (int, string 등)
//   - 복잡한 타입에 EqualFunc 사용 (슬라이스, 구조체 등)
//
// EqualMaps vs. basic.Equal:
//   - Functionally identical (basic.Equal calls EqualMaps)
//   - EqualMaps is the implementation, basic.Equal is the alias
//   - Use whichever is more semantically clear
//
// EqualMaps 대 basic.Equal:
//   - 기능적으로 동일 (basic.Equal이 EqualMaps 호출)
//   - EqualMaps가 구현, basic.Equal이 별칭
//   - 의미적으로 더 명확한 것 사용
//
// Diff vs. Compare:
//   - Diff: Single map with all differences
//   - Compare: Three maps (added, removed, modified)
//   - Diff simpler when you don't need categorization
//   - Compare better for detailed change tracking
//
// Diff 대 Compare:
//   - Diff: 모든 차이를 가진 단일 맵
//   - Compare: 세 맵 (added, removed, modified)
//   - 분류가 필요하지 않을 때 Diff가 더 간단
//   - 상세한 변경 추적에 Compare가 더 좋음
//
// Diff vs. DiffKeys:
//   - Diff: Returns map with differing entries (keys + values)
//   - DiffKeys: Returns slice with differing keys only
//   - Use Diff when you need values
//   - Use DiffKeys for key-only change detection (more efficient)
//
// Diff 대 DiffKeys:
//   - Diff: 다른 항목을 가진 맵 반환 (키 + 값)
//   - DiffKeys: 다른 키만 가진 슬라이스 반환
//   - 값이 필요할 때 Diff 사용
//   - 키만의 변경 감지에 DiffKeys 사용 (더 효율적)
//
// CommonKeys vs. Intersection:
//   - CommonKeys: Returns []K (keys only)
//   - Intersection: Returns map[K]V (keys + values from first map)
//   - Use CommonKeys when values not needed
//   - Use Intersection when you need resulting map
//
// CommonKeys 대 Intersection:
//   - CommonKeys: []K 반환 (키만)
//   - Intersection: map[K]V 반환 (첫 맵의 키 + 값)
//   - 값이 필요하지 않을 때 CommonKeys 사용
//   - 결과 맵이 필요할 때 Intersection 사용
//
// AllKeys vs. Union (Merge):
//   - AllKeys: Returns []K (unique keys only)
//   - Union: Returns map[K]V (merged map with all entries)
//   - Use AllKeys for key enumeration
//   - Use Union for actual map merging
//
// AllKeys 대 Union (Merge):
//   - AllKeys: []K 반환 (고유 키만)
//   - Union: map[K]V 반환 (모든 항목이 있는 병합 맵)
//   - 키 나열에 AllKeys 사용
//   - 실제 맵 병합에 Union 사용
//
// Compare vs. merge.Difference:
//   - Compare: Three-way split (added, removed, modified)
//   - Difference: Keys in m1 but not m2 (removed only)
//   - Compare for full change analysis
//   - Difference for simple removal detection
//
// Compare 대 merge.Difference:
//   - Compare: 3방향 분할 (added, removed, modified)
//   - Difference: m1에는 있지만 m2에는 없는 키 (제거만)
//   - 전체 변경 분석에 Compare
//   - 단순 제거 감지에 Difference
//
// # Performance Characteristics | 성능 특성
//
// Time Complexities:
//   - O(n): EqualMaps, EqualFunc (n = map size)
//   - O(n+m): Diff, DiffKeys, Compare (n, m = sizes of two maps)
//   - O(n*m): CommonKeys, AllKeys (n = number of maps, m = average size)
//
// 시간 복잡도:
//   - O(n): EqualMaps, EqualFunc (n = 맵 크기)
//   - O(n+m): Diff, DiffKeys, Compare (n, m = 두 맵의 크기)
//   - O(n*m): CommonKeys, AllKeys (n = 맵 개수, m = 평균 크기)
//
// Space Complexities:
//   - O(1): EqualMaps, EqualFunc (no allocations)
//   - O(d): Diff, DiffKeys (d = number of differences)
//   - O(a+r+mo): Compare (a=added, r=removed, mo=modified)
//   - O(k): CommonKeys, AllKeys (k = result keys)
//
// 공간 복잡도:
//   - O(1): EqualMaps, EqualFunc (할당 없음)
//   - O(d): Diff, DiffKeys (d = 차이 개수)
//   - O(a+r+mo): Compare (a=추가, r=제거, mo=수정)
//   - O(k): CommonKeys, AllKeys (k = 결과 키)
//
// Optimization Tips:
//   - Use EqualMaps over EqualFunc when possible (no function call overhead)
//   - Check size equality before EqualMaps for quick inequality detection
//   - Use DiffKeys instead of Diff when values not needed
//   - Use Compare for detailed analysis, Diff for simple difference
//   - Cache CommonKeys/AllKeys results if used repeatedly
//
// 최적화 팁:
//   - 가능하면 EqualFunc 대신 EqualMaps 사용 (함수 호출 오버헤드 없음)
//   - 빠른 불일치 감지를 위해 EqualMaps 전 크기 동등성 확인
//   - 값이 필요하지 않을 때 Diff 대신 DiffKeys 사용
//   - 상세 분석에 Compare, 단순 차이에 Diff 사용
//   - 반복 사용될 경우 CommonKeys/AllKeys 결과 캐시
//
// # Common Usage Patterns | 일반적인 사용 패턴
//
// 1. Configuration Change Detection | 구성 변경 감지
//
//	oldConfig := map[string]int{"timeout": 30, "retries": 3, "cache": 100}
//	newConfig := map[string]int{"timeout": 60, "retries": 3, "maxConn": 50}
//	added, removed, modified := maputil.Compare(oldConfig, newConfig)
//	fmt.Printf("Added: %v\n", added)       // {"maxConn": 50}
//	fmt.Printf("Removed: %v\n", removed)   // {"cache": 100}
//	fmt.Printf("Modified: %v\n", modified) // {"timeout": 60}
//
// 2. Validate Configuration Equality | 구성 동등성 검증
//
//	expected := map[string]string{"host": "localhost", "port": "8080"}
//	actual := map[string]string{"host": "localhost", "port": "8080"}
//	if !maputil.EqualMaps(expected, actual) {
//	    log.Fatal("Configuration mismatch!")
//	}
//
// 3. Compare Slices in Maps | 맵의 슬라이스 비교
//
//	permissions1 := map[string][]string{
//	    "admin": {"read", "write", "delete"},
//	    "user":  {"read"},
//	}
//	permissions2 := map[string][]string{
//	    "admin": {"read", "write", "delete"},
//	    "user":  {"read"},
//	}
//	equal := maputil.EqualFunc(permissions1, permissions2, func(v1, v2 []string) bool {
//	    if len(v1) != len(v2) {
//	        return false
//	    }
//	    for i := range v1 {
//	        if v1[i] != v2[i] {
//	            return false
//	        }
//	    }
//	    return true
//	})
//	// equal = true
//
// 4. Find Changed Configuration Keys | 변경된 구성 키 찾기
//
//	state1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	state2 := map[string]int{"a": 1, "b": 20, "d": 4}
//	changedKeys := maputil.DiffKeys(state1, state2)
//	for _, key := range changedKeys {
//	    fmt.Printf("Key changed: %s\n", key)
//	}
//	// Output: Key changed: b, Key changed: c, Key changed: d
//
// 5. Find Common Schema Fields | 공통 스키마 필드 찾기
//
//	schema1 := map[string]string{"id": "int", "name": "string", "age": "int"}
//	schema2 := map[string]string{"name": "string", "age": "int", "email": "string"}
//	schema3 := map[string]string{"age": "int", "email": "string", "phone": "string"}
//	commonFields := maputil.CommonKeys(schema1, schema2, schema3)
//	// commonFields = ["age"] - only age in all schemas
//
// 6. Collect All Possible Keys | 모든 가능한 키 수집
//
//	defaults := map[string]int{"timeout": 30, "retries": 3}
//	userConfig := map[string]int{"retries": 5, "cache": 100}
//	envConfig := map[string]int{"cache": 200, "debug": 1}
//	allKeys := maputil.AllKeys(defaults, userConfig, envConfig)
//	// allKeys = ["timeout", "retries", "cache", "debug"]
//	// Use to validate all possible configuration keys
//
// 7. Audit Log Changes | 감사 로그 변경사항
//
//	before := map[string]interface{}{"status": "active", "balance": 1000}
//	after := map[string]interface{}{"status": "inactive", "balance": 500, "locked": true}
//	diff := maputil.Diff(before, after)
//	for key, value := range diff {
//	    log.Printf("Field %s changed to %v", key, value)
//	}
//
// 8. Detect Missing Required Keys | 필수 키 누락 감지
//
//	requiredKeys := map[string]bool{"host": true, "port": true, "user": true}
//	actualConfig := map[string]bool{"host": true, "port": true}
//	missing := maputil.Difference(requiredKeys, actualConfig) // from merge.go
//	if len(missing) > 0 {
//	    log.Fatalf("Missing required keys: %v", maputil.KeysSlice(missing))
//	}
//
// 9. Version Comparison | 버전 비교
//
//	v1Features := map[string]bool{"feature_a": true, "feature_b": true}
//	v2Features := map[string]bool{"feature_b": true, "feature_c": true}
//	added, removed, _ := maputil.Compare(v1Features, v2Features)
//	fmt.Printf("New features: %v\n", maputil.KeysSlice(added))
//	fmt.Printf("Deprecated: %v\n", maputil.KeysSlice(removed))
//
// 10. Fuzzy Equality with Tolerance | 허용 오차를 사용한 퍼지 동등성
//
//	measurements1 := map[string]float64{"temp": 25.0, "humidity": 60.0}
//	measurements2 := map[string]float64{"temp": 25.05, "humidity": 60.1}
//	tolerance := 0.2
//	equal := maputil.EqualFunc(measurements1, measurements2, func(v1, v2 float64) bool {
//	    return math.Abs(v1-v2) <= tolerance
//	})
//	// equal = true (within tolerance)
//
// # Edge Cases and Nil Handling | 엣지 케이스와 Nil 처리
//
// Empty Maps:
//   - EqualMaps: Two empty maps are equal
//   - Diff, DiffKeys, Compare: Return empty results
//   - CommonKeys: Returns empty slice
//   - AllKeys: Returns empty slice
//
// 빈 맵:
//   - EqualMaps: 두 빈 맵은 동일
//   - Diff, DiffKeys, Compare: 빈 결과 반환
//   - CommonKeys: 빈 슬라이스 반환
//   - AllKeys: 빈 슬라이스 반환
//
// Nil Maps:
//   - All functions treat nil as empty maps
//   - EqualMaps: nil == nil and nil == empty both true
//   - Diff/Compare: Treats nil as no entries
//
// Nil 맵:
//   - 모든 함수가 nil을 빈 맵으로 처리
//   - EqualMaps: nil == nil과 nil == empty 모두 true
//   - Diff/Compare: nil을 항목 없음으로 처리
//
// Identical Maps:
//   - EqualMaps: Returns true
//   - Diff, DiffKeys: Return empty results
//   - Compare: Returns three empty maps
//
// 동일한 맵:
//   - EqualMaps: true 반환
//   - Diff, DiffKeys: 빈 결과 반환
//   - Compare: 세 개의 빈 맵 반환
//
// Completely Different Maps:
//   - EqualMaps: Returns false
//   - Diff: Returns all entries from both maps
//   - DiffKeys: Returns all keys from both maps
//   - Compare: added = m2, removed = m1, modified = empty
//
// 완전히 다른 맵:
//   - EqualMaps: false 반환
//   - Diff: 두 맵의 모든 항목 반환
//   - DiffKeys: 두 맵의 모든 키 반환
//   - Compare: added = m2, removed = m1, modified = 빈 맵
//
// Single Map Input:
//   - CommonKeys: Returns all keys from that map
//   - AllKeys: Returns all keys from that map
//
// 단일 맵 입력:
//   - CommonKeys: 해당 맵의 모든 키 반환
//   - AllKeys: 해당 맵의 모든 키 반환
//
// # Thread Safety | 스레드 안전성
//
// Read-Only Operations (Safe for Concurrent Reads):
//   - All functions in this file are read-only
//   - EqualMaps, EqualFunc, Diff, DiffKeys, Compare, CommonKeys, AllKeys
//   - Safe when input maps have concurrent readers
//   - Do not modify input maps
//
// 읽기 전용 작업 (동시 읽기 안전):
//   - 이 파일의 모든 함수는 읽기 전용
//   - EqualMaps, EqualFunc, Diff, DiffKeys, Compare, CommonKeys, AllKeys
//   - 입력 맵에 동시 읽기가 있을 때 안전
//   - 입력 맵을 수정하지 않음
//
// Concurrent Modification Warning:
//   - Do not modify input maps while comparison functions read them
//   - Use sync.RWMutex for concurrent access patterns
//
// 동시 수정 경고:
//   - 비교 함수가 읽는 동안 입력 맵 수정 금지
//   - 동시 접근 패턴에 sync.RWMutex 사용
//
// # See Also | 참고
//
// Related files in maputil package:
//   - basic.go: Equal function (alias for EqualMaps)
//   - merge.go: Difference, SymmetricDifference for set operations
//   - keys.go: KeysSlice for extracting keys from diff results
//   - filter.go: Pick, Omit for key-based filtering
//
// maputil 패키지의 관련 파일:
//   - basic.go: Equal 함수 (EqualMaps의 별칭)
//   - merge.go: 집합 연산을 위한 Difference, SymmetricDifference
//   - keys.go: diff 결과에서 키 추출을 위한 KeysSlice
//   - filter.go: 키 기반 필터링을 위한 Pick, Omit

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
