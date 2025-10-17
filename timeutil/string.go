package timeutil

import (
	"fmt"
	"time"
)

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/string.go
// Purpose: String-based convenience wrappers for time operations
//          시간 연산을 위한 문자열 기반 편의 래퍼
//
// This file provides string-based versions of time operations that accept
// string parameters instead of time.Time. All functions automatically parse
// input strings using ParseAny, which supports 40+ common date/time formats.
// This eliminates the need for manual parsing in application code and provides
// a more convenient API when working with string-based date/time data from
// databases, APIs, user input, or configuration files.
//
// 이 파일은 time.Time 대신 문자열 매개변수를 받는 시간 연산의 문자열 기반
// 버전을 제공합니다. 모든 함수는 40개 이상의 일반적인 날짜/시간 형식을 지원하는
// ParseAny를 사용하여 입력 문자열을 자동으로 파싱합니다. 이는 애플리케이션
// 코드에서 수동 파싱이 필요 없으며 데이터베이스, API, 사용자 입력 또는 구성
// 파일에서 문자열 기반 날짜/시간 데이터를 작업할 때 더 편리한 API를 제공합니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. AUTOMATIC STRING PARSING (자동 문자열 파싱)
//    - All functions use ParseAny internally
//      모든 함수가 내부적으로 ParseAny 사용
//    - Supports 40+ date/time formats
//      40개 이상의 날짜/시간 형식 지원
//    - No manual parsing required
//      수동 파싱 불필요
//
// 2. COMPREHENSIVE OPERATION COVERAGE (포괄적인 연산 범위)
//    - Difference calculations (SubTimeString, DiffIn*String)
//      차이 계산
//    - Age calculations (AgeString, AgeInYearsString)
//      나이 계산
//    - Relative time (RelativeTimeString)
//      상대 시간
//    - Business days (IsBusinessDayString)
//      영업일
//    - Arithmetic (Add*String, Sub*String)
//      산술
//    - Formatting (Format*String)
//      형식화
//    - Timezone conversion (ConvertTimezoneString)
//      타임존 변환
//    - Boundary operations (StartOf*String, EndOf*String)
//      경계 연산
//    - Weekday operations (Weekday*String)
//      평일 연산
//    - Week/month/year operations (WeekOf*String, DaysIn*String)
//      주/월/연도 연산
//    - Month operations (Month*String, QuarterString)
//      월 연산
//    - Comparisons (IsSame*String, IsBefore/AfterString, IsBetweenString)
//      비교
//
// 3. ERROR HANDLING (오류 처리)
//    - Returns error if string parsing fails
//      문자열 파싱 실패 시 오류 반환
//    - Wrapped errors with context
//      컨텍스트와 함께 래핑된 오류
//    - Clear error messages
//      명확한 오류 메시지
//
// 4. TYPE FLEXIBILITY (타입 유연성)
//    - Accept strings, return appropriate types
//      문자열 수용, 적절한 타입 반환
//    - Some return time.Time, others return primitives
//      일부는 time.Time 반환, 다른 일부는 기본 타입 반환
//    - Consistent with base function signatures
//      기본 함수 서명과 일치
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. CONVENIENCE OVER PERFORMANCE (성능보다 편의성)
//    - Prioritize ease of use
//      사용 편의성 우선
//    - Parsing overhead acceptable for most use cases
//      대부분의 사용 사례에서 파싱 오버헤드 허용
//    - Ideal for web applications, scripts, configuration
//      웹 애플리케이션, 스크립트, 구성에 이상적
//
// 2. CONSISTENT NAMING (일관된 명명)
//    - Base function name + "String" suffix
//      기본 함수 이름 + "String" 접미사
//    - Example: AddDays → AddDaysString
//      예시: AddDays → AddDaysString
//    - Easy to discover and remember
//      발견 및 기억 용이
//
// 3. PARSEANY FOR FLEXIBILITY (유연성을 위한 ParseAny)
//    - Uses most flexible parser (ParseAny)
//      가장 유연한 파서 사용 (ParseAny)
//    - Accepts many date/time formats
//      많은 날짜/시간 형식 수용
//    - No need to know exact format
//      정확한 형식을 알 필요 없음
//
// 4. PRESERVE BASE FUNCTION BEHAVIOR (기본 함수 동작 보존)
//    - String wrappers delegate to base functions
//      문자열 래퍼가 기본 함수에 위임
//    - Same logic, same results
//      동일한 로직, 동일한 결과
//    - Only difference is input parsing
//      유일한 차이는 입력 파싱
//
// 5. ERROR PROPAGATION (오류 전파)
//    - Parse errors returned with context
//      컨텍스트와 함께 파싱 오류 반환
//    - Base function errors also propagated
//      기본 함수 오류도 전파
//    - Clear which string failed to parse
//      어떤 문자열이 파싱에 실패했는지 명확
//
// 6. NO ASSUMPTIONS ABOUT FORMAT (형식에 대한 가정 없음)
//    - Don't require specific format
//      특정 형식 요구하지 않음
//    - Let ParseAny handle format detection
//      ParseAny가 형식 감지 처리하도록 허용
//    - Flexible for diverse data sources
//      다양한 데이터 소스에 유연
//
// ============================================================================
// STRING OPERATIONS OVERVIEW / 문자열 연산 개요
// ============================================================================
//
// DIFFERENCE & AGE (차이 및 나이) - 7 functions
// ├─ SubTimeString         : Structured time difference
// ├─ DiffInDaysString      : Difference in days
// ├─ DiffInHoursString     : Difference in hours
// ├─ DiffInMinutesString   : Difference in minutes
// ├─ AgeString             : Detailed age from birth date
// ├─ AgeInYearsString      : Age in years
// └─ RelativeTimeString    : Human-readable relative time
//
// BUSINESS DAYS & WEEKENDS (영업일 및 주말) - 2 functions
// ├─ IsBusinessDayString : Check if business day
// └─ IsWeekendString     : Check if weekend
//
// ARITHMETIC (산술) - 6 functions
// ├─ AddDaysString       : Add days
// ├─ AddHoursString      : Add hours
// ├─ AddMinutesString    : Add minutes
// ├─ SubDaysString       : Subtract days
// ├─ SubHoursString      : Subtract hours
// └─ SubMinutesString    : Subtract minutes
//
// FORMATTING (형식화) - 4 functions
// ├─ FormatString         : Format with custom layout
// ├─ FormatDateString     : Format as YYYY-MM-DD
// ├─ FormatDateTimeString : Format as YYYY-MM-DD HH:mm:ss
// └─ FormatISO8601String  : Format as ISO 8601
//
// TIMEZONE (타임존) - 1 function
// └─ ConvertTimezoneString : Convert to different timezone
//
// BOUNDARIES (경계) - 8 functions
// ├─ StartOfDayString    : Start of day (00:00:00)
// ├─ EndOfDayString      : End of day (23:59:59)
// ├─ StartOfWeekString   : Start of week (Monday 00:00)
// ├─ EndOfWeekString     : End of week (Sunday 23:59)
// ├─ StartOfMonthString  : Start of month (1st 00:00)
// ├─ EndOfMonthString    : End of month (last 23:59)
// ├─ StartOfYearString   : Start of year (Jan 1 00:00)
// └─ EndOfYearString     : End of year (Dec 31 23:59)
//
// WEEKDAY (평일) - 5 functions
// ├─ WeekdayString            : Weekday name (English)
// ├─ WeekdayKoreanString      : Weekday name (Korean)
// ├─ WeekdayShortString       : Short weekday (3 letters)
// ├─ WeekdayShortKoreanString : Short Korean weekday
// └─ WeekdayNumberString      : Weekday number (0-6)
//
// WEEK & DAY COUNTS (주 및 일 수) - 4 functions
// ├─ WeekOfYearString   : ISO week number (1-53)
// ├─ WeekOfMonthString  : Week of month (1-6)
// ├─ DaysInMonthString  : Days in month (28-31)
// └─ DaysInYearString   : Days in year (365-366)
//
// MONTH & QUARTER (월 및 분기) - 4 functions
// ├─ MonthKoreanString    : Korean month name
// ├─ MonthNameString      : English month name
// ├─ MonthNameShortString : Short English month
// └─ QuarterString        : Quarter (1-4)
//
// COMPARISONS (비교) - 6 functions
// ├─ IsLeapYearString : Check if leap year
// ├─ IsSameDayString  : Check if same day
// ├─ IsBeforeString   : Check if before
// ├─ IsAfterString    : Check if after
// └─ IsBetweenString  : Check if between
//
// Total: 47 string-based functions
// 총: 47개의 문자열 기반 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
//
// ALL FUNCTIONS: O(1) + parsing overhead
//   모든 함수: O(1) + 파싱 오버헤드
//
// PARSING OVERHEAD:
// 파싱 오버헤드:
// - ParseAny tries up to 40+ formats
//   ParseAny는 최대 40개 이상의 형식 시도
// - Early exit on first match
//   첫 매칭 시 조기 종료
// - Common formats matched quickly (~1-10 microseconds)
//   일반 형식은 빠르게 매칭 (~1-10 마이크로초)
// - Uncommon formats may take longer (~10-100 microseconds)
//   드문 형식은 더 오래 걸릴 수 있음 (~10-100 마이크로초)
//
// OPERATION OVERHEAD:
// 연산 오버헤드:
// - After parsing, delegates to base function
//   파싱 후 기본 함수에 위임
// - Base function performance same as non-string version
//   기본 함수 성능은 비문자열 버전과 동일
//
// SPACE COMPLEXITY (공간 복잡도):
// - O(1) for most operations
//   대부분의 연산에서 O(1)
// - Parsing creates temporary time.Time
//   파싱이 임시 time.Time 생성
// - Result allocation depends on operation
//   결과 할당은 연산에 따라 다름
//
// PERFORMANCE COMPARISON:
// 성능 비교:
//   AddDays(t, 5)           :  ~50 ns/op
//   AddDaysString("...", 5) : ~10,000 ns/op (200x slower due to parsing)
//                             파싱으로 인해 200배 느림
//
// WHEN TO USE:
// 사용 시기:
// - Use string versions for: API endpoints, scripts, configuration, user input
//   문자열 버전 사용: API 엔드포인트, 스크립트, 구성, 사용자 입력
// - Use base versions for: Performance-critical loops, high-frequency operations
//   기본 버전 사용: 성능 중요 루프, 고빈도 연산
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: API Request Handling (API 요청 처리)
// Use case: Calculate age from query parameter
// 사용 사례: 쿼리 매개변수에서 나이 계산
//
//   // GET /user/age?birthdate=1990-01-15
//   func handleAgeRequest(w http.ResponseWriter, r *http.Request) {
//       birthDateStr := r.URL.Query().Get("birthdate")
//
//       age, err := timeutil.AgeInYearsString(birthDateStr)
//       if err != nil {
//           http.Error(w, "Invalid birth date", http.StatusBadRequest)
//           return
//       }
//
//       json.NewEncoder(w).Encode(map[string]int{"age": age})
//   }
//
// PATTERN 2: Database Query Result (데이터베이스 쿼리 결과)
// Use case: Format date string from database
// 사용 사례: 데이터베이스에서 날짜 문자열 형식화
//
//   type Event struct {
//       ID        int
//       Name      string
//       StartTime string  // "2024-10-14 15:30:00"
//   }
//
//   event := getEventFromDB(id)
//
//   // Format for display
//   displayTime, _ := timeutil.FormatDateTimeString(event.StartTime)
//   relativeTime, _ := timeutil.RelativeTimeString(event.StartTime)
//
//   fmt.Printf("Event: %s\n", event.Name)
//   fmt.Printf("Time: %s (%s)\n", displayTime, relativeTime)
//
// PATTERN 3: Configuration File (구성 파일)
// Use case: Parse dates from YAML/JSON config
// 사용 사례: YAML/JSON 구성에서 날짜 파싱
//
//   type Config struct {
//       StartDate string `yaml:"start_date"`
//       EndDate   string `yaml:"end_date"`
//   }
//
//   config := loadConfig()
//
//   // Calculate duration
//   diff, err := timeutil.SubTimeString(config.StartDate, config.EndDate)
//   if err != nil {
//       log.Fatalf("Invalid dates in config: %v", err)
//   }
//
//   fmt.Printf("Project duration: %s\n", diff.String())
//
// PATTERN 4: User Input Validation (사용자 입력 검증)
// Use case: Validate date range from form
// 사용 사례: 양식에서 날짜 범위 검증
//
//   func validateDateRange(startStr, endStr string) error {
//       isAfter, err := timeutil.IsAfterString(endStr, startStr)
//       if err != nil {
//           return fmt.Errorf("invalid date format: %w", err)
//       }
//
//       if !isAfter {
//           return fmt.Errorf("end date must be after start date")
//       }
//
//       return nil
//   }
//
// PATTERN 5: Report Generation (보고서 생성)
// Use case: Calculate days between report dates
// 사용 사례: 보고서 날짜 간 일수 계산
//
//   type Report struct {
//       StartDate string
//       EndDate   string
//   }
//
//   report := &Report{
//       StartDate: "2024-01-01",
//       EndDate:   "2024-12-31",
//   }
//
//   days, _ := timeutil.DiffInDaysString(report.StartDate, report.EndDate)
//   fmt.Printf("Report covers %.0f days\n", days)
//
// PATTERN 6: Date Arithmetic from Strings (문자열에서 날짜 산술)
// Use case: Add business days to string date
// 사용 사례: 문자열 날짜에 영업일 추가
//
//   orderDate := "2024-10-14"
//
//   // Add 7 days
//   deliveryDate, err := timeutil.AddDaysString(orderDate, 7)
//   if err != nil {
//       log.Fatal(err)
//   }
//
//   fmt.Printf("Expected delivery: %s\n",
//       timeutil.FormatDate(deliveryDate))
//
// PATTERN 7: Timezone Conversion from String (문자열에서 타임존 변환)
// Use case: Convert meeting time to user timezone
// 사용 사례: 회의 시간을 사용자 타임존으로 변환
//
//   meetingTimeUTC := "2024-10-14 15:00:00"  // UTC
//   userTimezone := "America/New_York"
//
//   localTime, err := timeutil.ConvertTimezoneString(meetingTimeUTC, userTimezone)
//   if err != nil {
//       log.Fatal(err)
//   }
//
//   fmt.Printf("Meeting time in your timezone: %s\n",
//       timeutil.FormatDateTime(localTime))
//
// PATTERN 8: Weekday Localization (평일 현지화)
// Use case: Display weekday in user's language
// 사용 사례: 사용자의 언어로 평일 표시
//
//   dateStr := "2024-10-14"
//   language := "ko"
//
//   var weekday string
//   var err error
//
//   if language == "ko" {
//       weekday, err = timeutil.WeekdayKoreanString(dateStr)
//   } else {
//       weekday, err = timeutil.WeekdayString(dateStr)
//   }
//
//   if err == nil {
//       fmt.Printf("Weekday: %s\n", weekday)
//   }
//
// PATTERN 9: Batch Date Processing (배치 날짜 처리)
// Use case: Process multiple date strings
// 사용 사례: 여러 날짜 문자열 처리
//
//   dates := []string{
//       "2024-01-15",
//       "2024-06-20",
//       "2024-12-31",
//   }
//
//   for _, dateStr := range dates {
//       quarter, err := timeutil.QuarterString(dateStr)
//       if err != nil {
//           log.Printf("Invalid date %s: %v", dateStr, err)
//           continue
//       }
//
//       fmt.Printf("%s is in Q%d\n", dateStr, quarter)
//   }
//
// PATTERN 10: Comparison Operations (비교 연산)
// Use case: Check if date is in range
// 사용 사례: 날짜가 범위 내에 있는지 확인
//
//   func isInPromotionPeriod(dateStr string) (bool, error) {
//       startDate := "2024-10-01"
//       endDate := "2024-10-31"
//
//       return timeutil.IsBetweenString(dateStr, startDate, endDate)
//   }
//
//   valid, err := isInPromotionPeriod("2024-10-15")
//   if err == nil && valid {
//       fmt.Println("Date is in promotion period")
//   }
//
// ============================================================================
// ERROR HANDLING / 오류 처리
// ============================================================================
//
// PARSING ERRORS:
// 파싱 오류:
//
// All string functions can return errors from ParseAny:
// 모든 문자열 함수는 ParseAny에서 오류 반환 가능:
//
//   result, err := timeutil.AddDaysString("invalid date", 5)
//   if err != nil {
//       // Error: "failed to parse date string: ..."
//       // 오류: "날짜 문자열 파싱 실패: ..."
//   }
//
// WRAPPED ERRORS:
// 래핑된 오류:
//
// Errors are wrapped with context:
// 오류는 컨텍스트와 함께 래핑됨:
//
//   - "failed to parse date string: ..."
//   - "failed to parse first time string: ..."
//   - "failed to parse second time string: ..."
//   - "failed to parse birth date: ..."
//
// TIMEZONE ERRORS:
// 타임존 오류:
//
// ConvertTimezoneString can also return timezone errors:
// ConvertTimezoneString은 타임존 오류도 반환 가능:
//
//   result, err := timeutil.ConvertTimezoneString("2024-10-14", "Invalid/Timezone")
//   if err != nil {
//       // Error from ConvertTimezone
//       // ConvertTimezone의 오류
//   }
//
// BEST PRACTICE:
// 모범 사례:
//
// Always check errors:
// 항상 오류 확인:
//
//   result, err := timeutil.AddDaysString(dateStr, 5)
//   if err != nil {
//       log.Printf("Date operation failed: %v", err)
//       return
//   }
//   // Use result
//
// ============================================================================
// EDGE CASES / 경계 사례
// ============================================================================
//
// EMPTY STRING:
// 빈 문자열:
//   result, err := timeutil.FormatDateString("")
//   // Returns parsing error
//   // 파싱 오류 반환
//
// INVALID FORMAT:
// 잘못된 형식:
//   result, err := timeutil.AddDaysString("not a date", 5)
//   // Returns parsing error
//   // 파싱 오류 반환
//
// PARTIALLY VALID:
// 부분적으로 유효:
//   // First string valid, second invalid
//   diff, err := timeutil.SubTimeString("2024-10-14", "invalid")
//   // Returns error for second string
//   // 두 번째 문자열에 대한 오류 반환
//
// TIMEZONE IN STRING:
// 문자열의 타임존:
//   // ParseAny handles timezone in string
//   result, _ := timeutil.FormatDateString("2024-10-14T15:30:00Z")
//   // Works correctly
//   // 올바르게 작동
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// THREAD-SAFE:
// 스레드 안전:
// - All string functions are thread-safe
//   모든 문자열 함수는 스레드 안전
// - ParseAny is thread-safe
//   ParseAny는 스레드 안전
// - Base functions are thread-safe
//   기본 함수는 스레드 안전
// - Safe for concurrent use
//   동시 사용에 안전
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM parse.go:
// - ParseAny: Universal date/time parser
//
// FROM ALL OTHER FILES:
// - Wraps nearly every function in timeutil package
//   timeutil 패키지의 거의 모든 함수 래핑
//
// USED BY (사용처):
// - Web APIs (string-based request/response)
//   웹 API (문자열 기반 요청/응답)
// - CLI tools (command-line arguments)
//   CLI 도구 (명령줄 인수)
// - Configuration (YAML/JSON dates)
//   구성 (YAML/JSON 날짜)
// - Scripts (quick date calculations)
//   스크립트 (빠른 날짜 계산)
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE FOR STRING-BASED DATA SOURCES
//    문자열 기반 데이터 소스에 사용
//    age, _ := timeutil.AgeString(userInput)
//
// 2. ALWAYS CHECK ERRORS
//    항상 오류 확인
//    if result, err := timeutil.AddDaysString(...); err != nil { /* handle */ }
//
// 3. USE BASE FUNCTIONS IN PERFORMANCE-CRITICAL CODE
//    성능 중요 코드에서 기본 함수 사용
//    // Parse once, use many times
//    t := timeutil.ParseAny(dateStr)
//    for i := 0; i < 1000; i++ { timeutil.AddDays(t, i) }
//
// 4. PREFER WHEN CONVENIENCE MATTERS
//    편의성이 중요할 때 선호
//    // Quick scripts, APIs, configuration
//
// 5. DOCUMENT EXPECTED DATE FORMAT
//    예상 날짜 형식 문서화
//    // Even though ParseAny is flexible, document expected format
//
// 6. USE FOR API ENDPOINT PARAMETERS
//    API 엔드포인트 매개변수에 사용
//    // Perfect for query parameters, form data
//
// 7. COMBINE WITH VALIDATION
//    검증과 결합
//    // Use IsBetweenString, IsAfterString for validation
//
// 8. LEVERAGE FLEXIBLE PARSING
//    유연한 파싱 활용
//    // Users can input various formats, all work
//
// ============================================================================

