package sliceutil

// diff.go provides difference comparison and equality testing operations for slices.
//
// This file implements operations for comparing slices to detect changes,
// verify equality regardless of order, and check for duplicate elements:
//
// Difference Detection:
//
// Diff (Value-Based Comparison):
//   - Diff(old, new): Compare two slices by element values
//     Time: O(n+m), Space: O(n+m)
//     Returns DiffResult with Added, Removed, Unchanged categories
//     Added: Elements in new but not in old (new additions)
//     Removed: Elements in old but not in new (deletions)
//     Unchanged: Elements present in both (preserved)
//     Uses hash sets for O(1) membership testing
//     Automatically deduplicates unchanged elements
//     Example: Diff([1,2,3,4], [2,3,4,5]) → Added:[5], Removed:[1], Unchanged:[2,3,4]
//     Use cases: Version comparison, change tracking, sync detection
//
// DiffBy (Key-Based Comparison):
//   - DiffBy(old, new, keyFunc): Compare slices using extracted keys
//     Time: O(n+m), Space: O(n+m)
//     More flexible than Diff for complex types
//     Key function extracts comparable identifier from each element
//     Compares elements by key equality, not value equality
//     Returns full elements (not just keys) in DiffResult
//     Example: DiffBy(users, newUsers, u → u.ID) compares by ID
//     Use cases: Comparing structs, database sync, entity tracking
//
// DiffResult Structure:
//   - Added []T: New elements (in new but not old)
//   - Removed []T: Deleted elements (in old but not new)
//   - Unchanged []T: Preserved elements (in both)
//   - Can compute: Total changes = len(Added) + len(Removed)
//   - Can compute: Stability ratio = len(Unchanged) / (len(old) + len(new))
//
// Equality Testing:
//
// EqualUnordered (Set Equality):
//   - EqualUnordered(a, b): Check equality ignoring order
//     Time: O(n), Space: O(n)
//     Compares element frequencies (multiset equality)
//     Order doesn't matter: [1,2,3] == [3,2,1]
//     Frequencies matter: [1,2,2] != [1,2,3]
//     Different lengths always return false (early exit)
//     Uses Frequencies internally (element → count map)
//     Example: EqualUnordered([1,2,3,2], [3,2,1,2]) = true
//     Use cases: Test assertions, set comparison, unordered validation
//
// Duplicate Detection:
//
// HasDuplicates (Uniqueness Check):
//   - HasDuplicates(slice): Check if any element appears multiple times
//     Time: O(n) worst case, O(k) average (k = position of first dup), Space: O(n)
//     Short-circuits on first duplicate found
//     Uses hash set for O(1) lookups
//     Returns true immediately when duplicate detected
//     Empty slice returns false (no duplicates)
//     Example: HasDuplicates([1,2,3,2,4]) = true
//     Use cases: Uniqueness validation, constraint checking, data quality
//
// Comparison: Diff vs DiffBy:
//
// Diff:
//   - Direct value comparison
//   - Requires comparable type
//   - Simple element matching
//   - Example: Diff([1,2,3], [2,3,4])
//
// DiffBy:
//   - Key-based comparison
//   - Element type can be any, key must be comparable
//   - Flexible matching criteria
//   - Useful for structs: DiffBy(users, newUsers, u → u.ID)
//   - Can ignore certain fields: Compare by ID, ignore Name changes
//
// When to Use Each:
//   - Diff: For simple types (int, string) or when entire value matters
//   - DiffBy: For structs/complex types with identifier field
//   - DiffBy: When comparing by specific property (ID, name, etc.)
//   - DiffBy: When element type is not comparable but has comparable key
//
// Design Principles:
//   - Structured results: DiffResult provides organized output
//   - Hash-based: O(1) lookups for efficiency
//   - Deduplication: Automatically handles duplicates in input
//   - Type safety: Generic constraints ensure correctness
//   - Immutability: Original slices never modified
//
// Performance Characteristics:
//
// Time Complexity:
//   - Diff: O(n+m) - build two sets + iterate both slices
//   - DiffBy: O(n+m) - build two maps + iterate both slices
//   - EqualUnordered: O(n) - build frequency maps + compare (assumes len(a)==len(b))
//   - HasDuplicates: O(n) worst case, O(k) average with short-circuit
//
// Space Complexity:
//   - Diff: O(n+m) - two hash sets + result slices
//   - DiffBy: O(n+m) - two hash maps + result slices
//   - EqualUnordered: O(n) - two frequency maps
//   - HasDuplicates: O(k) where k = unique elements seen before first duplicate
//
// Memory Allocation:
//   - Diff/DiffBy: Dynamic allocation for result slices (size unknown)
//   - EqualUnordered: Pre-check length mismatch (early exit)
//   - HasDuplicates: Minimal allocation if duplicate found early
//
// Common Usage Patterns:
//
//	// Track changes between versions
//	diff := sliceutil.Diff(oldTags, newTags)
//	fmt.Printf("Added: %v, Removed: %v\n", diff.Added, diff.Removed)
//
//	// Compare user lists by ID
//	type User struct { ID int; Name string; Email string }
//	diff := sliceutil.DiffBy(oldUsers, newUsers, func(u User) int {
//	    return u.ID
//	})
//	for _, added := range diff.Added {
//	    fmt.Printf("New user: %s\n", added.Name)
//	}
//
//	// Test equality ignoring order
//	expected := []int{1, 2, 3}
//	actual := []int{3, 1, 2}
//	if sliceutil.EqualUnordered(expected, actual) {
//	    fmt.Println("Test passed")
//	}
//
//	// Validate uniqueness
//	if sliceutil.HasDuplicates(usernames) {
//	    return errors.New("duplicate usernames not allowed")
//	}
//
// Change Detection Patterns:
//
//	// Sync databases
//	diff := sliceutil.DiffBy(dbRecords, apiRecords, r → r.ID)
//	// Insert diff.Added to database
//	// Delete diff.Removed from database
//	// Update diff.Unchanged if needed
//
//	// Track inventory changes
//	diff := sliceutil.Diff(previousStock, currentStock)
//	restockNeeded := diff.Removed
//	newItems := diff.Added
//
// Testing Patterns:
//
//	// Flexible test assertions
//	func TestUnorderedResults(t *testing.T) {
//	    result := GetItems()
//	    expected := []string{"a", "b", "c"}
//	    if !sliceutil.EqualUnordered(result, expected) {
//	        t.Errorf("got %v, want %v (any order)", result, expected)
//	    }
//	}
//
// Comparison with Standard Library:
//   - More structured than manual loop comparisons
//   - Clearer intent than nested loops with maps
//   - Type-safe with generics
//   - Functional composition style
//
// Alternative Approaches:
//   - For sorted slices: Can use two-pointer merge algorithm (more efficient)
//   - For small slices: Linear search might be simpler
//   - For streaming data: Consider incremental diff algorithms
//
// diff.go는 슬라이스에 대한 차이 비교 및 동등성 테스트 작업을 제공합니다.
//
// 이 파일은 슬라이스를 비교하여 변경 사항을 감지하고, 순서와 관계없이
// 동등성을 검증하며, 중복 요소를 확인하는 작업을 구현합니다:
//
// 차이 감지:
//
// Diff (값 기반 비교):
//   - Diff(old, new): 요소 값으로 두 슬라이스 비교
//     시간: O(n+m), 공간: O(n+m)
//     Added, Removed, Unchanged 카테고리로 DiffResult 반환
//     Added: new에만 있는 요소 (새 추가)
//     Removed: old에만 있는 요소 (삭제)
//     Unchanged: 양쪽에 있는 요소 (보존)
//     O(1) 멤버십 테스트를 위해 해시 집합 사용
//     변경되지 않은 요소 자동 중복 제거
//     예: Diff([1,2,3,4], [2,3,4,5]) → Added:[5], Removed:[1], Unchanged:[2,3,4]
//     사용 사례: 버전 비교, 변경 추적, 동기화 감지
//
// DiffBy (키 기반 비교):
//   - DiffBy(old, new, keyFunc): 추출된 키로 슬라이스 비교
//     시간: O(n+m), 공간: O(n+m)
//     복잡한 타입에 대해 Diff보다 유연
//     키 함수가 각 요소에서 comparable 식별자 추출
//     값 동등성이 아닌 키 동등성으로 요소 비교
//     DiffResult에 전체 요소 반환 (키만이 아님)
//     예: DiffBy(users, newUsers, u → u.ID) ID로 비교
//     사용 사례: 구조체 비교, 데이터베이스 동기화, 엔티티 추적
//
// DiffResult 구조:
//   - Added []T: 새 요소 (new에만 있음)
//   - Removed []T: 삭제된 요소 (old에만 있음)
//   - Unchanged []T: 보존된 요소 (양쪽 모두에 있음)
//   - 계산 가능: 총 변경 = len(Added) + len(Removed)
//   - 계산 가능: 안정성 비율 = len(Unchanged) / (len(old) + len(new))
//
// 동등성 테스트:
//
// EqualUnordered (집합 동등성):
//   - EqualUnordered(a, b): 순서 무시하고 동등성 확인
//     시간: O(n), 공간: O(n)
//     요소 빈도 비교 (다중집합 동등성)
//     순서 무관: [1,2,3] == [3,2,1]
//     빈도 중요: [1,2,2] != [1,2,3]
//     길이 다르면 항상 false (조기 종료)
//     내부적으로 Frequencies 사용 (요소 → 개수 맵)
//     예: EqualUnordered([1,2,3,2], [3,2,1,2]) = true
//     사용 사례: 테스트 단언, 집합 비교, 순서 없는 검증
//
// 중복 감지:
//
// HasDuplicates (고유성 확인):
//   - HasDuplicates(slice): 요소가 여러 번 나타나는지 확인
//     시간: O(n) 최악, O(k) 평균 (k = 첫 중복 위치), 공간: O(n)
//     첫 번째 중복 발견 시 단락
//     O(1) 조회를 위해 해시 집합 사용
//     중복 감지 즉시 true 반환
//     빈 슬라이스는 false 반환 (중복 없음)
//     예: HasDuplicates([1,2,3,2,4]) = true
//     사용 사례: 고유성 검증, 제약 확인, 데이터 품질
//
// 비교: Diff vs DiffBy:
//
// Diff:
//   - 직접 값 비교
//   - comparable 타입 필요
//   - 단순 요소 일치
//   - 예: Diff([1,2,3], [2,3,4])
//
// DiffBy:
//   - 키 기반 비교
//   - 요소 타입은 any, 키는 comparable 필요
//   - 유연한 일치 기준
//   - 구조체 유용: DiffBy(users, newUsers, u → u.ID)
//   - 특정 필드 무시 가능: ID로 비교, Name 변경 무시
//
// 각각 사용 시기:
//   - Diff: 단순 타입 (int, string) 또는 전체 값 중요할 때
//   - DiffBy: 식별자 필드가 있는 구조체/복잡한 타입
//   - DiffBy: 특정 속성으로 비교 (ID, 이름 등)
//   - DiffBy: 요소 타입이 comparable이 아니지만 comparable 키가 있을 때
//
// 설계 원칙:
//   - 구조화된 결과: DiffResult가 정리된 출력 제공
//   - 해시 기반: 효율성을 위한 O(1) 조회
//   - 중복 제거: 입력의 중복 자동 처리
//   - 타입 안전성: 제네릭 제약 조건으로 정확성 보장
//   - 불변성: 원본 슬라이스 절대 수정 안 함
//
// 성능 특성:
//
// 시간 복잡도:
//   - Diff: O(n+m) - 두 집합 구축 + 두 슬라이스 반복
//   - DiffBy: O(n+m) - 두 맵 구축 + 두 슬라이스 반복
//   - EqualUnordered: O(n) - 빈도 맵 구축 + 비교 (len(a)==len(b) 가정)
//   - HasDuplicates: O(n) 최악, 단락으로 평균 O(k)
//
// 공간 복잡도:
//   - Diff: O(n+m) - 두 해시 집합 + 결과 슬라이스
//   - DiffBy: O(n+m) - 두 해시 맵 + 결과 슬라이스
//   - EqualUnordered: O(n) - 두 빈도 맵
//   - HasDuplicates: O(k) (k = 첫 중복 전에 본 고유 요소)
//
// 메모리 할당:
//   - Diff/DiffBy: 결과 슬라이스 동적 할당 (크기 미지)
//   - EqualUnordered: 길이 불일치 사전 확인 (조기 종료)
//   - HasDuplicates: 중복 일찍 발견 시 최소 할당
//
// 일반적인 사용 패턴:
//
//	// 버전 간 변경 추적
//	diff := sliceutil.Diff(oldTags, newTags)
//	fmt.Printf("추가: %v, 제거: %v\n", diff.Added, diff.Removed)
//
//	// ID로 사용자 목록 비교
//	type User struct { ID int; Name string; Email string }
//	diff := sliceutil.DiffBy(oldUsers, newUsers, func(u User) int {
//	    return u.ID
//	})
//	for _, added := range diff.Added {
//	    fmt.Printf("새 사용자: %s\n", added.Name)
//	}
//
//	// 순서 무시하고 동등성 테스트
//	expected := []int{1, 2, 3}
//	actual := []int{3, 1, 2}
//	if sliceutil.EqualUnordered(expected, actual) {
//	    fmt.Println("테스트 통과")
//	}
//
//	// 고유성 검증
//	if sliceutil.HasDuplicates(usernames) {
//	    return errors.New("중복 사용자 이름 불가")
//	}
//
// 변경 감지 패턴:
//
//	// 데이터베이스 동기화
//	diff := sliceutil.DiffBy(dbRecords, apiRecords, r → r.ID)
//	// diff.Added를 데이터베이스에 삽입
//	// diff.Removed를 데이터베이스에서 삭제
//	// 필요 시 diff.Unchanged 업데이트
//
//	// 재고 변경 추적
//	diff := sliceutil.Diff(previousStock, currentStock)
//	restockNeeded := diff.Removed
//	newItems := diff.Added
//
// 테스트 패턴:
//
//	// 유연한 테스트 단언
//	func TestUnorderedResults(t *testing.T) {
//	    result := GetItems()
//	    expected := []string{"a", "b", "c"}
//	    if !sliceutil.EqualUnordered(result, expected) {
//	        t.Errorf("got %v, want %v (임의 순서)", result, expected)
//	    }
//	}
//
// 표준 라이브러리와 비교:
//   - 수동 루프 비교보다 구조화
//   - 맵을 사용한 중첩 루프보다 의도 명확
//   - 제네릭으로 타입 안전
//   - 함수형 구성 스타일
//
// 대안 접근법:
//   - 정렬된 슬라이스: 두 포인터 병합 알고리즘 사용 가능 (더 효율적)
//   - 작은 슬라이스: 선형 검색이 더 단순할 수 있음
//   - 스트리밍 데이터: 증분 diff 알고리즘 고려

