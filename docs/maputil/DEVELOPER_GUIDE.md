# Maputil Package - Developer Guide / 개발자 가이드

**Version / 버전**: v1.8.005
**Package / 패키지**: `github.com/arkd0ng/go-utils/maputil`
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

The Maputil package follows these core design principles:

Maputil 패키지는 다음과 같은 핵심 설계 원칙을 따릅니다:

1. **Extreme Simplicity / 극도의 간결함**: Reduce 20+ lines of repetitive code to 1-2 lines / 20줄 이상의 반복 코드를 1-2줄로 축소
2. **Type Safety / 타입 안전성**: Leverage Go 1.18+ generics for compile-time type checking / Go 1.18+ 제네릭을 활용한 컴파일 타임 타입 검사
3. **Immutability / 불변성**: All operations return new maps (except Assign), never modifying originals / 모든 작업은 새 맵을 반환하며 (Assign 제외) 원본을 절대 수정하지 않음
4. **Zero Dependencies / 제로 의존성**: Only uses Go standard library / Go 표준 라이브러리만 사용
5. **Functional Programming / 함수형 프로그래밍**: Higher-order functions (Map, Filter, Reduce) / 고차 함수 (Map, Filter, Reduce)
6. **Performance / 성능**: Efficient algorithms with minimal allocations / 최소 할당으로 효율적인 알고리즘
7. **Comprehensive Coverage / 포괄적인 커버리지**: 81 functions across 10 categories / 10개 카테고리에 걸쳐 81개 함수

### High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────────────────┐
│                        Maputil Package                          │
│                     github.com/arkd0ng/go-utils/maputil         │
└─────────────────────────────────────────────────────────────────┘
                                  │
                ┌─────────────────┴─────────────────┐
                │                                   │
        ┌───────▼────────┐                 ┌───────▼────────┐
        │  Type System   │                 │   Functions    │
        │                │                 │   (81 total)   │
        │ - Entry[K,V]   │                 │                │
        │   struct       │                 │ 10 Categories: │
        │                │                 │                │
        │ - Number       │                 │ 1. Basic (11)  │
        │   constraint   │                 │ 2. Transform   │
        │                │                 │    (10)        │
        │ - Ordered      │                 │ 3. Aggregate   │
        │   constraint   │                 │    (9)         │
        │                │                 │ 4. Merge (8)   │
        │ - comparable   │                 │ 5. Filter (7)  │
        │   (built-in)   │                 │ 6. Conversion  │
        │                │                 │    (9)         │
        │ - any          │                 │ 7. Predicate   │
        │   (built-in)   │                 │    (7)         │
        └────────────────┘                 │ 8. Keys (8)    │
                                          │ 9. Values (7)  │
                                          │ 10. Comparison │
                                          │     (6)        │
                                          └────────────────┘
```

### Component Interaction / 컴포넌트 상호작용

```
User Code / 사용자 코드
    ↓
┌───────────────────────────┐
│   Public API Functions    │  ← Type-safe generic functions
│   (81 functions)          │    타입 안전 제네릭 함수
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Type Constraints        │  ← Number, Ordered, comparable
│   (Generic Validation)    │    제네릭 검증
│   + Entry[K,V] struct     │    + Entry[K,V] 구조체
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Core Algorithm          │  ← Efficient map operations
│   Implementation          │    효율적인 맵 작업
│   + Functional methods    │    + 함수형 메서드
│   + Set operations        │    + 집합 작업
└───────────────────────────┘
    ↓
┌───────────────────────────┐
│   Go Standard Library     │  ← sort, encoding/json
│                           │    + logging (version)
└───────────────────────────┘
```

---

## Package Structure / 패키지 구조

### File Organization / 파일 구성

```
maputil/
├── maputil.go             # Package documentation, types, constraints
│                          # 패키지 문서, 타입, 제약조건
├── maputil_test.go        # Package-level tests (version, import)
│                          # 패키지 레벨 테스트 (버전, 임포트)
├── basic.go               # Basic operations (11 functions)
│                          # 기본 작업 (11개 함수)
├── basic_test.go          # Tests for basic operations
│                          # 기본 작업 테스트
├── transform.go           # Transformation functions (10 functions)
│                          # 변환 함수 (10개 함수)
├── transform_test.go      # Tests for transformation functions
│                          # 변환 함수 테스트
├── aggregate.go           # Aggregation functions (9 functions)
│                          # 집계 함수 (9개 함수)
├── aggregate_test.go      # Tests for aggregation functions
│                          # 집계 함수 테스트
├── merge.go               # Merge operations (8 functions)
│                          # 병합 작업 (8개 함수)
├── merge_test.go          # Tests for merge operations
│                          # 병합 작업 테스트
├── filter.go              # Filter operations (7 functions)
│                          # 필터 작업 (7개 함수)
├── filter_test.go         # Tests for filter operations
│                          # 필터 작업 테스트
├── convert.go             # Conversion functions (9 functions)
│                          # 변환 함수 (9개 함수)
├── convert_test.go        # Tests for conversion functions
│                          # 변환 함수 테스트
├── predicate.go           # Predicate functions (7 functions)
│                          # 조건자 함수 (7개 함수)
├── predicate_test.go      # Tests for predicate functions
│                          # 조건자 함수 테스트
├── keys.go                # Key operations (8 functions)
│                          # 키 작업 (8개 함수)
├── keys_test.go           # Tests for key operations
│                          # 키 작업 테스트
├── values.go              # Value operations (7 functions)
│                          # 값 작업 (7개 함수)
├── values_test.go         # Tests for value operations
│                          # 값 작업 테스트
├── comparison.go          # Comparison operations (6 functions)
│                          # 비교 작업 (6개 함수)
├── comparison_test.go     # Tests for comparison operations
│                          # 비교 작업 테스트
└── README.md              # Package README
                           # 패키지 README
```

### File Responsibilities / 파일별 책임

| File / 파일 | Purpose / 목적 | Functions / 함수 | Lines / 줄 수 |
|-------------|---------------|-----------------|--------------|
| `maputil.go` | Package documentation, type definitions / 패키지 문서, 타입 정의 | Types: Entry, Number, Ordered | ~127 |
| `basic.go` | Get, set, delete, has operations / 가져오기, 설정, 삭제, 확인 작업 | 11 functions | ~225 |
| `transform.go` | Map, filter, invert, flatten / Map, 필터, 반전, 평탄화 | 10 functions | ~265 |
| `aggregate.go` | Reduce, sum, min, max, groupby / Reduce, 합계, 최소, 최대, 그룹화 | 9 functions | ~296 |
| `merge.go` | Merge, union, intersection / 병합, 합집합, 교집합 | 8 functions | ~247 |
| `filter.go` | Filter, omit, pick operations / 필터, 생략, 선택 작업 | 7 functions | ~177 |
| `convert.go` | Keys, values, entries, JSON / 키, 값, 항목, JSON | 9 functions | ~209 |
| `predicate.go` | Every, some, none, has checks / 모두, 어느, 없음, 확인 검사 | 7 functions | ~172 |
| `keys.go` | Key manipulation operations / 키 조작 작업 | 8 functions | ~235 |
| `values.go` | Value manipulation operations / 값 조작 작업 | 7 functions | ~228 |
| `comparison.go` | Diff, compare, common keys / 차이, 비교, 공통 키 | 6 functions | ~283 |
| `*_test.go` | Comprehensive tests for each file / 각 파일에 대한 종합 테스트 | Test functions | ~3,500 total |

**Total Package Size / 전체 패키지 크기**: ~5,964 lines (implementation + tests) / ~5,964줄 (구현 + 테스트)

---

## Core Components / 핵심 컴포넌트

### 1. Type Constraints and Structs / 타입 제약 및 구조체

**Location / 위치**: `maputil/maputil.go`

The package defines type constraints and structured types for generics:

패키지는 제네릭을 위한 타입 제약 및 구조화된 타입을 정의합니다:

```go
// Entry represents a key-value pair in a map
// Entry는 맵의 키-값 쌍을 나타냅니다
type Entry[K comparable, V any] struct {
    Key   K
    Value V
}

// Number constraint for numeric operations (arithmetic, statistical)
// 숫자 작업을 위한 Number 제약 (산술, 통계)
type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
        ~float32 | ~float64
}

