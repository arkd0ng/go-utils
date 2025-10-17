package validation

import (
	"testing"
)

func TestOneOf(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		allowed   []interface{}
		wantError bool
	}{
		// Valid - value matches one of allowed
		{"string match first", "active", []interface{}{"active", "inactive"}, false},
		{"string match last", "inactive", []interface{}{"active", "inactive"}, false},
		{"string match middle", "pending", []interface{}{"active", "pending", "inactive"}, false},
		{"int match", 1, []interface{}{1, 2, 3}, false},
		{"int match zero", 0, []interface{}{0, 1, 2}, false},
		{"float match", 1.5, []interface{}{1.0, 1.5, 2.0}, false},
		{"bool match true", true, []interface{}{true, false}, false},
		{"bool match false", false, []interface{}{true, false}, false},
		{"single allowed", "only", []interface{}{"only"}, false},

		// Invalid - value doesn't match any allowed
		{"string no match", "invalid", []interface{}{"active", "inactive"}, true},
		{"int no match", 5, []interface{}{1, 2, 3}, true},
		{"float no match", 2.5, []interface{}{1.0, 1.5, 2.0}, true},
		{"case sensitive", "Active", []interface{}{"active", "inactive"}, true},
		{"empty string no match", "", []interface{}{"active", "inactive"}, true},

		// Edge cases
		{"nil value with nil allowed", nil, []interface{}{nil, "value"}, false},
		{"nil value no match", nil, []interface{}{"value1", "value2"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.OneOf(tt.allowed...)

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

func TestOneOfNoValues(t *testing.T) {
	v := New("test", "field")
	v.OneOf() // No values provided

	if len(v.GetErrors()) == 0 {
		t.Errorf("expected error when no values provided")
	}
}

func TestNotOneOf(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		forbidden []interface{}
		wantError bool
	}{
		// Valid - value doesn't match any forbidden
		{"string not in list", "active", []interface{}{"admin", "root"}, false},
		{"int not in list", 5, []interface{}{1, 2, 3}, false},
		{"float not in list", 2.5, []interface{}{1.0, 1.5, 2.0}, false},
		{"case sensitive different", "Active", []interface{}{"active", "admin"}, false},
		{"empty string allowed", "", []interface{}{"admin", "root"}, false},

		// Invalid - value matches forbidden
		{"string match first", "admin", []interface{}{"admin", "root"}, true},
		{"string match last", "root", []interface{}{"admin", "root"}, true},
		{"string match middle", "root", []interface{}{"admin", "root", "superuser"}, true},
		{"int match", 1, []interface{}{1, 2, 3}, true},
		{"float match", 1.5, []interface{}{1.0, 1.5, 2.0}, true},
		{"bool match", true, []interface{}{true}, true},
		{"single forbidden match", "only", []interface{}{"only"}, true},

		// Edge cases
		{"nil value with nil forbidden", nil, []interface{}{nil, "value"}, true},
		{"nil value allowed", nil, []interface{}{"value1", "value2"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.NotOneOf(tt.forbidden...)

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

func TestNotOneOfNoValues(t *testing.T) {
	v := New("test", "field")
	v.NotOneOf() // No values provided

	if len(v.GetErrors()) == 0 {
		t.Errorf("expected error when no values provided")
	}
}

func TestWhen(t *testing.T) {
	t.Run("predicate true - validation executes", func(t *testing.T) {
		v := New("", "email")
		isRequired := true

		v.When(isRequired, func(val *Validator) {
			val.Required()
		})

		if len(v.GetErrors()) == 0 {
			t.Errorf("expected error from Required validation")
		}
	})

	t.Run("predicate false - validation skipped", func(t *testing.T) {
		v := New("", "email")
		isRequired := false

		v.When(isRequired, func(val *Validator) {
			val.Required()
		})

		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no error when predicate is false, got %v", v.GetErrors())
		}
	})

	t.Run("complex validation when true", func(t *testing.T) {
		v := New("test", "password")
		isStrict := true

		v.When(isStrict, func(val *Validator) {
			val.MinLength(8)
		})

		if len(v.GetErrors()) == 0 {
			t.Errorf("expected error from MinLength validation")
		}
	})

	t.Run("complex validation when false", func(t *testing.T) {
		v := New("test", "password")
		isStrict := false

		v.When(isStrict, func(val *Validator) {
			val.MinLength(8)
		})

		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no error when predicate is false")
		}
	})

	t.Run("multiple validations in When", func(t *testing.T) {
		v := New("test@", "email")
		shouldValidate := true

		v.When(shouldValidate, func(val *Validator) {
			val.Required().Email().MinLength(10)
		})

		// Should have errors because email is invalid and too short
		if len(v.GetErrors()) == 0 {
			t.Errorf("expected errors from validation chain")
		}
	})
}

func TestUnless(t *testing.T) {
	t.Run("predicate false - validation executes", func(t *testing.T) {
		v := New("", "email")
		isGuest := false

		v.Unless(isGuest, func(val *Validator) {
			val.Required()
		})

		if len(v.GetErrors()) == 0 {
			t.Errorf("expected error from Required validation")
		}
	})

	t.Run("predicate true - validation skipped", func(t *testing.T) {
		v := New("", "email")
		isGuest := true

		v.Unless(isGuest, func(val *Validator) {
			val.Required()
		})

		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no error when predicate is true, got %v", v.GetErrors())
		}
	})

	t.Run("complex validation unless true", func(t *testing.T) {
		v := New("test", "password")
		isSimple := true

		v.Unless(isSimple, func(val *Validator) {
			val.MinLength(8)
		})

		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no error when predicate is true")
		}
	})

	t.Run("complex validation unless false", func(t *testing.T) {
		v := New("test", "password")
		isSimple := false

		v.Unless(isSimple, func(val *Validator) {
			val.MinLength(8)
		})

		if len(v.GetErrors()) == 0 {
			t.Errorf("expected error from MinLength validation")
		}
	})

	t.Run("multiple validations in Unless", func(t *testing.T) {
		v := New("test@", "email")
		skipValidation := false

		v.Unless(skipValidation, func(val *Validator) {
			val.Required().Email().MinLength(10)
		})

		// Should have errors because email is invalid and too short
		if len(v.GetErrors()) == 0 {
			t.Errorf("expected errors from validation chain")
		}
	})
}

// Test StopOnError behavior for Logical validators
func TestLogicalValidatorsStopOnError(t *testing.T) {
	t.Run("OneOf StopOnError", func(t *testing.T) {
		v := New("invalid", "status").StopOnError()
		v.OneOf("active", "inactive")
		v.OneOf("yes", "no") // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("NotOneOf StopOnError", func(t *testing.T) {
		v := New("admin", "username").StopOnError()
		v.NotOneOf("admin", "root")
		v.NotOneOf("superuser") // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("When with StopOnError", func(t *testing.T) {
		v := New("", "email").StopOnError()
		v.Required()
		v.When(true, func(val *Validator) {
			val.Email() // Should not execute due to StopOnError
		})
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Unless with StopOnError", func(t *testing.T) {
		v := New("", "email").StopOnError()
		v.Required()
		v.Unless(false, func(val *Validator) {
			val.Email() // Should not execute due to StopOnError
		})
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})
}

// Test logical validation with chaining
func TestLogicalChaining(t *testing.T) {
	t.Run("OneOf with other validators", func(t *testing.T) {
		v := New("active", "status")
		v.Required().OneOf("active", "inactive", "pending")
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("NotOneOf with other validators", func(t *testing.T) {
		v := New("user123", "username")
		v.Required().NotOneOf("admin", "root").MinLength(5)
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("When with chaining inside", func(t *testing.T) {
		v := New("user@example.com", "email")
		shouldValidateEmail := true

		v.Required().When(shouldValidateEmail, func(val *Validator) {
			val.Email().MinLength(5)
		})

		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("Unless with chaining inside", func(t *testing.T) {
		v := New("shortpw", "password")
		isGuest := false

		v.Required().Unless(isGuest, func(val *Validator) {
			val.MinLength(8) // Should execute and fail
		})

		if len(v.GetErrors()) == 0 {
			t.Errorf("expected error from MinLength")
		}
	})
}

// Test edge cases
func TestLogicalEdgeCases(t *testing.T) {
	t.Run("OneOf with mixed types", func(t *testing.T) {
		v := New(1, "field")
		v.OneOf(1, "1", true) // int, string, bool
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors for matching int")
		}
	})

	t.Run("NotOneOf with mixed types", func(t *testing.T) {
		v := New("test", "field")
		v.NotOneOf(1, true, 3.14) // Different types
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors when types differ")
		}
	})

	t.Run("When with nested When", func(t *testing.T) {
		v := New("test", "field")
		condition1 := true
		condition2 := true

		v.When(condition1, func(val *Validator) {
			val.When(condition2, func(val2 *Validator) {
				val2.MinLength(10) // Should fail
			})
		})

		if len(v.GetErrors()) == 0 {
			t.Errorf("expected error from nested When")
		}
	})

	t.Run("Unless with nested Unless", func(t *testing.T) {
		v := New("test", "field")
		condition1 := false
		condition2 := false

		v.Unless(condition1, func(val *Validator) {
			val.Unless(condition2, func(val2 *Validator) {
				val2.MinLength(10) // Should fail
			})
		})

		if len(v.GetErrors()) == 0 {
			t.Errorf("expected error from nested Unless")
		}
	})

	t.Run("When and Unless combined", func(t *testing.T) {
		v := New("test@example.com", "email")
		isUser := true
		isGuest := false

		v.When(isUser, func(val *Validator) {
			val.Required()
		}).Unless(isGuest, func(val *Validator) {
			val.Email()
		})

		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})
}
