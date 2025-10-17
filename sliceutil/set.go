package sliceutil

// set.go provides mathematical set operations for slices.
//
// This file implements set theory operations treating slices as sets,
// with automatic deduplication and membership testing:
//
// Set Combination Operations:
//
// Union (A ∪ B):
//   - Union(a, b): All unique elements from both slices
//     Time: O(n+m), Space: O(n+m)
//     Combines both slices, removing duplicates
//     Preserves order: Elements from 'a' first, then unique from 'b'
//     Uses hash map for O(1) duplicate detection
//     Example: Union([1,2,3], [3,4,5]) = [1,2,3,4,5]
//     Mathematical: {x | x ∈ A ∨ x ∈ B}
//     Use cases: Combining lists, merging tags, aggregating results
//
// Intersection (A ∩ B):
//   - Intersection(a, b): Elements present in both slices
//     Time: O(n+m), Space: O(min(n,m))
//     Only elements appearing in both slices
//     Preserves order from first slice ('a')
//     Builds set from 'b' for O(1) lookups
//     Result size: 0 to min(len(a), len(b))
//     Example: Intersection([1,2,3,4], [3,4,5,6]) = [3,4]
//     Mathematical: {x | x ∈ A ∧ x ∈ B}
//     Use cases: Finding common elements, overlap detection, shared features
//
// Difference (A \ B or A - B):
//   - Difference(a, b): Elements in 'a' but not in 'b'
//     Time: O(n+m), Space: O(n)
//     Relative complement of B in A
//     Only elements unique to first slice
//     Preserves order from first slice ('a')
//     Builds set from 'b' for O(1) exclusion checks
//     Example: Difference([1,2,3,4], [3,4,5,6]) = [1,2]
//     Mathematical: {x | x ∈ A ∧ x ∉ B}
//     Use cases: Finding missing elements, exclusion filtering, delta computation
//
// SymmetricDifference (A △ B):
//   - SymmetricDifference(a, b): Elements in either but not both
//     Time: O(n+m), Space: O(n+m)
//     Elements unique to each slice (exclusive or)
//     Equivalent to: (A \ B) ∪ (B \ A)
//     Equivalent to: (A ∪ B) \ (A ∩ B)
//     Order: Unique from 'a' first, then unique from 'b'
//     Example: SymmetricDifference([1,2,3,4], [3,4,5,6]) = [1,2,5,6]
//     Mathematical: {x | (x ∈ A ∧ x ∉ B) ∨ (x ∈ B ∧ x ∉ A)}
//     Use cases: Finding differences, change detection, XOR-like operations
//
// Set Relationship Testing:
//
// IsSubset (A ⊆ B):
//   - IsSubset(a, b): Check if all elements of 'a' are in 'b'
//     Time: O(n+m), Space: O(m)
//     Returns true if A is subset of B
//     Empty set is subset of any set
//     Set is subset of itself
//     Short-circuits on first missing element
//     Example: IsSubset([1,2], [1,2,3,4]) = true
//     Mathematical: ∀x ∈ A, x ∈ B
//     Use cases: Permission checking, requirement validation, containment testing
//
// IsSuperset (A ⊇ B):
//   - IsSuperset(a, b): Check if all elements of 'b' are in 'a'
//     Time: O(n+m), Space: O(n)
//     Returns true if A is superset of B
//     Implemented as: IsSuperset(a, b) = IsSubset(b, a)
//     Set is superset of empty set
//     Set is superset of itself
//     Example: IsSuperset([1,2,3,4], [1,2]) = true
//     Mathematical: ∀x ∈ B, x ∈ A
//     Use cases: Capability checking, coverage validation, containment testing
//
// Set Theory Relationships:
//
// Properties:
//   - Union: Commutative (A ∪ B = B ∪ A), Associative ((A ∪ B) ∪ C = A ∪ (B ∪ C))
//   - Intersection: Commutative (A ∩ B = B ∩ A), Associative ((A ∩ B) ∩ C = A ∩ (B ∩ C))
//   - Difference: NOT commutative (A \ B ≠ B \ A)
//   - SymmetricDifference: Commutative (A △ B = B △ A), Associative
//
// Identities:
//   - Union(A, ∅) = A (identity element: empty set)
//   - Intersection(A, ∅) = ∅ (empty set dominates)
//   - Difference(A, ∅) = A
//   - Difference(A, A) = ∅
//   - SymmetricDifference(A, ∅) = A
//   - SymmetricDifference(A, A) = ∅
//
// De Morgan's Laws:
//   - Not implemented directly but mathematically:
//     (A ∪ B)' = A' ∩ B'
//     (A ∩ B)' = A' ∪ B'
//
// Design Principles:
//   - Deduplication: All operations automatically remove duplicates
//   - Order preservation: Results maintain order from input slices where possible
//   - Hash-based: Use maps for O(1) membership testing
//   - Type safety: Generic comparable constraint ensures correctness
//   - Immutability: Original slices never modified
//
// Performance Characteristics:
//
// Time Complexity:
//   - Union: O(n+m) - process both slices once
//   - Intersection: O(n+m) - build set + scan first slice
//   - Difference: O(n+m) - build set + scan first slice
//   - SymmetricDifference: O(n+m) - build both sets + scan both
//   - IsSubset: O(n+m) best case O(m) with short-circuit
//   - IsSuperset: O(n+m) best case O(n) with short-circuit
//
// Space Complexity:
//   - Union: O(n+m) - map + result can be up to len(a)+len(b)
//   - Intersection: O(m) - map for b + result up to min(n,m)
//   - Difference: O(m) - map for b + result up to n
//   - SymmetricDifference: O(n+m) - maps for both + result
//   - IsSubset: O(m) - map for b only
//   - IsSuperset: O(n) - map for a only
//
// Memory Efficiency Considerations:
//   - All operations create hash maps for membership testing
//   - Result slices allocated dynamically (except Union pre-allocates)
//   - Large slices require significant memory for hash maps
//   - Consider sorting + merging for very large datasets
//
// Common Usage Patterns:
//
//	// Merge tag lists
//	allTags := sliceutil.Union(article1.Tags, article2.Tags)
//
//	// Find common interests
//	common := sliceutil.Intersection(user1.Interests, user2.Interests)
//
//	// Find missing permissions
//	missing := sliceutil.Difference(required, actual)
//
//	// Find changes between versions
//	changes := sliceutil.SymmetricDifference(oldItems, newItems)
//
//	// Verify permissions
//	if sliceutil.IsSubset(requiredPermissions, userPermissions) {
//	    fmt.Println("User has all required permissions")
//	}
//
//	// Check capabilities
//	if sliceutil.IsSuperset(systemCapabilities, requestedFeatures) {
//	    fmt.Println("System supports all requested features")
//	}
//
// Combining Set Operations:
//
//	// Multiple unions (associative)
//	allElements := sliceutil.Union(
//	    sliceutil.Union(set1, set2),
//	    set3,
//	)
//
//	// Find elements unique to first of three sets
//	uniqueToFirst := sliceutil.Difference(
//	    set1,
//	    sliceutil.Union(set2, set3),
//	)
//
// Comparison with Standard Library:
//   - More intuitive than manual map manipulation
//   - Clearer intent than nested loops
//   - Type-safe with generics
//   - Functional composition style
//
// Alternative Implementations:
//   - For sorted slices: Can use merge-based algorithms (potentially faster)
//   - For small slices: Linear search might be faster (no map overhead)
//   - For repeated operations: Pre-convert to actual set data structure
//
// set.go는 슬라이스에 대한 수학적 집합 연산을 제공합니다.
//
// 이 파일은 슬라이스를 집합으로 취급하여 집합 이론 연산을 구현하며,
// 자동 중복 제거와 멤버십 테스트를 포함합니다:
//
// 집합 결합 작업:
//
// Union (A ∪ B):
//   - Union(a, b): 두 슬라이스의 모든 고유 요소
//     시간: O(n+m), 공간: O(n+m)
//     두 슬라이스를 결합하고 중복 제거
//     순서 유지: 'a'의 요소 먼저, 그 다음 'b'의 고유 요소
//     O(1) 중복 감지를 위해 해시 맵 사용
//     예: Union([1,2,3], [3,4,5]) = [1,2,3,4,5]
//     수학적: {x | x ∈ A ∨ x ∈ B}
//     사용 사례: 목록 결합, 태그 병합, 결과 집계
//
// Intersection (A ∩ B):
//   - Intersection(a, b): 두 슬라이스에 모두 존재하는 요소
//     시간: O(n+m), 공간: O(min(n,m))
//     두 슬라이스 모두에 나타나는 요소만
//     첫 번째 슬라이스('a')의 순서 유지
//     O(1) 조회를 위해 'b'에서 집합 구축
//     결과 크기: 0 ~ min(len(a), len(b))
//     예: Intersection([1,2,3,4], [3,4,5,6]) = [3,4]
//     수학적: {x | x ∈ A ∧ x ∈ B}
//     사용 사례: 공통 요소 찾기, 겹침 감지, 공유 기능
//
// Difference (A \ B 또는 A - B):
//   - Difference(a, b): 'a'에 있지만 'b'에 없는 요소
//     시간: O(n+m), 공간: O(n)
//     A에서 B의 상대 여집합
//     첫 번째 슬라이스에 고유한 요소만
//     첫 번째 슬라이스('a')의 순서 유지
//     O(1) 배제 확인을 위해 'b'에서 집합 구축
//     예: Difference([1,2,3,4], [3,4,5,6]) = [1,2]
//     수학적: {x | x ∈ A ∧ x ∉ B}
//     사용 사례: 누락 요소 찾기, 배제 필터링, 델타 계산
//
// SymmetricDifference (A △ B):
//   - SymmetricDifference(a, b): 둘 중 하나에만 있는 요소
//     시간: O(n+m), 공간: O(n+m)
//     각 슬라이스에 고유한 요소 (배타적 논리합)
//     동등: (A \ B) ∪ (B \ A)
//     동등: (A ∪ B) \ (A ∩ B)
//     순서: 'a'의 고유 요소 먼저, 그 다음 'b'의 고유 요소
//     예: SymmetricDifference([1,2,3,4], [3,4,5,6]) = [1,2,5,6]
//     수학적: {x | (x ∈ A ∧ x ∉ B) ∨ (x ∈ B ∧ x ∉ A)}
//     사용 사례: 차이점 찾기, 변경 감지, XOR 유사 작업
//
// 집합 관계 테스트:
//
// IsSubset (A ⊆ B):
//   - IsSubset(a, b): 'a'의 모든 요소가 'b'에 있는지 확인
//     시간: O(n+m), 공간: O(m)
//     A가 B의 부분집합이면 true 반환
//     공집합은 모든 집합의 부분집합
//     집합은 자기 자신의 부분집합
//     첫 번째 누락 요소에서 단락
//     예: IsSubset([1,2], [1,2,3,4]) = true
//     수학적: ∀x ∈ A, x ∈ B
//     사용 사례: 권한 확인, 요구사항 검증, 포함 테스트
//
// IsSuperset (A ⊇ B):
//   - IsSuperset(a, b): 'b'의 모든 요소가 'a'에 있는지 확인
//     시간: O(n+m), 공간: O(n)
//     A가 B의 상위집합이면 true 반환
//     구현: IsSuperset(a, b) = IsSubset(b, a)
//     집합은 공집합의 상위집합
//     집합은 자기 자신의 상위집합
//     예: IsSuperset([1,2,3,4], [1,2]) = true
//     수학적: ∀x ∈ B, x ∈ A
//     사용 사례: 기능 확인, 커버리지 검증, 포함 테스트
//
// 집합 이론 관계:
//
// 속성:
//   - Union: 교환법칙 (A ∪ B = B ∪ A), 결합법칙 ((A ∪ B) ∪ C = A ∪ (B ∪ C))
//   - Intersection: 교환법칙 (A ∩ B = B ∩ A), 결합법칙 ((A ∩ B) ∩ C = A ∩ (B ∩ C))
//   - Difference: 교환법칙 없음 (A \ B ≠ B \ A)
//   - SymmetricDifference: 교환법칙 (A △ B = B △ A), 결합법칙
//
// 항등원:
//   - Union(A, ∅) = A (항등원: 공집합)
//   - Intersection(A, ∅) = ∅ (공집합이 지배)
//   - Difference(A, ∅) = A
//   - Difference(A, A) = ∅
//   - SymmetricDifference(A, ∅) = A
//   - SymmetricDifference(A, A) = ∅
//
// 드 모르간 법칙:
//   - 직접 구현은 안 했지만 수학적으로:
//     (A ∪ B)' = A' ∩ B'
//     (A ∩ B)' = A' ∪ B'
//
// 설계 원칙:
//   - 중복 제거: 모든 작업이 자동으로 중복 제거
//   - 순서 보존: 결과가 가능한 경우 입력 슬라이스의 순서 유지
//   - 해시 기반: O(1) 멤버십 테스트를 위해 맵 사용
//   - 타입 안전성: 제네릭 comparable 제약 조건으로 정확성 보장
//   - 불변성: 원본 슬라이스 절대 수정 안 함
//
// 성능 특성:
//
// 시간 복잡도:
//   - Union: O(n+m) - 두 슬라이스를 한 번 처리
//   - Intersection: O(n+m) - 집합 구축 + 첫 번째 슬라이스 스캔
//   - Difference: O(n+m) - 집합 구축 + 첫 번째 슬라이스 스캔
//   - SymmetricDifference: O(n+m) - 두 집합 구축 + 두 슬라이스 스캔
//   - IsSubset: O(n+m) 최선 O(m) 단락 있음
//   - IsSuperset: O(n+m) 최선 O(n) 단락 있음
//
// 공간 복잡도:
//   - Union: O(n+m) - 맵 + 결과는 len(a)+len(b)까지 가능
//   - Intersection: O(m) - b용 맵 + 결과는 min(n,m)까지
//   - Difference: O(m) - b용 맵 + 결과는 n까지
//   - SymmetricDifference: O(n+m) - 두 맵 + 결과
//   - IsSubset: O(m) - b용 맵만
//   - IsSuperset: O(n) - a용 맵만
//
// 메모리 효율성 고려사항:
//   - 모든 작업이 멤버십 테스트를 위한 해시 맵 생성
//   - 결과 슬라이스 동적 할당 (Union은 사전 할당 제외)
//   - 큰 슬라이스는 해시 맵에 상당한 메모리 필요
//   - 매우 큰 데이터셋은 정렬 + 병합 고려
//
// 일반적인 사용 패턴:
//
//	// 태그 목록 병합
//	allTags := sliceutil.Union(article1.Tags, article2.Tags)
//
//	// 공통 관심사 찾기
//	common := sliceutil.Intersection(user1.Interests, user2.Interests)
//
//	// 누락된 권한 찾기
//	missing := sliceutil.Difference(required, actual)
//
//	// 버전 간 변경사항 찾기
//	changes := sliceutil.SymmetricDifference(oldItems, newItems)
//
//	// 권한 검증
//	if sliceutil.IsSubset(requiredPermissions, userPermissions) {
//	    fmt.Println("사용자가 필요한 모든 권한 보유")
//	}
//
//	// 기능 확인
//	if sliceutil.IsSuperset(systemCapabilities, requestedFeatures) {
//	    fmt.Println("시스템이 요청된 모든 기능 지원")
//	}
//
// 집합 연산 결합:
//
//	// 여러 합집합 (결합법칙)
//	allElements := sliceutil.Union(
//	    sliceutil.Union(set1, set2),
//	    set3,
//	)
//
//	// 세 집합 중 첫 번째에 고유한 요소 찾기
//	uniqueToFirst := sliceutil.Difference(
//	    set1,
//	    sliceutil.Union(set2, set3),
//	)
//
// 표준 라이브러리와 비교:
//   - 수동 맵 조작보다 직관적
//   - 중첩 루프보다 의도 명확
//   - 제네릭으로 타입 안전
//   - 함수형 구성 스타일
//
// 대안 구현:
//   - 정렬된 슬라이스: 병합 기반 알고리즘 사용 가능 (잠재적으로 더 빠름)
//   - 작은 슬라이스: 선형 검색이 더 빠를 수 있음 (맵 오버헤드 없음)
//   - 반복 작업: 실제 집합 데이터 구조로 미리 변환

