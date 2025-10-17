package validation

import (
	"testing"
	"time"
)

// TestDateFormat tests the DateFormat validator
func TestDateFormat(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		format    string
		wantError bool
	}{
		// Valid date formats - ISO 8601
		{"valid ISO 8601", "2025-10-17", "2006-01-02", false},
		{"valid ISO 8601 jan", "2025-01-01", "2006-01-02", false},
		{"valid ISO 8601 dec", "2025-12-31", "2006-01-02", false},

		// Valid date formats - US format
		{"valid US format", "10/17/2025", "01/02/2006", false},
		{"valid US format jan", "01/01/2025", "01/02/2006", false},

		// Valid date formats - EU format
		{"valid EU format", "17/10/2025", "02/01/2006", false},
		{"valid EU format jan", "01/01/2025", "02/01/2006", false},

		// Valid date formats - slash separator
		{"valid slash format", "2025/10/17", "2006/01/02", false},

		// Invalid date formats
		{"invalid format mismatch", "2025-10-17", "01/02/2006", true},
		{"invalid date string", "not-a-date", "2006-01-02", true},
		{"invalid incomplete", "2025-10", "2006-01-02", true},
		{"invalid day", "2025-02-30", "2006-01-02", true},
		{"invalid month", "2025-13-01", "2006-01-02", true},
		{"invalid empty string", "", "2006-01-02", true},

		// Invalid types
		{"invalid type int", 20251017, "2006-01-02", true},
		{"invalid type nil", nil, "2006-01-02", true},
		{"invalid type time.Time", time.Now(), "2006-01-02", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "date")
			v.DateFormat(tt.format)
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("DateFormat() expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("DateFormat() unexpected error: %v", err)
			}
		})
	}
}

// TestDateFormat_StopOnError tests DateFormat with StopOnError
func TestDateFormat_StopOnError(t *testing.T) {
	v := New("invalid", "date")
	v.StopOnError()
	v.DateFormat("2006-01-02")
	v.DateFormat("2006-01-02") // Should not add second error

	err := v.Validate()
	if err == nil {
		t.Fatal("expected error but got none")
	}

	errors := v.GetErrors()
	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(errors))
	}
}

// TestTimeFormat tests the TimeFormat validator
func TestTimeFormat(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		format    string
		wantError bool
	}{
		// Valid time formats - 24-hour
		{"valid 24h HMS", "14:30:00", "15:04:05", false},
		{"valid 24h HM", "14:30", "15:04", false},
		{"valid 24h midnight", "00:00:00", "15:04:05", false},
		{"valid 24h noon", "12:00:00", "15:04:05", false},
		{"valid 24h end of day", "23:59:59", "15:04:05", false},

		// Valid time formats - 12-hour
		{"valid 12h AM HMS", "02:30:00 AM", "03:04:05 PM", false},
		{"valid 12h PM HMS", "02:30:00 PM", "03:04:05 PM", false},
		{"valid 12h AM HM", "02:30 AM", "03:04 PM", false},
		{"valid 12h PM HM", "02:30 PM", "03:04 PM", false},

		// Invalid time formats
		{"invalid format mismatch", "14:30:00", "15:04", true},
		{"invalid time string", "not-a-time", "15:04:05", true},
		{"invalid hour", "25:00:00", "15:04:05", true},
		{"invalid minute", "14:60:00", "15:04:05", true},
		{"invalid second", "14:30:60", "15:04:05", true},
		{"invalid incomplete", "14:30:", "15:04:05", true},
		{"invalid empty string", "", "15:04:05", true},

		// Invalid types
		{"invalid type int", 143000, "15:04:05", true},
		{"invalid type nil", nil, "15:04:05", true},
		{"invalid type time.Time", time.Now(), "15:04:05", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "time")
			v.TimeFormat(tt.format)
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("TimeFormat() expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("TimeFormat() unexpected error: %v", err)
			}
		})
	}
}

// TestTimeFormat_StopOnError tests TimeFormat with StopOnError
func TestTimeFormat_StopOnError(t *testing.T) {
	v := New("invalid", "time")
	v.StopOnError()
	v.TimeFormat("15:04:05")
	v.TimeFormat("15:04:05") // Should not add second error

	err := v.Validate()
	if err == nil {
		t.Fatal("expected error but got none")
	}

	errors := v.GetErrors()
	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(errors))
	}
}

// TestDateBefore tests the DateBefore validator
func TestDateBefore(t *testing.T) {
	baseDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name      string
		value     interface{}
		before    time.Time
		wantError bool
	}{
		// Valid - time.Time values
		{"valid time.Time 1 day before", time.Date(2025, 10, 16, 12, 0, 0, 0, time.UTC), baseDate, false},
		{"valid time.Time 1 month before", time.Date(2025, 9, 17, 12, 0, 0, 0, time.UTC), baseDate, false},
		{"valid time.Time 1 year before", time.Date(2024, 10, 17, 12, 0, 0, 0, time.UTC), baseDate, false},
		{"valid time.Time far past", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), baseDate, false},

		// Valid - RFC3339 strings
		{"valid RFC3339", "2025-10-16T12:00:00Z", baseDate, false},
		{"valid RFC3339 with timezone", "2025-10-16T12:00:00+09:00", baseDate, false},

		// Valid - ISO 8601 strings
		{"valid ISO 8601", "2025-10-16", baseDate, false},
		{"valid ISO 8601 far past", "2000-01-01", baseDate, false},

		// Invalid - dates on or after baseDate (time.Time)
		{"invalid same date", time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC), baseDate, true},
		{"invalid 1 day after", time.Date(2025, 10, 18, 12, 0, 0, 0, time.UTC), baseDate, true},
		{"invalid 1 month after", time.Date(2025, 11, 17, 12, 0, 0, 0, time.UTC), baseDate, true},

		// Invalid - dates on or after baseDate (string)
		{"invalid same date RFC3339", "2025-10-17T12:00:00Z", baseDate, true},
		{"invalid same date ISO 8601", "2025-10-17T23:59:59Z", baseDate, true}, // Use time component to ensure it's after 12:00
		{"invalid after date ISO 8601", "2025-10-18", baseDate, true},

		// Invalid - malformed strings
		{"invalid malformed string", "not-a-date", baseDate, true},
		{"invalid incomplete date", "2025-10", baseDate, true},
		{"invalid empty string", "", baseDate, true},

		// Invalid types
		{"invalid type int", 20251017, baseDate, true},
		{"invalid type nil", nil, baseDate, true},
		{"invalid type float", 20251017.0, baseDate, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "date")
			v.DateBefore(tt.before)
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("DateBefore() expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("DateBefore() unexpected error: %v", err)
			}
		})
	}
}

