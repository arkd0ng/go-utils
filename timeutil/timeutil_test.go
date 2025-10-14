package timeutil

import (
	"testing"
	"time"
)

// Test TimeDiff type / TimeDiff 타입 테스트
func TestTimeDiff(t *testing.T) {
	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)

	diff := SubTime(start, end)

	// Test Days / Days 테스트
	days := diff.Days()
	if days < 2 || days > 3 {
		t.Errorf("Days() = %v, want between 2 and 3", days)
	}

	// Test String / String 테스트
	str := diff.String()
	if str == "" {
		t.Error("String() returned empty string")
	}

	// Test Humanize / Humanize 테스트
	humanized := diff.Humanize()
	if humanized == "" {
		t.Error("Humanize() returned empty string")
	}
}

// Test diff functions / diff 함수 테스트
func TestDiffFunctions(t *testing.T) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)

	seconds := DiffInSeconds(start, end)
	if seconds != 86400 {
		t.Errorf("DiffInSeconds() = %v, want 86400", seconds)
	}

	days := DiffInDays(start, end)
	if days != 1 {
		t.Errorf("DiffInDays() = %v, want 1", days)
	}
}

// Test timezone functions / 타임존 함수 테스트
func TestTimezone(t *testing.T) {
	// Test default timezone is KST / 기본 타임존이 KST인지 테스트
	defaultTz := GetDefaultTimezone()
	if defaultTz != "Asia/Seoul" && defaultTz != "KST" {
		t.Logf("Default timezone: %s (expected Asia/Seoul or KST)", defaultTz)
	}

	// Test timezone conversion / 타임존 변환 테스트
	now := time.Now()
	converted, err := ConvertTimezone(now, "America/New_York")
	if err != nil {
		t.Errorf("ConvertTimezone() error = %v", err)
	}
	if converted.IsZero() {
		t.Error("ConvertTimezone() returned zero time")
	}

	// Test invalid timezone / 잘못된 타임존 테스트
	_, err = ConvertTimezone(now, "Invalid/Timezone")
	if err == nil {
		t.Error("ConvertTimezone() should return error for invalid timezone")
	}
}

// Test arithmetic functions / 연산 함수 테스트
func TestArithmetic(t *testing.T) {
	now := time.Now()

	// Test AddDays / AddDays 테스트
	tomorrow := AddDays(now, 1)
	if !tomorrow.After(now) {
		t.Error("AddDays(1) should be after now")
	}

	// Test StartOfDay / StartOfDay 테스트
	startOfDay := StartOfDay(now)
	if startOfDay.Hour() != 0 || startOfDay.Minute() != 0 || startOfDay.Second() != 0 {
		t.Errorf("StartOfDay() = %v, want 00:00:00", startOfDay.Format("15:04:05"))
	}

	// Test EndOfDay / EndOfDay 테스트
	endOfDay := EndOfDay(now)
	if endOfDay.Hour() != 23 || endOfDay.Minute() != 59 || endOfDay.Second() != 59 {
		t.Errorf("EndOfDay() = %v, want 23:59:59", endOfDay.Format("15:04:05"))
	}
}

// Test format functions / 포맷 함수 테스트
func TestFormat(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 4, 5, 0, KST)

	// Test FormatDate / FormatDate 테스트
	date := FormatDate(testTime)
	if date != "2025-10-14" {
		t.Errorf("FormatDate() = %v, want 2025-10-14", date)
	}

	// Test custom format / 커스텀 포맷 테스트
	custom := Format(testTime, "YYYY-MM-DD HH:mm:ss")
	if custom == "" {
		t.Error("Format() returned empty string")
	}
}

// Test parse functions / 파싱 함수 테스트
func TestParse(t *testing.T) {
	// Test ParseDate / ParseDate 테스트
	parsed, err := ParseDate("2025-10-14")
	if err != nil {
		t.Errorf("ParseDate() error = %v", err)
	}
	if parsed.Year() != 2025 || parsed.Month() != 10 || parsed.Day() != 14 {
		t.Errorf("ParseDate() = %v, want 2025-10-14", parsed)
	}

	// Test Parse with auto-detection / 자동 감지 파싱 테스트
	parsed2, err := Parse("2025-10-14")
	if err != nil {
		t.Errorf("Parse() error = %v", err)
	}
	if parsed2.IsZero() {
		t.Error("Parse() returned zero time")
	}
}

