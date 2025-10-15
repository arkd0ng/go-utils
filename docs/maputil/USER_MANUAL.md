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

### Category 4: Merge Operations / 병합 작업

Functions for merging and combining maps.
맵을 병합하고 결합하는 함수입니다.

#### 4.1 Merge

```go
func Merge[K comparable, V any](maps ...map[K]V) map[K]V
```

Merges multiple maps into a single map. If duplicate keys exist, the value from the last map wins.
여러 맵을 단일 맵으로 병합합니다. 중복 키가 있으면 마지막 맵의 값이 우선합니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}
m3 := map[string]int{"c": 5, "d": 6}
result := maputil.Merge(m1, m2, m3)
// Result: map[string]int{"a": 1, "b": 3, "c": 5, "d": 6}
```

**Use Case / 사용 사례**: Merge configuration from multiple sources (defaults, environment, user settings) / 여러 소스의 설정 병합 (기본값, 환경, 사용자 설정)

---

#### 4.2 MergeWith

```go
func MergeWith[K comparable, V any](fn func(V, V) V, maps ...map[K]V) map[K]V
```

Merges multiple maps using a custom resolver function for duplicate keys. The resolver receives the old and new values and returns the value to use.
중복 키에 대해 사용자 정의 해결 함수를 사용하여 여러 맵을 병합합니다. 해결 함수는 이전 값과 새 값을 받아 사용할 값을 반환합니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}
result := maputil.MergeWith(
    func(old, new int) int { return old + new }, // Sum on conflict
    m1, m2,
)
// Result: map[string]int{"a": 1, "b": 5, "c": 4}
```

**Use Case / 사용 사례**: Aggregate metrics from multiple sources (sum counters, take max, etc.) / 여러 소스의 메트릭 집계 (카운터 합산, 최댓값 선택 등)

---

#### 4.3 DeepMerge

```go
func DeepMerge(maps ...map[string]interface{}) map[string]interface{}
```

Performs a deep merge of nested string maps. Nested maps are recursively merged. For non-map values, the last value wins.
중첩된 문자열 맵을 깊이 병합합니다. 중첩된 맵은 재귀적으로 병합됩니다. 맵이 아닌 값의 경우 마지막 값이 우선합니다.

**Example / 예제:**
```go
m1 := map[string]interface{}{
    "user": map[string]interface{}{"name": "Alice", "age": 25},
}
m2 := map[string]interface{}{
    "user": map[string]interface{}{"age": 26, "city": "Seoul"},
}
result := maputil.DeepMerge(m1, m2)
// Result: {"user": {"name": "Alice", "age": 26, "city": "Seoul"}}
```

**Use Case / 사용 사례**: Merge nested configuration objects / 중첩된 설정 객체 병합

---

#### 4.4 Union

```go
func Union[K comparable, V any](maps ...map[K]V) map[K]V
```

Alias for Merge. Returns a union of all input maps.
Merge의 별칭입니다. 모든 입력 맵의 합집합을 반환합니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"c": 3, "d": 4}
result := maputil.Union(m1, m2)
// Result: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
```

**Use Case / 사용 사례**: Combine feature flags, permissions, or tags from multiple sources / 여러 소스의 기능 플래그, 권한 또는 태그 결합

---

#### 4.5 Intersection

```go
func Intersection[K comparable, V any](maps ...map[K]V) map[K]V
```

Returns a map containing only keys that exist in all input maps. Values are taken from the first map.
모든 입력 맵에 존재하는 키만 포함하는 맵을 반환합니다. 값은 첫 번째 맵에서 가져옵니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 20, "c": 30, "d": 40}
m3 := map[string]int{"c": 100, "d": 200}
result := maputil.Intersection(m1, m2, m3)
// Result: map[string]int{"c": 3}
```

**Use Case / 사용 사례**: Find common permissions across roles, or shared feature flags / 역할 간 공통 권한 찾기 또는 공유 기능 플래그

---

#### 4.6 Difference

```go
func Difference[K comparable, V any](m1, m2 map[K]V) map[K]V
```

