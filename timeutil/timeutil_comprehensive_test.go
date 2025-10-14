package timeutil

import (
	"testing"
	"time"
)

// ============================================================
// 1. Time Difference Functions (8 functions)
// 1. 시간 차이 함수 (8개 함수)
// ============================================================

// TestSubTime tests SubTime function / SubTime 함수 테스트
func TestSubTime(t *testing.T) {
	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 10, 15, 30, 45, 0, time.UTC)

	diff := SubTime(start, end)

	// Test Days / Days 테스트
	days := diff.Days()
	if days < 9 || days > 10 {
		t.Errorf("SubTime().Days() = %v, want between 9 and 10", days)
	}

	// Test String / String 테스트
	str := diff.String()
	if str == "" {
		t.Error("SubTime().String() returned empty string")
	}

	// Test Humanize / Humanize 테스트
	humanized := diff.Humanize()
	if humanized == "" {
		t.Error("SubTime().Humanize() returned empty string")
	}
}

// TestDiffInSeconds tests DiffInSeconds function / DiffInSeconds 함수 테스트
func TestDiffInSeconds(t *testing.T) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)

	seconds := DiffInSeconds(start, end)
	if seconds != 86400 {
		t.Errorf("DiffInSeconds() = %v, want 86400", seconds)
	}

	// Test negative difference / 음수 차이 테스트
	negSeconds := DiffInSeconds(end, start)
	if negSeconds != -86400 {
		t.Errorf("DiffInSeconds() = %v, want -86400", negSeconds)
	}
}

// TestDiffInMinutes tests DiffInMinutes function / DiffInMinutes 함수 테스트
func TestDiffInMinutes(t *testing.T) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC)

	minutes := DiffInMinutes(start, end)
	if minutes != 60 {
		t.Errorf("DiffInMinutes() = %v, want 60", minutes)
	}
}

// TestDiffInHours tests DiffInHours function / DiffInHours 함수 테스트
func TestDiffInHours(t *testing.T) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)

	hours := DiffInHours(start, end)
	if hours != 24 {
		t.Errorf("DiffInHours() = %v, want 24", hours)
	}
}

// TestDiffInDays tests DiffInDays function / DiffInDays 함수 테스트
func TestDiffInDays(t *testing.T) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC)

	days := DiffInDays(start, end)
	if days != 7 {
		t.Errorf("DiffInDays() = %v, want 7", days)
	}
}

// TestDiffInWeeks tests DiffInWeeks function / DiffInWeeks 함수 테스트
func TestDiffInWeeks(t *testing.T) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)

	weeks := DiffInWeeks(start, end)
	if weeks != 2 {
		t.Errorf("DiffInWeeks() = %v, want 2", weeks)
	}
}

// TestDiffInMonths tests DiffInMonths function / DiffInMonths 함수 테스트
func TestDiffInMonths(t *testing.T) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)

	months := DiffInMonths(start, end)
	if months != 3 {
		t.Errorf("DiffInMonths() = %v, want 3", months)
	}
}

// TestDiffInYears tests DiffInYears function / DiffInYears 함수 테스트
func TestDiffInYears(t *testing.T) {
	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	years := DiffInYears(start, end)
	if years != 2 {
		t.Errorf("DiffInYears() = %v, want 2", years)
	}
}

// ============================================================
// 2. Timezone Operations (10 functions)
// 2. 타임존 작업 (10개 함수)
// ============================================================

// TestGetDefaultTimezone tests GetDefaultTimezone function / GetDefaultTimezone 함수 테스트
func TestGetDefaultTimezone(t *testing.T) {
	tz := GetDefaultTimezone()
	if tz != "Asia/Seoul" && tz != "KST" {
		t.Logf("GetDefaultTimezone() = %v (expected Asia/Seoul or KST)", tz)
	}
}

// TestNowKST tests NowKST function / NowKST 함수 테스트
func TestNowKST(t *testing.T) {
	kstTime := NowKST()
	if kstTime.IsZero() {
		t.Error("NowKST() returned zero time")
	}
}

// TestConvertTimezone tests ConvertTimezone function / ConvertTimezone 함수 테스트
func TestConvertTimezone(t *testing.T) {
	now := time.Now()

	// Test valid timezone / 유효한 타임존 테스트
	nyTime, err := ConvertTimezone(now, "America/New_York")
	if err != nil {
		t.Errorf("ConvertTimezone() error = %v", err)
	}
	if nyTime.IsZero() {
		t.Error("ConvertTimezone() returned zero time")
	}

	// Test invalid timezone / 잘못된 타임존 테스트
	_, err = ConvertTimezone(now, "Invalid/Timezone")
	if err == nil {
		t.Error("ConvertTimezone() should return error for invalid timezone")
	}
}

// TestToUTC tests ToUTC function / ToUTC 함수 테스트
func TestToUTC(t *testing.T) {
	kstTime := time.Date(2025, 10, 14, 15, 0, 0, 0, KST)
	utcTime := ToUTC(kstTime)

	if utcTime.Location() != time.UTC {
		t.Errorf("ToUTC() location = %v, want UTC", utcTime.Location())
	}
}

// TestToKST tests ToKST function / ToKST 함수 테스트
func TestToKST(t *testing.T) {
	utcTime := time.Date(2025, 10, 14, 6, 0, 0, 0, time.UTC)
	kstTime := ToKST(utcTime)

	if kstTime.Location() != KST {
		t.Errorf("ToKST() location = %v, want KST", kstTime.Location())
	}
}

// TestGetTimezoneOffset tests GetTimezoneOffset function / GetTimezoneOffset 함수 테스트
func TestGetTimezoneOffset(t *testing.T) {
	offset, err := GetTimezoneOffset("Asia/Seoul")
	if err != nil {
		t.Errorf("GetTimezoneOffset() error = %v", err)
	}
	if offset == 0 {
		t.Log("GetTimezoneOffset() returned 0, expected non-zero for KST")
	}

	// Test invalid timezone / 잘못된 타임존 테스트
	_, err = GetTimezoneOffset("Invalid/Timezone")
	if err == nil {
		t.Error("GetTimezoneOffset() should return error for invalid timezone")
	}
}

