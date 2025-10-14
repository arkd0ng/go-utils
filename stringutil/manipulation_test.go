package stringutil

import "testing"

func TestRepeat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		count    int
		expected string
	}{
		{"basic repeat", "abc", 3, "abcabcabc"},
		{"single repeat", "hello", 1, "hello"},
		{"zero repeat", "test", 0, ""},
		{"negative count", "test", -1, ""},
		{"unicode repeat", "안녕", 2, "안녕안녕"},
		{"emoji repeat", "🎉", 3, "🎉🎉🎉"},
		{"empty string", "", 5, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Repeat(tt.input, tt.count)
			if result != tt.expected {
				t.Errorf("Repeat(%q, %d) = %q, want %q", tt.input, tt.count, result, tt.expected)
			}
		})
	}
}

func TestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		start    int
		end      int
		expected string
	}{
		{"basic substring", "hello world", 0, 5, "hello"},
		{"middle substring", "hello world", 6, 11, "world"},
		{"unicode substring", "안녕하세요", 0, 2, "안녕"},
		{"out of bounds", "hello", 0, 100, "hello"},
		{"negative start", "hello", -5, 3, "hel"},
		{"swapped indices", "hello", 5, 2, "llo"},
		{"empty result", "hello", 2, 2, ""},
		{"emoji substring", "👋🌍🎉", 1, 3, "🌍🎉"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Substring(tt.input, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("Substring(%q, %d, %d) = %q, want %q", tt.input, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		n        int
		expected string
	}{
		{"basic left", "hello world", 5, "hello"},
		{"unicode left", "안녕하세요", 2, "안녕"},
		{"n greater than length", "hello", 10, "hello"},
		{"zero n", "hello", 0, ""},
		{"negative n", "hello", -1, ""},
		{"emoji left", "👋🌍🎉", 2, "👋🌍"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Left(tt.input, tt.n)
			if result != tt.expected {
				t.Errorf("Left(%q, %d) = %q, want %q", tt.input, tt.n, result, tt.expected)
			}
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		n        int
		expected string
	}{
		{"basic right", "hello world", 5, "world"},
		{"unicode right", "안녕하세요", 2, "세요"},
		{"n greater than length", "hello", 10, "hello"},
		{"zero n", "hello", 0, ""},
		{"negative n", "hello", -1, ""},
		{"emoji right", "👋🌍🎉", 2, "🌍🎉"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Right(tt.input, tt.n)
			if result != tt.expected {
				t.Errorf("Right(%q, %d) = %q, want %q", tt.input, tt.n, result, tt.expected)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{"Hello World", 8, "Hello Wo..."},
		{"Hello", 10, "Hello"},
		{"안녕하세요", 3, "안녕하..."},
	}

	for _, tt := range tests {
		result := Truncate(tt.input, tt.length)
		if result != tt.expected {
			t.Errorf("Truncate(%q, %d) = %q; want %q", tt.input, tt.length, result, tt.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"안녕", "녕안"},
		{"", ""},
	}

	for _, tt := range tests {
		result := Reverse(tt.input)
		if result != tt.expected {
			t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  hello   world  ", "hello world"},
		{"\t\nhello\t\nworld", "hello world"},
		{"", ""},
	}

	for _, tt := range tests {
		result := Clean(tt.input)
		if result != tt.expected {
			t.Errorf("Clean(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
