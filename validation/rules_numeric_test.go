package validation

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		min     float64
		wantErr bool
	}{
		{"int valid", 10, 5.0, false},
		{"int invalid", 3, 5.0, true},
		{"float valid", 10.5, 5.0, false},
		{"float invalid", 3.5, 5.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.Min(tt.min)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Min() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		max     float64
		wantErr bool
	}{
		{"int valid", 3, 5.0, false},
		{"int invalid", 10, 5.0, true},
		{"float valid", 3.5, 5.0, false},
		{"float invalid", 10.5, 5.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.Max(tt.max)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Max() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		min     float64
		max     float64
		wantErr bool
	}{
		{"valid", 5, 1.0, 10.0, false},
		{"at min", 1, 1.0, 10.0, false},
		{"at max", 10, 1.0, 10.0, false},
		{"below min", 0, 1.0, 10.0, true},
		{"above max", 11, 1.0, 10.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.Between(tt.min, tt.max)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Between() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPositive(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"positive int", 5, false},
		{"positive float", 5.5, false},
		{"zero", 0, true},
		{"negative", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.Positive()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Positive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNegative(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"negative int", -5, false},
		{"negative float", -5.5, false},
		{"zero", 0, true},
		{"positive", 5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.Negative()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Negative() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestZero(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"zero int", 0, false},
		{"zero float", 0.0, false},
		{"non-zero positive", 5, true},
		{"non-zero negative", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.Zero()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Zero() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNonZero(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"positive", 5, false},
		{"negative", -5, false},
		{"zero int", 0, true},
		{"zero float", 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.NonZero()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("NonZero() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEven(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"even positive", 4, false},
		{"even negative", -4, false},
		{"zero", 0, false},
		{"odd positive", 5, true},
		{"odd negative", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.Even()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Even() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOdd(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"odd positive", 5, false},
		{"odd negative", -5, false},
		{"even positive", 4, true},
		{"even negative", -4, true},
		{"zero", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.Odd()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Odd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMultipleOf(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		n       float64
		wantErr bool
	}{
		{"multiple of 5", 10, 5.0, false},
		{"multiple of 3", 9, 3.0, false},
		{"not multiple", 7, 3.0, true},
		{"zero divisor", 10, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.MultipleOf(tt.n)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("MultipleOf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNumericChaining(t *testing.T) {
	v := New(25, "age")
	v.Positive().Min(18).Max(120)

	err := v.Validate()
	if err != nil {
		t.Errorf("Chaining should pass, got error: %v", err)
	}
}

func TestNumericChainingWithErrors(t *testing.T) {
	v := New(150, "age")
	v.Positive().Min(18).Max(120)

	err := v.Validate()
	if err == nil {
		t.Error("Expected validation error for age > 120")
	}
}
