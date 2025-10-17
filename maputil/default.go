package maputil

// Package maputil/default.go provides operations for handling default values in maps.
// This file contains functions for setting, retrieving, and applying default values,
// enabling safe initialization and fallback behavior in map operations.
//
// maputil/default.go 패키지는 맵의 기본값 처리를 위한 작업을 제공합니다.
// 이 파일은 기본값 설정, 검색, 적용을 위한 함수들을 포함하여, 맵 작업에서 안전한
// 초기화와 폴백 동작을 가능하게 합니다.
//
// # Overview | 개요
//
// The default.go file provides 3 default value operations:
//
// default.go 파일은 3개 기본값 작업을 제공합니다:
//
// 1. MUTABLE OPERATIONS | 가변 작업
//   - GetOrSet: Get existing value or set default (O(1))
//   - SetDefault: Set default only if key missing (O(1))
//
// 2. IMMUTABLE OPERATIONS | 불변 작업
//   - Defaults: Create new map with defaults applied (O(n+d))
//
// # Design Principles | 설계 원칙
//
// 1. LAZY INITIALIZATION | 지연 초기화
//   - GetOrSet: Initializes value on first access
//   - Avoids pre-allocating unused keys
//   - Efficient for sparse maps with many possible keys
//   - Common in caching and memoization patterns
//
//   GetOrSet: 첫 접근 시 값 초기화
//   사용하지 않는 키를 미리 할당하지 않음
//   많은 가능한 키를 가진 희소 맵에 효율적
//   캐싱과 메모이제이션 패턴에서 일반적
//
// 2. NON-OVERWRITING | 비덮어쓰기
//   - SetDefault, Defaults: Never overwrite existing values
//   - Original values have priority over defaults
//   - Safe for applying fallback configurations
//   - Prevents accidental data loss
//
//   SetDefault, Defaults: 기존 값을 절대 덮어쓰지 않음
//   원본 값이 기본값보다 우선순위
//   폴백 구성 적용에 안전
//   실수로 인한 데이터 손실 방지
//
// 3. MUTATION VS. IMMUTABILITY | 가변 대 불변
//   - GetOrSet, SetDefault: Mutate input map (in-place)
//   - Defaults: Create new map (immutable)
//   - Choose based on whether side effects are acceptable
//   - Mutable operations more efficient, immutable operations safer
//
//   GetOrSet, SetDefault: 입력 맵 변경 (제자리)
//   Defaults: 새 맵 생성 (불변)
//   부작용이 허용되는지에 따라 선택
//   가변 작업이 더 효율적, 불변 작업이 더 안전
//
// 4. SINGLE VS. BULK DEFAULTS | 단일 대 대량 기본값
//   - GetOrSet, SetDefault: Single key-value pair
//   - Defaults: Multiple key-value pairs from map
//   - Use single operations in hot loops
//   - Use bulk operation for initial configuration
//
//   GetOrSet, SetDefault: 단일 키-값 쌍
//   Defaults: 맵의 여러 키-값 쌍
//   핫 루프에서 단일 작업 사용
//   초기 구성에 대량 작업 사용
//
// # Function Categories | 함수 카테고리
//
// MUTABLE OPERATIONS | 가변 작업
//
// GetOrSet(m, key, defaultValue) retrieves a value if it exists, or sets and
// returns the default value if the key is missing. This is a mutating operation
// that modifies the input map when the key doesn't exist.
//
// GetOrSet(m, key, defaultValue)는 값이 존재하면 검색하거나, 키가 누락되면
// 기본값을 설정하고 반환합니다. 키가 존재하지 않을 때 입력 맵을 수정하는 가변 작업입니다.
//
// Time Complexity: O(1) - hash lookup and optional insertion
// Space Complexity: O(1) - modifies existing map
// Mutating: YES - modifies input map when key missing
// Return Value: Existing value or default value
// Side Effect: May insert new key-value pair
// Use Case: Caching, lazy initialization, memoization
//
// 시간 복잡도: O(1) - 해시 조회 및 선택적 삽입
// 공간 복잡도: O(1) - 기존 맵 수정
// 가변: 예 - 키 누락 시 입력 맵 수정
// 반환값: 기존 값 또는 기본값
// 부작용: 새 키-값 쌍을 삽입할 수 있음
// 사용 사례: 캐싱, 지연 초기화, 메모이제이션
//
// Behavior:
//   - If key exists: Return existing value, no modification
//   - If key missing: Insert (key, defaultValue), return defaultValue
//
// 동작:
//   - 키 존재: 기존 값 반환, 수정 없음
//   - 키 누락: (키, 기본값) 삽입, 기본값 반환
//
// Example:
//   cache := map[string]int{"a": 1}
//   val1 := GetOrSet(cache, "a", 10) // val1 = 1, cache unchanged
//   val2 := GetOrSet(cache, "b", 10) // val2 = 10, cache now {"a": 1, "b": 10}
//
// Comparison with GetOr (basic.go):
//   - GetOr: Read-only, returns default without modifying map
//   - GetOrSet: Mutating, inserts default when missing
//   - GetOr for read-only access, GetOrSet for initialization
//
// GetOr (basic.go)와 비교:
//   - GetOr: 읽기 전용, 맵을 수정하지 않고 기본값 반환
//   - GetOrSet: 가변, 누락 시 기본값 삽입
//   - 읽기 전용 접근에 GetOr, 초기화에 GetOrSet
//
// SetDefault(m, key, defaultValue) sets a key to a default value only if the key
// doesn't already exist. Returns true if the key was set (new), false if the key
// already existed (no change). This is a mutating operation.
//
// SetDefault(m, key, defaultValue)는 키가 아직 존재하지 않는 경우에만 키를 기본값으로
// 설정합니다. 키가 설정되었으면 true (새로 생성), 키가 이미 존재했으면 false (변경 없음)를
// 반환합니다. 가변 작업입니다.
//
// Time Complexity: O(1) - hash lookup and optional insertion
// Space Complexity: O(1) - modifies existing map
// Mutating: YES - modifies input map when key missing
// Return Value: true if set (new key), false if exists (no change)
// Side Effect: May insert new key-value pair
// Non-Overwriting: Never overwrites existing values
//
// 시간 복잡도: O(1) - 해시 조회 및 선택적 삽입
// 공간 복잡도: O(1) - 기존 맵 수정
// 가변: 예 - 키 누락 시 입력 맵 수정
// 반환값: 설정되었으면 true (새 키), 존재하면 false (변경 없음)
// 부작용: 새 키-값 쌍을 삽입할 수 있음
// 비덮어쓰기: 기존 값을 절대 덮어쓰지 않음
//
// Behavior:
//   - If key exists: Return false, no modification
//   - If key missing: Insert (key, defaultValue), return true
//
// 동작:
//   - 키 존재: false 반환, 수정 없음
//   - 키 누락: (키, 기본값) 삽입, true 반환
//
// Example:
//   config := map[string]string{"host": "localhost"}
//   set1 := SetDefault(config, "port", "8080")    // set1 = true, adds port
//   set2 := SetDefault(config, "host", "0.0.0.0") // set2 = false, no change
//   // config = {"host": "localhost", "port": "8080"}
//
// Use Case: Safe initialization, conditional defaults, idempotent setup
// 사용 사례: 안전한 초기화, 조건부 기본값, 멱등 설정
//
// Comparison with GetOrSet:
//   - GetOrSet: Returns value (existing or default)
//   - SetDefault: Returns bool (whether key was set)
//   - GetOrSet when you need the value immediately
//   - SetDefault when you only care if initialization happened
//
// GetOrSet과 비교:
//   - GetOrSet: 값 반환 (기존 또는 기본값)
//   - SetDefault: bool 반환 (키가 설정되었는지)
//   - 값이 즉시 필요할 때 GetOrSet
//   - 초기화 발생 여부만 관심 있을 때 SetDefault
//
// IMMUTABLE OPERATIONS | 불변 작업
//
// Defaults(m, defaults) creates a new map that combines the original map with
// default values. Keys in the original map take precedence over keys in the
// defaults map. This is an immutable operation that doesn't modify inputs.
//
// Defaults(m, defaults)는 원본 맵과 기본값을 결합한 새 맵을 생성합니다.
// 원본 맵의 키가 defaults 맵의 키보다 우선합니다. 입력을 수정하지 않는 불변 작업입니다.
//
// Time Complexity: O(n+d) where n = size of m, d = size of defaults
// Space Complexity: O(n+d) - creates new map
// Mutating: NO - creates new map, originals unchanged
// Return Value: New map with defaults applied
// Precedence: Original map values > Default values
// Use Case: Configuration layering, fallback values
//
// 시간 복잡도: O(n+d) 여기서 n = m 크기, d = defaults 크기
// 공간 복잡도: O(n+d) - 새 맵 생성
// 가변: 아니오 - 새 맵 생성, 원본 변경 없음
// 반환값: 기본값이 적용된 새 맵
// 우선순위: 원본 맵 값 > 기본값
// 사용 사례: 구성 계층화, 폴백 값
//
// Behavior:
//   - Creates new map with all defaults keys
//   - Overwrites with all original map keys (original wins)
//   - Result has union of keys from both maps
//
// 동작:
//   - 모든 defaults 키를 가진 새 맵 생성
//   - 모든 원본 맵 키로 덮어쓰기 (원본 우선)
//   - 결과는 두 맵의 키 합집합
//
// Example:
//   config := map[string]int{"timeout": 60}
//   defaults := map[string]int{"timeout": 30, "retries": 3, "cache": 100}
//   result := Defaults(config, defaults)
//   // result = {"timeout": 60, "retries": 3, "cache": 100}
//   // timeout from config (60), retries and cache from defaults
//
// Comparison with Merge:
//   - Defaults: Original map values take precedence
//   - Merge: Last (rightmost) map values take precedence
//   - Defaults(config, defaults) ≈ Merge(defaults, config)
//   - Use Defaults for clearer "original + fallbacks" semantics
//
// Merge와 비교:
//   - Defaults: 원본 맵 값 우선
//   - Merge: 마지막 (가장 오른쪽) 맵 값 우선
//   - Defaults(config, defaults) ≈ Merge(defaults, config)
//   - 더 명확한 "원본 + 폴백" 의미론을 위해 Defaults 사용
//
// # Comparisons with Related Functions | 관련 함수와 비교
//
// GetOrSet vs. GetOr (basic.go):
//   - GetOrSet: Mutating, inserts default when missing
//   - GetOr: Read-only, returns default without modifying map
//   - GetOrSet for initialization, GetOr for safe reading
//   - GetOrSet has side effects, GetOr has none
//
// GetOrSet 대 GetOr (basic.go):
//   - GetOrSet: 가변, 누락 시 기본값 삽입
//   - GetOr: 읽기 전용, 맵을 수정하지 않고 기본값 반환
//   - 초기화에 GetOrSet, 안전한 읽기에 GetOr
//   - GetOrSet는 부작용 있음, GetOr는 없음
//
// GetOrSet vs. SetDefault:
//   - GetOrSet: Returns V (value), always modifies when missing
//   - SetDefault: Returns bool (was set?), always modifies when missing
//   - GetOrSet when you need the value immediately
//   - SetDefault when you only care about initialization status
//
// GetOrSet 대 SetDefault:
//   - GetOrSet: V (값) 반환, 누락 시 항상 수정
//   - SetDefault: bool (설정되었는지?) 반환, 누락 시 항상 수정
//   - 값이 즉시 필요할 때 GetOrSet
//   - 초기화 상태만 관심 있을 때 SetDefault
//
// SetDefault vs. basic.Set:
//   - SetDefault: Only sets if key doesn't exist (conditional)
//   - basic.Set: Always overwrites (unconditional)
//   - SetDefault for safe defaults, basic.Set for forced updates
//
// SetDefault 대 basic.Set:
//   - SetDefault: 키가 존재하지 않을 때만 설정 (조건부)
//   - basic.Set: 항상 덮어쓰기 (무조건)
//   - 안전한 기본값에 SetDefault, 강제 업데이트에 basic.Set
//
// Defaults vs. Merge:
//   - Defaults: Original map precedence (config wins over defaults)
//   - Merge: Last map precedence (rightmost wins)
//   - Defaults(m, d) ≈ Merge(d, m) but clearer semantics
//   - Use Defaults when semantics matter
//
// Defaults 대 Merge:
//   - Defaults: 원본 맵 우선 (config가 defaults보다 우선)
//   - Merge: 마지막 맵 우선 (가장 오른쪽 우선)
//   - Defaults(m, d) ≈ Merge(d, m) 하지만 의미론이 더 명확
//   - 의미론이 중요할 때 Defaults 사용
//
// Defaults vs. MergeWith:
//   - Defaults: Simple precedence (original > defaults)
//   - MergeWith: Custom conflict resolution function
//   - Defaults for simple fallbacks, MergeWith for custom logic
//
// Defaults 대 MergeWith:
//   - Defaults: 단순 우선순위 (원본 > 기본값)
//   - MergeWith: 커스텀 충돌 해결 함수
//   - 단순 폴백에 Defaults, 커스텀 로직에 MergeWith
//
// GetOrSet vs. Caching Libraries:
//   - GetOrSet: Simple, inline caching
//   - Cache Libraries: Thread-safe, expiration, eviction policies
//   - GetOrSet for simple single-threaded caching
//   - Cache Libraries for production concurrent caching
//
// GetOrSet 대 캐싱 라이브러리:
//   - GetOrSet: 간단한 인라인 캐싱
//   - 캐시 라이브러리: 스레드 안전, 만료, 제거 정책
//   - 단순 단일 스레드 캐싱에 GetOrSet
//   - 프로덕션 동시 캐싱에 캐시 라이브러리
//
// # Performance Characteristics | 성능 특성
//
// Time Complexities:
//   - O(1): GetOrSet, SetDefault (hash lookup + optional insert)
//   - O(n+d): Defaults (iterates both maps)
//
// 시간 복잡도:
//   - O(1): GetOrSet, SetDefault (해시 조회 + 선택적 삽입)
//   - O(n+d): Defaults (두 맵 모두 순회)
//
// Space Complexities:
//   - O(1): GetOrSet, SetDefault (modifies in-place)
//   - O(n+d): Defaults (creates new map)
//
// 공간 복잡도:
//   - O(1): GetOrSet, SetDefault (제자리 수정)
//   - O(n+d): Defaults (새 맵 생성)
//
// Optimization Tips:
//   - Use GetOrSet/SetDefault in hot loops (O(1) vs O(n+d) for Defaults)
//   - Use Defaults for initial setup, mutable operations for runtime
//   - Pre-allocate maps with capacity if size known
//   - Cache Defaults result if defaults don't change
//
// 최적화 팁:
//   - 핫 루프에서 GetOrSet/SetDefault 사용 (Defaults의 O(n+d) 대 O(1))
//   - 초기 설정에 Defaults, 런타임에 가변 작업
//   - 크기가 알려진 경우 용량으로 맵 사전 할당
//   - 기본값이 변경되지 않으면 Defaults 결과 캐시
//
// # Common Usage Patterns | 일반적인 사용 패턴
//
// 1. Lazy Cache Initialization | 지연 캐시 초기화
//
//	cache := make(map[string]ExpensiveResult)
//	func compute(key string) ExpensiveResult {
//	    return maputil.GetOrSet(cache, key, func() ExpensiveResult {
//	        // Expensive computation here
//	        return computeExpensive(key)
//	    }())
//	}
//	// First call computes and caches, subsequent calls retrieve from cache
//
// 2. Configuration with Defaults | 기본값이 있는 구성
//
//	userConfig := map[string]int{"timeout": 60}
//	systemDefaults := map[string]int{
//	    "timeout": 30,
//	    "retries": 3,
//	    "maxConn": 100,
//	}
//	finalConfig := maputil.Defaults(userConfig, systemDefaults)
//	// finalConfig = {"timeout": 60, "retries": 3, "maxConn": 100}
//
// 3. Safe Initialization of Nested Maps | 중첩 맵의 안전한 초기화
//
//	stats := make(map[string]map[string]int)
//	func recordStat(category, name string, value int) {
//	    categoryStats := maputil.GetOrSet(stats, category, make(map[string]int))
//	    categoryStats[name] = value
//	}
//
// 4. Idempotent Configuration Setup | 멱등 구성 설정
//
//	config := make(map[string]string)
//	maputil.SetDefault(config, "host", "localhost")
//	maputil.SetDefault(config, "port", "8080")
//	maputil.SetDefault(config, "host", "0.0.0.0") // No effect, host already set
//	// config = {"host": "localhost", "port": "8080"}
//
// 5. Multi-Layer Configuration | 다중 계층 구성
//
//	hardcodedDefaults := map[string]int{"timeout": 10, "retries": 1}
//	systemDefaults := map[string]int{"timeout": 30, "retries": 3}
//	userConfig := map[string]int{"timeout": 60}
//	finalConfig := maputil.Defaults(
//	    maputil.Defaults(userConfig, systemDefaults),
//	    hardcodedDefaults,
//	)
//	// Priority: userConfig > systemDefaults > hardcodedDefaults
//
// 6. Memoization Pattern | 메모이제이션 패턴
//
//	fibCache := map[int]int{0: 0, 1: 1}
//	func fib(n int) int {
//	    if val, ok := fibCache[n]; ok {
//	        return val
//	    }
//	    result := fib(n-1) + fib(n-2)
//	    fibCache[n] = result
//	    return result
//	}
//	// More concise with GetOrSet:
//	func fibWithGetOrSet(n int) int {
//	    if n <= 1 {
//	        return n
//	    }
//	    return maputil.GetOrSet(fibCache, n, fibWithGetOrSet(n-1)+fibWithGetOrSet(n-2))
//	}
//
// 7. Template Rendering with Defaults | 기본값이 있는 템플릿 렌더링
//
//	templateVars := map[string]string{"title": "My Page"}
//	defaultVars := map[string]string{
//	    "title":       "Default Title",
//	    "description": "Default Description",
//	    "author":      "Unknown",
//	}
//	allVars := maputil.Defaults(templateVars, defaultVars)
//	renderTemplate(template, allVars)
//
// 8. Counting with Default Zero | 기본 제로로 카운팅
//
//	counts := make(map[string]int)
//	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
//	for _, word := range words {
//	    counts[word] = maputil.GetOrSet(counts, word, 0) + 1
//	}
//	// counts = {"apple": 3, "banana": 2, "cherry": 1}
//
// 9. Feature Flags with Defaults | 기본값이 있는 기능 플래그
//
//	userFlags := map[string]bool{"new_ui": true}
//	defaultFlags := map[string]bool{
//	    "new_ui":       false,
//	    "dark_mode":    false,
//	    "experimental": false,
//	}
//	activeFlags := maputil.Defaults(userFlags, defaultFlags)
//	if activeFlags["new_ui"] {
//	    // Show new UI
//	}
//
// 10. Conditional API Parameter Defaults | 조건부 API 매개변수 기본값
//
//	params := map[string]string{"format": "json"}
//	maputil.SetDefault(params, "format", "xml")   // No effect, format already set
//	maputil.SetDefault(params, "version", "v1")   // Sets version=v1
//	maputil.SetDefault(params, "compress", "true") // Sets compress=true
//	// params = {"format": "json", "version": "v1", "compress": "true"}
//
// # Edge Cases and Nil Handling | 엣지 케이스와 Nil 처리
//
// Nil Maps:
//   - GetOrSet: Panics (cannot insert into nil map)
//   - SetDefault: Panics (cannot insert into nil map)
//   - Defaults: If m is nil, returns copy of defaults; if defaults is nil, returns copy of m
//
// Nil 맵:
//   - GetOrSet: 패닉 (nil 맵에 삽입 불가)
//   - SetDefault: 패닉 (nil 맵에 삽입 불가)
//   - Defaults: m이 nil이면 defaults 복사본 반환; defaults가 nil이면 m 복사본 반환
//
// Empty Maps:
//   - GetOrSet: Inserts default if key missing
//   - SetDefault: Inserts default if key missing
//   - Defaults: If m empty, returns copy of defaults
//
// 빈 맵:
//   - GetOrSet: 키 누락 시 기본값 삽입
//   - SetDefault: 키 누락 시 기본값 삽입
//   - Defaults: m이 비었으면 defaults 복사본 반환
//
// Zero Values as Defaults:
//   - Valid to use zero values (0, "", false, nil)
//   - Functions don't distinguish between zero and non-zero defaults
//
// 기본값으로서의 제로값:
//   - 제로값 사용 유효 (0, "", false, nil)
//   - 함수는 제로와 비제로 기본값을 구별하지 않음
//
// Existing Zero Values:
//   - GetOrSet: Returns existing zero value, doesn't overwrite
//   - SetDefault: Doesn't set default if zero value already exists
//   - Zero value ≠ missing key
//
// 기존 제로값:
//   - GetOrSet: 기존 제로값 반환, 덮어쓰지 않음
//   - SetDefault: 제로값이 이미 존재하면 기본값 설정 안 함
//   - 제로값 ≠ 누락 키
//
// # Thread Safety | 스레드 안전성
//
// Mutable Operations (NOT Safe):
//   - GetOrSet: Modifies map, requires external synchronization
//   - SetDefault: Modifies map, requires external synchronization
//   - Use sync.Mutex or sync.RWMutex for concurrent access
//
// 가변 작업 (안전하지 않음):
//   - GetOrSet: 맵 수정, 외부 동기화 필요
//   - SetDefault: 맵 수정, 외부 동기화 필요
//   - 동시 접근에 sync.Mutex 또는 sync.RWMutex 사용
//
// Immutable Operations (Safe for Concurrent Reads):
//   - Defaults: Creates new map, safe when inputs have concurrent readers
//   - Returned map should not be shared until fully initialized
//
// 불변 작업 (동시 읽기 안전):
//   - Defaults: 새 맵 생성, 입력에 동시 읽기가 있을 때 안전
//   - 반환된 맵은 완전히 초기화될 때까지 공유하지 않아야 함
//
// Concurrent Access Pattern:
//
//	var mu sync.RWMutex
//	cache := make(map[string]int)
//
//	func getCached(key string, compute func() int) int {
//	    mu.Lock()
//	    defer mu.Unlock()
//	    return maputil.GetOrSet(cache, key, compute())
//	}
//
// # See Also | 참고
//
// Related files in maputil package:
//   - basic.go: GetOr (read-only default), Set (unconditional set)
//   - merge.go: Merge (combines maps with last-wins)
//   - nested.go: GetNested, SetNested (nested defaults)
//
// maputil 패키지의 관련 파일:
//   - basic.go: GetOr (읽기 전용 기본값), Set (무조건 설정)
//   - merge.go: Merge (마지막 우선으로 맵 결합)
//   - nested.go: GetNested, SetNested (중첩 기본값)

