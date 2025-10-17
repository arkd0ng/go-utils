package stringutil

import (
	"fmt"
	"strings"
)

// =============================================================================
// File: formatting.go
// Purpose: Advanced String Formatting and Display Operations
// 파일: formatting.go
// 목적: 고급 문자열 포맷팅 및 디스플레이 연산
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The formatting.go file provides advanced string formatting operations for
// displaying data in human-readable formats. These functions handle common
// formatting tasks like number formatting, byte size conversion, pluralization,
// masking sensitive data, text wrapping, and indentation. These utilities are
// particularly useful for CLI applications, logging, report generation, and
// user-facing displays where data needs to be presented in a clear and
// professional manner.
//
// formatting.go 파일은 사람이 읽을 수 있는 형식으로 데이터를 표시하기 위한
// 고급 문자열 포맷팅 연산을 제공합니다. 이러한 함수는 숫자 포맷팅, 바이트
// 크기 변환, 복수형화, 민감한 데이터 마스킹, 텍스트 래핑 및 들여쓰기와 같은
// 일반적인 포맷팅 작업을 처리합니다. 이러한 유틸리티는 CLI 애플리케이션,
// 로깅, 보고서 생성 및 데이터를 명확하고 전문적인 방식으로 제시해야 하는
// 사용자 대면 디스플레이에 특히 유용합니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Human-Readability**: Format data for human consumption, not machines
//    **사람 가독성**: 기계가 아닌 사람이 소비할 수 있도록 데이터 포맷
//
// 2. **Locale-Awareness**: Support different number/date formats (customizable separators)
//    **로케일 인식**: 다양한 숫자/날짜 형식 지원 (사용자 정의 구분자)
//
// 3. **Privacy Protection**: Provide safe masking for sensitive information
//    **개인정보 보호**: 민감한 정보를 위한 안전한 마스킹 제공
//
// 4. **Text Presentation**: Handle text layout for various display contexts
//    **텍스트 표현**: 다양한 디스플레이 컨텍스트를 위한 텍스트 레이아웃 처리
//
// 5. **Flexibility**: Allow customization of formatting parameters
//    **유연성**: 포맷팅 매개변수의 사용자 정의 허용
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. NUMBER FORMATTING (숫자 포맷팅)
//    - FormatNumber: Format integers with thousand separators
//      FormatNumber: 천 단위 구분자로 정수 포맷
//    - FormatBytes: Convert bytes to human-readable sizes (KB, MB, GB)
//      FormatBytes: 바이트를 사람이 읽을 수 있는 크기로 변환 (KB, MB, GB)
//
// 2. PLURALIZATION (복수형화)
//    - Pluralize: Return singular or plural form based on count
//      Pluralize: count에 따라 단수형 또는 복수형 반환
//    - FormatWithCount: Format count with pluralized noun
//      FormatWithCount: 복수형 명사와 함께 count 포맷
//
// 3. TRUNCATION AND ELLIPSIS (자르기 및 생략 부호)
//    - Ellipsis: Truncate with ellipsis in the middle
//      Ellipsis: 중간에 생략 부호로 자르기
//
// 4. MASKING (마스킹)
//    - Mask: Mask string with custom character
//      Mask: 사용자 정의 문자로 문자열 마스킹
//    - MaskEmail: Mask email address
//      MaskEmail: 이메일 주소 마스킹
//    - MaskCreditCard: Mask credit card number
//      MaskCreditCard: 신용카드 번호 마스킹
//
// 5. TEXT LAYOUT (텍스트 레이아웃)
//    - AddLineNumbers: Add line numbers to text
//      AddLineNumbers: 텍스트에 줄 번호 추가
//    - Indent: Indent each line with prefix
//      Indent: 접두사로 각 줄 들여쓰기
//    - Dedent: Remove common leading whitespace
//      Dedent: 공통 선행 공백 제거
//    - WrapText: Wrap text to specified width
//      WrapText: 지정된 너비로 텍스트 래핑
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// FormatNumber(n int, separator string) string
// - Purpose: Format integer with thousand separators
// - 목적: 천 단위 구분자로 정수 포맷
// - Time Complexity: O(log n) - proportional to number of digits
// - 시간 복잡도: O(log n) - 자릿수에 비례
// - Space Complexity: O(log n)
// - 공간 복잡도: O(log n)
// - Customization: User-defined separator (comma, period, space)
// - 사용자 정의: 사용자 정의 구분자 (쉼표, 마침표, 공백)
// - Use Cases: Financial reports, statistics, dashboards
// - 사용 사례: 재무 보고서, 통계, 대시보드
//
// FormatBytes(bytes int64) string
// - Purpose: Convert bytes to human-readable size with unit
// - 목적: 바이트를 단위와 함께 사람이 읽을 수 있는 크기로 변환
// - Time Complexity: O(log n)
// - 시간 복잡도: O(log n)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Units: B, KB, MB, GB, TB, PB, EB (binary: 1024 base)
// - 단위: B, KB, MB, GB, TB, PB, EB (바이너리: 1024 기반)
// - Precision: 1 decimal place
// - 정밀도: 소수점 1자리
// - Use Cases: File sizes, memory usage, disk space, download progress
// - 사용 사례: 파일 크기, 메모리 사용량, 디스크 공간, 다운로드 진행률
//
// Pluralize(count int, singular, plural string) string
// - Purpose: Return singular or plural form based on count
// - 목적: count에 따라 단수형 또는 복수형 반환
// - Time Complexity: O(1)
// - 시간 복잡도: O(1)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Rule: count == 1 → singular, otherwise → plural
// - 규칙: count == 1 → 단수형, 그 외 → 복수형
// - Use Cases: User messages, notifications, item counts
// - 사용 사례: 사용자 메시지, 알림, 항목 수
//
// FormatWithCount(count int, singular, plural string) string
// - Purpose: Format count with pluralized noun
// - 목적: 복수형 명사와 함께 count 포맷
// - Time Complexity: O(1)
// - 시간 복잡도: O(1)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Output: "N item" or "N items"
// - 출력: "N item" 또는 "N items"
// - Use Cases: List summaries, status messages, UI labels
// - 사용 사례: 목록 요약, 상태 메시지, UI 레이블
//
// Ellipsis(s string, maxLen int) string
// - Purpose: Truncate string with ellipsis in middle
// - 목적: 중간에 생략 부호로 문자열 자르기
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: Preserves start and end, replaces middle with "..."
// - 동작: 시작과 끝 유지, 중간을 "..."로 대체
// - Use Cases: Long filenames, URLs, preserving file extensions
// - 사용 사례: 긴 파일명, URL, 파일 확장자 유지
//
// Mask(s string, first, last int, maskChar string) string
// - Purpose: Mask string revealing only first and last characters
// - 목적: 처음과 마지막 문자만 표시하고 문자열 마스킹
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Customization: Configurable mask character and reveal length
// - 사용자 정의: 구성 가능한 마스크 문자 및 표시 길이
// - Use Cases: Privacy protection, sensitive data display
// - 사용 사례: 개인정보 보호, 민감한 데이터 표시
//
// MaskEmail(email string) string
// - Purpose: Mask email address preserving structure
// - 목적: 구조를 유지하면서 이메일 주소 마스킹
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Pattern: "j******e@example.com"
// - 패턴: "j******e@example.com"
// - Use Cases: Privacy displays, confirmation messages
// - 사용 사례: 개인정보 표시, 확인 메시지
//
// MaskCreditCard(card string) string
// - Purpose: Mask credit card number showing last 4 digits
// - 목적: 마지막 4자리만 표시하고 신용카드 번호 마스킹
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Preserves: Hyphens and spaces in original format
// - 유지: 원본 형식의 하이픈 및 공백
// - Use Cases: Payment confirmations, transaction history
// - 사용 사례: 결제 확인, 거래 내역
//
// AddLineNumbers(s string) string
// - Purpose: Add line numbers to each line of text
// - 목적: 텍스트의 각 줄에 줄 번호 추가
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Format: "1: line1\n2: line2"
// - 형식: "1: line1\n2: line2"
// - Use Cases: Code display, log viewing, debugging
// - 사용 사례: 코드 표시, 로그 보기, 디버깅
//
// Indent(s string, prefix string) string
// - Purpose: Indent each line with specified prefix
// - 목적: 지정된 접두사로 각 줄 들여쓰기
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Customization: Any prefix (spaces, tabs, custom strings)
// - 사용자 정의: 모든 접두사 (공백, 탭, 사용자 정의 문자열)
// - Use Cases: Code generation, nested text, quotations
// - 사용 사례: 코드 생성, 중첩 텍스트, 인용구
//
// Dedent(s string) string
// - Purpose: Remove common leading whitespace
// - 목적: 공통 선행 공백 제거
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: Finds minimum indent and removes from all lines
// - 동작: 최소 들여쓰기를 찾아 모든 줄에서 제거
// - Use Cases: Template processing, multi-line strings, code cleanup
// - 사용 사례: 템플릿 처리, 여러 줄 문자열, 코드 정리
//
// WrapText(s string, width int) string
// - Purpose: Wrap text to specified line width
// - 목적: 지정된 줄 너비로 텍스트 래핑
// - Time Complexity: O(n) where n is number of words
// - 시간 복잡도: O(n), n은 단어 수
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: Breaks at word boundaries, preserves words
// - 동작: 단어 경계에서 나눔, 단어 유지
// - Use Cases: Terminal output, text formatting, responsive displays
// - 사용 사례: 터미널 출력, 텍스트 포맷팅, 반응형 디스플레이
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - FormatNumber: O(log n) - iterate digits
//   FormatNumber: O(log n) - 자릿수 반복
// - FormatBytes: O(log n) - divide by unit until < 1024
//   FormatBytes: O(log n) - 1024보다 작을 때까지 단위로 나누기
// - Pluralize/FormatWithCount: O(1) - simple comparison
//   Pluralize/FormatWithCount: O(1) - 단순 비교
// - Ellipsis: O(n) - rune conversion + slicing
//   Ellipsis: O(n) - rune 변환 + 슬라이싱
// - Mask: O(n) - iterate characters
//   Mask: O(n) - 문자 반복
// - MaskEmail/MaskCreditCard: O(n) - string operations
//   MaskEmail/MaskCreditCard: O(n) - 문자열 연산
// - AddLineNumbers: O(n) - split lines + format
//   AddLineNumbers: O(n) - 줄 분리 + 포맷
// - Indent/Dedent: O(n) - line-by-line processing
//   Indent/Dedent: O(n) - 줄별 처리
// - WrapText: O(w) where w is number of words
//   WrapText: O(w), w는 단어 수
//
// Space Complexities:
// 공간 복잡도:
// - FormatNumber: O(log n) for result string
//   FormatNumber: O(log n), 결과 문자열용
// - FormatBytes: O(1) - fixed format
//   FormatBytes: O(1) - 고정 형식
// - Mask functions: O(n) - new string
//   마스크 함수: O(n) - 새 문자열
// - Text layout functions: O(n) - modified text
//   텍스트 레이아웃 함수: O(n) - 수정된 텍스트
//
// Optimization Tips:
// 최적화 팁:
// 1. Cache formatted numbers if displaying repeatedly
//    반복 표시 시 포맷된 숫자 캐시
// 2. Use strings.Builder for concatenation in loops
//    루프의 연결에 strings.Builder 사용
// 3. Pre-allocate builders with estimated capacity
//    예상 용량으로 빌더 미리 할당
// 4. For large text, process in chunks for better memory locality
//    대용량 텍스트의 경우 더 나은 메모리 지역성을 위해 청크로 처리
// 5. Consider lazy formatting for large datasets
//    대규모 데이터셋의 경우 지연 포맷팅 고려
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Strings:
// 빈 문자열:
// - FormatNumber: Handles negative numbers
//   FormatNumber: 음수 처리
// - Pluralize: count == 0 uses plural form
//   Pluralize: count == 0은 복수형 사용
// - Ellipsis("", 10) returns ""
//   Ellipsis("", 10)는 "" 반환
// - WrapText("", 10) returns ""
//   WrapText("", 10)는 "" 반환
//
// Small Numbers/Strings:
// 작은 숫자/문자열:
// - FormatNumber(123, ",") returns "123" (no separator needed)
//   FormatNumber(123, ",")는 "123" 반환 (구분자 불필요)
// - FormatBytes(512) returns "512 B"
//   FormatBytes(512)는 "512 B" 반환
// - Ellipsis with maxLen <= string length returns original
//   maxLen <= 문자열 길이인 Ellipsis는 원본 반환
//
// Masking:
// 마스킹:
// - Mask with first+last >= length returns original
//   first+last >= 길이인 Mask는 원본 반환
// - MaskEmail with invalid format returns original
//   잘못된 형식의 MaskEmail은 원본 반환
// - MaskCreditCard preserves separators (hyphens, spaces)
//   MaskCreditCard는 구분자 유지 (하이픈, 공백)
//
// Text Layout:
// 텍스트 레이아웃:
// - Indent with empty lines: only non-empty lines indented
//   빈 줄이 있는 Indent: 비어 있지 않은 줄만 들여쓰기
// - Dedent: empty lines don't affect minimum indent calculation
//   Dedent: 빈 줄은 최소 들여쓰기 계산에 영향 없음
// - WrapText with width <= 0 returns original
//   width <= 0인 WrapText는 원본 반환
// - WrapText: long words not broken, may exceed width
//   WrapText: 긴 단어는 나누지 않음, 너비 초과 가능
//
// Negative Numbers:
// 음수:
// - FormatNumber correctly handles negatives: "-1,000,000"
//   FormatNumber는 음수를 올바르게 처리: "-1,000,000"
//
// Zero Values:
// 제로값:
// - Pluralize(0, "item", "items") returns "items"
//   Pluralize(0, "item", "items")는 "items" 반환
// - FormatBytes(0) returns "0 B"
//   FormatBytes(0)는 "0 B" 반환
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Formatting File Sizes
//    파일 크기 포맷팅:
//
//    fileSize := int64(1536000)
//    readable := stringutil.FormatBytes(fileSize)
//    // "1.5 MB"
//    fmt.Printf("File size: %s\n", readable)
//    // Display in UI, logs, progress bars
//    // UI, 로그, 진행률 표시줄에 표시
//
// 2. Displaying Item Counts
//    항목 수 표시:
//
//    itemCount := 5
//    message := stringutil.FormatWithCount(itemCount, "file", "files")
//    // "5 files"
//    fmt.Printf("Found %s\n", message)
//    // User-friendly notifications
//    // 사용자 친화적 알림
//
// 3. Formatting Financial Numbers
//    재무 숫자 포맷팅:
//
//    amount := 1234567
//    formatted := stringutil.FormatNumber(amount, ",")
//    // "1,234,567"
//    fmt.Printf("Total: $%s\n", formatted)
//    // Reports, invoices, dashboards
//    // 보고서, 송장, 대시보드
//
// 4. Masking Credit Card in Receipts
//    영수증의 신용카드 마스킹:
//
//    cardNumber := "1234-5678-9012-3456"
//    masked := stringutil.MaskCreditCard(cardNumber)
//    // "****-****-****-3456"
//    fmt.Printf("Card: %s\n", masked)
//    // Privacy-safe displays
//    // 개인정보 안전 표시
//
// 5. Masking Email for Privacy
//    개인정보를 위한 이메일 마스킹:
//
//    email := "john.doe@example.com"
//    masked := stringutil.MaskEmail(email)
//    // "j******e@example.com"
//    fmt.Printf("Sent to: %s\n", masked)
//    // Confirmation messages
//    // 확인 메시지
//
// 6. Truncating Long Filenames
//    긴 파일명 자르기:
//
//    filename := "very_long_document_name_v2_final.pdf"
//    short := stringutil.Ellipsis(filename, 20)
//    // "very_lon...nal.pdf"
//    // Preserves extension for clarity
//    // 명확성을 위해 확장자 유지
//
// 7. Wrapping Text for Terminal
//    터미널을 위한 텍스트 래핑:
//
//    longText := "This is a very long line of text that needs wrapping"
//    wrapped := stringutil.WrapText(longText, 30)
//    // "This is a very long line of\ntext that needs wrapping"
//    fmt.Println(wrapped)
//    // CLI output, help messages
//    // CLI 출력, 도움말 메시지
//
// 8. Adding Line Numbers to Code
//    코드에 줄 번호 추가:
//
//    code := "func main() {\n    fmt.Println(\"Hello\")\n}"
//    numbered := stringutil.AddLineNumbers(code)
//    // "1: func main() {\n2:     fmt.Println(\"Hello\")\n3: }"
//    // Code reviews, debugging output
//    // 코드 리뷰, 디버깅 출력
//
// 9. Indenting Quoted Text
//    인용된 텍스트 들여쓰기:
//
//    quote := "This is a quote\nfrom someone"
//    indented := stringutil.Indent(quote, "> ")
//    // "> This is a quote\n> from someone"
//    // Email replies, markdown quotes
//    // 이메일 답장, 마크다운 인용
//
// 10. Cleaning Template Indentation
//     템플릿 들여쓰기 정리:
//
//     template := "    line1\n    line2\n    line3"
//     clean := stringutil.Dedent(template)
//     // "line1\nline2\nline3"
//     // Multi-line string literals
//     // 여러 줄 문자열 리터럴
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// FormatBytes vs Manual Formatting
// - FormatBytes: Automatic unit selection, 1 decimal precision
//   FormatBytes: 자동 단위 선택, 소수점 1자리 정밀도
// - Manual: More control but verbose
//   수동: 더 많은 제어 but 장황함
// - Use FormatBytes for: Standard file size display
//   FormatBytes 사용: 표준 파일 크기 표시
//
// Ellipsis (formatting.go) vs Truncate (manipulation.go)
// - Ellipsis: Middle truncation, preserves end
//   Ellipsis: 중간 자르기, 끝 유지
// - Truncate: End truncation, preserves start
//   Truncate: 끝 자르기, 시작 유지
// - Use Ellipsis for: Filenames, URLs (preserve extension)
//   Ellipsis 사용: 파일명, URL (확장자 유지)
// - Use Truncate for: General text preview
//   Truncate 사용: 일반 텍스트 미리보기
//
// Mask vs MaskEmail/MaskCreditCard
// - Mask: Generic, configurable masking
//   Mask: 범용, 구성 가능한 마스킹
// - Specific masks: Optimized for common patterns
//   특정 마스크: 일반 패턴에 최적화
// - Use Mask for: Custom masking needs
//   Mask 사용: 사용자 정의 마스킹 필요
// - Use specific for: Email, credit cards
//   특정 사용: 이메일, 신용카드
//
// FormatNumber vs fmt.Sprintf
// - FormatNumber: Adds thousand separators, customizable
//   FormatNumber: 천 단위 구분자 추가, 사용자 정의 가능
// - fmt.Sprintf: Basic number formatting
//   fmt.Sprintf: 기본 숫자 포맷팅
// - FormatNumber: Better for human-readable numbers
//   FormatNumber: 사람이 읽을 수 있는 숫자에 더 좋음
//
// WrapText vs strings.Split
// - WrapText: Intelligent word-boundary wrapping
//   WrapText: 지능적인 단어 경계 래핑
// - strings.Split: Manual line breaking
//   strings.Split: 수동 줄 나누기
// - Use WrapText for: Automatic text layout
//   WrapText 사용: 자동 텍스트 레이아웃
//
// LOCALE CONSIDERATIONS
// 로케일 고려사항
// --------------------
// - FormatNumber: Separator is customizable (comma, period, space)
//   FormatNumber: 구분자는 사용자 정의 가능 (쉼표, 마침표, 공백)
//   * US: 1,000,000 (comma)
//   * 미국: 1,000,000 (쉼표)
//   * Europe: 1.000.000 (period) or 1 000 000 (space)
//   * 유럽: 1.000.000 (마침표) 또는 1 000 000 (공백)
//
// - FormatBytes: Uses binary units (1024), not decimal (1000)
//   FormatBytes: 10진수(1000)가 아닌 2진수 단위(1024) 사용
//   * More accurate for file systems
//   * 파일 시스템에 더 정확
//
// - Pluralize: English-only, simple count == 1 rule
//   Pluralize: 영어 전용, 간단한 count == 1 규칙
//   * For other languages, use i18n libraries
//   * 다른 언어의 경우 i18n 라이브러리 사용
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
//         formatted := stringutil.FormatNumber(1000000, ",")
//     }()
//
//     go func() {
//         size := stringutil.FormatBytes(1048576)
//     }()
//
//     // All formatting functions safe for concurrent use
//     // 모든 포맷팅 함수는 동시 사용에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - manipulation.go: String manipulation (Truncate, Reverse, etc.)
//   manipulation.go: 문자열 조작 (Truncate, Reverse 등)
// - case.go: Case conversion (ToSnakeCase, ToCamelCase, etc.)
//   case.go: 케이스 변환 (ToSnakeCase, ToCamelCase 등)
// - validation.go: String validation (IsEmail, IsURL, etc.)
//   validation.go: 문자열 검증 (IsEmail, IsURL 등)
//
// =============================================================================

