package validation

import (
	"testing"
)

func TestASCII(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid ASCII
		{"valid ASCII text", "Hello World 123", false},
		{"valid ASCII with symbols", "!@#$%^&*()", false},
		{"valid ASCII numbers", "0123456789", false},
		{"valid ASCII with newline", "Hello\nWorld", false},
		{"valid ASCII with tab", "Hello\tWorld", false},
		{"empty string", "", false},
		{"ASCII control chars", "\x00\x01\x7F", false},

		// Invalid ASCII (contains non-ASCII characters)
		{"invalid UTF-8 한글", "Hello 한글", true},
		{"invalid UTF-8 emoji", "Hello 😀", true},
		{"invalid UTF-8 Chinese", "你好", true},
		{"invalid UTF-8 Japanese", "こんにちは", true},

		// Type errors
		{"non-string int", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "ascii_field")
			v.ASCII()

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

func TestPrintable(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid printable ASCII
		{"valid printable text", "Hello World 123", false},
		{"valid printable with symbols", "!@#$%^&*()", false},
		{"valid printable numbers", "0123456789", false},
		{"valid printable space", "   ", false},
		{"empty string", "", false},

		// Invalid (contains control characters)
		{"invalid with newline", "Hello\nWorld", true},
		{"invalid with tab", "Hello\tWorld", true},
		{"invalid with NULL", "Hello\x00World", true},
		{"invalid with DEL", "Hello\x7FWorld", true},
		{"invalid UTF-8 한글", "Hello 한글", true},

		// Type errors
		{"non-string int", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "printable_field")
			v.Printable()

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

func TestWhitespace(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid whitespace
		{"valid single space", " ", false},
		{"valid multiple spaces", "   ", false},
		{"valid tab", "\t", false},
		{"valid newline", "\n", false},
		{"valid mixed whitespace", " \t\n  ", false},
		{"valid carriage return", "\r", false},

		// Invalid (contains non-whitespace)
		{"invalid with letter", " a ", true},
		{"invalid with number", " 1 ", true},
		{"invalid text", "Hello World", true},
		{"empty string", "", true},

		// Type errors
		{"non-string int", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "whitespace_field")
			v.Whitespace()

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

func TestAlphaSpace(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid alpha + space
		{"valid name", "John Doe", false},
		{"valid with spaces", "Hello World", false},
		{"valid only letters", "HelloWorld", false},
		{"valid only space", " ", false},
		{"valid multiple spaces", "Hello  World", false},
		{"empty string", "", false},

		// Invalid (contains numbers or special chars)
		{"invalid with number", "Hello 123", true},
		{"invalid with symbol", "Hello!", true},
		{"invalid with dash", "Hello-World", true},
		{"invalid with underscore", "Hello_World", true},
		{"invalid with tab", "Hello\tWorld", true},
		{"invalid with newline", "Hello\nWorld", true},

		// Type errors
		{"non-string int", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "alpha_space_field")
			v.AlphaSpace()

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

// Test StopOnError behavior for Data validators
func TestDataValidatorsStopOnError(t *testing.T) {
	t.Run("ASCII StopOnError", func(t *testing.T) {
		v := New("한글", "ascii_field").StopOnError()
		v.ASCII()
		v.ASCII() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Printable StopOnError", func(t *testing.T) {
		v := New("Hello\nWorld", "printable_field").StopOnError()
		v.Printable()
		v.Printable() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Whitespace StopOnError", func(t *testing.T) {
		v := New("not whitespace", "whitespace_field").StopOnError()
		v.Whitespace()
		v.Whitespace() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("AlphaSpace StopOnError", func(t *testing.T) {
		v := New("Hello123", "alpha_space_field").StopOnError()
		v.AlphaSpace()
		v.AlphaSpace() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})
}

// Test data format validation with chaining
func TestDataFormatChaining(t *testing.T) {
	t.Run("Valid ASCII chain", func(t *testing.T) {
		v := New("Hello World", "text")
		v.Required().ASCII().Printable()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("Invalid ASCII chain stops on first error", func(t *testing.T) {
		v := New("한글", "text").StopOnError()
		v.Required().ASCII().Printable()
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Valid AlphaSpace chain", func(t *testing.T) {
		v := New("John Doe", "name")
		v.Required().AlphaSpace()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("Invalid AlphaSpace accumulates errors without StopOnError", func(t *testing.T) {
		v := New("", "name")
		v.Required().AlphaSpace()
		errors := v.GetErrors()
		if len(errors) < 1 {
			t.Errorf("expected at least 1 error without StopOnError, got %d", len(errors))
		}
	})
}

// Test edge cases
func TestDataFormatEdgeCases(t *testing.T) {
	t.Run("ASCII with all control characters", func(t *testing.T) {
		v := New("\x00\x01\x02\x03\x04\x05", "control")
		v.ASCII()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors for control characters in ASCII, got %v", v.GetErrors())
		}
	})

	t.Run("Printable rejects control characters", func(t *testing.T) {
		v := New("\x00\x01\x02\x03\x04\x05", "control")
		v.Printable()
		if len(v.GetErrors()) == 0 {
			t.Errorf("expected error for control characters in Printable")
		}
	})

	t.Run("ASCII boundary character 127", func(t *testing.T) {
		v := New("\x7F", "boundary")
		v.ASCII()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors for ASCII boundary character 127, got %v", v.GetErrors())
		}
	})

	t.Run("ASCII character 128 fails", func(t *testing.T) {
		v := New("\x80", "boundary")
		v.ASCII()
		if len(v.GetErrors()) == 0 {
			t.Errorf("expected error for non-ASCII character 128")
		}
	})

	t.Run("Printable boundary character 32 (space)", func(t *testing.T) {
		v := New(" ", "space")
		v.Printable()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors for printable boundary character 32, got %v", v.GetErrors())
		}
	})

	t.Run("Printable boundary character 126 (~)", func(t *testing.T) {
		v := New("~", "tilde")
		v.Printable()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors for printable boundary character 126, got %v", v.GetErrors())
		}
	})

	t.Run("AlphaSpace with Unicode letters", func(t *testing.T) {
		v := New("Café", "unicode")
		v.AlphaSpace()
		// é is Unicode letter, should pass
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors for Unicode letters, got %v", v.GetErrors())
		}
	})
}
