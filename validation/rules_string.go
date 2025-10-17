package validation

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Required validates that the value is not empty or whitespace-only.
// It trims whitespace before checking, so spaces alone will fail validation.
//
// Required는 값이 비어있지 않거나 공백만 있지 않은지 검증합니다.
// 확인 전에 공백을 제거하므로 공백만 있으면 검증 실패합니다.
//
// Behavior / 동작:
//   - Trims leading and trailing whitespace
//     앞뒤 공백 제거
//   - Fails if resulting string is empty
//     결과 문자열이 비어있으면 실패
//   - Fails if value is not a string type
//     값이 문자열 타입이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Form field validation / 폼 필드 검증
//   - Required API parameters / 필수 API 매개변수
//   - User input validation / 사용자 입력 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//
// Example / 예제:
//
//	v := validation.New("", "name")
//	v.Required()  // Fails / 실패
//
//	v := validation.New("   ", "name")
//	v.Required()  // Fails (whitespace only) / 실패 (공백만)
//
//	v := validation.New("John", "name")
//	v.Required()  // Passes / 성공
func (v *Validator) Required() *Validator {
	return validateString(v, "required", func(s string) bool {
		return len(strings.TrimSpace(s)) > 0
	}, fmt.Sprintf("%s is required / %s은(는) 필수입니다", v.fieldName, v.fieldName))
}

// MinLength validates that the string has at least n characters.
// Uses rune count for accurate Unicode character counting.
//
// MinLength는 문자열이 최소 n자 이상인지 검증합니다.
// 정확한 유니코드 문자 계수를 위해 rune 수를 사용합니다.
//
// Parameters / 매개변수:
//   - n: Minimum number of characters required
//     필요한 최소 문자 수
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Counts Unicode runes, not bytes
//     바이트가 아닌 유니코드 rune 계산
//   - Handles multi-byte characters correctly (emoji, CJK, etc.)
//     다중 바이트 문자 올바르게 처리 (이모지, 한중일 문자 등)
//   - Fails if string has fewer than n characters
//     n자보다 적으면 실패
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Password length validation / 비밀번호 길이 검증
//   - Username minimum length / 사용자명 최소 길이
//   - Comment/description length / 댓글/설명 길이
//   - International text input / 국제 텍스트 입력
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Converts to rune slice (allocates memory)
//     rune 슬라이스로 변환 (메모리 할당)
//
// Example / 예제:
//
//	v := validation.New("Hi", "message")
//	v.MinLength(5)  // Fails / 실패
//
//	v := validation.New("Hello", "message")
//	v.MinLength(5)  // Passes / 성공
//
//	v := validation.New("안녕하세요", "greeting")
//	v.MinLength(3)  // Passes (5 runes) / 성공 (5 rune)
//
//	// With multi-byte characters / 다중 바이트 문자
//	v := validation.New("👋🌍", "emoji")
//	v.MinLength(2)  // Passes (2 runes) / 성공 (2 rune)
func (v *Validator) MinLength(n int) *Validator {
	return validateString(v, "minlength", func(s string) bool {
		return len([]rune(s)) >= n
	}, fmt.Sprintf("%s must be at least %d characters / %s은(는) 최소 %d자 이상이어야 합니다", v.fieldName, n, v.fieldName, n))
}

// MaxLength validates that the string has at most n characters.
// Uses rune count for accurate Unicode character counting.
//
// MaxLength는 문자열이 최대 n자 이하인지 검증합니다.
// 정확한 유니코드 문자 계수를 위해 rune 수를 사용합니다.
//
// Parameters / 매개변수:
//   - n: Maximum number of characters allowed
//     허용되는 최대 문자 수
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Counts Unicode runes, not bytes
//     바이트가 아닌 유니코드 rune 계산
//   - Handles multi-byte characters correctly (emoji, CJK, etc.)
//     다중 바이트 문자 올바르게 처리 (이모지, 한중일 문자 등)
//   - Fails if string has more than n characters
//     n자보다 많으면 실패
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Database column length constraints / 데이터베이스 컬럼 길이 제약
//   - Username maximum length / 사용자명 최대 길이
//   - Tweet/message length limits / 트윗/메시지 길이 제한
//   - Form field restrictions / 폼 필드 제한
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Converts to rune slice (allocates memory)
//     rune 슬라이스로 변환 (메모리 할당)
//
// Example / 예제:
//
//	v := validation.New("TooLongMessage", "message")
//	v.MaxLength(5)  // Fails / 실패
//
//	v := validation.New("Short", "message")
//	v.MaxLength(10)  // Passes / 성공
//
//	v := validation.New("안녕하세요반갑습니다", "greeting")
//	v.MaxLength(5)  // Fails (10 runes) / 실패 (10 rune)
//
//	// Database VARCHAR(50) constraint / 데이터베이스 VARCHAR(50) 제약
//	v := validation.New(username, "username")
//	v.Required().MaxLength(50)
func (v *Validator) MaxLength(n int) *Validator {
	return validateString(v, "maxlength", func(s string) bool {
		return len([]rune(s)) <= n
	}, fmt.Sprintf("%s must be at most %d characters / %s은(는) 최대 %d자 이하여야 합니다", v.fieldName, n, v.fieldName, n))
}