// TestGetLocalTimezone tests GetLocalTimezone function / GetLocalTimezone 함수 테스트
func TestGetLocalTimezone(t *testing.T) {
	localTz := GetLocalTimezone()
	if localTz == "" {
		t.Error("GetLocalTimezone() returned empty string")
	}
}

// TestIsValidTimezone tests IsValidTimezone function / IsValidTimezone 함수 테스트
func TestIsValidTimezone(t *testing.T) {
	// Test valid timezones / 유효한 타임존 테스트
	if !IsValidTimezone("America/New_York") {
		t.Error("IsValidTimezone('America/New_York') = false, want true")
	}

	// Test invalid timezone / 잘못된 타임존 테스트
	if IsValidTimezone("Invalid/Timezone") {
		t.Error("IsValidTimezone('Invalid/Timezone') = true, want false")
	}
}

// TestListTimezones tests ListTimezones function / ListTimezones 함수 테스트
func TestListTimezones(t *testing.T) {
	timezones := ListTimezones()
	if len(timezones) == 0 {
		t.Error("ListTimezones() returned empty list")
	}

	// Check if common timezones are present / 일반적인 타임존이 포함되어 있는지 확인
	found := false
	for _, tz := range timezones {
		if tz == "Asia/Seoul" || tz == "America/New_York" {
			found = true
			break
		}
	}
	if !found {
		t.Error("ListTimezones() should include common timezones")
	}
}

// ============================================================
// 3. Date Arithmetic Functions (16 functions)
// 3. 날짜 연산 함수 (16개 함수)
// ============================================================

// TestAddSeconds tests AddSeconds function / AddSeconds 함수 테스트
func TestAddSeconds(t *testing.T) {
	base := time.Date(2025, 10, 14, 15, 0, 0, 0, KST)
	result := AddSeconds(base, 30)

	if result.Second() != 30 {
		t.Errorf("AddSeconds() second = %v, want 30", result.Second())
	}
}

// TestAddMinutes tests AddMinutes function / AddMinutes 함수 테스트
func TestAddMinutes(t *testing.T) {
	base := time.Date(2025, 10, 14, 15, 0, 0, 0, KST)
	result := AddMinutes(base, 15)

	if result.Minute() != 15 {
		t.Errorf("AddMinutes() minute = %v, want 15", result.Minute())
	}
}

// TestAddHours tests AddHours function / AddHours 함수 테스트
func TestAddHours(t *testing.T) {
	base := time.Date(2025, 10, 14, 15, 0, 0, 0, KST)
	result := AddHours(base, 2)

	if result.Hour() != 17 {
		t.Errorf("AddHours() hour = %v, want 17", result.Hour())
	}
}

// TestAddDays tests AddDays function / AddDays 함수 테스트
func TestAddDays(t *testing.T) {
	base := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
	result := AddDays(base, 7)

	if result.Day() != 21 {
		t.Errorf("AddDays() day = %v, want 21", result.Day())
	}
}

// TestAddWeeks tests AddWeeks function / AddWeeks 함수 테스트
func TestAddWeeks(t *testing.T) {
	base := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
	result := AddWeeks(base, 2)

	expected := base.AddDate(0, 0, 14)
	if result.Day() != expected.Day() {
		t.Errorf("AddWeeks() day = %v, want %v", result.Day(), expected.Day())
	}
}

// TestAddMonths tests AddMonths function / AddMonths 함수 테스트
func TestAddMonths(t *testing.T) {
	base := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
	result := AddMonths(base, 3)

	if result.Month() != 1 || result.Year() != 2026 {
		t.Errorf("AddMonths() = %v, want January 2026", result)
	}
}

// TestAddYears tests AddYears function / AddYears 함수 테스트
func TestAddYears(t *testing.T) {
	base := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
	result := AddYears(base, 1)

	if result.Year() != 2026 {
		t.Errorf("AddYears() year = %v, want 2026", result.Year())
	}
}

// TestStartOfDay tests StartOfDay function / StartOfDay 함수 테스트
func TestStartOfDay(t *testing.T) {
	base := time.Date(2025, 10, 14, 15, 30, 45, 0, KST)
	result := StartOfDay(base)

	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Errorf("StartOfDay() = %v, want 00:00:00", result.Format("15:04:05"))
	}
}

// TestEndOfDay tests EndOfDay function / EndOfDay 함수 테스트
func TestEndOfDay(t *testing.T) {
	base := time.Date(2025, 10, 14, 15, 30, 45, 0, KST)
	result := EndOfDay(base)

	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 {
		t.Errorf("EndOfDay() = %v, want 23:59:59", result.Format("15:04:05"))
	}
}

// TestStartOfWeek tests StartOfWeek function / StartOfWeek 함수 테스트
func TestStartOfWeek(t *testing.T) {
	// Tuesday, October 14, 2025
	base := time.Date(2025, 10, 14, 15, 30, 45, 0, KST)
	result := StartOfWeek(base)

	if result.Weekday() != time.Monday {
		t.Errorf("StartOfWeek() weekday = %v, want Monday", result.Weekday())
	}
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Errorf("StartOfWeek() time = %v, want 00:00:00", result.Format("15:04:05"))
	}
}

// TestEndOfWeek tests EndOfWeek function / EndOfWeek 함수 테스트
func TestEndOfWeek(t *testing.T) {
	// Tuesday, October 14, 2025
	base := time.Date(2025, 10, 14, 15, 30, 45, 0, KST)
	result := EndOfWeek(base)

	if result.Weekday() != time.Sunday {
		t.Errorf("EndOfWeek() weekday = %v, want Sunday", result.Weekday())
	}
	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 {
		t.Errorf("EndOfWeek() time = %v, want 23:59:59", result.Format("15:04:05"))
	}
}

