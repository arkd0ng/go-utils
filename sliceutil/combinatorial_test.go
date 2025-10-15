package sliceutil

import (
	"reflect"
	"testing"
)

// TestPermutations tests the Permutations function.
// TestPermutations는 Permutations 함수를 테스트합니다.
func TestPermutations(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int // expected number of permutations
	}{
		{
			name:     "Empty slice / 빈 슬라이스",
			input:    []int{},
			expected: 1, // [[]]
		},
		{
			name:     "Single element / 단일 요소",
			input:    []int{1},
			expected: 1, // [[1]]
		},
		{
			name:     "Two elements / 두 요소",
			input:    []int{1, 2},
			expected: 2, // [[1,2], [2,1]]
		},
		{
			name:     "Three elements / 세 요소",
			input:    []int{1, 2, 3},
			expected: 6, // 3! = 6
		},
		{
			name:     "Four elements / 네 요소",
			input:    []int{1, 2, 3, 4},
			expected: 24, // 4! = 24
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Permutations(tt.input)

			// Check count / 개수 확인
			if len(result) != tt.expected {
				t.Errorf("Permutations() returned %d permutations, want %d", len(result), tt.expected)
			}

			// For non-empty input, verify all permutations have correct length
			// 비어 있지 않은 입력의 경우 모든 순열이 올바른 길이를 가지는지 확인
			if len(tt.input) > 0 {
				for i, perm := range result {
					if len(perm) != len(tt.input) {
						t.Errorf("Permutation %d has length %d, want %d", i, len(perm), len(tt.input))
					}
				}
			}
		})
	}
}

// TestPermutationsString tests Permutations with string slices.
// TestPermutationsString은 문자열 슬라이스로 Permutations를 테스트합니다.
func TestPermutationsString(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := Permutations(input)

	// Should have 3! = 6 permutations
	// 3! = 6개의 순열을 가져야 함
	if len(result) != 6 {
		t.Errorf("Permutations() returned %d permutations, want 6", len(result))
	}

	// Check one specific permutation exists
	// 특정 순열이 존재하는지 확인
	expectedPerm := []string{"a", "b", "c"}
	found := false
	for _, perm := range result {
		if reflect.DeepEqual(perm, expectedPerm) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected permutation %v not found in results", expectedPerm)
	}
}

// TestPermutationsUniqueness verifies all permutations are unique.
// TestPermutationsUniqueness는 모든 순열이 고유한지 확인합니다.
func TestPermutationsUniqueness(t *testing.T) {
	input := []int{1, 2, 3}
	result := Permutations(input)

	// Create a map to track seen permutations
	// 본 순열을 추적하기 위한 맵 생성
	seen := make(map[string]bool)
	for _, perm := range result {
		// Convert to string for map key
		// 맵 키를 위해 문자열로 변환
		key := ""
		for _, v := range perm {
			key += string(rune(v + '0'))
		}

		if seen[key] {
			t.Errorf("Duplicate permutation found: %v", perm)
		}
		seen[key] = true
	}
}

// TestCombinations tests the Combinations function.
// TestCombinations는 Combinations 함수를 테스트합니다.
func TestCombinations(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected int // expected number of combinations
	}{
		{
			name:     "Empty slice / 빈 슬라이스",
			input:    []int{},
			k:        0,
			expected: 1, // [[]]
		},
		{
			name:     "k = 0 / k = 0",
			input:    []int{1, 2, 3},
			k:        0,
			expected: 1, // [[]]
		},
		{
			name:     "k = n / k = n",
			input:    []int{1, 2, 3},
			k:        3,
			expected: 1, // [[1,2,3]]
		},
		{
			name:     "C(4, 2) / C(4, 2)",
			input:    []int{1, 2, 3, 4},
			k:        2,
			expected: 6, // 4!/(2!*2!) = 6
		},
		{
			name:     "C(5, 3) / C(5, 3)",
			input:    []int{1, 2, 3, 4, 5},
			k:        3,
			expected: 10, // 5!/(3!*2!) = 10
		},
		{
			name:     "k > n / k > n",
			input:    []int{1, 2},
			k:        3,
			expected: 0, // impossible
		},
		{
			name:     "k < 0 / k < 0",
			input:    []int{1, 2, 3},
			k:        -1,
			expected: 0, // invalid
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Combinations(tt.input, tt.k)

			// Check count / 개수 확인
			if len(result) != tt.expected {
				t.Errorf("Combinations() returned %d combinations, want %d", len(result), tt.expected)
			}

			// For valid combinations, verify all have correct length
			// 유효한 조합의 경우 모두 올바른 길이를 가지는지 확인
			if tt.k >= 0 && tt.k <= len(tt.input) {
				for i, comb := range result {
					if len(comb) != tt.k {
						t.Errorf("Combination %d has length %d, want %d", i, len(comb), tt.k)
					}
				}
			}
		})
	}
}

// TestCombinationsString tests Combinations with string slices.
// TestCombinationsString은 문자열 슬라이스로 Combinations를 테스트합니다.
func TestCombinationsString(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	result := Combinations(input, 2)

	// C(4, 2) = 6
	if len(result) != 6 {
		t.Errorf("Combinations() returned %d combinations, want 6", len(result))
	}

	// Check one specific combination exists
	// 특정 조합이 존재하는지 확인
	expectedComb := []string{"a", "b"}
	found := false
	for _, comb := range result {
		if reflect.DeepEqual(comb, expectedComb) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected combination %v not found in results", expectedComb)
	}
}

// TestCombinationsUniqueness verifies all combinations are unique.
// TestCombinationsUniqueness는 모든 조합이 고유한지 확인합니다.
func TestCombinationsUniqueness(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := Combinations(input, 2)

	// Create a map to track seen combinations
	// 본 조합을 추적하기 위한 맵 생성
	seen := make(map[string]bool)
	for _, comb := range result {
		// Convert to string for map key
		// 맵 키를 위해 문자열로 변환
		key := ""
		for _, v := range comb {
			key += string(rune(v + '0'))
		}

		if seen[key] {
			t.Errorf("Duplicate combination found: %v", comb)
		}
		seen[key] = true
	}
}

// TestCombinationsContent verifies specific combinations are correct.
// TestCombinationsContent는 특정 조합이 올바른지 확인합니다.
func TestCombinationsContent(t *testing.T) {
	input := []int{1, 2, 3}
	result := Combinations(input, 2)

	// C(3, 2) = 3: [1,2], [1,3], [2,3]
	expected := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}

	if len(result) != len(expected) {
		t.Fatalf("Combinations() returned %d combinations, want %d", len(result), len(expected))
	}

	// Check all expected combinations exist
	// 모든 예상 조합이 존재하는지 확인
	for _, exp := range expected {
		found := false
		for _, comb := range result {
			if reflect.DeepEqual(comb, exp) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected combination %v not found in results", exp)
		}
	}
}

// BenchmarkPermutations benchmarks the Permutations function.
// BenchmarkPermutations는 Permutations 함수를 벤치마크합니다.
func BenchmarkPermutations(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Permutations(input)
	}
}

// BenchmarkCombinations benchmarks the Combinations function.
// BenchmarkCombinations는 Combinations 함수를 벤치마크합니다.
func BenchmarkCombinations(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Combinations(input, 5)
	}
}
