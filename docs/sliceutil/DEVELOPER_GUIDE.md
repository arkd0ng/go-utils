# Sliceutil Package - Developer Guide / 개발자 가이드

**Version / 버전**: v1.7.017
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
4. **Zero Dependencies / 제로 의존성**: Only uses Go standard library / Go 표준 라이브러리만 사용
5. **Functional Programming / 함수형 프로그래밍**: Higher-order functions (Map, Filter, Reduce) / 고차 함수 (Map, Filter, Reduce)
6. **Performance / 성능**: Efficient algorithms with minimal allocations / 최소 할당으로 효율적인 알고리즘
7. **Comprehensive Coverage / 포괄적인 커버리지**: 60 functions across 8 categories / 8개 카테고리에 걸쳐 60개 함수

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
        │                │                 │   (60 total)   │
        │ - Number       │                 │                │
        │ - Ordered      │                 │ 8 Categories:  │
        │ - comparable   │                 │                │
        │ - any          │                 │ 1. Basic Ops   │
        └────────────────┘                 │ 2. Transform   │
                                          │ 3. Aggregate   │
                                          │ 4. Slicing     │
                                          │ 5. Set Ops     │
                                          │ 6. Sorting     │
                                          │ 7. Predicates  │
                                          │ 8. Utilities   │
                                          └────────────────┘
```

### Component Interaction / 컴포넌트 상호작용

```
User Code / 사용자 코드
    ↓
┌───────────────────────────┐
│   Public API Functions    │  ← Type-safe generic functions
│   (60 functions)          │    타입 안전 제네릭 함수
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Type Constraints        │  ← Number, Ordered, comparable
│   (Generic Validation)    │    제네릭 검증
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Core Algorithm          │  ← Efficient slice operations
│   Implementation          │    효율적인 슬라이스 작업
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Go Standard Library     │  ← slices, sort, math/rand
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
└── README.md              # Package README
                           # 패키지 README
