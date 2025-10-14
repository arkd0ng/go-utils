package stringutil

import (
	"fmt"
	"strings"
)

// FormatNumber formats an integer with thousand separators.
// FormatNumber는 천 단위 구분 기호로 정수를 포맷합니다.
//
// Example / 예제:
//
//	FormatNumber(1000000, ",")     // "1,000,000"
//	FormatNumber(1234567, ".")     // "1.234.567"
//	FormatNumber(1234567, " ")     // "1 234 567"
//	FormatNumber(123, ",")         // "123"
func FormatNumber(n int, separator string) string {
	// Handle negative numbers / 음수 처리
	negative := n < 0
	if negative {
		n = -n
	}

	s := fmt.Sprintf("%d", n)
	result := ""

	// Add separators from right to left / 오른쪽에서 왼쪽으로 구분 기호 추가
	for i := len(s); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}

		if result != "" {
			result = separator + result
		}
		result = s[start:i] + result
	}

	if negative {
		result = "-" + result
	}

	return result
}

// FormatBytes formats bytes as human-readable size (KB, MB, GB, etc.).
// FormatBytes는 바이트를 사람이 읽을 수 있는 크기로 포맷합니다 (KB, MB, GB 등).
//
// Example / 예제:
//
//	FormatBytes(1024)                 // "1.0 KB"
//	FormatBytes(1536)                 // "1.5 KB"
//	FormatBytes(1048576)              // "1.0 MB"
//	FormatBytes(1073741824)           // "1.0 GB"
//	FormatBytes(1099511627776)        // "1.0 TB"
func FormatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	units := []string{"KB", "MB", "GB", "TB", "PB", "EB"}
	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), units[exp])
}

// Pluralize returns the singular or plural form based on count.
// Pluralize는 count에 따라 단수형 또는 복수형을 반환합니다.
//
// Example / 예제:
//
//	Pluralize(1, "item", "items")      // "item"
//	Pluralize(5, "item", "items")      // "items"
//	Pluralize(0, "item", "items")      // "items"
//	Pluralize(1, "person", "people")   // "person"
//	Pluralize(5, "person", "people")   // "people"
func Pluralize(count int, singular, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

// FormatWithCount returns a formatted string with count and pluralized noun.
// FormatWithCount는 count와 복수형 명사를 포함한 포맷된 문자열을 반환합니다.
//
// Example / 예제:
//
//	FormatWithCount(1, "item", "items")    // "1 item"
//	FormatWithCount(5, "item", "items")    // "5 items"
//	FormatWithCount(0, "item", "items")    // "0 items"
func FormatWithCount(count int, singular, plural string) string {
	return fmt.Sprintf("%d %s", count, Pluralize(count, singular, plural))
}

// Ellipsis truncates a string and adds ellipsis in the middle.
// Ellipsis는 문자열을 자르고 중간에 ellipsis를 추가합니다.
//
// If maxLen is less than or equal to 3, just truncates without ellipsis.
// maxLen이 3 이하이면 ellipsis 없이 자릅니다.
//
// Example / 예제:
//
//	Ellipsis("verylongfilename.txt", 15)  // "verylo...me.txt"
//	Ellipsis("short.txt", 20)             // "short.txt"
//	Ellipsis("abcdefgh", 3)               // "abc"
func Ellipsis(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}

	if maxLen <= 3 {
		return string(runes[:maxLen])
	}

	// Calculate split point / 분할 지점 계산
	ellipsisLen := 3 // "..."
	leftLen := (maxLen - ellipsisLen) / 2
	rightLen := maxLen - ellipsisLen - leftLen

	return string(runes[:leftLen]) + "..." + string(runes[len(runes)-rightLen:])
}

// Mask masks a string with a character, revealing only first and last n characters.
// Mask는 문자열을 문자로 마스크하고, 처음과 마지막 n개 문자만 표시합니다.
//
// If first+last is greater than or equal to string length, returns original string.
// first+last가 문자열 길이보다 크거나 같으면 원본 문자열을 반환합니다.
//
// Example / 예제:
//
//	Mask("1234567890", 2, 2, "*")      // "12******90"
//	Mask("hello@example.com", 2, 4, "*")  // "he*****.com"
//	Mask("secret", 1, 1, "#")          // "s####t"
//	Mask("short", 2, 2, "*")           // "short"
func Mask(s string, first, last int, maskChar string) string {
	runes := []rune(s)
	length := len(runes)

	if first+last > length {
		return s
	}

	var result strings.Builder
	result.WriteString(string(runes[:first]))

	maskLen := length - first - last
	for i := 0; i < maskLen; i++ {
		result.WriteString(maskChar)
	}

	result.WriteString(string(runes[length-last:]))
	return result.String()
}

