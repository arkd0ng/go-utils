# maputil - Extreme Simplicity Map Utilities

**Reduce 20+ lines of repetitive map manipulation code to just 1-2 lines.**

**20줄 이상의 반복적인 맵 조작 코드를 단 1-2줄로 줄입니다.**

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview / 개요

The `maputil` package provides 81 type-safe functions across 10 categories to simplify common map operations in Go. Built with Go 1.18+ generics for compile-time type safety.

`maputil` 패키지는 10개 카테고리에 걸쳐 81개의 타입 안전 함수를 제공하여 Go의 일반적인 맵 작업을 단순화합니다. Go 1.18+ 제네릭으로 컴파일 타임 타입 안전성을 제공합니다.

## Features / 주요 기능

- ✅ **81 Functions** across 10 categories / 10개 카테고리에 걸쳐 81개 함수
- ✅ **Type-safe** with Go 1.18+ generics / Go 1.18+ 제네릭으로 타입 안전
- ✅ **Functional programming** style (Map, Filter, Reduce) / 함수형 프로그래밍 스타일
- ✅ **Immutable operations** (original maps unchanged) / 불변 작업 (원본 맵 변경 없음)
- ✅ **Zero external dependencies** / 외부 의존성 제로
- ✅ **100% test coverage** / 100% 테스트 커버리지
- ✅ **Bilingual documentation** (English/Korean) / 이중 언어 문서 (영문/한글)

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/maputil
```

## Quick Start / 빠른 시작

```go
package main

import (
	"fmt"
	"github.com/arkd0ng/go-utils/maputil"
)

