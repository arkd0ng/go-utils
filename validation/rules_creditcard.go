package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// CreditCard validates that the value is a valid credit card number using the Luhn algorithm.
// Accepts numbers with spaces or hyphens, validates length (13-19 digits), and performs checksum verification.
//
// CreditCard는 Luhn 알고리즘을 사용하여 값이 유효한 신용카드 번호인지 검증합니다.
// 공백이나 하이픈이 포함된 번호를 허용하고, 길이(13-19자리)를 검증하며, 체크섬 확인을 수행합니다.
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
//   - Removes spaces and hyphens before validation
//     검증 전에 공백과 하이픈 제거
//   - Accepts only numeric characters
//     숫자만 허용
//   - Validates length: 13-19 digits (standard card length)
//     길이 검증: 13-19자리 (표준 카드 길이)
//   - Applies Luhn algorithm checksum validation
//     Luhn 알고리즘 체크섬 검증 적용
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Payment form validation / 결제 양식 검증
//   - Credit card input validation / 신용카드 입력 검증
//   - Card number format verification / 카드 번호 형식 확인
//   - E-commerce checkout validation / 전자상거래 체크아웃 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = number length
//     시간 복잡도: O(n), n = 번호 길이
//   - Single pass Luhn algorithm
//     단일 패스 Luhn 알고리즘
//
// Example / 예제:
//
//	// Valid card numbers / 유효한 카드 번호
//	v := validation.New("4532015112830366", "card")
//	v.CreditCard()  // Passes (valid Visa)
//
//	// With formatting / 형식 포함
//	v = validation.New("4532-0151-1283-0366", "card")
//	v.CreditCard()  // Passes (hyphens removed)
//
//	v = validation.New("4532 0151 1283 0366", "card")
//	v.CreditCard()  // Passes (spaces removed)
//
//	// Invalid checksum / 무효한 체크섬
//	v = validation.New("4532015112830367", "card")
//	v.CreditCard()  // Fails (Luhn check fails)
//
//	// Invalid length / 무효한 길이
//	v = validation.New("123", "card")
//	v.CreditCard()  // Fails (too short)
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

// CreditCardType validates that the value is a valid credit card number of the specified type.
// Checks card-specific patterns and applies Luhn algorithm verification.
//
// CreditCardType는 값이 지정된 타입의 유효한 신용카드 번호인지 검증합니다.
// 카드별 패턴을 확인하고 Luhn 알고리즘 검증을 적용합니다.
//
// Parameters / 매개변수:
//   - cardType: Card type to validate (visa, mastercard, amex, discover, jcb, dinersclub, unionpay)
//     검증할 카드 타입 (visa, mastercard, amex, discover, jcb, dinersclub, unionpay)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Removes spaces and hyphens before validation
//     검증 전에 공백과 하이픈 제거
//   - Case-insensitive card type matching
//     대소문자 구분 없는 카드 타입 매칭
//   - Validates card-specific patterns:
//     카드별 패턴 검증:
//   - Visa: Starts with 4, 13 or 16 digits
//   - MasterCard: Starts with 51-55, 16 digits
//   - Amex: Starts with 34 or 37, 15 digits
//   - Discover: Starts with 6011 or 65, 16 digits
//   - JCB: Starts with 2131, 1800, or 35, 16 digits
//   - Diners Club: Starts with 300-305 or 36/38, 14 digits
//   - UnionPay: Starts with 62, 16-19 digits
//   - Applies Luhn algorithm after pattern check
//     패턴 확인 후 Luhn 알고리즘 적용
//   - Fails if card type unknown
//     카드 타입이 알 수 없으면 실패
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Card type-specific validation / 카드 타입별 검증
//   - Payment processor integration / 결제 처리 통합
//   - Card brand detection / 카드 브랜드 감지
//   - Merchant-specific card acceptance / 가맹점별 카드 수용
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = number length
//     시간 복잡도: O(n), n = 번호 길이
//   - Single regex match + Luhn check
//     단일 정규식 매칭 + Luhn 확인
//
// Example / 예제:
//
//	// Visa card validation / Visa 카드 검증
//	v := validation.New("4532015112830366", "card")
//	v.CreditCardType("visa")  // Passes
//
//	// Mastercard validation / Mastercard 검증
//	v = validation.New("5425233430109903", "card")
//	v.CreditCardType("mastercard")  // Passes
//
//	// American Express / American Express
//	v = validation.New("374245455400126", "card")
//	v.CreditCardType("amex")  // Passes
//
//	// Wrong type / 잘못된 타입
//	v = validation.New("4532015112830366", "card")
//	v.CreditCardType("mastercard")  // Fails (Visa card)
//
//	// Case-insensitive / 대소문자 구분 없음
//	v = validation.New("4532015112830366", "card")
//	v.CreditCardType("VISA")  // Passes
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

