package timeutil

import (
	"fmt"
	"strings"
	"time"
)

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/parse.go
// Purpose: Time and date parsing operations
//          시간 및 날짜 파싱 작업
//
// This file provides comprehensive time parsing functionality to convert time strings
// into time.Time objects. It offers flexible parsing methods from strict format-specific
// parsers to intelligent auto-detection parsers that can handle various formats.
//
// 이 파일은 시간 문자열을 time.Time 객체로 변환하는 포괄적인 시간 파싱 기능을 제공합니다.
// 엄격한 포맷별 파서부터 다양한 포맷을 처리할 수 있는 지능형 자동 감지 파서까지
// 유연한 파싱 방법을 제공합니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. STANDARD FORMAT PARSING (표준 포맷 파싱)
//    - ParseISO8601: Parses ISO 8601 format (2006-01-02T15:04:05Z07:00)
//      ISO 8601 포맷 파싱
//    - ParseRFC3339: Parses RFC 3339 format (identical to ISO 8601)
//      RFC 3339 포맷 파싱 (ISO 8601과 동일)
//    - Returns KST by default for international formats
//      국제 포맷의 경우 기본적으로 KST 반환
//
// 2. COMMON FORMAT PARSING (일반 포맷 파싱)
//    - ParseDate: Parses date-only strings (YYYY-MM-DD)
//      날짜만 파싱
//    - ParseDateTime: Parses datetime strings (YYYY-MM-DD HH:mm:ss)
//      날짜시간 파싱
//    - Uses KST as default timezone
//      기본 타임존으로 KST 사용
//
// 3. AUTO-DETECTION PARSING (자동 감지 파싱)
//    - Parse: Intelligent parser that auto-detects common formats
//      일반적인 포맷을 자동 감지하는 지능형 파서
//    - Tries ISO8601, RFC3339, DateTime, and Date formats
//      ISO8601, RFC3339, DateTime, Date 포맷 시도
//    - Sequential detection based on string patterns (T, space, colon, dash)
//      문자열 패턴 기반 순차 감지 (T, 공백, 콜론, 대시)
//
// 4. TIMEZONE-SPECIFIC PARSING (타임존별 파싱)
//    - ParseWithTimezone: Parses in specific timezone
//      특정 타임존에서 파싱
//    - Tries multiple formats until successful
//      성공할 때까지 여러 포맷 시도
//    - Useful for parsing times from different regions
//      다른 지역의 시간을 파싱하는 데 유용
//
// 5. CUSTOM LAYOUT PARSING (커스텀 레이아웃 파싱)
//    - ParseWithLayout: Parses with user-specified Go layout
//      사용자 지정 Go 레이아웃으로 파싱
//    - Full control over format
//      포맷에 대한 완전한 제어
//    - Supports any Go time format string
//      모든 Go 시간 포맷 문자열 지원
//
// 6. HIGH-PRECISION PARSING (고정밀 파싱)
//    - ParseMillis: Parses with millisecond precision (.SSS)
//      밀리초 정밀도로 파싱
//    - ParseMicros: Parses with microsecond precision (.SSSSSS)
//      마이크로초 정밀도로 파싱
//    - Useful for database timestamps and logs
//      데이터베이스 타임스탬프와 로그에 유용
//
// 7. UNIVERSAL PARSING (범용 파싱)
//    - ParseAny: Tries all known formats (40+ layouts)
//      알려진 모든 포맷 시도 (40개 이상의 레이아웃)
//    - Handles international formats (ISO8601, RFC3339, RFC822, RFC1123, etc.)
//      국제 포맷 처리
//    - Handles regional formats (US: MM/DD/YYYY, EU: DD-MM-YYYY, DE: DD.MM.YYYY)
//      지역별 포맷 처리
//    - Handles Korean formats (2006년 01월 02일, with/without 오전/오후)
//      한글 포맷 처리
//    - Handles month names (Jan, January, 월 이름)
//      월 이름 처리
//    - Last resort when format is unknown
//      포맷을 알 수 없을 때 최후의 수단
//
// ============================================================================
// PARSING STRATEGIES / 파싱 전략
// ============================================================================
//
// This file implements a TIERED PARSING APPROACH:
// 이 파일은 계층적 파싱 접근 방식을 구현합니다:
//
// TIER 1: STRICT FORMAT-SPECIFIC PARSERS (엄격한 포맷별 파서)
// - ParseISO8601, ParseRFC3339, ParseDate, ParseDateTime
// - Fastest performance (단일 포맷 시도)
// - Use when format is known (포맷을 알 때 사용)
// - Provides clear error messages (명확한 오류 메시지)
// - Best practice: Use these for API inputs with known formats
//
// TIER 2: SMART AUTO-DETECTION (스마트 자동 감지)
// - Parse: Auto-detects among 4 common formats
// - Good performance (최대 4번의 파싱 시도)
// - Use for user input or mixed sources (사용자 입력 또는 혼합 소스용)
// - Pattern-based detection (contains T, space, colon, dash)
// - Best practice: Use for flexible user inputs
//
// TIER 3: TIMEZONE-AWARE PARSING (타임존 인식 파싱)
// - ParseWithTimezone: Parses in specific timezone
// - Moderate performance (여러 포맷 시도)
// - Use for multi-timezone applications
// - Best practice: Use when timezone is known from context
//
// TIER 4: CUSTOM LAYOUT PARSING (커스텀 레이아웃 파싱)
// - ParseWithLayout: Maximum flexibility
// - Fast performance (단일 레이아웃 시도)
// - Use for non-standard formats
// - Best practice: Use for legacy system integration
//
// TIER 5: HIGH-PRECISION PARSING (고정밀 파싱)
// - ParseMillis, ParseMicros: Sub-second precision
// - Use for database timestamps and logs
// - Best practice: Use when precision matters
//
// TIER 6: UNIVERSAL FALLBACK (범용 대체)
// - ParseAny: Tries all known formats (40+ layouts)
// - Slowest performance (최대 40개 이상의 포맷 시도)
// - Use when format is completely unknown
// - Best practice: Use as last resort or for exploratory parsing
// - Performance warning: Avoid in hot paths (성능 주의: 빈번한 경로에서 피할 것)
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. INVERSE OF FORMATTING (포맷팅의 역연산)
//    - Parsing is the inverse operation of formatting
//      파싱은 포맷팅의 역연산입니다
//    - Each Parse* function corresponds to a Format* function in format.go
//      각 Parse* 함수는 format.go의 Format* 함수에 대응됩니다
//    - Example: ParseISO8601 ↔ FormatISO8601
//      예시: ParseISO8601 ↔ FormatISO8601
//    - Example: ParseDate ↔ FormatDate
//      예시: ParseDate ↔ FormatDate
//
// 2. TIMEZONE HANDLING (타임존 처리)
//    - International formats (ISO8601, RFC3339): Parse with timezone, then convert to KST
//      국제 포맷: 타임존으로 파싱 후 KST로 변환
//    - Local formats (Date, DateTime): Parse directly in KST
//      로컬 포맷: KST에서 직접 파싱
//    - Reason: Consistent with format.go which converts to KST before formatting
//      이유: KST로 변환 후 포맷팅하는 format.go와 일관성 유지
//
// 3. ERROR HANDLING (오류 처리)
//    - All parsing errors are wrapped with context
//      모든 파싱 오류는 컨텍스트와 함께 래핑됩니다
//    - Error messages include the attempted format
//      오류 메시지에 시도한 포맷 포함
//    - Helps debugging: "failed to parse ISO8601: ..." vs generic error
//      디버깅에 도움: "failed to parse ISO8601: ..." vs 일반 오류
//
// 4. FLEXIBILITY SPECTRUM (유연성 스펙트럼)
//    - Strict parsers (ParseDate) → Moderate (Parse) → Flexible (ParseAny)
//      엄격한 파서 → 보통 → 유연한 파서
//    - Trade-off: Performance vs flexibility
//      트레이드오프: 성능 vs 유연성
//    - Choose based on use case:
//      사용 사례에 따라 선택:
//      * Known format: Use strict parsers (성능 우선)
//      * User input: Use Parse or ParseAny (유연성 우선)
//      * API with contract: Use strict parsers (계약 준수)
//
// 5. KOREAN LOCALIZATION (한글 로컬라이제이션)
//    - ParseAny includes Korean date formats
//      ParseAny는 한글 날짜 포맷 포함
//    - Supports both padded and unpadded numbers
//      0으로 채워진 숫자와 채워지지 않은 숫자 모두 지원
//    - Supports 오전/오후 (AM/PM) notation
//      오전/오후 표기 지원
//    - Example: "2024년 1월 5일", "2024년 01월 05일 오후 3시"
//      예시: "2024년 1월 5일", "2024년 01월 05일 오후 3시"
//
// ============================================================================
// PARSING FUNCTIONS OVERVIEW / 파싱 함수 개요
// ============================================================================
//
// STANDARD FORMAT PARSERS (표준 포맷 파서) - 2 functions
// ├─ ParseISO8601      : ISO 8601 format (2006-01-02T15:04:05Z07:00)
// └─ ParseRFC3339      : RFC 3339 format (identical to ISO 8601)
//
// COMMON FORMAT PARSERS (일반 포맷 파서) - 2 functions
// ├─ ParseDate         : Date only (YYYY-MM-DD)
// └─ ParseDateTime     : Date and time (YYYY-MM-DD HH:mm:ss)
//
// INTELLIGENT PARSERS (지능형 파서) - 1 function
// └─ Parse             : Auto-detects among 4 common formats
//
// TIMEZONE-SPECIFIC PARSERS (타임존별 파서) - 1 function
// └─ ParseWithTimezone : Parses in specified timezone
//
// CUSTOM LAYOUT PARSERS (커스텀 레이아웃 파서) - 1 function
// └─ ParseWithLayout   : Parses with custom Go layout
//
// HIGH-PRECISION PARSERS (고정밀 파서) - 2 functions
// ├─ ParseMillis       : Millisecond precision (.SSS)
// └─ ParseMicros       : Microsecond precision (.SSSSSS)
//
// UNIVERSAL PARSER (범용 파서) - 1 function
// └─ ParseAny          : Tries 40+ formats (last resort)
//
// Total: 10 parsing functions
// 총: 10개의 파싱 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
// - ParseISO8601, ParseRFC3339, ParseDate, ParseDateTime: O(n)
//   where n = length of input string
//   입력 문자열의 길이
//
// - Parse: O(m * n)
//   where m = number of formats tried (최대 4)
//         n = length of input string
//   시도하는 포맷 수 (최대 4)
//
// - ParseWithTimezone: O(f * n)
//   where f = number of formats tried (4)
//         n = length of input string
//   시도하는 포맷 수 (4)
//
// - ParseWithLayout: O(n)
//   where n = length of input string
//   입력 문자열의 길이
//
// - ParseMillis, ParseMicros: O(n)
//   where n = length of input string
//   입력 문자열의 길이
//
// - ParseAny: O(f * n)
//   where f = number of formats tried (40+)
//         n = length of input string
//   시도하는 포맷 수 (40개 이상)
//   WARNING: Can be slow for unknown formats
//   경고: 알 수 없는 포맷의 경우 느릴 수 있음
//
// SPACE COMPLEXITY (공간 복잡도):
// - All functions: O(1) - constant space (상수 공간)
// - No additional memory allocation except for result
//   결과를 제외한 추가 메모리 할당 없음
//
// PERFORMANCE TIPS (성능 팁):
// 1. Use strict parsers when format is known
//    포맷을 알 때는 엄격한 파서 사용
// 2. Avoid ParseAny in hot paths
//    빈번한 경로에서 ParseAny 피하기
// 3. Cache timezone locations if using ParseWithTimezone frequently
//    ParseWithTimezone을 자주 사용할 경우 타임존 위치 캐시
// 4. Pre-validate input format before parsing if possible
//    가능하면 파싱 전에 입력 포맷 사전 검증
//
// ============================================================================
// FORMAT DETECTION PATTERNS / 포맷 감지 패턴
// ============================================================================
//
// Parse() uses PATTERN-BASED DETECTION:
// Parse()는 패턴 기반 감지를 사용합니다:
//
// Pattern: Contains "T"
// 패턴: "T" 포함
// ├─ Likely: ISO8601 or RFC3339
// └─ Example: "2024-01-15T14:30:00Z"
//
// Pattern: Contains space AND colon
// 패턴: 공백과 콜론 포함
// ├─ Likely: DateTime format
// └─ Example: "2024-01-15 14:30:00"
//
// Pattern: Two dashes, no space/colon
// 패턴: 대시 2개, 공백/콜론 없음
// ├─ Likely: Date format
// └─ Example: "2024-01-15"
//
// Pattern: None of above
// 패턴: 위의 어느 것도 아님
// ├─ Result: Parse fails
// └─ Solution: Use ParseAny() for unknown formats
//
// ============================================================================
// GO TIME FORMAT REFERENCE / Go 시간 포맷 참조
// ============================================================================
//
// Go uses a REFERENCE TIME for layouts:
// Go는 레이아웃에 참조 시간을 사용합니다:
//
// Reference time: Mon Jan 2 15:04:05 MST 2006
// 참조 시간: Mon Jan 2 15:04:05 MST 2006
//
// This represents: 01/02 03:04:05 PM '06 -0700
// 이것은 다음을 나타냅니다: 01/02 03:04:05 PM '06 -0700
//
// Mnemonic: 1, 2, 3, 4, 5, 6, 7
// 기억법: 1, 2, 3, 4, 5, 6, 7
//
// Components:
// 구성 요소:
// - Month      : 01, Jan, January (월)
// - Day        : 02, 2, _2 (일)
// - Hour       : 15 (24-hour), 03 (12-hour) (시)
// - Minute     : 04 (분)
// - Second     : 05 (초)
// - Year       : 2006, 06 (년)
// - Timezone   : MST, -0700, Z07:00 (타임존)
// - Sub-second : .000 (ms), .999999 (μs), .999999999 (ns)
//
// Examples of layouts:
// 레이아웃 예시:
// - "2006-01-02"              → YYYY-MM-DD
// - "2006-01-02 15:04:05"     → YYYY-MM-DD HH:mm:ss
// - "2006-01-02T15:04:05Z07:00" → ISO8601
// - "Jan 02, 2006"            → Mon DD, YYYY
// - "2006년 01월 02일"        → Korean date
//
// ============================================================================
// TIMEZONE BEHAVIOR / 타임존 동작
// ============================================================================
//
// INTERNATIONAL FORMATS (ISO8601, RFC3339):
// 국제 포맷 (ISO8601, RFC3339):
// 1. Parse with embedded timezone
//    포함된 타임존으로 파싱
// 2. Convert to KST (Asia/Seoul)
//    KST로 변환
// 3. Return time in KST
//    KST 시간 반환
// Example:
//   Input: "2024-01-15T14:30:00+09:00" (KST)
//   Parse: 2024-01-15 14:30:00 +0900 KST
//   Return: 2024-01-15 14:30:00 +0900 KST
//
// Example with conversion:
//   Input: "2024-01-15T05:30:00Z" (UTC)
//   Parse: 2024-01-15 05:30:00 +0000 UTC
//   Convert: 2024-01-15 14:30:00 +0900 KST
//   Return: 2024-01-15 14:30:00 +0900 KST
//
// LOCAL FORMATS (Date, DateTime):
// 로컬 포맷 (Date, DateTime):
// 1. Parse directly in KST
//    KST에서 직접 파싱
// 2. No conversion needed
//    변환 불필요
// 3. Return time in KST
//    KST 시간 반환
// Example:
//   Input: "2024-01-15 14:30:00"
//   Parse in KST: 2024-01-15 14:30:00 +0900 KST
//   Return: 2024-01-15 14:30:00 +0900 KST
//
// CUSTOM TIMEZONE (ParseWithTimezone):
// 커스텀 타임존 (ParseWithTimezone):
// 1. Load specified timezone
//    지정된 타임존 로드
// 2. Parse in that timezone
//    해당 타임존에서 파싱
// 3. Return time in specified timezone (no conversion to KST)
//    지정된 타임존의 시간 반환 (KST로 변환 안 함)
// Example:
//   Input: "2024-01-15 14:30:00", timezone: "America/New_York"
//   Parse in EST: 2024-01-15 14:30:00 -0500 EST
//   Return: 2024-01-15 14:30:00 -0500 EST
//
// ============================================================================
// ERROR HANDLING / 오류 처리
// ============================================================================
//
// All parsing functions return (time.Time, error):
// 모든 파싱 함수는 (time.Time, error)를 반환합니다:
//
// SUCCESS CASE:
// 성공 케이스:
// - Returns valid time.Time with proper timezone
//   올바른 타임존이 있는 유효한 time.Time 반환
// - Error is nil
//   오류는 nil
//
// FAILURE CASE:
// 실패 케이스:
// - Returns zero time.Time (time.Time{})
//   제로 time.Time 반환
// - Error contains context: "failed to parse [format]: [details]"
//   오류에 컨텍스트 포함: "failed to parse [format]: [details]"
//
// ERROR CHECKING:
// 오류 확인:
//   t, err := timeutil.ParseDate("2024-01-15")
//   if err != nil {
//       log.Printf("Failed to parse: %v", err)
//       return err
//   }
//   // Use t safely
//
// COMMON ERRORS:
// 일반 오류:
// - Invalid format: String doesn't match expected layout
//   잘못된 포맷: 문자열이 예상 레이아웃과 일치하지 않음
// - Out of range: Values like month 13, day 32
//   범위 초과: 월 13, 일 32 같은 값
// - Invalid timezone: Unknown timezone name
//   잘못된 타임존: 알 수 없는 타임존 이름
// - Empty string: Input is empty or whitespace only
//   빈 문자열: 입력이 비어 있거나 공백만 있음
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: API Input Validation (API 입력 검증)
// Use case: Parsing ISO8601 timestamps from REST API
// 사용 사례: REST API에서 ISO8601 타임스탬프 파싱
//
//   // Strict format expected
//   t, err := timeutil.ParseISO8601(request.Timestamp)
//   if err != nil {
//       return fmt.Errorf("invalid timestamp format: %w", err)
//   }
//
//   // Now t is in KST, ready for database
//   db.Save(Event{Time: t})
//
// PATTERN 2: Database Timestamp Parsing (데이터베이스 타임스탬프 파싱)
// Use case: Parsing MySQL datetime with milliseconds
// 사용 사례: MySQL datetime 밀리초 포함 파싱
//
//   // High-precision timestamp
//   t, err := timeutil.ParseMillis(row.CreatedAt)
//   if err != nil {
//       log.Printf("Failed to parse timestamp: %v", err)
//       return err
//   }
//
//   // Use for calculations
//   elapsed := time.Since(t)
//
// PATTERN 3: User Input Parsing (사용자 입력 파싱)
// Use case: Flexible date input from web form
// 사용 사례: 웹 폼에서 유연한 날짜 입력
//
//   // Auto-detect common formats
//   t, err := timeutil.Parse(userInput)
//   if err != nil {
//       // Fallback to universal parser
//       t, err = timeutil.ParseAny(userInput)
//       if err != nil {
//           return fmt.Errorf("unable to parse date: %w", err)
//       }
//   }
//
//   // Process the date
//   processDate(t)
//
// PATTERN 4: Multi-Timezone Application (다중 타임존 애플리케이션)
// Use case: Parsing times from different timezone sources
// 사용 사례: 다른 타임존 소스에서 시간 파싱
//
//   // Parse in specific timezone
//   nyTime, err := timeutil.ParseWithTimezone(
//       "2024-01-15 14:30:00",
//       "America/New_York",
//   )
//   if err != nil {
//       return err
//   }
//
//   // Convert to KST for comparison
//   kstTime := nyTime.In(timeutil.KST)
//   fmt.Printf("NY: %s, KST: %s\n",
//       nyTime.Format("15:04"),
//       kstTime.Format("15:04"))
//
// PATTERN 5: Legacy System Integration (레거시 시스템 통합)
// Use case: Parsing non-standard date formats
// 사용 사례: 비표준 날짜 포맷 파싱
//
//   // Custom layout for legacy format
//   layout := "02/01/2006 15:04:05"  // DD/MM/YYYY HH:mm:ss
//   t, err := timeutil.ParseWithLayout(legacyDate, layout)
//   if err != nil {
//       return fmt.Errorf("failed to parse legacy date: %w", err)
//   }
//
//   // Convert to standard format
//   standardDate := timeutil.FormatISO8601(t)
//
// PATTERN 6: Log Parsing (로그 파싱)
// Use case: Parsing various log timestamp formats
// 사용 사례: 다양한 로그 타임스탬프 포맷 파싱
//
//   // Try all known formats
//   t, err := timeutil.ParseAny(logLine.Timestamp)
//   if err != nil {
//       log.Printf("Unknown timestamp format: %s", logLine.Timestamp)
//       return err
//   }
//
//   // Filter logs by time range
//   if t.After(startTime) && t.Before(endTime) {
//       processLog(logLine)
//   }
//
// PATTERN 7: Date Range Parsing (날짜 범위 파싱)
// Use case: Parsing start and end dates for queries
// 사용 사례: 쿼리를 위한 시작 및 종료 날짜 파싱
//
//   // Parse date-only strings
//   startDate, err := timeutil.ParseDate(request.StartDate)
//   if err != nil {
//       return fmt.Errorf("invalid start date: %w", err)
//   }
//
//   endDate, err := timeutil.ParseDate(request.EndDate)
//   if err != nil {
//       return fmt.Errorf("invalid end date: %w", err)
//   }
//
//   // Query database
//   results := db.Query(
//       "SELECT * FROM events WHERE created_at BETWEEN ? AND ?",
//       startDate, endDate,
//   )
//
// PATTERN 8: Korean Date Input (한글 날짜 입력)
// Use case: Parsing Korean-formatted dates from Korean users
// 사용 사례: 한국 사용자로부터 한글 포맷 날짜 파싱
//
//   // Handle Korean date formats
//   koreanDate := "2024년 1월 15일 오후 3시 30분"
//   t, err := timeutil.ParseAny(koreanDate)
//   if err != nil {
//       return fmt.Errorf("한글 날짜 형식을 파싱할 수 없습니다: %w", err)
//   }
//
//   // Convert to standard format for storage
//   isoDate := timeutil.FormatISO8601(t)
//   db.Save(isoDate)
//
// PATTERN 9: Batch Parsing with Error Handling (오류 처리가 있는 일괄 파싱)
// Use case: Parsing multiple timestamps with graceful error handling
// 사용 사례: 우아한 오류 처리로 여러 타임스탬프 파싱
//
//   var parsedTimes []time.Time
//   var errors []error
//
//   for _, ts := range timestamps {
//       t, err := timeutil.Parse(ts)
//       if err != nil {
//           // Try ParseAny as fallback
//           t, err = timeutil.ParseAny(ts)
//           if err != nil {
//               errors = append(errors, fmt.Errorf("failed to parse %s: %w", ts, err))
//               continue
//           }
//       }
//       parsedTimes = append(parsedTimes, t)
//   }
//
//   if len(errors) > 0 {
//       log.Printf("Parsing errors: %v", errors)
//   }
//
// PATTERN 10: Performance-Critical Parsing (성능 중요 파싱)
// Use case: High-throughput parsing with known format
// 사용 사례: 알려진 포맷으로 높은 처리량 파싱
//
//   // Use strict parser for best performance
//   var times []time.Time
//   for _, dateStr := range largeDateList {
//       t, err := timeutil.ParseDateTime(dateStr)
//       if err != nil {
//           log.Printf("Invalid date: %s, error: %v", dateStr, err)
//           continue
//       }
//       times = append(times, t)
//   }
//
//   // Avoid ParseAny() in loops - too slow
//   // 루프에서 ParseAny() 피하기 - 너무 느림
//
// ============================================================================
// COMPARISON WITH FORMAT.GO / format.go와의 비교
// ============================================================================
//
// INVERSE OPERATIONS (역연산):
// Parse and Format functions are inverses:
// Parse와 Format 함수는 역연산입니다:
//
// ParseISO8601 ↔ FormatISO8601
//   Parse:  "2024-01-15T14:30:00+09:00" → time.Time
//   Format: time.Time → "2024-01-15T14:30:00+09:00"
//
// ParseRFC3339 ↔ FormatRFC3339
//   Parse:  "2024-01-15T14:30:00+09:00" → time.Time
//   Format: time.Time → "2024-01-15T14:30:00+09:00"
//
// ParseDate ↔ FormatDate
//   Parse:  "2024-01-15" → time.Time
//   Format: time.Time → "2024-01-15"
//
// ParseDateTime ↔ FormatDateTime
//   Parse:  "2024-01-15 14:30:00" → time.Time
//   Format: time.Time → "2024-01-15 14:30:00"
//
// ParseWithLayout ↔ Format (with custom tokens)
//   Parse:  Custom layout → time.Time
//   Format: time.Time → Custom token format
//
// COMPLEMENTARY FUNCTIONS (보완 함수):
// Some parse functions don't have direct format equivalents:
// 일부 파싱 함수는 직접적인 포맷 대응이 없습니다:
//
// Parse:
//   No direct equivalent - auto-detects format
//   직접 대응 없음 - 포맷 자동 감지
//
// ParseAny:
//   No direct equivalent - tries all formats
//   직접 대응 없음 - 모든 포맷 시도
//
// ParseWithTimezone:
//   Complement: FormatWithTimezone
//   보완: FormatWithTimezone
//
// ParseMillis/ParseMicros:
//   No direct equivalent - use Format with custom layout
//   직접 대응 없음 - 커스텀 레이아웃으로 Format 사용
//
// USAGE RECOMMENDATION (사용 권장사항):
// - For known formats: Use matching Parse/Format pair
//   알려진 포맷: 일치하는 Parse/Format 쌍 사용
// - For flexible input: Use Parse or ParseAny
//   유연한 입력: Parse 또는 ParseAny 사용
// - For API contracts: Use strict format pairs
//   API 계약: 엄격한 포맷 쌍 사용
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// All parsing functions in this file are THREAD-SAFE:
// 이 파일의 모든 파싱 함수는 스레드 안전합니다:
//
// 1. READ-ONLY CONSTANTS (읽기 전용 상수)
//    - All layout constants are immutable
//      모든 레이아웃 상수는 불변입니다
//    - defaultLocation is pre-loaded and read-only
//      defaultLocation은 미리 로드되고 읽기 전용입니다
//    - No mutation of shared state
//      공유 상태의 변경 없음
//
// 2. NO SHARED STATE (공유 상태 없음)
//    - Each function operates on input parameters only
//      각 함수는 입력 매개변수만 사용합니다
//    - No global variables are modified
//      전역 변수가 수정되지 않습니다
//    - Pure functions with no side effects
//      부작용이 없는 순수 함수
//
// 3. SAFE CONCURRENT USAGE (안전한 동시 사용)
//    - Multiple goroutines can call parsing functions simultaneously
//      여러 고루틴이 동시에 파싱 함수를 호출할 수 있습니다
//    - No locks or synchronization needed
//      잠금이나 동기화가 필요 없습니다
//    - Example:
//      예시:
//        for _, dateStr := range dates {
//            go func(s string) {
//                t, err := timeutil.ParseDate(s)
//                // Process t
//            }(dateStr)
//        }
//
// 4. GO'S TIME PACKAGE SAFETY (Go의 time 패키지 안전성)
//    - Go's time.Parse* functions are thread-safe
//      Go의 time.Parse* 함수는 스레드 안전합니다
//    - time.Time is immutable and safe for concurrent use
//      time.Time은 불변이며 동시 사용에 안전합니다
//
// IMPORTANT NOTES (중요 사항):
// - loadTimezone() caches locations internally (Go's time.LoadLocation)
//   loadTimezone()은 내부적으로 위치를 캐시합니다 (Go의 time.LoadLocation)
// - First call to LoadLocation for a timezone may be slower
//   타임존에 대한 첫 LoadLocation 호출은 느릴 수 있습니다
// - Subsequent calls use cached location (thread-safe cache in Go runtime)
//   후속 호출은 캐시된 위치를 사용합니다 (Go 런타임의 스레드 안전 캐시)
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM constants.go:
// - ISO8601Layout: "2006-01-02T15:04:05Z07:00"
// - RFC3339Layout: time.RFC3339
// - DateLayout: "2006-01-02"
// - DateTimeLayout: "2006-01-02 15:04:05"
// - defaultLocation: KST (Asia/Seoul)
//
// FROM timezone.go:
// - loadTimezone(): Loads timezone location
//
// USED BY (사용처):
// - Application code for converting strings to time.Time
//   문자열을 time.Time으로 변환하는 애플리케이션 코드
// - API handlers for request parsing
//   요청 파싱을 위한 API 핸들러
// - Database layer for timestamp parsing
//   타임스탬프 파싱을 위한 데이터베이스 계층
// - Log parsers and analytics
//   로그 파서 및 분석
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. PREFER STRICT PARSERS FOR KNOWN FORMATS
//    알려진 포맷에는 엄격한 파서 선호
//    ✓ Good: timeutil.ParseDate("2024-01-15")
//    ✗ Avoid: timeutil.ParseAny("2024-01-15")  // Slower
//
// 2. VALIDATE FORMAT BEFORE PARSING IF POSSIBLE
//    가능하면 파싱 전에 포맷 검증
//    if strings.Contains(input, "T") {
//        t, err = timeutil.ParseISO8601(input)
//    } else {
//        t, err = timeutil.ParseDate(input)
//    }
//
// 3. HANDLE ERRORS APPROPRIATELY
//    오류를 적절하게 처리
//    t, err := timeutil.ParseDate(input)
//    if err != nil {
//        return fmt.Errorf("invalid date: %w", err)
//    }
//
// 4. USE PARSEANY AS LAST RESORT
//    ParseAny를 최후의 수단으로 사용
//    // Try specific parser first
//    t, err := timeutil.Parse(input)
//    if err != nil {
//        // Fallback to ParseAny
//        t, err = timeutil.ParseAny(input)
//    }
//
// 5. CACHE TIMEZONE LOCATIONS
//    타임존 위치 캐시
//    // Go's time.LoadLocation caches automatically
//    // Go의 time.LoadLocation은 자동으로 캐시합니다
//
// 6. USE PARSEWITHOUTPUT FOR LOGGING
//    로깅에는 ParseWithLayout 사용
//    // Clear and explicit
//    layout := "2006-01-02 15:04:05.000"
//    t, err := timeutil.ParseWithLayout(logTimestamp, layout)
//
// 7. DOCUMENT EXPECTED FORMATS IN API
//    API에서 예상 포맷 문서화
//    // In API documentation:
//    // "timestamp field must be in ISO8601 format"
//    // Use ParseISO8601 for strict validation
//
// 8. TEST WITH VARIOUS INPUTS
//    다양한 입력으로 테스트
//    // Test valid, invalid, edge cases
//    // Empty string, malformed, timezone variations
//
// ============================================================================

