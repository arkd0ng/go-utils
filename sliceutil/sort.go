package sliceutil

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// sort.go provides sorting and sort validation operations for slices.
//
// This file implements immutable sorting operations with multiple strategies:
// simple ordering, custom key extraction, and multi-criteria sorting.
//
// Basic Sorting Operations:
//
// Sort (Ascending Order):
//   - Sort(slice): Sort in ascending order (smallest to largest)
//     Time: O(n log n), Space: O(n)
//     Uses Go's standard sort.Slice (introsort algorithm)
//     Creates new slice (original unchanged)
//     Stable sort: Equal elements maintain relative order
//     Example: Sort([3,1,4,1,5]) → [1,1,3,4,5]
//     Use cases: Numeric ordering, alphabetical sorting, default ordering
//
// SortDesc (Descending Order):
//   - SortDesc(slice): Sort in descending order (largest to smallest)
//     Time: O(n log n), Space: O(n)
//     Same algorithm as Sort but reversed comparison
//     Creates new slice (original unchanged)
//     Stable sort: Equal elements maintain relative order
//     Example: SortDesc([3,1,4,1,5]) → [5,4,3,1,1]
//     Use cases: Top-N rankings, reverse chronological, priority ordering
//
// Custom Sorting Operations:
//
// SortBy (Key-Based Sorting):
//   - SortBy(slice, keyFunc): Sort by extracted key in ascending order
//     Time: O(n log n), Space: O(n)
//     Key function extracts comparable value from each element
//     Sorts by key, not entire element
//     Useful for structs: Sort people by age, products by price
//     Key function called O(n log n) times (once per comparison)
//     Example: SortBy(people, p → p.Age) sorts by age ascending
//     Use cases: Sorting complex types, property-based ordering
//
// SortByMulti (Multi-Criteria Sorting):
//   - SortByMulti(slice, less): Sort with custom comparison function
//     Time: O(n log n), Space: O(n)
//     Most flexible sorting option
//     Less function compares two elements directly
//     Supports multi-key sorting (primary, secondary, tertiary keys)
//     Can mix ascending/descending per criterion
//     Example: Sort by name (asc), then age (asc), then score (desc)
//     Use cases: Complex sorting logic, tie-breaking, database-like ORDER BY
//
// Sort Validation Operations:
//
// IsSorted (Ascending Check):
//   - IsSorted(slice): Check if already sorted in ascending order
//     Time: O(n) worst case, O(1) best case, Space: O(1)
//     Short-circuits on first out-of-order pair
//     Empty/single-element slices are sorted (vacuous truth)
//     Example: IsSorted([1,2,3,4,5]) = true
//     Use cases: Precondition checking, binary search validation, optimization
//
// IsSortedDesc (Descending Check):
//   - IsSortedDesc(slice): Check if already sorted in descending order
//     Time: O(n) worst case, O(1) best case, Space: O(1)
//     Short-circuits on first out-of-order pair
//     Empty/single-element slices are sorted
//     Example: IsSortedDesc([5,4,3,2,1]) = true
//     Use cases: Precondition checking, reverse order validation
//
// Comparison: Sorting Strategies:
//
// Sort vs SortDesc:
//   - Sort: Ascending (< comparison)
//   - SortDesc: Descending (> comparison)
//   - Same performance, just reversed order
//
// Sort vs SortBy:
//   - Sort: Direct element comparison (Ordered types only)
//   - SortBy: Indirect via key function (any element type)
//   - SortBy: More flexible but slightly more overhead
//   - SortBy: Can sort by derived/computed values
//
// SortBy vs SortByMulti:
//   - SortBy: Single criterion, ascending only
//   - SortByMulti: Multiple criteria, any comparison logic
//   - SortByMulti: Can implement tie-breaking
//   - SortByMulti: Can mix ascending/descending
//
// When to Use Each:
//   - Sort/SortDesc: Simple types (int, string, float) with default order
//   - SortBy: Structs sorted by single field
//   - SortByMulti: Complex sorting (multiple fields, custom logic)
//   - IsSorted: Verify preconditions, avoid redundant sorting
//
// Design Principles:
//   - Immutability: All operations create new slices
//   - Stability: Equal elements maintain relative order
//   - Performance: O(n log n) time, O(n) space (optimal)
//   - Flexibility: Multiple sorting strategies for different needs
//   - Type safety: Generic constraints ensure correctness
//
// Performance Characteristics:
//
// Time Complexity:
//   - All sorting operations: O(n log n) average and worst case
//   - Go uses introsort (quicksort + heapsort + insertion sort)
//   - Validation operations: O(n) worst case, O(1) best with short-circuit
//
// Space Complexity:
//   - All sorting operations: O(n) for result slice
//   - Validation operations: O(1) - no allocation
//
// Stability:
//   - All sorting operations use sort.Slice which is stable
//   - Equal elements maintain their original relative order
//   - Important for multi-pass sorts or preserving secondary ordering
//
// Key Function Overhead:
//   - SortBy: Key function called O(n log n) times
//   - For expensive key functions, consider precomputing keys
//   - For simple extractions (p.Age), overhead is negligible
//
// Memory Allocation:
//   - Sorting: Single allocation for result slice (len(slice) capacity)
//   - Cloning: Full deep copy of slice data
//   - For in-place sorting: Use sort.Slice directly on slice
//
// Common Usage Patterns:
//
//	// Simple numeric sorting
//	ascending := sliceutil.Sort(numbers)
//	descending := sliceutil.SortDesc(numbers)
//
//	// Sort structs by property
//	type Product struct { Name string; Price float64 }
//	byPrice := sliceutil.SortBy(products, func(p Product) float64 {
//	    return p.Price
//	})
//
//	// Multi-criteria sorting
//	type Student struct { Grade int; Name string; Score float64 }
//	sorted := sliceutil.SortByMulti(students, func(a, b Student) bool {
//	    if a.Grade != b.Grade {
//	        return a.Grade < b.Grade  // Primary: Grade ascending
//	    }
//	    if a.Score != b.Score {
//	        return a.Score > b.Score  // Secondary: Score descending
//	    }
//	    return a.Name < b.Name        // Tertiary: Name ascending
//	})
//
//	// Validate before binary search
//	if !sliceutil.IsSorted(data) {
//	    data = sliceutil.Sort(data)
//	}
//	index := sort.SearchInts(data, target)
//
// Optimization Patterns:
//
//	// Avoid redundant sorting
//	if !sliceutil.IsSorted(data) {
//	    data = sliceutil.Sort(data)
//	}
//
//	// Precompute expensive keys
//	type CachedItem struct {
//	    Item Item
//	    Key  int
//	}
//	cached := sliceutil.Map(items, func(i Item) CachedItem {
//	    return CachedItem{i, expensiveKeyFunc(i)}
//	})
//	sorted := sliceutil.SortBy(cached, func(c CachedItem) int {
//	    return c.Key
//	})
//	result := sliceutil.Map(sorted, func(c CachedItem) Item {
//	    return c.Item
//	})
//
// Comparison with Standard Library:
//   - sort.Slice: In-place, modifies original
//   - sliceutil.Sort: Immutable, creates new slice
//   - sort.Slice: Requires manual cloning for immutability
//   - sliceutil: More convenient for functional style
//
// Multi-Criteria Sorting Examples:
//
//	// Sort by multiple fields with different orders
//	sorted := sliceutil.SortByMulti(employees, func(a, b Employee) bool {
//	    // Primary: Department (ascending)
//	    if a.Department != b.Department {
//	        return a.Department < b.Department
//	    }
//	    // Secondary: Salary (descending)
//	    if a.Salary != b.Salary {
//	        return a.Salary > b.Salary
//	    }
//	    // Tertiary: Name (ascending)
//	    return a.Name < b.Name
//	})
//
// sort.go는 슬라이스에 대한 정렬 및 정렬 검증 작업을 제공합니다.
//
// 이 파일은 여러 전략을 사용한 불변 정렬 작업을 구현합니다:
// 단순 순서, 사용자 정의 키 추출, 다중 기준 정렬.
//
// 기본 정렬 작업:
//
// Sort (오름차순):
//   - Sort(slice): 오름차순 정렬 (가장 작은 것부터 큰 것)
//     시간: O(n log n), 공간: O(n)
//     Go의 표준 sort.Slice 사용 (introsort 알고리즘)
//     새 슬라이스 생성 (원본 변경 안 함)
//     안정 정렬: 같은 요소는 상대적 순서 유지
//     예: Sort([3,1,4,1,5]) → [1,1,3,4,5]
//     사용 사례: 숫자 순서, 알파벳 정렬, 기본 순서
//
// SortDesc (내림차순):
//   - SortDesc(slice): 내림차순 정렬 (가장 큰 것부터 작은 것)
//     시간: O(n log n), 공간: O(n)
//     Sort와 같은 알고리즘이지만 비교 반전
//     새 슬라이스 생성 (원본 변경 안 함)
//     안정 정렬: 같은 요소는 상대적 순서 유지
//     예: SortDesc([3,1,4,1,5]) → [5,4,3,1,1]
//     사용 사례: Top-N 순위, 역시간순, 우선순위 순서
//
// 사용자 정의 정렬 작업:
//
// SortBy (키 기반 정렬):
//   - SortBy(slice, keyFunc): 추출된 키로 오름차순 정렬
//     시간: O(n log n), 공간: O(n)
//     키 함수가 각 요소에서 comparable 값 추출
//     전체 요소가 아닌 키로 정렬
//     구조체 유용: 나이로 사람 정렬, 가격으로 제품 정렬
//     키 함수는 O(n log n)번 호출 (비교당 한 번)
//     예: SortBy(people, p → p.Age) 나이로 오름차순 정렬
//     사용 사례: 복잡한 타입 정렬, 속성 기반 순서
//
// SortByMulti (다중 기준 정렬):
//   - SortByMulti(slice, less): 사용자 정의 비교 함수로 정렬
//     시간: O(n log n), 공간: O(n)
//     가장 유연한 정렬 옵션
//     Less 함수가 두 요소를 직접 비교
//     다중 키 정렬 지원 (1차, 2차, 3차 키)
//     기준별로 오름차순/내림차순 혼합 가능
//     예: 이름(오름차순), 나이(오름차순), 점수(내림차순) 정렬
//     사용 사례: 복잡한 정렬 논리, 동점 처리, 데이터베이스 같은 ORDER BY
//
// 정렬 검증 작업:
//
// IsSorted (오름차순 확인):
//   - IsSorted(slice): 이미 오름차순으로 정렬되었는지 확인
//     시간: O(n) 최악, O(1) 최선, 공간: O(1)
//     첫 번째 순서 위반 쌍에서 단락
//     빈/단일 요소 슬라이스는 정렬됨 (공허한 진리)
//     예: IsSorted([1,2,3,4,5]) = true
//     사용 사례: 전제 조건 확인, 이진 검색 검증, 최적화
//
// IsSortedDesc (내림차순 확인):
//   - IsSortedDesc(slice): 이미 내림차순으로 정렬되었는지 확인
//     시간: O(n) 최악, O(1) 최선, 공간: O(1)
//     첫 번째 순서 위반 쌍에서 단락
//     빈/단일 요소 슬라이스는 정렬됨
//     예: IsSortedDesc([5,4,3,2,1]) = true
//     사용 사례: 전제 조건 확인, 역순 검증
//
// 비교: 정렬 전략:
//
// Sort vs SortDesc:
//   - Sort: 오름차순 (< 비교)
//   - SortDesc: 내림차순 (> 비교)
//   - 성능 같음, 순서만 반전
//
// Sort vs SortBy:
//   - Sort: 직접 요소 비교 (Ordered 타입만)
//   - SortBy: 키 함수를 통한 간접 (모든 요소 타입)
//   - SortBy: 더 유연하지만 약간 더 많은 오버헤드
//   - SortBy: 파생/계산된 값으로 정렬 가능
//
// SortBy vs SortByMulti:
//   - SortBy: 단일 기준, 오름차순만
//   - SortByMulti: 다중 기준, 모든 비교 논리
//   - SortByMulti: 동점 처리 구현 가능
//   - SortByMulti: 오름차순/내림차순 혼합 가능
//
// 각각 사용 시기:
//   - Sort/SortDesc: 기본 순서의 단순 타입 (int, string, float)
//   - SortBy: 단일 필드로 정렬되는 구조체
//   - SortByMulti: 복잡한 정렬 (여러 필드, 사용자 정의 논리)
//   - IsSorted: 전제 조건 확인, 중복 정렬 방지
//
// 설계 원칙:
//   - 불변성: 모든 작업이 새 슬라이스 생성
//   - 안정성: 같은 요소는 상대적 순서 유지
//   - 성능: O(n log n) 시간, O(n) 공간 (최적)
//   - 유연성: 다른 요구에 대한 여러 정렬 전략
//   - 타입 안전성: 제네릭 제약 조건으로 정확성 보장
//
// 성능 특성:
//
// 시간 복잡도:
//   - 모든 정렬 작업: 평균 및 최악 O(n log n)
//   - Go는 introsort 사용 (quicksort + heapsort + insertion sort)
//   - 검증 작업: 최악 O(n), 단락으로 최선 O(1)
//
// 공간 복잡도:
//   - 모든 정렬 작업: 결과 슬라이스를 위한 O(n)
//   - 검증 작업: O(1) - 할당 없음
//
// 안정성:
//   - 모든 정렬 작업은 안정적인 sort.Slice 사용
//   - 같은 요소는 원래 상대적 순서 유지
//   - 다중 패스 정렬이나 2차 순서 보존에 중요
//
// 키 함수 오버헤드:
//   - SortBy: 키 함수 O(n log n)번 호출
//   - 비싼 키 함수는 키 사전 계산 고려
//   - 단순 추출 (p.Age)은 오버헤드 무시 가능
//
// 메모리 할당:
//   - 정렬: 결과 슬라이스를 위한 단일 할당 (len(slice) 용량)
//   - 복제: 슬라이스 데이터의 전체 깊은 복사
//   - 제자리 정렬: 슬라이스에서 직접 sort.Slice 사용
//
// 일반적인 사용 패턴:
//
//	// 단순 숫자 정렬
//	ascending := sliceutil.Sort(numbers)
//	descending := sliceutil.SortDesc(numbers)
//
//	// 속성으로 구조체 정렬
//	type Product struct { Name string; Price float64 }
//	byPrice := sliceutil.SortBy(products, func(p Product) float64 {
//	    return p.Price
//	})
//
//	// 다중 기준 정렬
//	type Student struct { Grade int; Name string; Score float64 }
//	sorted := sliceutil.SortByMulti(students, func(a, b Student) bool {
//	    if a.Grade != b.Grade {
//	        return a.Grade < b.Grade  // 1차: 학년 오름차순
//	    }
//	    if a.Score != b.Score {
//	        return a.Score > b.Score  // 2차: 점수 내림차순
//	    }
//	    return a.Name < b.Name        // 3차: 이름 오름차순
//	})
//
//	// 이진 검색 전 검증
//	if !sliceutil.IsSorted(data) {
//	    data = sliceutil.Sort(data)
//	}
//	index := sort.SearchInts(data, target)
//
// 최적화 패턴:
//
//	// 중복 정렬 방지
//	if !sliceutil.IsSorted(data) {
//	    data = sliceutil.Sort(data)
//	}
//
//	// 비싼 키 사전 계산
//	type CachedItem struct {
//	    Item Item
//	    Key  int
//	}
//	cached := sliceutil.Map(items, func(i Item) CachedItem {
//	    return CachedItem{i, expensiveKeyFunc(i)}
//	})
//	sorted := sliceutil.SortBy(cached, func(c CachedItem) int {
//	    return c.Key
//	})
//	result := sliceutil.Map(sorted, func(c CachedItem) Item {
//	    return c.Item
//	})
//
// 표준 라이브러리와 비교:
//   - sort.Slice: 제자리, 원본 수정
//   - sliceutil.Sort: 불변, 새 슬라이스 생성
//   - sort.Slice: 불변성을 위한 수동 복제 필요
//   - sliceutil: 함수형 스타일에 더 편리
//
// 다중 기준 정렬 예제:
//
//	// 다른 순서로 여러 필드 정렬
//	sorted := sliceutil.SortByMulti(employees, func(a, b Employee) bool {
//	    // 1차: 부서 (오름차순)
//	    if a.Department != b.Department {
//	        return a.Department < b.Department
//	    }
//	    // 2차: 급여 (내림차순)
//	    if a.Salary != b.Salary {
//	        return a.Salary > b.Salary
//	    }
//	    // 3차: 이름 (오름차순)
//	    return a.Name < b.Name
//	})

