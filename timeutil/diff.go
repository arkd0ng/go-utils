package timeutil

import "time"

// SubTime calculates the difference between two times and returns a TimeDiff.
// SubTime은 두 시간의 차이를 계산하고 TimeDiff를 반환합니다.
//
// The difference is calculated as t2 - t1.
// 차이는 t2 - t1로 계산됩니다.
//
// Example / 예제:
//
//	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)
//	diff := timeutil.SubTime(start, end)
//	fmt.Println(diff.Days())    // 2.270833...
//	fmt.Println(diff.String())  // "2 days 6 hours 30 minutes"
func SubTime(t1, t2 time.Time) *TimeDiff {
	return &TimeDiff{Duration: t2.Sub(t1)}
}

// DiffInSeconds returns the difference between two times in seconds.
// DiffInSeconds는 두 시간의 차이를 초 단위로 반환합니다.
//
// Example / 예제:
//
//	seconds := timeutil.DiffInSeconds(start, end)
func DiffInSeconds(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Seconds()
}

// DiffInMinutes returns the difference between two times in minutes.
// DiffInMinutes는 두 시간의 차이를 분 단위로 반환합니다.
//
// Example / 예제:
//
//	minutes := timeutil.DiffInMinutes(start, end)
func DiffInMinutes(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Minutes()
}

// DiffInHours returns the difference between two times in hours.
// DiffInHours는 두 시간의 차이를 시간 단위로 반환합니다.
//
// Example / 예제:
//
//	hours := timeutil.DiffInHours(start, end)
func DiffInHours(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Hours()
}

// DiffInDays returns the difference between two times in days.
// DiffInDays는 두 시간의 차이를 일 단위로 반환합니다.
//
// Example / 예제:
//
//	days := timeutil.DiffInDays(start, end)
func DiffInDays(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Hours() / 24
}

// DiffInWeeks returns the difference between two times in weeks.
// DiffInWeeks는 두 시간의 차이를 주 단위로 반환합니다.
//
// Example / 예제:
//
//	weeks := timeutil.DiffInWeeks(start, end)
func DiffInWeeks(t1, t2 time.Time) float64 {
	return DiffInDays(t1, t2) / 7
}

// DiffInMonths returns the difference between two times in months.
// DiffInMonths는 두 시간의 차이를 월 단위로 반환합니다.
//
// This function calculates the difference in calendar months.
// 이 함수는 달력 월 단위의 차이를 계산합니다.
//
// Example / 예제:
//
//	months := timeutil.DiffInMonths(start, end)
func DiffInMonths(t1, t2 time.Time) int {
	years := t2.Year() - t1.Year()
	months := int(t2.Month()) - int(t1.Month())

	// Adjust if day of month is earlier / 일이 더 이르면 조정
	if t2.Day() < t1.Day() {
		months--
	}

	return years*MonthsPerYear + months
}

// DiffInYears returns the difference between two times in years.
// DiffInYears는 두 시간의 차이를 년 단위로 반환합니다.
//
// This function calculates the difference in calendar years.
// 이 함수는 달력 년 단위의 차이를 계산합니다.
//
// Example / 예제:
//
//	years := timeutil.DiffInYears(start, end)
func DiffInYears(t1, t2 time.Time) int {
	years := t2.Year() - t1.Year()

	// Adjust if month/day is earlier / 월/일이 더 이르면 조정
	if t2.Month() < t1.Month() || (t2.Month() == t1.Month() && t2.Day() < t1.Day()) {
		years--
	}

	return years
}