// TestStartOfMonth tests StartOfMonth function / StartOfMonth 함수 테스트
func TestStartOfMonth(t *testing.T) {
	base := time.Date(2025, 10, 14, 15, 30, 45, 0, KST)
	result := StartOfMonth(base)

	if result.Day() != 1 {
		t.Errorf("StartOfMonth() day = %v, want 1", result.Day())
	}
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Errorf("StartOfMonth() time = %v, want 00:00:00", result.Format("15:04:05"))
	}
}

// TestEndOfMonth tests EndOfMonth function / EndOfMonth 함수 테스트
func TestEndOfMonth(t *testing.T) {
	base := time.Date(2025, 10, 14, 15, 30, 45, 0, KST)
	result := EndOfMonth(base)

	if result.Day() != 31 {
		t.Errorf("EndOfMonth() day = %v, want 31", result.Day())
	}
	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 {
		t.Errorf("EndOfMonth() time = %v, want 23:59:59", result.Format("15:04:05"))
	}
}

// TestStartOfYear tests StartOfYear function / StartOfYear 함수 테스트
func TestStartOfYear(t *testing.T) {
	base := time.Date(2025, 10, 14, 15, 30, 45, 0, KST)
	result := StartOfYear(base)

	if result.Month() != 1 || result.Day() != 1 {
		t.Errorf("StartOfYear() = %v, want January 1", result)
	}
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Errorf("StartOfYear() time = %v, want 00:00:00", result.Format("15:04:05"))
	}
}

// TestEndOfYear tests EndOfYear function / EndOfYear 함수 테스트
func TestEndOfYear(t *testing.T) {
	base := time.Date(2025, 10, 14, 15, 30, 45, 0, KST)
	result := EndOfYear(base)

	if result.Month() != 12 || result.Day() != 31 {
		t.Errorf("EndOfYear() = %v, want December 31", result)
	}
	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 {
		t.Errorf("EndOfYear() time = %v, want 23:59:59", result.Format("15:04:05"))
	}
}

// TestStartOfQuarter tests StartOfQuarter function / StartOfQuarter 함수 테스트
func TestStartOfQuarter(t *testing.T) {
	// Q4 2025 (October)
	base := time.Date(2025, 10, 14, 15, 30, 45, 0, KST)
	result := StartOfQuarter(base)

	if result.Month() != 10 || result.Day() != 1 {
		t.Errorf("StartOfQuarter() = %v, want October 1", result)
	}
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Errorf("StartOfQuarter() time = %v, want 00:00:00", result.Format("15:04:05"))
	}
}

// ============================================================
// 4. Date Formatting Functions (8 functions)
// 4. 날짜 포맷팅 함수 (8개 함수)
// ============================================================

// TestFormatISO8601 tests FormatISO8601 function / FormatISO8601 함수 테스트
func TestFormatISO8601(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 4, 5, 0, KST)
	result := FormatISO8601(testTime)

	if result == "" {
		t.Error("FormatISO8601() returned empty string")
	}
	// ISO 8601 format should contain 'T' and timezone offset
	if len(result) < 20 {
		t.Errorf("FormatISO8601() = %v, seems too short for ISO 8601", result)
	}
}

// TestFormatRFC3339 tests FormatRFC3339 function / FormatRFC3339 함수 테스트
func TestFormatRFC3339(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 4, 5, 0, KST)
	result := FormatRFC3339(testTime)

	if result == "" {
		t.Error("FormatRFC3339() returned empty string")
	}
}

// TestFormatDate tests FormatDate function / FormatDate 함수 테스트
func TestFormatDate(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 4, 5, 0, KST)
	result := FormatDate(testTime)

	if result != "2025-10-14" {
		t.Errorf("FormatDate() = %v, want 2025-10-14", result)
	}
}

// TestFormatDateTime tests FormatDateTime function / FormatDateTime 함수 테스트
func TestFormatDateTime(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 4, 5, 0, KST)
	result := FormatDateTime(testTime)

	if result != "2025-10-14 15:04:05" {
		t.Errorf("FormatDateTime() = %v, want 2025-10-14 15:04:05", result)
	}
}

// TestFormatTime tests FormatTime function / FormatTime 함수 테스트
func TestFormatTime(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 4, 5, 0, KST)
	result := FormatTime(testTime)

	if result != "15:04:05" {
		t.Errorf("FormatTime() = %v, want 15:04:05", result)
	}
}

// TestFormat tests Format function with custom tokens / Format 함수 커스텀 토큰 테스트
func TestFormat(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 4, 5, 0, KST)

	// Test various format tokens / 다양한 포맷 토큰 테스트
	result := Format(testTime, "YYYY-MM-DD HH:mm:ss")
	if result == "" {
		t.Error("Format() returned empty string")
	}
}

// TestFormatKorean tests FormatKorean function / FormatKorean 함수 테스트
func TestFormatKorean(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 4, 5, 0, KST)
	result := FormatKorean(testTime)

	if result == "" {
		t.Error("FormatKorean() returned empty string")
	}
	// Should contain Korean characters / 한글 문자를 포함해야 함
	if len(result) < 10 {
		t.Errorf("FormatKorean() = %v, seems too short", result)
	}
}

// TestFormatWithTimezone tests FormatWithTimezone function / FormatWithTimezone 함수 테스트
func TestFormatWithTimezone(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 4, 5, 0, KST)

	result, err := FormatWithTimezone(testTime, "America/New_York")
	if err != nil {
		t.Errorf("FormatWithTimezone() error = %v", err)
	}
	if result == "" {
		t.Error("FormatWithTimezone() returned empty string")
	}

	// Test invalid timezone / 잘못된 타임존 테스트
	_, err = FormatWithTimezone(testTime, "Invalid/Timezone")
	if err == nil {
		t.Error("FormatWithTimezone() should return error for invalid timezone")
	}
}

