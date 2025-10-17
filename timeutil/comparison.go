package timeutil

import "time"

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/comparison.go
// Purpose: Time and date comparison operations
//          시간 및 날짜 비교 연산
//
// This file provides comprehensive time comparison functionality for determining
// temporal relationships between times. It offers simple comparisons (before/after),
// contextual comparisons (today/yesterday), period-based comparisons (same week/month),
// and special date classifications (weekend/weekday, leap year, past/future).
//
// 이 파일은 시간 간의 시간적 관계를 결정하기 위한 포괄적인 시간 비교 기능을 제공합니다.
// 간단한 비교(전/후), 컨텍스트 비교(오늘/어제), 기간 기반 비교(같은 주/월),
// 특수 날짜 분류(주말/평일, 윤년, 과거/미래)를 제공합니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. BASIC COMPARISONS (기본 비교)
//    - IsBefore: Checks if one time is before another
//      한 시간이 다른 시간보다 이전인지 확인
//    - IsAfter: Checks if one time is after another
//      한 시간이 다른 시간보다 이후인지 확인
//    - IsBetween: Checks if time is within a range (inclusive)
//      시간이 범위 내에 있는지 확인 (포함)
//
// 2. CONTEXTUAL COMPARISONS (컨텍스트 비교)
//    - IsToday: Checks if time is today in KST
//      KST로 오늘인지 확인
//    - IsYesterday: Checks if time is yesterday in KST
//      KST로 어제인지 확인
//    - IsTomorrow: Checks if time is tomorrow in KST
//      KST로 내일인지 확인
//    - All use KST timezone for consistency
//      일관성을 위해 모두 KST 타임존 사용
//
// 3. PERIOD-BASED COMPARISONS (기간 기반 비교)
//    - IsThisWeek: Checks if time is in current week (KST)
//      현재 주인지 확인
//    - IsThisMonth: Checks if time is in current month (KST)
//      현재 월인지 확인
//    - IsThisYear: Checks if time is in current year (KST)
//      현재 년인지 확인
//
// 4. DAY TYPE CLASSIFICATIONS (요일 분류)
//    - IsWeekend: Checks if time is Saturday or Sunday
//      토요일 또는 일요일인지 확인
//    - IsWeekday: Checks if time is Monday to Friday
//      월요일부터 금요일인지 확인
//    - Useful for business logic
//      비즈니스 로직에 유용
//
// 5. SAMENESS COMPARISONS (동일성 비교)
//    - IsSameDay: Checks if two times are on the same day
//      두 시간이 같은 날인지 확인
//    - IsSameWeek: Checks if two times are in the same week
//      두 시간이 같은 주인지 확인
//    - IsSameMonth: Checks if two times are in the same month
//      두 시간이 같은 달인지 확인
//    - IsSameYear: Checks if two times are in the same year
//      두 시간이 같은 년인지 확인
//    - All ignore time component, compare date parts only
//      모두 시간 구성 요소를 무시하고 날짜 부분만 비교
//
// 6. SPECIAL DATE CHECKS (특수 날짜 확인)
//    - IsLeapYear: Checks if time is in a leap year
//      윤년인지 확인
//    - Leap year rules: divisible by 4, except centuries unless divisible by 400
//      윤년 규칙: 4로 나누어떨어지되, 세기는 400으로 나누어떨어져야 함
//
// 7. TEMPORAL DIRECTION (시간 방향)
//    - IsPast: Checks if time is before now
//      현재보다 과거인지 확인
//    - IsFuture: Checks if time is after now
//      현재보다 미래인지 확인
//    - Compared against time.Now() at call time
//      호출 시점의 time.Now()와 비교
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. SEMANTIC CLARITY (의미적 명확성)
//    - Function names clearly express their purpose
//      함수 이름이 목적을 명확히 표현합니다
//    - IsToday vs comparing dates manually - more readable
//      IsToday가 수동 날짜 비교보다 가독성이 높습니다
//    - IsWeekend vs checking weekday values - more intuitive
//      IsWeekend가 요일 값 확인보다 직관적입니다
//
// 2. KST AS DEFAULT TIMEZONE (KST를 기본 타임존으로)
//    - All contextual comparisons use KST
//      모든 컨텍스트 비교는 KST를 사용합니다
//    - IsToday checks against KST's current date
//      IsToday는 KST의 현재 날짜와 비교합니다
//    - Ensures consistent behavior across the package
//      패키지 전체의 일관된 동작 보장
//    - Example: Two users in different timezones see same "IsToday" result
//      예시: 다른 타임존의 두 사용자가 동일한 "IsToday" 결과를 봅니다
//
// 3. INCLUSIVE RANGE COMPARISONS (포괄적 범위 비교)
//    - IsBetween includes both start and end times
//      IsBetween은 시작과 끝 시간을 모두 포함합니다
//    - Uses >= and <= semantics, not > and <
//      > 와 < 가 아닌 >= 와 <= 의미 사용
//    - More intuitive for most use cases
//      대부분의 사용 사례에서 더 직관적
//
// 4. DATE-ONLY COMPARISONS (날짜만 비교)
//    - IsSameDay, IsSameWeek, IsSameMonth, IsSameYear ignore time component
//      IsSameDay, IsSameWeek, IsSameMonth, IsSameYear는 시간 구성 요소를 무시합니다
//    - 14:30 on Jan 15 and 09:00 on Jan 15 are "same day"
//      1월 15일 14:30과 1월 15일 09:00은 "같은 날"입니다
//    - Practical for business logic and user interfaces
//      비즈니스 로직과 사용자 인터페이스에 실용적
//
// 5. WEEK STARTS ON MONDAY (월요일 시작 주)
//    - IsSameWeek uses Monday as start of week
//      IsSameWeek는 월요일을 주의 시작으로 사용합니다
//    - Follows ISO 8601 standard
//      ISO 8601 표준을 따릅니다
//    - Consistent with StartOfWeek in arithmetic.go
//      arithmetic.go의 StartOfWeek와 일치합니다
//
// 6. NO MUTATION (변경 없음)
//    - All functions are read-only
//      모든 함수는 읽기 전용입니다
//    - Input times are never modified
//      입력 시간은 절대 수정되지 않습니다
//    - Safe for concurrent use
//      동시 사용에 안전합니다
//
// ============================================================================
// COMPARISON OPERATIONS OVERVIEW / 비교 연산 개요
// ============================================================================
//
// BASIC COMPARISONS (기본 비교) - 3 functions
// ├─ IsBefore          : t1 < t2
// ├─ IsAfter           : t1 > t2
// └─ IsBetween         : start <= t <= end
//
// CONTEXTUAL COMPARISONS (컨텍스트 비교) - 3 functions
// ├─ IsToday           : Same day as now (KST)
// ├─ IsYesterday       : One day before now (KST)
// └─ IsTomorrow        : One day after now (KST)
//
// PERIOD COMPARISONS (기간 비교) - 3 functions
// ├─ IsThisWeek        : Same week as now (KST)
// ├─ IsThisMonth       : Same month as now (KST)
// └─ IsThisYear        : Same year as now (KST)
//
// DAY TYPE CHECKS (요일 확인) - 2 functions
// ├─ IsWeekend         : Saturday or Sunday
// └─ IsWeekday         : Monday to Friday
//
// SAMENESS COMPARISONS (동일성 비교) - 4 functions
// ├─ IsSameDay         : Same date (year, month, day)
// ├─ IsSameWeek        : Same week (Monday-Sunday)
// ├─ IsSameMonth       : Same month (year, month)
// └─ IsSameYear        : Same year
//
// SPECIAL CHECKS (특수 확인) - 1 function
// └─ IsLeapYear        : Year is leap year
//
// TEMPORAL DIRECTION (시간 방향) - 2 functions
// ├─ IsPast            : Before now
// └─ IsFuture          : After now
//
// Total: 18 comparison functions
// 총: 18개의 비교 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
// All functions: O(1) - constant time
// 모든 함수: O(1) - 상수 시간
//
// BASIC COMPARISONS (IsBefore, IsAfter, IsBetween):
// - Uses Go's time.Before() and time.After()
//   Go의 time.Before()와 time.After() 사용
// - Compares Unix timestamps internally
//   내부적으로 Unix 타임스탬프 비교
// - Very fast: ~5-10 nanoseconds
//   매우 빠름: ~5-10나노초
//
// CONTEXTUAL COMPARISONS (IsToday, IsYesterday, IsTomorrow):
// - Calls time.Now() once
//   time.Now()를 한 번 호출
// - Timezone conversion to KST
//   KST로 타임존 변환
// - Date component extraction
//   날짜 구성 요소 추출
// - Fast: ~100-200 nanoseconds
//   빠름: ~100-200나노초
//
// PERIOD COMPARISONS (IsThisWeek, IsThisMonth, IsThisYear):
// - IsThisWeek: Calls StartOfWeek (moderate)
//   IsThisWeek: StartOfWeek 호출 (보통)
// - IsThisMonth/IsThisYear: Simple field comparison (fast)
//   IsThisMonth/IsThisYear: 간단한 필드 비교 (빠름)
// - Fast: ~50-200 nanoseconds
//   빠름: ~50-200나노초
//
// DAY TYPE CHECKS (IsWeekend, IsWeekday):
// - Extracts weekday value
//   요일 값 추출
// - Simple integer comparison
//   간단한 정수 비교
// - Very fast: ~20-30 nanoseconds
//   매우 빠름: ~20-30나노초
//
// SAMENESS COMPARISONS (IsSameDay, IsSameWeek, IsSameMonth, IsSameYear):
// - IsSameDay: Extracts year, month, day (fast)
//   IsSameDay: 년, 월, 일 추출 (빠름)
// - IsSameWeek: Calls StartOfWeek twice (moderate)
//   IsSameWeek: StartOfWeek를 두 번 호출 (보통)
// - IsSameMonth/IsSameYear: Field comparison (very fast)
//   IsSameMonth/IsSameYear: 필드 비교 (매우 빠름)
// - Fast: ~30-200 nanoseconds
//   빠름: ~30-200나노초
//
// SPECIAL CHECKS (IsLeapYear):
// - Integer modulo operations
//   정수 나머지 연산
// - Very fast: ~5-10 nanoseconds
//   매우 빠름: ~5-10나노초
//
// TEMPORAL DIRECTION (IsPast, IsFuture):
// - Calls time.Now() once
//   time.Now()를 한 번 호출
// - Uses time.Before() or time.After()
//   time.Before() 또는 time.After() 사용
// - Fast: ~50-100 nanoseconds
//   빠름: ~50-100나노초
//
// SPACE COMPLEXITY (공간 복잡도):
// - All functions: O(1) - constant space
//   모든 함수: O(1) - 상수 공간
// - No heap allocations
//   힙 할당 없음
// - Returns boolean value only
//   불린 값만 반환
//
// PERFORMANCE TIPS (성능 팁):
// 1. Cache time.Now() if calling multiple comparison functions
//    여러 비교 함수를 호출할 경우 time.Now() 캐시
// 2. Use basic comparisons (IsBefore, IsAfter) when possible
//    가능하면 기본 비교 사용 (IsBefore, IsAfter)
// 3. Avoid repeated IsSameWeek calls - cache StartOfWeek result
//    반복된 IsSameWeek 호출 피하기 - StartOfWeek 결과 캐시
//
// ============================================================================
// COMPARISON SEMANTICS / 비교 의미
// ============================================================================
//
// TIME PRECISION IN COMPARISONS (비교의 시간 정밀도):
// Go's time.Time includes nanosecond precision:
// Go의 time.Time은 나노초 정밀도를 포함합니다:
//
// IsBefore/IsAfter:
// - Compares full timestamp including nanoseconds
//   나노초를 포함한 전체 타임스탬프 비교
// - Example:
//   t1 = 14:30:00.000000000
//   t2 = 14:30:00.000000001
//   IsBefore(t1, t2) = true (1 nanosecond difference)
//
// IsSameDay/IsSameMonth/IsSameYear:
// - Ignores time component, compares date only
//   시간 구성 요소 무시, 날짜만 비교
// - Example:
//   t1 = 2024-01-15 14:30:00
//   t2 = 2024-01-15 09:00:00
//   IsSameDay(t1, t2) = true (same date, different time)
//
// INCLUSIVE VS EXCLUSIVE RANGES (포괄적 vs 배타적 범위):
// IsBetween is INCLUSIVE on both ends:
// IsBetween은 양쪽 끝을 포함합니다:
//
// IsBetween(t, start, end) checks:
//   start <= t <= end
//
// Not:
//   start < t < end  (exclusive)
//
// This is more intuitive for most use cases:
// 이것은 대부분의 사용 사례에서 더 직관적입니다:
// - "Events from 9 AM to 5 PM" includes both 9 AM and 5 PM
//   "오전 9시부터 오후 5시까지의 이벤트"는 오전 9시와 오후 5시를 모두 포함합니다
//
// LEAP YEAR RULES (윤년 규칙):
// IsLeapYear implements the standard leap year algorithm:
// IsLeapYear는 표준 윤년 알고리즘을 구현합니다:
//
// 1. Year divisible by 4: LEAP
//    4로 나누어떨어지는 년: 윤년
// 2. EXCEPT if divisible by 100: NOT LEAP
//    단, 100으로 나누어떨어지면: 윤년 아님
// 3. EXCEPT if divisible by 400: LEAP
//    단, 400으로 나누어떨어지면: 윤년
//
// Examples:
// 예시:
// - 2024: divisible by 4, not by 100 → LEAP (윤년)
// - 2100: divisible by 4 and 100, not by 400 → NOT LEAP (평년)
// - 2000: divisible by 4, 100, and 400 → LEAP (윤년)
//
// TIMEZONE CONSIDERATIONS (타임존 고려사항):
// Contextual comparisons use KST:
// 컨텍스트 비교는 KST를 사용합니다:
//
// IsToday example:
//   User in New York: 2024-01-15 10:00 EST (00:00 KST next day)
//   IsToday returns false in Korean context
//   한국 컨텍스트에서 IsToday는 false 반환
//
// This ensures all Korean users see consistent results
// 이는 모든 한국 사용자가 일관된 결과를 보도록 보장합니다
//
// WEEK DEFINITION (주 정의):
// IsSameWeek and IsThisWeek use ISO 8601 week definition:
// IsSameWeek와 IsThisWeek는 ISO 8601 주 정의를 사용합니다:
//
// - Week starts on Monday
//   주는 월요일에 시작합니다
// - Week ends on Sunday
//   주는 일요일에 끝납니다
//
// Example:
// 예시:
//   Monday, Jan 15:    Week 1
//   Tuesday, Jan 16:   Week 1
//   ...
//   Sunday, Jan 21:    Week 1
//   Monday, Jan 22:    Week 2
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Event Filtering (이벤트 필터링)
// Use case: Filter events by date range
// 사용 사례: 날짜 범위로 이벤트 필터링
//
//   start := time.Date(2024, 1, 1, 0, 0, 0, 0, timeutil.KST)
//   end := time.Date(2024, 1, 31, 23, 59, 59, 0, timeutil.KST)
//
//   var filteredEvents []Event
//   for _, event := range allEvents {
//       if timeutil.IsBetween(event.Time, start, end) {
//           filteredEvents = append(filteredEvents, event)
//       }
//   }
//
// PATTERN 2: Recent Activity Detection (최근 활동 감지)
// Use case: Show "today" badge on recent posts
// 사용 사례: 최근 게시물에 "오늘" 배지 표시
//
//   for _, post := range posts {
//       if timeutil.IsToday(post.CreatedAt) {
//           fmt.Printf("[TODAY] %s\n", post.Title)
//       } else if timeutil.IsYesterday(post.CreatedAt) {
//           fmt.Printf("[YESTERDAY] %s\n", post.Title)
//       } else {
//           fmt.Printf("[%s] %s\n",
//               timeutil.FormatDate(post.CreatedAt), post.Title)
//       }
//   }
//
// PATTERN 3: Business Hours Check (영업 시간 확인)
// Use case: Verify if current time is within business hours
// 사용 사례: 현재 시간이 영업 시간 내인지 확인
//
//   now := time.Now()
//   if timeutil.IsWeekday(now) {
//       openTime := time.Date(now.Year(), now.Month(), now.Day(),
//           9, 0, 0, 0, timeutil.KST)
//       closeTime := time.Date(now.Year(), now.Month(), now.Day(),
//           18, 0, 0, 0, timeutil.KST)
//
//       if timeutil.IsBetween(now, openTime, closeTime) {
//           fmt.Println("Currently within business hours")
//       } else {
//           fmt.Println("Outside business hours")
//       }
//   } else {
//       fmt.Println("Weekend - closed")
//   }
//
// PATTERN 4: Expiration Check (만료 확인)
// Use case: Check if subscription has expired
// 사용 사례: 구독이 만료되었는지 확인
//
//   subscription := getSubscription()
//   if timeutil.IsPast(subscription.ExpiresAt) {
//       fmt.Println("Subscription expired")
//       promptRenewal()
//   } else if timeutil.IsBetween(time.Now(),
//       timeutil.AddDays(subscription.ExpiresAt, -7),
//       subscription.ExpiresAt) {
//       fmt.Println("Subscription expiring soon!")
//   }
//
// PATTERN 5: Grouping by Period (기간별 그룹화)
// Use case: Group events by week/month/year
// 사용 사례: 주/월/년별로 이벤트 그룹화
//
//   type PeriodGroup struct {
//       Period string
//       Events []Event
//   }
//
//   var groups []PeriodGroup
//   currentGroup := PeriodGroup{Period: "This Week"}
//
//   for _, event := range events {
//       if timeutil.IsThisWeek(event.Time) {
//           currentGroup.Events = append(currentGroup.Events, event)
//       }
//   }
//   groups = append(groups, currentGroup)
//
// PATTERN 6: Duplicate Detection (중복 감지)
// Use case: Check if two timestamps represent the same day
// 사용 사례: 두 타임스탬프가 같은 날을 나타내는지 확인
//
//   lastVisit := user.LastVisitAt
//   currentVisit := time.Now()
//
//   if !timeutil.IsSameDay(lastVisit, currentVisit) {
//       user.VisitCount++
//       user.LastVisitAt = currentVisit
//       db.Update(user)
//   }
//
// PATTERN 7: Schedule Validation (일정 검증)
// Use case: Validate that start time is before end time
// 사용 사례: 시작 시간이 종료 시간보다 이른지 검증
//
//   if !timeutil.IsBefore(event.StartTime, event.EndTime) {
//       return errors.New("start time must be before end time")
//   }
//
//   if timeutil.IsPast(event.StartTime) {
//       return errors.New("cannot schedule events in the past")
//   }
//
// PATTERN 8: Weekend Pricing (주말 가격)
// Use case: Apply weekend pricing rates
// 사용 사례: 주말 가격 적용
//
//   basePrice := 100.0
//   if timeutil.IsWeekend(bookingDate) {
//       price = basePrice * 1.5  // 50% weekend surcharge
//       fmt.Println("Weekend rate applied")
//   } else {
//       price = basePrice
//   }
//
// PATTERN 9: Age Verification (나이 확인)
// Use case: Check if person is old enough (18 years)
// 사용 사례: 사람이 충분히 나이가 들었는지 확인 (18세)
//
//   eighteenYearsAgo := timeutil.AddYears(time.Now(), -18)
//   if timeutil.IsBefore(birthdate, eighteenYearsAgo) {
//       fmt.Println("Age verified: 18+")
//   } else {
//       fmt.Println("Must be 18 or older")
//   }
//
// PATTERN 10: Leap Year Handling (윤년 처리)
// Use case: Display February with correct number of days
// 사용 사례: 올바른 일수로 2월 표시
//
//   year := 2024
//   testDate := time.Date(year, 2, 1, 0, 0, 0, 0, timeutil.KST)
//
//   if timeutil.IsLeapYear(testDate) {
//       fmt.Printf("February %d has 29 days\n", year)
//   } else {
//       fmt.Printf("February %d has 28 days\n", year)
//   }
//
// ============================================================================
// COMPARISON WITH MANUAL APPROACHES / 수동 접근법과의 비교
// ============================================================================
//
// CHECKING IF TODAY: IsToday vs manual
// IsToday 대 수동 확인:
//
// timeutil.IsToday(t)
// ✓ One line (한 줄)
// ✓ Timezone-aware (타임존 인식)
// ✓ Clear intent (명확한 의도)
// ✓ Tested and reliable (테스트되고 신뢰할 수 있음)
//
// Manual:
//   now := time.Now().In(timeutil.KST)
//   t = t.In(timeutil.KST)
//   y1, m1, d1 := now.Date()
//   y2, m2, d2 := t.Date()
//   isToday := y1 == y2 && m1 == m2 && d1 == d2
// ✗ Multiple lines (여러 줄)
// ✗ Easy to make mistakes (실수하기 쉬움)
// ✗ Need to remember timezone conversion (타임존 변환을 기억해야 함)
// ✗ Less readable (가독성 낮음)
//
// CHECKING IF WEEKEND: IsWeekend vs manual
// IsWeekend 대 수동 확인:
//
// timeutil.IsWeekend(t)
// ✓ Semantic clarity (의미적 명확성)
// ✓ Self-documenting (자기 문서화)
// ✓ One line (한 줄)
//
// Manual:
//   weekday := t.Weekday()
//   isWeekend := weekday == time.Saturday || weekday == time.Sunday
// ✗ Need to remember weekday constants (요일 상수를 기억해야 함)
// ✗ Verbose (장황함)
// ✗ Less semantic (의미적으로 덜 명확)
//
// RANGE CHECK: IsBetween vs manual
// IsBetween 대 수동 확인:
//
// timeutil.IsBetween(t, start, end)
// ✓ Clear and concise (명확하고 간결)
// ✓ Handles edge cases (경계 케이스 처리)
// ✓ Inclusive semantics (포괄적 의미)
//
// Manual:
//   inRange := (t.After(start) || t.Equal(start)) &&
//              (t.Before(end) || t.Equal(end))
// ✗ Easy to get wrong (틀리기 쉬움)
// ✗ Verbose (장황함)
// ✗ Must remember to include Equal checks (Equal 확인 포함을 기억해야 함)
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
//    - Functions return boolean values only
//      함수는 불린 값만 반환합니다
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
//                if timeutil.IsToday(e.Time) {
//                    processEvent(e)
//                }
//            }(event)
//        }
//        wg.Wait()
//
// 4. TIME.NOW() CONSIDERATIONS (time.Now() 고려사항)
//    - Functions using time.Now() (IsToday, IsPast, etc.) get current time at call moment
//      time.Now()를 사용하는 함수(IsToday, IsPast 등)는 호출 시점의 현재 시간을 가져옵니다
//    - Result may differ between concurrent calls if time advances
//      시간이 진행되면 동시 호출 간에 결과가 다를 수 있습니다
//    - This is expected behavior, not a thread safety issue
//      이것은 예상된 동작이며, 스레드 안전성 문제가 아닙니다
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM constants.go:
// - defaultLocation: KST timezone for contextual comparisons
//
// FROM arithmetic.go:
// - StartOfWeek: For IsSameWeek calculation
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.Before(): For IsBefore
// - time.After(): For IsAfter
// - time.Equal(): For exact equality
// - time.Now(): For contextual comparisons
//
// USED BY (사용처):
// - Application business logic
//   애플리케이션 비즈니스 로직
// - Conditional rendering in UI
//   UI의 조건부 렌더링
// - Event filtering and sorting
//   이벤트 필터링 및 정렬
// - Access control and permissions
//   액세스 제어 및 권한
// - Scheduling and calendar systems
//   스케줄링 및 캘린더 시스템
// - relative.go (for relative time calculations)
//   relative.go (상대 시간 계산용)
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE SEMANTIC COMPARISON FUNCTIONS
//    의미적 비교 함수 사용
//    ✓ Good: timeutil.IsToday(t)
//    ✗ Avoid: Manual date extraction and comparison
//
// 2. CACHE TIME.NOW() FOR MULTIPLE COMPARISONS
//    여러 비교를 위해 time.Now() 캐시
//    now := time.Now()
//    if timeutil.IsWeekday(now) && !timeutil.IsPast(event.StartTime) {
//        // Process
//    }
//
// 3. USE ISBETWEEN FOR RANGE CHECKS
//    범위 확인에 IsBetween 사용
//    if timeutil.IsBetween(t, start, end) {
//        // Within range
//    }
//
// 4. REMEMBER INCLUSIVE SEMANTICS
//    포괄적 의미 기억
//    // IsBetween includes both endpoints
//    // IsBetween은 양쪽 끝을 포함합니다
//
// 5. USE CONTEXTUAL COMPARISONS FOR UI
//    UI에 컨텍스트 비교 사용
//    if timeutil.IsToday(post.CreatedAt) {
//        badge = "NEW"
//    }
//
// 6. VALIDATE SCHEDULES WITH COMPARISONS
//    비교로 일정 검증
//    if !timeutil.IsBefore(start, end) {
//        return errors.New("invalid time range")
//    }
//
// 7. USE ISWEEKEND FOR BUSINESS LOGIC
//    비즈니스 로직에 IsWeekend 사용
//    if timeutil.IsWeekend(date) {
//        applyWeekendPricing()
//    }
//
// 8. COMBINE WITH ARITHMETIC FUNCTIONS
//    산술 함수와 결합
//    oneWeekAgo := timeutil.AddDays(time.Now(), -7)
//    if timeutil.IsAfter(event.Time, oneWeekAgo) {
//        // Event is recent
//    }
//
// ============================================================================

