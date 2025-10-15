# Maputil Package - User Manual / 사용자 매뉴얼

**Version / 버전**: v1.8.003
**Package / 패키지**: `github.com/arkd0ng/go-utils/maputil`
**Go Version Required / 필요한 Go 버전**: 1.18+

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
   - [What is Maputil? / Maputil이란?](#what-is-maputil--maputil이란)
   - [Design Philosophy / 설계 철학](#design-philosophy--설계-철학)
   - [Key Features / 주요 기능](#key-features--주요-기능)
   - [Use Cases / 사용 사례](#use-cases--사용-사례)
2. [Installation / 설치](#installation--설치)
   - [Prerequisites / 전제 조건](#prerequisites--전제-조건)
   - [Installation / 설치](#installation--설치-1)
   - [Importing / 임포트](#importing--임포트)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [API Reference / API 참조](#api-reference--api-참조)
   - [Category 1: Basic Operations](#category-1-basic-operations--기본-작업)
   - [Category 2: Transformation](#category-2-transformation--변환)
   - [Category 3: Aggregation](#category-3-aggregation--집계)
   - [Category 4: Merge Operations](#category-4-merge-operations--병합-작업)
   - [Category 5: Filter Operations](#category-5-filter-operations--필터-작업)
   - [Category 6: Conversion](#category-6-conversion--변환)
   - [Category 7: Predicate Checks](#category-7-predicate-checks--조건-검사)
   - [Category 8: Key Operations](#category-8-key-operations--키-작업)
   - [Category 9: Value Operations](#category-9-value-operations--값-작업)
   - [Category 10: Comparison](#category-10-comparison--비교)
5. [Usage Patterns / 사용 패턴](#usage-patterns--사용-패턴)
6. [Common Use Cases / 일반적인 사용 사례](#common-use-cases--일반적인-사용-사례)
7. [Best Practices / 모범 사례](#best-practices--모범-사례)
8. [Performance Considerations / 성능 고려사항](#performance-considerations--성능-고려사항)
9. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)
10. [FAQ / 자주 묻는 질문](#faq--자주-묻는-질문)
11. [Additional Resources / 추가 자료](#additional-resources--추가-자료)

---

## Introduction / 소개

### What is Maputil? / Maputil이란?

Maputil is an **extreme simplicity** map utility package for Go that provides 81 type-safe functions to simplify common map operations. It reduces 20+ lines of repetitive boilerplate code to just 1-2 lines while maintaining full type safety through Go 1.18+ generics.

Maputil은 Go를 위한 **극도로 간단한** 맵 유틸리티 패키지로, 일반적인 맵 작업을 단순화하는 81개의 타입 안전 함수를 제공합니다. Go 1.18+ 제네릭을 통해 완전한 타입 안전성을 유지하면서 20줄 이상의 반복적인 보일러플레이트 코드를 단 1-2줄로 줄입니다.

**Before Maputil / Maputil 사용 전:**
```go
// 20+ lines of code to filter a map
filtered := make(map[string]int)
for k, v := range data {
    if v > 100 {
        filtered[k] = v
    }
}
```

**After Maputil / Maputil 사용 후:**
```go
// 1 line with maputil
filtered := maputil.Filter(data, func(k string, v int) bool { return v > 100 })
```

### Design Philosophy / 설계 철학

The maputil package is built on these core principles:
maputil 패키지는 다음과 같은 핵심 원칙을 기반으로 합니다:

1. **Extreme Simplicity / 극도의 간결함**: Reduce 20+ lines to 1-2 lines
   20줄 이상의 코드를 1-2줄로 축소

2. **Type Safety / 타입 안전성**: Full compile-time type checking using Go 1.18+ generics
   Go 1.18+ 제네릭을 사용한 완전한 컴파일 타임 타입 체크

3. **Immutability / 불변성**: All operations return new maps without modifying originals (except `Assign`)
   모든 작업이 원본을 수정하지 않고 새 맵을 반환 (Assign 제외)

4. **Functional Programming / 함수형 프로그래밍**: Map, Filter, Reduce patterns inspired by JavaScript, Python, Ruby
   JavaScript, Python, Ruby에서 영감을 받은 Map, Filter, Reduce 패턴

5. **Zero Dependencies / 제로 의존성**: No external dependencies, only standard library
   외부 의존성 없음, 표준 라이브러리만 사용

6. **100% Test Coverage / 100% 테스트 커버리지**: Every function is thoroughly tested with benchmarks
   모든 함수가 벤치마크와 함께 철저히 테스트됨

### Key Features / 주요 기능

- ✅ **81 Functions / 81개 함수**: Covering 10 categories of map operations
  10개 카테고리의 맵 작업을 다루는 81개 함수

- ✅ **Type-Safe Generics / 타입 안전 제네릭**: Full type safety with Go 1.18+ generics
  Go 1.18+ 제네릭으로 완전한 타입 안전성

- ✅ **Immutable Operations / 불변 작업**: Original maps remain unchanged
  원본 맵이 변경되지 않음

- ✅ **Functional Style / 함수형 스타일**: Map, Filter, Reduce, GroupBy patterns
  Map, Filter, Reduce, GroupBy 패턴

- ✅ **Zero Dependencies / 제로 의존성**: Only standard library
  표준 라이브러리만 사용

- ✅ **High Performance / 고성능**: Optimized algorithms with O(n) complexity
  O(n) 복잡도의 최적화된 알고리즘

- ✅ **Comprehensive Tests / 포괄적인 테스트**: 80+ test functions, 60+ benchmarks
  80개 이상의 테스트 함수, 60개 이상의 벤치마크

- ✅ **Well Documented / 잘 문서화됨**: Every function has detailed examples
  모든 함수에 상세한 예제 포함

### Use Cases / 사용 사례

Maputil is perfect for:
Maputil은 다음과 같은 경우에 완벽합니다:

1. **Data Transformation / 데이터 변환**: Convert, filter, and transform map data
   맵 데이터 변환, 필터링 및 변형

2. **Configuration Management / 설정 관리**: Merge configs, pick/omit keys, validate settings
   설정 병합, 키 선택/제외, 설정 검증

3. **API Response Processing / API 응답 처리**: Filter sensitive fields, transform responses
   민감한 필드 필터링, 응답 변환

4. **Data Analysis / 데이터 분석**: GroupBy, aggregate, statistical operations
   GroupBy, 집계, 통계 작업

5. **State Management / 상태 관리**: Clone state, diff changes, merge updates
   상태 복제, 변경 비교, 업데이트 병합

6. **Testing / 테스트**: Compare maps, check subsets, validate data
   맵 비교, 서브셋 확인, 데이터 검증

---

## Installation / 설치

### Prerequisites / 전제 조건

- **Go 1.18 or higher** / **Go 1.18 이상** (required for generics / 제네릭 필요)
- No external dependencies / 외부 의존성 없음

### Installation / 설치

Install the package using `go get`:
`go get`을 사용하여 패키지를 설치합니다:

```bash
go get github.com/arkd0ng/go-utils/maputil
```

### Importing / 임포트

```go
import "github.com/arkd0ng/go-utils/maputil"
```

---

## Quick Start / 빠른 시작

Here are 10 quick examples to get you started:
시작하기 위한 10가지 빠른 예제입니다:

### 1. Filter Map / 맵 필터링

```go
data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

// Filter values > 2 / 값이 2보다 큰 항목 필터링
filtered := maputil.Filter(data, func(k string, v int) bool {
    return v > 2
})
// Result: map[string]int{"c": 3, "d": 4}
```

### 2. Transform Values / 값 변환

```go
prices := map[string]int{"apple": 10, "banana": 20, "orange": 30}

// Double all prices / 모든 가격을 2배로
doubled := maputil.MapValues(prices, func(v int) int {
    return v * 2
})
// Result: map[string]int{"apple": 20, "banana": 40, "orange": 60}
```

### 3. Merge Maps / 맵 병합

```go
defaults := map[string]int{"timeout": 30, "retries": 3}
config := map[string]int{"timeout": 60}

// Merge with config overriding defaults / 설정으로 기본값 덮어쓰기
merged := maputil.Merge(defaults, config)
// Result: map[string]int{"timeout": 60, "retries": 3}
```

### 4. Get Keys and Values / 키와 값 가져오기

```go
data := map[string]int{"a": 1, "b": 2, "c": 3}

keys := maputil.Keys(data)      // []string{"a", "b", "c"}
values := maputil.Values(data)  // []int{1, 2, 3}
```

### 5. Group By / 그룹화

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

// Group users by city / 도시별로 사용자 그룹화
byCity := maputil.GroupBy(users, func(u User) string {
    return u.City
})
// Result: map[string][]User{
//     "Seoul": []User{{Name: "Alice", ...}, {Name: "Bob", ...}},
//     "Busan": []User{{Name: "Charlie", ...}},
// }
```

### 6. Pick/Omit Keys / 키 선택/제외

```go
user := map[string]string{
    "id":       "123",
    "name":     "Alice",
    "password": "secret",
    "email":    "alice@example.com",
}

// Pick only public fields / 공개 필드만 선택
public := maputil.Pick(user, "id", "name", "email")
// Result: map[string]string{"id": "123", "name": "Alice", "email": "alice@example.com"}

// Omit sensitive fields / 민감한 필드 제외
safe := maputil.Omit(user, "password")
// Result: map[string]string{"id": "123", "name": "Alice", "email": "alice@example.com"}
```

### 7. Sum/Average / 합계/평균

```go
sales := map[string]int{"Jan": 1000, "Feb": 1500, "Mar": 1200}

total := maputil.Sum(sales)        // 3700
average := maputil.Average(sales)  // 1233.33
```

### 8. Find Min/Max / 최소/최대 찾기

```go
scores := map[string]int{"Alice": 85, "Bob": 92, "Charlie": 78}

// Find highest score / 최고 점수 찾기
name, score, _ := maputil.Max(scores)
// name = "Bob", score = 92
```

### 9. Check Conditions / 조건 확인

```go
ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}

// Check if all adults (>=18) / 모두 성인인지 확인 (>=18)
allAdults := maputil.Every(ages, func(k string, v int) bool {
    return v >= 18
})
// Result: true

// Check if any seniors (>=65) / 노인이 있는지 확인 (>=65)
hasSeniors := maputil.Some(ages, func(k string, v int) bool {
    return v >= 65
})
// Result: false
```

### 10. Diff Maps / 맵 차이 비교

```go
old := map[string]int{"a": 1, "b": 2, "c": 3}
new := map[string]int{"a": 1, "b": 20, "d": 4}

diff := maputil.Diff(old, new)
// Result: map[string]int{"b": 20, "c": 3, "d": 4}
// (changed, removed, and added keys)
```

---

## API Reference / API 참조

This section documents all 81 functions across 10 categories.
이 섹션은 10개 카테고리의 모든 81개 함수를 문서화합니다.

### Category 1: Basic Operations / 기본 작업

Basic map operations for common tasks.
일반적인 작업을 위한 기본 맵 작업입니다.

#### 1.1 Get

```go
func Get[K comparable, V any](m map[K]V, key K) (V, bool)
```

Retrieves a value by key with existence check.
존재 여부 확인과 함께 키로 값을 가져옵니다.

**Example / 예제:**
```go
data := map[string]int{"a": 1, "b": 2}
value, exists := maputil.Get(data, "a")
// value = 1, exists = true

value, exists = maputil.Get(data, "z")
// value = 0 (zero value), exists = false
```

**Use Case / 사용 사례**: Safe value retrieval without panic / 패닉 없이 안전하게 값 가져오기

---

#### 1.2 GetOr

```go
func GetOr[K comparable, V any](m map[K]V, key K, defaultValue V) V
```

Gets a value by key, returns default if key doesn't exist.
키로 값을 가져오고, 키가 없으면 기본값을 반환합니다.

**Example / 예제:**
```go
config := map[string]int{"timeout": 30}
timeout := maputil.GetOr(config, "timeout", 60)    // 30
retries := maputil.GetOr(config, "retries", 3)     // 3 (default)
```

**Use Case / 사용 사례**: Provide sensible defaults for missing config values / 누락된 설정값에 대한 적절한 기본값 제공

---

#### 1.3 Set

```go
func Set[K comparable, V any](m map[K]V, key K, value V) map[K]V
```

Creates a new map with the key-value pair added/updated (immutable).
키-값 쌍이 추가/업데이트된 새 맵을 생성합니다 (불변).

**Example / 예제:**
```go
original := map[string]int{"a": 1, "b": 2}
updated := maputil.Set(original, "c", 3)
// original = map[string]int{"a": 1, "b": 2} (unchanged)
// updated = map[string]int{"a": 1, "b": 2, "c": 3}
```

**Use Case / 사용 사례**: Immutable updates for concurrent scenarios / 동시성 시나리오를 위한 불변 업데이트

---

#### 1.4 Delete

```go
func Delete[K comparable, V any](m map[K]V, key K) map[K]V
```

Creates a new map without the specified key (immutable).
지정된 키가 제거된 새 맵을 생성합니다 (불변).

**Example / 예제:**
```go
data := map[string]int{"a": 1, "b": 2, "c": 3}
filtered := maputil.Delete(data, "b")
// filtered = map[string]int{"a": 1, "c": 3}
```

**Use Case / 사용 사례**: Remove deprecated configuration keys / 더 이상 사용되지 않는 설정 키 제거

---

#### 1.5 Has

```go
func Has[K comparable, V any](m map[K]V, key K) bool
```

Checks if a key exists in the map.
맵에 키가 있는지 확인합니다.

**Example / 예제:**
```go
data := map[string]int{"a": 1, "b": 2}
hasA := maputil.Has(data, "a")  // true
hasZ := maputil.Has(data, "z")  // false
```

**Use Case / 사용 사례**: Validate required keys in configuration / 설정에서 필수 키 검증

---

#### 1.6 IsEmpty

```go
func IsEmpty[K comparable, V any](m map[K]V) bool
```

Checks if map has no elements.
맵에 요소가 없는지 확인합니다.

**Example / 예제:**
```go
empty := map[string]int{}
data := map[string]int{"a": 1}

maputil.IsEmpty(empty)  // true
maputil.IsEmpty(data)   // false
```

**Use Case / 사용 사례**: Validate data before processing / 처리 전 데이터 검증

---

#### 1.7 IsNotEmpty

```go
func IsNotEmpty[K comparable, V any](m map[K]V) bool
```

Checks if map has elements (inverse of IsEmpty).
맵에 요소가 있는지 확인합니다 (IsEmpty의 반대).

**Example / 예제:**
```go
data := map[string]int{"a": 1}
if maputil.IsNotEmpty(data) {
    // Process data
}
```

**Use Case / 사용 사례**: Guard clauses in functions / 함수의 가드 절

---

#### 1.8 Len

```go
func Len[K comparable, V any](m map[K]V) int
```

Gets the number of elements in the map.
맵의 요소 개수를 가져옵니다.

**Example / 예제:**
```go
data := map[string]int{"a": 1, "b": 2, "c": 3}
count := maputil.Len(data)  // 3
```

**Use Case / 사용 사례**: Pagination, statistics, validation / 페이지네이션, 통계, 검증

---

#### 1.9 Clear

```go
func Clear[K comparable, V any](m map[K]V) map[K]V
```

Creates an empty map (immutable).
빈 맵을 생성합니다 (불변).

**Example / 예제:**
```go
data := map[string]int{"a": 1, "b": 2}
cleared := maputil.Clear(data)
// cleared = map[string]int{} (empty)
// data still has original values
```

**Use Case / 사용 사례**: Reset state while preserving map reference / 맵 참조를 유지하면서 상태 재설정

---

#### 1.10 Clone

```go
func Clone[K comparable, V any](m map[K]V) map[K]V
```

Creates a deep copy of the map.
맵의 깊은 복사본을 생성합니다.

**Example / 예제:**
```go
original := map[string]int{"a": 1, "b": 2}
cloned := maputil.Clone(original)
// Modifying cloned doesn't affect original
cloned["c"] = 3
// original = map[string]int{"a": 1, "b": 2}
// cloned = map[string]int{"a": 1, "b": 2, "c": 3}
```

**Use Case / 사용 사례**: Create snapshots, protect against mutations / 스냅샷 생성, 변경으로부터 보호

---

#### 1.11 Equal

```go
func Equal[K comparable, V comparable](m1, m2 map[K]V) bool
```

Compares two maps for deep equality.
두 맵의 깊은 동등성을 비교합니다.

**Example / 예제:**
```go
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"a": 1, "b": 2}
map3 := map[string]int{"a": 1, "b": 3}

maputil.Equal(map1, map2)  // true
maputil.Equal(map1, map3)  // false
```

**Use Case / 사용 사례**: Testing, validation, cache comparisons / 테스트, 검증, 캐시 비교

---

### Category 2: Transformation / 변환

Functions for transforming map structure and values.
맵 구조와 값을 변환하는 함수입니다.

#### 2.1 Map

```go
func Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R
```

Transforms map values using a function, returning a new map with different value type.
함수를 사용하여 맵 값을 변환하고, 다른 값 타입의 새 맵을 반환합니다.

**Example / 예제:**
```go
scores := map[string]int{"math": 85, "english": 92}
grades := maputil.Map(scores, func(subject string, score int) string {
    if score >= 90 {
        return "A"
    }
    return "B"
})
// Result: map[string]string{"math": "B", "english": "A"}
```

**Use Case / 사용 사례**: Convert price integers to formatted strings / 가격 정수를 형식화된 문자열로 변환

---

#### 2.2 MapKeys

```go
func MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V
```

Transforms map keys using a function. If duplicate keys are produced, last wins.
함수를 사용하여 맵 키를 변환합니다. 중복 키가 생성되면 마지막 값이 우선합니다.

**Example / 예제:**
```go
data := map[string]int{"apple": 10, "banana": 20}
uppercased := maputil.MapKeys(data, func(k string, v int) string {
    return strings.ToUpper(k)
})
// Result: map[string]int{"APPLE": 10, "BANANA": 20}
```

**Use Case / 사용 사례**: Standardize key naming conventions / 키 명명 규칙 표준화

---

#### 2.3 MapValues

```go
func MapValues[K comparable, V any, R any](m map[K]V, fn func(V) R) map[K]R
```

Transforms map values using a value-only function.
값만 사용하는 함수로 맵 값을 변환합니다.

**Example / 예제:**
```go
prices := map[string]int{"apple": 10, "banana": 20}
doubled := maputil.MapValues(prices, func(v int) int {
    return v * 2
})
// Result: map[string]int{"apple": 20, "banana": 40}
```

**Use Case / 사용 사례**: Apply discounts, tax calculations / 할인 적용, 세금 계산

---

#### 2.4 MapEntries

```go
func MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](
    m map[K1]V1,
    fn func(K1, V1) (K2, V2),
) map[K2]V2
```

Transforms both keys and values, creating a completely new map type.
키와 값을 모두 변환하여 완전히 새로운 맵 타입을 생성합니다.

**Example / 예제:**
```go
scores := map[string]int{"Alice": 85, "Bob": 92}
reversed := maputil.MapEntries(scores, func(name string, score int) (int, string) {
    return score, name  // Swap key-value
})
// Result: map[int]string{85: "Alice", 92: "Bob"}
```

**Use Case / 사용 사례**: Create reverse lookups, indexes / 역방향 조회, 인덱스 생성

---

#### 2.5 Invert

```go
func Invert[K comparable, V comparable](m map[K]V) map[V]K
```

Swaps keys and values. If multiple keys have the same value, last wins.
키와 값을 교환합니다. 여러 키가 같은 값을 가지면 마지막 키가 우선합니다.

**Example / 예제:**
```go
original := map[string]int{"a": 1, "b": 2, "c": 3}
inverted := maputil.Invert(original)
// Result: map[int]string{1: "a", 2: "b", 3: "c"}
```

**Use Case / 사용 사례**: Bidirectional lookups / 양방향 조회

---

#### 2.6 Flatten

```go
func Flatten[K comparable, V any](m map[K]map[string]V, delimiter string) map[string]V
```

Flattens a nested map into a flat map using a delimiter.
구분자를 사용하여 중첩된 맵을 평면 맵으로 평탄화합니다.

**Example / 예제:**
```go
nested := map[string]map[string]int{
    "user1": {"age": 25, "score": 100},
    "user2": {"age": 30, "score": 95},
}
flat := maputil.Flatten(nested, ".")
// Result: map[string]int{
//     "user1.age": 25,
//     "user1.score": 100,
//     "user2.age": 30,
//     "user2.score": 95,
// }
```

**Use Case / 사용 사례**: Configuration flattening, database denormalization / 설정 평탄화, 데이터베이스 비정규화

---

#### 2.7 Unflatten

```go
func Unflatten[V any](m map[string]V, delimiter string) map[string]interface{}
```

Converts a flat map to a nested structure using a delimiter.
구분자를 사용하여 평면 맵을 중첩 구조로 변환합니다.

**Example / 예제:**
```go
flat := map[string]int{
    "user.name": 1,
    "user.age": 25,
    "admin.name": 2,
}
nested := maputil.Unflatten(flat, ".")
// Result: map[string]interface{}{
//     "user": map[string]interface{}{"name": 1, "age": 25},
//     "admin": map[string]interface{}{"name": 2},
// }
```

**Use Case / 사용 사례**: Parse dotted configuration keys / 점으로 구분된 설정 키 파싱

---

#### 2.8 Chunk

```go
func Chunk[K comparable, V any](m map[K]V, size int) []map[K]V
```

Splits a map into multiple smaller maps of the specified size.
맵을 지정된 크기의 여러 작은 맵으로 분할합니다.

**Example / 예제:**
```go
data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
chunks := maputil.Chunk(data, 2)
// Result: []map[string]int{
//     {"a": 1, "b": 2},
//     {"c": 3, "d": 4},
//     {"e": 5},
// }
```

**Use Case / 사용 사례**: Parallel processing, rate limiting / 병렬 처리, 속도 제한

---

#### 2.9 Partition

```go
func Partition[K comparable, V any](m map[K]V, fn func(K, V) bool) (map[K]V, map[K]V)
```

Splits a map into two maps based on a predicate: (passing, failing).
조건 함수를 기반으로 맵을 두 개로 분할합니다: (통과, 실패).

**Example / 예제:**
```go
scores := map[string]int{"math": 85, "english": 92, "science": 78}
passing, failing := maputil.Partition(scores, func(k string, v int) bool {
    return v >= 80
})
// passing = map[string]int{"math": 85, "english": 92}
// failing = map[string]int{"science": 78}
```

**Use Case / 사용 사례**: Filter data into categories / 데이터를 카테고리로 필터링

---

#### 2.10 Compact

```go
func Compact[K comparable, V comparable](m map[K]V) map[K]V
```

Removes entries with zero values from the map.
맵에서 제로 값을 가진 항목을 제거합니다.

**Example / 예제:**
```go
data := map[string]int{"a": 1, "b": 0, "c": 3, "d": 0}
compacted := maputil.Compact(data)
// Result: map[string]int{"a": 1, "c": 3}
```

**Use Case / 사용 사례**: Remove null/empty values before JSON serialization / JSON 직렬화 전 null/빈 값 제거

---

### Category 3: Aggregation / 집계

Functions for aggregating, grouping, and statistical operations.
집계, 그룹화 및 통계 작업을 위한 함수입니다.

#### 3.1 Reduce

```go
func Reduce[K comparable, V any, R any](m map[K]V, initial R, fn func(R, K, V) R) R
```

Reduces a map to a single value using an accumulator function.
누산기 함수를 사용하여 맵을 단일 값으로 축소합니다.

**Example / 예제:**
```go
sales := map[string]int{"Jan": 1000, "Feb": 1500, "Mar": 1200}
total := maputil.Reduce(sales, 0, func(acc int, month string, amount int) int {
    return acc + amount
})
// Result: 3700
```

**Use Case / 사용 사례**: Complex calculations, custom aggregations / 복잡한 계산, 사용자 정의 집계

---

#### 3.2 Sum

```go
func Sum[K comparable, V Number](m map[K]V) V
```

Calculates the sum of all numeric values.
모든 숫자 값의 합을 계산합니다.

**Example / 예제:**
```go
data := map[string]int{"a": 10, "b": 20, "c": 30}
total := maputil.Sum(data)  // 60
```

**Use Case / 사용 사례**: Financial totals, inventory counts / 재무 총계, 재고 개수

---

#### 3.3 Min

```go
func Min[K comparable, V Ordered](m map[K]V) (K, V, bool)
```

Finds the key-value pair with the minimum value. Returns (key, value, true) if map is not empty.
최소값을 가진 키-값 쌍을 찾습니다. 맵이 비어있지 않으면 (키, 값, true)를 반환합니다.

**Example / 예제:**
```go
prices := map[string]int{"apple": 30, "banana": 10, "orange": 20}
item, price, ok := maputil.Min(prices)
// item = "banana", price = 10, ok = true
```

**Use Case / 사용 사례**: Find worst performer, lowest price / 최악의 성능, 최저 가격 찾기

---

#### 3.4 Max

```go
func Max[K comparable, V Ordered](m map[K]V) (K, V, bool)
```

Finds the key-value pair with the maximum value.
최대값을 가진 키-값 쌍을 찾습니다.

**Example / 예제:**
```go
scores := map[string]int{"Alice": 85, "Bob": 92, "Charlie": 78}
name, score, ok := maputil.Max(scores)
// name = "Bob", score = 92, ok = true
```

**Use Case / 사용 사례**: Find best performer, highest price / 최고 성능, 최고 가격 찾기

---

#### 3.5 MinBy

```go
func MinBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool)
```

Finds the minimum value according to a custom scoring function.
사용자 정의 점수 함수에 따라 최소값을 찾습니다.

**Example / 예제:**
```go
type Product struct {
    Name  string
    Price int
    Stock int
}

products := map[string]Product{
    "p1": {Name: "A", Price: 100, Stock: 5},
    "p2": {Name: "B", Price: 50, Stock: 10},
}

// Find product with lowest stock
id, product, _ := maputil.MinBy(products, func(p Product) float64 {
    return float64(p.Stock)
})
// id = "p1", product.Stock = 5
```

**Use Case / 사용 사례**: Custom scoring for minimum selection / 최소 선택을 위한 사용자 정의 점수

---

#### 3.6 MaxBy

```go
func MaxBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool)
```

Finds the maximum value according to a custom scoring function.
사용자 정의 점수 함수에 따라 최대값을 찾습니다.

**Example / 예제:**
```go
products := map[string]Product{
    "p1": {Name: "A", Price: 100},
    "p2": {Name: "B", Price: 50},
}

// Find most expensive product
id, product, _ := maputil.MaxBy(products, func(p Product) float64 {
    return float64(p.Price)
})
// id = "p1", product.Price = 100
```

**Use Case / 사용 사례**: Custom scoring for maximum selection / 최대 선택을 위한 사용자 정의 점수

---

#### 3.7 Average

```go
func Average[K comparable, V Number](m map[K]V) float64
```

Calculates the average of all numeric values. Returns 0 if map is empty.
모든 숫자 값의 평균을 계산합니다. 맵이 비어있으면 0을 반환합니다.

**Example / 예제:**
```go
scores := map[string]int{"test1": 80, "test2": 90, "test3": 85}
avg := maputil.Average(scores)  // 85.0
```

**Use Case / 사용 사례**: Statistics, performance metrics / 통계, 성능 메트릭

---

#### 3.8 GroupBy

```go
func GroupBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G][]V
```

Groups a slice of elements by a key extracted from each element.
각 요소에서 추출한 키로 요소 슬라이스를 그룹화합니다.

**Example / 예제:**
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

byCity := maputil.GroupBy(users, func(u User) string {
    return u.City
})
// Result: map[string][]User{
//     "Seoul": []User{{Name: "Alice", ...}, {Name: "Bob", ...}},
//     "Busan": []User{{Name: "Charlie", ...}},
// }
```

**Use Case / 사용 사례**: Data categorization, reporting / 데이터 분류, 보고

---

#### 3.9 CountBy

```go
func CountBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G]int
```

Counts the number of elements in a slice for each key extracted from each element.
각 요소에서 추출한 키별로 슬라이스의 요소 수를 계산합니다.

**Example / 예제:**
```go
type Transaction struct {
    Type   string
    Amount int
}

txns := []Transaction{
    {Type: "income", Amount: 1000},
    {Type: "expense", Amount: 500},
    {Type: "income", Amount: 1500},
}

counts := maputil.CountBy(txns, func(t Transaction) string {
    return t.Type
})
// Result: map[string]int{"income": 2, "expense": 1}
```

**Use Case / 사용 사례**: Statistics, frequency analysis, histograms / 통계, 빈도 분석, 히스토그램

---

*[The USER_MANUAL continues with Categories 4-10, Usage Patterns, Common Use Cases, Best Practices, Performance Considerations, Troubleshooting, FAQ, and Additional Resources sections...]*

*Due to length limitations, this is the first part of the USER_MANUAL. The complete document would continue with the remaining categories and sections following the same detailed format.*

