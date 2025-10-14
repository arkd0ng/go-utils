// Package timeutil provides extreme simplicity time and date utility functions.
// timeutil 패키지는 극도로 간단한 시간 및 날짜 유틸리티 함수를 제공합니다.
//
// This package reduces 10-20 lines of repetitive time manipulation code
// to a single function call. All functions are thread-safe and have no
// external dependencies (standard library only).
//
// 이 패키지는 10-20줄의 반복적인 시간 조작 코드를 단일 함수 호출로 줄입니다.
// 모든 함수는 스레드 안전하며 외부 의존성이 없습니다 (표준 라이브러리만).
//
// Design Philosophy: "20 lines → 1 line"
// 설계 철학: "20줄 → 1줄"
//
// Categories:
// - Time Difference: SubTime, DiffInSeconds, DiffInMinutes, DiffInHours
// - Timezone Operations: ConvertTimezone, GetTimezoneOffset, ListTimezones
// - Date Arithmetic: AddDays, AddWeeks, AddMonths, StartOfDay, EndOfDay
// - Date Formatting: FormatISO8601, FormatRFC3339, Format
// - Time Parsing: ParseISO8601, ParseDate, Parse
// - Business Days: AddBusinessDays, IsBusinessDay, CountBusinessDays
// - Time Comparisons: IsToday, IsYesterday, IsBefore, IsAfter
// - Age Calculations: AgeInYears, Age
// - Relative Time: RelativeTime, TimeAgo
// - Unix Timestamp: Now, NowMilli, FromUnix, ToUnix
//
// Example:
//
//	import "github.com/arkd0ng/go-utils/timeutil"
//
//	// Time difference / 시간 차이
//	diff := timeutil.SubTime(start, end)
//	fmt.Println(diff.Days())    // 2
//	fmt.Println(diff.String())  // "2 days 6 hours 30 minutes"
//
//	// Timezone conversion / 타임존 변환
//	seoulTime, _ := timeutil.ConvertTimezone(time.Now(), "Asia/Seoul")
//
//	// Date arithmetic / 날짜 연산
//	tomorrow := timeutil.AddDays(time.Now(), 1)
//	startOfMonth := timeutil.StartOfMonth(time.Now())
//
//	// Date formatting / 날짜 포맷팅
//	iso := timeutil.FormatISO8601(time.Now())
//	custom := timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
//
//	// Time comparisons / 시간 비교
//	if timeutil.IsToday(someDate) {
//	    fmt.Println("It's today!")
//	}
package timeutil

import (
	"fmt"
	"time"
)

// TimeDiff represents the difference between two times.
// TimeDiff는 두 시간 사이의 차이를 나타냅니다.
type TimeDiff struct {
	Duration time.Duration
}

// Seconds returns the time difference in seconds.
// Seconds는 시간 차이를 초 단위로 반환합니다.
func (td *TimeDiff) Seconds() float64 {
	return td.Duration.Seconds()
}

// Minutes returns the time difference in minutes.
// Minutes는 시간 차이를 분 단위로 반환합니다.
func (td *TimeDiff) Minutes() float64 {
	return td.Duration.Minutes()
}

// Hours returns the time difference in hours.
// Hours는 시간 차이를 시간 단위로 반환합니다.
func (td *TimeDiff) Hours() float64 {
	return td.Duration.Hours()
}

// Days returns the time difference in days.
// Days는 시간 차이를 일 단위로 반환합니다.
func (td *TimeDiff) Days() float64 {
	return td.Hours() / 24
}

// Weeks returns the time difference in weeks.
// Weeks는 시간 차이를 주 단위로 반환합니다.
func (td *TimeDiff) Weeks() float64 {
	return td.Days() / 7
}