// String versions of time functions that accept string parameters instead of time.Time.
// These functions automatically parse the input strings using ParseAny.
// 문자열 매개변수를 받는 시간 함수의 문자열 버전입니다.
// 이 함수들은 ParseAny를 사용하여 입력 문자열을 자동으로 파싱합니다.

// SubTimeString calculates the time difference between two time strings.
// SubTimeString은 두 시간 문자열 사이의 시간 차이를 계산합니다.
//
// Example
// 예제:
//
//	diff, err := timeutil.SubTimeString("2024-10-04 08:34:42", "2024-10-14 14:56:23")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(diff.String()) // "10 days 6 hours 21 minutes"
func SubTimeString(s1, s2 string) (*TimeDiff, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return nil, fmt.Errorf("failed to parse first time string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return nil, fmt.Errorf("failed to parse second time string: %w", err)
	}

	return SubTime(t1, t2), nil
}

// DiffInDaysString returns the number of days between two time strings.
// DiffInDaysString은 두 시간 문자열 사이의 일수를 반환합니다.
func DiffInDaysString(s1, s2 string) (float64, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return 0, fmt.Errorf("failed to parse first time string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return 0, fmt.Errorf("failed to parse second time string: %w", err)
	}

	return DiffInDays(t1, t2), nil
}

