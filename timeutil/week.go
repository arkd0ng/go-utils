package timeutil

import "time"

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/week.go
// Purpose: Week and day counting operations for calendar calculations
//          달력 계산을 위한 주 및 일 계산 연산
//
// This file provides operations for determining week numbers (of year and month)
// and counting days in periods (months and years). It follows ISO 8601 week
// numbering standards where weeks start on Monday and the first week of the year
// contains January 4th. These operations are essential for calendar displays,
// scheduling systems, date pickers, and any application that needs to understand
// the structure of calendar time periods.
//
// 이 파일은 주 번호(연도 및 월)를 결정하고 기간(월 및 년)의 일 수를 계산하기
// 위한 연산을 제공합니다. 주가 월요일에 시작하고 연도의 첫 주가 1월 4일을
// 포함하는 ISO 8601 주 번호 표준을 따릅니다. 이러한 연산은 달력 표시, 스케줄링
// 시스템, 날짜 선택기 및 달력 기간 구조를 이해해야 하는 모든 애플리케이션에
// 필수적입니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. WEEK OF YEAR (연도의 주)
//    - WeekOfYear: ISO 8601 week number (1-53)
//      ISO 8601 주 번호 (1-53)
//    - Weeks start on Monday
//      주는 월요일에 시작
//    - First week contains January 4th
//      첫 주는 1월 4일 포함
//
// 2. WEEK OF MONTH (월의 주)
//    - WeekOfMonth: Week within month (1-6)
//      월 내 주 (1-6)
//    - Monday-based week numbering
//      월요일 기반 주 번호
//    - Maximum 6 weeks per month
//      월당 최대 6주
//
// 3. DAYS IN MONTH (월의 일 수)
//    - DaysInMonth: Number of days in month (28-31)
//      월의 일 수 (28-31)
//    - Handles leap years automatically
//      윤년 자동 처리
//    - Varies by month and year
//      월과 연도에 따라 변동
//
// 4. DAYS IN YEAR (연도의 일 수)
//    - DaysInYear: Number of days in year (365 or 366)
//      연도의 일 수 (365 또는 366)
//    - Leap year aware
//      윤년 인식
//    - Uses IsLeapYear internally
//      내부적으로 IsLeapYear 사용
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. ISO 8601 WEEK NUMBERING (ISO 8601 주 번호)
//    - International standard for week numbering
//      주 번호의 국제 표준
//    - Week starts on Monday (not Sunday)
//      주는 월요일에 시작 (일요일 아님)
//    - First week has at least 4 days (contains Jan 4)
//      첫 주는 최소 4일 (1월 4일 포함)
//    - Used in Europe, Asia, most of the world
//      유럽, 아시아, 대부분의 세계에서 사용
//
// 2. MONDAY AS WEEK START (주 시작으로 월요일)
//    - Consistent with other timeutil functions
//      다른 timeutil 함수와 일치
//    - Follows ISO 8601 standard
//      ISO 8601 표준 따름
//    - Monday = 0, Sunday = 6 (adjusted from Go's default)
//      월요일 = 0, 일요일 = 6 (Go의 기본값에서 조정)
//
// 3. LEAP YEAR HANDLING (윤년 처리)
//    - DaysInMonth accounts for February leap years
//      DaysInMonth는 2월 윤년 고려
//    - DaysInYear returns 366 for leap years
//      DaysInYear는 윤년에 366 반환
//    - Uses existing IsLeapYear function
//      기존 IsLeapYear 함수 사용
//
// 4. MONTH BOUNDARIES (월 경계)
//    - WeekOfMonth handles partial weeks at month start
//      WeekOfMonth는 월 시작의 부분 주 처리
//    - Days before first Monday are week 1 (not week 0)
//      첫 월요일 이전의 날은 주 1 (주 0이 아님)
//    - Maximum 6 weeks per month (rare, 31-day month starting on Saturday)
//      월당 최대 6주 (드물게, 토요일에 시작하는 31일 월)
//
// 5. ACCURATE CALCULATIONS (정확한 계산)
//    - DaysInMonth uses next-month-minus-one-day method
//      DaysInMonth는 다음 월 마이너스 1일 방법 사용
//    - Handles month overflow automatically
//      월 오버플로 자동 처리
//    - No hardcoded day counts (except constants)
//      하드코딩된 일 수 없음 (상수 제외)
//
// ============================================================================
// WEEK AND DAY OPERATIONS OVERVIEW / 주 및 일 연산 개요
// ============================================================================
//
// WEEK NUMBERING (주 번호) - 2 functions
// ├─ WeekOfYear  : ISO 8601 week of year (1-53)
// └─ WeekOfMonth : Week of month (1-6)
//
// DAY COUNTING (일 계산) - 2 functions
// ├─ DaysInMonth : Days in month (28-31)
// └─ DaysInYear  : Days in year (365-366)
//
// Total: 4 week/day functions
// 총: 4개의 주/일 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
//
// WEEKOFYEAR: O(1)
//   Calls time.ISOWeek() (built-in)
//   time.ISOWeek() 호출 (내장)
//   ~50-100 nanoseconds
//
// WEEKOFMONTH: O(1)
//   Arithmetic on day of month and first weekday
//   월의 일과 첫 평일의 산술
//   ~100-200 nanoseconds
//
// DAYSINMONTH: O(1)
//   Creates next month date and subtracts
//   다음 월 날짜 생성 및 빼기
//   ~50-100 nanoseconds
//
// DAYSINYEAR: O(1)
//   Calls IsLeapYear + conditional
//   IsLeapYear 호출 + 조건
//   ~20-50 nanoseconds
//
// SPACE COMPLEXITY (공간 복잡도):
// - All functions: O(1) - no additional allocation
//   모든 함수: O(1) - 추가 할당 없음
//
// PERFORMANCE NOTES:
// 성능 참고:
// 1. DaysInYear is fastest (simple conditional)
//    DaysInYear가 가장 빠름 (간단한 조건)
// 2. WeekOfYear uses optimized built-in
//    WeekOfYear는 최적화된 내장 함수 사용
// 3. All operations are very efficient
//    모든 연산이 매우 효율적
//
// ============================================================================
// ISO 8601 WEEK NUMBERING / ISO 8601 주 번호
// ============================================================================
//
// ISO 8601 STANDARD:
// ISO 8601 표준:
//
// RULES:
// 규칙:
// 1. Weeks start on Monday
//    주는 월요일에 시작
// 2. Week 1 is the first week with at least 4 days in the new year
//    주 1은 새해에 최소 4일이 있는 첫 주
// 3. Equivalently: Week 1 contains January 4th
//    동등하게: 주 1은 1월 4일 포함
// 4. Last week is week 52 or 53
//    마지막 주는 주 52 또는 53
//
// EXAMPLES:
// 예시:
//
// Year 2024:
// - Jan 1 (Mon) is in Week 1 (starts on Monday, has 7 days in Jan)
//   1월 1일 (월)은 주 1 (월요일 시작, 1월에 7일)
// - Jan 4 (Thu) is in Week 1 (confirms Week 1)
//   1월 4일 (목)은 주 1 (주 1 확인)
//
// Year 2023:
// - Jan 1 (Sun) is in Week 52 of 2022 (only 1 day in Jan, less than 4)
//   1월 1일 (일)은 2022년의 주 52 (1월에 1일만, 4보다 적음)
// - Jan 2 (Mon) starts Week 1 of 2023
//   1월 2일 (월)은 2023년의 주 1 시작
//
// WHY ISO 8601:
// ISO 8601을 사용하는 이유:
// - International standard
//   국제 표준
// - Consistent week lengths (mostly 7 days)
//   일관된 주 길이 (대부분 7일)
// - No split weeks at year boundary
//   연도 경계에서 분할된 주 없음
// - Used in business and logistics
//   비즈니스 및 물류에 사용
//
// ============================================================================
// WEEK OF MONTH CALCULATION / 월의 주 계산
// ============================================================================
//
// ALGORITHM:
// 알고리즘:
//
// 1. Find first day of month
//    월의 첫 날 찾기
// 2. Determine weekday of first day
//    첫 날의 평일 결정
// 3. Adjust weekday: Monday=0, Sunday=6
//    평일 조정: 월요일=0, 일요일=6
// 4. Calculate: (day + firstWeekday - 1) / 7 + 1
//    계산: (일 + 첫평일 - 1) / 7 + 1
//
// EXAMPLES:
// 예시:
//
// October 2024 (starts on Tuesday):
// - Oct 1 (Tue): (1 + 1 - 1) / 7 + 1 = 1/7 + 1 = 0 + 1 = 1 (Week 1)
// - Oct 7 (Mon): (7 + 1 - 1) / 7 + 1 = 7/7 + 1 = 1 + 1 = 2 (Week 2)
// - Oct 14 (Mon): (14 + 1 - 1) / 7 + 1 = 14/7 + 1 = 2 + 1 = 3 (Week 3)
//
// February 2024 (starts on Thursday):
// - Feb 1 (Thu): (1 + 3 - 1) / 7 + 1 = 3/7 + 1 = 0 + 1 = 1 (Week 1)
// - Feb 5 (Mon): (5 + 3 - 1) / 7 + 1 = 7/7 + 1 = 1 + 1 = 2 (Week 2)
// - Feb 29 (Thu): (29 + 3 - 1) / 7 + 1 = 31/7 + 1 = 4 + 1 = 5 (Week 5)
//
// WEEK COUNT BY MONTH:
// 월별 주 수:
// - Minimum: 4 weeks (rare, 28-day month starting on Monday)
//   최소: 4주 (드물게, 월요일에 시작하는 28일 월)
// - Typical: 5 weeks (most months)
//   일반: 5주 (대부분의 월)
// - Maximum: 6 weeks (31-day month starting on Saturday)
//   최대: 6주 (토요일에 시작하는 31일 월)
//
// ============================================================================
// DAYS IN MONTH / 월의 일 수
// ============================================================================
//
// MONTH LENGTHS:
// 월 길이:
//
// 31 DAYS (7 months):
// - January, March, May, July, August, October, December
//   1월, 3월, 5월, 7월, 8월, 10월, 12월
//
// 30 DAYS (4 months):
// - April, June, September, November
//   4월, 6월, 9월, 11월
//
// 28/29 DAYS (1 month):
// - February: 28 days (normal year), 29 days (leap year)
//   2월: 28일 (평년), 29일 (윤년)
//
// MNEMONIC (기억법):
// "Thirty days hath September, April, June, and November..."
// "30일은 9월, 4월, 6월, 그리고 11월..."
//
// LEAP YEAR FEBRUARY:
// 윤년 2월:
// - Normal year: 28 days
//   평년: 28일
// - Leap year: 29 days
//   윤년: 29일
// - Leap year rules:
//   윤년 규칙:
//   * Divisible by 4: Leap year
//     4로 나누어떨어짐: 윤년
//   * Except divisible by 100: Not leap year
//     100으로 나누어떨어지는 경우 제외: 윤년 아님
//   * Except divisible by 400: Leap year
//     400으로 나누어떨어지는 경우 제외: 윤년
//
// CALCULATION METHOD:
// 계산 방법:
// - Get first day of next month
//   다음 월의 첫 날 가져오기
// - Subtract one day
//   하루 빼기
// - Day number is days in month
//   일 번호가 월의 일 수
//
// Example:
// 예시:
//   Feb 2024: Next month is Mar 1, subtract 1 day = Feb 29 → 29 days
//   Feb 2023: Next month is Mar 1, subtract 1 day = Feb 28 → 28 days
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Calendar Grid Generation (달력 그리드 생성)
// Use case: Display monthly calendar
// 사용 사례: 월별 달력 표시
//
//   year := 2024
//   month := time.October
//
//   firstDay := time.Date(year, month, 1, 0, 0, 0, 0, timeutil.KST)
//   daysInMonth := timeutil.DaysInMonth(firstDay)
//   startWeekday := int(firstDay.Weekday())
//
//   // Generate calendar grid
//   fmt.Printf("Calendar for %s %d\n", month, year)
//   fmt.Println("Mon Tue Wed Thu Fri Sat Sun")
//
//   // Print leading spaces
//   for i := 0; i < startWeekday-1; i++ {
//       fmt.Print("    ")
//   }
//
//   // Print days
//   for day := 1; day <= daysInMonth; day++ {
//       fmt.Printf("%3d ", day)
//       if (day+startWeekday-1)%7 == 0 {
//           fmt.Println()
//       }
//   }
//
// PATTERN 2: Week Number Display (주 번호 표시)
// Use case: Show week number in calendar
// 사용 사례: 달력에 주 번호 표시
//
//   today := time.Now()
//   weekOfYear := timeutil.WeekOfYear(today)
//   weekOfMonth := timeutil.WeekOfMonth(today)
//
//   fmt.Printf("Today is Week %d of %d\n", weekOfYear, today.Year())
//   fmt.Printf("Week %d of %s\n", weekOfMonth, timeutil.MonthName(today))
//
// PATTERN 3: Date Range Validation (날짜 범위 검증)
// Use case: Validate day is within month
// 사용 사례: 일이 월 내에 있는지 검증
//
//   func isValidDate(year int, month time.Month, day int) bool {
//       t := time.Date(year, month, 1, 0, 0, 0, 0, timeutil.KST)
//       maxDays := timeutil.DaysInMonth(t)
//       return day >= 1 && day <= maxDays
//   }
//
//   valid := isValidDate(2024, time.February, 29)  // true (leap year)
//   invalid := isValidDate(2023, time.February, 29) // false
//
// PATTERN 4: Weekly Report Grouping (주간 보고서 그룹화)
// Use case: Group data by week of year
// 사용 사례: 연도의 주별로 데이터 그룹화
//
//   type WeeklyData struct {
//       Year  int
//       Week  int
//       Value float64
//   }
//
//   dataByWeek := make(map[string]*WeeklyData)
//
//   for _, record := range records {
//       week := timeutil.WeekOfYear(record.Date)
//       year := record.Date.Year()
//       key := fmt.Sprintf("%d-W%02d", year, week)
//
//       if _, exists := dataByWeek[key]; !exists {
//           dataByWeek[key] = &WeeklyData{
//               Year: year,
//               Week: week,
//           }
//       }
//       dataByWeek[key].Value += record.Amount
//   }
//
// PATTERN 5: Days Remaining in Month (월의 남은 일 수)
// Use case: Calculate days until end of month
// 사용 사례: 월말까지 남은 일 수 계산
//
//   today := time.Now()
//   daysInMonth := timeutil.DaysInMonth(today)
//   currentDay := today.Day()
//   daysRemaining := daysInMonth - currentDay
//
//   fmt.Printf("%d days remaining in %s\n",
//       daysRemaining,
//       timeutil.MonthName(today))
//
// PATTERN 6: Year Progress (연도 진행률)
// Use case: Calculate year completion percentage
// 사용 사례: 연도 완료 비율 계산
//
//   today := time.Now()
//   dayOfYear := today.YearDay()
//   daysInYear := timeutil.DaysInYear(today)
//
//   progress := float64(dayOfYear) / float64(daysInYear) * 100
//   fmt.Printf("Year is %.1f%% complete\n", progress)
//
// PATTERN 7: Date Picker Constraint (날짜 선택기 제약)
// Use case: Limit selectable dates in month
// 사용 사례: 월의 선택 가능한 날짜 제한
//
//   func getDaysInMonth(year int, month time.Month) []int {
//       t := time.Date(year, month, 1, 0, 0, 0, 0, timeutil.KST)
//       maxDays := timeutil.DaysInMonth(t)
//
//       days := make([]int, maxDays)
//       for i := 0; i < maxDays; i++ {
//           days[i] = i + 1
//       }
//       return days
//   }
//
//   // For date picker dropdown
//   availableDays := getDaysInMonth(2024, time.February)
//   // [1, 2, 3, ..., 29] for leap year
//
// PATTERN 8: Billing Cycle (청구 주기)
// Use case: Calculate monthly billing period
// 사용 사례: 월별 청구 기간 계산
//
//   billingDate := user.BillingStartDate
//   year := billingDate.Year()
//   month := billingDate.Month()
//
//   // Calculate next billing date (same day next month)
//   daysInCurrentMonth := timeutil.DaysInMonth(billingDate)
//   daysInNextMonth := timeutil.DaysInMonth(
//       time.Date(year, month+1, 1, 0, 0, 0, 0, timeutil.KST))
//
//   nextBillingDay := billingDate.Day()
//   if nextBillingDay > daysInNextMonth {
//       nextBillingDay = daysInNextMonth  // Cap to last day of month
//   }
//
//   nextBilling := time.Date(year, month+1, nextBillingDay, 0, 0, 0, 0, timeutil.KST)
//
// PATTERN 9: Week-based Schedule (주 기반 스케줄)
// Use case: Schedule tasks by week number
// 사용 사례: 주 번호로 작업 스케줄
//
//   type WeeklyTask struct {
//       Week int
//       Task string
//   }
//
//   tasks := []WeeklyTask{
//       {Week: 10, Task: "Q1 Review"},
//       {Week: 26, Task: "Mid-year Review"},
//       {Week: 52, Task: "Year-end Review"},
//   }
//
//   currentWeek := timeutil.WeekOfYear(time.Now())
//   for _, task := range tasks {
//       if task.Week == currentWeek {
//           fmt.Printf("Task this week: %s\n", task.Task)
//       }
//   }
//
// PATTERN 10: Month Iteration (월 반복)
// Use case: Iterate through all days in month
// 사용 사례: 월의 모든 날짜 반복
//
//   year := 2024
//   month := time.March
//
//   firstDay := time.Date(year, month, 1, 0, 0, 0, 0, timeutil.KST)
//   daysInMonth := timeutil.DaysInMonth(firstDay)
//
//   for day := 1; day <= daysInMonth; day++ {
//       date := time.Date(year, month, day, 0, 0, 0, 0, timeutil.KST)
//       if timeutil.IsBusinessDay(date) {
//           processBusinessDay(date)
//       }
//   }
//
// ============================================================================
// EDGE CASES / 경계 사례
// ============================================================================
//
// FEBRUARY LEAP YEAR:
// 2월 윤년:
//   t2024 := time.Date(2024, 2, 1, 0, 0, 0, 0, timeutil.KST)
//   days2024 := timeutil.DaysInMonth(t2024)  // 29
//
//   t2023 := time.Date(2023, 2, 1, 0, 0, 0, 0, timeutil.KST)
//   days2023 := timeutil.DaysInMonth(t2023)  // 28
//
// WEEK 53:
// 주 53:
//   // Some years have 53 weeks
//   // 일부 연도는 53주를 가짐
//   t := time.Date(2020, 12, 31, 0, 0, 0, 0, timeutil.KST)
//   week := timeutil.WeekOfYear(t)  // 53
//
// JANUARY WEEK NUMBERS:
// 1월 주 번호:
//   // First few days of January might be in last week of previous year
//   // 1월의 처음 며칠은 이전 연도의 마지막 주에 있을 수 있음
//   t := time.Date(2023, 1, 1, 0, 0, 0, 0, timeutil.KST)
//   week := timeutil.WeekOfYear(t)  // 52 (of 2022)
//
// MONTH STARTING ON SATURDAY:
// 토요일에 시작하는 월:
//   // Maximum 6 weeks in month
//   // 월의 최대 6주
//   t := time.Date(2024, 6, 1, 0, 0, 0, 0, timeutil.KST)  // Saturday
//   // June 1-30: Weeks 1-6
//
// ZERO TIME:
// 제로 시간:
//   var t time.Time
//   week := timeutil.WeekOfYear(t)   // 1
//   days := timeutil.DaysInMonth(t)  // 31 (January)
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// THREAD-SAFE FUNCTIONS:
// 스레드 안전 함수:
// - All week/day functions are thread-safe
//   모든 주/일 함수는 스레드 안전
// - Read-only operations
//   읽기 전용 연산
// - No shared mutable state
//   공유 변경 가능 상태 없음
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
// FROM constants.go:
// - DaysPerWeek: For week calculations (value: 7)
// - DaysPerYear: For DaysInYear (value: 365)
// - DaysPerLeapYear: For DaysInYear (value: 366)
//
// FROM comparison.go:
// - IsLeapYear: For DaysInYear and DaysInMonth (February)
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.ISOWeek(): For WeekOfYear
// - time.Weekday: For weekday calculations
//
// USED BY (사용처):
// - Calendar applications
//   달력 애플리케이션
// - Date pickers
//   날짜 선택기
// - Scheduling systems
//   스케줄링 시스템
// - Reports (weekly/monthly grouping)
//   보고서 (주별/월별 그룹화)
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE WEEKOFYEAR FOR ISO STANDARD WEEKS
//    ISO 표준 주에 WeekOfYear 사용
//    week := timeutil.WeekOfYear(date)
//
// 2. USE WEEKOFMONTH FOR MONTHLY CALENDARS
//    월별 달력에 WeekOfMonth 사용
//    weekInMonth := timeutil.WeekOfMonth(date)
//
// 3. VALIDATE DATES WITH DAYSINMONTH
//    DaysInMonth로 날짜 검증
//    maxDays := timeutil.DaysInMonth(t)
//    if day > maxDays { /* invalid */ }
//
// 4. CALCULATE PROGRESS WITH DAYSINYEAR
//    DaysInYear로 진행률 계산
//    progress := dayOfYear / timeutil.DaysInYear(t)
//
// 5. REMEMBER ISO 8601 WEEKS START ON MONDAY
//    ISO 8601 주가 월요일에 시작함을 기억
//    // Not Sunday like in some systems
//
// 6. HANDLE FEBRUARY LEAP YEARS
//    2월 윤년 처리
//    // DaysInMonth automatically handles this
//
// 7. USE FOR CALENDAR GRID GENERATION
//    달력 그리드 생성에 사용
//    // Combine with weekday for proper layout
//
// 8. GROUP DATA BY WEEK NUMBER
//    주 번호로 데이터 그룹화
//    weeklyData[timeutil.WeekOfYear(date)] = value
//
// ============================================================================

