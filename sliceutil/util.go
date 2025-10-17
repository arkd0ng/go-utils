package sliceutil

import (
	"fmt"
	"strings"
)

// ForEach executes a function for each element in the slice.
// ForEach는 슬라이스의 각 요소에 대해 함수를 실행합니다.
//
// The function is executed for its side effects; ForEach does not return a value.
// 함수는 부수 효과를 위해 실행됩니다; ForEach는 값을 반환하지 않습니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	sliceutil.ForEach(numbers, func(n int) {
//	    fmt.Println(n * 2)
//	})
//	// Output: 2, 4, 6, 8, 10 (each on a new line)
func ForEach[T any](slice []T, fn func(T)) {
	for _, item := range slice {
		fn(item)
	}
}

// ForEachIndexed executes a function for each element in the slice with its index.
// ForEachIndexed는 슬라이스의 각 요소와 인덱스에 대해 함수를 실행합니다.
//
// The function receives both the index and the element.
// 함수는 인덱스와 요소를 모두 받습니다.
//
// Example
// 예제:
//
//	words := []string{"apple", "banana", "cherry"}
//	sliceutil.ForEachIndexed(words, func(i int, word string) {
//	    fmt.Printf("%d: %s\n", i, word)
//	})
//	// Output:
//	// 0: apple
//	// 1: banana
//	// 2: cherry
func ForEachIndexed[T any](slice []T, fn func(int, T)) {
	for i, item := range slice {
		fn(i, item)
	}
}

// Join converts all elements to strings and joins them with the separator.
// Join은 모든 요소를 문자열로 변환하고 구분자로 결합합니다.
//
// Uses fmt.Sprint to convert elements to strings.
// fmt.Sprint를 사용하여 요소를 문자열로 변환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := sliceutil.Join(numbers, ", ")
//	// result: "1, 2, 3, 4, 5"
//
//	words := []string{"apple", "banana", "cherry"}
//	result2 := sliceutil.Join(words, "-")
//	// result2: "apple-banana-cherry"
func Join[T any](slice []T, separator string) string {
	if len(slice) == 0 {
		return ""
	}

	// Convert all elements to strings
	// 모든 요소를 문자열로 변환
	strSlice := make([]string, len(slice))
	for i, item := range slice {
		strSlice[i] = fmt.Sprint(item)
	}

	return strings.Join(strSlice, separator)
}

// Clone creates a shallow copy of the slice.
// Clone은 슬라이스의 얕은 복사본을 생성합니다.
//
// The returned slice has the same elements but is a different underlying array.
// 반환된 슬라이스는 동일한 요소를 가지지만 다른 기본 배열입니다.
//
// Example
// 예제:
//
//	original := []int{1, 2, 3, 4, 5}
//	cloned := sliceutil.Clone(original)
//
// cloned[0] = 99
// original: [1, 2, 3, 4, 5] (unchanged
// 변경되지 않음)
//
//	// cloned: [99, 2, 3, 4, 5]
func Clone[T any](slice []T) []T {
	if slice == nil {
		return nil
	}

	result := make([]T, len(slice))
	copy(result, slice)
	return result
}

// Fill replaces all elements in the slice with the specified value.
// Fill은 슬라이스의 모든 요소를 지정된 값으로 바꿉니다.
//
// Returns a new slice with all elements set to the value.
// 모든 요소가 값으로 설정된 새 슬라이스를 반환합니다.
//
// Example
// 예제:
//
//	slice := []int{1, 2, 3, 4, 5}
//	filled := sliceutil.Fill(slice, 0)
//
// // filled: [0, 0, 0, 0, 0]
// slice: [1, 2, 3, 4, 5] (unchanged
// 변경되지 않음)
func Fill[T any](slice []T, value T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	result := make([]T, len(slice))
	for i := range result {
		result[i] = value
	}
	return result
}

