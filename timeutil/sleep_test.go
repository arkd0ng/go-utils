package timeutil

import (
	"testing"
	"time"
)

// TestSleepUntil tests SleepUntil function
// SleepUntil 함수 테스트
func TestSleepUntil(t *testing.T) {
	// Test with past time (should return immediately)
	// 과거 시간으로 테스트 (즉시 반환되어야 함)
	t.Run("Past time", func(t *testing.T) {
		past := time.Now().Add(-1 * time.Second)
		start := time.Now()
		SleepUntil(past)
		duration := time.Since(start)

		// Should return almost immediately (< 10ms)
		// 거의 즉시 반환되어야 함 (< 10ms)
		if duration > 10*time.Millisecond {
			t.Errorf("SleepUntil() with past time took %v, want < 10ms", duration)
		}
	})

	// Test with future time
	// 미래 시간으로 테스트
	t.Run("Future time", func(t *testing.T) {
		sleepDuration := 100 * time.Millisecond
		future := time.Now().Add(sleepDuration)
		start := time.Now()
		SleepUntil(future)
		duration := time.Since(start)

		// Should sleep approximately the specified duration (±20ms tolerance)
		// 지정된 기간만큼 sleep해야 함 (±20ms 허용)
		if duration < sleepDuration-20*time.Millisecond || duration > sleepDuration+20*time.Millisecond {
			t.Errorf("SleepUntil() slept for %v, want ~%v", duration, sleepDuration)
		}
	})
}

// TestSleepUntilNextHour tests SleepUntilNextHour function
// SleepUntilNextHour 함수 테스트
func TestSleepUntilNextHour(t *testing.T) {
	t.Skip("Skipping SleepUntilNextHour test - takes too long")

	// This test would take up to an hour, so we skip it in normal testing
	// 이 테스트는 최대 1시간이 걸릴 수 있으므로 일반 테스팅에서 건너뜀
	now := time.Now()
	expected := now.Truncate(time.Hour).Add(time.Hour)

	start := time.Now()
	SleepUntilNextHour()
	actualWakeTime := time.Now()

	// Check if we woke up at the expected time (±1 second tolerance)
	// 예상 시간에 깨어났는지 확인 (±1초 허용)
	if actualWakeTime.Before(expected.Add(-1*time.Second)) || actualWakeTime.After(expected.Add(1*time.Second)) {
		t.Errorf("SleepUntilNextHour() woke at %v, want ~%v", actualWakeTime, expected)
	}

	t.Logf("Slept for %v", time.Since(start))
}

// TestSleepUntilNextDay tests SleepUntilNextDay function
// SleepUntilNextDay 함수 테스트
func TestSleepUntilNextDay(t *testing.T) {
	t.Skip("Skipping SleepUntilNextDay test - takes too long")

	// This test would take up to a day, so we skip it in normal testing
	// 이 테스트는 최대 하루가 걸릴 수 있으므로 일반 테스팅에서 건너뜀
	now := time.Now()
	expected := AddDays(StartOfDay(now), 1)

	start := time.Now()
	SleepUntilNextDay()
	actualWakeTime := time.Now()

	// Check if we woke up at the expected time (±1 second tolerance)
	// 예상 시간에 깨어났는지 확인 (±1초 허용)
	if actualWakeTime.Before(expected.Add(-1*time.Second)) || actualWakeTime.After(expected.Add(1*time.Second)) {
		t.Errorf("SleepUntilNextDay() woke at %v, want ~%v", actualWakeTime, expected)
	}

	t.Logf("Slept for %v", time.Since(start))
}

// TestSleepUntilNextWeek tests SleepUntilNextWeek function
// SleepUntilNextWeek 함수 테스트
func TestSleepUntilNextWeek(t *testing.T) {
	t.Skip("Skipping SleepUntilNextWeek test - takes too long")

	// This test would take up to a week, so we skip it in normal testing
	// 이 테스트는 최대 일주일이 걸릴 수 있으므로 일반 테스팅에서 건너뜀
	now := time.Now()
	expected := StartOfWeek(now).AddDate(0, 0, DaysPerWeek)

	start := time.Now()
	SleepUntilNextWeek()
	actualWakeTime := time.Now()

	// Check if we woke up at the expected time (±1 second tolerance)
	// 예상 시간에 깨어났는지 확인 (±1초 허용)
	if actualWakeTime.Before(expected.Add(-1*time.Second)) || actualWakeTime.After(expected.Add(1*time.Second)) {
		t.Errorf("SleepUntilNextWeek() woke at %v, want ~%v", actualWakeTime, expected)
	}

	t.Logf("Slept for %v", time.Since(start))
}

// Benchmark tests for sleep functions
// sleep 함수 벤치마크 테스트

// BenchmarkSleepUntil benchmarks SleepUntil function
// SleepUntil 함수 벤치마크
func BenchmarkSleepUntil(b *testing.B) {
	// Use past time to benchmark the overhead
	// 과거 시간을 사용하여 오버헤드 벤치마크
	past := time.Now().Add(-1 * time.Second)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SleepUntil(past)
	}
}
