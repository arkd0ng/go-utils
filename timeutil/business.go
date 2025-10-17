package timeutil

import (
	"sync"
	"time"
)

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/business.go
// Purpose: Business day calculations with holiday support
//          공휴일 지원이 있는 영업일 계산
//
// This file provides business day operations that are essential for commercial
// applications. It distinguishes between business days (weekdays excluding holidays)
// and non-business days (weekends and holidays), with support for custom holiday
// calendars. These operations are critical for SLA calculations, delivery estimates,
// payment processing, and any business logic that respects working days.
//
// 이 파일은 상업 애플리케이션에 필수적인 영업일 연산을 제공합니다. 영업일(공휴일을
// 제외한 평일)과 비영업일(주말 및 공휴일)을 구별하며, 커스텀 공휴일 캘린더를 지원합니다.
// 이러한 연산은 SLA 계산, 배송 예상, 결제 처리 및 근무일을 고려하는 모든 비즈니스
// 로직에 중요합니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. HOLIDAY MANAGEMENT (공휴일 관리)
//    - SetHolidays: Define custom holidays
//      커스텀 공휴일 정의
//    - GetHolidays: Retrieve configured holidays
//      설정된 공휴일 검색
//    - ClearHolidays: Remove all holidays
//      모든 공휴일 제거
//    - IsHoliday: Check if date is a holiday
//      날짜가 공휴일인지 확인
//    - Thread-safe with RWMutex
//      RWMutex로 스레드 안전
//
// 2. BUSINESS DAY CHECKING (영업일 확인)
//    - IsBusinessDay: True if weekday and not holiday
//      평일이고 공휴일이 아니면 true
//    - Excludes weekends automatically
//      주말 자동 제외
//    - Respects custom holiday calendar
//      커스텀 공휴일 캘린더 준수
//
// 3. BUSINESS DAY ARITHMETIC (영업일 산술)
//    - AddBusinessDays: Add/subtract business days
//      영업일 더하기/빼기
//    - Skips weekends and holidays
//      주말과 공휴일 건너뛰기
//    - Supports negative values (subtract)
//      음수 값 지원 (빼기)
//
// 4. NAVIGATION (탐색)
//    - NextBusinessDay: Find next business day
//      다음 영업일 찾기
//    - PreviousBusinessDay: Find previous business day
//      이전 영업일 찾기
//    - Useful for deadline calculations
//      마감일 계산에 유용
//
// 5. COUNTING (계산)
//    - CountBusinessDays: Count business days in range
//      범위 내 영업일 수 계산
//    - Inclusive start, exclusive end
//      시작 포함, 끝 제외
//    - Essential for SLA tracking
//      SLA 추적에 필수
//
// 6. KOREAN HOLIDAYS (한국 공휴일)
//    - AddKoreanHolidays: Add Korean public holidays
//      한국 공휴일 추가
//    - Includes fixed holidays (8 days)
//      고정 공휴일 포함 (8일)
//    - Note: Does not include lunar calendar holidays
//      참고: 음력 공휴일 미포함
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. CONFIGURABLE HOLIDAY CALENDAR (설정 가능한 공휴일 캘린더)
//    - No hardcoded holidays (except AddKoreanHolidays helper)
//      하드코딩된 공휴일 없음 (AddKoreanHolidays 헬퍼 제외)
//    - Application controls holiday list
//      애플리케이션이 공휴일 목록 제어
//    - Supports any country/region
//      모든 국가/지역 지원
//    - Can be updated at runtime
//      런타임에 업데이트 가능
//
// 2. THREAD-SAFE HOLIDAY STORAGE (스레드 안전 공휴일 저장)
//    - Uses sync.RWMutex for concurrent access
//      동시 액세스를 위해 sync.RWMutex 사용
//    - Multiple readers, single writer
//      여러 읽기, 단일 쓰기
//    - Safe for web applications
//      웹 애플리케이션에 안전
//
// 3. BUSINESS DAY = WEEKDAY - HOLIDAYS (영업일 = 평일 - 공휴일)
//    - Simple definition: Monday-Friday, not holiday
//      간단한 정의: 월요일-금요일, 공휴일 아님
//    - Matches most business practices
//      대부분의 비즈니스 관행과 일치
//    - Weekends always excluded
//      주말 항상 제외
//
// 4. DATE-BASED HOLIDAY MATCHING (날짜 기반 공휴일 매칭)
//    - Holidays matched by date only (YYYY-MM-DD)
//      공휴일은 날짜만으로 매칭 (YYYY-MM-DD)
//    - Time component ignored
//      시간 구성 요소 무시
//    - 2024-01-01 00:00 and 2024-01-01 14:30 both match holiday
//      2024-01-01 00:00과 2024-01-01 14:30 모두 공휴일과 매칭
//
// 5. INCLUSIVE/EXCLUSIVE RANGES (포함/배타 범위)
//    - CountBusinessDays: [start, end) - start included, end excluded
//      CountBusinessDays: [start, end) - 시작 포함, 끝 제외
//    - Standard interval convention
//      표준 간격 규칙
//    - Consistent with Go conventions
//      Go 규칙과 일치
//
// 6. KST AS REFERENCE TIMEZONE (참조 타임존으로 KST)
//    - All date comparisons in KST
//      모든 날짜 비교는 KST에서
//    - Input times converted to KST
//      입력 시간이 KST로 변환됨
//    - Ensures consistent holiday matching
//      일관된 공휴일 매칭 보장
//
// ============================================================================
// BUSINESS DAY OPERATIONS OVERVIEW / 영업일 연산 개요
// ============================================================================
//
// HOLIDAY MANAGEMENT (공휴일 관리) - 4 functions
// ├─ SetHolidays       : Set custom holidays
// ├─ GetHolidays       : Get configured holidays
// ├─ ClearHolidays     : Clear all holidays
// └─ IsHoliday         : Check if date is holiday
//
// BUSINESS DAY CHECKS (영업일 확인) - 1 function
// └─ IsBusinessDay     : Check if date is business day
//
// BUSINESS DAY ARITHMETIC (영업일 산술) - 1 function
// └─ AddBusinessDays   : Add/subtract business days
//
// NAVIGATION (탐색) - 2 functions
// ├─ NextBusinessDay     : Find next business day
// └─ PreviousBusinessDay : Find previous business day
//
// COUNTING (계산) - 1 function
// └─ CountBusinessDays : Count business days in range
//
// KOREAN HOLIDAYS (한국 공휴일) - 1 function
// └─ AddKoreanHolidays : Add Korean public holidays
//
// Total: 10 business day functions
// 총: 10개의 영업일 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
//
// SETHOLIDAYS: O(n) where n = number of holidays
//   n = 공휴일 수
//
// GETHOLIDAYS: O(n) where n = number of holidays
//   n = 공휴일 수
//
// CLEARHOLIDAYS: O(1) - just creates new empty map
//   단순히 새 빈 맵 생성
//
// ISHOLIDAY: O(1) - hash map lookup
//   해시 맵 조회
//
// ISBUSINESSDAY: O(1) - weekday check + holiday lookup
//   평일 확인 + 공휴일 조회
//
// ADDBUSINESSDAYS: O(d) where d = number of days to add
//   d = 추가할 일 수
//   Must iterate through each day
//   각 날짜를 반복해야 함
//
// NEXTBUSINESSDAY: O(k) where k = days until next business day
//   k = 다음 영업일까지의 일 수
//   Worst case: 3 days (Friday → Monday)
//   최악의 경우: 3일 (금요일 → 월요일)
//
// PREVIOUSBUSINESSDAY: O(k) where k = days to previous business day
//   k = 이전 영업일까지의 일 수
//   Worst case: 3 days (Monday → Friday)
//   최악의 경우: 3일 (월요일 → 금요일)
//
// COUNTBUSINESSDAYS: O(n) where n = |end - start| in days
//   n = 일 단위로 |end - start|
//   Must check each day in range
//   범위의 각 날짜를 확인해야 함
//
// ADDKOREANHOLIDAYS: O(1) - adds fixed 8 holidays
//   고정된 8개 공휴일 추가
//
// SPACE COMPLEXITY (공간 복잡도):
// - Holiday storage: O(h) where h = number of holidays
//   공휴일 저장: O(h), h = 공휴일 수
// - All functions: O(1) additional space
//   모든 함수: O(1) 추가 공간
//
// PERFORMANCE TIPS (성능 팁):
// 1. Keep holiday list reasonable (<1000 dates)
//    공휴일 목록을 합리적으로 유지 (<1000일)
// 2. Cache CountBusinessDays results if reused
//    재사용 시 CountBusinessDays 결과 캐시
// 3. Avoid AddBusinessDays with large day counts
//    큰 일 수로 AddBusinessDays 피하기
// 4. Use IsBusinessDay for simple checks (fastest)
//    간단한 확인에 IsBusinessDay 사용 (가장 빠름)
//
// ============================================================================
// HOLIDAY CALENDAR MANAGEMENT / 공휴일 캘린더 관리
// ============================================================================
//
// HOLIDAY STORAGE FORMAT:
// 공휴일 저장 형식:
//
// Internal representation: map[string]bool
// 내부 표현: map[string]bool
// Key: Date string in "YYYY-MM-DD" format
// 키: "YYYY-MM-DD" 형식의 날짜 문자열
// Value: Always true (presence indicates holiday)
// 값: 항상 true (존재가 공휴일을 나타냄)
//
// Example:
// 예시:
//   holidays = {
//       "2024-01-01": true,  // New Year
//       "2024-12-25": true,  // Christmas
//   }
//
// WHY MAP INSTEAD OF SLICE:
// 슬라이스 대신 맵을 사용하는 이유:
// - O(1) lookup vs O(n) for slice
//   슬라이스의 O(n)에 비해 O(1) 조회
// - Better performance for IsHoliday checks
//   IsHoliday 확인의 더 나은 성능
// - Trade-off: Slightly more memory
//   트레이드오프: 약간 더 많은 메모리
//
// THREAD SAFETY:
// 스레드 안전성:
// - sync.RWMutex protects holiday map
//   sync.RWMutex가 공휴일 맵 보호
// - Read lock for IsHoliday (concurrent reads OK)
//   IsHoliday를 위한 읽기 잠금 (동시 읽기 가능)
// - Write lock for SetHolidays/ClearHolidays
//   SetHolidays/ClearHolidays를 위한 쓰기 잠금
// - Safe for web applications with multiple goroutines
//   여러 고루틴이 있는 웹 애플리케이션에 안전
//
// ============================================================================
// KOREAN HOLIDAYS / 한국 공휴일
// ============================================================================
//
// AddKoreanHolidays adds 8 FIXED holidays:
// AddKoreanHolidays는 8개의 고정 공휴일을 추가합니다:
//
// 1. 신정 (New Year's Day): January 1
//    1월 1일
// 2. 3.1절 (Independence Movement Day): March 1
//    3월 1일
// 3. 어린이날 (Children's Day): May 5
//    5월 5일
// 4. 현충일 (Memorial Day): June 6
//    6월 6일
// 5. 광복절 (Liberation Day): August 15
//    8월 15일
// 6. 개천절 (National Foundation Day): October 3
//    10월 3일
// 7. 한글날 (Hangul Day): October 9
//    10월 9일
// 8. 크리스마스 (Christmas): December 25
//    12월 25일
//
// NOT INCLUDED (포함되지 않음):
// - 설날 (Seollal, Lunar New Year): Varies by lunar calendar
//   음력에 따라 변동
// - 석가탄신일 (Buddha's Birthday): Varies by lunar calendar
//   음력에 따라 변동
// - 추석 (Chuseok, Korean Thanksgiving): Varies by lunar calendar
//   음력에 따라 변동
//
// WHY LUNAR HOLIDAYS ARE EXCLUDED:
// 음력 공휴일이 제외된 이유:
// - Require lunar calendar calculations
//   음력 계산 필요
// - Different date each year
//   매년 날짜가 다름
// - Complex algorithm
//   복잡한 알고리즘
// - Can be added manually via SetHolidays
//   SetHolidays를 통해 수동으로 추가 가능
//
// USAGE:
// 사용법:
//   // Add 2024 Korean holidays
//   timeutil.AddKoreanHolidays(2024)
//
//   // Add 2025 Korean holidays
//   timeutil.AddKoreanHolidays(2025)
//
//   // Manually add lunar holidays
//   lunarHolidays := []time.Time{
//       time.Date(2024, 2, 10, 0, 0, 0, 0, timeutil.KST), // Seollal 2024
//       time.Date(2024, 9, 17, 0, 0, 0, 0, timeutil.KST), // Chuseok 2024
//   }
//   existing := timeutil.GetHolidays()
//   all := append(existing, lunarHolidays...)
//   timeutil.SetHolidays(all)
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Delivery Estimation (배송 예상)
// Use case: Calculate delivery date (5 business days)
// 사용 사례: 배송 날짜 계산 (5 영업일)
//
//   orderDate := time.Now()
//   deliveryDate := timeutil.AddBusinessDays(orderDate, 5)
//   fmt.Printf("Expected delivery: %s\n",
//       timeutil.FormatDate(deliveryDate))
//
// PATTERN 2: SLA Tracking (SLA 추적)
// Use case: Track response time in business days
// 사용 사례: 영업일로 응답 시간 추적
//
//   ticketCreated := ticket.CreatedAt
//   now := time.Now()
//   businessDays := timeutil.CountBusinessDays(ticketCreated, now)
//
//   if businessDays > 3 {
//       escalateTicket(ticket)
//       fmt.Printf("SLA breached: %d business days\n", businessDays)
//   }
//
// PATTERN 3: Payment Processing (결제 처리)
// Use case: Calculate payment settlement date
// 사용 사례: 결제 정산 날짜 계산
//
//   transactionDate := payment.ProcessedAt
//   settlementDate := timeutil.AddBusinessDays(transactionDate, 2)
//
//   if timeutil.IsBusinessDay(time.Now()) {
//       processSettlement(payment, settlementDate)
//   }
//
// PATTERN 4: Holiday Calendar Setup (공휴일 캘린더 설정)
// Use case: Configure holidays at application startup
// 사용 사례: 애플리케이션 시작 시 공휴일 설정
//
//   func init() {
//       // Add Korean holidays for current year
//       currentYear := time.Now().Year()
//       timeutil.AddKoreanHolidays(currentYear)
//
//       // Add custom company holidays
//       companyHolidays := []time.Time{
//           time.Date(currentYear, 8, 14, 0, 0, 0, 0, timeutil.KST),
//       }
//       existing := timeutil.GetHolidays()
//       all := append(existing, companyHolidays...)
//       timeutil.SetHolidays(all)
//   }
//
// PATTERN 5: Deadline Calculator (마감일 계산기)
// Use case: Calculate project deadline
// 사용 사례: 프로젝트 마감일 계산
//
//   startDate := project.StartDate
//   workingDays := 20  // 4 weeks of work
//
//   deadline := timeutil.AddBusinessDays(startDate, workingDays)
//   fmt.Printf("Project deadline: %s\n",
//       timeutil.FormatDate(deadline))
//
//   // Check if today is past deadline
//   today := time.Now()
//   if timeutil.IsBusinessDay(today) && today.After(deadline) {
//       sendDeadlineAlert(project)
//   }
//
// PATTERN 6: Working Day Counter (근무일 계산기)
// Use case: Count working days in a month
// 사용 사례: 월의 근무일 수 계산
//
//   now := time.Now()
//   monthStart := timeutil.StartOfMonth(now)
//   monthEnd := timeutil.EndOfMonth(now)
//
//   businessDays := timeutil.CountBusinessDays(monthStart, monthEnd.AddDate(0, 0, 1))
//   fmt.Printf("Working days this month: %d\n", businessDays)
//
// PATTERN 7: Next Available Appointment (다음 이용 가능한 예약)
// Use case: Find next business day for appointment
// 사용 사례: 예약을 위한 다음 영업일 찾기
//
//   requestedDate := appointment.RequestedDate
//
//   if !timeutil.IsBusinessDay(requestedDate) {
//       // Move to next business day
//       appointment.ScheduledDate = timeutil.NextBusinessDay(requestedDate)
//       fmt.Println("Requested date not available, moved to next business day")
//   } else {
//       appointment.ScheduledDate = requestedDate
//   }
//
// PATTERN 8: Reporting Period (보고 기간)
// Use case: Calculate report date range in business days
// 사용 사례: 영업일로 보고서 날짜 범위 계산
//
//   endDate := time.Now()
//   startDate := timeutil.AddBusinessDays(endDate, -30)
//
//   businessDays := timeutil.CountBusinessDays(startDate, endDate)
//   report := generateReport(startDate, endDate)
//   fmt.Printf("Report covers %d business days\n", businessDays)
//
// PATTERN 9: Holiday Checker (공휴일 확인기)
// Use case: Check if specific date is holiday
// 사용 사례: 특정 날짜가 공휴일인지 확인
//
//   checkDate := time.Date(2024, 12, 25, 0, 0, 0, 0, timeutil.KST)
//
//   if timeutil.IsHoliday(checkDate) {
//       fmt.Println("This is a holiday - office closed")
//   } else if timeutil.IsBusinessDay(checkDate) {
//       fmt.Println("Regular business day")
//   } else {
//       fmt.Println("Weekend")
//   }
//
// PATTERN 10: Auto-adjust Dates (날짜 자동 조정)
// Use case: Automatically adjust non-business days to next business day
// 사용 사례: 비영업일을 다음 영업일로 자동 조정
//
//   dueDate := invoice.DueDate
//
//   // If due date falls on non-business day, move to next business day
//   if !timeutil.IsBusinessDay(dueDate) {
//       dueDate = timeutil.NextBusinessDay(dueDate)
//       invoice.DueDate = dueDate
//       invoice.Adjusted = true
//       fmt.Printf("Due date adjusted to %s\n",
//           timeutil.FormatDate(dueDate))
//   }
//
// ============================================================================
// EDGE CASES / 경계 사례
// ============================================================================
//
// CONSECUTIVE HOLIDAYS:
// 연속 공휴일:
//
// Example: Long weekend with holiday on Friday
// 예시: 금요일 공휴일이 있는 긴 주말
//   Thursday → NextBusinessDay → Monday (skips Fri/Sat/Sun)
//   목요일 → NextBusinessDay → 월요일 (금/토/일 건너뛰기)
//
// ADDBUSINESSDAYS WITH ZERO:
// 0으로 AddBusinessDays:
//   AddBusinessDays(t, 0) returns t unchanged
//   AddBusinessDays(t, 0)는 t를 변경하지 않고 반환
//
// NEGATIVE BUSINESS DAYS:
// 음수 영업일:
//   AddBusinessDays(t, -5) subtracts 5 business days
//   AddBusinessDays(t, -5)는 5 영업일을 뺍니다
//
// HOLIDAY ON WEEKEND:
// 주말의 공휴일:
//   If holiday falls on Saturday, it's not a business day anyway
//   공휴일이 토요일에 해당하면 어쨌든 영업일이 아닙니다
//   No special handling needed
//   특별한 처리 불필요
//
// COUNTBUSINESSDAYS WITH REVERSED DATES:
// 역순 날짜로 CountBusinessDays:
//   Function swaps dates if start > end
//   start > end이면 함수가 날짜를 교환합니다
//   Always returns positive count
//   항상 양수 개수 반환
//
// EMPTY HOLIDAY CALENDAR:
// 빈 공휴일 캘린더:
//   If no holidays set, only weekends excluded
//   공휴일이 설정되지 않으면 주말만 제외됩니다
//   IsBusinessDay checks weekday only
//   IsBusinessDay는 평일만 확인합니다
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// THREAD-SAFE FUNCTIONS:
// 스레드 안전 함수:
// - IsHoliday, IsBusinessDay, AddBusinessDays, etc. (read operations)
//   IsHoliday, IsBusinessDay, AddBusinessDays 등 (읽기 연산)
// - Use RLock for concurrent reads
//   동시 읽기를 위해 RLock 사용
// - Safe to call from multiple goroutines
//   여러 고루틴에서 호출 안전
//
// WRITE OPERATIONS:
// 쓰기 연산:
// - SetHolidays, ClearHolidays, AddKoreanHolidays
// - Use Lock for exclusive write access
//   독점 쓰기 액세스를 위해 Lock 사용
// - Block concurrent reads during write
//   쓰기 중 동시 읽기 차단
//
// SAFE CONCURRENT USAGE:
// 안전한 동시 사용:
//   // Multiple goroutines checking business days
//   var wg sync.WaitGroup
//   for _, date := range dates {
//       wg.Add(1)
//       go func(d time.Time) {
//           defer wg.Done()
//           if timeutil.IsBusinessDay(d) {
//               processDate(d)
//           }
//       }(date)
//   }
//   wg.Wait()
//
// INITIALIZATION PATTERN:
// 초기화 패턴:
//   // Set holidays once at startup
//   func init() {
//       timeutil.AddKoreanHolidays(2024)
//   }
//
//   // Then use from multiple goroutines safely
//   // 그런 다음 여러 고루틴에서 안전하게 사용
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM constants.go:
// - defaultLocation: KST timezone
// - KST: *time.Location for Korean holidays
//
// FROM format.go:
// - FormatDate: For holiday key formatting
//
// FROM parse.go:
// - ParseDate: For GetHolidays
//
// FROM comparison.go:
// - IsWeekend: For IsBusinessDay check
//
// FROM arithmetic.go:
// - StartOfDay: For CountBusinessDays
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - sync.RWMutex: For thread-safe holiday storage
//
// USED BY (사용처):
// - E-commerce (delivery estimates)
//   전자상거래 (배송 예상)
// - Project management (deadline tracking)
//   프로젝트 관리 (마감일 추적)
// - Financial systems (settlement dates)
//   금융 시스템 (정산 날짜)
// - SLA monitoring
//   SLA 모니터링
// - Appointment scheduling
//   예약 스케줄링
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. INITIALIZE HOLIDAYS AT STARTUP
//    시작 시 공휴일 초기화
//    func init() {
//        timeutil.AddKoreanHolidays(2024)
//    }
//
// 2. USE ADDBUSINESSDAYS FOR ESTIMATES
//    예상에 AddBusinessDays 사용
//    deliveryDate := timeutil.AddBusinessDays(orderDate, 5)
//
// 3. USE COUNTBUSINESSDAYS FOR SLA
//    SLA에 CountBusinessDays 사용
//    elapsed := timeutil.CountBusinessDays(ticketCreated, now)
//
// 4. CHECK ISBUSINESSDAY BEFORE PROCESSING
//    처리 전에 IsBusinessDay 확인
//    if timeutil.IsBusinessDay(today) {
//        processOrders()
//    }
//
// 5. UPDATE HOLIDAYS ANNUALLY
//    매년 공휴일 업데이트
//    // Add new year's holidays
//    timeutil.AddKoreanHolidays(time.Now().Year())
//
// 6. INCLUDE LUNAR HOLIDAYS MANUALLY
//    음력 공휴일 수동 포함
//    // Calculate and add Seollal, Chuseok, etc.
//
// 7. HANDLE TIMEZONE CONSISTENTLY
//    타임존을 일관되게 처리
//    // All dates in KST for Korea-based application
//
// 8. CACHE COUNTBUSINESSDAYS RESULTS
//    CountBusinessDays 결과 캐시
//    // If querying same range multiple times
//
// ============================================================================

