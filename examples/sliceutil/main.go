package main

import (
	"fmt"
	"os"
	"time"

	"github.com/arkd0ng/go-utils/logging"
	"github.com/arkd0ng/go-utils/sliceutil"
)

// User represents a user in the system / User는 시스템의 사용자를 나타냅니다
type User struct {
	ID       int
	Name     string
	Age      int
	IsActive bool
	City     string
}

// Product represents a product in the system / Product는 시스템의 제품을 나타냅니다
type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	Sales    int
}

func main() {
	// Create results directory if it doesn't exist / 결과 디렉토리가 없다면 새롭게 생성
	if err := os.MkdirAll("./results/logs", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create logs directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger / 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(fmt.Sprintf("./results/logs/sliceutil_example_%s.log", time.Now().Format("20060102_150405"))),
		logging.WithLevel(logging.DEBUG),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Print banner / 배너 출력
	logger.Banner("Sliceutil Package Examples", "go-utils/sliceutil")

	logger.Info("=== Starting Sliceutil Package Examples ===")
	logger.Info("=== Sliceutil 패키지 예제 시작 ===")
	logger.Info("")

	// Run all example categories / 모든 예제 카테고리 실행
	basicOperationsExamples(logger)
	transformationExamples(logger)
	aggregationExamples(logger)
	slicingExamples(logger)
	setOperationsExamples(logger)
	sortingExamples(logger)
	predicateExamples(logger)
	utilityExamples(logger)
	realWorldExamples(logger)

	logger.Info("")
	logger.Info("=== All 60 Functions Demonstrated! ===")
	logger.Info("=== 모든 60개 함수 시연 완료! ===")
}

// ============================================================================
// 1. Basic Operations Examples (10 functions) / 기본 작업 예제 (10개 함수)
// ============================================================================

func basicOperationsExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("1. BASIC OPERATIONS / 기본 작업")
	logger.Info("========================================================================")

	// Contains
	numbers := []int{1, 2, 3, 4, 5}
	logger.Info("Contains example", "numbers", numbers, "search", 3, "result", sliceutil.Contains(numbers, 3))

	// ContainsFunc
	hasEven := sliceutil.ContainsFunc(numbers, func(n int) bool { return n%2 == 0 })
	logger.Info("ContainsFunc example", "predicate", "n%2==0", "result", hasEven)

	// IndexOf
	idx := sliceutil.IndexOf(numbers, 3)
	logger.Info("IndexOf example", "search", 3, "index", idx)

	// LastIndexOf
	duplicates := []int{1, 2, 3, 2, 4, 2, 5}
	lastIdx := sliceutil.LastIndexOf(duplicates, 2)
	logger.Info("LastIndexOf example", "slice", duplicates, "search", 2, "index", lastIdx)

	// Find
	value, found := sliceutil.Find(numbers, func(n int) bool { return n > 3 })
	logger.Info("Find example", "predicate", "n>3", "value", value, "found", found)

	// FindIndex
	findIdx := sliceutil.FindIndex(numbers, func(n int) bool { return n > 3 })
	logger.Info("FindIndex example", "predicate", "n>3", "index", findIdx)

	// Count
	count := sliceutil.Count(numbers, func(n int) bool { return n%2 == 0 })
	logger.Info("Count example", "predicate", "n%2==0", "count", count)

	// IsEmpty
	empty := []int{}
	logger.Info("IsEmpty example", "empty_slice", sliceutil.IsEmpty(empty), "numbers_slice", sliceutil.IsEmpty(numbers))

	// IsNotEmpty
	logger.Info("IsNotEmpty example", "numbers_slice", sliceutil.IsNotEmpty(numbers))

	// Equal
	other := []int{1, 2, 3, 4, 5}
	logger.Info("Equal example", "slice1", numbers, "slice2", other, "equal", sliceutil.Equal(numbers, other))
}