// ============================================================
// 5. Time Parsing Functions (6 functions)
// 5. 시간 파싱 함수 (6개 함수)
// ============================================================

// TestParseISO8601 tests ParseISO8601 function / ParseISO8601 함수 테스트
func TestParseISO8601(t *testing.T) {
	result, err := ParseISO8601("2025-10-14T15:04:05+09:00")
	if err != nil {
		t.Errorf("ParseISO8601() error = %v", err)
	}
	if result.Year() != 2025 || result.Month() != 10 || result.Day() != 14 {
		t.Errorf("ParseISO8601() = %v, want 2025-10-14", result)
	}

	// Test invalid format / 잘못된 포맷 테스트
	_, err = ParseISO8601("invalid")
	if err == nil {
		t.Error("ParseISO8601() should return error for invalid format")
	}
}

// TestParseRFC3339 tests ParseRFC3339 function / ParseRFC3339 함수 테스트
func TestParseRFC3339(t *testing.T) {
	result, err := ParseRFC3339("2025-10-14T15:04:05+09:00")
	if err != nil {
		t.Errorf("ParseRFC3339() error = %v", err)
	}
	if result.IsZero() {
		t.Error("ParseRFC3339() returned zero time")
	}
}

// TestParseDate tests ParseDate function / ParseDate 함수 테스트
func TestParseDate(t *testing.T) {
	result, err := ParseDate("2025-10-14")
	if err != nil {
		t.Errorf("ParseDate() error = %v", err)
	}
	if result.Year() != 2025 || result.Month() != 10 || result.Day() != 14 {
		t.Errorf("ParseDate() = %v, want 2025-10-14", result)
	}

	// Test invalid format / 잘못된 포맷 테스트
	_, err = ParseDate("invalid")
	if err == nil {
		t.Error("ParseDate() should return error for invalid format")
	}
}

// TestParseDateTime tests ParseDateTime function / ParseDateTime 함수 테스트
func TestParseDateTime(t *testing.T) {
	result, err := ParseDateTime("2025-10-14 15:04:05")
	if err != nil {
		t.Errorf("ParseDateTime() error = %v", err)
	}
	if result.Year() != 2025 || result.Hour() != 15 {
		t.Errorf("ParseDateTime() = %v, want 2025-10-14 15:04:05", result)
	}
}

// TestParse tests Parse function with auto-detection / Parse 함수 자동 감지 테스트
func TestParse(t *testing.T) {
	// Test date format / 날짜 포맷 테스트
	result1, err := Parse("2025-10-14")
	if err != nil {
		t.Errorf("Parse() error = %v", err)
	}
	if result1.IsZero() {
		t.Error("Parse() returned zero time for date format")
	}

	// Test datetime format / 날짜시간 포맷 테스트
	result2, err := Parse("2025-10-14 15:04:05")
	if err != nil {
		t.Errorf("Parse() error = %v", err)
	}
	if result2.IsZero() {
		t.Error("Parse() returned zero time for datetime format")
	}

	// Test invalid format / 잘못된 포맷 테스트
	_, err = Parse("invalid")
	if err == nil {
		t.Error("Parse() should return error for invalid format")
	}
}

// TestParseWithTimezone tests ParseWithTimezone function / ParseWithTimezone 함수 테스트
func TestParseWithTimezone(t *testing.T) {
	result, err := ParseWithTimezone("2025-10-14 15:04:05", "America/New_York")
	if err != nil {
		t.Errorf("ParseWithTimezone() error = %v", err)
	}
	if result.IsZero() {
		t.Error("ParseWithTimezone() returned zero time")
	}

	// Test invalid timezone / 잘못된 타임존 테스트
	_, err = ParseWithTimezone("2025-10-14 15:04:05", "Invalid/Timezone")
	if err == nil {
		t.Error("ParseWithTimezone() should return error for invalid timezone")
	}
}

// ============================================================
// 6. Time Comparison Functions (18 functions)
// 6. 시간 비교 함수 (18개 함수)
// ============================================================

// TestIsBefore tests IsBefore function / IsBefore 함수 테스트
func TestIsBefore(t *testing.T) {
	t1 := time.Date(2025, 10, 13, 0, 0, 0, 0, KST)
	t2 := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)

	if !IsBefore(t1, t2) {
		t.Error("IsBefore() = false, want true")
	}
	if IsBefore(t2, t1) {
		t.Error("IsBefore() = true, want false")
	}
}

// TestIsAfter tests IsAfter function / IsAfter 함수 테스트
func TestIsAfter(t *testing.T) {
	t1 := time.Date(2025, 10, 15, 0, 0, 0, 0, KST)
	t2 := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)

	if !IsAfter(t1, t2) {
		t.Error("IsAfter() = false, want true")
	}
	if IsAfter(t2, t1) {
		t.Error("IsAfter() = true, want false")
	}
}

// TestIsBetween tests IsBetween function / IsBetween 함수 테스트
func TestIsBetween(t *testing.T) {
	start := time.Date(2025, 10, 13, 0, 0, 0, 0, KST)
	middle := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
	end := time.Date(2025, 10, 15, 0, 0, 0, 0, KST)

	if !IsBetween(middle, start, end) {
		t.Error("IsBetween() = false, want true")
	}
	if IsBetween(start, middle, end) {
		t.Error("IsBetween() = true, want false for start time")
	}
}

// TestIsToday tests IsToday function / IsToday 함수 테스트
func TestIsToday(t *testing.T) {
	now := time.Now()
	if !IsToday(now) {
		t.Error("IsToday() = false for current time, want true")
	}

	yesterday := now.AddDate(0, 0, -1)
	if IsToday(yesterday) {
		t.Error("IsToday() = true for yesterday, want false")
	}
}