// GetOrSet retrieves a value if it exists, otherwise sets and returns the default value.
// GetOrSet은 값이 존재하면 검색하고, 그렇지 않으면 기본값을 설정하고 반환합니다.
//
// This function provides a convenient way to ensure a key always has a value.
// If the key exists, it returns the existing value. If not, it sets the key
// to the default value and returns it. The original map is modified.
//
// 이 함수는 키가 항상 값을 가지도록 보장하는 편리한 방법을 제공합니다.
// 키가 존재하면 기존 값을 반환합니다. 존재하지 않으면 키를
// 기본값으로 설정하고 반환합니다. 원본 맵이 수정됩니다.
//
// Time Complexity
// 시간 복잡도: O(1)
// Space Complexity
// 공간 복잡도: O(1)
//
// Parameters
// 매개변수:
// - m: The input map (will be modified)
// 입력 맵 (수정됨)
// - key: The key to look up
// 조회할 키
// - defaultValue: The default value to set if key doesn't exist
// 키가 존재하지 않으면 설정할 기본값
//
// Returns
// 반환값:
// - V: The existing value or the default value
// 기존 값 또는 기본값
//
// Example
// 예제:
//
//	cache := map[string]int{"a": 1, "b": 2}
//	maputil.GetOrSet(cache, "a", 10) // Returns 1 (exists)
//	maputil.GetOrSet(cache, "c", 10) // Returns 10 (doesn't exist, sets c=10)
//	// cache is now: map[string]int{"a": 1, "b": 2, "c": 10}
//
// Use Case
// 사용 사례:
// - Cache initialization
// 캐시 초기화
// - Default value management
// 기본값 관리
// - Configuration with fallbacks
// 폴백이 있는 설정
// - Lazy initialization
// 지연 초기화
func GetOrSet[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if value, exists := m[key]; exists {
		return value
	}
	m[key] = defaultValue
	return defaultValue
}

