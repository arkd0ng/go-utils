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
// This package provides 60 functions across 8 categories for common slice operations:
//
// 이 패키지는 일반적인 슬라이스 작업을 위한 8개 카테고리에 걸쳐 60개의 함수를 제공합니다:
//
//  1. Basic Operations (10 functions) - Contains, IndexOf, Find, etc.
//  2. Transformation (8 functions) - Map, Filter, Unique, Reverse, etc.
//  3. Aggregation (7 functions) - Reduce, Sum, Min, Max, GroupBy, etc.
//  4. Slicing (7 functions) - Chunk, Take, Drop, Sample, etc.
//  5. Set Operations (6 functions) - Union, Intersection, Difference, etc.
//  6. Sorting (5 functions) - Sort, SortBy, IsSorted, etc.
//  7. Predicates (6 functions) - All, Any, None, AllEqual, etc.
//  8. Utilities (11 functions) - ForEach, Join, Clone, Shuffle, Zip, etc.
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
// Current version: v1.7.x
//
// For more information, see:
//   - Package README: https://github.com/arkd0ng/go-utils/tree/main/sliceutil
//   - User Manual: https://github.com/arkd0ng/go-utils/tree/main/docs/sliceutil/USER_MANUAL.md
//   - Developer Guide: https://github.com/arkd0ng/go-utils/tree/main/docs/sliceutil/DEVELOPER_GUIDE.md
package sliceutil

// Version is the current package version.
// Version은 현재 패키지 버전입니다.
const Version = "1.7.015"
