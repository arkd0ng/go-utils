package timeutil

import (
	"strings"
	"time"
)

// =============================================================================
// File: format.go
// Purpose: Time and Date Formatting Operations
// 파일: format.go
// 목적: 시간 및 날짜 형식 지정 연산
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The format.go file provides comprehensive time formatting functions that
// convert time.Time values into human-readable string representations. It
// supports standard formats (ISO8601, RFC3339), custom token-based formatting,
// and specialized Korean language formats. These functions simplify the often
// confusing Go time formatting by providing intuitive function names and
// user-friendly format tokens (YYYY-MM-DD) instead of Go's reference time.
//
// format.go 파일은 time.Time 값을 사람이 읽을 수 있는 문자열 표현으로
// 변환하는 포괄적인 시간 형식 지정 함수를 제공합니다. 표준 형식 (ISO8601,
// RFC3339), 커스텀 토큰 기반 형식 지정 및 특수 한국어 형식을 지원합니다.
// 이러한 함수는 직관적인 함수 이름과 Go의 참조 시간 대신 사용자 친화적
// 형식 토큰 (YYYY-MM-DD)을 제공하여 종종 혼란스러운 Go 시간 형식 지정을
// 단순화합니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Intuitive Names**: Clear function names (FormatDate, FormatTime)
//    **직관적 이름**: 명확한 함수 이름 (FormatDate, FormatTime)
//
// 2. **User-Friendly Tokens**: YYYY-MM-DD instead of 2006-01-02
//    **사용자 친화적 토큰**: 2006-01-02 대신 YYYY-MM-DD
//
// 3. **Standard Formats**: Support common international standards
//    **표준 형식**: 일반적인 국제 표준 지원
//
// 4. **Localization**: Korean language formatting for local market
//    **현지화**: 로컬 시장을 위한 한국어 형식 지정
//
// 5. **Default Timezone**: All functions use KST unless specified
//    **기본 타임존**: 별도 지정 없으면 모든 함수가 KST 사용
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. STANDARD FORMATS (표준 형식)
//    - FormatISO8601: ISO 8601 format (2006-01-02T15:04:05Z07:00)
//      FormatISO8601: ISO 8601 형식 (2006-01-02T15:04:05Z07:00)
//    - FormatRFC3339: RFC 3339 format (same as ISO8601)
//      FormatRFC3339: RFC 3339 형식 (ISO8601과 동일)
//
// 2. COMMON FORMATS (일반 형식)
//    - FormatDate: Date only (YYYY-MM-DD)
//      FormatDate: 날짜만 (YYYY-MM-DD)
//    - FormatDateTime: Date and time (YYYY-MM-DD HH:mm:ss)
//      FormatDateTime: 날짜 및 시간 (YYYY-MM-DD HH:mm:ss)
//    - FormatTime: Time only (HH:mm:ss)
//      FormatTime: 시간만 (HH:mm:ss)
//
// 3. CUSTOM FORMATTING (커스텀 형식 지정)
//    - Format: Custom token-based formatting (YYYY-MM-DD HH:mm:ss)
//      Format: 커스텀 토큰 기반 형식 지정 (YYYY-MM-DD HH:mm:ss)
//    - FormatCustom: Alias for Format
//      FormatCustom: Format의 별칭
//
// 4. TIMEZONE-SPECIFIC (타임존 특정)
//    - FormatWithTimezone: Format in specific timezone
//      FormatWithTimezone: 특정 타임존에서 형식 지정
//
// 5. KOREAN FORMATS (한국어 형식)
//    - FormatKorean: Full Korean format (YYYY년 MM월 DD일 HH시 mm분 ss초)
//      FormatKorean: 전체 한국어 형식 (YYYY년 MM월 DD일 HH시 mm분 ss초)
//    - FormatKoreanDate: Korean date (YYYY년 MM월 DD일)
//      FormatKoreanDate: 한국어 날짜 (YYYY년 MM월 DD일)
//    - FormatKoreanDateTime: Korean with weekday
//      FormatKoreanDateTime: 요일 포함 한국어
//    - FormatKoreanDateWithWeekday: Date with full weekday
//      FormatKoreanDateWithWeekday: 전체 요일 포함 날짜
//    - FormatKoreanDateShort: Date with short weekday
//      FormatKoreanDateShort: 짧은 요일 포함 날짜
//
// 6. WEEKDAY HELPERS (요일 헬퍼)
//    - WeekdayKorean: Full Korean weekday name (월요일)
//      WeekdayKorean: 전체 한국어 요일 이름 (월요일)
//    - WeekdayKoreanShort: Short Korean weekday (월)
//      WeekdayKoreanShort: 짧은 한국어 요일 (월)
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// FormatISO8601(t time.Time) string
// - Purpose: Format time in ISO 8601 international standard
// - 목적: ISO 8601 국제 표준으로 시간 형식 지정
// - Format: YYYY-MM-DDTHH:mm:ss±hh:mm
// - 형식: YYYY-MM-DDTHH:mm:ss±hh:mm
// - Example Output: "2024-03-15T14:30:00+09:00"
// - 예제 출력: "2024-03-15T14:30:00+09:00"
// - Timezone: Converts to KST (default timezone)
// - 타임존: KST (기본 타임존)로 변환
// - Use Cases: API responses, data interchange, logging, international standards
// - 사용 사례: API 응답, 데이터 교환, 로깅, 국제 표준
//
// FormatRFC3339(t time.Time) string
// - Purpose: Format time in RFC 3339 standard (same as ISO8601)
// - 목적: RFC 3339 표준으로 시간 형식 지정 (ISO8601과 동일)
// - Standard: RFC 3339 (subset of ISO 8601)
// - 표준: RFC 3339 (ISO 8601의 하위 집합)
// - Use Cases: Internet timestamps, JSON serialization
// - 사용 사례: 인터넷 타임스탬프, JSON 직렬화
//
// FormatDate(t time.Time) string
// - Purpose: Format date only (no time component)
// - 목적: 날짜만 형식 지정 (시간 구성요소 없음)
// - Format: YYYY-MM-DD
// - 형식: YYYY-MM-DD
// - Example Output: "2024-03-15"
// - 예제 출력: "2024-03-15"
// - Use Cases: Birthday fields, calendar dates, date-only storage
// - 사용 사례: 생년월일 필드, 달력 날짜, 날짜만 저장
//
// FormatDateTime(t time.Time) string
// - Purpose: Format date and time in human-readable format
// - 목적: 사람이 읽을 수 있는 형식으로 날짜 및 시간 형식 지정
// - Format: YYYY-MM-DD HH:mm:ss
// - 형식: YYYY-MM-DD HH:mm:ss
// - Example Output: "2024-03-15 14:30:45"
// - 예제 출력: "2024-03-15 14:30:45"
// - Use Cases: Log timestamps, display dates, database storage
// - 사용 사례: 로그 타임스탬프, 날짜 표시, 데이터베이스 저장
//
// FormatTime(t time.Time) string
// - Purpose: Format time only (no date component)
// - 목적: 시간만 형식 지정 (날짜 구성요소 없음)
// - Format: HH:mm:ss
// - 형식: HH:mm:ss
// - Example Output: "14:30:45"
// - 예제 출력: "14:30:45"
// - Use Cases: Clock displays, time-only fields, scheduling
// - 사용 사례: 시계 표시, 시간만 필드, 일정 관리
//
// Format(t time.Time, layout string) string
// - Purpose: Custom formatting with user-friendly tokens
// - 목적: 사용자 친화적 토큰으로 커스텀 형식 지정
// - Tokens: YYYY, YY, MM, M, DD, D, HH, hh, mm, ss
// - 토큰: YYYY, YY, MM, M, DD, D, HH, hh, mm, ss
// - Example Input: "YYYY-MM-DD HH:mm:ss"
// - 예제 입력: "YYYY-MM-DD HH:mm:ss"
// - Conversion: Tokens converted to Go's reference time
// - 변환: 토큰이 Go의 참조 시간으로 변환
// - Example: "YYYY-MM-DD" → "2006-01-02"
// - 예: "YYYY-MM-DD" → "2006-01-02"
// - Use Cases: Custom display formats, user-defined formatting, flexible layouts
// - 사용 사례: 커스텀 표시 형식, 사용자 정의 형식 지정, 유연한 레이아웃
//
// FormatWithTimezone(t time.Time, tz string) (string, error)
// - Purpose: Format time in specific timezone
// - 목적: 특정 타임존에서 시간 형식 지정
// - Timezone: Any IANA timezone name (e.g., "America/New_York")
// - 타임존: 모든 IANA 타임존 이름 (예: "America/New_York")
// - Format: Uses FormatDateTime format (YYYY-MM-DD HH:mm:ss)
// - 형식: FormatDateTime 형식 사용 (YYYY-MM-DD HH:mm:ss)
// - Error Handling: Returns error if timezone invalid
// - 오류 처리: 타임존이 잘못되면 오류 반환
// - Use Cases: Multi-timezone applications, global services
// - 사용 사례: 다중 타임존 애플리케이션, 글로벌 서비스
//
// FormatKorean(t time.Time) string
// - Purpose: Full Korean language format
// - 목적: 전체 한국어 형식
// - Format: YYYY년 MM월 DD일 HH시 mm분 ss초
// - 형식: YYYY년 MM월 DD일 HH시 mm분 ss초
// - Example Output: "2024년 03월 15일 14시 30분 45초"
// - 예제 출력: "2024년 03월 15일 14시 30분 45초"
// - Use Cases: Korean UI, local applications, Korean users
// - 사용 사례: 한국어 UI, 로컬 애플리케이션, 한국 사용자
//
// FormatKoreanDate(t time.Time) string
// - Purpose: Korean date format (date only)
// - 목적: 한국어 날짜 형식 (날짜만)
// - Format: YYYY년 MM월 DD일
// - 형식: YYYY년 MM월 DD일
// - Example Output: "2024년 03월 15일"
// - 예제 출력: "2024년 03월 15일"
// - Use Cases: Date displays, calendar labels
// - 사용 사례: 날짜 표시, 달력 레이블
//
// FormatKoreanDateTime(t time.Time) string
// - Purpose: Korean format with full weekday name
// - 목적: 전체 요일 이름 포함 한국어 형식
// - Format: YYYY년 MM월 DD일 (요일) HH시 mm분 ss초
// - 형식: YYYY년 MM월 DD일 (요일) HH시 mm분 ss초
// - Example Output: "2024년 03월 15일 (금요일) 14시 30분 45초"
// - 예제 출력: "2024년 03월 15일 (금요일) 14시 30분 45초"
// - Use Cases: Detailed Korean displays, reports
// - 사용 사례: 상세 한국어 표시, 보고서
//
// FormatKoreanDateWithWeekday(t time.Time) string
// - Purpose: Korean date with full weekday name
// - 목적: 전체 요일 이름 포함 한국어 날짜
// - Format: YYYY년 MM월 DD일 (요일)
// - 형식: YYYY년 MM월 DD일 (요일)
// - Example Output: "2024년 03월 15일 (금요일)"
// - 예제 출력: "2024년 03월 15일 (금요일)"
// - Use Cases: Calendar displays, schedules
// - 사용 사례: 달력 표시, 일정
//
// FormatKoreanDateShort(t time.Time) string
// - Purpose: Korean date with short weekday
// - 목적: 짧은 요일 포함 한국어 날짜
// - Format: YYYY년 MM월 DD일 (요일)
// - 형식: YYYY년 MM월 DD일 (요일)
// - Example Output: "2024년 03월 15일 (금)"
// - 예제 출력: "2024년 03월 15일 (금)"
// - Use Cases: Compact displays, mobile UI
// - 사용 사례: 컴팩트 표시, 모바일 UI
//
// WeekdayKorean(t time.Time) string
// - Purpose: Get full Korean weekday name
// - 목적: 전체 한국어 요일 이름 가져오기
// - Return Values: 일요일, 월요일, 화요일, 수요일, 목요일, 금요일, 토요일
// - 반환 값: 일요일, 월요일, 화요일, 수요일, 목요일, 금요일, 토요일
// - Use Cases: Weekday displays, custom Korean formatting
// - 사용 사례: 요일 표시, 커스텀 한국어 형식 지정
//
// WeekdayKoreanShort(t time.Time) string
// - Purpose: Get short Korean weekday name
// - 목적: 짧은 한국어 요일 이름 가져오기
// - Return Values: 일, 월, 화, 수, 목, 금, 토
// - 반환 값: 일, 월, 화, 수, 목, 금, 토
// - Use Cases: Compact weekday displays, abbreviations
// - 사용 사례: 컴팩트 요일 표시, 약어
//
// CUSTOM FORMAT TOKENS
// 커스텀 형식 토큰
// --------------------
// The Format() function supports user-friendly tokens:
// Format() 함수는 사용자 친화적 토큰 지원:
//
// Token | Meaning          | Go Equivalent | Example
// ------|------------------|---------------|--------
// YYYY  | 4-digit year     | 2006          | 2024
// YY    | 2-digit year     | 06            | 24
// MM    | 2-digit month    | 01            | 03
// M     | 1-2 digit month  | 1             | 3
// DD    | 2-digit day      | 02            | 15
// D     | 1-2 digit day    | 2             | 15
// HH    | 2-digit hour 24h | 15            | 14
// hh    | 2-digit hour 12h | 03            | 02
// mm    | 2-digit minute   | 04            | 30
// ss    | 2-digit second   | 05            | 45
//
// Why Custom Tokens?
// 왜 커스텀 토큰?
// - Go's reference time (2006-01-02 15:04:05) is confusing
//   Go의 참조 시간 (2006-01-02 15:04:05)은 혼란스러움
// - YYYY-MM-DD is more intuitive and familiar to users
//   YYYY-MM-DD가 사용자에게 더 직관적이고 친숙함
// - Matches format conventions from other languages (PHP, JavaScript)
//   다른 언어 (PHP, JavaScript)의 형식 규칙과 일치
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - All format functions: O(n) where n is format string length
//   모든 형식 함수: O(n), n은 형식 문자열 길이
// - Format (token replacement): O(n * t) where t is number of tokens
//   Format (토큰 대체): O(n * t), t는 토큰 개수
//
// Space Complexities:
// 공간 복잡도:
// - All functions: O(n) - creates new formatted string
//   모든 함수: O(n) - 새 형식 지정 문자열 생성
//
// Optimization Tips:
// 최적화 팁:
// 1. For repeated formatting, cache format string conversion
//    반복 형식 지정의 경우 형식 문자열 변환 캐시
// 2. Standard formats (FormatDate, FormatTime) are faster than Format()
//    표준 형식 (FormatDate, FormatTime)이 Format()보다 빠름
// 3. Pre-convert custom tokens to Go layout if using repeatedly
//    반복 사용 시 커스텀 토큰을 Go 레이아웃으로 사전 변환
// 4. Korean weekday functions use array lookup (O(1))
//    한국어 요일 함수는 배열 조회 사용 (O(1))
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. API Response Formatting
//    API 응답 형식 지정:
//
//    response := map[string]string{
//        "created_at": timeutil.FormatISO8601(time.Now()),
//    }
//    // {"created_at": "2024-03-15T14:30:00+09:00"}
//    // International standard format
//    // 국제 표준 형식
//
// 2. Log Timestamp
//    로그 타임스탬프:
//
//    log.Printf("[%s] User logged in", timeutil.FormatDateTime(time.Now()))
//    // [2024-03-15 14:30:45] User logged in
//    // Human-readable timestamp
//    // 사람이 읽을 수 있는 타임스탬프
//
// 3. Date-Only Display
//    날짜만 표시:
//
//    birthday := timeutil.FormatDate(user.BirthDate)
//    // "1990-05-20"
//    // No time component
//    // 시간 구성요소 없음
//
// 4. Custom Format
//    커스텀 형식:
//
//    formatted := timeutil.Format(time.Now(), "YYYY/MM/DD HH:mm")
//    // "2024/03/15 14:30"
//    // Custom separator and precision
//    // 커스텀 구분자 및 정밀도
//
// 5. Korean UI Display
//    한국어 UI 표시:
//
//    display := timeutil.FormatKoreanDateWithWeekday(time.Now())
//    // "2024년 03월 15일 (금요일)"
//    // Localized format
//    // 현지화된 형식
//
// 6. Multi-Timezone Display
//    다중 타임존 표시:
//
//    nyTime, _ := timeutil.FormatWithTimezone(time.Now(), "America/New_York")
//    // "2024-03-15 01:30:45" (EST)
//    // Convert and format in one call
//    // 한 번의 호출로 변환 및 형식 지정
//
// 7. Weekday-Only Display
//    요일만 표시:
//
//    weekday := timeutil.WeekdayKorean(time.Now())
//    fmt.Printf("오늘은 %s입니다", weekday)
//    // "오늘은 금요일입니다"
//    // Weekday name only
//    // 요일 이름만
//
// 8. Short Date Format
//    짧은 날짜 형식:
//
//    shortDate := timeutil.Format(time.Now(), "YY/MM/DD")
//    // "24/03/15"
//    // 2-digit year
//    // 2자리 연도
//
// 9. Time-Only Clock
//    시간만 시계:
//
//    currentTime := timeutil.FormatTime(time.Now())
//    fmt.Println("Current time:", currentTime)
//    // "Current time: 14:30:45"
//    // Clock display
//    // 시계 표시
//
// 10. Full Korean DateTime
//     전체 한국어 날짜시간:
//
//     fullFormat := timeutil.FormatKorean(time.Now())
//     // "2024년 03월 15일 14시 30분 45초"
//     // Complete Korean format
//     // 완전한 한국어 형식
//
// TIMEZONE HANDLING
// 타임존 처리
// -----------------
// All formatting functions convert time to KST (default timezone) before
// formatting, unless otherwise specified:
//
// 모든 형식 지정 함수는 별도 지정이 없으면 형식 지정 전에 시간을 KST
// (기본 타임존)로 변환합니다:
//
// Example:
// 예:
//     utcTime := time.Now().UTC()
//     formatted := timeutil.FormatDateTime(utcTime)
//     // Automatically converted to KST before formatting
//     // 형식 지정 전에 자동으로 KST로 변환
//
// For explicit timezone formatting:
// 명시적 타임존 형식 지정:
//     formatted, _ := timeutil.FormatWithTimezone(time.Now(), "America/New_York")
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All formatting functions are thread-safe as they operate on immutable
// time.Time values and use read-only constants/maps.
//
// 모든 형식 지정 함수는 불변 time.Time 값에서 작동하고 읽기 전용 상수/맵을
// 사용하므로 스레드 안전합니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     go func() {
//         formatted := timeutil.FormatDateTime(time.Now())
//     }()
//
//     go func() {
//         korean := timeutil.FormatKorean(time.Now())
//     }()
//
//     // All format functions safe for concurrent use
//     // 모든 형식 함수는 동시 사용에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - constants.go: Format layouts and custom tokens
//   constants.go: 형식 레이아웃 및 커스텀 토큰
// - parse.go: Parsing (inverse of formatting)
//   parse.go: 파싱 (형식 지정의 역)
// - timezone.go: Timezone conversion
//   timezone.go: 타임존 변환
//
// =============================================================================

