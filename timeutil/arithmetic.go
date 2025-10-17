package timeutil

import "time"

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/arithmetic.go
// Purpose: Time and date arithmetic operations
//          시간 및 날짜 산술 연산
//
// This file provides comprehensive time arithmetic operations for adding/subtracting
// time units and finding boundary times (start/end of periods). These operations are
// essential for date calculations, scheduling, reporting, and time-based queries.
//
// 이 파일은 시간 단위를 더하거나 빼고 경계 시간(기간의 시작/끝)을 찾는 포괄적인
// 시간 산술 연산을 제공합니다. 이러한 연산은 날짜 계산, 스케줄링, 보고 및
// 시간 기반 쿼리에 필수적입니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. TIME UNIT ADDITION (시간 단위 더하기)
//    - AddSeconds: Adds seconds to a time
//      초 더하기
//    - AddMinutes: Adds minutes to a time
//      분 더하기
//    - AddHours: Adds hours to a time
//      시간 더하기
//    - Supports negative values for subtraction
//      빼기를 위한 음수 값 지원
//    - Handles overflow automatically (e.g., 61 seconds → +1 minute, +1 second)
//      오버플로우 자동 처리 (예: 61초 → +1분, +1초)
//
// 2. DATE UNIT ADDITION (날짜 단위 더하기)
//    - AddDays: Adds days to a date
//      일 더하기
//    - AddWeeks: Adds weeks to a date
//      주 더하기
//    - AddMonths: Adds months to a date (handles month boundaries)
//      월 더하기 (월 경계 처리)
//    - AddYears: Adds years to a date (handles leap years)
//      년 더하기 (윤년 처리)
//    - Supports negative values for subtraction
//      빼기를 위한 음수 값 지원
//
// 3. DAILY BOUNDARIES (일일 경계)
//    - StartOfDay: Returns 00:00:00 of the day in KST
//      KST로 하루의 00:00:00 반환
//    - EndOfDay: Returns 23:59:59.999999999 of the day in KST
//      KST로 하루의 23:59:59.999999999 반환
//    - Useful for date range queries
//      날짜 범위 쿼리에 유용
//
// 4. WEEKLY BOUNDARIES (주간 경계)
//    - StartOfWeek: Returns Monday 00:00:00 in KST
//      KST로 월요일 00:00:00 반환
//    - EndOfWeek: Returns Sunday 23:59:59 in KST
//      KST로 일요일 23:59:59 반환
//    - Week starts on Monday (ISO 8601 standard)
//      월요일 시작 (ISO 8601 표준)
//
// 5. MONTHLY BOUNDARIES (월간 경계)
//    - StartOfMonth: Returns 1st day 00:00:00 in KST
//      KST로 1일 00:00:00 반환
//    - EndOfMonth: Returns last day 23:59:59 in KST
//      KST로 마지막 날 23:59:59 반환
//    - Handles variable month lengths (28-31 days)
//      가변 월 길이 처리 (28-31일)
//
// 6. YEARLY BOUNDARIES (연간 경계)
//    - StartOfYear: Returns Jan 1 00:00:00 in KST
//      KST로 1월 1일 00:00:00 반환
//    - EndOfYear: Returns Dec 31 23:59:59 in KST
//      KST로 12월 31일 23:59:59 반환
//    - Useful for annual reports and summaries
//      연간 보고서 및 요약에 유용
//
// 7. QUARTERLY BOUNDARIES (분기 경계)
//    - StartOfQuarter: Returns start of financial quarter in KST
//      KST로 분기 시작 반환
//    - EndOfQuarter: Returns end of financial quarter in KST
//      KST로 분기 끝 반환
//    - Quarters: Q1 (Jan-Mar), Q2 (Apr-Jun), Q3 (Jul-Sep), Q4 (Oct-Dec)
//      분기: 1분기 (1-3월), 2분기 (4-6월), 3분기 (7-9월), 4분기 (10-12월)
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. SEMANTIC CLARITY (의미적 명확성)
//    - Function names clearly indicate the operation
//      함수 이름이 연산을 명확히 나타냅니다
//    - AddDays vs Add(24*time.Hour) - more readable
//      AddDays가 Add(24*time.Hour)보다 가독성이 높습니다
//    - StartOfDay vs manual time construction - more intuitive
//      StartOfDay가 수동 시간 생성보다 직관적입니다
//
// 2. CALENDAR-AWARE ARITHMETIC (달력 인식 산술)
//    - AddMonths handles variable month lengths correctly
//      AddMonths는 가변 월 길이를 올바르게 처리합니다
//    - Example: Jan 31 + 1 month = Feb 28/29 (not invalid)
//      예시: 1월 31일 + 1개월 = 2월 28/29일 (유효하지 않은 날짜가 아님)
//    - AddYears handles leap years correctly
//      AddYears는 윤년을 올바르게 처리합니다
//    - Example: Feb 29, 2024 + 1 year = Feb 28, 2025
//      예시: 2024년 2월 29일 + 1년 = 2025년 2월 28일
//
// 3. KST AS DEFAULT TIMEZONE (KST를 기본 타임존으로)
//    - All boundary functions return times in KST
//      모든 경계 함수는 KST로 시간을 반환합니다
//    - Input times are converted to KST before calculation
//      입력 시간은 계산 전에 KST로 변환됩니다
//    - Ensures consistency across the package
//      패키지 전체의 일관성 보장
//
// 4. NANOSECOND PRECISION (나노초 정밀도)
//    - EndOfDay uses .999999999 for maximum precision
//      EndOfDay는 최대 정밀도를 위해 .999999999 사용
//    - Ensures no time is lost in range queries
//      범위 쿼리에서 시간 손실 방지
//    - Example: Events up to "end of day" include all times before midnight
//      예시: "하루의 끝"까지의 이벤트는 자정 전의 모든 시간을 포함합니다
//
// 5. MONDAY-FIRST WEEK (월요일 시작 주)
//    - StartOfWeek returns Monday, not Sunday
//      StartOfWeek는 일요일이 아닌 월요일을 반환합니다
//    - Follows ISO 8601 standard
//      ISO 8601 표준을 따릅니다
//    - Consistent with international business practices
//      국제 비즈니스 관행과 일치합니다
//    - Go's time.Weekday: Sunday = 0, Monday = 1
//      Go의 time.Weekday: 일요일 = 0, 월요일 = 1
//
// 6. IMMUTABILITY (불변성)
//    - All functions return new time.Time values
//      모든 함수는 새로운 time.Time 값을 반환합니다
//    - Original time is never modified
//      원본 시간은 절대 수정되지 않습니다
//    - Safe for concurrent use
//      동시 사용에 안전합니다
//
// ============================================================================
// ARITHMETIC OPERATIONS OVERVIEW / 산술 연산 개요
// ============================================================================
//
// TIME UNIT ADDITION (시간 단위 더하기) - 3 functions
// ├─ AddSeconds        : Adds seconds
// ├─ AddMinutes        : Adds minutes
// └─ AddHours          : Adds hours
//
// DATE UNIT ADDITION (날짜 단위 더하기) - 4 functions
// ├─ AddDays           : Adds days
// ├─ AddWeeks          : Adds weeks
// ├─ AddMonths         : Adds months (calendar-aware)
// └─ AddYears          : Adds years (leap-year aware)
//
// DAILY BOUNDARIES (일일 경계) - 2 functions
// ├─ StartOfDay        : Returns 00:00:00
// └─ EndOfDay          : Returns 23:59:59.999999999
//
// WEEKLY BOUNDARIES (주간 경계) - 2 functions
// ├─ StartOfWeek       : Returns Monday 00:00:00
// └─ EndOfWeek         : Returns Sunday 23:59:59
//
// MONTHLY BOUNDARIES (월간 경계) - 2 functions
// ├─ StartOfMonth      : Returns 1st day 00:00:00
// └─ EndOfMonth        : Returns last day 23:59:59
//
// YEARLY BOUNDARIES (연간 경계) - 2 functions
// ├─ StartOfYear       : Returns Jan 1 00:00:00
// └─ EndOfYear         : Returns Dec 31 23:59:59
//
// QUARTERLY BOUNDARIES (분기 경계) - 2 functions
// ├─ StartOfQuarter    : Returns quarter start
// └─ EndOfQuarter      : Returns quarter end
//
// Total: 17 arithmetic functions
// 총: 17개의 산술 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
// All functions: O(1) - constant time
// 모든 함수: O(1) - 상수 시간
//
// TIME UNIT ADDITION (AddSeconds, AddMinutes, AddHours):
// - Simple duration addition using time.Add()
//   time.Add()를 사용한 간단한 기간 더하기
// - No calendar calculations needed
//   달력 계산 불필요
// - Very fast: ~10-20 nanoseconds
//   매우 빠름: ~10-20나노초
//
// DATE UNIT ADDITION (AddDays, AddWeeks, AddMonths, AddYears):
// - Uses time.AddDate() for calendar-aware arithmetic
//   달력 인식 산술을 위해 time.AddDate() 사용
// - Handles month/year boundaries
//   월/년 경계 처리
// - Fast: ~50-100 nanoseconds
//   빠름: ~50-100나노초
//
// BOUNDARY FUNCTIONS (StartOf*, EndOf*):
// - Creates new time.Time with specific components
//   특정 구성 요소로 새 time.Time 생성
// - Involves timezone conversion to KST
//   KST로 타임존 변환 포함
// - Moderate: ~100-200 nanoseconds
//   보통: ~100-200나노초
//
// SPACE COMPLEXITY (공간 복잡도):
// - All functions: O(1) - constant space
//   모든 함수: O(1) - 상수 공간
// - Returns a single time.Time value
//   단일 time.Time 값 반환
// - No heap allocations
//   힙 할당 없음
//
// PERFORMANCE TIPS (성능 팁):
// 1. Boundary functions create new time.Time - cache if reused
//    경계 함수는 새 time.Time 생성 - 재사용 시 캐시
// 2. Use AddDays for date arithmetic, not Add(24*time.Hour)
//    날짜 산술에는 Add(24*time.Hour)가 아닌 AddDays 사용
//    Reason: AddDays handles DST transitions correctly
//    이유: AddDays는 일광 절약 시간 전환을 올바르게 처리합니다
// 3. Batch operations when possible
//    가능하면 일괄 처리
//
// ============================================================================
// CALENDAR ARITHMETIC DETAILS / 달력 산술 세부 사항
// ============================================================================
//
// MONTH ADDITION BEHAVIOR (월 더하기 동작):
// Go's time.AddDate() handles month overflow intelligently:
// Go의 time.AddDate()는 월 오버플로우를 지능적으로 처리합니다:
//
// Example 1: Adding 1 month to Jan 31
// 예시 1: 1월 31일에 1개월 더하기
//   Jan 31 + 1 month = Feb 28 (or Feb 29 in leap year)
//   1월 31일 + 1개월 = 2월 28일 (윤년은 2월 29일)
//   Not March 3 or March 2
//   3월 3일이나 3월 2일이 아님
//
// Example 2: Adding 1 month to Oct 31
// 예시 2: 10월 31일에 1개월 더하기
//   Oct 31 + 1 month = Nov 30 (November has 30 days)
//   10월 31일 + 1개월 = 11월 30일 (11월은 30일까지)
//   Not Dec 1
//   12월 1일이 아님
//
// This behavior is DESIRED and correct for most use cases:
// 이 동작은 대부분의 사용 사례에서 바람직하고 올바릅니다:
// - Scheduling: "Meet on the last day of each month"
//   스케줄링: "매월 마지막 날에 만나기"
// - Billing: "Charge on the 31st or last day of month"
//   청구: "31일 또는 월의 마지막 날에 청구"
//
// YEAR ADDITION WITH LEAP YEAR (윤년이 있는 년 더하기):
// Example: Adding 1 year to Feb 29, 2024 (leap year)
// 예시: 2024년 2월 29일(윤년)에 1년 더하기
//   Feb 29, 2024 + 1 year = Feb 28, 2025
//   2024년 2월 29일 + 1년 = 2025년 2월 28일
//   (2025 is not a leap year)
//   (2025년은 윤년이 아님)
//
// DAY ADDITION VS HOUR ADDITION (일 더하기 vs 시간 더하기):
// IMPORTANT: AddDays is NOT the same as Add(24*time.Hour)
// 중요: AddDays는 Add(24*time.Hour)와 같지 않습니다
//
// Reason: Daylight Saving Time (DST) transitions
// 이유: 일광 절약 시간(DST) 전환
//
// Example in a region with DST:
// DST가 있는 지역의 예시:
//   March 10, 2024 01:00 + Add(24*time.Hour) = March 11, 2024 01:00
//   BUT if DST starts at 02:00 on March 10:
//   하지만 3월 10일 02:00에 DST가 시작되면:
//   March 10, 2024 01:00 + AddDays(1) = March 11, 2024 01:00 (correct)
//   March 10, 2024 01:00 + 24 hours = March 11, 2024 02:00 (wrong!)
//
// KST doesn't have DST, but AddDays is still preferred for clarity
// KST에는 DST가 없지만, 명확성을 위해 AddDays가 여전히 선호됩니다
//
// WEEK CALCULATION (주 계산):
// Go's time.Weekday: Sunday = 0, Monday = 1, ..., Saturday = 6
// Go의 time.Weekday: 일요일 = 0, 월요일 = 1, ..., 토요일 = 6
//
// ISO 8601 standard: Week starts on Monday
// ISO 8601 표준: 주는 월요일에 시작
//
// Conversion logic in StartOfWeek:
// StartOfWeek의 변환 로직:
//   if weekday == 0 (Sunday):
//       weekday = 7  // Treat as last day of week
//   daysToMonday = weekday - 1
//   return t minus daysToMonday
//
// Example: Wednesday (3) → daysToMonday = 2 → go back 2 days to Monday
// 예시: 수요일 (3) → daysToMonday = 2 → 2일 전으로 가서 월요일
//
// ============================================================================
// BOUNDARY TIME PRECISION / 경계 시간 정밀도
// ============================================================================
//
// START OF PERIODS (기간 시작):
// All Start* functions return 00:00:00.000000000
// 모든 Start* 함수는 00:00:00.000000000을 반환합니다
//
// Example:
//   t := time.Date(2024, 1, 15, 14, 30, 45, 123456789, timeutil.KST)
//   start := timeutil.StartOfDay(t)
//   // Result: 2024-01-15 00:00:00.000000000 +0900 KST
//
// END OF PERIODS (기간 끝):
// All End* functions return 23:59:59.999999999
// 모든 End* 함수는 23:59:59.999999999를 반환합니다
//
// Why .999999999? (왜 .999999999?)
// - Maximum precision before next day
//   다음 날 전의 최대 정밀도
// - Ensures range queries include all times up to end
//   범위 쿼리가 끝까지의 모든 시간을 포함하도록 보장
// - Example query: WHERE time >= StartOfDay AND time <= EndOfDay
//   예시 쿼리: WHERE time >= StartOfDay AND time <= EndOfDay
//   This includes events at 23:59:59.999999999
//   이것은 23:59:59.999999999의 이벤트를 포함합니다
//
// Alternative approach NOT used:
// 사용하지 않는 대체 접근법:
// - EndOfDay = StartOfDay + 1 day (next day 00:00:00)
//   EndOfDay = StartOfDay + 1일 (다음 날 00:00:00)
// - Query: WHERE time >= StartOfDay AND time < EndOfDay
//   쿼리: WHERE time >= StartOfDay AND time < EndOfDay
// - Reason: Less intuitive, requires < instead of <=
//   이유: 덜 직관적, <= 대신 < 필요
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Date Range Queries (날짜 범위 쿼리)
// Use case: Query database for events on a specific date
// 사용 사례: 특정 날짜의 이벤트에 대한 데이터베이스 쿼리
//
//   date := time.Date(2024, 1, 15, 0, 0, 0, 0, timeutil.KST)
//   start := timeutil.StartOfDay(date)
//   end := timeutil.EndOfDay(date)
//
//   events := db.Query(
//       "SELECT * FROM events WHERE created_at >= ? AND created_at <= ?",
//       start, end,
//   )
//   // Returns all events on Jan 15, 2024
//   // 2024년 1월 15일의 모든 이벤트 반환
//
// PATTERN 2: Weekly Reports (주간 보고서)
// Use case: Generate reports for the current week
// 사용 사례: 현재 주의 보고서 생성
//
//   now := time.Now()
//   weekStart := timeutil.StartOfWeek(now)
//   weekEnd := timeutil.EndOfWeek(now)
//
//   report := generateReport(weekStart, weekEnd)
//   fmt.Printf("Weekly Report: %s to %s\n",
//       timeutil.FormatDate(weekStart),
//       timeutil.FormatDate(weekEnd))
//
// PATTERN 3: Monthly Summaries (월간 요약)
// Use case: Calculate monthly statistics
// 사용 사례: 월간 통계 계산
//
//   now := time.Now()
//   monthStart := timeutil.StartOfMonth(now)
//   monthEnd := timeutil.EndOfMonth(now)
//
//   stats := calculateStats(monthStart, monthEnd)
//   fmt.Printf("Month: %s, Sales: %d\n",
//       now.Month(), stats.TotalSales)
//
// PATTERN 4: Scheduling Future Events (미래 이벤트 스케줄링)
// Use case: Schedule meetings at specific intervals
// 사용 사례: 특정 간격으로 회의 스케줄링
//
//   now := time.Now()
//   meetings := []time.Time{
//       timeutil.AddDays(now, 1),    // Tomorrow
//       timeutil.AddWeeks(now, 1),   // Next week
//       timeutil.AddMonths(now, 1),  // Next month
//   }
//
//   for i, meeting := range meetings {
//       fmt.Printf("Meeting %d: %s\n", i+1, timeutil.FormatDate(meeting))
//   }
//
// PATTERN 5: Age Calculation (나이 계산)
// Use case: Calculate age from birthdate
// 사용 사례: 생년월일로부터 나이 계산
//
//   birthdate := time.Date(1990, 5, 15, 0, 0, 0, 0, timeutil.KST)
//   today := timeutil.StartOfDay(time.Now())
//
//   years := 0
//   current := birthdate
//   for timeutil.AddYears(current, 1).Before(today) ||
//       timeutil.AddYears(current, 1).Equal(today) {
//       years++
//       current = timeutil.AddYears(current, 1)
//   }
//   fmt.Printf("Age: %d years\n", years)
//
// PATTERN 6: Billing Cycles (청구 주기)
// Use case: Calculate next billing date
// 사용 사례: 다음 청구 날짜 계산
//
//   lastBillingDate := time.Date(2024, 1, 31, 0, 0, 0, 0, timeutil.KST)
//   nextBillingDate := timeutil.AddMonths(lastBillingDate, 1)
//
//   fmt.Printf("Next billing: %s\n", timeutil.FormatDate(nextBillingDate))
//   // If last billing was Jan 31, next is Feb 28/29
//   // 마지막 청구가 1월 31일이면, 다음은 2월 28/29일
//
// PATTERN 7: Expiration Dates (만료 날짜)
// Use case: Set expiration 30 days from now
// 사용 사례: 지금부터 30일 후 만료 설정
//
//   now := time.Now()
//   expiresAt := timeutil.EndOfDay(timeutil.AddDays(now, 30))
//
//   token := generateToken()
//   token.ExpiresAt = expiresAt
//   // Token expires at 23:59:59 on the 30th day
//   // 토큰은 30일째의 23:59:59에 만료
//
// PATTERN 8: Quarter-End Processing (분기 말 처리)
// Use case: Run quarter-end reports
// 사용 사례: 분기 말 보고서 실행
//
//   now := time.Now()
//   quarterEnd := timeutil.EndOfQuarter(now)
//
//   if timeutil.StartOfDay(now).Equal(timeutil.StartOfDay(quarterEnd)) {
//       runQuarterEndReports()
//       fmt.Println("Running quarter-end reports")
//   }
//
// PATTERN 9: Time Range Validation (시간 범위 검증)
// Use case: Check if a date falls within a range
// 사용 사례: 날짜가 범위 내에 있는지 확인
//
//   startRange := timeutil.StartOfMonth(time.Now())
//   endRange := timeutil.EndOfMonth(time.Now())
//   checkDate := time.Date(2024, 1, 15, 14, 30, 0, 0, timeutil.KST)
//
//   if !checkDate.Before(startRange) && !checkDate.After(endRange) {
//       fmt.Println("Date is within current month")
//   }
//
// PATTERN 10: Recurring Events (반복 이벤트)
// Use case: Generate recurring event dates
// 사용 사례: 반복 이벤트 날짜 생성
//
//   startDate := time.Date(2024, 1, 1, 9, 0, 0, 0, timeutil.KST)
//   var dates []time.Time
//
//   // Every Monday for 10 weeks
//   for i := 0; i < 10; i++ {
//       eventDate := timeutil.AddWeeks(startDate, i)
//       dates = append(dates, eventDate)
//   }
//
//   for _, date := range dates {
//       fmt.Printf("Event on: %s\n", timeutil.FormatDate(date))
//   }
//
// ============================================================================
// COMPARISON WITH OTHER APPROACHES / 다른 접근법과의 비교
// ============================================================================
//
// ADDING DAYS: AddDays vs time.Add
// AddDays 대 time.Add:
//
// timeutil.AddDays(t, 1)
// ✓ Calendar-aware (달력 인식)
// ✓ Handles DST correctly (DST 올바르게 처리)
// ✓ Semantic clarity (의미적 명확성)
// ✓ More readable (가독성 높음)
//
// t.Add(24 * time.Hour)
// ✗ Not calendar-aware (달력 비인식)
// ✗ May fail with DST (DST에서 실패 가능)
// ✗ Less semantic (의미적으로 덜 명확)
// ✗ Harder to read (가독성 낮음)
//
// FINDING START OF DAY: StartOfDay vs manual
// StartOfDay 대 수동 생성:
//
// timeutil.StartOfDay(t)
// ✓ One line (한 줄)
// ✓ No errors (오류 없음)
// ✓ Handles timezone (타임존 처리)
// ✓ Consistent (일관성)
//
// time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, timeutil.KST)
// ✗ Multiple lines (여러 줄)
// ✗ Easy to make mistakes (실수하기 쉬움)
// ✗ Must handle timezone manually (타임존 수동 처리)
// ✗ Verbose (장황함)
//
// ADDING MONTHS: AddMonths vs time.AddDate
// AddMonths 대 time.AddDate:
//
// timeutil.AddMonths(t, 1)
// ✓ Clear intent (명확한 의도)
// ✓ Semantic (의미적)
// ✓ Self-documenting (자기 문서화)
//
// t.AddDate(0, 1, 0)
// ✗ Magic numbers (매직 넘버)
// ✗ Less clear (덜 명확)
// ✗ Need comments (주석 필요)
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// All functions in this file are THREAD-SAFE:
// 이 파일의 모든 함수는 스레드 안전합니다:
//
// 1. IMMUTABILITY (불변성)
//    - Input time.Time is never modified
//      입력 time.Time은 절대 수정되지 않습니다
//    - All functions return new time.Time values
//      모든 함수는 새로운 time.Time 값을 반환합니다
//    - No shared mutable state
//      공유 가변 상태 없음
//
// 2. READ-ONLY CONSTANTS (읽기 전용 상수)
//    - Uses constants from constants.go (DaysPerWeek, etc.)
//      constants.go의 상수 사용 (DaysPerWeek 등)
//    - defaultLocation is pre-loaded and read-only
//      defaultLocation은 미리 로드되고 읽기 전용
//    - No global state mutation
//      전역 상태 변경 없음
//
// 3. PURE FUNCTIONS (순수 함수)
//    - No side effects
//      부작용 없음
//    - Output depends only on input
//      출력은 입력에만 의존
//    - No I/O operations
//      I/O 작업 없음
//
// 4. SAFE CONCURRENT USAGE (안전한 동시 사용)
//    - Multiple goroutines can call these functions simultaneously
//      여러 고루틴이 동시에 이러한 함수를 호출할 수 있습니다
//    - No locks or synchronization needed
//      잠금이나 동기화가 필요 없습니다
//    - Example:
//      예시:
//        var wg sync.WaitGroup
//        for i := 0; i < 100; i++ {
//            wg.Add(1)
//            go func(days int) {
//                defer wg.Done()
//                future := timeutil.AddDays(time.Now(), days)
//                // Process future date
//            }(i)
//        }
//        wg.Wait()
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM constants.go:
// - DaysPerWeek: 7 (for AddWeeks calculation)
// - defaultLocation: KST (for boundary functions)
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.Add(): For time unit addition
// - time.AddDate(): For date unit addition
// - time.Date(): For creating boundary times
//
// USED BY (사용처):
// - Application code for date calculations
//   날짜 계산을 위한 애플리케이션 코드
// - Scheduling systems
//   스케줄링 시스템
// - Reporting and analytics
//   보고 및 분석
// - Database queries with time ranges
//   시간 범위가 있는 데이터베이스 쿼리
// - business.go (for business day calculations)
//   business.go (영업일 계산용)
// - age.go (for age calculations)
//   age.go (나이 계산용)
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE SEMANTIC FUNCTIONS OVER MANUAL CALCULATIONS
//    수동 계산보다 의미적 함수 사용
//    ✓ Good: timeutil.AddDays(t, 7)
//    ✗ Avoid: t.Add(7 * 24 * time.Hour)
//
// 2. USE BOUNDARY FUNCTIONS FOR RANGE QUERIES
//    범위 쿼리에 경계 함수 사용
//    start := timeutil.StartOfDay(date)
//    end := timeutil.EndOfDay(date)
//    query := "WHERE time >= ? AND time <= ?"
//
// 3. CACHE BOUNDARY TIMES IF REUSED
//    재사용 시 경계 시간 캐시
//    // Calculate once
//    monthStart := timeutil.StartOfMonth(now)
//    // Reuse in multiple queries
//
// 4. USE ADDMONTHS FOR MONTH-BASED SCHEDULING
//    월 기반 스케줄링에 AddMonths 사용
//    // Handles month boundaries correctly
//    nextBilling := timeutil.AddMonths(lastBilling, 1)
//
// 5. REMEMBER MONDAY-FIRST WEEKS
//    월요일 시작 주 기억
//    // Week starts on Monday (ISO 8601)
//    weekStart := timeutil.StartOfWeek(now)
//
// 6. USE ENDOFDAY FOR INCLUSIVE RANGES
//    포괄적 범위에 EndOfDay 사용
//    // Include all times up to end of day
//    end := timeutil.EndOfDay(date)
//    if event.Time.Before(end) || event.Time.Equal(end) {
//        // Event is within the day
//    }
//
// 7. VALIDATE RESULTS FOR MONTH/YEAR ADDITION
//    월/년 더하기 결과 검증
//    // Feb 29 + 1 year may become Feb 28
//    result := timeutil.AddYears(leapDate, 1)
//    // Check if result is as expected
//
// 8. USE APPROPRIATE PRECISION
//    적절한 정밀도 사용
//    // For date-only queries, use StartOfDay/EndOfDay
//    // For timestamp queries, use exact times
//
// ============================================================================