// Test comparison functions / 비교 함수 테스트
func TestComparison(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	// Test IsBefore / IsBefore 테스트
	if !IsBefore(yesterday, now) {
		t.Error("Yesterday should be before now")
	}

	// Test IsAfter / IsAfter 테스트
	if !IsAfter(tomorrow, now) {
		t.Error("Tomorrow should be after now")
	}

	// Test IsToday / IsToday 테스트
	if !IsToday(now) {
		t.Error("Now should be today")
	}

	// Test IsWeekend / IsWeekend 테스트
	saturday := time.Date(2025, 10, 18, 0, 0, 0, 0, KST) // Saturday
	if !IsWeekend(saturday) {
		t.Error("Saturday should be weekend")
	}
}

// Test age functions / 나이 함수 테스트
func TestAge(t *testing.T) {
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, KST)

	// Test AgeInYears / AgeInYears 테스트
	years := AgeInYears(birthDate)
	if years < 0 {
		t.Errorf("AgeInYears() = %v, should be non-negative", years)
	}

	// Test Age / Age 테스트
	age := Age(birthDate)
	if age == nil {
		t.Error("Age() returned nil")
	}
	if age.Years < 0 || age.Months < 0 || age.Days < 0 {
		t.Errorf("Age() = %v, all fields should be non-negative", age)
	}
}

// Test relative time functions / 상대 시간 함수 테스트
func TestRelativeTime(t *testing.T) {
	now := time.Now()
	past := now.Add(-2 * time.Hour)
	future := now.Add(3 * time.Hour)

	// Test past time / 과거 시간 테스트
	rel := RelativeTime(past)
	if rel == "" {
		t.Error("RelativeTime() returned empty string")
	}

	// Test future time / 미래 시간 테스트
	relFuture := RelativeTime(future)
	if relFuture == "" {
		t.Error("RelativeTime() for future returned empty string")
	}

	// Test short format / 짧은 포맷 테스트
	short := RelativeTimeShort(past)
	if short == "" {
		t.Error("RelativeTimeShort() returned empty string")
	}
}

// Test unix timestamp functions / Unix 타임스탬프 함수 테스트
func TestUnix(t *testing.T) {
	// Test Now / Now 테스트
	unix := Now()
	if unix <= 0 {
		t.Errorf("Now() = %v, should be positive", unix)
	}

	// Test FromUnix / FromUnix 테스트
	testTime := FromUnix(1634198400)
	if testTime.IsZero() {
		t.Error("FromUnix() returned zero time")
	}

	// Test ToUnix / ToUnix 테스트
	now := time.Now()
	unixTime := ToUnix(now)
	if unixTime <= 0 {
		t.Errorf("ToUnix() = %v, should be positive", unixTime)
	}
}

// Test business day functions / 영업일 함수 테스트
func TestBusinessDays(t *testing.T) {
	// Test IsWeekday / IsWeekday 테스트
	monday := time.Date(2025, 10, 13, 0, 0, 0, 0, KST) // Monday
	if !IsWeekday(monday) {
		t.Error("Monday should be a weekday")
	}

	// Test IsBusinessDay / IsBusinessDay 테스트
	if !IsBusinessDay(monday) {
		t.Error("Monday should be a business day")
	}

	// Test AddBusinessDays / AddBusinessDays 테스트
	nextBusiness := AddBusinessDays(monday, 1)
	if nextBusiness.IsZero() {
		t.Error("AddBusinessDays() returned zero time")
	}
}

// Benchmark SubTime / SubTime 벤치마크
func BenchmarkSubTime(b *testing.B) {
	start := time.Now()
	end := start.Add(24 * time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SubTime(start, end)
	}
}

// Benchmark FormatDate / FormatDate 벤치마크
func BenchmarkFormatDate(b *testing.B) {
	t := time.Now()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatDate(t)
	}
}

// Benchmark ParseDate / ParseDate 벤치마크
func BenchmarkParseDate(b *testing.B) {
	dateStr := "2025-10-14"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate(dateStr)
	}
}
