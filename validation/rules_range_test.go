package validation

import (
	"testing"
	"time"
)

// TestIntRange tests the IntRange validator
func TestIntRange(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		min       int
		max       int
		wantError bool
	}{
		// Valid int values
		{"valid int at min", 10, 10, 20, false},
		{"valid int at max", 20, 10, 20, false},
		{"valid int in middle", 15, 10, 20, false},
		{"valid int zero range", 0, 0, 0, false},
		{"valid int negative range", -5, -10, 0, false},

		// Valid different int types
		{"valid int8", int8(15), 10, 20, false},
		{"valid int16", int16(15), 10, 20, false},
		{"valid int32", int32(15), 10, 20, false},
		{"valid int64", int64(15), 10, 20, false},
		{"valid uint", uint(15), 10, 20, false},
		{"valid uint8", uint8(15), 10, 20, false},
		{"valid uint16", uint16(15), 10, 20, false},
		{"valid uint32", uint32(15), 10, 20, false},
		{"valid uint64", uint64(15), 10, 20, false},

		// Invalid - out of range
		{"invalid below min", 5, 10, 20, true},
		{"invalid above max", 25, 10, 20, true},
		{"invalid zero below range", -1, 0, 10, true},
		{"invalid negative below range", -15, -10, 0, true},

		// Invalid types
		{"invalid type float", 15.5, 10, 20, true},
		{"invalid type string", "15", 10, 20, true},
		{"invalid type nil", nil, 10, 20, true},
		{"invalid type bool", true, 10, 20, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "value")
			v.IntRange(tt.min, tt.max)
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("IntRange() expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("IntRange() unexpected error: %v", err)
			}
		})
	}
}

// TestIntRange_StopOnError tests IntRange with StopOnError
func TestIntRange_StopOnError(t *testing.T) {
	v := New(5, "value")
	v.StopOnError()
	v.IntRange(10, 20)
	v.IntRange(10, 20) // Should not add second error

	err := v.Validate()
	if err == nil {
		t.Fatal("expected error but got none")
	}

	errors := v.GetErrors()
	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(errors))
	}
}

// TestFloatRange tests the FloatRange validator
func TestFloatRange(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		min       float64
		max       float64
		wantError bool
	}{
		// Valid float values
		{"valid float at min", 10.0, 10.0, 20.0, false},
		{"valid float at max", 20.0, 10.0, 20.0, false},
		{"valid float in middle", 15.5, 10.0, 20.0, false},
		{"valid float32", float32(15.5), 10.0, 20.0, false},
		{"valid float64", float64(15.5), 10.0, 20.0, false},
		{"valid float zero", 0.0, 0.0, 10.0, false},
		{"valid float negative", -5.5, -10.0, 0.0, false},

		// Valid int values (converted to float)
		{"valid int as float", 15, 10.0, 20.0, false},
		{"valid int8 as float", int8(15), 10.0, 20.0, false},
		{"valid int16 as float", int16(15), 10.0, 20.0, false},
		{"valid int32 as float", int32(15), 10.0, 20.0, false},
		{"valid int64 as float", int64(15), 10.0, 20.0, false},
		{"valid uint as float", uint(15), 10.0, 20.0, false},

		// Invalid - out of range
		{"invalid below min", 5.0, 10.0, 20.0, true},
		{"invalid above max", 25.0, 10.0, 20.0, true},
		{"invalid float below", 9.99, 10.0, 20.0, true},
		{"invalid float above", 20.01, 10.0, 20.0, true},

		// Invalid types
		{"invalid type string", "15.5", 10.0, 20.0, true},
		{"invalid type nil", nil, 10.0, 20.0, true},
		{"invalid type bool", true, 10.0, 20.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "value")
			v.FloatRange(tt.min, tt.max)
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("FloatRange() expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("FloatRange() unexpected error: %v", err)
			}
		})
	}
}

// TestFloatRange_StopOnError tests FloatRange with StopOnError
func TestFloatRange_StopOnError(t *testing.T) {
	v := New(5.0, "value")
	v.StopOnError()
	v.FloatRange(10.0, 20.0)
	v.FloatRange(10.0, 20.0) // Should not add second error

	err := v.Validate()
	if err == nil {
		t.Fatal("expected error but got none")
	}

	errors := v.GetErrors()
	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(errors))
	}
}

