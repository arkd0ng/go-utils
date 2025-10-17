package maputil

import "sort"

// Package maputil/values.go provides comprehensive value-focused operations for maps.
// This file contains functions for extracting, transforming, filtering, searching,
// and analyzing map values independently of their keys.
//
// maputil/values.go 패키지는 맵의 값 중심 작업을 위한 포괄적인 기능을 제공합니다.
// 이 파일은 키와 무관하게 맵 값을 추출, 변환, 필터링, 검색, 분석하는 함수들을 포함합니다.
//
// # Overview | 개요
//
// The values.go file provides 12 value-centric operations organized into 5 categories:
//
// values.go 파일은 5개 카테고리로 구성된 12개 값 중심 작업을 제공합니다:
//
// 1. VALUE EXTRACTION | 값 추출
//   - ValuesSlice: Extract all values as a slice (O(n))
//   - ValuesSorted: Extract values in sorted order (O(n log n))
//   - ValuesBy: Extract values matching predicate (O(n))
//   - UniqueValues: Extract deduplicated values (O(n))
//
// 2. VALUE SEARCHING | 값 검색
//   - FindValue: Find first value satisfying predicate (O(n), early termination)
//
// 3. VALUE TRANSFORMATION | 값 변환
//   - ReplaceValue: Replace all occurrences of a value (O(n))
//   - UpdateValues: Transform all values with key-value function (O(n))
//
// 4. MATHEMATICAL OPERATIONS | 수학 연산
//   - MinValue: Find minimum value (O(n))
//   - MaxValue: Find maximum value (O(n))
//   - SumValues: Calculate sum of all values (O(n))
//
// 5. SPECIALIZED | 특수 기능
//   - Values comparison and validation operations
//
// # Design Principles | 설계 원칙
//
// 1. VALUE INDEPENDENCE | 값 독립성
//   - Operations focus on values, minimizing key consideration
//   - Counterpart to keys.go functions
//   - Key context provided only when necessary (ValuesBy, UpdateValues)
//
//   작업은 값에 초점을 맞추며 키 고려를 최소화합니다
//   keys.go 함수의 대응 기능
//   필요시에만 키 컨텍스트 제공 (ValuesBy, UpdateValues)
//
// 2. IMMUTABILITY | 불변성
//   - Value modification functions return new maps
//   - Original maps remain unchanged
//   - Safe for concurrent read access
//
//   값 수정 함수는 새 맵을 반환합니다
//   원본 맵은 변경되지 않습니다
//   동시 읽기 접근에 안전합니다
//
// 3. TYPE CONSTRAINTS | 타입 제약
//   - ValuesSorted requires Ordered constraint for sorting
//   - UniqueValues, ReplaceValue require comparable values for deduplication
//   - Mathematical operations require Number constraint
//   - Most functions accept any value type
//
//   ValuesSorted는 정렬을 위해 Ordered 제약이 필요합니다
//   UniqueValues, ReplaceValue는 중복 제거를 위해 comparable 값 필요
//   수학 연산은 Number 제약이 필요합니다
//   대부분의 함수는 모든 값 타입을 허용합니다
//
// 4. PERFORMANCE OPTIMIZATION | 성능 최적화
//   - Early termination in FindValue when match found
//   - UniqueValues uses map for O(1) deduplication
//   - Pre-allocated slices where sizes are known
//   - Single iteration for most operations
//
//   FindValue는 매치 발견 시 조기 종료
//   UniqueValues는 O(1) 중복 제거를 위해 맵 사용
//   크기가 알려진 경우 슬라이스를 미리 할당
//   대부분의 작업이 단일 순회
//
// # Function Categories | 함수 카테고리
//
// VALUE EXTRACTION OPERATIONS | 값 추출 작업
//
// ValuesSlice(m) extracts all values from a map as a slice. It's an alias for
// the core Values function, providing semantic clarity when you explicitly want
// a slice of values.
//
// ValuesSlice(m)는 맵의 모든 값을 슬라이스로 추출합니다. 명시적으로 값 슬라이스를
// 원할 때 의미적 명확성을 제공하는 핵심 Values 함수의 별칭입니다.
//
// Time Complexity: O(n) - iterates through all map entries
// Space Complexity: O(n) - creates new slice with all values
// Order: Non-deterministic (map iteration order is random)
// Nil Handling: Returns empty slice for nil maps
//
// 시간 복잡도: O(n) - 모든 맵 항목 순회
// 공간 복잡도: O(n) - 모든 값을 포함하는 새 슬라이스 생성
// 순서: 비결정적 (맵 순회 순서는 무작위)
// Nil 처리: nil 맵에 대해 빈 슬라이스 반환
//
// Use Case: Extract values when key information is not needed
// 사용 사례: 키 정보가 필요없을 때 값 추출
//
// ValuesSorted(m) extracts all values and returns them in sorted ascending order.
// This requires the Ordered constraint, meaning values must support comparison
// operators (<, >, ==).
//
// ValuesSorted(m)는 모든 값을 추출하고 오름차순 정렬하여 반환합니다.
// Ordered 제약이 필요하며, 값이 비교 연산자를 지원해야 합니다.
//
// Time Complexity: O(n log n) - extraction O(n) + sorting O(n log n)
// Space Complexity: O(n) - creates new sorted slice
// Order: Deterministic ascending order
// Nil Handling: Returns empty slice for nil maps
//
// 시간 복잡도: O(n log n) - 추출 O(n) + 정렬 O(n log n)
// 공간 복잡도: O(n) - 새 정렬된 슬라이스 생성
// 순서: 결정적 오름차순
// Nil 처리: nil 맵에 대해 빈 슬라이스 반환
//
// Use Case: When you need consistent, ordered value results
// 사용 사례: 일관되고 정렬된 값 결과가 필요할 때
//
// ValuesBy(m, fn) extracts only values that satisfy a predicate function.
// Unlike ValuesSlice which returns all values, this filters based on a condition
// that has access to both key and value.
//
// ValuesBy(m, fn)는 조건 함수를 만족하는 값만 추출합니다.
// 모든 값을 반환하는 ValuesSlice와 달리, 키와 값 모두에 접근하는 조건으로 필터링합니다.
//
// Time Complexity: O(n) - evaluates predicate for each entry
// Space Complexity: O(k) where k is the number of matching values
// Order: Non-deterministic (map iteration order)
// Nil Handling: Returns empty slice for nil maps
//
// 시간 복잡도: O(n) - 각 항목에 대해 조건 평가
// 공간 복잡도: O(k) 여기서 k는 매칭되는 값의 개수
// 순서: 비결정적 (맵 순회 순서)
// Nil 처리: nil 맵에 대해 빈 슬라이스 반환
//
// Use Case: Extract values based on key-value conditions
// 사용 사례: 키-값 조건에 기반한 값 추출
//
// UniqueValues(m) returns deduplicated values from a map. Maps can have multiple
// keys with the same value; this function ensures each unique value appears only
// once in the result.
//
// UniqueValues(m)는 맵에서 중복 제거된 값을 반환합니다. 맵은 같은 값을 가진 여러 키를
// 가질 수 있으며, 이 함수는 각 고유값이 결과에 한 번만 나타나도록 보장합니다.
//
// Time Complexity: O(n) - single iteration with map-based deduplication
// Space Complexity: O(u) where u is the number of unique values
// Deduplication: Uses temporary map for O(1) uniqueness checking
// Order: Non-deterministic (first occurrence of each unique value)
// Nil Handling: Returns empty slice for nil maps
//
// 시간 복잡도: O(n) - 맵 기반 중복 제거로 단일 순회
// 공간 복잡도: O(u) 여기서 u는 고유값의 개수
// 중복 제거: O(1) 고유성 검사를 위한 임시 맵 사용
// 순서: 비결정적 (각 고유값의 첫 발생)
// Nil 처리: nil 맵에 대해 빈 슬라이스 반환
//
// Use Case: Get distinct values when duplicates exist across keys
// 사용 사례: 키 간 중복이 존재할 때 고유값 얻기
//
// VALUE SEARCHING OPERATIONS | 값 검색 작업
//
// FindValue(m, fn) finds the first value that satisfies a predicate function.
// It provides early termination optimization, stopping as soon as a match is found.
//
// FindValue(m, fn)는 조건 함수를 만족하는 첫 번째 값을 찾습니다.
// 조기 종료 최적화를 제공하여 매치가 발견되는 즉시 중단합니다.
//
// Time Complexity: O(n) worst case, O(1) best case with early termination
// Space Complexity: O(1) - returns single value
// Return Value: (value, true) if found, (zero, false) if not found
// Order: Non-deterministic (depends on map iteration order)
// Early Termination: Stops at first match for efficiency
// Nil Handling: Returns (zero, false) for nil maps
//
// 시간 복잡도: 최악 O(n), 조기 종료로 최선 O(1)
// 공간 복잡도: O(1) - 단일 값 반환
// 반환값: 찾으면 (값, true), 못 찾으면 (제로값, false)
// 순서: 비결정적 (맵 순회 순서에 의존)
// 조기 종료: 효율성을 위해 첫 매치에서 중단
// Nil 처리: nil 맵에 대해 (제로값, false) 반환
//
// Use Case: Check if any value satisfies a condition and retrieve it
// 사용 사례: 조건을 만족하는 값이 있는지 확인하고 검색
//
// VALUE TRANSFORMATION OPERATIONS | 값 변환 작업
//
// ReplaceValue(m, old, new) creates a new map with all occurrences of a specific
// value replaced by a new value. Keys remain unchanged; only value substitution occurs.
//
// ReplaceValue(m, old, new)는 특정 값의 모든 발생을 새 값으로 대체한 새 맵을 생성합니다.
// 키는 변경되지 않고 값만 대체됩니다.
//
// Time Complexity: O(n) - iterates through all entries
// Space Complexity: O(n) - creates new map with same size
// Immutability: Returns new map, original unchanged
// Comparison: Requires comparable constraint for value equality check
// Replacement: All matching values replaced (not just first)
// Nil Handling: Returns empty map for nil maps
//
// 시간 복잡도: O(n) - 모든 항목 순회
// 공간 복잡도: O(n) - 같은 크기의 새 맵 생성
// 불변성: 새 맵 반환, 원본 변경 없음
// 비교: 값 동등성 검사를 위해 comparable 제약 필요
// 대체: 모든 매칭 값 대체 (첫 번째만이 아님)
// Nil 처리: nil 맵에 대해 빈 맵 반환
//
// Use Case: Normalize values, fix data inconsistencies
// 사용 사례: 값 정규화, 데이터 불일치 수정
//
// UpdateValues(m, fn) applies a transformation function to all values in a map,
// creating a new map with transformed values. The function receives both key and
// value, allowing key-dependent transformations.
//
// UpdateValues(m, fn)는 맵의 모든 값에 변환 함수를 적용하여 변환된 값을 가진
// 새 맵을 생성합니다. 함수는 키와 값 모두를 받아 키 의존적 변환을 허용합니다.
//
// Time Complexity: O(n) - applies transformation to each value
// Space Complexity: O(n) - creates new map
// Immutability: Returns new map, original unchanged
// Type Preservation: Input and output value types are the same
// Key Access: Function receives key context for conditional logic
// Nil Handling: Returns empty map for nil maps
//
// 시간 복잡도: O(n) - 각 값에 변환 적용
// 공간 복잡도: O(n) - 새 맵 생성
// 불변성: 새 맵 반환, 원본 변경 없음
// 타입 보존: 입력과 출력 값 타입이 동일
// 키 접근: 함수가 조건 로직을 위한 키 컨텍스트 수신
// Nil 처리: nil 맵에 대해 빈 맵 반환
//
// Comparison with MapValues: UpdateValues accepts (K, V) -> V, MapValues accepts V -> T
// MapValues와 비교: UpdateValues는 (K, V) -> V, MapValues는 V -> T 허용
//
// Use Case: Transform values based on both key and value context
// 사용 사례: 키와 값 컨텍스트 모두에 기반한 값 변환
//
// MATHEMATICAL OPERATIONS | 수학 연산
//
// MinValue(m) finds the minimum value in a map. Returns the smallest value
// according to the natural ordering defined by the Ordered constraint.
//
// MinValue(m)는 맵에서 최솟값을 찾습니다. Ordered 제약으로 정의된
// 자연 순서에 따른 가장 작은 값을 반환합니다.
//
// Time Complexity: O(n) - examines all values
// Space Complexity: O(1) - tracks only minimum
// Return Value: (min, true) if found, (zero, false) if empty
// Constraint: Requires Ordered constraint for comparison
// Empty Map: Returns (zero, false)
// Nil Handling: Returns (zero, false) for nil maps
//
// 시간 복잡도: O(n) - 모든 값 검사
// 공간 복잡도: O(1) - 최솟값만 추적
// 반환값: 찾으면 (최솟값, true), 비었으면 (제로값, false)
// 제약: 비교를 위해 Ordered 제약 필요
// 빈 맵: (제로값, false) 반환
// Nil 처리: nil 맵에 대해 (제로값, false) 반환
//
// Use Case: Find lowest value (temperature, price, score, etc.)
// 사용 사례: 최저값 찾기 (온도, 가격, 점수 등)
//
// MaxValue(m) finds the maximum value in a map. Returns the largest value
// according to the natural ordering defined by the Ordered constraint.
//
// MaxValue(m)는 맵에서 최댓값을 찾습니다. Ordered 제약으로 정의된
// 자연 순서에 따른 가장 큰 값을 반환합니다.
//
// Time Complexity: O(n) - examines all values
// Space Complexity: O(1) - tracks only maximum
// Return Value: (max, true) if found, (zero, false) if empty
// Constraint: Requires Ordered constraint for comparison
// Empty Map: Returns (zero, false)
// Nil Handling: Returns (zero, false) for nil maps
//
// 시간 복잡도: O(n) - 모든 값 검사
// 공간 복잡도: O(1) - 최댓값만 추적
// 반환값: 찾으면 (최댓값, true), 비었으면 (제로값, false)
// 제약: 비교를 위해 Ordered 제약 필요
// 빈 맵: (제로값, false) 반환
// Nil 처리: nil 맵에 대해 (제로값, false) 반환
//
// Use Case: Find highest value (temperature, price, score, etc.)
// 사용 사례: 최고값 찾기 (온도, 가격, 점수 등)
//
// SumValues(m) calculates the sum of all values in a map. Works with numeric types
// that implement the Number constraint (integers, floats).
//
// SumValues(m)는 맵의 모든 값의 합을 계산합니다. Number 제약을 구현하는
// 숫자 타입(정수, 실수)과 함께 작동합니다.
//
// Time Complexity: O(n) - sums all values
// Space Complexity: O(1) - accumulates in single variable
// Return Value: Sum of all values, 0 for empty map
// Constraint: Requires Number constraint for addition operator
// Empty Map: Returns zero value
// Overflow: No overflow protection; use appropriate type size
// Nil Handling: Returns zero for nil maps
//
// 시간 복잡도: O(n) - 모든 값 합산
// 공간 복잡도: O(1) - 단일 변수에 누적
// 반환값: 모든 값의 합, 빈 맵은 0
// 제약: 덧셈 연산자를 위해 Number 제약 필요
// 빈 맵: 제로값 반환
// 오버플로: 오버플로 보호 없음; 적절한 타입 크기 사용
// Nil 처리: nil 맵에 대해 제로값 반환
//
// Use Case: Calculate totals (prices, counts, scores)
// 사용 사례: 합계 계산 (가격, 카운트, 점수)
//
// # Comparisons with Related Functions | 관련 함수와 비교
//
// ValuesSlice vs. KeysSlice:
//   - ValuesSlice returns values, KeysSlice returns keys
//   - Both are O(n) with non-deterministic order
//   - Use ValuesSlice when keys are irrelevant
//   - Use KeysSlice when you need to iterate by keys
//
// ValuesSlice 대 KeysSlice:
//   - ValuesSlice는 값 반환, KeysSlice는 키 반환
//   - 둘 다 O(n)이며 비결정적 순서
//   - 키가 무관할 때 ValuesSlice 사용
//   - 키로 순회가 필요할 때 KeysSlice 사용
//
// ValuesSorted vs. KeysSorted:
//   - ValuesSorted sorts values, KeysSorted sorts keys
//   - Both are O(n log n) with deterministic order
//   - ValuesSorted requires values to be Ordered
//   - Use ValuesSorted for value-based sorted access
//
// ValuesSorted 대 KeysSorted:
//   - ValuesSorted는 값 정렬, KeysSorted는 키 정렬
//   - 둘 다 O(n log n)이며 결정적 순서
//   - ValuesSorted는 값이 Ordered여야 함
//   - 값 기반 정렬 접근에 ValuesSorted 사용
//
// ValuesBy vs. FilterValues:
//   - ValuesBy returns values matching predicate
//   - FilterValues returns map entries with matching values
//   - ValuesBy returns []V, FilterValues returns map[K]V
//   - Use ValuesBy when you only need values
//   - Use FilterValues when key association is needed
//
// ValuesBy 대 FilterValues:
//   - ValuesBy는 조건 매칭 값 반환
//   - FilterValues는 매칭 값을 가진 맵 항목 반환
//   - ValuesBy는 []V, FilterValues는 map[K]V 반환
//   - 값만 필요할 때 ValuesBy 사용
//   - 키 연관이 필요할 때 FilterValues 사용
//
// FindValue vs. FindKey:
//   - FindValue finds and returns a value
//   - FindKey finds and returns a key
//   - Both use early termination optimization
//   - Use FindValue when you want the value itself
//   - Use FindKey when you want the key (to access value later)
//
// FindValue 대 FindKey:
//   - FindValue는 값을 찾아 반환
//   - FindKey는 키를 찾아 반환
//   - 둘 다 조기 종료 최적화 사용
//   - 값 자체가 필요할 때 FindValue 사용
//   - 키가 필요할 때 (나중에 값 접근) FindKey 사용
//
// ReplaceValue vs. UpdateValues:
//   - ReplaceValue substitutes specific value: old -> new
//   - UpdateValues transforms all values: V -> f(K, V)
//   - ReplaceValue simpler for exact replacements
//   - UpdateValues more flexible for computed transformations
//
// ReplaceValue 대 UpdateValues:
//   - ReplaceValue는 특정 값 대체: 기존 -> 새 값
//   - UpdateValues는 모든 값 변환: V -> f(K, V)
//   - ReplaceValue가 정확한 대체에 더 간단
//   - UpdateValues가 계산된 변환에 더 유연
//
// UpdateValues vs. MapValues:
//   - UpdateValues: (K, V) -> V (key-aware, same value type)
//   - MapValues: V -> T (key-independent, can change value type)
//   - UpdateValues when key context needed
//   - MapValues when simple value transformation
//
// UpdateValues 대 MapValues:
//   - UpdateValues: (K, V) -> V (키 인식, 같은 값 타입)
//   - MapValues: V -> T (키 독립적, 값 타입 변경 가능)
//   - 키 컨텍스트 필요시 UpdateValues
//   - 단순 값 변환시 MapValues
//
// UniqueValues vs. Frequencies:
//   - UniqueValues returns []V with distinct values
//   - Frequencies returns map[V]int with counts
//   - UniqueValues when you only need distinct values
//   - Frequencies when you need occurrence counts
//
// UniqueValues 대 Frequencies:
//   - UniqueValues는 고유값을 가진 []V 반환
//   - Frequencies는 카운트를 가진 map[V]int 반환
//   - 고유값만 필요할 때 UniqueValues
//   - 발생 횟수가 필요할 때 Frequencies
//
// MinValue/MaxValue vs. MinBy/MaxBy:
//   - MinValue/MaxValue: Direct value comparison (Ordered constraint)
//   - MinBy/MaxBy: Custom score extraction (any comparable score)
//   - Use MinValue/MaxValue for simple value extremes
//   - Use MinBy/MaxBy for complex scoring logic
//
// MinValue/MaxValue 대 MinBy/MaxBy:
//   - MinValue/MaxValue: 직접 값 비교 (Ordered 제약)
//   - MinBy/MaxBy: 커스텀 점수 추출 (비교 가능한 모든 점수)
//   - 단순 값 극값에 MinValue/MaxValue 사용
//   - 복잡한 점수 로직에 MinBy/MaxBy 사용
//
// SumValues vs. Reduce:
//   - SumValues: Specialized for addition (Number constraint)
//   - Reduce: General-purpose accumulation (any operation)
//   - SumValues simpler and clearer for summation
//   - Reduce more flexible for custom aggregations
//
// SumValues 대 Reduce:
//   - SumValues: 덧셈 특화 (Number 제약)
//   - Reduce: 범용 누적 (모든 연산)
//   - SumValues가 합산에 더 간단하고 명확
//   - Reduce가 커스텀 집계에 더 유연
//
// # Performance Characteristics | 성능 특성
//
// Time Complexities:
//   - O(1): None (all operations iterate)
//   - O(n): ValuesSlice, ValuesBy, UniqueValues, FindValue (early termination),
//           ReplaceValue, UpdateValues, MinValue, MaxValue, SumValues
//   - O(n log n): ValuesSorted (due to sorting)
//
// 시간 복잡도:
//   - O(1): 없음 (모든 연산이 순회)
//   - O(n): ValuesSlice, ValuesBy, UniqueValues, FindValue (조기 종료),
//           ReplaceValue, UpdateValues, MinValue, MaxValue, SumValues
//   - O(n log n): ValuesSorted (정렬 때문)
//
// Space Complexities:
//   - O(1): MinValue, MaxValue, SumValues
//   - O(k): ValuesBy (k matching values), UniqueValues (u unique values)
//   - O(n): ValuesSlice, ValuesSorted, ReplaceValue, UpdateValues
//
// 공간 복잡도:
//   - O(1): MinValue, MaxValue, SumValues
//   - O(k): ValuesBy (k개 매칭 값), UniqueValues (u개 고유값)
//   - O(n): ValuesSlice, ValuesSorted, ReplaceValue, UpdateValues
//
// Optimization Tips:
//   - Use ValuesSlice over ValuesSorted when order doesn't matter (O(n) vs O(n log n))
//   - Use FindValue with early termination instead of ValuesBy + check
//   - Use ReplaceValue for simple substitutions, UpdateValues for complex transformations
//   - Use MinValue/MaxValue directly instead of ValuesSorted + indexing
//   - Cache ValuesSorted results if used multiple times
//
// 최적화 팁:
//   - 순서가 중요하지 않을 때 ValuesSorted 대신 ValuesSlice 사용 (O(n) 대 O(n log n))
//   - ValuesBy + 검사 대신 조기 종료하는 FindValue 사용
//   - 단순 대체는 ReplaceValue, 복잡한 변환은 UpdateValues 사용
//   - ValuesSorted + 인덱싱 대신 MinValue/MaxValue 직접 사용
//   - 여러 번 사용될 경우 ValuesSorted 결과 캐시
//
// # Common Usage Patterns | 일반적인 사용 패턴
//
// 1. Extract and Sort Values | 값 추출 및 정렬
//
//	scores := map[string]int{"alice": 95, "bob": 87, "charlie": 92}
//	sortedScores := maputil.ValuesSorted(scores)
//	// sortedScores = []int{87, 92, 95}
//	fmt.Printf("Lowest: %d, Highest: %d\n", sortedScores[0], sortedScores[len(sortedScores)-1])
//
// 2. Filter Values by Threshold | 임계값으로 값 필터링
//
//	temperatures := map[string]float64{
//	    "room1": 22.5, "room2": 18.0, "room3": 25.0, "room4": 19.5,
//	}
//	coldRooms := maputil.ValuesBy(temperatures, func(room string, temp float64) bool {
//	    return temp < 20.0
//	})
//	// coldRooms = []float64{18.0, 19.5} (order may vary)
//
// 3. Deduplicate Values | 값 중복 제거
//
//	userGroups := map[string]string{
//	    "alice": "admin", "bob": "user", "charlie": "admin", "dave": "moderator",
//	}
//	roles := maputil.UniqueValues(userGroups)
//	// roles = []string{"admin", "user", "moderator"} (order may vary)
//
// 4. Find Value Matching Condition | 조건 매칭 값 찾기
//
//	inventory := map[string]int{"apples": 50, "bananas": 0, "oranges": 30}
//	outOfStock, found := maputil.FindValue(inventory, func(item string, qty int) bool {
//	    return qty == 0
//	})
//	if found {
//	    fmt.Printf("Out of stock quantity: %d\n", outOfStock)
//	}
//
// 5. Replace Invalid Values | 잘못된 값 대체
//
//	config := map[string]int{"timeout": -1, "retries": 3, "maxConn": -1}
//	fixedConfig := maputil.ReplaceValue(config, -1, 0)
//	// fixedConfig = map[string]int{"timeout": 0, "retries": 3, "maxConn": 0}
//
// 6. Transform Values Based on Keys | 키 기반 값 변환
//
//	prices := map[string]float64{"book": 10.0, "pen": 2.0, "notebook": 5.0}
//	taxRates := map[string]float64{"book": 0.0, "pen": 0.08, "notebook": 0.08}
//
//	finalPrices := maputil.UpdateValues(prices, func(item string, price float64) float64 {
//	    if tax, ok := taxRates[item]; ok {
//	        return price * (1 + tax)
//	    }
//	    return price
//	})
//	// finalPrices = map[string]float64{"book": 10.0, "pen": 2.16, "notebook": 5.4}
//
// 7. Find Minimum and Maximum | 최솟값과 최댓값 찾기
//
//	scores := map[string]int{"alice": 95, "bob": 87, "charlie": 92}
//	minScore, _ := maputil.MinValue(scores) // minScore = 87
//	maxScore, _ := maputil.MaxValue(scores) // maxScore = 95
//	fmt.Printf("Range: %d - %d\n", minScore, maxScore)
//
// 8. Calculate Total | 합계 계산
//
//	cart := map[string]float64{"item1": 10.50, "item2": 5.25, "item3": 8.75}
//	total := maputil.SumValues(cart)
//	// total = 24.50
//	fmt.Printf("Cart Total: $%.2f\n", total)
//
// 9. Chain Value Operations | 값 작업 체인
//
//	data := map[string]int{"a": 1, "b": -2, "c": 3, "d": -4, "e": 5}
//
//	// Step 1: Replace negative values with 0
//	positive := maputil.ReplaceValue(data, -2, 0)
//	positive = maputil.ReplaceValue(positive, -4, 0)
//
//	// Step 2: Extract non-zero values
//	nonZero := maputil.ValuesBy(positive, func(k string, v int) bool {
//	    return v > 0
//	})
//
//	// Step 3: Sort values
//	sorted := maputil.ValuesSorted(maputil.MapValues(positive, func(v int) int {
//	    if v > 0 {
//	        return v
//	    }
//	    return 0
//	}))
//	// Result: sorted positive values
//
// 10. Aggregate Statistics | 집계 통계
//
//	grades := map[string]int{"math": 85, "science": 92, "history": 78, "english": 88}
//	total := maputil.SumValues(grades)
//	count := len(grades)
//	average := float64(total) / float64(count)
//	min, _ := maputil.MinValue(grades)
//	max, _ := maputil.MaxValue(grades)
//
//	fmt.Printf("Average: %.2f, Min: %d, Max: %d, Total: %d\n",
//	    average, min, max, total)
//	// Output: Average: 85.75, Min: 78, Max: 92, Total: 343
//
// # Edge Cases and Nil Handling | 엣지 케이스와 Nil 처리
//
// Empty Maps:
//   - ValuesSlice, ValuesSorted, ValuesBy, UniqueValues: Return empty slices
//   - FindValue: Returns (zero, false)
//   - ReplaceValue, UpdateValues: Return empty maps
//   - MinValue, MaxValue: Return (zero, false)
//   - SumValues: Returns zero value
//
// 빈 맵:
//   - ValuesSlice, ValuesSorted, ValuesBy, UniqueValues: 빈 슬라이스 반환
//   - FindValue: (제로값, false) 반환
//   - ReplaceValue, UpdateValues: 빈 맵 반환
//   - MinValue, MaxValue: (제로값, false) 반환
//   - SumValues: 제로값 반환
//
// Nil Maps:
//   - All functions handle nil maps gracefully
//   - Treated same as empty maps
//   - No panics or nil pointer dereferences
//
// Nil 맵:
//   - 모든 함수가 nil 맵을 우아하게 처리
//   - 빈 맵과 동일하게 처리
//   - 패닉이나 nil 포인터 역참조 없음
//
// Duplicate Values:
//   - ValuesSlice, ValuesSorted: Include all duplicates
//   - UniqueValues: Removes duplicates
//   - ReplaceValue: Replaces all occurrences
//
// 중복 값:
//   - ValuesSlice, ValuesSorted: 모든 중복 포함
//   - UniqueValues: 중복 제거
//   - ReplaceValue: 모든 발생 대체
//
// No Matching Values:
//   - ValuesBy: Returns empty slice
//   - FindValue: Returns (zero, false)
//
// 매칭 값 없음:
//   - ValuesBy: 빈 슬라이스 반환
//   - FindValue: (제로값, false) 반환
//
// # Thread Safety | 스레드 안전성
//
// Read Operations (Safe for Concurrent Reads):
//   - ValuesSlice, ValuesSorted, ValuesBy, UniqueValues
//   - FindValue
//   - MinValue, MaxValue, SumValues
//   - All extraction and mathematical operations are read-only
//
// 읽기 작업 (동시 읽기 안전):
//   - ValuesSlice, ValuesSorted, ValuesBy, UniqueValues
//   - FindValue
//   - MinValue, MaxValue, SumValues
//   - 모든 추출 및 수학 작업은 읽기 전용
//
// Write Operations (Return New Maps):
//   - ReplaceValue, UpdateValues
//   - Create new maps, don't modify originals
//   - Safe when original map has concurrent readers
//
// 쓰기 작업 (새 맵 반환):
//   - ReplaceValue, UpdateValues
//   - 새 맵 생성, 원본 수정 없음
//   - 원본 맵에 동시 읽기가 있을 때 안전
//
// Concurrent Modification Warning:
//   - Do not iterate map (ValuesSlice, etc.) while another goroutine modifies it
//   - Use explicit synchronization (mutex, RWMutex) for concurrent access
//
// 동시 수정 경고:
//   - 다른 고루틴이 수정하는 동안 맵 순회 (ValuesSlice 등) 금지
//   - 동시 접근에 명시적 동기화 (뮤텍스, RWMutex) 사용
//
// # See Also | 참고
//
// Related files in maputil package:
//   - keys.go: Key-focused operations (KeysSlice, KeysSorted, RenameKey, etc.)
//   - transform.go: Map transformation (MapValues, MapKeys, Invert, Flatten, etc.)
//   - filter.go: Filtering operations (Filter, Pick, Omit, etc.)
//   - aggregate.go: Aggregation operations (Reduce, GroupBy, Frequencies, etc.)
//   - basic.go: Fundamental map operations (Get, Set, Clone, Equal, etc.)
//
// maputil 패키지의 관련 파일:
//   - keys.go: 키 중심 작업 (KeysSlice, KeysSorted, RenameKey 등)
//   - transform.go: 맵 변환 (MapValues, MapKeys, Invert, Flatten 등)
//   - filter.go: 필터링 작업 (Filter, Pick, Omit 등)
//   - aggregate.go: 집계 작업 (Reduce, GroupBy, Frequencies 등)
//   - basic.go: 기본 맵 작업 (Get, Set, Clone, Equal 등)

