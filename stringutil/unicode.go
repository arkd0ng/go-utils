package stringutil

import (
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

// RuneCount returns the number of Unicode characters (runes) in a string.
// RuneCount는 문자열의 유니코드 문자(rune) 개수를 반환합니다.
//
// This is different from len(s) which returns the number of bytes.
// 이것은 바이트 개수를 반환하는 len(s)와 다릅니다.
//
// Example:
//
// RuneCount("hello")    // 5 / RuneCount("안녕하세요")  // 5 (not 15 bytes)
//	RuneCount("🔥🔥")      // 2 (not 8 bytes)
func RuneCount(s string) int {
	return utf8.RuneCountInString(s)
}

// Width returns the display width of a string.
// Width는 문자열의 디스플레이 너비를 반환합니다.
//
// This considers East Asian Width (EAW) properties:
// 동아시아 너비(EAW) 속성을 고려합니다:
// - ASCII characters (a-z, 0-9): width 1
// - CJK characters (한글, 漢字, etc): width 2
//   - Emoji: typically width 2
//
// Example:
//
// Width("hello")      // 5 / Width("안녕")        // 4 (2 characters × 2 width each)
//	Width("hello세계")   // 9 (5 + 4)
func Width(s string) int {
	totalWidth := 0
	for _, r := range s {
		prop := width.LookupRune(r)
		switch prop.Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth:
			totalWidth += 2
		default:
			totalWidth += 1
		}
	}
	return totalWidth
}

// Normalize normalizes a Unicode string to the specified form.
// Normalize는 유니코드 문자열을 지정된 형식으로 정규화합니다.
//
// Normalization Form:
// 정규화 형식:
//   - "NFC": Canonical Decomposition followed by Canonical Composition
//   - "NFD": Canonical Decomposition
//   - "NFKC": Compatibility Decomposition followed by Canonical Composition
//   - "NFKD": Compatibility Decomposition
//
// Default is NFC if form is empty or invalid.
// form이 비어있거나 유효하지 않으면 기본값은 NFC입니다.
//
// Example:
//
//	Normalize("café", "NFC")   // "café" (composed é)
//	Normalize("café", "NFD")   // "café" (decomposed e + ́)
//	Normalize("①②③", "NFKC")  // "123" (compatibility)
func Normalize(s string, form string) string {
	var normalizer norm.Form

	switch form {
	case "NFC":
		normalizer = norm.NFC
	case "NFD":
		normalizer = norm.NFD
	case "NFKC":
		normalizer = norm.NFKC
	case "NFKD":
		normalizer = norm.NFKD
	default:
		// Default to NFC
		// 기본값은 NFC
		normalizer = norm.NFC
	}

	return normalizer.String(s)
}