// SetDefault sets a key to a default value only if the key doesn't already exist.
// SetDefault는 키가 아직 존재하지 않는 경우에만 키를 기본값으로 설정합니다.
//
// This function provides a way to initialize a key with a default value without
// overwriting existing values. It returns whether the key was set (true if the
// key didn't exist before, false if it already existed).
//
// 이 함수는 기존 값을 덮어쓰지 않고 기본값으로 키를 초기화하는 방법을 제공합니다.
// 키가 설정되었는지 여부를 반환합니다 (키가 이전에 존재하지 않았으면 true,
// 이미 존재했으면 false).
//
// Time Complexity
// 시간 복잡도: O(1)
// Space Complexity
// 공간 복잡도: O(1)
//
// Parameters
// 매개변수:
// - m: The input map (will be modified)
// 입력 맵 (수정됨)
// - key: The key to set
// 설정할 키
// - defaultValue: The default value to set if key doesn't exist
// 키가 존재하지 않으면 설정할 기본값
//
// Returns
// 반환값:
// - bool: true if key was set, false if key already existed
// 키가 설정되었으면 true, 이미 존재했으면 false
//
// Example
// 예제:
//
//	config := map[string]string{"host": "localhost"}
//	maputil.SetDefault(config, "port", "8080")    // Returns true, sets port=8080
//	maputil.SetDefault(config, "host", "0.0.0.0") // Returns false, host remains localhost
//	// config: map[string]string{"host": "localhost", "port": "8080"}
//
// Use Case
// 사용 사례:
// - Safe initialization without overwriting
// 덮어쓰지 않는 안전한 초기화
// - Default configuration setup
// 기본 설정 설정
// - Conditional key creation
// 조건부 키 생성
func SetDefault[K comparable, V any](m map[K]V, key K, defaultValue V) bool {
	if _, exists := m[key]; exists {
		return false
	}
	m[key] = defaultValue
	return true
}

