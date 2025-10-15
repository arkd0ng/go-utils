package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
	"github.com/arkd0ng/go-utils/timeutil"
)

func main() {
	// Setup log file with backup management / 백업 관리와 함께 로그 파일 설정
	logFilePath := "logs/timeutil-example.log"

	// Check if previous log file exists / 이전 로그 파일 존재 여부 확인
	if fileutil.Exists(logFilePath) {
		// Get modification time of existing log file / 기존 로그 파일의 수정 시간 가져오기
		modTime, err := fileutil.ModTime(logFilePath)
		if err == nil {
			// Create backup filename with timestamp / 타임스탬프와 함께 백업 파일명 생성
			backupName := fmt.Sprintf("logs/timeutil-example-%s.log", modTime.Format("20060102-150405"))

			// Backup existing log file / 기존 로그 파일 백업
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
				// Delete original log file to prevent content duplication / 내용 중복 방지를 위해 원본 로그 파일 삭제
				fileutil.DeleteFile(logFilePath)
			}
		}

		// Cleanup old backup files - keep only 5 most recent / 오래된 백업 파일 정리 - 최근 5개만 유지
		backupPattern := "logs/timeutil-example-*.log"
		backupFiles, err := filepath.Glob(backupPattern)
		if err == nil && len(backupFiles) > 5 {
			// Sort by modification time / 수정 시간으로 정렬
			type fileInfo struct {
				path    string
				modTime time.Time
			}
			var files []fileInfo
			for _, f := range backupFiles {
				if mt, err := fileutil.ModTime(f); err == nil {
					files = append(files, fileInfo{path: f, modTime: mt})
				}
			}

			// Sort oldest first / 가장 오래된 것부터 정렬
			for i := 0; i < len(files)-1; i++ {
				for j := i + 1; j < len(files); j++ {
					if files[i].modTime.After(files[j].modTime) {
						files[i], files[j] = files[j], files[i]
					}
				}
			}

			// Delete oldest files to keep only 5 / 5개만 유지하도록 가장 오래된 파일 삭제
			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Printf("🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	// Initialize logger with fixed filename / 고정 파일명으로 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithLevel(logging.DEBUG),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Printf("Failed to create logger: %v\n", err)
		return
	}
	defer logger.Close()

	// Print banner / 배너 출력
	logger.Banner("Timeutil Package Examples", "v1.6.006")
	logger.Info("Starting comprehensive timeutil examples with all 102 functions")
	logger.Info("시작: 모든 102개 함수를 포함한 포괄적인 timeutil 예제")

	// ============================================================
	// 1. Time Difference Functions (8 functions)
	// 1. 시간 차이 함수 (8개 함수)
	// ============================================================
	logger.Info("=" + repeat("=", 60))
	logger.Info("1. Time Difference Functions / 시간 차이 함수 (8 functions)")
	logger.Info("=" + repeat("=", 60))

	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 10, 15, 30, 45, 0, time.UTC)

	// SubTime - Returns TimeDiff object / TimeDiff 객체 반환
	diff := timeutil.SubTime(start, end)
	logger.Info("SubTime(start, end)", "result", diff.String())
	logger.Info("  → Days", "value", fmt.Sprintf("%.2f", diff.Days()))
	logger.Info("  → Humanized", "value", diff.Humanize())

	// DiffInSeconds - Seconds between two times / 두 시간 사이의 초
	seconds := timeutil.DiffInSeconds(start, end)
	logger.Info("DiffInSeconds(start, end)", "seconds", fmt.Sprintf("%.0f", seconds))

	// DiffInMinutes - Minutes between two times / 두 시간 사이의 분
	minutes := timeutil.DiffInMinutes(start, end)
	logger.Info("DiffInMinutes(start, end)", "minutes", fmt.Sprintf("%.0f", minutes))

	// DiffInHours - Hours between two times / 두 시간 사이의 시간
	hours := timeutil.DiffInHours(start, end)
	logger.Info("DiffInHours(start, end)", "hours", fmt.Sprintf("%.2f", hours))

	// DiffInDays - Days between two times / 두 시간 사이의 일
	days := timeutil.DiffInDays(start, end)
	logger.Info("DiffInDays(start, end)", "days", fmt.Sprintf("%.2f", days))

	// DiffInWeeks - Weeks between two times / 두 시간 사이의 주
	weeks := timeutil.DiffInWeeks(start, end)
	logger.Info("DiffInWeeks(start, end)", "weeks", fmt.Sprintf("%.2f", weeks))

	// DiffInMonths - Months between two times / 두 시간 사이의 월
	months := timeutil.DiffInMonths(start, end)
	logger.Info("DiffInMonths(start, end)", "months", months)

	// DiffInYears - Years between two times / 두 시간 사이의 년
	years := timeutil.DiffInYears(start, end)
	logger.Info("DiffInYears(start, end)", "years", years)

	// ============================================================
	// 2. Timezone Operations (10 functions)
	// 2. 타임존 작업 (10개 함수)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("2. Timezone Operations / 타임존 작업 (10 functions)")
	logger.Info("=" + repeat("=", 60))

	now := time.Now()

	// GetDefaultTimezone - Get current default timezone / 현재 기본 타임존 가져오기
	defaultTz := timeutil.GetDefaultTimezone()
	logger.Info("GetDefaultTimezone()", "timezone", defaultTz)

	// NowKST - Get current time in KST / KST로 현재 시간 가져오기
	kstNow := timeutil.NowKST()
	logger.Info("NowKST()", "time", timeutil.FormatDateTime(kstNow))

	// ConvertTimezone - Convert to different timezone / 다른 타임존으로 변환
	tokyoTime, _ := timeutil.ConvertTimezone(now, "Asia/Tokyo")
	logger.Info("ConvertTimezone(now, 'Asia/Tokyo')", "time", timeutil.FormatDateTime(tokyoTime))

	nyTime, _ := timeutil.ConvertTimezone(now, "America/New_York")
	logger.Info("ConvertTimezone(now, 'America/New_York')", "time", timeutil.FormatDateTime(nyTime))

	londonTime, _ := timeutil.ConvertTimezone(now, "Europe/London")
	logger.Info("ConvertTimezone(now, 'Europe/London')", "time", timeutil.FormatDateTime(londonTime))

	// ToUTC - Convert to UTC / UTC로 변환
	utcTime := timeutil.ToUTC(now)
	logger.Info("ToUTC(now)", "time", timeutil.FormatDateTime(utcTime))

	// ToKST - Convert to KST / KST로 변환
	kstTime := timeutil.ToKST(now)
	logger.Info("ToKST(now)", "time", timeutil.FormatDateTime(kstTime))

	// GetTimezoneOffset - Get timezone offset in seconds / 타임존 오프셋(초)
	offset, _ := timeutil.GetTimezoneOffset("Asia/Seoul")
	logger.Info("GetTimezoneOffset('Asia/Seoul')", "offset_seconds", offset, "offset_hours", offset/3600)

	// GetLocalTimezone - Get system's local timezone / 시스템의 로컬 타임존
	localTz := timeutil.GetLocalTimezone()
	logger.Info("GetLocalTimezone()", "timezone", localTz)

	// IsValidTimezone - Check if timezone is valid / 타임존이 유효한지 확인
	validTz := timeutil.IsValidTimezone("America/New_York")
	invalidTz := timeutil.IsValidTimezone("Invalid/Timezone")
	logger.Info("IsValidTimezone('America/New_York')", "valid", validTz)
	logger.Info("IsValidTimezone('Invalid/Timezone')", "valid", invalidTz)

	// ListTimezones - Get list of common timezones / 일반적인 타임존 목록
	timezones := timeutil.ListTimezones()
	logger.Info("ListTimezones()", "count", len(timezones), "first_5", fmt.Sprintf("%v", timezones[:5]))

	// ============================================================
	// 3. Date Arithmetic Functions (16 functions)
	// 3. 날짜 연산 함수 (16개 함수)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("3. Date Arithmetic Functions / 날짜 연산 함수 (16 functions)")
	logger.Info("=" + repeat("=", 60))

	baseTime := time.Date(2025, 10, 14, 15, 30, 45, 0, timeutil.KST)
	logger.Info("Base time", "time", timeutil.FormatDateTime(baseTime))

	// AddSeconds - Add seconds / 초 더하기
	logger.Info("AddSeconds(baseTime, 30)", "result", timeutil.FormatDateTime(timeutil.AddSeconds(baseTime, 30)))

	// AddMinutes - Add minutes / 분 더하기
	logger.Info("AddMinutes(baseTime, 15)", "result", timeutil.FormatDateTime(timeutil.AddMinutes(baseTime, 15)))

	// AddHours - Add hours / 시간 더하기
	logger.Info("AddHours(baseTime, 2)", "result", timeutil.FormatDateTime(timeutil.AddHours(baseTime, 2)))

	// AddDays - Add days / 일 더하기
	logger.Info("AddDays(baseTime, 7)", "result", timeutil.FormatDate(timeutil.AddDays(baseTime, 7)))

	// AddWeeks - Add weeks / 주 더하기
	logger.Info("AddWeeks(baseTime, 2)", "result", timeutil.FormatDate(timeutil.AddWeeks(baseTime, 2)))

	// AddMonths - Add months / 월 더하기
	logger.Info("AddMonths(baseTime, 3)", "result", timeutil.FormatDate(timeutil.AddMonths(baseTime, 3)))

	// AddYears - Add years / 년 더하기
	logger.Info("AddYears(baseTime, 1)", "result", timeutil.FormatDate(timeutil.AddYears(baseTime, 1)))

	// StartOfDay - Get start of day (00:00:00) / 하루의 시작 (00:00:00)
	startOfDay := timeutil.StartOfDay(baseTime)
	logger.Info("StartOfDay(baseTime)", "result", timeutil.FormatDateTime(startOfDay))

	// EndOfDay - Get end of day (23:59:59) / 하루의 끝 (23:59:59)
	endOfDay := timeutil.EndOfDay(baseTime)
	logger.Info("EndOfDay(baseTime)", "result", timeutil.FormatDateTime(endOfDay))

	// StartOfWeek - Get start of week (Monday) / 주의 시작 (월요일)
	startOfWeek := timeutil.StartOfWeek(baseTime)
	logger.Info("StartOfWeek(baseTime)", "result", timeutil.FormatDateTime(startOfWeek))

	// EndOfWeek - Get end of week (Sunday) / 주의 끝 (일요일)
	endOfWeek := timeutil.EndOfWeek(baseTime)
	logger.Info("EndOfWeek(baseTime)", "result", timeutil.FormatDateTime(endOfWeek))

	// StartOfMonth - Get start of month / 월의 시작
	startOfMonth := timeutil.StartOfMonth(baseTime)
	logger.Info("StartOfMonth(baseTime)", "result", timeutil.FormatDateTime(startOfMonth))

	// EndOfMonth - Get end of month / 월의 끝
	endOfMonth := timeutil.EndOfMonth(baseTime)
	logger.Info("EndOfMonth(baseTime)", "result", timeutil.FormatDateTime(endOfMonth))

	// StartOfYear - Get start of year / 년의 시작
	startOfYear := timeutil.StartOfYear(baseTime)
	logger.Info("StartOfYear(baseTime)", "result", timeutil.FormatDateTime(startOfYear))

	// EndOfYear - Get end of year / 년의 끝
	endOfYear := timeutil.EndOfYear(baseTime)
	logger.Info("EndOfYear(baseTime)", "result", timeutil.FormatDateTime(endOfYear))

	// StartOfQuarter - Get start of quarter / 분기의 시작
	startOfQuarter := timeutil.StartOfQuarter(baseTime)
	logger.Info("StartOfQuarter(baseTime)", "result", timeutil.FormatDateTime(startOfQuarter))

	// ============================================================
	// 4. Date Formatting Functions (13 functions)
	// 4. 날짜 포맷팅 함수 (13개 함수)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("4. Date Formatting Functions / 날짜 포맷팅 함수 (13 functions)")
	logger.Info("=" + repeat("=", 60))

	formatTime := time.Date(2025, 10, 14, 15, 4, 5, 0, timeutil.KST)
	logger.Info("Format time", "time", formatTime.String())

	// FormatISO8601 - Format as ISO 8601 (YYYY-MM-DD) / ISO 8601 포맷
	logger.Info("FormatISO8601(formatTime)", "result", timeutil.FormatISO8601(formatTime))

	// FormatRFC3339 - Format as RFC 3339 / RFC 3339 포맷
	logger.Info("FormatRFC3339(formatTime)", "result", timeutil.FormatRFC3339(formatTime))

	// FormatDate - Format date only / 날짜만 포맷
	logger.Info("FormatDate(formatTime)", "result", timeutil.FormatDate(formatTime))

	// FormatDateTime - Format date and time / 날짜와 시간 포맷
	logger.Info("FormatDateTime(formatTime)", "result", timeutil.FormatDateTime(formatTime))

	// FormatTime - Format time only / 시간만 포맷
	logger.Info("FormatTime(formatTime)", "result", timeutil.FormatTime(formatTime))

	// Format - Custom format with tokens / 토큰으로 커스텀 포맷
	logger.Info("Format(formatTime, 'YYYY-MM-DD HH:mm:ss')", "result", timeutil.Format(formatTime, "YYYY-MM-DD HH:mm:ss"))
	logger.Info("Format(formatTime, 'YYYY/MM/DD')", "result", timeutil.Format(formatTime, "YYYY/MM/DD"))
	logger.Info("Format(formatTime, 'DD-MM-YYYY')", "result", timeutil.Format(formatTime, "DD-MM-YYYY"))

	// FormatKorean - Format in Korean style / 한국어 스타일 포맷
	logger.Info("FormatKorean(formatTime)", "result", timeutil.FormatKorean(formatTime))

	// WeekdayKorean - Get Korean weekday name / 한글 요일 이름 가져오기
	// Returns full Korean weekday name (e.g., "월요일", "화요일")
	// 전체 한글 요일 이름을 반환합니다 (예: "월요일", "화요일")
	weekdayKorean := timeutil.WeekdayKorean(formatTime)
	logger.Info("WeekdayKorean(formatTime)", "result", weekdayKorean)

	// WeekdayKoreanShort - Get short Korean weekday name / 짧은 한글 요일 이름 가져오기
	// Returns short Korean weekday name (e.g., "월", "화")
	// 짧은 한글 요일 이름을 반환합니다 (예: "월", "화")
	weekdayShort := timeutil.WeekdayKoreanShort(formatTime)
	logger.Info("WeekdayKoreanShort(formatTime)", "result", weekdayShort)

	// FormatKoreanDateTime - Format date/time with Korean weekday / 한글 요일이 포함된 날짜/시간 포맷
	// Format: YYYY년 MM월 DD일 (요일) HH시 mm분 ss초
	// Expected output: "2025년 10월 14일 (화요일) 15시 04분 05초"
	koreanDateTime := timeutil.FormatKoreanDateTime(formatTime)
	logger.Info("FormatKoreanDateTime(formatTime)", "result", koreanDateTime)

	// FormatKoreanDateWithWeekday - Format date with full Korean weekday / 전체 한글 요일이 포함된 날짜 포맷
	// Format: YYYY년 MM월 DD일 (요일)
	// Expected output: "2025년 10월 14일 (화요일)"
	koreanDateFull := timeutil.FormatKoreanDateWithWeekday(formatTime)
	logger.Info("FormatKoreanDateWithWeekday(formatTime)", "result", koreanDateFull)

	// FormatKoreanDateShort - Format date with short Korean weekday / 짧은 한글 요일이 포함된 날짜 포맷
	// Format: YYYY년 MM월 DD일 (요일)
	// Expected output: "2025년 10월 14일 (화)"
	koreanDateShort := timeutil.FormatKoreanDateShort(formatTime)
	logger.Info("FormatKoreanDateShort(formatTime)", "result", koreanDateShort)

	// FormatWithTimezone - Format with specific timezone / 특정 타임존으로 포맷
	formatted, _ := timeutil.FormatWithTimezone(formatTime, "America/New_York")
	logger.Info("FormatWithTimezone(formatTime, 'America/New_York')", "result", formatted)

	// ============================================================
	// 5. Time Parsing Functions (6 functions)
	// 5. 시간 파싱 함수 (6개 함수)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("5. Time Parsing Functions / 시간 파싱 함수 (6 functions)")
	logger.Info("=" + repeat("=", 60))

	// ParseISO8601 - Parse ISO 8601 format / ISO 8601 포맷 파싱
	parsed1, _ := timeutil.ParseISO8601("2025-10-14")
	logger.Info("ParseISO8601('2025-10-14')", "result", timeutil.FormatDateTime(parsed1))

	// ParseRFC3339 - Parse RFC 3339 format / RFC 3339 포맷 파싱
	parsed2, _ := timeutil.ParseRFC3339("2025-10-14T15:04:05+09:00")
	logger.Info("ParseRFC3339('2025-10-14T15:04:05+09:00')", "result", timeutil.FormatDateTime(parsed2))

	// ParseDate - Parse date string / 날짜 문자열 파싱
	parsed3, _ := timeutil.ParseDate("2025-10-14")
	logger.Info("ParseDate('2025-10-14')", "result", timeutil.FormatDateTime(parsed3))

	// ParseDateTime - Parse datetime string / 날짜시간 문자열 파싱
	parsed4, _ := timeutil.ParseDateTime("2025-10-14 15:04:05")
	logger.Info("ParseDateTime('2025-10-14 15:04:05')", "result", timeutil.FormatDateTime(parsed4))

	// Parse - Auto-detect format and parse / 자동 포맷 감지 및 파싱
	parsed5, _ := timeutil.Parse("2025-10-14")
	logger.Info("Parse('2025-10-14') - auto-detect", "result", timeutil.FormatDateTime(parsed5))

	parsed6, _ := timeutil.Parse("2025-10-14 15:04:05")
	logger.Info("Parse('2025-10-14 15:04:05') - auto-detect", "result", timeutil.FormatDateTime(parsed6))

	// ParseWithTimezone - Parse with specific timezone / 특정 타임존으로 파싱
	parsed7, _ := timeutil.ParseWithTimezone("2025-10-14 15:04:05", "America/New_York")
	logger.Info("ParseWithTimezone('2025-10-14 15:04:05', 'America/New_York')", "result", timeutil.FormatDateTime(parsed7))

	// ============================================================
	// 6. Time Comparison Functions (18 functions)
	// 6. 시간 비교 함수 (18개 함수)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("6. Time Comparison Functions / 시간 비교 함수 (18 functions)")
	logger.Info("=" + repeat("=", 60))

	compareTime := time.Now()
	yesterday := timeutil.AddDays(compareTime, -1)
	tomorrow := timeutil.AddDays(compareTime, 1)

	// IsBefore - Check if time is before another / 이전인지 확인
	logger.Info("IsBefore(yesterday, compareTime)", "result", timeutil.IsBefore(yesterday, compareTime))

	// IsAfter - Check if time is after another / 이후인지 확인
	logger.Info("IsAfter(tomorrow, compareTime)", "result", timeutil.IsAfter(tomorrow, compareTime))

	// IsBetween - Check if time is between two times / 두 시간 사이인지 확인
	logger.Info("IsBetween(compareTime, yesterday, tomorrow)", "result", timeutil.IsBetween(compareTime, yesterday, tomorrow))

	// IsToday - Check if time is today / 오늘인지 확인
	logger.Info("IsToday(compareTime)", "result", timeutil.IsToday(compareTime))
	logger.Info("IsToday(yesterday)", "result", timeutil.IsToday(yesterday))

	// IsYesterday - Check if time is yesterday / 어제인지 확인
	logger.Info("IsYesterday(yesterday)", "result", timeutil.IsYesterday(yesterday))

	// IsTomorrow - Check if time is tomorrow / 내일인지 확인
	logger.Info("IsTomorrow(tomorrow)", "result", timeutil.IsTomorrow(tomorrow))

	// IsWeekend - Check if time is weekend / 주말인지 확인
	saturday := time.Date(2025, 10, 18, 0, 0, 0, 0, timeutil.KST) // Saturday
	logger.Info("IsWeekend(saturday)", "result", timeutil.IsWeekend(saturday))
	logger.Info("IsWeekend(compareTime)", "result", timeutil.IsWeekend(compareTime))

	// IsWeekday - Check if time is weekday / 평일인지 확인
	monday := time.Date(2025, 10, 13, 0, 0, 0, 0, timeutil.KST) // Monday
	logger.Info("IsWeekday(monday)", "result", timeutil.IsWeekday(monday))

	// IsThisWeek - Check if time is this week / 이번 주인지 확인
	logger.Info("IsThisWeek(compareTime)", "result", timeutil.IsThisWeek(compareTime))

	// IsThisMonth - Check if time is this month / 이번 달인지 확인
	logger.Info("IsThisMonth(compareTime)", "result", timeutil.IsThisMonth(compareTime))

	// IsThisYear - Check if time is this year / 올해인지 확인
	logger.Info("IsThisYear(compareTime)", "result", timeutil.IsThisYear(compareTime))

	// IsSameDay - Check if two times are same day / 같은 날인지 확인
	logger.Info("IsSameDay(compareTime, compareTime)", "result", timeutil.IsSameDay(compareTime, compareTime))
	logger.Info("IsSameDay(compareTime, yesterday)", "result", timeutil.IsSameDay(compareTime, yesterday))

	// IsSameWeek - Check if two times are same week / 같은 주인지 확인
	logger.Info("IsSameWeek(compareTime, yesterday)", "result", timeutil.IsSameWeek(compareTime, yesterday))

	// IsSameMonth - Check if two times are same month / 같은 달인지 확인
	logger.Info("IsSameMonth(compareTime, yesterday)", "result", timeutil.IsSameMonth(compareTime, yesterday))

	// IsSameYear - Check if two times are same year / 같은 년인지 확인
	logger.Info("IsSameYear(compareTime, yesterday)", "result", timeutil.IsSameYear(compareTime, yesterday))

	// IsLeapYear - Check if year is leap year / 윤년인지 확인
	leapYear := time.Date(2024, 1, 1, 0, 0, 0, 0, timeutil.KST)
	nonLeapYear := time.Date(2025, 1, 1, 0, 0, 0, 0, timeutil.KST)
	logger.Info("IsLeapYear(2024)", "result", timeutil.IsLeapYear(leapYear))
	logger.Info("IsLeapYear(2025)", "result", timeutil.IsLeapYear(nonLeapYear))

	// IsPast - Check if time is in the past / 과거인지 확인
	logger.Info("IsPast(yesterday)", "result", timeutil.IsPast(yesterday))
	logger.Info("IsPast(tomorrow)", "result", timeutil.IsPast(tomorrow))

	// IsFuture - Check if time is in the future / 미래인지 확인
	logger.Info("IsFuture(tomorrow)", "result", timeutil.IsFuture(tomorrow))
	logger.Info("IsFuture(yesterday)", "result", timeutil.IsFuture(yesterday))

	// ============================================================
	// 7. Age Calculation Functions (4 functions)
	// 7. 나이 계산 함수 (4개 함수)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("7. Age Calculation Functions / 나이 계산 함수 (4 functions)")
	logger.Info("=" + repeat("=", 60))

	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, timeutil.KST)
	logger.Info("Birth date", "date", timeutil.FormatDate(birthDate))

	// AgeInYears - Get age in years / 년 단위 나이
	ageYears := timeutil.AgeInYears(birthDate)
	logger.Info("AgeInYears(birthDate)", "years", ageYears)

	// AgeInMonths - Get age in months / 월 단위 나이
	ageMonths := timeutil.AgeInMonths(birthDate)
	logger.Info("AgeInMonths(birthDate)", "months", ageMonths)

	// AgeInDays - Get age in days / 일 단위 나이
	ageDays := timeutil.AgeInDays(birthDate)
	logger.Info("AgeInDays(birthDate)", "days", ageDays)

	// Age - Get detailed age / 상세 나이
	age := timeutil.Age(birthDate)
	logger.Info("Age(birthDate)", "age", age.String(), "years", age.Years, "months", age.Months, "days", age.Days)

	// ============================================================
	// 8. Relative Time Functions (4 functions)
	// 8. 상대 시간 함수 (4개 함수)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("8. Relative Time Functions / 상대 시간 함수 (4 functions)")
	logger.Info("=" + repeat("=", 60))

	// Test various past times / 다양한 과거 시간 테스트
	past30Sec := time.Now().Add(-30 * time.Second)
	past2Hours := time.Now().Add(-2 * time.Hour)
	past3Days := time.Now().Add(-3 * 24 * time.Hour)
	past2Weeks := time.Now().Add(-2 * 7 * 24 * time.Hour)
	past3Months := time.Now().Add(-3 * 30 * 24 * time.Hour)

	// RelativeTime - Get relative time string (long format) / 상대 시간 문자열 (긴 포맷)
	logger.Info("RelativeTime(30 seconds ago)", "result", timeutil.RelativeTime(past30Sec))
	logger.Info("RelativeTime(2 hours ago)", "result", timeutil.RelativeTime(past2Hours))
	logger.Info("RelativeTime(3 days ago)", "result", timeutil.RelativeTime(past3Days))
	logger.Info("RelativeTime(2 weeks ago)", "result", timeutil.RelativeTime(past2Weeks))
	logger.Info("RelativeTime(3 months ago)", "result", timeutil.RelativeTime(past3Months))

	// Test future times / 미래 시간 테스트
	future5Min := time.Now().Add(5 * time.Minute)
	future3Hours := time.Now().Add(3 * time.Hour)
	future2Days := time.Now().Add(2 * 24 * time.Hour)

	logger.Info("RelativeTime(in 5 minutes)", "result", timeutil.RelativeTime(future5Min))
	logger.Info("RelativeTime(in 3 hours)", "result", timeutil.RelativeTime(future3Hours))
	logger.Info("RelativeTime(in 2 days)", "result", timeutil.RelativeTime(future2Days))

	// RelativeTimeShort - Get relative time string (short format) / 상대 시간 문자열 (짧은 포맷)
	logger.Info("RelativeTimeShort(2 hours ago)", "result", timeutil.RelativeTimeShort(past2Hours))
	logger.Info("RelativeTimeShort(3 days ago)", "result", timeutil.RelativeTimeShort(past3Days))
	logger.Info("RelativeTimeShort(in 3 hours)", "result", timeutil.RelativeTimeShort(future3Hours))

	// TimeAgo - Alias for RelativeTime / RelativeTime의 별칭
	logger.Info("TimeAgo(2 hours ago)", "result", timeutil.TimeAgo(past2Hours))

	// HumanizeDuration - Convert duration to human-readable string / 기간을 사람이 읽기 쉬운 문자열로
	duration := 2*time.Hour + 30*time.Minute + 45*time.Second
	logger.Info("HumanizeDuration(2h 30m 45s)", "result", timeutil.HumanizeDuration(duration))

	// ============================================================
	// 9. Unix Timestamp Functions (12 functions)
	// 9. Unix 타임스탬프 함수 (12개 함수)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("9. Unix Timestamp Functions / Unix 타임스탬프 함수 (12 functions)")
	logger.Info("=" + repeat("=", 60))

	// Now - Get current Unix timestamp (seconds) / 현재 Unix 타임스탬프 (초)
	unixNow := timeutil.Now()
	logger.Info("Now()", "unix_seconds", unixNow)

	// NowMilli - Get current Unix timestamp (milliseconds) / 현재 Unix 타임스탬프 (밀리초)
	unixMilli := timeutil.NowMilli()
	logger.Info("NowMilli()", "unix_milliseconds", unixMilli)

	// NowMicro - Get current Unix timestamp (microseconds) / 현재 Unix 타임스탬프 (마이크로초)
	unixMicro := timeutil.NowMicro()
	logger.Info("NowMicro()", "unix_microseconds", unixMicro)

	// NowNano - Get current Unix timestamp (nanoseconds) / 현재 Unix 타임스탬프 (나노초)
	unixNano := timeutil.NowNano()
	logger.Info("NowNano()", "unix_nanoseconds", unixNano)

	// FromUnix - Convert Unix timestamp (seconds) to time.Time / Unix 타임스탬프(초)를 time.Time으로
	fromUnix := timeutil.FromUnix(1634198400)
	logger.Info("FromUnix(1634198400)", "result", timeutil.FormatDateTime(fromUnix))

	// FromUnixMilli - Convert Unix timestamp (milliseconds) to time.Time / Unix 타임스탬프(밀리초)를 time.Time으로
	fromUnixMilli := timeutil.FromUnixMilli(1634198400123)
	logger.Info("FromUnixMilli(1634198400123)", "result", timeutil.FormatDateTime(fromUnixMilli))

	// FromUnixMicro - Convert Unix timestamp (microseconds) to time.Time / Unix 타임스탬프(마이크로초)를 time.Time으로
	fromUnixMicro := timeutil.FromUnixMicro(1634198400123456)
	logger.Info("FromUnixMicro(1634198400123456)", "result", timeutil.FormatDateTime(fromUnixMicro))

	// FromUnixNano - Convert Unix timestamp (nanoseconds) to time.Time / Unix 타임스탬프(나노초)를 time.Time으로
	fromUnixNano := timeutil.FromUnixNano(1634198400123456789)
	logger.Info("FromUnixNano(1634198400123456789)", "result", timeutil.FormatDateTime(fromUnixNano))

	// ToUnix - Convert time.Time to Unix timestamp (seconds) / time.Time을 Unix 타임스탬프(초)로
	testTimeForUnix := time.Date(2025, 10, 14, 15, 4, 5, 0, timeutil.KST)
	toUnix := timeutil.ToUnix(testTimeForUnix)
	logger.Info("ToUnix(2025-10-14 15:04:05)", "unix_seconds", toUnix)

	// ToUnixMilli - Convert time.Time to Unix timestamp (milliseconds) / time.Time을 Unix 타임스탬프(밀리초)로
	toUnixMilli := timeutil.ToUnixMilli(testTimeForUnix)
	logger.Info("ToUnixMilli(2025-10-14 15:04:05)", "unix_milliseconds", toUnixMilli)

	// ToUnixMicro - Convert time.Time to Unix timestamp (microseconds) / time.Time을 Unix 타임스탬프(마이크로초)로
	toUnixMicro := timeutil.ToUnixMicro(testTimeForUnix)
	logger.Info("ToUnixMicro(2025-10-14 15:04:05)", "unix_microseconds", toUnixMicro)

	// ToUnixNano - Convert time.Time to Unix timestamp (nanoseconds) / time.Time을 Unix 타임스탬프(나노초)로
	toUnixNano := timeutil.ToUnixNano(testTimeForUnix)
	logger.Info("ToUnixNano(2025-10-14 15:04:05)", "unix_nanoseconds", toUnixNano)

	// ============================================================
	// 10. Business Day Functions (7 functions + holiday management)
	// 10. 영업일 함수 (7개 함수 + 공휴일 관리)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("10. Business Day Functions / 영업일 함수 (7 functions)")
	logger.Info("=" + repeat("=", 60))

	bizMonday := time.Date(2025, 10, 13, 0, 0, 0, 0, timeutil.KST)   // Monday
	bizSaturday := time.Date(2025, 10, 18, 0, 0, 0, 0, timeutil.KST) // Saturday
	newYearDay := time.Date(2025, 1, 1, 0, 0, 0, 0, timeutil.KST)    // New Year

	// IsBusinessDay - Check if day is business day / 영업일인지 확인
	logger.Info("IsBusinessDay(Monday)", "result", timeutil.IsBusinessDay(bizMonday))
	logger.Info("IsBusinessDay(Saturday)", "result", timeutil.IsBusinessDay(bizSaturday))

	// AddKoreanHolidays - Add Korean public holidays / 한국 공휴일 추가
	timeutil.AddKoreanHolidays(2025)
	logger.Info("AddKoreanHolidays(2025)", "status", "completed")

	// IsHoliday - Check if day is holiday / 공휴일인지 확인
	logger.Info("IsHoliday(Jan 1, 2025)", "result", timeutil.IsHoliday(newYearDay))
	logger.Info("IsHoliday(Monday Oct 13)", "result", timeutil.IsHoliday(bizMonday))

	// IsBusinessDay after adding holidays / 공휴일 추가 후 영업일 확인
	logger.Info("IsBusinessDay(Jan 1, 2025) - after adding holidays", "result", timeutil.IsBusinessDay(newYearDay))

	// AddBusinessDays - Add business days / 영업일 더하기
	nextBiz1 := timeutil.AddBusinessDays(bizMonday, 1)
	logger.Info("AddBusinessDays(Monday, 1)", "result", timeutil.FormatDate(nextBiz1))

	nextBiz5 := timeutil.AddBusinessDays(bizMonday, 5)
	logger.Info("AddBusinessDays(Monday, 5)", "result", timeutil.FormatDate(nextBiz5))

	// NextBusinessDay - Get next business day / 다음 영업일
	nextBiz := timeutil.NextBusinessDay(bizSaturday)
	logger.Info("NextBusinessDay(Saturday)", "result", timeutil.FormatDate(nextBiz))

	// PreviousBusinessDay - Get previous business day / 이전 영업일
	prevBiz := timeutil.PreviousBusinessDay(bizMonday)
	logger.Info("PreviousBusinessDay(Monday)", "result", timeutil.FormatDate(prevBiz))

	// CountBusinessDays - Count business days between two dates / 두 날짜 사이의 영업일 수
	startDate := time.Date(2025, 10, 13, 0, 0, 0, 0, timeutil.KST) // Monday
	endDate := time.Date(2025, 10, 24, 0, 0, 0, 0, timeutil.KST)   // Friday (2 weeks later)
	bizDayCount := timeutil.CountBusinessDays(startDate, endDate)
	logger.Info("CountBusinessDays(Oct 13 ~ Oct 24)", "business_days", bizDayCount)

	// GetHolidays - Get list of registered holidays / 등록된 공휴일 목록
	holidays := timeutil.GetHolidays()
	logger.Info("GetHolidays()", "count", len(holidays))

	// SetHolidays - Set custom holidays / 커스텀 공휴일 설정
	customHolidays := []time.Time{
		time.Date(2025, 12, 25, 0, 0, 0, 0, timeutil.KST), // Christmas
	}
	timeutil.SetHolidays(customHolidays)
	logger.Info("SetHolidays([Christmas])", "status", "completed")

	// ClearHolidays - Clear all holidays / 모든 공휴일 지우기
	timeutil.ClearHolidays()
	logger.Info("ClearHolidays()", "status", "completed")

	// Re-add Korean holidays for final demo / 최종 데모를 위해 한국 공휴일 재추가
	timeutil.AddKoreanHolidays(2025)
	logger.Info("AddKoreanHolidays(2025) - restored", "status", "completed")

	// ============================================================
	// Summary / 요약
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("Summary / 요약")
	logger.Info("=" + repeat("=", 60))

	logger.Info("All 102 timeutil functions demonstrated successfully!")
	logger.Info("모든 102개 timeutil 함수를 성공적으로 시연했습니다!")
	logger.Info("")
	logger.Info("Categories covered / 다뤄진 카테고리:")
	logger.Info("  1. Time Difference (8 functions) / 시간 차이 (8개 함수)")
	logger.Info("  2. Timezone Operations (10 functions) / 타임존 작업 (10개 함수)")
	logger.Info("  3. Date Arithmetic (16 functions) / 날짜 연산 (16개 함수)")
	logger.Info("  4. Date Formatting (13 functions) / 날짜 포맷팅 (13개 함수)")
	logger.Info("  5. Time Parsing (6 functions) / 시간 파싱 (6개 함수)")
	logger.Info("  6. Time Comparisons (18 functions) / 시간 비교 (18개 함수)")
	logger.Info("  7. Age Calculations (4 functions) / 나이 계산 (4개 함수)")
	logger.Info("  8. Relative Time (4 functions) / 상대 시간 (4개 함수)")
	logger.Info("  9. Unix Timestamp (12 functions) / Unix 타임스탬프 (12개 함수)")
	logger.Info("  10. Business Days (7 functions) / 영업일 (7개 함수)")
	// ============================================================
	// 13. String Parameter Functions (New in v1.6.008!)
	// 13. 문자열 매개변수 함수 (v1.6.008 신규!)
	// ============================================================
	logger.Info("")
	logger.Info("=" + repeat("=", 60))
	logger.Info("13. String Parameter Functions / 문자열 매개변수 함수")
	logger.Info("    (NEW in v1.6.008 - Parse any format automatically!)")
	logger.Info("    (v1.6.008 신규 - 모든 포맷 자동 파싱!)")
	logger.Info("=" + repeat("=", 60))

	// ParseAny - Automatically detect and parse any time format / 모든 시간 포맷 자동 감지 및 파싱
	logger.Info("")
	logger.Info("--- ParseAny: Automatic Format Detection ---")
	logger.Info("--- ParseAny: 자동 포맷 감지 ---")

	formats := []string{
		"2024-10-04 08:34:42.324",   // MySQL with milliseconds
		"2024-10-04T08:34:42+09:00", // ISO8601
		"2024/10/04",                // Date with slashes
		"Oct 04, 2024",              // Month name
		"2024-10-04 08:34:42",       // Standard DateTime
	}

	for _, f := range formats {
		parsed, err := timeutil.ParseAny(f)
		if err == nil {
			logger.Info("ParseAny", "input", f, "result", parsed.Format("2006-01-02 15:04:05"))
		}
	}

	// SubTimeString - Calculate difference between two time strings / 두 시간 문자열 사이의 차이 계산
	logger.Info("")
	logger.Info("--- SubTimeString: String to TimeDiff ---")
	logger.Info("--- SubTimeString: 문자열을 TimeDiff로 ---")

	diffStr, err := timeutil.SubTimeString("2024-10-04 08:34:42", "2024-10-14 14:56:23")
	if err == nil {
		logger.Info("SubTimeString",
			"from", "2024-10-04 08:34:42",
			"to", "2024-10-14 14:56:23",
			"result", diffStr.String(),
			"humanized", diffStr.Humanize())
	}

	// DiffInDaysString - Days between two date strings / 두 날짜 문자열 사이의 일수
	logger.Info("")
	logger.Info("--- Date Difference Functions (String versions) ---")
	logger.Info("--- 날짜 차이 함수 (문자열 버전) ---")

	daysStr, _ := timeutil.DiffInDaysString("2024-10-04", "Oct 14, 2024")
	logger.Info("DiffInDaysString", "from", "2024-10-04", "to", "Oct 14, 2024", "days", fmt.Sprintf("%.2f", daysStr))

	hoursStr, _ := timeutil.DiffInHoursString("2024-10-04 08:00", "2024-10-04 14:30")
	logger.Info("DiffInHoursString", "from", "08:00", "to", "14:30", "hours", fmt.Sprintf("%.2f", hoursStr))

	// AgeInYearsString - Calculate age from birth date string / 생년월일 문자열로부터 나이 계산
	logger.Info("")
	logger.Info("--- Age Calculation (String version) ---")
	logger.Info("--- 나이 계산 (문자열 버전) ---")

	ageYears, errAge := timeutil.AgeInYearsString("1990-01-15")
	if errAge == nil {
		logger.Info("AgeInYearsString", "birthDate", "1990-01-15", "age", ageYears)
	}

	ageDetail, errAge2 := timeutil.AgeString("Jan 15, 1990")
	if errAge2 == nil {
		logger.Info("AgeString", "birthDate", "Jan 15, 1990",
			"years", ageDetail.Years,
			"months", ageDetail.Months,
			"days", ageDetail.Days)
	}

	// RelativeTimeString - Human-readable relative time / 사람이 읽기 쉬운 상대 시간
	logger.Info("")
	logger.Info("--- Relative Time (String version) ---")
	logger.Info("--- 상대 시간 (문자열 버전) ---")

	relStr, _ := timeutil.RelativeTimeString("2024-10-13 15:30:00")
	logger.Info("RelativeTimeString", "time", "2024-10-13 15:30:00", "relative", relStr)

	// IsBusinessDayString - Check if date is business day / 영업일 확인
	logger.Info("")
	logger.Info("--- Business Day Check (String version) ---")
	logger.Info("--- 영업일 확인 (문자열 버전) ---")

	isBizDay, _ := timeutil.IsBusinessDayString("2024-10-14") // Monday
	logger.Info("IsBusinessDayString", "date", "2024-10-14 (Monday)", "isBusinessDay", isBizDay)

	isWeekend, _ := timeutil.IsWeekendString("2024-10-12") // Saturday
	logger.Info("IsWeekendString", "date", "2024-10-12 (Saturday)", "isWeekend", isWeekend)

	// AddDaysString - Add days to date string / 날짜 문자열에 일수 더하기
	logger.Info("")
	logger.Info("--- Date Arithmetic (String versions) ---")
	logger.Info("--- 날짜 연산 (문자열 버전) ---")

	futureDate, _ := timeutil.AddDaysString("2024-10-04", 7)
	logger.Info("AddDaysString", "date", "2024-10-04", "add", "7 days", "result", timeutil.FormatDate(futureDate))

	pastDate, _ := timeutil.SubDaysString("2024-10-14", 7)
	logger.Info("SubDaysString", "date", "2024-10-14", "subtract", "7 days", "result", timeutil.FormatDate(pastDate))

	// FormatString - Convert between different formats / 다른 포맷으로 변환
	logger.Info("")
	logger.Info("--- Format Conversion (String versions) ---")
	logger.Info("--- 포맷 변환 (문자열 버전) ---")

	formattedStr, _ := timeutil.FormatString("Oct 04, 2024", "2006-01-02")
	logger.Info("FormatString", "input", "Oct 04, 2024", "format", "2006-01-02", "result", formattedStr)

	dateOnlyStr, _ := timeutil.FormatDateString("2024-10-04 15:30:00")
	logger.Info("FormatDateString", "input", "2024-10-04 15:30:00", "result", dateOnlyStr)

	iso8601String, _ := timeutil.FormatISO8601String("Oct 04, 2024")
	logger.Info("FormatISO8601String", "input", "Oct 04, 2024", "result", iso8601String)

	// WeekdayKoreanString - Get Korean weekday name / 한글 요일 이름 가져오기
	logger.Info("")
	logger.Info("--- Weekday Names (String versions) ---")
	logger.Info("--- 요일 이름 (문자열 버전) ---")

	weekdayKor, _ := timeutil.WeekdayKoreanString("2024-10-14") // Monday
	logger.Info("WeekdayKoreanString", "date", "2024-10-14", "weekday", weekdayKor)

	weekdayShortKor, _ := timeutil.WeekdayShortKoreanString("Oct 14, 2024")
	logger.Info("WeekdayShortKoreanString", "date", "Oct 14, 2024", "weekday", weekdayShortKor)

	// StartOfDayString / EndOfDayString - Get start/end of day / 하루의 시작/끝 가져오기
	logger.Info("")
	logger.Info("--- Time Boundaries (String versions) ---")
	logger.Info("--- 시간 경계 (문자열 버전) ---")

	startOfDayStr, _ := timeutil.StartOfDayString("2024-10-04 15:30:45")
	logger.Info("StartOfDayString", "input", "2024-10-04 15:30:45", "result", timeutil.FormatDateTime(startOfDayStr))

	endOfDayStr, _ := timeutil.EndOfDayString("2024-10-04")
	logger.Info("EndOfDayString", "input", "2024-10-04", "result", timeutil.FormatDateTime(endOfDayStr))

	// IsSameDayString / IsBeforeString / IsAfterString - Date comparisons / 날짜 비교
	logger.Info("")
	logger.Info("--- Date Comparisons (String versions) ---")
	logger.Info("--- 날짜 비교 (문자열 버전) ---")

	sameDay, _ := timeutil.IsSameDayString("2024-10-04 08:00", "Oct 04, 2024")
	logger.Info("IsSameDayString", "date1", "2024-10-04 08:00", "date2", "Oct 04, 2024", "same", sameDay)

	isBefore, _ := timeutil.IsBeforeString("2024-10-04", "2024-10-14")
	logger.Info("IsBeforeString", "date1", "2024-10-04", "date2", "2024-10-14", "before", isBefore)

	isBetween, _ := timeutil.IsBetweenString("2024-10-10", "2024-10-04", "2024-10-14")
	logger.Info("IsBetweenString", "date", "2024-10-10", "start", "2024-10-04", "end", "2024-10-14", "between", isBetween)

	logger.Info("")
	logger.Info("Total: 102+ functions demonstrated (including new String versions!)")
	logger.Info("총 102개 이상 함수 시연 완료 (새로운 String 버전 포함!)")
	logger.Info("")
	logger.Info("Check the log file at ./logs/timeutil-example.log for detailed output")
	logger.Info("상세한 출력은 ./logs/timeutil-example.log 파일을 확인하세요")

	logger.Info("=" + repeat("=", 60))
	logger.Info("Examples completed! / 예제 완료!")
	logger.Info("=" + repeat("=", 60))
}

// Helper function to repeat a string / 문자열 반복 헬퍼 함수
func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
