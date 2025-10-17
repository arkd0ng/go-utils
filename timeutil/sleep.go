package timeutil

import "time"

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/sleep.go
// Purpose: Sleep utilities for time-based pausing and scheduling
//          시간 기반 일시 중지 및 스케줄링을 위한 Sleep 유틸리티
//
// This file provides convenient sleep operations that allow goroutines to
// pause until specific time points or boundaries. Unlike raw time.Sleep()
// which requires duration calculations, these functions accept target times
// or automatically calculate sleep durations to the next hour, day, or week.
// These operations are useful for scheduled tasks, periodic jobs, and
// time-synchronized processing.
//
// 이 파일은 고루틴이 특정 시점이나 경계까지 일시 중지할 수 있도록 하는 편리한
// sleep 연산을 제공합니다. 기간 계산이 필요한 원시 time.Sleep()와 달리, 이러한
// 함수는 대상 시간을 받거나 다음 시간, 일 또는 주까지 sleep 기간을 자동으로
// 계산합니다. 이러한 연산은 예약된 작업, 주기적 작업 및 시간 동기화 처리에
// 유용합니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. SLEEP UNTIL SPECIFIC TIME (특정 시간까지 Sleep)
//    - SleepUntil: Sleep until a given time.Time
//      주어진 time.Time까지 sleep
//    - Handles past times (returns immediately)
//      과거 시간 처리 (즉시 반환)
//    - Safe for dynamic scheduling
//      동적 스케줄링에 안전
//
// 2. SLEEP UNTIL NEXT BOUNDARY (다음 경계까지 Sleep)
//    - SleepUntilNextHour: Sleep until top of next hour (HH:00:00)
//      다음 시간의 시작까지 sleep (HH:00:00)
//    - SleepUntilNextDay: Sleep until midnight (00:00:00)
//      자정까지 sleep (00:00:00)
//    - SleepUntilNextWeek: Sleep until next Monday midnight
//      다음 월요일 자정까지 sleep
//    - Useful for periodic tasks
//      주기적 작업에 유용
//
// 3. AUTOMATIC DURATION CALCULATION (자동 기간 계산)
//    - No manual duration math required
//      수동 기간 계산 불필요
//    - Uses time.Until() internally
//      내부적으로 time.Until() 사용
//    - Precise to nanosecond
//      나노초까지 정밀
//
// 4. SAFE FOR PAST TIMES (과거 시간에 안전)
//    - SleepUntil returns immediately if target is past
//      대상이 과거인 경우 SleepUntil이 즉시 반환
//    - No negative sleep durations
//      음수 sleep 기간 없음
//    - Prevents infinite waits
//      무한 대기 방지
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. TIME-BASED INSTEAD OF DURATION-BASED (기간 기반 대신 시간 기반)
//    - Specify "when" not "how long"
//      "얼마나 오래"가 아니라 "언제" 지정
//    - More intuitive for scheduling
//      스케줄링에 더 직관적
//    - Avoids drift in loops
//      루프에서 드리프트 방지
//    - Example: Sleep until 15:00, not "sleep 2.5 hours"
//      예시: 15:00까지 sleep, "2.5시간 sleep"이 아님
//
// 2. BOUNDARY-AWARE SLEEP (경계 인식 Sleep)
//    - SleepUntilNextHour, NextDay, NextWeek understand time boundaries
//      SleepUntilNextHour, NextDay, NextWeek는 시간 경계 이해
//    - No manual hour/day calculation needed
//      수동 시간/일 계산 불필요
//    - Uses KST for consistent boundaries
//      일관된 경계를 위해 KST 사용
//
// 3. NON-BLOCKING FOR PAST TIMES (과거 시간에 대해 비차단)
//    - If target time already passed, return immediately
//      대상 시간이 이미 지났으면 즉시 반환
//    - Prevents goroutine deadlock
//      고루틴 교착 상태 방지
//    - Safe for dynamic target times
//      동적 대상 시간에 안전
//
// 4. PRECISION PRESERVATION (정밀도 보존)
//    - Uses time.Until() which is nanosecond-precise
//      나노초 정밀한 time.Until() 사용
//    - No rounding errors
//      반올림 오류 없음
//    - Wakes up at exact target time
//      정확한 대상 시간에 깨어남
//
// 5. GOROUTINE-FRIENDLY (고루틴 친화적)
//    - Only sleeps the calling goroutine
//      호출 고루틴만 sleep
//    - Doesn't block other goroutines
//      다른 고루틴 차단하지 않음
//    - Safe for concurrent use
//      동시 사용에 안전
//
// ============================================================================
// SLEEP OPERATIONS OVERVIEW / Sleep 연산 개요
// ============================================================================
//
// SLEEP TO SPECIFIC TIME (특정 시간까지 Sleep) - 1 function
// └─ SleepUntil : Sleep until given time.Time
//
// SLEEP TO NEXT BOUNDARY (다음 경계까지 Sleep) - 3 functions
// ├─ SleepUntilNextHour : Sleep until start of next hour (HH:00:00)
// ├─ SleepUntilNextDay  : Sleep until start of next day (00:00:00)
// └─ SleepUntilNextWeek : Sleep until next Monday (00:00:00)
//
// Total: 4 sleep functions
// 총: 4개의 sleep 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
//
// SLEEPUNTIL: O(1) computation + sleep duration
//   Duration calculation is instant
//   기간 계산은 즉각적
//   Sleep time is the actual wait time
//   Sleep 시간은 실제 대기 시간
//
// SLEEPUNTILNEXTHOUR: O(1) computation + sleep duration
//   Truncate and add operations are instant
//   Truncate 및 add 연산은 즉각적
//
// SLEEPUNTILNEXTDAY: O(1) computation + sleep duration
//   StartOfDay and AddDays are O(1)
//   StartOfDay 및 AddDays는 O(1)
//
// SLEEPUNTILNEXTWEEK: O(1) computation + sleep duration
//   StartOfWeek is O(1)
//   StartOfWeek는 O(1)
//
// SPACE COMPLEXITY (공간 복잡도):
// - All functions: O(1) - no additional allocation
//   모든 함수: O(1) - 추가 할당 없음
//
// PERFORMANCE NOTES:
// 성능 참고:
// 1. Duration calculation is negligible (<100 ns)
//    기간 계산은 무시할 만함 (<100 ns)
// 2. Actual sleep uses OS-level timer
//    실제 sleep은 OS 수준 타이머 사용
// 3. Wakeup precision: ~1 millisecond on most systems
//    깨우기 정밀도: 대부분 시스템에서 ~1 밀리초
// 4. No CPU usage while sleeping
//    sleep 중 CPU 사용 없음
//
// ============================================================================
// SLEEP SEMANTICS / Sleep 의미론
// ============================================================================
//
// SLEEPUNTIL BEHAVIOR:
// SleepUntil 동작:
// - target > now: Sleep for (target - now) duration
//   대상 > 현재: (대상 - 현재) 기간 동안 sleep
// - target <= now: Return immediately (no sleep)
//   대상 <= 현재: 즉시 반환 (sleep 없음)
// - Equivalent to: time.Sleep(time.Until(t))
//   동등: time.Sleep(time.Until(t))
//
// SLEEPUNTILNEXTHOUR:
// - Sleeps until HH:00:00 of next hour
//   다음 시간의 HH:00:00까지 sleep
// - Example: 14:37:22 → sleeps until 15:00:00
//   예시: 14:37:22 → 15:00:00까지 sleep
// - Duration: Always less than 1 hour
//   기간: 항상 1시간 미만
//
// SLEEPUNTILNEXTDAY:
// - Sleeps until 00:00:00 of next day (in KST)
//   다음 날 00:00:00까지 sleep (KST)
// - Example: 2024-01-15 18:30 → sleeps until 2024-01-16 00:00:00
//   예시: 2024-01-15 18:30 → 2024-01-16 00:00:00까지 sleep
// - Duration: Always less than 24 hours
//   기간: 항상 24시간 미만
//
// SLEEPUNTILNEXTWEEK:
// - Sleeps until Monday 00:00:00 of next week (in KST)
//   다음 주 월요일 00:00:00까지 sleep (KST)
// - Monday = week start (ISO 8601)
//   월요일 = 주 시작 (ISO 8601)
// - Example: Friday 10:00 → sleeps until next Monday 00:00
//   예시: 금요일 10:00 → 다음 월요일 00:00까지 sleep
// - Duration: Up to 7 days
//   기간: 최대 7일
//
// GOROUTINE SLEEP:
// 고루틴 Sleep:
// - Only the calling goroutine sleeps
//   호출 고루틴만 sleep
// - Other goroutines continue execution
//   다른 고루틴은 계속 실행
// - No global effect
//   전역 효과 없음
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Hourly Task (시간마다 작업)
// Use case: Run task at the top of every hour
// 사용 사례: 매 시간마다 작업 실행
//
//   func runHourlyTask() {
//       for {
//           // Wait until next hour
//           timeutil.SleepUntilNextHour()
//
//           // Run task
//           fmt.Println("Hourly task at", timeutil.Now())
//           processHourlyData()
//       }
//   }
//
//   go runHourlyTask()
//
// PATTERN 2: Daily Batch Job (일일 배치 작업)
// Use case: Run report generation at midnight
// 사용 사례: 자정에 보고서 생성 실행
//
//   func dailyReportJob() {
//       for {
//           // Wait until midnight
//           timeutil.SleepUntilNextDay()
//
//           // Generate daily report
//           fmt.Println("Generating daily report for", timeutil.FormatDate(time.Now()))
//           generateDailyReport()
//       }
//   }
//
//   go dailyReportJob()
//
// PATTERN 3: Weekly Maintenance (주간 유지보수)
// Use case: Weekly database cleanup
// 사용 사례: 주간 데이터베이스 정리
//
//   func weeklyMaintenance() {
//       for {
//           // Wait until next Monday
//           timeutil.SleepUntilNextWeek()
//
//           fmt.Println("Starting weekly maintenance")
//           cleanupOldRecords()
//           vacuumDatabase()
//       }
//   }
//
//   go weeklyMaintenance()
//
// PATTERN 4: Scheduled Event (예약된 이벤트)
// Use case: Sleep until specific meeting time
// 사용 사례: 특정 회의 시간까지 sleep
//
//   meetingTime := time.Date(2024, 1, 15, 15, 0, 0, 0, timeutil.KST)
//
//   fmt.Println("Waiting for meeting...")
//   timeutil.SleepUntil(meetingTime)
//
//   fmt.Println("Meeting started!")
//   sendMeetingNotifications()
//
// PATTERN 5: Rate Limiting with Fixed Time Window (고정 시간 창이 있는 속도 제한)
// Use case: Reset rate limit counter every hour
// 사용 사례: 매 시간마다 속도 제한 카운터 재설정
//
//   func rateLimitReset() {
//       requestCount := 0
//       mu := &sync.Mutex{}
//
//       // Reset counter every hour
//       go func() {
//           for {
//               timeutil.SleepUntilNextHour()
//               mu.Lock()
//               requestCount = 0
//               mu.Unlock()
//               fmt.Println("Rate limit reset")
//           }
//       }()
//   }
//
// PATTERN 6: Delayed Start (지연된 시작)
// Use case: Start service at specific time
// 사용 사례: 특정 시간에 서비스 시작
//
//   startTime := time.Date(2024, 1, 15, 9, 0, 0, 0, timeutil.KST)
//
//   fmt.Println("Service scheduled to start at", timeutil.FormatDateTime(startTime))
//   timeutil.SleepUntil(startTime)
//
//   fmt.Println("Starting service now!")
//   startService()
//
// PATTERN 7: Time-Aligned Polling (시간 정렬 폴링)
// Use case: Poll API at exact minute boundaries
// 사용 사례: 정확한 분 경계에서 API 폴링
//
//   func pollAtMinuteBoundaries() {
//       for {
//           // Sleep until next minute
//           now := time.Now()
//           nextMinute := now.Truncate(time.Minute).Add(time.Minute)
//           timeutil.SleepUntil(nextMinute)
//
//           // Poll API
//           data := pollAPI()
//           processData(data)
//       }
//   }
//
// PATTERN 8: Synchronized Multi-Goroutine Task (동기화된 다중 고루틴 작업)
// Use case: Multiple workers start at the same time
// 사용 사례: 여러 작업자가 동시에 시작
//
//   startTime := time.Now().Add(5 * time.Second)
//
//   for i := 0; i < 10; i++ {
//       go func(workerID int) {
//           // All workers wait until same time
//           timeutil.SleepUntil(startTime)
//
//           fmt.Printf("Worker %d started\n", workerID)
//           doWork(workerID)
//       }(i)
//   }
//
// PATTERN 9: Conditional Sleep (조건부 Sleep)
// Use case: Sleep only if target time is in future
// 사용 사례: 대상 시간이 미래인 경우에만 sleep
//
//   func processScheduledTask(task Task) {
//       if task.ScheduledAt.After(time.Now()) {
//           fmt.Printf("Waiting until %s\n", timeutil.FormatDateTime(task.ScheduledAt))
//           timeutil.SleepUntil(task.ScheduledAt)
//       } else {
//           fmt.Println("Task is overdue, executing immediately")
//       }
//
//       executeTask(task)
//   }
//
// PATTERN 10: Loop Without Drift (드리프트 없는 루프)
// Use case: Execute task exactly every hour, no drift
// 사용 사례: 드리프트 없이 정확히 매 시간마다 작업 실행
//
//   func preciselySleepingLoop() {
//       for {
//           // Precise: Always wakes at HH:00:00
//           timeutil.SleepUntilNextHour()
//           executeTask()
//       }
//   }
//
//   // Compare with drift-prone version:
//   // 드리프트가 발생하기 쉬운 버전과 비교:
//   func driftingLoop() {
//       for {
//           time.Sleep(1 * time.Hour)  // Drifts over time
//           executeTask()              // 시간이 지나면서 드리프트
//       }
//   }
//
// ============================================================================
// EDGE CASES / 경계 사례
// ============================================================================
//
// PAST TARGET TIME:
// 과거 대상 시간:
//   pastTime := time.Now().Add(-5 * time.Minute)
//   timeutil.SleepUntil(pastTime)  // Returns immediately, no sleep
//                                  // 즉시 반환, sleep 없음
//
// VERY SHORT SLEEP:
// 매우 짧은 Sleep:
//   // Sleep for 1 nanosecond
//   target := time.Now().Add(1 * time.Nanosecond)
//   timeutil.SleepUntil(target)
//   // May not be precisely 1 nanosecond due to OS scheduler
//   // OS 스케줄러로 인해 정확히 1 나노초가 아닐 수 있음
//
// VERY LONG SLEEP:
// 매우 긴 Sleep:
//   // Sleep for 1 year
//   nextYear := time.Now().AddDate(1, 0, 0)
//   timeutil.SleepUntil(nextYear)
//   // Valid, but consider using time.Ticker for very long durations
//   // 유효하지만 매우 긴 기간에는 time.Ticker 사용 고려
//
// MIDNIGHT ON DST TRANSITION:
// DST 전환 시 자정:
//   // In timezones with DST, midnight may be ambiguous
//   // SleepUntilNextDay uses KST which has no DST
//   // DST가 있는 타임존에서 자정은 모호할 수 있음
//   // SleepUntilNextDay는 DST가 없는 KST 사용
//
// GOROUTINE CANCELLATION:
// 고루틴 취소:
//   // Sleep cannot be cancelled once started
//   // Use context.Context for cancellable waits
//   // Sleep은 시작되면 취소할 수 없음
//   // 취소 가능한 대기에 context.Context 사용
//   ctx, cancel := context.WithDeadline(context.Background(), targetTime)
//   <-ctx.Done()
//
// ZERO TIME:
// 제로 시간:
//   var t time.Time
//   timeutil.SleepUntil(t)  // Returns immediately (zero time is past)
//                           // 즉시 반환 (제로 시간은 과거)
//
// ============================================================================
// COMPARISON WITH ALTERNATIVES / 대안과의 비교
// ============================================================================
//
// VS. time.Sleep():
// - timeutil.SleepUntil: Specify target time
//   timeutil.SleepUntil: 대상 시간 지정
// - time.Sleep: Specify duration
//   time.Sleep: 기간 지정
//
// Example:
// 예시:
//   // With SleepUntil (clearer intent)
//   target := time.Date(2024, 1, 15, 15, 0, 0, 0, timeutil.KST)
//   timeutil.SleepUntil(target)
//
//   // With time.Sleep (manual calculation)
//   target := time.Date(2024, 1, 15, 15, 0, 0, 0, timeutil.KST)
//   duration := time.Until(target)
//   if duration > 0 {
//       time.Sleep(duration)
//   }
//
// VS. time.Ticker:
// - timeutil.SleepUntil*: One-time sleep to target
//   timeutil.SleepUntil*: 대상까지 일회성 sleep
// - time.Ticker: Repeated intervals
//   time.Ticker: 반복 간격
//
// Use SleepUntil for: Scheduled one-time events
// SleepUntil 사용: 예약된 일회성 이벤트
// Use Ticker for: Continuous periodic tasks (with drift)
// Ticker 사용: 지속적인 주기 작업 (드리프트 있음)
//
// VS. time.After():
// - timeutil.SleepUntil: Blocks goroutine
//   timeutil.SleepUntil: 고루틴 차단
// - time.After: Returns channel for select
//   time.After: select용 채널 반환
//
// Use SleepUntil for: Simple sequential flow
// SleepUntil 사용: 간단한 순차 흐름
// Use time.After for: Select statements with multiple channels
// time.After 사용: 여러 채널이 있는 select 문
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// THREAD-SAFE FUNCTIONS:
// 스레드 안전 함수:
// - All sleep functions are thread-safe
//   모든 sleep 함수는 스레드 안전
// - Only affect calling goroutine
//   호출 고루틴에만 영향
// - No shared mutable state
//   공유 변경 가능 상태 없음
// - Safe for concurrent use
//   동시 사용에 안전
//
// CONCURRENT USAGE:
// 동시 사용:
//   // Multiple goroutines can call SleepUntil concurrently
//   for i := 0; i < 100; i++ {
//       go func(id int) {
//           timeutil.SleepUntilNextHour()
//           fmt.Printf("Goroutine %d woke up\n", id)
//       }(i)
//   }
//   // All 100 goroutines wake at the same time (next hour)
//   // 모든 100개 고루틴이 동시에 깨어남 (다음 시간)
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM arithmetic.go:
// - StartOfDay: For SleepUntilNextDay
// - AddDays: For SleepUntilNextDay
// - StartOfWeek: For SleepUntilNextWeek
//
// FROM constants.go:
// - DaysPerWeek: For SleepUntilNextWeek
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.Until: Calculate duration until target
// - time.Sleep: Actual sleep operation
// - time.Truncate: For hour rounding
//
// USED BY (사용처):
// - Scheduled tasks (cron-like jobs)
//   예약된 작업 (cron과 유사한 작업)
// - Periodic batch jobs
//   주기적 배치 작업
// - Rate limiters (time window resets)
//   속도 제한기 (시간 창 재설정)
// - Synchronized goroutines
//   동기화된 고루틴
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE SLEEPUNTIL FOR SCHEDULED TASKS
//    예약된 작업에 SleepUntil 사용
//    timeutil.SleepUntil(scheduledTime)
//
// 2. USE SLEEPUNTILNEXT* FOR PERIODIC JOBS
//    주기적 작업에 SleepUntilNext* 사용
//    for { timeutil.SleepUntilNextHour(); runTask() }
//
// 3. RUN IN SEPARATE GOROUTINE FOR NON-BLOCKING
//    비차단을 위해 별도의 고루틴에서 실행
//    go func() { timeutil.SleepUntilNextDay(); doTask() }()
//
// 4. CHECK IF TARGET IS PAST BEFORE SLEEPING
//    sleep 전에 대상이 과거인지 확인
//    if target.After(time.Now()) { timeutil.SleepUntil(target) }
//
// 5. USE CONTEXT FOR CANCELLABLE WAITS
//    취소 가능한 대기에 context 사용
//    select { case <-ctx.Done(): return; case <-time.After(...): }
//
// 6. AVOID VERY LONG SLEEPS IN PRODUCTION
//    프로덕션에서 매우 긴 sleep 방지
//    // Consider time.Ticker or cron scheduler instead
//
// 7. LOG WHEN SLEEPING FOR DEBUGGING
//    디버깅을 위해 sleep 시 로그
//    fmt.Printf("Sleeping until %s\n", timeutil.FormatDateTime(target))
//
// 8. PREFER SLEEPUNTILNEXT* OVER MANUAL CALCULATION
//    수동 계산보다 SleepUntilNext* 선호
//    // Clearer intent, no math errors
//
// ============================================================================

