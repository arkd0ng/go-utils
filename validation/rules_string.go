package validation

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Required validates that the value is not empty.
// Required는 값이 비어있지 않은지 검증합니다.
func (v *Validator) Required() *Validator {
	return validateString(v, "required", func(s string) bool {
		return len(strings.TrimSpace(s)) > 0
	}, fmt.Sprintf("%s is required / %s은(는) 필수입니다", v.fieldName, v.fieldName))
}

// MinLength validates that the string has at least n characters.
// MinLength는 문자열이 최소 n자 이상인지 검증합니다.
func (v *Validator) MinLength(n int) *Validator {
	return validateString(v, "minlength", func(s string) bool {
		return len([]rune(s)) >= n
	}, fmt.Sprintf("%s must be at least %d characters / %s은(는) 최소 %d자 이상이어야 합니다", v.fieldName, n, v.fieldName, n))
}

// MaxLength validates that the string has at most n characters.
// MaxLength는 문자열이 최대 n자 이하인지 검증합니다.
func (v *Validator) MaxLength(n int) *Validator {
	return validateString(v, "maxlength", func(s string) bool {
		return len([]rune(s)) <= n
	}, fmt.Sprintf("%s must be at most %d characters / %s은(는) 최대 %d자 이하여야 합니다", v.fieldName, n, v.fieldName, n))
}

// Length validates that the string has exactly n characters.
// Length는 문자열이 정확히 n자인지 검증합니다.
func (v *Validator) Length(n int) *Validator {
	return validateString(v, "length", func(s string) bool {
		return len([]rune(s)) == n
	}, fmt.Sprintf("%s must be exactly %d characters / %s은(는) 정확히 %d자여야 합니다", v.fieldName, n, v.fieldName, n))
}

// Email validates that the string is a valid email address.
// Email은 문자열이 유효한 이메일 주소인지 검증합니다.
func (v *Validator) Email() *Validator {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return validateString(v, "email", func(s string) bool {
		return emailRegex.MatchString(s)
	}, fmt.Sprintf("%s must be a valid email / %s은(는) 유효한 이메일이어야 합니다", v.fieldName, v.fieldName))
}

// URL validates that the string is a valid URL.
// URL은 문자열이 유효한 URL인지 검증합니다.
func (v *Validator) URL() *Validator {
	urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	return validateString(v, "url", func(s string) bool {
		return urlRegex.MatchString(s)
	}, fmt.Sprintf("%s must be a valid URL / %s은(는) 유효한 URL이어야 합니다", v.fieldName, v.fieldName))
}

// Alpha validates that the string contains only letters.
// Alpha는 문자열이 문자만 포함하는지 검증합니다.
func (v *Validator) Alpha() *Validator {
	return validateString(v, "alpha", func(s string) bool {
		for _, r := range s {
			if !unicode.IsLetter(r) {
				return false
			}
		}
		return true
	}, fmt.Sprintf("%s must contain only letters / %s은(는) 문자만 포함해야 합니다", v.fieldName, v.fieldName))
}

// Alphanumeric validates that the string contains only letters and numbers.
// Alphanumeric은 문자열이 문자와 숫자만 포함하는지 검증합니다.
func (v *Validator) Alphanumeric() *Validator {
	return validateString(v, "alphanumeric", func(s string) bool {
		for _, r := range s {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				return false
			}
		}
		return true
	}, fmt.Sprintf("%s must contain only letters and numbers / %s은(는) 문자와 숫자만 포함해야 합니다", v.fieldName, v.fieldName))
}

// Numeric validates that the string contains only numbers.
// Numeric은 문자열이 숫자만 포함하는지 검증합니다.
func (v *Validator) Numeric() *Validator {
	return validateString(v, "numeric", func(s string) bool {
		for _, r := range s {
			if !unicode.IsDigit(r) {
				return false
			}
		}
		return len(s) > 0
	}, fmt.Sprintf("%s must contain only numbers / %s은(는) 숫자만 포함해야 합니다", v.fieldName, v.fieldName))
}

// StartsWith validates that the string starts with the given prefix.
// StartsWith는 문자열이 주어진 접두사로 시작하는지 검증합니다.
func (v *Validator) StartsWith(prefix string) *Validator {
	return validateString(v, "startswith", func(s string) bool {
		return strings.HasPrefix(s, prefix)
	}, fmt.Sprintf("%s must start with '%s' / %s은(는) '%s'로 시작해야 합니다", v.fieldName, prefix, v.fieldName, prefix))
}

// EndsWith validates that the string ends with the given suffix.
// EndsWith는 문자열이 주어진 접미사로 끝나는지 검증합니다.
func (v *Validator) EndsWith(suffix string) *Validator {
	return validateString(v, "endswith", func(s string) bool {
		return strings.HasSuffix(s, suffix)
	}, fmt.Sprintf("%s must end with '%s' / %s은(는) '%s'로 끝나야 합니다", v.fieldName, suffix, v.fieldName, suffix))
}

