# maputil Package Design Plan / maputil 패키지 설계 계획

## Overview / 개요

The `maputil` package provides extreme simplicity map utilities - reduce 20+ lines of repetitive map manipulation code to just 1-2 lines with **80+ type-safe functions**.

`maputil` 패키지는 극도로 간단한 맵 유틸리티를 제공합니다 - 20줄 이상의 반복적인 맵 조작 코드를 단 1-2줄로 줄이며, **80개 이상의 타입 안전 함수**를 제공합니다.

## Design Philosophy / 설계 철학

**"20 lines → 1 line"**

### Before (Standard Go) / 이전 (표준 Go)
```go
// Filtering a map / 맵 필터링
data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := make(map[string]int)
for k, v := range data {
    if v > 2 {
        result[k] = v
    }
}
// 7+ lines of boilerplate

// Merging maps / 맵 병합
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"b": 3, "c": 4}
merged := make(map[string]int)
for k, v := range map1 {
    merged[k] = v
}
for k, v := range map2 {
    merged[k] = v
}
// 10+ lines
```

### After (maputil) / 이후 (maputil)
```go
// Filtering / 필터링
result := maputil.Filter(data, func(k string, v int) bool { return v > 2 })

// Merging / 병합
merged := maputil.Merge(map1, map2)
```

## Core Features / 핵심 기능

1. **Type Safety / 타입 안전성**
   - Go 1.18+ generics for compile-time type checking
   - Go 1.18+ 제네릭으로 컴파일 타임 타입 체크
   - Type constraints: comparable (keys), any (values)
   - 타입 제약조건: comparable (키), any (값)

2. **Functional Programming / 함수형 프로그래밍**
   - Inspired by JavaScript, Python, Lodash map methods
   - JavaScript, Python, Lodash 맵 메서드에서 영감
   - Higher-order functions: Map, Filter, Reduce
   - 고차 함수: Map, Filter, Reduce

3. **Immutable Operations / 불변 작업**
   - All functions return new maps (original unchanged)
   - 모든 함수는 새 맵을 반환 (원본 변경 없음)
   - Prevents side effects
   - 부작용 방지

4. **Zero Dependencies / 제로 의존성**
   - Standard library only
   - 표준 라이브러리만 사용
   - No external packages
   - 외부 패키지 없음

5. **High Performance / 고성능**
   - Efficient algorithms (mostly O(n))
   - 효율적인 알고리즘 (대부분 O(n))
   - Optimized for common use cases
   - 일반 사용 사례에 최적화

## Package Structure / 패키지 구조

```
maputil/
├── maputil.go          # Package documentation and version / 패키지 문서 및 버전
├── basic.go            # Basic operations (11 functions)
├── transform.go        # Transformation functions (10 functions)
├── aggregate.go        # Aggregation functions (9 functions)
├── merge.go            # Merge and combine operations (8 functions)
├── filter.go           # Filter and partition operations (7 functions)
├── convert.go          # Conversion functions (8 functions)
├── predicate.go        # Predicate checks (7 functions)
├── keys.go             # Key operations (8 functions)
├── values.go           # Value operations (7 functions)
├── comparison.go       # Comparison functions (6 functions)
├── maputil_test.go     # Comprehensive tests (100% coverage)
└── README.md           # Package documentation
```

## Function Categories / 함수 카테고리

### 1. Basic Operations (11 functions)
**basic.go**

- `Get[K comparable, V any](m map[K]V, key K) (V, bool)` - Get value with existence check
- `GetOr[K comparable, V any](m map[K]V, key K, defaultValue V) V` - Get with default
- `Set[K comparable, V any](m map[K]V, key K, value V) map[K]V` - Set key-value (returns new map)
- `Delete[K comparable, V any](m map[K]V, keys ...K) map[K]V` - Delete keys (returns new map)
- `Has[K comparable, V any](m map[K]V, key K) bool` - Check key existence
- `IsEmpty[K comparable, V any](m map[K]V) bool` - Check if empty
- `IsNotEmpty[K comparable, V any](m map[K]V) bool` - Check if not empty
- `Len[K comparable, V any](m map[K]V) int` - Get length
- `Clear[K comparable, V any](m map[K]V) map[K]V` - Clear all entries (returns empty map)
- `Clone[K comparable, V any](m map[K]V) map[K]V` - Deep clone
- `Equal[K comparable, V comparable](m1, m2 map[K]V) bool` - Check equality

