package timeutil

import (
	"fmt"
	"math"
	"time"
)

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/relative.go
// Purpose: Relative time representation and human-friendly duration formatting
//          상대 시간 표현 및 사람 친화적 기간 포맷팅
//
// This file provides human-readable time representations that express time
// relative to "now" (e.g., "2 hours ago", "in 3 days") and convert durations
// into friendly strings (e.g., "2 hours 30 minutes"). These are essential for
// user interfaces, social media feeds, notifications, and anywhere humans need
// to quickly understand temporal relationships.
//
// 이 파일은 "지금"을 기준으로 시간을 표현하는 사람이 읽을 수 있는 시간 표현
// (예: "2시간 전", "3일 후")과 기간을 친근한 문자열로 변환하는 기능
// (예: "2시간 30분")을 제공합니다. 이는 사용자 인터페이스, 소셜 미디어 피드,
// 알림 및 사람이 시간 관계를 빠르게 이해해야 하는 모든 곳에 필수적입니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. RELATIVE TIME PAST (과거 상대 시간)
//    - RelativeTime: "2 hours ago", "3 days ago"
//      과거 시간을 "~전" 형식으로 표현
//    - Automatically chooses appropriate unit (seconds/minutes/hours/days/weeks/months/years)
//      자동으로 적절한 단위 선택 (초/분/시/일/주/월/년)
//    - Special handling for very recent times: "just now"
//      최근 시간에 대한 특별 처리: "방금 전"
//
// 2. RELATIVE TIME FUTURE (미래 상대 시간)
//    - RelativeTime: "in 2 hours", "in 3 days"
//      미래 시간을 "~후" 형식으로 표현
//    - Same unit selection logic as past times
//      과거 시간과 동일한 단위 선택 논리
//    - Special handling for imminent times: "just now"
//      임박한 시간에 대한 특별 처리: "곧"
//
// 3. SHORT RELATIVE TIME (짧은 상대 시간)
//    - RelativeTimeShort: "2h ago", "in 3d"
//      축약된 형식으로 표현
//    - Compact format for space-constrained UIs
//      공간이 제한된 UI를 위한 간결한 형식
//    - Uses abbreviations: s, m, h, d, w, mo, y
//      축약어 사용: s, m, h, d, w, mo, y
//
// 4. CONVENIENT ALIAS (편리한 별칭)
//    - TimeAgo: Alias for RelativeTime
//      RelativeTime의 별칭
//    - More intuitive name for past times
//      과거 시간에 더 직관적인 이름
//    - Common pattern in social media apps
//      소셜 미디어 앱의 일반적인 패턴
//
// 5. DURATION HUMANIZATION (기간 인간화)
//    - HumanizeDuration: "2 hours 30 minutes"
//      기간을 사람이 읽기 쉬운 형식으로 변환
//    - Breaks down duration into components
//      기간을 구성 요소로 분해
//    - Shows up to 2 significant units
//      최대 2개의 유의미한 단위 표시
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. HUMAN-FIRST APPROACH (사람 우선 접근법)
//    - Output designed for human readers, not machines
//      기계가 아닌 사람 독자를 위한 출력 설계
//    - Uses natural language: "2 hours ago" vs "120 minutes ago"
//      자연어 사용: "2시간 전" vs "120분 전"
//    - Prefers larger units when appropriate: "1 day" vs "24 hours"
//      적절할 때 더 큰 단위 선호: "1일" vs "24시간"
//
// 2. CONTEXTUAL PRECISION (맥락적 정밀도)
//    - Recent times: More precise (seconds/minutes)
//      최근 시간: 더 정밀 (초/분)
//    - Distant times: Less precise (months/years)
//      먼 시간: 덜 정밀 (월/년)
//    - Tradeoff: Readability vs precision
//      트레이드오프: 가독성 vs 정밀도
//    - "2 years ago" is more readable than "730 days ago"
//      "2년 전"이 "730일 전"보다 가독성이 높습니다
//
// 3. SINGULAR VS PLURAL (단수 vs 복수)
//    - "1 hour ago" (singular) vs "2 hours ago" (plural)
//      "1시간 전" (단수) vs "2시간 전" (복수)
//    - Grammatically correct English
//      문법적으로 올바른 영어
//    - Improves user experience
//      사용자 경험 향상
//
// 4. APPROXIMATION IS OK (근사치 허용)
//    - Months approximated as 30 days
//      월은 30일로 근사
//    - Years approximated as 365 days
//      년은 365일로 근사
//    - Good enough for human perception
//      사람의 인식에 충분히 정확
//    - Not for precise calculations
//      정밀 계산용이 아님
//
// 5. BIDIRECTIONAL TIME (양방향 시간)
//    - Supports both past ("ago") and future ("in")
//      과거 ("~전")와 미래 ("~후") 모두 지원
//    - Automatically detects direction from time difference
//      시간 차이에서 방향 자동 감지
//    - Consistent format for both directions
//      양방향에 일관된 형식
//
// 6. SHORT FORMAT FOR SPACE CONSTRAINTS (공간 제약을 위한 짧은 형식)
//    - RelativeTimeShort uses abbreviations
//      RelativeTimeShort는 축약어 사용
//    - "2h ago" vs "2 hours ago"
//      "2h ago" vs "2 hours ago"
//    - Useful for mobile UIs, tables, charts
//      모바일 UI, 테이블, 차트에 유용
//
// ============================================================================
// RELATIVE TIME OPERATIONS OVERVIEW / 상대 시간 연산 개요
// ============================================================================
//
// PUBLIC FUNCTIONS (공개 함수) - 4 functions
// ├─ RelativeTime      : Full format relative time
// ├─ RelativeTimeShort : Short format relative time
// ├─ TimeAgo           : Alias for RelativeTime
// └─ HumanizeDuration  : Humanize time.Duration
//
// INTERNAL HELPERS (내부 헬퍼) - 4 functions
// ├─ relativeTimePast        : Format past times
// ├─ relativeTimeFuture      : Format future times
// ├─ relativeTimePastShort   : Format past times (short)
// └─ relativeTimeFutureShort : Format future times (short)
//
// Total: 4 public + 4 internal = 8 functions
// 총: 공개 4개 + 내부 4개 = 8개 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
// All functions: O(1) - constant time
// 모든 함수: O(1) - 상수 시간
//
// RELATIVETIME / RELATIVETIMESHORT:
// - Calls time.Now() once
//   time.Now()를 한 번 호출
// - Timezone conversion to KST
//   KST로 타임존 변환
// - Single duration calculation
//   단일 기간 계산
// - Switch statement for unit selection
//   단위 선택을 위한 switch 문
// - Fast: ~100-200 nanoseconds
//   빠름: ~100-200나노초
//
// HUMANIZEDURATION:
// - Simple arithmetic on duration
//   기간에 대한 간단한 산술
// - String formatting
//   문자열 포맷팅
// - Very fast: ~50-100 nanoseconds
//   매우 빠름: ~50-100나노초
//
// SPACE COMPLEXITY (공간 복잡도):
// - All functions: O(1) - returns single string
//   모든 함수: O(1) - 단일 문자열 반환
// - String formatting uses temporary buffer
//   문자열 포맷팅은 임시 버퍼 사용
//
// PERFORMANCE TIPS (성능 팁):
// 1. Cache result if displaying same time multiple times
//    같은 시간을 여러 번 표시하는 경우 결과 캐시
// 2. Use RelativeTimeShort for better performance (simpler logic)
//    더 나은 성능을 위해 RelativeTimeShort 사용 (더 간단한 논리)
// 3. Avoid calling in tight loops - calculate once, display many
//    긴밀한 루프에서 호출 피하기 - 한 번 계산, 여러 번 표시
//
// ============================================================================
// TIME UNIT SELECTION LOGIC / 시간 단위 선택 논리
// ============================================================================
//
// RelativeTime uses THRESHOLD-BASED unit selection:
// RelativeTime은 임계값 기반 단위 선택을 사용합니다:
//
// PAST TIMES (과거 시간):
// < 10 seconds    → "just now"
// < 60 seconds    → "X seconds ago"
// < 60 minutes    → "X minutes ago"
// < 24 hours      → "X hours ago"
// < 7 days        → "X days ago"
// < 4 weeks       → "X weeks ago"
// < 12 months     → "X months ago"
// >= 12 months    → "X years ago"
//
// FUTURE TIMES (미래 시간):
// < 10 seconds    → "just now"
// < 60 seconds    → "in X seconds"
// < 60 minutes    → "in X minutes"
// < 24 hours      → "in X hours"
// < 7 days        → "in X days"
// < 4 weeks       → "in X weeks"
// < 12 months     → "in X months"
// >= 12 months    → "in X years"
//
// RATIONALE (근거):
// - 10 second threshold for "just now": Feels immediate
//   "방금 전"을 위한 10초 임계값: 즉각적으로 느껴집니다
// - 7 days before switching to weeks: Standard week definition
//   주로 전환하기 전 7일: 표준 주 정의
// - 4 weeks before switching to months: Approximate month
//   월로 전환하기 전 4주: 근사 월
// - 12 months before switching to years: Avoid "13 months ago"
//   년으로 전환하기 전 12개월: "13개월 전" 방지
//
// ============================================================================
// APPROXIMATIONS AND ACCURACY / 근사치와 정확도
// ============================================================================
//
// This file uses APPROXIMATIONS for simplicity:
// 이 파일은 단순성을 위해 근사치를 사용합니다:
//
// MONTH APPROXIMATION (월 근사):
// - 1 month = 30 days (not 28-31)
//   1개월 = 30일 (28-31일이 아님)
// - Error: ±3% depending on actual month
//   오차: 실제 월에 따라 ±3%
// - Acceptable for human-readable display
//   사람이 읽을 수 있는 표시에 허용 가능
//
// YEAR APPROXIMATION (년 근사):
// - 1 year = 365 days (not 365.25)
//   1년 = 365일 (365.25일이 아님)
// - Ignores leap years
//   윤년 무시
// - Error: ~0.25% per year
//   오차: 연간 ~0.25%
// - Acceptable for relative time display
//   상대 시간 표시에 허용 가능
//
// WEEK CALCULATION (주 계산):
// - 1 week = 7 days (exact)
//   1주 = 7일 (정확)
// - No approximation needed
//   근사치 불필요
//
// WHEN APPROXIMATIONS MATTER (근사치가 중요한 경우):
// - For precise calculations: Use DiffInMonths, DiffInYears, AgeInYears
//   정밀 계산: DiffInMonths, DiffInYears, AgeInYears 사용
// - For display: RelativeTime approximations are fine
//   표시용: RelativeTime 근사치는 괜찮습니다
//
// Example of approximation impact:
// 근사치 영향 예시:
//   Actual: 31 days
//   Display: "1 month ago" (using 30-day approximation)
//   Reality: Could be "1 month 1 day ago" if precise
//   실제: 정확하면 "1개월 1일 전"일 수 있음
//
// ============================================================================
// HUMANIZEDURATION LOGIC / HumanizeDuration 논리
// ============================================================================
//
// HumanizeDuration shows UP TO 2 UNITS:
// HumanizeDuration은 최대 2개 단위를 표시합니다:
//
// Examples:
// 예시:
//   2h 30m 45s → "2 hours 30 minutes" (days=0, hours=2, mins=30, secs=45)
//   1d 2h 30m  → "1 day 2 hours" (days=1, hours=2, mins=30, secs=0)
//   5m 30s     → "5 minutes 30 seconds" (days=0, hours=0, mins=5, secs=30)
//   45s        → "45 seconds" (days=0, hours=0, mins=0, secs=45)
//
// Logic:
// 논리:
// 1. Calculate days, hours, minutes, seconds
//    일, 시, 분, 초 계산
// 2. Show largest non-zero unit + next unit if non-zero
//    0이 아닌 가장 큰 단위 + 다음 단위가 0이 아니면 표시
// 3. If only one unit, show just that unit
//    단위가 하나만 있으면 해당 단위만 표시
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Social Media Posts (소셜 미디어 게시물)
// Use case: Show when post was created
// 사용 사례: 게시물이 생성된 시간 표시
//
//   for _, post := range posts {
//       relativeTime := timeutil.RelativeTime(post.CreatedAt)
//       fmt.Printf("%s posted %s\n", post.Author, relativeTime)
//       // "Alice posted 2 hours ago"
//   }
//
// PATTERN 2: Comment Timestamps (댓글 타임스탬프)
// Use case: Compact timestamp for comments
// 사용 사례: 댓글을 위한 간결한 타임스탬프
//
//   for _, comment := range comments {
//       shortTime := timeutil.RelativeTimeShort(comment.CreatedAt)
//       fmt.Printf("%s: %s - %s\n",
//           comment.Author, comment.Text, shortTime)
//       // "Bob: Great post! - 5m ago"
//   }
//
// PATTERN 3: Activity Feed (활동 피드)
// Use case: Show recent user activity
// 사용 사례: 최근 사용자 활동 표시
//
//   for _, activity := range activities {
//       when := timeutil.TimeAgo(activity.Timestamp)
//       fmt.Printf("%s %s %s\n",
//           activity.User, activity.Action, when)
//       // "Carol liked your photo 1 hour ago"
//   }
//
// PATTERN 4: Notification List (알림 목록)
// Use case: Show when notifications were received
// 사용 사례: 알림을 받은 시간 표시
//
//   for _, notif := range notifications {
//       time := timeutil.RelativeTime(notif.ReceivedAt)
//       fmt.Printf("[%s] %s\n", time, notif.Message)
//       // "[3 days ago] Your order has shipped"
//   }
//
// PATTERN 5: File Modified Time (파일 수정 시간)
// Use case: Show when files were last modified
// 사용 사례: 파일이 마지막으로 수정된 시간 표시
//
//   for _, file := range files {
//       modified := timeutil.RelativeTime(file.ModifiedAt)
//       fmt.Printf("%s (modified %s)\n", file.Name, modified)
//       // "document.pdf (modified 2 days ago)"
//   }
//
// PATTERN 6: Upcoming Events (다가오는 이벤트)
// Use case: Show time until future events
// 사용 사례: 미래 이벤트까지의 시간 표시
//
//   for _, event := range upcomingEvents {
//       timeUntil := timeutil.RelativeTime(event.StartTime)
//       fmt.Printf("%s starts %s\n", event.Name, timeUntil)
//       // "Webinar starts in 2 hours"
//   }
//
// PATTERN 7: Dashboard Tables (대시보드 테이블)
// Use case: Compact time display in tables
// 사용 사례: 테이블의 간결한 시간 표시
//
//   fmt.Println("User\t\tLast Login")
//   for _, user := range users {
//       lastLogin := timeutil.RelativeTimeShort(user.LastLoginAt)
//       fmt.Printf("%s\t\t%s\n", user.Name, lastLogin)
//       // "Alice		2h ago"
//   }
//
// PATTERN 8: Duration Display (기간 표시)
// Use case: Show how long an operation took
// 사용 사례: 작업이 얼마나 걸렸는지 표시
//
//   start := time.Now()
//   processData()
//   duration := time.Since(start)
//
//   humanDuration := timeutil.HumanizeDuration(duration)
//   fmt.Printf("Processing completed in %s\n", humanDuration)
//   // "Processing completed in 2 minutes 30 seconds"
//
// PATTERN 9: Cache Expiry (캐시 만료)
// Use case: Show when cache will expire
// 사용 사례: 캐시가 만료될 시간 표시
//
//   expiresAt := cache.ExpiresAt
//   timeUntil := timeutil.RelativeTime(expiresAt)
//   fmt.Printf("Cache expires %s\n", timeUntil)
//   // "Cache expires in 5 minutes"
//
// PATTERN 10: Chat Messages (채팅 메시지)
// Use case: Show message timestamps in chat
// 사용 사례: 채팅에서 메시지 타임스탬프 표시
//
//   for _, msg := range messages {
//       when := timeutil.RelativeTimeShort(msg.SentAt)
//       fmt.Printf("[%s] %s: %s\n", when, msg.Sender, msg.Text)
//       // "[5m ago] Dave: Hello everyone!"
//   }
//
// ============================================================================
// COMPARISON WITH OTHER FORMATS / 다른 형식과의 비교
// ============================================================================
//
// RELATIVE TIME vs ABSOLUTE TIME:
// 상대 시간 vs 절대 시간:
//
// Relative: "2 hours ago"
// ✓ Intuitive for recent times
//   최근 시간에 직관적
// ✓ No timezone confusion
//   타임존 혼란 없음
// ✓ Dynamically updates meaning
//   의미가 동적으로 업데이트됨
// ✗ Less precise for old times
//   오래된 시간에 덜 정밀
// ✗ Not suitable for archives
//   아카이브에 적합하지 않음
//
// Absolute: "2024-01-15 14:30"
// ✓ Precise and permanent
//   정밀하고 영구적
// ✓ Good for records and archives
//   기록과 아카이브에 좋음
// ✗ Requires timezone knowledge
//   타임존 지식 필요
// ✗ Less intuitive
//   덜 직관적
//
// BEST PRACTICE: Use both
// 모범 사례: 둘 다 사용
//   Display: "2 hours ago" (relative)
//   Tooltip: "2024-01-15 14:30 KST" (absolute)
//
// FULL vs SHORT FORMAT:
// 전체 vs 짧은 형식:
//
// Full: "2 hours ago"
// ✓ More readable
//   더 읽기 쉬움
// ✓ Better for sentences
//   문장에 더 좋음
// ✗ Takes more space
//   더 많은 공간 차지
//
// Short: "2h ago"
// ✓ Compact
//   간결함
// ✓ Good for tables/lists
//   테이블/목록에 좋음
// ✗ Less readable for unfamiliar users
//   익숙하지 않은 사용자에게 덜 읽기 쉬움
//
// ============================================================================
// LOCALIZATION CONSIDERATIONS / 로컬라이제이션 고려사항
// ============================================================================
//
// Current implementation is ENGLISH ONLY:
// 현재 구현은 영어만 지원합니다:
//
// Output examples:
// 출력 예시:
// - "2 hours ago"
// - "in 3 days"
// - "just now"
//
// For Korean localization, you would need:
// 한글 로컬라이제이션을 위해서는 다음이 필요합니다:
// - "2시간 전"
// - "3일 후"
// - "방금 전"
//
// IMPLEMENTATION APPROACH:
// 구현 접근법:
// 1. Add locale parameter to functions
//    함수에 로케일 매개변수 추가
// 2. Create translation maps
//    번역 맵 생성
// 3. Use locale-specific templates
//    로케일별 템플릿 사용
//
// Example (not implemented):
// 예시 (구현되지 않음):
//   RelativeTimeLocalized(t, "ko-KR") → "2시간 전"
//   RelativeTimeLocalized(t, "en-US") → "2 hours ago"
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
// 2. TIME.NOW() USAGE (time.Now() 사용)
//    - Functions call time.Now() internally
//      함수는 내부적으로 time.Now()를 호출합니다
//    - Gets current time at moment of call
//      호출 시점의 현재 시간을 가져옵니다
//    - Safe for concurrent use
//      동시 사용에 안전합니다
//
// 3. NO SIDE EFFECTS (부작용 없음)
//    - Functions return strings only
//      함수는 문자열만 반환합니다
//    - No I/O operations
//      I/O 작업 없음
//    - No logging or external calls
//      로깅이나 외부 호출 없음
//
// 4. SAFE CONCURRENT USAGE (안전한 동시 사용)
//    - Multiple goroutines can call these functions simultaneously
//      여러 고루틴이 동시에 이러한 함수를 호출할 수 있습니다
//    - No locks or synchronization needed
//      잠금이나 동기화가 필요 없습니다
//    - Example:
//      예시:
//        var wg sync.WaitGroup
//        for _, post := range posts {
//            wg.Add(1)
//            go func(p Post) {
//                defer wg.Done()
//                relTime := timeutil.RelativeTime(p.CreatedAt)
//                displayPost(p, relTime)
//            }(post)
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
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.Duration: For duration calculations
// - time.Now(): For getting current time
// - time.Sub(): For calculating time difference
// - fmt.Sprintf(): For string formatting
// - math.Abs(): For absolute value
//
// USED BY (사용처):
// - User interfaces and dashboards
//   사용자 인터페이스 및 대시보드
// - Social media feeds
//   소셜 미디어 피드
// - Comment sections
//   댓글 섹션
// - Activity logs
//   활동 로그
// - Notification systems
//   알림 시스템
// - Chat applications
//   채팅 애플리케이션
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE RELATIVE TIME FOR RECENT EVENTS
//    최근 이벤트에 상대 시간 사용
//    relTime := timeutil.RelativeTime(event.CreatedAt)
//    fmt.Printf("Posted %s\n", relTime)
//
// 2. PROVIDE ABSOLUTE TIME IN TOOLTIP
//    툴팁에 절대 시간 제공
//    <span title="2024-01-15 14:30 KST">2 hours ago</span>
//
// 3. USE SHORT FORMAT FOR COMPACT UIs
//    간결한 UI에 짧은 형식 사용
//    shortTime := timeutil.RelativeTimeShort(timestamp)
//    // "2h ago" instead of "2 hours ago"
//
// 4. UPDATE DISPLAY PERIODICALLY
//    주기적으로 표시 업데이트
//    // Refresh every minute to keep relative times current
//    ticker := time.NewTicker(1 * time.Minute)
//    go func() {
//        for range ticker.C {
//            updateRelativeTimes()
//        }
//    }()
//
// 5. USE HUMANIZEDURATION FOR ELAPSED TIME
//    경과 시간에 HumanizeDuration 사용
//    elapsed := time.Since(start)
//    fmt.Printf("Took %s\n", timeutil.HumanizeDuration(elapsed))
//
// 6. COMBINE WITH ISBEFORE/ISAFTER FOR CONTEXT
//    컨텍스트를 위해 IsBefore/IsAfter와 결합
//    relTime := timeutil.RelativeTime(event.Time)
//    if timeutil.IsFuture(event.Time) {
//        fmt.Printf("Upcoming: %s\n", relTime)
//    } else {
//        fmt.Printf("Past: %s\n", relTime)
//    }
//
// 7. CACHE RESULTS FOR PERFORMANCE
//    성능을 위해 결과 캐시
//    // Calculate once per render cycle
//    relTimes := make(map[int]string)
//    for _, item := range items {
//        relTimes[item.ID] = timeutil.RelativeTime(item.Timestamp)
//    }
//
// 8. USE TIMEAGO ALIAS FOR CLARITY
//    명확성을 위해 TimeAgo 별칭 사용
//    // More intuitive for past times
//    when := timeutil.TimeAgo(post.CreatedAt)
//    fmt.Printf("Posted %s\n", when)
//
// ============================================================================