Returns a map containing keys from the first map that are not in the second map.
첫 번째 맵에는 있지만 두 번째 맵에는 없는 키를 포함하는 맵을 반환합니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 20, "d": 40}
result := maputil.Difference(m1, m2)
// Result: map[string]int{"a": 1, "c": 3}
```

**Use Case / 사용 사례**: Find removed items, deprecated keys, or revoked permissions / 제거된 항목, 더 이상 사용되지 않는 키 또는 취소된 권한 찾기

---

#### 4.7 SymmetricDifference

```go
func SymmetricDifference[K comparable, V any](m1, m2 map[K]V) map[K]V
```

Returns a map containing keys that exist in either map but not in both (XOR operation).
어느 한 맵에는 있지만 두 맵 모두에는 없는 키를 포함하는 맵을 반환합니다 (XOR 연산).

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 20, "c": 30, "d": 40}
result := maputil.SymmetricDifference(m1, m2)
// Result: map[string]int{"a": 1, "d": 40}
```

**Use Case / 사용 사례**: Find unique features between two systems, or exclusive permissions / 두 시스템 간 고유 기능 또는 배타적 권한 찾기

---

#### 4.8 Assign

```go
func Assign[K comparable, V any](target map[K]V, sources ...map[K]V) map[K]V
```

Merges source maps into the target map (mutating operation). **WARNING**: This function modifies the target map in place.
소스 맵을 대상 맵에 병합합니다 (변경 작업). **경고**: 이 함수는 대상 맵을 직접 수정합니다.

**Example / 예제:**
```go
target := map[string]int{"a": 1, "b": 2}
source := map[string]int{"b": 3, "c": 4}
result := maputil.Assign(target, source)
// target is modified: map[string]int{"a": 1, "b": 3, "c": 4}
// result points to the same map as target
```

**Use Case / 사용 사례**: Performance-critical code where mutation is acceptable and memory allocation should be minimized / 변경이 허용되고 메모리 할당을 최소화해야 하는 성능 중심 코드

---

### Category 5: Filter Operations / 필터 작업

Functions for filtering map entries based on predicates.
조건에 따라 맵 항목을 필터링하는 함수입니다.

#### 5.1 Filter

```go
func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V
```

Returns a new map containing only the entries that satisfy the predicate.
조건을 만족하는 항목만 포함하는 새 맵을 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := maputil.Filter(m, func(k string, v int) bool {
    return v > 2
})
// Result: map[string]int{"c": 3, "d": 4}
```

**Use Case / 사용 사례**: Filter user data by criteria, extract active sessions, or remove invalid entries / 기준에 따라 사용자 데이터 필터링, 활성 세션 추출 또는 잘못된 항목 제거

---

#### 5.2 FilterKeys

```go
func FilterKeys[K comparable, V any](m map[K]V, fn func(K) bool) map[K]V
```

Returns a new map containing only the entries with keys that satisfy the predicate.
조건을 만족하는 키를 가진 항목만 포함하는 새 맵을 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
result := maputil.FilterKeys(m, func(k string) bool {
    return len(k) > 5
})
// Result: map[string]int{"banana": 2, "cherry": 3}
```

**Use Case / 사용 사례**: Filter by naming patterns, extract prefixed keys, or select keys matching regex / 명명 패턴으로 필터링, 접두사가 있는 키 추출 또는 정규식과 일치하는 키 선택

---

#### 5.3 FilterValues

```go
func FilterValues[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V
```

Returns a new map containing only the entries with values that satisfy the predicate.
조건을 만족하는 값을 가진 항목만 포함하는 새 맵을 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := maputil.FilterValues(m, func(v int) bool {
    return v%2 == 0
})
// Result: map[string]int{"b": 2, "d": 4}
```

**Use Case / 사용 사례**: Filter even/odd values, extract values within range, or select by value type / 짝수/홀수 값 필터링, 범위 내 값 추출 또는 값 타입으로 선택

---

#### 5.4 Pick

```go
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V
```

Returns a new map containing only the specified keys. Keys that don't exist in the original map are ignored.
지정된 키만 포함하는 새 맵을 반환합니다. 원본 맵에 존재하지 않는 키는 무시됩니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := maputil.Pick(m, "a", "c", "e")
// Result: map[string]int{"a": 1, "c": 3}
```

**Use Case / 사용 사례**: Extract public fields from user data, select specific columns, or whitelist keys / 사용자 데이터에서 공개 필드 추출, 특정 열 선택 또는 화이트리스트 키

---

#### 5.5 Omit

```go
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V
```

