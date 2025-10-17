package stringutil

import (
	"regexp"
	"strings"
	"unicode"
)

// =============================================================================
// File: manipulation.go
// Purpose: String Manipulation and Transformation Operations
// 파일: manipulation.go
// 목적: 문자열 조작 및 변환 연산
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The manipulation.go file provides fundamental string manipulation operations
// that are frequently needed in everyday programming. These functions handle
// common tasks like truncating text, reversing strings, removing duplicates,
// cleaning whitespace, and extracting substrings. All functions are Unicode-safe
// and designed to work correctly with international characters.
//
// manipulation.go 파일은 일상 프로그래밍에서 자주 필요한 기본적인 문자열 조작
// 연산을 제공합니다. 이러한 함수는 텍스트 자르기, 문자열 뒤집기, 중복 제거,
// 공백 정리, 부분 문자열 추출과 같은 일반적인 작업을 처리합니다. 모든 함수는
// 유니코드 안전하며 국제 문자와 올바르게 작동하도록 설계되었습니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Unicode-First**: All operations use rune-based indexing, not bytes
//    **유니코드 우선**: 모든 연산은 바이트가 아닌 rune 기반 인덱싱 사용
//
// 2. **Simplicity**: Convert 10-20 lines of repetitive code into single calls
//    **단순성**: 10-20줄의 반복 코드를 단일 호출로 변환
//
// 3. **Safe Defaults**: Functions handle edge cases gracefully (empty strings, out-of-bounds)
//    **안전한 기본값**: 함수는 엣지 케이스를 우아하게 처리 (빈 문자열, 범위 초과)
//
// 4. **No External Dependencies**: Uses only standard library
//    **외부 의존성 없음**: 표준 라이브러리만 사용
//
// 5. **Predictable Behavior**: Consistent handling of nil/empty strings across all functions
//    **예측 가능한 동작**: 모든 함수에서 nil/빈 문자열의 일관된 처리
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. TRUNCATION OPERATIONS (자르기 연산)
//    - Truncate: Truncate string to length with "..." suffix
//      Truncate: 길이로 문자열 자르고 "..." 추가
//    - TruncateWithSuffix: Truncate with custom suffix
//      TruncateWithSuffix: 사용자 정의 suffix로 자르기
//
// 2. TRANSFORMATION OPERATIONS (변환 연산)
//    - Reverse: Reverse string character order
//      Reverse: 문자열 순서 뒤집기
//    - SwapCase: Swap uppercase and lowercase
//      SwapCase: 대소문자 반전
//    - Repeat: Repeat string n times
//      Repeat: 문자열 n번 반복
//
// 3. CAPITALIZATION OPERATIONS (대문자화 연산)
//    - Capitalize: Capitalize first letter of each word
//      Capitalize: 각 단어의 첫 글자 대문자화
//    - CapitalizeFirst: Capitalize only first letter
//      CapitalizeFirst: 첫 글자만 대문자화
//
// 4. CLEANUP OPERATIONS (정리 연산)
//    - Clean: Trim and deduplicate whitespace
//      Clean: 공백 제거 및 중복 공백 정리
//    - RemoveSpaces: Remove all whitespace
//      RemoveSpaces: 모든 공백 제거
//    - RemoveDuplicates: Remove duplicate characters
//      RemoveDuplicates: 중복 문자 제거
//    - RemoveSpecialChars: Remove non-alphanumeric characters
//      RemoveSpecialChars: 영숫자가 아닌 문자 제거
//
// 5. EXTRACTION OPERATIONS (추출 연산)
//    - Substring: Extract substring by index range
//      Substring: 인덱스 범위로 부분 문자열 추출
//    - Left: Extract leftmost n characters
//      Left: 가장 왼쪽 n개 문자 추출
//    - Right: Extract rightmost n characters
//      Right: 가장 오른쪽 n개 문자 추출
//
// 6. INSERTION OPERATIONS (삽입 연산)
//    - Insert: Insert string at index
//      Insert: 인덱스에 문자열 삽입
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// Truncate(s string, length int) string
// - Purpose: Truncate string to specified length with "..." suffix
// - 목적: 지정된 길이로 문자열 자르고 "..." 추가
// - Time Complexity: O(n) where n is length
// - 시간 복잡도: O(n), n은 길이
// - Space Complexity: O(n) for new string
// - 공간 복잡도: O(n), 새 문자열용
// - Unicode-Safe: Uses rune count, not bytes
// - 유니코드 안전: 바이트가 아닌 rune 수 사용
// - Use Cases: Display preview text, limit output length, UI truncation
// - 사용 사례: 미리보기 텍스트 표시, 출력 길이 제한, UI 자르기
//
// Reverse(s string) string
// - Purpose: Reverse character order in string
// - 목적: 문자열의 문자 순서 뒤집기
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Unicode-Safe: Correctly handles multi-byte characters
// - 유니코드 안전: 멀티바이트 문자 올바르게 처리
// - Use Cases: Palindrome checking, text effects, algorithm puzzles
// - 사용 사례: 회문 확인, 텍스트 효과, 알고리즘 퍼즐
//
// Clean(s string) string
// - Purpose: Trim and normalize whitespace
// - 목적: 공백 제거 및 정규화
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: Trims leading/trailing spaces, deduplicates internal spaces
// - 동작: 앞뒤 공백 제거, 내부 공백 중복 제거
// - Use Cases: User input sanitization, text formatting, data cleaning
// - 사용 사례: 사용자 입력 정제, 텍스트 포맷팅, 데이터 정리
//
// Substring(s string, start, end int) string
// - Purpose: Extract substring by index range
// - 목적: 인덱스 범위로 부분 문자열 추출
// - Time Complexity: O(n) where n is substring length
// - 시간 복잡도: O(n), n은 부분 문자열 길이
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Auto-Adjustment: Handles out-of-bounds indices gracefully
// - 자동 조정: 범위 초과 인덱스를 우아하게 처리
// - Use Cases: Text parsing, data extraction, string slicing
// - 사용 사례: 텍스트 파싱, 데이터 추출, 문자열 슬라이싱
//
// RemoveDuplicates(s string) string
// - Purpose: Remove duplicate characters, keeping first occurrence
// - 목적: 중복 문자 제거, 첫 번째 발생만 유지
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n) for map and result
// - 공간 복잡도: O(n), 맵과 결과용
// - Order: Preserves first occurrence order
// - 순서: 첫 번째 발생 순서 유지
// - Use Cases: Unique character detection, data deduplication
// - 사용 사례: 고유 문자 감지, 데이터 중복 제거
//
// Insert(s string, index int, insert string) string
// - Purpose: Insert string at specified index
// - 목적: 지정된 인덱스에 문자열 삽입
// - Time Complexity: O(n + m) where n is string length, m is insert length
// - 시간 복잡도: O(n + m), n은 문자열 길이, m은 삽입 길이
// - Space Complexity: O(n + m)
// - 공간 복잡도: O(n + m)
// - Auto-Adjustment: Clamps index to valid range [0, len]
// - 자동 조정: 인덱스를 유효한 범위 [0, len]로 제한
// - Use Cases: Template insertion, text editing, string building
// - 사용 사례: 템플릿 삽입, 텍스트 편집, 문자열 구성
//
// UNICODE HANDLING
// 유니코드 처리
// ----------------
// All functions in this file are Unicode-safe, meaning they correctly handle:
// - Multi-byte UTF-8 characters (e.g., emoji: 😀, Chinese: 你好)
// - Grapheme clusters (combining characters)
// - Right-to-left scripts (Arabic, Hebrew)
// - Zero-width characters
//
// 이 파일의 모든 함수는 유니코드 안전하며, 다음을 올바르게 처리합니다:
// - 멀티바이트 UTF-8 문자 (예: 이모지: 😀, 중국어: 你好)
// - 그래핌 클러스터 (결합 문자)
// - 오른쪽에서 왼쪽 스크립트 (아랍어, 히브리어)
// - 너비 0 문자
//
// Implementation: Functions use []rune conversion for Unicode-safe indexing.
// 구현: 함수는 유니코드 안전 인덱싱을 위해 []rune 변환 사용.
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - Truncate/TruncateWithSuffix: O(n) - rune conversion
//   Truncate/TruncateWithSuffix: O(n) - rune 변환
// - Reverse: O(n) - in-place swap after rune conversion
//   Reverse: O(n) - rune 변환 후 제자리 교환
// - Capitalize/CapitalizeFirst: O(n) - single pass
//   Capitalize/CapitalizeFirst: O(n) - 단일 패스
// - RemoveDuplicates: O(n) - hash map lookup
//   RemoveDuplicates: O(n) - 해시 맵 조회
// - RemoveSpaces: O(n) - string replacement
//   RemoveSpaces: O(n) - 문자열 치환
// - RemoveSpecialChars: O(n) - regex replacement
//   RemoveSpecialChars: O(n) - 정규식 치환
// - Clean: O(n) - trim + regex
//   Clean: O(n) - 제거 + 정규식
// - Repeat: O(n * count) - concatenation
//   Repeat: O(n * count) - 연결
// - Substring/Left/Right: O(n) - rune slicing
//   Substring/Left/Right: O(n) - rune 슬라이싱
// - Insert: O(n + m) - rune concatenation
//   Insert: O(n + m) - rune 연결
// - SwapCase: O(n) - single pass
//   SwapCase: O(n) - 단일 패스
//
// Space Complexities:
// 공간 복잡도:
// - Most functions: O(n) for rune array + result string
//   대부분의 함수: O(n), rune 배열 + 결과 문자열용
// - RemoveDuplicates: O(n) for map + result
//   RemoveDuplicates: O(n), 맵 + 결과용
// - Insert: O(n + m) for new string
//   Insert: O(n + m), 새 문자열용
//
// Optimization Tips:
// 최적화 팁:
// 1. Avoid multiple truncations on same string - cache result
//    동일한 문자열에 여러 번 자르기 피하기 - 결과 캐시
// 2. Use Clean instead of multiple trim/replace calls
//    여러 trim/replace 호출 대신 Clean 사용
// 3. For repeated operations, consider strings.Builder
//    반복 연산의 경우 strings.Builder 고려
// 4. RemoveSpecialChars compiles regex on each call - cache if needed
//    RemoveSpecialChars는 각 호출마다 정규식 컴파일 - 필요 시 캐시
// 5. For ASCII-only strings, byte operations may be faster
//    ASCII 전용 문자열의 경우 바이트 연산이 더 빠를 수 있음
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Strings:
// 빈 문자열:
// - All functions safely handle empty strings
//   모든 함수는 빈 문자열을 안전하게 처리
// - Truncate("", 10) returns ""
//   Truncate("", 10)는 "" 반환
// - Reverse("") returns ""
//   Reverse("")는 "" 반환
// - Clean("") returns ""
//   Clean("")는 "" 반환
//
// Out-of-Bounds Indices:
// 범위 초과 인덱스:
// - Substring auto-adjusts indices to valid range
//   Substring은 인덱스를 유효한 범위로 자동 조정
// - Left/Right return entire string if n > length
//   Left/Right는 n > length이면 전체 문자열 반환
// - Insert clamps index to [0, len(s)]
//   Insert는 인덱스를 [0, len(s)]로 제한
//
// Negative Indices:
// 음수 인덱스:
// - Substring treats negative indices as 0
//   Substring은 음수 인덱스를 0으로 처리
// - Insert treats negative index as 0
//   Insert는 음수 인덱스를 0으로 처리
//
// Negative Count:
// 음수 카운트:
// - Repeat returns empty string for count < 0
//   Repeat는 count < 0일 때 빈 문자열 반환
//
// Whitespace-Only Strings:
// 공백만 있는 문자열:
// - Clean("   ") returns ""
//   Clean("   ")는 "" 반환
// - RemoveSpaces("   ") returns ""
//   RemoveSpaces("   ")는 "" 반환
//
// Special Characters:
// 특수 문자:
// - RemoveSpecialChars keeps only [a-zA-Z0-9\s]
//   RemoveSpecialChars는 [a-zA-Z0-9\s]만 유지
// - Unicode letters outside ASCII range are removed
//   ASCII 범위 밖의 유니코드 문자는 제거됨
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Truncating Long Text for Display
//    디스플레이를 위한 긴 텍스트 자르기:
//
//    description := "This is a very long product description..."
//    preview := stringutil.Truncate(description, 50)
//    // "This is a very long product description that ..."
//    // Useful for list views, cards, previews
//    // 목록 뷰, 카드, 미리보기에 유용
//
// 2. Cleaning User Input
//    사용자 입력 정리:
//
//    userInput := "  hello    world  \t\n"
//    cleaned := stringutil.Clean(userInput)
//    // "hello world"
//    // Removes extra whitespace, tabs, newlines
//    // 추가 공백, 탭, 개행 제거
//
// 3. Extracting Substring Safely
//    안전하게 부분 문자열 추출:
//
//    text := "Hello, World!"
//    // No need to check bounds
//    // 범위 확인 불필요
//    part := stringutil.Substring(text, 0, 100)
//    // "Hello, World!" (auto-adjusted)
//    // "Hello, World!" (자동 조정됨)
//
// 4. Reversing Strings for Algorithms
//    알고리즘을 위한 문자열 뒤집기:
//
//    word := "racecar"
//    reversed := stringutil.Reverse(word)
//    isPalindrome := word == reversed
//    // true
//    // Works with Unicode: Reverse("안녕") == "녕안"
//    // 유니코드 작동: Reverse("안녕") == "녕안"
//
// 5. Removing Duplicate Characters
//    중복 문자 제거:
//
//    input := "programming"
//    unique := stringutil.RemoveDuplicates(input)
//    // "progamin" (keeps first occurrence)
//    // "progamin" (첫 번째 발생 유지)
//
// 6. Capitalizing Text for Titles
//    제목을 위한 텍스트 대문자화:
//
//    title := "the quick brown fox"
//    formatted := stringutil.Capitalize(title)
//    // "The Quick Brown Fox"
//    // For first letter only: CapitalizeFirst
//    // 첫 글자만: CapitalizeFirst
//
// 7. Inserting Text at Position
//    위치에 텍스트 삽입:
//
//    greeting := "Hello, World!"
//    modified := stringutil.Insert(greeting, 7, "Beautiful ")
//    // "Hello, Beautiful World!"
//    // Useful for template processing
//    // 템플릿 처리에 유용
//
// 8. Extracting First/Last N Characters
//    처음/마지막 N개 문자 추출:
//
//    filename := "document.pdf"
//    extension := stringutil.Right(filename, 4)
//    // ".pdf"
//    prefix := stringutil.Left(filename, 3)
//    // "doc"
//
// 9. Sanitizing Input for Database
//    데이터베이스용 입력 정제:
//
//    username := "user@#$123"
//    safe := stringutil.RemoveSpecialChars(username)
//    // "user123" (alphanumeric only)
//    // "user123" (영숫자만)
//
// 10. Creating Repeated Patterns
//     반복 패턴 생성:
//
//     separator := stringutil.Repeat("-", 40)
//     // "----------------------------------------"
//     // Useful for borders, separators in CLI output
//     // CLI 출력의 테두리, 구분선에 유용
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// Truncate vs Substring
// - Truncate: Adds suffix ("..."), for display
//   Truncate: 접미사 추가 ("..."), 디스플레이용
// - Substring: Exact extraction, no suffix
//   Substring: 정확한 추출, 접미사 없음
// - Use Truncate for: User-facing truncation
//   Truncate 사용: 사용자 대면 자르기
// - Use Substring for: Exact slicing operations
//   Substring 사용: 정확한 슬라이싱 연산
//
// Clean vs RemoveSpaces
// - Clean: Trims + deduplicates spaces, keeps single spaces
//   Clean: 제거 + 공백 중복 제거, 단일 공백 유지
// - RemoveSpaces: Removes all spaces completely
//   RemoveSpaces: 모든 공백 완전 제거
// - Use Clean for: Normalizing human-readable text
//   Clean 사용: 사람이 읽을 수 있는 텍스트 정규화
// - Use RemoveSpaces for: Removing all whitespace
//   RemoveSpaces 사용: 모든 공백 제거
//
// Left vs Substring(s, 0, n)
// - Left: Simpler API, handles overflow
//   Left: 더 간단한 API, 오버플로 처리
// - Substring: More flexible with start/end
//   Substring: start/end로 더 유연함
// - Performance: Identical
//   성능: 동일
// - Use Left for: Simple prefix extraction
//   Left 사용: 간단한 접두사 추출
//
// Capitalize vs CapitalizeFirst
// - Capitalize: Capitalizes every word
//   Capitalize: 모든 단어 대문자화
// - CapitalizeFirst: Only first character
//   CapitalizeFirst: 첫 문자만
// - Use Capitalize for: Titles, headings
//   Capitalize 사용: 제목, 헤딩
// - Use CapitalizeFirst for: Sentences
//   CapitalizeFirst 사용: 문장
//
// RemoveDuplicates vs sliceutil.Unique
// - RemoveDuplicates: For strings (characters)
//   RemoveDuplicates: 문자열용 (문자)
// - sliceutil.Unique: For slices of any type
//   sliceutil.Unique: 모든 타입의 슬라이스용
// - Both preserve order of first occurrence
//   둘 다 첫 번째 발생 순서 유지
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All functions in this file are thread-safe for read-only operations since they
// don't modify the input string (strings are immutable in Go). However, if you're
// using shared mutable state (e.g., caching regex patterns), you need synchronization.
//
// 이 파일의 모든 함수는 입력 문자열을 수정하지 않으므로 읽기 전용 연산에
// 스레드 안전합니다 (Go에서 문자열은 불변). 그러나 공유 가변 상태
// (예: 정규식 패턴 캐싱)를 사용하는 경우 동기화가 필요합니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     // Safe - strings are immutable
//     // 안전 - 문자열은 불변
//     go func() {
//         result := stringutil.Truncate(sharedString, 10)
//     }()
//
//     // Safe - no shared state
//     // 안전 - 공유 상태 없음
//     go func() {
//         cleaned := stringutil.Clean(userInput)
//     }()
//
// Not Thread-Safe:
// 스레드 안전하지 않음:
//
//     // If you cache compiled regex patterns, use sync.Map or mutex
//     // 컴파일된 정규식 패턴을 캐시하는 경우 sync.Map 또는 mutex 사용
//     var regexCache = make(map[string]*regexp.Regexp)
//     // Needs synchronization for concurrent access
//     // 동시 접근을 위한 동기화 필요
//
// RELATED FILES
// 관련 파일
// -------------
// - case.go: Case conversion operations (ToSnakeCase, ToCamelCase, etc.)
//   case.go: 케이스 변환 연산 (ToSnakeCase, ToCamelCase 등)
// - validation.go: String validation functions (IsEmail, IsURL, etc.)
//   validation.go: 문자열 검증 함수 (IsEmail, IsURL 등)
// - search.go: Search and matching operations
//   search.go: 검색 및 매칭 연산
// - comparison.go: String comparison utilities
//   comparison.go: 문자열 비교 유틸리티
// - formatting.go: Advanced formatting operations
//   formatting.go: 고급 포맷팅 연산
// - unicode.go: Unicode-specific operations
//   unicode.go: 유니코드 전용 연산
//
// =============================================================================

