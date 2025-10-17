// Package maputil provides extreme simplicity map utilities for Go.
// 패키지 maputil은 Go를 위한 극도로 간단한 맵 유틸리티를 제공합니다.
//
// This package offers 80+ type-safe functions across 10 categories to simplify
// common map operations. Reduce 20+ lines of repetitive code to just 1-2 lines.
//
// 이 패키지는 10개 카테고리에 걸쳐 80개 이상의 타입 안전 함수를 제공하여
// 일반적인 맵 작업을 단순화합니다. 20줄 이상의 반복 코드를 단 1-2줄로 줄입니다.
//
// # Key Features
// 주요 기능
//
// - Type-safe operations using Go 1.18+ generics
// Go 1.18+ 제네릭을 사용한 타입 안전 작업
// - Functional programming style (Map, Filter, Reduce)
// 함수형 프로그래밍 스타일
// - Immutable operations (original maps unchanged)
// 불변 작업 (원본 맵 변경 없음)
// - Zero external dependencies
// 외부 의존성 제로
// - 100% test coverage
// 100% 테스트 커버리지
//
// # Categories
// 카테고리
//
//   - Basic Operations (11 functions): Get, Set, Delete, Clone, etc.
//   - Transformation (10 functions): Map, MapKeys, Invert, Flatten, etc.
//   - Aggregation (9 functions): Reduce, Sum, Min, Max, Average, etc.
//   - Merge Operations (8 functions): Merge, Union, Intersection, Difference, etc.
//   - Filter Operations (7 functions): Filter, Pick, Omit, Partition, etc.
//   - Conversion (8 functions): Keys, Values, Entries, FromSlice, ToJSON, etc.
//   - Predicate Checks (7 functions): Every, Some, None, HasValue, etc.
//   - Key Operations (8 functions): KeysSorted, RenameKey, SwapKeys, etc.
//   - Value Operations (7 functions): ValuesSorted, UniqueValues, ReplaceValue, etc.
//   - Comparison (6 functions): Diff, Compare, CommonKeys, etc.
//
// # Example Usage
// 사용 예제
//
// // Filter map
// 맵 필터링
//	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	result := maputil.Filter(data, func(k string, v int) bool {
//	    return v > 2
//	}) // map[string]int{"c": 3, "d": 4}
//
// // Transform values
// 값 변환
//	doubled := maputil.MapValues(data, func(v int) int {
//	    return v * 2
//	}) // map[string]int{"a": 2, "b": 4, "c": 6, "d": 8}
//
// // Merge maps
// 맵 병합
//	map1 := map[string]int{"a": 1, "b": 2}
//	map2 := map[string]int{"b": 3, "c": 4}
//	merged := maputil.Merge(map1, map2) // map[string]int{"a": 1, "b": 3, "c": 4}
//
// // Group slice by key
// 키로 슬라이스 그룹화
//	users := []User{{Name: "Alice", City: "Seoul"}, {Name: "Bob", City: "Seoul"}}
//	byCity := maputil.GroupBy(users, func(u User) string { return u.City })
//
// For more examples and detailed documentation, see the package README and USER_MANUAL.
// 더 많은 예제와 상세한 문서는 패키지 README와 USER_MANUAL을 참조하세요.
package maputil

import "github.com/arkd0ng/go-utils/internal/version"

// Version is the current version of the maputil package.
// Version은 maputil 패키지의 현재 버전입니다.
//
// The version is automatically loaded from cfg/app.yaml.
// If the file cannot be loaded, it returns "unknown".
//
// 버전은 cfg/app.yaml에서 자동으로 로드됩니다.
// 파일을 로드할 수 없으면 "unknown"을 반환합니다.
var Version = version.Get()

// Entry represents a key-value pair in a map.
// Entry는 맵의 키-값 쌍을 나타냅니다.
//
// This type is used for functions that work with map entries,
// such as Entries() and FromEntries().
//
// 이 타입은 Entries() 및 FromEntries()와 같이
// 맵 항목을 다루는 함수에 사용됩니다.
//
// Example
// 예제:
//
//	entries := []Entry[string, int]{
//	    {Key: "a", Value: 1},
//	    {Key: "b", Value: 2},
//	}
//	m := maputil.FromEntries(entries) // map[string]int{"a": 1, "b": 2}
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Number is a constraint that permits any numeric type.
// Number는 모든 숫자 타입을 허용하는 제약조건입니다.
//
// This constraint is used for functions that perform arithmetic operations,
// such as Sum() and Average().
//
// 이 제약조건은 Sum() 및 Average()와 같이
// 산술 연산을 수행하는 함수에 사용됩니다.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Ordered is a constraint that permits any ordered type.
// Ordered는 순서가 있는 모든 타입을 허용하는 제약조건입니다.
//
// This constraint is used for functions that perform comparisons,
// such as Min(), Max(), and sorting operations.
//
// 이 제약조건은 Min(), Max() 및 정렬 작업과 같이
// 비교를 수행하는 함수에 사용됩니다.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}
