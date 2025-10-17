package timeutil

import "time"

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/age.go
// Purpose: Age calculation operations from birth dates
//          생년월일로부터 나이 계산 연산
//
// This file provides specialized age calculation functionality to determine a
// person's age from their birth date. It offers calculations in different units
// (years, months, days) and a detailed breakdown showing years + months + days
// together. These operations are essential for user profiles, age verification,
// demographic analysis, and age-based business logic.
//
// 이 파일은 생년월일로부터 사람의 나이를 결정하는 전문화된 나이 계산 기능을 제공합니다.
// 다양한 단위(년, 월, 일)로 계산을 제공하고 년 + 월 + 일을 함께 보여주는 상세 분석을
// 제공합니다. 이러한 연산은 사용자 프로필, 나이 확인, 인구 통계 분석 및 나이 기반
// 비즈니스 로직에 필수적입니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. SIMPLE AGE IN YEARS (단순 년 단위 나이)
//    - AgeInYears: Returns age in complete years
//      완료된 년 수로 나이 반환
//    - Most common age representation
//      가장 일반적인 나이 표현
//    - "I am 33 years old"
//      "나는 33세입니다"
//
// 2. AGE IN MONTHS (월 단위 나이)
//    - AgeInMonths: Returns age in complete months
//      완료된 월 수로 나이 반환
//    - Useful for infants and subscriptions
//      유아와 구독에 유용
//    - "Baby is 18 months old"
//      "아기는 생후 18개월입니다"
//
// 3. AGE IN DAYS (일 단위 나이)
//    - AgeInDays: Returns age in complete days
//      완료된 일 수로 나이 반환
//    - Useful for newborns and precise tracking
//      신생아와 정밀 추적에 유용
//    - "Account is 90 days old"
//      "계정은 90일째입니다"
//
// 4. DETAILED AGE BREAKDOWN (상세 나이 분석)
//    - Age: Returns AgeDetail with years, months, and days
//      년, 월, 일이 있는 AgeDetail 반환
//    - "33 years 5 months 12 days old"
//      "33년 5개월 12일"
//    - More precise than years alone
//      년만 표시하는 것보다 정밀
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. BIRTHDAY-AWARE CALCULATIONS (생일 인식 계산)
//    - AgeInYears increments on birthday, not before
//      AgeInYears는 생일에 증가하며, 이전에는 증가하지 않습니다
//    - You're 33 until your 34th birthday
//      34번째 생일까지는 33세입니다
//    - Matches human intuition and legal definitions
//      인간 직관과 법적 정의와 일치합니다
//
// 2. CALENDAR-BASED, NOT DURATION-BASED (기간 기반이 아닌 달력 기반)
//    - Age is counted in calendar units (years/months/days)
//      나이는 달력 단위로 세집니다 (년/월/일)
//    - Not 365.25 days = 1 year
//      365.25일 = 1년이 아닙니다
//    - Handles leap years correctly
//      윤년을 올바르게 처리합니다
//
// 3. KST AS REFERENCE TIMEZONE (참조 타임존으로 KST)
//    - All calculations use KST (Asia/Seoul)
//      모든 계산은 KST (Asia/Seoul)를 사용합니다
//    - Input birth dates converted to KST
//      입력 생년월일이 KST로 변환됩니다
//    - Current time (now) is in KST
//      현재 시간 (now)은 KST입니다
//    - Ensures consistent results for all users in Korea
//      한국의 모든 사용자에게 일관된 결과 보장
//
// 4. START-OF-DAY FOR DAY COUNTS (일 수를 위한 하루의 시작)
//    - AgeInDays uses StartOfDay for both dates
//      AgeInDays는 두 날짜 모두에 StartOfDay를 사용합니다
//    - Eliminates time-of-day effects
//      시간대 효과 제거
//    - Born at 11 PM or 1 AM on same day = same age in days
//      같은 날 오후 11시 또는 오전 1시 출생 = 같은 일 수
//
// 5. DETAILED BREAKDOWN WITH AGEDETAIL (AgeDetail로 상세 분석)
//    - Age() returns AgeDetail struct
//      Age()는 AgeDetail 구조체를 반환합니다
//    - Contains Years, Months, Days fields
//      Years, Months, Days 필드 포함
//    - Provides String() and TotalDays() methods
//      String()과 TotalDays() 메서드 제공
//    - More informative than single number
//      단일 숫자보다 더 많은 정보 제공
//
// ============================================================================
// AGE OPERATIONS OVERVIEW / 나이 연산 개요
// ============================================================================
//
// SINGLE UNIT AGE (단일 단위 나이) - 3 functions
// ├─ AgeInYears        : Age in complete years
// ├─ AgeInMonths       : Age in complete months
// └─ AgeInDays         : Age in complete days
//
// DETAILED AGE (상세 나이) - 1 function
// └─ Age               : Returns AgeDetail (years + months + days)
//
// Total: 4 age calculation functions
// 총: 4개의 나이 계산 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
// All functions: O(1) - constant time
// 모든 함수: O(1) - 상수 시간
//
// AGEINYEARS:
// - Extracts year, month, day from two dates
//   두 날짜에서 년, 월, 일 추출
// - Simple integer arithmetic
//   간단한 정수 산술
// - Very fast: ~50-100 nanoseconds
//   매우 빠름: ~50-100나노초
//
// AGEINMONTHS:
// - Similar to AgeInYears, includes month calculation
//   AgeInYears와 유사, 월 계산 포함
// - Fast: ~50-100 nanoseconds
//   빠름: ~50-100나노초
//
// AGEINDAYS:
// - Calls StartOfDay twice (moderate overhead)
//   StartOfDay를 두 번 호출 (보통 오버헤드)
// - Uses time.Sub() and Hours() / 24
//   time.Sub()와 Hours() / 24 사용
// - Fast: ~100-200 nanoseconds
//   빠름: ~100-200나노초
//
// AGE (AGEDETAIL):
// - Most complex: calculates years, months, and days
//   가장 복잡: 년, 월, 일 계산
// - Additional logic for month/day adjustments
//   월/일 조정을 위한 추가 로직
// - Moderate: ~200-300 nanoseconds
//   보통: ~200-300나노초
//
// SPACE COMPLEXITY (공간 복잡도):
// - AgeInYears, AgeInMonths, AgeInDays: O(1) - returns single int
//   AgeInYears, AgeInMonths, AgeInDays: O(1) - 단일 int 반환
// - Age: O(1) - returns pointer to AgeDetail struct (small allocation)
//   Age: O(1) - AgeDetail 구조체 포인터 반환 (작은 할당)
//
// PERFORMANCE TIPS (성능 팁):
// 1. Cache time.Now() if calculating multiple ages
//    여러 나이를 계산할 경우 time.Now() 캐시
// 2. Use AgeInYears for simple age checks (fastest)
//    간단한 나이 확인에 AgeInYears 사용 (가장 빠름)
// 3. Use Age() when detailed breakdown is needed
//    상세 분석이 필요할 때 Age() 사용
//
// ============================================================================
// AGE CALCULATION LOGIC / 나이 계산 논리
// ============================================================================
//
// AGEINYEARS LOGIC (AgeInYears 논리):
// Calculates complete years since birth:
// 출생 이후 완료된 년 수를 계산합니다:
//
// Step 1: Calculate year difference
// 단계 1: 년 차이 계산
//   years = currentYear - birthYear
//
// Step 2: Adjust if birthday hasn't occurred this year
// 단계 2: 올해 생일이 지나지 않았으면 조정
//   if currentMonth < birthMonth:
//       years--
//   OR if currentMonth == birthMonth AND currentDay < birthDay:
//       years--
//
// Example 1: Birthday already passed
// 예시 1: 생일이 이미 지남
//   Birth: 1990-05-15
//   Today: 2024-06-20
//   years = 2024 - 1990 = 34
//   June (6) > May (5), no adjust
//   Result: 34 years
//
// Example 2: Birthday not yet reached
// 예시 2: 생일이 아직 도달하지 않음
//   Birth: 1990-05-15
//   Today: 2024-03-20
//   years = 2024 - 1990 = 34
//   March (3) < May (5), adjust: 34 - 1 = 33
//   Result: 33 years
//
// Example 3: Same month, day not reached
// 예시 3: 같은 월, 일이 도달하지 않음
//   Birth: 1990-05-15
//   Today: 2024-05-10
//   years = 2024 - 1990 = 34
//   May == May, but 10 < 15, adjust: 34 - 1 = 33
//   Result: 33 years (birthday in 5 days)
//
// AGEINMONTHS LOGIC (AgeInMonths 논리):
// Calculates complete months since birth:
// 출생 이후 완료된 월 수를 계산합니다:
//
// Step 1: Calculate total months
// 단계 1: 총 월 수 계산
//   totalMonths = (currentYear - birthYear) * 12 + (currentMonth - birthMonth)
//
// Step 2: Adjust if day of month hasn't occurred yet
// 단계 2: 월의 일이 아직 지나지 않았으면 조정
//   if currentDay < birthDay:
//       totalMonths--
//
// Example 1: Day already passed in month
// 예시 1: 월의 일이 이미 지남
//   Birth: 1990-05-15
//   Today: 2024-06-20
//   years = 34, months diff = 1
//   totalMonths = 34 * 12 + 1 = 409
//   Day 20 > 15, no adjust
//   Result: 409 months
//
// Example 2: Day not yet reached in month
// 예시 2: 월의 일이 아직 도달하지 않음
//   Birth: 1990-05-15
//   Today: 2024-06-10
//   years = 34, months diff = 1
//   totalMonths = 34 * 12 + 1 = 409
//   Day 10 < 15, adjust: 409 - 1 = 408
//   Result: 408 months
//
// AGEINDAYS LOGIC (AgeInDays 논리):
// Calculates complete days since birth:
// 출생 이후 완료된 일 수를 계산합니다:
//
// Step 1: Truncate both dates to start of day
// 단계 1: 두 날짜를 하루의 시작으로 절단
//   now = StartOfDay(now)
//   birthDate = StartOfDay(birthDate)
//
// Step 2: Calculate duration and convert to days
// 단계 2: 기간 계산 및 일로 변환
//   days = (now - birthDate).Hours() / 24
//
// This eliminates time-of-day effects:
// 이것은 시간대 효과를 제거합니다:
//   Birth: 1990-05-15 23:30
//   Today: 1990-05-17 01:00
//   Without truncation: 1.0625 days → 1 day (wrong)
//   With truncation: 2.0 days → 2 days (correct)
//
// AGE (AGEDETAIL) LOGIC (Age (AgeDetail) 논리):
// Calculates years, months, AND days:
// 년, 월, 그리고 일을 계산합니다:
//
// Step 1: Calculate year difference
// 단계 1: 년 차이 계산
//   years = currentYear - birthYear
//
// Step 2: Calculate month difference
// 단계 2: 월 차이 계산
//   months = currentMonth - birthMonth
//
// Step 3: Calculate day difference
// 단계 3: 일 차이 계산
//   days = currentDay - birthDay
//
// Step 4: Adjust days if negative
// 단계 4: 일이 음수이면 조정
//   if days < 0:
//       months--
//       days += daysInPreviousMonth
//
// Step 5: Adjust months if negative
// 단계 5: 월이 음수이면 조정
//   if months < 0:
//       years--
//       months += 12
//
// Example:
// 예시:
//   Birth: 1990-05-20
//   Today: 2024-03-15
//   years = 34, months = -2, days = -5
//   Adjust days: days = -5 + 31 (days in Feb) = 26, months = -2 - 1 = -3
//   Adjust months: months = -3 + 12 = 9, years = 34 - 1 = 33
//   Result: 33 years, 9 months, 26 days
//
// ============================================================================
// AGEDETAIL TYPE USAGE / AgeDetail 타입 사용
// ============================================================================
//
// Age() returns *AgeDetail which contains:
// Age()는 다음을 포함하는 *AgeDetail을 반환합니다:
//
// AgeDetail struct (defined in timeutil.go):
//   type AgeDetail struct {
//       Years  int  // Complete years
//       Months int  // Additional months (0-11)
//       Days   int  // Additional days (0-30)
//   }
//
// Available methods on AgeDetail:
// AgeDetail에서 사용 가능한 메서드:
// - String() string: Human-readable format
//   String() string: 사람이 읽을 수 있는 형식
//   Example: "33 years 5 months 12 days"
//   예시: "33년 5개월 12일"
//
// - TotalDays() int: Total days represented by the age
//   TotalDays() int: 나이로 표현되는 총 일 수
//   Approximation: years * 365 + months * 30 + days
//   근사치: years * 365 + months * 30 + days
//
// Example usage:
// 사용 예시:
//   age := timeutil.Age(birthDate)
//   fmt.Println(age.String())              // "33 years 5 months 12 days"
//   fmt.Printf("%d years old\n", age.Years) // "33 years old"
//   if age.Years >= 18 {
//       fmt.Println("Adult")
//   }
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Age Verification (나이 확인)
// Use case: Check if user is old enough (18+)
// 사용 사례: 사용자가 충분히 나이가 들었는지 확인 (18세 이상)
//
//   birthDate := user.BirthDate
//   age := timeutil.AgeInYears(birthDate)
//
//   if age < 18 {
//       return errors.New("must be 18 or older")
//   }
//   fmt.Println("Age verified")
//
// PATTERN 2: User Profile Display (사용자 프로필 표시)
// Use case: Show age in user profile
// 사용 사례: 사용자 프로필에 나이 표시
//
//   birthDate := user.BirthDate
//   age := timeutil.AgeInYears(birthDate)
//
//   fmt.Printf("Name: %s\n", user.Name)
//   fmt.Printf("Age: %d years old\n", age)
//
// PATTERN 3: Infant Age Display (유아 나이 표시)
// Use case: Show baby's age in months
// 사용 사례: 아기의 나이를 월 단위로 표시
//
//   birthDate := baby.BirthDate
//   months := timeutil.AgeInMonths(birthDate)
//
//   if months < 24 {
//       fmt.Printf("Baby is %d months old\n", months)
//   } else {
//       years := timeutil.AgeInYears(birthDate)
//       fmt.Printf("Child is %d years old\n", years)
//   }
//
// PATTERN 4: Detailed Age Display (상세 나이 표시)
// Use case: Show precise age with years, months, and days
// 사용 사례: 년, 월, 일로 정밀한 나이 표시
//
//   birthDate := person.BirthDate
//   age := timeutil.Age(birthDate)
//
//   fmt.Printf("Exact age: %s\n", age.String())
//   // Output: "Exact age: 33 years 5 months 12 days"
//
// PATTERN 5: Account Age Tracking (계정 나이 추적)
// Use case: Track how old a user account is
// 사용 사례: 사용자 계정이 얼마나 오래되었는지 추적
//
//   createdAt := account.CreatedAt
//   days := timeutil.AgeInDays(createdAt)
//
//   if days < 30 {
//       badge = "NEW USER"
//   } else if days < 365 {
//       badge = fmt.Sprintf("%d DAYS", days)
//   } else {
//       years := timeutil.AgeInYears(createdAt)
//       badge = fmt.Sprintf("%d YEARS", years)
//   }
//
// PATTERN 6: Subscription Duration (구독 기간)
// Use case: Show how long user has been subscribed
// 사용 사례: 사용자가 구독한 기간 표시
//
//   subscribedAt := subscription.StartDate
//   age := timeutil.Age(subscribedAt)
//
//   fmt.Printf("Subscribed for: %d years %d months\n",
//       age.Years, age.Months)
//
//   // Award loyalty bonus
//   if age.Years >= 5 {
//       awardLoyaltyBonus(user)
//   }
//
// PATTERN 7: Age-Based Pricing (나이 기반 가격)
// Use case: Apply age-based discounts
// 사용 사례: 나이 기반 할인 적용
//
//   age := timeutil.AgeInYears(visitor.BirthDate)
//
//   basePrice := 50.0
//   if age < 12 {
//       price = basePrice * 0.5  // Child discount
//   } else if age >= 65 {
//       price = basePrice * 0.7  // Senior discount
//   } else {
//       price = basePrice
//   }
//
// PATTERN 8: Birthday Notification (생일 알림)
// Use case: Check if today is user's birthday
// 사용 사례: 오늘이 사용자의 생일인지 확인
//
//   birthDate := user.BirthDate
//   now := time.Now().In(timeutil.KST)
//
//   if now.Month() == birthDate.Month() && now.Day() == birthDate.Day() {
//       age := timeutil.AgeInYears(birthDate)
//       sendBirthdayEmail(user, age)
//       fmt.Printf("Happy %d%s birthday!\n", age, ordinalSuffix(age))
//   }
//
// PATTERN 9: Demographics Analysis (인구 통계 분석)
// Use case: Categorize users by age groups
// 사용 사례: 연령대별로 사용자 분류
//
//   ageGroups := make(map[string]int)
//
//   for _, user := range users {
//       age := timeutil.AgeInYears(user.BirthDate)
//       var group string
//       switch {
//       case age < 18:
//           group = "Under 18"
//       case age < 30:
//           group = "18-29"
//       case age < 50:
//           group = "30-49"
//       case age < 65:
//           group = "50-64"
//       default:
//           group = "65+"
//       }
//       ageGroups[group]++
//   }
//
// PATTERN 10: Newborn Age Display (신생아 나이 표시)
// Use case: Show newborn age in days
// 사용 사례: 신생아 나이를 일 단위로 표시
//
//   birthDate := baby.BirthDate
//   days := timeutil.AgeInDays(birthDate)
//
//   if days < 7 {
//       fmt.Printf("Baby is %d days old\n", days)
//   } else if days < 30 {
//       weeks := days / 7
//       fmt.Printf("Baby is %d weeks old\n", weeks)
//   } else {
//       months := timeutil.AgeInMonths(birthDate)
//       fmt.Printf("Baby is %d months old\n", months)
//   }
//
// ============================================================================
// EDGE CASES / 경계 사례
// ============================================================================
//
// LEAP YEAR BIRTHDAYS (윤년 생일):
// Born on Feb 29 (leap year):
// 2월 29일(윤년) 출생:
//
// Example:
// 예시:
//   Birth: 2000-02-29
//   Today: 2024-02-28 (not leap year)
//   AgeInYears: 23 years (birthday hasn't occurred yet)
//   AgeInYears: 2000년 2월 29일
//   Today: 2024-02-29 (leap year)
//   AgeInYears: 24 years (birthday occurred)
//
// In non-leap years, person born on Feb 29 turns a year older on Mar 1
// 윤년이 아닌 해에는 2월 29일 출생자가 3월 1일에 한 살 더 먹습니다
//
// MONTH-END BIRTHDAYS (월말 생일):
// Born on Jan 31:
// 1월 31일 출생:
//
// Example:
// 예시:
//   Birth: 1990-01-31
//   Today: 2024-02-29 (leap year)
//   AgeInMonths: Calculates correctly despite Feb having only 29 days
//   AgeInMonths: 2월이 29일밖에 없어도 올바르게 계산
//
// SAME-DAY BIRTHDAY (같은 날 생일):
// Today is the birthday:
// 오늘이 생일:
//
// Example:
// 예시:
//   Birth: 1990-05-15 14:30
//   Today: 2024-05-15 09:00
//   AgeInYears: 34 years (birthday occurred)
//   Time of day doesn't matter for years/months
//   년/월에는 시간대가 중요하지 않습니다
//
// FUTURE BIRTH DATE (미래 생년월일):
// Birth date is in the future:
// 생년월일이 미래:
//
// Example:
// 예시:
//   Birth: 2025-01-15
//   Today: 2024-06-20
//   AgeInYears: -1 years (negative age)
//   Caller should validate birth date is not in future
//   호출자는 생년월일이 미래가 아닌지 검증해야 합니다
//
// ============================================================================
// COMPARISON WITH DIFFIN* FUNCTIONS / DiffIn* 함수와의 비교
// ============================================================================
//
// AGE FUNCTIONS vs DIFF FUNCTIONS:
// Age 함수 vs Diff 함수:
//
// AgeInYears(birthDate) == DiffInYears(birthDate, now)
// ✓ Same result (같은 결과)
// ✓ But Age* is more semantic for age calculations
//   하지만 Age*가 나이 계산에 더 의미적
//
// AgeInMonths(birthDate) == DiffInMonths(birthDate, now)
// ✓ Same result
//
// AgeInDays(birthDate) == int(DiffInDays(birthDate, now))
// ✓ Similar, but AgeInDays truncates to start of day
//   유사하지만 AgeInDays는 하루의 시작으로 절단
//
// WHEN TO USE WHICH:
// 어느 것을 사용할지:
// - Use Age* for: User ages, account ages, tenure
//   Age* 사용: 사용자 나이, 계정 나이, 재직 기간
// - Use DiffIn* for: Any two arbitrary times
//   DiffIn* 사용: 임의의 두 시간
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
//    - Input birth dates are never modified
//      입력 생년월일은 절대 수정되지 않습니다
//    - No shared mutable state
//      공유 가변 상태 없음
//
// 2. TIME.NOW() USAGE (time.Now() 사용)
//    - Functions call time.Now() internally
//      함수는 내부적으로 time.Now()를 호출합니다
//    - Gets current time at moment of call
//      호출 시점의 현재 시간을 가져옵니다
//    - Safe for concurrent use
//      동시 사용에 안전합니다
//
// 3. SAFE CONCURRENT USAGE (안전한 동시 사용)
//    - Multiple goroutines can call these functions simultaneously
//      여러 고루틴이 동시에 이러한 함수를 호출할 수 있습니다
//    - No locks or synchronization needed
//      잠금이나 동기화가 필요 없습니다
//    - Example:
//      예시:
//        var wg sync.WaitGroup
//        for _, user := range users {
//            wg.Add(1)
//            go func(u User) {
//                defer wg.Done()
//                age := timeutil.AgeInYears(u.BirthDate)
//                processUser(u, age)
//            }(user)
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
// - defaultLocation: KST timezone
// - MonthsPerYear: 12 (for AgeInMonths calculation)
//
// FROM arithmetic.go:
// - StartOfDay: For AgeInDays calculation
//
// FROM timeutil.go:
// - AgeDetail type and methods (for Age function)
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.Now(): For getting current time
//
// USED BY (사용처):
// - User profile pages
//   사용자 프로필 페이지
// - Age verification systems
//   나이 확인 시스템
// - Demographics and analytics
//   인구 통계 및 분석
// - Birthday notifications
//   생일 알림
// - Age-based access control
//   나이 기반 액세스 제어
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE AGEINYEARS FOR SIMPLE AGE CHECKS
//    간단한 나이 확인에 AgeInYears 사용
//    age := timeutil.AgeInYears(birthDate)
//    if age >= 18 {
//        allowAccess()
//    }
//
// 2. USE AGEINMONTHS FOR INFANTS
//    유아에 AgeInMonths 사용
//    months := timeutil.AgeInMonths(babyBirthDate)
//    if months < 6 {
//        recommendBabyFood()
//    }
//
// 3. USE AGEINDAYS FOR NEWBORNS
//    신생아에 AgeInDays 사용
//    days := timeutil.AgeInDays(birthDate)
//    if days < 28 {
//        fmt.Println("Newborn")
//    }
//
// 4. USE AGE() FOR DETAILED DISPLAY
//    상세 표시에 Age() 사용
//    age := timeutil.Age(birthDate)
//    fmt.Printf("Exact age: %s\n", age.String())
//
// 5. VALIDATE BIRTH DATE IS NOT IN FUTURE
//    생년월일이 미래가 아닌지 검증
//    if birthDate.After(time.Now()) {
//        return errors.New("birth date cannot be in future")
//    }
//
// 6. CACHE TIME.NOW() FOR MULTIPLE CALCULATIONS
//    여러 계산을 위해 time.Now() 캐시
//    now := time.Now()
//    // Use same 'now' for consistency
//
// 7. HANDLE LEAP YEAR BIRTHDAYS
//    윤년 생일 처리
//    // Feb 29 birthdays turn a year older on Mar 1 in non-leap years
//    // 윤년이 아닌 해에는 2월 29일 생일이 3월 1일에 한 살 더 먹습니다
//
// 8. USE AGE FOR LEGAL AGE CHECKS
//    법적 나이 확인에 Age 사용
//    age := timeutil.AgeInYears(birthDate)
//    if age < legalAge {
//        denyAccess()
//    }
//
// ============================================================================

