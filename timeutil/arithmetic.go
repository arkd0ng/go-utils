package timeutil

import "time"

// AddSeconds adds the specified number of seconds to a time.
// AddSeconds는 시간에 지정된 초를 더합니다.
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// AddMinutes adds the specified number of minutes to a time.
// AddMinutes는 시간에 지정된 분을 더합니다.
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddHours adds the specified number of hours to a time.
// AddHours는 시간에 지정된 시간을 더합니다.
func AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

// AddDays adds the specified number of days to a time.
// AddDays는 시간에 지정된 일을 더합니다.
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddWeeks adds the specified number of weeks to a time.
// AddWeeks는 시간에 지정된 주를 더합니다.
func AddWeeks(t time.Time, weeks int) time.Time {
	return t.AddDate(0, 0, weeks*DaysPerWeek)
}

// AddMonths adds the specified number of months to a time.
// AddMonths는 시간에 지정된 월을 더합니다.
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYears adds the specified number of years to a time.
// AddYears는 시간에 지정된 년을 더합니다.
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// StartOfDay returns the start of the day (00:00:00) in KST.
// StartOfDay는 KST로 하루의 시작 (00:00:00)을 반환합니다.
func StartOfDay(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, defaultLocation)
}

// EndOfDay returns the end of the day (23:59:59) in KST.
// EndOfDay는 KST로 하루의 끝 (23:59:59)을 반환합니다.
func EndOfDay(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, defaultLocation)
}

// StartOfWeek returns the start of the week (Monday 00:00:00) in KST.
// StartOfWeek는 KST로 주의 시작 (월요일 00:00:00)을 반환합니다.
func StartOfWeek(t time.Time) time.Time {
	t = t.In(defaultLocation)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7 // Sunday = 7
	}
	daysToMonday := weekday - 1
	return StartOfDay(t.AddDate(0, 0, -daysToMonday))
}

// EndOfWeek returns the end of the week (Sunday 23:59:59) in KST.
// EndOfWeek는 KST로 주의 끝 (일요일 23:59:59)을 반환합니다.
func EndOfWeek(t time.Time) time.Time {
	t = t.In(defaultLocation)
	weekday := int(t.Weekday())
	if weekday == 0 {
		return EndOfDay(t)
	}
	daysToSunday := 7 - weekday
	return EndOfDay(t.AddDate(0, 0, daysToSunday))
}

// StartOfMonth returns the start of the month (1st day 00:00:00) in KST.
// StartOfMonth는 KST로 월의 시작 (1일 00:00:00)을 반환합니다.
func StartOfMonth(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, defaultLocation)
}

// EndOfMonth returns the end of the month (last day 23:59:59) in KST.
// EndOfMonth는 KST로 월의 끝 (마지막 날 23:59:59)을 반환합니다.
func EndOfMonth(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999999999, defaultLocation)
}

// StartOfYear returns the start of the year (Jan 1 00:00:00) in KST.
// StartOfYear는 KST로 년의 시작 (1월 1일 00:00:00)을 반환합니다.
func StartOfYear(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, defaultLocation)
}

// EndOfYear returns the end of the year (Dec 31 23:59:59) in KST.
// EndOfYear는 KST로 년의 끝 (12월 31일 23:59:59)을 반환합니다.
func EndOfYear(t time.Time) time.Time {
	t = t.In(defaultLocation)
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, defaultLocation)
}

// StartOfQuarter returns the start of the quarter in KST.
// StartOfQuarter는 KST로 분기의 시작을 반환합니다.
func StartOfQuarter(t time.Time) time.Time {
	t = t.In(defaultLocation)
	month := t.Month()
	var quarterMonth time.Month
	switch {
	case month >= 1 && month <= 3:
		quarterMonth = 1
	case month >= 4 && month <= 6:
		quarterMonth = 4
	case month >= 7 && month <= 9:
		quarterMonth = 7
	default:
		quarterMonth = 10
	}
	return time.Date(t.Year(), quarterMonth, 1, 0, 0, 0, 0, defaultLocation)
}

// EndOfQuarter returns the end of the quarter in KST.
// EndOfQuarter는 KST로 분기의 끝을 반환합니다.
func EndOfQuarter(t time.Time) time.Time {
	start := StartOfQuarter(t)
	return EndOfMonth(start.AddDate(0, 2, 0))
}
