package stringutil

import (
	"math"
	"testing"
)

// Test Levenshtein distance
// Levenshtein 거리 테스트
func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected int
	}{
		{"", "", 0},
		{"", "hello", 5},
		{"hello", "", 5},
		{"hello", "hello", 0},
		{"kitten", "sitting", 3},
		{"saturday", "sunday", 3},
		{"hello", "hallo", 1},
		{"안녕", "안녕하세요", 3},
		{"abc", "xyz", 3},
	}

	for _, tt := range tests {
		result := LevenshteinDistance(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("LevenshteinDistance(%q, %q) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

// Test Similarity
// 유사도 테스트
func TestSimilarity(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected float64
		delta    float64
	}{
		{"", "", 1.0, 0.001},
		{"hello", "hello", 1.0, 0.001},
		{"hello", "hallo", 0.8, 0.001},
		{"hello", "world", 0.2, 0.001},
		{"kitten", "sitting", 0.571, 0.001},
		{"안녕하세요", "안녕하세요", 1.0, 0.001},
		{"abc", "xyz", 0.0, 0.001},
	}

	for _, tt := range tests {
		result := Similarity(tt.a, tt.b)
		if math.Abs(result-tt.expected) > tt.delta {
			t.Errorf("Similarity(%q, %q) = %f, want %f", tt.a, tt.b, result, tt.expected)
		}
	}
}

// Test Hamming distance
// Hamming 거리 테스트
func TestHammingDistance(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected int
	}{
		{"", "", 0},
		{"hello", "hello", 0},
		{"karolin", "kathrin", 3},
		{"hello", "world", 4},
		{"hello", "hi", -1},    // different lengths
		{"abc", "abcd", -1},    // different lengths
		{"1011101", "1001001", 2},
		{"2173896", "2233796", 3},
	}

	for _, tt := range tests {
		result := HammingDistance(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("HammingDistance(%q, %q) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

// Test Jaro-Winkler similarity
// Jaro-Winkler 유사도 테스트
func TestJaroWinklerSimilarity(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected float64
		delta    float64
	}{
		{"", "", 1.0, 0.001},
		{"hello", "hello", 1.0, 0.001},
		{"martha", "marhta", 0.961, 0.001},
		{"DIXON", "DICKSONX", 0.813, 0.001},
		{"hello", "world", 0.466, 0.001},
		{"abc", "xyz", 0.0, 0.001},
	}

	for _, tt := range tests {
		result := JaroWinklerSimilarity(tt.a, tt.b)
		if math.Abs(result-tt.expected) > tt.delta {
			t.Errorf("JaroWinklerSimilarity(%q, %q) = %f, want %f", tt.a, tt.b, result, tt.expected)
		}
	}
}

// Test symmetry
// 대칭성 테스트
func TestDistanceSymmetry(t *testing.T) {
	pairs := [][]string{
		{"hello", "world"},
		{"kitten", "sitting"},
		{"안녕", "안녕하세요"},
	}

	for _, pair := range pairs {
		a, b := pair[0], pair[1]

		// Levenshtein should be symmetric
		d1 := LevenshteinDistance(a, b)
		d2 := LevenshteinDistance(b, a)
		if d1 != d2 {
			t.Errorf("LevenshteinDistance not symmetric: %q vs %q: %d != %d", a, b, d1, d2)
		}

		// Similarity should be symmetric
		s1 := Similarity(a, b)
		s2 := Similarity(b, a)
		if math.Abs(s1-s2) > 0.001 {
			t.Errorf("Similarity not symmetric: %q vs %q: %f != %f", a, b, s1, s2)
		}

		// Hamming (if same length)
		if len([]rune(a)) == len([]rune(b)) {
			h1 := HammingDistance(a, b)
			h2 := HammingDistance(b, a)
			if h1 != h2 {
				t.Errorf("HammingDistance not symmetric: %q vs %q: %d != %d", a, b, h1, h2)
			}
		}

		// Jaro-Winkler should be symmetric
		jw1 := JaroWinklerSimilarity(a, b)
		jw2 := JaroWinklerSimilarity(b, a)
		if math.Abs(jw1-jw2) > 0.001 {
			t.Errorf("JaroWinklerSimilarity not symmetric: %q vs %q: %f != %f", a, b, jw1, jw2)
		}
	}
}

// Test triangle inequality for Levenshtein
// Levenshtein의 삼각 부등식 테스트
func TestLevenshteinTriangleInequality(t *testing.T) {
	triplets := [][]string{
		{"hello", "hallo", "hillo"},
		{"abc", "def", "ghi"},
	}

	for _, triplet := range triplets {
		a, b, c := triplet[0], triplet[1], triplet[2]

		dAB := LevenshteinDistance(a, b)
		dBC := LevenshteinDistance(b, c)
		dAC := LevenshteinDistance(a, c)

		// Triangle inequality: d(a,c) <= d(a,b) + d(b,c)
		if dAC > dAB+dBC {
			t.Errorf("Triangle inequality violated: d(%q,%q)=%d > d(%q,%q)+d(%q,%q)=%d+%d",
				a, c, dAC, a, b, b, c, dAB, dBC)
		}
	}
}

// Test Unicode support
// 유니코드 지원 테스트
func TestDistanceUnicode(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
	}{
		{"Korean", "안녕하세요", "안녕"},
		{"Japanese", "こんにちは", "こんにちは"},
		{"Emoji", "🎉🎊🎈", "🎉🎊"},
		{"Mixed", "hello世界", "hello🌍"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just check that functions don't panic with Unicode
			_ = LevenshteinDistance(tt.a, tt.b)
			_ = Similarity(tt.a, tt.b)
			_ = HammingDistance(tt.a, tt.b)
			_ = JaroWinklerSimilarity(tt.a, tt.b)
		})
	}
}

// Test edge cases
// 엣지 케이스 테스트
func TestDistanceEdgeCases(t *testing.T) {
	// Empty strings
	// 빈 문자열
	if LevenshteinDistance("", "") != 0 {
		t.Error("Levenshtein distance of empty strings should be 0")
	}
	if Similarity("", "") != 1.0 {
		t.Error("Similarity of empty strings should be 1.0")
	}
	if HammingDistance("", "") != 0 {
		t.Error("Hamming distance of empty strings should be 0")
	}
	if JaroWinklerSimilarity("", "") != 1.0 {
		t.Error("Jaro-Winkler similarity of empty strings should be 1.0")
	}

	// One empty string
	// 한 쪽이 빈 문자열
	if LevenshteinDistance("hello", "") != 5 {
		t.Error("Levenshtein distance to empty string should be length of non-empty string")
	}
	if Similarity("hello", "") != 0.0 {
		t.Error("Similarity to empty string should be 0.0")
	}

	// Identical strings
	// 동일한 문자열
	if LevenshteinDistance("hello", "hello") != 0 {
		t.Error("Levenshtein distance of identical strings should be 0")
	}
	if Similarity("hello", "hello") != 1.0 {
		t.Error("Similarity of identical strings should be 1.0")
	}
	if HammingDistance("hello", "hello") != 0 {
		t.Error("Hamming distance of identical strings should be 0")
	}
	if JaroWinklerSimilarity("hello", "hello") != 1.0 {
		t.Error("Jaro-Winkler similarity of identical strings should be 1.0")
	}
}

// Benchmark distance functions
// 거리 함수 벤치마크
func BenchmarkLevenshteinDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = LevenshteinDistance("kitten", "sitting")
	}
}

func BenchmarkLevenshteinDistanceLong(b *testing.B) {
	a := "The quick brown fox jumps over the lazy dog"
	b2 := "The quack brown fox jumped over the crazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = LevenshteinDistance(a, b2)
	}
}

func BenchmarkSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Similarity("hello", "hallo")
	}
}

func BenchmarkHammingDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HammingDistance("karolin", "kathrin")
	}
}

func BenchmarkJaroWinklerSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = JaroWinklerSimilarity("martha", "marhta")
	}
}