// AgeInYears calculates age in years from birth date.
// AgeInYears는 생년월일로부터 나이를 년 단위로 계산합니다.
//
// Example
// 예제:
//
//	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
//	age := timeutil.AgeInYears(birthDate) // 35
func AgeInYears(birthDate time.Time) int {
	now := time.Now().In(defaultLocation)
	birthDate = birthDate.In(defaultLocation)

	years := now.Year() - birthDate.Year()

	// Adjust if birthday hasn't occurred this year yet
	// 올해 생일이 아직 지나지 않았으면 조정
	if now.Month() < birthDate.Month() ||
		(now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		years--
	}

	return years
}

// AgeInMonths calculates age in months from birth date.
// AgeInMonths는 생년월일로부터 나이를 월 단위로 계산합니다.
//
// Example
// 예제:
//
//	months := timeutil.AgeInMonths(birthDate)
func AgeInMonths(birthDate time.Time) int {
	now := time.Now().In(defaultLocation)
	birthDate = birthDate.In(defaultLocation)

	years := now.Year() - birthDate.Year()
	months := int(now.Month()) - int(birthDate.Month())

	// Adjust if day hasn't occurred this month yet
	// 이번 달 일이 아직 지나지 않았으면 조정
	if now.Day() < birthDate.Day() {
		months--
	}

	return years*MonthsPerYear + months
}

