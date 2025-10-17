package validation

import "testing"

func TestRequired(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "test", false},
		{"empty", "", true},
		{"spaces only", "   ", true},
		{"with spaces", "  test  ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Required()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Required() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinLength(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		min     int
		wantErr bool
	}{
		{"valid", "test", 3, false},
		{"exact", "test", 4, false},
		{"too short", "ab", 3, true},
		{"unicode", "안녕", 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.MinLength(tt.min)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("MinLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMaxLength(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		max     int
		wantErr bool
	}{
		{"valid", "test", 10, false},
		{"exact", "test", 4, false},
		{"too long", "toolong", 3, true},
		{"unicode", "안녕하세요", 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.MaxLength(tt.max)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("MaxLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		length  int
		wantErr bool
	}{
		{"exact", "test", 4, false},
		{"too short", "ab", 4, true},
		{"too long", "toolong", 4, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Length(tt.length)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Length() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmail(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "test@example.com", false},
		{"valid with dots", "test.user@example.com", false},
		{"valid with plus", "test+tag@example.com", false},
		{"invalid no @", "testexample.com", true},
		{"invalid no domain", "test@", true},
		{"invalid no username", "@example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "email")
			v.Email()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Email() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestURL(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid http", "http://example.com", false},
		{"valid https", "https://example.com", false},
		{"valid with path", "https://example.com/path", false},
		{"invalid no protocol", "example.com", true},
		{"invalid ftp", "ftp://example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "url")
			v.URL()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("URL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlpha(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "abc", false},
		{"valid uppercase", "ABC", false},
		{"valid mixed", "AbC", false},
		{"invalid with numbers", "abc123", true},
		{"invalid with spaces", "a b c", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Alpha()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Alpha() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlphanumeric(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid letters", "abc", false},
		{"valid numbers", "123", false},
		{"valid mixed", "abc123", false},
		{"invalid with space", "abc 123", true},
		{"invalid with symbol", "abc@123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Alphanumeric()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Alphanumeric() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNumeric(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "123", false},
		{"invalid with letters", "123abc", true},
		{"invalid empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Numeric()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Numeric() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		prefix  string
		wantErr bool
	}{
		{"valid", "hello world", "hello", false},
		{"invalid", "world hello", "hello", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.StartsWith(tt.prefix)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("StartsWith() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		suffix  string
		wantErr bool
	}{
		{"valid", "hello world", "world", false},
		{"invalid", "world hello", "world", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.EndsWith(tt.suffix)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("EndsWith() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		substring string
		wantErr   bool
	}{
		{"valid", "hello world", "world", false},
		{"invalid", "hello", "world", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Contains(tt.substring)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Contains() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegex(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		pattern string
		wantErr bool
	}{
		{"valid", "abc123", `^[a-z]+[0-9]+$`, false},
		{"invalid", "123abc", `^[a-z]+[0-9]+$`, true},
		{"bad pattern", "test", `[`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Regex(tt.pattern)
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Regex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUUID(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "123e4567-e89b-12d3-a456-426614174000", false},
		{"valid uppercase", "123E4567-E89B-12D3-A456-426614174000", false},
		{"invalid", "not-a-uuid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "uuid")
			v.UUID()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSON(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid object", `{"key":"value"}`, false},
		{"valid array", `[1,2,3]`, false},
		{"invalid", `not json`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "json")
			v.JSON()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("JSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBase64(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "aGVsbG8=", false},
		{"invalid", "not-base64!", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "base64")
			v.Base64()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Base64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLowercase(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "hello", false},
		{"invalid", "Hello", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Lowercase()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Lowercase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUppercase(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "HELLO", false},
		{"invalid", "Hello", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "field")
			v.Uppercase()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Uppercase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPhone(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "+1234567890", false},
		{"valid with dashes", "123-456-7890", false},
		{"valid with spaces", "123 456 7890", false},
		{"too short", "12345", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "phone")
			v.Phone()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Phone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreditCard(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "1234567890123456", false},
		{"valid with spaces", "1234 5678 9012 3456", false},
		{"valid with dashes", "1234-5678-9012-3456", false},
		{"too short", "123456789012", true},
		{"too long", "12345678901234567890", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "card")
			v.CreditCard()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("CreditCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChaining(t *testing.T) {
	v := New("test", "username")
	v.Required().MinLength(3).MaxLength(20).Alphanumeric()

	err := v.Validate()
	if err != nil {
		t.Errorf("Chaining should pass, got error: %v", err)
	}
}

func TestChainingWithErrors(t *testing.T) {
	v := New("", "username")
	v.Required().MinLength(3).MaxLength(20)

	err := v.Validate()
	if err == nil {
		t.Error("Expected validation errors")
	}

	verrs, ok := err.(ValidationErrors)
	if !ok {
		t.Fatal("Expected ValidationErrors type")
	}

	// Should have at least one error (Required)
	if len(verrs) < 1 {
		t.Errorf("Expected at least 1 error, got %d", len(verrs))
	}
}
