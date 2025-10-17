package maputil

// aggregate.go provides map aggregation and reduction operations for Go.
//
// This file implements functions that reduce maps to single values or aggregate
// data in various ways. Operations include mathematical calculations, statistical
// analysis, and grouping operations.
//
// Aggregation Categories:
//
// General Reduction:
//   - Reduce(m, initial, fn): Flexible aggregation with accumulator
//     Time: O(n), Space: O(1)
//     Most general aggregation function
//     Function receives accumulator, key, and value
//     Returns final accumulated result
//     Example: Reduce({"a": 1, "b": 2}, 0, func(acc, k, v) { return acc + v }) = 3
//     Use cases: Custom aggregations, complex calculations, data transformation
//
// Mathematical Aggregations:
//   - Sum(m): Total of all numeric values
//     Time: O(n), Space: O(1)
//     Simple addition of all values
//     Returns zero for empty map
//     Example: Sum({"a": 1, "b": 2, "c": 3}) = 6
//     Use cases: Total calculations, simple statistics, data summaries
//
//   - Average(m): Mean of all numeric values
//     Time: O(n), Space: O(1)
//     Sum divided by count
//     Returns 0 for empty map
//     Example: Average({"a": 1, "b": 2, "c": 3, "d": 4}) = 2.5
//     Use cases: Statistical analysis, performance metrics, grade calculations
//
// Min/Max Operations:
//   - Min(m): Find minimum value entry
//     Time: O(n), Space: O(1)
//     Returns (key, value, true) for min entry
//     Returns (zero, zero, false) for empty map
//     Requires Ordered constraint
//     Example: Min({"a": 3, "b": 1, "c": 2}) = ("b", 1, true)
//     Use cases: Finding lowest value, range analysis, optimization
//
//   - Max(m): Find maximum value entry
//     Time: O(n), Space: O(1)
//     Returns (key, value, true) for max entry
//     Returns (zero, zero, false) for empty map
//     Requires Ordered constraint
//     Example: Max({"a": 3, "b": 1, "c": 2}) = ("a", 3, true)
//     Use cases: Finding highest value, peak detection, best selection
//
//   - MinBy(m, fn): Find minimum by custom score
//     Time: O(n), Space: O(1)
//     Function extracts float64 score from value
//     Returns entry with lowest score
//     Example: MinBy(users, func(u) { return float64(u.Age) }) = youngest user
//     Use cases: Custom comparisons, complex min, sorting criteria
//
//   - MaxBy(m, fn): Find maximum by custom score
//     Time: O(n), Space: O(1)
//     Function extracts float64 score from value
//     Returns entry with highest score
//     Example: MaxBy(users, func(u) { return float64(u.Score) }) = highest scorer
//     Use cases: Custom comparisons, complex max, ranking
//
// Statistical Operations:
//   - Median(m): Middle value of sorted values
//     Time: O(n log n), Space: O(n)
//     Sorts all values, returns middle
//     For even count: average of two middle values
//     Returns (0, false) for empty map
//     Example: Median({"a": 1, "b": 3, "c": 2}) = (2.0, true)
//     Use cases: Central tendency, robust statistics, outlier resistance
//
//   - Frequencies(m): Count occurrences of each value
//     Time: O(n), Space: O(u) where u = unique values
//     Inverts map: value → count
//     Returns new map of value frequencies
//     Example: Frequencies({"a": "X", "b": "Y", "c": "X"}) = {"X": 2, "Y": 1}
//     Use cases: Distribution analysis, duplicate detection, histogram data
//
// Grouping Operations:
//   - GroupBy(slice, fn): Group slice elements by key
//     Time: O(n), Space: O(n)
//     Not a map→map function (slice→map)
//     Function extracts grouping key from element
//     Returns map[Key][]Element
//     Example: GroupBy(users, func(u) { return u.City }) = map by city
//     Use cases: Data categorization, SQL-like GROUP BY, data organization
//
//   - CountBy(slice, fn): Count elements by key
//     Time: O(n), Space: O(u) where u = unique keys
//     Similar to GroupBy but returns counts
//     Function extracts grouping key from element
//     Returns map[Key]int
//     Example: CountBy(users, func(u) { return u.City }) = city counts
//     Use cases: Frequency analysis, distribution counting, aggregation stats
//
// Design Principles:
//   - Generality: Reduce provides foundation for custom aggregations
//   - Type Safety: Number and Ordered constraints for mathematical ops
//   - Consistency: Empty maps return sensible defaults (0, false, etc.)
//   - Flexibility: MinBy/MaxBy for complex custom comparisons
//   - Efficiency: O(n) operations except Median (requires sorting)
//
// Comparison: Sum vs Reduce:
//   - Sum: Specialized, only addition, simpler
//   - Reduce: General, any operation, more flexible
//   - Sum: More readable for simple totals
//   - Reduce: Necessary for complex aggregations
//
// Comparison: Min/Max vs MinBy/MaxBy:
//   - Min/Max: Direct value comparison
//   - MinBy/MaxBy: Custom score extraction
//   - Min/Max: Simpler, requires Ordered
//   - MinBy/MaxBy: More flexible, works with any type
//
// Comparison: Average vs Median:
//   - Average: Mean, affected by outliers
//   - Median: Middle value, robust to outliers
//   - Average: O(n), simpler calculation
//   - Median: O(n log n), requires sorting
//   - Average: Better for normal distributions
//   - Median: Better for skewed data
//
// Comparison: GroupBy vs Frequencies:
//   - GroupBy: Collects elements into groups
//   - Frequencies: Counts occurrences
//   - GroupBy: Returns []V (preserves data)
//   - Frequencies: Returns int (just counts)
//   - GroupBy: More memory, full data access
//   - Frequencies: Less memory, summary only
//
// Performance Characteristics:
//
// Time Complexity:
//   - Reduce/Sum/Average: O(n) - Single pass
//   - Min/Max/MinBy/MaxBy: O(n) - Single pass
//   - Frequencies/GroupBy/CountBy: O(n) - Single pass
//   - Median: O(n log n) - Requires sorting
//
// Space Complexity:
//   - Reduce/Sum/Average: O(1) - No allocation
//   - Min/Max/MinBy/MaxBy: O(1) - No allocation
//   - Frequencies: O(u) where u = unique values
//   - GroupBy: O(n) - Stores all elements
//   - CountBy: O(u) where u = unique keys
//   - Median: O(n) - Copies all values for sorting
//
// Memory Allocation:
//   - Math operations (Sum, Average, Min, Max): No heap allocation
//   - Median: Allocates slice for all values
//   - GroupBy: Allocates slice for each group
//   - Frequencies/CountBy: Allocates result map
//   - Reduce: Depends on accumulator function
//
// Common Usage Patterns:
//
//	// Calculate total
//	prices := map[string]float64{"apple": 1.5, "banana": 0.8, "orange": 2.0}
//	total := maputil.Sum(prices) // 4.3
//
//	// Find best/worst
//	scores := map[string]int{"Alice": 85, "Bob": 92, "Charlie": 78}
//	bestName, bestScore, _ := maputil.Max(scores) // "Bob", 92
//	worstName, worstScore, _ := maputil.Min(scores) // "Charlie", 78
//
//	// Statistical analysis
//	testScores := map[string]int{"t1": 85, "t2": 90, "t3": 75, "t4": 95, "t5": 80}
//	avg := maputil.Average(testScores) // 85.0
//	med, _ := maputil.Median(testScores) // 85.0
//
//	// Custom aggregation with Reduce
//	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	product := maputil.Reduce(data, 1, func(acc int, k string, v int) int {
//	    return acc * v
//	}) // 1 * 1 * 2 * 3 * 4 = 24
//
//	// Group users by city
//	type User struct { Name string; City string; Age int }
//	users := []User{
//	    {Name: "Alice", City: "Seoul", Age: 30},
//	    {Name: "Bob", City: "Seoul", Age: 25},
//	    {Name: "Charlie", City: "Busan", Age: 35},
//	}
//	byCity := maputil.GroupBy(users, func(u User) string {
//	    return u.City
//	})
//	// byCity["Seoul"] = [Alice, Bob]
//	// byCity["Busan"] = [Charlie]
//
//	// Count distribution
//	grades := []string{"A", "B", "A", "C", "B", "A", "D"}
//	counts := maputil.CountBy(grades, func(g string) string { return g })
//	// counts = {"A": 3, "B": 2, "C": 1, "D": 1}
//
//	// Find youngest/oldest with custom comparison
//	userMap := map[string]User{
//	    "u1": {Name: "Alice", Age: 30},
//	    "u2": {Name: "Bob", Age: 25},
//	    "u3": {Name: "Charlie", Age: 35},
//	}
//	_, youngest, _ := maputil.MinBy(userMap, func(u User) float64 {
//	    return float64(u.Age)
//	}) // youngest = Bob (age 25)
//
//	// Value frequency analysis
//	colors := map[string]string{
//	    "item1": "red", "item2": "blue", "item3": "red",
//	    "item4": "green", "item5": "red", "item6": "blue",
//	}
//	colorFreq := maputil.Frequencies(colors)
//	// colorFreq = {"red": 3, "blue": 2, "green": 1}
//
// Advanced Reduce Patterns:
//
//	// Concatenate all values
//	names := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie"}
//	allNames := maputil.Reduce(names, "", func(acc string, k int, v string) string {
//	    if acc == "" {
//	        return v
//	    }
//	    return acc + ", " + v
//	}) // "Alice, Bob, Charlie" (order may vary)
//
//	// Build complex structure
//	scores := map[string]int{"Alice": 85, "Bob": 92}
//	summary := maputil.Reduce(scores, map[string]interface{}{},
//	    func(acc map[string]interface{}, k string, v int) map[string]interface{} {
//	        acc[k] = map[string]interface{}{
//	            "score": v,
//	            "pass": v >= 80,
//	        }
//	        return acc
//	    })
//
// Edge Cases:
//   - Sum on empty map: Returns 0 (zero value)
//   - Average on empty map: Returns 0
//   - Min/Max on empty map: Returns (zero, zero, false)
//   - MinBy/MaxBy on empty map: Returns (zero, zero, false)
//   - Median on empty map: Returns (0, false)
//   - Median on single element: Returns that element
//   - Median on two elements: Returns average of both
//   - Frequencies on empty map: Returns empty map
//   - GroupBy on empty slice: Returns empty map
//   - CountBy on empty slice: Returns empty map
//   - Reduce on empty map: Returns initial value unchanged
//
// Nil Map Behavior:
//   - Sum(nil): Returns 0
//   - Average(nil): Returns 0
//   - Min(nil): Returns (zero, zero, false)
//   - Max(nil): Returns (zero, zero, false)
//   - MinBy(nil, fn): Returns (zero, zero, false)
//   - MaxBy(nil, fn): Returns (zero, zero, false)
//   - Median(nil): Returns (0, false)
//   - Frequencies(nil): Returns empty map
//   - Reduce(nil, init, fn): Returns init unchanged
//
// Thread Safety:
//   - All read-only operations (safe for concurrent reads of input)
//   - Not safe if input map is modified concurrently
//   - Reduce/MinBy/MaxBy functions must be goroutine-safe
//   - GroupBy/CountBy process slices (no map concurrency issues)
//
// Performance Tips:
//   - Sum/Average are faster than Reduce for simple math
//   - Min/Max are faster than MinBy when direct comparison works
//   - Median is expensive (O(n log n)) - cache if used multiple times
//   - GroupBy allocates slices - consider CountBy if only counts needed
//   - Reduce is flexible but has function call overhead
//   - For large maps, consider streaming aggregations
//
// Statistical Considerations:
//   - Average: Sensitive to outliers, use Median for robustness
//   - Median: Better for skewed distributions
//   - Median on even count: Returns average of middle two
//   - Frequencies: Useful for finding mode (most common value)
//   - Consider caching Median result (expensive to recompute)
//
// aggregate.go는 Go를 위한 맵 집계 및 축소 작업을 제공합니다.
//
// 이 파일은 맵을 단일 값으로 축소하거나 다양한 방식으로 데이터를 집계하는
// 함수를 구현합니다. 작업에는 수학적 계산, 통계 분석 및 그룹화 작업이 포함됩니다.
//
// 집계 카테고리:
//
// 일반 축소:
//   - Reduce(m, initial, fn): 누산기를 사용한 유연한 집계
//     시간: O(n), 공간: O(1)
//     가장 일반적인 집계 함수
//     함수가 누산기, 키, 값 수신
//     최종 누적 결과 반환
//     예: Reduce({"a": 1, "b": 2}, 0, func(acc, k, v) { return acc + v }) = 3
//     사용 사례: 사용자 정의 집계, 복잡한 계산, 데이터 변환
//
// 수학적 집계:
//   - Sum(m): 모든 숫자 값의 합계
//     시간: O(n), 공간: O(1)
//     모든 값의 단순 덧셈
//     빈 맵의 경우 zero 반환
//     예: Sum({"a": 1, "b": 2, "c": 3}) = 6
//     사용 사례: 합계 계산, 간단한 통계, 데이터 요약
//
//   - Average(m): 모든 숫자 값의 평균
//     시간: O(n), 공간: O(1)
//     합계를 개수로 나눔
//     빈 맵의 경우 0 반환
//     예: Average({"a": 1, "b": 2, "c": 3, "d": 4}) = 2.5
//     사용 사례: 통계 분석, 성능 메트릭, 성적 계산
//
// Min/Max 작업:
//   - Min(m): 최소값 항목 찾기
//     시간: O(n), 공간: O(1)
//     최소 항목에 대해 (키, 값, true) 반환
//     빈 맵의 경우 (zero, zero, false) 반환
//     Ordered 제약 필요
//     예: Min({"a": 3, "b": 1, "c": 2}) = ("b", 1, true)
//     사용 사례: 최저값 찾기, 범위 분석, 최적화
//
//   - Max(m): 최대값 항목 찾기
//     시간: O(n), 공간: O(1)
//     최대 항목에 대해 (키, 값, true) 반환
//     빈 맵의 경우 (zero, zero, false) 반환
//     Ordered 제약 필요
//     예: Max({"a": 3, "b": 1, "c": 2}) = ("a", 3, true)
//     사용 사례: 최고값 찾기, 피크 감지, 최선 선택
//
//   - MinBy(m, fn): 사용자 정의 점수로 최소값 찾기
//     시간: O(n), 공간: O(1)
//     함수가 값에서 float64 점수 추출
//     가장 낮은 점수의 항목 반환
//     예: MinBy(users, func(u) { return float64(u.Age) }) = 가장 어린 사용자
//     사용 사례: 사용자 정의 비교, 복잡한 최소값, 정렬 기준
//
//   - MaxBy(m, fn): 사용자 정의 점수로 최대값 찾기
//     시간: O(n), 공간: O(1)
//     함수가 값에서 float64 점수 추출
//     가장 높은 점수의 항목 반환
//     예: MaxBy(users, func(u) { return float64(u.Score) }) = 최고 득점자
//     사용 사례: 사용자 정의 비교, 복잡한 최대값, 순위 매기기
//
// 통계 작업:
//   - Median(m): 정렬된 값의 중앙값
//     시간: O(n log n), 공간: O(n)
//     모든 값 정렬, 중간 반환
//     짝수 개수의 경우: 두 중간 값의 평균
//     빈 맵의 경우 (0, false) 반환
//     예: Median({"a": 1, "b": 3, "c": 2}) = (2.0, true)
//     사용 사례: 중심 경향, 강건한 통계, 이상값 저항
//
//   - Frequencies(m): 각 값의 출현 횟수 계산
//     시간: O(n), 공간: O(u) (u = 고유 값)
//     맵 반전: 값 → 개수
//     값 빈도의 새 맵 반환
//     예: Frequencies({"a": "X", "b": "Y", "c": "X"}) = {"X": 2, "Y": 1}
//     사용 사례: 분포 분석, 중복 감지, 히스토그램 데이터
//
// 그룹화 작업:
//   - GroupBy(slice, fn): 키로 슬라이스 요소 그룹화
//     시간: O(n), 공간: O(n)
//     맵→맵 함수가 아님 (슬라이스→맵)
//     함수가 요소에서 그룹화 키 추출
//     map[Key][]Element 반환
//     예: GroupBy(users, func(u) { return u.City }) = 도시별 맵
//     사용 사례: 데이터 분류, SQL 같은 GROUP BY, 데이터 구성
//
//   - CountBy(slice, fn): 키별 요소 개수 계산
//     시간: O(n), 공간: O(u) (u = 고유 키)
//     GroupBy와 유사하지만 개수 반환
//     함수가 요소에서 그룹화 키 추출
//     map[Key]int 반환
//     예: CountBy(users, func(u) { return u.City }) = 도시별 개수
//     사용 사례: 빈도 분석, 분포 계산, 집계 통계
//
// 설계 원칙:
//   - 일반성: Reduce가 사용자 정의 집계의 기초 제공
//   - 타입 안전성: 수학 연산을 위한 Number 및 Ordered 제약
//   - 일관성: 빈 맵은 합리적인 기본값 반환 (0, false 등)
//   - 유연성: 복잡한 사용자 정의 비교를 위한 MinBy/MaxBy
//   - 효율성: Median 제외 O(n) 작업 (정렬 필요)
//
// 비교: Sum vs Reduce:
//   - Sum: 특화됨, 덧셈만, 더 단순
//   - Reduce: 일반적, 모든 작업, 더 유연
//   - Sum: 간단한 합계에 더 읽기 쉬움
//   - Reduce: 복잡한 집계에 필요
//
// 비교: Min/Max vs MinBy/MaxBy:
//   - Min/Max: 직접 값 비교
//   - MinBy/MaxBy: 사용자 정의 점수 추출
//   - Min/Max: 더 단순, Ordered 필요
//   - MinBy/MaxBy: 더 유연, 모든 타입 작동
//
// 비교: Average vs Median:
//   - Average: 평균, 이상값 영향 받음
//   - Median: 중앙값, 이상값에 강건
//   - Average: O(n), 더 단순한 계산
//   - Median: O(n log n), 정렬 필요
//   - Average: 정규 분포에 더 좋음
//   - Median: 왜곡된 데이터에 더 좋음
//
// 비교: GroupBy vs Frequencies:
//   - GroupBy: 요소를 그룹으로 수집
//   - Frequencies: 출현 횟수 계산
//   - GroupBy: []V 반환 (데이터 보존)
//   - Frequencies: int 반환 (개수만)
//   - GroupBy: 더 많은 메모리, 전체 데이터 접근
//   - Frequencies: 더 적은 메모리, 요약만
//
// 성능 특성:
//
// 시간 복잡도:
//   - Reduce/Sum/Average: O(n) - 단일 패스
//   - Min/Max/MinBy/MaxBy: O(n) - 단일 패스
//   - Frequencies/GroupBy/CountBy: O(n) - 단일 패스
//   - Median: O(n log n) - 정렬 필요
//
// 공간 복잡도:
//   - Reduce/Sum/Average: O(1) - 할당 없음
//   - Min/Max/MinBy/MaxBy: O(1) - 할당 없음
//   - Frequencies: O(u) (u = 고유 값)
//   - GroupBy: O(n) - 모든 요소 저장
//   - CountBy: O(u) (u = 고유 키)
//   - Median: O(n) - 정렬을 위해 모든 값 복사
//
// 메모리 할당:
//   - 수학 연산 (Sum, Average, Min, Max): 힙 할당 없음
//   - Median: 모든 값에 대한 슬라이스 할당
//   - GroupBy: 각 그룹에 대한 슬라이스 할당
//   - Frequencies/CountBy: 결과 맵 할당
//   - Reduce: 누산기 함수에 따라 다름
//
// 일반적인 사용 패턴:
//
//	// 합계 계산
//	prices := map[string]float64{"apple": 1.5, "banana": 0.8, "orange": 2.0}
//	total := maputil.Sum(prices) // 4.3
//
//	// 최고/최저 찾기
//	scores := map[string]int{"Alice": 85, "Bob": 92, "Charlie": 78}
//	bestName, bestScore, _ := maputil.Max(scores) // "Bob", 92
//	worstName, worstScore, _ := maputil.Min(scores) // "Charlie", 78
//
//	// 통계 분석
//	testScores := map[string]int{"t1": 85, "t2": 90, "t3": 75, "t4": 95, "t5": 80}
//	avg := maputil.Average(testScores) // 85.0
//	med, _ := maputil.Median(testScores) // 85.0
//
//	// Reduce로 사용자 정의 집계
//	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	product := maputil.Reduce(data, 1, func(acc int, k string, v int) int {
//	    return acc * v
//	}) // 1 * 1 * 2 * 3 * 4 = 24
//
//	// 도시별로 사용자 그룹화
//	type User struct { Name string; City string; Age int }
//	users := []User{
//	    {Name: "Alice", City: "Seoul", Age: 30},
//	    {Name: "Bob", City: "Seoul", Age: 25},
//	    {Name: "Charlie", City: "Busan", Age: 35},
//	}
//	byCity := maputil.GroupBy(users, func(u User) string {
//	    return u.City
//	})
//	// byCity["Seoul"] = [Alice, Bob]
//	// byCity["Busan"] = [Charlie]
//
//	// 분포 계산
//	grades := []string{"A", "B", "A", "C", "B", "A", "D"}
//	counts := maputil.CountBy(grades, func(g string) string { return g })
//	// counts = {"A": 3, "B": 2, "C": 1, "D": 1}
//
//	// 사용자 정의 비교로 최연소/최고령 찾기
//	userMap := map[string]User{
//	    "u1": {Name: "Alice", Age: 30},
//	    "u2": {Name: "Bob", Age: 25},
//	    "u3": {Name: "Charlie", Age: 35},
//	}
//	_, youngest, _ := maputil.MinBy(userMap, func(u User) float64 {
//	    return float64(u.Age)
//	}) // youngest = Bob (age 25)
//
//	// 값 빈도 분석
//	colors := map[string]string{
//	    "item1": "red", "item2": "blue", "item3": "red",
//	    "item4": "green", "item5": "red", "item6": "blue",
//	}
//	colorFreq := maputil.Frequencies(colors)
//	// colorFreq = {"red": 3, "blue": 2, "green": 1}
//
// 고급 Reduce 패턴:
//
//	// 모든 값 연결
//	names := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie"}
//	allNames := maputil.Reduce(names, "", func(acc string, k int, v string) string {
//	    if acc == "" {
//	        return v
//	    }
//	    return acc + ", " + v
//	}) // "Alice, Bob, Charlie" (순서는 다를 수 있음)
//
//	// 복잡한 구조 구축
//	scores := map[string]int{"Alice": 85, "Bob": 92}
//	summary := maputil.Reduce(scores, map[string]interface{}{},
//	    func(acc map[string]interface{}, k string, v int) map[string]interface{} {
//	        acc[k] = map[string]interface{}{
//	            "score": v,
//	            "pass": v >= 80,
//	        }
//	        return acc
//	    })
//
// 엣지 케이스:
//   - 빈 맵에 Sum: 0 반환 (zero 값)
//   - 빈 맵에 Average: 0 반환
//   - 빈 맵에 Min/Max: (zero, zero, false) 반환
//   - 빈 맵에 MinBy/MaxBy: (zero, zero, false) 반환
//   - 빈 맵에 Median: (0, false) 반환
//   - 단일 요소에 Median: 그 요소 반환
//   - 두 요소에 Median: 둘의 평균 반환
//   - 빈 맵에 Frequencies: 빈 맵 반환
//   - 빈 슬라이스에 GroupBy: 빈 맵 반환
//   - 빈 슬라이스에 CountBy: 빈 맵 반환
//   - 빈 맵에 Reduce: 초기값 변경 없이 반환
//
// Nil 맵 동작:
//   - Sum(nil): 0 반환
//   - Average(nil): 0 반환
//   - Min(nil): (zero, zero, false) 반환
//   - Max(nil): (zero, zero, false) 반환
//   - MinBy(nil, fn): (zero, zero, false) 반환
//   - MaxBy(nil, fn): (zero, zero, false) 반환
//   - Median(nil): (0, false) 반환
//   - Frequencies(nil): 빈 맵 반환
//   - Reduce(nil, init, fn): init 변경 없이 반환
//
// 스레드 안전성:
//   - 모든 읽기 전용 작업 (입력의 동시 읽기 안전)
//   - 입력 맵이 동시에 수정되면 안전하지 않음
//   - Reduce/MinBy/MaxBy 함수는 고루틴 안전해야 함
//   - GroupBy/CountBy는 슬라이스 처리 (맵 동시성 문제 없음)
//
// 성능 팁:
//   - 간단한 수학의 경우 Reduce보다 Sum/Average가 빠름
//   - 직접 비교가 가능하면 MinBy보다 Min/Max가 빠름
//   - Median은 비쌈 (O(n log n)) - 여러 번 사용 시 캐시
//   - GroupBy는 슬라이스 할당 - 개수만 필요하면 CountBy 고려
//   - Reduce는 유연하지만 함수 호출 오버헤드 있음
//   - 큰 맵의 경우 스트리밍 집계 고려
//
// 통계 고려사항:
//   - Average: 이상값에 민감, 강건성을 위해 Median 사용
//   - Median: 왜곡된 분포에 더 좋음
//   - 짝수 개수에 Median: 중간 두 개의 평균 반환
//   - Frequencies: 최빈값 찾기에 유용 (가장 흔한 값)
//   - Median 결과 캐싱 고려 (재계산 비쌈)

