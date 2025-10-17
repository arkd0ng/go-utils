package sliceutil

// basic.go provides fundamental slice operations for searching, finding, and counting elements.
//
// This file implements the most commonly used slice operations that serve as building blocks
// for more complex slice manipulations. All functions are generic and type-safe using Go 1.18+ generics.
//
// Categories of Operations:
//
// Element Existence:
//   - Contains(slice, item): Check if exact item exists
//     Time: O(n), Space: O(1)
//     Returns boolean indicating presence
//   - ContainsFunc(slice, predicate): Check if any item satisfies condition
//     Time: O(n), Space: O(1)
//     Flexible predicate-based search
//
// Index Finding:
//   - IndexOf(slice, item): Find first occurrence index
//     Time: O(n), Space: O(1)
//     Returns -1 if not found
//   - LastIndexOf(slice, item): Find last occurrence index
//     Time: O(n), Space: O(1)
//     Reverse search from end
//
// Element Finding:
//   - Find(slice, predicate): Get first element matching condition
//     Time: O(n), Space: O(1)
//     Returns (element, true) or (zero, false)
//   - FindLast(slice, predicate): Get last element matching condition
//     Time: O(n), Space: O(1)
//     Reverse search with predicate
//   - FindIndex(slice, predicate): Get index of first match
//     Time: O(n), Space: O(1)
//     Returns -1 if not found
//
// Counting:
//   - Count(slice, item): Count exact occurrences
//     Time: O(n), Space: O(1)
//     Returns number of matches
//   - CountFunc(slice, predicate): Count elements satisfying condition
//     Time: O(n), Space: O(1)
//     Flexible predicate-based counting
//
// Position Retrieval:
//   - First(slice): Get first element
//     Time: O(1), Space: O(1)
//     Returns (element, true) or (zero, false) for empty slice
//   - Last(slice): Get last element
//     Time: O(1), Space: O(1)
//     Safe access to last element
//
// Design Principles:
//   - All functions are non-mutating (return results, don't modify input)
//   - Consistent error handling via boolean returns or -1 for not found
//   - Generic implementations work with any comparable/any type
//   - Linear time complexity for search operations (unavoidable for unsorted slices)
//   - Predicate functions allow flexible custom matching logic
//
// Nil and Empty Slice Handling:
//   - All functions safely handle nil slices (treated as empty)
//   - Empty slices return false for Contains, -1 for IndexOf, etc.
//   - No panics on nil/empty input
//
// Performance Characteristics:
//   - Contains/Find: Early termination on first match (best case O(1), worst O(n))
//   - LastIndexOf/FindLast: Must scan entire slice (always O(n))
//   - Count: Must scan entire slice (always O(n))
//   - First/Last: Constant time O(1) index access
//
// Common Usage Patterns:
//
//	// Check existence
//	if sliceutil.Contains(userIDs, targetID) {
//	    // User exists
//	}
//
//	// Find complex objects
//	user, found := sliceutil.Find(users, func(u User) bool {
//	    return u.Email == "john@example.com"
//	})
//
//	// Count occurrences
//	errorCount := sliceutil.CountFunc(logs, func(log Log) bool {
//	    return log.Level == "ERROR"
//	})
//
//	// Safe access
//	if first, ok := sliceutil.First(items); ok {
//	    // Process first item
//	}
//
// basic.go는 요소 검색, 찾기, 개수 세기를 위한 기본 슬라이스 작업을 제공합니다.
//
// 이 파일은 더 복잡한 슬라이스 조작을 위한 빌딩 블록 역할을 하는 가장 일반적으로
// 사용되는 슬라이스 작업을 구현합니다. 모든 함수는 Go 1.18+ 제네릭을 사용하여
// 제네릭하고 타입 안전합니다.
//
// 작업 카테고리:
//
// 요소 존재:
//   - Contains(slice, item): 정확한 항목 존재 확인
//     시간: O(n), 공간: O(1)
//     존재 여부를 나타내는 부울 반환
//   - ContainsFunc(slice, predicate): 조건을 만족하는 항목이 있는지 확인
//     시간: O(n), 공간: O(1)
//     유연한 조건자 기반 검색
//
// 인덱스 찾기:
//   - IndexOf(slice, item): 첫 번째 발생 인덱스 찾기
//     시간: O(n), 공간: O(1)
//     찾지 못하면 -1 반환
//   - LastIndexOf(slice, item): 마지막 발생 인덱스 찾기
//     시간: O(n), 공간: O(1)
//     끝에서 역방향 검색
//
// 요소 찾기:
//   - Find(slice, predicate): 조건과 일치하는 첫 번째 요소 가져오기
//     시간: O(n), 공간: O(1)
//     (요소, true) 또는 (제로, false) 반환
//   - FindLast(slice, predicate): 조건과 일치하는 마지막 요소 가져오기
//     시간: O(n), 공간: O(1)
//     조건자와 함께 역방향 검색
//   - FindIndex(slice, predicate): 첫 번째 일치 인덱스 가져오기
//     시간: O(n), 공간: O(1)
//     찾지 못하면 -1 반환
//
// 개수 세기:
//   - Count(slice, item): 정확한 발생 횟수 세기
//     시간: O(n), 공간: O(1)
//     일치 수 반환
//   - CountFunc(slice, predicate): 조건을 만족하는 요소 개수 세기
//     시간: O(n), 공간: O(1)
//     유연한 조건자 기반 개수 세기
//
// 위치 검색:
//   - First(slice): 첫 번째 요소 가져오기
//     시간: O(1), 공간: O(1)
//     빈 슬라이스의 경우 (요소, true) 또는 (제로, false) 반환
//   - Last(slice): 마지막 요소 가져오기
//     시간: O(1), 공간: O(1)
//     마지막 요소에 안전하게 접근
//
// 설계 원칙:
//   - 모든 함수는 비변이성 (결과 반환, 입력 수정 안 함)
//   - 부울 반환 또는 찾지 못한 경우 -1을 통한 일관된 에러 처리
//   - 제네릭 구현은 모든 comparable/any 타입과 작동
//   - 검색 작업의 선형 시간 복잡도 (정렬되지 않은 슬라이스의 경우 불가피)
//   - 조건자 함수는 유연한 커스텀 매칭 로직 허용
//
// Nil 및 빈 슬라이스 처리:
//   - 모든 함수는 nil 슬라이스를 안전하게 처리 (빈 것으로 취급)
//   - 빈 슬라이스는 Contains에 대해 false, IndexOf에 대해 -1 등을 반환
//   - nil/빈 입력에 대한 패닉 없음
//
// 성능 특성:
//   - Contains/Find: 첫 번째 일치 시 조기 종료 (최선 O(1), 최악 O(n))
//   - LastIndexOf/FindLast: 전체 슬라이스 스캔 필요 (항상 O(n))
//   - Count: 전체 슬라이스 스캔 필요 (항상 O(n))
//   - First/Last: 상수 시간 O(1) 인덱스 접근
//
// 일반적인 사용 패턴:
//
//	// 존재 확인
//	if sliceutil.Contains(userIDs, targetID) {
//	    // 사용자 존재
//	}
//
//	// 복잡한 객체 찾기
//	user, found := sliceutil.Find(users, func(u User) bool {
//	    return u.Email == "john@example.com"
//	})
//
//	// 발생 횟수 세기
//	errorCount := sliceutil.CountFunc(logs, func(log Log) bool {
//	    return log.Level == "ERROR"
//	})
//
//	// 안전한 접근
//	if first, ok := sliceutil.First(items); ok {
//	    // 첫 번째 항목 처리
//	}

