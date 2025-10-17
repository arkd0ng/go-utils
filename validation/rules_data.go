package validation

import (
	"fmt"
	"unicode"
)

// ============================================================================
// DATA FORMAT VALIDATORS
// ============================================================================
//
// This file provides data format validation functions for:
// - ASCII (ASCII-only characters)
// - Printable (printable ASCII characters)
// - Whitespace (whitespace-only strings)
// - AlphaSpace (letters and spaces only)
//
// 이 파일은 다음을 위한 데이터 형식 검증 함수를 제공합니다:
// - ASCII (ASCII 문자만)
// - Printable (인쇄 가능한 ASCII 문자)
// - Whitespace (공백 문자만)
// - AlphaSpace (문자와 공백만)
//
// ============================================================================

// ASCII validates that the value contains only ASCII characters (0-127).
// Checks that all characters are in the standard ASCII range.
//
// ASCII는 값이 ASCII 문자(0-127)만 포함하는지 검증합니다.
// 모든 문자가 표준 ASCII 범위에 있는지 확인합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Checks ASCII range: 0-127
//     ASCII 범위 확인: 0-127
//   - Includes control characters (0-31, 127)
//     제어 문자 포함 (0-31, 127)
//   - Includes printable characters (32-126)
//     인쇄 가능한 문자 포함 (32-126)
//   - Fails on any non-ASCII character
//     ASCII가 아닌 문자가 있으면 실패
//
// Use Cases / 사용 사례:
//   - Legacy system compatibility / 레거시 시스템 호환성
//   - File format restrictions / 파일 형식 제한
//   - Protocol compliance / 프로토콜 준수
//   - Data encoding validation / 데이터 인코딩 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single pass character check
//     단일 패스 문자 확인
//
// ASCII characters include all printable and control characters in the range 0-127.
// ASCII 문자는 0-127 범위의 모든 인쇄 가능 및 제어 문자를 포함합니다.
//
// Example / 예시:
//
//	// Valid ASCII / 유효한 ASCII
//	text := "Hello World 123"
//	v := validation.New(text, "ascii_text")
//	v.ASCII()  // Passes
//
//	// With newline / 줄바꿈 포함
//	v = validation.New("Line1\nLine2", "text")
//	v.ASCII()  // Passes (control char allowed)
//
//	// Invalid - contains Unicode / 무효 - 유니코드 포함
//	v = validation.New("Hello 世界", "text")
//	v.ASCII()  // Fails (非ASCII)
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - All characters must be in ASCII range (0-127) / 모든 문자가 ASCII 범위(0-127)
func (v *Validator) ASCII() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("ascii", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if all characters are ASCII (0-127)
	for _, r := range str {
		if r > 127 {
			v.addError("ascii", fmt.Sprintf("%s must contain only ASCII characters / %s은(는) ASCII 문자만 포함해야 합니다", v.fieldName, v.fieldName))
			return v
		}
	}

	return v
}

// Printable validates that the value contains only printable ASCII characters (32-126).
// Excludes control characters, allowing only visible characters.
//
// Printable은 값이 인쇄 가능한 ASCII 문자(32-126)만 포함하는지 검증합니다.
// 제어 문자를 제외하고 표시 가능한 문자만 허용합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Printable range: 32-126
//     인쇄 가능 범위: 32-126
//   - Includes: letters, digits, punctuation, space
//     포함: 문자, 숫자, 구두점, 공백
//   - Excludes: control characters (0-31, 127)
//     제외: 제어 문자 (0-31, 127)
//   - Excludes: non-ASCII (>127)
//     제외: ASCII 이외 (>127)
//
// Use Cases / 사용 사례:
//   - Display text validation / 표시 텍스트 검증
//   - User input sanitization / 사용자 입력 정제
//   - Form field validation / 양식 필드 검증
//   - Terminal-safe output / 터미널 안전 출력
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single pass character check
//     단일 패스 문자 확인
//
// Printable ASCII characters include letters, numbers, symbols, and space, but not control characters.
// 인쇄 가능한 ASCII 문자는 문자, 숫자, 기호, 공백을 포함하지만 제어 문자는 포함하지 않습니다.
//
// Example / 예시:
//
//	// Valid printable / 유효한 인쇄 가능 문자
//	text := "Hello World! 123"
//	v := validation.New(text, "display_text")
//	v.Printable()  // Passes
//
//	// With symbols / 기호 포함
//	v = validation.New("Test@#$%123", "text")
//	v.Printable()  // Passes
//
//	// Invalid - contains newline / 무효 - 줄바꿈 포함
//	v = validation.New("Line1\nLine2", "text")
//	v.Printable()  // Fails (control char)
//
//	// Invalid - contains tab / 무효 - 탭 포함
//	v = validation.New("Col1\tCol2", "text")
//	v.Printable()  // Fails (control char)
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - All characters must be printable ASCII (32-126) / 모든 문자가 인쇄 가능한 ASCII (32-126)
//   - No control characters (0-31, 127) / 제어 문자 없음 (0-31, 127)
func (v *Validator) Printable() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("printable", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if all characters are printable ASCII (32-126)
	for _, r := range str {
		if r < 32 || r > 126 {
			v.addError("printable", fmt.Sprintf("%s must contain only printable characters / %s은(는) 인쇄 가능한 문자만 포함해야 합니다", v.fieldName, v.fieldName))
			return v
		}
	}

	return v
}

