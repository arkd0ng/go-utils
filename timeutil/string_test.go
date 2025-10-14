package timeutil

import (
	"testing"
)

// TestSubTimeString tests SubTimeString function.
// TestSubTimeString은 SubTimeString 함수를 테스트합니다.
func TestSubTimeString(t *testing.T) {
	tests := []struct {
		name       string
		s1         string
		s2         string
		wantErr    bool
		wantDays   float64
		wantHours  float64
		tolerance  float64 // for floating point comparison
	}{
		{
			name:      "Same day",
			s1:        "2024-10-04 08:00:00",
			s2:        "2024-10-04 14:30:00",
			wantErr:   false,
			wantHours: 6.5,
			tolerance: 0.1,
		},
		{
			name:      "Multiple days with milliseconds",
			s1:        "2024-10-04 08:34:42.324",
			s2:        "2024-10-14 14:56:23.789",
			wantErr:   false,
			wantDays:  10.27,
			tolerance: 0.01,
		},
		{
			name:    "Invalid first time",
			s1:      "invalid",
			s2:      "2024-10-14 14:56:23",
			wantErr: true,
		},
		{
			name:    "Invalid second time",
			s1:      "2024-10-04 08:34:42",
			s2:      "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SubTimeString(tt.s1, tt.s2)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubTimeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if tt.wantDays != 0 {
					days := result.Days()
					if days < tt.wantDays-tt.tolerance || days > tt.wantDays+tt.tolerance {
						t.Errorf("SubTimeString() days = %v, want %v (±%v)", days, tt.wantDays, tt.tolerance)
					}
				}
				if tt.wantHours != 0 {
					hours := result.Hours()
					if hours < tt.wantHours-tt.tolerance || hours > tt.wantHours+tt.tolerance {
						t.Errorf("SubTimeString() hours = %v, want %v (±%v)", hours, tt.wantHours, tt.tolerance)
					}
				}
			}
		})
	}
}

// TestDiffInDaysString tests DiffInDaysString function.
// TestDiffInDaysString은 DiffInDaysString 함수를 테스트합니다.
func TestDiffInDaysString(t *testing.T) {
	tests := []struct {
		name      string
		s1        string
		s2        string
		want      float64
		tolerance float64
		wantErr   bool
	}{
		{
			name:      "Same day",
			s1:        "2024-10-04",
			s2:        "2024-10-04",
			want:      0,
			tolerance: 0.01,
			wantErr:   false,
		},
		{
			name:      "One week",
			s1:        "2024-10-04",
			s2:        "2024-10-11",
			want:      7,
			tolerance: 0.01,
			wantErr:   false,
		},
		{
			name:      "Different formats",
			s1:        "2024-10-04",
			s2:        "Oct 11, 2024",
			want:      7,
			tolerance: 0.01,
			wantErr:   false,
		},
		{
			name:    "Invalid first date",
			s1:      "invalid",
			s2:      "2024-10-11",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DiffInDaysString(tt.s1, tt.s2)
			if (err != nil) != tt.wantErr {
				t.Errorf("DiffInDaysString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result < tt.want-tt.tolerance || result > tt.want+tt.tolerance {
					t.Errorf("DiffInDaysString() = %v, want %v (±%v)", result, tt.want, tt.tolerance)
				}
			}
		})
	}
}

// TestAgeString tests AgeString function.
// TestAgeString은 AgeString 함수를 테스트합니다.
func TestAgeString(t *testing.T) {
	tests := []struct {
		name      string
		birthDate string
		wantErr   bool
		minAge    int // minimum expected age
		maxAge    int // maximum expected age
	}{
		{
			name:      "Valid birth date",
			birthDate: "1990-01-15",
			wantErr:   false,
			minAge:    34,
			maxAge:    36,
		},
		{
			name:      "Different format",
			birthDate: "Jan 15, 1990",
			wantErr:   false,
			minAge:    34,
			maxAge:    36,
		},
		{
			name:      "Invalid date",
			birthDate: "invalid",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AgeString(tt.birthDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result.Years < tt.minAge || result.Years > tt.maxAge {
					t.Errorf("AgeString() = %v, want between %v and %v", result.Years, tt.minAge, tt.maxAge)
				}
			}
		})
	}
}

// TestAgeInYearsString tests AgeInYearsString function.
// TestAgeInYearsString은 AgeInYearsString 함수를 테스트합니다.
func TestAgeInYearsString(t *testing.T) {
	tests := []struct {
		name      string
		birthDate string
		wantErr   bool
		minAge    int // minimum expected age
		maxAge    int // maximum expected age
	}{
		{
			name:      "Valid birth date",
			birthDate: "1990-01-15",
			wantErr:   false,
			minAge:    34,
			maxAge:    36,
		},
		{
			name:      "Different format",
			birthDate: "Jan 15, 1990",
			wantErr:   false,
			minAge:    34,
			maxAge:    36,
		},
		{
			name:      "Invalid date",
			birthDate: "invalid",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AgeInYearsString(tt.birthDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgeInYearsString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result < tt.minAge || result > tt.maxAge {
					t.Errorf("AgeInYearsString() = %v, want between %v and %v", result, tt.minAge, tt.maxAge)
				}
			}
		})
	}
}