var (
	// holidays stores custom holidays
	// holidays는 커스텀 공휴일을 저장합니다
	holidays   = make(map[string]bool)
	holidaysMu sync.RWMutex
)

// SetHolidays sets custom holidays.
// SetHolidays는 커스텀 공휴일을 설정합니다.
//
// Example
// 예제:
//
//	holidays := []time.Time{
//	    time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),  // New Year
//	    time.Date(2025, 12, 25, 0, 0, 0, 0, time.UTC), // Christmas
//	}
//	timeutil.SetHolidays(holidays)
func SetHolidays(dates []time.Time) {
	holidaysMu.Lock()
	defer holidaysMu.Unlock()

	holidays = make(map[string]bool)
	for _, date := range dates {
		date = date.In(defaultLocation)
		key := FormatDate(date)
		holidays[key] = true
	}
}

// GetHolidays returns the list of custom holidays.
// GetHolidays는 커스텀 공휴일 목록을 반환합니다.
func GetHolidays() []time.Time {
	holidaysMu.RLock()
	defer holidaysMu.RUnlock()

	result := make([]time.Time, 0, len(holidays))
	for key := range holidays {
		date, _ := ParseDate(key)
		result = append(result, date)
	}
	return result
}

// ClearHolidays clears all custom holidays.
// ClearHolidays는 모든 커스텀 공휴일을 지웁니다.
func ClearHolidays() {
	holidaysMu.Lock()
	defer holidaysMu.Unlock()
	holidays = make(map[string]bool)
}