Returns a new map excluding the specified keys.
지정된 키를 제외한 새 맵을 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := maputil.Omit(m, "b", "d")
// Result: map[string]int{"a": 1, "c": 3}
```

**Use Case / 사용 사례**: Remove sensitive fields (password, token), exclude internal fields, or blacklist keys / 민감한 필드 제거 (비밀번호, 토큰), 내부 필드 제외 또는 블랙리스트 키

---

#### 5.6 PickBy

```go
func PickBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V
```

Returns a new map containing only entries that satisfy the predicate. This is an alias for Filter.
조건을 만족하는 항목만 포함하는 새 맵을 반환합니다. 이것은 Filter의 별칭입니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := maputil.PickBy(m, func(k string, v int) bool {
    return v > 2
})
// Result: map[string]int{"c": 3, "d": 4}
```

**Use Case / 사용 사례**: Same as Filter - conditional data selection / Filter와 동일 - 조건부 데이터 선택

---

#### 5.7 OmitBy

```go
func OmitBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V
```

Returns a new map excluding entries that satisfy the predicate. This is the inverse of Filter.
조건을 만족하는 항목을 제외한 새 맵을 반환합니다. 이것은 Filter의 반대입니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := maputil.OmitBy(m, func(k string, v int) bool {
    return v%2 == 0
})
// Result: map[string]int{"a": 1, "c": 3}
```

**Use Case / 사용 사례**: Remove entries matching criteria, filter out invalid data, or exclude by condition / 기준과 일치하는 항목 제거, 잘못된 데이터 필터링 또는 조건별 제외

---

### Category 6: Conversion / 변환

Functions for converting between maps, slices, and JSON.
맵, 슬라이스 및 JSON 간 변환을 위한 함수입니다.

#### 6.1 Keys

```go
func Keys[K comparable, V any](m map[K]V) []K
```

Returns all keys from the map as a slice. The order is not guaranteed (maps are unordered).
맵의 모든 키를 슬라이스로 반환합니다. 순서는 보장되지 않습니다 (맵은 순서가 없음).

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
keys := maputil.Keys(m)
// Result: []string{"a", "b", "c"} (order may vary)
```

**Use Case / 사용 사례**: Extract all keys for iteration, validation, or display / 반복, 검증 또는 표시를 위해 모든 키 추출

---

#### 6.2 Values

```go
func Values[K comparable, V any](m map[K]V) []V
```

Returns all values from the map as a slice. The order is not guaranteed.
맵의 모든 값을 슬라이스로 반환합니다. 순서는 보장되지 않습니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
values := maputil.Values(m)
// Result: []int{1, 2, 3} (order may vary)
```

**Use Case / 사용 사례**: Extract all values for aggregation, statistics, or bulk operations / 집계, 통계 또는 대량 작업을 위해 모든 값 추출

---

#### 6.3 Entries

```go
func Entries[K comparable, V any](m map[K]V) []Entry[K, V]
```

Returns all key-value pairs from the map as a slice of Entry structs.
맵의 모든 키-값 쌍을 Entry 구조체의 슬라이스로 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2}
entries := maputil.Entries(m)
// Result: []Entry[string, int]{
//     {Key: "a", Value: 1},
//     {Key: "b", Value: 2},
// } (order may vary)
```

**Use Case / 사용 사례**: Convert map to sortable slice, serialize to structured format / 맵을 정렬 가능한 슬라이스로 변환, 구조화된 형식으로 직렬화

---

#### 6.4 FromEntries

```go
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V
```

Creates a map from a slice of Entry structs. If duplicate keys exist, the last entry wins.
Entry 구조체의 슬라이스로부터 맵을 생성합니다. 중복 키가 있으면 마지막 항목이 우선합니다.

**Example / 예제:**
```go
entries := []Entry[string, int]{
    {Key: "a", Value: 1},
    {Key: "b", Value: 2},
}
m := maputil.FromEntries(entries)
// Result: map[string]int{"a": 1, "b": 2}
```

**Use Case / 사용 사례**: Reconstruct map from sorted entries, deserialize structured data / 정렬된 항목에서 맵 재구성, 구조화된 데이터 역직렬화

---

#### 6.5 ToJSON

```go
func ToJSON[K comparable, V any](m map[K]V) (string, error)
```