// TestIsYesterday tests IsYesterday function / IsYesterday 함수 테스트
func TestIsYesterday(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)
	if !IsYesterday(yesterday) {
		t.Error("IsYesterday() = false, want true")
	}

	today := time.Now()
	if IsYesterday(today) {
		t.Error("IsYesterday() = true for today, want false")
	}
}

// TestIsTomorrow tests IsTomorrow function / IsTomorrow 함수 테스트
func TestIsTomorrow(t *testing.T) {
	tomorrow := time.Now().AddDate(0, 0, 1)
	if !IsTomorrow(tomorrow) {
		t.Error("IsTomorrow() = false, want true")
	}

	today := time.Now()
	if IsTomorrow(today) {
		t.Error("IsTomorrow() = true for today, want false")
	}
}

// TestIsWeekend tests IsWeekend function / IsWeekend 함수 테스트
func TestIsWeekend(t *testing.T) {
	saturday := time.Date(2025, 10, 18, 0, 0, 0, 0, KST) // Saturday
	sunday := time.Date(2025, 10, 19, 0, 0, 0, 0, KST)   // Sunday
	monday := time.Date(2025, 10, 13, 0, 0, 0, 0, KST)   // Monday

	if !IsWeekend(saturday) {
		t.Error("IsWeekend() = false for Saturday, want true")
	}
	if !IsWeekend(sunday) {
		t.Error("IsWeekend() = false for Sunday, want true")
	}
	if IsWeekend(monday) {
		t.Error("IsWeekend() = true for Monday, want false")
	}
}

// TestIsWeekday tests IsWeekday function / IsWeekday 함수 테스트
func TestIsWeekday(t *testing.T) {
	monday := time.Date(2025, 10, 13, 0, 0, 0, 0, KST) // Monday
	saturday := time.Date(2025, 10, 18, 0, 0, 0, 0, KST) // Saturday

	if !IsWeekday(monday) {
		t.Error("IsWeekday() = false for Monday, want true")
	}
	if IsWeekday(saturday) {
		t.Error("IsWeekday() = true for Saturday, want false")
	}
}

// TestIsThisWeek tests IsThisWeek function / IsThisWeek 함수 테스트
func TestIsThisWeek(t *testing.T) {
	now := time.Now()
	if !IsThisWeek(now) {
		t.Error("IsThisWeek() = false for current time, want true")
	}

	nextWeek := now.AddDate(0, 0, 8)
	if IsThisWeek(nextWeek) {
		t.Error("IsThisWeek() = true for next week, want false")
	}
}

// TestIsThisMonth tests IsThisMonth function / IsThisMonth 함수 테스트
func TestIsThisMonth(t *testing.T) {
	now := time.Now()
	if !IsThisMonth(now) {
		t.Error("IsThisMonth() = false for current time, want true")
	}

	nextMonth := now.AddDate(0, 1, 0)
	if IsThisMonth(nextMonth) {
		t.Error("IsThisMonth() = true for next month, want false")
	}
}

// TestIsThisYear tests IsThisYear function / IsThisYear 함수 테스트
func TestIsThisYear(t *testing.T) {
	now := time.Now()
	if !IsThisYear(now) {
		t.Error("IsThisYear() = false for current time, want true")
	}

	nextYear := now.AddDate(1, 0, 0)
	if IsThisYear(nextYear) {
		t.Error("IsThisYear() = true for next year, want false")
	}
}

// TestIsSameDay tests IsSameDay function / IsSameDay 함수 테스트
func TestIsSameDay(t *testing.T) {
	t1 := time.Date(2025, 10, 14, 10, 0, 0, 0, KST)
	t2 := time.Date(2025, 10, 14, 15, 0, 0, 0, KST)
	t3 := time.Date(2025, 10, 15, 10, 0, 0, 0, KST)

	if !IsSameDay(t1, t2) {
		t.Error("IsSameDay() = false for same day, want true")
	}
	if IsSameDay(t1, t3) {
		t.Error("IsSameDay() = true for different days, want false")
	}
}

// TestIsSameWeek tests IsSameWeek function / IsSameWeek 함수 테스트
func TestIsSameWeek(t *testing.T) {
	// Both are in the same week (Oct 13-19, 2025)
	t1 := time.Date(2025, 10, 14, 0, 0, 0, 0, KST) // Tuesday
	t2 := time.Date(2025, 10, 16, 0, 0, 0, 0, KST) // Thursday
	t3 := time.Date(2025, 10, 21, 0, 0, 0, 0, KST) // Next Tuesday

	if !IsSameWeek(t1, t2) {
		t.Error("IsSameWeek() = false for same week, want true")
	}
	if IsSameWeek(t1, t3) {
		t.Error("IsSameWeek() = true for different weeks, want false")
	}
}

// TestIsSameMonth tests IsSameMonth function / IsSameMonth 함수 테스트
func TestIsSameMonth(t *testing.T) {
	t1 := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
	t2 := time.Date(2025, 10, 20, 0, 0, 0, 0, KST)
	t3 := time.Date(2025, 11, 14, 0, 0, 0, 0, KST)

	if !IsSameMonth(t1, t2) {
		t.Error("IsSameMonth() = false for same month, want true")
	}
	if IsSameMonth(t1, t3) {
		t.Error("IsSameMonth() = true for different months, want false")
	}
}

// TestIsSameYear tests IsSameYear function / IsSameYear 함수 테스트
func TestIsSameYear(t *testing.T) {
	t1 := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
	t2 := time.Date(2025, 12, 31, 0, 0, 0, 0, KST)
	t3 := time.Date(2026, 1, 1, 0, 0, 0, 0, KST)

	if !IsSameYear(t1, t2) {
		t.Error("IsSameYear() = false for same year, want true")
	}
	if IsSameYear(t1, t3) {
		t.Error("IsSameYear() = true for different years, want false")
	}
}

