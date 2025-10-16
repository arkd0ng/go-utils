// Package sliceutil provides extreme simplicity for slice operations in Go.
//
// # Design Philosophy / 설계 철학
//
// "20 lines → 1 line" - Reduce repetitive slice manipulation code to simple function calls.
//
// "20줄 → 1줄" - 반복적인 슬라이스 조작 코드를 간단한 함수 호출로 줄입니다.
//
// # Overview / 개요
//
// This package provides 95 functions across 14 categories for common slice operations:
//
// 이 패키지는 일반적인 슬라이스 작업을 위한 14개 카테고리에 걸쳐 95개의 함수를 제공합니다:
//
//  1. Basic Operations (11 functions) - Contains, IndexOf, Find, FindLast, Count, etc.
//  2. Transformation (8 functions) - Map, Filter, Unique, Reverse, Flatten, FlatMap, etc.
//  3. Aggregation (11 functions) - Reduce, ReduceRight, Sum, Min, Max, MinBy, MaxBy, Average, GroupBy, CountBy, Partition
//  4. Slicing (10 functions) - Chunk, Take, TakeLast, Drop, DropLast, TakeWhile, DropWhile, Slice, Sample, Interleave
//  5. Set Operations (6 functions) - Union, Intersection, Difference, SymmetricDifference, IsSubset, IsSuperset
//  6. Sorting (6 functions) - Sort, SortDesc, SortBy, SortByMulti, IsSorted, IsSortedDesc
//  7. Predicates (6 functions) - All, Any, None, AllEqual, ContainsAll, IsSortedBy
//  8. Utilities (13 functions) - ForEach, ForEachIndexed, Join, Clone, Fill, Insert, Remove, RemoveAll, Shuffle, Zip, Unzip, Window, Tap
//  9. Combinatorial (2 functions) - Permutations, Combinations
//
// 10. Statistics (8 functions) - Median, Mode, Frequencies, Percentile, StandardDeviation, Variance, MostCommon, LeastCommon
// 11. Diff/Comparison (4 functions) - Diff, DiffBy, EqualUnordered, HasDuplicates
// 12. Index-based (3 functions) - FindIndices, AtIndices, RemoveIndices
// 13. Conditional (3 functions) - ReplaceIf, ReplaceAll, UpdateWhere
// 14. Advanced (4 functions) - Scan, ZipWith, RotateLeft, RotateRight
//
// # Key Features / 주요 기능
//
//   - Type-safe with Go 1.18+ generics / Go 1.18+ 제네릭으로 타입 안전
//   - Functional programming style / 함수형 프로그래밍 스타일
//   - Immutable operations / 불변 작업
//   - Zero external dependencies / 제로 외부 의존성
//   - Comprehensive coverage of common operations / 일반적인 작업의 포괄적인 커버리지
//
// # Example / 예제
//
// Before (Standard Go):
//
//	// Filter even numbers
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	var evens []int
//	for _, n := range numbers {
//	    if n%2 == 0 {
//	        evens = append(evens, n)
//	    }
//	}
//	// 8+ lines of code
//
// After (This Package):
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
//	// 1 line of code (vs 8+)
//
// # Usage / 사용법
//
// Import the package:
//
//	import "github.com/arkd0ng/go-utils/sliceutil"
//
// Use any function:
//
//	// Basic Operations / 기본 작업
//	found := sliceutil.Contains([]int{1, 2, 3}, 2)
//	index := sliceutil.IndexOf([]string{"a", "b", "c"}, "b")
//
//	// Transformation / 변환
//	doubled := sliceutil.Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
//	evens := sliceutil.Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
//	unique := sliceutil.Unique([]int{1, 2, 2, 3, 3, 3})
//
//	// Aggregation / 집계
//	sum := sliceutil.Sum([]int{1, 2, 3, 4, 5})
//	max, _ := sliceutil.Max([]int{1, 5, 3, 9, 2})
//	grouped := sliceutil.GroupBy(users, func(u User) string { return u.City })
//
//	// Set Operations / 집합 작업
//	union := sliceutil.Union([]int{1, 2, 3}, []int{3, 4, 5})
//	intersection := sliceutil.Intersection([]int{1, 2, 3}, []int{2, 3, 4})
//
// # Performance / 성능
//
// All functions are implemented with performance in mind, using efficient algorithms
// and minimal allocations. For performance-critical code, benchmarks are provided.
//
// 모든 함수는 효율적인 알고리즘과 최소한의 할당을 사용하여 성능을 고려하여 구현되었습니다.
// 성능이 중요한 코드의 경우 벤치마크가 제공됩니다.
//
// # Version / 버전
//
// Current version is loaded automatically from cfg/app.yaml.
// 현재 버전은 cfg/app.yaml에서 자동으로 로드됩니다.
//
// For more information, see:
//   - Package README: https://github.com/arkd0ng/go-utils/tree/main/sliceutil
//   - User Manual: https://github.com/arkd0ng/go-utils/tree/main/docs/sliceutil/USER_MANUAL.md
//   - Developer Guide: https://github.com/arkd0ng/go-utils/tree/main/docs/sliceutil/DEVELOPER_GUIDE.md
package sliceutil

import (
	"math/rand"
	"sync"
	"time"

	"github.com/arkd0ng/go-utils/logging"
)

// Version is the current package version loaded from cfg/app.yaml.
// Version은 cfg/app.yaml에서 로드되는 현재 패키지 버전입니다.
var Version = getVersion()

// getVersion loads the application version from cfg/app.yaml.
// getVersion은 cfg/app.yaml에서 애플리케이션 버전을 로드합니다.
func getVersion() string {
	version := logging.TryLoadAppVersion()
	if version == "" {
		return "unknown"
	}
	return version
}

// Global random number generator for thread-safe random operations
// 스레드 안전 랜덤 작업을 위한 전역 랜덤 번호 생성기
var (
	rng     = rand.New(rand.NewSource(time.Now().UnixNano()))
	rngLock sync.Mutex
)