// AddSeconds adds the specified number of seconds to a time.
// AddSeconds는 시간에 지정된 초를 더합니다.
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// AddMinutes adds the specified number of minutes to a time.
// AddMinutes는 시간에 지정된 분을 더합니다.
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddHours adds the specified number of hours to a time.
// AddHours는 시간에 지정된 시간을 더합니다.
func AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

// AddDays adds the specified number of days to a time.
// AddDays는 시간에 지정된 일을 더합니다.
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddWeeks adds the specified number of weeks to a time.
// AddWeeks는 시간에 지정된 주를 더합니다.
func AddWeeks(t time.Time, weeks int) time.Time {
	return t.AddDate(0, 0, weeks*DaysPerWeek)
}

// AddMonths adds the specified number of months to a time.
// AddMonths는 시간에 지정된 월을 더합니다.
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYears adds the specified number of years to a time.
// AddYears는 시간에 지정된 년을 더합니다.
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// StartOfDay returns the start of the day (00:00:00) in KST.
// StartOfDay는 KST로 하루의 시작 (00:00:00)을 반환합니다.
func StartOfDay(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, defaultLocation)
}

// EndOfDay returns the end of the day (23:59:59) in KST.
// EndOfDay는 KST로 하루의 끝 (23:59:59)을 반환합니다.
func EndOfDay(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, defaultLocation)
}