// FormatNumber formats an integer with thousand separators.
// FormatNumber는 천 단위 구분 기호로 정수를 포맷합니다.
//
// Example
// 예제:
//
//	FormatNumber(1000000, ",")     // "1,000,000"
//	FormatNumber(1234567, ".")     // "1.234.567"
//	FormatNumber(1234567, " ")     // "1 234 567"
//	FormatNumber(123, ",")         // "123"
func FormatNumber(n int, separator string) string {
	// Handle negative numbers
	// 음수 처리
	negative := n < 0
	if negative {
		n = -n
	}

	s := fmt.Sprintf("%d", n)
	result := ""

	// Add separators from right to left
	// 오른쪽에서 왼쪽으로 구분 기호 추가
	for i := len(s); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}

		if result != "" {
			result = separator + result
		}
		result = s[start:i] + result
	}

	if negative {
		result = "-" + result
	}

	return result
}

// FormatBytes formats bytes as human-readable size (KB, MB, GB, etc.).
// FormatBytes는 바이트를 사람이 읽을 수 있는 크기로 포맷합니다 (KB, MB, GB 등).
//
// Example
// 예제:
//
//	FormatBytes(1024)                 // "1.0 KB"
//	FormatBytes(1536)                 // "1.5 KB"
//	FormatBytes(1048576)              // "1.0 MB"
//	FormatBytes(1073741824)           // "1.0 GB"
//	FormatBytes(1099511627776)        // "1.0 TB"
func FormatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	units := []string{"KB", "MB", "GB", "TB", "PB", "EB"}
	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), units[exp])
}

