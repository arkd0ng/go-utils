package stringutil

import "strings"

// =============================================================================
// File: search.go
// Purpose: String Search and Replacement Operations
// 파일: search.go
// 목적: 문자열 검색 및 치환 연산
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The search.go file provides enhanced string searching and replacement functions
// that extend the standard library's capabilities. These functions handle common
// patterns like searching for multiple substrings at once, case-insensitive
// operations, and batch replacements. They simplify code that would otherwise
// require loops or complex logic, making string searching more ergonomic and
// expressive.
//
// search.go 파일은 표준 라이브러리의 기능을 확장하는 향상된 문자열 검색 및
// 치환 함수를 제공합니다. 이러한 함수는 여러 부분 문자열을 한 번에 검색하거나
// 대소문자를 무시하는 연산, 배치 치환과 같은 일반적인 패턴을 처리합니다.
// 이들은 그렇지 않으면 루프나 복잡한 로직이 필요한 코드를 단순화하여 문자열
// 검색을 더 인체공학적이고 표현력 있게 만듭니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Batch Operations**: Handle multiple search/replace operations in one call
//    **배치 연산**: 한 번의 호출로 여러 검색/치환 연산 처리
//
// 2. **Ergonomic API**: Reduce boilerplate for common searching patterns
//    **인체공학적 API**: 일반적인 검색 패턴의 보일러플레이트 감소
//
// 3. **Flexible Matching**: Support various matching strategies (any, all, case-insensitive)
//    **유연한 매칭**: 다양한 매칭 전략 지원 (any, all, 대소문자 무시)
//
// 4. **Composability**: Functions work well with other string utilities
//    **조합 가능성**: 다른 문자열 유틸리티와 잘 작동하는 함수
//
// 5. **Performance-Conscious**: Minimize unnecessary operations
//    **성능 의식**: 불필요한 연산 최소화
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. MULTI-SUBSTRING SEARCH (다중 부분 문자열 검색)
//    - ContainsAny: Check if string contains any of given substrings
//      ContainsAny: 문자열에 주어진 부분 문자열 중 하나라도 포함되는지 확인
//    - ContainsAll: Check if string contains all given substrings
//      ContainsAll: 문자열에 주어진 모든 부분 문자열이 포함되는지 확인
//
// 2. PREFIX/SUFFIX MATCHING (접두사/접미사 매칭)
//    - StartsWithAny: Check if string starts with any prefix
//      StartsWithAny: 문자열이 접두사 중 하나로 시작하는지 확인
//    - EndsWithAny: Check if string ends with any suffix
//      EndsWithAny: 문자열이 접미사 중 하나로 끝나는지 확인
//
// 3. BATCH REPLACEMENT (배치 치환)
//    - ReplaceAll: Replace multiple substrings using a map
//      ReplaceAll: 맵을 사용하여 여러 부분 문자열 치환
//
// 4. CASE-INSENSITIVE OPERATIONS (대소문자 무시 연산)
//    - ReplaceIgnoreCase: Replace substring ignoring case
//      ReplaceIgnoreCase: 대소문자를 무시하고 부분 문자열 치환
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// ContainsAny(s string, substrs []string) bool
// - Purpose: Check if string contains any of given substrings
// - 목적: 문자열에 주어진 부분 문자열 중 하나라도 포함되는지 확인
// - Time Complexity: O(n * m) where n is number of substrings, m is avg length
// - 시간 복잡도: O(n * m), n은 부분 문자열 수, m은 평균 길이
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Early Exit: Returns true on first match
// - 조기 종료: 첫 번째 일치 시 true 반환
// - Use Cases: Checking for blacklisted words, file type detection, keyword matching
// - 사용 사례: 차단 단어 확인, 파일 타입 감지, 키워드 매칭
//
// ContainsAll(s string, substrs []string) bool
// - Purpose: Check if string contains all given substrings
// - 목적: 문자열에 주어진 모든 부분 문자열이 포함되는지 확인
// - Time Complexity: O(n * m)
// - 시간 복잡도: O(n * m)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Early Exit: Returns false on first non-match
// - 조기 종료: 첫 번째 불일치 시 false 반환
// - Use Cases: Validation, requirement checking, filter criteria
// - 사용 사례: 검증, 요구사항 확인, 필터 기준
//
// StartsWithAny(s string, prefixes []string) bool
// - Purpose: Check if string starts with any of given prefixes
// - 목적: 문자열이 주어진 접두사 중 하나로 시작하는지 확인
// - Time Complexity: O(n * p) where n is number of prefixes, p is avg prefix length
// - 시간 복잡도: O(n * p), n은 접두사 수, p는 평균 접두사 길이
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Protocol detection (http://, https://), file path checking, command parsing
// - 사용 사례: 프로토콜 감지 (http://, https://), 파일 경로 확인, 명령 파싱
//
// EndsWithAny(s string, suffixes []string) bool
// - Purpose: Check if string ends with any of given suffixes
// - 목적: 문자열이 주어진 접미사 중 하나로 끝나는지 확인
// - Time Complexity: O(n * s) where n is number of suffixes, s is avg suffix length
// - 시간 복잡도: O(n * s), n은 접미사 수, s는 평균 접미사 길이
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: File extension detection, URL routing, content type detection
// - 사용 사례: 파일 확장자 감지, URL 라우팅, 콘텐츠 타입 감지
//
// ReplaceAll(s string, replacements map[string]string) string
// - Purpose: Replace multiple substrings using a map
// - 목적: 맵을 사용하여 여러 부분 문자열 치환
// - Time Complexity: O(n * m) where n is number of replacements, m is string length
// - 시간 복잡도: O(n * m), n은 치환 수, m은 문자열 길이
// - Space Complexity: O(m) for result string
// - 공간 복잡도: O(m), 결과 문자열용
// - Behavior: Sequential replacement (order matters)
// - 동작: 순차 치환 (순서 중요)
// - Use Cases: Template expansion, placeholder replacement, text normalization
// - 사용 사례: 템플릿 확장, 플레이스홀더 치환, 텍스트 정규화
//
// ReplaceIgnoreCase(s, old, new string) string
// - Purpose: Replace substring ignoring case
// - 목적: 대소문자를 무시하고 부분 문자열 치환
// - Time Complexity: O(n) where n is string length
// - 시간 복잡도: O(n), n은 문자열 길이
// - Space Complexity: O(n) for result string
// - 공간 복잡도: O(n), 결과 문자열용
// - Case Handling: Matches any case, replaces with exact new string
// - 케이스 처리: 모든 케이스 일치, 정확한 new 문자열로 치환
// - Use Cases: User input normalization, case-insensitive search-and-replace
// - 사용 사례: 사용자 입력 정규화, 대소문자 무시 검색 및 치환
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - ContainsAny/ContainsAll: O(n * m) - iterate substrings, each O(m) search
//   ContainsAny/ContainsAll: O(n * m) - 부분 문자열 반복, 각 O(m) 검색
// - StartsWithAny/EndsWithAny: O(n * p) - iterate prefixes/suffixes, each O(p) check
//   StartsWithAny/EndsWithAny: O(n * p) - 접두사/접미사 반복, 각 O(p) 확인
// - ReplaceAll: O(n * m) - n replacements, each scanning m-length string
//   ReplaceAll: O(n * m) - n개 치환, 각각 m 길이 문자열 스캔
// - ReplaceIgnoreCase: O(m) - single pass with case-insensitive comparison
//   ReplaceIgnoreCase: O(m) - 대소문자 무시 비교를 사용한 단일 패스
//
// Space Complexities:
// 공간 복잡도:
// - ContainsAny/ContainsAll/StartsWithAny/EndsWithAny: O(1) - no extra space
//   ContainsAny/ContainsAll/StartsWithAny/EndsWithAny: O(1) - 추가 공간 없음
// - ReplaceAll: O(m) - result string
//   ReplaceAll: O(m) - 결과 문자열
// - ReplaceIgnoreCase: O(m) - result string + lowercase copies
//   ReplaceIgnoreCase: O(m) - 결과 문자열 + 소문자 복사본
//
// Optimization Tips:
// 최적화 팁:
// 1. For ContainsAny, put most likely matches first for early exit
//    ContainsAny의 경우 조기 종료를 위해 가장 가능성 높은 일치를 먼저 배치
// 2. For ContainsAll, put least likely matches first for early exit
//    ContainsAll의 경우 조기 종료를 위해 가능성 낮은 일치를 먼저 배치
// 3. Use ReplaceAll carefully - order of replacements matters
//    ReplaceAll 신중히 사용 - 치환 순서 중요
// 4. For case-insensitive search, consider precomputing lowercase versions
//    대소문자 무시 검색의 경우 소문자 버전 미리 계산 고려
// 5. For large replacement maps, consider more efficient algorithms (e.g., Aho-Corasick)
//    큰 치환 맵의 경우 더 효율적인 알고리즘 고려 (예: Aho-Corasick)
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Inputs:
// 빈 입력:
// - ContainsAny("", []) returns false (empty slices)
//   ContainsAny("", [])는 false 반환 (빈 슬라이스)
// - ContainsAll("", []) returns true (vacuous truth)
//   ContainsAll("", [])는 true 반환 (공허한 진리)
// - StartsWithAny/EndsWithAny("", []) returns false
//   StartsWithAny/EndsWithAny("", [])는 false 반환
// - ReplaceAll("", {}) returns ""
//   ReplaceAll("", {})는 "" 반환
//
// Empty Substrings:
// 빈 부분 문자열:
// - ContainsAny with empty substring: always true
//   빈 부분 문자열의 ContainsAny: 항상 true
// - ReplaceIgnoreCase with empty old: returns original
//   빈 old의 ReplaceIgnoreCase: 원본 반환
//
// Empty Replacement Map:
// 빈 치환 맵:
// - ReplaceAll(s, {}) returns s unchanged
//   ReplaceAll(s, {})는 변경되지 않은 s 반환
//
// Case Sensitivity:
// 대소문자 구분:
// - Only ReplaceIgnoreCase is case-insensitive
//   ReplaceIgnoreCase만 대소문자 무시
// - All others are case-sensitive
//   다른 모든 것은 대소문자 구분
//
// Overlapping Replacements:
// 겹치는 치환:
// - ReplaceAll applies replacements sequentially
//   ReplaceAll은 순차적으로 치환 적용
// - Order matters: ReplaceAll("abc", {"a":"x", "ab":"y"}) != ReplaceAll("abc", {"ab":"y", "a":"x"})
//   순서 중요: ReplaceAll("abc", {"a":"x", "ab":"y"}) != ReplaceAll("abc", {"ab":"y", "a":"x"})
// - First replacement can affect subsequent ones
//   첫 번째 치환이 이후 치환에 영향 가능
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Checking for Multiple File Extensions
//    여러 파일 확장자 확인:
//
//    filename := "document.pdf"
//    isDocument := stringutil.EndsWithAny(filename, []string{".pdf", ".doc", ".docx", ".txt"})
//    // true
//    // Categorize files by extension
//    // 확장자로 파일 분류
//
// 2. Protocol Detection
//    프로토콜 감지:
//
//    url := "https://example.com"
//    isSecure := stringutil.StartsWithAny(url, []string{"https://", "wss://"})
//    // true
//    // Detect secure protocols
//    // 보안 프로토콜 감지
//
// 3. Blacklist Checking
//    차단 목록 확인:
//
//    content := "This contains spam keywords"
//    blacklist := []string{"spam", "viagra", "casino"}
//    isBad := stringutil.ContainsAny(content, blacklist)
//    // true
//    // Content filtering
//    // 콘텐츠 필터링
//
// 4. Validation with Required Terms
//    필수 용어로 검증:
//
//    terms := "I agree to terms and conditions"
//    required := []string{"agree", "terms", "conditions"}
//    isValid := stringutil.ContainsAll(terms, required)
//    // true
//    // Ensure all required keywords present
//    // 모든 필수 키워드 존재 확인
//
// 5. Template Expansion
//    템플릿 확장:
//
//    template := "Hello {{name}}, your code is {{code}}"
//    replacements := map[string]string{
//        "{{name}}": "Alice",
//        "{{code}}": "12345",
//    }
//    result := stringutil.ReplaceAll(template, replacements)
//    // "Hello Alice, your code is 12345"
//    // Simple template engine
//    // 간단한 템플릿 엔진
//
// 6. Text Normalization
//    텍스트 정규화:
//
//    text := "hello-world_test"
//    replacements := map[string]string{
//        "-": "_",
//        "_": "_",
//    }
//    normalized := stringutil.ReplaceAll(text, replacements)
//    // Normalize separators
//    // 구분자 정규화
//
// 7. Case-Insensitive Word Replacement
//    대소문자 무시 단어 치환:
//
//    text := "Hello WORLD, hello world"
//    result := stringutil.ReplaceIgnoreCase(text, "hello", "hi")
//    // "hi WORLD, hi world"
//    // Replace all variations
//    // 모든 변형 치환
//
// 8. Multiple Keyword Search
//    다중 키워드 검색:
//
//    log := "Error: connection timeout"
//    errorKeywords := []string{"error", "fail", "timeout", "crash"}
//    hasError := stringutil.ContainsAny(strings.ToLower(log), errorKeywords)
//    // true
//    // Log level detection
//    // 로그 레벨 감지
//
// 9. File Type Detection
//    파일 타입 감지:
//
//    filename := "script.sh"
//    isExecutable := stringutil.EndsWithAny(filename, []string{".sh", ".bat", ".exe", ".bin"})
//    // true
//    // Detect executable files
//    // 실행 파일 감지
//
// 10. Route Matching
//     라우트 매칭:
//
//     path := "/api/users/123"
//     isAPI := stringutil.StartsWithAny(path, []string{"/api/", "/v1/", "/v2/"})
//     // true
//     // API route detection
//     // API 라우트 감지
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// ContainsAny vs Multiple strings.Contains Calls
// - ContainsAny: Single function call, cleaner code
//   ContainsAny: 단일 함수 호출, 더 깔끔한 코드
// - Multiple Contains: More verbose with if statements
//   다중 Contains: if 문으로 더 장황함
// - Use ContainsAny for: Checking multiple substrings
//   ContainsAny 사용: 여러 부분 문자열 확인
//
// ReplaceAll (search.go) vs strings.ReplaceAll (stdlib)
// - search.go ReplaceAll: Multiple replacements with map
//   search.go ReplaceAll: 맵으로 여러 치환
// - strings.ReplaceAll: Single replacement
//   strings.ReplaceAll: 단일 치환
// - Use search.go for: Batch replacements
//   search.go 사용: 배치 치환
//
// ReplaceIgnoreCase vs strings.ReplaceAll + strings.ToLower
// - ReplaceIgnoreCase: Direct case-insensitive replacement
//   ReplaceIgnoreCase: 직접 대소문자 무시 치환
// - ToLower + ReplaceAll: Loses case information
//   ToLower + ReplaceAll: 케이스 정보 손실
// - Use ReplaceIgnoreCase for: Preserving original case
//   ReplaceIgnoreCase 사용: 원본 케이스 보존
//
// StartsWithAny vs regexp
// - StartsWithAny: Faster for literal prefixes
//   StartsWithAny: 리터럴 접두사에 더 빠름
// - regexp: More powerful for patterns
//   regexp: 패턴에 더 강력
// - Use StartsWithAny for: Simple prefix matching
//   StartsWithAny 사용: 간단한 접두사 매칭
// - Use regexp for: Complex patterns
//   regexp 사용: 복잡한 패턴
//
// REPLACEMENT ORDER CAVEAT
// 치환 순서 주의사항
// ------------------------
// ReplaceAll applies replacements sequentially, which can lead to unexpected results:
// ReplaceAll은 순차적으로 치환을 적용하므로 예상치 못한 결과가 나올 수 있습니다:
//
// Example:
// 예제:
//     s := "a"
//     // First replaces "a" → "b", then "b" → "c"
//     // 먼저 "a" → "b" 치환, 그 다음 "b" → "c" 치환
//     result := ReplaceAll(s, map[string]string{"a": "b", "b": "c"})
//     // Result could be "c" (if "a"→"b" happens first) or "b" (if "b"→"c" happens first)
//     // 결과는 "c" (만약 "a"→"b"가 먼저) 또는 "b" (만약 "b"→"c"가 먼저)일 수 있음
//
// Note: Go maps are unordered, so iteration order is not guaranteed.
// 참고: Go 맵은 순서가 없으므로 반복 순서가 보장되지 않습니다.
//
// For predictable results, use replacements that don't overlap:
// 예측 가능한 결과를 위해 겹치지 않는 치환 사용:
//     ReplaceAll(s, map[string]string{"{{name}}": "Alice", "{{code}}": "123"})
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
//         result := stringutil.ContainsAny(text, keywords)
//     }()
//
//     go func() {
//         result := stringutil.ReplaceAll(text, replacements)
//     }()
//
//     // All search functions safe for concurrent use
//     // 모든 검색 함수는 동시 사용에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - comparison.go: String comparison operations
//   comparison.go: 문자열 비교 연산
// - validation.go: String validation functions
//   validation.go: 문자열 검증 함수
// - manipulation.go: String manipulation (Clean, Truncate, etc.)
//   manipulation.go: 문자열 조작 (Clean, Truncate 등)
//
// =============================================================================

// ContainsAny returns true if the string contains any of the substrings.
// ContainsAny는 문자열이 부분 문자열 중 하나라도 포함하면 true를 반환합니다.
//
// Example:
//
//	ContainsAny("hello world", []string{"foo", "world"})  // true
//	ContainsAny("hello world", []string{"foo", "bar"})    // false
func ContainsAny(s string, substrs []string) bool {
	for _, substr := range substrs {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// ContainsAll returns true if the string contains all of the substrings.
// ContainsAll은 문자열이 모든 부분 문자열을 포함하면 true를 반환합니다.
//
// Example:
//
//	ContainsAll("hello world", []string{"hello", "world"})  // true
//	ContainsAll("hello world", []string{"hello", "foo"})    // false
func ContainsAll(s string, substrs []string) bool {
	for _, substr := range substrs {
		if !strings.Contains(s, substr) {
			return false
		}
	}
	return true
}

// StartsWithAny returns true if the string starts with any of the prefixes.
// StartsWithAny는 문자열이 접두사 중 하나로 시작하면 true를 반환합니다.
//
// Example:
//
//	StartsWithAny("https://example.com", []string{"http://", "https://"})  // true
//	StartsWithAny("ftp://example.com", []string{"http://", "https://"})    // false
func StartsWithAny(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// EndsWithAny returns true if the string ends with any of the suffixes.
// EndsWithAny는 문자열이 접미사 중 하나로 끝나면 true를 반환합니다.
//
// Example:
//
//	EndsWithAny("file.txt", []string{".txt", ".md"})  // true
//	EndsWithAny("file.jpg", []string{".txt", ".md"})  // false
func EndsWithAny(s string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// ReplaceAll replaces multiple strings at once using a replacement map.
// ReplaceAll은 치환 맵을 사용하여 여러 문자열을 한 번에 치환합니다.
//
// Example:
//
//	ReplaceAll("a b c", map[string]string{"a": "x", "b": "y"})  // "x y c"
//	ReplaceAll("hello world", map[string]string{"hello": "hi", "world": "there"})  // "hi there"
func ReplaceAll(s string, replacements map[string]string) string {
	for old, new := range replacements {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}

// ReplaceIgnoreCase replaces a substring ignoring case.
// ReplaceIgnoreCase는 대소문자를 무시하고 부분 문자열을 치환합니다.
//
// Example:
//
//	ReplaceIgnoreCase("Hello World", "hello", "hi")  // "hi World"
//	ReplaceIgnoreCase("HELLO World", "hello", "hi")  // "hi World"
func ReplaceIgnoreCase(s, old, new string) string {
	// Handle empty old string
	// 빈 old 문자열 처리
	if old == "" {
		return s
	}

	lowerS := strings.ToLower(s)
	lowerOld := strings.ToLower(old)

	var result strings.Builder
	for len(lowerS) > 0 {
		index := strings.Index(lowerS, lowerOld)
		if index == -1 {
			result.WriteString(s)
			break
		}

		// Write everything before the match
		result.WriteString(s[:index])
		// Write the replacement
		result.WriteString(new)

		// Move past the match
		s = s[index+len(old):]
		lowerS = lowerS[index+len(lowerOld):]
	}

	return result.String()
}