// DiffInHoursString returns the number of hours between two time strings.
// DiffInHoursString은 두 시간 문자열 사이의 시간수를 반환합니다.
func DiffInHoursString(s1, s2 string) (float64, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return 0, fmt.Errorf("failed to parse first time string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return 0, fmt.Errorf("failed to parse second time string: %w", err)
	}

	return DiffInHours(t1, t2), nil
}

// DiffInMinutesString returns the number of minutes between two time strings.
// DiffInMinutesString은 두 시간 문자열 사이의 분수를 반환합니다.
func DiffInMinutesString(s1, s2 string) (float64, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return 0, fmt.Errorf("failed to parse first time string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return 0, fmt.Errorf("failed to parse second time string: %w", err)
	}

	return DiffInMinutes(t1, t2), nil
}

// AgeString calculates the age from a birth date string.
// AgeString은 생년월일 문자열로부터 나이를 계산합니다.
//
// Example
// 예제:
//
//	age, err := timeutil.AgeString("1990-01-15")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(age.Years) // 35 (assuming current year is 2025)
func AgeString(birthDate string) (*AgeDetail, error) {
	t, err := ParseAny(birthDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse birth date: %w", err)
	}

	return Age(t), nil
}

// AgeInYearsString calculates the age in years from a birth date string.
// AgeInYearsString은 생년월일 문자열로부터 나이(년)를 계산합니다.
func AgeInYearsString(birthDate string) (int, error) {
	t, err := ParseAny(birthDate)
	if err != nil {
		return 0, fmt.Errorf("failed to parse birth date: %w", err)
	}

	return AgeInYears(t), nil
}