// Sort returns a new slice with elements sorted in ascending order.
// Sort는 오름차순으로 정렬된 요소를 포함하는 새 슬라이스를 반환합니다.
//
// The original slice is not modified. Uses Go's standard sort algorithm.
// 원본 슬라이스는 수정되지 않습니다. Go의 표준 정렬 알고리즘을 사용합니다.
//
// Example
// 예제:
//
//	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
//	sorted := sliceutil.Sort(numbers)
//
// // sorted: [1, 1, 2, 3, 4, 5, 6, 9]
// numbers: [3, 1, 4, 1, 5, 9, 2, 6] (unchanged
// 변경되지 않음)
func Sort[T constraints.Ordered](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	// Clone the slice to avoid modifying the original
	// 원본 수정을 피하기 위해 슬라이스를 복제합니다
	result := make([]T, len(slice))
	copy(result, slice)

	// Sort in ascending order
	// 오름차순으로 정렬
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

// SortDesc returns a new slice with elements sorted in descending order.
// SortDesc는 내림차순으로 정렬된 요소를 포함하는 새 슬라이스를 반환합니다.
//
// The original slice is not modified. Uses Go's standard sort algorithm.
// 원본 슬라이스는 수정되지 않습니다. Go의 표준 정렬 알고리즘을 사용합니다.
//
// Example
// 예제:
//
//	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
//	sorted := sliceutil.SortDesc(numbers)
//
// // sorted: [9, 6, 5, 4, 3, 2, 1, 1]
// numbers: [3, 1, 4, 1, 5, 9, 2, 6] (unchanged
// 변경되지 않음)
func SortDesc[T constraints.Ordered](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	// Clone the slice to avoid modifying the original
	// 원본 수정을 피하기 위해 슬라이스를 복제합니다
	result := make([]T, len(slice))
	copy(result, slice)

	// Sort in descending order
	// 내림차순으로 정렬
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return result
}

// SortBy returns a new slice sorted by the key extracted by keyFunc in ascending order.
// SortBy는 keyFunc으로 추출한 키를 기준으로 오름차순 정렬된 새 슬라이스를 반환합니다.
//
// The original slice is not modified. The keyFunc extracts a comparable key from each element.
// 원본 슬라이스는 수정되지 않습니다. keyFunc은 각 요소에서 비교 가능한 키를 추출합니다.
//
// Example
// 예제:
//
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	people := []Person{
//	    {"Alice", 30},
//	    {"Bob", 25},
//	    {"Charlie", 35},
//	}
//	sortedByAge := sliceutil.SortBy(people, func(p Person) int { return p.Age })
//	// sortedByAge: [{Bob 25}, {Alice 30}, {Charlie 35}]
func SortBy[T any, K constraints.Ordered](slice []T, keyFunc func(T) K) []T {
	if len(slice) == 0 {
		return []T{}
	}

	// Clone the slice to avoid modifying the original
	// 원본 수정을 피하기 위해 슬라이스를 복제합니다
	result := make([]T, len(slice))
	copy(result, slice)

	// Sort by key extracted by keyFunc
	// keyFunc으로 추출한 키를 기준으로 정렬
	sort.Slice(result, func(i, j int) bool {
		return keyFunc(result[i]) < keyFunc(result[j])
	})

	return result
}

// IsSorted checks if the slice is sorted in ascending order.
// IsSorted는 슬라이스가 오름차순으로 정렬되어 있는지 확인합니다.
//
// Returns true if the slice is sorted in ascending order, false otherwise.
// 슬라이스가 오름차순으로 정렬되어 있으면 true를, 그렇지 않으면 false를 반환합니다.
//
// An empty slice or a slice with one element is considered sorted.
// 비어있는 슬라이스나 요소가 하나인 슬라이스는 정렬된 것으로 간주됩니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	isSorted := sliceutil.IsSorted(numbers)  // true
//
//	numbers2 := []int{5, 4, 3, 2, 1}
//	isSorted2 := sliceutil.IsSorted(numbers2)  // false
func IsSorted[T constraints.Ordered](slice []T) bool {
	if len(slice) <= 1 {
		return true
	}

	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}

	return true
}

// IsSortedDesc checks if the slice is sorted in descending order.
// IsSortedDesc는 슬라이스가 내림차순으로 정렬되어 있는지 확인합니다.
//
// Returns true if the slice is sorted in descending order, false otherwise.
// 슬라이스가 내림차순으로 정렬되어 있으면 true를, 그렇지 않으면 false를 반환합니다.
//
// An empty slice or a slice with one element is considered sorted.
// 비어있는 슬라이스나 요소가 하나인 슬라이스는 정렬된 것으로 간주됩니다.
//
// Example
// 예제:
//
//	numbers := []int{5, 4, 3, 2, 1}
//	isSortedDesc := sliceutil.IsSortedDesc(numbers)  // true
//
//	numbers2 := []int{1, 2, 3, 4, 5}
//	isSortedDesc2 := sliceutil.IsSortedDesc(numbers2)  // false
func IsSortedDesc[T constraints.Ordered](slice []T) bool {
	if len(slice) <= 1 {
		return true
	}

	for i := 1; i < len(slice); i++ {
		if slice[i] > slice[i-1] {
			return false
		}
	}

	return true
}

// SortByMulti returns a new slice sorted by using a custom comparison function.
// SortByMulti는 사용자 정의 비교 함수를 사용하여 정렬된 새 슬라이스를 반환합니다.
//
// The original slice is not modified. The less function should return true if element i
// should sort before element j. This allows for multi-key sorting by comparing multiple fields.
// 원본 슬라이스는 수정되지 않습니다. less 함수는 요소 i가 요소 j보다 앞에 정렬되어야 하면 true를 반환해야 합니다.
// 이를 통해 여러 필드를 비교하여 다중 키 정렬이 가능합니다.
//
// Example
// 예제:
//
//	type Person struct {
//	    Name string
//	    Age  int
//	    Score float64
//	}
//	people := []Person{
//	    {"Alice", 30, 95.5},
//	    {"Bob", 25, 88.0},
//	    {"Alice", 25, 92.0},
//	    {"Bob", 25, 90.0},
//	}
//
// // Sort by Name (ascending), then Age (ascending), then Score (descending)
// Name(오름차순), Age(오름차순), Score(내림차순) 순으로 정렬
//
//	sorted := sliceutil.SortByMulti(people, func(i, j Person) bool {
//	    if i.Name != j.Name {
//	        return i.Name < j.Name
//	    }
//	    if i.Age != j.Age {
//	        return i.Age < j.Age
//	    }
//	    return i.Score > j.Score // Descending
//	})
//	// sorted: [{Alice 25 92.0}, {Alice 30 95.5}, {Bob 25 90.0}, {Bob 25 88.0}]
func SortByMulti[T any](slice []T, less func(i, j T) bool) []T {
	if len(slice) == 0 {
		return []T{}
	}

	// Clone the slice to avoid modifying the original
	// 원본 수정을 피하기 위해 슬라이스를 복제합니다
	result := make([]T, len(slice))
	copy(result, slice)

	// Sort using the provided less function
	// 제공된 less 함수를 사용하여 정렬
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})

	return result
}
