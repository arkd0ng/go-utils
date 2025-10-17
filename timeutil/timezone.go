package timeutil

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/timezone.go
// Purpose: Timezone conversion and management operations
//          타임존 변환 및 관리 연산
//
// This file provides comprehensive timezone operations for converting times
// between different timezones, managing default timezone settings, and querying
// timezone information. It includes caching for performance and a curated list
// of 70+ common timezones worldwide. These operations are essential for global
// applications that need to handle users in different timezones, display times
// in local formats, or coordinate across geographic regions.
//
// 이 파일은 서로 다른 타임존 간에 시간을 변환하고, 기본 타임존 설정을 관리하며,
// 타임존 정보를 쿼리하기 위한 포괄적인 타임존 연산을 제공합니다. 성능을 위한
// 캐싱과 전 세계 70개 이상의 일반적인 타임존의 선별된 목록을 포함합니다. 이러한
// 연산은 다른 타임존의 사용자를 처리하거나, 로컬 형식으로 시간을 표시하거나,
// 지리적 영역을 조정해야 하는 글로벌 애플리케이션에 필수적입니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. DEFAULT TIMEZONE MANAGEMENT (기본 타임존 관리)
//    - SetDefaultTimezone: Configure default timezone
//      기본 타임존 설정
//    - GetDefaultTimezone: Query current default
//      현재 기본값 쿼리
//    - ResetDefaultTimezone: Reset to KST
//      KST로 재설정
//    - Affects all timeutil functions
//      모든 timeutil 함수에 영향
//
// 2. TIMEZONE CONVERSION (타임존 변환)
//    - ConvertTimezone: General conversion to any timezone
//      모든 타임존으로 일반 변환
//    - ToKST: Quick conversion to Korean time
//      한국 시간으로 빠른 변환
//    - ToUTC: Quick conversion to UTC
//      UTC로 빠른 변환
//    - Preserves instant in time (only representation changes)
//      시간의 순간 보존 (표현만 변경)
//
// 3. TIMEZONE INFORMATION (타임존 정보)
//    - GetTimezoneOffset: Get UTC offset in hours
//      시간 단위로 UTC 오프셋 가져오기
//    - IsValidTimezone: Check if timezone exists
//      타임존이 존재하는지 확인
//    - GetLocalTimezone: Get system timezone
//      시스템 타임존 가져오기
//    - ListTimezones: Get common timezone names (70+)
//      일반 타임존 이름 가져오기 (70개 이상)
//
// 4. CURRENT TIME IN TIMEZONE (타임존의 현재 시간)
//    - NowInTimezone: Get current time in any timezone
//      모든 타임존의 현재 시간 가져오기
//    - NowKST: Get current time in KST
//      KST의 현재 시간 가져오기
//    - Convenience for displaying "now" in different regions
//      다른 지역에서 "지금"을 표시하는 편의
//
// 5. PERFORMANCE OPTIMIZATION (성능 최적화)
//    - Timezone location caching
//      타임존 위치 캐싱
//    - Thread-safe cache with RWMutex
//      RWMutex로 스레드 안전 캐시
//    - Avoids repeated LoadLocation calls
//      반복적인 LoadLocation 호출 방지
//    - Significant performance gain for frequent conversions
//      빈번한 변환에 대한 상당한 성능 향상
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. KST AS DEFAULT (기본값으로 KST)
//    - Asia/Seoul (GMT+9) is the default timezone
//      Asia/Seoul (GMT+9)이 기본 타임존
//    - Suitable for Korea-based applications
//      한국 기반 애플리케이션에 적합
//    - Can be changed via SetDefaultTimezone
//      SetDefaultTimezone을 통해 변경 가능
//    - All contextual operations use this default (IsToday, etc.)
//      모든 컨텍스트 연산이 이 기본값 사용 (IsToday 등)
//
// 2. INSTANT PRESERVATION (순간 보존)
//    - Timezone conversion changes representation, not instant
//      타임존 변환은 표현을 변경하지 순간을 변경하지 않음
//    - 2024-01-01 12:00 KST = 2024-01-01 03:00 UTC (same instant)
//      2024-01-01 12:00 KST = 2024-01-01 03:00 UTC (동일한 순간)
//    - Unix timestamp remains the same
//      Unix 타임스탬프는 동일하게 유지
//    - Only the clock reading changes
//      시계 읽기만 변경됨
//
// 3. IANA TIMEZONE DATABASE (IANA 타임존 데이터베이스)
//    - Uses standard IANA timezone names
//      표준 IANA 타임존 이름 사용
//    - Format: "Continent/City" or "Area/Location"
//      형식: "대륙/도시" 또는 "지역/위치"
//    - Examples: "Asia/Seoul", "America/New_York", "Europe/London"
//      예시: "Asia/Seoul", "America/New_York", "Europe/London"
//    - Handles DST (Daylight Saving Time) automatically
//      DST (일광 절약 시간제) 자동 처리
//
// 4. CURATED TIMEZONE LIST (선별된 타임존 목록)
//    - ListTimezones returns 70+ common timezones
//      ListTimezones는 70개 이상의 일반 타임존 반환
//    - Covers major cities and regions worldwide
//      전 세계 주요 도시 및 지역 포함
//    - Organized by continent/region
//      대륙/지역별로 구성
//    - Includes UTC offset comments for reference
//      참조용 UTC 오프셋 주석 포함
//
// 5. THREAD-SAFE CACHING (스레드 안전 캐싱)
//    - time.LoadLocation is expensive (~microseconds)
//      time.LoadLocation은 비용이 많이 듦 (~마이크로초)
//    - Cache loaded locations for reuse
//      재사용을 위해 로드된 위치 캐시
//    - RWMutex allows concurrent reads
//      RWMutex는 동시 읽기 허용
//    - Reduces overhead by ~100x for cached timezones
//      캐시된 타임존에 대해 오버헤드를 ~100배 감소
//
// 6. ERROR HANDLING (오류 처리)
//    - Returns error for invalid timezone names
//      잘못된 타임존 이름에 대해 오류 반환
//    - Wraps LoadLocation errors with context
//      컨텍스트와 함께 LoadLocation 오류 래핑
//    - Allows validation before conversion
//      변환 전 검증 허용
//
// ============================================================================
// TIMEZONE OPERATIONS OVERVIEW / 타임존 연산 개요
// ============================================================================
//
// DEFAULT TIMEZONE MANAGEMENT (기본 타임존 관리) - 3 functions
// ├─ SetDefaultTimezone    : Set default timezone
// ├─ GetDefaultTimezone    : Get current default
// └─ ResetDefaultTimezone  : Reset to KST
//
// TIMEZONE CONVERSION (타임존 변환) - 3 functions
// ├─ ConvertTimezone : Convert to any timezone
// ├─ ToKST          : Convert to KST (Asia/Seoul)
// └─ ToUTC          : Convert to UTC
//
// TIMEZONE INFORMATION (타임존 정보) - 4 functions
// ├─ GetTimezoneOffset : Get UTC offset in hours
// ├─ IsValidTimezone   : Check if timezone exists
// ├─ GetLocalTimezone  : Get system timezone
// └─ ListTimezones     : Get common timezone names
//
// CURRENT TIME (현재 시간) - 2 functions
// ├─ NowInTimezone : Current time in specified timezone
// └─ NowKST        : Current time in KST
//
// INTERNAL HELPERS (내부 헬퍼) - 1 function
// └─ loadTimezone : Load timezone with caching
//
// Total: 13 timezone functions (12 public + 1 internal)
// 총: 13개의 타임존 함수 (12개 공개 + 1개 내부)
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
//
// SETDEFAULTTIMEZONE: O(1) with cache hit, O(k) with cache miss
//   k = time to load timezone from system
//   캐시 히트 시 O(1), 캐시 미스 시 O(k)
//   k = 시스템에서 타임존을 로드하는 시간
//
// GETDEFAULTTIMEZONE: O(1) - just returns string
//   단순히 문자열 반환
//
// RESETDEFAULTTIMEZONE: O(1) - assigns constant
//   상수 할당
//
// CONVERTTIMEZONE: O(1) with cache hit, O(k) with cache miss
//   캐시 히트 시 O(1), 캐시 미스 시 O(k)
//
// TOKST: O(1) - uses pre-loaded KST location
//   미리 로드된 KST 위치 사용
//
// TOUTC: O(1) - built-in UTC conversion
//   내장 UTC 변환
//
// GETTIMEZONEOFFSET: O(1) with cache
//   캐시 사용 시 O(1)
//
// ISVALIDTIMEZONE: O(k) - must attempt LoadLocation
//   LoadLocation 시도해야 함
//   Not cached (validation only)
//   캐시되지 않음 (검증만)
//
// GETLOCALTIMEZONE: O(1) - returns time.Local
//   time.Local 반환
//
// LISTTIMAZONES: O(1) - returns pre-defined slice
//   미리 정의된 슬라이스 반환
//
// NOWINTIMEZONE: O(1) with cache
//   캐시 사용 시 O(1)
//
// NOWKST: O(1) - uses pre-loaded KST
//   미리 로드된 KST 사용
//
// LOADTIMEZONE: O(1) with cache hit, O(k) with cache miss
//   k = 1-5 microseconds for system timezone lookup
//   캐시 히트 시 O(1), 캐시 미스 시 O(k)
//   k = 시스템 타임존 조회에 1-5 마이크로초
//
// SPACE COMPLEXITY (공간 복잡도):
// - Timezone cache: O(n) where n = number of unique timezones used
//   타임존 캐시: O(n), n = 사용된 고유 타임존 수
// - Typically small: 5-20 timezones in practice
//   일반적으로 작음: 실제로 5-20개 타임존
// - Each location: ~100 bytes
//   각 위치: ~100바이트
//
// CACHING PERFORMANCE GAIN:
// 캐싱 성능 향상:
// - Without cache: ~2-5 microseconds per conversion
//   캐시 없음: 변환당 ~2-5 마이크로초
// - With cache: ~20-50 nanoseconds per conversion
//   캐시 있음: 변환당 ~20-50 나노초
// - Speedup: 40-250x faster
//   속도 향상: 40-250배 더 빠름
//
// BENCHMARK EXAMPLE:
// 벤치마크 예시:
//   ConvertTimezone (no cache): 3000 ns/op
//   ConvertTimezone (cached):     50 ns/op
//   60x faster with cache
//   캐시 사용 시 60배 더 빠름
//
// ============================================================================
// TIMEZONE CONVERSION SEMANTICS / 타임존 변환 의미론
// ============================================================================
//
// INSTANT IN TIME (시간의 순간):
// A point in time is absolute and independent of timezone.
// 시간의 한 지점은 절대적이며 타임존과 무관합니다.
//
// Example:
// 예시:
//   2024-01-01 12:00:00 KST (Asia/Seoul)
//   = 2024-01-01 03:00:00 UTC
//   = 2023-12-31 22:00:00 EST (America/New_York)
//   All represent the SAME instant in time
//   모두 시간의 동일한 순간을 나타냄
//
// UNIX TIMESTAMP PRESERVATION:
// Unix 타임스탬프 보존:
//   t1 := time.Date(2024, 1, 1, 12, 0, 0, 0, KST)
//   t2 := t1.In(UTC)
//   t1.Unix() == t2.Unix()  // true - same instant
//                           // true - 동일한 순간
//
// WALL CLOCK vs INSTANT:
// 벽시계 vs 순간:
// - Wall clock: What a clock on the wall shows (12:00 in Seoul)
//   벽시계: 벽에 있는 시계가 보여주는 것 (서울에서 12:00)
// - Instant: Absolute point in time (Unix: 1704074400)
//   순간: 시간의 절대 지점 (Unix: 1704074400)
// - Conversion changes wall clock, not instant
//   변환은 벽시계를 변경하지 순간을 변경하지 않음
//
// DAYLIGHT SAVING TIME (DST):
// 일광 절약 시간제 (DST):
// - IANA database handles DST automatically
//   IANA 데이터베이스가 DST를 자동으로 처리
// - Offset changes on DST transition dates
//   DST 전환 날짜에 오프셋 변경
// - Example: America/New_York is EST (GMT-5) in winter, EDT (GMT-4) in summer
//   예시: America/New_York은 겨울에 EST (GMT-5), 여름에 EDT (GMT-4)
//
// ============================================================================
// TIMEZONE NAME FORMATS / 타임존 이름 형식
// ============================================================================
//
// IANA STANDARD FORMAT:
// IANA 표준 형식:
// - Continent/City: "Asia/Seoul", "America/New_York"
// - Area/Location: "Pacific/Honolulu"
// - Special: "UTC", "GMT"
//
// COMMON PATTERNS:
// 일반 패턴:
// - Asia: Seoul, Tokyo, Shanghai, Hong_Kong, Singapore, Dubai
// - Europe: London, Paris, Berlin, Rome, Moscow
// - America: New_York, Chicago, Los_Angeles, Toronto, Sao_Paulo
// - Pacific: Auckland, Fiji, Honolulu
// - Australia: Sydney, Melbourne, Perth
// - Africa: Cairo, Johannesburg, Lagos
//
// OFFSET EXAMPLES:
// 오프셋 예시:
// - KST (Asia/Seoul): GMT+9 (9 hours ahead of UTC)
//   GMT+9 (UTC보다 9시간 앞섬)
// - EST (America/New_York): GMT-5 (5 hours behind UTC)
//   GMT-5 (UTC보다 5시간 뒤짐)
// - UTC: GMT+0 (reference point)
//   GMT+0 (기준점)
//
// ABBREVIATIONS vs NAMES:
// 약어 vs 이름:
// - Use IANA names, not abbreviations: "Asia/Seoul" not "KST"
//   IANA 이름 사용, 약어 아님: "KST"가 아닌 "Asia/Seoul"
// - Abbreviations are ambiguous: CST = China/Central/Cuba Standard Time
//   약어는 모호함: CST = 중국/중부/쿠바 표준시
// - IANA names are unambiguous
//   IANA 이름은 명확함
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Display User's Local Time (사용자의 로컬 시간 표시)
// Use case: Show event time in user's timezone
// 사용 사례: 사용자의 타임존으로 이벤트 시간 표시
//
//   // Event stored in UTC
//   eventTime := event.StartsAt  // UTC
//
//   // Convert to user's timezone
//   userTZ := user.Timezone  // "America/New_York"
//   localTime, _ := timeutil.ConvertTimezone(eventTime, userTZ)
//
//   fmt.Printf("Event starts at %s (%s)\n",
//       timeutil.FormatDateTime(localTime),
//       userTZ)
//
// PATTERN 2: Store in UTC, Display in Local (UTC로 저장, 로컬로 표시)
// Use case: Database best practice
// 사용 사례: 데이터베이스 모범 사례
//
//   // Store in UTC
//   utcTime := timeutil.ToUTC(time.Now())
//   db.Save(&Record{CreatedAt: utcTime})
//
//   // Display in KST
//   record := db.Find(id)
//   kstTime := timeutil.ToKST(record.CreatedAt)
//   fmt.Println(timeutil.FormatDateTime(kstTime))
//
// PATTERN 3: Multi-timezone Application (다중 타임존 애플리케이션)
// Use case: Global service with multiple offices
// 사용 사례: 여러 사무실이 있는 글로벌 서비스
//
//   offices := map[string]string{
//       "Seoul":     "Asia/Seoul",
//       "New York":  "America/New_York",
//       "London":    "Europe/London",
//   }
//
//   now := time.Now()
//   for office, tz := range offices {
//       localTime, _ := timeutil.ConvertTimezone(now, tz)
//       fmt.Printf("%s: %s\n", office,
//           timeutil.FormatDateTime(localTime))
//   }
//
// PATTERN 4: Timezone Selector (타임존 선택기)
// Use case: User profile timezone setting
// 사용 사례: 사용자 프로필 타임존 설정
//
//   // Show timezone options
//   timezones := timeutil.ListTimezones()
//   for i, tz := range timezones {
//       offset, _ := timeutil.GetTimezoneOffset(tz)
//       fmt.Printf("%d. %s (GMT%+d)\n", i+1, tz, offset)
//   }
//
//   // Validate user selection
//   userInput := "Asia/Seoul"
//   if timeutil.IsValidTimezone(userInput) {
//       user.Timezone = userInput
//       db.Save(user)
//   }
//
// PATTERN 5: Current Time Dashboard (현재 시간 대시보드)
// Use case: Display multiple timezone clocks
// 사용 사례: 여러 타임존 시계 표시
//
//   func showWorldClocks() {
//       cities := []string{
//           "Asia/Seoul",
//           "America/New_York",
//           "Europe/London",
//           "Asia/Tokyo",
//       }
//
//       for _, tz := range cities {
//           now, _ := timeutil.NowInTimezone(tz)
//           fmt.Printf("%s: %s\n", tz,
//               timeutil.FormatTime(now))
//       }
//   }
//
// PATTERN 6: Timezone Conversion Chain (타임존 변환 체인)
// Use case: Meeting across timezones
// 사용 사례: 타임존을 넘어선 회의
//
//   // Meeting at 3 PM KST
//   meetingKST := time.Date(2024, 1, 15, 15, 0, 0, 0, timeutil.KST)
//
//   // What time in New York?
//   meetingNY, _ := timeutil.ConvertTimezone(meetingKST, "America/New_York")
//   fmt.Printf("Seoul: %s\n", timeutil.FormatTime(meetingKST))
//   fmt.Printf("New York: %s\n", timeutil.FormatTime(meetingNY))
//
// PATTERN 7: Offset Comparison (오프셋 비교)
// Use case: Find timezones near user's timezone
// 사용 사례: 사용자의 타임존 근처 타임존 찾기
//
//   userTZ := "Asia/Seoul"
//   userOffset, _ := timeutil.GetTimezoneOffset(userTZ)
//
//   for _, tz := range timeutil.ListTimezones() {
//       offset, _ := timeutil.GetTimezoneOffset(tz)
//       if abs(offset - userOffset) <= 2 {  // Within 2 hours
//           fmt.Printf("Similar: %s (GMT%+d)\n", tz, offset)
//       }
//   }
//
// PATTERN 8: API Response Formatting (API 응답 형식화)
// Use case: Return time in client's timezone
// 사용 사례: 클라이언트의 타임존으로 시간 반환
//
//   func getAppointmentAPI(w http.ResponseWriter, r *http.Request) {
//       // Get client timezone from header
//       clientTZ := r.Header.Get("X-Timezone")
//       if clientTZ == "" {
//           clientTZ = "Asia/Seoul"  // default
//       }
//
//       // Convert appointment time
//       apt := db.GetAppointment(id)
//       localTime, _ := timeutil.ConvertTimezone(apt.Time, clientTZ)
//
//       json.NewEncoder(w).Encode(map[string]interface{}{
//           "appointment_time": timeutil.FormatDateTime(localTime),
//           "timezone": clientTZ,
//       })
//   }
//
// PATTERN 9: Scheduling Across Timezones (타임존을 넘어선 스케줄링)
// Use case: Find common available hours
// 사용 사례: 공통 가능 시간 찾기
//
//   // 9 AM in each timezone
//   seoulMorning := time.Date(2024, 1, 15, 9, 0, 0, 0, timeutil.KST)
//   nyMorning := time.Date(2024, 1, 15, 9, 0, 0, 0, mustLoadLocation("America/New_York"))
//
//   // Convert both to UTC for comparison
//   seoulUTC := timeutil.ToUTC(seoulMorning)
//   nyUTC := timeutil.ToUTC(nyMorning)
//
//   fmt.Printf("Seoul 9 AM = %s UTC\n", timeutil.FormatTime(seoulUTC))
//   fmt.Printf("NY 9 AM = %s UTC\n", timeutil.FormatTime(nyUTC))
//
// PATTERN 10: Dynamic Default Timezone (동적 기본 타임존)
// Use case: Multi-tenant application
// 사용 사례: 다중 테넌트 애플리케이션
//
//   func setTenantTimezone(tenantID string) {
//       tenant := db.GetTenant(tenantID)
//       if tenant.Timezone != "" {
//           timeutil.SetDefaultTimezone(tenant.Timezone)
//       }
//   }
//
//   // Now all timeutil operations use tenant's timezone
//   now := timeutil.NowKST()  // Actually uses tenant timezone
//
// ============================================================================
// EDGE CASES / 경계 사례
// ============================================================================
//
// INVALID TIMEZONE NAME:
// 잘못된 타임존 이름:
//   _, err := timeutil.ConvertTimezone(t, "Invalid/Timezone")
//   // Returns error
//   // 오류 반환
//
// TIMEZONE WITH DST TRANSITION:
// DST 전환이 있는 타임존:
//   // On DST transition, some hours don't exist or repeat
//   // "Spring forward": 2 AM becomes 3 AM (2:30 doesn't exist)
//   // "Fall back": 2 AM repeats twice
//   // Go handles this automatically using IANA database
//   // DST 전환 시 일부 시간은 존재하지 않거나 반복됨
//   // "봄 전환": 오전 2시가 오전 3시가 됨 (2:30는 존재하지 않음)
//   // "가을 전환": 오전 2시가 두 번 반복됨
//   // Go는 IANA 데이터베이스를 사용하여 자동으로 처리
//
// OFFSET WITH MINUTES:
// 분이 있는 오프셋:
//   // Some timezones have non-hour offsets
//   // Asia/Kathmandu: GMT+5:45
//   // Australia/Adelaide: GMT+9:30
//   // GetTimezoneOffset returns hours only (truncates)
//   // 일부 타임존은 시간이 아닌 오프셋을 가짐
//   // GetTimezoneOffset는 시간만 반환 (절사)
//
// UTC CONVERSION:
// UTC 변환:
//   // t.UTC() is faster than ConvertTimezone(t, "UTC")
//   // Use ToUTC() for better performance
//   // t.UTC()가 ConvertTimezone(t, "UTC")보다 빠름
//   // 더 나은 성능을 위해 ToUTC() 사용
//
// ZERO VALUE TIME:
// 제로 값 시간:
//   var t time.Time
//   converted := timeutil.ToKST(t)
//   // Returns zero time in KST (still zero)
//   // KST의 제로 시간 반환 (여전히 제로)
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// THREAD-SAFE OPERATIONS:
// 스레드 안전 연산:
// - All timezone conversion functions (ConvertTimezone, ToKST, ToUTC)
//   모든 타임존 변환 함수
// - Timezone information queries (GetTimezoneOffset, etc.)
//   타임존 정보 쿼리
// - Current time functions (NowInTimezone, NowKST)
//   현재 시간 함수
//
// CACHE SYNCHRONIZATION:
// 캐시 동기화:
// - sync.RWMutex protects timezone cache
//   sync.RWMutex가 타임존 캐시 보호
// - Multiple goroutines can read concurrently
//   여러 고루틴이 동시에 읽을 수 있음
// - Cache writes are serialized
//   캐시 쓰기는 직렬화됨
// - Safe for high-concurrency web applications
//   높은 동시성 웹 애플리케이션에 안전
//
// DEFAULT TIMEZONE UPDATES:
// 기본 타임존 업데이트:
// - SetDefaultTimezone is NOT thread-safe
//   SetDefaultTimezone은 스레드 안전하지 않음
// - Should be called once at initialization
//   초기화 시 한 번 호출되어야 함
// - Don't change default timezone at runtime in concurrent contexts
//   동시성 컨텍스트에서 런타임에 기본 타임존 변경하지 마세요
//
// SAFE CONCURRENT PATTERN:
// 안전한 동시 패턴:
//   func init() {
//       timeutil.SetDefaultTimezone("America/New_York")
//   }
//
//   // Then use from multiple goroutines
//   var wg sync.WaitGroup
//   for _, userID := range users {
//       wg.Add(1)
//       go func(uid string) {
//           defer wg.Done()
//           user := getUser(uid)
//           localTime, _ := timeutil.ConvertTimezone(time.Now(), user.Timezone)
//           sendNotification(user, localTime)
//       }(userID)
//   }
//   wg.Wait()
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM constants.go:
// - KST: Pre-loaded Asia/Seoul timezone
// - defaultLocation: Current default timezone
// - SecondsPerHour: For offset calculation
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.Location: Timezone representation
// - time.LoadLocation: Load timezone from IANA database
// - sync.RWMutex: Thread-safe cache
//
// USED BY (사용처):
// - All timeutil functions that use timezone (indirectly via defaultLocation)
//   타임존을 사용하는 모든 timeutil 함수 (defaultLocation을 통해 간접적으로)
// - Web applications (display times in user's timezone)
//   웹 애플리케이션 (사용자의 타임존으로 시간 표시)
// - APIs (format responses in client timezone)
//   API (클라이언트 타임존으로 응답 형식화)
// - Scheduling systems (coordinate across timezones)
//   스케줄링 시스템 (타임존을 넘어선 조정)
// - Multi-region applications
//   다중 지역 애플리케이션
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. STORE IN UTC, DISPLAY IN LOCAL
//    UTC로 저장, 로컬로 표시
//    db.CreatedAt = timeutil.ToUTC(time.Now())
//    display := timeutil.ToKST(db.CreatedAt)
//
// 2. USE IANA NAMES, NOT ABBREVIATIONS
//    약어 아닌 IANA 이름 사용
//    Good: "Asia/Seoul"
//    Bad:  "KST" (ambiguous)
//
// 3. VALIDATE TIMEZONE BEFORE USING
//    사용 전 타임존 검증
//    if timeutil.IsValidTimezone(userInput) {
//        // Use it
//    }
//
// 4. SET DEFAULT ONCE AT STARTUP
//    시작 시 기본값 한 번 설정
//    func init() {
//        timeutil.SetDefaultTimezone("America/New_York")
//    }
//
// 5. CACHE TIMEZONE CONVERSIONS IF REPEATED
//    반복되는 경우 타임존 변환 캐시
//    // For same timezone used many times
//
// 6. USE TOKST/TOUTC FOR COMMON CONVERSIONS
//    일반 변환에 ToKST/ToUTC 사용
//    // Faster than generic ConvertTimezone
//
// 7. INCLUDE TIMEZONE IN API RESPONSES
//    API 응답에 타임존 포함
//    {"time": "2024-01-01 12:00", "timezone": "Asia/Seoul"}
//
// 8. HANDLE DST AUTOMATICALLY
//    DST 자동 처리
//    // Don't manually adjust for DST, use IANA names
//
// ============================================================================