// Whitespace validates that the value contains only whitespace characters.
// Uses unicode.IsSpace to check for all Unicode whitespace characters.
//
// Whitespace는 값이 공백 문자만 포함하는지 검증합니다.
// unicode.IsSpace를 사용하여 모든 유니코드 공백 문자를 확인합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Must not be empty string
//     빈 문자열이면 안 됨
//   - All characters must be whitespace
//     모든 문자가 공백이어야 함
//   - Includes: space, tab, newline, carriage return
//     포함: 공백, 탭, 줄바꿈, 캐리지 리턴
//   - Includes Unicode whitespace characters
//     유니코드 공백 문자 포함
//
// Use Cases / 사용 사례:
//   - Blank line detection / 빈 줄 감지
//   - Formatting validation / 형식 검증
//   - Padding verification / 패딩 확인
//   - Template processing / 템플릿 처리
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single pass Unicode check
//     단일 패스 유니코드 확인
//
// This validator checks for space, tab, newline, and other Unicode whitespace characters.
// 이 검증기는 공백, 탭, 줄바꿈 및 기타 유니코드 공백 문자를 확인합니다.
//
// Example / 예시:
//
//	// Valid whitespace / 유효한 공백
//	text := "   \t\n  "
//	v := validation.New(text, "whitespace_field")
//	v.Whitespace()  // Passes
//
//	// Only spaces / 공백만
//	v = validation.New("     ", "padding")
//	v.Whitespace()  // Passes
//
//	// Only tabs / 탭만
//	v = validation.New("\t\t\t", "indent")
//	v.Whitespace()  // Passes
//
//	// Invalid - contains text / 무효 - 텍스트 포함
//	v = validation.New("  hello  ", "text")
//	v.Whitespace()  // Fails
//
//	// Invalid - empty string / 무효 - 빈 문자열
//	v = validation.New("", "empty")
//	v.Whitespace()  // Fails
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - All characters must be whitespace / 모든 문자가 공백이어야 함
//   - String must not be empty / 문자열이 비어있지 않아야 함
func (v *Validator) Whitespace() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("whitespace", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Must not be empty
	if len(str) == 0 {
		v.addError("whitespace", fmt.Sprintf("%s must not be empty / %s은(는) 비어있을 수 없습니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if all characters are whitespace
	for _, r := range str {
		if !unicode.IsSpace(r) {
			v.addError("whitespace", fmt.Sprintf("%s must contain only whitespace / %s은(는) 공백 문자만 포함해야 합니다", v.fieldName, v.fieldName))
			return v
		}
	}

	return v
}

// AlphaSpace validates that the value contains only alphabetic characters and spaces.
// Uses unicode.IsLetter for comprehensive letter detection including Unicode characters.
//
// AlphaSpace는 값이 알파벳 문자와 공백만 포함하는지 검증합니다.
// 유니코드 문자를 포함한 포괄적인 문자 감지를 위해 unicode.IsLetter를 사용합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Allows Unicode letters (all languages)
//     유니코드 문자 허용 (모든 언어)
//   - Allows space character
//     공백 문자 허용
//   - Rejects numbers
//     숫자 거부
//   - Rejects punctuation and symbols
//     구두점 및 기호 거부
//
// Use Cases / 사용 사례:
//   - Full name validation / 전체 이름 검증
//   - Person name fields / 개인 이름 필드
//   - Text-only input / 텍스트 전용 입력
//   - Multi-language name support / 다국어 이름 지원
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single pass Unicode check
//     단일 패스 유니코드 확인
//
// This validator checks for letters (a-z, A-Z) and space characters only.
// No numbers or special characters are allowed.
// 이 검증기는 문자(a-z, A-Z)와 공백 문자만 확인합니다.
// 숫자나 특수 문자는 허용되지 않습니다.
//
// Example / 예시:
//
//	// Valid name / 유효한 이름
//	name := "John Doe"
//	v := validation.New(name, "full_name")
//	v.AlphaSpace()  // Passes
//
//	// Single word / 단일 단어
//	v = validation.New("Alice", "first_name")
//	v.AlphaSpace()  // Passes
//
//	// Unicode letters / 유니코드 문자
//	v = validation.New("김철수", "name")
//	v.AlphaSpace()  // Passes
//
//	// Invalid - contains number / 무효 - 숫자 포함
//	v = validation.New("John123", "name")
//	v.AlphaSpace()  // Fails
//
//	// Invalid - contains punctuation / 무효 - 구두점 포함
//	v = validation.New("O'Brien", "name")
//	v.AlphaSpace()  // Fails
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Only letters and spaces allowed / 문자와 공백만 허용
//   - No numbers or special characters / 숫자나 특수 문자 없음
func (v *Validator) AlphaSpace() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("alpha_space", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if all characters are letters or spaces
	for _, r := range str {
		if !unicode.IsLetter(r) && r != ' ' {
			v.addError("alpha_space", fmt.Sprintf("%s must contain only letters and spaces / %s은(는) 문자와 공백만 포함해야 합니다", v.fieldName, v.fieldName))
			return v
		}
	}

	return v
}