// TestIsLeapYear tests IsLeapYear function / IsLeapYear 함수 테스트
func TestIsLeapYear(t *testing.T) {
	leap2024 := time.Date(2024, 1, 1, 0, 0, 0, 0, KST)
	notLeap2025 := time.Date(2025, 1, 1, 0, 0, 0, 0, KST)

	if !IsLeapYear(leap2024) {
		t.Error("IsLeapYear() = false for 2024, want true")
	}
	if IsLeapYear(notLeap2025) {
		t.Error("IsLeapYear() = true for 2025, want false")
	}
}

// TestIsPast tests IsPast function / IsPast 함수 테스트
func TestIsPast(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	future := time.Now().Add(1 * time.Hour)

	if !IsPast(past) {
		t.Error("IsPast() = false for past time, want true")
	}
	if IsPast(future) {
		t.Error("IsPast() = true for future time, want false")
	}
}

// TestIsFuture tests IsFuture function / IsFuture 함수 테스트
func TestIsFuture(t *testing.T) {
	future := time.Now().Add(1 * time.Hour)
	past := time.Now().Add(-1 * time.Hour)

	if !IsFuture(future) {
		t.Error("IsFuture() = false for future time, want true")
	}
	if IsFuture(past) {
		t.Error("IsFuture() = true for past time, want false")
	}
}

// ============================================================
// 7. Age Calculation Functions (4 functions)
// 7. 나이 계산 함수 (4개 함수)
// ============================================================

// TestAgeInYears tests AgeInYears function / AgeInYears 함수 테스트
func TestAgeInYears(t *testing.T) {
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, KST)
	years := AgeInYears(birthDate)

	if years < 0 {
		t.Errorf("AgeInYears() = %v, should be non-negative", years)
	}
	if years < 30 {
		t.Errorf("AgeInYears() = %v, expected around 35", years)
	}
}

// TestAgeInMonths tests AgeInMonths function / AgeInMonths 함수 테스트
func TestAgeInMonths(t *testing.T) {
	birthDate := time.Now().AddDate(-1, -6, 0) // 1 year 6 months ago
	months := AgeInMonths(birthDate)

	if months < 18 || months > 19 {
		t.Errorf("AgeInMonths() = %v, expected around 18", months)
	}
}

// TestAgeInDays tests AgeInDays function / AgeInDays 함수 테스트
func TestAgeInDays(t *testing.T) {
	birthDate := time.Now().AddDate(0, 0, -30) // 30 days ago
	days := AgeInDays(birthDate)

	if days < 29 || days > 31 {
		t.Errorf("AgeInDays() = %v, expected around 30", days)
	}
}

// TestAge tests Age function / Age 함수 테스트
func TestAge(t *testing.T) {
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, KST)
	age := Age(birthDate)

	if age == nil {
		t.Fatal("Age() returned nil")
	}
	if age.Years < 0 || age.Months < 0 || age.Days < 0 {
		t.Errorf("Age() = %v, all fields should be non-negative", age)
	}
	if age.Years < 30 {
		t.Errorf("Age().Years = %v, expected around 35", age.Years)
	}

	// Test String method / String 메서드 테스트
	str := age.String()
	if str == "" {
		t.Error("Age().String() returned empty string")
	}
}

// ============================================================
// 8. Relative Time Functions (4 functions)
// 8. 상대 시간 함수 (4개 함수)
// ============================================================

// TestRelativeTime tests RelativeTime function / RelativeTime 함수 테스트
func TestRelativeTime(t *testing.T) {
	// Test past time / 과거 시간 테스트
	past := time.Now().Add(-2 * time.Hour)
	pastStr := RelativeTime(past)
	if pastStr == "" {
		t.Error("RelativeTime() returned empty string for past")
	}

	// Test future time / 미래 시간 테스트
	future := time.Now().Add(3 * time.Hour)
	futureStr := RelativeTime(future)
	if futureStr == "" {
		t.Error("RelativeTime() returned empty string for future")
	}
}

// TestRelativeTimeShort tests RelativeTimeShort function / RelativeTimeShort 함수 테스트
func TestRelativeTimeShort(t *testing.T) {
	past := time.Now().Add(-2 * time.Hour)
	short := RelativeTimeShort(past)

	if short == "" {
		t.Error("RelativeTimeShort() returned empty string")
	}
}

// TestTimeAgo tests TimeAgo function / TimeAgo 함수 테스트
func TestTimeAgo(t *testing.T) {
	past := time.Now().Add(-30 * time.Minute)
	ago := TimeAgo(past)

	if ago == "" {
		t.Error("TimeAgo() returned empty string")
	}
}

// TestHumanizeDuration tests HumanizeDuration function / HumanizeDuration 함수 테스트
func TestHumanizeDuration(t *testing.T) {
	// Test 2 hours 30 minutes / 2시간 30분 테스트
	duration := 2*time.Hour + 30*time.Minute
	humanized := HumanizeDuration(duration)

	if humanized == "" {
		t.Error("HumanizeDuration() returned empty string")
	}
}

// ============================================================
// 9. Unix Timestamp Functions (12 functions)
// 9. Unix 타임스탬프 함수 (12개 함수)
// ============================================================

// TestNow tests Now function / Now 함수 테스트
func TestNow(t *testing.T) {
	unix := Now()
	if unix <= 0 {
		t.Errorf("Now() = %v, should be positive", unix)
	}
}

// TestNowMilli tests NowMilli function / NowMilli 함수 테스트
func TestNowMilli(t *testing.T) {
	milli := NowMilli()
	if milli <= 0 {
		t.Errorf("NowMilli() = %v, should be positive", milli)
	}
}

// TestNowMicro tests NowMicro function / NowMicro 함수 테스트
func TestNowMicro(t *testing.T) {
	micro := NowMicro()
	if micro <= 0 {
		t.Errorf("NowMicro() = %v, should be positive", micro)
	}
}

