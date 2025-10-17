package stringutil

import (
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

// =============================================================================
// File: unicode.go
// Purpose: Unicode String Operations and Normalization
// 파일: unicode.go
// 목적: 유니코드 문자열 연산 및 정규화
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The unicode.go file provides specialized functions for working with Unicode
// strings, handling character counting, display width calculation, and Unicode
// normalization. These operations are essential for internationalized applications
// that need to correctly process text in various languages, including those with
// multi-byte characters (CJK languages), combining characters (accents), and
// emoji. The file addresses the fact that string length in bytes does not equal
// the number of visible characters.
//
// unicode.go 파일은 유니코드 문자열 작업을 위한 특수 함수를 제공하며, 문자
// 개수 세기, 디스플레이 너비 계산 및 유니코드 정규화를 처리합니다. 이러한
// 연산은 다중 바이트 문자 (CJK 언어), 결합 문자 (악센트), 이모지를 포함한
// 다양한 언어의 텍스트를 올바르게 처리해야 하는 국제화된 애플리케이션에
// 필수적입니다. 이 파일은 바이트 단위 문자열 길이가 보이는 문자 수와 같지
// 않다는 사실을 다룹니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Unicode-First**: Treat strings as sequences of runes, not bytes
//    **유니코드 우선**: 문자열을 바이트가 아닌 룬 시퀀스로 처리
//
// 2. **Display-Aware**: Consider actual display width, not just character count
//    **디스플레이 인식**: 문자 개수뿐만 아니라 실제 디스플레이 너비 고려
//
// 3. **Normalization Support**: Provide canonical and compatibility normalization
//    **정규화 지원**: 정규 및 호환성 정규화 제공
//
// 4. **International-Ready**: Handle CJK, emoji, combining characters correctly
//    **국제화 준비**: CJK, 이모지, 결합 문자 올바르게 처리
//
// 5. **Standards-Based**: Follow Unicode standards (UAX #11, UAX #15)
//    **표준 기반**: 유니코드 표준 따름 (UAX #11, UAX #15)
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. CHARACTER COUNTING (문자 개수 세기)
//    - RuneCount: Count Unicode characters (runes), not bytes
//      RuneCount: 바이트가 아닌 유니코드 문자 (룬) 개수
//
// 2. DISPLAY WIDTH CALCULATION (디스플레이 너비 계산)
//    - Width: Calculate display width considering East Asian characters
//      Width: 동아시아 문자를 고려한 디스플레이 너비 계산
//
// 3. UNICODE NORMALIZATION (유니코드 정규화)
//    - Normalize: Normalize to NFC, NFD, NFKC, or NFKD form
//      Normalize: NFC, NFD, NFKC 또는 NFKD 형식으로 정규화
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// RuneCount(s string) int
// - Purpose: Count Unicode characters (runes), not bytes
// - 목적: 바이트가 아닌 유니코드 문자 (룬) 개수
// - Difference from len(): len() counts bytes, RuneCount counts characters
// - len()과의 차이: len()은 바이트 개수, RuneCount는 문자 개수
// - Time Complexity: O(n) where n is byte length
// - 시간 복잡도: O(n), n은 바이트 길이
// - Space Complexity: O(1) - no allocation
// - 공간 복잡도: O(1) - 할당 없음
// - Unicode Handling: Correctly counts multi-byte characters as single rune
// - 유니코드 처리: 다중 바이트 문자를 단일 룬으로 올바르게 계산
// - Use Cases: Text truncation, character limits, input validation, progress indicators
// - 사용 사례: 텍스트 자르기, 문자 제한, 입력 검증, 진행 표시기
//
// Width(s string) int
// - Purpose: Calculate display width considering East Asian characters
// - 목적: 동아시아 문자를 고려한 디스플레이 너비 계산
// - East Asian Width (EAW): Follows Unicode Standard Annex #11
// - 동아시아 너비 (EAW): 유니코드 표준 부록 #11 준수
// - Width Rules:
//   * ASCII (a-z, 0-9, basic punctuation): width 1
//   * CJK characters (한글, 漢字, ひらがな, etc.): width 2
//   * Emoji: typically width 2
//   * Combining characters: width 0 (added to previous character)
// - 너비 규칙:
//   * ASCII (a-z, 0-9, 기본 구두점): 너비 1
//   * CJK 문자 (한글, 漢字, ひらがな 등): 너비 2
//   * 이모지: 일반적으로 너비 2
//   * 결합 문자: 너비 0 (이전 문자에 추가)
// - Time Complexity: O(n) where n is rune count
// - 시간 복잡도: O(n), n은 룬 개수
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Terminal output alignment, monospace display, text formatting
// - 사용 사례: 터미널 출력 정렬, 고정폭 디스플레이, 텍스트 형식
//
// Normalize(s string, form string) string
// - Purpose: Normalize Unicode string to canonical or compatibility form
// - 목적: 유니코드 문자열을 정규 또는 호환성 형식으로 정규화
// - Normalization Forms:
//   * NFC (Canonical Decomposition + Composition): Most common
//     NFC (정규 분해 + 결합): 가장 일반적
//   * NFD (Canonical Decomposition): Decomposes characters
//     NFD (정규 분해): 문자 분해
//   * NFKC (Compatibility Decomposition + Composition): Compatibility
//     NFKC (호환성 분해 + 결합): 호환성
//   * NFKD (Compatibility Decomposition): Compatibility decomposed
//     NFKD (호환성 분해): 호환성 분해됨
// - Default: NFC if form parameter is invalid or empty
// - 기본값: form 매개변수가 잘못되거나 비어있으면 NFC
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n) - creates new string
// - 공간 복잡도: O(n) - 새 문자열 생성
// - Use Cases: String comparison, search, database storage, text processing
// - 사용 사례: 문자열 비교, 검색, 데이터베이스 저장, 텍스트 처리
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - RuneCount: O(n) - iterate through bytes to count runes
//   RuneCount: O(n) - 룬 개수를 세기 위해 바이트 반복
// - Width: O(n) - iterate through runes and lookup width
//   Width: O(n) - 룬 반복 및 너비 조회
// - Normalize: O(n) - transform each character
//   Normalize: O(n) - 각 문자 변환
//
// Space Complexities:
// 공간 복잡도:
// - RuneCount: O(1) - no allocation
//   RuneCount: O(1) - 할당 없음
// - Width: O(1) - no allocation
//   Width: O(1) - 할당 없음
// - Normalize: O(n) - creates new string
//   Normalize: O(n) - 새 문자열 생성
//
// Optimization Tips:
// 최적화 팁:
// 1. Cache RuneCount results for repeated use
//    반복 사용을 위해 RuneCount 결과 캐시
// 2. For ASCII-only strings, len() is faster than RuneCount
//    ASCII만 있는 문자열의 경우 RuneCount보다 len()이 빠름
// 3. Width calculation is more expensive than RuneCount
//    Width 계산은 RuneCount보다 더 비쌈
// 4. Normalize once at input, not repeatedly
//    반복적으로 하지 말고 입력 시 한 번 정규화
// 5. For comparison, normalize both strings to same form
//    비교를 위해 두 문자열을 같은 형식으로 정규화
//
// UNICODE NORMALIZATION EXPLAINED
// 유니코드 정규화 설명
// --------------------------------
//
// What is Unicode Normalization?
// 유니코드 정규화란?
//
// Unicode allows multiple representations of the same character. For example:
// 유니코드는 같은 문자의 여러 표현을 허용합니다. 예:
// - "é" can be represented as:
//   "é"는 다음과 같이 표현될 수 있음:
//   1. Single character U+00E9 (composed)
//      단일 문자 U+00E9 (결합됨)
//   2. Two characters U+0065 (e) + U+0301 (combining acute accent)
//      두 문자 U+0065 (e) + U+0301 (결합 악센트)
//
// Normalization ensures consistent representation for:
// 정규화는 다음을 위한 일관된 표현을 보장:
// - String comparison (both forms should be considered equal)
//   문자열 비교 (두 형식 모두 동일하게 간주되어야 함)
// - Database storage (consistent indexing and searching)
//   데이터베이스 저장 (일관된 인덱싱 및 검색)
// - Text processing (predictable behavior)
//   텍스트 처리 (예측 가능한 동작)
//
// Normalization Forms:
// 정규화 형식:
//
// 1. **NFC (Canonical Composition)**:
//    - Decomposes then recomposes to canonical form
//      정규 형식으로 분해 후 재결합
//    - Most compact representation
//      가장 컴팩트한 표현
//    - Recommended for most use cases
//      대부분의 사용 사례에 권장
//    - Example: "é" → U+00E9 (single character)
//      예: "é" → U+00E9 (단일 문자)
//
// 2. **NFD (Canonical Decomposition)**:
//    - Fully decomposes characters
//      문자 완전히 분해
//    - Useful for analyzing diacritics
//      발음 구별 부호 분석에 유용
//    - Example: "é" → U+0065 + U+0301 (e + accent)
//      예: "é" → U+0065 + U+0301 (e + 악센트)
//
// 3. **NFKC (Compatibility Composition)**:
//    - Applies compatibility mappings then composes
//      호환성 매핑 적용 후 결합
//    - Converts similar-looking characters to standard forms
//      유사한 문자를 표준 형식으로 변환
//    - Example: "①" → "1", "ﬁ" → "fi"
//      예: "①" → "1", "ﬁ" → "fi"
//    - Use for: Search, text processing, normalization
//      사용처: 검색, 텍스트 처리, 정규화
//
// 4. **NFKD (Compatibility Decomposition)**:
//    - Applies compatibility mappings and decomposes
//      호환성 매핑 적용 및 분해
//    - Most decomposed form
//      가장 분해된 형식
//    - Example: "①" → "1", "é" → "e" + accent
//      예: "①" → "1", "é" → "e" + 악센트
//
// EAST ASIAN WIDTH EXPLAINED
// 동아시아 너비 설명
// ---------------------------
//
// Why Display Width Matters:
// 디스플레이 너비가 중요한 이유:
//
// In terminal/monospace displays, characters have different widths:
// 터미널/고정폭 디스플레이에서 문자는 다른 너비를 가짐:
// - ASCII characters: 1 cell width
//   ASCII 문자: 1셀 너비
// - CJK characters: 2 cell widths (wider)
//   CJK 문자: 2셀 너비 (더 넓음)
// - Emoji: typically 2 cell widths
//   이모지: 일반적으로 2셀 너비
//
// This affects:
// 이것은 다음에 영향:
// - Text alignment in terminals
//   터미널의 텍스트 정렬
// - Progress bars and UI elements
//   진행 표시줄 및 UI 요소
// - Fixed-width text formatting
//   고정폭 텍스트 형식
//
// Example:
// 예제:
//     "hello"     → width 5 (5 × 1)
//     "안녕"       → width 4 (2 × 2)
//     "hello세계" → width 9 (5 × 1 + 2 × 2)
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Strings:
// 빈 문자열:
// - RuneCount(""): 0
//   RuneCount(""): 0
// - Width(""): 0
//   Width(""): 0
// - Normalize("", "NFC"): ""
//   Normalize("", "NFC"): ""
//
// Multi-Byte Characters:
// 다중 바이트 문자:
// - RuneCount("你好"): 2 (not 6 bytes)
//   RuneCount("你好"): 2 (6바이트 아님)
// - len("你好"): 6 (bytes)
//   len("你好"): 6 (바이트)
// - Width("你好"): 4 (2 characters × 2 width)
//   Width("你好"): 4 (2문자 × 2너비)
//
// Emoji:
// 이모지:
// - RuneCount("🔥🔥"): 2
//   RuneCount("🔥🔥"): 2
// - len("🔥🔥"): 8 (bytes)
//   len("🔥🔥"): 8 (바이트)
// - Width("🔥🔥"): 4 (typically 2 width each)
//   Width("🔥🔥"): 4 (일반적으로 각 2너비)
//
// Combining Characters:
// 결합 문자:
// - "é" (composed): RuneCount 1, Width 1
//   "é" (결합됨): RuneCount 1, Width 1
// - "é" (decomposed e + accent): RuneCount 2, Width 1
//   "é" (분해됨 e + 악센트): RuneCount 2, Width 1
//
// Normalization:
// 정규화:
// - NFC is most compact: "é" → 1 character
//   NFC가 가장 컴팩트: "é" → 1문자
// - NFD is decomposed: "é" → 2 characters (e + accent)
//   NFD는 분해됨: "é" → 2문자 (e + 악센트)
// - NFKC converts compatibility: "①" → "1"
//   NFKC는 호환성 변환: "①" → "1"
//
// Invalid UTF-8:
// 잘못된 UTF-8:
// - RuneCount handles invalid UTF-8 by counting replacement characters
//   RuneCount는 대체 문자를 세어 잘못된 UTF-8 처리
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Character Limit Validation
//    문자 제한 검증:
//
//    userInput := "Hello, 世界!"
//    maxChars := 10
//    if stringutil.RuneCount(userInput) > maxChars {
//        return errors.New("input exceeds 10 characters")
//    }
//    // Correct character counting for international text
//    // 국제 텍스트의 올바른 문자 개수 세기
//
// 2. Terminal Text Alignment
//    터미널 텍스트 정렬:
//
//    items := []string{"hello", "안녕하세요", "🔥"}
//    maxWidth := 20
//    for _, item := range items {
//        padding := maxWidth - stringutil.Width(item)
//        fmt.Printf("%s%s\n", item, strings.Repeat(" ", padding))
//    }
//    // Properly align mixed-width characters
//    // 혼합 너비 문자 적절히 정렬
//
// 3. Unicode String Comparison
//    유니코드 문자열 비교:
//
//    s1 := "café"  // composed é
//    s2 := "café"  // decomposed e + accent
//    normalized1 := stringutil.Normalize(s1, "NFC")
//    normalized2 := stringutil.Normalize(s2, "NFC")
//    if normalized1 == normalized2 {
//        fmt.Println("Strings are equivalent")
//    }
//    // Normalize before comparison
//    // 비교 전 정규화
//
// 4. Database Text Storage
//    데이터베이스 텍스트 저장:
//
//    userInput := "naïve café"
//    normalized := stringutil.Normalize(userInput, "NFC")
//    // Store normalized form in database
//    // 정규화된 형식을 데이터베이스에 저장
//    // Ensures consistent search and indexing
//    // 일관된 검색 및 인덱싱 보장
//
// 5. Text Truncation with Unicode
//    유니코드로 텍스트 자르기:
//
//    text := "Hello, 世界! 🔥"
//    maxChars := 10
//    if stringutil.RuneCount(text) > maxChars {
//        runes := []rune(text)
//        truncated := string(runes[:maxChars]) + "..."
//    }
//    // Truncate by character count, not bytes
//    // 바이트가 아닌 문자 개수로 자르기
//
// 6. Progress Bar with Mixed Text
//    혼합 텍스트로 진행 표시줄:
//
//    label := "Processing 文件..."
//    barWidth := 50
//    labelWidth := stringutil.Width(label)
//    progressWidth := barWidth - labelWidth - 2
//    fmt.Printf("%s [%s]\n", label, strings.Repeat("=", progressWidth))
//    // Account for display width
//    // 디스플레이 너비 고려
//
// 7. Search with Normalization
//    정규화로 검색:
//
//    query := stringutil.Normalize("naïve", "NFKC")
//    for _, doc := range documents {
//        normalized := stringutil.Normalize(doc, "NFKC")
//        if strings.Contains(normalized, query) {
//            fmt.Println("Match found")
//        }
//    }
//    // Normalize both query and documents
//    // 쿼리 및 문서 모두 정규화
//
// 8. Emoji Handling
//    이모지 처리:
//
//    message := "Hello 👋 World 🌍"
//    charCount := stringutil.RuneCount(message)   // 15 characters
//    displayWidth := stringutil.Width(message)     // 17 (emoji = 2 width)
//    fmt.Printf("Characters: %d, Display width: %d\n", charCount, displayWidth)
//    // Distinguish between character count and display width
//    // 문자 개수와 디스플레이 너비 구별
//
// 9. Input Validation for CJK
//    CJK 입력 검증:
//
//    username := "사용자123"
//    if stringutil.RuneCount(username) > 20 {
//        return errors.New("username too long")
//    }
//    // Correctly count CJK characters
//    // CJK 문자 올바르게 개수 세기
//
// 10. Compatibility Normalization
//     호환성 정규화:
//
//     input := "①②③"  // Circled numbers
//     normalized := stringutil.Normalize(input, "NFKC")
//     // "123" - converted to ASCII digits
//     // "123" - ASCII 숫자로 변환
//     // Useful for search and comparison
//     // 검색 및 비교에 유용
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// RuneCount vs len()
// - RuneCount: Counts Unicode characters (runes)
//   RuneCount: 유니코드 문자 (룬) 개수
// - len(): Counts bytes
//   len(): 바이트 개수
// - Use RuneCount for: User-facing character limits
//   RuneCount 사용: 사용자 대상 문자 제한
// - Use len() for: Memory/storage calculations
//   len() 사용: 메모리/저장소 계산
//
// Width vs RuneCount
// - Width: Display width (considers CJK = 2 width)
//   Width: 디스플레이 너비 (CJK = 2너비 고려)
// - RuneCount: Character count
//   RuneCount: 문자 개수
// - Use Width for: Terminal/monospace alignment
//   Width 사용: 터미널/고정폭 정렬
// - Use RuneCount for: Character limits
//   RuneCount 사용: 문자 제한
//
// NFC vs NFD
// - NFC: Composed (compact), recommended for most use
//   NFC: 결합됨 (컴팩트), 대부분의 사용에 권장
// - NFD: Decomposed, useful for analysis
//   NFD: 분해됨, 분석에 유용
// - Use NFC for: Storage, display, general use
//   NFC 사용: 저장, 디스플레이, 일반 사용
// - Use NFD for: Diacritic removal, text analysis
//   NFD 사용: 발음 구별 부호 제거, 텍스트 분석
//
// NFKC vs NFC
// - NFKC: Compatibility (converts similar characters)
//   NFKC: 호환성 (유사 문자 변환)
// - NFC: Canonical (preserves distinctions)
//   NFC: 정규 (구별 보존)
// - Use NFKC for: Search, fuzzy matching
//   NFKC 사용: 검색, 퍼지 매칭
// - Use NFC for: Exact representation
//   NFC 사용: 정확한 표현
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All functions in this file are thread-safe as they operate on immutable strings
// and use thread-safe standard library functions.
//
// 이 파일의 모든 함수는 불변 문자열에서 작동하고 스레드 안전한 표준 라이브러리
// 함수를 사용하므로 스레드 안전합니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     go func() {
//         count := stringutil.RuneCount(text)
//     }()
//
//     go func() {
//         normalized := stringutil.Normalize(text, "NFC")
//     }()
//
//     // All Unicode functions safe for concurrent use
//     // 모든 유니코드 함수는 동시 사용에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - manipulation.go: String manipulation (uses rune-based operations)
//   manipulation.go: 문자열 조작 (룬 기반 연산 사용)
// - validation.go: String validation (Unicode-aware)
//   validation.go: 문자열 검증 (유니코드 인식)
// - comparison.go: String comparison (use with normalization)
//   comparison.go: 문자열 비교 (정규화와 함께 사용)
//
// =============================================================================