// SleepUntil sleeps until the specified time.
// SleepUntil은 지정된 시간까지 sleep합니다.
//
// If the target time is in the past, it returns immediately.
// 대상 시간이 과거인 경우 즉시 반환합니다.
//
// Example
// 예제:
//
//	target := time.Now().Add(5 * time.Second)
//	timeutil.SleepUntil(target) // Sleeps for 5 seconds
func SleepUntil(t time.Time) {
	duration := time.Until(t)
	if duration > 0 {
		time.Sleep(duration)
	}
}

// SleepUntilNextHour sleeps until the start of the next hour.
// SleepUntilNextHour는 다음 시간의 시작까지 sleep합니다.
//
// Example
// 예제:
//
//	// If current time is 14:30:45, sleeps until 15:00:00
//	timeutil.SleepUntilNextHour()
func SleepUntilNextHour() {
	now := time.Now()
	next := now.Truncate(time.Hour).Add(time.Hour)
	SleepUntil(next)
}

// SleepUntilNextDay sleeps until the start of the next day (00:00:00).
// SleepUntilNextDay는 다음 날의 시작(00:00:00)까지 sleep합니다.
//
// Example
// 예제:
//
//	// Sleeps until midnight
//	timeutil.SleepUntilNextDay()
func SleepUntilNextDay() {
	tomorrow := AddDays(StartOfDay(time.Now()), 1)
	SleepUntil(tomorrow)
}

// SleepUntilNextWeek sleeps until the start of the next week (Monday 00:00:00).
// SleepUntilNextWeek는 다음 주의 시작(월요일 00:00:00)까지 sleep합니다.
//
// Example
// 예제:
//
//	// Sleeps until next Monday midnight
//	timeutil.SleepUntilNextWeek()
func SleepUntilNextWeek() {
	now := time.Now()
	nextWeekStart := StartOfWeek(now).AddDate(0, 0, DaysPerWeek)
	SleepUntil(nextWeekStart)
}