// DiffResult holds the result of a diff operation between two slices.
// DiffResult는 두 슬라이스 간의 차이 연산 결과를 저장합니다.
type DiffResult[T any] struct {
	// Elements in new but not in old
	// 새 슬라이스에만 있는 요소
	Added []T
	// Elements in old but not in new
	// 이전 슬라이스에만 있는 요소
	Removed []T
	// Elements in both old and new
	// 양쪽 모두에 있는 요소
	Unchanged []T
}

// Diff compares two slices and returns the differences.
// Returns elements that were added (in new but not old), removed (in old but not new),
// and unchanged (in both).
//
// Diff는 두 슬라이스를 비교하여 차이를 반환합니다.
// 추가된 요소(새 슬라이스에만 있음), 제거된 요소(이전 슬라이스에만 있음),
// 변경되지 않은 요소(양쪽 모두에 있음)를 반환합니다.
//
// Example:
//
//	old := []int{1, 2, 3, 4}
//	new := []int{2, 3, 4, 5}
//	diff := sliceutil.Diff(old, new)
//	// diff.Added: [5]
//	// diff.Removed: [1]
//	// diff.Unchanged: [2, 3, 4]
func Diff[T comparable](old, new []T) DiffResult[T] {
	oldSet := make(map[T]bool)
	for _, v := range old {
		oldSet[v] = true
	}

	newSet := make(map[T]bool)
	for _, v := range new {
		newSet[v] = true
	}

	var added, removed, unchanged []T

	// Find added elements (in new but not in old)
	for _, v := range new {
		if !oldSet[v] {
			added = append(added, v)
		} else {
			unchanged = append(unchanged, v)
		}
	}

	// Find removed elements (in old but not in new)
	for _, v := range old {
		if !newSet[v] {
			removed = append(removed, v)
		}
	}

	// Remove duplicates from unchanged
	unchanged = Unique(unchanged)

	return DiffResult[T]{
		Added:     added,
		Removed:   removed,
		Unchanged: unchanged,
	}
}