// RelativeTimeString returns a human-readable relative time string.
// RelativeTimeString은 사람이 읽기 쉬운 상대 시간 문자열을 반환합니다.
//
// Example
// 예제:
//
//	rel, err := timeutil.RelativeTimeString("2024-10-13 15:30:00")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(rel) // "1 day ago"
func RelativeTimeString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return RelativeTime(t), nil
}

// IsBusinessDayString checks if a date string is a business day (Monday-Friday, not a holiday).
// IsBusinessDayString은 날짜 문자열이 영업일인지 확인합니다 (월-금, 공휴일 제외).
func IsBusinessDayString(s string) (bool, error) {
	t, err := ParseAny(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse date string: %w", err)
	}

	return IsBusinessDay(t), nil
}

// IsWeekendString checks if a date string falls on a weekend (Saturday or Sunday).
// IsWeekendString은 날짜 문자열이 주말인지 확인합니다 (토요일 또는 일요일).
func IsWeekendString(s string) (bool, error) {
	t, err := ParseAny(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse date string: %w", err)
	}

	return IsWeekend(t), nil
}

// AddDaysString adds a number of days to a date string.
// AddDaysString은 날짜 문자열에 일수를 더합니다.
//
// Example
// 예제:
//
//	result, err := timeutil.AddDaysString("2024-10-14", 7)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result) // 2024-10-21 00:00:00 +0900 KST
func AddDaysString(s string, days int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return AddDays(t, days), nil
}

