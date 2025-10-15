# Sliceutil Package - Developer Guide / 개발자 가이드

**Version / 버전**: v1.7.023
**Package / 패키지**: `github.com/arkd0ng/go-utils/sliceutil`
**Go Version / Go 버전**: 1.18+

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#architecture-overview--아키텍처-개요)
2. [Package Structure / 패키지 구조](#package-structure--패키지-구조)
3. [Core Components / 핵심 컴포넌트](#core-components--핵심-컴포넌트)
4. [Design Patterns / 디자인 패턴](#design-patterns--디자인-패턴)
5. [Internal Implementation / 내부 구현](#internal-implementation--내부-구현)
6. [Adding New Features / 새 기능 추가](#adding-new-features--새-기능-추가)
7. [Testing Guide / 테스트 가이드](#testing-guide--테스트-가이드)
8. [Performance / 성능](#performance--성능)
9. [Contributing Guidelines / 기여 가이드라인](#contributing-guidelines--기여-가이드라인)
10. [Code Style / 코드 스타일](#code-style--코드-스타일)

---

## Architecture Overview / 아키텍처 개요

### Design Principles / 설계 원칙

The Sliceutil package follows these core design principles:

Sliceutil 패키지는 다음과 같은 핵심 설계 원칙을 따릅니다:

1. **Extreme Simplicity / 극도의 간결함**: Reduce 20+ lines of repetitive code to 1 line / 20줄 이상의 반복 코드를 1줄로 축소
2. **Type Safety / 타입 안전성**: Leverage Go 1.18+ generics for compile-time type checking / Go 1.18+ 제네릭을 활용한 컴파일 타임 타입 검사
3. **Immutability / 불변성**: All operations return new slices, never modifying originals / 모든 작업은 새 슬라이스를 반환하며 원본을 절대 수정하지 않음
4. **Zero Dependencies / 제로 의존성**: Only uses Go standard library (except golang.org/x/exp) / Go 표준 라이브러리만 사용 (golang.org/x/exp 제외)
5. **Functional Programming / 함수형 프로그래밍**: Higher-order functions (Map, Filter, Reduce, Scan) / 고차 함수 (Map, Filter, Reduce, Scan)
6. **Performance / 성능**: Efficient algorithms with minimal allocations / 최소 할당으로 효율적인 알고리즘
7. **Comprehensive Coverage / 포괄적인 커버리지**: 95 functions across 14 categories / 14개 카테고리에 걸쳐 95개 함수

### High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────────────────┐
│                        Sliceutil Package                         │
│                     github.com/arkd0ng/go-utils/sliceutil        │
└─────────────────────────────────────────────────────────────────┘
                                  │
                ┌─────────────────┴─────────────────┐
                │                                   │
        ┌───────▼────────┐                 ┌───────▼────────┐
        │  Type System   │                 │   Functions    │
        │                │                 │   (95 total)   │
        │ - Number       │                 │                │
        │ - Ordered      │                 │ 14 Categories: │
        │ - comparable   │                 │                │
        │ - any          │                 │ 1. Basic       │
        │                │                 │ 2. Transform   │
        │ - DiffResult   │                 │ 3. Aggregate   │
        │   (struct)     │                 │ 4. Slicing     │
        └────────────────┘                 │ 5. Set Ops     │
                                          │ 6. Sorting     │
                                          │ 7. Predicates  │
                                          │ 8. Utilities   │
                                          │ 9. Statistics  │
                                          │ 10. Diff/Comp  │
                                          │ 11. Index Ops  │
                                          │ 12. Conditional│
                                          │ 13. Advanced   │
                                          │ 14. Combinator │
                                          └────────────────┘
```

### Component Interaction / 컴포넌트 상호작용

```
User Code / 사용자 코드
    ↓
┌───────────────────────────┐
│   Public API Functions    │  ← Type-safe generic functions
│   (95 functions)          │    타입 안전 제네릭 함수
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Type Constraints        │  ← Number, Ordered, comparable
│   (Generic Validation)    │    제네릭 검증
│   + DiffResult struct     │    + DiffResult 구조체
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Core Algorithm          │  ← Efficient slice operations
│   Implementation          │    효율적인 슬라이스 작업
│   + Statistical methods   │    + 통계 메서드
│   + Combinatorial algo    │    + 조합 알고리즘
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Go Standard Library     │  ← slices, sort, math/rand, errors
│   + golang.org/x/exp      │    + golang.org/x/exp/constraints
└───────────────────────────┘
```

---

## Package Structure / 패키지 구조

### File Organization / 파일 구성

```
sliceutil/
├── sliceutil.go           # Package documentation, types, constraints
│                          # 패키지 문서, 타입, 제약조건
├── sliceutil_test.go      # Package-level tests (version, import)
│                          # 패키지 레벨 테스트 (버전, 임포트)
├── basic.go               # Basic operations (10 functions)
│                          # 기본 작업 (10개 함수)
├── basic_test.go          # Tests for basic operations
│                          # 기본 작업 테스트
├── transform.go           # Transformation functions (8 functions)
│                          # 변환 함수 (8개 함수)
├── transform_test.go      # Tests for transformation functions
│                          # 변환 함수 테스트
├── aggregate.go           # Aggregation functions (7 functions)
│                          # 집계 함수 (7개 함수)
├── aggregate_test.go      # Tests for aggregation functions
│                          # 집계 함수 테스트
├── slice.go               # Slicing functions (7 functions)
│                          # 슬라이싱 함수 (7개 함수)
├── slice_test.go          # Tests for slicing functions
│                          # 슬라이싱 함수 테스트
├── set.go                 # Set operations (6 functions)
│                          # 집합 작업 (6개 함수)
├── set_test.go            # Tests for set operations
│                          # 집합 작업 테스트
├── sort.go                # Sorting functions (5 functions)
│                          # 정렬 함수 (5개 함수)
├── sort_test.go           # Tests for sorting functions
│                          # 정렬 함수 테스트
├── predicate.go           # Predicate functions (6 functions)
│                          # 조건자 함수 (6개 함수)
├── predicate_test.go      # Tests for predicate functions
│                          # 조건자 함수 테스트
├── util.go                # Utility functions (11 functions)
│                          # 유틸리티 함수 (11개 함수)
├── util_test.go           # Tests for utility functions
│                          # 유틸리티 함수 테스트
├── statistics.go          # Statistical functions (8 functions)
│                          # 통계 함수 (8개 함수)
├── statistics_test.go     # Tests for statistical functions
│                          # 통계 함수 테스트
├── diff.go                # Diff/comparison functions (4 functions)
│                          # 차이/비교 함수 (4개 함수)
├── diff_test.go           # Tests for diff operations
│                          # 차이 작업 테스트
├── index.go               # Index-based operations (3 functions)
│                          # 인덱스 기반 작업 (3개 함수)
├── index_test.go          # Tests for index operations
│                          # 인덱스 작업 테스트
├── conditional.go         # Conditional operations (3 functions)
│                          # 조건부 작업 (3개 함수)
├── conditional_test.go    # Tests for conditional operations
│                          # 조건부 작업 테스트
├── advanced.go            # Advanced functional operations (4 functions)
│                          # 고급 함수형 작업 (4개 함수)
├── advanced_test.go       # Tests for advanced operations
│                          # 고급 작업 테스트
├── combinatorial.go       # Combinatorial operations (2 functions)
│                          # 조합 작업 (2개 함수)
├── combinatorial_test.go  # Tests for combinatorial operations
│                          # 조합 작업 테스트
└── README.md              # Package README
                           # 패키지 README
```

### File Responsibilities / 파일별 책임

| File / 파일 | Purpose / 목적 | Functions / 함수 | Lines / 줄 수 |
|-------------|---------------|-----------------|--------------|
| `sliceutil.go` | Package documentation, type definitions / 패키지 문서, 타입 정의 | Types: Number, Ordered, DiffResult | ~150 |
| `basic.go` | Search, find, count operations / 검색, 찾기, 개수 세기 작업 | 10 functions | ~300 |
| `transform.go` | Map, filter, unique, reverse / Map, 필터, 고유값, 역순 | 8 functions | ~250 |
| `aggregate.go` | Reduce, sum, min, max, groupby / Reduce, 합계, 최소, 최대, 그룹화 | 7 functions | ~220 |
| `slice.go` | Chunk, take, drop, partition / 청크, 가져오기, 제거, 파티션 | 7 functions | ~200 |
| `set.go` | Union, intersection, difference / 합집합, 교집합, 차집합 | 6 functions | ~180 |
| `sort.go` | Sort, sort by, is sorted / 정렬, 정렬 기준, 정렬 확인 | 5 functions | ~150 |
| `predicate.go` | All, any, none, equal / 모두, 어느, 없음, 동일 | 6 functions | ~150 |
| `util.go` | ForEach, join, shuffle, zip / ForEach, 결합, 섞기, 압축 | 11 functions | ~300 |
| `statistics.go` | Statistical operations / 통계 작업 | 8 functions | ~330 |
| `diff.go` | Diff and comparison operations / 차이 및 비교 작업 | 4 functions + DiffResult | ~190 |
| `index.go` | Index-based operations / 인덱스 기반 작업 | 3 functions | ~95 |
| `conditional.go` | Conditional transformations / 조건부 변환 | 3 functions | ~95 |
| `advanced.go` | Advanced functional programming / 고급 함수형 프로그래밍 | 4 functions | ~125 |
| `combinatorial.go` | Permutations and combinations / 순열 및 조합 | 2 functions | ~110 |
| `*_test.go` | Comprehensive tests for each file / 각 파일에 대한 종합 테스트 | Test functions | ~4,070 total |

**Total Package Size / 전체 패키지 크기**: ~6,915 lines (implementation + tests) / ~6,915줄 (구현 + 테스트)

---

## Core Components / 핵심 컴포넌트

### 1. Type Constraints / 타입 제약조건

**Location / 위치**: `sliceutil/sliceutil.go`, `sliceutil/statistics.go`

The package defines type constraints and structured types for generics:

패키지는 제네릭을 위한 타입 제약 및 구조화된 타입을 정의합니다:

```go
// Number constraint for numeric operations (statistical, arithmetic)
// 숫자 작업을 위한 Number 제약 (통계, 산술)
// Uses golang.org/x/exp/constraints for standard numeric types
// 표준 숫자 타입을 위해 golang.org/x/exp/constraints 사용
type Number interface {
	constraints.Integer | constraints.Float
}

// Ordered constraint for sorting operations
// 정렬 작업을 위한 Ordered 제약
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

// DiffResult holds the result of a diff operation between two slices
// DiffResult는 두 슬라이스 간의 차이 연산 결과를 저장합니다
type DiffResult[T any] struct {
	Added     []T // Elements in new but not in old / 새 슬라이스에만 있는 요소
	Removed   []T // Elements in old but not in new / 이전 슬라이스에만 있는 요소
	Unchanged []T // Elements in both old and new / 양쪽 모두에 있는 요소
}
```

**Design Rationale / 설계 근거**:
- `Number`: Used for statistical and mathematical operations (Median, Variance, Sum, Average) / 통계 및 수학적 작업에 사용 (Median, Variance, Sum, Average)
  - Leverages `golang.org/x/exp/constraints` for standard numeric type definitions / 표준 숫자 타입 정의를 위해 `golang.org/x/exp/constraints` 활용
  - Supports both integers and floating-point types / 정수 및 부동 소수점 타입 모두 지원
- `Ordered`: Used for comparison and sorting (Sort, IsSorted) / 비교 및 정렬에 사용 (Sort, IsSorted)
- `comparable`: Built-in Go constraint for equality checks (Contains, Unique, Diff) / 동등성 검사를 위한 내장 Go 제약 (Contains, Unique, Diff)
- `any`: Used for type-agnostic operations (Map, Filter, ForEach, Scan) / 타입에 구애받지 않는 작업에 사용 (Map, Filter, ForEach, Scan)
- `DiffResult[T]`: Structured type for returning comprehensive diff results / 포괄적인 차이 결과를 반환하기 위한 구조화된 타입
  - Encapsulates added, removed, and unchanged elements / 추가, 제거, 변경되지 않은 요소를 캡슐화
  - Type-safe with generics / 제네릭으로 타입 안전

### 2. Function Categories / 함수 카테고리

#### Category 1: Basic Operations (10 functions) / 기본 작업 (10개 함수)

**File / 파일**: `sliceutil/basic.go`

**Purpose / 목적**: Fundamental slice operations for searching, finding, and counting.

기본적인 슬라이스 작업: 검색, 찾기, 개수 세기.

**Key Functions / 주요 함수**:
- `Contains[T comparable](slice []T, element T) bool`
- `IndexOf[T comparable](slice []T, element T) int`
- `Find[T any](slice []T, predicate func(T) bool) (T, bool)`
- `Count[T any](slice []T, predicate func(T) bool) int`

**Implementation Pattern / 구현 패턴**:
```go
// Simple linear search pattern
// 간단한 선형 검색 패턴
func Contains[T comparable](slice []T, element T) bool {
    for _, item := range slice {
        if item == element {
            return true
        }
    }
    return false
}
```

**Time Complexity / 시간 복잡도**: O(n) for most operations / 대부분의 작업에 대해 O(n)

#### Category 2: Transformation (8 functions) / 변환 (8개 함수)

**File / 파일**: `sliceutil/transform.go`

**Purpose / 목적**: Transform slices into different forms.

슬라이스를 다양한 형태로 변환.

**Key Functions / 주요 함수**:
- `Map[T any, R any](slice []T, mapper func(T) R) []R`
- `Filter[T any](slice []T, predicate func(T) bool) []T`
- `Unique[T comparable](slice []T) []T`
- `Flatten[T any](slice [][]T) []T`

**Implementation Pattern / 구현 패턴**:
```go
// Map: Transform each element
// Map: 각 요소 변환
func Map[T any, R any](slice []T, mapper func(T) R) []R {
    if len(slice) == 0 {
        return []R{}
    }

    result := make([]R, len(slice))
    for i, item := range slice {
        result[i] = mapper(item)
    }
    return result
}
```

**Time Complexity / 시간 복잡도**: O(n) typical, O(n²) for Unique with large datasets / 일반적으로 O(n), 대용량 데이터셋의 Unique는 O(n²)

#### Category 3: Aggregation (7 functions) / 집계 (7개 함수)

**File / 파일**: `sliceutil/aggregate.go`

**Purpose / 목적**: Aggregate data from slices.

슬라이스에서 데이터 집계.

**Key Functions / 주요 함수**:
- `Reduce[T any, R any](slice []T, initial R, reducer func(R, T) R) R`
- `Sum[T Number](slice []T) T`
- `GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T`

**Implementation Pattern / 구현 패턴**:
```go
// Reduce: Flexible aggregation pattern
// Reduce: 유연한 집계 패턴
func Reduce[T any, R any](slice []T, initial R, reducer func(R, T) R) R {
    result := initial
    for _, item := range slice {
        result = reducer(result, item)
    }
    return result
}

// Sum: Specialized aggregation
// Sum: 특수화된 집계
func Sum[T Number](slice []T) T {
    var sum T
    for _, item := range slice {
        sum += item
    }
    return sum
}
```

**Time Complexity / 시간 복잡도**: O(n) for most, O(n) space for GroupBy / 대부분 O(n), GroupBy는 O(n) 공간

#### Category 4: Slicing (7 functions) / 슬라이싱 (7개 함수)

**File / 파일**: `sliceutil/slice.go`

**Purpose / 목적**: Extract portions of slices.

슬라이스의 일부 추출.

**Key Functions / 주요 함수**:
- `Chunk[T any](slice []T, size int) [][]T`
- `Take[T any](slice []T, n int) []T`
- `Partition[T any](slice []T, predicate func(T) bool) ([]T, []T)`

**Implementation Pattern / 구현 패턴**:
```go
// Chunk: Split into fixed-size pieces
// Chunk: 고정 크기 조각으로 분할
func Chunk[T any](slice []T, size int) [][]T {
    if size <= 0 || len(slice) == 0 {
        return [][]T{}
    }

    chunks := make([][]T, 0, (len(slice)+size-1)/size)
    for i := 0; i < len(slice); i += size {
        end := i + size
        if end > len(slice) {
            end = len(slice)
        }
        chunks = append(chunks, slice[i:end])
    }
    return chunks
}
```

**Time Complexity / 시간 복잡도**: O(n) typical / 일반적으로 O(n)

#### Category 5: Set Operations (6 functions) / 집합 작업 (6개 함수)

**File / 파일**: `sliceutil/set.go`

**Purpose / 목적**: Set-theory operations on slices.

슬라이스에 대한 집합론 작업.

**Key Functions / 주요 함수**:
- `Union[T comparable](slice1, slice2 []T) []T`
- `Intersection[T comparable](slice1, slice2 []T) []T`
- `Difference[T comparable](slice1, slice2 []T) []T`

**Implementation Pattern / 구현 패턴**:
```go
// Union: Combine unique elements from both slices
// Union: 두 슬라이스의 고유 요소 결합
func Union[T comparable](slice1, slice2 []T) []T {
    seen := make(map[T]bool)
    result := make([]T, 0, len(slice1)+len(slice2))

    for _, item := range slice1 {
        if !seen[item] {
            seen[item] = true
            result = append(result, item)
        }
    }

    for _, item := range slice2 {
        if !seen[item] {
            seen[item] = true
            result = append(result, item)
        }
    }

    return result
}
```

**Time Complexity / 시간 복잡도**: O(n + m) with hash map / 해시 맵 사용 시 O(n + m)

#### Category 6: Sorting (5 functions) / 정렬 (5개 함수)

**File / 파일**: `sliceutil/sort.go`

**Purpose / 목적**: Sort slices and check sort status.

슬라이스 정렬 및 정렬 상태 확인.

**Key Functions / 주요 함수**:
- `Sort[T Ordered](slice []T) []T`
- `SortBy[T any](slice []T, less func(a, b T) bool) []T`
- `IsSorted[T Ordered](slice []T) bool`

**Implementation Pattern / 구현 패턴**:
```go
// Sort: Use standard library sort
// Sort: 표준 라이브러리 정렬 사용
func Sort[T Ordered](slice []T) []T {
    if len(slice) == 0 {
        return []T{}
    }

    result := make([]T, len(slice))
    copy(result, slice)
    sort.Slice(result, func(i, j int) bool {
        return result[i] < result[j]
    })
    return result
}
```

**Time Complexity / 시간 복잡도**: O(n log n) for sorting / 정렬은 O(n log n)

#### Category 7: Predicates (6 functions) / 조건자 (6개 함수)

**File / 파일**: `sliceutil/predicate.go`

**Purpose / 목적**: Test conditions on slices.

슬라이스에 대한 조건 테스트.

**Key Functions / 주요 함수**:
- `All[T any](slice []T, predicate func(T) bool) bool`
- `Any[T any](slice []T, predicate func(T) bool) bool`
- `Equal[T comparable](slice1, slice2 []T) bool`

**Implementation Pattern / 구현 패턴**:
```go
// All: Check if all elements satisfy predicate
// All: 모든 요소가 조건자를 만족하는지 확인
func All[T any](slice []T, predicate func(T) bool) bool {
    for _, item := range slice {
        if !predicate(item) {
            return false
        }
    }
    return true
}

// Short-circuit evaluation for efficiency
// 효율성을 위한 단락 평가
```

**Time Complexity / 시간 복잡도**: O(n) worst case, but often short-circuits / 최악의 경우 O(n), 하지만 종종 단락 평가

#### Category 8: Utilities (11 functions) / 유틸리티 (11개 함수)

**File / 파일**: `sliceutil/util.go`

**Purpose / 목적**: Miscellaneous utility functions.

기타 유틸리티 함수.

**Key Functions / 주요 함수**:
- `ForEach[T any](slice []T, fn func(T))`
- `Join[T any](slice []T, separator string) string`
- `Shuffle[T any](slice []T) []T`
- `Zip[T, U any](slice1 []T, slice2 []U) [][2]interface{}`

**Implementation Pattern / 구현 패턴**:
```go
// ForEach: Execute function for each element
// ForEach: 각 요소에 대해 함수 실행
func ForEach[T any](slice []T, fn func(T)) {
    for _, item := range slice {
        fn(item)
    }
}

// Shuffle: Fisher-Yates algorithm
// Shuffle: Fisher-Yates 알고리즘
func Shuffle[T any](slice []T) []T {
    if len(slice) == 0 {
        return []T{}
    }

    result := make([]T, len(slice))
    copy(result, slice)

    for i := len(result) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        result[i], result[j] = result[j], result[i]
    }

    return result
}
```

**Time Complexity / 시간 복잡도**: Varies by function / 함수에 따라 다름

#### Category 9: Statistics (8 functions) / 통계 (8개 함수)

**File / 파일**: `sliceutil/statistics.go`

**Purpose / 목적**: Statistical analysis and frequency operations on numeric slices.

숫자 슬라이스에 대한 통계 분석 및 빈도 작업.

**Key Functions / 주요 함수**:
- `Median[T Number](slice []T) (float64, error)`
- `Mode[T comparable](slice []T) (T, error)`
- `Frequencies[T comparable](slice []T) map[T]int`
- `Percentile[T Number](slice []T, p float64) (float64, error)`
- `StandardDeviation[T Number](slice []T) (float64, error)`
- `Variance[T Number](slice []T) (float64, error)`
- `MostCommon[T comparable](slice []T, n int) []T`
- `LeastCommon[T comparable](slice []T, n int) []T`

**Implementation Pattern / 구현 패턴**:
```go
// Median: Sort and find middle value(s)
// Median: 정렬하고 중간 값 찾기
func Median[T Number](slice []T) (float64, error) {
    if len(slice) == 0 {
        return 0, errors.New("cannot calculate median of empty slice")
    }

    // Create sorted copy to preserve immutability
    // 불변성을 유지하기 위해 정렬된 복사본 생성
    sorted := Sort(slice)
    length := len(sorted)

    if length%2 == 0 {
        // Even: average of two middle values
        // 짝수: 두 중간 값의 평균
        mid1 := sorted[length/2-1]
        mid2 := sorted[length/2]
        return (float64(mid1) + float64(mid2)) / 2.0, nil
    }

    // Odd: middle value
    // 홀수: 중간 값
    return float64(sorted[length/2]), nil
}

// Variance: Calculate population variance
// Variance: 모집단 분산 계산
func Variance[T Number](slice []T) (float64, error) {
    if len(slice) == 0 {
        return 0, errors.New("cannot calculate variance of empty slice")
    }

    mean := Average(slice)
    var sumSquaredDiff float64
    for _, v := range slice {
        diff := float64(v) - mean
        sumSquaredDiff += diff * diff
    }

    return sumSquaredDiff / float64(len(slice)), nil
}
```

**Design Decisions / 설계 결정**:
- Uses `Number` constraint defined in `statistics.go` with `golang.org/x/exp/constraints` / `golang.org/x/exp/constraints`를 사용한 `statistics.go`의 `Number` 제약 사용
- Returns `error` for empty slices to ensure safe statistical calculations / 안전한 통계 계산을 보장하기 위해 빈 슬라이스에 대해 `error` 반환
- Population formulas (divide by N) for variance/stddev, not sample formulas / 분산/표준편차에 대해 표본 공식이 아닌 모집단 공식 사용 (N으로 나눔)
- Linear interpolation for percentile calculation / 백분위수 계산에 선형 보간 사용

**Time Complexity / 시간 복잡도**: O(n log n) for Median/Percentile (sorting), O(n) for others / Median/Percentile은 O(n log n) (정렬), 기타는 O(n)

#### Category 10: Diff/Comparison (4 functions) / 차이/비교 (4개 함수)

**File / 파일**: `sliceutil/diff.go`

**Purpose / 목적**: Compare slices and find differences between them.

슬라이스를 비교하고 차이점을 찾습니다.

**Key Functions / 주요 함수**:
- `Diff[T comparable](old, new []T) DiffResult[T]`
- `DiffBy[T any, K comparable](old, new []T, keyFunc func(T) K) DiffResult[T]`
- `EqualUnordered[T comparable](a, b []T) bool`
- `HasDuplicates[T comparable](slice []T) bool`

**Implementation Pattern / 구현 패턴**:
```go
// Diff: Compare two slices using hash maps
// Diff: 해시 맵을 사용하여 두 슬라이스 비교
func Diff[T comparable](old, new []T) DiffResult[T] {
    // Build hash sets for O(1) lookup
    // O(1) 조회를 위한 해시 집합 구축
    oldSet := make(map[T]bool)
    for _, v := range old {
        oldSet[v] = true
    }

    newSet := make(map[T]bool)
    for _, v := range new {
        newSet[v] = true
    }

    var added, removed, unchanged []T

    // Find added elements (in new but not in old)
    // 추가된 요소 찾기 (new에는 있지만 old에는 없음)
    for _, v := range new {
        if !oldSet[v] {
            added = append(added, v)
        } else {
            unchanged = append(unchanged, v)
        }
    }

    // Find removed elements (in old but not in new)
    // 제거된 요소 찾기 (old에는 있지만 new에는 없음)
    for _, v := range old {
        if !newSet[v] {
            removed = append(removed, v)
        }
    }

    // Remove duplicates from unchanged
    // unchanged에서 중복 제거
    unchanged = Unique(unchanged)

    return DiffResult[T]{
        Added:     added,
        Removed:   removed,
        Unchanged: unchanged,
    }
}
```

**Design Decisions / 설계 결정**:
- `DiffResult[T]` struct encapsulates three slices for clarity / 명확성을 위해 `DiffResult[T]` 구조체가 세 개의 슬라이스를 캡슐화
- `DiffBy` allows comparing complex types using key function / `DiffBy`는 키 함수를 사용하여 복잡한 타입 비교 가능
- Hash map approach ensures O(n + m) complexity / 해시 맵 접근 방식은 O(n + m) 복잡도 보장

**Time Complexity / 시간 복잡도**: O(n + m) for Diff/DiffBy, O(n) for EqualUnordered/HasDuplicates / Diff/DiffBy는 O(n + m), EqualUnordered/HasDuplicates는 O(n)

#### Category 11: Index Operations (3 functions) / 인덱스 작업 (3개 함수)

**File / 파일**: `sliceutil/index.go`

**Purpose / 목적**: Work with slice elements using index-based operations.

인덱스 기반 작업을 사용하여 슬라이스 요소 작업.

**Key Functions / 주요 함수**:
- `FindIndices[T any](slice []T, predicate func(T) bool) []int`
- `AtIndices[T any](slice []T, indices []int) []T`
- `RemoveIndices[T any](slice []T, indices []int) []T`

**Implementation Pattern / 구현 패턴**:
```go
// FindIndices: Collect all matching indices
// FindIndices: 모든 일치하는 인덱스 수집
func FindIndices[T any](slice []T, predicate func(T) bool) []int {
    var indices []int
    for i, v := range slice {
        if predicate(v) {
            indices = append(indices, i)
        }
    }
    return indices
}

// RemoveIndices: Skip specified indices efficiently
// RemoveIndices: 지정된 인덱스를 효율적으로 건너뜀
func RemoveIndices[T any](slice []T, indices []int) []T {
    if len(indices) == 0 {
        return Clone(slice)
    }

    // Create hash set for O(1) lookup
    // O(1) 조회를 위한 해시 집합 생성
    toRemove := make(map[int]bool)
    for _, idx := range indices {
        if idx >= 0 && idx < len(slice) {
            toRemove[idx] = true
        }
    }

    // Build result, skipping marked indices
    // 표시된 인덱스를 건너뛰고 결과 구축
    result := make([]T, 0, len(slice)-len(toRemove))
    for i, v := range slice {
        if !toRemove[i] {
            result = append(result, v)
        }
    }

    return result
}
```

**Design Decisions / 설계 결정**:
- Out-of-bounds indices are silently skipped (no panic) / 범위를 벗어난 인덱스는 자동으로 건너뜀 (패닉 없음)
- Negative indices are not supported / 음수 인덱스는 지원되지 않음
- Hash map for efficient index removal / 효율적인 인덱스 제거를 위한 해시 맵

**Time Complexity / 시간 복잡도**: O(n) for all operations / 모든 작업에 대해 O(n)

#### Category 12: Conditional Operations (3 functions) / 조건부 작업 (3개 함수)

**File / 파일**: `sliceutil/conditional.go`

**Purpose / 목적**: Conditional transformations and replacements.

조건부 변환 및 교체.

**Key Functions / 주요 함수**:
- `ReplaceIf[T any](slice []T, predicate func(T) bool, newValue T) []T`
- `ReplaceAll[T comparable](slice []T, oldValue, newValue T) []T`
- `UpdateWhere[T any](slice []T, predicate func(T) bool, updater func(T) T) []T`

**Implementation Pattern / 구현 패턴**:
```go
// ReplaceIf: Conditional replacement with predicate
// ReplaceIf: 조건자를 사용한 조건부 교체
func ReplaceIf[T any](slice []T, predicate func(T) bool, newValue T) []T {
    result := make([]T, len(slice))
    for i, v := range slice {
        if predicate(v) {
            result[i] = newValue
        } else {
            result[i] = v
        }
    }
    return result
}

// UpdateWhere: Conditional update with updater function
// UpdateWhere: updater 함수를 사용한 조건부 업데이트
func UpdateWhere[T any](slice []T, predicate func(T) bool, updater func(T) T) []T {
    result := make([]T, len(slice))
    for i, v := range slice {
        if predicate(v) {
            result[i] = updater(v)
        } else {
            result[i] = v
        }
    }
    return result
}
```

**Design Decisions / 설계 결정**:
- Immutable operations (return new slice) / 불변 작업 (새 슬라이스 반환)
- `UpdateWhere` is more flexible than `ReplaceIf` (custom transformation) / `UpdateWhere`는 `ReplaceIf`보다 더 유연함 (사용자 정의 변환)
- Works well with structs and complex types / 구조체 및 복잡한 타입과 잘 작동

**Time Complexity / 시간 복잡도**: O(n) for all operations / 모든 작업에 대해 O(n)

#### Category 13: Advanced Functional (4 functions) / 고급 함수형 (4개 함수)

**File / 파일**: `sliceutil/advanced.go`

**Purpose / 목적**: Advanced functional programming patterns.

고급 함수형 프로그래밍 패턴.

**Key Functions / 주요 함수**:
- `Scan[T any](slice []T, initial T, accumulator func(T, T) T) []T`
- `ZipWith[T, U, R any](a []T, b []U, zipper func(T, U) R) []R`
- `RotateLeft[T any](slice []T, n int) []T`
- `RotateRight[T any](slice []T, n int) []T`

**Implementation Pattern / 구현 패턴**:
```go
// Scan: Like Reduce but returns all intermediate results
// Scan: Reduce와 비슷하지만 모든 중간 결과 반환
func Scan[T any](slice []T, initial T, accumulator func(T, T) T) []T {
    result := make([]T, len(slice)+1)
    result[0] = initial

    acc := initial
    for i, v := range slice {
        acc = accumulator(acc, v)
        result[i+1] = acc
    }

    return result
}

// ZipWith: Combine two slices element-wise with custom function
// ZipWith: 사용자 정의 함수로 두 슬라이스를 요소별로 결합
func ZipWith[T, U, R any](a []T, b []U, zipper func(T, U) R) []R {
    length := len(a)
    if len(b) < length {
        length = len(b)
    }

    result := make([]R, length)
    for i := 0; i < length; i++ {
        result[i] = zipper(a[i], b[i])
    }

    return result
}

// RotateLeft: Circular rotation to the left
// RotateLeft: 왼쪽으로 원형 회전
func RotateLeft[T any](slice []T, n int) []T {
    if len(slice) == 0 {
        return []T{}
    }

    // Normalize n to be within [0, len)
    // n을 [0, len) 범위로 정규화
    length := len(slice)
    n = ((n % length) + length) % length

    result := make([]T, length)
    for i := 0; i < length; i++ {
        result[i] = slice[(i+n)%length]
    }

    return result
}
```

**Design Decisions / 설계 결정**:
- `Scan` is useful for cumulative operations (running sum, factorial) / `Scan`은 누적 작업에 유용 (누적 합, 팩토리얼)
- `ZipWith` generalizes Zip to any combining function / `ZipWith`는 Zip을 임의의 결합 함수로 일반화
- Rotation functions handle negative n correctly / 회전 함수는 음수 n을 올바르게 처리

**Time Complexity / 시간 복잡도**: O(n) for Scan/Rotate, O(min(len(a), len(b))) for ZipWith / Scan/Rotate는 O(n), ZipWith는 O(min(len(a), len(b)))

#### Category 14: Combinatorial (2 functions) / 조합 (2개 함수)

**File / 파일**: `sliceutil/combinatorial.go`

**Purpose / 목적**: Generate permutations and combinations.

순열 및 조합 생성.

**Key Functions / 주요 함수**:
- `Permutations[T any](slice []T) [][]T`
- `Combinations[T any](slice []T, k int) [][]T`

**Implementation Pattern / 구현 패턴**:
```go
// Permutations: Generate all permutations using Heap's algorithm
// Permutations: Heap의 알고리즘을 사용하여 모든 순열 생성
func Permutations[T any](slice []T) [][]T {
    if len(slice) == 0 {
        return [][]T{{}}
    }

    result := [][]T{}
    permute(slice, 0, &result)
    return result
}

// Helper function using Heap's algorithm
// Heap의 알고리즘을 사용하는 헬퍼 함수
func permute[T any](slice []T, k int, result *[][]T) {
    if k == len(slice)-1 {
        perm := make([]T, len(slice))
        copy(perm, slice)
        *result = append(*result, perm)
        return
    }

    for i := k; i < len(slice); i++ {
        slice[k], slice[i] = slice[i], slice[k]
        permute(slice, k+1, result)
        slice[k], slice[i] = slice[i], slice[k] // backtrack
    }
}

// Combinations: Generate all k-combinations recursively
// Combinations: 재귀적으로 모든 k-조합 생성
func Combinations[T any](slice []T, k int) [][]T {
    if k < 0 || k > len(slice) {
        return [][]T{}
    }
    if k == 0 {
        return [][]T{{}}
    }
    if k == len(slice) {
        return [][]T{append([]T{}, slice...)}
    }

    result := [][]T{}
    combine(slice, k, 0, []T{}, &result)
    return result
}
```

**Design Decisions / 설계 결정**:
- Uses Heap's algorithm for permutations (efficient, in-place) / 순열에 Heap의 알고리즘 사용 (효율적, 제자리)
- Recursive approach for combinations / 조합에 재귀 접근 방식
- **WARNING**: Complexity grows factorially/exponentially / **경고**: 복잡도가 팩토리얼/지수적으로 증가
  - Permutations: O(n!) - n=10 produces 3,628,800 results / 순열: O(n!) - n=10일 때 3,628,800개 결과 생성
  - Combinations: O(C(n,k)) - C(20,10) produces 184,756 results / 조합: O(C(n,k)) - C(20,10)일 때 184,756개 결과 생성

**Time Complexity / 시간 복잡도**: O(n!) for Permutations, O(C(n,k)) for Combinations / 순열은 O(n!), 조합은 O(C(n,k))

---

## Design Patterns / 디자인 패턴

### Pattern 1: Generic Type Parameters / 제네릭 타입 매개변수

**Purpose / 목적**: Provide type-safe operations without code duplication.

코드 중복 없이 타입 안전 작업 제공.

**Implementation / 구현**:
```go
// Single generic type parameter
// 단일 제네릭 타입 매개변수
func Filter[T any](slice []T, predicate func(T) bool) []T {
    // Type T is inferred from the slice argument
    // 타입 T는 슬라이스 인수에서 추론됨
}

// Multiple generic type parameters (input and output types differ)
// 여러 제네릭 타입 매개변수 (입력 및 출력 타입이 다름)
func Map[T any, R any](slice []T, mapper func(T) R) []R {
    // T is input type, R is output type
    // T는 입력 타입, R은 출력 타입
}

// Constrained generic type (only numeric types)
// 제약된 제네릭 타입 (숫자 타입만)
func Sum[T Number](slice []T) T {
    // T must satisfy Number constraint
    // T는 Number 제약을 만족해야 함
}
```

**Benefits / 이점**:
- Type safety at compile time / 컴파일 타임에 타입 안전성
- No runtime type assertions / 런타임 타입 단언 없음
- Code reusability / 코드 재사용성
- Better IDE support and autocomplete / 더 나은 IDE 지원 및 자동 완성

### Pattern 2: Functional Programming with Higher-Order Functions / 고차 함수를 사용한 함수형 프로그래밍

**Purpose / 목적**: Enable functional programming patterns in Go.

Go에서 함수형 프로그래밍 패턴 활성화.

**Implementation / 구현**:
```go
// Higher-order function: takes function as parameter
// 고차 함수: 함수를 매개변수로 받음
func Map[T any, R any](slice []T, mapper func(T) R) []R {
    result := make([]R, len(slice))
    for i, item := range slice {
        result[i] = mapper(item)
    }
    return result
}

// Usage: pass lambda/anonymous function
// 사용: 람다/익명 함수 전달
doubled := Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
```

**Common Higher-Order Functions / 일반적인 고차 함수**:
- `Map`: Transform elements / 요소 변환
- `Filter`: Select elements / 요소 선택
- `Reduce`: Aggregate elements / 요소 집계
- `Find`: Search with predicate / 조건자로 검색
- `All/Any/None`: Test conditions / 조건 테스트

**Benefits / 이점**:
- Declarative code style / 선언적 코드 스타일
- Composable operations / 구성 가능한 작업
- Less boilerplate / 보일러플레이트 감소

### Pattern 3: Immutability by Return Value / 반환값을 통한 불변성

**Purpose / 목적**: Ensure original slices are never modified.

원본 슬라이스가 절대 수정되지 않도록 보장.

**Implementation / 구현**:
```go
// Always create new slice instead of modifying original
// 원본을 수정하는 대신 항상 새 슬라이스 생성
func Reverse[T any](slice []T) []T {
    if len(slice) == 0 {
        return []T{}
    }

    // Create new slice
    // 새 슬라이스 생성
    result := make([]T, len(slice))

    // Copy in reverse order
    // 역순으로 복사
    for i := range slice {
        result[i] = slice[len(slice)-1-i]
    }

    return result
}

// Original slice remains unchanged
// 원본 슬라이스는 변경되지 않음
original := []int{1, 2, 3}
reversed := Reverse(original)
// original: [1 2 3]
// reversed: [3 2 1]
```

**Benefits / 이점**:
- No side effects / 부작용 없음
- Safe for concurrent use / 동시 사용에 안전
- Easier to reason about code / 코드를 이해하기 쉬움
- Prevents unexpected bugs / 예기치 않은 버그 방지

### Pattern 4: Empty Slice Handling / 빈 슬라이스 처리

**Purpose / 목적**: Consistent behavior for edge cases.

엣지 케이스에 대한 일관된 동작.

**Implementation / 구현**:
```go
// Pattern 1: Return empty slice (not nil)
// 패턴 1: 빈 슬라이스 반환 (nil 아님)
func Filter[T any](slice []T, predicate func(T) bool) []T {
    if len(slice) == 0 {
        return []T{} // Empty slice, not nil
    }
    // ... implementation
}

// Pattern 2: Return zero value and false for not found
// 패턴 2: 찾지 못한 경우 제로값과 false 반환
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
    for _, item := range slice {
        if predicate(item) {
            return item, true
        }
    }
    var zero T
    return zero, false // Not found
}

// Pattern 3: Return sensible defaults for aggregation
// 패턴 3: 집계에 대해 합리적인 기본값 반환
func Average[T Number](slice []T) float64 {
    if len(slice) == 0 {
        return 0.0 // Reasonable default
    }
    return float64(Sum(slice)) / float64(len(slice))
}
```

**Benefits / 이점**:
- Predictable behavior / 예측 가능한 동작
- Avoids panics / 패닉 방지
- Consistent API / 일관된 API

### Pattern 5: Builder Pattern for Chaining (Future Enhancement) / 체이닝을 위한 빌더 패턴 (향후 개선)

**Purpose / 목적**: Enable fluent API for chaining operations.

작업 체이닝을 위한 유창한 API 활성화.

**Potential Implementation / 잠재적 구현**:
```go
// Future enhancement: Slice builder for method chaining
// 향후 개선: 메서드 체이닝을 위한 슬라이스 빌더
type SliceBuilder[T any] struct {
    data []T
}

func NewBuilder[T any](slice []T) *SliceBuilder[T] {
    return &SliceBuilder[T]{data: slice}
}

func (b *SliceBuilder[T]) Filter(predicate func(T) bool) *SliceBuilder[T] {
    b.data = Filter(b.data, predicate)
    return b
}

func (b *SliceBuilder[T]) Map(mapper func(T) T) *SliceBuilder[T] {
    b.data = Map(b.data, mapper)
    return b
}

func (b *SliceBuilder[T]) Build() []T {
    return b.data
}

// Usage:
// result := NewBuilder([]int{1,2,3,4,5}).
//     Filter(func(n int) bool { return n > 2 }).
//     Map(func(n int) int { return n * 2 }).
//     Build()
```

**Note / 참고**: This is a potential future enhancement not yet implemented.

이것은 아직 구현되지 않은 잠재적인 향후 개선 사항입니다.

---

## Internal Implementation / 내부 구현

### Memory Management / 메모리 관리

#### Slice Allocation Strategy / 슬라이스 할당 전략

**Pre-allocation when size is known / 크기를 알 때 사전 할당**:
```go
// Good: Pre-allocate with known size
// 좋음: 알려진 크기로 사전 할당
func Map[T any, R any](slice []T, mapper func(T) R) []R {
    result := make([]R, len(slice)) // Pre-allocate exact size
    for i, item := range slice {
        result[i] = mapper(item)
    }
    return result
}
```

**Gradual growth when size is unknown / 크기를 모를 때 점진적 증가**:
```go
// Reasonable: Start with capacity estimate
// 합리적: 용량 추정으로 시작
func Filter[T any](slice []T, predicate func(T) bool) []T {
    if len(slice) == 0 {
        return []T{}
    }

    // Estimate capacity (assume ~50% will match)
    // 용량 추정 (~50%가 일치할 것으로 가정)
    result := make([]T, 0, len(slice)/2)

    for _, item := range slice {
        if predicate(item) {
            result = append(result, item)
        }
    }
    return result
}
```

#### Copy vs Reference / 복사 vs 참조

**When to copy / 복사해야 할 때**:
```go
// Copy: When returning modified slice
// 복사: 수정된 슬라이스를 반환할 때
func Sort[T Ordered](slice []T) []T {
    result := make([]T, len(slice))
    copy(result, slice) // Copy to preserve original
    sort.Slice(result, func(i, j int) bool {
        return result[i] < result[j]
    })
    return result
}
```

**When to reference / 참조해야 할 때**:
```go
// Reference: When only reading
// 참조: 읽기만 할 때
func Contains[T comparable](slice []T, element T) bool {
    for _, item := range slice { // No copy needed
        if item == element {
            return true
        }
    }
    return false
}
```

### Algorithm Complexity / 알고리즘 복잡도

| Operation / 작업 | Time / 시간 | Space / 공간 | Notes / 참고 |
|------------------|------------|-------------|-------------|
| **Basic** | | | |
| Contains | O(n) | O(1) | Linear search / 선형 검색 |
| IndexOf | O(n) | O(1) | Linear search / 선형 검색 |
| Find | O(n) | O(1) | Early return on match / 일치 시 조기 반환 |
| **Transform** | | | |
| Map | O(n) | O(n) | New slice created / 새 슬라이스 생성 |
| Filter | O(n) | O(n) | New slice created / 새 슬라이스 생성 |
| Unique | O(n) | O(n) | Hash map for dedup / 중복 제거를 위한 해시 맵 |
| Flatten | O(n*m) | O(n*m) | n slices, m avg length / n 슬라이스, m 평균 길이 |
| **Aggregate** | | | |
| Reduce | O(n) | O(1) | Single accumulator / 단일 누산기 |
| Sum | O(n) | O(1) | Specialized reduce / 특수화된 reduce |
| Min/Max | O(n) | O(1) | Single pass / 단일 패스 |
| GroupBy | O(n) | O(n) | Hash map / 해시 맵 |
| **Slicing** | | | |
| Chunk | O(n) | O(n) | Creates sub-slices / 하위 슬라이스 생성 |
| Take/Drop | O(n) | O(n) | Sub-slice copy / 하위 슬라이스 복사 |
| **Set Ops** | | | |
| Union | O(n+m) | O(n+m) | Hash map / 해시 맵 |
| Intersection | O(n+m) | O(min(n,m)) | Hash map / 해시 맵 |
| **Sorting** | | | |
| Sort | O(n log n) | O(n) | Go's sort.Slice / Go의 sort.Slice |
| **Utilities** | | | |
| Shuffle | O(n) | O(n) | Fisher-Yates / Fisher-Yates |
| **Statistics** | | | |
| Median | O(n log n) | O(n) | Requires sorting / 정렬 필요 |
| Mode | O(n) | O(n) | Hash map frequency / 해시 맵 빈도 |
| Percentile | O(n log n) | O(n) | Requires sorting / 정렬 필요 |
| Variance/StdDev | O(n) | O(1) | Two passes / 두 번의 패스 |
| MostCommon | O(n log n) | O(n) | Frequency + sort / 빈도 + 정렬 |
| **Diff/Comparison** | | | |
| Diff | O(n+m) | O(n+m) | Hash maps / 해시 맵 |
| DiffBy | O(n+m) | O(n+m) | Hash maps / 해시 맵 |
| EqualUnordered | O(n) | O(n) | Frequency maps / 빈도 맵 |
| HasDuplicates | O(n) | O(n) | Hash set / 해시 집합 |
| **Index Ops** | | | |
| FindIndices | O(n) | O(k) | k = matches / k = 일치 항목 |
| AtIndices | O(k) | O(k) | k = indices / k = 인덱스 |
| RemoveIndices | O(n+k) | O(n) | Hash set + filter / 해시 집합 + 필터 |
| **Conditional** | | | |
| ReplaceIf | O(n) | O(n) | New slice / 새 슬라이스 |
| UpdateWhere | O(n) | O(n) | New slice / 새 슬라이스 |
| **Advanced** | | | |
| Scan | O(n) | O(n) | Intermediate results / 중간 결과 |
| ZipWith | O(min(n,m)) | O(min(n,m)) | Shorter length / 짧은 길이 |
| Rotate | O(n) | O(n) | New slice / 새 슬라이스 |
| **Combinatorial** | | | |
| Permutations | O(n!) | O(n!) | Factorial growth! / 팩토리얼 증가! |
| Combinations | O(C(n,k)) | O(C(n,k)) | Binomial coefficient / 이항 계수 |

### Optimization Techniques / 최적화 기법

#### 1. Early Return / 조기 반환

```go
// Stop as soon as condition is met
// 조건이 충족되면 즉시 중지
func Any[T any](slice []T, predicate func(T) bool) bool {
    for _, item := range slice {
        if predicate(item) {
            return true // Early return
        }
    }
    return false
}
```

#### 2. Pre-allocation / 사전 할당

```go
// Allocate exact size to avoid reallocation
// 재할당을 피하기 위해 정확한 크기 할당
func Map[T any, R any](slice []T, mapper func(T) R) []R {
    result := make([]R, len(slice)) // Pre-allocate
    // ...
}
```

#### 3. Hash Map for O(1) Lookup / O(1) 조회를 위한 해시 맵

```go
// Use map for fast membership testing
// 빠른 멤버십 테스트를 위해 맵 사용
func Union[T comparable](slice1, slice2 []T) []T {
    seen := make(map[T]bool) // O(1) lookup
    // ...
}
```

#### 4. Avoid Unnecessary Allocations / 불필요한 할당 방지

```go
// Return empty slice literal for zero-length inputs
// 길이가 0인 입력에 대해 빈 슬라이스 리터럴 반환
func Filter[T any](slice []T, predicate func(T) bool) []T {
    if len(slice) == 0 {
        return []T{} // No allocation needed
    }
    // ...
}
```

---

## Adding New Features / 새 기능 추가

### Step-by-Step Guide / 단계별 가이드

#### Step 1: Design the Function / 함수 설계

**Questions to ask / 질문할 사항**:
1. What category does this function belong to? / 이 함수는 어떤 카테고리에 속하나요?
2. What type constraints are needed? / 어떤 타입 제약이 필요한가요?
3. What is the expected input/output? / 예상 입력/출력은 무엇인가요?
4. Should it return (value, bool) or just value? / (값, bool)을 반환해야 하나요 아니면 값만 반환해야 하나요?
5. What edge cases exist (empty slice, nil, invalid input)? / 어떤 엣지 케이스가 존재하나요 (빈 슬라이스, nil, 잘못된 입력)?

#### Step 2: Choose the Correct File / 올바른 파일 선택

| Category / 카테고리 | File / 파일 | Example / 예제 |
|---------------------|------------|---------------|
| Search/Find / 검색/찾기 | `basic.go` | Contains, Find, IndexOf |
| Transform / 변환 | `transform.go` | Map, Filter, Unique |
| Aggregate / 집계 | `aggregate.go` | Sum, GroupBy, Reduce |
| Slice / 슬라이스 | `slice.go` | Chunk, Take, Partition |
| Set / 집합 | `set.go` | Union, Intersection |
| Sort / 정렬 | `sort.go` | Sort, IsSorted |
| Test / 테스트 | `predicate.go` | All, Any, Equal |
| Utility / 유틸리티 | `util.go` | ForEach, Shuffle, Zip |
| Statistics / 통계 | `statistics.go` | Median, Variance, Percentile |
| Diff/Compare / 차이/비교 | `diff.go` | Diff, DiffBy, EqualUnordered |
| Index / 인덱스 | `index.go` | FindIndices, AtIndices |
| Conditional / 조건부 | `conditional.go` | ReplaceIf, UpdateWhere |
| Advanced / 고급 | `advanced.go` | Scan, ZipWith, Rotate |
| Combinatorial / 조합 | `combinatorial.go` | Permutations, Combinations |

#### Step 3: Implement the Function / 함수 구현

**Template / 템플릿**:
```go
// FunctionName does X and returns Y.
// FunctionName은 X를 수행하고 Y를 반환합니다.
//
// Example / 예제:
//
//	input := []int{1, 2, 3}
//	result := FunctionName(input)
//	// result: [expected output]
//
// Time Complexity: O(?)
// Space Complexity: O(?)
func FunctionName[T Constraint](slice []T, params...) ReturnType {
    // 1. Handle edge cases
    // 1. 엣지 케이스 처리
    if len(slice) == 0 {
        return /* appropriate zero value */
    }

    // 2. Allocate result
    // 2. 결과 할당
    result := make(ReturnType, /* appropriate size */)

    // 3. Main logic
    // 3. 메인 로직
    for _, item := range slice {
        // ... implementation
    }

    // 4. Return
    // 4. 반환
    return result
}
```

#### Step 4: Write Tests / 테스트 작성

**Test Template / 테스트 템플릿**:
```go
func TestFunctionName(t *testing.T) {
    // Test Case 1: Normal case
    // 테스트 케이스 1: 정상 케이스
    t.Run("normal case", func(t *testing.T) {
        input := []int{1, 2, 3}
        expected := /* expected result */
        result := FunctionName(input)

        if !reflect.DeepEqual(result, expected) {
            t.Errorf("FunctionName() = %v, want %v", result, expected)
        }
    })

    // Test Case 2: Empty slice
    // 테스트 케이스 2: 빈 슬라이스
    t.Run("empty slice", func(t *testing.T) {
        input := []int{}
        expected := /* expected for empty */
        result := FunctionName(input)

        if !reflect.DeepEqual(result, expected) {
            t.Errorf("FunctionName(empty) = %v, want %v", result, expected)
        }
    })

    // Test Case 3: Edge case
    // 테스트 케이스 3: 엣지 케이스
    t.Run("edge case", func(t *testing.T) {
        // ... test edge case
    })
}
```

#### Step 5: Update Documentation / 문서 업데이트

**Files to update / 업데이트할 파일**:
1. Function godoc comments / 함수 godoc 주석
2. `README.md` - Add to function list / 함수 목록에 추가
3. `USER_MANUAL.md` - Add usage example / 사용 예제 추가
4. `DEVELOPER_GUIDE.md` - Update this file / 이 파일 업데이트
5. `examples/sliceutil/main.go` - Add example / 예제 추가

#### Step 6: Run Tests and Benchmarks / 테스트 및 벤치마크 실행

```bash
# Run all tests
# 모든 테스트 실행
go test ./sliceutil -v

# Run specific test
# 특정 테스트 실행
go test ./sliceutil -run TestFunctionName -v

# Run benchmarks
# 벤치마크 실행
go test ./sliceutil -bench=BenchmarkFunctionName -benchmem

# Check coverage
# 커버리지 확인
go test ./sliceutil -cover
```

#### Step 7: Commit Changes / 변경 사항 커밋

```bash
git add sliceutil/
git commit -m "Feat: Add FunctionName to sliceutil package

- Implement FunctionName for [purpose]
- Add comprehensive tests
- Update documentation
- Add example usage

🤖 Generated with Claude Code
Co-Authored-By: Claude <noreply@anthropic.com>"
```

### Adding Statistical Functions / 통계 함수 추가

**Special Considerations for Statistics / 통계에 대한 특별 고려 사항**:

1. **Use Number Constraint / Number 제약 사용**:
   ```go
   // Import constraints from golang.org/x/exp
   // golang.org/x/exp에서 constraints 임포트
   import "golang.org/x/exp/constraints"

   // Define Number constraint if not already defined
   // Number 제약이 아직 정의되지 않은 경우 정의
   type Number interface {
       constraints.Integer | constraints.Float
   }

   // Use in function signature
   // 함수 시그니처에 사용
   func NewStatFunc[T Number](slice []T) (float64, error) {
       // Implementation
   }
   ```

2. **Return (value, error) Pattern / (값, error) 패턴 반환**:
   ```go
   // Always validate input and return error for empty slices
   // 항상 입력을 검증하고 빈 슬라이스에 대해 error 반환
   if len(slice) == 0 {
       return 0, errors.New("cannot calculate X of empty slice")
   }
   ```

3. **Preserve Immutability / 불변성 유지**:
   ```go
   // If sorting is needed, create a copy first
   // 정렬이 필요한 경우 먼저 복사본 생성
   sorted := Sort(slice) // Uses existing Sort function
   ```

4. **Population vs Sample Formulas / 모집단 vs 표본 공식**:
   - Use population formulas (divide by N) by default / 기본적으로 모집단 공식 사용 (N으로 나눔)
   - Document clearly in comments / 주석에 명확히 문서화
   - Consider adding sample variants if needed / 필요한 경우 표본 변형 추가 고려

**Example Statistical Function / 통계 함수 예제**:
```go
// NewStatistic calculates [description].
// Uses [formula type] formula.
// Returns an error if the slice is empty.
//
// NewStatistic은 [설명]을 계산합니다.
// [공식 유형] 공식을 사용합니다.
// 슬라이스가 비어 있으면 에러를 반환합니다.
//
// Example:
//
//	numbers := []float64{1, 2, 3, 4, 5}
//	result, err := sliceutil.NewStatistic(numbers) // Expected result
//
// Time Complexity: O(?)
func NewStatistic[T Number](slice []T) (float64, error) {
    if len(slice) == 0 {
        return 0, errors.New("cannot calculate statistic of empty slice")
    }

    // Implementation
    // 구현

    return result, nil
}
```

### Adding Diff/Comparison Functions / 차이/비교 함수 추가

**Special Considerations for Diff Operations / 차이 작업에 대한 특별 고려 사항**:

1. **Use Hash Maps for Efficiency / 효율성을 위해 해시 맵 사용**:
   ```go
   // Build hash sets for O(1) lookup
   // O(1) 조회를 위한 해시 집합 구축
   oldSet := make(map[T]bool)
   newSet := make(map[T]bool)
   ```

2. **Consider DiffResult Struct / DiffResult 구조체 고려**:
   ```go
   // For comprehensive diff results
   // 포괄적인 차이 결과를 위해
   return DiffResult[T]{
       Added:     added,
       Removed:   removed,
       Unchanged: unchanged,
   }
   ```

3. **Support Key-Based Comparison / 키 기반 비교 지원**:
   ```go
   // For complex types, provide a "By" variant
   // 복잡한 타입의 경우 "By" 변형 제공
   func DiffBy[T any, K comparable](old, new []T, keyFunc func(T) K) DiffResult[T]
   ```

4. **Handle Duplicates Carefully / 중복을 신중하게 처리**:
   ```go
   // Decide: preserve duplicates or use Unique()?
   // 결정: 중복을 유지할 것인지 Unique() 사용할 것인지?
   unchanged = Unique(unchanged) // If deduplication is needed
   ```

**Example Diff Function / 차이 함수 예제**:
```go
// NewDiff compares two slices and returns [what].
// Uses hash maps for O(n + m) complexity.
//
// NewDiff는 두 슬라이스를 비교하고 [무엇을] 반환합니다.
// O(n + m) 복잡도를 위해 해시 맵을 사용합니다.
//
// Example:
//
//	a := []int{1, 2, 3}
//	b := []int{2, 3, 4}
//	result := sliceutil.NewDiff(a, b)
//
// Time Complexity: O(n + m)
// Space Complexity: O(n + m)
func NewDiff[T comparable](a, b []T) SomeResult[T] {
    // Build hash sets
    // 해시 집합 구축
    setA := make(map[T]bool)
    for _, v := range a {
        setA[v] = true
    }

    setB := make(map[T]bool)
    for _, v := range b {
        setB[v] = true
    }

    // Compare and categorize
    // 비교 및 분류

    return result
}
```

---

## Testing Guide / 테스트 가이드

### Test Structure / 테스트 구조

The sliceutil package has comprehensive test coverage (100%) with tests organized by file:

sliceutil 패키지는 파일별로 구성된 포괄적인 테스트 커버리지(100%)를 가지고 있습니다:

```
sliceutil/
├── sliceutil_test.go      # Package-level tests
├── basic_test.go          # Tests for basic operations (10 functions)
├── transform_test.go      # Tests for transformation functions (8 functions)
├── aggregate_test.go      # Tests for aggregation functions (7 functions)
├── slice_test.go          # Tests for slicing functions (7 functions)
├── set_test.go            # Tests for set operations (6 functions)
├── sort_test.go           # Tests for sorting functions (5 functions)
├── predicate_test.go      # Tests for predicate functions (6 functions)
├── util_test.go           # Tests for utility functions (11 functions)
├── statistics_test.go     # Tests for statistical functions (8 functions)
├── diff_test.go           # Tests for diff/comparison (4 functions)
├── index_test.go          # Tests for index operations (3 functions)
├── conditional_test.go    # Tests for conditional operations (3 functions)
├── advanced_test.go       # Tests for advanced functional ops (4 functions)
└── combinatorial_test.go  # Tests for combinatorial ops (2 functions)
```

**Test Coverage Details / 테스트 커버리지 세부 정보**:
- Total test files: 15 / 총 테스트 파일: 15개
- Total test lines: ~4,070 / 총 테스트 라인: ~4,070줄
- Coverage: 100% across all 95 functions / 커버리지: 95개 함수 모두 100%
- Test cases: 260+ comprehensive test scenarios / 테스트 케이스: 260개 이상의 종합 테스트 시나리오

### Test Categories / 테스트 카테고리

#### 1. Functional Tests / 기능 테스트

**Purpose / 목적**: Verify correct behavior for normal inputs.

정상 입력에 대한 올바른 동작 확인.

```go
func TestMap_NormalCase(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    expected := []int{2, 4, 6, 8, 10}

    result := Map(input, func(n int) int {
        return n * 2
    })

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Map() = %v, want %v", result, expected)
    }
}
```

#### 2. Edge Case Tests / 엣지 케이스 테스트

**Purpose / 목적**: Test boundary conditions and special cases.

경계 조건 및 특수 케이스 테스트.

```go
func TestMap_EmptySlice(t *testing.T) {
    input := []int{}
    result := Map(input, func(n int) int { return n * 2 })

    if len(result) != 0 {
        t.Errorf("Map(empty) should return empty slice, got %v", result)
    }
}

func TestMap_SingleElement(t *testing.T) {
    input := []int{5}
    expected := []int{10}
    result := Map(input, func(n int) int { return n * 2 })

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Map(single) = %v, want %v", result, expected)
    }
}
```

#### 3. Type Tests / 타입 테스트

**Purpose / 목적**: Verify generics work with different types.

제네릭이 다양한 타입과 작동하는지 확인.

```go
func TestMap_DifferentTypes(t *testing.T) {
    // Test with strings
    // 문자열 테스트
    strings := []string{"a", "b", "c"}
    upper := Map(strings, func(s string) string {
        return strings.ToUpper(s)
    })
    // Verify result...

    // Test with floats
    // 실수 테스트
    floats := []float64{1.1, 2.2, 3.3}
    doubled := Map(floats, func(f float64) float64 {
        return f * 2
    })
    // Verify result...
}
```

#### 4. Immutability Tests / 불변성 테스트

**Purpose / 목적**: Ensure original slices are not modified.

원본 슬라이스가 수정되지 않도록 보장.

```go
func TestMap_ImmutabilityGuarantee(t *testing.T) {
    original := []int{1, 2, 3}
    originalCopy := make([]int, len(original))
    copy(originalCopy, original)

    // Call Map
    _ = Map(original, func(n int) int { return n * 2 })

    // Verify original unchanged
    // 원본이 변경되지 않았는지 확인
    if !reflect.DeepEqual(original, originalCopy) {
        t.Errorf("Map() modified original slice")
    }
}
```

#### 5. Benchmark Tests / 벤치마크 테스트

**Purpose / 목적**: Measure performance.

성능 측정.

```go
func BenchmarkMap_SmallSlice(b *testing.B) {
    input := []int{1, 2, 3, 4, 5}
    mapper := func(n int) int { return n * 2 }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = Map(input, mapper)
    }
}

func BenchmarkMap_LargeSlice(b *testing.B) {
    input := make([]int, 10000)
    for i := range input {
        input[i] = i
    }
    mapper := func(n int) int { return n * 2 }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = Map(input, mapper)
    }
}
```

### Running Tests / 테스트 실행

```bash
# Run all tests with verbose output
# 상세 출력으로 모든 테스트 실행
go test ./sliceutil -v

# Run specific test
# 특정 테스트 실행
go test ./sliceutil -run TestMap -v

# Run tests with coverage
# 커버리지와 함께 테스트 실행
go test ./sliceutil -cover

# Generate coverage report
# 커버리지 보고서 생성
go test ./sliceutil -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks
# 벤치마크 실행
go test ./sliceutil -bench=. -benchmem

# Run specific benchmark
# 특정 벤치마크 실행
go test ./sliceutil -bench=BenchmarkMap -benchmem
```

### Test Coverage Requirements / 테스트 커버리지 요구사항

**Minimum Requirements / 최소 요구사항**:
- Overall coverage: ≥ 90% / 전체 커버리지: ≥ 90%
- Per-file coverage: ≥ 85% / 파일별 커버리지: ≥ 85%
- Critical paths: 100% / 중요 경로: 100%

**Current Status / 현재 상태** (v1.7.023):
- Overall: 100% ✅
- All files: 100% ✅
- Total functions tested: 95/95 ✅
- Test cases: 260+ ✅

---

## Performance / 성능

### Performance Characteristics / 성능 특성

#### Small Slices (< 100 elements) / 작은 슬라이스 (< 100 요소)

- **Overhead**: Minimal function call overhead / 최소 함수 호출 오버헤드
- **Optimization**: Compiler inlining possible / 컴파일러 인라인 가능
- **Memory**: Pre-allocation beneficial / 사전 할당이 유익함

#### Medium Slices (100 - 10,000 elements) / 중간 슬라이스 (100 - 10,000 요소)

- **Performance**: Linear scaling / 선형 확장
- **Memory**: Efficient with pre-allocation / 사전 할당으로 효율적
- **Recommendation**: Use appropriate functions / 적절한 함수 사용

#### Large Slices (> 10,000 elements) / 큰 슬라이스 (> 10,000 요소)

- **Performance**: May benefit from chunking / 청킹으로 이익을 얻을 수 있음
- **Memory**: Watch for allocations / 할당에 주의
- **Recommendation**: Consider batch processing / 배치 처리 고려

### Benchmark Results / 벤치마크 결과

**Example Benchmarks (on Apple M1) / 예제 벤치마크 (Apple M1 기준)**:

```
# Basic Operations / 기본 작업
BenchmarkContains/small-8         	100000000	        10.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkContains/large-8         	  500000	      2450 ns/op	       0 B/op	       0 allocs/op

# Transform / 변환
BenchmarkMap/small-8              	 50000000	        35.1 ns/op	      40 B/op	       1 allocs/op
BenchmarkMap/large-8              	    50000	     28540 ns/op	   81920 B/op	       1 allocs/op

BenchmarkFilter/small-8           	 30000000	        42.3 ns/op	      32 B/op	       1 allocs/op
BenchmarkFilter/large-8           	    30000	     38920 ns/op	   40960 B/op	       1 allocs/op

# Sorting / 정렬
BenchmarkSort/small-8             	 10000000	       125 ns/op	      40 B/op	       1 allocs/op
BenchmarkSort/large-8             	     5000	    289450 ns/op	   81920 B/op	       1 allocs/op

# Statistics / 통계
BenchmarkMedian/small-8           	  8000000	       145 ns/op	      40 B/op	       1 allocs/op
BenchmarkMedian/large-8           	     4000	    298230 ns/op	   81920 B/op	       1 allocs/op

BenchmarkVariance/small-8         	 80000000	        18.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkVariance/large-8         	    70000	     17200 ns/op	       0 B/op	       0 allocs/op

# Diff/Comparison / 차이/비교
BenchmarkDiff/small-8             	 20000000	        75.2 ns/op	     128 B/op	       3 allocs/op
BenchmarkDiff/large-8             	    15000	     82300 ns/op	  163840 B/op	       3 allocs/op

# Advanced / 고급
BenchmarkScan/small-8             	 40000000	        42.0 ns/op	      48 B/op	       1 allocs/op
BenchmarkScan/large-8             	    45000	     29100 ns/op	   90112 B/op	       1 allocs/op

# Combinatorial (WARNING: Small sizes only!) / 조합 (경고: 작은 크기만!)
BenchmarkPermutations/n=5-8       	   500000	      3420 ns/op	    1920 B/op	     120 allocs/op
BenchmarkCombinations/n=10_k=5-8  	   100000	     18500 ns/op	    6144 B/op	     252 allocs/op
```

**Performance Notes for New Functions / 새 함수에 대한 성능 참고사항**:

1. **Statistical Functions / 통계 함수**:
   - Median/Percentile require sorting: O(n log n) / Median/Percentile은 정렬 필요: O(n log n)
   - Variance/StdDev are O(n) but require two passes / Variance/StdDev는 O(n)이지만 두 번의 패스 필요
   - Use for analysis, not hot paths / 분석용으로 사용, 핫 패스에는 사용 안 함

2. **Diff Operations / 차이 작업**:
   - Efficient O(n + m) with hash maps / 해시 맵으로 효율적인 O(n + m)
   - Good for comparing datasets / 데이터셋 비교에 적합
   - Memory overhead for hash maps / 해시 맵을 위한 메모리 오버헤드

3. **Combinatorial Functions / 조합 함수**:
   - **USE WITH EXTREME CAUTION** / **극도로 주의하여 사용**
   - Factorial/exponential growth / 팩토리얼/지수 증가
   - Only for small inputs (n ≤ 10) / 작은 입력에만 사용 (n ≤ 10)
   - Consider generating on-demand instead of all at once / 한 번에 모두 생성하는 대신 주문형 생성 고려

### Optimization Tips / 최적화 팁

#### 1. Use Specific Functions / 특정 함수 사용

```go
// ✅ Good - Use Sum (O(n), specialized)
// 좋음 - Sum 사용 (O(n), 특수화됨)
total := Sum(numbers)

// ❌ Less efficient - Use Reduce (O(n), but slower)
// 덜 효율적 - Reduce 사용 (O(n), 하지만 느림)
total := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
```

#### 2. Avoid Unnecessary Operations / 불필요한 작업 방지

```go
// ✅ Good - Single pass
// 좋음 - 단일 패스
evens, odds := Partition(numbers, func(n int) bool { return n%2 == 0 })

// ❌ Less efficient - Two passes
// 덜 효율적 - 두 번의 패스
evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
odds := Filter(numbers, func(n int) bool { return n%2 != 0 })
```

#### 3. Use Chunking for Large Data / 대용량 데이터에 청킹 사용

```go
// ✅ Good - Process in chunks
// 좋음 - 청크로 처리
chunks := Chunk(largeData, 1000)
for _, chunk := range chunks {
    result := Map(chunk, expensiveOperation)
    // Process result...
}
```

#### 4. Check Before Expensive Operations / 비용이 많이 드는 작업 전에 확인

```go
// ✅ Good - Check first
// 좋음 - 먼저 확인
if !IsSorted(data) {
    data = Sort(data)
}

// ❌ Wasteful - Always sort
// 낭비 - 항상 정렬
data = Sort(data)
```

---

## Contributing Guidelines / 기여 가이드라인

### How to Contribute / 기여 방법

1. **Fork the Repository / 저장소 포크**
   - Go to https://github.com/arkd0ng/go-utils
   - Click "Fork" button

2. **Create Feature Branch / 기능 브랜치 생성**
   ```bash
   git checkout -b feature/my-new-function
   ```

3. **Make Changes / 변경 수행**
   - Add function implementation
   - Write comprehensive tests
   - Update documentation

4. **Test Your Changes / 변경 사항 테스트**
   ```bash
   go test ./sliceutil -v
   go test ./sliceutil -cover
   ```

5. **Commit with Conventional Commits / 규칙적 커밋으로 커밋**
   ```bash
   git commit -m "Feat: Add NewFunction to sliceutil

   - Implement NewFunction for [purpose]
   - Add tests with 100% coverage
   - Update documentation"
   ```

6. **Push and Create Pull Request / 푸시 및 풀 리퀘스트 생성**
   ```bash
   git push origin feature/my-new-function
   ```

### Code Review Checklist / 코드 리뷰 체크리스트

- [ ] Function follows design principles / 함수가 설계 원칙을 따름
- [ ] Comprehensive tests (≥90% coverage) / 포괄적인 테스트 (≥90% 커버리지)
- [ ] Bilingual documentation (English/Korean) / 이중 언어 문서 (영문/한글)
- [ ] Godoc comments with examples / 예제가 있는 Godoc 주석
- [ ] No breaking changes to existing API / 기존 API에 대한 파괴적 변경 없음
- [ ] Benchmarks for performance-critical code / 성능 중요 코드에 대한 벤치마크
- [ ] Edge cases handled / 엣지 케이스 처리됨
- [ ] Immutability guaranteed / 불변성 보장됨

### Commit Message Format / 커밋 메시지 형식

```
Type: Short description (imperative mood)

- Detailed change 1
- Detailed change 2
- Detailed change 3

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude <noreply@anthropic.com>
```

**Types / 타입**:
- `Feat`: New feature / 새 기능
- `Fix`: Bug fix / 버그 수정
- `Docs`: Documentation / 문서
- `Test`: Tests / 테스트
- `Refactor`: Code refactoring / 코드 리팩토링
- `Perf`: Performance improvement / 성능 개선

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

#### Functions / 함수

- Use PascalCase for exported functions / 내보낸 함수에 PascalCase 사용
- Use descriptive names (Map, Filter, Reduce) / 설명적 이름 사용 (Map, Filter, Reduce)
- Avoid abbreviations unless universally known / 보편적으로 알려지지 않은 약어 피함

```go
// ✅ Good
func Map[T any, R any](slice []T, mapper func(T) R) []R

// ❌ Bad
func Mp[T any, R any](s []T, f func(T) R) []R
```

#### Variables / 변수

- Use camelCase for local variables / 로컬 변수에 camelCase 사용
- Use short names for short scopes / 짧은 범위에 짧은 이름 사용
- Use descriptive names for long scopes / 긴 범위에 설명적 이름 사용

```go
// ✅ Good
for i, item := range slice {
    result[i] = mapper(item)
}

// ❌ Bad
for index, currentSliceElement := range slice {
    result[index] = mapper(currentSliceElement)
}
```

#### Type Parameters / 타입 매개변수

- Use single letters for simple generics (T, R, K, V) / 간단한 제네릭에 단일 문자 사용 (T, R, K, V)
- T: Type (general) / 타입 (일반)
- R: Return type (when different from T) / 반환 타입 (T와 다를 때)
- K: Key type (for maps) / 키 타입 (맵용)
- V: Value type (for maps) / 값 타입 (맵용)

```go
func Map[T any, R any](slice []T, mapper func(T) R) []R
func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T
```

### Documentation Style / 문서 스타일

#### Function Comments / 함수 주석

```go
// FunctionName does X and returns Y. (English first)
// FunctionName은 X를 수행하고 Y를 반환합니다. (Korean second)
//
// Additional explanation if needed. (English)
// 필요한 경우 추가 설명. (Korean)
//
// Example / 예제:
//
//	input := []int{1, 2, 3}
//	result := FunctionName(input, predicate)
//	// Output: [expected]
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func FunctionName[T any](slice []T) ReturnType {
    // Implementation
}
```

#### Code Comments / 코드 주석

```go
// Bilingual inline comments for clarity
// 명확성을 위한 이중 언어 인라인 주석
func Example() {
    // Pre-allocate with known size / 알려진 크기로 사전 할당
    result := make([]int, len(slice))

    // Early return for edge case / 엣지 케이스에 대한 조기 반환
    if len(slice) == 0 {
        return []int{}
    }
}
```

### Error Handling / 에러 처리

Since sliceutil functions operate on in-memory slices, they typically don't return errors. Instead:

sliceutil 함수는 메모리 내 슬라이스에서 작동하므로 일반적으로 오류를 반환하지 않습니다. 대신:

- Return zero values for not found / 찾지 못한 경우 제로값 반환
- Return empty slices for no matches / 일치하는 항목이 없으면 빈 슬라이스 반환
- Use (value, bool) pattern for optional returns / 선택적 반환에 (값, bool) 패턴 사용

```go
// Pattern: (value, found)
// 패턴: (값, 발견됨)
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
    for _, item := range slice {
        if predicate(item) {
            return item, true
        }
    }
    var zero T
    return zero, false
}
```

---

## Conclusion / 결론

The sliceutil package (v1.7.023) provides a comprehensive, type-safe, and performant solution for slice operations in Go. By following this developer guide, you can understand the internal architecture, contribute effectively, and use the package optimally.

sliceutil 패키지 (v1.7.023)는 Go에서 슬라이스 작업을 위한 포괄적이고 타입 안전하며 성능이 뛰어난 솔루션을 제공합니다. 이 개발자 가이드를 따르면 내부 아키텍처를 이해하고 효과적으로 기여하며 패키지를 최적으로 사용할 수 있습니다.

**Key Takeaways / 주요 요점**:
- **Extreme simplicity through generics** / 제네릭을 통한 극도의 간결함
  - 95 functions across 14 categories / 14개 카테고리에 걸쳐 95개 함수
  - Reduce 20+ lines to 1 line / 20줄 이상을 1줄로 축소
- **Immutability for safety** / 안전을 위한 불변성
  - All operations return new slices / 모든 작업이 새 슬라이스 반환
  - No side effects / 부작용 없음
- **Comprehensive test coverage (100%)** / 포괄적인 테스트 커버리지 (100%)
  - 260+ test scenarios / 260개 이상의 테스트 시나리오
  - All 95 functions fully tested / 95개 함수 모두 완전히 테스트됨
- **Performance-optimized implementations** / 성능 최적화된 구현
  - Efficient algorithms (hash maps, early returns) / 효율적인 알고리즘 (해시 맵, 조기 반환)
  - Minimal allocations / 최소 할당
- **Minimal external dependencies** / 최소 외부 의존성
  - Only golang.org/x/exp for constraints / constraints를 위해 golang.org/x/exp만 사용
  - Standard library first / 표준 라이브러리 우선

**Version History / 버전 히스토리**:
- **v1.7.018**: Initial release with 60 functions (8 categories) / 60개 함수로 첫 릴리스 (8개 카테고리)
- **v1.7.023** (current): Extended to 95 functions (14 categories) / 95개 함수로 확장 (14개 카테고리)
  - Added: Statistics (8), Diff/Comparison (4), Index Operations (3) / 추가: 통계 (8개), 차이/비교 (4개), 인덱스 작업 (3개)
  - Added: Conditional (3), Advanced (4), Combinatorial (2) / 추가: 조건부 (3개), 고급 (4개), 조합 (2개)
  - Enhanced: Number constraint with golang.org/x/exp / 개선: golang.org/x/exp로 Number 제약
  - New types: DiffResult struct / 새 타입: DiffResult 구조체

**What's New in v1.7.023 / v1.7.023의 새로운 기능**:
1. **Statistical Analysis / 통계 분석**: Median, Mode, Variance, Percentile, StandardDeviation, MostCommon, LeastCommon
2. **Advanced Diff/Compare / 고급 차이/비교**: Diff, DiffBy with DiffResult, EqualUnordered, HasDuplicates
3. **Index-Based Operations / 인덱스 기반 작업**: FindIndices, AtIndices, RemoveIndices
4. **Conditional Transformations / 조건부 변환**: ReplaceIf, ReplaceAll, UpdateWhere
5. **Advanced Functional / 고급 함수형**: Scan, ZipWith, RotateLeft, RotateRight
6. **Combinatorial Math / 조합 수학**: Permutations, Combinations (use cautiously!)

For user documentation, see the [User Manual](USER_MANUAL.md).

사용자 문서는 [사용자 매뉴얼](USER_MANUAL.md)을 참조하세요.

**Contributing / 기여하기**:
We welcome contributions! Please follow the guidelines in this document and ensure:
- 100% test coverage for new functions / 새 함수에 대해 100% 테스트 커버리지
- Bilingual documentation (English/Korean) / 이중 언어 문서 (영문/한글)
- Immutability guarantees / 불변성 보장
- Performance benchmarks / 성능 벤치마크

기여를 환영합니다! 이 문서의 가이드라인을 따르고 다음을 확인하세요:

---

**End of Developer Guide / 개발자 가이드 끝**