### 2. Transformation (10 functions)
**transform.go**

- `Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R` - Transform values
- `MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V` - Transform keys
- `MapValues[K comparable, V any, R any](m map[K]V, fn func(V) R) map[K]R` - Transform values (value-only)
- `MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, fn func(K1, V1) (K2, V2)) map[K2]V2` - Transform both keys and values
- `Invert[K comparable, V comparable](m map[K]V) map[V]K` - Swap keys and values
- `Flatten[K comparable, V any](m map[K]map[string]V, delimiter string) map[string]V` - Flatten nested map
- `Unflatten[V any](m map[string]V, delimiter string) map[string]interface{}` - Unflatten to nested map
- `Chunk[K comparable, V any](m map[K]V, size int) []map[K]V` - Split into chunks
- `Partition[K comparable, V any](m map[K]V, fn func(K, V) bool) (map[K]V, map[K]V)` - Partition by predicate
- `Compact[K comparable, V any](m map[K]V) map[K]V` - Remove zero values

### 3. Aggregation (9 functions)
**aggregate.go**

- `Reduce[K comparable, V any, R any](m map[K]V, initial R, fn func(R, K, V) R) R` - Reduce to single value
- `Sum[K comparable, V Number](m map[K]V) V` - Sum all values
- `Min[K comparable, V Ordered](m map[K]V) (K, V, bool)` - Find minimum value
- `Max[K comparable, V Ordered](m map[K]V) (K, V, bool)` - Find maximum value
- `MinBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool)` - Find min by function
- `MaxBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool)` - Find max by function
- `Average[K comparable, V Number](m map[K]V) float64` - Calculate average
- `GroupBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G][]V` - Group slice by key
- `CountBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G]int` - Count by key

### 4. Merge Operations (8 functions)
**merge.go**

- `Merge[K comparable, V any](maps ...map[K]V) map[K]V` - Merge multiple maps (last wins)
- `MergeWith[K comparable, V any](fn func(V, V) V, maps ...map[K]V) map[K]V` - Merge with custom resolver
- `DeepMerge(maps ...map[string]interface{}) map[string]interface{}` - Deep merge nested maps
- `Union[K comparable, V any](maps ...map[K]V) map[K]V` - Union (same as Merge)
- `Intersection[K comparable, V any](maps ...map[K]V) map[K]V` - Intersection (common keys)
- `Difference[K comparable, V any](m1, m2 map[K]V) map[K]V` - Difference (keys in m1 not in m2)
- `SymmetricDifference[K comparable, V any](m1, m2 map[K]V) map[K]V` - Symmetric difference
- `Assign[K comparable, V any](target map[K]V, sources ...map[K]V) map[K]V` - Assign to target (mutating)

### 5. Filter Operations (7 functions)
**filter.go**

- `Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V` - Filter by predicate
- `FilterKeys[K comparable, V any](m map[K]V, fn func(K) bool) map[K]V` - Filter by keys
- `FilterValues[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V` - Filter by values
- `Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V` - Omit specified keys
- `Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V` - Pick specified keys
- `OmitBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V` - Omit by predicate
- `PickBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V` - Pick by predicate

### 6. Conversion (8 functions)
**convert.go**

- `Keys[K comparable, V any](m map[K]V) []K` - Get all keys
- `Values[K comparable, V any](m map[K]V) []V` - Get all values
- `Entries[K comparable, V any](m map[K]V) []Entry[K, V]` - Get key-value pairs
- `FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V` - Create from entries
- `FromSlice[K comparable, V any](slice []V, fn func(V) K) map[K]V` - Create from slice
- `FromSliceBy[K comparable, V any, R any](slice []V, keyFn func(V) K, valueFn func(V) R) map[K]R` - Create from slice with transform
- `ToSlice[K comparable, V any, R any](m map[K]V, fn func(K, V) R) []R` - Convert to slice
- `ToJSON[K comparable, V any](m map[K]V) (string, error)` - Convert to JSON string