// IsHoliday checks if t is a custom holiday.
// IsHoliday는 t가 커스텀 공휴일인지 확인합니다.
func IsHoliday(t time.Time) bool {
	holidaysMu.RLock()
	defer holidaysMu.RUnlock()

	t = t.In(defaultLocation)
	key := FormatDate(t)
	return holidays[key]
}

// IsBusinessDay checks if t is a business day (Monday-Friday, excluding holidays).
// IsBusinessDay는 t가 영업일인지 확인합니다 (월-금, 공휴일 제외).
func IsBusinessDay(t time.Time) bool {
	t = t.In(defaultLocation)

	// Check if weekend
	// 주말 확인
	if IsWeekend(t) {
		return false
	}

	// Check if holiday
	// 공휴일 확인
	if IsHoliday(t) {
		return false
	}

	return true
}

// AddBusinessDays adds the specified number of business days to a time.
// AddBusinessDays는 시간에 지정된 영업일을 더합니다.
//
// Business days are Monday-Friday, excluding weekends and custom holidays.
// 영업일은 월-금이며 주말과 커스텀 공휴일을 제외합니다.
//
// Example
// 예제:
//
//	nextBusinessDay := timeutil.AddBusinessDays(time.Now(), 5)
func AddBusinessDays(t time.Time, days int) time.Time {
	t = t.In(defaultLocation)

	if days == 0 {
		return t
	}

	direction := 1
	if days < 0 {
		direction = -1
		days = -days
	}

	count := 0
	current := t

	for count < days {
		current = current.AddDate(0, 0, direction)
		if IsBusinessDay(current) {
			count++
		}
	}

	return current
}

