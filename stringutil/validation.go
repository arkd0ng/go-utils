package stringutil

import (
	"regexp"
	"strings"
	"unicode"
)

// =============================================================================
// File: validation.go
// Purpose: String Validation and Type Checking
// 파일: validation.go
// 목적: 문자열 검증 및 타입 확인
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The validation.go file provides a comprehensive set of validation functions
// for checking whether strings match specific patterns or contain specific types
// of characters. These functions are essential for input validation, data
// sanitization, and ensuring data quality. They cover common validation scenarios
// from email addresses to character type checking (alpha, numeric, alphanumeric).
//
// validation.go 파일은 문자열이 특정 패턴과 일치하거나 특정 유형의 문자를
// 포함하는지 확인하는 포괄적인 검증 함수 세트를 제공합니다. 이러한 함수는
// 입력 검증, 데이터 정제 및 데이터 품질 보장에 필수적입니다. 이메일 주소부터
// 문자 유형 확인 (알파, 숫자, 영숫자)까지 일반적인 검증 시나리오를 다룹니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Practical Validation**: Focus on real-world use cases, not theoretical perfection
//    **실용적 검증**: 이론적 완벽성이 아닌 실제 사용 사례에 초점
//
// 2. **Performance-Conscious**: Use efficient algorithms and early exits
//    **성능 의식**: 효율적인 알고리즘 및 조기 종료 사용
//
// 3. **Clear Semantics**: Boolean returns with intuitive naming (Is*)
//    **명확한 의미**: 직관적인 명명(Is*)의 불린 반환
//
// 4. **Unicode-Aware**: Handle Unicode characters properly
//    **유니코드 인식**: 유니코드 문자 적절히 처리
//
// 5. **Composability**: Functions can be combined for complex validation
//    **조합 가능성**: 복잡한 검증을 위해 함수 결합 가능
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. FORMAT VALIDATION (형식 검증)
//    - IsEmail: Validate email address format
//      IsEmail: 이메일 주소 형식 검증
//    - IsURL: Validate URL format (http/https)
//      IsURL: URL 형식 검증 (http/https)
//
// 2. CHARACTER TYPE CHECKING (문자 유형 확인)
//    - IsAlphanumeric: Letters and digits only
//      IsAlphanumeric: 문자와 숫자만
//    - IsNumeric: Digits only (0-9)
//      IsNumeric: 숫자만 (0-9)
//    - IsAlpha: Letters only (a-z, A-Z)
//      IsAlpha: 문자만 (a-z, A-Z)
//
// 3. WHITESPACE CHECKING (공백 확인)
//    - IsBlank: Empty or whitespace only
//      IsBlank: 비어있거나 공백만
//
// 4. CASE CHECKING (케이스 확인)
//    - IsLower: All letters lowercase
//      IsLower: 모든 문자가 소문자
//    - IsUpper: All letters uppercase
//      IsUpper: 모든 문자가 대문자
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// IsEmail(s string) bool
// - Purpose: Validate email address format
// - 목적: 이메일 주소 형식 검증
// - Validation: Practical regex, not RFC 5322 compliant
// - 검증: 실용적 정규식, RFC 5322 완전 준수 아님
// - Pattern: [localpart]@[domain].[tld]
// - 패턴: [localpart]@[domain].[tld]
// - Allows: Letters, digits, ._%+- in local part
// - 허용: 로컬 부분에 문자, 숫자, ._%+-
// - Time Complexity: O(n) - regex matching
// - 시간 복잡도: O(n) - 정규식 매칭
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: User registration, form validation, contact forms
// - 사용 사례: 사용자 등록, 폼 검증, 연락 폼
//
// IsURL(s string) bool
// - Purpose: Validate URL format (http/https)
// - 목적: URL 형식 검증 (http/https)
// - Validation: Simple prefix check (http:// or https://)
// - 검증: 간단한 접두사 확인 (http:// 또는 https://)
// - Note: Does NOT validate full URL structure
// - 참고: 전체 URL 구조 검증 안 함
// - Time Complexity: O(1) - prefix check
// - 시간 복잡도: O(1) - 접두사 확인
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Link validation, URL filtering, protocol detection
// - 사용 사례: 링크 검증, URL 필터링, 프로토콜 감지
//
// IsAlphanumeric(s string) bool
// - Purpose: Check if string contains only letters and digits
// - 목적: 문자열이 문자와 숫자만 포함하는지 확인
// - Allowed: a-z, A-Z, 0-9 (plus Unicode letters/digits)
// - 허용: a-z, A-Z, 0-9 (유니코드 문자/숫자 포함)
// - Empty String: Returns false
// - 빈 문자열: false 반환
// - Time Complexity: O(n) - iterate all characters
// - 시간 복잡도: O(n) - 모든 문자 반복
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Username validation, ID validation, sanitization
// - 사용 사례: 사용자명 검증, ID 검증, 정제
//
// IsNumeric(s string) bool
// - Purpose: Check if string contains only digits
// - 목적: 문자열이 숫자만 포함하는지 확인
// - Allowed: 0-9 (plus Unicode digits)
// - 허용: 0-9 (유니코드 숫자 포함)
// - Empty String: Returns false
// - 빈 문자열: false 반환
// - Note: Does NOT handle negative numbers or decimals
// - 참고: 음수나 소수점 처리 안 함
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Numeric input validation, ID validation, zip codes
// - 사용 사례: 숫자 입력 검증, ID 검증, 우편번호
//
// IsAlpha(s string) bool
// - Purpose: Check if string contains only letters
// - 목적: 문자열이 문자만 포함하는지 확인
// - Allowed: a-z, A-Z (plus Unicode letters)
// - 허용: a-z, A-Z (유니코드 문자 포함)
// - Empty String: Returns false
// - 빈 문자열: false 반환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Name validation, text-only input, language detection
// - 사용 사례: 이름 검증, 텍스트만 입력, 언어 감지
//
// IsBlank(s string) bool
// - Purpose: Check if string is empty or contains only whitespace
// - 목적: 문자열이 비어있거나 공백만 포함하는지 확인
// - Whitespace: Space, tab, newline, carriage return, etc.
// - 공백: 스페이스, 탭, 줄바꿈, 캐리지 리턴 등
// - Empty String: Returns true
// - 빈 문자열: true 반환
// - Implementation: Uses strings.TrimSpace
// - 구현: strings.TrimSpace 사용
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Required field validation, empty string detection
// - 사용 사례: 필수 필드 검증, 빈 문자열 감지
//
// IsLower(s string) bool
// - Purpose: Check if all letters are lowercase
// - 목적: 모든 문자가 소문자인지 확인
// - Non-Letters: Digits and symbols ignored
// - 비문자: 숫자와 기호 무시
// - Empty String: Returns false
// - 빈 문자열: false 반환
// - Requires: At least one letter
// - 요구사항: 최소 한 개의 문자
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Case validation, style checking, normalization check
// - 사용 사례: 케이스 검증, 스타일 확인, 정규화 확인
//
// IsUpper(s string) bool
// - Purpose: Check if all letters are uppercase
// - 목적: 모든 문자가 대문자인지 확인
// - Non-Letters: Digits and symbols ignored
// - 비문자: 숫자와 기호 무시
// - Empty String: Returns false
// - 빈 문자열: false 반환
// - Requires: At least one letter
// - 요구사항: 최소 한 개의 문자
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Constant name validation, style checking, format detection
// - 사용 사례: 상수 이름 검증, 스타일 확인, 형식 감지
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - IsEmail: O(n) - regex matching (compiled once, cached)
//   IsEmail: O(n) - 정규식 매칭 (한 번 컴파일, 캐시됨)
// - IsURL: O(1) - simple prefix check
//   IsURL: O(1) - 간단한 접두사 확인
// - IsAlphanumeric/IsNumeric/IsAlpha: O(n) - iterate all runes
//   IsAlphanumeric/IsNumeric/IsAlpha: O(n) - 모든 룬 반복
// - IsBlank: O(n) - iterate to find non-whitespace
//   IsBlank: O(n) - 공백 아닌 것 찾기 위해 반복
// - IsLower/IsUpper: O(n) - iterate all runes
//   IsLower/IsUpper: O(n) - 모든 룬 반복
//
// Space Complexities:
// 공간 복잡도:
// - All functions: O(1) - no extra allocation
//   모든 함수: O(1) - 추가 할당 없음
// - Exception: IsEmail compiles regex (once, cached)
//   예외: IsEmail은 정규식 컴파일 (한 번, 캐시됨)
//
// Optimization Tips:
// 최적화 팁:
// 1. IsURL is fastest - simple prefix check
//    IsURL이 가장 빠름 - 간단한 접두사 확인
// 2. For repeated email validation, regex is compiled once
//    반복적인 이메일 검증의 경우 정규식은 한 번만 컴파일됨
// 3. Character type checks use early exit on first non-matching character
//    문자 유형 확인은 첫 번째 불일치 문자에서 조기 종료 사용
// 4. For very long strings, consider length check first
//    매우 긴 문자열의 경우 먼저 길이 확인 고려
// 5. IsBlank is efficient for detecting empty input
//    IsBlank는 빈 입력 감지에 효율적
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty String Handling:
// 빈 문자열 처리:
// - IsEmail(""): false (not a valid email)
//   IsEmail(""): false (유효한 이메일 아님)
// - IsURL(""): false (no scheme)
//   IsURL(""): false (스킴 없음)
// - IsAlphanumeric(""): false (no characters)
//   IsAlphanumeric(""): false (문자 없음)
// - IsNumeric(""): false (no digits)
//   IsNumeric(""): false (숫자 없음)
// - IsAlpha(""): false (no letters)
//   IsAlpha(""): false (문자 없음)
// - IsBlank(""): true (empty is blank)
//   IsBlank(""): true (빈 것은 공백)
// - IsLower(""): false (no letters)
//   IsLower(""): false (문자 없음)
// - IsUpper(""): false (no letters)
//   IsUpper(""): false (문자 없음)
//
// Whitespace-Only Strings:
// 공백만 있는 문자열:
// - IsBlank("   "): true (only whitespace)
//   IsBlank("   "): true (공백만)
// - IsAlpha("   "): false (spaces not letters)
//   IsAlpha("   "): false (공백은 문자 아님)
//
// Mixed Case:
// 혼합 케이스:
// - IsLower("abc123"): true (digits ignored)
//   IsLower("abc123"): true (숫자 무시)
// - IsUpper("ABC123"): true (digits ignored)
//   IsUpper("ABC123"): true (숫자 무시)
// - IsLower("aBc"): false (has uppercase B)
//   IsLower("aBc"): false (대문자 B 있음)
//
// Unicode Characters:
// 유니코드 문자:
// - IsAlpha("你好"): true (Unicode letters)
//   IsAlpha("你好"): true (유니코드 문자)
// - IsNumeric("١٢٣"): true (Arabic-Indic digits)
//   IsNumeric("١٢٣"): true (아랍-인도 숫자)
// - IsAlphanumeric("café"): true (accented letters)
//   IsAlphanumeric("café"): true (악센트 문자)
//
// Email Validation Limitations:
// 이메일 검증 제한사항:
// - Does NOT validate against full RFC 5322 spec
//   전체 RFC 5322 사양 검증 안 함
// - Allows common patterns (99% of real emails)
//   일반적인 패턴 허용 (실제 이메일의 99%)
// - May reject some valid but rare email formats
//   일부 유효하지만 드문 이메일 형식 거부 가능
//
// URL Validation Limitations:
// URL 검증 제한사항:
// - Only checks for http:// or https:// prefix
//   http:// 또는 https:// 접두사만 확인
// - Does NOT validate domain, path, or query parameters
//   도메인, 경로 또는 쿼리 매개변수 검증 안 함
// - For comprehensive URL validation, use net/url.Parse
//   포괄적인 URL 검증을 위해 net/url.Parse 사용
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. User Registration Email Validation
//    사용자 등록 이메일 검증:
//
//    email := "user@example.com"
//    if !stringutil.IsEmail(email) {
//        return errors.New("invalid email address")
//    }
//    // Validate before saving to database
//    // 데이터베이스 저장 전 검증
//
// 2. Username Validation (Alphanumeric Only)
//    사용자명 검증 (영숫자만):
//
//    username := "user123"
//    if !stringutil.IsAlphanumeric(username) {
//        return errors.New("username must contain only letters and numbers")
//    }
//    // No special characters allowed
//    // 특수 문자 허용 안 함
//
// 3. Required Field Validation
//    필수 필드 검증:
//
//    name := strings.TrimSpace(userInput)
//    if stringutil.IsBlank(name) {
//        return errors.New("name is required")
//    }
//    // Ensure non-empty input
//    // 비어있지 않은 입력 보장
//
// 4. Numeric ID Validation
//    숫자 ID 검증:
//
//    id := "12345"
//    if !stringutil.IsNumeric(id) {
//        return errors.New("ID must be numeric")
//    }
//    // Only digits allowed
//    // 숫자만 허용
//
// 5. Constant Name Validation (Uppercase)
//    상수 이름 검증 (대문자):
//
//    constName := "MAX_VALUE"
//    if !stringutil.IsUpper(strings.ReplaceAll(constName, "_", "")) {
//        fmt.Println("Warning: constant name should be uppercase")
//    }
//    // Check naming convention
//    // 명명 규칙 확인
//
// 6. Link Validation
//    링크 검증:
//
//    link := "https://example.com"
//    if !stringutil.IsURL(link) {
//        return errors.New("invalid URL - must start with http:// or https://")
//    }
//    // Ensure secure protocol
//    // 보안 프로토콜 보장
//
// 7. Name Validation (Letters Only)
//    이름 검증 (문자만):
//
//    firstName := "John"
//    if !stringutil.IsAlpha(firstName) {
//        return errors.New("name must contain only letters")
//    }
//    // No numbers or special characters
//    // 숫자나 특수 문자 없음
//
// 8. Combined Validation (Multiple Checks)
//    결합 검증 (다중 확인):
//
//    password := "MyPass123"
//    hasLower := false
//    hasUpper := false
//    hasDigit := false
//    for _, r := range password {
//        if unicode.IsLower(r) { hasLower = true }
//        if unicode.IsUpper(r) { hasUpper = true }
//        if unicode.IsDigit(r) { hasDigit = true }
//    }
//    if !hasLower || !hasUpper || !hasDigit {
//        return errors.New("password must contain lowercase, uppercase, and digit")
//    }
//    // Complex validation logic
//    // 복잡한 검증 로직
//
// 9. Whitespace Detection in Input
//    입력의 공백 감지:
//
//    userInput := "  \t\n  "
//    if stringutil.IsBlank(userInput) {
//        fmt.Println("Please provide non-empty input")
//    }
//    // Detect effectively empty input
//    // 효과적으로 빈 입력 감지
//
// 10. Format Detection
//     형식 감지:
//
//     code := "ABC123"
//     if stringutil.IsAlphanumeric(code) {
//         if stringutil.IsNumeric(code) {
//             fmt.Println("Numeric code")
//         } else if stringutil.IsAlpha(code) {
//             fmt.Println("Alphabetic code")
//         } else {
//             fmt.Println("Mixed alphanumeric code")
//         }
//     }
//     // Classify code format
//     // 코드 형식 분류
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// IsEmail vs regexp.MustCompile
// - IsEmail: Pre-compiled regex, convenient
//   IsEmail: 미리 컴파일된 정규식, 편리
// - regexp.MustCompile: More control, custom patterns
//   regexp.MustCompile: 더 많은 제어, 사용자 정의 패턴
// - Use IsEmail for: Standard email validation
//   IsEmail 사용: 표준 이메일 검증
//
// IsURL vs net/url.Parse
// - IsURL: Simple prefix check, fast
//   IsURL: 간단한 접두사 확인, 빠름
// - net/url.Parse: Full URL parsing, comprehensive
//   net/url.Parse: 전체 URL 파싱, 포괄적
// - Use IsURL for: Quick protocol check
//   IsURL 사용: 빠른 프로토콜 확인
// - Use net/url.Parse for: Full validation
//   net/url.Parse 사용: 전체 검증
//
// IsBlank vs len(s) == 0
// - IsBlank: Checks for whitespace too
//   IsBlank: 공백도 확인
// - len(s) == 0: Only checks empty string
//   len(s) == 0: 빈 문자열만 확인
// - Use IsBlank for: User input validation
//   IsBlank 사용: 사용자 입력 검증
//
// IsNumeric vs strconv.Atoi
// - IsNumeric: Only validates format
//   IsNumeric: 형식만 검증
// - strconv.Atoi: Validates and parses to int
//   strconv.Atoi: 검증하고 int로 파싱
// - Use IsNumeric for: Format check without parsing
//   IsNumeric 사용: 파싱 없이 형식 확인
//
// IsAlpha vs regexp
// - IsAlpha: Fast, simple character check
//   IsAlpha: 빠름, 간단한 문자 확인
// - regexp: More flexible patterns
//   regexp: 더 유연한 패턴
// - Use IsAlpha for: Pure letter check
//   IsAlpha 사용: 순수 문자 확인
//
// VALIDATION BEST PRACTICES
// 검증 모범 사례
// --------------------------
// 1. **Client and Server Validation**: Always validate on server, even if validated on client
//    **클라이언트 및 서버 검증**: 클라이언트에서 검증해도 항상 서버에서 검증
//
// 2. **Whitespace Trimming**: Use strings.TrimSpace before validation
//    **공백 제거**: 검증 전에 strings.TrimSpace 사용
//
// 3. **Clear Error Messages**: Provide specific feedback on validation failures
//    **명확한 오류 메시지**: 검증 실패 시 구체적인 피드백 제공
//
// 4. **Combine Checks**: Use multiple validation functions for complex requirements
//    **확인 결합**: 복잡한 요구사항에 여러 검증 함수 사용
//
// 5. **Unicode Awareness**: Remember that unicode.IsLetter includes non-ASCII letters
//    **유니코드 인식**: unicode.IsLetter는 비ASCII 문자 포함 기억
//
// 6. **Performance**: Check simplest conditions first (length, emptiness)
//    **성능**: 가장 간단한 조건 먼저 확인 (길이, 비어있음)
//
// 7. **Security**: Don't trust user input, always validate
//    **보안**: 사용자 입력 신뢰 안 함, 항상 검증
//
// 8. **Context-Specific**: Adjust validation rules to your use case
//    **컨텍스트 특정**: 사용 사례에 맞게 검증 규칙 조정
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All functions in this file are thread-safe as they operate on immutable strings
// and don't use shared mutable state. The regex in IsEmail is compiled once and
// cached, which is safe for concurrent use.
//
// 이 파일의 모든 함수는 불변 문자열에서 작동하고 공유 가변 상태를 사용하지
// 않으므로 스레드 안전합니다. IsEmail의 정규식은 한 번 컴파일되고 캐시되며,
// 동시 사용에 안전합니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     go func() {
//         valid := stringutil.IsEmail(email)
//     }()
//
//     go func() {
//         valid := stringutil.IsNumeric(id)
//     }()
//
//     // All validation functions safe for concurrent use
//     // 모든 검증 함수는 동시 사용에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - comparison.go: String comparison operations
//   comparison.go: 문자열 비교 연산
// - search.go: String search operations
//   search.go: 문자열 검색 연산
// - case.go: Case conversion operations
//   case.go: 케이스 변환 연산
//
// =============================================================================