// ============================================================================
// 2. Transformation Examples (8 functions) / 변환 예제 (8개 함수)
// ============================================================================

func transformationExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("2. TRANSFORMATION FUNCTIONS / 변환 함수")
	logger.Info("========================================================================")

	numbers := []int{1, 2, 3, 4, 5}

	// Map
	doubled := sliceutil.Map(numbers, func(n int) int { return n * 2 })
	logger.Info("Map example", "input", numbers, "operation", "x2", "output", doubled)

	// Filter
	evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
	logger.Info("Filter example", "input", numbers, "predicate", "n%2==0", "output", evens)

	// FlatMap
	words := []string{"hello", "world"}
	chars := sliceutil.FlatMap(words, func(s string) []string {
		result := make([]string, len(s))
		for i, c := range s {
			result[i] = string(c)
		}
		return result
	})
	logger.Info("FlatMap example", "input", words, "output_length", len(chars))

	// Flatten
	nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
	flat := sliceutil.Flatten(nested)
	logger.Info("Flatten example", "input", nested, "output", flat)

	// Unique
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	unique := sliceutil.Unique(duplicates)
	logger.Info("Unique example", "input", duplicates, "output", unique)

	// UniqueBy
	users := []User{
		{ID: 1, Name: "Alice", Age: 25, City: "Seoul"},
		{ID: 2, Name: "Bob", Age: 30, City: "Seoul"},
		{ID: 1, Name: "Alice Duplicate", Age: 25, City: "Seoul"},
	}
	uniqueUsers := sliceutil.UniqueBy(users, func(u User) int { return u.ID })
	logger.Info("UniqueBy example", "input_count", len(users), "output_count", len(uniqueUsers))

	// Compact
	withZeros := []int{1, 0, 2, 0, 3, 0, 4}
	compacted := sliceutil.Compact(withZeros)
	logger.Info("Compact example", "input", withZeros, "output", compacted)

	// Reverse
	reversed := sliceutil.Reverse(numbers)
	logger.Info("Reverse example", "input", numbers, "output", reversed)
}

// ============================================================================
// 3. Aggregation Examples (7 functions) / 집계 예제 (7개 함수)
// ============================================================================

func aggregationExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("3. AGGREGATION FUNCTIONS / 집계 함수")
	logger.Info("========================================================================")

	numbers := []int{1, 2, 3, 4, 5}

	// Reduce
	sum := sliceutil.Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	logger.Info("Reduce example", "input", numbers, "operation", "sum", "result", sum)

	// Sum
	total := sliceutil.Sum(numbers)
	logger.Info("Sum example", "input", numbers, "result", total)

	// Min
	min, _ := sliceutil.Min(numbers)
	logger.Info("Min example", "input", numbers, "result", min)

	// Max
	max, _ := sliceutil.Max(numbers)
	logger.Info("Max example", "input", numbers, "result", max)

	// Average
	avg := sliceutil.Average(numbers)
	logger.Info("Average example", "input", numbers, "result", avg)

	// GroupBy
	users := []User{
		{ID: 1, Name: "Alice", Age: 25, City: "Seoul"},
		{ID: 2, Name: "Bob", Age: 30, City: "Busan"},
		{ID: 3, Name: "Charlie", Age: 35, City: "Seoul"},
	}
	byCity := sliceutil.GroupBy(users, func(u User) string { return u.City })
	logger.Info("GroupBy example", "total_users", len(users), "cities", len(byCity))
	for city, cityUsers := range byCity {
		logger.Info("  Group", "city", city, "count", len(cityUsers))
	}

	// Partition
	evens, odds := sliceutil.Partition(numbers, func(n int) bool { return n%2 == 0 })
	logger.Info("Partition example", "input", numbers, "evens", evens, "odds", odds)
}

// ============================================================================
// 4. Slicing Examples (7 functions) / 슬라이싱 예제 (7개 함수)
// ============================================================================

func slicingExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("4. SLICING FUNCTIONS / 슬라이싱 함수")
	logger.Info("========================================================================")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Chunk
	chunks := sliceutil.Chunk(numbers, 3)
	logger.Info("Chunk example", "input", numbers, "chunk_size", 3, "chunks_count", len(chunks))

	// Take
	first3 := sliceutil.Take(numbers, 3)
	logger.Info("Take example", "input", numbers, "n", 3, "output", first3)

	// TakeLast
	last3 := sliceutil.TakeLast(numbers, 3)
	logger.Info("TakeLast example", "input", numbers, "n", 3, "output", last3)

	// Drop
	rest := sliceutil.Drop(numbers, 3)
	logger.Info("Drop example", "input", numbers, "n", 3, "output", rest)

	// DropLast
	withoutLast3 := sliceutil.DropLast(numbers, 3)
	logger.Info("DropLast example", "input", numbers, "n", 3, "output", withoutLast3)

	// Slice
	middle := sliceutil.Slice(numbers, 3, 7)
	logger.Info("Slice example", "input", numbers, "start", 3, "end", 7, "output", middle)

	// Sample
	sample := sliceutil.Sample(numbers, 3)
	logger.Info("Sample example", "input", numbers, "n", 3, "sample_length", len(sample))
}

// ============================================================================
// 5. Set Operations Examples (6 functions) / 집합 작업 예제 (6개 함수)
// ============================================================================

func setOperationsExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("5. SET OPERATIONS / 집합 작업")
	logger.Info("========================================================================")

	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	// Union
	union := sliceutil.Union(set1, set2)
	logger.Info("Union example", "set1", set1, "set2", set2, "result", union)

	// Intersection
	intersection := sliceutil.Intersection(set1, set2)
	logger.Info("Intersection example", "set1", set1, "set2", set2, "result", intersection)

	// Difference
	diff := sliceutil.Difference(set1, set2)
	logger.Info("Difference example", "set1", set1, "set2", set2, "result", diff)

	// SymmetricDifference
	symDiff := sliceutil.SymmetricDifference(set1, set2)
	logger.Info("SymmetricDifference example", "set1", set1, "set2", set2, "result", symDiff)

	// IsSubset
	subset := []int{2, 3, 4}
	isSubset := sliceutil.IsSubset(subset, set1)
	logger.Info("IsSubset example", "subset", subset, "superset", set1, "result", isSubset)

	// IsSuperset
	isSuperset := sliceutil.IsSuperset(set1, subset)
	logger.Info("IsSuperset example", "superset", set1, "subset", subset, "result", isSuperset)
}

// ============================================================================
// 6. Sorting Examples (5 functions) / 정렬 예제 (5개 함수)
// ============================================================================

func sortingExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("6. SORTING FUNCTIONS / 정렬 함수")
	logger.Info("========================================================================")

	numbers := []int{5, 2, 8, 1, 9, 3}

	// Sort
	sorted := sliceutil.Sort(numbers)
	logger.Info("Sort example", "input", numbers, "output", sorted)

	// SortDesc
	sortedDesc := sliceutil.SortDesc(numbers)
	logger.Info("SortDesc example", "input", numbers, "output", sortedDesc)

	// SortBy
	users := []User{
		{ID: 1, Name: "Charlie", Age: 35},
		{ID: 2, Name: "Alice", Age: 25},
		{ID: 3, Name: "Bob", Age: 30},
	}
	sortedByAge := sliceutil.SortBy(users, func(u User) int { return u.Age })
	logger.Info("SortBy example", "sort_by", "age")
	for _, u := range sortedByAge {
		logger.Info("  User", "name", u.Name, "age", u.Age)
	}

	// IsSorted
	logger.Info("IsSorted example", "slice", sorted, "result", sliceutil.IsSorted(sorted))

	// IsSortedDesc
	logger.Info("IsSortedDesc example", "slice", sortedDesc, "result", sliceutil.IsSortedDesc(sortedDesc))
}