// Luhn validates that the value passes the Luhn algorithm checksum verification.
// Generic Luhn validation for any numeric string (credit cards, IMEI, etc.).
//
// Luhn은 값이 Luhn 알고리즘 체크섬 검증을 통과하는지 검증합니다.
// 모든 숫자 문자열(신용카드, IMEI 등)에 대한 일반 Luhn 검증입니다.
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
//   - Removes spaces and hyphens before validation
//     검증 전에 공백과 하이픈 제거
//   - Accepts only numeric characters
//     숫자만 허용
//   - Applies Luhn algorithm (mod 10 checksum)
//     Luhn 알고리즘 적용 (mod 10 체크섬)
//   - No length restrictions (unlike CreditCard)
//     길이 제한 없음 (CreditCard와 달리)
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Algorithm / 알고리즘:
//  1. Starting from rightmost digit, double every second digit
//     가장 오른쪽 숫자부터 시작하여 두 번째 숫자마다 두 배로
//  2. If doubling results in > 9, subtract 9
//     두 배 결과가 9보다 크면 9를 뺌
//  3. Sum all digits
//     모든 숫자 합계
//  4. Valid if sum % 10 == 0
//     합계 % 10 == 0이면 유효
//
// Use Cases / 사용 사례:
//   - Credit card validation / 신용카드 검증
//   - IMEI number validation / IMEI 번호 검증
//   - National ID validation (some countries) / 주민등록번호 검증 (일부 국가)
//   - Generic checksum validation / 일반 체크섬 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = number length
//     시간 복잡도: O(n), n = 번호 길이
//   - Single pass algorithm
//     단일 패스 알고리즘
//
// Example / 예제:
//
//	// Valid Luhn numbers / 유효한 Luhn 번호
//	v := validation.New("79927398713", "number")
//	v.Luhn()  // Passes (valid checksum)
//
//	// Credit card number / 신용카드 번호
//	v = validation.New("4532015112830366", "card")
//	v.Luhn()  // Passes
//
//	// With formatting / 형식 포함
//	v = validation.New("4532-0151-1283-0366", "card")
//	v.Luhn()  // Passes (hyphens removed)
//
//	// Invalid checksum / 무효한 체크섬
//	v = validation.New("79927398714", "number")
//	v.Luhn()  // Fails (checksum error)
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

// luhnCheck performs the Luhn algorithm checksum validation on a string of digits.
// Internal helper function that implements the Luhn mod 10 algorithm.
//
// luhnCheck는 숫자 문자열에 대해 Luhn 알고리즘 체크섬 검증을 수행합니다.
// Luhn mod 10 알고리즘을 구현하는 내부 헬퍼 함수입니다.
//
// Parameters / 매개변수:
//   - number: String of digits to validate (no spaces or hyphens)
//     검증할 숫자 문자열 (공백이나 하이픈 없음)
//
// Returns / 반환:
//   - bool: true if checksum is valid, false otherwise
//     bool: 체크섬이 유효하면 true, 그렇지 않으면 false
//
// Algorithm / 알고리즘:
//  1. Process digits from right to left
//     오른쪽에서 왼쪽으로 숫자 처리
//  2. Double every second digit based on parity
//     패리티에 따라 두 번째 숫자마다 두 배로
//  3. If doubled digit > 9, subtract 9 (equivalent to summing digits)
//     두 배한 숫자가 9보다 크면 9를 뺌 (자릿수 합계와 동일)
//  4. Sum all processed digits
//     처리된 모든 숫자 합계
//  5. Valid if sum % 10 == 0
//     합계 % 10 == 0이면 유효
//
// Behavior / 동작:
//   - Returns false if non-digit character encountered
//     숫자가 아닌 문자를 만나면 false 반환
//   - Uses parity to determine which digits to double
//     패리티를 사용하여 두 배할 숫자 결정
//   - Applies mod 10 checksum validation
//     mod 10 체크섬 검증 적용
//
// Performance / 성능:
//   - Time complexity: O(n), n = number length
//     시간 복잡도: O(n), n = 번호 길이
//   - Space complexity: O(1)
//     공간 복잡도: O(1)
//   - Single pass through digits
//     숫자를 통한 단일 패스
//
// Example / 예제:
//
//	luhnCheck("79927398713")  // true (valid)
//	luhnCheck("79927398714")  // false (invalid checksum)
//	luhnCheck("4532015112830366")  // true (valid Visa)
//	luhnCheck("abc123")  // false (contains non-digits)
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