// String returns a human-readable string representation of the time difference.
// String은 시간 차이를 사람이 읽기 쉬운 문자열로 반환합니다.
//
// Examples:
//   - "2 days 3 hours 15 minutes"
//   - "5 hours 30 minutes"
//   - "45 minutes 20 seconds"
func (td *TimeDiff) String() string {
	d := td.Duration
	if d < 0 {
		d = -d
	}

	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	var result string
	if days > 0 {
		result = fmt.Sprintf("%d days", days)
		if hours > 0 {
			result += fmt.Sprintf(" %d hours", hours)
		}
		if minutes > 0 {
			result += fmt.Sprintf(" %d minutes", minutes)
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

// Humanize returns a short human-readable string representation.
// Humanize는 짧은 사람이 읽기 쉬운 문자열을 반환합니다.
//
// Examples:
//   - "2d 3h 15m"
//   - "5h 30m"
//   - "45m 20s"
func (td *TimeDiff) Humanize() string {
	d := td.Duration
	if d < 0 {
		d = -d
	}

	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	var result string
	if days > 0 {
		result = fmt.Sprintf("%dd", days)
		if hours > 0 {
			result += fmt.Sprintf(" %dh", hours)
		}
		if minutes > 0 {
			result += fmt.Sprintf(" %dm", minutes)
		}
	} else if hours > 0 {
		result = fmt.Sprintf("%dh", hours)
		if minutes > 0 {
			result += fmt.Sprintf(" %dm", minutes)
		}
	} else if minutes > 0 {
		result = fmt.Sprintf("%dm", minutes)
		if seconds > 0 {
			result += fmt.Sprintf(" %ds", seconds)
		}
	} else {
		result = fmt.Sprintf("%ds", seconds)
	}

	return result
}

// Abs returns the absolute value of the time difference.
// Abs는 시간 차이의 절대값을 반환합니다.
func (td *TimeDiff) Abs() *TimeDiff {
	if td.Duration < 0 {
		return &TimeDiff{Duration: -td.Duration}
	}
	return td
}

// AgeDetail represents a person's age in years, months, and days.
// AgeDetail은 년, 월, 일 단위의 나이를 나타냅니다.
type AgeDetail struct {
	Years  int
	Months int
	Days   int
}

// String returns a human-readable string representation of the age.
// String은 나이를 사람이 읽기 쉬운 문자열로 반환합니다.
//
// Examples:
//   - "35 years 5 months 14 days"
//   - "1 year 2 months"
//   - "6 months 10 days"
func (a *AgeDetail) String() string {
	var result string
	if a.Years > 0 {
		if a.Years == 1 {
			result = "1 year"
		} else {
			result = fmt.Sprintf("%d years", a.Years)
		}
		if a.Months > 0 {
			if a.Months == 1 {
				result += " 1 month"
			} else {
				result += fmt.Sprintf(" %d months", a.Months)
			}
		}
		if a.Days > 0 {
			if a.Days == 1 {
				result += " 1 day"
			} else {
				result += fmt.Sprintf(" %d days", a.Days)
			}
		}
	} else if a.Months > 0 {
		if a.Months == 1 {
			result = "1 month"
		} else {
			result = fmt.Sprintf("%d months", a.Months)
		}
		if a.Days > 0 {
			if a.Days == 1 {
				result += " 1 day"
			} else {
				result += fmt.Sprintf(" %d days", a.Days)
			}
		}
	} else {
		if a.Days == 1 {
			result = "1 day"
		} else {
			result = fmt.Sprintf("%d days", a.Days)
		}
	}
	return result
}

// Humanize returns a short human-readable string representation of the age.
// Humanize는 나이를 짧은 사람이 읽기 쉬운 문자열로 반환합니다.
//
// Examples:
//   - "35y 5m 14d"
//   - "1y 2m"
//   - "6m 10d"
func (a *AgeDetail) Humanize() string {
	var result string
	if a.Years > 0 {
		result = fmt.Sprintf("%dy", a.Years)
		if a.Months > 0 {
			result += fmt.Sprintf(" %dm", a.Months)
		}
		if a.Days > 0 {
			result += fmt.Sprintf(" %dd", a.Days)
		}
	} else if a.Months > 0 {
		result = fmt.Sprintf("%dm", a.Months)
		if a.Days > 0 {
			result += fmt.Sprintf(" %dd", a.Days)
		}
	} else {
		result = fmt.Sprintf("%dd", a.Days)
	}
	return result
}