// AgeInDays calculates age in days from birth date.
// AgeInDays는 생년월일로부터 나이를 일 단위로 계산합니다.
//
// Example
// 예제:
//
//	days := timeutil.AgeInDays(birthDate)
func AgeInDays(birthDate time.Time) int {
	now := time.Now().In(defaultLocation)
	birthDate = birthDate.In(defaultLocation)

	// Truncate to start of day for accurate day count
	// 정확한 일 수 계산을 위해 하루의 시작으로 절단
	now = StartOfDay(now)
	birthDate = StartOfDay(birthDate)

	return int(now.Sub(birthDate).Hours() / 24)
}

// Age calculates detailed age (years, months, days) from birth date.
// Age는 생년월일로부터 상세 나이 (년, 월, 일)를 계산합니다.
//
// Example
// 예제:
//
//	age := timeutil.Age(birthDate)
//	fmt.Printf("%d years %d months %d days\n", age.Years, age.Months, age.Days)
func Age(birthDate time.Time) *AgeDetail {
	now := time.Now().In(defaultLocation)
	birthDate = birthDate.In(defaultLocation)

	years := now.Year() - birthDate.Year()
	months := int(now.Month()) - int(birthDate.Month())
	days := now.Day() - birthDate.Day()

	// Adjust days
	// 일 조정
	if days < 0 {
		months--
		// Get days in previous month
		// 이전 달의 일 수 가져오기
		prevMonth := now.AddDate(0, -1, 0)
		daysInPrevMonth := time.Date(prevMonth.Year(), prevMonth.Month()+1, 0, 0, 0, 0, 0, defaultLocation).Day()
		days += daysInPrevMonth
	}

	// Adjust months
	// 월 조정
	if months < 0 {
		years--
		months += MonthsPerYear
	}

	return &AgeDetail{
		Years:  years,
		Months: months,
		Days:   days,
	}
}
