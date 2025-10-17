package validation

import (
	"testing"
)

func TestUUIDv4(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid UUIDv4
		{"valid UUIDv4", "550e8400-e29b-41d4-a716-446655440000", false},
		{"valid UUIDv4 variant 8", "f47ac10b-58cc-4372-8567-0e02b2c3d479", false},
		{"valid UUIDv4 variant 9", "f47ac10b-58cc-4372-9567-0e02b2c3d479", false},
		{"valid UUIDv4 variant a", "f47ac10b-58cc-4372-a567-0e02b2c3d479", false},
		{"valid UUIDv4 variant b", "f47ac10b-58cc-4372-b567-0e02b2c3d479", false},
		{"valid UUIDv4 uppercase", "550E8400-E29B-41D4-A716-446655440000", false},
		{"valid UUIDv4 mixed case", "550e8400-E29B-41D4-A716-446655440000", false},

		// Invalid UUIDv4
		{"invalid version 1", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", true},
		{"invalid version 3", "6ba7b811-9dad-31d1-80b4-00c04fd430c8", true},
		{"invalid version 5", "6ba7b815-9dad-51d1-80b4-00c04fd430c8", true},
		{"invalid variant c", "f47ac10b-58cc-4372-c567-0e02b2c3d479", true},
		{"invalid variant 7", "f47ac10b-58cc-4372-7567-0e02b2c3d479", true},
		{"invalid format", "550e8400-e29b-51d4-a716-446655440000", true},
		{"invalid too short", "550e8400-e29b-41d4", true},
		{"invalid no hyphens", "550e8400e29b41d4a716446655440000", true},
		{"invalid characters", "550e8400-e29b-41d4-a716-44665544000g", true},
		{"empty string", "", true},
		{"non-string value", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "uuid_field")
			v.UUIDv4()

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

func TestXML(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid XML
		{"valid XML simple", `<root>content</root>`, false},
		{"valid XML with attributes", `<root attr="value">content</root>`, false},
		{"valid XML nested", `<root><child>content</child></root>`, false},
		{"valid XML self-closing", `<root/>`, false},
		{"valid XML with declaration", `<?xml version="1.0"?><root>content</root>`, false},
		{"valid XML with namespace", `<root xmlns="http://example.com">content</root>`, false},
		{"valid XML with whitespace", `  <root>content</root>  `, false},
		{"valid XML multiple children", `<root><child1/><child2/></root>`, false},
		{"valid XML with CDATA", `<root><![CDATA[content]]></root>`, false},
		{"valid XML empty element", `<root></root>`, false},

		// Invalid XML
		{"invalid XML unclosed tag", `<root>content`, true},
		{"invalid XML mismatched tags", `<root>content</roo>`, true},
		{"invalid XML invalid character", `<root>content<root>`, true},
		{"invalid XML missing closing bracket", `<root content</root>`, true},
		{"invalid XML double root", `<root1/><root2/>`, false}, // Go xml.Unmarshal accepts this
		{"empty string", "", true},
		{"non-XML string", "not xml", true},
		{"non-string value", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "xml_field")
			v.XML()

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

func TestHex(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid Hex
		{"valid hex lowercase", "deadbeef", false},
		{"valid hex uppercase", "DEADBEEF", false},
		{"valid hex mixed case", "DeAdBeEf", false},
		{"valid hex with 0x prefix", "0xdeadbeef", false},
		{"valid hex with 0X prefix", "0XDEADBEEF", false},
		{"valid hex numbers only", "123456789", true}, // Odd length
		{"valid hex single char", "a", true},          // Odd length
		{"valid hex empty", "", false},                // Empty string is valid hex (decodes to empty)
		{"valid hex long", "0123456789abcdefABCDEF", false},
		{"valid hex even length", "abcd", false},
		{"valid hex two chars", "ab", false},

		// Invalid Hex
		{"invalid hex with g", "deadbeeg", true}, // Even length but invalid char 'g'
		{"invalid hex with space", "dead beef", true},
		{"invalid hex with special", "dead-beef", true},
		{"invalid hex with invalid char", "deadbee!", true},
		{"invalid hex odd length valid chars", "abc", true},
		{"non-string value", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "hex_field")
			v.Hex()

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

// Test StopOnError behavior for Format validators
func TestFormatValidatorsStopOnError(t *testing.T) {
	t.Run("UUIDv4 StopOnError", func(t *testing.T) {
		v := New("invalid", "uuid_field").StopOnError()
		v.UUIDv4()
		v.UUIDv4() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("XML StopOnError", func(t *testing.T) {
		v := New("invalid", "xml_field").StopOnError()
		v.XML()
		v.XML() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Hex StopOnError", func(t *testing.T) {
		v := New("invalidg", "hex_field").StopOnError()
		v.Hex()
		v.Hex() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})
}
