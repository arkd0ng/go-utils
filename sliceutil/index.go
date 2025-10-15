package sliceutil

// index.go contains index-based slice operations.
// index.go는 인덱스 기반 슬라이스 작업을 포함합니다.

// FindIndices returns all indices where the predicate returns true.
// Returns an empty slice if no elements match.
//
// FindIndices는 조건이 true를 반환하는 모든 인덱스를 반환합니다.
// 일치하는 요소가 없으면 빈 슬라이스를 반환합니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evenIndices := sliceutil.FindIndices(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // [1, 3, 5] (indices of 2, 4, 6)
func FindIndices[T any](slice []T, predicate func(T) bool) []int {
	var indices []int
	for i, v := range slice {
		if predicate(v) {
			indices = append(indices, i)
		}
	}
	return indices
}

// AtIndices returns elements at the specified indices.
// Indices that are out of bounds are silently skipped.
// Negative indices are not supported and will be skipped.
//
// AtIndices는 지정된 인덱스의 요소를 반환합니다.
// 범위를 벗어난 인덱스는 자동으로 건너뜁니다.
// 음수 인덱스는 지원되지 않으며 건너뜁니다.
//
// Example:
//
//	numbers := []int{10, 20, 30, 40, 50}
//	selected := sliceutil.AtIndices(numbers, []int{0, 2, 4})
//	// [10, 30, 50]
//
//	// Out of bounds indices are skipped
//	selected := sliceutil.AtIndices(numbers, []int{0, 10, 2})
//	// [10, 30]
func AtIndices[T any](slice []T, indices []int) []T {
	result := make([]T, 0, len(indices))
	for _, idx := range indices {
		if idx >= 0 && idx < len(slice) {
			result = append(result, slice[idx])
		}
	}
	return result
}

// RemoveIndices returns a new slice with elements at the specified indices removed.
// Indices that are out of bounds are silently skipped.
// Negative indices are not supported and will be skipped.
// The original slice is not modified.
//
// RemoveIndices는 지정된 인덱스의 요소가 제거된 새 슬라이스를 반환합니다.
// 범위를 벗어난 인덱스는 자동으로 건너뜁니다.
// 음수 인덱스는 지원되지 않으며 건너뜁니다.
// 원본 슬라이스는 수정되지 않습니다.
//
// Example:
//
//	numbers := []int{10, 20, 30, 40, 50}
//	result := sliceutil.RemoveIndices(numbers, []int{1, 3})
//	// [10, 30, 50] (removed 20 and 40)
//
//	// Out of bounds indices are skipped
//	result := sliceutil.RemoveIndices(numbers, []int{1, 10, 3})
//	// [10, 30, 50] (index 10 is skipped)
func RemoveIndices[T any](slice []T, indices []int) []T {
	if len(indices) == 0 {
		return Clone(slice)
	}

	// Create a set of indices to remove for O(1) lookup
	toRemove := make(map[int]bool)
	for _, idx := range indices {
		if idx >= 0 && idx < len(slice) {
			toRemove[idx] = true
		}
	}

	// Build result slice
	result := make([]T, 0, len(slice)-len(toRemove))
	for i, v := range slice {
		if !toRemove[i] {
			result = append(result, v)
		}
	}

	return result
}
