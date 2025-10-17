package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
	"github.com/arkd0ng/go-utils/sliceutil"
)

// User represents a user in the system
// User는 시스템의 사용자를 나타냅니다
type User struct {
	ID   int
	Name string
	Age  int
	City string
}

func main() {
	// Setup log file with backup management
	// 백업 관리와 함께 로그 파일 설정
	logFilePath := "logs/sliceutil-example.log"

	// Check if previous log file exists
	// 이전 로그 파일 존재 여부 확인
	if fileutil.Exists(logFilePath) {
		// Get modification time of existing log file
		// 기존 로그 파일의 수정 시간 가져오기
		modTime, err := fileutil.ModTime(logFilePath)
		if err == nil {
			// Create backup filename with timestamp
			// 타임스탬프와 함께 백업 파일명 생성
			backupName := fmt.Sprintf("logs/sliceutil-example-%s.log", modTime.Format("20060102-150405"))

			// Backup existing log file
			// 기존 로그 파일 백업
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
				// Delete original log file to prevent content duplication
				// 내용 중복 방지를 위해 원본 로그 파일 삭제
				fileutil.DeleteFile(logFilePath)
			}
		}

		// Cleanup old backup files - keep only 5 most recent
		// 오래된 백업 파일 정리 - 최근 5개만 유지
		backupPattern := "logs/sliceutil-example-*.log"
		backupFiles, err := filepath.Glob(backupPattern)
		if err == nil && len(backupFiles) > 5 {
			// Sort by modification time
			// 수정 시간으로 정렬
			type fileInfo struct {
				path    string
				modTime time.Time
			}
			var files []fileInfo
			for _, f := range backupFiles {
				if mt, err := fileutil.ModTime(f); err == nil {
					files = append(files, fileInfo{path: f, modTime: mt})
				}
			}

			// Sort oldest first
			// 가장 오래된 것부터 정렬
			for i := 0; i < len(files)-1; i++ {
				for j := i + 1; j < len(files); j++ {
					if files[i].modTime.After(files[j].modTime) {
						files[i], files[j] = files[j], files[i]
					}
				}
			}

			// Delete oldest files to keep only 5
			// 5개만 유지하도록 가장 오래된 파일 삭제
			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Printf("🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	// Initialize logger with fixed filename
	// 고정 파일명으로 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithLevel(logging.DEBUG),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Printf("Failed to create logger: %v\n", err)
		return
	}
	defer logger.Close()

	// Print banner
	// 배너 출력
	logger.Banner("Sliceutil Package - Comprehensive Examples", "v1.9.013")
	logger.Info("")

	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║            Sliceutil Package - Comprehensive Examples                      ║")
	logger.Info("║            Sliceutil 패키지 - 종합 예제                                     ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	logger.Info("📋 Package Information / 패키지 정보")
	logger.Info("   Package: github.com/arkd0ng/go-utils/sliceutil")
	logger.Info("   Description: Extremely simple slice utilities (20 lines → 1 line)")
	logger.Info("   설명: 극도로 간단한 슬라이스 유틸리티 (20줄 → 1줄)")
	logger.Info("   Total Functions: 95 functions across 14 categories")
	logger.Info("   Go 1.18+ Generics: Type-safe slice operations")
	logger.Info("   Zero Dependencies: Standard library only (golang.org/x/exp excluded)")
	logger.Info("")

	logger.Info("🌟 Key Features / 주요 기능")
	logger.Info("   • Type Safety: Go 1.18+ generics with compile-time type checking")
	logger.Info("   • Functional Style: Inspired by JavaScript, Python, Ruby array methods")
	logger.Info("   • Immutability: All operations return new slices (no mutation)")
	logger.Info("   • Generic Constraints: Number, Ordered, comparable constraints")
	logger.Info("   • High Performance: Efficient algorithms (mostly O(n) time)")
	logger.Info("   • 100% Test Coverage: Complete tests for all functions")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣  Basic Operations (11 functions)")
	logger.Info("   기본 작업 (11개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1.1 Filter() - Filter elements by predicate")
	logger.Info("    조건에 맞는 요소 필터링")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Filter[T any](slice []T, predicate func(T) bool) []T")
	logger.Info("")

	logger.Info("📖 Description / 설명:")
	logger.Info("   Returns new slice containing only elements that satisfy the predicate")
	logger.Info("   조건을 만족하는 요소만 포함하는 새 슬라이스 반환")
	logger.Info("   • Generic function works with any type")
	logger.Info("   • Immutable: original slice unchanged")
	logger.Info("")

	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Filter active users (활성 사용자 필터링)")
	logger.Info("   • Remove invalid data (잘못된 데이터 제거)")
	logger.Info("   • Select items by criteria (조건별 항목 선택)")
	logger.Info("   • Data validation (데이터 검증)")
	logger.Info("")

	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Type-safe generic implementation")
	logger.Info("   • Custom predicate function")
	logger.Info("   • O(n) time complexity")
	logger.Info("   • Returns new slice (immutable)")
	logger.Info("")

	logger.Info("▶️  Executing / 실행 중:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
	logger.Info(fmt.Sprintf("   Original: %v", numbers))
	logger.Info(fmt.Sprintf("   Filter(even numbers): %v", evens))
	logger.Info("")

	logger.Info("✅ Results Analysis / 결과 분석:")
	logger.Info(fmt.Sprintf("   • Original slice unchanged: %v", numbers))
	logger.Info(fmt.Sprintf("   • Filtered result: %v (only even numbers)", evens))
	logger.Info("   • New slice allocated (immutable operation)")
	logger.Info("   • Type-safe: compiler ensures int → int")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1.2 Map() - Transform elements")
	logger.Info("    요소 변환")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Map[T any, R any](slice []T, mapper func(T) R) []R")
	logger.Info("")

	logger.Info("📖 Description / 설명:")
	logger.Info("   Transforms each element using mapper function, returns new slice")
	logger.Info("   매퍼 함수를 사용하여 각 요소를 변환하고 새 슬라이스 반환")
	logger.Info("")

	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Extract properties (속성 추출)")
	logger.Info("   • Type conversion (타입 변환)")
	logger.Info("   • Data transformation (데이터 변환)")
	logger.Info("   • Formatting (포맷팅)")
	logger.Info("")

	logger.Info("▶️  Executing / 실행 중:")
	users := []User{
		{ID: 1, Name: "Alice", Age: 25, City: "Seoul"},
		{ID: 2, Name: "Bob", Age: 30, City: "Busan"},
		{ID: 3, Name: "Charlie", Age: 35, City: "Incheon"},
	}
	names := sliceutil.Map(users, func(u User) string { return u.Name })
	logger.Info(fmt.Sprintf("   Users: %d users", len(users)))
	logger.Info(fmt.Sprintf("   Map(extract names): %v", names))
	logger.Info("")

	logger.Info("📝 Additional Basic Functions:")
	logger.Info("   1.3 Contains() - Check if slice contains element")
	logger.Info("   1.4 ContainsFunc() - Check with predicate")
	logger.Info("   1.5 IndexOf() - Find first index of element")
	logger.Info("   1.6 LastIndexOf() - Find last index")
	logger.Info("   1.7 Find() - Find first matching element")
	logger.Info("   1.8 FindLast() - Find last matching element")
	logger.Info("   1.9 FindIndex() - Find index by predicate")
	logger.Info("   1.10 Count() - Count matching elements")
	logger.Info("   1.11 Equal() - Check slice equality")
	logger.Info("")

	contains := sliceutil.Contains(numbers, 5)
	indexOf := sliceutil.IndexOf(numbers, 7)
	logger.Info(fmt.Sprintf("   Contains(5): %v", contains))
	logger.Info(fmt.Sprintf("   IndexOf(7): %d", indexOf))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("2️⃣  Transformation (8 functions)")
	logger.Info("   변환 (8개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Transformation Functions:")
	logger.Info("   2.1 Map() - Transform elements (already demonstrated)")
	logger.Info("   2.2 Filter() - Filter by predicate (already demonstrated)")
	logger.Info("   2.3 FlatMap() - Map and flatten nested slices")
	logger.Info("   2.4 Flatten() - Flatten nested slices")
	logger.Info("   2.5 Unique() - Remove duplicates")
	logger.Info("   2.6 UniqueBy() - Remove duplicates by key")
	logger.Info("   2.7 Compact() - Remove zero values")
	logger.Info("   2.8 Reverse() - Reverse slice order")
	logger.Info("")

	logger.Info("▶️  Executing Transformation / 변환 실행:")
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	unique := sliceutil.Unique(duplicates)
	logger.Info(fmt.Sprintf("   Unique(%v) = %v", duplicates, unique))

	withZeros := []int{1, 0, 2, 0, 3, 0, 4}
	compact := sliceutil.Compact(withZeros)
	logger.Info(fmt.Sprintf("   Compact(%v) = %v", withZeros, compact))

	reversed := sliceutil.Reverse(numbers[0:5])
	logger.Info(fmt.Sprintf("   Reverse([1 2 3 4 5]) = %v", reversed))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("3️⃣  Aggregation (11 functions)")
	logger.Info("   집계 (11개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("3.1 Reduce() - Reduce to single value")
	logger.Info("    단일 값으로 축소")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Reduce[T any, R any](slice []T, initial R, reducer func(R, T) R) R")
	logger.Info("")

	logger.Info("📖 Description / 설명:")
	logger.Info("   Reduces slice to single value by applying reducer function")
	logger.Info("   리듀서 함수를 적용하여 슬라이스를 단일 값으로 축소")
	logger.Info("")

	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Sum all numbers (모든 숫자 합계)")
	logger.Info("   • Calculate totals (총계 계산)")
	logger.Info("   • Concatenate strings (문자열 연결)")
	logger.Info("   • Complex aggregations (복잡한 집계)")
	logger.Info("")

	logger.Info("▶️  Executing / 실행 중:")
	sum := sliceutil.Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	logger.Info(fmt.Sprintf("   Reduce(sum): %d", sum))

	product := sliceutil.Reduce([]int{1, 2, 3, 4, 5}, 1, func(acc, n int) int { return acc * n })
	logger.Info(fmt.Sprintf("   Reduce(product 1-5): %d", product))
	logger.Info("")

	logger.Info("📝 Additional Aggregation Functions:")
	logger.Info("   3.2 ReduceRight() - Reduce from right to left")
	logger.Info("   3.3 Sum() - Sum all numbers")
	logger.Info("   3.4 Min() - Find minimum value")
	logger.Info("   3.5 Max() - Find maximum value")
	logger.Info("   3.6 MinBy() - Find min by key function")
	logger.Info("   3.7 MaxBy() - Find max by key function")
	logger.Info("   3.8 Average() - Calculate average")
	logger.Info("   3.9 GroupBy() - Group elements by key")
	logger.Info("   3.10 CountBy() - Count elements by key")
	logger.Info("   3.11 Partition() - Split into two slices by predicate")
	logger.Info("")

	minVal, _ := sliceutil.Min(numbers)
	maxVal, _ := sliceutil.Max(numbers)
	avg := sliceutil.Average(numbers)
	logger.Info(fmt.Sprintf("   Min: %d, Max: %d, Average: %.2f", minVal, maxVal, avg))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("4️⃣  Slicing Operations (11 functions)")
	logger.Info("   슬라이싱 작업 (11개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Slicing Functions:")
	logger.Info("   4.1 Chunk() - Split into chunks of size n")
	logger.Info("   4.2 Slice() - Safe slice with bounds checking")
	logger.Info("   4.3 Take() - Take first n elements")
	logger.Info("   4.4 TakeLast() - Take last n elements")
	logger.Info("   4.5 TakeWhile() - Take while predicate is true")
	logger.Info("   4.6 Drop() - Drop first n elements")
	logger.Info("   4.7 DropLast() - Drop last n elements")
	logger.Info("   4.8 DropWhile() - Drop while predicate is true")
	logger.Info("   4.9 Sample() - Random sample of n elements")
	logger.Info("   4.10 Window() - Sliding window of size n")
	logger.Info("   4.11 Interleave() - Interleave multiple slices")
	logger.Info("")

	logger.Info("▶️  Executing Slicing / 슬라이싱 실행:")
	chunks := sliceutil.Chunk(numbers, 3)
	logger.Info(fmt.Sprintf("   Chunk(size=3): %v", chunks))

	first3 := sliceutil.Take(numbers, 3)
	logger.Info(fmt.Sprintf("   Take(3): %v", first3))

	last3 := sliceutil.TakeLast(numbers, 3)
	logger.Info(fmt.Sprintf("   TakeLast(3): %v", last3))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("5️⃣  Set Operations (6 functions)")
	logger.Info("   집합 작업 (6개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Set Functions:")
	logger.Info("   5.1 Union() - Union of two slices")
	logger.Info("   5.2 Intersection() - Intersection of two slices")
	logger.Info("   5.3 Difference() - Elements in first but not second")
	logger.Info("   5.4 SymmetricDifference() - Elements in either but not both")
	logger.Info("   5.5 IsSubset() - Check if first is subset of second")
	logger.Info("   5.6 IsSuperset() - Check if first is superset of second")
	logger.Info("")

	logger.Info("▶️  Executing Set Operations / 집합 작업 실행:")
	setA := []int{1, 2, 3, 4, 5}
	setB := []int{4, 5, 6, 7, 8}

	union := sliceutil.Union(setA, setB)
	logger.Info(fmt.Sprintf("   Union(%v, %v) = %v", setA, setB, union))

	intersection := sliceutil.Intersection(setA, setB)
	logger.Info(fmt.Sprintf("   Intersection = %v", intersection))

	difference := sliceutil.Difference(setA, setB)
	logger.Info(fmt.Sprintf("   Difference(A-B) = %v", difference))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("6️⃣  Sorting (6 functions)")
	logger.Info("   정렬 (6개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Sorting Functions:")
	logger.Info("   6.1 Sort() - Sort in ascending order")
	logger.Info("   6.2 SortDesc() - Sort in descending order")
	logger.Info("   6.3 SortBy() - Sort by key function")
	logger.Info("   6.4 SortByMulti() - Sort by multiple keys")
	logger.Info("   6.5 IsSorted() - Check if sorted ascending")
	logger.Info("   6.6 IsSortedDesc() - Check if sorted descending")
	logger.Info("")

	logger.Info("▶️  Executing Sorting / 정렬 실행:")
	unsorted := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	sorted := sliceutil.Sort(unsorted)
	logger.Info(fmt.Sprintf("   Sort(%v) = %v", unsorted, sorted))

	sortedDesc := sliceutil.SortDesc(unsorted)
	logger.Info(fmt.Sprintf("   SortDesc = %v", sortedDesc))

	usersSorted := sliceutil.SortBy(users, func(u User) int { return u.Age })
	for _, u := range usersSorted {
		logger.Info(fmt.Sprintf("   - %s (age %d)", u.Name, u.Age))
	}
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("7️⃣  Predicates (6 functions)")
	logger.Info("   조건 검사 (6개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Predicate Functions:")
	logger.Info("   7.1 All() - Check if all elements satisfy predicate")
	logger.Info("   7.2 Any() - Check if any element satisfies predicate")
	logger.Info("   7.3 None() - Check if no elements satisfy predicate")
	logger.Info("   7.4 AllEqual() - Check if all elements are equal")
	logger.Info("   7.5 ContainsAll() - Check if contains all elements")
	logger.Info("   7.6 IsSortedBy() - Check if sorted by key")
	logger.Info("")

	logger.Info("▶️  Executing Predicates / 조건 검사 실행:")
	allPositive := sliceutil.All(numbers, func(n int) bool { return n > 0 })
	logger.Info(fmt.Sprintf("   All positive: %v", allPositive))

	anyGreater := sliceutil.Any(numbers, func(n int) bool { return n > 5 })
	logger.Info(fmt.Sprintf("   Any > 5: %v", anyGreater))

	noneNegative := sliceutil.None(numbers, func(n int) bool { return n < 0 })
	logger.Info(fmt.Sprintf("   None negative: %v", noneNegative))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("8️⃣  Utilities (12 functions)")
	logger.Info("   유틸리티 (12개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Utility Functions:")
	logger.Info("   8.1 ForEach() - Iterate with side effects")
	logger.Info("   8.2 ForEachIndexed() - Iterate with index")
	logger.Info("   8.3 Tap() - Execute function and return original")
	logger.Info("   8.4 Clone() - Deep copy slice")
	logger.Info("   8.5 Fill() - Fill with value")
	logger.Info("   8.6 Insert() - Insert at index")
	logger.Info("   8.7 Remove() - Remove element")
	logger.Info("   8.8 RemoveAll() - Remove all matching")
	logger.Info("   8.9 Join() - Join to string")
	logger.Info("   8.10 Shuffle() - Random shuffle")
	logger.Info("   8.11 Zip() - Zip two slices")
	logger.Info("   8.12 Unzip() - Unzip pairs")
	logger.Info("")

	logger.Info("▶️  Executing Utilities / 유틸리티 실행:")
	cloned := sliceutil.Clone(numbers[0:5])
	logger.Info(fmt.Sprintf("   Clone([1 2 3 4 5]): %v", cloned))

	inserted := sliceutil.Insert(numbers[0:5], 2, 99)
	logger.Info(fmt.Sprintf("   Insert(99 at index 2): %v", inserted))

	joined := sliceutil.Join([]string{"Go", "is", "awesome"}, " ")
	logger.Info(fmt.Sprintf("   Join(['Go', 'is', 'awesome'], ' '): '%s'", joined))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📊 Summary / 요약")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("This example demonstrated comprehensive slice utilities:")
	logger.Info("본 예제는 포괄적인 슬라이스 유틸리티를 시연했습니다:")
	logger.Info("")

	logger.Info("  1️⃣  Basic Operations (11 functions) - Filter, Map, Contains, Find, etc.")
	logger.Info("     기본 작업 (11개 함수) - 필터, 맵, 포함, 찾기 등")
	logger.Info("  2️⃣  Transformation (8 functions) - Map, Filter, Unique, Reverse, etc.")
	logger.Info("     변환 (8개 함수) - 맵, 필터, 중복 제거, 역순 등")
	logger.Info("  3️⃣  Aggregation (11 functions) - Reduce, Sum, Min, Max, Average, etc.")
	logger.Info("     집계 (11개 함수) - 축소, 합계, 최소, 최대, 평균 등")
	logger.Info("  4️⃣  Slicing (11 functions) - Chunk, Take, Drop, Window, etc.")
	logger.Info("     슬라이싱 (11개 함수) - 청크, 가져오기, 버리기, 윈도우 등")
	logger.Info("  5️⃣  Set Operations (6 functions) - Union, Intersection, Difference, etc.")
	logger.Info("     집합 작업 (6개 함수) - 합집합, 교집합, 차집합 등")
	logger.Info("  6️⃣  Sorting (6 functions) - Sort, SortBy, IsSorted, etc.")
	logger.Info("     정렬 (6개 함수) - 정렬, 키별 정렬, 정렬 확인 등")
	logger.Info("  7️⃣  Predicates (6 functions) - All, Any, None, etc.")
	logger.Info("     조건 검사 (6개 함수) - 모두, 하나라도, 없음 등")
	logger.Info("  8️⃣  Utilities (12 functions) - ForEach, Clone, Insert, Join, etc.")
	logger.Info("     유틸리티 (12개 함수) - 반복, 복제, 삽입, 결합 등")
	logger.Info("")

	logger.Info("  Plus 6 more categories with 24 additional functions:")
	logger.Info("  추가로 6개 카테고리, 24개 함수:")
	logger.Info("  • Combinatorial (2 functions) - Permutations, Combinations")
	logger.Info("  • Statistics (8 functions) - Median, Mode, StandardDeviation, etc.")
	logger.Info("  • Diff/Comparison (4 functions) - Diff, EqualUnordered, etc.")
	logger.Info("  • Index-based (3 functions) - FindIndices, AtIndices, RemoveIndices")
	logger.Info("  • Conditional (3 functions) - ReplaceIf, ReplaceAll, UpdateWhere")
	logger.Info("  • Advanced (4 functions) - Scan, ZipWith, RotateLeft, RotateRight")
	logger.Info("")

	logger.Info("✨ Key Takeaways / 주요 포인트:")
	logger.Info("   • All 95 functions production-ready (95개 함수 모두 프로덕션 준비 완료)")
	logger.Info("   • Type-safe with Go 1.18+ generics (Go 1.18+ 제네릭으로 타입 안전)")
	logger.Info("   • Functional programming patterns (함수형 프로그래밍 패턴)")
	logger.Info("   • Immutable operations (불변 작업)")
	logger.Info("   • Zero dependencies (제로 의존성)")
	logger.Info("   • 100% test coverage (100% 테스트 커버리지)")
	logger.Info("   • JavaScript/Python/Ruby-inspired API (JavaScript/Python/Ruby에서 영감)")
	logger.Info("")

	logger.Info("All examples completed successfully!")
	logger.Info("모든 예제가 성공적으로 완료되었습니다!")
	logger.Info("")
}
