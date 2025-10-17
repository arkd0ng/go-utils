package timeutil

import "time"

// =============================================================================
// File: constants.go
// Purpose: Time and Date Constants and Format Definitions
// 파일: constants.go
// 목적: 시간 및 날짜 상수와 형식 정의
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The constants.go file provides centralized definitions for time-related
// constants, timezone settings, and format layouts used throughout the timeutil
// package. These constants ensure consistency, reduce magic numbers, improve
// code readability, and provide a single source of truth for common time
// operations and formatting patterns.
//
// constants.go 파일은 timeutil 패키지 전체에서 사용되는 시간 관련 상수,
// 타임존 설정 및 형식 레이아웃에 대한 중앙 집중식 정의를 제공합니다. 이러한
// 상수는 일관성을 보장하고 매직 넘버를 줄이며 코드 가독성을 향상시키고
// 일반적인 시간 연산 및 형식 패턴에 대한 단일 정보 소스를 제공합니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Single Source of Truth**: All constants defined in one place
//    **단일 정보 소스**: 모든 상수를 한 곳에 정의
//
// 2. **No Magic Numbers**: Named constants instead of hardcoded values
//    **매직 넘버 없음**: 하드코딩된 값 대신 명명된 상수
//
// 3. **Default Timezone**: KST (Asia/Seoul) as default for Korean context
//    **기본 타임존**: 한국 컨텍스트를 위한 기본값으로 KST (Asia/Seoul)
//
// 4. **Standard Formats**: Common format layouts for consistent formatting
//    **표준 형식**: 일관된 형식 지정을 위한 일반 형식 레이아웃
//
// 5. **User-Friendly Tokens**: Custom format tokens (YYYY, MM, DD) for easier use
//    **사용자 친화적 토큰**: 더 쉬운 사용을 위한 커스텀 형식 토큰 (YYYY, MM, DD)
//
// CONSTANT CATEGORIES
// 상수 범주
// --------------------
//
// 1. DEFAULT TIMEZONE (기본 타임존)
//    - DefaultTimezone: "Asia/Seoul" (KST, GMT+9)
//      DefaultTimezone: "Asia/Seoul" (KST, GMT+9)
//    - DefaultLocation: Alias for DefaultTimezone
//      DefaultLocation: DefaultTimezone의 별칭
//    - KST: *time.Location for Asia/Seoul
//      KST: Asia/Seoul의 *time.Location
//
// 2. TIME CONVERSION CONSTANTS (시간 변환 상수)
//    - SecondsPerMinute: 60
//    - SecondsPerHour: 3600
//    - SecondsPerDay: 86400
//    - MinutesPerHour: 60
//    - HoursPerDay: 24
//
// 3. DATE CONSTANTS (날짜 상수)
//    - DaysPerWeek: 7
//    - DaysPerYear: 365
//    - DaysPerLeapYear: 366
//    - MonthsPerYear: 12
//    - MonthsPerQuarter: 3
//    - QuartersPerYear: 4
//    - WeeksPerYear: 52
//
// 4. STANDARD FORMAT LAYOUTS (표준 형식 레이아웃)
//    - ISO8601Layout: "2006-01-02T15:04:05Z07:00"
//    - RFC3339Layout: Same as ISO8601
//    - DateLayout: "2006-01-02"
//    - DateTimeLayout: "2006-01-02 15:04:05"
//    - TimeLayout: "15:04:05"
//
// 5. CUSTOM FORMAT TOKENS (커스텀 형식 토큰)
//    - customFormatTokens: Map of user-friendly tokens to Go layouts
//      customFormatTokens: 사용자 친화적 토큰과 Go 레이아웃의 맵
//
// KEY CONSTANTS SUMMARY
// 주요 상수 요약
// ---------------------
//
// DefaultTimezone / DefaultLocation
// - Value: "Asia/Seoul"
// - 값: "Asia/Seoul"
// - Purpose: Default timezone for all timeutil operations
// - 목적: 모든 timeutil 연산의 기본 타임존
// - Timezone: KST (Korea Standard Time), GMT+9
// - 타임존: KST (한국 표준시), GMT+9
// - Why KST: Package designed for Korean market/users
// - KST 이유: 한국 시장/사용자를 위해 설계된 패키지
// - Fallback: UTC if KST cannot be loaded
// - 폴백: KST를 로드할 수 없으면 UTC
//
// KST (*time.Location)
// - Type: *time.Location
// - 타입: *time.Location
// - Initialized: In init() function
// - 초기화: init() 함수에서
// - Purpose: Pre-loaded timezone location for performance
// - 목적: 성능을 위해 미리 로드된 타임존 위치
// - Usage: time.Now().In(KST)
// - 사용: time.Now().In(KST)
//
// Time Conversion Constants
// 시간 변환 상수
// - SecondsPerMinute: 60 (1 minute = 60 seconds)
//   SecondsPerMinute: 60 (1분 = 60초)
// - SecondsPerHour: 3600 (1 hour = 3600 seconds)
//   SecondsPerHour: 3600 (1시간 = 3600초)
// - SecondsPerDay: 86400 (1 day = 86400 seconds)
//   SecondsPerDay: 86400 (1일 = 86400초)
// - MinutesPerHour: 60 (1 hour = 60 minutes)
//   MinutesPerHour: 60 (1시간 = 60분)
// - HoursPerDay: 24 (1 day = 24 hours)
//   HoursPerDay: 24 (1일 = 24시간)
// - Use Cases: Time unit conversions, duration calculations
//   사용 사례: 시간 단위 변환, 기간 계산
//
// Date Constants
// 날짜 상수
// - DaysPerWeek: 7 (1 week = 7 days)
//   DaysPerWeek: 7 (1주 = 7일)
// - DaysPerYear: 365 (non-leap year)
//   DaysPerYear: 365 (윤년 아님)
// - DaysPerLeapYear: 366 (leap year)
//   DaysPerLeapYear: 366 (윤년)
// - MonthsPerYear: 12 (1 year = 12 months)
//   MonthsPerYear: 12 (1년 = 12개월)
// - MonthsPerQuarter: 3 (1 quarter = 3 months)
//   MonthsPerQuarter: 3 (1분기 = 3개월)
// - QuartersPerYear: 4 (1 year = 4 quarters)
//   QuartersPerYear: 4 (1년 = 4분기)
// - WeeksPerYear: 52 (approximate, actual: 52.14)
//   WeeksPerYear: 52 (근사값, 실제: 52.14)
// - Use Cases: Date calculations, calendar operations
//   사용 사례: 날짜 계산, 달력 연산
//
// ISO8601Layout / RFC3339Layout
// - Value: "2006-01-02T15:04:05Z07:00"
// - 값: "2006-01-02T15:04:05Z07:00"
// - Standard: ISO 8601 / RFC 3339 international standard
// - 표준: ISO 8601 / RFC 3339 국제 표준
// - Format: YYYY-MM-DDTHH:mm:ss±hh:mm
// - 형식: YYYY-MM-DDTHH:mm:ss±hh:mm
// - Use Cases: API responses, data interchange, logging
// - 사용 사례: API 응답, 데이터 교환, 로깅
//
// DateLayout
// - Value: "2006-01-02"
// - 값: "2006-01-02"
// - Format: YYYY-MM-DD (date only, no time)
// - 형식: YYYY-MM-DD (날짜만, 시간 없음)
// - Use Cases: Date-only storage, birthday fields, calendar dates
// - 사용 사례: 날짜만 저장, 생년월일 필드, 달력 날짜
//
// DateTimeLayout
// - Value: "2006-01-02 15:04:05"
// - 값: "2006-01-02 15:04:05"
// - Format: YYYY-MM-DD HH:mm:ss (human-readable)
// - 형식: YYYY-MM-DD HH:mm:ss (사람이 읽기 쉬움)
// - Use Cases: Log timestamps, display formats, database storage
// - 사용 사례: 로그 타임스탬프, 표시 형식, 데이터베이스 저장
//
// TimeLayout
// - Value: "15:04:05"
// - 값: "15:04:05"
// - Format: HH:mm:ss (time only, no date)
// - 형식: HH:mm:ss (시간만, 날짜 없음)
// - Use Cases: Time-only fields, clock displays, time comparisons
// - 사용 사례: 시간만 필드, 시계 표시, 시간 비교
//
// Custom Format Tokens (customFormatTokens)
// - YYYY: 4-digit year (2006) → "2006"
//   YYYY: 4자리 연도 (2006) → "2006"
// - YY: 2-digit year (06) → "06"
//   YY: 2자리 연도 (06) → "06"
// - MM: 2-digit month (01-12) → "01"
//   MM: 2자리 월 (01-12) → "01"
// - M: 1-2 digit month (1-12) → "1"
//   M: 1-2자리 월 (1-12) → "1"
// - DD: 2-digit day (01-31) → "02"
//   DD: 2자리 일 (01-31) → "02"
// - D: 1-2 digit day (1-31) → "2"
//   D: 1-2자리 일 (1-31) → "2"
// - HH: 2-digit hour 24h (00-23) → "15"
//   HH: 2자리 시간 24시간 (00-23) → "15"
// - hh: 2-digit hour 12h (01-12) → "03"
//   hh: 2자리 시간 12시간 (01-12) → "03"
// - mm: 2-digit minute (00-59) → "04"
//   mm: 2자리 분 (00-59) → "04"
// - ss: 2-digit second (00-59) → "05"
//   ss: 2자리 초 (00-59) → "05"
// - Purpose: User-friendly format strings (YYYY-MM-DD) instead of Go's reference time
//   목적: Go의 참조 시간 대신 사용자 친화적 형식 문자열 (YYYY-MM-DD)
//
// WHY KST AS DEFAULT?
// 왜 KST가 기본값?
// -------------------
// The timeutil package defaults to KST (Asia/Seoul, GMT+9) because:
// timeutil 패키지가 KST (Asia/Seoul, GMT+9)를 기본값으로 하는 이유:
//
// 1. **Target Audience**: Primarily Korean users and applications
//    **대상 청중**: 주로 한국 사용자 및 애플리케이션
//
// 2. **Convenience**: No need to specify timezone for Korean developers
//    **편의성**: 한국 개발자가 타임존 지정할 필요 없음
//
// 3. **Common Use Case**: Most Go applications in Korea use KST
//    **일반 사용 사례**: 한국의 대부분 Go 애플리케이션은 KST 사용
//
// 4. **Explicit Override**: Users can still specify other timezones when needed
//    **명시적 재정의**: 필요시 사용자가 여전히 다른 타임존 지정 가능
//
// 5. **Consistency**: All functions use same default, avoiding confusion
//    **일관성**: 모든 함수가 같은 기본값 사용, 혼란 방지
//
// Note: For global applications, consider using UTC as default instead.
// 참고: 글로벌 애플리케이션의 경우 대신 UTC를 기본값으로 사용 고려.
//
// GO TIME FORMAT REFERENCE
// Go 시간 형식 참조
// ------------------------
// Go uses a reference time for format layouts: "Mon Jan 2 15:04:05 MST 2006"
// Go는 형식 레이아웃에 참조 시간 사용: "Mon Jan 2 15:04:05 MST 2006"
//
// This represents: 01/02 03:04:05 PM '06 -0700
// 이것은 다음을 나타냄: 01/02 03:04:05 PM '06 -0700
//
// Reference time components:
// 참조 시간 구성요소:
// - Month: 01 (January, 1st month)
//   월: 01 (1월, 첫 번째 월)
// - Day: 02 (2nd day)
//   일: 02 (2일)
// - Hour: 03 (3 PM in 12h) or 15 (in 24h)
//   시간: 03 (12시간 형식 오후 3시) 또는 15 (24시간 형식)
// - Minute: 04 (4th minute)
//   분: 04 (4분)
// - Second: 05 (5th second)
//   초: 05 (5초)
// - Year: 2006 (2006 year)
//   연도: 2006 (2006년)
// - Timezone: -0700 (MST, 7 hours behind UTC)
//   타임존: -0700 (MST, UTC보다 7시간 뒤)
//
// This mnemonic helps remember the format: 1, 2, 3, 4, 5, 6, 7
// 이 니모닉은 형식을 기억하는 데 도움: 1, 2, 3, 4, 5, 6, 7
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Using Default Timezone
//    기본 타임존 사용:
//
//    now := time.Now().In(timeutil.KST)
//    // Current time in KST
//    // KST의 현재 시간
//
// 2. Time Unit Conversion
//    시간 단위 변환:
//
//    hours := seconds / timeutil.SecondsPerHour
//    days := hours / timeutil.HoursPerDay
//    // Convert seconds to days
//    // 초를 일로 변환
//
// 3. Standard Format Usage
//    표준 형식 사용:
//
//    formatted := time.Now().Format(timeutil.ISO8601Layout)
//    // "2024-03-15T14:30:00+09:00"
//    // ISO 8601 format
//    // ISO 8601 형식
//
// 4. Date-Only Formatting
//    날짜만 형식 지정:
//
//    dateStr := time.Now().Format(timeutil.DateLayout)
//    // "2024-03-15"
//    // Date without time
//    // 시간 없이 날짜만
//
// 5. Custom Format with Tokens
//    토큰으로 커스텀 형식:
//
//    // Internal use by Format() function
//    // Format() 함수의 내부 사용
//    customFormat := "YYYY-MM-DD HH:mm:ss"
//    // Converted to "2006-01-02 15:04:05"
//    // "2006-01-02 15:04:05"로 변환
//
// 6. Calculating Days in Period
//    기간의 일 수 계산:
//
//    days := 2 * timeutil.WeeksPerYear // 104 days (2 weeks)
//    quarterDays := timeutil.DaysPerYear / timeutil.QuartersPerYear
//    // Calculate approximate quarter length
//    // 대략적인 분기 길이 계산
//
// 7. Duration Calculations
//    기간 계산:
//
//    totalSeconds := 3 * timeutil.HoursPerDay * timeutil.SecondsPerHour
//    // Total seconds in 3 days
//    // 3일의 총 초
//
// DESIGN RATIONALE
// 설계 근거
// -----------------
//
// Named Constants vs Magic Numbers:
// 명명된 상수 vs 매직 넘버:
// - Bad: hours := seconds / 3600
//   나쁨: hours := seconds / 3600
// - Good: hours := seconds / timeutil.SecondsPerHour
//   좋음: hours := seconds / timeutil.SecondsPerHour
// - Benefits: Self-documenting, searchable, maintainable
//   이점: 자체 문서화, 검색 가능, 유지 관리 가능
//
// Centralized Formats:
// 중앙 집중식 형식:
// - Single source for format strings
//   형식 문자열의 단일 소스
// - Easy to update across entire codebase
//   전체 코드베이스에서 업데이트 쉬움
// - Consistent formatting everywhere
//   모든 곳에서 일관된 형식
//
// Pre-loaded Locations:
// 미리 로드된 위치:
// - KST loaded once in init()
//   init()에서 KST 한 번 로드
// - Avoid repeated LoadLocation() calls
//   반복된 LoadLocation() 호출 방지
// - Better performance
//   더 나은 성능
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All constants and pre-loaded locations in this file are immutable and
// thread-safe for concurrent read access.
//
// 이 파일의 모든 상수와 미리 로드된 위치는 불변이며 동시 읽기 액세스에
// 스레드 안전합니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     go func() {
//         formatted := time.Now().Format(timeutil.ISO8601Layout)
//     }()
//
//     go func() {
//         kstTime := time.Now().In(timeutil.KST)
//     }()
//
//     // All constants safe for concurrent read
//     // 모든 상수는 동시 읽기에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - format.go: Uses format layouts and custom tokens
//   format.go: 형식 레이아웃 및 커스텀 토큰 사용
// - parse.go: Uses format layouts for parsing
//   parse.go: 파싱을 위해 형식 레이아웃 사용
// - timezone.go: Uses KST and timezone constants
//   timezone.go: KST 및 타임존 상수 사용
// - arithmetic.go: Uses time conversion constants
//   arithmetic.go: 시간 변환 상수 사용
//
// =============================================================================

