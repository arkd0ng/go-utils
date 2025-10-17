package timeutil

import (
	"fmt"
	"time"
)

// String versions of time functions that accept string parameters instead of time.Time.
// These functions automatically parse the input strings using ParseAny.
// 문자열 매개변수를 받는 시간 함수의 문자열 버전입니다.
// 이 함수들은 ParseAny를 사용하여 입력 문자열을 자동으로 파싱합니다.

// SubTimeString calculates the time difference between two time strings.
// SubTimeString은 두 시간 문자열 사이의 시간 차이를 계산합니다.
//
// Example
// 예제:
//
//	diff, err := timeutil.SubTimeString("2024-10-04 08:34:42", "2024-10-14 14:56:23")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(diff.String()) // "10 days 6 hours 21 minutes"
func SubTimeString(s1, s2 string) (*TimeDiff, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return nil, fmt.Errorf("failed to parse first time string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return nil, fmt.Errorf("failed to parse second time string: %w", err)
	}

	return SubTime(t1, t2), nil
}

// DiffInDaysString returns the number of days between two time strings.
// DiffInDaysString은 두 시간 문자열 사이의 일수를 반환합니다.
func DiffInDaysString(s1, s2 string) (float64, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return 0, fmt.Errorf("failed to parse first time string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return 0, fmt.Errorf("failed to parse second time string: %w", err)
	}

	return DiffInDays(t1, t2), nil
}

// DiffInHoursString returns the number of hours between two time strings.
// DiffInHoursString은 두 시간 문자열 사이의 시간수를 반환합니다.
func DiffInHoursString(s1, s2 string) (float64, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return 0, fmt.Errorf("failed to parse first time string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return 0, fmt.Errorf("failed to parse second time string: %w", err)
	}

	return DiffInHours(t1, t2), nil
}

// DiffInMinutesString returns the number of minutes between two time strings.
// DiffInMinutesString은 두 시간 문자열 사이의 분수를 반환합니다.
func DiffInMinutesString(s1, s2 string) (float64, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return 0, fmt.Errorf("failed to parse first time string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return 0, fmt.Errorf("failed to parse second time string: %w", err)
	}

	return DiffInMinutes(t1, t2), nil
}

// AgeString calculates the age from a birth date string.
// AgeString은 생년월일 문자열로부터 나이를 계산합니다.
//
// Example
// 예제:
//
//	age, err := timeutil.AgeString("1990-01-15")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(age.Years) // 35 (assuming current year is 2025)
func AgeString(birthDate string) (*AgeDetail, error) {
	t, err := ParseAny(birthDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse birth date: %w", err)
	}

	return Age(t), nil
}

// AgeInYearsString calculates the age in years from a birth date string.
// AgeInYearsString은 생년월일 문자열로부터 나이(년)를 계산합니다.
func AgeInYearsString(birthDate string) (int, error) {
	t, err := ParseAny(birthDate)
	if err != nil {
		return 0, fmt.Errorf("failed to parse birth date: %w", err)
	}

	return AgeInYears(t), nil
}

// RelativeTimeString returns a human-readable relative time string.
// RelativeTimeString은 사람이 읽기 쉬운 상대 시간 문자열을 반환합니다.
//
// Example
// 예제:
//
//	rel, err := timeutil.RelativeTimeString("2024-10-13 15:30:00")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(rel) // "1 day ago"
func RelativeTimeString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return RelativeTime(t), nil
}

// IsBusinessDayString checks if a date string is a business day (Monday-Friday, not a holiday).
// IsBusinessDayString은 날짜 문자열이 영업일인지 확인합니다 (월-금, 공휴일 제외).
func IsBusinessDayString(s string) (bool, error) {
	t, err := ParseAny(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse date string: %w", err)
	}

	return IsBusinessDay(t), nil
}

// IsWeekendString checks if a date string falls on a weekend (Saturday or Sunday).
// IsWeekendString은 날짜 문자열이 주말인지 확인합니다 (토요일 또는 일요일).
func IsWeekendString(s string) (bool, error) {
	t, err := ParseAny(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse date string: %w", err)
	}

	return IsWeekend(t), nil
}

// AddDaysString adds a number of days to a date string.
// AddDaysString은 날짜 문자열에 일수를 더합니다.
//
// Example
// 예제:
//
//	result, err := timeutil.AddDaysString("2024-10-14", 7)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result) // 2024-10-21 00:00:00 +0900 KST
func AddDaysString(s string, days int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return AddDays(t, days), nil
}

// AddHoursString adds a number of hours to a datetime string.
// AddHoursString은 날짜시간 문자열에 시간수를 더합니다.
func AddHoursString(s string, hours int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime string: %w", err)
	}

	return AddHours(t, hours), nil
}

// AddMinutesString adds a number of minutes to a datetime string.
// AddMinutesString은 날짜시간 문자열에 분수를 더합니다.
func AddMinutesString(s string, minutes int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime string: %w", err)
	}

	return AddMinutes(t, minutes), nil
}

