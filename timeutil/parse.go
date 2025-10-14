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
// Supported formats / 지원되는 포맷:
//   - ISO8601: 2006-01-02T15:04:05Z07:00
//   - RFC3339: 2006-01-02T15:04:05Z07:00
//   - Date: 2006-01-02
//   - DateTime: 2006-01-02 15:04:05
func Parse(s string) (time.Time, error) {
	s = strings.TrimSpace(s)

	// Try ISO8601/RFC3339 / ISO8601/RFC3339 시도
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

	// Try DateTime / DateTime 시도
	if strings.Contains(s, " ") && strings.Contains(s, ":") {
		t, err := time.ParseInLocation(DateTimeLayout, s, defaultLocation)
		if err == nil {
			return t, nil
		}
	}

	// Try Date / Date 시도
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

	// Try different formats / 다른 포맷 시도
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