// AddHoursString adds a number of hours to a datetime string.
// AddHoursString은 날짜시간 문자열에 시간수를 더합니다.
func AddHoursString(s string, hours int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime string: %w", err)
	}

	return AddHours(t, hours), nil
}

// AddMinutesString adds a number of minutes to a datetime string.
// AddMinutesString은 날짜시간 문자열에 분수를 더합니다.
func AddMinutesString(s string, minutes int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime string: %w", err)
	}

	return AddMinutes(t, minutes), nil
}

// SubDaysString subtracts a number of days from a date string.
// SubDaysString은 날짜 문자열에서 일수를 뺍니다.
func SubDaysString(s string, days int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return AddDays(t, -days), nil
}

// SubHoursString subtracts a number of hours from a datetime string.
// SubHoursString은 날짜시간 문자열에서 시간수를 뺍니다.
func SubHoursString(s string, hours int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime string: %w", err)
	}

	return AddHours(t, -hours), nil
}

// SubMinutesString subtracts a number of minutes from a datetime string.
// SubMinutesString은 날짜시간 문자열에서 분수를 뺍니다.
func SubMinutesString(s string, minutes int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime string: %w", err)
	}

	return AddMinutes(t, -minutes), nil
}

// FormatString parses a time string and formats it with the given layout.
// FormatString은 시간 문자열을 파싱하여 주어진 레이아웃으로 포맷합니다.
//
// Example
// 예제:
//
//	result, err := timeutil.FormatString("2024-10-14 15:30:00", "2006-01-02")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result) // "2024-10-14"
func FormatString(s, layout string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return t.Format(layout), nil
}