// Pluralize returns the singular or plural form based on count.
// Pluralize는 count에 따라 단수형 또는 복수형을 반환합니다.
//
// Example
// 예제:
//
//	Pluralize(1, "item", "items")      // "item"
//	Pluralize(5, "item", "items")      // "items"
//	Pluralize(0, "item", "items")      // "items"
//	Pluralize(1, "person", "people")   // "person"
//	Pluralize(5, "person", "people")   // "people"
func Pluralize(count int, singular, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

// FormatWithCount returns a formatted string with count and pluralized noun.
// FormatWithCount는 count와 복수형 명사를 포함한 포맷된 문자열을 반환합니다.
//
// Example
// 예제:
//
//	FormatWithCount(1, "item", "items")    // "1 item"
//	FormatWithCount(5, "item", "items")    // "5 items"
//	FormatWithCount(0, "item", "items")    // "0 items"
func FormatWithCount(count int, singular, plural string) string {
	return fmt.Sprintf("%d %s", count, Pluralize(count, singular, plural))
}

// Ellipsis truncates a string and adds ellipsis in the middle.
// Ellipsis는 문자열을 자르고 중간에 ellipsis를 추가합니다.
//
// If maxLen is less than or equal to 3, just truncates without ellipsis.
// maxLen이 3 이하이면 ellipsis 없이 자릅니다.
//
// Example
// 예제:
//
//	Ellipsis("verylongfilename.txt", 15)  // "verylo...me.txt"
//	Ellipsis("short.txt", 20)             // "short.txt"
//	Ellipsis("abcdefgh", 3)               // "abc"
func Ellipsis(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}

	if maxLen <= 3 {
		return string(runes[:maxLen])
	}

	// Calculate split point
	// 분할 지점 계산
	ellipsisLen := 3 // "..."
	leftLen := (maxLen - ellipsisLen) / 2
	rightLen := maxLen - ellipsisLen - leftLen

	return string(runes[:leftLen]) + "..." + string(runes[len(runes)-rightLen:])
}

