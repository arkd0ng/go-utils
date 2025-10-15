# sliceutil - Slice Utilities / 슬라이스 유틸리티

**v1.7.023** - Extreme Simplicity for Slice Operations! 🎉

Extreme simplicity slice utility functions for Go - reduce 10-20 lines of slice manipulation code to just 1 line.

극도로 간단한 Go용 슬라이스 유틸리티 함수 - 10-20줄의 슬라이스 조작 코드를 단 1줄로 줄입니다.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview / 개요

The `sliceutil` package provides **95 type-safe functions** for common slice operations in Go. Stop writing repetitive loops and start using functional programming style.

`sliceutil` 패키지는 Go에서 일반적인 슬라이스 작업을 위한 **95개의 타입 안전 함수**를 제공합니다. 반복적인 루프 작성을 멈추고 함수형 프로그래밍 스타일을 사용하세요.

### Design Philosophy: "20 lines → 1 line" / 설계 철학: "20줄 → 1줄"

**Before (Standard Go) / 이전 (표준 Go)**:
```go
// Filter even numbers / 짝수 필터링
numbers := []int{1, 2, 3, 4, 5, 6}
var evens []int
for _, n := range numbers {
    if n%2 == 0 {
        evens = append(evens, n)
    }
}

// Map to strings / 문자열로 매핑
var strings []string
for _, n := range evens {
    strings = append(strings, fmt.Sprintf("num_%d", n))
}

// Check if contains / 포함 확인
found := false
for _, s := range strings {
    if s == "num_4" {
        found = true
        break
    }
}
// 20+ lines of code
```

**After (This Package) / 이후 (이 패키지)**:
```go
numbers := []int{1, 2, 3, 4, 5, 6}
evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
strings := sliceutil.Map(evens, func(n int) string { return fmt.Sprintf("num_%d", n) })
found := sliceutil.Contains(strings, "num_4")
// 3 lines of code (vs 20+)
```

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/sliceutil
```

**Requirements / 요구사항**:
- Go 1.18 or higher (for generics support) / Go 1.18 이상 (제네릭 지원)

## Quick Start / 빠른 시작

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/sliceutil"
)

func main() {
    // Basic Operations / 기본 작업
    numbers := []int{1, 2, 3, 4, 5, 6}

    // Check if contains / 포함 확인
    hasThree := sliceutil.Contains(numbers, 3) // true

    // Find index / 인덱스 찾기
    index := sliceutil.IndexOf(numbers, 4) // 3

    // Transformation / 변환
    // Filter even numbers / 짝수 필터링
    evens := sliceutil.Filter(numbers, func(n int) bool {
        return n%2 == 0
    }) // [2, 4, 6]

    // Map to double / 2배로 매핑
    doubled := sliceutil.Map(numbers, func(n int) int {
        return n * 2
    }) // [2, 4, 6, 8, 10, 12]

    // Remove duplicates / 중복 제거
    items := []int{1, 2, 2, 3, 3, 3, 4}
    unique := sliceutil.Unique(items) // [1, 2, 3, 4]

    // Aggregation / 집계
    // Sum all numbers / 모든 숫자의 합
    sum := sliceutil.Sum(numbers) // 21

    // Find max / 최대값 찾기
    max, _ := sliceutil.Max(numbers) // 6

    // Group by condition / 조건으로 그룹화
    trues, falses := sliceutil.Partition(numbers, func(n int) bool {
        return n > 3
    }) // trues: [4, 5, 6], falses: [1, 2, 3]

    // Slicing / 슬라이싱
    // Take first 3 items / 처음 3개 항목
    first3 := sliceutil.Take(numbers, 3) // [1, 2, 3]

    // Split into chunks / 청크로 분할
    chunks := sliceutil.Chunk(numbers, 2) // [[1, 2], [3, 4], [5, 6]]

    // Set Operations / 집합 작업
    a := []int{1, 2, 3, 4}
    b := []int{3, 4, 5, 6}

    union := sliceutil.Union(a, b) // [1, 2, 3, 4, 5, 6]
    intersection := sliceutil.Intersection(a, b) // [3, 4]
    difference := sliceutil.Difference(a, b) // [1, 2]

    // Sorting / 정렬
    unsorted := []int{3, 1, 4, 1, 5, 9, 2}
    sorted := sliceutil.Sort(unsorted) // [1, 1, 2, 3, 4, 5, 9]

    // Predicates / 조건 검사
    // Check if all are positive / 모두 양수인지 확인
    allPositive := sliceutil.All(numbers, func(n int) bool {
        return n > 0
    }) // true

    // Check if any is greater than 5 / 5보다 큰 것이 있는지 확인
    anyGreater5 := sliceutil.Any(numbers, func(n int) bool {
        return n > 5
    }) // true

    // Utilities / 유틸리티
    // Clone slice / 슬라이스 복제
    cloned := sliceutil.Clone(numbers)

    // Shuffle / 섞기
    shuffled := sliceutil.Shuffle(numbers)

    // Join to string / 문자열로 결합
    joined := sliceutil.Join(numbers, ", ") // "1, 2, 3, 4, 5, 6"

    fmt.Println(hasThree, evens, sum, sorted, allPositive)
}
```

