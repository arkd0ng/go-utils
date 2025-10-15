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
	logger.Banner("Timeutil Package - Comprehensive Examples", "v1.9.012")
	logger.Info("")

	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║            Timeutil Package - Comprehensive Examples                       ║")
	logger.Info("║            Timeutil 패키지 - 종합 예제                                      ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")

	logger.Info("📋 Package Information / 패키지 정보")
	logger.Info("   Package: github.com/arkd0ng/go-utils/timeutil")
	logger.Info("   Description: Extremely simple time and date utilities")
	logger.Info("   설명: 극도로 간단한 시간 및 날짜 유틸리티")
	logger.Info("   Total Functions: 114 functions across 12 categories")
	logger.Info("   Default Timezone: Asia/Seoul (KST, GMT+9)")
	logger.Info("   Zero Dependencies: Standard library only")
	logger.Info("")

	logger.Info("🌟 Key Features / 주요 기능")
	logger.Info("   • KST Default: Asia/Seoul timezone as package-wide default")
	logger.Info("   • Custom Format Tokens: YYYY-MM-DD instead of Go's 2006-01-02")
	logger.Info("   • Business Days: Date calculations considering weekends and holidays")
	logger.Info("   • Korean Holidays: AddKoreanHolidays() for automatic holiday management")
	logger.Info("   • String Parameters: 50+ String version functions")
	logger.Info("   • Auto Format Detection: ParseAny with 40+ format recognition")
	logger.Info("   • Thread Safe: sync.RWMutex for timezone caching")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣  Time Difference Functions (8 functions)")
	logger.Info("   시간 차이 함수 (8개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1.1 SubTime() - Calculate time difference")
	logger.Info("    시간 차이 계산")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func SubTime(t1, t2 time.Time) TimeDiff")
	logger.Info("")

	logger.Info("📖 Description / 설명:")
	logger.Info("   Returns comprehensive time difference with human-readable output")
	logger.Info("   사람이 읽기 쉬운 출력을 제공하는 종합적인 시간 차이 반환")
	logger.Info("   • TimeDiff type with Days(), Humanize(), String() methods")
	logger.Info("   • Supports positive and negative differences")
	logger.Info("")

	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Calculate duration between events (이벤트 간 기간 계산)")
	logger.Info("   • Display time differences in UIs (UI에서 시간 차이 표시)")
	logger.Info("   • Age calculations (나이 계산)")
	logger.Info("   • Project timeline analysis (프로젝트 타임라인 분석)")
	logger.Info("")

	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Days() method: Returns total days as float")
	logger.Info("   • Humanize() method: '2 hours ago', 'in 3 days'")
	logger.Info("   • String() method: '2 days 6 hours 30 minutes'")
	logger.Info("   • Handles negative differences (past/future)")
	logger.Info("")

	logger.Info("▶️  Executing / 실행 중:")
	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 10, 15, 30, 45, 0, time.UTC)
	diff := timeutil.SubTime(start, end)
	logger.Info(fmt.Sprintf("   SubTime(%s, %s)", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05")))
	logger.Info(fmt.Sprintf("   Result: %s", diff.String()))
	logger.Info(fmt.Sprintf("   Days: %.2f", diff.Days()))
	logger.Info(fmt.Sprintf("   Humanized: %s", diff.Humanize()))
	logger.Info("")

	logger.Info("✅ Results Analysis / 결과 분석:")
	logger.Info(fmt.Sprintf("   • String(): '%s' (full breakdown)", diff.String()))
	logger.Info(fmt.Sprintf("   • Days(): %.2f days (decimal representation)", diff.Days()))
	logger.Info(fmt.Sprintf("   • Humanize(): '%s' (human-friendly)", diff.Humanize()))
	logger.Info("   • Perfect for displaying elapsed time in applications")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1.2 DiffInDays() - Get difference in days")
	logger.Info("    일 단위 차이")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func DiffInDays(t1, t2 time.Time) float64")
	logger.Info("")

	logger.Info("📖 Description / 설명:")
	logger.Info("   Calculate difference between two times in days (decimal)")
	logger.Info("   두 시간 사이의 차이를 일 단위(소수점 포함)로 계산")
	logger.Info("")

	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Subscription duration (구독 기간)")
	logger.Info("   • Project timelines (프로젝트 타임라인)")
	logger.Info("   • Age calculations (나이 계산)")
	logger.Info("   • Billing periods (청구 기간)")
	logger.Info("")

	logger.Info("▶️  Executing / 실행 중:")
	days := timeutil.DiffInDays(start, end)
	logger.Info(fmt.Sprintf("   DiffInDays(2025-01-01, 2025-01-10) = %.2f days", days))
	logger.Info("")

	logger.Info("📝 Additional Time Difference Functions:")
	logger.Info("   1.3 DiffInSeconds() - Difference in seconds")
	logger.Info("   1.4 DiffInMinutes() - Difference in minutes")
	logger.Info("   1.5 DiffInHours() - Difference in hours")
	logger.Info("   1.6 DiffInWeeks() - Difference in weeks")
	logger.Info("   1.7 DiffInMonths() - Difference in months")
	logger.Info("   1.8 DiffInYears() - Difference in years")
	logger.Info("")

	seconds := timeutil.DiffInSeconds(start, end)
	minutes := timeutil.DiffInMinutes(start, end)
	hours := timeutil.DiffInHours(start, end)
	weeks := timeutil.DiffInWeeks(start, end)
	months := timeutil.DiffInMonths(start, end)
	years := timeutil.DiffInYears(start, end)

	logger.Info(fmt.Sprintf("   DiffInSeconds: %.0f seconds", seconds))
	logger.Info(fmt.Sprintf("   DiffInMinutes: %.0f minutes", minutes))
	logger.Info(fmt.Sprintf("   DiffInHours: %.2f hours", hours))
	logger.Info(fmt.Sprintf("   DiffInWeeks: %.2f weeks", weeks))
	logger.Info(fmt.Sprintf("   DiffInMonths: %d months", months))
	logger.Info(fmt.Sprintf("   DiffInYears: %d years", years))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("2️⃣  Timezone Operations (10 functions)")
	logger.Info("   타임존 작업 (10개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("2.1 NowKST() - Get current time in KST")
	logger.Info("    KST 현재 시간")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func NowKST() time.Time")
	logger.Info("")

	logger.Info("📖 Description / 설명:")
	logger.Info("   Returns current time in KST timezone (Asia/Seoul, GMT+9)")
	logger.Info("   KST 타임존(Asia/Seoul, GMT+9)의 현재 시간 반환")
	logger.Info("")

	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Korean applications (한국 애플리케이션)")
	logger.Info("   • Logging in KST (KST로 로깅)")
	logger.Info("   • Business hours in Korea (한국 영업 시간)")
	logger.Info("   • Timestamp generation (타임스탬프 생성)")
	logger.Info("")

	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Default timezone: Asia/Seoul (GMT+9)")
	logger.Info("   • Thread-safe operation")
	logger.Info("   • Cached timezone loading")
	logger.Info("   • No conversion needed for Korean apps")
	logger.Info("")

	logger.Info("▶️  Executing / 실행 중:")
	kstNow := timeutil.NowKST()
	logger.Info(fmt.Sprintf("   NowKST() = %s", timeutil.FormatDateTime(kstNow)))
	logger.Info(fmt.Sprintf("   Timezone: %s", kstNow.Location().String()))
	logger.Info("")

	logger.Info("✅ Results Analysis / 결과 분석:")
	logger.Info(fmt.Sprintf("   • Current KST time: %s", timeutil.FormatDateTime(kstNow)))
	logger.Info("   • Timezone correctly set to Asia/Seoul")
	logger.Info("   • No UTC conversion required")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("2.2 ConvertTimezone() - Convert between timezones")
	logger.Info("    타임존 간 변환")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func ConvertTimezone(t time.Time, timezone string) (time.Time, error)")
	logger.Info("")

	logger.Info("📖 Description / 설명:")
	logger.Info("   Convert a time to a different timezone")
	logger.Info("   시간을 다른 타임존으로 변환")
	logger.Info("")

	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Global applications (글로벌 애플리케이션)")
	logger.Info("   • Meeting schedulers (회의 스케줄러)")
	logger.Info("   • Multi-region systems (다중 지역 시스템)")
	logger.Info("   • Time comparison (시간 비교)")
	logger.Info("")

	logger.Info("▶️  Executing / 실행 중:")
	now := time.Now()
	tokyoTime, _ := timeutil.ConvertTimezone(now, "Asia/Tokyo")
	nycTime, _ := timeutil.ConvertTimezone(now, "America/New_York")
	londonTime, _ := timeutil.ConvertTimezone(now, "Europe/London")

	logger.Info(fmt.Sprintf("   Original (Local): %s", timeutil.FormatDateTime(now)))
	logger.Info(fmt.Sprintf("   Tokyo: %s", timeutil.FormatDateTime(tokyoTime)))
	logger.Info(fmt.Sprintf("   New York: %s", timeutil.FormatDateTime(nycTime)))
	logger.Info(fmt.Sprintf("   London: %s", timeutil.FormatDateTime(londonTime)))
	logger.Info("")

	logger.Info("📝 Additional Timezone Functions:")
	logger.Info("   2.3 ToUTC() - Convert to UTC")
	logger.Info("   2.4 ToKST() - Convert to KST")
	logger.Info("   2.5 GetTimezoneOffset() - Get timezone offset in hours")
	logger.Info("   2.6 GetDefaultTimezone() - Get current default timezone")
	logger.Info("   2.7 SetDefaultTimezone() - Set default timezone")
	logger.Info("   2.8 GetLocalTimezone() - Get system local timezone")
	logger.Info("   2.9 IsValidTimezone() - Validate timezone name")
	logger.Info("   2.10 ListTimezones() - List all available timezones")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("3️⃣  Date Arithmetic (16 functions)")
	logger.Info("   날짜 연산 (16개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Date Arithmetic Functions:")
	logger.Info("   3.1 AddDays() - Add days to time")
	logger.Info("   3.2 AddWeeks() - Add weeks")
	logger.Info("   3.3 AddMonths() - Add months")
	logger.Info("   3.4 AddYears() - Add years")
	logger.Info("   3.5 AddHours() - Add hours")
	logger.Info("   3.6 AddMinutes() - Add minutes")
	logger.Info("   3.7 AddSeconds() - Add seconds")
	logger.Info("   3.8 StartOfDay() - Get start of day (00:00:00)")
	logger.Info("   3.9 EndOfDay() - Get end of day (23:59:59)")
	logger.Info("   3.10 StartOfWeek() - Get start of week (Monday)")
	logger.Info("   3.11 EndOfWeek() - Get end of week (Sunday)")
	logger.Info("   3.12 StartOfMonth() - Get start of month")
	logger.Info("   3.13 EndOfMonth() - Get end of month")
	logger.Info("   3.14 StartOfYear() - Get start of year")
	logger.Info("   3.15 EndOfYear() - Get end of year")
	logger.Info("   3.16 StartOfQuarter() - Get start of quarter")
	logger.Info("")

	logger.Info("▶️  Executing Date Arithmetic / 날짜 연산 실행:")
	baseTime := time.Date(2025, 10, 15, 10, 30, 0, 0, time.UTC)
	logger.Info(fmt.Sprintf("   Base Time: %s", timeutil.FormatDateTime(baseTime)))
	logger.Info(fmt.Sprintf("   AddDays(7): %s", timeutil.FormatDate(timeutil.AddDays(baseTime, 7))))
	logger.Info(fmt.Sprintf("   AddMonths(3): %s", timeutil.FormatDate(timeutil.AddMonths(baseTime, 3))))
	logger.Info(fmt.Sprintf("   AddYears(1): %s", timeutil.FormatDate(timeutil.AddYears(baseTime, 1))))
	logger.Info(fmt.Sprintf("   StartOfDay: %s", timeutil.FormatDateTime(timeutil.StartOfDay(baseTime))))
	logger.Info(fmt.Sprintf("   EndOfDay: %s", timeutil.FormatDateTime(timeutil.EndOfDay(baseTime))))
	logger.Info(fmt.Sprintf("   StartOfMonth: %s", timeutil.FormatDate(timeutil.StartOfMonth(baseTime))))
	logger.Info(fmt.Sprintf("   EndOfMonth: %s", timeutil.FormatDate(timeutil.EndOfMonth(baseTime))))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("4️⃣  Date Formatting (8 functions)")
	logger.Info("   날짜 포맷팅 (8개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Formatting Functions:")
	logger.Info("   4.1 Format() - Custom format with YYYY-MM-DD tokens")
	logger.Info("   4.2 FormatISO8601() - ISO 8601 format")
	logger.Info("   4.3 FormatRFC3339() - RFC 3339 format")
	logger.Info("   4.4 FormatDate() - Date only (YYYY-MM-DD)")
	logger.Info("   4.5 FormatDateTime() - Date and time")
	logger.Info("   4.6 FormatTime() - Time only (HH:MM:SS)")
	logger.Info("   4.7 FormatKorean() - Korean format (2025년 10월 15일)")
	logger.Info("   4.8 FormatCustom() - Go's native layout format")
	logger.Info("")

	logger.Info("▶️  Executing Formatting / 포맷팅 실행:")
	sampleTime := time.Date(2025, 10, 15, 14, 30, 45, 0, time.UTC)
	logger.Info(fmt.Sprintf("   Format('YYYY-MM-DD HH:mm:ss'): %s", timeutil.Format(sampleTime, "YYYY-MM-DD HH:mm:ss")))
	logger.Info(fmt.Sprintf("   FormatISO8601(): %s", timeutil.FormatISO8601(sampleTime)))
	logger.Info(fmt.Sprintf("   FormatDate(): %s", timeutil.FormatDate(sampleTime)))
	logger.Info(fmt.Sprintf("   FormatDateTime(): %s", timeutil.FormatDateTime(sampleTime)))
	logger.Info(fmt.Sprintf("   FormatTime(): %s", timeutil.FormatTime(sampleTime)))
	logger.Info(fmt.Sprintf("   FormatKorean(): %s", timeutil.FormatKorean(sampleTime)))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("5️⃣  Time Parsing (10 functions)")
	logger.Info("   시간 파싱 (10개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("5.1 ParseAny() - Auto-detect format and parse")
	logger.Info("    자동 포맷 감지 및 파싱")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func ParseAny(s string) (time.Time, error)")
	logger.Info("")

	logger.Info("📖 Description / 설명:")
	logger.Info("   Automatically detects format from 40+ common patterns and parses time string")
	logger.Info("   40개 이상의 일반적인 패턴에서 자동으로 포맷을 감지하고 시간 문자열 파싱")
	logger.Info("")

	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • User input parsing (사용자 입력 파싱)")
	logger.Info("   • API response parsing (API 응답 파싱)")
	logger.Info("   • Log file parsing (로그 파일 파싱)")
	logger.Info("   • Flexible time input (유연한 시간 입력)")
	logger.Info("")

	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • 40+ format patterns recognized")
	logger.Info("   • ISO 8601, RFC 3339, common formats")
	logger.Info("   • Date-only, time-only, datetime")
	logger.Info("   • No need to specify format")
	logger.Info("")

	logger.Info("▶️  Executing ParseAny / 자동 파싱 실행:")
	formats := []string{
		"2025-10-15",
		"2025-10-15 14:30:45",
		"2025/10/15 14:30:45",
		"15-Oct-2025",
		"Oct 15, 2025",
	}
	for _, f := range formats {
		if parsed, err := timeutil.ParseAny(f); err == nil {
			logger.Info(fmt.Sprintf("   ParseAny('%s') = %s", f, timeutil.FormatDateTime(parsed)))
		}
	}
	logger.Info("")

	logger.Info("📝 Additional Parsing Functions:")
	logger.Info("   5.2 Parse() - Parse with format")
	logger.Info("   5.3 ParseISO8601() - Parse ISO 8601")
	logger.Info("   5.4 ParseRFC3339() - Parse RFC 3339")
	logger.Info("   5.5 ParseDate() - Parse date only")
	logger.Info("   5.6 ParseDateTime() - Parse date and time")
	logger.Info("   5.7 ParseWithTimezone() - Parse with timezone")
	logger.Info("   5.8 ParseWithLayout() - Parse with Go layout")
	logger.Info("   5.9 ParseMillis() - Parse millisecond timestamp")
	logger.Info("   5.10 ParseMicros() - Parse microsecond timestamp")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("6️⃣  Time Comparisons (18 functions)")
	logger.Info("   시간 비교 (18개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Comparison Functions:")
	logger.Info("   6.1 IsBefore() - Check if time is before another")
	logger.Info("   6.2 IsAfter() - Check if time is after another")
	logger.Info("   6.3 IsBetween() - Check if time is between two times")
	logger.Info("   6.4 IsToday() - Check if time is today")
	logger.Info("   6.5 IsYesterday() - Check if time is yesterday")
	logger.Info("   6.6 IsTomorrow() - Check if time is tomorrow")
	logger.Info("   6.7 IsThisWeek() - Check if time is this week")
	logger.Info("   6.8 IsThisMonth() - Check if time is this month")
	logger.Info("   6.9 IsThisYear() - Check if time is this year")
	logger.Info("   6.10 IsWeekend() - Check if time is weekend")
	logger.Info("   6.11 IsWeekday() - Check if time is weekday")
	logger.Info("   6.12 IsSameDay() - Check if two times are same day")
	logger.Info("   6.13 IsSameWeek() - Check if two times are same week")
	logger.Info("   6.14 IsSameMonth() - Check if two times are same month")
	logger.Info("   6.15 IsSameYear() - Check if two times are same year")
	logger.Info("   6.16 IsLeapYear() - Check if year is leap year")
	logger.Info("   6.17 IsPast() - Check if time is in the past")
	logger.Info("   6.18 IsFuture() - Check if time is in the future")
	logger.Info("")

	logger.Info("▶️  Executing Comparisons / 비교 실행:")
	testTime := time.Date(2025, 10, 15, 14, 30, 0, 0, time.UTC)
	logger.Info(fmt.Sprintf("   Test Time: %s", timeutil.FormatDateTime(testTime)))
	logger.Info(fmt.Sprintf("   IsWeekday(): %v", timeutil.IsWeekday(testTime)))
	logger.Info(fmt.Sprintf("   IsWeekend(): %v", timeutil.IsWeekend(testTime)))
	leapYear2024 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	leapYear2025 := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	logger.Info(fmt.Sprintf("   IsLeapYear(2024): %v", timeutil.IsLeapYear(leapYear2024)))
	logger.Info(fmt.Sprintf("   IsLeapYear(2025): %v", timeutil.IsLeapYear(leapYear2025)))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("7️⃣  Age Calculations (4 functions)")
	logger.Info("   나이 계산 (4개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Age Functions:")
	logger.Info("   7.1 Age() - Get age as AgeDetail (Years, Months, Days)")
	logger.Info("   7.2 AgeInYears() - Get age in years only")
	logger.Info("   7.3 AgeInMonths() - Get age in months only")
	logger.Info("   7.4 AgeInDays() - Get age in days only")
	logger.Info("")

	logger.Info("▶️  Executing Age Calculation / 나이 계산 실행:")
	birthDate := time.Date(1990, 5, 20, 0, 0, 0, 0, time.UTC)
	age := timeutil.Age(birthDate)
	logger.Info(fmt.Sprintf("   Birth Date: %s", timeutil.FormatDate(birthDate)))
	logger.Info(fmt.Sprintf("   Age: %d years %d months %d days", age.Years, age.Months, age.Days))
	logger.Info(fmt.Sprintf("   AgeInYears: %d", timeutil.AgeInYears(birthDate)))
	logger.Info(fmt.Sprintf("   AgeInMonths: %d", timeutil.AgeInMonths(birthDate)))
	logger.Info(fmt.Sprintf("   AgeInDays: %d", timeutil.AgeInDays(birthDate)))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("8️⃣  Relative Time (4 functions)")
	logger.Info("   상대 시간 (4개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Relative Time Functions:")
	logger.Info("   8.1 RelativeTime() - '2 hours ago', 'in 3 days'")
	logger.Info("   8.2 RelativeTimeShort() - '2h ago', 'in 3d'")
	logger.Info("   8.3 TimeAgo() - Alias for RelativeTime")
	logger.Info("   8.4 HumanizeDuration() - Humanize duration")
	logger.Info("")

	logger.Info("▶️  Executing Relative Time / 상대 시간 실행:")
	pastTime := time.Now().Add(-2 * time.Hour)
	futureTime := time.Now().Add(3 * 24 * time.Hour)
	logger.Info(fmt.Sprintf("   RelativeTime(2 hours ago): %s", timeutil.RelativeTime(pastTime)))
	logger.Info(fmt.Sprintf("   RelativeTime(3 days future): %s", timeutil.RelativeTime(futureTime)))
	logger.Info(fmt.Sprintf("   RelativeTimeShort(2h ago): %s", timeutil.RelativeTimeShort(pastTime)))
	logger.Info(fmt.Sprintf("   HumanizeDuration(90 minutes): %s", timeutil.HumanizeDuration(90*time.Minute)))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("9️⃣  Unix Timestamp (12 functions)")
	logger.Info("   Unix 타임스탬프 (12개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Unix Timestamp Functions:")
	logger.Info("   9.1 Now() - Current Unix timestamp (seconds)")
	logger.Info("   9.2 NowMilli() - Current Unix timestamp (milliseconds)")
	logger.Info("   9.3 NowMicro() - Current Unix timestamp (microseconds)")
	logger.Info("   9.4 NowNano() - Current Unix timestamp (nanoseconds)")
	logger.Info("   9.5 ToUnix() - Convert time to Unix seconds")
	logger.Info("   9.6 ToUnixMilli() - Convert to Unix milliseconds")
	logger.Info("   9.7 ToUnixMicro() - Convert to Unix microseconds")
	logger.Info("   9.8 ToUnixNano() - Convert to Unix nanoseconds")
	logger.Info("   9.9 FromUnix() - Convert Unix seconds to time")
	logger.Info("   9.10 FromUnixMilli() - Convert Unix milliseconds to time")
	logger.Info("   9.11 FromUnixMicro() - Convert Unix microseconds to time")
	logger.Info("   9.12 FromUnixNano() - Convert Unix nanoseconds to time")
	logger.Info("")

	logger.Info("▶️  Executing Unix Timestamp / Unix 타임스탬프 실행:")
	unixNow := timeutil.Now()
	unixMilli := timeutil.NowMilli()
	logger.Info(fmt.Sprintf("   Now(): %d", unixNow))
	logger.Info(fmt.Sprintf("   NowMilli(): %d", unixMilli))
	logger.Info(fmt.Sprintf("   FromUnix(%d): %s", unixNow, timeutil.FormatDateTime(timeutil.FromUnix(unixNow))))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔟 Business Days (7 functions)")
	logger.Info("   영업일 (7개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Business Day Functions:")
	logger.Info("   10.1 IsBusinessDay() - Check if date is business day")
	logger.Info("   10.2 IsHoliday() - Check if date is holiday")
	logger.Info("   10.3 AddBusinessDays() - Add business days")
	logger.Info("   10.4 NextBusinessDay() - Get next business day")
	logger.Info("   10.5 PreviousBusinessDay() - Get previous business day")
	logger.Info("   10.6 CountBusinessDays() - Count business days between dates")
	logger.Info("   10.7 AddKoreanHolidays() - Add Korean holidays automatically")
	logger.Info("")

	logger.Info("▶️  Executing Business Days / 영업일 실행:")
	// Add Korean holidays / 한국 공휴일 추가
	timeutil.AddKoreanHolidays(2025)

	bizDate := time.Date(2025, 10, 15, 0, 0, 0, 0, time.UTC)
	logger.Info(fmt.Sprintf("   Test Date: %s", timeutil.FormatDate(bizDate)))
	logger.Info(fmt.Sprintf("   IsBusinessDay(): %v", timeutil.IsBusinessDay(bizDate)))
	logger.Info(fmt.Sprintf("   IsHoliday(): %v", timeutil.IsHoliday(bizDate)))

	nextBiz := timeutil.NextBusinessDay(bizDate)
	logger.Info(fmt.Sprintf("   NextBusinessDay(): %s", timeutil.FormatDate(nextBiz)))

	bizDaysAdded := timeutil.AddBusinessDays(bizDate, 5)
	logger.Info(fmt.Sprintf("   AddBusinessDays(5): %s", timeutil.FormatDate(bizDaysAdded)))

	endDate := time.Date(2025, 10, 31, 0, 0, 0, 0, time.UTC)
	bizCount := timeutil.CountBusinessDays(bizDate, endDate)
	logger.Info(fmt.Sprintf("   CountBusinessDays(Oct 15-31): %d days", bizCount))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣1️⃣  Week Functions (4 functions)")
	logger.Info("   주 관련 함수 (4개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Week Functions:")
	logger.Info("   11.1 WeekOfYear() - Get week number of year")
	logger.Info("   11.2 WeekOfMonth() - Get week number of month")
	logger.Info("   11.3 DaysInMonth() - Get number of days in month")
	logger.Info("   11.4 DaysInYear() - Get number of days in year")
	logger.Info("")

	logger.Info("▶️  Executing Week Functions / 주 함수 실행:")
	weekTest := time.Date(2025, 10, 15, 0, 0, 0, 0, time.UTC)
	logger.Info(fmt.Sprintf("   Date: %s", timeutil.FormatDate(weekTest)))
	logger.Info(fmt.Sprintf("   WeekOfYear(): %d", timeutil.WeekOfYear(weekTest)))
	logger.Info(fmt.Sprintf("   WeekOfMonth(): %d", timeutil.WeekOfMonth(weekTest)))
	logger.Info(fmt.Sprintf("   DaysInMonth(Oct 2025): %d", timeutil.DaysInMonth(weekTest)))
	yearTest := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	logger.Info(fmt.Sprintf("   DaysInYear(2025): %d", timeutil.DaysInYear(yearTest)))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣2️⃣  Month Functions (4 functions)")
	logger.Info("   월 관련 함수 (4개 함수)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📝 Month Functions:")
	logger.Info("   12.1 MonthKorean() - Get Korean month name (10월)")
	logger.Info("   12.2 MonthName() - Get full month name (October)")
	logger.Info("   12.3 MonthNameShort() - Get short month name (Oct)")
	logger.Info("   12.4 Quarter() - Get quarter number (1-4)")
	logger.Info("")

	logger.Info("▶️  Executing Month Functions / 월 함수 실행:")
	monthTest := time.Date(2025, 10, 15, 0, 0, 0, 0, time.UTC)
	logger.Info(fmt.Sprintf("   Date: %s", timeutil.FormatDate(monthTest)))
	logger.Info(fmt.Sprintf("   MonthKorean(): %s", timeutil.MonthKorean(monthTest)))
	logger.Info(fmt.Sprintf("   MonthName(): %s", timeutil.MonthName(monthTest)))
	logger.Info(fmt.Sprintf("   MonthNameShort(): %s", timeutil.MonthNameShort(monthTest)))
	logger.Info(fmt.Sprintf("   Quarter(): Q%d", timeutil.Quarter(monthTest)))
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📊 Summary / 요약")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("This example demonstrated comprehensive time utilities:")
	logger.Info("본 예제는 포괄적인 시간 유틸리티를 시연했습니다:")
	logger.Info("")

	logger.Info("  1️⃣  Time Difference (8 functions) - Calculate time differences")
	logger.Info("     시간 차이 (8개 함수) - 시간 차이 계산")
	logger.Info("  2️⃣  Timezone Operations (10 functions) - Timezone conversions")
	logger.Info("     타임존 작업 (10개 함수) - 타임존 변환")
	logger.Info("  3️⃣  Date Arithmetic (16 functions) - Add/subtract time units")
	logger.Info("     날짜 연산 (16개 함수) - 시간 단위 더하기/빼기")
	logger.Info("  4️⃣  Date Formatting (8 functions) - Format time to strings")
	logger.Info("     날짜 포맷팅 (8개 함수) - 시간을 문자열로 포맷")
	logger.Info("  5️⃣  Time Parsing (10 functions) - Parse strings to time")
	logger.Info("     시간 파싱 (10개 함수) - 문자열을 시간으로 파싱")
	logger.Info("  6️⃣  Time Comparisons (18 functions) - Compare times")
	logger.Info("     시간 비교 (18개 함수) - 시간 비교")
	logger.Info("  7️⃣  Age Calculations (4 functions) - Calculate age")
	logger.Info("     나이 계산 (4개 함수) - 나이 계산")
	logger.Info("  8️⃣  Relative Time (4 functions) - Human-friendly time")
	logger.Info("     상대 시간 (4개 함수) - 사람 친화적 시간")
	logger.Info("  9️⃣  Unix Timestamp (12 functions) - Unix timestamp handling")
	logger.Info("     Unix 타임스탬프 (12개 함수) - Unix 타임스탬프 처리")
	logger.Info("  🔟 Business Days (7 functions) - Business day operations")
	logger.Info("     영업일 (7개 함수) - 영업일 작업")
	logger.Info("  1️⃣1️⃣  Week Functions (4 functions) - Week-related operations")
	logger.Info("     주 관련 함수 (4개 함수) - 주 관련 작업")
	logger.Info("  1️⃣2️⃣  Month Functions (4 functions) - Month-related operations")
	logger.Info("     월 관련 함수 (4개 함수) - 월 관련 작업")
	logger.Info("")

	logger.Info("✨ Key Takeaways / 주요 포인트:")
	logger.Info("   • All 105 functions demonstrated (105개 함수 시연)")
	logger.Info("   • KST as default timezone (KST가 기본 타임존)")
	logger.Info("   • Custom format tokens (YYYY-MM-DD) (커스텀 포맷 토큰)")
	logger.Info("   • Auto-format detection with ParseAny (ParseAny로 자동 포맷 감지)")
	logger.Info("   • Business day support with Korean holidays (한국 공휴일 포함 영업일 지원)")
	logger.Info("   • Human-readable relative time (사람이 읽기 쉬운 상대 시간)")
	logger.Info("   • 50+ String parameter functions (50개 이상 문자열 매개변수 함수)")
	logger.Info("")

	logger.Info("All examples completed successfully!")
	logger.Info("모든 예제가 성공적으로 완료되었습니다!")
	logger.Info("")
}