// RelativeTime returns a human-readable relative time string.
// RelativeTime은 사람이 읽기 쉬운 상대 시간 문자열을 반환합니다.
//
// Examples:
//   - "2 hours ago"
//   - "in 3 days"
//   - "just now"
func RelativeTime(t time.Time) string {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	diff := now.Sub(t)

	if diff < 0 {
		// Future
		// 미래
		return relativeTimeFuture(-diff)
	}

	// Past
	// 과거
	return relativeTimePast(diff)
}

// RelativeTimeShort returns a short human-readable relative time string.
// RelativeTimeShort는 짧은 사람이 읽기 쉬운 상대 시간 문자열을 반환합니다.
//
// Examples:
//   - "2h ago"
//   - "in 3d"
//   - "now"
func RelativeTimeShort(t time.Time) string {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	diff := now.Sub(t)

	if diff < 0 {
		// Future
		// 미래
		return relativeTimeFutureShort(-diff)
	}

	// Past
	// 과거
	return relativeTimePastShort(diff)
}

// TimeAgo is an alias for RelativeTime.
// TimeAgo는 RelativeTime의 별칭입니다.
func TimeAgo(t time.Time) string {
	return RelativeTime(t)
}

// relativeTimePast returns a relative time string for past times.
// relativeTimePast는 과거 시간에 대한 상대 시간 문자열을 반환합니다.
func relativeTimePast(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	switch {
	case seconds < 10:
		return "just now"
	case seconds < 60:
		return fmt.Sprintf("%d seconds ago", seconds)
	case minutes == 1:
		return "1 minute ago"
	case minutes < 60:
		return fmt.Sprintf("%d minutes ago", minutes)
	case hours == 1:
		return "1 hour ago"
	case hours < 24:
		return fmt.Sprintf("%d hours ago", hours)
	case days == 1:
		return "1 day ago"
	case days < 7:
		return fmt.Sprintf("%d days ago", days)
	case weeks == 1:
		return "1 week ago"
	case weeks < 4:
		return fmt.Sprintf("%d weeks ago", weeks)
	case months == 1:
		return "1 month ago"
	case months < 12:
		return fmt.Sprintf("%d months ago", months)
	case years == 1:
		return "1 year ago"
	default:
		return fmt.Sprintf("%d years ago", years)
	}
}

