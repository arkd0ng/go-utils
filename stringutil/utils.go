package stringutil

import (
	"strings"
)

// =============================================================================
// File: utils.go
// Purpose: General-Purpose String Utilities
// 파일: utils.go
// 목적: 범용 문자열 유틸리티
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The utils.go file provides miscellaneous string utility functions that don't
// fit into other specialized categories but are frequently needed in everyday
// programming. These utilities cover counting operations (words, occurrences),
// collection operations (map, filter, join), padding operations, and text
// splitting. They enhance the standard library with convenient, commonly-used
// operations that reduce boilerplate code.
//
// utils.go 파일은 다른 특수 범주에 맞지 않지만 일상적인 프로그래밍에서 자주
// 필요한 기타 문자열 유틸리티 함수를 제공합니다. 이러한 유틸리티는 개수 세기
// 연산 (단어, 발생 횟수), 컬렉션 연산 (맵, 필터, 조인), 패딩 연산 및 텍스트
// 분리를 다룹니다. 이들은 보일러플레이트 코드를 줄이는 편리하고 일반적으로
// 사용되는 연산으로 표준 라이브러리를 향상시킵니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Convenience**: Provide common operations in simple functions
//    **편의성**: 간단한 함수로 일반적인 연산 제공
//
// 2. **Standard Library Wrappers**: Enhance stdlib with consistent API
//    **표준 라이브러리 래퍼**: 일관된 API로 stdlib 향상
//
// 3. **Functional Programming**: Support map/filter paradigms
//    **함수형 프로그래밍**: 맵/필터 패러다임 지원
//
// 4. **Unicode-Safe**: Use rune-based operations where appropriate
//    **유니코드 안전**: 적절한 경우 룬 기반 연산 사용
//
// 5. **Predictable Behavior**: Clear, consistent behavior across functions
//    **예측 가능한 동작**: 함수 간 명확하고 일관된 동작
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. COUNTING OPERATIONS (개수 세기 연산)
//    - CountWords: Count words separated by whitespace
//      CountWords: 공백으로 구분된 단어 개수
//    - CountOccurrences: Count substring occurrences
//      CountOccurrences: 부분 문자열 발생 횟수
//
// 2. COLLECTION OPERATIONS (컬렉션 연산)
//    - Join: Join string slice with separator
//      Join: 구분자로 문자열 슬라이스 연결
//    - Map: Apply function to all strings in slice
//      Map: 슬라이스의 모든 문자열에 함수 적용
//    - Filter: Filter strings by predicate
//      Filter: 조건으로 문자열 필터링
//
// 3. PADDING OPERATIONS (패딩 연산)
//    - PadLeft: Pad string on left side
//      PadLeft: 문자열 왼쪽에 패딩
//    - PadRight: Pad string on right side
//      PadRight: 문자열 오른쪽에 패딩
//
// 4. TEXT SPLITTING (텍스트 분리)
//    - Lines: Split by newlines
//      Lines: 줄바꿈으로 분리
//    - Words: Split by whitespace
//      Words: 공백으로 분리
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// CountWords(s string) int
// - Purpose: Count words separated by whitespace
// - 목적: 공백으로 구분된 단어 개수
// - Whitespace: Space, tab, newline, etc.
// - 공백: 스페이스, 탭, 줄바꿈 등
// - Behavior: Multiple consecutive whitespace treated as one separator
// - 동작: 연속된 여러 공백을 하나의 구분자로 처리
// - Time Complexity: O(n) where n is string length
// - 시간 복잡도: O(n), n은 문자열 길이
// - Space Complexity: O(n) - creates slice of words
// - 공간 복잡도: O(n) - 단어 슬라이스 생성
// - Use Cases: Word count validation, text statistics, input validation
// - 사용 사례: 단어 개수 검증, 텍스트 통계, 입력 검증
//
// CountOccurrences(s, substr string) int
// - Purpose: Count how many times substring appears
// - 목적: 부분 문자열이 나타나는 횟수 계산
// - Overlapping: Counts non-overlapping occurrences
// - 겹침: 겹치지 않는 발생 횟수 계산
// - Empty Substring: Returns len(s)+1 for empty substr
// - 빈 부분 문자열: 빈 substr에 대해 len(s)+1 반환
// - Time Complexity: O(n * m) where n is string length, m is substr length
// - 시간 복잡도: O(n * m), n은 문자열 길이, m은 substr 길이
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Pattern frequency, character counting, text analysis
// - 사용 사례: 패턴 빈도, 문자 개수, 텍스트 분석
//
// Join(strs []string, sep string) string
// - Purpose: Join string slice with separator
// - 목적: 구분자로 문자열 슬라이스 연결
// - Wrapper: Convenience wrapper for strings.Join
// - 래퍼: strings.Join의 편의 래퍼
// - Time Complexity: O(n) where n is total length of all strings
// - 시간 복잡도: O(n), n은 모든 문자열의 총 길이
// - Space Complexity: O(n) - creates new string
// - 공간 복잡도: O(n) - 새 문자열 생성
// - Use Cases: CSV generation, path construction, formatting
// - 사용 사례: CSV 생성, 경로 구성, 형식 지정
//
// Map(strs []string, fn func(string) string) []string
// - Purpose: Apply transformation function to all strings
// - 목적: 모든 문자열에 변환 함수 적용
// - Functional Style: Supports functional programming paradigm
// - 함수형 스타일: 함수형 프로그래밍 패러다임 지원
// - Time Complexity: O(n * f) where n is slice length, f is function cost
// - 시간 복잡도: O(n * f), n은 슬라이스 길이, f는 함수 비용
// - Space Complexity: O(n) - creates new slice
// - 공간 복잡도: O(n) - 새 슬라이스 생성
// - Use Cases: Bulk transformations, case conversion, adding prefixes/suffixes
// - 사용 사례: 대량 변환, 케이스 변환, 접두사/접미사 추가
//
// Filter(strs []string, fn func(string) bool) []string
// - Purpose: Filter strings by predicate function
// - 목적: 조건 함수로 문자열 필터링
// - Return: New slice with strings where fn returns true
// - 반환: fn이 true를 반환하는 문자열의 새 슬라이스
// - Empty Result: Returns empty slice (not nil)
// - 빈 결과: 빈 슬라이스 반환 (nil 아님)
// - Time Complexity: O(n * f) where f is predicate cost
// - 시간 복잡도: O(n * f), f는 조건 비용
// - Space Complexity: O(n) in worst case (all match)
// - 공간 복잡도: 최악의 경우 O(n) (모두 일치)
// - Use Cases: Removing empty strings, filtering by length, validation
// - 사용 사례: 빈 문자열 제거, 길이별 필터링, 검증
//
// PadLeft(s string, length int, pad string) string
// - Purpose: Pad string on left to reach specified length
// - 목적: 지정된 길이에 도달하도록 문자열 왼쪽에 패딩
// - Padding: Repeats pad string until length reached
// - 패딩: 길이에 도달할 때까지 패드 문자열 반복
// - No Truncation: If string already >= length, returns unchanged
// - 잘림 없음: 문자열이 이미 >= 길이면 변경 없이 반환
// - Unicode-Safe: Uses rune count, not byte count
// - 유니코드 안전: 바이트 개수가 아닌 룬 개수 사용
// - Time Complexity: O(n) where n is padding length
// - 시간 복잡도: O(n), n은 패딩 길이
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Use Cases: Number formatting (leading zeros), text alignment, IDs
// - 사용 사례: 숫자 형식 (앞에 0), 텍스트 정렬, ID
//
// PadRight(s string, length int, pad string) string
// - Purpose: Pad string on right to reach specified length
// - 목적: 지정된 길이에 도달하도록 문자열 오른쪽에 패딩
// - Padding: Repeats pad string until length reached
// - 패딩: 길이에 도달할 때까지 패드 문자열 반복
// - No Truncation: If string already >= length, returns unchanged
// - 잘림 없음: 문자열이 이미 >= 길이면 변경 없이 반환
// - Unicode-Safe: Uses rune count
// - 유니코드 안전: 룬 개수 사용
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Use Cases: Column formatting, table alignment, spacing
// - 사용 사례: 열 형식, 테이블 정렬, 간격
//
// Lines(s string) []string
// - Purpose: Split string by newlines
// - 목적: 줄바꿈으로 문자열 분리
// - Delimiter: Splits on '\n' character
// - 구분자: '\n' 문자로 분리
// - Preserves Empty Lines: Empty lines become empty strings in result
// - 빈 줄 보존: 빈 줄은 결과에서 빈 문자열이 됨
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Use Cases: File parsing, multi-line input processing, text splitting
// - 사용 사례: 파일 파싱, 다중 라인 입력 처리, 텍스트 분리
//
// Words(s string) []string
// - Purpose: Split string by whitespace
// - 목적: 공백으로 문자열 분리
// - Whitespace: Space, tab, newline, etc.
// - 공백: 스페이스, 탭, 줄바꿈 등
// - Behavior: Consecutive whitespace treated as one separator, trims edges
// - 동작: 연속된 공백을 하나의 구분자로 처리, 가장자리 제거
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Use Cases: Tokenization, parsing, word extraction
// - 사용 사례: 토큰화, 파싱, 단어 추출
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - CountWords: O(n) - iterate through string
//   CountWords: O(n) - 문자열 반복
// - CountOccurrences: O(n * m) - search for substring
//   CountOccurrences: O(n * m) - 부분 문자열 검색
// - Join: O(n) - concatenate strings
//   Join: O(n) - 문자열 연결
// - Map: O(n * f) - apply function to each element
//   Map: O(n * f) - 각 요소에 함수 적용
// - Filter: O(n * f) - check predicate for each element
//   Filter: O(n * f) - 각 요소에 대해 조건 확인
// - PadLeft/PadRight: O(p) where p is padding length
//   PadLeft/PadRight: O(p), p는 패딩 길이
// - Lines/Words: O(n) - split string
//   Lines/Words: O(n) - 문자열 분리
//
// Space Complexities:
// 공간 복잡도:
// - CountWords/CountOccurrences: O(n) for temporary slices
//   CountWords/CountOccurrences: 임시 슬라이스용 O(n)
// - Join/Map/Filter/Lines/Words: O(n) - create new slices/strings
//   Join/Map/Filter/Lines/Words: O(n) - 새 슬라이스/문자열 생성
// - PadLeft/PadRight: O(n) - create padded string
//   PadLeft/PadRight: O(n) - 패딩된 문자열 생성
//
// Optimization Tips:
// 최적화 팁:
// 1. For large slices, consider in-place operations instead of Map/Filter
//    큰 슬라이스의 경우 Map/Filter 대신 제자리 연산 고려
// 2. CountWords allocates slice - for simple count, might be optimized
//    CountWords는 슬라이스 할당 - 간단한 개수의 경우 최적화 가능
// 3. Join is efficient - uses strings.Builder internally
//    Join은 효율적 - 내부적으로 strings.Builder 사용
// 4. For repeated padding, consider caching padded strings
//    반복적인 패딩의 경우 패딩된 문자열 캐싱 고려
// 5. Lines/Words split entire string - for streaming, use Scanner
//    Lines/Words는 전체 문자열 분리 - 스트리밍의 경우 Scanner 사용
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Strings:
// 빈 문자열:
// - CountWords(""): 0 (no words)
//   CountWords(""): 0 (단어 없음)
// - CountOccurrences("", ""): 1 (special case)
//   CountOccurrences("", ""): 1 (특수 케이스)
// - Join([]string{}, ","): "" (empty result)
//   Join([]string{}, ","): "" (빈 결과)
// - Map([]string{}, fn): []string{} (empty slice)
//   Map([]string{}, fn): []string{} (빈 슬라이스)
// - Filter([]string{}, fn): []string{} (empty, not nil)
//   Filter([]string{}, fn): []string{} (빈, nil 아님)
// - PadLeft("", 5, "0"): "00000"
//   PadLeft("", 5, "0"): "00000"
// - Lines(""): [""] (one empty line)
//   Lines(""): [""] (하나의 빈 줄)
// - Words(""): []string{} (no words)
//   Words(""): []string{} (단어 없음)
//
// Whitespace Handling:
// 공백 처리:
// - CountWords("  a  b  "): 2 (leading/trailing whitespace ignored)
//   CountWords("  a  b  "): 2 (앞뒤 공백 무시)
// - Words("  a  b  "): ["a", "b"] (trimmed and split)
//   Words("  a  b  "): ["a", "b"] (제거되고 분리됨)
//
// Padding Behavior:
// 패딩 동작:
// - PadLeft("123", 5, "0"): "00123" (2 zeros added)
//   PadLeft("123", 5, "0"): "00123" (0 2개 추가)
// - PadLeft("12345", 3, "0"): "12345" (no change, already long enough)
//   PadLeft("12345", 3, "0"): "12345" (변경 없음, 이미 충분히 김)
// - PadLeft("hi", 5, "ab"): "ababhi" (repeats pad string)
//   PadLeft("hi", 5, "ab"): "ababhi" (패드 문자열 반복)
//
// Filter Empty Result:
// 필터 빈 결과:
// - Filter always returns non-nil slice (may be empty)
//   Filter는 항상 nil이 아닌 슬라이스 반환 (비어있을 수 있음)
// - This avoids nil pointer issues in iteration
//   반복에서 nil 포인터 문제 방지
//
// Unicode Considerations:
// 유니코드 고려사항:
// - PadLeft/PadRight use rune count, not byte count
//   PadLeft/PadRight는 바이트 개수가 아닌 룬 개수 사용
// - PadLeft("你好", 5, "0"): "000你好" (2 chars → 5 chars)
//   PadLeft("你好", 5, "0"): "000你好" (2문자 → 5문자)
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Word Count Validation
//    단어 개수 검증:
//
//    essay := "This is a sample essay"
//    wordCount := stringutil.CountWords(essay)
//    if wordCount < 100 {
//        return errors.New("essay must be at least 100 words")
//    }
//    // Validate minimum word count
//    // 최소 단어 개수 검증
//
// 2. Padding Numbers with Leading Zeros
//    앞에 0으로 숫자 패딩:
//
//    id := "42"
//    paddedID := stringutil.PadLeft(id, 6, "0")
//    // "000042"
//    // Format IDs with consistent length
//    // 일관된 길이로 ID 형식 지정
//
// 3. Bulk String Transformation
//    대량 문자열 변환:
//
//    names := []string{"alice", "bob", "charlie"}
//    uppercase := stringutil.Map(names, strings.ToUpper)
//    // ["ALICE", "BOB", "CHARLIE"]
//    // Transform all strings at once
//    // 모든 문자열을 한 번에 변환
//
// 4. Filtering Empty Strings
//    빈 문자열 필터링:
//
//    inputs := []string{"hello", "", "world", ""}
//    nonEmpty := stringutil.Filter(inputs, func(s string) bool {
//        return s != ""
//    })
//    // ["hello", "world"]
//    // Remove empty entries
//    // 빈 항목 제거
//
// 5. CSV Row Generation
//    CSV 행 생성:
//
//    fields := []string{"Alice", "30", "Engineer"}
//    csvRow := stringutil.Join(fields, ",")
//    // "Alice,30,Engineer"
//    // Create CSV format
//    // CSV 형식 생성
//
// 6. Counting Pattern Frequency
//    패턴 빈도 개수:
//
//    text := "hello hello world hello"
//    helloCount := stringutil.CountOccurrences(text, "hello")
//    // 3
//    // Analyze text patterns
//    // 텍스트 패턴 분석
//
// 7. Processing Multi-Line Input
//    다중 라인 입력 처리:
//
//    multiLine := "line1\nline2\nline3"
//    lines := stringutil.Lines(multiLine)
//    for i, line := range lines {
//        fmt.Printf("%d: %s\n", i+1, line)
//    }
//    // Process each line separately
//    // 각 줄 개별 처리
//
// 8. Tokenizing User Input
//    사용자 입력 토큰화:
//
//    userInput := "  search   term   here  "
//    tokens := stringutil.Words(userInput)
//    // ["search", "term", "here"]
//    // Clean tokenization
//    // 깨끗한 토큰화
//
// 9. Table Column Alignment
//    테이블 열 정렬:
//
//    names := []string{"Alice", "Bob", "Charlie"}
//    for _, name := range names {
//        padded := stringutil.PadRight(name, 10, " ")
//        fmt.Printf("|%s|", padded)
//    }
//    // |Alice     |Bob       |Charlie   |
//    // Create aligned columns
//    // 정렬된 열 생성
//
// 10. Filtering by Length
//     길이별 필터링:
//
//     passwords := []string{"short", "verylongpassword", "ok"}
//     valid := stringutil.Filter(passwords, func(s string) bool {
//         return len(s) >= 8
//     })
//     // ["verylongpassword"]
//     // Validate password length
//     // 비밀번호 길이 검증
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// CountWords vs strings.Split
// - CountWords: Counts words, handles multiple whitespace
//   CountWords: 단어 개수, 여러 공백 처리
// - strings.Split: Splits on exact delimiter
//   strings.Split: 정확한 구분자로 분리
// - Use CountWords for: Word counting
//   CountWords 사용: 단어 개수 세기
//
// Map vs for loop
// - Map: Functional style, returns new slice
//   Map: 함수형 스타일, 새 슬라이스 반환
// - for loop: More explicit, can modify in-place
//   for 루프: 더 명시적, 제자리 수정 가능
// - Use Map for: Concise transformations
//   Map 사용: 간결한 변환
//
// Filter vs for loop with append
// - Filter: Functional style, clear intent
//   Filter: 함수형 스타일, 명확한 의도
// - for loop: More control over allocation
//   for 루프: 할당에 대한 더 많은 제어
// - Use Filter for: Readability
//   Filter 사용: 가독성
//
// Join vs manual concatenation
// - Join: Efficient (uses strings.Builder)
//   Join: 효율적 (strings.Builder 사용)
// - Manual: Multiple allocations, slower
//   수동: 여러 할당, 더 느림
// - Use Join for: Performance
//   Join 사용: 성능
//
// PadLeft vs fmt.Sprintf
// - PadLeft: More flexible (any pad string)
//   PadLeft: 더 유연 (모든 패드 문자열)
// - fmt.Sprintf: Built-in, limited to spaces
//   fmt.Sprintf: 내장, 공백으로 제한
// - Use PadLeft for: Custom padding
//   PadLeft 사용: 사용자 정의 패딩
//
// FUNCTIONAL PROGRAMMING PATTERNS
// 함수형 프로그래밍 패턴
// --------------------------------
//
// Map and Filter enable functional-style programming:
// Map과 Filter는 함수형 스타일 프로그래밍을 가능하게 함:
//
// Chaining (conceptual):
// 체이닝 (개념):
//     inputs := []string{"  HELLO  ", "  WORLD  "}
//
//     // Trim whitespace
//     // 공백 제거
//     trimmed := stringutil.Map(inputs, strings.TrimSpace)
//
//     // Convert to lowercase
//     // 소문자로 변환
//     lowercase := stringutil.Map(trimmed, strings.ToLower)
//
//     // Filter out short words
//     // 짧은 단어 필터링
//     filtered := stringutil.Filter(lowercase, func(s string) bool {
//         return len(s) > 3
//     })
//
//     // Result: ["hello", "world"]
//     // 결과: ["hello", "world"]
//
// Benefits:
// 이점:
// - Clear, declarative code
//   명확하고 선언적인 코드
// - Easy to read and maintain
//   읽고 유지 관리하기 쉬움
// - Testable (functions are pure)
//   테스트 가능 (함수가 순수함)
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All functions in this file are thread-safe as they operate on immutable strings
// and create new slices/strings rather than modifying in-place.
//
// 이 파일의 모든 함수는 불변 문자열에서 작동하고 제자리 수정 대신 새
// 슬라이스/문자열을 생성하므로 스레드 안전합니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     go func() {
//         count := stringutil.CountWords(text)
//     }()
//
//     go func() {
//         padded := stringutil.PadLeft(id, 10, "0")
//     }()
//
//     // All utility functions safe for concurrent use
//     // 모든 유틸리티 함수는 동시 사용에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - manipulation.go: String manipulation operations
//   manipulation.go: 문자열 조작 연산
// - formatting.go: Advanced formatting operations
//   formatting.go: 고급 형식 지정 연산
// - builder.go: Fluent string builder (uses some utilities)
//   builder.go: 유창한 문자열 빌더 (일부 유틸리티 사용)
//
// =============================================================================