// Length validates that the string has exactly n characters.
// Uses rune count for accurate Unicode character counting.
//
// Length는 문자열이 정확히 n자인지 검증합니다.
// 정확한 유니코드 문자 계수를 위해 rune 수를 사용합니다.
//
// Parameters / 매개변수:
//   - n: Exact number of characters required
//     필요한 정확한 문자 수
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Counts Unicode runes, not bytes
//     바이트가 아닌 유니코드 rune 계산
//   - Handles multi-byte characters correctly (emoji, CJK, etc.)
//     다중 바이트 문자 올바르게 처리 (이모지, 한중일 문자 등)
//   - Fails if string length is not exactly n
//     문자열 길이가 정확히 n이 아니면 실패
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Fixed-length codes (postal codes, product codes)
//     고정 길이 코드 (우편번호, 제품 코드)
//   - PIN codes / PIN 코드
//   - Verification codes / 인증 코드
//   - Country codes (ISO 3166) / 국가 코드 (ISO 3166)
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Converts to rune slice (allocates memory)
//     rune 슬라이스로 변환 (메모리 할당)
//
// Example / 예제:
//
//	v := validation.New("12345", "zipcode")
//	v.Length(5)  // Passes / 성공
//
//	v := validation.New("1234", "zipcode")
//	v.Length(5)  // Fails (too short) / 실패 (너무 짧음)
//
//	v := validation.New("123456", "zipcode")
//	v.Length(5)  // Fails (too long) / 실패 (너무 김)
//
//	// PIN code validation / PIN 코드 검증
//	v := validation.New(pin, "pin")
//	v.Length(4).Numeric()
//
//	// Korean phone number (11 digits) / 한국 전화번호 (11자리)
//	v := validation.New("01012345678", "phone")
//	v.Length(11).Numeric()
func (v *Validator) Length(n int) *Validator {
	return validateString(v, "length", func(s string) bool {
		return len([]rune(s)) == n
	}, fmt.Sprintf("%s must be exactly %d characters / %s은(는) 정확히 %d자여야 합니다", v.fieldName, n, v.fieldName, n))
}

// Email validates that the string is a valid email address format.
// It uses a regex pattern that covers most common email formats per RFC 5322.
//
// Email은 문자열이 유효한 이메일 주소 형식인지 검증합니다.
// RFC 5322에 따라 대부분의 일반적인 이메일 형식을 다루는 정규식 패턴을 사용합니다.
//
// Format Rules / 형식 규칙:
//   - Local part: alphanumeric, dots, underscores, percent, plus, hyphen
//     로컬 부분: 영숫자, 점, 언더스코어, 퍼센트, 플러스, 하이픈
//   - @ symbol required / @ 기호 필수
//   - Domain: alphanumeric, dots, hyphen
//     도메인: 영숫자, 점, 하이픈
//   - TLD: at least 2 letters
//     최상위 도메인: 최소 2글자
//
// Valid Examples / 유효한 예:
//   - user@example.com
//   - john.doe@company.co.uk
//   - test+tag@domain.com
//   - user_123@sub.domain.com
//
// Invalid Examples / 유효하지 않은 예:
//   - missing@domain (no TLD)
//   - @example.com (no local part)
//   - user@.com (invalid domain)
//   - user..name@example.com (consecutive dots)
//
// Behavior / 동작:
//   - Case-insensitive validation
//     대소문자 구분 없는 검증
//   - Does not validate if email actually exists
//     이메일이 실제로 존재하는지는 검증하지 않음
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Uses compiled regex (safe for concurrent use)
//     스레드 안전: 컴파일된 정규식 사용 (동시 사용 안전)
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Regex is pre-compiled for efficiency
//     효율성을 위해 정규식 사전 컴파일됨
//
// Limitations / 제한사항:
//   - Simplified regex, may not catch all edge cases
//     간소화된 정규식, 모든 엣지 케이스를 잡지 못할 수 있음
//   - International domain names (IDN) need special handling
//     국제 도메인 이름(IDN)은 특별한 처리 필요
//
// Example / 예제:
//
//	v := validation.New("user@example.com", "email")
//	v.Email()  // Passes / 성공
//
//	v := validation.New("invalid-email", "email")
//	v.Email()  // Fails / 실패
//
//	// Chaining with other validations / 다른 검증과 체이닝
//	v := validation.New(email, "email")
//	v.Required().Email().MaxLength(100)
func (v *Validator) Email() *Validator {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return validateString(v, "email", func(s string) bool {
		return emailRegex.MatchString(s)
	}, fmt.Sprintf("%s must be a valid email / %s은(는) 유효한 이메일이어야 합니다", v.fieldName, v.fieldName))
}

// URL validates that the string is a valid URL format.
// Supports HTTP and HTTPS protocols only.
//
// URL은 문자열이 유효한 URL 형식인지 검증합니다.
// HTTP 및 HTTPS 프로토콜만 지원합니다.
//
// Format Rules / 형식 규칙:
//   - Must start with http:// or https://
//     http:// 또는 https://로 시작해야 함
//   - Must have valid domain/host
//     유효한 도메인/호스트 필요
//   - May include path, query string, fragment
//     경로, 쿼리 문자열, 프래그먼트 포함 가능
//   - No whitespace allowed
//     공백 허용 안 됨
//
// Valid Examples / 유효한 예:
//   - http://example.com
//   - https://www.example.com/path
//   - https://api.example.com:8080/v1/users?id=123
//   - https://example.com/page#section
//
// Invalid Examples / 유효하지 않은 예:
//   - ftp://example.com (wrong protocol)
//   - www.example.com (missing protocol)
//   - https:// (missing domain)
//   - https://example .com (contains space)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Validates format only, not actual URL accessibility
//     형식만 검증하며 실제 URL 접근성은 검증하지 않음
//   - Case-sensitive for protocol (must be lowercase)
//     프로토콜은 대소문자 구분 (소문자여야 함)
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Website URL input / 웹사이트 URL 입력
//   - API endpoint validation / API 엔드포인트 검증
//   - Webhook URL / 웹훅 URL
//   - External resource links / 외부 리소스 링크
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Uses compiled regex (safe for concurrent use)
//     스레드 안전: 컴파일된 정규식 사용 (동시 사용 안전)
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Regex compilation happens on each call
//     각 호출마다 정규식 컴파일 발생
//
// Limitations / 제한사항:
//   - Only HTTP/HTTPS protocols supported
//     HTTP/HTTPS 프로토콜만 지원
//   - Does not validate if URL actually exists
//     URL이 실제로 존재하는지 검증하지 않음
//   - Simplified validation, may not catch all edge cases
//     간소화된 검증, 모든 엣지 케이스를 잡지 못할 수 있음
//
// Example / 예제:
//
//	v := validation.New("https://example.com", "website")
//	v.URL()  // Passes / 성공
//
//	v := validation.New("not-a-url", "website")
//	v.URL()  // Fails / 실패
//
//	// API endpoint validation / API 엔드포인트 검증
//	v := validation.New(webhookURL, "webhook")
//	v.Required().URL().MaxLength(200)
func (v *Validator) URL() *Validator {
	urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	return validateString(v, "url", func(s string) bool {
		return urlRegex.MatchString(s)
	}, fmt.Sprintf("%s must be a valid URL / %s은(는) 유효한 URL이어야 합니다", v.fieldName, v.fieldName))
}