Converts a map to a JSON string. Returns an error if the map cannot be marshaled to JSON.
맵을 JSON 문자열로 변환합니다. 맵을 JSON으로 마샬링할 수 없으면 에러를 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
json, err := maputil.ToJSON(m)
// Result: `{"a":1,"b":2,"c":3}`
```

**Use Case / 사용 사례**: Serialize map for API response, logging, or storage / API 응답, 로깅 또는 저장소를 위해 맵 직렬화

---

#### 6.6 FromJSON

```go
func FromJSON[K comparable, V any](jsonStr string, m *map[K]V) error
```

Parses a JSON string into a map. Returns an error if the JSON string cannot be unmarshaled.
JSON 문자열을 맵으로 파싱합니다. JSON 문자열을 언마샬링할 수 없으면 에러를 반환합니다.

**Example / 예제:**
```go
var m map[string]int
err := maputil.FromJSON(`{"a":1,"b":2,"c":3}`, &m)
// m = map[string]int{"a": 1, "b": 2, "c": 3}
```

**Use Case / 사용 사례**: Parse API request body, configuration files, or stored data / API 요청 본문, 설정 파일 또는 저장된 데이터 파싱

---

#### 6.7 ToSlice

```go
func ToSlice[K comparable, V any, R any](m map[K]V, fn func(K, V) R) []R
```

Converts a map to a slice using a transformation function.
변환 함수를 사용하여 맵을 슬라이스로 변환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
slice := maputil.ToSlice(m, func(k string, v int) string {
    return fmt.Sprintf("%s=%d", k, v)
})
// Result: []string{"a=1", "b=2", "c=3"} (order may vary)
```

**Use Case / 사용 사례**: Convert map to formatted strings, create display list, or transform to DTOs / 맵을 형식화된 문자열로 변환, 표시 목록 생성 또는 DTO로 변환

---

#### 6.8 FromSlice

```go
func FromSlice[K comparable, V any](slice []V, fn func(V) K) map[K]V
```

Creates a map from a slice using a key extraction function. If duplicate keys exist, the last value wins.
키 추출 함수를 사용하여 슬라이스로부터 맵을 생성합니다. 중복 키가 있으면 마지막 값이 우선합니다.

**Example / 예제:**
```go
type User struct { ID int; Name string }
users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
m := maputil.FromSlice(users, func(u User) int {
    return u.ID
})
// Result: map[int]User{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}}
```

**Use Case / 사용 사례**: Index slice by ID, create lookup tables, or convert list to map / ID로 슬라이스 인덱싱, 조회 테이블 생성 또는 목록을 맵으로 변환

---

#### 6.9 FromSliceBy

```go
func FromSliceBy[K comparable, V any, R any](slice []V, keyFn func(V) K, valueFn func(V) R) map[K]R
```

Creates a map from a slice using key and value extraction functions.
키와 값 추출 함수를 사용하여 슬라이스로부터 맵을 생성합니다.

**Example / 예제:**
```go
type User struct { ID int; Name string; Age int }
users := []User{{ID: 1, Name: "Alice", Age: 25}, {ID: 2, Name: "Bob", Age: 30}}
m := maputil.FromSliceBy(users,
    func(u User) int { return u.ID },
    func(u User) string { return u.Name },
)
// Result: map[int]string{1: "Alice", 2: "Bob"}
```

**Use Case / 사용 사례**: Create ID-to-name mappings, extract specific fields, or build custom indexes / ID-이름 매핑 생성, 특정 필드 추출 또는 사용자 정의 인덱스 구축

---

### Category 7: Predicate Checks / 조건 검사

Functions for checking conditions on map entries.
맵 항목에 대한 조건을 확인하는 함수입니다.

#### 7.1 Every

```go
func Every[K comparable, V any](m map[K]V, fn func(K, V) bool) bool
```

Checks whether all entries in the map satisfy the predicate. Returns true for empty maps.
맵의 모든 항목이 조건을 만족하는지 확인합니다. 빈 맵의 경우 true를 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 2, "b": 4, "c": 6}
allEven := maputil.Every(m, func(k string, v int) bool {
    return v%2 == 0
})
// Result: true
```

**Use Case / 사용 사례**: Validate all entries meet criteria, check all users are active, or verify data integrity / 모든 항목이 기준을 충족하는지 검증, 모든 사용자가 활성인지 확인 또는 데이터 무결성 확인

---

#### 7.2 Some

```go
func Some[K comparable, V any](m map[K]V, fn func(K, V) bool) bool
```

Checks whether at least one entry in the map satisfies the predicate. Returns false for empty maps.
맵의 최소 하나의 항목이 조건을 만족하는지 확인합니다. 빈 맵의 경우 false를 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
hasEven := maputil.Some(m, func(k string, v int) bool {
    return v%2 == 0
})
// Result: true (because of "b": 2)
```

