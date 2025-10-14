package timeutil

import (
	"testing"
	"time"
)

// TestWeekOfYear tests WeekOfYear function / WeekOfYear 함수 테스트
func TestWeekOfYear(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected int
	}{
		{
			name:     "Week 1 of 2025",
			time:     time.Date(2025, 1, 1, 0, 0, 0, 0, KST),
			expected: 1,
		},
		{
			name:     "Week 42 of 2025",
			time:     time.Date(2025, 10, 14, 0, 0, 0, 0, KST),
			expected: 42,
		},
		{
			name:     "Last week of 2025",
			time:     time.Date(2025, 12, 31, 0, 0, 0, 0, KST),
			expected: 1, // Week 1 of 2026 in ISO 8601
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WeekOfYear(tt.time)
			if result != tt.expected {
				t.Errorf("WeekOfYear() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestWeekOfMonth tests WeekOfMonth function / WeekOfMonth 함수 테스트
func TestWeekOfMonth(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected int
	}{
		{
			name:     "First day of month (Wednesday)",
			time:     time.Date(2025, 10, 1, 0, 0, 0, 0, KST),
			expected: 1,
		},
		{
			name:     "Day 6 (Monday, start of week 2)",
			time:     time.Date(2025, 10, 6, 0, 0, 0, 0, KST),
			expected: 2,
		},
		{
			name:     "Day 14 (Tuesday, in week 3)",
			time:     time.Date(2025, 10, 14, 0, 0, 0, 0, KST),
			expected: 3,
		},
		{
			name:     "Day 31 of month (Friday, in week 5)",
			time:     time.Date(2025, 10, 31, 0, 0, 0, 0, KST),
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WeekOfMonth(tt.time)
			if result != tt.expected {
				t.Errorf("WeekOfMonth() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestDaysInMonth tests DaysInMonth function / DaysInMonth 함수 테스트
func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected int
	}{
		{
			name:     "January (31 days)",
			time:     time.Date(2025, 1, 1, 0, 0, 0, 0, KST),
			expected: 31,
		},
		{
			name:     "February 2025 (28 days)",
			time:     time.Date(2025, 2, 1, 0, 0, 0, 0, KST),
			expected: 28,
		},
		{
			name:     "February 2024 (29 days, leap year)",
			time:     time.Date(2024, 2, 1, 0, 0, 0, 0, KST),
			expected: 29,
		},
		{
			name:     "April (30 days)",
			time:     time.Date(2025, 4, 1, 0, 0, 0, 0, KST),
			expected: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DaysInMonth(tt.time)
			if result != tt.expected {
				t.Errorf("DaysInMonth() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestDaysInYear tests DaysInYear function / DaysInYear 함수 테스트
func TestDaysInYear(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected int
	}{
		{
			name:     "2025 (365 days)",
			time:     time.Date(2025, 1, 1, 0, 0, 0, 0, KST),
			expected: 365,
		},
		{
			name:     "2024 (366 days, leap year)",
			time:     time.Date(2024, 1, 1, 0, 0, 0, 0, KST),
			expected: 366,
		},
		{
			name:     "2000 (366 days, leap year)",
			time:     time.Date(2000, 1, 1, 0, 0, 0, 0, KST),
			expected: 366,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DaysInYear(tt.time)
			if result != tt.expected {
				t.Errorf("DaysInYear() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Benchmark tests for week functions / week 함수 벤치마크 테스트

// BenchmarkWeekOfYear benchmarks WeekOfYear function / WeekOfYear 함수 벤치마크
func BenchmarkWeekOfYear(b *testing.B) {
	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WeekOfYear(t)
	}
}

// BenchmarkWeekOfMonth benchmarks WeekOfMonth function / WeekOfMonth 함수 벤치마크
func BenchmarkWeekOfMonth(b *testing.B) {
	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WeekOfMonth(t)
	}
}

// BenchmarkDaysInMonth benchmarks DaysInMonth function / DaysInMonth 함수 벤치마크
func BenchmarkDaysInMonth(b *testing.B) {
	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DaysInMonth(t)
	}
}
