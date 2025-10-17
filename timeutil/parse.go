package timeutil

import (
	"fmt"
	"strings"
	"time"
)

// ParseISO8601 parses a time string in ISO8601 format.
// ParseISO8601은 ISO8601 포맷의 시간 문자열을 파싱합니다.
func ParseISO8601(s string) (time.Time, error) {
	t, err := time.Parse(ISO8601Layout, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse ISO8601: %w", err)
	}
	return t.In(defaultLocation), nil
}

// ParseRFC3339 parses a time string in RFC3339 format.
// ParseRFC3339는 RFC3339 포맷의 시간 문자열을 파싱합니다.
func ParseRFC3339(s string) (time.Time, error) {
	t, err := time.Parse(RFC3339Layout, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse RFC3339: %w", err)
	}
	return t.In(defaultLocation), nil
}

// ParseDate parses a date string (YYYY-MM-DD).
// ParseDate는 날짜 문자열을 파싱합니다 (YYYY-MM-DD).
func ParseDate(s string) (time.Time, error) {
	t, err := time.ParseInLocation(DateLayout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %w", err)
	}
	return t, nil
}

// ParseDateTime parses a datetime string (YYYY-MM-DD HH:mm:ss).
// ParseDateTime은 날짜시간 문자열을 파싱합니다 (YYYY-MM-DD HH:mm:ss).
func ParseDateTime(s string) (time.Time, error) {
	t, err := time.ParseInLocation(DateTimeLayout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime: %w", err)
	}
	return t, nil
}

