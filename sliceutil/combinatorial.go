package sliceutil

// combinatorial.go provides combinatorial mathematics operations for slices.
//
// This file implements operations for generating permutations and combinations,
// fundamental concepts in combinatorial mathematics:
//
// Permutation Operations:
//
// Permutations (All Orderings):
//   - Permutations(slice): Generate all possible orderings of elements
//     Time: O(n! * n), Space: O(n! * n)
//     Factorial growth: n=5 → 120, n=10 → 3,628,800
//     Uses Heap's algorithm for efficient generation
//     Each permutation is independent copy (safe to modify)
//     Order matters: [1,2,3] ≠ [3,2,1]
//     All elements used exactly once
//     Example: Permutations([1,2,3]) → [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
//     Mathematical: P(n) = n!
//     Use cases: Trying all orderings, brute-force search, scheduling
//
// Combination Operations:
//
// Combinations (k-Element Subsets):
//   - Combinations(slice, k): Generate all k-element subsets
//     Time: O(C(n,k) * k), Space: O(C(n,k) * k)
//     Binomial coefficient growth: C(n,k) = n! / (k! * (n-k)!)
//     C(10,5) = 252, C(20,10) = 184,756
//     Order doesn't matter: [1,2] == [2,1] (only one appears)
//     Each element used at most once
//     Example: Combinations([1,2,3,4], 2) → [[1,2], [1,3], [1,4], [2,3], [2,4], [3,4]]
//     Mathematical: C(n,k) = n! / (k! * (n-k)!)
//     Use cases: Selecting subsets, feature combinations, team selection
//
// Permutations vs Combinations:
//
// Permutations:
//   - Order matters: [A,B] ≠ [B,A]
//   - Uses all elements (n elements)
//   - Count: n!
//   - Example: Race finish order (Gold, Silver, Bronze)
//   - Larger result set
//
// Combinations:
//   - Order doesn't matter: [A,B] == [B,A]
//   - Uses k elements (k ≤ n)
//   - Count: C(n,k) = n! / (k! * (n-k)!)
//   - Example: Choosing k winners (order doesn't matter)
//   - Smaller result set for k < n
//
// When to Use Each:
//   - Permutations: Order matters (rankings, sequences, arrangements)
//   - Combinations: Order doesn't matter (selections, subsets, teams)
//
// Performance and Growth Rates:
//
// Permutations Growth (n!):
//   - n=1: 1
//   - n=3: 6
//   - n=5: 120
//   - n=7: 5,040
//   - n=10: 3,628,800
//   - n=12: 479,001,600 (nearly half billion!)
//   - WARNING: Extremely rapid growth, use with caution
//
// Combinations Growth (C(n,k)):
//   - C(5,2): 10
//   - C(10,3): 120
//   - C(10,5): 252
//   - C(20,10): 184,756
//   - C(30,15): 155,117,520
//   - Maximum at k ≈ n/2
//   - Smaller than permutations but still grows fast
//
// Memory Considerations:
//   - Each permutation: n elements
//   - Total permutations: n! * n elements
//   - Each combination: k elements
//   - Total combinations: C(n,k) * k elements
//   - Large n can exhaust memory quickly
//
// Practical Limits:
//   - Permutations: n ≤ 10 practical, n ≤ 12 max
//   - Combinations: Depends on k, but C(n,k) should be < millions
//   - Consider iterative generation for large sets (not implemented here)
//
// Design Principles:
//   - Independence: Each result is separate copy (safe to modify)
//   - Immutability: Original slice never modified
//   - Completeness: Generates all possibilities
//   - Type safety: Generic implementations for any type
//
// Algorithm Details:
//
// Permutations - Heap's Algorithm:
//   - In-place generation with backtracking
//   - Minimal swaps: O(n) swaps per permutation
//   - Efficient: One of the fastest permutation algorithms
//   - Lexicographic order not guaranteed
//
// Combinations - Recursive Generation:
//   - Depth-first search approach
//   - Builds combinations incrementally
//   - Lexicographic order maintained
//   - No duplicates generated
//
// Common Usage Patterns:
//
//	// Generate all orderings
//	tasks := []string{"A", "B", "C"}
//	allOrderings := sliceutil.Permutations(tasks)
//	for _, ordering := range allOrderings {
//	    fmt.Printf("Try order: %v\n", ordering)
//	}
//
//	// Select teams
//	players := []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}
//	teams := sliceutil.Combinations(players, 3)
//	fmt.Printf("%d possible teams of 3\n", len(teams))
//
//	// Find best arrangement
//	items := []Item{...}
//	bestScore := 0.0
//	var bestOrder []Item
//	for _, perm := range sliceutil.Permutations(items) {
//	    score := evaluate(perm)
//	    if score > bestScore {
//	        bestScore = score
//	        bestOrder = perm
//	    }
//	}
//
// Safety Checks:
//
//	// Check size before generating
//	if len(items) <= 10 {
//	    perms := sliceutil.Permutations(items)
//	    // Safe to process
//	} else {
//	    // Too large, use different approach
//	}
//
//	// Combinations with invalid k
//	combs := sliceutil.Combinations(items, -1)  // Returns [][]T{}
//	combs := sliceutil.Combinations(items, 100) // Returns [][]T{}
//
// Special Cases:
//   - Empty slice permutations: [[]] (one empty permutation)
//   - Combinations(slice, 0): [[]] (one empty combination)
//   - Combinations(slice, len(slice)): [[slice]] (one full combination)
//   - Combinations with k > len(slice): [] (no valid combinations)
//
// Comparison with Alternatives:
//   - Iterative generation: Memory efficient but more complex
//   - Lazy evaluation: Generate on demand (not implemented here)
//   - Recursive (current): Simple, complete generation
//
// combinatorial.go는 슬라이스에 대한 조합 수학 작업을 제공합니다.
//
// 이 파일은 순열과 조합 생성 작업을 구현하며, 조합 수학의 기본 개념입니다:
//
// 순열 작업:
//
// Permutations (모든 순서):
//   - Permutations(slice): 요소의 모든 가능한 순서 생성
//     시간: O(n! * n), 공간: O(n! * n)
//     팩토리얼 증가: n=5 → 120, n=10 → 3,628,800
//     효율적인 생성을 위해 Heap의 알고리즘 사용
//     각 순열은 독립적인 복사본 (수정 안전)
//     순서 중요: [1,2,3] ≠ [3,2,1]
//     모든 요소 정확히 한 번 사용
//     예: Permutations([1,2,3]) → [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
//     수학적: P(n) = n!
//     사용 사례: 모든 순서 시도, 무차별 대입 검색, 스케줄링
//
// 조합 작업:
//
// Combinations (k-요소 부분집합):
//   - Combinations(slice, k): 모든 k-요소 부분집합 생성
//     시간: O(C(n,k) * k), 공간: O(C(n,k) * k)
//     이항 계수 증가: C(n,k) = n! / (k! * (n-k)!)
//     C(10,5) = 252, C(20,10) = 184,756
//     순서 무관: [1,2] == [2,1] (하나만 나타남)
//     각 요소 최대 한 번 사용
//     예: Combinations([1,2,3,4], 2) → [[1,2], [1,3], [1,4], [2,3], [2,4], [3,4]]
//     수학적: C(n,k) = n! / (k! * (n-k)!)
//     사용 사례: 부분집합 선택, 기능 조합, 팀 선택
//
// 순열 vs 조합:
//
// 순열:
//   - 순서 중요: [A,B] ≠ [B,A]
//   - 모든 요소 사용 (n개 요소)
//   - 개수: n!
//   - 예: 경주 완주 순서 (금, 은, 동)
//   - 더 큰 결과 집합
//
// 조합:
//   - 순서 무관: [A,B] == [B,A]
//   - k개 요소 사용 (k ≤ n)
//   - 개수: C(n,k) = n! / (k! * (n-k)!)
//   - 예: k명 우승자 선택 (순서 무관)
//   - k < n일 때 더 작은 결과 집합
//
// 각각 사용 시기:
//   - 순열: 순서 중요 (순위, 시퀀스, 배열)
//   - 조합: 순서 무관 (선택, 부분집합, 팀)
//
// 성능 및 증가율:
//
// 순열 증가 (n!):
//   - n=1: 1
//   - n=3: 6
//   - n=5: 120
//   - n=7: 5,040
//   - n=10: 3,628,800
//   - n=12: 479,001,600 (거의 5억!)
//   - 경고: 매우 빠른 증가, 주의해서 사용
//
// 조합 증가 (C(n,k)):
//   - C(5,2): 10
//   - C(10,3): 120
//   - C(10,5): 252
//   - C(20,10): 184,756
//   - C(30,15): 155,117,520
//   - k ≈ n/2에서 최대
//   - 순열보다 작지만 여전히 빠르게 증가
//
// 메모리 고려사항:
//   - 각 순열: n개 요소
//   - 총 순열: n! * n 요소
//   - 각 조합: k개 요소
//   - 총 조합: C(n,k) * k 요소
//   - 큰 n은 메모리 빠르게 소진
//
// 실용적 한계:
//   - 순열: n ≤ 10 실용적, n ≤ 12 최대
//   - 조합: k에 따라 다름, C(n,k)는 수백만 미만이어야 함
//   - 큰 집합은 반복적 생성 고려 (여기 구현 안 됨)
//
// 설계 원칙:
//   - 독립성: 각 결과는 별도 복사본 (수정 안전)
//   - 불변성: 원본 슬라이스 절대 수정 안 함
//   - 완전성: 모든 가능성 생성
//   - 타입 안전성: 모든 타입에 대한 제네릭 구현
//
// 알고리즘 세부사항:
//
// 순열 - Heap의 알고리즘:
//   - 백트래킹으로 제자리 생성
//   - 최소 스왑: 순열당 O(n) 스왑
//   - 효율적: 가장 빠른 순열 알고리즘 중 하나
//   - 사전순 보장 안 함
//
// 조합 - 재귀 생성:
//   - 깊이 우선 탐색 접근
//   - 조합을 점진적으로 구축
//   - 사전순 유지
//   - 중복 생성 안 함
//
// 일반적인 사용 패턴:
//
//	// 모든 순서 생성
//	tasks := []string{"A", "B", "C"}
//	allOrderings := sliceutil.Permutations(tasks)
//	for _, ordering := range allOrderings {
//	    fmt.Printf("순서 시도: %v\n", ordering)
//	}
//
//	// 팀 선택
//	players := []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}
//	teams := sliceutil.Combinations(players, 3)
//	fmt.Printf("3명 팀 %d개 가능\n", len(teams))
//
//	// 최적 배열 찾기
//	items := []Item{...}
//	bestScore := 0.0
//	var bestOrder []Item
//	for _, perm := range sliceutil.Permutations(items) {
//	    score := evaluate(perm)
//	    if score > bestScore {
//	        bestScore = score
//	        bestOrder = perm
//	    }
//	}
//
// 안전 확인:
//
//	// 생성 전 크기 확인
//	if len(items) <= 10 {
//	    perms := sliceutil.Permutations(items)
//	    // 처리 안전
//	} else {
//	    // 너무 큼, 다른 접근법 사용
//	}
//
//	// 잘못된 k로 조합
//	combs := sliceutil.Combinations(items, -1)  // [][]T{} 반환
//	combs := sliceutil.Combinations(items, 100) // [][]T{} 반환
//
// 특수 경우:
//   - 빈 슬라이스 순열: [[]] (하나의 빈 순열)
//   - Combinations(slice, 0): [[]] (하나의 빈 조합)
//   - Combinations(slice, len(slice)): [[slice]] (하나의 전체 조합)
//   - k > len(slice)인 조합: [] (유효한 조합 없음)
//
// 대안과 비교:
//   - 반복적 생성: 메모리 효율적이지만 더 복잡
//   - 지연 평가: 요청 시 생성 (여기 구현 안 됨)
//   - 재귀 (현재): 단순, 완전 생성