**Use Case / 사용 사례**: Check if any entry meets criteria, detect at least one error, or find any match / 항목이 기준을 충족하는지 확인, 최소 하나의 오류 감지 또는 일치 항목 찾기

---

#### 7.3 None

```go
func None[K comparable, V any](m map[K]V, fn func(K, V) bool) bool
```

Checks whether no entries in the map satisfy the predicate. Returns true for empty maps.
맵의 어떤 항목도 조건을 만족하지 않는지 확인합니다. 빈 맵의 경우 true를 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 3, "c": 5}
noEven := maputil.None(m, func(k string, v int) bool {
    return v%2 == 0
})
// Result: true
```

**Use Case / 사용 사례**: Verify no errors exist, check no invalid entries, or confirm absence of condition / 오류가 없는지 확인, 잘못된 항목이 없는지 확인 또는 조건의 부재 확인

---

#### 7.4 HasValue

```go
func HasValue[K comparable, V comparable](m map[K]V, value V) bool
```

Checks whether a value exists in the map.
맵에 값이 존재하는지 확인합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
exists := maputil.HasValue(m, 2)  // true
exists = maputil.HasValue(m, 5)   // false
```

**Use Case / 사용 사례**: Check if specific value is present, validate value existence, or search by value / 특정 값이 있는지 확인, 값 존재 검증 또는 값으로 검색

---

#### 7.5 HasEntry

```go
func HasEntry[K comparable, V comparable](m map[K]V, key K, value V) bool
```

Checks whether a specific key-value pair exists in the map.
특정 키-값 쌍이 맵에 존재하는지 확인합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
exists := maputil.HasEntry(m, "b", 2)  // true
exists = maputil.HasEntry(m, "b", 3)   // false
```

**Use Case / 사용 사례**: Validate exact key-value match, check configuration setting, or verify state / 정확한 키-값 일치 검증, 설정 확인 또는 상태 확인

---

#### 7.6 IsSubset

```go
func IsSubset[K comparable, V comparable](subset, superset map[K]V) bool
```

Checks whether the first map is a subset of the second map. A map is a subset if all its key-value pairs exist in the superset.
첫 번째 맵이 두 번째 맵의 부분집합인지 확인합니다. 모든 키-값 쌍이 상위집합에 존재하면 부분집합입니다.

**Example / 예제:**
```go
subset := map[string]int{"a": 1, "b": 2}
superset := map[string]int{"a": 1, "b": 2, "c": 3}
result := maputil.IsSubset(subset, superset)
// Result: true
```

**Use Case / 사용 사례**: Verify permissions are subset of allowed permissions, check feature subset, or validate configuration / 권한이 허용된 권한의 부분집합인지 확인, 기능 부분집합 확인 또는 설정 검증

---

#### 7.7 IsSuperset

```go
func IsSuperset[K comparable, V comparable](superset, subset map[K]V) bool
```

Checks whether a map is a superset of another map. A map is a superset if it contains all key-value pairs from the subset.
맵이 다른 맵의 상위집합인지 확인합니다. 맵이 subset의 모든 키-값 쌍을 포함하면 상위집합입니다.

**Example / 예제:**
```go
superset := map[string]int{"a": 1, "b": 2, "c": 3}
subset := map[string]int{"a": 1, "b": 2}
result := maputil.IsSuperset(superset, subset)
// Result: true
```

**Use Case / 사용 사례**: Check if permissions include required set, verify feature availability, or validate coverage / 권한에 필수 세트가 포함되어 있는지 확인, 기능 가용성 확인 또는 범위 검증

---

### Category 8: Key Operations / 키 작업

Functions for manipulating and querying map keys.
맵 키를 조작하고 쿼리하는 함수입니다.

#### 8.1 KeysSorted

```go
func KeysSorted[K Ordered, V any](m map[K]V) []K
```

Returns all keys from the map as a sorted slice.
맵의 모든 키를 정렬된 슬라이스로 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"c": 3, "a": 1, "b": 2}
keys := maputil.KeysSorted(m)
// Result: []string{"a", "b", "c"}
```

**Use Case / 사용 사례**: Sorted iteration, deterministic output, or ordered display / 정렬된 반복, 결정적 출력 또는 정렬된 표시

---

#### 8.2 FindKey

```go
func FindKey[K comparable, V any](m map[K]V, fn func(K, V) bool) (K, bool)
```

