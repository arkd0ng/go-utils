package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// CreditCard checks if the value is a valid credit card number (passes Luhn algorithm).
// CreditCard는 값이 유효한 신용카드 번호인지 확인합니다 (Luhn 알고리즘 통과).
func (v *Validator) CreditCard() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("credit_card", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Remove spaces and hyphens
	cleaned := strings.ReplaceAll(strings.ReplaceAll(str, " ", ""), "-", "")

	// Check if it contains only digits
	if !regexp.MustCompile(`^\d+$`).MatchString(cleaned) {
		v.addError("credit_card", fmt.Sprintf("%s must contain only digits / %s은(는) 숫자만 포함해야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check length (credit cards are typically 13-19 digits)
	if len(cleaned) < 13 || len(cleaned) > 19 {
		v.addError("credit_card", fmt.Sprintf("%s must be 13-19 digits long / %s은(는) 13-19자리여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Validate using Luhn algorithm
	if !luhnCheck(cleaned) {
		v.addError("credit_card", fmt.Sprintf("%s must be a valid credit card number / %s은(는) 유효한 신용카드 번호여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// CreditCardType checks if the value is a valid credit card number of the specified type.
// CreditCardType는 값이 지정된 타입의 유효한 신용카드 번호인지 확인합니다.
func (v *Validator) CreditCardType(cardType string) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("credit_card_type", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Remove spaces and hyphens
	cleaned := strings.ReplaceAll(strings.ReplaceAll(str, " ", ""), "-", "")

	// Define card type patterns
	patterns := map[string]*regexp.Regexp{
		"visa":       regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`),
		"mastercard": regexp.MustCompile(`^5[1-5][0-9]{14}$`),
		"amex":       regexp.MustCompile(`^3[47][0-9]{13}$`),
		"discover":   regexp.MustCompile(`^6(?:011|5[0-9]{2})[0-9]{12}$`),
		"jcb":        regexp.MustCompile(`^(?:2131|1800)\d{12}$|^35\d{14}$`), // 2131/1800 + 12 digits = 16, 35 + 14 digits = 16
		"dinersclub": regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`),
		"unionpay":   regexp.MustCompile(`^(62[0-9]{14,17})$`),
	}

	pattern, exists := patterns[strings.ToLower(cardType)]
	if !exists {
		v.addError("credit_card_type", fmt.Sprintf("unknown credit card type: %s / 알 수 없는 신용카드 타입: %s", cardType, cardType))
		return v
	}

	if !pattern.MatchString(cleaned) {
		v.addError("credit_card_type", fmt.Sprintf("%s must be a valid %s card number / %s은(는) 유효한 %s 카드 번호여야 합니다", v.fieldName, cardType, v.fieldName, cardType))
		return v
	}

	// Also validate using Luhn algorithm
	if !luhnCheck(cleaned) {
		v.addError("credit_card_type", fmt.Sprintf("%s must be a valid credit card number / %s은(는) 유효한 신용카드 번호여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// Luhn checks if the value passes the Luhn algorithm (checksum validation).
// Luhn은 값이 Luhn 알고리즘을 통과하는지 확인합니다 (체크섬 검증).
func (v *Validator) Luhn() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("luhn", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Remove spaces and hyphens
	cleaned := strings.ReplaceAll(strings.ReplaceAll(str, " ", ""), "-", "")

	// Check if it contains only digits
	if !regexp.MustCompile(`^\d+$`).MatchString(cleaned) {
		v.addError("luhn", fmt.Sprintf("%s must contain only digits / %s은(는) 숫자만 포함해야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if !luhnCheck(cleaned) {
		v.addError("luhn", fmt.Sprintf("%s must pass Luhn algorithm check / %s은(는) Luhn 알고리즘 검사를 통과해야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// luhnCheck performs the Luhn algorithm check on a string of digits.
// luhnCheck는 숫자 문자열에 대해 Luhn 알고리즘 검사를 수행합니다.
func luhnCheck(number string) bool {
	var sum int
	parity := len(number) % 2

	for i, digit := range number {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}

		if i%2 == parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}

		sum += d
	}

	return sum%10 == 0
}
