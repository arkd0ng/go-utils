package timeutil

import (
	"testing"
	"time"
)

// TestParseWithLayout tests parsing with custom layouts.
// TestParseWithLayout은 사용자 지정 레이아웃으로 파싱을 테스트합니다.
func TestParseWithLayout(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		layout  string
		wantErr bool
	}{
		{
			name:    "Milliseconds",
			input:   "2024-10-04 08:34:42.324",
			layout:  "2006-01-02 15:04:05.000",
			wantErr: false,
		},
		{
			name:    "Microseconds",
			input:   "2024-10-04 08:34:42.324567",
			layout:  "2006-01-02 15:04:05.999999",
			wantErr: false,
		},
		{
			name:    "Date with slashes",
			input:   "2024/10/04",
			layout:  "2006/01/02",
			wantErr: false,
		},
		{
			name:    "Invalid format",
			input:   "2024-10-04",
			layout:  "2006/01/02",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseWithLayout(tt.input, tt.layout)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseWithLayout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result.IsZero() {
				t.Errorf("ParseWithLayout() returned zero time")
			}
		})
	}
}

// TestParseMillis tests parsing datetime with milliseconds.
// TestParseMillis는 밀리초를 포함한 날짜시간 파싱을 테스트합니다.
func TestParseMillis(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid milliseconds",
			input:   "2024-10-04 08:34:42.324",
			wantErr: false,
		},
		{
			name:    "Zero milliseconds",
			input:   "2024-10-04 08:34:42.000",
			wantErr: false,
		},
		{
			name:    "Invalid format",
			input:   "2024-10-04 08:34:42",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseMillis(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMillis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result.IsZero() {
				t.Errorf("ParseMillis() returned zero time")
			}
		})
	}
}

// TestParseMicros tests parsing datetime with microseconds.
// TestParseMicros는 마이크로초를 포함한 날짜시간 파싱을 테스트합니다.
func TestParseMicros(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid microseconds",
			input:   "2024-10-04 08:34:42.324567",
			wantErr: false,
		},
		{
			name:    "Zero microseconds",
			input:   "2024-10-04 08:34:42.000000",
			wantErr: false,
		},
		{
			name:    "Invalid format",
			input:   "2024-10-04 08:34:42",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseMicros(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMicros() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result.IsZero() {
				t.Errorf("ParseMicros() returned zero time")
			}
		})
	}
}