// FormatDateString parses a time string and formats it as a date (YYYY-MM-DD).
// FormatDateString은 시간 문자열을 파싱하여 날짜 형식으로 포맷합니다 (YYYY-MM-DD).
func FormatDateString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return FormatDate(t), nil
}

// FormatDateTimeString parses a time string and formats it as datetime (YYYY-MM-DD HH:mm:ss).
// FormatDateTimeString은 시간 문자열을 파싱하여 날짜시간 형식으로 포맷합니다 (YYYY-MM-DD HH:mm:ss).
func FormatDateTimeString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return FormatDateTime(t), nil
}

// FormatISO8601String parses a time string and formats it in ISO8601 format.
// FormatISO8601String은 시간 문자열을 파싱하여 ISO8601 형식으로 포맷합니다.
func FormatISO8601String(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return FormatISO8601(t), nil
}

// ConvertTimezoneString parses a time string and converts it to a different timezone.
// ConvertTimezoneString은 시간 문자열을 파싱하여 다른 타임존으로 변환합니다.
//
// Example
// 예제:
//
//	result, err := timeutil.ConvertTimezoneString("2024-10-14 15:30:00", "America/New_York")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result)
func ConvertTimezoneString(s, tz string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time string: %w", err)
	}

	return ConvertTimezone(t, tz)
}

