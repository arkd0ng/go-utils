package sliceutil

// combinatorial.go contains combinatorial operations for slices.
// combinatorial.go는 슬라이스 조합 작업을 포함합니다.

// Permutations returns all possible permutations of the slice.
// Returns a slice of slices, where each sub-slice is a permutation.
// Warning: The number of permutations grows factorially (n!).
// Permutations는 슬라이스의 모든 가능한 순열을 반환합니다.
// 각 하위 슬라이스가 순열인 슬라이스의 슬라이스를 반환합니다.
// 경고: 순열의 수는 팩토리얼로 증가합니다 (n!).
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3}
//	perms := sliceutil.Permutations(numbers)
//	// perms: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
//
//	letters := []string{"a", "b"}
//	perms := sliceutil.Permutations(letters)
//	// perms: [["a","b"], ["b","a"]]
//
// Performance note / 성능 참고:
//   n=5: 120 permutations
//   n=10: 3,628,800 permutations
//   Use with caution for large slices!
func Permutations[T any](slice []T) [][]T {
	if len(slice) == 0 {
		return [][]T{{}}
	}

	// Create a copy to avoid modifying the original slice
	// 원본 슬라이스 수정을 방지하기 위해 복사본 생성
	sliceCopy := make([]T, len(slice))
	copy(sliceCopy, slice)

	result := [][]T{}
	permute(sliceCopy, 0, &result)
	return result
}

// permute is a helper function for generating permutations using Heap's algorithm.
// permute는 Heap의 알고리즘을 사용하여 순열을 생성하는 헬퍼 함수입니다.
func permute[T any](slice []T, k int, result *[][]T) {
	if k == len(slice)-1 {
		// Make a copy of the current permutation
		perm := make([]T, len(slice))
		copy(perm, slice)
		*result = append(*result, perm)
		return
	}

	for i := k; i < len(slice); i++ {
		slice[k], slice[i] = slice[i], slice[k]
		permute(slice, k+1, result)
		slice[k], slice[i] = slice[i], slice[k] // backtrack
	}
}

// Combinations returns all possible combinations of k elements from the slice.
// Returns a slice of slices, where each sub-slice is a combination.
// Warning: The number of combinations is C(n, k) = n! / (k! * (n-k)!).
// Combinations는 슬라이스에서 k개 요소의 모든 가능한 조합을 반환합니다.
// 각 하위 슬라이스가 조합인 슬라이스의 슬라이스를 반환합니다.
// 경고: 조합의 수는 C(n, k) = n! / (k! * (n-k)!)입니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4}
//	combs := sliceutil.Combinations(numbers, 2)
//	// combs: [[1,2], [1,3], [1,4], [2,3], [2,4], [3,4]]
//
//	letters := []string{"a", "b", "c"}
//	combs := sliceutil.Combinations(letters, 2)
//	// combs: [["a","b"], ["a","c"], ["b","c"]]
//
// Performance note / 성능 참고:
//   C(10, 5) = 252 combinations
//   C(20, 10) = 184,756 combinations
//   Use with caution for large values!
func Combinations[T any](slice []T, k int) [][]T {
	if k < 0 || k > len(slice) {
		return [][]T{}
	}
	if k == 0 {
		return [][]T{{}}
	}
	if k == len(slice) {
		return [][]T{append([]T{}, slice...)}
	}

	result := [][]T{}
	combine(slice, k, 0, []T{}, &result)
	return result
}

// combine is a helper function for generating combinations recursively.
// combine는 재귀적으로 조합을 생성하는 헬퍼 함수입니다.
func combine[T any](slice []T, k, start int, current []T, result *[][]T) {
	if len(current) == k {
		// Make a copy of the current combination
		comb := make([]T, len(current))
		copy(comb, current)
		*result = append(*result, comb)
		return
	}

	for i := start; i < len(slice); i++ {
		combine(slice, k, i+1, append(current, slice[i]), result)
	}
}
