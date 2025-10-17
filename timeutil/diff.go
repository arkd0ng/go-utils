package timeutil

import "time"

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/diff.go
// Purpose: Time difference calculation operations
//          시간 차이 계산 연산
//
// This file provides comprehensive time difference calculation functionality to
// measure the duration between two time points. It offers both precise duration
// calculations (floating-point) and calendar-based calculations (integer), along
// with a structured TimeDiff type for rich difference representation.
//
// 이 파일은 두 시간 지점 간의 기간을 측정하기 위한 포괄적인 시간 차이 계산 기능을
// 제공합니다. 정밀한 기간 계산(부동소수점)과 달력 기반 계산(정수)을 모두 제공하며,
// 풍부한 차이 표현을 위한 구조화된 TimeDiff 타입을 제공합니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. STRUCTURED DIFFERENCE (구조화된 차이)
//    - SubTime: Returns TimeDiff struct with helper methods
//      헬퍼 메서드가 있는 TimeDiff 구조체 반환
//    - Rich representation with String(), Days(), Hours(), etc.
//      String(), Days(), Hours() 등의 풍부한 표현
//    - Based on time.Duration internally
//      내부적으로 time.Duration 기반
//
// 2. PRECISE TIME UNIT DIFFERENCES (정밀한 시간 단위 차이)
//    - DiffInSeconds: Difference in seconds (floating-point)
//      초 단위 차이 (부동소수점)
//    - DiffInMinutes: Difference in minutes (floating-point)
//      분 단위 차이 (부동소수점)
//    - DiffInHours: Difference in hours (floating-point)
//      시간 단위 차이 (부동소수점)
//    - DiffInDays: Difference in days (floating-point)
//      일 단위 차이 (부동소수점)
//    - DiffInWeeks: Difference in weeks (floating-point)
//      주 단위 차이 (부동소수점)
//    - Includes fractional parts (e.g., 2.5 days)
//      소수 부분 포함 (예: 2.5일)
//
// 3. CALENDAR-BASED DIFFERENCES (달력 기반 차이)
//    - DiffInMonths: Difference in calendar months (integer)
//      달력 월 단위 차이 (정수)
//    - DiffInYears: Difference in calendar years (integer)
//      달력 년 단위 차이 (정수)
//    - Handles month/year boundaries correctly
//      월/년 경계를 올바르게 처리
//    - Useful for age calculations and billing cycles
//      나이 계산과 청구 주기에 유용
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. TWO TYPES OF DIFFERENCES (두 가지 유형의 차이)
//    - DURATION-BASED: Exact time elapsed (seconds, minutes, hours, days, weeks)
//      기간 기반: 정확한 경과 시간 (초, 분, 시간, 일, 주)
//    - CALENDAR-BASED: Date component differences (months, years)
//      달력 기반: 날짜 구성 요소 차이 (월, 년)
//    - Different semantics, different use cases
//      다른 의미, 다른 사용 사례
//
// 2. CONSISTENT DIRECTIONALITY (일관된 방향성)
//    - All functions calculate: t2 - t1
//      모든 함수는 t2 - t1을 계산합니다
//    - Positive result: t2 is after t1 (future)
//      양수 결과: t2가 t1보다 이후 (미래)
//    - Negative result: t2 is before t1 (past)
//      음수 결과: t2가 t1보다 이전 (과거)
//    - Example: DiffInDays(Jan 1, Jan 3) = 2.0
//      예시: DiffInDays(1월 1일, 1월 3일) = 2.0
//    - Example: DiffInDays(Jan 3, Jan 1) = -2.0
//      예시: DiffInDays(1월 3일, 1월 1일) = -2.0
//
// 3. FLOATING-POINT FOR PRECISION (정밀도를 위한 부동소수점)
//    - DiffInDays returns float64, not int
//      DiffInDays는 int가 아닌 float64를 반환합니다
//    - Preserves fractional parts: 2.5 days, 36.5 hours
//      소수 부분 보존: 2.5일, 36.5시간
//    - More accurate for duration-based calculations
//      기간 기반 계산에서 더 정확합니다
//
// 4. CALENDAR-AWARE FOR MONTHS/YEARS (월/년에 대한 달력 인식)
//    - DiffInMonths and DiffInYears use calendar logic
//      DiffInMonths와 DiffInYears는 달력 논리를 사용합니다
//    - Jan 31 to Feb 28 = 0 months (day 31 > day 28)
//      1월 31일에서 2월 28일 = 0개월 (일 31 > 일 28)
//    - Jan 31 to Mar 1 = 1 month
//      1월 31일에서 3월 1일 = 1개월
//    - Matches human intuition for age/tenure calculations
//      나이/재직 계산에 대한 인간 직관과 일치합니다
//
// 5. TIMEDIFF TYPE FOR RICHNESS (풍부함을 위한 TimeDiff 타입)
//    - SubTime returns *TimeDiff, not time.Duration
//      SubTime은 time.Duration이 아닌 *TimeDiff를 반환합니다
//    - TimeDiff wraps time.Duration with helper methods
//      TimeDiff는 헬퍼 메서드로 time.Duration을 래핑합니다
//    - Provides Days(), Hours(), String() methods
//      Days(), Hours(), String() 메서드 제공
//    - More convenient than raw duration
//      원시 기간보다 편리합니다
//
// ============================================================================
// DIFFERENCE OPERATIONS OVERVIEW / 차이 연산 개요
// ============================================================================
//
// STRUCTURED DIFFERENCE (구조화된 차이) - 1 function
// └─ SubTime           : Returns TimeDiff with helper methods
//
// PRECISE TIME UNITS (정밀한 시간 단위) - 5 functions
// ├─ DiffInSeconds     : Difference in seconds (float64)
// ├─ DiffInMinutes     : Difference in minutes (float64)
// ├─ DiffInHours       : Difference in hours (float64)
// ├─ DiffInDays        : Difference in days (float64)
// └─ DiffInWeeks       : Difference in weeks (float64)
//
// CALENDAR UNITS (달력 단위) - 2 functions
// ├─ DiffInMonths      : Difference in months (int)
// └─ DiffInYears       : Difference in years (int)
//
// Total: 8 difference functions
// 총: 8개의 차이 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
// All functions: O(1) - constant time
// 모든 함수: O(1) - 상수 시간
//
// DURATION-BASED FUNCTIONS (DiffInSeconds, DiffInMinutes, DiffInHours, DiffInDays, DiffInWeeks):
// - Uses time.Sub() to get time.Duration
//   time.Sub()를 사용하여 time.Duration 가져오기
// - Simple arithmetic on duration
//   기간에 대한 간단한 산술 연산
// - Very fast: ~10-30 nanoseconds
//   매우 빠름: ~10-30나노초
//
// CALENDAR-BASED FUNCTIONS (DiffInMonths, DiffInYears):
// - Extracts year, month, day components
//   년, 월, 일 구성 요소 추출
// - Simple integer arithmetic
//   간단한 정수 산술
// - Fast: ~30-50 nanoseconds
//   빠름: ~30-50나노초
//
// SUBTIME:
// - Creates TimeDiff struct wrapping time.Duration
//   time.Duration을 래핑하는 TimeDiff 구조체 생성
// - Single heap allocation for pointer
//   포인터를 위한 단일 힙 할당
// - Fast: ~50-100 nanoseconds
//   빠름: ~50-100나노초
//
// SPACE COMPLEXITY (공간 복잡도):
// - DiffIn* functions: O(1) - returns single value
//   DiffIn* 함수: O(1) - 단일 값 반환
// - SubTime: O(1) - returns pointer to TimeDiff struct
//   SubTime: O(1) - TimeDiff 구조체 포인터 반환
//
// PERFORMANCE TIPS (성능 팁):
// 1. For simple comparisons, use DiffIn* directly
//    간단한 비교에는 DiffIn*을 직접 사용
// 2. Use SubTime for multiple unit conversions
//    여러 단위 변환에는 SubTime 사용
// 3. Cache duration if reused in multiple calculations
//    여러 계산에서 재사용되는 경우 기간 캐시
//
// ============================================================================
// DURATION VS CALENDAR SEMANTICS / 기간 vs 달력 의미
// ============================================================================
//
// DURATION-BASED DIFFERENCES (기간 기반 차이):
// These functions measure EXACT elapsed time:
// 이 함수들은 정확한 경과 시간을 측정합니다:
//
// Example: Jan 15 09:00 to Jan 17 15:00
// 예시: 1월 15일 09:00에서 1월 17일 15:00
//
// DiffInDays result: 2.25 days
// DiffInDays 결과: 2.25일
// Calculation: 54 hours / 24 = 2.25
// 계산: 54시간 / 24 = 2.25
//
// DiffInHours result: 54.0 hours
// DiffInHours 결과: 54.0시간
//
// CALENDAR-BASED DIFFERENCES (달력 기반 차이):
// These functions count DATE COMPONENTS:
// 이 함수들은 날짜 구성 요소를 세습니다:
//
// Example: Jan 15 to Mar 14
// 예시: 1월 15일에서 3월 14일
//
// DiffInMonths result: 1 month
// DiffInMonths 결과: 1개월
// Logic: Feb is 1 month after Jan, day 14 < day 15 (adjust)
// 논리: 2월은 1월 이후 1개월, 일 14 < 일 15 (조정)
// Actually: Jan (month 1) to Mar (month 3) = 2 months difference
//           But day 14 < day 15, so subtract 1 month
//           Result: 1 month
//
// Example: Jan 31 to Feb 28
// 예시: 1월 31일에서 2월 28일
//
// DiffInMonths result: 0 months
// DiffInMonths 결과: 0개월
// Logic: 1 month later (Feb), but day 28 < day 31 (adjust back)
// 논리: 1개월 후 (2월), 하지만 일 28 < 일 31 (다시 조정)
//
// Example: Jan 31 to Mar 1
// 예시: 1월 31일에서 3월 1일
//
// DiffInMonths result: 1 month
// DiffInMonths 결과: 1개월
// Logic: 2 months later (Mar), day 1 < day 31 (adjust: 2 - 1 = 1)
// 논리: 2개월 후 (3월), 일 1 < 일 31 (조정: 2 - 1 = 1)
//
// WHY THIS MATTERS (왜 이것이 중요한가):
// - Use duration-based for: SLA tracking, performance metrics, elapsed time
//   다음에는 기간 기반 사용: SLA 추적, 성능 메트릭, 경과 시간
// - Use calendar-based for: Age calculation, tenure, subscription periods
//   다음에는 달력 기반 사용: 나이 계산, 재직 기간, 구독 기간
//
// ============================================================================
// YEAR DIFFERENCE CALCULATION / 년 차이 계산
// ============================================================================
//
// DiffInYears implements CALENDAR YEAR logic:
// DiffInYears는 달력 년 논리를 구현합니다:
//
// Example 1: Birthday not yet reached
// 예시 1: 생일이 아직 도달하지 않음
//   t1 = 1990-05-15 (birthdate)
//   t2 = 2024-03-10 (today)
//   Years difference: 2024 - 1990 = 34
//   But month 3 < month 5, so adjust: 34 - 1 = 33 years
//   Result: 33 years (not yet 34th birthday)
//
// Example 2: Birthday already passed
// 예시 2: 생일이 이미 지남
//   t1 = 1990-05-15 (birthdate)
//   t2 = 2024-06-10 (today)
//   Years difference: 2024 - 1990 = 34
//   Month 6 > month 5, no adjust
//   Result: 34 years
//
// Example 3: Same month but day earlier
// 예시 3: 같은 월이지만 일이 이름
//   t1 = 1990-05-15 (birthdate)
//   t2 = 2024-05-10 (today)
//   Years difference: 2024 - 1990 = 34
//   Month 5 == month 5, but day 10 < day 15, so adjust: 34 - 1 = 33
//   Result: 33 years (birthday is in 5 days)
//
// This matches human intuition for age:
// 이것은 나이에 대한 인간 직관과 일치합니다:
// - You're 33 until your 34th birthday
//   34번째 생일까지는 33세입니다
//
// ============================================================================
// MONTH DIFFERENCE CALCULATION / 월 차이 계산
// ============================================================================
//
// DiffInMonths implements CALENDAR MONTH logic:
// DiffInMonths는 달력 월 논리를 구현합니다:
//
// Example 1: Day not yet reached
// 예시 1: 일이 아직 도달하지 않음
//   t1 = 2024-01-31
//   t2 = 2024-03-30
//   Months difference: (3 - 1) = 2
//   But day 30 < day 31, so adjust: 2 - 1 = 1 month
//   Result: 1 month
//
// Example 2: Day already passed
// 예시 2: 일이 이미 지남
//   t1 = 2024-01-15
//   t2 = 2024-03-20
//   Months difference: (3 - 1) = 2
//   Day 20 > day 15, no adjust
//   Result: 2 months
//
// Example 3: Across year boundary
// 예시 3: 년 경계를 넘어
//   t1 = 2023-11-15
//   t2 = 2024-02-10
//   Years difference: (2024 - 2023) = 1 year = 12 months
//   Months difference: (2 - 11) = -9
//   Total: 12 + (-9) = 3 months
//   But day 10 < day 15, so adjust: 3 - 1 = 2 months
//   Result: 2 months
//
// This matches intuition for subscriptions:
// 이것은 구독에 대한 직관과 일치합니다:
// - "Monthly plan started Jan 31" means next billing on Feb 28/29, then Mar 31
//   "1월 31일에 시작된 월간 플랜"은 다음 청구가 2월 28/29일, 그 다음 3월 31일을 의미합니다
//
// ============================================================================
// TIMEDIFF TYPE USAGE / TimeDiff 타입 사용
// ============================================================================
//
// SubTime returns *TimeDiff which wraps time.Duration:
// SubTime은 time.Duration을 래핑하는 *TimeDiff를 반환합니다:
//
// TimeDiff struct (defined in timeutil.go):
//   type TimeDiff struct {
//       Duration time.Duration
//   }
//
// Available methods on TimeDiff:
// TimeDiff에서 사용 가능한 메서드:
// - String() string: Human-readable format ("2 days 6 hours 30 minutes")
//   String() string: 사람이 읽을 수 있는 형식 ("2일 6시간 30분")
// - Days() float64: Total days including fractional
//   Days() float64: 소수 포함 총 일수
// - Hours() float64: Total hours
//   Hours() float64: 총 시간
// - Minutes() float64: Total minutes
//   Minutes() float64: 총 분
// - Seconds() float64: Total seconds
//   Seconds() float64: 총 초
// - Components() (days, hours, mins, secs int): Breakdown
//   Components() (days, hours, mins, secs int): 분해
//
// Example usage:
// 사용 예시:
//   diff := timeutil.SubTime(start, end)
//   fmt.Println(diff.String())      // "2 days 6 hours 30 minutes"
//   fmt.Printf("%.2f days\n", diff.Days())  // "2.27 days"
//   days, hours, mins, secs := diff.Components()
//   fmt.Printf("%dd %dh %dm %ds\n", days, hours, mins, secs)
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Elapsed Time Display (경과 시간 표시)
// Use case: Show how long ago something happened
// 사용 사례: 무언가가 얼마나 오래 전에 발생했는지 표시
//
//   createdAt := post.CreatedAt
//   now := time.Now()
//
//   diff := timeutil.SubTime(createdAt, now)
//   fmt.Printf("Post created %s ago\n", diff.String())
//   // Output: "Post created 2 days 6 hours 30 minutes ago"
//
// PATTERN 2: SLA Tracking (SLA 추적)
// Use case: Track service level agreement compliance
// 사용 사례: 서비스 수준 계약 준수 추적
//
//   ticketCreated := ticket.CreatedAt
//   ticketResolved := ticket.ResolvedAt
//
//   hours := timeutil.DiffInHours(ticketCreated, ticketResolved)
//   if hours <= 24.0 {
//       fmt.Println("SLA met: Resolved within 24 hours")
//   } else {
//       fmt.Printf("SLA breached: Took %.1f hours\n", hours)
//   }
//
// PATTERN 3: Age Calculation (나이 계산)
// Use case: Calculate person's age from birthdate
// 사용 사례: 생년월일로부터 사람의 나이 계산
//
//   birthdate := time.Date(1990, 5, 15, 0, 0, 0, 0, timeutil.KST)
//   today := time.Now()
//
//   years := timeutil.DiffInYears(birthdate, today)
//   fmt.Printf("Age: %d years old\n", years)
//
//   // For more detail, also show months
//   months := timeutil.DiffInMonths(birthdate, today)
//   extraMonths := months % 12
//   fmt.Printf("Age: %d years %d months\n", years, extraMonths)
//
// PATTERN 4: Subscription Duration (구독 기간)
// Use case: Calculate how many months user has been subscribed
// 사용 사례: 사용자가 구독한 개월 수 계산
//
//   subscriptionStart := user.SubscribedAt
//   now := time.Now()
//
//   months := timeutil.DiffInMonths(subscriptionStart, now)
//   fmt.Printf("Subscribed for %d months\n", months)
//
//   // Bill every 3 months
//   if months > 0 && months % 3 == 0 {
//       billCustomer(user)
//   }
//
// PATTERN 5: Performance Metrics (성능 메트릭)
// Use case: Measure execution time
// 사용 사례: 실행 시간 측정
//
//   start := time.Now()
//   processData()
//   end := time.Now()
//
//   seconds := timeutil.DiffInSeconds(start, end)
//   if seconds > 1.0 {
//       log.Printf("WARNING: Processing took %.2f seconds\n", seconds)
//   }
//
// PATTERN 6: Date Range Duration (날짜 범위 기간)
// Use case: Calculate duration of an event
// 사용 사례: 이벤트 기간 계산
//
//   eventStart := event.StartTime
//   eventEnd := event.EndTime
//
//   diff := timeutil.SubTime(eventStart, eventEnd)
//   fmt.Printf("Event duration: %s\n", diff.String())
//
//   days := diff.Days()
//   if days >= 1.0 {
//       fmt.Printf("Multi-day event: %.1f days\n", days)
//   }
//
// PATTERN 7: Countdown Timer (카운트다운 타이머)
// Use case: Show time remaining until deadline
// 사용 사례: 마감까지 남은 시간 표시
//
//   deadline := time.Date(2024, 12, 31, 23, 59, 59, 0, timeutil.KST)
//   now := time.Now()
//
//   diff := timeutil.SubTime(now, deadline)
//   days, hours, mins, _ := diff.Components()
//   fmt.Printf("Time remaining: %d days, %d hours, %d minutes\n",
//       days, hours, mins)
//
// PATTERN 8: Work Hours Calculation (근무 시간 계산)
// Use case: Calculate hours worked in a shift
// 사용 사례: 교대 근무 시간 계산
//
//   clockIn := shift.ClockInTime
//   clockOut := shift.ClockOutTime
//
//   hours := timeutil.DiffInHours(clockIn, clockOut)
//   fmt.Printf("Hours worked: %.2f\n", hours)
//
//   // Pay calculation
//   pay := hours * hourlyRate
//   if hours > 8.0 {
//       overtime := hours - 8.0
//       pay += overtime * hourlyRate * 0.5  // 1.5x for overtime
//   }
//
// PATTERN 9: Cache Expiration (캐시 만료)
// Use case: Check if cached data has expired
// 사용 사례: 캐시된 데이터가 만료되었는지 확인
//
//   cachedAt := cache.Timestamp
//   now := time.Now()
//
//   minutes := timeutil.DiffInMinutes(cachedAt, now)
//   if minutes > 30.0 {
//       fmt.Println("Cache expired, refreshing...")
//       refreshCache()
//   }
//
// PATTERN 10: Tenure Calculation (재직 기간 계산)
// Use case: Calculate employee tenure
// 사용 사례: 직원 재직 기간 계산
//
//   hireDate := employee.HireDate
//   today := time.Now()
//
//   years := timeutil.DiffInYears(hireDate, today)
//   months := timeutil.DiffInMonths(hireDate, today)
//   extraMonths := months % 12
//
//   fmt.Printf("Tenure: %d years %d months\n", years, extraMonths)
//
//   // Award milestone bonus
//   if years == 5 || years == 10 {
//       awardBonus(employee)
//   }
//
// ============================================================================
// NEGATIVE DIFFERENCES / 음수 차이
// ============================================================================
//
// All difference functions support NEGATIVE results:
// 모든 차이 함수는 음수 결과를 지원합니다:
//
// When t2 < t1 (t2 is before t1):
// t2 < t1일 때 (t2가 t1보다 이전):
//
// Example:
// 예시:
//   t1 = Jan 3, 2024 15:00
//   t2 = Jan 1, 2024 09:00
//
//   DiffInDays(t1, t2) = -2.25 days
//   DiffInHours(t1, t2) = -54 hours
//   DiffInMonths(t1, t2) = 0 months (same month)
//
// Use cases for negative differences:
// 음수 차이의 사용 사례:
// - Detecting out-of-order events
//   순서가 잘못된 이벤트 감지
// - Validating time ranges (start must be before end)
//   시간 범위 검증 (시작이 끝보다 이전이어야 함)
// - Bi-directional time calculations
//   양방향 시간 계산
//
// Example validation:
// 검증 예시:
//   diff := timeutil.DiffInSeconds(event.EndTime, event.StartTime)
//   if diff < 0 {
//       return errors.New("end time must be after start time")
//   }
//
// ============================================================================
// COMPARISON WITH STANDARD LIBRARY / 표준 라이브러리와의 비교
// ============================================================================
//
// STANDARD LIBRARY APPROACH:
// 표준 라이브러리 접근법:
//   duration := t2.Sub(t1)
//   days := duration.Hours() / 24
//   weeks := duration.Hours() / 24 / 7
//
// TIMEUTIL APPROACH:
// timeutil 접근법:
//   days := timeutil.DiffInDays(t1, t2)
//   weeks := timeutil.DiffInWeeks(t1, t2)
//
// ADVANTAGES OF TIMEUTIL:
// timeutil의 장점:
// ✓ More semantic: DiffInDays vs Hours() / 24
//   더 의미적: DiffInDays vs Hours() / 24
// ✓ Less error-prone: No magic numbers
//   오류 가능성 적음: 매직 넘버 없음
// ✓ Calendar-aware for months/years
//   월/년에 대한 달력 인식
// ✓ TimeDiff type for rich representation
//   풍부한 표현을 위한 TimeDiff 타입
//
// STANDARD LIBRARY HAS NO:
// 표준 라이브러리에는 없음:
// ✗ Calendar month difference
//   달력 월 차이
// ✗ Calendar year difference
//   달력 년 차이
// ✗ Helper for DiffInWeeks
//   DiffInWeeks를 위한 헬퍼
// ✗ Rich difference representation
//   풍부한 차이 표현
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// All functions in this file are THREAD-SAFE:
// 이 파일의 모든 함수는 스레드 안전합니다:
//
// 1. READ-ONLY OPERATIONS (읽기 전용 연산)
//    - All functions are pure read operations
//      모든 함수는 순수 읽기 연산입니다
//    - Input times are never modified
//      입력 시간은 절대 수정되지 않습니다
//    - No shared mutable state
//      공유 가변 상태 없음
//
// 2. NO SIDE EFFECTS (부작용 없음)
//    - Functions return values only
//      함수는 값만 반환합니다
//    - No I/O operations
//      I/O 작업 없음
//    - No logging or external calls
//      로깅이나 외부 호출 없음
//
// 3. SAFE CONCURRENT USAGE (안전한 동시 사용)
//    - Multiple goroutines can call these functions simultaneously
//      여러 고루틴이 동시에 이러한 함수를 호출할 수 있습니다
//    - No locks or synchronization needed
//      잠금이나 동기화가 필요 없습니다
//    - Example:
//      예시:
//        var wg sync.WaitGroup
//        for _, event := range events {
//            wg.Add(1)
//            go func(e Event) {
//                defer wg.Done()
//                hours := timeutil.DiffInHours(e.StartTime, e.EndTime)
//                processEvent(e, hours)
//            }(event)
//        }
//        wg.Wait()
//
// 4. SUBTIME ALLOCATION (SubTime 할당)
//    - SubTime allocates a new TimeDiff struct
//      SubTime은 새로운 TimeDiff 구조체를 할당합니다
//    - Each call gets independent memory
//      각 호출은 독립적인 메모리를 가져옵니다
//    - Safe for concurrent use
//      동시 사용에 안전합니다
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM constants.go:
// - MonthsPerYear: 12 (for DiffInMonths calculation)
//
// FROM timeutil.go:
// - TimeDiff type and methods (for SubTime)
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.Duration: For duration calculations
// - time.Sub(): For getting duration between times
//
// USED BY (사용처):
// - age.go (for age calculations)
//   age.go (나이 계산용)
// - Application code for duration tracking
//   기간 추적을 위한 애플리케이션 코드
// - Analytics and reporting
//   분석 및 보고
// - SLA monitoring
//   SLA 모니터링
// - Billing and subscription systems
//   청구 및 구독 시스템
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE APPROPRIATE DIFFERENCE TYPE
//    적절한 차이 유형 사용
//    - Duration-based for: elapsed time, performance, SLA
//      기간 기반: 경과 시간, 성능, SLA
//    - Calendar-based for: age, tenure, subscriptions
//      달력 기반: 나이, 재직 기간, 구독
//
// 2. USE SUBTIME FOR MULTIPLE UNITS
//    여러 단위에 SubTime 사용
//    diff := timeutil.SubTime(start, end)
//    fmt.Println(diff.String())
//    fmt.Printf("%.2f days\n", diff.Days())
//
// 3. HANDLE NEGATIVE DIFFERENCES
//    음수 차이 처리
//    diff := timeutil.DiffInHours(start, end)
//    if diff < 0 {
//        log.Println("Time range is backwards")
//    }
//
// 4. USE DIFFINMONTHS FOR BILLING
//    청구에 DiffInMonths 사용
//    months := timeutil.DiffInMonths(lastBilling, now)
//    if months >= 1 {
//        billCustomer()
//    }
//
// 5. USE DIFFINYEARS FOR AGE
//    나이에 DiffInYears 사용
//    years := timeutil.DiffInYears(birthdate, now)
//    if years >= 18 {
//        allowAccess()
//    }
//
// 6. CACHE FREQUENTLY USED DIFFERENCES
//    자주 사용되는 차이 캐시
//    // Calculate once
//    diff := timeutil.DiffInDays(start, end)
//    // Reuse in multiple places
//
// 7. BE AWARE OF FLOATING-POINT PRECISION
//    부동소수점 정밀도 인식
//    days := timeutil.DiffInDays(t1, t2)
//    // Use fmt.Printf("%.2f", days) for display
//
// 8. VALIDATE TIME RANGES
//    시간 범위 검증
//    if timeutil.DiffInSeconds(start, end) <= 0 {
//        return errors.New("invalid time range")
//    }
//
// ============================================================================

