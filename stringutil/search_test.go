package stringutil

import (
	"testing"
)

// TestContainsAny tests the ContainsAny function / ContainsAny 함수 테스트
func TestContainsAny(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		substrs  []string
		expected bool
	}{
		{"contains one", "hello world", []string{"hello", "foo"}, true},
		{"contains multiple", "hello world", []string{"hello", "world"}, true},
		{"contains none", "hello world", []string{"foo", "bar"}, false},
		{"empty substrs", "hello world", []string{}, false},
		{"empty string", "", []string{"hello"}, false},
		{"case sensitive", "Hello World", []string{"hello"}, false},
		{"partial match", "hello world", []string{"hel"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsAny(tt.s, tt.substrs)
			if result != tt.expected {
				t.Errorf("ContainsAny(%q, %v) = %v, want %v",
					tt.s, tt.substrs, result, tt.expected)
			}
		})
	}
}

// TestContainsAll tests the ContainsAll function / ContainsAll 함수 테스트
func TestContainsAll(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		substrs  []string
		expected bool
	}{
		{"contains all", "hello world foo", []string{"hello", "world"}, true},
		{"missing one", "hello world", []string{"hello", "foo"}, false},
		{"empty substrs", "hello world", []string{}, true},
		{"empty string", "", []string{"hello"}, false},
		{"case sensitive", "Hello World", []string{"Hello", "world"}, false},
		{"partial matches", "hello world", []string{"hel", "wor"}, true},
		{"duplicate substrs", "hello world", []string{"hello", "hello"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsAll(tt.s, tt.substrs)
			if result != tt.expected {
				t.Errorf("ContainsAll(%q, %v) = %v, want %v",
					tt.s, tt.substrs, result, tt.expected)
			}
		})
	}
}

// TestStartsWithAny tests the StartsWithAny function / StartsWithAny 함수 테스트
func TestStartsWithAny(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		prefixes []string
		expected bool
	}{
		{"starts with one", "hello world", []string{"hello", "foo"}, true},
		{"starts with none", "hello world", []string{"foo", "bar"}, false},
		{"empty prefixes", "hello world", []string{}, false},
		{"empty string", "", []string{"hello"}, false},
		{"case sensitive", "Hello World", []string{"hello"}, false},
		{"exact match", "hello", []string{"hello"}, true},
		{"partial prefix", "hello world", []string{"hel"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartsWithAny(tt.s, tt.prefixes)
			if result != tt.expected {
				t.Errorf("StartsWithAny(%q, %v) = %v, want %v",
					tt.s, tt.prefixes, result, tt.expected)
			}
		})
	}
}

// TestEndsWithAny tests the EndsWithAny function / EndsWithAny 함수 테스트
func TestEndsWithAny(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		suffixes []string
		expected bool
	}{
		{"ends with one", "hello world", []string{"world", "foo"}, true},
		{"ends with none", "hello world", []string{"foo", "bar"}, false},
		{"empty suffixes", "hello world", []string{}, false},
		{"empty string", "", []string{"world"}, false},
		{"case sensitive", "Hello World", []string{"world"}, false},
		{"exact match", "world", []string{"world"}, true},
		{"partial suffix", "hello world", []string{"rld"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndsWithAny(tt.s, tt.suffixes)
			if result != tt.expected {
				t.Errorf("EndsWithAny(%q, %v) = %v, want %v",
					tt.s, tt.suffixes, result, tt.expected)
			}
		})
	}
}

// TestReplaceAll tests the ReplaceAll function / ReplaceAll 함수 테스트
func TestReplaceAll(t *testing.T) {
	tests := []struct {
		name         string
		s            string
		replacements map[string]string
		expected     string
	}{
		{
			"single replacement",
			"hello world",
			map[string]string{"hello": "hi"},
			"hi world",
		},
		{
			"multiple replacements",
			"hello world foo bar",
			map[string]string{"hello": "hi", "world": "universe"},
			"hi universe foo bar",
		},
		{
			"overlapping replacements",
			"hello hello world",
			map[string]string{"hello": "hi", "world": "universe"},
			"hi hi universe",
		},
		{
			"empty replacements",
			"hello world",
			map[string]string{},
			"hello world",
		},
		{
			"no matches",
			"hello world",
			map[string]string{"foo": "bar"},
			"hello world",
		},
		{
			"empty string",
			"",
			map[string]string{"hello": "hi"},
			"",
		},
		{
			"replace with empty",
			"hello world",
			map[string]string{"hello": "", "world": ""},
			" ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceAll(tt.s, tt.replacements)
			if result != tt.expected {
				t.Errorf("ReplaceAll(%q, %v) = %q, want %q",
					tt.s, tt.replacements, result, tt.expected)
			}
		})
	}
}

// TestReplaceIgnoreCase tests the ReplaceIgnoreCase function / ReplaceIgnoreCase 함수 테스트
func TestReplaceIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		old      string
		new      string
		expected string
	}{
		{"lowercase match", "hello world", "hello", "hi", "hi world"},
		{"uppercase match", "HELLO WORLD", "hello", "hi", "hi WORLD"},
		{"mixed case match", "HeLLo WoRLD", "hello", "hi", "hi WoRLD"},
		{"multiple matches", "hello HELLO HeLLo", "hello", "hi", "hi hi hi"},
		{"no match", "hello world", "foo", "bar", "hello world"},
		{"empty old", "hello world", "", "hi", "hello world"},
		{"empty string", "", "hello", "hi", ""},
		{"case insensitive partial", "hello world", "ELLO", "i", "hi world"},
		{"unicode match", "café CAFÉ", "café", "coffee", "coffee coffee"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceIgnoreCase(tt.s, tt.old, tt.new)
			if result != tt.expected {
				t.Errorf("ReplaceIgnoreCase(%q, %q, %q) = %q, want %q",
					tt.s, tt.old, tt.new, result, tt.expected)
			}
		})
	}
}

// Benchmarks / 벤치마크

func BenchmarkContainsAny(b *testing.B) {
	s := "hello world foo bar baz"
	substrs := []string{"foo", "qux", "quux"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsAny(s, substrs)
	}
}

func BenchmarkContainsAll(b *testing.B) {
	s := "hello world foo bar baz"
	substrs := []string{"hello", "world", "foo"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsAll(s, substrs)
	}
}

func BenchmarkReplaceAll(b *testing.B) {
	s := "hello world foo bar baz"
	replacements := map[string]string{
		"hello": "hi",
		"world": "universe",
		"foo":   "bar",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReplaceAll(s, replacements)
	}
}

func BenchmarkReplaceIgnoreCase(b *testing.B) {
	s := "Hello WORLD HeLLo world"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReplaceIgnoreCase(s, "hello", "hi")
	}
}
