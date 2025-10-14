package main

import (
	"fmt"

	"github.com/arkd0ng/go-utils/sliceutil"
)

func main() {
	fmt.Println("=== Sliceutil Package Examples / Sliceutil 패키지 예제 ===")
	fmt.Println()

	// Run all example categories / 모든 예제 카테고리 실행
	basicOperationsExamples()
	transformationExamples()
	aggregationExamples()
	slicingExamples()
	setOperationsExamples()
	sortingExamples()
	predicateExamples()
	utilityExamples()
	realWorldExamples()
}

// ============================================================================
// 1. Basic Operations Examples (10 functions) / 기본 작업 예제 (10개 함수)
// ============================================================================

func basicOperationsExamples() {
	fmt.Println("1. BASIC OPERATIONS / 기본 작업")
	fmt.Println("=" + string(make([]byte, 70)))

	// Contains
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Contains 3: %v\n", sliceutil.Contains(numbers, 3))

	// ContainsFunc
	hasEven := sliceutil.ContainsFunc(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Has even number: %v\n", hasEven)

	// IndexOf
	idx := sliceutil.IndexOf(numbers, 3)
	fmt.Printf("Index of 3: %d\n", idx)

	// LastIndexOf
	duplicates := []int{1, 2, 3, 2, 4, 2, 5}
	lastIdx := sliceutil.LastIndexOf(duplicates, 2)
	fmt.Printf("Last index of 2: %d\n", lastIdx)

	// Find
	value, found := sliceutil.Find(numbers, func(n int) bool { return n > 3 })
	fmt.Printf("First value > 3: %v (found: %v)\n", value, found)

	// FindIndex
	findIdx := sliceutil.FindIndex(numbers, func(n int) bool { return n > 3 })
	fmt.Printf("Index of first value > 3: %d\n", findIdx)

	// Count
	count := sliceutil.Count(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Count of even numbers: %d\n", count)

	// IsEmpty
	empty := []int{}
	fmt.Printf("Empty slice is empty: %v\n", sliceutil.IsEmpty(empty))
	fmt.Printf("Numbers slice is empty: %v\n", sliceutil.IsEmpty(numbers))

	// IsNotEmpty
	fmt.Printf("Numbers slice is not empty: %v\n", sliceutil.IsNotEmpty(numbers))

	// Equal
	other := []int{1, 2, 3, 4, 5}
	fmt.Printf("numbers == other: %v\n", sliceutil.Equal(numbers, other))

	fmt.Println()
}

// ============================================================================
// 2. Transformation Examples (8 functions) / 변환 예제 (8개 함수)
// ============================================================================

func transformationExamples() {
	fmt.Println("2. TRANSFORMATION FUNCTIONS / 변환 함수")
	fmt.Println("=" + string(make([]byte, 70)))

	numbers := []int{1, 2, 3, 4, 5}

	// Map
	doubled := sliceutil.Map(numbers, func(n int) int { return n * 2 })
	fmt.Printf("Map (x2): %v\n", doubled)

	// Filter
	evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Filter (evens): %v\n", evens)

	// FlatMap
	words := []string{"hello", "world"}
	chars := sliceutil.FlatMap(words, func(s string) []string {
		result := make([]string, len(s))
		for i, c := range s {
			result[i] = string(c)
		}
		return result
	})
	fmt.Printf("FlatMap (chars): %v\n", chars)

	// Flatten
	nested := [][]int{{1, 2}, {3, 4}, {5}}
	flat := sliceutil.Flatten(nested)
	fmt.Printf("Flatten: %v\n", flat)

	// Unique
	withDupes := []int{1, 2, 2, 3, 3, 3, 4}
	unique := sliceutil.Unique(withDupes)
	fmt.Printf("Unique: %v\n", unique)

	// UniqueBy
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
	}
	uniqueByAge := sliceutil.UniqueBy(people, func(p Person) int { return p.Age })
	fmt.Printf("UniqueBy age: %v\n", uniqueByAge)

	// Compact
	consecutive := []int{1, 1, 2, 2, 2, 3, 3, 4}
	compacted := sliceutil.Compact(consecutive)
	fmt.Printf("Compact: %v\n", compacted)

	// Reverse
	reversed := sliceutil.Reverse(numbers)
	fmt.Printf("Reverse: %v\n", reversed)

	fmt.Println()
}

// ============================================================================
// 3. Aggregation Examples (7 functions) / 집계 예제 (7개 함수)
// ============================================================================

func aggregationExamples() {
	fmt.Println("3. AGGREGATION FUNCTIONS / 집계 함수")
	fmt.Println("=" + string(make([]byte, 70)))

	numbers := []int{1, 2, 3, 4, 5}

	// Reduce
	sum := sliceutil.Reduce(numbers, 0, func(acc int, n int) int { return acc + n })
	fmt.Printf("Reduce (sum): %d\n", sum)

	// Sum
	total := sliceutil.Sum(numbers)
	fmt.Printf("Sum: %d\n", total)

	// Min
	min, _ := sliceutil.Min(numbers)
	fmt.Printf("Min: %d\n", min)

	// Max
	max, _ := sliceutil.Max(numbers)
	fmt.Printf("Max: %d\n", max)

	// Average
	avg := sliceutil.Average(numbers)
	fmt.Printf("Average: %.2f\n", avg)

	// GroupBy
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
		{"David", 25},
	}
	grouped := sliceutil.GroupBy(people, func(p Person) int { return p.Age })
	fmt.Printf("GroupBy age: %v\n", grouped)

	// Partition
	evens, odds := sliceutil.Partition(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Partition (evens/odds): %v / %v\n", evens, odds)

	fmt.Println()
}

// ============================================================================
// 4. Slicing Examples (7 functions) / 슬라이싱 예제 (7개 함수)
// ============================================================================

func slicingExamples() {
	fmt.Println("4. SLICING FUNCTIONS / 슬라이싱 함수")
	fmt.Println("=" + string(make([]byte, 70)))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Chunk
	chunks := sliceutil.Chunk(numbers, 3)
	fmt.Printf("Chunk (size 3): %v\n", chunks)

	// Take
	first3 := sliceutil.Take(numbers, 3)
	fmt.Printf("Take (first 3): %v\n", first3)

	// TakeLast
	last3 := sliceutil.TakeLast(numbers, 3)
	fmt.Printf("TakeLast (last 3): %v\n", last3)

	// Drop
	dropped := sliceutil.Drop(numbers, 3)
	fmt.Printf("Drop (first 3): %v\n", dropped)

	// DropLast
	droppedLast := sliceutil.DropLast(numbers, 3)
	fmt.Printf("DropLast (last 3): %v\n", droppedLast)

	// Slice
	middle := sliceutil.Slice(numbers, 3, 7)
	fmt.Printf("Slice (3:7): %v\n", middle)

	// Slice with negative indices / 음수 인덱스로 슬라이스
	lastFew := sliceutil.Slice(numbers, -3, -1)
	fmt.Printf("Slice (-3:-1): %v\n", lastFew)

	// Sample
	sample := sliceutil.Sample(numbers, 3)
	fmt.Printf("Sample (3 random): %v\n", sample)

	fmt.Println()
}

// ============================================================================
// 5. Set Operations Examples (6 functions) / 집합 작업 예제 (6개 함수)
// ============================================================================

func setOperationsExamples() {
	fmt.Println("5. SET OPERATIONS / 집합 작업")
	fmt.Println("=" + string(make([]byte, 70)))

	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6, 7, 8}

	// Union
	union := sliceutil.Union(a, b)
	fmt.Printf("Union: %v\n", union)

	// Intersection
	intersection := sliceutil.Intersection(a, b)
	fmt.Printf("Intersection: %v\n", intersection)

	// Difference
	diff := sliceutil.Difference(a, b)
	fmt.Printf("Difference (a - b): %v\n", diff)

	// SymmetricDifference
	symDiff := sliceutil.SymmetricDifference(a, b)
	fmt.Printf("SymmetricDifference: %v\n", symDiff)

	// IsSubset
	subset := []int{2, 3, 4}
	isSubset := sliceutil.IsSubset(subset, a)
	fmt.Printf("%v is subset of %v: %v\n", subset, a, isSubset)

	// IsSuperset
	isSuperset := sliceutil.IsSuperset(a, subset)
	fmt.Printf("%v is superset of %v: %v\n", a, subset, isSuperset)

	fmt.Println()
}