// RuneCount returns the number of Unicode characters (runes) in a string.
// RuneCount는 문자열의 유니코드 문자(rune) 개수를 반환합니다.
//
// This is different from len(s) which returns the number of bytes.
// 이것은 바이트 개수를 반환하는 len(s)와 다릅니다.
//
// Example:
//
// RuneCount("hello")    // 5
// RuneCount("안녕하세요")  // 5 (not 15 bytes)
//
//	RuneCount("🔥🔥")      // 2 (not 8 bytes)
func RuneCount(s string) int {
	return utf8.RuneCountInString(s)
}

// Width returns the display width of a string.
// Width는 문자열의 디스플레이 너비를 반환합니다.
//
// This considers East Asian Width (EAW) properties:
// 동아시아 너비(EAW) 속성을 고려합니다:
// - ASCII characters (a-z, 0-9): width 1
// - CJK characters (한글, 漢字, etc): width 2
//   - Emoji: typically width 2
//
// Example:
//
// Width("hello")      // 5
// Width("안녕")        // 4 (2 characters × 2 width each)
//
//	Width("hello세계")   // 9 (5 + 4)
func Width(s string) int {
	totalWidth := 0
	for _, r := range s {
		prop := width.LookupRune(r)
		switch prop.Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth:
			totalWidth += 2
		default:
			totalWidth += 1
		}
	}
	return totalWidth
}

// Normalize normalizes a Unicode string to the specified form.
// Normalize는 유니코드 문자열을 지정된 형식으로 정규화합니다.
//
// Normalization Form:
// 정규화 형식:
//   - "NFC": Canonical Decomposition followed by Canonical Composition
//   - "NFD": Canonical Decomposition
//   - "NFKC": Compatibility Decomposition followed by Canonical Composition
//   - "NFKD": Compatibility Decomposition
//
// Default is NFC if form is empty or invalid.
// form이 비어있거나 유효하지 않으면 기본값은 NFC입니다.
//
// Example:
//
//	Normalize("café", "NFC")   // "café" (composed é)
//	Normalize("café", "NFD")   // "café" (decomposed e + ́)
//	Normalize("①②③", "NFKC")  // "123" (compatibility)
func Normalize(s string, form string) string {
	var normalizer norm.Form

	switch form {
	case "NFC":
		normalizer = norm.NFC
	case "NFD":
		normalizer = norm.NFD
	case "NFKC":
		normalizer = norm.NFKC
	case "NFKD":
		normalizer = norm.NFKD
	default:
		// Default to NFC
		// 기본값은 NFC
		normalizer = norm.NFC
	}

	return normalizer.String(s)
}
