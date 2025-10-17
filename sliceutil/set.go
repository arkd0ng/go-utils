package sliceutil

// set.go contains set operations for slices.
// set.go는 슬라이스 집합 작업을 포함합니다.

// Union returns the union of two slices (all unique elements from both).
// Union은 두 슬라이스의 합집합을 반환합니다 (양쪽의 모든 고유 요소).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3}
//	b := []int{3, 4, 5}
//	union := sliceutil.Union(a, b) // [1, 2, 3, 4, 5]
//
//	words1 := []string{"apple", "banana"}
//	words2 := []string{"banana", "cherry"}
//	union := sliceutil.Union(words1, words2) // ["apple", "banana", "cherry"]
func Union[T comparable](a, b []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(a)+len(b))

	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	for _, v := range b {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Intersection returns the intersection of two slices (common elements).
// Intersection은 두 슬라이스의 교집합을 반환합니다 (공통 요소).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{3, 4, 5, 6}
//	common := sliceutil.Intersection(a, b) // [3, 4]
//
//	words1 := []string{"apple", "banana", "cherry"}
//	words2 := []string{"banana", "cherry", "date"}
//	common := sliceutil.Intersection(words1, words2) // ["banana", "cherry"]
func Intersection[T comparable](a, b []T) []T {
	setB := make(map[T]bool)
	for _, v := range b {
		setB[v] = true
	}

	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, v := range a {
		if setB[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Difference returns the difference of two slices (elements in a but not in b).
// Difference는 두 슬라이스의 차집합을 반환합니다 (a에 있지만 b에 없는 요소).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{3, 4, 5, 6}
//	diff := sliceutil.Difference(a, b) // [1, 2]
//
//	words1 := []string{"apple", "banana", "cherry"}
//	words2 := []string{"banana", "date"}
//	diff := sliceutil.Difference(words1, words2) // ["apple", "cherry"]
func Difference[T comparable](a, b []T) []T {
	setB := make(map[T]bool)
	for _, v := range b {
		setB[v] = true
	}

	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, v := range a {
		if !setB[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// SymmetricDifference returns the symmetric difference of two slices.
// Elements that are in either a or b, but not in both.
// SymmetricDifference는 두 슬라이스의 대칭 차집합을 반환합니다.
// a 또는 b에 있지만 둘 다에 있지 않은 요소.
//
// Example
// 예제:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{3, 4, 5, 6}
//	symDiff := sliceutil.SymmetricDifference(a, b) // [1, 2, 5, 6]
//
//	words1 := []string{"apple", "banana"}
//	words2 := []string{"banana", "cherry"}
//	symDiff := sliceutil.SymmetricDifference(words1, words2) // ["apple", "cherry"]
func SymmetricDifference[T comparable](a, b []T) []T {
	setA := make(map[T]bool)
	setB := make(map[T]bool)

	for _, v := range a {
		setA[v] = true
	}
	for _, v := range b {
		setB[v] = true
	}

	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, v := range a {
		if !setB[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	for _, v := range b {
		if !setA[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// IsSubset returns true if a is a subset of b (all elements of a are in b).
// IsSubset은 a가 b의 부분집합이면 true를 반환합니다 (a의 모든 요소가 b에 있음).
//
// Example
// 예제:
//
//	a := []int{1, 2}
//	b := []int{1, 2, 3, 4}
//	isSubset := sliceutil.IsSubset(a, b) // true
//
//	a := []int{1, 5}
//	b := []int{1, 2, 3, 4}
//	isSubset := sliceutil.IsSubset(a, b) // false
func IsSubset[T comparable](a, b []T) bool {
	setB := make(map[T]bool)
	for _, v := range b {
		setB[v] = true
	}

	for _, v := range a {
		if !setB[v] {
			return false
		}
	}

	return true
}

// IsSuperset returns true if a is a superset of b (all elements of b are in a).
// IsSuperset은 a가 b의 상위집합이면 true를 반환합니다 (b의 모든 요소가 a에 있음).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{1, 2}
//	isSuperset := sliceutil.IsSuperset(a, b) // true
//
//	a := []int{1, 2, 3}
//	b := []int{1, 5}
//	isSuperset := sliceutil.IsSuperset(a, b) // false
func IsSuperset[T comparable](a, b []T) bool {
	return IsSubset(b, a)
}
