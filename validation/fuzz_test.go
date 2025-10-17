package validation

import (
	"testing"
)

// FuzzEmail tests the Email validator with random inputs
// FuzzEmail은 랜덤 입력으로 Email 검증기를 테스트합니다
func FuzzEmail(f *testing.F) {
	// Seed corpus with valid and invalid emails
	// 유효한 이메일과 유효하지 않은 이메일로 시드 코퍼스 생성
	testcases := []string{
		"test@example.com",
		"user.name+tag@example.co.uk",
		"invalid",
		"@example.com",
		"test@",
		"",
		"test@@example.com",
		"test@example",
	}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, email string) {
		// Should not panic on any input
		// 어떤 입력에도 패닉하지 않아야 함
		v := New(email, "email")
		v.Email()
		_ = v.Validate()
	})
}

// FuzzURL tests the URL validator with random inputs
// FuzzURL은 랜덤 입력으로 URL 검증기를 테스트합니다
func FuzzURL(f *testing.F) {
	// Seed corpus
	testcases := []string{
		"https://example.com",
		"http://localhost:8080",
		"ftp://files.example.com",
		"invalid-url",
		"",
		"://missing-protocol",
		"http://",
	}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, url string) {
		// Should not panic
		v := New(url, "url")
		v.URL()
		_ = v.Validate()
	})
}

// FuzzMinLength tests MinLength validator with random inputs
// FuzzMinLength는 랜덤 입력으로 MinLength 검증기를 테스트합니다
func FuzzMinLength(f *testing.F) {
	testcases := []struct {
		str    string
		minLen int
	}{
		{"hello", 3},
		{"", 0},
		{"test", 10},
		{"안녕하세요", 3},
		{"Hello🎉", 5},
	}

	for _, tc := range testcases {
		f.Add(tc.str, tc.minLen)
	}

	f.Fuzz(func(t *testing.T, str string, minLen int) {
		// Avoid negative lengths which would cause expected errors
		if minLen < 0 {
			return
		}

		// Should not panic
		v := New(str, "field")
		v.MinLength(minLen)
		_ = v.Validate()
	})
}

// FuzzMaxLength tests MaxLength validator with random inputs
// FuzzMaxLength는 랜덤 입력으로 MaxLength 검증기를 테스트합니다
func FuzzMaxLength(f *testing.F) {
	testcases := []struct {
		str    string
		maxLen int
	}{
		{"hello", 10},
		{"test", 3},
		{"안녕하세요", 10},
		{"", 0},
	}

	for _, tc := range testcases {
		f.Add(tc.str, tc.maxLen)
	}

	f.Fuzz(func(t *testing.T, str string, maxLen int) {
		if maxLen < 0 {
			return
		}

		// Should not panic
		v := New(str, "field")
		v.MaxLength(maxLen)
		_ = v.Validate()
	})
}

// FuzzRegex tests the Matches validator with random inputs
// FuzzRegex는 랜덤 입력으로 Matches 검증기를 테스트합니다
func FuzzRegex(f *testing.F) {
	testcases := []struct {
		str     string
		pattern string
	}{
		{"test123", `^[a-z]+\d+$`},
		{"hello", `[a-z]+`},
		{"", `.*`},
		{"test", `[`}, // Invalid regex
	}

	for _, tc := range testcases {
		f.Add(tc.str, tc.pattern)
	}

	f.Fuzz(func(t *testing.T, str, pattern string) {
		// Should not panic even with invalid regex
		// 유효하지 않은 정규식이라도 패닉하지 않아야 함
		v := New(str, "field")
		v.Regex(pattern)
		_ = v.Validate()
	})
}

// FuzzAlpha tests Alpha validator with random inputs
// FuzzAlpha는 랜덤 입력으로 Alpha 검증기를 테스트합니다
func FuzzAlpha(f *testing.F) {
	testcases := []string{
		"hello",
		"WORLD",
		"Hello123",
		"안녕하세요",
		"Test!@#",
		"",
		"  spaces  ",
	}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, str string) {
		// Should not panic
		v := New(str, "field")
		v.Alpha()
		_ = v.Validate()
	})
}

// FuzzAlphaNumeric tests AlphaNumeric validator
// FuzzAlphaNumeric은 AlphaNumeric 검증기를 테스트합니다
func FuzzAlphaNumeric(f *testing.F) {
	testcases := []string{
		"hello123",
		"Test456",
		"!@#$%",
		"안녕123",
		"",
	}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, str string) {
		v := New(str, "field")
		v.Alphanumeric()
		_ = v.Validate()
	})
}

// FuzzNumeric tests Numeric validator
// FuzzNumeric은 Numeric 검증기를 테스트합니다
func FuzzNumeric(f *testing.F) {
	testcases := []string{
		"12345",
		"0",
		"-123",
		"12.34",
		"abc",
		"",
	}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, str string) {
		v := New(str, "field")
		v.Numeric()
		_ = v.Validate()
	})
}