// FormatISO8601 formats a time in ISO8601 format.
// FormatISO8601은 시간을 ISO8601 포맷으로 포맷합니다.
func FormatISO8601(t time.Time) string {
	return t.In(defaultLocation).Format(ISO8601Layout)
}

// FormatRFC3339 formats a time in RFC3339 format.
// FormatRFC3339는 시간을 RFC3339 포맷으로 포맷합니다.
func FormatRFC3339(t time.Time) string {
	return t.In(defaultLocation).Format(RFC3339Layout)
}

// FormatDate formats a time as date only (YYYY-MM-DD).
// FormatDate는 시간을 날짜만으로 포맷합니다 (YYYY-MM-DD).
func FormatDate(t time.Time) string {
	return t.In(defaultLocation).Format(DateLayout)
}

// FormatDateTime formats a time as date and time (YYYY-MM-DD HH:mm:ss).
// FormatDateTime은 시간을 날짜 및 시간으로 포맷합니다 (YYYY-MM-DD HH:mm:ss).
func FormatDateTime(t time.Time) string {
	return t.In(defaultLocation).Format(DateTimeLayout)
}

// FormatTime formats a time as time only (HH:mm:ss).
// FormatTime은 시간을 시간만으로 포맷합니다 (HH:mm:ss).
func FormatTime(t time.Time) string {
	return t.In(defaultLocation).Format(TimeLayout)
}

