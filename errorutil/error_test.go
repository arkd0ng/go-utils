package errorutil

import (
	"testing"
)

// TestNew tests the New function
// TestNew는 New 함수를 테스트합니다
func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
	}{
		{
			name:    "simple message",
			message: "something went wrong",
			want:    "something went wrong",
		},
		{
			name:    "empty message",
			message: "",
			want:    "",
		},
		{
			name:    "message with special characters",
			message: "error: failed! @#$%",
			want:    "error: failed! @#$%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New(tt.message)
			if err == nil {
				t.Fatal("New() returned nil, want error")
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("New() = %q, want %q", got, tt.want)
			}

			// Verify that the error is a wrappedError type
			// 에러가 wrappedError 타입인지 확인
			if _, ok := err.(*wrappedError); !ok {
				t.Errorf("New() returned %T, want *wrappedError", err)
			}

			// Verify Unwrap returns nil (no cause)
			// Unwrap이 nil을 반환하는지 확인 (cause 없음)
			if unwrapper, ok := err.(Unwrapper); ok {
				if cause := unwrapper.Unwrap(); cause != nil {
					t.Errorf("Unwrap() = %v, want nil", cause)
				}
			}
		})
	}
}

// TestNewf tests the Newf function
// TestNewf는 Newf 함수를 테스트합니다
func TestNewf(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
		want   string
	}{
		{
			name:   "simple format",
			format: "error: %s",
			args:   []interface{}{"test"},
			want:   "error: test",
		},
		{
			name:   "multiple arguments",
			format: "failed to process user %d: %s",
			args:   []interface{}{123, "invalid data"},
			want:   "failed to process user 123: invalid data",
		},
		{
			name:   "no arguments",
			format: "simple error",
			args:   []interface{}{},
			want:   "simple error",
		},
		{
			name:   "integer format",
			format: "error code: %d",
			args:   []interface{}{500},
			want:   "error code: 500",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Newf(tt.format, tt.args...)
			if err == nil {
				t.Fatal("Newf() returned nil, want error")
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Newf() = %q, want %q", got, tt.want)
			}

			// Verify that the error is a wrappedError type
			// 에러가 wrappedError 타입인지 확인
			if _, ok := err.(*wrappedError); !ok {
				t.Errorf("Newf() returned %T, want *wrappedError", err)
			}
		})
	}
}

// TestWithCode tests the WithCode function
// TestWithCode는 WithCode 함수를 테스트합니다
func TestWithCode(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		message  string
		wantMsg  string
		wantCode string
	}{
		{
			name:     "standard error code",
			code:     "ERR001",
			message:  "invalid input",
			wantMsg:  "[ERR001] invalid input",
			wantCode: "ERR001",
		},
		{
			name:     "descriptive error code",
			code:     "VALIDATION_ERROR",
			message:  "field is required",
			wantMsg:  "[VALIDATION_ERROR] field is required",
			wantCode: "VALIDATION_ERROR",
		},
		{
			name:     "empty code",
			code:     "",
			message:  "test error",
			wantMsg:  "[] test error",
			wantCode: "",
		},
		{
			name:     "empty message",
			code:     "TEST001",
			message:  "",
			wantMsg:  "[TEST001] ",
			wantCode: "TEST001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WithCode(tt.code, tt.message)
			if err == nil {
				t.Fatal("WithCode() returned nil, want error")
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test Code() method / Code() 메서드 테스트
			if coder, ok := err.(Coder); ok {
				if got := coder.Code(); got != tt.wantCode {
					t.Errorf("Code() = %q, want %q", got, tt.wantCode)
				}
			} else {
				t.Error("WithCode() did not return a Coder")
			}

			// Verify that the error is a codedError type
			// 에러가 codedError 타입인지 확인
			if _, ok := err.(*codedError); !ok {
				t.Errorf("WithCode() returned %T, want *codedError", err)
			}
		})
	}
}

