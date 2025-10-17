package stringutil

import "strings"

// =============================================================================
// File: comparison.go
// Purpose: String Comparison Operations
// 파일: comparison.go
// 목적: 문자열 비교 연산
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The comparison.go file provides essential string comparison utilities that
// wrap and enhance the standard library's comparison functions. These functions
// focus on common comparison patterns like case-insensitive equality checking,
// prefix/suffix matching, and provide a consistent, ergonomic API for string
// comparisons throughout the stringutil package.
//
// comparison.go 파일은 표준 라이브러리의 비교 함수를 래핑하고 향상시키는
// 필수적인 문자열 비교 유틸리티를 제공합니다. 이러한 함수는 대소문자 무시
// 동등성 확인, 접두사/접미사 매칭과 같은 일반적인 비교 패턴에 초점을 맞추며,
// stringutil 패키지 전체에서 일관되고 인체공학적인 API를 제공합니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Standard Library Wrapping**: Provide consistent wrappers for stdlib functions
//    **표준 라이브러리 래핑**: stdlib 함수의 일관된 래퍼 제공
//
// 2. **Discoverability**: Make comparison functions easier to find
//    **발견 가능성**: 비교 함수를 찾기 쉽게 만듦
//
// 3. **Consistency**: Uniform naming and behavior across stringutil
//    **일관성**: stringutil 전체에서 균일한 명명 및 동작
//
// 4. **Future Extensions**: Easy to extend with more comparison functions
//    **미래 확장**: 더 많은 비교 함수로 쉽게 확장
//
// 5. **Unicode-Aware**: Proper handling of Unicode case folding
//    **유니코드 인식**: 유니코드 케이스 폴딩 적절히 처리
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. CASE-INSENSITIVE COMPARISON (대소문자 무시 비교)
//    - EqualFold: Check if two strings are equal ignoring case
//      EqualFold: 대소문자를 무시하고 두 문자열이 같은지 확인
//
// 2. PREFIX MATCHING (접두사 매칭)
//    - HasPrefix: Check if string starts with prefix
//      HasPrefix: 문자열이 접두사로 시작하는지 확인
//
// 3. SUFFIX MATCHING (접미사 매칭)
//    - HasSuffix: Check if string ends with suffix
//      HasSuffix: 문자열이 접미사로 끝나는지 확인
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// EqualFold(s1, s2 string) bool
// - Purpose: Compare two strings case-insensitively
// - 목적: 두 문자열을 대소문자 무시하고 비교
// - Time Complexity: O(n) where n is length of shorter string
// - 시간 복잡도: O(n), n은 짧은 문자열의 길이
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Unicode Handling: Uses Unicode case folding (not simple ToLower)
// - 유니코드 처리: 유니코드 케이스 폴딩 사용 (단순 ToLower 아님)
// - Early Exit: Returns false on first mismatch
// - 조기 종료: 첫 번째 불일치 시 false 반환
// - Use Cases: Case-insensitive comparison, user input validation, protocol detection
// - 사용 사례: 대소문자 무시 비교, 사용자 입력 검증, 프로토콜 감지
//
// HasPrefix(s, prefix string) bool
// - Purpose: Check if string starts with prefix
// - 목적: 문자열이 접두사로 시작하는지 확인
// - Time Complexity: O(p) where p is length of prefix
// - 시간 복잡도: O(p), p는 접두사 길이
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Case Sensitivity: Case-sensitive
// - 대소문자 구분: 대소문자 구분함
// - Empty Prefix: Returns true for any string
// - 빈 접두사: 모든 문자열에 대해 true 반환
// - Use Cases: Protocol checking (http://), path validation, command parsing
// - 사용 사례: 프로토콜 확인 (http://), 경로 검증, 명령 파싱
//
// HasSuffix(s, suffix string) bool
// - Purpose: Check if string ends with suffix
// - 목적: 문자열이 접미사로 끝나는지 확인
// - Time Complexity: O(s) where s is length of suffix
// - 시간 복잡도: O(s), s는 접미사 길이
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Case Sensitivity: Case-sensitive
// - 대소문자 구분: 대소문자 구분함
// - Empty Suffix: Returns true for any string
// - 빈 접미사: 모든 문자열에 대해 true 반환
// - Use Cases: File extension checking, URL routing, content type detection
// - 사용 사례: 파일 확장자 확인, URL 라우팅, 콘텐츠 타입 감지
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - EqualFold: O(n) where n is length of shorter string
//   EqualFold: O(n), n은 짧은 문자열의 길이
//   - Early exit on length mismatch or first character difference
//     길이 불일치 또는 첫 문자 차이 시 조기 종료
// - HasPrefix: O(p) where p is prefix length
//   HasPrefix: O(p), p는 접두사 길이
//   - Early exit if prefix longer than string
//     접두사가 문자열보다 길면 조기 종료
// - HasSuffix: O(s) where s is suffix length
//   HasSuffix: O(s), s는 접미사 길이
//   - Early exit if suffix longer than string
//     접미사가 문자열보다 길면 조기 종료
//
// Space Complexities:
// 공간 복잡도:
// - All functions: O(1) - no extra space allocation
//   모든 함수: O(1) - 추가 공간 할당 없음
//
// Optimization Tips:
// 최적화 팁:
// 1. EqualFold is faster than strings.ToLower(s1) == strings.ToLower(s2)
//    EqualFold가 strings.ToLower(s1) == strings.ToLower(s2)보다 빠름
//    - No temporary string allocation
//      임시 문자열 할당 없음
// 2. For repeated prefix/suffix checks, consider precomputing
//    반복적인 접두사/접미사 확인은 미리 계산 고려
// 3. Use HasPrefix/HasSuffix instead of string slicing for checks
//    확인에는 문자열 슬라이싱 대신 HasPrefix/HasSuffix 사용
// 4. For multiple prefixes, use StartsWithAny from search.go
//    여러 접두사는 search.go의 StartsWithAny 사용
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Strings:
// 빈 문자열:
// - EqualFold("", "") returns true
//   EqualFold("", "")는 true 반환
// - HasPrefix("", "") returns true (vacuous truth)
//   HasPrefix("", "")는 true 반환 (공허한 진리)
// - HasSuffix("", "") returns true
//   HasSuffix("", "")는 true 반환
// - HasPrefix("hello", "") returns true (empty prefix matches all)
//   HasPrefix("hello", "")는 true 반환 (빈 접두사는 모두 일치)
// - HasSuffix("hello", "") returns true (empty suffix matches all)
//   HasSuffix("hello", "")는 true 반환 (빈 접미사는 모두 일치)
//
// Length Mismatches:
// 길이 불일치:
// - HasPrefix returns false if prefix longer than string
//   접두사가 문자열보다 길면 HasPrefix는 false 반환
// - HasSuffix returns false if suffix longer than string
//   접미사가 문자열보다 길면 HasSuffix는 false 반환
// - EqualFold returns false if lengths differ (for ASCII)
//   길이가 다르면 EqualFold는 false 반환 (ASCII의 경우)
//
// Unicode Considerations:
// 유니코드 고려사항:
// - EqualFold uses proper Unicode case folding
//   EqualFold는 적절한 유니코드 케이스 폴딩 사용
// - Example: "ß" (German sharp s) vs "SS"
//   예: "ß" (독일어 샤프 s) vs "SS"
//   - EqualFold("ß", "ss") may return true depending on normalization
//     EqualFold("ß", "ss")는 정규화에 따라 true 반환 가능
// - HasPrefix/HasSuffix are byte-based (UTF-8 aware)
//   HasPrefix/HasSuffix는 바이트 기반 (UTF-8 인식)
//
// Case Sensitivity:
// 대소문자 구분:
// - Only EqualFold is case-insensitive
//   EqualFold만 대소문자 무시
// - HasPrefix/HasSuffix are case-sensitive
//   HasPrefix/HasSuffix는 대소문자 구분
// - For case-insensitive prefix/suffix, use:
//   대소문자 무시 접두사/접미사는 다음 사용:
//   strings.ToLower(s) then HasPrefix/HasSuffix
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Case-Insensitive Command Checking
//    대소문자 무시 명령 확인:
//
//    userInput := "EXIT"
//    if stringutil.EqualFold(userInput, "exit") {
//        // Exit application
//        // 애플리케이션 종료
//    }
//    // Handle commands regardless of case
//    // 대소문자 관계없이 명령 처리
//
// 2. Protocol Detection
//    프로토콜 감지:
//
//    url := "https://example.com"
//    isHTTPS := stringutil.HasPrefix(url, "https://")
//    // true
//    // Determine if URL uses secure protocol
//    // URL이 보안 프로토콜 사용하는지 확인
//
// 3. File Extension Checking
//    파일 확장자 확인:
//
//    filename := "document.pdf"
//    isPDF := stringutil.HasSuffix(filename, ".pdf")
//    // true
//    // Validate file type
//    // 파일 타입 검증
//
// 4. Configuration Value Validation
//    설정 값 검증:
//
//    logLevel := "DEBUG"
//    isDebug := stringutil.EqualFold(logLevel, "debug")
//    // true
//    // Case-insensitive config matching
//    // 대소문자 무시 설정 매칭
//
// 5. Path Validation
//    경로 검증:
//
//    path := "/api/users/123"
//    isAPI := stringutil.HasPrefix(path, "/api/")
//    // true
//    // Validate API endpoint
//    // API 엔드포인트 검증
//
// 6. Email Domain Checking
//    이메일 도메인 확인:
//
//    email := "user@example.com"
//    parts := strings.Split(email, "@")
//    if len(parts) == 2 {
//        isCompanyEmail := stringutil.HasSuffix(parts[1], "company.com")
//        // Check corporate email
//        // 회사 이메일 확인
//    }
//
// 7. Header Comparison (HTTP, etc.)
//    헤더 비교 (HTTP 등):
//
//    header := "Content-Type"
//    isContentType := stringutil.EqualFold(header, "content-type")
//    // true
//    // HTTP headers are case-insensitive
//    // HTTP 헤더는 대소문자 무시
//
// 8. File Path Checking
//    파일 경로 확인:
//
//    path := "/home/user/documents/file.txt"
//    isHome := stringutil.HasPrefix(path, "/home/")
//    // true
//    // Validate file location
//    // 파일 위치 검증
//
// 9. Content Type Detection
//    콘텐츠 타입 감지:
//
//    contentType := "application/json; charset=utf-8"
//    isJSON := stringutil.HasPrefix(contentType, "application/json")
//    // true
//    // Detect JSON response
//    // JSON 응답 감지
//
// 10. Boolean Configuration Parsing
//     불린 설정 파싱:
//
//     value := "TRUE"
//     isTrue := stringutil.EqualFold(value, "true") ||
//               stringutil.EqualFold(value, "yes") ||
//               stringutil.EqualFold(value, "1")
//     // Parse various true representations
//     // 다양한 true 표현 파싱
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// EqualFold vs strings.ToLower Comparison
// - EqualFold: No allocation, faster
//   EqualFold: 할당 없음, 더 빠름
// - ToLower + ==: Two allocations, slower
//   ToLower + ==: 두 번의 할당, 더 느림
// - Use EqualFold for: Performance-critical comparisons
//   EqualFold 사용: 성능 중요 비교
//
// EqualFold vs == operator
// - EqualFold: Case-insensitive
//   EqualFold: 대소문자 무시
// - ==: Case-sensitive
//   ==: 대소문자 구분
// - Use EqualFold for: User input, configuration, protocols
//   EqualFold 사용: 사용자 입력, 설정, 프로토콜
//
// HasPrefix vs strings.Index
// - HasPrefix: Only checks start, O(p)
//   HasPrefix: 시작만 확인, O(p)
// - Index: Searches entire string, O(n)
//   Index: 전체 문자열 검색, O(n)
// - Use HasPrefix for: Prefix detection
//   HasPrefix 사용: 접두사 감지
//
// HasSuffix vs string slicing
// - HasSuffix: Clear intent, safe
//   HasSuffix: 명확한 의도, 안전
// - Slicing: Requires bounds checking, error-prone
//   슬라이싱: 경계 확인 필요, 오류 발생 가능
// - Use HasSuffix for: Cleaner code
//   HasSuffix 사용: 더 깔끔한 코드
//
// HasPrefix/HasSuffix vs regexp
// - HasPrefix/HasSuffix: Faster for literal strings
//   HasPrefix/HasSuffix: 리터럴 문자열에 더 빠름
// - regexp: More powerful for patterns
//   regexp: 패턴에 더 강력
// - Use HasPrefix/HasSuffix for: Simple literal matching
//   HasPrefix/HasSuffix 사용: 간단한 리터럴 매칭
//
// UNICODE AND CASE FOLDING
// 유니코드 및 케이스 폴딩
// ------------------------
// EqualFold uses Unicode case folding, which is more sophisticated than simple
// case conversion. Case folding is designed for case-insensitive string matching
// and handles special characters correctly:
//
// EqualFold는 단순 케이스 변환보다 더 정교한 유니코드 케이스 폴딩을 사용합니다.
// 케이스 폴딩은 대소문자 무시 문자열 매칭을 위해 설계되었으며 특수 문자를
// 올바르게 처리합니다:
//
// Examples:
// 예제:
// - EqualFold("Hello", "hello") = true
//   EqualFold("Hello", "hello") = true
// - EqualFold("İstanbul", "istanbul") may vary by locale
//   EqualFold("İstanbul", "istanbul")는 로케일에 따라 다를 수 있음
// - EqualFold("Straße", "STRASSE") depends on normalization
//   EqualFold("Straße", "STRASSE")는 정규화에 따라 다름
//
// Note: For ASCII strings, EqualFold is equivalent to case-insensitive comparison.
// 참고: ASCII 문자열의 경우 EqualFold는 대소문자 무시 비교와 동일합니다.
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All functions in this file are thread-safe as they operate on immutable strings
// and don't use shared mutable state.
//
// 이 파일의 모든 함수는 불변 문자열에서 작동하고 공유 가변 상태를 사용하지
// 않으므로 스레드 안전합니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     go func() {
//         result := stringutil.EqualFold(s1, s2)
//     }()
//
//     go func() {
//         result := stringutil.HasPrefix(text, prefix)
//     }()
//
//     // All comparison functions safe for concurrent use
//     // 모든 비교 함수는 동시 사용에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - search.go: Multi-substring search (ContainsAny, StartsWithAny, etc.)
//   search.go: 다중 부분 문자열 검색 (ContainsAny, StartsWithAny 등)
// - validation.go: String validation functions
//   validation.go: 문자열 검증 함수
// - case.go: Case conversion operations
//   case.go: 케이스 변환 연산
//
// =============================================================================

// EqualFold compares two strings case-insensitively.
// EqualFold는 두 문자열을 대소문자 구분 없이 비교합니다.
//
// Returns true if strings are equal ignoring case.
// 대소문자를 무시하고 문자열이 같으면 true를 반환합니다.
//
// Example:
//
//	EqualFold("hello", "HELLO")  // true
//	EqualFold("GoLang", "golang") // true
//	EqualFold("hello", "world")  // false
func EqualFold(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

// HasPrefix checks if string starts with the given prefix.
// HasPrefix는 문자열이 주어진 접두사로 시작하는지 확인합니다.
//
// Example:
//
//	HasPrefix("hello world", "hello")  // true
//	HasPrefix("golang", "go")          // true
//	HasPrefix("hello", "world")        // false
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix checks if string ends with the given suffix.
// HasSuffix는 문자열이 주어진 접미사로 끝나는지 확인합니다.
//
// Example:
//
//	HasSuffix("hello world", "world")  // true
//	HasSuffix("golang", "lang")        // true
//	HasSuffix("hello", "world")        // false
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