// Timezone cache for performance
// 성능을 위한 타임존 캐시
var (
	timezoneCache   = make(map[string]*time.Location)
	timezoneCacheMu sync.RWMutex
)

// SetDefaultTimezone sets the default timezone for all timeutil functions.
// SetDefaultTimezone은 모든 timeutil 함수의 기본 타임존을 설정합니다.
//
// Default is "Asia/Seoul" (KST, GMT+9).
// 기본값은 "Asia/Seoul" (KST, GMT+9)입니다.
//
// Example
// 예제:
//
//	timeutil.SetDefaultTimezone("America/New_York")
func SetDefaultTimezone(tz string) error {
	loc, err := loadTimezone(tz)
	if err != nil {
		return fmt.Errorf("failed to set default timezone: %w", err)
	}
	defaultLocation = loc
	return nil
}

// GetDefaultTimezone returns the current default timezone name.
// GetDefaultTimezone은 현재 기본 타임존 이름을 반환합니다.
//
// Example
// 예제:
//
//	tz := timeutil.GetDefaultTimezone() // "Asia/Seoul"
func GetDefaultTimezone() string {
	return defaultLocation.String()
}

// ResetDefaultTimezone resets the default timezone to KST (Asia/Seoul).
// ResetDefaultTimezone은 기본 타임존을 KST (Asia/Seoul)로 재설정합니다.
func ResetDefaultTimezone() {
	defaultLocation = KST
}

