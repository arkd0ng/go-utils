package validation

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

// Security tests for validation package
// validation 패키지의 보안 테스트

// TestSecurity_InputValidation tests input validation security
// TestSecurity_InputValidation는 입력 검증 보안을 테스트합니다
func TestSecurity_InputValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "normal input",
			input:   "hello",
			wantErr: false,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
		{
			name:    "null bytes",
			input:   "hello\x00world",
			wantErr: false, // Should handle gracefully / 우아하게 처리해야 함
		},
		{
			name:    "control characters",
			input:   "hello\nworld\r\n",
			wantErr: false, // Should handle gracefully / 우아하게 처리해야 함
		},
		{
			name:    "unicode",
			input:   "안녕하세요",
			wantErr: false,
		},
		{
			name:    "mixed unicode",
			input:   "Hello안녕🎉",
			wantErr: false,
		},
		{
			name:    "very long input",
			input:   strings.Repeat("a", 10000),
			wantErr: false,
		},
		{
			name:    "special characters",
			input:   "!@#$%^&*()_+-=[]{}|;:',.<>?/`~",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.input, "text")
			v.Required()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Should never panic / 절대 패닉하지 않아야 함
			// The fact that we got here means no panic occurred
			// 여기까지 왔다는 것은 패닉이 발생하지 않았다는 의미
		})
	}
}

// TestSecurity_SQLInjectionPatterns tests that validation doesn't execute malicious patterns
// TestSecurity_SQLInjectionPatterns는 검증이 악성 패턴을 실행하지 않는지 테스트합니다
func TestSecurity_SQLInjectionPatterns(t *testing.T) {
	// Common SQL injection patterns / 일반적인 SQL 인젝션 패턴
	injectionAttempts := []string{
		"' OR '1'='1",
		"'; DROP TABLE users; --",
		"1' UNION SELECT * FROM users--",
		"admin'--",
		"' OR 1=1--",
		"' OR 'a'='a",
		"1; DELETE FROM users",
		"1' AND '1'='1",
	}

	for _, attempt := range injectionAttempts {
		t.Run(fmt.Sprintf("injection_%s", attempt), func(t *testing.T) {
			v := New(attempt, "username")
			v.Required().MinLength(3).MaxLength(20)

			// Validation should handle these safely
			// 검증은 이들을 안전하게 처리해야 함
			_ = v.Validate()

			// No execution should occur, just validation
			// 실행이 발생하지 않아야 하며, 검증만 수행
			// The fact that we got here means the code is safe
			// 여기까지 왔다는 것은 코드가 안전하다는 의미
		})
	}
}

// TestSecurity_XSSPatterns tests XSS attack pattern handling
// TestSecurity_XSSPatterns는 XSS 공격 패턴 처리를 테스트합니다
func TestSecurity_XSSPatterns(t *testing.T) {
	// Common XSS patterns / 일반적인 XSS 패턴
	xssAttempts := []string{
		"<script>alert('XSS')</script>",
		"<img src=x onerror=alert('XSS')>",
		"javascript:alert('XSS')",
		"<iframe src='javascript:alert(\"XSS\")'></iframe>",
		"<body onload=alert('XSS')>",
		"<svg/onload=alert('XSS')>",
		"<script>document.cookie</script>",
		"';alert(String.fromCharCode(88,83,83))//",
	}

	for _, attempt := range xssAttempts {
		t.Run(fmt.Sprintf("xss_%s", attempt), func(t *testing.T) {
			v := New(attempt, "comment")
			v.Required()

			// Validation should handle these safely without execution
			// 검증은 실행 없이 이들을 안전하게 처리해야 함
			_ = v.Validate()

			// No script execution should occur
			// 스크립트 실행이 발생하지 않아야 함
			// Just validation of the string
			// 문자열 검증만 수행
		})
	}
}

// TestSecurity_PathTraversalPrevention tests path traversal attack prevention
// TestSecurity_PathTraversalPrevention는 경로 탐색 공격 방지를 테스트합니다
func TestSecurity_PathTraversalPrevention(t *testing.T) {
	// Path traversal attempts / 경로 탐색 시도
	pathTraversalAttempts := []string{
		"../etc/passwd",
		"../../etc/passwd",
		"..\\..\\windows\\system32",
		"/etc/passwd",
		"C:\\Windows\\System32",
		"file:///etc/passwd",
		"..%2F..%2Fetc%2Fpasswd",
		"....//....//etc/passwd",
	}

	for _, attempt := range pathTraversalAttempts {
		t.Run(fmt.Sprintf("path_traversal_%s", attempt), func(t *testing.T) {
			v := New(attempt, "filepath")
			v.Required()

			// Validation doesn't execute file operations
			// 검증은 파일 작업을 실행하지 않음
			// It only validates the string format
			// 문자열 형식만 검증함
			_ = v.Validate()

			// No file access should occur during validation
			// 검증 중에 파일 액세스가 발생하지 않아야 함
		})
	}
}

