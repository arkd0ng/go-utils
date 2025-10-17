package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ISBN validates International Standard Book Number (ISBN-10 or ISBN-13).
// Validates both formats with checksum verification.
//
// ISBN는 국제 표준 도서 번호(ISBN-10 또는 ISBN-13)를 검증합니다.
// 체크섬 확인과 함께 두 형식을 모두 검증합니다.
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
//   - Accepts ISBN-10 (10 digits)
//     ISBN-10 허용 (10자리)
//   - Accepts ISBN-13 (13 digits)
//     ISBN-13 허용 (13자리)
//   - Allows hyphens and spaces (removed for validation)
//     하이픈 및 공백 허용 (검증 시 제거)
//   - Validates checksum for both formats
//     두 형식 모두 체크섬 검증
//
// Use Cases / 사용 사례:
//   - Book cataloging systems / 도서 목록 시스템
//   - E-commerce book stores / 전자상거래 서점
//   - Library management / 도서관 관리
//   - Publishing platforms / 출판 플랫폼
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = 10 or 13
//     시간 복잡도: O(n), n = 10 또는 13
//   - Checksum calculation required
//     체크섬 계산 필요
//
// Supports both ISBN-10 and ISBN-13 formats with or without hyphens.
// ISBN-10과 ISBN-13 형식을 모두 지원하며 하이픈 포함/미포함 가능합니다.
//
// Example / 예시:
//
//	// Valid ISBN-13 / 유효한 ISBN-13
//	v := validation.New("978-0-596-52068-7", "isbn")
//	v.ISBN()  // Passes
//
//	// Valid ISBN-10 / 유효한 ISBN-10
//	v = validation.New("0-596-52068-9", "isbn")
//	v.ISBN()  // Passes
//
//	// Without hyphens / 하이픈 없이
//	v = validation.New("9780596520687", "isbn")
//	v.ISBN()  // Passes
//
//	// Invalid checksum / 무효한 체크섬
//	v = validation.New("978-0-596-52068-0", "isbn")
//	v.ISBN()  // Fails
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Must be 10 or 13 digits (excluding hyphens/spaces) / 10 또는 13자리 (하이픈/공백 제외)
//   - Must pass checksum validation / 체크섬 검증 통과
func (v *Validator) ISBN() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("isbn", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Remove hyphens and spaces
	cleaned := strings.ReplaceAll(strings.ReplaceAll(str, "-", ""), " ", "")

	// Check if it's ISBN-10 or ISBN-13
	if len(cleaned) == 10 {
		if !isValidISBN10(cleaned) {
			v.addError("isbn", fmt.Sprintf("%s must be a valid ISBN-10 / %s은(는) 유효한 ISBN-10이어야 합니다", v.fieldName, v.fieldName))
		}
	} else if len(cleaned) == 13 {
		if !isValidISBN13(cleaned) {
			v.addError("isbn", fmt.Sprintf("%s must be a valid ISBN-13 / %s은(는) 유효한 ISBN-13이어야 합니다", v.fieldName, v.fieldName))
		}
	} else {
		v.addError("isbn", fmt.Sprintf("%s must be a valid ISBN (10 or 13 digits) / %s은(는) 유효한 ISBN이어야 합니다 (10 또는 13자리)", v.fieldName, v.fieldName))
	}

	return v
}