// IsBefore checks if t1 is before t2.
// IsBefore는 t1이 t2보다 이전인지 확인합니다.
func IsBefore(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// IsAfter checks if t1 is after t2.
// IsAfter는 t1이 t2보다 이후인지 확인합니다.
func IsAfter(t1, t2 time.Time) bool {
	return t1.After(t2)
}

// IsBetween checks if t is between start and end.
// IsBetween은 t가 start와 end 사이에 있는지 확인합니다.
func IsBetween(t, start, end time.Time) bool {
	return (t.After(start) || t.Equal(start)) && (t.Before(end) || t.Equal(end))
}

// IsToday checks if t is today in KST.
// IsToday는 t가 KST로 오늘인지 확인합니다.
func IsToday(t time.Time) bool {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	return IsSameDay(t, now)
}

// IsYesterday checks if t is yesterday in KST.
// IsYesterday는 t가 KST로 어제인지 확인합니다.
func IsYesterday(t time.Time) bool {
	yesterday := time.Now().In(defaultLocation).AddDate(0, 0, -1)
	t = t.In(defaultLocation)
	return IsSameDay(t, yesterday)
}

// IsTomorrow checks if t is tomorrow in KST.
// IsTomorrow는 t가 KST로 내일인지 확인합니다.
func IsTomorrow(t time.Time) bool {
	tomorrow := time.Now().In(defaultLocation).AddDate(0, 0, 1)
	t = t.In(defaultLocation)
	return IsSameDay(t, tomorrow)
}

// IsThisWeek checks if t is in the current week in KST.
// IsThisWeek는 t가 KST로 이번 주인지 확인합니다.
func IsThisWeek(t time.Time) bool {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	return IsSameWeek(t, now)
}

// IsThisMonth checks if t is in the current month in KST.
// IsThisMonth는 t가 KST로 이번 달인지 확인합니다.
func IsThisMonth(t time.Time) bool {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	return IsSameMonth(t, now)
}

// IsThisYear checks if t is in the current year in KST.
// IsThisYear는 t가 KST로 올해인지 확인합니다.
func IsThisYear(t time.Time) bool {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	return IsSameYear(t, now)
}

// IsWeekend checks if t is on a weekend (Saturday or Sunday).
// IsWeekend는 t가 주말인지 (토요일 또는 일요일) 확인합니다.
func IsWeekend(t time.Time) bool {
	t = t.In(defaultLocation)
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsWeekday checks if t is on a weekday (Monday to Friday).
// IsWeekday는 t가 평일인지 (월요일부터 금요일) 확인합니다.
func IsWeekday(t time.Time) bool {
	return !IsWeekend(t)
}

// IsSameDay checks if t1 and t2 are on the same day.
// IsSameDay는 t1과 t2가 같은 날인지 확인합니다.
func IsSameDay(t1, t2 time.Time) bool {
	t1 = t1.In(defaultLocation)
	t2 = t2.In(defaultLocation)
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsSameWeek checks if t1 and t2 are in the same week.
// IsSameWeek는 t1과 t2가 같은 주인지 확인합니다.
func IsSameWeek(t1, t2 time.Time) bool {
	t1 = t1.In(defaultLocation)
	t2 = t2.In(defaultLocation)
	start1 := StartOfWeek(t1)
	start2 := StartOfWeek(t2)
	return IsSameDay(start1, start2)
}

// IsSameMonth checks if t1 and t2 are in the same month.
// IsSameMonth는 t1과 t2가 같은 달인지 확인합니다.
func IsSameMonth(t1, t2 time.Time) bool {
	t1 = t1.In(defaultLocation)
	t2 = t2.In(defaultLocation)
	return t1.Year() == t2.Year() && t1.Month() == t2.Month()
}

// IsSameYear checks if t1 and t2 are in the same year.
// IsSameYear는 t1과 t2가 같은 년인지 확인합니다.
func IsSameYear(t1, t2 time.Time) bool {
	t1 = t1.In(defaultLocation)
	t2 = t2.In(defaultLocation)
	return t1.Year() == t2.Year()
}

// IsLeapYear checks if t is in a leap year.
// IsLeapYear는 t가 윤년인지 확인합니다.
func IsLeapYear(t time.Time) bool {
	year := t.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// IsPast checks if t is in the past.
// IsPast는 t가 과거인지 확인합니다.
func IsPast(t time.Time) bool {
	return t.Before(time.Now())
}

// IsFuture checks if t is in the future.
// IsFuture는 t가 미래인지 확인합니다.
func IsFuture(t time.Time) bool {
	return t.After(time.Now())
}
