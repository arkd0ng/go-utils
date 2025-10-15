package main

import (
	"fmt"

	"github.com/arkd0ng/go-utils/maputil"
)

func main() {
	fmt.Println("=== maputil Package Examples ===")
	fmt.Println()

	// 1. Basic Operations / 기본 작업
	fmt.Println("1. Basic Operations / 기본 작업")
	basicOperations()
	fmt.Println()

	// 2. Transformation / 변환
	fmt.Println("2. Transformation / 변환")
	transformations()
	fmt.Println()

	// 3. Aggregation / 집계
	fmt.Println("3. Aggregation / 집계")
	aggregations()
	fmt.Println()

	// 4. Merge Operations / 병합 작업
	fmt.Println("4. Merge Operations / 병합 작업")
	mergeOperations()
	fmt.Println()

	// 5. Filter Operations / 필터 작업
	fmt.Println("5. Filter Operations / 필터 작업")
	filterOperations()
	fmt.Println()

	// 6. Conversion / 변환
	fmt.Println("6. Conversion / 변환")
	conversions()
	fmt.Println()

	// 7. Predicate Checks / 조건 검사
	fmt.Println("7. Predicate Checks / 조건 검사")
	predicates()
	fmt.Println()

	// 8. Key Operations / 키 작업
	fmt.Println("8. Key Operations / 키 작업")
	keyOperations()
	fmt.Println()

	// 9. Value Operations / 값 작업
	fmt.Println("9. Value Operations / 값 작업")
	valueOperations()
	fmt.Println()

	// 10. Comparison / 비교
	fmt.Println("10. Comparison / 비교")
	comparisons()
	fmt.Println()
}

func basicOperations() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// Get / 가져오기
	value, ok := maputil.Get(m, "a")
	fmt.Printf("Get('a'): %d, exists: %v\n", value, ok)

	// GetOr / 기본값과 함께 가져오기
	value = maputil.GetOr(m, "d", 10)
	fmt.Printf("GetOr('d', 10): %d\n", value)

	// Set / 설정
	result := maputil.Set(m, "d", 4)
	fmt.Printf("Set('d', 4): %v\n", result)

	// Delete / 삭제
	result = maputil.Delete(m, "a")
	fmt.Printf("Delete('a'): %v\n", result)

	// Clone / 복제
	cloned := maputil.Clone(m)
	fmt.Printf("Clone: %v\n", cloned)

	// Equal / 동등 비교
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	equal := maputil.Equal(m, m2)
	fmt.Printf("Equal: %v\n", equal)
}

func transformations() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// Map / 맵
	result := maputil.Map(m, func(k string, v int) string {
		return fmt.Sprintf("%s=%d", k, v)
	})
	fmt.Printf("Map: %v\n", result)

	// MapValues / 값 변환
	doubled := maputil.MapValues(m, func(v int) int {
		return v * 2
	})
	fmt.Printf("MapValues (doubled): %v\n", doubled)

	// Invert / 반전
	inverted := maputil.Invert(m)
	fmt.Printf("Invert: %v\n", inverted)

	// Partition / 분할
	even, odd := maputil.Partition(m, func(k string, v int) bool {
		return v%2 == 0
	})
	fmt.Printf("Partition - even: %v, odd: %v\n", even, odd)

	// Compact / 압축 (zero 값 제거)
	sparse := map[string]int{"a": 1, "b": 0, "c": 3, "d": 0}
	compacted := maputil.Compact(sparse)
	fmt.Printf("Compact: %v\n", compacted)
}

func aggregations() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Sum / 합계
	total := maputil.Sum(m)
	fmt.Printf("Sum: %d\n", total)

	// Average / 평균
	avg := maputil.Average(m)
	fmt.Printf("Average: %.2f\n", avg)

	// Min / 최소값
	minKey, minVal, _ := maputil.Min(m)
	fmt.Printf("Min: %s=%d\n", minKey, minVal)

	// Max / 최대값
	maxKey, maxVal, _ := maputil.Max(m)
	fmt.Printf("Max: %s=%d\n", maxKey, maxVal)

	// Reduce / 축소
	sum := maputil.Reduce(m, 0, func(acc int, k string, v int) int {
		return acc + v
	})
	fmt.Printf("Reduce (sum): %d\n", sum)

	// GroupBy / 그룹화
	type User struct {
		Name string
		City string
	}
	users := []User{
		{Name: "Alice", City: "Seoul"},
		{Name: "Bob", City: "Seoul"},
		{Name: "Charlie", City: "Busan"},
	}
	byCity := maputil.GroupBy[string, User, string](users, func(u User) string {
		return u.City
	})
	fmt.Printf("GroupBy City: %v\n", byCity)

	// CountBy / 개수 세기
	counts := maputil.CountBy[string, User, string](users, func(u User) string {
		return u.City
	})
	fmt.Printf("CountBy City: %v\n", counts)
}

func mergeOperations() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	m3 := map[string]int{"c": 5, "d": 6}

	// Merge / 병합
	merged := maputil.Merge(m1, m2, m3)
	fmt.Printf("Merge: %v\n", merged)

	// MergeWith / 사용자 정의 병합
	mergedWith := maputil.MergeWith(func(old, new int) int {
		return old + new // Sum on conflict / 충돌 시 합산
	}, m1, m2)
	fmt.Printf("MergeWith (sum): %v\n", mergedWith)

	// Intersection / 교집합
	m4 := map[string]int{"a": 1, "b": 2, "c": 3}
	m5 := map[string]int{"b": 20, "c": 30, "d": 40}
	intersection := maputil.Intersection(m4, m5)
	fmt.Printf("Intersection: %v\n", intersection)

	// Difference / 차집합
	diff := maputil.Difference(m4, m5)
	fmt.Printf("Difference: %v\n", diff)

	// SymmetricDifference / 대칭 차집합
	symDiff := maputil.SymmetricDifference(m4, m5)
	fmt.Printf("SymmetricDifference: %v\n", symDiff)
}

