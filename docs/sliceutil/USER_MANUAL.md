# Sliceutil Package - User Manual / 사용자 매뉴얼

**Version / 버전**: v1.7.016
**Package / 패키지**: `github.com/arkd0ng/go-utils/sliceutil`
**Go Version / Go 버전**: 1.18+

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Function Reference / 함수 참조](#function-reference--함수-참조)
   - [Basic Operations / 기본 작업](#basic-operations--기본-작업)
   - [Transformation / 변환](#transformation--변환)
   - [Aggregation / 집계](#aggregation--집계)
   - [Slicing / 슬라이싱](#slicing--슬라이싱)
   - [Set Operations / 집합 작업](#set-operations--집합-작업)
   - [Sorting / 정렬](#sorting--정렬)
   - [Predicates / 조건자](#predicates--조건자)
   - [Utilities / 유틸리티](#utilities--유틸리티)
5. [Common Use Cases / 일반적인 사용 사례](#common-use-cases--일반적인-사용-사례)
6. [Best Practices / 모범 사례](#best-practices--모범-사례)
7. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)
8. [FAQ / 자주 묻는 질문](#faq--자주-묻는-질문)

---

## Introduction / 소개

### What is Sliceutil? / Sliceutil이란?

Sliceutil is a Go package that provides **extreme simplicity** for slice operations, reducing repetitive code from 20+ lines to just 1 line. It offers 60 type-safe functions across 8 categories for common slice manipulations.

Sliceutil은 슬라이스 작업을 위한 **극도의 간결함**을 제공하는 Go 패키지로, 반복적인 코드를 20줄 이상에서 단 1줄로 줄입니다. 일반적인 슬라이스 조작을 위해 8개 카테고리에 걸쳐 60개의 타입 안전 함수를 제공합니다.

### Design Philosophy / 설계 철학

**"20 lines → 1 line"** - Transform verbose, repetitive slice manipulation code into simple, readable function calls.

**"20줄 → 1줄"** - 장황하고 반복적인 슬라이스 조작 코드를 간단하고 읽기 쉬운 함수 호출로 변환합니다.

**Before (Standard Go) / 이전 (표준 Go)**:
```go
// Filter even numbers and double them
numbers := []int{1, 2, 3, 4, 5, 6}
var evens []int
for _, n := range numbers {
    if n%2 == 0 {
        evens = append(evens, n)
    }
}
var doubled []int
for _, n := range evens {
    doubled = append(doubled, n*2)
}
// 12+ lines of code
```

**After (With Sliceutil) / 이후 (Sliceutil 사용)**:
```go
numbers := []int{1, 2, 3, 4, 5, 6}
evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
doubled := sliceutil.Map(evens, func(n int) int { return n * 2 })
// 2 lines of code (vs 12+)
```

### Key Features / 주요 기능

- **Type-Safe with Go 1.18+ Generics** / **Go 1.18+ 제네릭으로 타입 안전**: Compile-time type checking eliminates runtime errors / 컴파일 타임 타입 검사로 런타임 오류 제거
- **Functional Programming Style** / **함수형 프로그래밍 스타일**: Map, Filter, Reduce, and more / Map, Filter, Reduce 등
- **Immutable Operations** / **불변 작업**: Original slices are never modified / 원본 슬라이스는 절대 수정되지 않음
- **Zero External Dependencies** / **제로 외부 의존성**: Only uses Go standard library / Go 표준 라이브러리만 사용
- **Comprehensive Coverage** / **포괄적인 커버리지**: 60 functions for all common operations / 모든 일반적인 작업을 위한 60개 함수
- **Performance Optimized** / **성능 최적화**: Efficient algorithms with minimal allocations / 최소 할당으로 효율적인 알고리즘

### Use Cases / 사용 사례

Sliceutil is perfect for:
- Data transformation and filtering / 데이터 변환 및 필터링
- Collection processing pipelines / 컬렉션 처리 파이프라인
- Business logic simplification / 비즈니스 로직 간소화
- Reducing boilerplate code / 보일러플레이트 코드 감소
- Functional programming in Go / Go에서의 함수형 프로그래밍

---

## Installation / 설치

### Prerequisites / 전제 조건

- **Go 1.18 or higher** (required for generics support) / **Go 1.18 이상** (제네릭 지원 필요)
- Basic understanding of Go slices / Go 슬라이스에 대한 기본 이해

### Install the Package / 패키지 설치

```bash
go get github.com/arkd0ng/go-utils/sliceutil
```

### Import in Your Code / 코드에서 임포트

```go
import "github.com/arkd0ng/go-utils/sliceutil"
```

### Verify Installation / 설치 확인

Create a simple test file / 간단한 테스트 파일 생성:

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    doubled := sliceutil.Map(numbers, func(n int) int { return n * 2 })
    fmt.Println(doubled) // Output: [2 4 6 8 10]
}
```

Run the test / 테스트 실행:
```bash
go run main.go
```

---

## Quick Start / 빠른 시작

Here are 5 examples showing the most common use cases / 가장 일반적인 사용 사례를 보여주는 5가지 예제:

### Example 1: Filter and Map / 필터링 및 매핑

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

func main() {
    // Filter even numbers and double them
    // 짝수를 필터링하고 2배로 만들기
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    evens := sliceutil.Filter(numbers, func(n int) bool {
        return n%2 == 0
    })

    doubled := sliceutil.Map(evens, func(n int) int {
        return n * 2
    })

    fmt.Println(doubled) // [4 8 12 16 20]
}
```

### Example 2: Find and Contains / 찾기 및 포함 확인

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

type User struct {
    ID   int
    Name string
    Age  int
}

func main() {
    users := []User{
        {ID: 1, Name: "Alice", Age: 25},
        {ID: 2, Name: "Bob", Age: 30},
        {ID: 3, Name: "Charlie", Age: 35},
    }

    // Find user with ID 2
    // ID가 2인 사용자 찾기
    user, found := sliceutil.Find(users, func(u User) bool {
        return u.ID == 2
    })

    if found {
        fmt.Printf("Found: %s (Age: %d)\n", user.Name, user.Age)
        // Output: Found: Bob (Age: 30)
    }

    // Check if any user is over 30
    // 30세 이상인 사용자가 있는지 확인
    hasOver30 := sliceutil.Any(users, func(u User) bool {
        return u.Age > 30
    })

    fmt.Println("Has user over 30:", hasOver30) // true
}
```

### Example 3: GroupBy and Aggregate / 그룹화 및 집계

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

type Product struct {
    Name     string
    Category string
    Price    float64
}

func main() {
    products := []Product{
        {Name: "Laptop", Category: "Electronics", Price: 999.99},
        {Name: "Mouse", Category: "Electronics", Price: 29.99},
        {Name: "Desk", Category: "Furniture", Price: 199.99},
        {Name: "Chair", Category: "Furniture", Price: 149.99},
    }

    // Group by category
    // 카테고리별로 그룹화
    grouped := sliceutil.GroupBy(products, func(p Product) string {
        return p.Category
    })

    // Calculate average price per category
    // 카테고리별 평균 가격 계산
    for category, items := range grouped {
        prices := sliceutil.Map(items, func(p Product) float64 {
            return p.Price
        })
        avg := sliceutil.Average(prices)
        fmt.Printf("%s: $%.2f\n", category, avg)
    }
    // Output:
    // Electronics: $514.99
    // Furniture: $174.99
}
```

### Example 4: Unique and Set Operations / 고유값 및 집합 작업

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

func main() {
    // Remove duplicates
    // 중복 제거
    numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
    unique := sliceutil.Unique(numbers)
    fmt.Println("Unique:", unique) // [1 2 3 4 5]

    // Set operations
    // 집합 작업
    set1 := []int{1, 2, 3, 4, 5}
    set2 := []int{4, 5, 6, 7, 8}

    union := sliceutil.Union(set1, set2)
    fmt.Println("Union:", union) // [1 2 3 4 5 6 7 8]

    intersection := sliceutil.Intersection(set1, set2)
    fmt.Println("Intersection:", intersection) // [4 5]

    difference := sliceutil.Difference(set1, set2)
    fmt.Println("Difference:", difference) // [1 2 3]
}
```

### Example 5: Chunk and Partition / 청크 및 파티션

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Split into chunks of 3
    // 3개씩 청크로 분할
    chunks := sliceutil.Chunk(numbers, 3)
    fmt.Println("Chunks:", chunks)
    // Output: [[1 2 3] [4 5 6] [7 8 9] [10]]

    // Partition by even/odd
    // 짝수/홀수로 파티션
    evens, odds := sliceutil.Partition(numbers, func(n int) bool {
        return n%2 == 0
    })

    fmt.Println("Evens:", evens) // [2 4 6 8 10]
    fmt.Println("Odds:", odds)   // [1 3 5 7 9]
}
```

---

## Function Reference / 함수 참조

### Basic Operations / 기본 작업

This category provides fundamental slice operations for searching, finding, and checking elements.

이 카테고리는 요소 검색, 찾기 및 확인을 위한 기본적인 슬라이스 작업을 제공합니다.

#### 1. Contains

Check if a slice contains a specific element / 슬라이스에 특정 요소가 포함되어 있는지 확인

**Signature / 시그니처**:
```go
func Contains[T comparable](slice []T, element T) bool
```

**Parameters / 매개변수**:
- `slice`: The slice to search / 검색할 슬라이스
- `element`: The element to find / 찾을 요소

**Returns / 반환값**: `true` if element is found, `false` otherwise / 요소가 발견되면 `true`, 그렇지 않으면 `false`

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}
found := sliceutil.Contains(numbers, 3) // true

names := []string{"Alice", "Bob", "Charlie"}
found = sliceutil.Contains(names, "David") // false
```

#### 2. IndexOf

Find the index of the first occurrence of an element / 요소의 첫 번째 발생 인덱스 찾기

**Signature / 시그니처**:
```go
func IndexOf[T comparable](slice []T, element T) int
```

**Parameters / 매개변수**:
- `slice`: The slice to search / 검색할 슬라이스
- `element`: The element to find / 찾을 요소

**Returns / 반환값**: Index of element, or -1 if not found / 요소의 인덱스, 또는 발견되지 않으면 -1

**Example / 예제**:
```go
numbers := []int{10, 20, 30, 40, 50}
index := sliceutil.IndexOf(numbers, 30) // 2

names := []string{"Alice", "Bob", "Charlie"}
index = sliceutil.IndexOf(names, "Bob") // 1
index = sliceutil.IndexOf(names, "David") // -1
```

#### 3. LastIndexOf

Find the index of the last occurrence of an element / 요소의 마지막 발생 인덱스 찾기

**Signature / 시그니처**:
```go
func LastIndexOf[T comparable](slice []T, element T) int
```

**Parameters / 매개변수**:
- `slice`: The slice to search / 검색할 슬라이스
- `element`: The element to find / 찾을 요소

**Returns / 반환값**: Last index of element, or -1 if not found / 요소의 마지막 인덱스, 또는 발견되지 않으면 -1

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 2, 1}
index := sliceutil.LastIndexOf(numbers, 2) // 3 (not 1)

letters := []string{"a", "b", "c", "b", "a"}
index = sliceutil.LastIndexOf(letters, "b") // 3
```

#### 4. Find

Find the first element matching a predicate / 조건자와 일치하는 첫 번째 요소 찾기

**Signature / 시그니처**:
```go
func Find[T any](slice []T, predicate func(T) bool) (T, bool)
```

**Parameters / 매개변수**:
- `slice`: The slice to search / 검색할 슬라이스
- `predicate`: Function that returns true for the desired element / 원하는 요소에 대해 true를 반환하는 함수

**Returns / 반환값**: The element and `true` if found, zero value and `false` otherwise / 발견되면 요소와 `true`, 그렇지 않으면 제로값과 `false`

**Example / 예제**:
```go
type User struct {
    ID   int
    Name string
    Age  int
}

users := []User{
    {ID: 1, Name: "Alice", Age: 25},
    {ID: 2, Name: "Bob", Age: 30},
}

user, found := sliceutil.Find(users, func(u User) bool {
    return u.Age > 28
})

if found {
    fmt.Println(user.Name) // "Bob"
}
```

#### 5. FindIndex

Find the index of the first element matching a predicate / 조건자와 일치하는 첫 번째 요소의 인덱스 찾기

**Signature / 시그니처**:
```go
func FindIndex[T any](slice []T, predicate func(T) bool) int
```

**Parameters / 매개변수**:
- `slice`: The slice to search / 검색할 슬라이스
- `predicate`: Function that returns true for the desired element / 원하는 요소에 대해 true를 반환하는 함수

**Returns / 반환값**: Index of the element, or -1 if not found / 요소의 인덱스, 또는 발견되지 않으면 -1

**Example / 예제**:
```go
numbers := []int{1, 3, 5, 8, 10}
index := sliceutil.FindIndex(numbers, func(n int) bool {
    return n%2 == 0 // First even number
})
fmt.Println(index) // 3 (value is 8)
```

#### 6. FindLast

Find the last element matching a predicate / 조건자와 일치하는 마지막 요소 찾기

**Signature / 시그니처**:
```go
func FindLast[T any](slice []T, predicate func(T) bool) (T, bool)
```

**Parameters / 매개변수**:
- `slice`: The slice to search / 검색할 슬라이스
- `predicate`: Function that returns true for the desired element / 원하는 요소에 대해 true를 반환하는 함수

**Returns / 반환값**: The last matching element and `true` if found, zero value and `false` otherwise / 발견되면 마지막 일치 요소와 `true`, 그렇지 않으면 제로값과 `false`

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5, 6}
last, found := sliceutil.FindLast(numbers, func(n int) bool {
    return n%2 == 0 // Last even number
})

if found {
    fmt.Println(last) // 6
}
```

#### 7. FindLastIndex

Find the index of the last element matching a predicate / 조건자와 일치하는 마지막 요소의 인덱스 찾기

**Signature / 시그니처**:
```go
func FindLastIndex[T any](slice []T, predicate func(T) bool) int
```

**Parameters / 매개변수**:
- `slice`: The slice to search / 검색할 슬라이스
- `predicate`: Function that returns true for the desired element / 원하는 요소에 대해 true를 반환하는 함수

**Returns / 반환값**: Last index of matching element, or -1 if not found / 일치하는 요소의 마지막 인덱스, 또는 발견되지 않으면 -1

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5, 6}
index := sliceutil.FindLastIndex(numbers, func(n int) bool {
    return n < 4
})
fmt.Println(index) // 2 (value is 3)
```

#### 8. Count

Count elements matching a predicate / 조건자와 일치하는 요소 개수 세기

**Signature / 시그니처**:
```go
func Count[T any](slice []T, predicate func(T) bool) int
```

**Parameters / 매개변수**:
- `slice`: The slice to count from / 개수를 셀 슬라이스
- `predicate`: Function that returns true for elements to count / 세려는 요소에 대해 true를 반환하는 함수

**Returns / 반환값**: Number of matching elements / 일치하는 요소의 개수

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evenCount := sliceutil.Count(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(evenCount) // 5
```

#### 9. First

Get the first element of a slice / 슬라이스의 첫 번째 요소 가져오기

**Signature / 시그니처**:
```go
func First[T any](slice []T) (T, bool)
```

**Parameters / 매개변수**:
- `slice`: The slice / 슬라이스

**Returns / 반환값**: The first element and `true` if slice is not empty, zero value and `false` otherwise / 슬라이스가 비어있지 않으면 첫 번째 요소와 `true`, 그렇지 않으면 제로값과 `false`

**Example / 예제**:
```go
numbers := []int{10, 20, 30}
first, ok := sliceutil.First(numbers)
if ok {
    fmt.Println(first) // 10
}

empty := []int{}
_, ok = sliceutil.First(empty) // ok is false
```

#### 10. Last

Get the last element of a slice / 슬라이스의 마지막 요소 가져오기

**Signature / 시그니처**:
```go
func Last[T any](slice []T) (T, bool)
```

**Parameters / 매개변수**:
- `slice`: The slice / 슬라이스

**Returns / 반환값**: The last element and `true` if slice is not empty, zero value and `false` otherwise / 슬라이스가 비어있지 않으면 마지막 요소와 `true`, 그렇지 않으면 제로값과 `false`

**Example / 예제**:
```go
numbers := []int{10, 20, 30}
last, ok := sliceutil.Last(numbers)
if ok {
    fmt.Println(last) // 30
}
```

---

### Transformation / 변환

This category provides functions for transforming slices into different forms.

이 카테고리는 슬라이스를 다양한 형태로 변환하는 함수를 제공합니다.

#### 1. Map

Transform each element using a mapper function / 매퍼 함수를 사용하여 각 요소 변환

**Signature / 시그니처**:
```go
func Map[T any, R any](slice []T, mapper func(T) R) []R
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `mapper`: Function to transform each element / 각 요소를 변환하는 함수

**Returns / 반환값**: New slice with transformed elements / 변환된 요소로 구성된 새 슬라이스

**Example / 예제**:
```go
// Double all numbers
numbers := []int{1, 2, 3, 4, 5}
doubled := sliceutil.Map(numbers, func(n int) int {
    return n * 2
})
fmt.Println(doubled) // [2 4 6 8 10]

// Convert to strings
strings := sliceutil.Map(numbers, func(n int) string {
    return fmt.Sprintf("Number: %d", n)
})
// ["Number: 1", "Number: 2", ...]
```

#### 2. Filter

Keep only elements matching a predicate / 조건자와 일치하는 요소만 유지

**Signature / 시그니처**:
```go
func Filter[T any](slice []T, predicate func(T) bool) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `predicate`: Function that returns true for elements to keep / 유지할 요소에 대해 true를 반환하는 함수

**Returns / 반환값**: New slice with only matching elements / 일치하는 요소만 포함된 새 슬라이스

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5, 6}
evens := sliceutil.Filter(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(evens) // [2 4 6]
```

#### 3. Unique

Remove duplicate elements / 중복 요소 제거

**Signature / 시그니처**:
```go
func Unique[T comparable](slice []T) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: New slice with unique elements only / 고유 요소만 포함된 새 슬라이스

**Example / 예제**:
```go
numbers := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
unique := sliceutil.Unique(numbers)
fmt.Println(unique) // [1 2 3 4 5]
```

#### 4. UniqueBy

Remove duplicates based on a key function / 키 함수를 기반으로 중복 제거

**Signature / 시그니처**:
```go
func UniqueBy[T any, K comparable](slice []T, keyFunc func(T) K) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `keyFunc`: Function to extract comparison key from element / 요소에서 비교 키를 추출하는 함수

**Returns / 반환값**: New slice with unique elements based on key / 키를 기반으로 고유한 요소로 구성된 새 슬라이스

**Example / 예제**:
```go
type User struct {
    ID   int
    Name string
}

users := []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 1, Name: "Alice Duplicate"},
}

unique := sliceutil.UniqueBy(users, func(u User) int {
    return u.ID
})
// Result: [{1 Alice} {2 Bob}]
```

#### 5. Reverse

Reverse the order of elements / 요소의 순서 뒤집기

**Signature / 시그니처**:
```go
func Reverse[T any](slice []T) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: New slice with elements in reverse order / 요소가 역순으로 배열된 새 슬라이스

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}
reversed := sliceutil.Reverse(numbers)
fmt.Println(reversed) // [5 4 3 2 1]
```

#### 6. Flatten

Flatten a slice of slices into a single slice / 슬라이스의 슬라이스를 단일 슬라이스로 평탄화

**Signature / 시그니처**:
```go
func Flatten[T any](slice [][]T) []T
```

**Parameters / 매개변수**:
- `slice`: Slice of slices / 슬라이스의 슬라이스

**Returns / 반환값**: Single flattened slice / 평탄화된 단일 슬라이스

**Example / 예제**:
```go
nested := [][]int{
    {1, 2, 3},
    {4, 5},
    {6, 7, 8, 9},
}
flat := sliceutil.Flatten(nested)
fmt.Println(flat) // [1 2 3 4 5 6 7 8 9]
```

#### 7. FlatMap

Map each element to a slice and flatten the result / 각 요소를 슬라이스로 매핑하고 결과 평탄화

**Signature / 시그니처**:
```go
func FlatMap[T any, R any](slice []T, mapper func(T) []R) []R
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `mapper`: Function that maps each element to a slice / 각 요소를 슬라이스로 매핑하는 함수

**Returns / 반환값**: Flattened slice of mapped results / 매핑된 결과의 평탄화된 슬라이스

**Example / 예제**:
```go
words := []string{"hello", "world"}
chars := sliceutil.FlatMap(words, func(w string) []rune {
    return []rune(w)
})
// Result: ['h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd']
```

#### 8. Compact

Remove zero values from a slice / 슬라이스에서 제로값 제거

**Signature / 시그니처**:
```go
func Compact[T comparable](slice []T) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: New slice with zero values removed / 제로값이 제거된 새 슬라이스

**Example / 예제**:
```go
numbers := []int{1, 0, 2, 0, 3, 0, 4}
compacted := sliceutil.Compact(numbers)
fmt.Println(compacted) // [1 2 3 4]

strings := []string{"a", "", "b", "", "c"}
compacted2 := sliceutil.Compact(strings)
fmt.Println(compacted2) // ["a" "b" "c"]
```

---

### Aggregation / 집계

This category provides functions for aggregating slice data.

이 카테고리는 슬라이스 데이터를 집계하는 함수를 제공합니다.

#### 1. Reduce

Reduce a slice to a single value / 슬라이스를 단일 값으로 축소

**Signature / 시그니처**:
```go
func Reduce[T any, R any](slice []T, initial R, reducer func(R, T) R) R
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `initial`: Initial accumulator value / 초기 누산기 값
- `reducer`: Function to combine accumulator and element / 누산기와 요소를 결합하는 함수

**Returns / 반환값**: Final accumulated value / 최종 누산된 값

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}
sum := sliceutil.Reduce(numbers, 0, func(acc, n int) int {
    return acc + n
})
fmt.Println(sum) // 15

// String concatenation
words := []string{"Hello", "World", "!"}
sentence := sliceutil.Reduce(words, "", func(acc, w string) string {
    if acc == "" {
        return w
    }
    return acc + " " + w
})
fmt.Println(sentence) // "Hello World !"
```

#### 2. Sum

Calculate the sum of numeric elements / 숫자 요소의 합 계산

**Signature / 시그니처**:
```go
func Sum[T Number](slice []T) T
```

**Parameters / 매개변수**:
- `slice`: Slice of numbers / 숫자 슬라이스

**Returns / 반환값**: Sum of all elements / 모든 요소의 합

**Example / 예제**:
```go
integers := []int{1, 2, 3, 4, 5}
total := sliceutil.Sum(integers)
fmt.Println(total) // 15

prices := []float64{9.99, 19.99, 29.99}
totalPrice := sliceutil.Sum(prices)
fmt.Println(totalPrice) // 59.97
```

#### 3. Average

Calculate the average of numeric elements / 숫자 요소의 평균 계산

**Signature / 시그니처**:
```go
func Average[T Number](slice []T) float64
```

**Parameters / 매개변수**:
- `slice`: Slice of numbers / 숫자 슬라이스

**Returns / 반환값**: Average value, or 0 if slice is empty / 평균 값, 또는 슬라이스가 비어있으면 0

**Example / 예제**:
```go
numbers := []int{10, 20, 30, 40, 50}
avg := sliceutil.Average(numbers)
fmt.Println(avg) // 30.0

scores := []float64{85.5, 90.0, 78.5, 92.0}
avgScore := sliceutil.Average(scores)
fmt.Println(avgScore) // 86.5
```

#### 4. Min

Find the minimum value / 최소값 찾기

**Signature / 시그니처**:
```go
func Min[T Number](slice []T) (T, bool)
```

**Parameters / 매개변수**:
- `slice`: Slice of numbers / 숫자 슬라이스

**Returns / 반환값**: Minimum value and `true`, or zero value and `false` if empty / 최소값과 `true`, 또는 비어있으면 제로값과 `false`

**Example / 예제**:
```go
numbers := []int{5, 2, 8, 1, 9}
min, ok := sliceutil.Min(numbers)
if ok {
    fmt.Println(min) // 1
}
```

#### 5. Max

Find the maximum value / 최대값 찾기

**Signature / 시그니처**:
```go
func Max[T Number](slice []T) (T, bool)
```

**Parameters / 매개변수**:
- `slice`: Slice of numbers / 숫자 슬라이스

**Returns / 반환값**: Maximum value and `true`, or zero value and `false` if empty / 최대값과 `true`, 또는 비어있으면 제로값과 `false`

**Example / 예제**:
```go
numbers := []int{5, 2, 8, 1, 9}
max, ok := sliceutil.Max(numbers)
if ok {
    fmt.Println(max) // 9
}
```

#### 6. GroupBy

Group elements by a key function / 키 함수로 요소 그룹화

**Signature / 시그니처**:
```go
func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `keyFunc`: Function to extract grouping key / 그룹화 키를 추출하는 함수

**Returns / 반환값**: Map of keys to slices of elements / 키를 슬라이스 요소에 매핑하는 맵

**Example / 예제**:
```go
type Person struct {
    Name string
    Age  int
    City string
}

people := []Person{
    {Name: "Alice", Age: 25, City: "Seoul"},
    {Name: "Bob", Age: 30, City: "Busan"},
    {Name: "Charlie", Age: 25, City: "Seoul"},
}

byCity := sliceutil.GroupBy(people, func(p Person) string {
    return p.City
})
// Result: map[Seoul:[{Alice 25 Seoul} {Charlie 25 Seoul}] Busan:[{Bob 30 Busan}]]
```

#### 7. CountBy

Count elements by a key function / 키 함수로 요소 개수 세기

**Signature / 시그니처**:
```go
func CountBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K]int
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `keyFunc`: Function to extract counting key / 카운팅 키를 추출하는 함수

**Returns / 반환값**: Map of keys to counts / 키를 개수에 매핑하는 맵

**Example / 예제**:
```go
words := []string{"apple", "banana", "apricot", "blueberry", "avocado"}
byFirstLetter := sliceutil.CountBy(words, func(w string) rune {
    return rune(w[0])
})
// Result: map['a':3 'b':2]
```

---

### Slicing / 슬라이싱

This category provides functions for extracting portions of slices.

이 카테고리는 슬라이스의 일부를 추출하는 함수를 제공합니다.

#### 1. Chunk

Split a slice into chunks of specified size / 슬라이스를 지정된 크기의 청크로 분할

**Signature / 시그니처**:
```go
func Chunk[T any](slice []T, size int) [][]T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `size`: Size of each chunk / 각 청크의 크기

**Returns / 반환값**: Slice of chunks / 청크의 슬라이스

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
chunks := sliceutil.Chunk(numbers, 3)
// Result: [[1 2 3] [4 5 6] [7 8 9]]

chunks2 := sliceutil.Chunk(numbers, 4)
// Result: [[1 2 3 4] [5 6 7 8] [9]]
```

#### 2. Take

Take the first n elements / 처음 n개 요소 가져오기

**Signature / 시그니처**:
```go
func Take[T any](slice []T, n int) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `n`: Number of elements to take / 가져올 요소 개수

**Returns / 반환값**: New slice with first n elements / 처음 n개 요소로 구성된 새 슬라이스

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}
first3 := sliceutil.Take(numbers, 3)
fmt.Println(first3) // [1 2 3]
```

#### 3. TakeWhile

Take elements while predicate is true / 조건자가 true인 동안 요소 가져오기

**Signature / 시그니처**:
```go
func TakeWhile[T any](slice []T, predicate func(T) bool) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `predicate`: Function to test elements / 요소를 테스트하는 함수

**Returns / 반환값**: Elements from start while predicate is true / 조건자가 true인 동안 시작부터의 요소

**Example / 예제**:
```go
numbers := []int{2, 4, 6, 7, 8, 10}
evens := sliceutil.TakeWhile(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(evens) // [2 4 6]
```

#### 4. Drop

Drop the first n elements / 처음 n개 요소 제거

**Signature / 시그니처**:
```go
func Drop[T any](slice []T, n int) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `n`: Number of elements to drop / 제거할 요소 개수

**Returns / 반환값**: New slice without first n elements / 처음 n개 요소가 제거된 새 슬라이스

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}
rest := sliceutil.Drop(numbers, 2)
fmt.Println(rest) // [3 4 5]
```

#### 5. DropWhile

Drop elements while predicate is true / 조건자가 true인 동안 요소 제거

**Signature / 시그니처**:
```go
func DropWhile[T any](slice []T, predicate func(T) bool) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `predicate`: Function to test elements / 요소를 테스트하는 함수

**Returns / 반환값**: Elements from first false predicate onward / 첫 번째 false 조건자부터의 요소

**Example / 예제**:
```go
numbers := []int{2, 4, 6, 7, 8, 10}
afterOdd := sliceutil.DropWhile(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(afterOdd) // [7 8 10]
```

#### 6. Partition

Partition elements into two slices based on predicate / 조건자를 기반으로 두 슬라이스로 요소 분할

**Signature / 시그니처**:
```go
func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T)
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `predicate`: Function to test elements / 요소를 테스트하는 함수

**Returns / 반환값**: Two slices: matching and non-matching / 두 슬라이스: 일치하는 것과 일치하지 않는 것

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
evens, odds := sliceutil.Partition(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(evens) // [2 4 6 8]
fmt.Println(odds)  // [1 3 5 7]
```

#### 7. Sample

Get n random elements from a slice / 슬라이스에서 n개의 무작위 요소 가져오기

**Signature / 시그니처**:
```go
func Sample[T any](slice []T, n int) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `n`: Number of elements to sample / 샘플링할 요소 개수

**Returns / 반환값**: Random sample of n elements / n개 요소의 무작위 샘플

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sample := sliceutil.Sample(numbers, 3)
fmt.Println(sample) // e.g., [3 7 1] (random)
```

---

### Set Operations / 집합 작업

This category provides set-theory operations on slices.

이 카테고리는 슬라이스에 대한 집합론 작업을 제공합니다.

#### 1. Union

Get the union of two slices (all unique elements) / 두 슬라이스의 합집합 가져오기 (모든 고유 요소)

**Signature / 시그니처**:
```go
func Union[T comparable](slice1, slice2 []T) []T
```

**Parameters / 매개변수**:
- `slice1`: First slice / 첫 번째 슬라이스
- `slice2`: Second slice / 두 번째 슬라이스

**Returns / 반환값**: Union of both slices / 두 슬라이스의 합집합

**Example / 예제**:
```go
set1 := []int{1, 2, 3, 4}
set2 := []int{3, 4, 5, 6}
union := sliceutil.Union(set1, set2)
fmt.Println(union) // [1 2 3 4 5 6]
```

#### 2. Intersection

Get the intersection of two slices (common elements) / 두 슬라이스의 교집합 가져오기 (공통 요소)

**Signature / 시그니처**:
```go
func Intersection[T comparable](slice1, slice2 []T) []T
```

**Parameters / 매개변수**:
- `slice1`: First slice / 첫 번째 슬라이스
- `slice2`: Second slice / 두 번째 슬라이스

**Returns / 반환값**: Intersection of both slices / 두 슬라이스의 교집합

**Example / 예제**:
```go
set1 := []int{1, 2, 3, 4}
set2 := []int{3, 4, 5, 6}
intersection := sliceutil.Intersection(set1, set2)
fmt.Println(intersection) // [3 4]
```

#### 3. Difference

Get elements in first slice but not in second / 첫 번째 슬라이스에는 있지만 두 번째에는 없는 요소 가져오기

**Signature / 시그니처**:
```go
func Difference[T comparable](slice1, slice2 []T) []T
```

**Parameters / 매개변수**:
- `slice1`: First slice / 첫 번째 슬라이스
- `slice2`: Second slice / 두 번째 슬라이스

**Returns / 반환값**: Elements in slice1 but not in slice2 / slice1에는 있지만 slice2에는 없는 요소

**Example / 예제**:
```go
set1 := []int{1, 2, 3, 4}
set2 := []int{3, 4, 5, 6}
diff := sliceutil.Difference(set1, set2)
fmt.Println(diff) // [1 2]
```

#### 4. SymmetricDifference

Get elements in either slice but not in both / 어느 한 슬라이스에는 있지만 둘 다에는 없는 요소 가져오기

**Signature / 시그니처**:
```go
func SymmetricDifference[T comparable](slice1, slice2 []T) []T
```

**Parameters / 매개변수**:
- `slice1`: First slice / 첫 번째 슬라이스
- `slice2`: Second slice / 두 번째 슬라이스

**Returns / 반환값**: Elements in either but not both / 어느 한 쪽에는 있지만 둘 다에는 없는 요소

**Example / 예제**:
```go
set1 := []int{1, 2, 3, 4}
set2 := []int{3, 4, 5, 6}
symDiff := sliceutil.SymmetricDifference(set1, set2)
fmt.Println(symDiff) // [1 2 5 6]
```

#### 5. IsSubset

Check if first slice is a subset of second / 첫 번째 슬라이스가 두 번째의 부분집합인지 확인

**Signature / 시그니처**:
```go
func IsSubset[T comparable](subset, superset []T) bool
```

**Parameters / 매개변수**:
- `subset`: Potential subset / 잠재적 부분집합
- `superset`: Potential superset / 잠재적 상위집합

**Returns / 반환값**: `true` if subset is a subset of superset / subset이 superset의 부분집합이면 `true`

**Example / 예제**:
```go
subset := []int{1, 2}
superset := []int{1, 2, 3, 4, 5}
isSubset := sliceutil.IsSubset(subset, superset)
fmt.Println(isSubset) // true
```

#### 6. IsDisjoint

Check if two slices have no common elements / 두 슬라이스에 공통 요소가 없는지 확인

**Signature / 시그니처**:
```go
func IsDisjoint[T comparable](slice1, slice2 []T) bool
```

**Parameters / 매개변수**:
- `slice1`: First slice / 첫 번째 슬라이스
- `slice2`: Second slice / 두 번째 슬라이스

**Returns / 반환값**: `true` if slices have no common elements / 슬라이스에 공통 요소가 없으면 `true`

**Example / 예제**:
```go
set1 := []int{1, 2, 3}
set2 := []int{4, 5, 6}
disjoint := sliceutil.IsDisjoint(set1, set2)
fmt.Println(disjoint) // true

set3 := []int{3, 4, 5}
disjoint2 := sliceutil.IsDisjoint(set1, set3)
fmt.Println(disjoint2) // false (3 is common)
```

---

### Sorting / 정렬

This category provides sorting-related functions.

이 카테고리는 정렬 관련 함수를 제공합니다.

#### 1. Sort

Sort a slice in ascending order / 슬라이스를 오름차순으로 정렬

**Signature / 시그니처**:
```go
func Sort[T Ordered](slice []T) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: New sorted slice / 정렬된 새 슬라이스

**Example / 예제**:
```go
numbers := []int{5, 2, 8, 1, 9}
sorted := sliceutil.Sort(numbers)
fmt.Println(sorted) // [1 2 5 8 9]

words := []string{"banana", "apple", "cherry"}
sortedWords := sliceutil.Sort(words)
fmt.Println(sortedWords) // ["apple" "banana" "cherry"]
```

#### 2. SortBy

Sort a slice using a custom comparison function / 사용자 정의 비교 함수를 사용하여 슬라이스 정렬

**Signature / 시그니처**:
```go
func SortBy[T any](slice []T, less func(a, b T) bool) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `less`: Function that returns true if a < b / a < b이면 true를 반환하는 함수

**Returns / 반환값**: New sorted slice / 정렬된 새 슬라이스

**Example / 예제**:
```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
}

// Sort by age
byAge := sliceutil.SortBy(people, func(a, b Person) bool {
    return a.Age < b.Age
})
// Result: [{Bob 25} {Alice 30} {Charlie 35}]
```

#### 3. SortDesc

Sort a slice in descending order / 슬라이스를 내림차순으로 정렬

**Signature / 시그니처**:
```go
func SortDesc[T Ordered](slice []T) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: New sorted slice in descending order / 내림차순으로 정렬된 새 슬라이스

**Example / 예제**:
```go
numbers := []int{5, 2, 8, 1, 9}
sorted := sliceutil.SortDesc(numbers)
fmt.Println(sorted) // [9 8 5 2 1]
```

#### 4. IsSorted

Check if a slice is sorted in ascending order / 슬라이스가 오름차순으로 정렬되어 있는지 확인

**Signature / 시그니처**:
```go
func IsSorted[T Ordered](slice []T) bool
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: `true` if sorted, `false` otherwise / 정렬되어 있으면 `true`, 그렇지 않으면 `false`

**Example / 예제**:
```go
sorted := []int{1, 2, 3, 4, 5}
fmt.Println(sliceutil.IsSorted(sorted)) // true

unsorted := []int{1, 3, 2, 4}
fmt.Println(sliceutil.IsSorted(unsorted)) // false
```

#### 5. IsSortedBy

Check if a slice is sorted according to a comparison function / 비교 함수에 따라 슬라이스가 정렬되어 있는지 확인

**Signature / 시그니처**:
```go
func IsSortedBy[T any](slice []T, less func(a, b T) bool) bool
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `less`: Function that returns true if a < b / a < b이면 true를 반환하는 함수

**Returns / 반환값**: `true` if sorted according to less function / less 함수에 따라 정렬되어 있으면 `true`

**Example / 예제**:
```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {Name: "Bob", Age: 25},
    {Name: "Alice", Age: 30},
    {Name: "Charlie", Age: 35},
}

isSorted := sliceutil.IsSortedBy(people, func(a, b Person) bool {
    return a.Age < b.Age
})
fmt.Println(isSorted) // true
```

---

### Predicates / 조건자

This category provides predicate functions for testing slice conditions.

이 카테고리는 슬라이스 조건을 테스트하는 조건자 함수를 제공합니다.

#### 1. All

Check if all elements satisfy a predicate / 모든 요소가 조건자를 만족하는지 확인

**Signature / 시그니처**:
```go
func All[T any](slice []T, predicate func(T) bool) bool
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `predicate`: Function to test elements / 요소를 테스트하는 함수

**Returns / 반환값**: `true` if all elements match, `false` otherwise / 모든 요소가 일치하면 `true`, 그렇지 않으면 `false`

**Example / 예제**:
```go
numbers := []int{2, 4, 6, 8}
allEven := sliceutil.All(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(allEven) // true
```

#### 2. Any

Check if any element satisfies a predicate / 어떤 요소라도 조건자를 만족하는지 확인

**Signature / 시그니처**:
```go
func Any[T any](slice []T, predicate func(T) bool) bool
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `predicate`: Function to test elements / 요소를 테스트하는 함수

**Returns / 반환값**: `true` if any element matches, `false` otherwise / 어떤 요소라도 일치하면 `true`, 그렇지 않으면 `false`

**Example / 예제**:
```go
numbers := []int{1, 3, 5, 8}
hasEven := sliceutil.Any(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(hasEven) // true (8 is even)
```

#### 3. None

Check if no elements satisfy a predicate / 모든 요소가 조건자를 만족하지 않는지 확인

**Signature / 시그니처**:
```go
func None[T any](slice []T, predicate func(T) bool) bool
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `predicate`: Function to test elements / 요소를 테스트하는 함수

**Returns / 반환값**: `true` if no elements match, `false` otherwise / 어떤 요소도 일치하지 않으면 `true`, 그렇지 않으면 `false`

**Example / 예제**:
```go
numbers := []int{1, 3, 5, 7}
noEvens := sliceutil.None(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(noEvens) // true
```

#### 4. AllEqual

Check if all elements are equal / 모든 요소가 동일한지 확인

**Signature / 시그니처**:
```go
func AllEqual[T comparable](slice []T) bool
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: `true` if all elements are equal, `false` otherwise / 모든 요소가 동일하면 `true`, 그렇지 않으면 `false`

**Example / 예제**:
```go
same := []int{5, 5, 5, 5}
fmt.Println(sliceutil.AllEqual(same)) // true

different := []int{5, 5, 5, 6}
fmt.Println(sliceutil.AllEqual(different)) // false
```

#### 5. IsEmpty

Check if a slice is empty / 슬라이스가 비어있는지 확인

**Signature / 시그니처**:
```go
func IsEmpty[T any](slice []T) bool
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: `true` if slice is empty, `false` otherwise / 슬라이스가 비어있으면 `true`, 그렇지 않으면 `false`

**Example / 예제**:
```go
empty := []int{}
fmt.Println(sliceutil.IsEmpty(empty)) // true

notEmpty := []int{1, 2, 3}
fmt.Println(sliceutil.IsEmpty(notEmpty)) // false
```

#### 6. Equal

Check if two slices are equal / 두 슬라이스가 동일한지 확인

**Signature / 시그니처**:
```go
func Equal[T comparable](slice1, slice2 []T) bool
```

**Parameters / 매개변수**:
- `slice1`: First slice / 첫 번째 슬라이스
- `slice2`: Second slice / 두 번째 슬라이스

**Returns / 반환값**: `true` if slices are equal, `false` otherwise / 슬라이스가 동일하면 `true`, 그렇지 않으면 `false`

**Example / 예제**:
```go
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
slice3 := []int{1, 2, 4}

fmt.Println(sliceutil.Equal(slice1, slice2)) // true
fmt.Println(sliceutil.Equal(slice1, slice3)) // false
```

---

### Utilities / 유틸리티

This category provides miscellaneous utility functions.

이 카테고리는 기타 유틸리티 함수를 제공합니다.

#### 1. ForEach

Execute a function for each element / 각 요소에 대해 함수 실행

**Signature / 시그니처**:
```go
func ForEach[T any](slice []T, fn func(T))
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `fn`: Function to execute for each element / 각 요소에 대해 실행할 함수

**Returns / 반환값**: None / 없음

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}
sliceutil.ForEach(numbers, func(n int) {
    fmt.Printf("%d ", n)
})
// Output: 1 2 3 4 5
```

#### 2. ForEachIndexed

Execute a function for each element with index / 인덱스와 함께 각 요소에 대해 함수 실행

**Signature / 시그니처**:
```go
func ForEachIndexed[T any](slice []T, fn func(int, T))
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `fn`: Function to execute with index and element / 인덱스와 요소로 실행할 함수

**Returns / 반환값**: None / 없음

**Example / 예제**:
```go
words := []string{"Hello", "World", "!"}
sliceutil.ForEachIndexed(words, func(i int, w string) {
    fmt.Printf("%d: %s\n", i, w)
})
// Output:
// 0: Hello
// 1: World
// 2: !
```

#### 3. Join

Join slice elements into a string / 슬라이스 요소를 문자열로 결합

**Signature / 시그니처**:
```go
func Join[T any](slice []T, separator string) string
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `separator`: Separator between elements / 요소 사이의 구분자

**Returns / 반환값**: Joined string / 결합된 문자열

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}
result := sliceutil.Join(numbers, ", ")
fmt.Println(result) // "1, 2, 3, 4, 5"

words := []string{"Hello", "World"}
sentence := sliceutil.Join(words, " ")
fmt.Println(sentence) // "Hello World"
```

#### 4. Clone

Create a shallow copy of a slice / 슬라이스의 얕은 복사본 생성

**Signature / 시그니처**:
```go
func Clone[T any](slice []T) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: New slice with copied elements / 복사된 요소로 구성된 새 슬라이스

**Example / 예제**:
```go
original := []int{1, 2, 3}
copy := sliceutil.Clone(original)

copy[0] = 99
fmt.Println(original) // [1 2 3] (unchanged)
fmt.Println(copy)     // [99 2 3]
```

#### 5. Fill

Fill a slice with a value / 슬라이스를 값으로 채우기

**Signature / 시그니처**:
```go
func Fill[T any](slice []T, value T) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `value`: Value to fill with / 채울 값

**Returns / 반환값**: New slice filled with value / 값으로 채워진 새 슬라이스

**Example / 예제**:
```go
numbers := make([]int, 5)
filled := sliceutil.Fill(numbers, 7)
fmt.Println(filled) // [7 7 7 7 7]
```

#### 6. Repeat

Create a slice with a value repeated n times / 값이 n번 반복된 슬라이스 생성

**Signature / 시그니처**:
```go
func Repeat[T any](value T, count int) []T
```

**Parameters / 매개변수**:
- `value`: Value to repeat / 반복할 값
- `count`: Number of repetitions / 반복 횟수

**Returns / 반환값**: New slice with repeated value / 반복된 값으로 구성된 새 슬라이스

**Example / 예제**:
```go
repeated := sliceutil.Repeat("Go", 3)
fmt.Println(repeated) // ["Go" "Go" "Go"]

zeros := sliceutil.Repeat(0, 5)
fmt.Println(zeros) // [0 0 0 0 0]
```

#### 7. Shuffle

Randomly shuffle a slice / 슬라이스를 무작위로 섞기

**Signature / 시그니처**:
```go
func Shuffle[T any](slice []T) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스

**Returns / 반환값**: New shuffled slice / 섞인 새 슬라이스

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}
shuffled := sliceutil.Shuffle(numbers)
fmt.Println(shuffled) // e.g., [3 1 5 2 4] (random order)
```

#### 8. Zip

Combine two slices into pairs / 두 슬라이스를 쌍으로 결합

**Signature / 시그니처**:
```go
func Zip[T, U any](slice1 []T, slice2 []U) [][2]interface{}
```

**Parameters / 매개변수**:
- `slice1`: First slice / 첫 번째 슬라이스
- `slice2`: Second slice / 두 번째 슬라이스

**Returns / 반환값**: Slice of pairs / 쌍의 슬라이스

**Example / 예제**:
```go
names := []string{"Alice", "Bob", "Charlie"}
ages := []int{25, 30, 35}

pairs := sliceutil.Zip(names, ages)
// Result: [{"Alice", 25}, {"Bob", 30}, {"Charlie", 35}]
```

#### 9. Unzip

Split pairs into two slices / 쌍을 두 슬라이스로 분할

**Signature / 시그니처**:
```go
func Unzip[T, U any](pairs [][2]interface{}) ([]T, []U)
```

**Parameters / 매개변수**:
- `pairs`: Slice of pairs / 쌍의 슬라이스

**Returns / 반환값**: Two slices / 두 슬라이스

**Example / 예제**:
```go
pairs := [][2]interface{}{
    {"Alice", 25},
    {"Bob", 30},
}

names, ages := sliceutil.Unzip[string, int](pairs)
// names: ["Alice", "Bob"]
// ages: [25, 30]
```

#### 10. Window

Create sliding windows of specified size / 지정된 크기의 슬라이딩 윈도우 생성

**Signature / 시그니처**:
```go
func Window[T any](slice []T, size int) [][]T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `size`: Window size / 윈도우 크기

**Returns / 반환값**: Slice of windows / 윈도우의 슬라이스

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}
windows := sliceutil.Window(numbers, 3)
// Result: [[1 2 3] [2 3 4] [3 4 5]]
```

#### 11. Tap

Execute a function on each element and return the slice / 각 요소에 함수를 실행하고 슬라이스 반환

**Signature / 시그니처**:
```go
func Tap[T any](slice []T, fn func(T)) []T
```

**Parameters / 매개변수**:
- `slice`: The input slice / 입력 슬라이스
- `fn`: Function to execute for each element / 각 요소에 대해 실행할 함수

**Returns / 반환값**: The original slice (for chaining) / 원본 슬라이스 (체이닝용)

**Example / 예제**:
```go
numbers := []int{1, 2, 3, 4, 5}

// Tap for logging without breaking the chain
result := sliceutil.Tap(numbers, func(n int) {
    fmt.Printf("Processing: %d\n", n)
})

// result is still [1 2 3 4 5]
```

---

## Common Use Cases / 일반적인 사용 사례

### Use Case 1: Data Filtering Pipeline / 데이터 필터링 파이프라인

**Scenario / 시나리오**: Filter, transform, and aggregate user data / 사용자 데이터 필터링, 변환 및 집계

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

type User struct {
    ID       int
    Name     string
    Age      int
    IsActive bool
    Salary   float64
}

func main() {
    users := []User{
        {ID: 1, Name: "Alice", Age: 28, IsActive: true, Salary: 75000},
        {ID: 2, Name: "Bob", Age: 35, IsActive: false, Salary: 85000},
        {ID: 3, Name: "Charlie", Age: 42, IsActive: true, Salary: 95000},
        {ID: 4, Name: "Diana", Age: 30, IsActive: true, Salary: 70000},
        {ID: 5, Name: "Eve", Age: 25, IsActive: false, Salary: 65000},
    }

    // Filter active users over 30
    // 30세 이상의 활성 사용자 필터링
    activeOver30 := sliceutil.Filter(users, func(u User) bool {
        return u.IsActive && u.Age > 30
    })

    // Extract names
    // 이름 추출
    names := sliceutil.Map(activeOver30, func(u User) string {
        return u.Name
    })

    // Calculate average salary
    // 평균 급여 계산
    salaries := sliceutil.Map(activeOver30, func(u User) float64 {
        return u.Salary
    })
    avgSalary := sliceutil.Average(salaries)

    fmt.Println("Active users over 30:", names)
    fmt.Printf("Average salary: $%.2f\n", avgSalary)
    // Output:
    // Active users over 30: [Charlie]
    // Average salary: $95000.00
}
```

### Use Case 2: E-commerce Product Management / 전자상거래 제품 관리

**Scenario / 시나리오**: Group products by category and find best sellers / 카테고리별로 제품을 그룹화하고 베스트셀러 찾기

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

type Product struct {
    ID       int
    Name     string
    Category string
    Price    float64
    Sales    int
}

func main() {
    products := []Product{
        {ID: 1, Name: "Laptop Pro", Category: "Electronics", Price: 1299.99, Sales: 450},
        {ID: 2, Name: "Mouse", Category: "Electronics", Price: 29.99, Sales: 1200},
        {ID: 3, Name: "Desk Chair", Category: "Furniture", Price: 249.99, Sales: 300},
        {ID: 4, Name: "Monitor", Category: "Electronics", Price: 399.99, Sales: 600},
        {ID: 5, Name: "Standing Desk", Category: "Furniture", Price: 499.99, Sales: 150},
    }

    // Group by category
    // 카테고리별로 그룹화
    byCategory := sliceutil.GroupBy(products, func(p Product) string {
        return p.Category
    })

    // Find best seller in each category
    // 각 카테고리에서 베스트셀러 찾기
    for category, items := range byCategory {
        sorted := sliceutil.SortBy(items, func(a, b Product) bool {
            return a.Sales > b.Sales // Descending order
        })

        bestSeller, _ := sliceutil.First(sorted)
        fmt.Printf("%s Best Seller: %s (%d sales)\n",
            category, bestSeller.Name, bestSeller.Sales)
    }
    // Output:
    // Electronics Best Seller: Mouse (1200 sales)
    // Furniture Best Seller: Desk Chair (300 sales)
}
```

### Use Case 3: Batch Processing / 배치 처리

**Scenario / 시나리오**: Process large dataset in chunks / 대용량 데이터셋을 청크로 처리

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
    "time"
)

func processChunk(chunk []int) {
    // Simulate processing
    // 처리 시뮬레이션
    time.Sleep(100 * time.Millisecond)
    sum := sliceutil.Sum(chunk)
    fmt.Printf("Processed chunk: sum=%d\n", sum)
}

func main() {
    // Large dataset
    // 대용량 데이터셋
    data := make([]int, 1000)
    for i := range data {
        data[i] = i + 1
    }

    // Process in batches of 100
    // 100개씩 배치 처리
    chunks := sliceutil.Chunk(data, 100)

    fmt.Printf("Processing %d chunks...\n", len(chunks))

    sliceutil.ForEachIndexed(chunks, func(i int, chunk []int) {
        fmt.Printf("Batch %d/%d: ", i+1, len(chunks))
        processChunk(chunk)
    })

    fmt.Println("All batches processed!")
}
```

### Use Case 4: Data Deduplication / 데이터 중복 제거

**Scenario / 시나리오**: Remove duplicate records based on unique identifier / 고유 식별자를 기반으로 중복 레코드 제거

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

type Record struct {
    ID        string
    Name      string
    Timestamp int64
}

func main() {
    records := []Record{
        {ID: "A001", Name: "Alice", Timestamp: 1000},
        {ID: "B002", Name: "Bob", Timestamp: 2000},
        {ID: "A001", Name: "Alice Updated", Timestamp: 3000}, // Duplicate ID
        {ID: "C003", Name: "Charlie", Timestamp: 4000},
        {ID: "B002", Name: "Bob Updated", Timestamp: 5000}, // Duplicate ID
    }

    // Keep only latest record for each ID
    // 각 ID에 대해 최신 레코드만 유지
    sorted := sliceutil.SortBy(records, func(a, b Record) bool {
        return a.Timestamp > b.Timestamp // Newest first
    })

    unique := sliceutil.UniqueBy(sorted, func(r Record) string {
        return r.ID
    })

    fmt.Println("Deduplicated records:")
    sliceutil.ForEach(unique, func(r Record) {
        fmt.Printf("ID: %s, Name: %s, Time: %d\n", r.ID, r.Name, r.Timestamp)
    })
    // Output:
    // ID: B002, Name: Bob Updated, Time: 5000
    // ID: A001, Name: Alice Updated, Time: 3000
    // ID: C003, Name: Charlie, Time: 4000
}
```

### Use Case 5: Statistical Analysis / 통계 분석

**Scenario / 시나리오**: Calculate statistics on numerical data / 숫자 데이터에 대한 통계 계산

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

func main() {
    // Test scores
    // 시험 점수
    scores := []float64{85.5, 92.0, 78.5, 88.0, 95.5, 82.0, 90.0, 87.5, 93.0, 86.0}

    // Calculate statistics
    // 통계 계산
    avg := sliceutil.Average(scores)
    min, _ := sliceutil.Min(scores)
    max, _ := sliceutil.Max(scores)
    sum := sliceutil.Sum(scores)

    // Count by grade ranges
    // 등급 범위별 개수
    gradeA := sliceutil.Count(scores, func(s float64) bool { return s >= 90 })
    gradeB := sliceutil.Count(scores, func(s float64) bool { return s >= 80 && s < 90 })
    gradeC := sliceutil.Count(scores, func(s float64) bool { return s < 80 })

    // Check if any failed (< 70)
    // 낙제(< 70) 여부 확인
    anyFailed := sliceutil.Any(scores, func(s float64) bool { return s < 70 })

    fmt.Println("=== Score Statistics ===")
    fmt.Printf("Average: %.2f\n", avg)
    fmt.Printf("Min: %.2f\n", min)
    fmt.Printf("Max: %.2f\n", max)
    fmt.Printf("Total: %.2f\n", sum)
    fmt.Printf("\nGrade Distribution:\n")
    fmt.Printf("A (90-100): %d students\n", gradeA)
    fmt.Printf("B (80-89): %d students\n", gradeB)
    fmt.Printf("C (<80): %d students\n", gradeC)
    fmt.Printf("\nAny failures: %v\n", anyFailed)
}
```

### Use Case 6: Set Operations for Access Control / 접근 제어를 위한 집합 작업

**Scenario / 시나리오**: Manage user permissions with set operations / 집합 작업으로 사용자 권한 관리

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

func main() {
    // User permissions
    // 사용자 권한
    adminPerms := []string{"read", "write", "delete", "admin"}
    userPerms := []string{"read", "write"}
    guestPerms := []string{"read"}

    // Check if user is subset of admin (has subset of permissions)
    // 사용자가 관리자의 부분집합인지 확인 (권한의 부분집합 보유)
    isValidUser := sliceutil.IsSubset(userPerms, adminPerms)
    fmt.Printf("User permissions valid: %v\n", isValidUser) // true

    // Find common permissions between user and guest
    // 사용자와 게스트 간의 공통 권한 찾기
    commonPerms := sliceutil.Intersection(userPerms, guestPerms)
    fmt.Printf("Common permissions: %v\n", commonPerms) // [read]

    // Find permissions user has but guest doesn't
    // 사용자는 있지만 게스트는 없는 권한 찾기
    extraPerms := sliceutil.Difference(userPerms, guestPerms)
    fmt.Printf("Extra user permissions: %v\n", extraPerms) // [write]

    // Combine all unique permissions
    // 모든 고유 권한 결합
    allPerms := sliceutil.Union(sliceutil.Union(adminPerms, userPerms), guestPerms)
    fmt.Printf("All permissions: %v\n", allPerms)
    // [read write delete admin]
}
```

### Use Case 7: Data Transformation Chain / 데이터 변환 체인

**Scenario / 시나리오**: Complex multi-step data transformation / 복잡한 다단계 데이터 변환

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
    "strings"
)

func main() {
    // Raw data
    // 원시 데이터
    rawData := []string{
        "  Alice-25-Engineer  ",
        "BOB-30-MANAGER",
        "  charlie-28-developer  ",
        "  DIANA-35-designer  ",
    }

    // Multi-step transformation
    // 다단계 변환
    result := sliceutil.Map(rawData, func(s string) string {
        // 1. Trim whitespace / 공백 제거
        s = strings.TrimSpace(s)

        // 2. Convert to lowercase / 소문자로 변환
        s = strings.ToLower(s)

        // 3. Split and reconstruct / 분할 및 재구성
        parts := strings.Split(s, "-")
        if len(parts) == 3 {
            // Capitalize name
            name := strings.Title(parts[0])
            age := parts[1]
            role := strings.Title(parts[2])
            return fmt.Sprintf("%s (%s years, %s)", name, age, role)
        }
        return s
    })

    fmt.Println("Transformed data:")
    sliceutil.ForEach(result, func(s string) {
        fmt.Println("-", s)
    })
    // Output:
    // - Alice (25 years, Engineer)
    // - Bob (30 years, Manager)
    // - Charlie (28 years, Developer)
    // - Diana (35 years, Designer)
}
```

### Use Case 8: Finding Outliers / 이상치 찾기

**Scenario / 시나리오**: Detect outliers in numerical data / 숫자 데이터에서 이상치 감지

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
    "math"
)

func main() {
    data := []float64{10, 12, 11, 13, 12, 100, 11, 10, 12, 13} // 100 is an outlier

    // Calculate average and filter outliers
    // 평균 계산 및 이상치 필터링
    avg := sliceutil.Average(data)

    // Define outlier as value > 2x average
    // 이상치를 평균의 2배 이상인 값으로 정의
    threshold := avg * 2

    outliers := sliceutil.Filter(data, func(v float64) bool {
        return math.Abs(v) > threshold
    })

    normalData := sliceutil.Filter(data, func(v float64) bool {
        return math.Abs(v) <= threshold
    })

    fmt.Printf("Average: %.2f\n", avg)
    fmt.Printf("Threshold: %.2f\n", threshold)
    fmt.Printf("Outliers: %v\n", outliers)
    fmt.Printf("Normal data: %v\n", normalData)
    fmt.Printf("Normal average: %.2f\n", sliceutil.Average(normalData))
    // Output:
    // Average: 19.40
    // Threshold: 38.80
    // Outliers: [100]
    // Normal data: [10 12 11 13 12 11 10 12 13]
    // Normal average: 11.56
}
```

---

## Best Practices / 모범 사례

### 1. Prefer Immutability / 불변성 선호

All Sliceutil functions return new slices without modifying the original. Always use the returned value.

모든 Sliceutil 함수는 원본을 수정하지 않고 새 슬라이스를 반환합니다. 항상 반환값을 사용하세요.

```go
// ✅ Good / 좋음
numbers := []int{1, 2, 3, 4, 5}
doubled := sliceutil.Map(numbers, func(n int) int { return n * 2 })
fmt.Println(doubled) // [2 4 6 8 10]
fmt.Println(numbers) // [1 2 3 4 5] (unchanged)

// ❌ Bad - Ignoring return value / 나쁨 - 반환값 무시
sliceutil.Map(numbers, func(n int) int { return n * 2 })
// numbers is still [1 2 3 4 5]
```

### 2. Chain Operations / 작업 체이닝

Chain multiple operations for readable pipelines.

읽기 쉬운 파이프라인을 위해 여러 작업을 체이닝하세요.

```go
// ✅ Good - Chained operations / 좋음 - 체이닝 작업
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

result := sliceutil.Map(
    sliceutil.Filter(numbers, func(n int) bool {
        return n%2 == 0
    }),
    func(n int) int {
        return n * n
    },
)
// result: [4 16 36 64 100]

// ❌ Less readable / 덜 읽기 쉬움
evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
squared := sliceutil.Map(evens, func(n int) int { return n * n })
```

### 3. Use Type Inference / 타입 추론 사용

Let Go infer types when possible for cleaner code.

가능한 경우 Go가 타입을 추론하도록 하여 더 깔끔한 코드를 작성하세요.

```go
// ✅ Good - Type inference / 좋음 - 타입 추론
numbers := []int{1, 2, 3}
doubled := sliceutil.Map(numbers, func(n int) int { return n * 2 })

// ❌ Unnecessary explicit types / 불필요한 명시적 타입
numbers := []int{1, 2, 3}
doubled := sliceutil.Map[int, int](numbers, func(n int) int { return n * 2 })
```

### 4. Handle Empty Slices / 빈 슬라이스 처리

Always check for empty slices when using functions that may fail.

실패할 수 있는 함수를 사용할 때는 항상 빈 슬라이스를 확인하세요.

```go
// ✅ Good - Check before use / 좋음 - 사용 전 확인
numbers := []int{}
if !sliceutil.IsEmpty(numbers) {
    first, _ := sliceutil.First(numbers)
    fmt.Println(first)
} else {
    fmt.Println("Empty slice")
}

// ✅ Good - Use ok pattern / 좋음 - ok 패턴 사용
if first, ok := sliceutil.First(numbers); ok {
    fmt.Println(first)
}
```

### 5. Prefer Specific Functions Over Generic Ones / 제네릭 함수보다 특정 함수 선호

Use specialized functions like `Sum`, `Average`, `Min`, `Max` instead of `Reduce` when available.

가능한 경우 `Reduce` 대신 `Sum`, `Average`, `Min`, `Max`와 같은 특수화된 함수를 사용하세요.

```go
// ✅ Good - Use specific function / 좋음 - 특정 함수 사용
numbers := []int{1, 2, 3, 4, 5}
sum := sliceutil.Sum(numbers)

// ❌ Less readable - Using Reduce / 덜 읽기 쉬움 - Reduce 사용
sum := sliceutil.Reduce(numbers, 0, func(acc, n int) int {
    return acc + n
})
```

### 6. Use GroupBy for Complex Aggregations / 복잡한 집계에 GroupBy 사용

When you need to aggregate data by categories, use `GroupBy`.

카테고리별로 데이터를 집계해야 할 때는 `GroupBy`를 사용하세요.

```go
// ✅ Good - GroupBy for category-based operations / 좋음 - 카테고리 기반 작업에 GroupBy
type Product struct {
    Name     string
    Category string
    Price    float64
}

products := []Product{...}
byCategory := sliceutil.GroupBy(products, func(p Product) string {
    return p.Category
})

for category, items := range byCategory {
    prices := sliceutil.Map(items, func(p Product) float64 { return p.Price })
    avg := sliceutil.Average(prices)
    fmt.Printf("%s: $%.2f\n", category, avg)
}
```

### 7. Avoid Modifying Elements in ForEach / ForEach에서 요소 수정 방지

`ForEach` is for side effects only. Use `Map` for transformations.

`ForEach`는 부작용을 위한 것입니다. 변환에는 `Map`을 사용하세요.

```go
// ✅ Good - Map for transformation / 좋음 - 변환에 Map
numbers := []int{1, 2, 3}
doubled := sliceutil.Map(numbers, func(n int) int { return n * 2 })

// ❌ Bad - Trying to modify in ForEach / 나쁨 - ForEach에서 수정 시도
result := make([]int, 0)
sliceutil.ForEach(numbers, func(n int) {
    result = append(result, n*2) // Side effect, less clear
})
```

### 8. Use Partition Instead of Multiple Filters / 여러 필터 대신 Partition 사용

When you need both matching and non-matching elements, use `Partition`.

일치하는 요소와 일치하지 않는 요소가 모두 필요한 경우 `Partition`을 사용하세요.

```go
// ✅ Good - Single pass with Partition / 좋음 - Partition으로 단일 패스
numbers := []int{1, 2, 3, 4, 5, 6}
evens, odds := sliceutil.Partition(numbers, func(n int) bool {
    return n%2 == 0
})

// ❌ Less efficient - Two passes / 덜 효율적 - 두 번의 패스
evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
odds := sliceutil.Filter(numbers, func(n int) bool { return n%2 != 0 })
```

### 9. Use UniqueBy for Complex Uniqueness / 복잡한 고유성에 UniqueBy 사용

For structs or complex types, use `UniqueBy` with a key function.

구조체나 복잡한 타입의 경우 키 함수와 함께 `UniqueBy`를 사용하세요.

```go
// ✅ Good - UniqueBy for structs / 좋음 - 구조체에 UniqueBy
type User struct {
    ID   int
    Name string
}

users := []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 1, Name: "Alice Duplicate"},
}

unique := sliceutil.UniqueBy(users, func(u User) int {
    return u.ID
})
```

### 10. Combine Set Operations / 집합 작업 결합

Use set operations for permission checks, data reconciliation, etc.

권한 확인, 데이터 조정 등에 집합 작업을 사용하세요.

```go
// ✅ Good - Set operations for permissions / 좋음 - 권한 확인에 집합 작업
requiredPerms := []string{"read", "write", "delete"}
userPerms := []string{"read", "write"}

hasAllPerms := sliceutil.IsSubset(requiredPerms, userPerms)
missingPerms := sliceutil.Difference(requiredPerms, userPerms)

if !hasAllPerms {
    fmt.Println("Missing permissions:", missingPerms)
}
```

### 11. Use Chunk for Batch Processing / 배치 처리에 Chunk 사용

When processing large datasets, use `Chunk` to process in batches.

대용량 데이터셋을 처리할 때는 `Chunk`를 사용하여 배치로 처리하세요.

```go
// ✅ Good - Process in chunks / 좋음 - 청크로 처리
largeDataset := make([]int, 10000)
chunks := sliceutil.Chunk(largeDataset, 100)

sliceutil.ForEach(chunks, func(chunk []int) {
    processBatch(chunk) // Process 100 items at a time
})
```

### 12. Check Sort Status Before Sorting / 정렬 전에 정렬 상태 확인

Use `IsSorted` to avoid unnecessary sorting.

불필요한 정렬을 피하기 위해 `IsSorted`를 사용하세요.

```go
// ✅ Good - Check before sorting / 좋음 - 정렬 전 확인
numbers := []int{1, 2, 3, 4, 5}

if !sliceutil.IsSorted(numbers) {
    numbers = sliceutil.Sort(numbers)
}
```

---

## Troubleshooting / 문제 해결

### Issue 1: Type Mismatch Errors / 타입 불일치 오류

**Problem / 문제**: Compiler error "type does not implement comparable"

**Solution / 해결책**: Some functions like `Contains`, `Unique`, `Union` require comparable types. For custom types, implement equality differently.

일부 함수(예: `Contains`, `Unique`, `Union`)는 비교 가능한 타입이 필요합니다. 사용자 정의 타입의 경우 동등성을 다르게 구현하세요.

```go
// ❌ Won't compile - struct is not comparable / 컴파일 안됨 - 구조체는 비교 불가
type User struct {
    ID   int
    Name string
    Tags []string // Slice makes it non-comparable
}

users := []User{...}
unique := sliceutil.Unique(users) // ❌ Error

// ✅ Solution - Use UniqueBy / 해결책 - UniqueBy 사용
unique := sliceutil.UniqueBy(users, func(u User) int {
    return u.ID
})
```

### Issue 2: Empty Slice Panics / 빈 슬라이스 패닉

**Problem / 문제**: Panic when accessing elements of empty slice

**Solution / 해결책**: Always check if slice is empty or use functions that return (value, bool).

항상 슬라이스가 비어있는지 확인하거나 (값, bool)을 반환하는 함수를 사용하세요.

```go
// ❌ May panic / 패닉 가능
numbers := []int{}
first, _ := sliceutil.First(numbers) // Returns zero value, false
fmt.Println(first * 2) // 0 * 2 = 0 (works, but misleading)

// ✅ Safe approach / 안전한 접근
if first, ok := sliceutil.First(numbers); ok {
    fmt.Println(first * 2)
} else {
    fmt.Println("Empty slice")
}
```

### Issue 3: Performance with Large Slices / 대용량 슬라이스 성능

**Problem / 문제**: Slow performance with very large slices

**Solution / 해결책**: Use chunking for batch processing and avoid unnecessary copies.

배치 처리를 위해 청킹을 사용하고 불필요한 복사를 피하세요.

```go
// ❌ Slow - Processing entire large slice at once / 느림 - 전체 대용량 슬라이스 한 번에 처리
largeData := make([]int, 1000000)
result := sliceutil.Map(largeData, expensiveOperation)

// ✅ Better - Process in chunks / 더 나음 - 청크로 처리
chunks := sliceutil.Chunk(largeData, 1000)
var results []int
for _, chunk := range chunks {
    chunkResult := sliceutil.Map(chunk, expensiveOperation)
    results = append(results, chunkResult...)
}
```

### Issue 4: Unexpected Nil Slices / 예기치 않은 Nil 슬라이스

**Problem / 문제**: Functions return nil instead of empty slice

**Solution / 해결책**: Most functions return empty slices, but check for nil if needed.

대부분의 함수는 빈 슬라이스를 반환하지만 필요한 경우 nil을 확인하세요.

```go
// ✅ Check for nil / nil 확인
var numbers []int // nil slice
if numbers == nil {
    numbers = []int{} // Initialize to empty slice
}

result := sliceutil.Filter(numbers, someCondition)
// result will be an empty slice, not nil
```

### Issue 5: Incorrect Generic Type Inference / 잘못된 제네릭 타입 추론

**Problem / 문제**: Compiler cannot infer types

**Solution / 해결책**: Explicitly specify type parameters when needed.

필요한 경우 타입 매개변수를 명시적으로 지정하세요.

```go
// ❌ Type inference may fail / 타입 추론 실패 가능
result := sliceutil.Map(data, func(x interface{}) string {
    return fmt.Sprint(x)
})

// ✅ Explicit types / 명시적 타입
result := sliceutil.Map[interface{}, string](data, func(x interface{}) string {
    return fmt.Sprint(x)
})
```

### Issue 6: Modifying Original Slice / 원본 슬라이스 수정

**Problem / 문제**: Accidentally modifying original slice

**Solution / 해결책**: Remember that Sliceutil functions are immutable - they don't modify the original.

Sliceutil 함수는 불변이라는 것을 기억하세요 - 원본을 수정하지 않습니다.

```go
// ❌ Expecting original to change / 원본 변경 기대
numbers := []int{1, 2, 3, 4, 5}
sliceutil.Filter(numbers, func(n int) bool { return n > 3 })
fmt.Println(numbers) // Still [1 2 3 4 5]

// ✅ Correct - Use return value / 올바름 - 반환값 사용
numbers = sliceutil.Filter(numbers, func(n int) bool { return n > 3 })
fmt.Println(numbers) // [4 5]
```

---

## FAQ / 자주 묻는 질문

### Q1: Do I need Go 1.18+? / Go 1.18+가 필요한가요?

**A**: Yes, Sliceutil requires Go 1.18 or higher because it uses generics. Generics were introduced in Go 1.18.

**A**: 예, Sliceutil은 제네릭을 사용하므로 Go 1.18 이상이 필요합니다. 제네릭은 Go 1.18에서 도입되었습니다.

### Q2: Are these functions thread-safe? / 이 함수들은 스레드 안전한가요?

**A**: Yes, all functions are thread-safe because they don't modify the original slice and don't use shared state. However, if you're accessing the same slice from multiple goroutines, you still need synchronization.

**A**: 예, 모든 함수는 원본 슬라이스를 수정하지 않고 공유 상태를 사용하지 않으므로 스레드 안전합니다. 그러나 여러 고루틴에서 동일한 슬라이스에 액세스하는 경우에는 여전히 동기화가 필요합니다.

### Q3: What's the performance impact? / 성능 영향은 어떤가요?

**A**: Sliceutil functions are optimized for common cases and typically have O(n) complexity. They create new slices, so there's a memory allocation cost, but this ensures immutability and safety.

**A**: Sliceutil 함수는 일반적인 경우에 최적화되어 있으며 일반적으로 O(n) 복잡도를 갖습니다. 새 슬라이스를 생성하므로 메모리 할당 비용이 있지만, 이는 불변성과 안전성을 보장합니다.

### Q4: Can I use this with custom types? / 사용자 정의 타입과 함께 사용할 수 있나요?

**A**: Yes! Most functions work with any type (`any` constraint). Functions requiring comparison use `comparable` constraint. For complex comparisons, use functions like `FindIndex`, `UniqueBy`, `SortBy` that accept custom predicates.

**A**: 예! 대부분의 함수는 모든 타입(`any` 제약)과 작동합니다. 비교가 필요한 함수는 `comparable` 제약을 사용합니다. 복잡한 비교의 경우 사용자 정의 조건자를 허용하는 `FindIndex`, `UniqueBy`, `SortBy`와 같은 함수를 사용하세요.

### Q5: How do I contribute or report bugs? / 기여하거나 버그를 보고하려면 어떻게 하나요?

**A**: Visit our GitHub repository at https://github.com/arkd0ng/go-utils and open an issue or pull request.

**A**: https://github.com/arkd0ng/go-utils의 GitHub 저장소를 방문하여 이슈나 풀 리퀘스트를 여세요.

### Q6: Why immutable operations? / 왜 불변 작업인가요?

**A**: Immutability prevents accidental bugs, makes code easier to reason about, and allows safe concurrent access. It's a core principle of functional programming.

**A**: 불변성은 우발적인 버그를 방지하고, 코드를 더 쉽게 이해할 수 있게 하며, 안전한 동시 액세스를 허용합니다. 이는 함수형 프로그래밍의 핵심 원칙입니다.

### Q7: Can I chain operations? / 작업을 체이닝할 수 있나요?

**A**: Yes, you can chain operations by nesting function calls or assigning intermediate results.

**A**: 예, 함수 호출을 중첩하거나 중간 결과를 할당하여 작업을 체이닝할 수 있습니다.

```go
result := sliceutil.Map(
    sliceutil.Filter(data, condition),
    transformation,
)
```

### Q8: What's the difference between Filter and Find? / Filter와 Find의 차이점은 무엇인가요?

**A**: `Filter` returns all matching elements as a slice. `Find` returns only the first matching element and a boolean indicating if it was found.

**A**: `Filter`는 일치하는 모든 요소를 슬라이스로 반환합니다. `Find`는 첫 번째 일치 요소만 반환하고 발견 여부를 나타내는 불린 값을 반환합니다.

### Q9: Can I use this in production? / 프로덕션에서 사용할 수 있나요?

**A**: Yes, Sliceutil is designed for production use with comprehensive tests, benchmarks, and zero external dependencies.

**A**: 예, Sliceutil은 종합적인 테스트, 벤치마크 및 제로 외부 의존성으로 프로덕션 사용을 위해 설계되었습니다.

### Q10: How does Sliceutil compare to other libraries? / Sliceutil은 다른 라이브러리와 어떻게 비교되나요?

**A**: Sliceutil focuses on extreme simplicity, type safety with generics, zero dependencies, and comprehensive coverage of 60 functions across 8 categories. It's designed specifically for Go 1.18+ with modern best practices.

**A**: Sliceutil은 극도의 간결함, 제네릭을 사용한 타입 안전성, 제로 의존성, 8개 카테고리에 걸쳐 60개 함수의 포괄적인 커버리지에 중점을 둡니다. Go 1.18+ 및 최신 모범 사례를 위해 특별히 설계되었습니다.

---

## Additional Resources / 추가 리소스

- **Package README**: https://github.com/arkd0ng/go-utils/tree/main/sliceutil
- **Developer Guide**: https://github.com/arkd0ng/go-utils/tree/main/docs/sliceutil/DEVELOPER_GUIDE.md
- **Examples**: https://github.com/arkd0ng/go-utils/tree/main/examples/sliceutil
- **GitHub Repository**: https://github.com/arkd0ng/go-utils
- **Issue Tracker**: https://github.com/arkd0ng/go-utils/issues

---

**End of User Manual / 사용자 매뉴얼 끝**

For developer information, see the [Developer Guide](DEVELOPER_GUIDE.md).

개발자 정보는 [개발자 가이드](DEVELOPER_GUIDE.md)를 참조하세요.