// Insert inserts items at the specified index.
// Insert는 지정된 인덱스에 항목을 삽입합니다.
//
// Returns a new slice with items inserted at the index.
// 인덱스에 항목이 삽입된 새 슬라이스를 반환합니다.
//
// If index is negative or greater than slice length, items are appended.
// 인덱스가 음수이거나 슬라이스 길이보다 크면 항목이 추가됩니다.
//
// Example
// 예제:
//
//	slice := []int{1, 2, 5, 6}
//	result := sliceutil.Insert(slice, 2, 3, 4)
//	// result: [1, 2, 3, 4, 5, 6]
//
//	result2 := sliceutil.Insert(slice, 0, 0)
//	// result2: [0, 1, 2, 5, 6]
func Insert[T any](slice []T, index int, items ...T) []T {
	if len(items) == 0 {
		return Clone(slice)
	}

	// Clamp index to valid range
	// 인덱스를 유효한 범위로 제한
	if index < 0 {
		index = 0
	}
	if index > len(slice) {
		index = len(slice)
	}

	// Create new slice with enough capacity
	// 충분한 용량으로 새 슬라이스 생성
	result := make([]T, 0, len(slice)+len(items))
	result = append(result, slice[:index]...)
	result = append(result, items...)
	result = append(result, slice[index:]...)

	return result
}

// Remove removes the element at the specified index.
// Remove는 지정된 인덱스의 요소를 제거합니다.
//
// Returns a new slice with the element removed.
// 요소가 제거된 새 슬라이스를 반환합니다.
//
// If index is out of bounds, returns a copy of the original slice.
// 인덱스가 범위를 벗어나면 원본 슬라이스의 복사본을 반환합니다.
//
// Example
// 예제:
//
//	slice := []int{1, 2, 3, 4, 5}
//	result := sliceutil.Remove(slice, 2)
//
// // result: [1, 2, 4, 5]
// slice: [1, 2, 3, 4, 5] (unchanged
// 변경되지 않음)
func Remove[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return Clone(slice)
	}

	result := make([]T, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)

	return result
}

// RemoveAll removes all occurrences of the specified item.
// RemoveAll은 지정된 항목의 모든 발생을 제거합니다.
//
// Returns a new slice with all occurrences of the item removed.
// 항목의 모든 발생이 제거된 새 슬라이스를 반환합니다.
//
// Example
// 예제:
//
//	slice := []int{1, 2, 3, 2, 4, 2, 5}
//	result := sliceutil.RemoveAll(slice, 2)
//	// result: [1, 3, 4, 5]
//
//	words := []string{"apple", "banana", "apple", "cherry"}
//	result2 := sliceutil.RemoveAll(words, "apple")
//	// result2: ["banana", "cherry"]
func RemoveAll[T comparable](slice []T, item T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != item {
			result = append(result, v)
		}
	}
	return result
}

