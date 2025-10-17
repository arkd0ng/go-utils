package validation

import (
	"os"
	"testing"
	"time"
)

// TestFileWritable_Directory tests FileWritable with a directory path
func TestFileWritable_Directory(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "test_dir_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	v := New(tmpDir, "path")
	v.FileWritable()

	// Should fail because it's a directory, not a file
	if len(v.GetErrors()) == 0 {
		t.Error("expected error for directory path")
	}
}

// TestFileWritable_ReadOnlyFile tests FileWritable with a read-only file
func TestFileWritable_ReadOnlyFile(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test_readonly_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// Make it read-only
	if err := os.Chmod(tmpFile.Name(), 0444); err != nil {
		t.Fatal(err)
	}

	v := New(tmpFile.Name(), "file")
	v.FileWritable()

	// Restore permissions for cleanup
	os.Chmod(tmpFile.Name(), 0644)

	// Should fail because file is read-only
	if len(v.GetErrors()) == 0 {
		t.Error("expected error for read-only file")
	}
}

// TestFileWritable_InvalidParentDir tests FileWritable with invalid parent directory
func TestFileWritable_InvalidParentDir(t *testing.T) {
	v := New("/nonexistent/path/file.txt", "file")
	v.FileWritable()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for non-existent parent directory")
	}
}