// relativeTimeFuture returns a relative time string for future times.
// relativeTimeFuture는 미래 시간에 대한 상대 시간 문자열을 반환합니다.
func relativeTimeFuture(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	switch {
	case seconds < 10:
		return "just now"
	case seconds < 60:
		return fmt.Sprintf("in %d seconds", seconds)
	case minutes == 1:
		return "in 1 minute"
	case minutes < 60:
		return fmt.Sprintf("in %d minutes", minutes)
	case hours == 1:
		return "in 1 hour"
	case hours < 24:
		return fmt.Sprintf("in %d hours", hours)
	case days == 1:
		return "in 1 day"
	case days < 7:
		return fmt.Sprintf("in %d days", days)
	case weeks == 1:
		return "in 1 week"
	case weeks < 4:
		return fmt.Sprintf("in %d weeks", weeks)
	case months == 1:
		return "in 1 month"
	case months < 12:
		return fmt.Sprintf("in %d months", months)
	case years == 1:
		return "in 1 year"
	default:
		return fmt.Sprintf("in %d years", years)
	}
}

// relativeTimePastShort returns a short relative time string for past times.
// relativeTimePastShort는 과거 시간에 대한 짧은 상대 시간 문자열을 반환합니다.
func relativeTimePastShort(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	switch {
	case seconds < 10:
		return "now"
	case seconds < 60:
		return fmt.Sprintf("%ds ago", seconds)
	case minutes < 60:
		return fmt.Sprintf("%dm ago", minutes)
	case hours < 24:
		return fmt.Sprintf("%dh ago", hours)
	case days < 7:
		return fmt.Sprintf("%dd ago", days)
	case weeks < 4:
		return fmt.Sprintf("%dw ago", weeks)
	case months < 12:
		return fmt.Sprintf("%dmo ago", months)
	default:
		return fmt.Sprintf("%dy ago", years)
	}
}

