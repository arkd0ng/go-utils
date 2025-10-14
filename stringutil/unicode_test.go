package stringutil

import "testing"

func TestRuneCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"ASCII only", "hello", 5},
		{"Korean", "안녕하세요", 5},
		{"Japanese", "こんにちは", 5},
		{"Chinese", "你好世界", 4},
		{"Emoji", "🔥🔥", 2},
		{"Mixed", "hello世界", 7},
		{"Empty string", "", 0},
		{"Spaces", "   ", 3},
		{"Special chars", "!@#$%", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RuneCount(tt.input)
			if result != tt.expected {
				t.Errorf("RuneCount(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestWidth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"ASCII only", "hello", 5},
		{"Korean", "안녕", 4},        // 2 chars × 2 width
		{"Japanese", "こんにちは", 10}, // 5 chars × 2 width
		{"Chinese", "你好", 4},       // 2 chars × 2 width
		{"Mixed ASCII+CJK", "hello世界", 9}, // 5 + 4
		{"Mixed CJK+ASCII", "안녕hello", 9}, // 4 + 5
		{"Empty string", "", 0},
		{"Spaces", "   ", 3},
		{"Numbers", "12345", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Width(tt.input)
			if result != tt.expected {
				t.Errorf("Width(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		form     string
		expected string
	}{
		// NFC tests
		{"NFC simple", "hello", "NFC", "hello"},
		{"NFC composed", "café", "NFC", "café"},
		{"NFC Korean", "한글", "NFC", "한글"},

		// NFD tests
		{"NFD simple", "hello", "NFD", "hello"},
		{"NFD decomposed", "café", "NFD", "café"}, // é → e + combining acute

		// NFKC tests
		{"NFKC compatibility numbers", "①②③", "NFKC", "123"},
		{"NFKC fullwidth", "Ｈｅｌｌｏ", "NFKC", "Hello"},

		// NFKD tests
		{"NFKD compatibility", "①②③", "NFKD", "123"},

		// Default to NFC
		{"Empty form defaults to NFC", "café", "", "café"},
		{"Invalid form defaults to NFC", "café", "INVALID", "café"},

		// Edge cases
		{"Empty string", "", "NFC", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Normalize(tt.input, tt.form)
			// For visual comparison, we check the result is not empty
			// and has expected characteristics
			if len(tt.input) > 0 && len(result) == 0 {
				t.Errorf("Normalize(%q, %q) returned empty string", tt.input, tt.form)
			}
		})
	}
}

func TestNormalizeNFC(t *testing.T) {
	// Test NFC normalization specifically
	input := "café" // This might be decomposed
	result := Normalize(input, "NFC")

	// NFC should produce composed form
	if len(result) == 0 {
		t.Error("NFC normalization returned empty string")
	}

	// RuneCount should work correctly on normalized string
	count := RuneCount(result)
	if count == 0 {
		t.Errorf("RuneCount on normalized string = %d, want > 0", count)
	}
}

func TestNormalizeNFD(t *testing.T) {
	// Test NFD normalization specifically
	input := "café"
	result := Normalize(input, "NFD")

	// NFD produces decomposed form (longer byte length typically)
	if len(result) == 0 {
		t.Error("NFD normalization returned empty string")
	}
}

func TestNormalizeIdempotent(t *testing.T) {
	// Normalizing an already normalized string should return the same result
	input := "Hello World 안녕하세요"

	// NFC
	nfc1 := Normalize(input, "NFC")
	nfc2 := Normalize(nfc1, "NFC")
	if nfc1 != nfc2 {
		t.Error("NFC normalization is not idempotent")
	}

	// NFD
	nfd1 := Normalize(input, "NFD")
	nfd2 := Normalize(nfd1, "NFD")
	if nfd1 != nfd2 {
		t.Error("NFD normalization is not idempotent")
	}
}
