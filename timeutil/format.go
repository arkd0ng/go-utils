package timeutil

import (
	"strings"
	"time"
)

// FormatISO8601 formats a time in ISO8601 format.
// FormatISO8601은 시간을 ISO8601 포맷으로 포맷합니다.
func FormatISO8601(t time.Time) string {
	return t.In(defaultLocation).Format(ISO8601Layout)
}

// FormatRFC3339 formats a time in RFC3339 format.
// FormatRFC3339는 시간을 RFC3339 포맷으로 포맷합니다.
func FormatRFC3339(t time.Time) string {
	return t.In(defaultLocation).Format(RFC3339Layout)
}

// FormatDate formats a time as date only (YYYY-MM-DD).
// FormatDate는 시간을 날짜만으로 포맷합니다 (YYYY-MM-DD).
func FormatDate(t time.Time) string {
	return t.In(defaultLocation).Format(DateLayout)
}

// FormatDateTime formats a time as date and time (YYYY-MM-DD HH:mm:ss).
// FormatDateTime은 시간을 날짜 및 시간으로 포맷합니다 (YYYY-MM-DD HH:mm:ss).
func FormatDateTime(t time.Time) string {
	return t.In(defaultLocation).Format(DateTimeLayout)
}

// FormatTime formats a time as time only (HH:mm:ss).
// FormatTime은 시간을 시간만으로 포맷합니다 (HH:mm:ss).
func FormatTime(t time.Time) string {
	return t.In(defaultLocation).Format(TimeLayout)
}

// Format formats a time using custom format tokens.
// Format은 커스텀 포맷 토큰을 사용하여 시간을 포맷합니다.
//
// Supported tokens
// 지원되는 토큰:
//
//	YYYY - 4-digit year
//	YY   - 2-digit year
//	MM   - 2-digit month
//	M    - 1 or 2-digit month
//	DD   - 2-digit day
//	D    - 1 or 2-digit day
//	HH   - 2-digit hour (24h)
//	hh   - 2-digit hour (12h)
//	mm   - 2-digit minute
//	ss   - 2-digit second
//
// Example
// 예제:
//
// timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
// timeutil.Format(time.Now(), "YYYY년 MM월 DD일")
func Format(t time.Time, layout string) string {
	t = t.In(defaultLocation)
	goLayout := layout
	for token, goToken := range customFormatTokens {
		goLayout = strings.ReplaceAll(goLayout, token, goToken)
	}
	return t.Format(goLayout)
}

// FormatCustom is an alias for Format.
// FormatCustom은 Format의 별칭입니다.
func FormatCustom(t time.Time, layout string) string {
	return Format(t, layout)
}

// FormatWithTimezone formats a time in a specific timezone.
// FormatWithTimezone은 특정 타임존에서 시간을 포맷합니다.
func FormatWithTimezone(t time.Time, tz string) (string, error) {
	converted, err := ConvertTimezone(t, tz)
	if err != nil {
		return "", err
	}
	return FormatDateTime(converted), nil
}

// FormatKorean formats a time in Korean format (YYYY년 MM월 DD일 HH시 mm분 ss초).
// FormatKorean은 시간을 한국어 포맷으로 포맷합니다 (YYYY년 MM월 DD일 HH시 mm분 ss초).
func FormatKorean(t time.Time) string {
	return Format(t, "YYYY년 MM월 DD일 HH시 mm분 ss초")
}

// FormatKoreanDate formats a time in Korean date format (YYYY년 MM월 DD일).
// FormatKoreanDate는 시간을 한국어 날짜 포맷으로 포맷합니다 (YYYY년 MM월 DD일).
func FormatKoreanDate(t time.Time) string {
	return Format(t, "YYYY년 MM월 DD일")
}

// WeekdayKorean returns the Korean name of the weekday.
// WeekdayKorean은 요일의 한글 이름을 반환합니다.
//
// Returns
// 반환값:
//   - "일요일" for Sunday
//   - "월요일" for Monday
//   - "화요일" for Tuesday
//   - "수요일" for Wednesday
//   - "목요일" for Thursday
//   - "금요일" for Friday
//   - "토요일" for Saturday
//
// Example
// 예제:
//
// t := time.Now()
// fmt.Println(timeutil.WeekdayKorean(t))  // Output: 월요일
func WeekdayKorean(t time.Time) string {
	weekdays := []string{
		"일요일", // Sunday
		"월요일", // Monday
		"화요일", // Tuesday
		"수요일", // Wednesday
		"목요일", // Thursday
		"금요일", // Friday
		"토요일", // Saturday
	}
	return weekdays[t.Weekday()]
}

// WeekdayKoreanShort returns the short Korean name of the weekday.
// WeekdayKoreanShort는 요일의 짧은 한글 이름을 반환합니다.
//
// Returns
// 반환값:
//   - "일" for Sunday
//   - "월" for Monday
//   - "화" for Tuesday
//   - "수" for Wednesday
//   - "목" for Thursday
//   - "금" for Friday
//   - "토" for Saturday
//
// Example
// 예제:
//
// t := time.Now()
// fmt.Println(timeutil.WeekdayKoreanShort(t))  // Output: 월
func WeekdayKoreanShort(t time.Time) string {
	weekdays := []string{
		"일", // Sunday
		"월", // Monday
		"화", // Tuesday
		"수", // Wednesday
		"목", // Thursday
		"금", // Friday
		"토", // Saturday
	}
	return weekdays[t.Weekday()]
}

// FormatKoreanDateTime formats a time in Korean format with weekday.
// FormatKoreanDateTime은 요일을 포함한 한국어 포맷으로 시간을 포맷합니다.
//
// Format
// 포맷: YYYY년 MM월 DD일 (요일) HH시 mm분 ss초
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 15, 30, 0, 0, time.UTC)
//
// fmt.Println(timeutil.FormatKoreanDateTime(t))
// Output: 2025년 10월 14일 (화요일) 15시 30분 00초
func FormatKoreanDateTime(t time.Time) string {
	t = t.In(defaultLocation)
	return t.Format("2006년 01월 02일") + " (" + WeekdayKorean(t) + ") " + t.Format("15시 04분 05초")
}

// FormatKoreanDateWithWeekday formats a date in Korean format with weekday.
// FormatKoreanDateWithWeekday는 요일을 포함한 한국어 날짜 포맷으로 포맷합니다.
//
// Format
// 포맷: YYYY년 MM월 DD일 (요일)
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, time.UTC)
//
// fmt.Println(timeutil.FormatKoreanDateWithWeekday(t))
// Output: 2025년 10월 14일 (화요일)
func FormatKoreanDateWithWeekday(t time.Time) string {
	t = t.In(defaultLocation)
	return t.Format("2006년 01월 02일") + " (" + WeekdayKorean(t) + ")"
}

// FormatKoreanDateShort formats a date in Korean format with short weekday.
// FormatKoreanDateShort는 짧은 요일을 포함한 한국어 날짜 포맷으로 포맷합니다.
//
// Format
// 포맷: YYYY년 MM월 DD일 (요일)
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, time.UTC)
//
// fmt.Println(timeutil.FormatKoreanDateShort(t))
// Output: 2025년 10월 14일 (화)
func FormatKoreanDateShort(t time.Time) string {
	t = t.In(defaultLocation)
	return t.Format("2006년 01월 02일") + " (" + WeekdayKoreanShort(t) + ")"
}