// NextBusinessDay returns the next business day.
// NextBusinessDay는 다음 영업일을 반환합니다.
func NextBusinessDay(t time.Time) time.Time {
	t = t.In(defaultLocation)
	next := t.AddDate(0, 0, 1)

	for !IsBusinessDay(next) {
		next = next.AddDate(0, 0, 1)
	}

	return next
}

// PreviousBusinessDay returns the previous business day.
// PreviousBusinessDay는 이전 영업일을 반환합니다.
func PreviousBusinessDay(t time.Time) time.Time {
	t = t.In(defaultLocation)
	prev := t.AddDate(0, 0, -1)

	for !IsBusinessDay(prev) {
		prev = prev.AddDate(0, 0, -1)
	}

	return prev
}

// CountBusinessDays counts the number of business days between two dates.
// CountBusinessDays는 두 날짜 사이의 영업일 수를 계산합니다.
//
// The count is inclusive of start but exclusive of end.
// 카운트는 start를 포함하지만 end는 제외합니다.
//
// Example
// 예제:
//
//	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC)
//	days := timeutil.CountBusinessDays(start, end)
func CountBusinessDays(start, end time.Time) int {
	start = start.In(defaultLocation)
	end = end.In(defaultLocation)

	// Ensure start is before end
	// start가 end보다 앞에 있도록 확인
	if start.After(end) {
		start, end = end, start
	}

	// Truncate to start of day
	// 하루의 시작으로 절단
	start = StartOfDay(start)
	end = StartOfDay(end)

	count := 0
	current := start

	for current.Before(end) {
		if IsBusinessDay(current) {
			count++
		}
		current = current.AddDate(0, 0, 1)
	}

	return count
}

