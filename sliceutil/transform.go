package sliceutil

// transform.go provides functional transformation operations for slices.
//
// This file implements core functional programming patterns for transforming,
// filtering, and restructuring slices:
//
// Transformation Operations:
//
// Map (Element-wise Transform):
//   - Map(slice, fn): Apply function to each element
//     Time: O(n), Space: O(n)
//     Creates new slice with transformed elements
//     Type-safe: Can transform T → R (different types)
//     Pure function: Original slice unchanged
//     Example: [1,2,3] with n*2 → [2,4,6]
//     Use cases: Data transformation, type conversion, calculations
//
// Filter (Selective Inclusion):
//   - Filter(slice, predicate): Keep elements matching condition
//     Time: O(n), Space: O(k) where k = matching elements
//     Result slice contains only elements where predicate returns true
//     Maintains original order
//     Efficient allocation: Pre-allocates with capacity len(slice)
//     Example: [1,2,3,4,5,6] with even → [2,4,6]
//     Use cases: Data filtering, validation, subset extraction
//
// FlatMap (Map + Flatten):
//   - FlatMap(slice, fn): Map then flatten results
//     Time: O(n*m) where m = avg result length, Space: O(n*m)
//     Each element can produce multiple results
//     All results concatenated into single slice
//     Example: ["ab","cd"] with chars → ['a','b','c','d']
//     Use cases: Nested data expansion, one-to-many transformations
//
// Flatten (Nest Reduction):
//   - Flatten(slice): Convert [][]T to []T
//     Time: O(n*m) where m = avg inner length, Space: O(n*m)
//     Concatenates all inner slices
//     Single-level flattening only (not recursive)
//     Example: [[1,2],[3,4],[5]] → [1,2,3,4,5]
//     Use cases: Batch result merging, nested data simplification
//
// Unique (Duplicate Removal):
//   - Unique(slice): Remove duplicate elements
//     Time: O(n), Space: O(n)
//     Preserves first occurrence, removes subsequent duplicates
//     Maintains original order
//     Uses map for O(1) lookups
//     Requires comparable type
//     Example: [1,2,2,3,1,4] → [1,2,3,4]
//     Use cases: Deduplication, set conversion
//
// Reverse (Order Inversion):
//   - Reverse(slice): Reverse element order
//     Time: O(n), Space: O(n)
//     Creates new slice in reverse order
//     Original slice unchanged
//     Example: [1,2,3,4,5] → [5,4,3,2,1]
//     Use cases: Stack operations, reverse chronological order
//
// Compact (Remove Consecutive Duplicates):
//   - Compact(slice): Remove adjacent duplicates
//     Time: O(n), Space: O(k) where k = unique sequences
//     Only removes consecutive duplicates, not all
//     Maintains order and one of each sequence
//     Example: [1,1,2,2,2,3,1,1] → [1,2,3,1]
//     Use cases: Run-length encoding prep, sequence simplification
//
// MapIndex (Indexed Transform):
//   - MapIndex(slice, fn): Apply function with index access
//     Time: O(n), Space: O(n)
//     Function receives both element and its index
//     Useful when transformation depends on position
//     Example: ["a","b","c"] with (i,v) → v+i → ["a0","b1","c2"]
//     Use cases: Enumeration, position-aware transformations
//
// Design Principles:
//   - Immutability: All functions return new slices
//   - Composition: Functions can be chained
//   - Type safety: Generic implementations for any type
//   - Predictable allocation: Pre-allocate when size known
//   - Nil handling: nil input produces nil output
//
// Performance Characteristics:
//   - Map/MapIndex: Always allocates full output size
//   - Filter: Allocates capacity for worst case, actual size may be smaller
//   - FlatMap/Flatten: Cannot predict size, grows dynamically
//   - Unique: Requires hash map for seen elements
//   - Reverse: Simple allocation and copy
//   - Compact: Dynamic allocation based on unique sequences
//
// Memory Efficiency:
//   - Map: Same size as input (len(result) = len(input))
//   - Filter: Same or smaller (0 <= len(result) <= len(input))
//   - FlatMap: Typically larger (len(result) >= len(input))
//   - Unique: Same or smaller, plus map overhead
//   - Reverse: Same size as input
//   - Compact: Same or smaller (len(result) <= len(input))
//
// Common Usage Patterns:
//
//	// Transform user data
//	userIDs := sliceutil.Map(users, func(u User) int {
//	    return u.ID
//	})
//
//	// Filter valid items
//	active := sliceutil.Filter(items, func(i Item) bool {
//	    return i.Status == "active"
//	})
//
//	// Expand nested structures
//	allTags := sliceutil.FlatMap(articles, func(a Article) []string {
//	    return a.Tags
//	})
//
//	// Remove duplicates from search results
//	unique := sliceutil.Unique(searchResults)
//
//	// Chain operations (composability)
//	result := sliceutil.Map(
//	    sliceutil.Filter(
//	        sliceutil.Unique(data),
//	        isValid,
//	    ),
//	    transform,
//	)
//
// Comparison with Standard Library:
//   - More concise than for loops
//   - Type-safe with generics (vs interface{})
//   - Functional composition style
//   - Predictable memory allocation
//
// transform.go는 슬라이스에 대한 함수형 변환 작업을 제공합니다.
//
// 이 파일은 슬라이스를 변환, 필터링 및 재구성하기 위한 핵심 함수형
// 프로그래밍 패턴을 구현합니다:
//
// 변환 작업:
//
// Map (요소별 변환):
//   - Map(slice, fn): 각 요소에 함수 적용
//     시간: O(n), 공간: O(n)
//     변환된 요소로 새 슬라이스 생성
//     타입 안전: T → R (다른 타입)로 변환 가능
//     순수 함수: 원본 슬라이스 변경 안 함
//     예: [1,2,3] n*2로 → [2,4,6]
//     사용 사례: 데이터 변환, 타입 변환, 계산
//
// Filter (선택적 포함):
//   - Filter(slice, predicate): 조건과 일치하는 요소 유지
//     시간: O(n), 공간: O(k) (k = 일치하는 요소)
//     결과 슬라이스는 조건자가 true를 반환하는 요소만 포함
//     원래 순서 유지
//     효율적인 할당: len(slice) 용량으로 사전 할당
//     예: [1,2,3,4,5,6] 짝수로 → [2,4,6]
//     사용 사례: 데이터 필터링, 검증, 부분 집합 추출
//
// FlatMap (Map + Flatten):
//   - FlatMap(slice, fn): 매핑 후 결과 평탄화
//     시간: O(n*m) (m = 평균 결과 길이), 공간: O(n*m)
//     각 요소가 여러 결과 생성 가능
//     모든 결과를 단일 슬라이스로 연결
//     예: ["ab","cd"] 문자로 → ['a','b','c','d']
//     사용 사례: 중첩 데이터 확장, 일대다 변환
//
// Flatten (중첩 축소):
//   - Flatten(slice): [][]T를 []T로 변환
//     시간: O(n*m) (m = 평균 내부 길이), 공간: O(n*m)
//     모든 내부 슬라이스 연결
//     단일 레벨 평탄화만 (재귀 아님)
//     예: [[1,2],[3,4],[5]] → [1,2,3,4,5]
//     사용 사례: 배치 결과 병합, 중첩 데이터 단순화
//
// Unique (중복 제거):
//   - Unique(slice): 중복 요소 제거
//     시간: O(n), 공간: O(n)
//     첫 번째 발생 유지, 후속 중복 제거
//     원래 순서 유지
//     O(1) 조회를 위해 맵 사용
//     comparable 타입 필요
//     예: [1,2,2,3,1,4] → [1,2,3,4]
//     사용 사례: 중복 제거, 집합 변환
//
// Reverse (순서 반전):
//   - Reverse(slice): 요소 순서 반전
//     시간: O(n), 공간: O(n)
//     역순으로 새 슬라이스 생성
//     원본 슬라이스 변경 안 함
//     예: [1,2,3,4,5] → [5,4,3,2,1]
//     사용 사례: 스택 작업, 역시간순
//
// Compact (연속 중복 제거):
//   - Compact(slice): 인접 중복 제거
//     시간: O(n), 공간: O(k) (k = 고유 시퀀스)
//     연속 중복만 제거, 모든 중복 아님
//     순서와 각 시퀀스 중 하나 유지
//     예: [1,1,2,2,2,3,1,1] → [1,2,3,1]
//     사용 사례: 런 길이 인코딩 준비, 시퀀스 단순화
//
// MapIndex (인덱스 변환):
//   - MapIndex(slice, fn): 인덱스 접근과 함께 함수 적용
//     시간: O(n), 공간: O(n)
//     함수가 요소와 인덱스 모두 수신
//     변환이 위치에 따라 달라질 때 유용
//     예: ["a","b","c"] (i,v) → v+i → ["a0","b1","c2"]
//     사용 사례: 열거, 위치 인식 변환
//
// 설계 원칙:
//   - 불변성: 모든 함수는 새 슬라이스 반환
//   - 구성: 함수 체인 가능
//   - 타입 안전성: 모든 타입에 대한 제네릭 구현
//   - 예측 가능한 할당: 크기를 알 때 사전 할당
//   - Nil 처리: nil 입력은 nil 출력 생성
//
// 성능 특성:
//   - Map/MapIndex: 항상 전체 출력 크기 할당
//   - Filter: 최악의 경우 용량 할당, 실제 크기는 더 작을 수 있음
//   - FlatMap/Flatten: 크기 예측 불가, 동적으로 성장
//   - Unique: 본 요소를 위한 해시 맵 필요
//   - Reverse: 간단한 할당 및 복사
//   - Compact: 고유 시퀀스에 따라 동적 할당
//
// 메모리 효율성:
//   - Map: 입력과 같은 크기 (len(result) = len(input))
//   - Filter: 같거나 작음 (0 <= len(result) <= len(input))
//   - FlatMap: 일반적으로 더 큼 (len(result) >= len(input))
//   - Unique: 같거나 작음, 맵 오버헤드 추가
//   - Reverse: 입력과 같은 크기
//   - Compact: 같거나 작음 (len(result) <= len(input))
//
// 일반적인 사용 패턴:
//
//	// 사용자 데이터 변환
//	userIDs := sliceutil.Map(users, func(u User) int {
//	    return u.ID
//	})
//
//	// 유효한 항목 필터링
//	active := sliceutil.Filter(items, func(i Item) bool {
//	    return i.Status == "active"
//	})
//
//	// 중첩 구조 확장
//	allTags := sliceutil.FlatMap(articles, func(a Article) []string {
//	    return a.Tags
//	})
//
//	// 검색 결과에서 중복 제거
//	unique := sliceutil.Unique(searchResults)
//
//	// 작업 체인 (구성성)
//	result := sliceutil.Map(
//	    sliceutil.Filter(
//	        sliceutil.Unique(data),
//	        isValid,
//	    ),
//	    transform,
//	)
//
// 표준 라이브러리와 비교:
//   - for 루프보다 간결
//   - 제네릭으로 타입 안전 (vs interface{})
//   - 함수형 구성 스타일
//   - 예측 가능한 메모리 할당