// Reduce reduces a map to a single value using an accumulator function.
// Reduce는 누산기 함수를 사용하여 맵을 단일 값으로 축소합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	sum := maputil.Reduce(m, 0, func(acc int, k string, v int) int {
//	    return acc + v
//	}) // sum = 6
func Reduce[K comparable, V any, R any](m map[K]V, initial R, fn func(R, K, V) R) R {
	result := initial
	for k, v := range m {
		result = fn(result, k, v)
	}
	return result
}

// Sum calculates the sum of all numeric values in the map.
// Sum은 맵의 모든 숫자 값의 합을 계산합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	total := maputil.Sum(m) // total = 6
func Sum[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

// Min finds the key-value pair with the minimum value.
// Min은 최소값을 가진 키-값 쌍을 찾습니다.
//
// Returns the key, value, and true if the map is not empty.
// Returns zero values and false if the map is empty.
//
// 맵이 비어 있지 않으면 키, 값, true를 반환합니다.
// 맵이 비어 있으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	key, value, ok := maputil.Min(m) // key = "b", value = 1, ok = true
func Min[K comparable, V Ordered](m map[K]V) (K, V, bool) {
	if len(m) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var minKey K
	var minValue V
	first := true

	for k, v := range m {
		if first || v < minValue {
			minKey = k
			minValue = v
			first = false
		}
	}

	return minKey, minValue, true
}

// Max finds the key-value pair with the maximum value.
// Max는 최대값을 가진 키-값 쌍을 찾습니다.
//
// Returns the key, value, and true if the map is not empty.
// Returns zero values and false if the map is empty.
//
// 맵이 비어 있지 않으면 키, 값, true를 반환합니다.
// 맵이 비어 있으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	key, value, ok := maputil.Max(m) // key = "a", value = 3, ok = true
func Max[K comparable, V Ordered](m map[K]V) (K, V, bool) {
	if len(m) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var maxKey K
	var maxValue V
	first := true

	for k, v := range m {
		if first || v > maxValue {
			maxKey = k
			maxValue = v
			first = false
		}
	}

	return maxKey, maxValue, true
}

// MinBy finds the key-value pair with the minimum value according to a custom function.
// MinBy는 사용자 정의 함수에 따라 최소값을 가진 키-값 쌍을 찾습니다.
//
// Returns the key, value, and true if the map is not empty.
// Returns zero values and false if the map is empty.
//
// 맵이 비어 있지 않으면 키, 값, true를 반환합니다.
// 맵이 비어 있으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { Name string; Age int }
//	m := map[string]User{
//	    "alice": {Name: "Alice", Age: 30},
//	    "bob":   {Name: "Bob", Age: 25},
//	}
//	key, user, ok := maputil.MinBy(m, func(u User) float64 {
//	    return float64(u.Age)
//	}) // key = "bob", user.Age = 25, ok = true
func MinBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool) {
	if len(m) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var minKey K
	var minValue V
	var minScore float64
	first := true

	for k, v := range m {
		score := fn(v)
		if first || score < minScore {
			minKey = k
			minValue = v
			minScore = score
			first = false
		}
	}

	return minKey, minValue, true
}

