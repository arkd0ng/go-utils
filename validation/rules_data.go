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
// ASCII는 값이 ASCII 문자(0-127)만 포함하는지 검증합니다.
//
// ASCII characters include all printable and control characters in the range 0-127.
// ASCII 문자는 0-127 범위의 모든 인쇄 가능 및 제어 문자를 포함합니다.
//
// Example / 예시:
//
//	text := "Hello World 123"
//	v := validation.New(text, "ascii_text")
//	v.ASCII()
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
// Printable은 값이 인쇄 가능한 ASCII 문자(32-126)만 포함하는지 검증합니다.
//
// Printable ASCII characters include letters, numbers, symbols, and space, but not control characters.
// 인쇄 가능한 ASCII 문자는 문자, 숫자, 기호, 공백을 포함하지만 제어 문자는 포함하지 않습니다.
//
// Example / 예시:
//
//	text := "Hello World! 123"
//	v := validation.New(text, "display_text")
//	v.Printable()
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
// Whitespace는 값이 공백 문자만 포함하는지 검증합니다.
//
// This validator checks for space, tab, newline, and other Unicode whitespace characters.
// 이 검증기는 공백, 탭, 줄바꿈 및 기타 유니코드 공백 문자를 확인합니다.
//
// Example / 예시:
//
//	text := "   \t\n  "
//	v := validation.New(text, "whitespace_field")
//	v.Whitespace()
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
// AlphaSpace는 값이 알파벳 문자와 공백만 포함하는지 검증합니다.
//
// This validator checks for letters (a-z, A-Z) and space characters only.
// No numbers or special characters are allowed.
// 이 검증기는 문자(a-z, A-Z)와 공백 문자만 확인합니다.
// 숫자나 특수 문자는 허용되지 않습니다.
//
// Example / 예시:
//
//	name := "John Doe"
//	v := validation.New(name, "full_name")
//	v.AlphaSpace()
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
