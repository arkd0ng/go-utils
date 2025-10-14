package timeutil

import "time"

// IsBefore checks if t1 is before t2.
// IsBefore는 t1이 t2보다 이전인지 확인합니다.
func IsBefore(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// IsAfter checks if t1 is after t2.
// IsAfter는 t1이 t2보다 이후인지 확인합니다.
func IsAfter(t1, t2 time.Time) bool {
	return t1.After(t2)
}

// IsBetween checks if t is between start and end.
// IsBetween은 t가 start와 end 사이에 있는지 확인합니다.
func IsBetween(t, start, end time.Time) bool {
	return (t.After(start) || t.Equal(start)) && (t.Before(end) || t.Equal(end))
}

// IsToday checks if t is today in KST.
// IsToday는 t가 KST로 오늘인지 확인합니다.
func IsToday(t time.Time) bool {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	return IsSameDay(t, now)
}

// IsYesterday checks if t is yesterday in KST.
// IsYesterday는 t가 KST로 어제인지 확인합니다.
func IsYesterday(t time.Time) bool {
	yesterday := time.Now().In(defaultLocation).AddDate(0, 0, -1)
	t = t.In(defaultLocation)
	return IsSameDay(t, yesterday)
}

// IsTomorrow checks if t is tomorrow in KST.
// IsTomorrow는 t가 KST로 내일인지 확인합니다.
func IsTomorrow(t time.Time) bool {
	tomorrow := time.Now().In(defaultLocation).AddDate(0, 0, 1)
	t = t.In(defaultLocation)
	return IsSameDay(t, tomorrow)
}

// IsThisWeek checks if t is in the current week in KST.
// IsThisWeek는 t가 KST로 이번 주인지 확인합니다.
func IsThisWeek(t time.Time) bool {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	return IsSameWeek(t, now)
}

// IsThisMonth checks if t is in the current month in KST.
// IsThisMonth는 t가 KST로 이번 달인지 확인합니다.
func IsThisMonth(t time.Time) bool {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	return IsSameMonth(t, now)
}

// IsThisYear checks if t is in the current year in KST.
// IsThisYear는 t가 KST로 올해인지 확인합니다.
func IsThisYear(t time.Time) bool {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	return IsSameYear(t, now)
}

// IsWeekend checks if t is on a weekend (Saturday or Sunday).
// IsWeekend는 t가 주말인지 (토요일 또는 일요일) 확인합니다.
func IsWeekend(t time.Time) bool {
	t = t.In(defaultLocation)
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsWeekday checks if t is on a weekday (Monday to Friday).
// IsWeekday는 t가 평일인지 (월요일부터 금요일) 확인합니다.
func IsWeekday(t time.Time) bool {
	return !IsWeekend(t)
}

// IsSameDay checks if t1 and t2 are on the same day.
// IsSameDay는 t1과 t2가 같은 날인지 확인합니다.
func IsSameDay(t1, t2 time.Time) bool {
	t1 = t1.In(defaultLocation)
	t2 = t2.In(defaultLocation)
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsSameWeek checks if t1 and t2 are in the same week.
// IsSameWeek는 t1과 t2가 같은 주인지 확인합니다.
func IsSameWeek(t1, t2 time.Time) bool {
	t1 = t1.In(defaultLocation)
	t2 = t2.In(defaultLocation)
	start1 := StartOfWeek(t1)
	start2 := StartOfWeek(t2)
	return IsSameDay(start1, start2)
}

// IsSameMonth checks if t1 and t2 are in the same month.
// IsSameMonth는 t1과 t2가 같은 달인지 확인합니다.
func IsSameMonth(t1, t2 time.Time) bool {
	t1 = t1.In(defaultLocation)
	t2 = t2.In(defaultLocation)
	return t1.Year() == t2.Year() && t1.Month() == t2.Month()
}

// IsSameYear checks if t1 and t2 are in the same year.
// IsSameYear는 t1과 t2가 같은 년인지 확인합니다.
func IsSameYear(t1, t2 time.Time) bool {
	t1 = t1.In(defaultLocation)
	t2 = t2.In(defaultLocation)
	return t1.Year() == t2.Year()
}

// IsLeapYear checks if t is in a leap year.
// IsLeapYear는 t가 윤년인지 확인합니다.
func IsLeapYear(t time.Time) bool {
	year := t.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// IsPast checks if t is in the past.
// IsPast는 t가 과거인지 확인합니다.
func IsPast(t time.Time) bool {
	return t.Before(time.Now())
}

// IsFuture checks if t is in the future.
// IsFuture는 t가 미래인지 확인합니다.
func IsFuture(t time.Time) bool {
	return t.After(time.Now())
}
