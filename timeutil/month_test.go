package timeutil

import (
	"testing"
	"time"
)

// TestMonthKorean tests MonthKorean function
// MonthKorean 함수 테스트
func TestMonthKorean(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{
			name:     "January",
			time:     time.Date(2025, 1, 1, 0, 0, 0, 0, KST),
			expected: "1월",
		},
		{
			name:     "October",
			time:     time.Date(2025, 10, 14, 0, 0, 0, 0, KST),
			expected: "10월",
		},
		{
			name:     "December",
			time:     time.Date(2025, 12, 31, 0, 0, 0, 0, KST),
			expected: "12월",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MonthKorean(tt.time)
			if result != tt.expected {
				t.Errorf("MonthKorean() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestMonthName tests MonthName function
// MonthName 함수 테스트
func TestMonthName(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{
			name:     "January",
			time:     time.Date(2025, 1, 1, 0, 0, 0, 0, KST),
			expected: "January",
		},
		{
			name:     "October",
			time:     time.Date(2025, 10, 14, 0, 0, 0, 0, KST),
			expected: "October",
		},
		{
			name:     "December",
			time:     time.Date(2025, 12, 31, 0, 0, 0, 0, KST),
			expected: "December",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MonthName(tt.time)
			if result != tt.expected {
				t.Errorf("MonthName() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestMonthNameShort tests MonthNameShort function
// MonthNameShort 함수 테스트
func TestMonthNameShort(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{
			name:     "January",
			time:     time.Date(2025, 1, 1, 0, 0, 0, 0, KST),
			expected: "Jan",
		},
		{
			name:     "October",
			time:     time.Date(2025, 10, 14, 0, 0, 0, 0, KST),
			expected: "Oct",
		},
		{
			name:     "December",
			time:     time.Date(2025, 12, 31, 0, 0, 0, 0, KST),
			expected: "Dec",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MonthNameShort(tt.time)
			if result != tt.expected {
				t.Errorf("MonthNameShort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestQuarter tests Quarter function
// Quarter 함수 테스트
func TestQuarter(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected int
	}{
		{
			name:     "Q1 - January",
			time:     time.Date(2025, 1, 1, 0, 0, 0, 0, KST),
			expected: 1,
		},
		{
			name:     "Q1 - March",
			time:     time.Date(2025, 3, 31, 0, 0, 0, 0, KST),
			expected: 1,
		},
		{
			name:     "Q2 - April",
			time:     time.Date(2025, 4, 1, 0, 0, 0, 0, KST),
			expected: 2,
		},
		{
			name:     "Q2 - June",
			time:     time.Date(2025, 6, 30, 0, 0, 0, 0, KST),
			expected: 2,
		},
		{
			name:     "Q3 - July",
			time:     time.Date(2025, 7, 1, 0, 0, 0, 0, KST),
			expected: 3,
		},
		{
			name:     "Q3 - September",
			time:     time.Date(2025, 9, 30, 0, 0, 0, 0, KST),
			expected: 3,
		},
		{
			name:     "Q4 - October",
			time:     time.Date(2025, 10, 14, 0, 0, 0, 0, KST),
			expected: 4,
		},
		{
			name:     "Q4 - December",
			time:     time.Date(2025, 12, 31, 0, 0, 0, 0, KST),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Quarter(tt.time)
			if result != tt.expected {
				t.Errorf("Quarter() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Benchmark tests for month functions
// month 함수 벤치마크 테스트

// BenchmarkMonthKorean benchmarks MonthKorean function
// MonthKorean 함수 벤치마크
func BenchmarkMonthKorean(b *testing.B) {
	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MonthKorean(t)
	}
}

// BenchmarkMonthName benchmarks MonthName function
// MonthName 함수 벤치마크
func BenchmarkMonthName(b *testing.B) {
	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MonthName(t)
	}
}

// BenchmarkQuarter benchmarks Quarter function
// Quarter 함수 벤치마크
func BenchmarkQuarter(b *testing.B) {
	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Quarter(t)
	}
}
