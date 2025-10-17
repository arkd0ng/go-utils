package validation

import (
	"testing"
)

func TestTrue(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid
		{"true value", true, false},

		// Invalid
		{"false value", false, true},
		{"not boolean string", "true", true},
		{"not boolean int", 1, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.True()

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

func TestFalse(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid
		{"false value", false, false},

		// Invalid
		{"true value", true, true},
		{"not boolean string", "false", true},
		{"not boolean int", 0, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.False()

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

func TestNil(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid
		{"nil value", nil, false},
		{"nil pointer", (*string)(nil), false},
		{"nil slice", ([]int)(nil), false},
		{"nil map", (map[string]int)(nil), false},
		{"nil interface", (interface{})(nil), false},

		// Invalid
		{"non-nil string", "value", true},
		{"non-nil int", 0, true},
		{"non-nil bool", false, true},
		{"non-nil pointer", func() interface{} { s := "test"; return &s }(), true},
		{"non-nil slice", []int{}, true},
		{"non-nil map", map[string]int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Nil()

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

func TestNotNil(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid
		{"non-nil string", "value", false},
		{"non-nil int", 0, false},
		{"non-nil bool", false, false},
		{"non-nil pointer", func() interface{} { s := "test"; return &s }(), false},
		{"non-nil slice", []int{}, false},
		{"non-nil map", map[string]int{}, false},

		// Invalid
		{"nil value", nil, true},
		{"nil pointer", (*string)(nil), true},
		{"nil slice", ([]int)(nil), true},
		{"nil map", (map[string]int)(nil), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.NotNil()

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

func TestType(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		typeName  string
		wantError bool
	}{
		// Valid
		{"string type", "hello", "string", false},
		{"int type", 42, "int", false},
		{"bool type", true, "bool", false},
		{"float64 type", 3.14, "float64", false},
		{"slice type", []int{1, 2, 3}, "slice", false},
		{"map type", map[string]int{"a": 1}, "map", false},

		// Invalid
		{"string vs int", "hello", "int", true},
		{"int vs string", 42, "string", true},
		{"bool vs int", true, "int", true},
		{"nil value", nil, "string", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Type(tt.typeName)

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

func TestEmpty(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid (empty/zero values)
		{"empty string", "", false},
		{"zero int", 0, false},
		{"zero float", 0.0, false},
		{"false bool", false, false},
		{"nil value", nil, false},
		{"nil slice", ([]int)(nil), false},
		{"empty slice", []int{}, false},
		{"nil map", (map[string]int)(nil), false},
		{"empty map", map[string]int{}, false},
		{"nil pointer", (*string)(nil), false},

		// Invalid (non-empty values)
		{"non-empty string", "hello", true},
		{"non-zero int", 42, true},
		{"non-zero float", 3.14, true},
		{"true bool", true, true},
		{"non-empty slice", []int{1, 2, 3}, true},
		{"non-empty map", map[string]int{"a": 1}, true},
		{"non-nil pointer", func() interface{} { s := "test"; return &s }(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Empty()

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

func TestNotEmpty(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid (non-empty values)
		{"non-empty string", "hello", false},
		{"non-zero int", 42, false},
		{"non-zero float", 3.14, false},
		{"true bool", true, false},
		{"non-empty slice", []int{1, 2, 3}, false},
		{"non-empty map", map[string]int{"a": 1}, false},
		{"non-nil pointer", func() interface{} { s := "test"; return &s }(), false},

		// Invalid (empty/zero values)
		{"empty string", "", true},
		{"zero int", 0, true},
		{"zero float", 0.0, true},
		{"false bool", false, true},
		{"nil value", nil, true},
		{"nil slice", ([]int)(nil), true},
		{"empty slice", []int{}, true},
		{"nil map", (map[string]int)(nil), true},
		{"empty map", map[string]int{}, true},
		{"nil pointer", (*string)(nil), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.NotEmpty()

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

// Test StopOnError behavior for Type validators
func TestTypeValidatorsStopOnError(t *testing.T) {
	t.Run("True StopOnError", func(t *testing.T) {
		v := New(false, "field").StopOnError()
		v.True()
		v.True() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Nil StopOnError", func(t *testing.T) {
		v := New("not nil", "field").StopOnError()
		v.Nil()
		v.Nil() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Type StopOnError", func(t *testing.T) {
		v := New(123, "field").StopOnError()
		v.Type("string")
		v.Type("bool") // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})
}

// Test type validation with chaining
func TestTypeChaining(t *testing.T) {
	t.Run("True with chaining", func(t *testing.T) {
		v := New(true, "accepted")
		v.True()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("NotNil with chaining", func(t *testing.T) {
		str := "value"
		ptr := &str
		v := New(ptr, "pointer")
		v.NotNil()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("NotEmpty with chaining", func(t *testing.T) {
		v := New("hello", "text")
		v.Required().NotEmpty().MinLength(5)
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})
}

// Test edge cases
func TestTypeEdgeCases(t *testing.T) {
	t.Run("Empty slice vs nil slice", func(t *testing.T) {
		// Both should be considered empty
		v1 := New(([]int)(nil), "nil_slice")
		v1.Empty()
		if len(v1.GetErrors()) > 0 {
			t.Errorf("nil slice should be empty")
		}

		v2 := New([]int{}, "empty_slice")
		v2.Empty()
		if len(v2.GetErrors()) > 0 {
			t.Errorf("empty slice should be empty")
		}
	})

	t.Run("Empty map vs nil map", func(t *testing.T) {
		// Both should be considered empty
		v1 := New((map[string]int)(nil), "nil_map")
		v1.Empty()
		if len(v1.GetErrors()) > 0 {
			t.Errorf("nil map should be empty")
		}

		v2 := New(map[string]int{}, "empty_map")
		v2.Empty()
		if len(v2.GetErrors()) > 0 {
			t.Errorf("empty map should be empty")
		}
	})

	t.Run("Zero vs empty for numbers", func(t *testing.T) {
		// Zero should be considered empty
		v := New(0, "zero")
		v.Empty()
		if len(v.GetErrors()) > 0 {
			t.Errorf("zero should be empty")
		}
	})

	t.Run("False vs empty for bool", func(t *testing.T) {
		// False should be considered empty
		v := New(false, "false_bool")
		v.Empty()
		if len(v.GetErrors()) > 0 {
			t.Errorf("false should be empty")
		}
	})

	t.Run("Pointer to zero value", func(t *testing.T) {
		// Pointer to zero value should not be nil
		zero := 0
		v := New(&zero, "ptr_to_zero")
		v.NotNil()
		if len(v.GetErrors()) > 0 {
			t.Errorf("pointer to zero should not be nil")
		}
	})
}

// Test complex scenarios
func TestComplexTypeValidation(t *testing.T) {
	t.Run("Terms acceptance validation", func(t *testing.T) {
		accepted := true
		v := New(accepted, "terms_accepted")
		v.True()
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected valid terms acceptance")
		}
	})

	t.Run("Optional pointer validation", func(t *testing.T) {
		// Optional pointer can be nil
		var ptr *string
		v1 := New(ptr, "optional")
		v1.Nil()
		if len(v1.GetErrors()) > 0 {
			t.Errorf("expected nil pointer to be valid")
		}

		// Or have a value
		str := "value"
		v2 := New(&str, "optional")
		v2.NotNil()
		if len(v2.GetErrors()) > 0 {
			t.Errorf("expected non-nil pointer to be valid")
		}
	})

	t.Run("Collection validation", func(t *testing.T) {
		// Empty slice validation
		emptySlice := []string{}
		v1 := New(emptySlice, "tags")
		v1.Empty()
		if len(v1.GetErrors()) > 0 {
			t.Errorf("expected empty slice to be valid")
		}

		// Non-empty slice validation
		nonEmptySlice := []string{"tag1", "tag2"}
		v2 := New(nonEmptySlice, "tags")
		v2.NotEmpty()
		if len(v2.GetErrors()) > 0 {
			t.Errorf("expected non-empty slice to be valid")
		}
	})

	t.Run("Type safety validation", func(t *testing.T) {
		// Ensure type is correct
		value := "hello"
		v := New(value, "text")
		v.Type("string").MinLength(5)
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected valid string type")
		}
	})
}