// Mask masks a string with a character, revealing only first and last n characters.
// Mask는 문자열을 문자로 마스크하고, 처음과 마지막 n개 문자만 표시합니다.
//
// If first+last is greater than or equal to string length, returns original string.
// first+last가 문자열 길이보다 크거나 같으면 원본 문자열을 반환합니다.
//
// Example
// 예제:
//
//	Mask("1234567890", 2, 2, "*")      // "12******90"
//	Mask("hello@example.com", 2, 4, "*")  // "he*****.com"
//	Mask("secret", 1, 1, "#")          // "s####t"
//	Mask("short", 2, 2, "*")           // "short"
func Mask(s string, first, last int, maskChar string) string {
	runes := []rune(s)
	length := len(runes)

	if first+last > length {
		return s
	}

	var result strings.Builder
	result.WriteString(string(runes[:first]))

	maskLen := length - first - last
	for i := 0; i < maskLen; i++ {
		result.WriteString(maskChar)
	}

	result.WriteString(string(runes[length-last:]))
	return result.String()
}

// MaskEmail masks an email address, revealing only the first character and domain.
// MaskEmail은 이메일 주소를 마스크하고, 첫 문자와 도메인만 표시합니다.
//
// Example
// 예제:
//
//	MaskEmail("john.doe@example.com")  // "j******e@example.com"
//	MaskEmail("a@example.com")         // "a@example.com"
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}

	local := parts[0]
	domain := parts[1]

	if len(local) <= 2 {
		return email
	}

	maskedLocal := string(local[0]) + strings.Repeat("*", len(local)-2) + string(local[len(local)-1])
	return maskedLocal + "@" + domain
}