// MaxBy finds the key-value pair with the maximum value according to a custom function.
// MaxBy는 사용자 정의 함수에 따라 최대값을 가진 키-값 쌍을 찾습니다.
//
// Returns the key, value, and true if the map is not empty.
// Returns zero values and false if the map is empty.
//
// 맵이 비어 있지 않으면 키, 값, true를 반환합니다.
// 맵이 비어 있으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { Name string; Score int }
//	m := map[string]User{
//	    "alice": {Name: "Alice", Score: 95},
//	    "bob":   {Name: "Bob", Score: 88},
//	}
//	key, user, ok := maputil.MaxBy(m, func(u User) float64 {
//	    return float64(u.Score)
//	}) // key = "alice", user.Score = 95, ok = true
func MaxBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool) {
	if len(m) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var maxKey K
	var maxValue V
	var maxScore float64
	first := true

	for k, v := range m {
		score := fn(v)
		if first || score > maxScore {
			maxKey = k
			maxValue = v
			maxScore = score
			first = false
		}
	}

	return maxKey, maxValue, true
}

// Average calculates the average of all numeric values in the map.
// Average는 맵의 모든 숫자 값의 평균을 계산합니다.
//
// Returns 0 if the map is empty.
// 맵이 비어 있으면 0을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	avg := maputil.Average(m) // avg = 2.5
func Average[K comparable, V Number](m map[K]V) float64 {
	if len(m) == 0 {
		return 0
	}

	var sum V
	for _, v := range m {
		sum += v
	}

	return float64(sum) / float64(len(m))
}