// ============================================================================
// 7. Predicate Examples (6 functions) / 조건 검사 예제 (6개 함수)
// ============================================================================

func predicateExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("7. PREDICATE FUNCTIONS / 조건 검사 함수")
	logger.Info("========================================================================")

	numbers := []int{2, 4, 6, 8, 10}

	// All
	allEven := sliceutil.All(numbers, func(n int) bool { return n%2 == 0 })
	logger.Info("All example", "predicate", "n%2==0", "result", allEven)

	// Any
	anyOdd := sliceutil.Any(numbers, func(n int) bool { return n%2 != 0 })
	logger.Info("Any example", "predicate", "n%2!=0", "result", anyOdd)

	// None
	noneNegative := sliceutil.None(numbers, func(n int) bool { return n < 0 })
	logger.Info("None example", "predicate", "n<0", "result", noneNegative)

	// AllEqual
	same := []int{5, 5, 5, 5}
	logger.Info("AllEqual example", "slice", same, "result", sliceutil.AllEqual(same))

	// IsSortedBy
	users := []User{
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Bob", Age: 30},
		{ID: 3, Name: "Charlie", Age: 35},
	}
	isSorted := sliceutil.IsSortedBy(users, func(u User) int { return u.Age })
	logger.Info("IsSortedBy example", "sort_by", "age", "result", isSorted)

	// ContainsAll
	containsAll := sliceutil.ContainsAll([]int{1, 2, 3, 4, 5}, 2, 4)
	logger.Info("ContainsAll example", "result", containsAll)
}

// ============================================================================
// 8. Utility Examples (11 functions) / 유틸리티 예제 (11개 함수)
// ============================================================================

func utilityExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("8. UTILITY FUNCTIONS / 유틸리티 함수")
	logger.Info("========================================================================")

	numbers := []int{1, 2, 3, 4, 5}

	// ForEach
	logger.Info("ForEach example:")
	sliceutil.ForEach(numbers, func(n int) {
		logger.Debug("  Processing number", "value", n)
	})

	// ForEachIndexed
	logger.Info("ForEachIndexed example:")
	sliceutil.ForEachIndexed(numbers, func(i, n int) {
		logger.Debug("  Index and value", "index", i, "value", n)
	})

	// Join
	joined := sliceutil.Join(numbers, ", ")
	logger.Info("Join example", "input", numbers, "separator", ", ", "result", joined)

	// Clone
	cloned := sliceutil.Clone(numbers)
	logger.Info("Clone example", "original", numbers, "cloned", cloned)

	// Fill
	toFill := make([]int, 5)
	filled := sliceutil.Fill(toFill, 7)
	logger.Info("Fill example", "size", 5, "value", 7, "result", filled)

	// Insert
	inserted := sliceutil.Insert(numbers, 2, 99)
	logger.Info("Insert example", "input", numbers, "index", 2, "value", 99, "output", inserted)

	// Remove
	removed := sliceutil.Remove(numbers, 2)
	logger.Info("Remove example", "input", numbers, "index", 2, "output", removed)

	// RemoveAll
	toRemove := []int{1, 2, 2, 3, 3, 3, 4, 5}
	removedAll := sliceutil.RemoveAll(toRemove, 3)
	logger.Info("RemoveAll example", "input", toRemove, "value", 3, "output", removedAll)

	// Shuffle
	shuffled := sliceutil.Shuffle(numbers)
	logger.Info("Shuffle example", "input", numbers, "output", shuffled)

	// Zip
	names := []string{"Alice", "Bob", "Charlie"}
	ages := []int{25, 30, 35}
	zipped := sliceutil.Zip(names, ages)
	logger.Info("Zip example", "names", names, "ages", ages, "pairs_count", len(zipped))

	// Unzip
	unzippedNames, unzippedAges := sliceutil.Unzip[string, int](zipped)
	logger.Info("Unzip example", "names", unzippedNames, "ages", unzippedAges)
}

