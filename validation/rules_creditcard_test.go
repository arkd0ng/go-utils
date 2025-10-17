package validation

import (
	"testing"
)

func TestCreditCard(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid credit card numbers (using standard test card numbers)
		{"valid Visa 16 digits", "4532015112830366", false},
		{"valid Visa 13 digits", "4532015112830", false},
		{"valid Mastercard", "5425233430109903", false},
		{"valid Amex", "374245455400126", false},
		{"valid Discover", "6011111111111117", false},
		{"valid JCB", "3530111333300000", false},
		{"valid Diners Club", "30569309025904", false},

		// Valid with spaces and hyphens (should be cleaned)
		{"valid with spaces", "4532 0151 1283 0366", false},
		{"valid with hyphens", "4532-0151-1283-0366", false},
		{"valid with mixed", "4532 0151-1283 0366", false},

		// Invalid credit card numbers
		{"invalid Luhn check", "4532015112830367", true}, // Last digit changed
		{"invalid too short", "453201511283", true},      // 12 digits
		{"invalid too long", "45320151128303661234", true}, // 20 digits
		{"invalid non-digits", "453201511283036a", true},
		{"invalid special chars", "4532-0151-1283-036!", true},
		{"invalid empty", "", true},
		{"invalid all zeros", "0000000000000000", false}, // Technically passes Luhn (sum=0, 0%10=0)

		// Type errors
		{"non-string int", 123, true},
		{"non-string float", 123.45, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "card_field")
			v.CreditCard()

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

func TestCreditCardType(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		cardType  string
		wantError bool
	}{
		// Visa tests (starts with 4, 13 or 16 digits)
		{"valid Visa 16", "4532015112830366", "visa", false},
		{"valid Visa 13", "4532015112830", "visa", false},
		{"valid Visa uppercase", "4532015112830366", "VISA", false},
		{"invalid Visa wrong pattern", "4532015112830366", "mastercard", true},

		// Mastercard tests (starts with 51-55, 16 digits)
		{"valid Mastercard 51", "5105105105105100", "mastercard", false},
		{"valid Mastercard 55", "5555555555554444", "mastercard", false},
		{"invalid Mastercard wrong pattern", "5105105105105100", "visa", true},

		// Amex tests (starts with 34 or 37, 15 digits)
		{"valid Amex 34", "340000000000009", "amex", false},
		{"valid Amex 37", "378282246310005", "amex", false},
		{"invalid Amex wrong pattern", "378282246310005", "visa", true},

		// Discover tests (starts with 6011 or 65, 16 digits)
		{"valid Discover 6011", "6011111111111117", "discover", false},
		{"valid Discover 65", "6500000000000002", "discover", false},
		{"invalid Discover wrong pattern", "6011111111111117", "visa", true},

		// JCB tests (starts with 2131, 1800, or 35, 16 digits)
		{"valid JCB 35", "3530111333300000", "jcb", false},
		{"valid JCB 2131", "2131000000000008", "jcb", false},
		{"valid JCB 1800", "1800000000000000", "jcb", false}, // Fixed: correct Luhn checksum
		{"invalid JCB wrong pattern", "3530111333300000", "visa", true},

		// Diners Club tests (starts with 300-305 or 36 or 38, 14 digits)
		{"valid Diners Club 30", "30569309025904", "dinersclub", false},
		{"valid Diners Club 36", "36000000000008", "dinersclub", false},
		{"valid Diners Club 38", "38000000000006", "dinersclub", false},
		{"invalid Diners Club wrong pattern", "30569309025904", "visa", true},

		// UnionPay tests (starts with 62, 16-19 digits)
		{"valid UnionPay 16", "6200000000000005", "unionpay", false},
		{"valid UnionPay 19", "6200000000000000000", "unionpay", false}, // Fixed: correct Luhn checksum
		{"invalid UnionPay wrong pattern", "6200000000000005", "visa", true},

		// Invalid card type
		{"unknown card type", "4532015112830366", "unknown", true},
		{"empty card type", "4532015112830366", "", true},

		// Invalid formats
		{"invalid Luhn", "4532015112830367", "visa", true},
		{"invalid length for type", "453201511283036", "visa", true}, // 15 digits for Visa
		{"empty string", "", "visa", true},

		// Type errors
		{"non-string value", 123, "visa", true},
		{"nil value", nil, "visa", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "card_field")
			v.CreditCardType(tt.cardType)

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

func TestLuhn(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid Luhn numbers
		{"valid credit card", "4532015112830366", false},
		{"valid with spaces", "4532 0151 1283 0366", false},
		{"valid with hyphens", "4532-0151-1283-0366", false},
		{"valid short", "79927398713", false},
		{"valid long", "378282246310005", false},

		// Invalid Luhn numbers
		{"invalid last digit", "4532015112830367", true},
		{"invalid all zeros", "0000000000000000", false}, // Technically passes Luhn (sum=0, 0%10=0)
		{"invalid checksum", "1234567812345670", false}, // Actually passes Luhn!

		// Invalid formats
		{"invalid non-digits", "453201511283036a", true},
		{"invalid special chars", "4532-0151-1283-036!", true},
		{"empty string", "", true},

		// Type errors
		{"non-string value", 123456789, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "luhn_field")
			v.Luhn()

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

// Test StopOnError behavior for Credit Card validators
func TestCreditCardValidatorsStopOnError(t *testing.T) {
	t.Run("CreditCard StopOnError", func(t *testing.T) {
		v := New("invalid", "card_field").StopOnError()
		v.CreditCard()
		v.CreditCard() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("CreditCardType StopOnError", func(t *testing.T) {
		v := New("invalid", "card_field").StopOnError()
		v.CreditCardType("visa")
		v.CreditCardType("visa") // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Luhn StopOnError", func(t *testing.T) {
		v := New("invalid", "luhn_field").StopOnError()
		v.Luhn()
		v.Luhn() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})
}

// Test luhnCheck helper function directly
func TestLuhnCheck(t *testing.T) {
	tests := []struct {
		name   string
		number string
		want   bool
	}{
		{"valid Visa", "4532015112830366", true},
		{"valid Mastercard", "5425233430109903", true},
		{"valid Amex", "374245455400126", true},
		{"invalid last digit", "4532015112830367", false},
		{"all zeros passes Luhn", "0000000000000000", true}, // Technically valid by Luhn (sum=0, 0%10=0)
		{"valid short", "79927398713", true},
		{"checksum passes Luhn", "1234567812345670", true}, // This actually passes Luhn
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := luhnCheck(tt.number)
			if result != tt.want {
				t.Errorf("luhnCheck(%s) = %v, want %v", tt.number, result, tt.want)
			}
		})
	}
}

// Test credit card validation with chaining
func TestCreditCardChaining(t *testing.T) {
	t.Run("Valid credit card chain", func(t *testing.T) {
		v := New("4532015112830366", "card")
		v.Required().CreditCard().CreditCardType("visa")
		if len(v.GetErrors()) > 0 {
			t.Errorf("expected no errors but got %v", v.GetErrors())
		}
	})

	t.Run("Invalid credit card chain stops on first error", func(t *testing.T) {
		v := New("invalid", "card").StopOnError()
		v.Required().CreditCard().CreditCardType("visa")
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("Invalid credit card accumulates errors without StopOnError", func(t *testing.T) {
		v := New("invalid", "card")
		v.CreditCard().Luhn()
		errors := v.GetErrors()
		if len(errors) < 2 {
			t.Errorf("expected multiple errors without StopOnError, got %d", len(errors))
		}
	})
}