// StartOfDayString returns the start of the day (00:00:00) for a date string.
// StartOfDayString은 날짜 문자열의 시작 시간(00:00:00)을 반환합니다.
func StartOfDayString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return StartOfDay(t), nil
}

// EndOfDayString returns the end of the day (23:59:59) for a date string.
// EndOfDayString은 날짜 문자열의 종료 시간(23:59:59)을 반환합니다.
func EndOfDayString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return EndOfDay(t), nil
}

// StartOfWeekString returns the start of the week (Monday 00:00:00) for a date string.
// StartOfWeekString은 날짜 문자열의 주 시작 시간(월요일 00:00:00)을 반환합니다.
func StartOfWeekString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return StartOfWeek(t), nil
}

// EndOfWeekString returns the end of the week (Sunday 23:59:59) for a date string.
// EndOfWeekString은 날짜 문자열의 주 종료 시간(일요일 23:59:59)을 반환합니다.
func EndOfWeekString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return EndOfWeek(t), nil
}

// StartOfMonthString returns the start of the month (day 1, 00:00:00) for a date string.
// StartOfMonthString은 날짜 문자열의 월 시작 시간(1일 00:00:00)을 반환합니다.
func StartOfMonthString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return StartOfMonth(t), nil
}

// EndOfMonthString returns the end of the month (last day, 23:59:59) for a date string.
// EndOfMonthString은 날짜 문자열의 월 종료 시간(마지막 날 23:59:59)을 반환합니다.
func EndOfMonthString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return EndOfMonth(t), nil
}

// StartOfYearString returns the start of the year (Jan 1, 00:00:00) for a date string.
// StartOfYearString은 날짜 문자열의 연 시작 시간(1월 1일 00:00:00)을 반환합니다.
func StartOfYearString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return StartOfYear(t), nil
}

// EndOfYearString returns the end of the year (Dec 31, 23:59:59) for a date string.
// EndOfYearString은 날짜 문자열의 연 종료 시간(12월 31일 23:59:59)을 반환합니다.
func EndOfYearString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return EndOfYear(t), nil
}

// WeekdayString returns the weekday name for a date string.
// WeekdayString은 날짜 문자열의 요일 이름을 반환합니다.
func WeekdayString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return t.Weekday().String(), nil
}

// WeekdayKoreanString returns the Korean weekday name for a date string.
// WeekdayKoreanString은 날짜 문자열의 한글 요일 이름을 반환합니다.
func WeekdayKoreanString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return WeekdayKorean(t), nil
}