// Alpha validates that the string contains only letters (alphabetic characters).
// Supports Unicode letters from all languages including CJK, Cyrillic, Arabic, etc.
//
// Alpha는 문자열이 문자(알파벳)만 포함하는지 검증합니다.
// 한중일, 키릴 문자, 아랍 문자 등 모든 언어의 유니코드 문자를 지원합니다.
//
// Character Rules / 문자 규칙:
//   - Only Unicode letters allowed (L category)
//     유니코드 문자만 허용 (L 카테고리)
//   - No digits, spaces, punctuation, or special characters
//     숫자, 공백, 구두점, 특수 문자 불허
//   - Supports all language scripts
//     모든 언어 스크립트 지원
//
// Valid Examples / 유효한 예:
//   - "Hello"
//   - "안녕하세요"
//   - "Привет" (Russian / 러시아어)
//   - "مرحبا" (Arabic / 아랍어)
//
// Invalid Examples / 유효하지 않은 예:
//   - "Hello123" (contains digits)
//   - "Hello World" (contains space)
//   - "Hello!" (contains punctuation)
//   - "" (empty string)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Iterates through each rune
//     각 rune을 반복
//   - Uses unicode.IsLetter() for validation
//     unicode.IsLetter()를 사용하여 검증
//   - Empty string passes validation (use Required() to prevent)
//     빈 문자열은 검증 통과 (방지하려면 Required() 사용)
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Name fields (first name, last name)
//     이름 필드 (성, 이름)
//   - Language-only input / 언어 전용 입력
//   - Text-only fields / 텍스트 전용 필드
//   - International names / 국제 이름
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Checks each rune individually
//     각 rune을 개별적으로 검사
//
// Example / 예제:
//
//	v := validation.New("John", "firstname")
//	v.Alpha()  // Passes / 성공
//
//	v := validation.New("John123", "firstname")
//	v.Alpha()  // Fails / 실패
//
//	v := validation.New("김철수", "name")
//	v.Alpha()  // Passes / 성공
//
//	// Name validation / 이름 검증
//	v := validation.New(name, "name")
//	v.Required().Alpha().MinLength(2).MaxLength(50)
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
// Supports Unicode letters and digits from all languages.
//
// Alphanumeric은 문자열이 문자와 숫자만 포함하는지 검증합니다.
// 모든 언어의 유니코드 문자와 숫자를 지원합니다.
//
// Character Rules / 문자 규칙:
//   - Unicode letters (L category) allowed
//     유니코드 문자 허용 (L 카테고리)
//   - Unicode digits (Nd category) allowed
//     유니코드 숫자 허용 (Nd 카테고리)
//   - No spaces, punctuation, or special characters
//     공백, 구두점, 특수 문자 불허
//   - Supports international characters and digits
//     국제 문자와 숫자 지원
//
// Valid Examples / 유효한 예:
//   - "abc123"
//   - "User123"
//   - "김철수123"
//   - "١٢٣abc" (Arabic digits / 아랍 숫자)
//
// Invalid Examples / 유효하지 않은 예:
//   - "abc 123" (contains space)
//   - "user_123" (contains underscore)
//   - "user-123" (contains hyphen)
//   - "user@123" (contains special char)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Iterates through each rune
//     각 rune을 반복
//   - Uses unicode.IsLetter() and unicode.IsDigit()
//     unicode.IsLetter() 및 unicode.IsDigit() 사용
//   - Empty string passes validation (use Required() to prevent)
//     빈 문자열은 검증 통과 (방지하려면 Required() 사용)
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Username validation / 사용자명 검증
//   - Product codes / 제품 코드
//   - Reference numbers / 참조 번호
//   - Identifiers without special characters / 특수 문자 없는 식별자
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Checks each rune individually
//     각 rune을 개별적으로 검사
//
// Example / 예제:
//
//	v := validation.New("User123", "username")
//	v.Alphanumeric()  // Passes / 성공
//
//	v := validation.New("User_123", "username")
//	v.Alphanumeric()  // Fails / 실패
//
//	v := validation.New("사용자123", "username")
//	v.Alphanumeric()  // Passes / 성공
//
//	// Username validation / 사용자명 검증
//	v := validation.New(username, "username")
//	v.Required().Alphanumeric().MinLength(3).MaxLength(20)
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