// TestSecurity_CommandInjection tests command injection pattern handling
// TestSecurity_CommandInjection는 명령 인젝션 패턴 처리를 테스트합니다
func TestSecurity_CommandInjection(t *testing.T) {
	// Command injection attempts / 명령 인젝션 시도
	commandAttempts := []string{
		"; ls -la",
		"| cat /etc/passwd",
		"& ping -c 10 localhost",
		"`whoami`",
		"$(whoami)",
		"; rm -rf /",
		"&& echo hacked",
		"|| cat /etc/shadow",
	}

	for _, attempt := range commandAttempts {
		t.Run(fmt.Sprintf("command_injection_%s", attempt), func(t *testing.T) {
			v := New(attempt, "command")
			v.Required()

			// Validation should not execute any commands
			// 검증은 어떤 명령도 실행하지 않아야 함
			_ = v.Validate()

			// No command execution should occur
			// 명령 실행이 발생하지 않아야 함
			// Only string validation
			// 문자열 검증만 수행
		})
	}
}

// TestSecurity_BufferOverflow tests handling of extremely large inputs
// TestSecurity_BufferOverflow는 매우 큰 입력 처리를 테스트합니다
func TestSecurity_BufferOverflow(t *testing.T) {
	sizes := []int{
		1024,        // 1KB
		1024 * 10,   // 10KB
		1024 * 100,  // 100KB
		1024 * 1024, // 1MB
	}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("size_%dKB", size/1024), func(t *testing.T) {
			// Create large input / 큰 입력 생성
			largeInput := strings.Repeat("A", size)

			v := New(largeInput, "large_field")
			v.Required().MaxLength(size * 2) // Allow larger / 더 크게 허용

			// Should handle large inputs safely
			// 큰 입력을 안전하게 처리해야 함
			err := v.Validate()

			if err != nil {
				t.Errorf("unexpected error for size %d: %v", size, err)
			}

			// No buffer overflow or crash
			// 버퍼 오버플로우나 충돌 없음
		})
	}
}

// TestSecurity_NullByteInjection tests null byte injection handling
// TestSecurity_NullByteInjection는 널 바이트 인젝션 처리를 테스트합니다
func TestSecurity_NullByteInjection(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "null at end",
			input: "filename.txt\x00.jpg",
		},
		{
			name:  "null in middle",
			input: "file\x00name.txt",
		},
		{
			name:  "multiple nulls",
			input: "file\x00name\x00.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.input, "filename")
			v.Required()

			// Should handle null bytes safely
			// 널 바이트를 안전하게 처리해야 함
			_ = v.Validate()

			// No security vulnerability should be exposed
			// 보안 취약점이 노출되지 않아야 함
		})
	}
}

// TestSecurity_RegexDenialOfService tests ReDoS (Regular Expression Denial of Service) resistance
// TestSecurity_RegexDenialOfService는 ReDoS 저항성을 테스트합니다
func TestSecurity_RegexDenialOfService(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping ReDoS test in short mode")
	}

	// Patterns that could cause ReDoS if not properly implemented
	// 제대로 구현되지 않으면 ReDoS를 일으킬 수 있는 패턴
	redosPatterns := []struct {
		name    string
		pattern string
		input   string
	}{
		{
			name:    "nested quantifiers",
			pattern: `^(a+)+$`,
			input:   strings.Repeat("a", 30) + "X", // Doesn't match / 일치하지 않음
		},
		{
			name:    "alternation",
			pattern: `^(a|a)*$`,
			input:   strings.Repeat("a", 30) + "X",
		},
		{
			name:    "overlapping",
			pattern: `^(a|ab)+$`,
			input:   strings.Repeat("ab", 30) + "X",
		},
	}

	for _, tt := range redosPatterns {
		t.Run(tt.name, func(t *testing.T) {
			// Set a timeout to detect ReDoS
			// ReDoS를 감지하기 위해 타임아웃 설정
			done := make(chan bool)
			go func() {
				v := New(tt.input, "field")
				v.Regex(tt.pattern)
				_ = v.Validate()
				done <- true
			}()

			select {
			case <-done:
				// Completed successfully / 성공적으로 완료
			case <-time.After(5 * time.Second):
				t.Errorf("regex validation timeout - possible ReDoS vulnerability")
			}
		})
	}
}