// Parse attempts to parse a time string by auto-detecting the format.
// Parse는 포맷을 자동 감지하여 시간 문자열을 파싱합니다.
//
// Supported formats
// 지원되는 포맷:
//   - ISO8601: 2006-01-02T15:04:05Z07:00
//   - RFC3339: 2006-01-02T15:04:05Z07:00
//   - Date: 2006-01-02
//   - DateTime: 2006-01-02 15:04:05
func Parse(s string) (time.Time, error) {
	s = strings.TrimSpace(s)

	// Try ISO8601/RFC3339
	// ISO8601/RFC3339 시도
	if strings.Contains(s, "T") {
		t, err := time.Parse(time.RFC3339, s)
		if err == nil {
			return t.In(defaultLocation), nil
		}
		t, err = time.Parse(ISO8601Layout, s)
		if err == nil {
			return t.In(defaultLocation), nil
		}
	}

	// Try DateTime
	// DateTime 시도
	if strings.Contains(s, " ") && strings.Contains(s, ":") {
		t, err := time.ParseInLocation(DateTimeLayout, s, defaultLocation)
		if err == nil {
			return t, nil
		}
	}

	// Try Date
	// Date 시도
	if strings.Count(s, "-") == 2 {
		t, err := time.ParseInLocation(DateLayout, s, defaultLocation)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time string: %s", s)
}

// ParseWithTimezone parses a time string in a specific timezone.
// ParseWithTimezone은 특정 타임존에서 시간 문자열을 파싱합니다.
func ParseWithTimezone(s, tz string) (time.Time, error) {
	loc, err := loadTimezone(tz)
	if err != nil {
		return time.Time{}, err
	}

	// Try different formats
	// 다른 포맷 시도
	formats := []string{
		ISO8601Layout,
		RFC3339Layout,
		DateTimeLayout,
		DateLayout,
	}

	for _, layout := range formats {
		t, err := time.ParseInLocation(layout, s, loc)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time string in timezone %s: %s", tz, s)
}

// ParseWithLayout parses a time string with a custom layout.
// ParseWithLayout은 사용자 지정 레이아웃으로 시간 문자열을 파싱합니다.
//
// Example layouts
// 레이아웃 예제:
//   - "2006-01-02 15:04:05.000" for milliseconds / 밀리초용
//   - "2006-01-02 15:04:05.999999" for microseconds / 마이크로초용
//   - "2006/01/02" for date with slashes / 슬래시 구분 날짜
func ParseWithLayout(s, layout string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse with layout %s: %w", layout, err)
	}
	return t, nil
}

// ParseMillis parses a datetime string with milliseconds (YYYY-MM-DD HH:mm:ss.SSS).
// ParseMillis는 밀리초를 포함한 날짜시간 문자열을 파싱합니다 (YYYY-MM-DD HH:mm:ss.SSS).
func ParseMillis(s string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.000"
	t, err := time.ParseInLocation(layout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse milliseconds: %w", err)
	}
	return t, nil
}

// ParseMicros parses a datetime string with microseconds (YYYY-MM-DD HH:mm:ss.SSSSSS).
// ParseMicros는 마이크로초를 포함한 날짜시간 문자열을 파싱합니다 (YYYY-MM-DD HH:mm:ss.SSSSSS).
func ParseMicros(s string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.999999"
	t, err := time.ParseInLocation(layout, s, defaultLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse microseconds: %w", err)
	}
	return t, nil
}

// ParseAny attempts to parse a time string by trying all common formats.
// ParseAny는 모든 일반적인 포맷을 시도하여 시간 문자열을 파싱합니다.
//
// Supported formats
// 지원되는 포맷:
//   - ISO8601: 2006-01-02T15:04:05Z07:00
//   - RFC3339: 2006-01-02T15:04:05Z07:00
//   - DateTime with milliseconds: 2006-01-02 15:04:05.000
//   - DateTime with microseconds: 2006-01-02 15:04:05.999999
//   - DateTime with nanoseconds: 2006-01-02 15:04:05.999999999
//   - DateTime: 2006-01-02 15:04:05
//   - Date: 2006-01-02
//   - Date with slashes: 2006/01/02
//   - DateTime with slashes: 2006/01/02 15:04:05
//   - US format: 01/02/2006
//   - US format with time: 01/02/2006 15:04:05
//   - Month name: Jan 02, 2006
//   - Short month: 02-Jan-2006
//   - RFC822: 02 Jan 06 15:04 MST
//   - RFC1123: Mon, 02 Jan 2006 15:04:05 MST
//   - ANSIC: Mon Jan _2 15:04:05 2006
//   - UnixDate: Mon Jan _2 15:04:05 MST 2006
//   - RubyDate: Mon Jan 02 15:04:05 -0700 2006
func ParseAny(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, fmt.Errorf("empty time string")
	}

	// All formats to try
	// 시도할 모든 포맷
	formats := []string{
		// ISO8601 and RFC3339
		// ISO8601 및 RFC3339
		time.RFC3339,
		time.RFC3339Nano,
		ISO8601Layout,
		"2006-01-02T15:04:05Z0700",
		"2006-01-02T15:04:05",

		// DateTime with sub-seconds
		// 밀리초/마이크로초/나노초 포함
		"2006-01-02 15:04:05.999999999",
		"2006-01-02 15:04:05.999999",
		"2006-01-02 15:04:05.000",

		// DateTime
		// 날짜시간
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"2006-01-02 15:04",
		"2006/01/02 15:04",

		// Date only
		// 날짜만
		"2006-01-02",
		"2006/01/02",
		"01/02/2006", // US format
		"02-01-2006", // EU format
		"02.01.2006", // DE format

		// Korean formats
		// 한글 포맷
		"2006년 01월 02일 15시 04분 05초",
		"2006년 01월 02일 15시 04분",
		"2006년 01월 02일 15시",
		"2006년 01월 02일",
		"2006년 1월 2일 15시 4분 5초",
		"2006년 1월 2일 15시 4분",
		"2006년 1월 2일 15시",
		"2006년 1월 2일",
		"2006년 01월 02일 오후 3시 04분 05초",
		"2006년 01월 02일 오후 3시 04분",
		"2006년 01월 02일 오후 3시",
		"2006년 01월 02일 오전 9시 04분 05초",
		"2006년 01월 02일 오전 9시 04분",
		"2006년 01월 02일 오전 9시",

		// With month names
		// 월 이름 포함
		"Jan 02, 2006",
		"January 02, 2006",
		"02-Jan-2006",
		"02-January-2006",
		"02 Jan 2006",
		"02 January 2006",

		// Standard Go time formats
		// Go 표준 시간 포맷
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,

		// Additional common formats
		// 추가 일반 포맷
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05 -0700",
		"Mon, 02 Jan 2006 15:04:05 -0700",
	}

	// Try parsing with each format
	// 각 포맷으로 파싱 시도
	for _, layout := range formats {
		// Try with default location first
		// 기본 타임존으로 먼저 시도
		t, err := time.ParseInLocation(layout, s, defaultLocation)
		if err == nil {
			return t, nil
		}

		// Try without location
		// 타임존 없이 시도
		t, err = time.Parse(layout, s)
		if err == nil {
			return t.In(defaultLocation), nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time string with any known format: %s", s)
}