// ============================================================================
// 6. Sorting Examples (5 functions) / 정렬 예제 (5개 함수)
// ============================================================================

func sortingExamples() {
	fmt.Println("6. SORTING FUNCTIONS / 정렬 함수")
	fmt.Println("=" + string(make([]byte, 70)))

	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// Sort
	sorted := sliceutil.Sort(numbers)
	fmt.Printf("Sort: %v\n", sorted)

	// SortDesc
	sortedDesc := sliceutil.SortDesc(numbers)
	fmt.Printf("SortDesc: %v\n", sortedDesc)

	// SortBy
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sortedByAge := sliceutil.SortBy(people, func(p Person) int { return p.Age })
	fmt.Printf("SortBy age: %v\n", sortedByAge)

	// IsSorted
	isSorted := sliceutil.IsSorted(sorted)
	fmt.Printf("Is sorted: %v\n", isSorted)

	// IsSortedDesc
	isSortedDesc := sliceutil.IsSortedDesc(sortedDesc)
	fmt.Printf("Is sorted desc: %v\n", isSortedDesc)

	fmt.Println()
}

// ============================================================================
// 7. Predicate Examples (6 functions) / 조건 검사 예제 (6개 함수)
// ============================================================================

func predicateExamples() {
	fmt.Println("7. PREDICATE FUNCTIONS / 조건 검사 함수")
	fmt.Println("=" + string(make([]byte, 70)))

	numbers := []int{2, 4, 6, 8}

	// All
	allEven := sliceutil.All(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("All even: %v\n", allEven)

	// Any
	mixed := []int{1, 3, 5, 6}
	anyEven := sliceutil.Any(mixed, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Any even: %v\n", anyEven)

	// None
	odds := []int{1, 3, 5, 7}
	noneEven := sliceutil.None(odds, func(n int) bool { return n%2 == 0 })
	fmt.Printf("None even: %v\n", noneEven)

	// AllEqual
	sameValues := []int{5, 5, 5, 5}
	allEqual := sliceutil.AllEqual(sameValues)
	fmt.Printf("All equal: %v\n", allEqual)

	// IsSortedBy
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	sortedByAge := sliceutil.IsSortedBy(people, func(p Person) int { return p.Age })
	fmt.Printf("Is sorted by age: %v\n", sortedByAge)

	// ContainsAll
	all := []int{1, 2, 3, 4, 5}
	hasAll := sliceutil.ContainsAll(all, 2, 4)
	fmt.Printf("Contains 2 and 4: %v\n", hasAll)

	fmt.Println()
}

// ============================================================================
// 8. Utility Examples (11 functions) / 유틸리티 예제 (11개 함수)
// ============================================================================

func utilityExamples() {
	fmt.Println("8. UTILITY FUNCTIONS / 유틸리티 함수")
	fmt.Println("=" + string(make([]byte, 70)))

	numbers := []int{1, 2, 3, 4, 5}

	// ForEach
	fmt.Print("ForEach (print): ")
	sliceutil.ForEach(numbers, func(n int) {
		fmt.Printf("%d ", n)
	})
	fmt.Println()

	// ForEachIndexed
	fmt.Print("ForEachIndexed: ")
	sliceutil.ForEachIndexed(numbers, func(i int, n int) {
		fmt.Printf("[%d]=%d ", i, n)
	})
	fmt.Println()

	// Join
	joined := sliceutil.Join(numbers, ", ")
	fmt.Printf("Join: %s\n", joined)

	// Clone
	cloned := sliceutil.Clone(numbers)
	fmt.Printf("Clone: %v\n", cloned)

	// Fill
	filled := sliceutil.Fill(numbers, 0)
	fmt.Printf("Fill (0): %v\n", filled)

	// Insert
	inserted := sliceutil.Insert(numbers, 2, 99, 100)
	fmt.Printf("Insert at 2: %v\n", inserted)

	// Remove
	removed := sliceutil.Remove(numbers, 2)
	fmt.Printf("Remove at 2: %v\n", removed)

	// RemoveAll
	withDupes := []int{1, 2, 3, 2, 4, 2, 5}
	removedAll := sliceutil.RemoveAll(withDupes, 2)
	fmt.Printf("RemoveAll (2): %v\n", removedAll)

	// Shuffle
	shuffled := sliceutil.Shuffle(numbers)
	fmt.Printf("Shuffle: %v\n", shuffled)

	// Zip
	words := []string{"one", "two", "three"}
	zipped := sliceutil.Zip(numbers[:3], words)
	fmt.Printf("Zip: %v\n", zipped)

	// Unzip
	nums, strs := sliceutil.Unzip[int, string](zipped)
	fmt.Printf("Unzip: %v / %v\n", nums, strs)

	fmt.Println()
}

// ============================================================================
// 9. Real-World Examples / 실제 사용 시나리오
// ============================================================================

func realWorldExamples() {
	fmt.Println("9. REAL-WORLD SCENARIOS / 실제 사용 시나리오")
	fmt.Println("=" + string(make([]byte, 70)))

	// Scenario 1: Processing user data / 시나리오 1: 사용자 데이터 처리
	type User struct {
		Name     string
		Age      int
		IsActive bool
	}

	users := []User{
		{"Alice", 30, true},
		{"Bob", 25, false},
		{"Charlie", 35, true},
		{"David", 28, true},
		{"Eve", 32, false},
	}

	// Get active users over 30 / 30세 이상의 활성 사용자 가져오기
	activeOver30 := sliceutil.Filter(users, func(u User) bool {
		return u.IsActive && u.Age > 30
	})
	fmt.Printf("Active users over 30: %v\n", activeOver30)

	// Get average age of active users / 활성 사용자의 평균 나이 구하기
	activeUsers := sliceutil.Filter(users, func(u User) bool { return u.IsActive })
	ages := sliceutil.Map(activeUsers, func(u User) int { return u.Age })
	avgAge := sliceutil.Average(ages)
	fmt.Printf("Average age of active users: %.1f\n", avgAge)

	// Scenario 2: Processing product data / 시나리오 2: 제품 데이터 처리
	type Product struct {
		Name  string
		Price float64
		Stock int
	}

	products := []Product{
		{"Laptop", 999.99, 10},
		{"Mouse", 29.99, 50},
		{"Keyboard", 79.99, 30},
		{"Monitor", 299.99, 15},
		{"Webcam", 89.99, 0},
	}

	// Get in-stock products sorted by price / 재고가 있는 제품을 가격순으로 정렬
	inStock := sliceutil.Filter(products, func(p Product) bool { return p.Stock > 0 })
	sortedByPrice := sliceutil.SortBy(inStock, func(p Product) float64 { return p.Price })
	fmt.Printf("In-stock products (by price): %v\n", sortedByPrice)

	// Get total inventory value / 총 재고 가치 구하기
	values := sliceutil.Map(products, func(p Product) float64 {
		return p.Price * float64(p.Stock)
	})
	totalValue := sliceutil.Sum(values)
	fmt.Printf("Total inventory value: $%.2f\n", totalValue)

	// Scenario 3: Data analysis pipeline / 시나리오 3: 데이터 분석 파이프라인
	temperatures := []float64{23.5, 25.1, 24.8, 26.3, 25.9, 24.2, 23.8}

	// Calculate statistics / 통계 계산
	avgTemp := sliceutil.Average(temperatures)
	minTemp, _ := sliceutil.Min(temperatures)
	maxTemp, _ := sliceutil.Max(temperatures)

	fmt.Printf("\nTemperature Statistics:\n")
	fmt.Printf("  Average: %.2f°C\n", avgTemp)
	fmt.Printf("  Min: %.2f°C\n", minTemp)
	fmt.Printf("  Max: %.2f°C\n", maxTemp)

	// Find outliers (>2°C from average) / 이상값 찾기 (평균에서 >2°C)
	outliers := sliceutil.Filter(temperatures, func(t float64) bool {
		diff := t - avgTemp
		if diff < 0 {
			diff = -diff
		}
		return diff > 2.0
	})
	fmt.Printf("  Outliers: %v\n", outliers)

	fmt.Println()
	fmt.Println("=== All 60 Functions Demonstrated! / 모든 60개 함수 시연 완료! ===")
}
