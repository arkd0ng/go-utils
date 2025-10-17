package sliceutil

// index.go provides index-based navigation and manipulation operations for slices.
//
// This file implements operations for finding element positions, extracting elements
// by index, and removing elements at specified positions:
//
// Index Finding Operations:
//
// FindIndices (Predicate-Based Search):
//   - FindIndices(slice, predicate): Find all positions matching condition
//     Time: O(n), Space: O(k) where k = matching count
//     Returns slice of all matching indices
//     Returns empty slice if no matches (not nil)
//     Indices in ascending order (natural scan order)
//     No short-circuiting (always scans entire slice)
//     Example: FindIndices([1,2,3,4,5,6], isEven) → [1,3,5]
//     Use cases: Multi-match finding, pattern detection, batch operations
//
// Index-Based Extraction:
//
// AtIndices (Select by Position):
//   - AtIndices(slice, indices): Extract elements at specified positions
//     Time: O(m) where m = len(indices), Space: O(m)
//     Returns elements in order of indices provided
//     Out-of-bounds indices silently skipped (no panic)
//     Negative indices skipped (not supported)
//     Duplicate indices allowed (element appears multiple times in result)
//     Pre-allocates result with capacity len(indices)
//     Example: AtIndices([10,20,30,40,50], [0,2,4]) → [10,30,50]
//     Use cases: Selective extraction, reordering, sampling by position
//
// Index-Based Removal:
//
// RemoveIndices (Delete by Position):
//   - RemoveIndices(slice, indices): Remove elements at specified positions
//     Time: O(n+m) where m = len(indices), Space: O(n)
//     Returns new slice with specified positions removed
//     Original slice not modified (immutable operation)
//     Out-of-bounds indices silently skipped
//     Negative indices skipped
//     Duplicate indices treated as single removal
//     Uses hash set for O(1) removal checks
//     Example: RemoveIndices([10,20,30,40,50], [1,3]) → [10,30,50]
//     Use cases: Batch deletion, selective removal, index-based filtering
//
// Comparison with Value-Based Operations:
//
// FindIndices vs Find:
//   - FindIndices: Returns positions (indices), finds all matches
//   - Find: Returns values, finds first match only
//   - FindIndices: Useful when position matters
//   - Find: Useful when only value matters
//
// AtIndices vs Filter:
//   - AtIndices: Select by position (index-based)
//   - Filter: Select by condition (value-based)
//   - AtIndices: Known positions, can reorder
//   - Filter: Unknown positions, preserves order
//
// RemoveIndices vs Filter:
//   - RemoveIndices: Remove by position (index-based)
//   - Filter: Keep by condition (value-based, inverse logic)
//   - RemoveIndices: Known positions to delete
//   - Filter: Condition-based inclusion
//
// Design Principles:
//   - Safe indexing: Out-of-bounds indices handled gracefully
//   - No panics: Invalid indices skipped, not errors
//   - Immutability: Original slices never modified
//   - Flexible: Support duplicate indices, arbitrary order
//   - Type-safe: Generic implementations for any type
//
// Index Handling:
//   - Valid range: 0 to len(slice)-1
//   - Out of bounds: Silently skipped (no error, no panic)
//   - Negative: Not supported, silently skipped
//   - Duplicates: Allowed, processed multiple times or deduplicated
//
// Performance Characteristics:
//
// FindIndices:
//   - Time: O(n) - must scan entire slice
//   - Space: O(k) - stores k matching indices
//   - No early termination (finds all matches)
//   - Dynamic allocation as matches found
//
// AtIndices:
//   - Time: O(m) - iterate through m indices
//   - Space: O(m) - pre-allocate result capacity
//   - Each index lookup is O(1)
//   - Efficient for small m relative to n
//
// RemoveIndices:
//   - Time: O(n+m) - build removal set (O(m)) + scan slice (O(n))
//   - Space: O(n) - worst case result is nearly full slice
//   - Hash set for O(1) removal checks
//   - Single pass through slice
//
// Memory Allocation:
//   - FindIndices: Dynamic growth (starts empty)
//   - AtIndices: Pre-allocate len(indices) capacity
//   - RemoveIndices: Pre-allocate len(slice)-len(validIndices)
//
// Common Usage Patterns:
//
//	// Find all error positions
//	errorIndices := sliceutil.FindIndices(results, func(r Result) bool {
//	    return r.Error != nil
//	})
//	fmt.Printf("Errors at positions: %v\n", errorIndices)
//
//	// Extract specific elements by position
//	samples := sliceutil.AtIndices(data, []int{0, 10, 20, 30})
//
//	// Reorder elements
//	reordered := sliceutil.AtIndices(items, []int{2, 0, 3, 1})
//
//	// Duplicate elements by repeating indices
//	doubled := sliceutil.AtIndices(items, []int{0, 0, 1, 1, 2, 2})
//
//	// Remove errors by position
//	cleaned := sliceutil.RemoveIndices(data, errorIndices)
//
//	// Remove multiple specific positions
//	pruned := sliceutil.RemoveIndices(items, []int{1, 3, 5, 7})
//
// Combining Operations:
//
//	// Find then extract
//	matchingIndices := sliceutil.FindIndices(data, isValid)
//	validItems := sliceutil.AtIndices(data, matchingIndices)
//	// Equivalent to Filter, but index-based
//
//	// Find then remove
//	badIndices := sliceutil.FindIndices(data, isInvalid)
//	cleaned := sliceutil.RemoveIndices(data, badIndices)
//	// Equivalent to Filter with negated predicate
//
// Safe Indexing Examples:
//
//	// Out-of-bounds indices handled gracefully
//	nums := []int{1, 2, 3}
//	result := sliceutil.AtIndices(nums, []int{0, 5, 2})
//	// Returns [1, 3] - index 5 skipped, no panic
//
//	// Empty indices
//	result := sliceutil.AtIndices(nums, []int{})
//	// Returns [] - empty result
//
//	// All invalid indices
//	result := sliceutil.AtIndices(nums, []int{-1, 10, 20})
//	// Returns [] - all skipped
//
// Comparison with Standard Library:
//   - More convenient than manual index tracking
//   - Safer than direct slice indexing (no panics)
//   - More expressive than loop-based extraction
//   - Type-safe with generics
//
// Alternative Approaches:
//   - For single index: Direct slice indexing (slice[i])
//   - For contiguous ranges: Slice expressions (slice[start:end])
//   - For filtering: Use Filter for value-based selection
//   - For complex patterns: Combine with Map, Filter
//
// index.go는 슬라이스에 대한 인덱스 기반 탐색 및 조작 작업을 제공합니다.
//
// 이 파일은 요소 위치 찾기, 인덱스로 요소 추출, 지정된 위치의
// 요소 제거 작업을 구현합니다:
//
// 인덱스 찾기 작업:
//
// FindIndices (조건자 기반 검색):
//   - FindIndices(slice, predicate): 조건과 일치하는 모든 위치 찾기
//     시간: O(n), 공간: O(k) (k = 일치 개수)
//     일치하는 모든 인덱스의 슬라이스 반환
//     일치 없으면 빈 슬라이스 반환 (nil 아님)
//     인덱스를 오름차순 (자연 스캔 순서)
//     단락 없음 (항상 전체 슬라이스 스캔)
//     예: FindIndices([1,2,3,4,5,6], isEven) → [1,3,5]
//     사용 사례: 다중 일치 찾기, 패턴 감지, 배치 작업
//
// 인덱스 기반 추출:
//
// AtIndices (위치로 선택):
//   - AtIndices(slice, indices): 지정된 위치의 요소 추출
//     시간: O(m) (m = len(indices)), 공간: O(m)
//     제공된 인덱스 순서로 요소 반환
//     범위 벗어난 인덱스 자동 스킵 (패닉 없음)
//     음수 인덱스 스킵 (지원 안 함)
//     중복 인덱스 허용 (결과에 요소가 여러 번 나타남)
//     len(indices) 용량으로 결과 사전 할당
//     예: AtIndices([10,20,30,40,50], [0,2,4]) → [10,30,50]
//     사용 사례: 선택적 추출, 재정렬, 위치별 샘플링
//
// 인덱스 기반 제거:
//
// RemoveIndices (위치로 삭제):
//   - RemoveIndices(slice, indices): 지정된 위치의 요소 제거
//     시간: O(n+m) (m = len(indices)), 공간: O(n)
//     지정된 위치가 제거된 새 슬라이스 반환
//     원본 슬라이스 수정 안 함 (불변 작업)
//     범위 벗어난 인덱스 자동 스킵
//     음수 인덱스 스킵
//     중복 인덱스를 단일 제거로 처리
//     O(1) 제거 확인을 위해 해시 집합 사용
//     예: RemoveIndices([10,20,30,40,50], [1,3]) → [10,30,50]
//     사용 사례: 배치 삭제, 선택적 제거, 인덱스 기반 필터링
//
// 값 기반 작업과 비교:
//
// FindIndices vs Find:
//   - FindIndices: 위치(인덱스) 반환, 모든 일치 찾기
//   - Find: 값 반환, 첫 번째 일치만 찾기
//   - FindIndices: 위치가 중요할 때 유용
//   - Find: 값만 중요할 때 유용
//
// AtIndices vs Filter:
//   - AtIndices: 위치로 선택 (인덱스 기반)
//   - Filter: 조건으로 선택 (값 기반)
//   - AtIndices: 알려진 위치, 재정렬 가능
//   - Filter: 알 수 없는 위치, 순서 유지
//
// RemoveIndices vs Filter:
//   - RemoveIndices: 위치로 제거 (인덱스 기반)
//   - Filter: 조건으로 유지 (값 기반, 역논리)
//   - RemoveIndices: 삭제할 위치 알려짐
//   - Filter: 조건 기반 포함
//
// 설계 원칙:
//   - 안전한 인덱싱: 범위 벗어난 인덱스 우아하게 처리
//   - 패닉 없음: 잘못된 인덱스 스킵, 에러 없음
//   - 불변성: 원본 슬라이스 절대 수정 안 함
//   - 유연성: 중복 인덱스, 임의 순서 지원
//   - 타입 안전성: 모든 타입에 대한 제네릭 구현
//
// 인덱스 처리:
//   - 유효 범위: 0 ~ len(slice)-1
//   - 범위 벗어남: 자동 스킵 (에러 없음, 패닉 없음)
//   - 음수: 지원 안 함, 자동 스킵
//   - 중복: 허용, 여러 번 처리되거나 중복 제거
//
// 성능 특성:
//
// FindIndices:
//   - 시간: O(n) - 전체 슬라이스 스캔 필요
//   - 공간: O(k) - k개 일치 인덱스 저장
//   - 조기 종료 없음 (모든 일치 찾기)
//   - 일치 발견 시 동적 할당
//
// AtIndices:
//   - 시간: O(m) - m개 인덱스 반복
//   - 공간: O(m) - 결과 용량 사전 할당
//   - 각 인덱스 조회는 O(1)
//   - n에 비해 작은 m에 효율적
//
// RemoveIndices:
//   - 시간: O(n+m) - 제거 집합 구축 (O(m)) + 슬라이스 스캔 (O(n))
//   - 공간: O(n) - 최악의 경우 결과가 거의 전체 슬라이스
//   - O(1) 제거 확인을 위한 해시 집합
//   - 슬라이스를 한 번만 통과
//
// 메모리 할당:
//   - FindIndices: 동적 성장 (빈 것으로 시작)
//   - AtIndices: len(indices) 용량 사전 할당
//   - RemoveIndices: len(slice)-len(validIndices) 사전 할당
//
// 일반적인 사용 패턴:
//
//	// 모든 에러 위치 찾기
//	errorIndices := sliceutil.FindIndices(results, func(r Result) bool {
//	    return r.Error != nil
//	})
//	fmt.Printf("에러 위치: %v\n", errorIndices)
//
//	// 위치로 특정 요소 추출
//	samples := sliceutil.AtIndices(data, []int{0, 10, 20, 30})
//
//	// 요소 재정렬
//	reordered := sliceutil.AtIndices(items, []int{2, 0, 3, 1})
//
//	// 인덱스 반복으로 요소 복제
//	doubled := sliceutil.AtIndices(items, []int{0, 0, 1, 1, 2, 2})
//
//	// 위치로 에러 제거
//	cleaned := sliceutil.RemoveIndices(data, errorIndices)
//
//	// 여러 특정 위치 제거
//	pruned := sliceutil.RemoveIndices(items, []int{1, 3, 5, 7})
//
// 작업 결합:
//
//	// 찾기 후 추출
//	matchingIndices := sliceutil.FindIndices(data, isValid)
//	validItems := sliceutil.AtIndices(data, matchingIndices)
//	// Filter와 동등하지만 인덱스 기반
//
//	// 찾기 후 제거
//	badIndices := sliceutil.FindIndices(data, isInvalid)
//	cleaned := sliceutil.RemoveIndices(data, badIndices)
//	// 조건자를 부정한 Filter와 동등
//
// 안전한 인덱싱 예제:
//
//	// 범위 벗어난 인덱스 우아하게 처리
//	nums := []int{1, 2, 3}
//	result := sliceutil.AtIndices(nums, []int{0, 5, 2})
//	// [1, 3] 반환 - 인덱스 5 스킵, 패닉 없음
//
//	// 빈 인덱스
//	result := sliceutil.AtIndices(nums, []int{})
//	// [] 반환 - 빈 결과
//
//	// 모두 잘못된 인덱스
//	result := sliceutil.AtIndices(nums, []int{-1, 10, 20})
//	// [] 반환 - 모두 스킵
//
// 표준 라이브러리와 비교:
//   - 수동 인덱스 추적보다 편리
//   - 직접 슬라이스 인덱싱보다 안전 (패닉 없음)
//   - 루프 기반 추출보다 표현력 높음
//   - 제네릭으로 타입 안전
//
// 대안 접근법:
//   - 단일 인덱스: 직접 슬라이스 인덱싱 (slice[i])
//   - 연속 범위: 슬라이스 표현식 (slice[start:end])
//   - 필터링: 값 기반 선택을 위한 Filter 사용
//   - 복잡한 패턴: Map, Filter와 결합

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