// Defaults returns a new map with default values for missing keys.
// Defaults는 누락된 키에 대해 기본값이 있는 새 맵을 반환합니다.
//
// This function creates a new map that contains all entries from the original map
// plus any keys from the defaults map that weren't in the original. Existing keys
// in the original map are not overwritten.
//
// 이 함수는 원본 맵의 모든 항목과 원본에 없는 defaults 맵의 키를
// 포함하는 새 맵을 생성합니다. 원본 맵의 기존 키는 덮어쓰지 않습니다.
//
// Time Complexity
// 시간 복잡도: O(n + d) where n is original map size, d is defaults size
// Space Complexity
// 공간 복잡도: O(n + d)
//
// Parameters
// 매개변수:
// - m: The original map
// 원본 맵
// - defaults: Map containing default key-value pairs
// 기본 키-값 쌍을 포함하는 맵
//
// Returns
// 반환값:
// - map[K]V: New map with defaults applied
// 기본값이 적용된 새 맵
//
// Example
// 예제:
//
//	config := map[string]string{"host": "localhost"}
//	defaults := map[string]string{"host": "0.0.0.0", "port": "8080", "timeout": "30s"}
//	result := maputil.Defaults(config, defaults)
//	// result: map[string]string{"host": "localhost", "port": "8080", "timeout": "30s"}
//	// Note: host keeps original value, port and timeout are added from defaults
//
// Use Case
// 사용 사례:
// - Configuration management with defaults
// 기본값이 있는 설정 관리
// - Template rendering with default values
// 기본값이 있는 템플릿 렌더링
// - API response with fallback values
// 폴백 값이 있는 API 응답
// - User preferences with system defaults
// 시스템 기본값이 있는 사용자 기본 설정
func Defaults[K comparable, V any](m, defaults map[K]V) map[K]V {
	result := make(map[K]V, len(m)+len(defaults))

	// First, copy all defaults
	for k, v := range defaults {
		result[k] = v
	}

	// Then, overwrite with original values (original takes precedence)
	for k, v := range m {
		result[k] = v
	}

	return result
}
