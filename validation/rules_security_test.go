package validation

import (
	"testing"
)

// ============================================================================
// JWT VALIDATOR TESTS
// ============================================================================

func TestJWT(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid JWT tokens
		{"valid JWT", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", false},
		{"valid JWT with long payload", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U", false},

		// Invalid JWT tokens
		{"empty string", "", true},
		{"only dots", "...", true},
		{"two parts only", "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0In0", true},
		{"four parts", "part1.part2.part3.part4", true},
		{"invalid base64 in header", "invalid!.eyJzdWIiOiIxMjM0In0.signature", true},
		{"invalid base64 in payload", "eyJhbGciOiJIUzI1NiJ9.invalid!.signature", true},
		{"empty header", ".eyJzdWIiOiIxMjM0In0.signature", true},
		{"empty payload", "eyJhbGciOiJIUzI1NiJ9..signature", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "jwt_field")
			v.JWT()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

func TestJWTWithStopOnError(t *testing.T) {
	v := New("invalid", "jwt_field").StopOnError()
	v.JWT().Required()

	err := v.Validate()
	if err == nil {
		t.Error("expected error but got none")
	}

	// Should only have one error due to StopOnError
	if len(v.errors) > 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(v.errors))
	}
}

// ============================================================================
// BCRYPT VALIDATOR TESTS
// ============================================================================

func TestBCrypt(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid bcrypt hashes
		{"valid bcrypt 2a", "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", false},
		{"valid bcrypt 2b", "$2b$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", false},
		{"valid bcrypt 2x", "$2x$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", false},
		{"valid bcrypt 2y", "$2y$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", false},
		{"valid bcrypt cost 04", "$2a$04$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", false},
		{"valid bcrypt cost 12", "$2a$12$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", false},

		// Invalid bcrypt hashes
		{"empty string", "", true},
		{"wrong prefix", "$3a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", true},
		{"wrong format", "$2c$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", true},
		{"too short", "$2a$10$N9qo8uLOickgx2ZMRZoMye", true},
		{"too long", "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWyExtra", true},
		{"invalid cost single digit", "$2a$1$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", true},
		{"invalid cost three digits", "$2a$100$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", true},
		{"contains invalid char", "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh@y", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "bcrypt_field")
			v.BCrypt()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// ============================================================================
// MD5 VALIDATOR TESTS
// ============================================================================

func TestMD5(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid MD5 hashes
		{"valid MD5 lowercase", "5d41402abc4b2a76b9719d911017c592", false},
		{"valid MD5 uppercase", "5D41402ABC4B2A76B9719D911017C592", false},
		{"valid MD5 mixed case", "5d41402ABC4b2a76B9719d911017C592", false},
		{"valid MD5 all zeros", "00000000000000000000000000000000", false},
		{"valid MD5 all f", "ffffffffffffffffffffffffffffffff", false},

		// Invalid MD5 hashes
		{"empty string", "", true},
		{"too short", "5d41402abc4b2a76b9719d911017c59", true},
		{"too long", "5d41402abc4b2a76b9719d911017c5922", true},
		{"invalid character g", "5d41402abc4b2a76b9719d911017c59g", true},
		{"invalid character z", "5d41402abc4b2a76b9719d911017c59z", true},
		{"contains space", "5d41402abc4b2a76b9719d91 1017c592", true},
		{"contains dash", "5d41402a-bc4b2a76b9719d911017c592", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "md5_field")
			v.MD5()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// ============================================================================
// SHA1 VALIDATOR TESTS
// ============================================================================

func TestSHA1(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid SHA1 hashes
		{"valid SHA1 lowercase", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", false},
		{"valid SHA1 uppercase", "AAF4C61DDCC5E8A2DABEDE0F3B482CD9AEA9434D", false},
		{"valid SHA1 mixed case", "aaf4C61ddCC5e8a2DABEDE0f3b482CD9aea9434D", false},
		{"valid SHA1 all zeros", "0000000000000000000000000000000000000000", false},
		{"valid SHA1 all f", "ffffffffffffffffffffffffffffffffffffffff", false},

		// Invalid SHA1 hashes
		{"empty string", "", true},
		{"too short", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434", true},
		{"too long", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434dd", true},
		{"invalid character g", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434g", true},
		{"invalid character z", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434z", true},
		{"contains space", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea 434d", true},
		{"contains dash", "aaf4c61d-dcc5e8a2dabede0f3b482cd9aea9434d", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "sha1_field")
			v.SHA1()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// ============================================================================
// SHA256 VALIDATOR TESTS
// ============================================================================

func TestSHA256(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid SHA256 hashes
		{"valid SHA256 lowercase", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", false},
		{"valid SHA256 uppercase", "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855", false},
		{"valid SHA256 mixed case", "e3b0C44298FC1c149afBF4c8996fb92427AE41e4649b934CA495991b7852B855", false},
		{"valid SHA256 all zeros", "0000000000000000000000000000000000000000000000000000000000000000", false},
		{"valid SHA256 all f", "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", false},

		// Invalid SHA256 hashes
		{"empty string", "", true},
		{"too short", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85", true},
		{"too long", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b8555", true},
		{"invalid character g", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85g", true},
		{"invalid character z", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85z", true},
		{"contains space", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b 7852b855", true},
		{"contains dash", "e3b0c442-98fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "sha256_field")
			v.SHA256()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// ============================================================================
// SHA512 VALIDATOR TESTS
// ============================================================================

func TestSHA512(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid SHA512 hashes
		{"valid SHA512 lowercase", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", false},
		{"valid SHA512 uppercase", "CF83E1357EEFB8BDF1542850D66D8007D620E4050B5715DC83F4A921D36CE9CE47D0D13C5D85F2B0FF8318D2877EEC2F63B931BD47417A81A538327AF927DA3E", false},
		{"valid SHA512 mixed case", "cf83E1357eefB8bdf1542850D66d8007d620E4050b5715DC83f4a921D36ce9ce47D0d13c5D85f2b0ff8318D2877eec2F63b931BD47417a81A538327af927DA3e", false},
		{"valid SHA512 all zeros", "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", false},
		{"valid SHA512 all f", "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", false},

		// Invalid SHA512 hashes
		{"empty string", "", true},
		{"too short", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3", true},
		{"too long", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3ee", true},
		{"invalid character g", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3g", true},
		{"invalid character z", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3z", true},
		{"contains space", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927 a3e", true},
		{"contains dash", "cf83e135-7eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "sha512_field")
			v.SHA512()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// ============================================================================
// MULTI-FIELD SECURITY VALIDATION TESTS
// ============================================================================

func TestMultiFieldSecurityValidation(t *testing.T) {
	type SecurityData struct {
		JWTToken      string
		BCryptHash    string
		MD5Hash       string
		SHA256Hash    string
	}

	data := SecurityData{
		JWTToken:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U",
		BCryptHash: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
		MD5Hash:    "5d41402abc4b2a76b9719d911017c592",
		SHA256Hash: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	}

	mv := NewValidator()
	mv.Field(data.JWTToken, "jwt_token").JWT()
	mv.Field(data.BCryptHash, "bcrypt_hash").BCrypt()
	mv.Field(data.MD5Hash, "md5_hash").MD5()
	mv.Field(data.SHA256Hash, "sha256_hash").SHA256()

	err := mv.Validate()
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}

func TestSecurityValidationChaining(t *testing.T) {
	// Test chaining with other validators
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U"

	v := New(jwt, "token")
	v.Required().JWT().MinLength(10)

	err := v.Validate()
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}