// Ordered constraint for sorting and comparison operations
// 정렬 및 비교 작업을 위한 Ordered 제약
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
        ~float32 | ~float64 | ~string
}
```

**Design Rationale / 설계 근거**:
- `Entry[K comparable, V any]`: Generic struct for representing key-value pairs / 키-값 쌍을 나타내는 제네릭 구조체
  - Used by `Entries()` and `FromEntries()` functions / `Entries()` 및 `FromEntries()` 함수에서 사용
  - Keys must be `comparable` (map requirement) / 키는 `comparable`이어야 함 (맵 요구사항)
  - Values can be `any` type / 값은 `any` 타입 가능
- `Number`: Used for arithmetic operations (Sum, Average) / 산술 작업에 사용 (Sum, Average)
  - Includes all integer and floating-point types / 모든 정수 및 부동 소수점 타입 포함
  - Uses `~` for underlying type matching / 기본 타입 매칭에 `~` 사용
- `Ordered`: Used for comparison and sorting (Min, Max, KeysSorted) / 비교 및 정렬에 사용 (Min, Max, KeysSorted)
  - Includes numeric types and strings / 숫자 타입 및 문자열 포함
- `comparable`: Built-in Go constraint for equality checks / 동등성 검사를 위한 내장 Go 제약
- `any`: Used for type-agnostic operations (Map, Filter, ForEach) / 타입에 구애받지 않는 작업에 사용 (Map, Filter, ForEach)

### 2. Function Categories / 함수 카테고리

#### Category 1: Basic Operations (11 functions) / 기본 작업 (11개 함수)

**File / 파일**: `maputil/basic.go`

**Purpose / 목적**: Fundamental map operations for accessing, modifying, and checking entries.

기본적인 맵 작업: 항목 접근, 수정, 확인.

**Key Functions / 주요 함수**:
- `Get[K comparable, V any](m map[K]V, key K) (V, bool)`
- `GetOr[K comparable, V any](m map[K]V, key K, defaultValue V) V`
- `Set[K comparable, V any](m map[K]V, key K, value V) map[K]V`
- `Delete[K comparable, V any](m map[K]V, keys ...K) map[K]V`
- `Has[K comparable, V any](m map[K]V, key K) bool`
- `IsEmpty[K comparable, V any](m map[K]V) bool`
- `Clone[K comparable, V any](m map[K]V) map[K]V`
- `Equal[K comparable, V comparable](m1, m2 map[K]V) bool`

**Implementation Pattern / 구현 패턴**:
```go
// Get: Safe key lookup with existence check
// Get: 존재 확인과 함께 안전한 키 조회
func Get[K comparable, V any](m map[K]V, key K) (V, bool) {
    value, ok := m[key]
    return value, ok
}

// Clone: Create shallow copy (immutability pattern)
// Clone: 얕은 복사본 생성 (불변성 패턴)
func Clone[K comparable, V any](m map[K]V) map[K]V {
    if m == nil {
        return nil
    }
    result := make(map[K]V, len(m))
    for k, v := range m {
        result[k] = v
    }
    return result
}
```

**Time Complexity / 시간 복잡도**: O(1) for Get/Has, O(n) for Clone/Delete/Equal / Get/Has는 O(1), Clone/Delete/Equal은 O(n)

#### Category 2: Transformation (10 functions) / 변환 (10개 함수)

**File / 파일**: `maputil/transform.go`

**Purpose / 목적**: Transform maps into different forms (keys, values, or both).

맵을 다양한 형태로 변환 (키, 값, 또는 둘 다).

**Key Functions / 주요 함수**:
- `Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R`
- `MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V`
- `MapValues[K comparable, V any, R any](m map[K]V, fn func(V) R) map[K]R`
- `MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, fn func(K1, V1) (K2, V2)) map[K2]V2`
- `Invert[K comparable, V comparable](m map[K]V) map[V]K`
- `Flatten[K comparable, V any](m map[K]map[string]V, delimiter string) map[string]V`
- `Unflatten[V any](m map[string]V, delimiter string) map[string]interface{}`
- `Chunk[K comparable, V any](m map[K]V, size int) []map[K]V`
- `Partition[K comparable, V any](m map[K]V, fn func(K, V) bool) (map[K]V, map[K]V)`
- `Compact[K comparable, V comparable](m map[K]V) map[K]V`

**Implementation Pattern / 구현 패턴**:
```go
// Map: Transform values using key-value function
// Map: 키-값 함수를 사용하여 값 변환
func Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R {
    result := make(map[K]R, len(m))
    for k, v := range m {
        result[k] = fn(k, v)
    }
    return result
}

// Invert: Swap keys and values
// Invert: 키와 값 교환
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
    result := make(map[V]K, len(m))
    for k, v := range m {
        result[v] = k
    }
    return result
}

// Flatten: Convert nested map to flat map with delimiter
// Flatten: 구분자를 사용하여 중첩 맵을 평면 맵으로 변환
func Flatten[K comparable, V any](m map[K]map[string]V, delimiter string) map[string]V {
    result := make(map[string]V)
    for k, nested := range m {
        prefix := fmt.Sprintf("%v", k)
        for nestedKey, nestedValue := range nested {
            flatKey := prefix + delimiter + nestedKey
            result[flatKey] = nestedValue
        }
    }
    return result
}
```

**Time Complexity / 시간 복잡도**: O(n) for most, O(n*m) for Flatten/Unflatten / 대부분 O(n), Flatten/Unflatten은 O(n*m)

#### Category 3: Aggregation (9 functions) / 집계 (9개 함수)

**File / 파일**: `maputil/aggregate.go`

**Purpose / 목적**: Aggregate data from maps (reduce, sum, grouping).

맵에서 데이터 집계 (reduce, sum, 그룹화).

**Key Functions / 주요 함수**:
- `Reduce[K comparable, V any, R any](m map[K]V, initial R, fn func(R, K, V) R) R`
- `Sum[K comparable, V Number](m map[K]V) V`
- `Min[K comparable, V Ordered](m map[K]V) (K, V, bool)`
- `Max[K comparable, V Ordered](m map[K]V) (K, V, bool)`
- `MinBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool)`
- `MaxBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool)`
- `Average[K comparable, V Number](m map[K]V) float64`
- `GroupBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G][]V`
- `CountBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G]int`

**Implementation Pattern / 구현 패턴**:
```go
// Reduce: Flexible aggregation pattern
// Reduce: 유연한 집계 패턴
func Reduce[K comparable, V any, R any](m map[K]V, initial R, fn func(R, K, V) R) R {
    result := initial
    for k, v := range m {
        result = fn(result, k, v)
    }
    return result
}

// Sum: Specialized aggregation for numbers
// Sum: 숫자를 위한 특수화된 집계
func Sum[K comparable, V Number](m map[K]V) V {
    var sum V
    for _, v := range m {
        sum += v
    }
    return sum
}

// Min: Find key-value pair with minimum value
// Min: 최소값을 가진 키-값 쌍 찾기
func Min[K comparable, V Ordered](m map[K]V) (K, V, bool) {
    if len(m) == 0 {
        var zeroK K
        var zeroV V
        return zeroK, zeroV, false
    }

    var minKey K
    var minValue V
    first := true

    for k, v := range m {
        if first || v < minValue {
            minKey = k
            minValue = v
            first = false
        }
    }

    return minKey, minValue, true
}

// GroupBy: Group slice elements by key
// GroupBy: 키로 슬라이스 요소 그룹화
func GroupBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G][]V {
    result := make(map[G][]V)

    for _, item := range slice {
        key := fn(item)
        result[key] = append(result[key], item)
    }

    return result
}
```

**Time Complexity / 시간 복잡도**: O(n) for all aggregation functions / 모든 집계 함수에 대해 O(n)

#### Category 4: Merge Operations (8 functions) / 병합 작업 (8개 함수)

**File / 파일**: `maputil/merge.go`

**Purpose / 목적**: Merge multiple maps using various strategies.

다양한 전략을 사용하여 여러 맵 병합.

**Key Functions / 주요 함수**:
- `Merge[K comparable, V any](maps ...map[K]V) map[K]V`
- `MergeWith[K comparable, V any](fn func(V, V) V, maps ...map[K]V) map[K]V`
- `DeepMerge(maps ...map[string]interface{}) map[string]interface{}`
- `Union[K comparable, V any](maps ...map[K]V) map[K]V`
- `Intersection[K comparable, V any](maps ...map[K]V) map[K]V`
- `Difference[K comparable, V any](m1, m2 map[K]V) map[K]V`
- `SymmetricDifference[K comparable, V any](m1, m2 map[K]V) map[K]V`
- `Assign[K comparable, V any](target map[K]V, sources ...map[K]V) map[K]V`

**Implementation Pattern / 구현 패턴**:
```go
// Merge: Last value wins for duplicate keys
// Merge: 중복 키에 대해 마지막 값이 우선
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
    totalSize := 0
    for _, m := range maps {
        totalSize += len(m)
    }

    result := make(map[K]V, totalSize)
    for _, m := range maps {
        for k, v := range m {
            result[k] = v
        }
    }

    return result
}

// MergeWith: Custom resolver for conflicts
// MergeWith: 충돌에 대한 사용자 정의 해결자
func MergeWith[K comparable, V any](fn func(V, V) V, maps ...map[K]V) map[K]V {
    totalSize := 0
    for _, m := range maps {
        totalSize += len(m)
    }

    result := make(map[K]V, totalSize)
    for _, m := range maps {
        for k, v := range m {
            if existing, exists := result[k]; exists {
                result[k] = fn(existing, v)
            } else {
                result[k] = v
            }
        }
    }

    return result
}

// Intersection: Keys that exist in ALL maps
// Intersection: 모든 맵에 존재하는 키
func Intersection[K comparable, V any](maps ...map[K]V) map[K]V {
    if len(maps) == 0 {
        return make(map[K]V)
    }
    if len(maps) == 1 {
        return Clone(maps[0])
    }

    result := make(map[K]V)

    // Iterate over the first map
    for k, v := range maps[0] {
        existsInAll := true

        // Check if key exists in all other maps
        for i := 1; i < len(maps); i++ {
            if _, exists := maps[i][k]; !exists {
                existsInAll = false
                break
            }
        }

        if existsInAll {
            result[k] = v
        }
    }

    return result
}