// Permutations returns all possible permutations of the slice.
// Returns a slice of slices, where each sub-slice is a permutation.
// Warning: The number of permutations grows factorially (n!).
// Permutations는 슬라이스의 모든 가능한 순열을 반환합니다.
// 각 하위 슬라이스가 순열인 슬라이스의 슬라이스를 반환합니다.
// 경고: 순열의 수는 팩토리얼로 증가합니다 (n!).
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3}
//	perms := sliceutil.Permutations(numbers)
//	// perms: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
//
//	letters := []string{"a", "b"}
//	perms := sliceutil.Permutations(letters)
//	// perms: [["a","b"], ["b","a"]]
//
// Performance note
// 성능 참고:
//
//	n=5: 120 permutations
//	n=10: 3,628,800 permutations
//	Use with caution for large slices!
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
// Warning: The number of combinations is C(n, k) = n!
// (k! * (n-k)!).
// Combinations는 슬라이스에서 k개 요소의 모든 가능한 조합을 반환합니다.
// 각 하위 슬라이스가 조합인 슬라이스의 슬라이스를 반환합니다.
// 경고: 조합의 수는 C(n, k) = n!
// (k! * (n-k)!)입니다.
//
// Example
// 예제:
//
//	numbers := []int{1, 2, 3, 4}
//	combs := sliceutil.Combinations(numbers, 2)
//	// combs: [[1,2], [1,3], [1,4], [2,3], [2,4], [3,4]]
//
//	letters := []string{"a", "b", "c"}
//	combs := sliceutil.Combinations(letters, 2)
//	// combs: [["a","b"], ["a","c"], ["b","c"]]
//
// Performance note
// 성능 참고:
//
//	C(10, 5) = 252 combinations
//	C(20, 10) = 184,756 combinations
//	Use with caution for large values!
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