// Contains checks if slice contains the specified item.
// Contains는 슬라이스에 지정된 항목이 있는지 확인합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasThree := sliceutil.Contains(numbers, 3) // true
//	hasTen := sliceutil.Contains(numbers, 10)  // false
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// ContainsFunc checks if slice contains an item that satisfies the predicate.
// ContainsFunc는 조건을 만족하는 항목이 슬라이스에 있는지 확인합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasEven := sliceutil.ContainsFunc(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // true (has 2, 4)
func ContainsFunc[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of item in slice.
// Returns -1 if item is not found.
//
// IndexOf는 슬라이스에서 항목의 첫 번째 발생 인덱스를 반환합니다.
// 항목을 찾을 수 없으면 -1을 반환합니다.
//
// Example
// 예제:
//
//	fruits := []string{"apple", "banana", "cherry", "banana"}
//	index := sliceutil.IndexOf(fruits, "banana") // 1
//	index2 := sliceutil.IndexOf(fruits, "grape") // -1
func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of item in slice.
// Returns -1 if item is not found.
//
// LastIndexOf는 슬라이스에서 항목의 마지막 발생 인덱스를 반환합니다.
// 항목을 찾을 수 없으면 -1을 반환합니다.
//
// Example
// 예제:
//
//	fruits := []string{"apple", "banana", "cherry", "banana"}
//	index := sliceutil.LastIndexOf(fruits, "banana") // 3
//	index2 := sliceutil.LastIndexOf(fruits, "grape") // -1
func LastIndexOf[T comparable](slice []T, item T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// Find returns the first item in slice that satisfies the predicate.
// Returns the found item and true if found, zero value and false otherwise.
//
// Find는 조건을 만족하는 슬라이스의 첫 번째 항목을 반환합니다.
// 찾은 경우 항목과 true를 반환하고, 그렇지 않으면 제로 값과 false를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	even, found := sliceutil.Find(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // even = 2, found = true
//
//	negative, found := sliceutil.Find(numbers, func(n int) bool {
//	    return n < 0
//	}) // negative = 0, found = false
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, v := range slice {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindLast returns the last item in slice that satisfies the predicate.
// Returns the found item and true if found, zero value and false otherwise.
//
// FindLast는 조건을 만족하는 슬라이스의 마지막 항목을 반환합니다.
// 찾은 경우 항목과 true를 반환하고, 그렇지 않으면 제로 값과 false를 반환합니다.
//
// Similar to Find, but searches from right to left.
// Find와 유사하지만 오른쪽에서 왼쪽으로 검색합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	even, found := sliceutil.FindLast(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // even = 6, found = true (last even number)
//
//	words := []string{"apple", "banana", "apricot", "cherry"}
//	startsWithA, found := sliceutil.FindLast(words, func(s string) bool {
//	    return len(s) > 0 && s[0] == 'a'
//	}) // startsWithA = "apricot", found = true
func FindLast[T any](slice []T, predicate func(T) bool) (T, bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return slice[i], true
		}
	}
	var zero T
	return zero, false
}

// FindIndex returns the index of the first item that satisfies the predicate.
// Returns -1 if no item is found.
//
// FindIndex는 조건을 만족하는 첫 번째 항목의 인덱스를 반환합니다.
// 항목을 찾을 수 없으면 -1을 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	index := sliceutil.FindIndex(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // 1 (index of 2)
func FindIndex[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// Count returns the number of items that satisfy the predicate.
// Count는 조건을 만족하는 항목의 개수를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evenCount := sliceutil.Count(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // 3 (2, 4, 6)
func Count[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
		}
	}
	return count
}