// Assign: MUTATING operation (WARNING)
// Assign: 변경 작업 (경고)
func Assign[K comparable, V any](target map[K]V, sources ...map[K]V) map[K]V {
    for _, source := range sources {
        for k, v := range source {
            target[k] = v
        }
    }
    return target
}
```

**Design Note / 설계 참고**: `Assign` is the ONLY mutating operation in maputil. All other functions return new maps. / `Assign`은 maputil의 유일한 변경 작업입니다. 다른 모든 함수는 새 맵을 반환합니다.

**Time Complexity / 시간 복잡도**: O(n*m) for multi-map operations / 다중 맵 작업에 대해 O(n*m)

#### Category 5: Filter Operations (7 functions) / 필터 작업 (7개 함수)

**File / 파일**: `maputil/filter.go`

**Purpose / 목적**: Filter maps based on predicates or key selections.

조건자 또는 키 선택에 따라 맵 필터링.

**Key Functions / 주요 함수**:
- `Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V`
- `FilterKeys[K comparable, V any](m map[K]V, fn func(K) bool) map[K]V`
- `FilterValues[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V`
- `Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V`
- `Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V`
- `OmitBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V`
- `PickBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V`

**Implementation Pattern / 구현 패턴**:
```go
// Filter: Keep entries that satisfy predicate
// Filter: 조건자를 만족하는 항목 유지
func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
    result := make(map[K]V)

    for k, v := range m {
        if fn(k, v) {
            result[k] = v
        }
    }

    return result
}

// Pick: Select specific keys
// Pick: 특정 키 선택
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V {
    result := make(map[K]V, len(keys))

    for _, key := range keys {
        if value, exists := m[key]; exists {
            result[key] = value
        }
    }

    return result
}

// Omit: Exclude specific keys
// Omit: 특정 키 제외
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V {
    if len(keys) == 0 {
        return Clone(m)
    }

    // Create hash set for O(1) lookup
    toOmit := make(map[K]struct{}, len(keys))
    for _, key := range keys {
        toOmit[key] = struct{}{}
    }

    result := make(map[K]V)
    for k, v := range m {
        if _, shouldOmit := toOmit[k]; !shouldOmit {
            result[k] = v
        }
    }

    return result
}
```

**Time Complexity / 시간 복잡도**: O(n) for Filter operations, O(k) for Pick (k = number of keys) / Filter 작업은 O(n), Pick은 O(k) (k = 키 개수)

#### Category 6: Conversion (9 functions) / 변환 (9개 함수)

**File / 파일**: `maputil/convert.go`

**Purpose / 목적**: Convert between maps, slices, entries, and JSON.

맵, 슬라이스, 항목, JSON 간 변환.

**Key Functions / 주요 함수**:
- `Keys[K comparable, V any](m map[K]V) []K`
- `Values[K comparable, V any](m map[K]V) []V`
- `Entries[K comparable, V any](m map[K]V) []Entry[K, V]`
- `FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V`
- `FromSlice[K comparable, V any](slice []V, fn func(V) K) map[K]V`
- `FromSliceBy[K comparable, V any, R any](slice []V, keyFn func(V) K, valueFn func(V) R) map[K]R`
- `ToSlice[K comparable, V any, R any](m map[K]V, fn func(K, V) R) []R`
- `ToJSON[K comparable, V any](m map[K]V) (string, error)`
- `FromJSON[K comparable, V any](jsonStr string, m *map[K]V) error`

**Implementation Pattern / 구현 패턴**:
```go
// Keys: Extract all keys as slice
// Keys: 모든 키를 슬라이스로 추출
func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

// Entries: Convert map to slice of Entry structs
// Entries: 맵을 Entry 구조체의 슬라이스로 변환
func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
    entries := make([]Entry[K, V], 0, len(m))
    for k, v := range m {
        entries = append(entries, Entry[K, V]{Key: k, Value: v})
    }
    return entries
}

// FromEntries: Create map from Entry slice
// FromEntries: Entry 슬라이스로부터 맵 생성
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
    result := make(map[K]V, len(entries))
    for _, entry := range entries {
        result[entry.Key] = entry.Value
    }
    return result
}

// FromSlice: Create map from slice using key extractor
// FromSlice: 키 추출기를 사용하여 슬라이스로부터 맵 생성
func FromSlice[K comparable, V any](slice []V, fn func(V) K) map[K]V {
    result := make(map[K]V, len(slice))
    for _, item := range slice {
        key := fn(item)
        result[key] = item
    }
    return result
}

// ToJSON: Marshal map to JSON string
// ToJSON: 맵을 JSON 문자열로 마샬링
func ToJSON[K comparable, V any](m map[K]V) (string, error) {
    bytes, err := json.Marshal(m)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}
```

**Time Complexity / 시간 복잡도**: O(n) for all conversion operations / 모든 변환 작업에 대해 O(n)

#### Category 7: Predicate Checks (7 functions) / 조건 검사 (7개 함수)

**File / 파일**: `maputil/predicate.go`

**Purpose / 목적**: Test conditions on map entries.

맵 항목에 대한 조건 테스트.

**Key Functions / 주요 함수**:
- `Every[K comparable, V any](m map[K]V, fn func(K, V) bool) bool`
- `Some[K comparable, V any](m map[K]V, fn func(K, V) bool) bool`
- `None[K comparable, V any](m map[K]V, fn func(K, V) bool) bool`
- `HasKey[K comparable, V any](m map[K]V, key K) bool`
- `HasValue[K comparable, V comparable](m map[K]V, value V) bool`
- `HasEntry[K comparable, V comparable](m map[K]V, key K, value V) bool`
- `IsSubset[K comparable, V comparable](subset, superset map[K]V) bool`
- `IsSuperset[K comparable, V comparable](superset, subset map[K]V) bool`

**Implementation Pattern / 구현 패턴**:
```go
// Every: Check if all entries satisfy predicate
// Every: 모든 항목이 조건자를 만족하는지 확인
func Every[K comparable, V any](m map[K]V, fn func(K, V) bool) bool {
    for k, v := range m {
        if !fn(k, v) {
            return false
        }
    }
    return true
}

// Some: Check if at least one entry satisfies predicate
// Some: 최소 하나의 항목이 조건자를 만족하는지 확인
func Some[K comparable, V any](m map[K]V, fn func(K, V) bool) bool {
    for k, v := range m {
        if fn(k, v) {
            return true
        }
    }
    return false
}

// HasValue: Check if value exists in map (O(n) search)
// HasValue: 맵에 값이 존재하는지 확인 (O(n) 검색)
func HasValue[K comparable, V comparable](m map[K]V, value V) bool {
    for _, v := range m {
        if v == value {
            return true
        }
    }
    return false
}

// IsSubset: Check if all key-value pairs exist in superset
// IsSubset: 모든 키-값 쌍이 상위집합에 존재하는지 확인
func IsSubset[K comparable, V comparable](subset, superset map[K]V) bool {
    if len(subset) > len(superset) {
        return false
    }

    for k, v := range subset {
        if superV, exists := superset[k]; !exists || v != superV {
            return false
        }
    }

    return true
}
```

**Time Complexity / 시간 복잡도**: O(1) for HasKey, O(n) for others with short-circuit / HasKey는 O(1), 나머지는 단락 평가와 함께 O(n)

#### Category 8: Key Operations (8 functions) / 키 작업 (8개 함수)

**File / 파일**: `maputil/keys.go`

**Purpose / 목적**: Manipulate and transform map keys.

맵 키 조작 및 변환.

**Key Functions / 주요 함수**:
- `KeysSorted[K Ordered, V any](m map[K]V) []K`
- `KeysBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []K`
- `RenameKey[K comparable, V any](m map[K]V, oldKey, newKey K) map[K]V`
- `RenameKeys[K comparable, V any](m map[K]V, mapping map[K]K) map[K]V`
- `SwapKeys[K comparable, V any](m map[K]V, key1, key2 K) map[K]V`
- `FindKey[K comparable, V any](m map[K]V, fn func(K, V) bool) (K, bool)`
- `PrefixKeys[V any](m map[string]V, prefix string) map[string]V`
- `TransformKeys[K comparable, V any](m map[K]V, fn func(K) K) map[K]V`

**Implementation Pattern / 구현 패턴**:
```go
// KeysSorted: Return keys in sorted order
// KeysSorted: 정렬된 순서로 키 반환
func KeysSorted[K Ordered, V any](m map[K]V) []K {
    keys := Keys(m)
    sort.Slice(keys, func(i, j int) bool {
        return keys[i] < keys[j]
    })
    return keys
}

// RenameKey: Rename a single key
// RenameKey: 단일 키 이름 변경
func RenameKey[K comparable, V any](m map[K]V, oldKey, newKey K) map[K]V {
    result := Clone(m)

    if value, exists := result[oldKey]; exists {
        delete(result, oldKey)
        result[newKey] = value
    }

    return result
}

// SwapKeys: Exchange two keys' values
// SwapKeys: 두 키의 값 교환
func SwapKeys[K comparable, V any](m map[K]V, key1, key2 K) map[K]V {
    result := Clone(m)

    val1, exists1 := result[key1]
    val2, exists2 := result[key2]

    if exists1 && exists2 {
        result[key1] = val2
        result[key2] = val1
    }

    return result
}