// ParseISO8601 parses a time string in ISO8601 format.
// ParseISO8601은 ISO8601 포맷의 시간 문자열을 파싱합니다.
func ParseISO8601(s string) (time.Time, error) {
	t, err := time.Parse(ISO8601Layout, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse ISO8601: %w", err)
	}
	return t.In(defaultLocation), nil
}

// ParseRFC3339 parses a time string in RFC3339 format.
// ParseRFC3339는 RFC3339 포맷의 시간 문자열을 파싱합니다.
func ParseRFC3339(s string) (time.Time, error) {
	t, err := time.Parse(RFC3339Layout, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse RFC3339: %w", err)
	}
	return t.In(defaultLocation), nil
}

// ParseDate parses a date string (YYYY-MM-DD).
// ParseDate는 날짜 문자열을 파싱합니다 (YYYY-MM-DD).
func ParseDate(s string) (time.Time, error) {
	t, err := time.ParseInLocation(DateLayout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %w", err)
	}
	return t, nil
}

// ParseDateTime parses a datetime string (YYYY-MM-DD HH:mm:ss).
// ParseDateTime은 날짜시간 문자열을 파싱합니다 (YYYY-MM-DD HH:mm:ss).
func ParseDateTime(s string) (time.Time, error) {
	t, err := time.ParseInLocation(DateTimeLayout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime: %w", err)
	}
	return t, nil
}

