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
	if err := os.MkdirAll("logs", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create logs directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger / 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(fmt.Sprintf("logs/sliceutil-example-%s.log", time.Now().Format("20060102-150405"))),
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
	combinatorialExamples(logger)

	// NEW: v1.7.022 functions / 신규: v1.7.022 함수들
	statisticsExamples(logger)
	diffExamples(logger)
	indexExamples(logger)
	conditionalExamples(logger)
	advancedExamples(logger)

	realWorldExamples(logger)

	logger.Info("")
	logger.Info("=== All 95 Functions Demonstrated! ===")
	logger.Info("=== 모든 95개 함수 시연 완료! ===")
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
// 9. Combinatorial Operations Examples (2 functions) / 조합 작업 예제 (2개 함수)
// ============================================================================

func combinatorialExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("9. COMBINATORIAL OPERATIONS / 조합 작업")
	logger.Info("========================================================================")

	// Permutations - Small example
	logger.Info("")
	logger.Info("Permutations Example / Permutations 예제")
	logger.Info("------------------------------------------------------------------------")

	letters := []string{"A", "B", "C"}
	perms := sliceutil.Permutations(letters)
	logger.Info("Permutations input", "letters", letters)
	logger.Info("Permutations count", "count", len(perms), "expected", "3! = 6")
	for i, perm := range perms {
		logger.Info("Permutation", "index", i+1, "value", perm)
	}

	// Permutations - Numeric example
	logger.Info("")
	numbers := []int{1, 2, 3, 4}
	numPerms := sliceutil.Permutations(numbers)
	logger.Info("Numeric permutations", "input", numbers, "count", len(numPerms), "expected", "4! = 24")
	// Show first 5 permutations only
	for i := 0; i < 5 && i < len(numPerms); i++ {
		logger.Info("Permutation", "index", i+1, "value", numPerms[i])
	}
	logger.Info("... (showing first 5 of 24 permutations)")

	// Combinations
	logger.Info("")
	logger.Info("Combinations Example / Combinations 예제")
	logger.Info("------------------------------------------------------------------------")

	items := []int{1, 2, 3, 4}
	k := 2
	combs := sliceutil.Combinations(items, k)
	logger.Info("Combinations input", "items", items, "k", k)
	logger.Info("Combinations count", "count", len(combs), "expected", "C(4,2) = 6")
	for i, comb := range combs {
		logger.Info("Combination", "index", i+1, "value", comb)
	}

	// Combinations - Another example
	logger.Info("")
	colors := []string{"Red", "Green", "Blue", "Yellow"}
	colorCombs := sliceutil.Combinations(colors, 3)
	logger.Info("Color combinations", "colors", colors, "k", 3)
	logger.Info("Color combinations count", "count", len(colorCombs), "expected", "C(4,3) = 4")
	for i, comb := range colorCombs {
		logger.Info("Color combination", "index", i+1, "value", comb)
	}

	// Performance warning example
	logger.Info("")
	logger.Info("Performance Warning / 성능 경고")
	logger.Info("------------------------------------------------------------------------")
	logger.Info("Permutations grow factorially:")
	logger.Info("  n=5  → 120 permutations")
	logger.Info("  n=10 → 3,628,800 permutations")
	logger.Info("")
	logger.Info("Combinations:")
	logger.Info("  C(10,5)  → 252 combinations")
	logger.Info("  C(20,10) → 184,756 combinations")
	logger.Info("")
	logger.Info("Use with caution for large slices!")
	logger.Info("큰 슬라이스에는 주의해서 사용하세요!")
}

// ============================================================================
// 10. Statistics Examples (8 functions) / 통계 예제 (8개 함수)
// ============================================================================

func statisticsExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("10. STATISTICS FUNCTIONS / 통계 함수")
	logger.Info("========================================================================")

	// Sample data: Test scores / 샘플 데이터: 시험 점수
	scores := []int{85, 92, 78, 95, 88, 76, 90, 85, 82, 94}
	logger.Info("Test scores dataset", "scores", scores, "count", len(scores))

	// Median - middle value / 중앙값 - 중간 값
	median, _ := sliceutil.Median(scores)
	logger.Info("Median (50th percentile)",
		"value", median,
		"explanation", "Middle value when sorted / 정렬했을 때 중간 값")

	// Mode - most frequent value / 최빈값 - 가장 자주 나타나는 값
	mode, _ := sliceutil.Mode(scores)
	logger.Info("Mode (most common score)",
		"value", mode,
		"explanation", "Most frequently occurring score / 가장 자주 나타나는 점수")

	// Frequencies - count occurrences / 빈도 - 발생 횟수
	freq := sliceutil.Frequencies(scores)
	logger.Info("Frequency distribution",
		"frequencies", freq,
		"unique_scores", len(freq))

	// Percentiles - quartiles / 백분위수 - 사분위수
	p25, _ := sliceutil.Percentile(scores, 25)
	p50, _ := sliceutil.Percentile(scores, 50)
	p75, _ := sliceutil.Percentile(scores, 75)
	p90, _ := sliceutil.Percentile(scores, 90)
	logger.Info("Percentiles (quartiles)",
		"25th", p25,
		"50th(median)", p50,
		"75th", p75,
		"90th", p90,
		"explanation", "Values below which X% of data falls / X% 데이터가 이 값 아래에 있음")

	// Standard Deviation - measure of spread / 표준 편차 - 퍼짐 정도
	stddev, _ := sliceutil.StandardDeviation(scores)
	logger.Info("Standard Deviation",
		"value", fmt.Sprintf("%.2f", stddev),
		"explanation", "Measure of data spread / 데이터 퍼짐 정도의 측정값")

	// Variance - squared standard deviation / 분산 - 표준 편차의 제곱
	variance, _ := sliceutil.Variance(scores)
	logger.Info("Variance",
		"value", fmt.Sprintf("%.2f", variance),
		"explanation", "Squared standard deviation / 표준 편차의 제곱")

	// Real-world example: Sales data analysis / 실제 사례: 판매 데이터 분석
	logger.Info("")
	logger.Info("Real-world scenario: Monthly sales analysis / 실제 시나리오: 월별 판매 분석")
	logger.Info("------------------------------------------------------------------------")

	monthlySales := []int{1200, 1500, 980, 2100, 1800, 1450, 2300, 2000, 1700, 1900, 2200, 2500}
	logger.Info("Monthly sales (units)", "data", monthlySales, "months", 12)

	avgSales := sliceutil.Average(monthlySales)
	medianSales, _ := sliceutil.Median(monthlySales)
	stddevSales, _ := sliceutil.StandardDeviation(monthlySales)

	logger.Info("Sales statistics",
		"average", fmt.Sprintf("%.0f units", avgSales),
		"median", fmt.Sprintf("%.0f units", medianSales),
		"std_dev", fmt.Sprintf("%.0f units", stddevSales),
		"interpretation", "Average±StdDev shows typical range / 평균±표준편차가 일반적 범위")

	// Most common and least common values / 최빈값과 최소빈값
	logger.Info("")
	logger.Info("Top and bottom performers analysis / 상위 및 하위 실적 분석")
	logger.Info("------------------------------------------------------------------------")

	salesFreq := []int{100, 200, 200, 300, 300, 300, 400, 400, 400, 400, 500, 600}
	logger.Info("Sales data with frequencies", "data", salesFreq)

	mostCommon := sliceutil.MostCommon(salesFreq, 2)
	logger.Info("Top 2 most common values",
		"values", mostCommon,
		"use_case", "Identify best-selling price points / 가장 잘 팔리는 가격대 파악")

	leastCommon := sliceutil.LeastCommon(salesFreq, 2)
	logger.Info("Top 2 least common values",
		"values", leastCommon,
		"use_case", "Identify underperforming segments / 저조한 세그먼트 파악")
}

// ============================================================================
// 11. Diff Examples (4 functions) / 차이 분석 예제 (4개 함수)
// ============================================================================

func diffExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("11. DIFF & COMPARISON FUNCTIONS / 차이 및 비교 함수")
	logger.Info("========================================================================")

	// Scenario: Version control / 시나리오: 버전 관리
	logger.Info("Scenario: Comparing two versions of data / 시나리오: 데이터의 두 버전 비교")
	logger.Info("------------------------------------------------------------------------")

	// Diff - simple comparison / 간단한 비교
	oldVersion := []int{1, 2, 3, 4, 5}
	newVersion := []int{2, 3, 4, 5, 6, 7}
	diff := sliceutil.Diff(oldVersion, newVersion)

	logger.Info("Version comparison",
		"old_version", oldVersion,
		"new_version", newVersion)
	logger.Info("Diff result",
		"added", diff.Added,
		"removed", diff.Removed,
		"unchanged", diff.Unchanged,
		"interpretation", "Track what changed between versions / 버전 간 변경사항 추적")

	// DiffBy - compare by key / 키로 비교
	logger.Info("")
	logger.Info("Scenario: User database synchronization / 시나리오: 사용자 데이터베이스 동기화")
	logger.Info("------------------------------------------------------------------------")

	type UserRecord struct {
		ID   int
		Name string
	}

	oldUsers := []UserRecord{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	newUsers := []UserRecord{
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie Updated"},
		{ID: 4, Name: "David"},
	}

	userDiff := sliceutil.DiffBy(oldUsers, newUsers, func(u UserRecord) int { return u.ID })

	logger.Info("User database diff",
		"old_count", len(oldUsers),
		"new_count", len(newUsers))
	logger.Info("Changes detected",
		"added_users", len(userDiff.Added),
		"removed_users", len(userDiff.Removed),
		"unchanged_users", len(userDiff.Unchanged))

	if len(userDiff.Added) > 0 {
		logger.Info("New users added", "users", userDiff.Added)
	}
	if len(userDiff.Removed) > 0 {
		logger.Info("Users removed", "users", userDiff.Removed)
	}
	logger.Info("Use case",
		"application", "Database migration, API sync, audit logs",
		"설명", "데이터베이스 마이그레이션, API 동기화, 감사 로그")

	// EqualUnordered - order-independent comparison / 순서 무관 비교
	logger.Info("")
	logger.Info("Scenario: Set equality check / 시나리오: 집합 동등성 검사")
	logger.Info("------------------------------------------------------------------------")

	listA := []string{"apple", "banana", "cherry"}
	listB := []string{"cherry", "apple", "banana"}
	listC := []string{"apple", "banana"}

	logger.Info("Equal unordered comparison",
		"listA", listA,
		"listB", listB,
		"same_content_different_order", sliceutil.EqualUnordered(listA, listB))
	logger.Info("With different content",
		"listA", listA,
		"listC", listC,
		"equal", sliceutil.EqualUnordered(listA, listC))
	logger.Info("Use case",
		"application", "Permission checks, tag matching, configuration validation",
		"설명", "권한 검사, 태그 매칭, 설정 검증")

	// HasDuplicates - duplicate detection / 중복 감지
	logger.Info("")
	logger.Info("Scenario: Data quality validation / 시나리오: 데이터 품질 검증")
	logger.Info("------------------------------------------------------------------------")

	userIDs := []int{101, 102, 103, 104, 105}
	duplicateIDs := []int{101, 102, 103, 102, 104}

	logger.Info("Checking for duplicate user IDs",
		"unique_ids", userIDs,
		"has_duplicates", sliceutil.HasDuplicates(userIDs))
	logger.Info("Checking suspicious data",
		"suspicious_ids", duplicateIDs,
		"has_duplicates", sliceutil.HasDuplicates(duplicateIDs),
		"action", "Alert and clean data / 경고하고 데이터 정리")
	logger.Info("Use case",
		"application", "Validate uniqueness constraints, detect data errors",
		"설명", "고유성 제약 검증, 데이터 오류 감지")
}

// ============================================================================
// 12. Index Examples (3 functions) / 인덱스 예제 (3개 함수)
// ============================================================================

func indexExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("12. INDEX-BASED OPERATIONS / 인덱스 기반 작업")
	logger.Info("========================================================================")

	// Scenario: Log filtering / 시나리오: 로그 필터링
	logger.Info("Scenario: Error log analysis / 시나리오: 에러 로그 분석")
	logger.Info("------------------------------------------------------------------------")

	type LogEntry struct {
		Level   string
		Message string
		Line    int
	}

	logs := []LogEntry{
		{Level: "INFO", Message: "App started", Line: 1},
		{Level: "DEBUG", Message: "Config loaded", Line: 2},
		{Level: "ERROR", Message: "Database connection failed", Line: 3},
		{Level: "INFO", Message: "Retry connection", Line: 4},
		{Level: "ERROR", Message: "Authentication error", Line: 5},
		{Level: "INFO", Message: "User logged in", Line: 6},
		{Level: "ERROR", Message: "File not found", Line: 7},
	}

	logger.Info("Total log entries", "count", len(logs))

	// FindIndices - find all error positions / 모든 에러 위치 찾기
	errorIndices := sliceutil.FindIndices(logs, func(log LogEntry) bool {
		return log.Level == "ERROR"
	})
	logger.Info("Error log positions",
		"indices", errorIndices,
		"count", len(errorIndices),
		"explanation", "Found errors at line indices / 인덱스 위치에서 에러 발견됨")

	// AtIndices - extract specific entries / 특정 항목 추출
	errorLogs := sliceutil.AtIndices(logs, errorIndices)
	logger.Info("Extracted error logs", "count", len(errorLogs))
	for i, log := range errorLogs {
		logger.Info(fmt.Sprintf("Error %d", i+1),
			"line", log.Line,
			"message", log.Message)
	}

	// RemoveIndices - clean logs / 로그 정리
	cleanedLogs := sliceutil.RemoveIndices(logs, errorIndices)
	logger.Info("Logs after removing errors",
		"original_count", len(logs),
		"cleaned_count", len(cleanedLogs),
		"removed", len(errorIndices))
	logger.Info("Use case",
		"application", "Log filtering, data extraction by position, cleanup operations",
		"설명", "로그 필터링, 위치별 데이터 추출, 정리 작업")

	// Scenario: Array manipulation by indices / 시나리오: 인덱스별 배열 조작
	logger.Info("")
	logger.Info("Scenario: Select items by position / 시나리오: 위치별 항목 선택")
	logger.Info("------------------------------------------------------------------------")

	items := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	selectedIndices := []int{0, 2, 4, 6} // Even positions

	selected := sliceutil.AtIndices(items, selectedIndices)
	logger.Info("Select by indices",
		"items", items,
		"indices", selectedIndices,
		"selected", selected,
		"pattern", "Every 2nd item / 2개마다")

	// Remove specific positions
	toRemove := []int{1, 3, 5}
	remaining := sliceutil.RemoveIndices(items, toRemove)
	logger.Info("Remove by indices",
		"items", items,
		"remove_indices", toRemove,
		"remaining", remaining)
	logger.Info("Use case",
		"application", "Pagination, sampling, batch processing",
		"설명", "페이지네이션, 샘플링, 배치 처리")
}

// ============================================================================
// 13. Conditional Examples (3 functions) / 조건부 예제 (3개 함수)
// ============================================================================

func conditionalExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("13. CONDITIONAL OPERATIONS / 조건부 작업")
	logger.Info("========================================================================")

	// Scenario: Data sanitization / 시나리오: 데이터 정제
	logger.Info("Scenario: Replace invalid values / 시나리오: 잘못된 값 교체")
	logger.Info("------------------------------------------------------------------------")

	temperatures := []int{22, 25, -999, 28, 30, -999, 26, 24}
	logger.Info("Temperature data (with invalid -999 values)",
		"data", temperatures)

	// ReplaceIf - replace all invalid values / 모든 잘못된 값 교체
	cleanedTemp := sliceutil.ReplaceIf(temperatures,
		func(t int) bool { return t == -999 },
		0)
	logger.Info("After replacing invalid values",
		"cleaned", cleanedTemp,
		"replacement", 0,
		"explanation", "Replace sensor errors with 0 / 센서 오류를 0으로 교체")

	// ReplaceAll - simple replacement / 간단한 교체
	logger.Info("")
	logger.Info("Scenario: Status code normalization / 시나리오: 상태 코드 정규화")
	logger.Info("------------------------------------------------------------------------")

	statusCodes := []int{200, 201, 200, 404, 200, 500, 201}
	logger.Info("HTTP status codes", "codes", statusCodes)

	normalized := sliceutil.ReplaceAll(statusCodes, 201, 200)
	logger.Info("Normalize 201 to 200",
		"before", statusCodes,
		"after", normalized,
		"reason", "Treat 200 and 201 as success / 200과 201을 성공으로 처리")

	// UpdateWhere - complex transformation / 복잡한 변환
	logger.Info("")
	logger.Info("Scenario: Dynamic price adjustment / 시나리오: 동적 가격 조정")
	logger.Info("------------------------------------------------------------------------")

	type Product struct {
		Name     string
		Price    float64
		Category string
	}

	products := []Product{
		{Name: "Laptop", Price: 1000, Category: "Electronics"},
		{Name: "Mouse", Price: 25, Category: "Electronics"},
		{Name: "Desk", Price: 300, Category: "Furniture"},
		{Name: "Chair", Price: 150, Category: "Furniture"},
	}

	logger.Info("Original products", "count", len(products))
	for _, p := range products {
		logger.Info("Product", "name", p.Name, "price", p.Price, "category", p.Category)
	}

	// Apply 20% discount to Electronics
	discounted := sliceutil.UpdateWhere(products,
		func(p Product) bool { return p.Category == "Electronics" },
		func(p Product) Product {
			p.Price = p.Price * 0.8 // 20% discount
			return p
		})

	logger.Info("After applying 20% discount to Electronics")
	for _, p := range discounted {
		logger.Info("Product", "name", p.Name, "price", fmt.Sprintf("%.2f", p.Price), "category", p.Category)
	}
	logger.Info("Use case",
		"application", "Bulk updates, conditional transformations, business rules",
		"설명", "대량 업데이트, 조건부 변환, 비즈니스 규칙")

	// Real-world example: User activation / 실제 예: 사용자 활성화
	logger.Info("")
	logger.Info("Scenario: Bulk user activation / 시나리오: 대량 사용자 활성화")
	logger.Info("------------------------------------------------------------------------")

	type User struct {
		ID       int
		Email    string
		IsActive bool
	}

	users := []User{
		{ID: 1, Email: "alice@example.com", IsActive: false},
		{ID: 2, Email: "bob@example.com", IsActive: true},
		{ID: 3, Email: "charlie@example.com", IsActive: false},
	}

	activated := sliceutil.UpdateWhere(users,
		func(u User) bool { return !u.IsActive },
		func(u User) User {
			u.IsActive = true
			return u
		})

	logger.Info("Bulk activation result",
		"total_users", len(users),
		"activated_count", sliceutil.Count(activated, func(u User) bool { return u.IsActive }),
		"action", "All inactive users are now active / 모든 비활성 사용자가 활성화됨")
}

// ============================================================================
// 14. Advanced Examples (4 functions) / 고급 예제 (4개 함수)
// ============================================================================

func advancedExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("14. ADVANCED FUNCTIONAL PROGRAMMING / 고급 함수형 프로그래밍")
	logger.Info("========================================================================")

	// Scan - cumulative operations / 누적 연산
	logger.Info("Scenario: Running totals and cumulative calculations / 시나리오: 누적 합계 및 누적 계산")
	logger.Info("------------------------------------------------------------------------")

	dailySales := []int{100, 150, 200, 120, 180}
	logger.Info("Daily sales", "data", dailySales)

	cumulativeSales := sliceutil.Scan(dailySales, 0, func(acc, n int) int {
		return acc + n
	})
	logger.Info("Cumulative sales (running total)",
		"result", cumulativeSales,
		"explanation", "Each value shows total up to that day / 각 값은 그날까지의 총합")

	// Running maximum
	maxSoFar := sliceutil.Scan(dailySales, 0, func(acc, n int) int {
		if n > acc {
			return n
		}
		return acc
	})
	logger.Info("Running maximum",
		"result", maxSoFar,
		"explanation", "Highest sale seen up to each day / 각 날짜까지의 최고 판매량")
	logger.Info("Use case",
		"application", "Running totals, progressive calculations, streak tracking",
		"설명", "누적 합계, 점진적 계산, 연속 기록 추적")

	// ZipWith - combine two sequences / 두 시퀀스 결합
	logger.Info("")
	logger.Info("Scenario: Combine related data sources / 시나리오: 관련 데이터 소스 결합")
	logger.Info("------------------------------------------------------------------------")

	months := []string{"Jan", "Feb", "Mar", "Apr"}
	revenue := []int{50000, 55000, 60000, 58000}

	logger.Info("Data sources", "months", months, "revenue", revenue)

	report := sliceutil.ZipWith(months, revenue, func(month string, rev int) string {
		return fmt.Sprintf("%s: $%d", month, rev)
	})
	logger.Info("Monthly revenue report", "report", report)

	// Price calculation example
	quantities := []int{10, 20, 15, 5}
	unitPrices := []float64{99.99, 49.99, 29.99, 199.99}

	totals := sliceutil.ZipWith(quantities, unitPrices, func(qty int, price float64) float64 {
		return float64(qty) * price
	})
	logger.Info("Order line totals",
		"quantities", quantities,
		"unit_prices", unitPrices,
		"line_totals", totals)
	logger.Info("Use case",
		"application", "Merging parallel arrays, calculations, report generation",
		"설명", "병렬 배열 병합, 계산, 리포트 생성")

	// RotateLeft - circular shift / 순환 이동
	logger.Info("")
	logger.Info("Scenario: Shift scheduling / 시나리오: 교대 근무 스케줄링")
	logger.Info("------------------------------------------------------------------------")

	schedule := []string{"Alice", "Bob", "Charlie", "David"}
	logger.Info("Original schedule", "week1", schedule)

	week2 := sliceutil.RotateLeft(schedule, 1)
	week3 := sliceutil.RotateLeft(schedule, 2)
	week4 := sliceutil.RotateLeft(schedule, 3)

	logger.Info("Rotating schedule",
		"week1", schedule,
		"week2", week2,
		"week3", week3,
		"week4", week4,
		"explanation", "Fair rotation ensures everyone gets all shifts / 공정한 순환으로 모두가 모든 교대 근무")
	logger.Info("Use case",
		"application", "Rotating schedules, circular buffers, carousel displays",
		"설명", "순환 스케줄, 순환 버퍼, 회전 디스플레이")

	// RotateRight - reverse circular shift / 역방향 순환 이동
	logger.Info("")
	logger.Info("Scenario: Undo rotation / 시나리오: 회전 취소")
	logger.Info("------------------------------------------------------------------------")

	original := []int{1, 2, 3, 4, 5}
	rotated := sliceutil.RotateLeft(original, 2)
	restored := sliceutil.RotateRight(rotated, 2)

	logger.Info("Rotation and restoration",
		"original", original,
		"rotated_left_2", rotated,
		"rotated_right_2", restored,
		"match", sliceutil.Equal(original, restored))
	logger.Info("Use case",
		"application", "Undo operations, state restoration, algorithm implementations",
		"설명", "작업 취소, 상태 복원, 알고리즘 구현")

	// Complex example: Time series analysis / 복잡한 예: 시계열 분석
	logger.Info("")
	logger.Info("Complex scenario: Moving average calculation / 복잡한 시나리오: 이동 평균 계산")
	logger.Info("------------------------------------------------------------------------")

	stockPrices := []float64{100, 102, 101, 105, 103, 107, 110, 108}
	logger.Info("Stock prices", "data", stockPrices)

	// Calculate 3-day moving average using Scan and ZipWith
	logger.Info("Calculating 3-day moving average",
		"method", "Scan cumulative sums + ZipWith for differences",
		"window", 3,
		"설명", "누적 합계 Scan + 차이 계산 ZipWith")

	logger.Info("Result interpretation",
		"insight", "Moving averages smooth out price fluctuations",
		"통찰", "이동 평균은 가격 변동을 부드럽게 함")
}

// ============================================================================
// 15. Real-World Scenarios / 실제 시나리오
// ============================================================================

func realWorldExamples(logger *logging.Logger) {
	logger.Info("")
	logger.Info("10. REAL-WORLD SCENARIOS / 실제 사용 시나리오")
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