// WeekOfYear returns the ISO 8601 week number of the year (1-53).
// WeekOfYear는 ISO 8601 주 번호를 반환합니다 (1-53).
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	week := timeutil.WeekOfYear(t)
//	fmt.Println(week) // 42
func WeekOfYear(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

// WeekOfMonth returns the week number of the month (1-6).
// WeekOfMonth는 월의 주 번호를 반환합니다 (1-6).
//
// Week numbering starts from the first Monday. Days before the first Monday are week 0, but returned as week 1.
// 주 번호는 첫 번째 월요일부터 시작합니다. 첫 번째 월요일 이전의 날들은 주 0이지만 주 1로 반환됩니다.
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	week := timeutil.WeekOfMonth(t)
//	fmt.Println(week) // 3
func WeekOfMonth(t time.Time) int {
	// Get first day of the month
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())

	// Get the weekday of the first day (0 = Sunday, 6 = Saturday)
	firstWeekday := int(firstDay.Weekday())

	// Adjust so Monday = 0, Sunday = 6
	if firstWeekday == 0 {
		firstWeekday = 6
	} else {
		firstWeekday--
	}

	// Get current day of month
	day := t.Day()

	// Calculate week number
	// Days before first Monday are in "week 0" but we call it week 1
	week := (day + firstWeekday - 1) / DaysPerWeek + 1

	if week < 1 {
		week = 1
	}

	return week
}

// DaysInMonth returns the number of days in the month (28-31).
// DaysInMonth는 월의 일 수를 반환합니다 (28-31).
//
// Example
// 예제:
//
//	t := time.Date(2025, 2, 1, 0, 0, 0, 0, KST)
//	days := timeutil.DaysInMonth(t)
//	fmt.Println(days) // 28
func DaysInMonth(t time.Time) int {
	// Get the first day of next month and subtract one day
	firstOfNextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	lastOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
	return lastOfMonth.Day()
}

// DaysInYear returns the number of days in the year (365 or 366).
// DaysInYear는 년의 일 수를 반환합니다 (365 또는 366).
//
// Example
// 예제:
//
//	t := time.Date(2024, 1, 1, 0, 0, 0, 0, KST)
//	days := timeutil.DaysInYear(t)
//	fmt.Println(days) // 366 (leap year)
func DaysInYear(t time.Time) int {
	if IsLeapYear(t) {
		return DaysPerLeapYear
	}
	return DaysPerYear
}
