package errorutil

import (
	"errors"
	"testing"
)

// TestWrappedError tests the wrappedError type
// TestWrappedError는 wrappedError 타입을 테스트합니다
func TestWrappedError(t *testing.T) {
	tests := []struct {
		name      string
		msg       string
		cause     error
		wantMsg   string
		wantCause error
	}{
		{
			name:      "error without cause",
			msg:       "test error",
			cause:     nil,
			wantMsg:   "test error",
			wantCause: nil,
		},
		{
			name:      "error with cause",
			msg:       "wrapper",
			cause:     errors.New("original"),
			wantMsg:   "wrapper: original",
			wantCause: errors.New("original"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &wrappedError{
				msg:   tt.msg,
				cause: tt.cause,
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test Unwrap() method / Unwrap() 메서드 테스트
			if got := err.Unwrap(); got != tt.wantCause {
				if got == nil || tt.wantCause == nil {
					t.Errorf("Unwrap() = %v, want %v", got, tt.wantCause)
				} else if got.Error() != tt.wantCause.Error() {
					t.Errorf("Unwrap() = %v, want %v", got, tt.wantCause)
				}
			}
		})
	}
}

// TestCodedError tests the codedError type
// TestCodedError는 codedError 타입을 테스트합니다
func TestCodedError(t *testing.T) {
	tests := []struct {
		name      string
		msg       string
		code      string
		cause     error
		wantMsg   string
		wantCode  string
		wantCause error
	}{
		{
			name:      "coded error without cause",
			msg:       "validation failed",
			code:      "ERR001",
			cause:     nil,
			wantMsg:   "[ERR001] validation failed",
			wantCode:  "ERR001",
			wantCause: nil,
		},
		{
			name:      "coded error with cause",
			msg:       "database error",
			code:      "DB500",
			cause:     errors.New("connection failed"),
			wantMsg:   "[DB500] database error: connection failed",
			wantCode:  "DB500",
			wantCause: errors.New("connection failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &codedError{
				msg:   tt.msg,
				code:  tt.code,
				cause: tt.cause,
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test Code() method / Code() 메서드 테스트
			if got := err.Code(); got != tt.wantCode {
				t.Errorf("Code() = %q, want %q", got, tt.wantCode)
			}

			// Test Unwrap() method / Unwrap() 메서드 테스트
			if got := err.Unwrap(); got != tt.wantCause {
				if got == nil || tt.wantCause == nil {
					t.Errorf("Unwrap() = %v, want %v", got, tt.wantCause)
				} else if got.Error() != tt.wantCause.Error() {
					t.Errorf("Unwrap() = %v, want %v", got, tt.wantCause)
				}
			}
		})
	}
}

// TestNumericCodedError tests the numericCodedError type
// TestNumericCodedError는 numericCodedError 타입을 테스트합니다
func TestNumericCodedError(t *testing.T) {
	tests := []struct {
		name      string
		msg       string
		code      int
		cause     error
		wantMsg   string
		wantCode  int
		wantCause error
	}{
		{
			name:      "numeric coded error without cause",
			msg:       "not found",
			code:      404,
			cause:     nil,
			wantMsg:   "[404] not found",
			wantCode:  404,
			wantCause: nil,
		},
		{
			name:      "numeric coded error with cause",
			msg:       "internal error",
			code:      500,
			cause:     errors.New("database failure"),
			wantMsg:   "[500] internal error: database failure",
			wantCode:  500,
			wantCause: errors.New("database failure"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &numericCodedError{
				msg:   tt.msg,
				code:  tt.code,
				cause: tt.cause,
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test Code() method / Code() 메서드 테스트
			if got := err.Code(); got != tt.wantCode {
				t.Errorf("Code() = %d, want %d", got, tt.wantCode)
			}

			// Test Unwrap() method / Unwrap() 메서드 테스트
			if got := err.Unwrap(); got != tt.wantCause {
				if got == nil || tt.wantCause == nil {
					t.Errorf("Unwrap() = %v, want %v", got, tt.wantCause)
				} else if got.Error() != tt.wantCause.Error() {
					t.Errorf("Unwrap() = %v, want %v", got, tt.wantCause)
				}
			}
		})
	}
}

// TestStackError tests the stackError type
// TestStackError는 stackError 타입을 테스트합니다
func TestStackError(t *testing.T) {
	stack := []Frame{
		{File: "test.go", Line: 10, Function: "TestFunc"},
		{File: "main.go", Line: 20, Function: "main"},
	}

	tests := []struct {
		name       string
		msg        string
		stack      []Frame
		cause      error
		wantMsg    string
		wantStack  []Frame
		wantCause  error
	}{
		{
			name:       "stack error without cause",
			msg:        "error occurred",
			stack:      stack,
			cause:      nil,
			wantMsg:    "error occurred",
			wantStack:  stack,
			wantCause:  nil,
		},
		{
			name:       "stack error with cause",
			msg:        "wrapped error",
			stack:      stack,
			cause:      errors.New("original error"),
			wantMsg:    "wrapped error: original error",
			wantStack:  stack,
			wantCause:  errors.New("original error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &stackError{
				msg:   tt.msg,
				stack: tt.stack,
				cause: tt.cause,
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test StackTrace() method / StackTrace() 메서드 테스트
			got := err.StackTrace()
			if len(got) != len(tt.wantStack) {
				t.Errorf("StackTrace() length = %d, want %d", len(got), len(tt.wantStack))
			}
			for i := range got {
				if got[i] != tt.wantStack[i] {
					t.Errorf("StackTrace()[%d] = %v, want %v", i, got[i], tt.wantStack[i])
				}
			}

			// Test Unwrap() method / Unwrap() 메서드 테스트
			if gotErr := err.Unwrap(); gotErr != tt.wantCause {
				if gotErr == nil || tt.wantCause == nil {
					t.Errorf("Unwrap() = %v, want %v", gotErr, tt.wantCause)
				} else if gotErr.Error() != tt.wantCause.Error() {
					t.Errorf("Unwrap() = %v, want %v", gotErr, tt.wantCause)
				}
			}
		})
	}
}

// TestContextError tests the contextError type
// TestContextError는 contextError 타입을 테스트합니다
func TestContextError(t *testing.T) {
	ctx := map[string]interface{}{
		"user_id": 123,
		"action":  "login",
	}

	tests := []struct {
		name      string
		msg       string
		ctx       map[string]interface{}
		cause     error
		wantMsg   string
		wantCtx   map[string]interface{}
		wantCause error
	}{
		{
			name:      "context error without cause",
			msg:       "operation failed",
			ctx:       ctx,
			cause:     nil,
			wantMsg:   "operation failed",
			wantCtx:   ctx,
			wantCause: nil,
		},
		{
			name:      "context error with cause",
			msg:       "wrapped operation failed",
			ctx:       ctx,
			cause:     errors.New("database error"),
			wantMsg:   "wrapped operation failed: database error",
			wantCtx:   ctx,
			wantCause: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &contextError{
				msg:   tt.msg,
				ctx:   tt.ctx,
				cause: tt.cause,
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test Context() method / Context() 메서드 테스트
			got := err.Context()
			if len(got) != len(tt.wantCtx) {
				t.Errorf("Context() length = %d, want %d", len(got), len(tt.wantCtx))
			}
			for k, v := range tt.wantCtx {
				if got[k] != v {
					t.Errorf("Context()[%s] = %v, want %v", k, got[k], v)
				}
			}

			// Test that Context() returns a copy (immutability)
			// Context()가 복사본을 반환하는지 테스트 (불변성)
			got["new_key"] = "new_value"
			got2 := err.Context()
			if _, exists := got2["new_key"]; exists {
				t.Error("Context() should return a copy, not the original map")
			}

			// Test Unwrap() method / Unwrap() 메서드 테스트
			if gotErr := err.Unwrap(); gotErr != tt.wantCause {
				if gotErr == nil || tt.wantCause == nil {
					t.Errorf("Unwrap() = %v, want %v", gotErr, tt.wantCause)
				} else if gotErr.Error() != tt.wantCause.Error() {
					t.Errorf("Unwrap() = %v, want %v", gotErr, tt.wantCause)
				}
			}
		})
	}
}

// TestCompositeError tests the compositeError type
// TestCompositeError는 compositeError 타입을 테스트합니다
func TestCompositeError(t *testing.T) {
	stack := []Frame{
		{File: "test.go", Line: 10, Function: "TestFunc"},
	}
	ctx := map[string]interface{}{
		"user_id": 456,
	}

	tests := []struct {
		name         string
		msg          string
		code         string
		numCode      int
		stack        []Frame
		ctx          map[string]interface{}
		cause        error
		wantMsg      string
		wantCode     string
		wantNumCode  int
		wantStack    []Frame
		wantCtx      map[string]interface{}
		wantCause    error
	}{
		{
			name:         "composite with string code",
			msg:          "validation error",
			code:         "VAL001",
			numCode:      0,
			stack:        stack,
			ctx:          ctx,
			cause:        nil,
			wantMsg:      "[VAL001] validation error",
			wantCode:     "VAL001",
			wantNumCode:  0,
			wantStack:    stack,
			wantCtx:      ctx,
			wantCause:    nil,
		},
		{
			name:         "composite with numeric code",
			msg:          "not found",
			code:         "",
			numCode:      404,
			stack:        stack,
			ctx:          ctx,
			cause:        errors.New("resource missing"),
			wantMsg:      "[404] not found: resource missing",
			wantCode:     "",
			wantNumCode:  404,
			wantStack:    stack,
			wantCtx:      ctx,
			wantCause:    errors.New("resource missing"),
		},
		{
			name:         "composite without codes",
			msg:          "generic error",
			code:         "",
			numCode:      0,
			stack:        stack,
			ctx:          ctx,
			cause:        nil,
			wantMsg:      "generic error",
			wantCode:     "",
			wantNumCode:  0,
			wantStack:    stack,
			wantCtx:      ctx,
			wantCause:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &compositeError{
				msg:     tt.msg,
				code:    tt.code,
				numCode: tt.numCode,
				stack:   tt.stack,
				ctx:     tt.ctx,
				cause:   tt.cause,
			}

			// Test Error() method / Error() 메서드 테스트
			if got := err.Error(); got != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", got, tt.wantMsg)
			}

			// Test Code() method / Code() 메서드 테스트
			if got := err.Code(); got != tt.wantCode {
				t.Errorf("Code() = %q, want %q", got, tt.wantCode)
			}

			// Test NumericCode() method / NumericCode() 메서드 테스트
			if got := err.NumericCode(); got != tt.wantNumCode {
				t.Errorf("NumericCode() = %d, want %d", got, tt.wantNumCode)
			}

			// Test StackTrace() method / StackTrace() 메서드 테스트
			gotStack := err.StackTrace()
			if len(gotStack) != len(tt.wantStack) {
				t.Errorf("StackTrace() length = %d, want %d", len(gotStack), len(tt.wantStack))
			}

			// Test Context() method / Context() 메서드 테스트
			gotCtx := err.Context()
			if len(gotCtx) != len(tt.wantCtx) {
				t.Errorf("Context() length = %d, want %d", len(gotCtx), len(tt.wantCtx))
			}
			for k, v := range tt.wantCtx {
				if gotCtx[k] != v {
					t.Errorf("Context()[%s] = %v, want %v", k, gotCtx[k], v)
				}
			}

			// Test Unwrap() method / Unwrap() 메서드 테스트
			if gotErr := err.Unwrap(); gotErr != tt.wantCause {
				if gotErr == nil || tt.wantCause == nil {
					t.Errorf("Unwrap() = %v, want %v", gotErr, tt.wantCause)
				} else if gotErr.Error() != tt.wantCause.Error() {
					t.Errorf("Unwrap() = %v, want %v", gotErr, tt.wantCause)
				}
			}
		})
	}
}

// TestFrame tests the Frame type
// TestFrame는 Frame 타입을 테스트합니다
func TestFrame(t *testing.T) {
	tests := []struct {
		name     string
		frame    Frame
		wantStr  string
	}{
		{
			name: "normal frame",
			frame: Frame{
				File:     "/path/to/file.go",
				Line:     42,
				Function: "github.com/user/pkg.Function",
			},
			wantStr: "/path/to/file.go:42 github.com/user/pkg.Function",
		},
		{
			name: "empty frame",
			frame: Frame{
				File:     "",
				Line:     0,
				Function: "",
			},
			wantStr: ":0 ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.frame.String(); got != tt.wantStr {
				t.Errorf("Frame.String() = %q, want %q", got, tt.wantStr)
			}
		})
	}
}