// DiffBy compares two slices using a key function and returns the differences.
// The key function extracts a comparable key from each element for comparison.
// This is useful for comparing slices of structs or complex types.
//
// DiffBy는 키 함수를 사용하여 두 슬라이스를 비교하고 차이를 반환합니다.
// 키 함수는 각 요소에서 비교 가능한 키를 추출합니다.
// 이는 구조체나 복잡한 타입의 슬라이스를 비교할 때 유용합니다.
//
// Example:
//
//	type User struct { ID int; Name string }
//	old := []User{{1, "Alice"}, {2, "Bob"}}
//	new := []User{{2, "Bob"}, {3, "Charlie"}}
//	diff := sliceutil.DiffBy(old, new, func(u User) int { return u.ID })
//	// diff.Added: [{3, "Charlie"}]
//	// diff.Removed: [{1, "Alice"}]
//	// diff.Unchanged: [{2, "Bob"}]
func DiffBy[T any, K comparable](old, new []T, keyFunc func(T) K) DiffResult[T] {
	oldMap := make(map[K]T)
	for _, v := range old {
		oldMap[keyFunc(v)] = v
	}

	newMap := make(map[K]T)
	for _, v := range new {
		newMap[keyFunc(v)] = v
	}

	var added, removed, unchanged []T

	// Find added elements (keys in new but not in old)
	for key, v := range newMap {
		if _, exists := oldMap[key]; !exists {
			added = append(added, v)
		} else {
			unchanged = append(unchanged, v)
		}
	}

	// Find removed elements (keys in old but not in new)
	for key, v := range oldMap {
		if _, exists := newMap[key]; !exists {
			removed = append(removed, v)
		}
	}

	return DiffResult[T]{
		Added:     added,
		Removed:   removed,
		Unchanged: unchanged,
	}
}