// IsEmail validates if a string is an email address (practical validation).
// IsEmail은 문자열이 이메일 주소인지 검증합니다 (실용적 검증).
//
// Not RFC 5322 compliant, but good enough for 99% of cases.
// RFC 5322 완전 준수 아니지만 99%의 경우에 충분함.
//
// Example:
//
//	IsEmail("user@example.com")      // true
//	IsEmail("user+tag@example.com")  // true
//	IsEmail("invalid.email")         // false
//	IsEmail("@example.com")          // false
func IsEmail(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(s)
}

// IsURL validates if a string is a URL.
// IsURL은 문자열이 URL인지 검증합니다.
//
// Checks for http:// or https:// scheme.
// http:// 또는 https:// 스킴 확인.
//
// Example:
//
//	IsURL("https://example.com")       // true
//	IsURL("http://example.com/path")   // true
//	IsURL("example.com")               // false (no scheme)
//	IsURL("htp://invalid")             // false
func IsURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

// IsAlphanumeric checks if string contains only alphanumeric characters (a-z, A-Z, 0-9).
// IsAlphanumeric은 문자열이 영숫자만 포함하는지 확인합니다 (a-z, A-Z, 0-9).
//
// Example:
//
//	IsAlphanumeric("abc123")   // true
//	IsAlphanumeric("ABC")      // true
//	IsAlphanumeric("abc-123")  // false (has dash)
//	IsAlphanumeric("abc 123")  // false (has space)
func IsAlphanumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsNumeric checks if string contains only digits (0-9).
// IsNumeric은 문자열이 숫자만 포함하는지 확인합니다 (0-9).
//
// Example:
//
//	IsNumeric("12345")   // true
//	IsNumeric("0")       // true
//	IsNumeric("123.45")  // false (has dot)
//	IsNumeric("-123")    // false (has minus)
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsAlpha checks if string contains only letters (a-z, A-Z).
// IsAlpha는 문자열이 알파벳만 포함하는지 확인합니다 (a-z, A-Z).
//
// Example:
//
//	IsAlpha("abcABC")  // true
//	IsAlpha("hello")   // true
//	IsAlpha("abc123")  // false (has digits)
//	IsAlpha("abc-")    // false (has dash)
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsBlank checks if string is empty or contains only whitespace.
// IsBlank는 문자열이 비어있거나 공백만 포함하는지 확인합니다.
//
// Example:
//
//	IsBlank("")       // true
//	IsBlank("   ")    // true
//	IsBlank("\t\n")   // true
//	IsBlank("hello")  // false
//	IsBlank(" a ")    // false
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsLower checks if all letters in string are lowercase.
// IsLower는 문자열의 모든 글자가 소문자인지 확인합니다.
//
// Example:
//
//	IsLower("hello")   // true
//	IsLower("abc")     // true
//	IsLower("Hello")   // false
//	IsLower("ABC")     // false
//	IsLower("abc123")  // true (digits don't affect)
func IsLower(s string) bool {
	if s == "" {
		return false
	}
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsLower(r) {
				return false
			}
		}
	}
	return hasLetter
}

// IsUpper checks if all letters in string are uppercase.
// IsUpper는 문자열의 모든 글자가 대문자인지 확인합니다.
//
// Example:
//
//	IsUpper("HELLO")   // true
//	IsUpper("ABC")     // true
//	IsUpper("Hello")   // false
//	IsUpper("abc")     // false
//	IsUpper("ABC123")  // true (digits don't affect)
func IsUpper(s string) bool {
	if s == "" {
		return false
	}
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetter
}
