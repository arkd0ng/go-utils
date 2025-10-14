package main

import (
	"fmt"
	"time"

	"github.com/arkd0ng/go-utils/timeutil"
)

func main() {
	fmt.Println("=== Timeutil Package Examples / Timeutil 패키지 예제 ===\n")

	// 1. Time Difference / 시간 차이
	fmt.Println("1. Time Difference / 시간 차이")
	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)
	diff := timeutil.SubTime(start, end)
	fmt.Printf("   Difference: %s\n", diff.String())
	fmt.Printf("   Days: %.2f\n", diff.Days())
	fmt.Printf("   Humanized: %s\n\n", diff.Humanize())

	// 2. Timezone Operations / 타임존 작업
	fmt.Println("2. Timezone Operations / 타임존 작업")
	fmt.Printf("   Default timezone: %s\n", timeutil.GetDefaultTimezone())
	now := time.Now()
	seoulTime, _ := timeutil.ConvertTimezone(now, "Asia/Seoul")
	nyTime, _ := timeutil.ConvertTimezone(now, "America/New_York")
	fmt.Printf("   Seoul time: %s\n", timeutil.FormatDateTime(seoulTime))
	fmt.Printf("   New York time: %s\n", timeutil.FormatDateTime(nyTime))
	kstNow := timeutil.NowKST()
	fmt.Printf("   KST now: %s\n\n", timeutil.FormatDateTime(kstNow))

	// 3. Date Arithmetic / 날짜 연산
	fmt.Println("3. Date Arithmetic / 날짜 연산")
	tomorrow := timeutil.AddDays(time.Now(), 1)
	nextWeek := timeutil.AddWeeks(time.Now(), 1)
	nextMonth := timeutil.AddMonths(time.Now(), 1)
	fmt.Printf("   Tomorrow: %s\n", timeutil.FormatDate(tomorrow))
	fmt.Printf("   Next week: %s\n", timeutil.FormatDate(nextWeek))
	fmt.Printf("   Next month: %s\n", timeutil.FormatDate(nextMonth))
	fmt.Printf("   Start of month: %s\n", timeutil.FormatDateTime(timeutil.StartOfMonth(time.Now())))
	fmt.Printf("   End of month: %s\n\n", timeutil.FormatDateTime(timeutil.EndOfMonth(time.Now())))

	// 4. Date Formatting / 날짜 포맷팅
	fmt.Println("4. Date Formatting / 날짜 포맷팅")
	testTime := time.Now()
	fmt.Printf("   ISO8601: %s\n", timeutil.FormatISO8601(testTime))
	fmt.Printf("   Date: %s\n", timeutil.FormatDate(testTime))
	fmt.Printf("   DateTime: %s\n", timeutil.FormatDateTime(testTime))
	fmt.Printf("   Custom: %s\n", timeutil.Format(testTime, "YYYY-MM-DD HH:mm:ss"))
	fmt.Printf("   Korean: %s\n\n", timeutil.FormatKorean(testTime))

	// 5. Time Parsing / 시간 파싱
	fmt.Println("5. Time Parsing / 시간 파싱")
	parsed1, _ := timeutil.ParseDate("2025-10-14")
	parsed2, _ := timeutil.ParseDateTime("2025-10-14 15:04:05")
	parsed3, _ := timeutil.Parse("2025-10-14")
	fmt.Printf("   Parsed date: %s\n", timeutil.FormatDate(parsed1))
	fmt.Printf("   Parsed datetime: %s\n", timeutil.FormatDateTime(parsed2))
	fmt.Printf("   Auto-parsed: %s\n\n", timeutil.FormatDate(parsed3))

	// 6. Time Comparisons / 시간 비교
	fmt.Println("6. Time Comparisons / 시간 비교")
	fmt.Printf("   Is today: %v\n", timeutil.IsToday(time.Now()))
	fmt.Printf("   Is weekend: %v\n", timeutil.IsWeekend(time.Now()))
	fmt.Printf("   Is weekday: %v\n", timeutil.IsWeekday(time.Now()))
	fmt.Printf("   Is this month: %v\n", timeutil.IsThisMonth(time.Now()))
	fmt.Printf("   Is this year: %v\n\n", timeutil.IsThisYear(time.Now()))

	// 7. Age Calculations / 나이 계산
	fmt.Println("7. Age Calculations / 나이 계산")
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	years := timeutil.AgeInYears(birthDate)
	months := timeutil.AgeInMonths(birthDate)
	days := timeutil.AgeInDays(birthDate)
	age := timeutil.Age(birthDate)
	fmt.Printf("   Age in years: %d\n", years)
	fmt.Printf("   Age in months: %d\n", months)
	fmt.Printf("   Age in days: %d\n", days)
	fmt.Printf("   Detailed age: %s\n\n", age.String())

	// 8. Relative Time / 상대 시간
	fmt.Println("8. Relative Time / 상대 시간")
	past := time.Now().Add(-2 * time.Hour)
	future := time.Now().Add(3 * time.Hour)
	fmt.Printf("   2 hours ago: %s\n", timeutil.RelativeTime(past))
	fmt.Printf("   3 hours from now: %s\n", timeutil.RelativeTime(future))
	fmt.Printf("   Short format: %s\n\n", timeutil.RelativeTimeShort(past))

	// 9. Unix Timestamp / Unix 타임스탬프
	fmt.Println("9. Unix Timestamp / Unix 타임스탬프")
	unix := timeutil.Now()
	unixMilli := timeutil.NowMilli()
	fmt.Printf("   Unix timestamp (seconds): %d\n", unix)
	fmt.Printf("   Unix timestamp (milliseconds): %d\n", unixMilli)
	fromUnix := timeutil.FromUnix(unix)
	fmt.Printf("   From unix: %s\n\n", timeutil.FormatDateTime(fromUnix))

	// 10. Business Days / 영업일
	fmt.Println("10. Business Days / 영업일")
	fmt.Printf("   Is business day: %v\n", timeutil.IsBusinessDay(time.Now()))
	nextBusinessDay := timeutil.AddBusinessDays(time.Now(), 5)
	fmt.Printf("   5 business days later: %s\n", timeutil.FormatDate(nextBusinessDay))

	// Add Korean holidays for 2025 / 2025년 한국 공휴일 추가
	timeutil.AddKoreanHolidays(2025)
	newYear := time.Date(2025, 1, 1, 0, 0, 0, 0, timeutil.KST)
	fmt.Printf("   Is Jan 1 holiday: %v\n", timeutil.IsHoliday(newYear))
	fmt.Printf("   Is Jan 1 business day: %v\n\n", timeutil.IsBusinessDay(newYear))

	fmt.Println("=== All examples completed! / 모든 예제 완료! ===")
}