// Numeric validates that the string contains only numeric digits.
// Supports Unicode digits from all languages. Empty strings fail validation.
//
// Numeric은 문자열이 숫자만 포함하는지 검증합니다.
// 모든 언어의 유니코드 숫자를 지원합니다. 빈 문자열은 검증 실패합니다.
//
// Character Rules / 문자 규칙:
//   - Only Unicode digits (Nd category) allowed
//     유니코드 숫자만 허용 (Nd 카테고리)
//   - No letters, spaces, punctuation, or special characters
//     문자, 공백, 구두점, 특수 문자 불허
//   - No decimal points or negative signs
//     소수점이나 음수 기호 불허
//   - Empty string fails (unlike Alpha/Alphanumeric)
//     빈 문자열 실패 (Alpha/Alphanumeric과 다름)
//
// Valid Examples / 유효한 예:
//   - "123"
//   - "0"
//   - "١٢٣" (Arabic-Indic digits / 아랍-인도 숫자)
//   - "१२३" (Devanagari digits / 데바나가리 숫자)
//
// Invalid Examples / 유효하지 않은 예:
//   - "" (empty string)
//   - "12.34" (contains decimal)
//   - "-123" (contains minus sign)
//   - "123abc" (contains letters)
//   - "1 2 3" (contains spaces)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Iterates through each rune
//     각 rune을 반복
//   - Uses unicode.IsDigit() for validation
//     unicode.IsDigit()를 사용하여 검증
//   - Empty string explicitly fails (len check)
//     빈 문자열은 명시적으로 실패 (길이 검사)
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - ID numbers / ID 번호
//   - Numeric codes / 숫자 코드
//   - PIN validation / PIN 검증
//   - Integer-only input / 정수 전용 입력
//   - Quantity fields / 수량 필드
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Checks each rune individually
//     각 rune을 개별적으로 검사
//
// Note / 참고:
//   - For decimal numbers, use a different validator
//     소수를 위해서는 다른 검증기 사용
//   - For negative numbers, use a different validator
//     음수를 위해서는 다른 검증기 사용
//   - This is for digit-only strings, not numeric values
//     이것은 숫자 값이 아닌 숫자 전용 문자열용
//
// Example / 예제:
//
//	v := validation.New("12345", "code")
//	v.Numeric()  // Passes / 성공
//
//	v := validation.New("123abc", "code")
//	v.Numeric()  // Fails / 실패
//
//	v := validation.New("", "code")
//	v.Numeric()  // Fails (empty) / 실패 (빈 문자열)
//
//	// PIN code validation / PIN 코드 검증
//	v := validation.New(pin, "pin")
//	v.Numeric().Length(6)
//
//	// Product code / 제품 코드
//	v := validation.New(code, "product_code")
//	v.Required().Numeric().MinLength(8).MaxLength(12)
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
// Case-sensitive comparison.
//
// StartsWith는 문자열이 주어진 접두사로 시작하는지 검증합니다.
// 대소문자를 구분합니다.
//
// Parameters / 매개변수:
//   - prefix: The string that must appear at the beginning
//     시작 부분에 나타나야 하는 문자열
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Case-sensitive comparison
//     대소문자 구분 비교
//   - Uses strings.HasPrefix() internally
//     내부적으로 strings.HasPrefix() 사용
//   - Empty prefix always passes
//     빈 접두사는 항상 통과
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - URL scheme validation (http://, https://)
//     URL 스킴 검증 (http://, https://)
//   - File path validation (/, ./)
//     파일 경로 검증 (/, ./)
//   - Code/ID prefix requirements (USER_, PROD_)
//     코드/ID 접두사 요구사항 (USER_, PROD_)
//   - Command validation (/, !)
//     명령어 검증 (/, !)
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is prefix length
//     시간 복잡도: O(n) (n은 접두사 길이)
//   - Optimal performance for prefix checking
//     접두사 검사에 최적의 성능
//
// Example / 예제:
//
//	v := validation.New("https://example.com", "url")
//	v.StartsWith("https://")  // Passes / 성공
//
//	v := validation.New("http://example.com", "url")
//	v.StartsWith("https://")  // Fails / 실패
//
//	v := validation.New("USER_12345", "user_id")
//	v.StartsWith("USER_")  // Passes / 성공
//
//	// Command prefix validation / 명령어 접두사 검증
//	v := validation.New(command, "command")
//	v.StartsWith("/").MinLength(2)
//
//	// Case sensitivity / 대소문자 구분
//	v := validation.New("Hello", "greeting")
//	v.StartsWith("hello")  // Fails (case mismatch) / 실패 (대소문자 불일치)
func (v *Validator) StartsWith(prefix string) *Validator {
	return validateString(v, "startswith", func(s string) bool {
		return strings.HasPrefix(s, prefix)
	}, fmt.Sprintf("%s must start with '%s' / %s은(는) '%s'로 시작해야 합니다", v.fieldName, prefix, v.fieldName, prefix))
}

// EndsWith validates that the string ends with the given suffix.
// Case-sensitive comparison.
//
// EndsWith는 문자열이 주어진 접미사로 끝나는지 검증합니다.
// 대소문자를 구분합니다.
//
// Parameters / 매개변수:
//   - suffix: The string that must appear at the end
//     끝 부분에 나타나야 하는 문자열
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Case-sensitive comparison
//     대소문자 구분 비교
//   - Uses strings.HasSuffix() internally
//     내부적으로 strings.HasSuffix() 사용
//   - Empty suffix always passes
//     빈 접미사는 항상 통과
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - File extension validation (.jpg, .pdf, .txt)
//     파일 확장자 검증 (.jpg, .pdf, .txt)
//   - Domain validation (.com, .org, .net)
//     도메인 검증 (.com, .org, .net)
//   - Email domain restriction (@company.com)
//     이메일 도메인 제한 (@company.com)
//   - Code/ID suffix requirements (_TEST, _PROD)
//     코드/ID 접미사 요구사항 (_TEST, _PROD)
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is suffix length
//     시간 복잡도: O(n) (n은 접미사 길이)
//   - Optimal performance for suffix checking
//     접미사 검사에 최적의 성능
//
// Example / 예제:
//
//	v := validation.New("document.pdf", "filename")
//	v.EndsWith(".pdf")  // Passes / 성공
//
//	v := validation.New("document.txt", "filename")
//	v.EndsWith(".pdf")  // Fails / 실패
//
//	v := validation.New("user@company.com", "email")
//	v.EndsWith("@company.com")  // Passes / 성공
//
//	// File extension validation / 파일 확장자 검증
//	v := validation.New(filename, "upload")
//	v.Required().EndsWith(".jpg")
//
//	// Multiple allowed extensions / 여러 허용 확장자
//	filename := "image.png"
//	if !strings.HasSuffix(filename, ".jpg") && !strings.HasSuffix(filename, ".png") {
//	    // validation fails / 검증 실패
//	}
//
//	// Case sensitivity / 대소문자 구분
//	v := validation.New("FILE.PDF", "filename")
//	v.EndsWith(".pdf")  // Fails (case mismatch) / 실패 (대소문자 불일치)
func (v *Validator) EndsWith(suffix string) *Validator {
	return validateString(v, "endswith", func(s string) bool {
		return strings.HasSuffix(s, suffix)
	}, fmt.Sprintf("%s must end with '%s' / %s은(는) '%s'로 끝나야 합니다", v.fieldName, suffix, v.fieldName, suffix))
}

// Contains validates that the string contains the given substring.
// Case-sensitive search.
//
// Contains는 문자열이 주어진 부분 문자열을 포함하는지 검증합니다.
// 대소문자를 구분합니다.
//
// Parameters / 매개변수:
//   - substring: The string that must be found within the value
//     값 내에서 찾아야 하는 문자열
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Case-sensitive search
//     대소문자 구분 검색
//   - Uses strings.Contains() internally
//     내부적으로 strings.Contains() 사용
//   - Empty substring always passes
//     빈 부분 문자열은 항상 통과
//   - Finds substring at any position
//     어느 위치에서든 부분 문자열 찾기
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Keyword filtering / 키워드 필터링
//   - Content validation / 콘텐츠 검증
//   - Required text presence / 필수 텍스트 존재
//   - Substring matching / 부분 문자열 매칭
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n*m) where n is string length, m is substring length
//     시간 복잡도: O(n*m) (n은 문자열 길이, m은 부분 문자열 길이)
//   - Uses Boyer-Moore-like algorithm internally
//     내부적으로 Boyer-Moore 유사 알고리즘 사용
//
// Example / 예제:
//
//	v := validation.New("Hello World", "message")
//	v.Contains("World")  // Passes / 성공
//
//	v := validation.New("Hello World", "message")
//	v.Contains("world")  // Fails (case mismatch) / 실패 (대소문자 불일치)
//
//	v := validation.New("user@example.com", "email")
//	v.Contains("@")  // Passes / 성공
//
//	// Keyword validation / 키워드 검증
//	v := validation.New(description, "description")
//	v.Required().Contains("important").MinLength(10)
//
//	// Multiple keywords (need separate validators) / 여러 키워드 (별도 검증기 필요)
//	v1 := validation.New(text, "content").Contains("keyword1")
//	v2 := validation.New(text, "content").Contains("keyword2")
func (v *Validator) Contains(substring string) *Validator {
	return validateString(v, "contains", func(s string) bool {
		return strings.Contains(s, substring)
	}, fmt.Sprintf("%s must contain '%s' / %s은(는) '%s'를 포함해야 합니다", v.fieldName, substring, v.fieldName, substring))
}