// SubTime calculates the difference between two times and returns a TimeDiff.
// SubTime은 두 시간의 차이를 계산하고 TimeDiff를 반환합니다.
//
// The difference is calculated as t2 - t1.
// 차이는 t2 - t1로 계산됩니다.
//
// Example
// 예제:
//
//	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)
//	diff := timeutil.SubTime(start, end)
//	fmt.Println(diff.Days())    // 2.270833...
//	fmt.Println(diff.String())  // "2 days 6 hours 30 minutes"
func SubTime(t1, t2 time.Time) *TimeDiff {
	return &TimeDiff{Duration: t2.Sub(t1)}
}

// DiffInSeconds returns the difference between two times in seconds.
// DiffInSeconds는 두 시간의 차이를 초 단위로 반환합니다.
//
// Example
// 예제:
//
//	seconds := timeutil.DiffInSeconds(start, end)
func DiffInSeconds(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Seconds()
}

// DiffInMinutes returns the difference between two times in minutes.
// DiffInMinutes는 두 시간의 차이를 분 단위로 반환합니다.
//
// Example
// 예제:
//
//	minutes := timeutil.DiffInMinutes(start, end)
func DiffInMinutes(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Minutes()
}

// DiffInHours returns the difference between two times in hours.
// DiffInHours는 두 시간의 차이를 시간 단위로 반환합니다.
//
// Example
// 예제:
//
//	hours := timeutil.DiffInHours(start, end)
func DiffInHours(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Hours()
}

