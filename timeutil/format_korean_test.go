package timeutil

import (
	"testing"
	"time"
)

// TestWeekdayKorean tests WeekdayKorean function
// WeekdayKorean 함수 테스트
func TestWeekdayKorean(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{
			name:     "Sunday",
			time:     time.Date(2025, 10, 12, 0, 0, 0, 0, KST), // Sunday
			expected: "일요일",
		},
		{
			name:     "Monday",
			time:     time.Date(2025, 10, 13, 0, 0, 0, 0, KST), // Monday
			expected: "월요일",
		},
		{
			name:     "Tuesday",
			time:     time.Date(2025, 10, 14, 0, 0, 0, 0, KST), // Tuesday
			expected: "화요일",
		},
		{
			name:     "Wednesday",
			time:     time.Date(2025, 10, 15, 0, 0, 0, 0, KST), // Wednesday
			expected: "수요일",
		},
		{
			name:     "Thursday",
			time:     time.Date(2025, 10, 16, 0, 0, 0, 0, KST), // Thursday
			expected: "목요일",
		},
		{
			name:     "Friday",
			time:     time.Date(2025, 10, 17, 0, 0, 0, 0, KST), // Friday
			expected: "금요일",
		},
		{
			name:     "Saturday",
			time:     time.Date(2025, 10, 18, 0, 0, 0, 0, KST), // Saturday
			expected: "토요일",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WeekdayKorean(tt.time)
			if result != tt.expected {
				t.Errorf("WeekdayKorean() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestWeekdayKoreanShort tests WeekdayKoreanShort function
// WeekdayKoreanShort 함수 테스트
func TestWeekdayKoreanShort(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{
			name:     "Sunday",
			time:     time.Date(2025, 10, 12, 0, 0, 0, 0, KST), // Sunday
			expected: "일",
		},
		{
			name:     "Monday",
			time:     time.Date(2025, 10, 13, 0, 0, 0, 0, KST), // Monday
			expected: "월",
		},
		{
			name:     "Tuesday",
			time:     time.Date(2025, 10, 14, 0, 0, 0, 0, KST), // Tuesday
			expected: "화",
		},
		{
			name:     "Wednesday",
			time:     time.Date(2025, 10, 15, 0, 0, 0, 0, KST), // Wednesday
			expected: "수",
		},
		{
			name:     "Thursday",
			time:     time.Date(2025, 10, 16, 0, 0, 0, 0, KST), // Thursday
			expected: "목",
		},
		{
			name:     "Friday",
			time:     time.Date(2025, 10, 17, 0, 0, 0, 0, KST), // Friday
			expected: "금",
		},
		{
			name:     "Saturday",
			time:     time.Date(2025, 10, 18, 0, 0, 0, 0, KST), // Saturday
			expected: "토",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WeekdayKoreanShort(tt.time)
			if result != tt.expected {
				t.Errorf("WeekdayKoreanShort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestFormatKoreanDateTime tests FormatKoreanDateTime function
// FormatKoreanDateTime 함수 테스트
func TestFormatKoreanDateTime(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 30, 45, 0, KST) // Tuesday

	result := FormatKoreanDateTime(testTime)
	expected := "2025년 10월 14일 (화요일) 15시 30분 45초"

	if result != expected {
		t.Errorf("FormatKoreanDateTime() = %v, want %v", result, expected)
	}
}

// TestFormatKoreanDateWithWeekday tests FormatKoreanDateWithWeekday function
// FormatKoreanDateWithWeekday 함수 테스트
func TestFormatKoreanDateWithWeekday(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 30, 45, 0, KST) // Tuesday

	result := FormatKoreanDateWithWeekday(testTime)
	expected := "2025년 10월 14일 (화요일)"

	if result != expected {
		t.Errorf("FormatKoreanDateWithWeekday() = %v, want %v", result, expected)
	}
}

// TestFormatKoreanDateShort tests FormatKoreanDateShort function
// FormatKoreanDateShort 함수 테스트
func TestFormatKoreanDateShort(t *testing.T) {
	testTime := time.Date(2025, 10, 14, 15, 30, 45, 0, KST) // Tuesday

	result := FormatKoreanDateShort(testTime)
	expected := "2025년 10월 14일 (화)"

	if result != expected {
		t.Errorf("FormatKoreanDateShort() = %v, want %v", result, expected)
	}
}

// BenchmarkWeekdayKorean benchmarks WeekdayKorean function
// WeekdayKorean 함수 벤치마크
func BenchmarkWeekdayKorean(b *testing.B) {
	testTime := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WeekdayKorean(testTime)
	}
}

// BenchmarkWeekdayKoreanShort benchmarks WeekdayKoreanShort function
// WeekdayKoreanShort 함수 벤치마크
func BenchmarkWeekdayKoreanShort(b *testing.B) {
	testTime := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WeekdayKoreanShort(testTime)
	}
}

// BenchmarkFormatKoreanDateTime benchmarks FormatKoreanDateTime function
// FormatKoreanDateTime 함수 벤치마크
func BenchmarkFormatKoreanDateTime(b *testing.B) {
	testTime := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatKoreanDateTime(testTime)
	}
}