// Truncate truncates a string to the specified length and appends "...".
// Truncate는 문자열을 지정된 길이로 자르고 "..."를 추가합니다.
//
// Unicode-safe: uses rune count, not byte count.
// 유니코드 안전: 바이트 수가 아닌 rune 수 사용.
//
// Example:
//
// Truncate("Hello World", 8)    // "Hello..."
// Truncate("안녕하세요", 3)        // "안녕하..."
func Truncate(s string, length int) string {
	return TruncateWithSuffix(s, length, "...")
}

// TruncateWithSuffix truncates a string with a custom suffix.
// TruncateWithSuffix는 사용자 정의 suffix로 문자열을 자릅니다.
//
// Example:
//
// TruncateWithSuffix("Hello World", 8, "…")  // "Hello Wo…"
// TruncateWithSuffix("안녕하세요", 3, "…")      // "안녕하…"
func TruncateWithSuffix(s string, length int, suffix string) string {
	runes := []rune(s)
	if len(runes) <= length {
		return s
	}
	return string(runes[:length]) + suffix
}

// Reverse reverses a string (Unicode-safe).
// Reverse는 문자열을 뒤집습니다 (유니코드 안전).
//
// Example:
//
// Reverse("hello")  // "olleh"
// Reverse("안녕")    // "녕안"
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Capitalize capitalizes the first letter of each word.
// Capitalize는 각 단어의 첫 글자를 대문자로 만듭니다.
//
// Example:
//
//	Capitalize("hello world")  // "Hello World"
//	Capitalize("hello-world")  // "Hello-World"
func Capitalize(s string) string {
	return strings.Title(s)
}