// IsEmpty checks if the slice is empty or nil.
// IsEmpty는 슬라이스가 비어있거나 nil인지 확인합니다.
//
// Example
// 예제:
//
//	empty := []int{}
//	sliceutil.IsEmpty(empty) // true
//	sliceutil.IsEmpty(nil)   // true
//
//	nonEmpty := []int{1, 2, 3}
//	sliceutil.IsEmpty(nonEmpty) // false
func IsEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// IsNotEmpty checks if the slice has at least one item.
// IsNotEmpty는 슬라이스에 최소한 하나의 항목이 있는지 확인합니다.
//
// Example
// 예제:
//
//	empty := []int{}
//	sliceutil.IsNotEmpty(empty) // false
//
//	nonEmpty := []int{1, 2, 3}
//	sliceutil.IsNotEmpty(nonEmpty) // true
func IsNotEmpty[T any](slice []T) bool {
	return len(slice) > 0
}

// Equal checks if two slices are equal (same length and same elements in same order).
// Equal은 두 슬라이스가 같은지 확인합니다 (같은 길이와 같은 순서의 같은 요소).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3}
//	b := []int{1, 2, 3}
//	c := []int{1, 2, 4}
//
//	sliceutil.Equal(a, b) // true
//	sliceutil.Equal(a, c) // false
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