// Default timezone
// 기본 타임존
// All functions use KST (Asia/Seoul, GMT+9) as default timezone unless specified.
// 모든 함수는 별도 지정이 없으면 KST (Asia/Seoul, GMT+9)를 기본 타임존으로 사용합니다.
const (
	DefaultTimezone = "Asia/Seoul"
	DefaultLocation = "Asia/Seoul" // KST, GMT+9
)

var (
	// KST is the default timezone location (Asia/Seoul, GMT+9)
	// KST는 기본 타임존 위치입니다 (Asia/Seoul, GMT+9)
	KST *time.Location

	// defaultLocation is the current default location, can be changed
	// defaultLocation은 현재 기본 위치이며 변경 가능합니다
	defaultLocation *time.Location
)

// init loads the default timezone (KST)
// init은 기본 타임존(KST)을 로드합니다
func init() {
	var err error
	KST, err = time.LoadLocation(DefaultTimezone)
	if err != nil {
		// Fallback to UTC if KST cannot be loaded
		// KST를 로드할 수 없으면 UTC로 폴백
		KST = time.UTC
	}
	defaultLocation = KST
}

// Time constants
// 시간 상수
const (
	SecondsPerMinute  = 60
	SecondsPerHour    = 3600
	SecondsPerDay     = 86400
	MinutesPerHour    = 60
	HoursPerDay       = 24
	DaysPerWeek       = 7
	DaysPerYear       = 365
	DaysPerLeapYear   = 366
	MonthsPerYear     = 12
	MonthsPerQuarter  = 3
	QuartersPerYear   = 4
	WeeksPerYear      = 52
)

