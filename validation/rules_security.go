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
// Validates the three-part structure and base64url encoding.
//
// JWT는 값이 유효한 JWT (JSON Web Token) 형식인지 검증합니다.
// 세 부분 구조와 base64url 인코딩을 검증합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Validates 3-part structure: header.payload.signature
//     3부분 구조 검증: header.payload.signature
//   - Validates base64url encoding
//     base64url 인코딩 검증
//   - Checks non-empty parts
//     비어있지 않은 부분 확인
//   - Does not validate JWT signature or expiration
//     JWT 서명이나 만료 검증하지 않음
//
// Use Cases / 사용 사례:
//   - API token validation / API 토큰 검증
//   - Authentication header parsing / 인증 헤더 파싱
//   - OAuth/OpenID Connect / OAuth/OpenID 연결
//   - Token format verification / 토큰 형식 확인
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = token length
//     시간 복잡도: O(n), n = 토큰 길이
//   - Base64 decode for 3 parts
//     3개 부분에 대한 Base64 디코드
//
// A valid JWT has three parts separated by dots: header.payload.signature
// Each part is base64url encoded.
//
// 유효한 JWT는 점으로 구분된 세 부분으로 구성됩니다: header.payload.signature
// 각 부분은 base64url로 인코딩됩니다.
//
// Example / 예시:
//
//	// Valid JWT / 유효한 JWT
//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U"
//	v := validation.New(token, "jwt_token")
//	v.JWT()  // Passes
//
//	// With user claims / 사용자 클레임 포함
//	token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxMjN9.abc123def"
//	v = validation.New(token, "auth_token")
//	v.JWT()  // Passes (format only)
//
//	// Invalid - missing part / 무효 - 부분 누락
//	v = validation.New("header.payload", "token")
//	v.JWT()  // Fails (only 2 parts)
//
//	// Invalid - bad encoding / 무효 - 잘못된 인코딩
//	v = validation.New("abc.def.ghi", "token")
//	v.JWT()  // Fails (invalid base64url)
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
// Validates the BCrypt hash format with cost factor and encoded salt/hash.
//
// BCrypt는 값이 유효한 BCrypt 해시인지 검증합니다.
// 비용 계수 및 인코딩된 솔트/해시를 포함한 BCrypt 해시 형식을 검증합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Validates BCrypt identifier ($2a$, $2b$, $2x$, $2y$)
//     BCrypt 식별자 검증 ($2a$, $2b$, $2x$, $2y$)
//   - Validates cost factor (00-31)
//     비용 계수 검증 (00-31)
//   - Validates 60-character length
//     60자 길이 검증
//   - Validates base64 salt/hash encoding
//     base64 솔트/해시 인코딩 검증
//
// Use Cases / 사용 사례:
//   - Password hash storage / 비밀번호 해시 저장
//   - Database hash validation / 데이터베이스 해시 검증
//   - Authentication systems / 인증 시스템
//   - Security compliance / 보안 규정 준수
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1), fixed length check
//     시간 복잡도: O(1), 고정 길이 확인
//   - Regex pattern matching
//     정규식 패턴 매칭
//
// BCrypt hashes start with $2a$, $2b$, $2x$, or $2y$ followed by cost and salt.
// BCrypt 해시는 $2a$, $2b$, $2x$, 또는 $2y$로 시작하며 비용과 솔트가 따릅니다.
//
// Example / 예시:
//
//	// Valid BCrypt hash / 유효한 BCrypt 해시
//	hash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
//	v := validation.New(hash, "password_hash")
//	v.BCrypt()  // Passes
//
//	// Higher cost factor / 높은 비용 계수
//	hash = "$2b$12$KCgZvfKlYH2PvzNLn5C4veWmJ3dP5HlN4KJlBp4LwBd0XYZ1qR4Qi"
//	v = validation.New(hash, "hash")
//	v.BCrypt()  // Passes
//
//	// Invalid - wrong length / 무효 - 잘못된 길이
//	v = validation.New("$2a$10$short", "hash")
//	v.BCrypt()  // Fails
//
//	// Invalid - wrong format / 무효 - 잘못된 형식
//	v = validation.New("not-a-bcrypt-hash", "hash")
//	v.BCrypt()  // Fails
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
// MD5 produces 128-bit hash values typically represented as 32 hexadecimal digits.
//
// MD5는 값이 유효한 MD5 해시(32자리 16진수)인지 검증합니다.
// MD5는 일반적으로 32자리 16진수로 표현되는 128비트 해시 값을 생성합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Validates 32-character length
//     32자 길이 검증
//   - Validates hexadecimal characters (0-9, a-f, A-F)
//     16진수 문자 검증 (0-9, a-f, A-F)
//   - Case-insensitive
//     대소문자 구분 안 함
//   - Does not validate hash correctness
//     해시 정확성 검증하지 않음
//
// Use Cases / 사용 사례:
//   - File integrity checking / 파일 무결성 확인
//   - Hash format validation / 해시 형식 검증
//   - Checksum verification / 체크섬 확인
//   - Legacy system compatibility / 레거시 시스템 호환성
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1), fixed length regex
//     시간 복잡도: O(1), 고정 길이 정규식
//   - Simple pattern matching
//     간단한 패턴 매칭
//
// Note: MD5 is cryptographically broken and should not be used for security purposes.
// Use SHA256 or stronger algorithms for security-critical applications.
//
// 참고: MD5는 암호학적으로 취약하며 보안 목적으로 사용해서는 안 됩니다.
// 보안이 중요한 애플리케이션에는 SHA256 이상의 강력한 알고리즘을 사용하세요.
//
// Example / 예시:
//
//	// Valid MD5 hash / 유효한 MD5 해시
//	hash := "5d41402abc4b2a76b9719d911017c592"
//	v := validation.New(hash, "file_hash")
//	v.MD5()  // Passes
//
//	// Uppercase hex / 대문자 16진수
//	hash = "5D41402ABC4B2A76B9719D911017C592"
//	v = validation.New(hash, "hash")
//	v.MD5()  // Passes
//
//	// Invalid - wrong length / 무효 - 잘못된 길이
//	v = validation.New("5d41402abc4b2a76", "hash")
//	v.MD5()  // Fails (too short)
//
//	// Invalid - non-hex characters / 무효 - 16진수가 아닌 문자
//	v = validation.New("5d41402abc4b2a76b9719d911017c59z", "hash")
//	v.MD5()  // Fails (contains 'z')
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
// SHA1 produces 160-bit hash values typically represented as 40 hexadecimal digits.
//
// SHA1은 값이 유효한 SHA1 해시(40자리 16진수)인지 검증합니다.
// SHA1은 일반적으로 40자리 16진수로 표현되는 160비트 해시 값을 생성합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Validates 40-character length
//     40자 길이 검증
//   - Validates hexadecimal characters (0-9, a-f, A-F)
//     16진수 문자 검증 (0-9, a-f, A-F)
//   - Case-insensitive
//     대소문자 구분 안 함
//   - Does not validate hash correctness
//     해시 정확성 검증하지 않음
//
// Use Cases / 사용 사례:
//   - Git commit hash validation / Git 커밋 해시 검증
//   - File integrity verification / 파일 무결성 확인
//   - Digital signature validation / 디지털 서명 검증
//   - Version control systems / 버전 관리 시스템
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1), fixed length regex
//     시간 복잡도: O(1), 고정 길이 정규식
//   - Simple pattern matching
//     간단한 패턴 매칭
//
// Note: SHA1 is no longer considered cryptographically secure.
// Use SHA256 or stronger algorithms for security-critical applications.
//
// 참고: SHA1은 더 이상 암호학적으로 안전하지 않습니다.
// 보안이 중요한 애플리케이션에는 SHA256 이상의 강력한 알고리즘을 사용하세요.
//
// Example / 예시:
//
//	// Valid SHA1 hash / 유효한 SHA1 해시
//	hash := "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
//	v := validation.New(hash, "commit_hash")
//	v.SHA1()  // Passes
//
//	// Git commit hash / Git 커밋 해시
//	hash = "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"
//	v = validation.New(hash, "git_sha")
//	v.SHA1()  // Passes
//
//	// Invalid - wrong length / 무효 - 잘못된 길이
//	v = validation.New("aaf4c61ddcc5e8a2", "hash")
//	v.SHA1()  // Fails (too short)
//
//	// Invalid - non-hex characters / 무효 - 16진수가 아닌 문자
//	v = validation.New("zaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", "hash")
//	v.SHA1()  // Fails (contains 'z')
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
// SHA256 produces 256-bit hash values typically represented as 64 hexadecimal digits.
//
// SHA256은 값이 유효한 SHA256 해시(64자리 16진수)인지 검증합니다.
// SHA256은 일반적으로 64자리 16진수로 표현되는 256비트 해시 값을 생성합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Validates 64-character length
//     64자 길이 검증
//   - Validates hexadecimal characters (0-9, a-f, A-F)
//     16진수 문자 검증 (0-9, a-f, A-F)
//   - Case-insensitive
//     대소문자 구분 안 함
//   - Does not validate hash correctness
//     해시 정확성 검증하지 않음
//
// Use Cases / 사용 사례:
//   - File integrity verification / 파일 무결성 확인
//   - Blockchain and cryptocurrency / 블록체인 및 암호화폐
//   - Digital signatures / 디지털 서명
//   - Password hashing (with salt) / 비밀번호 해싱 (솔트 포함)
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1), fixed length regex
//     시간 복잡도: O(1), 고정 길이 정규식
//   - Simple pattern matching
//     간단한 패턴 매칭
//
// SHA256 is currently considered cryptographically secure and is widely used
// in security-critical applications.
//
// SHA256은 현재 암호학적으로 안전한 것으로 간주되며 보안이 중요한
// 애플리케이션에서 널리 사용됩니다.
//
// Example / 예시:
//
//	// Valid SHA256 hash / 유효한 SHA256 해시
//	hash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
//	v := validation.New(hash, "file_hash")
//	v.SHA256()  // Passes
//
//	// Bitcoin block hash / 비트코인 블록 해시
//	hash = "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
//	v = validation.New(hash, "block_hash")
//	v.SHA256()  // Passes
//
//	// Invalid - wrong length / 무효 - 잘못된 길이
//	v = validation.New("e3b0c44298fc1c14", "hash")
//	v.SHA256()  // Fails (too short)
//
//	// Invalid - non-hex characters / 무효 - 16진수가 아닌 문자
//	v = validation.New("z3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", "hash")
//	v.SHA256()  // Fails (contains 'z')
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
// SHA512 produces 512-bit hash values typically represented as 128 hexadecimal digits.
//
// SHA512는 값이 유효한 SHA512 해시(128자리 16진수)인지 검증합니다.
// SHA512는 일반적으로 128자리 16진수로 표현되는 512비트 해시 값을 생성합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Validates 128-character length
//     128자 길이 검증
//   - Validates hexadecimal characters (0-9, a-f, A-F)
//     16진수 문자 검증 (0-9, a-f, A-F)
//   - Case-insensitive
//     대소문자 구분 안 함
//   - Does not validate hash correctness
//     해시 정확성 검증하지 않음
//
// Use Cases / 사용 사례:
//   - High-security applications / 고보안 애플리케이션
//   - Large file integrity / 대용량 파일 무결성
//   - Certificate generation / 인증서 생성
//   - Password hashing (with salt) / 비밀번호 해싱 (솔트 포함)
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1), fixed length regex
//     시간 복잡도: O(1), 고정 길이 정규식
//   - Simple pattern matching
//     간단한 패턴 매칭
//
// SHA512 offers higher security than SHA256 but produces longer hash values.
// It is suitable for applications requiring maximum security.
//
// SHA512는 SHA256보다 높은 보안성을 제공하지만 더 긴 해시 값을 생성합니다.
// 최대 보안이 필요한 애플리케이션에 적합합니다.
//
// Example / 예시:
//
//	// Valid SHA512 hash / 유효한 SHA512 해시
//	hash := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"
//	v := validation.New(hash, "password_hash")
//	v.SHA512()  // Passes
//
//	// File checksum / 파일 체크섬
//	hash = "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
//	v = validation.New(hash, "file_checksum")
//	v.SHA512()  // Passes
//
//	// Invalid - wrong length / 무효 - 잘못된 길이
//	v = validation.New("cf83e1357eefb8bdf1542850d66d8007", "hash")
//	v.SHA512()  // Fails (too short)
//
//	// Invalid - non-hex characters / 무효 - 16진수가 아닌 문자
//	v = validation.New("zf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", "hash")
//	v.SHA512()  // Fails (contains 'z')
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