// TestRelativeTimeString tests RelativeTimeString function.
// TestRelativeTimeString은 RelativeTimeString 함수를 테스트합니다.
func TestRelativeTimeString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid date",
			input:   "2024-10-13 15:30:00",
			wantErr: false,
		},
		{
			name:    "Different format",
			input:   "Oct 13, 2024",
			wantErr: false,
		},
		{
			name:    "Invalid date",
			input:   "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RelativeTimeString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RelativeTimeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == "" {
				t.Errorf("RelativeTimeString() returned empty string")
			}
		})
	}
}

// TestIsBusinessDayString tests IsBusinessDayString function.
// TestIsBusinessDayString은 IsBusinessDayString 함수를 테스트합니다.
func TestIsBusinessDayString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		wantErr bool
	}{
		{
			name:    "Monday (business day)",
			input:   "2024-10-14", // Monday
			want:    true,
			wantErr: false,
		},
		{
			name:    "Saturday (weekend)",
			input:   "2024-10-12", // Saturday
			want:    false,
			wantErr: false,
		},
		{
			name:    "Different format",
			input:   "Oct 14, 2024",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Invalid date",
			input:   "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := IsBusinessDayString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsBusinessDayString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.want {
				t.Errorf("IsBusinessDayString() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestAddDaysString tests AddDaysString function.
// TestAddDaysString은 AddDaysString 함수를 테스트합니다.
func TestAddDaysString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		days    int
		wantDay int
		wantErr bool
	}{
		{
			name:    "Add 7 days",
			input:   "2024-10-04",
			days:    7,
			wantDay: 11,
			wantErr: false,
		},
		{
			name:    "Subtract 7 days (negative)",
			input:   "2024-10-14",
			days:    -7,
			wantDay: 7,
			wantErr: false,
		},
		{
			name:    "Different format",
			input:   "Oct 04, 2024",
			days:    7,
			wantDay: 11,
			wantErr: false,
		},
		{
			name:    "Invalid date",
			input:   "invalid",
			days:    7,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AddDaysString(tt.input, tt.days)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddDaysString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result.Day() != tt.wantDay {
				t.Errorf("AddDaysString() day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

// TestFormatString tests FormatString function.
// TestFormatString은 FormatString 함수를 테스트합니다.
func TestFormatString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		layout  string
		want    string
		wantErr bool
	}{
		{
			name:    "Format as date",
			input:   "2024-10-04 15:30:00",
			layout:  "2006-01-02",
			want:    "2024-10-04",
			wantErr: false,
		},
		{
			name:    "Format with different input format",
			input:   "Oct 04, 2024",
			layout:  "2006-01-02",
			want:    "2024-10-04",
			wantErr: false,
		},
		{
			name:    "Invalid input",
			input:   "invalid",
			layout:  "2006-01-02",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FormatString(tt.input, tt.layout)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.want {
				t.Errorf("FormatString() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestFormatDateString tests FormatDateString function.
// TestFormatDateString은 FormatDateString 함수를 테스트합니다.
func TestFormatDateString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "DateTime to Date",
			input:   "2024-10-04 15:30:00",
			want:    "2024-10-04",
			wantErr: false,
		},
		{
			name:    "Different format",
			input:   "Oct 04, 2024",
			want:    "2024-10-04",
			wantErr: false,
		},
		{
			name:    "Invalid input",
			input:   "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FormatDateString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatDateString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.want {
				t.Errorf("FormatDateString() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestWeekdayKoreanString tests WeekdayKoreanString function.
// TestWeekdayKoreanString은 WeekdayKoreanString 함수를 테스트합니다.
func TestWeekdayKoreanString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "Monday",
			input:   "2024-10-14", // Monday
			want:    "월요일",
			wantErr: false,
		},
		{
			name:    "Different format",
			input:   "Oct 14, 2024",
			want:    "월요일",
			wantErr: false,
		},
		{
			name:    "Invalid date",
			input:   "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := WeekdayKoreanString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeekdayKoreanString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.want {
				t.Errorf("WeekdayKoreanString() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestStartOfDayString tests StartOfDayString function.
// TestStartOfDayString은 StartOfDayString 함수를 테스트합니다.
func TestStartOfDayString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid date",
			input:   "2024-10-04 15:30:45",
			wantErr: false,
		},
		{
			name:    "Different format",
			input:   "Oct 04, 2024",
			wantErr: false,
		},
		{
			name:    "Invalid date",
			input:   "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StartOfDayString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartOfDayString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
					t.Errorf("StartOfDayString() = %v, want 00:00:00", result.Format("15:04:05"))
				}
			}
		})
	}
}

// TestIsSameDayString tests IsSameDayString function.
// TestIsSameDayString은 IsSameDayString 함수를 테스트합니다.
func TestIsSameDayString(t *testing.T) {
	tests := []struct {
		name    string
		s1      string
		s2      string
		want    bool
		wantErr bool
	}{
		{
			name:    "Same day, different times",
			s1:      "2024-10-04 08:00:00",
			s2:      "2024-10-04 15:30:00",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Different days",
			s1:      "2024-10-04",
			s2:      "2024-10-05",
			want:    false,
			wantErr: false,
		},
		{
			name:    "Different formats, same day",
			s1:      "2024-10-04",
			s2:      "Oct 04, 2024",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Invalid first date",
			s1:      "invalid",
			s2:      "2024-10-04",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := IsSameDayString(tt.s1, tt.s2)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSameDayString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.want {
				t.Errorf("IsSameDayString() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestIsBeforeString tests IsBeforeString function.
// TestIsBeforeString은 IsBeforeString 함수를 테스트합니다.
func TestIsBeforeString(t *testing.T) {
	tests := []struct {
		name    string
		s1      string
		s2      string
		want    bool
		wantErr bool
	}{
		{
			name:    "s1 before s2",
			s1:      "2024-10-04",
			s2:      "2024-10-05",
			want:    true,
			wantErr: false,
		},
		{
			name:    "s1 after s2",
			s1:      "2024-10-05",
			s2:      "2024-10-04",
			want:    false,
			wantErr: false,
		},
		{
			name:    "Different formats",
			s1:      "2024-10-04",
			s2:      "Oct 05, 2024",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Invalid first date",
			s1:      "invalid",
			s2:      "2024-10-04",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := IsBeforeString(tt.s1, tt.s2)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsBeforeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.want {
				t.Errorf("IsBeforeString() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestIsBetweenString tests IsBetweenString function.
// TestIsBetweenString은 IsBetweenString 함수를 테스트합니다.
func TestIsBetweenString(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		start   string
		end     string
		want    bool
		wantErr bool
	}{
		{
			name:    "Between dates",
			s:       "2024-10-10",
			start:   "2024-10-04",
			end:     "2024-10-14",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Before start",
			s:       "2024-10-03",
			start:   "2024-10-04",
			end:     "2024-10-14",
			want:    false,
			wantErr: false,
		},
		{
			name:    "After end",
			s:       "2024-10-15",
			start:   "2024-10-04",
			end:     "2024-10-14",
			want:    false,
			wantErr: false,
		},
		{
			name:    "Different formats",
			s:       "Oct 10, 2024",
			start:   "2024-10-04",
			end:     "2024-10-14",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Invalid date",
			s:       "invalid",
			start:   "2024-10-04",
			end:     "2024-10-14",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := IsBetweenString(tt.s, tt.start, tt.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsBetweenString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.want {
				t.Errorf("IsBetweenString() = %v, want %v", result, tt.want)
			}
		})
	}
}

// BenchmarkSubTimeString benchmarks SubTimeString function.
// BenchmarkSubTimeString은 SubTimeString 함수를 벤치마크합니다.
func BenchmarkSubTimeString(b *testing.B) {
	s1 := "2024-10-04 08:34:42.324"
	s2 := "2024-10-14 14:56:23.789"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SubTimeString(s1, s2)
	}
}

// BenchmarkDiffInDaysString benchmarks DiffInDaysString function.
// BenchmarkDiffInDaysString은 DiffInDaysString 함수를 벤치마크합니다.
func BenchmarkDiffInDaysString(b *testing.B) {
	s1 := "2024-10-04"
	s2 := "2024-10-14"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DiffInDaysString(s1, s2)
	}
}

// BenchmarkFormatDateString benchmarks FormatDateString function.
// BenchmarkFormatDateString은 FormatDateString 함수를 벤치마크합니다.
func BenchmarkFormatDateString(b *testing.B) {
	input := "2024-10-04 15:30:00"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FormatDateString(input)
	}
}

// BenchmarkIsBusinessDayString benchmarks IsBusinessDayString function.
// BenchmarkIsBusinessDayString은 IsBusinessDayString 함수를 벤치마크합니다.
func BenchmarkIsBusinessDayString(b *testing.B) {
	input := "2024-10-14"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = IsBusinessDayString(input)
	}
}