Finds the first key that satisfies the predicate. Returns the key and true if found, or zero value and false otherwise.
조건을 만족하는 첫 번째 키를 찾습니다. 찾으면 키와 true를 반환하고, 그렇지 않으면 zero 값과 false를 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
key, found := maputil.FindKey(m, func(k string, v int) bool {
    return v > 2
})
// key = "c", found = true (may vary due to map iteration order)
```

**Use Case / 사용 사례**: Find first match, search by condition, or locate specific entry / 첫 번째 일치 항목 찾기, 조건으로 검색 또는 특정 항목 찾기

---

#### 8.3 FindKeys

```go
func FindKeys[K comparable, V any](m map[K]V, fn func(K, V) bool) []K
```

Finds all keys that satisfy the predicate.
조건을 만족하는 모든 키를 찾습니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
keys := maputil.FindKeys(m, func(k string, v int) bool {
    return v > 2
})
// Result: []string{"c", "d"} (order may vary)
```

**Use Case / 사용 사례**: Find all matching keys, extract keys by criteria, or filter key set / 일치하는 모든 키 찾기, 기준별 키 추출 또는 키 세트 필터링

---

#### 8.4 RenameKey

```go
func RenameKey[K comparable, V any](m map[K]V, oldKey, newKey K) map[K]V
```

Creates a new map with a key renamed. If the old key doesn't exist, returns a clone. If the new key already exists, it will be overwritten.
키의 이름이 변경된 새 맵을 생성합니다. 이전 키가 존재하지 않으면 복사본을 반환합니다. 새 키가 이미 존재하면 덮어씁니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
result := maputil.RenameKey(m, "b", "B")
// Result: map[string]int{"a": 1, "B": 2, "c": 3}
```

**Use Case / 사용 사례**: Standardize key names, migrate field names, or rename API fields / 키 이름 표준화, 필드 이름 마이그레이션 또는 API 필드 이름 변경

---

#### 8.5 SwapKeys

```go
func SwapKeys[K comparable, V any](m map[K]V, key1, key2 K) map[K]V
```

Creates a new map with two keys swapped. If either key doesn't exist, returns a clone.
두 키가 교환된 새 맵을 생성합니다. 어느 한 키라도 존재하지 않으면 복사본을 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
result := maputil.SwapKeys(m, "a", "b")
// Result: map[string]int{"a": 2, "b": 1, "c": 3}
```

**Use Case / 사용 사례**: Swap configuration values, exchange priorities, or reorder entries / 설정 값 교환, 우선순위 교환 또는 항목 재정렬

---

#### 8.6 PrefixKeys

```go
func PrefixKeys[V any](m map[string]V, prefix string) map[string]V
```

Adds a prefix to all keys in a map.
맵의 모든 키에 접두사를 추가합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2}
result := maputil.PrefixKeys(m, "key_")
// Result: map[string]int{"key_a": 1, "key_b": 2}
```

**Use Case / 사용 사례**: Namespace keys, add environment prefix, or categorize entries / 키 네임스페이스 지정, 환경 접두사 추가 또는 항목 분류

---

#### 8.7 SuffixKeys

```go
func SuffixKeys[V any](m map[string]V, suffix string) map[string]V
```

Adds a suffix to all keys in a map.
맵의 모든 키에 접미사를 추가합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2}
result := maputil.SuffixKeys(m, "_key")
// Result: map[string]int{"a_key": 1, "b_key": 2}
```

**Use Case / 사용 사례**: Add version suffix, append unit suffix, or tag entries / 버전 접미사 추가, 단위 접미사 추가 또는 항목 태그 지정

---

#### 8.8 TransformKeys

```go
func TransformKeys[K comparable, V any](m map[K]V, fn func(K) K) map[K]V
```

Transforms all keys in a map using a function. If multiple keys map to the same transformed key, the last value wins.
함수를 사용하여 맵의 모든 키를 변환합니다. 여러 키가 같은 변환된 키로 매핑되면 마지막 값이 우선합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
result := maputil.TransformKeys(m, func(k string) string {
    return strings.ToUpper(k)
})
// Result: map[string]int{"A": 1, "B": 2, "C": 3}
```

**Use Case / 사용 사례**: Normalize key casing, sanitize keys, or apply key transformations / 키 대소문자 정규화, 키 정리 또는 키 변환 적용

---

### Category 9: Value Operations / 값 작업

Functions for manipulating and querying map values.
맵 값을 조작하고 쿼리하는 함수입니다.

#### 9.1 ValuesSorted

```go
func ValuesSorted[K comparable, V Ordered](m map[K]V) []V
```

Returns all values from the map as a sorted slice.
맵의 모든 값을 정렬된 슬라이스로 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 3, "b": 1, "c": 2}
values := maputil.ValuesSorted(m)
// Result: []int{1, 2, 3}
```

