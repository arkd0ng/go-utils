package sliceutil

// DiffResult holds the result of a diff operation between two slices.
// DiffResult는 두 슬라이스 간의 차이 연산 결과를 저장합니다.
type DiffResult[T any] struct {
	// Elements in new but not in old
	// 새 슬라이스에만 있는 요소
	Added []T
	// Elements in old but not in new
	// 이전 슬라이스에만 있는 요소
	Removed []T
	// Elements in both old and new
	// 양쪽 모두에 있는 요소
	Unchanged []T
}

// Diff compares two slices and returns the differences.
// Returns elements that were added (in new but not old), removed (in old but not new),
// and unchanged (in both).
//
// Diff는 두 슬라이스를 비교하여 차이를 반환합니다.
// 추가된 요소(새 슬라이스에만 있음), 제거된 요소(이전 슬라이스에만 있음),
// 변경되지 않은 요소(양쪽 모두에 있음)를 반환합니다.
//
// Example:
//
//	old := []int{1, 2, 3, 4}
//	new := []int{2, 3, 4, 5}
//	diff := sliceutil.Diff(old, new)
//	// diff.Added: [5]
//	// diff.Removed: [1]
//	// diff.Unchanged: [2, 3, 4]
func Diff[T comparable](old, new []T) DiffResult[T] {
	oldSet := make(map[T]bool)
	for _, v := range old {
		oldSet[v] = true
	}

	newSet := make(map[T]bool)
	for _, v := range new {
		newSet[v] = true
	}

	var added, removed, unchanged []T

	// Find added elements (in new but not in old)
	for _, v := range new {
		if !oldSet[v] {
			added = append(added, v)
		} else {
			unchanged = append(unchanged, v)
		}
	}

	// Find removed elements (in old but not in new)
	for _, v := range old {
		if !newSet[v] {
			removed = append(removed, v)
		}
	}

	// Remove duplicates from unchanged
	unchanged = Unique(unchanged)

	return DiffResult[T]{
		Added:     added,
		Removed:   removed,
		Unchanged: unchanged,
	}
}

// DiffBy compares two slices using a key function and returns the differences.
// The key function extracts a comparable key from each element for comparison.
// This is useful for comparing slices of structs or complex types.
//
// DiffBy는 키 함수를 사용하여 두 슬라이스를 비교하고 차이를 반환합니다.
// 키 함수는 각 요소에서 비교 가능한 키를 추출합니다.
// 이는 구조체나 복잡한 타입의 슬라이스를 비교할 때 유용합니다.
//
// Example:
//
//	type User struct { ID int; Name string }
//	old := []User{{1, "Alice"}, {2, "Bob"}}
//	new := []User{{2, "Bob"}, {3, "Charlie"}}
//	diff := sliceutil.DiffBy(old, new, func(u User) int { return u.ID })
//	// diff.Added: [{3, "Charlie"}]
//	// diff.Removed: [{1, "Alice"}]
//	// diff.Unchanged: [{2, "Bob"}]
func DiffBy[T any, K comparable](old, new []T, keyFunc func(T) K) DiffResult[T] {
	oldMap := make(map[K]T)
	for _, v := range old {
		oldMap[keyFunc(v)] = v
	}

	newMap := make(map[K]T)
	for _, v := range new {
		newMap[keyFunc(v)] = v
	}

	var added, removed, unchanged []T

	// Find added elements (keys in new but not in old)
	for key, v := range newMap {
		if _, exists := oldMap[key]; !exists {
			added = append(added, v)
		} else {
			unchanged = append(unchanged, v)
		}
	}

	// Find removed elements (keys in old but not in new)
	for key, v := range oldMap {
		if _, exists := newMap[key]; !exists {
			removed = append(removed, v)
		}
	}

	return DiffResult[T]{
		Added:     added,
		Removed:   removed,
		Unchanged: unchanged,
	}
}

// EqualUnordered checks if two slices contain the same elements, regardless of order.
// Returns true if both slices have the same elements with the same frequencies.
//
// EqualUnordered는 순서에 관계없이 두 슬라이스가 같은 요소를 포함하는지 확인합니다.
// 양쪽 슬라이스가 같은 빈도의 같은 요소를 가지면 true를 반환합니다.
//
// Example:
//
//	a := []int{1, 2, 3, 2}
//	b := []int{3, 2, 1, 2}
//	equal := sliceutil.EqualUnordered(a, b) // true
//
//	c := []int{1, 2, 3}
//	d := []int{1, 2, 3, 3}
//	equal := sliceutil.EqualUnordered(c, d) // false (different frequencies)
func EqualUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	// Get frequencies of both slices
	freqA := Frequencies(a)
	freqB := Frequencies(b)

	// Check if frequencies match
	if len(freqA) != len(freqB) {
		return false
	}

	for key, countA := range freqA {
		countB, exists := freqB[key]
		if !exists || countA != countB {
			return false
		}
	}

	return true
}

// HasDuplicates checks if a slice contains any duplicate elements.
// Returns true if there are duplicates, false otherwise.
//
// HasDuplicates는 슬라이스에 중복 요소가 있는지 확인합니다.
// 중복이 있으면 true를, 없으면 false를 반환합니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 2, 4}
//	hasDups := sliceutil.HasDuplicates(numbers) // true
//
//	unique := []int{1, 2, 3, 4, 5}
//	hasDups := sliceutil.HasDuplicates(unique) // false
func HasDuplicates[T comparable](slice []T) bool {
	seen := make(map[T]bool)
	for _, v := range slice {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