// PrefixKeys: Add prefix to all string keys
// PrefixKeys: 모든 문자열 키에 접두사 추가
func PrefixKeys[V any](m map[string]V, prefix string) map[string]V {
    result := make(map[string]V, len(m))
    for k, v := range m {
        result[prefix+k] = v
    }
    return result
}
```

**Time Complexity / 시간 복잡도**: O(n log n) for KeysSorted, O(n) for most others / KeysSorted는 O(n log n), 대부분은 O(n)

#### Category 9: Value Operations (7 functions) / 값 작업 (7개 함수)

**File / 파일**: `maputil/values.go`

**Purpose / 목적**: Manipulate and analyze map values.

맵 값 조작 및 분석.

**Key Functions / 주요 함수**:
- `ValuesSorted[K comparable, V Ordered](m map[K]V) []V`
- `ValuesBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []V`
- `UniqueValues[K comparable, V comparable](m map[K]V) []V`
- `FindValue[K comparable, V any](m map[K]V, fn func(K, V) bool) (V, bool)`
- `ReplaceValue[K comparable, V comparable](m map[K]V, oldValue, newValue V) map[K]V`
- `UpdateValues[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V`
- `MinValue[K comparable, V Ordered](m map[K]V) (V, bool)`
- `MaxValue[K comparable, V Ordered](m map[K]V) (V, bool)`

**Implementation Pattern / 구현 패턴**:
```go
// ValuesSorted: Return values in sorted order
// ValuesSorted: 정렬된 순서로 값 반환
func ValuesSorted[K comparable, V Ordered](m map[K]V) []V {
    values := Values(m)
    sort.Slice(values, func(i, j int) bool {
        return values[i] < values[j]
    })
    return values
}

// UniqueValues: Get unique values from map
// UniqueValues: 맵에서 고유한 값 가져오기
func UniqueValues[K comparable, V comparable](m map[K]V) []V {
    seen := make(map[V]struct{})
    result := make([]V, 0)

    for _, v := range m {
        if _, exists := seen[v]; !exists {
            seen[v] = struct{}{}
            result = append(result, v)
        }
    }

    return result
}

// ReplaceValue: Replace all occurrences of a value
// ReplaceValue: 특정 값의 모든 발생 대체
func ReplaceValue[K comparable, V comparable](m map[K]V, oldValue, newValue V) map[K]V {
    result := make(map[K]V, len(m))
    for k, v := range m {
        if v == oldValue {
            result[k] = newValue
        } else {
            result[k] = v
        }
    }
    return result
}

// UpdateValues: Transform all values with function
// UpdateValues: 함수로 모든 값 변환
func UpdateValues[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V {
    result := make(map[K]V, len(m))
    for k, v := range m {
        result[k] = fn(k, v)
    }
    return result
}
```

**Time Complexity / 시간 복잡도**: O(n log n) for ValuesSorted, O(n) for most others / ValuesSorted는 O(n log n), 대부분은 O(n)

#### Category 10: Comparison (6 functions) / 비교 (6개 함수)

**File / 파일**: `maputil/comparison.go`

**Purpose / 목적**: Compare maps and find differences.

맵을 비교하고 차이점 찾기.

**Key Functions / 주요 함수**:
- `EqualMaps[K comparable, V comparable](m1, m2 map[K]V) bool`
- `EqualFunc[K comparable, V any](m1, m2 map[K]V, eq func(V, V) bool) bool`
- `Diff[K comparable, V comparable](m1, m2 map[K]V) map[K]V`
- `DiffKeys[K comparable, V comparable](m1, m2 map[K]V) []K`
- `CommonKeys[K comparable, V any](maps ...map[K]V) []K`
- `AllKeys[K comparable, V any](maps ...map[K]V) []K`
- `Compare[K comparable, V comparable](m1, m2 map[K]V) (added, removed, modified map[K]V)`

**Implementation Pattern / 구현 패턴**:
```go
// Diff: Find all differences between two maps
// Diff: 두 맵 간의 모든 차이점 찾기
func Diff[K comparable, V comparable](m1, m2 map[K]V) map[K]V {
    result := make(map[K]V)

    // Check m1 against m2
    for k, v1 := range m1 {
        if v2, exists := m2[k]; !exists || v1 != v2 {
            if exists {
                result[k] = v2 // Use value from m2
            } else {
                result[k] = v1 // Use value from m1
            }
        }
    }

    // Check m2 for keys not in m1
    for k, v2 := range m2 {
        if _, exists := m1[k]; !exists {
            result[k] = v2
        }
    }

    return result
}

// Compare: Detailed comparison with three result maps
// Compare: 세 개의 결과 맵으로 상세 비교
func Compare[K comparable, V comparable](m1, m2 map[K]V) (added, removed, modified map[K]V) {
    added = make(map[K]V)
    removed = make(map[K]V)
    modified = make(map[K]V)

    // Check m1 against m2
    for k, v1 := range m1 {
        if v2, exists := m2[k]; exists {
            if v1 != v2 {
                modified[k] = v2
            }
        } else {
            removed[k] = v1
        }
    }

    // Check m2 for keys not in m1
    for k, v2 := range m2 {
        if _, exists := m1[k]; !exists {
            added[k] = v2
        }
    }

    return added, removed, modified
}

// CommonKeys: Find keys that exist in ALL maps
// CommonKeys: 모든 맵에 존재하는 키 찾기
func CommonKeys[K comparable, V any](maps ...map[K]V) []K {
    if len(maps) == 0 {
        return []K{}
    }
    if len(maps) == 1 {
        return Keys(maps[0])
    }

    common := make([]K, 0)

    // Iterate over the first map
    for k := range maps[0] {
        existsInAll := true

        // Check if key exists in all other maps
        for i := 1; i < len(maps); i++ {
            if _, exists := maps[i][k]; !exists {
                existsInAll = false
                break
            }
        }

        if existsInAll {
            common = append(common, k)
        }
    }

    return common
}
```

**Time Complexity / 시간 복잡도**: O(n+m) for Diff/Compare, O(n*m) for CommonKeys / Diff/Compare는 O(n+m), CommonKeys는 O(n*m)

---

## Design Patterns / 디자인 패턴

### Pattern 1: Generic Type Parameters / 제네릭 타입 매개변수

**Purpose / 목적**: Provide type-safe operations without code duplication.

코드 중복 없이 타입 안전 작업 제공.

**Implementation / 구현**:
```go
// Single generic type for key and value
// 키와 값에 대한 단일 제네릭 타입
func Clone[K comparable, V any](m map[K]V) map[K]V {
    // K must be comparable (map requirement)
    // V can be any type
}

// Multiple generic types (different input/output)
// 여러 제네릭 타입 (다른 입력/출력)
func Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R {
    // K: key type (must be comparable)
    // V: input value type
    // R: output value type
}

// Constrained generic type (only numeric values)
// 제약된 제네릭 타입 (숫자 값만)
func Sum[K comparable, V Number](m map[K]V) V {
    // V must satisfy Number constraint
}
```

**Benefits / 이점**:
- Type safety at compile time / 컴파일 타임에 타입 안전성
- No runtime type assertions / 런타임 타입 단언 없음
- Code reusability / 코드 재사용성
- Better IDE support and autocomplete / 더 나은 IDE 지원 및 자동 완성

### Pattern 2: Functional Programming with Higher-Order Functions / 고차 함수를 사용한 함수형 프로그래밍

**Purpose / 목적**: Enable functional programming patterns for maps.

맵에 대한 함수형 프로그래밍 패턴 활성화.

**Implementation / 구현**:
```go
// Higher-order function: takes function as parameter
// 고차 함수: 함수를 매개변수로 받음
func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
    result := make(map[K]V)
    for k, v := range m {
        if fn(k, v) {
            result[k] = v
        }
    }
    return result
}

// Usage: pass lambda/anonymous function
// 사용: 람다/익명 함수 전달
data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
filtered := Filter(data, func(k string, v int) bool {
    return v > 2
}) // map[string]int{"c": 3, "d": 4}
```

**Benefits / 이점**:
- Declarative code (what, not how) / 선언적 코드 (무엇을, 어떻게가 아니라)
- Composable operations / 조합 가능한 작업
- Immutable transformations / 불변 변환
- Chainable operations (via intermediate maps) / 체이닝 가능한 작업 (중간 맵을 통해)

### Pattern 3: Immutability Pattern / 불변성 패턴

**Purpose / 목적**: Ensure original maps are never modified (except Assign).

원본 맵이 절대 수정되지 않도록 보장 (Assign 제외).

**Implementation / 구현**:
```go
// Always return new map (immutable)
// 항상 새 맵 반환 (불변)
func Set[K comparable, V any](m map[K]V, key K, value V) map[K]V {
    result := make(map[K]V, len(m)+1)
    for k, v := range m {
        result[k] = v
    }
    result[key] = value
    return result
}

// Exception: Assign is MUTATING (clearly documented)
// 예외: Assign은 변경 작업 (명확하게 문서화됨)
func Assign[K comparable, V any](target map[K]V, sources ...map[K]V) map[K]V {
    // WARNING: Modifies target map in place
    // 경고: 대상 맵을 직접 수정
    for _, source := range sources {
        for k, v := range source {
            target[k] = v
        }
    }
    return target
}
```

**Benefits / 이점**:
- Predictable behavior / 예측 가능한 동작
- Safe concurrent reads (original unchanged) / 안전한 동시 읽기 (원본 변경 없음)
- No side effects / 부작용 없음
- Easier debugging / 더 쉬운 디버깅

### Pattern 4: Safe Error Handling / 안전한 에러 처리

**Purpose / 목적**: Handle edge cases gracefully without panics.

패닉 없이 엣지 케이스를 우아하게 처리.

**Implementation / 구현**:
```go
// Return zero values and false for empty maps
// 빈 맵에 대해 zero 값과 false 반환
func Min[K comparable, V Ordered](m map[K]V) (K, V, bool) {
    if len(m) == 0 {
        var zeroK K
        var zeroV V
        return zeroK, zeroV, false // Safe return
    }
    // ... find min
}

// Return nil for nil maps
// nil 맵에 대해 nil 반환
func Clone[K comparable, V any](m map[K]V) map[K]V {
    if m == nil {
        return nil
    }
    // ... clone logic
}

// Return empty results for invalid input
// 잘못된 입력에 대해 빈 결과 반환
func Chunk[K comparable, V any](m map[K]V, size int) []map[K]V {
    if size <= 0 || len(m) == 0 {
        return []map[K]V{}
    }
    // ... chunk logic
}
```

**Benefits / 이점**:
- No panics, predictable behavior / 패닉 없음, 예측 가능한 동작
- Easy to check for errors / 에러 확인 용이
- Follows Go idioms / Go 관용구 준수

---

## Internal Implementation / 내부 구현

### Implementation Example 1: Map Function / Map 함수

**Purpose / 목적**: Transform map values using a function.

함수를 사용하여 맵 값 변환.

**Code / 코드**:
```go
func Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R {
    result := make(map[K]R, len(m))
    for k, v := range m {
        result[k] = fn(k, v)
    }
    return result
}
```

**Flow Diagram / 흐름 다이어그램**:
```
Input: map[K]V, func(K,V) R
    ↓
1. Create new map[K]R with capacity len(m)
   새 map[K]R 생성 (용량 len(m))
    ↓
2. Iterate over input map
   입력 맵 반복
    ↓
3. For each (k, v):
   각 (k, v)에 대해:
   - Apply function: fn(k, v) → R
     함수 적용: fn(k, v) → R
   - Store in result[k]
     result[k]에 저장
    ↓
4. Return new map[K]R
   새 map[K]R 반환
```

**Complexity / 복잡도**:
- Time: O(n) where n = len(m) / 시간: O(n) 여기서 n = len(m)
- Space: O(n) for result map / 공간: 결과 맵에 대해 O(n)

### Implementation Example 2: Filter Function / Filter 함수

**Purpose / 목적**: Keep only entries that satisfy a predicate.

조건자를 만족하는 항목만 유지.

**Code / 코드**:
```go
func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
    result := make(map[K]V)

    for k, v := range m {
        if fn(k, v) {
            result[k] = v
        }
    }

    return result
}
```

**Flow Diagram / 흐름 다이어그램**:
```
Input: map[K]V, func(K,V) bool
    ↓