```

### File Responsibilities / 파일별 책임

| File / 파일 | Purpose / 목적 | Functions / 함수 | Lines / 줄 수 |
|-------------|---------------|-----------------|--------------|
| `sliceutil.go` | Package documentation, type definitions / 패키지 문서, 타입 정의 | Types: Number, Ordered | ~100 |
| `basic.go` | Search, find, count operations / 검색, 찾기, 개수 세기 작업 | 10 functions | ~300 |
| `transform.go` | Map, filter, unique, reverse / Map, 필터, 고유값, 역순 | 8 functions | ~250 |
| `aggregate.go` | Reduce, sum, min, max, groupby / Reduce, 합계, 최소, 최대, 그룹화 | 7 functions | ~220 |
| `slice.go` | Chunk, take, drop, partition / 청크, 가져오기, 제거, 파티션 | 7 functions | ~200 |
| `set.go` | Union, intersection, difference / 합집합, 교집합, 차집합 | 6 functions | ~180 |
| `sort.go` | Sort, sort by, is sorted / 정렬, 정렬 기준, 정렬 확인 | 5 functions | ~150 |
| `predicate.go` | All, any, none, equal / 모두, 어느, 없음, 동일 | 6 functions | ~150 |
| `util.go` | ForEach, join, shuffle, zip / ForEach, 결합, 섞기, 압축 | 11 functions | ~300 |
| `*_test.go` | Comprehensive tests for each file / 각 파일에 대한 종합 테스트 | Test functions | ~2,500 total |

**Total Package Size / 전체 패키지 크기**: ~4,350 lines (implementation + tests) / ~4,350줄 (구현 + 테스트)

---

## Core Components / 핵심 컴포넌트

### 1. Type Constraints / 타입 제약조건

**Location / 위치**: `sliceutil/sliceutil.go`

The package defines three key type constraints for generics:

패키지는 제네릭을 위한 세 가지 주요 타입 제약조건을 정의합니다:

```go
// Number constraint for numeric operations (sum, average, min, max)
// 숫자 작업을 위한 Number 제약 (합계, 평균, 최소, 최대)
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Ordered constraint for sorting operations
// 정렬 작업을 위한 Ordered 제약
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}
```

**Design Rationale / 설계 근거**:
- `Number`: Used for mathematical operations (Sum, Average, Min, Max) / 수학적 작업에 사용 (Sum, Average, Min, Max)
- `Ordered`: Used for comparison and sorting (Sort, IsSorted) / 비교 및 정렬에 사용 (Sort, IsSorted)
- `comparable`: Built-in Go constraint for equality checks (Contains, Unique) / 동등성 검사를 위한 내장 Go 제약 (Contains, Unique)
- `any`: Used for type-agnostic operations (Map, Filter, ForEach) / 타입에 구애받지 않는 작업에 사용 (Map, Filter, ForEach)

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
| Contains | O(n) | O(1) | Linear search / 선형 검색 |
| IndexOf | O(n) | O(1) | Linear search / 선형 검색 |
| Find | O(n) | O(1) | Early return on match / 일치 시 조기 반환 |
| Map | O(n) | O(n) | New slice created / 새 슬라이스 생성 |
| Filter | O(n) | O(n) | New slice created / 새 슬라이스 생성 |
| Unique | O(n) | O(n) | Hash map for dedup / 중복 제거를 위한 해시 맵 |
| Flatten | O(n*m) | O(n*m) | n slices, m avg length / n 슬라이스, m 평균 길이 |
| Reduce | O(n) | O(1) | Single accumulator / 단일 누산기 |
| Sum | O(n) | O(1) | Specialized reduce / 특수화된 reduce |
| Min/Max | O(n) | O(1) | Single pass / 단일 패스 |
| GroupBy | O(n) | O(n) | Hash map / 해시 맵 |
| Chunk | O(n) | O(n) | Creates sub-slices / 하위 슬라이스 생성 |
| Union | O(n+m) | O(n+m) | Hash map / 해시 맵 |
| Intersection | O(n+m) | O(min(n,m)) | Hash map / 해시 맵 |
| Sort | O(n log n) | O(n) | Go's sort.Slice / Go의 sort.Slice |
| Shuffle | O(n) | O(n) | Fisher-Yates / Fisher-Yates |

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

---

## Testing Guide / 테스트 가이드

### Test Structure / 테스트 구조

The sliceutil package has comprehensive test coverage (99.5%) with tests organized by file:

sliceutil 패키지는 파일별로 구성된 포괄적인 테스트 커버리지(99.5%)를 가지고 있습니다:

```
sliceutil/
├── basic_test.go          # Tests for basic operations
├── transform_test.go      # Tests for transformation functions
├── aggregate_test.go      # Tests for aggregation functions
├── slice_test.go          # Tests for slicing functions
├── set_test.go            # Tests for set operations
├── sort_test.go           # Tests for sorting functions
├── predicate_test.go      # Tests for predicate functions
└── util_test.go           # Tests for utility functions
```

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

**Current Status / 현재 상태**:
- Overall: 99.5% ✅
- All files: > 95% ✅

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
BenchmarkContains/small-8         	100000000	        10.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkContains/large-8         	  500000	      2450 ns/op	       0 B/op	       0 allocs/op

BenchmarkMap/small-8              	 50000000	        35.1 ns/op	      40 B/op	       1 allocs/op
BenchmarkMap/large-8              	    50000	     28540 ns/op	   81920 B/op	       1 allocs/op

BenchmarkFilter/small-8           	 30000000	        42.3 ns/op	      32 B/op	       1 allocs/op
BenchmarkFilter/large-8           	    30000	     38920 ns/op	   40960 B/op	       1 allocs/op

BenchmarkSort/small-8             	 10000000	       125 ns/op	      40 B/op	       1 allocs/op
BenchmarkSort/large-8             	     5000	    289450 ns/op	   81920 B/op	       1 allocs/op
```

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

The sliceutil package provides a comprehensive, type-safe, and performant solution for slice operations in Go. By following this developer guide, you can understand the internal architecture, contribute effectively, and use the package optimally.

sliceutil 패키지는 Go에서 슬라이스 작업을 위한 포괄적이고 타입 안전하며 성능이 뛰어난 솔루션을 제공합니다. 이 개발자 가이드를 따르면 내부 아키텍처를 이해하고 효과적으로 기여하며 패키지를 최적으로 사용할 수 있습니다.

**Key Takeaways / 주요 요점**:
- Extreme simplicity through generics / 제네릭을 통한 극도의 간결함
- Immutability for safety / 안전을 위한 불변성
- Comprehensive test coverage (99.5%) / 포괄적인 테스트 커버리지 (99.5%)
- Performance-optimized implementations / 성능 최적화된 구현
- Zero external dependencies / 제로 외부 의존성

For user documentation, see the [User Manual](USER_MANUAL.md).

사용자 문서는 [사용자 매뉴얼](USER_MANUAL.md)을 참조하세요.

---

**End of Developer Guide / 개발자 가이드 끝**