// MaskEmail masks an email address, revealing only the first character and domain.
// MaskEmail은 이메일 주소를 마스크하고, 첫 문자와 도메인만 표시합니다.
//
// Example / 예제:
//
//	MaskEmail("john.doe@example.com")  // "j******e@example.com"
//	MaskEmail("a@example.com")         // "a@example.com"
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}

	local := parts[0]
	domain := parts[1]

	if len(local) <= 2 {
		return email
	}

	maskedLocal := string(local[0]) + strings.Repeat("*", len(local)-2) + string(local[len(local)-1])
	return maskedLocal + "@" + domain
}

// MaskCreditCard masks a credit card number, revealing only the last 4 digits.
// MaskCreditCard는 신용카드 번호를 마스크하고, 마지막 4자리만 표시합니다.
//
// Example / 예제:
//
//	MaskCreditCard("1234567890123456")  // "************3456"
//	MaskCreditCard("1234-5678-9012-3456")  // "****-****-****-3456"
func MaskCreditCard(card string) string {
	runes := []rune(card)
	length := len(runes)

	if length <= 4 {
		return card
	}

	var result strings.Builder
	for i := 0; i < length-4; i++ {
		if runes[i] == '-' || runes[i] == ' ' {
			result.WriteRune(runes[i])
		} else {
			result.WriteRune('*')
		}
	}
	result.WriteString(string(runes[length-4:]))

	return result.String()
}

// AddLineNumbers adds line numbers to each line of text.
// AddLineNumbers는 텍스트의 각 줄에 줄 번호를 추가합니다.
//
// Example / 예제:
//
//	AddLineNumbers("line1\nline2\nline3")
//	// "1: line1\n2: line2\n3: line3"
func AddLineNumbers(s string) string {
	lines := strings.Split(s, "\n")
	var result strings.Builder

	for i, line := range lines {
		if i > 0 {
			result.WriteRune('\n')
		}
		result.WriteString(fmt.Sprintf("%d: %s", i+1, line))
	}

	return result.String()
}

// Indent indents each line with the specified prefix.
// Indent는 각 줄을 지정된 접두사로 들여쓰기합니다.
//
// Example / 예제:
//
//	Indent("line1\nline2", "  ")  // "  line1\n  line2"
//	Indent("line1\nline2", "\t")  // "\tline1\n\tline2"
func Indent(s string, prefix string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if line != "" || i < len(lines)-1 {
			lines[i] = prefix + line
		}
	}
	return strings.Join(lines, "\n")
}

// Dedent removes common leading whitespace from each line.
// Dedent는 각 줄에서 공통 선행 공백을 제거합니다.
//
// Example / 예제:
//
//	Dedent("  line1\n  line2")  // "line1\nline2"
//	Dedent("    line1\n  line2")  // "  line1\nline2"
func Dedent(s string) string {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return s
	}

	// Find minimum indentation / 최소 들여쓰기 찾기
	minIndent := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		indent := 0
		for _, r := range line {
			if r != ' ' && r != '\t' {
				break
			}
			indent++
		}

		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}

	if minIndent <= 0 {
		return s
	}

	// Remove common indentation / 공통 들여쓰기 제거
	for i, line := range lines {
		if len(line) >= minIndent {
			lines[i] = line[minIndent:]
		}
	}

	return strings.Join(lines, "\n")
}

// WrapText wraps text to the specified line width.
// WrapText는 텍스트를 지정된 줄 너비로 줄바꿈합니다.
//
// Example / 예제:
//
//	WrapText("The quick brown fox jumps", 10)
//	// "The quick\nbrown fox\njumps"
func WrapText(s string, width int) string {
	if width <= 0 {
		return s
	}

	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}

	var result strings.Builder
	lineLen := 0

	for i, word := range words {
		wordLen := len([]rune(word))

		if lineLen == 0 {
			// First word on line / 줄의 첫 단어
			result.WriteString(word)
			lineLen = wordLen
		} else if lineLen+1+wordLen <= width {
			// Word fits on current line / 현재 줄에 단어 맞춤
			result.WriteRune(' ')
			result.WriteString(word)
			lineLen += 1 + wordLen
		} else {
			// Start new line / 새 줄 시작
			result.WriteRune('\n')
			result.WriteString(word)
			lineLen = wordLen
		}

		_ = i // avoid unused variable warning
	}

	return result.String()
}