1. Create empty result map
   빈 결과 맵 생성
    ↓
2. Iterate over input map
   입력 맵 반복
    ↓
3. For each (k, v):
   각 (k, v)에 대해:
   - Evaluate: fn(k, v) → bool
     평가: fn(k, v) → bool
   - If true: add (k,v) to result
     true이면: (k,v)를 결과에 추가
   - If false: skip
     false이면: 건너뜀
    ↓
4. Return filtered map
   필터링된 맵 반환
```

**Complexity / 복잡도**:
- Time: O(n) where n = len(m) / 시간: O(n) 여기서 n = len(m)
- Space: O(k) where k = number of matching entries / 공간: O(k) 여기서 k = 일치하는 항목 수

### Implementation Example 3: Reduce Function / Reduce 함수

**Purpose / 목적**: Aggregate map into a single value.

맵을 단일 값으로 집계.

**Code / 코드**:
```go
func Reduce[K comparable, V any, R any](m map[K]V, initial R, fn func(R, K, V) R) R {
    result := initial
    for k, v := range m {
        result = fn(result, k, v)
    }
    return result
}
```

**Flow Diagram / 흐름 다이어그램**:
```
Input: map[K]V, initial R, func(R,K,V) R
    ↓
1. Set accumulator = initial
   누산기 = initial로 설정
    ↓
2. Iterate over input map
   입력 맵 반복
    ↓
3. For each (k, v):
   각 (k, v)에 대해:
   - Update accumulator: fn(acc, k, v) → acc'
     누산기 업데이트: fn(acc, k, v) → acc'
    ↓
4. Return final accumulator
   최종 누산기 반환
```

**Example Usage / 사용 예제**:
```go
// Sum all values
// 모든 값 합산
m := map[string]int{"a": 1, "b": 2, "c": 3}
sum := Reduce(m, 0, func(acc int, k string, v int) int {
    return acc + v
}) // sum = 6

// Concatenate all keys
// 모든 키 연결
keys := Reduce(m, "", func(acc string, k string, v int) string {
    return acc + k
}) // keys = "abc" (order may vary)
```

**Complexity / 복잡도**:
- Time: O(n) where n = len(m) / 시간: O(n) 여기서 n = len(m)
- Space: O(1) (no extra space) / 공간: O(1) (추가 공간 없음)

### Implementation Example 4: GroupBy Function / GroupBy 함수

**Purpose / 목적**: Group slice elements by key into map.

키로 슬라이스 요소를 맵으로 그룹화.

**Code / 코드**:
```go
func GroupBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G][]V {
    result := make(map[G][]V)

    for _, item := range slice {
        key := fn(item)
        result[key] = append(result[key], item)
    }

    return result
}
```

**Flow Diagram / 흐름 다이어그램**:
```
Input: []V, func(V) G
    ↓
1. Create empty map[G][]V
   빈 map[G][]V 생성
    ↓
2. Iterate over slice
   슬라이스 반복
    ↓
3. For each item:
   각 item에 대해:
   - Extract key: fn(item) → G
     키 추출: fn(item) → G
   - Append item to result[key]
     item을 result[key]에 추가
    ↓
4. Return grouped map
   그룹화된 맵 반환
```

**Example Usage / 사용 예제**:
```go
type User struct {
    Name string
    City string
}

users := []User{
    {Name: "Alice", City: "Seoul"},
    {Name: "Bob", City: "Seoul"},
    {Name: "Charlie", City: "Busan"},
}

byCity := GroupBy(users, func(u User) string {
    return u.City
})
// byCity = map[string][]User{
//     "Seoul": []User{{Name: "Alice", ...}, {Name: "Bob", ...}},
//     "Busan": []User{{Name: "Charlie", ...}},
// }
```

**Complexity / 복잡도**:
- Time: O(n) where n = len(slice) / 시간: O(n) 여기서 n = len(slice)
- Space: O(n) for result map and slices / 공간: 결과 맵 및 슬라이스에 대해 O(n)

### Implementation Example 5: Intersection Function / Intersection 함수

**Purpose / 목적**: Find keys that exist in ALL input maps.

모든 입력 맵에 존재하는 키 찾기.

**Code / 코드**:
```go
func Intersection[K comparable, V any](maps ...map[K]V) map[K]V {
    if len(maps) == 0 {
        return make(map[K]V)
    }
    if len(maps) == 1 {
        return Clone(maps[0])
    }

    result := make(map[K]V)

    // Iterate over the first map
    for k, v := range maps[0] {
        existsInAll := true

        // Check if key exists in all other maps
        for i := 1; i < len(maps); i++ {
            if _, exists := maps[i][k]; !exists {
                existsInAll = false
                break
            }
        }

        if existsInAll {
            result[k] = v
        }
    }

    return result
}
```

**Flow Diagram / 흐름 다이어그램**:
```
Input: ...map[K]V (variadic)
    ↓
1. Handle edge cases:
   엣지 케이스 처리:
   - 0 maps → return empty
     0개 맵 → 빈 맵 반환
   - 1 map → return clone
     1개 맵 → 복사본 반환
    ↓
2. Use first map as baseline
   첫 번째 맵을 기준으로 사용
    ↓
3. For each key in first map:
   첫 번째 맵의 각 키에 대해:
   - Check if exists in ALL other maps
     다른 모든 맵에 존재하는지 확인
   - If yes: add to result
     예이면: 결과에 추가
   - If no: skip
     아니오이면: 건너뜀
    ↓
4. Return intersection map
   교집합 맵 반환
```

**Complexity / 복잡도**:
- Time: O(n*m) where n = len(maps), m = avg map size / 시간: O(n*m) 여기서 n = len(maps), m = 평균 맵 크기
- Space: O(k) where k = size of intersection / 공간: O(k) 여기서 k = 교집합 크기

---

## Adding New Features / 새 기능 추가

### Step-by-Step Guide / 단계별 가이드

**Step 1: Identify the Category / 카테고리 식별**

Determine which of the 10 categories your function belongs to:

함수가 속한 10개 카테고리 중 하나를 결정합니다:

1. **Basic Operations** / 기본 작업: Get, Set, Has, Clone, etc. → `basic.go`
2. **Transformation** / 변환: Map, Invert, Flatten, etc. → `transform.go`
3. **Aggregation** / 집계: Reduce, Sum, GroupBy, etc. → `aggregate.go`
4. **Merge Operations** / 병합 작업: Merge, Union, Intersection, etc. → `merge.go`
5. **Filter Operations** / 필터 작업: Filter, Pick, Omit, etc. → `filter.go`
6. **Conversion** / 변환: Keys, Values, Entries, JSON, etc. → `convert.go`
7. **Predicate Checks** / 조건 검사: Every, Some, HasValue, etc. → `predicate.go`
8. **Key Operations** / 키 작업: RenameKey, PrefixKeys, etc. → `keys.go`
9. **Value Operations** / 값 작업: ReplaceValue, UniqueValues, etc. → `values.go`
10. **Comparison** / 비교: Diff, Compare, CommonKeys, etc. → `comparison.go`

**Step 2: Define Function Signature / 함수 시그니처 정의**

Choose appropriate generic constraints:

적절한 제네릭 제약 선택:

```go
// Example 1: Simple key-value operation
// 예제 1: 간단한 키-값 작업
func YourFunction[K comparable, V any](m map[K]V, ...) ReturnType

