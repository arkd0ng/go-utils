package validation

import (
	"testing"
)

func TestISBN(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid ISBN-10
		{"valid ISBN-10", "0-596-52068-9", false},
		{"valid ISBN-10 no hyphens", "0596520689", false},
		{"valid ISBN-10 with X", "043942089X", false},
		{"valid ISBN-10 with spaces", "0 596 52068 9", false},

		// Valid ISBN-13
		{"valid ISBN-13", "978-0-596-52068-7", false},
		{"valid ISBN-13 no hyphens", "9780596520687", false},
		{"valid ISBN-13 with spaces", "978 0 596 52068 7", false},
		{"valid ISBN-13 example 2", "978-3-16-148410-0", false},

		// Invalid ISBN
		{"invalid ISBN-10 checksum", "0-596-52068-0", true},
		{"invalid ISBN-13 checksum", "978-0-596-52068-0", true},
		{"invalid length", "123456", true},
		{"invalid ISBN-10 with letters", "059652068A", true},
		{"invalid ISBN-13 with letters", "978059652068A", true},
		{"empty string", "", true},

		// Type errors
		{"non-string int", 1234567890, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "isbn_field")
			v.ISBN()

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

func TestISSN(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid ISSN
		{"valid ISSN", "2049-3630", false},
		{"valid ISSN no hyphen", "20493630", false},
		{"valid ISSN with X", "0378-5955", false},
		{"valid ISSN with spaces", "2049 3630", false},
		{"valid ISSN ending with X", "0000-006X", false}, // Fixed: correct checksum

		// Invalid ISSN
		{"invalid ISSN checksum", "2049-3631", true},
		{"invalid length too short", "2049363", true},
		{"invalid length too long", "204936301", true},
		{"invalid with letters", "2049363A", true},
		{"empty string", "", true},

		// Type errors
		{"non-string int", 20493630, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "issn_field")
			v.ISSN()

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

func TestEAN(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid EAN-8
		{"valid EAN-8", "96385074", false},
		{"valid EAN-8 example 2", "73513537", false},

		// Valid EAN-13
		{"valid EAN-13", "4006381333931", false},
		{"valid EAN-13 example 2", "5901234123457", false},
		{"valid EAN-13 with hyphens", "400-6381-333-931", false},
		{"valid EAN-13 with spaces", "4006 381 333 931", false},

		// Invalid EAN
		{"invalid EAN-8 checksum", "96385075", true},
		{"invalid EAN-13 checksum", "4006381333932", true},
		{"invalid length", "12345", true},
		{"invalid EAN-8 with letters", "9638507A", true},
		{"invalid EAN-13 with letters", "400638133393A", true},
		{"empty string", "", true},

		// Type errors
		{"non-string int", 96385074, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "ean_field")
			v.EAN()

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

// Test StopOnError behavior for Business validators
func TestBusinessValidatorsStopOnError(t *testing.T) {
	t.Run("ISBN StopOnError", func(t *testing.T) {
		v := New("invalid", "isbn_field").StopOnError()
		v.ISBN()
		v.ISBN() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("ISSN StopOnError", func(t *testing.T) {
		v := New("invalid", "issn_field").StopOnError()
		v.ISSN()
		v.ISSN() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("EAN StopOnError", func(t *testing.T) {
		v := New("invalid", "ean_field").StopOnError()
		v.EAN()
		v.EAN() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})
}

// Test helper functions directly
func TestISBN10Validation(t *testing.T) {
	tests := []struct {
		name  string
		isbn  string
		want  bool
	}{
		{"valid ISBN-10", "0596520689", true},
		{"valid ISBN-10 with X", "043942089X", true},
		{"invalid checksum", "0596520680", false},
		{"invalid length", "059652068", false},
		{"invalid with letters", "059A520689", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidISBN10(tt.isbn)
			if result != tt.want {
				t.Errorf("isValidISBN10(%s) = %v, want %v", tt.isbn, result, tt.want)
			}
		})
	}
}

func TestISBN13Validation(t *testing.T) {
	tests := []struct {
		name  string
		isbn  string
		want  bool
	}{
		{"valid ISBN-13", "9780596520687", true},
		{"valid ISBN-13 example 2", "9783161484100", true},
		{"invalid checksum", "9780596520680", false},
		{"invalid length", "978059652068", false},
		{"invalid with letters", "978059652068A", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidISBN13(tt.isbn)
			if result != tt.want {
				t.Errorf("isValidISBN13(%s) = %v, want %v", tt.isbn, result, tt.want)
			}
		})
	}
}

func TestISSNValidation(t *testing.T) {
	tests := []struct {
		name string
		issn string
		want bool
	}{
		{"valid ISSN", "20493630", true},
		{"valid ISSN with X", "0000006X", true}, // Fixed: correct checksum
		{"invalid checksum", "20493631", false},
		{"invalid length", "2049363", false},
		{"invalid with letters", "2049363A", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidISSN(tt.issn)
			if result != tt.want {
				t.Errorf("isValidISSN(%s) = %v, want %v", tt.issn, result, tt.want)
			}
		})
	}
}

func TestEAN8Validation(t *testing.T) {
	tests := []struct {
		name string
		ean  string
		want bool
	}{
		{"valid EAN-8", "96385074", true},
		{"valid EAN-8 example 2", "73513537", true},
		{"invalid checksum", "96385075", false},
		{"invalid length", "9638507", false},
		{"invalid with letters", "9638507A", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidEAN8(tt.ean)
			if result != tt.want {
				t.Errorf("isValidEAN8(%s) = %v, want %v", tt.ean, result, tt.want)
			}
		})
	}
}

func TestEAN13Validation(t *testing.T) {
	tests := []struct {
		name string
		ean  string
		want bool
	}{
		{"valid EAN-13", "4006381333931", true},
		{"valid EAN-13 example 2", "5901234123457", true},
		{"invalid checksum", "4006381333932", false},
		{"invalid length", "400638133393", false},
		{"invalid with letters", "400638133393A", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidEAN13(tt.ean)
			if result != tt.want {
				t.Errorf("isValidEAN13(%s) = %v, want %v", tt.ean, result, tt.want)
			}
		})
	}
}

// Test business ID validation with chaining
func TestBusinessIDChaining(t *testing.T) {
	t.Run("Valid ISBN chain", func(t *testing.T) {
		v := New("978-0-596-52068-7", "isbn")
		v.Required().ISBN()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("Invalid ISBN chain stops on first error", func(t *testing.T) {
		v := New("invalid", "isbn").StopOnError()
		v.Required().ISBN()
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Invalid ISBN accumulates errors without StopOnError", func(t *testing.T) {
		v := New("", "isbn")
		v.Required().ISBN()
		errors := v.GetErrors()
		if len(errors) < 2 {
			t.Errorf("expected multiple errors without StopOnError, got %d", len(errors))
		}
	})
}