// MaskCreditCard masks a credit card number, revealing only the last 4 digits.
// MaskCreditCard는 신용카드 번호를 마스크하고, 마지막 4자리만 표시합니다.
//
// Example
// 예제:
//
//	MaskCreditCard("1234567890123456")  // "************3456"
//	MaskCreditCard("1234-5678-9012-3456")  // "****-****-****-3456"
func MaskCreditCard(card string) string {
	runes := []rune(card)
	length := len(runes)

	if length <= 4 {
		return card
	}

	var result strings.Builder
	for i := 0; i < length-4; i++ {
		if runes[i] == '-' || runes[i] == ' ' {
			result.WriteRune(runes[i])
		} else {
			result.WriteRune('*')
		}
	}
	result.WriteString(string(runes[length-4:]))

	return result.String()
}

// AddLineNumbers adds line numbers to each line of text.
// AddLineNumbers는 텍스트의 각 줄에 줄 번호를 추가합니다.
//
// Example
// 예제:
//
//	AddLineNumbers("line1\nline2\nline3")
//	// "1: line1\n2: line2\n3: line3"
func AddLineNumbers(s string) string {
	lines := strings.Split(s, "\n")
	var result strings.Builder

	for i, line := range lines {
		if i > 0 {
			result.WriteRune('\n')
		}
		result.WriteString(fmt.Sprintf("%d: %s", i+1, line))
	}

	return result.String()
}