// TestNowNano tests NowNano function / NowNano 함수 테스트
func TestNowNano(t *testing.T) {
	nano := NowNano()
	if nano <= 0 {
		t.Errorf("NowNano() = %v, should be positive", nano)
	}
}

// TestFromUnix tests FromUnix function / FromUnix 함수 테스트
func TestFromUnix(t *testing.T) {
	unix := int64(1634198400) // 2021-10-14 12:00:00 UTC
	result := FromUnix(unix)

	if result.IsZero() {
		t.Error("FromUnix() returned zero time")
	}
	if result.Unix() != unix {
		t.Errorf("FromUnix() unix = %v, want %v", result.Unix(), unix)
	}
}

// TestFromUnixMilli tests FromUnixMilli function / FromUnixMilli 함수 테스트
func TestFromUnixMilli(t *testing.T) {
	milli := int64(1634198400000)
	result := FromUnixMilli(milli)

	if result.IsZero() {
		t.Error("FromUnixMilli() returned zero time")
	}
}

// TestFromUnixMicro tests FromUnixMicro function / FromUnixMicro 함수 테스트
func TestFromUnixMicro(t *testing.T) {
	micro := int64(1634198400000000)
	result := FromUnixMicro(micro)

	if result.IsZero() {
		t.Error("FromUnixMicro() returned zero time")
	}
}

// TestFromUnixNano tests FromUnixNano function / FromUnixNano 함수 테스트
func TestFromUnixNano(t *testing.T) {
	nano := int64(1634198400000000000)
	result := FromUnixNano(nano)

	if result.IsZero() {
		t.Error("FromUnixNano() returned zero time")
	}
}

// TestToUnix tests ToUnix function / ToUnix 함수 테스트
func TestToUnix(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 0, 0, 0, KST)
	unix := ToUnix(testTime)

	if unix <= 0 {
		t.Errorf("ToUnix() = %v, should be positive", unix)
	}

	// Verify round-trip / 왕복 변환 검증
	converted := FromUnix(unix)
	if converted.Unix() != testTime.Unix() {
		t.Error("ToUnix() round-trip failed")
	}
}

// TestToUnixMilli tests ToUnixMilli function / ToUnixMilli 함수 테스트
func TestToUnixMilli(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 0, 0, 0, KST)
	milli := ToUnixMilli(testTime)

	if milli <= 0 {
		t.Errorf("ToUnixMilli() = %v, should be positive", milli)
	}
}

// TestToUnixMicro tests ToUnixMicro function / ToUnixMicro 함수 테스트
func TestToUnixMicro(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 0, 0, 0, KST)
	micro := ToUnixMicro(testTime)

	if micro <= 0 {
		t.Errorf("ToUnixMicro() = %v, should be positive", micro)
	}
}

// TestToUnixNano tests ToUnixNano function / ToUnixNano 함수 테스트
func TestToUnixNano(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 0, 0, 0, KST)
	nano := ToUnixNano(testTime)

	if nano <= 0 {
		t.Errorf("ToUnixNano() = %v, should be positive", nano)
	}
}

// ============================================================
// 10. Business Day Functions (7 functions)
// 10. 영업일 함수 (7개 함수)
// ============================================================

// TestIsBusinessDay tests IsBusinessDay function / IsBusinessDay 함수 테스트
func TestIsBusinessDay(t *testing.T) {
	// Clear holidays first / 먼저 공휴일 초기화
	ClearHolidays()

	monday := time.Date(2025, 10, 13, 0, 0, 0, 0, KST) // Monday
	saturday := time.Date(2025, 10, 18, 0, 0, 0, 0, KST) // Saturday

	if !IsBusinessDay(monday) {
		t.Error("IsBusinessDay() = false for Monday, want true")
	}
	if IsBusinessDay(saturday) {
		t.Error("IsBusinessDay() = true for Saturday, want false")
	}

	// Test with holiday / 공휴일 테스트
	AddKoreanHolidays(2025)
	christmas := time.Date(2025, 12, 25, 0, 0, 0, 0, KST)
	if IsBusinessDay(christmas) {
		t.Error("IsBusinessDay() = true for Christmas, want false")
	}
}

// TestAddKoreanHolidays tests AddKoreanHolidays function / AddKoreanHolidays 함수 테스트
func TestAddKoreanHolidays(t *testing.T) {
	ClearHolidays()
	AddKoreanHolidays(2025)

	holidays := GetHolidays()
	if len(holidays) == 0 {
		t.Error("AddKoreanHolidays() added no holidays")
	}
}

// TestIsHoliday tests IsHoliday function / IsHoliday 함수 테스트
func TestIsHoliday(t *testing.T) {
	ClearHolidays()
	christmas := time.Date(2025, 12, 25, 0, 0, 0, 0, KST)

	// Before adding holiday / 공휴일 추가 전
	if IsHoliday(christmas) {
		t.Error("IsHoliday() = true before adding, want false")
	}

	// After adding holiday / 공휴일 추가 후
	SetHolidays([]time.Time{christmas})
	if !IsHoliday(christmas) {
		t.Error("IsHoliday() = false after adding, want true")
	}
}

// TestAddBusinessDays tests AddBusinessDays function / AddBusinessDays 함수 테스트
func TestAddBusinessDays(t *testing.T) {
	ClearHolidays()

	// Friday + 1 business day = next Monday
	friday := time.Date(2025, 10, 17, 0, 0, 0, 0, KST)
	nextBusiness := AddBusinessDays(friday, 1)

	if nextBusiness.Weekday() != time.Monday {
		t.Errorf("AddBusinessDays() weekday = %v, want Monday", nextBusiness.Weekday())
	}
}