// CapitalizeFirst capitalizes only the first letter of the string.
// CapitalizeFirst는 문자열의 첫 글자만 대문자로 만듭니다.
//
// Example:
//
//	CapitalizeFirst("hello world")  // "Hello world"
//	CapitalizeFirst("HELLO WORLD")  // "HELLO WORLD"
func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// RemoveDuplicates removes duplicate characters from a string.
// RemoveDuplicates는 문자열에서 중복 문자를 제거합니다.
//
// Example:
//
//	RemoveDuplicates("hello")  // "helo"
//	RemoveDuplicates("aabbcc")  // "abc"
func RemoveDuplicates(s string) string {
	seen := make(map[rune]bool)
	var result []rune
	for _, r := range s {
		if !seen[r] {
			seen[r] = true
			result = append(result, r)
		}
	}
	return string(result)
}

// RemoveSpaces removes all whitespace from a string.
// RemoveSpaces는 문자열에서 모든 공백을 제거합니다.
//
// Example:
//
//	RemoveSpaces("h e l l o")  // "hello"
//	RemoveSpaces("  hello world  ")  // "helloworld"
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\t", ""), "\n", "")
}

// RemoveSpecialChars removes special characters, keeping only alphanumeric and spaces.
// RemoveSpecialChars는 특수 문자를 제거하고 영숫자와 공백만 유지합니다.
//
// Example:
//
//	RemoveSpecialChars("hello@#$123")  // "hello123"
//	RemoveSpecialChars("a!b@c#123")    // "abc123"
func RemoveSpecialChars(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	return re.ReplaceAllString(s, "")
}