// Regex validates that the string matches the given regular expression pattern.
// If the pattern is invalid, adds an error and returns without validation.
//
// Regex는 문자열이 주어진 정규식 패턴과 일치하는지 검증합니다.
// 패턴이 유효하지 않으면 오류를 추가하고 검증 없이 반환합니다.
//
// Parameters / 매개변수:
//   - pattern: Regular expression pattern (Go regex syntax)
//     정규식 패턴 (Go 정규식 문법)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Compiles regex pattern on each call
//     각 호출마다 정규식 패턴 컴파일
//   - Returns error if pattern is invalid
//     패턴이 유효하지 않으면 오류 반환
//   - Uses Go's regexp package (RE2 syntax)
//     Go의 regexp 패키지 사용 (RE2 문법)
//   - Full string match not required (use ^$ for full match)
//     전체 문자열 매칭 불필요 (전체 매칭은 ^$ 사용)
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Custom format validation / 사용자 정의 형식 검증
//   - Complex pattern matching / 복잡한 패턴 매칭
//   - Business-specific rules / 비즈니스 특정 규칙
//   - Advanced string validation / 고급 문자열 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Regex compilation is safe
//     스레드 안전: 정규식 컴파일 안전
//   - Compiled regex is safe for concurrent use
//     컴파일된 정규식은 동시 사용 안전
//
// Performance / 성능:
//   - Time complexity: O(n) for typical patterns
//     시간 복잡도: 일반 패턴의 경우 O(n)
//   - Regex compilation overhead on each call
//     각 호출마다 정규식 컴파일 오버헤드
//   - Consider pre-compiled regex for frequent use
//     자주 사용하는 경우 사전 컴파일된 정규식 고려
//
// Regex Syntax / 정규식 문법:
//   - Go uses RE2 syntax (no backreferences)
//     Go는 RE2 문법 사용 (역참조 없음)
//   - Common patterns: ^, $, ., *, +, ?, [], (), |
//     일반 패턴: ^, $, ., *, +, ?, [], (), |
//   - Character classes: \d, \w, \s
//     문자 클래스: \d, \w, \s
//
// Example / 예제:
//
//	// Korean phone number / 한국 전화번호
//	v := validation.New("010-1234-5678", "phone")
//	v.Regex(`^010-\d{4}-\d{4}$`)  // Passes / 성공
//
//	// Alphanumeric with hyphens / 하이픈이 있는 영숫자
//	v := validation.New("ABC-123", "code")
//	v.Regex(`^[A-Z]+-\d+$`)  // Passes / 성공
//
//	// Invalid pattern / 유효하지 않은 패턴
//	v := validation.New("test", "value")
//	v.Regex(`[invalid(`)  // Error added / 오류 추가
//
//	// Partial match (no anchors) / 부분 매칭 (앵커 없음)
//	v := validation.New("abc123def", "value")
//	v.Regex(`\d+`)  // Passes (contains digits) / 성공 (숫자 포함)
//
//	// Full string match / 전체 문자열 매칭
//	v := validation.New("abc123", "value")
//	v.Regex(`^[a-z]+$`)  // Fails (has digits) / 실패 (숫자 있음)
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