// GroupBy groups a slice of elements by a key extracted from each element.
// GroupBy는 각 요소에서 추출한 키로 요소 슬라이스를 그룹화합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { Name string; City string }
//	users := []User{
//	    {Name: "Alice", City: "Seoul"},
//	    {Name: "Bob", City: "Seoul"},
//	    {Name: "Charlie", City: "Busan"},
//	}
//	byCity := maputil.GroupBy(users, func(u User) string {
//	    return u.City
//	})
//	// byCity = map[string][]User{
//	//     "Seoul": []User{{Name: "Alice", ...}, {Name: "Bob", ...}},
//	//     "Busan": []User{{Name: "Charlie", ...}},
//	// }
func GroupBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G][]V {
	result := make(map[G][]V)

	for _, item := range slice {
		key := fn(item)
		result[key] = append(result[key], item)
	}

	return result
}

// CountBy counts the number of elements in a slice for each key extracted from each element.
// CountBy는 각 요소에서 추출한 키별로 슬라이스의 요소 수를 계산합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { Name string; City string }
//	users := []User{
//	    {Name: "Alice", City: "Seoul"},
//	    {Name: "Bob", City: "Seoul"},
//	    {Name: "Charlie", City: "Busan"},
//	}
//	counts := maputil.CountBy(users, func(u User) string {
//	    return u.City
//	}) // counts = map[string]int{"Seoul": 2, "Busan": 1}
func CountBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G]int {
	result := make(map[G]int)

	for _, item := range slice {
		key := fn(item)
		result[key]++
	}

	return result
}

