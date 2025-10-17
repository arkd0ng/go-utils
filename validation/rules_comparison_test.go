package validation

import (
	"testing"
	"time"
)

func TestEquals(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		expected  interface{}
		wantError bool
	}{
		{"equal strings", "hello", "hello", false},
		{"not equal strings", "hello", "world", true},
		{"equal ints", 42, 42, false},
		{"not equal ints", 42, 10, true},
		{"equal floats", 3.14, 3.14, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Equals(tt.expected)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("Equals() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestNotEquals(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		forbidden interface{}
		wantError bool
	}{
		{"not equal strings", "hello", "world", false},
		{"equal strings", "hello", "hello", true},
		{"not equal ints", 42, 10, false},
		{"equal ints", 42, 42, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.NotEquals(tt.forbidden)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("NotEquals() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestGreaterThan(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		threshold float64
		wantError bool
	}{
		{"int greater", 10, 5.0, false},
		{"int equal", 10, 10.0, true},
		{"int less", 5, 10.0, true},
		{"float greater", 10.5, 10.0, false},
		{"float equal", 10.0, 10.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.GreaterThan(tt.threshold)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("GreaterThan() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		threshold float64
		wantError bool
	}{
		{"int greater", 10, 5.0, false},
		{"int equal", 10, 10.0, false},
		{"int less", 5, 10.0, true},
		{"float greater", 10.5, 10.0, false},
		{"float equal", 10.0, 10.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.GreaterThanOrEqual(tt.threshold)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("GreaterThanOrEqual() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestLessThan(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		threshold float64
		wantError bool
	}{
		{"int less", 5, 10.0, false},
		{"int equal", 10, 10.0, true},
		{"int greater", 15, 10.0, true},
		{"float less", 9.5, 10.0, false},
		{"float equal", 10.0, 10.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.LessThan(tt.threshold)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("LessThan() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestLessThanOrEqual(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		threshold float64
		wantError bool
	}{
		{"int less", 5, 10.0, false},
		{"int equal", 10, 10.0, false},
		{"int greater", 15, 10.0, true},
		{"float less", 9.5, 10.0, false},
		{"float equal", 10.0, 10.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.LessThanOrEqual(tt.threshold)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("LessThanOrEqual() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestBefore(t *testing.T) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)

	tests := []struct {
		name      string
		value     interface{}
		threshold time.Time
		wantError bool
	}{
		{"before", yesterday, now, false},
		{"after", tomorrow, now, true},
		{"equal", now, now, true},
		{"not a time", "string", now, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Before(tt.threshold)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("Before() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestAfter(t *testing.T) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)

	tests := []struct {
		name      string
		value     interface{}
		threshold time.Time
		wantError bool
	}{
		{"after", tomorrow, now, false},
		{"before", yesterday, now, true},
		{"equal", now, now, true},
		{"not a time", "string", now, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.After(tt.threshold)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("After() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestBeforeOrEqual(t *testing.T) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)

	tests := []struct {
		name      string
		value     interface{}
		threshold time.Time
		wantError bool
	}{
		{"before", yesterday, now, false},
		{"equal", now, now, false},
		{"after", tomorrow, now, true},
		{"not a time", "string", now, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.BeforeOrEqual(tt.threshold)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("BeforeOrEqual() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestAfterOrEqual(t *testing.T) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)

	tests := []struct {
		name      string
		value     interface{}
		threshold time.Time
		wantError bool
	}{
		{"after", tomorrow, now, false},
		{"equal", now, now, false},
		{"before", yesterday, now, true},
		{"not a time", "string", now, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.AfterOrEqual(tt.threshold)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("AfterOrEqual() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestComparisonChaining(t *testing.T) {
	t.Run("valid numeric comparison chaining", func(t *testing.T) {
		v := New(50, "score")
		v.GreaterThan(0).LessThan(100).NotEquals(75)
		err := v.Validate()
		if err != nil {
			t.Errorf("Valid comparison chaining failed: %v", err)
		}
	})

	t.Run("invalid numeric comparison chaining", func(t *testing.T) {
		v := New(150, "score")
		v.GreaterThan(0).LessThan(100)
		err := v.Validate()
		if err == nil {
			t.Error("Expected error for value exceeding threshold")
		}
	})

	t.Run("valid time comparison chaining", func(t *testing.T) {
		now := time.Now()
		yesterday := now.Add(-24 * time.Hour)
		tomorrow := now.Add(24 * time.Hour)

		v := New(now, "date")
		v.After(yesterday).Before(tomorrow)
		err := v.Validate()
		if err != nil {
			t.Errorf("Valid time comparison chaining failed: %v", err)
		}
	})
}

// TestBetweenTime tests the BetweenTime validator
// TestBetweenTime는 BetweenTime 검증기를 테스트합니다
func TestBetweenTime(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	middle := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	before := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
	after := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name      string
		value     interface{}
		start     time.Time
		end       time.Time
		shouldErr bool
	}{
		{"valid time in range", middle, start, end, false},
		{"valid at start boundary", start, start, end, false},
		{"valid at end boundary", end, start, end, false},
		{"invalid before start", before, start, end, true},
		{"invalid after end", after, start, end, true},
		{"invalid non-time value", "2024-06-15", start, end, true},
		{"invalid nil value", nil, start, end, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "date")
			v.BetweenTime(tt.start, tt.end)
			err := v.Validate()

			if tt.shouldErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

// TestBetweenTimeStopOnError tests BetweenTime with StopOnError
// TestBetweenTimeStopOnError는 StopOnError가 있는 BetweenTime을 테스트합니다
func TestBetweenTimeStopOnError(t *testing.T) {
	t.Run("BetweenTime StopOnError", func(t *testing.T) {
		start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
		after := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

		v := New(after, "date")
		v.StopOnError().BetweenTime(start, end).After(end) // After should not run

		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("Expected 1 error (stopped after first), got %d", len(errors))
		}
	})
}

// TestBetweenTimeChaining tests BetweenTime with chaining
// TestBetweenTimeChaining는 체이닝이 있는 BetweenTime을 테스트합니다
func TestBetweenTimeChaining(t *testing.T) {
	t.Run("BetweenTime with chaining", func(t *testing.T) {
		start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
		middle := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)

		v := New(middle, "date")
		v.BetweenTime(start, end).After(start).Before(end)
		err := v.Validate()
		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}
	})
}

// TestBetweenTimeEdgeCases tests BetweenTime edge cases
// TestBetweenTimeEdgeCases는 BetweenTime 경계 사례를 테스트합니다
func TestBetweenTimeEdgeCases(t *testing.T) {
	t.Run("BetweenTime with same start and end", func(t *testing.T) {
		now := time.Now()
		v := New(now, "date")
		v.BetweenTime(now, now)
		err := v.Validate()
		if err != nil {
			t.Errorf("Expected no error for time equal to start and end, got: %v", err)
		}
	})

	t.Run("BetweenTime with inverted range", func(t *testing.T) {
		start := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
		end := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		middle := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)

		v := New(middle, "date")
		v.BetweenTime(start, end) // Inverted range
		err := v.Validate()
		if err == nil {
			t.Errorf("Expected error for inverted time range")
		}
	})
}
