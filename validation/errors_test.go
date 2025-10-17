package validation

import (
	"testing"
)

func TestValidationErrorError(t *testing.T) {
	tests := []struct {
		name     string
		err      ValidationError
		expected string
	}{
		{
			name: "with custom message",
			err: ValidationError{
				Field:   "email",
				Value:   "invalid",
				Rule:    "email",
				Message: "email must be valid",
			},
			expected: "email must be valid",
		},
		{
			name: "without custom message",
			err: ValidationError{
				Field: "name",
				Value: "",
				Rule:  "required",
			},
			expected: "required validation failed for field 'name'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.expected {
				t.Errorf("Error() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestValidationErrorsError(t *testing.T) {
	tests := []struct {
		name     string
		errors   ValidationErrors
		expected string
	}{
		{
			name:     "empty errors",
			errors:   ValidationErrors{},
			expected: "",
		},
		{
			name: "single error",
			errors: ValidationErrors{
				{Field: "email", Rule: "email", Message: "email must be valid"},
			},
			expected: "email must be valid",
		},
		{
			name: "multiple errors",
			errors: ValidationErrors{
				{Field: "email", Rule: "email", Message: "email must be valid"},
				{Field: "name", Rule: "required", Message: "name is required"},
			},
			expected: "email must be valid; name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.errors.Error()
			if got != tt.expected {
				t.Errorf("Error() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestValidationErrorsHasField(t *testing.T) {
	errors := ValidationErrors{
		{Field: "email", Rule: "email"},
		{Field: "name", Rule: "required"},
	}

	tests := []struct {
		field    string
		expected bool
	}{
		{"email", true},
		{"name", true},
		{"age", false},
	}

	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			got := errors.HasField(tt.field)
			if got != tt.expected {
				t.Errorf("HasField(%q) = %v, want %v", tt.field, got, tt.expected)
			}
		})
	}
}

func TestValidationErrorsGetField(t *testing.T) {
	errors := ValidationErrors{
		{Field: "email", Rule: "email", Message: "invalid email"},
		{Field: "email", Rule: "required", Message: "email required"},
		{Field: "name", Rule: "required", Message: "name required"},
	}

	emailErrors := errors.GetField("email")
	if len(emailErrors) != 2 {
		t.Errorf("Expected 2 email errors, got %d", len(emailErrors))
	}

	nameErrors := errors.GetField("name")
	if len(nameErrors) != 1 {
		t.Errorf("Expected 1 name error, got %d", len(nameErrors))
	}

	ageErrors := errors.GetField("age")
	if len(ageErrors) != 0 {
		t.Errorf("Expected 0 age errors, got %d", len(ageErrors))
	}
}

func TestValidationErrorsToMap(t *testing.T) {
	errors := ValidationErrors{
		{Field: "email", Rule: "email", Message: "invalid email"},
		{Field: "email", Rule: "required", Message: "email required"},
		{Field: "name", Rule: "required", Message: "name required"},
	}

	m := errors.ToMap()

	if len(m) != 2 {
		t.Errorf("Expected map with 2 keys, got %d", len(m))
	}

	if len(m["email"]) != 2 {
		t.Errorf("Expected 2 email messages, got %d", len(m["email"]))
	}

	if len(m["name"]) != 1 {
		t.Errorf("Expected 1 name message, got %d", len(m["name"]))
	}
}

func TestValidationErrorsFirst(t *testing.T) {
	t.Run("with errors", func(t *testing.T) {
		errors := ValidationErrors{
			{Field: "email", Rule: "email", Message: "invalid email"},
			{Field: "name", Rule: "required", Message: "name required"},
		}

		first := errors.First()
		if first == nil {
			t.Fatal("Expected non-nil error")
		}

		if first.Field != "email" {
			t.Errorf("Expected first error field 'email', got %q", first.Field)
		}
	})

	t.Run("empty errors", func(t *testing.T) {
		errors := ValidationErrors{}
		first := errors.First()
		if first != nil {
			t.Error("Expected nil for empty errors")
		}
	})
}

func TestValidationErrorsCount(t *testing.T) {
	tests := []struct {
		name     string
		errors   ValidationErrors
		expected int
	}{
		{
			name:     "empty",
			errors:   ValidationErrors{},
			expected: 0,
		},
		{
			name: "single error",
			errors: ValidationErrors{
				{Field: "email", Rule: "email"},
			},
			expected: 1,
		},
		{
			name: "multiple errors",
			errors: ValidationErrors{
				{Field: "email", Rule: "email"},
				{Field: "name", Rule: "required"},
				{Field: "age", Rule: "min"},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.errors.Count()
			if got != tt.expected {
				t.Errorf("Count() = %d, want %d", got, tt.expected)
			}
		})
	}
}