// ValuesSlice returns all values from the map as a slice (alias for Values).
// ValuesSlice는 맵의 모든 값을 슬라이스로 반환합니다 (Values의 별칭).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
func ValuesSlice[K comparable, V any](m map[K]V) []V {
	return Values(m)
}

// ValuesSorted returns all values from the map as a sorted slice.
// ValuesSorted는 맵의 모든 값을 정렬된 슬라이스로 반환합니다.
//
// Time complexity: O(n log n)
// 시간 복잡도: O(n log n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	values := maputil.ValuesSorted(m) // []int{1, 2, 3}
func ValuesSorted[K comparable, V Ordered](m map[K]V) []V {
	values := Values(m)
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}

// ValuesBy returns values from the map that satisfy the predicate.
// ValuesBy는 조건을 만족하는 맵의 값을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	values := maputil.ValuesBy(m, func(k string, v int) bool {
//	    return v > 2
//	}) // []int{3, 4} (order may vary)
func ValuesBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []V {
	values := make([]V, 0)
	for k, v := range m {
		if fn(k, v) {
			values = append(values, v)
		}
	}
	return values
}

// UniqueValues returns a slice of unique values from the map.
// UniqueValues는 맵의 고유한 값들의 슬라이스를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3, "e": 2}
//	unique := maputil.UniqueValues(m) // []int{1, 2, 3} (order may vary)
func UniqueValues[K comparable, V comparable](m map[K]V) []V {
	seen := make(map[V]struct{})
	result := make([]V, 0)

	for _, v := range m {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// FindValue finds the first value that satisfies the predicate.
// FindValue는 조건을 만족하는 첫 번째 값을 찾습니다.
//
// Returns the value and true if found, or zero value and false otherwise.
// 찾으면 값과 true를 반환하고, 그렇지 않으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n) worst case, early termination when found
// 시간 복잡도: 최악의 경우 O(n), 찾으면 조기 종료
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	value, found := maputil.FindValue(m, func(k string, v int) bool {
//	    return v > 2
//	}) // value = 3, found = true (may vary due to map iteration order)
func FindValue[K comparable, V any](m map[K]V, fn func(K, V) bool) (V, bool) {
	for k, v := range m {
		if fn(k, v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

// ReplaceValue creates a new map with all occurrences of a value replaced.
// ReplaceValue는 특정 값의 모든 발생을 대체한 새 맵을 생성합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3}
//	result := maputil.ReplaceValue(m, 1, 10) // map[string]int{"a": 10, "b": 2, "c": 10, "d": 3}
func ReplaceValue[K comparable, V comparable](m map[K]V, oldValue, newValue V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		if v == oldValue {
			result[k] = newValue
		} else {
			result[k] = v
		}
	}
	return result
}

// UpdateValues creates a new map with all values transformed by the function.
// UpdateValues는 함수로 모든 값을 변환한 새 맵을 생성합니다.
//
// This is similar to MapValues but uses key-value pairs.
// 이것은 MapValues와 유사하지만 키-값 쌍을 사용합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.UpdateValues(m, func(k string, v int) int {
//	    return v * 10
//	}) // map[string]int{"a": 10, "b": 20, "c": 30}
func UpdateValues[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = fn(k, v)
	}
	return result
}

// MinValue returns the minimum value from a map.
// MinValue는 맵에서 최솟값을 반환합니다.
//
// Returns the minimum value and true if found, or zero value and false if map is empty.
// 찾은 경우 최솟값과 true를, 맵이 비어있으면 제로값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	min, found := maputil.MinValue(m) // min = 1, found = true
func MinValue[K comparable, V Ordered](m map[K]V) (V, bool) {
	if len(m) == 0 {
		var zero V
		return zero, false
	}

	var min V
	first := true
	for _, v := range m {
		if first || v < min {
			min = v
			first = false
		}
	}
	return min, true
}

// MaxValue returns the maximum value from a map.
// MaxValue는 맵에서 최댓값을 반환합니다.
//
// Returns the maximum value and true if found, or zero value and false if map is empty.
// 찾은 경우 최댓값과 true를, 맵이 비어있으면 제로값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	max, found := maputil.MaxValue(m) // max = 3, found = true
func MaxValue[K comparable, V Ordered](m map[K]V) (V, bool) {
	if len(m) == 0 {
		var zero V
		return zero, false
	}

	var max V
	first := true
	for _, v := range m {
		if first || v > max {
			max = v
			first = false
		}
	}
	return max, true
}

// SumValues returns the sum of all values in a map.
// SumValues는 맵의 모든 값의 합을 반환합니다.
//
// Returns zero if the map is empty.
// 맵이 비어있으면 0을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	sum := maputil.SumValues(m) // sum = 6
func SumValues[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