**Use Case / 사용 사례**: Sorted statistics, ordered display, or ranked lists / 정렬된 통계, 정렬된 표시 또는 순위 목록

---

#### 9.2 UniqueValues

```go
func UniqueValues[K comparable, V comparable](m map[K]V) []V
```

Returns a slice of unique values from the map.
맵의 고유한 값들의 슬라이스를 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3, "e": 2}
unique := maputil.UniqueValues(m)
// Result: []int{1, 2, 3} (order may vary)
```

**Use Case / 사용 사례**: Find distinct values, remove duplicates, or get value set / 고유 값 찾기, 중복 제거 또는 값 세트 가져오기

---

#### 9.3 ReplaceValue

```go
func ReplaceValue[K comparable, V comparable](m map[K]V, oldValue, newValue V) map[K]V
```

Creates a new map with all occurrences of a value replaced.
특정 값의 모든 발생을 대체한 새 맵을 생성합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3}
result := maputil.ReplaceValue(m, 1, 10)
// Result: map[string]int{"a": 10, "b": 2, "c": 10, "d": 3}
```

**Use Case / 사용 사례**: Replace default values, normalize data, or fix incorrect values / 기본값 교체, 데이터 정규화 또는 잘못된 값 수정

---

#### 9.4 UpdateValues

```go
func UpdateValues[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V
```

Creates a new map with all values transformed by the function.
함수로 모든 값을 변환한 새 맵을 생성합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
result := maputil.UpdateValues(m, func(k string, v int) int {
    return v * 10
})
// Result: map[string]int{"a": 10, "b": 20, "c": 30}
```

**Use Case / 사용 사례**: Bulk value updates, apply transformations, or recalculate values / 대량 값 업데이트, 변환 적용 또는 값 재계산

---

#### 9.5 MinValue

```go
func MinValue[K comparable, V Ordered](m map[K]V) (V, bool)
```

Returns the minimum value from a map. Returns the value and true if found, or zero value and false if map is empty.
맵에서 최솟값을 반환합니다. 찾은 경우 값과 true를, 맵이 비어있으면 제로값과 false를 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 3, "b": 1, "c": 2}
min, found := maputil.MinValue(m)
// min = 1, found = true
```

**Use Case / 사용 사례**: Find lowest price, minimum score, or smallest quantity / 최저 가격, 최소 점수 또는 최소 수량 찾기

---

#### 9.6 MaxValue

```go
func MaxValue[K comparable, V Ordered](m map[K]V) (V, bool)
```

Returns the maximum value from a map. Returns the value and true if found, or zero value and false if map is empty.
맵에서 최댓값을 반환합니다. 찾은 경우 값과 true를, 맵이 비어있으면 제로값과 false를 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 3, "b": 1, "c": 2}
max, found := maputil.MaxValue(m)
// max = 3, found = true
```

**Use Case / 사용 사례**: Find highest price, maximum score, or largest quantity / 최고 가격, 최대 점수 또는 최대 수량 찾기

---

#### 9.7 SumValues

```go
func SumValues[K comparable, V Number](m map[K]V) V
```

Returns the sum of all values in a map. Returns zero if the map is empty.
맵의 모든 값의 합을 반환합니다. 맵이 비어있으면 0을 반환합니다.

**Example / 예제:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
sum := maputil.SumValues(m)
// sum = 6
```

**Use Case / 사용 사례**: Total sales, aggregate counts, or sum metrics / 총 판매, 집계 개수 또는 메트릭 합계

---

### Category 10: Comparison / 비교

Functions for comparing maps and finding differences.
맵을 비교하고 차이를 찾는 함수입니다.

#### 10.1 Diff

```go
func Diff[K comparable, V comparable](m1, m2 map[K]V) map[K]V
```