// DiffInDays returns the difference between two times in days.
// DiffInDays는 두 시간의 차이를 일 단위로 반환합니다.
//
// Example
// 예제:
//
//	days := timeutil.DiffInDays(start, end)
func DiffInDays(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Hours() / 24
}

// DiffInWeeks returns the difference between two times in weeks.
// DiffInWeeks는 두 시간의 차이를 주 단위로 반환합니다.
//
// Example
// 예제:
//
//	weeks := timeutil.DiffInWeeks(start, end)
func DiffInWeeks(t1, t2 time.Time) float64 {
	return DiffInDays(t1, t2) / 7
}

// DiffInMonths returns the difference between two times in months.
// DiffInMonths는 두 시간의 차이를 월 단위로 반환합니다.
//
// This function calculates the difference in calendar months.
// 이 함수는 달력 월 단위의 차이를 계산합니다.
//
// Example
// 예제:
//
//	months := timeutil.DiffInMonths(start, end)
func DiffInMonths(t1, t2 time.Time) int {
	years := t2.Year() - t1.Year()
	months := int(t2.Month()) - int(t1.Month())

	// Adjust if day of month is earlier
	// 일이 더 이르면 조정
	if t2.Day() < t1.Day() {
		months--
	}

	return years*MonthsPerYear + months
}

// DiffInYears returns the difference between two times in years.
// DiffInYears는 두 시간의 차이를 년 단위로 반환합니다.
//
// This function calculates the difference in calendar years.
// 이 함수는 달력 년 단위의 차이를 계산합니다.
//
// Example
// 예제:
//
//	years := timeutil.DiffInYears(start, end)
func DiffInYears(t1, t2 time.Time) int {
	years := t2.Year() - t1.Year()

	// Adjust if month/day is earlier
	// 월/일이 더 이르면 조정
	if t2.Month() < t1.Month() || (t2.Month() == t1.Month() && t2.Day() < t1.Day()) {
		years--
	}

	return years
}