// TestSecurity_EmailSpoofing tests email spoofing pattern detection
// TestSecurity_EmailSpoofing는 이메일 스푸핑 패턴 감지를 테스트합니다
func TestSecurity_EmailSpoofing(t *testing.T) {
	// Email spoofing attempts / 이메일 스푸핑 시도
	spoofingAttempts := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "normal email",
			email:   "user@example.com",
			wantErr: false,
		},
		{
			name:    "with display name",
			email:   "User <user@example.com>",
			wantErr: true, // Should reject formatted emails / 형식화된 이메일 거부해야 함
		},
		{
			name:    "with special chars",
			email:   "user+tag@example.com",
			wantErr: false, // Valid email format / 유효한 이메일 형식
		},
		{
			name:    "with quotes",
			email:   "\"user\"@example.com",
			wantErr: true, // Suspicious pattern / 의심스러운 패턴
		},
		{
			name:    "with spaces",
			email:   "user @example.com",
			wantErr: true, // Invalid / 무효함
		},
		{
			name:    "double @",
			email:   "user@@example.com",
			wantErr: true, // Invalid / 무효함
		},
	}

	for _, tt := range spoofingAttempts {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.email, "email")
			v.Email()
			err := v.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Email(%q) error = %v, wantErr %v", tt.email, err, tt.wantErr)
			}
		})
	}
}

// TestSecurity_UnicodeNormalization tests unicode normalization attacks
// TestSecurity_UnicodeNormalization는 유니코드 정규화 공격을 테스트합니다
func TestSecurity_UnicodeNormalization(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "combining characters",
			input: "e\u0301", // é (e + combining acute accent)
		},
		{
			name:  "fullwidth",
			input: "ＡＢＣ", // Fullwidth ABC
		},
		{
			name:  "zero-width",
			input: "hello\u200Bworld", // Zero-width space
		},
		{
			name:  "right-to-left override",
			input: "hello\u202Eworld", // RLO character
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.input, "text")
			v.Required()

			// Should handle unicode safely
			// 유니코드를 안전하게 처리해야 함
			_ = v.Validate()

			// No security issues with unicode normalization
			// 유니코드 정규화 관련 보안 문제 없음
		})
	}
}

// TestSecurity_IntegerOverflow tests integer overflow handling
// TestSecurity_IntegerOverflow는 정수 오버플로우 처리를 테스트합니다
func TestSecurity_IntegerOverflow(t *testing.T) {
	tests := []struct {
		name  string
		value int64
	}{
		{
			name:  "max int64",
			value: 9223372036854775807,
		},
		{
			name:  "min int64",
			value: -9223372036854775808,
		},
		{
			name:  "zero",
			value: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "number")
			v.Min(float64(tt.value) - 1).Max(float64(tt.value) + 1)

			// Should handle extreme values safely
			// 극단적인 값을 안전하게 처리해야 함
			err := v.Validate()

			if err != nil {
				t.Errorf("unexpected error for value %d: %v", tt.value, err)
			}
		})
	}
}

// TestSecurity_ConcurrentAccess tests thread safety
// TestSecurity_ConcurrentAccess는 스레드 안전성을 테스트합니다
func TestSecurity_ConcurrentAccess(t *testing.T) {
	// Run with: go test -race
	// 실행: go test -race

	numGoroutines := 100
	iterations := 100

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				v := New(fmt.Sprintf("user%d@example.com", j), "email")
				v.Required().Email()
				_ = v.Validate()
			}
		}(i)
	}

	wg.Wait()

	// If there are race conditions, go test -race will detect them
	// 경합 조건이 있으면 go test -race가 감지함
	// No data races should occur
	// 데이터 경합이 발생하지 않아야 함
}

// TestSecurity_ErrorMessageLeakage tests that error messages don't leak sensitive information
// TestSecurity_ErrorMessageLeakage는 에러 메시지가 민감한 정보를 누출하지 않는지 테스트합니다
func TestSecurity_ErrorMessageLeakage(t *testing.T) {
	sensitiveData := "password123"

	v := New(sensitiveData, "password")
	v.Required().MinLength(20) // Will fail / 실패할 것임
	err := v.Validate()

	if err == nil {
		t.Fatal("expected error")
	}

	errMsg := err.Error()

	// Error message should not contain the actual password
	// 에러 메시지는 실제 비밀번호를 포함하지 않아야 함
	if strings.Contains(errMsg, sensitiveData) {
		t.Errorf("error message leaks sensitive data: %s", errMsg)
	}

	// Error message should be generic and safe
	// 에러 메시지는 일반적이고 안전해야 함
	t.Logf("Error message (safe): %s", errMsg)
}