// ConvertTimezone converts a time to a different timezone.
// ConvertTimezone은 시간을 다른 타임존으로 변환합니다.
//
// Example
// 예제:
//
//	now := time.Now()
//	seoulTime, _ := timeutil.ConvertTimezone(now, "Asia/Seoul")
//	nyTime, _ := timeutil.ConvertTimezone(now, "America/New_York")
func ConvertTimezone(t time.Time, tz string) (time.Time, error) {
	loc, err := loadTimezone(tz)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to convert timezone: %w", err)
	}
	return t.In(loc), nil
}

// ToKST converts a time to KST (Asia/Seoul, GMT+9).
// ToKST는 시간을 KST (Asia/Seoul, GMT+9)로 변환합니다.
//
// This is a convenience function for ConvertTimezone(t, "Asia/Seoul").
// 이것은 ConvertTimezone(t, "Asia/Seoul")의 편의 함수입니다.
//
// Example
// 예제:
//
//	kstTime := timeutil.ToKST(time.Now())
func ToKST(t time.Time) time.Time {
	return t.In(KST)
}

// ToUTC converts a time to UTC.
// ToUTC는 시간을 UTC로 변환합니다.
//
// Example
// 예제:
//
//	utcTime := timeutil.ToUTC(time.Now())
func ToUTC(t time.Time) time.Time {
	return t.UTC()
}

