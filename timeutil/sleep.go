package timeutil

import "time"

// SleepUntil sleeps until the specified time.
// SleepUntil은 지정된 시간까지 sleep합니다.
//
// If the target time is in the past, it returns immediately.
// 대상 시간이 과거인 경우 즉시 반환합니다.
//
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
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
// Example / 예제:
//
//	// Sleeps until next Monday midnight
//	timeutil.SleepUntilNextWeek()
func SleepUntilNextWeek() {
	now := time.Now()
	nextWeekStart := StartOfWeek(now).AddDate(0, 0, DaysPerWeek)
	SleepUntil(nextWeekStart)
}