// Format formats a time using custom format tokens.
// Format은 커스텀 포맷 토큰을 사용하여 시간을 포맷합니다.
//
// Supported tokens
// 지원되는 토큰:
//
//	YYYY - 4-digit year
//	YY   - 2-digit year
//	MM   - 2-digit month
//	M    - 1 or 2-digit month
//	DD   - 2-digit day
//	D    - 1 or 2-digit day
//	HH   - 2-digit hour (24h)
//	hh   - 2-digit hour (12h)
//	mm   - 2-digit minute
//	ss   - 2-digit second
//
// Example
// 예제:
//
// timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
// timeutil.Format(time.Now(), "YYYY년 MM월 DD일")
func Format(t time.Time, layout string) string {
	t = t.In(defaultLocation)
	goLayout := layout
	for token, goToken := range customFormatTokens {
		goLayout = strings.ReplaceAll(goLayout, token, goToken)
	}
	return t.Format(goLayout)
}

// FormatCustom is an alias for Format.
// FormatCustom은 Format의 별칭입니다.
func FormatCustom(t time.Time, layout string) string {
	return Format(t, layout)
}

// FormatWithTimezone formats a time in a specific timezone.
// FormatWithTimezone은 특정 타임존에서 시간을 포맷합니다.
func FormatWithTimezone(t time.Time, tz string) (string, error) {
	converted, err := ConvertTimezone(t, tz)
	if err != nil {
		return "", err
	}
	return FormatDateTime(converted), nil
}