// UUID validates that the string is a valid UUID (Universally Unique Identifier).
// Supports standard UUID format (8-4-4-4-12 hexadecimal digits).
//
// UUID는 문자열이 유효한 UUID(범용 고유 식별자)인지 검증합니다.
// 표준 UUID 형식 (8-4-4-4-12 16진수)을 지원합니다.
//
// Format Rules / 형식 규칙:
//   - Five groups of hexadecimal digits
//     5개 그룹의 16진수
//   - Separated by hyphens: 8-4-4-4-12
//     하이픈으로 구분: 8-4-4-4-12
//   - Total 36 characters (32 hex + 4 hyphens)
//     총 36자 (32 16진수 + 4 하이픈)
//   - Case-insensitive (converted to lowercase)
//     대소문자 구분 없음 (소문자로 변환)
//
// Valid Examples / 유효한 예:
//   - "550e8400-e29b-41d4-a716-446655440000"
//   - "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
//   - "550E8400-E29B-41D4-A716-446655440000" (uppercase)
//
// Invalid Examples / 유효하지 않은 예:
//   - "550e8400e29b41d4a716446655440000" (no hyphens)
//   - "550e8400-e29b-41d4-a716" (incomplete)
//   - "550e8400-e29b-41d4-a716-44665544000g" (invalid hex)
//   - "550e8400-e29b-41d4-a716-4466554400000" (wrong length)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Converts to lowercase before validation
//     검증 전 소문자로 변환
//   - Uses regex for format validation
//     형식 검증에 정규식 사용
//   - Does not validate UUID version or variant
//     UUID 버전이나 변형은 검증하지 않음
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Database primary keys / 데이터베이스 기본 키
//   - Unique identifiers / 고유 식별자
//   - Session IDs / 세션 ID
//   - Resource identifiers in APIs / API의 리소스 식별자
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Uses compiled regex (safe for concurrent use)
//     스레드 안전: 컴파일된 정규식 사용 (동시 사용 안전)
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - String lowercase conversion overhead
//     문자열 소문자 변환 오버헤드
//
// Limitations / 제한사항:
//   - Does not distinguish UUID versions (v1, v4, etc.)
//     UUID 버전 구분하지 않음 (v1, v4 등)
//   - Does not validate timestamp in v1 UUIDs
//     v1 UUID의 타임스탬프 검증하지 않음
//   - Format validation only, not uniqueness
//     형식 검증만, 고유성은 검증하지 않음
//
// Example / 예제:
//
//	v := validation.New("550e8400-e29b-41d4-a716-446655440000", "id")
//	v.UUID()  // Passes / 성공
//
//	v := validation.New("not-a-uuid", "id")
//	v.UUID()  // Fails / 실패
//
//	v := validation.New("550E8400-E29B-41D4-A716-446655440000", "id")
//	v.UUID()  // Passes (case insensitive) / 성공 (대소문자 구분 없음)
//
//	// API resource ID / API 리소스 ID
//	v := validation.New(resourceID, "resource_id")
//	v.Required().UUID()
func (v *Validator) UUID() *Validator {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return validateString(v, "uuid", func(s string) bool {
		return uuidRegex.MatchString(strings.ToLower(s))
	}, fmt.Sprintf("%s must be a valid UUID / %s은(는) 유효한 UUID여야 합니다", v.fieldName, v.fieldName))
}

// JSON validates that the string is valid JSON format.
// Can validate JSON objects, arrays, strings, numbers, booleans, and null.
//
// JSON은 문자열이 유효한 JSON 형식인지 검증합니다.
// JSON 객체, 배열, 문자열, 숫자, 불린, null을 검증할 수 있습니다.
//
// Valid JSON Types / 유효한 JSON 타입:
//   - Objects: {"key": "value"}
//     객체: {"key": "value"}
//   - Arrays: [1, 2, 3]
//     배열: [1, 2, 3]
//   - Strings: "text"
//     문자열: "text"
//   - Numbers: 123, 12.34
//     숫자: 123, 12.34
//   - Booleans: true, false
//     불린: true, false
//   - Null: null
//     null: null
//
// Valid Examples / 유효한 예:
//   - `{"name": "John", "age": 30}`
//   - `[1, 2, 3, 4, 5]`
//   - `"simple string"`
//   - `123.45`
//   - `true`
//   - `null`
//
// Invalid Examples / 유효하지 않은 예:
//   - `{invalid json}` (invalid syntax)
//   - `{'key': 'value'}` (single quotes)
//   - `{name: "John"}` (unquoted key)
//   - “ (empty string)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses json.Unmarshal() for validation
//     json.Unmarshal()을 사용하여 검증
//   - Validates JSON syntax only, not schema
//     JSON 문법만 검증하며 스키마는 검증하지 않음
//   - Accepts any valid JSON value (not just objects/arrays)
//     모든 유효한 JSON 값 허용 (객체/배열만이 아님)
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - API request body validation / API 요청 본문 검증
//   - Configuration file validation / 설정 파일 검증
//   - JSON payload validation / JSON 페이로드 검증
//   - Webhook data validation / 웹훅 데이터 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is JSON string length
//     시간 복잡도: O(n) (n은 JSON 문자열 길이)
//   - Full JSON parsing overhead
//     전체 JSON 파싱 오버헤드
//   - Memory allocation for unmarshaling
//     언마샬링을 위한 메모리 할당
//
// Limitations / 제한사항:
//   - Does not validate JSON schema
//     JSON 스키마 검증하지 않음
//   - Does not validate specific structure
//     특정 구조 검증하지 않음
//   - Accepts any valid JSON type
//     모든 유효한 JSON 타입 허용
//
// Example / 예제:
//
//	v := validation.New(`{"name": "John"}`, "data")
//	v.JSON()  // Passes / 성공
//
//	v := validation.New(`{invalid}`, "data")
//	v.JSON()  // Fails / 실패
//
//	v := validation.New(`[1, 2, 3]`, "array")
//	v.JSON()  // Passes / 성공
//
//	v := validation.New(`"simple string"`, "text")
//	v.JSON()  // Passes (valid JSON string) / 성공 (유효한 JSON 문자열)
//
//	// API request body / API 요청 본문
//	v := validation.New(requestBody, "body")
//	v.Required().JSON()
func (v *Validator) JSON() *Validator {
	return validateString(v, "json", func(s string) bool {
		var js interface{}
		return json.Unmarshal([]byte(s), &js) == nil
	}, fmt.Sprintf("%s must be valid JSON / %s은(는) 유효한 JSON이어야 합니다", v.fieldName, v.fieldName))
}