// SubDaysString subtracts a number of days from a date string.
// SubDaysString은 날짜 문자열에서 일수를 뺍니다.
func SubDaysString(s string, days int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return AddDays(t, -days), nil
}

// SubHoursString subtracts a number of hours from a datetime string.
// SubHoursString은 날짜시간 문자열에서 시간수를 뺍니다.
func SubHoursString(s string, hours int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime string: %w", err)
	}

	return AddHours(t, -hours), nil
}

// SubMinutesString subtracts a number of minutes from a datetime string.
// SubMinutesString은 날짜시간 문자열에서 분수를 뺍니다.
func SubMinutesString(s string, minutes int) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime string: %w", err)
	}

	return AddMinutes(t, -minutes), nil
}

// FormatString parses a time string and formats it with the given layout.
// FormatString은 시간 문자열을 파싱하여 주어진 레이아웃으로 포맷합니다.
//
// Example
// 예제:
//
//	result, err := timeutil.FormatString("2024-10-14 15:30:00", "2006-01-02")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result) // "2024-10-14"
func FormatString(s, layout string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return t.Format(layout), nil
}

// FormatDateString parses a time string and formats it as a date (YYYY-MM-DD).
// FormatDateString은 시간 문자열을 파싱하여 날짜 형식으로 포맷합니다 (YYYY-MM-DD).
func FormatDateString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return FormatDate(t), nil
}

// FormatDateTimeString parses a time string and formats it as datetime (YYYY-MM-DD HH:mm:ss).
// FormatDateTimeString은 시간 문자열을 파싱하여 날짜시간 형식으로 포맷합니다 (YYYY-MM-DD HH:mm:ss).
func FormatDateTimeString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return FormatDateTime(t), nil
}

// FormatISO8601String parses a time string and formats it in ISO8601 format.
// FormatISO8601String은 시간 문자열을 파싱하여 ISO8601 형식으로 포맷합니다.
func FormatISO8601String(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse time string: %w", err)
	}

	return FormatISO8601(t), nil
}

// ConvertTimezoneString parses a time string and converts it to a different timezone.
// ConvertTimezoneString은 시간 문자열을 파싱하여 다른 타임존으로 변환합니다.
//
// Example
// 예제:
//
//	result, err := timeutil.ConvertTimezoneString("2024-10-14 15:30:00", "America/New_York")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result)
func ConvertTimezoneString(s, tz string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time string: %w", err)
	}

	return ConvertTimezone(t, tz)
}

// StartOfDayString returns the start of the day (00:00:00) for a date string.
// StartOfDayString은 날짜 문자열의 시작 시간(00:00:00)을 반환합니다.
func StartOfDayString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return StartOfDay(t), nil
}

// EndOfDayString returns the end of the day (23:59:59) for a date string.
// EndOfDayString은 날짜 문자열의 종료 시간(23:59:59)을 반환합니다.
func EndOfDayString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return EndOfDay(t), nil
}

// StartOfWeekString returns the start of the week (Monday 00:00:00) for a date string.
// StartOfWeekString은 날짜 문자열의 주 시작 시간(월요일 00:00:00)을 반환합니다.
func StartOfWeekString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return StartOfWeek(t), nil
}

// EndOfWeekString returns the end of the week (Sunday 23:59:59) for a date string.
// EndOfWeekString은 날짜 문자열의 주 종료 시간(일요일 23:59:59)을 반환합니다.
func EndOfWeekString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return EndOfWeek(t), nil
}

// StartOfMonthString returns the start of the month (day 1, 00:00:00) for a date string.
// StartOfMonthString은 날짜 문자열의 월 시작 시간(1일 00:00:00)을 반환합니다.
func StartOfMonthString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return StartOfMonth(t), nil
}

// EndOfMonthString returns the end of the month (last day, 23:59:59) for a date string.
// EndOfMonthString은 날짜 문자열의 월 종료 시간(마지막 날 23:59:59)을 반환합니다.
func EndOfMonthString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return EndOfMonth(t), nil
}

// StartOfYearString returns the start of the year (Jan 1, 00:00:00) for a date string.
// StartOfYearString은 날짜 문자열의 연 시작 시간(1월 1일 00:00:00)을 반환합니다.
func StartOfYearString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return StartOfYear(t), nil
}

// EndOfYearString returns the end of the year (Dec 31, 23:59:59) for a date string.
// EndOfYearString은 날짜 문자열의 연 종료 시간(12월 31일 23:59:59)을 반환합니다.
func EndOfYearString(s string) (time.Time, error) {
	t, err := ParseAny(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %w", err)
	}

	return EndOfYear(t), nil
}

// WeekdayString returns the weekday name for a date string.
// WeekdayString은 날짜 문자열의 요일 이름을 반환합니다.
func WeekdayString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return t.Weekday().String(), nil
}

// WeekdayKoreanString returns the Korean weekday name for a date string.
// WeekdayKoreanString은 날짜 문자열의 한글 요일 이름을 반환합니다.
func WeekdayKoreanString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return WeekdayKorean(t), nil
}