## Key Features / 주요 기능

### 1. Type-Safe with Generics / 제네릭으로 타입 안전

```go
// Works with any type / 모든 타입과 작동
numbers := []int{1, 2, 3}
strings := []string{"a", "b", "c"}
users := []User{{Name: "Alice"}, {Name: "Bob"}}

// Type-safe operations / 타입 안전 작업
sliceutil.Contains(numbers, 2)  // int
sliceutil.Contains(strings, "b") // string
sliceutil.Map(users, func(u User) string { return u.Name }) // User → string
```

### 2. Functional Programming Style / 함수형 프로그래밍 스타일

```go
// Chaining operations / 작업 체이닝
result := sliceutil.Filter(
    sliceutil.Map(
        sliceutil.Unique([]int{1, 2, 2, 3, 3, 4}),
        func(n int) int { return n * 2 },
    ),
    func(n int) bool { return n > 4 },
) // [6, 8]
```

### 3. Immutable Operations / 불변 작업

```go
// Original slice is never modified / 원본 슬라이스는 변경되지 않음
original := []int{1, 2, 3}
doubled := sliceutil.Map(original, func(n int) int { return n * 2 })
// original: [1, 2, 3] (unchanged)
// doubled: [2, 4, 6] (new slice)
```

### 4. Zero External Dependencies / 제로 외부 의존성

All functions use only the standard library.

모든 함수는 표준 라이브러리만 사용합니다.

### 5. Comprehensive Coverage / 포괄적인 커버리지

95 functions across 14 categories cover 99% of common slice operations.

14개 카테고리의 95개 함수가 일반적인 슬라이스 작업의 99%를 커버합니다.

## Function Categories / 함수 카테고리

### 1. Basic Operations (11 functions) / 기본 작업 (11개 함수)

Essential operations for searching and checking slices.

슬라이스 검색 및 확인을 위한 필수 작업.

```go
sliceutil.Contains(slice, item)           // Check if contains
sliceutil.ContainsFunc(slice, predicate)  // Check with predicate
sliceutil.IndexOf(slice, item)            // Find first index
sliceutil.LastIndexOf(slice, item)        // Find last index
sliceutil.Find(slice, predicate)          // Find first match
sliceutil.FindLast(slice, predicate)      // Find last match
sliceutil.FindIndex(slice, predicate)     // Find first match index
sliceutil.Count(slice, predicate)         // Count matches
sliceutil.IsEmpty(slice)                  // Check if empty
sliceutil.IsNotEmpty(slice)               // Check if not empty
sliceutil.Equal(a, b)                     // Compare two slices
```

### 2. Transformation (8 functions) / 변환 (8개 함수)

Transform slices into different forms.

슬라이스를 다른 형태로 변환.

```go
sliceutil.Map(slice, mapper)              // Transform each item
sliceutil.Filter(slice, predicate)        // Keep matching items
sliceutil.FlatMap(slice, mapper)          // Map and flatten
sliceutil.Flatten(slice)                  // Flatten nested slices
sliceutil.Unique(slice)                   // Remove duplicates
sliceutil.UniqueBy(slice, keyFunc)        // Remove duplicates by key
sliceutil.Compact(slice)                  // Remove zero values
sliceutil.Reverse(slice)                  // Reverse order
```

### 3. Aggregation (11 functions) / 집계 (11개 함수)

