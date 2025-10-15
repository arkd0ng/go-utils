package main

import (
	"context"
	"fmt"
	"time"

	"github.com/arkd0ng/go-utils/logging"
	"github.com/arkd0ng/go-utils/maputil"
)

func main() {
	// Initialize logger / 로거 초기화
	logger, err := logging.New(
		logging.WithLevel(logging.DEBUG),
		logging.WithFilePath(fmt.Sprintf("logs/maputil-example-%s.log", time.Now().Format("20060102_150405"))),
		logging.WithStdout(true),
	)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	// Display banner / 배너 표시
	logger.Banner("maputil Package Examples", maputil.Version)
	logger.Info("==============================================================================")
	logger.Info("This example demonstrates all 81 functions in the maputil package")
	logger.Info("이 예제는 maputil 패키지의 모든 81개 함수를 시연합니다")
	logger.Info("==============================================================================")
	logger.Info("")

	// Run all examples / 모든 예제 실행
	ctx := context.Background()

	// Category 1: Basic Operations (11 functions) / 기본 작업 (11개 함수)
	basicOperations(ctx, logger)

	// Category 2: Transformation (10 functions) / 변환 (10개 함수)
	transformations(ctx, logger)

	// Category 3: Aggregation (9 functions) / 집계 (9개 함수)
	aggregations(ctx, logger)

	// Category 4: Merge Operations (8 functions) / 병합 작업 (8개 함수)
	mergeOperations(ctx, logger)

	// Category 5: Filter Operations (7 functions) / 필터 작업 (7개 함수)
	filterOperations(ctx, logger)

	// Category 6: Conversion (8 functions) / 변환 (8개 함수)
	conversions(ctx, logger)

	// Category 7: Predicate Checks (7 functions) / 조건 검사 (7개 함수)
	predicates(ctx, logger)

	// Category 8: Key Operations (8 functions) / 키 작업 (8개 함수)
	keyOperations(ctx, logger)

	// Category 9: Value Operations (7 functions) / 값 작업 (7개 함수)
	valueOperations(ctx, logger)

	// Category 10: Comparison (6 functions) / 비교 (6개 함수)
	comparisons(ctx, logger)

	// Category 11: Utility Functions (NEW) / 유틸리티 함수 (신규)
	utilityFunctions(ctx, logger)

	// Advanced: Real-World Use Cases / 고급: 실제 사용 사례
	realWorldExamples(ctx, logger)

	logger.Info("")
	logger.Info("==============================================================================")
	logger.Info("✅ All examples completed successfully!")
	logger.Info("✅ 모든 예제가 성공적으로 완료되었습니다!")
	logger.Info("==============================================================================")
}

// ============================================================================
// Category 1: Basic Operations (11 functions) / 기본 작업 (11개 함수)
// ============================================================================
func basicOperations(ctx context.Context, logger *logging.Logger) {
	logger.Info("")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📦 Category 1: Basic Operations (11 functions)")
	logger.Info("📦 카테고리 1: 기본 작업 (11개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	// Sample data / 샘플 데이터
	products := map[string]int{
		"laptop":  1200,
		"mouse":   25,
		"keyword": 80,
	}
	logger.Info("📋 Sample data (product prices):", "products", products)
	logger.Info("")

	// 1. Get - Retrieve value by key / 키로 값 가져오기
	logger.Info("1️⃣  Get() - Retrieve value by key / 키로 값 가져오기")
	logger.Info("   Purpose: Safely get a value with existence check")
	logger.Info("   목적: 존재 여부 확인과 함께 안전하게 값 가져오기")
	price, exists := maputil.Get(products, "laptop")
	logger.Info("   Result:", "price", price, "exists", exists)
	logger.Info("   💡 Use case: Check if key exists before using value")
	logger.Info("")

	// 2. GetOr - Get value with default / 기본값과 함께 가져오기
	logger.Info("2️⃣  GetOr() - Get value with default fallback / 기본값으로 대체하여 가져오기")
	logger.Info("   Purpose: Get value or return default if key doesn't exist")
	logger.Info("   목적: 키가 없으면 기본값 반환")
	price = maputil.GetOr(products, "monitor", 300)
	logger.Info("   Result (non-existent key 'monitor'):", "price", price)
	logger.Info("   💡 Use case: Provide sensible defaults for missing config values")
	logger.Info("")

	// 3. Set - Add or update value / 값 추가 또는 업데이트
	logger.Info("3️⃣  Set() - Add or update a key-value pair / 키-값 쌍 추가 또는 업데이트")
	logger.Info("   Purpose: Create a new map with added/updated value (immutable)")
	logger.Info("   목적: 값이 추가/업데이트된 새 맵 생성 (불변)")
	newProducts := maputil.Set(products, "monitor", 300)
	logger.Info("   Original map:", "products", products)
	logger.Info("   New map:", "newProducts", newProducts)
	logger.Info("   💡 Use case: Immutable updates for concurrent scenarios")
	logger.Info("")

	// 4. Delete - Remove key / 키 제거
	logger.Info("4️⃣  Delete() - Remove a key from map / 맵에서 키 제거")
	logger.Info("   Purpose: Create new map without specified key")
	logger.Info("   목적: 지정된 키가 제거된 새 맵 생성")
	filtered := maputil.Delete(products, "mouse")
	logger.Info("   Result (removed 'mouse'):", "filtered", filtered)
	logger.Info("   💡 Use case: Remove deprecated configuration keys")
	logger.Info("")

	// 5. Has - Check if key exists / 키 존재 확인
	logger.Info("5️⃣  Has() - Check if key exists in map / 맵에 키가 있는지 확인")
	logger.Info("   Purpose: Boolean check for key existence")
	logger.Info("   목적: 키 존재 여부 불리언 확인")
	hasLaptop := maputil.Has(products, "laptop")
	hasMonitor := maputil.Has(products, "monitor")
	logger.Info("   Has 'laptop':", "exists", hasLaptop)
	logger.Info("   Has 'monitor':", "exists", hasMonitor)
	logger.Info("   💡 Use case: Validate required keys in configuration")
	logger.Info("")

	// 6. IsEmpty - Check if map is empty / 맵이 비어있는지 확인
	logger.Info("6️⃣  IsEmpty() - Check if map has no elements / 맵에 요소가 없는지 확인")
	logger.Info("   Purpose: Quick emptiness check")
	logger.Info("   목적: 빠른 비어있음 확인")
	empty := map[string]int{}
	logger.Info("   Empty map:", "isEmpty", maputil.IsEmpty(empty))
	logger.Info("   Products map:", "isEmpty", maputil.IsEmpty(products))
	logger.Info("   💡 Use case: Validate data before processing")
	logger.Info("")

	// 7. IsNotEmpty - Check if map has elements / 맵에 요소가 있는지 확인
	logger.Info("7️⃣  IsNotEmpty() - Check if map has elements / 맵에 요소가 있는지 확인")
	logger.Info("   Purpose: Inverse of IsEmpty for readability")
	logger.Info("   목적: 가독성을 위한 IsEmpty의 반대")
	logger.Info("   Products map:", "isNotEmpty", maputil.IsNotEmpty(products))
	logger.Info("   💡 Use case: Guard clauses in functions")
	logger.Info("")

	// 8. Len - Get map length / 맵 길이 가져오기
	logger.Info("8️⃣  Len() - Get number of elements in map / 맵의 요소 개수 가져오기")
	logger.Info("   Purpose: Count key-value pairs")
	logger.Info("   목적: 키-값 쌍 개수 세기")
	length := maputil.Len(products)
	logger.Info("   Length:", "count", length)
	logger.Info("   💡 Use case: Pagination, statistics, validation")
	logger.Info("")

	// 9. Clear - Remove all elements / 모든 요소 제거
	logger.Info("9️⃣  Clear() - Remove all elements from map / 맵의 모든 요소 제거")
	logger.Info("   Purpose: Create empty map (immutable)")
	logger.Info("   목적: 빈 맵 생성 (불변)")
	cleared := maputil.Clear(products)
	logger.Info("   Cleared map:", "cleared", cleared, "length", len(cleared))
	logger.Info("   💡 Use case: Reset state while preserving map reference")
	logger.Info("")

	// 10. Clone - Deep copy / 깊은 복사
	logger.Info("🔟 Clone() - Create deep copy of map / 맵의 깊은 복사본 생성")
	logger.Info("   Purpose: Independent copy for safe modifications")
	logger.Info("   목적: 안전한 수정을 위한 독립적인 복사본")
	cloned := maputil.Clone(products)
	logger.Info("   Cloned map:", "cloned", cloned)
	logger.Info("   Are they equal?:", "equal", maputil.Equal(products, cloned))
	logger.Info("   💡 Use case: Create snapshots, protect against mutations")
	logger.Info("")

	// 11. Equal - Compare two maps / 두 맵 비교
	logger.Info("1️⃣1️⃣ Equal() - Compare two maps for equality / 두 맵의 동등성 비교")
	logger.Info("   Purpose: Deep equality check")
	logger.Info("   목적: 깊은 동등성 확인")
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}
	map3 := map[string]int{"a": 1, "b": 3}
	logger.Info("   map1 == map2:", "equal", maputil.Equal(map1, map2))
	logger.Info("   map1 == map3:", "equal", maputil.Equal(map1, map3))
	logger.Info("   💡 Use case: Testing, validation, cache comparisons")
	logger.Info("")
}