// Base64 validates that the string is valid Base64 encoding.
// Uses standard Base64 encoding (RFC 4648).
//
// Base64는 문자열이 유효한 Base64 인코딩인지 검증합니다.
// 표준 Base64 인코딩 (RFC 4648)을 사용합니다.
//
// Format Rules / 형식 규칙:
//   - Characters: A-Z, a-z, 0-9, +, /
//     문자: A-Z, a-z, 0-9, +, /
//   - Padding: = character for alignment
//     패딩: 정렬을 위한 = 문자
//   - Length must be multiple of 4
//     길이는 4의 배수여야 함
//   - Whitespace not allowed
//     공백 허용 안 됨
//
// Valid Examples / 유효한 예:
//   - "SGVsbG8=" (encodes "Hello")
//   - "SGVsbG8gV29ybGQ=" (encodes "Hello World")
//   - "YWJjMTIz" (encodes "abc123")
//   - "MTIzNDU2Nzg5MA==" (with padding)
//
// Invalid Examples / 유효하지 않은 예:
//   - "Hello!" (not Base64)
//   - "SGVsbG8" (missing padding)
//   - "SGVs bG8=" (contains space)
//   - "SGVsbG8==" (incorrect padding)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses base64.StdEncoding.DecodeString()
//     base64.StdEncoding.DecodeString() 사용
//   - Attempts to decode the string
//     문자열 디코딩 시도
//   - Fails if decoding returns error
//     디코딩이 오류 반환하면 실패
//   - Does not validate decoded content
//     디코딩된 내용은 검증하지 않음
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - File upload (base64 encoded files)
//     파일 업로드 (base64 인코딩 파일)
//   - Image data validation / 이미지 데이터 검증
//   - Encoded credentials / 인코딩된 자격 증명
//   - Binary data transmission / 바이너리 데이터 전송
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Full decode operation overhead
//     전체 디코드 작업 오버헤드
//   - Memory allocation for decoded bytes
//     디코딩된 바이트를 위한 메모리 할당
//
// Encoding Standards / 인코딩 표준:
//   - Standard Base64 (RFC 4648): +, /
//     표준 Base64 (RFC 4648): +, /
//   - URL-safe Base64 not supported: -, _
//     URL 안전 Base64 미지원: -, _
//   - Base64URL requires different validator
//     Base64URL은 다른 검증기 필요
//
// Example / 예제:
//
//	v := validation.New("SGVsbG8gV29ybGQ=", "encoded")
//	v.Base64()  // Passes / 성공
//
//	v := validation.New("Not Base64!", "encoded")
//	v.Base64()  // Fails / 실패
//
//	v := validation.New("SGVsbG8", "encoded")
//	v.Base64()  // Fails (missing padding) / 실패 (패딩 누락)
//
//	// Image upload validation / 이미지 업로드 검증
//	v := validation.New(imageData, "image")
//	v.Required().Base64().MaxLength(1048576) // 1MB limit
//
//	// Decode after validation / 검증 후 디코딩
//	if v.IsValid() {
//	    decoded, _ := base64.StdEncoding.DecodeString(imageData)
//	    // Process decoded data / 디코딩된 데이터 처리
//	}
func (v *Validator) Base64() *Validator {
	return validateString(v, "base64", func(s string) bool {
		_, err := base64.StdEncoding.DecodeString(s)
		return err == nil
	}, fmt.Sprintf("%s must be valid Base64 / %s은(는) 유효한 Base64여야 합니다", v.fieldName, v.fieldName))
}

// Lowercase validates that all letters in the string are lowercase.
// Non-letter characters are ignored.
//
// Lowercase는 문자열의 모든 문자가 소문자인지 검증합니다.
// 문자가 아닌 문자는 무시됩니다.
//
// Character Rules / 문자 규칙:
//   - All letters must be lowercase
//     모든 문자는 소문자여야 함
//   - Numbers and symbols are ignored
//     숫자와 기호는 무시됨
//   - Empty string passes validation
//     빈 문자열은 검증 통과
//   - Unicode lowercase letters supported
//     유니코드 소문자 지원
//
// Valid Examples / 유효한 예:
//   - "hello"
//   - "hello123"
//   - "hello-world"
//   - "안녕하세요" (no case distinction)
//   - "123" (no letters)
//
// Invalid Examples / 유효하지 않은 예:
//   - "Hello" (has uppercase H)
//   - "hEllo" (has uppercase E)
//   - "HELLO" (all uppercase)
//   - "HeLLo123" (mixed case)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Compares string with its lowercase version
//     문자열을 소문자 버전과 비교
//   - Uses strings.ToLower() for comparison
//     strings.ToLower()를 사용하여 비교
//   - Passes if string equals its lowercase form
//     문자열이 소문자 형태와 같으면 통과
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Username format enforcement / 사용자명 형식 강제
//   - Email local part validation / 이메일 로컬 부분 검증
//   - Lowercase-only fields / 소문자 전용 필드
//   - Normalized identifiers / 정규화된 식별자
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - String allocation for lowercase conversion
//     소문자 변환을 위한 문자열 할당
//
// Example / 예제:
//
//	v := validation.New("hello", "username")
//	v.Lowercase()  // Passes / 성공
//
//	v := validation.New("Hello", "username")
//	v.Lowercase()  // Fails / 실패
//
//	v := validation.New("hello123", "username")
//	v.Lowercase()  // Passes (numbers ignored) / 성공 (숫자 무시)
//
//	v := validation.New("hello-world", "slug")
//	v.Lowercase()  // Passes (hyphen ignored) / 성공 (하이픈 무시)
//
//	// Username validation / 사용자명 검증
//	v := validation.New(username, "username")
//	v.Required().Lowercase().Alphanumeric().MinLength(3)
func (v *Validator) Lowercase() *Validator {
	return validateString(v, "lowercase", func(s string) bool {
		return s == strings.ToLower(s)
	}, fmt.Sprintf("%s must be lowercase / %s은(는) 소문자여야 합니다", v.fieldName, v.fieldName))
}

