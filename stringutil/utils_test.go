package stringutil

import (
	"reflect"
	"strings"
	"testing"
)

// TestCountWords tests the CountWords function
// CountWords 함수 테스트
func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"simple sentence", "hello world", 2},
		{"multiple spaces", "hello   world  foo", 3},
		{"leading/trailing spaces", "  hello world  ", 2},
		{"single word", "hello", 1},
		{"empty string", "", 0},
		{"only spaces", "   ", 0},
		{"unicode words", "안녕 세계", 2},
		{"mixed", "hello 세계 world", 3},
		{"with punctuation", "hello, world!", 2},
		{"newlines", "hello\nworld\nfoo", 3},
		{"tabs", "hello\tworld", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.s)
			if result != tt.expected {
				t.Errorf("CountWords(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}

// TestCountOccurrences tests the CountOccurrences function
// CountOccurrences 함수 테스트
func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		substr   string
		expected int
	}{
		{"single occurrence", "hello world", "world", 1},
		{"multiple occurrences", "hello hello hello", "hello", 3},
		{"overlapping", "aaaa", "aa", 2}, // strings.Count doesn't count overlapping
		{"no occurrence", "hello world", "foo", 0},
		{"empty substr", "hello", "", 6}, // len("hello") + 1
		{"empty string", "", "hello", 0},
		{"case sensitive", "Hello hello HELLO", "hello", 1},
		{"unicode", "안녕하세요 안녕", "안녕", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountOccurrences(tt.s, tt.substr)
			if result != tt.expected {
				t.Errorf("CountOccurrences(%q, %q) = %d, want %d",
					tt.s, tt.substr, result, tt.expected)
			}
		})
	}
}

// TestJoin tests the Join function
// Join 함수 테스트
func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		sep      string
		expected string
	}{
		{"simple join", []string{"hello", "world"}, " ", "hello world"},
		{"comma separator", []string{"a", "b", "c"}, ",", "a,b,c"},
		{"empty separator", []string{"hello", "world"}, "", "helloworld"},
		{"single element", []string{"hello"}, ",", "hello"},
		{"empty slice", []string{}, ",", ""},
		{"with empty strings", []string{"hello", "", "world"}, " ", "hello  world"},
		{"unicode", []string{"안녕", "세계"}, " ", "안녕 세계"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Join(tt.strs, tt.sep)
			if result != tt.expected {
				t.Errorf("Join(%v, %q) = %q, want %q",
					tt.strs, tt.sep, result, tt.expected)
			}
		})
	}
}

// TestMap tests the Map function
// Map 함수 테스트
func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		fn       func(string) string
		expected []string
	}{
		{
			"uppercase",
			[]string{"hello", "world"},
			strings.ToUpper,
			[]string{"HELLO", "WORLD"},
		},
		{
			"lowercase",
			[]string{"HELLO", "WORLD"},
			strings.ToLower,
			[]string{"hello", "world"},
		},
		{
			"add prefix",
			[]string{"hello", "world"},
			func(s string) string { return "pre_" + s },
			[]string{"pre_hello", "pre_world"},
		},
		{
			"empty slice",
			[]string{},
			strings.ToUpper,
			[]string{},
		},
		{
			"identity function",
			[]string{"hello", "world"},
			func(s string) string { return s },
			[]string{"hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.strs, tt.fn)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Map(%v, fn) = %v, want %v",
					tt.strs, result, tt.expected)
			}
		})
	}
}

// TestFilter tests the Filter function
// Filter 함수 테스트
func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		fn       func(string) bool
		expected []string
	}{
		{
			"filter non-empty",
			[]string{"hello", "", "world", ""},
			func(s string) bool { return s != "" },
			[]string{"hello", "world"},
		},
		{
			"filter by length",
			[]string{"a", "bb", "ccc", "dddd"},
			func(s string) bool { return len(s) > 2 },
			[]string{"ccc", "dddd"},
		},
		{
			"filter by prefix",
			[]string{"hello", "world", "help", "foo"},
			func(s string) bool { return strings.HasPrefix(s, "he") },
			[]string{"hello", "help"},
		},
		{
			"no matches",
			[]string{"hello", "world"},
			func(s string) bool { return false },
			[]string{},
		},
		{
			"all matches",
			[]string{"hello", "world"},
			func(s string) bool { return true },
			[]string{"hello", "world"},
		},
		{
			"empty slice",
			[]string{},
			func(s string) bool { return true },
			[]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter(tt.strs, tt.fn)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Filter(%v, fn) = %v, want %v",
					tt.strs, result, tt.expected)
			}
		})
	}
}

