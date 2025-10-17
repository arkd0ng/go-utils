package validation

import (
	"testing"
)

func TestLatitude(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid latitudes
		{"valid zero", 0.0, false},
		{"valid positive", 37.5665, false},
		{"valid negative", -37.5665, false},
		{"valid max", 90.0, false},
		{"valid min", -90.0, false},
		{"valid float32", float32(45.0), false},
		{"valid int", 45, false},
		{"valid int64", int64(45), false},
		{"valid string", "37.5665", false},
		{"valid string negative", "-37.5665", false},

		// Invalid latitudes
		{"invalid too high", 90.1, true},
		{"invalid too low", -90.1, true},
		{"invalid way too high", 180.0, true},
		{"invalid way too low", -180.0, true},
		{"invalid string non-numeric", "abc", true},
		{"invalid string empty", "", true},

		// Type errors
		{"invalid bool", true, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "latitude_field")
			v.Latitude()

			if tt.wantError {
				if len(v.GetErrors()) == 0 {
					t.Errorf("expected error but got none")
				}
			} else {
				if len(v.GetErrors()) > 0 {
					t.Errorf("expected no error but got %v", v.GetErrors())
				}
			}
		})
	}
}

func TestLongitude(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid longitudes
		{"valid zero", 0.0, false},
		{"valid positive", 126.9780, false},
		{"valid negative", -122.4194, false},
		{"valid max", 180.0, false},
		{"valid min", -180.0, false},
		{"valid float32", float32(90.0), false},
		{"valid int", 90, false},
		{"valid int64", int64(90), false},
		{"valid string", "126.9780", false},
		{"valid string negative", "-122.4194", false},

		// Invalid longitudes
		{"invalid too high", 180.1, true},
		{"invalid too low", -180.1, true},
		{"invalid way too high", 360.0, true},
		{"invalid way too low", -360.0, true},
		{"invalid string non-numeric", "xyz", true},
		{"invalid string empty", "", true},

		// Type errors
		{"invalid bool", false, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "longitude_field")
			v.Longitude()

			if tt.wantError {
				if len(v.GetErrors()) == 0 {
					t.Errorf("expected error but got none")
				}
			} else {
				if len(v.GetErrors()) > 0 {
					t.Errorf("expected no error but got %v", v.GetErrors())
				}
			}
		})
	}
}

func TestCoordinate(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid coordinates
		{"valid Seoul", "37.5665,126.9780", false},
		{"valid with space", "37.5665, 126.9780", false},
		{"valid New York", "40.7128,-74.0060", false},
		{"valid London", "51.5074,-0.1278", false},
		{"valid zero", "0,0", false},
		{"valid max lat", "90,0", false},
		{"valid min lat", "-90,0", false},
		{"valid max lon", "0,180", false},
		{"valid min lon", "0,-180", false},
		{"valid max both", "90,180", false},
		{"valid min both", "-90,-180", false},

		// Invalid coordinates - format
		{"invalid no comma", "37.5665 126.9780", true},
		{"invalid too many parts", "37.5665,126.9780,100", true},
		{"invalid one part", "37.5665", true},
		{"invalid empty", "", true},

		// Invalid coordinates - latitude out of range
		{"invalid lat too high", "90.1,126.9780", true},
		{"invalid lat too low", "-90.1,126.9780", true},
		{"invalid lat way too high", "180,126.9780", true},

		// Invalid coordinates - longitude out of range
		{"invalid lon too high", "37.5665,180.1", true},
		{"invalid lon too low", "37.5665,-180.1", true},
		{"invalid lon way too high", "37.5665,360", true},

		// Invalid coordinates - non-numeric
		{"invalid lat non-numeric", "abc,126.9780", true},
		{"invalid lon non-numeric", "37.5665,xyz", true},
		{"invalid both non-numeric", "abc,xyz", true},

		// Type errors
		{"non-string int", 123, true},
		{"non-string float", 123.45, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "coordinate_field")
			v.Coordinate()

			if tt.wantError {
				if len(v.GetErrors()) == 0 {
					t.Errorf("expected error but got none")
				}
			} else {
				if len(v.GetErrors()) > 0 {
					t.Errorf("expected no error but got %v", v.GetErrors())
				}
			}
		})
	}
}

// Test StopOnError behavior for Geographic validators
func TestGeographicValidatorsStopOnError(t *testing.T) {
	t.Run("Latitude StopOnError", func(t *testing.T) {
		v := New(200.0, "lat_field").StopOnError()
		v.Latitude()
		v.Latitude() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Longitude StopOnError", func(t *testing.T) {
		v := New(200.0, "lon_field").StopOnError()
		v.Longitude()
		v.Longitude() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Coordinate StopOnError", func(t *testing.T) {
		v := New("invalid", "coord_field").StopOnError()
		v.Coordinate()
		v.Coordinate() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})
}

// Test geographic validation with chaining
func TestGeographicChaining(t *testing.T) {
	t.Run("Valid coordinate chain", func(t *testing.T) {
		v := New("37.5665,126.9780", "location")
		v.Required().Coordinate()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("Invalid coordinate chain stops on first error", func(t *testing.T) {
		v := New("invalid", "location").StopOnError()
		v.Required().Coordinate()
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Multi-field geographic validation", func(t *testing.T) {
		mv := NewValidator()
		mv.Field(37.5665, "latitude").Latitude()
		mv.Field(126.9780, "longitude").Longitude()
		mv.Field("37.5665,126.9780", "coordinate").Coordinate()

		err := mv.Validate()
		if err != nil {
			t.Errorf("expected no errors but got %v", err)
		}
	})
}

// Test edge cases
func TestGeographicEdgeCases(t *testing.T) {
	t.Run("Latitude exactly at boundaries", func(t *testing.T) {
		v1 := New(90.0, "lat")
		v1.Latitude()
		if len(v1.GetErrors()) > 0 {
			t.Error("90.0 should be valid latitude")
		}

		v2 := New(-90.0, "lat")
		v2.Latitude()
		if len(v2.GetErrors()) > 0 {
			t.Error("-90.0 should be valid latitude")
		}
	})

	t.Run("Longitude exactly at boundaries", func(t *testing.T) {
		v1 := New(180.0, "lon")
		v1.Longitude()
		if len(v1.GetErrors()) > 0 {
			t.Error("180.0 should be valid longitude")
		}

		v2 := New(-180.0, "lon")
		v2.Longitude()
		if len(v2.GetErrors()) > 0 {
			t.Error("-180.0 should be valid longitude")
		}
	})

	t.Run("Coordinate with extra spaces", func(t *testing.T) {
		v := New("  37.5665  ,  126.9780  ", "coord")
		v.Coordinate()
		if len(v.GetErrors()) > 0 {
			t.Error("Coordinate with extra spaces should be valid")
		}
	})

	t.Run("Coordinate at boundaries", func(t *testing.T) {
		v := New("90,180", "coord")
		v.Coordinate()
		if len(v.GetErrors()) > 0 {
			t.Error("90,180 should be valid coordinate")
		}
	})
}