Reduce slices to single values or group items.

슬라이스를 단일 값으로 줄이거나 항목 그룹화.

```go
sliceutil.Reduce(slice, initial, reducer) // Custom aggregation
sliceutil.ReduceRight(slice, initial, reducer) // Reduce from right
sliceutil.Sum(slice)                      // Sum of numbers
sliceutil.Min(slice)                      // Minimum value
sliceutil.Max(slice)                      // Maximum value
sliceutil.MinBy(slice, keyFunc)           // Minimum by custom key
sliceutil.MaxBy(slice, keyFunc)           // Maximum by custom key
sliceutil.Average(slice)                  // Average of numbers
sliceutil.GroupBy(slice, keyFunc)         // Group by key
sliceutil.CountBy(slice, keyFunc)         // Count occurrences by key
sliceutil.Partition(slice, predicate)     // Split by condition
```

### 4. Slicing (11 functions) / 슬라이싱 (11개 함수)

Extract portions of slices.

슬라이스의 일부 추출.

```go
sliceutil.Chunk(slice, size)              // Split into chunks
sliceutil.Take(slice, n)                  // Take first n items
sliceutil.TakeLast(slice, n)              // Take last n items
sliceutil.TakeWhile(slice, predicate)     // Take while predicate is true
sliceutil.Drop(slice, n)                  // Skip first n items
sliceutil.DropLast(slice, n)              // Skip last n items
sliceutil.DropWhile(slice, predicate)     // Drop while predicate is true
sliceutil.Slice(slice, start, end)        // Extract range
sliceutil.Sample(slice, n)                // Random sampling
sliceutil.Window(slice, size)             // Create sliding windows
sliceutil.Interleave(slices...)           // Interleave multiple slices
```

### 5. Set Operations (6 functions) / 집합 작업 (6개 함수)

Treat slices as mathematical sets.

슬라이스를 수학적 집합으로 처리.

```go
sliceutil.Union(a, b)                     // Union of two slices
sliceutil.Intersection(a, b)              // Intersection of two slices
sliceutil.Difference(a, b)                // Items in a but not b
sliceutil.SymmetricDifference(a, b)       // Items in either but not both
sliceutil.IsSubset(a, b)                  // Check if a ⊆ b
sliceutil.IsSuperset(a, b)                // Check if a ⊇ b
```

### 6. Sorting (6 functions) / 정렬 (6개 함수)

Sort and check sorting order.

정렬 및 정렬 순서 확인.

```go
sliceutil.Sort(slice)                     // Sort ascending
sliceutil.SortDesc(slice)                 // Sort descending
sliceutil.SortBy(slice, keyFunc)          // Sort by custom key
sliceutil.SortByMulti(slice, lessFunc)    // Sort by multiple keys
sliceutil.IsSorted(slice)                 // Check if sorted ascending
sliceutil.IsSortedDesc(slice)             // Check if sorted descending
```

### 7. Predicates (6 functions) / 조건 검사 (6개 함수)

Check conditions across all or some items.

모든 항목 또는 일부 항목에서 조건 확인.

```go
sliceutil.All(slice, predicate)           // Check if all match
sliceutil.Any(slice, predicate)           // Check if any matches
sliceutil.None(slice, predicate)          // Check if none matches
sliceutil.AllEqual(slice)                 // Check if all items equal
sliceutil.IsSortedBy(slice, keyFunc)      // Check if sorted by key
sliceutil.ContainsAll(slice, items...)    // Check if contains all
```

### 8. Utilities (12 functions) / 유틸리티 (12개 함수)

Miscellaneous helpful operations.

기타 유용한 작업.

```go
sliceutil.ForEach(slice, fn)              // Iterate with side effects
sliceutil.ForEachIndexed(slice, fn)       // Iterate with index
sliceutil.Tap(slice, fn)                  // Execute side effect and return slice
sliceutil.Join(slice, separator)          // Convert to string
sliceutil.Clone(slice)                    // Deep copy
sliceutil.Fill(slice, value)              // Fill with value
sliceutil.Insert(slice, index, items...)  // Insert at index
sliceutil.Remove(slice, index)            // Remove at index
sliceutil.RemoveAll(slice, item)          // Remove all occurrences
sliceutil.Shuffle(slice)                  // Randomize order
sliceutil.Zip(a, b)                       // Combine two slices
sliceutil.Unzip(slice)                    // Split pairs
```

