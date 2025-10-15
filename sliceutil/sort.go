package sliceutil

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Sort returns a new slice with elements sorted in ascending order.
// Sort는 오름차순으로 정렬된 요소를 포함하는 새 슬라이스를 반환합니다.
//
// The original slice is not modified. Uses Go's standard sort algorithm.
// 원본 슬라이스는 수정되지 않습니다. Go의 표준 정렬 알고리즘을 사용합니다.
//
// Example / 예제:
//
//	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
//	sorted := sliceutil.Sort(numbers)
//	// sorted: [1, 1, 2, 3, 4, 5, 6, 9]
//	// numbers: [3, 1, 4, 1, 5, 9, 2, 6] (unchanged / 변경되지 않음)
func Sort[T constraints.Ordered](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	// Clone the slice to avoid modifying the original
	// 원본 수정을 피하기 위해 슬라이스를 복제합니다
	result := make([]T, len(slice))
	copy(result, slice)

	// Sort in ascending order / 오름차순으로 정렬
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
// Example / 예제:
//
//	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
//	sorted := sliceutil.SortDesc(numbers)
//	// sorted: [9, 6, 5, 4, 3, 2, 1, 1]
//	// numbers: [3, 1, 4, 1, 5, 9, 2, 6] (unchanged / 변경되지 않음)
func SortDesc[T constraints.Ordered](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	// Clone the slice to avoid modifying the original
	// 원본 수정을 피하기 위해 슬라이스를 복제합니다
	result := make([]T, len(slice))
	copy(result, slice)

	// Sort in descending order / 내림차순으로 정렬
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
// Example / 예제:
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

	// Sort by key extracted by keyFunc / keyFunc으로 추출한 키를 기준으로 정렬
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
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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
//	// Sort by Name (ascending), then Age (ascending), then Score (descending)
//	// Name(오름차순), Age(오름차순), Score(내림차순) 순으로 정렬
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