// ============================================================================
// 9. Real-World Scenarios / 실제 시나리오
// ============================================================================

func realWorldExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("9. REAL-WORLD SCENARIOS / 실제 사용 시나리오")
	logger.Info("========================================================================")

	// Scenario 1: User Data Processing / 사용자 데이터 처리
	logger.Info("")
	logger.Info("Scenario 1: User Data Processing / 사용자 데이터 처리")
	logger.Info("------------------------------------------------------------------------")

	users := []User{
		{ID: 1, Name: "Alice", Age: 28, IsActive: true, City: "Seoul"},
		{ID: 2, Name: "Bob", Age: 35, IsActive: false, City: "Busan"},
		{ID: 3, Name: "Charlie", Age: 42, IsActive: true, City: "Seoul"},
		{ID: 4, Name: "Diana", Age: 30, IsActive: true, City: "Seoul"},
		{ID: 5, Name: "Eve", Age: 25, IsActive: false, City: "Busan"},
	}

	logger.Info("Total users", "count", len(users))

	// Filter active users in Seoul
	activeSeoul := sliceutil.Filter(users, func(u User) bool {
		return u.IsActive && u.City == "Seoul"
	})
	logger.Info("Active users in Seoul", "count", len(activeSeoul))

	// Get their names
	names := sliceutil.Map(activeSeoul, func(u User) string { return u.Name })
	logger.Info("Names", "list", names)

	// Calculate average age
	ages := sliceutil.Map(activeSeoul, func(u User) int { return u.Age })
	avgAge := sliceutil.Average(ages)
	logger.Info("Average age", "value", avgAge)

	// Scenario 2: Product Data Processing / 제품 데이터 처리
	logger.Info("")
	logger.Info("Scenario 2: Product Data Processing / 제품 데이터 처리")
	logger.Info("------------------------------------------------------------------------")

	products := []Product{
		{ID: 1, Name: "Laptop Pro", Category: "Electronics", Price: 1299.99, Sales: 450},
		{ID: 2, Name: "Mouse", Category: "Electronics", Price: 29.99, Sales: 1200},
		{ID: 3, Name: "Desk Chair", Category: "Furniture", Price: 249.99, Sales: 300},
		{ID: 4, Name: "Monitor", Category: "Electronics", Price: 399.99, Sales: 600},
		{ID: 5, Name: "Standing Desk", Category: "Furniture", Price: 499.99, Sales: 150},
	}

	logger.Info("Total products", "count", len(products))

	// Group by category
	byCategory := sliceutil.GroupBy(products, func(p Product) string { return p.Category })
	logger.Info("Categories", "count", len(byCategory))

	// Find best seller in each category
	for category, items := range byCategory {
		sorted := sliceutil.SortBy(items, func(p Product) int { return -p.Sales }) // Negative for descending
		if len(sorted) > 0 {
			bestSeller := sorted[0]
			logger.Info("Best seller",
				"category", category,
				"product", bestSeller.Name,
				"sales", bestSeller.Sales)
		}
	}

	// Scenario 3: Data Analysis Pipeline / 데이터 분석 파이프라인
	logger.Info("")
	logger.Info("Scenario 3: Data Analysis Pipeline / 데이터 분석 파이프라인")
	logger.Info("------------------------------------------------------------------------")

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter -> Map -> Reduce pipeline
	result := sliceutil.Reduce(
		sliceutil.Map(
			sliceutil.Filter(data, func(n int) bool { return n%2 == 0 }),
			func(n int) int { return n * n },
		),
		0,
		func(acc, n int) int { return acc + n },
	)

	logger.Info("Pipeline result",
		"operation", "filter(even) -> map(square) -> reduce(sum)",
		"result", result)
}