// TestPadLeft tests the PadLeft function
// PadLeft 함수 테스트
func TestPadLeft(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		length   int
		pad      string
		expected string
	}{
		{"pad with spaces", "hello", 10, " ", "     hello"},
		{"pad with zeros", "123", 5, "0", "00123"},
		{"pad with dashes", "test", 8, "-", "----test"},
		{"no padding needed", "hello", 3, " ", "hello"},
		{"exact length", "hello", 5, " ", "hello"},
		{"empty string", "", 5, "*", "*****"},
		{"unicode pad", "hello", 10, "★", "★★★★★hello"},
		{"multi-char pad", "hi", 10, "ab", "ababababababababhi"}, // "ab" repeated 8 times = 16 chars
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PadLeft(tt.s, tt.length, tt.pad)
			if result != tt.expected {
				t.Errorf("PadLeft(%q, %d, %q) = %q, want %q",
					tt.s, tt.length, tt.pad, result, tt.expected)
			}
		})
	}
}

// TestPadRight tests the PadRight function
// PadRight 함수 테스트
func TestPadRight(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		length   int
		pad      string
		expected string
	}{
		{"pad with spaces", "hello", 10, " ", "hello     "},
		{"pad with zeros", "123", 5, "0", "12300"},
		{"pad with dashes", "test", 8, "-", "test----"},
		{"no padding needed", "hello", 3, " ", "hello"},
		{"exact length", "hello", 5, " ", "hello"},
		{"empty string", "", 5, "*", "*****"},
		{"unicode pad", "hello", 10, "★", "hello★★★★★"},
		{"multi-char pad", "hi", 10, "ab", "hiabababababababab"}, // "ab" repeated 8 times = 16 chars
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PadRight(tt.s, tt.length, tt.pad)
			if result != tt.expected {
				t.Errorf("PadRight(%q, %d, %q) = %q, want %q",
					tt.s, tt.length, tt.pad, result, tt.expected)
			}
		})
	}
}

// TestLines tests the Lines function
// Lines 함수 테스트
func TestLines(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []string
	}{
		{"simple lines", "hello\nworld", []string{"hello", "world"}},
		{"multiple newlines", "a\nb\nc\n", []string{"a", "b", "c", ""}},
		{"windows newlines", "hello\r\nworld", []string{"hello\r", "world"}},
		{"single line", "hello", []string{"hello"}},
		{"empty string", "", []string{""}},
		{"only newline", "\n", []string{"", ""}},
		{"trailing newline", "hello\nworld\n", []string{"hello", "world", ""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Lines(tt.s)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Lines(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}

// TestWords tests the Words function
// Words 함수 테스트
func TestWords(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []string
	}{
		{"simple words", "hello world", []string{"hello", "world"}},
		{"multiple spaces", "hello   world  foo", []string{"hello", "world", "foo"}},
		{"leading/trailing spaces", "  hello world  ", []string{"hello", "world"}},
		{"single word", "hello", []string{"hello"}},
		{"empty string", "", []string{}},
		{"only spaces", "   ", []string{}},
		{"unicode words", "안녕 세계", []string{"안녕", "세계"}},
		{"tabs and newlines", "hello\tworld\nfoo", []string{"hello", "world", "foo"}},
		{"mixed whitespace", "  hello\n\tworld  \nfoo  ", []string{"hello", "world", "foo"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Words(tt.s)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Words(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}

// Benchmarks
// 벤치마크

func BenchmarkCountWords(b *testing.B) {
	s := "hello world foo bar baz qux quux corge grault garply"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountWords(s)
	}
}

func BenchmarkCountOccurrences(b *testing.B) {
	s := "hello hello hello world world world foo bar baz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountOccurrences(s, "hello")
	}
}

func BenchmarkJoin(b *testing.B) {
	strs := []string{"hello", "world", "foo", "bar", "baz"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(strs, ",")
	}
}

func BenchmarkMap(b *testing.B) {
	strs := []string{"hello", "world", "foo", "bar", "baz"}
	fn := strings.ToUpper
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map(strs, fn)
	}
}

func BenchmarkFilter(b *testing.B) {
	strs := []string{"hello", "world", "foo", "bar", "baz", "qux", "quux"}
	fn := func(s string) bool { return len(s) > 3 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(strs, fn)
	}
}

func BenchmarkPadLeft(b *testing.B) {
	s := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PadLeft(s, 20, " ")
	}
}

func BenchmarkPadRight(b *testing.B) {
	s := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PadRight(s, 20, " ")
	}
}
