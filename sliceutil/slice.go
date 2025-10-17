package sliceutil

// slice.go provides slicing, sampling, and segmentation operations for slices.
//
// This file implements operations for extracting, splitting, and manipulating
// contiguous or sampled portions of slices:
//
// Slicing Operations:
//
// Chunk (Split into Equal Parts):
//   - Chunk(slice, size): Split into fixed-size chunks
//     Time: O(n), Space: O(n)
//     Last chunk may be smaller if slice not evenly divisible
//     Returns [][]T (slice of slices)
//     Useful for: Batch processing, pagination, data partitioning
//     Example: [1,2,3,4,5,6,7] chunks of 3 → [[1,2,3], [4,5,6], [7]]
//
// Take/Drop (Prefix/Suffix Operations):
//   - Take(slice, n): Get first n elements
//     Time: O(n), Space: O(n)
//     Returns copy, safe to modify
//     If n > len, returns entire slice
//   - TakeLast(slice, n): Get last n elements
//     Time: O(n), Space: O(n)
//     Extracts suffix
//   - Drop(slice, n): Remove first n elements
//     Time: O(n), Space: O(n)
//     Returns remaining elements
//   - DropLast(slice, n): Remove last n elements
//     Time: O(n), Space: O(n)
//     Returns prefix
//
// Conditional Take/Drop:
//   - TakeWhile(slice, predicate): Take elements while condition holds
//     Time: O(k) where k = elements taken, Space: O(k)
//     Stops at first false predicate
//     Example: TakeWhile([1,2,3,4,1,2], n < 4) → [1,2,3]
//   - DropWhile(slice, predicate): Skip elements while condition holds
//     Time: O(k) where k = elements dropped, Space: O(n-k)
//     Returns remainder after first false
//     Example: DropWhile([1,2,3,4,5], n < 4) → [4,5]
//
// Generic Slicing:
//   - Slice(slice, start, end): Extract arbitrary range
//     Time: O(end-start), Space: O(end-start)
//     Python-like slicing [start:end]
//     Handles negative indices (from end)
//     Clamps out-of-bounds indices
//     Returns copy of range
//
// Sampling:
//   - Sample(slice, n): Random sample of n elements
//     Time: O(n), Space: O(n)
//     Without replacement (unique elements)
//     If n >= len, returns shuffled copy
//     Uses global thread-safe RNG
//     Useful for: Random selection, testing, data sampling
//
// Interleaving:
//   - Interleave(slices...): Merge multiple slices round-robin
//     Time: O(total elements), Space: O(total elements)
//     Takes one element from each slice in order
//     Continues until all slices exhausted
//     Example: [1,2,3] + [a,b] + [x,y,z] → [1,a,x,2,b,y,3,z]
//     Useful for: Merging parallel streams, fair scheduling
//
// Design Principles:
//   - All functions return new slices (immutable operations)
//   - Safe to modify returned slices without affecting input
//   - Consistent handling of edge cases (empty, nil, oversized n)
//   - Take/Drop operations: n <= 0 returns empty, n >= len returns all/empty
//   - Conditional operations: Predicate evaluated left-to-right
//
// Edge Case Handling:
//   - Empty slices: Return empty results (no panic)
//   - Nil slices: Treated as empty
//   - Negative n: Treated as 0 for Take/Drop
//   - Oversized n: Clamped to slice length
//   - Slice with negative indices: -1 = last element, -2 = second-to-last, etc.
//
// Performance Characteristics:
//   - Chunk: Single pass, allocates chunks
//   - Take/Drop: Copy operations, linear in result size
//   - TakeWhile/DropWhile: Early termination possible (best case O(1))
//   - Sample: Shuffle algorithm (Fisher-Yates), always O(n)
//   - Interleave: Single pass over all input slices
//
// Memory Allocation:
//   - All operations allocate new slice for result
//   - Chunk: Allocates outer slice + n inner slices
//   - Take/Drop: Single allocation for result slice
//   - Sample: Allocates indices array + result slice
//
// Common Usage Patterns:
//
//	// Batch processing
//	items := []Task{...} // 1000 tasks
//	batches := sliceutil.Chunk(items, 50)
//	for _, batch := range batches {
//	    processBatch(batch) // Process 50 at a time
//	}
//
//	// Pagination
//	page := sliceutil.Slice(allItems, pageNum*pageSize, (pageNum+1)*pageSize)
//
//	// Take first valid items
//	valid := sliceutil.TakeWhile(items, func(item Item) bool {
//	    return item.IsValid()
//	})
//
//	// Random sampling for testing
//	testData := sliceutil.Sample(allData, 100)
//
//	// Merge parallel data sources
//	merged := sliceutil.Interleave(stream1, stream2, stream3)
//
// Thread Safety:
//   - Sample uses global RNG with mutex (thread-safe)
//   - Other functions are pure (no shared state)
//   - Safe for concurrent use with different slices
//
// slice.go는 슬라이스의 슬라이싱, 샘플링 및 분할 작업을 제공합니다.
//
// 이 파일은 슬라이스의 연속적이거나 샘플링된 부분을 추출, 분할 및 조작하는
// 작업을 구현합니다:
//
// 슬라이싱 작업:
//
// Chunk (동일한 부분으로 분할):
//   - Chunk(slice, size): 고정 크기 청크로 분할
//     시간: O(n), 공간: O(n)
//     슬라이스가 균등하게 나누어지지 않으면 마지막 청크가 더 작을 수 있음
//     [][]T (슬라이스의 슬라이스) 반환
//     용도: 배치 처리, 페이지네이션, 데이터 분할
//     예: [1,2,3,4,5,6,7] 3 크기 청크 → [[1,2,3], [4,5,6], [7]]
//
// Take/Drop (접두사/접미사 작업):
//   - Take(slice, n): 첫 n개 요소 가져오기
//     시간: O(n), 공간: O(n)
//     복사본 반환, 수정 안전
//     n > len이면 전체 슬라이스 반환
//   - TakeLast(slice, n): 마지막 n개 요소 가져오기
//     시간: O(n), 공간: O(n)
//     접미사 추출
//   - Drop(slice, n): 첫 n개 요소 제거
//     시간: O(n), 공간: O(n)
//     나머지 요소 반환
//   - DropLast(slice, n): 마지막 n개 요소 제거
//     시간: O(n), 공간: O(n)
//     접두사 반환
//
// 조건부 Take/Drop:
//   - TakeWhile(slice, predicate): 조건이 유지되는 동안 요소 가져오기
//     시간: O(k) (k = 가져온 요소 수), 공간: O(k)
//     첫 번째 false 조건자에서 중지
//     예: TakeWhile([1,2,3,4,1,2], n < 4) → [1,2,3]
//   - DropWhile(slice, predicate): 조건이 유지되는 동안 요소 건너뛰기
//     시간: O(k) (k = 건너뛴 요소 수), 공간: O(n-k)
//     첫 번째 false 후 나머지 반환
//     예: DropWhile([1,2,3,4,5], n < 4) → [4,5]
//
// 제네릭 슬라이싱:
//   - Slice(slice, start, end): 임의 범위 추출
//     시간: O(end-start), 공간: O(end-start)
//     Python과 유사한 슬라이싱 [start:end]
//     음수 인덱스 처리 (끝에서부터)
//     범위 밖 인덱스 클램핑
//     범위 복사본 반환
//
// 샘플링:
//   - Sample(slice, n): n개 요소의 무작위 샘플
//     시간: O(n), 공간: O(n)
//     비복원 추출 (고유 요소)
//     n >= len이면 섞인 복사본 반환
//     전역 스레드 안전 RNG 사용
//     용도: 무작위 선택, 테스트, 데이터 샘플링
//
// 인터리빙:
//   - Interleave(slices...): 여러 슬라이스를 라운드 로빈으로 병합
//     시간: O(총 요소), 공간: O(총 요소)
//     각 슬라이스에서 순서대로 하나씩 가져옴
//     모든 슬라이스가 소진될 때까지 계속
//     예: [1,2,3] + [a,b] + [x,y,z] → [1,a,x,2,b,y,3,z]
//     용도: 병렬 스트림 병합, 공정한 스케줄링
//
// 설계 원칙:
//   - 모든 함수는 새 슬라이스 반환 (불변 작업)
//   - 입력에 영향을 주지 않고 반환된 슬라이스 수정 안전
//   - 엣지 케이스의 일관된 처리 (빈, nil, 과다 n)
//   - Take/Drop 작업: n <= 0은 빈 반환, n >= len은 전체/빈 반환
//   - 조건부 작업: 조건자는 왼쪽에서 오른쪽으로 평가
//
// 엣지 케이스 처리:
//   - 빈 슬라이스: 빈 결과 반환 (패닉 없음)
//   - Nil 슬라이스: 빈 것으로 취급
//   - 음수 n: Take/Drop에 대해 0으로 취급
//   - 과다 n: 슬라이스 길이로 클램핑
//   - 음수 인덱스가 있는 슬라이스: -1 = 마지막 요소, -2 = 마지막에서 두 번째 등
//
// 성능 특성:
//   - Chunk: 단일 패스, 청크 할당
//   - Take/Drop: 복사 작업, 결과 크기에 선형
//   - TakeWhile/DropWhile: 조기 종료 가능 (최선 O(1))
//   - Sample: 셔플 알고리즘 (Fisher-Yates), 항상 O(n)
//   - Interleave: 모든 입력 슬라이스에 대한 단일 패스
//
// 메모리 할당:
//   - 모든 작업은 결과를 위해 새 슬라이스 할당
//   - Chunk: 외부 슬라이스 + n개 내부 슬라이스 할당
//   - Take/Drop: 결과 슬라이스를 위한 단일 할당
//   - Sample: 인덱스 배열 + 결과 슬라이스 할당
//
// 일반적인 사용 패턴:
//
//	// 배치 처리
//	items := []Task{...} // 1000개 작업
//	batches := sliceutil.Chunk(items, 50)
//	for _, batch := range batches {
//	    processBatch(batch) // 한 번에 50개 처리
//	}
//
//	// 페이지네이션
//	page := sliceutil.Slice(allItems, pageNum*pageSize, (pageNum+1)*pageSize)
//
//	// 첫 번째 유효 항목 가져오기
//	valid := sliceutil.TakeWhile(items, func(item Item) bool {
//	    return item.IsValid()
//	})
//
//	// 테스트를 위한 무작위 샘플링
//	testData := sliceutil.Sample(allData, 100)
//
//	// 병렬 데이터 소스 병합
//	merged := sliceutil.Interleave(stream1, stream2, stream3)
//
// 스레드 안전성:
//   - Sample은 뮤텍스가 있는 전역 RNG 사용 (스레드 안전)
//   - 다른 함수는 순수 (공유 상태 없음)
//   - 다른 슬라이스로 동시 사용 안전

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