func main() {
	// Filter map / 맵 필터링
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	filtered := maputil.Filter(data, func(k string, v int) bool {
		return v > 2
	})
	fmt.Println(filtered) // map[c:3 d:4]

	// Transform values / 값 변환
	doubled := maputil.MapValues(data, func(v int) int {
		return v * 2
	})
	fmt.Println(doubled) // map[a:2 b:4 c:6 d:8]

	// Merge maps / 맵 병합
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	merged := maputil.Merge(m1, m2)
	fmt.Println(merged) // map[a:1 b:3 c:4]

	// Group by key / 키로 그룹화
	type User struct {
		Name string
		City string
	}
	users := []User{
		{Name: "Alice", City: "Seoul"},
		{Name: "Bob", City: "Seoul"},
		{Name: "Charlie", City: "Busan"},
	}
	byCity := maputil.GroupBy[string, User, string](users, func(u User) string {
		return u.City
	})
	fmt.Println(byCity)
	// map[Seoul:[{Alice Seoul} {Bob Seoul}] Busan:[{Charlie Busan}]]
}
```

## Function Categories / 함수 카테고리

### 1. Basic Operations (11 functions)

**Core map operations / 핵심 맵 작업**

- `Get` - Get value with existence check / 존재 확인과 함께 값 가져오기
- `GetOr` - Get value with default / 기본값과 함께 값 가져오기
- `Set` - Set key-value pair / 키-값 쌍 설정
- `Delete` - Delete keys / 키 삭제
- `Has` - Check key existence / 키 존재 확인
- `IsEmpty` - Check if empty / 비어 있는지 확인
- `IsNotEmpty` - Check if not empty / 비어 있지 않은지 확인
- `Len` - Get length / 길이 가져오기
- `Clear` - Clear all entries / 모든 항목 제거
- `Clone` - Clone map / 맵 복제
- `Equal` - Check equality / 동등성 확인

### 2. Transformation (10 functions)

**Transform map structure and values / 맵 구조와 값 변환**

- `Map` - Transform values / 값 변환
- `MapKeys` - Transform keys / 키 변환
- `MapValues` - Transform values (value-only) / 값 변환 (값만)
- `MapEntries` - Transform both keys and values / 키와 값 모두 변환
- `Invert` - Swap keys and values / 키와 값 교환
- `Flatten` - Flatten nested maps / 중첩 맵 평면화
- `Unflatten` - Unflatten to nested map / 중첩 맵으로 펼치기
- `Chunk` - Split into chunks / 청크로 분할
- `Partition` - Partition by predicate / 조건으로 분할
- `Compact` - Remove zero values / zero 값 제거

### 3. Aggregation (9 functions)

**Aggregate and calculate statistics / 집계 및 통계 계산**

- `Reduce` - Reduce to single value / 단일 값으로 축소
- `Sum` - Sum all values / 모든 값 합산
- `Min` - Find minimum / 최소값 찾기
- `Max` - Find maximum / 최대값 찾기
- `MinBy` - Find minimum by function / 함수로 최소값 찾기
- `MaxBy` - Find maximum by function / 함수로 최대값 찾기
- `Average` - Calculate average / 평균 계산
- `GroupBy` - Group slice by key / 키로 슬라이스 그룹화
- `CountBy` - Count by key / 키별 개수 세기

### 4. Merge Operations (8 functions)

**Combine multiple maps / 여러 맵 결합**

- `Merge` - Merge maps (last wins) / 맵 병합 (마지막 우선)
- `MergeWith` - Merge with custom resolver / 사용자 정의 해결자로 병합
- `DeepMerge` - Deep merge nested maps / 중첩 맵 깊이 병합
- `Union` - Union (alias for Merge) / 합집합 (Merge 별칭)
- `Intersection` - Intersection (common keys) / 교집합 (공통 키)
- `Difference` - Difference / 차집합
- `SymmetricDifference` - Symmetric difference / 대칭 차집합
- `Assign` - Assign to target (mutating) / 대상에 할당 (변경)

### 5. Filter Operations (7 functions)

**Filter and select entries / 항목 필터링 및 선택**

- `Filter` - Filter by predicate / 조건으로 필터
- `FilterKeys` - Filter by keys / 키로 필터
- `FilterValues` - Filter by values / 값으로 필터
- `Omit` - Omit specified keys / 지정된 키 제외
- `Pick` - Pick specified keys / 지정된 키 선택
- `OmitBy` - Omit by predicate / 조건으로 제외
- `PickBy` - Pick by predicate / 조건으로 선택

### 6. Conversion (8 functions)

**Convert between maps and other types / 맵과 다른 타입 간 변환**

- `Keys` - Get all keys / 모든 키 가져오기
- `Values` - Get all values / 모든 값 가져오기
- `Entries` - Get key-value pairs / 키-값 쌍 가져오기
- `FromEntries` - Create from entries / 항목으로부터 생성
- `FromSlice` - Create from slice / 슬라이스로부터 생성
- `FromSliceBy` - Create from slice with transform / 변환과 함께 슬라이스로부터 생성
- `ToSlice` - Convert to slice / 슬라이스로 변환
- `ToJSON` - Convert to JSON string / JSON 문자열로 변환

### 7. Predicate Checks (7 functions)

**Check conditions on map entries / 맵 항목의 조건 확인**

- `Every` - Check if all match / 모두 일치하는지 확인
- `Some` - Check if any match / 하나라도 일치하는지 확인
- `None` - Check if none match / 아무것도 일치하지 않는지 확인
- `HasKey` - Check key existence / 키 존재 확인
- `HasValue` - Check value existence / 값 존재 확인
- `HasEntry` - Check key-value pair / 키-값 쌍 확인
- `IsSubset` - Check if subset / 부분집합인지 확인

### 8. Key Operations (8 functions)

**Manipulate map keys / 맵 키 조작**

- `KeysSlice` - Get keys as slice / 키를 슬라이스로
- `KeysSorted` - Get sorted keys / 정렬된 키 가져오기
- `KeysBy` - Get keys by predicate / 조건으로 키 가져오기
- `RenameKey` - Rename key / 키 이름 변경
- `RenameKeys` - Rename multiple keys / 여러 키 이름 변경
- `SwapKeys` - Swap two keys / 두 키 교환
- `FindKey` - Find key by predicate / 조건으로 키 찾기
- `FindKeys` - Find all keys by predicate / 조건으로 모든 키 찾기

### 9. Value Operations (7 functions)

**Manipulate map values / 맵 값 조작**

- `ValuesSlice` - Get values as slice / 값을 슬라이스로
- `ValuesSorted` - Get sorted values / 정렬된 값 가져오기
- `ValuesBy` - Get values by predicate / 조건으로 값 가져오기
- `UniqueValues` - Get unique values / 고유한 값 가져오기
- `FindValue` - Find value by predicate / 조건으로 값 찾기
- `ReplaceValue` - Replace value / 값 대체
- `UpdateValues` - Update all values / 모든 값 업데이트

### 10. Comparison (6 functions)

**Compare and diff maps / 맵 비교 및 차이**

- `EqualFunc` - Check equality with custom comparator / 사용자 정의 비교자로 동등성 확인
- `Diff` - Get differing entries / 차이 나는 항목 가져오기
- `DiffKeys` - Get differing keys / 차이 나는 키 가져오기
- `CommonKeys` - Get common keys / 공통 키 가져오기
- `AllKeys` - Get all unique keys / 모든 고유 키 가져오기
- `Compare` - Detailed comparison / 상세 비교

## Before vs After / 이전 vs 이후

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
// 7+ lines of boilerplate / 7줄 이상의 보일러플레이트

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
// 10+ lines / 10줄 이상
```