// Shuffle returns a new slice with elements in random order.
// Shuffle은 요소가 무작위 순서로 있는 새 슬라이스를 반환합니다.
//
// Uses a default random source seeded with the current time.
// 현재 시간으로 시드된 기본 랜덤 소스를 사용합니다.
//
// The original slice is not modified.
// 원본 슬라이스는 수정되지 않습니다.
//
// Example
// 예제:
//
//	slice := []int{1, 2, 3, 4, 5}
//
// shuffled := sliceutil.Shuffle(slice)
// shuffled: [3, 1, 5, 2, 4] (random order
// 무작위 순서)
// // slice: [1, 2, 3, 4, 5] (unchanged
// 변경되지 않음)
func Shuffle[T any](slice []T) []T {
	if len(slice) <= 1 {
		return Clone(slice)
	}

	result := Clone(slice)
	rngLock.Lock()
	defer rngLock.Unlock()

	// Fisher-Yates shuffle algorithm
	// Fisher-Yates 셔플 알고리즘
	for i := len(result) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// Zip combines two slices into a slice of pairs.
// Zip은 두 슬라이스를 쌍의 슬라이스로 결합합니다.
//
// Returns a slice of [2]any where each element is a pair of elements from the two slices.
// 각 요소가 두 슬라이스의 요소 쌍인 [2]any의 슬라이스를 반환합니다.
//
// The resulting slice length is the minimum of the two input slices.
// 결과 슬라이스 길이는 두 입력 슬라이스의 최소값입니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3}
//	words := []string{"one", "two", "three"}
//	zipped := sliceutil.Zip(numbers, words)
//	// zipped: [[1, "one"], [2, "two"], [3, "three"]]
func Zip[T, U any](a []T, b []U) [][2]any {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	result := make([][2]any, minLen)
	for i := 0; i < minLen; i++ {
		result[i] = [2]any{a[i], b[i]}
	}

	return result
}

// Unzip separates a slice of pairs into two slices.
// Unzip은 쌍의 슬라이스를 두 슬라이스로 분리합니다.
//
// Returns two slices: one with first elements and one with second elements.
// 두 개의 슬라이스를 반환합니다: 첫 번째 요소를 가진 슬라이스와 두 번째 요소를 가진 슬라이스.
//
// IMPORTANT: Type assertions will panic if the slice contains elements
// that are not of types T and U. Ensure all pairs are correctly typed before calling this function.
//
// 중요: 슬라이스에 T 및 U 타입이 아닌 요소가 포함되어 있으면 타입 단언이 패닉을 발생시킵니다.
// 이 함수를 호출하기 전에 모든 쌍이 올바르게 타입이 지정되었는지 확인하세요.
//
// Example
// 예제:
//
// // ✅ CORRECT usage
// 올바른 사용:
//
//	zipped := [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}
//	numbers, words := sliceutil.Unzip[int, string](zipped)
//	// numbers: [1, 2, 3]
//	// words: ["one", "two", "three"]
//
// // ❌ INCORRECT usage (will panic!)
// 잘못된 사용 (패닉 발생!):
//
//	badZipped := [][2]any{{1, "one"}, {"wrong", 2}} // Wrong types / 잘못된 타입
//	nums, words := sliceutil.Unzip[int, string](badZipped) // PANIC!
func Unzip[T, U any](slice [][2]any) ([]T, []U) {
	if len(slice) == 0 {
		return []T{}, []U{}
	}

	first := make([]T, len(slice))
	second := make([]U, len(slice))

	for i, pair := range slice {
		first[i] = pair[0].(T)
		second[i] = pair[1].(U)
	}

	return first, second
}

// Window returns a slice of sliding windows of the specified size.
// Window는 지정된 크기의 슬라이딩 윈도우 슬라이스를 반환합니다.
//
// Returns a slice of slices, where each sub-slice is a window of the specified size.
// 각 하위 슬라이스가 지정된 크기의 윈도우인 슬라이스의 슬라이스를 반환합니다.
//
// If size is less than or equal to 0, or greater than slice length, returns empty slice.
// size가 0 이하이거나 슬라이스 길이보다 크면 빈 슬라이스를 반환합니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	windows := sliceutil.Window(numbers, 3)
//	// windows: [[1, 2, 3], [2, 3, 4], [3, 4, 5]]
//
//	words := []string{"a", "b", "c", "d"}
//	windows2 := sliceutil.Window(words, 2)
//	// windows2: [["a", "b"], ["b", "c"], ["c", "d"]]
func Window[T any](slice []T, size int) [][]T {
	if size <= 0 || size > len(slice) {
		return [][]T{}
	}

	numWindows := len(slice) - size + 1
	result := make([][]T, numWindows)

	for i := 0; i < numWindows; i++ {
		window := make([]T, size)
		copy(window, slice[i:i+size])
		result[i] = window
	}

	return result
}

// Tap executes a function on the slice and returns the slice unchanged.
// Tap은 슬라이스에 함수를 실행하고 슬라이스를 변경하지 않고 반환합니다.
//
// Useful for debugging or side effects in method chains.
// 메서드 체인에서 디버깅이나 부수 효과에 유용합니다.
//
// The function receives the entire slice and can perform any operation,
// but the original slice is returned unchanged.
// 함수는 전체 슬라이스를 받고 모든 작업을 수행할 수 있지만
// 원본 슬라이스는 변경되지 않고 반환됩니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := sliceutil.Tap(numbers, func(s []int) {
//	    fmt.Printf("Current slice: %v\n", s)
//	})
//	// Output: Current slice: [1 2 3 4 5]
//	// result: [1, 2, 3, 4, 5]
//
// // Useful in chains
// 체인에서 유용
//
//	result2 := sliceutil.Map(
//	    sliceutil.Tap(
//	        sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 }),
//	        func(s []int) { fmt.Printf("Filtered: %v\n", s) },
//	    ),
//	    func(n int) int { return n * 2 },
//	)
//	// Output: Filtered: [2 4]
//	// result2: [4, 8]
func Tap[T any](slice []T, fn func([]T)) []T {
	fn(slice)
	return slice
}