// Clean trims whitespace and deduplicates spaces.
// Clean은 공백을 제거하고 중복 공백을 정리합니다.
//
// Example:
//
//	Clean("  hello   world  ")  // "hello world"
//	Clean("\t\nhello\t\nworld")  // "hello world"
func Clean(s string) string {
	// Trim leading/trailing spaces
	// 앞뒤 공백 제거
	s = strings.TrimSpace(s)

	// Replace multiple spaces with single space
	// 중복 공백을 단일 공백으로
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")

	return s
}

// Repeat repeats a string n times.
// Repeat는 문자열을 n번 반복합니다.
//
// Unicode-safe: works correctly with all Unicode characters.
// 유니코드 안전: 모든 유니코드 문자와 정상 작동.
//
// Example:
//
// Repeat("hello", 3)  // "hellohellohello"
// Repeat("안녕", 2)     // "안녕안녕"
//
//	Repeat("*", 5)      // "*****"
func Repeat(s string, count int) string {
	if count < 0 {
		return ""
	}
	return strings.Repeat(s, count)
}

// Substring extracts a substring from start to end index (Unicode-safe).
// Substring은 start부터 end 인덱스까지 부분 문자열을 추출합니다 (유니코드 안전).
//
// Parameters:
// - start: starting index (inclusive)
// - end: ending index (exclusive)
//
// If indices are out of bounds, they are adjusted to valid range.
// 인덱스가 범위를 벗어나면 유효한 범위로 조정됩니다.
//
// Example:
//
//	Substring("hello world", 0, 5)   // "hello"
//
// Substring("hello world", 6, 11)  // "world"
// Substring("안녕하세요", 0, 2)       // "안녕"
//
//	Substring("hello", 0, 100)       // "hello" (auto-adjusted)
func Substring(s string, start, end int) string {
	runes := []rune(s)
	length := len(runes)

	// Adjust negative indices
	// 음수 인덱스 조정
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}

	// Adjust out-of-bounds indices
	// 범위 초과 인덱스 조정
	if start > length {
		start = length
	}
	if end > length {
		end = length
	}

	// Ensure start <= end
	// start <= end 보장
	if start > end {
		start, end = end, start
	}

	return string(runes[start:end])
}