// Common format layouts
// 일반 포맷 레이아웃
const (
	// ISO8601 format
	// ISO8601 포맷
	ISO8601Layout = "2006-01-02T15:04:05Z07:00"

	// RFC3339 format
	// RFC3339 포맷
	RFC3339Layout = "2006-01-02T15:04:05Z07:00"

	// Date only format
	// 날짜만 포맷
	DateLayout = "2006-01-02"

	// DateTime format
	// 날짜시간 포맷
	DateTimeLayout = "2006-01-02 15:04:05"

	// Time only format
	// 시간만 포맷
	TimeLayout = "15:04:05"
)

// Custom format tokens for user-friendly formatting
// 사용자 친화적 포맷팅을 위한 커스텀 포맷 토큰
// These tokens are translated to Go's standard layout format.
// 이 토큰들은 Go의 표준 레이아웃 포맷으로 변환됩니다.
//
// Supported tokens
// 지원되는 토큰:
//   YYYY - 4-digit year (2006)
//   YY   - 2-digit year (06)
//   MM   - 2-digit month (01-12)
//   M    - 1 or 2-digit month (1-12)
//   DD   - 2-digit day (01-31)
//   D    - 1 or 2-digit day (1-31)
//   HH   - 2-digit hour 24h format (00-23)
//   hh   - 2-digit hour 12h format (01-12)
//   mm   - 2-digit minute (00-59)
//   ss   - 2-digit second (00-59)
var customFormatTokens = map[string]string{
	"YYYY": "2006",
	"YY":   "06",
	"MM":   "01",
	"M":    "1",
	"DD":   "02",
	"D":    "2",
	"HH":   "15",
	"hh":   "03",
	"mm":   "04",
	"ss":   "05",
}