// WeekdayShortString returns the short weekday name for a date string.
// WeekdayShortString은 날짜 문자열의 짧은 요일 이름을 반환합니다.
func WeekdayShortString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return t.Format("Mon"), nil
}

// WeekdayShortKoreanString returns the short Korean weekday name for a date string.
// WeekdayShortKoreanString은 날짜 문자열의 짧은 한글 요일 이름을 반환합니다.
func WeekdayShortKoreanString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return WeekdayKoreanShort(t), nil
}

// WeekdayNumberString returns the weekday number (0=Sunday, 6=Saturday) for a date string.
// WeekdayNumberString은 날짜 문자열의 요일 번호(0=일요일, 6=토요일)를 반환합니다.
func WeekdayNumberString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return int(t.Weekday()), nil
}

// WeekOfYearString returns the ISO week number for a date string.
// WeekOfYearString은 날짜 문자열의 ISO 주 번호를 반환합니다.
func WeekOfYearString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return WeekOfYear(t), nil
}

// WeekOfMonthString returns the week number within the month for a date string.
// WeekOfMonthString은 날짜 문자열의 월 내 주 번호를 반환합니다.
func WeekOfMonthString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return WeekOfMonth(t), nil
}

// DaysInMonthString returns the number of days in the month for a date string.
// DaysInMonthString은 날짜 문자열의 월의 일수를 반환합니다.
func DaysInMonthString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return DaysInMonth(t), nil
}

// DaysInYearString returns the number of days in the year for a date string.
// DaysInYearString은 날짜 문자열의 연도의 일수를 반환합니다.
func DaysInYearString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return DaysInYear(t), nil
}

// MonthKoreanString returns the Korean month name for a date string.
// MonthKoreanString은 날짜 문자열의 한글 월 이름을 반환합니다.
func MonthKoreanString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return MonthKorean(t), nil
}

// MonthNameString returns the full month name for a date string.
// MonthNameString은 날짜 문자열의 전체 월 이름을 반환합니다.
func MonthNameString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return MonthName(t), nil
}

// MonthNameShortString returns the short month name for a date string.
// MonthNameShortString은 날짜 문자열의 짧은 월 이름을 반환합니다.
func MonthNameShortString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return MonthNameShort(t), nil
}

// QuarterString returns the quarter (1-4) for a date string.
// QuarterString은 날짜 문자열의 분기(1-4)를 반환합니다.
func QuarterString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return Quarter(t), nil
}

// IsLeapYearString checks if the year in a date string is a leap year.
// IsLeapYearString은 날짜 문자열의 연도가 윤년인지 확인합니다.
func IsLeapYearString(s string) (bool, error) {
	t, err := ParseAny(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse date string: %w", err)
	}

	return IsLeapYear(t), nil
}

// IsSameDayString checks if two date strings are on the same day.
// IsSameDayString은 두 날짜 문자열이 같은 날인지 확인합니다.
func IsSameDayString(s1, s2 string) (bool, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return false, fmt.Errorf("failed to parse first date string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return false, fmt.Errorf("failed to parse second date string: %w", err)
	}

	return IsSameDay(t1, t2), nil
}

// IsBeforeString checks if the first date string is before the second.
// IsBeforeString은 첫 번째 날짜 문자열이 두 번째보다 이전인지 확인합니다.
func IsBeforeString(s1, s2 string) (bool, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return false, fmt.Errorf("failed to parse first date string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return false, fmt.Errorf("failed to parse second date string: %w", err)
	}

	return IsBefore(t1, t2), nil
}

// IsAfterString checks if the first date string is after the second.
// IsAfterString은 첫 번째 날짜 문자열이 두 번째보다 이후인지 확인합니다.
func IsAfterString(s1, s2 string) (bool, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return false, fmt.Errorf("failed to parse first date string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return false, fmt.Errorf("failed to parse second date string: %w", err)
	}

	return IsAfter(t1, t2), nil
}

// IsBetweenString checks if a date string is between two other dates.
// IsBetweenString은 날짜 문자열이 두 날짜 사이에 있는지 확인합니다.
func IsBetweenString(s, start, end string) (bool, error) {
	t, err := ParseAny(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse date string: %w", err)
	}

	startTime, err := ParseAny(start)
	if err != nil {
		return false, fmt.Errorf("failed to parse start date string: %w", err)
	}

	endTime, err := ParseAny(end)
	if err != nil {
		return false, fmt.Errorf("failed to parse end date string: %w", err)
	}

	return IsBetween(t, startTime, endTime), nil
}