// Example 2: Numeric values only
// 예제 2: 숫자 값만
func YourFunction[K comparable, V Number](m map[K]V, ...) ReturnType

// Example 3: Ordered values (for sorting/comparison)
// 예제 3: 순서가 있는 값 (정렬/비교용)
func YourFunction[K comparable, V Ordered](m map[K]V, ...) ReturnType

// Example 4: Comparable values (for equality checks)
// 예제 4: 비교 가능한 값 (동등성 검사용)
func YourFunction[K comparable, V comparable](m map[K]V, ...) ReturnType

// Example 5: Transform to different type
// 예제 5: 다른 타입으로 변환
func YourFunction[K comparable, V any, R any](m map[K]V, ...) map[K]R
```

**Step 3: Write Implementation / 구현 작성**

Follow these guidelines:

다음 지침을 따르세요:

```go
// YourFunction does something useful with a map
// YourFunction은 맵으로 유용한 작업을 수행합니다
//
// Longer description of what it does, edge cases, behavior.
// 수행하는 작업, 엣지 케이스, 동작에 대한 더 긴 설명.
//
// Time complexity: O(?)
// 시간 복잡도: O(?)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	result := maputil.YourFunction(m, ...)
//	// result = ...
func YourFunction[K comparable, V any](m map[K]V, params ...) ReturnType {
    // 1. Handle edge cases first
    // 1. 엣지 케이스를 먼저 처리
    if len(m) == 0 {
        // Return appropriate zero/empty value
        // 적절한 zero/빈 값 반환
    }

    // 2. Pre-allocate result if possible
    // 2. 가능하면 결과를 미리 할당
    result := make(map[K]V, len(m))

    // 3. Implement core logic
    // 3. 핵심 로직 구현
    for k, v := range m {
        // Process each entry
        // 각 항목 처리
    }

    // 4. Return result (immutable - new map)
    // 4. 결과 반환 (불변 - 새 맵)
    return result
}
```

**Step 4: Write Tests / 테스트 작성**

Create comprehensive tests in `*_test.go`:

`*_test.go`에 종합 테스트 생성:

```go
func TestYourFunction(t *testing.T) {
    // Test 1: Normal case
    // 테스트 1: 정상 케이스
    t.Run("Normal case", func(t *testing.T) {
        m := map[string]int{"a": 1, "b": 2, "c": 3}
        result := YourFunction(m, ...)
        expected := map[string]int{...}

        if !reflect.DeepEqual(result, expected) {
            t.Errorf("Expected %v, got %v", expected, result)
        }
    })

    // Test 2: Empty map
    // 테스트 2: 빈 맵
    t.Run("Empty map", func(t *testing.T) {
        m := map[string]int{}
        result := YourFunction(m, ...)
        // Assert result is correct for empty input
    })

    // Test 3: Nil map
    // 테스트 3: nil 맵
    t.Run("Nil map", func(t *testing.T) {
        var m map[string]int
        result := YourFunction(m, ...)
        // Assert result handles nil gracefully
    })

    // Test 4: Single entry
    // 테스트 4: 단일 항목
    t.Run("Single entry", func(t *testing.T) {
        m := map[string]int{"a": 1}
        result := YourFunction(m, ...)
        // Assert correct behavior
    })

    // Test 5: Large map (performance)
    // 테스트 5: 큰 맵 (성능)
    t.Run("Large map", func(t *testing.T) {
        m := make(map[int]int, 10000)
        for i := 0; i < 10000; i++ {
            m[i] = i * 2
        }
        result := YourFunction(m, ...)
        // Assert correctness and performance
    })
}
```

**Step 5: Update Documentation / 문서 업데이트**

Update these files:

다음 파일을 업데이트하세요:

1. **Package README** (`maputil/README.md`):
   - Add function to category list / 카테고리 목록에 함수 추가
   - Add usage example / 사용 예제 추가
   - Update function count / 함수 개수 업데이트

2. **USER_MANUAL.md** (`docs/maputil/USER_MANUAL.md`):
   - Add detailed usage examples / 상세 사용 예제 추가
   - Add to appropriate section / 적절한 섹션에 추가

3. **DEVELOPER_GUIDE.md** (this file):
   - Update function count in overview / 개요에서 함수 개수 업데이트
   - Add implementation details if complex / 복잡한 경우 구현 세부 사항 추가

**Example: Adding a `Tap` Function / 예제: `Tap` 함수 추가**

```go
// File: maputil/util.go (if it exists, or create new category file)
// 파일: maputil/util.go (존재하면, 또는 새 카테고리 파일 생성)

// Tap executes a function for each entry and returns the original map unchanged.
// Tap은 각 항목에 대해 함수를 실행하고 변경되지 않은 원본 맵을 반환합니다.
//
// This is useful for debugging or side effects without modifying the map.
// 이것은 맵을 수정하지 않고 디버깅이나 부작용에 유용합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	result := maputil.Tap(m, func(k string, v int) {
//	    fmt.Printf("%s=%d\n", k, v)
//	})
//	// Prints: a=1, b=2
//	// Returns: same map (immutable - returns input)
func Tap[K comparable, V any](m map[K]V, fn func(K, V)) map[K]V {
    for k, v := range m {
        fn(k, v)
    }
    return m
}
```

**Test / 테스트**:
```go
func TestTap(t *testing.T) {
    t.Run("Executes function for each entry", func(t *testing.T) {
        m := map[string]int{"a": 1, "b": 2, "c": 3}
        count := 0

        result := Tap(m, func(k string, v int) {
            count++
        })

        if count != 3 {
            t.Errorf("Expected 3 calls, got %d", count)
        }

        // Verify original map unchanged
        if !reflect.DeepEqual(result, m) {
            t.Error("Tap should return original map")
        }
    })
}
```

---

## Testing Guide / 테스트 가이드

### Test Structure / 테스트 구조

Each `*_test.go` file contains:

각 `*_test.go` 파일은 다음을 포함합니다:

1. **Unit tests for each function** / 각 함수에 대한 단위 테스트
2. **Table-driven tests** / 테이블 주도 테스트
3. **Edge case tests** / 엣지 케이스 테스트
4. **Benchmark tests** / 벤치마크 테스트

**Example Test File Structure / 예제 테스트 파일 구조**:
```go
package maputil

import (
    "testing"
    "reflect"
)

// TestFunctionName tests the FunctionName function
// TestFunctionName은 FunctionName 함수를 테스트합니다
func TestFunctionName(t *testing.T) {
    // Table-driven test / 테이블 주도 테스트
    tests := []struct {
        name     string
        input    map[string]int
        expected map[string]int
    }{
        {
            name:     "Normal case",
            input:    map[string]int{"a": 1, "b": 2},
            expected: map[string]int{"a": 2, "b": 4},
        },
        {
            name:     "Empty map",
            input:    map[string]int{},
            expected: map[string]int{},
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FunctionName(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("Expected %v, got %v", tt.expected, result)
            }
        })
    }
}

// BenchmarkFunctionName benchmarks the FunctionName function
// BenchmarkFunctionName은 FunctionName 함수를 벤치마크합니다
func BenchmarkFunctionName(b *testing.B) {
    m := make(map[int]int, 1000)
    for i := 0; i < 1000; i++ {
        m[i] = i
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        FunctionName(m)
    }
}
```

### Running Tests / 테스트 실행

```bash
# Run all tests / 모든 테스트 실행
go test ./maputil -v

# Run specific test / 특정 테스트 실행
go test ./maputil -v -run TestFilter

# Run tests with coverage / 커버리지와 함께 테스트 실행
go test ./maputil -cover

# Generate coverage report / 커버리지 보고서 생성
go test ./maputil -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks / 벤치마크 실행
go test ./maputil -bench=.

# Run specific benchmark / 특정 벤치마크 실행
go test ./maputil -bench=BenchmarkMap