// Map applies a function to each element and returns a new slice with the results.
// Map은 각 요소에 함수를 적용하고 결과로 새 슬라이스를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	doubled := sliceutil.Map(numbers, func(n int) int {
//	    return n * 2
//	}) // [2, 4, 6, 8, 10]
//
//	words := []string{"hello", "world"}
//	lengths := sliceutil.Map(words, func(s string) int {
//	    return len(s)
//	}) // [5, 5]
func Map[T any, R any](slice []T, fn func(T) R) []R {
	if slice == nil {
		return nil
	}
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter returns a new slice containing only elements that satisfy the predicate.
// Filter는 조건을 만족하는 요소만 포함하는 새 슬라이스를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens := sliceutil.Filter(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // [2, 4, 6]
//
//	words := []string{"apple", "a", "banana", "pear", "ab"}
//	long := sliceutil.Filter(words, func(s string) bool {
//	    return len(s) > 2
//	}) // ["apple", "banana", "pear"]
func Filter[T any](slice []T, predicate func(T) bool) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// FlatMap applies a function to each element and flattens the results into a single slice.
// FlatMap은 각 요소에 함수를 적용하고 결과를 하나의 슬라이스로 평탄화합니다.
//
// Example
// 예제:
//
//	words := []string{"hello", "world"}
//	chars := sliceutil.FlatMap(words, func(s string) []rune {
//	    return []rune(s)
//	}) // ['h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd']
//
//	numbers := []int{1, 2, 3}
//	pairs := sliceutil.FlatMap(numbers, func(n int) []int {
//	    return []int{n, n * 2}
//	}) // [1, 2, 2, 4, 3, 6]
func FlatMap[T any, R any](slice []T, fn func(T) []R) []R {
	if slice == nil {
		return nil
	}
	result := make([]R, 0, len(slice))
	for _, v := range slice {
		result = append(result, fn(v)...)
	}
	return result
}

// Flatten flattens a slice of slices into a single slice.
// Flatten은 슬라이스의 슬라이스를 하나의 슬라이스로 평탄화합니다.
//
// Example
// 예제:
//
//	nested := [][]int{{1, 2}, {3, 4}, {5}}
//	flat := sliceutil.Flatten(nested) // [1, 2, 3, 4, 5]
//
//	words := [][]string{{"hello", "world"}, {"foo", "bar"}}
//	flat := sliceutil.Flatten(words) // ["hello", "world", "foo", "bar"]
func Flatten[T any](slice [][]T) []T {
	if slice == nil {
		return nil
	}
	// Calculate total length for pre-allocation
	totalLen := 0
	for _, sub := range slice {
		totalLen += len(sub)
	}
	result := make([]T, 0, totalLen)
	for _, sub := range slice {
		result = append(result, sub...)
	}
	return result
}

// Unique returns a new slice with duplicate elements removed.
// Unique는 중복 요소가 제거된 새 슬라이스를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4}
//	unique := sliceutil.Unique(numbers) // [1, 2, 3, 4]
//
//	words := []string{"apple", "banana", "apple", "cherry"}
//	unique := sliceutil.Unique(words) // ["apple", "banana", "cherry"]
func Unique[T comparable](slice []T) []T {
	if slice == nil {
		return nil
	}
	seen := make(map[T]bool, len(slice))
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// UniqueBy returns a new slice with duplicate elements removed based on a key function.
// UniqueBy는 키 함수를 기반으로 중복 요소가 제거된 새 슬라이스를 반환합니다.
//
// Example
// 예제:
//
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	people := []Person{
//	    {"Alice", 25},
//	    {"Bob", 30},
//	    {"Alice", 28},
//	}
//	unique := sliceutil.UniqueBy(people, func(p Person) string {
//	    return p.Name
//	}) // [{"Alice", 25}, {"Bob", 30}]
func UniqueBy[T any, K comparable](slice []T, keyFunc func(T) K) []T {
	if slice == nil {
		return nil
	}
	seen := make(map[K]bool, len(slice))
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		key := keyFunc(v)
		if !seen[key] {
			seen[key] = true
			result = append(result, v)
		}
	}
	return result
}

// Compact removes consecutive duplicate elements from the slice.
// Compact는 슬라이스에서 연속된 중복 요소를 제거합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
//	compact := sliceutil.Compact(numbers) // [1, 2, 3, 4, 5]
//
//	words := []string{"a", "a", "b", "b", "c"}
//	compact := sliceutil.Compact(words) // ["a", "b", "c"]
func Compact[T comparable](slice []T) []T {
	if slice == nil {
		return nil
	}
	if len(slice) == 0 {
		return []T{}
	}
	result := make([]T, 0, len(slice))
	result = append(result, slice[0])
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			result = append(result, slice[i])
		}
	}
	return result
}

// Reverse returns a new slice with elements in reverse order.
// Reverse는 요소가 역순으로 된 새 슬라이스를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	reversed := sliceutil.Reverse(numbers) // [5, 4, 3, 2, 1]
//
//	words := []string{"hello", "world"}
//	reversed := sliceutil.Reverse(words) // ["world", "hello"]
func Reverse[T any](slice []T) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}
