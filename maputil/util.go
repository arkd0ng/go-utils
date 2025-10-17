package maputil

// =============================================================================
// File: util.go
// Purpose: General-Purpose Map Utility Operations
// 파일: util.go
// 목적: 범용 맵 유틸리티 연산
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The util.go file provides a collection of general-purpose utility functions
// for maps that don't fit into specific categories like filtering, transforming,
// or merging. These utilities handle common patterns like batch operations,
// iteration with side effects, validation, and functional transformations that
// are frequently needed but don't warrant their own dedicated files.
//
// util.go 파일은 필터링, 변환 또는 병합과 같은 특정 범주에 속하지 않는
// 범용 맵 유틸리티 함수 모음을 제공합니다. 이러한 유틸리티는 배치 연산,
// 부수 효과가 있는 반복, 검증 및 자주 필요하지만 별도의 전용 파일이
// 필요하지 않은 기능적 변환과 같은 일반적인 패턴을 처리합니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Convenience First**: Provide commonly-needed operations in simple, easy-to-use functions
//    **편의성 우선**: 간단하고 사용하기 쉬운 함수로 자주 필요한 연산 제공
//
// 2. **Functional Patterns**: Support functional programming patterns like Apply and Tap
//    **함수형 패턴**: Apply와 Tap 같은 함수형 프로그래밍 패턴 지원
//
// 3. **Batch Operations**: Enable efficient batch processing with GetMany and SetMany
//    **배치 연산**: GetMany와 SetMany로 효율적인 배치 처리 가능
//
// 4. **Method Chaining**: Facilitate fluent interfaces with operations like Tap
//    **메서드 체이닝**: Tap 같은 연산으로 유창한 인터페이스 촉진
//
// 5. **Side-Effect Support**: Provide safe ways to perform side effects with ForEach and Tap
//    **부수 효과 지원**: ForEach와 Tap으로 안전한 부수 효과 수행 방법 제공
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. ITERATION OPERATIONS (반복 연산)
//    - ForEach: Execute side-effect function for each key-value pair
//      ForEach: 각 키-값 쌍에 대해 부수 효과 함수 실행
//    - Apply: Transform all values using a mapping function
//      Apply: 매핑 함수를 사용하여 모든 값 변환
//
// 2. BATCH OPERATIONS (배치 연산)
//    - GetMany: Retrieve multiple values by keys at once
//      GetMany: 한 번에 여러 키로 값 조회
//    - SetMany: Set multiple key-value pairs at once
//      SetMany: 한 번에 여러 키-값 쌍 설정
//
// 3. VALIDATION OPERATIONS (검증 연산)
//    - ContainsAllKeys: Check if all specified keys exist
//      ContainsAllKeys: 지정된 모든 키가 존재하는지 확인
//
// 4. FUNCTIONAL UTILITIES (함수형 유틸리티)
//    - Tap: Execute side-effect function and return original map (for chaining)
//      Tap: 부수 효과 함수 실행 후 원본 맵 반환 (체이닝용)
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// ForEach[K comparable, V any](m map[K]V, fn func(K, V))
// - Purpose: Iterate over map entries and execute side-effect function
// - 목적: 맵 항목을 반복하고 부수 효과 함수 실행
// - Time Complexity: O(n) where n is number of entries
// - 시간 복잡도: O(n), n은 항목 수
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Logging, printing, accumulating statistics, validation
// - 사용 사례: 로깅, 출력, 통계 누적, 검증
//
// GetMany[K comparable, V any](m map[K]V, keys ...K) []V
// - Purpose: Retrieve values for multiple keys in one call
// - 목적: 한 번의 호출로 여러 키의 값 조회
// - Time Complexity: O(k) where k is number of keys
// - 시간 복잡도: O(k), k는 키 개수
// - Space Complexity: O(k) for result slice
// - 공간 복잡도: O(k), 결과 슬라이스용
// - Behavior: Returns zero values for missing keys
// - 동작: 누락된 키에 대해 제로값 반환
// - Use Cases: Batch retrieval, multiple parameter access, data gathering
// - 사용 사례: 배치 조회, 다중 매개변수 접근, 데이터 수집
//
// SetMany[K comparable, V any](m map[K]V, entries ...Entry[K, V]) map[K]V
// - Purpose: Create new map with multiple entries added
// - 목적: 여러 항목이 추가된 새 맵 생성
// - Time Complexity: O(n + e) where n is original size, e is number of entries
// - 시간 복잡도: O(n + e), n은 원본 크기, e는 항목 수
// - Space Complexity: O(n + e)
// - 공간 복잡도: O(n + e)
// - Immutability: Returns new map, original unchanged
// - 불변성: 새 맵 반환, 원본 변경 없음
// - Use Cases: Bulk initialization, configuration setup, batch updates
// - 사용 사례: 대량 초기화, 설정 구성, 배치 업데이트
//
// Tap[K comparable, V any](m map[K]V, fn func(map[K]V)) map[K]V
// - Purpose: Execute side-effect function and return original map
// - 목적: 부수 효과 함수 실행 후 원본 맵 반환
// - Time Complexity: O(f) where f is complexity of provided function
// - 시간 복잡도: O(f), f는 제공된 함수의 복잡도
// - Space Complexity: O(1) excluding function's space
// - 공간 복잡도: O(1), 함수의 공간 제외
// - Pattern: Enables method chaining while performing side effects
// - 패턴: 부수 효과 수행 중 메서드 체이닝 가능
// - Use Cases: Debugging, logging, validation in chains, collecting metrics
// - 사용 사례: 디버깅, 로깅, 체인의 검증, 메트릭 수집
//
// ContainsAllKeys[K comparable, V any](m map[K]V, keys []K) bool
// - Purpose: Validate that all specified keys exist in map
// - 목적: 지정된 모든 키가 맵에 존재하는지 검증
// - Time Complexity: O(k) where k is number of keys
// - 시간 복잡도: O(k), k는 키 개수
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Edge Case: Empty keys slice returns true (vacuous truth)
// - 엣지 케이스: 빈 키 슬라이스는 true 반환 (공허한 진리)
// - Use Cases: Configuration validation, API response checking, prerequisite validation
// - 사용 사례: 설정 검증, API 응답 확인, 전제 조건 검증
//
// Apply[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V
// - Purpose: Transform all values using key-value function
// - 목적: 키-값 함수를 사용하여 모든 값 변환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n) for new map
// - 공간 복잡도: O(n), 새 맵용
// - Pattern: Functional map transformation
// - 패턴: 함수형 맵 변환
// - Use Cases: Price calculations, unit conversions, data normalization, applying formulas
// - 사용 사례: 가격 계산, 단위 변환, 데이터 정규화, 수식 적용
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// ForEach vs MapValues (transform.go)
// - ForEach: Side effects only, returns nothing
//   ForEach: 부수 효과만, 아무것도 반환하지 않음
// - MapValues: Pure transformation, returns new map
//   MapValues: 순수 변환, 새 맵 반환
// - Use ForEach for: Logging, printing, accumulating external state
//   ForEach 사용: 로깅, 출력, 외부 상태 누적
// - Use MapValues for: Transforming values into new map
//   MapValues 사용: 새 맵으로 값 변환
//
// GetMany vs multiple Get calls
// - GetMany: Single function call, cleaner code
//   GetMany: 단일 함수 호출, 더 깔끔한 코드
// - Multiple Gets: More verbose, harder to read
//   다중 Get: 더 장황함, 읽기 어려움
// - Performance: Similar O(k), but GetMany has better ergonomics
//   성능: 유사한 O(k), 하지만 GetMany가 더 나은 인체공학
// - Use GetMany when: Retrieving multiple related values
//   GetMany 사용 시기: 여러 관련 값 조회
//
// SetMany vs Merge (merge.go)
// - SetMany: Takes Entry structs, more explicit
//   SetMany: Entry 구조체 사용, 더 명시적
// - Merge: Takes map, more natural for map-to-map merging
//   Merge: 맵 사용, 맵 간 병합에 더 자연스러움
// - Use SetMany for: Literal entry initialization
//   SetMany 사용: 리터럴 항목 초기화
// - Use Merge for: Combining existing maps
//   Merge 사용: 기존 맵 결합
//
// Tap vs ForEach
// - Tap: Returns map for chaining, accepts map as parameter
//   Tap: 체이닝을 위해 맵 반환, 맵을 매개변수로 받음
// - ForEach: Returns nothing, accepts key-value pairs
//   ForEach: 아무것도 반환하지 않음, 키-값 쌍 받음
// - Use Tap for: Method chaining, debugging pipelines
//   Tap 사용: 메서드 체이닝, 파이프라인 디버깅
// - Use ForEach for: Simple iteration with side effects
//   ForEach 사용: 부수 효과가 있는 단순 반복
//
// Apply vs MapValues (transform.go)
// - Apply: Transformation can use both key and value
//   Apply: 변환이 키와 값 모두 사용 가능
// - MapValues: Transformation uses only value
//   MapValues: 변환이 값만 사용
// - Use Apply when: Key context is needed for transformation
//   Apply 사용 시기: 변환에 키 컨텍스트가 필요할 때
// - Use MapValues when: Only value matters
//   MapValues 사용 시기: 값만 중요할 때
//
// ContainsAllKeys vs HasKey (basic.go)
// - ContainsAllKeys: Checks multiple keys at once
//   ContainsAllKeys: 여러 키를 한 번에 확인
// - HasKey: Checks single key
//   HasKey: 단일 키 확인
// - ContainsAllKeys: O(k) for k keys
//   ContainsAllKeys: k개 키에 대해 O(k)
// - HasKey: O(1) per call
//   HasKey: 호출당 O(1)
// - Use ContainsAllKeys for: Validating multiple prerequisites
//   ContainsAllKeys 사용: 여러 전제 조건 검증
// - Use HasKey for: Single key existence check
//   HasKey 사용: 단일 키 존재 확인
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - ForEach: O(n) - iterates over all entries
//   ForEach: O(n) - 모든 항목 반복
// - GetMany: O(k) - k lookups, each O(1)
//   GetMany: O(k) - k개 조회, 각각 O(1)
// - SetMany: O(n + e) - copy original map + add entries
//   SetMany: O(n + e) - 원본 맵 복사 + 항목 추가
// - Tap: O(f) - depends on provided function
//   Tap: O(f) - 제공된 함수에 의존
// - ContainsAllKeys: O(k) - k existence checks
//   ContainsAllKeys: O(k) - k개 존재 확인
// - Apply: O(n) - transforms all entries
//   Apply: O(n) - 모든 항목 변환
//
// Space Complexities:
// 공간 복잡도:
// - ForEach: O(1) - no additional space
//   ForEach: O(1) - 추가 공간 없음
// - GetMany: O(k) - result slice
//   GetMany: O(k) - 결과 슬라이스
// - SetMany: O(n + e) - new map
//   SetMany: O(n + e) - 새 맵
// - Tap: O(1) - returns original
//   Tap: O(1) - 원본 반환
// - ContainsAllKeys: O(1) - no additional space
//   ContainsAllKeys: O(1) - 추가 공간 없음
// - Apply: O(n) - new map
//   Apply: O(n) - 새 맵
//
// Optimization Tips:
// 최적화 팁:
// 1. Use GetMany instead of multiple individual Gets for cleaner code
//    더 깔끔한 코드를 위해 여러 개별 Get 대신 GetMany 사용
// 2. Pre-allocate maps when using SetMany with known size
//    알려진 크기로 SetMany 사용 시 맵 미리 할당
// 3. Avoid expensive operations in ForEach function parameter
//    ForEach 함수 매개변수에서 비용이 많이 드는 연산 피하기
// 4. Use Apply only when transformation is needed; otherwise use ForEach
//    변환이 필요할 때만 Apply 사용; 그렇지 않으면 ForEach 사용
// 5. Cache ContainsAllKeys results if checking same keys repeatedly
//    같은 키를 반복적으로 확인할 경우 ContainsAllKeys 결과 캐시
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Nil Maps:
// nil 맵:
// - ForEach: Safe, does nothing for nil map
//   ForEach: 안전함, nil 맵에 대해 아무것도 하지 않음
// - GetMany: Returns slice of zero values for all keys
//   GetMany: 모든 키에 대해 제로값 슬라이스 반환
// - SetMany: Creates new map with entries (handles nil gracefully)
//   SetMany: 항목이 있는 새 맵 생성 (nil을 우아하게 처리)
// - Tap: Safe if function handles nil, returns nil
//   Tap: 함수가 nil을 처리하면 안전함, nil 반환
// - ContainsAllKeys: Returns false if map is nil (unless keys empty)
//   ContainsAllKeys: 맵이 nil이면 false 반환 (키가 비어있지 않는 한)
// - Apply: Creates new empty map (handles nil gracefully)
//   Apply: 새 빈 맵 생성 (nil을 우아하게 처리)
//
// Empty Maps:
// 빈 맵:
// - ForEach: Does nothing, no iterations
//   ForEach: 아무것도 하지 않음, 반복 없음
// - GetMany: Returns zero values for all requested keys
//   GetMany: 요청된 모든 키에 대해 제로값 반환
// - SetMany: Returns map with only new entries
//   SetMany: 새 항목만 있는 맵 반환
// - Tap: Executes function with empty map, returns empty map
//   Tap: 빈 맵으로 함수 실행, 빈 맵 반환
// - ContainsAllKeys: Returns true if keys slice is empty (vacuous truth)
//   ContainsAllKeys: 키 슬라이스가 비어있으면 true 반환 (공허한 진리)
// - Apply: Returns empty map
//   Apply: 빈 맵 반환
//
// Missing Keys:
// 누락된 키:
// - GetMany: Returns zero value for missing keys (no error)
//   GetMany: 누락된 키에 대해 제로값 반환 (오류 없음)
// - ContainsAllKeys: Returns false if any key is missing
//   ContainsAllKeys: 키가 하나라도 누락되면 false 반환
//
// Duplicate Keys in SetMany:
// SetMany의 중복 키:
// - Last entry wins (overwrites previous)
//   마지막 항목이 승리 (이전 항목 덮어쓰기)
//
// Empty Keys Slice:
// 빈 키 슬라이스:
// - GetMany: Returns empty slice
//   GetMany: 빈 슬라이스 반환
// - ContainsAllKeys: Returns true (vacuous truth: all zero keys exist)
//   ContainsAllKeys: true 반환 (공허한 진리: 모든 0개 키가 존재)
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Logging Map Contents
//    맵 내용 로깅:
//
//    scores := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}
//    maputil.ForEach(scores, func(name string, score int) {
//        log.Printf("Student %s scored %d", name, score)
//    })
//    // Logs each student's score
//    // 각 학생의 점수 로깅
//
// 2. Batch Retrieval of Configuration Values
//    설정 값의 배치 조회:
//
//    config := map[string]string{
//        "host": "localhost",
//        "port": "8080",
//        "timeout": "30s",
//    }
//    values := maputil.GetMany(config, "host", "port", "timeout")
//    host, port, timeout := values[0], values[1], values[2]
//    // Retrieve multiple config values at once
//    // 여러 설정 값을 한 번에 조회
//
// 3. Initializing Map with Multiple Entries
//    여러 항목으로 맵 초기화:
//
//    base := map[string]int{"a": 1}
//    result := maputil.SetMany(base,
//        maputil.Entry[string, int]{Key: "b", Value: 2},
//        maputil.Entry[string, int]{Key: "c", Value: 3},
//        maputil.Entry[string, int]{Key: "d", Value: 4},
//    )
//    // result: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//    // Bulk initialization with entries
//    // 항목으로 대량 초기화
//
// 4. Debugging in Method Chain
//    메서드 체인에서 디버깅:
//
//    result := maputil.Filter(data, predicate).
//        Pipe(maputil.Tap(func(m map[string]int) {
//            fmt.Printf("After filter: %v\n", m)
//        })).
//        Pipe(maputil.MapValues(transform))
//    // Debug intermediate results without breaking chain
//    // 체인을 끊지 않고 중간 결과 디버깅
//
// 5. Validating Required Fields
//    필수 필드 검증:
//
//    request := map[string]interface{}{
//        "username": "alice",
//        "email": "alice@example.com",
//        "age": 30,
//    }
//    required := []string{"username", "email", "password"}
//    if !maputil.ContainsAllKeys(request, required) {
//        return errors.New("missing required fields")
//    }
//    // Validate all required fields present
//    // 모든 필수 필드가 있는지 검증
//
// 6. Applying Discount to All Prices
//    모든 가격에 할인 적용:
//
//    prices := map[string]float64{
//        "apple": 1.50,
//        "banana": 0.75,
//        "orange": 2.00,
//    }
//    discounted := maputil.Apply(prices, func(item string, price float64) float64 {
//        return price * 0.9 // 10% discount
//    })
//    // Apply discount to all items
//    // 모든 항목에 할인 적용
//
// 7. Accumulating Statistics with ForEach
//    ForEach로 통계 누적:
//
//    sales := map[string]int{"Q1": 1000, "Q2": 1500, "Q3": 1200, "Q4": 1800}
//    var total int
//    var count int
//    maputil.ForEach(sales, func(quarter string, amount int) {
//        total += amount
//        count++
//    })
//    average := total / count
//    // Calculate statistics from map
//    // 맵에서 통계 계산
//
// 8. Converting Units with Apply
//    Apply로 단위 변환:
//
//    distancesKm := map[string]float64{
//        "route1": 10.0,
//        "route2": 25.5,
//        "route3": 8.2,
//    }
//    distancesMiles := maputil.Apply(distancesKm, func(route string, km float64) float64 {
//        return km * 0.621371 // km to miles
//    })
//    // Convert all distances to different unit
//    // 모든 거리를 다른 단위로 변환
//
// 9. Checking API Response Completeness
//    API 응답 완전성 확인:
//
//    response := map[string]interface{}{
//        "id": 123,
//        "name": "Product",
//        "price": 29.99,
//        "stock": 50,
//    }
//    expectedFields := []string{"id", "name", "price", "stock", "category"}
//    if !maputil.ContainsAllKeys(response, expectedFields) {
//        log.Println("Warning: API response missing expected fields")
//    }
//    // Validate API response structure
//    // API 응답 구조 검증
//
// 10. Chained Transformations with Tap for Monitoring
//     모니터링을 위한 Tap을 사용한 체인 변환:
//
//     processed := maputil.Filter(rawData, isValid).
//         Pipe(maputil.Tap(func(m map[string]int) {
//             log.Printf("After validation: %d items", len(m))
//         })).
//         Pipe(maputil.Apply(m, normalize)).
//         Pipe(maputil.Tap(func(m map[string]int) {
//             log.Printf("After normalization: %d items", len(m))
//         }))
//     // Monitor pipeline stages without interrupting flow
//     // 흐름을 중단하지 않고 파이프라인 단계 모니터링
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// Most functions in this file are NOT thread-safe when the map is being
// modified concurrently. Use synchronization mechanisms when needed:
//
// 이 파일의 대부분의 함수는 맵이 동시에 수정될 때 스레드 안전하지 않습니다.
// 필요할 때 동기화 메커니즘을 사용하세요:
//
// Thread-Safe Operations:
// 스레드 안전 연산:
// - Read-only operations (GetMany, ContainsAllKeys) are safe if map not modified
//   읽기 전용 연산 (GetMany, ContainsAllKeys)은 맵이 수정되지 않으면 안전
// - Immutable operations (SetMany, Apply) return new maps, safe for concurrent reads
//   불변 연산 (SetMany, Apply)은 새 맵 반환, 동시 읽기에 안전
//
// Not Thread-Safe:
// 스레드 안전하지 않음:
// - ForEach: Not safe if map is modified during iteration
//   ForEach: 반복 중 맵이 수정되면 안전하지 않음
// - Tap: Not safe if function modifies map or map modified externally
//   Tap: 함수가 맵을 수정하거나 외부에서 맵이 수정되면 안전하지 않음
//
// Safe Concurrent Access Pattern:
// 안전한 동시 접근 패턴:
//
//     var mu sync.RWMutex
//     var sharedMap = make(map[string]int)
//
//     // Reading with GetMany
//     // GetMany로 읽기
//     mu.RLock()
//     values := maputil.GetMany(sharedMap, "key1", "key2")
//     mu.RUnlock()
//
//     // Updating with SetMany
//     // SetMany로 업데이트
//     mu.Lock()
//     sharedMap = maputil.SetMany(sharedMap,
//         maputil.Entry[string, int]{Key: "key3", Value: 3},
//     )
//     mu.Unlock()
//
// RELATED FILES
// 관련 파일
// -------------
// - basic.go: Basic map operations (Get, Set, Delete, Has, etc.)
//   basic.go: 기본 맵 연산 (Get, Set, Delete, Has 등)
// - transform.go: Value transformation operations (MapValues, MapKeys, etc.)
//   transform.go: 값 변환 연산 (MapValues, MapKeys 등)
// - filter.go: Filtering and selection operations
//   filter.go: 필터링 및 선택 연산
// - aggregate.go: Aggregation and reduction operations
//   aggregate.go: 집계 및 축소 연산
// - merge.go: Merging and combining operations
//   merge.go: 병합 및 결합 연산
// - convert.go: Conversion to other types (ToJSON, FromSlice, etc.)
//   convert.go: 다른 타입으로 변환 (ToJSON, FromSlice 등)
//
// =============================================================================

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
// Time Complexity
// 시간 복잡도: O(n)
// Space Complexity
// 공간 복잡도: O(1)
//
// Parameters
// 매개변수:
// - m: The input map
// 입력 맵
// - fn: Function to execute for each entry
// 각 항목에 대해 실행할 함수
//
// Example
// 예제:
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
// Use Case
// 사용 사례:
// - Logging all map entries
// 모든 맵 항목 로깅
// - Debugging map contents
// 맵 내용 디버깅
// - Performing side effects (e.g., sending notifications)
// 부수 효과 수행 (예: 알림 전송)
// - Collecting statistics without modifying the map
// 맵을 수정하지 않고 통계 수집
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
// Time Complexity
// 시간 복잡도: O(k) where k is the number of keys
// k는 키의 개수
// Space Complexity
// 공간 복잡도: O(k)
//
// Parameters
// 매개변수:
// - m: The input map
// 입력 맵
// - keys: Variable number of keys to retrieve
// 검색할 키의 가변 개수
//
// Returns
// 반환값:
// - []V: Slice of values corresponding to the keys
// 키에 해당하는 값의 슬라이스
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	values := maputil.GetMany(m, "a", "c", "d")
//	// values: [1, 3, 0] (d doesn't exist, returns zero value)
//
// Use Case
// 사용 사례:
// - Batch retrieval of multiple values
// 여러 값의 일괄 검색
// - Configuration lookups
// 설정 조회
// - Data extraction for processing
// 처리를 위한 데이터 추출
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
// Time Complexity
// 시간 복잡도: O(n + e) where n is map size, e is entries
// n은 맵 크기, e는 항목 개수
// Space Complexity
// 공간 복잡도: O(n + e)
//
// Parameters
// 매개변수:
// - m: The input map
// 입력 맵
// - entries: Variable number of Entry structs to set
// 설정할 Entry 구조체의 가변 개수
//
// Returns
// 반환값:
// - map[K]V: New map with updated entries
// 업데이트된 항목이 있는 새 맵
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	result := maputil.SetMany(m,
//	    maputil.Entry[string, int]{Key: "c", Value: 3},
//	    maputil.Entry[string, int]{Key: "d", Value: 4},
//	)
//	// result: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//
// Use Case
// 사용 사례:
// - Batch updates to configuration
// 설정에 대한 일괄 업데이트
// - Initializing map with multiple values
// 여러 값으로 맵 초기화
// - Merging multiple entries at once
// 여러 항목을 한 번에 병합
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
// Time Complexity
// 시간 복잡도: O(n) - depends on fn
// fn에 따라 다름
// Space Complexity
// 공간 복잡도: O(1)
//
// Parameters
// 매개변수:
// - m: The input map
// 입력 맵
// - fn: Side-effect function to execute
// 실행할 부수 효과 함수
//
// Returns
// 반환값:
// - map[K]V: The original map (unchanged)
// 원본 맵 (변경되지 않음)
//
// Example
// 예제:
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
// Use Case
// 사용 사례:
// - Debugging in method chains
// 메서드 체인에서 디버깅
// - Logging intermediate results
// 중간 결과 로깅
// - Collecting statistics without breaking chain
// 체인을 끊지 않고 통계 수집
// - Performing validation or assertions
// 유효성 검사 또는 단언 수행
func Tap[K comparable, V any](m map[K]V, fn func(map[K]V)) map[K]V {
	fn(m)
	return m
}

// ContainsAllKeys checks if the map contains all the specified keys.
// ContainsAllKeys는 맵에 지정된 모든 키가 포함되어 있는지 확인합니다.
//
// This function returns true only if all provided keys exist in the map.
// If the keys slice is empty, it returns true (vacuous truth).
//
// 이 함수는 제공된 모든 키가 맵에 존재하는 경우에만 true를 반환합니다.
// 키 슬라이스가 비어 있으면 true를 반환합니다 (공허한 진리).
//
// Time Complexity
// 시간 복잡도: O(k) where k is the number of keys
// k는 키의 개수
// Space Complexity
// 공간 복잡도: O(1)
//
// Parameters
// 매개변수:
// - m: The input map
// 입력 맵
// - keys: Slice of keys to check
// 확인할 키의 슬라이스
//
// Returns
// 반환값:
// - bool: true if all keys exist, false otherwise
// 모든 키가 존재하면 true, 그렇지 않으면 false
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	maputil.ContainsAllKeys(m, []string{"a", "b"})      // true
//	maputil.ContainsAllKeys(m, []string{"a", "d"})      // false
//	maputil.ContainsAllKeys(m, []string{})              // true (empty keys)
//
// Use Case
// 사용 사례:
// - Validating required configuration keys
// 필수 설정 키 검증
// - Checking API response completeness
// API 응답 완전성 확인
// - Ensuring all dependencies are present
// 모든 종속성이 존재하는지 확인
// - Form validation
// 폼 검증
func ContainsAllKeys[K comparable, V any](m map[K]V, keys []K) bool {
	for _, key := range keys {
		if _, exists := m[key]; !exists {
			return false
		}
	}
	return true
}

// Apply modifies the map by applying a function to each key-value pair.
// Apply는 각 키-값 쌍에 함수를 적용하여 맵을 수정합니다.
//
// This function creates a new map where each value is the result of applying
// the transformation function to the original key-value pair. The keys remain
// the same, but the values are transformed.
//
// 이 함수는 각 값이 원본 키-값 쌍에 변환 함수를 적용한 결과인 새 맵을 생성합니다.
// 키는 동일하게 유지되지만 값이 변환됩니다.
//
// Time Complexity
// 시간 복잡도: O(n)
// Space Complexity
// 공간 복잡도: O(n)
//
// Parameters
// 매개변수:
// - m: The input map
// 입력 맵
// - fn: Transformation function
// 변환 함수
//
// Returns
// 반환값:
// - map[K]V: New map with transformed values
// 변환된 값이 있는 새 맵
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.Apply(m, func(k string, v int) int {
//	    return v * 2
//	})
//	// result: map[string]int{"a": 2, "b": 4, "c": 6}
//
// Use Case
// 사용 사례:
// - Transforming all values in a map
// 맵의 모든 값 변환
// - Applying discounts to prices
// 가격에 할인 적용
// - Normalizing data values
// 데이터 값 정규화
// - Converting units (e.g., km to miles)
// 단위 변환 (예: km를 마일로)
func Apply[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = fn(k, v)
	}
	return result
}