// TestFileWritable_WritableFile tests FileWritable with an existing writable file (success case)
func TestFileWritable_WritableFile(t *testing.T) {
	// Create a temporary writable file
	tmpFile, err := os.CreateTemp("", "test_writable_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	v := New(tmpFile.Name(), "file")
	v.FileWritable()

	// Should pass because file is writable
	if len(v.GetErrors()) != 0 {
		t.Errorf("expected no error for writable file, got %v", v.GetErrors())
	}
}

// TestFileWritable_NewFileInWritableDir tests FileWritable with new file in writable directory (success case)
func TestFileWritable_NewFileInWritableDir(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "test_dir_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Test a new file in this directory
	newFile := tmpDir + "/newfile.txt"
	v := New(newFile, "file")
	v.FileWritable()

	// Should pass because parent directory is writable
	if len(v.GetErrors()) != 0 {
		t.Errorf("expected no error for new file in writable directory, got %v", v.GetErrors())
	}
}

// TestFileWritable_WithStopOnError tests FileWritable with stopOnError flag
func TestFileWritable_WithStopOnError(t *testing.T) {
	v := New("/some/path/file.txt", "file")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.FileWritable()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestEmpty_WithStopOnError tests Empty with stopOnError flag
func TestEmpty_WithStopOnError(t *testing.T) {
	v := New("not empty", "field")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.Empty()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestNotEmpty_WithStopOnError tests NotEmpty with stopOnError flag
func TestNotEmpty_WithStopOnError(t *testing.T) {
	v := New("", "field")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.NotEmpty()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestEmpty_WithNonEmptyValue tests Empty with non-empty value
func TestEmpty_WithNonEmptyValue(t *testing.T) {
	v := New("not empty", "field")
	v.Empty()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for non-empty string")
	}
}

// TestNotEmpty_WithEmptyValue tests NotEmpty with empty value
func TestNotEmpty_WithEmptyValue(t *testing.T) {
	v := New("", "field")
	v.NotEmpty()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for empty string")
	}
}

// TestIsEmptyValue_AllTypes tests isEmptyValue with all supported types
func TestIsEmptyValue_AllTypes(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		// Nil
		{"nil", nil, true},

		// String
		{"empty string", "", true},
		{"non-empty string", "hello", false},

		// Bool
		{"false bool", false, true},
		{"true bool", true, false},

		// Int types
		{"zero int", int(0), true},
		{"non-zero int", int(42), false},
		{"zero int8", int8(0), true},
		{"non-zero int8", int8(42), false},
		{"zero int16", int16(0), true},
		{"non-zero int16", int16(42), false},
		{"zero int32", int32(0), true},
		{"non-zero int32", int32(42), false},
		{"zero int64", int64(0), true},
		{"non-zero int64", int64(42), false},

		// Uint types
		{"zero uint", uint(0), true},
		{"non-zero uint", uint(42), false},
		{"zero uint8", uint8(0), true},
		{"non-zero uint8", uint8(42), false},
		{"zero uint16", uint16(0), true},
		{"non-zero uint16", uint16(42), false},
		{"zero uint32", uint32(0), true},
		{"non-zero uint32", uint32(42), false},
		{"zero uint64", uint64(0), true},
		{"non-zero uint64", uint64(42), false},
		{"zero uintptr", uintptr(0), true},
		{"non-zero uintptr", uintptr(42), false},

		// Float types
		{"zero float32", float32(0.0), true},
		{"non-zero float32", float32(3.14), false},
		{"zero float64", float64(0.0), true},
		{"non-zero float64", float64(3.14), false},

		// Pointer
		{"nil pointer", (*int)(nil), true},
		{"non-nil pointer", new(int), false},

		// Interface
		{"nil interface", interface{}(nil), true},

		// Slice
		{"nil slice", []int(nil), true},
		{"empty slice", []int{}, true},
		{"non-empty slice", []int{1, 2}, false},

		// Map
		{"nil map", map[string]int(nil), true},
		{"empty map", map[string]int{}, true},
		{"non-empty map", map[string]int{"a": 1}, false},

		// Channel (unbuffered channel has Len() == 0, so it's considered empty)
		{"nil chan", (chan int)(nil), true},
		{"non-nil unbuffered chan", make(chan int), true},
		{"non-nil buffered chan with data", func() chan int { ch := make(chan int, 1); ch <- 1; return ch }(), false},

		// Array
		{"empty array", [0]int{}, true},
		{"non-empty array", [2]int{1, 2}, false},

		// Struct (unsupported type, should return false)
		{"struct", struct{}{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isEmptyValue(tt.value)
			if result != tt.expected {
				t.Errorf("isEmptyValue(%v) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

// TestBetweenTime_WithStopOnError tests BetweenTime with stopOnError flag
func TestBetweenTime_WithStopOnError(t *testing.T) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now.Add(1 * time.Hour)

	v := New(now, "datetime")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.BetweenTime(start, end)

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestBetweenTime_WithNonTimeValue tests BetweenTime with non-time value
func TestBetweenTime_WithNonTimeValue(t *testing.T) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now.Add(1 * time.Hour)

	v := New("not a time", "datetime")
	v.BetweenTime(start, end)

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for non-time value")
	}
}

// TestBetweenTime_BeforeRange tests BetweenTime with value before range
func TestBetweenTime_BeforeRange(t *testing.T) {
	now := time.Now()
	start := now.Add(1 * time.Hour)
	end := now.Add(2 * time.Hour)

	v := New(now, "datetime")
	v.BetweenTime(start, end)

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for time before range")
	}
}

// TestBetweenTime_AfterRange tests BetweenTime with value after range
func TestBetweenTime_AfterRange(t *testing.T) {
	now := time.Now()
	start := now.Add(-2 * time.Hour)
	end := now.Add(-1 * time.Hour)

	v := New(now, "datetime")
	v.BetweenTime(start, end)

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for time after range")
	}
}

// TestHexColor_WithStopOnError tests HexColor with stopOnError flag
func TestHexColor_WithStopOnError(t *testing.T) {
	v := New("#FFF", "color")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.HexColor()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestRGBA_WithStopOnError tests RGBA with stopOnError flag
func TestRGBA_WithStopOnError(t *testing.T) {
	v := New("rgba(255,0,0,1)", "color")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.RGBA()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestRGBA_InvalidRGBValue tests RGBA with invalid RGB value (>255)
func TestRGBA_InvalidRGBValue(t *testing.T) {
	v := New("rgba(256,0,0,1)", "color")
	v.RGBA()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for RGB value > 255")
	}
}

// TestHSL_WithStopOnError tests HSL with stopOnError flag
func TestHSL_WithStopOnError(t *testing.T) {
	v := New("hsl(120,100%,50%)", "color")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.HSL()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestHSL_InvalidHueValue tests HSL with invalid hue value (>360)
func TestHSL_InvalidHueValue(t *testing.T) {
	v := New("hsl(361,100%,50%)", "color")
	v.HSL()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for hue value > 360")
	}
}

// TestEAN_InvalidChecksum tests EAN with invalid checksum (EAN8 format)
func TestEAN_InvalidChecksum8(t *testing.T) {
	// Valid format but wrong checksum
	// "12345678" has wrong check digit (should be calculated from first 7 digits)
	v := New("12345678", "ean")
	v.EAN()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for invalid EAN8 checksum")
	}
}

// TestEAN_InvalidChecksum tests EAN with invalid checksum (EAN13 format)
func TestEAN_InvalidChecksum13(t *testing.T) {
	// Valid format but wrong checksum
	v := New("1234567890120", "ean")
	v.EAN()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for invalid EAN13 checksum")
	}
}

// TestEAN_ValidNumbers tests EAN with valid EAN8 and EAN13 numbers
func TestEAN_ValidNumbers(t *testing.T) {
	validEANs := []string{
		"40123455",       // Valid EAN-8
		"5901234123457",  // Valid EAN-13
		"4006381333931",  // Valid EAN-13 (Coca-Cola)
		"0000000000000",  // Valid EAN-13 (remainder == 0 case)
	}

	for _, ean := range validEANs {
		v := New(ean, "ean")
		v.EAN()

		if len(v.GetErrors()) != 0 {
			t.Errorf("expected no error for valid EAN %s, got %v", ean, v.GetErrors())
		}
	}
}

// TestLuhnCheck_InvalidChecksum tests credit card Luhn check with invalid checksum
func TestLuhnCheck_InvalidChecksum(t *testing.T) {
	// Valid format but fails Luhn check
	v := New("4111111111111112", "card")
	v.CreditCard()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for invalid Luhn checksum")
	}
}

// TestCreditCard_WithNonNumeric tests CreditCard with non-numeric characters (letters)
func TestCreditCard_WithNonNumeric(t *testing.T) {
	v := New("4111ABCD11111111", "card")
	v.CreditCard()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for card number with letters")
	}
}