func filterOperations() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Filter / 필터
	filtered := maputil.Filter(m, func(k string, v int) bool {
		return v > 2
	})
	fmt.Printf("Filter (v > 2): %v\n", filtered)

	// FilterKeys / 키로 필터
	filteredKeys := maputil.FilterKeys(m, func(k string) bool {
		return k >= "b"
	})
	fmt.Printf("FilterKeys (k >= 'b'): %v\n", filteredKeys)

	// FilterValues / 값으로 필터
	filteredValues := maputil.FilterValues(m, func(v int) bool {
		return v%2 == 0
	})
	fmt.Printf("FilterValues (even): %v\n", filteredValues)

	// Pick / 선택
	picked := maputil.Pick(m, "a", "c")
	fmt.Printf("Pick ('a', 'c'): %v\n", picked)

	// Omit / 제외
	omitted := maputil.Omit(m, "b", "d")
	fmt.Printf("Omit ('b', 'd'): %v\n", omitted)
}

func conversions() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// Keys / 키
	keys := maputil.Keys(m)
	fmt.Printf("Keys: %v\n", keys)

	// Values / 값
	values := maputil.Values(m)
	fmt.Printf("Values: %v\n", values)

	// Entries / 항목
	entries := maputil.Entries(m)
	fmt.Printf("Entries: %v\n", entries)

	// FromEntries / 항목으로부터
	newMap := maputil.FromEntries(entries)
	fmt.Printf("FromEntries: %v\n", newMap)

	// ToJSON / JSON으로
	json, _ := maputil.ToJSON(m)
	fmt.Printf("ToJSON: %s\n", json)

	// ToSlice / 슬라이스로
	slice := maputil.ToSlice(m, func(k string, v int) string {
		return fmt.Sprintf("%s=%d", k, v)
	})
	fmt.Printf("ToSlice: %v\n", slice)
}

func predicates() {
	m := map[string]int{"a": 2, "b": 4, "c": 6}

	// Every / 모두
	allEven := maputil.Every(m, func(k string, v int) bool {
		return v%2 == 0
	})
	fmt.Printf("Every (even): %v\n", allEven)

	// Some / 일부
	hasGreaterThan3 := maputil.Some(m, func(k string, v int) bool {
		return v > 3
	})
	fmt.Printf("Some (v > 3): %v\n", hasGreaterThan3)

	// None / 없음
	noOdd := maputil.None(m, func(k string, v int) bool {
		return v%2 == 1
	})
	fmt.Printf("None (odd): %v\n", noOdd)

	// HasValue / 값 존재
	hasValue := maputil.HasValue(m, 4)
	fmt.Printf("HasValue(4): %v\n", hasValue)

	// HasEntry / 항목 존재
	hasEntry := maputil.HasEntry(m, "b", 4)
	fmt.Printf("HasEntry('b', 4): %v\n", hasEntry)
}

func keyOperations() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// KeysSorted / 정렬된 키
	keys := maputil.KeysSorted(m)
	fmt.Printf("KeysSorted: %v\n", keys)

	// FindKey / 키 찾기
	key, found := maputil.FindKey(m, func(k string, v int) bool {
		return v > 2
	})
	fmt.Printf("FindKey (v > 2): %s, found: %v\n", key, found)

	// RenameKey / 키 이름 변경
	renamed := maputil.RenameKey(m, "b", "B")
	fmt.Printf("RenameKey ('b' -> 'B'): %v\n", renamed)

	// SwapKeys / 키 교환
	swapped := maputil.SwapKeys(m, "a", "c")
	fmt.Printf("SwapKeys ('a', 'c'): %v\n", swapped)
}

func valueOperations() {
	m := map[string]int{"a": 3, "b": 1, "c": 2}

	// ValuesSorted / 정렬된 값
	values := maputil.ValuesSorted(m)
	fmt.Printf("ValuesSorted: %v\n", values)

	// UniqueValues / 고유 값
	sparse := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3, "e": 2}
	unique := maputil.UniqueValues(sparse)
	fmt.Printf("UniqueValues: %v\n", unique)

	// ReplaceValue / 값 대체
	replaced := maputil.ReplaceValue(sparse, 1, 10)
	fmt.Printf("ReplaceValue (1 -> 10): %v\n", replaced)

	// UpdateValues / 값 업데이트
	updated := maputil.UpdateValues(m, func(k string, v int) int {
		return v * 10
	})
	fmt.Printf("UpdateValues (*10): %v\n", updated)
}

func comparisons() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 20, "d": 4}

	// Diff / 차이
	diff := maputil.Diff(m1, m2)
	fmt.Printf("Diff: %v\n", diff)

	// DiffKeys / 차이 키
	diffKeys := maputil.DiffKeys(m1, m2)
	fmt.Printf("DiffKeys: %v\n", diffKeys)

	// Compare / 비교
	added, removed, modified := maputil.Compare(m1, m2)
	fmt.Printf("Compare:\n")
	fmt.Printf("  Added: %v\n", added)
	fmt.Printf("  Removed: %v\n", removed)
	fmt.Printf("  Modified: %v\n", modified)

	// CommonKeys / 공통 키
	m3 := map[string]int{"b": 100, "c": 200}
	common := maputil.CommonKeys(m1, m2, m3)
	fmt.Printf("CommonKeys: %v\n", common)

	// AllKeys / 모든 키
	allKeys := maputil.AllKeys(m1, m2, m3)
	fmt.Printf("AllKeys: %v\n", allKeys)
}