### 9. Combinatorial Operations (2 functions) / 조합 작업 (2개 함수)

Generate permutations and combinations from slices.

슬라이스에서 순열과 조합 생성.

```go
sliceutil.Permutations(slice)             // All possible permutations (n!)
sliceutil.Combinations(slice, k)          // All k-combinations C(n,k)
```

**Performance Warning / 성능 경고**:
- Permutations grow factorially: n=5 → 120, n=10 → 3,628,800
- Combinations: C(10,5) = 252, C(20,10) = 184,756
- Use with caution for large slices!

**성능 경고**:
- 순열은 팩토리얼로 증가: n=5 → 120, n=10 → 3,628,800
- 조합: C(10,5) = 252, C(20,10) = 184,756
- 큰 슬라이스에는 주의해서 사용하세요!

### 10. Statistics (8 functions) / 통계 (8개 함수)

Perform statistical operations on numeric slices.

숫자 슬라이스에 대한 통계 작업 수행.

```go
sliceutil.Median(slice)                   // Calculate median
sliceutil.Mode(slice)                     // Find most frequent value
sliceutil.Frequencies(slice)              // Get frequency map
sliceutil.Percentile(slice, p)            // Calculate percentile
sliceutil.StandardDeviation(slice)        // Calculate std dev
sliceutil.Variance(slice)                 // Calculate variance
sliceutil.MostCommon(slice, n)            // Get n most common values
sliceutil.LeastCommon(slice, n)           // Get n least common values
```

### 11. Diff/Comparison (4 functions) / 차이/비교 (4개 함수)

Compare slices and detect differences.

슬라이스 비교 및 차이 감지.

```go
sliceutil.Diff(old, new)                  // Find added/removed/unchanged
sliceutil.DiffBy(old, new, keyFunc)       // Diff by custom key
sliceutil.EqualUnordered(a, b)            // Compare ignoring order
sliceutil.HasDuplicates(slice)            // Check for duplicates
```

### 12. Index-based (3 functions) / 인덱스 기반 (3개 함수)

Operate on slices using indices.

인덱스를 사용한 슬라이스 작업.

```go
sliceutil.FindIndices(slice, predicate)   // Find all matching indices
sliceutil.AtIndices(slice, indices)       // Get elements at indices
sliceutil.RemoveIndices(slice, indices)   // Remove elements at indices
```

### 13. Conditional (3 functions) / 조건부 (3개 함수)

Transform slices based on conditions.

조건에 따른 슬라이스 변환.

```go
sliceutil.ReplaceIf(slice, predicate, value)  // Replace matching items
sliceutil.ReplaceAll(slice, old, new)         // Replace all occurrences
sliceutil.UpdateWhere(slice, predicate, fn)   // Update matching items
```

### 14. Advanced (4 functions) / 고급 (4개 함수)

Advanced functional programming operations.

고급 함수형 프로그래밍 작업.

```go
sliceutil.Scan(slice, initial, fn)        // Cumulative aggregation
sliceutil.ZipWith(a, b, fn)               // Combine with custom function
sliceutil.RotateLeft(slice, n)            // Rotate left by n positions
sliceutil.RotateRight(slice, n)           // Rotate right by n positions
```

## Real-World Examples / 실제 사용 예제

### Example 1: Data Processing Pipeline / 데이터 처리 파이프라인

```go
// Process user data / 사용자 데이터 처리
users := []User{
    {Name: "Alice", Age: 25, City: "Seoul"},
    {Name: "Bob", Age: 30, City: "Busan"},
    {Name: "Charlie", Age: 35, City: "Seoul"},
}

// Get names of Seoul users over 20 / 20세 이상의 서울 사용자 이름 가져오기
seoulAdultNames := sliceutil.Map(
    sliceutil.Filter(users, func(u User) bool {
        return u.City == "Seoul" && u.Age > 20
    }),
    func(u User) string { return u.Name },
) // ["Alice", "Charlie"]
```

### Example 2: Data Validation / 데이터 검증

