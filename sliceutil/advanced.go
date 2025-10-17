package sliceutil

// advanced.go provides advanced functional programming and transformation operations.
//
// This file implements sophisticated slice operations inspired by functional programming
// languages, offering powerful tools for data transformation and manipulation:
//
// Advanced Operations:
//
// Scan (Cumulative Reduce):
//   - Scan(slice, initial, accumulator): Track all intermediate reduction steps
//     Time: O(n), Space: O(n)
//     Like Reduce but returns all intermediate values
//     First element is initial value, then successive accumulated results
//     Useful for: Running totals, cumulative sums, factorial sequences
//     Example: [1,2,3,4] with sum → [0,1,3,6,10]
//
// ZipWith (Custom Combining):
//   - ZipWith(a, b, zipper): Combine two slices with custom function
//     Time: O(min(len(a), len(b))), Space: O(min(len(a), len(b)))
//     More flexible than Zip - custom merge logic
//     Result length = shorter input slice
//     Useful for: Element-wise operations, merging data structures
//     Example: [1,2,3] + ["a","b","c"] → ["1:a","2:b","3:c"]
//
// RotateLeft (Circular Shift):
//   - RotateLeft(slice, n): Shift elements left circularly
//     Time: O(n), Space: O(n)
//     Elements falling off left reappear on right
//     Handles negative n (rotates right)
//     Normalizes n to [0, len) range
//     Useful for: Circular buffers, round-robin scheduling
//     Example: [1,2,3,4,5] rotated 2 → [3,4,5,1,2]
//
// RotateRight (Circular Shift):
//   - RotateRight(slice, n): Shift elements right circularly
//     Time: O(n), Space: O(n)
//     Elements falling off right reappear on left
//     Equivalent to RotateLeft(slice, -n)
//     Useful for: Shifting windows, cyclic permutations
//     Example: [1,2,3,4,5] rotated 2 → [4,5,1,2,3]
//
// Key Features:
//   - Functional programming style with higher-order functions
//   - Immutable operations (original slice unchanged)
//   - Generic implementations for any type
//   - Efficient algorithms with minimal allocations
//
// Scan vs Reduce:
//   - Reduce: Returns single final result
//   - Scan: Returns slice of all intermediate results
//   - Scan result length = input length + 1 (includes initial)
//
// ZipWith vs Zip:
//   - Zip: Combines into tuples/pairs (fixed structure)
//   - ZipWith: Custom combination function (flexible result type)
//   - Both stop at shorter slice length
//
// Rotation Behavior:
//   - Both rotations handle edge cases:
//     * Empty slice → empty result
//     * n = 0 → clone of original
//     * n >= len → modulo wrap-around
//     * Negative n → opposite direction
//   - Rotations are pure functions (no mutation)
//
// Performance Characteristics:
//   - Scan: Single pass, linear time, allocates result array
//   - ZipWith: Single pass over shorter slice, linear time
//   - RotateLeft/Right: Two slice operations (split + concatenate)
//   - All operations: One allocation for result
//
// Common Use Cases:
//
//	// Running totals for analytics
//	sales := []float64{100, 150, 200, 50}
//	runningTotal := sliceutil.Scan(sales, 0.0, func(sum, sale float64) float64 {
//	    return sum + sale
//	})
//	// [0, 100, 250, 450, 500]
//
//	// Merge parallel data streams
//	ids := []int{1, 2, 3}
//	names := []string{"Alice", "Bob", "Charlie"}
//	users := sliceutil.ZipWith(ids, names, func(id int, name string) User {
//	    return User{ID: id, Name: name}
//	})
//
//	// Circular buffer operations
//	buffer := []byte{1, 2, 3, 4, 5}
//	shifted := sliceutil.RotateLeft(buffer, 2)  // [3,4,5,1,2]
//
//	// Round-robin task distribution
//	workers := []string{"W1", "W2", "W3"}
//	nextShift := sliceutil.RotateLeft(workers, 1)
//
// Thread Safety:
//   - All functions are thread-safe (pure, no shared state)
//   - Input slices not modified
//   - Safe for concurrent use with different slices
//
// advanced.go는 고급 함수형 프로그래밍 및 변환 작업을 제공합니다.
//
// 이 파일은 함수형 프로그래밍 언어에서 영감을 받은 정교한 슬라이스 작업을
// 구현하여 데이터 변환 및 조작을 위한 강력한 도구를 제공합니다:
//
// 고급 작업:
//
// Scan (누적 Reduce):
//   - Scan(slice, initial, accumulator): 모든 중간 축소 단계 추적
//     시간: O(n), 공간: O(n)
//     Reduce와 비슷하지만 모든 중간 값 반환
//     첫 번째 요소는 초기값, 그 다음 연속 누적 결과
//     용도: 누적 합계, 팩토리얼 시퀀스
//     예: [1,2,3,4] 합계로 → [0,1,3,6,10]
//
// ZipWith (커스텀 결합):
//   - ZipWith(a, b, zipper): 커스텀 함수로 두 슬라이스 결합
//     시간: O(min(len(a), len(b))), 공간: O(min(len(a), len(b)))
//     Zip보다 유연 - 커스텀 병합 로직
//     결과 길이 = 더 짧은 입력 슬라이스
//     용도: 요소별 연산, 데이터 구조 병합
//     예: [1,2,3] + ["a","b","c"] → ["1:a","2:b","3:c"]
//
// RotateLeft (순환 시프트):
//   - RotateLeft(slice, n): 요소를 왼쪽으로 순환 이동
//     시간: O(n), 공간: O(n)
//     왼쪽으로 떨어진 요소가 오른쪽에 다시 나타남
//     음수 n 처리 (오른쪽 회전)
//     n을 [0, len) 범위로 정규화
//     용도: 순환 버퍼, 라운드 로빈 스케줄링
//     예: [1,2,3,4,5] 2 회전 → [3,4,5,1,2]
//
// RotateRight (순환 시프트):
//   - RotateRight(slice, n): 요소를 오른쪽으로 순환 이동
//     시간: O(n), 공간: O(n)
//     오른쪽으로 떨어진 요소가 왼쪽에 다시 나타남
//     RotateLeft(slice, -n)과 동등
//     용도: 윈도우 이동, 순환 순열
//     예: [1,2,3,4,5] 2 회전 → [4,5,1,2,3]
//
// 주요 기능:
//   - 고차 함수를 사용한 함수형 프로그래밍 스타일
//   - 불변 작업 (원본 슬라이스 변경 안 함)
//   - 모든 타입에 대한 제네릭 구현
//   - 최소 할당으로 효율적인 알고리즘
//
// Scan vs Reduce:
//   - Reduce: 단일 최종 결과 반환
//   - Scan: 모든 중간 결과의 슬라이스 반환
//   - Scan 결과 길이 = 입력 길이 + 1 (초기값 포함)
//
// ZipWith vs Zip:
//   - Zip: 튜플/쌍으로 결합 (고정 구조)
//   - ZipWith: 커스텀 결합 함수 (유연한 결과 타입)
//   - 둘 다 더 짧은 슬라이스 길이에서 중지
//
// 회전 동작:
//   - 두 회전 모두 엣지 케이스 처리:
//     * 빈 슬라이스 → 빈 결과
//     * n = 0 → 원본 복제
//     * n >= len → 모듈로 랩어라운드
//     * 음수 n → 반대 방향
//   - 회전은 순수 함수 (변이 없음)
//
// 성능 특성:
//   - Scan: 단일 패스, 선형 시간, 결과 배열 할당
//   - ZipWith: 더 짧은 슬라이스에 대한 단일 패스, 선형 시간
//   - RotateLeft/Right: 두 슬라이스 작업 (분할 + 연결)
//   - 모든 작업: 결과에 대한 하나의 할당
//
// 일반적인 사용 사례:
//
//	// 분석을 위한 누적 합계
//	sales := []float64{100, 150, 200, 50}
//	runningTotal := sliceutil.Scan(sales, 0.0, func(sum, sale float64) float64 {
//	    return sum + sale
//	})
//	// [0, 100, 250, 450, 500]
//
//	// 병렬 데이터 스트림 병합
//	ids := []int{1, 2, 3}
//	names := []string{"Alice", "Bob", "Charlie"}
//	users := sliceutil.ZipWith(ids, names, func(id int, name string) User {
//	    return User{ID: id, Name: name}
//	})
//
//	// 순환 버퍼 작업
//	buffer := []byte{1, 2, 3, 4, 5}
//	shifted := sliceutil.RotateLeft(buffer, 2)  // [3,4,5,1,2]
//
//	// 라운드 로빈 작업 배포
//	workers := []string{"W1", "W2", "W3"}
//	nextShift := sliceutil.RotateLeft(workers, 1)
//
// 스레드 안전성:
//   - 모든 함수는 스레드 안전 (순수, 공유 상태 없음)
//   - 입력 슬라이스 수정 안 함
//   - 다른 슬라이스로 동시 사용 안전

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
