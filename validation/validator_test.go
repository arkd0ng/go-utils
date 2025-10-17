package validation

import (
	"testing"
)

func TestNew(t *testing.T) {
	v := New("test", "name")

	if v == nil {
		t.Fatal("Expected non-nil validator")
	}

	if v.value != "test" {
		t.Errorf("Expected value 'test', got %v", v.value)
	}

	if v.fieldName != "name" {
		t.Errorf("Expected fieldName 'name', got %v", v.fieldName)
	}

	if len(v.errors) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(v.errors))
	}

	if v.stopOnError {
		t.Error("Expected stopOnError to be false")
	}
}

func TestValidate(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		v := New("test", "name")
		err := v.Validate()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("with errors", func(t *testing.T) {
		v := New("", "name")
		v.addError("required", "name is required")
		err := v.Validate()

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		verrs, ok := err.(ValidationErrors)
		if !ok {
			t.Fatal("Expected ValidationErrors type")
		}

		if len(verrs) != 1 {
			t.Errorf("Expected 1 error, got %d", len(verrs))
		}
	})
}

func TestGetErrors(t *testing.T) {
	v := New("", "name")
	v.addError("required", "name is required")
	v.addError("minlength", "name too short")

	errors := v.GetErrors()

	if len(errors) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(errors))
	}
}

func TestStopOnError(t *testing.T) {
	v := New("", "name")
	v.StopOnError()

	if !v.stopOnError {
		t.Error("Expected stopOnError to be true")
	}

	// Add first error
	v.addError("required", "name is required")
	if len(v.errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(v.errors))
	}

	// Try to add second error - should be ignored
	v.addError("minlength", "name too short")
	if len(v.errors) != 1 {
		t.Errorf("Expected still 1 error due to StopOnError, got %d", len(v.errors))
	}
}

func TestWithMessage(t *testing.T) {
	v := New("", "name")
	v.addError("required", "name is required")
	v.WithMessage("Custom error message")

	if len(v.errors) != 1 {
		t.Fatalf("Expected 1 error, got %d", len(v.errors))
	}

	if v.errors[0].Message != "Custom error message" {
		t.Errorf("Expected custom message, got %q", v.errors[0].Message)
	}
}

func TestAddError(t *testing.T) {
	v := New("test", "name")
	v.addError("custom", "custom error")

	if len(v.errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(v.errors))
	}

	err := v.errors[0]
	if err.Field != "name" {
		t.Errorf("Expected field 'name', got %q", err.Field)
	}

	if err.Value != "test" {
		t.Errorf("Expected value 'test', got %v", err.Value)
	}

	if err.Rule != "custom" {
		t.Errorf("Expected rule 'custom', got %q", err.Rule)
	}

	if err.Message != "custom error" {
		t.Errorf("Expected message 'custom error', got %q", err.Message)
	}
}

func TestCustom(t *testing.T) {
	t.Run("passing validation", func(t *testing.T) {
		v := New("test", "name")
		v.Custom(func(val interface{}) bool {
			s, ok := val.(string)
			return ok && len(s) > 0
		}, "must not be empty")

		err := v.Validate()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("failing validation", func(t *testing.T) {
		v := New("", "name")
		v.Custom(func(val interface{}) bool {
			s, ok := val.(string)
			return ok && len(s) > 0
		}, "must not be empty")

		err := v.Validate()
		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		verrs := err.(ValidationErrors)
		if len(verrs) != 1 {
			t.Errorf("Expected 1 error, got %d", len(verrs))
		}

		if verrs[0].Message != "must not be empty" {
			t.Errorf("Expected custom message, got %q", verrs[0].Message)
		}
	})
}

func TestNewValidator(t *testing.T) {
	mv := NewValidator()

	if mv == nil {
		t.Fatal("Expected non-nil MultiValidator")
	}

	if len(mv.validators) != 0 {
		t.Errorf("Expected 0 validators, got %d", len(mv.validators))
	}

	if len(mv.errors) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(mv.errors))
	}
}

func TestMultiValidatorField(t *testing.T) {
	mv := NewValidator()
	v := mv.Field("test", "name")

	if v == nil {
		t.Fatal("Expected non-nil validator")
	}

	if len(mv.validators) != 1 {
		t.Errorf("Expected 1 validator, got %d", len(mv.validators))
	}

	if v.fieldName != "name" {
		t.Errorf("Expected fieldName 'name', got %q", v.fieldName)
	}

	if v.value != "test" {
		t.Errorf("Expected value 'test', got %v", v.value)
	}
}

func TestMultiValidatorValidate(t *testing.T) {
	t.Run("all valid", func(t *testing.T) {
		mv := NewValidator()
		mv.Field("test", "name")
		mv.Field("test@example.com", "email")

		err := mv.Validate()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("with errors", func(t *testing.T) {
		mv := NewValidator()
		v1 := mv.Field("", "name")
		v1.addError("required", "name is required")

		v2 := mv.Field("", "email")
		v2.addError("required", "email is required")

		err := mv.Validate()
		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		verrs, ok := err.(ValidationErrors)
		if !ok {
			t.Fatal("Expected ValidationErrors type")
		}

		if len(verrs) != 2 {
			t.Errorf("Expected 2 errors, got %d", len(verrs))
		}
	})
}

func TestMultiValidatorGetErrors(t *testing.T) {
	mv := NewValidator()
	v1 := mv.Field("", "name")
	v1.addError("required", "name is required")

	v2 := mv.Field("", "email")
	v2.addError("required", "email is required")

	// Must call Validate first to collect errors
	mv.Validate()

	errors := mv.GetErrors()
	if len(errors) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(errors))
	}
}

func TestValidateString(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		v := New("test", "name")
		validateString(v, "test", func(s string) bool {
			return len(s) > 0
		}, "must not be empty")

		if len(v.errors) != 0 {
			t.Errorf("Expected 0 errors, got %d", len(v.errors))
		}
	})

	t.Run("invalid string", func(t *testing.T) {
		v := New("", "name")
		validateString(v, "test", func(s string) bool {
			return len(s) > 0
		}, "must not be empty")

		if len(v.errors) != 1 {
			t.Errorf("Expected 1 error, got %d", len(v.errors))
		}
	})

	t.Run("non-string value", func(t *testing.T) {
		v := New(123, "name")
		validateString(v, "test", func(s string) bool {
			return true
		}, "must not be empty")

		if len(v.errors) != 1 {
			t.Errorf("Expected 1 error, got %d", len(v.errors))
		}

		if v.errors[0].Message != "name must be a string" {
			t.Errorf("Expected type error, got %q", v.errors[0].Message)
		}
	})
}

func TestValidateNumeric(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		valid bool
	}{
		{"int", 42, true},
		{"int8", int8(42), true},
		{"int16", int16(42), true},
		{"int32", int32(42), true},
		{"int64", int64(42), true},
		{"uint", uint(42), true},
		{"uint8", uint8(42), true},
		{"uint16", uint16(42), true},
		{"uint32", uint32(42), true},
		{"uint64", uint64(42), true},
		{"float32", float32(42.5), true},
		{"float64", float64(42.5), true},
		{"string", "42", false},
		{"bool", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			validateNumeric(v, "test", func(n float64) bool {
				return n > 0
			}, "must be positive")

			if tt.valid {
				if len(v.errors) != 0 {
					t.Errorf("Expected 0 errors for valid %s, got %d", tt.name, len(v.errors))
				}
			} else {
				if len(v.errors) != 1 {
					t.Errorf("Expected 1 error for invalid %s, got %d", tt.name, len(v.errors))
				}
			}
		})
	}
}
