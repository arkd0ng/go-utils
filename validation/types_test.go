package validation

import "testing"

func TestValidatorStruct(t *testing.T) {
	v := &Validator{
		value:     "test",
		fieldName: "name",
		errors:    []ValidationError{},
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
}

func TestMultiValidatorStruct(t *testing.T) {
	mv := &MultiValidator{
		validators: []*Validator{},
		errors:     []ValidationError{},
	}

	if len(mv.validators) != 0 {
		t.Errorf("Expected 0 validators, got %d", len(mv.validators))
	}

	if len(mv.errors) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(mv.errors))
	}
}

func TestRuleFunc(t *testing.T) {
	rule := func(v interface{}) bool {
		s, ok := v.(string)
		return ok && len(s) > 0
	}

	if !rule("test") {
		t.Error("Expected rule to return true for 'test'")
	}

	if rule("") {
		t.Error("Expected rule to return false for empty string")
	}

	if rule(123) {
		t.Error("Expected rule to return false for non-string")
	}
}