// Contains validates that the string contains the given substring.
// Contains는 문자열이 주어진 부분 문자열을 포함하는지 검증합니다.
func (v *Validator) Contains(substring string) *Validator {
	return validateString(v, "contains", func(s string) bool {
		return strings.Contains(s, substring)
	}, fmt.Sprintf("%s must contain '%s' / %s은(는) '%s'를 포함해야 합니다", v.fieldName, substring, v.fieldName, substring))
}

// Regex validates that the string matches the given regular expression pattern.
// Regex는 문자열이 주어진 정규식 패턴과 일치하는지 검증합니다.
func (v *Validator) Regex(pattern string) *Validator {
	re, err := regexp.Compile(pattern)
	if err != nil {
		v.addError("regex", fmt.Sprintf("invalid regex pattern: %v", err))
		return v
	}

	return validateString(v, "regex", func(s string) bool {
		return re.MatchString(s)
	}, fmt.Sprintf("%s must match pattern '%s' / %s은(는) 패턴 '%s'와 일치해야 합니다", v.fieldName, pattern, v.fieldName, pattern))
}

// UUID validates that the string is a valid UUID.
// UUID는 문자열이 유효한 UUID인지 검증합니다.
func (v *Validator) UUID() *Validator {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return validateString(v, "uuid", func(s string) bool {
		return uuidRegex.MatchString(strings.ToLower(s))
	}, fmt.Sprintf("%s must be a valid UUID / %s은(는) 유효한 UUID여야 합니다", v.fieldName, v.fieldName))
}

// JSON validates that the string is valid JSON.
// JSON은 문자열이 유효한 JSON인지 검증합니다.
func (v *Validator) JSON() *Validator {
	return validateString(v, "json", func(s string) bool {
		var js interface{}
		return json.Unmarshal([]byte(s), &js) == nil
	}, fmt.Sprintf("%s must be valid JSON / %s은(는) 유효한 JSON이어야 합니다", v.fieldName, v.fieldName))
}

// Base64 validates that the string is valid Base64 encoding.
// Base64는 문자열이 유효한 Base64 인코딩인지 검증합니다.
func (v *Validator) Base64() *Validator {
	return validateString(v, "base64", func(s string) bool {
		_, err := base64.StdEncoding.DecodeString(s)
		return err == nil
	}, fmt.Sprintf("%s must be valid Base64 / %s은(는) 유효한 Base64여야 합니다", v.fieldName, v.fieldName))
}

// Lowercase validates that the string is all lowercase.
// Lowercase는 문자열이 모두 소문자인지 검증합니다.
func (v *Validator) Lowercase() *Validator {
	return validateString(v, "lowercase", func(s string) bool {
		return s == strings.ToLower(s)
	}, fmt.Sprintf("%s must be lowercase / %s은(는) 소문자여야 합니다", v.fieldName, v.fieldName))
}

// Uppercase validates that the string is all uppercase.
// Uppercase는 문자열이 모두 대문자인지 검증합니다.
func (v *Validator) Uppercase() *Validator {
	return validateString(v, "uppercase", func(s string) bool {
		return s == strings.ToUpper(s)
	}, fmt.Sprintf("%s must be uppercase / %s은(는) 대문자여야 합니다", v.fieldName, v.fieldName))
}

// Phone validates that the string is a valid phone number.
// Phone은 문자열이 유효한 전화번호인지 검증합니다.
func (v *Validator) Phone() *Validator {
	// Simple phone validation - can be extended
	phoneRegex := regexp.MustCompile(`^[+]?[(]?[0-9]{1,4}[)]?[-\s.]?[(]?[0-9]{1,4}[)]?[-\s.]?[0-9]{1,9}$`)
	return validateString(v, "phone", func(s string) bool {
		// Remove common separators
		cleaned := strings.ReplaceAll(s, " ", "")
		cleaned = strings.ReplaceAll(cleaned, "-", "")
		cleaned = strings.ReplaceAll(cleaned, "(", "")
		cleaned = strings.ReplaceAll(cleaned, ")", "")
		cleaned = strings.ReplaceAll(cleaned, ".", "")
		return phoneRegex.MatchString(s) && len(cleaned) >= 10
	}, fmt.Sprintf("%s must be a valid phone number / %s은(는) 유효한 전화번호여야 합니다", v.fieldName, v.fieldName))
}

// CreditCard validates basic credit card format (Luhn algorithm not implemented).
// CreditCard는 기본 신용카드 형식을 검증합니다 (Luhn 알고리즘 미구현).
func (v *Validator) CreditCard() *Validator {
	ccRegex := regexp.MustCompile(`^[0-9]{13,19}$`)
	return validateString(v, "creditcard", func(s string) bool {
		// Remove spaces and dashes
		cleaned := strings.ReplaceAll(s, " ", "")
		cleaned = strings.ReplaceAll(cleaned, "-", "")
		return ccRegex.MatchString(cleaned)
	}, fmt.Sprintf("%s must be a valid credit card number / %s은(는) 유효한 신용카드 번호여야 합니다", v.fieldName, v.fieldName))
}