// ============================================================================
// Category 2: Transformation (10 functions) / 변환 (10개 함수)
// ============================================================================
func transformations(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔄 Category 2: Transformation (10 functions)")
	logger.Info("🔄 카테고리 2: 변환 (10개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	scores := map[string]int{"math": 85, "english": 92, "science": 78}
	logger.Info("📋 Sample data (test scores):", "scores", scores)
	logger.Info("")

	// 1. Map - Transform to new type / 새 타입으로 변환
	logger.Info("1️⃣  Map() - Transform map values to a new type / 맵 값을 새 타입으로 변환")
	logger.Info("   Purpose: Convert map[K]V to map[K]R")
	logger.Info("   목적: map[K]V를 map[K]R로 변환")
	grades := maputil.Map(scores, func(subject string, score int) string {
		if score >= 90 {
			return "A"
		} else if score >= 80 {
			return "B"
		}
		return "C"
	})
	logger.Info("   Grades:", "grades", grades)
	logger.Info("   💡 Use case: Convert price integers to formatted strings")
	logger.Info("")

	// 2. MapKeys - Transform keys / 키 변환
	logger.Info("2️⃣  MapKeys() - Transform all keys with a function / 모든 키를 함수로 변환")
	logger.Info("   Purpose: Change key format/type")
	logger.Info("   목적: 키 형식/타입 변경")
	uppercaseKeys := maputil.MapKeys(scores, func(k string, v int) string {
		return k + "_SCORE"
	})
	logger.Info("   Uppercase keys:", "result", uppercaseKeys)
	logger.Info("   💡 Use case: Standardize key naming conventions")
	logger.Info("")

	// 3. MapValues - Transform values / 값 변환
	logger.Info("3️⃣  MapValues() - Transform all values with a function / 모든 값을 함수로 변환")
	logger.Info("   Purpose: Apply operation to all values")
	logger.Info("   목적: 모든 값에 작업 적용")
	bonusScores := maputil.MapValues(scores, func(score int) int {
		return score + 5 // +5 bonus points / +5 보너스 점수
	})
	logger.Info("   With bonus:", "bonusScores", bonusScores)
	logger.Info("   💡 Use case: Apply discounts, tax calculations")
	logger.Info("")

	// 4. MapEntries - Transform both keys and values / 키와 값 모두 변환
	logger.Info("4️⃣  MapEntries() - Transform both keys and values / 키와 값 모두 변환")
	logger.Info("   Purpose: Complete transformation of map structure")
	logger.Info("   목적: 맵 구조의 완전한 변환")
	reversed := maputil.MapEntries(scores, func(k string, v int) (int, string) {
		return v, k // Swap key-value / 키-값 교환
	})
	logger.Info("   Reversed (score -> subject):", "reversed", reversed)
	logger.Info("   💡 Use case: Create reverse lookups, indexes")
	logger.Info("")

	// 5. Invert - Swap keys and values / 키와 값 교환
	logger.Info("5️⃣  Invert() - Swap keys and values / 키와 값 교환")
	logger.Info("   Purpose: Create reverse mapping")
	logger.Info("   목적: 역방향 매핑 생성")
	inverted := maputil.Invert(scores)
	logger.Info("   Inverted:", "inverted", inverted)
	logger.Info("   💡 Use case: Bidirectional lookups")
	logger.Info("")

	// 6. Flatten - Flatten nested map / 중첩 맵 평탄화
	logger.Info("6️⃣  Flatten() - Flatten nested map structure / 중첩된 맵 구조 평탄화")
	logger.Info("   Purpose: Convert nested maps to flat structure")
	logger.Info("   목적: 중첩 맵을 평면 구조로 변환")
	nested := map[string]map[string]int{
		"class_a": {"math": 85, "english": 90},
		"class_b": {"math": 78, "english": 88},
	}
	logger.Info("   Nested data:", "nested", nested)
	flattened := maputil.Flatten(nested, ".")
	logger.Info("   Flattened:", "flattened", flattened)
	logger.Info("   💡 Use case: Configuration flattening, database denormalization")
	logger.Info("")

	// 7. Unflatten - Create nested structure / 중첩 구조 생성
	logger.Info("7️⃣  Unflatten() - Create nested map from flat keys / 평면 키로부터 중첩 맵 생성")
	logger.Info("   Purpose: Convert flat keys to nested structure")
	logger.Info("   목적: 평면 키를 중첩 구조로 변환")
	flat := map[string]int{
		"class_a.math":    85,
		"class_a.english": 90,
		"class_b.math":    78,
	}
	logger.Info("   Flat data:", "flat", flat)
	unflattened := maputil.Unflatten(flat, ".")
	logger.Info("   Unflattened:", "unflattened", unflattened)
	logger.Info("   💡 Use case: Parse dotted configuration keys")
	logger.Info("")

	// 8. Chunk - Split into smaller maps / 작은 맵으로 분할
	logger.Info("8️⃣  Chunk() - Split map into chunks of specified size / 지정된 크기의 청크로 분할")
	logger.Info("   Purpose: Batch processing")
	logger.Info("   목적: 배치 처리")
	large := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	chunks := maputil.Chunk(large, 2)
	logger.Info("   Chunks (size=2):", "count", len(chunks))
	for i, chunk := range chunks {
		logger.Info(fmt.Sprintf("   Chunk %d:", i), "data", chunk)
	}
	logger.Info("   💡 Use case: Parallel processing, rate limiting")
	logger.Info("")

	// 9. Partition - Split by predicate / 조건으로 분할
	logger.Info("9️⃣  Partition() - Split map into two by predicate / 조건으로 두 개로 분할")
	logger.Info("   Purpose: Separate passing and failing items")
	logger.Info("   목적: 통과 및 실패 항목 분리")
	passing, failing := maputil.Partition(scores, func(k string, v int) bool {
		return v >= 80
	})
	logger.Info("   Passing (>=80):", "passing", passing)
	logger.Info("   Failing (<80):", "failing", failing)
	logger.Info("   💡 Use case: Filter data into categories")
	logger.Info("")

	// 10. Compact - Remove zero values / 제로 값 제거
	logger.Info("🔟 Compact() - Remove zero values from map / 맵에서 제로 값 제거")
	logger.Info("   Purpose: Clean sparse data")
	logger.Info("   목적: 희소 데이터 정리")
	sparse := map[string]int{"a": 1, "b": 0, "c": 3, "d": 0, "e": 5}
	logger.Info("   Original (with zeros):", "sparse", sparse)
	compacted := maputil.Compact(sparse)
	logger.Info("   Compacted:", "compacted", compacted)
	logger.Info("   💡 Use case: Remove null/empty values before JSON serialization")
	logger.Info("")
}

// ============================================================================
// Category 3: Aggregation (9 functions) / 집계 (9개 함수)
// ============================================================================
func aggregations(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📊 Category 3: Aggregation (9 functions)")
	logger.Info("📊 카테고리 3: 집계 (9개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	sales := map[string]int{
		"January":  1200,
		"February": 1500,
		"March":    1100,
		"April":    1800,
	}
	logger.Info("📋 Sample data (monthly sales in $):", "sales", sales)
	logger.Info("")

	// 1. Reduce - Custom aggregation / 사용자 정의 집계
	logger.Info("1️⃣  Reduce() - Custom aggregation with accumulator / 누산기를 사용한 사용자 정의 집계")
	logger.Info("   Purpose: Flexible aggregation pattern")
	logger.Info("   목적: 유연한 집계 패턴")
	total := maputil.Reduce(sales, 0, func(acc int, month string, amount int) int {
		return acc + amount
	})
	logger.Info("   Total sales:", "total", total)
	logger.Info("   💡 Use case: Complex calculations, custom aggregations")
	logger.Info("")

	// 2. Sum - Sum all values / 모든 값 합산
	logger.Info("2️⃣  Sum() - Sum all numeric values / 모든 숫자 값 합산")
	logger.Info("   Purpose: Quick sum calculation")
	logger.Info("   목적: 빠른 합계 계산")
	totalSales := maputil.Sum(sales)
	logger.Info("   Total sales:", "sum", totalSales)
	logger.Info("   💡 Use case: Financial totals, inventory counts, statistics")
	logger.Info("")

	// 3. Min - Find minimum / 최솟값 찾기
	logger.Info("3️⃣  Min() - Find entry with minimum value / 최솟값을 가진 항목 찾기")
	logger.Info("   Purpose: Identify lowest value and its key")
	logger.Info("   목적: 최저값과 해당 키 식별")
	minMonth, minSales, _ := maputil.Min(sales)
	logger.Info("   Minimum sales month:", "month", minMonth, "sales", minSales)
	logger.Info("   💡 Use case: Find worst performer, lowest price")
	logger.Info("")

	// 4. Max - Find maximum / 최댓값 찾기
	logger.Info("4️⃣  Max() - Find entry with maximum value / 최댓값을 가진 항목 찾기")
	logger.Info("   Purpose: Identify highest value and its key")
	logger.Info("   목적: 최고값과 해당 키 식별")
	maxMonth, maxSales, _ := maputil.Max(sales)
	logger.Info("   Maximum sales month:", "month", maxMonth, "sales", maxSales)
	logger.Info("   💡 Use case: Find best performer, highest price")
	logger.Info("")

	// 5. MinBy - Find minimum by custom function / 사용자 정의 함수로 최솟값 찾기
	logger.Info("5️⃣  MinBy() - Find minimum by custom score function / 사용자 정의 점수 함수로 최솟값 찾기")
	logger.Info("   Purpose: Custom minimum logic based on score")
	logger.Info("   목적: 점수 기반 사용자 정의 최소값 로직")
	employees := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	youngest, youngestAge, _ := maputil.MinBy(employees, func(age int) float64 {
		return float64(age)
	})
	logger.Info("   Youngest employee:", "name", youngest, "age", youngestAge)
	logger.Info("   💡 Use case: Custom scoring for minimum selection")
	logger.Info("")

	// 6. MaxBy - Find maximum by custom function / 사용자 정의 함수로 최댓값 찾기
	logger.Info("6️⃣  MaxBy() - Find maximum by custom score function / 사용자 정의 점수 함수로 최댓값 찾기")
	logger.Info("   Purpose: Custom maximum logic based on score")
	logger.Info("   목적: 점수 기반 사용자 정의 최대값 로직")
	oldest, oldestAge, _ := maputil.MaxBy(employees, func(age int) float64 {
		return float64(age)
	})
	logger.Info("   Oldest employee:", "name", oldest, "age", oldestAge)
	logger.Info("   💡 Use case: Custom scoring for maximum selection")
	logger.Info("")

	// 7. Average - Calculate average / 평균 계산
	logger.Info("7️⃣  Average() - Calculate average of all values / 모든 값의 평균 계산")
	logger.Info("   Purpose: Mean value calculation")
	logger.Info("   목적: 평균값 계산")
	avgSales := maputil.Average(sales)
	logger.Info("   Average monthly sales:", "average", fmt.Sprintf("$%.2f", avgSales))
	logger.Info("   💡 Use case: Statistics, performance metrics")
	logger.Info("")

	// 8. GroupBy - Group by key function / 키 함수로 그룹화
	logger.Info("8️⃣  GroupBy() - Group slice elements by key function / 키 함수로 슬라이스 요소 그룹화")
	logger.Info("   Purpose: Create categorical groups")
	logger.Info("   목적: 범주별 그룹 생성")
	type Transaction struct {
		ID     int
		Type   string
		Amount int
	}
	transactions := []Transaction{
		{ID: 1, Type: "income", Amount: 1000},
		{ID: 2, Type: "expense", Amount: 500},
		{ID: 3, Type: "income", Amount: 1500},
		{ID: 4, Type: "expense", Amount: 300},
	}
	logger.Info("   Sample transactions:", "count", len(transactions))
	grouped := maputil.GroupBy[string, Transaction, string](transactions, func(t Transaction) string {
		return t.Type
	})
	logger.Info("   Grouped by type:")
	logger.Info("   - Income transactions:", "count", len(grouped["income"]))
	logger.Info("   - Expense transactions:", "count", len(grouped["expense"]))
	logger.Info("   💡 Use case: Data categorization, reporting")
	logger.Info("")

	// 9. CountBy - Count by key function / 키 함수로 개수 세기
	logger.Info("9️⃣  CountBy() - Count slice elements by key function / 키 함수로 슬라이스 요소 개수 세기")
	logger.Info("   Purpose: Get count for each category")
	logger.Info("   목적: 각 범주별 개수 가져오기")
	counts := maputil.CountBy[string, Transaction, string](transactions, func(t Transaction) string {
		return t.Type
	})
	logger.Info("   Transaction counts by type:", "counts", counts)
	logger.Info("   💡 Use case: Statistics, frequency analysis, histograms")
	logger.Info("")
}

// ============================================================================
// Category 4: Merge Operations (8 functions) / 병합 작업 (8개 함수)
// ============================================================================
func mergeOperations(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔗 Category 4: Merge Operations (8 functions)")
	logger.Info("🔗 카테고리 4: 병합 작업 (8개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	defaultConfig := map[string]int{"timeout": 30, "retries": 3}
	userConfig := map[string]int{"timeout": 60, "maxConn": 10}
	logger.Info("📋 Sample data:")
	logger.Info("   Default config:", "config", defaultConfig)
	logger.Info("   User config:", "config", userConfig)
	logger.Info("")

	// 1. Merge - Combine multiple maps / 여러 맵 결합
	logger.Info("1️⃣  Merge() - Combine multiple maps (last wins) / 여러 맵 결합 (마지막 우선)")
	logger.Info("   Purpose: Simple map merging")
	logger.Info("   목적: 간단한 맵 병합")
	merged := maputil.Merge(defaultConfig, userConfig)
	logger.Info("   Merged config:", "config", merged)
	logger.Info("   💡 Use case: Configuration override, settings merge")
	logger.Info("")

	// 2. MergeWith - Custom merge strategy / 사용자 정의 병합 전략
	logger.Info("2️⃣  MergeWith() - Merge with custom conflict resolver / 사용자 정의 충돌 해결로 병합")
	logger.Info("   Purpose: Control how conflicts are resolved")
	logger.Info("   목적: 충돌 해결 방법 제어")
	inventory1 := map[string]int{"apple": 10, "banana": 5}
	inventory2 := map[string]int{"apple": 15, "orange": 8}
	combined := maputil.MergeWith(func(old, new int) int {
		return old + new // Sum quantities / 수량 합산
	}, inventory1, inventory2)
	logger.Info("   Combined inventory:", "inventory", combined)
	logger.Info("   💡 Use case: Inventory management, data consolidation")
	logger.Info("")

	// 3. DeepMerge - Recursive merge / 재귀적 병합
	logger.Info("3️⃣  DeepMerge() - Recursively merge nested maps / 중첩 맵 재귀적 병합")
	logger.Info("   Purpose: Merge nested structures")
	logger.Info("   목적: 중첩 구조 병합")
	config1 := map[string]interface{}{
		"server": map[string]interface{}{"host": "localhost", "port": 8080},
		"db":     map[string]interface{}{"name": "mydb"},
	}
	config2 := map[string]interface{}{
		"server": map[string]interface{}{"port": 9090, "ssl": true},
		"cache":  map[string]interface{}{"ttl": 300},
	}
	logger.Info("   Config 1:", "config", config1)
	logger.Info("   Config 2:", "config", config2)
	deepMerged := maputil.DeepMerge(config1, config2)
	logger.Info("   Deep merged:", "config", deepMerged)
	logger.Info("   💡 Use case: Complex configuration merging")
	logger.Info("")

	// 4. Union - Combine all keys / 모든 키 결합
	logger.Info("4️⃣  Union() - Combine all maps (alias for Merge) / 모든 맵 결합 (Merge 별칭)")
	logger.Info("   Purpose: Set union operation")
	logger.Info("   목적: 집합 합집합 연산")
	set1 := map[string]int{"a": 1, "b": 2}
	set2 := map[string]int{"b": 3, "c": 4}
	union := maputil.Union(set1, set2)
	logger.Info("   Union:", "result", union)
	logger.Info("   💡 Use case: Combine datasets")
	logger.Info("")

	// 5. Intersection - Common keys only / 공통 키만
	logger.Info("5️⃣  Intersection() - Keep only common keys / 공통 키만 유지")
	logger.Info("   Purpose: Set intersection operation")
	logger.Info("   목적: 집합 교집합 연산")
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"b": 20, "c": 30, "d": 40}
	intersection := maputil.Intersection(map1, map2)
	logger.Info("   Map 1:", "map", map1)
	logger.Info("   Map 2:", "map", map2)
	logger.Info("   Intersection (common keys):", "result", intersection)
	logger.Info("   💡 Use case: Find common elements, shared permissions")
	logger.Info("")

	// 6. Difference - Keys in first but not in second / 첫 번째에만 있는 키
	logger.Info("6️⃣  Difference() - Keys in first map but not in second / 첫 번째 맵에만 있는 키")
	logger.Info("   Purpose: Set difference operation")
	logger.Info("   목적: 집합 차집합 연산")
	difference := maputil.Difference(map1, map2)
	logger.Info("   Difference (map1 - map2):", "result", difference)
	logger.Info("   💡 Use case: Find missing items, removed permissions")
	logger.Info("")

	// 7. SymmetricDifference - Keys in either but not both / 한쪽에만 있는 키
	logger.Info("7️⃣  SymmetricDifference() - Keys in either map but not both / 한 맵에만 있는 키")
	logger.Info("   Purpose: Symmetric difference operation")
	logger.Info("   목적: 대칭 차집합 연산")
	symDiff := maputil.SymmetricDifference(map1, map2)
	logger.Info("   Symmetric difference:", "result", symDiff)
	logger.Info("   💡 Use case: Find changes, detect discrepancies")
	logger.Info("")

	// 8. Assign - Mutating merge / 변경하는 병합
	logger.Info("8️⃣  Assign() - Merge into target (MUTATING!) / 대상에 병합 (변경됨!)")
	logger.Info("   Purpose: In-place merge (modifies first map)")
	logger.Info("   목적: 제자리 병합 (첫 번째 맵 수정)")
	target := map[string]int{"a": 1, "b": 2}
	source := map[string]int{"b": 3, "c": 4}
	logger.Info("   Before assign - Target:", "target", target)
	maputil.Assign(target, source)
	logger.Info("   After assign - Target:", "target", target)
	logger.Warn("   ⚠️  Warning: This function mutates the input map!")
	logger.Info("   💡 Use case: Performance-critical updates")
	logger.Info("")
}

// ============================================================================
// Category 5: Filter Operations (7 functions) / 필터 작업 (7개 함수)
// ============================================================================
func filterOperations(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔍 Category 5: Filter Operations (7 functions)")
	logger.Info("🔍 카테고리 5: 필터 작업 (7개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	products := map[string]int{
		"laptop":   1200,
		"mouse":    25,
		"keyboard": 80,
		"monitor":  300,
		"headset":  150,
	}
	logger.Info("📋 Sample data (product prices):", "products", products)
	logger.Info("")

	// 1. Filter - Filter by predicate / 조건으로 필터
	logger.Info("1️⃣  Filter() - Keep entries matching predicate / 조건에 맞는 항목만 유지")
	logger.Info("   Purpose: General purpose filtering")
	logger.Info("   목적: 범용 필터링")
	expensive := maputil.Filter(products, func(name string, price int) bool {
		return price > 100
	})
	logger.Info("   Expensive products (>$100):", "products", expensive)
	logger.Info("   💡 Use case: Price ranges, status filtering")
	logger.Info("")

	// 2. FilterKeys - Filter by key predicate / 키 조건으로 필터
	logger.Info("2️⃣  FilterKeys() - Filter by key predicate only / 키 조건으로만 필터")
	logger.Info("   Purpose: Key-based filtering")
	logger.Info("   목적: 키 기반 필터링")
	startsWithM := maputil.FilterKeys(products, func(name string) bool {
		return name[0] == 'm'
	})
	logger.Info("   Products starting with 'm':", "products", startsWithM)
	logger.Info("   💡 Use case: Name patterns, prefix matching")
	logger.Info("")

	// 3. FilterValues - Filter by value predicate / 값 조건으로 필터
	logger.Info("3️⃣  FilterValues() - Filter by value predicate only / 값 조건으로만 필터")
	logger.Info("   Purpose: Value-based filtering")
	logger.Info("   목적: 값 기반 필터링")
	affordable := maputil.FilterValues(products, func(price int) bool {
		return price <= 100
	})
	logger.Info("   Affordable products (≤$100):", "products", affordable)
	logger.Info("   💡 Use case: Threshold filtering, range queries")
	logger.Info("")

	// 4. Pick - Select specific keys / 특정 키 선택
	logger.Info("4️⃣  Pick() - Select specific keys only / 특정 키만 선택")
	logger.Info("   Purpose: Whitelist approach")
	logger.Info("   목적: 화이트리스트 방식")
	selected := maputil.Pick(products, "laptop", "monitor")
	logger.Info("   Picked (laptop, monitor):", "products", selected)
	logger.Info("   💡 Use case: Extract subset, API response shaping")
	logger.Info("")

	// 5. Omit - Exclude specific keys / 특정 키 제외
	logger.Info("5️⃣  Omit() - Exclude specific keys / 특정 키 제외")
	logger.Info("   Purpose: Blacklist approach")
	logger.Info("   목적: 블랙리스트 방식")
	filtered := maputil.Omit(products, "mouse", "keyboard")
	logger.Info("   Omitted (mouse, keyboard):", "products", filtered)
	logger.Info("   💡 Use case: Remove sensitive fields, hide internals")
	logger.Info("")

	// 6. PickBy - Pick by predicate / 조건으로 선택
	logger.Info("6️⃣  PickBy() - Pick entries matching predicate / 조건에 맞는 항목 선택")
	logger.Info("   Purpose: Dynamic whitelist")
	logger.Info("   목적: 동적 화이트리스트")
	midRange := maputil.PickBy(products, func(name string, price int) bool {
		return price >= 50 && price <= 200
	})
	logger.Info("   Mid-range products ($50-$200):", "products", midRange)
	logger.Info("   💡 Use case: Complex selection criteria")
	logger.Info("")

	// 7. OmitBy - Omit by predicate / 조건으로 제외
	logger.Info("7️⃣  OmitBy() - Omit entries matching predicate / 조건에 맞는 항목 제외")
	logger.Info("   Purpose: Dynamic blacklist")
	logger.Info("   목적: 동적 블랙리스트")
	notExpensive := maputil.OmitBy(products, func(name string, price int) bool {
		return price > 500
	})
	logger.Info("   Not expensive (≤$500):", "products", notExpensive)
	logger.Info("   💡 Use case: Exclude outliers, remove invalid data")
	logger.Info("")
}

// ============================================================================
// Category 6: Conversion (8 functions) / 변환 (8개 함수)
// ============================================================================
func conversions(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔄 Category 6: Conversion (8 functions)")
	logger.Info("🔄 카테고리 6: 변환 (8개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	config := map[string]int{"timeout": 30, "retries": 3, "maxConn": 10}
	logger.Info("📋 Sample data (configuration):", "config", config)
	logger.Info("")

	// 1. Keys - Extract all keys / 모든 키 추출
	logger.Info("1️⃣  Keys() - Extract all keys as slice / 모든 키를 슬라이스로 추출")
	logger.Info("   Purpose: Get key list")
	logger.Info("   목적: 키 목록 가져오기")
	keys := maputil.Keys(config)
	logger.Info("   Keys:", "keys", keys)
	logger.Info("   💡 Use case: Validation, iteration, display")
	logger.Info("")

	// 2. Values - Extract all values / 모든 값 추출
	logger.Info("2️⃣  Values() - Extract all values as slice / 모든 값을 슬라이스로 추출")
	logger.Info("   Purpose: Get value list")
	logger.Info("   목적: 값 목록 가져오기")
	values := maputil.Values(config)
	logger.Info("   Values:", "values", values)
	logger.Info("   💡 Use case: Statistics, data processing")
	logger.Info("")

	// 3. Entries - Convert to key-value pairs / 키-값 쌍으로 변환
	logger.Info("3️⃣  Entries() - Convert to Entry slice / Entry 슬라이스로 변환")
	logger.Info("   Purpose: Structured key-value pairs")
	logger.Info("   목적: 구조화된 키-값 쌍")
	entries := maputil.Entries(config)
	logger.Info("   Entries:", "entries", entries)
	logger.Info("   💡 Use case: Serialization, iteration with both key and value")
	logger.Info("")

	// 4. FromEntries - Build map from entries / 항목에서 맵 생성
	logger.Info("4️⃣  FromEntries() - Build map from Entry slice / Entry 슬라이스에서 맵 생성")
	logger.Info("   Purpose: Reverse of Entries()")
	logger.Info("   목적: Entries()의 역")
	reconstructed := maputil.FromEntries(entries)
	logger.Info("   Reconstructed map:", "map", reconstructed)
	logger.Info("   💡 Use case: Deserialization, map construction")
	logger.Info("")

	// 5. ToJSON - Convert to JSON string / JSON 문자열로 변환
	logger.Info("5️⃣  ToJSON() - Convert map to JSON string / 맵을 JSON 문자열로 변환")
	logger.Info("   Purpose: Serialize to JSON")
	logger.Info("   목적: JSON으로 직렬화")
	jsonStr, err := maputil.ToJSON(config)
	if err != nil {
		logger.Error("JSON conversion failed", "error", err)
	} else {
		logger.Info("   JSON:", "json", jsonStr)
	}
	logger.Info("   💡 Use case: API responses, configuration export")
	logger.Info("")

	// 6. FromJSON - Parse JSON string / JSON 문자열 파싱
	logger.Info("6️⃣  FromJSON() - Parse JSON string to map / JSON 문자열을 맵으로 파싱")
	logger.Info("   Purpose: Deserialize from JSON")
	logger.Info("   목적: JSON에서 역직렬화")
	var parsed map[string]int
	err = maputil.FromJSON(`{"timeout":60,"retries":5}`, &parsed)
	if err != nil {
		logger.Error("JSON parsing failed", "error", err)
	} else {
		logger.Info("   Parsed from JSON:", "map", parsed)
	}
	logger.Info("   💡 Use case: API requests, configuration import")
	logger.Info("")

	// 7. ToSlice - Convert to custom slice / 사용자 정의 슬라이스로 변환
	logger.Info("7️⃣  ToSlice() - Convert map to custom slice / 맵을 사용자 정의 슬라이스로 변환")
	logger.Info("   Purpose: Custom transformation to slice")
	logger.Info("   목적: 슬라이스로 사용자 정의 변환")
	formatted := maputil.ToSlice(config, func(key string, value int) string {
		return fmt.Sprintf("%s=%d", key, value)
	})
	logger.Info("   Formatted strings:", "strings", formatted)
	logger.Info("   💡 Use case: Display formatting, CSV export")
	logger.Info("")

	// 8. FromSlice - Build map from slice / 슬라이스에서 맵 생성
	logger.Info("8️⃣  FromSlice() - Build map from slice with key extractor / 키 추출 함수로 슬라이스에서 맵 생성")
	logger.Info("   Purpose: Index slice by key")
	logger.Info("   목적: 키로 슬라이스 인덱싱")
	type User struct {
		ID   int
		Name string
	}
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}
	logger.Info("   Sample users:", "count", len(users))
	userMap := maputil.FromSlice(users, func(u User) int {
		return u.ID
	})
	logger.Info("   User map (indexed by ID):", "map", userMap)
	logger.Info("   💡 Use case: Create lookups, build indexes")
	logger.Info("")
}

// ============================================================================
// Category 7: Predicate Checks (7 functions) / 조건 검사 (7개 함수)
// ============================================================================
func predicates(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("✅ Category 7: Predicate Checks (7 functions)")
	logger.Info("✅ 카테고리 7: 조건 검사 (7개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	scores := map[string]int{"math": 90, "english": 85, "science": 92}
	logger.Info("📋 Sample data (test scores):", "scores", scores)
	logger.Info("")

	// 1. Every - All match predicate / 모두 조건 충족
	logger.Info("1️⃣  Every() - Check if all entries match predicate / 모든 항목이 조건 충족 확인")
	logger.Info("   Purpose: Universal quantification")
	logger.Info("   목적: 전체 한정")
	allPassing := maputil.Every(scores, func(subject string, score int) bool {
		return score >= 80
	})
	logger.Info("   All scores >= 80?:", "result", allPassing)
	logger.Info("   💡 Use case: Validation, quality checks")
	logger.Info("")

	// 2. Some - At least one matches / 하나 이상 일치
	logger.Info("2️⃣  Some() - Check if any entry matches predicate / 어떤 항목이라도 조건 충족 확인")
	logger.Info("   Purpose: Existential quantification")
	logger.Info("   목적: 존재 한정")
	hasExcellent := maputil.Some(scores, func(subject string, score int) bool {
		return score >= 90
	})
	logger.Info("   Any score >= 90?:", "result", hasExcellent)
	logger.Info("   💡 Use case: Find if condition exists")
	logger.Info("")

	// 3. None - No entries match / 일치하는 항목 없음
	logger.Info("3️⃣  None() - Check if no entries match predicate / 조건에 맞는 항목이 없는지 확인")
	logger.Info("   Purpose: Negative existential")
	logger.Info("   목적: 부정 존재")
	noFailing := maputil.None(scores, func(subject string, score int) bool {
		return score < 60
	})
	logger.Info("   No score < 60?:", "result", noFailing)
	logger.Info("   💡 Use case: Ensure absence of bad data")
	logger.Info("")

	// 4. HasKey - Check if key exists / 키 존재 확인
	logger.Info("4️⃣  HasKey() - Check if specific key exists / 특정 키 존재 확인")
	logger.Info("   Purpose: Key membership test")
	logger.Info("   목적: 키 멤버십 테스트")
	hasMath := maputil.HasKey(scores, "math")
	hasHistory := maputil.HasKey(scores, "history")
	logger.Info("   Has 'math'?:", "exists", hasMath)
	logger.Info("   Has 'history'?:", "exists", hasHistory)
	logger.Info("   💡 Use case: Required field validation")
	logger.Info("")

	// 5. HasValue - Check if value exists / 값 존재 확인
	logger.Info("5️⃣  HasValue() - Check if specific value exists / 특정 값 존재 확인")
	logger.Info("   Purpose: Value membership test")
	logger.Info("   목적: 값 멤버십 테스트")
	has90 := maputil.HasValue(scores, 90)
	has100 := maputil.HasValue(scores, 100)
	logger.Info("   Has value 90?:", "exists", has90)
	logger.Info("   Has value 100?:", "exists", has100)
	logger.Info("   💡 Use case: Find if specific value is present")
	logger.Info("")

	// 6. HasEntry - Check if key-value pair exists / 키-값 쌍 존재 확인
	logger.Info("6️⃣  HasEntry() - Check if specific key-value pair exists / 특정 키-값 쌍 존재 확인")
	logger.Info("   Purpose: Exact entry match")
	logger.Info("   목적: 정확한 항목 일치")
	hasMath90 := maputil.HasEntry(scores, "math", 90)
	hasMath85 := maputil.HasEntry(scores, "math", 85)
	logger.Info("   Has entry ('math', 90)?:", "exists", hasMath90)
	logger.Info("   Has entry ('math', 85)?:", "exists", hasMath85)
	logger.Info("   💡 Use case: Verify specific state")
	logger.Info("")

	// 7. IsSubset - Check if subset / 부분집합 확인
	logger.Info("7️⃣  IsSubset() - Check if first map is subset of second / 첫 맵이 두 번째 맵의 부분집합인지 확인")
	logger.Info("   Purpose: Subset relationship test")
	logger.Info("   목적: 부분집합 관계 테스트")
	subset := map[string]int{"math": 90, "english": 85}
	superset := map[string]int{"math": 90, "english": 85, "science": 92}
	isSubset := maputil.IsSubset(subset, superset)
	logger.Info("   Is subset?:", "result", isSubset)
	logger.Info("   💡 Use case: Permission checks, capability testing")
	logger.Info("")
}

// ============================================================================
// Category 8: Key Operations (8 functions) / 키 작업 (8개 함수)
// ============================================================================
func keyOperations(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔑 Category 8: Key Operations (8 functions)")
	logger.Info("🔑 카테고리 8: 키 작업 (8개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	data := map[string]int{"zebra": 3, "apple": 1, "mango": 2}
	logger.Info("📋 Sample data:", "data", data)
	logger.Info("")

	// 1. KeysSorted - Get sorted keys / 정렬된 키 가져오기
	logger.Info("1️⃣  KeysSorted() - Get keys in sorted order / 정렬된 순서로 키 가져오기")
	logger.Info("   Purpose: Deterministic key ordering")
	logger.Info("   목적: 결정적 키 순서")
	sortedKeys := maputil.KeysSorted(data)
	logger.Info("   Sorted keys:", "keys", sortedKeys)
	logger.Info("   💡 Use case: Consistent output, alphabetical display")
	logger.Info("")

	// 2. FindKey - Find first matching key / 첫 번째 일치 키 찾기
	logger.Info("2️⃣  FindKey() - Find first key matching predicate / 조건에 맞는 첫 키 찾기")
	logger.Info("   Purpose: Search for key by condition")
	logger.Info("   목적: 조건으로 키 검색")
	key, found := maputil.FindKey(data, func(k string, v int) bool {
		return v > 1
	})
	logger.Info("   First key with value > 1:", "key", key, "found", found)
	logger.Info("   💡 Use case: Find first matching item")
	logger.Info("")

	// 3. FindKeys - Find all matching keys / 일치하는 모든 키 찾기
	logger.Info("3️⃣  FindKeys() - Find all keys matching predicate / 조건에 맞는 모든 키 찾기")
	logger.Info("   Purpose: Search for multiple keys")
	logger.Info("   목적: 여러 키 검색")
	keys := maputil.FindKeys(data, func(k string, v int) bool {
		return v >= 2
	})
	logger.Info("   Keys with value >= 2:", "keys", keys)
	logger.Info("   💡 Use case: Batch selection")
	logger.Info("")

	// 4. RenameKey - Rename a key / 키 이름 변경
	logger.Info("4️⃣  RenameKey() - Rename a specific key / 특정 키 이름 변경")
	logger.Info("   Purpose: Change key name while preserving value")
	logger.Info("   목적: 값을 유지하면서 키 이름 변경")
	renamed := maputil.RenameKey(data, "apple", "APPLE")
	logger.Info("   Renamed 'apple' to 'APPLE':", "result", renamed)
	logger.Info("   💡 Use case: API field mapping, normalization")
	logger.Info("")

	// 5. SwapKeys - Swap two key values / 두 키의 값 교환
	logger.Info("5️⃣  SwapKeys() - Swap values of two keys / 두 키의 값 교환")
	logger.Info("   Purpose: Exchange values between keys")
	logger.Info("   목적: 키 간 값 교환")
	swapped := maputil.SwapKeys(data, "apple", "mango")
	logger.Info("   Swapped 'apple' and 'mango':", "result", swapped)
	logger.Info("   💡 Use case: Reorder priorities, swap positions")
	logger.Info("")

	// 6. PrefixKeys - Add prefix to all keys / 모든 키에 접두사 추가
	logger.Info("6️⃣  PrefixKeys() - Add prefix to all keys / 모든 키에 접두사 추가")
	logger.Info("   Purpose: Namespace keys")
	logger.Info("   목적: 키 네임스페이스화")
	prefixed := maputil.PrefixKeys(data, "fruit_")
	logger.Info("   With prefix 'fruit_':", "result", prefixed)
	logger.Info("   💡 Use case: Avoid key collisions, categorization")
	logger.Info("")

	// 7. SuffixKeys - Add suffix to all keys / 모든 키에 접미사 추가
	logger.Info("7️⃣  SuffixKeys() - Add suffix to all keys / 모든 키에 접미사 추가")
	logger.Info("   Purpose: Add common suffix")
	logger.Info("   목적: 공통 접미사 추가")
	suffixed := maputil.SuffixKeys(data, "_count")
	logger.Info("   With suffix '_count':", "result", suffixed)
	logger.Info("   💡 Use case: Type indication, unit labeling")
	logger.Info("")

	// 8. TransformKeys - Transform all keys / 모든 키 변환
	logger.Info("8️⃣  TransformKeys() - Transform all keys with function / 함수로 모든 키 변환")
	logger.Info("   Purpose: Custom key transformation")
	logger.Info("   목적: 사용자 정의 키 변환")
	transformed := maputil.TransformKeys(data, func(k string) string {
		return fmt.Sprintf("[%s]", k)
	})
	logger.Info("   Transformed keys (brackets):", "result", transformed)
	logger.Info("   💡 Use case: Format conversion, standardization")
	logger.Info("")
}

// ============================================================================
// Category 9: Value Operations (7 functions) / 값 작업 (7개 함수)
// ============================================================================
func valueOperations(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("💎 Category 9: Value Operations (7 functions)")
	logger.Info("💎 카테고리 9: 값 작업 (7개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	prices := map[string]int{"item_a": 30, "item_b": 10, "item_c": 20, "item_d": 10}
	logger.Info("📋 Sample data (prices):", "prices", prices)
	logger.Info("")

	// 1. ValuesSorted - Get sorted values / 정렬된 값 가져오기
	logger.Info("1️⃣  ValuesSorted() - Get values in sorted order / 정렬된 순서로 값 가져오기")
	logger.Info("   Purpose: Ordered value list")
	logger.Info("   목적: 순서가 정해진 값 목록")
	sortedValues := maputil.ValuesSorted(prices)
	logger.Info("   Sorted values:", "values", sortedValues)
	logger.Info("   💡 Use case: Price sorting, ranking")
	logger.Info("")

	// 2. UniqueValues - Get unique values / 고유 값 가져오기
	logger.Info("2️⃣  UniqueValues() - Get unique values only / 고유 값만 가져오기")
	logger.Info("   Purpose: Remove duplicate values")
	logger.Info("   목적: 중복 값 제거")
	unique := maputil.UniqueValues(prices)
	logger.Info("   Unique values:", "values", unique)
	logger.Info("   💡 Use case: Find distinct values, deduplicate")
	logger.Info("")

	// 3. ReplaceValue - Replace all occurrences of a value / 값의 모든 발생 대체
	logger.Info("3️⃣  ReplaceValue() - Replace all occurrences of a value / 특정 값의 모든 발생 대체")
	logger.Info("   Purpose: Bulk value replacement")
	logger.Info("   목적: 대량 값 교체")
	replaced := maputil.ReplaceValue(prices, 10, 15)
	logger.Info("   Replaced 10 with 15:", "result", replaced)
	logger.Info("   💡 Use case: Price updates, status corrections")
	logger.Info("")

	// 4. UpdateValues - Transform all values / 모든 값 변환
	logger.Info("4️⃣  UpdateValues() - Transform all values with function / 함수로 모든 값 변환")
	logger.Info("   Purpose: Apply operation to all values")
	logger.Info("   목적: 모든 값에 작업 적용")
	discounted := maputil.UpdateValues(prices, func(k string, price int) int {
		return price * 90 / 100 // 10% discount / 10% 할인
	})
	logger.Info("   With 10% discount:", "discounted", discounted)
	logger.Info("   💡 Use case: Bulk calculations, transformations")
	logger.Info("")

	// 5. MinValue - Find minimum value / 최솟값 찾기
	logger.Info("5️⃣  MinValue() - Find minimum value in map / 맵에서 최솟값 찾기")
	logger.Info("   Purpose: Get lowest value")
	logger.Info("   목적: 최저값 가져오기")
	minPrice, found := maputil.MinValue(prices)
	if found {
		logger.Info("   Minimum price:", "price", minPrice)
	}
	logger.Info("   💡 Use case: Find lowest price, minimum threshold")
	logger.Info("")

	// 6. MaxValue - Find maximum value / 최댓값 찾기
	logger.Info("6️⃣  MaxValue() - Find maximum value in map / 맵에서 최댓값 찾기")
	logger.Info("   Purpose: Get highest value")
	logger.Info("   목적: 최고값 가져오기")
	maxPrice, found := maputil.MaxValue(prices)
	if found {
		logger.Info("   Maximum price:", "price", maxPrice)
	}
	logger.Info("   💡 Use case: Find highest price, maximum limit")
	logger.Info("")

	// 7. SumValues - Sum all values / 모든 값 합산
	logger.Info("7️⃣  SumValues() - Sum all numeric values / 모든 숫자 값 합산")
	logger.Info("   Purpose: Total calculation")
	logger.Info("   목적: 총계 계산")
	totalPrice := maputil.SumValues(prices)
	logger.Info("   Total price:", "total", totalPrice)
	logger.Info("   💡 Use case: Shopping cart total, revenue calculation")
	logger.Info("")
}

// ============================================================================
// Category 10: Comparison (6 functions) / 비교 (6개 함수)
// ============================================================================
func comparisons(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("⚖️  Category 10: Comparison (6 functions)")
	logger.Info("⚖️  카테고리 10: 비교 (6개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	oldConfig := map[string]int{"timeout": 30, "retries": 3, "workers": 5}
	newConfig := map[string]int{"timeout": 30, "retries": 5, "maxConn": 10}
	logger.Info("📋 Sample data:")
	logger.Info("   Old config:", "config", oldConfig)
	logger.Info("   New config:", "config", newConfig)
	logger.Info("")

	// 1. Diff - Find all differences / 모든 차이점 찾기
	logger.Info("1️⃣  Diff() - Find all differences between two maps / 두 맵 간 모든 차이점 찾기")
	logger.Info("   Purpose: Detect any changes")
	logger.Info("   목적: 모든 변경사항 감지")
	diff := maputil.Diff(oldConfig, newConfig)
	logger.Info("   Differences:", "diff", diff)
	logger.Info("   💡 Use case: Change detection, audit logs")
	logger.Info("")

	// 2. DiffKeys - Find keys that differ / 다른 키들 찾기
	logger.Info("2️⃣  DiffKeys() - Find keys that differ / 다른 키들 찾기")
	logger.Info("   Purpose: List of changed keys")
	logger.Info("   목적: 변경된 키 목록")
	diffKeys := maputil.DiffKeys(oldConfig, newConfig)
	logger.Info("   Different keys:", "keys", diffKeys)
	logger.Info("   💡 Use case: Track changed fields")
	logger.Info("")

	// 3. Compare - Detailed comparison / 상세 비교
	logger.Info("3️⃣  Compare() - Detailed three-way comparison / 상세한 3방향 비교")
	logger.Info("   Purpose: Categorize changes")
	logger.Info("   목적: 변경사항 분류")
	added, removed, modified := maputil.Compare(oldConfig, newConfig)
	logger.Info("   Added keys:", "added", added)
	logger.Info("   Removed keys:", "removed", removed)
	logger.Info("   Modified keys:", "modified", modified)
	logger.Info("   💡 Use case: Migration planning, version control")
	logger.Info("")

	// 4. CommonKeys - Find common keys / 공통 키 찾기
	logger.Info("4️⃣  CommonKeys() - Find keys present in all maps / 모든 맵에 존재하는 키 찾기")
	logger.Info("   Purpose: Find intersection of keys")
	logger.Info("   목적: 키의 교집합 찾기")
	thirdConfig := map[string]int{"timeout": 60, "retries": 5}
	common := maputil.CommonKeys(oldConfig, newConfig, thirdConfig)
	logger.Info("   Common keys across 3 configs:", "keys", common)
	logger.Info("   💡 Use case: Find shared fields, required keys")
	logger.Info("")

	// 5. AllKeys - Get all unique keys / 모든 고유 키 가져오기
	logger.Info("5️⃣  AllKeys() - Get all unique keys from all maps / 모든 맵의 고유 키 가져오기")
	logger.Info("   Purpose: Union of all keys")
	logger.Info("   목적: 모든 키의 합집합")
	allKeys := maputil.AllKeys(oldConfig, newConfig, thirdConfig)
	logger.Info("   All unique keys:", "keys", allKeys)
	logger.Info("   💡 Use case: Schema discovery, field collection")
	logger.Info("")

	// 6. EqualMaps - Check equality / 동등성 확인
	logger.Info("6️⃣  EqualMaps() - Check if two maps are exactly equal / 두 맵이 정확히 같은지 확인")
	logger.Info("   Purpose: Exact equality test")
	logger.Info("   목적: 정확한 동등성 테스트")
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}
	map3 := map[string]int{"a": 1, "b": 3}
	logger.Info("   map1 == map2:", "equal", maputil.EqualMaps(map1, map2))
	logger.Info("   map1 == map3:", "equal", maputil.EqualMaps(map1, map3))
	logger.Info("   💡 Use case: Testing, validation, caching")
	logger.Info("")
}

// ============================================================================
// Advanced: Real-World Use Cases / 고급: 실제 사용 사례
// ============================================================================
func realWorldExamples(ctx context.Context, logger *logging.Logger) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🌟 Advanced: Real-World Use Cases")
	logger.Info("🌟 고급: 실제 사용 사례")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	// Use Case 1: Configuration Management / 설정 관리
	logger.Info("📌 Use Case 1: Configuration Management / 설정 관리")
	logger.Info("   Scenario: Merge default, environment, and user configs")
	logger.Info("   시나리오: 기본, 환경 및 사용자 설정 병합")

	defaultCfg := map[string]int{"timeout": 30, "retries": 3, "poolSize": 10}
	envCfg := map[string]int{"timeout": 60, "debug": 1}
	userCfg := map[string]int{"poolSize": 20}

	logger.Info("   Step 1: Default config", "config", defaultCfg)
	logger.Info("   Step 2: Environment override", "config", envCfg)
	logger.Info("   Step 3: User override", "config", userCfg)

	finalCfg := maputil.Merge(defaultCfg, envCfg, userCfg)
	logger.Info("   ✅ Final configuration:", "config", finalCfg)
	logger.Info("")

	// Use Case 2: Data Validation Pipeline / 데이터 검증 파이프라인
	logger.Info("📌 Use Case 2: Data Validation Pipeline / 데이터 검증 파이프라인")
	logger.Info("   Scenario: Validate and clean user input")
	logger.Info("   시나리오: 사용자 입력 검증 및 정리")

	userInput := map[string]int{
		"age":    25,
		"salary": 50000,
		"bonus":  0,
		"tax":    -100, // Invalid negative / 잘못된 음수
	}
	logger.Info("   Raw user input:", "input", userInput)

	// Step 1: Remove zero values / 제로 값 제거
	cleaned := maputil.Compact(userInput)
	logger.Info("   Step 1 - Removed zeros:", "cleaned", cleaned)

	// Step 2: Filter out invalid values / 잘못된 값 필터
	validated := maputil.Filter(cleaned, func(k string, v int) bool {
		return v > 0
	})
	logger.Info("   Step 2 - Filtered negatives:", "validated", validated)

	// Step 3: Ensure required fields / 필수 필드 확인
	required := []string{"age", "salary"}
	hasAllRequired := true
	for _, field := range required {
		if !maputil.Has(validated, field) {
			hasAllRequired = false
			logger.Warn("   Missing required field:", "field", field)
		}
	}
	logger.Info("   ✅ Validation complete:", "hasAllRequired", hasAllRequired)
	logger.Info("")

	// Use Case 3: Shopping Cart with Discounts / 할인이 적용된 장바구니
	logger.Info("📌 Use Case 3: Shopping Cart with Discounts / 할인이 적용된 장바구니")
	logger.Info("   Scenario: Apply tiered discounts based on quantity")
	logger.Info("   시나리오: 수량 기반 단계별 할인 적용")

	cart := map[string]int{
		"laptop":   1,
		"mouse":    2,
		"keyboard": 1,
		"monitor":  1,
	}
	prices := map[string]int{
		"laptop":   1000,
		"mouse":    25,
		"keyboard": 80,
		"monitor":  300,
	}

	logger.Info("   Cart:", "cart", cart)
	logger.Info("   Prices:", "prices", prices)

	// Calculate subtotal / 소계 계산
	subtotal := 0
	for item, qty := range cart {
		if price, ok := prices[item]; ok {
			subtotal += price * qty
		}
	}
	logger.Info("   Subtotal:", "amount", subtotal)

	// Apply discounts: 10% if qty > 1 / 할인 적용: 수량 > 1이면 10%
	discountedCart := maputil.MapValues(cart, func(qty int) int {
		if qty > 1 {
			return qty * 90 / 100 // 10% off / 10% 할인
		}
		return qty
	})
	logger.Info("   After quantity discount:", "cart", discountedCart)

	// Calculate final total / 최종 합계 계산
	total := 0
	for item, qty := range discountedCart {
		if price, ok := prices[item]; ok {
			total += price * qty
		}
	}
	savings := subtotal - total
	logger.Info("   ✅ Final total:", "total", total, "saved", savings)
	logger.Info("")

	// Use Case 4: API Response Filtering / API 응답 필터링
	logger.Info("📌 Use Case 4: API Response Filtering / API 응답 필터링")
	logger.Info("   Scenario: Filter sensitive fields from API response")
	logger.Info("   시나리오: API 응답에서 민감한 필드 필터링")

	userProfile := map[string]interface{}{
		"id":         123,
		"name":       "Alice",
		"email":      "alice@example.com",
		"password":   "hashed_password",
		"ssn":        "123-45-6789",
		"created_at": "2024-01-01",
	}
	logger.Info("   Raw profile:", "profile", userProfile)

	// Remove sensitive fields / 민감한 필드 제거
	publicProfile := maputil.Omit(userProfile, "password", "ssn")
	logger.Info("   ✅ Public profile:", "profile", publicProfile)
	logger.Info("")

	// Use Case 5: Performance Monitoring / 성능 모니터링
	logger.Info("📌 Use Case 5: Performance Monitoring / 성능 모니터링")
	logger.Info("   Scenario: Analyze response times across services")
	logger.Info("   시나리오: 서비스 전체의 응답 시간 분석")

	responseTimes := map[string]int{
		"auth_service":    45,
		"user_service":    120,
		"payment_service": 250,
		"email_service":   380,
		"search_service":  95,
	}
	logger.Info("   Response times (ms):", "times", responseTimes)

	// Find slow services (> 200ms) / 느린 서비스 찾기 (> 200ms)
	slow := maputil.Filter(responseTimes, func(service string, ms int) bool {
		return ms > 200
	})
	logger.Info("   Slow services (>200ms):", "services", slow)

	// Calculate statistics / 통계 계산
	avgTime := maputil.Average(responseTimes)
	slowestService, slowestTime, _ := maputil.Max(responseTimes)
	fastestService, fastestTime, _ := maputil.Min(responseTimes)

	logger.Info("   ✅ Statistics:")
	logger.Info("   - Average:", "ms", fmt.Sprintf("%.1f", avgTime))
	logger.Info("   - Fastest:", "service", fastestService, "ms", fastestTime)
	logger.Info("   - Slowest:", "service", slowestService, "ms", slowestTime)
	logger.Info("")
}

// ============================================================================
// Category 11: Utility Functions (NEW) / 유틸리티 함수 (신규)
// ============================================================================
func utilityFunctions(ctx context.Context, logger *logging.Logger) {
	logger.Info("")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("Category 11: Utility Functions (NEW) / 유틸리티 함수 (신규)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	// 1. ForEach - Iterate over map entries / 맵 항목 순회
	logger.Info("1️⃣  ForEach() - Execute function for each entry / 각 항목에 대해 함수 실행")
	logger.Info("   Purpose: Perform side effects for each key-value pair")
	logger.Info("   목적: 각 키-값 쌍에 대해 부수 효과 수행")

	m := map[string]int{"errors": 5, "warnings": 12, "info": 100}
	logger.Info("   Input map:", "map", m)

	logger.Info("   Iterating with ForEach:")
	maputil.ForEach(m, func(level string, count int) {
		logger.Info("   - Log level:", "level", level, "count", count)
	})

	// Example: Collect all keys
	var keys []string
	maputil.ForEach(m, func(k string, v int) {
		keys = append(keys, k)
	})
	logger.Info("   ✅ Collected keys:", "keys", keys)
	logger.Info("   💡 Use case: Logging, debugging, collecting data without creating new maps")
	logger.Info("")

	// 2. GetMany - Retrieve multiple values at once / 여러 값을 한 번에 검색
	logger.Info("2️⃣  GetMany() - Get multiple values at once / 여러 값을 한 번에 가져오기")
	logger.Info("   Purpose: Batch retrieval of multiple values by keys")
	logger.Info("   목적: 키로 여러 값을 일괄 검색")

	config := map[string]string{
		"host":     "localhost",
		"port":     "5432",
		"database": "mydb",
		"username": "admin",
		"password": "secret",
	}
	logger.Info("   Input map:", "config", config)

	// Get multiple configuration values
	values := maputil.GetMany(config, "host", "port", "database", "timeout")
	logger.Info("   ✅ Retrieved values (host, port, database, timeout):", "values", values)
	logger.Info("   Note: 'timeout' doesn't exist, returns empty string (zero value)")
	logger.Info("   💡 Use case: Batch config lookups, multi-key data extraction")
	logger.Info("")

	// 3. SetMany - Set multiple key-value pairs at once / 여러 키-값 쌍을 한 번에 설정
	logger.Info("3️⃣  SetMany() - Set multiple key-value pairs at once / 여러 키-값 쌍을 한 번에 설정")
	logger.Info("   Purpose: Batch updates to map entries")
	logger.Info("   목적: 맵 항목에 대한 일괄 업데이트")

	settings := map[string]string{
		"theme": "dark",
		"lang":  "en",
	}
	logger.Info("   Input map:", "settings", settings)

	// Add multiple settings at once
	updated := maputil.SetMany(settings,
		maputil.Entry[string, string]{Key: "font", Value: "monospace"},
		maputil.Entry[string, string]{Key: "size", Value: "14"},
		maputil.Entry[string, string]{Key: "theme", Value: "light"}, // Update existing
	)
	logger.Info("   ✅ Updated map:", "updated", updated)
	logger.Info("   Note: Original map unchanged (immutable), theme value updated")
	logger.Info("   💡 Use case: Batch config updates, map initialization, merging multiple entries")
	logger.Info("")

	// 4. Tap - Execute side effect and return map / 부수 효과를 실행하고 맵 반환
	logger.Info("4️⃣  Tap() - Execute side effect and return map / 부수 효과를 실행하고 맵 반환")
	logger.Info("   Purpose: Debugging in method chains without breaking the chain")
	logger.Info("   목적: 체인을 끊지 않고 메서드 체인에서 디버깅")

	prices := map[string]int{"apple": 100, "banana": 80, "cherry": 150}
	logger.Info("   Input map:", "prices", prices)

	// Use Tap for debugging in a chain
	result := maputil.Tap(prices, func(m map[string]int) {
		logger.Info("   [Tap] Intermediate state:", "map", m)
		sum := 0
		for _, v := range m {
			sum += v
		}
		logger.Info("   [Tap] Total price:", "sum", sum)
	})

	logger.Info("   ✅ Returned map (unchanged):", "result", result)
	logger.Info("   Note: Original map passed through, side effect performed")
	logger.Info("   💡 Use case: Logging in pipelines, collecting stats, validation in chains")
	logger.Info("")

	// 5. ContainsAllKeys - Check if all keys exist / 모든 키가 존재하는지 확인
	logger.Info("5️⃣  ContainsAllKeys() - Check if all keys exist / 모든 키가 존재하는지 확인")
	logger.Info("   Purpose: Validate required keys in a map")
	logger.Info("   목적: 맵에서 필수 키 검증")

	apiResponse := map[string]interface{}{
		"status": "success",
		"data":   map[string]interface{}{"id": 123, "name": "Alice"},
		"code":   200,
	}
	logger.Info("   Input map:", "apiResponse", apiResponse)

	requiredKeys := []string{"status", "data", "code"}
	hasAll := maputil.ContainsAllKeys(apiResponse, requiredKeys)
	logger.Info("   ✅ Contains all required keys:", "hasAll", hasAll)

	missingKeys := []string{"status", "data", "timestamp"}
	hasAllMissing := maputil.ContainsAllKeys(apiResponse, missingKeys)
	logger.Info("   ❌ Contains all keys (with missing 'timestamp'):", "hasAll", hasAllMissing)

	emptyKeys := []string{}
	hasEmpty := maputil.ContainsAllKeys(apiResponse, emptyKeys)
	logger.Info("   ✅ Empty keys slice (vacuous truth):", "hasAll", hasEmpty)

	logger.Info("   💡 Use case: API response validation, required config checks, form validation")
	logger.Info("")

	// 6. Apply - Transform all values in place / 모든 값을 제자리에서 변환
	logger.Info("6️⃣  Apply() - Transform all values / 모든 값 변환")
	logger.Info("   Purpose: Apply a function to all values in the map")
	logger.Info("   목적: 맵의 모든 값에 함수 적용")

	productPrices := map[string]int{"laptop": 1000, "mouse": 20, "keyboard": 50}
	logger.Info("   Input map:", "productPrices", productPrices)

	// Apply 10% discount
	discounted := maputil.Apply(productPrices, func(k string, v int) int {
		return int(float64(v) * 0.9) // 10% discount
	})
	logger.Info("   ✅ After 10% discount:", "discounted", discounted)

	// Apply key-dependent transformation
	adjusted := maputil.Apply(productPrices, func(k string, v int) int {
		if k == "laptop" {
			return v + 100 // Add $100 to laptop
		}
		return v
	})
	logger.Info("   ✅ After key-dependent adjustment:", "adjusted", adjusted)

	logger.Info("   Note: Original map unchanged (immutable)")
	logger.Info("   💡 Use case: Bulk price adjustments, data normalization, unit conversions")
	logger.Info("")

	// 7. GetOrSet - Get value or set default / 값 가져오기 또는 기본값 설정
	logger.Info("7️⃣  GetOrSet() - Get value or set default / 값 가져오기 또는 기본값 설정")
	logger.Info("   Purpose: Ensure a key always has a value")
	logger.Info("   목적: 키가 항상 값을 가지도록 보장")

	cache := map[string]int{"a": 1, "b": 2}
	logger.Info("   Input map:", "cache", cache)

	// Get existing value
	val1 := maputil.GetOrSet(cache, "a", 10)
	logger.Info("   ✅ Get existing key 'a':", "value", val1)

	// Set and get new value
	val2 := maputil.GetOrSet(cache, "c", 10)
	logger.Info("   ✅ Get new key 'c' (sets to 10):", "value", val2)
	logger.Info("   Updated cache:", "cache", cache)

	logger.Info("   Note: Map is modified in-place, useful for lazy initialization")
	logger.Info("   💡 Use case: Cache initialization, default value management, lazy loading")
	logger.Info("")

	// 8. SetDefault - Set key only if not exists / 키가 존재하지 않을 때만 설정
	logger.Info("8️⃣  SetDefault() - Set key only if not exists / 키가 존재하지 않을 때만 설정")
	logger.Info("   Purpose: Initialize keys without overwriting")
	logger.Info("   목적: 덮어쓰지 않고 키 초기화")

	configMap := map[string]string{"host": "localhost"}
	logger.Info("   Input map:", "config", configMap)

	// Set new key
	wasSet1 := maputil.SetDefault(configMap, "port", "8080")
	logger.Info("   ✅ Set new key 'port':", "wasSet", wasSet1, "config", configMap)

	// Try to overwrite existing key (won't work)
	wasSet2 := maputil.SetDefault(configMap, "host", "0.0.0.0")
	logger.Info("   ❌ Try to overwrite 'host':", "wasSet", wasSet2, "config", configMap)

	logger.Info("   Note: Returns true if key was set, false if already existed")
	logger.Info("   💡 Use case: Safe config initialization, default value setup")
	logger.Info("")

	// 9. Defaults - Merge with default values / 기본값과 병합
	logger.Info("9️⃣  Defaults() - Merge with default values / 기본값과 병합")
	logger.Info("   Purpose: Apply default values for missing keys")
	logger.Info("   목적: 누락된 키에 대해 기본값 적용")

	userConfig := map[string]string{"host": "localhost"}
	defaultConfig := map[string]string{
		"host":    "0.0.0.0",
		"port":    "8080",
		"timeout": "30s",
	}
	logger.Info("   User config:", "userConfig", userConfig)
	logger.Info("   Default config:", "defaultConfig", defaultConfig)

	fullConfig := maputil.Defaults(userConfig, defaultConfig)
	logger.Info("   ✅ Merged config:", "fullConfig", fullConfig)
	logger.Info("   Note: User values take precedence, new map created (immutable)")

	// Empty user config case
	emptyConfig := map[string]string{}
	allDefaults := maputil.Defaults(emptyConfig, defaultConfig)
	logger.Info("   ✅ Empty config merged with defaults:", "result", allDefaults)

	logger.Info("   💡 Use case: Config management, user preferences + system defaults, template rendering")
	logger.Info("")
}
