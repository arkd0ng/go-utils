package sliceutil

import (
	"math/rand"
	"time"
)

// slice.go contains slicing operations for slices.
// slice.go는 슬라이스 슬라이싱 작업을 포함합니다.

// Chunk splits a slice into chunks of the specified size.
// The last chunk may be smaller than the specified size.
// Chunk는 슬라이스를 지정된 크기의 청크로 분할합니다.
// 마지막 청크는 지정된 크기보다 작을 수 있습니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5, 6, 7}
//	chunks := sliceutil.Chunk(numbers, 3) // [[1, 2, 3], [4, 5, 6], [7]]
//
//	words := []string{"a", "b", "c", "d", "e"}
//	chunks := sliceutil.Chunk(words, 2) // [["a", "b"], ["c", "d"], ["e"]]
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	if len(slice) == 0 {
		return [][]T{}
	}

	numChunks := (len(slice) + size - 1) / size
	result := make([][]T, 0, numChunks)

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}

	return result
}

// Take returns the first n elements from the slice.
// If n is greater than the slice length, returns the entire slice.
// Take는 슬라이스에서 첫 n개 요소를 반환합니다.
// n이 슬라이스 길이보다 크면 전체 슬라이스를 반환합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	first3 := sliceutil.Take(numbers, 3) // [1, 2, 3]
//
//	words := []string{"hello", "world", "foo"}
//	first2 := sliceutil.Take(words, 2) // ["hello", "world"]
func Take[T any](slice []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(slice) {
		return append([]T{}, slice...)
	}
	return append([]T{}, slice[:n]...)
}

// TakeLast returns the last n elements from the slice.
// If n is greater than the slice length, returns the entire slice.
// TakeLast는 슬라이스에서 마지막 n개 요소를 반환합니다.
// n이 슬라이스 길이보다 크면 전체 슬라이스를 반환합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	last3 := sliceutil.TakeLast(numbers, 3) // [3, 4, 5]
//
//	words := []string{"hello", "world", "foo"}
//	last2 := sliceutil.TakeLast(words, 2) // ["world", "foo"]
func TakeLast[T any](slice []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(slice) {
		return append([]T{}, slice...)
	}
	return append([]T{}, slice[len(slice)-n:]...)
}

// Drop returns a slice without the first n elements.
// If n is greater than the slice length, returns an empty slice.
// Drop은 첫 n개 요소를 제외한 슬라이스를 반환합니다.
// n이 슬라이스 길이보다 크면 빈 슬라이스를 반환합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	rest := sliceutil.Drop(numbers, 2) // [3, 4, 5]
//
//	words := []string{"hello", "world", "foo"}
//	rest := sliceutil.Drop(words, 1) // ["world", "foo"]
func Drop[T any](slice []T, n int) []T {
	if n <= 0 {
		return append([]T{}, slice...)
	}
	if n >= len(slice) {
		return []T{}
	}
	return append([]T{}, slice[n:]...)
}

// DropLast returns a slice without the last n elements.
// If n is greater than the slice length, returns an empty slice.
// DropLast는 마지막 n개 요소를 제외한 슬라이스를 반환합니다.
// n이 슬라이스 길이보다 크면 빈 슬라이스를 반환합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	rest := sliceutil.DropLast(numbers, 2) // [1, 2, 3]
//
//	words := []string{"hello", "world", "foo"}
//	rest := sliceutil.DropLast(words, 1) // ["hello", "world"]
func DropLast[T any](slice []T, n int) []T {
	if n <= 0 {
		return append([]T{}, slice...)
	}
	if n >= len(slice) {
		return []T{}
	}
	return append([]T{}, slice[:len(slice)-n]...)
}

// Slice returns a slice of elements from start to end index.
// Negative indices count from the end of the slice.
// Slice는 start 인덱스에서 end 인덱스까지의 요소 슬라이스를 반환합니다.
// 음수 인덱스는 슬라이스 끝에서부터 계산합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	sub := sliceutil.Slice(numbers, 1, 4) // [2, 3, 4]
//
//	sub := sliceutil.Slice(numbers, -3, -1) // [3, 4]
func Slice[T any](slice []T, start, end int) []T {
	length := len(slice)

	// Handle negative indices
	if start < 0 {
		start = length + start
		if start < 0 {
			start = 0
		}
	}
	if end < 0 {
		end = length + end
		if end < 0 {
			end = 0
		}
	}

	// Clamp to valid range
	if start > length {
		start = length
	}
	if end > length {
		end = length
	}

	// Ensure start <= end
	if start > end {
		return []T{}
	}

	return append([]T{}, slice[start:end]...)
}

// Sample returns n random elements from the slice.
// Elements are selected without replacement (no duplicates).
// If n is greater than the slice length, returns all elements in random order.
// Sample은 슬라이스에서 n개의 랜덤 요소를 반환합니다.
// 요소는 중복 없이 선택됩니다 (복원 추출 없음).
// n이 슬라이스 길이보다 크면 모든 요소를 랜덤 순서로 반환합니다.
//
// Example / 예제:
//
//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	sample := sliceutil.Sample(numbers, 3) // e.g., [7, 2, 9]
//
//	words := []string{"apple", "banana", "cherry", "date"}
//	sample := sliceutil.Sample(words, 2) // e.g., ["cherry", "apple"]
func Sample[T any](slice []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(slice) {
		// Return all elements in random order
		result := append([]T{}, slice...)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(result), func(i, j int) {
			result[i], result[j] = result[j], result[i]
		})
		return result
	}

	// Use Fisher-Yates shuffle variant for sampling
	result := append([]T{}, slice...)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		j := i + r.Intn(len(result)-i)
		result[i], result[j] = result[j], result[i]
	}

	return result[:n]
}