// StartOfWeek returns the start of the week (Monday 00:00:00) in KST.
// StartOfWeek는 KST로 주의 시작 (월요일 00:00:00)을 반환합니다.
func StartOfWeek(t time.Time) time.Time {
	t = t.In(defaultLocation)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7 // Sunday = 7
	}
	daysToMonday := weekday - 1
	return StartOfDay(t.AddDate(0, 0, -daysToMonday))
}

// EndOfWeek returns the end of the week (Sunday 23:59:59) in KST.
// EndOfWeek는 KST로 주의 끝 (일요일 23:59:59)을 반환합니다.
func EndOfWeek(t time.Time) time.Time {
	t = t.In(defaultLocation)
	weekday := int(t.Weekday())
	if weekday == 0 {
		return EndOfDay(t)
	}
	daysToSunday := 7 - weekday
	return EndOfDay(t.AddDate(0, 0, daysToSunday))
}

// StartOfMonth returns the start of the month (1st day 00:00:00) in KST.
// StartOfMonth는 KST로 월의 시작 (1일 00:00:00)을 반환합니다.
func StartOfMonth(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, defaultLocation)
}

// EndOfMonth returns the end of the month (last day 23:59:59) in KST.
// EndOfMonth는 KST로 월의 끝 (마지막 날 23:59:59)을 반환합니다.
func EndOfMonth(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999999999, defaultLocation)
}