// TestWithCodef tests the WithCodef function
// TestWithCodef는 WithCodef 함수를 테스트합니다
func TestWithCodef(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		format   string
		args     []interface{}
		wantMsg  string
		wantCode string
	}{
		{
			name:     "with single argument",
			code:     "ERR001",
			format:   "invalid user: %d",
			args:     []interface{}{123},
			wantMsg:  "[ERR001] invalid user: 123",
			wantCode: "ERR001",
		},
		{
			name:     "with multiple arguments",
			code:     "DB_ERROR",
			format:   "query failed: %s (%d rows)",
			args:     []interface{}{"timeout", 0},
			wantMsg:  "[DB_ERROR] query failed: timeout (0 rows)",
			wantCode: "DB_ERROR",
		},
		{
			name:     "no arguments",
			code:     "SIMPLE",
			format:   "simple error",
			args:     []interface{}{},
			wantMsg:  "[SIMPLE] simple error",
			wantCode: "SIMPLE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WithCodef(tt.code, tt.format, tt.args...)
			if err == nil {
				t.Fatal("WithCodef() returned nil, want error")
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test Code() method / Code() 메서드 테스트
			if coder, ok := err.(Coder); ok {
				if got := coder.Code(); got != tt.wantCode {
					t.Errorf("Code() = %q, want %q", got, tt.wantCode)
				}
			} else {
				t.Error("WithCodef() did not return a Coder")
			}
		})
	}
}

// TestWithNumericCode tests the WithNumericCode function
// TestWithNumericCode는 WithNumericCode 함수를 테스트합니다
func TestWithNumericCode(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		message  string
		wantMsg  string
		wantCode int
	}{
		{
			name:     "HTTP 404 error",
			code:     404,
			message:  "user not found",
			wantMsg:  "[404] user not found",
			wantCode: 404,
		},
		{
			name:     "HTTP 500 error",
			code:     500,
			message:  "internal server error",
			wantMsg:  "[500] internal server error",
			wantCode: 500,
		},
		{
			name:     "zero code",
			code:     0,
			message:  "test error",
			wantMsg:  "[0] test error",
			wantCode: 0,
		},
		{
			name:     "negative code",
			code:     -1,
			message:  "negative code test",
			wantMsg:  "[-1] negative code test",
			wantCode: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WithNumericCode(tt.code, tt.message)
			if err == nil {
				t.Fatal("WithNumericCode() returned nil, want error")
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test Code() method / Code() 메서드 테스트
			if coder, ok := err.(NumericCoder); ok {
				if got := coder.Code(); got != tt.wantCode {
					t.Errorf("Code() = %d, want %d", got, tt.wantCode)
				}
			} else {
				t.Error("WithNumericCode() did not return a NumericCoder")
			}

			// Verify that the error is a numericCodedError type
			// 에러가 numericCodedError 타입인지 확인
			if _, ok := err.(*numericCodedError); !ok {
				t.Errorf("WithNumericCode() returned %T, want *numericCodedError", err)
			}
		})
	}
}

// TestWithNumericCodef tests the WithNumericCodef function
// TestWithNumericCodef는 WithNumericCodef 함수를 테스트합니다
func TestWithNumericCodef(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		format   string
		args     []interface{}
		wantMsg  string
		wantCode int
	}{
		{
			name:     "HTTP 404 with user ID",
			code:     404,
			format:   "user %d not found",
			args:     []interface{}{123},
			wantMsg:  "[404] user 123 not found",
			wantCode: 404,
		},
		{
			name:     "HTTP 500 with error details",
			code:     500,
			format:   "database error: %s",
			args:     []interface{}{"connection timeout"},
			wantMsg:  "[500] database error: connection timeout",
			wantCode: 500,
		},
		{
			name:     "multiple arguments",
			code:     400,
			format:   "validation failed: field %s must be %d characters",
			args:     []interface{}{"username", 8},
			wantMsg:  "[400] validation failed: field username must be 8 characters",
			wantCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WithNumericCodef(tt.code, tt.format, tt.args...)
			if err == nil {
				t.Fatal("WithNumericCodef() returned nil, want error")
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test Code() method / Code() 메서드 테스트
			if coder, ok := err.(NumericCoder); ok {
				if got := coder.Code(); got != tt.wantCode {
					t.Errorf("Code() = %d, want %d", got, tt.wantCode)
				}
			} else {
				t.Error("WithNumericCodef() did not return a NumericCoder")
			}
		})
	}
}
