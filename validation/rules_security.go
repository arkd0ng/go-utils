package validation

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

// ============================================================================
// SECURITY VALIDATORS
// ============================================================================
//
// This file provides security-related validation functions for:
// - JWT (JSON Web Token) format validation
// - BCrypt hash validation
// - Hash format validation (MD5, SHA1, SHA256, SHA512)
//
// 이 파일은 다음을 위한 보안 관련 검증 함수를 제공합니다:
// - JWT (JSON Web Token) 형식 검증
// - BCrypt 해시 검증
// - 해시 형식 검증 (MD5, SHA1, SHA256, SHA512)
//
// ============================================================================

// JWT validates that the value is a valid JWT (JSON Web Token) format.
// JWT는 값이 유효한 JWT (JSON Web Token) 형식인지 검증합니다.
//
// A valid JWT has three parts separated by dots: header.payload.signature
// Each part is base64url encoded.
//
// 유효한 JWT는 점으로 구분된 세 부분으로 구성됩니다: header.payload.signature
// 각 부분은 base64url로 인코딩됩니다.
//
// Example / 예시:
//
//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U"
//	v := validation.New(token, "jwt_token")
//	v.JWT()
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Must have exactly 3 parts separated by dots / 점으로 구분된 정확히 3개 부분
//   - Each part must be valid base64url encoding / 각 부분은 유효한 base64url 인코딩
//   - Header and payload must be non-empty / 헤더와 페이로드는 비어있지 않아야 함
func (v *Validator) JWT() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("jwt", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// JWT format: header.payload.signature
	parts := strings.Split(str, ".")
	if len(parts) != 3 {
		v.addError("jwt", fmt.Sprintf("%s must be a valid JWT with 3 parts (header.payload.signature) / %s은(는) 3개 부분으로 구성된 유효한 JWT여야 합니다 (header.payload.signature)", v.fieldName, v.fieldName))
		return v
	}

	// Validate each part is non-empty and base64url encoded
	for i, part := range parts {
		if part == "" {
			partName := []string{"header", "payload", "signature"}[i]
			v.addError("jwt", fmt.Sprintf("%s JWT %s cannot be empty / %s JWT %s은(는) 비어있을 수 없습니다", v.fieldName, partName, v.fieldName, partName))
			return v
		}

		// Base64url uses - and _ instead of + and /, and no padding
		// Convert to standard base64 for validation
		base64Str := strings.ReplaceAll(strings.ReplaceAll(part, "-", "+"), "_", "/")

		// Add padding if needed
		switch len(base64Str) % 4 {
		case 2:
			base64Str += "=="
		case 3:
			base64Str += "="
		}

		// Try to decode
		_, err := base64.StdEncoding.DecodeString(base64Str)
		if err != nil {
			partName := []string{"header", "payload", "signature"}[i]
			v.addError("jwt", fmt.Sprintf("%s JWT %s must be valid base64url encoding / %s JWT %s은(는) 유효한 base64url 인코딩이어야 합니다", v.fieldName, partName, v.fieldName, partName))
			return v
		}
	}

	return v
}