// TestParseAny tests parsing with automatic format detection.
// TestParseAny는 자동 포맷 감지 파싱을 테스트합니다.
func TestParseAny(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantErr    bool
		wantYear   int
		wantMonth  time.Month
		wantDay    int
		wantHour   int
		wantMinute int
		wantSecond int
	}{
		{
			name:       "ISO8601",
			input:      "2024-10-04T08:34:42+09:00",
			wantErr:    false,
			wantYear:   2024,
			wantMonth:  time.October,
			wantDay:    4,
			wantHour:   8,
			wantMinute: 34,
			wantSecond: 42,
		},
		{
			name:       "RFC3339",
			input:      "2024-10-04T08:34:42Z",
			wantErr:    false,
			wantYear:   2024,
			wantMonth:  time.October,
			wantDay:    4,
			wantHour:   17, // 08:00 UTC = 17:00 KST
			wantMinute: 34,
			wantSecond: 42,
		},
		{
			name:       "DateTime with milliseconds",
			input:      "2024-10-04 08:34:42.324",
			wantErr:    false,
			wantYear:   2024,
			wantMonth:  time.October,
			wantDay:    4,
			wantHour:   8,
			wantMinute: 34,
			wantSecond: 42,
		},
		{
			name:       "DateTime with microseconds",
			input:      "2024-10-04 08:34:42.324567",
			wantErr:    false,
			wantYear:   2024,
			wantMonth:  time.October,
			wantDay:    4,
			wantHour:   8,
			wantMinute: 34,
			wantSecond: 42,
		},
		{
			name:       "DateTime",
			input:      "2024-10-04 08:34:42",
			wantErr:    false,
			wantYear:   2024,
			wantMonth:  time.October,
			wantDay:    4,
			wantHour:   8,
			wantMinute: 34,
			wantSecond: 42,
		},
		{
			name:      "Date only",
			input:     "2024-10-04",
			wantErr:   false,
			wantYear:  2024,
			wantMonth: time.October,
			wantDay:   4,
		},
		{
			name:      "Date with slashes",
			input:     "2024/10/04",
			wantErr:   false,
			wantYear:  2024,
			wantMonth: time.October,
			wantDay:   4,
		},
		{
			name:      "US format",
			input:     "10/04/2024",
			wantErr:   false,
			wantYear:  2024,
			wantMonth: time.October,
			wantDay:   4,
		},
		{
			name:      "Month name",
			input:     "Oct 04, 2024",
			wantErr:   false,
			wantYear:  2024,
			wantMonth: time.October,
			wantDay:   4,
		},
		{
			name:      "Full month name",
			input:     "October 04, 2024",
			wantErr:   false,
			wantYear:  2024,
			wantMonth: time.October,
			wantDay:   4,
		},
		{
			name:      "Short month format",
			input:     "04-Oct-2024",
			wantErr:   false,
			wantYear:  2024,
			wantMonth: time.October,
			wantDay:   4,
		},
		{
			name:    "Empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:    "Invalid format",
			input:   "not a date",
			wantErr: true,
		},
		{
			name:    "Invalid date",
			input:   "2024-13-45",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseAny(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAny() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if result.Year() != tt.wantYear {
					t.Errorf("ParseAny() year = %d, want %d", result.Year(), tt.wantYear)
				}
				if result.Month() != tt.wantMonth {
					t.Errorf("ParseAny() month = %v, want %v", result.Month(), tt.wantMonth)
				}
				if result.Day() != tt.wantDay {
					t.Errorf("ParseAny() day = %d, want %d", result.Day(), tt.wantDay)
				}
				if tt.wantHour != 0 && result.Hour() != tt.wantHour {
					t.Errorf("ParseAny() hour = %d, want %d", result.Hour(), tt.wantHour)
				}
				if tt.wantMinute != 0 && result.Minute() != tt.wantMinute {
					t.Errorf("ParseAny() minute = %d, want %d", result.Minute(), tt.wantMinute)
				}
				if tt.wantSecond != 0 && result.Second() != tt.wantSecond {
					t.Errorf("ParseAny() second = %d, want %d", result.Second(), tt.wantSecond)
				}
			}
		})
	}
}

// TestParseAnyDatabaseFormats tests common database timestamp formats.
// TestParseAnyDatabaseFormats는 일반적인 데이터베이스 타임스탬프 포맷을 테스트합니다.
func TestParseAnyDatabaseFormats(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		desc    string
	}{
		{
			name:    "MySQL DATETIME",
			input:   "2024-10-04 08:34:42",
			wantErr: false,
			desc:    "Standard MySQL DATETIME format",
		},
		{
			name:    "MySQL DATETIME with milliseconds",
			input:   "2024-10-04 08:34:42.324",
			wantErr: false,
			desc:    "MySQL DATETIME(3) with milliseconds",
		},
		{
			name:    "MySQL DATETIME with microseconds",
			input:   "2024-10-04 08:34:42.324567",
			wantErr: false,
			desc:    "MySQL DATETIME(6) with microseconds",
		},
		{
			name:    "PostgreSQL TIMESTAMP",
			input:   "2024-10-04 08:34:42.324567",
			wantErr: false,
			desc:    "PostgreSQL TIMESTAMP format",
		},
		{
			name:    "SQLite DATETIME",
			input:   "2024-10-04 08:34:42",
			wantErr: false,
			desc:    "SQLite DATETIME format",
		},
		{
			name:    "ISO8601 with timezone",
			input:   "2024-10-04T08:34:42+09:00",
			wantErr: false,
			desc:    "ISO8601 format with timezone (common in APIs)",
		},
		{
			name:    "RFC3339",
			input:   "2024-10-04T08:34:42Z",
			wantErr: false,
			desc:    "RFC3339 format (common in JSON APIs)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseAny(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAny() for %s error = %v, wantErr %v", tt.desc, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result.IsZero() {
				t.Errorf("ParseAny() for %s returned zero time", tt.desc)
			}
		})
	}
}