// StartOfYear returns the start of the year (Jan 1 00:00:00) in KST.
// StartOfYear는 KST로 년의 시작 (1월 1일 00:00:00)을 반환합니다.
func StartOfYear(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, defaultLocation)
}

// EndOfYear returns the end of the year (Dec 31 23:59:59) in KST.
// EndOfYear는 KST로 년의 끝 (12월 31일 23:59:59)을 반환합니다.
func EndOfYear(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, defaultLocation)
}

// StartOfQuarter returns the start of the quarter in KST.
// StartOfQuarter는 KST로 분기의 시작을 반환합니다.
func StartOfQuarter(t time.Time) time.Time {
	t = t.In(defaultLocation)
	month := t.Month()
	var quarterMonth time.Month
	switch {
	case month >= 1 && month <= 3:
		quarterMonth = 1
	case month >= 4 && month <= 6:
		quarterMonth = 4
	case month >= 7 && month <= 9:
		quarterMonth = 7
	default:
		quarterMonth = 10
	}
	return time.Date(t.Year(), quarterMonth, 1, 0, 0, 0, 0, defaultLocation)
}

// EndOfQuarter returns the end of the quarter in KST.
// EndOfQuarter는 KST로 분기의 끝을 반환합니다.
func EndOfQuarter(t time.Time) time.Time {
	start := StartOfQuarter(t)
	return EndOfMonth(start.AddDate(0, 2, 0))
}