// Left returns the leftmost n characters of a string (Unicode-safe).
// Left는 문자열의 가장 왼쪽 n개 문자를 반환합니다 (유니코드 안전).
//
// If n is greater than string length, returns the entire string.
// n이 문자열 길이보다 크면 전체 문자열을 반환합니다.
//
// Example:
//
// Left("hello world", 5)  // "hello"
// Left("안녕하세요", 2)      // "안녕"
//
//	Left("hello", 10)       // "hello"
func Left(s string, n int) string {
	if n <= 0 {
		return ""
	}
	runes := []rune(s)
	if n >= len(runes) {
		return s
	}
	return string(runes[:n])
}

// Right returns the rightmost n characters of a string (Unicode-safe).
// Right는 문자열의 가장 오른쪽 n개 문자를 반환합니다 (유니코드 안전).
//
// If n is greater than string length, returns the entire string.
// n이 문자열 길이보다 크면 전체 문자열을 반환합니다.
//
// Example:
//
// Right("hello world", 5)  // "world"
// Right("안녕하세요", 2)       // "세요"
//
//	Right("hello", 10)       // "hello"
func Right(s string, n int) string {
	if n <= 0 {
		return ""
	}
	runes := []rune(s)
	length := len(runes)
	if n >= length {
		return s
	}
	return string(runes[length-n:])
}

