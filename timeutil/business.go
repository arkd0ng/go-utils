package timeutil

import (
	"sync"
	"time"
)

var (
	// holidays stores custom holidays
	// holidays는 커스텀 공휴일을 저장합니다
	holidays   = make(map[string]bool)
	holidaysMu sync.RWMutex
)

// SetHolidays sets custom holidays.
// SetHolidays는 커스텀 공휴일을 설정합니다.
//
// Example / 예제:
//
//	holidays := []time.Time{
//	    time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),  // New Year
//	    time.Date(2025, 12, 25, 0, 0, 0, 0, time.UTC), // Christmas
//	}
//	timeutil.SetHolidays(holidays)
func SetHolidays(dates []time.Time) {
	holidaysMu.Lock()
	defer holidaysMu.Unlock()

	holidays = make(map[string]bool)
	for _, date := range dates {
		date = date.In(defaultLocation)
		key := FormatDate(date)
		holidays[key] = true
	}
}

// GetHolidays returns the list of custom holidays.
// GetHolidays는 커스텀 공휴일 목록을 반환합니다.
func GetHolidays() []time.Time {
	holidaysMu.RLock()
	defer holidaysMu.RUnlock()

	result := make([]time.Time, 0, len(holidays))
	for key := range holidays {
		date, _ := ParseDate(key)
		result = append(result, date)
	}
	return result
}

// ClearHolidays clears all custom holidays.
// ClearHolidays는 모든 커스텀 공휴일을 지웁니다.
func ClearHolidays() {
	holidaysMu.Lock()
	defer holidaysMu.Unlock()
	holidays = make(map[string]bool)
}

// IsHoliday checks if t is a custom holiday.
// IsHoliday는 t가 커스텀 공휴일인지 확인합니다.
func IsHoliday(t time.Time) bool {
	holidaysMu.RLock()
	defer holidaysMu.RUnlock()

	t = t.In(defaultLocation)
	key := FormatDate(t)
	return holidays[key]
}

// IsBusinessDay checks if t is a business day (Monday-Friday, excluding holidays).
// IsBusinessDay는 t가 영업일인지 확인합니다 (월-금, 공휴일 제외).
func IsBusinessDay(t time.Time) bool {
	t = t.In(defaultLocation)

	// Check if weekend / 주말 확인
	if IsWeekend(t) {
		return false
	}

	// Check if holiday / 공휴일 확인
	if IsHoliday(t) {
		return false
	}

	return true
}

// AddBusinessDays adds the specified number of business days to a time.
// AddBusinessDays는 시간에 지정된 영업일을 더합니다.
//
// Business days are Monday-Friday, excluding weekends and custom holidays.
// 영업일은 월-금이며 주말과 커스텀 공휴일을 제외합니다.
//
// Example / 예제:
//
//	nextBusinessDay := timeutil.AddBusinessDays(time.Now(), 5)
func AddBusinessDays(t time.Time, days int) time.Time {
	t = t.In(defaultLocation)

	if days == 0 {
		return t
	}

	direction := 1
	if days < 0 {
		direction = -1
		days = -days
	}

	count := 0
	current := t

	for count < days {
		current = current.AddDate(0, 0, direction)
		if IsBusinessDay(current) {
			count++
		}
	}

	return current
}

// NextBusinessDay returns the next business day.
// NextBusinessDay는 다음 영업일을 반환합니다.
func NextBusinessDay(t time.Time) time.Time {
	t = t.In(defaultLocation)
	next := t.AddDate(0, 0, 1)

	for !IsBusinessDay(next) {
		next = next.AddDate(0, 0, 1)
	}

	return next
}

// PreviousBusinessDay returns the previous business day.
// PreviousBusinessDay는 이전 영업일을 반환합니다.
func PreviousBusinessDay(t time.Time) time.Time {
	t = t.In(defaultLocation)
	prev := t.AddDate(0, 0, -1)

	for !IsBusinessDay(prev) {
		prev = prev.AddDate(0, 0, -1)
	}

	return prev
}

// CountBusinessDays counts the number of business days between two dates.
// CountBusinessDays는 두 날짜 사이의 영업일 수를 계산합니다.
//
// The count is inclusive of start but exclusive of end.
// 카운트는 start를 포함하지만 end는 제외합니다.
//
// Example / 예제:
//
//	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC)
//	days := timeutil.CountBusinessDays(start, end)
func CountBusinessDays(start, end time.Time) int {
	start = start.In(defaultLocation)
	end = end.In(defaultLocation)

	// Ensure start is before end / start가 end보다 앞에 있도록 확인
	if start.After(end) {
		start, end = end, start
	}

	// Truncate to start of day / 하루의 시작으로 절단
	start = StartOfDay(start)
	end = StartOfDay(end)

	count := 0
	current := start

	for current.Before(end) {
		if IsBusinessDay(current) {
			count++
		}
		current = current.AddDate(0, 0, 1)
	}

	return count
}

// AddKoreanHolidays adds common Korean public holidays for a given year.
// AddKoreanHolidays는 주어진 년도의 일반적인 한국 공휴일을 추가합니다.
//
// This function adds fixed holidays only (New Year's Day, Independence Movement Day,
// Liberation Day, National Foundation Day, Hangul Day, Christmas).
// 이 함수는 고정 공휴일만 추가합니다 (신정, 3.1절, 광복절, 개천절, 한글날, 크리스마스).
//
// Note: Lunar calendar holidays (Seollal, Chuseok) are not included.
// 참고: 음력 공휴일 (설날, 추석)은 포함되지 않습니다.
//
// Example / 예제:
//
//	timeutil.AddKoreanHolidays(2025)
func AddKoreanHolidays(year int) {
	koreanHolidays := []time.Time{
		time.Date(year, 1, 1, 0, 0, 0, 0, KST),   // New Year's Day / 신정
		time.Date(year, 3, 1, 0, 0, 0, 0, KST),   // Independence Movement Day / 3.1절
		time.Date(year, 5, 5, 0, 0, 0, 0, KST),   // Children's Day / 어린이날
		time.Date(year, 6, 6, 0, 0, 0, 0, KST),   // Memorial Day / 현충일
		time.Date(year, 8, 15, 0, 0, 0, 0, KST),  // Liberation Day / 광복절
		time.Date(year, 10, 3, 0, 0, 0, 0, KST),  // National Foundation Day / 개천절
		time.Date(year, 10, 9, 0, 0, 0, 0, KST),  // Hangul Day / 한글날
		time.Date(year, 12, 25, 0, 0, 0, 0, KST), // Christmas / 크리스마스
	}

	// Append to existing holidays / 기존 공휴일에 추가
	existing := GetHolidays()
	all := append(existing, koreanHolidays...)
	SetHolidays(all)
}