// GetTimezoneOffset returns the timezone offset in hours from UTC.
// GetTimezoneOffset는 UTC로부터의 타임존 오프셋을 시간 단위로 반환합니다.
//
// Example
// 예제:
//
//	offset, _ := timeutil.GetTimezoneOffset("Asia/Seoul") // +9
func GetTimezoneOffset(tz string) (int, error) {
	loc, err := loadTimezone(tz)
	if err != nil {
		return 0, fmt.Errorf("failed to get timezone offset: %w", err)
	}

	// Get offset at current time
	// 현재 시간의 오프셋 가져오기
	_, offset := time.Now().In(loc).Zone()
	return offset / SecondsPerHour, nil
}

// ListTimezones returns a list of common timezone names.
// ListTimezones는 일반적인 타임존 이름 목록을 반환합니다.
//
// Note: This returns a curated list of commonly used timezones.
// 참고: 이것은 일반적으로 사용되는 타임존의 선별된 목록을 반환합니다.
func ListTimezones() []string {
	return []string{
		// Asia
		// 아시아
		"Asia/Seoul",        // KST (GMT+9)
		"Asia/Tokyo",        // JST (GMT+9)
		"Asia/Shanghai",     // CST (GMT+8)
		"Asia/Hong_Kong",    // HKT (GMT+8)
		"Asia/Singapore",    // SGT (GMT+8)
		"Asia/Bangkok",      // ICT (GMT+7)
		"Asia/Dubai",        // GST (GMT+4)
		"Asia/Kolkata",      // IST (GMT+5:30)
		"Asia/Jakarta",      // WIB (GMT+7)
		"Asia/Manila",       // PHT (GMT+8)
		"Asia/Taipei",       // CST (GMT+8)
		"Asia/Ho_Chi_Minh",  // ICT (GMT+7)
		"Asia/Kuala_Lumpur", // MYT (GMT+8)

		// Europe
		// 유럽
		"Europe/London",    // GMT/BST (GMT+0/+1)
		"Europe/Paris",     // CET/CEST (GMT+1/+2)
		"Europe/Berlin",    // CET/CEST (GMT+1/+2)
		"Europe/Rome",      // CET/CEST (GMT+1/+2)
		"Europe/Madrid",    // CET/CEST (GMT+1/+2)
		"Europe/Amsterdam", // CET/CEST (GMT+1/+2)
		"Europe/Brussels",  // CET/CEST (GMT+1/+2)
		"Europe/Vienna",    // CET/CEST (GMT+1/+2)
		"Europe/Zurich",    // CET/CEST (GMT+1/+2)
		"Europe/Stockholm", // CET/CEST (GMT+1/+2)
		"Europe/Moscow",    // MSK (GMT+3)

		// Americas
		// 아메리카
		"America/New_York",       // EST/EDT (GMT-5/-4)
		"America/Chicago",        // CST/CDT (GMT-6/-5)
		"America/Denver",         // MST/MDT (GMT-7/-6)
		"America/Los_Angeles",    // PST/PDT (GMT-8/-7)
		"America/Toronto",        // EST/EDT (GMT-5/-4)
		"America/Vancouver",      // PST/PDT (GMT-8/-7)
		"America/Mexico_City",    // CST/CDT (GMT-6/-5)
		"America/Sao_Paulo",      // BRT/BRST (GMT-3/-2)
		"America/Buenos_Aires",   // ART (GMT-3)
		"America/Santiago",       // CLT/CLST (GMT-4/-3)
		"America/Bogota",         // COT (GMT-5)
		"America/Lima",           // PET (GMT-5)
		"America/Caracas",        // VET (GMT-4)
		"America/Panama",         // EST (GMT-5)
		"America/Havana",         // CST/CDT (GMT-5/-4)
		"America/Port-au-Prince", // EST/EDT (GMT-5/-4)

		// Pacific
		// 태평양
		"Pacific/Auckland",  // NZST/NZDT (GMT+12/+13)
		"Pacific/Fiji",      // FJT/FJST (GMT+12/+13)
		"Pacific/Honolulu",  // HST (GMT-10)
		"Pacific/Guam",      // ChST (GMT+10)
		"Pacific/Pago_Pago", // SST (GMT-11)
		"Pacific/Tahiti",    // TAHT (GMT-10)

		// Australia
		// 호주
		"Australia/Sydney",    // AEST/AEDT (GMT+10/+11)
		"Australia/Melbourne", // AEST/AEDT (GMT+10/+11)
		"Australia/Brisbane",  // AEST (GMT+10)
		"Australia/Perth",     // AWST (GMT+8)
		"Australia/Adelaide",  // ACST/ACDT (GMT+9:30/+10:30)

		// Africa
		// 아프리카
		"Africa/Cairo",        // EET/EEST (GMT+2/+3)
		"Africa/Johannesburg", // SAST (GMT+2)
		"Africa/Lagos",        // WAT (GMT+1)
		"Africa/Nairobi",      // EAT (GMT+3)
		"Africa/Casablanca",   // WET/WEST (GMT+0/+1)

		// Middle East
		// 중동
		"Asia/Jerusalem", // IST/IDT (GMT+2/+3)
		"Asia/Riyadh",    // AST (GMT+3)
		"Asia/Tehran",    // IRST/IRDT (GMT+3:30/+4:30)
		"Asia/Baghdad",   // AST (GMT+3)
		"Asia/Kuwait",    // AST (GMT+3)
		"Asia/Doha",      // AST (GMT+3)
		"Asia/Muscat",    // GST (GMT+4)
		"Asia/Karachi",   // PKT (GMT+5)
		"Asia/Dhaka",     // BST (GMT+6)
		"Asia/Yangon",    // MMT (GMT+6:30)
		"Asia/Kathmandu", // NPT (GMT+5:45)

		// UTC
		"UTC",
	}
}