// Median calculates the median of all numeric values in the map.
// Median는 맵의 모든 숫자 값의 중앙값을 계산합니다.
//
// This function collects all values, sorts them, and returns the middle value.
// For even-length maps, it returns the average of the two middle values.
// Returns (0, false) for empty maps.
//
// 이 함수는 모든 값을 수집하고 정렬한 후 중간 값을 반환합니다.
// 짝수 길이 맵의 경우 두 중간 값의 평균을 반환합니다.
// 빈 맵의 경우 (0, false)를 반환합니다.
//
// Time Complexity
// 시간 복잡도: O(n log n) due to sorting
// Space Complexity
// 공간 복잡도: O(n) for collecting values
//
// Parameters
// 매개변수:
// - m: The input map with numeric values
// 숫자 값이 있는 입력 맵
//
// Returns
// 반환값:
// - float64: The median value
// 중앙값
// - bool: false if map is empty, true otherwise
// 맵이 비어있으면 false, 그렇지 않으면 true
//
// Example
// 예제:
//
//	scores := map[string]int{
//		"Alice":   85,
//		"Bob":     90,
//		"Charlie": 75,
//		"Diana":   95,
//		"Eve":     80,
//	}
//	median, ok := maputil.Median(scores) // median = 85.0 (middle value when sorted)
//
//	evenScores := map[string]int{
//		"Alice": 80,
//		"Bob":   90,
//		"Charlie": 70,
//		"Diana": 100,
//	}
//	median2, ok2 := maputil.Median(evenScores) // median2 = 85.0 (average of 80 and 90)
//
// Use Case
// 사용 사례:
// - Statistical analysis
// 통계 분석
// - Grade distribution analysis
// 성적 분포 분석
// - Performance metrics
// 성능 메트릭
// - Finding typical values
// 대표값 찾기
func Median[K comparable, V Number](m map[K]V) (float64, bool) {
	if len(m) == 0 {
		return 0, false
	}

	// Collect all values into a slice
	values := make([]float64, 0, len(m))
	for _, v := range m {
		values = append(values, float64(v))
	}

	// Sort values
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}

	n := len(values)
	if n%2 == 1 {
		// Odd length: return middle value
		return values[n/2], true
	}

	// Even length: return average of two middle values
	mid1 := values[n/2-1]
	mid2 := values[n/2]
	return (mid1 + mid2) / 2, true
}

