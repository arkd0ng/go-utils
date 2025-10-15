package sliceutil

// advanced.go contains advanced functional programming operations.
// advanced.go는 고급 함수형 프로그래밍 작업을 포함합니다.

// Scan returns a slice containing the successive reduced values.
// It's like Reduce but returns all intermediate results.
// The first element of the result is always the initial value.
//
// Scan은 연속적인 reduce 값을 포함하는 슬라이스를 반환합니다.
// Reduce와 비슷하지만 모든 중간 결과를 반환합니다.
// 결과의 첫 번째 요소는 항상 초기값입니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := sliceutil.Scan(numbers, 0, func(acc, n int) int {
//	    return acc + n
//	})
//	// [0, 1, 3, 6, 10, 15] (cumulative sum)
//
//	// Running product
//	result := sliceutil.Scan(numbers, 1, func(acc, n int) int {
//	    return acc * n
//	})
//	// [1, 1, 2, 6, 24, 120] (factorial-like)
func Scan[T any](slice []T, initial T, accumulator func(T, T) T) []T {
	result := make([]T, len(slice)+1)
	result[0] = initial

	acc := initial
	for i, v := range slice {
		acc = accumulator(acc, v)
		result[i+1] = acc
	}

	return result
}

// ZipWith combines two slices into one using a zipper function.
// The resulting slice has the length of the shorter input slice.
//
// ZipWith는 zipper 함수를 사용하여 두 슬라이스를 하나로 결합합니다.
// 결과 슬라이스는 더 짧은 입력 슬라이스의 길이를 가집니다.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	strings := []string{"a", "b", "c"}
//	result := sliceutil.ZipWith(numbers, strings, func(n int, s string) string {
//	    return fmt.Sprintf("%d:%s", n, s)
//	})
//	// ["1:a", "2:b", "3:c"]
//
//	// Sum corresponding elements
//	a := []int{1, 2, 3, 4}
//	b := []int{10, 20, 30}
//	result := sliceutil.ZipWith(a, b, func(x, y int) int {
//	    return x + y
//	})
//	// [11, 22, 33] (length 3, stops at shorter slice)
func ZipWith[T, U, R any](a []T, b []U, zipper func(T, U) R) []R {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}

	result := make([]R, length)
	for i := 0; i < length; i++ {
		result[i] = zipper(a[i], b[i])
	}

	return result
}

// RotateLeft rotates the slice to the left by n positions.
// Elements that fall off the left are appended to the right.
// Returns a new slice, the original is not modified.
//
// RotateLeft는 슬라이스를 왼쪽으로 n 위치만큼 회전합니다.
// 왼쪽으로 떨어진 요소는 오른쪽에 추가됩니다.
// 새 슬라이스를 반환하며 원본은 수정되지 않습니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := sliceutil.RotateLeft(numbers, 2)
//	// [3, 4, 5, 1, 2]
//
//	// Negative n rotates right
//	result := sliceutil.RotateLeft(numbers, -1)
//	// [5, 1, 2, 3, 4]
func RotateLeft[T any](slice []T, n int) []T {
	if len(slice) == 0 {
		return []T{}
	}

	// Normalize n to be within [0, len)
	length := len(slice)
	n = ((n % length) + length) % length

	result := make([]T, length)
	for i := 0; i < length; i++ {
		result[i] = slice[(i+n)%length]
	}

	return result
}

// RotateRight rotates the slice to the right by n positions.
// Elements that fall off the right are prepended to the left.
// Returns a new slice, the original is not modified.
//
// RotateRight는 슬라이스를 오른쪽으로 n 위치만큼 회전합니다.
// 오른쪽으로 떨어진 요소는 왼쪽에 추가됩니다.
// 새 슬라이스를 반환하며 원본은 수정되지 않습니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := sliceutil.RotateRight(numbers, 2)
//	// [4, 5, 1, 2, 3]
//
//	// Negative n rotates left
//	result := sliceutil.RotateRight(numbers, -1)
//	// [2, 3, 4, 5, 1]
func RotateRight[T any](slice []T, n int) []T {
	// Rotating right by n is the same as rotating left by -n
	return RotateLeft(slice, -n)
}