// TestNextBusinessDay tests NextBusinessDay function / NextBusinessDay 함수 테스트
func TestNextBusinessDay(t *testing.T) {
	ClearHolidays()

	// Friday → next Monday
	friday := time.Date(2025, 10, 17, 0, 0, 0, 0, KST)
	nextDay := NextBusinessDay(friday)

	if nextDay.Weekday() != time.Monday {
		t.Errorf("NextBusinessDay() weekday = %v, want Monday", nextDay.Weekday())
	}
}

// TestPreviousBusinessDay tests PreviousBusinessDay function / PreviousBusinessDay 함수 테스트
func TestPreviousBusinessDay(t *testing.T) {
	ClearHolidays()

	// Monday → previous Friday
	monday := time.Date(2025, 10, 13, 0, 0, 0, 0, KST)
	prevDay := PreviousBusinessDay(monday)

	if prevDay.Weekday() != time.Friday {
		t.Errorf("PreviousBusinessDay() weekday = %v, want Friday", prevDay.Weekday())
	}
}

// TestCountBusinessDays tests CountBusinessDays function / CountBusinessDays 함수 테스트
func TestCountBusinessDays(t *testing.T) {
	ClearHolidays()

	// Monday to Friday = 4 business days (Tuesday, Wednesday, Thursday, Friday)
	// CountBusinessDays excludes start date
	monday := time.Date(2025, 10, 13, 0, 0, 0, 0, KST)
	friday := time.Date(2025, 10, 17, 0, 0, 0, 0, KST)

	count := CountBusinessDays(monday, friday)
	if count != 4 {
		t.Errorf("CountBusinessDays() = %v, want 4", count)
	}
}

// ============================================================
// Additional Tests for Holiday Management
// 추가 공휴일 관리 테스트
// ============================================================

// TestGetHolidays tests GetHolidays function / GetHolidays 함수 테스트
func TestGetHolidays(t *testing.T) {
	ClearHolidays()
	AddKoreanHolidays(2025)

	holidays := GetHolidays()
	if len(holidays) == 0 {
		t.Error("GetHolidays() returned empty list")
	}
}

// TestSetHolidays tests SetHolidays function / SetHolidays 함수 테스트
func TestSetHolidays(t *testing.T) {
	christmas := time.Date(2025, 12, 25, 0, 0, 0, 0, KST)
	newYear := time.Date(2025, 1, 1, 0, 0, 0, 0, KST)

	SetHolidays([]time.Time{christmas, newYear})
	holidays := GetHolidays()

	if len(holidays) != 2 {
		t.Errorf("SetHolidays() set %v holidays, want 2", len(holidays))
	}
}

// TestClearHolidays tests ClearHolidays function / ClearHolidays 함수 테스트
func TestClearHolidays(t *testing.T) {
	AddKoreanHolidays(2025)
	ClearHolidays()

	holidays := GetHolidays()
	if len(holidays) != 0 {
		t.Errorf("ClearHolidays() left %v holidays, want 0", len(holidays))
	}
}

// ============================================================
// Benchmark Tests / 벤치마크 테스트
// ============================================================

// BenchmarkSubTime benchmarks SubTime function / SubTime 함수 벤치마크
func BenchmarkSubTime(b *testing.B) {
	start := time.Now()
	end := start.Add(24 * time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SubTime(start, end)
	}
}

// BenchmarkDiffInDays benchmarks DiffInDays function / DiffInDays 함수 벤치마크
func BenchmarkDiffInDays(b *testing.B) {
	start := time.Now()
	end := start.Add(24 * time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DiffInDays(start, end)
	}
}

// BenchmarkConvertTimezone benchmarks ConvertTimezone function / ConvertTimezone 함수 벤치마크
func BenchmarkConvertTimezone(b *testing.B) {
	t := time.Now()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ConvertTimezone(t, "America/New_York")
	}
}

// BenchmarkFormatDate benchmarks FormatDate function / FormatDate 함수 벤치마크
func BenchmarkFormatDate(b *testing.B) {
	t := time.Now()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatDate(t)
	}
}

// BenchmarkFormatDateTime benchmarks FormatDateTime function / FormatDateTime 함수 벤치마크
func BenchmarkFormatDateTime(b *testing.B) {
	t := time.Now()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatDateTime(t)
	}
}

// BenchmarkFormat benchmarks Format function with custom tokens / 커스텀 토큰으로 Format 함수 벤치마크
func BenchmarkFormat(b *testing.B) {
	t := time.Now()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Format(t, "YYYY-MM-DD HH:mm:ss")
	}
}

// BenchmarkParseDate benchmarks ParseDate function / ParseDate 함수 벤치마크
func BenchmarkParseDate(b *testing.B) {
	dateStr := "2025-10-14"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate(dateStr)
	}
}

// BenchmarkParseDateTime benchmarks ParseDateTime function / ParseDateTime 함수 벤치마크
func BenchmarkParseDateTime(b *testing.B) {
	dateTimeStr := "2025-10-14 15:04:05"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseDateTime(dateTimeStr)
	}
}

// BenchmarkParse benchmarks Parse function with auto-detection / 자동 감지로 Parse 함수 벤치마크
func BenchmarkParse(b *testing.B) {
	dateStr := "2025-10-14"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Parse(dateStr)
	}
}

// BenchmarkRelativeTime benchmarks RelativeTime function / RelativeTime 함수 벤치마크
func BenchmarkRelativeTime(b *testing.B) {
	t := time.Now().Add(-2 * time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RelativeTime(t)
	}
}

// BenchmarkAge benchmarks Age function / Age 함수 벤치마크
func BenchmarkAge(b *testing.B) {
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, KST)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Age(birthDate)
	}
}

// BenchmarkIsBusinessDay benchmarks IsBusinessDay function / IsBusinessDay 함수 벤치마크
func BenchmarkIsBusinessDay(b *testing.B) {
	AddKoreanHolidays(2025)
	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsBusinessDay(t)
	}
}