// BenchmarkParseWithLayout benchmarks ParseWithLayout function.
// BenchmarkParseWithLayout은 ParseWithLayout 함수를 벤치마크합니다.
func BenchmarkParseWithLayout(b *testing.B) {
	input := "2024-10-04 08:34:42.324"
	layout := "2006-01-02 15:04:05.000"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseWithLayout(input, layout)
	}
}

// BenchmarkParseMillis benchmarks ParseMillis function.
// BenchmarkParseMillis는 ParseMillis 함수를 벤치마크합니다.
func BenchmarkParseMillis(b *testing.B) {
	input := "2024-10-04 08:34:42.324"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseMillis(input)
	}
}

// BenchmarkParseMicros benchmarks ParseMicros function.
// BenchmarkParseMicros는 ParseMicros 함수를 벤치마크합니다.
func BenchmarkParseMicros(b *testing.B) {
	input := "2024-10-04 08:34:42.324567"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseMicros(input)
	}
}

// BenchmarkParseAny benchmarks ParseAny function.
// BenchmarkParseAny는 ParseAny 함수를 벤치마크합니다.
func BenchmarkParseAny(b *testing.B) {
	input := "2024-10-04 08:34:42.324"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseAny(input)
	}
}

// BenchmarkParseAnyISO8601 benchmarks ParseAny with ISO8601 format.
// BenchmarkParseAnyISO8601은 ISO8601 포맷으로 ParseAny를 벤치마크합니다.
func BenchmarkParseAnyISO8601(b *testing.B) {
	input := "2024-10-04T08:34:42+09:00"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseAny(input)
	}
}

// BenchmarkParseAnyDate benchmarks ParseAny with date-only format.
// BenchmarkParseAnyDate는 날짜만 있는 포맷으로 ParseAny를 벤치마크합니다.
func BenchmarkParseAnyDate(b *testing.B) {
	input := "2024-10-04"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseAny(input)
	}
}

// TestParseAnyKoreanFormats tests Korean date/time formats.
// TestParseAnyKoreanFormats는 한글 날짜/시간 포맷을 테스트합니다.
func TestParseAnyKoreanFormats(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantYear int
	}{
		{
			name:     "Korean DateTime Full",
			input:    "2024년 10월 04일 15시 30분 45초",
			wantErr:  false,
			wantYear: 2024,
		},
		{
			name:     "Korean DateTime without seconds",
			input:    "2024년 10월 04일 15시 30분",
			wantErr:  false,
			wantYear: 2024,
		},
		{
			name:     "Korean DateTime hour only",
			input:    "2024년 10월 04일 15시",
			wantErr:  false,
			wantYear: 2024,
		},
		{
			name:     "Korean Date only",
			input:    "2024년 10월 04일",
			wantErr:  false,
			wantYear: 2024,
		},
		{
			name:     "Korean without leading zeros",
			input:    "2024년 1월 4일 9시 5분 3초",
			wantErr:  false,
			wantYear: 2024,
		},
		{
			name:     "Korean with AM (오전)",
			input:    "2024년 10월 04일 오전 9시 30분 45초",
			wantErr:  false,
			wantYear: 2024,
		},
		{
			name:     "Korean with PM (오후)",
			input:    "2024년 10월 04일 오후 3시 30분 45초",
			wantErr:  false,
			wantYear: 2024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseAny(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAny() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result.Year() != tt.wantYear {
					t.Errorf("ParseAny() year = %d, want %d", result.Year(), tt.wantYear)
				}
			}
		})
	}
}