// ISSN validates International Standard Serial Number (ISSN-8).
// Validates 8-digit format with checksum verification.
//
// ISSN는 국제 표준 연속 간행물 번호(ISSN-8)를 검증합니다.
// 체크섬 확인과 함께 8자리 형식을 검증합니다.
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
//   - Validates 8-digit format (7 digits + checksum)
//     8자리 형식 검증 (7자리 + 체크섬)
//   - Checksum can be 0-9 or X
//     체크섬은 0-9 또는 X
//   - Allows optional hyphen after 4th digit
//     4번째 자리 뒤 선택적 하이픈 허용
//   - Validates checksum algorithm
//     체크섬 알고리즘 검증
//
// Use Cases / 사용 사례:
//   - Journal cataloging / 저널 목록 작성
//   - Magazine subscription systems / 잡지 구독 시스템
//   - Periodical databases / 정기 간행물 데이터베이스
//   - Library serial management / 도서관 연속 간행물 관리
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1), fixed length
//     시간 복잡도: O(1), 고정 길이
//   - Checksum calculation: 7 operations
//     체크섬 계산: 7회 연산
//
// ISSN format: XXXX-XXXX (8 digits with optional hyphen after 4th digit)
// ISSN 형식: XXXX-XXXX (4번째 자리 뒤에 선택적 하이픈이 있는 8자리)
//
// Example / 예시:
//
//	// Valid with hyphen / 하이픈 포함 유효
//	v := validation.New("2049-3630", "issn")
//	v.ISSN()  // Passes
//
//	// Valid without hyphen / 하이픈 없이 유효
//	v = validation.New("20493630", "issn")
//	v.ISSN()  // Passes
//
//	// With X checksum / X 체크섬 포함
//	v = validation.New("0378-595X", "issn")
//	v.ISSN()  // Passes
//
//	// Invalid checksum / 무효한 체크섬
//	v = validation.New("2049-3631", "issn")
//	v.ISSN()  // Fails
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Must be 8 digits (excluding hyphen) / 8자리 (하이픈 제외)
//   - Last digit can be 0-9 or X / 마지막 자리는 0-9 또는 X
//   - Must pass checksum validation / 체크섬 검증 통과
func (v *Validator) ISSN() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("issn", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Remove hyphen and spaces
	cleaned := strings.ReplaceAll(strings.ReplaceAll(str, "-", ""), " ", "")

	if !isValidISSN(cleaned) {
		v.addError("issn", fmt.Sprintf("%s must be a valid ISSN / %s은(는) 유효한 ISSN이어야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// EAN validates European Article Number (EAN-8 or EAN-13).
// Validates product barcodes with checksum verification.
//
// EAN는 유럽 상품 코드(EAN-8 또는 EAN-13)를 검증합니다.
// 체크섬 확인과 함께 상품 바코드를 검증합니다.
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
//   - Supports EAN-8 (8 digits)
//     EAN-8 지원 (8자리)
//   - Supports EAN-13 (13 digits)
//     EAN-13 지원 (13자리)
//   - Allows hyphens and spaces (removed for validation)
//     하이픈 및 공백 허용 (검증 시 제거)
//   - Validates checksum using EAN algorithm
//     EAN 알고리즘을 사용한 체크섬 검증
//
// Use Cases / 사용 사례:
//   - Retail product scanning / 소매 상품 스캔
//   - Inventory management / 재고 관리
//   - Point of sale systems / 판매 시점 시스템
//   - E-commerce product identification / 전자상거래 상품 식별
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = 8 or 13
//     시간 복잡도: O(n), n = 8 또는 13
//   - Checksum calculation required
//     체크섬 계산 필요
//
// Supports both EAN-8 and EAN-13 formats with checksum validation.
// 체크섬 검증과 함께 EAN-8 및 EAN-13 형식을 모두 지원합니다.
//
// Example / 예시:
//
//	// Valid EAN-13 / 유효한 EAN-13
//	v := validation.New("4006381333931", "ean")
//	v.EAN()  // Passes
//
//	// Valid EAN-8 / 유효한 EAN-8
//	v = validation.New("96385074", "ean")
//	v.EAN()  // Passes
//
//	// With spaces / 공백 포함
//	v = validation.New("4006381 333931", "ean")
//	v.EAN()  // Passes
//
//	// Invalid checksum / 무효한 체크섬
//	v = validation.New("4006381333930", "ean")
//	v.EAN()  // Fails
//
//	// Invalid length / 무효한 길이
//	v = validation.New("123456", "ean")
//	v.EAN()  // Fails (not 8 or 13 digits)
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Must be 8 or 13 digits (excluding hyphens/spaces) / 8 또는 13자리 (하이픈/공백 제외)
//   - Must pass checksum validation / 체크섬 검증 통과
func (v *Validator) EAN() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("ean", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Remove spaces and hyphens
	cleaned := strings.ReplaceAll(strings.ReplaceAll(str, "-", ""), " ", "")

	// Check if it's EAN-8 or EAN-13
	if len(cleaned) == 8 {
		if !isValidEAN8(cleaned) {
			v.addError("ean", fmt.Sprintf("%s must be a valid EAN-8 / %s은(는) 유효한 EAN-8이어야 합니다", v.fieldName, v.fieldName))
		}
	} else if len(cleaned) == 13 {
		if !isValidEAN13(cleaned) {
			v.addError("ean", fmt.Sprintf("%s must be a valid EAN-13 / %s은(는) 유효한 EAN-13이어야 합니다", v.fieldName, v.fieldName))
		}
	} else {
		v.addError("ean", fmt.Sprintf("%s must be a valid EAN (8 or 13 digits) / %s은(는) 유효한 EAN이어야 합니다 (8 또는 13자리)", v.fieldName, v.fieldName))
	}

	return v
}

// isValidISBN10 validates ISBN-10 using checksum algorithm.
// isValidISBN10은 체크섬 알고리즘을 사용하여 ISBN-10을 검증합니다.
func isValidISBN10(isbn string) bool {
	// ISBN-10 must be 10 characters (9 digits + checksum which can be 0-9 or X)
	if len(isbn) != 10 {
		return false
	}

	// Check format: 9 digits followed by digit or X
	isbnRegex := regexp.MustCompile(`^\d{9}[\dXx]$`)
	if !isbnRegex.MatchString(isbn) {
		return false
	}

	// Calculate checksum
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(isbn[i]))
		sum += digit * (10 - i)
	}

	// Last character
	lastChar := strings.ToUpper(string(isbn[9]))
	var checkDigit int
	if lastChar == "X" {
		checkDigit = 10
	} else {
		checkDigit, _ = strconv.Atoi(lastChar)
	}

	sum += checkDigit

	return sum%11 == 0
}

// isValidISBN13 validates ISBN-13 using checksum algorithm.
// isValidISBN13은 체크섬 알고리즘을 사용하여 ISBN-13을 검증합니다.
func isValidISBN13(isbn string) bool {
	// ISBN-13 must be 13 digits
	if len(isbn) != 13 {
		return false
	}

	// Check if all characters are digits
	isbnRegex := regexp.MustCompile(`^\d{13}$`)
	if !isbnRegex.MatchString(isbn) {
		return false
	}

	// Calculate checksum
	sum := 0
	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(isbn[i]))
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}

	checkDigit, _ := strconv.Atoi(string(isbn[12]))
	remainder := sum % 10
	var expectedCheckDigit int
	if remainder == 0 {
		expectedCheckDigit = 0
	} else {
		expectedCheckDigit = 10 - remainder
	}

	return checkDigit == expectedCheckDigit
}

