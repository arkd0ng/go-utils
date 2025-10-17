package sliceutil

import (
	"fmt"
	"strings"
)

// util.go provides utility operations for slices.
//
// This file implements miscellaneous utility functions for iteration,
// manipulation, transformation, and debugging of slices:
//
// Iteration Operations:
//
// ForEach (Side Effect Iteration):
//   - ForEach(slice, fn): Execute function for each element
//     Time: O(n), Space: O(1)
//     No return value (side effects only)
//     Function called once per element in order
//     Cannot break early (processes all elements)
//     Example: ForEach([1,2,3], print) prints each element
//     Use cases: Printing, logging, mutation outside slice
//
// ForEachIndexed (Indexed Iteration):
//   - ForEachIndexed(slice, fn): Execute function with index
//     Time: O(n), Space: O(1)
//     Function receives (index, element) pairs
//     Useful when position matters
//     Example: ForEachIndexed(items, func(i, item) { fmt.Printf("%d: %v\n", i, item) })
//     Use cases: Enumeration, index-aware operations, numbered output
//
// String Operations:
//
// Join (Concatenate to String):
//   - Join(slice, separator): Convert elements to strings and join
//     Time: O(n), Space: O(n)
//     Uses fmt.Sprint to convert each element
//     Works with any type (not just strings)
//     Similar to strings.Join but type-generic
//     Example: Join([1,2,3], ", ") = "1, 2, 3"
//     Use cases: String representation, CSV output, display formatting
//
// Slice Manipulation:
//
// Clone (Shallow Copy):
//   - Clone(slice): Create independent copy
//     Time: O(n), Space: O(n)
//     Shallow copy: New slice, same elements
//     Nil input returns nil (preserves nil semantics)
//     Safe to modify result without affecting original
//     Example: Clone([1,2,3]) = [1,2,3] (different backing array)
//     Use cases: Immutable operations, defensive copying, fork data
//
// Fill (Uniform Value):
//   - Fill(slice, value): Replace all elements with value
//     Time: O(n), Space: O(n)
//     Returns new slice (immutable)
//     All positions set to same value
//     Empty input returns empty slice
//     Example: Fill([1,2,3,4,5], 0) = [0,0,0,0,0]
//     Use cases: Initialization, reset, placeholder creation
//
// Insert (Add Elements):
//   - Insert(slice, index, items...): Insert items at position
//     Time: O(n), Space: O(n+m) where m = len(items)
//     Returns new slice with items inserted
//     Index clamped to [0, len(slice)]
//     Negative index treated as 0
//     Index > len treated as append
//     Example: Insert([1,2,5,6], 2, 3, 4) = [1,2,3,4,5,6]
//     Use cases: Adding elements mid-slice, building sequences
//
// Remove (Delete by Index):
//   - Remove(slice, index): Remove element at position
//     Time: O(n), Space: O(n)
//     Returns new slice without element
//     Out-of-bounds index returns clone (no error)
//     Example: Remove([1,2,3,4,5], 2) = [1,2,4,5]
//     Use cases: Deleting specific position, index-based removal
//
// RemoveAll (Delete by Value):
//   - RemoveAll(slice, item): Remove all occurrences
//     Time: O(n), Space: O(n)
//     Removes every matching element
//     Returns new slice (immutable)
//     Example: RemoveAll([1,2,3,2,4,2], 2) = [1,3,4]
//     Use cases: Filtering specific value, cleanup
//
// Randomization:
//
// Shuffle (Random Order):
//   - Shuffle(slice): Randomize element order
//     Time: O(n), Space: O(n)
//     Uses Fisher-Yates shuffle algorithm
//     Cryptographically weak randomness
//     Thread-safe (uses global RNG with lock)
//     Returns new slice (immutable)
//     Example: Shuffle([1,2,3,4,5]) = [3,1,5,2,4] (random)
//     Use cases: Card games, random sampling, lottery selection
//
// Pairing Operations:
//
// Zip (Combine Two Slices):
//   - Zip(a, b): Create pairs from two slices
//     Time: O(min(n,m)), Space: O(min(n,m))
//     Result length = min(len(a), len(b))
//     Longer slice truncated
//     Returns [][2]any (type erasure for flexibility)
//     Example: Zip([1,2,3], ["a","b","c"]) = [[1,"a"], [2,"b"], [3,"c"]]
//     Use cases: Parallel iteration, dictionary construction, coordinate pairs
//
// Unzip (Separate Pairs):
//   - Unzip(slice): Split pairs into two slices
//     Time: O(n), Space: O(n)
//     Inverse of Zip
//     Type assertions may panic if types wrong
//     Must specify types explicitly: Unzip[int, string](pairs)
//     Example: Unzip([[1,"a"], [2,"b"]]) = ([1,2], ["a","b"])
//     Use cases: Unpacking paired data, separating coordinates
//
// Window Operations:
//
// Window (Sliding Windows):
//   - Window(slice, size): Generate overlapping subsequences
//     Time: O(n*size), Space: O(n*size)
//     Each window is independent copy
//     Number of windows = len(slice) - size + 1
//     Invalid size returns empty slice
//     Example: Window([1,2,3,4,5], 3) = [[1,2,3], [2,3,4], [3,4,5]]
//     Use cases: Moving averages, pattern detection, n-gram analysis
//
// Debugging:
//
// Tap (Debug Inspection):
//   - Tap(slice, fn): Execute function and return slice unchanged
//     Time: O(1) + fn execution, Space: O(1)
//     Passthrough operation for side effects
//     Function receives whole slice
//     Original slice returned (not copy)
//     Useful in method chains for debugging
//     Example: Tap(data, func(s) { fmt.Println(s) }) returns data
//     Use cases: Debugging chains, logging, inspection, breakpoints
//
// Design Principles:
//   - Immutability: Most operations return new slices
//   - Consistency: Predictable behavior for edge cases
//   - Flexibility: Generic implementations for any type
//   - Safety: Bounds checking, nil handling
//   - Utility: Commonly needed but not category-specific
//
// Comparison: ForEach vs Map/Filter:
//   - ForEach: Side effects only, no return value
//   - Map/Filter: Transform data, return new slice
//   - ForEach: Use when result doesn't matter (logging, printing)
//   - Map/Filter: Use when building new data structure
//
// Comparison: Remove vs RemoveAll:
//   - Remove: Single element by index (position-based)
//   - RemoveAll: All occurrences by value (value-based)
//   - Remove: O(n), removes exactly one
//   - RemoveAll: O(n), removes zero or more
//
// Comparison: Clone vs Copy:
//   - Clone: Allocates new slice, copies data
//   - copy(): Requires pre-allocated destination
//   - Clone: More convenient, returns new slice
//   - copy(): More efficient if destination exists
//
// Performance Characteristics:
//
// Time Complexity:
//   - ForEach/ForEachIndexed: O(n)
//   - Join: O(n) string conversions + concatenation
//   - Clone/Fill: O(n)
//   - Insert/Remove/RemoveAll: O(n)
//   - Shuffle: O(n)
//   - Zip/Unzip: O(n)
//   - Window: O(n * size)
//   - Tap: O(1) + function cost
//
// Space Complexity:
//   - ForEach/ForEachIndexed: O(1)
//   - Join: O(n) for result string
//   - Clone/Fill/Remove/RemoveAll: O(n)
//   - Insert: O(n+m)
//   - Shuffle: O(n)
//   - Zip/Unzip: O(n)
//   - Window: O(n * size)
//   - Tap: O(1)
//
// Common Usage Patterns:
//
//	// Print all elements
//	sliceutil.ForEach(items, func(item Item) {
//	    fmt.Printf("Item: %v\n", item)
//	})
//
//	// Create display string
//	display := sliceutil.Join(tags, ", ")
//	fmt.Printf("Tags: %s\n", display)
//
//	// Safe modification
//	modified := sliceutil.Clone(original)
//	// Modify 'modified' without affecting 'original'
//
//	// Initialize array
//	zeros := sliceutil.Fill(template, 0)
//
//	// Build sequence
//	data := []int{1, 2, 5, 6}
//	complete := sliceutil.Insert(data, 2, 3, 4)
//
//	// Shuffle deck
//	shuffled := sliceutil.Shuffle(cards)
//
//	// Combine data
//	keys := []string{"name", "age", "city"}
//	values := []any{"Alice", 30, "NYC"}
//	pairs := sliceutil.Zip(keys, values)
//
//	// Sliding analysis
//	temps := []float64{20, 21, 22, 23, 24, 25}
//	windows := sliceutil.Window(temps, 3)
//	for _, w := range windows {
//	    avg := sliceutil.Average(w)
//	    fmt.Printf("Window avg: %.1f\n", avg)
//	}
//
//	// Debug chain
//	result := sliceutil.Map(
//	    sliceutil.Tap(
//	        sliceutil.Filter(data, isValid),
//	        func(s []int) { log.Printf("After filter: %v", s) },
//	    ),
//	    transform,
//	)
//
// Method Chaining Patterns:
//
//	// Use Tap for debugging chains
//	result := sliceutil.Map(
//	    sliceutil.Tap(sliceutil.Filter(data, pred), log),
//	    mapper,
//	)
//
//	// Clone before mutation
//	safe := sliceutil.Clone(original)
//	// Mutate 'safe' freely
//
// Edge Cases:
//   - ForEach on empty: No-op (function never called)
//   - Join on empty: Returns "" (empty string)
//   - Clone of nil: Returns nil (preserves nil semantics)
//   - Fill of empty: Returns []T{} (empty slice)
//   - Insert with invalid index: Clamped to valid range
//   - Remove with invalid index: Returns clone
//   - Shuffle of <=1: Returns clone (nothing to shuffle)
//   - Zip with different lengths: Uses minimum length
//   - Window with invalid size: Returns [][]T{} (empty)
//
// Thread Safety:
//   - Shuffle: Thread-safe (uses locked global RNG)
//   - Other operations: Not thread-safe if slice modified concurrently
//   - Clone creates independent copy (safe from original's mutations)
//
// util.go는 슬라이스에 대한 유틸리티 작업을 제공합니다.
//
// 이 파일은 슬라이스의 반복, 조작, 변환 및 디버깅을 위한
// 기타 유틸리티 함수를 구현합니다:
//
// 반복 작업:
//
// ForEach (부수 효과 반복):
//   - ForEach(slice, fn): 각 요소에 대해 함수 실행
//     시간: O(n), 공간: O(1)
//     반환 값 없음 (부수 효과만)
//     순서대로 요소당 한 번 함수 호출
//     조기 중단 불가 (모든 요소 처리)
//     예: ForEach([1,2,3], print) 각 요소 출력
//     사용 사례: 출력, 로깅, 슬라이스 외부 변경
//
// ForEachIndexed (인덱스 반복):
//   - ForEachIndexed(slice, fn): 인덱스와 함께 함수 실행
//     시간: O(n), 공간: O(1)
//     함수가 (인덱스, 요소) 쌍 수신
//     위치가 중요할 때 유용
//     예: ForEachIndexed(items, func(i, item) { fmt.Printf("%d: %v\n", i, item) })
//     사용 사례: 열거, 인덱스 인식 작업, 번호 출력
//
// 문자열 작업:
//
// Join (문자열로 연결):
//   - Join(slice, separator): 요소를 문자열로 변환하고 결합
//     시간: O(n), 공간: O(n)
//     각 요소를 변환하기 위해 fmt.Sprint 사용
//     모든 타입 작동 (문자열만이 아님)
//     strings.Join과 유사하지만 타입 제네릭
//     예: Join([1,2,3], ", ") = "1, 2, 3"
//     사용 사례: 문자열 표현, CSV 출력, 디스플레이 포맷
//
// 슬라이스 조작:
//
// Clone (얕은 복사):
//   - Clone(slice): 독립 복사본 생성
//     시간: O(n), 공간: O(n)
//     얕은 복사: 새 슬라이스, 같은 요소
//     Nil 입력은 nil 반환 (nil 의미론 유지)
//     원본에 영향 없이 결과 수정 안전
//     예: Clone([1,2,3]) = [1,2,3] (다른 백킹 배열)
//     사용 사례: 불변 작업, 방어적 복사, 데이터 포크
//
// Fill (균일 값):
//   - Fill(slice, value): 모든 요소를 값으로 교체
//     시간: O(n), 공간: O(n)
//     새 슬라이스 반환 (불변)
//     모든 위치를 같은 값으로 설정
//     빈 입력은 빈 슬라이스 반환
//     예: Fill([1,2,3,4,5], 0) = [0,0,0,0,0]
//     사용 사례: 초기화, 재설정, 플레이스홀더 생성
//
// Insert (요소 추가):
//   - Insert(slice, index, items...): 위치에 항목 삽입
//     시간: O(n), 공간: O(n+m) (m = len(items))
//     항목이 삽입된 새 슬라이스 반환
//     인덱스 [0, len(slice)]로 제한
//     음수 인덱스는 0으로 처리
//     len보다 큰 인덱스는 append로 처리
//     예: Insert([1,2,5,6], 2, 3, 4) = [1,2,3,4,5,6]
//     사용 사례: 슬라이스 중간에 요소 추가, 시퀀스 구축
//
// Remove (인덱스로 삭제):
//   - Remove(slice, index): 위치의 요소 제거
//     시간: O(n), 공간: O(n)
//     요소 없이 새 슬라이스 반환
//     범위 벗어난 인덱스는 복제 반환 (에러 없음)
//     예: Remove([1,2,3,4,5], 2) = [1,2,4,5]
//     사용 사례: 특정 위치 삭제, 인덱스 기반 제거
//
// RemoveAll (값으로 삭제):
//   - RemoveAll(slice, item): 모든 발생 제거
//     시간: O(n), 공간: O(n)
//     일치하는 모든 요소 제거
//     새 슬라이스 반환 (불변)
//     예: RemoveAll([1,2,3,2,4,2], 2) = [1,3,4]
//     사용 사례: 특정 값 필터링, 정리
//
// 무작위화:
//
// Shuffle (무작위 순서):
//   - Shuffle(slice): 요소 순서 무작위화
//     시간: O(n), 공간: O(n)
//     Fisher-Yates 셔플 알고리즘 사용
//     암호학적으로 약한 무작위성
//     스레드 안전 (잠금이 있는 전역 RNG 사용)
//     새 슬라이스 반환 (불변)
//     예: Shuffle([1,2,3,4,5]) = [3,1,5,2,4] (무작위)
//     사용 사례: 카드 게임, 무작위 샘플링, 복권 선택
//
// 쌍 작업:
//
// Zip (두 슬라이스 결합):
//   - Zip(a, b): 두 슬라이스에서 쌍 생성
//     시간: O(min(n,m)), 공간: O(min(n,m))
//     결과 길이 = min(len(a), len(b))
//     긴 슬라이스 잘림
//     [][2]any 반환 (유연성을 위한 타입 소거)
//     예: Zip([1,2,3], ["a","b","c"]) = [[1,"a"], [2,"b"], [3,"c"]]
//     사용 사례: 병렬 반복, 사전 구성, 좌표 쌍
//
// Unzip (쌍 분리):
//   - Unzip(slice): 쌍을 두 슬라이스로 분할
//     시간: O(n), 공간: O(n)
//     Zip의 역
//     타입 잘못되면 타입 단언이 패닉 가능
//     타입 명시적으로 지정 필요: Unzip[int, string](pairs)
//     예: Unzip([[1,"a"], [2,"b"]]) = ([1,2], ["a","b"])
//     사용 사례: 쌍 데이터 언팩, 좌표 분리
//
// 윈도우 작업:
//
// Window (슬라이딩 윈도우):
//   - Window(slice, size): 겹치는 부분 시퀀스 생성
//     시간: O(n*size), 공간: O(n*size)
//     각 윈도우는 독립 복사본
//     윈도우 수 = len(slice) - size + 1
//     잘못된 크기는 빈 슬라이스 반환
//     예: Window([1,2,3,4,5], 3) = [[1,2,3], [2,3,4], [3,4,5]]
//     사용 사례: 이동 평균, 패턴 감지, n-그램 분석
//
// 디버깅:
//
// Tap (디버그 검사):
//   - Tap(slice, fn): 함수 실행하고 슬라이스 변경 없이 반환
//     시간: O(1) + fn 실행, 공간: O(1)
//     부수 효과를 위한 패스스루 작업
//     함수가 전체 슬라이스 수신
//     원본 슬라이스 반환 (복사본 아님)
//     메서드 체인에서 디버깅에 유용
//     예: Tap(data, func(s) { fmt.Println(s) }) data 반환
//     사용 사례: 체인 디버깅, 로깅, 검사, 중단점
//
// 설계 원칙:
//   - 불변성: 대부분 작업이 새 슬라이스 반환
//   - 일관성: 엣지 케이스에 예측 가능한 동작
//   - 유연성: 모든 타입에 대한 제네릭 구현
//   - 안전성: 범위 확인, nil 처리
//   - 유틸리티: 일반적으로 필요하지만 카테고리별이 아님
//
// 비교: ForEach vs Map/Filter:
//   - ForEach: 부수 효과만, 반환 값 없음
//   - Map/Filter: 데이터 변환, 새 슬라이스 반환
//   - ForEach: 결과가 중요하지 않을 때 사용 (로깅, 출력)
//   - Map/Filter: 새 데이터 구조 구축 시 사용
//
// 비교: Remove vs RemoveAll:
//   - Remove: 인덱스로 단일 요소 (위치 기반)
//   - RemoveAll: 값으로 모든 발생 (값 기반)
//   - Remove: O(n), 정확히 하나 제거
//   - RemoveAll: O(n), 0개 이상 제거
//
// 비교: Clone vs Copy:
//   - Clone: 새 슬라이스 할당, 데이터 복사
//   - copy(): 사전 할당된 대상 필요
//   - Clone: 더 편리, 새 슬라이스 반환
//   - copy(): 대상 존재 시 더 효율적
//
// 성능 특성:
//
// 시간 복잡도:
//   - ForEach/ForEachIndexed: O(n)
//   - Join: O(n) 문자열 변환 + 연결
//   - Clone/Fill: O(n)
//   - Insert/Remove/RemoveAll: O(n)
//   - Shuffle: O(n)
//   - Zip/Unzip: O(n)
//   - Window: O(n * size)
//   - Tap: O(1) + 함수 비용
//
// 공간 복잡도:
//   - ForEach/ForEachIndexed: O(1)
//   - Join: 결과 문자열을 위한 O(n)
//   - Clone/Fill/Remove/RemoveAll: O(n)
//   - Insert: O(n+m)
//   - Shuffle: O(n)
//   - Zip/Unzip: O(n)
//   - Window: O(n * size)
//   - Tap: O(1)
//
// 일반적인 사용 패턴:
//
//	// 모든 요소 출력
//	sliceutil.ForEach(items, func(item Item) {
//	    fmt.Printf("항목: %v\n", item)
//	})
//
//	// 디스플레이 문자열 생성
//	display := sliceutil.Join(tags, ", ")
//	fmt.Printf("태그: %s\n", display)
//
//	// 안전한 수정
//	modified := sliceutil.Clone(original)
//	// 'original'에 영향 없이 'modified' 수정
//
//	// 배열 초기화
//	zeros := sliceutil.Fill(template, 0)
//
//	// 시퀀스 구축
//	data := []int{1, 2, 5, 6}
//	complete := sliceutil.Insert(data, 2, 3, 4)
//
//	// 덱 셔플
//	shuffled := sliceutil.Shuffle(cards)
//
//	// 데이터 결합
//	keys := []string{"name", "age", "city"}
//	values := []any{"Alice", 30, "NYC"}
//	pairs := sliceutil.Zip(keys, values)
//
//	// 슬라이딩 분석
//	temps := []float64{20, 21, 22, 23, 24, 25}
//	windows := sliceutil.Window(temps, 3)
//	for _, w := range windows {
//	    avg := sliceutil.Average(w)
//	    fmt.Printf("윈도우 평균: %.1f\n", avg)
//	}
//
//	// 디버그 체인
//	result := sliceutil.Map(
//	    sliceutil.Tap(
//	        sliceutil.Filter(data, isValid),
//	        func(s []int) { log.Printf("필터 후: %v", s) },
//	    ),
//	    transform,
//	)
//
// 메서드 체인 패턴:
//
//	// 디버깅 체인을 위한 Tap 사용
//	result := sliceutil.Map(
//	    sliceutil.Tap(sliceutil.Filter(data, pred), log),
//	    mapper,
//	)
//
//	// 변경 전 복제
//	safe := sliceutil.Clone(original)
//	// 'safe' 자유롭게 변경
//
// 엣지 케이스:
//   - 빈 것에 ForEach: 작업 없음 (함수 절대 호출 안 됨)
//   - 빈 것에 Join: "" 반환 (빈 문자열)
//   - nil 복제: nil 반환 (nil 의미론 유지)
//   - 빈 것 Fill: []T{} 반환 (빈 슬라이스)
//   - 잘못된 인덱스로 Insert: 유효 범위로 제한
//   - 잘못된 인덱스로 Remove: 복제 반환
//   - <=1 Shuffle: 복제 반환 (셔플할 것 없음)
//   - 다른 길이로 Zip: 최소 길이 사용
//   - 잘못된 크기로 Window: [][]T{} 반환 (빈 것)
//
// 스레드 안전성:
//   - Shuffle: 스레드 안전 (잠긴 전역 RNG 사용)
//   - 다른 작업: 슬라이스가 동시에 수정되면 스레드 안전하지 않음
//   - Clone은 독립 복사본 생성 (원본 변경으로부터 안전)

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
// badZipped := [][2]any{{1, "one"}, {"wrong", 2}} // Wrong types
// 잘못된 타입
//
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
