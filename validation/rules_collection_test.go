package validation

import (
	"testing"
)

func TestIn(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		allowed   []interface{}
		wantError bool
	}{
		{"string in list", "red", []interface{}{"red", "green", "blue"}, false},
		{"string not in list", "yellow", []interface{}{"red", "green", "blue"}, true},
		{"int in list", 2, []interface{}{1, 2, 3}, false},
		{"int not in list", 5, []interface{}{1, 2, 3}, true},
		{"empty allowed list", "any", []interface{}{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.In(tt.allowed...)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("In() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestNotIn(t *testing.T) {
	tests := []struct {
		name       string
		value      interface{}
		forbidden  []interface{}
		wantError  bool
	}{
		{"string not in forbidden", "yellow", []interface{}{"red", "green"}, false},
		{"string in forbidden", "red", []interface{}{"red", "green"}, true},
		{"int not in forbidden", 5, []interface{}{1, 2, 3}, false},
		{"int in forbidden", 2, []interface{}{1, 2, 3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.NotIn(tt.forbidden...)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("NotIn() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestArrayLength(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		length    int
		wantError bool
	}{
		{"exact length", []int{1, 2, 3}, 3, false},
		{"wrong length - too short", []int{1, 2}, 3, true},
		{"wrong length - too long", []int{1, 2, 3, 4}, 3, true},
		{"empty array with length 0", []string{}, 0, false},
		{"not an array", "string", 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.ArrayLength(tt.length)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("ArrayLength() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestArrayMinLength(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		min       int
		wantError bool
	}{
		{"exact min length", []int{1, 2, 3}, 3, false},
		{"above min length", []int{1, 2, 3, 4}, 3, false},
		{"below min length", []int{1, 2}, 3, true},
		{"empty array with min 0", []string{}, 0, false},
		{"empty array with min 1", []string{}, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.ArrayMinLength(tt.min)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("ArrayMinLength() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestArrayMaxLength(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		max       int
		wantError bool
	}{
		{"exact max length", []int{1, 2, 3}, 3, false},
		{"below max length", []int{1, 2}, 3, false},
		{"above max length", []int{1, 2, 3, 4}, 3, true},
		{"empty array", []string{}, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.ArrayMaxLength(tt.max)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("ArrayMaxLength() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestArrayNotEmpty(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		{"non-empty array", []int{1, 2, 3}, false},
		{"single element", []string{"a"}, false},
		{"empty array", []int{}, true},
		{"not an array", "string", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.ArrayNotEmpty()
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("ArrayNotEmpty() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestArrayUnique(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		{"unique integers", []int{1, 2, 3, 4}, false},
		{"duplicate integers", []int{1, 2, 2, 3}, true},
		{"unique strings", []string{"a", "b", "c"}, false},
		{"duplicate strings", []string{"a", "b", "a"}, true},
		{"empty array", []int{}, false},
		{"single element", []string{"a"}, false},
		{"not an array", "string", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.ArrayUnique()
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("ArrayUnique() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestMapHasKey(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		key       interface{}
		wantError bool
	}{
		{"map has key", map[string]int{"age": 25}, "age", false},
		{"map missing key", map[string]int{"age": 25}, "name", true},
		{"empty map", map[string]string{}, "key", true},
		{"not a map", "string", "key", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.MapHasKey(tt.key)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("MapHasKey() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestMapHasKeys(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		keys      []interface{}
		wantError bool
	}{
		{
			"map has all keys",
			map[string]int{"name": 1, "age": 25, "city": 3},
			[]interface{}{"name", "age"},
			false,
		},
		{
			"map missing one key",
			map[string]int{"name": 1, "age": 25},
			[]interface{}{"name", "email"},
			true,
		},
		{
			"map missing multiple keys",
			map[string]int{"name": 1},
			[]interface{}{"age", "email"},
			true,
		},
		{
			"empty keys list",
			map[string]int{"name": 1},
			[]interface{}{},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.MapHasKeys(tt.keys...)
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("MapHasKeys() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestMapNotEmpty(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		{"non-empty map", map[string]int{"key": 1}, false},
		{"empty map", map[string]string{}, true},
		{"not a map", "string", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.MapNotEmpty()
			err := v.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("MapNotEmpty() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestCollectionChaining(t *testing.T) {
	// Test chaining multiple collection validators
	// 여러 컬렉션 검증자 체이닝 테스트
	t.Run("valid array chaining", func(t *testing.T) {
		v := New([]int{1, 2, 3}, "numbers")
		v.ArrayNotEmpty().ArrayMinLength(2).ArrayMaxLength(5).ArrayUnique()
		err := v.Validate()
		if err != nil {
			t.Errorf("Valid array chaining failed: %v", err)
		}
	})

	t.Run("invalid array chaining", func(t *testing.T) {
		v := New([]int{1, 2, 2, 3}, "numbers")
		v.ArrayNotEmpty().ArrayLength(4).ArrayUnique()
		err := v.Validate()
		if err == nil {
			t.Error("Expected error for duplicate elements")
		}
	})

	t.Run("valid map chaining", func(t *testing.T) {
		data := map[string]int{"name": 1, "age": 25, "city": 3}
		v := New(data, "user")
		v.MapNotEmpty().MapHasKeys("name", "age")
		err := v.Validate()
		if err != nil {
			t.Errorf("Valid map chaining failed: %v", err)
		}
	})

	t.Run("invalid map chaining", func(t *testing.T) {
		data := map[string]int{"name": 1}
		v := New(data, "user")
		v.MapNotEmpty().MapHasKeys("name", "age", "email")
		err := v.Validate()
		if err == nil {
			t.Error("Expected error for missing keys")
		}
	})
}
