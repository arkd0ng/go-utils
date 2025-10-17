package sliceutil

// slice.go contains slicing operations for slices.
// slice.go는 슬라이스 슬라이싱 작업을 포함합니다.

// Chunk splits a slice into chunks of the specified size.
// The last chunk may be smaller than the specified size.
// Chunk는 슬라이스를 지정된 크기의 청크로 분할합니다.
// 마지막 청크는 지정된 크기보다 작을 수 있습니다.
//
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
// Example
// 예제:
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
		rngLock.Lock()
		defer rngLock.Unlock()
		rng.Shuffle(len(result), func(i, j int) {
			result[i], result[j] = result[j], result[i]
		})
		return result
	}

	// Use Fisher-Yates shuffle variant for sampling
	result := append([]T{}, slice...)
	rngLock.Lock()
	defer rngLock.Unlock()

	for i := 0; i < n; i++ {
		j := i + rng.Intn(len(result)-i)
		result[i], result[j] = result[j], result[i]
	}

	return result[:n]
}

// TakeWhile returns elements from the beginning while predicate is true.
// Stops at the first element that doesn't satisfy the predicate.
// TakeWhile은 조건이 참인 동안 처음부터 요소를 반환합니다.
// 조건을 만족하지 않는 첫 번째 요소에서 중지합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 1, 2}
//	result := sliceutil.TakeWhile(numbers, func(n int) bool {
//	    return n < 4
//	}) // [1, 2, 3]
//
//	words := []string{"apple", "apricot", "banana", "avocado"}
//	result := sliceutil.TakeWhile(words, func(s string) bool {
//	    return len(s) > 0 && s[0] == 'a'
//	}) // ["apple", "apricot"]
func TakeWhile[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !predicate(v) {
			break
		}
		result = append(result, v)
	}
	return result
}

// DropWhile returns elements after dropping the beginning while predicate is true.
// Starts including elements from the first one that doesn't satisfy the predicate.
// DropWhile은 조건이 참인 동안 처음 요소를 제거한 후 요소를 반환합니다.
// 조건을 만족하지 않는 첫 번째 요소부터 포함하기 시작합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 1, 2}
//	result := sliceutil.DropWhile(numbers, func(n int) bool {
//	    return n < 4
//	}) // [4, 1, 2]
//
//	words := []string{"apple", "apricot", "banana", "avocado"}
//	result := sliceutil.DropWhile(words, func(s string) bool {
//	    return len(s) > 0 && s[0] == 'a'
//	}) // ["banana", "avocado"]
func DropWhile[T any](slice []T, predicate func(T) bool) []T {
	for i, v := range slice {
		if !predicate(v) {
			return append([]T{}, slice[i:]...)
		}
	}
	return []T{}
}

// Interleave merges multiple slices by taking one element from each in turn.
// Continues until all slices are exhausted.
// Interleave은 각 슬라이스에서 차례로 하나씩 요소를 가져와 여러 슬라이스를 병합합니다.
// 모든 슬라이스가 소진될 때까지 계속합니다.
//
// Example
// 예제:
//
//	a := []int{1, 2, 3}
//	b := []int{10, 20, 30}
//	c := []int{100, 200}
//	result := sliceutil.Interleave(a, b, c)
//	// result: [1, 10, 100, 2, 20, 200, 3, 30]
//
//	words1 := []string{"a", "b"}
//	words2 := []string{"x", "y", "z"}
//	result := sliceutil.Interleave(words1, words2)
//	// result: ["a", "x", "b", "y", "z"]
func Interleave[T any](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}

	// Calculate total length
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, 0, totalLen)
	maxLen := 0
	for _, s := range slices {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	for i := 0; i < maxLen; i++ {
		for _, s := range slices {
			if i < len(s) {
				result = append(result, s[i])
			}
		}
	}

	return result
}