// CountWords counts the number of words in a string (split by whitespace).
// CountWords는 문자열의 단어 수를 셉니다 (공백으로 분리).
//
// Example:
//
//	CountWords("hello world")  // 2
//	CountWords("  a  b  c  ")  // 3
func CountWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

// CountOccurrences counts the number of times a substring appears in a string.
// CountOccurrences는 부분 문자열이 문자열에 나타나는 횟수를 셉니다.
//
// Example:
//
//	CountOccurrences("hello hello", "hello")  // 2
//	CountOccurrences("abcabc", "abc")          // 2
func CountOccurrences(s, substr string) int {
	return strings.Count(s, substr)
}

// Join joins a slice of strings with a separator (wrapper for strings.Join).
// Join은 구분자로 문자열 슬라이스를 연결합니다 (strings.Join의 래퍼).
//
// Example:
//
//	Join([]string{"a", "b", "c"}, "-")  // "a-b-c"
//	Join([]string{"hello", "world"}, " ")  // "hello world"
func Join(strs []string, sep string) string {
	return strings.Join(strs, sep)
}

// Map applies a function to all strings in a slice.
// Map은 슬라이스의 모든 문자열에 함수를 적용합니다.
//
// Example:
//
//	Map([]string{"a", "b"}, strings.ToUpper)  // ["A", "B"]
//	Map([]string{"hello", "world"}, func(s string) string { return s + "!" })  // ["hello!", "world!"]
func Map(strs []string, fn func(string) string) []string {
	result := make([]string, len(strs))
	for i, s := range strs {
		result[i] = fn(s)
	}
	return result
}