# Verbose benchmarks with memory stats / 메모리 통계와 함께 상세 벤치마크
go test ./maputil -bench=. -benchmem
```

### Writing Good Tests / 좋은 테스트 작성하기

**Test Checklist / 테스트 체크리스트**:

- [ ] Normal/happy path case / 정상/행복 경로 케이스
- [ ] Empty map case / 빈 맵 케이스
- [ ] Nil map case / nil 맵 케이스
- [ ] Single entry case / 단일 항목 케이스
- [ ] Large map case (performance) / 큰 맵 케이스 (성능)
- [ ] Duplicate keys/values (if relevant) / 중복 키/값 (관련 있는 경우)
- [ ] Invalid parameters (negative, zero, etc.) / 잘못된 매개변수 (음수, 0 등)
- [ ] Benchmark test / 벤치마크 테스트
- [ ] Immutability check (original unchanged) / 불변성 검사 (원본 변경 없음)
- [ ] Type safety with different types / 다양한 타입으로 타입 안전성

**Example: Comprehensive Test / 예제: 종합 테스트**:
```go
func TestMap(t *testing.T) {
    // Test 1: Transform values
    t.Run("Transform values", func(t *testing.T) {
        m := map[string]int{"a": 1, "b": 2, "c": 3}
        result := Map(m, func(k string, v int) int {
            return v * 2
        })
        expected := map[string]int{"a": 2, "b": 4, "c": 6}

        if !reflect.DeepEqual(result, expected) {
            t.Errorf("Expected %v, got %v", expected, result)
        }
    })

    // Test 2: Empty map
    t.Run("Empty map", func(t *testing.T) {
        m := map[string]int{}
        result := Map(m, func(k string, v int) int {
            return v * 2
        })

        if len(result) != 0 {
            t.Errorf("Expected empty map, got %v", result)
        }
    })

    // Test 3: Immutability
    t.Run("Original map unchanged", func(t *testing.T) {
        m := map[string]int{"a": 1, "b": 2}
        original := Clone(m)

        Map(m, func(k string, v int) int {
            return v * 100
        })

        if !reflect.DeepEqual(m, original) {
            t.Error("Original map was modified")
        }
    })

    // Test 4: Type transformation
    t.Run("Transform to different type", func(t *testing.T) {
        m := map[string]int{"a": 1, "b": 2}
        result := Map(m, func(k string, v int) string {
            return fmt.Sprintf("%s=%d", k, v)
        })

        if len(result) != 2 {
            t.Errorf("Expected 2 entries, got %d", len(result))
        }
    })
}
```

### Test Coverage Goals / 테스트 커버리지 목표

**Target / 목표**: 100% code coverage / 100% 코드 커버리지

**Current Status / 현재 상태**: ~100% (all functions tested) / ~100% (모든 함수 테스트됨)

**How to Achieve / 달성 방법**:
1. Test every function / 모든 함수 테스트
2. Test every branch (if/else) / 모든 분기 테스트 (if/else)
3. Test every loop iteration type / 모든 루프 반복 타입 테스트
4. Test edge cases / 엣지 케이스 테스트
5. Test error conditions / 에러 조건 테스트

---

## Performance / 성능

### Time Complexity Summary / 시간 복잡도 요약

| Category / 카테고리 | Functions / 함수 | Time Complexity / 시간 복잡도 |
|---------------------|-----------------|------------------------------|
| **Basic** | Get, Has, IsEmpty, Len | O(1) |
| **Basic** | Clone, Set, Delete, Equal | O(n) |
| **Transformation** | Map, MapKeys, MapValues, MapEntries, Invert, Compact | O(n) |
| **Transformation** | Flatten, Unflatten | O(n*m) nested maps |
| **Transformation** | Chunk, Partition | O(n) |
| **Aggregation** | Reduce, Sum, Average, Min, Max, MinBy, MaxBy | O(n) |
| **Aggregation** | GroupBy, CountBy | O(n) |
| **Merge** | Merge, Union, Difference, SymmetricDifference, Assign | O(n+m) |
| **Merge** | MergeWith, DeepMerge | O(n*m) |
| **Merge** | Intersection | O(n*m) multi-map |
| **Filter** | Filter, FilterKeys, FilterValues, OmitBy, PickBy | O(n) |
| **Filter** | Omit | O(n) |
| **Filter** | Pick | O(k) k=num keys |
| **Conversion** | Keys, Values, Entries, FromEntries, FromSlice, ToSlice | O(n) |
| **Conversion** | ToJSON, FromJSON | O(n) |
| **Predicate** | Every, Some, None | O(n) worst, short-circuit |
| **Predicate** | HasKey | O(1) |
| **Predicate** | HasValue, HasEntry, IsSubset, IsSuperset | O(n) |
| **Keys** | KeysBy, FindKey, PrefixKeys, SuffixKeys, TransformKeys | O(n) |
| **Keys** | KeysSorted | O(n log n) |
| **Keys** | RenameKey, RenameKeys, SwapKeys | O(n) |
| **Values** | ValuesBy, FindValue, ReplaceValue, UpdateValues | O(n) |
| **Values** | ValuesSorted | O(n log n) |
| **Values** | UniqueValues, MinValue, MaxValue | O(n) |
| **Comparison** | Diff, DiffKeys, Compare | O(n+m) |
| **Comparison** | EqualMaps, EqualFunc | O(n) |
| **Comparison** | CommonKeys, AllKeys | O(n*m) multi-map |

**Key Notation / 주요 표기법**:
- `n` = size of input map(s) / 입력 맵의 크기
- `m` = number of maps (for multi-map operations) / 맵 개수 (다중 맵 작업용)
- `k` = number of keys (for Pick operation) / 키 개수 (Pick 작업용)

### Space Complexity / 공간 복잡도

**Most operations**: O(n) space for result map / 대부분의 작업: 결과 맵에 대해 O(n) 공간

**Exceptions / 예외**:
- `Reduce`, `Sum`, `Average`, `Min`, `Max`: O(1) / O(1)
- `Every`, `Some`, `None`, `Has*`: O(1) / O(1)
- `Chunk`: O(n) for all chunks combined / 모든 청크를 합쳐서 O(n)
- `GroupBy`: O(n) for all groups combined / 모든 그룹을 합쳐서 O(n)

### Performance Tips / 성능 팁

**1. Pre-allocate Maps / 맵 미리 할당**

```go
// Good: Pre-allocate with expected size
// 좋음: 예상 크기로 미리 할당
result := make(map[K]V, len(input))

// Bad: Let map grow dynamically
// 나쁨: 맵이 동적으로 증가하도록 둠
result := make(map[K]V)
```

**2. Use Appropriate Functions / 적절한 함수 사용**

```go
// Pick is O(k) where k = number of keys
// Pick은 O(k) 여기서 k = 키 개수
picked := Pick(largeMap, "key1", "key2") // Fast for few keys

// Filter is O(n) where n = map size
// Filter는 O(n) 여기서 n = 맵 크기
filtered := Filter(largeMap, func(k, v) bool { ... }) // Scans entire map
```

**3. Avoid Unnecessary Cloning / 불필요한 복제 방지**

```go
// If you know the result will be different, don't clone first
// 결과가 다를 것을 알고 있다면, 먼저 복제하지 마세요
// Good: Direct creation
result := make(map[K]V, len(m))
for k, v := range m {
    result[k] = transform(v)
}

// Bad: Clone then modify (wasteful)
// result := Clone(m)
// for k := range result {
//     result[k] = transform(result[k])
// }
```

**4. Use Hash Sets for Lookups / 조회에 해시 집합 사용**

```go
// Good: O(1) lookup with hash set
// 좋음: 해시 집합으로 O(1) 조회
toDelete := make(map[K]struct{})
for _, key := range keysToDelete {
    toDelete[key] = struct{}{}
}
for k, v := range m {
    if _, shouldDelete := toDelete[k]; !shouldDelete {
        result[k] = v
    }
}

// Bad: O(n*m) with slice Contains
// 나쁨: 슬라이스 Contains로 O(n*m)
// for k, v := range m {
//     if !contains(keysToDelete, k) {
//         result[k] = v
//     }
// }
```

**5. Short-Circuit When Possible / 가능하면 단락 평가**

```go
// Every/Some/None return early
// Every/Some/None은 조기 반환
func Every[K, V](m map[K]V, fn func(K, V) bool) bool {
    for k, v := range m {
        if !fn(k, v) {
            return false // Exit immediately
        }
    }
    return true
}
```

### Benchmark Results / 벤치마크 결과

**Typical performance on modern hardware / 최신 하드웨어에서의 일반적인 성능**:

```
BenchmarkClone/size=100-8         	  500000	      2500 ns/op
BenchmarkClone/size=1000-8        	   50000	     25000 ns/op
BenchmarkClone/size=10000-8       	    5000	    250000 ns/op

BenchmarkFilter/size=100-8        	  300000	      3500 ns/op
BenchmarkFilter/size=1000-8       	   30000	     35000 ns/op
BenchmarkFilter/size=10000-8      	    3000	    350000 ns/op

BenchmarkMap/size=100-8           	  400000	      3000 ns/op
BenchmarkMap/size=1000-8          	   40000	     30000 ns/op
BenchmarkMap/size=10000-8         	    4000	    300000 ns/op

BenchmarkMerge/2maps-8            	  200000	      6000 ns/op
BenchmarkMerge/5maps-8            	   80000	     15000 ns/op
BenchmarkMerge/10maps-8           	   40000	     30000 ns/op
```

**Note / 참고**: Actual performance depends on hardware, Go version, and data characteristics.

실제 성능은 하드웨어, Go 버전, 데이터 특성에 따라 다릅니다.

---

## Contributing Guidelines / 기여 가이드라인

### Pull Request Process / 풀 리퀘스트 프로세스

**Step 1: Fork and Clone / 포크 및 복제**

```bash
# Fork the repository on GitHub
# GitHub에서 저장소 포크

# Clone your fork
# 포크한 저장소 복제
git clone https://github.com/yourusername/go-utils.git
cd go-utils
```

**Step 2: Create Feature Branch / 기능 브랜치 생성**

```bash
git checkout -b feature/maputil-new-function
```

**Step 3: Make Changes / 변경 사항 만들기**

1. Add your function to the appropriate file / 적절한 파일에 함수 추가
2. Write comprehensive tests / 종합 테스트 작성
3. Update documentation (README, USER_MANUAL) / 문서 업데이트 (README, USER_MANUAL)
4. Ensure code style compliance / 코드 스타일 준수 확인

**Step 4: Run Tests / 테스트 실행**

```bash
# Run all tests
go test ./maputil -v

# Check coverage
go test ./maputil -cover

# Run benchmarks
go test ./maputil -bench=.

# Format code
go fmt ./maputil/...