// Indent indents each line with the specified prefix.
// Indent는 각 줄을 지정된 접두사로 들여쓰기합니다.
//
// Example
// 예제:
//
//	Indent("line1\nline2", "  ")  // "  line1\n  line2"
//	Indent("line1\nline2", "\t")  // "\tline1\n\tline2"
func Indent(s string, prefix string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if line != "" || i < len(lines)-1 {
			lines[i] = prefix + line
		}
	}
	return strings.Join(lines, "\n")
}

// Dedent removes common leading whitespace from each line.
// Dedent는 각 줄에서 공통 선행 공백을 제거합니다.
//
// Example
// 예제:
//
//	Dedent("  line1\n  line2")  // "line1\nline2"
//	Dedent("    line1\n  line2")  // "  line1\nline2"
func Dedent(s string) string {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return s
	}

	// Find minimum indentation
	// 최소 들여쓰기 찾기
	minIndent := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		indent := 0
		for _, r := range line {
			if r != ' ' && r != '\t' {
				break
			}
			indent++
		}

		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}

	if minIndent <= 0 {
		return s
	}

	// Remove common indentation
	// 공통 들여쓰기 제거
	for i, line := range lines {
		if len(line) >= minIndent {
			lines[i] = line[minIndent:]
		}
	}

	return strings.Join(lines, "\n")
}

// WrapText wraps text to the specified line width.
// WrapText는 텍스트를 지정된 줄 너비로 줄바꿈합니다.
//
// Example
// 예제:
//
//	WrapText("The quick brown fox jumps", 10)
//	// "The quick\nbrown fox\njumps"
func WrapText(s string, width int) string {
	if width <= 0 {
		return s
	}

	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}

	var result strings.Builder
	lineLen := 0

	for i, word := range words {
		wordLen := len([]rune(word))

		if lineLen == 0 {
			// First word on line
			// 줄의 첫 단어
			result.WriteString(word)
			lineLen = wordLen
		} else if lineLen+1+wordLen <= width {
			// Word fits on current line
			// 현재 줄에 단어 맞춤
			result.WriteRune(' ')
			result.WriteString(word)
			lineLen += 1 + wordLen
		} else {
			// Start new line
			// 새 줄 시작
			result.WriteRune('\n')
			result.WriteString(word)
			lineLen = wordLen
		}

		_ = i // avoid unused variable warning
	}

	return result.String()
}