// isValidISSN validates ISSN using checksum algorithm.
// isValidISSN은 체크섬 알고리즘을 사용하여 ISSN을 검증합니다.
func isValidISSN(issn string) bool {
	// ISSN must be 8 characters (7 digits + checksum which can be 0-9 or X)
	if len(issn) != 8 {
		return false
	}

	// Check format: 7 digits followed by digit or X
	issnRegex := regexp.MustCompile(`^\d{7}[\dXx]$`)
	if !issnRegex.MatchString(issn) {
		return false
	}

	// Calculate checksum
	sum := 0
	for i := 0; i < 7; i++ {
		digit, _ := strconv.Atoi(string(issn[i]))
		sum += digit * (8 - i)
	}

	// Last character
	lastChar := strings.ToUpper(string(issn[7]))
	var checkDigit int
	if lastChar == "X" {
		checkDigit = 10
	} else {
		checkDigit, _ = strconv.Atoi(lastChar)
	}

	sum += checkDigit

	return sum%11 == 0
}

// isValidEAN8 validates EAN-8 using checksum algorithm.
// isValidEAN8은 체크섬 알고리즘을 사용하여 EAN-8을 검증합니다.
func isValidEAN8(ean string) bool {
	// EAN-8 must be 8 digits
	if len(ean) != 8 {
		return false
	}

	// Check if all characters are digits
	eanRegex := regexp.MustCompile(`^\d{8}$`)
	if !eanRegex.MatchString(ean) {
		return false
	}

	// Calculate checksum
	sum := 0
	for i := 0; i < 7; i++ {
		digit, _ := strconv.Atoi(string(ean[i]))
		if i%2 == 0 {
			sum += digit * 3
		} else {
			sum += digit
		}
	}

	checkDigit, _ := strconv.Atoi(string(ean[7]))
	remainder := sum % 10
	var expectedCheckDigit int
	if remainder == 0 {
		expectedCheckDigit = 0
	} else {
		expectedCheckDigit = 10 - remainder
	}

	return checkDigit == expectedCheckDigit
}

// isValidEAN13 validates EAN-13 using checksum algorithm.
// isValidEAN13은 체크섬 알고리즘을 사용하여 EAN-13을 검증합니다.
func isValidEAN13(ean string) bool {
	// EAN-13 must be 13 digits
	if len(ean) != 13 {
		return false
	}

	// Check if all characters are digits
	eanRegex := regexp.MustCompile(`^\d{13}$`)
	if !eanRegex.MatchString(ean) {
		return false
	}

	// Calculate checksum
	sum := 0
	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(ean[i]))
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}

	checkDigit, _ := strconv.Atoi(string(ean[12]))
	remainder := sum % 10
	var expectedCheckDigit int
	if remainder == 0 {
		expectedCheckDigit = 0
	} else {
		expectedCheckDigit = 10 - remainder
	}

	return checkDigit == expectedCheckDigit
}