// WeekdayShortString returns the short weekday name for a date string.
// WeekdayShortString은 날짜 문자열의 짧은 요일 이름을 반환합니다.
func WeekdayShortString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return t.Format("Mon"), nil
}

// WeekdayShortKoreanString returns the short Korean weekday name for a date string.
// WeekdayShortKoreanString은 날짜 문자열의 짧은 한글 요일 이름을 반환합니다.
func WeekdayShortKoreanString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return WeekdayKoreanShort(t), nil
}

// WeekdayNumberString returns the weekday number (0=Sunday, 6=Saturday) for a date string.
// WeekdayNumberString은 날짜 문자열의 요일 번호(0=일요일, 6=토요일)를 반환합니다.
func WeekdayNumberString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return int(t.Weekday()), nil
}

// WeekOfYearString returns the ISO week number for a date string.
// WeekOfYearString은 날짜 문자열의 ISO 주 번호를 반환합니다.
func WeekOfYearString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return WeekOfYear(t), nil
}

// WeekOfMonthString returns the week number within the month for a date string.
// WeekOfMonthString은 날짜 문자열의 월 내 주 번호를 반환합니다.
func WeekOfMonthString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return WeekOfMonth(t), nil
}

// DaysInMonthString returns the number of days in the month for a date string.
// DaysInMonthString은 날짜 문자열의 월의 일수를 반환합니다.
func DaysInMonthString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return DaysInMonth(t), nil
}

// DaysInYearString returns the number of days in the year for a date string.
// DaysInYearString은 날짜 문자열의 연도의 일수를 반환합니다.
func DaysInYearString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return DaysInYear(t), nil
}

// MonthKoreanString returns the Korean month name for a date string.
// MonthKoreanString은 날짜 문자열의 한글 월 이름을 반환합니다.
func MonthKoreanString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return MonthKorean(t), nil
}

// MonthNameString returns the full month name for a date string.
// MonthNameString은 날짜 문자열의 전체 월 이름을 반환합니다.
func MonthNameString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return MonthName(t), nil
}

// MonthNameShortString returns the short month name for a date string.
// MonthNameShortString은 날짜 문자열의 짧은 월 이름을 반환합니다.
func MonthNameShortString(s string) (string, error) {
	t, err := ParseAny(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse date string: %w", err)
	}

	return MonthNameShort(t), nil
}

// QuarterString returns the quarter (1-4) for a date string.
// QuarterString은 날짜 문자열의 분기(1-4)를 반환합니다.
func QuarterString(s string) (int, error) {
	t, err := ParseAny(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date string: %w", err)
	}

	return Quarter(t), nil
}

// IsLeapYearString checks if the year in a date string is a leap year.
// IsLeapYearString은 날짜 문자열의 연도가 윤년인지 확인합니다.
func IsLeapYearString(s string) (bool, error) {
	t, err := ParseAny(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse date string: %w", err)
	}

	return IsLeapYear(t), nil
}

// IsSameDayString checks if two date strings are on the same day.
// IsSameDayString은 두 날짜 문자열이 같은 날인지 확인합니다.
func IsSameDayString(s1, s2 string) (bool, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return false, fmt.Errorf("failed to parse first date string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return false, fmt.Errorf("failed to parse second date string: %w", err)
	}

	return IsSameDay(t1, t2), nil
}

// IsBeforeString checks if the first date string is before the second.
// IsBeforeString은 첫 번째 날짜 문자열이 두 번째보다 이전인지 확인합니다.
func IsBeforeString(s1, s2 string) (bool, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return false, fmt.Errorf("failed to parse first date string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return false, fmt.Errorf("failed to parse second date string: %w", err)
	}

	return IsBefore(t1, t2), nil
}

// IsAfterString checks if the first date string is after the second.
// IsAfterString은 첫 번째 날짜 문자열이 두 번째보다 이후인지 확인합니다.
func IsAfterString(s1, s2 string) (bool, error) {
	t1, err := ParseAny(s1)
	if err != nil {
		return false, fmt.Errorf("failed to parse first date string: %w", err)
	}

	t2, err := ParseAny(s2)
	if err != nil {
		return false, fmt.Errorf("failed to parse second date string: %w", err)
	}

	return IsAfter(t1, t2), nil
}

// IsBetweenString checks if a date string is between two other dates.
// IsBetweenString은 날짜 문자열이 두 날짜 사이에 있는지 확인합니다.
func IsBetweenString(s, start, end string) (bool, error) {
	t, err := ParseAny(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse date string: %w", err)
	}

	startTime, err := ParseAny(start)
	if err != nil {
		return false, fmt.Errorf("failed to parse start date string: %w", err)
	}

	endTime, err := ParseAny(end)
	if err != nil {
		return false, fmt.Errorf("failed to parse end date string: %w", err)
	}

	return IsBetween(t, startTime, endTime), nil
}
