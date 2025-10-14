package timeutil

import "time"

// WeekOfYear returns the ISO 8601 week number of the year (1-53).
// WeekOfYear는 ISO 8601 주 번호를 반환합니다 (1-53).
//
// Example / 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	week := timeutil.WeekOfYear(t)
//	fmt.Println(week) // 42
func WeekOfYear(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

// WeekOfMonth returns the week number of the month (1-6).
// WeekOfMonth는 월의 주 번호를 반환합니다 (1-6).
//
// Week numbering starts from the first Monday. Days before the first Monday are week 0, but returned as week 1.
// 주 번호는 첫 번째 월요일부터 시작합니다. 첫 번째 월요일 이전의 날들은 주 0이지만 주 1로 반환됩니다.
//
// Example / 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	week := timeutil.WeekOfMonth(t)
//	fmt.Println(week) // 3
func WeekOfMonth(t time.Time) int {
	// Get first day of the month
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())

	// Get the weekday of the first day (0 = Sunday, 6 = Saturday)
	firstWeekday := int(firstDay.Weekday())

	// Adjust so Monday = 0, Sunday = 6
	if firstWeekday == 0 {
		firstWeekday = 6
	} else {
		firstWeekday--
	}

	// Get current day of month
	day := t.Day()

	// Calculate week number
	// Days before first Monday are in "week 0" but we call it week 1
	week := (day + firstWeekday - 1) / DaysPerWeek + 1

	if week < 1 {
		week = 1
	}

	return week
}

// DaysInMonth returns the number of days in the month (28-31).
// DaysInMonth는 월의 일 수를 반환합니다 (28-31).
//
// Example / 예제:
//
//	t := time.Date(2025, 2, 1, 0, 0, 0, 0, KST)
//	days := timeutil.DaysInMonth(t)
//	fmt.Println(days) // 28
func DaysInMonth(t time.Time) int {
	// Get the first day of next month and subtract one day
	firstOfNextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	lastOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
	return lastOfMonth.Day()
}

// DaysInYear returns the number of days in the year (365 or 366).
// DaysInYear는 년의 일 수를 반환합니다 (365 또는 366).
//
// Example / 예제:
//
//	t := time.Date(2024, 1, 1, 0, 0, 0, 0, KST)
//	days := timeutil.DaysInYear(t)
//	fmt.Println(days) // 366 (leap year)
func DaysInYear(t time.Time) int {
	if IsLeapYear(t) {
		return DaysPerLeapYear
	}
	return DaysPerYear
}