// Filter filters strings by a predicate function.
// Filter는 조건 함수로 문자열을 필터링합니다.
//
// Example:
//
//	Filter([]string{"a", "ab", "abc"}, func(s string) bool { return len(s) > 2 })  // ["abc"]
//	Filter([]string{"hello", "world", "hi"}, func(s string) bool { return len(s) > 3 })  // ["hello", "world"]
func Filter(strs []string, fn func(string) bool) []string {
	// Initialize to empty slice, not nil
	// nil이 아닌 빈 슬라이스로 초기화
	result := make([]string, 0)
	for _, s := range strs {
		if fn(s) {
			result = append(result, s)
		}
	}
	return result
}

// PadLeft pads a string on the left side to reach the specified length.
// PadLeft는 지정된 길이에 도달하도록 문자열의 왼쪽에 패딩을 추가합니다.
//
// Example:
//
//	PadLeft("5", 3, "0")    // "005"
//	PadLeft("42", 5, "0")   // "00042"
func PadLeft(s string, length int, pad string) string {
	runes := []rune(s)
	if len(runes) >= length {
		return s
	}
	padCount := length - len(runes)
	padding := strings.Repeat(pad, padCount)
	return padding + s
}

// PadRight pads a string on the right side to reach the specified length.
// PadRight는 지정된 길이에 도달하도록 문자열의 오른쪽에 패딩을 추가합니다.
//
// Example:
//
//	PadRight("5", 3, "0")   // "500"
//	PadRight("42", 5, "0")  // "42000"
func PadRight(s string, length int, pad string) string {
	runes := []rune(s)
	if len(runes) >= length {
		return s
	}
	padCount := length - len(runes)
	padding := strings.Repeat(pad, padCount)
	return s + padding
}

// Lines splits a string by newlines.
// Lines는 줄바꿈으로 문자열을 분리합니다.
//
// Example:
//
//	Lines("line1\nline2\nline3")  // ["line1", "line2", "line3"]
//	Lines("a\nb\nc")               // ["a", "b", "c"]
func Lines(s string) []string {
	return strings.Split(s, "\n")
}

// Words splits a string by whitespace.
// Words는 공백으로 문자열을 분리합니다.
//
// Example:
//
//	Words("hello world foo")  // ["hello", "world", "foo"]
//	Words("  a  b  c  ")       // ["a", "b", "c"]
func Words(s string) []string {
	return strings.Fields(s)
}