# Lint code (if golangci-lint installed)
golangci-lint run ./maputil/...
```

**Step 5: Commit and Push / 커밋 및 푸시**

```bash
git add .
git commit -m "Feat: Add NewFunction to maputil package"
git push origin feature/maputil-new-function
```

**Step 6: Create Pull Request / 풀 리퀘스트 생성**

- Go to GitHub and create a PR from your branch / GitHub로 가서 브랜치에서 PR 생성
- Fill in the PR template / PR 템플릿 작성
- Wait for review / 리뷰 대기

### PR Checklist / PR 체크리스트

Before submitting a PR, ensure:

PR을 제출하기 전에 다음을 확인하세요:

- [ ] All tests pass (`go test ./maputil -v`) / 모든 테스트 통과
- [ ] Code coverage is 100% / 코드 커버리지 100%
- [ ] Benchmarks added for new functions / 새 함수에 대한 벤치마크 추가
- [ ] Documentation updated (README, USER_MANUAL, DEVELOPER_GUIDE) / 문서 업데이트
- [ ] Code follows style guide (see below) / 코드가 스타일 가이드를 따름 (아래 참조)
- [ ] Function has bilingual comments (English/Korean) / 함수에 이중 언어 주석 (영문/한글)
- [ ] Examples added to README / README에 예제 추가
- [ ] No breaking changes (or clearly documented) / 중단 변경 없음 (또는 명확하게 문서화됨)
- [ ] Immutability preserved (except for Assign) / 불변성 유지 (Assign 제외)
- [ ] Generic constraints appropriate / 제네릭 제약 적절

### Code Review Guidelines / 코드 리뷰 가이드라인

**Reviewers should check / 리뷰어가 확인해야 할 사항**:

1. **Correctness / 정확성**: Does the code do what it claims? / 코드가 주장하는 대로 작동하는가?
2. **Tests / 테스트**: Are all edge cases tested? / 모든 엣지 케이스가 테스트되었는가?
3. **Performance / 성능**: Is the algorithm efficient? / 알고리즘이 효율적인가?
4. **Documentation / 문서**: Are comments clear and bilingual? / 주석이 명확하고 이중 언어인가?
5. **Style / 스타일**: Does it follow the style guide? / 스타일 가이드를 따르는가?
6. **Immutability / 불변성**: Does it preserve immutability? / 불변성을 유지하는가?

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

**1. Function Names / 함수 이름**

- Use PascalCase for exported functions / 내보낸 함수에 PascalCase 사용
- Use descriptive, action-oriented names / 설명적이고 행동 지향적인 이름 사용
- Avoid abbreviations / 약어 사용 안 함

```go
// Good
func Map[K, V, R](...) ...
func Filter[K, V](...) ...
func GroupBy[K, V, G](...) ...

// Bad
func M[K, V, R](...) ...        // Too short
func Flt[K, V](...) ...         // Abbreviation
func MapTheValues[K, V, R](...) // Too verbose
```

**2. Parameter Names / 매개변수 이름**

- Use short, clear names / 짧고 명확한 이름 사용
- `m` for map parameters / 맵 매개변수에 `m`
- `fn` for function parameters / 함수 매개변수에 `fn`
- `k`, `v` for key-value pairs / 키-값 쌍에 `k`, `v`

```go
func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
    for k, v := range m {
        if fn(k, v) {
            // ...
        }
    }
}
```

**3. Variable Names / 변수 이름**

- Use descriptive names for important variables / 중요한 변수에 설명적인 이름 사용
- Use single letters for loop indices / 루프 인덱스에 단일 문자 사용

```go
// Good
result := make(map[K]V, len(m))
seen := make(map[K]bool)
for i := 0; i < len(keys); i++ {
    // ...
}

// Bad
r := make(map[K]V, len(m))      // Too short
resultMap := make(map[K]V)      // Redundant "Map"
index := 0                       // Too verbose for loop
```

### Documentation Style / 문서 스타일

**1. Function Comments / 함수 주석**

Every exported function must have:

모든 내보낸 함수는 다음을 가져야 합니다:

```go
// FunctionName does something useful
// FunctionName은 유용한 작업을 수행합니다
//
// Longer description with details about behavior, edge cases, etc.
// 동작, 엣지 케이스 등에 대한 세부 사항이 포함된 더 긴 설명.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example / 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	result := maputil.FunctionName(m, ...)
//	// result = ...
func FunctionName[K comparable, V any](m map[K]V, ...) ReturnType {
    // Implementation
}
```

**2. Inline Comments / 인라인 주석**

- Use bilingual comments for important logic / 중요한 로직에 이중 언어 주석 사용
- Explain WHY, not WHAT / 무엇이 아닌 왜를 설명

```go
// Good
// Create hash set for O(1) lookup
// O(1) 조회를 위한 해시 집합 생성
seen := make(map[K]bool)

// Bad
// Create a map
seen := make(map[K]bool)
```

### Code Organization / 코드 구성

**1. Import Order / 임포트 순서**

```go
import (
    // Standard library first
    // 표준 라이브러리 먼저
    "encoding/json"
    "fmt"
    "sort"

    // Third-party packages second
    // 서드파티 패키지 두 번째
    // (none for maputil)

    // Local packages last
    // 로컬 패키지 마지막
    "github.com/arkd0ng/go-utils/logging"
)
```

**2. Function Order in Files / 파일의 함수 순서**

1. Package documentation / 패키지 문서
2. Type definitions / 타입 정의
3. Public functions (alphabetical) / 공개 함수 (알파벳순)
4. Private helper functions / 비공개 헬퍼 함수

**3. Error Handling / 에러 처리**

- Handle edge cases at the start of functions / 함수 시작 부분에서 엣지 케이스 처리
- Return appropriate zero/empty values / 적절한 zero/빈 값 반환
- No panics (except for truly exceptional cases) / 패닉 없음 (진정으로 예외적인 경우 제외)

```go
func ExampleFunction[K comparable, V any](m map[K]V, ...) ReturnType {
    // Handle edge cases first
    if m == nil {
        return nil // or appropriate zero value
    }
    if len(m) == 0 {
        return EmptyResult
    }

    // Main logic
    result := make(map[K]V, len(m))
    // ...
    return result
}
```

### Best Practices / 모범 사례

**1. Immutability / 불변성**

Always return new maps (except `Assign`):

항상 새 맵 반환 (`Assign` 제외):

```go
// Good: Returns new map
func Set[K, V](m map[K]V, key K, value V) map[K]V {
    result := make(map[K]V, len(m)+1)
    for k, v := range m {
        result[k] = v
    }
    result[key] = value
    return result
}

// Bad: Modifies original
// func Set[K, V](m map[K]V, key K, value V) map[K]V {
//     m[key] = value
//     return m
// }
```

**2. Pre-allocation / 미리 할당**

Pre-allocate maps when size is known:

크기를 알고 있을 때 맵 미리 할당:

```go
// Good: Pre-allocate with capacity
result := make(map[K]V, len(m))

// Acceptable: Unknown size
result := make(map[K]V)
```

**3. Generic Constraints / 제네릭 제약**

Use the most specific constraint:

가장 구체적인 제약 사용:

```go
// Good: Specific constraints
func Sum[K comparable, V Number](m map[K]V) V { ... }
func Min[K comparable, V Ordered](m map[K]V) (K, V, bool) { ... }

// Bad: Too generic
// func Sum[K comparable, V any](m map[K]V) V { ... } // Won't compile
```

**4. Test Coverage / 테스트 커버리지**

Aim for 100% coverage with meaningful tests:

의미 있는 테스트로 100% 커버리지 목표:

- Normal cases / 정상 케이스
- Edge cases (empty, nil, single) / 엣지 케이스 (빈, nil, 단일)
- Error conditions / 에러 조건
- Immutability checks / 불변성 검사
- Benchmarks / 벤치마크

---

## Conclusion / 결론

The Maputil package provides 81 type-safe, immutable, and efficient functions for map operations in Go. By following this developer guide, you can:

Maputil 패키지는 Go에서 맵 작업을 위한 81개의 타입 안전하고 불변적이며 효율적인 함수를 제공합니다. 이 개발자 가이드를 따르면 다음을 할 수 있습니다:

1. **Understand the architecture** / 아키텍처 이해
2. **Navigate the codebase** / 코드베이스 탐색
3. **Add new features** / 새 기능 추가
4. **Write comprehensive tests** / 종합 테스트 작성
5. **Optimize performance** / 성능 최적화
6. **Contribute effectively** / 효과적으로 기여

**Key Takeaways / 주요 요점**:
- All operations are immutable (except Assign) / 모든 작업은 불변 (Assign 제외)
- Type safety through generics / 제네릭을 통한 타입 안전성
- Comprehensive test coverage (100%) / 포괄적인 테스트 커버리지 (100%)
- Bilingual documentation / 이중 언어 문서
- Zero external dependencies / 외부 의존성 제로

For more information, see:

자세한 내용은 다음을 참조하세요:

- **USER_MANUAL.md**: User-facing documentation with examples / 예제가 포함된 사용자 대상 문서
- **README.md**: Quick start and overview / 빠른 시작 및 개요
- **Source code**: `maputil/*.go` files / `maputil/*.go` 파일

Happy coding! / 즐거운 코딩!

---

**Document Version / 문서 버전**: v1.8.005
**Last Updated / 마지막 업데이트**: 2025-10-15
**Maintained By / 유지 관리자**: go-utils team