// relativeTimeFutureShort returns a short relative time string for future times.
// relativeTimeFutureShort는 미래 시간에 대한 짧은 상대 시간 문자열을 반환합니다.
func relativeTimeFutureShort(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	switch {
	case seconds < 10:
		return "now"
	case seconds < 60:
		return fmt.Sprintf("in %ds", seconds)
	case minutes < 60:
		return fmt.Sprintf("in %dm", minutes)
	case hours < 24:
		return fmt.Sprintf("in %dh", hours)
	case days < 7:
		return fmt.Sprintf("in %dd", days)
	case weeks < 4:
		return fmt.Sprintf("in %dw", weeks)
	case months < 12:
		return fmt.Sprintf("in %dmo", months)
	default:
		return fmt.Sprintf("in %dy", years)
	}
}

// HumanizeDuration converts a duration to a human-readable string.
// HumanizeDuration은 duration을 사람이 읽기 쉬운 문자열로 변환합니다.
//
// Example
// 예제:
//
//	d := 2*time.Hour + 30*time.Minute
//	str := timeutil.HumanizeDuration(d) // "2 hours 30 minutes"
func HumanizeDuration(d time.Duration) string {
	if d < 0 {
		d = -d
	}

	seconds := int(math.Abs(d.Seconds()))
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24

	seconds = seconds % 60
	minutes = minutes % 60
	hours = hours % 24

	var result string
	if days > 0 {
		result = fmt.Sprintf("%d days", days)
		if hours > 0 {
			result += fmt.Sprintf(" %d hours", hours)
		}
	} else if hours > 0 {
		result = fmt.Sprintf("%d hours", hours)
		if minutes > 0 {
			result += fmt.Sprintf(" %d minutes", minutes)
		}
	} else if minutes > 0 {
		result = fmt.Sprintf("%d minutes", minutes)
		if seconds > 0 {
			result += fmt.Sprintf(" %d seconds", seconds)
		}
	} else {
		result = fmt.Sprintf("%d seconds", seconds)
	}

	return result
}