Returns a map containing entries that differ between two maps. An entry differs if the key exists in both maps but values are different, or the key exists in only one map.
두 맵 간에 다른 항목을 포함하는 맵을 반환합니다. 키가 두 맵 모두에 존재하지만 값이 다르거나, 키가 한 맵에만 존재하면 항목이 다릅니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"a": 1, "b": 20, "d": 4}
diff := maputil.Diff(m1, m2)
// Result: map[string]int{"b": 20, "c": 3, "d": 4}
```

**Use Case / 사용 사례**: Detect configuration changes, track state differences, or find updates / 설정 변경 감지, 상태 차이 추적 또는 업데이트 찾기

---

#### 10.2 DiffKeys

```go
func DiffKeys[K comparable, V comparable](m1, m2 map[K]V) []K
```

Returns a slice of keys that differ between two maps. A key differs if it exists in one map but not the other, or if it exists in both maps but values are different.
두 맵 간에 다른 키를 슬라이스로 반환합니다. 한 맵에는 존재하지만 다른 맵에는 없거나, 두 맵 모두에 존재하지만 값이 다르면 키가 다릅니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"a": 1, "b": 20, "d": 4}
keys := maputil.DiffKeys(m1, m2)
// Result: []string{"b", "c", "d"} (order may vary)
```

**Use Case / 사용 사례**: Find changed keys, detect schema differences, or identify updates / 변경된 키 찾기, 스키마 차이 감지 또는 업데이트 식별

---

#### 10.3 Compare

```go
func Compare[K comparable, V comparable](m1, m2 map[K]V) (added, removed, modified map[K]V)
```

Performs a detailed comparison of two maps. Returns three maps: added (entries in m2 but not m1), removed (entries in m1 but not m2), and modified (entries in both but with different values from m2).
두 맵을 상세하게 비교합니다. 세 개의 맵을 반환합니다: added (m2에는 있지만 m1에는 없음), removed (m1에는 있지만 m2에는 없음), modified (두 맵 모두에 존재하지만 m2의 값이 다름).

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"a": 1, "b": 20, "d": 4}
added, removed, modified := maputil.Compare(m1, m2)
// added = {"d": 4}
// removed = {"c": 3}
// modified = {"b": 20}
```

**Use Case / 사용 사례**: Generate change logs, track data migrations, or show before/after differences / 변경 로그 생성, 데이터 마이그레이션 추적 또는 이전/이후 차이 표시

---

#### 10.4 CommonKeys

```go
func CommonKeys[K comparable, V any](maps ...map[K]V) []K
```

Returns a slice of keys that exist in all input maps.
모든 입력 맵에 존재하는 키를 슬라이스로 반환합니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 20, "c": 30, "d": 40}
m3 := map[string]int{"c": 100, "d": 200}
common := maputil.CommonKeys(m1, m2, m3)
// Result: []string{"c"}
```

**Use Case / 사용 사례**: Find shared keys, intersection analysis, or common fields / 공유 키 찾기, 교집합 분석 또는 공통 필드

---

#### 10.5 AllKeys

```go
func AllKeys[K comparable, V any](maps ...map[K]V) []K
```

Returns a slice of all unique keys from all input maps.
모든 입력 맵의 고유한 키를 슬라이스로 반환합니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 20, "c": 30}
m3 := map[string]int{"c": 100, "d": 200}
all := maputil.AllKeys(m1, m2, m3)
// Result: []string{"a", "b", "c", "d"} (order may vary)
```

**Use Case / 사용 사례**: Get complete key set, union analysis, or all fields / 완전한 키 세트 가져오기, 합집합 분석 또는 모든 필드

---

#### 10.6 EqualMaps

```go
func EqualMaps[K comparable, V comparable](m1, m2 map[K]V) bool
```

Checks whether two maps are equal. Two maps are equal if they have the same keys and values.
두 맵이 동일한지 확인합니다. 두 맵은 같은 키와 값을 가지면 동일합니다.

**Example / 예제:**
```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"a": 1, "b": 2, "c": 3}
equal := maputil.EqualMaps(m1, m2)
// Result: true
```

**Use Case / 사용 사례**: Testing, cache validation, or state comparison / 테스트, 캐시 검증 또는 상태 비교

---

## Usage Patterns / 사용 패턴

*[Content to be added in full version]*

---

## Common Use Cases / 일반적인 사용 사례

*[Content to be added in full version]*

---

## Best Practices / 모범 사례

*[Content to be added in full version]*

---

## Performance Considerations / 성능 고려사항

*[Content to be added in full version]*

---

## Troubleshooting / 문제 해결

*[Content to be added in full version]*

---

## FAQ / 자주 묻는 질문

*[Content to be added in full version]*

---

## Additional Resources / 추가 자료

*[Content to be added in full version]*

---