// BCrypt validates that the value is a valid BCrypt hash.
// BCrypt는 값이 유효한 BCrypt 해시인지 검증합니다.
//
// BCrypt hashes start with $2a$, $2b$, $2x$, or $2y$ followed by cost and salt.
// BCrypt 해시는 $2a$, $2b$, $2x$, 또는 $2y$로 시작하며 비용과 솔트가 따릅니다.
//
// Example / 예시:
//
//	hash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
//	v := validation.New(hash, "password_hash")
//	v.BCrypt()
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Must start with $2a$, $2b$, $2x$, or $2y$ / $2a$, $2b$, $2x$, 또는 $2y$로 시작
//   - Must be 60 characters long / 60자여야 함
//   - Must match BCrypt format / BCrypt 형식과 일치해야 함
func (v *Validator) BCrypt() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("bcrypt", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// BCrypt hash format: $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy
	// Length should be 60 characters
	// Pattern: $2[abxy]$[0-9]{2}$[./A-Za-z0-9]{53}
	bcryptRegex := regexp.MustCompile(`^\$2[abxy]\$[0-9]{2}\$[./A-Za-z0-9]{53}$`)
	if !bcryptRegex.MatchString(str) {
		v.addError("bcrypt", fmt.Sprintf("%s must be a valid BCrypt hash / %s은(는) 유효한 BCrypt 해시여야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// MD5 validates that the value is a valid MD5 hash (32 hexadecimal characters).
// MD5는 값이 유효한 MD5 해시(32자리 16진수)인지 검증합니다.
//
// Example / 예시:
//
//	hash := "5d41402abc4b2a76b9719d911017c592"
//	v := validation.New(hash, "password_hash")
//	v.MD5()
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Must be exactly 32 characters / 정확히 32자여야 함
//   - Must contain only hexadecimal characters (0-9, a-f, A-F) / 16진수 문자만 포함
func (v *Validator) MD5() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("md5", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	md5Regex := regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
	if !md5Regex.MatchString(str) {
		v.addError("md5", fmt.Sprintf("%s must be a valid MD5 hash (32 hexadecimal characters) / %s은(는) 유효한 MD5 해시여야 합니다 (32자리 16진수)", v.fieldName, v.fieldName))
	}

	return v
}

// SHA1 validates that the value is a valid SHA1 hash (40 hexadecimal characters).
// SHA1은 값이 유효한 SHA1 해시(40자리 16진수)인지 검증합니다.
//
// Example / 예시:
//
//	hash := "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
//	v := validation.New(hash, "commit_hash")
//	v.SHA1()
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Must be exactly 40 characters / 정확히 40자여야 함
//   - Must contain only hexadecimal characters (0-9, a-f, A-F) / 16진수 문자만 포함
func (v *Validator) SHA1() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("sha1", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	sha1Regex := regexp.MustCompile(`^[a-fA-F0-9]{40}$`)
	if !sha1Regex.MatchString(str) {
		v.addError("sha1", fmt.Sprintf("%s must be a valid SHA1 hash (40 hexadecimal characters) / %s은(는) 유효한 SHA1 해시여야 합니다 (40자리 16진수)", v.fieldName, v.fieldName))
	}

	return v
}

// SHA256 validates that the value is a valid SHA256 hash (64 hexadecimal characters).
// SHA256은 값이 유효한 SHA256 해시(64자리 16진수)인지 검증합니다.
//
// Example / 예시:
//
//	hash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
//	v := validation.New(hash, "file_hash")
//	v.SHA256()
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Must be exactly 64 characters / 정확히 64자여야 함
//   - Must contain only hexadecimal characters (0-9, a-f, A-F) / 16진수 문자만 포함
func (v *Validator) SHA256() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("sha256", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	sha256Regex := regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
	if !sha256Regex.MatchString(str) {
		v.addError("sha256", fmt.Sprintf("%s must be a valid SHA256 hash (64 hexadecimal characters) / %s은(는) 유효한 SHA256 해시여야 합니다 (64자리 16진수)", v.fieldName, v.fieldName))
	}

	return v
}

// SHA512 validates that the value is a valid SHA512 hash (128 hexadecimal characters).
// SHA512는 값이 유효한 SHA512 해시(128자리 16진수)인지 검증합니다.
//
// Example / 예시:
//
//	hash := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"
//	v := validation.New(hash, "password_hash")
//	v.SHA512()
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Must be exactly 128 characters / 정확히 128자여야 함
//   - Must contain only hexadecimal characters (0-9, a-f, A-F) / 16진수 문자만 포함
func (v *Validator) SHA512() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("sha512", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	sha512Regex := regexp.MustCompile(`^[a-fA-F0-9]{128}$`)
	if !sha512Regex.MatchString(str) {
		v.addError("sha512", fmt.Sprintf("%s must be a valid SHA512 hash (128 hexadecimal characters) / %s은(는) 유효한 SHA512 해시여야 합니다 (128자리 16진수)", v.fieldName, v.fieldName))
	}

	return v
}