// FormatKorean formats a time in Korean format (YYYY년 MM월 DD일 HH시 mm분 ss초).
// FormatKorean은 시간을 한국어 포맷으로 포맷합니다 (YYYY년 MM월 DD일 HH시 mm분 ss초).
func FormatKorean(t time.Time) string {
	return Format(t, "YYYY년 MM월 DD일 HH시 mm분 ss초")
}

// FormatKoreanDate formats a time in Korean date format (YYYY년 MM월 DD일).
// FormatKoreanDate는 시간을 한국어 날짜 포맷으로 포맷합니다 (YYYY년 MM월 DD일).
func FormatKoreanDate(t time.Time) string {
	return Format(t, "YYYY년 MM월 DD일")
}

// WeekdayKorean returns the Korean name of the weekday.
// WeekdayKorean은 요일의 한글 이름을 반환합니다.
//
// Returns
// 반환값:
//   - "일요일" for Sunday
//   - "월요일" for Monday
//   - "화요일" for Tuesday
//   - "수요일" for Wednesday
//   - "목요일" for Thursday
//   - "금요일" for Friday
//   - "토요일" for Saturday
//
// Example
// 예제:
//
// t := time.Now()
// fmt.Println(timeutil.WeekdayKorean(t))  // Output: 월요일
func WeekdayKorean(t time.Time) string {
	weekdays := []string{
		"일요일", // Sunday
		"월요일", // Monday
		"화요일", // Tuesday
		"수요일", // Wednesday
		"목요일", // Thursday
		"금요일", // Friday
		"토요일", // Saturday
	}
	return weekdays[t.Weekday()]
}