// Uppercase validates that all letters in the string are uppercase.
// Non-letter characters are ignored.
//
// Uppercase는 문자열의 모든 문자가 대문자인지 검증합니다.
// 문자가 아닌 문자는 무시됩니다.
//
// Character Rules / 문자 규칙:
//   - All letters must be uppercase
//     모든 문자는 대문자여야 함
//   - Numbers and symbols are ignored
//     숫자와 기호는 무시됨
//   - Empty string passes validation
//     빈 문자열은 검증 통과
//   - Unicode uppercase letters supported
//     유니코드 대문자 지원
//
// Valid Examples / 유효한 예:
//   - "HELLO"
//   - "HELLO123"
//   - "HELLO-WORLD"
//   - "안녕하세요" (no case distinction)
//   - "123" (no letters)
//
// Invalid Examples / 유효하지 않은 예:
//   - "Hello" (has lowercase ello)
//   - "hELLO" (has lowercase h)
//   - "hello" (all lowercase)
//   - "HeLLo123" (mixed case)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Compares string with its uppercase version
//     문자열을 대문자 버전과 비교
//   - Uses strings.ToUpper() for comparison
//     strings.ToUpper()를 사용하여 비교
//   - Passes if string equals its uppercase form
//     문자열이 대문자 형태와 같으면 통과
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Country codes (ISO 3166)
//     국가 코드 (ISO 3166)
//   - Currency codes (ISO 4217)
//     통화 코드 (ISO 4217)
//   - Language codes (ISO 639)
//     언어 코드 (ISO 639)
//   - Constant identifiers / 상수 식별자
//   - Acronyms and abbreviations / 약어 및 약자
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - String allocation for uppercase conversion
//     대문자 변환을 위한 문자열 할당
//
// Example / 예제:
//
//	v := validation.New("HELLO", "code")
//	v.Uppercase()  // Passes / 성공
//
//	v := validation.New("Hello", "code")
//	v.Uppercase()  // Fails / 실패
//
//	v := validation.New("HELLO123", "code")
//	v.Uppercase()  // Passes (numbers ignored) / 성공 (숫자 무시)
//
//	v := validation.New("USA", "country_code")
//	v.Uppercase()  // Passes / 성공
//
//	// Country code validation (ISO 3166) / 국가 코드 검증 (ISO 3166)
//	v := validation.New(countryCode, "country_code")
//	v.Required().Uppercase().Length(2).Alpha()
//
//	// Currency code (ISO 4217) / 통화 코드 (ISO 4217)
//	v := validation.New(currencyCode, "currency")
//	v.Required().Uppercase().Length(3).Alpha()
func (v *Validator) Uppercase() *Validator {
	return validateString(v, "uppercase", func(s string) bool {
		return s == strings.ToUpper(s)
	}, fmt.Sprintf("%s must be uppercase / %s은(는) 대문자여야 합니다", v.fieldName, v.fieldName))
}

// Phone validates that the string is a valid phone number format.
// Supports international formats with various separators.
//
// Phone은 문자열이 유효한 전화번호 형식인지 검증합니다.
// 다양한 구분자를 사용하는 국제 형식을 지원합니다.
//
// Format Rules / 형식 규칙:
//   - Minimum 10 digits after removing separators
//     구분자 제거 후 최소 10자리
//   - Optional country code prefix (+)
//     선택적 국가 코드 접두사 (+)
//   - Allowed separators: space, hyphen, parentheses, dot
//     허용되는 구분자: 공백, 하이픈, 괄호, 점
//   - 1-4 digit groups separated by optional separators
//     선택적 구분자로 구분된 1-4자리 그룹
//
// Valid Examples / 유효한 예:
//   - "010-1234-5678" (Korean format)
//   - "+82-10-1234-5678" (with country code)
//   - "(02) 1234-5678" (with area code)
//   - "1234567890" (no separators)
//   - "+1-555-123-4567" (US format)
//   - "02.1234.5678" (dot separators)
//
// Invalid Examples / 유효하지 않은 예:
//   - "123456789" (< 10 digits)
//   - "abc-defg-hijk" (non-numeric)
//   - "123 456" (< 10 digits after cleaning)
//   - "" (empty)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Validates format using regex
//     정규식을 사용하여 형식 검증
//   - Removes separators for digit count check
//     자릿수 확인을 위해 구분자 제거
//   - Requires minimum 10 digits
//     최소 10자리 필요
//   - Does not validate if number is active
//     번호가 활성 상태인지 검증하지 않음
//   - Skips validation if value is not a string
//     값이 문자열이 아니면 검증 건너뜀
//
// Use Cases / 사용 사례:
//   - Contact information / 연락처 정보
//   - User registration / 사용자 등록
//   - Phone verification / 전화 인증
//   - International phone numbers / 국제 전화번호
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Regex compilation is safe
//     스레드 안전: 정규식 컴파일 안전
//
// Performance / 성능:
//   - Time complexity: O(n) where n is string length
//     시간 복잡도: O(n) (n은 문자열 길이)
//   - Multiple string replacements for cleaning
//     정리를 위한 여러 문자열 교체
//   - Regex compilation on each call
//     각 호출마다 정규식 컴파일
//
// Limitations / 제한사항:
//   - Simplified validation, may not match all formats
//     간소화된 검증, 모든 형식과 일치하지 않을 수 있음
//   - Does not validate against country-specific rules
//     국가별 규칙에 대해 검증하지 않음
//   - Does not verify if number actually exists
//     번호가 실제로 존재하는지 검증하지 않음
//
// Example / 예제:
//
//	v := validation.New("010-1234-5678", "phone")
//	v.Phone()  // Passes / 성공
//
//	v := validation.New("+82-10-1234-5678", "phone")
//	v.Phone()  // Passes / 성공
//
//	v := validation.New("12345", "phone")
//	v.Phone()  // Fails (< 10 digits) / 실패 (10자리 미만)
//
//	// Korean mobile / 한국 휴대폰
//	v := validation.New(phone, "mobile")
//	v.Required().Phone().StartsWith("010")
//
//	// International phone / 국제 전화
//	v := validation.New(phone, "phone")
//	v.Required().Phone().StartsWith("+")
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