// TestCreditCard_ValidNumbers tests CreditCard with valid card numbers to cover all Luhn paths
func TestCreditCard_ValidNumbers(t *testing.T) {
	validCards := []string{
		"4111111111111111", // Visa
		"5500000000000004", // Mastercard
		"340000000000009",  // Amex
		"6011000000000004", // Discover
		"378282246310005",  // Amex (more digits > 5 for doubling > 9 case)
		"5555555555554444", // Mastercard (many 5s for d > 9 coverage)
	}

	for _, card := range validCards {
		v := New(card, "card")
		v.CreditCard()

		if len(v.GetErrors()) != 0 {
			t.Errorf("expected no error for valid card %s, got %v", card, v.GetErrors())
		}
	}
}

// TestJWT_WithStopOnError tests JWT with stopOnError flag
func TestJWT_WithStopOnError(t *testing.T) {
	v := New("valid.jwt.token", "token")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.JWT()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestBCrypt_WithStopOnError tests BCrypt with stopOnError flag
func TestBCrypt_WithStopOnError(t *testing.T) {
	v := New("$2a$10$valid", "hash")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.BCrypt()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestMD5_WithStopOnError tests MD5 with stopOnError flag
func TestMD5_WithStopOnError(t *testing.T) {
	v := New("5d41402abc4b2a76b9719d911017c592", "hash")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.MD5()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestSHA1_WithStopOnError tests SHA1 with stopOnError flag
func TestSHA1_WithStopOnError(t *testing.T) {
	v := New("aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", "hash")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.SHA1()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestSHA256_WithStopOnError tests SHA256 with stopOnError flag
func TestSHA256_WithStopOnError(t *testing.T) {
	v := New("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", "hash")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.SHA256()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestSHA512_WithStopOnError tests SHA512 with stopOnError flag
func TestSHA512_WithStopOnError(t *testing.T) {
	v := New("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", "hash")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.SHA512()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestFalse_WithStopOnError tests False with stopOnError flag
func TestFalse_WithStopOnError(t *testing.T) {
	v := New(false, "flag")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.False()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}

// TestFalse_WithNonBoolValue tests False with non-bool value
func TestFalse_WithNonBoolValue(t *testing.T) {
	v := New("not a bool", "flag")
	v.False()

	if len(v.GetErrors()) == 0 {
		t.Error("expected error for non-bool value")
	}
}

// TestNotNil_WithStopOnError tests NotNil with stopOnError flag
func TestNotNil_WithStopOnError(t *testing.T) {
	v := New("not nil", "value")
	v.stopOnError = true
	v.addError("previous", "previous error")

	// Should skip validation due to stopOnError
	v.NotNil()

	if len(v.GetErrors()) != 1 {
		t.Errorf("expected 1 error (previous), got %d", len(v.GetErrors()))
	}
}