// Parse attempts to parse a time string by auto-detecting the format.
// Parse는 포맷을 자동 감지하여 시간 문자열을 파싱합니다.
//
// Supported formats
// 지원되는 포맷:
//   - ISO8601: 2006-01-02T15:04:05Z07:00
//   - RFC3339: 2006-01-02T15:04:05Z07:00
//   - Date: 2006-01-02
//   - DateTime: 2006-01-02 15:04:05
func Parse(s string) (time.Time, error) {
	s = strings.TrimSpace(s)

	// Try ISO8601/RFC3339
	// ISO8601/RFC3339 시도
	if strings.Contains(s, "T") {
		t, err := time.Parse(time.RFC3339, s)
		if err == nil {
			return t.In(defaultLocation), nil
		}
		t, err = time.Parse(ISO8601Layout, s)
		if err == nil {
			return t.In(defaultLocation), nil
		}
	}

	// Try DateTime
	// DateTime 시도
	if strings.Contains(s, " ") && strings.Contains(s, ":") {
		t, err := time.ParseInLocation(DateTimeLayout, s, defaultLocation)
		if err == nil {
			return t, nil
		}
	}

	// Try Date
	// Date 시도
	if strings.Count(s, "-") == 2 {
		t, err := time.ParseInLocation(DateLayout, s, defaultLocation)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time string: %s", s)
}

// ParseWithTimezone parses a time string in a specific timezone.
// ParseWithTimezone은 특정 타임존에서 시간 문자열을 파싱합니다.
func ParseWithTimezone(s, tz string) (time.Time, error) {
	loc, err := loadTimezone(tz)
	if err != nil {
		return time.Time{}, err
	}

	// Try different formats
	// 다른 포맷 시도
	formats := []string{
		ISO8601Layout,
		RFC3339Layout,
		DateTimeLayout,
		DateLayout,
	}

	for _, layout := range formats {
		t, err := time.ParseInLocation(layout, s, loc)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time string in timezone %s: %s", tz, s)
}

// ParseWithLayout parses a time string with a custom layout.
// ParseWithLayout은 사용자 지정 레이아웃으로 시간 문자열을 파싱합니다.
//
// Example layouts
// 레이아웃 예제:
//   - "2006-01-02 15:04:05.000" for milliseconds
//
// 밀리초용
//   - "2006-01-02 15:04:05.999999" for microseconds
//
// 마이크로초용
//   - "2006/01/02" for date with slashes
//
// 슬래시 구분 날짜
func ParseWithLayout(s, layout string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse with layout %s: %w", layout, err)
	}
	return t, nil
}

// ParseMillis parses a datetime string with milliseconds (YYYY-MM-DD HH:mm:ss.SSS).
// ParseMillis는 밀리초를 포함한 날짜시간 문자열을 파싱합니다 (YYYY-MM-DD HH:mm:ss.SSS).
func ParseMillis(s string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.000"
	t, err := time.ParseInLocation(layout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse milliseconds: %w", err)
	}
	return t, nil
}

// ParseMicros parses a datetime string with microseconds (YYYY-MM-DD HH:mm:ss.SSSSSS).
// ParseMicros는 마이크로초를 포함한 날짜시간 문자열을 파싱합니다 (YYYY-MM-DD HH:mm:ss.SSSSSS).
func ParseMicros(s string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.999999"
	t, err := time.ParseInLocation(layout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse microseconds: %w", err)
	}
	return t, nil
}

// ParseAny attempts to parse a time string by trying all common formats.
// ParseAny는 모든 일반적인 포맷을 시도하여 시간 문자열을 파싱합니다.
//
// Supported formats
// 지원되는 포맷:
//   - ISO8601: 2006-01-02T15:04:05Z07:00
//   - RFC3339: 2006-01-02T15:04:05Z07:00
//   - DateTime with milliseconds: 2006-01-02 15:04:05.000
//   - DateTime with microseconds: 2006-01-02 15:04:05.999999
//   - DateTime with nanoseconds: 2006-01-02 15:04:05.999999999
//   - DateTime: 2006-01-02 15:04:05
//   - Date: 2006-01-02
//   - Date with slashes: 2006/01/02
//   - DateTime with slashes: 2006/01/02 15:04:05
//   - US format: 01/02/2006
//   - US format with time: 01/02/2006 15:04:05
//   - Month name: Jan 02, 2006
//   - Short month: 02-Jan-2006
//   - RFC822: 02 Jan 06 15:04 MST
//   - RFC1123: Mon, 02 Jan 2006 15:04:05 MST
//   - ANSIC: Mon Jan _2 15:04:05 2006
//   - UnixDate: Mon Jan _2 15:04:05 MST 2006
//   - RubyDate: Mon Jan 02 15:04:05 -0700 2006
func ParseAny(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, fmt.Errorf("empty time string")
	}

	// All formats to try
	// 시도할 모든 포맷
	formats := []string{
		// ISO8601 and RFC3339
		// ISO8601 및 RFC3339
		time.RFC3339,
		time.RFC3339Nano,
		ISO8601Layout,
		"2006-01-02T15:04:05Z0700",
		"2006-01-02T15:04:05",

		// DateTime with sub-seconds
		// 밀리초/마이크로초/나노초 포함
		"2006-01-02 15:04:05.999999999",
		"2006-01-02 15:04:05.999999",
		"2006-01-02 15:04:05.000",

		// DateTime
		// 날짜시간
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"2006-01-02 15:04",
		"2006/01/02 15:04",

		// Date only
		// 날짜만
		"2006-01-02",
		"2006/01/02",
		"01/02/2006", // US format
		"02-01-2006", // EU format
		"02.01.2006", // DE format

		// Korean formats
		// 한글 포맷
		"2006년 01월 02일 15시 04분 05초",
		"2006년 01월 02일 15시 04분",
		"2006년 01월 02일 15시",
		"2006년 01월 02일",
		"2006년 1월 2일 15시 4분 5초",
		"2006년 1월 2일 15시 4분",
		"2006년 1월 2일 15시",
		"2006년 1월 2일",
		"2006년 01월 02일 오후 3시 04분 05초",
		"2006년 01월 02일 오후 3시 04분",
		"2006년 01월 02일 오후 3시",
		"2006년 01월 02일 오전 9시 04분 05초",
		"2006년 01월 02일 오전 9시 04분",
		"2006년 01월 02일 오전 9시",

		// With month names
		// 월 이름 포함
		"Jan 02, 2006",
		"January 02, 2006",
		"02-Jan-2006",
		"02-January-2006",
		"02 Jan 2006",
		"02 January 2006",

		// Standard Go time formats
		// Go 표준 시간 포맷
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,

		// Additional common formats
		// 추가 일반 포맷
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05 -0700",
		"Mon, 02 Jan 2006 15:04:05 -0700",
	}

	// Try parsing with each format
	// 각 포맷으로 파싱 시도
	for _, layout := range formats {
		// Try with default location first
		// 기본 타임존으로 먼저 시도
		t, err := time.ParseInLocation(layout, s, defaultLocation)
		if err == nil {
			return t, nil
		}

		// Try without location
		// 타임존 없이 시도
		t, err = time.Parse(layout, s)
		if err == nil {
			return t.In(defaultLocation), nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time string with any known format: %s", s)
}