// Frequencies counts the occurrence of each unique value in the map.
// Frequencies는 맵의 각 고유 값의 출현 빈도를 계산합니다.
//
// This function inverts the map structure, creating a map where keys are the
// original values and values are their occurrence counts. Useful for finding
// duplicate values or analyzing value distributions.
//
// 이 함수는 맵 구조를 반전하여 키가 원래 값이고 값이 출현 횟수인 맵을 생성합니다.
// 중복 값을 찾거나 값 분포를 분석하는 데 유용합니다.
//
// Time Complexity
// 시간 복잡도: O(n)
// Space Complexity
// 공간 복잡도: O(u) where u is unique values
//
// Parameters
// 매개변수:
// - m: The input map
// 입력 맵
//
// Returns
// 반환값:
// - map[V]int: Map of value → count
// 값 → 개수 맵
//
// Example
// 예제:
//
//	grades := map[string]string{
//		"Alice":   "A",
//		"Bob":     "B",
//		"Charlie": "A",
//		"Diana":   "C",
//		"Eve":     "B",
//		"Frank":   "A",
//	}
//	freq := maputil.Frequencies(grades)
//	// freq = map[string]int{"A": 3, "B": 2, "C": 1}
//
//	scores := map[string]int{
//		"test1": 85,
//		"test2": 90,
//		"test3": 85,
//		"test4": 90,
//		"test5": 75,
//	}
//	scoreFreq := maputil.Frequencies(scores)
//	// scoreFreq = map[int]int{85: 2, 90: 2, 75: 1}
//
// Use Case
// 사용 사례:
// - Finding duplicate values
// 중복 값 찾기
// - Value distribution analysis
// 값 분포 분석
// - Histogram generation
// 히스토그램 생성
// - Data quality checks
// 데이터 품질 확인
func Frequencies[K comparable, V comparable](m map[K]V) map[V]int {
	result := make(map[V]int)

	for _, value := range m {
		result[value]++
	}

	return result
}