// TestDateRange tests the DateRange validator
func TestDateRange(t *testing.T) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)

	tests := []struct {
		name      string
		value     interface{}
		start     time.Time
		end       time.Time
		wantError bool
	}{
		// Valid - time.Time values
		{"valid time.Time at start", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), start, end, false},
		{"valid time.Time at end", time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC), start, end, false},
		{"valid time.Time in middle", time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC), start, end, false},

		// Valid - RFC3339 strings
		{"valid RFC3339 at start", "2025-01-01T00:00:00Z", start, end, false},
		{"valid RFC3339 in middle", "2025-06-15T12:00:00Z", start, end, false},
		{"valid RFC3339 with timezone", "2025-06-15T12:00:00+09:00", start, end, false},

		// Valid - ISO 8601 strings
		{"valid ISO 8601 at start", "2025-01-01", start, end, false},
		{"valid ISO 8601 in middle", "2025-06-15", start, end, false},
		{"valid ISO 8601 at end", "2025-12-31", start, end, false},

		// Invalid - before start (time.Time)
		{"invalid time.Time before start", time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC), start, end, true},
		{"invalid time.Time way before", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), start, end, true},

		// Invalid - after end (time.Time)
		{"invalid time.Time after end", time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), start, end, true},
		{"invalid time.Time way after", time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC), start, end, true},

		// Invalid - before start (string)
		{"invalid RFC3339 before start", "2024-12-31T23:59:59Z", start, end, true},
		{"invalid ISO 8601 before start", "2024-12-31", start, end, true},

		// Invalid - after end (string)
		{"invalid RFC3339 after end", "2026-01-01T00:00:00Z", start, end, true},
		{"invalid ISO 8601 after end", "2026-01-01", start, end, true},

		// Invalid - malformed strings
		{"invalid malformed string", "not-a-date", start, end, true},
		{"invalid incomplete date", "2025-06", start, end, true},
		{"invalid empty string", "", start, end, true},

		// Invalid types
		{"invalid type int", 20250615, start, end, true},
		{"invalid type nil", nil, start, end, true},
		{"invalid type float", 20250615.0, start, end, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "date")
			v.DateRange(tt.start, tt.end)
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("DateRange() expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("DateRange() unexpected error: %v", err)
			}
		})
	}
}

// TestDateRange_StopOnError tests DateRange with StopOnError
func TestDateRange_StopOnError(t *testing.T) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	beforeStart := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	v := New(beforeStart, "date")
	v.StopOnError()
	v.DateRange(start, end)
	v.DateRange(start, end) // Should not add second error

	err := v.Validate()
	if err == nil {
		t.Fatal("expected error but got none")
	}

	errors := v.GetErrors()
	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(errors))
	}
}

// TestRangeValidators_Combined tests combining range validators
func TestRangeValidators_Combined(t *testing.T) {
	t.Run("valid int and float ranges", func(t *testing.T) {
		mv := NewValidator()
		mv.Field(25, "age").IntRange(18, 65)
		mv.Field(98.6, "temperature").FloatRange(95.0, 105.0)
		err := mv.Validate()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid int range", func(t *testing.T) {
		mv := NewValidator()
		mv.Field(15, "age").IntRange(18, 65)
		err := mv.Validate()

		if err == nil {
			t.Error("expected error but got none")
		}
	})

	t.Run("valid date and numeric ranges", func(t *testing.T) {
		start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
		testDate := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)

		mv := NewValidator()
		mv.Field(testDate, "event_date").DateRange(start, end)
		mv.Field(50, "capacity").IntRange(10, 100)
		mv.Field(25.5, "price").FloatRange(10.0, 100.0)
		err := mv.Validate()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

// TestToInt64Helper tests the toInt64 helper function
func TestToInt64Helper(t *testing.T) {
	v := &Validator{}

	tests := []struct {
		name      string
		value     interface{}
		expected  int64
		wantError bool
	}{
		{"int", 42, 42, false},
		{"int8", int8(42), 42, false},
		{"int16", int16(42), 42, false},
		{"int32", int32(42), 42, false},
		{"int64", int64(42), 42, false},
		{"uint", uint(42), 42, false},
		{"uint8", uint8(42), 42, false},
		{"uint16", uint16(42), 42, false},
		{"uint32", uint32(42), 42, false},
		{"uint64", uint64(42), 42, false},
		{"string", "42", 0, true},
		{"float", 42.5, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := v.toInt64(tt.value)

			if tt.wantError {
				if err == nil {
					t.Error("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("got %d, want %d", result, tt.expected)
				}
			}
		})
	}
}

// TestToFloat64Helper tests the toFloat64 helper function
func TestToFloat64Helper(t *testing.T) {
	v := &Validator{}

	tests := []struct {
		name      string
		value     interface{}
		expected  float64
		wantError bool
	}{
		{"float32", float32(42.5), 42.5, false},
		{"float64", float64(42.5), 42.5, false},
		{"int", 42, 42.0, false},
		{"int8", int8(42), 42.0, false},
		{"int16", int16(42), 42.0, false},
		{"int32", int32(42), 42.0, false},
		{"int64", int64(42), 42.0, false},
		{"uint", uint(42), 42.0, false},
		{"uint8", uint8(42), 42.0, false},
		{"uint16", uint16(42), 42.0, false},
		{"uint32", uint32(42), 42.0, false},
		{"uint64", uint64(42), 42.0, false},
		{"string", "42.5", 0, true},
		{"bool", true, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := v.toFloat64(tt.value)

			if tt.wantError {
				if err == nil {
					t.Error("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("got %f, want %f", result, tt.expected)
				}
			}
		})
	}
}