```go
// Check if all prices are positive / 모든 가격이 양수인지 확인
prices := []float64{10.5, 20.0, 30.25, 15.75}
allPositive := sliceutil.All(prices, func(p float64) bool {
    return p > 0
}) // true

// Check if any price is over 100 / 100 이상의 가격이 있는지 확인
anyExpensive := sliceutil.Any(prices, func(p float64) bool {
    return p > 100
}) // false
```

### Example 3: Data Aggregation / 데이터 집계

```go
// Group orders by status / 상태별 주문 그룹화
orders := []Order{
    {ID: 1, Status: "pending"},
    {ID: 2, Status: "shipped"},
    {ID: 3, Status: "pending"},
}

byStatus := sliceutil.GroupBy(orders, func(o Order) string {
    return o.Status
})
// map[string][]Order{
//     "pending": [{1, "pending"}, {3, "pending"}],
//     "shipped": [{2, "shipped"}],
// }
```

### Example 4: Batch Processing / 배치 처리

```go
// Process items in batches of 100 / 100개씩 배치 처리
items := make([]int, 1000)
batches := sliceutil.Chunk(items, 100) // 10 batches of 100 items

for _, batch := range batches {
    processBatch(batch)
}
```

### Example 5: Set Operations / 집합 작업

```go
// Find common tags between posts / 게시물 간 공통 태그 찾기
post1Tags := []string{"go", "programming", "backend"}
post2Tags := []string{"go", "web", "programming"}

commonTags := sliceutil.Intersection(post1Tags, post2Tags)
// ["go", "programming"]

allTags := sliceutil.Union(post1Tags, post2Tags)
// ["go", "programming", "backend", "web"]
```

## Performance / 성능

All functions are implemented with performance in mind:

모든 함수는 성능을 고려하여 구현되었습니다:

- Efficient algorithms / 효율적인 알고리즘
- Minimal allocations / 최소한의 할당
- Benchmark tests included / 벤치마크 테스트 포함

For performance-critical code, consider:

성능이 중요한 코드의 경우 다음을 고려하세요:

1. Use preallocated slices when possible / 가능한 경우 미리 할당된 슬라이스 사용
2. Avoid unnecessary transformations / 불필요한 변환 방지
3. Check benchmarks for your use case / 사용 사례에 대한 벤치마크 확인

## Documentation / 문서

For more detailed documentation:

더 자세한 문서:

- **[User Manual](../docs/sliceutil/USER_MANUAL.md)** - Comprehensive guide with examples / 예제가 있는 포괄적인 가이드
- **[Developer Guide](../docs/sliceutil/DEVELOPER_GUIDE.md)** - Architecture and implementation details / 아키텍처 및 구현 세부사항
- **[Design Plan](../docs/sliceutil/DESIGN_PLAN.md)** - Design philosophy and decisions / 설계 철학 및 결정
- **[Work Plan](../docs/sliceutil/WORK_PLAN.md)** - Implementation roadmap / 구현 로드맵

## Testing / 테스트

Run tests:

테스트 실행:

```bash
# Run all tests / 모든 테스트 실행
go test ./sliceutil -v

# Run with coverage / 커버리지와 함께 실행
go test ./sliceutil -cover

# Run benchmarks / 벤치마크 실행
go test ./sliceutil -bench=.
```

## Contributing / 기여하기

Contributions are welcome! Please see [CONTRIBUTING.md](../CONTRIBUTING.md) for details.

기여를 환영합니다! 자세한 내용은 [CONTRIBUTING.md](../CONTRIBUTING.md)를 참조하세요.

## License / 라이선스

MIT License - see [LICENSE](../LICENSE) for details.

MIT 라이선스 - 자세한 내용은 [LICENSE](../LICENSE)를 참조하세요.

## Version / 버전

Current version: **v1.7.023**

For version history, see [CHANGELOG-v1.7.md](../docs/CHANGELOG/CHANGELOG-v1.7.md).

버전 히스토리는 [CHANGELOG-v1.7.md](../docs/CHANGELOG/CHANGELOG-v1.7.md)를 참조하세요.

## Support / 지원

- GitHub Issues: https://github.com/arkd0ng/go-utils/issues
- Documentation: https://github.com/arkd0ng/go-utils/tree/main/sliceutil

---

**Made with ❤️ by arkd0ng**

**Part of [go-utils](https://github.com/arkd0ng/go-utils) - A collection of utility packages for Go**