// Insert inserts a string at the specified index (Unicode-safe).
// Insert는 지정된 인덱스에 문자열을 삽입합니다 (유니코드 안전).
//
// If index is negative or greater than length, it's adjusted to valid range.
// 인덱스가 음수이거나 길이보다 크면 유효한 범위로 조정됩니다.
//
// Example:
//
//	Insert("hello world", 5, ",")    // "hello, world"
//
// Insert("hello", 0, "say ")       // "say hello"
// Insert("안녕하세요", 2, " 반갑습니다 ")  // "안녕 반갑습니다 하세요"
func Insert(s string, index int, insert string) string {
	runes := []rune(s)
	length := len(runes)

	// Adjust negative index
	// 음수 인덱스 조정
	if index < 0 {
		index = 0
	}
	// Adjust out-of-bounds index
	// 범위 초과 인덱스 조정
	if index > length {
		index = length
	}

	// Build result
	// 결과 생성
	result := make([]rune, 0, length+len([]rune(insert)))
	result = append(result, runes[:index]...)
	result = append(result, []rune(insert)...)
	result = append(result, runes[index:]...)

	return string(result)
}

// SwapCase swaps the case of all letters in a string.
// SwapCase는 문자열의 모든 글자의 대소문자를 반전합니다.
//
// Uppercase becomes lowercase and vice versa.
// 대문자는 소문자로, 소문자는 대문자로 변환됩니다.
//
// Example:
//
//	SwapCase("Hello World")  // "hELLO wORLD"
//	SwapCase("GoLang")       // "gOlANG"
//	SwapCase("ABC123xyz")    // "abc123XYZ"
func SwapCase(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return string(runes)
}