// Union returns the union of two slices (all unique elements from both).
// Union은 두 슬라이스의 합집합을 반환합니다 (양쪽의 모든 고유 요소).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3}
//	b := []int{3, 4, 5}
//	union := sliceutil.Union(a, b) // [1, 2, 3, 4, 5]
//
//	words1 := []string{"apple", "banana"}
//	words2 := []string{"banana", "cherry"}
//	union := sliceutil.Union(words1, words2) // ["apple", "banana", "cherry"]
func Union[T comparable](a, b []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(a)+len(b))

	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	for _, v := range b {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Intersection returns the intersection of two slices (common elements).
// Intersection은 두 슬라이스의 교집합을 반환합니다 (공통 요소).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{3, 4, 5, 6}
//	common := sliceutil.Intersection(a, b) // [3, 4]
//
//	words1 := []string{"apple", "banana", "cherry"}
//	words2 := []string{"banana", "cherry", "date"}
//	common := sliceutil.Intersection(words1, words2) // ["banana", "cherry"]
func Intersection[T comparable](a, b []T) []T {
	setB := make(map[T]bool)
	for _, v := range b {
		setB[v] = true
	}

	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, v := range a {
		if setB[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Difference returns the difference of two slices (elements in a but not in b).
// Difference는 두 슬라이스의 차집합을 반환합니다 (a에 있지만 b에 없는 요소).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{3, 4, 5, 6}
//	diff := sliceutil.Difference(a, b) // [1, 2]
//
//	words1 := []string{"apple", "banana", "cherry"}
//	words2 := []string{"banana", "date"}
//	diff := sliceutil.Difference(words1, words2) // ["apple", "cherry"]
func Difference[T comparable](a, b []T) []T {
	setB := make(map[T]bool)
	for _, v := range b {
		setB[v] = true
	}

	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, v := range a {
		if !setB[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// SymmetricDifference returns the symmetric difference of two slices.
// Elements that are in either a or b, but not in both.
// SymmetricDifference는 두 슬라이스의 대칭 차집합을 반환합니다.
// a 또는 b에 있지만 둘 다에 있지 않은 요소.
//
// Example
// 예제:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{3, 4, 5, 6}
//	symDiff := sliceutil.SymmetricDifference(a, b) // [1, 2, 5, 6]
//
//	words1 := []string{"apple", "banana"}
//	words2 := []string{"banana", "cherry"}
//	symDiff := sliceutil.SymmetricDifference(words1, words2) // ["apple", "cherry"]
func SymmetricDifference[T comparable](a, b []T) []T {
	setA := make(map[T]bool)
	setB := make(map[T]bool)

	for _, v := range a {
		setA[v] = true
	}
	for _, v := range b {
		setB[v] = true
	}

	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, v := range a {
		if !setB[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	for _, v := range b {
		if !setA[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// IsSubset returns true if a is a subset of b (all elements of a are in b).
// IsSubset은 a가 b의 부분집합이면 true를 반환합니다 (a의 모든 요소가 b에 있음).
//
// Example
// 예제:
//
//	a := []int{1, 2}
//	b := []int{1, 2, 3, 4}
//	isSubset := sliceutil.IsSubset(a, b) // true
//
//	a := []int{1, 5}
//	b := []int{1, 2, 3, 4}
//	isSubset := sliceutil.IsSubset(a, b) // false
func IsSubset[T comparable](a, b []T) bool {
	setB := make(map[T]bool)
	for _, v := range b {
		setB[v] = true
	}

	for _, v := range a {
		if !setB[v] {
			return false
		}
	}

	return true
}

// IsSuperset returns true if a is a superset of b (all elements of b are in a).
// IsSuperset은 a가 b의 상위집합이면 true를 반환합니다 (b의 모든 요소가 a에 있음).
//
// Example
// 예제:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{1, 2}
//	isSuperset := sliceutil.IsSuperset(a, b) // true
//
//	a := []int{1, 2, 3}
//	b := []int{1, 5}
//	isSuperset := sliceutil.IsSuperset(a, b) // false
func IsSuperset[T comparable](a, b []T) bool {
	return IsSubset(b, a)
}