// EqualUnordered checks if two slices contain the same elements, regardless of order.
// Returns true if both slices have the same elements with the same frequencies.
//
// EqualUnordered는 순서에 관계없이 두 슬라이스가 같은 요소를 포함하는지 확인합니다.
// 양쪽 슬라이스가 같은 빈도의 같은 요소를 가지면 true를 반환합니다.
//
// Example:
//
//	a := []int{1, 2, 3, 2}
//	b := []int{3, 2, 1, 2}
//	equal := sliceutil.EqualUnordered(a, b) // true
//
//	c := []int{1, 2, 3}
//	d := []int{1, 2, 3, 3}
//	equal := sliceutil.EqualUnordered(c, d) // false (different frequencies)
func EqualUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	// Get frequencies of both slices
	freqA := Frequencies(a)
	freqB := Frequencies(b)

	// Check if frequencies match
	if len(freqA) != len(freqB) {
		return false
	}

	for key, countA := range freqA {
		countB, exists := freqB[key]
		if !exists || countA != countB {
			return false
		}
	}

	return true
}

// HasDuplicates checks if a slice contains any duplicate elements.
// Returns true if there are duplicates, false otherwise.
//
// HasDuplicates는 슬라이스에 중복 요소가 있는지 확인합니다.
// 중복이 있으면 true를, 없으면 false를 반환합니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 2, 4}
//	hasDups := sliceutil.HasDuplicates(numbers) // true
//
//	unique := []int{1, 2, 3, 4, 5}
//	hasDups := sliceutil.HasDuplicates(unique) // false
func HasDuplicates[T comparable](slice []T) bool {
	seen := make(map[T]bool)
	for _, v := range slice {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