// AddKoreanHolidays adds common Korean public holidays for a given year.
// AddKoreanHolidays는 주어진 년도의 일반적인 한국 공휴일을 추가합니다.
//
// This function adds fixed holidays only (New Year's Day, Independence Movement Day,
// Liberation Day, National Foundation Day, Hangul Day, Christmas).
// 이 함수는 고정 공휴일만 추가합니다 (신정, 3.1절, 광복절, 개천절, 한글날, 크리스마스).
//
// Note: Lunar calendar holidays (Seollal, Chuseok) are not included.
// 참고: 음력 공휴일 (설날, 추석)은 포함되지 않습니다.
//
// Example
// 예제:
//
//	timeutil.AddKoreanHolidays(2025)
func AddKoreanHolidays(year int) {
	koreanHolidays := []time.Time{
		// New Year's Day
		// 신정
		time.Date(year, 1, 1, 0, 0, 0, 0, KST),
		// Independence Movement Day
		// 3.1절
		time.Date(year, 3, 1, 0, 0, 0, 0, KST),
		// Children's Day
		// 어린이날
		time.Date(year, 5, 5, 0, 0, 0, 0, KST),
		// Memorial Day
		// 현충일
		time.Date(year, 6, 6, 0, 0, 0, 0, KST),
		// Liberation Day
		// 광복절
		time.Date(year, 8, 15, 0, 0, 0, 0, KST),
		// National Foundation Day
		// 개천절
		time.Date(year, 10, 3, 0, 0, 0, 0, KST),
		// Hangul Day
		// 한글날
		time.Date(year, 10, 9, 0, 0, 0, 0, KST),
		// Christmas
		// 크리스마스
		time.Date(year, 12, 25, 0, 0, 0, 0, KST),
	}

	// Append to existing holidays
	// 기존 공휴일에 추가
	existing := GetHolidays()
	all := append(existing, koreanHolidays...)
	SetHolidays(all)
}
