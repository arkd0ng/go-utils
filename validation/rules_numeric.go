package validation

import "fmt"

// Min validates that the numeric value is at least n.
// Min은 숫자 값이 최소 n 이상인지 검증합니다.
func (v *Validator) Min(n float64) *Validator {
	return validateNumeric(v, "min", func(val float64) bool {
		return val >= n
	}, fmt.Sprintf("%s must be at least %v / %s은(는) 최소 %v 이상이어야 합니다", v.fieldName, n, v.fieldName, n))
}

// Max validates that the numeric value is at most n.
// Max는 숫자 값이 최대 n 이하인지 검증합니다.
func (v *Validator) Max(n float64) *Validator {
	return validateNumeric(v, "max", func(val float64) bool {
		return val <= n
	}, fmt.Sprintf("%s must be at most %v / %s은(는) 최대 %v 이하여야 합니다", v.fieldName, n, v.fieldName, n))
}

// Between validates that the numeric value is between min and max (inclusive).
// Between은 숫자 값이 min과 max 사이인지 검증합니다 (포함).
func (v *Validator) Between(min, max float64) *Validator {
	return validateNumeric(v, "between", func(val float64) bool {
		return val >= min && val <= max
	}, fmt.Sprintf("%s must be between %v and %v / %s은(는) %v와 %v 사이여야 합니다", v.fieldName, min, max, v.fieldName, min, max))
}

// Positive validates that the numeric value is greater than 0.
// Positive는 숫자 값이 0보다 큰지 검증합니다.
func (v *Validator) Positive() *Validator {
	return validateNumeric(v, "positive", func(val float64) bool {
		return val > 0
	}, fmt.Sprintf("%s must be positive / %s은(는) 양수여야 합니다", v.fieldName, v.fieldName))
}

// Negative validates that the numeric value is less than 0.
// Negative는 숫자 값이 0보다 작은지 검증합니다.
func (v *Validator) Negative() *Validator {
	return validateNumeric(v, "negative", func(val float64) bool {
		return val < 0
	}, fmt.Sprintf("%s must be negative / %s은(는) 음수여야 합니다", v.fieldName, v.fieldName))
}

// Zero validates that the numeric value is equal to 0.
// Zero는 숫자 값이 0과 같은지 검증합니다.
func (v *Validator) Zero() *Validator {
	return validateNumeric(v, "zero", func(val float64) bool {
		return val == 0
	}, fmt.Sprintf("%s must be zero / %s은(는) 0이어야 합니다", v.fieldName, v.fieldName))
}

// NonZero validates that the numeric value is not equal to 0.
// NonZero는 숫자 값이 0이 아닌지 검증합니다.
func (v *Validator) NonZero() *Validator {
	return validateNumeric(v, "nonzero", func(val float64) bool {
		return val != 0
	}, fmt.Sprintf("%s must not be zero / %s은(는) 0이 아니어야 합니다", v.fieldName, v.fieldName))
}

// Even validates that the numeric value is an even number.
// Even은 숫자 값이 짝수인지 검증합니다.
func (v *Validator) Even() *Validator {
	return validateNumeric(v, "even", func(val float64) bool {
		return int(val)%2 == 0
	}, fmt.Sprintf("%s must be an even number / %s은(는) 짝수여야 합니다", v.fieldName, v.fieldName))
}

// Odd validates that the numeric value is an odd number.
// Odd는 숫자 값이 홀수인지 검증합니다.
func (v *Validator) Odd() *Validator {
	return validateNumeric(v, "odd", func(val float64) bool {
		return int(val)%2 != 0
	}, fmt.Sprintf("%s must be an odd number / %s은(는) 홀수여야 합니다", v.fieldName, v.fieldName))
}

// MultipleOf validates that the numeric value is a multiple of n.
// MultipleOf는 숫자 값이 n의 배수인지 검증합니다.
func (v *Validator) MultipleOf(n float64) *Validator {
	return validateNumeric(v, "multipleof", func(val float64) bool {
		if n == 0 {
			return false
		}
		return int(val)%int(n) == 0
	}, fmt.Sprintf("%s must be a multiple of %v / %s은(는) %v의 배수여야 합니다", v.fieldName, n, v.fieldName, n))
}