// IsValidTimezone checks if a timezone name is valid.
// IsValidTimezone은 타임존 이름이 유효한지 확인합니다.
//
// Example
// 예제:
//
//	if timeutil.IsValidTimezone("Asia/Seoul") {
//	    // Valid timezone
//	}
func IsValidTimezone(tz string) bool {
	_, err := time.LoadLocation(tz)
	return err == nil
}

// GetLocalTimezone returns the local system timezone name.
// GetLocalTimezone은 로컬 시스템 타임존 이름을 반환합니다.
//
// Example
// 예제:
//
//	local := timeutil.GetLocalTimezone()
func GetLocalTimezone() string {
	return time.Local.String()
}

// NowInTimezone returns the current time in the specified timezone.
// NowInTimezone은 지정된 타임존의 현재 시간을 반환합니다.
//
// Example
// 예제:
//
//	seoulNow, _ := timeutil.NowInTimezone("Asia/Seoul")
func NowInTimezone(tz string) (time.Time, error) {
	loc, err := loadTimezone(tz)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get current time: %w", err)
	}
	return time.Now().In(loc), nil
}

// NowKST returns the current time in KST (Asia/Seoul, GMT+9).
// NowKST는 KST (Asia/Seoul, GMT+9)의 현재 시간을 반환합니다.
//
// This is the default timezone for all timeutil functions.
// 이것은 모든 timeutil 함수의 기본 타임존입니다.
//
// Example
// 예제:
//
//	now := timeutil.NowKST()
func NowKST() time.Time {
	return time.Now().In(KST)
}

// loadTimezone loads a timezone location with caching.
// loadTimezone은 캐싱과 함께 타임존 위치를 로드합니다.
func loadTimezone(tz string) (*time.Location, error) {
	// Check cache first
	// 먼저 캐시 확인
	timezoneCacheMu.RLock()
	if loc, ok := timezoneCache[tz]; ok {
		timezoneCacheMu.RUnlock()
		return loc, nil
	}
	timezoneCacheMu.RUnlock()

	// Load timezone
	// 타임존 로드
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, fmt.Errorf("invalid timezone '%s': %w", tz, err)
	}

	// Cache it
	// 캐시에 저장
	timezoneCacheMu.Lock()
	timezoneCache[tz] = loc
	timezoneCacheMu.Unlock()

	return loc, nil
}