### 7. Predicate Checks (7 functions)
**predicate.go**

- `Every[K comparable, V any](m map[K]V, fn func(K, V) bool) bool` - Check if all match
- `Some[K comparable, V any](m map[K]V, fn func(K, V) bool) bool` - Check if any match
- `None[K comparable, V any](m map[K]V, fn func(K, V) bool) bool` - Check if none match
- `HasKey[K comparable, V any](m map[K]V, key K) bool` - Check key existence (alias)
- `HasValue[K comparable, V comparable](m map[K]V, value V) bool` - Check value existence
- `HasEntry[K comparable, V comparable](m map[K]V, key K, value V) bool` - Check key-value pair
- `IsSubset[K comparable, V comparable](subset, superset map[K]V) bool` - Check if subset

### 8. Key Operations (8 functions)
**keys.go**

- `KeysSlice[K comparable, V any](m map[K]V) []K` - Get keys as slice (alias)
- `KeysSorted[K Ordered, V any](m map[K]V) []K` - Get sorted keys
- `KeysBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []K` - Get keys by predicate
- `RenameKey[K comparable, V any](m map[K]V, oldKey, newKey K) map[K]V` - Rename key
- `RenameKeys[K comparable, V any](m map[K]V, mapping map[K]K) map[K]V` - Rename multiple keys
- `SwapKeys[K comparable, V any](m map[K]V, key1, key2 K) map[K]V` - Swap two keys
- `FindKey[K comparable, V any](m map[K]V, fn func(K, V) bool) (K, bool)` - Find key by predicate
- `FindKeys[K comparable, V any](m map[K]V, fn func(K, V) bool) []K` - Find all keys by predicate

### 9. Value Operations (7 functions)
**values.go**

- `ValuesSlice[K comparable, V any](m map[K]V) []V` - Get values as slice (alias)
- `ValuesSorted[K comparable, V Ordered](m map[K]V) []V` - Get sorted values
- `ValuesBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []V` - Get values by predicate
- `UniqueValues[K comparable, V comparable](m map[K]V) []V` - Get unique values
- `FindValue[K comparable, V any](m map[K]V, fn func(K, V) bool) (V, bool)` - Find value by predicate
- `ReplaceValue[K comparable, V comparable](m map[K]V, oldValue, newValue V) map[K]V` - Replace all occurrences
- `UpdateValues[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V` - Update all values

### 10. Comparison (6 functions)
**comparison.go**

- `EqualFunc[K comparable, V any](m1, m2 map[K]V, eq func(V, V) bool) bool` - Check equality with custom comparator
- `Diff[K comparable, V comparable](m1, m2 map[K]V) map[K]V` - Get differing entries
- `DiffKeys[K comparable, V any](m1, m2 map[K]V) []K` - Get differing keys
- `CommonKeys[K comparable, V any](maps ...map[K]V) []K` - Get common keys
- `AllKeys[K comparable, V any](maps ...map[K]V) []K` - Get all unique keys
- `Compare[K comparable, V comparable](m1, m2 map[K]V) (added, removed, modified map[K]V)` - Detailed comparison

## Type Definitions / 타입 정의

```go
// Entry represents a key-value pair / Entry는 키-값 쌍을 나타냅니다
type Entry[K comparable, V any] struct {
    Key   K
    Value V
}

// Number constraint for numeric operations / 숫자 작업을 위한 Number 제약조건
type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
        ~float32 | ~float64
}

// Ordered constraint for comparison operations / 비교 작업을 위한 Ordered 제약조건
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
        ~float32 | ~float64 | ~string
}
```

## Design Principles / 설계 원칙

1. **Consistency with sliceutil / sliceutil과의 일관성**
   - Similar naming conventions
   - 유사한 명명 규칙
   - Parallel function signatures
   - 병렬 함수 시그니처

2. **Functional Programming / 함수형 프로그래밍**
   - Higher-order functions
   - 고차 함수
   - Pure functions (no side effects)
   - 순수 함수 (부작용 없음)
   - Composability
   - 조합 가능성