// TestDateBefore_StopOnError tests DateBefore with StopOnError
func TestDateBefore_StopOnError(t *testing.T) {
	baseDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)
	afterDate := time.Date(2025, 10, 18, 12, 0, 0, 0, time.UTC)

	v := New(afterDate, "date")
	v.StopOnError()
	v.DateBefore(baseDate)
	v.DateBefore(baseDate) // Should not add second error

	err := v.Validate()
	if err == nil {
		t.Fatal("expected error but got none")
	}

	errors := v.GetErrors()
	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(errors))
	}
}

// TestDateAfter tests the DateAfter validator
func TestDateAfter(t *testing.T) {
	baseDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name      string
		value     interface{}
		after     time.Time
		wantError bool
	}{
		// Valid - time.Time values
		{"valid time.Time 1 day after", time.Date(2025, 10, 18, 12, 0, 0, 0, time.UTC), baseDate, false},
		{"valid time.Time 1 month after", time.Date(2025, 11, 17, 12, 0, 0, 0, time.UTC), baseDate, false},
		{"valid time.Time 1 year after", time.Date(2026, 10, 17, 12, 0, 0, 0, time.UTC), baseDate, false},
		{"valid time.Time far future", time.Date(2030, 12, 31, 23, 59, 59, 0, time.UTC), baseDate, false},

		// Valid - RFC3339 strings
		{"valid RFC3339", "2025-10-18T12:00:00Z", baseDate, false},
		{"valid RFC3339 with timezone", "2025-10-18T12:00:00+09:00", baseDate, false},

		// Valid - ISO 8601 strings
		{"valid ISO 8601", "2025-10-18", baseDate, false},
		{"valid ISO 8601 far future", "2030-12-31", baseDate, false},

		// Invalid - dates on or before baseDate (time.Time)
		{"invalid same date", time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC), baseDate, true},
		{"invalid 1 day before", time.Date(2025, 10, 16, 12, 0, 0, 0, time.UTC), baseDate, true},
		{"invalid 1 month before", time.Date(2025, 9, 17, 12, 0, 0, 0, time.UTC), baseDate, true},

		// Invalid - dates on or before baseDate (string)
		{"invalid same date RFC3339", "2025-10-17T12:00:00Z", baseDate, true},
		{"invalid same date ISO 8601", "2025-10-17", baseDate, true},
		{"invalid before date ISO 8601", "2025-10-16", baseDate, true},

		// Invalid - malformed strings
		{"invalid malformed string", "not-a-date", baseDate, true},
		{"invalid incomplete date", "2025-10", baseDate, true},
		{"invalid empty string", "", baseDate, true},

		// Invalid types
		{"invalid type int", 20251017, baseDate, true},
		{"invalid type nil", nil, baseDate, true},
		{"invalid type float", 20251017.0, baseDate, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "date")
			v.DateAfter(tt.after)
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("DateAfter() expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("DateAfter() unexpected error: %v", err)
			}
		})
	}
}

// TestDateAfter_StopOnError tests DateAfter with StopOnError
func TestDateAfter_StopOnError(t *testing.T) {
	baseDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)
	beforeDate := time.Date(2025, 10, 16, 12, 0, 0, 0, time.UTC)

	v := New(beforeDate, "date")
	v.StopOnError()
	v.DateAfter(baseDate)
	v.DateAfter(baseDate) // Should not add second error

	err := v.Validate()
	if err == nil {
		t.Fatal("expected error but got none")
	}

	errors := v.GetErrors()
	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(errors))
	}
}

// TestDateTimeValidators_Combined tests combining date/time validators
func TestDateTimeValidators_Combined(t *testing.T) {
	t.Run("valid date format and range", func(t *testing.T) {
		minDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		maxDate := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
		testDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)

		v := New(testDate, "date")
		v.DateAfter(minDate).DateBefore(maxDate)
		err := v.Validate()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid date range", func(t *testing.T) {
		minDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		maxDate := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
		testDate := time.Date(2024, 10, 17, 12, 0, 0, 0, time.UTC) // Before minDate

		v := New(testDate, "date")
		v.DateAfter(minDate).DateBefore(maxDate)
		err := v.Validate()

		if err == nil {
			t.Error("expected error but got none")
		}
	})

	t.Run("valid date string format and parsing", func(t *testing.T) {
		v := New("2025-10-17", "date")
		v.DateFormat("2006-01-02")
		err := v.Validate()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