// WeekdayKoreanShort returns the short Korean name of the weekday.
// WeekdayKoreanShort는 요일의 짧은 한글 이름을 반환합니다.
//
// Returns
// 반환값:
//   - "일" for Sunday
//   - "월" for Monday
//   - "화" for Tuesday
//   - "수" for Wednesday
//   - "목" for Thursday
//   - "금" for Friday
//   - "토" for Saturday
//
// Example
// 예제:
//
// t := time.Now()
// fmt.Println(timeutil.WeekdayKoreanShort(t))  // Output: 월
func WeekdayKoreanShort(t time.Time) string {
	weekdays := []string{
		"일", // Sunday
		"월", // Monday
		"화", // Tuesday
		"수", // Wednesday
		"목", // Thursday
		"금", // Friday
		"토", // Saturday
	}
	return weekdays[t.Weekday()]
}

// FormatKoreanDateTime formats a time in Korean format with weekday.
// FormatKoreanDateTime은 요일을 포함한 한국어 포맷으로 시간을 포맷합니다.
//
// Format
// 포맷: YYYY년 MM월 DD일 (요일) HH시 mm분 ss초
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 15, 30, 0, 0, time.UTC)
//
// fmt.Println(timeutil.FormatKoreanDateTime(t))
// Output: 2025년 10월 14일 (화요일) 15시 30분 00초
func FormatKoreanDateTime(t time.Time) string {
	t = t.In(defaultLocation)
	return t.Format("2006년 01월 02일") + " (" + WeekdayKorean(t) + ") " + t.Format("15시 04분 05초")
}

// FormatKoreanDateWithWeekday formats a date in Korean format with weekday.
// FormatKoreanDateWithWeekday는 요일을 포함한 한국어 날짜 포맷으로 포맷합니다.
//
// Format
// 포맷: YYYY년 MM월 DD일 (요일)
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, time.UTC)
//
// fmt.Println(timeutil.FormatKoreanDateWithWeekday(t))
// Output: 2025년 10월 14일 (화요일)
func FormatKoreanDateWithWeekday(t time.Time) string {
	t = t.In(defaultLocation)
	return t.Format("2006년 01월 02일") + " (" + WeekdayKorean(t) + ")"
}

// FormatKoreanDateShort formats a date in Korean format with short weekday.
// FormatKoreanDateShort는 짧은 요일을 포함한 한국어 날짜 포맷으로 포맷합니다.
//
// Format
// 포맷: YYYY년 MM월 DD일 (요일)
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, time.UTC)
//
// fmt.Println(timeutil.FormatKoreanDateShort(t))
// Output: 2025년 10월 14일 (화)
func FormatKoreanDateShort(t time.Time) string {
	t = t.In(defaultLocation)
	return t.Format("2006년 01월 02일") + " (" + WeekdayKoreanShort(t) + ")"
}