3. **Performance / 성능**
   - Avoid unnecessary allocations
   - 불필요한 할당 방지
   - Efficient algorithms
   - 효율적인 알고리즘
   - Benchmark all functions
   - 모든 함수 벤치마크

4. **Type Safety / 타입 안전성**
   - Leverage generics
   - 제네릭 활용
   - Compile-time guarantees
   - 컴파일 타임 보장
   - Clear constraints
   - 명확한 제약조건

5. **Immutability / 불변성**
   - Never modify input maps
   - 입력 맵을 절대 수정하지 않음
   - Always return new maps
   - 항상 새 맵 반환
   - Prevent side effects
   - 부작용 방지

## Testing Strategy / 테스트 전략

1. **Unit Tests / 단위 테스트**
   - Test each function independently
   - 각 함수를 독립적으로 테스트
   - Edge cases (empty maps, nil values)
   - 엣지 케이스 (빈 맵, nil 값)
   - Type variations
   - 타입 변형

2. **Integration Tests / 통합 테스트**
   - Function composition
   - 함수 조합
   - Complex pipelines
   - 복잡한 파이프라인

3. **Benchmark Tests / 벤치마크 테스트**
   - Performance measurement
   - 성능 측정
   - Memory allocation tracking
   - 메모리 할당 추적

4. **Coverage / 커버리지**
   - Target: 100% code coverage
   - 목표: 100% 코드 커버리지
   - All branches tested
   - 모든 분기 테스트

## Documentation / 문서화

1. **Code Comments / 코드 주석**
   - Bilingual (English/Korean)
   - 이중 언어 (영문/한글)
   - Clear examples
   - 명확한 예제
   - Time complexity notes
   - 시간 복잡도 주석

2. **Package README / 패키지 README**
   - Quick start guide
   - 빠른 시작 가이드
   - Common use cases
   - 일반적인 사용 사례
   - Function categorization
   - 함수 분류

3. **Comprehensive Manuals / 포괄적인 매뉴얼**
   - USER_MANUAL.md (~3000+ lines)
   - DEVELOPER_GUIDE.md (~2000+ lines)
   - PERFORMANCE_BENCHMARKS.md

## Implementation Roadmap / 구현 로드맵

### Phase 1: Core Functions
1. basic.go - Basic operations
2. transform.go - Transformation functions
3. aggregate.go - Aggregation functions

### Phase 2: Advanced Functions
4. merge.go - Merge operations
5. filter.go - Filter operations
6. convert.go - Conversion functions

### Phase 3: Specialized Functions
7. predicate.go - Predicate checks
8. keys.go - Key operations
9. values.go - Value operations
10. comparison.go - Comparison functions

### Phase 4: Testing & Documentation
- Comprehensive tests (100% coverage)
- Benchmark tests
- Examples
- Documentation

## Success Metrics / 성공 지표

1. **Code Reduction / 코드 감소**
   - 20+ lines → 1-2 lines
   - 20줄 이상 → 1-2줄

2. **Test Coverage / 테스트 커버리지**
   - 100% code coverage
   - 100% 코드 커버리지

3. **Performance / 성능**
   - O(n) for most operations
   - 대부분의 작업에서 O(n)
   - Minimal allocations
   - 최소 할당

4. **Usability / 사용성**
   - Intuitive API
   - 직관적인 API
   - Clear error messages
   - 명확한 에러 메시지
   - Comprehensive examples
   - 포괄적인 예제

## Estimated Function Count / 예상 함수 수

- **Total: 80+ functions** across 10 categories
- **총 80개 이상 함수** 10개 카테고리에 걸쳐

## Version / 버전

- Initial release: **v1.8.001**
- 초기 릴리스: **v1.8.001**

## Conclusion / 결론

The maputil package will provide a comprehensive, type-safe, and performant solution for map manipulation in Go, following the successful pattern established by sliceutil.

maputil 패키지는 sliceutil에서 확립된 성공적인 패턴을 따라 Go에서 맵 조작을 위한 포괄적이고 타입 안전하며 고성능인 솔루션을 제공할 것입니다.