### After (maputil) / 이후 (maputil)

```go
// Filtering / 필터링
result := maputil.Filter(data, func(k string, v int) bool { return v > 2 })
// 1 line! / 1줄!

// Merging / 병합
merged := maputil.Merge(map1, map2)
// 1 line! / 1줄!
```

## Real-World Examples / 실제 사용 예제

### Example 1: User Management / 사용자 관리

```go
type User struct {
	ID    int
	Name  string
	Email string
	City  string
}

users := []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com", City: "Seoul"},
	{ID: 2, Name: "Bob", Email: "bob@example.com", City: "Seoul"},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com", City: "Busan"},
}

// Create a map of users by ID / ID로 사용자 맵 생성
userMap := maputil.FromSlice(users, func(u User) int {
	return u.ID
})

// Group users by city / 도시별로 사용자 그룹화
byCity := maputil.GroupBy[int, User, string](users, func(u User) string {
	return u.City
})

// Count users by city / 도시별 사용자 수 계산
counts := maputil.CountBy[int, User, string](users, func(u User) string {
	return u.City
})

fmt.Println(counts) // map[Seoul:2 Busan:1]
```

### Example 2: Configuration Management / 설정 관리

```go
// Default configuration / 기본 설정
defaults := map[string]string{
	"host":    "localhost",
	"port":    "8080",
	"timeout": "30s",
}

// User configuration / 사용자 설정
userConfig := map[string]string{
	"port":    "3000",
	"timeout": "60s",
}

// Merge with user config taking precedence / 사용자 설정이 우선하는 병합
config := maputil.Merge(defaults, userConfig)

fmt.Println(config)
// map[host:localhost port:3000 timeout:60s]
```

### Example 3: Data Transformation / 데이터 변환

```go
// API response / API 응답
data := map[string]interface{}{
	"user_id":    1,
	"user_name":  "Alice",
	"user_email": "alice@example.com",
}

// Convert keys to camelCase / 키를 camelCase로 변환
camelCase := maputil.MapKeys(data, func(k string, v interface{}) string {
	return strings.ReplaceAll(k, "_", "")
})

fmt.Println(camelCase)
// map[userid:1 username:Alice useremail:alice@example.com]
```

## Performance / 성능

All functions are optimized for performance:

모든 함수는 성능에 최적화되어 있습니다:

- Most operations are **O(n)** time complexity / 대부분의 작업은 **O(n)** 시간 복잡도
- Minimal memory allocations / 최소한의 메모리 할당
- Efficient algorithms / 효율적인 알고리즘

```bash
BenchmarkFilter-8       	  500000	      2847 ns/op
BenchmarkMerge-8        	  300000	      4521 ns/op
BenchmarkMap-8          	  500000	      3012 ns/op
```

## Documentation / 문서

- [Package Documentation](https://pkg.go.dev/github.com/arkd0ng/go-utils/maputil)
- [Examples](../examples/maputil/)
- [User Manual](../docs/maputil/USER_MANUAL.md) - Comprehensive guide (coming soon)
- [Developer Guide](../docs/maputil/DEVELOPER_GUIDE.md) - Technical documentation (coming soon)

## Testing / 테스트

```bash
# Run all tests / 모든 테스트 실행
go test ./maputil -v

# Run with coverage / 커버리지와 함께 실행
go test ./maputil -cover

# Run benchmarks / 벤치마크 실행
go test ./maputil -bench=.
```

## Contributing / 기여하기

Contributions are welcome! Please feel free to submit a Pull Request.

기여를 환영합니다! Pull Request를 제출해 주세요.

## License / 라이선스

MIT License - see [LICENSE](../LICENSE) file for details.

MIT 라이선스 - 자세한 내용은 [LICENSE](../LICENSE) 파일을 참조하세요.

## Version / 버전

Current version: **v1.8.001**

현재 버전: **v1.8.001**

## Author / 작성자

**arkd0ng**

- GitHub: [@arkd0ng](https://github.com/arkd0ng)

## See Also / 관련 패키지

- [sliceutil](../sliceutil/) - Slice utilities (95 functions) / 슬라이스 유틸리티 (95개 함수)
- [stringutil](../stringutil/) - String utilities (53 functions) / 문자열 유틸리티 (53개 함수)
- [timeutil](../timeutil/) - Time utilities (114 functions) / 시간 유틸리티 (114개 함수)
